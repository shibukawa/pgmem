package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_wc_isupper(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v35 int32
	_ = v35
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_wc_isupper[0]))
	switch v3 - int32(1) {
	case 0:
		v15 = Fn13968(m, l0, int32(65), int32(_a_F_pg_wc_isupper_0), int32(_a_F_pg_wc_isupper_1), int32(655))
		mBase = m.M
		return v15
	case 1:
		v20 = F_towlower(m, l0)
		mBase = m.M
		return base.B2i32(v20 != l0)
	case 2:
		if base.Ui32(l0) <= base.Ui32(int32(255)) {
			v35 = base.B2i32(base.B2i32(base.Ui32(l0-int32(65)) < base.Ui32(int32(26))) != int32(0))
		} else {
			v35 = int32(0)
		}
		return v35
	default:
		return base.B2i32(base.Ui32(l0-int32(65)) < base.Ui32(int32(26)))
	}
}
