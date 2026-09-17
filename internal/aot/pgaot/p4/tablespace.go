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
				v12 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace_maintenance_io_concurrency[0]))
				v13 = v12
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_get_tablespace_maintenance_io_concurrency[0]))
			v13 = v12
		}
		return v13
	}
}
func F_sendTablespace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	v10 = m.G0
	v12 = v10 - int32(1152)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(_a_F_sendTablespace_0)
	v18 = v12 + int32(128)
	v23 = F_pg_snprintf(m, v18, int32(1024), int32(_a_F_sendTablespace_1), v12+int32(16))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int64(0)
	} else {
		v31 = F___fstatat(m, int32(-100), v18, v12+int32(32), int32(256))
		mBase = m.M
		if v31 != 0 {
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_sendTablespace[0]))
			if v33 == int32(44) {
				v66 = int64(0)
				m.G0 = v12 + int32(1152)
				return v66
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v18
						F_errmsg(m, int32(_a_F_sendTablespace_2), v12)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_sendTablespace_3), int32(1160), int32(_a_F_sendTablespace_4))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
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
			F__tarWriteHeader(m, l0, int32(_a_F_sendTablespace_0), int32(0), v12+int32(32), l3)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int64(0)
			} else {
				v59 = F_strlen(m, l1)
				mBase = m.M
				v62 = F_sendDir(m, l0, v12+int32(128), v59, l3, int32(0), int32(1), l4, l2, l5)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int64(0)
				} else {
					v66 = v62 + int64(512)
					m.G0 = v12 + int32(1152)
					return v66
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
	v7 = F_build_reloptions(m, l0, l1, int32(128), int32(32), int32(_a_F_tablespace_reloptions_0), int32(4))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
