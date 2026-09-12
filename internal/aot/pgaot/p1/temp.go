package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TempTablespacePath(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	if l1 != 0 {
		v13 = base.B2i32(base.Ui32(int32(2)) <= base.Ui32(l1-int32(1663)))
	} else {
		v13 = int32(0)
	}
	if v13 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(235304)
		v22 = F_pg_snprintf(m, l0, int32(1024), int32(177065), v6+int32(16))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			m.G0 = v6 + int32(32)
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(235304)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(561568)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(488699)
		v33 = F_pg_snprintf(m, l0, int32(1024), int32(176961), v6)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			m.G0 = v6 + int32(32)
			return
		}
	}
}
func F_isTempNamespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	return base.B2i32(v4 != int32(0)) & base.B2i32(l0 == v4)
}
func F_isTempOrTempToastNamespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v5 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	if v5 != 0 {
		v6 = int32(1)
		if l0 == v5 {
			v13 = v6
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, _consts[318]))
			if v9 == l0 {
				v13 = v6
			} else {
				v13 = int32(0)
			}
		}
	} else {
		v13 = int32(0)
	}
	return v13
}
