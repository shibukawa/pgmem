package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MemoryContextAllocExtended(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l2&int32(1) != 0 {
		if int32(0) <= l1 {
			v15 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v15)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v23 = m.T0[v22].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if base.B2i32(l2&int32(4) == v15)|base.B2i32(v23 == int32(0)) != 0 {
				} else {
					if l1&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(l1)) == int32(0) {
						if l1 == int32(0) {
						} else {
							v41 = v23 + l1
							v43 = v23 + int32(4)
							if base.Ui32(v43) < base.Ui32(v41) {
								v45 = v41
							} else {
								v45 = v43
							}
							v50 = (v23^int32(-1)+v45)&int32(-4) + int32(4)
							if v50 == int32(0) {
							} else {
								base.MemoryFill(m, v23, int32(0), v50)
							}
						}
					} else {
						if l1 == int32(0) {
						} else {
							base.MemoryFill(m, v23, int32(0), l1)
						}
					}
				}
				m.G0 = v7 + int32(16)
				return v23
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(_a_F_MemoryContextAllocExtended_0), v7)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_MemoryContextAllocExtended_1), int32(1254), int32(_a_F_MemoryContextAllocExtended_2))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
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
		if base.Ui32(int32(1073741824)) <= base.Ui32(l1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(_a_F_MemoryContextAllocExtended_0), v7)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_MemoryContextAllocExtended_1), int32(1254), int32(_a_F_MemoryContextAllocExtended_2))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v15 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v15)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			v23 = m.T0[v22].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if base.B2i32(l2&int32(4) == v15)|base.B2i32(v23 == int32(0)) != 0 {
				} else {
					if l1&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(l1)) == int32(0) {
						if l1 == int32(0) {
						} else {
							v41 = v23 + l1
							v43 = v23 + int32(4)
							if base.Ui32(v43) < base.Ui32(v41) {
								v45 = v41
							} else {
								v45 = v43
							}
							v50 = (v23^int32(-1)+v45)&int32(-4) + int32(4)
							if v50 == int32(0) {
							} else {
								base.MemoryFill(m, v23, int32(0), v50)
							}
						}
					} else {
						if l1 == int32(0) {
						} else {
							base.MemoryFill(m, v23, int32(0), l1)
						}
					}
				}
				m.G0 = v7 + int32(16)
				return v23
			}
		}
	}
}
func F_MemoryContextAllocHuge(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v3 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v3)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_MemoryContextAllocationFailure(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l2&int32(2) == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_MemoryContextAllocationFailure[0]))
		if v14 != 0 {
			F_MemoryContextStats(m, v14)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(_a_F_MemoryContextAllocationFailure_0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_MemoryContextAllocationFailure_1), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v30
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
							F_errdetail(m, int32(_a_F_MemoryContextAllocationFailure_2), v7)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_MemoryContextAllocationFailure_3), int32(1164), int32(_a_F_MemoryContextAllocationFailure_4))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(_a_F_MemoryContextAllocationFailure_0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_MemoryContextAllocationFailure_1), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v30
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
						F_errdetail(m, int32(_a_F_MemoryContextAllocationFailure_2), v7)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_MemoryContextAllocationFailure_3), int32(1164), int32(_a_F_MemoryContextAllocationFailure_4))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
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
	} else {
		m.G0 = v7 + int32(16)
		return int32(0)
	}
}
func F_MemoryContextSetParent(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v6 != l1 {
		if v6 == int32(0) {
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v11 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v10
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v10
			}
			if v10 == int32(0) {
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v16
			}
		}
		if l1 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l1
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v23
			if v23 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = l0
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = l0
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
			return
		}
	} else {
		return
	}
}
