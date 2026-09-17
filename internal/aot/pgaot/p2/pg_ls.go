package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_ls_archive_statusdir(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_pg_ls_dir_files(m, l0, int32(_a_F_pg_ls_archive_statusdir_0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_ls_logicalmapdir(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_pg_ls_dir_files(m, l0, int32(_a_F_pg_ls_logicalmapdir_0), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_ls_logicalsnapdir(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_pg_ls_dir_files(m, l0, int32(_a_F_pg_ls_logicalsnapdir_0), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_pg_ls_tmpdir(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(1040)
	m.G0 = v7
	v13 = F_SearchSysCacheExists(m, int32(69), l1, v3, v3, v3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg(m, int32(_a_F_pg_ls_tmpdir_0), v7)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pg_ls_tmpdir_1), int32(657), int32(_a_F_pg_ls_tmpdir_2))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v34 = v7 + int32(16)
			F_TempTablespacePath(m, v34, l1)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_pg_ls_dir_files(m, l0, v34, int32(1))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					m.G0 = v7 + int32(1040)
					return
				}
			}
		}
	}
}
