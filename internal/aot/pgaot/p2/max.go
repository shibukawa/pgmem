package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_max_stack_depth(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v38 int32
	_ = v38
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_check_max_stack_depth[0]))
	if v12 == int32(0) {
		v17 = v8 + int32(16)
		if v17 != 0 {
			switch int32(0) {
			case 0:
				v25 = m.G2
				v26 = m.G1
				v28 = base.I64_extend_i32_u(v25 - v26)
				*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v28
				*(*int64)(unsafe.Add(mBase, uint32(v17))) = v28
			default:
				v31 = int64(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v31
				*(*int64)(unsafe.Add(mBase, uint32(v17))) = v31
			case 4:
				v21 = int64(4096)
				*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v21
				*(*int64)(unsafe.Add(mBase, uint32(v17))) = v21
			}
		} else {
		}
		v38 = F___syscall_ret(m, int32(0))
		mBase = m.M
		v42 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
		if base.Ui64(int64(2147483646)) < base.Ui64(v42) {
			v46 = int32(2147483647)
		} else {
			v46 = base.I32_wrap_i64(v42)
		}
		if v38 < int32(0) {
			v49 = int32(-1)
		} else {
			v49 = v46
		}
		*(*int32)(unsafe.Add(mBase, _c_F_check_max_stack_depth[0])) = v49
		v51 = v49
	} else {
		v51 = v12
	}
	v53 = int32(1)
	if v51 <= int32(0) {
		v87 = v53
		m.G0 = v8 + int32(32)
		return v87
	} else {
		v57 = v51 - int32(_a_F_check_max_stack_depth_0)
		if v10<<(uint(int32(10))%32) <= v57 {
			v87 = v53
			m.G0 = v8 + int32(32)
			return v87
		} else {
			v63 = *(*int32)(unsafe.Add(mBase, _c_F_check_max_stack_depth[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_check_max_stack_depth[2])) = v63
			v67 = base.I32_div_s(v57, int32(1024))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v67
			v71 = F_format_elog_string(m, int32(_a_F_check_max_stack_depth_1), v8)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_check_max_stack_depth[3])) = v71
				v77 = *(*int32)(unsafe.Add(mBase, _c_F_check_max_stack_depth[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_check_max_stack_depth[2])) = v77
				v83 = F_format_elog_string(m, int32(_a_F_check_max_stack_depth_2), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_check_max_stack_depth[4])) = v83
					v87 = int32(0)
					m.G0 = v8 + int32(32)
					return v87
				}
			}
		}
	}
}
