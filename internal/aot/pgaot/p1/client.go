package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcessClientReadInterrupt(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v4 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _consts[826])))
	if v6 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[48]))
		if v8 != 0 {
			F_ProcessInterrupts(m)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, _consts[827]))
				if v12 != 0 {
					F_ProcessCatchupInterrupt(m)
					mBase = m.M
					v14 = m.ExcPending
					if v14 != 0 {
						return
					} else {
						v16 = *(*int32)(unsafe.Add(mBase, _consts[183]))
						if v16 == int32(0) {
							*(*int32)(unsafe.Add(mBase, _consts[137])) = v4
							return
						} else {
							F_ProcessNotifyInterrupt(m, int32(1))
							mBase = m.M
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[137])) = v4
								return
							}
						}
					}
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, _consts[183]))
					if v16 == int32(0) {
						*(*int32)(unsafe.Add(mBase, _consts[137])) = v4
						return
					} else {
						F_ProcessNotifyInterrupt(m, int32(1))
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[137])) = v4
							return
						}
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[827]))
			if v12 != 0 {
				F_ProcessCatchupInterrupt(m)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, _consts[183]))
					if v16 == int32(0) {
						*(*int32)(unsafe.Add(mBase, _consts[137])) = v4
						return
					} else {
						F_ProcessNotifyInterrupt(m, int32(1))
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[137])) = v4
							return
						}
					}
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _consts[183]))
				if v16 == int32(0) {
					*(*int32)(unsafe.Add(mBase, _consts[137])) = v4
					return
				} else {
					F_ProcessNotifyInterrupt(m, int32(1))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[137])) = v4
						return
					}
				}
			}
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, _consts[763]))
		if v25 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[137])) = v4
			return
		} else {
			if l0 != 0 {
				v29 = *(*int32)(unsafe.Add(mBase, _consts[48]))
				if v29 == int32(0) {
					*(*int32)(unsafe.Add(mBase, _consts[137])) = v4
					return
				} else {
					F_ProcessInterrupts(m)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[137])) = v4
						return
					}
				}
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, _consts[516]))
				F_SetLatch(m, v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[137])) = v4
					return
				}
			}
		}
	}
}
func F_assign_client_encoding(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[20]))
	if int32(0) <= v8 {
		m.G0 = v5 + int32(16)
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v12 = F_SetClientEncoding(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if int32(0) <= v12 {
				m.G0 = v5 + int32(16)
				return
			} else {
				v18 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					if v18 == int32(0) {
						m.G0 = v5 + int32(16)
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v11
						F_errmsg_internal(m, int32(453826), v5)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_errfinish(m, int32(498468), int32(799), int32(335256))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								m.G0 = v5 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
