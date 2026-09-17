package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DeleteSharedComments(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn13826(m, l0, l1, int32(2397), int32(2396))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_LockSharedObjectForSession(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+14)) = uint16(v7)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+12)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(1262)
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v2
	v19 = F_LockAcquire(m, v5, int32(8), int32(1), v2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_SharedFileSetAttach(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(1)
	v9 = l0 + int32(44)
	if v5 != 0 {
		F_s_lock(m, v9, int32(_a_F_SharedFileSetAttach_0), int32(60), int32(_a_F_SharedFileSetAttach_1))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v15 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v15 + int32(1)
				F_on_dsm_detach(m, l1, int32(1096), l0)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					return
				}
			} else {
				v24 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v24
				F_errstart_cold(m, int32(21), v24)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_SharedFileSetAttach_2), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_SharedFileSetAttach_0), int32(73), int32(_a_F_SharedFileSetAttach_1))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
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
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v15 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v15 + int32(1)
			F_on_dsm_detach(m, l1, int32(1096), l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				return
			}
		} else {
			v24 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v24
			F_errstart_cold(m, int32(21), v24)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_SharedFileSetAttach_2), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_SharedFileSetAttach_0), int32(73), int32(_a_F_SharedFileSetAttach_1))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
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
}
func F_UnlockSharedObject(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v9)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v4
	v18 = F_LockRelease(m, v7, l2, v4)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_deleteSharedDependencyRecordsFor(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v7 = F_table_open(m, int32(1214), int32(3))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = int32(0)
		F_shdepDropDependency(m, v7, l0, l1, l2, base.B2i32(l2 == v9), v9, v9, v9)
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			F_relation_close(m, v7, int32(3))
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_shared_dependency_comparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v6) < base.Ui32(v7) {
		return int32(-1)
	} else {
		v11 = int32(1)
		if base.Ui32(v7) < base.Ui32(v6) {
			v34 = v11
			return v34
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if base.Ui32(v13) < base.Ui32(v14) {
				return int32(-1)
			} else {
				if base.Ui32(v14) < base.Ui32(v13) {
					v34 = v11
					return v34
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if base.Ui32(v19) < base.Ui32(v20) {
						return int32(-1)
					} else {
						if base.Ui32(v20) < base.Ui32(v19) {
							v34 = v11
						} else {
							v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+12)))
							v27 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+12)))
							if v26 < v27 {
								v34 = int32(-1)
							} else {
								v34 = base.B2i32(v27 < v26)
							}
						}
						return v34
					}
				}
			}
		}
	}
}
