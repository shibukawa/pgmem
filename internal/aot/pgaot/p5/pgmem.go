package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gen_pgmem_iv_size(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	return v4
}
func F_pgmem_aes_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+92))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if base.Ui32(l2) < base.Ui32(int32(17)) {
		v21 = int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v21
		if l2 != 0 {
			base.MemoryCopy(m, v8+int32(4), l1, l2)
		} else {
		}
		v28 = v8 + int32(68)
		if l3 != 0 {
			if v10 != 0 {
				base.MemoryCopy(m, v28, l3, v10)
			} else {
			}
			return int32(0)
		} else {
			v32 = int32(0)
			if v10 == v32 {
				v39 = v32
			} else {
				base.MemoryFill(m, v28, int32(0), v10)
				v39 = v32
			}
			return v39
		}
	} else {
		if base.Ui32(l2) < base.Ui32(int32(25)) {
			v21 = int32(24)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v21
			if l2 != 0 {
				base.MemoryCopy(m, v8+int32(4), l1, l2)
			} else {
			}
			v28 = v8 + int32(68)
			if l3 != 0 {
				if v10 != 0 {
					base.MemoryCopy(m, v28, l3, v10)
				} else {
				}
				return int32(0)
			} else {
				v32 = int32(0)
				if v10 == v32 {
					v39 = v32
				} else {
					base.MemoryFill(m, v28, int32(0), v10)
					v39 = v32
				}
				return v39
			}
		} else {
			v17 = int32(32)
			if base.Ui32(v17) < base.Ui32(l2) {
				v39 = int32(-7)
				return v39
			} else {
				v21 = v17
				*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v21
				if l2 != 0 {
					base.MemoryCopy(m, v8+int32(4), l1, l2)
				} else {
				}
				v28 = v8 + int32(68)
				if l3 != 0 {
					if v10 != 0 {
						base.MemoryCopy(m, v28, l3, v10)
					} else {
					}
					return int32(0)
				} else {
					v32 = int32(0)
					if v10 == v32 {
						v39 = v32
					} else {
						base.MemoryFill(m, v28, int32(0), v10)
						v39 = v32
					}
					return v39
				}
			}
		}
	}
}
func F_pgmem_call_sighandler(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	m.T0[l0].(func(*base.Module, int32))(m, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_pgmem_des3_init(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+92))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v9 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+84)) = v9
	v11 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v6)+12)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(v6)+20)) = v11
	if base.Ui32(v9) <= base.Ui32(l2) {
		v20 = v9
	} else {
		v20 = l2
	}
	if v20 != 0 {
		base.MemoryCopy(m, v6+int32(4), l1, v20)
	} else {
	}
	v25 = v6 + int32(68)
	if l3 != 0 {
		if v8 == int32(0) {
			return int32(0)
		} else {
			base.MemoryCopy(m, v25, l3, v8)
			return int32(0)
		}
	} else {
		if v8 == int32(0) {
		} else {
			base.MemoryFill(m, v25, int32(0), v8)
		}
		return int32(0)
	}
}
func F_pgmem_init(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_init[0])) = int32(_a_F_pgmem_init_0)
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_init[1])) = int32(_a_F_pgmem_init_1)
	*(*int32)(unsafe.Add(mBase, _c_F_pgmem_init[2])) = int32(_a_F_pgmem_init_2)
	return
}
func F_pgmem_recv(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = m.Env.Pgmem_sock_recv(m, l0, l1, l2)
	mBase = m.M
	if int32(0) <= v4 {
		return v4
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_pgmem_recv[0])) = int32(0) - v4
		return int32(-1)
	}
}
func F_pgmem_send(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = m.Env.Pgmem_sock_send(m, l0, l1, l2)
	mBase = m.M
	if int32(0) <= v4 {
		return v4
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_pgmem_send[0])) = int32(0) - v4
		return int32(-1)
	}
}
func F_pgmem_shmdt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = m.Env.Pgmem_shmdt(m, l0)
	mBase = m.M
	if int32(0) <= v2 {
		return v2
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_pgmem_shmdt[0])) = int32(0) - v2
		return int32(-1)
	}
}
