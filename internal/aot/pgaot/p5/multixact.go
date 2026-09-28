package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_multixact_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v84 int32
	_ = v84
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)))
	if base.Ui32(v14) <= base.Ui32(int32(31)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(80)
	return
L2:
	;
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v17
	F_appendStringInfo(m, l0, int32(_a_F_multixact_desc_0), v10)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v25 = v14&int32(240) - int32(32)
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	return
L6:
	;
	goto L1
L7:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v76
	F_appendStringInfo(m, l0, int32(_a_F_multixact_desc_1), v10-int32(-64))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L24
	}
L8:
	;
	if v25 == int32(16) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v28
	F_appendStringInfo(m, l0, int32(_a_F_multixact_desc_2), v10+int32(32))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L14
	}
L11:
	;
	goto L7
L12:
	;
	goto L1
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v39 <= int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v46 = int32(0)
	goto L16
L16:
	;
	v54 = v13 + int32(20) + v46<<(uint(int32(3))%32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v55
	F_appendStringInfo(m, l0, int32(_a_F_multixact_desc_3), v10+int32(16))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	goto L1
L18:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if base.Ui32(v62) <= base.Ui32(int32(5)) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62<<(uint(int32(2))%32))+uint32(_c_F_multixact_desc[0])))
	v69 = v67
	goto L21
L20:
	;
	v69 = int32(_a_F_multixact_desc_4)
	goto L21
L21:
	;
	F_appendStringInfoString(m, l0, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v73 = v46 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v73 < v74 {
		v46 = v73
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	goto L1
}
