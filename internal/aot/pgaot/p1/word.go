package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_compareWORD(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	if v5 == int32(0) {
		v13 = int32(0)
		if v13 < v7 {
			v16 = int32(-1)
		} else {
			v16 = v13
		}
		v33 = v16
	} else {
		if v7 == int32(0) {
			v33 = base.B2i32(int32(0) < v5)
		} else {
			if base.Ui32(v5) < base.Ui32(v7) {
				v22 = v5
			} else {
				v22 = v7
			}
			v23 = F_memcmp(m, v4, v6, v22)
			mBase = m.M
			if v23 != 0 {
				v31 = v23
				v33 = v31
			} else {
				if v5 == v7 {
					v33 = int32(0)
				} else {
					if v5 < v7 {
						v30 = int32(-1)
					} else {
						v30 = int32(1)
					}
					v31 = v30
					v33 = v31
				}
			}
		}
	}
	if v33 != 0 {
		v44 = v33
	} else {
		v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
		v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
		if v35 == v36 {
			v44 = int32(0)
		} else {
			if base.Ui32(v36) < base.Ui32(v35) {
				v41 = int32(1)
			} else {
				v41 = int32(-1)
			}
			v44 = v41
		}
	}
	return v44
}
func F_compareWordEntryPos(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v4 = int32(_a_F_compareWordEntryPos_0)
	v5 = v3 & v4
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v8 = v6 & v4
	return base.B2i32(base.Ui32(v8) < base.Ui32(v5)) - base.B2i32(base.Ui32(v5) < base.Ui32(v8))
}
func F_word_similarity_dist_op(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14022(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
