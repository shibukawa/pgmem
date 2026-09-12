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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1035]))
	if v4 != 0 {
		v5 = F_fclose(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1035])) = int32(0)
			v11 = *(*int32)(unsafe.Add(mBase, _consts[1036]))
			if v11 != 0 {
				v12 = F_fflush(m, v11)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, _consts[1036]))
					v16 = F_fclose(m, v15)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[1036])) = int32(0)
						*(*int32)(unsafe.Add(mBase, _consts[1037])) = int32(-1)
						*(*int32)(unsafe.Add(mBase, _consts[670])) = int32(1)
						m.Env.Exit(m, l0)
						mBase = m.M
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[1037])) = int32(-1)
				*(*int32)(unsafe.Add(mBase, _consts[670])) = int32(1)
				m.Env.Exit(m, l0)
				mBase = m.M
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[1036]))
		if v11 != 0 {
			v12 = F_fflush(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _consts[1036]))
				v16 = F_fclose(m, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[1036])) = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[1037])) = int32(-1)
					*(*int32)(unsafe.Add(mBase, _consts[670])) = int32(1)
					m.Env.Exit(m, l0)
					mBase = m.M
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1037])) = int32(-1)
			*(*int32)(unsafe.Add(mBase, _consts[670])) = int32(1)
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1033]))
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
		*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(52)
		return int32(0)
	}
}
func F_pgl_send(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1040]))
	v7 = m.T0[v6].(func(*base.Module, int32, int32) int32)(m, l1, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_pgl_set_pclose_fn(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[1034])) = l0
	return
}
func F_pgl_set_popen_fn(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[1033])) = l0
	return
}
func F_pgl_shmget(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v8 = *(*int32)(unsafe.Add(mBase, _consts[1038]))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	if l2&int32(512) != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if l0 == v15 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	return v17
L7:
	;
	goto L8
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v19 != 0 {
		v12 = v19
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	v32 = int32(65536)
	goto L13
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(44)
	return int32(-1)
L13:
	;
	if base.Ui32(v32) < base.Ui32(l1) {
		v32 = v32 << (uint(int32(1)) % 32)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v38 = F_emscripten_builtin_malloc(m, v32)
	mBase = m.M
	if v38 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(48)
	return int32(-1)
L17:
	;
	goto L18
L18:
	;
	v47 = F_emscripten_builtin_malloc(m, int32(24))
	mBase = m.M
	if v47 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_emscripten_builtin_free(m, v38)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(48)
	return int32(-1)
L20:
	;
	goto L21
L21:
	;
	v56 = int32(4141164)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[1039]))
	*(*int32)(unsafe.Add(mBase, _consts[1039])) = v58 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v58
	v67 = int32(4554840)
	v68 = *(*int32)(unsafe.Add(mBase, _consts[1038]))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v68
	*(*int32)(unsafe.Add(mBase, _consts[1038])) = v47
	return v58
}
