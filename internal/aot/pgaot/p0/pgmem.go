package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gen_pgmem_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if int32(0) < v5 {
		m.Env.Pgmem_cipher_free(m, v5)
		mBase = m.M
	} else {
	}
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+96))
	if v9 != 0 {
		F_ResourceOwnerForget(m, v9, v4, int32(_a_F_gen_pgmem_free_0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			base.MemoryFill(m, v4, int32(0), int32(100))
			F_pfree(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		base.MemoryFill(m, v4, int32(0), int32(100))
		F_pfree(m, v4)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		}
	}
}
