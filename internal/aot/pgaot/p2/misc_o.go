package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OpenTransientFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_reserveAllocatedDesc(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[0]))
	if v15 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L19
	}
L6:
	;
	v55 = F_BasicOpenFilePerm(m, l0, l1, l2)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L14
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[1]))
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[2]))
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[3]))
	if v21+(v23+v15) < v19 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[4]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	F_LruDelete(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[0]))
	if v38 <= int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[1]))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[2]))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[3]))
	if v42 <= v44+(v46+v38) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	if int32(0) <= v55 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[5]))
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[3]))
	v65 = v60 + v62*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = int32(3)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[6]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	goto L18
L16:
	;
	goto L17
L17:
	;
	m.G0 = v8 + int32(16)
	return v55
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v71
	v73 = int32(_a_F_OpenTransientFilePerm_0)
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[3])) = v75 + int32(1)
	goto L17
L19:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTransientFilePerm[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v93
	F_errmsg(m, int32(_a_F_OpenTransientFilePerm_1), v8)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_OpenTransientFilePerm_2), int32(2720), int32(_a_F_OpenTransientFilePerm_3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_offsethash_create(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13999(m, l0, l1, l2, int32(_a_F_offsethash_create_0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_offsethash_insert_hash_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v36 int32
	_ = v36
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	goto L1
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v26) <= base.Ui32(v25) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L1
L4:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v212 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v211 + v212
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)) = uint8(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = l1
	v218 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v218)
	return v206
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L10
	} else {
		goto L45
	}
L6:
	;
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v28 == int64(4294967296) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v38 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v41 = v40 & l2
	v44 = v39 + v41<<(uint(int32(3))%32)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+4)))
	if v45 == v38 {
		v206 = v44
		goto L4
	} else {
		goto L12
	}
L9:
	;
	F_offsethash_grow(m, l0, v28<<(uint(int64(1))%64))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L8
L12:
	;
	v53 = v38
	v54 = v41
	v55 = v44
	goto L13
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if l1 == v60 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v206 = v184
	goto L4
L15:
	;
	v62 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v62)
	return v55
L16:
	;
	goto L17
L17:
	;
	v66 = v54 + int32(1)
	v67 = int32(16)
	v71 = (int32(base.Ui32(v60)>>(uint(v67)%32)) ^ v60) * int32(-2048144789)
	v76 = (int32(base.Ui32(v71)>>(uint(int32(13))%32)) ^ v71) * int32(-1028477387)
	v80 = v40 & (int32(base.Ui32(v76)>>(uint(v67)%32)) ^ v76)
	if base.Ui32(v54) < base.Ui32(v80) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v84 = v54 + v82
	goto L20
L19:
	;
	v84 = v54
	goto L20
L20:
	;
	if base.Ui32(v84-v80) < base.Ui32(v53) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v88 = v66 & v40
	v91 = v39 + v88<<(uint(int32(3))%32)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+4)))
	if v92 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v171 = v53 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v171) {
		goto L40
	} else {
		goto L41
	}
L24:
	;
	v97 = int32(0)
	v98 = v88
	goto L27
L25:
	;
	v128 = v88
	v129 = v91
	goto L26
L26:
	;
	if v128 != v54 {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	if base.Ui32(int32(150)) <= base.Ui32(v97) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v128 = v118
	v129 = v121
	goto L26
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v107), base.F64_convert_i64_u(v109)), float64(0.1)) != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v114 = int32(1)
	v118 = (v98 + v114) & v40
	v121 = v39 + v118<<(uint(int32(3))%32)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)))
	if v122 != 0 {
		v97 = v97 + v114
		v98 = v118
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	goto L28
L34:
	;
	v141 = v128
	v142 = v129
	goto L37
L35:
	;
	goto L36
L36:
	;
	v206 = v55
	goto L4
L37:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v151 = v148 & (v141 - int32(1))
	v154 = v39 + v151<<(uint(int32(3))%32)
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v154)))
	*(*int64)(unsafe.Add(mBase, uint32(v142))) = v155
	if v151 != v54 {
		v141 = v151
		v142 = v154
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L36
L39:
	;
	goto L38
L40:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v176 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v174), base.F64_convert_i64_u(v176)), float64(0.1)) != 0 {
		goto L3
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v181 = v66 & v40
	v184 = v39 + v181<<(uint(int32(3))%32)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+4)))
	if v185 != 0 {
		v53 = v171
		v54 = v181
		v55 = v184
		goto L13
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	goto L14
L45:
	;
	F_errmsg_internal(m, int32(_a_F_offsethash_insert_hash_internal_0), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_offsethash_insert_hash_internal_1), int32(630), int32(_a_F_offsethash_insert_hash_internal_2))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_oidge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(base.Ui32(v3) <= base.Ui32(v2))
}
func F_oidout(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13954(m, l0, int32(_a_F_oidout_0), int32(12))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_oidsmaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.Ui32(v3) < base.Ui32(v4) {
		v6 = v3
	} else {
		v6 = v4
	}
	return v6
}
func F_oidvectorne(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_btoidvectorcmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v2 != int32(0))
	}
}
func F_oidvectortypes(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_check_valid_oidvector(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v19 = v17 * int32(20)
	v21 = v19 | int32(1)
	v22 = F_palloc(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v24)
	if v17 <= v24 {
		v100 = v22
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v111 = F_cstring_to_text(m, v100)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L22
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v31 = F_format_type_extended(m, v28, int32(-1), int32(2))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v33 = F_strlen(m, v31)
	mBase = m.M
	v35 = v33 + int32(2)
	if base.Ui32(v19) < base.Ui32(v35) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v38 = v35 + v21
	v39 = F_repalloc(m, v22, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v41 = v19
	v42 = v21
	v43 = v22
	goto L9
L9:
	;
	v44 = F_strlen(m, v43)
	mBase = m.M
	v46 = F_strcpy(m, v44+v43, v31)
	mBase = m.M
	goto L11
L10:
	;
	v41 = v19 + v35
	v42 = v38
	v43 = v39
	goto L9
L11:
	;
	if v17 == int32(1) {
		v100 = v43
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oidvectortypes[0])))
	v55 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_oidvectortypes[1])))
	v57 = v43
	v58 = v41 - v33
	v59 = int32(1)
	v61 = v42
	goto L13
L13:
	;
	v68 = int32(2)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(24)+v59<<(uint(v68)%32))))
	v74 = F_format_type_extended(m, v71, int32(-1), v68)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v100 = v84
	goto L4
L15:
	;
	v76 = F_strlen(m, v74)
	mBase = m.M
	v78 = v76 + int32(2)
	if base.Ui32(v58) < base.Ui32(v78) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v81 = v78 + v61
	v82 = F_repalloc(m, v57, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v84 = v57
	v85 = v58
	v86 = v61
	goto L18
L18:
	;
	v87 = F_strlen(m, v84)
	mBase = m.M
	v88 = v87 + v84
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+2)) = uint8(v53)
	*(*uint16)(unsafe.Add(mBase, uint32(v88))) = uint16(v55)
	v94 = F_strlen(m, v84)
	mBase = m.M
	v96 = F_strcpy(m, v94+v84, v74)
	mBase = m.M
	goto L20
L19:
	;
	v84 = v82
	v85 = v58 + v78
	v86 = v81
	goto L18
L20:
	;
	v98 = v59 + int32(1)
	if v98 != v17 {
		v57 = v84
		v58 = v85 - v76 - int32(2)
		v59 = v98
		v61 = v86
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	return v111
}
func F_okeys_array_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	if v8 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_okeys_array_start_0)
				F_errmsg(m, int32(_a_F_okeys_array_start_1), v5)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_okeys_array_start_2), int32(818), int32(_a_F_okeys_array_start_3))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		m.G0 = v5 + int32(16)
		return int32(0)
	}
}
func F_okeys_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v9 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_okeys_scalar_0)
				F_errmsg(m, int32(_a_F_okeys_scalar_1), v6)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_okeys_scalar_2), int32(833), int32(_a_F_okeys_scalar_3))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
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
		m.G0 = v6 + int32(16)
		return int32(0)
	}
}
func F_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1&int32(64) == int32(0) {
		v14 = int32(_a_F_open_0)
		if l1&v14 != v14 {
			v22 = int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2 + int32(4)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v22 = v21
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2 + int32(4)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v22 = v21
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v22
	v27 = m.Env.X__syscall_openat(m, int32(-100), l0, l1|int32(_a_F_open_1), v7)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v27) {
		*(*int32)(unsafe.Add(mBase, _c_F_open[0])) = int32(0) - v27
		v35 = int32(-1)
	} else {
		v35 = v27
	}
	m.G0 = v7 + int32(16)
	return v35
}
func F_open_auth_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	if int32(11) <= l2 {
		v14 = F_errstart(m, l1, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v14 != 0 {
				F_errcode_for_file_access(m)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
					F_errmsg(m, int32(_a_F_open_auth_file_0), v7+int32(-48))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_open_auth_file_1), int32(612), int32(_a_F_open_auth_file_2))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							if l3 == int32(0) {
								v84 = v5
								m.G0 = v9 - int32(-64)
								return v84
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
								v35 = F_psprintf(m, int32(_a_F_open_auth_file_0), v9)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v35
									v84 = v5
									m.G0 = v9 - int32(-64)
									return v84
								}
							}
						}
					}
				}
			} else {
				if l3 == int32(0) {
					v84 = v5
					m.G0 = v9 - int32(-64)
					return v84
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
					v35 = F_psprintf(m, int32(_a_F_open_auth_file_0), v9)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v35
						v84 = v5
						m.G0 = v9 - int32(-64)
						return v84
					}
				}
			}
		}
	} else {
		v39 = F_AllocateFile(m, l0, int32(_a_F_open_auth_file_3))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			if v39 == int32(0) {
				v44 = *(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0]))
				v46 = F_errstart(m, l1, int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					if v46 != 0 {
						F_errcode_for_file_access(m)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l0
							F_errmsg(m, int32(_a_F_open_auth_file_4), v7+int32(-16))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_open_auth_file_1), int32(627), int32(_a_F_open_auth_file_2))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									if l3 != 0 {
										*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0])) = v44
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0
										v67 = F_psprintf(m, int32(_a_F_open_auth_file_4), v7+int32(-32))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v67
											*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0])) = v44
											v84 = int32(0)
											m.G0 = v9 - int32(-64)
											return v84
										}
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0])) = v44
										v84 = int32(0)
										m.G0 = v9 - int32(-64)
										return v84
									}
								}
							}
						}
					} else {
						if l3 != 0 {
							*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0])) = v44
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0
							v67 = F_psprintf(m, int32(_a_F_open_auth_file_4), v7+int32(-32))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v67
								*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0])) = v44
								v84 = int32(0)
								m.G0 = v9 - int32(-64)
								return v84
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[0])) = v44
							v84 = int32(0)
							m.G0 = v9 - int32(-64)
							return v84
						}
					}
				}
			} else {
				if l2 != 0 {
					v84 = v39
					m.G0 = v9 - int32(-64)
					return v84
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[1]))
					v80 = F_AllocSetContextCreateInternal(m, v75, int32(_a_F_open_auth_file_5), int32(0), int32(1024), int32(_a_F_open_auth_file_6))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_open_auth_file[2])) = v80
						v84 = v39
						m.G0 = v9 - int32(-64)
						return v84
					}
				}
			}
		}
	}
}
func F_or_arg_index_match_cmp_group(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v6 < v7 {
		v18 = int32(-1)
	} else {
		if v7 < v6 {
			v18 = int32(1)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			if v12 < v13 {
				v18 = int32(-1)
			} else {
				v18 = base.B2i32(v13 < v12)
			}
		}
	}
	return v18
}
func F_ordered_set_transition_multi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	v2 = int32(0)
	v12 = l0 + int32(20)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v13 == int32(1) {
		v17 = F_ordered_set_startup(m, l0, int32(1))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = v17
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
			m.T0[v26].(func(*base.Module, int32))(m, v24)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
				if v29 < int32(2) {
					v114 = v2
				} else {
					v33 = v29 - int32(1)
					v34 = int32(0)
					if v29 != int32(2) {
						v41 = v34
						v50 = v2
						for {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
							v52 = int32(2)
							v56 = v41 | int32(1)
							v57 = int32(3)
							v59 = v12 + v56<<(uint(v57)%32)
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
							*(*int32)(unsafe.Add(mBase, uint32(v51+v41<<(uint(v52)%32)))) = v60
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
							v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
							*(*uint8)(unsafe.Add(mBase, uint32(v62+v41))) = uint8(v64)
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
							v71 = v41 + v52
							v74 = v12 + v71<<(uint(v57)%32)
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
							*(*int32)(unsafe.Add(mBase, uint32(v66+v56<<(uint(v52)%32)))) = v75
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)))
							*(*uint8)(unsafe.Add(mBase, uint32(v56+v77))) = uint8(v79)
							v82 = v50 + v52
							if v82 != v33&int32(-2) {
								v41 = v71
								v50 = v82
								continue
							} else {
								break
							}
							break
						}
						if v33&int32(1) == int32(0) {
							v114 = v33
						} else {
							v86 = v71
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
							v102 = v12 + v86<<(uint(int32(3))%32)
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v96+v86<<(uint(int32(2))%32)))) = v103
							v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
							v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+12)))
							*(*uint8)(unsafe.Add(mBase, uint32(v105+v86))) = uint8(v107)
							v114 = v33
						}
					} else {
						v86 = v34
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
						v102 = v12 + v86<<(uint(int32(3))%32)
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v96+v86<<(uint(int32(2))%32)))) = v103
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
						v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+12)))
						*(*uint8)(unsafe.Add(mBase, uint32(v105+v86))) = uint8(v107)
						v114 = v33
					}
				}
				v119 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
				v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+50)))
				if v121 == int32(104) {
					v124 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
					v128 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v124+v114<<(uint(int32(2))%32)))) = v128
					v130 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
					*(*uint8)(unsafe.Add(mBase, uint32(v130+v114))) = uint8(v128)
				} else {
				}
				v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
				v136 = v134 & int32(_a_F_ordered_set_transition_multi_0)
				*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)) = uint16(v136)
				v138 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
				v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
				*(*uint16)(unsafe.Add(mBase, uint32(v24)+6)) = uint16(v139)
				v141 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
				F_tuplesort_puttupleslot(m, v141, v24)
				mBase = m.M
				v143 = m.ExcPending
				if v143 != 0 {
					return int32(0)
				} else {
					v144 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v144 + int64(1)
					return v22
				}
			}
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v22 = v21
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
		m.T0[v26].(func(*base.Module, int32))(m, v24)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
			if v29 < int32(2) {
				v114 = v2
			} else {
				v33 = v29 - int32(1)
				v34 = int32(0)
				if v29 != int32(2) {
					v41 = v34
					v50 = v2
					for {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
						v52 = int32(2)
						v56 = v41 | int32(1)
						v57 = int32(3)
						v59 = v12 + v56<<(uint(v57)%32)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
						*(*int32)(unsafe.Add(mBase, uint32(v51+v41<<(uint(v52)%32)))) = v60
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
						*(*uint8)(unsafe.Add(mBase, uint32(v62+v41))) = uint8(v64)
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
						v71 = v41 + v52
						v74 = v12 + v71<<(uint(v57)%32)
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
						*(*int32)(unsafe.Add(mBase, uint32(v66+v56<<(uint(v52)%32)))) = v75
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)))
						*(*uint8)(unsafe.Add(mBase, uint32(v56+v77))) = uint8(v79)
						v82 = v50 + v52
						if v82 != v33&int32(-2) {
							v41 = v71
							v50 = v82
							continue
						} else {
							break
						}
						break
					}
					if v33&int32(1) == int32(0) {
						v114 = v33
					} else {
						v86 = v71
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
						v102 = v12 + v86<<(uint(int32(3))%32)
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v96+v86<<(uint(int32(2))%32)))) = v103
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
						v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+12)))
						*(*uint8)(unsafe.Add(mBase, uint32(v105+v86))) = uint8(v107)
						v114 = v33
					}
				} else {
					v86 = v34
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
					v102 = v12 + v86<<(uint(int32(3))%32)
					v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v96+v86<<(uint(int32(2))%32)))) = v103
					v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
					v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+12)))
					*(*uint8)(unsafe.Add(mBase, uint32(v105+v86))) = uint8(v107)
					v114 = v33
				}
			}
			v119 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
			v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+50)))
			if v121 == int32(104) {
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
				v128 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v124+v114<<(uint(int32(2))%32)))) = v128
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
				*(*uint8)(unsafe.Add(mBase, uint32(v130+v114))) = uint8(v128)
			} else {
			}
			v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
			v136 = v134 & int32(_a_F_ordered_set_transition_multi_0)
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)) = uint16(v136)
			v138 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
			v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
			*(*uint16)(unsafe.Add(mBase, uint32(v24)+6)) = uint16(v139)
			v141 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			F_tuplesort_puttupleslot(m, v141, v24)
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return int32(0)
			} else {
				v144 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v144 + int64(1)
				return v22
			}
		}
	}
}
func F_overlaps_timetz(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13955(m, l0, int32(1272), int32(1271))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
