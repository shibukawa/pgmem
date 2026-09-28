package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyTriggerDesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v2 = int32(0)
	if l0 == v2 {
		v135 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v135
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= int32(0) {
		v135 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = F_palloc(m, int32(32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v20
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v22
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = F_palloc(m, v26*int32(60))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = v31 * int32(60)
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryCopy(m, v29, v34, v33)
	goto L9
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v29
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v37 <= int32(0) {
		v135 = v14
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v41 = v29
	v46 = v2
	goto L11
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v48 = F_pstrdup(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	v135 = v14
	goto L1
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v48
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+36)))
	if int32(0) < v51 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = F_palloc(m, v51<<(uint(int32(1))%32)&int32(_a_F_CopyTriggerDesc_0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+34)))
	if int32(0) < v68 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+36)))
	v62 = v60 << (uint(int32(1)) % 32)
	if v62 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v41)+40))
	base.MemoryCopy(m, v58, v63, v62)
	goto L20
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+40)) = v58
	goto L16
L21:
	;
	v73 = F_palloc(m, v68<<(uint(int32(2))%32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	if v114 != 0 {
		goto L32
	} else {
		goto L33
	}
L24:
	;
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+34)))
	if int32(0) < v75 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v81 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+44)) = v73
	goto L23
L28:
	;
	v87 = v81 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v41)+44))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89+v87)))
	v92 = F_pstrdup(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73+v87))) = v92
	v96 = v81 + int32(1)
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+34)))
	if v96 < v97 {
		v81 = v96
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v115 = F_pstrdup(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v41)+52))
	if v118 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+48)) = v115
	goto L34
L36:
	;
	v119 = F_pstrdup(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	if v122 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+52)) = v119
	goto L38
L40:
	;
	v123 = F_pstrdup(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v129 = v46 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v129 < v130 {
		v41 = v41 + int32(60)
		v46 = v129
		goto L11
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+56)) = v123
	goto L42
L44:
	;
	goto L12
}
func F_TriggerSetParentTrigger(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	v15 = v12 + int32(48)
	F_ScanKeyInit(m, v15, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		v23 = int32(1)
		v26 = F_systable_beginscan(m, l0, int32(2702), v23, int32(0), v23, v15)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = F_systable_getnext(m, v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				if v28 != 0 {
					v30 = F_heap_copytuple(m, v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
						v34 = v32 + v33
						if l2 != 0 {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
							if v35 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
									F_errmsg_internal(m, int32(_a_F_TriggerSetParentTrigger_0), v12+int32(16))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_TriggerSetParentTrigger_1), int32(1270), int32(_a_F_TriggerSetParentTrigger_2))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = l2
								F_CatalogTupleUpdate(m, l0, v28+int32(4), v30)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									v41 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v41
									*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = l1
									v44 = int32(2620)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v44
									*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v41
									*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v44
									v52 = v12 + int32(36)
									v54 = v12 + int32(24)
									F_recordDependencyOn(m, v52, v54, int32(80))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l3
										*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(1259)
										F_recordDependencyOn(m, v52, v54, int32(83))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											F_pfree(m, v30)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												F_systable_endscan(m, v26)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													m.G0 = v12 + int32(112)
													return
												}
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = int32(0)
							F_CatalogTupleUpdate(m, l0, v28+int32(4), v30)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								v72 = int32(2620)
								v75 = F_deleteDependencyRecordsForClass(m, v72, l1, v72, int32(80))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									v80 = F_deleteDependencyRecordsForClass(m, int32(2620), l1, int32(1259), int32(83))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										F_pfree(m, v30)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return
										} else {
											F_systable_endscan(m, v26)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												m.G0 = v12 + int32(112)
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
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
						F_errmsg_internal(m, int32(_a_F_TriggerSetParentTrigger_3), v12)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_TriggerSetParentTrigger_1), int32(1262), int32(_a_F_TriggerSetParentTrigger_2))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
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
