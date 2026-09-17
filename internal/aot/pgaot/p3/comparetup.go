package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_comparetup_heap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v9 == int32(1) {
		if v7&int32(1) != 0 {
			v44 = F_comparetup_heap_tiebreak(m, l0, l1, l2)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v47 = v44
				return v47
			}
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+9)))
			if v16 != 0 {
				v17 = int32(-1)
			} else {
				v17 = int32(1)
			}
			return v17
		}
	} else {
		if v7&int32(1) != 0 {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+9)))
			if v23 != 0 {
				v24 = int32(1)
			} else {
				v24 = int32(-1)
			}
			return v24
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
			v29 = m.T0[v28].(func(*base.Module, int32, int32, int32) int32)(m, v26, v27, v8)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = int32(1)
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+8)))
				if v34 != v33 {
					v41 = v29
					if v41 != 0 {
						v47 = v41
						return v47
					} else {
						v44 = F_comparetup_heap_tiebreak(m, l0, l1, l2)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v47 = v44
							return v47
						}
					}
				} else {
					if v29 < int32(0) {
						v47 = v33
						return v47
					} else {
						v41 = int32(0) - v29
						if v41 != 0 {
							v47 = v41
							return v47
						} else {
							v44 = F_comparetup_heap_tiebreak(m, l0, l1, l2)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v47 = v44
								return v47
							}
						}
					}
				}
			}
		}
	}
}
