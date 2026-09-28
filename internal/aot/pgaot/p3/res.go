package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResOwnerPrintRelCache(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))+48))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9 + int32(4)
	v14 = F_psprintf(m, int32(_a_F_ResOwnerPrintRelCache_0), v6)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v14
	}
}
func F_ResOwnerReleasePGMEMCipher(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	v2 = int32(0)
	v4 = base.I32_wrap_i64(l0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+96)) = v2
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v7 <= v2 {
		base.MemoryFill(m, v4, int32(0), int32(100))
		F_pfree(m, v4)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			return
		}
	} else {
		m.Env.Pgmem_cipher_free(m, v7)
		mBase = m.M
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+96))
		if v11 == int32(0) {
			base.MemoryFill(m, v4, int32(0), int32(100))
			F_pfree(m, v4)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				return
			}
		} else {
			F_ResourceOwnerForget(m, v11, l0&int64(4294967295), int32(_a_F_ResOwnerReleasePGMEMCipher_0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				base.MemoryFill(m, v4, int32(0), int32(100))
				F_pfree(m, v4)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
