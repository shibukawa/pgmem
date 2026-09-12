package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_systable_getnext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = F_index_getnext_slot(m, v12, int32(1), v10)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v14 == int32(0) {
				v69 = v2
				F_HandleConcurrentAbort(m)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v69
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v24 = F_ExecFetchSlotHeapTuple(m, v20, int32(0), v8+int32(15))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+72)))
					if v27 != int32(1) {
						v69 = v24
						F_HandleConcurrentAbort(m)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return v69
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(428902), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(478028), int32(536), int32(61157))
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
		*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v45
		v48 = *(*int32)(unsafe.Add(mBase, _consts[114]))
		if v48 != 0 {
			v50 = int32(*(*uint8)(unsafe.Add(mBase, _consts[115])))
			if v50&int32(1) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(322411), int32(0))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(313135), int32(1034), int32(80961))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+188))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
				v59 = m.T0[v58].(func(*base.Module, int32, int32, int32) int32)(m, v43, int32(1), v10)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					if v59 == int32(0) {
						v69 = v2
						F_HandleConcurrentAbort(m)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return v69
						}
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v67 = F_ExecFetchSlotHeapTuple(m, v63, int32(0), v8+int32(14))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v69 = v67
							F_HandleConcurrentAbort(m)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(16)
								return v69
							}
						}
					}
				}
			}
		} else {
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+188))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
			v59 = m.T0[v58].(func(*base.Module, int32, int32, int32) int32)(m, v43, int32(1), v10)
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				if v59 == int32(0) {
					v69 = v2
					F_HandleConcurrentAbort(m)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v69
					}
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v67 = F_ExecFetchSlotHeapTuple(m, v63, int32(0), v8+int32(14))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v69 = v67
						F_HandleConcurrentAbort(m)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return v69
						}
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
							F_errmsg_internal(m, int32(428902), int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(478028), int32(742), int32(433897))
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
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+68))
	F_LockBuffer(m, v8, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		F_UnlockTuple(m, v5, v7+int32(4), int32(7))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[116])) = int32(0)
			F_systable_endscan(m, l0)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				return
			}
		}
	}
}
