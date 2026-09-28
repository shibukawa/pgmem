package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InvalidateSyncingRelStates(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_InvalidateSyncingRelStates[0])) = int32(0)
	return
}
func F_InvalidateVictimBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v45 int64
	_ = v45
	var v56 int64
	_ = v56
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v76 int64
	_ = v76
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
	var v141 int64
	_ = v141
	var v145 int64
	_ = v145
	var v151 int64
	_ = v151
	var v157 int64
	_ = v157
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v173 int64
	_ = v173
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v16
	v18 = F_BufTableHashCode(m, v10)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateVictimBuffer[0]))
	v30 = v23 + v18&int32(127)<<(uint(int32(7))%32) + int32(_a_F_InvalidateVictimBuffer_0)
	v32 = F_LWLockAcquire(m, v30, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = int64(4194304)
	v36 = base.AtomicRmwOr64(m, l0, int32(24), v34)
	if v36&v34 != int64(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v45 = v36
	goto L7
L5:
	;
	v122 = v36
	goto L6
L6:
	;
	v126 = v122 & int64(8650751)
	if v126 != int64(1) {
		goto L29
	} else {
		goto L30
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(_a_F_InvalidateVictimBuffer_1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(_a_F_InvalidateVictimBuffer_2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(_a_F_InvalidateVictimBuffer_3)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(0)
	v56 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v56
	if v45&int64(4194304) != v56 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v122 = v113
	goto L6
L9:
	;
	goto L12
L10:
	;
	goto L11
L11:
	;
	v91 = int32(_a_F_InvalidateVictimBuffer_4)
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateVictimBuffer[1]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(24))+8))
	if v94 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	F_perform_spin_delay(m, v10+int32(24))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	v73 = int64(0)
	v76 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v73, v73)
	if v76&int64(4194304) != v73 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v111 = int64(4194304)
	v113 = base.AtomicRmwOr64(m, l0, int32(24), v111)
	if v113&v111 != int64(0) {
		v45 = v113
		goto L7
	} else {
		goto L27
	}
L17:
	;
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InvalidateVictimBuffer[1])) = v109
	goto L17
L19:
	;
	if int32(999) < v92 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v92 < int32(11) {
		goto L17
	} else {
		goto L26
	}
L22:
	;
	v99 = int32(900)
	if v99 <= v92 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v102 = v99
	goto L25
L24:
	;
	v102 = v92
	goto L25
L25:
	;
	v109 = v102 + int32(100)
	goto L18
L26:
	;
	v109 = v92 - int32(1)
	goto L18
L27:
	;
	goto L8
L28:
	;
	m.G0 = v10 + int32(48)
	return base.B2i32(v126 == int64(1))
L29:
	;
	v131 = base.AtomicRmwSub64(m, l0, int32(24), int64(4194304))
	F_LWLockRelease(m, v30)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v141 = v122 | int64(4194304)
	v145 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v141, v122&int64(-17179869183))
	if v145 != v141 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L28
L33:
	;
	v151 = v145
	goto L36
L34:
	;
	goto L35
L35:
	;
	F_BufTableDelete(m, v10, v18)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L39
	}
L36:
	;
	v157 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v151, v151&int64(-17179607041))
	if v151 != v157 {
		v151 = v157
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	goto L37
L39:
	;
	F_LWLockRelease(m, v30)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v170 = int64(0)
	v173 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v170, v170)
	goto L28
}
