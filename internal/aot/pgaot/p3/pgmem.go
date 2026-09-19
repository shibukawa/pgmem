package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gen_pgmem_decrypt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
	if v9 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v28 = v10
		v29 = int32(0)
		v30 = m.Env.Pgmem_cipher_run(m, v28, v29, l1, l2, l3, l4, l3)
		mBase = m.M
		if v30 < v29 {
			return int32(-18)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l5))) = v30
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
			v29 = int32(0)
			v30 = m.Env.Pgmem_cipher_run(m, v28, v29, l1, l2, l3, l4, l3)
			mBase = m.M
			if v30 < v29 {
				return int32(-18)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l5))) = v30
				return int32(0)
			}
		}
	}
}
func F_pgmem_shmget(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = m.Env.Pgmem_shmget(m, l0, l1, l2)
	mBase = m.M
	if int32(0) <= v4 {
		return v4
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_pgmem_shmget[0])) = int32(0) - v4
		return int32(-1)
	}
}
func F_pgmem_waitpid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
	v14 = m.Env.Pgmem_waitpid(m, int32(-1), v6+int32(12), int32(1))
	mBase = m.M
	if l0 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v15
	} else {
	}
	if int32(0) <= v14 {
		v24 = v14
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_pgmem_waitpid[0])) = int32(0) - v14
		v24 = int32(-1)
	}
	m.G0 = v6 + int32(16)
	return v24
}
