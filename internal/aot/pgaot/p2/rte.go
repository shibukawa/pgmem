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
	var v11 int32
	_ = v11
	var v17 int64
	_ = v17
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	v7 = F_palloc(m, int32(136))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		base.MemoryCopy(m, v7, l2, int32(136))
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+128)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+120)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+104)) = v11
		v17 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+96)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v7)+76)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v7)+60)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = v11
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v30 = F_lappend(m, v29, v7)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v30
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			switch v33 {
			case 0:
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
				v38 = v34
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v40 = F_lappend_oid(m, v39, v38)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v40
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v44 != 0 {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
						v47 = v45
					} else {
						v47 = int32(0)
					}
					v48 = F_bms_add_member(m, v43, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
						if v52 != 0 {
							v53 = F_getRTEPermissionInfo(m, l1, v7)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
								v59 = F_addRTEPermissionInfo(m, l0+int32(36), v7)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									v61 = *(*int64)(unsafe.Add(mBase, uint32(v53)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v59)+32)) = v61
									v63 = *(*int64)(unsafe.Add(mBase, uint32(v53)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v59)+24)) = v63
									v65 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
									*(*int64)(unsafe.Add(mBase, uint32(v59)+16)) = v65
									v67 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v59)+8)) = v67
									v69 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
									*(*int64)(unsafe.Add(mBase, uint32(v59))) = v69
									return
								}
							}
						} else {
							return
						}
					}
				}
			case 1:
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
				if v35 == int32(0) {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
					if v52 != 0 {
						v53 = F_getRTEPermissionInfo(m, l1, v7)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
							v59 = F_addRTEPermissionInfo(m, l0+int32(36), v7)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v61 = *(*int64)(unsafe.Add(mBase, uint32(v53)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v59)+32)) = v61
								v63 = *(*int64)(unsafe.Add(mBase, uint32(v53)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v59)+24)) = v63
								v65 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v59)+16)) = v65
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v59)+8)) = v67
								v69 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
								*(*int64)(unsafe.Add(mBase, uint32(v59))) = v69
								return
							}
						}
					} else {
						return
					}
				} else {
					v38 = v35
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v40 = F_lappend_oid(m, v39, v38)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v40
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v44 != 0 {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
							v47 = v45
						} else {
							v47 = int32(0)
						}
						v48 = F_bms_add_member(m, v43, v47)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v48
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
							if v52 != 0 {
								v53 = F_getRTEPermissionInfo(m, l1, v7)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
									v59 = F_addRTEPermissionInfo(m, l0+int32(36), v7)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										v61 = *(*int64)(unsafe.Add(mBase, uint32(v53)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v59)+32)) = v61
										v63 = *(*int64)(unsafe.Add(mBase, uint32(v53)+24))
										*(*int64)(unsafe.Add(mBase, uint32(v59)+24)) = v63
										v65 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
										*(*int64)(unsafe.Add(mBase, uint32(v59)+16)) = v65
										v67 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v59)+8)) = v67
										v69 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
										*(*int64)(unsafe.Add(mBase, uint32(v59))) = v69
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
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
				if v52 != 0 {
					v53 = F_getRTEPermissionInfo(m, l1, v7)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
						v59 = F_addRTEPermissionInfo(m, l0+int32(36), v7)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							v61 = *(*int64)(unsafe.Add(mBase, uint32(v53)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v59)+32)) = v61
							v63 = *(*int64)(unsafe.Add(mBase, uint32(v53)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v59)+24)) = v63
							v65 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v59)+16)) = v65
							v67 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v59)+8)) = v67
							v69 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
							*(*int64)(unsafe.Add(mBase, uint32(v59))) = v69
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
			F_errmsg_internal(m, int32(_a_F_getRTEPermissionInfo_0), v7)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_getRTEPermissionInfo_1), int32(4016), int32(_a_F_getRTEPermissionInfo_2))
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
				F_errmsg_internal(m, int32(_a_F_getRTEPermissionInfo_0), v7)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_getRTEPermissionInfo_1), int32(4016), int32(_a_F_getRTEPermissionInfo_2))
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
					F_errmsg_internal(m, int32(_a_F_getRTEPermissionInfo_0), v7)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_getRTEPermissionInfo_1), int32(4016), int32(_a_F_getRTEPermissionInfo_2))
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
						F_errmsg_internal(m, int32(_a_F_getRTEPermissionInfo_3), v7+int32(16))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_getRTEPermissionInfo_1), int32(4021), int32(_a_F_getRTEPermissionInfo_2))
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
