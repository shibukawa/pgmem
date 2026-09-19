package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pointerhash_iterate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v10 = v7
	goto L1
L1:
	;
	if v10&int32(1) != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v27
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v23 = v19 & (v20 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23
	v27 = v18 + v20<<(uint(int32(3))%32)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v29 = v19 & v28
	if v29 == v23 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v31)
	goto L8
L7:
	;
	goto L8
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)))
	if v34 != int32(1) {
		v10 = base.B2i32(v23 == v29)
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L2
}
func F_pointerhash_lookup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = int32(16)
	v12 = (int32(base.Ui32(l1)>>(uint(v8)%32)) ^ l1) * int32(-2048144789)
	v17 = (int32(base.Ui32(v12)>>(uint(int32(13))%32)) ^ v12) * int32(-1028477387)
	v21 = v7 & (int32(base.Ui32(v17)>>(uint(v8)%32)) ^ v17)
	v24 = v6 + v21<<(uint(int32(3))%32)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
	if v25 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v30 = v24
	v32 = v21
	goto L5
L4:
	;
	return v30
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v35 == l1 {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v39 = (v32 + int32(1)) & v7
	v42 = v6 + v39<<(uint(int32(3))%32)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+4)))
	if v43 != 0 {
		v30 = v42
		v32 = v39
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
func F_pointerhash_stat(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	Fn13976(m, l0, int32(_a_F_pointerhash_stat_0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
