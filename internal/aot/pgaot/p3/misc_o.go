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
	F_errmsg_internal(m, int32(473085), v10)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(491249), int32(226), int32(172770))
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
	F_errmsg(m, int32(541761), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(491249), int32(216), int32(172770))
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
			v69 = F_query_tree_walker_impl(m, l0, int32(1048), v9+int32(8), int32(0))
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
				F_ResourceOwnerRemember(m, v71, v68, int32(1609712))
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
						F_ResourceOwnerRemember(m, v71, v68, int32(1609712))
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
				v14 = int32(4403072)
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
							F_ResourceOwnerRemember(m, v71, v68, int32(1609712))
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
									F_ResourceOwnerRemember(m, v71, v68, int32(1609712))
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
							F_ResourceOwnerRemember(m, v71, v68, int32(1609712))
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v275 int32
	_ = v275
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int64
	_ = v463
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v539 int32
	_ = v539
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1155 int32
	_ = v1155
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
	return v1155
L2:
	;
	v24 = F_pstrdup(m, int32(738731))
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
	v1155 = v7
	goto L1
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1144
	v1147 = F___memset(m, v55, int32(0), l2)
	mBase = m.M
	goto L348
L8:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1048 != 0 {
		goto L325
	} else {
		goto L326
	}
L9:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L5
	} else {
		goto L320
	}
L10:
	;
	v31 = F_strlen(m, l1)
	mBase = m.M
	if v31 == l2 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L5
	} else {
		goto L315
	}
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v33 {
	case 0:
		goto L17
	case 1:
		goto L19
	default:
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L5
	} else {
		goto L310
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L5
	} else {
		goto L305
	}
L17:
	;
	v55 = F_pstrdup(m, l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L29
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L22
	}
L19:
	;
	if l2 != int32(1) {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v36 != int32(1) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v39 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v39
	v1155 = v39
	goto L1
L22:
	;
	F_errmsg_internal(m, int32(350767), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(494660), int32(200), int32(400427))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L5
	} else {
		goto L299
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L5
	} else {
		goto L293
	}
L27:
	;
	v81 = v55 + int32(1)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v82 != int32(44) {
		goto L25
	} else {
		goto L35
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L30
	}
L29:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	switch v57 - int32(110) {
	case 0, 11:
		goto L27
	default:
		goto L26
	case 2:
		goto L28
	}
L30:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	F_errdetail(m, int32(636402), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(494660), int32(221), int32(400427))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+2)))
	if v85 != int32(44) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L5
	} else {
		goto L289
	}
L37:
	;
	if v85 == int32(97) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v117 = v55 + int32(3)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v118 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v101 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55)+2)))
	F_sanitize_char_1(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(4385740)
	F_errdetail(m, int32(623699), v15+int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(494660), int32(256), int32(400427))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v122 = v55 + int32(4)
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v123 != 0 {
		goto L57
	} else {
		goto L58
	}
L48:
	;
	goto L49
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L5
	} else {
		goto L283
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L5
	} else {
		goto L278
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L5
	} else {
		goto L273
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L5
	} else {
		goto L268
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L5
	} else {
		goto L263
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L5
	} else {
		goto L258
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L5
	} else {
		goto L253
	}
L56:
	;
	if v130 != 0 {
		goto L126
	} else {
		goto L127
	}
L57:
	;
	v130 = int32(0)
	v133 = v122
	goto L60
L58:
	;
	goto L59
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L5
	} else {
		goto L121
	}
L60:
	;
	v137 = int32(1)
	v138 = F___strchrnul(m, v133, v137)
	mBase = m.M
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v140 == v137 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L59
L62:
	;
	if v144 == int32(0) {
		goto L50
	} else {
		goto L66
	}
L63:
	;
	v144 = v138
	goto L65
L64:
	;
	v144 = int32(0)
	goto L65
L65:
	;
	goto L62
L66:
	;
	v147 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v147)
	if v133 == v144 {
		goto L56
	} else {
		goto L67
	}
L67:
	;
	v150 = int32(61)
	v151 = F___strchrnul(m, v133, v150)
	mBase = m.M
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v153 == v150 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v157 == int32(0) {
		goto L51
	} else {
		goto L72
	}
L69:
	;
	v157 = v151
	goto L71
L70:
	;
	v157 = int32(0)
	goto L71
L71:
	;
	goto L68
L72:
	;
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v160)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v162 == v160 {
		goto L52
	} else {
		goto L73
	}
L73:
	;
	v165 = int32(504684)
	v169 = m.G0
	v171 = v169 - int32(32)
	v172 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v171)+24)) = v172
	*(*int64)(unsafe.Add(mBase, uint32(v171)+16)) = v172
	*(*int64)(unsafe.Add(mBase, uint32(v171)+8)) = v172
	*(*int64)(unsafe.Add(mBase, uint32(v171))) = v172
	v180 = int32(*(*uint8)(unsafe.Add(mBase, _consts[546])))
	if v180 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248+v133))))
	if v250 != 0 {
		goto L53
	} else {
		goto L95
	}
L75:
	;
	v248 = int32(0)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, _consts[547])))
	if v184 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v188 = v133
	goto L81
L79:
	;
	goto L80
L80:
	;
	v198 = v165
	v199 = v180
	goto L84
L81:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v194 == v180 {
		v188 = v188 + int32(1)
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v248 = v188 - v133
	goto L74
L83:
	;
	goto L82
L84:
	;
	v206 = v171 + int32(base.Ui32(v199)>>(uint(int32(3))%32))&int32(28)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v208 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v207 | v208<<(uint(v199)%32)
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
	if v212 != 0 {
		v198 = v198 + v208
		v199 = v212
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v215 == int32(0) {
		v240 = v133
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	v248 = v240 - v133
	goto L74
L88:
	;
	v219 = v133
	v220 = v215
	goto L89
L89:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v171+int32(base.Ui32(v220)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v228)>>(uint(v220)%32))&int32(1) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v240 = v236
	goto L87
L91:
	;
	v240 = v219
	goto L87
L92:
	;
	goto L93
L93:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	v236 = v219 + int32(1)
	if v234 != 0 {
		v219 = v236
		v220 = v234
		goto L89
	} else {
		goto L94
	}
L94:
	;
	goto L90
L95:
	;
	v252 = v157 + int32(1)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	if v253 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v255 = v253
	v261 = v252
	goto L99
L97:
	;
	goto L98
L98:
	;
	v300 = int32(318497)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, _consts[548])))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v304 == int32(0) {
		v323 = v303
		v324 = v304
		goto L109
	} else {
		goto L110
	}
L99:
	;
	if base.Ui32((v255-int32(127))&int32(255)) <= base.Ui32(int32(161)) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L98
L101:
	;
	v275 = v255&int32(255) - int32(9)
	if base.Ui32(int32(23)) < base.Ui32(v275) {
		goto L54
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v286 = v261 + int32(1)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v287 != 0 {
		v255 = v287
		v261 = v286
		goto L99
	} else {
		goto L106
	}
L104:
	;
	if int32(1)<<(uint(v275)%32)&int32(8388627) == int32(0) {
		goto L54
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	goto L100
L107:
	;
	v328 = v144 + int32(1)
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	if v329 != 0 {
		v130 = v326
		v133 = v328
		goto L60
	} else {
		goto L120
	}
L108:
	;
	if v324-v323 != 0 {
		goto L116
	} else {
		goto L117
	}
L109:
	;
	goto L108
L110:
	;
	if v303 != v304 {
		v323 = v303
		v324 = v304
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v308 = v133
	v309 = v300
	goto L112
L112:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+1)))
	if v313 == int32(0) {
		v323 = v312
		v324 = v313
		goto L109
	} else {
		goto L114
	}
L113:
	;
	v323 = v312
	v324 = v313
	goto L109
L114:
	;
	v316 = int32(1)
	if v312 == v313 {
		v308 = v308 + v316
		v309 = v309 + v316
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v326 = v130
	goto L107
L117:
	;
	goto L118
L118:
	;
	if v130 != 0 {
		goto L55
	} else {
		goto L119
	}
L119:
	;
	v326 = v252
	goto L107
L120:
	;
	goto L61
L121:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	F_errdetail(m, int32(588585), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(494660), int32(473), int32(318479))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	if v362 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L5
	} else {
		goto L248
	}
L129:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v365 == int32(0) {
		goto L8
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L5
	} else {
		goto L243
	}
L132:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v373 = v130
	v374 = int32(720067)
	v375 = int32(7)
	goto L134
L133:
	;
	if v420 != 0 {
		goto L149
	} else {
		goto L150
	}
L134:
	;
	if v375 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v420 = int32(0)
	goto L133
L136:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
	if v378 == v379 {
		v401 = v378
		goto L139
	} else {
		goto L140
	}
L137:
	;
	goto L138
L138:
	;
	goto L135
L139:
	;
	v403 = int32(1)
	if v401 != 0 {
		v373 = v373 + v403
		v374 = v374 + v403
		v375 = v375 - v403
		goto L134
	} else {
		goto L148
	}
L140:
	;
	if base.Ui32((v378-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v389 = v378 | int32(32)
	goto L143
L142:
	;
	v389 = v378
	goto L143
L143:
	;
	if base.Ui32((v379-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v398 = v379 | int32(32)
	goto L146
L145:
	;
	v398 = v379
	goto L146
L146:
	;
	if v389 == v398 {
		v401 = v389
		goto L139
	} else {
		goto L147
	}
L147:
	;
	v420 = v389 - v398
	goto L133
L148:
	;
	goto L138
L149:
	;
	v423 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L5
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v432 = v130 + int32(7)
	goto L154
L152:
	;
	if v423 == int32(0) {
		goto L8
	} else {
		goto L153
	}
L153:
	;
	v1011 = int32(586)
	v1017 = int32(617049)
	goto L9
L154:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	if v443 != int32(32) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	if v443 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v432 = v432 + int32(1)
	goto L154
L159:
	;
	v450 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L5
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v456 = int32(552388)
	v460 = m.G0
	v462 = v460 - int32(32)
	v463 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v462)+24)) = v463
	*(*int64)(unsafe.Add(mBase, uint32(v462)+16)) = v463
	*(*int64)(unsafe.Add(mBase, uint32(v462)+8)) = v463
	*(*int64)(unsafe.Add(mBase, uint32(v462))) = v463
	v471 = int32(*(*uint8)(unsafe.Add(mBase, _consts[549])))
	if v471 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	if v450 == int32(0) {
		goto L8
	} else {
		goto L163
	}
L163:
	;
	v1011 = int32(603)
	v1017 = int32(552472)
	goto L9
L164:
	;
	v547 = v539
	goto L185
L165:
	;
	v539 = int32(0)
	goto L164
L166:
	;
	goto L167
L167:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, _consts[550])))
	if v475 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v479 = v432
	goto L171
L169:
	;
	goto L170
L170:
	;
	v489 = v456
	v490 = v471
	goto L174
L171:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if v485 == v471 {
		v479 = v479 + int32(1)
		goto L171
	} else {
		goto L173
	}
L172:
	;
	v539 = v479 - v432
	goto L164
L173:
	;
	goto L172
L174:
	;
	v497 = v462 + int32(base.Ui32(v490)>>(uint(int32(3))%32))&int32(28)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)))
	v499 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v497))) = v498 | v499<<(uint(v490)%32)
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489)+1)))
	if v503 != 0 {
		v489 = v489 + v499
		v490 = v503
		goto L174
	} else {
		goto L176
	}
L175:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	if v506 == int32(0) {
		v531 = v432
		goto L177
	} else {
		goto L178
	}
L176:
	;
	goto L175
L177:
	;
	v539 = v531 - v432
	goto L164
L178:
	;
	v510 = v432
	v511 = v506
	goto L179
L179:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v462+int32(base.Ui32(v511)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v519)>>(uint(v511)%32))&int32(1) == int32(0) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v531 = v527
	goto L177
L181:
	;
	v531 = v510
	goto L177
L182:
	;
	goto L183
L183:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+1)))
	v527 = v510 + int32(1)
	if v525 != 0 {
		v510 = v527
		v511 = v525
		goto L179
	} else {
		goto L184
	}
L184:
	;
	goto L180
L185:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432+v547))))
	if v553 != int32(61) {
		goto L187
	} else {
		goto L188
	}
L186:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L5
	} else {
		goto L239
	}
L187:
	;
	if v553 != 0 {
		goto L191
	} else {
		goto L192
	}
L188:
	;
	v547 = v547 + int32(1)
	goto L185
L189:
	;
	goto L186
L190:
	;
	goto L189
L191:
	;
	v558 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L5
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _consts[551]))
	if v565 == int32(0) {
		goto L190
	} else {
		goto L196
	}
L194:
	;
	if v558 == int32(0) {
		goto L8
	} else {
		goto L195
	}
L195:
	;
	v1011 = int32(625)
	v1017 = int32(564370)
	goto L9
L196:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v565)+12))
	if v568 == int32(0) {
		goto L190
	} else {
		goto L197
	}
L197:
	;
	v572 = F_palloc0(m, int32(8))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L5
	} else {
		goto L198
	}
L198:
	;
	v575 = *(*int32)(unsafe.Add(mBase, _consts[552]))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v368)+364))
	v578 = *(*int32)(unsafe.Add(mBase, _consts[551]))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+12))
	v580 = m.T0[v579].(func(*base.Module, int32, int32, int32, int32) int32)(m, v575, v432, v576, v572)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L5
	} else {
		goto L199
	}
L199:
	;
	if v580 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v586 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L5
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	if v602 != 0 {
		goto L208
	} else {
		goto L209
	}
L203:
	;
	if v586 == int32(0) {
		goto L8
	} else {
		goto L204
	}
L204:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L5
	} else {
		goto L205
	}
L205:
	;
	F_errmsg(m, int32(380569), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(494660), int32(665), int32(354518))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L5
	} else {
		goto L207
	}
L207:
	;
	goto L8
L208:
	;
	F_set_authn_id(m, v368, v602)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L5
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572))))
	if v605 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	goto L210
L212:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	if v669 != 0 {
		goto L233
	} else {
		goto L234
	}
L213:
	;
	v608 = int32(0)
	v611 = F_errstart(m, int32(15), v608)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L5
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v368)+380))
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632)+416)))
	if v633 != 0 {
		v667 = int32(1)
		goto L212
	} else {
		goto L221
	}
L216:
	;
	if v611 == int32(0) {
		v667 = v608
		goto L212
	} else {
		goto L217
	}
L217:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v368)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v615
	F_errmsg(m, int32(682744), v15+int32(48))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L5
	} else {
		goto L218
	}
L218:
	;
	F_errdetail_log(m, int32(601626), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L5
	} else {
		goto L219
	}
L219:
	;
	F_errfinish(m, int32(494660), int32(681), int32(354518))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L5
	} else {
		goto L220
	}
L220:
	;
	v667 = v608
	goto L212
L221:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	if v634 != 0 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v632)+300))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v368)+364))
	v662 = *(*int32)(unsafe.Add(mBase, _consts[118]))
	v663 = F_check_usermap(m, v659, v660, v662)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L5
	} else {
		goto L232
	}
L223:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	if v635 != 0 {
		goto L222
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v636 = int32(0)
	v639 = F_errstart(m, int32(15), v636)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L5
	} else {
		goto L227
	}
L226:
	;
	goto L225
L227:
	;
	if v639 == int32(0) {
		v667 = v636
		goto L212
	} else {
		goto L228
	}
L228:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v368)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v643
	F_errmsg(m, int32(682744), v15+int32(32))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L5
	} else {
		goto L229
	}
L229:
	;
	F_errdetail_log(m, int32(553087), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L5
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(494660), int32(705), int32(354518))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L5
	} else {
		goto L231
	}
L231:
	;
	v667 = v636
	goto L212
L232:
	;
	v667 = base.B2i32(v663 == int32(0))
	goto L212
L233:
	;
	F_pfree(m, v669)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L5
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	F_pfree(m, v572)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L5
	} else {
		goto L237
	}
L236:
	;
	goto L235
L237:
	;
	if v667 == int32(0) {
		goto L8
	} else {
		goto L238
	}
L238:
	;
	v1139 = int32(1)
	v1144 = int32(2)
	goto L7
L239:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L5
	} else {
		goto L240
	}
L240:
	;
	F_errmsg(m, int32(460179), int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L5
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(494660), int32(656), int32(354518))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L5
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L5
	} else {
		goto L244
	}
L244:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L5
	} else {
		goto L245
	}
L245:
	;
	F_errdetail(m, int32(588524), int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(494660), int32(280), int32(400427))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L5
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L5
	} else {
		goto L249
	}
L249:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L5
	} else {
		goto L250
	}
L250:
	;
	F_errdetail(m, int32(610471), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L5
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(494660), int32(273), int32(400427))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L5
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L5
	} else {
		goto L254
	}
L254:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L5
	} else {
		goto L255
	}
L255:
	;
	F_errdetail(m, int32(577698), int32(0))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L5
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(494660), int32(453), int32(318479))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L5
	} else {
		goto L257
	}
L257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L258:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L5
	} else {
		goto L259
	}
L259:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L5
	} else {
		goto L260
	}
L260:
	;
	F_errdetail(m, int32(610715), int32(0))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L5
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(494660), int32(376), int32(212401))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L5
	} else {
		goto L262
	}
L262:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L263:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L5
	} else {
		goto L264
	}
L264:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L5
	} else {
		goto L265
	}
L265:
	;
	F_errdetail(m, int32(617351), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L5
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(494660), int32(350), int32(212401))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L5
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L5
	} else {
		goto L269
	}
L269:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L5
	} else {
		goto L270
	}
L270:
	;
	F_errdetail(m, int32(617315), int32(0))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L5
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(494660), int32(343), int32(212401))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L5
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L5
	} else {
		goto L274
	}
L274:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L5
	} else {
		goto L275
	}
L275:
	;
	F_errdetail(m, int32(610750), int32(0))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L5
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(494660), int32(439), int32(318479))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L5
	} else {
		goto L277
	}
L277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L278:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L5
	} else {
		goto L279
	}
L279:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L5
	} else {
		goto L280
	}
L280:
	;
	F_errdetail(m, int32(589194), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L5
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(494660), int32(421), int32(318479))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L5
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L5
	} else {
		goto L284
	}
L284:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L5
	} else {
		goto L285
	}
L285:
	;
	v869 = int32(*(*int8)(unsafe.Add(mBase, uint32(v117))))
	F_sanitize_char_1(m, v869)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L5
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = int32(4385740)
	F_errdetail(m, int32(647203), v15-int32(-64))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L5
	} else {
		goto L287
	}
L287:
	;
	F_errfinish(m, int32(494660), int32(265), int32(400427))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L5
	} else {
		goto L288
	}
L288:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L289:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L5
	} else {
		goto L290
	}
L290:
	;
	F_errmsg(m, int32(439203), int32(0))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L5
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(494660), int32(250), int32(400427))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L5
	} else {
		goto L292
	}
L292:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L293:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L5
	} else {
		goto L294
	}
L294:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L5
	} else {
		goto L295
	}
L295:
	;
	F_sanitize_char_1(m, base.I32_extend8_s(v57))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(4385740)
	F_errdetail(m, int32(648267), v15)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L5
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(494660), int32(241), int32(400427))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L5
	} else {
		goto L298
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L5
	} else {
		goto L300
	}
L300:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L5
	} else {
		goto L301
	}
L301:
	;
	v935 = int32(*(*int8)(unsafe.Add(mBase, uint32(v81))))
	F_sanitize_char_1(m, v935)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L5
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = int32(4385740)
	F_errdetail(m, int32(647259), v15+int32(80))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L5
	} else {
		goto L303
	}
L303:
	;
	F_errfinish(m, int32(494660), int32(232), int32(400427))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L5
	} else {
		goto L304
	}
L304:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L305:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L5
	} else {
		goto L306
	}
L306:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L5
	} else {
		goto L307
	}
L307:
	;
	F_errdetail(m, int32(613832), int32(0))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L5
	} else {
		goto L308
	}
L308:
	;
	F_errfinish(m, int32(494660), int32(193), int32(400427))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
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
	v976 = m.ExcPending
	if v976 != 0 {
		goto L5
	} else {
		goto L311
	}
L311:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L5
	} else {
		goto L312
	}
L312:
	;
	F_errdetail(m, int32(605797), int32(0))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L5
	} else {
		goto L313
	}
L313:
	;
	F_errfinish(m, int32(494660), int32(175), int32(400427))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L5
	} else {
		goto L314
	}
L314:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L315:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L5
	} else {
		goto L316
	}
L316:
	;
	F_errmsg(m, int32(402684), int32(0))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L5
	} else {
		goto L317
	}
L317:
	;
	F_errdetail(m, int32(552495), int32(0))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L5
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(494660), int32(170), int32(400427))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L5
	} else {
		goto L319
	}
L319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L320:
	;
	F_errmsg(m, int32(280580), int32(0))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L5
	} else {
		goto L321
	}
L321:
	;
	F_errdetail_log(m, v1017, int32(0))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L5
	} else {
		goto L322
	}
L322:
	;
	F_errfinish(m, int32(494660), v1011, int32(111357))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L5
	} else {
		goto L323
	}
L323:
	;
	goto L8
L324:
	;
	F_initStringInfo(m, v15+int32(96))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L5
	} else {
		goto L334
	}
L325:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1049 != 0 {
		goto L324
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L5
	} else {
		goto L329
	}
L328:
	;
	goto L327
L329:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L5
	} else {
		goto L330
	}
L330:
	;
	F_errmsg(m, int32(216658), int32(0))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L5
	} else {
		goto L331
	}
L331:
	;
	F_errdetail_log(m, int32(609317), int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L5
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(494660), int32(500), int32(359027))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L5
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_appendStringInfoString(m, v15+int32(96), v1076)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L5
	} else {
		goto L335
	}
L335:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1081 = F_strstr(m, v1079, int32(552325))
	mBase = m.M
	if v1081 == int32(0) {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	F_appendStringInfoString(m, v15+int32(96), int32(257709))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L5
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	F_initStringInfo(m, v15+int32(112))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L5
	} else {
		goto L340
	}
L339:
	;
	goto L338
L340:
	;
	F_appendStringInfoString(m, v15+int32(112), int32(727462))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L5
	} else {
		goto L341
	}
L341:
	;
	F_appendStringInfoString(m, v15+int32(112), int32(727197))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L5
	} else {
		goto L342
	}
L342:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	F_escape_json(m, v15+int32(112), v1105)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L5
	} else {
		goto L343
	}
L343:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	F_pfree(m, v1108)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L5
	} else {
		goto L344
	}
L344:
	;
	F_appendStringInfoString(m, v15+int32(112), int32(727233))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L5
	} else {
		goto L345
	}
L345:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_escape_json(m, v15+int32(112), v1118)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L5
	} else {
		goto L346
	}
L346:
	;
	F_appendStringInfoString(m, v15+int32(112), int32(6954))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L5
	} else {
		goto L347
	}
L347:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1126
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v15)+116))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1128
	v1139 = int32(0)
	v1144 = int32(1)
	goto L7
L348:
	;
	v1155 = v1139
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
	F_appendStringInfo(m, l0, int32(59295), v6)
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
	v6 = F_uint32in_subr(m, v2, int32(0), int32(432623), v5)
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
	v50 = F_pg_sprintf(m, v43, int32(59295), v9)
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
	F_errmsg(m, int32(207391), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(496560), int32(131), int32(207369))
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
					F_errmsg(m, int32(501270), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(496560), int32(245), int32(35732))
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
						F_errmsg(m, int32(501270), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496560), int32(245), int32(35732))
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
							F_errmsg(m, int32(501270), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(496560), int32(245), int32(35732))
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
								F_errmsg(m, int32(501270), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(496560), int32(245), int32(35732))
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
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(330479)
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
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(330479)
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
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(331104)
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
