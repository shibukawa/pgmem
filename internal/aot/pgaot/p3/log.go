package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_log_destination(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _consts[1119])) = v4
	return
}
func F_check_log_timezone(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_pg_tzset(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v51 = v4
			m.G0 = v8 + int32(16)
			return v51
		} else {
			v17 = F_pg_tz_acceptable(m, v11)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v22 = *(*int32)(unsafe.Add(mBase, _consts[43]))
					*(*int32)(unsafe.Add(mBase, _consts[508])) = v22
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v25
					v29 = F_format_elog_string(m, int32(183231), v8)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[527])) = v29
						v33 = *(*int32)(unsafe.Add(mBase, _consts[43]))
						*(*int32)(unsafe.Add(mBase, _consts[508])) = v33
						v39 = F_format_elog_string(m, int32(630664), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[509])) = v39
							v51 = v4
							m.G0 = v8 + int32(16)
							return v51
						}
					}
				} else {
					v43 = F_guc_malloc(m, int32(4))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v43
						if v43 == int32(0) {
							v51 = v4
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v43))) = v11
							v51 = int32(1)
						}
						m.G0 = v8 + int32(16)
						return v51
					}
				}
			}
		}
	}
}
