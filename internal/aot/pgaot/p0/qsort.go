package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_qsort_partition_list_value_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v9 = F_FunctionCall2Coll(m, v4, v6, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v9)
	}
}
