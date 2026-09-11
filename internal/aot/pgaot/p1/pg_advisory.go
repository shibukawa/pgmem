package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_advisory_lock_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(34209794)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v9
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_pg_advisory_lock_int4[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v15
	v20 = F_LockAcquire(m, v6, int32(7), int32(1), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return int32(0)
	}
}
func F_pg_try_advisory_lock_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(34209793)
	*(*uint32)(unsafe.Add(mBase, uint32(v6)+8)) = uint32(v9)
	v14 = int64(base.Ui64(v9) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v6)+4)) = uint32(v14)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_try_advisory_lock_int8[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v17
	v20 = int32(1)
	v22 = F_LockAcquire(m, v6, int32(7), v20, v20)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return base.B2i32(v22 != int32(0))
	}
}
func F_pg_try_advisory_lock_shared_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(34209794)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v9
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_pg_try_advisory_lock_shared_int4[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v15
	v18 = int32(1)
	v20 = F_LockAcquire(m, v6, int32(5), v18, v18)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return base.B2i32(v20 != int32(0))
	}
}
