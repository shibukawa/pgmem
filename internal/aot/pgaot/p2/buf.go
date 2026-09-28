package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufFileTell(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v4
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v6 + v7
	return
}
func F_BufTableDelete(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_BufTableDelete[0]))
	v7 = F_hash_search_with_hash_value(m, v4, l0, l1, int32(2), int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		if v7 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_BufTableDelete_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_BufTableDelete_1), int32(166), int32(_a_F_BufTableDelete_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			return
		}
	}
}
func F_BufTableHashCode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_BufTableHashCode[0]))
	v4 = F_get_hash_value(m, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_LockBufHdr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v17 int64
	_ = v17
	var v26 int64
	_ = v26
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v84 int64
	_ = v84
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = int64(4194304)
	v10 = base.AtomicRmwOr64(m, l0, int32(24), v8)
	if v10&v8 != int64(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = v10
	goto L4
L2:
	;
	v84 = v10
	goto L3
L3:
	;
	m.G0 = v6 + int32(32)
	return v84 | int64(4194304)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(_a_F_LockBufHdr_0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(_a_F_LockBufHdr_1)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(_a_F_LockBufHdr_2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(0)
	v26 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v26
	if v17&int64(4194304) != v26 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v84 = v77
	goto L3
L6:
	;
	goto L9
L7:
	;
	goto L8
L8:
	;
	v55 = int32(_a_F_LockBufHdr_3)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_LockBufHdr[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(8))+8))
	if v58 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	F_perform_spin_delay(m, v6+int32(8))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	return int64(0)
L12:
	;
	v41 = int64(0)
	v44 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v41, v41)
	if v44&int64(4194304) != v41 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v75 = int64(4194304)
	v77 = base.AtomicRmwOr64(m, l0, int32(24), v75)
	if v77&v75 != int64(0) {
		v17 = v77
		goto L4
	} else {
		goto L25
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LockBufHdr[0])) = v73
	goto L15
L17:
	;
	if int32(999) < v56 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v56 < int32(11) {
		goto L15
	} else {
		goto L24
	}
L20:
	;
	v63 = int32(900)
	if v63 <= v56 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v66 = v63
	goto L23
L22:
	;
	v66 = v56
	goto L23
L23:
	;
	v73 = v66 + int32(100)
	goto L16
L24:
	;
	v73 = v56 - int32(1)
	goto L16
L25:
	;
	goto L5
}
