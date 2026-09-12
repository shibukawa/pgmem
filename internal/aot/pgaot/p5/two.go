package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TwoPhaseGetXidByVirtualXID(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	v3 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v16 = F_LWLockAcquire(m, v12+int32(2304), int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[101]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v22 <= int32(0) {
		v70 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	F_LWLockRelease(m, v79+int32(2304))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L14
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = *(*int32)(unsafe.Add(mBase, _consts[102]))
	v32 = int32(0)
	v34 = v3
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(8)+v32<<(uint(int32(2))%32))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+44)))
	if v46 != int32(1) {
		v61 = v34
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v66 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v66)
	v70 = v34
	goto L3
L7:
	;
	goto L6
L8:
	;
	v64 = v32 + int32(1)
	if v64 != v22 {
		v32 = v64
		v34 = v61
		goto L5
	} else {
		goto L13
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v55 = v49 + v50*int32(640) + int32(52)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v28 != v56 {
		v61 = v34
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v27 != v58 {
		v61 = v34
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v34 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v45)+32))
	v61 = v60
	goto L8
L13:
	;
	v70 = v61
	goto L3
L14:
	;
	return v70
}
func F_TwoPhaseTransactionGid(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(568627), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errfinish(m, int32(522203), int32(2689), int32(457740))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		v31 = F_pg_snprintf(m, l2, int32(200), int32(40045), v7)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
