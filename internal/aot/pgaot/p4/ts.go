package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TSDictionaryIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, int32(76), l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v148
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v20)
	v148 = v3
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(41852), v11)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(520004), int32(2950), int32(69586))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v39 = v35 + v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v40 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L47
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[388]))
	v45 = int32(0)
	if v44 == v45 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[388]))
	if v87 == int32(0) {
		v138 = v3
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v83 == int32(0) {
		v138 = v3
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v83 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v51 <= int32(0) {
		v76 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v76
	goto L18
L23:
	;
	v54 = int32(0)
	if v54 < v51 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v57 = v51
	goto L26
L25:
	;
	v57 = v54
	goto L26
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v60 = int32(0)
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58+v60<<(uint(int32(2))%32))))
	v69 = base.B2i32(v68 == v40)
	if v68 == v40 {
		v76 = v69
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v76 = v69
	goto L22
L29:
	;
	v71 = v60 + int32(1)
	if v71 != v57 {
		v60 = v71
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v90 < v91 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v96 = v90
	goto L36
L34:
	;
	goto L35
L35:
	;
	v138 = v3
	goto L14
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v96<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, _consts[260]))
	if v108 != v110 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	if v108 == v40 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v120 = v96 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v120 < v121 {
		v96 = v120
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v138 = int32(1)
	goto L14
L42:
	;
	goto L43
L43:
	;
	v115 = int32(0)
	v117 = F_SearchSysCacheExists(m, int32(75), v39+int32(4), v108, v115, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	if v117 != 0 {
		v138 = v3
		goto L14
	} else {
		goto L45
	}
L45:
	;
	goto L40
L46:
	;
	goto L37
L47:
	;
	v148 = v138
	goto L1
}
func F_check_commit_ts_buffers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_check_slru_buffers(m, int32(142600), l0)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_getTSCurrentConfig(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[1391]))
	if v8 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[1392]))
		if v12 != 0 {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v13 != 0 {
				v30 = *(*int32)(unsafe.Add(mBase, _consts[1393]))
				if v30 != 0 {
					v57 = *(*int32)(unsafe.Add(mBase, _consts[1392]))
					v59 = F_stringToQualifiedNameList(m, v57, int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						v62 = F_get_ts_config_oid(m, v59, int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[1391])) = v62
							v65 = v62
							m.G0 = v5 + int32(48)
							return v65
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v5)+16)) = int64(85899345924)
					v37 = F_hash_create(m, int32(415378), int32(16), v5, int32(40))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[1393])) = v37
						F_CacheRegisterSyscacheCallback(m, int32(74), int32(1613), v37)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, _consts[1393]))
							F_CacheRegisterSyscacheCallback(m, int32(72), int32(1613), v47)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, _consts[382]))
								if v51 != 0 {
									v57 = *(*int32)(unsafe.Add(mBase, _consts[1392]))
									v59 = F_stringToQualifiedNameList(m, v57, int32(0))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v62 = F_get_ts_config_oid(m, v59, int32(0))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[1391])) = v62
											v65 = v62
											m.G0 = v5 + int32(48)
											return v65
										}
									}
								} else {
									F_CreateCacheMemoryContext(m)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										v57 = *(*int32)(unsafe.Add(mBase, _consts[1392]))
										v59 = F_stringToQualifiedNameList(m, v57, int32(0))
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v62 = F_get_ts_config_oid(m, v59, int32(0))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[1391])) = v62
												v65 = v62
												m.G0 = v5 + int32(48)
												return v65
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(112787), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519722), int32(568), int32(351570))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(112787), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(519722), int32(568), int32(351570))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
		v65 = v8
		m.G0 = v5 + int32(48)
		return v65
	}
}
func F_get_ts_dict_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	F_DeconstructQualifiedName(m, l0, v10+int32(12), v10+int32(8))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	m.G0 = v10 + int32(16)
	return v106
L4:
	;
	if l1 != 0 {
		v106 = v83
		goto L3
	} else {
		goto L26
	}
L5:
	;
	v83 = int32(0)
	goto L4
L6:
	;
	v22 = F_LookupExplicitNamespace(m, v20, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L15
	}
L9:
	;
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v24 = int32(0)
	goto L12
L11:
	;
	v24 = l1
	goto L12
L12:
	;
	if v24 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v27 = int32(0)
	v29 = F_GetSysCacheOid(m, int32(75), v26, v22, v27, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v83 = v29
	goto L4
L15:
	;
	v33 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[388]))
	if v35 == v33 {
		v83 = v33
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v38 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v39 <= v38 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v46 = v38
	goto L18
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v46<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, _consts[260]))
	if v54 != v56 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L5
L20:
	;
	v59 = int32(0)
	v61 = F_GetSysCacheOid(m, int32(75), v42, v54, v59, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v65 = v46 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v65 < v66 {
		v46 = v65
		goto L18
	} else {
		goto L25
	}
L23:
	;
	if v61 != 0 {
		v106 = v61
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L19
L26:
	;
	if v83 != 0 {
		v106 = v83
		goto L3
	} else {
		goto L27
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v91 = F_NameListToString(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v91
	F_errmsg(m, int32(75911), v10)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(520004), int32(2910), int32(450909))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_ts_template_func(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = F_defGetQualifiedName(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int64(9796820404457)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v13
		v18 = int32(4)
		if l1 == v18 {
			v21 = int32(1)
		} else {
			v21 = v18
		}
		v25 = F_LookupFuncName(m, v9, v21, v7+int32(16), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = F_get_func_rettype(m, v25)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 != int32(2281) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(117833860))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v41 = F_func_signature_string(m, v9, v21, int32(0), v7+int32(16))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v44 = F_format_type_be(m, int32(2281))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v44
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v41
									F_errmsg(m, int32(200496), v7)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(514289), int32(643), int32(508640))
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
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
				} else {
					m.G0 = v7 + int32(32)
					return v25
				}
			}
		}
	}
}
func F_ts_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if base.Ui64(int64(2)) <= base.Ui64(v4-int64(9223372036854775807)) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
		if base.Ui64(int64(1)) < base.Ui64(v10-int64(9223372036854775807)) {
			v28 = F_DirectFunctionCall2Coll(m, int32(1517), int32(0), v3, v9)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = F_abs_interval(m, v28)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					return v30
				}
			}
		} else {
			v17 = F_palloc(m, int32(16))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(9223372036854775807)
				*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(9223372034707292159)
				return v17
			}
		}
	} else {
		v17 = F_palloc(m, int32(16))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(9223372036854775807)
			*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(9223372034707292159)
			return v17
		}
	}
}
func F_ts_headline(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = F_getTSCurrentConfig(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_DirectFunctionCall3Coll(m, int32(1187), int32(0), v4, v8, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_ts_headline_json_byid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = F_DirectFunctionCall3Coll(m, int32(1190), int32(0), v4, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_ts_headline_json_byid_opt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v26 < int32(4) {
		v36 = int32(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v38 = F_palloc0(m, int32(24))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v29 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v30 == v29 {
		v36 = v29
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v33 = F_pg_detoast_datum(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v36 = v33
	goto L3
L7:
	;
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+36)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v16)+28)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+20)) = v40
	v51 = F_palloc(m, int32(512))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v16 + int32(12)
	v57 = F_lookup_ts_config_cache(m, v18)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v61 = F_lookup_ts_parser_cache(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v61
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v65 = F_deserialize_deflist(m, v36)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v68 = v61
	v69 = int32(0)
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	if v71 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v68 = v67
	v69 = v65
	goto L13
L15:
	;
	v72 = m.G0
	v74 = v72 - int32(80)
	m.G0 = v74
	v77 = F_palloc0(m, int32(40))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L63
	}
L18:
	;
	v80 = F_palloc0(m, int32(16))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v84 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v86 = int32(1)
	v87 = v84 + v86
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v92 = v90 & v86
	if v92 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v93 = v87
	goto L23
L22:
	;
	v93 = v84 + int32(4)
	goto L23
L23:
	;
	if v90 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	goto L35
L25:
	;
	v96 = int32(4)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v98&int32(254) == int32(2) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v111 = int32(1)
	if v92 != 0 {
		v121 = int32(base.Ui32(v90)>>(uint(v111)%32)) - v111
		goto L24
	} else {
		goto L34
	}
L28:
	;
	v107 = v96
	goto L30
L29:
	;
	v107 = base.B2i32(v98 == int32(18)) << (uint(v96) % 32)
	goto L30
L30:
	;
	if v98 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v110 = v96
	goto L33
L32:
	;
	v110 = v107
	goto L33
L33:
	;
	v121 = v110
	goto L24
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v121 = int32(base.Ui32(v115)>>(uint(int32(2))%32)) - int32(4)
	goto L24
L35:
	;
	v126 = F_makeJsonLexContextCstringLen(m, v74+int32(12), v93, v121, v124, int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v126
	v129 = F_makeStringInfo(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = int32(1188)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v77)+36)) = int32(1396)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = int32(1397)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(1398)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = int32(1399)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = int32(1400)
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v77)+28)) = int32(1401)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = int32(1402)
	v152 = F_pg_parse_json(m, v74+int32(12), v77)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v152 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_json_errsave_error(m, v152, v74+int32(12), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_freeJsonLexContext(m, v74+int32(12))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v166 = F_cstring_to_text_with_len(m, v164, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	m.G0 = v74 + int32(80)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v171 != v20 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_pfree(m, v20)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v175 != v24 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	F_pfree(m, v24)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v36 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	F_pfree(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v36 == v181 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	F_pfree(m, v36)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+20)))
	if v188 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	F_pfree(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	m.G0 = v16 + int32(48)
	return v166
L61:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	F_pfree(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(275344), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(515233), int32(471), int32(89593))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ts_headline_jsonb_byid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = F_DirectFunctionCall3Coll(m, int32(1189), int32(0), v4, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_ts_headline_jsonb_byid_opt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v24 < int32(4) {
		v34 = int32(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v36 = F_palloc0(m, int32(24))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v27 = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v28 == v27 {
		v34 = v27
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v31 = F_pg_detoast_datum(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v34 = v31
	goto L3
L7:
	;
	v38 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+36)) = v38
	*(*int64)(unsafe.Add(mBase, uint32(v14)+28)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+20)) = v38
	v49 = F_palloc(m, int32(512))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v14 + int32(12)
	v55 = F_lookup_ts_config_cache(m, v16)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v59 = F_lookup_ts_parser_cache(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v59
	if v34 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v63 = F_deserialize_deflist(m, v34)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v66 = v59
	v67 = v2
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	if v69 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v66 = v65
	v67 = v63
	goto L13
L15:
	;
	v70 = m.G0
	v72 = v70 - int32(32)
	m.G0 = v72
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = int32(0)
	v78 = F_JsonbIteratorInit(m, v18+int32(4))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L73
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+28)) = v78
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+8)))
	v87 = F_JsonbIteratorNext(m, v72+int32(28), v72+int32(8), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v87 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v90 = v87
	goto L23
L21:
	;
	v180 = v2
	goto L22
L22:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	if v188 == int32(16) {
		goto L51
	} else {
		goto L52
	}
L23:
	;
	if v90&int32(-2) != int32(2) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v180 = v168
	goto L22
L25:
	;
	v168 = F_pushJsonbValue(m, v72+int32(4), v90, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L48
	}
L26:
	;
	if base.Ui32(v90) < base.Ui32(int32(4)) {
		goto L45
	} else {
		goto L46
	}
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v106 != int32(1) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v111 = F_headline_json_value(m, v36, v109, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v113 = F_pg_detoast_datum_packed(m, v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v115 = int32(1)
	v116 = v113 + v115
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v119&v115 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v122 = v116
	goto L33
L32:
	;
	v122 = v113 + int32(4)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+16)) = v122
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v124 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v154
	v167 = v72 + int32(8)
	goto L25
L35:
	;
	v127 = int32(4)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v129&int32(254) == int32(2) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v142 = int32(1)
	if v124&v142 != 0 {
		v154 = int32(base.Ui32(v124)>>(uint(v142)%32)) - v142
		goto L34
	} else {
		goto L44
	}
L38:
	;
	v138 = v127
	goto L40
L39:
	;
	v138 = base.B2i32(v129 == int32(18)) << (uint(v127) % 32)
	goto L40
L40:
	;
	if v129 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v141 = v127
	goto L43
L42:
	;
	v141 = v138
	goto L43
L43:
	;
	v154 = v141
	goto L34
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v154 = int32(base.Ui32(v148)>>(uint(int32(2))%32)) - int32(4)
	goto L34
L45:
	;
	v163 = v72 + int32(8)
	goto L47
L46:
	;
	v163 = int32(0)
	goto L47
L47:
	;
	v167 = v163
	goto L25
L48:
	;
	v175 = F_JsonbIteratorNext(m, v72+int32(28), v72+int32(8), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if v175 != 0 {
		v90 = v175
		goto L23
	} else {
		goto L50
	}
L50:
	;
	goto L24
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v180)+12)) = uint8(v81)
	goto L53
L52:
	;
	goto L53
L53:
	;
	v192 = F_JsonbValueToJsonb(m, v180)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	m.G0 = v72 + int32(32)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v197 != v18 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_pfree(m, v18)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v201 != v22 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	F_pfree(m, v22)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v34 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_pfree(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L67
	}
L64:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v34 == v207 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	F_pfree(m, v34)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+20)))
	if v214 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	F_pfree(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	m.G0 = v14 + int32(48)
	return v192
L71:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	F_pfree(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(275344), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(515233), int32(394), int32(89640))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ts_parse_byid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	if v5 != 0 {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
		v26 = F_prs_process_call(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v26 != 0 {
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
				*(*int64)(unsafe.Add(mBase, uint32(v25))) = v28 + int64(1)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = int32(1)
				return v26
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(2)
					v41 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
					return v26
				}
			}
		}
	} else {
		v7 = l0 + int32(28)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v9 = F_pg_detoast_datum_packed(m, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				F_prs_setup_firstcall(m, v13, l0, v15, v9)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					if v9 == v18 {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
						v26 = F_prs_process_call(m, v25)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							if v26 != 0 {
								v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
								*(*int64)(unsafe.Add(mBase, uint32(v25))) = v28 + int64(1)
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = int32(1)
								return v26
							} else {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(2)
									v41 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
									return v26
								}
							}
						}
					} else {
						F_pfree(m, v9)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
							v26 = F_prs_process_call(m, v25)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								if v26 != 0 {
									v28 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
									*(*int64)(unsafe.Add(mBase, uint32(v25))) = v28 + int64(1)
									v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = int32(1)
									return v26
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int32(0)
									} else {
										v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(2)
										v41 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
										return v26
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
func F_ts_rank_wtt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 float32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = l0 + int32(28)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v20 = F_pg_detoast_datum(m, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			F_getWeights(m, v13, v10)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v26 = F_calc_rank(m, v10, v20, v22, int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v28 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
							if v32 != v20 {
								F_pfree(m, v20)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v36 != v22 {
										F_pfree(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											m.G0 = v10 + int32(16)
											return base.I32_reinterpret_f32(v26)
										}
									} else {
										m.G0 = v10 + int32(16)
										return base.I32_reinterpret_f32(v26)
									}
								}
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v36 != v22 {
									F_pfree(m, v22)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I32_reinterpret_f32(v26)
									}
								} else {
									m.G0 = v10 + int32(16)
									return base.I32_reinterpret_f32(v26)
								}
							}
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
						if v32 != v20 {
							F_pfree(m, v20)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v36 != v22 {
									F_pfree(m, v22)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return base.I32_reinterpret_f32(v26)
									}
								} else {
									m.G0 = v10 + int32(16)
									return base.I32_reinterpret_f32(v26)
								}
							}
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v36 != v22 {
								F_pfree(m, v22)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(16)
									return base.I32_reinterpret_f32(v26)
								}
							} else {
								m.G0 = v10 + int32(16)
								return base.I32_reinterpret_f32(v26)
							}
						}
					}
				}
			}
		}
	}
}
func F_ts_stat2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = l0 + int32(28)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v19 = F_pg_detoast_datum_packed(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_SPI_connect_ext(m, int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
						v27 = F_ts_stat_sql(m, v26, v12, v19)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v29 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
									if v33 != v19 {
										F_pfree(m, v19)
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return int32(0)
										} else {
											F_ts_setup_firstcall(m, l0, v21, v27)
											mBase = m.M
											v38 = m.ExcPending
											if v38 != 0 {
												return int32(0)
											} else {
												v39 = F_SPI_finish(m)
												mBase = m.M
												v40 = m.ExcPending
												if v40 != 0 {
													return int32(0)
												} else {
													v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
													v48 = F_ts_process_call(m, v47)
													mBase = m.M
													v49 = m.ExcPending
													if v49 != 0 {
														return int32(0)
													} else {
														if v48 != 0 {
															v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
															*(*int64)(unsafe.Add(mBase, uint32(v47))) = v50 + int64(1)
															v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = int32(1)
															return v48
														} else {
															F_end_MultiFuncCall(m, l0)
															mBase = m.M
															v59 = m.ExcPending
															if v59 != 0 {
																return int32(0)
															} else {
																v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = int32(2)
																v63 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v63)
																return v48
															}
														}
													}
												}
											}
										}
									} else {
										F_ts_setup_firstcall(m, l0, v21, v27)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return int32(0)
										} else {
											v39 = F_SPI_finish(m)
											mBase = m.M
											v40 = m.ExcPending
											if v40 != 0 {
												return int32(0)
											} else {
												v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
												v48 = F_ts_process_call(m, v47)
												mBase = m.M
												v49 = m.ExcPending
												if v49 != 0 {
													return int32(0)
												} else {
													if v48 != 0 {
														v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
														*(*int64)(unsafe.Add(mBase, uint32(v47))) = v50 + int64(1)
														v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = int32(1)
														return v48
													} else {
														F_end_MultiFuncCall(m, l0)
														mBase = m.M
														v59 = m.ExcPending
														if v59 != 0 {
															return int32(0)
														} else {
															v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = int32(2)
															v63 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v63)
															return v48
														}
													}
												}
											}
										}
									}
								}
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
								if v33 != v19 {
									F_pfree(m, v19)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										F_ts_setup_firstcall(m, l0, v21, v27)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return int32(0)
										} else {
											v39 = F_SPI_finish(m)
											mBase = m.M
											v40 = m.ExcPending
											if v40 != 0 {
												return int32(0)
											} else {
												v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
												v48 = F_ts_process_call(m, v47)
												mBase = m.M
												v49 = m.ExcPending
												if v49 != 0 {
													return int32(0)
												} else {
													if v48 != 0 {
														v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
														*(*int64)(unsafe.Add(mBase, uint32(v47))) = v50 + int64(1)
														v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = int32(1)
														return v48
													} else {
														F_end_MultiFuncCall(m, l0)
														mBase = m.M
														v59 = m.ExcPending
														if v59 != 0 {
															return int32(0)
														} else {
															v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = int32(2)
															v63 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v63)
															return v48
														}
													}
												}
											}
										}
									}
								} else {
									F_ts_setup_firstcall(m, l0, v21, v27)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int32(0)
									} else {
										v39 = F_SPI_finish(m)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int32(0)
										} else {
											v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
											v48 = F_ts_process_call(m, v47)
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												if v48 != 0 {
													v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
													*(*int64)(unsafe.Add(mBase, uint32(v47))) = v50 + int64(1)
													v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = int32(1)
													return v48
												} else {
													F_end_MultiFuncCall(m, l0)
													mBase = m.M
													v59 = m.ExcPending
													if v59 != 0 {
														return int32(0)
													} else {
														v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = int32(2)
														v63 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v63)
														return v48
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
		}
	} else {
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
		v48 = F_ts_process_call(m, v47)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			if v48 != 0 {
				v50 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
				*(*int64)(unsafe.Add(mBase, uint32(v47))) = v50 + int64(1)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = int32(1)
				return v48
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = int32(2)
					v63 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v63)
					return v48
				}
			}
		}
	}
}
