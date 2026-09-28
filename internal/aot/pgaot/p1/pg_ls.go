package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_ls_dir_files(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	v8 = m.G0
	v10 = v8 - int32(2224)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = F_AllocateDir(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v10 + int32(2224)
	return
L4:
	;
	v18 = int32(0)
	if v16|base.B2i32(l2 == v18) == v18 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ls_dir_files[0]))
	if v24 == int32(44) {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v27 = F_ReadDir(m, v16, l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v29 = v27
	goto L13
L11:
	;
	goto L12
L12:
	;
	F_FreeDir(m, v16)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L33
	}
L13:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+19)))
	if v36 == int32(46) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v109 = F_ReadDir(m, v16, l1)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L31
	}
L16:
	;
	v40 = v29 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
	v44 = v10 + int32(128)
	v49 = F_pg_snprintf(m, v44, int32(2048), int32(_a_F_pg_ls_dir_files_0), v10+int32(16))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v55 = F___fstatat(m, int32(-100), v44, v10+int32(32), int32(0))
	mBase = m.M
	goto L18
L18:
	;
	if v55 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ls_dir_files[0]))
	if v59 == int32(44) {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	if v77&int32(_a_F_pg_ls_dir_files_1) != int32(_a_F_pg_ls_dir_files_2) {
		goto L15
	} else {
		goto L27
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v44
	F_errmsg(m, int32(_a_F_pg_ls_dir_files_3), v10)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_pg_ls_dir_files_4), int32(613), int32(_a_F_pg_ls_dir_files_5))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v82 = F_cstring_to_text(m, v40)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+2192)) = base.I64_extend_i32_u(v82)
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v10)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+2200)) = v86
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v10)+88))
	goto L29
L29:
	;
	v93 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+2188)) = uint16(v93)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+2208)) = v88*int64(1000000) - int64(946684800000000)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+2190)) = uint8(v93)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_tuplestore_putvalues(m, v98, v99, v10+int32(2192), v10+int32(2188))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L15
L31:
	;
	if v109 != 0 {
		v29 = v109
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L14
L33:
	;
	goto L3
}
func F_pg_ls_logdir(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ls_logdir[0]))
	F_pg_ls_dir_files(m, l0, v3, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_ls_waldir(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	F_pg_ls_dir_files(m, l0, int32(_a_F_pg_ls_waldir_0), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
