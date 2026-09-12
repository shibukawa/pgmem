package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ObjectsInPublicationToOids(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L12
	} else {
		goto L24
	}
L2:
	;
	m.G0 = v10 + int32(16)
	return
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = v4
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v22<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	switch v29 {
	case 0:
		goto L8
	case 1:
		goto L11
	case 2:
		goto L10
	default:
		goto L9
	}
L6:
	;
	goto L2
L7:
	;
	v73 = v22 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v73 < v74 {
		v22 = v73
		goto L5
	} else {
		goto L23
	}
L8:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v67 = F_lappend(m, v65, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L22
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L19
	}
L10:
	;
	v39 = F_fetch_search_path(m, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L15
	}
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v32 = F_get_namespace_oid(m, v30, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v35 = F_list_append_unique_oid(m, v34, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v35
	goto L7
L15:
	;
	if v39 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	F_list_free(m, v39)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v48 = F_list_append_unique_oid(m, v47, v44)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v48
	goto L7
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v55
	F_errmsg_internal(m, int32(457670), v10)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(475273), int32(226), int32(165198))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v67
	goto L7
L23:
	;
	goto L6
L24:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(525013), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(475273), int32(216), int32(165198))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_OffsetVarNodes(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l1
	if l0 == v3 {
		v73 = F_OffsetVarNodes_walker(m, l0, v9+int32(8))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return
		} else {
			m.G0 = v9 + int32(16)
			return
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v16 != int32(67) {
			v73 = F_OffsetVarNodes_walker(m, l0, v9+int32(8))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			if v19 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1 + v19
			} else {
			}
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			if v22 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = l1 + v22
			} else {
			}
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			if v25 == int32(0) {
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
				if v28 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = l1 + v28
				}
			}
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
			if v34 == int32(0) {
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
				if v37 <= int32(0) {
				} else {
					v43 = int32(0)
					for {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v43<<(uint(int32(2))%32))))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v52 + l1
						v56 = v43 + int32(1)
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
						if v56 < v57 {
							v43 = v56
							continue
						} else {
							break
						}
						break
					}
				}
			}
			v69 = F_query_tree_walker_impl(m, l0, int32(1047), v9+int32(8), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
func F_OpclassnameGetOpcid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	v3 = int32(0)
	F_recomputeNamespacePath(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _consts[428]))
	if v11 == int32(0) {
		v47 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v47
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v14 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = v3
	goto L8
L6:
	;
	goto L7
L7:
	;
	v47 = int32(0)
	goto L3
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v21<<(uint(int32(2))%32))))
	v28 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	if v26 != v28 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v32 = F_GetSysCacheOid(m, int32(13), l0, l1, v26, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v36 = v21 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v36 < v37 {
		v21 = v36
		goto L8
	} else {
		goto L15
	}
L13:
	;
	if v32 != 0 {
		v47 = v32
		goto L3
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L9
}
func F_OpenTemporaryFile(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	if l0 != 0 {
		v50 = *(*int32)(unsafe.Add(mBase, _consts[804]))
		if v50 != 0 {
			v52 = v50
		} else {
			v52 = int32(1663)
		}
		v54 = F_OpenTemporaryFileInTablespace(m, v52, int32(1))
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, _consts[413]))
			v62 = v57 + v54*int32(48) + int32(4)
			v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62))))
			v65 = v63 | int32(5)
			*(*uint16)(unsafe.Add(mBase, uint32(v62))) = uint16(v65)
			if l0 != 0 {
				v93 = v54
				return v93
			} else {
				v68 = v54
				v71 = *(*int32)(unsafe.Add(mBase, _consts[170]))
				F_ResourceOwnerRemember(m, v71, v68, int32(1590544))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, _consts[413]))
					v79 = v76 + v68*int32(48)
					v81 = *(*int32)(unsafe.Add(mBase, _consts[170]))
					*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v81
					v84 = v79 + int32(4)
					v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84))))
					v87 = v85 | int32(2)
					*(*uint16)(unsafe.Add(mBase, uint32(v84))) = uint16(v87)
					v90 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _consts[805])) = uint8(v90)
					v93 = v68
					return v93
				}
			}
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, _consts[170]))
		F_ResourceOwnerEnlarge(m, v5)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[806]))
			if v11 <= int32(0) {
				v50 = *(*int32)(unsafe.Add(mBase, _consts[804]))
				if v50 != 0 {
					v52 = v50
				} else {
					v52 = int32(1663)
				}
				v54 = F_OpenTemporaryFileInTablespace(m, v52, int32(1))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, _consts[413]))
					v62 = v57 + v54*int32(48) + int32(4)
					v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62))))
					v65 = v63 | int32(5)
					*(*uint16)(unsafe.Add(mBase, uint32(v62))) = uint16(v65)
					if l0 != 0 {
						v93 = v54
						return v93
					} else {
						v68 = v54
						v71 = *(*int32)(unsafe.Add(mBase, _consts[170]))
						F_ResourceOwnerRemember(m, v71, v68, int32(1590544))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, _consts[413]))
							v79 = v76 + v68*int32(48)
							v81 = *(*int32)(unsafe.Add(mBase, _consts[170]))
							*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v81
							v84 = v79 + int32(4)
							v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84))))
							v87 = v85 | int32(2)
							*(*uint16)(unsafe.Add(mBase, uint32(v84))) = uint16(v87)
							v90 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _consts[805])) = uint8(v90)
							v93 = v68
							return v93
						}
					}
				}
			} else {
				v14 = int32(4365552)
				v16 = *(*int32)(unsafe.Add(mBase, _consts[807]))
				v18 = v16 + int32(1)
				if v18 < v11 {
					v21 = v18
				} else {
					v21 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, _consts[807])) = v21
				v24 = *(*int32)(unsafe.Add(mBase, _consts[808]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v21<<(uint(int32(2))%32))))
				if v28 == int32(0) {
					v50 = *(*int32)(unsafe.Add(mBase, _consts[804]))
					if v50 != 0 {
						v52 = v50
					} else {
						v52 = int32(1663)
					}
					v54 = F_OpenTemporaryFileInTablespace(m, v52, int32(1))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, _consts[413]))
						v62 = v57 + v54*int32(48) + int32(4)
						v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62))))
						v65 = v63 | int32(5)
						*(*uint16)(unsafe.Add(mBase, uint32(v62))) = uint16(v65)
						if l0 != 0 {
							v93 = v54
							return v93
						} else {
							v68 = v54
							v71 = *(*int32)(unsafe.Add(mBase, _consts[170]))
							F_ResourceOwnerRemember(m, v71, v68, int32(1590544))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, _consts[413]))
								v79 = v76 + v68*int32(48)
								v81 = *(*int32)(unsafe.Add(mBase, _consts[170]))
								*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v81
								v84 = v79 + int32(4)
								v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84))))
								v87 = v85 | int32(2)
								*(*uint16)(unsafe.Add(mBase, uint32(v84))) = uint16(v87)
								v90 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[805])) = uint8(v90)
								v93 = v68
								return v93
							}
						}
					}
				} else {
					v32 = F_OpenTemporaryFileInTablespace(m, v28, int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						if v32 <= int32(0) {
							v50 = *(*int32)(unsafe.Add(mBase, _consts[804]))
							if v50 != 0 {
								v52 = v50
							} else {
								v52 = int32(1663)
							}
							v54 = F_OpenTemporaryFileInTablespace(m, v52, int32(1))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, _consts[413]))
								v62 = v57 + v54*int32(48) + int32(4)
								v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62))))
								v65 = v63 | int32(5)
								*(*uint16)(unsafe.Add(mBase, uint32(v62))) = uint16(v65)
								if l0 != 0 {
									v93 = v54
									return v93
								} else {
									v68 = v54
									v71 = *(*int32)(unsafe.Add(mBase, _consts[170]))
									F_ResourceOwnerRemember(m, v71, v68, int32(1590544))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, _consts[413]))
										v79 = v76 + v68*int32(48)
										v81 = *(*int32)(unsafe.Add(mBase, _consts[170]))
										*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v81
										v84 = v79 + int32(4)
										v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84))))
										v87 = v85 | int32(2)
										*(*uint16)(unsafe.Add(mBase, uint32(v84))) = uint16(v87)
										v90 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _consts[805])) = uint8(v90)
										v93 = v68
										return v93
									}
								}
							}
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, _consts[413]))
							v42 = v37 + v32*int32(48) + int32(4)
							v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42))))
							v45 = v43 | int32(5)
							*(*uint16)(unsafe.Add(mBase, uint32(v42))) = uint16(v45)
							v68 = v32
							v71 = *(*int32)(unsafe.Add(mBase, _consts[170]))
							F_ResourceOwnerRemember(m, v71, v68, int32(1590544))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, _consts[413]))
								v79 = v76 + v68*int32(48)
								v81 = *(*int32)(unsafe.Add(mBase, _consts[170]))
								*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v81
								v84 = v79 + int32(4)
								v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84))))
								v87 = v85 | int32(2)
								*(*uint16)(unsafe.Add(mBase, uint32(v84))) = uint16(v87)
								v90 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _consts[805])) = uint8(v90)
								v93 = v68
								return v93
							}
						}
					}
				}
			}
		}
	}
}
func F_OpernameGetOprid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	F_DeconstructQualifiedName(m, l0, v15+int32(12), v15+int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v15 + int32(16)
	return v151
L4:
	;
	v27 = F_LookupExplicitNamespace(m, v25, int32(1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v46 = F_SearchSysCacheList(m, int32(39), int32(3), v45, l1, l2)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	if v27 == int32(0) {
		v151 = v4
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v33 = F_SearchSysCache4(m, int32(39), v32, l1, l2, v27)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v33 == int32(0) {
		v151 = v4
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38)))
	F_ReleaseCatCache(m, v33)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v151 = v40
	goto L3
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v48 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_ReleaseCatCacheList(m, v46)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v151 = v4
	goto L3
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[428]))
	if v56 == int32(0) {
		v137 = v4
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_ReleaseCatCacheList(m, v46)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L36
	}
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v59 <= int32(0) {
		v137 = v4
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v62 = int32(0)
	if v62 < v59 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v65 = v59
	goto L23
L22:
	;
	v65 = v62
	goto L23
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v77 = v4
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v70+v77<<(uint(int32(2))%32))))
	if v86 == v69 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v137 = v4
	goto L18
L26:
	;
	v130 = v77 + int32(1)
	if v130 != v65 {
		v77 = v130
		goto L24
	} else {
		goto L35
	}
L27:
	;
	v88 = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v89 <= v88 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v92 = v88
	goto L29
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(48)+v92<<(uint(int32(2))%32))))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+56))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+22)))
	v110 = v108 + v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+68))
	if v86 != v111 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v137 = v116
	goto L18
L31:
	;
	v114 = v92 + int32(1)
	if v89 != v114 {
		v92 = v114
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	goto L26
L35:
	;
	goto L25
L36:
	;
	v151 = v137
	goto L3
}
func F_oauth_exchange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int64
	_ = v228
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v445 int32
	_ = v445
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v488 int32
	_ = v488
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int64
	_ = v519
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1073 int32
	_ = v1073
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1125 int32
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1211 int32
	_ = v1211
	v7 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(128)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(-1)
	if l1 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(128)
	return v1211
L2:
	;
	v24 = F_pstrdup(m, int32(719562))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v24
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v1211 = v7
	goto L1
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1200
	v1203 = F___memset(m, v111, int32(0), l2)
	mBase = m.M
	goto L365
L8:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1104 != 0 {
		goto L342
	} else {
		goto L343
	}
L9:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L5
	} else {
		goto L337
	}
L10:
	;
	if l1&int32(3) == int32(0) {
		v54 = l1
		goto L15
	} else {
		goto L16
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L5
	} else {
		goto L332
	}
L13:
	;
	if v87 == l2 {
		goto L30
	} else {
		goto L31
	}
L14:
	;
	v87 = v79 - l1
	goto L13
L15:
	;
	v58 = v54
	goto L24
L16:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v38 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v87 = int32(0)
	goto L13
L18:
	;
	goto L19
L19:
	;
	v43 = l1
	goto L20
L20:
	;
	v47 = v43 + int32(1)
	if v47&int32(3) == int32(0) {
		v54 = v47
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v79 = v47
	goto L14
L22:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v52 != 0 {
		v43 = v47
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v67 = int32(-2139062144)
	if (int32(16843008)-v64|v64)&v67 == v67 {
		v58 = v58 + int32(4)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v73 = v58
	goto L27
L26:
	;
	goto L25
L27:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v77 != 0 {
		v73 = v73 + int32(1)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v79 = v73
	goto L14
L29:
	;
	goto L28
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v89 {
	case 0:
		goto L34
	case 1:
		goto L36
	default:
		goto L35
	}
L31:
	;
	goto L32
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L5
	} else {
		goto L327
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L5
	} else {
		goto L322
	}
L34:
	;
	v111 = F_pstrdup(m, l1)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L46
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L39
	}
L36:
	;
	if l2 != int32(1) {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v92 != int32(1) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v95 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v95
	v1211 = v95
	goto L1
L39:
	;
	F_errmsg_internal(m, int32(338763), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(478550), int32(200), int32(386978))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L5
	} else {
		goto L316
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L5
	} else {
		goto L310
	}
L44:
	;
	v137 = v111 + int32(1)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v138 != int32(44) {
		goto L42
	} else {
		goto L52
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L47
	}
L46:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	switch v113 - int32(110) {
	case 0, 11:
		goto L44
	default:
		goto L43
	case 2:
		goto L45
	}
L47:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	F_errdetail(m, int32(617633), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(478550), int32(221), int32(386978))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+2)))
	if v141 != int32(44) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L5
	} else {
		goto L306
	}
L54:
	;
	if v141 == int32(97) {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v173 = v111 + int32(3)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	if v174 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	v157 = int32(*(*int8)(unsafe.Add(mBase, uint32(v111)+2)))
	F_sanitize_char_1(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(4348220)
	F_errdetail(m, int32(605335), v15+int32(16))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(478550), int32(256), int32(386978))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v178 = v111 + int32(4)
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v179 != 0 {
		goto L74
	} else {
		goto L75
	}
L65:
	;
	goto L66
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L5
	} else {
		goto L300
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L5
	} else {
		goto L295
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L5
	} else {
		goto L290
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L5
	} else {
		goto L285
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L5
	} else {
		goto L280
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L5
	} else {
		goto L275
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L5
	} else {
		goto L270
	}
L73:
	;
	if v186 != 0 {
		goto L143
	} else {
		goto L144
	}
L74:
	;
	v186 = int32(0)
	v189 = v178
	goto L77
L75:
	;
	goto L76
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L5
	} else {
		goto L138
	}
L77:
	;
	v193 = int32(1)
	v194 = F___strchrnul(m, v189, v193)
	mBase = m.M
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v196 == v193 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L76
L79:
	;
	if v200 == int32(0) {
		goto L67
	} else {
		goto L83
	}
L80:
	;
	v200 = v194
	goto L82
L81:
	;
	v200 = int32(0)
	goto L82
L82:
	;
	goto L79
L83:
	;
	v203 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v200))) = uint8(v203)
	if v189 == v200 {
		goto L73
	} else {
		goto L84
	}
L84:
	;
	v206 = int32(61)
	v207 = F___strchrnul(m, v189, v206)
	mBase = m.M
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v209 == v206 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v213 == int32(0) {
		goto L68
	} else {
		goto L89
	}
L86:
	;
	v213 = v207
	goto L88
L87:
	;
	v213 = int32(0)
	goto L88
L88:
	;
	goto L85
L89:
	;
	v216 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v216)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v218 == v216 {
		goto L69
	} else {
		goto L90
	}
L90:
	;
	v221 = int32(488274)
	v225 = m.G0
	v227 = v225 - int32(32)
	v228 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v227)+24)) = v228
	*(*int64)(unsafe.Add(mBase, uint32(v227)+16)) = v228
	*(*int64)(unsafe.Add(mBase, uint32(v227)+8)) = v228
	*(*int64)(unsafe.Add(mBase, uint32(v227))) = v228
	v236 = int32(*(*uint8)(unsafe.Add(mBase, _consts[546])))
	if v236 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304+v189))))
	if v306 != 0 {
		goto L70
	} else {
		goto L112
	}
L92:
	;
	v304 = int32(0)
	goto L91
L93:
	;
	goto L94
L94:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, _consts[547])))
	if v240 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v244 = v189
	goto L98
L96:
	;
	goto L97
L97:
	;
	v254 = v221
	v255 = v236
	goto L101
L98:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	if v250 == v236 {
		v244 = v244 + int32(1)
		goto L98
	} else {
		goto L100
	}
L99:
	;
	v304 = v244 - v189
	goto L91
L100:
	;
	goto L99
L101:
	;
	v262 = v227 + int32(base.Ui32(v255)>>(uint(int32(3))%32))&int32(28)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	v264 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v263 | v264<<(uint(v255)%32)
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)))
	if v268 != 0 {
		v254 = v254 + v264
		v255 = v268
		goto L101
	} else {
		goto L103
	}
L102:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v271 == int32(0) {
		v296 = v189
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L102
L104:
	;
	v304 = v296 - v189
	goto L91
L105:
	;
	v275 = v189
	v276 = v271
	goto L106
L106:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v227+int32(base.Ui32(v276)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v284)>>(uint(v276)%32))&int32(1) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v296 = v292
	goto L104
L108:
	;
	v296 = v275
	goto L104
L109:
	;
	goto L110
L110:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+1)))
	v292 = v275 + int32(1)
	if v290 != 0 {
		v275 = v292
		v276 = v290
		goto L106
	} else {
		goto L111
	}
L111:
	;
	goto L107
L112:
	;
	v308 = v213 + int32(1)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	if v309 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v311 = v309
	v317 = v308
	goto L116
L114:
	;
	goto L115
L115:
	;
	v356 = int32(307479)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, _consts[548])))
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v360 == int32(0) {
		v379 = v359
		v380 = v360
		goto L126
	} else {
		goto L127
	}
L116:
	;
	if base.Ui32((v311-int32(127))&int32(255)) <= base.Ui32(int32(161)) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L115
L118:
	;
	v331 = v311&int32(255) - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v331) {
		goto L71
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v342 = v317 + int32(1)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	if v343 != 0 {
		v311 = v343
		v317 = v342
		goto L116
	} else {
		goto L123
	}
L121:
	;
	if int32(1)<<(uint(v331)%32)&int32(8388627) == int32(0) {
		goto L71
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	goto L117
L124:
	;
	v384 = v200 + int32(1)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v385 != 0 {
		v186 = v382
		v189 = v384
		goto L77
	} else {
		goto L137
	}
L125:
	;
	if v380-v379 != 0 {
		goto L133
	} else {
		goto L134
	}
L126:
	;
	goto L125
L127:
	;
	if v359 != v360 {
		v379 = v359
		v380 = v360
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v364 = v189
	v365 = v356
	goto L129
L129:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+1)))
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+1)))
	if v369 == int32(0) {
		v379 = v368
		v380 = v369
		goto L126
	} else {
		goto L131
	}
L130:
	;
	v379 = v368
	v380 = v369
	goto L126
L131:
	;
	v372 = int32(1)
	if v368 == v369 {
		v364 = v364 + v372
		v365 = v365 + v372
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v382 = v186
	goto L124
L134:
	;
	goto L135
L135:
	;
	if v186 != 0 {
		goto L72
	} else {
		goto L136
	}
L136:
	;
	v382 = v308
	goto L124
L137:
	;
	goto L78
L138:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L5
	} else {
		goto L139
	}
L139:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L5
	} else {
		goto L140
	}
L140:
	;
	F_errdetail(m, int32(570392), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(478550), int32(473), int32(307461))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L5
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	if v418 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	goto L145
L145:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L5
	} else {
		goto L265
	}
L146:
	;
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v421 == int32(0) {
		goto L8
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L5
	} else {
		goto L260
	}
L149:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v429 = v186
	v430 = int32(700924)
	v431 = int32(7)
	goto L151
L150:
	;
	if v476 != 0 {
		goto L166
	} else {
		goto L167
	}
L151:
	;
	if v431 != 0 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v476 = int32(0)
	goto L150
L153:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	if v434 == v435 {
		v457 = v434
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L155
L155:
	;
	goto L152
L156:
	;
	v459 = int32(1)
	if v457 != 0 {
		v429 = v429 + v459
		v430 = v430 + v459
		v431 = v431 - v459
		goto L151
	} else {
		goto L165
	}
L157:
	;
	if base.Ui32((v434-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v445 = v434 | int32(32)
	goto L160
L159:
	;
	v445 = v434
	goto L160
L160:
	;
	if base.Ui32((v435-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v454 = v435 | int32(32)
	goto L163
L162:
	;
	v454 = v435
	goto L163
L163:
	;
	if v445 == v454 {
		v457 = v445
		goto L156
	} else {
		goto L164
	}
L164:
	;
	v476 = v445 - v454
	goto L150
L165:
	;
	goto L155
L166:
	;
	v479 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L5
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v488 = v186 + int32(7)
	goto L171
L169:
	;
	if v479 == int32(0) {
		goto L8
	} else {
		goto L170
	}
L170:
	;
	v1067 = int32(586)
	v1073 = int32(598685)
	goto L9
L171:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	if v499 != int32(32) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	if v499 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v488 = v488 + int32(1)
	goto L171
L176:
	;
	v506 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L5
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v512 = int32(534718)
	v516 = m.G0
	v518 = v516 - int32(32)
	v519 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v518)+24)) = v519
	*(*int64)(unsafe.Add(mBase, uint32(v518)+16)) = v519
	*(*int64)(unsafe.Add(mBase, uint32(v518)+8)) = v519
	*(*int64)(unsafe.Add(mBase, uint32(v518))) = v519
	v527 = int32(*(*uint8)(unsafe.Add(mBase, _consts[549])))
	if v527 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	if v506 == int32(0) {
		goto L8
	} else {
		goto L180
	}
L180:
	;
	v1067 = int32(603)
	v1073 = int32(534802)
	goto L9
L181:
	;
	v603 = v595
	goto L202
L182:
	;
	v595 = int32(0)
	goto L181
L183:
	;
	goto L184
L184:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, _consts[550])))
	if v531 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v535 = v488
	goto L188
L186:
	;
	goto L187
L187:
	;
	v545 = v512
	v546 = v527
	goto L191
L188:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	if v541 == v527 {
		v535 = v535 + int32(1)
		goto L188
	} else {
		goto L190
	}
L189:
	;
	v595 = v535 - v488
	goto L181
L190:
	;
	goto L189
L191:
	;
	v553 = v518 + int32(base.Ui32(v546)>>(uint(int32(3))%32))&int32(28)
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	v555 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v553))) = v554 | v555<<(uint(v546)%32)
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+1)))
	if v559 != 0 {
		v545 = v545 + v555
		v546 = v559
		goto L191
	} else {
		goto L193
	}
L192:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	if v562 == int32(0) {
		v587 = v488
		goto L194
	} else {
		goto L195
	}
L193:
	;
	goto L192
L194:
	;
	v595 = v587 - v488
	goto L181
L195:
	;
	v566 = v488
	v567 = v562
	goto L196
L196:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v518+int32(base.Ui32(v567)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v575)>>(uint(v567)%32))&int32(1) == int32(0) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v587 = v583
	goto L194
L198:
	;
	v587 = v566
	goto L194
L199:
	;
	goto L200
L200:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566)+1)))
	v583 = v566 + int32(1)
	if v581 != 0 {
		v566 = v583
		v567 = v581
		goto L196
	} else {
		goto L201
	}
L201:
	;
	goto L197
L202:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488+v603))))
	if v609 != int32(61) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L5
	} else {
		goto L256
	}
L204:
	;
	if v609 != 0 {
		goto L208
	} else {
		goto L209
	}
L205:
	;
	v603 = v603 + int32(1)
	goto L202
L206:
	;
	goto L203
L207:
	;
	goto L206
L208:
	;
	v614 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L5
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v621 = *(*int32)(unsafe.Add(mBase, _consts[551]))
	if v621 == int32(0) {
		goto L207
	} else {
		goto L213
	}
L211:
	;
	if v614 == int32(0) {
		goto L8
	} else {
		goto L212
	}
L212:
	;
	v1067 = int32(625)
	v1073 = int32(546687)
	goto L9
L213:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v621)+12))
	if v624 == int32(0) {
		goto L207
	} else {
		goto L214
	}
L214:
	;
	v628 = F_palloc0(m, int32(8))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L5
	} else {
		goto L215
	}
L215:
	;
	v631 = *(*int32)(unsafe.Add(mBase, _consts[552]))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v424)+364))
	v634 = *(*int32)(unsafe.Add(mBase, _consts[551]))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)+12))
	v636 = m.T0[v635].(func(*base.Module, int32, int32, int32, int32) int32)(m, v631, v488, v632, v628)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L5
	} else {
		goto L216
	}
L216:
	;
	if v636 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v642 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L5
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v628)+4))
	if v658 != 0 {
		goto L225
	} else {
		goto L226
	}
L220:
	;
	if v642 == int32(0) {
		goto L8
	} else {
		goto L221
	}
L221:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L5
	} else {
		goto L222
	}
L222:
	;
	F_errmsg(m, int32(367258), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L5
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(478550), int32(665), int32(342514))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L5
	} else {
		goto L224
	}
L224:
	;
	goto L8
L225:
	;
	F_set_authn_id(m, v424, v658)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L5
	} else {
		goto L228
	}
L226:
	;
	goto L227
L227:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628))))
	if v661 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L228:
	;
	goto L227
L229:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v628)+4))
	if v725 != 0 {
		goto L250
	} else {
		goto L251
	}
L230:
	;
	v664 = int32(0)
	v667 = F_errstart(m, int32(15), v664)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L5
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v424)+380))
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v688)+416)))
	if v689 != 0 {
		v723 = int32(1)
		goto L229
	} else {
		goto L238
	}
L233:
	;
	if v667 == int32(0) {
		v723 = v664
		goto L229
	} else {
		goto L234
	}
L234:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v424)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v671
	F_errmsg(m, int32(663680), v15+int32(48))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L5
	} else {
		goto L235
	}
L235:
	;
	F_errdetail_log(m, int32(583399), int32(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L5
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(478550), int32(681), int32(342514))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L5
	} else {
		goto L237
	}
L237:
	;
	v723 = v664
	goto L229
L238:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v628)+4))
	if v690 != 0 {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v688)+300))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v424)+364))
	v718 = *(*int32)(unsafe.Add(mBase, _consts[118]))
	v719 = F_check_usermap(m, v715, v716, v718)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L5
	} else {
		goto L249
	}
L240:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690))))
	if v691 != 0 {
		goto L239
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v692 = int32(0)
	v695 = F_errstart(m, int32(15), v692)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L5
	} else {
		goto L244
	}
L243:
	;
	goto L242
L244:
	;
	if v695 == int32(0) {
		v723 = v692
		goto L229
	} else {
		goto L245
	}
L245:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v424)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v699
	F_errmsg(m, int32(663680), v15+int32(32))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	F_errdetail_log(m, int32(535417), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L5
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(478550), int32(705), int32(342514))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L5
	} else {
		goto L248
	}
L248:
	;
	v723 = v692
	goto L229
L249:
	;
	v723 = base.B2i32(v719 == int32(0))
	goto L229
L250:
	;
	F_pfree(m, v725)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L5
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	F_pfree(m, v628)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L5
	} else {
		goto L254
	}
L253:
	;
	goto L252
L254:
	;
	if v723 == int32(0) {
		goto L8
	} else {
		goto L255
	}
L255:
	;
	v1195 = int32(1)
	v1200 = int32(2)
	goto L7
L256:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L5
	} else {
		goto L257
	}
L257:
	;
	F_errmsg(m, int32(445294), int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L5
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(478550), int32(656), int32(342514))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L5
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L5
	} else {
		goto L261
	}
L261:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L5
	} else {
		goto L262
	}
L262:
	;
	F_errdetail(m, int32(570331), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L5
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(478550), int32(280), int32(386978))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L5
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L5
	} else {
		goto L266
	}
L266:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L5
	} else {
		goto L267
	}
L267:
	;
	F_errdetail(m, int32(592244), int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L5
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(478550), int32(273), int32(386978))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L5
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L5
	} else {
		goto L271
	}
L271:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L5
	} else {
		goto L272
	}
L272:
	;
	F_errdetail(m, int32(559884), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L5
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(478550), int32(453), int32(307461))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L5
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L5
	} else {
		goto L276
	}
L276:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L5
	} else {
		goto L277
	}
L277:
	;
	F_errdetail(m, int32(592351), int32(0))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L5
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(478550), int32(376), int32(204430))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L5
	} else {
		goto L279
	}
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L5
	} else {
		goto L281
	}
L281:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L5
	} else {
		goto L282
	}
L282:
	;
	F_errdetail(m, int32(598987), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L5
	} else {
		goto L283
	}
L283:
	;
	F_errfinish(m, int32(478550), int32(350), int32(204430))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L5
	} else {
		goto L284
	}
L284:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L285:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L5
	} else {
		goto L286
	}
L286:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L5
	} else {
		goto L287
	}
L287:
	;
	F_errdetail(m, int32(598951), int32(0))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L5
	} else {
		goto L288
	}
L288:
	;
	F_errfinish(m, int32(478550), int32(343), int32(204430))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L5
	} else {
		goto L289
	}
L289:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L290:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L5
	} else {
		goto L291
	}
L291:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L5
	} else {
		goto L292
	}
L292:
	;
	F_errdetail(m, int32(592386), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L5
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(478550), int32(439), int32(307461))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L5
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L5
	} else {
		goto L297
	}
L297:
	;
	F_errdetail(m, int32(571001), int32(0))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L5
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(478550), int32(421), int32(307461))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L5
	} else {
		goto L299
	}
L299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L300:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L5
	} else {
		goto L301
	}
L301:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L5
	} else {
		goto L302
	}
L302:
	;
	v925 = int32(*(*int8)(unsafe.Add(mBase, uint32(v173))))
	F_sanitize_char_1(m, v925)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L5
	} else {
		goto L303
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = int32(4348220)
	F_errdetail(m, int32(628389), v15-int32(-64))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L5
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(478550), int32(265), int32(386978))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L5
	} else {
		goto L305
	}
L305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L306:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L5
	} else {
		goto L307
	}
L307:
	;
	F_errmsg(m, int32(424625), int32(0))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L5
	} else {
		goto L308
	}
L308:
	;
	F_errfinish(m, int32(478550), int32(250), int32(386978))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L5
	} else {
		goto L309
	}
L309:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L310:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L5
	} else {
		goto L311
	}
L311:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L5
	} else {
		goto L312
	}
L312:
	;
	F_sanitize_char_1(m, base.I32_extend8_s(v113))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L5
	} else {
		goto L313
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(4348220)
	F_errdetail(m, int32(629453), v15)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L5
	} else {
		goto L314
	}
L314:
	;
	F_errfinish(m, int32(478550), int32(241), int32(386978))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L5
	} else {
		goto L315
	}
L315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L316:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L5
	} else {
		goto L317
	}
L317:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L5
	} else {
		goto L318
	}
L318:
	;
	v991 = int32(*(*int8)(unsafe.Add(mBase, uint32(v137))))
	F_sanitize_char_1(m, v991)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L5
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = int32(4348220)
	F_errdetail(m, int32(628445), v15+int32(80))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L5
	} else {
		goto L320
	}
L320:
	;
	F_errfinish(m, int32(478550), int32(232), int32(386978))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L5
	} else {
		goto L321
	}
L321:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L322:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L5
	} else {
		goto L323
	}
L323:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L5
	} else {
		goto L324
	}
L324:
	;
	F_errdetail(m, int32(595468), int32(0))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L5
	} else {
		goto L325
	}
L325:
	;
	F_errfinish(m, int32(478550), int32(193), int32(386978))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L5
	} else {
		goto L326
	}
L326:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L327:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L5
	} else {
		goto L328
	}
L328:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L5
	} else {
		goto L329
	}
L329:
	;
	F_errdetail(m, int32(587570), int32(0))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L5
	} else {
		goto L330
	}
L330:
	;
	F_errfinish(m, int32(478550), int32(175), int32(386978))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L5
	} else {
		goto L331
	}
L331:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L332:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L5
	} else {
		goto L333
	}
L333:
	;
	F_errmsg(m, int32(389235), int32(0))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L5
	} else {
		goto L334
	}
L334:
	;
	F_errdetail(m, int32(534825), int32(0))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L5
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(478550), int32(170), int32(386978))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L5
	} else {
		goto L336
	}
L336:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L337:
	;
	F_errmsg(m, int32(270454), int32(0))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L5
	} else {
		goto L338
	}
L338:
	;
	F_errdetail_log(m, v1073, int32(0))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L5
	} else {
		goto L339
	}
L339:
	;
	F_errfinish(m, int32(478550), v1067, int32(105353))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L5
	} else {
		goto L340
	}
L340:
	;
	goto L8
L341:
	;
	F_initStringInfo(m, v15+int32(96))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L5
	} else {
		goto L351
	}
L342:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1105 != 0 {
		goto L341
	} else {
		goto L345
	}
L343:
	;
	goto L344
L344:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L5
	} else {
		goto L346
	}
L345:
	;
	goto L344
L346:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L5
	} else {
		goto L347
	}
L347:
	;
	F_errmsg(m, int32(208607), int32(0))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L5
	} else {
		goto L348
	}
L348:
	;
	F_errdetail_log(m, int32(591090), int32(0))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L5
	} else {
		goto L349
	}
L349:
	;
	F_errfinish(m, int32(478550), int32(500), int32(346913))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L5
	} else {
		goto L350
	}
L350:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L351:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_appendStringInfoString(m, v15+int32(96), v1132)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L5
	} else {
		goto L352
	}
L352:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1137 = F_strstr(m, v1135, int32(534655))
	mBase = m.M
	if v1137 == int32(0) {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	F_appendStringInfoString(m, v15+int32(96), int32(248668))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L5
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	F_initStringInfo(m, v15+int32(112))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L5
	} else {
		goto L357
	}
L356:
	;
	goto L355
L357:
	;
	F_appendStringInfoString(m, v15+int32(112), int32(708299))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L5
	} else {
		goto L358
	}
L358:
	;
	F_appendStringInfoString(m, v15+int32(112), int32(708034))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L5
	} else {
		goto L359
	}
L359:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	F_escape_json(m, v15+int32(112), v1161)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L5
	} else {
		goto L360
	}
L360:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	F_pfree(m, v1164)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L5
	} else {
		goto L361
	}
L361:
	;
	F_appendStringInfoString(m, v15+int32(112), int32(708070))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L5
	} else {
		goto L362
	}
L362:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_escape_json(m, v15+int32(112), v1174)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L5
	} else {
		goto L363
	}
L363:
	;
	F_appendStringInfoString(m, v15+int32(112), int32(6925))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L5
	} else {
		goto L364
	}
L364:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1182
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v15)+116))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1184
	v1195 = int32(0)
	v1200 = int32(1)
	goto L7
L365:
	;
	v1211 = v1195
	goto L1
}
func F_offset_elem_desc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	F_appendStringInfo(m, l0, int32(57422), v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func F_oidin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_uint32in_subr(m, v2, int32(0), int32(418045), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_oidlarger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.Ui32(v4) < base.Ui32(v3) {
		v6 = v3
	} else {
		v6 = v4
	}
	return v6
}
func F_oidvectoreq(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_btoidvectorcmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v2 == int32(0))
	}
}
func F_oidvectorout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L20
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v16 != int32(26) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v24 = F_palloc(m, v19*int32(12)|int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if int32(0) < v19 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v33 = v24
	v34 = int32(0)
	goto L10
L8:
	;
	v64 = v24
	goto L9
L9:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v70)
	m.G0 = v9 + int32(16)
	return v24
L10:
	;
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v64 = v59
	goto L9
L12:
	;
	v39 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v39)
	v43 = v33 + int32(1)
	goto L14
L13:
	;
	v43 = v33
	goto L14
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(24)+v34<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v47
	v50 = F_pg_sprintf(m, v43, int32(57422), v9)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v52 = v43
	goto L16
L16:
	;
	v59 = v52 + int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v60 != 0 {
		v52 = v59
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v62 = v34 + int32(1)
	if v62 != v19 {
		v33 = v59
		v34 = v62
		goto L10
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	goto L11
L20:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(199547), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(480421), int32(131), int32(199525))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_oidvectorrecv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+13)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v9
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+44)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+36)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(26)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+28)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v8
	v26 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+22)) = uint16(v26)
	v30 = F_array_recv(m, v6+int32(4))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
		if v34 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(485043), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(480421), int32(245), int32(33976))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
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
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
			if v37 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50462850))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(485043), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(480421), int32(245), int32(33976))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
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
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				if v38 != int32(26) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50462850))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(485043), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(480421), int32(245), int32(33976))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
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
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
					if v41 == int32(0) {
						m.G0 = v6 + int32(48)
						return v30
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50462850))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(485043), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(480421), int32(245), int32(33976))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
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
}
func F_okeys_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+32))
	if v5 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 <= v9 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v8 << (uint(int32(1)) % 32)
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v17 = F_repalloc(m, v14, v8<<(uint(int32(3))%32))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17
				v22 = F_pstrdup(m, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v24 + int32(1)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v28+v24<<(uint(int32(2))%32)))) = v22
					return int32(0)
				}
			}
		} else {
			v22 = F_pstrdup(m, l1)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v24 + int32(1)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v28+v24<<(uint(int32(2))%32)))) = v22
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func F_order_qual_clauses(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 float64
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 float64
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 float64
	_ = v111
	var v112 int64
	_ = v112
	var v115 int32
	_ = v115
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 float64
	_ = v135
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l1 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v191
L2:
	;
	v191 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 < int32(2) {
		v191 = l1
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v25 = F_palloc(m, v20*int32(24))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v29 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v34 = v3
	goto L11
L9:
	;
	goto L10
L10:
	;
	v89 = int32(2)
	if v20 <= v89 {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v34<<(uint(int32(2))%32))))
	F_cost_qual_eval_node(m, v15, v48, l0)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v53 = v25 + v34*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v48
	v55 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v53)+8)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v58 != int32(318) {
		v71 = int32(0)
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v71
	v74 = v34 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v74 < v75 {
		v34 = v74
		goto L11
	} else {
		goto L20
	}
L15:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+13)))
	if v61 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v66 = *(*float64)(unsafe.Add(mBase, _consts[588]))
	if base.F64_lt(v55, base.F64_mul(v66, float64(10))) != 0 {
		v71 = int32(0)
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	v71 = v70
	goto L14
L19:
	;
	goto L18
L20:
	;
	goto L12
L21:
	;
	v92 = v89
	goto L23
L22:
	;
	v92 = v20
	goto L23
L23:
	;
	v101 = int32(1)
	goto L24
L24:
	;
	v108 = v25 + v101*int32(24)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	v111 = *(*float64)(unsafe.Add(mBase, uint32(v108)+8))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v108)))
	v115 = v101
	goto L27
L25:
	;
	v163 = int32(1)
	if v20 <= v163 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v155 = v25 + v151*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v155)+20)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v155)+16)) = v110
	*(*float64)(unsafe.Add(mBase, uint32(v155)+8)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v155))) = v112
	v161 = v101 + int32(1)
	if v161 != v92 {
		v101 = v161
		goto L24
	} else {
		goto L35
	}
L27:
	;
	v127 = v25 + v115*int32(24)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127-int32(8))))
	if base.Ui32(v130) < base.Ui32(v110) {
		v151 = v115
		goto L26
	} else {
		goto L29
	}
L28:
	;
	v151 = int32(0)
	goto L26
L29:
	;
	if v130 == v110 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v127-int32(16))))
	if base.F64_ge(v111, v135) != 0 {
		v151 = v115
		goto L26
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v138 = v127 - int32(24)
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v139
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v138)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v127)+16)) = v141
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v127)+8)) = v143
	v145 = int32(1)
	if v145 < v115 {
		v115 = v115 - v145
		goto L27
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	goto L28
L35:
	;
	goto L25
L36:
	;
	v166 = v163
	goto L38
L37:
	;
	v166 = v20
	goto L38
L38:
	;
	v167 = int32(0)
	v170 = v167
	v171 = v167
	goto L39
L39:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v25+v171*int32(24))))
	v185 = F_lappend(m, v170, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L6
	} else {
		goto L41
	}
L40:
	;
	v191 = v185
	goto L1
L41:
	;
	v188 = v171 + int32(1)
	if v188 != v166 {
		v170 = v185
		v171 = v188
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
}
func F_owningrel_does_not_exist_skipping(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	if l0 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = v5 - int32(1)
	} else {
		v9 = int32(-1)
	}
	v10 = F_list_copy_head(m, l0, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_makeRangeVarFromNameList(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
			if v16 == int32(0) {
				v25 = F_makeRangeVarFromNameList(m, v10)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = int32(0)
					v31 = F_RangeVarGetRelidExtended(m, v25, v27, int32(1), v27, v27)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 != 0 {
							v41 = int32(0)
							return v41
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(318726)
							v35 = F_NameListToString(m, v10)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v38 = v35
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v38
								v41 = int32(1)
								return v41
							}
						}
					}
				}
			} else {
				v19 = F_LookupNamespaceNoError(m, v16)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v19 != 0 {
						v25 = F_makeRangeVarFromNameList(m, v10)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = int32(0)
							v31 = F_RangeVarGetRelidExtended(m, v25, v27, int32(1), v27, v27)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								if v31 != 0 {
									v41 = int32(0)
									return v41
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(318726)
									v35 = F_NameListToString(m, v10)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										v38 = v35
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v38
										v41 = int32(1)
										return v41
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(319351)
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
						v38 = v23
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v38
						v41 = int32(1)
						return v41
					}
				}
			}
		}
	}
}
