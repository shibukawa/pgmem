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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l2&int32(1) != 0 {
		if int32(0) <= l1 {
			v15 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v15)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v19 = m.T0[v18].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if l2&int32(4) == int32(0) {
				} else {
					if v19 == int32(0) {
					} else {
						if base.Ui32(int32(1024)) < base.Ui32(l1) {
							v51 = F__emscripten_memset_bulkmem(m, v19, base.I32_extend8_s(int32(0)), l1)
							mBase = m.M
						} else {
							if l1&int32(3) != 0 {
								v51 = F__emscripten_memset_bulkmem(m, v19, base.I32_extend8_s(int32(0)), l1)
								mBase = m.M
							} else {
								v33 = v19 + l1
								if base.Ui32(v33) <= base.Ui32(v19) {
								} else {
									v39 = v19 + int32(4)
									if base.Ui32(v39) < base.Ui32(v33) {
										v41 = v33
									} else {
										v41 = v39
									}
									v48 = F__emscripten_memset_bulkmem(m, v19, base.I32_extend8_s(int32(0)), (v19^int32(-1)+v41)&int32(-4)+int32(4))
									mBase = m.M
								}
							}
						}
					}
				}
				m.G0 = v7 + int32(16)
				return v19
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(37298), v7)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492304), int32(1254), int32(461279))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
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
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(37298), v7)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(492304), int32(1254), int32(461279))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
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
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v19 = m.T0[v18].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if l2&int32(4) == int32(0) {
				} else {
					if v19 == int32(0) {
					} else {
						if base.Ui32(int32(1024)) < base.Ui32(l1) {
							v51 = F__emscripten_memset_bulkmem(m, v19, base.I32_extend8_s(int32(0)), l1)
							mBase = m.M
						} else {
							if l1&int32(3) != 0 {
								v51 = F__emscripten_memset_bulkmem(m, v19, base.I32_extend8_s(int32(0)), l1)
								mBase = m.M
							} else {
								v33 = v19 + l1
								if base.Ui32(v33) <= base.Ui32(v19) {
								} else {
									v39 = v19 + int32(4)
									if base.Ui32(v39) < base.Ui32(v33) {
										v41 = v33
									} else {
										v41 = v39
									}
									v48 = F__emscripten_memset_bulkmem(m, v19, base.I32_extend8_s(int32(0)), (v19^int32(-1)+v41)&int32(-4)+int32(4))
									mBase = m.M
								}
							}
						}
					}
				}
				m.G0 = v7 + int32(16)
				return v19
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
		v14 = *(*int32)(unsafe.Add(mBase, _consts[11]))
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
					F_errcode(m, int32(8389))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(13904), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v30
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
							F_errdetail(m, int32(664432), v7)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(492304), int32(1164), int32(363389))
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
				F_errcode(m, int32(8389))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(13904), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v30
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
						F_errdetail(m, int32(664432), v7)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492304), int32(1164), int32(363389))
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
