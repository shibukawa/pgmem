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
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
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
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[0]))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+22)))
	if base.Ui32(int32(32)) <= base.Ui32(v10) {
		F_pgaio_submit_staged(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[0]))
			v19 = v18
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
			if v20 == int32(0) {
				v23 = int32(_a_F_pgaio_io_acquire_nb_0)
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[1])) = v25 + int32(1)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
				if v29 == int32(0) {
					v78 = int32(0)
					v80 = int32(_a_F_pgaio_io_acquire_nb_0)
					v82 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[1]))
					*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[1])) = v82 - int32(1)
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
						v48 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v43
						if l0 != 0 {
							v51 = l0 + int32(608)
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+612))
							if v52 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+612)) = v51
								*(*int32)(unsafe.Add(mBase, uint32(l0)+608)) = v51
							} else {
							}
							v58 = v32 + int32(12)
							*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v51
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
							*(*int32)(unsafe.Add(mBase, uint32(v58))) = v60
							*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v58
							*(*int32)(unsafe.Add(mBase, uint32(v51))) = v58
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
						v80 = int32(_a_F_pgaio_io_acquire_nb_0)
						v82 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[1]))
						*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[1])) = v82 - int32(1)
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
					F_errmsg_internal(m, int32(_a_F_pgaio_io_acquire_nb_1), int32(0))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pgaio_io_acquire_nb_2), int32(199), int32(_a_F_pgaio_io_acquire_nb_3))
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
			v23 = int32(_a_F_pgaio_io_acquire_nb_0)
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[1])) = v25 + int32(1)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
			if v29 == int32(0) {
				v78 = int32(0)
				v80 = int32(_a_F_pgaio_io_acquire_nb_0)
				v82 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[1])) = v82 - int32(1)
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
					v48 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v43
					if l0 != 0 {
						v51 = l0 + int32(608)
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+612))
						if v52 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+612)) = v51
							*(*int32)(unsafe.Add(mBase, uint32(l0)+608)) = v51
						} else {
						}
						v58 = v32 + int32(12)
						*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v51
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
						*(*int32)(unsafe.Add(mBase, uint32(v58))) = v60
						*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v58
						*(*int32)(unsafe.Add(mBase, uint32(v51))) = v58
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
					v80 = int32(_a_F_pgaio_io_acquire_nb_0)
					v82 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[1]))
					*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_acquire_nb[1])) = v82 - int32(1)
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
				F_errmsg_internal(m, int32(_a_F_pgaio_io_acquire_nb_1), int32(0))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pgaio_io_acquire_nb_2), int32(199), int32(_a_F_pgaio_io_acquire_nb_3))
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
		F_errmsg_internal(m, int32(_a_F_pgaio_sync_submit_0), int32(0))
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(_a_F_pgaio_sync_submit_1), int32(44), int32(_a_F_pgaio_sync_submit_2))
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
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v2 = int32(_a_F_pgaio_worker_shmem_init_0)
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_shmem_init[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(64)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_shmem_init[0]))
	v8 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = v8
	v10 = int32(_a_F_pgaio_worker_shmem_init_1)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_shmem_init[1]))
	v12 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v12)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_shmem_init[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v8
	base.MemoryFill(m, v15+int32(28), int32(255), int32(128))
	return
}
func F_pgaio_wref_clear(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-1)
	return
}
