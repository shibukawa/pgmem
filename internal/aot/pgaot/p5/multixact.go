package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_multixact_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v84 int32
	_ = v84
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)))
	if base.Ui32(v15) <= base.Ui32(int32(31)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 - int32(-64)
	return
L2:
	;
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v18
	F_appendStringInfo(m, l0, int32(450322), v11)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	switch v15&int32(240) - int32(32) {
	case 0:
		goto L8
	default:
		goto L1
	case 16:
		goto L7
	}
L5:
	;
	return
L6:
	;
	goto L1
L7:
	;
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v14)+4))
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v76
	F_appendStringInfo(m, l0, int32(701233), v9+int32(-16))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L19
	}
L8:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v27
	F_appendStringInfo(m, l0, int32(778565), v9+int32(-32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v36 <= int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v43 = int32(0)
	goto L11
L11:
	;
	v52 = v14 + int32(12) + v43<<(uint(int32(3))%32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v53
	F_appendStringInfo(m, l0, int32(764180), v9+int32(-48))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L1
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if base.Ui32(v61) <= base.Ui32(int32(5)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61<<(uint(int32(2))%32))+uint32(_consts[63])))
	v69 = v68
	goto L16
L15:
	;
	v69 = int32(778965)
	goto L16
L16:
	;
	F_appendStringInfoString(m, l0, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v73 = v43 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v73 < v74 {
		v43 = v73
		goto L11
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	goto L1
}
