package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_on_exit_clear_xact_state(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v10 int32
	_ = v10
	v4 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_on_exit_clear_xact_state[0])) = v4
	*(*int64)(unsafe.Add(mBase, _c_F_on_exit_clear_xact_state[1])) = v4
	v10 = int32(0)
	*(*uint16)(unsafe.Add(mBase, _c_F_on_exit_clear_xact_state[2])) = uint16(v10)
	return
}
func F_on_shmem_exit(m *base.Module, l0 int32, l1 int64) {
	var v10 int32
	_ = v10
	Fn14230(m, l0, l1, int32(_a_F_on_shmem_exit_0), int32(377), int32(_a_F_on_shmem_exit_1), int32(_a_F_on_shmem_exit_2), int32(_a_F_on_shmem_exit_3), int32(_a_F_on_shmem_exit_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_on_sl(m *base.Module, l0 int32) int64 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_line_contain_point(m, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 != 0 {
			v11 = F_line_contain_point(m, v3, v4+int32(16))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int64(0)
			} else {
				v15 = base.I64_extend_i32_u(v11)
				return v15
			}
		} else {
			v15 = int64(0)
			return v15
		}
	}
}
