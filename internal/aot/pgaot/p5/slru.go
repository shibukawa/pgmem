package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlruScanDirCbDeleteAll(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	var v6 int64
	_ = v6
	var v10 int32
	_ = v10
	v6 = base.I64_div_s(l2, int64(32))
	F_SlruInternalDeleteSegment(m, l0, v6)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_SlruScanDirCbFindEarliest(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	if v6 != int64(-1) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v10 = m.T0[v9].(func(*base.Module, int64, int64) int32)(m, l2, v6)
		mBase = m.M
		if v10 == int32(0) {
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l3))) = l2
		}
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(l3))) = l2
	}
	return int32(0)
}
func F_check_slru_buffers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = v8 & int32(15)
	if v10 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[87]))
		*(*int32)(unsafe.Add(mBase, _consts[88])) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		v20 = F_format_elog_string(m, int32(681946), v6)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[89])) = v20
			m.G0 = v6 + int32(16)
			return base.B2i32(v10 == int32(0))
		}
	} else {
		m.G0 = v6 + int32(16)
		return base.B2i32(v10 == int32(0))
	}
}
