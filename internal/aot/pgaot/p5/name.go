package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NameOfDatum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 != 0 {
		v9 = v3
		return v9
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v5 = F_NameListToString(m, v4)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = v5
			return v9
		}
	}
}
