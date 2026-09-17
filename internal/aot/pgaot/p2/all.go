package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_show_all_file_settings(m *base.Module, l0 int32) int32 {
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
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	v6 = m.G0
	v8 = v6 - int32(48)
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
	return int32(0)
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
	m.G0 = v8 + int32(48)
	return int32(0)
L7:
	;
	v27 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v27
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v8)+11)) = v33
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v39 == v33 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v22
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v52 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+9)) = uint8(v52)
	goto L9
L11:
	;
	v42 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+8)) = uint8(v42)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v44 = F_cstring_to_text(m, v39)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v44
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v47 == int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v50
	goto L9
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v61 != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v56 = F_cstring_to_text(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v59)
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v56
	goto L16
L21:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+21)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v69 != 0 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v62 = F_cstring_to_text(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v65 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+12)) = uint8(v65)
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v62
	goto L21
L26:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_tuplestore_putvalues(m, v75, v76, v8+int32(16), v8+int32(8))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L31
	}
L27:
	;
	v70 = F_cstring_to_text(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v73 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)) = uint8(v73)
	goto L26
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v70
	goto L26
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	if v85 != 0 {
		v22 = v22 + int32(1)
		v24 = v85
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L8
}
