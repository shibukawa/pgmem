package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AlterConstrEnforceabilityRecurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l6)+16))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22+v23)))
	F_ScanKeyInit(m, v17, int32(12), int32(3), int32(184), v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = int32(1)
	v32 = F_systable_beginscan(m, l2, int32(2579), v29, int32(0), v29, v17)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	goto L4
L4:
	;
	v48 = F_systable_getnext(m, v32)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_systable_endscan(m, v32)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L11
	}
L6:
	;
	if v48 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v50 = F_ATExecAlterConstrEnforceability(m, l0, l1, l2, l3, l4, l5, v48, l7, l8, l9, l10, l11)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	goto L5
L10:
	;
	goto L4
L11:
	;
	m.G0 = v17 + int32(48)
	return
}
func F_AlterDomainAddConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = *(*int64)(unsafe.Add(mBase, _consts[224]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[225]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v18
	v21 = F_makeTypeNameFromNameList(m, l1)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		v23 = F_typenameTypeId(m, int32(0), v21)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v27 = F_table_open(m, int32(1247), int32(3))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v31 = F_SearchSysCacheCopy(m, int32(82), v23, int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					if v31 != 0 {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
						F_checkDomainOwner(m, v31)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							if v37 != int32(161) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return
								} else {
									v106 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v106
									F_errmsg_internal(m, int32(478564), v12+int32(16))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										F_errfinish(m, int32(486893), int32(2964), int32(89943))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v40 = v33 + v34
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
								switch v41 - int32(1) {
								case 0:
									v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+130)))
									if v59 != 0 {
										F_sequence_close(m, v27, int32(3))
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											m.G0 = v12 + int32(32)
											return
										}
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v40)+68))
										F_domainAddNotNullConstraint(m, v23, v60, l2, v40+int32(4), l3)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+15)))
											if v65 == int32(0) {
												F_validateDomainNotNullConstraint(m, v23)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													v70 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v40)+130)) = uint8(v70)
													F_CatalogTupleUpdate(m, v27, v31+int32(4), v31)
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v23
														*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
														F_sequence_close(m, v27, int32(3))
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return
														} else {
															m.G0 = v12 + int32(32)
															return
														}
													}
												}
											} else {
												v70 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v40)+130)) = uint8(v70)
												F_CatalogTupleUpdate(m, v27, v31+int32(4), v31)
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v23
													*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
													F_sequence_close(m, v27, int32(3))
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return
													} else {
														m.G0 = v12 + int32(32)
														return
													}
												}
											}
										}
									}
								default:
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v23
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
									F_sequence_close(m, v27, int32(3))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										m.G0 = v12 + int32(32)
										return
									}
								case 4:
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+68))
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)+132))
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+136))
									v49 = F_domainAddCheckConstraint(m, v23, v44, v45, v46, l2, v40+int32(4), l3)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+15)))
										if v51 == int32(0) {
											F_validateDomainCheckConstraint(m, v23, v49)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return
											} else {
												F_CacheInvalidateHeapTuple(m, v27, v31, int32(0))
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v23
													*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
													F_sequence_close(m, v27, int32(3))
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return
													} else {
														m.G0 = v12 + int32(32)
														return
													}
												}
											}
										} else {
											F_CacheInvalidateHeapTuple(m, v27, v31, int32(0))
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v23
												*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
												F_sequence_close(m, v27, int32(3))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													m.G0 = v12 + int32(32)
													return
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
						v92 = m.ExcPending
						if v92 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v23
							F_errmsg_internal(m, int32(49916), v12)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								F_errfinish(m, int32(486893), int32(2956), int32(89943))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
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
		}
	}
}
func F_AlterDomainNotNull(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	v3 = l2
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v19 = *(*int64)(unsafe.Add(mBase, _consts[224]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v19
	v22 = *(*int32)(unsafe.Add(mBase, _consts[225]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v22
	v25 = F_makeTypeNameFromNameList(m, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = F_typenameTypeId(m, v4, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = F_SearchSysCacheCopy(m, int32(82), v27, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L51
	}
L6:
	;
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	F_checkDomainOwner(m, v35)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L48
	}
L10:
	;
	v41 = v38 + v37
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+130)))
	if v3 == v42 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	m.G0 = v16 + int32(32)
	return
L12:
	;
	F_sequence_close(m, v31, int32(3))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v3 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L11
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+130)) = uint8(v3)
	F_CatalogTupleUpdate(m, v31, v35+int32(4), v35)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L41
	}
L17:
	;
	v48 = F_palloc0(m, int32(108))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v64 = m.G0
	v66 = v64 - int32(48)
	m.G0 = v66
	v70 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+104)) = int32(-1)
	v52 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+16)) = uint8(v52)
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(4294967457)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	F_domainAddNotNullConstraint(m, v27, v56, v48, v41+int32(4), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_validateDomainNotNullConstraint(m, v27)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L16
L23:
	;
	F_ScanKeyInit(m, v66, int32(10), int32(3), int32(184), v27)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v78 = int32(1)
	v81 = F_systable_beginscan(m, v70, int32(2665), v78, int32(0), v78, v66)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	F_systable_endscan(m, v81)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L37
	}
L26:
	;
	v83 = F_systable_getnext(m, v81)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v83 == int32(0) {
		v123 = v4
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v90 = v83
	goto L29
L29:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+22)))
	v102 = v100 + v101
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+72)))
	if v103 != int32(110) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v123 = v4
	goto L25
L31:
	;
	v111 = F_systable_getnext(m, v81)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L35
	}
L32:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+76)))
	if v106 != int32(1) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v109 = F_heap_copytuple(m, v90)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v123 = v109
	goto L25
L35:
	;
	if v111 != 0 {
		v90 = v111
		goto L29
	} else {
		goto L36
	}
L36:
	;
	goto L30
L37:
	;
	F_sequence_close(m, v70, int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	m.G0 = v66 + int32(48)
	if v123 == int32(0) {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(2606)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+22)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138+v139)))
	v142 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v141
	F_performDeletion(m, v16+int32(20), v142, v142)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L16
L41:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v170 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v172 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1247), v27, v172, v172, v172)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
	F_pfree(m, v35)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	F_sequence_close(m, v31, int32(3))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L11
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v27
	F_errmsg_internal(m, int32(49916), v16)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(486893), int32(2761), int32(299395))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v41 + int32(4)
	F_errmsg_internal(m, int32(688021), v16+int32(16))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(486893), int32(2796), int32(299395))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_AlterSetting(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(128)
	m.G0 = v14
	v16 = F_ExtractSetVariableArgs(m, l2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = F_table_open(m, int32(2964), int32(3))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v14+int32(32), int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_ScanKeyInit(m, v14+int32(80), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v42 = F_systable_beginscan(m, v20, int32(2965), int32(1), int32(0), int32(2), v14+int32(32))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v44 = F_systable_getnext(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v46 == int32(5) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v265 != 0 {
		goto L70
	} else {
		goto L71
	}
L9:
	;
	if v44 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v44 != 0 {
		goto L46
	} else {
		goto L47
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v54 = F_heap_getattr_4(m, v44, v51, v14+int32(31))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+31)))
	if v56 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_CatalogTupleDelete(m, v20, v44+int32(4))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L45
	}
L15:
	;
	v57 = F_pg_detoast_datum(m, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v59 = m.G0
	v61 = v59 - int32(16)
	m.G0 = v61
	if v57 == int32(0) {
		v146 = v4
		goto L17
	} else {
		goto L18
	}
L17:
	;
	m.G0 = v61 + int32(16)
	if v146 == int32(0) {
		goto L14
	} else {
		goto L42
	}
L18:
	;
	v65 = F_superuser(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v65 != 0 {
		v146 = v4
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v67 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v67
	v72 = v57 + int32(16)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v73 <= int32(0) {
		v146 = v4
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v83 = v4
	goto L22
L22:
	;
	v91 = F_array_ref(m, v57, v61+int32(12), v61+int32(3))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v146 = v132
	goto L17
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v91
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+3)))
	if v94 != 0 {
		v132 = v83
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v135 = v133 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v135 <= v137 {
		v83 = v132
		goto L22
	} else {
		goto L41
	}
L26:
	;
	v95 = F_text_to_cstring(m, v91)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v97 = int32(61)
	v98 = F___strchrnul(m, v95, v97)
	mBase = m.M
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v100 == v97 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v104))) = uint8(v105)
	v109 = F_validate_option_array_item(m, v95, v105, int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L32
	}
L29:
	;
	v104 = v98
	goto L31
L30:
	;
	v104 = int32(0)
	goto L31
L31:
	;
	goto L28
L32:
	;
	if v109 != 0 {
		v132 = v83
		goto L25
	} else {
		goto L33
	}
L33:
	;
	if v83 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v125 + int32(1)
	F_pfree(m, v95)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L40
	}
L35:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v116 = F_array_set(m, v83, v61+int32(8), v113, int32(-1), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v122 = F_construct_array_builtin(m, v61+int32(4), int32(1), int32(25))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v124 = v116
	goto L34
L39:
	;
	v124 = v122
	goto L34
L40:
	;
	v132 = v124
	goto L25
L41:
	;
	goto L23
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v146
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)) = uint8(v156)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+8)) = uint16(v156)
	v160 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+10)) = uint8(v160)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v171 = F_heap_modify_tuple(m, v44, v164, v14+int32(16), v14+int32(12), v14+int32(8))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_CatalogTupleUpdate(m, v20, v44+int32(4), v171)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L8
L45:
	;
	goto L8
L46:
	;
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)) = uint8(v190)
	v192 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+10)) = uint8(v192)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+8)) = uint16(v190)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v199 = F_heap_getattr_4(m, v44, v196, v14+int32(31))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v16 == int32(0) {
		goto L8
	} else {
		goto L66
	}
L49:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+31)))
	if v201 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v204 = F_pg_detoast_datum(m, v199)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v206 = v4
	goto L52
L52:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v16 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v206 = v204
	goto L52
L54:
	;
	if v212 != 0 {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v208 = F_GUCArrayAdd(m, v206, v207, v16)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v210 = F_GUCArrayDelete(m, v206, v207)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v212 = v208
	goto L54
L59:
	;
	v212 = v210
	goto L54
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v212
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v223 = F_heap_modify_tuple(m, v44, v216, v14+int32(16), v14+int32(12), v14+int32(8))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_CatalogTupleDelete(m, v20, v44+int32(4))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	F_CatalogTupleUpdate(m, v20, v44+int32(4), v223)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L8
L65:
	;
	goto L8
L66:
	;
	v233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)) = uint8(v233)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+12)) = uint16(v233)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v239 = F_GUCArrayAdd(m, v233, v238, v16)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l0
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v249 = F_heap_form_tuple(m, v244, v14+int32(16), v14+int32(12))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_CatalogTupleInsert(m, v20, v249)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	goto L8
L70:
	;
	v267 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2964), l0, v267, l1, v267)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	F_systable_endscan(m, v42)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	F_sequence_close(m, v20, int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	m.G0 = v14 + int32(128)
	return
}
func F_transformAlterTableStmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
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
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v648 int32
	_ = v648
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v758 int32
	_ = v758
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v822 int32
	_ = v822
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v886 int32
	_ = v886
	var v899 int32
	_ = v899
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v949 int32
	_ = v949
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	v6 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(208)
	m.G0 = v23
	v26 = F_relation_open(m, l0, v6)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v32 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = l2
	v35 = int32(1)
	v37 = int32(0)
	v40 = F_addRangeTableEntryForRelation(m, v32, v26, v35, v37, v37, v35)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v43 = int32(1)
	F_addNSItemToQuery(m, v32, v40, int32(0), v43, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+136)) = v32
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+119)))
	v51 = base.B2i32(v49 == int32(102))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+156)) = uint8(v51)
	if v49 == int32(102) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v55 = int32(532551)
	goto L8
L7:
	;
	v55 = int32(532528)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+140)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v58 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+168)) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v23)+176)) = v58
	*(*int64)(unsafe.Add(mBase, uint32(v23)+184)) = v58
	v64 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v23)+160)) = v58
	v68 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+157)) = uint8(v68)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+152)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v23)+148)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = v57
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+204)) = uint8(v64)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v64
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+196)) = uint8(base.B2i32(v49 == int32(112)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v82 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v83 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v616 = v64
	v620 = v6
	v621 = v35
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+188)) = int32(0)
	F_transformIndexConstraints(m, v23+int32(136))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L141
	}
L12:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v23)+188))
	v616 = v595
	v620 = v606
	v621 = v600
	goto L11
L13:
	;
	v595 = v64
	v600 = v35
	goto L12
L14:
	;
	goto L15
L15:
	;
	v97 = v64
	v99 = v6
	v102 = v35
	goto L16
L16:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v99<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	switch v113 {
	case 0:
		goto L29
	default:
		goto L20
	case 16:
		goto L28
	case 24:
		goto L27
	case 59, 60:
		goto L24
	case 62:
		goto L26
	case 63:
		goto L25
	}
L17:
	;
	v595 = v571
	v600 = v576
	goto L12
L18:
	;
	v583 = v99 + int32(1)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v583 < v584 {
		v97 = v571
		v99 = v583
		v102 = v576
		goto L16
	} else {
		goto L140
	}
L19:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v505 = F_get_attnum(m, l0, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L121
	}
L20:
	;
	v482 = F_lappend(m, v97, v112)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L120
	}
L21:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v331)+8)) = v477
	goto L20
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L116
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L112
	}
L24:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v23)+148))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+48))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+119)))
	if v334 != int32(73) {
		goto L90
	} else {
		goto L91
	}
L25:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	if v256 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L26:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v229 = F_palloc0(m, int32(68))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L60
	}
L27:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+28))
	if v158 != 0 {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	if v128 == int32(161) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_transformColumnDefinition(m, v23+int32(136), v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v119 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+56)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v116)+28))
	v125 = F_lappend(m, v97, v112)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v571 = v125
	v576 = base.B2i32(v121 == v119) & v102
	goto L18
L32:
	;
	F_transformTableConstraint(m, v23+int32(136), v127)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v571 = v97
	v576 = base.B2i32(v136 != int32(9)) & v102
	goto L18
L36:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v145
	F_errmsg_internal(m, int32(478564), v23+int32(16))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(491952), int32(3631), int32(96437))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v160 = F_transformExpr(m, v32, v158, int32(35))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+131)))
	if v164 != 0 {
		goto L20
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157)+32)) = v160
	goto L41
L43:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v166 = F_get_attnum(m, l0, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v166 == int32(0) {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	if v166 <= int32(0) {
		goto L20
	} else {
		goto L46
	}
L46:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+int32(20)+v172<<(uint(int32(4))%32)+v166*int32(100)-int32(11)))))
	if v181 == int32(0) {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	v185 = F_getIdentitySequence(m, v26, v166, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	v188 = F_typenameTypeId(m, v32, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v191 = F_palloc0(m, int32(16))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = int32(190)
	v195 = F_get_rel_namespace(m, v185)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v197 = F_get_namespace_name(m, v195)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v199 = F_get_rel_name(m, v185)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v202 = F_makeRangeVar(m, v197, v199, int32(-1))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+4)) = v202
	v206 = F_makeTypeNameFromOid(m, v188)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v209 = F_makeDefElem(m, int32(171934), v206, int32(-1))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v23)+132)) = v209
	v216 = F_list_make1_impl(m, int32(1), v23+int32(44))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v218 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v191)+12)) = uint8(v218)
	*(*int32)(unsafe.Add(mBase, uint32(v191)+8)) = v216
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v23)+184))
	v222 = F_lappend(m, v221, v191)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+184)) = v222
	v225 = F_lappend(m, v97, v112)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v571 = v225
	v576 = v102
	goto L18
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = int32(90)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+4)) = v233
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v229)+36)) = uint8(v235)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v229
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v239 = F_get_attnum(m, l0, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v239 == int32(0) {
		goto L22
	} else {
		goto L62
	}
L62:
	;
	v245 = F_get_atttype(m, l0, v239)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v227)+48))
	v248 = int32(1)
	v250 = int32(0)
	F_generateSerialExtraStmts(m, v23+int32(136), v229, v245, v247, v248, v248, v250, v250)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v254 = F_lappend(m, v97, v112)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v571 = v254
	v576 = v102
	goto L18
L66:
	;
	v259 = int32(0)
	v497 = v259
	v499 = v259
	goto L19
L67:
	;
	goto L68
L68:
	;
	v261 = int32(0)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v264 <= v261 {
		v497 = v261
		v499 = v261
		goto L19
	} else {
		goto L69
	}
L69:
	;
	v273 = v261
	v280 = v261
	v282 = v261
	goto L70
L70:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v287+v273<<(uint(int32(2))%32))))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	v293 = int32(441354)
	v296 = int32(*(*uint8)(unsafe.Add(mBase, _consts[246])))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	if v297 == int32(0) {
		v316 = v296
		v317 = v297
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v497 = v325
	v499 = v326
	goto L19
L72:
	;
	v328 = v273 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v328 < v329 {
		v273 = v328
		v280 = v325
		v282 = v326
		goto L70
	} else {
		goto L86
	}
L73:
	;
	if v317-v316 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L74:
	;
	goto L73
L75:
	;
	if v296 != v297 {
		v316 = v296
		v317 = v297
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v301 = v292
	v302 = v293
	goto L77
L77:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+1)))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	if v306 == int32(0) {
		v316 = v305
		v317 = v306
		goto L74
	} else {
		goto L79
	}
L78:
	;
	v316 = v305
	v317 = v306
	goto L74
L79:
	;
	v309 = int32(1)
	if v305 == v306 {
		v301 = v301 + v309
		v302 = v302 + v309
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v321 = F_lappend(m, v280, v291)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v323 = F_lappend(m, v282, v291)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	v325 = v321
	v326 = v282
	goto L72
L85:
	;
	v325 = v280
	v326 = v323
	goto L72
L86:
	;
	goto L71
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L109
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L105
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L101
	}
L90:
	;
	switch v334 - int32(105) {
	case 0:
		goto L88
	default:
		goto L87
	case 7:
		goto L93
	case 9:
		goto L89
	}
L91:
	;
	goto L92
L92:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	if v346 == int32(0) {
		goto L21
	} else {
		goto L96
	}
L93:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	if v339 == int32(0) {
		goto L21
	} else {
		goto L94
	}
L94:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v23)+136))
	v343 = F_transformPartitionBound(m, v342, v332, v339)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v343
	goto L21
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v332)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v356 + int32(4)
	F_errmsg(m, int32(389002), v23+int32(96))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(491952), int32(4248), int32(423033))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v332)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v377 + int32(4)
	F_errmsg(m, int32(445779), v23+int32(112))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(491952), int32(4255), int32(423033))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v332)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v398 + int32(4)
	F_errmsg(m, int32(445749), v23+int32(128))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(491952), int32(4262), int32(423033))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v332)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v416 + int32(4)
	F_errmsg_internal(m, int32(27983), v23+int32(80))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(491952), int32(4267), int32(423033))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v437 + int32(4)
	F_errmsg(m, int32(70796), v23+int32(32))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(491952), int32(3663), int32(96437))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v460 + int32(4)
	F_errmsg(m, int32(70796), v23+int32(48))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(491952), int32(3703), int32(96437))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	v571 = v482
	v576 = v102
	goto L18
L121:
	;
	if v505 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v533 = F_getIdentitySequence(m, v26, v505, int32(1))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L129
	}
L125:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v516 + int32(4)
	F_errmsg(m, int32(70796), v23-int32(-64))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(491952), int32(3745), int32(96437))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	if v533 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v536 = F_palloc0(m, int32(16))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v497
	v560 = F_lappend(m, v97, v112)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L139
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536))) = int32(190)
	v540 = F_get_rel_namespace(m, v533)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v542 = F_get_namespace_name(m, v540)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v544 = F_get_rel_name(m, v533)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v547 = F_makeRangeVar(m, v542, v544, int32(-1))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v549 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v536)+12)) = uint16(v549)
	*(*int32)(unsafe.Add(mBase, uint32(v536)+8)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v536)+4)) = v547
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v23)+184))
	v554 = F_lappend(m, v553, v536)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+184)) = v554
	goto L132
L139:
	;
	v571 = v560
	v576 = v102
	goto L18
L140:
	;
	goto L17
L141:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v23)+172))
	if v633 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v23)+188))
	if v695 == int32(0) {
		v758 = v616
		goto L150
	} else {
		goto L151
	}
L143:
	;
	if v621 == int32(0) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	if v638 <= int32(0) {
		goto L142
	} else {
		goto L145
	}
L145:
	;
	v648 = int32(0)
	goto L146
L146:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v633)+12))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v662+v648<<(uint(int32(2))%32))))
	v667 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v666)+15)) = uint8(v667)
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v666)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v666)+16)) = uint8(v669)
	v672 = v648 + v667
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	if v672 < v673 {
		v648 = v672
		goto L146
	} else {
		goto L148
	}
L147:
	;
	goto L142
L148:
	;
	goto L147
L149:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L189
	}
L150:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v23)+164))
	if v769 == int32(0) {
		v822 = v758
		goto L163
	} else {
		goto L164
	}
L151:
	;
	v698 = int32(0)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v695)+4))
	if v699 <= v698 {
		v758 = v616
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v709 = v698
	v711 = v616
	goto L153
L153:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v695)+12))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v722+v709<<(uint(int32(2))%32))))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	if v727 != int32(204) {
		goto L149
	} else {
		goto L155
	}
L154:
	;
	v758 = v743
	goto L150
L155:
	;
	v730 = F_transformIndexStmt(m, l0, v726, l2)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v733 = F_palloc0(m, int32(32))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v733))) = int32(147)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v730)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v733)+20)) = v730
	if v737 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v741 = int32(21)
	goto L160
L159:
	;
	v741 = int32(14)
	goto L160
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v733)+4)) = v741
	v743 = F_lappend(m, v711, v733)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v746 = v709 + int32(1)
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v695)+4))
	if v746 < v747 {
		v709 = v746
		v711 = v743
		goto L153
	} else {
		goto L162
	}
L162:
	;
	goto L154
L163:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if v833 == int32(0) {
		v886 = v822
		goto L171
	} else {
		goto L172
	}
L164:
	;
	v772 = int32(0)
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v769)+4))
	if v773 <= v772 {
		v822 = v758
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v782 = v772
	v785 = v758
	goto L166
L166:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v769)+12))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v796+v782<<(uint(int32(2))%32))))
	v802 = F_palloc0(m, int32(32))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L168
	}
L167:
	;
	v822 = v807
	goto L163
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+20)) = v800
	*(*int64)(unsafe.Add(mBase, uint32(v802))) = int64(68719476883)
	v807 = F_lappend(m, v785, v802)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v810 = v782 + int32(1)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v769)+4))
	if v810 < v811 {
		v782 = v810
		v785 = v807
		goto L166
	} else {
		goto L170
	}
L170:
	;
	goto L167
L171:
	;
	if v633 == int32(0) {
		v949 = v886
		goto L179
	} else {
		goto L180
	}
L172:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v833)+4))
	if v836 <= int32(0) {
		v886 = v822
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v846 = int32(0)
	v849 = v822
	goto L174
L174:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v833)+12))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v860+v846<<(uint(int32(2))%32))))
	v866 = F_palloc0(m, int32(32))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L176
	}
L175:
	;
	v886 = v871
	goto L171
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v866)+20)) = v864
	*(*int64)(unsafe.Add(mBase, uint32(v866))) = int64(68719476883)
	v871 = F_lappend(m, v849, v866)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v874 = v846 + int32(1)
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v833)+4))
	if v874 < v875 {
		v846 = v874
		v849 = v871
		goto L174
	} else {
		goto L178
	}
L178:
	;
	goto L175
L179:
	;
	F_relation_close(m, v26, int32(0))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L1
	} else {
		goto L187
	}
L180:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	if v899 <= int32(0) {
		v949 = v886
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v909 = int32(0)
	v912 = v886
	goto L182
L182:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v633)+12))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v923+v909<<(uint(int32(2))%32))))
	v929 = F_palloc0(m, int32(32))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L184
	}
L183:
	;
	v949 = v934
	goto L179
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v929)+20)) = v927
	*(*int64)(unsafe.Add(mBase, uint32(v929))) = int64(68719476883)
	v934 = F_lappend(m, v912, v929)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v937 = v909 + int32(1)
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	if v937 < v938 {
		v909 = v937
		v912 = v934
		goto L182
	} else {
		goto L186
	}
L186:
	;
	goto L183
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v949
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v23)+184))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v964
	v967 = F_list_concat(m, int32(0), v620)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v967
	m.G0 = v23 + int32(208)
	return l1
L189:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v978
	F_errmsg_internal(m, int32(468829), v23)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(491952), int32(3838), int32(96437))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
