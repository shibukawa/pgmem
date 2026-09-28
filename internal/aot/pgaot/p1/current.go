package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetCurrentTransactionNestLevel(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_GetCurrentTransactionNestLevel[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+28))
	return v3
}
func F_current_user(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_current_user[0]))
	v7 = F_GetUserNameFromId(m, v5, v3)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v12 = F_DirectFunctionCall1Coll(m, int32(534), v3, base.I64_extend_i32_u(v7))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			return v12
		}
	}
}
