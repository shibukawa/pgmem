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
	var v15 int32
	_ = v15
	var v18 int64
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
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_AlterDomainAddConstraint[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15
	v18 = *(*int64)(unsafe.Add(mBase, _c_F_AlterDomainAddConstraint[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v18
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
									F_errmsg_internal(m, int32(_a_F_AlterDomainAddConstraint_0), v12+int32(16))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_AlterDomainAddConstraint_1), int32(2964), int32(_a_F_AlterDomainAddConstraint_2))
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
										F_relation_close(m, v27, int32(3))
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
														F_relation_close(m, v27, int32(3))
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
													F_relation_close(m, v27, int32(3))
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
									F_relation_close(m, v27, int32(3))
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
													F_relation_close(m, v27, int32(3))
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
												F_relation_close(m, v27, int32(3))
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
							F_errmsg_internal(m, int32(_a_F_AlterDomainAddConstraint_3), v12)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_AlterDomainAddConstraint_1), int32(2956), int32(_a_F_AlterDomainAddConstraint_2))
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
	var v19 int32
	_ = v19
	var v22 int64
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	v3 = l2
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_AlterDomainNotNull[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v19
	v22 = *(*int64)(unsafe.Add(mBase, _c_F_AlterDomainNotNull[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v22
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
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L50
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
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L47
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
	F_relation_close(m, v31, int32(3))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L46
	}
L12:
	;
	goto L11
L13:
	;
	goto L14
L14:
	;
	if v3 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+130)) = uint8(v3)
	F_CatalogTupleUpdate(m, v31, v35+int32(4), v35)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L40
	}
L16:
	;
	v45 = F_palloc0(m, int32(108))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v61 = m.G0
	v63 = v61 - int32(48)
	m.G0 = v63
	v67 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+104)) = int32(-1)
	v49 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+16)) = uint8(v49)
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = int64(4294967457)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	F_domainAddNotNullConstraint(m, v27, v53, v45, v41+int32(4), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_validateDomainNotNullConstraint(m, v27)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	F_ScanKeyInit(m, v63, int32(10), int32(3), int32(184), v27)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v75 = int32(1)
	v78 = F_systable_beginscan(m, v67, int32(2665), v75, int32(0), v75, v63)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	F_systable_endscan(m, v78)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L36
	}
L25:
	;
	v80 = F_systable_getnext(m, v78)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v80 == int32(0) {
		v122 = v4
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v92 = v80
	goto L28
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+22)))
	v99 = v97 + v98
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+72)))
	if v100 != int32(110) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v122 = v4
	goto L24
L30:
	;
	v108 = F_systable_getnext(m, v78)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L34
	}
L31:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+76)))
	if v103 != int32(1) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v106 = F_heap_copytuple(m, v92)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v122 = v106
	goto L24
L34:
	;
	if v108 != 0 {
		v92 = v108
		goto L28
	} else {
		goto L35
	}
L35:
	;
	goto L29
L36:
	;
	F_relation_close(m, v67, int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	m.G0 = v63 + int32(48)
	if v122 == int32(0) {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(2606)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+22)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v135+v136)))
	v139 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v138
	F_performDeletion(m, v16+int32(20), v139, v139)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L15
L40:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_AlterDomainNotNull[2]))
	if v167 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v169 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1247), v27, v169, v169, v169)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
	F_pfree(m, v35)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	goto L11
L46:
	;
	m.G0 = v16 + int32(32)
	return
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v27
	F_errmsg_internal(m, int32(_a_F_AlterDomainNotNull_0), v16)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_AlterDomainNotNull_1), int32(2761), int32(_a_F_AlterDomainNotNull_2))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v41 + int32(4)
	F_errmsg_internal(m, int32(_a_F_AlterDomainNotNull_3), v16+int32(16))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_AlterDomainNotNull_1), int32(2796), int32(_a_F_AlterDomainNotNull_2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
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
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	v15 = F_ExtractSetVariableArgs(m, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = F_table_open(m, int32(2964), int32(3))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = v13 + int32(32)
	F_ScanKeyInit(m, v22, int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_ScanKeyInit(m, v13+int32(80), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v39 = F_systable_beginscan(m, v19, int32(2965), int32(1), int32(0), int32(2), v22)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v41 = F_systable_getnext(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v43 == int32(5) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSetting[0]))
	if v256 != 0 {
		goto L70
	} else {
		goto L71
	}
L9:
	;
	if v41 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v41 != 0 {
		goto L46
	} else {
		goto L47
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v51 = F_heap_getattr_4(m, v41, v48, v13+int32(31))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+31)))
	if v53 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_simple_heap_delete(m, v19, v41+int32(4))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L45
	}
L15:
	;
	v54 = F_pg_detoast_datum(m, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v56 = m.G0
	v58 = v56 - int32(16)
	m.G0 = v58
	if v54 == int32(0) {
		v141 = v4
		goto L17
	} else {
		goto L18
	}
L17:
	;
	m.G0 = v58 + int32(16)
	if v141 == int32(0) {
		goto L14
	} else {
		goto L42
	}
L18:
	;
	v62 = F_superuser(m)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v62 != 0 {
		v141 = v4
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v64 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v64
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v68 <= int32(0) {
		v141 = v4
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v79 = v4
	goto L22
L22:
	;
	v85 = F_array_ref(m, v54, v58+int32(12), v58+int32(3))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v141 = v126
	goto L17
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v85
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+3)))
	if v88 != 0 {
		v126 = v79
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v129 = v127 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v129 <= v131 {
		v79 = v126
		goto L22
	} else {
		goto L41
	}
L26:
	;
	v89 = F_text_to_cstring(m, v85)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v91 = int32(61)
	v92 = F___strchrnul(m, v89, v91)
	mBase = m.M
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v94 == v91 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v99 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(v99)
	v103 = F_validate_option_array_item(m, v89, v99, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L32
	}
L29:
	;
	v98 = v92
	goto L31
L30:
	;
	v98 = int32(0)
	goto L31
L31:
	;
	goto L28
L32:
	;
	if v103 != 0 {
		v126 = v79
		goto L25
	} else {
		goto L33
	}
L33:
	;
	if v79 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v119 + int32(1)
	F_pfree(m, v89)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L40
	}
L35:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v110 = F_array_set(m, v79, v58+int32(8), v107, int32(-1), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v116 = F_construct_array_builtin(m, v58+int32(4), int32(1), int32(25))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v118 = v110
	goto L34
L39:
	;
	v118 = v116
	goto L34
L40:
	;
	v126 = v118
	goto L25
L41:
	;
	goto L23
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v141
	v149 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v149)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)) = uint16(v149)
	v153 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v153)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v164 = F_heap_modify_tuple(m, v41, v157, v13+int32(16), v13+int32(12), v13+int32(8))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_CatalogTupleUpdate(m, v19, v41+int32(4), v164)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
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
	v182 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v182)
	v184 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v184)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)) = uint16(v182)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v191 = F_heap_getattr_4(m, v41, v188, v13+int32(31))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v15 == int32(0) {
		goto L8
	} else {
		goto L66
	}
L49:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+31)))
	if v193 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v196 = F_pg_detoast_datum(m, v191)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v198 = v4
	goto L52
L52:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v15 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v198 = v196
	goto L52
L54:
	;
	if v204 != 0 {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v200 = F_GUCArrayAdd(m, v198, v199, v15)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v202 = F_GUCArrayDelete(m, v198, v199)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v204 = v200
	goto L54
L59:
	;
	v204 = v202
	goto L54
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v204
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v215 = F_heap_modify_tuple(m, v41, v208, v13+int32(16), v13+int32(12), v13+int32(8))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_simple_heap_delete(m, v19, v41+int32(4))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	F_CatalogTupleUpdate(m, v19, v41+int32(4), v215)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
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
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v225)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)) = uint16(v225)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v231 = F_GUCArrayAdd(m, v225, v230, v15)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l0
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v241 = F_heap_form_tuple(m, v236, v13+int32(16), v13+int32(12))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_CatalogTupleInsert(m, v19, v241)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	goto L8
L70:
	;
	v258 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2964), l0, v258, l1, v258)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	F_systable_endscan(m, v39)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	F_relation_close(m, v19, int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	m.G0 = v13 + int32(128)
	return
}
func F_transformAlterTableStmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v639 int32
	_ = v639
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v759 int32
	_ = v759
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v783 int32
	_ = v783
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
	var v821 int32
	_ = v821
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v883 int32
	_ = v883
	var v896 int32
	_ = v896
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v944 int32
	_ = v944
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	v6 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(208)
	m.G0 = v22
	v25 = F_relation_open(m, l0, v6)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v31 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = l2
	v34 = int32(1)
	v36 = int32(0)
	v39 = F_addRangeTableEntryForRelation(m, v31, v25, v34, v36, v36, v34)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v42 = int32(1)
	F_addNSItemToQuery(m, v31, v39, int32(0), v42, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+136)) = v31
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+119)))
	v50 = base.B2i32(v48 == int32(102))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+156)) = uint8(v50)
	if v48 == int32(102) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v54 = int32(_a_F_transformAlterTableStmt_0)
	goto L8
L7:
	;
	v54 = int32(_a_F_transformAlterTableStmt_1)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+140)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v57 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+160)) = v57
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+157)) = uint8(v59)
	v61 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+152)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v22)+148)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v22)+168)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v22)+176)) = v57
	*(*int64)(unsafe.Add(mBase, uint32(v22)+184)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v61
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+204)) = uint8(v61)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v61
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+196)) = uint8(base.B2i32(v48 == int32(112)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v81 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if int32(0) < v82 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v606 = v61
	v607 = v6
	v610 = v34
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = int32(0)
	F_transformIndexConstraints(m, v22+int32(136))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L138
	}
L12:
	;
	v93 = v61
	v97 = v34
	v102 = v6
	goto L15
L13:
	;
	v586 = v61
	v590 = v34
	goto L14
L14:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v22)+188))
	v606 = v586
	v607 = v597
	v610 = v590
	goto L11
L15:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v102<<(uint(int32(2))%32))))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	switch v109 {
	case 0:
		goto L28
	default:
		goto L18
	case 16:
		goto L27
	case 24:
		goto L26
	case 59, 60:
		goto L23
	case 62:
		goto L25
	case 63:
		goto L24
	}
L16:
	;
	v586 = v563
	v590 = v567
	goto L14
L17:
	;
	v575 = v102 + int32(1)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v575 < v576 {
		v93 = v563
		v97 = v567
		v102 = v575
		goto L15
	} else {
		goto L137
	}
L18:
	;
	v553 = F_lappend(m, v93, v108)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L136
	}
L19:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v22)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v379)+8)) = v548
	goto L18
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L132
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L128
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L124
	}
L23:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v22)+148))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+48))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+119)))
	if v382 != int32(73) {
		goto L102
	} else {
		goto L103
	}
L24:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	if v250 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L25:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	v223 = F_palloc0(m, int32(68))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L59
	}
L26:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+28))
	if v154 != 0 {
		goto L38
	} else {
		goto L39
	}
L27:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v124 == int32(161) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	F_transformColumnDefinition(m, v22+int32(136), v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v115 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+56)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
	v121 = F_lappend(m, v93, v108)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v563 = v121
	v567 = base.B2i32(v117 == v115) & v97
	goto L17
L31:
	;
	F_transformTableConstraint(m, v22+int32(136), v123)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v563 = v93
	v567 = base.B2i32(v132 != int32(9)) & v97
	goto L17
L35:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v141
	F_errmsg_internal(m, int32(_a_F_transformAlterTableStmt_2), v22+int32(16))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_transformAlterTableStmt_3), int32(3631), int32(_a_F_transformAlterTableStmt_4))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v156 = F_transformExpr(m, v31, v154, int32(35))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+131)))
	if v160 != 0 {
		goto L18
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+32)) = v156
	goto L40
L42:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v162 = F_get_attnum(m, l0, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v162 == int32(0) {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	if v162 <= int32(0) {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v168<<(uint(int32(4))%32)+v162*int32(100))+9)))
	if v175 == int32(0) {
		goto L18
	} else {
		goto L46
	}
L46:
	;
	v179 = F_getIdentitySequence(m, v25, v162, int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	v182 = F_typenameTypeId(m, v31, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v185 = F_palloc0(m, int32(16))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(190)
	v189 = F_get_rel_namespace(m, v179)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v191 = F_get_namespace_name(m, v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v193 = F_get_rel_name(m, v179)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v196 = F_makeRangeVar(m, v191, v193, int32(-1))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+4)) = v196
	v200 = F_makeTypeNameFromOid(m, v182)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v203 = F_makeDefElem(m, int32(_a_F_transformAlterTableStmt_5), v200, int32(-1))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v22)+132)) = v203
	v210 = F_list_make1_impl(m, int32(1), v22+int32(44))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v212 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+12)) = uint8(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+8)) = v210
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v22)+184))
	v216 = F_lappend(m, v215, v185)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v216
	v219 = F_lappend(m, v93, v108)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v563 = v219
	v567 = v97
	goto L17
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = int32(90)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+4)) = v227
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+36)) = uint8(v229)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+20)) = v223
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v233 = F_get_attnum(m, l0, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v233 == int32(0) {
		goto L21
	} else {
		goto L61
	}
L61:
	;
	v239 = F_get_atttype(m, l0, v233)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v221)+48))
	v242 = int32(1)
	v244 = int32(0)
	F_generateSerialExtraStmts(m, v22+int32(136), v223, v239, v241, v242, v242, v244, v244)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v248 = F_lappend(m, v93, v108)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v563 = v248
	v567 = v97
	goto L17
L65:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v345 = F_get_attnum(m, l0, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L86
	}
L66:
	;
	v253 = int32(0)
	v334 = v253
	v336 = v253
	goto L65
L67:
	;
	goto L68
L68:
	;
	v255 = int32(0)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	if v258 <= v255 {
		v334 = v255
		v336 = v255
		goto L65
	} else {
		goto L69
	}
L69:
	;
	v267 = v255
	v270 = v255
	v272 = v255
	goto L70
L70:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v250)+12))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v280+v267<<(uint(int32(2))%32))))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+8))
	v286 = int32(_a_F_transformAlterTableStmt_6)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_transformAlterTableStmt[0])))
	if base.B2i32(v289 == int32(0))|base.B2i32(v289 != v292) != 0 {
		v310 = v289
		v311 = v292
		goto L74
	} else {
		goto L75
	}
L71:
	;
	v334 = v319
	v336 = v320
	goto L65
L72:
	;
	v322 = v267 + int32(1)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	if v322 < v323 {
		v267 = v322
		v270 = v319
		v272 = v320
		goto L70
	} else {
		goto L85
	}
L73:
	;
	if v310-v311 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L74:
	;
	goto L73
L75:
	;
	v295 = v285
	v296 = v286
	goto L76
L76:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+1)))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+1)))
	if v300 == int32(0) {
		v310 = v300
		v311 = v299
		goto L74
	} else {
		goto L78
	}
L77:
	;
	v310 = v300
	v311 = v299
	goto L74
L78:
	;
	v303 = int32(1)
	if v300 == v299 {
		v295 = v295 + v303
		v296 = v296 + v303
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v315 = F_lappend(m, v270, v284)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v317 = F_lappend(m, v272, v284)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L84
	}
L83:
	;
	v319 = v315
	v320 = v272
	goto L72
L84:
	;
	v319 = v270
	v320 = v317
	goto L72
L85:
	;
	goto L71
L86:
	;
	if v345 == int32(0) {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	v350 = F_getIdentitySequence(m, v25, v345, int32(1))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	if v350 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v353 = F_palloc0(m, int32(16))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+20)) = v334
	v377 = F_lappend(m, v93, v108)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L98
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = int32(190)
	v357 = F_get_rel_namespace(m, v350)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v359 = F_get_namespace_name(m, v357)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v361 = F_get_rel_name(m, v350)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v364 = F_makeRangeVar(m, v359, v361, int32(-1))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v366 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v353)+12)) = uint16(v366)
	*(*int32)(unsafe.Add(mBase, uint32(v353)+8)) = v336
	*(*int32)(unsafe.Add(mBase, uint32(v353)+4)) = v364
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v22)+184))
	v371 = F_lappend(m, v370, v353)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v371
	goto L91
L98:
	;
	v563 = v377
	v567 = v97
	goto L17
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L121
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L117
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L113
	}
L102:
	;
	switch v382 - int32(105) {
	case 0:
		goto L100
	default:
		goto L99
	case 7:
		goto L105
	case 9:
		goto L101
	}
L103:
	;
	goto L104
L104:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v379)+8))
	if v394 == int32(0) {
		goto L19
	} else {
		goto L108
	}
L105:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v379)+8))
	if v387 == int32(0) {
		goto L19
	} else {
		goto L106
	}
L106:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v22)+136))
	v391 = F_transformPartitionBound(m, v390, v380, v387)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v391
	goto L19
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v380)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v404 + int32(4)
	F_errmsg(m, int32(_a_F_transformAlterTableStmt_7), v22+int32(96))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_transformAlterTableStmt_3), int32(_a_F_transformAlterTableStmt_8), int32(_a_F_transformAlterTableStmt_9))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v380)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v425 + int32(4)
	F_errmsg(m, int32(_a_F_transformAlterTableStmt_10), v22+int32(112))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_transformAlterTableStmt_3), int32(_a_F_transformAlterTableStmt_11), int32(_a_F_transformAlterTableStmt_9))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v380)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v446 + int32(4)
	F_errmsg(m, int32(_a_F_transformAlterTableStmt_12), v22+int32(128))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_transformAlterTableStmt_3), int32(_a_F_transformAlterTableStmt_13), int32(_a_F_transformAlterTableStmt_9))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v380)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v464 + int32(4)
	F_errmsg_internal(m, int32(_a_F_transformAlterTableStmt_14), v22+int32(80))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_transformAlterTableStmt_3), int32(_a_F_transformAlterTableStmt_15), int32(_a_F_transformAlterTableStmt_9))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v485 + int32(4)
	F_errmsg(m, int32(_a_F_transformAlterTableStmt_16), v22+int32(32))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_transformAlterTableStmt_3), int32(3663), int32(_a_F_transformAlterTableStmt_4))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v509
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v508 + int32(4)
	F_errmsg(m, int32(_a_F_transformAlterTableStmt_16), v22+int32(48))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_transformAlterTableStmt_3), int32(3703), int32(_a_F_transformAlterTableStmt_4))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v531 + int32(4)
	F_errmsg(m, int32(_a_F_transformAlterTableStmt_16), v22-int32(-64))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_transformAlterTableStmt_3), int32(3745), int32(_a_F_transformAlterTableStmt_4))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	v563 = v553
	v567 = v97
	goto L17
L137:
	;
	goto L16
L138:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v22)+172))
	v624 = int32(0)
	if base.B2i32(v623 == v624)|base.B2i32(v610 == v624) != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v22)+188))
	if v684 == int32(0) {
		v759 = v606
		goto L145
	} else {
		goto L146
	}
L140:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	if v629 <= int32(0) {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v639 = int32(0)
	goto L142
L142:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v623)+12))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v652+v639<<(uint(int32(2))%32))))
	v657 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v656)+15)) = uint8(v657)
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v656)+16)) = uint8(v659)
	v662 = v639 + v657
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	if v662 < v663 {
		v639 = v662
		goto L142
	} else {
		goto L144
	}
L143:
	;
	goto L139
L144:
	;
	goto L143
L145:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v22)+164))
	if v770 == int32(0) {
		v821 = v759
		goto L163
	} else {
		goto L164
	}
L146:
	;
	v687 = int32(0)
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v684)+4))
	if v688 <= v687 {
		v759 = v606
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v699 = v606
	v701 = v687
	goto L148
L148:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v684)+12))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v710+v701<<(uint(int32(2))%32))))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)))
	if v715 == int32(204) {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L160
	}
L150:
	;
	v718 = F_transformIndexStmt(m, l0, v714, l2)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	goto L149
L153:
	;
	v721 = F_palloc0(m, int32(32))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v721))) = int32(147)
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v718)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v721)+20)) = v718
	if v725 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v729 = int32(21)
	goto L157
L156:
	;
	v729 = int32(14)
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v721)+4)) = v729
	v731 = F_lappend(m, v699, v721)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v734 = v701 + int32(1)
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v684)+4))
	if v734 < v735 {
		v699 = v731
		v701 = v734
		goto L148
	} else {
		goto L159
	}
L159:
	;
	v759 = v731
	goto L145
L160:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v714)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v741
	F_errmsg_internal(m, int32(_a_F_transformAlterTableStmt_17), v22)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_transformAlterTableStmt_3), int32(3838), int32(_a_F_transformAlterTableStmt_4))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v22)+168))
	if v832 == int32(0) {
		v883 = v821
		goto L171
	} else {
		goto L172
	}
L164:
	;
	v773 = int32(0)
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v770)+4))
	if v774 <= v773 {
		v821 = v759
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v783 = v773
	v785 = v759
	goto L166
L166:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v770)+12))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v796+v783<<(uint(int32(2))%32))))
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
	v821 = v807
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
	v810 = v783 + int32(1)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v770)+4))
	if v810 < v811 {
		v783 = v810
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
	if v623 == int32(0) {
		v944 = v883
		goto L179
	} else {
		goto L180
	}
L172:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v832)+4))
	if v835 <= int32(0) {
		v883 = v821
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v845 = int32(0)
	v847 = v821
	goto L174
L174:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v832)+12))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v858+v845<<(uint(int32(2))%32))))
	v864 = F_palloc0(m, int32(32))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L176
	}
L175:
	;
	v883 = v869
	goto L171
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v864)+20)) = v862
	*(*int64)(unsafe.Add(mBase, uint32(v864))) = int64(68719476883)
	v869 = F_lappend(m, v847, v864)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v872 = v845 + int32(1)
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v832)+4))
	if v872 < v873 {
		v845 = v872
		v847 = v869
		goto L174
	} else {
		goto L178
	}
L178:
	;
	goto L175
L179:
	;
	F_relation_close(m, v25, int32(0))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L187
	}
L180:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	if v896 <= int32(0) {
		v944 = v883
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v906 = int32(0)
	v908 = v883
	goto L182
L182:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v623)+12))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v919+v906<<(uint(int32(2))%32))))
	v925 = F_palloc0(m, int32(32))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L184
	}
L183:
	;
	v944 = v930
	goto L179
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v925)+20)) = v923
	*(*int64)(unsafe.Add(mBase, uint32(v925))) = int64(68719476883)
	v930 = F_lappend(m, v908, v925)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v933 = v906 + int32(1)
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	if v933 < v934 {
		v906 = v933
		v908 = v930
		goto L182
	} else {
		goto L186
	}
L186:
	;
	goto L183
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v944
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v22)+184))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v959
	v962 = F_list_concat(m, int32(0), v607)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v962
	m.G0 = v22 + int32(208)
	return l1
}
