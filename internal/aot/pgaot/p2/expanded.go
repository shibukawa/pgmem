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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = F_FunctionCall2Coll(m, v6, v7, v8, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v36 = int32(-1)
			return v36
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = F_FunctionCall2Coll(m, v15, v16, v17, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 != 0 {
					v36 = int32(1)
					return v36
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v26 = F_FunctionCall2Coll(m, v22, v23, v24, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						if v26 != 0 {
							v36 = int32(-1)
							return v36
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v32 = F_FunctionCall2Coll(m, v28, v29, v30, v31)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v36 = base.B2i32(v32 != int32(0))
								return v36
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
							v42 = v17 << (uint(int32(2)) % 32)
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
						v24 = F_MemoryContextAlloc(m, v21, v17*int32(5))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v17
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v24
							v30 = v24 + v17<<(uint(int32(2))%32)
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
								v42 = v17 << (uint(int32(2)) % 32)
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
					v24 = F_MemoryContextAlloc(m, v21, v17*int32(5))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v17
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v24
						v30 = v24 + v17<<(uint(int32(2))%32)
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
							v42 = v17 << (uint(int32(2)) % 32)
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
						v42 = v17 << (uint(int32(2)) % 32)
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
					v24 = F_MemoryContextAlloc(m, v21, v17*int32(5))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v17
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v24
						v30 = v24 + v17<<(uint(int32(2))%32)
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
							v42 = v17 << (uint(int32(2)) % 32)
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
				v24 = F_MemoryContextAlloc(m, v21, v17*int32(5))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v17
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v24
					v30 = v24 + v17<<(uint(int32(2))%32)
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
						v42 = v17 << (uint(int32(2)) % 32)
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
