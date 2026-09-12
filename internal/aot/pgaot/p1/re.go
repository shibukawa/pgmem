package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecReScanBitmapAnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v2 < v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = v2
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+v9<<(uint(int32(2))%32))))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_UpdateChangedParamSet(m, v16, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	if v20 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	F_ExecReScan(m, v16)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v26 = v9 + int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v26 < v27 {
		v9 = v26
		goto L4
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	goto L5
}
func F_RE_wchar_execute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	v12 = F_pg_regexec(m, int32(4434420), l0, l1, l2, l3, l4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if base.Ui32(int32(2)) <= base.Ui32(v12) {
			v20 = F_pg_regerror(m, v12, v9+int32(16))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(302252162))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(16)
						F_errmsg(m, int32(195286), v9)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(476522), int32(304), int32(334624))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			m.G0 = v9 + int32(128)
			return base.B2i32(v12 == int32(0))
		}
	}
}
