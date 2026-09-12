package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DomainHasConstraints(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = F_lookup_type_cache(m, l0, int32(8192))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+308))
		return base.B2i32(v7 != int32(0))
	}
}
func F_domainAddNotNullConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v12 != 0 {
		v14 = F_ConstraintNameIsUsed(m, int32(1), l0, v12)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if v14 == int32(0) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v43 = v18
				v45 = int32(0)
				v47 = int32(1)
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+15)))
				v65 = int32(32)
				v78 = F_CreateConstraintEntry(m, v43, l1, int32(110), v45, v45, v47, (v48^int32(-1))&v47, v45, v45, v45, v45, v45, l0, v45, v45, v45, v45, v45, v45, v45, v65, v65, v45, v45, v65, v45, v45, v45, v47, v45, v45, v45, v45)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					if l4 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v78
						*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2606)
					} else {
					}
					m.G0 = v10 + int32(16)
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errcode(m, int32(290948))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v26
						F_errmsg(m, int32(123768), v10)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errfinish(m, int32(514319), int32(3684), int32(96454))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
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
		v37 = int32(0)
		v40 = F_ChooseConstraintName(m, l3, v37, int32(314098), l1, v37)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v40
			v43 = v40
			v45 = int32(0)
			v47 = int32(1)
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+15)))
			v65 = int32(32)
			v78 = F_CreateConstraintEntry(m, v43, l1, int32(110), v45, v45, v47, (v48^int32(-1))&v47, v45, v45, v45, v45, v45, l0, v45, v45, v45, v45, v45, v45, v45, v65, v65, v45, v45, v65, v45, v45, v45, v47, v45, v45, v45, v45)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return
			} else {
				if l4 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v78
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2606)
				} else {
				}
				m.G0 = v10 + int32(16)
				return
			}
		}
	}
}
func F_domain_state_setup(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_MemoryContextAlloc(m, l2, int32(84))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v19 = F_lookup_type_cache(m, l0, int32(4096))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+13)))
			if v21 == int32(100) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+300))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+304))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v25
				v28 = v14 + int32(8)
				v30 = v14 + int32(4)
				if l1 != 0 {
					F_getTypeBinaryInputInfo(m, v24, v30, v28)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
						F_fmgr_info_cxt(m, v35, v14+int32(16), l2)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_InitDomainConstraintRef(m, l0, v14+int32(44), l2, int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
								m.G0 = v11 + int32(16)
								return v14
							}
						}
					}
				} else {
					F_getTypeInputInfo(m, v24, v30, v28)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
						F_fmgr_info_cxt(m, v35, v14+int32(16), l2)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_InitDomainConstraintRef(m, l0, v14+int32(44), l2, int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
								m.G0 = v11 + int32(16)
								return v14
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67141764))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						v60 = F_format_type_be(m, l0)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v60
							F_errmsg(m, int32(288863), v11)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(513292), int32(96), int32(242246))
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
				}
			}
		}
	}
}
