package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitializeMaxBackends(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v9 = *(*int32)(unsafe.Add(mBase, _consts[310]))
	v11 = *(*int32)(unsafe.Add(mBase, _consts[326]))
	v13 = *(*int32)(unsafe.Add(mBase, _consts[325]))
	v15 = *(*int32)(unsafe.Add(mBase, _consts[1138]))
	v20 = v9 + (v11 + (v13 + v15)) + int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[114])) = v20
	if base.Ui32(int32(262144)) <= base.Ui32(v20) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				F_errmsg(m, int32(472452), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(262142)
					v38 = *(*int32)(unsafe.Add(mBase, _consts[325]))
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = v38
					v41 = *(*int32)(unsafe.Add(mBase, _consts[1138]))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v41
					v44 = *(*int32)(unsafe.Add(mBase, _consts[326]))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v44
					v47 = *(*int32)(unsafe.Add(mBase, _consts[310]))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v47
					F_errdetail(m, int32(684885), v5)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_errfinish(m, int32(517864), int32(570), int32(183811))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
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
	} else {
		m.G0 = v5 + int32(32)
		return
	}
}
func F_MaxLivePostmasterChildren(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, _consts[654]))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(112550), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(525863), int32(73), int32(295592))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		return v3
	}
}
