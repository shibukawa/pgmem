package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateOpFamily(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v97 int32
	_ = v97
	var v101 int64
	_ = v101
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v10 = m.G0
	v12 = v10 - int32(160)
	m.G0 = v12
	v16 = F_table_open(m, int32(2753), int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v20 = F_SearchSysCacheExists(m, int32(41), l4, l2, l3, int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 == int32(0) {
				v24 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v12)+124)) = uint8(v24)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = v24
				v30 = F_GetNewOidWithIndex(m, v16, int32(2755), int32(1))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+132)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v30
					v37 = F_strncpy(m, v12+int32(56), l2, int32(64))
					mBase = m.M
					v38 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v37)+63)) = uint8(v38)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+140)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v12)+136)) = v12 + int32(56)
					v45 = *(*int32)(unsafe.Add(mBase, _consts[31]))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v45
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
					v52 = F_heap_form_tuple(m, v47, v12+int32(128), v12+int32(120))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_CatalogTupleInsert(m, v16, v52)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							F_pfree(m, v52)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								v59 = l0 + int32(8)
								v60 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v59))) = v60
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v30
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2753)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v60
								*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(2601)
								F_recordDependencyOn(m, l0, v12+int32(44), int32(97))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(2615)
									F_recordDependencyOn(m, l0, v12+int32(44), int32(110))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, _consts[31]))
										F_recordDependencyOnOwner(m, int32(2753), v30, v87)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											F_recordDependencyOnCurrentExtension(m, l0, int32(0))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												v93 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
												*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v93
												v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												v97 = *(*int32)(unsafe.Add(mBase, _consts[351]))
												*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v97
												*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v95
												v101 = *(*int64)(unsafe.Add(mBase, _consts[350]))
												*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v101
												F_EventTriggerCollectSimpleCommand(m, v12+int32(32), v12+int32(16), l1)
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return
												} else {
													v110 = *(*int32)(unsafe.Add(mBase, _consts[297]))
													if v110 != 0 {
														v112 = int32(0)
														F_RunObjectPostCreateHook(m, int32(2753), v30, v112, v112)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return
														} else {
															F_sequence_close(m, v16, int32(3))
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return
															} else {
																m.G0 = v12 + int32(160)
																return
															}
														}
													} else {
														F_sequence_close(m, v16, int32(3))
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return
														} else {
															m.G0 = v12 + int32(160)
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
				v125 = m.ExcPending
				if v125 != 0 {
					return
				} else {
					F_errcode(m, int32(290948))
					mBase = m.M
					v128 = m.ExcPending
					if v128 != 0 {
						return
					} else {
						v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v129
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = l2
						F_errmsg(m, int32(117050), v12)
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return
						} else {
							F_errfinish(m, int32(493999), int32(268), int32(19990))
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
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
	v17 = int32(0)
	v19 = F_SearchSysCacheList(m, int32(3), int32(1), l0, v17, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_ReleaseCatCacheList(m, v19)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L8
	} else {
		goto L46
	}
L8:
	;
	return int32(0)
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	if v23 <= int32(0) {
		v141 = v4
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v29 = int32(0)
	v32 = v4
	goto L11
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(48)+v29<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+56))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v45 = v43 + v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	if v46 != int32(405) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v141 = v130
	goto L7
L13:
	;
	v135 = v29 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	if v135 < v136 {
		v29 = v135
		v32 = v130
		goto L11
	} else {
		goto L45
	}
L14:
	;
	v130 = v32
	goto L13
L15:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+16)))
	if v49 != int32(1) {
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
	if v117 != int32(2) {
		v130 = v114
		goto L13
	} else {
		goto L44
	}
L19:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v91 = F_SearchSysCache4(m, int32(5), v89, v85, v85, int32(1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L34
	}
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v56 = F_SearchSysCache4(m, int32(5), v53, v54, v54, int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
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
		goto L32
	}
L23:
	;
	if v56 == int32(0) {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60+v61)+20))
	F_ReleaseCatCache(m, v56)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v63
	if v63 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v69 = int32(2)
	goto L28
L27:
	;
	v69 = int32(4)
	goto L28
L28:
	;
	v70 = int32(0)
	v72 = base.B2i32(v63 != v70) | v32
	if l2 == v70 {
		v114 = v72
		v117 = v69
		goto L18
	} else {
		goto L29
	}
L29:
	;
	if v63 == int32(0) {
		v114 = v72
		v117 = v69
		goto L18
	} else {
		goto L30
	}
L30:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v77 != v78 {
		v85 = v77
		goto L19
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v63
	v141 = int32(1)
	goto L7
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v85 = v84
	goto L19
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v102
	if v102 != 0 {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	if v91 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v102 = int32(0)
	goto L33
L36:
	;
	goto L37
L37:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+22)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96+v97)+20))
	F_ReleaseCatCache(m, v91)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v102 = v99
	goto L33
L39:
	;
	v106 = int32(2)
	goto L41
L40:
	;
	v106 = int32(4)
	goto L41
L41:
	;
	v107 = int32(0)
	v109 = base.B2i32(v102 != v107) | v32
	if l1 == v107 {
		v114 = v109
		v117 = v106
		goto L18
	} else {
		goto L42
	}
L42:
	;
	if v102 == int32(0) {
		goto L17
	} else {
		goto L43
	}
L43:
	;
	v114 = v109
	v117 = v106
	goto L18
L44:
	;
	v141 = v114
	goto L7
L45:
	;
	goto L12
L46:
	;
	return v141 & int32(1)
}
func F_op_strict(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = F_SearchSysCache1(m, int32(40), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(68755), v7)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(498911), int32(1622), int32(109131))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16+v17)+100))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg_internal(m, int32(68755), v7)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(498911), int32(1622), int32(109131))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v25 = F_SearchSysCache1(m, int32(47), v19)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v19
								F_errmsg_internal(m, int32(44676), v7+int32(16))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(498911), int32(1908), int32(109141))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
							v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
							v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v30)+99)))
							F_ReleaseCatCache(m, v25)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(32)
								return v32
							}
						}
					}
				}
			}
		}
	}
}
