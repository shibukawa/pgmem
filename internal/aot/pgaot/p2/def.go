package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_defGetInt64(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		switch v10 - int32(465) {
		case 0:
			v40 = int64(*(*int32)(unsafe.Add(mBase, uint32(v9)+4)))
			v41 = v40
			m.G0 = v7 + int32(32)
			return v41
		case 1:
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v16 = F_DirectFunctionCall1Coll(m, int32(546), int32(0), v15)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				v20 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
				v41 = v20
				m.G0 = v7 + int32(32)
				return v41
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16801924))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v28
					F_errmsg(m, int32(344561), v7+int32(16))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(495662), int32(197), int32(548583))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int64(0)
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v53
				F_errmsg(m, int32(344561), v7)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(495662), int32(179), int32(548583))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int64(0)
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
func F_defGetTypeName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		if v10 != int32(68) {
			if v10 != int32(468) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16801924))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v55
						F_errmsg(m, int32(379420), v7+int32(16))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495662), int32(289), int32(379768))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v9
				*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v9
				v20 = F_list_make1_impl(m, int32(1), v7+int32(24))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = F_makeTypeNameFromNameList(m, v20)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = v24
						m.G0 = v7 + int32(32)
						return v26
					}
				}
			}
		} else {
			v26 = v9
			m.G0 = v7 + int32(32)
			return v26
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v38
				F_errmsg(m, int32(215803), v7)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495662), int32(277), int32(379768))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
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
}
func F_makeDefElem(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_palloc0(m, int32(24))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(93)
		return v6
	}
}
