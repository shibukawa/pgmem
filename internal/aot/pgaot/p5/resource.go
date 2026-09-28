package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResourceOwnerReleaseAll(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return
L2:
	;
	v28 = v22
	goto L7
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
	v22 = v15
	v23 = v16
	goto L2
L4:
	;
	goto L5
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	if v17 == int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v22 = v17
	v23 = l0 + int32(24)
	goto L2
L7:
	;
	v36 = v23 + v28<<(uint(int32(4))%32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36-int32(8))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if base.Ui32(l1) < base.Ui32(v40) {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L1
L9:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v36-int32(16))))
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v75 = v28 - int32(1)
	if v15 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	v60 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L17
	} else {
		goto L20
	}
L14:
	;
	v46 = m.T0[v45].(func(*base.Module, int64) int32)(m, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+20)) = uint32(v44)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v48
	v54 = F_psprintf(m, int32(_a_F_ResourceOwnerReleaseAll_0), v13+int32(16))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L17
	} else {
		goto L19
	}
L17:
	;
	return
L18:
	;
	v57 = v46
	goto L13
L19:
	;
	v57 = v54
	goto L13
L20:
	;
	if v60 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v57
	F_errmsg_internal(m, int32(_a_F_ResourceOwnerReleaseAll_1), v13)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_pfree(m, v57)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L17
	} else {
		goto L26
	}
L24:
	;
	F_errfinish(m, int32(_a_F_ResourceOwnerReleaseAll_2), int32(395), int32(_a_F_ResourceOwnerReleaseAll_3))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L12
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	m.T0[v80].(func(*base.Module, int64))(m, v44)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L17
	} else {
		goto L31
	}
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v75)
	goto L27
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v75
	goto L27
L31:
	;
	if v75 != 0 {
		v28 = v75
		goto L7
	} else {
		goto L32
	}
L32:
	;
	goto L8
}
