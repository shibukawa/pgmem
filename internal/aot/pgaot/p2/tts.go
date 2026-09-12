package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tts_buffer_heap_get_heap_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4 != 0 {
		v32 = v4
		return v32
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v6&int32(4) != 0 {
			v32 = int32(0)
			return v32
		} else {
			v9 = int32(4554128)
			v10 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
			v14 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v14)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v14
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v21 = F_heap_form_tuple(m, v18, v19, v20)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v21
				v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v28 = v26 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v28)
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v10
				v32 = v21
				return v32
			}
		}
	}
}
func F_tts_heap_copy_heap_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4 != 0 {
		v32 = v4
		v34 = F_heap_copytuple(m, v32)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			return v34
		}
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v6&int32(4) != 0 {
			v32 = int32(0)
			v34 = F_heap_copytuple(m, v32)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				return v34
			}
		} else {
			v9 = int32(4554128)
			v10 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
			v14 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v14
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v14)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v21 = F_heap_form_tuple(m, v18, v19, v20)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v21
				v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v28 = v26 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v28)
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v10
				v32 = v21
				v34 = F_heap_copytuple(m, v32)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					return v34
				}
			}
		}
	}
}
func F_tts_heap_materialize(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v4&int32(4) == int32(0) {
		v9 = int32(4554128)
		v10 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v14
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v14)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v18 == v14 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v24 = F_heap_form_tuple(m, v21, v22, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v28 = v24
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v28
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v32 = v30 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v32)
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v10
				return
			}
		} else {
			v26 = F_heap_copytuple(m, v18)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v28 = v26
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v28
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v32 = v30 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v32)
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v10
				return
			}
		}
	} else {
		return
	}
}
func F_tts_minimal_copy_minimal_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v6 != 0 {
		v42 = v6
		v45 = F_heap_copy_minimal_tuple(m, v42, l1)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			return v45
		}
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v8&int32(4) != 0 {
			v42 = int32(0)
			v45 = F_heap_copy_minimal_tuple(m, v42, l1)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				return v45
			}
		} else {
			v11 = int32(4554128)
			v12 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
			v16 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v16
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v16)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v24 = F_heap_form_minimal_tuple(m, v20, v21, v22, v16)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v24
				v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v31 = v29 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v31)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v34 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v24 - v34
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v33 + v34
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
				v42 = v24
				v45 = F_heap_copy_minimal_tuple(m, v42, l1)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					return v45
				}
			}
		}
	}
}
func F_tts_virtual_getsysattr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	F_errstart_cold(m, int32(21), int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(65895), int32(0))
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(516973), int32(147), int32(216631))
				v21 = m.ExcPending
				if v21 != 0 {
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
