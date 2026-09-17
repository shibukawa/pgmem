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
	var v12 int32
	_ = v12
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
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(57), l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v16 = v14 + v15
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+119)))
			v19 = v17 - int32(112)
			if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v19))|base.B2i32(int32(1)<<(uint(v19)%32)&int32(69) == int32(0)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v82
						F_errmsg(m, int32(_a_F_RangeVarCallbackForRenameRule_0), v9)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16)+119)))
							F_errdetail_relkind_not_supported(m, v87)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_RangeVarCallbackForRenameRule_1), int32(774), int32(_a_F_RangeVarCallbackForRenameRule_2))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
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
				v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarCallbackForRenameRule[0])))
				if v30 == int32(0) {
					v34 = int32(1)
					if base.Ui32(l1) < base.Ui32(int32(_a_F_RangeVarCallbackForRenameRule_3)) {
						v42 = v34
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
						if v37 == int32(99) {
							v42 = v34
						} else {
							v40 = F_isTempToastNamespace(m, v37)
							mBase = m.M
							v42 = v40
						}
					}
					if v42 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v102
								F_errmsg(m, int32(_a_F_RangeVarCallbackForRenameRule_4), v9+int32(16))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_RangeVarCallbackForRenameRule_1), int32(780), int32(_a_F_RangeVarCallbackForRenameRule_2))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
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
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForRenameRule[1]))
						v46 = F_object_ownercheck(m, int32(1259), l1, v45)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							if v46 == int32(0) {
								v51 = F_get_rel_relkind(m, l1)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									switch v51 - int32(73) {
									case 0, 32:
										v62 = int32(20)
										v64 = v62
									default:
										v62 = int32(41)
										v64 = v62
									case 10:
										v64 = int32(37)
									case 29:
										v64 = int32(18)
									case 36:
										v64 = int32(23)
									case 45:
										v64 = int32(51)
									}
									v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									F_aclcheck_error(m, int32(2), v64, v65)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v12)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v12)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForRenameRule[1]))
					v46 = F_object_ownercheck(m, int32(1259), l1, v45)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						if v46 == int32(0) {
							v51 = F_get_rel_relkind(m, l1)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								switch v51 - int32(73) {
								case 0, 32:
									v62 = int32(20)
									v64 = v62
								default:
									v62 = int32(41)
									v64 = v62
								case 10:
									v64 = int32(37)
								case 29:
									v64 = int32(18)
								case 36:
									v64 = int32(23)
								case 45:
									v64 = int32(51)
								}
								v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								F_aclcheck_error(m, int32(2), v64, v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									F_ReleaseCatCache(m, v12)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						} else {
							F_ReleaseCatCache(m, v12)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l1 != 0 {
		v10 = F_SearchSysCache1(m, int32(57), l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			if v10 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg_internal(m, int32(_a_F_RangeVarCallbackOwnsRelation_0), v7)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_RangeVarCallbackOwnsRelation_1), int32(_a_F_RangeVarCallbackOwnsRelation_2), int32(_a_F_RangeVarCallbackOwnsRelation_3))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackOwnsRelation[0]))
				v17 = F_object_ownercheck(m, int32(1259), l1, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					if v17 == int32(0) {
						v22 = F_get_rel_relkind(m, l1)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							switch v22 - int32(73) {
							case 0, 32:
								v33 = int32(20)
								v35 = v33
							default:
								v33 = int32(41)
								v35 = v33
							case 10:
								v35 = int32(37)
							case 29:
								v35 = int32(18)
							case 36:
								v35 = int32(23)
							case 45:
								v35 = int32(51)
							}
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							F_aclcheck_error(m, int32(2), v35, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarCallbackOwnsRelation[1])))
								if v40 == int32(0) {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
									v47 = int32(1)
									if base.Ui32(l1) < base.Ui32(int32(_a_F_RangeVarCallbackOwnsRelation_4)) {
										v55 = v47
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v43+v44)+68))
										if v50 == int32(99) {
											v55 = v47
										} else {
											v53 = F_isTempToastNamespace(m, v50)
											mBase = m.M
											v55 = v53
										}
									}
									if v55 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											F_errcode(m, int32(16797828))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v84
												F_errmsg(m, int32(_a_F_RangeVarCallbackOwnsRelation_5), v7+int32(16))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_RangeVarCallbackOwnsRelation_1), int32(_a_F_RangeVarCallbackOwnsRelation_6), int32(_a_F_RangeVarCallbackOwnsRelation_3))
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
									} else {
										F_ReleaseCatCache(m, v10)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											m.G0 = v7 + int32(32)
											return
										}
									}
								} else {
									F_ReleaseCatCache(m, v10)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										m.G0 = v7 + int32(32)
										return
									}
								}
							}
						}
					} else {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarCallbackOwnsRelation[1])))
						if v40 == int32(0) {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
							v47 = int32(1)
							if base.Ui32(l1) < base.Ui32(int32(_a_F_RangeVarCallbackOwnsRelation_4)) {
								v55 = v47
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v43+v44)+68))
								if v50 == int32(99) {
									v55 = v47
								} else {
									v53 = F_isTempToastNamespace(m, v50)
									mBase = m.M
									v55 = v53
								}
							}
							if v55 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v84
										F_errmsg(m, int32(_a_F_RangeVarCallbackOwnsRelation_5), v7+int32(16))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_RangeVarCallbackOwnsRelation_1), int32(_a_F_RangeVarCallbackOwnsRelation_6), int32(_a_F_RangeVarCallbackOwnsRelation_3))
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
							} else {
								F_ReleaseCatCache(m, v10)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									m.G0 = v7 + int32(32)
									return
								}
							}
						} else {
							F_ReleaseCatCache(m, v10)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
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
	return v147
L2:
	;
	F_AccessTempTableNamespace(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L47
	}
L3:
	;
	v142 = int32(0)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
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
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v94 == int32(116) {
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
	v73 = int32(0)
	v76 = F_GetSysCacheOid(m, int32(37), v42, v73, v73, v73)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	if v76 != 0 {
		v147 = v76
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v42
	F_errmsg(m, int32(_a_F_RangeVarGetCreationNamespace_1), v7)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_RangeVarGetCreationNamespace_2), int32(3547), int32(_a_F_RangeVarGetCreationNamespace_3))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
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
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarGetCreationNamespace[2])))
	if v101 != 0 {
		v142 = int32(1)
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetCreationNamespace[3]))
	if v103 != 0 {
		v147 = v103
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(_a_F_RangeVarGetCreationNamespace_4), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_RangeVarGetCreationNamespace_2), int32(704), int32(_a_F_RangeVarGetCreationNamespace_5))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
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
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v127 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v127
	F_errmsg(m, int32(_a_F_RangeVarGetCreationNamespace_6), v7+int32(16))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_RangeVarGetCreationNamespace_2), int32(668), int32(_a_F_RangeVarGetCreationNamespace_5))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetCreationNamespace[4]))
	v147 = v146
	goto L1
}
func F_make_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = F_range_serialize(m, l0, l1, l2, l3, l4)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l4 == int32(0) {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
			if v22 == int32(0) {
				v65 = v11
				m.G0 = v9 + int32(32)
				return v65
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v29 = int32(1)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(base.Ui32(v25)>>(uint(int32(2))%32))-v29))))
				if v31&v29 != 0 {
					v65 = v11
					m.G0 = v9 + int32(32)
					return v65
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l4
					v37 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v37)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v37)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v11
					v42 = int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v9)+22)) = uint16(v42)
					v45 = l0 + int32(240)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v45
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
					v50 = m.T0[v49].(func(*base.Module, int32) int32)(m, v9+int32(4))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						if l4 == int32(0) {
							v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
							if v59 == int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v74
									F_errmsg_internal(m, int32(_a_F_make_range_0), v9)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_make_range_1), int32(2048), int32(_a_F_make_range_2))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v62 = F_pg_detoast_datum(m, v50)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v65 = v62
									m.G0 = v9 + int32(32)
									return v65
								}
							}
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
							if v54 != int32(447) {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
								if v59 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v74
										F_errmsg_internal(m, int32(_a_F_make_range_0), v9)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_make_range_1), int32(2048), int32(_a_F_make_range_2))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v62 = F_pg_detoast_datum(m, v50)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v65 = v62
										m.G0 = v9 + int32(32)
										return v65
									}
								}
							} else {
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
								if v58 != 0 {
									v65 = int32(0)
									m.G0 = v9 + int32(32)
									return v65
								} else {
									v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
									if v59 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v74
											F_errmsg_internal(m, int32(_a_F_make_range_0), v9)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_make_range_1), int32(2048), int32(_a_F_make_range_2))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v62 = F_pg_detoast_datum(m, v50)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v65 = v62
											m.G0 = v9 + int32(32)
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
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
			if v17 != int32(447) {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
				if v22 == int32(0) {
					v65 = v11
					m.G0 = v9 + int32(32)
					return v65
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v29 = int32(1)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(base.Ui32(v25)>>(uint(int32(2))%32))-v29))))
					if v31&v29 != 0 {
						v65 = v11
						m.G0 = v9 + int32(32)
						return v65
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l4
						v37 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v37)
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v37)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v11
						v42 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v9)+22)) = uint16(v42)
						v45 = l0 + int32(240)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v45
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
						v50 = m.T0[v49].(func(*base.Module, int32) int32)(m, v9+int32(4))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							if l4 == int32(0) {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
								if v59 == int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v74
										F_errmsg_internal(m, int32(_a_F_make_range_0), v9)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_make_range_1), int32(2048), int32(_a_F_make_range_2))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v62 = F_pg_detoast_datum(m, v50)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v65 = v62
										m.G0 = v9 + int32(32)
										return v65
									}
								}
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								if v54 != int32(447) {
									v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
									if v59 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v74
											F_errmsg_internal(m, int32(_a_F_make_range_0), v9)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_make_range_1), int32(2048), int32(_a_F_make_range_2))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v62 = F_pg_detoast_datum(m, v50)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v65 = v62
											m.G0 = v9 + int32(32)
											return v65
										}
									}
								} else {
									v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
									if v58 != 0 {
										v65 = int32(0)
										m.G0 = v9 + int32(32)
										return v65
									} else {
										v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
										if v59 == int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
												*(*int32)(unsafe.Add(mBase, uint32(v9))) = v74
												F_errmsg_internal(m, int32(_a_F_make_range_0), v9)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_make_range_1), int32(2048), int32(_a_F_make_range_2))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v62 = F_pg_detoast_datum(m, v50)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												v65 = v62
												m.G0 = v9 + int32(32)
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
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
				if v21 != 0 {
					v65 = int32(0)
					m.G0 = v9 + int32(32)
					return v65
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
					if v22 == int32(0) {
						v65 = v11
						m.G0 = v9 + int32(32)
						return v65
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						v29 = int32(1)
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(base.Ui32(v25)>>(uint(int32(2))%32))-v29))))
						if v31&v29 != 0 {
							v65 = v11
							m.G0 = v9 + int32(32)
							return v65
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l4
							v37 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v37)
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v37)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v11
							v42 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v9)+22)) = uint16(v42)
							v45 = l0 + int32(240)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v45
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
							v50 = m.T0[v49].(func(*base.Module, int32) int32)(m, v9+int32(4))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								if l4 == int32(0) {
									v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
									if v59 == int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v74
											F_errmsg_internal(m, int32(_a_F_make_range_0), v9)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_make_range_1), int32(2048), int32(_a_F_make_range_2))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v62 = F_pg_detoast_datum(m, v50)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v65 = v62
											m.G0 = v9 + int32(32)
											return v65
										}
									}
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
									if v54 != int32(447) {
										v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
										if v59 == int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
												*(*int32)(unsafe.Add(mBase, uint32(v9))) = v74
												F_errmsg_internal(m, int32(_a_F_make_range_0), v9)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_make_range_1), int32(2048), int32(_a_F_make_range_2))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										} else {
											v62 = F_pg_detoast_datum(m, v50)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												v65 = v62
												m.G0 = v9 + int32(32)
												return v65
											}
										}
									} else {
										v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
										if v58 != 0 {
											v65 = int32(0)
											m.G0 = v9 + int32(32)
											return v65
										} else {
											v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
											if v59 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = v74
													F_errmsg_internal(m, int32(_a_F_make_range_0), v9)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_make_range_1), int32(2048), int32(_a_F_make_range_2))
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v62 = F_pg_detoast_datum(m, v50)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int32(0)
												} else {
													v65 = v62
													m.G0 = v9 + int32(32)
													return v65
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
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
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
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v317 != 0 {
		goto L110
	} else {
		goto L111
	}
L51:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_range_[0]))
	switch v163 - int32(1) {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	default:
		goto L58
	}
L52:
	;
	return v133
L53:
	;
	if base.B2i32(v213 == v157)|base.B2i32(base.Ui32(l1) <= base.Ui32(v213))&base.B2i32(base.Ui32(v213) <= base.Ui32(l2)) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L54:
	;
	v213 = v209
	goto L53
L55:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_range_[1]))
	if base.Ui32(int32(127)) < base.Ui32(v157) {
		goto L69
	} else {
		goto L70
	}
L56:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_range_[1]))
	if base.Ui32(int32(127)) < base.Ui32(v157) {
		goto L66
	} else {
		goto L67
	}
L57:
	;
	if base.Ui32(v157) <= base.Ui32(int32(127)) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	if base.Ui32(int32(127)) < base.Ui32(v157) {
		v209 = v157
		goto L54
	} else {
		goto L59
	}
L59:
	;
	v168 = F_pg_tolower(m, v157)
	mBase = m.M
	v213 = v168
	goto L53
L60:
	;
	v180 = v157<<(uint(int32(2))%32) + int32(_a_F_range__1)
	goto L62
L61:
	;
	v175 = F_case_index(m, v157)
	mBase = m.M
	v180 = v175<<(uint(int32(2))%32) + int32(_a_F_range__2)
	goto L62
L62:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	if v181 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v182 = v181
	goto L65
L64:
	;
	v182 = v157
	goto L65
L65:
	;
	v213 = v182
	goto L53
L66:
	;
	v194 = F_towlower(m, v157)
	mBase = m.M
	v213 = v194
	goto L53
L67:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+4)))
	if v187&int32(1) == int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v192 = F_pg_tolower(m, v157)
	mBase = m.M
	v213 = v192
	goto L53
L69:
	;
	if base.Ui32(int32(255)) < base.Ui32(v157) {
		v209 = v157
		goto L54
	} else {
		goto L72
	}
L70:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+4)))
	if v199&int32(1) == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v204 = F_pg_tolower(m, v157)
	mBase = m.M
	v213 = v204
	goto L53
L72:
	;
	v208 = F_tolower(m, v157)
	mBase = m.M
	v209 = v208
	goto L54
L73:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v222 <= v221 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_range_[0]))
	switch v235 - int32(1) {
	case 0:
		goto L83
	case 1:
		goto L82
	case 2:
		goto L81
	default:
		goto L84
	}
L76:
	;
	goto L50
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v221 + int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v227+v221<<(uint(int32(2))%32)))) = v213
	goto L75
L79:
	;
	if base.B2i32(v285 == v157)|base.B2i32(base.Ui32(l1) <= base.Ui32(v285))&base.B2i32(base.Ui32(v285) <= base.Ui32(l2)) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L80:
	;
	v285 = v281
	goto L79
L81:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_range_[1]))
	if base.Ui32(int32(127)) < base.Ui32(v157) {
		goto L95
	} else {
		goto L96
	}
L82:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_range_[1]))
	if base.Ui32(int32(127)) < base.Ui32(v157) {
		goto L92
	} else {
		goto L93
	}
L83:
	;
	if base.Ui32(v157) <= base.Ui32(int32(127)) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if base.Ui32(int32(127)) < base.Ui32(v157) {
		v281 = v157
		goto L80
	} else {
		goto L85
	}
L85:
	;
	v240 = F_pg_toupper(m, v157)
	mBase = m.M
	v285 = v240
	goto L79
L86:
	;
	v252 = v157<<(uint(int32(2))%32) + int32(_a_F_range__3)
	goto L88
L87:
	;
	v247 = F_case_index(m, v157)
	mBase = m.M
	v252 = v247<<(uint(int32(2))%32) + int32(_a_F_range__4)
	goto L88
L88:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	if v253 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v254 = v253
	goto L91
L90:
	;
	v254 = v157
	goto L91
L91:
	;
	v285 = v254
	goto L79
L92:
	;
	v266 = F_towupper(m, v157)
	mBase = m.M
	v285 = v266
	goto L79
L93:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+4)))
	if v259&int32(1) == int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v264 = F_pg_toupper(m, v157)
	mBase = m.M
	v285 = v264
	goto L79
L95:
	;
	if base.Ui32(int32(255)) < base.Ui32(v157) {
		v281 = v157
		goto L80
	} else {
		goto L98
	}
L96:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268)+4)))
	if v271&int32(1) == int32(0) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v276 = F_pg_toupper(m, v157)
	mBase = m.M
	v285 = v276
	goto L79
L98:
	;
	v280 = F_toupper(m, v157)
	mBase = m.M
	v281 = v280
	goto L80
L99:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v294 <= v293 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_range_[2]))
	if v306 != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	goto L50
L103:
	;
	goto L104
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v293 + int32(1)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v299+v293<<(uint(int32(2))%32)))) = v285
	goto L101
L105:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L17
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v310 = v157 + int32(1)
	if base.Ui32(v310) <= base.Ui32(l2) {
		v157 = v310
		goto L51
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	goto L52
L110:
	;
	v319 = v317
	goto L112
L111:
	;
	v319 = int32(19)
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v319
	return int32(0)
}
func F_range_adjacent_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v9 == v10 {
		F_range_deserialize(m, l0, l1, v7+int32(72), v7+int32(56), v7+int32(47))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_range_deserialize(m, l0, l2, v7-int32(-64), v7+int32(48), v7+int32(46))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(0)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+47)))
				if v31 != 0 {
					v56 = v30
					m.G0 = v7 + int32(80)
					return v56
				} else {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+46)))
					if v32&int32(1) != 0 {
						v56 = v30
						m.G0 = v7 + int32(80)
						return v56
					} else {
						v35 = *(*int64)(unsafe.Add(mBase, uint32(v7)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v35
						v37 = *(*int64)(unsafe.Add(mBase, uint32(v7)+64))
						*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v37
						v44 = F_bounds_adjacent(m, l0, v7+int32(32), v7+int32(24))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							if v44 != 0 {
								v56 = int32(1)
								m.G0 = v7 + int32(80)
								return v56
							} else {
								v46 = *(*int64)(unsafe.Add(mBase, uint32(v7)+48))
								*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v46
								v48 = *(*int64)(unsafe.Add(mBase, uint32(v7)+72))
								*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v48
								v54 = F_bounds_adjacent(m, l0, v7+int32(16), v7+int32(8))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = v54
									m.G0 = v7 + int32(80)
									return v56
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
		v64 = m.ExcPending
		if v64 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_range_adjacent_internal_0), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_adjacent_internal_1), int32(811), int32(_a_F_range_adjacent_internal_2))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
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
func F_range_agg_finalfn(m *base.Module, l0 int32) int32 {
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
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
	v144 = m.ExcPending
	if v144 != 0 {
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
	switch v19 - int32(429) {
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
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+368))
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
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
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
	v131 = m.ExcPending
	if v131 != 0 {
		goto L28
	} else {
		goto L46
	}
L19:
	;
	m.G0 = v11 + int32(16)
	return v123
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
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v123 = int32(0)
	goto L19
L24:
	;
	goto L23
L25:
	;
	v56 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
	v123 = int32(0)
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
	return int32(0)
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
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v76)+296))
	v113 = F_make_multirange(m, v60, v112, v53, v79)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L28
	} else {
		goto L45
	}
L41:
	;
	v93 = v84 << (uint(int32(2)) % 32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v95+v93)))
	v98 = F_pg_detoast_datum(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L28
	} else {
		goto L43
	}
L42:
	;
	goto L40
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79+v93))) = v98
	v102 = v84 + int32(1)
	if v102 != v53 {
		v84 = v102
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v123 = v113
	goto L19
L46:
	;
	F_errmsg_internal(m, int32(_a_F_range_agg_finalfn_1), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L28
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_range_agg_finalfn_2), int32(1384), int32(_a_F_range_agg_finalfn_3))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
	v148 = m.ExcPending
	if v148 != 0 {
		goto L28
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_range_agg_finalfn_2), int32(558), int32(_a_F_range_agg_finalfn_5))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
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
func F_range_agg_transfn(m *base.Module, l0 int32) int32 {
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
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
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
		switch v15 - int32(429) {
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
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+368))
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
			return int32(0)
		} else {
			v47 = F_type_is_range(m, v43)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				if v47 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_range_agg_transfn_0), int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_range_agg_transfn_1), int32(1352), int32(_a_F_range_agg_transfn_2))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
					if v51 == int32(1) {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						v56 = F_initArrayResult(m, v43, v54, int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v59 = v56
							v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
							if v60 == int32(0) {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
								v66 = F_accumArrayResult(m, v59, v63, int32(0), v43, v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(16)
									return v59
								}
							} else {
								m.G0 = v7 + int32(16)
								return v59
							}
						}
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v59 = v58
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v60 == int32(0) {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							v66 = F_accumArrayResult(m, v59, v63, int32(0), v43, v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(16)
								return v59
							}
						} else {
							m.G0 = v7 + int32(16)
							return v59
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_range_agg_transfn_3), int32(0))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_agg_transfn_1), int32(1348), int32(_a_F_range_agg_transfn_2))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
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
func F_range_contained_by_multirange(m *base.Module, l0 int32) int32 {
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
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
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
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_contained_by_multirange_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_contained_by_multirange_1), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_contained_by_multirange_2), int32(558), int32(_a_F_range_contained_by_multirange_3))
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
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = F_multirange_contains_range_internal(m, v33, v17, v12)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_contained_by_multirange_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_contained_by_multirange_1), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_contained_by_multirange_2), int32(558), int32(_a_F_range_contained_by_multirange_3))
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
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = F_multirange_contains_range_internal(m, v33, v17, v12)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_range_contains_multirange(m *base.Module, l0 int32) int32 {
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
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
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
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_contains_multirange_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_contains_multirange_1), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_contains_multirange_2), int32(558), int32(_a_F_range_contains_multirange_3))
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
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = F_range_contains_multirange_internal(m, v33, v12, v17)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_contains_multirange_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_contains_multirange_1), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_contains_multirange_2), int32(558), int32(_a_F_range_contains_multirange_3))
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
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = F_range_contains_multirange_internal(m, v33, v12, v17)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v34
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v10 == v11 {
		F_range_deserialize(m, l0, l1, v8+int32(40), v8+int32(24), v8+int32(15))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_range_deserialize(m, l0, l2, v8+int32(32), v8+int32(16), v8+int32(14))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = int32(1)
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
				if v32&v31&base.B2i32(v35 == v31) != 0 {
					v106 = v31
					m.G0 = v8 + int32(48)
					return v106 & int32(1)
				} else {
					v39 = int32(0)
					if v32 != v35 {
						v106 = v39
						m.G0 = v8 + int32(48)
						return v106 & int32(1)
					} else {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+44)))
						if v42 == int32(1) {
							if v41&int32(1) == int32(0) {
								v106 = v39
								m.G0 = v8 + int32(48)
								return v106 & int32(1)
							} else {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+46)))
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
								if v49 == v50 {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
									if v76 == int32(1) {
										if v75&int32(1) == int32(0) {
											v106 = v39
										} else {
											v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
											v106 = base.B2i32(v83 == v84)
										}
										m.G0 = v8 + int32(48)
										return v106 & int32(1)
									} else {
										if v75&int32(1) != 0 {
											v106 = v39
											m.G0 = v8 + int32(48)
											return v106 & int32(1)
										} else {
											v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
											v93 = F_FunctionCall2Coll(m, l0+int32(212), v90, v91, v92)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												if v93 != 0 {
													v106 = v39
												} else {
													v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+29)))
													v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)))
													if (v95|v96)&int32(1) != 0 {
														v106 = v95 & v96
													} else {
														v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
														v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
														v106 = base.B2i32(v101 == v102)
													}
												}
												m.G0 = v8 + int32(48)
												return v106 & int32(1)
											}
										}
									}
								} else {
									v106 = v39
									m.G0 = v8 + int32(48)
									return v106 & int32(1)
								}
							}
						} else {
							if v41&int32(1) != 0 {
								v106 = v39
								m.G0 = v8 + int32(48)
								return v106 & int32(1)
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
								v59 = F_FunctionCall2Coll(m, l0+int32(212), v56, v57, v58)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									if v59 != 0 {
										v106 = v39
										m.G0 = v8 + int32(48)
										return v106 & int32(1)
									} else {
										v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
										v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+45)))
										if v62 == int32(0) {
											if v61&int32(1) != 0 {
												v106 = v39
												m.G0 = v8 + int32(48)
												return v106 & int32(1)
											} else {
												v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+46)))
												v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
												if v67 == v68 {
													v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
													v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
													if v76 == int32(1) {
														if v75&int32(1) == int32(0) {
															v106 = v39
														} else {
															v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
															v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
															v106 = base.B2i32(v83 == v84)
														}
														m.G0 = v8 + int32(48)
														return v106 & int32(1)
													} else {
														if v75&int32(1) != 0 {
															v106 = v39
															m.G0 = v8 + int32(48)
															return v106 & int32(1)
														} else {
															v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
															v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
															v92 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
															v93 = F_FunctionCall2Coll(m, l0+int32(212), v90, v91, v92)
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int32(0)
															} else {
																if v93 != 0 {
																	v106 = v39
																} else {
																	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+29)))
																	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)))
																	if (v95|v96)&int32(1) != 0 {
																		v106 = v95 & v96
																	} else {
																		v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
																		v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
																		v106 = base.B2i32(v101 == v102)
																	}
																}
																m.G0 = v8 + int32(48)
																return v106 & int32(1)
															}
														}
													}
												} else {
													v106 = v39
													m.G0 = v8 + int32(48)
													return v106 & int32(1)
												}
											}
										} else {
											if v61&int32(1) == int32(0) {
												v106 = v39
												m.G0 = v8 + int32(48)
												return v106 & int32(1)
											} else {
												v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
												v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
												if v76 == int32(1) {
													if v75&int32(1) == int32(0) {
														v106 = v39
													} else {
														v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
														v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
														v106 = base.B2i32(v83 == v84)
													}
													m.G0 = v8 + int32(48)
													return v106 & int32(1)
												} else {
													if v75&int32(1) != 0 {
														v106 = v39
														m.G0 = v8 + int32(48)
														return v106 & int32(1)
													} else {
														v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
														v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
														v92 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
														v93 = F_FunctionCall2Coll(m, l0+int32(212), v90, v91, v92)
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															if v93 != 0 {
																v106 = v39
															} else {
																v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+29)))
																v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)))
																if (v95|v96)&int32(1) != 0 {
																	v106 = v95 & v96
																} else {
																	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
																	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
																	v106 = base.B2i32(v101 == v102)
																}
															}
															m.G0 = v8 + int32(48)
															return v106 & int32(1)
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
		v116 = m.ExcPending
		if v116 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_range_eq_internal_0), int32(0))
			mBase = m.M
			v120 = m.ExcPending
			if v120 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_eq_internal_1), int32(586), int32(_a_F_range_eq_internal_2))
				mBase = m.M
				v125 = m.ExcPending
				if v125 != 0 {
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
							F_errfinish(m, int32(_a_F_range_get_typcache_1), int32(1776), int32(_a_F_range_get_typcache_2))
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
						F_errfinish(m, int32(_a_F_range_get_typcache_1), int32(1776), int32(_a_F_range_get_typcache_2))
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
func F_range_intersect_agg_transfn(m *base.Module, l0 int32) int32 {
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
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
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
		switch v15 - int32(429) {
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
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)+368))
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
			return int32(0)
		} else {
			v47 = F_type_is_range(m, v43)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				if v47 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_range_intersect_agg_transfn_0), int32(0))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_range_intersect_agg_transfn_1), int32(1234), int32(_a_F_range_intersect_agg_transfn_2))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
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
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v65 = F_pg_detoast_datum(m, v64)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v68 = F_pg_detoast_datum(m, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									v70 = F_range_intersect_internal(m, v63, v65, v68)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										m.G0 = v7 + int32(16)
										return v70
									}
								}
							}
						} else {
							v56 = F_lookup_type_cache(m, v43, int32(2048))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+200))
								if v58 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v43
										F_errmsg_internal(m, int32(_a_F_range_intersect_agg_transfn_3), v7)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_range_intersect_agg_transfn_1), int32(1776), int32(_a_F_range_intersect_agg_transfn_4))
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int32(0)
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
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v65 = F_pg_detoast_datum(m, v64)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										v68 = F_pg_detoast_datum(m, v67)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v70 = F_range_intersect_internal(m, v63, v65, v68)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												m.G0 = v7 + int32(16)
												return v70
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
							return int32(0)
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+200))
							if v58 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v43
									F_errmsg_internal(m, int32(_a_F_range_intersect_agg_transfn_3), v7)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_range_intersect_agg_transfn_1), int32(1776), int32(_a_F_range_intersect_agg_transfn_4))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
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
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v65 = F_pg_detoast_datum(m, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									v68 = F_pg_detoast_datum(m, v67)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v70 = F_range_intersect_internal(m, v63, v65, v68)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											m.G0 = v7 + int32(16)
											return v70
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
		v79 = m.ExcPending
		if v79 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_range_intersect_agg_transfn_5), int32(0))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_intersect_agg_transfn_1), int32(1230), int32(_a_F_range_intersect_agg_transfn_2))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
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
func F_range_lower_inc(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13985(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
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
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	F_range_deserialize(m, l0, l1, v9+int32(-24), v9+int32(-40), v9+int32(-49))
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
	F_range_deserialize(m, l0, l2, v9+int32(-32), v9+int32(-48), v9+int32(-50))
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
		v370 = l1
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L194
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L190
	}
L6:
	;
	m.G0 = v11 - int32(-64)
	return v370
L7:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)))
	if v32&int32(1) != 0 {
		v370 = l1
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
	if v36 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)))
	if v173 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L10:
	;
	v170 = v164
	v171 = v167
	v172 = int32(0)
	goto L9
L11:
	;
	v118 = int32(1)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
	if v119 == v118 {
		goto L63
	} else {
		goto L64
	}
L12:
	;
	v111 = int32(1)
	if v69&v111 != 0 {
		goto L60
	} else {
		goto L61
	}
L13:
	;
	v95 = int32(1)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
	if v96 == v95 {
		goto L50
	} else {
		goto L51
	}
L14:
	;
	v90 = int32(1)
	if v39&v90 != 0 {
		goto L47
	} else {
		goto L48
	}
L15:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+46)))
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
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)))
	if v45 == v39 {
		v94 = int32(0)
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
	v94 = v51
	goto L13
L23:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)))
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
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
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
	v117 = v57
	goto L11
L29:
	;
	if v63 != 0 {
		v117 = v63
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+37)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+45)))
	if v66 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+46)))
	if v65&int32(1) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v65&int32(1) != 0 {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)))
	if v69 != v74 {
		goto L12
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v77 = int32(1)
	if v69&v77 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v117 = int32(0)
	goto L11
L38:
	;
	v81 = v77
	goto L40
L39:
	;
	v81 = int32(-1)
	goto L40
L40:
	;
	v117 = v81
	goto L11
L41:
	;
	v117 = int32(0)
	goto L11
L42:
	;
	goto L43
L43:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)))
	if v87 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v88 = int32(-1)
	goto L46
L45:
	;
	v88 = int32(1)
	goto L46
L46:
	;
	v117 = v88
	goto L11
L47:
	;
	v93 = int32(-1)
	goto L49
L48:
	;
	v93 = v90
	goto L49
L49:
	;
	v94 = v93
	goto L13
L50:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	if v99 == v39 {
		v170 = v94
		v171 = int32(0)
		v172 = v95
		goto L9
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v107 = int32(1)
	if v39&v107 != 0 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v102 = int32(1)
	if v39&v102 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v105 = int32(-1)
	goto L56
L55:
	;
	v105 = v102
	goto L56
L56:
	;
	v170 = v94
	v171 = v105
	v172 = v95
	goto L9
L57:
	;
	v110 = int32(-1)
	goto L59
L58:
	;
	v110 = v107
	goto L59
L59:
	;
	v164 = v94
	v167 = v110
	goto L10
L60:
	;
	v115 = v111
	goto L62
L61:
	;
	v115 = int32(-1)
	goto L62
L62:
	;
	v117 = v115
	goto L11
L63:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	if v124 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v126 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v132 = F_FunctionCall2Coll(m, l0+int32(212), v129, v130, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L69
	}
L66:
	;
	v125 = int32(1)
	goto L68
L67:
	;
	v125 = int32(-1)
	goto L68
L68:
	;
	v170 = v117
	v171 = v125
	v172 = v118
	goto L9
L69:
	;
	if v132 != 0 {
		v170 = v117
		v171 = v132
		v172 = v126
		goto L9
	} else {
		goto L70
	}
L70:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+45)))
	if v135 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+46)))
	if v134&int32(1) == int32(0) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if v134&int32(1) != 0 {
		goto L86
	} else {
		goto L87
	}
L74:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	if v143 == v138 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v151 = int32(1)
	if v138&v151 != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v170 = v117
	v171 = int32(0)
	v172 = v126
	goto L9
L78:
	;
	goto L79
L79:
	;
	v146 = int32(1)
	if v138&v146 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v150 = v146
	goto L82
L81:
	;
	v150 = int32(-1)
	goto L82
L82:
	;
	v164 = v117
	v167 = v150
	goto L10
L83:
	;
	v155 = v151
	goto L85
L84:
	;
	v155 = int32(-1)
	goto L85
L85:
	;
	v170 = v117
	v171 = v155
	v172 = v126
	goto L9
L86:
	;
	v170 = v117
	v171 = int32(0)
	v172 = v126
	goto L9
L87:
	;
	goto L88
L88:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	if v161 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v162 = int32(-1)
	goto L91
L90:
	;
	v162 = int32(1)
	goto L91
L91:
	;
	v164 = v117
	v167 = v162
	goto L10
L92:
	;
	v302 = int32(0)
	if base.B2i32(v300 < v302)|base.B2i32(v302 < v171) != 0 {
		v370 = l1
		goto L6
	} else {
		goto L179
	}
L93:
	;
	if int32(0) <= v170 {
		v299 = v292
		v300 = v293
		goto L92
	} else {
		goto L177
	}
L94:
	;
	if v172 != 0 {
		goto L148
	} else {
		goto L149
	}
L95:
	;
	v245 = int32(1)
	if v206&v245 != 0 {
		goto L145
	} else {
		goto L146
	}
L96:
	;
	if v172 != 0 {
		goto L133
	} else {
		goto L134
	}
L97:
	;
	v227 = int32(1)
	if v176&v227 != 0 {
		goto L130
	} else {
		goto L131
	}
L98:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)))
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
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)))
	if v176 == v182 {
		v231 = int32(0)
		goto L96
	} else {
		goto L102
	}
L102:
	;
	v185 = int32(1)
	if v176&v185 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v188 = int32(-1)
	goto L105
L104:
	;
	v188 = v185
	goto L105
L105:
	;
	v231 = v188
	goto L96
L106:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)))
	if v193 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	goto L108
L108:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v200 = F_FunctionCall2Coll(m, l0+int32(212), v197, v198, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L112
	}
L109:
	;
	v194 = int32(1)
	goto L111
L110:
	;
	v194 = int32(-1)
	goto L111
L111:
	;
	v251 = v194
	goto L94
L112:
	;
	if v200 != 0 {
		v251 = v200
		goto L94
	} else {
		goto L113
	}
L113:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+37)))
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)))
	if v203 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)))
	if v202&int32(1) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	if v202&int32(1) != 0 {
		goto L124
	} else {
		goto L125
	}
L117:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)))
	if v211 != v206 {
		goto L95
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v214 = int32(1)
	if v206&v214 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v251 = int32(0)
	goto L94
L121:
	;
	v218 = v214
	goto L123
L122:
	;
	v218 = int32(-1)
	goto L123
L123:
	;
	v251 = v218
	goto L94
L124:
	;
	v251 = int32(0)
	goto L94
L125:
	;
	goto L126
L126:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)))
	if v224 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v225 = int32(-1)
	goto L129
L128:
	;
	v225 = int32(1)
	goto L129
L129:
	;
	v251 = v225
	goto L94
L130:
	;
	v230 = int32(-1)
	goto L132
L131:
	;
	v230 = v227
	goto L132
L132:
	;
	v231 = v230
	goto L96
L133:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	if v232 == v176 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	v241 = int32(1)
	if v176&v241 != 0 {
		goto L142
	} else {
		goto L143
	}
L136:
	;
	v299 = int32(0)
	v300 = v231
	goto L92
L137:
	;
	goto L138
L138:
	;
	v236 = int32(1)
	if v176&v236 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v239 = int32(-1)
	goto L141
L140:
	;
	v239 = v236
	goto L141
L141:
	;
	v292 = v239
	v293 = v231
	goto L93
L142:
	;
	v244 = int32(-1)
	goto L144
L143:
	;
	v244 = v241
	goto L144
L144:
	;
	v292 = v244
	v293 = v231
	goto L93
L145:
	;
	v249 = v245
	goto L147
L146:
	;
	v249 = int32(-1)
	goto L147
L147:
	;
	v251 = v249
	goto L94
L148:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	if v254 != 0 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v261 = F_FunctionCall2Coll(m, l0+int32(212), v258, v259, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L154
	}
L151:
	;
	v255 = int32(1)
	goto L153
L152:
	;
	v255 = int32(-1)
	goto L153
L153:
	;
	v292 = v255
	v293 = v251
	goto L93
L154:
	;
	if v261 != 0 {
		v292 = v261
		v293 = v251
		goto L93
	} else {
		goto L155
	}
L155:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)))
	if v264 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)))
	if v263&int32(1) == int32(0) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L158
L158:
	;
	if v263&int32(1) != 0 {
		goto L171
	} else {
		goto L172
	}
L159:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	if v272 == v267 {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	goto L161
L161:
	;
	v280 = int32(1)
	if v267&v280 != 0 {
		goto L168
	} else {
		goto L169
	}
L162:
	;
	v299 = int32(0)
	v300 = v251
	goto L92
L163:
	;
	goto L164
L164:
	;
	v275 = int32(1)
	if v267&v275 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v279 = v275
	goto L167
L166:
	;
	v279 = int32(-1)
	goto L167
L167:
	;
	v292 = v279
	v293 = v251
	goto L93
L168:
	;
	v284 = v280
	goto L170
L169:
	;
	v284 = int32(-1)
	goto L170
L170:
	;
	v292 = v284
	v293 = v251
	goto L93
L171:
	;
	v299 = int32(0)
	v300 = v251
	goto L92
L172:
	;
	goto L173
L173:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	if v290 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v291 = int32(-1)
	goto L176
L175:
	;
	v291 = int32(1)
	goto L176
L176:
	;
	v292 = v291
	v293 = v251
	goto L93
L177:
	;
	if int32(0) < v292 {
		goto L5
	} else {
		goto L178
	}
L178:
	;
	v299 = v292
	v300 = v293
	goto L92
L179:
	;
	v307 = int32(0)
	if base.B2i32(v170 < v307)|base.B2i32(v307 < v299) == v307 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v314 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+62)) = uint8(v314)
	v316 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+60)) = uint16(v316)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v11)+51)) = v316
	v330 = F_make_range(m, l0, v9+int32(-8), v9+int32(-16), v314, v316)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v332 = int32(0)
	if base.B2i32(v332 < v170)|base.B2i32(v332 < v299) == v332 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v370 = v330
	goto L6
L184:
	;
	v339 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)) = uint8(v339)
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+37)))
	v343 = v341 ^ int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+37)) = uint8(v343)
	v351 = F_make_range(m, l0, v9+int32(-24), v9+int32(-32), v339, v339)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	if v299|v170 < int32(0) {
		goto L4
	} else {
		goto L188
	}
L187:
	;
	v370 = v351
	goto L6
L188:
	;
	v356 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v356)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)))
	v360 = v358 ^ v356
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(v360)
	v366 = int32(0)
	v368 = F_make_range(m, l0, v9+int32(-48), v9+int32(-40), v366, v366)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v370 = v368
	goto L6
L190:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_errmsg(m, int32(_a_F_range_minus_internal_3), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_range_minus_internal_1), int32(1023), int32(_a_F_range_minus_internal_2))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
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
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_range_minus_internal_1), int32(1045), int32(_a_F_range_minus_internal_2))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
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
func F_range_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_check_stack_depth(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v22 = F_get_range_io_data(m, l0, v16, int32(2))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_pq_getmsgbyte(m, v15)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24&int32(9) == int32(0) {
					v31 = F_pq_getmsgint(m, v15, int32(4))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = F_pq_getmsgbytes(m, v15, v31)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v36 = v12 + int32(8)
							F_initStringInfo(m, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_appendBinaryStringInfo(m, v36, v33, v31)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
									v44 = F_ReceiveFunctionCall(m, v22+int32(4), v36, v43, v14)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
										F_pfree(m, v46)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return int32(0)
										} else {
											v49 = v44
											*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v49
											if v24&int32(17) == int32(0) {
												v58 = F_pq_getmsgint(m, v15, int32(4))
												mBase = m.M
												v59 = m.ExcPending
												if v59 != 0 {
													return int32(0)
												} else {
													v60 = F_pq_getmsgbytes(m, v15, v58)
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return int32(0)
													} else {
														v63 = v12 + int32(8)
														F_initStringInfo(m, v63)
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return int32(0)
														} else {
															F_appendBinaryStringInfo(m, v63, v60, v58)
															mBase = m.M
															v67 = m.ExcPending
															if v67 != 0 {
																return int32(0)
															} else {
																v70 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
																v71 = F_ReceiveFunctionCall(m, v22+int32(4), v63, v70, v14)
																mBase = m.M
																v72 = m.ExcPending
																if v72 != 0 {
																	return int32(0)
																} else {
																	v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
																	F_pfree(m, v73)
																	mBase = m.M
																	v75 = m.ExcPending
																	if v75 != 0 {
																		return int32(0)
																	} else {
																		v79 = v71
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v79
																		F_pq_getmsgend(m, v15)
																		mBase = m.M
																		v82 = m.ExcPending
																		if v82 != 0 {
																			return int32(0)
																		} else {
																			v83 = int32(0)
																			*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v83)
																			v85 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)) = uint8(v85)
																			v90 = int32(base.Ui32(v24)>>(uint(int32(3))%32)) & v85
																			*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)) = uint8(v90)
																			v95 = int32(base.Ui32(v24)>>(uint(int32(2))%32)) & v85
																			*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v95)
																			v100 = int32(base.Ui32(v24)>>(uint(int32(4))%32)) & v85
																			*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v100)
																			v105 = int32(base.Ui32(v24)>>(uint(v85)%32)) & v85
																			*(*uint8)(unsafe.Add(mBase, uint32(v12)+29)) = uint8(v105)
																			v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
																			v115 = F_make_range(m, v107, v12+int32(24), v12+int32(8), v24&v85, v83)
																			mBase = m.M
																			v116 = m.ExcPending
																			if v116 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v12 + int32(32)
																				return v115
																			}
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v79 = v2
												*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v79
												F_pq_getmsgend(m, v15)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													v83 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v83)
													v85 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)) = uint8(v85)
													v90 = int32(base.Ui32(v24)>>(uint(int32(3))%32)) & v85
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)) = uint8(v90)
													v95 = int32(base.Ui32(v24)>>(uint(int32(2))%32)) & v85
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v95)
													v100 = int32(base.Ui32(v24)>>(uint(int32(4))%32)) & v85
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v100)
													v105 = int32(base.Ui32(v24)>>(uint(v85)%32)) & v85
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+29)) = uint8(v105)
													v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
													v115 = F_make_range(m, v107, v12+int32(24), v12+int32(8), v24&v85, v83)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return int32(0)
													} else {
														m.G0 = v12 + int32(32)
														return v115
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
					v49 = v2
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v49
					if v24&int32(17) == int32(0) {
						v58 = F_pq_getmsgint(m, v15, int32(4))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = F_pq_getmsgbytes(m, v15, v58)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v63 = v12 + int32(8)
								F_initStringInfo(m, v63)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									F_appendBinaryStringInfo(m, v63, v60, v58)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
										v71 = F_ReceiveFunctionCall(m, v22+int32(4), v63, v70, v14)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
											F_pfree(m, v73)
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int32(0)
											} else {
												v79 = v71
												*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v79
												F_pq_getmsgend(m, v15)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													v83 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v83)
													v85 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)) = uint8(v85)
													v90 = int32(base.Ui32(v24)>>(uint(int32(3))%32)) & v85
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)) = uint8(v90)
													v95 = int32(base.Ui32(v24)>>(uint(int32(2))%32)) & v85
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v95)
													v100 = int32(base.Ui32(v24)>>(uint(int32(4))%32)) & v85
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v100)
													v105 = int32(base.Ui32(v24)>>(uint(v85)%32)) & v85
													*(*uint8)(unsafe.Add(mBase, uint32(v12)+29)) = uint8(v105)
													v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
													v115 = F_make_range(m, v107, v12+int32(24), v12+int32(8), v24&v85, v83)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return int32(0)
													} else {
														m.G0 = v12 + int32(32)
														return v115
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v79 = v2
						*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v79
						F_pq_getmsgend(m, v15)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v83 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v83)
							v85 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)) = uint8(v85)
							v90 = int32(base.Ui32(v24)>>(uint(int32(3))%32)) & v85
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)) = uint8(v90)
							v95 = int32(base.Ui32(v24)>>(uint(int32(2))%32)) & v85
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v95)
							v100 = int32(base.Ui32(v24)>>(uint(int32(4))%32)) & v85
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v100)
							v105 = int32(base.Ui32(v24)>>(uint(v85)%32)) & v85
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+29)) = uint8(v105)
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
							v115 = F_make_range(m, v107, v12+int32(24), v12+int32(8), v24&v85, v83)
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								m.G0 = v12 + int32(32)
								return v115
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
func F_range_upper_inf(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13985(m, l0, int32(4))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
