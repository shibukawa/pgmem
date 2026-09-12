package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgaio_io_acquire_nb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v9 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+22)))
	if base.Ui32(int32(32)) <= base.Ui32(v10) {
		F_pgaio_submit_staged(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _consts[767]))
			v19 = v18
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
			if v20 == int32(0) {
				v23 = int32(4465212)
				v25 = *(*int32)(unsafe.Add(mBase, _consts[106]))
				*(*int32)(unsafe.Add(mBase, _consts[106])) = v25 + int32(1)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
				if v29 == int32(0) {
					v78 = int32(0)
					v80 = int32(4465212)
					v82 = *(*int32)(unsafe.Add(mBase, _consts[106]))
					*(*int32)(unsafe.Add(mBase, _consts[106])) = v82 - int32(1)
					return v78
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v34
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
					*(*int32)(unsafe.Add(mBase, uint32(v34))) = v36
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					v39 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v38 - v39
					v43 = v32 - int32(24)
					F_pgaio_io_update_state(m, v43, v39)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, _consts[767]))
						*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v43
						if l0 != 0 {
							v51 = v32 + int32(12)
							v53 = l0 + int32(352)
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
							if v54 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+356)) = v53
								*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v53
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v53
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
							*(*int32)(unsafe.Add(mBase, uint32(v51))) = v60
							*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v51
							*(*int32)(unsafe.Add(mBase, uint32(v53))) = v51
							*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = l0
						} else {
						}
						if l1 == int32(0) {
							v78 = v43
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = l1
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v71 & int32(-449)
							v78 = v43
						}
						v80 = int32(4465212)
						v82 = *(*int32)(unsafe.Add(mBase, _consts[106]))
						*(*int32)(unsafe.Add(mBase, _consts[106])) = v82 - int32(1)
						return v78
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(66661), int32(0))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(487614), int32(199), int32(494228))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
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
	} else {
		v19 = v9
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
		if v20 == int32(0) {
			v23 = int32(4465212)
			v25 = *(*int32)(unsafe.Add(mBase, _consts[106]))
			*(*int32)(unsafe.Add(mBase, _consts[106])) = v25 + int32(1)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
			if v29 == int32(0) {
				v78 = int32(0)
				v80 = int32(4465212)
				v82 = *(*int32)(unsafe.Add(mBase, _consts[106]))
				*(*int32)(unsafe.Add(mBase, _consts[106])) = v82 - int32(1)
				return v78
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v34
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
				*(*int32)(unsafe.Add(mBase, uint32(v34))) = v36
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
				v39 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v38 - v39
				v43 = v32 - int32(24)
				F_pgaio_io_update_state(m, v43, v39)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, _consts[767]))
					*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v43
					if l0 != 0 {
						v51 = v32 + int32(12)
						v53 = l0 + int32(352)
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
						if v54 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+356)) = v53
							*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v53
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v53
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
						*(*int32)(unsafe.Add(mBase, uint32(v51))) = v60
						*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v51
						*(*int32)(unsafe.Add(mBase, uint32(v53))) = v51
						*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = l0
					} else {
					}
					if l1 == int32(0) {
						v78 = v43
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = l1
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v71 & int32(-449)
						v78 = v43
					}
					v80 = int32(4465212)
					v82 = *(*int32)(unsafe.Add(mBase, _consts[106]))
					*(*int32)(unsafe.Add(mBase, _consts[106])) = v82 - int32(1)
					return v78
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(66661), int32(0))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(487614), int32(199), int32(494228))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
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
func F_pgaio_io_set_flag(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	v4 = v3 | l1
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v4)
	return
}
func F_pgaio_sync_submit(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	F_errstart_cold(m, int32(21), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		F_errmsg_internal(m, int32(19061), int32(0))
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(491380), int32(44), int32(99645))
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_pgaio_worker_shmem_init(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v16 = F_ShmemInitStruct(m, int32(342504), int32(272), v9+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[768])) = v16
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
		if v19 == int32(0) {
			*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(64)
		} else {
		}
		v31 = F_ShmemInitStruct(m, int32(296622), int32(264), v9+int32(15))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[769])) = v31
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
			if v34 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v31))) = int64(0)
				v40 = v31 + int32(12)
				v42 = v31 + int32(8)
				v47 = int32(0)
				for {
					v50 = v47 << (uint(int32(3)) % 32)
					v52 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v42+v50))) = v52
					*(*uint8)(unsafe.Add(mBase, uint32(v50+v40))) = uint8(v52)
					v58 = v50 | int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v42+v58))) = v52
					*(*uint8)(unsafe.Add(mBase, uint32(v40+v58))) = uint8(v52)
					v66 = v50 | int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(v42+v66))) = v52
					*(*uint8)(unsafe.Add(mBase, uint32(v40+v66))) = uint8(v52)
					v74 = v50 | int32(24)
					*(*int32)(unsafe.Add(mBase, uint32(v42+v74))) = v52
					*(*uint8)(unsafe.Add(mBase, uint32(v74+v40))) = uint8(v52)
					v82 = v47 + int32(4)
					if v82 != int32(32) {
						v47 = v82
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_pgaio_wref_clear(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-1)
	return
}
