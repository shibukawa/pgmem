package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dshash_delete_entry(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1-int32(4))))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = l1 - int32(8)
	v21 = v6 + int32(base.Ui32(v9)>>(uint(int32(32)-v11)%32))<<(uint(int32(2))%32)
	goto L3
L1:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_LWLockRelease(m, v53+v52*int32(20)+int32(8))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L10
	}
L2:
	;
	v52 = int32(base.Ui32(v9) >> (uint(int32(25)) % 32))
	goto L1
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v24 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	F_dsa_free(m, v32, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L9
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_dsa_get_address(m, v27, v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	if v28 != v18 {
		v21 = v28
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L4
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v31
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v39 = int32(base.Ui32(v9) >> (uint(int32(25)) % 32))
	v44 = v37 + v39*int32(20) + int32(24)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v45 - int32(1)
	v52 = v39
	goto L1
L10:
	;
	return
}
