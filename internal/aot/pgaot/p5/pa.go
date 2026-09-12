package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pa_set_fileset_state(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v3 != 0 {
		F_s_lock(m, l0, int32(495528), int32(1508), int32(351633))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(2)
			v14 = *(*int32)(unsafe.Add(mBase, _consts[509]))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v16
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v18
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v20
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+60)) = v22
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v24
			v26 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = v26
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(2)
		v14 = *(*int32)(unsafe.Add(mBase, _consts[509]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v18
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v20
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+60)) = v22
		v24 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v24
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = v26
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
		return
	}
}
func F_pa_switch_to_partial_serialize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if v10 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v13
			F_errmsg(m, int32(389121), v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(495528), int32(1223), int32(342263))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v23)
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					F_stream_start_internal(m, v26, v23)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						if l1 == int32(0) {
							v33 = *(*int32)(unsafe.Add(mBase, _consts[509]))
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
							F_LockApplyTransactionForSession(m, v34, v36, int32(0), int32(8))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
								*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(1)
								if v42 != 0 {
									F_s_lock(m, v41, int32(495528), int32(1508), int32(351633))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = int32(1)
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = int32(1)
									m.G0 = v6 + int32(16)
									return
								}
							}
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
							*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(1)
							if v42 != 0 {
								F_s_lock(m, v41, int32(495528), int32(1508), int32(351633))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = int32(1)
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = int32(1)
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				}
			}
		} else {
			v23 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v23)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			F_stream_start_internal(m, v26, v23)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				if l1 == int32(0) {
					v33 = *(*int32)(unsafe.Add(mBase, _consts[509]))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
					F_LockApplyTransactionForSession(m, v34, v36, int32(0), int32(8))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
						*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(1)
						if v42 != 0 {
							F_s_lock(m, v41, int32(495528), int32(1508), int32(351633))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = int32(1)
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = int32(1)
							m.G0 = v6 + int32(16)
							return
						}
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(1)
					if v42 != 0 {
						F_s_lock(m, v41, int32(495528), int32(1508), int32(351633))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = int32(1)
							m.G0 = v6 + int32(16)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v41)+32)) = int32(1)
						m.G0 = v6 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_pa_unlock_transaction(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _consts[509]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	F_UnlockApplyTransactionForSession(m, v4, l0, int32(1), int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
