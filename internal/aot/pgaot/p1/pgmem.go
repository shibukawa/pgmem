package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gen_pgmem_encrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
	if v9 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v28 = v10
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
		v33 = m.Env.Pgmem_cipher_run(m, v28, int32(1), l1, l2, l3, l4, v31+l3)
		mBase = m.M
		if v33 < int32(0) {
			return int32(-19)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l5))) = v33
			return int32(0)
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+84))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
		v20 = m.Env.Pgmem_cipher_create(m, v12, v13, v8+int32(4), v16, v8+int32(68), v19)
		mBase = m.M
		if v20 <= int32(0) {
			return int32(-8)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+88)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v20
			v28 = v20
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
			v33 = m.Env.Pgmem_cipher_run(m, v28, int32(1), l1, l2, l3, l4, v31+l3)
			mBase = m.M
			if v33 < int32(0) {
				return int32(-19)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l5))) = v33
				return int32(0)
			}
		}
	}
}
func F_gen_pgmem_key_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+92))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
	return v4
}
func F_pgmem_module_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = int32(0)
	if l0 < v2 {
		v13 = v2
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[1278]))
		if v6 <= l0 {
			v13 = v2
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(12))+uint32(_consts[1279])))
			v13 = v12
		}
	}
	return v13
}
func F_pgmem_reset_session(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_AbortOutOfAnyTransaction(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		if l0 != 0 {
			F_StartTransactionCommand(m)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(245)
				F_DiscardCommand(m, v5+int32(8), int32(1))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					F_CommitTransactionCommand(m)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v21 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _consts[320])) = uint8(v21)
						v24 = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[318])) = v24
						*(*int32)(unsafe.Add(mBase, _consts[316])) = v24
						*(*int32)(unsafe.Add(mBase, _consts[319])) = v24
						*(*uint8)(unsafe.Add(mBase, _consts[321])) = uint8(v24)
						m.G0 = v5 + int32(16)
						return
					}
				}
			}
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	}
}
