package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatalogTupleInsertWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	F_simple_heap_insert(m, l0, l1)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_CatalogIndexInsert(m, l2, l1, int32(1))
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_CatalogTupleUpdate(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(1)
	v13 = F_palloc0(m, int32(216))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(388)
		F_ExecOpenIndices(m, v13, v15)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			F_simple_heap_update(m, l0, l1, l2, v8+int32(12))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				F_CatalogIndexInsert(m, v13, l2, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_ExecCloseIndices(m, v13)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						F_pfree(m, v13)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_CatalogTupleUpdateWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(1)
	F_simple_heap_update(m, l0, l1, l2, v8+int32(12))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		F_CatalogIndexInsert(m, l3, l2, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	}
}
func F_IsCatalogNamespace(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(11))
}
