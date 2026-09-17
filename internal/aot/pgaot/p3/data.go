package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DataChecksumsEnabled(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_DataChecksumsEnabled[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+252))
	return base.B2i32(v3 != int32(0))
}
func F_copy_read_data(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_copy_read_data[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v15 != v16 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = v15 - v16
	if v18 < l2 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v28 = l2
	v29 = int32(0)
	goto L3
L3:
	;
	if base.B2i32(v28 <= int32(0))|base.B2i32(l1 <= v29) != 0 {
		v108 = v29
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v20 = v18
	goto L6
L5:
	;
	v20 = l2
	goto L6
L6:
	;
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	base.MemoryCopy(m, l0, v21+v16, v20)
	goto L9
L8:
	;
	goto L9
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v24 + v20
	v28 = l2 - v20
	v29 = v20
	goto L3
L10:
	;
	m.G0 = v11 + int32(16)
	return v108
L11:
	;
	v34 = l0
	v36 = v28
	v38 = v29
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(0)
	v46 = v34
	v48 = v36
	v50 = v38
	goto L14
L13:
	;
	v108 = v50
	goto L10
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_copy_read_data[1]))
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_copy_read_data[2]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
	v63 = m.T0[v62].(func(*base.Module, int32, int32, int32) int32)(m, v55, v11+int32(8), v11+int32(12))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_copy_read_data[3]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v97 = F_WaitLatchOrSocket(m, v93, v94, int32(1000), int32(134217759))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L16
	} else {
		goto L34
	}
L16:
	;
	return int32(0)
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_copy_read_data[4]))
	if v68 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v63 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	if v63 < int32(0) {
		v108 = v50
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L15
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_copy_read_data[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v73
	if v63 < v48 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v81 = v63
	goto L28
L27:
	;
	v81 = v48
	goto L28
L28:
	;
	if v81 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	base.MemoryCopy(m, v46, v73, v81)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v83 + v81
	v86 = v81 + v50
	v87 = v48 - v81
	if v87 <= int32(0) {
		v108 = v86
		goto L10
	} else {
		goto L32
	}
L32:
	;
	if v86 < l1 {
		v46 = v46 + v81
		v48 = v87
		v50 = v86
		goto L14
	} else {
		goto L33
	}
L33:
	;
	v108 = v86
	goto L10
L34:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_copy_read_data[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = int32(0)
	goto L35
L35:
	;
	if v50 < l1 {
		v34 = v46
		v36 = v48
		v38 = v50
		goto L12
	} else {
		goto L36
	}
L36:
	;
	goto L13
}
