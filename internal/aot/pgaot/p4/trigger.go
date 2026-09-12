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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v34 = v32 * int32(60)
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v38 <= int32(0) {
		v135 = v14
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v35 = F__emscripten_memcpy_bulkmem(m, v29, v31, v34)
	mBase = m.M
	v36 = v35
	goto L10
L9:
	;
	v36 = v29
	goto L10
L10:
	;
	goto L7
L11:
	;
	v42 = v29
	v45 = v2
	goto L12
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	v49 = F_pstrdup(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L14
	}
L13:
	;
	v135 = v14
	goto L1
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v49
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+36)))
	if int32(0) < v52 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v59 = F_palloc(m, v52<<(uint(int32(1))%32)&int32(65534))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+34)))
	if int32(0) < v68 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42)+40))
	v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+36)))
	v64 = v62 << (uint(int32(1)) % 32)
	if v64 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+40)) = v66
	goto L17
L20:
	;
	v65 = F__emscripten_memcpy_bulkmem(m, v59, v61, v64)
	mBase = m.M
	v66 = v65
	goto L22
L21:
	;
	v66 = v59
	goto L22
L22:
	;
	goto L19
L23:
	;
	v73 = F_palloc(m, v68<<(uint(int32(2))%32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	if v114 != 0 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+34)))
	if int32(0) < v75 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v81 = int32(0)
	goto L30
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+44)) = v73
	goto L25
L30:
	;
	v87 = v81 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v42)+44))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89+v87)))
	v92 = F_pstrdup(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73+v87))) = v92
	v96 = v81 + int32(1)
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+34)))
	if v96 < v97 {
		v81 = v96
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v115 = F_pstrdup(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
	if v118 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = v115
	goto L36
L38:
	;
	v119 = F_pstrdup(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v42)+56))
	if v122 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+52)) = v119
	goto L40
L42:
	;
	v123 = F_pstrdup(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v129 = v45 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v129 < v130 {
		v42 = v42 + int32(60)
		v45 = v129
		goto L12
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+56)) = v123
	goto L44
L46:
	;
	goto L13
}
func F_TriggerSetParentTrigger(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	F_ScanKeyInit(m, v12+int32(48), int32(1), int32(3), int32(184), l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v22 = int32(1)
		v27 = F_systable_beginscan(m, l0, int32(2702), v22, int32(0), v22, v12+int32(48))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v29 = F_systable_getnext(m, v27)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				if v29 != 0 {
					v31 = F_heap_copytuple(m, v29)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
						v35 = v33 + v34
						if l2 != 0 {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
							if v36 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
									F_errmsg_internal(m, int32(219644), v12+int32(16))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return
									} else {
										F_errfinish(m, int32(484380), int32(1255), int32(220269))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = l2
								F_CatalogTupleUpdate(m, l0, v29+int32(4), v31)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									v42 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v42
									*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = l1
									v45 = int32(2620)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v45
									*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v42
									*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v45
									F_recordDependencyOn(m, v12+int32(36), v12+int32(24), int32(80))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l3
										*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(1259)
										F_recordDependencyOn(m, v12+int32(36), v12+int32(24), int32(83))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											F_pfree(m, v31)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return
											} else {
												F_systable_endscan(m, v27)
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return
												} else {
													m.G0 = v12 + int32(96)
													return
												}
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = int32(0)
							F_CatalogTupleUpdate(m, l0, v29+int32(4), v31)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								v77 = int32(2620)
								v80 = F_deleteDependencyRecordsForClass(m, v77, l1, v77, int32(80))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v85 = F_deleteDependencyRecordsForClass(m, int32(2620), l1, int32(1259), int32(83))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										F_pfree(m, v31)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_systable_endscan(m, v27)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												m.G0 = v12 + int32(96)
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
					v97 = m.ExcPending
					if v97 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
						F_errmsg_internal(m, int32(42915), v12)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							F_errfinish(m, int32(484380), int32(1247), int32(220269))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
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
