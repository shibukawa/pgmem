package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecInitRangeTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l1
	if l1 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v9 = v8
	} else {
		v9 = int32(0)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v14 = F_palloc0(m, v9<<(uint(int32(2))%32))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v16
		return
	}
}
func F_RangeVarCallbackForRenameRule(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v17 = v15 + v16
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
			v20 = v18 - int32(112)
			if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v20))|base.B2i32(int32(1)<<(uint(v20)%32)&int32(69) == int32(0)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v83
						F_errmsg(m, int32(_a_F_RangeVarCallbackForRenameRule_0), v9)
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							v88 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+119)))
							F_errdetail_relkind_not_supported(m, v88)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_RangeVarCallbackForRenameRule_1), int32(760), int32(_a_F_RangeVarCallbackForRenameRule_2))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
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
			} else {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarCallbackForRenameRule[0])))
				if v31 == int32(0) {
					v35 = int32(1)
					if base.Ui32(l1) < base.Ui32(int32(_a_F_RangeVarCallbackForRenameRule_3)) {
						v43 = v35
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
						if v38 == int32(99) {
							v43 = v35
						} else {
							v41 = F_isTempToastNamespace(m, v38)
							mBase = m.M
							v43 = v41
						}
					}
					if v43 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v103
								F_errmsg(m, int32(_a_F_RangeVarCallbackForRenameRule_4), v9+int32(16))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_RangeVarCallbackForRenameRule_1), int32(766), int32(_a_F_RangeVarCallbackForRenameRule_2))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForRenameRule[1]))
						v47 = F_object_ownercheck(m, int32(1259), l1, v46)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							if v47 == int32(0) {
								v52 = F_get_rel_relkind(m, l1)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									switch v52 - int32(73) {
									case 0, 32:
										v63 = int32(20)
										v65 = v63
									default:
										v63 = int32(42)
										v65 = v63
									case 10:
										v65 = int32(38)
									case 29:
										v65 = int32(18)
									case 36:
										v65 = int32(23)
									case 45:
										v65 = int32(52)
									}
									v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									F_aclcheck_error(m, int32(2), v65, v66)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v13)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v13)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForRenameRule[1]))
					v47 = F_object_ownercheck(m, int32(1259), l1, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						if v47 == int32(0) {
							v52 = F_get_rel_relkind(m, l1)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								switch v52 - int32(73) {
								case 0, 32:
									v63 = int32(20)
									v65 = v63
								default:
									v63 = int32(42)
									v65 = v63
								case 10:
									v65 = int32(38)
								case 29:
									v65 = int32(18)
								case 36:
									v65 = int32(23)
								case 45:
									v65 = int32(52)
								}
								v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								F_aclcheck_error(m, int32(2), v65, v66)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v13)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						} else {
							F_ReleaseCatCache(m, v13)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			}
		} else {
			m.G0 = v9 + int32(32)
			return
		}
	}
}
func F_RangeVarCallbackOwnsRelation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l1 != 0 {
		v11 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l1))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			if v11 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg_internal(m, int32(_a_F_RangeVarCallbackOwnsRelation_0), v7)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_RangeVarCallbackOwnsRelation_1), int32(_a_F_RangeVarCallbackOwnsRelation_2), int32(_a_F_RangeVarCallbackOwnsRelation_3))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackOwnsRelation[0]))
				v18 = F_object_ownercheck(m, int32(1259), l1, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					if v18 == int32(0) {
						v23 = F_get_rel_relkind(m, l1)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							switch v23 - int32(73) {
							case 0, 32:
								v34 = int32(20)
								v36 = v34
							default:
								v34 = int32(42)
								v36 = v34
							case 10:
								v36 = int32(38)
							case 29:
								v36 = int32(18)
							case 36:
								v36 = int32(23)
							case 45:
								v36 = int32(52)
							}
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							F_aclcheck_error(m, int32(2), v36, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarCallbackOwnsRelation[1])))
								if v41 == int32(0) {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
									v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+22)))
									v48 = int32(1)
									if base.Ui32(l1) < base.Ui32(int32(_a_F_RangeVarCallbackOwnsRelation_4)) {
										v56 = v48
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(v44+v45)+68))
										if v51 == int32(99) {
											v56 = v48
										} else {
											v54 = F_isTempToastNamespace(m, v51)
											mBase = m.M
											v56 = v54
										}
									}
									if v56 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											F_errcode(m, int32(16797828))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v85
												F_errmsg(m, int32(_a_F_RangeVarCallbackOwnsRelation_5), v7+int32(16))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_RangeVarCallbackOwnsRelation_1), int32(_a_F_RangeVarCallbackOwnsRelation_6), int32(_a_F_RangeVarCallbackOwnsRelation_3))
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										F_ReleaseCatCache(m, v11)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											m.G0 = v7 + int32(32)
											return
										}
									}
								} else {
									F_ReleaseCatCache(m, v11)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								}
							}
						}
					} else {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarCallbackOwnsRelation[1])))
						if v41 == int32(0) {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+22)))
							v48 = int32(1)
							if base.Ui32(l1) < base.Ui32(int32(_a_F_RangeVarCallbackOwnsRelation_4)) {
								v56 = v48
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v44+v45)+68))
								if v51 == int32(99) {
									v56 = v48
								} else {
									v54 = F_isTempToastNamespace(m, v51)
									mBase = m.M
									v56 = v54
								}
							}
							if v56 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return
									} else {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v85
										F_errmsg(m, int32(_a_F_RangeVarCallbackOwnsRelation_5), v7+int32(16))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_RangeVarCallbackOwnsRelation_1), int32(_a_F_RangeVarCallbackOwnsRelation_6), int32(_a_F_RangeVarCallbackOwnsRelation_3))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v11)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									m.G0 = v7 + int32(32)
									return
								}
							}
						} else {
							F_ReleaseCatCache(m, v11)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								m.G0 = v7 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v7 + int32(32)
		return
	}
}
func F_RangeVarGetCreationNamespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v7 + int32(32)
	return v148
L2:
	;
	F_AccessTempTableNamespace(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L8
	} else {
		goto L47
	}
L3:
	;
	v143 = int32(0)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L43
	}
L5:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetCreationNamespace[0]))
	v12 = F_get_database_name(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v42 != 0 {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	return int32(0)
L9:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if base.B2i32(v18 == int32(0))|base.B2i32(v18 != v21) != 0 {
		v39 = v18
		v40 = v21
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v39-v40 != 0 {
		goto L4
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v24 = v9
	v25 = v12
	goto L13
L13:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 == int32(0) {
		v39 = v29
		v40 = v28
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v39 = v29
	v40 = v28
	goto L11
L15:
	;
	v32 = int32(1)
	if v29 == v28 {
		v24 = v24 + v32
		v25 = v25 + v32
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	goto L7
L18:
	;
	v43 = int32(_a_F_RangeVarGetCreationNamespace_0)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarGetCreationNamespace[1])))
	if base.B2i32(v46 == int32(0))|base.B2i32(v46 != v49) != 0 {
		v67 = v46
		v68 = v49
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v95 == int32(116) {
		goto L3
	} else {
		goto L35
	}
L21:
	;
	if v67-v68 == int32(0) {
		goto L3
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v52 = v42
	v53 = v43
	goto L24
L24:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v57 == int32(0) {
		v67 = v57
		v68 = v56
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v67 = v57
	v68 = v56
	goto L22
L26:
	;
	v60 = int32(1)
	if v57 == v56 {
		v52 = v52 + v60
		v53 = v53 + v60
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v74 = int64(0)
	v77 = F_GetSysCacheOid(m, int32(37), base.I64_extend_i32_u(v42), v74, v74, v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	if v77 != 0 {
		v148 = v77
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v42
	F_errmsg(m, int32(_a_F_RangeVarGetCreationNamespace_1), v7)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_RangeVarGetCreationNamespace_2), int32(3616), int32(_a_F_RangeVarGetCreationNamespace_3))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarGetCreationNamespace[2])))
	if v102 != 0 {
		v143 = int32(1)
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetCreationNamespace[3]))
	if v104 != 0 {
		v148 = v104
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(_a_F_RangeVarGetCreationNamespace_4), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_RangeVarGetCreationNamespace_2), int32(705), int32(_a_F_RangeVarGetCreationNamespace_5))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v128
	F_errmsg(m, int32(_a_F_RangeVarGetCreationNamespace_6), v7+int32(16))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_RangeVarGetCreationNamespace_2), int32(669), int32(_a_F_RangeVarGetCreationNamespace_5))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetCreationNamespace[4]))
	v148 = v147
	goto L1
}
func F_make_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = F_range_serialize(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if l4 == int32(0) {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
			if v23 == int32(0) {
				v68 = v12
				m.G0 = v10 + int32(48)
				return v68
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v30 = int32(1)
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v26)>>(uint(int32(2))%32))-v30))))
				if v32&v30 != 0 {
					v68 = v12
					m.G0 = v10 + int32(48)
					return v68
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l4
					v38 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v38)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)) = uint8(v38)
					v42 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v10)+26)) = uint16(v42)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_u(v12)
					v47 = l0 + int32(240)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v47
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
					v52 = m.T0[v51].(func(*base.Module, int32) int64)(m, v10+int32(8))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if l4 == int32(0) {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
							if v61 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v77
									F_errmsg_internal(m, int32(_a_F_make_range_0), v10)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_make_range_1), int32(2218), int32(_a_F_make_range_2))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v65 = F_pg_detoast_datum(m, base.I32_wrap_i64(v52))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v68 = v65
									m.G0 = v10 + int32(48)
									return v68
								}
							}
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							if v56 != int32(453) {
								v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
								if v61 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v77
										F_errmsg_internal(m, int32(_a_F_make_range_0), v10)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_make_range_1), int32(2218), int32(_a_F_make_range_2))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v65 = F_pg_detoast_datum(m, base.I32_wrap_i64(v52))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v68 = v65
										m.G0 = v10 + int32(48)
										return v68
									}
								}
							} else {
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
								if v60 != 0 {
									v68 = int32(0)
									m.G0 = v10 + int32(48)
									return v68
								} else {
									v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
									if v61 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v77
											F_errmsg_internal(m, int32(_a_F_make_range_0), v10)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_make_range_1), int32(2218), int32(_a_F_make_range_2))
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v65 = F_pg_detoast_datum(m, base.I32_wrap_i64(v52))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v68 = v65
											m.G0 = v10 + int32(48)
											return v68
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			if v18 != int32(453) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
				if v23 == int32(0) {
					v68 = v12
					m.G0 = v10 + int32(48)
					return v68
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v30 = int32(1)
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v26)>>(uint(int32(2))%32))-v30))))
					if v32&v30 != 0 {
						v68 = v12
						m.G0 = v10 + int32(48)
						return v68
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l4
						v38 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v38)
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)) = uint8(v38)
						v42 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v10)+26)) = uint16(v42)
						*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_u(v12)
						v47 = l0 + int32(240)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v47
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
						v52 = m.T0[v51].(func(*base.Module, int32) int64)(m, v10+int32(8))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if l4 == int32(0) {
								v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
								if v61 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = v77
										F_errmsg_internal(m, int32(_a_F_make_range_0), v10)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_make_range_1), int32(2218), int32(_a_F_make_range_2))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v65 = F_pg_detoast_datum(m, base.I32_wrap_i64(v52))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v68 = v65
										m.G0 = v10 + int32(48)
										return v68
									}
								}
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								if v56 != int32(453) {
									v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
									if v61 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v77
											F_errmsg_internal(m, int32(_a_F_make_range_0), v10)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_make_range_1), int32(2218), int32(_a_F_make_range_2))
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v65 = F_pg_detoast_datum(m, base.I32_wrap_i64(v52))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v68 = v65
											m.G0 = v10 + int32(48)
											return v68
										}
									}
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
									if v60 != 0 {
										v68 = int32(0)
										m.G0 = v10 + int32(48)
										return v68
									} else {
										v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
										if v61 == int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = v77
												F_errmsg_internal(m, int32(_a_F_make_range_0), v10)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_make_range_1), int32(2218), int32(_a_F_make_range_2))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v65 = F_pg_detoast_datum(m, base.I32_wrap_i64(v52))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v68 = v65
												m.G0 = v10 + int32(48)
												return v68
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
				if v22 != 0 {
					v68 = int32(0)
					m.G0 = v10 + int32(48)
					return v68
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
					if v23 == int32(0) {
						v68 = v12
						m.G0 = v10 + int32(48)
						return v68
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v30 = int32(1)
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v26)>>(uint(int32(2))%32))-v30))))
						if v32&v30 != 0 {
							v68 = v12
							m.G0 = v10 + int32(48)
							return v68
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l4
							v38 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v38)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)) = uint8(v38)
							v42 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v10)+26)) = uint16(v42)
							*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_extend_i32_u(v12)
							v47 = l0 + int32(240)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v47
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
							v52 = m.T0[v51].(func(*base.Module, int32) int64)(m, v10+int32(8))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if l4 == int32(0) {
									v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
									if v61 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v77
											F_errmsg_internal(m, int32(_a_F_make_range_0), v10)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_make_range_1), int32(2218), int32(_a_F_make_range_2))
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v65 = F_pg_detoast_datum(m, base.I32_wrap_i64(v52))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v68 = v65
											m.G0 = v10 + int32(48)
											return v68
										}
									}
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
									if v56 != int32(453) {
										v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
										if v61 == int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = v77
												F_errmsg_internal(m, int32(_a_F_make_range_0), v10)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_make_range_1), int32(2218), int32(_a_F_make_range_2))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v65 = F_pg_detoast_datum(m, base.I32_wrap_i64(v52))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v68 = v65
												m.G0 = v10 + int32(48)
												return v68
											}
										}
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
										if v60 != 0 {
											v68 = int32(0)
											m.G0 = v10 + int32(48)
											return v68
										} else {
											v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
											if v61 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
													*(*int32)(unsafe.Add(mBase, uint32(v10))) = v77
													F_errmsg_internal(m, int32(_a_F_make_range_0), v10)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_make_range_1), int32(2218), int32(_a_F_make_range_2))
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v65 = F_pg_detoast_datum(m, base.I32_wrap_i64(v52))
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													v68 = v65
													m.G0 = v10 + int32(48)
													return v68
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
}
func F_range_(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	if base.Ui32(l2) < base.Ui32(l1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v13 = v11
	goto L6
L5:
	;
	v13 = int32(11)
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v13
	return int32(0)
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v19 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	v82 = int32(_a_F_range__0)
	v83 = l2 - l1
	if base.Ui32(v82) <= base.Ui32(v83) {
		goto L29
	} else {
		goto L30
	}
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v62 != 0 {
		goto L26
	} else {
		goto L27
	}
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v20 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v38 = F_palloc_extended(m, int32(36), int32(2))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L17
	} else {
		goto L19
	}
L14:
	;
	F_pfree(m, v19)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v23 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = int32(-1)
	v28 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v28
	v60 = v19
	goto L10
L17:
	;
	return int32(0)
L18:
	;
	goto L13
L19:
	;
	if v38 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+12)) = int64(4294967296)
	v47 = v38 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v38
	v60 = v38
	goto L10
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v56 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v58 = v56
	goto L25
L24:
	;
	v58 = int32(12)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v58
	v60 = v53
	goto L10
L26:
	;
	return int32(0)
L27:
	;
	goto L28
L28:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v67 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v65+v66<<(uint(v67)%32)))) = l1
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v71+v72<<(uint(v67)%32))+4)) = l2
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = v77 + int32(1)
	return v60
L29:
	;
	v86 = v82
	goto L31
L30:
	;
	v86 = v83
	goto L31
L31:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v87 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v135 != 0 {
		goto L47
	} else {
		goto L48
	}
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v88 <= v86 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v102 = v86 + int32(1)
	v103 = int32(2)
	v104 = v102 << (uint(v103) % 32)
	v108 = F_palloc_extended(m, v104+int32(36), v103)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L17
	} else {
		goto L40
	}
L36:
	;
	F_pfree(m, v87)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L17
	} else {
		goto L39
	}
L37:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	if v90 <= int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = int32(-1)
	v95 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v95
	v133 = v87
	goto L32
L39:
	;
	goto L35
L40:
	;
	if v108 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+4)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v108)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v108)+12)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = int32(0)
	v118 = v108 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v108)+8)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v108)+20)) = v118 + v104
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v108
	v133 = v108
	goto L32
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v125 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v128 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v130 = v128
	goto L46
L45:
	;
	v130 = int32(12)
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v130
	v133 = v125
	goto L32
L47:
	;
	return int32(0)
L48:
	;
	goto L49
L49:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133)+20))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v140 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v138+v139<<(uint(v140)%32)))) = l1
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v133)+20))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v144+v145<<(uint(v140)%32))+4)) = l2
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+12)) = v150 + int32(1)
	v157 = l1
	goto L51
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v253 != 0 {
		goto L86
	} else {
		goto L87
	}
L51:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_range_[0]))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
	if v163 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	return v133
L53:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_range_[0]))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+2)))
	if v202 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L54:
	;
	if base.B2i32(v181 == v157)|base.B2i32(base.Ui32(l1) <= base.Ui32(v181))&base.B2i32(base.Ui32(v181) <= base.Ui32(l2)) != 0 {
		goto L53
	} else {
		goto L63
	}
L55:
	;
	if base.Ui32(int32(127)) < base.Ui32(v157) {
		goto L53
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+64))
	v179 = m.T0[v178].(func(*base.Module, int32, int32) int32)(m, v157, v162)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L17
	} else {
		goto L62
	}
L58:
	;
	if base.Ui32((v157-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v176 = v157 | int32(32)
	goto L61
L60:
	;
	v176 = v157
	goto L61
L61:
	;
	v181 = v176
	goto L54
L62:
	;
	v181 = v179
	goto L54
L63:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v188 <= v187 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	goto L50
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v187 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v193+v187<<(uint(int32(2))%32)))) = v181
	goto L53
L67:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_range_[1]))
	if v242 != 0 {
		goto L81
	} else {
		goto L82
	}
L68:
	;
	if base.B2i32(v222 == v157)|base.B2i32(base.Ui32(l1) <= base.Ui32(v222))&base.B2i32(base.Ui32(v222) <= base.Ui32(l2)) != 0 {
		goto L67
	} else {
		goto L77
	}
L69:
	;
	if base.Ui32(int32(127)) < base.Ui32(v157) {
		goto L67
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+60))
	v220 = m.T0[v219].(func(*base.Module, int32, int32) int32)(m, v157, v201)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L17
	} else {
		goto L76
	}
L72:
	;
	if base.Ui32((v157-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v215 = v157 + int32(224)
	goto L75
L74:
	;
	v215 = v157
	goto L75
L75:
	;
	v222 = v215 & int32(255)
	goto L68
L76:
	;
	v222 = v220
	goto L68
L77:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v229 <= v228 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	goto L50
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v228 + int32(1)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v234+v228<<(uint(int32(2))%32)))) = v222
	goto L67
L81:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L17
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v246 = v157 + int32(1)
	if base.Ui32(v246) <= base.Ui32(l2) {
		v157 = v246
		goto L51
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	goto L52
L86:
	;
	v255 = v253
	goto L88
L87:
	;
	v255 = int32(19)
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v255
	return int32(0)
}
func F_range_adjacent_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v11 == v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_range_deserialize(m, l0, l1, v9-int32(-64), v9+int32(32), v9+int32(15))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L57
	}
L4:
	;
	return int32(0)
L5:
	;
	F_range_deserialize(m, l0, l2, v9+int32(48), v9+int32(16), v9+int32(14))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v32 = int32(0)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
	if v33 != 0 {
		v198 = v32
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v9 + int32(112)
	return v198
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)))
	if v34&int32(1) != 0 {
		v198 = v32
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v37
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v9)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+104)) = v41
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v9)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+96)) = v43
	v45 = base.I32_wrap_i64(v41)
	if base.I32_wrap_i64(v39)&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v116
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v9)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = v118
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v9)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+104)) = v120
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v9)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+96)) = v122
	v124 = base.I32_wrap_i64(v120)
	if base.I32_wrap_i64(v118)&int32(1) != 0 {
		goto L37
	} else {
		goto L38
	}
L11:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+89)))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+105)))
	if v111 != v112 {
		v198 = int32(1)
		goto L7
	} else {
		goto L33
	}
L12:
	;
	if v70 != 0 {
		goto L10
	} else {
		goto L32
	}
L13:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v74 == int32(0) {
		goto L10
	} else {
		goto L29
	}
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+90)))
	if v45&int32(1) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if v45&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+106)))
	if v52 == v49 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v49&int32(1) != 0 {
		goto L13
	} else {
		goto L22
	}
L20:
	;
	if v49&int32(1) != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L10
L22:
	;
	goto L10
L23:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+106)))
	if v60 == int32(0) {
		goto L13
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v9)+80))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v9)+96))
	v68 = F_FunctionCall2Coll(m, l0+int32(212), v65, v66, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L27
	}
L26:
	;
	goto L10
L27:
	;
	v70 = base.I32_wrap_i64(v68)
	if int32(0) <= v70 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	goto L13
L29:
	;
	v77 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+90)) = uint8(v77)
	v80 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+106)) = uint8(v80)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+89)))
	v84 = v82 ^ v77
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+89)) = uint8(v84)
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+105)))
	v88 = v86 ^ v77
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+105)) = uint8(v88)
	v96 = F_make_range(m, l0, v9+int32(80), v9+int32(96), v80, v80)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v102 = int32(1)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+int32(base.Ui32(v98)>>(uint(int32(2))%32))-v102))))
	if v104&v102 == int32(0) {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v198 = v77
	goto L7
L32:
	;
	goto L11
L33:
	;
	goto L10
L34:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+89)))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+105)))
	v198 = base.B2i32(v193 != v194)
	goto L7
L35:
	;
	if v152 == int32(0) {
		goto L34
	} else {
		goto L56
	}
L36:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v156 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L37:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+90)))
	if v124&int32(1) != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if v124&int32(1) != 0 {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+106)))
	if v131 == v128 {
		goto L34
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v128&int32(1) != 0 {
		goto L36
	} else {
		goto L45
	}
L43:
	;
	if v128&int32(1) != 0 {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v198 = int32(0)
	goto L7
L45:
	;
	v198 = int32(0)
	goto L7
L46:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+106)))
	if v141 == int32(0) {
		goto L36
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v9)+80))
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v9)+96))
	v150 = F_FunctionCall2Coll(m, l0+int32(212), v147, v148, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L50
	}
L49:
	;
	v198 = int32(0)
	goto L7
L50:
	;
	v152 = base.I32_wrap_i64(v150)
	if int32(0) <= v152 {
		goto L35
	} else {
		goto L51
	}
L51:
	;
	goto L36
L52:
	;
	v198 = int32(0)
	goto L7
L53:
	;
	goto L54
L54:
	;
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+106)) = uint8(v160)
	v162 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+90)) = uint8(v162)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+89)))
	v166 = v164 ^ v162
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+89)) = uint8(v166)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+105)))
	v170 = v168 ^ v162
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+105)) = uint8(v170)
	v178 = F_make_range(m, l0, v9+int32(80), v9+int32(96), v160, v160)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v184 = int32(1)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+int32(base.Ui32(v180)>>(uint(int32(2))%32))-v184))))
	v198 = v186 & v184
	goto L7
L56:
	;
	v198 = int32(0)
	goto L7
L57:
	;
	F_errmsg_internal(m, int32(_a_F_range_adjacent_internal_0), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_range_adjacent_internal_1), int32(815), int32(_a_F_range_adjacent_internal_2))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_range_agg_finalfn(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v126 int64
	_ = v126
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = v11 + int32(12)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 == v2 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L28
	} else {
		goto L49
	}
L2:
	;
	if v44 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	v44 = v41
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
	v41 = v37
	goto L3
L5:
	;
	v33 = int32(0)
	if v14 == v33 {
		v41 = v33
		goto L3
	} else {
		goto L15
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	switch v19 - int32(435) {
	case 0:
		goto L8
	case 1:
		goto L7
	default:
		goto L5
	}
L7:
	;
	if v14 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	if v14 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = int32(1)
	goto L2
L10:
	;
	goto L11
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+168))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v36 = v26
	v37 = int32(1)
	goto L4
L12:
	;
	v44 = int32(2)
	goto L2
L13:
	;
	goto L14
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+376))
	v36 = v31
	v37 = int32(2)
	goto L4
L15:
	;
	v36 = v33
	v37 = v2
	goto L4
L16:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v45 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L28
	} else {
		goto L46
	}
L19:
	;
	m.G0 = v11 + int32(16)
	return v126
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	if v53 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v48 != 0 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v50 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v50)
	v126 = int64(0)
	goto L19
L24:
	;
	goto L23
L25:
	;
	v56 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
	v126 = int64(0)
	goto L19
L26:
	;
	goto L27
L27:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = F_get_fn_expr_rettype(m, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	return int64(0)
L29:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	if v65 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v79 = F_palloc0(m, v53<<(uint(int32(2))%32))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L28
	} else {
		goto L37
	}
L31:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v66 == v60 {
		v76 = v65
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v69 = F_lookup_type_cache(m, v60, int32(_a_F_range_agg_finalfn_0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L28
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+296))
	if v71 == int32(0) {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = v69
	v76 = v69
	goto L30
L37:
	;
	if int32(0) < v53 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v84 = int32(0)
	goto L41
L39:
	;
	goto L40
L40:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v76)+296))
	v115 = F_make_multirange(m, v60, v114, v53, v79)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L28
	} else {
		goto L45
	}
L41:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95+v84<<(uint(int32(3))%32))))
	v100 = F_pg_detoast_datum(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L28
	} else {
		goto L43
	}
L42:
	;
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79+v84<<(uint(int32(2))%32)))) = v100
	v104 = v84 + int32(1)
	if v104 != v53 {
		v84 = v104
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v126 = base.I64_extend_i32_u(v115)
	goto L19
L46:
	;
	F_errmsg_internal(m, int32(_a_F_range_agg_finalfn_1), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L28
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_range_agg_finalfn_2), int32(1460), int32(_a_F_range_agg_finalfn_3))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L28
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v60
	F_errmsg_internal(m, int32(_a_F_range_agg_finalfn_4), v11)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L28
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_range_agg_finalfn_2), int32(561), int32(_a_F_range_agg_finalfn_5))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L28
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_range_agg_transfn(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = v7 + int32(12)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == v2 {
		v29 = int32(0)
		if v10 == v29 {
			v37 = v29
		} else {
			v32 = v29
			v33 = v2
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
			v37 = v33
		}
		v40 = v37
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		switch v15 - int32(435) {
		case 0:
			if v10 == int32(0) {
				v40 = int32(1)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+168))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
				v32 = v22
				v33 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
				v37 = v33
				v40 = v37
			}
		case 1:
			if v10 == int32(0) {
				v40 = int32(2)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+376))
				v32 = v27
				v33 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
				v37 = v33
				v40 = v37
			}
		default:
			v29 = int32(0)
			if v10 == v29 {
				v37 = v29
			} else {
				v32 = v29
				v33 = v2
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
				v37 = v33
			}
			v40 = v37
		}
	}
	if v40 != 0 {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v43 = F_get_fn_expr_argtype(m, v41, int32(1))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int64(0)
		} else {
			v47 = F_type_is_range(m, v43)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int64(0)
			} else {
				if v47 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int64(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_range_agg_transfn_0), int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_range_agg_transfn_1), int32(1428), int32(_a_F_range_agg_transfn_2))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v51 == int32(1) {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						v56 = F_initArrayResult(m, v43, v54, int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							v59 = v56
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
							if v60 == int32(0) {
								v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
								v66 = F_accumArrayResult(m, v59, v63, int32(0), v43, v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int64(0)
								} else {
									m.G0 = v7 + int32(16)
									return base.I64_extend_i32_u(v59)
								}
							} else {
								m.G0 = v7 + int32(16)
								return base.I64_extend_i32_u(v59)
							}
						}
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v59 = v58
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
						if v60 == int32(0) {
							v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							v66 = F_accumArrayResult(m, v59, v63, int32(0), v43, v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int64(0)
							} else {
								m.G0 = v7 + int32(16)
								return base.I64_extend_i32_u(v59)
							}
						} else {
							m.G0 = v7 + int32(16)
							return base.I64_extend_i32_u(v59)
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_range_agg_transfn_3), int32(0))
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_range_agg_transfn_1), int32(1424), int32(_a_F_range_agg_transfn_2))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_range_contained_by_multirange(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = F_multirange_contains_range_internal(m, v33, v17, v12)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v34)
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_contained_by_multirange_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_contained_by_multirange_1), v9)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_range_contained_by_multirange_2), int32(561), int32(_a_F_range_contained_by_multirange_3))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = F_multirange_contains_range_internal(m, v33, v17, v12)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v34)
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_contained_by_multirange_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_contained_by_multirange_1), v9)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_range_contained_by_multirange_2), int32(561), int32(_a_F_range_contained_by_multirange_3))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = F_multirange_contains_range_internal(m, v33, v17, v12)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v34)
						}
					}
				}
			}
		}
	}
}
func F_range_contains_multirange(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = F_range_contains_multirange_internal(m, v33, v12, v17)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v34)
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_contains_multirange_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_contains_multirange_1), v9)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_range_contains_multirange_2), int32(561), int32(_a_F_range_contains_multirange_3))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = F_range_contains_multirange_internal(m, v33, v12, v17)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v34)
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_contains_multirange_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_contains_multirange_1), v9)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_range_contains_multirange_2), int32(561), int32(_a_F_range_contains_multirange_3))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = F_range_contains_multirange_internal(m, v33, v12, v17)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v34)
						}
					}
				}
			}
		}
	}
}
func F_range_eq_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v10 == v11 {
		F_range_deserialize(m, l0, l1, v8-int32(-64), v8+int32(32), v8+int32(15))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_range_deserialize(m, l0, l2, v8+int32(48), v8+int32(16), v8+int32(14))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = int32(1)
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
				if v32&v31&base.B2i32(v35 == v31) != 0 {
					v108 = v31
					m.G0 = v8 + int32(80)
					return v108 & int32(1)
				} else {
					v39 = int32(0)
					if v32 != v35 {
						v108 = v39
						m.G0 = v8 + int32(80)
						return v108 & int32(1)
					} else {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+56)))
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+72)))
						if v42 == int32(1) {
							if v41&int32(1) == int32(0) {
								v108 = v39
								m.G0 = v8 + int32(80)
								return v108 & int32(1)
							} else {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+74)))
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+58)))
								if v49 == v50 {
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+40)))
									if v77 == int32(1) {
										if v76&int32(1) == int32(0) {
											v108 = v39
										} else {
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+42)))
											v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
											v108 = base.B2i32(v84 == v85)
										}
										m.G0 = v8 + int32(80)
										return v108 & int32(1)
									} else {
										if v76&int32(1) != 0 {
											v108 = v39
											m.G0 = v8 + int32(80)
											return v108 & int32(1)
										} else {
											v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v92 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
											v93 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
											v94 = F_FunctionCall2Coll(m, l0+int32(212), v91, v92, v93)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												if base.I32_wrap_i64(v94) != 0 {
													v108 = v39
												} else {
													v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+41)))
													v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+25)))
													if (v97|v98)&int32(1) != 0 {
														v108 = v97 & v98
													} else {
														v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+42)))
														v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
														v108 = base.B2i32(v103 == v104)
													}
												}
												m.G0 = v8 + int32(80)
												return v108 & int32(1)
											}
										}
									}
								} else {
									v108 = v39
									m.G0 = v8 + int32(80)
									return v108 & int32(1)
								}
							}
						} else {
							if v41&int32(1) != 0 {
								v108 = v39
								m.G0 = v8 + int32(80)
								return v108 & int32(1)
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v57 = *(*int64)(unsafe.Add(mBase, uint32(v8)+64))
								v58 = *(*int64)(unsafe.Add(mBase, uint32(v8)+48))
								v59 = F_FunctionCall2Coll(m, l0+int32(212), v56, v57, v58)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									if base.I32_wrap_i64(v59) != 0 {
										v108 = v39
										m.G0 = v8 + int32(80)
										return v108 & int32(1)
									} else {
										v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+57)))
										v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+73)))
										if v63 == int32(0) {
											if v62&int32(1) != 0 {
												v108 = v39
												m.G0 = v8 + int32(80)
												return v108 & int32(1)
											} else {
												v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+74)))
												v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+58)))
												if v68 == v69 {
													v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
													v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+40)))
													if v77 == int32(1) {
														if v76&int32(1) == int32(0) {
															v108 = v39
														} else {
															v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+42)))
															v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
															v108 = base.B2i32(v84 == v85)
														}
														m.G0 = v8 + int32(80)
														return v108 & int32(1)
													} else {
														if v76&int32(1) != 0 {
															v108 = v39
															m.G0 = v8 + int32(80)
															return v108 & int32(1)
														} else {
															v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
															v92 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
															v93 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
															v94 = F_FunctionCall2Coll(m, l0+int32(212), v91, v92, v93)
															mBase = m.M
															v95 = m.ExcPending
															if v95 != 0 {
																return int32(0)
															} else {
																if base.I32_wrap_i64(v94) != 0 {
																	v108 = v39
																} else {
																	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+41)))
																	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+25)))
																	if (v97|v98)&int32(1) != 0 {
																		v108 = v97 & v98
																	} else {
																		v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+42)))
																		v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
																		v108 = base.B2i32(v103 == v104)
																	}
																}
																m.G0 = v8 + int32(80)
																return v108 & int32(1)
															}
														}
													}
												} else {
													v108 = v39
													m.G0 = v8 + int32(80)
													return v108 & int32(1)
												}
											}
										} else {
											if v62&int32(1) == int32(0) {
												v108 = v39
												m.G0 = v8 + int32(80)
												return v108 & int32(1)
											} else {
												v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
												v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+40)))
												if v77 == int32(1) {
													if v76&int32(1) == int32(0) {
														v108 = v39
													} else {
														v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+42)))
														v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
														v108 = base.B2i32(v84 == v85)
													}
													m.G0 = v8 + int32(80)
													return v108 & int32(1)
												} else {
													if v76&int32(1) != 0 {
														v108 = v39
														m.G0 = v8 + int32(80)
														return v108 & int32(1)
													} else {
														v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
														v92 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
														v93 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
														v94 = F_FunctionCall2Coll(m, l0+int32(212), v91, v92, v93)
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return int32(0)
														} else {
															if base.I32_wrap_i64(v94) != 0 {
																v108 = v39
															} else {
																v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+41)))
																v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+25)))
																if (v97|v98)&int32(1) != 0 {
																	v108 = v97 & v98
																} else {
																	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+42)))
																	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
																	v108 = base.B2i32(v103 == v104)
																}
															}
															m.G0 = v8 + int32(80)
															return v108 & int32(1)
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
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v118 = m.ExcPending
		if v118 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_range_eq_internal_0), int32(0))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_eq_internal_1), int32(590), int32(_a_F_range_eq_internal_2))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
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
func F_range_get_typcache(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v10 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		if v11 == l1 {
			v23 = v10
			m.G0 = v7 + int32(16)
			return v23
		} else {
			v14 = F_lookup_type_cache(m, l1, int32(2048))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+200))
				if v18 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
						F_errmsg_internal(m, int32(_a_F_range_get_typcache_0), v7)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_range_get_typcache_1), int32(1946), int32(_a_F_range_get_typcache_2))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v14
					v23 = v14
					m.G0 = v7 + int32(16)
					return v23
				}
			}
		}
	} else {
		v14 = F_lookup_type_cache(m, l1, int32(2048))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+200))
			if v18 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg_internal(m, int32(_a_F_range_get_typcache_0), v7)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_range_get_typcache_1), int32(1946), int32(_a_F_range_get_typcache_2))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v14
				v23 = v14
				m.G0 = v7 + int32(16)
				return v23
			}
		}
	}
}
func F_range_intersect_agg_transfn(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = v7 + int32(12)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == v2 {
		v29 = int32(0)
		if v10 == v29 {
			v37 = v29
		} else {
			v32 = v29
			v33 = v2
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
			v37 = v33
		}
		v40 = v37
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		switch v15 - int32(435) {
		case 0:
			if v10 == int32(0) {
				v40 = int32(1)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+168))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
				v32 = v22
				v33 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
				v37 = v33
				v40 = v37
			}
		case 1:
			if v10 == int32(0) {
				v40 = int32(2)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+376))
				v32 = v27
				v33 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
				v37 = v33
				v40 = v37
			}
		default:
			v29 = int32(0)
			if v10 == v29 {
				v37 = v29
			} else {
				v32 = v29
				v33 = v2
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v32
				v37 = v33
			}
			v40 = v37
		}
	}
	if v40 != 0 {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v43 = F_get_fn_expr_argtype(m, v41, int32(1))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int64(0)
		} else {
			v47 = F_type_is_range(m, v43)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int64(0)
			} else {
				if v47 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int64(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_range_intersect_agg_transfn_0), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_range_intersect_agg_transfn_1), int32(1404), int32(_a_F_range_intersect_agg_transfn_2))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
					if v52 != 0 {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
						if v53 == v43 {
							v63 = v52
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v65 = F_pg_detoast_datum(m, v64)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int64(0)
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								v68 = F_pg_detoast_datum(m, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int64(0)
								} else {
									v70 = F_range_intersect_internal(m, v63, v65, v68)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int64(0)
									} else {
										m.G0 = v7 + int32(16)
										return base.I64_extend_i32_u(v70)
									}
								}
							}
						} else {
							v56 = F_lookup_type_cache(m, v43, int32(2048))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+200))
								if v58 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v43
										F_errmsg_internal(m, int32(_a_F_range_intersect_agg_transfn_3), v7)
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_range_intersect_agg_transfn_1), int32(1946), int32(_a_F_range_intersect_agg_transfn_4))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v56
									v63 = v56
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v65 = F_pg_detoast_datum(m, v64)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int64(0)
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										v68 = F_pg_detoast_datum(m, v67)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int64(0)
										} else {
											v70 = F_range_intersect_internal(m, v63, v65, v68)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int64(0)
											} else {
												m.G0 = v7 + int32(16)
												return base.I64_extend_i32_u(v70)
											}
										}
									}
								}
							}
						}
					} else {
						v56 = F_lookup_type_cache(m, v43, int32(2048))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+200))
							if v58 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v43
									F_errmsg_internal(m, int32(_a_F_range_intersect_agg_transfn_3), v7)
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_range_intersect_agg_transfn_1), int32(1946), int32(_a_F_range_intersect_agg_transfn_4))
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v56
								v63 = v56
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v65 = F_pg_detoast_datum(m, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int64(0)
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v68 = F_pg_detoast_datum(m, v67)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int64(0)
									} else {
										v70 = F_range_intersect_internal(m, v63, v65, v68)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int64(0)
										} else {
											m.G0 = v7 + int32(16)
											return base.I64_extend_i32_u(v70)
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
		v80 = m.ExcPending
		if v80 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_range_intersect_agg_transfn_5), int32(0))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_range_intersect_agg_transfn_1), int32(1400), int32(_a_F_range_intersect_agg_transfn_2))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_range_lower_inc(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14367(m, l0, int64(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_range_minus_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v264 int64
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int64
	_ = v322
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	F_range_deserialize(m, l0, l1, v11-int32(-64), v11+int32(32), v11+int32(15))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_range_deserialize(m, l0, l2, v11+int32(48), v11+int32(16), v11+int32(14))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v31 != 0 {
		v374 = l1
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L194
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L190
	}
L6:
	;
	m.G0 = v11 + int32(112)
	return v374
L7:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)))
	if v32&int32(1) != 0 {
		v374 = l1
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+56)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+72)))
	if v36 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)))
	if v175 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L10:
	;
	v172 = v166
	v173 = v169
	v174 = int32(0)
	goto L9
L11:
	;
	v119 = int32(1)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)))
	if v120 == v119 {
		goto L63
	} else {
		goto L64
	}
L12:
	;
	v112 = int32(1)
	if v70&v112 != 0 {
		goto L60
	} else {
		goto L61
	}
L13:
	;
	v96 = int32(1)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)))
	if v97 == v96 {
		goto L50
	} else {
		goto L51
	}
L14:
	;
	v91 = int32(1)
	if v39&v91 != 0 {
		goto L47
	} else {
		goto L48
	}
L15:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+74)))
	if v35&int32(1) == int32(0) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v35&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)))
	if v45 == v39 {
		v95 = int32(0)
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v48 = int32(1)
	if v39&v48 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v51 = int32(-1)
	goto L22
L21:
	;
	v51 = v48
	goto L22
L22:
	;
	v95 = v51
	goto L13
L23:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)))
	if v56 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v11)+64))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
	v63 = F_FunctionCall2Coll(m, l0+int32(212), v60, v61, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L29
	}
L26:
	;
	v57 = int32(1)
	goto L28
L27:
	;
	v57 = int32(-1)
	goto L28
L28:
	;
	v118 = v57
	goto L11
L29:
	;
	v65 = base.I32_wrap_i64(v63)
	if v65 != 0 {
		v118 = v65
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+57)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+73)))
	if v67 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+74)))
	if v66&int32(1) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v66&int32(1) != 0 {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)))
	if v70 != v75 {
		goto L12
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v78 = int32(1)
	if v70&v78 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v118 = int32(0)
	goto L11
L38:
	;
	v82 = v78
	goto L40
L39:
	;
	v82 = int32(-1)
	goto L40
L40:
	;
	v118 = v82
	goto L11
L41:
	;
	v118 = int32(0)
	goto L11
L42:
	;
	goto L43
L43:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)))
	if v88 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v89 = int32(-1)
	goto L46
L45:
	;
	v89 = int32(1)
	goto L46
L46:
	;
	v118 = v89
	goto L11
L47:
	;
	v94 = int32(-1)
	goto L49
L48:
	;
	v94 = v91
	goto L49
L49:
	;
	v95 = v94
	goto L13
L50:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
	if v100 == v39 {
		v172 = v95
		v173 = int32(0)
		v174 = v96
		goto L9
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v108 = int32(1)
	if v39&v108 != 0 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v103 = int32(1)
	if v39&v103 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v106 = int32(-1)
	goto L56
L55:
	;
	v106 = v103
	goto L56
L56:
	;
	v172 = v95
	v173 = v106
	v174 = v96
	goto L9
L57:
	;
	v111 = int32(-1)
	goto L59
L58:
	;
	v111 = v108
	goto L59
L59:
	;
	v166 = v95
	v169 = v111
	goto L10
L60:
	;
	v116 = v112
	goto L62
L61:
	;
	v116 = int32(-1)
	goto L62
L62:
	;
	v118 = v116
	goto L11
L63:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
	if v125 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v127 = int32(0)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v11)+64))
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
	v133 = F_FunctionCall2Coll(m, l0+int32(212), v130, v131, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L69
	}
L66:
	;
	v126 = int32(1)
	goto L68
L67:
	;
	v126 = int32(-1)
	goto L68
L68:
	;
	v172 = v118
	v173 = v126
	v174 = v119
	goto L9
L69:
	;
	v135 = base.I32_wrap_i64(v133)
	if v135 != 0 {
		v172 = v118
		v173 = v135
		v174 = v127
		goto L9
	} else {
		goto L70
	}
L70:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+73)))
	if v137 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+74)))
	if v136&int32(1) == int32(0) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if v136&int32(1) != 0 {
		goto L86
	} else {
		goto L87
	}
L74:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
	if v145 == v140 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v153 = int32(1)
	if v140&v153 != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v172 = v118
	v173 = int32(0)
	v174 = v127
	goto L9
L78:
	;
	goto L79
L79:
	;
	v148 = int32(1)
	if v140&v148 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v152 = v148
	goto L82
L81:
	;
	v152 = int32(-1)
	goto L82
L82:
	;
	v166 = v118
	v169 = v152
	goto L10
L83:
	;
	v157 = v153
	goto L85
L84:
	;
	v157 = int32(-1)
	goto L85
L85:
	;
	v172 = v118
	v173 = v157
	v174 = v127
	goto L9
L86:
	;
	v172 = v118
	v173 = int32(0)
	v174 = v127
	goto L9
L87:
	;
	goto L88
L88:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
	if v163 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v164 = int32(-1)
	goto L91
L90:
	;
	v164 = int32(1)
	goto L91
L91:
	;
	v166 = v118
	v169 = v164
	goto L10
L92:
	;
	v306 = int32(0)
	if base.B2i32(v304 < v306)|base.B2i32(v306 < v173) != 0 {
		v374 = l1
		goto L6
	} else {
		goto L179
	}
L93:
	;
	if int32(0) <= v172 {
		v303 = v296
		v304 = v297
		goto L92
	} else {
		goto L177
	}
L94:
	;
	if v174 != 0 {
		goto L148
	} else {
		goto L149
	}
L95:
	;
	v248 = int32(1)
	if v209&v248 != 0 {
		goto L145
	} else {
		goto L146
	}
L96:
	;
	if v174 != 0 {
		goto L133
	} else {
		goto L134
	}
L97:
	;
	v230 = int32(1)
	if v178&v230 != 0 {
		goto L130
	} else {
		goto L131
	}
L98:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)))
	if v35&int32(1) == int32(0) {
		goto L97
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if v35&int32(1) != 0 {
		goto L106
	} else {
		goto L107
	}
L101:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)))
	if v178 == v184 {
		v234 = int32(0)
		goto L96
	} else {
		goto L102
	}
L102:
	;
	v187 = int32(1)
	if v178&v187 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v190 = int32(-1)
	goto L105
L104:
	;
	v190 = v187
	goto L105
L105:
	;
	v234 = v190
	goto L96
L106:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)))
	if v195 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	goto L108
L108:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
	v202 = F_FunctionCall2Coll(m, l0+int32(212), v199, v200, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L112
	}
L109:
	;
	v196 = int32(1)
	goto L111
L110:
	;
	v196 = int32(-1)
	goto L111
L111:
	;
	v254 = v196
	goto L94
L112:
	;
	v204 = base.I32_wrap_i64(v202)
	if v204 != 0 {
		v254 = v204
		goto L94
	} else {
		goto L113
	}
L113:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+57)))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)))
	if v206 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)))
	if v205&int32(1) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	if v205&int32(1) != 0 {
		goto L124
	} else {
		goto L125
	}
L117:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)))
	if v214 != v209 {
		goto L95
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v217 = int32(1)
	if v209&v217 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v254 = int32(0)
	goto L94
L121:
	;
	v221 = v217
	goto L123
L122:
	;
	v221 = int32(-1)
	goto L123
L123:
	;
	v254 = v221
	goto L94
L124:
	;
	v254 = int32(0)
	goto L94
L125:
	;
	goto L126
L126:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)))
	if v227 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v228 = int32(-1)
	goto L129
L128:
	;
	v228 = int32(1)
	goto L129
L129:
	;
	v254 = v228
	goto L94
L130:
	;
	v233 = int32(-1)
	goto L132
L131:
	;
	v233 = v230
	goto L132
L132:
	;
	v234 = v233
	goto L96
L133:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
	if v235 == v178 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	v244 = int32(1)
	if v178&v244 != 0 {
		goto L142
	} else {
		goto L143
	}
L136:
	;
	v303 = int32(0)
	v304 = v234
	goto L92
L137:
	;
	goto L138
L138:
	;
	v239 = int32(1)
	if v178&v239 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v242 = int32(-1)
	goto L141
L140:
	;
	v242 = v239
	goto L141
L141:
	;
	v296 = v242
	v297 = v234
	goto L93
L142:
	;
	v247 = int32(-1)
	goto L144
L143:
	;
	v247 = v244
	goto L144
L144:
	;
	v296 = v247
	v297 = v234
	goto L93
L145:
	;
	v252 = v248
	goto L147
L146:
	;
	v252 = int32(-1)
	goto L147
L147:
	;
	v254 = v252
	goto L94
L148:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
	if v257 != 0 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
	v264 = F_FunctionCall2Coll(m, l0+int32(212), v261, v262, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L154
	}
L151:
	;
	v258 = int32(1)
	goto L153
L152:
	;
	v258 = int32(-1)
	goto L153
L153:
	;
	v296 = v258
	v297 = v254
	goto L93
L154:
	;
	v266 = base.I32_wrap_i64(v264)
	if v266 != 0 {
		v296 = v266
		v297 = v254
		goto L93
	} else {
		goto L155
	}
L155:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)))
	if v268 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)))
	if v267&int32(1) == int32(0) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L158
L158:
	;
	if v267&int32(1) != 0 {
		goto L171
	} else {
		goto L172
	}
L159:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
	if v276 == v271 {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	goto L161
L161:
	;
	v284 = int32(1)
	if v271&v284 != 0 {
		goto L168
	} else {
		goto L169
	}
L162:
	;
	v303 = int32(0)
	v304 = v254
	goto L92
L163:
	;
	goto L164
L164:
	;
	v279 = int32(1)
	if v271&v279 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v283 = v279
	goto L167
L166:
	;
	v283 = int32(-1)
	goto L167
L167:
	;
	v296 = v283
	v297 = v254
	goto L93
L168:
	;
	v288 = v284
	goto L170
L169:
	;
	v288 = int32(-1)
	goto L170
L170:
	;
	v296 = v288
	v297 = v254
	goto L93
L171:
	;
	v303 = int32(0)
	v304 = v254
	goto L92
L172:
	;
	goto L173
L173:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
	if v294 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v295 = int32(-1)
	goto L176
L175:
	;
	v295 = int32(1)
	goto L176
L176:
	;
	v296 = v295
	v297 = v254
	goto L93
L177:
	;
	if int32(0) < v296 {
		goto L5
	} else {
		goto L178
	}
L178:
	;
	v303 = v296
	v304 = v297
	goto L92
L179:
	;
	v311 = int32(0)
	if base.B2i32(v172 < v311)|base.B2i32(v311 < v303) == v311 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v318 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+106)) = uint8(v318)
	v320 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+104)) = uint16(v320)
	v322 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = v322
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v11)+87)) = v320
	v334 = F_make_range(m, l0, v11+int32(96), v11+int32(80), v318, v320)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v336 = int32(0)
	if base.B2i32(v336 < v172)|base.B2i32(v336 < v303) == v336 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v374 = v334
	goto L6
L184:
	;
	v343 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)) = uint8(v343)
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+57)))
	v347 = v345 ^ int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+57)) = uint8(v347)
	v355 = F_make_range(m, l0, v11-int32(-64), v11+int32(48), v343, v343)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	if v303|v172 < int32(0) {
		goto L4
	} else {
		goto L188
	}
L187:
	;
	v374 = v355
	goto L6
L188:
	;
	v360 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v360)
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)))
	v364 = v362 ^ v360
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(v364)
	v370 = int32(0)
	v372 = F_make_range(m, l0, v11+int32(16), v11+int32(32), v370, v370)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v374 = v372
	goto L6
L190:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_errmsg(m, int32(_a_F_range_minus_internal_3), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_range_minus_internal_1), int32(1027), int32(_a_F_range_minus_internal_2))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	F_errmsg_internal(m, int32(_a_F_range_minus_internal_0), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_range_minus_internal_1), int32(1049), int32(_a_F_range_minus_internal_2))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_range_minus_multi(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
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
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v249 int64
	_ = v249
	var v250 int64
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v311 int64
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int64
	_ = v485
	var v486 int32
	_ = v486
	var v487 int64
	_ = v487
	var v493 int64
	_ = v493
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v507 int64
	_ = v507
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v23 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L6
	} else {
		goto L211
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L6
	} else {
		goto L208
	}
L3:
	;
	v26 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+16))
	goto L203
L6:
	;
	return int64(0)
L7:
	;
	v30 = int32(_a_F_range_minus_multi_0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_range_minus_multi[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_range_minus_multi[0])) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v36 = F_pg_detoast_datum(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v39 = F_pg_detoast_datum(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v41 != v42 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v45 = F_palloc(m, int32(12))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v49 = F_lookup_type_cache(m, v47, int32(2048))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+200))
	if v51 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v56 = m.G0
	v58 = v56 - int32(80)
	m.G0 = v58
	F_range_deserialize(m, v49, v36, v58-int32(-64), v58+int32(32), v58+int32(15))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	F_range_deserialize(m, v49, v39, v58+int32(48), v58+int32(16), v58+int32(14))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+15)))
	if v76 != 0 {
		v444 = v2
		goto L18
	} else {
		goto L19
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v45
	*(*int32)(unsafe.Add(mBase, _c_F_range_minus_multi[0])) = v31
	goto L5
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L6
	} else {
		goto L199
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45+int32(8)))) = v444
	m.G0 = v58 + int32(80)
	goto L16
L19:
	;
	v77 = int32(1)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+14)))
	if v78 == v77 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v36
	v444 = v77
	goto L18
L21:
	;
	goto L22
L22:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+56)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+72)))
	if v83 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+40)))
	if v223 == int32(1) {
		goto L113
	} else {
		goto L114
	}
L24:
	;
	v220 = int32(0)
	v221 = v215
	v222 = v217
	goto L23
L25:
	;
	v167 = int32(1)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+24)))
	if v168 == v167 {
		goto L77
	} else {
		goto L78
	}
L26:
	;
	v159 = int32(1)
	if v117&v159 != 0 {
		goto L74
	} else {
		goto L75
	}
L27:
	;
	v143 = int32(1)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+24)))
	if v144 == v143 {
		goto L64
	} else {
		goto L65
	}
L28:
	;
	v138 = int32(1)
	if v86&v138 != 0 {
		goto L61
	} else {
		goto L62
	}
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+74)))
	if v82&int32(1) == int32(0) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v82&int32(1) != 0 {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+58)))
	if v86 == v92 {
		v142 = int32(0)
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v95 = int32(1)
	if v86&v95 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v98 = int32(-1)
	goto L36
L35:
	;
	v98 = v95
	goto L36
L36:
	;
	v142 = v98
	goto L27
L37:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+58)))
	if v103 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v49)+208))
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v58)+64))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v58)+48))
	v110 = F_FunctionCall2Coll(m, v49+int32(212), v107, v108, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L43
	}
L40:
	;
	v104 = int32(1)
	goto L42
L41:
	;
	v104 = int32(-1)
	goto L42
L42:
	;
	v166 = v104
	goto L25
L43:
	;
	v112 = base.I32_wrap_i64(v110)
	if v112 != 0 {
		v166 = v112
		goto L25
	} else {
		goto L44
	}
L44:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+57)))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+73)))
	if v114 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+74)))
	if v113&int32(1) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	if v113&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+58)))
	if v122 != v117 {
		goto L26
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v125 = int32(1)
	if v117&v125 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v166 = int32(0)
	goto L25
L52:
	;
	v129 = v125
	goto L54
L53:
	;
	v129 = int32(-1)
	goto L54
L54:
	;
	v166 = v129
	goto L25
L55:
	;
	v166 = int32(0)
	goto L25
L56:
	;
	goto L57
L57:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+58)))
	if v135 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v136 = int32(-1)
	goto L60
L59:
	;
	v136 = int32(1)
	goto L60
L60:
	;
	v166 = v136
	goto L25
L61:
	;
	v141 = int32(-1)
	goto L63
L62:
	;
	v141 = v138
	goto L63
L63:
	;
	v142 = v141
	goto L27
L64:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+26)))
	if v86 == v147 {
		v220 = v143
		v221 = v142
		v222 = v2
		goto L23
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v155 = int32(1)
	if v86&v155 != 0 {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	v150 = int32(1)
	if v86&v150 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v153 = int32(-1)
	goto L70
L69:
	;
	v153 = v150
	goto L70
L70:
	;
	v220 = v143
	v221 = v142
	v222 = v153
	goto L23
L71:
	;
	v158 = int32(-1)
	goto L73
L72:
	;
	v158 = v155
	goto L73
L73:
	;
	v215 = v142
	v217 = v158
	goto L24
L74:
	;
	v163 = v159
	goto L76
L75:
	;
	v163 = int32(-1)
	goto L76
L76:
	;
	v166 = v163
	goto L25
L77:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+26)))
	if v173 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v175 = int32(0)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v49)+208))
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v58)+64))
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v58)+16))
	v181 = F_FunctionCall2Coll(m, v49+int32(212), v178, v179, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L6
	} else {
		goto L83
	}
L80:
	;
	v174 = int32(1)
	goto L82
L81:
	;
	v174 = int32(-1)
	goto L82
L82:
	;
	v220 = v167
	v221 = v166
	v222 = v174
	goto L23
L83:
	;
	v183 = base.I32_wrap_i64(v181)
	if v183 != 0 {
		v220 = v175
		v221 = v166
		v222 = v183
		goto L23
	} else {
		goto L84
	}
L84:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+25)))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+73)))
	if v185 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+74)))
	if v184&int32(1) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	if v184&int32(1) != 0 {
		goto L100
	} else {
		goto L101
	}
L88:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+26)))
	if v193 == v188 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v201 = int32(1)
	if v188&v201 != 0 {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	v220 = v175
	v221 = v166
	v222 = int32(0)
	goto L23
L92:
	;
	goto L93
L93:
	;
	v196 = int32(1)
	if v188&v196 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v200 = v196
	goto L96
L95:
	;
	v200 = int32(-1)
	goto L96
L96:
	;
	v215 = v166
	v217 = v200
	goto L24
L97:
	;
	v205 = v201
	goto L99
L98:
	;
	v205 = int32(-1)
	goto L99
L99:
	;
	v220 = v175
	v221 = v166
	v222 = v205
	goto L23
L100:
	;
	v220 = v175
	v221 = v166
	v222 = int32(0)
	goto L23
L101:
	;
	goto L102
L102:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+26)))
	if v211 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v212 = int32(-1)
	goto L105
L104:
	;
	v212 = int32(1)
	goto L105
L105:
	;
	v215 = v166
	v217 = v212
	goto L24
L106:
	;
	v387 = int32(0)
	if base.B2i32(v222 <= v387)&base.B2i32(v387 <= v384) == v387 {
		goto L189
	} else {
		goto L190
	}
L107:
	;
	v384 = v380
	v385 = int32(0)
	goto L106
L108:
	;
	v344 = int32(0)
	if base.B2i32(v342 <= v344)|base.B2i32(v344 <= v221) != 0 {
		v384 = v341
		v385 = v342
		goto L106
	} else {
		goto L186
	}
L109:
	;
	if v220 != 0 {
		goto L161
	} else {
		goto L162
	}
L110:
	;
	v295 = int32(1)
	if v257&v295 != 0 {
		goto L158
	} else {
		goto L159
	}
L111:
	;
	if v220 != 0 {
		goto L148
	} else {
		goto L149
	}
L112:
	;
	v278 = int32(1)
	if v226&v278 != 0 {
		goto L145
	} else {
		goto L146
	}
L113:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+42)))
	if v82&int32(1) == int32(0) {
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if v82&int32(1) != 0 {
		goto L121
	} else {
		goto L122
	}
L116:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+58)))
	if v226 == v232 {
		v282 = int32(0)
		goto L111
	} else {
		goto L117
	}
L117:
	;
	v235 = int32(1)
	if v226&v235 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v238 = int32(-1)
	goto L120
L119:
	;
	v238 = v235
	goto L120
L120:
	;
	v282 = v238
	goto L111
L121:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+58)))
	if v243 != 0 {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v49)+208))
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v58)+32))
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v58)+48))
	v250 = F_FunctionCall2Coll(m, v49+int32(212), v247, v248, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L6
	} else {
		goto L127
	}
L124:
	;
	v244 = int32(1)
	goto L126
L125:
	;
	v244 = int32(-1)
	goto L126
L126:
	;
	v300 = v244
	goto L109
L127:
	;
	v252 = base.I32_wrap_i64(v250)
	if v252 != 0 {
		v300 = v252
		goto L109
	} else {
		goto L128
	}
L128:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+57)))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+41)))
	if v254 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+42)))
	if v253&int32(1) == int32(0) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	goto L131
L131:
	;
	if v253&int32(1) != 0 {
		goto L139
	} else {
		goto L140
	}
L132:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+58)))
	if v257 != v262 {
		goto L110
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v265 = int32(1)
	if v257&v265 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v300 = int32(0)
	goto L109
L136:
	;
	v269 = v265
	goto L138
L137:
	;
	v269 = int32(-1)
	goto L138
L138:
	;
	v300 = v269
	goto L109
L139:
	;
	v300 = int32(0)
	goto L109
L140:
	;
	goto L141
L141:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+58)))
	if v275 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v276 = int32(-1)
	goto L144
L143:
	;
	v276 = int32(1)
	goto L144
L144:
	;
	v300 = v276
	goto L109
L145:
	;
	v281 = int32(-1)
	goto L147
L146:
	;
	v281 = v278
	goto L147
L147:
	;
	v282 = v281
	goto L111
L148:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+26)))
	if v226 == v283 {
		v380 = v282
		goto L107
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v291 = int32(1)
	if v226&v291 != 0 {
		goto L155
	} else {
		goto L156
	}
L151:
	;
	v286 = int32(1)
	if v226&v286 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v289 = int32(-1)
	goto L154
L153:
	;
	v289 = v286
	goto L154
L154:
	;
	v341 = v282
	v342 = v289
	goto L108
L155:
	;
	v294 = int32(-1)
	goto L157
L156:
	;
	v294 = v291
	goto L157
L157:
	;
	v341 = v282
	v342 = v294
	goto L108
L158:
	;
	v299 = v295
	goto L160
L159:
	;
	v299 = int32(-1)
	goto L160
L160:
	;
	v300 = v299
	goto L109
L161:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+26)))
	if v304 != 0 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	goto L163
L163:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v49)+208))
	v309 = *(*int64)(unsafe.Add(mBase, uint32(v58)+32))
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v58)+16))
	v311 = F_FunctionCall2Coll(m, v49+int32(212), v308, v309, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L6
	} else {
		goto L167
	}
L164:
	;
	v305 = int32(1)
	goto L166
L165:
	;
	v305 = int32(-1)
	goto L166
L166:
	;
	v341 = v300
	v342 = v305
	goto L108
L167:
	;
	v313 = base.I32_wrap_i64(v311)
	if v313 != 0 {
		v341 = v300
		v342 = v313
		goto L108
	} else {
		goto L168
	}
L168:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+25)))
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+41)))
	if v315 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+42)))
	if v314&int32(1) == int32(0) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	if v314&int32(1) != 0 {
		v380 = v300
		goto L107
	} else {
		goto L182
	}
L172:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+26)))
	if v318 == v323 {
		v380 = v300
		goto L107
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v330 = int32(1)
	if v318&v330 != 0 {
		goto L179
	} else {
		goto L180
	}
L175:
	;
	v325 = int32(1)
	if v318&v325 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v329 = v325
	goto L178
L177:
	;
	v329 = int32(-1)
	goto L178
L178:
	;
	v341 = v300
	v342 = v329
	goto L108
L179:
	;
	v334 = v330
	goto L181
L180:
	;
	v334 = int32(-1)
	goto L181
L181:
	;
	v341 = v300
	v342 = v334
	goto L108
L182:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+26)))
	if v339 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v340 = int32(-1)
	goto L185
L184:
	;
	v340 = int32(1)
	goto L185
L185:
	;
	v341 = v300
	v342 = v340
	goto L108
L186:
	;
	v349 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+58)) = uint8(v349)
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+57)))
	v353 = v351 ^ int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+57)) = uint8(v353)
	v361 = F_make_range(m, v49, v58-int32(-64), v58+int32(48), v349, v349)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L6
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v361
	v364 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+26)) = uint8(v364)
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+25)))
	v368 = v366 ^ v364
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+25)) = uint8(v368)
	v374 = int32(0)
	v376 = F_make_range(m, v49, v58+int32(16), v58+int32(32), v374, v374)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L6
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v376
	v444 = int32(2)
	goto L18
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v36
	v444 = int32(1)
	goto L18
L190:
	;
	goto L191
L191:
	;
	v396 = int32(0)
	if base.B2i32(v385 <= v396)&base.B2i32(v396 <= v221) != 0 {
		v444 = v396
		goto L18
	} else {
		goto L192
	}
L192:
	;
	v402 = int32(0)
	if base.B2i32(v402 < v221)|base.B2i32(v402 < v385) == v402 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v409 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+58)) = uint8(v409)
	v411 = int32(1)
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+57)))
	v414 = v412 ^ v411
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+57)) = uint8(v414)
	v422 = F_make_range(m, v49, v58-int32(-64), v58+int32(48), v409, v409)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L6
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	if v385|v221 < int32(0) {
		goto L17
	} else {
		goto L197
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v422
	v444 = v411
	goto L18
L197:
	;
	v428 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+26)) = uint8(v428)
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+25)))
	v433 = v431 ^ v428
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+25)) = uint8(v433)
	v439 = int32(0)
	v441 = F_make_range(m, v49, v58+int32(16), v58+int32(32), v439, v439)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v441
	v444 = v428
	goto L18
L199:
	;
	F_errmsg_internal(m, int32(_a_F_range_minus_multi_5), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L6
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_range_minus_multi_2), int32(1384), int32(_a_F_range_minus_multi_6))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L6
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	m.G0 = v20 + int32(16)
	return v507
L203:
	;
	v485 = *(*int64)(unsafe.Add(mBase, uint32(v484)))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v484)+16))
	v487 = int64(*(*int32)(unsafe.Add(mBase, uint32(v486)+8)))
	if base.Ui64(v485) < base.Ui64(v487) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v493 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v486+base.I32_wrap_i64(v485)<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v484))) = v485 + int64(1)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v497)+20)) = int32(1)
	v507 = v493
	goto L202
L205:
	;
	goto L206
L206:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L6
	} else {
		goto L207
	}
L207:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v502)+20)) = int32(2)
	v505 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v505)
	v507 = int64(0)
	goto L202
L208:
	;
	F_errmsg_internal(m, int32(_a_F_range_minus_multi_1), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L6
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(_a_F_range_minus_multi_2), int32(1260), int32(_a_F_range_minus_multi_3))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L6
	} else {
		goto L210
	}
L210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v47
	F_errmsg_internal(m, int32(_a_F_range_minus_multi_4), v20)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L6
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_range_minus_multi_2), int32(1272), int32(_a_F_range_minus_multi_3))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L6
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_range_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int64
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	v9 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_check_stack_depth(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v23 = F_get_range_io_data(m, l0, v17, int32(2))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			v25 = F_pq_getmsgbyte(m, v16)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				if v25&int32(9) == int32(0) {
					v32 = F_pq_getmsgint(m, v16, int32(4))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						v34 = F_pq_getmsgbytes(m, v16, v32)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							v37 = v13 + int32(16)
							F_initStringInfo(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
							} else {
								F_appendBinaryStringInfo(m, v37, v34, v32)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int64(0)
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
									v45 = F_ReceiveFunctionCall(m, v23+int32(4), v37, v44, v15)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return int64(0)
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
										F_pfree(m, v47)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int64(0)
										} else {
											v53 = v45
											*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v53
											if v25&int32(17) == int32(0) {
												v60 = F_pq_getmsgint(m, v16, int32(4))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return int64(0)
												} else {
													v62 = F_pq_getmsgbytes(m, v16, v60)
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return int64(0)
													} else {
														F_initStringInfo(m, v13)
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return int64(0)
														} else {
															F_appendBinaryStringInfo(m, v13, v62, v60)
															mBase = m.M
															v67 = m.ExcPending
															if v67 != 0 {
																return int64(0)
															} else {
																v70 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
																v71 = F_ReceiveFunctionCall(m, v23+int32(4), v13, v70, v15)
																mBase = m.M
																v72 = m.ExcPending
																if v72 != 0 {
																	return int64(0)
																} else {
																	v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	F_pfree(m, v73)
																	mBase = m.M
																	v75 = m.ExcPending
																	if v75 != 0 {
																		return int64(0)
																	} else {
																		v78 = v71
																		*(*int64)(unsafe.Add(mBase, uint32(v13))) = v78
																		F_pq_getmsgend(m, v16)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return int64(0)
																		} else {
																			v82 = int32(0)
																			*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v82)
																			v84 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v13)+26)) = uint8(v84)
																			v89 = int32(base.Ui32(v25)>>(uint(int32(3))%32)) & v84
																			*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v89)
																			v94 = int32(base.Ui32(v25)>>(uint(int32(2))%32)) & v84
																			*(*uint8)(unsafe.Add(mBase, uint32(v13)+9)) = uint8(v94)
																			v99 = int32(base.Ui32(v25)>>(uint(int32(4))%32)) & v84
																			*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)) = uint8(v99)
																			v104 = int32(base.Ui32(v25)>>(uint(v84)%32)) & v84
																			*(*uint8)(unsafe.Add(mBase, uint32(v13)+25)) = uint8(v104)
																			v106 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
																			v112 = F_make_range(m, v106, v13+int32(16), v13, v25&v84, v82)
																			mBase = m.M
																			v113 = m.ExcPending
																			if v113 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v13 + int32(32)
																				return base.I64_extend_i32_u(v112)
																			}
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v78 = v9
												*(*int64)(unsafe.Add(mBase, uint32(v13))) = v78
												F_pq_getmsgend(m, v16)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int64(0)
												} else {
													v82 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v82)
													v84 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+26)) = uint8(v84)
													v89 = int32(base.Ui32(v25)>>(uint(int32(3))%32)) & v84
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v89)
													v94 = int32(base.Ui32(v25)>>(uint(int32(2))%32)) & v84
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+9)) = uint8(v94)
													v99 = int32(base.Ui32(v25)>>(uint(int32(4))%32)) & v84
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)) = uint8(v99)
													v104 = int32(base.Ui32(v25)>>(uint(v84)%32)) & v84
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+25)) = uint8(v104)
													v106 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
													v112 = F_make_range(m, v106, v13+int32(16), v13, v25&v84, v82)
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return int64(0)
													} else {
														m.G0 = v13 + int32(32)
														return base.I64_extend_i32_u(v112)
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
					v53 = v9
					*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v53
					if v25&int32(17) == int32(0) {
						v60 = F_pq_getmsgint(m, v16, int32(4))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int64(0)
						} else {
							v62 = F_pq_getmsgbytes(m, v16, v60)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int64(0)
							} else {
								F_initStringInfo(m, v13)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int64(0)
								} else {
									F_appendBinaryStringInfo(m, v13, v62, v60)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int64(0)
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
										v71 = F_ReceiveFunctionCall(m, v23+int32(4), v13, v70, v15)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int64(0)
										} else {
											v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
											F_pfree(m, v73)
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int64(0)
											} else {
												v78 = v71
												*(*int64)(unsafe.Add(mBase, uint32(v13))) = v78
												F_pq_getmsgend(m, v16)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int64(0)
												} else {
													v82 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v82)
													v84 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+26)) = uint8(v84)
													v89 = int32(base.Ui32(v25)>>(uint(int32(3))%32)) & v84
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v89)
													v94 = int32(base.Ui32(v25)>>(uint(int32(2))%32)) & v84
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+9)) = uint8(v94)
													v99 = int32(base.Ui32(v25)>>(uint(int32(4))%32)) & v84
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)) = uint8(v99)
													v104 = int32(base.Ui32(v25)>>(uint(v84)%32)) & v84
													*(*uint8)(unsafe.Add(mBase, uint32(v13)+25)) = uint8(v104)
													v106 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
													v112 = F_make_range(m, v106, v13+int32(16), v13, v25&v84, v82)
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return int64(0)
													} else {
														m.G0 = v13 + int32(32)
														return base.I64_extend_i32_u(v112)
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v78 = v9
						*(*int64)(unsafe.Add(mBase, uint32(v13))) = v78
						F_pq_getmsgend(m, v16)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int64(0)
						} else {
							v82 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v82)
							v84 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v13)+26)) = uint8(v84)
							v89 = int32(base.Ui32(v25)>>(uint(int32(3))%32)) & v84
							*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v89)
							v94 = int32(base.Ui32(v25)>>(uint(int32(2))%32)) & v84
							*(*uint8)(unsafe.Add(mBase, uint32(v13)+9)) = uint8(v94)
							v99 = int32(base.Ui32(v25)>>(uint(int32(4))%32)) & v84
							*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)) = uint8(v99)
							v104 = int32(base.Ui32(v25)>>(uint(v84)%32)) & v84
							*(*uint8)(unsafe.Add(mBase, uint32(v13)+25)) = uint8(v104)
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
							v112 = F_make_range(m, v106, v13+int32(16), v13, v25&v84, v82)
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int64(0)
							} else {
								m.G0 = v13 + int32(32)
								return base.I64_extend_i32_u(v112)
							}
						}
					}
				}
			}
		}
	}
}
func F_range_table_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v5 = int32(0)
	if l0 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v20 = v5
	goto L7
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v20<<(uint(int32(2))%32))))
	v27 = F_range_table_entry_walker_impl(m, v26, l1, l2, l3)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return v27
L9:
	;
	return int32(0)
L10:
	;
	if v27 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = v20 + int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v34 < v35 {
		v20 = v34
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L8
L14:
	;
	goto L13
}
func F_range_upper_inf(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14367(m, l0, int64(4))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
