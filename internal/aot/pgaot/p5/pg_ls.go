package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_ls_dir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = F_convert_and_check_filename(m, v14)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v20 != int32(3) {
		v34 = v2
		v35 = v2
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_InitMaterializedSRF(m, l0, int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v23 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v29 = base.B2i32(v26 != int32(0))
	goto L8
L7:
	;
	v29 = v2
	goto L8
L8:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v30 != 0 {
		v34 = v29
		v35 = v2
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v34 = v29
	v35 = base.B2i32(v31 != int32(0))
	goto L4
L10:
	;
	v39 = F_AllocateDir(m, v18)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	m.G0 = v10 + int32(16)
	return int32(0)
L12:
	;
	v41 = int32(0)
	if v39|base.B2i32(v34 == v41) == v41 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ls_dir[0]))
	if v47 == int32(44) {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v50 = F_ReadDir(m, v39, v18)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	if v50 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v52 = v50
	goto L21
L19:
	;
	goto L20
L20:
	;
	F_FreeDir(m, v39)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L34
	}
L21:
	;
	if v35 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L20
L23:
	;
	v87 = F_ReadDir(m, v39, v18)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L32
	}
L24:
	;
	v73 = F_cstring_to_text(m, v52+int32(19))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L30
	}
L25:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+19)))
	if v59 != int32(46) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+20)))
	if v62 == int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+20)))
	if v65 != int32(46) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+21)))
	if v68 == int32(0) {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	v75 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)) = uint8(v75)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v73
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_tuplestore_putvalues(m, v78, v79, v10+int32(12), v10+int32(11))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L23
L32:
	;
	if v87 != 0 {
		v52 = v87
		goto L21
	} else {
		goto L33
	}
L33:
	;
	goto L22
L34:
	;
	goto L11
}
func F_pg_ls_dir_1arg(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_pg_ls_dir(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_ls_summariesdir(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_pg_ls_dir_files(m, l0, int32(_a_F_pg_ls_summariesdir_0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_ls_tmpdir_noargs(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	F_pg_ls_tmpdir(m, l0, int32(1663))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
