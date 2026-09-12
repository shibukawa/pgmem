package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tts_buffer_heap_copy_minimal_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v5 != 0 {
		v33 = v5
		v35 = F_minimal_tuple_from_heap_tuple(m, v33, l1)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			return v35
		}
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v7&int32(4) != 0 {
			v33 = int32(0)
			v35 = F_minimal_tuple_from_heap_tuple(m, v33, l1)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				return v35
			}
		} else {
			v10 = int32(4549024)
			v11 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[28])) = v13
			v15 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v15)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v15
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v22 = F_heap_form_tuple(m, v19, v20, v21)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v22
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v29 = v27 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v29)
				*(*int32)(unsafe.Add(mBase, _consts[28])) = v11
				v33 = v22
				v35 = F_minimal_tuple_from_heap_tuple(m, v33, l1)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v35
				}
			}
		}
	}
}
func F_tts_heap_clear(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v3&int32(4) != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		F_pfree(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			v12 = v9 & int32(-5)
			v13 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v13)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v13)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
			v22 = v12 | int32(2)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v22)
			return
		}
	} else {
		v12 = v3
		v13 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v13)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v13)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(0)
		v22 = v12 | int32(2)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v22)
		return
	}
}
func F_tts_heap_copyslot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v5 = int32(4549024)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[28])) = v6
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v16 == int32(1646216) {
			v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			if v19&int32(4) != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				F_pfree(m, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
					v26 = v25
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
					v29 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v12
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v29)
					v35 = l0 + int32(32)
					*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v29)
					v39 = v26 & int32(-7)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v39)
					v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+8)))
					*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v41)
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v43
					v46 = v39 | int32(4)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v46)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v48
					return
				}
			} else {
				v26 = v19
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
				v29 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v12
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v29)
				v35 = l0 + int32(32)
				*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v29)
				v39 = v26 & int32(-7)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v39)
				v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v41)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v43
				v46 = v39 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v46)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v48
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(91551), int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errfinish(m, int32(514015), int32(1553), int32(400762))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
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
}
func F_tts_minimal_is_current_xact_tuple(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(400036), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(514015), int32(581), int32(399051))
				v19 = m.ExcPending
				if v19 != 0 {
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
func F_tts_virtual_getsomeattrs(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	F_errstart_cold(m, int32(21), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		F_errmsg_internal(m, int32(91728), int32(0))
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_errfinish(m, int32(514015), int32(132), int32(138398))
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_tts_virtual_is_current_xact_tuple(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(400036), int32(0))
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(514015), int32(163), int32(399017))
				v19 = m.ExcPending
				if v19 != 0 {
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
