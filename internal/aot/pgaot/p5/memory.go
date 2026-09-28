package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MemoryContextAllocAligned(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	if base.Ui32(l2) <= base.Ui32(int32(8)) {
		v9 = F_MemoryContextAllocExtended(m, l0, l1, l3)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v9
		}
	} else {
		v17 = F_MemoryContextAllocExtended(m, l0, l1+l2, l3&int32(-5))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 == int32(0) {
				return int32(0)
			} else {
				v26 = int32(0)
				v28 = (l2 + v17 + int32(7)) & (v26 - l2)
				v30 = v28 - int32(8)
				*(*int64)(unsafe.Add(mBase, uint32(v30))) = base.I64_extend_i32_u(l2)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v30-v17)<<(uint(int64(34))%64) | int64(6)
				if l3&int32(4) == v26 {
					return v28
				} else {
					if l1&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(l1)) == int32(0) {
						if l1 == int32(0) {
							return v28
						} else {
							v57 = v28 + l1
							v59 = v28 + int32(4)
							if base.Ui32(v59) < base.Ui32(v57) {
								v61 = v57
							} else {
								v61 = v59
							}
							v66 = (v28^int32(-1)+v61)&int32(-4) + int32(4)
							if v66 == int32(0) {
								return v28
							} else {
								base.MemoryFill(m, v28, int32(0), v66)
								return v28
							}
						}
					} else {
						if l1 == int32(0) {
						} else {
							base.MemoryFill(m, v28, int32(0), l1)
						}
						return v28
					}
				}
			}
		}
	}
}
func F_MemoryContextAllocZero(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v3 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v3)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = m.T0[v8].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, v3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if l1&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(l1)) == int32(0) {
			if l1 == int32(0) {
				return v9
			} else {
				v24 = v9 + l1
				v26 = v9 + int32(4)
				if base.Ui32(v26) < base.Ui32(v24) {
					v28 = v24
				} else {
					v28 = v26
				}
				v33 = (v9^int32(-1)+v28)&int32(-4) + int32(4)
				if v33 == int32(0) {
					return v9
				} else {
					base.MemoryFill(m, v9, int32(0), v33)
					return v9
				}
			}
		} else {
			if l1 == int32(0) {
			} else {
				base.MemoryFill(m, v9, int32(0), l1)
			}
			return v9
		}
	}
}
func F_MemoryContextRegisterResetCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v3
	v5 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1
	return
}
func F_MemoryContextResetOnly(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v8 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	m.T0[v16].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L9
	} else {
		goto L11
	}
L6:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	m.T0[v12].(func(*base.Module, int32))(m, v11)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	return
L10:
	;
	goto L4
L11:
	;
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v19)
	goto L3
}
