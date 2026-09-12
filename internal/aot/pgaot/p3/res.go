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
	v12 = F_psprintf(m, int32(671564), v5)
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 != 0 {
		if l0 < int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[776]))
			v13 = l0 ^ int32(-1)
			v16 = v11 + v13<<(uint(int32(2))%32)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v19 = v17 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
			if v19 == int32(0) {
				v23 = int32(4365464)
				v25 = *(*int32)(unsafe.Add(mBase, _consts[777]))
				v26 = int32(1)
				*(*int32)(unsafe.Add(mBase, _consts[777])) = v25 - v26
				v30 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				v35 = v30 + v13<<(uint(int32(6))%32) + int32(24)
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
				*(*int32)(unsafe.Add(mBase, uint32(v35))) = v36 - v26
			} else {
			}
			m.G0 = v6 + int32(16)
			return
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			F_UnpinBufferNoOwner(m, v42+l0<<(uint(int32(6))%32)+int32(-64))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0)
			F_errmsg_internal(m, int32(469193), v6)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				F_errfinish(m, int32(476033), int32(6561), int32(268234))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v2
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 <= v2 {
		v18 = F___memset(m, l0, int32(0), int32(100))
		mBase = m.M
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
			v18 = F___memset(m, l0, int32(0), int32(100))
			mBase = m.M
			F_pfree(m, l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				return
			}
		} else {
			F_ResourceOwnerForget(m, v9, l0, int32(4341756))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v18 = F___memset(m, l0, int32(0), int32(100))
				mBase = m.M
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
