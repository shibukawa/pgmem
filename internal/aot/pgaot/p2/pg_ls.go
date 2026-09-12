package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_ls_archive_statusdir(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_pg_ls_dir_files(m, l0, int32(122665), int32(1))
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
	F_pg_ls_dir_files(m, l0, int32(166587), int32(0))
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
	F_pg_ls_dir_files(m, l0, int32(127194), int32(0))
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
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(1040)
	m.G0 = v6
	v12 = F_SearchSysCacheExists(m, int32(69), l1, v3, v3, v3)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
					F_errmsg(m, int32(75499), v6)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errfinish(m, int32(525303), int32(657), int32(224878))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
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
			F_TempTablespacePath(m, v6+int32(16), l1)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				F_pg_ls_dir_files(m, l0, v6+int32(16), int32(1))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					m.G0 = v6 + int32(1040)
					return
				}
			}
		}
	}
}
