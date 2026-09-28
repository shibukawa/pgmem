package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PGSemaphoreLock(m *base.Module, l0 int32) {
	var v7 int32
	_ = v7
	Fn14220(m, l0, int32(_a_F_PGSemaphoreLock_0), int32(329), int32(_a_F_PGSemaphoreLock_1), int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_RemovePgTempFiles(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	v5 = m.G0
	v7 = v5 - int32(1120)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_RemovePgTempFiles_0)
	v12 = v7 + int32(48)
	v17 = F_pg_snprintf(m, v12, int32(1060), int32(_a_F_RemovePgTempFiles_1), v7+int32(32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_RemovePgTempFilesInDir(m, v12, int32(1), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_RemovePgTempRelationFiles(m, int32(_a_F_RemovePgTempFiles_2))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = F_AllocateDir(m, int32(_a_F_RemovePgTempFiles_3))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = F_ReadDirExtended(m, v27, int32(_a_F_RemovePgTempFiles_3), int32(15))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v34 = v31
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_FreeDir(m, v27)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L24
	}
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+19)))
	if v37 != int32(46) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L9
L12:
	;
	v85 = F_ReadDirExtended(m, v27, int32(_a_F_RemovePgTempFiles_3), int32(15))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L22
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(_a_F_RemovePgTempFiles_0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(_a_F_RemovePgTempFiles_4)
	v54 = v34 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(_a_F_RemovePgTempFiles_3)
	v59 = v7 + int32(48)
	v64 = F_pg_snprintf(m, v59, int32(1060), int32(_a_F_RemovePgTempFiles_5), v7+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L18
	}
L14:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+20)))
	if v40 == int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+20)))
	if v43 != int32(46) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+21)))
	if v46 == int32(0) {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	F_RemovePgTempFilesInDir(m, v59, int32(1), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(_a_F_RemovePgTempFiles_4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_RemovePgTempFiles_3)
	v77 = F_pg_snprintf(m, v59, int32(1060), int32(_a_F_RemovePgTempFiles_6), v7)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_RemovePgTempRelationFiles(m, v59)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L12
L22:
	;
	if v85 != 0 {
		v34 = v85
		goto L10
	} else {
		goto L23
	}
L23:
	;
	goto L11
L24:
	;
	m.G0 = v7 + int32(1120)
	return
}
func F_pg_any_to_server(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 <= int32(0) {
		v97 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v97
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_pg_any_to_server[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v15 != l2 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v17 = l2
	goto L5
L4:
	;
	v17 = int32(0)
	goto L5
L5:
	;
	if v17 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15*int32(28))+uint32(_c_F_pg_any_to_server[1])))
	v25 = m.T0[v24].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if v15 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	return int32(0)
L10:
	;
	if l1 == v25 {
		v97 = l0
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_report_invalid_encoding(m, v15, l0+v25, l1-v25)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
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
	v36 = int32(0)
	if base.B2i32(l2 == int32(7))|base.B2i32(base.Ui32(int32(34)) < base.Ui32(l2)) == v36 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_pg_any_to_server[2]))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v90 == l2 {
		goto L32
	} else {
		goto L33
	}
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2*int32(28))+uint32(_c_F_pg_any_to_server[1])))
	v49 = m.T0[v48].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v59 = v36
	goto L22
L19:
	;
	if l1 == v49 {
		v97 = l0
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_report_invalid_encoding(m, l2, l0+v49, l1-v49)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	v61 = l0 + v59
	v62 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61))))
	if int32(0) < v62 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L9
	} else {
		goto L28
	}
L24:
	;
	v66 = v59 + int32(1)
	if l1 != v66 {
		v59 = v66
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	v97 = l0
	goto L1
L28:
	;
	F_errcode(m, int32(17301634))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v75
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_pg_any_to_server[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v78
	F_errmsg(m, int32(_a_F_pg_any_to_server_0), v8)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_pg_any_to_server_1), int32(726), int32(_a_F_pg_any_to_server_2))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	v93 = F_perform_default_encoding_conversion(m, l0, l1, int32(1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v95 = F_pg_do_encoding_conversion(m, l0, l1, l2, v15)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L36
	}
L35:
	;
	v97 = v93
	goto L1
L36:
	;
	v97 = v95
	goto L1
}
func F_pg_ascii_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	v3 = int32(0)
	if base.B2i32(l0&int32(3) == v3)|base.B2i32(l1 == v3) != 0 {
		v35 = l0
		v37 = l1
		v38 = base.B2i32(l1 != v3)
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v109 != 0 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	v109 = int32(0)
	goto L1
L3:
	;
	v87 = v80
	v89 = v82
	goto L20
L4:
	;
	if v38 == int32(0) {
		goto L2
	} else {
		goto L11
	}
L5:
	;
	v18 = l0
	v20 = l1
	goto L6
L6:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v23 == int32(0) {
		v80 = v18
		v82 = v20
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v35 = v30
	v37 = v26
	v38 = v28
	goto L4
L8:
	;
	v25 = int32(1)
	v26 = v20 - v25
	v27 = int32(0)
	v28 = base.B2i32(v26 != v27)
	v30 = v18 + v25
	if v30&int32(3) == v27 {
		v35 = v30
		v37 = v26
		v38 = v28
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v26 != 0 {
		v18 = v30
		v20 = v26
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v43 = int32(0)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if base.B2i32(v43 == v44)|base.B2i32(base.Ui32(v37) < base.Ui32(int32(4))) == v43 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v53 = v35
	v55 = v37
	goto L15
L13:
	;
	v73 = v35
	v75 = v37
	goto L14
L14:
	;
	if v75 == int32(0) {
		goto L2
	} else {
		goto L19
	}
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v60 = v59 ^ int32(0)
	v63 = int32(-2139062144)
	if (int32(16843008)-v60|v60)&v63 != v63 {
		v80 = v53
		v82 = v55
		goto L3
	} else {
		goto L17
	}
L16:
	;
	v73 = v68
	v75 = v70
	goto L14
L17:
	;
	v67 = int32(4)
	v68 = v53 + v67
	v70 = v55 - v67
	if base.Ui32(int32(3)) < base.Ui32(v70) {
		v53 = v68
		v55 = v70
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v80 = v73
	v82 = v75
	goto L3
L20:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if int32(0) == v92 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L2
L22:
	;
	v109 = v87
	goto L1
L23:
	;
	goto L24
L24:
	;
	v94 = int32(1)
	v97 = v89 - v94
	if v97 != 0 {
		v87 = v87 + v94
		v89 = v97
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v111 = v109 - l0
	goto L28
L27:
	;
	v111 = l1
	goto L28
L28:
	;
	return v111
}
func F_pg_backend_pid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	v3 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_backend_pid[0])))
	return v3
}
func F_pg_base64_decode_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	if l1 == v5 {
		v202 = l2
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L30
	} else {
		goto L66
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L30
	} else {
		goto L58
	}
L3:
	;
	m.G0 = v15 + int32(48)
	return base.I64_extend_i32_s(v202 - l2)
L4:
	;
	v19 = l0 + l1
	v20 = l0
	v24 = l2
	v25 = v5
	v28 = v5
	v29 = v5
	goto L5
L5:
	;
	v33 = v20
	goto L8
L6:
	;
	v171 = int32(0)
	if base.B2i32(l3 == v171)|base.B2i32(v169 != int32(2)) == v171 {
		goto L51
	} else {
		goto L52
	}
L7:
	;
	goto L6
L8:
	;
	v44 = int32(1)
	v45 = v33 + v44
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v48 = v46 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v48))|base.B2i32(v44<<(uint(v48)%32)&int32(_a_F_pg_base64_decode_internal_0) == int32(0)) != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if l3 != 0 {
		goto L22
	} else {
		goto L23
	}
L10:
	;
	goto L9
L11:
	;
	if base.Ui32(v45) < base.Ui32(v19) {
		v33 = v45
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v166 = v24
	v167 = v25
	v169 = v28
	goto L7
L13:
	;
	if base.Ui32(v45) < base.Ui32(v19) {
		v20 = v45
		v24 = v159
		v25 = v160
		v28 = v162
		v29 = v163
		goto L5
	} else {
		goto L50
	}
L14:
	;
	v159 = v152
	v160 = v153
	v162 = int32(0)
	v163 = v156
	goto L13
L15:
	;
	v152 = v149
	v153 = int32(0)
	v156 = v148
	goto L14
L16:
	;
	v121 = v118 + v25<<(uint(int32(6))%32)
	v123 = v28 + int32(1)
	if v123 != int32(4) {
		goto L41
	} else {
		goto L42
	}
L17:
	;
	v115 = int32(*(*int8)(unsafe.Add(mBase, uint32(v114)+uint32(_c_F_pg_base64_decode_internal[0]))))
	if v115 < int32(0) {
		goto L1
	} else {
		goto L40
	}
L18:
	;
	v114 = int32(47)
	goto L17
L19:
	;
	if base.Ui32((v46-int32(1))&int32(255)) < base.Ui32(int32(126)) {
		v114 = v46
		goto L17
	} else {
		goto L39
	}
L20:
	;
	if v46 == int32(95) {
		goto L18
	} else {
		goto L38
	}
L21:
	;
	if v29 != 0 {
		v118 = int32(0)
		goto L16
	} else {
		goto L26
	}
L22:
	;
	switch v46 - int32(45) {
	case 0:
		v114 = int32(43)
		goto L17
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L19
	case 16:
		goto L21
	default:
		goto L20
	}
L23:
	;
	goto L24
L24:
	;
	if v46 != int32(61) {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	switch v28 - int32(2) {
	case 0:
		goto L29
	case 1:
		goto L28
	default:
		goto L27
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v72 = int32(2)
	v74 = int32(base.Ui32(v25) >> (uint(v72) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)) = uint8(v74)
	v77 = int32(base.Ui32(v25) >> (uint(int32(10)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v77)
	v148 = v72
	v149 = v24 + v72
	goto L15
L29:
	;
	v159 = v24
	v160 = v25 << (uint(int32(6)) % 32)
	v162 = int32(3)
	v163 = int32(1)
	goto L13
L30:
	;
	return int64(0)
L31:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	if l3 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v92 = int32(_a_F_pg_base64_decode_internal_1)
	goto L35
L34:
	;
	v92 = int32(_a_F_pg_base64_decode_internal_2)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v92
	F_errmsg(m, int32(_a_F_pg_base64_decode_internal_3), v15+int32(32))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_pg_base64_decode_internal_4), int32(551), int32(_a_F_pg_base64_decode_internal_5))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	goto L19
L39:
	;
	goto L1
L40:
	;
	v118 = v115
	goto L16
L41:
	;
	v159 = v24
	v160 = v121
	v162 = v123
	v163 = v29
	goto L13
L42:
	;
	goto L43
L43:
	;
	v127 = int32(base.Ui32(v121) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v127)
	v129 = int32(0)
	if v29 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v132 = int32(1)
	v152 = v24 + v132
	v153 = v129
	v156 = v132
	goto L14
L45:
	;
	goto L46
L46:
	;
	v136 = int32(base.Ui32(v121) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)) = uint8(v136)
	if v29 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v152 = v24 + int32(2)
	v153 = v129
	v156 = v29
	goto L14
L48:
	;
	goto L49
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)) = uint8(v121)
	v148 = int32(0)
	v149 = v24 + int32(3)
	goto L15
L50:
	;
	v166 = v159
	v167 = v160
	v169 = v162
	goto L7
L51:
	;
	v179 = int32(base.Ui32(v167) >> (uint(int32(4)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v179)
	v202 = v166 + int32(1)
	goto L3
L52:
	;
	goto L53
L53:
	;
	v183 = int32(0)
	if base.B2i32(l3 == v183)|base.B2i32(v169 != int32(3)) == v183 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v190 = int32(2)
	v191 = int32(base.Ui32(v167) >> (uint(v190) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)) = uint8(v191)
	v194 = int32(base.Ui32(v167) >> (uint(int32(10)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v194)
	v202 = v166 + v190
	goto L3
L55:
	;
	goto L56
L56:
	;
	if v169 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v202 = v166
	goto L3
L58:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L30
	} else {
		goto L59
	}
L59:
	;
	if l3 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v225 = int32(_a_F_pg_base64_decode_internal_1)
	goto L62
L61:
	;
	v225 = int32(_a_F_pg_base64_decode_internal_2)
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v225
	F_errmsg(m, int32(_a_F_pg_base64_decode_internal_6), v15+int32(16))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L30
	} else {
		goto L63
	}
L63:
	;
	F_errhint(m, int32(_a_F_pg_base64_decode_internal_7), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L30
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_pg_base64_decode_internal_4), int32(603), int32(_a_F_pg_base64_decode_internal_5))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L30
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L30
	} else {
		goto L67
	}
L67:
	;
	v249 = F_pg_mblen_range(m, v33, v19)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L30
	} else {
		goto L68
	}
L68:
	;
	if l3 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v253 = int32(_a_F_pg_base64_decode_internal_1)
	goto L71
L70:
	;
	v253 = int32(_a_F_pg_base64_decode_internal_2)
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v249
	F_errmsg(m, int32(_a_F_pg_base64_decode_internal_8), v15)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L30
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_pg_base64_decode_internal_4), int32(568), int32(_a_F_pg_base64_decode_internal_5))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L30
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_base64_enc_len(m *base.Module, l0 int32, l1 int32) int64 {
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v10 int64
	_ = v10
	v4 = base.I32_div_u_s(l1, int32(57))
	v7 = int64(2)
	v10 = base.I64_div_u_s(base.I64_extend_i32_u(l1)+v7, int64(3))
	return base.I64_extend_i32_u(v4) + v10<<(uint(v7)%64)
}
func F_pg_create_physical_replication_slot(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_get_call_result_type(m, l0, int32(0), v8+int32(8))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		if v16 == int32(1) {
			F_CheckSlotPermissions(m)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				F_CheckSlotRequirements(m, int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v28 = int32(0)
					if v10 == int64(0) {
						v33 = v28
					} else {
						v33 = int32(2)
					}
					v34 = int32(0)
					F_ReplicationSlotCreate(m, base.I32_wrap_i64(v12), v28, v33, v34, v34, v34, v34)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						if v11 != int64(0) {
							F_ReplicationSlotReserveWal(m)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								F_ReplicationSlotMarkDirty(m)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int64(0)
								} else {
									F_ReplicationSlotSave(m)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int64(0)
									} else {
										v49 = *(*int32)(unsafe.Add(mBase, _c_F_pg_create_physical_replication_slot[0]))
										v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)+104))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v50
										v56 = v49
										v57 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v57)
										v59 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v59)
										*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = base.I64_extend_i32_u(v56 + int32(24))
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
										v70 = F_heap_form_tuple(m, v65, v8+int32(16), v8+int32(14))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int64(0)
										} else {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
											v73 = F_HeapTupleHeaderGetDatum(m, v72)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int64(0)
											} else {
												F_ReplicationSlotRelease(m)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int64(0)
												} else {
													m.G0 = v8 + int32(32)
													return v73
												}
											}
										}
									}
								}
							}
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, _c_F_pg_create_physical_replication_slot[0]))
							v56 = v54
							v57 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v57)
							v59 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v59)
							*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = base.I64_extend_i32_u(v56 + int32(24))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
							v70 = F_heap_form_tuple(m, v65, v8+int32(16), v8+int32(14))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int64(0)
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
								v73 = F_HeapTupleHeaderGetDatum(m, v72)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int64(0)
								} else {
									F_ReplicationSlotRelease(m)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int64(0)
									} else {
										m.G0 = v8 + int32(32)
										return v73
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_create_physical_replication_slot_0), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_create_physical_replication_slot_1), int32(89), int32(_a_F_pg_create_physical_replication_slot_2))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int64(0)
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
func F_pg_cryptohash_final(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	if l0 == int32(0) {
		return int32(-1)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = m.Env.Pgmem_hash_final(m, v8, l1, l2)
		mBase = m.M
		if int32(0) <= v9 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
			return int32(-1)
		}
	}
}
func F_pg_cryptohash_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	if l0 != 0 {
		v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		m.Env.Pgmem_hash_free(m, v2)
		mBase = m.M
		F___memset(m, l0, int32(0), int32(12))
		mBase = m.M
		F_pfree(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_pg_cursor(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = v7 + int32(76)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_cursor[0]))
	F_hash_seq_init(m, v16, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = F_hash_seq_search(m, v16)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = v21
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v7 + int32(96)
	return int64(0)
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v28 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v28)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v28
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+136)))
	if v32 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v78 = F_hash_seq_search(m, v7+int32(76))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L16
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	if v35 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v39 = F_cstring_to_text(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = base.I64_extend_i32_u(v39)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	v44 = F_cstring_to_text(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = base.I64_extend_i32_u(v44)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	v49 = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = base.I64_extend_i32_u(v48 & v49)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = base.I64_extend_i32_u(int32(base.Ui32(v48)>>(uint(v49)%32)) & v49)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = base.I64_extend_i32_u(int32(base.Ui32(v48)>>(uint(int32(5))%32)) & v49)
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v27)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+56)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_tuplestore_putvalues(m, v67, v68, v7+int32(16), v7+int32(8))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	if v78 != 0 {
		v23 = v78
		goto L8
	} else {
		goto L17
	}
L17:
	;
	goto L9
}
func F_pg_database_locale(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_database_locale[0]))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_pg_database_locale_0), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_pg_database_locale_1), int32(1198), int32(_a_F_pg_database_locale_2))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		return v3
	}
}
func F_pg_database_size_oid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v4 = int64(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = base.I32_wrap_i64(v9)
	v17 = F_SearchSysCacheExists(m, int32(21), v9&int64(4294967295), v4, v4, v4)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		if v17 != 0 {
			v21 = F_calculate_database_size(m, v10)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				if v21 == int64(0) {
					v25 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
				} else {
				}
				m.G0 = v7 + int32(16)
				return v21
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
					F_errmsg(m, int32(_a_F_pg_database_size_oid_0), v7)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_database_size_oid_1), int32(180), int32(_a_F_pg_database_size_oid_2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int64(0)
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
func F_pg_euccn2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v18 = v4
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	switch v19 - int32(142) {
	case 0:
		goto L10
	case 1:
		goto L9
	default:
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(0)
	return v82
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v66
	v71 = v18 + int32(1)
	v73 = v14 + int32(4)
	v74 = v15 + v67
	if int32(0) < v74 {
		v13 = v68
		v14 = v73
		v15 = v74
		v18 = v71
		goto L4
	} else {
		goto L18
	}
L8:
	;
	if v19 == int32(0) {
		v78 = v14
		v82 = v18
		goto L6
	} else {
		goto L13
	}
L9:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v78 = v14
		v82 = v18
		goto L6
	} else {
		goto L12
	}
L10:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v78 = v14
		v82 = v18
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v28 = v24<<(uint(int32(8))%32) | int32(_a_F_pg_euccn2wchar_with_len_0)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v66 = v28 | v30
	v67 = int32(-3)
	v68 = v13 + int32(3)
	goto L7
L12:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v41 = v37<<(uint(int32(8))%32) | int32(_a_F_pg_euccn2wchar_with_len_1)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v41
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v66 = v41 | v43
	v67 = int32(-3)
	v68 = v13 + int32(3)
	goto L7
L13:
	;
	if base.I32_extend8_s(v19) < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v15 == int32(1) {
		v78 = v14
		v82 = v18
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v66 = v19
	v67 = int32(-1)
	v68 = v13 + int32(1)
	goto L7
L17:
	;
	v56 = v19 << (uint(int32(8)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v56
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v66 = v56 | v58
	v67 = int32(-2)
	v68 = v13 + int32(2)
	goto L7
L18:
	;
	v78 = v73
	v82 = v71
	goto L6
}
func F_pg_euctw_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v5 = int32(-1)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v6 - int32(142) {
	case 0:
		if l1 < int32(4) {
			v49 = v5
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if base.Ui32((v11+int32(88))&int32(255)) < base.Ui32(int32(249)) {
				v49 = v5
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
				if base.Ui32(int32(93)) < base.Ui32((v18+int32(95))&int32(255)) {
					v49 = v5
				} else {
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
					if base.Ui32(int32(94)) <= base.Ui32((v25+int32(95))&int32(255)) {
						v49 = v5
					} else {
						v47 = int32(4)
						v49 = v47
					}
				}
			}
		}
	case 1:
		v49 = v5
	default:
		if int32(0) <= base.I32_extend8_s(v6) {
			v47 = int32(1)
			v49 = v47
		} else {
			v37 = int32(2)
			if l1 < v37 {
				v49 = v5
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if base.Ui32(int32(93)) < base.Ui32((v40+int32(95))&int32(255)) {
					v49 = v5
				} else {
					v47 = v37
					v49 = v47
				}
			}
		}
	}
	return v49
}
func F_pg_freespace(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_relation_open(m, v10, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+119)))
		switch v17 - int32(83) {
		case 0, 22, 26, 31, 33:
			if base.Ui64(int64(4294967295)) <= base.Ui64(v9) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_freespace_0), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_freespace_1), int32(47), int32(_a_F_pg_freespace_2))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v62 = F_GetRecordedFreeSpace(m, v12, base.I32_wrap_i64(v9))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int64(0)
				} else {
					F_relation_close(m, v12, int32(1))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int64(0)
					} else {
						m.G0 = v7 + int32(16)
						return base.I64_extend16_s(base.I64_extend_i32_u(v62))
					}
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27 + int32(4)
					F_errmsg(m, int32(_a_F_pg_freespace_3), v7)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
						v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v34)+119)))
						F_errdetail_relkind_not_supported(m, v35)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_freespace_1), int32(42), int32(_a_F_pg_freespace_2))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
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
func F_pg_gb18030_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	if l1 <= int32(0) {
		v75 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v75 - l0
L2:
	;
	v9 = l1
	v11 = l0
	goto L3
L3:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v14 = base.I32_extend8_s(v13)
	if int32(0) <= v14 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v75 = v68
	goto L1
L5:
	;
	v68 = v66 + v11
	v69 = v9 - v66
	if int32(0) < v69 {
		v9 = v69
		v11 = v68
		goto L3
	} else {
		goto L22
	}
L6:
	;
	if v14 == int32(0) {
		v75 = v11
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v9) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v66 = int32(1)
	goto L5
L10:
	;
	if base.B2i32(v13 == int32(128))|base.B2i32(v13 == int32(255)) != 0 {
		v75 = v11
		goto L1
	} else {
		goto L19
	}
L11:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if base.Ui32(int32(9)) < base.Ui32((v22-int32(48))&int32(255)) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v9 == int32(1) {
		v75 = v11
		goto L1
	} else {
		goto L18
	}
L14:
	;
	if base.B2i32(v13 == int32(128))|base.B2i32(v13 == int32(255)) != 0 {
		v75 = v11
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+2)))
	if base.Ui32((v34+int32(1))&int32(255)) < base.Ui32(int32(130)) {
		v75 = v11
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+3)))
	if base.Ui32(int32(10)) <= base.Ui32((v41-int32(48))&int32(255)) {
		v75 = v11
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v66 = int32(4)
	goto L5
L18:
	;
	goto L10
L19:
	;
	v56 = int32(2)
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v57 < int32(-1) {
		v66 = v56
		goto L5
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32((v57-int32(127))&int32(255)) < base.Ui32(int32(193)) {
		v75 = v11
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v66 = v56
	goto L5
L22:
	;
	goto L4
}
func F_pg_generic_charinc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v8 = l0 + l1 - int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pg_generic_charinc[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11*int32(28))+uint32(_c_F_pg_generic_charinc[1])))
	goto L1
L1:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v22 != int32(255) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return base.B2i32(v22 != int32(255))
L3:
	;
	v26 = v22 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v26)
	v28 = m.T0[v16].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	return int32(0)
L7:
	;
	if v28 != l1 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_pg_get_catalog_foreign_keys(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int64
	_ = v112
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v13 == int32(0) {
		v16 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = int32(_a_F_pg_get_catalog_foreign_keys_0)
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_catalog_foreign_keys[0]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_pg_get_catalog_foreign_keys[0])) = v23
			v28 = F_get_call_result_type(m, l0, int32(0), v8+int32(-48))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				if v28 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int64(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pg_get_catalog_foreign_keys_1), int32(0))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_get_catalog_foreign_keys_2), int32(483), int32(_a_F_pg_get_catalog_foreign_keys_3))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
					v33 = F_BlessTupleDesc(m, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v33
						v38 = F_palloc(m, int32(28))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							F_fmgr_info(m, int32(750), v38)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v38
								*(*int32)(unsafe.Add(mBase, _c_F_pg_get_catalog_foreign_keys[0])) = v21
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
								if base.Ui64(v50) <= base.Ui64(int64(225)) {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
									v54 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v54)
									*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v54
									v60 = base.I32_wrap_i64(v50) * int32(20)
									v61 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_get_catalog_foreign_keys[1]))))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v61
									v63 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_get_catalog_foreign_keys[2]))))
									v65 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_get_catalog_foreign_keys[3]))))
									v68 = F_FunctionCall3Coll(m, v53, v54, v65, int64(25), int64(-1))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v63
										*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v68
										v72 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_get_catalog_foreign_keys[4]))))
										v74 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_get_catalog_foreign_keys[5]))))
										v77 = F_FunctionCall3Coll(m, v53, int32(0), v74, int64(25), int64(-1))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v72
											v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_get_catalog_foreign_keys[6]))))
											*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v80
											*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v77
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v49)+28))
											v88 = F_heap_form_tuple(m, v83, v8+int32(-48), v8+int32(-56))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int64(0)
											} else {
												v90 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
												*(*int64)(unsafe.Add(mBase, uint32(v49))) = v90 + int64(1)
												v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(1)
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
												v98 = F_HeapTupleHeaderGetDatum(m, v97)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return int64(0)
												} else {
													v112 = v98
													m.G0 = v10 - int32(-64)
													return v112
												}
											}
										}
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int64(0)
									} else {
										v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v102)+20)) = int32(2)
										v105 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
										v112 = int64(0)
										m.G0 = v10 - int32(-64)
										return v112
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
		v50 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
		if base.Ui64(v50) <= base.Ui64(int64(225)) {
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
			v54 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v54)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v54
			v60 = base.I32_wrap_i64(v50) * int32(20)
			v61 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_get_catalog_foreign_keys[1]))))
			*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v61
			v63 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_get_catalog_foreign_keys[2]))))
			v65 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_get_catalog_foreign_keys[3]))))
			v68 = F_FunctionCall3Coll(m, v53, v54, v65, int64(25), int64(-1))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int64(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v63
				*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v68
				v72 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_get_catalog_foreign_keys[4]))))
				v74 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_get_catalog_foreign_keys[5]))))
				v77 = F_FunctionCall3Coll(m, v53, int32(0), v74, int64(25), int64(-1))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v72
					v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_get_catalog_foreign_keys[6]))))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v80
					*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v77
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v49)+28))
					v88 = F_heap_form_tuple(m, v83, v8+int32(-48), v8+int32(-56))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int64(0)
					} else {
						v90 = *(*int64)(unsafe.Add(mBase, uint32(v49)))
						*(*int64)(unsafe.Add(mBase, uint32(v49))) = v90 + int64(1)
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(1)
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
						v98 = F_HeapTupleHeaderGetDatum(m, v97)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int64(0)
						} else {
							v112 = v98
							m.G0 = v10 - int32(-64)
							return v112
						}
					}
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return int64(0)
			} else {
				v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v102)+20)) = int32(2)
				v105 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
				v112 = int64(0)
				m.G0 = v10 - int32(-64)
				return v112
			}
		}
	}
}
func F_pg_get_indexdef_ext(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if v13 == int64(0) {
		v16 = int32(2)
	} else {
		v16 = int32(7)
	}
	v18 = F_pg_get_indexdef_worker(m, v3, v4, v2, base.B2i32(v4 != v2), v2, v2, v2, v16, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		if v18 == int32(0) {
			v24 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
			return int64(0)
		} else {
			v28 = F_cstring_to_text(m, v18)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v18)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v28)
				}
			}
		}
	}
}
func F_pg_get_multixact_stats(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v15 = F_get_call_result_type(m, l0, int32(0), v10+int32(76))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		if v15 == int32(1) {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_multixact_stats[0]))
			v24 = F_has_privs_of_role(m, v22, int32(3375))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				if v24 == int32(0) {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
					if v29 == int32(0) {
						v83 = v28
					} else {
						base.MemoryFill(m, v10+int32(28), int32(1), v29)
						v83 = v28
					}
					v92 = F_heap_form_tuple(m, v83, v10+int32(32), v10+int32(28))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int64(0)
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
						v95 = F_HeapTupleHeaderGetDatum(m, v94)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int64(0)
						} else {
							m.G0 = v10 + int32(80)
							return v95
						}
					}
				} else {
					v39 = v10 + int32(24)
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_multixact_stats[1]))
					v47 = F_LWLockAcquire(m, v43+int32(1664), int32(1))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int64(0)
					} else {
						v49 = int32(_a_F_pg_get_multixact_stats_0)
						v50 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_multixact_stats[2]))
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v10))) = v51
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v39))) = v53
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_multixact_stats[2]))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
						v58 = *(*int64)(unsafe.Add(mBase, uint32(v56)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v10+int32(8)))) = v58
						v61 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_multixact_stats[1]))
						F_LWLockRelease(m, v61+int32(1664))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int64(0)
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
							*(*int32)(unsafe.Add(mBase, uint32(v10+int32(20)))) = v57 - v66
							*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(0)
							v71 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v71
							v73 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)))
							*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v73
							v75 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
							v77 = v75 - v76
							*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v77
							*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v77 * int64(5)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
							v83 = v82
							v92 = F_heap_form_tuple(m, v83, v10+int32(32), v10+int32(28))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int64(0)
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
								v95 = F_HeapTupleHeaderGetDatum(m, v94)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int64(0)
								} else {
									m.G0 = v10 + int32(80)
									return v95
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_get_multixact_stats_1), int32(0))
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_get_multixact_stats_2), int32(109), int32(_a_F_pg_get_multixact_stats_3))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int64(0)
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
func F_pg_get_publication_sequences(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v9 == v2 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = F_text_to_cstring(m, v13)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					v21 = int32(_a_F_pg_get_publication_sequences_0)
					v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_publication_sequences[0]))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
					*(*int32)(unsafe.Add(mBase, _c_F_pg_get_publication_sequences[0])) = v24
					v27 = F_get_publication_oid(m, v17, int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						if v27 != 0 {
							v29 = F_GetPublication(m, v27)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int64(0)
							} else {
								v31 = v29
								v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+9)))
								if v32 == int32(1) {
									v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
									v38 = F_GetAllPublicationRelations(m, v35, int32(83), int32(0))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int64(0)
									} else {
										v41 = v38
										*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v41
										*(*int32)(unsafe.Add(mBase, _c_F_pg_get_publication_sequences[0])) = v22
										v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
										v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
										if v51 == int32(0) {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int64(0)
											} else {
												v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = int32(2)
												v76 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
												return int64(0)
											}
										} else {
											v54 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
											v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v51)+4)))
											if base.Ui64(v55) <= base.Ui64(v54) {
												F_end_MultiFuncCall(m, l0)
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int64(0)
												} else {
													v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = int32(2)
													v76 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
													return int64(0)
												}
											} else {
												v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
												v62 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57+base.I32_wrap_i64(v54)<<(uint(int32(2))%32)))))
												*(*int64)(unsafe.Add(mBase, uint32(v50))) = v54 + int64(1)
												v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = int32(1)
												return v62
											}
										}
									}
								} else {
									v41 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v41
									*(*int32)(unsafe.Add(mBase, _c_F_pg_get_publication_sequences[0])) = v22
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
									if v51 == int32(0) {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int64(0)
										} else {
											v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = int32(2)
											v76 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
											return int64(0)
										}
									} else {
										v54 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
										v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v51)+4)))
										if base.Ui64(v55) <= base.Ui64(v54) {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int64(0)
											} else {
												v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = int32(2)
												v76 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
												return int64(0)
											}
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
											v62 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57+base.I32_wrap_i64(v54)<<(uint(int32(2))%32)))))
											*(*int64)(unsafe.Add(mBase, uint32(v50))) = v54 + int64(1)
											v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = int32(1)
											return v62
										}
									}
								}
							}
						} else {
							v31 = v2
							v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+9)))
							if v32 == int32(1) {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
								v38 = F_GetAllPublicationRelations(m, v35, int32(83), int32(0))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int64(0)
								} else {
									v41 = v38
									*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v41
									*(*int32)(unsafe.Add(mBase, _c_F_pg_get_publication_sequences[0])) = v22
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
									if v51 == int32(0) {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int64(0)
										} else {
											v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = int32(2)
											v76 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
											return int64(0)
										}
									} else {
										v54 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
										v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v51)+4)))
										if base.Ui64(v55) <= base.Ui64(v54) {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int64(0)
											} else {
												v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = int32(2)
												v76 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
												return int64(0)
											}
										} else {
											v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
											v62 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57+base.I32_wrap_i64(v54)<<(uint(int32(2))%32)))))
											*(*int64)(unsafe.Add(mBase, uint32(v50))) = v54 + int64(1)
											v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = int32(1)
											return v62
										}
									}
								}
							} else {
								v41 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v41
								*(*int32)(unsafe.Add(mBase, _c_F_pg_get_publication_sequences[0])) = v22
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
								if v51 == int32(0) {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int64(0)
									} else {
										v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = int32(2)
										v76 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
										return int64(0)
									}
								} else {
									v54 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
									v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v51)+4)))
									if base.Ui64(v55) <= base.Ui64(v54) {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int64(0)
										} else {
											v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = int32(2)
											v76 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
											return int64(0)
										}
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
										v62 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57+base.I32_wrap_i64(v54)<<(uint(int32(2))%32)))))
										*(*int64)(unsafe.Add(mBase, uint32(v50))) = v54 + int64(1)
										v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = int32(1)
										return v62
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
		if v51 == int32(0) {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int64(0)
			} else {
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = int32(2)
				v76 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
				return int64(0)
			}
		} else {
			v54 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
			v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v51)+4)))
			if base.Ui64(v55) <= base.Ui64(v54) {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int64(0)
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = int32(2)
					v76 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v76)
					return int64(0)
				}
			} else {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
				v62 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57+base.I32_wrap_i64(v54)<<(uint(int32(2))%32)))))
				*(*int64)(unsafe.Add(mBase, uint32(v50))) = v54 + int64(1)
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = int32(1)
				return v62
			}
		}
	}
}
func F_pg_get_ruledef_ext(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v6 == int64(0) {
		v9 = int32(2)
	} else {
		v9 = int32(7)
	}
	v10 = F_pg_get_ruledef_worker(m, v3, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		if v10 == int32(0) {
			v16 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
			return int64(0)
		} else {
			v20 = F_cstring_to_text(m, v10)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v10)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v20)
				}
			}
		}
	}
}
func F_pg_get_statisticsobj_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
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
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v126 int64
	_ = v126
	var v127 int32
	_ = v127
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
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v569 int32
	_ = v569
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v672 int32
	_ = v672
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	v4 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(192)
	m.G0 = v24
	v28 = F_SearchSysCache1(m, int32(64), base.I64_extend_i32_u(l0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L4
	} else {
		goto L155
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L4
	} else {
		goto L152
	}
L3:
	;
	m.G0 = v24 + int32(192)
	return v672
L4:
	;
	return int32(0)
L5:
	;
	if v28 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if l2 != 0 {
		v672 = int32(0)
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v50 = F_heap_attisnull(m, v28, int32(9), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L13
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_statisticsobj_worker_0), v24)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_pg_get_statisticsobj_worker_1), int32(1682), int32(_a_F_pg_get_statisticsobj_worker_2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
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
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v54 = v52 + v53
	if v50 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	F_initStringInfo(m, v24+int32(120))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L25
	}
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v81 = v65
	v82 = v70
	v83 = v71
	v84 = v77
	v85 = int32(0)
	goto L14
L16:
	;
	v81 = v4
	v82 = v74
	v83 = v75
	v84 = v4
	v85 = int32(1)
	goto L14
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	v74 = v54 + int32(96)
	v75 = v57
	goto L16
L18:
	;
	goto L19
L19:
	;
	v60 = F_SysCacheGetAttrNotNull(m, int32(64), v28, int32(9))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v63 = F_text_to_cstring(m, base.I32_wrap_i64(v60))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v65 = F_stringToNode(m, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_pfree(m, v63)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v70 = v54 + int32(96)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	if v65 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v74 = v70
	v75 = v71
	goto L16
L25:
	;
	if l1 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v54)+72))
	v95 = F_get_namespace_name_or_temp(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v385 = int32(1)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v386 <= int32(0) {
		goto L104
	} else {
		goto L105
	}
L29:
	;
	v98 = v24 + int32(136)
	F_initStringInfo(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	if v95 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v101 = F_quote_identifier(m, v95)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v111 = F_quote_identifier(m, v54+int32(8))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v101
	F_appendStringInfo(m, v98, int32(_a_F_pg_get_statisticsobj_worker_3), v24+int32(112))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	F_appendStringInfoString(m, v24+int32(136), v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v24)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v115
	F_appendStringInfo(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_4), v24+int32(96))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v126 = F_SysCacheGetAttrNotNull(m, int32(64), v28, int32(8))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v129 = F_pg_detoast_datum(m, base.I32_wrap_i64(v126))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v131 != int32(1) {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	if v134 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	if v135 != int32(18) {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	if v138 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v292&(v288&v291)|base.B2i32(v83+v84 < int32(2)) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L45:
	;
	v288 = int32(0)
	v291 = v4
	v292 = v4
	goto L44
L46:
	;
	goto L47
L47:
	;
	v143 = v129 + int32(24)
	v145 = v138 & int32(3)
	v146 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v138) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v152 = v146
	v156 = v146
	v159 = v4
	v160 = v4
	v171 = v4
	goto L51
L49:
	;
	v225 = v146
	v229 = v146
	v232 = v4
	v233 = v4
	goto L50
L50:
	;
	v246 = v225
	v248 = v229
	v253 = v232
	v254 = v233
	v266 = v4
	goto L71
L51:
	;
	v173 = int32(1)
	v175 = v152 + v143
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	switch v176 - int32(100) {
	case 0:
		v183 = v173
		v184 = v159
		v185 = v160
		goto L53
	default:
		v181 = v159
		v182 = v160
		goto L54
	case 2:
		goto L56
	case 9:
		goto L55
	}
L52:
	;
	if v145 == int32(0) {
		v288 = v215
		v291 = v216
		v292 = v217
		goto L44
	} else {
		goto L70
	}
L53:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	switch v186 - int32(100) {
	case 0:
		v193 = v173
		v194 = v184
		v195 = v185
		goto L57
	default:
		v191 = v184
		v192 = v185
		goto L58
	case 2:
		goto L59
	case 9:
		goto L60
	}
L54:
	;
	v183 = v156
	v184 = v181
	v185 = v182
	goto L53
L55:
	;
	v181 = v159
	v182 = int32(1)
	goto L54
L56:
	;
	v181 = int32(1)
	v182 = v160
	goto L54
L57:
	;
	v196 = int32(1)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
	switch v198 - int32(100) {
	case 0:
		v205 = v196
		v206 = v194
		v207 = v195
		goto L61
	default:
		v203 = v194
		v204 = v195
		goto L62
	case 2:
		goto L63
	case 9:
		goto L64
	}
L58:
	;
	v193 = v183
	v194 = v191
	v195 = v192
	goto L57
L59:
	;
	v191 = int32(1)
	v192 = v185
	goto L58
L60:
	;
	v191 = v184
	v192 = int32(1)
	goto L58
L61:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+3)))
	switch v208 - int32(100) {
	case 0:
		v215 = v196
		v216 = v206
		v217 = v207
		goto L65
	default:
		v213 = v206
		v214 = v207
		goto L66
	case 2:
		goto L67
	case 9:
		goto L68
	}
L62:
	;
	v205 = v193
	v206 = v203
	v207 = v204
	goto L61
L63:
	;
	v203 = int32(1)
	v204 = v195
	goto L62
L64:
	;
	v203 = v194
	v204 = int32(1)
	goto L62
L65:
	;
	v218 = int32(4)
	v219 = v152 + v218
	v221 = v171 + v218
	if v221 != v138&int32(2147483644) {
		v152 = v219
		v156 = v215
		v159 = v216
		v160 = v217
		v171 = v221
		goto L51
	} else {
		goto L69
	}
L66:
	;
	v215 = v205
	v216 = v213
	v217 = v214
	goto L65
L67:
	;
	v213 = int32(1)
	v214 = v207
	goto L66
L68:
	;
	v213 = v206
	v214 = int32(1)
	goto L66
L69:
	;
	goto L52
L70:
	;
	v225 = v219
	v229 = v215
	v232 = v216
	v233 = v217
	goto L50
L71:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246+v143))))
	switch v269 - int32(100) {
	case 0:
		v276 = int32(1)
		v277 = v253
		v278 = v254
		goto L73
	default:
		v274 = v253
		v275 = v254
		goto L74
	case 2:
		goto L75
	case 9:
		goto L76
	}
L72:
	;
	v288 = v276
	v291 = v277
	v292 = v278
	goto L44
L73:
	;
	v279 = int32(1)
	v282 = v266 + v279
	if v282 != v145 {
		v246 = v246 + v279
		v248 = v276
		v253 = v277
		v254 = v278
		v266 = v282
		goto L71
	} else {
		goto L77
	}
L74:
	;
	v276 = v248
	v277 = v274
	v278 = v275
	goto L73
L75:
	;
	v274 = int32(1)
	v275 = v254
	goto L74
L76:
	;
	v274 = v253
	v275 = int32(1)
	goto L74
L77:
	;
	goto L72
L78:
	;
	v314 = v24 + int32(120)
	F_appendStringInfoString(m, v314, int32(_a_F_pg_get_statisticsobj_worker_5))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_appendStringInfoString(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_6))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L4
	} else {
		goto L102
	}
L81:
	;
	if v288&int32(1) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_appendStringInfoString(m, v314, int32(_a_F_pg_get_statisticsobj_worker_7))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L4
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v291 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L84
L86:
	;
	if v292 != 0 {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	v339 = v288
	goto L86
L88:
	;
	goto L89
L89:
	;
	v325 = int32(1)
	if v288&v325 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v330 = int32(_a_F_pg_get_statisticsobj_worker_8)
	goto L92
L91:
	;
	v330 = int32(_a_F_pg_get_statisticsobj_worker_9)
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v330
	F_appendStringInfo(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_10), v24+int32(80))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v339 = v325
	goto L86
L94:
	;
	if v339&int32(1) != 0 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	F_appendStringInfoChar(m, v24+int32(120), int32(41))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L4
	} else {
		goto L101
	}
L97:
	;
	v344 = int32(_a_F_pg_get_statisticsobj_worker_8)
	goto L99
L98:
	;
	v344 = int32(_a_F_pg_get_statisticsobj_worker_9)
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v344
	F_appendStringInfo(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_11), v24-int32(-64))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	goto L96
L101:
	;
	goto L80
L102:
	;
	goto L28
L103:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v470 = F_get_rel_name(m, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L4
	} else {
		goto L118
	}
L104:
	;
	v448 = int32(0)
	goto L103
L105:
	;
	goto L106
L106:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v393 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+104)))
	v395 = F_get_attname(m, v392, v393, int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v397 = F_quote_identifier(m, v395)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	F_appendStringInfoString(m, v24+int32(120), v397)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	if v401 < int32(2) {
		v448 = v385
		goto L103
	} else {
		goto L110
	}
L110:
	;
	v406 = v385
	goto L111
L111:
	;
	v430 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54+int32(104)+v406<<(uint(int32(1))%32)))))
	v432 = v24 + int32(120)
	F_appendStringInfoString(m, v432, int32(_a_F_pg_get_statisticsobj_worker_8))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L113
	}
L112:
	;
	v448 = v445
	goto L103
L113:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v438 = F_get_attname(m, v436, v430, int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	v440 = F_quote_identifier(m, v438)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	F_appendStringInfoString(m, v432, v440)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	v445 = v406 + int32(1)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v54)+96))
	if v445 < v446 {
		v406 = v445
		goto L111
	} else {
		goto L117
	}
L117:
	;
	goto L112
L118:
	;
	if v470 == int32(0) {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v476 = F_palloc0(m, int32(80))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	v479 = F_palloc0(m, int32(136))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v479)+24)) = int32(1)
	v483 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v479)+21)) = uint8(v483)
	*(*int32)(unsafe.Add(mBase, uint32(v479)+16)) = v474
	v486 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v479)+12)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v479))) = int32(101)
	v492 = F_makeAlias(m, v470, v486)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v479)+8)) = v492
	*(*int32)(unsafe.Add(mBase, uint32(v479)+4)) = v492
	v496 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v479)+124)) = uint16(v496)
	v498 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v479)+20)) = uint8(v498)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v24)+136)) = v479
	v505 = F_list_make1_impl(m, int32(1), v24+int32(60))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v507 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v476)+20)) = v507
	*(*int64)(unsafe.Add(mBase, uint32(v476)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v476))) = v505
	F_set_rtable_names(m, v476, v507, v507)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	F_set_simple_column_names(m, v476)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v476
	*(*int32)(unsafe.Add(mBase, uint32(v24)+176)) = v476
	v523 = F_list_make1_impl(m, int32(1), v24+int32(56))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	if v85 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if l1 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L128:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v525 <= int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v528 = v448
	v530 = v486
	goto L130
L130:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v549+v530<<(uint(int32(2))%32))))
	v555 = v24 + int32(176)
	F_initStringInfo(m, v555)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L4
	} else {
		goto L132
	}
L131:
	;
	goto L127
L132:
	;
	v558 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+168)) = uint8(v558)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+152)) = v558
	*(*int64)(unsafe.Add(mBase, uint32(v24)+144)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+140)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v24)+172)) = v558
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+171)) = uint8(v558)
	v569 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+169)) = uint16(v569)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+164)) = v558
	*(*int64)(unsafe.Add(mBase, uint32(v24)+156)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+136)) = v555
	F_get_rule_expr(m, v553, v24+int32(136), v558)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v24)+176))
	if int32(0) < v528 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	F_appendStringInfoString(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_8))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L4
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	if v553 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	goto L136
L138:
	;
	v607 = int32(1)
	v610 = v530 + v607
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v610 < v611 {
		v528 = v528 + v607
		v530 = v610
		goto L130
	} else {
		goto L145
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v581
	F_appendStringInfo(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_12), v24+int32(48))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L4
	} else {
		goto L144
	}
L140:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	switch v591 - int32(15) {
	case 0:
		goto L142
	default:
		goto L139
	case 4, 23, 24, 25, 26, 33:
		goto L141
	}
L141:
	;
	F_appendStringInfoString(m, v24+int32(120), v581)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L143
	}
L142:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v553)+16))
	switch v594 {
	case 0, 3:
		goto L141
	default:
		goto L139
	}
L143:
	;
	goto L138
L144:
	;
	goto L138
L145:
	;
	goto L131
L146:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v638 = F_generate_relation_name(m, v636, int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L4
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	F_ReleaseCatCache(m, v28)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L4
	} else {
		goto L151
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v638
	F_appendStringInfo(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_13), v24+int32(32))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v24)+120))
	v672 = v650
	goto L3
L152:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_statisticsobj_worker_14), int32(0))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_pg_get_statisticsobj_worker_1), int32(1731), int32(_a_F_pg_get_statisticsobj_worker_2))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v469
	F_errmsg_internal(m, int32(_a_F_pg_get_statisticsobj_worker_15), v24+int32(16))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_pg_get_statisticsobj_worker_1), int32(_a_F_pg_get_statisticsobj_worker_16), int32(_a_F_pg_get_statisticsobj_worker_17))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_statisticsobjdef(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14347(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_pg_get_statisticsobjdef_columns(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14347(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_pg_get_userbyid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v12 = F_palloc(m, int32(64))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v12))) = v16
		v33 = F_SearchSysCache1(m, int32(11), v10)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int64(0)
		} else {
			if v33 != 0 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
				v37 = v35 + v36
				v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+60))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v38
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v37)+52))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v40
				v42 = *(*int64)(unsafe.Add(mBase, uint32(v37)+44))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v42
				v44 = *(*int64)(unsafe.Add(mBase, uint32(v37)+36))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v44
				v46 = *(*int64)(unsafe.Add(mBase, uint32(v37)+28))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v46
				v48 = *(*int64)(unsafe.Add(mBase, uint32(v37)+20))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v48
				v50 = *(*int64)(unsafe.Add(mBase, uint32(v37)+12))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v50
				v52 = *(*int64)(unsafe.Add(mBase, uint32(v37)+4))
				*(*int64)(unsafe.Add(mBase, uint32(v12))) = v52
				F_ReleaseCatCache(m, v33)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int64(0)
				} else {
					m.G0 = v8 + int32(16)
					return base.I64_extend_i32_u(v12)
				}
			} else {
				*(*uint32)(unsafe.Add(mBase, uint32(v8))) = uint32(v10)
				v58 = F_pg_sprintf(m, v12, int32(_a_F_pg_get_userbyid_0), v8)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int64(0)
				} else {
					m.G0 = v8 + int32(16)
					return base.I64_extend_i32_u(v12)
				}
			}
		}
	}
}
func F_pg_getopt_start(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(_a_F_pg_getopt_start_0)
	v10 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v10
	return
}
func F_pg_hba_file_rules(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(784)
	m.G0 = v22
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+284)) = v32
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_pg_hba_file_rules[0]))
	v39 = F_open_auth_file(m, v35, int32(21), v32, v32)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_pg_hba_file_rules[0]))
	F_tokenize_auth_file(m, v42, v39, v22+int32(284), int32(12), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_pg_hba_file_rules[1]))
	v55 = F_AllocSetContextCreateInternal(m, v50, int32(_a_F_pg_hba_file_rules_0), int32(0), int32(1024), int32(_a_F_pg_hba_file_rules_1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v57 = int32(_a_F_pg_hba_file_rules_2)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_pg_hba_file_rules[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_hba_file_rules[1])) = v55
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v22)+284))
	if v61 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_free_auth_file(m, v39)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L1
	} else {
		goto L220
	}
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v72 = v22 + int32(547)
	v84 = v2
	v85 = v2
	goto L9
L9:
	;
	v92 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93+v84<<(uint(int32(2))%32))))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	if v98 == v92 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v102 = F_parse_hba_line(m, v97, int32(12))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v105 = v92
	v106 = v98
	goto L13
L13:
	;
	v107 = int64(*(*int32)(unsafe.Add(mBase, uint32(v97)+8)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v111 = int32(0)
	base.MemoryFill(m, v22+int32(560), v111, int32(88))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+551)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v22)+544)) = int64(0)
	v119 = v85 + int32(1)
	if v106 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	v105 = v102
	v106 = v104
	goto L13
L15:
	;
	v124 = F_cstring_to_text(m, v108)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	v120 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+544)) = uint8(v120)
	goto L15
L17:
	;
	goto L18
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+560)) = base.I64_extend_i32_s(v119)
	goto L15
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+576)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v22)+568)) = base.I64_extend_i32_u(v124)
	if v105 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v106 != 0 {
		goto L210
	} else {
		goto L211
	}
L21:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	if base.Ui32(v129) <= base.Ui32(int32(5)) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v827 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+3)) = v827
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v827
	goto L20
L24:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	if v141 != 0 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129<<(uint(int32(2))%32))+uint32(_c_F_pg_hba_file_rules[2])))
	v135 = F_cstring_to_text(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v139 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+547)) = uint8(v139)
	goto L24
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+584)) = base.I64_extend_i32_u(v135)
	goto L24
L29:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
	if v223 != 0 {
		goto L43
	} else {
		goto L44
	}
L30:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v142 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L32
L32:
	;
	v202 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+548)) = uint8(v202)
	goto L29
L33:
	;
	v198 = F_strlist_to_textarray(m, v182)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L41
	}
L34:
	;
	v182 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v146 = int32(0)
	v149 = v146
	v151 = v146
	goto L37
L37:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v149<<(uint(int32(2))%32))))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v173 = F_lappend(m, v151, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v182 = v173
	goto L33
L39:
	;
	v176 = v149 + int32(1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v176 < v177 {
		v149 = v176
		v151 = v173
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+592)) = base.I64_extend_i32_u(v198)
	goto L29
L42:
	;
	v306 = int32(0)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v105)+288))
	switch v307 {
	case 0:
		goto L60
	case 1:
		goto L58
	case 2:
		goto L57
	case 3:
		v379 = v306
		v380 = int32(_a_F_pg_hba_file_rules_3)
		goto L56
	default:
		v372 = v306
		goto L59
	}
L43:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v224 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	goto L45
L45:
	;
	v284 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+549)) = uint8(v284)
	goto L42
L46:
	;
	v280 = F_strlist_to_textarray(m, v264)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L54
	}
L47:
	;
	v264 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v228 = int32(0)
	v231 = v228
	v233 = v228
	goto L50
L50:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249+v231<<(uint(int32(2))%32))))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v255 = F_lappend(m, v233, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	v264 = v255
	goto L46
L52:
	;
	v258 = v231 + int32(1)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v258 < v259 {
		v231 = v258
		v233 = v255
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+600)) = base.I64_extend_i32_u(v280)
	goto L42
L55:
	;
	if v386 != 0 {
		goto L89
	} else {
		goto L90
	}
L56:
	;
	v382 = F_cstring_to_text(m, v380)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L87
	}
L57:
	;
	v379 = v306
	v380 = int32(_a_F_pg_hba_file_rules_4)
	goto L56
L58:
	;
	v379 = v306
	v380 = int32(_a_F_pg_hba_file_rules_5)
	goto L56
L59:
	;
	v375 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+550)) = uint8(v375)
	v386 = v372
	goto L55
L60:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v105)+292))
	if v308 != 0 {
		v379 = v306
		v380 = v308
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v309 = int32(0)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v105)+152))
	if v309 < v310 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v314 = v105 + int32(24)
	v316 = v22 + int32(288)
	v318 = int32(0)
	v321 = F_pg_getnameinfo_all(m, v314, v310, v316, int32(255), v318, v318, int32(1))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	v339 = v309
	goto L64
L64:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v105)+284))
	if int32(0) < v341 {
		goto L74
	} else {
		goto L75
	}
L65:
	;
	if v321 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v314))))
	if v325 != int32(10) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L68
L68:
	;
	v337 = F_pstrdup(m, v22+int32(288))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L73
	}
L69:
	;
	goto L68
L70:
	;
	goto L69
L71:
	;
	v329 = F_strchr(m, v316, int32(37))
	mBase = m.M
	if v329 == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v332 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v329))) = uint8(v332)
	goto L70
L73:
	;
	v339 = v337
	goto L64
L74:
	;
	v345 = v105 + int32(156)
	v347 = v22 + int32(288)
	v349 = int32(0)
	v352 = F_pg_getnameinfo_all(m, v345, v341, v347, int32(255), v349, v349, int32(1))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v370 = v306
	goto L76
L76:
	;
	if v339 != 0 {
		v379 = v370
		v380 = v339
		goto L56
	} else {
		goto L86
	}
L77:
	;
	if v352 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v345))))
	if v356 != int32(10) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L80
L80:
	;
	v368 = F_pstrdup(m, v22+int32(288))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L85
	}
L81:
	;
	goto L80
L82:
	;
	goto L81
L83:
	;
	v360 = F_strchr(m, v347, int32(37))
	mBase = m.M
	if v360 == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v363 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v360))) = uint8(v363)
	goto L82
L85:
	;
	v370 = v368
	goto L76
L86:
	;
	v372 = v370
	goto L59
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+608)) = base.I64_extend_i32_u(v382)
	v386 = v379
	goto L55
L88:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v105)+296))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v395<<(uint(int32(2))%32))+uint32(_c_F_pg_hba_file_rules[3])))
	goto L93
L89:
	;
	v389 = F_cstring_to_text(m, v386)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v393 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+551)) = uint8(v393)
	goto L88
L92:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+616)) = base.I64_extend_i32_u(v389)
	goto L88
L93:
	;
	v399 = F_cstring_to_text(m, v398)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+624)) = base.I64_extend_i32_u(v399)
	v403 = int32(0)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v105)+296))
	if base.Ui32(int32(1)) < base.Ui32(v404-int32(7)) {
		v437 = v403
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v105)+300))
	if v439 != 0 {
		goto L104
	} else {
		goto L105
	}
L96:
	;
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+368)))
	if v409 != int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v420 = v403
	v421 = v22 + int32(656)
	goto L99
L98:
	;
	v415 = F_cstring_to_text(m, int32(_a_F_pg_hba_file_rules_6))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L100
	}
L99:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v105)+364))
	if v422 == int32(0) {
		v437 = v420
		goto L95
	} else {
		goto L101
	}
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+656)) = base.I64_extend_i32_u(v415)
	v420 = int32(1)
	v421 = v22 + int32(656) | int32(8)
	goto L99
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+272)) = v422
	v429 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_7), v22+int32(272))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v431 = F_cstring_to_text(m, v429)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v421))) = base.I64_extend_i32_u(v431)
	v437 = v420 + int32(1)
	goto L95
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+256)) = v439
	v449 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_8), v22+int32(256))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	v457 = v437
	goto L106
L106:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v105)+356))
	if v458 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v451 = F_cstring_to_text(m, v449)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v437<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v451)
	v457 = v437 + int32(1)
	goto L106
L109:
	;
	if v458 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v481 = v457
	goto L111
L111:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v105)+304))
	if v482 != 0 {
		goto L117
	} else {
		goto L118
	}
L112:
	;
	v463 = int32(_a_F_pg_hba_file_rules_9)
	goto L114
L113:
	;
	v463 = int32(_a_F_pg_hba_file_rules_10)
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+240)) = v463
	v473 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_11), v22+int32(240))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v475 = F_cstring_to_text(m, v473)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v457<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v475)
	v481 = v457 + int32(1)
	goto L111
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+224)) = v482
	v492 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_12), v22+int32(224))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	v500 = v481
	goto L119
L119:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v105)+296))
	if v501 == int32(11) {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	v494 = F_cstring_to_text(m, v492)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v481<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v494)
	v500 = v481 + int32(1)
	goto L119
L122:
	;
	v825 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+553)) = uint8(v825)
	goto L20
L123:
	;
	v817 = F_construct_array_builtin(m, v22+int32(656), v812, int32(25))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L207
	}
L124:
	;
	if v808 == int32(0) {
		goto L122
	} else {
		goto L206
	}
L125:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v105)+316))
	if v504 != 0 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	v730 = v500
	v732 = v501
	goto L127
L127:
	;
	if v732 != int32(14) {
		v808 = v730
		goto L124
	} else {
		goto L187
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = v504
	v514 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_13), v22+int32(208))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	v522 = v500
	goto L130
L130:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v105)+320))
	if v523 != 0 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v516 = F_cstring_to_text(m, v514)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v500<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v516)
	v522 = v500 + int32(1)
	goto L130
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v523
	v533 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_14), v22+int32(192))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	v541 = v522
	goto L135
L135:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v105)+312))
	if v542 != 0 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v535 = F_cstring_to_text(m, v533)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v522<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v535)
	v541 = v522 + int32(1)
	goto L135
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v542
	v552 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_15), v22+int32(176))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	v560 = v541
	goto L140
L140:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+309)))
	if v561 == int32(1) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v554 = F_cstring_to_text(m, v552)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v541<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v554)
	v560 = v541 + int32(1)
	goto L140
L143:
	;
	v570 = F_cstring_to_text(m, int32(_a_F_pg_hba_file_rules_16))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L146
	}
L144:
	;
	v576 = v560
	goto L145
L145:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v105)+348))
	if v577 != 0 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v560<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v570)
	v576 = v560 + int32(1)
	goto L145
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v577
	v587 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_17), v22+int32(160))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	v595 = v576
	goto L149
L149:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v105)+352))
	if v596 != 0 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	v589 = F_cstring_to_text(m, v587)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v576<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v589)
	v595 = v576 + int32(1)
	goto L149
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v596
	v606 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_18), v22+int32(144))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L155
	}
L153:
	;
	v614 = v595
	goto L154
L154:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v105)+340))
	if v615 != 0 {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	v608 = F_cstring_to_text(m, v606)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v595<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v608)
	v614 = v595 + int32(1)
	goto L154
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v615
	v625 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_19), v22+int32(128))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	v633 = v614
	goto L159
L159:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v105)+324))
	if v634 != 0 {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	v627 = F_cstring_to_text(m, v625)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v614<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v627)
	v633 = v614 + int32(1)
	goto L159
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v634
	v644 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_20), v22+int32(112))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L165
	}
L163:
	;
	v652 = v633
	goto L164
L164:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v105)+328))
	if v653 != 0 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v646 = F_cstring_to_text(m, v644)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v633<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v646)
	v652 = v633 + int32(1)
	goto L164
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v653
	v663 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_21), v22+int32(96))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	v671 = v652
	goto L169
L169:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v105)+332))
	if v672 != 0 {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	v665 = F_cstring_to_text(m, v663)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v652<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v665)
	v671 = v652 + int32(1)
	goto L169
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v672
	v682 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_22), v22+int32(80))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	v690 = v671
	goto L174
L174:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v105)+336))
	if v691 != 0 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v684 = F_cstring_to_text(m, v682)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v671<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v684)
	v690 = v671 + int32(1)
	goto L174
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v691
	v701 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_23), v22-int32(-64))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L180
	}
L178:
	;
	v709 = v690
	goto L179
L179:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v105)+344))
	if v710 != 0 {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	v703 = F_cstring_to_text(m, v701)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v690<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v703)
	v709 = v690 + int32(1)
	goto L179
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v710
	v720 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_24), v22+int32(48))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L185
	}
L183:
	;
	v728 = v709
	goto L184
L184:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v105)+296))
	v730 = v728
	v732 = v729
	goto L127
L185:
	;
	v722 = F_cstring_to_text(m, v720)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v709<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v722)
	v728 = v709 + int32(1)
	goto L184
L187:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v105)+372))
	if v735 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v735
	v745 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_25), v22+int32(32))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L191
	}
L189:
	;
	v753 = v730
	goto L190
L190:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v105)+376))
	if v754 != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	v747 = F_cstring_to_text(m, v745)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v730<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v747)
	v753 = v730 + int32(1)
	goto L190
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v754
	v764 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_26), v22+int32(16))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L196
	}
L194:
	;
	v772 = v753
	goto L195
L195:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v105)+380))
	if v773 != 0 {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	v766 = F_cstring_to_text(m, v764)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v753<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v766)
	v772 = v753 + int32(1)
	goto L195
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v773
	v781 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_27), v22)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L201
	}
L199:
	;
	v789 = v772
	goto L200
L200:
	;
	v790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+384)))
	if v790 != int32(1) {
		v808 = v789
		goto L124
	} else {
		goto L203
	}
L201:
	;
	v783 = F_cstring_to_text(m, v781)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v772<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v783)
	v789 = v772 + int32(1)
	goto L200
L203:
	;
	v800 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_28), int32(0))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v802 = F_cstring_to_text(m, v800)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22+int32(656)+v789<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v802)
	v812 = v789 + int32(1)
	goto L123
L206:
	;
	v812 = v808
	goto L123
L207:
	;
	if v817 == int32(0) {
		goto L122
	} else {
		goto L208
	}
L208:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+632)) = base.I64_extend_i32_u(v817)
	goto L20
L209:
	;
	if v106 != 0 {
		goto L214
	} else {
		goto L215
	}
L210:
	;
	v850 = F_cstring_to_text(m, v106)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v854 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+554)) = uint8(v854)
	goto L209
L213:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+640)) = base.I64_extend_i32_u(v850)
	goto L209
L214:
	;
	v856 = v85
	goto L216
L215:
	;
	v856 = v119
	goto L216
L216:
	;
	v861 = F_heap_form_tuple(m, v30, v22+int32(560), v22+int32(544))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	F_tuplestore_puttuple(m, v31, v861)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v866 = v84 + int32(1)
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v866 < v867 {
		v84 = v866
		v85 = v856
		goto L9
	} else {
		goto L219
	}
L219:
	;
	goto L10
L220:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_hba_file_rules[1])) = v58
	F_MemoryContextDelete(m, v55)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v894 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v894)
	m.G0 = v22 + int32(784)
	return int64(0)
}
func F_pg_hmac_error(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	if l0 == int32(0) {
		return int32(_a_F_pg_hmac_error_0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7 != 0 {
			v19 = v7
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v11 == int32(2) {
				v14 = int32(_a_F_pg_hmac_error_1)
			} else {
				v14 = int32(_a_F_pg_hmac_error_2)
			}
			if v11 == int32(1) {
				v17 = int32(_a_F_pg_hmac_error_0)
			} else {
				v17 = v14
			}
			v19 = v17
		}
		return v19
	}
}
func F_pg_hmac_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
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
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	v4 = int32(0)
	v14 = int32(-1)
	if l0 == v4 {
		v229 = v14
		return v229
	} else {
		v18 = l0 + int32(152)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v21 = int32(0)
		v22 = base.B2i32(v20 == v21)
		if v22 == v21 {
			base.MemoryFill(m, v18, int32(92), v20)
		} else {
		}
		v28 = l0 + int32(24)
		if v22 == int32(0) {
			base.MemoryFill(m, v28, int32(54), v20)
		} else {
		}
		if base.Ui32(l2) <= base.Ui32(v20) {
			v92 = l2
			v93 = l1
			v94 = v4
			if v92 == int32(0) {
			} else {
				v97 = int32(0)
				if v92 != int32(1) {
					v106 = int32(0)
					v107 = v97
					for {
						v118 = v107 + v28
						v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
						v120 = v107 + v93
						v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
						v122 = v119 ^ v121
						*(*uint8)(unsafe.Add(mBase, uint32(v118))) = uint8(v122)
						v124 = v107 + v18
						v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
						v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
						v127 = v125 ^ v126
						*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v127)
						v130 = v107 | int32(1)
						v131 = v28 + v130
						v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
						v133 = v93 + v130
						v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
						v135 = v132 ^ v134
						*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v135)
						v137 = v18 + v130
						v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
						v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
						v140 = v138 ^ v139
						*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v140)
						v142 = int32(2)
						v143 = v107 + v142
						v145 = v106 + v142
						if v145 != v92&int32(-2) {
							v106 = v145
							v107 = v143
							continue
						} else {
							break
						}
						break
					}
					if v92&int32(1) == int32(0) {
					} else {
						v151 = v143
						v162 = v151 + v28
						v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
						v164 = v151 + v93
						v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
						v166 = v163 ^ v165
						*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v166)
						v168 = v151 + v18
						v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
						v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
						v171 = v169 ^ v170
						*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v171)
					}
				} else {
					v151 = v97
					v162 = v151 + v28
					v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
					v164 = v151 + v93
					v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
					v166 = v163 ^ v165
					*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v166)
					v168 = v151 + v18
					v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
					v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
					v171 = v169 ^ v170
					*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v171)
				}
			}
			v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v187 = F_pg_cryptohash_init(m, v186)
			mBase = m.M
			if int32(0) <= v187 {
				v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v192 = F_pg_cryptohash_update(m, v190, v28, v191)
				mBase = m.M
				if int32(0) <= v192 {
					v214 = int32(0)
					if v94 == v214 {
						v229 = v214
						return v229
					} else {
						v217 = v214
						F_pfree(m, v94)
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return int32(0)
						} else {
							v229 = v217
							return v229
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
					v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					if v197 == int32(0) {
						v212 = int32(_a_F_pg_hmac_init_0)
					} else {
						v204 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
						if v204 == int32(1) {
							v207 = int32(_a_F_pg_hmac_init_1)
						} else {
							v207 = int32(_a_F_pg_hmac_init_2)
						}
						if v204 == int32(2) {
							v210 = int32(_a_F_pg_hmac_init_0)
						} else {
							v210 = v207
						}
						v212 = v210
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v212
					if v94 != 0 {
						v217 = v14
						F_pfree(m, v94)
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return int32(0)
						} else {
							v229 = v217
							return v229
						}
					} else {
						v229 = v14
						return v229
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
				v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v197 == int32(0) {
					v212 = int32(_a_F_pg_hmac_init_0)
				} else {
					v204 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
					if v204 == int32(1) {
						v207 = int32(_a_F_pg_hmac_init_1)
					} else {
						v207 = int32(_a_F_pg_hmac_init_2)
					}
					if v204 == int32(2) {
						v210 = int32(_a_F_pg_hmac_init_0)
					} else {
						v210 = v207
					}
					v212 = v210
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v212
				if v94 != 0 {
					v217 = v14
					F_pfree(m, v94)
					mBase = m.M
					v219 = m.ExcPending
					if v219 != 0 {
						return int32(0)
					} else {
						v229 = v217
						return v229
					}
				} else {
					v229 = v14
					return v229
				}
			}
		} else {
			v34 = F_palloc(m, v19)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				if v34 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
					return int32(-1)
				} else {
					if v19 != 0 {
						base.MemoryFill(m, v34, int32(0), v19)
					} else {
					}
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v47 = F_pg_cryptohash_create(m, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						if v47 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
							F_pfree(m, v34)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								return int32(-1)
							}
						} else {
							v57 = F_pg_cryptohash_init(m, v47)
							mBase = m.M
							if v57 < int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
								if v47 == int32(0) {
									v82 = int32(_a_F_pg_hmac_init_0)
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
									if v74 == int32(1) {
										v77 = int32(_a_F_pg_hmac_init_1)
									} else {
										v77 = int32(_a_F_pg_hmac_init_2)
									}
									if v74 == int32(2) {
										v80 = int32(_a_F_pg_hmac_init_0)
									} else {
										v80 = v77
									}
									v82 = v80
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v82
								F_pg_cryptohash_free(m, v47)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v34)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										return int32(-1)
									}
								}
							} else {
								v60 = F_pg_cryptohash_update(m, v47, l1, l2)
								mBase = m.M
								if v60 < int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
									if v47 == int32(0) {
										v82 = int32(_a_F_pg_hmac_init_0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
										if v74 == int32(1) {
											v77 = int32(_a_F_pg_hmac_init_1)
										} else {
											v77 = int32(_a_F_pg_hmac_init_2)
										}
										if v74 == int32(2) {
											v80 = int32(_a_F_pg_hmac_init_0)
										} else {
											v80 = v77
										}
										v82 = v80
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v82
									F_pg_cryptohash_free(m, v47)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v34)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											return int32(-1)
										}
									}
								} else {
									v63 = F_pg_cryptohash_final(m, v47, v34, v19)
									mBase = m.M
									if int32(0) <= v63 {
										F_pg_cryptohash_free(m, v47)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											v92 = v19
											v93 = v34
											v94 = v34
											if v92 == int32(0) {
											} else {
												v97 = int32(0)
												if v92 != int32(1) {
													v106 = int32(0)
													v107 = v97
													for {
														v118 = v107 + v28
														v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
														v120 = v107 + v93
														v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
														v122 = v119 ^ v121
														*(*uint8)(unsafe.Add(mBase, uint32(v118))) = uint8(v122)
														v124 = v107 + v18
														v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
														v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
														v127 = v125 ^ v126
														*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v127)
														v130 = v107 | int32(1)
														v131 = v28 + v130
														v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
														v133 = v93 + v130
														v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
														v135 = v132 ^ v134
														*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v135)
														v137 = v18 + v130
														v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
														v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
														v140 = v138 ^ v139
														*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v140)
														v142 = int32(2)
														v143 = v107 + v142
														v145 = v106 + v142
														if v145 != v92&int32(-2) {
															v106 = v145
															v107 = v143
															continue
														} else {
															break
														}
														break
													}
													if v92&int32(1) == int32(0) {
													} else {
														v151 = v143
														v162 = v151 + v28
														v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
														v164 = v151 + v93
														v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
														v166 = v163 ^ v165
														*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v166)
														v168 = v151 + v18
														v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
														v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
														v171 = v169 ^ v170
														*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v171)
													}
												} else {
													v151 = v97
													v162 = v151 + v28
													v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
													v164 = v151 + v93
													v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
													v166 = v163 ^ v165
													*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v166)
													v168 = v151 + v18
													v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
													v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
													v171 = v169 ^ v170
													*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v171)
												}
											}
											v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v187 = F_pg_cryptohash_init(m, v186)
											mBase = m.M
											if int32(0) <= v187 {
												v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												v192 = F_pg_cryptohash_update(m, v190, v28, v191)
												mBase = m.M
												if int32(0) <= v192 {
													v214 = int32(0)
													if v94 == v214 {
														v229 = v214
														return v229
													} else {
														v217 = v214
														F_pfree(m, v94)
														mBase = m.M
														v219 = m.ExcPending
														if v219 != 0 {
															return int32(0)
														} else {
															v229 = v217
															return v229
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
													v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													if v197 == int32(0) {
														v212 = int32(_a_F_pg_hmac_init_0)
													} else {
														v204 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
														if v204 == int32(1) {
															v207 = int32(_a_F_pg_hmac_init_1)
														} else {
															v207 = int32(_a_F_pg_hmac_init_2)
														}
														if v204 == int32(2) {
															v210 = int32(_a_F_pg_hmac_init_0)
														} else {
															v210 = v207
														}
														v212 = v210
													}
													*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v212
													if v94 != 0 {
														v217 = v14
														F_pfree(m, v94)
														mBase = m.M
														v219 = m.ExcPending
														if v219 != 0 {
															return int32(0)
														} else {
															v229 = v217
															return v229
														}
													} else {
														v229 = v14
														return v229
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
												v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												if v197 == int32(0) {
													v212 = int32(_a_F_pg_hmac_init_0)
												} else {
													v204 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
													if v204 == int32(1) {
														v207 = int32(_a_F_pg_hmac_init_1)
													} else {
														v207 = int32(_a_F_pg_hmac_init_2)
													}
													if v204 == int32(2) {
														v210 = int32(_a_F_pg_hmac_init_0)
													} else {
														v210 = v207
													}
													v212 = v210
												}
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v212
												if v94 != 0 {
													v217 = v14
													F_pfree(m, v94)
													mBase = m.M
													v219 = m.ExcPending
													if v219 != 0 {
														return int32(0)
													} else {
														v229 = v217
														return v229
													}
												} else {
													v229 = v14
													return v229
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(2)
										if v47 == int32(0) {
											v82 = int32(_a_F_pg_hmac_init_0)
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
											if v74 == int32(1) {
												v77 = int32(_a_F_pg_hmac_init_1)
											} else {
												v77 = int32(_a_F_pg_hmac_init_2)
											}
											if v74 == int32(2) {
												v80 = int32(_a_F_pg_hmac_init_0)
											} else {
												v80 = v77
											}
											v82 = v80
										}
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v82
										F_pg_cryptohash_free(m, v47)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v34)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												return int32(-1)
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
func F_pg_input_error_info(m *base.Module, l0 int32) int64 {
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
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_pg_input_error_info[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = v19
			v22 = *(*int64)(unsafe.Add(mBase, _c_F_pg_input_error_info[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v22
			v27 = F_get_call_result_type(m, l0, int32(0), v8+int32(60))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				if v27 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v8)+69)) = uint8(v31)
					v35 = F_pg_input_is_valid_common(m, l0, v11, v16, v8-int32(-64))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						if v35 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(16843009)
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
							v120 = F_heap_form_tuple(m, v115, v8+int32(16), v8+int32(12))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int64(0)
							} else {
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
								v123 = F_HeapTupleHeaderGetDatum(m, v122)
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int64(0)
								} else {
									m.G0 = v8 + int32(80)
									return v123
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+32))
							v43 = F_cstring_to_text(m, v42)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = base.I64_extend_i32_u(v43)
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
								if v48 != 0 {
									v49 = F_cstring_to_text(m, v48)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = base.I64_extend_i32_u(v49)
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
										v56 = v53
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+44))
										if v57 != 0 {
											v58 = F_cstring_to_text(m, v57)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int64(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = base.I64_extend_i32_u(v58)
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
												v65 = v62
												v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
												v67 = int32(_a_F_pg_input_error_info_0)
												v68 = int32(63)
												v70 = int32(48)
												v71 = v66&v68 + v70
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[2])) = uint8(v71)
												v79 = int32(base.Ui32(v66)>>(uint(int32(24))%32))&v68 + v70
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[3])) = uint8(v79)
												v87 = int32(base.Ui32(v66)>>(uint(int32(18))%32))&v68 + v70
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[4])) = uint8(v87)
												v95 = int32(base.Ui32(v66)>>(uint(int32(12))%32))&v68 + v70
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[5])) = uint8(v95)
												v103 = int32(base.Ui32(v66)>>(uint(int32(6))%32))&v68 + v70
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[6])) = uint8(v103)
												v106 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[7])) = uint8(v106)
												v109 = F_cstring_to_text(m, v67)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int64(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = base.I64_extend_i32_u(v109)
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
													v120 = F_heap_form_tuple(m, v115, v8+int32(16), v8+int32(12))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int64(0)
													} else {
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
														v123 = F_HeapTupleHeaderGetDatum(m, v122)
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return int64(0)
														} else {
															m.G0 = v8 + int32(80)
															return v123
														}
													}
												}
											}
										} else {
											v63 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v63)
											v65 = v56
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
											v67 = int32(_a_F_pg_input_error_info_0)
											v68 = int32(63)
											v70 = int32(48)
											v71 = v66&v68 + v70
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[2])) = uint8(v71)
											v79 = int32(base.Ui32(v66)>>(uint(int32(24))%32))&v68 + v70
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[3])) = uint8(v79)
											v87 = int32(base.Ui32(v66)>>(uint(int32(18))%32))&v68 + v70
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[4])) = uint8(v87)
											v95 = int32(base.Ui32(v66)>>(uint(int32(12))%32))&v68 + v70
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[5])) = uint8(v95)
											v103 = int32(base.Ui32(v66)>>(uint(int32(6))%32))&v68 + v70
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[6])) = uint8(v103)
											v106 = int32(0)
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[7])) = uint8(v106)
											v109 = F_cstring_to_text(m, v67)
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int64(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = base.I64_extend_i32_u(v109)
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
												v120 = F_heap_form_tuple(m, v115, v8+int32(16), v8+int32(12))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int64(0)
												} else {
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
													v123 = F_HeapTupleHeaderGetDatum(m, v122)
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return int64(0)
													} else {
														m.G0 = v8 + int32(80)
														return v123
													}
												}
											}
										}
									}
								} else {
									v54 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v8)+13)) = uint8(v54)
									v56 = v47
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+44))
									if v57 != 0 {
										v58 = F_cstring_to_text(m, v57)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = base.I64_extend_i32_u(v58)
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
											v65 = v62
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
											v67 = int32(_a_F_pg_input_error_info_0)
											v68 = int32(63)
											v70 = int32(48)
											v71 = v66&v68 + v70
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[2])) = uint8(v71)
											v79 = int32(base.Ui32(v66)>>(uint(int32(24))%32))&v68 + v70
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[3])) = uint8(v79)
											v87 = int32(base.Ui32(v66)>>(uint(int32(18))%32))&v68 + v70
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[4])) = uint8(v87)
											v95 = int32(base.Ui32(v66)>>(uint(int32(12))%32))&v68 + v70
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[5])) = uint8(v95)
											v103 = int32(base.Ui32(v66)>>(uint(int32(6))%32))&v68 + v70
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[6])) = uint8(v103)
											v106 = int32(0)
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[7])) = uint8(v106)
											v109 = F_cstring_to_text(m, v67)
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int64(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = base.I64_extend_i32_u(v109)
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
												v120 = F_heap_form_tuple(m, v115, v8+int32(16), v8+int32(12))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int64(0)
												} else {
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
													v123 = F_HeapTupleHeaderGetDatum(m, v122)
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return int64(0)
													} else {
														m.G0 = v8 + int32(80)
														return v123
													}
												}
											}
										}
									} else {
										v63 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v63)
										v65 = v56
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
										v67 = int32(_a_F_pg_input_error_info_0)
										v68 = int32(63)
										v70 = int32(48)
										v71 = v66&v68 + v70
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[2])) = uint8(v71)
										v79 = int32(base.Ui32(v66)>>(uint(int32(24))%32))&v68 + v70
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[3])) = uint8(v79)
										v87 = int32(base.Ui32(v66)>>(uint(int32(18))%32))&v68 + v70
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[4])) = uint8(v87)
										v95 = int32(base.Ui32(v66)>>(uint(int32(12))%32))&v68 + v70
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[5])) = uint8(v95)
										v103 = int32(base.Ui32(v66)>>(uint(int32(6))%32))&v68 + v70
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[6])) = uint8(v103)
										v106 = int32(0)
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[7])) = uint8(v106)
										v109 = F_cstring_to_text(m, v67)
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = base.I64_extend_i32_u(v109)
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
											v120 = F_heap_form_tuple(m, v115, v8+int32(16), v8+int32(12))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int64(0)
											} else {
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
												v123 = F_HeapTupleHeaderGetDatum(m, v122)
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int64(0)
												} else {
													m.G0 = v8 + int32(80)
													return v123
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v132 = m.ExcPending
					if v132 != 0 {
						return int64(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pg_input_error_info_1), int32(0))
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_input_error_info_2), int32(699), int32(_a_F_pg_input_error_info_3))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return int64(0)
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
func F_pg_input_is_valid_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
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
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_text_to_cstring(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
		if v17 == int32(0) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
			v22 = F_MemoryContextAlloc(m, v20, int32(48))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v22
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
				v28 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v30 == v28 {
					v78 = v28
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
					if v36 == int32(0) {
						v78 = v28
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
						v41 = v39 - int32(11)
						if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v41))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v41)%32))&int32(1) == int32(0))|int32(0) != 0 {
							v78 = v28
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v41<<(uint(int32(2))%32))+uint32(_c_F_pg_input_is_valid_common[0])))
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v36+v56)))
							if v58 == int32(0) {
								v78 = v28
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								if v61 <= int32(1) {
									v78 = v28
								} else {
									v63 = int32(1)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+int32(4))))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
									switch v69 - int32(7) {
									case 0:
										v78 = v63
									case 1:
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
										if v72 == int32(0) {
											v78 = v63
										} else {
											v78 = int32(0)
										}
									default:
										v78 = int32(0)
									}
								}
							}
						}
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)) = uint8(v78)
				v80 = v27
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				if v81 != 0 {
					v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+8)))
					if v82 != 0 {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
						v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(8))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 + int32(16)
							return v117
						}
					} else {
						v83 = F_text_to_cstring(m, l2)
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							v85 = int32(4)
							v90 = F_parseTypeString(m, v83, v10+v85, v80+v85, int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
								if v92 == v93 {
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
									v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(8))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return v117
									}
								} else {
									F_getTypeInputInfo(m, v92, v80+int32(12), v80+int32(16))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int32(0)
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
										v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
										F_fmgr_info_cxt(m, v101, v80+int32(20), v105)
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int32(0)
										} else {
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
											v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(8))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												m.G0 = v10 + int32(16)
												return v117
											}
										}
									}
								}
							}
						}
					}
				} else {
					v83 = F_text_to_cstring(m, l2)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						v85 = int32(4)
						v90 = F_parseTypeString(m, v83, v10+v85, v80+v85, int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							if v92 == v93 {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
								v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(8))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return v117
								}
							} else {
								F_getTypeInputInfo(m, v92, v80+int32(12), v80+int32(16))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
									F_fmgr_info_cxt(m, v101, v80+int32(20), v105)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
										v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(8))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(16)
											return v117
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v80 = v17
			v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
			if v81 != 0 {
				v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+8)))
				if v82 != 0 {
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
					v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(8))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(16)
						return v117
					}
				} else {
					v83 = F_text_to_cstring(m, l2)
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						v85 = int32(4)
						v90 = F_parseTypeString(m, v83, v10+v85, v80+v85, int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							if v92 == v93 {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
								v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(8))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return v117
								}
							} else {
								F_getTypeInputInfo(m, v92, v80+int32(12), v80+int32(16))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
									F_fmgr_info_cxt(m, v101, v80+int32(20), v105)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
										v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(8))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(16)
											return v117
										}
									}
								}
							}
						}
					}
				}
			} else {
				v83 = F_text_to_cstring(m, l2)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					v85 = int32(4)
					v90 = F_parseTypeString(m, v83, v10+v85, v80+v85, int32(0))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
						if v92 == v93 {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
							v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(8))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v117
							}
						} else {
							F_getTypeInputInfo(m, v92, v80+int32(12), v80+int32(16))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
								v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
								F_fmgr_info_cxt(m, v101, v80+int32(20), v105)
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int32(0)
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
									v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(8))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return v117
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
func F_pg_is_ascii(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = l0
	for {
		v5 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3))))
		if int32(0) < v5 {
			v3 = v3 + int32(1)
			continue
		} else {
			break
		}
		break
	}
	return base.B2i32(v5 == int32(0))
}
func F_pg_listening_channels(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v6 != 0 {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
		if v35 == int32(0) {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int64(0)
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = int32(2)
				v59 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
				return int64(0)
			}
		} else {
			v38 = F_hash_seq_search(m, v35)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int64(0)
			} else {
				if v38 == int32(0) {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = int32(2)
						v59 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
						return int64(0)
					}
				} else {
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v42 + int64(1)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = int32(1)
					v49 = F_cstring_to_text(m, v38)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v49)
					}
				}
			}
		}
	} else {
		v7 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_pg_listening_channels[0]))
			if v12 != 0 {
				v13 = int32(_a_F_pg_listening_channels_0)
				v14 = *(*int32)(unsafe.Add(mBase, _c_F_pg_listening_channels[1]))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_pg_listening_channels[1])) = v16
				v19 = F_palloc(m, int32(20))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_listening_channels[0]))
					F_hash_seq_init(m, v19, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v19
						*(*int32)(unsafe.Add(mBase, _c_F_pg_listening_channels[1])) = v14
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
						if v35 == int32(0) {
							F_end_MultiFuncCall(m, l0)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = int32(2)
								v59 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
								return int64(0)
							}
						} else {
							v38 = F_hash_seq_search(m, v35)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
							} else {
								if v38 == int32(0) {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int64(0)
									} else {
										v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = int32(2)
										v59 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
										return int64(0)
									}
								} else {
									v42 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
									*(*int64)(unsafe.Add(mBase, uint32(v34))) = v42 + int64(1)
									v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = int32(1)
									v49 = F_cstring_to_text(m, v38)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int64(0)
									} else {
										return base.I64_extend_i32_u(v49)
									}
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(0)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
				if v35 == int32(0) {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int64(0)
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = int32(2)
						v59 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
						return int64(0)
					}
				} else {
					v38 = F_hash_seq_search(m, v35)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						if v38 == int32(0) {
							F_end_MultiFuncCall(m, l0)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = int32(2)
								v59 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
								return int64(0)
							}
						} else {
							v42 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
							*(*int64)(unsafe.Add(mBase, uint32(v34))) = v42 + int64(1)
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = int32(1)
							v49 = F_cstring_to_text(m, v38)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v49)
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_lock_status(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v313 int32
	_ = v313
	var v317 int64
	_ = v317
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v344 int64
	_ = v344
	var v346 int32
	_ = v346
	var v348 int64
	_ = v348
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int64
	_ = v406
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int64
	_ = v657
	var v658 int64
	_ = v658
	var v661 int64
	_ = v661
	var v662 int64
	_ = v662
	var v663 int64
	_ = v663
	var v664 int64
	_ = v664
	var v665 int64
	_ = v665
	var v666 int64
	_ = v666
	var v667 int64
	_ = v667
	var v668 int64
	_ = v668
	var v669 int64
	_ = v669
	var v670 int64
	_ = v670
	var v671 int64
	_ = v671
	var v672 int64
	_ = v672
	var v673 int64
	_ = v673
	var v674 int64
	_ = v674
	var v675 int64
	_ = v675
	var v676 int64
	_ = v676
	var v677 int64
	_ = v677
	var v678 int64
	_ = v678
	var v679 int64
	_ = v679
	var v680 int64
	_ = v680
	var v681 int64
	_ = v681
	var v682 int64
	_ = v682
	var v683 int64
	_ = v683
	var v684 int64
	_ = v684
	var v685 int64
	_ = v685
	var v686 int64
	_ = v686
	var v687 int64
	_ = v687
	var v688 int64
	_ = v688
	var v689 int64
	_ = v689
	var v690 int64
	_ = v690
	var v691 int64
	_ = v691
	var v723 int64
	_ = v723
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int64
	_ = v767
	var v769 int64
	_ = v769
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int64
	_ = v791
	var v794 int64
	_ = v794
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1054 int64
	_ = v1054
	var v1055 int64
	_ = v1055
	var v1058 int64
	_ = v1058
	var v1059 int64
	_ = v1059
	var v1060 int64
	_ = v1060
	var v1061 int64
	_ = v1061
	var v1062 int64
	_ = v1062
	var v1063 int64
	_ = v1063
	var v1064 int64
	_ = v1064
	var v1065 int64
	_ = v1065
	var v1066 int64
	_ = v1066
	var v1067 int64
	_ = v1067
	var v1068 int64
	_ = v1068
	var v1069 int64
	_ = v1069
	var v1070 int64
	_ = v1070
	var v1071 int64
	_ = v1071
	var v1072 int64
	_ = v1072
	var v1073 int64
	_ = v1073
	var v1074 int64
	_ = v1074
	var v1075 int64
	_ = v1075
	var v1076 int64
	_ = v1076
	var v1077 int64
	_ = v1077
	var v1078 int64
	_ = v1078
	var v1079 int64
	_ = v1079
	var v1080 int64
	_ = v1080
	var v1081 int64
	_ = v1081
	var v1082 int64
	_ = v1082
	var v1083 int64
	_ = v1083
	var v1084 int64
	_ = v1084
	var v1085 int64
	_ = v1085
	var v1086 int64
	_ = v1086
	var v1087 int64
	_ = v1087
	var v1088 int64
	_ = v1088
	var v1120 int64
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int64
	_ = v1163
	var v1165 int64
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1356 int32
	_ = v1356
	var v1359 int64
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1418 int32
	_ = v1418
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1452 int64
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1464 int64
	_ = v1464
	var v1466 int64
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1474 int64
	_ = v1474
	var v1476 int64
	_ = v1476
	var v1478 int64
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1494 int32
	_ = v1494
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1515 int64
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1527 int64
	_ = v1527
	var v1529 int64
	_ = v1529
	var v1531 int64
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1539 int64
	_ = v1539
	var v1541 int64
	_ = v1541
	var v1543 int64
	_ = v1543
	var v1547 int32
	_ = v1547
	var v1552 int64
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1586 int64
	_ = v1586
	var v1588 int64
	_ = v1588
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int64
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1606 int64
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1646 int64
	_ = v1646
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1674 int64
	_ = v1674
	var v1676 int64
	_ = v1676
	var v1680 int64
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1686 int64
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1694 int64
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1722 int32
	_ = v1722
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int64
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int64
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1744 int32
	_ = v1744
	var v1764 int64
	_ = v1764
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(272)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v25 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1328)+16))
	goto L186
L4:
	;
	return int64(0)
L5:
	;
	v32 = int32(_a_F_pg_lock_status_0)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[0]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[0])) = v35
	v38 = F_CreateTemplateTupleDesc(m, int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v38, int32(1), int32(_a_F_pg_lock_status_1), int32(25), int32(-1), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v38, int32(2), int32(_a_F_pg_lock_status_2), int32(26), int32(-1), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v38, int32(3), int32(_a_F_pg_lock_status_3), int32(26), int32(-1), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v38, int32(4), int32(_a_F_pg_lock_status_4), int32(23), int32(-1), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v38, int32(5), int32(_a_F_pg_lock_status_5), int32(21), int32(-1), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	F_TupleDescInitEntry(m, v38, int32(6), int32(_a_F_pg_lock_status_6), int32(25), int32(-1), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_TupleDescInitEntry(m, v38, int32(7), int32(_a_F_pg_lock_status_7), int32(28), int32(-1), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_TupleDescInitEntry(m, v38, int32(8), int32(_a_F_pg_lock_status_8), int32(26), int32(-1), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_TupleDescInitEntry(m, v38, int32(9), int32(_a_F_pg_lock_status_9), int32(26), int32(-1), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_TupleDescInitEntry(m, v38, int32(10), int32(_a_F_pg_lock_status_10), int32(21), int32(-1), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	F_TupleDescInitEntry(m, v38, int32(11), int32(_a_F_pg_lock_status_11), int32(25), int32(-1), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F_TupleDescInitEntry(m, v38, int32(12), int32(_a_F_pg_lock_status_12), int32(23), int32(-1), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_TupleDescInitEntry(m, v38, int32(13), int32(_a_F_pg_lock_status_13), int32(25), int32(-1), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_TupleDescInitEntry(m, v38, int32(14), int32(_a_F_pg_lock_status_14), int32(16), int32(-1), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_TupleDescInitEntry(m, v38, int32(15), int32(_a_F_pg_lock_status_15), int32(16), int32(-1), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_TupleDescInitEntry(m, v38, int32(16), int32(_a_F_pg_lock_status_16), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v152 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v152 < v161 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v239 = F_BlessTupleDesc(m, v38)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L42
	}
L24:
	;
	v165 = v38 + int32(28)
	v172 = v152
	v173 = v161
	v175 = v152
	goto L28
L25:
	;
	v229 = v152
	v236 = v161
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v229
	goto L23
L27:
	;
	v229 = v223
	v236 = v202
	goto L26
L28:
	;
	v181 = v165 + v161<<(uint(int32(3))%32) + v172*int32(100)
	v184 = v165 + v172<<(uint(int32(3))%32)
	if v161 != v173 {
		v202 = v173
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v223 = v161
	goto L27
L30:
	;
	v203 = int32(*(*int16)(unsafe.Add(mBase, uint32(v184)+2)))
	if v203 <= int32(0) {
		v223 = v172
		goto L27
	} else {
		goto L38
	}
L31:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+7)))
	if v186 != int32(118) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v202 = v172
	goto L30
L33:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+4)))
	if v189 != int32(1) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+6)))
	if v192&int32(6) != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v195 = int32(*(*int16)(unsafe.Add(mBase, uint32(v184)+2)))
	if v195 <= int32(0) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+90)))
	if v198 != int32(118) {
		v202 = v161
		goto L30
	} else {
		goto L37
	}
L37:
	;
	goto L32
L38:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+90)))
	if v206 == int32(118) {
		v223 = v172
		goto L27
	} else {
		goto L39
	}
L39:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+5)))
	v215 = (v175 + v209 - int32(1)) & (int32(0) - v209)
	if int32(_a_F_pg_lock_status_17) < v215 {
		v223 = v172
		goto L27
	} else {
		goto L40
	}
L40:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v184))) = uint16(v215)
	v221 = v172 + int32(1)
	if v221 != v161 {
		v172 = v221
		v173 = v202
		v175 = v215 + v203
		goto L28
	} else {
		goto L41
	}
L41:
	;
	goto L29
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v239
	v243 = F_palloc(m, int32(16))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v243
	v246 = m.G0
	v248 = v246 - int32(32)
	m.G0 = v248
	v251 = F_palloc(m, int32(8))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[1]))
	v256 = F_palloc_mul(m, int32(56), v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = v256
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[2]))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	if v261 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v263 = v260
	v266 = v2
	v268 = v255
	v273 = v2
	goto L49
L47:
	;
	v526 = v2
	v528 = v255
	goto L48
L48:
	;
	v542 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v546 = F_LWLockAcquire(m, v542+int32(_a_F_pg_lock_status_18), int32(1))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L83
	}
L49:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v284 = v281 + v273*int32(768)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	if v285 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v526 = v503
	v528 = v505
	goto L48
L51:
	;
	v287 = v284 + int32(548)
	v289 = F_LWLockAcquire(m, v287, int32(1))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	v500 = v263
	v503 = v266
	v505 = v268
	goto L53
L53:
	;
	v519 = v273 + int32(1)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v500)+16))
	if base.Ui32(v519) < base.Ui32(v520) {
		v263 = v500
		v266 = v503
		v268 = v505
		v273 = v519
		goto L49
	} else {
		goto L82
	}
L54:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[4]))
	if v292 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v295 = v292
	v298 = v266
	v300 = v268
	v301 = int32(0)
	goto L58
L56:
	;
	v437 = v266
	v439 = v268
	goto L57
L57:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+572)))
	if v452 != 0 {
		goto L74
	} else {
		goto L75
	}
L58:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v284)+564))
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v313+v301<<(uint(int32(3))%32))))
	if v317 != int64(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v437 = v415
	v439 = v417
	goto L57
L60:
	;
	v331 = v298
	v333 = v300
	v344 = int64(0)
	goto L63
L61:
	;
	v412 = v295
	v415 = v298
	v417 = v300
	goto L62
L62:
	;
	v431 = v301 + int32(1)
	if base.Ui32(v431) < base.Ui32(v412) {
		v295 = v412
		v298 = v415
		v300 = v417
		v301 = v431
		goto L58
	} else {
		goto L73
	}
L63:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v284)+564))
	v348 = *(*int64)(unsafe.Add(mBase, uint32(v346+v301&int32(268435455)<<(uint(int32(3))%32))))
	v354 = base.I32_wrap_i64(int64(base.Ui64(v348)>>(uint(v344*int64(3))%64))) & int32(7)
	if v354 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[4]))
	v412 = v410
	v415 = v401
	v417 = v402
	goto L62
L65:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	if v333 <= v331 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v401 = v331
	v402 = v333
	goto L67
L67:
	;
	v406 = v344 + int64(1)
	if v406 != int64(16) {
		v331 = v401
		v333 = v402
		v344 = v406
		goto L63
	} else {
		goto L72
	}
L68:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[1]))
	v359 = v358 + v333
	v362 = F_repalloc(m, v355, v359*int32(56))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L4
	} else {
		goto L71
	}
L69:
	;
	v365 = v355
	v366 = v333
	goto L70
L70:
	;
	v369 = v365 + v331*int32(56)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v284)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v370
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v284)+568))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v372+v301<<(uint(int32(6))%32)+base.I32_wrap_i64(v344)<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+20)) = int32(0)
	v381 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+16)) = v354 << (uint(v381) % 32)
	*(*int64)(unsafe.Add(mBase, uint32(v369)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v378
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v284)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+24)) = v387
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v284)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+28)) = v389
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v369)+48)) = uint8(v381)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+44)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v369)+40)) = v391
	*(*int64)(unsafe.Add(mBase, uint32(v369)+32)) = int64(0)
	v401 = v331 + v381
	v402 = v366
	goto L67
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = v362
	v365 = v362
	v366 = v359
	goto L70
L72:
	;
	goto L64
L73:
	;
	goto L59
L74:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	if v439 <= v437 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v491 = v437
	v492 = v439
	goto L76
L76:
	;
	F_LWLockRelease(m, v287)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L81
	}
L77:
	;
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[1]))
	v457 = v456 + v439
	v460 = F_repalloc(m, v453, v457*int32(56))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L80
	}
L78:
	;
	v463 = v453
	v464 = v439
	goto L79
L79:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v284)+40))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v284)+576))
	v469 = v463 + v437*int32(56)
	*(*int64)(unsafe.Add(mBase, uint32(v469)+16)) = int64(128)
	*(*int64)(unsafe.Add(mBase, uint32(v469)+8)) = int64(73746443898191872)
	*(*int32)(unsafe.Add(mBase, uint32(v469)+4)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v469))) = v465
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v284)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v469)+24)) = v476
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v284)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v469)+28)) = v478
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	v481 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+48)) = uint8(v481)
	*(*int32)(unsafe.Add(mBase, uint32(v469)+44)) = v480
	*(*int32)(unsafe.Add(mBase, uint32(v469)+40)) = v480
	*(*int64)(unsafe.Add(mBase, uint32(v469)+32)) = int64(0)
	v491 = v437 + v481
	v492 = v464
	goto L76
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = v460
	v463 = v460
	v464 = v457
	goto L79
L81:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[2]))
	v500 = v498
	v503 = v491
	v505 = v492
	goto L53
L82:
	;
	goto L50
L83:
	;
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v553 = F_LWLockAcquire(m, v549+int32(_a_F_pg_lock_status_19), int32(1))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v560 = F_LWLockAcquire(m, v556+int32(_a_F_pg_lock_status_20), int32(1))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v567 = F_LWLockAcquire(m, v563+int32(_a_F_pg_lock_status_21), int32(1))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v574 = F_LWLockAcquire(m, v570+int32(_a_F_pg_lock_status_22), int32(1))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	v577 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v581 = F_LWLockAcquire(m, v577+int32(_a_F_pg_lock_status_23), int32(1))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	v584 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v588 = F_LWLockAcquire(m, v584+int32(_a_F_pg_lock_status_24), int32(1))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v591 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v595 = F_LWLockAcquire(m, v591+int32(_a_F_pg_lock_status_25), int32(1))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	v598 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v602 = F_LWLockAcquire(m, v598+int32(_a_F_pg_lock_status_26), int32(1))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	v605 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v609 = F_LWLockAcquire(m, v605+int32(_a_F_pg_lock_status_27), int32(1))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v612 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v616 = F_LWLockAcquire(m, v612+int32(_a_F_pg_lock_status_28), int32(1))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v623 = F_LWLockAcquire(m, v619+int32(_a_F_pg_lock_status_29), int32(1))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	v626 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v630 = F_LWLockAcquire(m, v626+int32(_a_F_pg_lock_status_30), int32(1))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	v633 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v637 = F_LWLockAcquire(m, v633+int32(_a_F_pg_lock_status_31), int32(1))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v644 = F_LWLockAcquire(m, v640+int32(_a_F_pg_lock_status_32), int32(1))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v647 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v651 = F_LWLockAcquire(m, v647+int32(_a_F_pg_lock_status_33), int32(1))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	v654 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[5]))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v654)))
	v657 = *(*int64)(unsafe.Add(mBase, uint32(v656)+8))
	v658 = *(*int64)(unsafe.Add(mBase, uint32(v656)+808))
	if v658 != int64(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v725 = v526 + base.I32_wrap_i64(v723)
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v725
	if v528 < v725 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v661 = *(*int64)(unsafe.Add(mBase, uint32(v656)+752))
	v662 = *(*int64)(unsafe.Add(mBase, uint32(v656)+728))
	v663 = *(*int64)(unsafe.Add(mBase, uint32(v656)+704))
	v664 = *(*int64)(unsafe.Add(mBase, uint32(v656)+680))
	v665 = *(*int64)(unsafe.Add(mBase, uint32(v656)+656))
	v666 = *(*int64)(unsafe.Add(mBase, uint32(v656)+632))
	v667 = *(*int64)(unsafe.Add(mBase, uint32(v656)+608))
	v668 = *(*int64)(unsafe.Add(mBase, uint32(v656)+584))
	v669 = *(*int64)(unsafe.Add(mBase, uint32(v656)+560))
	v670 = *(*int64)(unsafe.Add(mBase, uint32(v656)+536))
	v671 = *(*int64)(unsafe.Add(mBase, uint32(v656)+512))
	v672 = *(*int64)(unsafe.Add(mBase, uint32(v656)+488))
	v673 = *(*int64)(unsafe.Add(mBase, uint32(v656)+464))
	v674 = *(*int64)(unsafe.Add(mBase, uint32(v656)+440))
	v675 = *(*int64)(unsafe.Add(mBase, uint32(v656)+416))
	v676 = *(*int64)(unsafe.Add(mBase, uint32(v656)+392))
	v677 = *(*int64)(unsafe.Add(mBase, uint32(v656)+368))
	v678 = *(*int64)(unsafe.Add(mBase, uint32(v656)+344))
	v679 = *(*int64)(unsafe.Add(mBase, uint32(v656)+320))
	v680 = *(*int64)(unsafe.Add(mBase, uint32(v656)+296))
	v681 = *(*int64)(unsafe.Add(mBase, uint32(v656)+272))
	v682 = *(*int64)(unsafe.Add(mBase, uint32(v656)+248))
	v683 = *(*int64)(unsafe.Add(mBase, uint32(v656)+224))
	v684 = *(*int64)(unsafe.Add(mBase, uint32(v656)+200))
	v685 = *(*int64)(unsafe.Add(mBase, uint32(v656)+176))
	v686 = *(*int64)(unsafe.Add(mBase, uint32(v656)+152))
	v687 = *(*int64)(unsafe.Add(mBase, uint32(v656)+128))
	v688 = *(*int64)(unsafe.Add(mBase, uint32(v656)+104))
	v689 = *(*int64)(unsafe.Add(mBase, uint32(v656)+80))
	v690 = *(*int64)(unsafe.Add(mBase, uint32(v656)+56))
	v691 = *(*int64)(unsafe.Add(mBase, uint32(v656)+32))
	v723 = v661 + (v662 + (v663 + (v664 + (v665 + (v666 + (v667 + (v668 + (v669 + (v670 + (v671 + (v672 + (v673 + (v674 + (v675 + (v676 + (v677 + (v678 + (v679 + (v680 + (v681 + (v682 + (v683 + (v684 + (v685 + (v686 + (v687 + (v688 + (v689 + (v690 + (v691 + v657))))))))))))))))))))))))))))))
	goto L102
L101:
	;
	v723 = v657
	goto L102
L102:
	;
	goto L99
L103:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	v731 = F_repalloc(m, v728, v725*int32(56))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v735 = v248 + int32(12)
	v737 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[5]))
	F_hash_seq_init(m, v735, v737)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L4
	} else {
		goto L107
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+4)) = v731
	goto L105
L107:
	;
	v740 = F_hash_seq_search(m, v735)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	if v740 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v743 = v740
	v746 = v526
	goto L112
L110:
	;
	goto L111
L111:
	;
	v822 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v822+int32(_a_F_pg_lock_status_33))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L4
	} else {
		goto L119
	}
L112:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v743)+4))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	v765 = v762 + v746*int32(56)
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v743)))
	v767 = *(*int64)(unsafe.Add(mBase, uint32(v766)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v765)+8)) = v767
	v769 = *(*int64)(unsafe.Add(mBase, uint32(v766)))
	*(*int64)(unsafe.Add(mBase, uint32(v765))) = v769
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v743)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v765)+16)) = v771
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v761)+384))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v743)))
	if v773 == v774 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L111
L114:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v761)+400))
	v778 = v776
	goto L116
L115:
	;
	v778 = int32(0)
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v765)+20)) = v778
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v761)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v765)+24)) = v780
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v761)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v765)+28)) = v782
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v761)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v765)+40)) = v784
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v743)+8))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v786)+12))
	v788 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v765)+48)) = uint8(v788)
	*(*int32)(unsafe.Add(mBase, uint32(v765)+44)) = v787
	v791 = int64(0)
	v794 = base.AtomicRmwCmpxchg64(m, v761, int32(408), v791, v791)
	*(*int64)(unsafe.Add(mBase, uint32(v765)+32)) = v794
	v800 = F_hash_seq_search(m, v248+int32(12))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	if v800 != 0 {
		v743 = v800
		v746 = v746 + int32(1)
		goto L112
	} else {
		goto L118
	}
L118:
	;
	goto L113
L119:
	;
	v828 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v828+int32(_a_F_pg_lock_status_32))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	v834 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v834+int32(_a_F_pg_lock_status_31))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	v840 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v840+int32(_a_F_pg_lock_status_30))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	v846 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v846+int32(_a_F_pg_lock_status_29))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v852 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v852+int32(_a_F_pg_lock_status_28))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v858 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v858+int32(_a_F_pg_lock_status_27))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v864 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v864+int32(_a_F_pg_lock_status_26))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	v870 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v870+int32(_a_F_pg_lock_status_25))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v876 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v876+int32(_a_F_pg_lock_status_24))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	v882 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v882+int32(_a_F_pg_lock_status_23))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	v888 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v888+int32(_a_F_pg_lock_status_22))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v894 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v894+int32(_a_F_pg_lock_status_21))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	v900 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v900+int32(_a_F_pg_lock_status_20))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v906 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v906+int32(_a_F_pg_lock_status_19))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	v912 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v912+int32(_a_F_pg_lock_status_18))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	v917 = int32(32)
	m.G0 = v248 + v917
	v920 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v243)+4)) = v920
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v251
	v924 = m.G0
	v926 = v924 - v917
	m.G0 = v926
	v929 = F_palloc(m, int32(12))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L4
	} else {
		goto L135
	}
L135:
	;
	v932 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v936 = F_LWLockAcquire(m, v932+int32(_a_F_pg_lock_status_34), int32(1))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v939 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v943 = F_LWLockAcquire(m, v939+int32(_a_F_pg_lock_status_35), int32(1))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	v946 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v950 = F_LWLockAcquire(m, v946+int32(_a_F_pg_lock_status_36), int32(1))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	v953 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v957 = F_LWLockAcquire(m, v953+int32(_a_F_pg_lock_status_37), int32(1))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	v960 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v964 = F_LWLockAcquire(m, v960+int32(_a_F_pg_lock_status_38), int32(1))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	v967 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v971 = F_LWLockAcquire(m, v967+int32(_a_F_pg_lock_status_39), int32(1))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	v974 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v978 = F_LWLockAcquire(m, v974+int32(_a_F_pg_lock_status_40), int32(1))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	v981 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v985 = F_LWLockAcquire(m, v981+int32(_a_F_pg_lock_status_41), int32(1))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	v988 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v992 = F_LWLockAcquire(m, v988+int32(_a_F_pg_lock_status_42), int32(1))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	v995 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v999 = F_LWLockAcquire(m, v995+int32(_a_F_pg_lock_status_43), int32(1))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v1006 = F_LWLockAcquire(m, v1002+int32(_a_F_pg_lock_status_44), int32(1))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v1013 = F_LWLockAcquire(m, v1009+int32(_a_F_pg_lock_status_45), int32(1))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v1020 = F_LWLockAcquire(m, v1016+int32(_a_F_pg_lock_status_46), int32(1))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v1027 = F_LWLockAcquire(m, v1023+int32(_a_F_pg_lock_status_47), int32(1))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v1034 = F_LWLockAcquire(m, v1030+int32(_a_F_pg_lock_status_48), int32(1))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v1041 = F_LWLockAcquire(m, v1037+int32(_a_F_pg_lock_status_49), int32(1))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v1048 = F_LWLockAcquire(m, v1044+int32(3584), int32(1))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[6]))
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v1051)))
	v1054 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+8))
	v1055 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+808))
	if v1055 != int64(0) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v1121 = base.I32_wrap_i64(v1120)
	*(*int32)(unsafe.Add(mBase, uint32(v929))) = v1121
	v1124 = F_palloc_mul(m, int32(16), v1121)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L4
	} else {
		goto L157
	}
L154:
	;
	v1058 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+752))
	v1059 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+728))
	v1060 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+704))
	v1061 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+680))
	v1062 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+656))
	v1063 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+632))
	v1064 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+608))
	v1065 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+584))
	v1066 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+560))
	v1067 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+536))
	v1068 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+512))
	v1069 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+488))
	v1070 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+464))
	v1071 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+440))
	v1072 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+416))
	v1073 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+392))
	v1074 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+368))
	v1075 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+344))
	v1076 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+320))
	v1077 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+296))
	v1078 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+272))
	v1079 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+248))
	v1080 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+224))
	v1081 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+200))
	v1082 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+176))
	v1083 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+152))
	v1084 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+128))
	v1085 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+104))
	v1086 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+80))
	v1087 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+56))
	v1088 = *(*int64)(unsafe.Add(mBase, uint32(v1053)+32))
	v1120 = v1058 + (v1059 + (v1060 + (v1061 + (v1062 + (v1063 + (v1064 + (v1065 + (v1066 + (v1067 + (v1068 + (v1069 + (v1070 + (v1071 + (v1072 + (v1073 + (v1074 + (v1075 + (v1076 + (v1077 + (v1078 + (v1079 + (v1080 + (v1081 + (v1082 + (v1083 + (v1084 + (v1085 + (v1086 + (v1087 + (v1088 + v1054))))))))))))))))))))))))))))))
	goto L156
L155:
	;
	v1120 = v1054
	goto L156
L156:
	;
	goto L153
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v929)+4)) = v1124
	v1128 = F_palloc_mul(m, int32(120), v1121)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v929)+8)) = v1128
	v1132 = v926 + int32(12)
	v1134 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[6]))
	F_hash_seq_init(m, v1132, v1134)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	v1137 = F_hash_seq_search(m, v1132)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	if v1137 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v1140 = v920
	v1142 = v1137
	goto L164
L162:
	;
	goto L163
L163:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1200+int32(3584))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L4
	} else {
		goto L168
	}
L164:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v929)+4))
	v1161 = v1158 + v1140<<(uint(int32(4))%32)
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1142)))
	v1163 = *(*int64)(unsafe.Add(mBase, uint32(v1162)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1161)+8)) = v1163
	v1165 = *(*int64)(unsafe.Add(mBase, uint32(v1162)))
	*(*int64)(unsafe.Add(mBase, uint32(v1161))) = v1165
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v929)+8))
	v1168 = int32(120)
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+4))
	base.MemoryCopy(m, v1167+v1140*v1168, v1171, v1168)
	v1178 = F_hash_seq_search(m, v926+int32(12))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L4
	} else {
		goto L166
	}
L165:
	;
	goto L163
L166:
	;
	if v1178 != 0 {
		v1140 = v1140 + int32(1)
		v1142 = v1178
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1206+int32(_a_F_pg_lock_status_49))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1212+int32(_a_F_pg_lock_status_48))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1218+int32(_a_F_pg_lock_status_47))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1224+int32(_a_F_pg_lock_status_46))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1230+int32(_a_F_pg_lock_status_45))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1236+int32(_a_F_pg_lock_status_44))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1242+int32(_a_F_pg_lock_status_43))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1248+int32(_a_F_pg_lock_status_42))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1254+int32(_a_F_pg_lock_status_41))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1260+int32(_a_F_pg_lock_status_40))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1266+int32(_a_F_pg_lock_status_39))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1272+int32(_a_F_pg_lock_status_38))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1278+int32(_a_F_pg_lock_status_37))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1284+int32(_a_F_pg_lock_status_36))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1290+int32(_a_F_pg_lock_status_35))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1296+int32(_a_F_pg_lock_status_34))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	m.G0 = v926 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v243)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = v929
	*(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[0])) = v33
	goto L3
L185:
	;
	m.G0 = v22 + int32(272)
	return v1764
L186:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+16))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+4))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1330)))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1332)))
	if v1331 < v1333 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	goto L190
L188:
	;
	goto L189
L189:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+12))
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+8))
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1636)))
	if v1635 < v1637 {
		goto L259
	} else {
		goto L260
	}
L190:
	;
	v1356 = int32(0)
	base.MemoryFill(m, v22+int32(80), v1356, int32(128))
	v1359 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v1359
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = v1359
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+4))
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+4))
	v1367 = v1363 + v1364*int32(56)
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+16))
	if v1368 == v1356 {
		goto L194
	} else {
		goto L195
	}
L191:
	;
	goto L189
L192:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+4))
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1332)))
	if v1613 < v1614 {
		goto L190
	} else {
		goto L258
	}
L193:
	;
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1367)+14)))
	if base.Ui32(v1427) <= base.Ui32(int32(11)) {
		goto L227
	} else {
		goto L228
	}
L194:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1330)+4)) = v1364 + int32(1)
	if v1418 == int32(0) {
		goto L192
	} else {
		goto L225
	}
L195:
	;
	if v1368&int32(1) != 0 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+16)) = v1414 & v1368
	v1425 = v1413
	v1426 = int32(1)
	goto L193
L197:
	;
	v1413 = int32(0)
	v1414 = int32(-2)
	goto L196
L198:
	;
	goto L199
L199:
	;
	if v1368&int32(2) != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1413 = int32(1)
	v1414 = int32(-3)
	goto L196
L201:
	;
	goto L202
L202:
	;
	if v1368&int32(4) != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v1413 = int32(2)
	v1414 = int32(-5)
	goto L196
L204:
	;
	goto L205
L205:
	;
	if v1368&int32(8) != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v1413 = int32(3)
	v1414 = int32(-9)
	goto L196
L207:
	;
	goto L208
L208:
	;
	if v1368&int32(16) != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1413 = int32(4)
	v1414 = int32(-17)
	goto L196
L210:
	;
	goto L211
L211:
	;
	if v1368&int32(32) != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1413 = int32(5)
	v1414 = int32(-33)
	goto L196
L213:
	;
	goto L214
L214:
	;
	if v1368&int32(64) != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1413 = int32(6)
	v1414 = int32(-65)
	goto L196
L216:
	;
	goto L217
L217:
	;
	if v1368&int32(128) != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1413 = int32(7)
	v1414 = int32(-129)
	goto L196
L219:
	;
	goto L220
L220:
	;
	if v1368&int32(256) != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v1413 = int32(8)
	v1414 = int32(-257)
	goto L196
L222:
	;
	goto L223
L223:
	;
	if v1368&int32(512) == int32(0) {
		goto L194
	} else {
		goto L224
	}
L224:
	;
	v1413 = int32(9)
	v1414 = int32(-513)
	goto L196
L225:
	;
	v1425 = v1418
	v1426 = int32(0)
	goto L193
L226:
	;
	v1444 = F_cstring_to_text(m, v1443)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L4
	} else {
		goto L231
	}
L227:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1427<<(uint(int32(2))%32))+uint32(_c_F_pg_lock_status[7])))
	v1443 = v1432
	goto L226
L228:
	;
	goto L229
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v1427
	v1435 = v22 + int32(208)
	v1440 = F_pg_snprintf(m, v1435, int32(32), int32(_a_F_pg_lock_status_50), v22+int32(48))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L4
	} else {
		goto L230
	}
L230:
	;
	v1443 = v1435
	goto L226
L231:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = base.I64_extend_i32_u(v1444)
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	v1449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1367)+14)))
	switch v1449 {
	case 0, 1:
		goto L241
	case 2:
		goto L240
	case 3:
		goto L239
	case 4:
		goto L238
	case 5:
		goto L237
	case 6:
		goto L236
	case 7:
		goto L235
	default:
		goto L233
	case 11:
		goto L234
	}
L232:
	;
	v1552 = *(*int64)(unsafe.Add(mBase, uint32(v1367)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v1552
	v1555 = v22 + int32(240)
	v1560 = F_pg_snprintf(m, v1555, int32(32), int32(_a_F_pg_lock_status_51), v22+int32(16))
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L4
	} else {
		goto L244
	}
L233:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = base.I64_extend_i32_u(v1448)
	v1539 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1367)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+136)) = v1539
	v1541 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1367)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+144)) = v1541
	v1543 = int64(*(*int16)(unsafe.Add(mBase, uint32(v1367)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+66)) = int32(16843009)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+152)) = v1543
	v1547 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+70)) = uint8(v1547)
	goto L232
L234:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = base.I64_extend_i32_u(v1448)
	v1527 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1367)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+144)) = v1527
	v1529 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1367)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = v1529
	v1531 = int64(*(*int16)(unsafe.Add(mBase, uint32(v1367)+12)))
	v1532 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+71)) = uint8(v1532)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+152)) = v1531
	*(*int32)(unsafe.Add(mBase, uint32(v22)+66)) = int32(16843009)
	goto L232
L235:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = base.I64_extend_i32_u(v1448)
	v1515 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1367)+4)))
	v1516 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+73)) = uint8(v1516)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+71)) = uint8(v1516)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+144)) = v1515
	*(*int32)(unsafe.Add(mBase, uint32(v22)+65)) = int32(16843009)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+69)) = uint8(v1516)
	goto L232
L236:
	;
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v1448
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v1494
	v1498 = v22 + int32(240)
	v1499 = int32(32)
	v1503 = F_pg_snprintf(m, v1498, v1499, int32(_a_F_pg_lock_status_51), v22+v1499)
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L4
	} else {
		goto L242
	}
L237:
	;
	v1484 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+73)) = uint8(v1484)
	v1486 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+71)) = uint16(v1486)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+65)) = int32(16843009)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+69)) = uint8(v1484)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = base.I64_extend_i32_u(v1448)
	goto L232
L238:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = base.I64_extend_i32_u(v1448)
	v1474 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1367)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = v1474
	v1476 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1367)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v1476
	v1478 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v1367)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+69)) = int32(16843009)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = v1478
	v1482 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+73)) = uint8(v1482)
	goto L232
L239:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = base.I64_extend_i32_u(v1448)
	v1464 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1367)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = v1464
	v1466 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1367)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = int32(16843009)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v1466
	v1470 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+72)) = uint16(v1470)
	goto L232
L240:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+66)) = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = base.I64_extend_i32_u(v1448)
	goto L232
L241:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = base.I64_extend_i32_u(v1448)
	v1452 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1367)+4)))
	v1453 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+67)) = v1453
	*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = v1452
	*(*int32)(unsafe.Add(mBase, uint32(v22)+70)) = v1453
	goto L232
L242:
	;
	v1505 = F_cstring_to_text(m, v1498)
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L4
	} else {
		goto L243
	}
L243:
	;
	v1507 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+70)) = v1507
	*(*int32)(unsafe.Add(mBase, uint32(v22)+65)) = v1507
	*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = base.I64_extend_i32_u(v1505)
	goto L232
L244:
	;
	v1562 = F_cstring_to_text(m, v1555)
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L4
	} else {
		goto L245
	}
L245:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+160)) = base.I64_extend_i32_u(v1562)
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+40))
	if v1566 != 0 {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	v1571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1367)+15)))
	v1572 = int32(2)
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1571<<(uint(v1572)%32))+uint32(_c_F_pg_lock_status[8])))
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+8))
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1575+v1425<<(uint(v1572)%32))))
	goto L250
L247:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+168)) = base.I64_extend_i32_s(v1566)
	goto L246
L248:
	;
	goto L249
L249:
	;
	v1569 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+75)) = uint8(v1569)
	goto L246
L250:
	;
	v1580 = F_cstring_to_text(m, v1579)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L4
	} else {
		goto L251
	}
L251:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+184)) = base.I64_extend_i32_u(v1426)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+176)) = base.I64_extend_i32_u(v1580)
	v1586 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1367)+48)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+192)) = v1586
	if v1426 != 0 {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+28))
	v1601 = F_heap_form_tuple(m, v1596, v22+int32(80), v22-int32(-64))
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L4
	} else {
		goto L256
	}
L253:
	;
	v1593 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+79)) = uint8(v1593)
	goto L252
L254:
	;
	v1588 = *(*int64)(unsafe.Add(mBase, uint32(v1367)+32))
	if v1588 == int64(0) {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+200)) = v1588
	goto L252
L256:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v1601)+16))
	v1604 = F_HeapTupleHeaderGetDatum(m, v1603)
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	v1606 = *(*int64)(unsafe.Add(mBase, uint32(v1329)))
	*(*int64)(unsafe.Add(mBase, uint32(v1329))) = v1606 + int64(1)
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+20)) = int32(1)
	v1764 = v1604
	goto L185
L258:
	;
	goto L191
L259:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1636)+4))
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1636)+8))
	v1643 = int32(0)
	base.MemoryFill(m, v22+int32(96), v1643, int32(112))
	v1646 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+216)) = v1646
	*(*int64)(unsafe.Add(mBase, uint32(v22)+208)) = v1646
	*(*int32)(unsafe.Add(mBase, uint32(v1330)+12)) = v1635 + int32(1)
	v1656 = v1639 + v1635<<(uint(int32(4))%32)
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1656)+12))
	if v1657 == v1643 {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	goto L261
L261:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L4
	} else {
		goto L282
	}
L262:
	;
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v1656)+8))
	v1663 = base.B2i32(v1660 != int32(-1))
	goto L264
L263:
	;
	v1663 = int32(2)
	goto L264
L264:
	;
	v1666 = v1635*int32(120) + v1640
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1663<<(uint(int32(2))%32))+uint32(_c_F_pg_lock_status[9])))
	v1670 = F_cstring_to_text(m, v1669)
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = base.I64_extend_i32_u(v1670)
	v1674 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1656))))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = v1674
	v1676 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1656)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = v1676
	if v1663 == int32(2) {
		goto L269
	} else {
		goto L270
	}
L266:
	;
	v1690 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+217)) = uint8(v1690)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+213)) = int32(16843009)
	v1694 = *(*int64)(unsafe.Add(mBase, uint32(v1666)))
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v1694
	v1697 = v22 + int32(240)
	v1700 = F_pg_snprintf(m, v1697, int32(32), int32(_a_F_pg_lock_status_51), v22)
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L4
	} else {
		goto L273
	}
L267:
	;
	v1688 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+211)) = uint8(v1688)
	goto L266
L268:
	;
	v1686 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1656)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v1686
	goto L266
L269:
	;
	v1680 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v1656)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = v1680
	goto L268
L270:
	;
	goto L271
L271:
	;
	v1682 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+212)) = uint8(v1682)
	if v1663 != v1682 {
		goto L267
	} else {
		goto L272
	}
L272:
	;
	goto L268
L273:
	;
	v1702 = F_cstring_to_text(m, v1697)
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L4
	} else {
		goto L274
	}
L274:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+160)) = base.I64_extend_i32_u(v1702)
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1666)+112))
	if v1706 != 0 {
		goto L276
	} else {
		goto L277
	}
L275:
	;
	v1712 = F_cstring_to_text(m, int32(_a_F_pg_lock_status_52))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L4
	} else {
		goto L279
	}
L276:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+168)) = base.I64_extend_i32_s(v1706)
	goto L275
L277:
	;
	goto L278
L278:
	;
	v1709 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+219)) = uint8(v1709)
	goto L275
L279:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+192)) = int64(0)
	v1716 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+223)) = uint8(v1716)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+176)) = base.I64_extend_i32_u(v1712)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+184)) = int64(1)
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+28))
	v1727 = F_heap_form_tuple(m, v1722, v22+int32(80), v22+int32(208))
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L4
	} else {
		goto L280
	}
L280:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1727)+16))
	v1730 = F_HeapTupleHeaderGetDatum(m, v1729)
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L4
	} else {
		goto L281
	}
L281:
	;
	v1732 = *(*int64)(unsafe.Add(mBase, uint32(v1329)))
	*(*int64)(unsafe.Add(mBase, uint32(v1329))) = v1732 + int64(1)
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1736)+20)) = int32(1)
	v1764 = v1730
	goto L185
L282:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1741)+20)) = int32(2)
	v1744 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1744)
	v1764 = int64(0)
	goto L185
}
func F_pg_log_backend_memory_contexts(m *base.Module, l0 int32) int64 {
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
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v94 int64
	_ = v94
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_BackendPidGetProc(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v94
L2:
	;
	v94 = int64(0)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
	F_errmsg(m, v80, v8)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L26
	}
L4:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_pg_log_backend_memory_contexts[0]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v67 = base.I32_div_s(v59-v64, int32(768))
	v68 = F_SendProcSignal(m, v10, int32(5), v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L22
	}
L5:
	;
	return int64(0)
L6:
	;
	if v11 != 0 {
		v59 = v11
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v15 = int32(0)
	if v10 == v15 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v50 != 0 {
		v59 = v50
		goto L4
	} else {
		goto L19
	}
L9:
	;
	v50 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_pg_log_backend_memory_contexts[1]))
	v27 = v15
	goto L14
L12:
	;
	v50 = v44
	goto L8
L13:
	;
	v44 = v34 + int32(768)
	goto L12
L14:
	;
	v30 = v27 * int32(768)
	v31 = v23 + v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v32 == v10 {
		v44 = v31
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v50 = int32(0)
	goto L8
L16:
	;
	v34 = v23 + v30
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+780))
	if v35 == v10 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v38 = v27 + int32(2)
	if v38 != int32(38) {
		v27 = v38
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v53 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	if v53 == int32(0) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v80 = int32(_a_F_pg_log_backend_memory_contexts_0)
	v81 = int32(296)
	goto L3
L22:
	;
	if int32(0) <= v68 {
		v94 = int64(1)
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v74 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	if v74 == int32(0) {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v80 = int32(_a_F_pg_log_backend_memory_contexts_1)
	v81 = int32(305)
	goto L3
L26:
	;
	F_errfinish(m, int32(_a_F_pg_log_backend_memory_contexts_2), v81, int32(_a_F_pg_log_backend_memory_contexts_3))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	goto L2
}
func F_pg_lsn_mi(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(288)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui64(v9) < base.Ui64(v10) {
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v10 - v9
		v18 = F_pg_snprintf(m, v7+int32(32), int32(256), int32(_a_F_pg_lsn_mi_0), v7)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v39 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), base.I64_extend_i32_u(v7+int32(32)), int64(0), int64(-1))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int64(0)
			} else {
				m.G0 = v7 + int32(288)
				return v39
			}
		}
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v9 - v10
		v30 = F_pg_snprintf(m, v7+int32(32), int32(256), int32(_a_F_pg_lsn_mi_1), v7+int32(16))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int64(0)
		} else {
			v39 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), base.I64_extend_i32_u(v7+int32(32)), int64(0), int64(-1))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int64(0)
			} else {
				m.G0 = v7 + int32(288)
				return v39
			}
		}
	}
}
func F_pg_ndistinct_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
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
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
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
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v251 int32
	_ = v251
	var v273 int32
	_ = v273
	var v274 float64
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v635 int32
	_ = v635
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v672 int32
	_ = v672
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v808 int64
	_ = v808
	v2 = int32(0)
	v18 = int64(0)
	v19 = m.G0
	v21 = v19 - int32(304)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+264)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v21)+260)) = v23
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+280)) = v18
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+276)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+272)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+256)) = int32(1671)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+252)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+248)) = int32(1672)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+244)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = int32(1673)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+236)) = int32(1674)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+232)) = int32(1675)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+228)) = int32(1676)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+224)) = int32(1677)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+220)) = v21 + int32(260)
	v55 = F_strlen(m, v23)
	mBase = m.M
	v58 = F_makeJsonLexContextCstringLen(m, v2, v23, v55, int32(6), int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v64 = F_pg_parse_json(m, v58, v21+int32(220))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_freeJsonLexContext(m, v58)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v64 != 0 {
		v754 = v2
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v21)+280))
	F_list_free(m, v770)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L128
	}
L6:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)+268))
	if v68 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v71 = v69
	goto L9
L8:
	;
	v71 = int32(0)
	goto L9
L9:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v21)+264))
	switch v72 {
	case 0:
		goto L12
	default:
		goto L11
	case 6:
		goto L13
	}
L10:
	;
	v141 = F_palloc(m, v71<<(uint(int32(4))%32)+int32(16))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L30
	}
L11:
	;
	v110 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v21)+272))
	v112 = F_errsave_start(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L24
	}
L12:
	;
	v86 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v21)+272))
	v88 = F_errsave_start(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v71 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errmsg_internal(m, int32(_a_F_pg_ndistinct_in_0), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_pg_ndistinct_in_1), int32(607), int32(_a_F_pg_ndistinct_in_2))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	if v88 == int32(0) {
		v754 = v86
		goto L5
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+208)) = v23
	F_errmsg(m, int32(_a_F_pg_ndistinct_in_3), v21+int32(208))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v103 = F_errdetail(m, int32(_a_F_pg_ndistinct_in_4), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errsave_finish(m, v87, int32(_a_F_pg_ndistinct_in_1), int32(615), int32(_a_F_pg_ndistinct_in_2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v754 = v86
	goto L5
L24:
	;
	if v112 == int32(0) {
		v754 = v110
		goto L5
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v23
	F_errmsg(m, int32(_a_F_pg_ndistinct_in_3), v21+int32(32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v21)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v125
	v130 = F_errdetail(m, int32(_a_F_pg_ndistinct_in_5), v21+int32(16))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errsave_finish(m, v111, int32(_a_F_pg_ndistinct_in_1), int32(623), int32(_a_F_pg_ndistinct_in_2))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v754 = v110
	goto L5
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+8)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v141))) = int64(7035076516)
	if int32(0) < v71 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v655 = v21 + int32(288)
	F_initStringInfo(m, v655)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L110
	}
L32:
	;
	F_pfree(m, v141)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L109
	}
L33:
	;
	v149 = v141 + int32(16)
	v154 = v2
	v161 = v2
	v165 = v2
	goto L36
L34:
	;
	goto L35
L35:
	;
	v614 = F_statext_ndistinct_serialize(m, v141)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L108
	}
L36:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v21)+268))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169+v154<<(uint(int32(2))%32))))
	if v154 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v289 = v149 + v283<<(uint(int32(4))%32)
	v296 = int32(0)
	goto L59
L38:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	v181 = int32(0)
	goto L41
L39:
	;
	goto L40
L40:
	;
	v273 = v149 + v154<<(uint(int32(4))%32)
	v274 = *(*float64)(unsafe.Add(mBase, uint32(v173)))
	*(*float64)(unsafe.Add(mBase, uint32(v273))) = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v273)+8)) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v273)+12)) = v278
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	v281 = base.B2i32(v161 < v280)
	if v161 < v280 {
		goto L51
	} else {
		goto L52
	}
L41:
	;
	v196 = v149 + v181<<(uint(int32(4))%32)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	if v174 != v197 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L40
L43:
	;
	v251 = v181 + int32(1)
	if v251 != v154 {
		v181 = v251
		goto L41
	} else {
		goto L50
	}
L44:
	;
	if v174 <= int32(0) {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v206 = int32(0)
	goto L46
L46:
	;
	v223 = v206 << (uint(int32(1)) % 32)
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202+v223))))
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201+v223))))
	if v225 != v227 {
		goto L43
	} else {
		goto L48
	}
L47:
	;
	goto L31
L48:
	;
	v230 = v206 + int32(1)
	if v174 != v230 {
		v206 = v230
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	goto L42
L51:
	;
	v282 = v280
	goto L53
L52:
	;
	v282 = v161
	goto L53
L53:
	;
	if v161 < v280 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v283 = v154
	goto L56
L55:
	;
	v283 = v165
	goto L56
L56:
	;
	v285 = v154 + int32(1)
	if v285 != v71 {
		v154 = v285
		v161 = v282
		v165 = v283
		goto L36
	} else {
		goto L57
	}
L57:
	;
	goto L37
L58:
	;
	v444 = v21 + int32(288)
	F_initStringInfo(m, v444)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L80
	}
L59:
	;
	if v296 == v283 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v396 = F_statext_ndistinct_serialize(m, v141)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L75
	}
L61:
	;
	v393 = v296 + int32(1)
	if v393 != v71 {
		v296 = v393
		goto L59
	} else {
		goto L74
	}
L62:
	;
	v312 = v149 + v296<<(uint(int32(4))%32)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+8))
	if v313 <= int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
	if v316 <= int32(0) {
		goto L58
	} else {
		goto L64
	}
L64:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v289)+12))
	v326 = int32(0)
	goto L65
L65:
	;
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v319+v326<<(uint(int32(1))%32)))))
	v347 = int32(0)
	goto L67
L66:
	;
	goto L61
L67:
	;
	v366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v320+v347<<(uint(int32(1))%32)))))
	if v366 != v343 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v372 = v326 + int32(1)
	if v372 != v313 {
		v326 = v372
		goto L65
	} else {
		goto L73
	}
L69:
	;
	v369 = v347 + int32(1)
	if v316 != v369 {
		v347 = v369
		goto L67
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	goto L68
L72:
	;
	goto L58
L73:
	;
	goto L66
L74:
	;
	goto L60
L75:
	;
	v402 = int32(0)
	goto L76
L76:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v141+v402<<(uint(int32(4))%32))+28))
	F_pfree(m, v419)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v618 = v396
	goto L32
L78:
	;
	v423 = v402 + int32(1)
	if v423 != v71 {
		v402 = v423
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	v448 = int32(*(*int16)(unsafe.Add(mBase, uint32(v447))))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v448
	F_appendStringInfo(m, v444, int32(_a_F_pg_ndistinct_in_6), v21+int32(128))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v312)+8))
	if int32(2) <= v455 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v461 = int32(1)
	goto L85
L83:
	;
	goto L84
L84:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v21)+288))
	v514 = v21 + int32(288)
	F_initStringInfo(m, v514)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L89
	}
L85:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	v481 = int32(*(*int16)(unsafe.Add(mBase, uint32(v477+v461<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v481
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_pg_ndistinct_in_7), v21+int32(112))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	goto L84
L87:
	;
	v491 = v461 + int32(1)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v312)+8))
	if v491 < v492 {
		v461 = v491
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v289)+12))
	v518 = int32(*(*int16)(unsafe.Add(mBase, uint32(v517))))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v518
	F_appendStringInfo(m, v514, int32(_a_F_pg_ndistinct_in_6), v21+int32(96))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
	if int32(2) <= v525 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v531 = int32(1)
	goto L94
L92:
	;
	goto L93
L93:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v21)+288))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v21)+272))
	v585 = F_errsave_start(m, v584)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L98
	}
L94:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v289)+12))
	v551 = int32(*(*int16)(unsafe.Add(mBase, uint32(v547+v531<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v551
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_pg_ndistinct_in_7), v21+int32(80))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L96
	}
L95:
	;
	goto L93
L96:
	;
	v561 = v531 + int32(1)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
	if v561 < v562 {
		v531 = v561
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	if v585 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	F_pfree(m, v512)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L106
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v23
	F_errmsg(m, int32(_a_F_pg_ndistinct_in_3), v21-int32(-64))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = int32(_a_F_pg_ndistinct_in_8)
	v603 = F_errdetail(m, int32(_a_F_pg_ndistinct_in_9), v21+int32(48))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errsave_finish(m, v584, int32(_a_F_pg_ndistinct_in_1), int32(703), int32(_a_F_pg_ndistinct_in_2))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	goto L101
L106:
	;
	F_pfree(m, v582)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v754 = int32(0)
	goto L5
L108:
	;
	v618 = v614
	goto L32
L109:
	;
	v754 = v618
	goto L5
L110:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v659 = int32(*(*int16)(unsafe.Add(mBase, uint32(v658))))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v659
	F_appendStringInfo(m, v655, int32(_a_F_pg_ndistinct_in_6), v21+int32(192))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	if int32(2) <= v666 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v672 = int32(1)
	goto L115
L113:
	;
	goto L114
L114:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v21)+288))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v21)+272))
	v726 = F_errsave_start(m, v725)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L119
	}
L115:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v692 = int32(*(*int16)(unsafe.Add(mBase, uint32(v688+v672<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v692
	F_appendStringInfo(m, v21+int32(288), int32(_a_F_pg_ndistinct_in_7), v21+int32(176))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L117
	}
L116:
	;
	goto L114
L117:
	;
	v702 = v672 + int32(1)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	if v702 < v703 {
		v672 = v702
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	if v726 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	F_pfree(m, v723)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L127
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+160)) = v23
	F_errmsg(m, int32(_a_F_pg_ndistinct_in_3), v21+int32(160))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+148)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v21)+144)) = int32(_a_F_pg_ndistinct_in_8)
	v743 = F_errdetail(m, int32(_a_F_pg_ndistinct_in_10), v21+int32(144))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errsave_finish(m, v725, int32(_a_F_pg_ndistinct_in_1), int32(652), int32(_a_F_pg_ndistinct_in_2))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	goto L122
L127:
	;
	v754 = int32(0)
	goto L5
L128:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v21)+268))
	F_list_free_deep(m, v773)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	if v754 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v808 = base.I64_extend_i32_u(v754)
	goto L132
L131:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v21)+272))
	if v777 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	m.G0 = v21 + int32(304)
	return v808
L133:
	;
	v804 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v804)
	v808 = int64(0)
	goto L132
L134:
	;
	v784 = F_errsave_start(m, v777)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L138
	}
L135:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v777)))
	if v780 != int32(453) {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777)+4)))
	if v783 != 0 {
		goto L133
	} else {
		goto L137
	}
L137:
	;
	goto L134
L138:
	;
	if v784 == int32(0) {
		goto L133
	} else {
		goto L139
	}
L139:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v23
	F_errmsg(m, int32(_a_F_pg_ndistinct_in_3), v21)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v797 = F_errdetail(m, int32(_a_F_pg_ndistinct_in_11), int32(0))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errsave_finish(m, v777, int32(_a_F_pg_ndistinct_in_1), int32(780), int32(_a_F_pg_ndistinct_in_12))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	goto L133
}
func F_pg_node_tree_recv(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_pg_node_tree_recv_0), int32(335), int32(_a_F_pg_node_tree_recv_1), int32(_a_F_pg_node_tree_recv_2), int32(_a_F_pg_node_tree_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_pg_parse_json(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	if l0 == int32(_a_F_pg_parse_json_0) {
		return int32(16)
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v8 != 0 {
			return int32(2)
		} else {
			v11 = F_json_lex(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v11 != 0 {
					v35 = v11
					return v35
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					switch v15 - int32(3) {
					case 0:
						v18 = F_parse_object(m, l0, l1)
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return int32(0)
						} else {
							v24 = v18
							if v24 != 0 {
								v35 = v24
								return v35
							} else {
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v25 == int32(12) {
									v28 = F_json_lex(m, l0)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return int32(0)
									} else {
										return v28
									}
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v33 != 0 {
										v34 = int32(9)
									} else {
										v34 = int32(11)
									}
									v35 = v34
									return v35
								}
							}
						}
					default:
						v22 = F_parse_scalar(m, l0, l1)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v24 = v22
							if v24 != 0 {
								v35 = v24
								return v35
							} else {
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v25 == int32(12) {
									v28 = F_json_lex(m, l0)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return int32(0)
									} else {
										return v28
									}
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v33 != 0 {
										v34 = int32(9)
									} else {
										v34 = int32(11)
									}
									v35 = v34
									return v35
								}
							}
						}
					case 2:
						v20 = F_parse_array(m, l0, l1)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							v24 = v20
							if v24 != 0 {
								v35 = v24
								return v35
							} else {
								v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v25 == int32(12) {
									v28 = F_json_lex(m, l0)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return int32(0)
									} else {
										return v28
									}
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v33 != 0 {
										v34 = int32(9)
									} else {
										v34 = int32(11)
									}
									v35 = v34
									return v35
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_plan_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 == int32(6) {
		v47 = int32(0)
		return v47
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_query[0])))
		if v11 == int32(1) {
			v14 = int32(_a_F_pg_plan_query_0)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_query[1])) = int64(4)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_query[2])) = int64(3)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_query[3])) = int64(2)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_query[4])) = int64(1)
			v24 = F___syscall_ret(m, int32(0))
			mBase = m.M
			F_gettimeofday(m, int32(_a_F_pg_plan_query_1))
			mBase = m.M
		} else {
		}
		v27 = F_planner(m, l0, l1, l2, l3, l4)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_query[0])))
			if v32 == int32(1) {
				F_ShowUsage(m, int32(_a_F_pg_plan_query_2))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_query[5])))
					if v39 != int32(1) {
						v47 = v27
						return v47
					} else {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_query[6])))
						F_elog_node_display(m, int32(_a_F_pg_plan_query_3), v27, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = v27
							return v47
						}
					}
				}
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_query[5])))
				if v39 != int32(1) {
					v47 = v27
					return v47
				} else {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_query[6])))
					F_elog_node_display(m, int32(_a_F_pg_plan_query_3), v27, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = v27
						return v47
					}
				}
			}
		}
	}
}
func F_pg_prepared_xact(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
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
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v259 int32
	_ = v259
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int64
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int64
	_ = v362
	var v364 int64
	_ = v364
	var v366 int64
	_ = v366
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int64
	_ = v376
	var v377 int32
	_ = v377
	var v378 int64
	_ = v378
	var v382 int32
	_ = v382
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v418 int64
	_ = v418
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v19 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+16))
	goto L52
L4:
	;
	return int64(0)
L5:
	;
	v26 = int32(_a_F_pg_prepared_xact_0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[0]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[0])) = v29
	v32 = F_CreateTemplateTupleDesc(m, int32(5))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v32, int32(1), int32(_a_F_pg_prepared_xact_1), int32(28), int32(-1), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v32, int32(2), int32(_a_F_pg_prepared_xact_2), int32(25), int32(-1), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v32, int32(3), int32(_a_F_pg_prepared_xact_3), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v32, int32(4), int32(_a_F_pg_prepared_xact_4), int32(26), int32(-1), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v32, int32(5), int32(_a_F_pg_prepared_xact_5), int32(26), int32(-1), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v69 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v69 < v78 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v156 = F_BlessTupleDesc(m, v32)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L31
	}
L13:
	;
	v82 = v32 + int32(28)
	v89 = v69
	v90 = v78
	v92 = v69
	goto L17
L14:
	;
	v146 = v69
	v153 = v78
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v146
	goto L12
L16:
	;
	v146 = v140
	v153 = v119
	goto L15
L17:
	;
	v98 = v82 + v78<<(uint(int32(3))%32) + v89*int32(100)
	v101 = v82 + v89<<(uint(int32(3))%32)
	if v78 != v90 {
		v119 = v90
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v140 = v78
	goto L16
L19:
	;
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101)+2)))
	if v120 <= int32(0) {
		v140 = v89
		goto L16
	} else {
		goto L27
	}
L20:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+7)))
	if v103 != int32(118) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v119 = v89
	goto L19
L22:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
	if v106 != int32(1) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+6)))
	if v109&int32(6) != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101)+2)))
	if v112 <= int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+90)))
	if v115 != int32(118) {
		v119 = v78
		goto L19
	} else {
		goto L26
	}
L26:
	;
	goto L21
L27:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+90)))
	if v123 == int32(118) {
		v140 = v89
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+5)))
	v132 = (v92 + v126 - int32(1)) & (int32(0) - v126)
	if int32(_a_F_pg_prepared_xact_6) < v132 {
		v140 = v89
		goto L16
	} else {
		goto L29
	}
L29:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v101))) = uint16(v132)
	v138 = v89 + int32(1)
	if v138 != v78 {
		v89 = v138
		v90 = v119
		v92 = v132 + v120
		goto L17
	} else {
		goto L30
	}
L30:
	;
	goto L18
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v156
	v160 = F_palloc(m, int32(12))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v160
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[1]))
	v168 = F_LWLockAcquire(m, v164+int32(2304), int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[2]))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	if v172 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = v172
	*(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[0])) = v27
	goto L3
L35:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[1]))
	F_LWLockRelease(m, v176+int32(2304))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v184 = F_palloc_mul(m, int32(256), v172)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = int32(0)
	goto L34
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v184
	if v172 <= int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[1]))
	F_LWLockRelease(m, v276+int32(2304))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L49
	}
L41:
	;
	v189 = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[2]))
	v193 = v191 + int32(8)
	if v172 != int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v201 = v189
	v208 = v2
	goto L45
L43:
	;
	v241 = v189
	goto L44
L44:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v193+v241<<(uint(int32(2))%32))))
	base.MemoryCopy(m, v184+v241<<(uint(int32(8))%32), v259, int32(256))
	goto L40
L45:
	;
	v213 = int32(8)
	v216 = int32(2)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v193+v201<<(uint(v216)%32))))
	v220 = int32(256)
	base.MemoryCopy(m, v184+v201<<(uint(v213)%32), v219, v220)
	v223 = v201 | int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v193+v223<<(uint(v216)%32))))
	base.MemoryCopy(m, v184+v223<<(uint(v213)%32), v230, v220)
	v234 = v201 + v216
	v236 = v208 + v216
	if v236 != v172&int32(2147483646) {
		v201 = v234
		v208 = v236
		goto L45
	} else {
		goto L47
	}
L46:
	;
	if v172&int32(1) == int32(0) {
		goto L40
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	v241 = v234
	goto L44
L49:
	;
	goto L34
L50:
	;
	m.G0 = v16 - int32(-64)
	return v418
L51:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L64
	}
L52:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+16))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	if v315 == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v314)+8))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	if v319 <= v318 {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[3]))
	v324 = v318
	goto L55
L55:
	;
	v336 = int32(1)
	v337 = v324 + v336
	*(*int32)(unsafe.Add(mBase, uint32(v314)+8)) = v337
	v341 = v315 + v324<<(uint(int32(8))%32)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v344 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+12)) = uint8(v344)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v344
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+48)))
	if v348 == v336 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L51
L57:
	;
	v353 = v343 + v342*int32(768)
	v354 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v353)+48)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v354
	v358 = F_cstring_to_text(m, v341+int32(51))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v337 < v319 {
		v324 = v337
		goto L55
	} else {
		goto L63
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = base.I64_extend_i32_u(v358)
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v341)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v362
	v364 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v341)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v364
	v366 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v353)+20)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v313)+28))
	v373 = F_heap_form_tuple(m, v368, v14+int32(-48), v14+int32(-56))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	v376 = F_HeapTupleHeaderGetDatum(m, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v313)))
	*(*int64)(unsafe.Add(mBase, uint32(v313))) = v378 + int64(1)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v382)+20)) = int32(1)
	v418 = v376
	goto L50
L63:
	;
	goto L56
L64:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v401)+20)) = int32(2)
	v404 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v404)
	v418 = int64(0)
	goto L50
}
func F_pg_promote(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v178 int64
	_ = v178
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_promote[0])))
	if v15 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v255 = F_unlink(m, int32(_a_F_pg_promote_0))
	mBase = m.M
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L13
	} else {
		goto L66
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L13
	} else {
		goto L62
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L13
	} else {
		goto L58
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L13
	} else {
		goto L54
	}
L5:
	;
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_pg_promote[1]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+308))
	v23 = base.B2i32(v21 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_promote[0])) = uint8(v23)
	v25 = v23
	goto L8
L7:
	;
	v25 = int32(0)
	goto L8
L8:
	;
	goto L5
L9:
	;
	v26 = base.I32_wrap_i64(v12)
	if v26 <= int32(0) {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L13
	} else {
		goto L49
	}
L12:
	;
	v31 = F_AllocateFile(m, int32(_a_F_pg_promote_0), int32(_a_F_pg_promote_1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int64(0)
L14:
	;
	if v31 == int32(0) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v37 = F_FreeFile(m, v31)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v37 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_pg_promote[2]))
	v42 = F_pgmem_kill(m, v40, int32(10))
	mBase = m.M
	if v42 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v43 = int64(1)
	if v11 == int64(0) {
		v178 = v43
		goto L19
	} else {
		goto L20
	}
L19:
	;
	m.G0 = v9 + int32(48)
	return v178
L20:
	;
	v49 = m.G0
	v50 = int32(16)
	v51 = v49 - v50
	m.G0 = v51
	F_gettimeofday(m, v51)
	mBase = m.M
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v51)+8)))
	m.G0 = v51 + v50
	goto L21
L21:
	;
	goto L23
L22:
	;
	v155 = int64(0)
	v158 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L45
	}
L23:
	;
	v78 = m.G0
	v79 = int32(16)
	v80 = v78 - v79
	m.G0 = v80
	F_gettimeofday(m, v80)
	mBase = m.M
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v80)))
	v84 = int64(*(*int32)(unsafe.Add(mBase, uint32(v80)+8)))
	m.G0 = v80 + v79
	goto L25
L24:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L13
	} else {
		goto L39
	}
L25:
	;
	if v55+v54*int64(1000000)-int64(946684800000000)+v12&int64(2147483647)*int64(1000000) <= v84+v83*int64(1000000)-int64(946684800000000) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_pg_promote[3]))
	v96 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v96
	v101 = base.AtomicRmwOr32(m, v96, int32(_a_F_pg_promote_2), v96)
	goto L27
L27:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_promote[0])))
	if v104 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v114 == int32(0) {
		v178 = v43
		goto L19
	} else {
		goto L32
	}
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_pg_promote[1]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+308))
	v112 = base.B2i32(v110 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_promote[0])) = uint8(v112)
	v114 = v112
	goto L31
L30:
	;
	v114 = int32(0)
	goto L31
L31:
	;
	goto L28
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_pg_promote[4]))
	if v118 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L13
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_pg_promote[3]))
	v126 = F_WaitLatch(m, v122, int32(25), int32(100), int32(134217771))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L13
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	if v126&int32(16) == int32(0) {
		goto L23
	} else {
		goto L38
	}
L38:
	;
	goto L24
L39:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(_a_F_pg_promote_3), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L13
	} else {
		goto L41
	}
L41:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	F_errcontext_msg(m, int32(_a_F_pg_promote_4), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_pg_promote_5), int32(760), int32(_a_F_pg_promote_6))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	if v158 == int32(0) {
		v178 = v155
		goto L19
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v26
	F_errmsg_plural(m, int32(_a_F_pg_promote_7), int32(_a_F_pg_promote_8), v26, v9+int32(16))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_pg_promote_5), int32(767), int32(_a_F_pg_promote_6))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	v178 = v155
	goto L19
L49:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L13
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_pg_promote_9), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	F_errhint(m, int32(_a_F_pg_promote_10), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L13
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_pg_promote_5), int32(700), int32(_a_F_pg_promote_6))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L13
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(_a_F_pg_promote_11), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_pg_promote_5), int32(705), int32(_a_F_pg_promote_6))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_pg_promote_0)
	F_errmsg(m, int32(_a_F_pg_promote_12), v9)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_pg_promote_5), int32(713), int32(_a_F_pg_promote_6))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(_a_F_pg_promote_0)
	F_errmsg(m, int32(_a_F_pg_promote_13), v9+int32(32))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_pg_promote_5), int32(719), int32(_a_F_pg_promote_6))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(517))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L13
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_pg_promote_14), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L13
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_pg_promote_5), int32(727), int32(_a_F_pg_promote_6))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_read_binary_file_all_missing(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = F_convert_and_check_filename(m, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v12 = int64(0)
			v16 = F_read_binary_file(m, v10, v12, int64(-1), base.B2i32(v9 != v12))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				if v16 == int32(0) {
					v20 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v16)
				}
			}
		}
	}
}
func F_pg_read_binary_file_off_len_missing(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		if int64(0) <= v11 {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
			v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
			v16 = F_convert_and_check_filename(m, v7)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				v20 = F_read_binary_file(m, v16, v15, v11, base.B2i32(v14 != int64(0)))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					if v20 == int32(0) {
						v24 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v20)
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_read_binary_file_off_len_missing_0), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_read_binary_file_off_len_missing_1), int32(269), int32(_a_F_pg_read_binary_file_off_len_missing_2))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
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
func F_pg_reg_getcolor(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
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
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
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
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v3 < v8 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if int32(2) <= v52 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v51 = v43
	goto L1
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v14 = v8
	v16 = v3
	goto L6
L4:
	;
	goto L5
L5:
	;
	v51 = int32(0)
	goto L1
L6:
	;
	v21 = base.I32_div_s(v14-v16, int32(2))
	v22 = v21 + v16
	v25 = v11 + v22*int32(12)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if base.Ui32(l1) < base.Ui32(v26) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	if v33 < v32 {
		v14 = v32
		v16 = v33
		goto L6
	} else {
		goto L13
	}
L9:
	;
	v32 = v22
	v33 = v16
	goto L8
L10:
	;
	goto L11
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if base.Ui32(l1) <= base.Ui32(v28) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v32 = v14
	v33 = v22 + int32(1)
	goto L8
L13:
	;
	goto L7
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v55 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v321 = int32(*(*int16)(unsafe.Add(mBase, uint32(v317+v51<<(uint(int32(1))%32)))))
	return v321
L17:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v82 == int32(0) {
		v106 = v81
		goto L30
	} else {
		goto L31
	}
L18:
	;
	v81 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)))
	if v62 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if base.Ui32(l1-int32(32)) < base.Ui32(int32(95)) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+40))
	v73 = m.T0[v72].(func(*base.Module, int32, int32) int32)(m, l1, v61)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v70 = v55
	goto L26
L25:
	;
	v70 = int32(0)
	goto L26
L26:
	;
	v81 = v70
	goto L17
L27:
	;
	return int32(0)
L28:
	;
	if v73 == int32(0) {
		v81 = int32(0)
		goto L17
	} else {
		goto L29
	}
L29:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v81 = v79
	goto L17
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v107 == int32(0) {
		v131 = v106
		goto L40
	} else {
		goto L41
	}
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+2)))
	if v87 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v106 = v102 | v81
	goto L30
L33:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l1) {
		v106 = v81
		goto L30
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
	v97 = m.T0[v96].(func(*base.Module, int32, int32) int32)(m, l1, v86)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L27
	} else {
		goto L38
	}
L36:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_pg_reg_getcolor[1]))))
	if v92&int32(3) != 0 {
		v102 = v82
		goto L32
	} else {
		goto L37
	}
L37:
	;
	v106 = v81
	goto L30
L38:
	;
	if v97 == int32(0) {
		v106 = v81
		goto L30
	} else {
		goto L39
	}
L39:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v102 = v101
	goto L32
L40:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v132 == int32(0) {
		v159 = v131
		goto L50
	} else {
		goto L51
	}
L41:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+2)))
	if v112 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v131 = v127 | v106
	goto L40
L43:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l1) {
		v131 = v106
		goto L40
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
	v122 = m.T0[v121].(func(*base.Module, int32, int32) int32)(m, l1, v111)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L27
	} else {
		goto L48
	}
L46:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_pg_reg_getcolor[1]))))
	if v117&int32(2) != 0 {
		v127 = v107
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v131 = v106
	goto L40
L48:
	;
	if v122 == int32(0) {
		v131 = v106
		goto L40
	} else {
		goto L49
	}
L49:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v127 = v126
	goto L42
L50:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v160 == int32(0) {
		v183 = v159
		goto L61
	} else {
		goto L62
	}
L51:
	;
	if l1 == int32(95) {
		v154 = v132
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v159 = v154 | v131
	goto L50
L53:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+2)))
	if v139 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l1) {
		v159 = v131
		goto L50
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+24))
	v149 = m.T0[v148].(func(*base.Module, int32, int32) int32)(m, l1, v138)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L27
	} else {
		goto L59
	}
L57:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_pg_reg_getcolor[1]))))
	if v144&int32(3) != 0 {
		v154 = v132
		goto L52
	} else {
		goto L58
	}
L58:
	;
	v159 = v131
	goto L50
L59:
	;
	if v149 == int32(0) {
		v159 = v131
		goto L50
	} else {
		goto L60
	}
L60:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v154 = v153
	goto L52
L61:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v184 == int32(0) {
		v208 = v183
		goto L70
	} else {
		goto L71
	}
L62:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+2)))
	if v165 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v183 = v179 | v159
	goto L61
L64:
	;
	if base.Ui32(l1-int32(48)) < base.Ui32(int32(10)) {
		v179 = v160
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+16))
	v174 = m.T0[v173].(func(*base.Module, int32, int32) int32)(m, l1, v164)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L27
	} else {
		goto L68
	}
L67:
	;
	v183 = v159
	goto L61
L68:
	;
	if v174 == int32(0) {
		v183 = v159
		goto L61
	} else {
		goto L69
	}
L69:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v179 = v178
	goto L63
L70:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v209 == int32(0) {
		v233 = v208
		goto L80
	} else {
		goto L81
	}
L71:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+2)))
	if v189 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v208 = v204 | v183
	goto L70
L73:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l1) {
		v208 = v183
		goto L70
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v188)+8))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+44))
	v199 = m.T0[v198].(func(*base.Module, int32, int32) int32)(m, l1, v188)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L27
	} else {
		goto L78
	}
L76:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_pg_reg_getcolor[1]))))
	if v194&int32(64) != 0 {
		v204 = v184
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v208 = v183
	goto L70
L78:
	;
	if v199 == int32(0) {
		v208 = v183
		goto L70
	} else {
		goto L79
	}
L79:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v204 = v203
	goto L72
L80:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v234 == int32(0) {
		v257 = v233
		goto L90
	} else {
		goto L91
	}
L81:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+2)))
	if v214 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v233 = v229 | v208
	goto L80
L83:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l1) {
		v233 = v208
		goto L80
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+48))
	v224 = m.T0[v223].(func(*base.Module, int32, int32) int32)(m, l1, v213)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L27
	} else {
		goto L88
	}
L86:
	;
	v219 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_pg_reg_getcolor[1]))))
	if v219 < int32(0) {
		v229 = v209
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v233 = v208
	goto L80
L88:
	;
	if v224 == int32(0) {
		v233 = v208
		goto L80
	} else {
		goto L89
	}
L89:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v229 = v228
	goto L82
L90:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v258 == int32(0) {
		v281 = v257
		goto L99
	} else {
		goto L100
	}
L91:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+2)))
	if v239 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v257 = v253 | v233
	goto L90
L93:
	;
	if base.Ui32(l1-int32(97)) < base.Ui32(int32(26)) {
		v253 = v234
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+32))
	v248 = m.T0[v247].(func(*base.Module, int32, int32) int32)(m, l1, v238)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L27
	} else {
		goto L97
	}
L96:
	;
	v257 = v233
	goto L90
L97:
	;
	if v248 == int32(0) {
		v257 = v233
		goto L90
	} else {
		goto L98
	}
L98:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v253 = v252
	goto L92
L99:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v282 == int32(0) {
		v305 = v281
		goto L108
	} else {
		goto L109
	}
L100:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+2)))
	if v263 == int32(1) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v281 = v277 | v257
	goto L99
L102:
	;
	if base.Ui32(l1-int32(65)) < base.Ui32(int32(26)) {
		v277 = v258
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+28))
	v272 = m.T0[v271].(func(*base.Module, int32, int32) int32)(m, l1, v262)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L27
	} else {
		goto L106
	}
L105:
	;
	v281 = v257
	goto L99
L106:
	;
	if v272 == int32(0) {
		v281 = v257
		goto L99
	} else {
		goto L107
	}
L107:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v277 = v276
	goto L101
L108:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v309 = int32(1)
	v315 = int32(*(*int16)(unsafe.Add(mBase, uint32(v306+v307*v51<<(uint(v309)%32)+v305<<(uint(v309)%32)))))
	return v315
L109:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+2)))
	if v287 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v305 = v301 | v281
	goto L108
L111:
	;
	if base.Ui32(l1-int32(33)) < base.Ui32(int32(94)) {
		v301 = v282
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v286)+8))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+36))
	v296 = m.T0[v295].(func(*base.Module, int32, int32) int32)(m, l1, v286)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L27
	} else {
		goto L115
	}
L114:
	;
	v305 = v281
	goto L108
L115:
	;
	if v296 == int32(0) {
		v305 = v281
		goto L108
	} else {
		goto L116
	}
L116:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v301 = v300
	goto L110
}
func F_pg_regerror(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v811 int32
	_ = v811
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v12 = int32(_a_F_pg_regerror_0)
	switch l0 - int32(101) {
	case 0:
		goto L5
	case 1:
		goto L4
	default:
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(144)
	return
L2:
	;
	v730 = F_strlen(m, v729)
	mBase = m.M
	if base.Ui32(v730+int32(1)) < base.Ui32(int32(100)) {
		goto L207
	} else {
		goto L208
	}
L3:
	;
	v701 = v12
	goto L197
L4:
	;
	v627 = l1
	goto L172
L5:
	;
	v15 = int32(0)
	v16 = int32(_a_F_pg_regerror_1)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[0])))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v19 == v15)|base.B2i32(v19 != v22) != 0 {
		v40 = v19
		v41 = v22
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v615
	v618 = v10 + int32(48)
	v622 = F_pg_sprintf(m, v618, int32(_a_F_pg_regerror_2), v10+int32(16))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L169
	} else {
		goto L170
	}
L7:
	;
	if v40-v41 == int32(0) {
		v615 = v15
		goto L6
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v25 = v16
	v26 = l1
	goto L10
L10:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v30 == int32(0) {
		v40 = v30
		v41 = v29
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v40 = v30
	v41 = v29
	goto L8
L12:
	;
	v33 = int32(1)
	if v30 == v29 {
		v25 = v25 + v33
		v26 = v26 + v33
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v46 = int32(_a_F_pg_regerror_3)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[1])))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v49 == int32(0))|base.B2i32(v49 != v52) != 0 {
		v70 = v49
		v71 = v52
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v70-v71 == int32(0) {
		v615 = int32(1)
		goto L6
	} else {
		goto L22
	}
L16:
	;
	goto L15
L17:
	;
	v55 = v46
	v56 = l1
	goto L18
L18:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v60 == int32(0) {
		v70 = v60
		v71 = v59
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v70 = v60
	v71 = v59
	goto L16
L20:
	;
	v63 = int32(1)
	if v60 == v59 {
		v55 = v55 + v63
		v56 = v56 + v63
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v76 = int32(_a_F_pg_regerror_4)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[2])))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v79 == int32(0))|base.B2i32(v79 != v82) != 0 {
		v100 = v79
		v101 = v82
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v100-v101 == int32(0) {
		v615 = int32(2)
		goto L6
	} else {
		goto L30
	}
L24:
	;
	goto L23
L25:
	;
	v85 = v76
	v86 = l1
	goto L26
L26:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v90 == int32(0) {
		v100 = v90
		v101 = v89
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v100 = v90
	v101 = v89
	goto L24
L28:
	;
	v93 = int32(1)
	if v90 == v89 {
		v85 = v85 + v93
		v86 = v86 + v93
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v106 = int32(_a_F_pg_regerror_5)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[3])))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v109 == int32(0))|base.B2i32(v109 != v112) != 0 {
		v130 = v109
		v131 = v112
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v130-v131 == int32(0) {
		v615 = int32(3)
		goto L6
	} else {
		goto L38
	}
L32:
	;
	goto L31
L33:
	;
	v115 = v106
	v116 = l1
	goto L34
L34:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	if v120 == int32(0) {
		v130 = v120
		v131 = v119
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v130 = v120
	v131 = v119
	goto L32
L36:
	;
	v123 = int32(1)
	if v120 == v119 {
		v115 = v115 + v123
		v116 = v116 + v123
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v136 = int32(_a_F_pg_regerror_6)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[4])))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v139 == int32(0))|base.B2i32(v139 != v142) != 0 {
		v160 = v139
		v161 = v142
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v160-v161 == int32(0) {
		v615 = int32(4)
		goto L6
	} else {
		goto L46
	}
L40:
	;
	goto L39
L41:
	;
	v145 = v136
	v146 = l1
	goto L42
L42:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+1)))
	if v150 == int32(0) {
		v160 = v150
		v161 = v149
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v160 = v150
	v161 = v149
	goto L40
L44:
	;
	v153 = int32(1)
	if v150 == v149 {
		v145 = v145 + v153
		v146 = v146 + v153
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v166 = int32(_a_F_pg_regerror_7)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[5])))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v169 == int32(0))|base.B2i32(v169 != v172) != 0 {
		v190 = v169
		v191 = v172
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v190-v191 == int32(0) {
		v615 = int32(5)
		goto L6
	} else {
		goto L54
	}
L48:
	;
	goto L47
L49:
	;
	v175 = v166
	v176 = l1
	goto L50
L50:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+1)))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	if v180 == int32(0) {
		v190 = v180
		v191 = v179
		goto L48
	} else {
		goto L52
	}
L51:
	;
	v190 = v180
	v191 = v179
	goto L48
L52:
	;
	v183 = int32(1)
	if v180 == v179 {
		v175 = v175 + v183
		v176 = v176 + v183
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v196 = int32(_a_F_pg_regerror_8)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[6])))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v199 == int32(0))|base.B2i32(v199 != v202) != 0 {
		v220 = v199
		v221 = v202
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v220-v221 == int32(0) {
		v615 = int32(6)
		goto L6
	} else {
		goto L62
	}
L56:
	;
	goto L55
L57:
	;
	v205 = v196
	v206 = l1
	goto L58
L58:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	if v210 == int32(0) {
		v220 = v210
		v221 = v209
		goto L56
	} else {
		goto L60
	}
L59:
	;
	v220 = v210
	v221 = v209
	goto L56
L60:
	;
	v213 = int32(1)
	if v210 == v209 {
		v205 = v205 + v213
		v206 = v206 + v213
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v226 = int32(_a_F_pg_regerror_9)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[7])))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v229 == int32(0))|base.B2i32(v229 != v232) != 0 {
		v250 = v229
		v251 = v232
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v250-v251 == int32(0) {
		v615 = int32(7)
		goto L6
	} else {
		goto L70
	}
L64:
	;
	goto L63
L65:
	;
	v235 = v226
	v236 = l1
	goto L66
L66:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+1)))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	if v240 == int32(0) {
		v250 = v240
		v251 = v239
		goto L64
	} else {
		goto L68
	}
L67:
	;
	v250 = v240
	v251 = v239
	goto L64
L68:
	;
	v243 = int32(1)
	if v240 == v239 {
		v235 = v235 + v243
		v236 = v236 + v243
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v256 = int32(_a_F_pg_regerror_10)
	v259 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[8])))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v259 == int32(0))|base.B2i32(v259 != v262) != 0 {
		v280 = v259
		v281 = v262
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v280-v281 == int32(0) {
		v615 = int32(8)
		goto L6
	} else {
		goto L78
	}
L72:
	;
	goto L71
L73:
	;
	v265 = v256
	v266 = l1
	goto L74
L74:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+1)))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+1)))
	if v270 == int32(0) {
		v280 = v270
		v281 = v269
		goto L72
	} else {
		goto L76
	}
L75:
	;
	v280 = v270
	v281 = v269
	goto L72
L76:
	;
	v273 = int32(1)
	if v270 == v269 {
		v265 = v265 + v273
		v266 = v266 + v273
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v286 = int32(_a_F_pg_regerror_11)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[9])))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v289 == int32(0))|base.B2i32(v289 != v292) != 0 {
		v310 = v289
		v311 = v292
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v310-v311 == int32(0) {
		v615 = int32(9)
		goto L6
	} else {
		goto L86
	}
L80:
	;
	goto L79
L81:
	;
	v295 = v286
	v296 = l1
	goto L82
L82:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+1)))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+1)))
	if v300 == int32(0) {
		v310 = v300
		v311 = v299
		goto L80
	} else {
		goto L84
	}
L83:
	;
	v310 = v300
	v311 = v299
	goto L80
L84:
	;
	v303 = int32(1)
	if v300 == v299 {
		v295 = v295 + v303
		v296 = v296 + v303
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v316 = int32(_a_F_pg_regerror_12)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[10])))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v319 == int32(0))|base.B2i32(v319 != v322) != 0 {
		v340 = v319
		v341 = v322
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v340-v341 == int32(0) {
		v615 = int32(10)
		goto L6
	} else {
		goto L94
	}
L88:
	;
	goto L87
L89:
	;
	v325 = v316
	v326 = l1
	goto L90
L90:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+1)))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+1)))
	if v330 == int32(0) {
		v340 = v330
		v341 = v329
		goto L88
	} else {
		goto L92
	}
L91:
	;
	v340 = v330
	v341 = v329
	goto L88
L92:
	;
	v333 = int32(1)
	if v330 == v329 {
		v325 = v325 + v333
		v326 = v326 + v333
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v346 = int32(_a_F_pg_regerror_13)
	v349 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[11])))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v349 == int32(0))|base.B2i32(v349 != v352) != 0 {
		v370 = v349
		v371 = v352
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v370-v371 == int32(0) {
		v615 = int32(11)
		goto L6
	} else {
		goto L102
	}
L96:
	;
	goto L95
L97:
	;
	v355 = v346
	v356 = l1
	goto L98
L98:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+1)))
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355)+1)))
	if v360 == int32(0) {
		v370 = v360
		v371 = v359
		goto L96
	} else {
		goto L100
	}
L99:
	;
	v370 = v360
	v371 = v359
	goto L96
L100:
	;
	v363 = int32(1)
	if v360 == v359 {
		v355 = v355 + v363
		v356 = v356 + v363
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v376 = int32(_a_F_pg_regerror_14)
	v379 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[12])))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v379 == int32(0))|base.B2i32(v379 != v382) != 0 {
		v400 = v379
		v401 = v382
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v400-v401 == int32(0) {
		v615 = int32(12)
		goto L6
	} else {
		goto L110
	}
L104:
	;
	goto L103
L105:
	;
	v385 = v376
	v386 = l1
	goto L106
L106:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+1)))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385)+1)))
	if v390 == int32(0) {
		v400 = v390
		v401 = v389
		goto L104
	} else {
		goto L108
	}
L107:
	;
	v400 = v390
	v401 = v389
	goto L104
L108:
	;
	v393 = int32(1)
	if v390 == v389 {
		v385 = v385 + v393
		v386 = v386 + v393
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v406 = int32(_a_F_pg_regerror_15)
	v409 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[13])))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v409 == int32(0))|base.B2i32(v409 != v412) != 0 {
		v430 = v409
		v431 = v412
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v430-v431 == int32(0) {
		v615 = int32(13)
		goto L6
	} else {
		goto L118
	}
L112:
	;
	goto L111
L113:
	;
	v415 = v406
	v416 = l1
	goto L114
L114:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+1)))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+1)))
	if v420 == int32(0) {
		v430 = v420
		v431 = v419
		goto L112
	} else {
		goto L116
	}
L115:
	;
	v430 = v420
	v431 = v419
	goto L112
L116:
	;
	v423 = int32(1)
	if v420 == v419 {
		v415 = v415 + v423
		v416 = v416 + v423
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v436 = int32(_a_F_pg_regerror_16)
	v439 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[14])))
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v439 == int32(0))|base.B2i32(v439 != v442) != 0 {
		v460 = v439
		v461 = v442
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if v460-v461 == int32(0) {
		v615 = int32(15)
		goto L6
	} else {
		goto L126
	}
L120:
	;
	goto L119
L121:
	;
	v445 = v436
	v446 = l1
	goto L122
L122:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+1)))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+1)))
	if v450 == int32(0) {
		v460 = v450
		v461 = v449
		goto L120
	} else {
		goto L124
	}
L123:
	;
	v460 = v450
	v461 = v449
	goto L120
L124:
	;
	v453 = int32(1)
	if v450 == v449 {
		v445 = v445 + v453
		v446 = v446 + v453
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v466 = int32(_a_F_pg_regerror_17)
	v469 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[15])))
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v469 == int32(0))|base.B2i32(v469 != v472) != 0 {
		v490 = v469
		v491 = v472
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v490-v491 == int32(0) {
		v615 = int32(16)
		goto L6
	} else {
		goto L134
	}
L128:
	;
	goto L127
L129:
	;
	v475 = v466
	v476 = l1
	goto L130
L130:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+1)))
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+1)))
	if v480 == int32(0) {
		v490 = v480
		v491 = v479
		goto L128
	} else {
		goto L132
	}
L131:
	;
	v490 = v480
	v491 = v479
	goto L128
L132:
	;
	v483 = int32(1)
	if v480 == v479 {
		v475 = v475 + v483
		v476 = v476 + v483
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	v496 = int32(_a_F_pg_regerror_18)
	v499 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[16])))
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v499 == int32(0))|base.B2i32(v499 != v502) != 0 {
		v520 = v499
		v521 = v502
		goto L136
	} else {
		goto L137
	}
L135:
	;
	if v520-v521 == int32(0) {
		v615 = int32(17)
		goto L6
	} else {
		goto L142
	}
L136:
	;
	goto L135
L137:
	;
	v505 = v496
	v506 = l1
	goto L138
L138:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+1)))
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505)+1)))
	if v510 == int32(0) {
		v520 = v510
		v521 = v509
		goto L136
	} else {
		goto L140
	}
L139:
	;
	v520 = v510
	v521 = v509
	goto L136
L140:
	;
	v513 = int32(1)
	if v510 == v509 {
		v505 = v505 + v513
		v506 = v506 + v513
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v526 = int32(_a_F_pg_regerror_19)
	v529 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[17])))
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v529 == int32(0))|base.B2i32(v529 != v532) != 0 {
		v550 = v529
		v551 = v532
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if v550-v551 == int32(0) {
		v615 = int32(18)
		goto L6
	} else {
		goto L150
	}
L144:
	;
	goto L143
L145:
	;
	v535 = v526
	v536 = l1
	goto L146
L146:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536)+1)))
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535)+1)))
	if v540 == int32(0) {
		v550 = v540
		v551 = v539
		goto L144
	} else {
		goto L148
	}
L147:
	;
	v550 = v540
	v551 = v539
	goto L144
L148:
	;
	v543 = int32(1)
	if v540 == v539 {
		v535 = v535 + v543
		v536 = v536 + v543
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	v556 = int32(_a_F_pg_regerror_20)
	v559 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[18])))
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v559 == int32(0))|base.B2i32(v559 != v562) != 0 {
		v580 = v559
		v581 = v562
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v580-v581 == int32(0) {
		v615 = int32(19)
		goto L6
	} else {
		goto L158
	}
L152:
	;
	goto L151
L153:
	;
	v565 = v556
	v566 = l1
	goto L154
L154:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566)+1)))
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+1)))
	if v570 == int32(0) {
		v580 = v570
		v581 = v569
		goto L152
	} else {
		goto L156
	}
L155:
	;
	v580 = v570
	v581 = v569
	goto L152
L156:
	;
	v573 = int32(1)
	if v570 == v569 {
		v565 = v565 + v573
		v566 = v566 + v573
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v587 = int32(_a_F_pg_regerror_21)
	v590 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_regerror[19])))
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if base.B2i32(v590 == int32(0))|base.B2i32(v590 != v593) != 0 {
		v611 = v590
		v612 = v593
		goto L160
	} else {
		goto L161
	}
L159:
	;
	if v611-v612 != 0 {
		goto L166
	} else {
		goto L167
	}
L160:
	;
	goto L159
L161:
	;
	v596 = v587
	v597 = l1
	goto L162
L162:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597)+1)))
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+1)))
	if v601 == int32(0) {
		v611 = v601
		v612 = v600
		goto L160
	} else {
		goto L164
	}
L163:
	;
	v611 = v601
	v612 = v600
	goto L160
L164:
	;
	v604 = int32(1)
	if v601 == v600 {
		v596 = v596 + v604
		v597 = v597 + v604
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	v614 = int32(-1)
	goto L168
L167:
	;
	v614 = int32(20)
	goto L168
L168:
	;
	v615 = v614
	goto L6
L169:
	;
	return
L170:
	;
	v729 = v618
	goto L2
L171:
	;
	v676 = v12
	goto L187
L172:
	;
	v632 = v627 + int32(1)
	v633 = int32(*(*int8)(unsafe.Add(mBase, uint32(v627))))
	v634 = F___isspace(m, v633)
	mBase = m.M
	if v634 != 0 {
		v627 = v632
		goto L172
	} else {
		goto L174
	}
L173:
	;
	v635 = int32(1)
	switch v633&int32(255) - int32(43) {
	case 0:
		v641 = v635
		goto L176
	default:
		v643 = v633
		v644 = v627
		v645 = v635
		goto L175
	case 2:
		goto L177
	}
L174:
	;
	goto L173
L175:
	;
	v646 = int32(0)
	v648 = v643 - int32(48)
	if base.Ui32(v648) <= base.Ui32(int32(9)) {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	v642 = int32(*(*int8)(unsafe.Add(mBase, uint32(v632))))
	v643 = v642
	v644 = v632
	v645 = v641
	goto L175
L177:
	;
	v641 = int32(0)
	goto L176
L178:
	;
	v651 = v646
	v652 = v648
	v653 = v644
	goto L181
L179:
	;
	v665 = v646
	goto L180
L180:
	;
	if v645 != 0 {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	v655 = int32(10)
	v657 = v651*v655 - v652
	v658 = int32(*(*int8)(unsafe.Add(mBase, uint32(v653)+1)))
	v662 = v658 - int32(48)
	if base.Ui32(v662) < base.Ui32(v655) {
		v651 = v657
		v652 = v662
		v653 = v653 + int32(1)
		goto L181
	} else {
		goto L183
	}
L182:
	;
	v665 = v657
	goto L180
L183:
	;
	goto L182
L184:
	;
	v671 = int32(0) - v665
	goto L186
L185:
	;
	v671 = v665
	goto L186
L186:
	;
	goto L171
L187:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v676)))
	if int32(0) <= v679 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	if int32(0) <= v679 {
		goto L193
	} else {
		goto L194
	}
L189:
	;
	if v671 != v679 {
		v676 = v676 + int32(12)
		goto L187
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	goto L188
L192:
	;
	goto L191
L193:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v676)+4))
	v729 = v688
	goto L2
L194:
	;
	goto L195
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v671
	v691 = v10 + int32(48)
	v695 = F_pg_sprintf(m, v691, int32(_a_F_pg_regerror_22), v10+int32(32))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L169
	} else {
		goto L196
	}
L196:
	;
	v729 = v691
	goto L2
L197:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v705 = int32(0)
	v706 = base.B2i32(v704 < v705)
	if v706 == v705 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if v706 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L199:
	;
	if l0 != v704 {
		v701 = v701 + int32(12)
		goto L197
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	goto L198
L202:
	;
	goto L201
L203:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v701)+8))
	v729 = v715
	goto L2
L204:
	;
	goto L205
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	v718 = v10 + int32(48)
	v720 = F_pg_sprintf(m, v718, int32(_a_F_pg_regerror_23), v10)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L169
	} else {
		goto L206
	}
L206:
	;
	v729 = v718
	goto L2
L207:
	;
	if (v729^l1)&int32(3) != 0 {
		goto L213
	} else {
		goto L214
	}
L208:
	;
	goto L209
L209:
	;
	base.MemoryCopy(m, l1, v729, int32(99))
	v811 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+99)) = uint8(v811)
	goto L1
L210:
	;
	goto L1
L211:
	;
	goto L210
L212:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v789))) = uint8(v788)
	if v788&int32(255) == int32(0) {
		goto L211
	} else {
		goto L227
	}
L213:
	;
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729))))
	v787 = v729
	v788 = v740
	v789 = l1
	goto L212
L214:
	;
	goto L215
L215:
	;
	if v729&int32(3) != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v744 = v729
	v746 = l1
	goto L219
L217:
	;
	v758 = v729
	v760 = l1
	goto L218
L218:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	v765 = int32(-2139062144)
	if (int32(16843008)-v762|v762)&v765 != v765 {
		v787 = v758
		v788 = v762
		v789 = v760
		goto L212
	} else {
		goto L223
	}
L219:
	;
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
	*(*uint8)(unsafe.Add(mBase, uint32(v746))) = uint8(v747)
	if v747 == int32(0) {
		goto L211
	} else {
		goto L221
	}
L220:
	;
	v758 = v754
	v760 = v752
	goto L218
L221:
	;
	v751 = int32(1)
	v752 = v746 + v751
	v754 = v744 + v751
	if v754&int32(3) != 0 {
		v744 = v754
		v746 = v752
		goto L219
	} else {
		goto L222
	}
L222:
	;
	goto L220
L223:
	;
	v770 = v758
	v771 = v762
	v772 = v760
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v772))) = v771
	v774 = int32(4)
	v775 = v772 + v774
	v777 = v770 + v774
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v770)+4))
	v782 = int32(-2139062144)
	if (int32(16843008)-v779|v779)&v782 == v782 {
		v770 = v777
		v771 = v779
		v772 = v775
		goto L224
	} else {
		goto L226
	}
L225:
	;
	v787 = v777
	v788 = v779
	v789 = v775
	goto L212
L226:
	;
	goto L225
L227:
	;
	v796 = v787
	v798 = v789
	goto L228
L228:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v798)+1)) = uint8(v799)
	v801 = int32(1)
	if v799 != 0 {
		v796 = v796 + v801
		v798 = v798 + v801
		goto L228
	} else {
		goto L230
	}
L229:
	;
	goto L211
L230:
	;
	goto L229
}
func F_pg_regfree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	if l0 != 0 {
		v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
		m.T0[v3].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_pg_sequence_parameters(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = base.I32_wrap_i64(v10)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sequence_parameters[0]))
	v15 = F_pg_class_aclcheck(m, v11, v13, int64(262))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		if v15 == int32(0) {
			v24 = F_get_call_result_type(m, l0, int32(0), v8+int32(92))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				if v24 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int64(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pg_sequence_parameters_0), int32(0))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_sequence_parameters_1), int32(1762), int32(_a_F_pg_sequence_parameters_2))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v28 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+27)) = v28
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v28
					v35 = F_SearchSysCache1(m, int32(61), v10&int64(4294967295))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						if v35 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
								F_errmsg_internal(m, int32(_a_F_pg_sequence_parameters_3), v8)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_sequence_parameters_1), int32(1768), int32(_a_F_pg_sequence_parameters_2))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+22)))
							v41 = v39 + v40
							v42 = *(*int64)(unsafe.Add(mBase, uint32(v41)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v42
							v44 = *(*int64)(unsafe.Add(mBase, uint32(v41)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v44
							v46 = *(*int64)(unsafe.Add(mBase, uint32(v41)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v46
							v48 = *(*int64)(unsafe.Add(mBase, uint32(v41)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v48
							v50 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v41)+48)))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v50
							v52 = *(*int64)(unsafe.Add(mBase, uint32(v41)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+72)) = v52
							v54 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v41)+4)))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+80)) = v54
							F_ReleaseCatCache(m, v35)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
								v63 = F_heap_form_tuple(m, v58, v8+int32(32), v8+int32(24))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int64(0)
								} else {
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
									v66 = F_HeapTupleHeaderGetDatum(m, v65)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int64(0)
									} else {
										m.G0 = v8 + int32(96)
										return v66
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int64(0)
				} else {
					v79 = F_get_rel_name(m, v11)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v79
						F_errmsg(m, int32(_a_F_pg_sequence_parameters_4), v8+int32(16))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_sequence_parameters_1), int32(1759), int32(_a_F_pg_sequence_parameters_2))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
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
func F_pg_server_to_any(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	v4 = int32(0)
	if base.B2i32(l2 == v4)|base.B2i32(l1 <= v4) != 0 {
		v40 = l0
		return v40
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_pg_server_to_any[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if l2 == v12 {
			v40 = l0
			return v40
		} else {
			if v12 == int32(0) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l2*int32(28))+uint32(_c_F_pg_server_to_any[1])))
				v21 = m.T0[v20].(func(*base.Module, int32, int32) int32)(m, l0, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if l1 == v21 {
						v40 = l0
						return v40
					} else {
						F_report_invalid_encoding(m, l2, l0+v21, l1-v21)
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
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_server_to_any[2]))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				if v32 == l2 {
					v35 = F_perform_default_encoding_conversion(m, l0, l1, int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						return v35
					}
				} else {
					v38 = F_pg_do_encoding_conversion(m, l0, l1, v12, l2)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = v38
						return v40
					}
				}
			}
		}
	}
}
func F_pg_set_timing_clock_source(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	v3 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_set_timing_clock_source[0])) = v3
	*(*int64)(unsafe.Add(mBase, _c_F_pg_set_timing_clock_source[1])) = v3
	*(*int32)(unsafe.Add(mBase, _c_F_pg_set_timing_clock_source[2])) = l0
	return
}
func F_pg_size_pretty_numeric(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var __phi24 int32
	_ = __phi24
	var v26 int32
	_ = v26
	var __phi26 int32
	_ = __phi26
	var v27 int32
	_ = v27
	var __phi27 int32
	_ = __phi27
	var v28 int32
	_ = v28
	var __phi28 int32
	_ = __phi28
	var v30 int32
	_ = v30
	var __phi30 int32
	_ = __phi30
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	__phi24 = int32(_a_F_pg_size_pretty_numeric_0)
	__phi26 = int32(_a_F_pg_size_pretty_numeric_1)
	__phi27 = v19
	__phi28 = int32(_a_F_pg_size_pretty_numeric_2)
	__phi30 = int32(_a_F_pg_size_pretty_numeric_3)
	v24 = __phi24
	v26 = __phi26
	v27 = __phi27
	v28 = __phi28
	v30 = __phi30
	goto L3
L3:
	;
	v33 = int32(0)
	v36 = base.I64_extend_i32_u(v27)
	v37 = F_DirectFunctionCall1Coll(m, int32(1404), v33, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+8)))
	if v82 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	goto L4
L6:
	;
	v40 = F_pg_detoast_datum(m, base.I32_wrap_i64(v37))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v43 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+4)))
	v44 = F_int64_to_numeric(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v47 = F_DirectFunctionCall2Coll(m, int32(1403), v33, base.I64_extend_i32_u(v40), base.I64_extend_i32_u(v44))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v47 != int64(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v73 = v24
	v76 = v27
	v78 = v26
	goto L5
L11:
	;
	goto L12
L12:
	;
	v56 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
	v57 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+21)))
	v58 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
	v59 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+20)))
	v64 = F_int64_to_numeric(m, int64(1)<<(uint(v56+(v57-(v58+v59)))%64))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v67 = F_DirectFunctionCall2Coll(m, int32(1405), int32(0), v36, base.I64_extend_i32_u(v64))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v71 = F_pg_detoast_datum(m, base.I32_wrap_i64(v67))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v69 != 0 {
		__phi24 = v30
		__phi26 = v28
		__phi27 = v71
		__phi28 = v69
		__phi30 = v30 + int32(12)
		v24 = __phi24
		v26 = __phi26
		v27 = __phi27
		v28 = __phi28
		v30 = __phi30
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v73 = v30
	v76 = v71
	v78 = v28
	goto L5
L17:
	;
	v86 = F_int64_to_numeric(m, int64(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v121 = v76
	goto L19
L19:
	;
	v123 = F_DirectFunctionCall1Coll(m, int32(664), int32(0), base.I64_extend_i32_u(v121))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L30
	}
L20:
	;
	v89 = F_int64_to_numeric(m, int64(1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v92 = F_int64_to_numeric(m, int64(2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v95 = int32(0)
	v100 = base.I64_extend_i32_u(v76)
	v102 = F_DirectFunctionCall2Coll(m, int32(1407), v95, v100, base.I64_extend_i32_u(v86))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v102 == int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v106 = int32(19)
	goto L26
L25:
	;
	v106 = int32(1406)
	goto L26
L26:
	;
	v109 = F_DirectFunctionCall2Coll(m, v106, int32(0), v100, base.I64_extend_i32_u(v89))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v112 = F_DirectFunctionCall2Coll(m, int32(1405), v95, v109, base.I64_extend_i32_u(v92))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v115 = F_pg_detoast_datum(m, base.I32_wrap_i64(v112))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v121 = v115
	goto L19
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v78
	*(*uint32)(unsafe.Add(mBase, uint32(v12))) = uint32(v123)
	v128 = F_psprintf(m, int32(_a_F_pg_size_pretty_numeric_4), v12)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v130 = F_cstring_to_text(m, v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	m.G0 = v12 + int32(16)
	return base.I64_extend_i32_u(v130)
}
func F_pg_strftime(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pg_strftime[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v5
	v16 = l0 + l1
	v19 = F__fmt(m, l2, l3, l0, v16, v10+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_pg_strftime[0])) = int32(61)
			if l1 == int32(0) {
				v41 = v5
			} else {
				v39 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v39)
				v41 = v5
			}
		} else {
			if v19 == v16 {
				*(*int32)(unsafe.Add(mBase, _c_F_pg_strftime[0])) = int32(68)
				if l1 != 0 {
					v39 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v39)
					v41 = v5
				} else {
					v41 = v5
				}
			} else {
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, _c_F_pg_strftime[0])) = v13
				v41 = v19 - l0
			}
		}
		m.G0 = v10 + int32(16)
		return v41
	}
}
func F_pg_strtoint32(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_pg_strtoint32_safe(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pg_strtoint32_safe(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v78 int64
	_ = v78
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v331 int64
	_ = v331
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v18 = base.B2i32(v16 == int32(45))
	v19 = l0 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v24 = (v20 - int32(48)) & int32(255)
	if base.Ui32(int32(10)) <= base.Ui32(v24) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v409
L2:
	;
	F_errsave_finish(m, l1, int32(_a_F_pg_strtoint32_safe_0), v401, int32(_a_F_pg_strtoint32_safe_1))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L81
	} else {
		goto L90
	}
L3:
	;
	v374 = int32(0)
	v375 = F_errsave_start(m, l1)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L81
	} else {
		goto L86
	}
L4:
	;
	v347 = int32(0)
	v348 = F_errsave_start(m, l1)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L81
	} else {
		goto L82
	}
L5:
	;
	v97 = l0
	v101 = v16
	goto L22
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v29 = v27 - int32(48)
	if base.Ui32(v29&int32(255)) <= base.Ui32(int32(9)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v39 = v19 + int32(1)
	v40 = v24
	v41 = v29
	goto L10
L8:
	;
	v64 = v27
	v66 = v24
	goto L9
L9:
	;
	if v64&int32(255) != 0 {
		goto L5
	} else {
		goto L14
	}
L10:
	;
	if base.Ui32(int32(214748364)) < base.Ui32(v40) {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v64 = v53
	v66 = v52
	goto L9
L12:
	;
	v48 = int32(10)
	v50 = int32(255)
	v52 = v40*v48 + v41&v50
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v57 = v53 - int32(48)
	if base.Ui32(v57&v50) < base.Ui32(v48) {
		v39 = v39 + int32(1)
		v40 = v52
		v41 = v57
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	if v16 == int32(45) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v78 = int64(0) - base.I64_extend_i32_u(v66)
	if v78 != base.I64_extend32_s(v78) {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v66 < int32(0) {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v409 = base.I32_wrap_i64(v78)
	goto L1
L19:
	;
	v409 = v66
	goto L1
L20:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v120 != int32(48) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v118 = v97 + int32(1)
	v119 = v18
	goto L20
L22:
	;
	if base.Ui32(v101-int32(9)) < base.Ui32(int32(5)) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v113 = int32(1)
	v118 = v97 + v113
	v119 = v113
	goto L20
L24:
	;
	goto L23
L25:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	v97 = v97 + int32(1)
	v101 = v110
	goto L22
L26:
	;
	switch v101 - int32(32) {
	case 0:
		goto L25
	default:
		v118 = v97
		v119 = v18
		goto L20
	case 11:
		goto L21
	case 13:
		goto L24
	}
L27:
	;
	if v298 == v299 {
		goto L3
	} else {
		goto L69
	}
L28:
	;
	v260 = v118
	v262 = int32(0)
	v263 = v120
	goto L60
L29:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	switch v123 - int32(66) {
	case 0, 32:
		goto L30
	default:
		goto L28
	case 13, 45:
		goto L31
	case 22, 54:
		goto L32
	}
L30:
	;
	v219 = v118 + int32(2)
	v222 = v219
	v224 = int32(0)
	goto L52
L31:
	;
	v179 = v118 + int32(2)
	v182 = v179
	v184 = int32(0)
	goto L44
L32:
	;
	v128 = v118 + int32(2)
	v131 = v128
	v133 = int32(0)
	goto L33
L33:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	goto L35
L34:
	;
	goto L3
L35:
	;
	if base.B2i32(base.Ui32(v139-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v139|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if base.Ui32(int32(134217728)) < base.Ui32(v133) {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v139 != int32(95) {
		v298 = v131
		v299 = v128
		v300 = v133
		v301 = v139
		goto L27
	} else {
		goto L40
	}
L39:
	;
	v155 = int32(*(*int8)(unsafe.Add(mBase, uint32(v139)+uint32(_c_F_pg_strtoint32_safe[0]))))
	v131 = v131 + int32(1)
	v133 = v155 + v133<<(uint(int32(4))%32)
	goto L33
L40:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)))
	if v161 == int32(0) {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	goto L42
L42:
	;
	if base.B2i32(base.Ui32(v161-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v161|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		v131 = v131 + int32(1)
		goto L33
	} else {
		goto L43
	}
L43:
	;
	goto L34
L44:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v190&int32(248) == int32(48) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L3
L46:
	;
	if base.Ui32(int32(268435456)) < base.Ui32(v184) {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v190 != int32(95) {
		v298 = v182
		v299 = v179
		v300 = v184
		v301 = v190
		goto L27
	} else {
		goto L50
	}
L49:
	;
	v182 = v182 + int32(1)
	v184 = (v190-int32(48))&int32(255) | v184<<(uint(int32(3))%32)
	goto L44
L50:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
	if base.Ui32(int32(248)) <= base.Ui32((v208-int32(56))&int32(255)) {
		v182 = v182 + int32(1)
		goto L44
	} else {
		goto L51
	}
L51:
	;
	goto L45
L52:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v230&int32(254) == int32(48) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L3
L54:
	;
	if base.Ui32(int32(1073741824)) < base.Ui32(v224) {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v230 != int32(95) {
		v298 = v222
		v299 = v219
		v300 = v224
		v301 = v230
		goto L27
	} else {
		goto L58
	}
L57:
	;
	v241 = int32(1)
	v222 = v222 + v241
	v224 = (v230-int32(48))&int32(255) | v224<<(uint(v241)%32)
	goto L52
L58:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	if base.Ui32(int32(254)) <= base.Ui32((v248-int32(50))&int32(255)) {
		v222 = v222 + int32(1)
		goto L52
	} else {
		goto L59
	}
L59:
	;
	goto L53
L60:
	;
	v271 = (v263 - int32(48)) & int32(255)
	if base.Ui32(v271) <= base.Ui32(int32(9)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L3
L62:
	;
	if base.Ui32(int32(214748364)) < base.Ui32(v262) {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v263&int32(255) != int32(95) {
		v298 = v260
		v299 = v118
		v300 = v262
		v301 = v263
		goto L27
	} else {
		goto L66
	}
L65:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	v260 = v260 + int32(1)
	v262 = v262*int32(10) + v271
	v263 = v279
	goto L60
L66:
	;
	if v260 == v118 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	if base.Ui32((v287-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v260 = v260 + int32(1)
		v263 = v287
		goto L60
	} else {
		goto L68
	}
L68:
	;
	goto L61
L69:
	;
	v309 = v298
	v312 = v301
	goto L70
L70:
	;
	v318 = v312 & int32(255)
	if base.B2i32(base.Ui32(v318-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v318 == int32(32)) != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L4
L72:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	v309 = v309 + int32(1)
	v312 = v326
	goto L70
L73:
	;
	if v318 != 0 {
		goto L3
	} else {
		goto L75
	}
L74:
	;
	goto L71
L75:
	;
	if v119 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v331 = int64(0) - base.I64_extend_i32_u(v300)
	if v331 != base.I64_extend32_s(v331) {
		goto L4
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if int32(0) <= v300 {
		v409 = v300
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v409 = base.I32_wrap_i64(v331)
	goto L1
L80:
	;
	goto L74
L81:
	;
	return int32(0)
L82:
	;
	if v348 == int32(0) {
		v409 = v347
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(_a_F_pg_strtoint32_safe_2)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
	F_errmsg(m, int32(_a_F_pg_strtoint32_safe_3), v13)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v395 = v347
	v401 = int32(611)
	goto L2
L86:
	;
	if v375 == int32(0) {
		v409 = v374
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L81
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_pg_strtoint32_safe_2)
	F_errmsg(m, int32(_a_F_pg_strtoint32_safe_4), v13+int32(16))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v395 = v374
	v401 = int32(617)
	goto L2
L90:
	;
	v409 = v395
	goto L1
}
func F_pg_strxfrm(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v6 == int32(0) {
		v9 = F_strlen(m, l1)
		mBase = m.M
		if base.Ui32(l2) <= base.Ui32(v9) {
			v21 = v9
			return v21
		} else {
			if v9 != 0 {
				base.MemoryCopy(m, l0, l1, v9)
			} else {
			}
			v13 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0+v9))) = uint8(v13)
			return v9
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v17 = m.T0[v16].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l2, l1, l3)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = v17
			return v21
		}
	}
}
func F_pg_table_size(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_try_relation_open(m, v4, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int64(0)
		} else {
			v16 = F_calculate_table_size(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				F_relation_close(m, v6, int32(1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					return v16
				}
			}
		}
	}
}
func F_pg_timezone_abbrevs_zone(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
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
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v191 int64
	_ = v191
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int64
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int64
	_ = v345
	var v351 int64
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int64
	_ = v359
	var v360 int64
	_ = v360
	var v363 int64
	_ = v363
	var v369 int32
	_ = v369
	var v371 int64
	_ = v371
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int64
	_ = v391
	var v392 int32
	_ = v392
	var v393 int64
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v439 int64
	_ = v439
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+46)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+44)) = uint16(v2)
	v16 = *(*int64)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[0]))
	v18 = base.I64_div_s(v16, int64(1000000))
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v18 + int64(946684800)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v23 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L6
	} else {
		goto L93
	}
L3:
	;
	v26 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	goto L12
L6:
	;
	return int64(0)
L7:
	;
	v30 = int32(_a_F_pg_timezone_abbrevs_zone_0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[1]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[1])) = v33
	v36 = F_palloc(m, int32(4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v38 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v36
	v44 = F_get_call_result_type(m, l0, v38, v9+int32(48))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v44 != int32(1) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v48
	*(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[1])) = v31
	goto L5
L11:
	;
	m.G0 = v9 + int32(80)
	return v439
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[2]))
	v60 = int32(0)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v62 < v60 {
		v78 = v60
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v78 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L13
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+268))
	if v65 <= v62 {
		v78 = v60
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v68 = int32(_a_F_pg_timezone_abbrevs_zone_1)
	v70 = F_strlen(m, v59+v62+v68)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v70 + v62 + int32(1)
	v78 = v59 + v68 + v62
	goto L14
L17:
	;
	v81 = v78
	goto L20
L18:
	;
	goto L19
L19:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L6
	} else {
		goto L92
	}
L20:
	;
	v85 = int32(_a_F_pg_timezone_abbrevs_zone_2)
	v89 = m.G0
	v91 = v89 - int32(32)
	v92 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v91)+24)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v91)+16)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v91)+8)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = v92
	v100 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[3])))
	if v100 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L19
L22:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[2]))
	v402 = int32(0)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v404 < v402 {
		v420 = v402
		goto L88
	} else {
		goto L89
	}
L23:
	;
	v169 = F_strlen(m, v81)
	mBase = m.M
	if v168 != v169 {
		goto L22
	} else {
		goto L42
	}
L24:
	;
	v168 = int32(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[4])))
	if v104 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v108 = v81
	goto L30
L28:
	;
	goto L29
L29:
	;
	v118 = v85
	v119 = v100
	goto L33
L30:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v114 == v100 {
		v108 = v108 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v168 = v108 - v81
	goto L23
L32:
	;
	goto L31
L33:
	;
	v126 = v91 + int32(base.Ui32(v119)>>(uint(int32(3))%32))&int32(28)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v128 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v127 | v128<<(uint(v119)%32)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	if v132 != 0 {
		v118 = v118 + v128
		v119 = v132
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v135 == int32(0) {
		v158 = v81
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v168 = v158 - v81
	goto L23
L37:
	;
	v139 = v81
	v140 = v135
	goto L38
L38:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v91+int32(base.Ui32(v140)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v148)>>(uint(v140)%32))&int32(1) == int32(0) {
		v158 = v139
		goto L36
	} else {
		goto L40
	}
L39:
	;
	v158 = v156
	goto L36
L40:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+1)))
	v156 = v139 + int32(1)
	if v154 != 0 {
		v139 = v156
		v140 = v154
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[2]))
	v179 = int32(0)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v178)+268))
	if v186 <= v179 {
		v339 = v179
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v339 == int32(0) {
		goto L22
	} else {
		goto L78
	}
L44:
	;
	goto L43
L45:
	;
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(32))))
	v194 = int32(0)
	goto L46
L46:
	;
	v205 = v194 + (v178 + int32(_a_F_pg_timezone_abbrevs_zone_1))
	v206 = F_strcmp(m, v81, v205)
	mBase = m.M
	if v206 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v178)+260))
	if v212 <= int32(0) {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v207 = F_strlen(m, v205)
	mBase = m.M
	v210 = v207 + v194 + int32(1)
	if v210 < v186 {
		v194 = v210
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	v339 = v179
	goto L44
L52:
	;
	v257 = v178 + int32(_a_F_pg_timezone_abbrevs_zone_3)
	v259 = v178 + int32(_a_F_pg_timezone_abbrevs_zone_4)
	v260 = v249
	goto L66
L53:
	;
	v249 = int32(0)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v219 = v212
	v224 = int32(0)
	goto L56
L56:
	;
	v232 = int32(1)
	v233 = (v219 + v224) >> (uint(v232) % 32)
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v178+int32(280)+v233<<(uint(int32(3))%32))))
	v240 = base.B2i32(v191 < v239)
	if v191 < v239 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v249 = v241
	goto L52
L58:
	;
	v241 = v224
	goto L60
L59:
	;
	v241 = v233 + v232
	goto L60
L60:
	;
	if v191 < v239 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v242 = v233
	goto L63
L62:
	;
	v242 = v219
	goto L63
L63:
	;
	if v241 < v242 {
		v219 = v242
		v224 = v241
		goto L56
	} else {
		goto L64
	}
L64:
	;
	goto L57
L65:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(28)))) = v324
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(24)))) = v326
	v339 = int32(1)
	goto L44
L66:
	;
	if int32(0) < v260 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v178)+uint32(_c_F_pg_timezone_abbrevs_zone[5])))
	v286 = v259 + v283<<(uint(int32(4))%32)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+8))
	if v287 == v194 {
		v318 = v286
		goto L65
	} else {
		goto L72
	}
L68:
	;
	v275 = v260 - int32(1)
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257+v275))))
	v280 = v259 + v277<<(uint(int32(4))%32)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	if v281 != v194 {
		v260 = v275
		goto L66
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	v318 = v280
	goto L65
L72:
	;
	if v212 <= v249 {
		v339 = v179
		goto L44
	} else {
		goto L73
	}
L73:
	;
	v295 = v249
	goto L74
L74:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v257))))
	v306 = v259 + v303<<(uint(int32(4))%32)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+8))
	if v307 == v194 {
		v318 = v306
		goto L65
	} else {
		goto L76
	}
L75:
	;
	v339 = v179
	goto L44
L76:
	;
	v310 = v295 + int32(1)
	if v212 != v310 {
		v295 = v310
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v343 = F_cstring_to_text(m, v81)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	v345 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v345
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v345
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = base.I64_extend_i32_u(v343)
	v351 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+28)))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v351 * int64(1000000)
	v356 = F_palloc(m, int32(16))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	v359 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+12)))
	v360 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+16)))
	v363 = v359 + v360*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v363-int64(2147483648)) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = base.I64_extend_i32_u(v356)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = base.I64_extend_i32_u(base.B2i32(v378 != int32(0)))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	v388 = F_heap_form_tuple(m, v383, v9+int32(48), v9+int32(44))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L6
	} else {
		goto L85
	}
L82:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v356)+12)) = uint32(v363)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v356)+8)) = v369
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	*(*int64)(unsafe.Add(mBase, uint32(v356))) = v371
	goto L84
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v388)+16))
	v391 = F_HeapTupleHeaderGetDatum(m, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v56)))
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = v393 + int64(1)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v397)+20)) = int32(1)
	v439 = v391
	goto L11
L87:
	;
	if v420 != 0 {
		v81 = v420
		goto L20
	} else {
		goto L91
	}
L88:
	;
	goto L87
L89:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v401)+268))
	if v407 <= v404 {
		v420 = v402
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v410 = int32(_a_F_pg_timezone_abbrevs_zone_1)
	v412 = F_strlen(m, v401+v404+v410)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v412 + v404 + int32(1)
	v420 = v401 + v410 + v404
	goto L88
L91:
	;
	goto L21
L92:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v429)+20)) = int32(2)
	v432 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v432)
	v439 = int64(0)
	goto L11
L93:
	;
	F_errmsg_internal(m, int32(_a_F_pg_timezone_abbrevs_zone_5), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_pg_timezone_abbrevs_zone_6), int32(_a_F_pg_timezone_abbrevs_zone_7), int32(_a_F_pg_timezone_abbrevs_zone_8))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_trigger_depth(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	v3 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_trigger_depth[0])))
	return v3
}
func F_pg_ts_dict_is_visible(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_TSDictionaryIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v25 = int64(0)
		} else {
			v25 = base.I64_extend_i32_u(v14)
		}
		m.G0 = v7 + int32(16)
		return v25
	}
}
func F_pg_tzenumerate_next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(2080)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 < v2 {
		v273 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L7
	} else {
		goto L65
	}
L2:
	;
	m.G0 = v11 + int32(2080)
	return v273
L3:
	;
	v19 = l0 + int32(88)
	v21 = l0 + int32(48)
	v23 = l0 + int32(8)
	v25 = v13
	goto L4
L4:
	;
	v33 = v25 << (uint(int32(2)) % 32)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23+v33)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v21)))
	v38 = F_ReadDir(m, v35, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v273 = int32(0)
	goto L2
L6:
	;
	if int32(0) <= v264 {
		v25 = v264
		goto L4
	} else {
		goto L64
	}
L7:
	;
	return int32(0)
L8:
	;
	if v38 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23+v44<<(uint(int32(2))%32))))
	F_FreeDir(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+19)))
	if v62 == int32(46) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v21+v51<<(uint(int32(2))%32))))
	F_pfree(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v60 = v58 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v60
	v264 = v60
	goto L6
L14:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v264 = v263
	goto L6
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v21+v65<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v38 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v69
	v75 = v11 + int32(32)
	v80 = F_pg_snprintf(m, v75, int32(2048), int32(_a_F_pg_tzenumerate_next_0), v11+int32(16))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v84 = F_get_dirent_type(m, v75, v38, int32(1), int32(21))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	if v84 == int32(3) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(9) <= v88 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v133 = F_tzload(m, v128+(v11+int32(32)), int32(0), l0+int32(344))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L7
	} else {
		goto L29
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88 + int32(1)
	v94 = F_pstrdup(m, v75)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21+v96<<(uint(int32(2))%32)))) = v94
	v101 = F_AllocateDir(m, v75)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v104 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v23+v103<<(uint(v104)%32)))) = v101
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v23+v108<<(uint(v104)%32))))
	if v112 != 0 {
		v264 = v108
		goto L6
	} else {
		goto L24
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v75
	F_errmsg(m, int32(_a_F_pg_tzenumerate_next_1), v11)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_pg_tzenumerate_next_2), int32(463), int32(_a_F_pg_tzenumerate_next_3))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	if v133 != 0 {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v135 = F_pg_tz_acceptable(m, v19)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	if v135 == int32(0) {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v142 = v139 + (v11 + int32(32))
	goto L36
L33:
	;
	v273 = v19
	goto L2
L34:
	;
	v259 = F_strlen(m, v248)
	mBase = m.M
	goto L33
L36:
	;
	goto L37
L37:
	;
	v149 = int32(255)
	if (v19^v142)&int32(3) != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v252 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v249))) = uint8(v252)
	goto L34
L39:
	;
	v233 = v228
	v234 = v229
	v235 = v230
	goto L60
L40:
	;
	if v223 == int32(0) {
		v248 = v221
		v249 = v222
		goto L38
	} else {
		goto L59
	}
L41:
	;
	v221 = v142
	v222 = v19
	v223 = v149
	goto L40
L42:
	;
	goto L43
L43:
	;
	v153 = int32(0)
	if base.B2i32(v142&int32(3) == v153)|int32(0) == v153 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v189 == int32(0) {
		v248 = v186
		v249 = v187
		goto L38
	} else {
		goto L53
	}
L45:
	;
	v165 = v142
	v166 = v19
	v167 = v149
	goto L48
L46:
	;
	goto L47
L47:
	;
	v186 = v142
	v187 = v19
	v188 = v149
	v189 = int32(1)
	goto L44
L48:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v169)
	if v169 == int32(0) {
		v228 = v165
		v229 = v166
		v230 = v167
		goto L39
	} else {
		goto L50
	}
L49:
	;
	v186 = v180
	v187 = v174
	v188 = v176
	v189 = v178
	goto L44
L50:
	;
	v173 = int32(1)
	v174 = v166 + v173
	v176 = v167 - v173
	v177 = int32(0)
	v178 = base.B2i32(v176 != v177)
	v180 = v165 + v173
	if v180&int32(3) == v177 {
		v186 = v180
		v187 = v174
		v188 = v176
		v189 = v178
		goto L44
	} else {
		goto L51
	}
L51:
	;
	if v176 != 0 {
		v165 = v180
		v166 = v174
		v167 = v176
		goto L48
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if base.B2i32(v192 == int32(0))|base.B2i32(base.Ui32(v188) < base.Ui32(int32(4))) != 0 {
		v221 = v186
		v222 = v187
		v223 = v188
		goto L40
	} else {
		goto L54
	}
L54:
	;
	v199 = v186
	v200 = v187
	v201 = v188
	goto L55
L55:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v207 = int32(-2139062144)
	if (int32(16843008)-v204|v204)&v207 != v207 {
		v228 = v199
		v229 = v200
		v230 = v201
		goto L39
	} else {
		goto L57
	}
L56:
	;
	v221 = v215
	v222 = v213
	v223 = v217
	goto L40
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v204
	v212 = int32(4)
	v213 = v200 + v212
	v215 = v199 + v212
	v217 = v201 - v212
	if base.Ui32(int32(3)) < base.Ui32(v217) {
		v199 = v215
		v200 = v213
		v201 = v217
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v228 = v221
	v229 = v222
	v230 = v223
	goto L39
L60:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	*(*uint8)(unsafe.Add(mBase, uint32(v234))) = uint8(v237)
	if v237 == int32(0) {
		v248 = v233
		v249 = v234
		goto L38
	} else {
		goto L62
	}
L61:
	;
	v248 = v244
	v249 = v242
	goto L38
L62:
	;
	v241 = int32(1)
	v242 = v234 + v241
	v244 = v233 + v241
	v246 = v235 - v241
	if v246 != 0 {
		v233 = v244
		v234 = v242
		v235 = v246
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L5
L65:
	;
	F_errmsg_internal(m, int32(_a_F_pg_tzenumerate_next_4), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_pg_tzenumerate_next_2), int32(455), int32(_a_F_pg_tzenumerate_next_3))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_tzset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(_a_F_pg_tzset_0)
	m.G0 = v8
	v10 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(int32(255)) < base.Ui32(v10) {
		v287 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(_a_F_pg_tzset_0)
	return v287
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_pg_tzset[0]))
	if v14 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+520)) = int64(102873056674048)
	v25 = F_hash_create(m, int32(_a_F_pg_tzset_1), int64(4), v8+int32(512), int32(24))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v32 = v14
	goto L5
L5:
	;
	v34 = v8 + int32(256)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_tzset[0])) = v25
	if v25 == int32(0) {
		v287 = v2
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v32 = v25
	goto L5
L9:
	;
	v36 = l0
	v38 = v34
	v40 = v35
	goto L12
L10:
	;
	v62 = v34
	v63 = v32
	goto L11
L11:
	;
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v65)
	v71 = F_hash_search(m, v63, v8+int32(256), v65, v65)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L19
	}
L12:
	;
	if base.Ui32((v40-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_pg_tzset[0]))
	v62 = v54
	v63 = v59
	goto L11
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v51)
	v53 = int32(1)
	v54 = v38 + v53
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v55 != 0 {
		v36 = v36 + v53
		v38 = v54
		v40 = v55
		goto L12
	} else {
		goto L18
	}
L15:
	;
	v49 = v40 - int32(32)
	goto L17
L16:
	;
	v49 = v40
	goto L17
L17:
	;
	v51 = v49 & int32(255)
	goto L14
L18:
	;
	goto L13
L19:
	;
	if v71 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v287 = v71 + int32(256)
	goto L1
L21:
	;
	goto L22
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v8)+256))
	if v75 == int32(_a_F_pg_tzset_2) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_pg_tzset[0]))
	v201 = F_hash_search(m, v196, v8+int32(256), int32(1), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L57
	}
L24:
	;
	v117 = v8 + int32(256)
	if (v117^v8)&int32(3) != 0 {
		goto L39
	} else {
		goto L40
	}
L25:
	;
	v83 = F_tzparse(m, v8+int32(256), v8+int32(512), int32(1))
	mBase = m.M
	if v83 != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v98 = v8 + int32(256)
	v100 = v8 + int32(512)
	v101 = F_tzload(m, v98, v8, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L32
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	F_errmsg_internal(m, int32(_a_F_pg_tzset_3), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_pg_tzset_4), int32(278), int32(_a_F_pg_tzset_5))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	if v101 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L33:
	;
	v105 = int32(0)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+256)))
	if v106 == int32(58) {
		v287 = v105
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v109 = int32(0)
	v110 = F_tzparse(m, v98, v100, v109)
	mBase = m.M
	if v110 == v109 {
		v287 = v105
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L24
L36:
	;
	goto L23
L37:
	;
	goto L36
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v172))) = uint8(v171)
	if v171&int32(255) == int32(0) {
		goto L37
	} else {
		goto L53
	}
L39:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v170 = v117
	v171 = v123
	v172 = v8
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v117&int32(3) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v127 = v117
	v129 = v8
	goto L45
L43:
	;
	v141 = v117
	v143 = v8
	goto L44
L44:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v148 = int32(-2139062144)
	if (int32(16843008)-v145|v145)&v148 != v148 {
		v170 = v141
		v171 = v145
		v172 = v143
		goto L38
	} else {
		goto L49
	}
L45:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v130)
	if v130 == int32(0) {
		goto L37
	} else {
		goto L47
	}
L46:
	;
	v141 = v137
	v143 = v135
	goto L44
L47:
	;
	v134 = int32(1)
	v135 = v129 + v134
	v137 = v127 + v134
	if v137&int32(3) != 0 {
		v127 = v137
		v129 = v135
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v153 = v141
	v154 = v145
	v155 = v143
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v154
	v157 = int32(4)
	v158 = v155 + v157
	v160 = v153 + v157
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v165 = int32(-2139062144)
	if (int32(16843008)-v162|v162)&v165 == v165 {
		v153 = v160
		v154 = v162
		v155 = v158
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v170 = v160
	v171 = v162
	v172 = v158
	goto L38
L52:
	;
	goto L51
L53:
	;
	v179 = v170
	v181 = v172
	goto L54
L54:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v181)+1)) = uint8(v182)
	v184 = int32(1)
	if v182 != 0 {
		v179 = v179 + v184
		v181 = v181 + v184
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L37
L56:
	;
	goto L55
L57:
	;
	v204 = v201 + int32(256)
	if (v8^v204)&int32(3) != 0 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v279 = int32(512)
	base.MemoryCopy(m, v201+v279, v8+v279, int32(_a_F_pg_tzset_6))
	v287 = v204
	goto L1
L59:
	;
	goto L58
L60:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v259))) = uint8(v258)
	if v258&int32(255) == int32(0) {
		goto L59
	} else {
		goto L75
	}
L61:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v257 = v8
	v258 = v210
	v259 = v204
	goto L60
L62:
	;
	goto L63
L63:
	;
	if v8&int32(3) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v214 = v8
	v216 = v204
	goto L67
L65:
	;
	v228 = v8
	v230 = v204
	goto L66
L66:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v235 = int32(-2139062144)
	if (int32(16843008)-v232|v232)&v235 != v235 {
		v257 = v228
		v258 = v232
		v259 = v230
		goto L60
	} else {
		goto L71
	}
L67:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v217)
	if v217 == int32(0) {
		goto L59
	} else {
		goto L69
	}
L68:
	;
	v228 = v224
	v230 = v222
	goto L66
L69:
	;
	v221 = int32(1)
	v222 = v216 + v221
	v224 = v214 + v221
	if v224&int32(3) != 0 {
		v214 = v224
		v216 = v222
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v240 = v228
	v241 = v232
	v242 = v230
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v241
	v244 = int32(4)
	v245 = v242 + v244
	v247 = v240 + v244
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	v252 = int32(-2139062144)
	if (int32(16843008)-v249|v249)&v252 == v252 {
		v240 = v247
		v241 = v249
		v242 = v245
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v257 = v247
	v258 = v249
	v259 = v245
	goto L60
L74:
	;
	goto L73
L75:
	;
	v266 = v257
	v268 = v259
	goto L76
L76:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+1)) = uint8(v269)
	v271 = int32(1)
	if v269 != 0 {
		v266 = v266 + v271
		v268 = v268 + v271
		goto L76
	} else {
		goto L78
	}
L77:
	;
	goto L59
L78:
	;
	goto L77
}
func F_pg_u_prop_case_ignorable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(517)
	v11 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_u_prop_case_ignorable[0]))))
	return int32(base.Ui32(v40&int32(16)) >> (uint(int32(4)) % 32))
L4:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 << (uint(int32(3)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_u_prop_case_ignorable[1])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	if v34 <= v33 {
		v10 = v33
		v11 = v34
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v33 = v10
	v34 = v16 + int32(1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_u_prop_case_ignorable[2])))
	if base.Ui32(v27) <= base.Ui32(l0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(1)
L11:
	;
	goto L12
L12:
	;
	v33 = v16 - int32(1)
	v34 = v11
	goto L6
L13:
	;
	goto L5
}
func F_pg_u_prop_cased(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v52 = int32(691)
	v53 = int32(0)
	goto L15
L2:
	;
	v10 = int32(3408)
	v11 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_pg_u_prop_cased[0]))))
	return int32(base.Ui32(v41&int32(8)) >> (uint(int32(3)) % 32))
L5:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 * int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_u_prop_cased[1])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_u_prop_cased[2]))))
	if v34 != int32(3) {
		goto L1
	} else {
		goto L14
	}
L7:
	;
	goto L6
L8:
	;
	if v32 <= v31 {
		v10 = v31
		v11 = v32
		goto L5
	} else {
		goto L13
	}
L9:
	;
	v31 = v10
	v32 = v16 + int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_c_F_pg_u_prop_cased[3])))
	if base.Ui32(v27) <= base.Ui32(l0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v31 = v16 - int32(1)
	v32 = v11
	goto L8
L13:
	;
	goto L1
L14:
	;
	return int32(1)
L15:
	;
	v58 = base.I32_div_s(v52+v53, int32(2))
	v60 = v58 << (uint(int32(3)) % 32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_u_prop_cased[4])))
	if base.Ui32(v63) < base.Ui32(l0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v81 = int32(659)
	v82 = int32(0)
	goto L25
L17:
	;
	if v76 <= v75 {
		v52 = v75
		v53 = v76
		goto L15
	} else {
		goto L24
	}
L18:
	;
	v75 = v52
	v76 = v58 + int32(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_pg_u_prop_cased[5])))
	if base.Ui32(v69) <= base.Ui32(l0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(1)
L22:
	;
	goto L23
L23:
	;
	v75 = v58 - int32(1)
	v76 = v53
	goto L17
L24:
	;
	goto L16
L25:
	;
	v87 = base.I32_div_s(v81+v82, int32(2))
	v89 = v87 << (uint(int32(3)) % 32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+uint32(_c_F_pg_u_prop_cased[6])))
	if base.Ui32(v92) < base.Ui32(l0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	return int32(0)
L27:
	;
	if v105 <= v104 {
		v81 = v104
		v82 = v105
		goto L25
	} else {
		goto L34
	}
L28:
	;
	v104 = v81
	v105 = v87 + int32(1)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+uint32(_c_F_pg_u_prop_cased[7])))
	if base.Ui32(v98) <= base.Ui32(l0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(1)
L32:
	;
	goto L33
L33:
	;
	v104 = v87 - int32(1)
	v105 = v82
	goto L27
L34:
	;
	goto L26
}
func F_pg_ulltoa_n(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	v3 = int32(0)
	if l0 == int64(0) {
		v12 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v12)
		return int32(1)
	} else {
		v20 = int32(1233)
		v25 = int32(base.Ui32((base.I32_wrap_i64(base.I64_clz(l0))^int32(63))*v20+v20) >> (uint(int32(12)) % 32))
		v28 = *(*int64)(unsafe.Add(mBase, uint32(v25<<(uint(int32(3))%32))+uint32(_c_F_pg_ulltoa_n[0])))
		v30 = v25 + base.B2i32(base.Ui64(v28) <= base.Ui64(l0))
		if base.Ui64(int64(100000000)) <= base.Ui64(l0) {
			v34 = l0
			v37 = v3
			for {
				v43 = l1 + v30 - v37
				v44 = int32(8)
				v47 = base.I64_div_u_s(v34, int64(100000000))
				v51 = base.I32_wrap_i64(v34 + v47*int64(4194967296))
				v53 = base.I32_div_u_s(v51, int32(_a_F_pg_ulltoa_n_0))
				v54 = int32(1)
				v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53<<(uint(v54)%32))+uint32(_c_F_pg_ulltoa_n[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v43-v44))) = uint16(v56)
				v60 = int32(_a_F_pg_ulltoa_n_1)
				v61 = base.I32_div_u_s(v51, v60)
				v62 = int32(100)
				v63 = base.I32_rem_u_s(v61, v62)
				v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63<<(uint(v54)%32))+uint32(_c_F_pg_ulltoa_n[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v43-int32(6)))) = uint16(v66)
				v72 = v51 - v61*v60
				v73 = int32(_a_F_pg_ulltoa_n_2)
				v76 = base.I32_div_u_s(v72&v73, v62)
				v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76<<(uint(v54)%32))+uint32(_c_F_pg_ulltoa_n[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v43-int32(4)))) = uint16(v79)
				v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v72-v76*v62)&v73<<(uint(v54)%32))+uint32(_c_F_pg_ulltoa_n[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v43-int32(2)))) = uint16(v90)
				v93 = v37 + v44
				if base.Ui64(int64(9999999999999999)) < base.Ui64(v34) {
					v34 = v47
					v37 = v93
					continue
				} else {
					break
				}
				break
			}
			v96 = v47
			v99 = v93
		} else {
			v96 = l0
			v99 = v3
		}
		v105 = base.I32_wrap_i64(v96)
		if base.Ui64(int64(10000)) <= base.Ui64(v96) {
			v109 = l1 + v30 - v99
			v110 = int32(4)
			v113 = base.I32_div_u_s(v105, int32(_a_F_pg_ulltoa_n_1))
			v116 = v105 + v113*int32(-10000)
			v117 = int32(100)
			v118 = base.I32_div_u_s(v116, v117)
			v119 = int32(1)
			v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118<<(uint(v119)%32))+uint32(_c_F_pg_ulltoa_n[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v109-v110))) = uint16(v121)
			v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v116-v118*v117)<<(uint(v119)%32))+uint32(_c_F_pg_ulltoa_n[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v109-int32(2)))) = uint16(v130)
			v134 = v113
			v135 = v99 | v110
		} else {
			v134 = v105
			v135 = v99
		}
		if base.Ui32(int32(100)) <= base.Ui32(v134) {
			v143 = int32(2)
			v145 = int32(_a_F_pg_ulltoa_n_2)
			v147 = int32(100)
			v148 = base.I32_div_u_s(v134&v145, v147)
			v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v134-v148*v147)&v145<<(uint(int32(1))%32))+uint32(_c_F_pg_ulltoa_n[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v30-v135-v143))) = uint16(v156)
			v160 = v148
			v161 = v135 + v143
		} else {
			v160 = v134
			v161 = v135
		}
		if base.Ui32(int32(10)) <= base.Ui32(v160) {
			v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160<<(uint(int32(1))%32))+uint32(_c_F_pg_ulltoa_n[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v30-v161-int32(2)))) = uint16(v170)
			return v30
		} else {
			v174 = v160 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v174)
			return v30
		}
	}
}
func F_pg_utf_dsplen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v29 int32
	_ = v29
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	v6 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	v8 = v6 & int32(255)
	if v6 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v8&int32(224) == int32(192) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v65 = v8
	goto L3
L3:
	;
	if v65 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v58))))
	v65 = v59 | v61&int32(63)
	goto L3
L5:
	;
	v58 = int32(1)
	v59 = v8 << (uint(int32(6)) % 32) & int32(1984)
	goto L4
L6:
	;
	goto L7
L7:
	;
	if v8&int32(240) == int32(224) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v58 = int32(2)
	v59 = v8<<(uint(int32(12))%32)&int32(_a_F_pg_utf_dsplen_0) | v29&int32(63)<<(uint(int32(6))%32)
	goto L4
L9:
	;
	goto L10
L10:
	;
	if v8&int32(248) != int32(240) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(-1)
L12:
	;
	goto L13
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v47 = int32(63)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v58 = int32(3)
	v59 = v8<<(uint(int32(18))%32)&int32(_a_F_pg_utf_dsplen_1) | v46&v47<<(uint(int32(12))%32) | v52&v47<<(uint(int32(6))%32)
	goto L4
L14:
	;
	return int32(0)
L15:
	;
	goto L16
L16:
	;
	if base.B2i32(base.Ui32(v65) < base.Ui32(int32(32)))|base.B2i32(base.Ui32(int32(_a_F_pg_utf_dsplen_2)) < base.Ui32(v65))|base.B2i32(base.Ui32(v65-int32(127)) < base.Ui32(int32(33))) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if base.Ui32(v65-int32(_a_F_pg_utf_dsplen_3)) < base.Ui32(int32(-917827)) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v155 = int32(-1)
	goto L19
L19:
	;
	return v155
L20:
	;
	return int32(1)
L21:
	;
	goto L22
L22:
	;
	v92 = int32(0)
	v94 = int32(340)
	goto L23
L23:
	;
	v99 = base.I32_div_s(v92+v94, int32(2))
	v101 = v99 << (uint(int32(3)) % 32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_pg_utf_dsplen[0])))
	if base.Ui32(v104) < base.Ui32(v65) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if base.Ui32(v65-int32(_a_F_pg_utf_dsplen_4)) < base.Ui32(int32(-257790)) {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	if v116 <= v117 {
		v92 = v116
		v94 = v117
		goto L23
	} else {
		goto L32
	}
L26:
	;
	v116 = v99 + int32(1)
	v117 = v94
	goto L25
L27:
	;
	goto L28
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_pg_utf_dsplen[1])))
	if base.Ui32(v110) <= base.Ui32(v65) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	goto L31
L31:
	;
	v116 = v92
	v117 = v99 - int32(1)
	goto L25
L32:
	;
	goto L24
L33:
	;
	return int32(1)
L34:
	;
	goto L35
L35:
	;
	v129 = int32(0)
	v130 = int32(122)
	goto L36
L36:
	;
	v134 = base.I32_div_s(v129+v130, int32(2))
	v136 = v134 << (uint(int32(3)) % 32)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_pg_utf_dsplen[2])))
	if base.Ui32(v139) < base.Ui32(v65) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v155 = int32(1)
	goto L19
L38:
	;
	if v151 <= v152 {
		v129 = v151
		v130 = v152
		goto L36
	} else {
		goto L45
	}
L39:
	;
	v151 = v134 + int32(1)
	v152 = v130
	goto L38
L40:
	;
	goto L41
L41:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_pg_utf_dsplen[3])))
	if base.Ui32(v145) <= base.Ui32(v65) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	return int32(2)
L43:
	;
	goto L44
L44:
	;
	v151 = v129
	v152 = v134 - int32(1)
	goto L38
L45:
	;
	goto L37
}
func F_pg_verify_mbstr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_c_F_pg_verify_mbstr[0])))
	v11 = m.T0[v10].(func(*base.Module, int32, int32) int32)(m, l1, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l3|base.B2i32(v11 == l2) == int32(0) {
			F_report_invalid_encoding(m, l0, l1+v11, l2-v11)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return base.B2i32(l2 == v11)
		}
	}
}
func F_pg_walfile_name(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_walfile_name[0])))
	if v13 == int32(1) {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name[1]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+308))
		v21 = base.B2i32(v19 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_walfile_name[0])) = uint8(v21)
		v23 = v21
	} else {
		v23 = int32(0)
	}
	if v23 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_walfile_name_0), int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_pg_walfile_name_1)
					F_errhint(m, int32(_a_F_pg_walfile_name_2), v8)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_walfile_name_3), int32(480), int32(_a_F_pg_walfile_name_4))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v48 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name[2])))
		v50 = *(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name[1]))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+300))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v51
		v53 = base.I64_div_u_s(v10, v48)
		v56 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name[2])))
		v57 = base.I64_div_u_s(int64(4294967296), v56)
		v58 = base.I64_div_u_s(v53, v57)
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+20)) = uint32(v58)
		v61 = v53 - v57*v58
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+24)) = uint32(v61)
		v64 = v8 + int32(32)
		v69 = F_pg_snprintf(m, v64, int32(64), int32(_a_F_pg_walfile_name_5), v8+int32(16))
		mBase = m.M
		v70 = m.ExcPending
		if v70 != 0 {
			return int64(0)
		} else {
			v71 = F_cstring_to_text(m, v64)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int64(0)
			} else {
				m.G0 = v8 + int32(96)
				return base.I64_extend_i32_u(v71)
			}
		}
	}
}
func F_pg_walfile_name_offset(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int64
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v169 int64
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v201 int32
	_ = v201
	v8 = m.G0
	v10 = v8 - int32(112)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[0])))
	if v15 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[1]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+308))
	v23 = base.B2i32(v21 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[0])) = uint8(v23)
	v25 = v23
	goto L4
L3:
	;
	v25 = int32(0)
	goto L4
L4:
	;
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v50 = F_CreateTemplateTupleDesc(m, int32(2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L14
	}
L8:
	;
	return int64(0)
L9:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_errmsg(m, int32(_a_F_pg_walfile_name_offset_0), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_pg_walfile_name_offset_1)
	F_errhint(m, int32(_a_F_pg_walfile_name_offset_2), v10)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_pg_walfile_name_offset_3), int32(421), int32(_a_F_pg_walfile_name_offset_4))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	F_TupleDescInitEntry(m, v50, int32(1), int32(_a_F_pg_walfile_name_offset_5), int32(25), int32(-1), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	F_TupleDescInitEntry(m, v50, int32(2), int32(_a_F_pg_walfile_name_offset_6), int32(23), int32(-1), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v66 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v66 < v75 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v153 = F_BlessTupleDesc(m, v50)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L8
	} else {
		goto L36
	}
L18:
	;
	v79 = v50 + int32(28)
	v86 = v66
	v87 = v75
	v89 = v66
	goto L22
L19:
	;
	v143 = v66
	v150 = v75
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v143
	goto L17
L21:
	;
	v143 = v137
	v150 = v116
	goto L20
L22:
	;
	v95 = v79 + v75<<(uint(int32(3))%32) + v86*int32(100)
	v98 = v79 + v86<<(uint(int32(3))%32)
	if v75 != v87 {
		v116 = v87
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v137 = v75
	goto L21
L24:
	;
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+2)))
	if v117 <= int32(0) {
		v137 = v86
		goto L21
	} else {
		goto L32
	}
L25:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+7)))
	if v100 != int32(118) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v116 = v86
	goto L24
L27:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+4)))
	if v103 != int32(1) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+6)))
	if v106&int32(6) != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+2)))
	if v109 <= int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+90)))
	if v112 != int32(118) {
		v116 = v75
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+90)))
	if v120 == int32(118) {
		v137 = v86
		goto L21
	} else {
		goto L33
	}
L33:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+5)))
	v129 = (v89 + v123 - int32(1)) & (int32(0) - v123)
	if int32(_a_F_pg_walfile_name_offset_7) < v129 {
		v137 = v86
		goto L21
	} else {
		goto L34
	}
L34:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v98))) = uint16(v129)
	v135 = v86 + int32(1)
	if v135 != v75 {
		v86 = v135
		v87 = v116
		v89 = v129 + v117
		goto L22
	} else {
		goto L35
	}
L35:
	;
	goto L23
L36:
	;
	v156 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[2])))
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[1]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+300))
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v159
	v161 = base.I64_div_u_s(v12, v156)
	v164 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[2])))
	v165 = base.I64_div_u_s(int64(4294967296), v164)
	v166 = base.I64_div_u_s(v161, v165)
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v166)
	v169 = v161 - v165*v166
	*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v169)
	v172 = v10 + int32(48)
	v177 = F_pg_snprintf(m, v172, int32(64), int32(_a_F_pg_walfile_name_offset_8), v10+int32(16))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v179 = F_cstring_to_text(m, v172)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v181 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+30)) = uint16(v181)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_u(v179)
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = base.I64_extend_i32_u(base.I32_wrap_i64(v12) & (v187 - int32(1)))
	v197 = F_heap_form_tuple(m, v153, v10+int32(32), v10+int32(30))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	v200 = F_HeapTupleHeaderGetDatum(m, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	m.G0 = v10 + int32(112)
	return v200
}
