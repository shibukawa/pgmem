package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_on_shmem_exit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v6 = *(*int32)(unsafe.Add(mBase, _consts[838]))
	if v6 < int32(20) {
		v10 = v6 << (uint(int32(3)) % 32)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[839]))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[840]))) = l0
		*(*int32)(unsafe.Add(mBase, _consts[838])) = v6 + int32(1)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, _consts[841])))
		if v22 == int32(0) {
			v30 = *(*int32)(unsafe.Add(mBase, _consts[842]))
			if v30 <= int32(31) {
				*(*int32)(unsafe.Add(mBase, _consts[842])) = v30 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v30<<(uint(int32(2))%32))+uint32(_consts[843]))) = int32(1103)
			} else {
			}
			v45 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[841])) = uint8(v45)
		} else {
		}
		return
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(126527), int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errfinish(m, int32(525084), int32(377), int32(106116))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_on_sl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_line_contain_point(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v12 = F_line_contain_point(m, v4, v5+int32(16))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = v12
				return v14
			}
		} else {
			v14 = int32(0)
			return v14
		}
	}
}
