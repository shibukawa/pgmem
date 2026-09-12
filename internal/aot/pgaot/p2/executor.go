package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecutorFinish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	v9 = *(*int32)(unsafe.Add(mBase, _consts[439]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.T0[v9].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v12 = int32(4553888)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	return
L6:
	;
	F_InstrStartNode(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+148))
	if v23 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+128)))
	if v74&int32(32) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v26 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v32 = int32(0)
	goto L13
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v32<<(uint(int32(2))%32))))
	goto L15
L14:
	;
	goto L10
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v15)+152))
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v64 = v32 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v64 < v65 {
		v32 = v64
		goto L13
	} else {
		goto L30
	}
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	F_MemoryContextReset(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	if v52 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	F_ExecReScan(m, v40)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, v40)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	if v56 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+4)))
	if v58&int32(2) == int32(0) {
		goto L15
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L16
L29:
	;
	goto L28
L30:
	;
	goto L14
L31:
	;
	F_AfterTriggerEndQuery(m, v15)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v81 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	F_InstrStopNode(m, v81, float64(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v13
	v87 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+136)) = uint8(v87)
	return
L38:
	;
	goto L37
}
func F_executor_errposition(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	if l0 == int32(0) {
		return
	} else {
		if l1 < int32(0) {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			if v7 == int32(0) {
				return
			} else {
				v10 = F_pg_mbstrlen_with_len(m, v7, l1)
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					v14 = F_errposition(m, v10+int32(1))
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
