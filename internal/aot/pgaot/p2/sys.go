package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetSysCacheHashValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if base.Ui32(l0) <= base.Ui32(int32(84)) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_GetSysCacheHashValue[0])))
		if v16 != 0 {
			v33 = m.G0
			v35 = v33 - int32(16)
			m.G0 = v35
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
			if v37 == int32(0) {
				F_CatalogCacheInitializeCache(m, v16)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
					switch v42 - int32(1) {
					case 0:
						v66 = v4
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
						v68 = m.T0[v67].(func(*base.Module, int32) int32)(m, l1)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							m.G0 = v35 + int32(16)
							m.G0 = v10 + int32(16)
							return v68 ^ v66
						}
					case 1:
						v59 = v4
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
						v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, l2)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v66 = base.I32_rotl(v61, int32(8)) ^ v59
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
							v68 = m.T0[v67].(func(*base.Module, int32) int32)(m, l1)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								m.G0 = v35 + int32(16)
								m.G0 = v10 + int32(16)
								return v68 ^ v66
							}
						}
					case 2:
						v51 = v4
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
						v54 = m.T0[v53].(func(*base.Module, int32) int32)(m, int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v59 = base.I32_rotl(v54, int32(16)) ^ v51
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
							v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, l2)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v66 = base.I32_rotl(v61, int32(8)) ^ v59
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
								v68 = m.T0[v67].(func(*base.Module, int32) int32)(m, l1)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									m.G0 = v35 + int32(16)
									m.G0 = v10 + int32(16)
									return v68 ^ v66
								}
							}
						}
					case 3:
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
						v47 = m.T0[v46].(func(*base.Module, int32) int32)(m, int32(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v51 = base.I32_rotl(v47, int32(24))
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
							v54 = m.T0[v53].(func(*base.Module, int32) int32)(m, int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v59 = base.I32_rotl(v54, int32(16)) ^ v51
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
								v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, l2)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									v66 = base.I32_rotl(v61, int32(8)) ^ v59
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
									v68 = m.T0[v67].(func(*base.Module, int32) int32)(m, l1)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										m.G0 = v35 + int32(16)
										m.G0 = v10 + int32(16)
										return v68 ^ v66
									}
								}
							}
						}
					default:
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v35))) = v42
							F_errmsg_internal(m, int32(_a_F_GetSysCacheHashValue_0), v35)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_GetSysCacheHashValue_1), int32(373), int32(_a_F_GetSysCacheHashValue_2))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
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
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
				switch v42 - int32(1) {
				case 0:
					v66 = v4
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
					v68 = m.T0[v67].(func(*base.Module, int32) int32)(m, l1)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						m.G0 = v35 + int32(16)
						m.G0 = v10 + int32(16)
						return v68 ^ v66
					}
				case 1:
					v59 = v4
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
					v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, l2)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v66 = base.I32_rotl(v61, int32(8)) ^ v59
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
						v68 = m.T0[v67].(func(*base.Module, int32) int32)(m, l1)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							m.G0 = v35 + int32(16)
							m.G0 = v10 + int32(16)
							return v68 ^ v66
						}
					}
				case 2:
					v51 = v4
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
					v54 = m.T0[v53].(func(*base.Module, int32) int32)(m, int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v59 = base.I32_rotl(v54, int32(16)) ^ v51
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
						v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, l2)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v66 = base.I32_rotl(v61, int32(8)) ^ v59
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
							v68 = m.T0[v67].(func(*base.Module, int32) int32)(m, l1)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								m.G0 = v35 + int32(16)
								m.G0 = v10 + int32(16)
								return v68 ^ v66
							}
						}
					}
				case 3:
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
					v47 = m.T0[v46].(func(*base.Module, int32) int32)(m, int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v51 = base.I32_rotl(v47, int32(24))
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
						v54 = m.T0[v53].(func(*base.Module, int32) int32)(m, int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v59 = base.I32_rotl(v54, int32(16)) ^ v51
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
							v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, l2)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v66 = base.I32_rotl(v61, int32(8)) ^ v59
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
								v68 = m.T0[v67].(func(*base.Module, int32) int32)(m, l1)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									m.G0 = v35 + int32(16)
									m.G0 = v10 + int32(16)
									return v68 ^ v66
								}
							}
						}
					}
				default:
					F_errstart_cold(m, int32(22), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v35))) = v42
						F_errmsg_internal(m, int32(_a_F_GetSysCacheHashValue_0), v35)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_GetSysCacheHashValue_1), int32(373), int32(_a_F_GetSysCacheHashValue_2))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(_a_F_GetSysCacheHashValue_3), v10)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetSysCacheHashValue_4), int32(669), int32(_a_F_GetSysCacheHashValue_5))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
			F_errmsg_internal(m, int32(_a_F_GetSysCacheHashValue_3), v10)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_GetSysCacheHashValue_4), int32(669), int32(_a_F_GetSysCacheHashValue_5))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
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
func F_SearchSysCache3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_SearchSysCache3[0])))
	v12 = F_SearchCatCacheInternal(m, v9, int32(3), l1, l2, l3, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
