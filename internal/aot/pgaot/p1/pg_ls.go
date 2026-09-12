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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
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
		goto L3
	}
L3:
	;
	if l2 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v9 + int32(2192)
	return
L5:
	;
	v23 = F_ReadDir(m, v15, l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	if v15 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v20 == int32(44) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v27 = v23
	goto L13
L11:
	;
	goto L12
L12:
	;
	F_FreeDir(m, v15)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L35
	}
L13:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+19)))
	if v33 == int32(46) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v111 = F_ReadDir(m, v15, l1)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L33
	}
L16:
	;
	v37 = v27 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
	v46 = F_pg_snprintf(m, v9+int32(128), int32(2048), int32(168545), v9+int32(16))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v54 = F___fstatat(m, int32(-100), v9+int32(128), v9+int32(32), int32(0))
	mBase = m.M
	goto L18
L18:
	;
	if v54 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v58 == int32(44) {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v78&int32(61440) != int32(32768) {
		goto L15
	} else {
		goto L27
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(128)
	F_errmsg(m, int32(285574), v9)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(479690), int32(612), int32(156524))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
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
	v83 = F_cstring_to_text(m, v37)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+2180)) = v83
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v9)+56))
	v87 = F_Int64GetDatum(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+2184)) = v87
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v9)+88))
	goto L30
L30:
	;
	v95 = F_Int64GetDatum(m, v90*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v97 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9+int32(2178)))) = uint8(v97)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+2176)) = uint16(v97)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+2188)) = v95
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v102, v103, v9+int32(2180), v9+int32(2176))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L15
L33:
	;
	if v111 != 0 {
		v27 = v111
		goto L13
	} else {
		goto L34
	}
L34:
	;
	goto L14
L35:
	;
	goto L4
}
func F_pg_ls_logdir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1077]))
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
	F_pg_ls_dir_files(m, l0, int32(295996), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
