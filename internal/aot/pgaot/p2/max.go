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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, _consts[1236]))
	if v12 == int32(0) {
		v15 = int32(16)
		v16 = v8 + v15
		v21 = m.G0
		v23 = v21 - v15
		m.G0 = v23
		if v16 != 0 {
			v25 = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(v16))) = v25
			*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v25
		} else {
		}
		v29 = int32(0)
		v30 = F___syscall_ret(m, v29)
		mBase = m.M
		if v30 == v29 {
			v71 = int32(0)
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, _consts[158]))
			if v34 != int32(52) {
				v71 = v30
			} else {
				v38 = v23 + int32(8)
				v39 = int64(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v38))) = v39
				*(*int64)(unsafe.Add(mBase, uint32(v38)+8)) = v39
				v43 = int32(0)
				v44 = F___syscall_ret(m, v43)
				mBase = m.M
				if v44 < v43 {
					v71 = int32(-1)
				} else {
					v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+8)))
					if v49 == int64(4294967295) {
						v52 = int64(-1)
					} else {
						v52 = v49
					}
					*(*int64)(unsafe.Add(mBase, uint32(v16))) = v52
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
					if v55 == int32(-1) {
						v59 = int64(-1)
					} else {
						v59 = base.I64_extend_i32_u(v55)
					}
					*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v59
					if v49 == int64(4294967295) {
						*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(-1)
					} else {
					}
					v65 = int32(0)
					if v55 != int32(-1) {
						v71 = v65
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(-1)
						v71 = v65
					}
				}
			}
		}
		m.G0 = v23 + int32(16)
		v79 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
		if base.Ui64(int64(2147483646)) < base.Ui64(v79) {
			v83 = int32(2147483647)
		} else {
			v83 = base.I32_wrap_i64(v79)
		}
		if v71 < int32(0) {
			v86 = int32(-1)
		} else {
			v86 = v83
		}
		*(*int32)(unsafe.Add(mBase, _consts[1236])) = v86
		v88 = v86
	} else {
		v88 = v12
	}
	v90 = int32(1)
	if v88 <= int32(0) {
		v124 = v90
		m.G0 = v8 + int32(32)
		return v124
	} else {
		v94 = v88 - int32(524288)
		if v10<<(uint(int32(10))%32) <= v94 {
			v124 = v90
			m.G0 = v8 + int32(32)
			return v124
		} else {
			v100 = *(*int32)(unsafe.Add(mBase, _consts[158]))
			*(*int32)(unsafe.Add(mBase, _consts[253])) = v100
			v104 = base.I32_div_s(v94, int32(1024))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v104
			v108 = F_format_elog_string(m, int32(661624), v8)
			mBase = m.M
			v111 = m.ExcPending
			if v111 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[254])) = v108
				v114 = *(*int32)(unsafe.Add(mBase, _consts[158]))
				*(*int32)(unsafe.Add(mBase, _consts[253])) = v114
				v120 = F_format_elog_string(m, int32(582590), int32(0))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[255])) = v120
					v124 = int32(0)
					m.G0 = v8 + int32(32)
					return v124
				}
			}
		}
	}
}
