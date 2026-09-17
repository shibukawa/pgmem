package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ss_get_location(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ss_get_location[0]))
	v14 = F_LWLockAcquire(m, v10+int32(3072), int32(0))
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ss_get_location[1]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v26 = v23
	goto L4
L3:
	;
	if v26 != v46 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v32 != v18 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v20
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v46 = v44
	goto L3
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v38 != 0 {
		v26 = v38
		goto L4
	} else {
		goto L10
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v34 != v19 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v36 == v20 {
		v46 = v23
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	goto L5
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v49 == v26 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_ss_get_location[0]))
	F_LWLockRelease(m, v68+int32(3072))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L20
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v48
	goto L16
L15:
	;
	goto L16
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v52
	if v52 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v54
	goto L19
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_ss_get_location[1]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v26
	goto L13
L20:
	;
	if base.Ui32(v66) < base.Ui32(l1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v75 = v66
	goto L23
L22:
	;
	v75 = int32(0)
	goto L23
L23:
	;
	return v75
}
