package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ObjectsInPublicationToOids(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L26
	}
L2:
	;
	m.G0 = v11 + int32(16)
	return
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v24 = v5
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v24<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	switch v31 {
	case 0:
		goto L12
	case 1:
		goto L8
	case 2:
		goto L11
	case 3:
		goto L10
	default:
		goto L9
	}
L6:
	;
	goto L2
L7:
	;
	v86 = v24 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v86 < v87 {
		v24 = v86
		goto L5
	} else {
		goto L25
	}
L8:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v76 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+16)) = uint8(v76)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v80 = F_lappend(m, v78, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L24
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L21
	}
L10:
	;
	v49 = F_fetch_search_path(m, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L17
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v42 = F_get_namespace_oid(m, v40, int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L15
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+16)) = uint8(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v37 = F_lappend(m, v35, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37
	goto L7
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v45 = F_list_append_unique_oid(m, v44, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v45
	goto L7
L17:
	;
	if v49 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	F_list_free(m, v49)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v58 = F_list_append_unique_oid(m, v57, v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v58
	goto L7
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v65
	F_errmsg_internal(m, int32(_a_F_ObjectsInPublicationToOids_0), v11)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_ObjectsInPublicationToOids_1), int32(230), int32(_a_F_ObjectsInPublicationToOids_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v80
	goto L7
L25:
	;
	goto L6
L26:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(_a_F_ObjectsInPublicationToOids_3), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_ObjectsInPublicationToOids_1), int32(220), int32(_a_F_ObjectsInPublicationToOids_2))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
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
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
				if v28 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = l1 + v28
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
			v69 = F_query_tree_walker_impl(m, l0, int32(1126), v9+int32(8), int32(0))
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v58 int32
	_ = v58
	v3 = int32(0)
	F_recomputeNamespacePath(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_OpclassnameGetOpcid[0]))
	if v13 == int32(0) {
		v58 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v58
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(0) < v16 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = v3
	goto L8
L6:
	;
	goto L7
L7:
	;
	v58 = int32(0)
	goto L3
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v24<<(uint(int32(2))%32))))
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_OpclassnameGetOpcid[1]))
	if v32 != v34 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v39 = F_GetSysCacheOid(m, int32(13), base.I64_extend_i32_u(l0), base.I64_extend_i32_u(l1), base.I64_extend_i32_u(v32), int64(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v43 = v24 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v43 < v44 {
		v24 = v43
		goto L8
	} else {
		goto L15
	}
L13:
	;
	if v39 != 0 {
		v58 = v39
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
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
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
				v88 = v52
				return v88
			} else {
				v64 = v52
				v67 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
				F_ResourceOwnerRemember(m, v67, base.I64_extend_i32_s(v64), int32(_a_F_OpenTemporaryFile_0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
					v76 = v73 + v64*int32(48)
					v78 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v78
					v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)))
					v82 = v80 | int32(2)
					*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)) = uint16(v82)
					v85 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[3])) = uint8(v85)
					v88 = v64
					return v88
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
						v88 = v52
						return v88
					} else {
						v64 = v52
						v67 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
						F_ResourceOwnerRemember(m, v67, base.I64_extend_i32_s(v64), int32(_a_F_OpenTemporaryFile_0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
							v76 = v73 + v64*int32(48)
							v78 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v78
							v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)))
							v82 = v80 | int32(2)
							*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)) = uint16(v82)
							v85 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[3])) = uint8(v85)
							v88 = v64
							return v88
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
							v88 = v52
							return v88
						} else {
							v64 = v52
							v67 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
							F_ResourceOwnerRemember(m, v67, base.I64_extend_i32_s(v64), int32(_a_F_OpenTemporaryFile_0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
								v76 = v73 + v64*int32(48)
								v78 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v78
								v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)))
								v82 = v80 | int32(2)
								*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)) = uint16(v82)
								v85 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[3])) = uint8(v85)
								v88 = v64
								return v88
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
									v88 = v52
									return v88
								} else {
									v64 = v52
									v67 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
									F_ResourceOwnerRemember(m, v67, base.I64_extend_i32_s(v64), int32(_a_F_OpenTemporaryFile_0))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v73 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
										v76 = v73 + v64*int32(48)
										v78 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
										*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v78
										v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)))
										v82 = v80 | int32(2)
										*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)) = uint16(v82)
										v85 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[3])) = uint8(v85)
										v88 = v64
										return v88
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
							F_ResourceOwnerRemember(m, v67, base.I64_extend_i32_s(v64), int32(_a_F_OpenTemporaryFile_0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[1]))
								v76 = v73 + v64*int32(48)
								v78 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[2]))
								*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v78
								v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)))
								v82 = v80 | int32(2)
								*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)) = uint16(v82)
								v85 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_OpenTemporaryFile[3])) = uint8(v85)
								v88 = v64
								return v88
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
	var v32 int64
	_ = v32
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
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
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
	return v156
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
	v48 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)))
	v51 = F_SearchSysCacheList(m, int32(39), int32(3), v48, base.I64_extend_i32_u(l1), base.I64_extend_i32_u(l2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	if v27 == int32(0) {
		v156 = v4
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v32 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)))
	v36 = F_SearchSysCache4(m, int32(39), v32, base.I64_extend_i32_u(l1), base.I64_extend_i32_u(l2), base.I64_extend_i32_u(v27))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v36 == int32(0) {
		v156 = v4
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40+v41)))
	F_ReleaseCatCache(m, v36)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v156 = v43
	goto L3
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+56))
	if v53 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_ReleaseCatCacheList(m, v51)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
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
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v156 = v4
	goto L3
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_OpernameGetOprid[0]))
	if v61 == int32(0) {
		v142 = v4
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_ReleaseCatCacheList(m, v51)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L36
	}
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 <= int32(0) {
		v142 = v4
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v67 = int32(0)
	if v67 < v64 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v70 = v64
	goto L23
L22:
	;
	v70 = v67
	goto L23
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_OpernameGetOprid[1]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v82 = v4
	goto L24
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v75+v82<<(uint(int32(2))%32))))
	if v91 == v74 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v142 = v4
	goto L18
L26:
	;
	v135 = v82 + int32(1)
	if v135 != v70 {
		v82 = v135
		goto L24
	} else {
		goto L35
	}
L27:
	;
	v93 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v51)+56))
	if v94 <= v93 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v97 = v93
	goto L29
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v51-int32(-64)+v97<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+72))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+22)))
	v115 = v113 + v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+68))
	if v91 != v116 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v142 = v121
	goto L18
L31:
	;
	v119 = v97 + int32(1)
	if v94 != v119 {
		v97 = v119
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
	v156 = v142
	goto L3
}
func F_oauth_exchange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int64
	_ = v203
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v312 int32
	_ = v312
	var v323 int32
	_ = v323
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int64
	_ = v592
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v630 int32
	_ = v630
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v691 int32
	_ = v691
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int64
	_ = v712
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v788 int32
	_ = v788
	var v795 int32
	_ = v795
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1110 int32
	_ = v1110
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1279 int32
	_ = v1279
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1362 int32
	_ = v1362
	var v1370 int32
	_ = v1370
	var v1382 int32
	_ = v1382
	var v1392 int32
	_ = v1392
	v7 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(160)
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(-1)
	if l1 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v21 + int32(160)
	return v1392
L2:
	;
	v30 = F_pstrdup(m, int32(_a_F_oauth_exchange_0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v30
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v1392 = v7
	goto L1
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v1382
	F___memset(m, v85, int32(0), l2)
	mBase = m.M
	goto L393
L8:
	;
	v1288 = m.G0
	v1290 = v1288 - int32(32)
	m.G0 = v1290
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1292 == int32(0) {
		goto L371
	} else {
		goto L372
	}
L9:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L5
	} else {
		goto L366
	}
L10:
	;
	v37 = F_strlen(m, l1)
	mBase = m.M
	if v37 == l2 {
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
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L5
	} else {
		goto L361
	}
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(2)) <= base.Ui32(v39-int32(1)) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L5
	} else {
		goto L356
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L5
	} else {
		goto L351
	}
L17:
	;
	v85 = F_pstrdup(m, l1)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L40
	}
L18:
	;
	if v39 == int32(0) {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if l2 != int32(1) {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errmsg_internal(m, int32(_a_F_oauth_exchange_1), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(223), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
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
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v61 != int32(1) {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v64 = int32(2)
	if v39 == v64 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v69 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3)
	v1392 = v64
	goto L1
L30:
	;
	if v69 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_4), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v80 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v80
	v1392 = v80
	goto L1
L34:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(212), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L5
	} else {
		goto L345
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L5
	} else {
		goto L339
	}
L38:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v110 != int32(44) {
		goto L36
	} else {
		goto L46
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L41
	}
L40:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	switch v87 - int32(110) {
	case 0, 11:
		goto L38
	default:
		goto L37
	case 2:
		goto L39
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
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v103 = F_errdetail(m, int32(_a_F_oauth_exchange_6), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(244), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+2)))
	if v113 != int32(44) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L5
	} else {
		goto L335
	}
L48:
	;
	if v113 == int32(97) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+3)))
	if v144 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	v129 = int32(*(*int8)(unsafe.Add(mBase, uint32(v85)+2)))
	F_sanitize_char_1(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = int32(_a_F_oauth_exchange_7)
	v137 = F_errdetail(m, int32(_a_F_oauth_exchange_8), v21+int32(16))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(279), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
	if v147 != 0 {
		goto L68
	} else {
		goto L69
	}
L59:
	;
	goto L60
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L5
	} else {
		goto L329
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L5
	} else {
		goto L324
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L5
	} else {
		goto L319
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L5
	} else {
		goto L314
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L5
	} else {
		goto L309
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L5
	} else {
		goto L304
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L5
	} else {
		goto L299
	}
L67:
	;
	if v158 != 0 {
		goto L133
	} else {
		goto L134
	}
L68:
	;
	v158 = v7
	v159 = v85 + int32(4)
	goto L71
L69:
	;
	goto L70
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L5
	} else {
		goto L128
	}
L71:
	;
	v168 = int32(1)
	v169 = F___strchrnul(m, v159, v168)
	mBase = m.M
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v171 == v168 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L70
L73:
	;
	if v175 == int32(0) {
		goto L61
	} else {
		goto L77
	}
L74:
	;
	v175 = v169
	goto L76
L75:
	;
	v175 = int32(0)
	goto L76
L76:
	;
	goto L73
L77:
	;
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v178)
	if v159 == v175 {
		goto L67
	} else {
		goto L78
	}
L78:
	;
	v181 = int32(61)
	v182 = F___strchrnul(m, v159, v181)
	mBase = m.M
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v184 == v181 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v188 == int32(0) {
		goto L62
	} else {
		goto L83
	}
L80:
	;
	v188 = v182
	goto L82
L81:
	;
	v188 = int32(0)
	goto L82
L82:
	;
	goto L79
L83:
	;
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v191)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v193 == v191 {
		goto L63
	} else {
		goto L84
	}
L84:
	;
	v196 = int32(_a_F_oauth_exchange_9)
	v200 = m.G0
	v202 = v200 - int32(32)
	v203 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v202)+24)) = v203
	*(*int64)(unsafe.Add(mBase, uint32(v202)+16)) = v203
	*(*int64)(unsafe.Add(mBase, uint32(v202)+8)) = v203
	*(*int64)(unsafe.Add(mBase, uint32(v202))) = v203
	v211 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oauth_exchange[0])))
	if v211 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279+v159))))
	if v281 != 0 {
		goto L64
	} else {
		goto L104
	}
L86:
	;
	v279 = int32(0)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oauth_exchange[1])))
	if v215 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v219 = v159
	goto L92
L90:
	;
	goto L91
L91:
	;
	v229 = v196
	v230 = v211
	goto L95
L92:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if v225 == v211 {
		v219 = v219 + int32(1)
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v279 = v219 - v159
	goto L85
L94:
	;
	goto L93
L95:
	;
	v237 = v202 + int32(base.Ui32(v230)>>(uint(int32(3))%32))&int32(28)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v239 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v238 | v239<<(uint(v230)%32)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+1)))
	if v243 != 0 {
		v229 = v229 + v239
		v230 = v243
		goto L95
	} else {
		goto L97
	}
L96:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v246 == int32(0) {
		v269 = v159
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L96
L98:
	;
	v279 = v269 - v159
	goto L85
L99:
	;
	v250 = v159
	v251 = v246
	goto L100
L100:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v202+int32(base.Ui32(v251)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v259)>>(uint(v251)%32))&int32(1) == int32(0) {
		v269 = v250
		goto L98
	} else {
		goto L102
	}
L101:
	;
	v269 = v267
	goto L98
L102:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+1)))
	v267 = v250 + int32(1)
	if v265 != 0 {
		v250 = v267
		v251 = v265
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v283 = v188 + int32(1)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	if v284 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v286 = v284
	v291 = v283
	goto L108
L106:
	;
	goto L107
L107:
	;
	v344 = int32(_a_F_oauth_exchange_10)
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oauth_exchange[2])))
	if base.B2i32(v347 == int32(0))|base.B2i32(v347 != v350) != 0 {
		v368 = v347
		v369 = v350
		goto L117
	} else {
		goto L118
	}
L108:
	;
	if base.Ui32((v286-int32(127))&int32(255)) <= base.Ui32(int32(161)) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L107
L110:
	;
	v312 = v286&int32(255) - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v312))|base.B2i32(int32(1)<<(uint(v312)%32)&int32(_a_F_oauth_exchange_11) == int32(0)) != 0 {
		goto L65
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+1)))
	if v323 != 0 {
		v286 = v323
		v291 = v291 + int32(1)
		goto L108
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	goto L109
L115:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	if v374 != 0 {
		v158 = v371
		v159 = v175 + int32(1)
		goto L71
	} else {
		goto L127
	}
L116:
	;
	if v368-v369 != 0 {
		goto L123
	} else {
		goto L124
	}
L117:
	;
	goto L116
L118:
	;
	v353 = v159
	v354 = v344
	goto L119
L119:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+1)))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+1)))
	if v358 == int32(0) {
		v368 = v358
		v369 = v357
		goto L117
	} else {
		goto L121
	}
L120:
	;
	v368 = v358
	v369 = v357
	goto L117
L121:
	;
	v361 = int32(1)
	if v358 == v357 {
		v353 = v353 + v361
		v354 = v354 + v361
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v371 = v158
	goto L115
L124:
	;
	goto L125
L125:
	;
	if v158 != 0 {
		goto L66
	} else {
		goto L126
	}
L126:
	;
	v371 = v283
	goto L115
L127:
	;
	goto L72
L128:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L5
	} else {
		goto L130
	}
L130:
	;
	v406 = F_errdetail(m, int32(_a_F_oauth_exchange_12), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L5
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(518), int32(_a_F_oauth_exchange_13))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+1)))
	if v413 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L5
	} else {
		goto L294
	}
L136:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)+380))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+388))
	if v418 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L137:
	;
	goto L138
L138:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L5
	} else {
		goto L289
	}
L139:
	;
	v614 = v158
	v615 = int32(_a_F_oauth_exchange_14)
	v616 = int32(7)
	goto L177
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v448
	v555 = F_psprintf(m, int32(_a_F_oauth_exchange_15), v21+int32(112))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L5
	} else {
		goto L163
	}
L141:
	;
	v529 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_oauth_exchange[3])) = uint8(v529)
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v531 != 0 {
		goto L139
	} else {
		goto L162
	}
L142:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	if v421 <= int32(0) {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_exchange[4]))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v418)+12))
	v440 = v7
	goto L144
L144:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v426+v440<<(uint(int32(2))%32))))
	if v425 == int32(0) {
		goto L140
	} else {
		goto L146
	}
L145:
	;
	goto L141
L146:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	if v451 <= int32(0) {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v425)+12))
	v457 = int32(0)
	goto L148
L148:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v454+v457<<(uint(int32(2))%32))))
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448))))
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	if base.B2i32(v480 == int32(0))|base.B2i32(v480 != v483) != 0 {
		v501 = v480
		v502 = v483
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v508 = v440 + int32(1)
	if v421 != v508 {
		v440 = v508
		goto L144
	} else {
		goto L161
	}
L150:
	;
	if v501-v502 != 0 {
		goto L157
	} else {
		goto L158
	}
L151:
	;
	goto L150
L152:
	;
	v486 = v448
	v487 = v477
	goto L153
L153:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+1)))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486)+1)))
	if v491 == int32(0) {
		v501 = v491
		v502 = v490
		goto L151
	} else {
		goto L155
	}
L154:
	;
	v501 = v491
	v502 = v490
	goto L151
L155:
	;
	v494 = int32(1)
	if v491 == v490 {
		v486 = v486 + v494
		v487 = v487 + v494
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	v505 = v457 + int32(1)
	if v505 != v451 {
		v457 = v505
		goto L148
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	goto L149
L160:
	;
	goto L140
L161:
	;
	goto L145
L162:
	;
	v1279 = int32(2)
	goto L8
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v555
	v561 = F_psprintf(m, int32(_a_F_oauth_exchange_16), v21+int32(96))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L5
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v561
	v566 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L5
	} else {
		goto L165
	}
L165:
	;
	if v566 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L5
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3)
	v1392 = int32(2)
	goto L1
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v555
	F_errmsg(m, int32(_a_F_oauth_exchange_16), v21+int32(80))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L5
	} else {
		goto L170
	}
L170:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v417)+380))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v577
	v583 = F_errdetail(m, int32(_a_F_oauth_exchange_17), v21-int32(-64))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L5
	} else {
		goto L171
	}
L171:
	;
	F_errhint(m, int32(_a_F_oauth_exchange_18), int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L5
	} else {
		goto L172
	}
L172:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L5
	} else {
		goto L173
	}
L173:
	;
	v592 = *(*int64)(unsafe.Add(mBase, uint32(v417)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = base.I64_rotl(v592, int64(32))
	F_errcontext_msg(m, int32(_a_F_oauth_exchange_19), v21+int32(48))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L5
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(1078), int32(_a_F_oauth_exchange_20))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L5
	} else {
		goto L175
	}
L175:
	;
	goto L168
L176:
	;
	if v661 != 0 {
		goto L192
	} else {
		goto L193
	}
L177:
	;
	if v616 != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v661 = int32(0)
	goto L176
L179:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615))))
	if v619 == v620 {
		v642 = v619
		goto L182
	} else {
		goto L183
	}
L180:
	;
	goto L181
L181:
	;
	goto L178
L182:
	;
	v644 = int32(1)
	if v642 != 0 {
		v614 = v614 + v644
		v615 = v615 + v644
		v616 = v616 - v644
		goto L177
	} else {
		goto L191
	}
L183:
	;
	if base.Ui32((v619-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v630 = v619 | int32(32)
	goto L186
L185:
	;
	v630 = v619
	goto L186
L186:
	;
	if base.Ui32((v620-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v639 = v620 | int32(32)
	goto L189
L188:
	;
	v639 = v620
	goto L189
L189:
	;
	if v630 == v639 {
		v642 = v630
		goto L182
	} else {
		goto L190
	}
L190:
	;
	v661 = v630 - v639
	goto L176
L191:
	;
	goto L181
L192:
	;
	v664 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L5
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v674 = v158 + int32(7)
	goto L199
L195:
	;
	if v664 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1279 = int32(1)
	goto L8
L197:
	;
	goto L198
L198:
	;
	v1238 = int32(620)
	v1243 = int32(_a_F_oauth_exchange_21)
	goto L9
L199:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v691 != int32(32) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	if v691 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L202:
	;
	v674 = v674 + int32(1)
	goto L199
L204:
	;
	v698 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L5
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v705 = int32(_a_F_oauth_exchange_22)
	v709 = m.G0
	v711 = v709 - int32(32)
	v712 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v711)+24)) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v711)+16)) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v711)+8)) = v712
	*(*int64)(unsafe.Add(mBase, uint32(v711))) = v712
	v720 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oauth_exchange[5])))
	if v720 == int32(0) {
		goto L212
	} else {
		goto L213
	}
L207:
	;
	if v698 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v1279 = int32(1)
	goto L8
L209:
	;
	goto L210
L210:
	;
	v1238 = int32(637)
	v1243 = int32(_a_F_oauth_exchange_23)
	goto L9
L211:
	;
	v795 = v788
	goto L230
L212:
	;
	v788 = int32(0)
	goto L211
L213:
	;
	goto L214
L214:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oauth_exchange[6])))
	if v724 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v728 = v674
	goto L218
L216:
	;
	goto L217
L217:
	;
	v738 = v705
	v739 = v720
	goto L221
L218:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v728))))
	if v734 == v720 {
		v728 = v728 + int32(1)
		goto L218
	} else {
		goto L220
	}
L219:
	;
	v788 = v728 - v674
	goto L211
L220:
	;
	goto L219
L221:
	;
	v746 = v711 + int32(base.Ui32(v739)>>(uint(int32(3))%32))&int32(28)
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)))
	v748 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v746))) = v747 | v748<<(uint(v739)%32)
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v738)+1)))
	if v752 != 0 {
		v738 = v738 + v748
		v739 = v752
		goto L221
	} else {
		goto L223
	}
L222:
	;
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v755 == int32(0) {
		v778 = v674
		goto L224
	} else {
		goto L225
	}
L223:
	;
	goto L222
L224:
	;
	v788 = v778 - v674
	goto L211
L225:
	;
	v759 = v674
	v760 = v755
	goto L226
L226:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v711+int32(base.Ui32(v760)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v768)>>(uint(v760)%32))&int32(1) == int32(0) {
		v778 = v759
		goto L224
	} else {
		goto L228
	}
L227:
	;
	v778 = v776
	goto L224
L228:
	;
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759)+1)))
	v776 = v759 + int32(1)
	if v774 != 0 {
		v759 = v776
		v760 = v774
		goto L226
	} else {
		goto L229
	}
L229:
	;
	goto L227
L230:
	;
	v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674+v795))))
	if v808 != int32(61) {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L5
	} else {
		goto L285
	}
L232:
	;
	if v808 != 0 {
		goto L236
	} else {
		goto L237
	}
L233:
	;
	v795 = v795 + int32(1)
	goto L230
L234:
	;
	goto L231
L235:
	;
	goto L234
L236:
	;
	v813 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L5
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v821 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_exchange[7]))
	if v821 == int32(0) {
		goto L235
	} else {
		goto L243
	}
L239:
	;
	if v813 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1279 = int32(1)
	goto L8
L241:
	;
	goto L242
L242:
	;
	v1238 = int32(659)
	v1243 = int32(_a_F_oauth_exchange_24)
	goto L9
L243:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v821)+12))
	if v824 == int32(0) {
		goto L235
	} else {
		goto L244
	}
L244:
	;
	v828 = F_palloc0(m, int32(12))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L5
	} else {
		goto L245
	}
L245:
	;
	v831 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_exchange[8]))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v416)+364))
	v834 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_exchange[7]))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v834)+12))
	v836 = m.T0[v835].(func(*base.Module, int32, int32, int32, int32) int32)(m, v831, v674, v832, v828)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L5
	} else {
		goto L246
	}
L246:
	;
	if v836 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v842 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L5
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	if v867 != 0 {
		goto L261
	} else {
		goto L262
	}
L250:
	;
	if v842 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L5
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v828)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v864
	v1279 = int32(1)
	goto L8
L254:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_25), int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L5
	} else {
		goto L255
	}
L255:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v828)+8))
	if v851 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v851
	F_errdetail_log(m, int32(_a_F_oauth_exchange_26), v21+int32(32))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L5
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(700), int32(_a_F_oauth_exchange_27))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L5
	} else {
		goto L260
	}
L259:
	;
	goto L258
L260:
	;
	goto L253
L261:
	;
	F_set_authn_id(m, v416, v867)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L5
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v828))))
	if v870 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L264:
	;
	goto L263
L265:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	if v897 != 0 {
		goto L279
	} else {
		goto L280
	}
L266:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v828)+8))
	if v873 != 0 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	goto L268
L268:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v416)+380))
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v879)+384)))
	if v880 != 0 {
		v896 = int32(1)
		goto L265
	} else {
		goto L272
	}
L269:
	;
	v875 = v873
	goto L271
L270:
	;
	v875 = int32(_a_F_oauth_exchange_28)
	goto L271
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v875
	v896 = int32(0)
	goto L265
L272:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	if v881 != 0 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v879)+300))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v416)+364))
	v889 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_exchange[9]))
	v890 = F_check_usermap(m, v886, v887, v889)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L5
	} else {
		goto L278
	}
L274:
	;
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881))))
	if v882 != 0 {
		goto L273
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(_a_F_oauth_exchange_29)
	v896 = int32(0)
	goto L265
L277:
	;
	goto L276
L278:
	;
	v896 = base.B2i32(v890 == int32(0))
	goto L265
L279:
	;
	F_pfree(m, v897)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L5
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	F_pfree(m, v828)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L5
	} else {
		goto L283
	}
L282:
	;
	goto L281
L283:
	;
	v902 = int32(1)
	if v896 != 0 {
		v1370 = v902
		v1382 = int32(3)
		goto L7
	} else {
		goto L284
	}
L284:
	;
	v1279 = v902
	goto L8
L285:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L5
	} else {
		goto L286
	}
L286:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_30), int32(0))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L5
	} else {
		goto L287
	}
L287:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(690), int32(_a_F_oauth_exchange_27))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
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
	F_errcode(m, int32(16908800))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L5
	} else {
		goto L290
	}
L290:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L5
	} else {
		goto L291
	}
L291:
	;
	v938 = F_errdetail(m, int32(_a_F_oauth_exchange_31), int32(0))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L5
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(303), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L5
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L294:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L5
	} else {
		goto L295
	}
L295:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	v958 = F_errdetail(m, int32(_a_F_oauth_exchange_32), int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L5
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(296), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
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
	v971 = m.ExcPending
	if v971 != 0 {
		goto L5
	} else {
		goto L300
	}
L300:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L5
	} else {
		goto L301
	}
L301:
	;
	v978 = F_errdetail(m, int32(_a_F_oauth_exchange_33), int32(0))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L5
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(498), int32(_a_F_oauth_exchange_13))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
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
	v991 = m.ExcPending
	if v991 != 0 {
		goto L5
	} else {
		goto L305
	}
L305:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L5
	} else {
		goto L306
	}
L306:
	;
	v998 = F_errdetail(m, int32(_a_F_oauth_exchange_34), int32(0))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L5
	} else {
		goto L307
	}
L307:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(421), int32(_a_F_oauth_exchange_35))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
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
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L5
	} else {
		goto L310
	}
L310:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L5
	} else {
		goto L311
	}
L311:
	;
	v1018 = F_errdetail(m, int32(_a_F_oauth_exchange_36), int32(0))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L5
	} else {
		goto L312
	}
L312:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(395), int32(_a_F_oauth_exchange_35))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
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
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L5
	} else {
		goto L315
	}
L315:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L5
	} else {
		goto L316
	}
L316:
	;
	v1038 = F_errdetail(m, int32(_a_F_oauth_exchange_37), int32(0))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L5
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(388), int32(_a_F_oauth_exchange_35))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L5
	} else {
		goto L318
	}
L318:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L319:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L5
	} else {
		goto L320
	}
L320:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L5
	} else {
		goto L321
	}
L321:
	;
	v1058 = F_errdetail(m, int32(_a_F_oauth_exchange_38), int32(0))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L5
	} else {
		goto L322
	}
L322:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(484), int32(_a_F_oauth_exchange_13))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L5
	} else {
		goto L323
	}
L323:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L324:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L5
	} else {
		goto L325
	}
L325:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L5
	} else {
		goto L326
	}
L326:
	;
	v1078 = F_errdetail(m, int32(_a_F_oauth_exchange_39), int32(0))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L5
	} else {
		goto L327
	}
L327:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(466), int32(_a_F_oauth_exchange_13))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L5
	} else {
		goto L328
	}
L328:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L329:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L5
	} else {
		goto L330
	}
L330:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L5
	} else {
		goto L331
	}
L331:
	;
	v1096 = int32(*(*int8)(unsafe.Add(mBase, uint32(v85)+3)))
	F_sanitize_char_1(m, v1096)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L5
	} else {
		goto L332
	}
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = int32(_a_F_oauth_exchange_7)
	v1104 = F_errdetail(m, int32(_a_F_oauth_exchange_40), v21+int32(128))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L5
	} else {
		goto L333
	}
L333:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(288), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L5
	} else {
		goto L334
	}
L334:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L335:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L5
	} else {
		goto L336
	}
L336:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_41), int32(0))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L5
	} else {
		goto L337
	}
L337:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(273), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L5
	} else {
		goto L338
	}
L338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L5
	} else {
		goto L340
	}
L340:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L5
	} else {
		goto L341
	}
L341:
	;
	F_sanitize_char_1(m, base.I32_extend8_s(v87))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L5
	} else {
		goto L342
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(_a_F_oauth_exchange_7)
	v1144 = F_errdetail(m, int32(_a_F_oauth_exchange_42), v21)
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L5
	} else {
		goto L343
	}
L343:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(264), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L5
	} else {
		goto L344
	}
L344:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L345:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L5
	} else {
		goto L346
	}
L346:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L5
	} else {
		goto L347
	}
L347:
	;
	v1162 = int32(*(*int8)(unsafe.Add(mBase, uint32(v85)+1)))
	F_sanitize_char_1(m, v1162)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L5
	} else {
		goto L348
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+144)) = int32(_a_F_oauth_exchange_7)
	v1170 = F_errdetail(m, int32(_a_F_oauth_exchange_43), v21+int32(144))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L5
	} else {
		goto L349
	}
L349:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(255), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
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
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L5
	} else {
		goto L352
	}
L352:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L5
	} else {
		goto L353
	}
L353:
	;
	v1190 = F_errdetail(m, int32(_a_F_oauth_exchange_44), int32(0))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L5
	} else {
		goto L354
	}
L354:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(204), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L5
	} else {
		goto L355
	}
L355:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L356:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L5
	} else {
		goto L357
	}
L357:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L5
	} else {
		goto L358
	}
L358:
	;
	v1210 = F_errdetail(m, int32(_a_F_oauth_exchange_45), int32(0))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L5
	} else {
		goto L359
	}
L359:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(185), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L5
	} else {
		goto L360
	}
L360:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L361:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L5
	} else {
		goto L362
	}
L362:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_5), int32(0))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L5
	} else {
		goto L363
	}
L363:
	;
	v1230 = F_errdetail(m, int32(_a_F_oauth_exchange_46), int32(0))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L5
	} else {
		goto L364
	}
L364:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(180), int32(_a_F_oauth_exchange_3))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L5
	} else {
		goto L365
	}
L365:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L366:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_47), int32(0))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L5
	} else {
		goto L367
	}
L367:
	;
	F_errdetail_log(m, v1243, int32(0))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L5
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), v1238, int32(_a_F_oauth_exchange_48))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L5
	} else {
		goto L369
	}
L369:
	;
	v1279 = int32(1)
	goto L8
L370:
	;
	v1370 = int32(0)
	v1382 = v1279
	goto L7
L371:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1346 = m.ExcPending
	if v1346 != 0 {
		goto L5
	} else {
		goto L388
	}
L372:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1295 == int32(0) {
		goto L371
	} else {
		goto L373
	}
L373:
	;
	F_initStringInfo(m, v1290)
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L5
	} else {
		goto L374
	}
L374:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_appendStringInfoString(m, v1290, v1300)
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L5
	} else {
		goto L375
	}
L375:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1305 = F_strstr(m, v1303, int32(_a_F_oauth_exchange_49))
	mBase = m.M
	if v1305 == int32(0) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	F_appendStringInfoString(m, v1290, int32(_a_F_oauth_exchange_50))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L5
	} else {
		goto L379
	}
L377:
	;
	goto L378
L378:
	;
	v1312 = v1290 + int32(16)
	F_initStringInfo(m, v1312)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L5
	} else {
		goto L380
	}
L379:
	;
	goto L378
L380:
	;
	F_appendStringInfoString(m, v1312, int32(_a_F_oauth_exchange_51))
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L5
	} else {
		goto L381
	}
L381:
	;
	F_appendStringInfoString(m, v1312, int32(_a_F_oauth_exchange_52))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L5
	} else {
		goto L382
	}
L382:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1290)))
	F_escape_json(m, v1312, v1321)
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L5
	} else {
		goto L383
	}
L383:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1290)))
	F_pfree(m, v1324)
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L5
	} else {
		goto L384
	}
L384:
	;
	F_appendStringInfoString(m, v1312, int32(_a_F_oauth_exchange_53))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L5
	} else {
		goto L385
	}
L385:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_escape_json(m, v1312, v1330)
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L5
	} else {
		goto L386
	}
L386:
	;
	F_appendStringInfoString(m, v1312, int32(_a_F_oauth_exchange_54))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L5
	} else {
		goto L387
	}
L387:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1290)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1336
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1290)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1338
	m.G0 = v1290 + int32(32)
	goto L370
L388:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L5
	} else {
		goto L389
	}
L389:
	;
	F_errmsg(m, int32(_a_F_oauth_exchange_55), int32(0))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L5
	} else {
		goto L390
	}
L390:
	;
	F_errdetail_log(m, int32(_a_F_oauth_exchange_56), int32(0))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L5
	} else {
		goto L391
	}
L391:
	;
	F_errfinish(m, int32(_a_F_oauth_exchange_2), int32(545), int32(_a_F_oauth_exchange_57))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L5
	} else {
		goto L392
	}
L392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L393:
	;
	v1392 = v1370
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
func F_oid8gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(base.Ui64(v3) < base.Ui64(v2)))
}
func F_oid8larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui64(v5) < base.Ui64(v4) {
		v7 = v4
	} else {
		v7 = v5
	}
	return v7
}
func F_oidin(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = F_uint32in_subr(m, v2, int32(0), int32(_a_F_oidin_0), v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v6)
	}
}
func F_oidlarger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v4) < base.Ui32(v3) {
		v6 = v3
	} else {
		v6 = v4
	}
	return base.I64_extend_i32_u(v6)
}
func F_oidvectoreq(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_btoidvectorcmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(base.I32_wrap_i64(v2) == int32(0)))
	}
}
func F_oidvectorout(m *base.Module, l0 int32) int64 {
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
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v87 = m.ExcPending
	if v87 != 0 {
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
	return int64(0)
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
	return base.I64_extend_i32_u(v25)
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
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(_a_F_oidvectorout_1), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_oidvectorout_2), int32(131), int32(_a_F_oidvectorout_3))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
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
func F_oidvectorrecv(m *base.Module, l0 int32) int64 {
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v8 = Fn14313(m, l0, int32(_a_F_oidvectorrecv_0), int32(245), int32(_a_F_oidvectorrecv_1), int32(_a_F_oidvectorrecv_2), int32(26), int64(26))
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		return v8
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
	if v58 != int32(320) {
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
