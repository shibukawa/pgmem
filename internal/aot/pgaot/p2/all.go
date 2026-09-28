package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_show_all_file_settings(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = F_ProcessConfigFileInternal(m, int32(2), int32(0), int32(12))
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
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = int32(1)
	v24 = v14
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v8 + int32(80)
	return int64(0)
L7:
	;
	v27 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v27
	v41 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v8)+11)) = v41
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v45 == v41 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L6
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = base.I64_extend_i32_s(v22)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v63 != 0 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+9)) = uint8(v59)
	goto L9
L11:
	;
	v48 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+8)) = uint8(v48)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v50 = F_cstring_to_text(m, v45)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = base.I64_extend_i32_u(v50)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v54 == int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v57 = int64(*(*int32)(unsafe.Add(mBase, uint32(v24)+16)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v57
	goto L9
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v70 != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v64 = F_cstring_to_text(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v68 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v68)
	goto L16
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = base.I64_extend_i32_u(v64)
	goto L16
L21:
	;
	v77 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v24)+21)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v79 != 0 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v71 = F_cstring_to_text(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v75 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v75)
	goto L21
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = base.I64_extend_i32_u(v71)
	goto L21
L26:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_tuplestore_putvalues(m, v86, v87, v8+int32(16), v8+int32(8))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L31
	}
L27:
	;
	v80 = F_cstring_to_text(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v84 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v84)
	goto L26
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = base.I64_extend_i32_u(v80)
	goto L26
L31:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	if v96 != 0 {
		v22 = v22 + int32(1)
		v24 = v96
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L8
}
