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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
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
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
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
			if base.Ui32(int32(6)) < base.Ui32(v19) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v81
						F_errmsg(m, int32(174315), v9)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							v86 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16)+119)))
							F_errdetail_relkind_not_supported(m, v86)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								F_errfinish(m, int32(525170), int32(774), int32(402760))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
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
				if int32(1)<<(uint(v19)%32)&int32(69) == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v81
							F_errmsg(m, int32(174315), v9)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								v86 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16)+119)))
								F_errdetail_relkind_not_supported(m, v86)
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_errfinish(m, int32(525170), int32(774), int32(402760))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
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
					v29 = int32(*(*uint8)(unsafe.Add(mBase, _consts[456])))
					if v29 == int32(0) {
						v33 = int32(1)
						if base.Ui32(l1) < base.Ui32(int32(12000)) {
							v41 = v33
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
							if v36 == int32(99) {
								v41 = v33
							} else {
								v39 = F_isTempToastNamespace(m, v36)
								mBase = m.M
								v41 = v39
							}
						}
						if v41 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								F_errcode(m, int32(16797828))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v101
									F_errmsg(m, int32(345099), v9+int32(16))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return
									} else {
										F_errfinish(m, int32(525170), int32(780), int32(402760))
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
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
							v44 = *(*int32)(unsafe.Add(mBase, _consts[168]))
							v45 = F_object_ownercheck(m, int32(1259), l1, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								if v45 == int32(0) {
									v50 = F_get_rel_relkind(m, l1)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										switch v50 - int32(73) {
										case 0, 32:
											v63 = int32(20)
										default:
											v61 = int32(41)
											v63 = v61
										case 10:
											v63 = int32(37)
										case 29:
											v61 = int32(18)
											v63 = v61
										case 36:
											v63 = int32(23)
										case 45:
											v63 = int32(51)
										}
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										F_aclcheck_error(m, int32(2), v63, v64)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											F_ReleaseCatCache(m, v12)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
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
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, _consts[168]))
						v45 = F_object_ownercheck(m, int32(1259), l1, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							if v45 == int32(0) {
								v50 = F_get_rel_relkind(m, l1)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									switch v50 - int32(73) {
									case 0, 32:
										v63 = int32(20)
									default:
										v61 = int32(41)
										v63 = v61
									case 10:
										v63 = int32(37)
									case 29:
										v61 = int32(18)
										v63 = v61
									case 36:
										v63 = int32(23)
									case 45:
										v63 = int32(51)
									}
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									F_aclcheck_error(m, int32(2), v63, v64)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v12)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
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
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
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
					F_errmsg_internal(m, int32(50136), v7)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						F_errfinish(m, int32(520068), int32(19565), int32(277820))
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
				v16 = *(*int32)(unsafe.Add(mBase, _consts[168]))
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
								v35 = int32(20)
							default:
								v33 = int32(41)
								v35 = v33
							case 10:
								v35 = int32(37)
							case 29:
								v33 = int32(18)
								v35 = v33
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
								v40 = int32(*(*uint8)(unsafe.Add(mBase, _consts[456])))
								if v40 == int32(0) {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
									v47 = int32(1)
									if base.Ui32(l1) < base.Ui32(int32(12000)) {
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
												F_errmsg(m, int32(345099), v7+int32(16))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return
												} else {
													F_errfinish(m, int32(520068), int32(19576), int32(277820))
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
						v40 = int32(*(*uint8)(unsafe.Add(mBase, _consts[456])))
						if v40 == int32(0) {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
							v47 = int32(1)
							if base.Ui32(l1) < base.Ui32(int32(12000)) {
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
										F_errmsg(m, int32(345099), v7+int32(16))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_errfinish(m, int32(520068), int32(19576), int32(277820))
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
	var v19 int32
	_ = v19
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
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
	return v145
L2:
	;
	F_AccessTempTableNamespace(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L8
	} else {
		goto L49
	}
L3:
	;
	v140 = int32(0)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L45
	}
L5:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _consts[100]))
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
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v41 != 0 {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	return int32(0)
L9:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v19 == int32(0) {
		v38 = v18
		v39 = v19
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v39-v38 != 0 {
		goto L4
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	if v18 != v19 {
		v38 = v18
		v39 = v19
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v23 = v9
	v24 = v12
	goto L14
L14:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v28 == int32(0) {
		v38 = v27
		v39 = v28
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v38 = v27
	v39 = v28
	goto L11
L16:
	;
	v31 = int32(1)
	if v27 == v28 {
		v23 = v23 + v31
		v24 = v24 + v31
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L7
L19:
	;
	v42 = int32(248157)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _consts[385])))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v46 == int32(0) {
		v65 = v45
		v66 = v46
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v92 == int32(116) {
		goto L3
	} else {
		goto L37
	}
L22:
	;
	if v66-v65 == int32(0) {
		goto L3
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	if v45 != v46 {
		v65 = v45
		v66 = v46
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v50 = v41
	v51 = v42
	goto L26
L26:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v55 == int32(0) {
		v65 = v54
		v66 = v55
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v65 = v54
	v66 = v55
	goto L23
L28:
	;
	v58 = int32(1)
	if v54 == v55 {
		v50 = v50 + v58
		v51 = v51 + v58
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v71 = int32(0)
	v74 = F_GetSysCacheOid(m, int32(37), v41, v71, v71, v71)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	if v74 != 0 {
		v145 = v74
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v41
	F_errmsg(m, int32(78807), v7)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(525873), int32(3547), int32(457175))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, _consts[386])))
	if v99 != 0 {
		v140 = int32(1)
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[387]))
	if v101 != 0 {
		v145 = v101
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(294425), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(525873), int32(704), int32(440261))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v126
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v125
	F_errmsg(m, int32(725159), v7+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(525873), int32(668), int32(440261))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L8
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
	v144 = *(*int32)(unsafe.Add(mBase, _consts[260]))
	v145 = v144
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
	var v34 int32
	_ = v34
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
					v34 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v34)
					*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l4
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v34)
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
									F_errmsg_internal(m, int32(559985), v9)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(519717), int32(2048), int32(422188))
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
										F_errmsg_internal(m, int32(559985), v9)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(519717), int32(2048), int32(422188))
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
											F_errmsg_internal(m, int32(559985), v9)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(519717), int32(2048), int32(422188))
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
						v34 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v34)
						*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l4
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v34)
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
										F_errmsg_internal(m, int32(559985), v9)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(519717), int32(2048), int32(422188))
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
											F_errmsg_internal(m, int32(559985), v9)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(519717), int32(2048), int32(422188))
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
												F_errmsg_internal(m, int32(559985), v9)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(519717), int32(2048), int32(422188))
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
							v34 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v34)
							*(*int64)(unsafe.Add(mBase, uint32(v9)+12)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l4
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v34)
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
											F_errmsg_internal(m, int32(559985), v9)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(519717), int32(2048), int32(422188))
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
												F_errmsg_internal(m, int32(559985), v9)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(519717), int32(2048), int32(422188))
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
													F_errmsg_internal(m, int32(559985), v9)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(519717), int32(2048), int32(422188))
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
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
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
	v82 = int32(99999)
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
	goto L50
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _consts[815]))
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
L51:
	;
	return v133
L52:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _consts[815]))
	switch v240 - int32(1) {
	case 0:
		goto L86
	case 1:
		goto L85
	case 2:
		goto L84
	default:
		goto L87
	}
L53:
	;
	if v213 == v157 {
		goto L52
	} else {
		goto L73
	}
L54:
	;
	v213 = v209
	goto L53
L55:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	if base.Ui32(int32(127)) < base.Ui32(v157) {
		goto L69
	} else {
		goto L70
	}
L56:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _consts[814]))
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
	v180 = v157<<(uint(int32(2))%32) + int32(1907524)
	goto L62
L61:
	;
	v175 = F_case_index(m, v157)
	mBase = m.M
	v180 = v175<<(uint(int32(2))%32) + int32(1907520)
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
	if base.B2i32(base.Ui32(l1) <= base.Ui32(v213))&base.B2i32(base.Ui32(v213) <= base.Ui32(l2)) != 0 {
		goto L52
	} else {
		goto L74
	}
L74:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v219 <= v218 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v223 != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v218 + int32(1)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v232+v218<<(uint(int32(2))%32)))) = v213
	goto L52
L78:
	;
	v225 = v223
	goto L80
L79:
	;
	v225 = int32(19)
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v225
	return int32(0)
L81:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v316 != 0 {
		goto L110
	} else {
		goto L111
	}
L82:
	;
	if v290 == v157 {
		goto L81
	} else {
		goto L102
	}
L83:
	;
	v290 = v286
	goto L82
L84:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	if base.Ui32(int32(127)) < base.Ui32(v157) {
		goto L98
	} else {
		goto L99
	}
L85:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _consts[814]))
	if base.Ui32(int32(127)) < base.Ui32(v157) {
		goto L95
	} else {
		goto L96
	}
L86:
	;
	if base.Ui32(v157) <= base.Ui32(int32(127)) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	if base.Ui32(int32(127)) < base.Ui32(v157) {
		v286 = v157
		goto L83
	} else {
		goto L88
	}
L88:
	;
	v245 = F_pg_toupper(m, v157)
	mBase = m.M
	v290 = v245
	goto L82
L89:
	;
	v257 = v157<<(uint(int32(2))%32) + int32(1921156)
	goto L91
L90:
	;
	v252 = F_case_index(m, v157)
	mBase = m.M
	v257 = v252<<(uint(int32(2))%32) + int32(1921152)
	goto L91
L91:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	if v258 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v259 = v258
	goto L94
L93:
	;
	v259 = v157
	goto L94
L94:
	;
	v290 = v259
	goto L82
L95:
	;
	v271 = F_towupper(m, v157)
	mBase = m.M
	v290 = v271
	goto L82
L96:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+4)))
	if v264&int32(1) == int32(0) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v269 = F_pg_toupper(m, v157)
	mBase = m.M
	v290 = v269
	goto L82
L98:
	;
	if base.Ui32(int32(255)) < base.Ui32(v157) {
		v286 = v157
		goto L83
	} else {
		goto L101
	}
L99:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+4)))
	if v276&int32(1) == int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v281 = F_pg_toupper(m, v157)
	mBase = m.M
	v290 = v281
	goto L82
L101:
	;
	v285 = F_toupper(m, v157)
	mBase = m.M
	v286 = v285
	goto L83
L102:
	;
	if base.B2i32(base.Ui32(l1) <= base.Ui32(v290))&base.B2i32(base.Ui32(v290) <= base.Ui32(l2)) != 0 {
		goto L81
	} else {
		goto L103
	}
L103:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v296 <= v295 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v300 != 0 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v295 + int32(1)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v309+v295<<(uint(int32(2))%32)))) = v290
	goto L81
L107:
	;
	v302 = v300
	goto L109
L108:
	;
	v302 = int32(19)
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v302
	return int32(0)
L110:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L17
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v320 = v157 + int32(1)
	if base.Ui32(v320) <= base.Ui32(l2) {
		v157 = v320
		goto L50
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	goto L51
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
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
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
					v54 = v30
					m.G0 = v7 + int32(80)
					return v54
				} else {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+46)))
					if v32 != 0 {
						v54 = v30
						m.G0 = v7 + int32(80)
						return v54
					} else {
						v33 = *(*int64)(unsafe.Add(mBase, uint32(v7)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v33
						v35 = *(*int64)(unsafe.Add(mBase, uint32(v7)+64))
						*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v35
						v42 = F_bounds_adjacent(m, l0, v7+int32(32), v7+int32(24))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							if v42 != 0 {
								v54 = int32(1)
								m.G0 = v7 + int32(80)
								return v54
							} else {
								v44 = *(*int64)(unsafe.Add(mBase, uint32(v7)+48))
								*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v44
								v46 = *(*int64)(unsafe.Add(mBase, uint32(v7)+72))
								*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v46
								v52 = F_bounds_adjacent(m, l0, v7+int32(16), v7+int32(8))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v54 = v52
									m.G0 = v7 + int32(80)
									return v54
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
		v62 = m.ExcPending
		if v62 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(342551), int32(0))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(519717), int32(811), int32(326806))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
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
	v69 = F_lookup_type_cache(m, v60, int32(65536))
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
	F_errmsg_internal(m, int32(67587), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L28
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(519712), int32(1384), int32(295309))
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
	F_errmsg_internal(m, int32(389826), v11)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L28
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(519712), int32(558), int32(419268))
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
						F_errmsg_internal(m, int32(423191), int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(519712), int32(1352), int32(295059))
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
			F_errmsg_internal(m, int32(67487), int32(0))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(519712), int32(1348), int32(295059))
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
					v25 = F_lookup_type_cache(m, v19, int32(65536))
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
								F_errmsg_internal(m, int32(389826), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(519712), int32(558), int32(419268))
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
				v25 = F_lookup_type_cache(m, v19, int32(65536))
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
							F_errmsg_internal(m, int32(389826), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(519712), int32(558), int32(419268))
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
					v25 = F_lookup_type_cache(m, v19, int32(65536))
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
								F_errmsg_internal(m, int32(389826), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(519712), int32(558), int32(419268))
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
				v25 = F_lookup_type_cache(m, v19, int32(65536))
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
							F_errmsg_internal(m, int32(389826), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(519712), int32(558), int32(419268))
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
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
					v107 = v31
					m.G0 = v8 + int32(48)
					return v107
				} else {
					v39 = int32(0)
					if v35 != v32&int32(255) {
						v107 = v39
						m.G0 = v8 + int32(48)
						return v107
					} else {
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+44)))
						if v44 == int32(1) {
							if v43&int32(1) == int32(0) {
								v107 = v39
								m.G0 = v8 + int32(48)
								return v107
							} else {
								v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+46)))
								v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
								if v51 == v52 {
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
									if v78 == int32(1) {
										if v77&int32(1) == int32(0) {
											v107 = v39
										} else {
											v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
											v107 = base.B2i32(v85 == v86)
										}
										m.G0 = v8 + int32(48)
										return v107
									} else {
										if v77&int32(1) != 0 {
											v107 = v39
											m.G0 = v8 + int32(48)
											return v107
										} else {
											v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
											v95 = F_FunctionCall2Coll(m, l0+int32(212), v92, v93, v94)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												if v95 != 0 {
													v107 = v39
												} else {
													v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+29)))
													v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)))
													v99 = v97 & v98
													if v97 != 0 {
														v107 = v99
													} else {
														if v98&int32(1) != 0 {
															v107 = v99
														} else {
															v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
															v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
															v107 = base.B2i32(v102 == v103)
														}
													}
												}
												m.G0 = v8 + int32(48)
												return v107
											}
										}
									}
								} else {
									v107 = v39
									m.G0 = v8 + int32(48)
									return v107
								}
							}
						} else {
							if v43&int32(1) != 0 {
								v107 = v39
								m.G0 = v8 + int32(48)
								return v107
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
								v61 = F_FunctionCall2Coll(m, l0+int32(212), v58, v59, v60)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									if v61 != 0 {
										v107 = v39
										m.G0 = v8 + int32(48)
										return v107
									} else {
										v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
										v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+45)))
										if v64 == int32(0) {
											if v63&int32(1) != 0 {
												v107 = v39
												m.G0 = v8 + int32(48)
												return v107
											} else {
												v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+46)))
												v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
												if v69 == v70 {
													v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
													v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
													if v78 == int32(1) {
														if v77&int32(1) == int32(0) {
															v107 = v39
														} else {
															v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
															v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
															v107 = base.B2i32(v85 == v86)
														}
														m.G0 = v8 + int32(48)
														return v107
													} else {
														if v77&int32(1) != 0 {
															v107 = v39
															m.G0 = v8 + int32(48)
															return v107
														} else {
															v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
															v93 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
															v94 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
															v95 = F_FunctionCall2Coll(m, l0+int32(212), v92, v93, v94)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return int32(0)
															} else {
																if v95 != 0 {
																	v107 = v39
																} else {
																	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+29)))
																	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)))
																	v99 = v97 & v98
																	if v97 != 0 {
																		v107 = v99
																	} else {
																		if v98&int32(1) != 0 {
																			v107 = v99
																		} else {
																			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
																			v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
																			v107 = base.B2i32(v102 == v103)
																		}
																	}
																}
																m.G0 = v8 + int32(48)
																return v107
															}
														}
													}
												} else {
													v107 = v39
													m.G0 = v8 + int32(48)
													return v107
												}
											}
										} else {
											if v63&int32(1) == int32(0) {
												v107 = v39
												m.G0 = v8 + int32(48)
												return v107
											} else {
												v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
												v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
												if v78 == int32(1) {
													if v77&int32(1) == int32(0) {
														v107 = v39
													} else {
														v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
														v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
														v107 = base.B2i32(v85 == v86)
													}
													m.G0 = v8 + int32(48)
													return v107
												} else {
													if v77&int32(1) != 0 {
														v107 = v39
														m.G0 = v8 + int32(48)
														return v107
													} else {
														v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
														v93 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
														v94 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
														v95 = F_FunctionCall2Coll(m, l0+int32(212), v92, v93, v94)
														mBase = m.M
														v96 = m.ExcPending
														if v96 != 0 {
															return int32(0)
														} else {
															if v95 != 0 {
																v107 = v39
															} else {
																v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+29)))
																v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)))
																v99 = v97 & v98
																if v97 != 0 {
																	v107 = v99
																} else {
																	if v98&int32(1) != 0 {
																		v107 = v99
																	} else {
																		v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+30)))
																		v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
																		v107 = base.B2i32(v102 == v103)
																	}
																}
															}
															m.G0 = v8 + int32(48)
															return v107
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
		v115 = m.ExcPending
		if v115 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(342551), int32(0))
			mBase = m.M
			v119 = m.ExcPending
			if v119 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(519717), int32(586), int32(327283))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
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
						F_errmsg_internal(m, int32(389953), v7)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(519717), int32(1776), int32(419273))
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
					F_errmsg_internal(m, int32(389953), v7)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519717), int32(1776), int32(419273))
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
						F_errmsg_internal(m, int32(423143), int32(0))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(519717), int32(1234), int32(294935))
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
										F_errmsg_internal(m, int32(389953), v7)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(519717), int32(1776), int32(419273))
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
									F_errmsg_internal(m, int32(389953), v7)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(519717), int32(1776), int32(419273))
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
			F_errmsg_internal(m, int32(67209), int32(0))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(519717), int32(1230), int32(294935))
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		v11 = int32(1)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3+int32(base.Ui32(v7)>>(uint(int32(2))%32))-v11))))
		return int32(base.Ui32(v13)>>(uint(v11)%32)) & v11
	}
}
func F_range_minus_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
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
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
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
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	F_range_deserialize(m, l0, l1, v10+int32(-24), v10+int32(-40), v10+int32(-49))
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
	F_range_deserialize(m, l0, l2, v10+int32(-32), v10+int32(-48), v10+int32(-50))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v32 != 0 {
		v373 = l1
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
		goto L196
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L192
	}
L6:
	;
	m.G0 = v12 - int32(-64)
	return v373
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)))
	if v33 != 0 {
		v373 = l1
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+36)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
	if v35 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+28)))
	if v178 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L10:
	;
	v175 = int32(0)
	v176 = v170
	v177 = v172
	goto L9
L11:
	;
	v121 = int32(1)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)))
	if v122 == v121 {
		goto L63
	} else {
		goto L64
	}
L12:
	;
	v114 = int32(1)
	if v70&v114 != 0 {
		goto L60
	} else {
		goto L61
	}
L13:
	;
	v96 = int32(1)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)))
	if v97 == v96 {
		goto L50
	} else {
		goto L51
	}
L14:
	;
	v91 = int32(1)
	if v38&v91 != 0 {
		goto L47
	} else {
		goto L48
	}
L15:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+46)))
	if v34&int32(1) == int32(0) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v34&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)))
	if v44 == v38&int32(255) {
		v95 = int32(0)
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v49 = int32(1)
	if v38&v49 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v52 = int32(-1)
	goto L22
L21:
	;
	v52 = v49
	goto L22
L22:
	;
	v95 = v52
	goto L13
L23:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)))
	if v57 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v64 = F_FunctionCall2Coll(m, l0+int32(212), v61, v62, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L29
	}
L26:
	;
	v58 = int32(1)
	goto L28
L27:
	;
	v58 = int32(-1)
	goto L28
L28:
	;
	v120 = v58
	goto L11
L29:
	;
	if v64 != 0 {
		v120 = v64
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+37)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+45)))
	if v67 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+46)))
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
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)))
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
	v120 = int32(0)
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
	v120 = v82
	goto L11
L41:
	;
	v120 = int32(0)
	goto L11
L42:
	;
	goto L43
L43:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)))
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
	v120 = v89
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
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	if v100 == v38&int32(255) {
		v175 = v96
		v176 = v95
		v177 = int32(0)
		goto L9
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v110 = int32(1)
	if v38&v110 != 0 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v105 = int32(1)
	if v38&v105 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v108 = int32(-1)
	goto L56
L55:
	;
	v108 = v105
	goto L56
L56:
	;
	v175 = v96
	v176 = v95
	v177 = v108
	goto L9
L57:
	;
	v113 = int32(-1)
	goto L59
L58:
	;
	v113 = v110
	goto L59
L59:
	;
	v170 = v95
	v172 = v113
	goto L10
L60:
	;
	v118 = v114
	goto L62
L61:
	;
	v118 = int32(-1)
	goto L62
L62:
	;
	v120 = v118
	goto L11
L63:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	if v127 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v129 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v135 = F_FunctionCall2Coll(m, l0+int32(212), v132, v133, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L69
	}
L66:
	;
	v128 = int32(1)
	goto L68
L67:
	;
	v128 = int32(-1)
	goto L68
L68:
	;
	v175 = v121
	v176 = v120
	v177 = v128
	goto L9
L69:
	;
	if v135 != 0 {
		v175 = v129
		v176 = v120
		v177 = v135
		goto L9
	} else {
		goto L70
	}
L70:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+45)))
	if v138 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+46)))
	if v137&int32(1) == int32(0) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if v137&int32(1) != 0 {
		goto L86
	} else {
		goto L87
	}
L74:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	if v146 == v141&int32(255) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v156 = int32(1)
	if v141&v156 != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v175 = v129
	v176 = v120
	v177 = int32(0)
	goto L9
L78:
	;
	goto L79
L79:
	;
	v151 = int32(1)
	if v141&v151 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v155 = v151
	goto L82
L81:
	;
	v155 = int32(-1)
	goto L82
L82:
	;
	v170 = v120
	v172 = v155
	goto L10
L83:
	;
	v160 = v156
	goto L85
L84:
	;
	v160 = int32(-1)
	goto L85
L85:
	;
	v175 = v129
	v176 = v120
	v177 = v160
	goto L9
L86:
	;
	v175 = v129
	v176 = v120
	v177 = int32(0)
	goto L9
L87:
	;
	goto L88
L88:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	if v166 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v167 = int32(-1)
	goto L91
L90:
	;
	v167 = int32(1)
	goto L91
L91:
	;
	v170 = v120
	v172 = v167
	goto L10
L92:
	;
	if int32(0) < v177 {
		v373 = l1
		goto L6
	} else {
		goto L179
	}
L93:
	;
	if int32(0) <= v176 {
		v307 = v299
		v308 = v300
		goto L92
	} else {
		goto L177
	}
L94:
	;
	if v175 != 0 {
		goto L148
	} else {
		goto L149
	}
L95:
	;
	v252 = int32(1)
	if v211&v252 != 0 {
		goto L145
	} else {
		goto L146
	}
L96:
	;
	if v175 != 0 {
		goto L133
	} else {
		goto L134
	}
L97:
	;
	v234 = int32(1)
	if v181&v234 != 0 {
		goto L130
	} else {
		goto L131
	}
L98:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)))
	if v34&int32(1) == int32(0) {
		goto L97
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if v34&int32(1) != 0 {
		goto L106
	} else {
		goto L107
	}
L101:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)))
	if v181 == v187 {
		v238 = int32(0)
		goto L96
	} else {
		goto L102
	}
L102:
	;
	v190 = int32(1)
	if v181&v190 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v193 = int32(-1)
	goto L105
L104:
	;
	v193 = v190
	goto L105
L105:
	;
	v238 = v193
	goto L96
L106:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)))
	if v198 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	goto L108
L108:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v205 = F_FunctionCall2Coll(m, l0+int32(212), v202, v203, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L112
	}
L109:
	;
	v199 = int32(1)
	goto L111
L110:
	;
	v199 = int32(-1)
	goto L111
L111:
	;
	v258 = v199
	goto L94
L112:
	;
	if v205 != 0 {
		v258 = v205
		goto L94
	} else {
		goto L113
	}
L113:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+37)))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+29)))
	if v208 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)))
	if v207&int32(1) == int32(0) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	if v207&int32(1) != 0 {
		goto L124
	} else {
		goto L125
	}
L117:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)))
	if v216 != v211&int32(255) {
		goto L95
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v221 = int32(1)
	if v211&v221 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v258 = int32(0)
	goto L94
L121:
	;
	v225 = v221
	goto L123
L122:
	;
	v225 = int32(-1)
	goto L123
L123:
	;
	v258 = v225
	goto L94
L124:
	;
	v258 = int32(0)
	goto L94
L125:
	;
	goto L126
L126:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)))
	if v231 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v232 = int32(-1)
	goto L129
L128:
	;
	v232 = int32(1)
	goto L129
L129:
	;
	v258 = v232
	goto L94
L130:
	;
	v237 = int32(-1)
	goto L132
L131:
	;
	v237 = v234
	goto L132
L132:
	;
	v238 = v237
	goto L96
L133:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	if v239 == v181 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	goto L135
L135:
	;
	v248 = int32(1)
	if v181&v248 != 0 {
		goto L142
	} else {
		goto L143
	}
L136:
	;
	v307 = int32(0)
	v308 = v238
	goto L92
L137:
	;
	goto L138
L138:
	;
	v243 = int32(1)
	if v181&v243 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v246 = int32(-1)
	goto L141
L140:
	;
	v246 = v243
	goto L141
L141:
	;
	v299 = v246
	v300 = v238
	goto L93
L142:
	;
	v251 = int32(-1)
	goto L144
L143:
	;
	v251 = v248
	goto L144
L144:
	;
	v299 = v251
	v300 = v238
	goto L93
L145:
	;
	v256 = v252
	goto L147
L146:
	;
	v256 = int32(-1)
	goto L147
L147:
	;
	v258 = v256
	goto L94
L148:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	if v261 != 0 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v268 = F_FunctionCall2Coll(m, l0+int32(212), v265, v266, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L154
	}
L151:
	;
	v262 = int32(1)
	goto L153
L152:
	;
	v262 = int32(-1)
	goto L153
L153:
	;
	v299 = v262
	v300 = v258
	goto L93
L154:
	;
	if v268 != 0 {
		v299 = v268
		v300 = v258
		goto L93
	} else {
		goto L155
	}
L155:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+29)))
	if v271 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+30)))
	if v270&int32(1) == int32(0) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L158
L158:
	;
	if v270&int32(1) != 0 {
		goto L171
	} else {
		goto L172
	}
L159:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	if v279 == v274 {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	goto L161
L161:
	;
	v287 = int32(1)
	if v274&v287 != 0 {
		goto L168
	} else {
		goto L169
	}
L162:
	;
	v307 = int32(0)
	v308 = v258
	goto L92
L163:
	;
	goto L164
L164:
	;
	v282 = int32(1)
	if v274&v282 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v286 = v282
	goto L167
L166:
	;
	v286 = int32(-1)
	goto L167
L167:
	;
	v299 = v286
	v300 = v258
	goto L93
L168:
	;
	v291 = v287
	goto L170
L169:
	;
	v291 = int32(-1)
	goto L170
L170:
	;
	v299 = v291
	v300 = v258
	goto L93
L171:
	;
	v307 = int32(0)
	v308 = v258
	goto L92
L172:
	;
	goto L173
L173:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	if v297 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v298 = int32(-1)
	goto L176
L175:
	;
	v298 = int32(1)
	goto L176
L176:
	;
	v299 = v298
	v300 = v258
	goto L93
L177:
	;
	if int32(0) < v299 {
		goto L5
	} else {
		goto L178
	}
L178:
	;
	v307 = v299
	v308 = v300
	goto L92
L179:
	;
	if v308 < int32(0) {
		v373 = l1
		goto L6
	} else {
		goto L180
	}
L180:
	;
	if v176 < int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	if int32(0) < v176 {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	if int32(0) < v307 {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v319 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+62)) = uint8(v319)
	v321 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+60)) = uint16(v321)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v321
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v321
	*(*int32)(unsafe.Add(mBase, uint32(v12)+51)) = v321
	v335 = F_make_range(m, l0, v10+int32(-8), v10+int32(-16), v319, v321)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v373 = v335
	goto L6
L185:
	;
	if v176 < int32(0) {
		goto L4
	} else {
		goto L189
	}
L186:
	;
	if int32(0) < v307 {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v341 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+38)) = uint8(v341)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+37)))
	v345 = v343 ^ int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+37)) = uint8(v345)
	v353 = F_make_range(m, l0, v10+int32(-24), v10+int32(-32), v341, v341)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	v373 = v353
	goto L6
L189:
	;
	if v307 < int32(0) {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	v359 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)) = uint8(v359)
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)))
	v363 = v361 ^ v359
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)) = uint8(v363)
	v369 = int32(0)
	v371 = F_make_range(m, l0, v10+int32(-48), v10+int32(-40), v369, v369)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v373 = v371
	goto L6
L192:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_errmsg(m, int32(122814), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(519717), int32(1023), int32(326955))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	F_errmsg_internal(m, int32(123019), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(519717), int32(1045), int32(326955))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_check_stack_depth(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v21 = F_get_range_io_data(m, l0, v14, int32(2))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = F_pq_getmsgbyte(m, v15)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23&int32(9) == int32(0) {
					v30 = F_pq_getmsgint(m, v15, int32(4))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = F_pq_getmsgbytes(m, v15, v30)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_initStringInfo(m, v11+int32(8))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								F_appendBinaryStringInfo(m, v11+int32(8), v32, v30)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
									v47 = F_ReceiveFunctionCall(m, v21+int32(4), v11+int32(8), v46, v13)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v49 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
										F_pfree(m, v49)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											v52 = v47
											*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v52
											if v23&int32(17) == int32(0) {
												v60 = F_pq_getmsgint(m, v15, int32(4))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return int32(0)
												} else {
													v62 = F_pq_getmsgbytes(m, v15, v60)
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return int32(0)
													} else {
														F_initStringInfo(m, v11+int32(8))
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															F_appendBinaryStringInfo(m, v11+int32(8), v62, v60)
															mBase = m.M
															v71 = m.ExcPending
															if v71 != 0 {
																return int32(0)
															} else {
																v76 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
																v77 = F_ReceiveFunctionCall(m, v21+int32(4), v11+int32(8), v76, v13)
																mBase = m.M
																v78 = m.ExcPending
																if v78 != 0 {
																	return int32(0)
																} else {
																	v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
																	F_pfree(m, v79)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return int32(0)
																	} else {
																		v83 = v77
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v83
																		F_pq_getmsgend(m, v15)
																		mBase = m.M
																		v86 = m.ExcPending
																		if v86 != 0 {
																			return int32(0)
																		} else {
																			v87 = int32(0)
																			*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v87)
																			v89 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v89)
																			v94 = int32(base.Ui32(v23)>>(uint(int32(3))%32)) & v89
																			*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v94)
																			v99 = int32(base.Ui32(v23)>>(uint(int32(2))%32)) & v89
																			*(*uint8)(unsafe.Add(mBase, uint32(v11)+13)) = uint8(v99)
																			v104 = int32(base.Ui32(v23)>>(uint(int32(4))%32)) & v89
																			*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)) = uint8(v104)
																			v109 = int32(base.Ui32(v23)>>(uint(v89)%32)) & v89
																			*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v109)
																			v111 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
																			v119 = F_make_range(m, v111, v11+int32(24), v11+int32(8), v23&v89, v87)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v11 + int32(32)
																				return v119
																			}
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v83 = v2
												*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v83
												F_pq_getmsgend(m, v15)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													v87 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v87)
													v89 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v89)
													v94 = int32(base.Ui32(v23)>>(uint(int32(3))%32)) & v89
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v94)
													v99 = int32(base.Ui32(v23)>>(uint(int32(2))%32)) & v89
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+13)) = uint8(v99)
													v104 = int32(base.Ui32(v23)>>(uint(int32(4))%32)) & v89
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)) = uint8(v104)
													v109 = int32(base.Ui32(v23)>>(uint(v89)%32)) & v89
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v109)
													v111 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
													v119 = F_make_range(m, v111, v11+int32(24), v11+int32(8), v23&v89, v87)
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														m.G0 = v11 + int32(32)
														return v119
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
					v52 = v2
					*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v52
					if v23&int32(17) == int32(0) {
						v60 = F_pq_getmsgint(m, v15, int32(4))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = F_pq_getmsgbytes(m, v15, v60)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_initStringInfo(m, v11+int32(8))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_appendBinaryStringInfo(m, v11+int32(8), v62, v60)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
										v77 = F_ReceiveFunctionCall(m, v21+int32(4), v11+int32(8), v76, v13)
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
											F_pfree(m, v79)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												v83 = v77
												*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v83
												F_pq_getmsgend(m, v15)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													v87 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v87)
													v89 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v89)
													v94 = int32(base.Ui32(v23)>>(uint(int32(3))%32)) & v89
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v94)
													v99 = int32(base.Ui32(v23)>>(uint(int32(2))%32)) & v89
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+13)) = uint8(v99)
													v104 = int32(base.Ui32(v23)>>(uint(int32(4))%32)) & v89
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)) = uint8(v104)
													v109 = int32(base.Ui32(v23)>>(uint(v89)%32)) & v89
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v109)
													v111 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
													v119 = F_make_range(m, v111, v11+int32(24), v11+int32(8), v23&v89, v87)
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														m.G0 = v11 + int32(32)
														return v119
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v83 = v2
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v83
						F_pq_getmsgend(m, v15)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							v87 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v87)
							v89 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v89)
							v94 = int32(base.Ui32(v23)>>(uint(int32(3))%32)) & v89
							*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v94)
							v99 = int32(base.Ui32(v23)>>(uint(int32(2))%32)) & v89
							*(*uint8)(unsafe.Add(mBase, uint32(v11)+13)) = uint8(v99)
							v104 = int32(base.Ui32(v23)>>(uint(int32(4))%32)) & v89
							*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)) = uint8(v104)
							v109 = int32(base.Ui32(v23)>>(uint(v89)%32)) & v89
							*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v109)
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							v119 = F_make_range(m, v111, v11+int32(24), v11+int32(8), v23&v89, v87)
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int32(0)
							} else {
								m.G0 = v11 + int32(32)
								return v119
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		v11 = int32(1)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3+int32(base.Ui32(v7)>>(uint(int32(2))%32))-v11))))
		return int32(base.Ui32(v13)>>(uint(int32(4))%32)) & v11
	}
}
