package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_add_rte_to_flat_rtable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int64
	_ = v18
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	v7 = F_palloc(m, int32(136))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v10 = F__emscripten_memcpy_bulkmem(m, v7, l2, int32(136))
		mBase = m.M
		v12 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = v12
		v18 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+96)) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v10)+76)) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v10)+52)) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v10)+60)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v12
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v31 = F_lappend(m, v30, v10)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v31
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			switch v34 {
			case 0:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
				v39 = v35
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v41 = F_lappend_oid(m, v40, v39)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v41
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v45 != 0 {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
						v48 = v46
					} else {
						v48 = int32(0)
					}
					v49 = F_bms_add_member(m, v44, v48)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v49
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
						if v53 != 0 {
							v54 = F_getRTEPermissionInfo(m, l1, v10)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(0)
								v60 = F_addRTEPermissionInfo(m, l0+int32(36), v10)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									v62 = *(*int64)(unsafe.Add(mBase, uint32(v54)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v62
									v64 = *(*int64)(unsafe.Add(mBase, uint32(v54)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v64
									v66 = *(*int64)(unsafe.Add(mBase, uint32(v54)+16))
									*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = v66
									v68 = *(*int64)(unsafe.Add(mBase, uint32(v54)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v68
									v70 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
									*(*int64)(unsafe.Add(mBase, uint32(v60))) = v70
									return
								}
							}
						} else {
							return
						}
					}
				}
			case 1:
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
				if v36 == int32(0) {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
					if v53 != 0 {
						v54 = F_getRTEPermissionInfo(m, l1, v10)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(0)
							v60 = F_addRTEPermissionInfo(m, l0+int32(36), v10)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								v62 = *(*int64)(unsafe.Add(mBase, uint32(v54)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v62
								v64 = *(*int64)(unsafe.Add(mBase, uint32(v54)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v64
								v66 = *(*int64)(unsafe.Add(mBase, uint32(v54)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = v66
								v68 = *(*int64)(unsafe.Add(mBase, uint32(v54)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v68
								v70 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
								*(*int64)(unsafe.Add(mBase, uint32(v60))) = v70
								return
							}
						}
					} else {
						return
					}
				} else {
					v39 = v36
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v41 = F_lappend_oid(m, v40, v39)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v41
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v45 != 0 {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							v48 = v46
						} else {
							v48 = int32(0)
						}
						v49 = F_bms_add_member(m, v44, v48)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v49
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
							if v53 != 0 {
								v54 = F_getRTEPermissionInfo(m, l1, v10)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(0)
									v60 = F_addRTEPermissionInfo(m, l0+int32(36), v10)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										v62 = *(*int64)(unsafe.Add(mBase, uint32(v54)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v62
										v64 = *(*int64)(unsafe.Add(mBase, uint32(v54)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v64
										v66 = *(*int64)(unsafe.Add(mBase, uint32(v54)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = v66
										v68 = *(*int64)(unsafe.Add(mBase, uint32(v54)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v68
										v70 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
										*(*int64)(unsafe.Add(mBase, uint32(v60))) = v70
										return
									}
								}
							} else {
								return
							}
						}
					}
				}
			default:
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
				if v53 != 0 {
					v54 = F_getRTEPermissionInfo(m, l1, v10)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(0)
						v60 = F_addRTEPermissionInfo(m, l0+int32(36), v10)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							v62 = *(*int64)(unsafe.Add(mBase, uint32(v54)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v62
							v64 = *(*int64)(unsafe.Add(mBase, uint32(v54)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v64
							v66 = *(*int64)(unsafe.Add(mBase, uint32(v54)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = v66
							v68 = *(*int64)(unsafe.Add(mBase, uint32(v54)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v68
							v70 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
							*(*int64)(unsafe.Add(mBase, uint32(v60))) = v70
							return
						}
					}
				} else {
					return
				}
			}
		}
	}
}
func F_getRTEPermissionInfo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v37
			F_errmsg_internal(m, int32(54399), v7)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493158), int32(4016), int32(241173))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v38
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v37
				F_errmsg_internal(m, int32(54399), v7)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493158), int32(4016), int32(241173))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if base.Ui32(v14) < base.Ui32(v11) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v38
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v37
					F_errmsg_internal(m, int32(54399), v7)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493158), int32(4016), int32(241173))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+v11<<(uint(int32(2))%32)-int32(4))))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				if v23 != v24 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v53
						F_errmsg_internal(m, int32(652691), v7+int32(16))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(493158), int32(4021), int32(241173))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					m.G0 = v7 + int32(32)
					return v22
				}
			}
		}
	}
}
