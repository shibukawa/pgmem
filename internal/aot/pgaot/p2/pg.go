package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_PGSemaphoreLock(m *base.Module, l0 int32) {
	var v7 int32
	_ = v7
	Fn13857(m, l0, int32(_a_F_PGSemaphoreLock_0), int32(335), int32(_a_F_PGSemaphoreLock_1), int32(2))
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 <= int32(0) {
		v92 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v92
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
		v92 = l0
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
	if base.Ui32(int32(34)) < base.Ui32(l2) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_pg_any_to_server[2]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v85 == l2 {
		goto L33
	} else {
		goto L34
	}
L16:
	;
	v53 = int32(0)
	goto L23
L17:
	;
	goto L16
L18:
	;
	goto L19
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2*int32(28))+uint32(_c_F_pg_any_to_server[1])))
	v44 = m.T0[v43].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if l1 == v44 {
		v92 = l0
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_report_invalid_encoding(m, l2, l0+v44, l1-v44)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v56 = l0 + v53
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(v56))))
	if int32(0) < v57 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L9
	} else {
		goto L29
	}
L25:
	;
	v61 = v53 + int32(1)
	if l1 != v61 {
		v53 = v61
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v92 = l0
	goto L1
L29:
	;
	F_errcode(m, int32(17301634))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_pg_any_to_server[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v73
	F_errmsg(m, int32(_a_F_pg_any_to_server_0), v8)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_pg_any_to_server_1), int32(724), int32(_a_F_pg_any_to_server_2))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v88 = F_perform_default_encoding_conversion(m, l0, l1, int32(1))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v90 = F_pg_do_encoding_conversion(m, l0, l1, l2, v15)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L37
	}
L36:
	;
	v92 = v88
	goto L1
L37:
	;
	v92 = v90
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
func F_pg_backend_pid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backend_pid[0]))
	return v3
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
func F_pg_buffercache_evict(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+34)) = uint16(v2)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_get_call_result_type(m, l0, v2, v8+int32(44))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L42
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L3
	} else {
		goto L38
	}
L3:
	;
	return int32(0)
L4:
	;
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v22 = F_superuser(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L3
	} else {
		goto L35
	}
L8:
	;
	if v22 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v12 <= int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict[0]))
	if v29 < v12 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v33 = m.G0
	v35 = v33 - int32(32)
	m.G0 = v35
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict[1]))
	F_ResourceOwnerEnlarge(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = int32(_a_F_pg_buffercache_evict_0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = int32(_a_F_pg_buffercache_evict_1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = int32(_a_F_pg_buffercache_evict_2)
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = int64(0)
	v57 = v44 + v12<<(uint(int32(6))%32)
	v59 = v57 - int32(40)
	v60 = int32(_a_F_pg_buffercache_evict_3)
	v62 = base.AtomicRmwOr32(m, v59, v51, v60)
	if v62&v60 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	goto L17
L15:
	;
	goto L16
L16:
	;
	v87 = int32(_a_F_pg_buffercache_evict_4)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict[3]))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v35+int32(8))+8))
	if v90 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	F_perform_spin_delay(m, v35+int32(8))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	v74 = int32(_a_F_pg_buffercache_evict_3)
	v76 = base.AtomicRmwOr32(m, v59, int32(0), v74)
	if v76&v74 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v109 = F_EvictUnpinnedBufferInternal(m, v57+int32(-64), v8+int32(33))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L32
	}
L22:
	;
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_evict[3])) = v105
	goto L22
L24:
	;
	if int32(999) < v88 {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v88 < int32(11) {
		goto L22
	} else {
		goto L31
	}
L27:
	;
	v95 = int32(900)
	if v95 <= v88 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = v95
	goto L30
L29:
	;
	v98 = v88
	goto L30
L30:
	;
	v105 = v98 + int32(100)
	goto L23
L31:
	;
	v105 = v88 - int32(1)
	goto L23
L32:
	;
	m.G0 = v35 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v109
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+33)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
	v122 = F_heap_form_tuple(m, v117, v8+int32(36), v8+int32(34))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v125 = F_HeapTupleHeaderGetDatum(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	m.G0 = v8 + int32(48)
	return v125
L35:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_evict_5), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_6), int32(691), int32(_a_F_pg_buffercache_evict_7))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
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
	F_errcode(m, int32(16797828))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(_a_F_pg_buffercache_evict_7)
	F_errmsg(m, int32(_a_F_pg_buffercache_evict_8), v8+int32(16))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_6), int32(672), int32(_a_F_pg_buffercache_evict_9))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_evict_10), v8)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_evict_6), int32(696), int32(_a_F_pg_buffercache_evict_7))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_create_physical_replication_slot(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_get_call_result_type(m, l0, int32(0), v8)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(1) {
			F_CheckSlotPermissions(m)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_CheckSlotRequirements(m)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = int32(0)
					if v10 != 0 {
						v27 = int32(2)
					} else {
						v27 = v24
					}
					v28 = int32(0)
					F_ReplicationSlotCreate(m, v12, v24, v27, v28, v28, v28)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v11 != 0 {
							F_ReplicationSlotReserveWal(m)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_ReplicationSlotMarkDirty(m)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									F_ReplicationSlotSave(m)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int32(0)
									} else {
										v39 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v8)+6)) = uint8(v39)
										v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_create_physical_replication_slot[0]))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v43 + int32(24)
										v47 = *(*int64)(unsafe.Add(mBase, uint32(v43)+104))
										v48 = F_Int64GetDatum(m, v47)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v48
											v59 = v39
											*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)) = uint8(v59)
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
											v67 = F_heap_form_tuple(m, v62, v8+int32(8), v8+int32(6))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
												v70 = F_HeapTupleHeaderGetDatum(m, v69)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_ReplicationSlotRelease(m)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														m.G0 = v8 + int32(16)
														return v70
													}
												}
											}
										}
									}
								}
							}
						} else {
							v51 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+6)) = uint8(v51)
							v54 = *(*int32)(unsafe.Add(mBase, _c_F_pg_create_physical_replication_slot[0]))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v54 + int32(24)
							v59 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)) = uint8(v59)
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v67 = F_heap_form_tuple(m, v62, v8+int32(8), v8+int32(6))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
								v70 = F_HeapTupleHeaderGetDatum(m, v69)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_ReplicationSlotRelease(m)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(16)
										return v70
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
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_create_physical_replication_slot_0), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_create_physical_replication_slot_1), int32(77), int32(_a_F_pg_create_physical_replication_slot_2))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
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
func F_pg_cursor(m *base.Module, l0 int32) int32 {
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v5 = m.G0
	v7 = v5 + int32(-64)
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
	return int32(0)
L2:
	;
	v16 = v5 + int32(-20)
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
	m.G0 = v7 - int32(-64)
	return int32(0)
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
	v75 = F_hash_seq_search(m, v5+int32(-20))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L17
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
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v39
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	v43 = F_cstring_to_text(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	v47 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v46 & v47
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(base.Ui32(v46)>>(uint(v47)%32)) & v47
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(base.Ui32(v46)>>(uint(int32(5))%32)) & v47
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v27)+128))
	v61 = F_Int64GetDatum(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v61
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	F_tuplestore_putvalues(m, v64, v65, v5+int32(-48), v5+int32(-56))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	if v75 != 0 {
		v23 = v75
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L9
}
func F_pg_database_size_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_SearchSysCacheExists(m, int32(21), v10, v2, v2, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v18 = F_calculate_database_size(m, v10)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int64(0) {
					v22 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
					v27 = int32(0)
					m.G0 = v7 + int32(16)
					return v27
				} else {
					v25 = F_Int64GetDatum(m, v18)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = v25
						m.G0 = v7 + int32(16)
						return v27
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
					F_errmsg(m, int32(_a_F_pg_database_size_oid_0), v7)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_database_size_oid_1), int32(180), int32(_a_F_pg_database_size_oid_2))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
func F_pg_freespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_relation_open(m, v11, int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
		switch v18 - int32(83) {
		case 0, 22, 26, 31, 33:
			if base.Ui64(int64(4294967295)) <= base.Ui64(v10) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_freespace_0), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_freespace_1), int32(47), int32(_a_F_pg_freespace_2))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
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
				v63 = F_GetRecordedFreeSpace(m, v13, base.I32_wrap_i64(v10))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_relation_close(m, v13, int32(1))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return base.I32_extend16_s(v63)
					}
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v28 + int32(4)
					F_errmsg(m, int32(_a_F_pg_freespace_3), v7)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
						v36 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+119)))
						F_errdetail_relkind_not_supported(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_freespace_1), int32(42), int32(_a_F_pg_freespace_2))
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
func F_pg_get_catalog_foreign_keys(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == int32(0) {
		v17 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(_a_F_pg_get_catalog_foreign_keys_0)
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_catalog_foreign_keys[0]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_pg_get_catalog_foreign_keys[0])) = v24
			v29 = F_get_call_result_type(m, l0, int32(0), v11+int32(16))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v29 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pg_get_catalog_foreign_keys_1), int32(0))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_get_catalog_foreign_keys_2), int32(510), int32(_a_F_pg_get_catalog_foreign_keys_3))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
					v34 = F_BlessTupleDesc(m, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v34
						v39 = F_palloc(m, int32(28))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_fmgr_info(m, int32(750), v39)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v39
								*(*int32)(unsafe.Add(mBase, _c_F_pg_get_catalog_foreign_keys[0])) = v22
								v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
								v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
								if base.Ui64(v51) <= base.Ui64(int64(218)) {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
									v55 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v55)
									*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v55
									v61 = base.I32_wrap_i64(v51) * int32(20)
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_pg_get_catalog_foreign_keys[1])))
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v62
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_pg_get_catalog_foreign_keys[2])))
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_pg_get_catalog_foreign_keys[3])))
									v69 = F_FunctionCall3Coll(m, v54, v55, v66, int32(25), int32(-1))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v64
										*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v69
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_pg_get_catalog_foreign_keys[4]))))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_pg_get_catalog_foreign_keys[5])))
										v78 = F_FunctionCall3Coll(m, v54, int32(0), v75, int32(25), int32(-1))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v73
											v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_pg_get_catalog_foreign_keys[6]))))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v81
											*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v78
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
											v89 = F_heap_form_tuple(m, v84, v11+int32(16), v11+int32(8))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												v91 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
												*(*int64)(unsafe.Add(mBase, uint32(v50))) = v91 + int64(1)
												v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v95)+20)) = int32(1)
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
												v99 = F_HeapTupleHeaderGetDatum(m, v98)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return int32(0)
												} else {
													v113 = v99
													m.G0 = v11 + int32(48)
													return v113
												}
											}
										}
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v103)+20)) = int32(2)
										v106 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v106)
										v113 = int32(0)
										m.G0 = v11 + int32(48)
										return v113
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
		v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
		if base.Ui64(v51) <= base.Ui64(int64(218)) {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
			v55 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v55)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v55
			v61 = base.I32_wrap_i64(v51) * int32(20)
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_pg_get_catalog_foreign_keys[1])))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v62
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_pg_get_catalog_foreign_keys[2])))
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_pg_get_catalog_foreign_keys[3])))
			v69 = F_FunctionCall3Coll(m, v54, v55, v66, int32(25), int32(-1))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v64
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v69
				v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_pg_get_catalog_foreign_keys[4]))))
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_pg_get_catalog_foreign_keys[5])))
				v78 = F_FunctionCall3Coll(m, v54, int32(0), v75, int32(25), int32(-1))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v73
					v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+uint32(_c_F_pg_get_catalog_foreign_keys[6]))))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v81
					*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v78
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
					v89 = F_heap_form_tuple(m, v84, v11+int32(16), v11+int32(8))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						v91 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
						*(*int64)(unsafe.Add(mBase, uint32(v50))) = v91 + int64(1)
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v95)+20)) = int32(1)
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
						v99 = F_HeapTupleHeaderGetDatum(m, v98)
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v113 = v99
							m.G0 = v11 + int32(48)
							return v113
						}
					}
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v103)+20)) = int32(2)
				v106 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v106)
				v113 = int32(0)
				m.G0 = v11 + int32(48)
				return v113
			}
		}
	}
}
func F_pg_get_indexdef_ext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v13 != 0 {
		v14 = int32(7)
	} else {
		v14 = int32(2)
	}
	v16 = F_pg_get_indexdef_worker(m, v3, v4, v2, base.B2i32(v4 != v2), v2, v2, v2, v14, int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			return int32(0)
		} else {
			v26 = F_cstring_to_text(m, v16)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v16)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					return v26
				}
			}
		}
	}
}
func F_pg_get_ruledef_ext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v6 != 0 {
		v7 = int32(7)
	} else {
		v7 = int32(2)
	}
	v8 = F_pg_get_ruledef_worker(m, v3, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 == int32(0) {
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
			return int32(0)
		} else {
			v18 = F_cstring_to_text(m, v8)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v8)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v18
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
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
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v566 int32
	_ = v566
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v669 int32
	_ = v669
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	v4 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(192)
	m.G0 = v24
	v27 = F_SearchSysCache1(m, int32(64), l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L4
	} else {
		goto L155
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L4
	} else {
		goto L152
	}
L3:
	;
	m.G0 = v24 + int32(192)
	return v669
L4:
	;
	return int32(0)
L5:
	;
	if v27 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if l2 != 0 {
		v669 = int32(0)
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v49 = F_heap_attisnull(m, v27, int32(9), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L13
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_statisticsobj_worker_0), v24)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_pg_get_statisticsobj_worker_1), int32(1680), int32(_a_F_pg_get_statisticsobj_worker_2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
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
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+22)))
	v53 = v51 + v52
	if v49 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	F_initStringInfo(m, v24+int32(120))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L25
	}
L15:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v79 = v63
	v80 = v68
	v81 = v69
	v82 = v75
	v83 = int32(0)
	goto L14
L16:
	;
	v79 = v4
	v80 = v72
	v81 = v73
	v82 = v4
	v83 = int32(1)
	goto L14
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	v72 = v53 + int32(96)
	v73 = v56
	goto L16
L18:
	;
	goto L19
L19:
	;
	v59 = F_SysCacheGetAttrNotNull(m, int32(64), v27, int32(9))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v61 = F_text_to_cstring(m, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v63 = F_stringToNode(m, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_pfree(m, v61)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v68 = v53 + int32(96)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	if v63 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v72 = v68
	v73 = v69
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
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v53)+72))
	v93 = F_get_namespace_name_or_temp(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v382 = int32(1)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v383 <= int32(0) {
		goto L104
	} else {
		goto L105
	}
L29:
	;
	v96 = v24 + int32(136)
	F_initStringInfo(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	if v93 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v99 = F_quote_identifier(m, v93)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v109 = F_quote_identifier(m, v53+int32(8))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+112)) = v99
	F_appendStringInfo(m, v96, int32(_a_F_pg_get_statisticsobj_worker_3), v24+int32(112))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	F_appendStringInfoString(m, v24+int32(136), v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v24)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v113
	F_appendStringInfo(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_4), v24+int32(96))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v124 = F_SysCacheGetAttrNotNull(m, int32(64), v27, int32(8))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v126 = F_pg_detoast_datum(m, v124)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v128 != int32(1) {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	if v131 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	if v132 != int32(18) {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	if v135 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v289&(v285&v288)|base.B2i32(v81+v82 < int32(2)) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L45:
	;
	v285 = int32(0)
	v288 = v4
	v289 = v4
	goto L44
L46:
	;
	goto L47
L47:
	;
	v140 = v126 + int32(24)
	v142 = v135 & int32(3)
	v143 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v135) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v149 = v143
	v153 = v143
	v156 = v4
	v157 = v4
	v168 = v4
	goto L51
L49:
	;
	v222 = v143
	v226 = v143
	v229 = v4
	v230 = v4
	goto L50
L50:
	;
	v243 = v222
	v245 = v226
	v250 = v229
	v251 = v230
	v263 = v4
	goto L71
L51:
	;
	v170 = int32(1)
	v172 = v149 + v140
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	switch v173 - int32(100) {
	case 0:
		v180 = v170
		v181 = v156
		v182 = v157
		goto L53
	default:
		v178 = v156
		v179 = v157
		goto L54
	case 2:
		goto L56
	case 9:
		goto L55
	}
L52:
	;
	if v142 == int32(0) {
		v285 = v212
		v288 = v213
		v289 = v214
		goto L44
	} else {
		goto L70
	}
L53:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
	switch v183 - int32(100) {
	case 0:
		v190 = v170
		v191 = v181
		v192 = v182
		goto L57
	default:
		v188 = v181
		v189 = v182
		goto L58
	case 2:
		goto L59
	case 9:
		goto L60
	}
L54:
	;
	v180 = v153
	v181 = v178
	v182 = v179
	goto L53
L55:
	;
	v178 = v156
	v179 = int32(1)
	goto L54
L56:
	;
	v178 = int32(1)
	v179 = v157
	goto L54
L57:
	;
	v193 = int32(1)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+2)))
	switch v195 - int32(100) {
	case 0:
		v202 = v193
		v203 = v191
		v204 = v192
		goto L61
	default:
		v200 = v191
		v201 = v192
		goto L62
	case 2:
		goto L63
	case 9:
		goto L64
	}
L58:
	;
	v190 = v180
	v191 = v188
	v192 = v189
	goto L57
L59:
	;
	v188 = int32(1)
	v189 = v182
	goto L58
L60:
	;
	v188 = v181
	v189 = int32(1)
	goto L58
L61:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+3)))
	switch v205 - int32(100) {
	case 0:
		v212 = v193
		v213 = v203
		v214 = v204
		goto L65
	default:
		v210 = v203
		v211 = v204
		goto L66
	case 2:
		goto L67
	case 9:
		goto L68
	}
L62:
	;
	v202 = v190
	v203 = v200
	v204 = v201
	goto L61
L63:
	;
	v200 = int32(1)
	v201 = v192
	goto L62
L64:
	;
	v200 = v191
	v201 = int32(1)
	goto L62
L65:
	;
	v215 = int32(4)
	v216 = v149 + v215
	v218 = v168 + v215
	if v218 != v135&int32(2147483644) {
		v149 = v216
		v153 = v212
		v156 = v213
		v157 = v214
		v168 = v218
		goto L51
	} else {
		goto L69
	}
L66:
	;
	v212 = v202
	v213 = v210
	v214 = v211
	goto L65
L67:
	;
	v210 = int32(1)
	v211 = v204
	goto L66
L68:
	;
	v210 = v203
	v211 = int32(1)
	goto L66
L69:
	;
	goto L52
L70:
	;
	v222 = v216
	v226 = v212
	v229 = v213
	v230 = v214
	goto L50
L71:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v140))))
	switch v266 - int32(100) {
	case 0:
		v273 = int32(1)
		v274 = v250
		v275 = v251
		goto L73
	default:
		v271 = v250
		v272 = v251
		goto L74
	case 2:
		goto L75
	case 9:
		goto L76
	}
L72:
	;
	v285 = v273
	v288 = v274
	v289 = v275
	goto L44
L73:
	;
	v276 = int32(1)
	v279 = v263 + v276
	if v279 != v142 {
		v243 = v243 + v276
		v245 = v273
		v250 = v274
		v251 = v275
		v263 = v279
		goto L71
	} else {
		goto L77
	}
L74:
	;
	v273 = v245
	v274 = v271
	v275 = v272
	goto L73
L75:
	;
	v271 = int32(1)
	v272 = v251
	goto L74
L76:
	;
	v271 = v250
	v272 = int32(1)
	goto L74
L77:
	;
	goto L72
L78:
	;
	v311 = v24 + int32(120)
	F_appendStringInfoString(m, v311, int32(_a_F_pg_get_statisticsobj_worker_5))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
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
	v360 = m.ExcPending
	if v360 != 0 {
		goto L4
	} else {
		goto L102
	}
L81:
	;
	if v285&int32(1) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_appendStringInfoString(m, v311, int32(_a_F_pg_get_statisticsobj_worker_7))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v288 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L84
L86:
	;
	if v289 != 0 {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	v336 = v285
	goto L86
L88:
	;
	goto L89
L89:
	;
	v322 = int32(1)
	if v285&v322 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v327 = int32(_a_F_pg_get_statisticsobj_worker_8)
	goto L92
L91:
	;
	v327 = int32(_a_F_pg_get_statisticsobj_worker_9)
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v327
	F_appendStringInfo(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_10), v24+int32(80))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v336 = v322
	goto L86
L94:
	;
	if v336&int32(1) != 0 {
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
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L101
	}
L97:
	;
	v341 = int32(_a_F_pg_get_statisticsobj_worker_8)
	goto L99
L98:
	;
	v341 = int32(_a_F_pg_get_statisticsobj_worker_9)
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v341
	F_appendStringInfo(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_11), v24-int32(-64))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
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
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v467 = F_get_rel_name(m, v466)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L4
	} else {
		goto L118
	}
L104:
	;
	v445 = int32(0)
	goto L103
L105:
	;
	goto L106
L106:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v390 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+104)))
	v392 = F_get_attname(m, v389, v390, int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v394 = F_quote_identifier(m, v392)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	F_appendStringInfoString(m, v24+int32(120), v394)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	if v398 < int32(2) {
		v445 = v382
		goto L103
	} else {
		goto L110
	}
L110:
	;
	v403 = v382
	goto L111
L111:
	;
	v427 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53+int32(104)+v403<<(uint(int32(1))%32)))))
	v429 = v24 + int32(120)
	F_appendStringInfoString(m, v429, int32(_a_F_pg_get_statisticsobj_worker_8))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L113
	}
L112:
	;
	v445 = v442
	goto L103
L113:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v435 = F_get_attname(m, v433, v427, int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	v437 = F_quote_identifier(m, v435)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	F_appendStringInfoString(m, v429, v437)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	v442 = v403 + int32(1)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	if v442 < v443 {
		v403 = v442
		goto L111
	} else {
		goto L117
	}
L117:
	;
	goto L112
L118:
	;
	if v467 == int32(0) {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v473 = F_palloc0(m, int32(80))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	v476 = F_palloc0(m, int32(136))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v476)+24)) = int32(1)
	v480 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v476)+21)) = uint8(v480)
	*(*int32)(unsafe.Add(mBase, uint32(v476)+16)) = v471
	v483 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v476)+12)) = v483
	*(*int32)(unsafe.Add(mBase, uint32(v476))) = int32(101)
	v489 = F_makeAlias(m, v467, v483)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v476)+8)) = v489
	*(*int32)(unsafe.Add(mBase, uint32(v476)+4)) = v489
	v493 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v476)+124)) = uint16(v493)
	v495 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v476)+20)) = uint8(v495)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = v476
	*(*int32)(unsafe.Add(mBase, uint32(v24)+136)) = v476
	v502 = F_list_make1_impl(m, int32(1), v24+int32(60))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v504 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v473)+20)) = v504
	*(*int64)(unsafe.Add(mBase, uint32(v473)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v473))) = v502
	F_set_rtable_names(m, v473, v504, v504)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	F_set_simple_column_names(m, v473)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v24)+176)) = v473
	v520 = F_list_make1_impl(m, int32(1), v24+int32(56))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	if v83 != 0 {
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
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v522 <= int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v525 = v445
	v527 = v483
	goto L130
L130:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v546+v527<<(uint(int32(2))%32))))
	v552 = v24 + int32(176)
	F_initStringInfo(m, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L132
	}
L131:
	;
	goto L127
L132:
	;
	v555 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+168)) = uint8(v555)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+152)) = v555
	*(*int64)(unsafe.Add(mBase, uint32(v24)+144)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+140)) = v520
	*(*int32)(unsafe.Add(mBase, uint32(v24)+172)) = v555
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+171)) = uint8(v555)
	v566 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+169)) = uint16(v566)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+164)) = v555
	*(*int64)(unsafe.Add(mBase, uint32(v24)+156)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+136)) = v552
	F_get_rule_expr(m, v550, v24+int32(136), v555)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v24)+176))
	if int32(0) < v525 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	F_appendStringInfoString(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_8))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L4
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	if v550 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	goto L136
L138:
	;
	v604 = int32(1)
	v607 = v527 + v604
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v607 < v608 {
		v525 = v525 + v604
		v527 = v607
		goto L130
	} else {
		goto L145
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v578
	F_appendStringInfo(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_12), v24+int32(48))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L144
	}
L140:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v550)))
	switch v588 - int32(15) {
	case 0:
		goto L142
	default:
		goto L139
	case 4, 23, 24, 25, 26, 33:
		goto L141
	}
L141:
	;
	F_appendStringInfoString(m, v24+int32(120), v578)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L4
	} else {
		goto L143
	}
L142:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v550)+16))
	switch v591 {
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
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v635 = F_generate_relation_name(m, v633, int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L4
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	F_ReleaseCatCache(m, v27)
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L4
	} else {
		goto L151
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v635
	F_appendStringInfo(m, v24+int32(120), int32(_a_F_pg_get_statisticsobj_worker_13), v24+int32(32))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v24)+120))
	v669 = v647
	goto L3
L152:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_statisticsobj_worker_14), int32(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(_a_F_pg_get_statisticsobj_worker_1), int32(1729), int32(_a_F_pg_get_statisticsobj_worker_2))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v466
	F_errmsg_internal(m, int32(_a_F_pg_get_statisticsobj_worker_15), v24+int32(16))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_pg_get_statisticsobj_worker_1), int32(_a_F_pg_get_statisticsobj_worker_16), int32(_a_F_pg_get_statisticsobj_worker_17))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
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
func F_pg_get_statisticsobjdef(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13984(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pg_get_statisticsobjdef_columns(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13984(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pg_get_userbyid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_palloc(m, int32(64))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v11))) = v15
		v32 = F_SearchSysCache1(m, int32(11), v9)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			if v32 != 0 {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
				v36 = v34 + v35
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)+60))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v37
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+52))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v36)+44))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v41
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v36)+36))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v43
				v45 = *(*int64)(unsafe.Add(mBase, uint32(v36)+28))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v45
				v47 = *(*int64)(unsafe.Add(mBase, uint32(v36)+20))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v47
				v49 = *(*int64)(unsafe.Add(mBase, uint32(v36)+12))
				*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v49
				v51 = *(*int64)(unsafe.Add(mBase, uint32(v36)+4))
				*(*int64)(unsafe.Add(mBase, uint32(v11))) = v51
				F_ReleaseCatCache(m, v32)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v11
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
				v57 = F_pg_sprintf(m, v11, int32(_a_F_pg_get_userbyid_0), v7)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v11
				}
			}
		}
	}
}
func F_pg_hba_file_rules(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
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
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(736)
	m.G0 = v21
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+348)) = v31
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_hba_file_rules[0]))
	v38 = F_open_auth_file(m, v34, int32(21), v31, v31)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_pg_hba_file_rules[0]))
	F_tokenize_auth_file(m, v41, v38, v21+int32(348), int32(12), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_pg_hba_file_rules[1]))
	v54 = F_AllocSetContextCreateInternal(m, v49, int32(_a_F_pg_hba_file_rules_0), int32(0), int32(1024), int32(_a_F_pg_hba_file_rules_1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v56 = int32(_a_F_pg_hba_file_rules_2)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_pg_hba_file_rules[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_hba_file_rules[1])) = v54
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v21)+348))
	if v60 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_free_auth_file(m, v38)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L240
	}
L7:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v63 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v71 = v21 + int32(611)
	v83 = v2
	v84 = v2
	goto L9
L9:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v83<<(uint(int32(2))%32))))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	if v96 == v90 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v100 = F_parse_hba_line(m, v95, int32(12))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v103 = v90
	v104 = v96
	goto L13
L13:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v107 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+624)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v21)+632)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v21)+640)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v21)+648)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v21)+656)) = v107
	v117 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+664)) = v117
	*(*int64)(unsafe.Add(mBase, uint32(v21)+608)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v21)+615)) = v117
	v124 = v84 + int32(1)
	if v104 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	v103 = v100
	v104 = v102
	goto L13
L15:
	;
	v128 = F_cstring_to_text(m, v106)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	v125 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+608)) = uint8(v125)
	goto L15
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+624)) = v124
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+632)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+628)) = v128
	if v103 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v104 != 0 {
		goto L230
	} else {
		goto L231
	}
L21:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	if base.Ui32(v132) <= base.Ui32(int32(5)) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v874 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+3)) = v874
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v874
	goto L20
L24:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	if v143 != 0 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132<<(uint(int32(2))%32))+uint32(_c_F_pg_hba_file_rules[2])))
	v138 = F_cstring_to_text(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v141 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+611)) = uint8(v141)
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+636)) = v138
	goto L24
L29:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
	if v221 != 0 {
		goto L43
	} else {
		goto L44
	}
L30:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v144 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L32
L32:
	;
	v201 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+612)) = uint8(v201)
	goto L29
L33:
	;
	v198 = F_strlist_to_textarray(m, v183)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L41
	}
L34:
	;
	v183 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v148 = int32(0)
	v151 = v148
	v153 = v148
	goto L37
L37:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168+v151<<(uint(int32(2))%32))))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v174 = F_lappend(m, v153, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v183 = v174
	goto L33
L39:
	;
	v177 = v151 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if v177 < v178 {
		v151 = v177
		v153 = v174
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+640)) = v198
	goto L29
L42:
	;
	v300 = int32(0)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v103)+288))
	switch v301 {
	case 0:
		goto L60
	case 1:
		goto L58
	case 2:
		goto L57
	case 3:
		v373 = v300
		v374 = int32(_a_F_pg_hba_file_rules_3)
		goto L56
	default:
		v366 = v300
		goto L59
	}
L43:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v222 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	goto L45
L45:
	;
	v279 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+613)) = uint8(v279)
	goto L42
L46:
	;
	v276 = F_strlist_to_textarray(m, v261)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L54
	}
L47:
	;
	v261 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v226 = int32(0)
	v229 = v226
	v231 = v226
	goto L50
L50:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v246+v229<<(uint(int32(2))%32))))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v252 = F_lappend(m, v231, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	v261 = v252
	goto L46
L52:
	;
	v255 = v229 + int32(1)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v255 < v256 {
		v229 = v255
		v231 = v252
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+644)) = v276
	goto L42
L55:
	;
	if v379 != 0 {
		goto L89
	} else {
		goto L90
	}
L56:
	;
	v376 = F_cstring_to_text(m, v374)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L87
	}
L57:
	;
	v373 = v300
	v374 = int32(_a_F_pg_hba_file_rules_4)
	goto L56
L58:
	;
	v373 = v300
	v374 = int32(_a_F_pg_hba_file_rules_5)
	goto L56
L59:
	;
	v369 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+614)) = uint8(v369)
	v379 = v366
	goto L55
L60:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v103)+292))
	if v302 != 0 {
		v373 = v300
		v374 = v302
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v303 = int32(0)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v103)+152))
	if v303 < v304 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v308 = v103 + int32(24)
	v310 = v21 + int32(352)
	v312 = int32(0)
	v315 = F_pg_getnameinfo_all(m, v308, v304, v310, int32(255), v312, v312, int32(1))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	v333 = v303
	goto L64
L64:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v103)+284))
	if int32(0) < v335 {
		goto L74
	} else {
		goto L75
	}
L65:
	;
	if v315 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v308))))
	if v319 != int32(10) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L68
L68:
	;
	v331 = F_pstrdup(m, v21+int32(352))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
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
	v323 = F_strchr(m, v310, int32(37))
	mBase = m.M
	if v323 == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v326 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v326)
	goto L70
L73:
	;
	v333 = v331
	goto L64
L74:
	;
	v339 = v103 + int32(156)
	v341 = v21 + int32(352)
	v343 = int32(0)
	v346 = F_pg_getnameinfo_all(m, v339, v335, v341, int32(255), v343, v343, int32(1))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v364 = v300
	goto L76
L76:
	;
	if v333 != 0 {
		v373 = v364
		v374 = v333
		goto L56
	} else {
		goto L86
	}
L77:
	;
	if v346 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v339))))
	if v350 != int32(10) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L80
L80:
	;
	v362 = F_pstrdup(m, v21+int32(352))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
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
	v354 = F_strchr(m, v341, int32(37))
	mBase = m.M
	if v354 == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v357 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v354))) = uint8(v357)
	goto L82
L85:
	;
	v364 = v362
	goto L76
L86:
	;
	v366 = v364
	goto L59
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+648)) = v376
	v379 = v373
	goto L55
L88:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v103)+296))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v387<<(uint(int32(2))%32))+uint32(_c_F_pg_hba_file_rules[3])))
	goto L93
L89:
	;
	v382 = F_cstring_to_text(m, v379)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v385 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+615)) = uint8(v385)
	goto L88
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+652)) = v382
	goto L88
L93:
	;
	v391 = F_cstring_to_text(m, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+656)) = v391
	v394 = int32(0)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v103)+296))
	if base.Ui32(int32(1)) < base.Ui32(v395-int32(7)) {
		v426 = v394
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v103)+300))
	if v428 != 0 {
		goto L104
	} else {
		goto L105
	}
L96:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+368)))
	if v400 != int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v410 = v394
	v411 = v21 + int32(672)
	goto L99
L98:
	;
	v406 = F_cstring_to_text(m, int32(_a_F_pg_hba_file_rules_6))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L100
	}
L99:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v103)+364))
	if v412 == int32(0) {
		v426 = v410
		goto L95
	} else {
		goto L101
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+672)) = v406
	v410 = int32(1)
	v411 = v21 + int32(672) | int32(4)
	goto L99
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+336)) = v412
	v419 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_7), v21+int32(336))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v421 = F_cstring_to_text(m, v419)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v411))) = v421
	v426 = v410 + int32(1)
	goto L95
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+320)) = v428
	v438 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_8), v21+int32(320))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	v445 = v426
	goto L106
L106:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v103)+356))
	if v446 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v440 = F_cstring_to_text(m, v438)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)|v426<<(uint(int32(2))%32)))) = v440
	v445 = v426 + int32(1)
	goto L106
L109:
	;
	if v446 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v468 = v445
	goto L111
L111:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v103)+304))
	if v469 != 0 {
		goto L117
	} else {
		goto L118
	}
L112:
	;
	v451 = int32(_a_F_pg_hba_file_rules_9)
	goto L114
L113:
	;
	v451 = int32(_a_F_pg_hba_file_rules_10)
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+304)) = v451
	v461 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_11), v21+int32(304))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v463 = F_cstring_to_text(m, v461)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v445<<(uint(int32(2))%32)))) = v463
	v468 = v445 + int32(1)
	goto L111
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+288)) = v469
	v479 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_12), v21+int32(288))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	v486 = v468
	goto L119
L119:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v103)+296))
	if v487 == int32(11) {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v481 = F_cstring_to_text(m, v479)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v468<<(uint(int32(2))%32)))) = v481
	v486 = v468 + int32(1)
	goto L119
L122:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v103)+296))
	if v784 != int32(15) {
		v856 = v782
		goto L206
	} else {
		goto L207
	}
L123:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v103)+316))
	if v490 != 0 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v704 = v486
	v706 = v487
	goto L125
L125:
	;
	if v706 != int32(13) {
		v782 = v704
		goto L122
	} else {
		goto L185
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+272)) = v490
	v500 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_13), v21+int32(272))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	v507 = v486
	goto L128
L128:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v103)+320))
	if v508 != 0 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	v502 = F_cstring_to_text(m, v500)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v486<<(uint(int32(2))%32)))) = v502
	v507 = v486 + int32(1)
	goto L128
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+256)) = v508
	v518 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_14), v21+int32(256))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L134
	}
L132:
	;
	v525 = v507
	goto L133
L133:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v103)+312))
	if v526 != 0 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v520 = F_cstring_to_text(m, v518)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v507<<(uint(int32(2))%32)))) = v520
	v525 = v507 + int32(1)
	goto L133
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+240)) = v526
	v536 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_15), v21+int32(240))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	v543 = v525
	goto L138
L138:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+309)))
	if v544 == int32(1) {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v538 = F_cstring_to_text(m, v536)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v525<<(uint(int32(2))%32)))) = v538
	v543 = v525 + int32(1)
	goto L138
L141:
	;
	v553 = F_cstring_to_text(m, int32(_a_F_pg_hba_file_rules_16))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	v558 = v543
	goto L143
L143:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v103)+348))
	if v559 != 0 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v543<<(uint(int32(2))%32)))) = v553
	v558 = v543 + int32(1)
	goto L143
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+224)) = v559
	v569 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_17), v21+int32(224))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	v576 = v558
	goto L147
L147:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v103)+352))
	if v577 != 0 {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v571 = F_cstring_to_text(m, v569)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v558<<(uint(int32(2))%32)))) = v571
	v576 = v558 + int32(1)
	goto L147
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+208)) = v577
	v587 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_18), v21+int32(208))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L153
	}
L151:
	;
	v594 = v576
	goto L152
L152:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v103)+340))
	if v595 != 0 {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	v589 = F_cstring_to_text(m, v587)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v576<<(uint(int32(2))%32)))) = v589
	v594 = v576 + int32(1)
	goto L152
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+192)) = v595
	v605 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_19), v21+int32(192))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	v612 = v594
	goto L157
L157:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v103)+324))
	if v613 != 0 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v607 = F_cstring_to_text(m, v605)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v594<<(uint(int32(2))%32)))) = v607
	v612 = v594 + int32(1)
	goto L157
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = v613
	v623 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_20), v21+int32(176))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L163
	}
L161:
	;
	v630 = v612
	goto L162
L162:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v103)+328))
	if v631 != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v625 = F_cstring_to_text(m, v623)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v612<<(uint(int32(2))%32)))) = v625
	v630 = v612 + int32(1)
	goto L162
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+160)) = v631
	v641 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_21), v21+int32(160))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	v648 = v630
	goto L167
L167:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v103)+332))
	if v649 != 0 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v643 = F_cstring_to_text(m, v641)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v630<<(uint(int32(2))%32)))) = v643
	v648 = v630 + int32(1)
	goto L167
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+144)) = v649
	v659 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_22), v21+int32(144))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	v666 = v648
	goto L172
L172:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v103)+336))
	if v667 != 0 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	v661 = F_cstring_to_text(m, v659)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v648<<(uint(int32(2))%32)))) = v661
	v666 = v648 + int32(1)
	goto L172
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v667
	v677 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_23), v21+int32(128))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	v684 = v666
	goto L177
L177:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v103)+344))
	if v685 != 0 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v679 = F_cstring_to_text(m, v677)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v666<<(uint(int32(2))%32)))) = v679
	v684 = v666 + int32(1)
	goto L177
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v685
	v695 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_24), v21+int32(112))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	v702 = v684
	goto L182
L182:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v103)+296))
	v704 = v702
	v706 = v703
	goto L125
L183:
	;
	v697 = F_cstring_to_text(m, v695)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v684<<(uint(int32(2))%32)))) = v697
	v702 = v684 + int32(1)
	goto L182
L185:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v103)+376))
	if v709 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v709
	v719 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_25), v21+int32(96))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L189
	}
L187:
	;
	v726 = v704
	goto L188
L188:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v103)+384))
	if v727 != 0 {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	v721 = F_cstring_to_text(m, v719)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v704<<(uint(int32(2))%32)))) = v721
	v726 = v704 + int32(1)
	goto L188
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v727
	v737 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_26), v21+int32(80))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L194
	}
L192:
	;
	v744 = v726
	goto L193
L193:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v103)+392))
	if v745 != 0 {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	v739 = F_cstring_to_text(m, v737)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v726<<(uint(int32(2))%32)))) = v739
	v744 = v726 + int32(1)
	goto L193
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v745
	v755 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_27), v21-int32(-64))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L199
	}
L197:
	;
	v762 = v744
	goto L198
L198:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v103)+400))
	if v763 == int32(0) {
		v782 = v762
		goto L122
	} else {
		goto L201
	}
L199:
	;
	v757 = F_cstring_to_text(m, v755)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v744<<(uint(int32(2))%32)))) = v757
	v762 = v744 + int32(1)
	goto L198
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v763
	v775 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_28), v21+int32(48))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v777 = F_cstring_to_text(m, v775)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v762<<(uint(int32(2))%32)))) = v777
	v782 = v762 + int32(1)
	goto L122
L204:
	;
	v872 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+617)) = uint8(v872)
	goto L20
L205:
	;
	v865 = F_construct_array_builtin(m, v21+int32(672), v860, int32(25))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L1
	} else {
		goto L227
	}
L206:
	;
	if v856 == int32(0) {
		goto L204
	} else {
		goto L226
	}
L207:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v103)+404))
	if v787 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v787
	v797 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_29), v21+int32(32))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L211
	}
L209:
	;
	v804 = v782
	goto L210
L210:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v103)+408))
	if v805 != 0 {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v799 = F_cstring_to_text(m, v797)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v782<<(uint(int32(2))%32)))) = v799
	v804 = v782 + int32(1)
	goto L210
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v805
	v815 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_30), v21+int32(16))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L216
	}
L214:
	;
	v822 = v804
	goto L215
L215:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v103)+412))
	if v823 != 0 {
		goto L218
	} else {
		goto L219
	}
L216:
	;
	v817 = F_cstring_to_text(m, v815)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v804<<(uint(int32(2))%32)))) = v817
	v822 = v804 + int32(1)
	goto L215
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v823
	v831 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_31), v21)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L221
	}
L219:
	;
	v838 = v822
	goto L220
L220:
	;
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+416)))
	if v839 != int32(1) {
		v856 = v838
		goto L206
	} else {
		goto L223
	}
L221:
	;
	v833 = F_cstring_to_text(m, v831)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v822<<(uint(int32(2))%32)))) = v833
	v838 = v822 + int32(1)
	goto L220
L223:
	;
	v849 = F_psprintf(m, int32(_a_F_pg_hba_file_rules_32), int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v851 = F_cstring_to_text(m, v849)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(672)+v838<<(uint(int32(2))%32)))) = v851
	v860 = v838 + int32(1)
	goto L205
L226:
	;
	v860 = v856
	goto L205
L227:
	;
	if v865 == int32(0) {
		goto L204
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+660)) = v865
	goto L20
L229:
	;
	if v104 != 0 {
		goto L234
	} else {
		goto L235
	}
L230:
	;
	v896 = F_cstring_to_text(m, v104)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L1
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v899 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+618)) = uint8(v899)
	goto L229
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+664)) = v896
	goto L229
L234:
	;
	v901 = v84
	goto L236
L235:
	;
	v901 = v124
	goto L236
L236:
	;
	v906 = F_heap_form_tuple(m, v29, v21+int32(624), v21+int32(608))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_tuplestore_puttuple(m, v30, v906)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	v911 = v83 + int32(1)
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v911 < v912 {
		v83 = v911
		v84 = v901
		goto L9
	} else {
		goto L239
	}
L239:
	;
	goto L10
L240:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_hba_file_rules[1])) = v57
	F_MemoryContextDelete(m, v54)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v938 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v938)
	m.G0 = v21 + int32(736)
	return int32(0)
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
func F_pg_hypot(m *base.Module, l0 float64, l1 float64) float64 {
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	var v9 float64
	_ = v9
	var v14 int64
	_ = v14
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v25 float64
	_ = v25
	var v30 float64
	_ = v30
	var v38 float64
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v5 = math.Float64frombits(uint64(0x7ff0000000000000))
	v6 = base.F64_abs(l0)
	if base.F64_eq(v6, v5) != 0 {
		v38 = v5
		return v38
	} else {
		v9 = base.F64_abs(l1)
		if base.F64_eq(v9, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v38 = v5
			return v38
		} else {
			v14 = int64(9218868437227405312)
			if base.B2i32(base.Ui64(v14) < base.Ui64(base.I64_reinterpret_f64(v6)))|base.B2i32(base.Ui64(v14) < base.Ui64(base.I64_reinterpret_f64(v9))) != 0 {
				v38 = math.Float64frombits(uint64(0x7ff8000000000000))
				return v38
			} else {
				v20 = base.F64_lt(v6, v9)
				if v20 != 0 {
					v21 = v9
				} else {
					v21 = v6
				}
				if v20 != 0 {
					v22 = v6
				} else {
					v22 = v9
				}
				if base.F64_eq(v22, float64(0)) != 0 {
					v38 = v21
					return v38
				} else {
					v25 = base.F64_div(v22, v21)
					v30 = base.F64_mul(v21, base.F64_sqrt(base.F64_add(base.F64_mul(v25, v25), float64(1))))
					if base.F64_eq(base.F64_abs(v30), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						v44 = m.ExcPending
						if v44 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if base.F64_eq(v30, float64(0)) != 0 {
							F_float_underflow_error(m)
							v46 = m.ExcPending
							if v46 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v38 = v30
							return v38
						}
					}
				}
			}
		}
	}
}
func F_pg_input_error_info(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_input_error_info[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v18
			v21 = *(*int64)(unsafe.Add(mBase, _c_F_pg_input_error_info[1]))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+48)) = v21
			v26 = F_get_call_result_type(m, l0, int32(0), v5+int32(-20))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(1) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+53)) = uint8(v30)
					v34 = F_pg_input_is_valid_common(m, l0, v10, v15, v5+int32(-16))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						if v34 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(16843009)
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
							v115 = F_heap_form_tuple(m, v110, v5+int32(-48), v5+int32(-52))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
								v118 = F_HeapTupleHeaderGetDatum(m, v117)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 - int32(-64)
									return v118
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
							v42 = F_cstring_to_text(m, v41)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v42
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+36))
								if v46 != 0 {
									v47 = F_cstring_to_text(m, v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v47
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
										v53 = v50
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
										if v54 != 0 {
											v55 = F_cstring_to_text(m, v54)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v55
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
												v61 = v58
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
												v63 = int32(_a_F_pg_input_error_info_0)
												v64 = int32(63)
												v66 = int32(48)
												v67 = v62&v64 + v66
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[2])) = uint8(v67)
												v75 = int32(base.Ui32(v62)>>(uint(int32(24))%32))&v64 + v66
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[3])) = uint8(v75)
												v83 = int32(base.Ui32(v62)>>(uint(int32(18))%32))&v64 + v66
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[4])) = uint8(v83)
												v91 = int32(base.Ui32(v62)>>(uint(int32(12))%32))&v64 + v66
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[5])) = uint8(v91)
												v99 = int32(base.Ui32(v62)>>(uint(int32(6))%32))&v64 + v66
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[6])) = uint8(v99)
												v102 = int32(0)
												*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[7])) = uint8(v102)
												v105 = F_cstring_to_text(m, v63)
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v105
													v110 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
													v115 = F_heap_form_tuple(m, v110, v5+int32(-48), v5+int32(-52))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return int32(0)
													} else {
														v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
														v118 = F_HeapTupleHeaderGetDatum(m, v117)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return int32(0)
														} else {
															m.G0 = v7 - int32(-64)
															return v118
														}
													}
												}
											}
										} else {
											v59 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v59)
											v61 = v53
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
											v63 = int32(_a_F_pg_input_error_info_0)
											v64 = int32(63)
											v66 = int32(48)
											v67 = v62&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[2])) = uint8(v67)
											v75 = int32(base.Ui32(v62)>>(uint(int32(24))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[3])) = uint8(v75)
											v83 = int32(base.Ui32(v62)>>(uint(int32(18))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[4])) = uint8(v83)
											v91 = int32(base.Ui32(v62)>>(uint(int32(12))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[5])) = uint8(v91)
											v99 = int32(base.Ui32(v62)>>(uint(int32(6))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[6])) = uint8(v99)
											v102 = int32(0)
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[7])) = uint8(v102)
											v105 = F_cstring_to_text(m, v63)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v105
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
												v115 = F_heap_form_tuple(m, v110, v5+int32(-48), v5+int32(-52))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
													v118 = F_HeapTupleHeaderGetDatum(m, v117)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 - int32(-64)
														return v118
													}
												}
											}
										}
									}
								} else {
									v51 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v51)
									v53 = v45
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
									if v54 != 0 {
										v55 = F_cstring_to_text(m, v54)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v55
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v7)+56))
											v61 = v58
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
											v63 = int32(_a_F_pg_input_error_info_0)
											v64 = int32(63)
											v66 = int32(48)
											v67 = v62&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[2])) = uint8(v67)
											v75 = int32(base.Ui32(v62)>>(uint(int32(24))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[3])) = uint8(v75)
											v83 = int32(base.Ui32(v62)>>(uint(int32(18))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[4])) = uint8(v83)
											v91 = int32(base.Ui32(v62)>>(uint(int32(12))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[5])) = uint8(v91)
											v99 = int32(base.Ui32(v62)>>(uint(int32(6))%32))&v64 + v66
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[6])) = uint8(v99)
											v102 = int32(0)
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[7])) = uint8(v102)
											v105 = F_cstring_to_text(m, v63)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v105
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
												v115 = F_heap_form_tuple(m, v110, v5+int32(-48), v5+int32(-52))
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
													v118 = F_HeapTupleHeaderGetDatum(m, v117)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int32(0)
													} else {
														m.G0 = v7 - int32(-64)
														return v118
													}
												}
											}
										}
									} else {
										v59 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v59)
										v61 = v53
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
										v63 = int32(_a_F_pg_input_error_info_0)
										v64 = int32(63)
										v66 = int32(48)
										v67 = v62&v64 + v66
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[2])) = uint8(v67)
										v75 = int32(base.Ui32(v62)>>(uint(int32(24))%32))&v64 + v66
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[3])) = uint8(v75)
										v83 = int32(base.Ui32(v62)>>(uint(int32(18))%32))&v64 + v66
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[4])) = uint8(v83)
										v91 = int32(base.Ui32(v62)>>(uint(int32(12))%32))&v64 + v66
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[5])) = uint8(v91)
										v99 = int32(base.Ui32(v62)>>(uint(int32(6))%32))&v64 + v66
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[6])) = uint8(v99)
										v102 = int32(0)
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_input_error_info[7])) = uint8(v102)
										v105 = F_cstring_to_text(m, v63)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v105
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
											v115 = F_heap_form_tuple(m, v110, v5+int32(-48), v5+int32(-52))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int32(0)
											} else {
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
												v118 = F_HeapTupleHeaderGetDatum(m, v117)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int32(0)
												} else {
													m.G0 = v7 - int32(-64)
													return v118
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
					v127 = m.ExcPending
					if v127 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pg_input_error_info_1), int32(0))
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_input_error_info_2), int32(726), int32(_a_F_pg_input_error_info_3))
							mBase = m.M
							v136 = m.ExcPending
							if v136 != 0 {
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
						v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
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
							v90 = F_parseTypeString(m, v83, v10+int32(8), v80+int32(4), int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
								if v92 == v93 {
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
									v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
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
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
											v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
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
						v90 = F_parseTypeString(m, v83, v10+int32(8), v80+int32(4), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							if v92 == v93 {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
								v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
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
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
										v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
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
					v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
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
						v90 = F_parseTypeString(m, v83, v10+int32(8), v80+int32(4), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
							if v92 == v93 {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
								v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
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
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
										v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
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
					v90 = F_parseTypeString(m, v83, v10+int32(8), v80+int32(4), int32(0))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
						if v92 == v93 {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
							v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
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
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
									v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
									v117 = F_InputFunctionCallSafe(m, v80+int32(20), v12, v113, v114, l3, v10+int32(12))
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
func F_pg_isblank(m *base.Module, l0 int32) int32 {
	return base.I32_wrap_i64(int64(base.Ui64(int64(4294976000))>>(uint(base.I64_extend_i32_s(l0))%64))) & base.B2i32(base.Ui32(l0) < base.Ui32(int32(33)))
}
func F_pg_listening_channels(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v6 == int32(0) {
		v9 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_pg_listening_channels[0]))
			if v16 == int32(0) {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = int32(2)
					v43 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
					return int32(0)
				}
			} else {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
				v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+4)))
				if base.Ui64(v20) <= base.Ui64(v19) {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = int32(2)
						v43 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
						return int32(0)
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22+base.I32_wrap_i64(v19)<<(uint(int32(2))%32))))
					*(*int64)(unsafe.Add(mBase, uint32(v14))) = v19 + int64(1)
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1)
					v34 = F_cstring_to_text(m, v27)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						return v34
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_pg_listening_channels[0]))
		if v16 == int32(0) {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = int32(2)
				v43 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
				return int32(0)
			}
		} else {
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
			v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+4)))
			if base.Ui64(v20) <= base.Ui64(v19) {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = int32(2)
					v43 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
					return int32(0)
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v22+base.I32_wrap_i64(v19)<<(uint(int32(2))%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v14))) = v19 + int64(1)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1)
				v34 = F_cstring_to_text(m, v27)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					return v34
				}
			}
		}
	}
}
func F_pg_lock_status(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v224 int32
	_ = v224
	var v228 int64
	_ = v228
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	var v258 int64
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int64
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int64
	_ = v669
	var v671 int64
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int64
	_ = v693
	var v696 int64
	_ = v696
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int64
	_ = v1062
	var v1064 int64
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1250 int64
	_ = v1250
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1325 int32
	_ = v1325
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1448 int64
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1480 int64
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int64
	_ = v1500
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1534 int64
	_ = v1534
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1597 int64
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1620 int32
	_ = v1620
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int64
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(208)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v24 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+16))
	goto L167
L4:
	;
	return int32(0)
L5:
	;
	v31 = int32(_a_F_pg_lock_status_0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[0]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[0])) = v34
	v37 = F_CreateTemplateTupleDesc(m, int32(16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v37, int32(1), int32(_a_F_pg_lock_status_1), int32(25), int32(-1), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v37, int32(2), int32(_a_F_pg_lock_status_2), int32(26), int32(-1), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v37, int32(3), int32(_a_F_pg_lock_status_3), int32(26), int32(-1), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v37, int32(4), int32(_a_F_pg_lock_status_4), int32(23), int32(-1), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v37, int32(5), int32(_a_F_pg_lock_status_5), int32(21), int32(-1), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	F_TupleDescInitEntry(m, v37, int32(6), int32(_a_F_pg_lock_status_6), int32(25), int32(-1), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_TupleDescInitEntry(m, v37, int32(7), int32(_a_F_pg_lock_status_7), int32(28), int32(-1), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_TupleDescInitEntry(m, v37, int32(8), int32(_a_F_pg_lock_status_8), int32(26), int32(-1), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_TupleDescInitEntry(m, v37, int32(9), int32(_a_F_pg_lock_status_9), int32(26), int32(-1), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	F_TupleDescInitEntry(m, v37, int32(10), int32(_a_F_pg_lock_status_10), int32(21), int32(-1), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	F_TupleDescInitEntry(m, v37, int32(11), int32(_a_F_pg_lock_status_11), int32(25), int32(-1), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F_TupleDescInitEntry(m, v37, int32(12), int32(_a_F_pg_lock_status_12), int32(23), int32(-1), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_TupleDescInitEntry(m, v37, int32(13), int32(_a_F_pg_lock_status_13), int32(25), int32(-1), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_TupleDescInitEntry(m, v37, int32(14), int32(_a_F_pg_lock_status_14), int32(16), int32(-1), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	F_TupleDescInitEntry(m, v37, int32(15), int32(_a_F_pg_lock_status_15), int32(16), int32(-1), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	F_TupleDescInitEntry(m, v37, int32(16), int32(_a_F_pg_lock_status_16), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v151 = F_BlessTupleDesc(m, v37)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v151
	v155 = F_palloc(m, int32(16))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v155
	v158 = m.G0
	v160 = v158 - int32(32)
	m.G0 = v160
	v163 = F_palloc(m, int32(8))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[1]))
	v169 = F_palloc(m, v166*int32(56))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v169
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[2]))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
	if v174 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v176 = v173
	v178 = v2
	v181 = v166
	v186 = v2
	goto L30
L28:
	;
	v432 = v2
	v435 = v166
	goto L29
L29:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v452 = F_LWLockAcquire(m, v448+int32(_a_F_pg_lock_status_17), int32(1))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L4
	} else {
		goto L64
	}
L30:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v196 = v193 + v186*int32(640)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+44))
	if v197 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v432 = v410
	v435 = v413
	goto L29
L32:
	;
	v199 = v196 + int32(584)
	v201 = F_LWLockAcquire(m, v199, int32(1))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	v408 = v176
	v410 = v178
	v413 = v181
	goto L34
L34:
	;
	v426 = v186 + int32(1)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v408)+16))
	if base.Ui32(v426) < base.Ui32(v427) {
		v176 = v408
		v178 = v410
		v181 = v413
		v186 = v426
		goto L30
	} else {
		goto L63
	}
L35:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[4]))
	if v204 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v207 = v204
	v209 = v178
	v212 = v181
	v213 = int32(0)
	goto L39
L37:
	;
	v345 = v178
	v348 = v181
	goto L38
L38:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+608)))
	if v360 != 0 {
		goto L55
	} else {
		goto L56
	}
L39:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v196)+600))
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v224+v213<<(uint(int32(3))%32))))
	if v228 != int64(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v345 = v324
	v348 = v327
	goto L38
L41:
	;
	v241 = v209
	v244 = v212
	v255 = int64(0)
	goto L44
L42:
	;
	v322 = v207
	v324 = v209
	v327 = v212
	goto L43
L43:
	;
	v340 = v213 + int32(1)
	if base.Ui32(v340) < base.Ui32(v322) {
		v207 = v322
		v209 = v324
		v212 = v327
		v213 = v340
		goto L39
	} else {
		goto L54
	}
L44:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v196)+600))
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v256+v213&int32(268435455)<<(uint(int32(3))%32))))
	v264 = base.I32_wrap_i64(int64(base.Ui64(v258)>>(uint(v255*int64(3))%64))) & int32(7)
	if v264 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[4]))
	v322 = v320
	v324 = v311
	v327 = v312
	goto L43
L46:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v244 <= v241 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v311 = v241
	v312 = v244
	goto L48
L48:
	;
	v316 = v255 + int64(1)
	if v316 != int64(16) {
		v241 = v311
		v244 = v312
		v255 = v316
		goto L44
	} else {
		goto L53
	}
L49:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[1]))
	v269 = v268 + v244
	v272 = F_repalloc(m, v265, v269*int32(56))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	v275 = v265
	v276 = v244
	goto L51
L51:
	;
	v279 = v275 + v241*int32(56)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v196)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = v280
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v196)+604))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282+v213<<(uint(int32(6))%32)+base.I32_wrap_i64(v255)<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = int32(0)
	v291 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v279)+16)) = v264 << (uint(v291) % 32)
	*(*int64)(unsafe.Add(mBase, uint32(v279)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v279)+4)) = v288
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v196)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+24)) = v297
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v196)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+28)) = v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v196)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v279)+48)) = uint8(v291)
	*(*int32)(unsafe.Add(mBase, uint32(v279)+44)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v279)+40)) = v301
	*(*int64)(unsafe.Add(mBase, uint32(v279)+32)) = int64(0)
	v311 = v241 + v291
	v312 = v276
	goto L48
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v272
	v275 = v272
	v276 = v269
	goto L51
L53:
	;
	goto L45
L54:
	;
	goto L40
L55:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v348 <= v345 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v398 = v345
	v400 = v348
	goto L57
L57:
	;
	F_LWLockRelease(m, v199)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L62
	}
L58:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[1]))
	v365 = v364 + v348
	v368 = F_repalloc(m, v361, v365*int32(56))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L61
	}
L59:
	;
	v371 = v361
	v372 = v348
	goto L60
L60:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v196)+52))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v196)+612))
	v377 = v371 + v345*int32(56)
	*(*int64)(unsafe.Add(mBase, uint32(v377)+16)) = int64(128)
	*(*int64)(unsafe.Add(mBase, uint32(v377)+8)) = int64(73746443898191872)
	*(*int32)(unsafe.Add(mBase, uint32(v377)+4)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v377))) = v373
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v196)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v377)+24)) = v384
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v196)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v377)+28)) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v196)+44))
	v389 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v377)+48)) = uint8(v389)
	*(*int32)(unsafe.Add(mBase, uint32(v377)+44)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v377)+40)) = v388
	*(*int64)(unsafe.Add(mBase, uint32(v377)+32)) = int64(0)
	v398 = v345 + v389
	v400 = v372
	goto L57
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v368
	v371 = v368
	v372 = v365
	goto L60
L62:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[2]))
	v408 = v406
	v410 = v398
	v413 = v400
	goto L34
L63:
	;
	goto L31
L64:
	;
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v459 = F_LWLockAcquire(m, v455+int32(_a_F_pg_lock_status_18), int32(1))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v462 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v466 = F_LWLockAcquire(m, v462+int32(_a_F_pg_lock_status_19), int32(1))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v473 = F_LWLockAcquire(m, v469+int32(_a_F_pg_lock_status_20), int32(1))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v480 = F_LWLockAcquire(m, v476+int32(_a_F_pg_lock_status_21), int32(1))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v487 = F_LWLockAcquire(m, v483+int32(_a_F_pg_lock_status_22), int32(1))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v494 = F_LWLockAcquire(m, v490+int32(_a_F_pg_lock_status_23), int32(1))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v501 = F_LWLockAcquire(m, v497+int32(_a_F_pg_lock_status_24), int32(1))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v508 = F_LWLockAcquire(m, v504+int32(_a_F_pg_lock_status_25), int32(1))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v515 = F_LWLockAcquire(m, v511+int32(_a_F_pg_lock_status_26), int32(1))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v522 = F_LWLockAcquire(m, v518+int32(_a_F_pg_lock_status_27), int32(1))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v529 = F_LWLockAcquire(m, v525+int32(_a_F_pg_lock_status_28), int32(1))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v536 = F_LWLockAcquire(m, v532+int32(_a_F_pg_lock_status_29), int32(1))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v543 = F_LWLockAcquire(m, v539+int32(_a_F_pg_lock_status_30), int32(1))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v546 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v550 = F_LWLockAcquire(m, v546+int32(_a_F_pg_lock_status_31), int32(1))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v557 = F_LWLockAcquire(m, v553+int32(_a_F_pg_lock_status_32), int32(1))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v560 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[5]))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v560)))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v562)+4))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v562)+412))
	if v564 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v628 = v627 + v432
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v628
	if v435 < v628 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v562)+376))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v562)+364))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v562)+352))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v562)+340))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v562)+328))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v562)+316))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v562)+304))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v562)+292))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v562)+280))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v562)+268))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v562)+256))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v562)+244))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v562)+232))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v562)+220))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v562)+208))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v562)+196))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v562)+184))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v562)+172))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v562)+160))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v562)+148))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v562)+136))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v562)+124))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v562)+112))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v562)+100))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v562)+88))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v562)+76))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v562)+64))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v562)+52))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v562)+40))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v562)+28))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v562)+16))
	v627 = v565 + (v566 + (v567 + (v568 + (v569 + (v570 + (v571 + (v572 + (v573 + (v574 + (v575 + (v576 + (v577 + (v578 + (v579 + (v580 + (v581 + (v582 + (v583 + (v584 + (v585 + (v586 + (v587 + (v588 + (v589 + (v590 + (v591 + (v592 + (v593 + (v594 + (v595 + v563))))))))))))))))))))))))))))))
	goto L83
L82:
	;
	v627 = v563
	goto L83
L83:
	;
	goto L80
L84:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v634 = F_repalloc(m, v631, v628*int32(56))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L4
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v638 = v160 + int32(12)
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[5]))
	F_hash_seq_init(m, v638, v640)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L4
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v634
	goto L86
L88:
	;
	v643 = F_hash_seq_search(m, v638)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	if v643 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v646 = v643
	v648 = v432
	goto L93
L91:
	;
	goto L92
L92:
	;
	v723 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v723+int32(_a_F_pg_lock_status_32))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L4
	} else {
		goto L100
	}
L93:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v646)+4))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v667 = v664 + v648*int32(56)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v646)))
	v669 = *(*int64)(unsafe.Add(mBase, uint32(v668)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v667)+8)) = v669
	v671 = *(*int64)(unsafe.Add(mBase, uint32(v668)))
	*(*int64)(unsafe.Add(mBase, uint32(v667))) = v671
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v646)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v667)+16)) = v673
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v663)+92))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v646)))
	if v675 == v676 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L92
L95:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v663)+100))
	v680 = v678
	goto L97
L96:
	;
	v680 = int32(0)
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v667)+20)) = v680
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v663)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v667)+24)) = v682
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v663)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v667)+28)) = v684
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v663)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v667)+40)) = v686
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v646)+8))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v688)+44))
	v690 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v667)+48)) = uint8(v690)
	*(*int32)(unsafe.Add(mBase, uint32(v667)+44)) = v689
	v693 = int64(0)
	v696 = base.AtomicRmwCmpxchg64(m, v663, int32(112), v693, v693)
	*(*int64)(unsafe.Add(mBase, uint32(v667)+32)) = v696
	v702 = F_hash_seq_search(m, v160+int32(12))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	if v702 != 0 {
		v646 = v702
		v648 = v648 + int32(1)
		goto L93
	} else {
		goto L99
	}
L99:
	;
	goto L94
L100:
	;
	v729 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v729+int32(_a_F_pg_lock_status_31))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	v735 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v735+int32(_a_F_pg_lock_status_30))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	v741 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v741+int32(_a_F_pg_lock_status_29))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v747 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v747+int32(_a_F_pg_lock_status_28))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v753 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v753+int32(_a_F_pg_lock_status_27))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v759 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v759+int32(_a_F_pg_lock_status_26))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	v765 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v765+int32(_a_F_pg_lock_status_25))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v771 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v771+int32(_a_F_pg_lock_status_24))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v777 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v777+int32(_a_F_pg_lock_status_23))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v783 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v783+int32(_a_F_pg_lock_status_22))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	v789 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v789+int32(_a_F_pg_lock_status_21))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	v795 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v795+int32(_a_F_pg_lock_status_20))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	v801 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v801+int32(_a_F_pg_lock_status_19))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	v807 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v807+int32(_a_F_pg_lock_status_18))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	v813 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v813+int32(_a_F_pg_lock_status_17))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	v818 = int32(32)
	m.G0 = v160 + v818
	v821 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v155)+4)) = v821
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = v163
	v825 = m.G0
	v827 = v825 - v818
	m.G0 = v827
	v830 = F_palloc(m, int32(12))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	v833 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v837 = F_LWLockAcquire(m, v833+int32(_a_F_pg_lock_status_33), int32(1))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	v840 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v844 = F_LWLockAcquire(m, v840+int32(_a_F_pg_lock_status_34), int32(1))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	v847 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v851 = F_LWLockAcquire(m, v847+int32(_a_F_pg_lock_status_35), int32(1))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v858 = F_LWLockAcquire(m, v854+int32(_a_F_pg_lock_status_36), int32(1))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	v861 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v865 = F_LWLockAcquire(m, v861+int32(_a_F_pg_lock_status_37), int32(1))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	v868 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v872 = F_LWLockAcquire(m, v868+int32(_a_F_pg_lock_status_38), int32(1))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	v875 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v879 = F_LWLockAcquire(m, v875+int32(_a_F_pg_lock_status_39), int32(1))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	v882 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v886 = F_LWLockAcquire(m, v882+int32(_a_F_pg_lock_status_40), int32(1))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v893 = F_LWLockAcquire(m, v889+int32(_a_F_pg_lock_status_41), int32(1))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v896 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v900 = F_LWLockAcquire(m, v896+int32(_a_F_pg_lock_status_42), int32(1))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	v903 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v907 = F_LWLockAcquire(m, v903+int32(_a_F_pg_lock_status_43), int32(1))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v910 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v914 = F_LWLockAcquire(m, v910+int32(_a_F_pg_lock_status_44), int32(1))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	v917 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v921 = F_LWLockAcquire(m, v917+int32(_a_F_pg_lock_status_45), int32(1))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	v924 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v928 = F_LWLockAcquire(m, v924+int32(_a_F_pg_lock_status_46), int32(1))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v931 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v935 = F_LWLockAcquire(m, v931+int32(_a_F_pg_lock_status_47), int32(1))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	v938 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v942 = F_LWLockAcquire(m, v938+int32(_a_F_pg_lock_status_48), int32(1))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	v945 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	v949 = F_LWLockAcquire(m, v945+int32(3584), int32(1))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	v952 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[6]))
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v952)))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v954)+4))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v954)+412))
	if v956 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v830))) = v1019
	v1023 = F_palloc(m, v1019<<(uint(int32(4))%32))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L4
	} else {
		goto L138
	}
L135:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v954)+376))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v954)+364))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v954)+352))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v954)+340))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v954)+328))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v954)+316))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v954)+304))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v954)+292))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v954)+280))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v954)+268))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v954)+256))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v954)+244))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v954)+232))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v954)+220))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v954)+208))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v954)+196))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v954)+184))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v954)+172))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v954)+160))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v954)+148))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v954)+136))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v954)+124))
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v954)+112))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v954)+100))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v954)+88))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v954)+76))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v954)+64))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v954)+52))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v954)+40))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v954)+28))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v954)+16))
	v1019 = v957 + (v958 + (v959 + (v960 + (v961 + (v962 + (v963 + (v964 + (v965 + (v966 + (v967 + (v968 + (v969 + (v970 + (v971 + (v972 + (v973 + (v974 + (v975 + (v976 + (v977 + (v978 + (v979 + (v980 + (v981 + (v982 + (v983 + (v984 + (v985 + (v986 + (v987 + v955))))))))))))))))))))))))))))))
	goto L137
L136:
	;
	v1019 = v955
	goto L137
L137:
	;
	goto L134
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v830)+4)) = v1023
	v1028 = F_palloc(m, v1019*int32(120))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v830)+8)) = v1028
	v1032 = v827 + int32(12)
	v1034 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[6]))
	F_hash_seq_init(m, v1032, v1034)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	v1037 = F_hash_seq_search(m, v1032)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	if v1037 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v1040 = v821
	v1043 = v1037
	goto L145
L143:
	;
	goto L144
L144:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1098+int32(3584))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L4
	} else {
		goto L149
	}
L145:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v830)+4))
	v1060 = v1057 + v1040<<(uint(int32(4))%32)
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1043)))
	v1062 = *(*int64)(unsafe.Add(mBase, uint32(v1061)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1060)+8)) = v1062
	v1064 = *(*int64)(unsafe.Add(mBase, uint32(v1061)))
	*(*int64)(unsafe.Add(mBase, uint32(v1060))) = v1064
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v830)+8))
	v1067 = int32(120)
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1043)+4))
	base.MemoryCopy(m, v1066+v1040*v1067, v1070, v1067)
	v1077 = F_hash_seq_search(m, v827+int32(12))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L4
	} else {
		goto L147
	}
L146:
	;
	goto L144
L147:
	;
	if v1077 != 0 {
		v1040 = v1040 + int32(1)
		v1043 = v1077
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1104+int32(_a_F_pg_lock_status_48))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1110+int32(_a_F_pg_lock_status_47))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1116+int32(_a_F_pg_lock_status_46))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1122+int32(_a_F_pg_lock_status_45))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1128+int32(_a_F_pg_lock_status_44))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L4
	} else {
		goto L154
	}
L154:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1134+int32(_a_F_pg_lock_status_43))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1140+int32(_a_F_pg_lock_status_42))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1146+int32(_a_F_pg_lock_status_41))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1152+int32(_a_F_pg_lock_status_40))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1158+int32(_a_F_pg_lock_status_39))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1164+int32(_a_F_pg_lock_status_38))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1170+int32(_a_F_pg_lock_status_37))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1176+int32(_a_F_pg_lock_status_36))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1182+int32(_a_F_pg_lock_status_35))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1188+int32(_a_F_pg_lock_status_34))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[3]))
	F_LWLockRelease(m, v1194+int32(_a_F_pg_lock_status_33))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	m.G0 = v827 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v155)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v155)+8)) = v830
	*(*int32)(unsafe.Add(mBase, _c_F_pg_lock_status[0])) = v32
	goto L3
L166:
	;
	m.G0 = v21 + int32(208)
	return v1646
L167:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1226)+16))
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+4))
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1227)))
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1229)))
	if v1228 < v1230 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	goto L171
L169:
	;
	goto L170
L170:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+12))
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+8))
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v1529)))
	if v1528 < v1530 {
		goto L241
	} else {
		goto L242
	}
L171:
	;
	v1250 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+136)) = v1250
	*(*int64)(unsafe.Add(mBase, uint32(v21)+128)) = v1250
	*(*int64)(unsafe.Add(mBase, uint32(v21)+120)) = v1250
	*(*int64)(unsafe.Add(mBase, uint32(v21)+112)) = v1250
	*(*int64)(unsafe.Add(mBase, uint32(v21)+104)) = v1250
	*(*int64)(unsafe.Add(mBase, uint32(v21)+96)) = v1250
	*(*int64)(unsafe.Add(mBase, uint32(v21)+88)) = v1250
	*(*int64)(unsafe.Add(mBase, uint32(v21)+80)) = v1250
	*(*int64)(unsafe.Add(mBase, uint32(v21)+72)) = v1250
	*(*int64)(unsafe.Add(mBase, uint32(v21)+64)) = v1250
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1229)+4))
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+4))
	v1274 = v1270 + v1271*int32(56)
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+16))
	if v1275 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	goto L170
L173:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1227)+4))
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v1229)))
	if v1507 < v1508 {
		goto L171
	} else {
		goto L240
	}
L174:
	;
	v1334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1274)+14)))
	if base.Ui32(v1334) <= base.Ui32(int32(11)) {
		goto L208
	} else {
		goto L209
	}
L175:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1227)+4)) = v1271 + int32(1)
	if v1325 == int32(0) {
		goto L173
	} else {
		goto L206
	}
L176:
	;
	if v1275&int32(1) != 0 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1274)+16)) = v1321 & v1275
	v1332 = v1320
	v1333 = int32(1)
	goto L174
L178:
	;
	v1320 = int32(0)
	v1321 = int32(-2)
	goto L177
L179:
	;
	goto L180
L180:
	;
	if v1275&int32(2) != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v1320 = int32(1)
	v1321 = int32(-3)
	goto L177
L182:
	;
	goto L183
L183:
	;
	if v1275&int32(4) != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1320 = int32(2)
	v1321 = int32(-5)
	goto L177
L185:
	;
	goto L186
L186:
	;
	if v1275&int32(8) != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v1320 = int32(3)
	v1321 = int32(-9)
	goto L177
L188:
	;
	goto L189
L189:
	;
	if v1275&int32(16) != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v1320 = int32(4)
	v1321 = int32(-17)
	goto L177
L191:
	;
	goto L192
L192:
	;
	if v1275&int32(32) != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1320 = int32(5)
	v1321 = int32(-33)
	goto L177
L194:
	;
	goto L195
L195:
	;
	if v1275&int32(64) != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1320 = int32(6)
	v1321 = int32(-65)
	goto L177
L197:
	;
	goto L198
L198:
	;
	if v1275&int32(128) != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v1320 = int32(7)
	v1321 = int32(-129)
	goto L177
L200:
	;
	goto L201
L201:
	;
	if v1275&int32(256) != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v1320 = int32(8)
	v1321 = int32(-257)
	goto L177
L203:
	;
	goto L204
L204:
	;
	if v1275&int32(512) == int32(0) {
		goto L175
	} else {
		goto L205
	}
L205:
	;
	v1320 = int32(9)
	v1321 = int32(-513)
	goto L177
L206:
	;
	v1332 = v1325
	v1333 = int32(0)
	goto L174
L207:
	;
	v1351 = F_cstring_to_text(m, v1350)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L4
	} else {
		goto L212
	}
L208:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1334<<(uint(int32(2))%32))+uint32(_c_F_pg_lock_status[7])))
	v1350 = v1339
	goto L207
L209:
	;
	goto L210
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v1334
	v1342 = v21 + int32(144)
	v1347 = F_pg_snprintf(m, v1342, int32(32), int32(_a_F_pg_lock_status_49), v21+int32(48))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L4
	} else {
		goto L211
	}
L211:
	;
	v1350 = v1342
	goto L207
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v1351
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1274)))
	v1355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1274)+14)))
	switch v1355 {
	case 0, 1:
		goto L222
	case 2:
		goto L221
	case 3:
		goto L220
	case 4:
		goto L219
	case 5:
		goto L218
	case 6:
		goto L217
	case 7:
		goto L216
	default:
		goto L214
	case 11:
		goto L215
	}
L213:
	;
	v1448 = *(*int64)(unsafe.Add(mBase, uint32(v1274)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v1448
	v1451 = v21 + int32(176)
	v1456 = F_pg_snprintf(m, v1451, int32(32), int32(_a_F_pg_lock_status_50), v21+int32(16))
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L4
	} else {
		goto L225
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1354
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+108)) = v1436
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v1438
	v1440 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1274)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+66)) = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = v1440
	v1444 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+70)) = uint8(v1444)
	goto L213
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1354
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v1425
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v1427
	v1429 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1274)+12)))
	v1430 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+71)) = uint8(v1430)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = v1429
	*(*int32)(unsafe.Add(mBase, uint32(v21)+66)) = int32(16843009)
	goto L213
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v1354
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+4))
	v1415 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+73)) = uint8(v1415)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+71)) = uint8(v1415)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v1414
	*(*int32)(unsafe.Add(mBase, uint32(v21)+65)) = int32(16843009)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+69)) = uint8(v1415)
	goto L213
L217:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v1395
	v1399 = v21 + int32(176)
	v1400 = int32(32)
	v1404 = F_pg_snprintf(m, v1399, v1400, int32(_a_F_pg_lock_status_50), v21+v1400)
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L4
	} else {
		goto L223
	}
L218:
	;
	v1386 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+73)) = uint8(v1386)
	v1388 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+71)) = uint16(v1388)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v1354
	*(*int32)(unsafe.Add(mBase, uint32(v21)+65)) = int32(16843009)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+69)) = uint8(v1386)
	goto L213
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1354
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v1376
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v1378
	v1380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1274)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+69)) = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v1380
	v1384 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+73)) = uint8(v1384)
	goto L213
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1354
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v1367
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v1369
	v1373 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+72)) = uint16(v1373)
	goto L213
L221:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+66)) = int64(72340172838076673)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1354
	goto L213
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1354
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+4))
	v1358 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+67)) = v1358
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v1357
	*(*int32)(unsafe.Add(mBase, uint32(v21)+70)) = v1358
	goto L213
L223:
	;
	v1406 = F_cstring_to_text(m, v1399)
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L4
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = v1406
	v1409 = int32(16843009)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+70)) = v1409
	*(*int32)(unsafe.Add(mBase, uint32(v21)+65)) = v1409
	goto L213
L225:
	;
	v1458 = F_cstring_to_text(m, v1451)
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L4
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+120)) = v1458
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+40))
	if v1461 != 0 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v1465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1274)+15)))
	v1466 = int32(2)
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1465<<(uint(v1466)%32))+uint32(_c_F_pg_lock_status[8])))
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1468)+8))
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1469+v1332<<(uint(v1466)%32))))
	goto L231
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+124)) = v1461
	goto L227
L229:
	;
	goto L230
L230:
	;
	v1463 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+75)) = uint8(v1463)
	goto L227
L231:
	;
	v1474 = F_cstring_to_text(m, v1473)
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L4
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+132)) = v1333
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v1474
	v1478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1274)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+136)) = v1478
	if v1333 != 0 {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1226)+28))
	v1495 = F_heap_form_tuple(m, v1490, v21+int32(80), v21-int32(-64))
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L4
	} else {
		goto L238
	}
L234:
	;
	v1487 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+79)) = uint8(v1487)
	goto L233
L235:
	;
	v1480 = *(*int64)(unsafe.Add(mBase, uint32(v1274)+32))
	if v1480 == int64(0) {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v1483 = F_Int64GetDatum(m, v1480)
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L4
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+140)) = v1483
	goto L233
L238:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1495)+16))
	v1498 = F_HeapTupleHeaderGetDatum(m, v1497)
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L4
	} else {
		goto L239
	}
L239:
	;
	v1500 = *(*int64)(unsafe.Add(mBase, uint32(v1226)))
	*(*int64)(unsafe.Add(mBase, uint32(v1226))) = v1500 + int64(1)
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1504)+20)) = int32(1)
	v1646 = v1498
	goto L166
L240:
	;
	goto L172
L241:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v1529)+4))
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1529)+8))
	v1534 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+136)) = v1534
	*(*int64)(unsafe.Add(mBase, uint32(v21)+128)) = v1534
	*(*int64)(unsafe.Add(mBase, uint32(v21)+120)) = v1534
	*(*int64)(unsafe.Add(mBase, uint32(v21)+112)) = v1534
	*(*int64)(unsafe.Add(mBase, uint32(v21)+104)) = v1534
	*(*int64)(unsafe.Add(mBase, uint32(v21)+96)) = v1534
	*(*int64)(unsafe.Add(mBase, uint32(v21)+88)) = v1534
	*(*int64)(unsafe.Add(mBase, uint32(v21)+80)) = v1534
	*(*int64)(unsafe.Add(mBase, uint32(v21)+152)) = v1534
	*(*int64)(unsafe.Add(mBase, uint32(v21)+144)) = v1534
	*(*int32)(unsafe.Add(mBase, uint32(v1227)+12)) = v1528 + int32(1)
	v1560 = v1532 + v1528<<(uint(int32(4))%32)
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+12))
	if v1561 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	goto L243
L243:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L4
	} else {
		goto L264
	}
L244:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+8))
	v1567 = base.B2i32(v1564 != int32(-1))
	goto L246
L245:
	;
	v1567 = int32(2)
	goto L246
L246:
	;
	v1570 = v1528*int32(120) + v1533
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1567<<(uint(int32(2))%32))+uint32(_c_F_pg_lock_status[9])))
	v1574 = F_cstring_to_text(m, v1573)
	mBase = m.M
	v1575 = m.ExcPending
	if v1575 != 0 {
		goto L4
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v1574
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1560)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v1577
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v1579
	if v1567 == int32(2) {
		goto L251
	} else {
		goto L252
	}
L248:
	;
	v1593 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+153)) = uint8(v1593)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+149)) = int32(16843009)
	v1597 = *(*int64)(unsafe.Add(mBase, uint32(v1570)))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v1597
	v1600 = v21 + int32(176)
	v1603 = F_pg_snprintf(m, v1600, int32(32), int32(_a_F_pg_lock_status_50), v21)
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L4
	} else {
		goto L255
	}
L249:
	;
	v1591 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+147)) = uint8(v1591)
	goto L248
L250:
	;
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v1589
	goto L248
L251:
	;
	v1583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1560)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v1583
	goto L250
L252:
	;
	goto L253
L253:
	;
	v1585 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+148)) = uint8(v1585)
	if v1567 != v1585 {
		goto L249
	} else {
		goto L254
	}
L254:
	;
	goto L250
L255:
	;
	v1605 = F_cstring_to_text(m, v1600)
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L4
	} else {
		goto L256
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+120)) = v1605
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1570)+112))
	if v1608 != 0 {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v1613 = F_cstring_to_text(m, int32(_a_F_pg_lock_status_51))
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L4
	} else {
		goto L261
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+124)) = v1608
	goto L257
L259:
	;
	goto L260
L260:
	;
	v1610 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+155)) = uint8(v1610)
	goto L257
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v1613
	v1616 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+159)) = uint8(v1616)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+132)) = int64(1)
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v1226)+28))
	v1625 = F_heap_form_tuple(m, v1620, v21+int32(80), v21+int32(144))
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L4
	} else {
		goto L262
	}
L262:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+16))
	v1628 = F_HeapTupleHeaderGetDatum(m, v1627)
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L4
	} else {
		goto L263
	}
L263:
	;
	v1630 = *(*int64)(unsafe.Add(mBase, uint32(v1226)))
	*(*int64)(unsafe.Add(mBase, uint32(v1226))) = v1630 + int64(1)
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1634)+20)) = int32(1)
	v1646 = v1628
	goto L166
L264:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1639)+20)) = int32(2)
	v1642 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1642)
	v1646 = int32(0)
	goto L166
}
func F_pg_log_backend_memory_contexts(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
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
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_BackendPidGetProc(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v92
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	F_errmsg(m, v82, v7)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L25
	}
L3:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_pg_log_backend_memory_contexts[0]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v67 = base.I32_div_s(v59-v64, int32(640))
	v68 = F_SendProcSignal(m, v9, int32(5), v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L21
	}
L4:
	;
	return int32(0)
L5:
	;
	if v10 != 0 {
		v59 = v10
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v14 = int32(0)
	if v9 == v14 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v49 != 0 {
		v59 = v49
		goto L3
	} else {
		goto L18
	}
L8:
	;
	v49 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_log_backend_memory_contexts[1]))
	v26 = v14
	goto L13
L11:
	;
	v49 = v43
	goto L7
L12:
	;
	v43 = v33 + int32(640)
	goto L11
L13:
	;
	v29 = v26 * int32(640)
	v30 = v22 + v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+44))
	if v31 == v9 {
		v43 = v30
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v49 = int32(0)
	goto L7
L15:
	;
	v33 = v22 + v29
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+684))
	if v34 == v9 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v37 = v26 + int32(2)
	if v37 != int32(38) {
		v26 = v37
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v50 = int32(0)
	v53 = F_errstart(m, int32(19), v50)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if v53 == int32(0) {
		v92 = v50
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v82 = int32(_a_F_pg_log_backend_memory_contexts_0)
	v83 = int32(293)
	goto L2
L21:
	;
	if int32(0) <= v68 {
		v92 = int32(1)
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v72 = int32(0)
	v75 = F_errstart(m, int32(19), v72)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v75 == int32(0) {
		v92 = v72
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v82 = int32(_a_F_pg_log_backend_memory_contexts_1)
	v83 = int32(302)
	goto L2
L25:
	;
	F_errfinish(m, int32(_a_F_pg_log_backend_memory_contexts_2), v83, int32(_a_F_pg_log_backend_memory_contexts_3))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v92 = int32(0)
	goto L1
}
func F_pg_md5_binary(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v5 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	v9 = F_pg_cryptohash_create(m, v5)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(_a_F_pg_md5_binary_0)
			return int32(0)
		} else {
			v34 = F_pg_cryptohash_init(m, v9)
			mBase = m.M
			if v34 < int32(0) {
				if v9 == int32(0) {
					v58 = int32(_a_F_pg_md5_binary_0)
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					if v50 == int32(1) {
						v53 = int32(_a_F_pg_md5_binary_1)
					} else {
						v53 = int32(_a_F_pg_md5_binary_2)
					}
					if v50 == int32(2) {
						v56 = int32(_a_F_pg_md5_binary_0)
					} else {
						v56 = v53
					}
					v58 = v56
				}
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v58
				F_pg_cryptohash_free(m, v9)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			} else {
				v37 = F_pg_cryptohash_update(m, v9, l0, l1)
				mBase = m.M
				if v37 < int32(0) {
					if v9 == int32(0) {
						v58 = int32(_a_F_pg_md5_binary_0)
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						if v50 == int32(1) {
							v53 = int32(_a_F_pg_md5_binary_1)
						} else {
							v53 = int32(_a_F_pg_md5_binary_2)
						}
						if v50 == int32(2) {
							v56 = int32(_a_F_pg_md5_binary_0)
						} else {
							v56 = v53
						}
						v58 = v56
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v58
					F_pg_cryptohash_free(m, v9)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					v41 = F_pg_cryptohash_final(m, v9, l2, int32(16))
					mBase = m.M
					if int32(0) <= v41 {
						F_pg_cryptohash_free(m, v9)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							return int32(1)
						}
					} else {
						if v9 == int32(0) {
							v58 = int32(_a_F_pg_md5_binary_0)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							if v50 == int32(1) {
								v53 = int32(_a_F_pg_md5_binary_1)
							} else {
								v53 = int32(_a_F_pg_md5_binary_2)
							}
							if v50 == int32(2) {
								v56 = int32(_a_F_pg_md5_binary_0)
							} else {
								v56 = v53
							}
							v58 = v56
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v58
						F_pg_cryptohash_free(m, v9)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				}
			}
		}
	}
}
func F_pg_ndistinct_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_pg_ndistinct_in_0), int32(343), int32(_a_F_pg_ndistinct_in_1), int32(_a_F_pg_ndistinct_in_2), int32(_a_F_pg_ndistinct_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_pg_node_tree_recv(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13874(m, l0, int32(_a_F_pg_node_tree_recv_0), int32(335), int32(_a_F_pg_node_tree_recv_1), int32(_a_F_pg_node_tree_recv_2), int32(_a_F_pg_node_tree_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
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
func F_pg_plan_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 == int32(6) {
		v46 = int32(0)
		return v46
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_query[0])))
		if v10 == int32(1) {
			v13 = int32(_a_F_pg_plan_query_0)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_query[1])) = int64(4)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_query[2])) = int64(3)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_query[3])) = int64(2)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_plan_query[4])) = int64(1)
			v23 = F___syscall_ret(m, int32(0))
			mBase = m.M
			F_gettimeofday(m, int32(_a_F_pg_plan_query_1))
			mBase = m.M
		} else {
		}
		v26 = F_planner(m, l0, l1, l2, l3)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_query[0])))
			if v31 == int32(1) {
				F_ShowUsage(m, int32(_a_F_pg_plan_query_2))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_query[5])))
					if v38 != int32(1) {
						v46 = v26
						return v46
					} else {
						v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_query[6])))
						F_elog_node_display(m, int32(_a_F_pg_plan_query_3), v26, v43)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = v26
							return v46
						}
					}
				}
			} else {
				v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_query[5])))
				if v38 != int32(1) {
					v46 = v26
					return v46
				} else {
					v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_plan_query[6])))
					F_elog_node_display(m, int32(_a_F_pg_plan_query_3), v26, v43)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v46 = v26
						return v46
					}
				}
			}
		}
	}
}
func F_pg_prepared_xact(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int64
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v290 int32
	_ = v290
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v18 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	goto L33
L4:
	;
	return int32(0)
L5:
	;
	v25 = int32(_a_F_pg_prepared_xact_0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[0])) = v28
	v31 = F_CreateTemplateTupleDesc(m, int32(5))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v31, int32(1), int32(_a_F_pg_prepared_xact_1), int32(28), int32(-1), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v31, int32(2), int32(_a_F_pg_prepared_xact_2), int32(25), int32(-1), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v31, int32(3), int32(_a_F_pg_prepared_xact_3), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v31, int32(4), int32(_a_F_pg_prepared_xact_4), int32(26), int32(-1), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v31, int32(5), int32(_a_F_pg_prepared_xact_5), int32(26), int32(-1), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v68 = F_BlessTupleDesc(m, v31)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v68
	v72 = F_palloc(m, int32(12))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v72
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[1]))
	v80 = F_LWLockAcquire(m, v76+int32(2304), int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[2]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v84 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v84
	*(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[0])) = v26
	goto L3
L16:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[1]))
	F_LWLockRelease(m, v88+int32(2304))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v97 = F_palloc(m, v84*int32(248))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(0)
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v97
	if v84 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[1]))
	F_LWLockRelease(m, v186+int32(2304))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L30
	}
L22:
	;
	v102 = int32(0)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[2]))
	v106 = v104 + int32(8)
	if v84 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v114 = v102
	v121 = v2
	goto L26
L24:
	;
	v153 = v102
	goto L25
L25:
	;
	v164 = int32(248)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v106+v153<<(uint(int32(2))%32))))
	base.MemoryCopy(m, v97+v153*v164, v170, v164)
	goto L21
L26:
	;
	v125 = int32(248)
	v128 = int32(2)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v106+v114<<(uint(v128)%32))))
	base.MemoryCopy(m, v97+v114*v125, v131, v125)
	v135 = v114 | int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v106+v135<<(uint(v128)%32))))
	base.MemoryCopy(m, v97+v135*v125, v142, v125)
	v146 = v114 + v128
	v148 = v121 + v128
	if v148 != v84&int32(2147483646) {
		v114 = v146
		v121 = v148
		goto L26
	} else {
		goto L28
	}
L27:
	;
	if v84&int32(1) == int32(0) {
		goto L21
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v153 = v146
	goto L25
L30:
	;
	goto L15
L31:
	;
	m.G0 = v15 + int32(48)
	return v315
L32:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L46
	}
L33:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+16))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	if v223 == int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v227 <= v226 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_xact[3]))
	v232 = v226
	goto L36
L36:
	;
	v243 = int32(1)
	v244 = v232 + v243
	*(*int32)(unsafe.Add(mBase, uint32(v222)+8)) = v244
	v248 = v223 + v232*int32(248)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v251 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v251)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v251
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+44)))
	if v255 == v243 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L32
L38:
	;
	v260 = v250 + v249*int32(640)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v261
	v265 = F_cstring_to_text(m, v248+int32(47))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v244 < v227 {
		v232 = v244
		goto L36
	} else {
		goto L45
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v265
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v248)+8))
	v269 = F_Int64GetDatum(m, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v269
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v248)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v272
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v260)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v274
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v221)+28))
	v281 = F_heap_form_tuple(m, v276, v15+int32(16), v15+int32(8))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v281)+16))
	v284 = F_HeapTupleHeaderGetDatum(m, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v221)))
	*(*int64)(unsafe.Add(mBase, uint32(v221))) = v286 + int64(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v290)+20)) = int32(1)
	v315 = v284
	goto L31
L45:
	;
	goto L37
L46:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v308)+20)) = int32(2)
	v311 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v311)
	v315 = int32(0)
	goto L31
}
func F_pg_promote(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_promote[0])))
	if v14 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v219 = F_unlink(m, int32(_a_F_pg_promote_0))
	mBase = m.M
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L13
	} else {
		goto L69
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L13
	} else {
		goto L65
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L13
	} else {
		goto L61
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L13
	} else {
		goto L57
	}
L5:
	;
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_pg_promote[1]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+316))
	v22 = base.B2i32(v20 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_promote[0])) = uint8(v22)
	v24 = v22
	goto L8
L7:
	;
	v24 = int32(0)
	goto L8
L8:
	;
	goto L5
L9:
	;
	if v10 <= int32(0) {
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
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L52
	}
L12:
	;
	v29 = F_AllocateFile(m, int32(_a_F_pg_promote_0), int32(_a_F_pg_promote_1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	if v29 == int32(0) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v35 = F_FreeFile(m, v29)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v35 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_pg_promote[2]))
	v40 = F_pgmem_kill(m, v38, int32(10))
	mBase = m.M
	if v40 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v11 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	m.G0 = v8 + int32(48)
	return v141
L20:
	;
	v141 = int32(1)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v44 = int32(0)
	v46 = v10 * int32(10)
	if v46 <= v44 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v120 = int32(0)
	v123 = F_errstart(m, int32(19), v120)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L13
	} else {
		goto L48
	}
L24:
	;
	v50 = v44
	goto L25
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_pg_promote[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = int32(0)
	goto L27
L26:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L13
	} else {
		goto L42
	}
L27:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_promote[0])))
	if v61 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v71 == int32(0) {
		v141 = int32(1)
		goto L19
	} else {
		goto L32
	}
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_pg_promote[1]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+316))
	v69 = base.B2i32(v67 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_promote[0])) = uint8(v69)
	v71 = v69
	goto L31
L30:
	;
	v71 = int32(0)
	goto L31
L31:
	;
	goto L28
L32:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_pg_promote[4]))
	if v75 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_pg_promote[3]))
	v83 = F_WaitLatch(m, v79, int32(25), int32(100), int32(134217771))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L13
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	if v83&int32(16) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v90 = v50 + int32(1)
	if v90 == v46 {
		goto L23
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	goto L26
L41:
	;
	v50 = v90
	goto L25
L42:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_pg_promote_2), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	F_errcontext_msg(m, int32(_a_F_pg_promote_3), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_pg_promote_4), int32(741), int32(_a_F_pg_promote_5))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	if v123 == int32(0) {
		v141 = v120
		goto L19
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	F_errmsg_plural(m, int32(_a_F_pg_promote_6), int32(_a_F_pg_promote_7), v10, v8+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L13
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_pg_promote_4), int32(748), int32(_a_F_pg_promote_5))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	v141 = v120
	goto L19
L52:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L13
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(_a_F_pg_promote_8), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	F_errhint(m, int32(_a_F_pg_promote_9), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_pg_promote_4), int32(681), int32(_a_F_pg_promote_5))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_pg_promote_10), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L13
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_pg_promote_4), int32(686), int32(_a_F_pg_promote_5))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_pg_promote_0)
	F_errmsg(m, int32(_a_F_pg_promote_11), v8)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_pg_promote_4), int32(694), int32(_a_F_pg_promote_5))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L13
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L13
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_pg_promote_0)
	F_errmsg(m, int32(_a_F_pg_promote_12), v8+int32(32))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L13
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_pg_promote_4), int32(700), int32(_a_F_pg_promote_5))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L13
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(517))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(_a_F_pg_promote_13), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_pg_promote_4), int32(708), int32(_a_F_pg_promote_5))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_read_binary_file_all_missing(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_convert_and_check_filename(m, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v16 = F_read_binary_file(m, v10, int64(0), int64(-1), base.B2i32(v9 != int32(0)))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v16 == int32(0) {
					v20 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
					return int32(0)
				} else {
					return v16
				}
			}
		}
	}
}
func F_pg_read_binary_file_off_len_missing(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		if int64(0) <= v12 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
			v18 = F_convert_and_check_filename(m, v7)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v22 = F_read_binary_file(m, v18, v17, v12, base.B2i32(v15 != int32(0)))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					if v22 == int32(0) {
						v26 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
						return int32(0)
					} else {
						return v22
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_read_binary_file_off_len_missing_0), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_read_binary_file_off_len_missing_1), int32(269), int32(_a_F_pg_read_binary_file_off_len_missing_2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
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
	var v15 int32
	_ = v15
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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v268 int32
	_ = v268
	var v291 int32
	_ = v291
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v527 int32
	_ = v527
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v688 int32
	_ = v688
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
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
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v856 int32
	_ = v856
	var v867 int32
	_ = v867
	var v879 int32
	_ = v879
	var v883 int32
	_ = v883
	var v890 int32
	_ = v890
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v930 int32
	_ = v930
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v962 int32
	_ = v962
	var v969 int32
	_ = v969
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1149 int32
	_ = v1149
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
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
	v14 = v3
	v15 = v8
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
	v21 = base.I32_div_s(v15-v14, int32(2))
	v22 = v21 + v14
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
	if v32 < v33 {
		v14 = v32
		v15 = v33
		goto L6
	} else {
		goto L13
	}
L9:
	;
	v32 = v14
	v33 = v22
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
	v32 = v22 + int32(1)
	v33 = v15
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
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v1204 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1200+v51<<(uint(int32(1))%32)))))
	return v1204
L17:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v307 == int32(0) {
		v468 = v306
		goto L91
	} else {
		goto L92
	}
L18:
	;
	v306 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v59 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	switch v61 - int32(1) {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	default:
		goto L22
	}
L21:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v306 = v305
	goto L17
L22:
	;
	if base.Ui32(l1-int32(127)) < base.Ui32(int32(-95)) {
		v306 = v59
		goto L17
	} else {
		goto L90
	}
L23:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v306 = v59
		goto L17
	} else {
		goto L87
	}
L24:
	;
	if base.Ui32(l1) <= base.Ui32(int32(254)) {
		goto L83
	} else {
		goto L84
	}
L25:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	if v254 != 0 {
		goto L21
	} else {
		goto L81
	}
L27:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l1) {
		goto L47
	} else {
		goto L48
	}
L28:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v108 != int32(15) {
		goto L27
	} else {
		goto L41
	}
L29:
	;
	v107 = v80 + int32(_a_F_pg_reg_getcolor_0)
	goto L28
L30:
	;
	v72 = int32(3367)
	v73 = int32(0)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v107 = l1<<(uint(int32(1))%32) + int32(_a_F_pg_reg_getcolor_1)
	goto L28
L33:
	;
	v78 = base.I32_div_s(v72+v73, int32(2))
	v80 = v78 * int32(12)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_c_F_pg_reg_getcolor[1])))
	if base.Ui32(v83) < base.Ui32(l1) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L27
L35:
	;
	if v94 <= v93 {
		v72 = v93
		v73 = v94
		goto L33
	} else {
		goto L40
	}
L36:
	;
	v93 = v72
	v94 = v78 + int32(1)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v80)+uint32(_c_F_pg_reg_getcolor[2])))
	if base.Ui32(v89) <= base.Ui32(l1) {
		goto L29
	} else {
		goto L39
	}
L39:
	;
	v93 = v78 - int32(1)
	v94 = v73
	goto L35
L40:
	;
	goto L34
L41:
	;
	v254 = int32(0)
	goto L26
L42:
	;
	v254 = v246
	goto L26
L43:
	;
	v246 = base.B2i32(v236&int32(255) == int32(12))
	goto L42
L44:
	;
	if int32(1)<<(uint(v168)%32)&int32(_a_F_pg_reg_getcolor_2) != 0 {
		goto L62
	} else {
		goto L63
	}
L45:
	;
	if l1 == int32(9) {
		v246 = v147
		goto L42
	} else {
		goto L60
	}
L46:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+uint32(_c_F_pg_reg_getcolor[3]))))
	v168 = v159
	goto L44
L47:
	;
	v122 = int32(3367)
	v123 = int32(0)
	goto L50
L48:
	;
	goto L49
L49:
	;
	v147 = int32(1)
	v150 = l1 << (uint(v147) % 32)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+uint32(_c_F_pg_reg_getcolor[4]))))
	if v147<<(uint(v151)%32)&int32(_a_F_pg_reg_getcolor_2) == int32(0) {
		goto L45
	} else {
		goto L58
	}
L50:
	;
	v128 = base.I32_div_s(v122+v123, int32(2))
	v130 = v128 * int32(12)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)+uint32(_c_F_pg_reg_getcolor[1])))
	if base.Ui32(v133) < base.Ui32(l1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v168 = int32(0)
	goto L44
L52:
	;
	if v144 <= v143 {
		v122 = v143
		v123 = v144
		goto L50
	} else {
		goto L57
	}
L53:
	;
	v143 = v122
	v144 = v128 + int32(1)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v130)+uint32(_c_F_pg_reg_getcolor[2])))
	if base.Ui32(v139) <= base.Ui32(l1) {
		goto L46
	} else {
		goto L56
	}
L56:
	;
	v143 = v128 - int32(1)
	v144 = v123
	goto L52
L57:
	;
	goto L51
L58:
	;
	if l1 != int32(9) {
		v236 = v151
		goto L43
	} else {
		goto L59
	}
L59:
	;
	v246 = v147
	goto L42
L60:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+uint32(_c_F_pg_reg_getcolor[5]))))
	if v164&int32(32) != 0 {
		v236 = v151
		goto L43
	} else {
		goto L61
	}
L61:
	;
	v246 = v147
	goto L42
L62:
	;
	v209 = int32(3367)
	v210 = int32(0)
	goto L73
L63:
	;
	v176 = int32(10)
	v177 = int32(0)
	goto L64
L64:
	;
	v182 = base.I32_div_s(v176+v177, int32(2))
	v184 = v182 << (uint(int32(3)) % 32)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v184)+uint32(_c_F_pg_reg_getcolor[6])))
	if base.Ui32(v187) < base.Ui32(l1) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v254 = int32(1)
	goto L26
L66:
	;
	if v198 <= v197 {
		v176 = v197
		v177 = v198
		goto L64
	} else {
		goto L71
	}
L67:
	;
	v197 = v176
	v198 = v182 + int32(1)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v184)+uint32(_c_F_pg_reg_getcolor[7])))
	if base.Ui32(v193) <= base.Ui32(l1) {
		goto L62
	} else {
		goto L70
	}
L70:
	;
	v197 = v182 - int32(1)
	v198 = v177
	goto L66
L71:
	;
	goto L65
L72:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+uint32(_c_F_pg_reg_getcolor[3]))))
	v236 = v234
	goto L43
L73:
	;
	v215 = base.I32_div_s(v209+v210, int32(2))
	v217 = v215 * int32(12)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v217)+uint32(_c_F_pg_reg_getcolor[1])))
	if base.Ui32(v220) < base.Ui32(l1) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v236 = int32(0)
	goto L43
L75:
	;
	if v231 <= v230 {
		v209 = v230
		v210 = v231
		goto L73
	} else {
		goto L80
	}
L76:
	;
	v230 = v209
	v231 = v215 + int32(1)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v217)+uint32(_c_F_pg_reg_getcolor[2])))
	if base.Ui32(v226) <= base.Ui32(l1) {
		goto L72
	} else {
		goto L79
	}
L79:
	;
	v230 = v215 - int32(1)
	v231 = v210
	goto L75
L80:
	;
	goto L74
L81:
	;
	v306 = v59
	goto L17
L82:
	;
	if v291 != 0 {
		goto L21
	} else {
		goto L86
	}
L83:
	;
	v291 = base.B2i32(base.Ui32(int32(32)) < base.Ui32((l1+int32(1))&int32(127)))
	goto L82
L84:
	;
	goto L85
L85:
	;
	v268 = int32(_a_F_pg_reg_getcolor_3)
	v291 = base.B2i32(l1&v268 != v268)&base.B2i32(base.Ui32(l1-int32(_a_F_pg_reg_getcolor_4)) < base.Ui32(int32(_a_F_pg_reg_getcolor_5))) | (base.B2i32(base.Ui32(l1-int32(_a_F_pg_reg_getcolor_6)) < base.Ui32(int32(_a_F_pg_reg_getcolor_7))) | base.B2i32(base.Ui32(l1) < base.Ui32(int32(_a_F_pg_reg_getcolor_8))) | base.B2i32(base.Ui32(l1-int32(_a_F_pg_reg_getcolor_9)) < base.Ui32(int32(_a_F_pg_reg_getcolor_10))))
	goto L82
L86:
	;
	v306 = v59
	goto L17
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(l1-int32(32)) < base.Ui32(int32(95)) {
		goto L21
	} else {
		goto L89
	}
L89:
	;
	v306 = v59
	goto L17
L90:
	;
	goto L21
L91:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v469 == int32(0) {
		v570 = v468
		goto L139
	} else {
		goto L140
	}
L92:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	switch v311 - int32(1) {
	case 0:
		goto L96
	case 1:
		goto L94
	case 2:
		goto L95
	default:
		goto L97
	}
L93:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v468 = v466 | v306
	goto L91
L94:
	;
	if base.Ui32(int32(10)) <= base.Ui32(l1-int32(48)) {
		goto L135
	} else {
		goto L136
	}
L95:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v468 = v306
		goto L91
	} else {
		goto L131
	}
L96:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[8]))
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+16)))
	v325 = (v321 ^ int32(-1)) & int32(1)
	if base.Ui32(int32(128)) <= base.Ui32(l1) {
		goto L105
	} else {
		goto L106
	}
L97:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		v468 = v306
		goto L91
	} else {
		goto L98
	}
L98:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_pg_reg_getcolor[9]))))
	if v316&int32(3) != 0 {
		goto L93
	} else {
		goto L99
	}
L99:
	;
	v468 = v306
	goto L91
L100:
	;
	if v435 != 0 {
		goto L93
	} else {
		goto L130
	}
L101:
	;
	v435 = v428
	goto L100
L102:
	;
	v428 = base.B2i32(v417&int32(255) == int32(9))
	goto L101
L103:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+uint32(_c_F_pg_reg_getcolor[4]))))
	v417 = v410
	goto L102
L104:
	;
	v435 = base.B2i32(base.Ui32(l1-int32(48)) < base.Ui32(int32(10)))
	goto L100
L105:
	;
	v335 = int32(1178)
	v336 = int32(0)
	goto L108
L106:
	;
	goto L107
L107:
	;
	v390 = int32(1)
	v392 = l1 << (uint(v390) % 32)
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+uint32(_c_F_pg_reg_getcolor[5]))))
	if v393&v390 != 0 {
		v428 = v390
		goto L101
	} else {
		goto L128
	}
L108:
	;
	v341 = base.I32_div_s(v335+v336, int32(2))
	v343 = v341 << (uint(int32(3)) % 32)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v343)+uint32(_c_F_pg_reg_getcolor[10])))
	if base.Ui32(v346) < base.Ui32(l1) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	if v325 != 0 {
		goto L104
	} else {
		goto L118
	}
L110:
	;
	if v358 <= v357 {
		v335 = v357
		v336 = v358
		goto L108
	} else {
		goto L117
	}
L111:
	;
	v357 = v335
	v358 = v341 + int32(1)
	goto L110
L112:
	;
	goto L113
L113:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v343)+uint32(_c_F_pg_reg_getcolor[11])))
	if base.Ui32(v352) <= base.Ui32(l1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v435 = int32(1)
	goto L100
L115:
	;
	goto L116
L116:
	;
	v357 = v341 - int32(1)
	v358 = v336
	goto L110
L117:
	;
	goto L109
L118:
	;
	v364 = int32(3367)
	v365 = int32(0)
	goto L120
L119:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+uint32(_c_F_pg_reg_getcolor[3]))))
	v417 = v389
	goto L102
L120:
	;
	v370 = base.I32_div_s(v364+v365, int32(2))
	v372 = v370 * int32(12)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v372)+uint32(_c_F_pg_reg_getcolor[1])))
	if base.Ui32(v375) < base.Ui32(l1) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v417 = int32(0)
	goto L102
L122:
	;
	if v386 <= v385 {
		v364 = v385
		v365 = v386
		goto L120
	} else {
		goto L127
	}
L123:
	;
	v385 = v364
	v386 = v370 + int32(1)
	goto L122
L124:
	;
	goto L125
L125:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v372)+uint32(_c_F_pg_reg_getcolor[2])))
	if base.Ui32(v381) <= base.Ui32(l1) {
		goto L119
	} else {
		goto L126
	}
L126:
	;
	v385 = v370 - int32(1)
	v386 = v365
	goto L122
L127:
	;
	goto L121
L128:
	;
	if v325 == int32(0) {
		goto L103
	} else {
		goto L129
	}
L129:
	;
	goto L104
L130:
	;
	v468 = v306
	goto L91
L131:
	;
	goto L132
L132:
	;
	if base.B2i32(base.Ui32(l1-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(l1|int32(32)-int32(97)) < base.Ui32(int32(26))) != 0 {
		goto L93
	} else {
		goto L133
	}
L133:
	;
	v468 = v306
	goto L91
L134:
	;
	if v463 == int32(0) {
		v468 = v306
		goto L91
	} else {
		goto L138
	}
L135:
	;
	v459 = F_iswalpha(m, l1)
	mBase = m.M
	v463 = base.B2i32(v459 != int32(0))
	goto L137
L136:
	;
	v463 = int32(1)
	goto L137
L137:
	;
	goto L134
L138:
	;
	goto L93
L139:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v571 == int32(0) {
		v617 = v570
		goto L168
	} else {
		goto L169
	}
L140:
	;
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	switch v473 - int32(1) {
	case 0:
		goto L144
	case 1:
		goto L143
	case 2:
		goto L142
	default:
		goto L145
	}
L141:
	;
	if v565 == int32(0) {
		v570 = v468
		goto L139
	} else {
		goto L167
	}
L142:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v570 = v468
		goto L139
	} else {
		goto L165
	}
L143:
	;
	if base.Ui32(l1) <= base.Ui32(int32(_a_F_pg_reg_getcolor_11)) {
		goto L162
	} else {
		goto L163
	}
L144:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		v570 = v468
		goto L139
	} else {
		goto L146
	}
L146:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_pg_reg_getcolor[9]))))
	v479 = int32(1)
	v565 = int32(base.Ui32(v478)>>(uint(v479)%32)) & v479
	goto L141
L147:
	;
	v565 = v527
	goto L141
L148:
	;
	v491 = int32(1178)
	v492 = int32(0)
	goto L151
L149:
	;
	goto L150
L150:
	;
	v517 = int32(1)
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1<<(uint(v517)%32))+uint32(_c_F_pg_reg_getcolor[5]))))
	v527 = v519 & v517
	goto L147
L151:
	;
	v497 = base.I32_div_s(v491+v492, int32(2))
	v499 = v497 << (uint(int32(3)) % 32)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v499)+uint32(_c_F_pg_reg_getcolor[10])))
	if base.Ui32(v502) < base.Ui32(l1) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v527 = int32(0)
	goto L147
L153:
	;
	if v514 <= v513 {
		v491 = v513
		v492 = v514
		goto L151
	} else {
		goto L160
	}
L154:
	;
	v513 = v491
	v514 = v497 + int32(1)
	goto L153
L155:
	;
	goto L156
L156:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v499)+uint32(_c_F_pg_reg_getcolor[11])))
	if base.Ui32(v508) <= base.Ui32(l1) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v527 = int32(1)
	goto L147
L158:
	;
	goto L159
L159:
	;
	v513 = v497 - int32(1)
	v514 = v492
	goto L153
L160:
	;
	goto L152
L161:
	;
	v565 = v551
	goto L141
L162:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l1)>>(uint(int32(8))%32)))+uint32(_c_F_pg_reg_getcolor[12]))))
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l1)>>(uint(int32(3))%32))&int32(31)|v539<<(uint(int32(5))%32))+uint32(_c_F_pg_reg_getcolor[12]))))
	v551 = int32(base.Ui32(v543)>>(uint(l1&int32(7))%32)) & int32(1)
	goto L161
L163:
	;
	goto L164
L164:
	;
	v551 = base.B2i32(base.Ui32(l1) < base.Ui32(int32(_a_F_pg_reg_getcolor_12)))
	goto L161
L165:
	;
	goto L166
L166:
	;
	v565 = base.B2i32(base.B2i32(base.Ui32(l1|int32(32)-int32(97)) < base.Ui32(int32(26))) != int32(0))
	goto L141
L167:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v570 = v568 | v468
	goto L139
L168:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v618 == int32(0) {
		v711 = v617
		goto L180
	} else {
		goto L181
	}
L169:
	;
	if l1 == int32(95) {
		v610 = int32(1)
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v612 == int32(0) {
		v617 = v570
		goto L168
	} else {
		goto L179
	}
L171:
	;
	v612 = v610
	goto L170
L172:
	;
	v578 = int32(0)
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	switch v580 - int32(1) {
	case 0:
		goto L175
	case 1:
		goto L174
	case 2:
		goto L173
	default:
		goto L176
	}
L173:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v610 = v578
		goto L171
	} else {
		goto L178
	}
L174:
	;
	v601 = F_iswalnum(m, l1)
	mBase = m.M
	v612 = v601
	goto L170
L175:
	;
	v591 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[8]))
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591)+16)))
	v597 = F_pg_u_isalnum(m, l1, (v592^int32(-1))&int32(1))
	mBase = m.M
	v612 = v597
	goto L170
L176:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		v610 = v578
		goto L171
	} else {
		goto L177
	}
L177:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_pg_reg_getcolor[9]))))
	v612 = base.B2i32(v585&int32(3) != int32(0))
	goto L170
L178:
	;
	v607 = F_isalnum(m, l1)
	mBase = m.M
	v610 = base.B2i32(v607 != int32(0))
	goto L171
L179:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v617 = v615 | v570
	goto L168
L180:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v712 == int32(0) {
		v912 = v711
		goto L211
	} else {
		goto L212
	}
L181:
	;
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	switch v622 - int32(1) {
	case 0:
		goto L186
	case 1:
		goto L185
	case 2:
		goto L184
	default:
		goto L183
	}
L182:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v711 = v709 | v617
	goto L180
L183:
	;
	if base.Ui32(l1-int32(58)) < base.Ui32(int32(-10)) {
		v711 = v617
		goto L180
	} else {
		goto L210
	}
L184:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v711 = v617
		goto L180
	} else {
		goto L207
	}
L185:
	;
	goto L205
L186:
	;
	v626 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[8]))
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+16)))
	if (v627^int32(-1))&int32(1) != 0 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	if v688 != 0 {
		goto L182
	} else {
		goto L204
	}
L188:
	;
	v688 = base.B2i32(base.Ui32(l1-int32(48)) < base.Ui32(int32(10)))
	goto L187
L189:
	;
	goto L190
L190:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	v688 = base.B2i32(v678&int32(255) == int32(9))
	goto L187
L192:
	;
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652)+uint32(_c_F_pg_reg_getcolor[3]))))
	v678 = v672
	goto L191
L193:
	;
	v644 = int32(0)
	v645 = int32(3367)
	goto L196
L194:
	;
	goto L195
L195:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_c_F_pg_reg_getcolor[4]))))
	v678 = v671
	goto L191
L196:
	;
	v650 = base.I32_div_s(v644+v645, int32(2))
	v652 = v650 * int32(12)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v652)+uint32(_c_F_pg_reg_getcolor[1])))
	if base.Ui32(v655) < base.Ui32(l1) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v678 = int32(0)
	goto L191
L198:
	;
	if v665 <= v666 {
		v644 = v665
		v645 = v666
		goto L196
	} else {
		goto L203
	}
L199:
	;
	v665 = v650 + int32(1)
	v666 = v645
	goto L198
L200:
	;
	goto L201
L201:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v652)+uint32(_c_F_pg_reg_getcolor[2])))
	if base.Ui32(v661) <= base.Ui32(l1) {
		goto L192
	} else {
		goto L202
	}
L202:
	;
	v665 = v644
	v666 = v650 - int32(1)
	goto L198
L203:
	;
	goto L197
L204:
	;
	v711 = v617
	goto L180
L205:
	;
	if base.Ui32(l1-int32(48)) < base.Ui32(int32(10)) {
		goto L182
	} else {
		goto L206
	}
L206:
	;
	v711 = v617
	goto L180
L207:
	;
	goto L208
L208:
	;
	if base.Ui32(l1-int32(48)) < base.Ui32(int32(10)) {
		goto L182
	} else {
		goto L209
	}
L209:
	;
	v711 = v617
	goto L180
L210:
	;
	goto L182
L211:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v913 == int32(0) {
		v989 = v912
		goto L270
	} else {
		goto L271
	}
L212:
	;
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	switch v716 - int32(1) {
	case 0:
		goto L216
	case 1:
		goto L215
	case 2:
		goto L214
	default:
		goto L217
	}
L213:
	;
	if v907 == int32(0) {
		v912 = v711
		goto L211
	} else {
		goto L269
	}
L214:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v912 = v711
		goto L211
	} else {
		goto L264
	}
L215:
	;
	if base.Ui32(l1) <= base.Ui32(int32(_a_F_pg_reg_getcolor_11)) {
		goto L261
	} else {
		goto L262
	}
L216:
	;
	v727 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[8]))
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727)+16)))
	v731 = int32(1)
	if (v728^int32(-1))&v731 != 0 {
		goto L226
	} else {
		goto L227
	}
L217:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		v912 = v711
		goto L211
	} else {
		goto L218
	}
L218:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_pg_reg_getcolor[9]))))
	v907 = int32(base.Ui32(v721)>>(uint(int32(6))%32)) & int32(1)
	goto L213
L219:
	;
	v907 = v867
	goto L213
L220:
	;
	v867 = base.B2i32(v731<<(uint(v856)%32)&int32(1073217536) != int32(0))
	goto L219
L221:
	;
	v842 = int32(1)
	v843 = l1 << (uint(v842) % 32)
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v843)+uint32(_c_F_pg_reg_getcolor[5]))))
	if v844&v842 != 0 {
		goto L257
	} else {
		goto L258
	}
L222:
	;
	v867 = base.B2i32(v731<<(uint(v836)%32)&int32(821559296) != int32(0))
	goto L219
L223:
	;
	v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v809)+uint32(_c_F_pg_reg_getcolor[3]))))
	v836 = v830
	goto L222
L224:
	;
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1<<(uint(int32(1))%32))+uint32(_c_F_pg_reg_getcolor[4]))))
	v836 = v829
	goto L222
L225:
	;
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779)+uint32(_c_F_pg_reg_getcolor[3]))))
	v856 = v826
	goto L220
L226:
	;
	if base.Ui32(l1) < base.Ui32(int32(128)) {
		goto L221
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	if base.Ui32(l1) <= base.Ui32(int32(127)) {
		goto L224
	} else {
		goto L248
	}
L229:
	;
	v743 = int32(0)
	v744 = int32(1178)
	goto L230
L230:
	;
	v749 = base.I32_div_s(v743+v744, int32(2))
	v751 = v749 << (uint(int32(3)) % 32)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v751)+uint32(_c_F_pg_reg_getcolor[10])))
	if base.Ui32(v754) < base.Ui32(l1) {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	v771 = int32(0)
	v772 = int32(3367)
	goto L240
L232:
	;
	if v765 <= v766 {
		v743 = v765
		v744 = v766
		goto L230
	} else {
		goto L239
	}
L233:
	;
	v765 = v749 + int32(1)
	v766 = v744
	goto L232
L234:
	;
	goto L235
L235:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v751)+uint32(_c_F_pg_reg_getcolor[11])))
	if base.Ui32(v760) <= base.Ui32(l1) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v867 = int32(0)
	goto L219
L237:
	;
	goto L238
L238:
	;
	v765 = v743
	v766 = v749 - int32(1)
	goto L232
L239:
	;
	goto L231
L240:
	;
	v777 = base.I32_div_s(v771+v772, int32(2))
	v779 = v777 * int32(12)
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v779)+uint32(_c_F_pg_reg_getcolor[1])))
	if base.Ui32(v782) < base.Ui32(l1) {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	v856 = int32(0)
	goto L220
L242:
	;
	if v792 <= v793 {
		v771 = v792
		v772 = v793
		goto L240
	} else {
		goto L247
	}
L243:
	;
	v792 = v777 + int32(1)
	v793 = v772
	goto L242
L244:
	;
	goto L245
L245:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v779)+uint32(_c_F_pg_reg_getcolor[2])))
	if base.Ui32(v788) <= base.Ui32(l1) {
		goto L225
	} else {
		goto L246
	}
L246:
	;
	v792 = v771
	v793 = v777 - int32(1)
	goto L242
L247:
	;
	goto L241
L248:
	;
	v801 = int32(0)
	v802 = int32(3367)
	goto L249
L249:
	;
	v807 = base.I32_div_s(v801+v802, int32(2))
	v809 = v807 * int32(12)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v809)+uint32(_c_F_pg_reg_getcolor[1])))
	if base.Ui32(v812) < base.Ui32(l1) {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	v836 = int32(0)
	goto L222
L251:
	;
	if v822 <= v823 {
		v801 = v822
		v802 = v823
		goto L249
	} else {
		goto L256
	}
L252:
	;
	v822 = v807 + int32(1)
	v823 = v802
	goto L251
L253:
	;
	goto L254
L254:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v809)+uint32(_c_F_pg_reg_getcolor[2])))
	if base.Ui32(v818) <= base.Ui32(l1) {
		goto L223
	} else {
		goto L255
	}
L255:
	;
	v822 = v801
	v823 = v807 - int32(1)
	goto L251
L256:
	;
	goto L250
L257:
	;
	v867 = int32(0)
	goto L219
L258:
	;
	goto L259
L259:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v843)+uint32(_c_F_pg_reg_getcolor[4]))))
	v856 = v850
	goto L220
L260:
	;
	v907 = v890
	goto L213
L261:
	;
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l1)>>(uint(int32(8))%32)))+uint32(_c_F_pg_reg_getcolor[13]))))
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l1)>>(uint(int32(3))%32))&int32(31)|v879<<(uint(int32(5))%32))+uint32(_c_F_pg_reg_getcolor[13]))))
	v890 = int32(base.Ui32(v883)>>(uint(l1&int32(7))%32)) & int32(1)
	goto L263
L262:
	;
	v890 = int32(0)
	goto L263
L263:
	;
	goto L260
L264:
	;
	if base.Ui32(l1-int32(33)) <= base.Ui32(int32(93)) {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	v907 = base.B2i32(base.B2i32(v902 == int32(0)) != int32(0))
	goto L213
L266:
	;
	v900 = F_isalnum(m, l1)
	mBase = m.M
	v902 = v900
	goto L268
L267:
	;
	v902 = int32(1)
	goto L268
L268:
	;
	goto L265
L269:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v912 = v910 | v711
	goto L211
L270:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v990 == int32(0) {
		v1022 = v989
		goto L299
	} else {
		goto L300
	}
L271:
	;
	v917 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	switch v917 - int32(1) {
	case 0:
		goto L275
	case 1:
		goto L274
	case 2:
		goto L273
	default:
		goto L276
	}
L272:
	;
	if v984 == int32(0) {
		v989 = v912
		goto L270
	} else {
		goto L298
	}
L273:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v989 = v912
		goto L270
	} else {
		goto L296
	}
L274:
	;
	if l1 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L275:
	;
	v930 = Fn13991(m, l1, int32(5), int32(32), int32(_a_F_pg_reg_getcolor_13), int32(_a_F_pg_reg_getcolor_14), int32(10))
	mBase = m.M
	goto L278
L276:
	;
	if base.Ui32(int32(127)) < base.Ui32(l1) {
		v989 = v912
		goto L270
	} else {
		goto L277
	}
L277:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_pg_reg_getcolor[9]))))
	v984 = int32(base.Ui32(v922) >> (uint(int32(7)) % 32))
	goto L272
L278:
	;
	v984 = v930
	goto L272
L279:
	;
	v984 = v969
	goto L272
L280:
	;
	v969 = int32(0)
	goto L279
L281:
	;
	goto L282
L282:
	;
	if l1 != 0 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	v969 = base.B2i32(v962 != int32(0))
	goto L279
L284:
	;
	v942 = int32(_a_F_pg_reg_getcolor_15)
	goto L287
L285:
	;
	goto L286
L286:
	;
	v952 = int32(_a_F_pg_reg_getcolor_15)
	v953 = F_wcslen(m, v952)
	mBase = m.M
	v962 = v953<<(uint(int32(2))%32) + v952
	goto L283
L287:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v942)))
	if v945 != 0 {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	if v945 != 0 {
		goto L293
	} else {
		goto L294
	}
L289:
	;
	if l1 != v945 {
		v942 = v942 + int32(4)
		goto L287
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	goto L288
L292:
	;
	goto L291
L293:
	;
	v951 = v942
	goto L295
L294:
	;
	v951 = int32(0)
	goto L295
L295:
	;
	v962 = v951
	goto L283
L296:
	;
	goto L297
L297:
	;
	v984 = base.B2i32(base.B2i32(l1 == int32(32))|base.B2i32(base.Ui32(l1-int32(9)) < base.Ui32(int32(5))) != int32(0))
	goto L272
L298:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v989 = v987 | v912
	goto L270
L299:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v1023 == int32(0) {
		v1055 = v1022
		goto L314
	} else {
		goto L315
	}
L300:
	;
	v994 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	switch v994 - int32(1) {
	case 0:
		goto L305
	case 1:
		goto L304
	case 2:
		goto L303
	default:
		goto L302
	}
L301:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v1022 = v1020 | v989
	goto L299
L302:
	;
	if base.Ui32(l1-int32(123)) < base.Ui32(int32(-26)) {
		v1022 = v989
		goto L299
	} else {
		goto L313
	}
L303:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v1022 = v989
		goto L299
	} else {
		goto L310
	}
L304:
	;
	v1005 = F_towupper(m, l1)
	mBase = m.M
	goto L308
L305:
	;
	v1001 = Fn13990(m, l1, int32(97), int32(_a_F_pg_reg_getcolor_16), int32(_a_F_pg_reg_getcolor_17), int32(689))
	mBase = m.M
	goto L306
L306:
	;
	if v1001 != 0 {
		goto L301
	} else {
		goto L307
	}
L307:
	;
	v1022 = v989
	goto L299
L308:
	;
	if v1005 != l1 {
		goto L301
	} else {
		goto L309
	}
L309:
	;
	v1022 = v989
	goto L299
L310:
	;
	goto L311
L311:
	;
	if base.Ui32(l1-int32(97)) < base.Ui32(int32(26)) {
		goto L301
	} else {
		goto L312
	}
L312:
	;
	v1022 = v989
	goto L299
L313:
	;
	goto L301
L314:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v1056 == int32(0) {
		v1188 = v1055
		goto L329
	} else {
		goto L330
	}
L315:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	switch v1027 - int32(1) {
	case 0:
		goto L320
	case 1:
		goto L319
	case 2:
		goto L318
	default:
		goto L317
	}
L316:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v1055 = v1053 | v1022
	goto L314
L317:
	;
	if base.Ui32(l1-int32(91)) < base.Ui32(int32(-26)) {
		v1055 = v1022
		goto L314
	} else {
		goto L328
	}
L318:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v1055 = v1022
		goto L314
	} else {
		goto L325
	}
L319:
	;
	v1038 = F_towlower(m, l1)
	mBase = m.M
	goto L323
L320:
	;
	v1034 = Fn13990(m, l1, int32(65), int32(_a_F_pg_reg_getcolor_18), int32(_a_F_pg_reg_getcolor_19), int32(655))
	mBase = m.M
	goto L321
L321:
	;
	if v1034 != 0 {
		goto L316
	} else {
		goto L322
	}
L322:
	;
	v1055 = v1022
	goto L314
L323:
	;
	if v1038 != l1 {
		goto L316
	} else {
		goto L324
	}
L324:
	;
	v1055 = v1022
	goto L314
L325:
	;
	goto L326
L326:
	;
	if base.Ui32(l1-int32(65)) < base.Ui32(int32(26)) {
		goto L316
	} else {
		goto L327
	}
L327:
	;
	v1055 = v1022
	goto L314
L328:
	;
	goto L316
L329:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v1192 = int32(1)
	v1198 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1189+v1190*v51<<(uint(v1192)%32)+v1188<<(uint(v1192)%32)))))
	return v1198
L330:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, _c_F_pg_reg_getcolor[0]))
	switch v1060 - int32(1) {
	case 0:
		goto L335
	case 1:
		goto L334
	case 2:
		goto L333
	default:
		goto L332
	}
L331:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1188 = v1186 | v1055
	goto L329
L332:
	;
	if base.Ui32(l1-int32(127)) < base.Ui32(int32(-94)) {
		v1188 = v1055
		goto L329
	} else {
		goto L372
	}
L333:
	;
	if base.Ui32(int32(255)) < base.Ui32(l1) {
		v1188 = v1055
		goto L329
	} else {
		goto L369
	}
L334:
	;
	v1167 = F_iswspace(m, l1)
	mBase = m.M
	if v1167 != 0 {
		goto L365
	} else {
		goto L366
	}
L335:
	;
	if base.Ui32(int32(128)) <= base.Ui32(l1) {
		goto L341
	} else {
		goto L342
	}
L336:
	;
	if v1163 != 0 {
		goto L331
	} else {
		goto L363
	}
L337:
	;
	v1163 = int32(0)
	goto L336
L338:
	;
	v1163 = v1149
	goto L336
L339:
	;
	v1115 = int32(0)
	if int32(1)<<(uint(v1114)%32)&int32(_a_F_pg_reg_getcolor_2) != 0 {
		v1149 = v1115
		goto L338
	} else {
		goto L354
	}
L340:
	;
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1079)+uint32(_c_F_pg_reg_getcolor[3]))))
	v1114 = v1111
	goto L339
L341:
	;
	v1071 = int32(3367)
	v1072 = int32(0)
	goto L344
L342:
	;
	goto L343
L343:
	;
	v1096 = int32(1)
	v1099 = l1 << (uint(v1096) % 32)
	v1100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099)+uint32(_c_F_pg_reg_getcolor[4]))))
	if v1096<<(uint(v1100)%32)&int32(_a_F_pg_reg_getcolor_2) != 0 {
		goto L337
	} else {
		goto L352
	}
L344:
	;
	v1077 = base.I32_div_s(v1071+v1072, int32(2))
	v1079 = v1077 * int32(12)
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+uint32(_c_F_pg_reg_getcolor[1])))
	if base.Ui32(v1082) < base.Ui32(l1) {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	v1114 = int32(0)
	goto L339
L346:
	;
	if v1093 <= v1092 {
		v1071 = v1092
		v1072 = v1093
		goto L344
	} else {
		goto L351
	}
L347:
	;
	v1092 = v1071
	v1093 = v1077 + int32(1)
	goto L346
L348:
	;
	goto L349
L349:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+uint32(_c_F_pg_reg_getcolor[2])))
	if base.Ui32(v1088) <= base.Ui32(l1) {
		goto L340
	} else {
		goto L350
	}
L350:
	;
	v1092 = v1077 - int32(1)
	v1093 = v1072
	goto L346
L351:
	;
	goto L345
L352:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099)+uint32(_c_F_pg_reg_getcolor[5]))))
	if v1106&int32(32) == int32(0) {
		v1149 = v1096
		goto L338
	} else {
		goto L353
	}
L353:
	;
	goto L337
L354:
	;
	v1122 = int32(10)
	v1123 = v1115
	goto L355
L355:
	;
	v1128 = base.I32_div_s(v1122+v1123, int32(2))
	v1130 = v1128 << (uint(int32(3)) % 32)
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+uint32(_c_F_pg_reg_getcolor[6])))
	if base.Ui32(v1133) < base.Ui32(l1) {
		goto L358
	} else {
		goto L359
	}
L356:
	;
	v1149 = int32(1)
	goto L338
L357:
	;
	if v1144 <= v1143 {
		v1122 = v1143
		v1123 = v1144
		goto L355
	} else {
		goto L362
	}
L358:
	;
	v1143 = v1122
	v1144 = v1128 + int32(1)
	goto L357
L359:
	;
	goto L360
L360:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+uint32(_c_F_pg_reg_getcolor[7])))
	if base.Ui32(v1139) <= base.Ui32(l1) {
		goto L337
	} else {
		goto L361
	}
L361:
	;
	v1143 = v1128 - int32(1)
	v1144 = v1123
	goto L357
L362:
	;
	goto L356
L363:
	;
	v1188 = v1055
	goto L329
L364:
	;
	if v1172 != 0 {
		goto L331
	} else {
		goto L368
	}
L365:
	;
	v1172 = int32(0)
	goto L367
L366:
	;
	v1169 = F_iswprint(m, l1)
	mBase = m.M
	v1172 = base.B2i32(v1169 != int32(0))
	goto L367
L367:
	;
	goto L364
L368:
	;
	v1188 = v1055
	goto L329
L369:
	;
	goto L370
L370:
	;
	if base.Ui32(l1-int32(33)) < base.Ui32(int32(94)) {
		goto L331
	} else {
		goto L371
	}
L371:
	;
	v1188 = v1055
	goto L329
L372:
	;
	goto L331
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
func F_pg_sequence_parameters(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_pg_sequence_parameters[0]))
	v13 = F_pg_class_aclcheck(m, v9, v11, int64(262))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			v22 = F_get_call_result_type(m, l0, int32(0), v5+int32(-4))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pg_sequence_parameters_0), int32(0))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_sequence_parameters_1), int32(1757), int32(_a_F_pg_sequence_parameters_2))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v26 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+27)) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v26
					v31 = F_SearchSysCache1(m, int32(61), v9)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
								F_errmsg_internal(m, int32(_a_F_pg_sequence_parameters_3), v7)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_sequence_parameters_1), int32(1763), int32(_a_F_pg_sequence_parameters_2))
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
							v37 = v35 + v36
							v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+8))
							v39 = F_Int64GetDatum(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v39
								v42 = *(*int64)(unsafe.Add(mBase, uint32(v37)+32))
								v43 = F_Int64GetDatum(m, v42)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v43
									v46 = *(*int64)(unsafe.Add(mBase, uint32(v37)+24))
									v47 = F_Int64GetDatum(m, v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v47
										v50 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
										v51 = F_Int64GetDatum(m, v50)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v51
											v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+48)))
											*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v54
											v56 = *(*int64)(unsafe.Add(mBase, uint32(v37)+40))
											v57 = F_Int64GetDatum(m, v56)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v57
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v60
												F_ReleaseCatCache(m, v31)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int32(0)
												} else {
													v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+60))
													v69 = F_heap_form_tuple(m, v64, v5+int32(-32), v5+int32(-40))
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
														v72 = F_HeapTupleHeaderGetDatum(m, v71)
														mBase = m.M
														v73 = m.ExcPending
														if v73 != 0 {
															return int32(0)
														} else {
															m.G0 = v7 - int32(-64)
															return v72
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					v85 = F_get_rel_name(m, v9)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v85
						F_errmsg(m, int32(_a_F_pg_sequence_parameters_4), v5+int32(-48))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_sequence_parameters_1), int32(1754), int32(_a_F_pg_sequence_parameters_2))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
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
func F_pg_size_pretty_numeric(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var __phi25 int32
	_ = __phi25
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
	var v32 int64
	_ = v32
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	__phi24 = int32(_a_F_pg_size_pretty_numeric_0)
	__phi25 = v19
	__phi27 = int32(_a_F_pg_size_pretty_numeric_1)
	__phi28 = int32(_a_F_pg_size_pretty_numeric_2)
	__phi30 = int32(_a_F_pg_size_pretty_numeric_3)
	v24 = __phi24
	v25 = __phi25
	v27 = __phi27
	v28 = __phi28
	v30 = __phi30
	goto L3
L3:
	;
	v32 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24)+4)))
	v34 = int32(0)
	v37 = F_DirectFunctionCall1Coll(m, int32(1276), v34, v25)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+8)))
	if v74 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	goto L4
L6:
	;
	v39 = F_pg_detoast_datum(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v41 = F_int64_to_numeric(m, v32)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v43 = F_DirectFunctionCall2Coll(m, int32(1275), v34, v39, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v65 = v24
	v67 = v25
	v70 = v27
	goto L5
L11:
	;
	goto L12
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v51 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
	v52 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+21)))
	v53 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
	v54 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+20)))
	v59 = F_int64_to_numeric(m, int64(1)<<(uint(v51+(v52-(v53+v54)))%64))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v61 = F_DirectFunctionCall2Coll(m, int32(1277), int32(0), v25, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v63 = F_pg_detoast_datum(m, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v47 != 0 {
		__phi24 = v30
		__phi25 = v63
		__phi27 = v28
		__phi28 = v47
		__phi30 = v30 + int32(12)
		v24 = __phi24
		v25 = __phi25
		v27 = __phi27
		v28 = __phi28
		v30 = __phi30
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v65 = v30
	v67 = v63
	v70 = v28
	goto L5
L17:
	;
	v78 = F_int64_to_numeric(m, int64(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v105 = v67
	goto L19
L19:
	;
	v106 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L30
	}
L20:
	;
	v81 = F_int64_to_numeric(m, int64(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v84 = F_int64_to_numeric(m, int64(2))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v87 = int32(0)
	v92 = F_DirectFunctionCall2Coll(m, int32(1279), v87, v67, v78)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v92 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v94 = int32(1278)
	goto L26
L25:
	;
	v94 = int32(18)
	goto L26
L26:
	;
	v96 = F_DirectFunctionCall2Coll(m, v94, int32(0), v67, v81)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v98 = F_DirectFunctionCall2Coll(m, int32(1277), v87, v96, v84)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v100 = F_pg_detoast_datum(m, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v105 = v100
	goto L19
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v106
	v111 = F_psprintf(m, int32(_a_F_pg_size_pretty_numeric_4), v12)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v113 = F_cstring_to_text(m, v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	m.G0 = v12 + int32(16)
	return v113
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
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
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
			v37 = v5
		} else {
			if v19 == v16 {
				*(*int32)(unsafe.Add(mBase, _c_F_pg_strftime[0])) = int32(68)
				v37 = v5
			} else {
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, _c_F_pg_strftime[0])) = v13
				v37 = v19 - l0
			}
		}
		m.G0 = v10 + int32(16)
		return v37
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
	v401 = int32(612)
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
	v401 = int32(618)
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
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = m.T0[v7].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l2, l1, int32(-1), l3)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_pg_table_size(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_try_relation_open(m, v4, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int32(0)
		} else {
			v16 = F_calculate_table_size(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_relation_close(m, v6, int32(1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = F_Int64GetDatum(m, v16)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						return v21
					}
				}
			}
		}
	}
}
func F_pg_timezone_abbrevs_zone(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v189 int64
	_ = v189
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int64
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int64
	_ = v343
	var v348 int64
	_ = v348
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int64
	_ = v358
	var v359 int64
	_ = v359
	var v362 int64
	_ = v362
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int64
	_ = v390
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+50)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+48)) = uint16(v2)
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[0]))
	v17 = base.I64_div_s(v15, int64(1000000))
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v17 + int64(946684800)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v22 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L6
	} else {
		goto L93
	}
L3:
	;
	v25 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	goto L12
L6:
	;
	return int32(0)
L7:
	;
	v29 = int32(_a_F_pg_timezone_abbrevs_zone_0)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[1]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[1])) = v32
	v35 = F_palloc(m, int32(4))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v35
	v43 = F_get_call_result_type(m, l0, v37, v6+int32(-56))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v43 != int32(1) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v47
	*(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[1])) = v30
	goto L5
L11:
	;
	m.G0 = v8 - int32(-64)
	return v433
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[2]))
	v59 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v61 < v59 {
		v77 = v59
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v77 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L13
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+268))
	if v64 <= v61 {
		v77 = v59
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v67 = int32(_a_F_pg_timezone_abbrevs_zone_1)
	v69 = F_strlen(m, v58+v61+v67)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v69 + v61 + int32(1)
	v77 = v58 + v67 + v61
	goto L14
L17:
	;
	v80 = v77
	goto L20
L18:
	;
	goto L19
L19:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L6
	} else {
		goto L92
	}
L20:
	;
	v83 = int32(_a_F_pg_timezone_abbrevs_zone_2)
	v87 = m.G0
	v89 = v87 - int32(32)
	v90 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+24)) = v90
	*(*int64)(unsafe.Add(mBase, uint32(v89)+16)) = v90
	*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v90
	*(*int64)(unsafe.Add(mBase, uint32(v89))) = v90
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[3])))
	if v98 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L19
L22:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[2]))
	v399 = int32(0)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v401 < v399 {
		v417 = v399
		goto L88
	} else {
		goto L89
	}
L23:
	;
	v167 = F_strlen(m, v80)
	mBase = m.M
	if v166 != v167 {
		goto L22
	} else {
		goto L42
	}
L24:
	;
	v166 = int32(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[4])))
	if v102 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v106 = v80
	goto L30
L28:
	;
	goto L29
L29:
	;
	v116 = v83
	v117 = v98
	goto L33
L30:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v112 == v98 {
		v106 = v106 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v166 = v106 - v80
	goto L23
L32:
	;
	goto L31
L33:
	;
	v124 = v89 + int32(base.Ui32(v117)>>(uint(int32(3))%32))&int32(28)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v126 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v125 | v126<<(uint(v117)%32)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v130 != 0 {
		v116 = v116 + v126
		v117 = v130
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v133 == int32(0) {
		v156 = v80
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v166 = v156 - v80
	goto L23
L37:
	;
	v137 = v80
	v138 = v133
	goto L38
L38:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v89+int32(base.Ui32(v138)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v146)>>(uint(v138)%32))&int32(1) == int32(0) {
		v156 = v137
		goto L36
	} else {
		goto L40
	}
L39:
	;
	v156 = v154
	goto L36
L40:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	v154 = v137 + int32(1)
	if v152 != 0 {
		v137 = v154
		v138 = v152
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_zone[2]))
	v177 = int32(0)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)+268))
	if v184 <= v177 {
		v337 = v177
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v337 == int32(0) {
		goto L22
	} else {
		goto L78
	}
L44:
	;
	goto L43
L45:
	;
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v6+int32(-24))))
	v192 = int32(0)
	goto L46
L46:
	;
	v203 = v192 + (v176 + int32(_a_F_pg_timezone_abbrevs_zone_1))
	v204 = F_strcmp(m, v80, v203)
	mBase = m.M
	if v204 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v176)+260))
	if v210 <= int32(0) {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v205 = F_strlen(m, v203)
	mBase = m.M
	v208 = v205 + v192 + int32(1)
	if v208 < v184 {
		v192 = v208
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
	v337 = v177
	goto L44
L52:
	;
	v255 = v176 + int32(_a_F_pg_timezone_abbrevs_zone_3)
	v257 = v176 + int32(_a_F_pg_timezone_abbrevs_zone_4)
	v258 = v247
	goto L66
L53:
	;
	v247 = int32(0)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v217 = v210
	v222 = int32(0)
	goto L56
L56:
	;
	v230 = int32(1)
	v231 = (v217 + v222) >> (uint(v230) % 32)
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v176+int32(280)+v231<<(uint(int32(3))%32))))
	v238 = base.B2i32(v189 < v237)
	if v189 < v237 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v247 = v239
	goto L52
L58:
	;
	v239 = v222
	goto L60
L59:
	;
	v239 = v231 + v230
	goto L60
L60:
	;
	if v189 < v237 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v240 = v231
	goto L63
L62:
	;
	v240 = v217
	goto L63
L63:
	;
	if v239 < v240 {
		v217 = v240
		v222 = v239
		goto L56
	} else {
		goto L64
	}
L64:
	;
	goto L57
L65:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	*(*int32)(unsafe.Add(mBase, uint32(v6+int32(-28)))) = v322
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v6+int32(-32)))) = v324
	v337 = int32(1)
	goto L44
L66:
	;
	if int32(0) < v258 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v176)+uint32(_c_F_pg_timezone_abbrevs_zone[5])))
	v284 = v257 + v281<<(uint(int32(4))%32)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	if v285 == v192 {
		v316 = v284
		goto L65
	} else {
		goto L72
	}
L68:
	;
	v273 = v258 - int32(1)
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255+v273))))
	v278 = v257 + v275<<(uint(int32(4))%32)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+8))
	if v279 != v192 {
		v258 = v273
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
	v316 = v278
	goto L65
L72:
	;
	if v210 <= v247 {
		v337 = v177
		goto L44
	} else {
		goto L73
	}
L73:
	;
	v293 = v247
	goto L74
L74:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293+v255))))
	v304 = v257 + v301<<(uint(int32(4))%32)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+8))
	if v305 == v192 {
		v316 = v304
		goto L65
	} else {
		goto L76
	}
L75:
	;
	v337 = v177
	goto L44
L76:
	;
	v308 = v293 + int32(1)
	if v210 != v308 {
		v293 = v308
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v341 = F_cstring_to_text(m, v80)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	v343 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v341
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v343
	v348 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+36)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v348 * int64(1000000)
	v353 = v6 + int32(-56)
	v355 = F_palloc(m, int32(16))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	v358 = int64(*(*int32)(unsafe.Add(mBase, uint32(v353)+12)))
	v359 = int64(*(*int32)(unsafe.Add(mBase, uint32(v353)+16)))
	v362 = v358 + v359*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v362-int64(2147483648)) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = v355
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+60)) = base.B2i32(v376 != int32(0))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	v385 = F_heap_form_tuple(m, v380, v6+int32(-12), v6+int32(-16))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L6
	} else {
		goto L85
	}
L82:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v355)+12)) = uint32(v362)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v353)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v355)+8)) = v368
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v353)))
	*(*int64)(unsafe.Add(mBase, uint32(v355))) = v370
	goto L84
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v385)+16))
	v388 = F_HeapTupleHeaderGetDatum(m, v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	v390 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v390 + int64(1)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v394)+20)) = int32(1)
	v433 = v388
	goto L11
L87:
	;
	if v417 != 0 {
		v80 = v417
		goto L20
	} else {
		goto L91
	}
L88:
	;
	goto L87
L89:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v398)+268))
	if v404 <= v401 {
		v417 = v399
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v407 = int32(_a_F_pg_timezone_abbrevs_zone_1)
	v409 = F_strlen(m, v398+v401+v407)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v409 + v401 + int32(1)
	v417 = v398 + v407 + v401
	goto L88
L91:
	;
	goto L21
L92:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v425)+20)) = int32(2)
	v428 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v428)
	v433 = int32(0)
	goto L11
L93:
	;
	F_errmsg_internal(m, int32(_a_F_pg_timezone_abbrevs_zone_5), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_pg_timezone_abbrevs_zone_6), int32(_a_F_pg_timezone_abbrevs_zone_7), int32(_a_F_pg_timezone_abbrevs_zone_8))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
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
func F_pg_trigger_depth(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_trigger_depth[0]))
	return v3
}
func F_pg_ts_dict_is_visible(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_TSDictionaryIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
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
	*(*int64)(unsafe.Add(mBase, uint32(v8)+528)) = int64(102873056674048)
	v25 = F_hash_create(m, int32(_a_F_pg_tzset_1), int32(4), v8+int32(512), int32(24))
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
	v94 = int32(333)
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
	v130 = int32(121)
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
func F_pg_walfile_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_walfile_name[0])))
	if v14 == int32(1) {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name[1]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+316))
		v22 = base.B2i32(v20 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_walfile_name[0])) = uint8(v22)
		v24 = v22
	} else {
		v24 = int32(0)
	}
	if v24 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_walfile_name_0), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_pg_walfile_name_1)
					F_errhint(m, int32(_a_F_pg_walfile_name_2), v8)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_walfile_name_3), int32(449), int32(_a_F_pg_walfile_name_4))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
	} else {
		v49 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name[2])))
		v51 = *(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name[1]))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+308))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v52
		v54 = base.I64_div_u_s(v11, v49)
		v57 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name[2])))
		v58 = base.I64_div_u_s(int64(4294967296), v57)
		v59 = base.I64_div_u_s(v54, v58)
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+20)) = uint32(v59)
		v62 = v54 - v58*v59
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+24)) = uint32(v62)
		v65 = v8 + int32(32)
		v70 = F_pg_snprintf(m, v65, int32(64), int32(_a_F_pg_walfile_name_5), v8+int32(16))
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int32(0)
		} else {
			v72 = F_cstring_to_text(m, v65)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(96)
				return v72
			}
		}
	}
}
func F_pg_walfile_name_offset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
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
	v8 = m.G0
	v10 = v8 - int32(112)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[0])))
	if v16 == int32(1) {
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[1]))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+316))
		v24 = base.B2i32(v22 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[0])) = uint8(v24)
		v26 = v24
	} else {
		v26 = int32(0)
	}
	if v26 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_walfile_name_offset_0), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_pg_walfile_name_offset_1)
					F_errhint(m, int32(_a_F_pg_walfile_name_offset_2), v10)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_walfile_name_offset_3), int32(391), int32(_a_F_pg_walfile_name_offset_4))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
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
	} else {
		v51 = F_CreateTemplateTupleDesc(m, int32(2))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			F_TupleDescInitEntry(m, v51, int32(1), int32(_a_F_pg_walfile_name_offset_5), int32(25), int32(-1), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				F_TupleDescInitEntry(m, v51, int32(2), int32(_a_F_pg_walfile_name_offset_6), int32(23), int32(-1), int32(0))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					v67 = F_BlessTupleDesc(m, v51)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v70 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[2])))
						v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[1]))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+308))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v73
						v75 = base.I64_div_u_s(v13, v70)
						v78 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[2])))
						v79 = base.I64_div_u_s(int64(4294967296), v78)
						v80 = base.I64_div_u_s(v75, v79)
						*(*uint32)(unsafe.Add(mBase, uint32(v10)+20)) = uint32(v80)
						v83 = v75 - v79*v80
						*(*uint32)(unsafe.Add(mBase, uint32(v10)+24)) = uint32(v83)
						v86 = v10 + int32(48)
						v91 = F_pg_snprintf(m, v86, int32(64), int32(_a_F_pg_walfile_name_offset_7), v10+int32(16))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							v93 = F_cstring_to_text(m, v86)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v93
								v96 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v10)+38)) = uint16(v96)
								v100 = *(*int32)(unsafe.Add(mBase, _c_F_pg_walfile_name_offset[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = base.I32_wrap_i64(v13) & (v100 - int32(1))
								v109 = F_heap_form_tuple(m, v67, v10+int32(40), v10+int32(38))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return int32(0)
								} else {
									v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
									v112 = F_HeapTupleHeaderGetDatum(m, v111)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(112)
										return v112
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
