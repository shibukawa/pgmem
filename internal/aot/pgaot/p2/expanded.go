package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_compare_expanded_ranges(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v10 = F_FunctionCall2Coll(m, v6, v7, v8, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != int64(0) {
			v42 = int32(-1)
			return v42
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v19 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			v21 = F_FunctionCall2Coll(m, v17, v18, v19, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v21 != int64(0) {
					v42 = int32(1)
					return v42
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					v29 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
					v30 = F_FunctionCall2Coll(m, v26, v27, v28, v29)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 != int64(0) {
							v42 = int32(-1)
							return v42
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							v36 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
							v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
							v38 = F_FunctionCall2Coll(m, v34, v35, v36, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v42 = base.B2i32(v38 != int64(0))
								return v42
							}
						}
					}
				}
			}
		}
	}
}
func F_deconstruct_expanded_record(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v6&int32(4) == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v11 == int32(0) {
			v14 = F_expanded_record_fetch_tupdesc(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = v14
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if v18 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					if v19 == v17 {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
						v33 = v18
						v34 = v32
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v35&int32(1) != 0 {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							F_heap_deform_tuple(m, v38, v16, v33, v34)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v50 | int32(4)
								return
							}
						} else {
							v42 = v17 << (uint(int32(3)) % 32)
							if v42 != 0 {
								base.MemoryFill(m, v33, int32(0), v42)
							} else {
							}
							if v17 == int32(0) {
							} else {
								base.MemoryFill(m, v34, int32(1), v17)
							}
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v50 | int32(4)
							return
						}
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v24 = F_MemoryContextAlloc(m, v21, v17*int32(9))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v17
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v24
							v30 = v24 + v17<<(uint(int32(3))%32)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v30
							v33 = v24
							v34 = v30
							v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
							if v35&int32(1) != 0 {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
								F_heap_deform_tuple(m, v38, v16, v33, v34)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v50 | int32(4)
									return
								}
							} else {
								v42 = v17 << (uint(int32(3)) % 32)
								if v42 != 0 {
									base.MemoryFill(m, v33, int32(0), v42)
								} else {
								}
								if v17 == int32(0) {
								} else {
									base.MemoryFill(m, v34, int32(1), v17)
								}
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v50 | int32(4)
								return
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v24 = F_MemoryContextAlloc(m, v21, v17*int32(9))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v17
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v24
						v30 = v24 + v17<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v30
						v33 = v24
						v34 = v30
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v35&int32(1) != 0 {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							F_heap_deform_tuple(m, v38, v16, v33, v34)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v50 | int32(4)
								return
							}
						} else {
							v42 = v17 << (uint(int32(3)) % 32)
							if v42 != 0 {
								base.MemoryFill(m, v33, int32(0), v42)
							} else {
							}
							if v17 == int32(0) {
							} else {
								base.MemoryFill(m, v34, int32(1), v17)
							}
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v50 | int32(4)
							return
						}
					}
				}
			}
		} else {
			v16 = v11
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			if v18 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				if v19 == v17 {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					v33 = v18
					v34 = v32
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v35&int32(1) != 0 {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						F_heap_deform_tuple(m, v38, v16, v33, v34)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v50 | int32(4)
							return
						}
					} else {
						v42 = v17 << (uint(int32(3)) % 32)
						if v42 != 0 {
							base.MemoryFill(m, v33, int32(0), v42)
						} else {
						}
						if v17 == int32(0) {
						} else {
							base.MemoryFill(m, v34, int32(1), v17)
						}
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v50 | int32(4)
						return
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v24 = F_MemoryContextAlloc(m, v21, v17*int32(9))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v17
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v24
						v30 = v24 + v17<<(uint(int32(3))%32)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v30
						v33 = v24
						v34 = v30
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v35&int32(1) != 0 {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							F_heap_deform_tuple(m, v38, v16, v33, v34)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v50 | int32(4)
								return
							}
						} else {
							v42 = v17 << (uint(int32(3)) % 32)
							if v42 != 0 {
								base.MemoryFill(m, v33, int32(0), v42)
							} else {
							}
							if v17 == int32(0) {
							} else {
								base.MemoryFill(m, v34, int32(1), v17)
							}
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v50 | int32(4)
							return
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v24 = F_MemoryContextAlloc(m, v21, v17*int32(9))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v17
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v24
					v30 = v24 + v17<<(uint(int32(3))%32)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v30
					v33 = v24
					v34 = v30
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v35&int32(1) != 0 {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
						F_heap_deform_tuple(m, v38, v16, v33, v34)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v50 | int32(4)
							return
						}
					} else {
						v42 = v17 << (uint(int32(3)) % 32)
						if v42 != 0 {
							base.MemoryFill(m, v33, int32(0), v42)
						} else {
						}
						if v17 == int32(0) {
						} else {
							base.MemoryFill(m, v34, int32(1), v17)
						}
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v50 | int32(4)
						return
					}
				}
			}
		}
	} else {
		return
	}
}
