package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_show_all_file_settings(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = F_ProcessConfigFileInternal(m, int32(2), int32(0), int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = int32(1)
	v30 = v16
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v10 + int32(48)
	return int32(0)
L7:
	;
	v35 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10+int32(32)))) = v35
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(40)))) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v10)+11)) = v37
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v47 == v37 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v28
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v63 != 0 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v60 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+9)) = uint8(v60)
	goto L9
L11:
	;
	v50 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+8)) = uint8(v50)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v52 = F_cstring_to_text(m, v47)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	if v55 == int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v58
	goto L9
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v69 != 0 {
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
	v67 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)) = uint8(v67)
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v64
	goto L16
L21:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+21)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v77 != 0 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v70 = F_cstring_to_text(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v73 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v73)
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v70
	goto L21
L26:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_tuplestore_putvalues(m, v83, v84, v10+int32(16), v10+int32(8))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L31
	}
L27:
	;
	v78 = F_cstring_to_text(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v81)
	goto L26
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v78
	goto L26
L31:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	if v93 != 0 {
		v28 = v28 + int32(1)
		v30 = v93
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L8
}
