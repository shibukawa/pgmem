package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_asyncQueueUnregister(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[448])))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v13 = F_LWLockAcquire(m, v9+int32(3456), int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	return
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[449]))
	v18 = v16 + int32(56)
	v19 = int32(4083424)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v21 = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v18+v20<<(uint(v21)%32)))) = int32(-1)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v16+v27<<(uint(v21)%32))+60)) = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	v35 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	if v33 != v35 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v65<<(uint(int32(5))%32)-int32(-64)))) = int32(-1)
	v78 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v78+int32(3456))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L16
	}
L7:
	;
	v37 = v33
	goto L10
L8:
	;
	goto L9
L9:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v16+v33<<(uint(int32(5))%32)-int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v63
	v65 = v33
	goto L6
L10:
	;
	if v37 == int32(-1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v18+v35<<(uint(int32(5))%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v54
	v57 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v65 = v57
	goto L6
L12:
	;
	v65 = v35
	goto L6
L13:
	;
	goto L14
L14:
	;
	v48 = v18 + v37<<(uint(int32(5))%32) + int32(8)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v49 != v35 {
		v37 = v49
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v84 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[448])) = uint8(v84)
	goto L3
}
