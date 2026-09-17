package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_ls_dir_files(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
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
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v118 int32
	_ = v118
	v7 = m.G0
	v9 = v7 - int32(2192)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
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
	return
L2:
	;
	v15 = F_AllocateDir(m, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v9 + int32(2192)
	return
L4:
	;
	v17 = int32(0)
	if v15|base.B2i32(l2 == v17) == v17 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ls_dir_files[0]))
	if v23 == int32(44) {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v26 = F_ReadDir(m, v15, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	if v26 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v28 = v26
	goto L13
L11:
	;
	goto L12
L12:
	;
	F_FreeDir(m, v15)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L35
	}
L13:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+19)))
	if v34 == int32(46) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v109 = F_ReadDir(m, v15, l1)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L33
	}
L16:
	;
	v38 = v28 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	v42 = v9 + int32(128)
	v47 = F_pg_snprintf(m, v42, int32(2048), int32(_a_F_pg_ls_dir_files_0), v9+int32(16))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v53 = F___fstatat(m, int32(-100), v42, v9+int32(32), int32(0))
	mBase = m.M
	goto L18
L18:
	;
	if v53 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ls_dir_files[0]))
	if v57 == int32(44) {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v75&int32(_a_F_pg_ls_dir_files_1) != int32(_a_F_pg_ls_dir_files_2) {
		goto L15
	} else {
		goto L27
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v42
	F_errmsg(m, int32(_a_F_pg_ls_dir_files_3), v9)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_pg_ls_dir_files_4), int32(612), int32(_a_F_pg_ls_dir_files_5))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
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
	v80 = F_cstring_to_text(m, v38)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+2180)) = v80
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v9)+56))
	v84 = F_Int64GetDatum(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+2184)) = v84
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v9)+88))
	goto L30
L30:
	;
	v92 = F_Int64GetDatum(m, v87*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v94 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+2176)) = uint16(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+2188)) = v92
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+2178)) = uint8(v94)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v99, v100, v9+int32(2180), v9+int32(2176))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L15
L33:
	;
	if v109 != 0 {
		v28 = v109
		goto L13
	} else {
		goto L34
	}
L34:
	;
	goto L14
L35:
	;
	goto L3
}
func F_pg_ls_logdir(m *base.Module, l0 int32) int32 {
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
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_ls_waldir(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_pg_ls_dir_files(m, l0, int32(_a_F_pg_ls_waldir_0), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
