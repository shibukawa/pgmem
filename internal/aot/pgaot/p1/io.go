package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IOContextForStrategy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = v8 - int32(1)
		if base.Ui32(int32(3)) <= base.Ui32(v10) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v29
				F_errmsg_internal(m, int32(_a_F_IOContextForStrategy_0), v6)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_IOContextForStrategy_1), int32(736), int32(_a_F_IOContextForStrategy_2))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v10<<(uint(int32(2))%32))+uint32(_c_F_IOContextForStrategy[0])))
			v18 = v15
			m.G0 = v6 + int32(16)
			return v18
		}
	} else {
		v18 = int32(3)
		m.G0 = v6 + int32(16)
		return v18
	}
}
func F_WaitIO(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v39 int64
	_ = v39
	var v48 int64
	_ = v48
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v67 int64
	_ = v67
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v116 int64
	_ = v116
	var v120 int64
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_WaitIO[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = v12 + v13<<(uint(int32(4))%32)
	F_ConditionVariablePrepareToSleep(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = l0 + int32(36)
	goto L3
L3:
	;
	v27 = int64(4194304)
	v29 = base.AtomicRmwOr64(m, l0, int32(24), v27)
	if v29&v27 != int64(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L39
	}
L5:
	;
	v39 = v29
	goto L8
L6:
	;
	v113 = v29
	goto L7
L7:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v114
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v116
	v120 = base.AtomicRmwSub64(m, l0, int32(24), int64(4194304))
	if v113&int64(67108864) != int64(0) {
		goto L29
	} else {
		goto L30
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(_a_F_WaitIO_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(_a_F_WaitIO_1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(_a_F_WaitIO_2)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(0)
	v48 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v48
	if v39&int64(4194304) != v48 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v113 = v103
	goto L7
L10:
	;
	goto L13
L11:
	;
	goto L12
L12:
	;
	v81 = int32(_a_F_WaitIO_3)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_WaitIO[1]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(24))+8))
	if v84 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	F_perform_spin_delay(m, v9+int32(24))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v64 = int64(0)
	v67 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v64, v64)
	if v67&int64(4194304) != v64 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v101 = int64(4194304)
	v103 = base.AtomicRmwOr64(m, l0, int32(24), v101)
	if v103&v101 != int64(0) {
		v39 = v103
		goto L8
	} else {
		goto L28
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitIO[1])) = v99
	goto L18
L20:
	;
	if int32(999) < v82 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v82 < int32(11) {
		goto L18
	} else {
		goto L27
	}
L23:
	;
	v89 = int32(900)
	if v89 <= v82 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v92 = v89
	goto L26
L25:
	;
	v92 = v82
	goto L26
L26:
	;
	v99 = v92 + int32(100)
	goto L19
L27:
	;
	v99 = v82 - int32(1)
	goto L19
L28:
	;
	goto L9
L29:
	;
	v126 = v9 + int32(8)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	goto L32
L30:
	;
	goto L31
L31:
	;
	goto L4
L32:
	;
	if v127 != int32(-1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_pgaio_wref_wait(m, v126)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_ConditionVariableSleep(m, v16, int32(134217736))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	F_ConditionVariablePrepareToSleep(m, v16)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L3
L38:
	;
	goto L3
L39:
	;
	m.G0 = v9 + int32(48)
	return
}
