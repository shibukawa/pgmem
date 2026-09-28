package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_ls_archive_statusdir(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	F_pg_ls_dir_files(m, l0, int32(_a_F_pg_ls_archive_statusdir_0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_ls_logicalmapdir(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	F_pg_ls_dir_files(m, l0, int32(_a_F_pg_ls_logicalmapdir_0), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_ls_logicalsnapdir(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	F_pg_ls_dir_files(m, l0, int32(_a_F_pg_ls_logicalsnapdir_0), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_ls_tmpdir(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(1040)
	m.G0 = v7
	v11 = int64(0)
	v14 = F_SearchSysCacheExists(m, int32(69), base.I64_extend_i32_u(l1), v11, v11, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg(m, int32(_a_F_pg_ls_tmpdir_0), v7)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pg_ls_tmpdir_1), int32(658), int32(_a_F_pg_ls_tmpdir_2))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
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
			v35 = v7 + int32(16)
			F_TempTablespacePath(m, v35, l1)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_pg_ls_dir_files(m, l0, v35, int32(1))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					m.G0 = v7 + int32(1040)
					return
				}
			}
		}
	}
}
