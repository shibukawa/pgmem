package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateOpFamily(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	v12 = m.G0
	v14 = v12 - int32(176)
	m.G0 = v14
	v18 = F_table_open(m, int32(2753), int32(3))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v21 = base.I64_extend_i32_u(l4)
		v23 = base.I64_extend_i32_u(l3)
		v25 = F_SearchSysCacheExists(m, int32(41), v21, base.I64_extend_i32_u(l2), v23, int64(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			if v25 == int32(0) {
				v29 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+124)) = uint8(v29)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+120)) = v29
				v35 = F_GetNewOidWithIndex(m, v18, int32(2755), int32(1))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v14)+136)) = v21
					*(*int64)(unsafe.Add(mBase, uint32(v14)+128)) = base.I64_extend_i32_u(v35)
					v41 = v14 + int32(56)
					v43 = F_strncpy(m, v41, l2, int32(64))
					mBase = m.M
					v44 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v43)+63)) = uint8(v44)
					*(*int64)(unsafe.Add(mBase, uint32(v14)+152)) = v23
					*(*int64)(unsafe.Add(mBase, uint32(v14)+144)) = base.I64_extend_i32_u(v41)
					v50 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_CreateOpFamily[0])))
					*(*int64)(unsafe.Add(mBase, uint32(v14)+160)) = v50
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
					v57 = F_heap_form_tuple(m, v52, v14+int32(128), v14+int32(120))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						F_CatalogTupleInsert(m, v18, v57)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							F_pfree(m, v57)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v63 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2753)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = int32(2601)
								v74 = v14 + int32(44)
								F_recordDependencyOn(m, l0, v74, int32(97))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = int32(2615)
									F_recordDependencyOn(m, l0, v74, int32(110))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, _c_F_CreateOpFamily[0]))
										F_recordDependencyOnOwner(m, int32(2753), v35, v88)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_recordDependencyOnCurrentExtension(m, l0, int32(0))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v94
												v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v96
												v99 = *(*int64)(unsafe.Add(mBase, _c_F_CreateOpFamily[1]))
												*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v99
												v102 = *(*int32)(unsafe.Add(mBase, _c_F_CreateOpFamily[2]))
												*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v102
												F_EventTriggerCollectSimpleCommand(m, v14+int32(32), v14+int32(16), l1)
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													v111 = *(*int32)(unsafe.Add(mBase, _c_F_CreateOpFamily[3]))
													if v111 != 0 {
														v113 = int32(0)
														F_RunObjectPostCreateHook(m, int32(2753), v35, v113, v113)
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															F_relation_close(m, v18, int32(3))
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return
															} else {
																m.G0 = v14 + int32(176)
																return
															}
														}
													} else {
														F_relation_close(m, v18, int32(3))
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return
														} else {
															m.G0 = v14 + int32(176)
															return
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_CreateOpFamily_0))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return
					} else {
						v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v130
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = l2
						F_errmsg(m, int32(_a_F_CreateOpFamily_1), v14)
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CreateOpFamily_2), int32(268), int32(_a_F_CreateOpFamily_3))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
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
func F_get_op_hash_functions(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v52 int32
	_ = v52
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	v4 = int32(0)
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v19 = int64(0)
	v21 = F_SearchSysCacheList(m, int32(3), int32(1), base.I64_extend_i32_u(l0), v19, v19)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_ReleaseCatCacheList(m, v21)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L8
	} else {
		goto L45
	}
L8:
	;
	return int32(0)
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v25 <= int32(0) {
		v148 = v4
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v31 = int32(0)
	v34 = v4
	goto L11
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)+v31<<(uint(int32(2))%32))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+72))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+22)))
	v48 = v46 + v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v49 != int32(405) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v148 = v137
	goto L7
L13:
	;
	v142 = v31 + int32(1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v142 < v143 {
		v31 = v142
		v34 = v137
		goto L11
	} else {
		goto L44
	}
L14:
	;
	v137 = v34
	goto L13
L15:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+16)))
	if v52 != int32(1) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if l1 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L14
L18:
	;
	if v123 != int32(2) {
		v137 = v121
		goto L13
	} else {
		goto L43
	}
L19:
	;
	v94 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+4)))
	v95 = base.I64_extend_i32_u(v89)
	v97 = F_SearchSysCache4(m, int32(5), v94, v95, v95, int64(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L33
	}
L20:
	;
	v56 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+4)))
	v57 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v48)+8)))
	v59 = F_SearchSysCache4(m, int32(5), v56, v57, v57, int64(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if l2 == int32(0) {
		goto L14
	} else {
		goto L31
	}
L23:
	;
	if v59 == int32(0) {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+22)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63+v64)+20))
	F_ReleaseCatCache(m, v59)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v66
	if v66 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v72 = int32(2)
	goto L28
L27:
	;
	v72 = int32(4)
	goto L28
L28:
	;
	v73 = int32(0)
	if base.B2i32(l2 == v73)|base.B2i32(v66 == v73) != 0 {
		v121 = base.B2i32(v66 != v73) | v34
		v123 = v72
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if v81 != v82 {
		v89 = v81
		goto L19
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v66
	v148 = int32(1)
	goto L7
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v89 = v88
	goto L19
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v109
	if v109 != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	if v97 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v109 = int32(0)
	goto L32
L35:
	;
	goto L36
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+22)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102+v103)+20))
	F_ReleaseCatCache(m, v97)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v109 = v105
	goto L32
L38:
	;
	v113 = int32(2)
	goto L40
L39:
	;
	v113 = int32(4)
	goto L40
L40:
	;
	v114 = int32(0)
	v116 = base.B2i32(v109 != v114) | v34
	if l1 == v114 {
		v121 = v116
		v123 = v113
		goto L18
	} else {
		goto L41
	}
L41:
	;
	if v109 == int32(0) {
		goto L17
	} else {
		goto L42
	}
L42:
	;
	v121 = v116
	v123 = v113
	goto L18
L43:
	;
	v148 = v121
	goto L7
L44:
	;
	goto L12
L45:
	;
	return v148 & int32(1)
}
func F_op_strict(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v11 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(_a_F_op_strict_0), v7)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_op_strict_1), int32(1794), int32(_a_F_op_strict_2))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18)+100))
			F_ReleaseCatCache(m, v11)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				if v20 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg_internal(m, int32(_a_F_op_strict_0), v7)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_op_strict_1), int32(1794), int32(_a_F_op_strict_2))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v27 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(v20))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v20
								F_errmsg_internal(m, int32(_a_F_op_strict_3), v7+int32(16))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_op_strict_1), int32(2080), int32(_a_F_op_strict_4))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
							v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v32)+99)))
							F_ReleaseCatCache(m, v27)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(32)
								return v34
							}
						}
					}
				}
			}
		}
	}
}
