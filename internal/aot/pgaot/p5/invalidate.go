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
	var v44 int32
	_ = v44
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
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
	v22 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v29 = v22 + v17&int32(127)<<(uint(int32(7))%32) + int32(6912)
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(241628)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(520922)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v44 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v43 | v44
	if v43&v44 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	goto L7
L5:
	;
	v67 = v43
	goto L6
L6:
	;
	v74 = int32(4164732)
	v75 = *(*int32)(unsafe.Add(mBase, _consts[602]))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(24))+8))
	if v77 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	F_perform_spin_delay(m, v9+int32(24))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v67 = v59
	goto L6
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v60 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v59 | v60
	if v59&v60 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v95 = v67 & int32(8650751)
	if v95 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[602])) = v92
	goto L12
L14:
	;
	if int32(999) < v75 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v75 < int32(11) {
		goto L12
	} else {
		goto L21
	}
L17:
	;
	v82 = int32(900)
	if v82 <= v75 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v85 = v82
	goto L20
L19:
	;
	v85 = v75
	goto L20
L20:
	;
	v92 = v85 + int32(100)
	goto L13
L21:
	;
	v92 = v75 - int32(1)
	goto L13
L22:
	;
	F_LWLockRelease(m, v29)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v67 & int32(-4194305)
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
	v110 = m.ExcPending
	if v110 != 0 {
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
	return base.B2i32(v95 == int32(1))
}
