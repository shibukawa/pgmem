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
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v7 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[738]))
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
				F_errmsg_internal(m, int32(464622), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					F_errfinish(m, int32(521899), int32(160), int32(367708))
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[738]))
	v4 = F_get_hash_value(m, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_LockBufHdr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(240134)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(517496)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v18 | v19
	if v18&v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L4
L2:
	;
	v41 = v18
	goto L3
L3:
	;
	v45 = int32(4160012)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[739]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(8))+8))
	if v48 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	F_perform_spin_delay(m, v6+int32(8))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v41 = v33
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v34 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v33 | v34
	if v33&v34 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	m.G0 = v6 + int32(32)
	return v41 | int32(4194304)
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[739])) = v63
	goto L10
L12:
	;
	if int32(999) < v46 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v46 < int32(11) {
		goto L10
	} else {
		goto L19
	}
L15:
	;
	v53 = int32(900)
	if v53 <= v46 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v56 = v53
	goto L18
L17:
	;
	v56 = v46
	goto L18
L18:
	;
	v63 = v56 + int32(100)
	goto L11
L19:
	;
	v63 = v46 - int32(1)
	goto L11
}
