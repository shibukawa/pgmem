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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int64
	_ = v89
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
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
					v35 = v12 + int32(56)
					v37 = F_strncpy(m, v35, l2, int32(64))
					mBase = m.M
					v38 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v37)+63)) = uint8(v38)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+140)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v12)+136)) = v35
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_CreateOpFamily[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v43
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
					v50 = F_heap_form_tuple(m, v45, v12+int32(128), v12+int32(120))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_CatalogTupleInsert(m, v16, v50)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							F_pfree(m, v50)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								v56 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v30
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2753)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(2601)
								v67 = v12 + int32(44)
								F_recordDependencyOn(m, l0, v67, int32(97))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(2615)
									F_recordDependencyOn(m, l0, v67, int32(110))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, _c_F_CreateOpFamily[0]))
										F_recordDependencyOnOwner(m, int32(2753), v30, v81)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											F_recordDependencyOnCurrentExtension(m, l0, int32(0))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v87
												v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v89
												v92 = *(*int64)(unsafe.Add(mBase, _c_F_CreateOpFamily[1]))
												*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v92
												v95 = *(*int32)(unsafe.Add(mBase, _c_F_CreateOpFamily[2]))
												*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v95
												F_EventTriggerCollectSimpleCommand(m, v12+int32(32), v12+int32(16), l1)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return
												} else {
													v104 = *(*int32)(unsafe.Add(mBase, _c_F_CreateOpFamily[3]))
													if v104 != 0 {
														v106 = int32(0)
														F_RunObjectPostCreateHook(m, int32(2753), v30, v106, v106)
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return
														} else {
															F_relation_close(m, v16, int32(3))
															mBase = m.M
															v112 = m.ExcPending
															if v112 != 0 {
																return
															} else {
																m.G0 = v12 + int32(160)
																return
															}
														}
													} else {
														F_relation_close(m, v16, int32(3))
														mBase = m.M
														v112 = m.ExcPending
														if v112 != 0 {
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
				v119 = m.ExcPending
				if v119 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_CreateOpFamily_0))
					mBase = m.M
					v122 = m.ExcPending
					if v122 != 0 {
						return
					} else {
						v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v123
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = l2
						F_errmsg(m, int32(_a_F_CreateOpFamily_1), v12)
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CreateOpFamily_2), int32(268), int32(_a_F_CreateOpFamily_3))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
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
	v147 = m.ExcPending
	if v147 != 0 {
		goto L8
	} else {
		goto L45
	}
L8:
	;
	return int32(0)
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	if v23 <= int32(0) {
		v139 = v4
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
	v139 = v129
	goto L7
L13:
	;
	v133 = v29 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	if v133 < v134 {
		v29 = v133
		v32 = v129
		goto L11
	} else {
		goto L44
	}
L14:
	;
	v129 = v32
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
	if v118 != int32(2) {
		v129 = v116
		goto L13
	} else {
		goto L43
	}
L19:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v92 = F_SearchSysCache4(m, int32(5), v90, v86, v86, int32(1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L33
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
		goto L31
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
	if base.B2i32(l2 == v70)|base.B2i32(v63 == v70) != 0 {
		v116 = base.B2i32(v63 != v70) | v32
		v118 = v69
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v78 != v79 {
		v86 = v78
		goto L19
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v63
	v139 = int32(1)
	goto L7
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v86 = v85
	goto L19
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v104
	if v104 != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	if v92 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v104 = int32(0)
	goto L32
L35:
	;
	goto L36
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+22)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97+v98)+20))
	F_ReleaseCatCache(m, v92)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v104 = v100
	goto L32
L38:
	;
	v108 = int32(2)
	goto L40
L39:
	;
	v108 = int32(4)
	goto L40
L40:
	;
	v109 = int32(0)
	v111 = base.B2i32(v104 != v109) | v32
	if l1 == v109 {
		v116 = v111
		v118 = v108
		goto L18
	} else {
		goto L41
	}
L41:
	;
	if v104 == int32(0) {
		goto L17
	} else {
		goto L42
	}
L42:
	;
	v116 = v111
	v118 = v108
	goto L18
L43:
	;
	v139 = v116
	goto L7
L44:
	;
	goto L12
L45:
	;
	return v139 & int32(1)
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
				F_errmsg_internal(m, int32(_a_F_op_strict_0), v7)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_op_strict_1), int32(1622), int32(_a_F_op_strict_2))
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
						F_errmsg_internal(m, int32(_a_F_op_strict_0), v7)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_op_strict_1), int32(1622), int32(_a_F_op_strict_2))
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
								F_errmsg_internal(m, int32(_a_F_op_strict_3), v7+int32(16))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_op_strict_1), int32(1908), int32(_a_F_op_strict_4))
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
