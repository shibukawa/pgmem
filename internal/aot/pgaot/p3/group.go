package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecEndGroup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_get_sort_group_operators(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	if l7 != 0 {
		v18 = int32(23)
	} else {
		v18 = int32(7)
	}
	v19 = F_lookup_type_cache(m, l0, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
		if l1&base.B2i32(v21 == int32(0)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				F_errcode(m, int32(52461700))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					v50 = F_format_type_be(m, l0)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
						F_errmsg(m, int32(200226), v14)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							F_errhint(m, int32(601305), int32(0))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								F_errfinish(m, int32(519759), int32(217), int32(139847))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
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
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
			if l3&base.B2i32(v25 == int32(0)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					F_errcode(m, int32(52461700))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v50 = F_format_type_be(m, l0)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v14))) = v50
							F_errmsg(m, int32(200226), v14)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								F_errhint(m, int32(601305), int32(0))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_errfinish(m, int32(519759), int32(217), int32(139847))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
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
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
				if v31 != 0 {
					v32 = int32(0)
				} else {
					v32 = l2
				}
				if v32 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							v72 = F_format_type_be(m, l0)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v72
								F_errmsg(m, int32(200067), v14+int32(16))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									F_errfinish(m, int32(519759), int32(222), int32(139847))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
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
				} else {
					if l4 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = v21
					} else {
					}
					if l5 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l5))) = v31
					} else {
					}
					if l6 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l6))) = v25
					} else {
					}
					if l7 != 0 {
						*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(base.B2i32(v29 != int32(0)))
					} else {
					}
					m.G0 = v14 + int32(32)
					return
				}
			}
		}
	}
}
