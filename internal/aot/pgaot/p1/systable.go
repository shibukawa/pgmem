package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_systable_getnext(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = F_index_getnext_slot(m, v11, int32(1), v9)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v13 == int32(0) {
				v62 = int32(0)
				F_HandleConcurrentAbort(m)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v62
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v24 = F_ExecFetchSlotHeapTuple(m, v20, int32(0), v7+int32(15))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+72)))
					if v27 != int32(1) {
						v62 = v24
						F_HandleConcurrentAbort(m)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(16)
							return v62
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_systable_getnext_0), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_systable_getnext_1), int32(537), int32(_a_F_systable_getnext_2))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
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
		}
	} else {
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+56))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v45
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+188))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
		v51 = m.T0[v50].(func(*base.Module, int32, int32, int32) int32)(m, v43, int32(1), v9)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			if v51 == int32(0) {
				v62 = int32(0)
				F_HandleConcurrentAbort(m)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v62
				}
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v60 = F_ExecFetchSlotHeapTuple(m, v56, int32(0), v7+int32(14))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v62 = v60
					F_HandleConcurrentAbort(m)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v62
					}
				}
			}
		}
	}
}
func F_systable_getnext_ordered(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_index_getnext_slot(m, v4, l1, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v37 = v3
			F_HandleConcurrentAbort(m)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				return v37
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v13 = int32(0)
			v15 = F_ExecFetchSlotHeapTuple(m, v12, v13, v13)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v15 == int32(0) {
					v37 = v3
					F_HandleConcurrentAbort(m)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						return v37
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+72)))
					if v20 != int32(1) {
						v37 = v15
						F_HandleConcurrentAbort(m)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							return v37
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_systable_getnext_ordered_0), int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_systable_getnext_ordered_1), int32(744), int32(_a_F_systable_getnext_ordered_2))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
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
		}
	}
}
func F_systable_inplace_update_cancel(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
	F_UnlockBuffer(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_UnlockTuple(m, v5, v7+int32(4), int32(7))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_cancel[0])) = int32(0)
			F_systable_endscan(m, l0)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				return
			}
		}
	}
}
