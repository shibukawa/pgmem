package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GenerationGetChunkContext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	v5 = l0 - int32(8)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	if v6&int64(16) != int64(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(40))+8))
		return v13
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v5-base.I32_wrap_i64(int64(base.Ui64(v6)>>(uint(int64(34))%64)))&int32(1073741822))+8))
		return v21
	}
}
func F_build_generation_expression(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = F_build_column_default(m, l0, l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v11+v12<<(uint(int32(4))%32)+l1*int32(100))+16))
			if v23 == int32(0) {
				v38 = v13
				m.G0 = v9 + int32(16)
				return v38
			} else {
				v26 = F_exprCollation(m, v13)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					if v26 == v23 {
						v38 = v13
						m.G0 = v9 + int32(16)
						return v38
					} else {
						v30 = F_palloc0(m, int32(16))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v23
							*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v13
							*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(31)
							v38 = v30
							m.G0 = v9 + int32(16)
							return v38
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v47 + int32(4)
				F_errmsg_internal(m, int32(721136), v9)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(496538), int32(4533), int32(269598))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
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
