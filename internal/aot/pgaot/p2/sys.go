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
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if base.Ui32(l0) <= base.Ui32(int32(84)) {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1181])))
		if v18 != 0 {
			v35 = m.G0
			v37 = v35 - int32(16)
			m.G0 = v37
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
			if v39 == int32(0) {
				F_CatalogCacheInitializeCache(m, v18)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
					switch v44 - int32(1) {
					case 0:
						v68 = v4
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
						v70 = m.T0[v69].(func(*base.Module, int32) int32)(m, l1)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							m.G0 = v37 + int32(16)
							m.G0 = v10 + int32(16)
							return v70 ^ v68
						}
					case 1:
						v61 = v4
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
						v63 = m.T0[v62].(func(*base.Module, int32) int32)(m, l2)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v68 = base.I32_rotl(v63, int32(8)) ^ v61
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
							v70 = m.T0[v69].(func(*base.Module, int32) int32)(m, l1)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								m.G0 = v37 + int32(16)
								m.G0 = v10 + int32(16)
								return v70 ^ v68
							}
						}
					case 2:
						v53 = v4
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
						v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v61 = base.I32_rotl(v56, int32(16)) ^ v53
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
							v63 = m.T0[v62].(func(*base.Module, int32) int32)(m, l2)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								v68 = base.I32_rotl(v63, int32(8)) ^ v61
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
								v70 = m.T0[v69].(func(*base.Module, int32) int32)(m, l1)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									m.G0 = v37 + int32(16)
									m.G0 = v10 + int32(16)
									return v70 ^ v68
								}
							}
						}
					case 3:
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
						v49 = m.T0[v48].(func(*base.Module, int32) int32)(m, int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							v53 = base.I32_rotl(v49, int32(24))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
							v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, int32(0))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v61 = base.I32_rotl(v56, int32(16)) ^ v53
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
								v63 = m.T0[v62].(func(*base.Module, int32) int32)(m, l2)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									v68 = base.I32_rotl(v63, int32(8)) ^ v61
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
									v70 = m.T0[v69].(func(*base.Module, int32) int32)(m, l1)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										m.G0 = v37 + int32(16)
										m.G0 = v10 + int32(16)
										return v70 ^ v68
									}
								}
							}
						}
					default:
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v37))) = v44
							F_errmsg_internal(m, int32(481544), v37)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(499007), int32(373), int32(347263))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
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
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
				switch v44 - int32(1) {
				case 0:
					v68 = v4
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
					v70 = m.T0[v69].(func(*base.Module, int32) int32)(m, l1)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						m.G0 = v37 + int32(16)
						m.G0 = v10 + int32(16)
						return v70 ^ v68
					}
				case 1:
					v61 = v4
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
					v63 = m.T0[v62].(func(*base.Module, int32) int32)(m, l2)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v68 = base.I32_rotl(v63, int32(8)) ^ v61
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
						v70 = m.T0[v69].(func(*base.Module, int32) int32)(m, l1)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							m.G0 = v37 + int32(16)
							m.G0 = v10 + int32(16)
							return v70 ^ v68
						}
					}
				case 2:
					v53 = v4
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
					v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						v61 = base.I32_rotl(v56, int32(16)) ^ v53
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
						v63 = m.T0[v62].(func(*base.Module, int32) int32)(m, l2)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v68 = base.I32_rotl(v63, int32(8)) ^ v61
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
							v70 = m.T0[v69].(func(*base.Module, int32) int32)(m, l1)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								m.G0 = v37 + int32(16)
								m.G0 = v10 + int32(16)
								return v70 ^ v68
							}
						}
					}
				case 3:
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
					v49 = m.T0[v48].(func(*base.Module, int32) int32)(m, int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v53 = base.I32_rotl(v49, int32(24))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
						v56 = m.T0[v55].(func(*base.Module, int32) int32)(m, int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v61 = base.I32_rotl(v56, int32(16)) ^ v53
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
							v63 = m.T0[v62].(func(*base.Module, int32) int32)(m, l2)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								v68 = base.I32_rotl(v63, int32(8)) ^ v61
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
								v70 = m.T0[v69].(func(*base.Module, int32) int32)(m, l1)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									m.G0 = v37 + int32(16)
									m.G0 = v10 + int32(16)
									return v70 ^ v68
								}
							}
						}
					}
				default:
					F_errstart_cold(m, int32(22), int32(0))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v37))) = v44
						F_errmsg_internal(m, int32(481544), v37)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(499007), int32(373), int32(347263))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
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
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(487718), v10)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499019), int32(669), int32(347326))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
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
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
			F_errmsg_internal(m, int32(487718), v10)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(499019), int32(669), int32(347326))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1181])))
	v12 = F_SearchCatCacheInternal(m, v9, int32(3), l1, l2, l3, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
