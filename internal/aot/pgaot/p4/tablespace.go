package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_tablespace_maintenance_io_concurrency(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = F_get_tablespace(m, l0)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
		if v6 != 0 {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
			if int32(0) <= v7 {
				v13 = v7
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, _consts[303]))
				v13 = v12
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[303]))
			v13 = v12
		}
		return v13
	}
}
func F_sendTablespace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v69 int64
	_ = v69
	v9 = m.G0
	v11 = v9 - int32(1152)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(575000)
	v22 = F_pg_snprintf(m, v11+int32(128), int32(1024), int32(181138), v11+int32(16))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int64(0)
	} else {
		v32 = F___fstatat(m, int32(-100), v11+int32(128), v11+int32(32), int32(256))
		mBase = m.M
		if v32 != 0 {
			v34 = *(*int32)(unsafe.Add(mBase, _consts[140]))
			if v34 == int32(44) {
				v69 = int64(0)
				m.G0 = v11 + int32(1152)
				return v69
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int64(0)
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(128)
						F_errmsg(m, int32(302505), v11)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(507264), int32(1160), int32(428855))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			F__tarWriteHeader(m, l0, int32(575000), int32(0), v11+int32(32), l3)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int64(0)
			} else {
				v62 = F_strlen(m, l1)
				mBase = m.M
				v65 = F_sendDir(m, l0, v11+int32(128), v62, l3, int32(0), int32(1), l4, l2, l5)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int64(0)
				} else {
					v69 = v65 + int64(512)
					m.G0 = v11 + int32(1152)
					return v69
				}
			}
		}
	}
}
func F_tablespace_reloptions(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(128), int32(32), int32(772176), int32(4))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
