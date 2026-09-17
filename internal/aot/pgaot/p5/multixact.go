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
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v82 int32
	_ = v82
	v8 = m.G0
	v10 = v8 + int32(-64)
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
	m.G0 = v10 - int32(-64)
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
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v13)+4))
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v75
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v74
	F_appendStringInfo(m, l0, int32(_a_F_multixact_desc_1), v8+int32(-16))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v28
	F_appendStringInfo(m, l0, int32(_a_F_multixact_desc_2), v8+int32(-32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
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
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v37 <= int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v44 = int32(0)
	goto L16
L16:
	;
	v52 = v13 + int32(12) + v44<<(uint(int32(3))%32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v53
	F_appendStringInfo(m, l0, int32(_a_F_multixact_desc_3), v8+int32(-48))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	goto L1
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if base.Ui32(v60) <= base.Ui32(int32(5)) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60<<(uint(int32(2))%32))+uint32(_c_F_multixact_desc[0])))
	v67 = v65
	goto L21
L20:
	;
	v67 = int32(_a_F_multixact_desc_4)
	goto L21
L21:
	;
	F_appendStringInfoString(m, l0, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v71 = v44 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v71 < v72 {
		v44 = v71
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
