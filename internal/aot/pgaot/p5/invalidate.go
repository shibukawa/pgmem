package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InvalidateVictimBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v15
	v17 = F_BufTableHashCode(m, v9)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateVictimBuffer[0]))
	v29 = v22 + v17&int32(127)<<(uint(int32(7))%32) + int32(_a_F_InvalidateVictimBuffer_0)
	v31 = F_LWLockAcquire(m, v29, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(_a_F_InvalidateVictimBuffer_1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(_a_F_InvalidateVictimBuffer_2)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(_a_F_InvalidateVictimBuffer_3)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(0)
	v43 = int32(_a_F_InvalidateVictimBuffer_4)
	v45 = base.AtomicRmwOr32(m, l0, int32(24), v43)
	if v45&v43 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	goto L7
L5:
	;
	v65 = v45
	goto L6
L6:
	;
	v72 = int32(_a_F_InvalidateVictimBuffer_5)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateVictimBuffer[1]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(24))+8))
	if v75 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	F_perform_spin_delay(m, v9+int32(24))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v65 = v60
	goto L6
L9:
	;
	v58 = int32(_a_F_InvalidateVictimBuffer_4)
	v60 = base.AtomicRmwOr32(m, l0, int32(24), v58)
	if v60&v58 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v93 = v65 & int32(_a_F_InvalidateVictimBuffer_6)
	if v93 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InvalidateVictimBuffer[1])) = v90
	goto L12
L14:
	;
	if int32(999) < v73 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v73 < int32(11) {
		goto L12
	} else {
		goto L21
	}
L17:
	;
	v80 = int32(900)
	if v80 <= v73 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v83 = v80
	goto L20
L19:
	;
	v83 = v73
	goto L20
L20:
	;
	v90 = v83 + int32(100)
	goto L13
L21:
	;
	v90 = v73 - int32(1)
	goto L13
L22:
	;
	F_LWLockRelease(m, v29)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v65 & int32(-4194305)
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(1)
	F_BufTableDelete(m, v9, v17)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	m.G0 = v9 + int32(48)
	return base.B2i32(v93 == int32(1))
}
