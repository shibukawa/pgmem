package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgl_exit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[0]))
	if v4 != 0 {
		v5 = F_fclose(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[0])) = int32(0)
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[1]))
			if v11 != 0 {
				v12 = F_fflush(m, v11)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[1]))
					v16 = F_fclose(m, v15)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[1])) = int32(0)
						*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[2])) = int32(-1)
						*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[3])) = int32(1)
						m.Env.Exit(m, l0)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[2])) = int32(-1)
				*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[3])) = int32(1)
				m.Env.Exit(m, l0)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[1]))
		if v11 != 0 {
			v12 = F_fflush(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[1]))
				v16 = F_fclose(m, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[1])) = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[2])) = int32(-1)
					*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[3])) = int32(1)
					m.Env.Exit(m, l0)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[2])) = int32(-1)
			*(*int32)(unsafe.Add(mBase, _c_F_pgl_exit[3])) = int32(1)
			m.Env.Exit(m, l0)
			mBase = m.M
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_pgl_popen(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_popen[0]))
	if v5 != 0 {
		v6 = m.T0[v5].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_pgl_popen[1])) = int32(52)
		return int32(0)
	}
}
