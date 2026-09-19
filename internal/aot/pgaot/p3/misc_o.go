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
	F_errmsg_internal(m, int32(_a_F_ObjectsInPublicationToOids_0), v10)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_ObjectsInPublicationToOids_1), int32(226), int32(_a_F_ObjectsInPublicationToOids_2))
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
	F_errmsg(m, int32(_a_F_ObjectsInPublicationToOids_3), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_ObjectsInPublicationToOids_1), int32(216), int32(_a_F_ObjectsInPublicationToOids_2))
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
	var v45 int32
	_ = v45
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
					v45 = int32(0)
					for {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v45<<(uint(int32(2))%32))))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v52 + l1
						v56 = v45 + int32(1)
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
						if v56 < v57 {
							v45 = v56
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
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_OpclassnameGetOpcid[0]))
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
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_OpclassnameGetOpcid[1]))
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	if l0 != 0 {
		v48 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[0]))
		if v48 != 0 {
			v50 = v48
		} else {
			v50 = int32(1663)
		}
		v52 = F_OpenTemporaryFileInTablespace(m, v50, int32(1))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
			v58 = v55 + v52*int32(48)
			v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
			v61 = v59 | int32(5)
			*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)) = uint16(v61)
			if l0 != 0 {
				v87 = v52
				return v87
			} else {
				v64 = v52
				v67 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
				F_ResourceOwnerRemember(m, v67, v64, int32(_a_F_OpenTemporaryFile_0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
					v75 = v72 + v64*int32(48)
					v77 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v77
					v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)))
					v81 = v79 | int32(2)
					*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)) = uint16(v81)
					v84 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[3])) = uint8(v84)
					v87 = v64
					return v87
				}
			}
		}
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
		F_ResourceOwnerEnlarge(m, v5)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[4]))
			if v11 <= int32(0) {
				v48 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[0]))
				if v48 != 0 {
					v50 = v48
				} else {
					v50 = int32(1663)
				}
				v52 = F_OpenTemporaryFileInTablespace(m, v50, int32(1))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
					v58 = v55 + v52*int32(48)
					v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
					v61 = v59 | int32(5)
					*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)) = uint16(v61)
					if l0 != 0 {
						v87 = v52
						return v87
					} else {
						v64 = v52
						v67 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
						F_ResourceOwnerRemember(m, v67, v64, int32(_a_F_OpenTemporaryFile_0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v72 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
							v75 = v72 + v64*int32(48)
							v77 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v77
							v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)))
							v81 = v79 | int32(2)
							*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)) = uint16(v81)
							v84 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[3])) = uint8(v84)
							v87 = v64
							return v87
						}
					}
				}
			} else {
				v14 = int32(_a_F_OpenTemporaryFile_1)
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[5]))
				v18 = v16 + int32(1)
				if v18 < v11 {
					v21 = v18
				} else {
					v21 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[5])) = v21
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[6]))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v21<<(uint(int32(2))%32))))
				if v28 == int32(0) {
					v48 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[0]))
					if v48 != 0 {
						v50 = v48
					} else {
						v50 = int32(1663)
					}
					v52 = F_OpenTemporaryFileInTablespace(m, v50, int32(1))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
						v58 = v55 + v52*int32(48)
						v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
						v61 = v59 | int32(5)
						*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)) = uint16(v61)
						if l0 != 0 {
							v87 = v52
							return v87
						} else {
							v64 = v52
							v67 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
							F_ResourceOwnerRemember(m, v67, v64, int32(_a_F_OpenTemporaryFile_0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
								v75 = v72 + v64*int32(48)
								v77 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v77
								v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)))
								v81 = v79 | int32(2)
								*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)) = uint16(v81)
								v84 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[3])) = uint8(v84)
								v87 = v64
								return v87
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
							v48 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[0]))
							if v48 != 0 {
								v50 = v48
							} else {
								v50 = int32(1663)
							}
							v52 = F_OpenTemporaryFileInTablespace(m, v50, int32(1))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
								v58 = v55 + v52*int32(48)
								v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
								v61 = v59 | int32(5)
								*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)) = uint16(v61)
								if l0 != 0 {
									v87 = v52
									return v87
								} else {
									v64 = v52
									v67 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
									F_ResourceOwnerRemember(m, v67, v64, int32(_a_F_OpenTemporaryFile_0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v72 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
										v75 = v72 + v64*int32(48)
										v77 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
										*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v77
										v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)))
										v81 = v79 | int32(2)
										*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)) = uint16(v81)
										v84 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[3])) = uint8(v84)
										v87 = v64
										return v87
									}
								}
							}
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
							v40 = v37 + v32*int32(48)
							v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)))
							v43 = v41 | int32(5)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)) = uint16(v43)
							v64 = v32
							v67 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
							F_ResourceOwnerRemember(m, v67, v64, int32(_a_F_OpenTemporaryFile_0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
								v75 = v72 + v64*int32(48)
								v77 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v77
								v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)))
								v81 = v79 | int32(2)
								*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)) = uint16(v81)
								v84 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[3])) = uint8(v84)
								v87 = v64
								return v87
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
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_OpernameGetOprid[0]))
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
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_OpernameGetOprid[1]))
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
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v271 int32
	_ = v271
	var v282 int32
	_ = v282
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int64
	_ = v461
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v537 int32
	_ = v537
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1121 int32
	_ = v1121
	var v1126 int32
	_ = v1126
	var v1137 int32
	_ = v1137
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
	return v1137
L2:
	;
	v24 = F_pstrdup(m, int32(_a_F_oauth_exchange_0))
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
	v1137 = v7
	goto L1
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1126
	F___memset(m, v55, int32(0), l2)
	mBase = m.M
	goto L342
L8:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1046 != 0 {
		goto L319
	} else {
		goto L320
	}
L9:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L5
	} else {
		goto L314
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
	v991 = m.ExcPending
	if v991 != 0 {
		goto L5
	} else {
		goto L309
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
	v971 = m.ExcPending
	if v971 != 0 {
		goto L5
	} else {
		goto L304
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L5
	} else {
		goto L299
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
	v1137 = v39
	goto L1
L22:
	;
	F_errmsg_internal(m, int32(_a_F_oauth_exchange_1), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(200), int32(_a_F_oauth_exchange_3))
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
	v925 = m.ExcPending
	if v925 != 0 {
		goto L5
	} else {
		goto L293
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L5
	} else {
		goto L287
	}
L27:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v80 != int32(44) {
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
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(221), int32(_a_F_oauth_exchange_3))
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
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+2)))
	if v83 != int32(44) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L5
	} else {
		goto L283
	}
L37:
	;
	if v83 == int32(97) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+3)))
	if v114 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v99 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55)+2)))
	F_sanitize_char_1(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(_a_F_oauth_exchange_6)
	F_errdetail(m, int32(_a_F_oauth_exchange_7), v15+int32(16))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(256), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
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
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+4)))
	if v117 != 0 {
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
	v859 = m.ExcPending
	if v859 != 0 {
		goto L5
	} else {
		goto L277
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L5
	} else {
		goto L272
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L5
	} else {
		goto L267
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L5
	} else {
		goto L262
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L5
	} else {
		goto L257
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L5
	} else {
		goto L252
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L5
	} else {
		goto L247
	}
L56:
	;
	if v126 != 0 {
		goto L122
	} else {
		goto L123
	}
L57:
	;
	v126 = int32(0)
	v131 = v55 + int32(4)
	goto L60
L58:
	;
	goto L59
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L5
	} else {
		goto L117
	}
L60:
	;
	v133 = int32(1)
	v134 = F___strchrnul(m, v131, v133)
	mBase = m.M
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v136 == v133 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L59
L62:
	;
	if v140 == int32(0) {
		goto L50
	} else {
		goto L66
	}
L63:
	;
	v140 = v134
	goto L65
L64:
	;
	v140 = int32(0)
	goto L65
L65:
	;
	goto L62
L66:
	;
	v143 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v143)
	if v131 == v140 {
		goto L56
	} else {
		goto L67
	}
L67:
	;
	v146 = int32(61)
	v147 = F___strchrnul(m, v131, v146)
	mBase = m.M
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v149 == v146 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v153 == int32(0) {
		goto L51
	} else {
		goto L72
	}
L69:
	;
	v153 = v147
	goto L71
L70:
	;
	v153 = int32(0)
	goto L71
L71:
	;
	goto L68
L72:
	;
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v156)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v158 == v156 {
		goto L52
	} else {
		goto L73
	}
L73:
	;
	v161 = int32(_a_F_oauth_exchange_8)
	v165 = m.G0
	v167 = v165 - int32(32)
	v168 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v167)+24)) = v168
	*(*int64)(unsafe.Add(mBase, uint32(v167)+16)) = v168
	*(*int64)(unsafe.Add(mBase, uint32(v167)+8)) = v168
	*(*int64)(unsafe.Add(mBase, uint32(v167))) = v168
	v176 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oauth_exchange[0])))
	if v176 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244+v131))))
	if v246 != 0 {
		goto L53
	} else {
		goto L93
	}
L75:
	;
	v244 = int32(0)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oauth_exchange[1])))
	if v180 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v184 = v131
	goto L81
L79:
	;
	goto L80
L80:
	;
	v194 = v161
	v195 = v176
	goto L84
L81:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v190 == v176 {
		v184 = v184 + int32(1)
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v244 = v184 - v131
	goto L74
L83:
	;
	goto L82
L84:
	;
	v202 = v167 + int32(base.Ui32(v195)>>(uint(int32(3))%32))&int32(28)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v204 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v203 | v204<<(uint(v195)%32)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	if v208 != 0 {
		v194 = v194 + v204
		v195 = v208
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v211 == int32(0) {
		v234 = v131
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	v244 = v234 - v131
	goto L74
L88:
	;
	v215 = v131
	v216 = v211
	goto L89
L89:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v167+int32(base.Ui32(v216)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v224)>>(uint(v216)%32))&int32(1) == int32(0) {
		v234 = v215
		goto L87
	} else {
		goto L91
	}
L90:
	;
	v234 = v232
	goto L87
L91:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	v232 = v215 + int32(1)
	if v230 != 0 {
		v215 = v232
		v216 = v230
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v248 = v153 + int32(1)
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if v249 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v251 = v249
	v257 = v248
	goto L97
L95:
	;
	goto L96
L96:
	;
	v297 = int32(_a_F_oauth_exchange_9)
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oauth_exchange[2])))
	if base.B2i32(v300 == int32(0))|base.B2i32(v300 != v303) != 0 {
		v321 = v300
		v322 = v303
		goto L106
	} else {
		goto L107
	}
L97:
	;
	if base.Ui32((v251-int32(127))&int32(255)) <= base.Ui32(int32(161)) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L96
L99:
	;
	v271 = v251&int32(255) - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v271))|base.B2i32(int32(1)<<(uint(v271)%32)&int32(_a_F_oauth_exchange_10) == int32(0)) != 0 {
		goto L54
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+1)))
	if v282 != 0 {
		v251 = v282
		v257 = v257 + int32(1)
		goto L97
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	goto L98
L104:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)))
	if v327 != 0 {
		v126 = v324
		v131 = v140 + int32(1)
		goto L60
	} else {
		goto L116
	}
L105:
	;
	if v321-v322 != 0 {
		goto L112
	} else {
		goto L113
	}
L106:
	;
	goto L105
L107:
	;
	v306 = v131
	v307 = v297
	goto L108
L108:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+1)))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+1)))
	if v311 == int32(0) {
		v321 = v311
		v322 = v310
		goto L106
	} else {
		goto L110
	}
L109:
	;
	v321 = v311
	v322 = v310
	goto L106
L110:
	;
	v314 = int32(1)
	if v311 == v310 {
		v306 = v306 + v314
		v307 = v307 + v314
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v324 = v126
	goto L104
L113:
	;
	goto L114
L114:
	;
	if v126 != 0 {
		goto L55
	} else {
		goto L115
	}
L115:
	;
	v324 = v248
	goto L104
L116:
	;
	goto L61
L117:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L5
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_11), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(473), int32(_a_F_oauth_exchange_12))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L5
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)))
	if v360 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L5
	} else {
		goto L242
	}
L125:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v363 == int32(0) {
		goto L8
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L5
	} else {
		goto L237
	}
L128:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v371 = v126
	v372 = int32(_a_F_oauth_exchange_13)
	v373 = int32(7)
	goto L130
L129:
	;
	if v418 != 0 {
		goto L145
	} else {
		goto L146
	}
L130:
	;
	if v373 != 0 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v418 = int32(0)
	goto L129
L132:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371))))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	if v376 == v377 {
		v399 = v376
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L134
L134:
	;
	goto L131
L135:
	;
	v401 = int32(1)
	if v399 != 0 {
		v371 = v371 + v401
		v372 = v372 + v401
		v373 = v373 - v401
		goto L130
	} else {
		goto L144
	}
L136:
	;
	if base.Ui32((v376-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v387 = v376 | int32(32)
	goto L139
L138:
	;
	v387 = v376
	goto L139
L139:
	;
	if base.Ui32((v377-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v396 = v377 | int32(32)
	goto L142
L141:
	;
	v396 = v377
	goto L142
L142:
	;
	if v387 == v396 {
		v399 = v387
		goto L135
	} else {
		goto L143
	}
L143:
	;
	v418 = v387 - v396
	goto L129
L144:
	;
	goto L134
L145:
	;
	v421 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L5
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v430 = v126 + int32(7)
	goto L150
L148:
	;
	if v421 == int32(0) {
		goto L8
	} else {
		goto L149
	}
L149:
	;
	v1009 = int32(586)
	v1015 = int32(_a_F_oauth_exchange_14)
	goto L9
L150:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	if v441 != int32(32) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if v441 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	v430 = v430 + int32(1)
	goto L150
L155:
	;
	v448 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L5
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v454 = int32(_a_F_oauth_exchange_15)
	v458 = m.G0
	v460 = v458 - int32(32)
	v461 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v460)+24)) = v461
	*(*int64)(unsafe.Add(mBase, uint32(v460)+16)) = v461
	*(*int64)(unsafe.Add(mBase, uint32(v460)+8)) = v461
	*(*int64)(unsafe.Add(mBase, uint32(v460))) = v461
	v469 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oauth_exchange[3])))
	if v469 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	if v448 == int32(0) {
		goto L8
	} else {
		goto L159
	}
L159:
	;
	v1009 = int32(603)
	v1015 = int32(_a_F_oauth_exchange_16)
	goto L9
L160:
	;
	v545 = v537
	goto L179
L161:
	;
	v537 = int32(0)
	goto L160
L162:
	;
	goto L163
L163:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oauth_exchange[4])))
	if v473 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v477 = v430
	goto L167
L165:
	;
	goto L166
L166:
	;
	v487 = v454
	v488 = v469
	goto L170
L167:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	if v483 == v469 {
		v477 = v477 + int32(1)
		goto L167
	} else {
		goto L169
	}
L168:
	;
	v537 = v477 - v430
	goto L160
L169:
	;
	goto L168
L170:
	;
	v495 = v460 + int32(base.Ui32(v488)>>(uint(int32(3))%32))&int32(28)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	v497 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v495))) = v496 | v497<<(uint(v488)%32)
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+1)))
	if v501 != 0 {
		v487 = v487 + v497
		v488 = v501
		goto L170
	} else {
		goto L172
	}
L171:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	if v504 == int32(0) {
		v527 = v430
		goto L173
	} else {
		goto L174
	}
L172:
	;
	goto L171
L173:
	;
	v537 = v527 - v430
	goto L160
L174:
	;
	v508 = v430
	v509 = v504
	goto L175
L175:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v460+int32(base.Ui32(v509)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v517)>>(uint(v509)%32))&int32(1) == int32(0) {
		v527 = v508
		goto L173
	} else {
		goto L177
	}
L176:
	;
	v527 = v525
	goto L173
L177:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+1)))
	v525 = v508 + int32(1)
	if v523 != 0 {
		v508 = v525
		v509 = v523
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430+v545))))
	if v551 != int32(61) {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L5
	} else {
		goto L233
	}
L181:
	;
	if v551 != 0 {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	v545 = v545 + int32(1)
	goto L179
L183:
	;
	goto L180
L184:
	;
	goto L183
L185:
	;
	v556 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L5
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_exchange[5]))
	if v563 == int32(0) {
		goto L184
	} else {
		goto L190
	}
L188:
	;
	if v556 == int32(0) {
		goto L8
	} else {
		goto L189
	}
L189:
	;
	v1009 = int32(625)
	v1015 = int32(_a_F_oauth_exchange_17)
	goto L9
L190:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v563)+12))
	if v566 == int32(0) {
		goto L184
	} else {
		goto L191
	}
L191:
	;
	v570 = F_palloc0(m, int32(8))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_exchange[6]))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v366)+364))
	v576 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_exchange[5]))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v576)+12))
	v578 = m.T0[v577].(func(*base.Module, int32, int32, int32, int32) int32)(m, v573, v430, v574, v570)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L5
	} else {
		goto L193
	}
L193:
	;
	if v578 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v584 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L5
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	if v600 != 0 {
		goto L202
	} else {
		goto L203
	}
L197:
	;
	if v584 == int32(0) {
		goto L8
	} else {
		goto L198
	}
L198:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L5
	} else {
		goto L199
	}
L199:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_18), int32(0))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L5
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(665), int32(_a_F_oauth_exchange_19))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L5
	} else {
		goto L201
	}
L201:
	;
	goto L8
L202:
	;
	F_set_authn_id(m, v366, v600)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L5
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570))))
	if v603 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	goto L204
L206:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	if v667 != 0 {
		goto L227
	} else {
		goto L228
	}
L207:
	;
	v606 = int32(0)
	v609 = F_errstart(m, int32(15), v606)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L5
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v366)+380))
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630)+416)))
	if v631 != 0 {
		v665 = int32(1)
		goto L206
	} else {
		goto L215
	}
L210:
	;
	if v609 == int32(0) {
		v665 = v606
		goto L206
	} else {
		goto L211
	}
L211:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v366)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v613
	F_errmsg(m, int32(_a_F_oauth_exchange_20), v15+int32(48))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L5
	} else {
		goto L212
	}
L212:
	;
	F_errdetail_log(m, int32(_a_F_oauth_exchange_21), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L5
	} else {
		goto L213
	}
L213:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(681), int32(_a_F_oauth_exchange_19))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L5
	} else {
		goto L214
	}
L214:
	;
	v665 = v606
	goto L206
L215:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	if v632 != 0 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v630)+300))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v366)+364))
	v660 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_exchange[7]))
	v661 = F_check_usermap(m, v657, v658, v660)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L5
	} else {
		goto L226
	}
L217:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632))))
	if v633 != 0 {
		goto L216
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v634 = int32(0)
	v637 = F_errstart(m, int32(15), v634)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L5
	} else {
		goto L221
	}
L220:
	;
	goto L219
L221:
	;
	if v637 == int32(0) {
		v665 = v634
		goto L206
	} else {
		goto L222
	}
L222:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v366)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v641
	F_errmsg(m, int32(_a_F_oauth_exchange_20), v15+int32(32))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L5
	} else {
		goto L223
	}
L223:
	;
	F_errdetail_log(m, int32(_a_F_oauth_exchange_22), int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L5
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(705), int32(_a_F_oauth_exchange_19))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L5
	} else {
		goto L225
	}
L225:
	;
	v665 = v634
	goto L206
L226:
	;
	v665 = base.B2i32(v661 == int32(0))
	goto L206
L227:
	;
	F_pfree(m, v667)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L5
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	F_pfree(m, v570)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L5
	} else {
		goto L231
	}
L230:
	;
	goto L229
L231:
	;
	if v665 == int32(0) {
		goto L8
	} else {
		goto L232
	}
L232:
	;
	v1121 = int32(1)
	v1126 = int32(2)
	goto L7
L233:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L5
	} else {
		goto L234
	}
L234:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_23), int32(0))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L5
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(656), int32(_a_F_oauth_exchange_19))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L5
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L5
	} else {
		goto L238
	}
L238:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L5
	} else {
		goto L239
	}
L239:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_24), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L5
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(280), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L5
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L5
	} else {
		goto L243
	}
L243:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L5
	} else {
		goto L244
	}
L244:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_25), int32(0))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L5
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(273), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L5
	} else {
		goto L248
	}
L248:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L5
	} else {
		goto L249
	}
L249:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_26), int32(0))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L5
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(453), int32(_a_F_oauth_exchange_12))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L5
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L5
	} else {
		goto L253
	}
L253:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L5
	} else {
		goto L254
	}
L254:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_27), int32(0))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L5
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(376), int32(_a_F_oauth_exchange_28))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L5
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L5
	} else {
		goto L258
	}
L258:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L5
	} else {
		goto L259
	}
L259:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_29), int32(0))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L5
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(350), int32(_a_F_oauth_exchange_28))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L5
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L262:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L5
	} else {
		goto L263
	}
L263:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L5
	} else {
		goto L264
	}
L264:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_30), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L5
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(343), int32(_a_F_oauth_exchange_28))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L5
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L5
	} else {
		goto L268
	}
L268:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L5
	} else {
		goto L269
	}
L269:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_31), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L5
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(439), int32(_a_F_oauth_exchange_12))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L5
	} else {
		goto L271
	}
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L272:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L5
	} else {
		goto L273
	}
L273:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L5
	} else {
		goto L274
	}
L274:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_32), int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L5
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(421), int32(_a_F_oauth_exchange_12))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L5
	} else {
		goto L276
	}
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L277:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L5
	} else {
		goto L278
	}
L278:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L5
	} else {
		goto L279
	}
L279:
	;
	v867 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55)+3)))
	F_sanitize_char_1(m, v867)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L5
	} else {
		goto L280
	}
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = int32(_a_F_oauth_exchange_6)
	F_errdetail(m, int32(_a_F_oauth_exchange_33), v15-int32(-64))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L5
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(265), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L5
	} else {
		goto L284
	}
L284:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_34), int32(0))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L5
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(250), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L5
	} else {
		goto L286
	}
L286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L287:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L5
	} else {
		goto L288
	}
L288:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L5
	} else {
		goto L289
	}
L289:
	;
	F_sanitize_char_1(m, base.I32_extend8_s(v57))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L5
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_oauth_exchange_6)
	F_errdetail(m, int32(_a_F_oauth_exchange_35), v15)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L5
	} else {
		goto L291
	}
L291:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(241), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
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
	v928 = m.ExcPending
	if v928 != 0 {
		goto L5
	} else {
		goto L294
	}
L294:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L5
	} else {
		goto L295
	}
L295:
	;
	v933 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55)+1)))
	F_sanitize_char_1(m, v933)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = int32(_a_F_oauth_exchange_6)
	F_errdetail(m, int32(_a_F_oauth_exchange_36), v15+int32(80))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L5
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(232), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
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
	v954 = m.ExcPending
	if v954 != 0 {
		goto L5
	} else {
		goto L300
	}
L300:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L5
	} else {
		goto L301
	}
L301:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_37), int32(0))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L5
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(193), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L5
	} else {
		goto L303
	}
L303:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L304:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L5
	} else {
		goto L305
	}
L305:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L5
	} else {
		goto L306
	}
L306:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_38), int32(0))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L5
	} else {
		goto L307
	}
L307:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(175), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L5
	} else {
		goto L308
	}
L308:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L309:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L5
	} else {
		goto L310
	}
L310:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L5
	} else {
		goto L311
	}
L311:
	;
	F_errdetail(m, int32(_a_F_oauth_exchange_39), int32(0))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L5
	} else {
		goto L312
	}
L312:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(170), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L5
	} else {
		goto L313
	}
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L314:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_40), int32(0))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L5
	} else {
		goto L315
	}
L315:
	;
	F_errdetail_log(m, v1015, int32(0))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L5
	} else {
		goto L316
	}
L316:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), v1009, int32(_a_F_oauth_exchange_41))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L5
	} else {
		goto L317
	}
L317:
	;
	goto L8
L318:
	;
	v1069 = v15 + int32(96)
	F_initStringInfo(m, v1069)
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L5
	} else {
		goto L328
	}
L319:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1047 != 0 {
		goto L318
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L5
	} else {
		goto L323
	}
L322:
	;
	goto L321
L323:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L5
	} else {
		goto L324
	}
L324:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_42), int32(0))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L5
	} else {
		goto L325
	}
L325:
	;
	F_errdetail_log(m, int32(_a_F_oauth_exchange_43), int32(0))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L5
	} else {
		goto L326
	}
L326:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(500), int32(_a_F_oauth_exchange_44))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L5
	} else {
		goto L327
	}
L327:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L328:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_appendStringInfoString(m, v1069, v1072)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L5
	} else {
		goto L329
	}
L329:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1077 = F_strstr(m, v1075, int32(_a_F_oauth_exchange_45))
	mBase = m.M
	if v1077 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	F_appendStringInfoString(m, v1069, int32(_a_F_oauth_exchange_46))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L5
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1084 = v15 + int32(112)
	F_initStringInfo(m, v1084)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L5
	} else {
		goto L334
	}
L333:
	;
	goto L332
L334:
	;
	F_appendStringInfoString(m, v1084, int32(_a_F_oauth_exchange_47))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L5
	} else {
		goto L335
	}
L335:
	;
	F_appendStringInfoString(m, v1084, int32(_a_F_oauth_exchange_48))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L5
	} else {
		goto L336
	}
L336:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	F_escape_json(m, v1084, v1093)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L5
	} else {
		goto L337
	}
L337:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	F_pfree(m, v1096)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L5
	} else {
		goto L338
	}
L338:
	;
	F_appendStringInfoString(m, v1084, int32(_a_F_oauth_exchange_49))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L5
	} else {
		goto L339
	}
L339:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_escape_json(m, v1084, v1102)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L5
	} else {
		goto L340
	}
L340:
	;
	F_appendStringInfoString(m, v1084, int32(_a_F_oauth_exchange_50))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L5
	} else {
		goto L341
	}
L341:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1108
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v15)+116))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1110
	v1121 = int32(0)
	v1126 = int32(1)
	goto L7
L342:
	;
	v1137 = v1121
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
	F_appendStringInfo(m, l0, int32(_a_F_offset_elem_desc_0), v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func F_offsethash_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = int32(16)
	v8 = (int32(base.Ui32(l1)>>(uint(v4)%32)) ^ l1) * int32(-2048144789)
	v13 = (int32(base.Ui32(v8)>>(uint(int32(13))%32)) ^ v8) * int32(-1028477387)
	v17 = F_offsethash_insert_hash_internal(m, l0, l1, int32(base.Ui32(v13)>>(uint(v4)%32))^v13, l2)
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		return v17
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
	v6 = F_uint32in_subr(m, v2, int32(0), int32(_a_F_oidin_0), v5)
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
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v7 = m.G0
	v9 = v7 - int32(32)
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
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L15
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
	v19 = int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v25 = F_palloc(m, v20*int32(12)|v19)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v20 <= int32(0) {
		v71 = v25
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v77)
	m.G0 = v9 + int32(32)
	return v25
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v31
	v36 = F_pg_sprintf(m, v25, int32(_a_F_oidvectorout_0), v9+int32(16))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v38 = int32(1)
	v39 = v25 + v38
	v40 = F_strlen(m, v39)
	mBase = m.M
	v41 = v40 + v39
	if v20 == v38 {
		v71 = v41
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v46 = v41
	v49 = v19
	goto L11
L11:
	;
	v52 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v52)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(24)+v49<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v57
	v62 = F_pg_sprintf(m, v46+int32(1), int32(_a_F_oidvectorout_0), v9)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v71 = v67
	goto L7
L13:
	;
	v65 = v46 + int32(2)
	v66 = F_strlen(m, v65)
	mBase = m.M
	v67 = v66 + v65
	v69 = v49 + int32(1)
	if v69 != v20 {
		v46 = v67
		v49 = v69
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(_a_F_oidvectorout_1), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_oidvectorout_2), int32(131), int32(_a_F_oidvectorout_3))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_oidvectorrecv(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13947(m, l0, int32(_a_F_oidvectorrecv_0), int32(245), int32(_a_F_oidvectorrecv_1), int32(_a_F_oidvectorrecv_2), int32(26))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
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
	var v94 int32
	_ = v94
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
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
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
	return v190
L2:
	;
	v190 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 < int32(2) {
		v190 = l1
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
	v66 = *(*float64)(unsafe.Add(mBase, _c_F_order_qual_clauses[0]))
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
	v94 = int32(1)
	goto L24
L24:
	;
	v108 = v25 + v94*int32(24)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	v111 = *(*float64)(unsafe.Add(mBase, uint32(v108)+8))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v108)))
	v115 = v94
	goto L27
L25:
	;
	v162 = int32(1)
	if v20 <= v162 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v154 = v25 + v150*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+20)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v154)+16)) = v110
	*(*float64)(unsafe.Add(mBase, uint32(v154)+8)) = v111
	*(*int64)(unsafe.Add(mBase, uint32(v154))) = v112
	v160 = v94 + int32(1)
	if v160 != v92 {
		v94 = v160
		goto L24
	} else {
		goto L35
	}
L27:
	;
	v127 = v25 + v115*int32(24)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127-int32(8))))
	if base.Ui32(v130) < base.Ui32(v110) {
		v150 = v115
		goto L26
	} else {
		goto L29
	}
L28:
	;
	v150 = int32(0)
	goto L26
L29:
	;
	if v110 == v130 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v127-int32(16))))
	if base.F64_ge(v111, v135) != 0 {
		v150 = v115
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
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v138)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v127)+16)) = v139
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v127)+8)) = v141
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v143
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
	v165 = v162
	goto L38
L37:
	;
	v165 = v20
	goto L38
L38:
	;
	v166 = int32(0)
	v169 = v166
	v170 = v166
	goto L39
L39:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v25+v170*int32(24))))
	v184 = F_lappend(m, v169, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L6
	} else {
		goto L41
	}
L40:
	;
	v190 = v184
	goto L1
L41:
	;
	v187 = v170 + int32(1)
	if v187 != v165 {
		v169 = v184
		v170 = v187
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
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a_F_owningrel_does_not_exist_skipping_0)
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
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a_F_owningrel_does_not_exist_skipping_0)
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
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(_a_F_owningrel_does_not_exist_skipping_1)
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
