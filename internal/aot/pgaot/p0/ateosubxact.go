package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOSubXact_LargeObject(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	if v10 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	if v14 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[362]))
	v22 = v14
	v24 = v18
	v25 = v4
	goto L4
L4:
	;
	v29 = v24 + v25<<(uint(int32(2))%32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v30 == int32(0) {
		v49 = v22
		v50 = v24
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v52 = v25 + int32(1)
	if v52 < v49 {
		v22 = v49
		v24 = v50
		v25 = v52
		goto L4
	} else {
		goto L18
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v33 != l1 {
		v49 = v22
		v50 = v24
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if l0 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = l2
	v49 = v22
	v50 = v24
	goto L6
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v38 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[365]))
	F_UnregisterSnapshotFromOwner(m, v38, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	F_pfree(m, v30)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L15
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	goto L14
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	v48 = *(*int32)(unsafe.Add(mBase, _consts[362]))
	v49 = v46
	v50 = v48
	goto L6
L18:
	;
	goto L5
}
