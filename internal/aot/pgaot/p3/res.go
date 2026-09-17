package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResOwnerPrintRelCache(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v7 + int32(4)
	v12 = F_psprintf(m, int32(_a_F_ResOwnerPrintRelCache_0), v5)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(16)
		return v12
	}
}
func F_ResOwnerReleaseBufferPin(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 != 0 {
		if l0 < int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferPin[0]))
			v13 = l0 ^ int32(-1)
			v16 = v11 + v13<<(uint(int32(2))%32)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v19 = v17 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
			if v19 == int32(0) {
				v23 = int32(_a_F_ResOwnerReleaseBufferPin_0)
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferPin[1]))
				v26 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferPin[1])) = v25 - v26
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferPin[2]))
				v33 = v30 + v13<<(uint(int32(6))%32)
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = v34 - v26
			} else {
			}
			m.G0 = v6 + int32(16)
			return
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, _c_F_ResOwnerReleaseBufferPin[3]))
			F_UnpinBufferNoOwner(m, v40+l0<<(uint(int32(6))%32)+int32(-64))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0)
			F_errmsg_internal(m, int32(_a_F_ResOwnerReleaseBufferPin_1), v6)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ResOwnerReleaseBufferPin_2), int32(_a_F_ResOwnerReleaseBufferPin_3), int32(_a_F_ResOwnerReleaseBufferPin_4))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
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
func F_ResOwnerReleasePGMEMCipher(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v2
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 <= v2 {
		base.MemoryFill(m, l0, int32(0), int32(100))
		F_pfree(m, l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			return
		}
	} else {
		m.Env.Pgmem_cipher_free(m, v5)
		mBase = m.M
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		if v9 == int32(0) {
			base.MemoryFill(m, l0, int32(0), int32(100))
			F_pfree(m, l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				return
			}
		} else {
			F_ResourceOwnerForget(m, v9, l0, int32(_a_F_ResOwnerReleasePGMEMCipher_0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				base.MemoryFill(m, l0, int32(0), int32(100))
				F_pfree(m, l0)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
