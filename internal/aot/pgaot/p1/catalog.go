package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatalogCacheComputeTupleHashValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch l1 - int32(1) {
	case 0:
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v34 = F_fastgetattr_4(m, l2, v31, v14, v12+int32(15))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			v82 = int32(0)
			v83 = v34
			v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v88 = m.T0[v87].(func(*base.Module, int32) int32)(m, v83)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				m.G0 = v12 + int32(16)
				return v88 ^ v82
			}
		}
	case 1:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v20 = F_fastgetattr_4(m, l2, v17, v14, v12+int32(15))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v28 = F_fastgetattr_4(m, l2, v25, v14, v12+int32(15))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v71 = int32(0)
				v72 = v28
				v74 = v20
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v77 = m.T0[v76].(func(*base.Module, int32) int32)(m, v74)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					v82 = base.I32_rotl(v77, int32(8)) ^ v71
					v83 = v72
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v88 = m.T0[v87].(func(*base.Module, int32) int32)(m, v83)
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						m.G0 = v12 + int32(16)
						return v88 ^ v82
					}
				}
			}
		}
	case 2:
		v41 = v4
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v45 = F_fastgetattr_4(m, l2, v42, v14, v12+int32(15))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			v50 = F_fastgetattr_4(m, l2, v47, v14, v12+int32(15))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v55 = F_fastgetattr_4(m, l2, v52, v14, v12+int32(15))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					if l1 == int32(4) {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v60 = m.T0[v59].(func(*base.Module, int32) int32)(m, v41)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v64 = base.I32_rotl(v60, int32(24))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v66 = m.T0[v65].(func(*base.Module, int32) int32)(m, v45)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v71 = v64 ^ base.I32_rotl(v66, int32(16))
								v72 = v55
								v74 = v50
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v77 = m.T0[v76].(func(*base.Module, int32) int32)(m, v74)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									v82 = base.I32_rotl(v77, int32(8)) ^ v71
									v83 = v72
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v88 = m.T0[v87].(func(*base.Module, int32) int32)(m, v83)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(16)
										return v88 ^ v82
									}
								}
							}
						}
					} else {
						v64 = v4
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v66 = m.T0[v65].(func(*base.Module, int32) int32)(m, v45)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v71 = v64 ^ base.I32_rotl(v66, int32(16))
							v72 = v55
							v74 = v50
							v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v77 = m.T0[v76].(func(*base.Module, int32) int32)(m, v74)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								v82 = base.I32_rotl(v77, int32(8)) ^ v71
								v83 = v72
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								v88 = m.T0[v87].(func(*base.Module, int32) int32)(m, v83)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									m.G0 = v12 + int32(16)
									return v88 ^ v82
								}
							}
						}
					}
				}
			}
		}
	case 3:
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v39 = F_fastgetattr_4(m, l2, v36, v14, v12+int32(15))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			v41 = v39
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v45 = F_fastgetattr_4(m, l2, v42, v14, v12+int32(15))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v50 = F_fastgetattr_4(m, l2, v47, v14, v12+int32(15))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v55 = F_fastgetattr_4(m, l2, v52, v14, v12+int32(15))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if l1 == int32(4) {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v60 = m.T0[v59].(func(*base.Module, int32) int32)(m, v41)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v64 = base.I32_rotl(v60, int32(24))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v66 = m.T0[v65].(func(*base.Module, int32) int32)(m, v45)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v71 = v64 ^ base.I32_rotl(v66, int32(16))
									v72 = v55
									v74 = v50
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v77 = m.T0[v76].(func(*base.Module, int32) int32)(m, v74)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										v82 = base.I32_rotl(v77, int32(8)) ^ v71
										v83 = v72
										v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										v88 = m.T0[v87].(func(*base.Module, int32) int32)(m, v83)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											m.G0 = v12 + int32(16)
											return v88 ^ v82
										}
									}
								}
							}
						} else {
							v64 = v4
							v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v66 = m.T0[v65].(func(*base.Module, int32) int32)(m, v45)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v71 = v64 ^ base.I32_rotl(v66, int32(16))
								v72 = v55
								v74 = v50
								v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v77 = m.T0[v76].(func(*base.Module, int32) int32)(m, v74)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									v82 = base.I32_rotl(v77, int32(8)) ^ v71
									v83 = v72
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v88 = m.T0[v87].(func(*base.Module, int32) int32)(m, v83)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(16)
										return v88 ^ v82
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
		v98 = m.ExcPending
		if v98 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
			F_errmsg_internal(m, int32(478605), v12)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(495970), int32(428), int32(345340))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
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
	return base.B2i32(base.Ui32(v2) < base.Ui32(int32(12000)))
}
func F_get_catalog_object_by_oid_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = *(*int32)(unsafe.Add(mBase, _consts[325]))
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L30
	} else {
		goto L45
	}
L2:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if int32(0) < v80 {
		goto L24
	} else {
		goto L25
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v16 == v13 {
		v77 = v15
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v23 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[325])) = v70
	v77 = v70
	goto L2
L8:
	;
	v28 = v23 * int32(40)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_consts[326])))
	if v13 != v31 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v70 = v28 + int32(747872)
	goto L7
L10:
	;
	if v23 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v38 = (v23 | int32(1)) * int32(40)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_consts[326])))
	if v13 == v41 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v70 = v38 + int32(747872)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v48 = (v23 | int32(2)) * int32(40)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[326])))
	if v13 == v51 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v70 = v48 + int32(747872)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v58 = (v23 | int32(3)) * int32(40)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+uint32(_consts[326])))
	if v13 == v61 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v70 = v58 + int32(747872)
	goto L7
L21:
	;
	v23 = v23 + int32(4)
	goto L8
L23:
	;
	m.G0 = v11 - int32(-64)
	return v117
L24:
	;
	if l3 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	F_ScanKeyInit(m, v9+int32(-48), l1, int32(3), int32(184), l2)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L30
	} else {
		goto L33
	}
L27:
	;
	v83 = F_SearchSysCacheLockedCopy1(m, v80, l2)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v88 = F_SearchSysCacheCopy(m, v80, l2, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L30
	} else {
		goto L32
	}
L30:
	;
	return int32(0)
L31:
	;
	v117 = v83
	goto L23
L32:
	;
	v117 = v88
	goto L23
L33:
	;
	v97 = int32(0)
	v98 = int32(1)
	v103 = F_systable_beginscan(m, l0, v90, v98, v97, v98, v9+int32(-48))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v105 = F_systable_getnext(m, v103)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	if v105 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if l3 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v114 = v97
	goto L38
L38:
	;
	F_systable_endscan(m, v103)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L30
	} else {
		goto L44
	}
L39:
	;
	F_LockTuple(m, l0, v105+int32(4), int32(7))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L30
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v112 = F_heap_copytuple(m, v105)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L30
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v114 = v112
	goto L38
L44:
	;
	v117 = v114
	goto L23
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
	F_errmsg_internal(m, int32(59243), v11)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L30
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(490554), int32(2777), int32(500708))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L30
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
