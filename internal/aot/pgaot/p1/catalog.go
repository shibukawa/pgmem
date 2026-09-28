package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatalogCacheComputeTupleHashValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch l1 - int32(1) {
	case 0:
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v33 = F_fastgetattr_4(m, l2, v30, v15, v13+int32(15))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v79 = int32(0)
			v81 = v33
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v86 = m.T0[v85].(func(*base.Module, int64) int32)(m, v81)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				m.G0 = v13 + int32(16)
				return v86 ^ v79
			}
		}
	case 1:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v20 = v13 + int32(15)
		v21 = F_fastgetattr_4(m, l2, v18, v15, v20)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v27 = F_fastgetattr_4(m, l2, v26, v15, v20)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v67 = int32(0)
				v69 = v27
				v70 = v21
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v74 = m.T0[v73].(func(*base.Module, int64) int32)(m, v70)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					v79 = base.I32_rotl(v74, int32(8)) ^ v67
					v81 = v69
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v86 = m.T0[v85].(func(*base.Module, int64) int32)(m, v81)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						m.G0 = v13 + int32(16)
						return v86 ^ v79
					}
				}
			}
		}
	case 2:
		v40 = int64(0)
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v43 = v13 + int32(15)
		v44 = F_fastgetattr_4(m, l2, v41, v15, v43)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v47 = F_fastgetattr_4(m, l2, v46, v15, v43)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v50 = F_fastgetattr_4(m, l2, v49, v15, v43)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					if l1 == int32(4) {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v55 = m.T0[v54].(func(*base.Module, int64) int32)(m, v40)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v60 = base.I32_rotl(v55, int32(24))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v62 = m.T0[v61].(func(*base.Module, int64) int32)(m, v44)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v67 = v60 ^ base.I32_rotl(v62, int32(16))
								v69 = v50
								v70 = v47
								v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v74 = m.T0[v73].(func(*base.Module, int64) int32)(m, v70)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									v79 = base.I32_rotl(v74, int32(8)) ^ v67
									v81 = v69
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v86 = m.T0[v85].(func(*base.Module, int64) int32)(m, v81)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										m.G0 = v13 + int32(16)
										return v86 ^ v79
									}
								}
							}
						}
					} else {
						v60 = int32(0)
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v62 = m.T0[v61].(func(*base.Module, int64) int32)(m, v44)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v67 = v60 ^ base.I32_rotl(v62, int32(16))
							v69 = v50
							v70 = v47
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v74 = m.T0[v73].(func(*base.Module, int64) int32)(m, v70)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v79 = base.I32_rotl(v74, int32(8)) ^ v67
								v81 = v69
								v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v86 = m.T0[v85].(func(*base.Module, int64) int32)(m, v81)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									m.G0 = v13 + int32(16)
									return v86 ^ v79
								}
							}
						}
					}
				}
			}
		}
	case 3:
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v38 = F_fastgetattr_4(m, l2, v35, v15, v13+int32(15))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = v38
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v43 = v13 + int32(15)
			v44 = F_fastgetattr_4(m, l2, v41, v15, v43)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v47 = F_fastgetattr_4(m, l2, v46, v15, v43)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v50 = F_fastgetattr_4(m, l2, v49, v15, v43)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						if l1 == int32(4) {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v55 = m.T0[v54].(func(*base.Module, int64) int32)(m, v40)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								v60 = base.I32_rotl(v55, int32(24))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v62 = m.T0[v61].(func(*base.Module, int64) int32)(m, v44)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v67 = v60 ^ base.I32_rotl(v62, int32(16))
									v69 = v50
									v70 = v47
									v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v74 = m.T0[v73].(func(*base.Module, int64) int32)(m, v70)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										v79 = base.I32_rotl(v74, int32(8)) ^ v67
										v81 = v69
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v86 = m.T0[v85].(func(*base.Module, int64) int32)(m, v81)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											m.G0 = v13 + int32(16)
											return v86 ^ v79
										}
									}
								}
							}
						} else {
							v60 = int32(0)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v62 = m.T0[v61].(func(*base.Module, int64) int32)(m, v44)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v67 = v60 ^ base.I32_rotl(v62, int32(16))
								v69 = v50
								v70 = v47
								v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v74 = m.T0[v73].(func(*base.Module, int64) int32)(m, v70)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									v79 = base.I32_rotl(v74, int32(8)) ^ v67
									v81 = v69
									v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v86 = m.T0[v85].(func(*base.Module, int64) int32)(m, v81)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										m.G0 = v13 + int32(16)
										return v86 ^ v79
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
		v96 = m.ExcPending
		if v96 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
			F_errmsg_internal(m, int32(_a_F_CatalogCacheComputeTupleHashValue_0), v13)
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_CatalogCacheComputeTupleHashValue_1), int32(440), int32(_a_F_CatalogCacheComputeTupleHashValue_2))
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_get_catalog_object_by_oid_extended[0]))
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L27
	} else {
		goto L43
	}
L2:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	if int32(0) <= v68 {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v18 == v15 {
		v62 = v17
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v28 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_catalog_object_by_oid_extended[0])) = v56
	v62 = v56
	goto L2
L8:
	;
	v56 = v32 + int32(_a_F_get_catalog_object_by_oid_extended_0)
	goto L7
L9:
	;
	v56 = v32 + int32(_a_F_get_catalog_object_by_oid_extended_1)
	goto L7
L10:
	;
	v56 = v32 + int32(_a_F_get_catalog_object_by_oid_extended_2)
	goto L7
L11:
	;
	v32 = v28 * int32(40)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_get_catalog_object_by_oid_extended[1])))
	if v15 != v33 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v56 = v32 + int32(_a_F_get_catalog_object_by_oid_extended_3)
	goto L7
L13:
	;
	if v28 == int32(36) {
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
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_get_catalog_object_by_oid_extended[2])))
	if v39 == v15 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_get_catalog_object_by_oid_extended[3])))
	if v41 == v15 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_get_catalog_object_by_oid_extended[4])))
	if v43 == v15 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v28 = v28 + int32(4)
	goto L11
L20:
	;
	m.G0 = v13 - int32(-64)
	return v110
L21:
	;
	v71 = base.I64_extend_i32_u(l2)
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
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v81 = v11 + int32(-56)
	F_ScanKeyInit(m, v81, l1, int32(3), int32(184), base.I64_extend_i32_u(l2))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L27
	} else {
		goto L30
	}
L24:
	;
	v72 = F_SearchSysCacheLockedCopy1(m, v68, v71)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v77 = F_SearchSysCacheCopy(m, v68, v71, int64(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L27
	} else {
		goto L29
	}
L27:
	;
	return int32(0)
L28:
	;
	v110 = v72
	goto L20
L29:
	;
	v110 = v77
	goto L20
L30:
	;
	v87 = int32(1)
	v90 = F_systable_beginscan(m, l0, v79, v87, int32(0), v87, v81)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v92 = F_systable_getnext(m, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	if v92 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_systable_endscan(m, v90)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L27
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if l3 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v110 = int32(0)
	goto L20
L37:
	;
	F_LockTuple(m, l0, v92+int32(4), int32(7))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L27
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v104 = F_heap_copytuple(m, v92)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L27
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	F_systable_endscan(m, v90)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L27
	} else {
		goto L42
	}
L42:
	;
	v110 = v104
	goto L20
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v15
	F_errmsg_internal(m, int32(_a_F_get_catalog_object_by_oid_extended_4), v13)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L27
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_get_catalog_object_by_oid_extended_5), int32(2827), int32(_a_F_get_catalog_object_by_oid_extended_6))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L27
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
