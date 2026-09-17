package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatalogCacheComputeTupleHashValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v29 int32
	_ = v29
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
	var v39 int32
	_ = v39
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
	var v47 int32
	_ = v47
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
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch l1 - int32(1) {
	case 0:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v32 = F_fastgetattr_4(m, l2, v29, v14, v12+int32(15))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v78 = int32(0)
			v79 = v32
			v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v85 = m.T0[v84].(func(*base.Module, int32) int32)(m, v79)
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return int32(0)
			} else {
				m.G0 = v12 + int32(16)
				return v85 ^ v78
			}
		}
	case 1:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v19 = v12 + int32(15)
		v20 = F_fastgetattr_4(m, l2, v17, v14, v19)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v26 = F_fastgetattr_4(m, l2, v25, v14, v19)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v66 = int32(0)
				v67 = v26
				v69 = v20
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v73 = m.T0[v72].(func(*base.Module, int32) int32)(m, v69)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					v78 = base.I32_rotl(v73, int32(8)) ^ v66
					v79 = v67
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v85 = m.T0[v84].(func(*base.Module, int32) int32)(m, v79)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						m.G0 = v12 + int32(16)
						return v85 ^ v78
					}
				}
			}
		}
	case 2:
		v39 = int32(0)
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v42 = v12 + int32(15)
		v43 = F_fastgetattr_4(m, l2, v40, v14, v42)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v46 = F_fastgetattr_4(m, l2, v45, v14, v42)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v49 = F_fastgetattr_4(m, l2, v48, v14, v42)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					if l1 == int32(4) {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v54 = m.T0[v53].(func(*base.Module, int32) int32)(m, v39)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v59 = base.I32_rotl(v54, int32(24))
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, v43)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v66 = v59 ^ base.I32_rotl(v61, int32(16))
								v67 = v49
								v69 = v46
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v73 = m.T0[v72].(func(*base.Module, int32) int32)(m, v69)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v78 = base.I32_rotl(v73, int32(8)) ^ v66
									v79 = v67
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v85 = m.T0[v84].(func(*base.Module, int32) int32)(m, v79)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(16)
										return v85 ^ v78
									}
								}
							}
						}
					} else {
						v59 = int32(0)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, v43)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v66 = v59 ^ base.I32_rotl(v61, int32(16))
							v67 = v49
							v69 = v46
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v73 = m.T0[v72].(func(*base.Module, int32) int32)(m, v69)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v78 = base.I32_rotl(v73, int32(8)) ^ v66
								v79 = v67
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v85 = m.T0[v84].(func(*base.Module, int32) int32)(m, v79)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									m.G0 = v12 + int32(16)
									return v85 ^ v78
								}
							}
						}
					}
				}
			}
		}
	case 3:
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v37 = F_fastgetattr_4(m, l2, v34, v14, v12+int32(15))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			v39 = v37
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v42 = v12 + int32(15)
			v43 = F_fastgetattr_4(m, l2, v40, v14, v42)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v46 = F_fastgetattr_4(m, l2, v45, v14, v42)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v49 = F_fastgetattr_4(m, l2, v48, v14, v42)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if l1 == int32(4) {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v54 = m.T0[v53].(func(*base.Module, int32) int32)(m, v39)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v59 = base.I32_rotl(v54, int32(24))
								v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, v43)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									v66 = v59 ^ base.I32_rotl(v61, int32(16))
									v67 = v49
									v69 = v46
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v73 = m.T0[v72].(func(*base.Module, int32) int32)(m, v69)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v78 = base.I32_rotl(v73, int32(8)) ^ v66
										v79 = v67
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v85 = m.T0[v84].(func(*base.Module, int32) int32)(m, v79)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											m.G0 = v12 + int32(16)
											return v85 ^ v78
										}
									}
								}
							}
						} else {
							v59 = int32(0)
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v61 = m.T0[v60].(func(*base.Module, int32) int32)(m, v43)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v66 = v59 ^ base.I32_rotl(v61, int32(16))
								v67 = v49
								v69 = v46
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v73 = m.T0[v72].(func(*base.Module, int32) int32)(m, v69)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									v78 = base.I32_rotl(v73, int32(8)) ^ v66
									v79 = v67
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v85 = m.T0[v84].(func(*base.Module, int32) int32)(m, v79)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(16)
										return v85 ^ v78
									}
								}
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v95 = m.ExcPending
		if v95 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
			F_errmsg_internal(m, int32(_a_F_CatalogCacheComputeTupleHashValue_0), v12)
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_CatalogCacheComputeTupleHashValue_1), int32(428), int32(_a_F_CatalogCacheComputeTupleHashValue_2))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
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
func F_IsCatalogRelation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	return base.B2i32(base.Ui32(v2) < base.Ui32(int32(_a_F_IsCatalogRelation_0)))
}
func F_get_catalog_object_by_oid_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_get_catalog_object_by_oid_extended[0]))
	if v16 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L27
	} else {
		goto L42
	}
L2:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if int32(0) < v65 {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v17 == v14 {
		v60 = v16
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v26 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_catalog_object_by_oid_extended[0])) = v54
	v60 = v54
	goto L2
L8:
	;
	v54 = v30 + int32(_a_F_get_catalog_object_by_oid_extended_0)
	goto L7
L9:
	;
	v54 = v30 + int32(_a_F_get_catalog_object_by_oid_extended_1)
	goto L7
L10:
	;
	v54 = v30 + int32(_a_F_get_catalog_object_by_oid_extended_2)
	goto L7
L11:
	;
	v30 = v26 * int32(40)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_get_catalog_object_by_oid_extended[1])))
	if v14 != v31 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v54 = v30 + int32(_a_F_get_catalog_object_by_oid_extended_3)
	goto L7
L13:
	;
	if v26 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_get_catalog_object_by_oid_extended[2])))
	if v37 == v14 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_get_catalog_object_by_oid_extended[3])))
	if v39 == v14 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_get_catalog_object_by_oid_extended[4])))
	if v41 == v14 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v26 = v26 + int32(4)
	goto L11
L20:
	;
	m.G0 = v12 - int32(-64)
	return v102
L21:
	;
	if l3 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v77 = v10 + int32(-48)
	F_ScanKeyInit(m, v77, l1, int32(3), int32(184), l2)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L27
	} else {
		goto L30
	}
L24:
	;
	v68 = F_SearchSysCacheLockedCopy1(m, v65, l2)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v73 = F_SearchSysCacheCopy(m, v65, l2, int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L27
	} else {
		goto L29
	}
L27:
	;
	return int32(0)
L28:
	;
	v102 = v68
	goto L20
L29:
	;
	v102 = v73
	goto L20
L30:
	;
	v82 = int32(0)
	v83 = int32(1)
	v86 = F_systable_beginscan(m, l0, v75, v83, v82, v83, v77)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v88 = F_systable_getnext(m, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	if v88 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if l3 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v97 = v82
	goto L35
L35:
	;
	F_systable_endscan(m, v86)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L27
	} else {
		goto L41
	}
L36:
	;
	F_LockTuple(m, l0, v88+int32(4), int32(7))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L27
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v95 = F_heap_copytuple(m, v88)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L27
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v97 = v95
	goto L35
L41:
	;
	v102 = v97
	goto L20
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v14
	F_errmsg_internal(m, int32(_a_F_get_catalog_object_by_oid_extended_4), v12)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L27
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_get_catalog_object_by_oid_extended_5), int32(2777), int32(_a_F_get_catalog_object_by_oid_extended_6))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L27
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
