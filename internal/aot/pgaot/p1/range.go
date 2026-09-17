package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_RangeVarCallbackForAttachIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	if v11 == int32(0) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
		F_LockRelationOid(m, v14, int32(1))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v18)
			if l1 == l2 {
				if l1 == int32(0) {
					m.G0 = v9 + int32(16)
					return
				} else {
					v33 = F_SearchSysCache1(m, int32(57), l1)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						if v33 == int32(0) {
							m.G0 = v9 + int32(16)
							return
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
							v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v38)+119)))
							if v40|int32(32) != int32(105) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v67
										F_errmsg(m, int32(_a_F_RangeVarCallbackForAttachIndex_0), v9)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_RangeVarCallbackForAttachIndex_1), int32(_a_F_RangeVarCallbackForAttachIndex_2), int32(_a_F_RangeVarCallbackForAttachIndex_3))
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
								}
							} else {
								F_ReleaseCatCache(m, v33)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									v48 = F_IndexGetRelation(m, l1, int32(0))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
										F_LockRelationOid(m, v48, int32(1))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											m.G0 = v9 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				if v21 == int32(0) {
					if l1 == int32(0) {
						m.G0 = v9 + int32(16)
						return
					} else {
						v33 = F_SearchSysCache1(m, int32(57), l1)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							if v33 == int32(0) {
								m.G0 = v9 + int32(16)
								return
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
								v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v38)+119)))
								if v40|int32(32) != int32(105) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v67
											F_errmsg(m, int32(_a_F_RangeVarCallbackForAttachIndex_0), v9)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_RangeVarCallbackForAttachIndex_1), int32(_a_F_RangeVarCallbackForAttachIndex_2), int32(_a_F_RangeVarCallbackForAttachIndex_3))
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
									}
								} else {
									F_ReleaseCatCache(m, v33)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										v48 = F_IndexGetRelation(m, l1, int32(0))
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
											F_LockRelationOid(m, v48, int32(1))
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return
											} else {
												m.G0 = v9 + int32(16)
												return
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_UnlockRelationOid(m, v21, int32(1))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
						if l1 == int32(0) {
							m.G0 = v9 + int32(16)
							return
						} else {
							v33 = F_SearchSysCache1(m, int32(57), l1)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								if v33 == int32(0) {
									m.G0 = v9 + int32(16)
									return
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
									v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
									v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v38)+119)))
									if v40|int32(32) != int32(105) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return
											} else {
												v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v9))) = v67
												F_errmsg(m, int32(_a_F_RangeVarCallbackForAttachIndex_0), v9)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_RangeVarCallbackForAttachIndex_1), int32(_a_F_RangeVarCallbackForAttachIndex_2), int32(_a_F_RangeVarCallbackForAttachIndex_3))
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
										}
									} else {
										F_ReleaseCatCache(m, v33)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											v48 = F_IndexGetRelation(m, l1, int32(0))
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
												F_LockRelationOid(m, v48, int32(1))
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return
												} else {
													m.G0 = v9 + int32(16)
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
		if l1 == l2 {
			if l1 == int32(0) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v33 = F_SearchSysCache1(m, int32(57), l1)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					if v33 == int32(0) {
						m.G0 = v9 + int32(16)
						return
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v38)+119)))
						if v40|int32(32) != int32(105) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v67
									F_errmsg(m, int32(_a_F_RangeVarCallbackForAttachIndex_0), v9)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_RangeVarCallbackForAttachIndex_1), int32(_a_F_RangeVarCallbackForAttachIndex_2), int32(_a_F_RangeVarCallbackForAttachIndex_3))
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
							}
						} else {
							F_ReleaseCatCache(m, v33)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								v48 = F_IndexGetRelation(m, l1, int32(0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
									F_LockRelationOid(m, v48, int32(1))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										m.G0 = v9 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			if v21 == int32(0) {
				if l1 == int32(0) {
					m.G0 = v9 + int32(16)
					return
				} else {
					v33 = F_SearchSysCache1(m, int32(57), l1)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						if v33 == int32(0) {
							m.G0 = v9 + int32(16)
							return
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
							v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v38)+119)))
							if v40|int32(32) != int32(105) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v67
										F_errmsg(m, int32(_a_F_RangeVarCallbackForAttachIndex_0), v9)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_RangeVarCallbackForAttachIndex_1), int32(_a_F_RangeVarCallbackForAttachIndex_2), int32(_a_F_RangeVarCallbackForAttachIndex_3))
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
								}
							} else {
								F_ReleaseCatCache(m, v33)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									v48 = F_IndexGetRelation(m, l1, int32(0))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
										F_LockRelationOid(m, v48, int32(1))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											m.G0 = v9 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_UnlockRelationOid(m, v21, int32(1))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
					if l1 == int32(0) {
						m.G0 = v9 + int32(16)
						return
					} else {
						v33 = F_SearchSysCache1(m, int32(57), l1)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							if v33 == int32(0) {
								m.G0 = v9 + int32(16)
								return
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
								v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v38)+119)))
								if v40|int32(32) != int32(105) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v67
											F_errmsg(m, int32(_a_F_RangeVarCallbackForAttachIndex_0), v9)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_RangeVarCallbackForAttachIndex_1), int32(_a_F_RangeVarCallbackForAttachIndex_2), int32(_a_F_RangeVarCallbackForAttachIndex_3))
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
									}
								} else {
									F_ReleaseCatCache(m, v33)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										v48 = F_IndexGetRelation(m, l1, int32(0))
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
											F_LockRelationOid(m, v48, int32(1))
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return
											} else {
												m.G0 = v9 + int32(16)
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
}
func F_RangeVarGetAndCheckCreationNamespace(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int64
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L67
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetAndCheckCreationNamespace[0]))
	v20 = F_get_database_name(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v51 = *(*int64)(unsafe.Add(mBase, _c_F_RangeVarGetAndCheckCreationNamespace[1]))
	v55 = v4
	v56 = v4
	v61 = v4
	v63 = v51
	goto L15
L5:
	;
	return int32(0)
L6:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v26 == int32(0))|base.B2i32(v26 != v29) != 0 {
		v47 = v26
		v48 = v29
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v47-v48 != 0 {
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v32 = v17
	v33 = v20
	goto L10
L10:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v37 == int32(0) {
		v47 = v37
		v48 = v36
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v47 = v37
	v48 = v36
	goto L8
L12:
	;
	v40 = int32(1)
	if v37 == v36 {
		v32 = v32 + v40
		v33 = v33 + v40
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L4
L15:
	;
	v64 = F_RangeVarGetCreationNamespace(m, l0)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L17
	}
L16:
	;
	F_RangeVarAdjustRelationPersistence(m, l0, v149)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L63
	}
L17:
	;
	if l2 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v68 = F_get_relname_relid(m, v67, v64)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	v70 = int32(0)
	goto L20
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetAndCheckCreationNamespace[2]))
	if v72 == int32(0) {
		v149 = v64
		v150 = v70
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v70 = v68
	goto L20
L22:
	;
	goto L16
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetAndCheckCreationNamespace[3]))
	v79 = F_object_aclcheck(m, int32(2615), v64, v77, int64(512))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	if v79 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v82 = F_get_namespace_name(m, v64)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v61 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	F_aclcheck_error(m, v79, int32(36), v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v64 != v55 {
		goto L41
	} else {
		goto L42
	}
L31:
	;
	if base.B2i32(v70 != v56)|base.B2i32(v64 != v55) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v149 = v55
	v150 = v56
	goto L22
L33:
	;
	goto L34
L34:
	;
	if v64 != v55 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_UnlockDatabaseObject(m, int32(2615), v55, int32(1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v98 = int32(0)
	if base.B2i32(l1 == v98)|base.B2i32(v56 == v98)|base.B2i32(v70 == v56) != 0 {
		goto L30
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	F_UnlockRelationOid(m, v56, l1)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	goto L30
L41:
	;
	F_LockDatabaseObject(m, int32(2615), v64, int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v112 = int32(0)
	if base.B2i32(l1 == v112)|base.B2i32(v70 == v112) != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	v147 = *(*int64)(unsafe.Add(mBase, _c_F_RangeVarGetAndCheckCreationNamespace[1]))
	if v63 != v147 {
		v55 = v64
		v56 = v70
		v61 = int32(1)
		v63 = v147
		goto L15
	} else {
		goto L62
	}
L46:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetAndCheckCreationNamespace[3]))
	v120 = F_object_ownercheck(m, int32(1259), v70, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	if v120 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v125 = F_get_rel_relkind(m, v70)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v70 == v56 {
		goto L45
	} else {
		goto L60
	}
L51:
	;
	switch v125 - int32(73) {
	case 0, 32:
		v136 = int32(20)
		goto L53
	default:
		goto L54
	case 10:
		goto L58
	case 29:
		goto L55
	case 36:
		goto L56
	case 45:
		goto L57
	}
L52:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_aclcheck_error(m, int32(2), v138, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L59
	}
L53:
	;
	v138 = v136
	goto L52
L54:
	;
	v136 = int32(41)
	goto L53
L55:
	;
	v138 = int32(18)
	goto L52
L56:
	;
	v138 = int32(23)
	goto L52
L57:
	;
	v138 = int32(51)
	goto L52
L58:
	;
	v138 = int32(37)
	goto L52
L59:
	;
	goto L50
L60:
	;
	F_LockRelationOid(m, v70, l1)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	goto L45
L62:
	;
	v149 = v64
	v150 = v70
	goto L22
L63:
	;
	if l2 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v150
	goto L66
L65:
	;
	goto L66
L66:
	;
	m.G0 = v15 + int32(16)
	return v149
L67:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v168 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v169
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v168
	F_errmsg(m, int32(_a_F_RangeVarGetAndCheckCreationNamespace_0), v15)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_RangeVarGetAndCheckCreationNamespace_1), int32(760), int32(_a_F_RangeVarGetAndCheckCreationNamespace_2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RangeVarGetRelidExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
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
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v255 int64
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int64
	_ = v273
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int64
	_ = v315
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L5
	} else {
		goto L101
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetRelidExtended[0]))
	v24 = F_get_database_name(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v55 = l2 & int32(1)
	v59 = *(*int64)(unsafe.Add(mBase, _c_F_RangeVarGetRelidExtended[1]))
	v66 = v6
	v74 = v6
	v76 = v59
	goto L17
L5:
	;
	return int32(0)
L6:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if base.B2i32(v30 == int32(0))|base.B2i32(v30 != v33) != 0 {
		v51 = v30
		v52 = v33
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v51-v52 != 0 {
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v36 = v21
	v37 = v24
	goto L10
L10:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v41 == int32(0) {
		v51 = v41
		v52 = v40
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v51 = v41
	v52 = v40
	goto L8
L12:
	;
	v44 = int32(1)
	if v41 == v40 {
		v36 = v36 + v44
		v37 = v37 + v44
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L4
L15:
	;
	m.G0 = v19 + int32(80)
	return v301
L16:
	;
	F_errfinish(m, int32(_a_F_RangeVarGetRelidExtended_0), v296, int32(_a_F_RangeVarGetRelidExtended_1))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L5
	} else {
		goto L100
	}
L17:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v77 == int32(116) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	if v257 != 0 {
		v301 = v257
		goto L15
	} else {
		goto L86
	}
L19:
	;
	if l3 != 0 {
		goto L48
	} else {
		goto L49
	}
L20:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L38
	}
L21:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v122 = F_get_relname_relid(m, v121, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L37
	}
L22:
	;
	v80 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetRelidExtended[2]))
	if v82 == v80 {
		v185 = v80
		goto L19
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v109 == int32(0) {
		goto L20
	} else {
		goto L33
	}
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v85 == int32(0) {
		v120 = v82
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v88 = F_LookupExplicitNamespace(m, v85, v55)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetRelidExtended[2]))
	if v88 == v91 {
		v120 = v88
		goto L21
	} else {
		goto L28
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_RangeVarGetRelidExtended_2), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_RangeVarGetRelidExtended_0), int32(519), int32(_a_F_RangeVarGetRelidExtended_1))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v112 = F_LookupExplicitNamespace(m, v109, v55)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	if v55 == int32(0) {
		v120 = v112
		goto L21
	} else {
		goto L35
	}
L35:
	;
	v116 = int32(0)
	if v112 == v116 {
		v185 = v116
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v120 = v112
	goto L21
L37:
	;
	v185 = v122
	goto L19
L38:
	;
	v127 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarGetRelidExtended[3]))
	if v129 == v127 {
		v185 = v127
		goto L19
	} else {
		goto L39
	}
L39:
	;
	v132 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v132 < v133 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v143 = v132
	goto L43
L41:
	;
	goto L42
L42:
	;
	v185 = int32(0)
	goto L19
L43:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152+v143<<(uint(int32(2))%32))))
	v157 = F_get_relname_relid(m, v124, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	if v157 != 0 {
		v185 = v157
		goto L19
	} else {
		goto L46
	}
L46:
	;
	v160 = v143 + int32(1)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v160 < v161 {
		v143 = v160
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	m.T0[l3].(func(*base.Module, int32, int32, int32, int32))(m, l0, v185, v66, l4)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L5
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if l1 == int32(0) {
		v257 = v185
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	goto L18
L53:
	;
	if v74 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v185 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L55:
	;
	if v185 == v66 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v257 = v66
	goto L52
L57:
	;
	goto L58
L58:
	;
	if v66 == int32(0) {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	F_UnlockRelationOid(m, v66, l1)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	goto L54
L61:
	;
	v255 = *(*int64)(unsafe.Add(mBase, _c_F_RangeVarGetRelidExtended[1]))
	if v76 != v255 {
		v66 = v185
		v74 = int32(1)
		v76 = v255
		goto L17
	} else {
		goto L85
	}
L62:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if l2&int32(6) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L61
L66:
	;
	F_LockRelationOid(m, v185, l1)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L5
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v215 = F_ConditionalLockRelationOid(m, v185, l1)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L70
	}
L69:
	;
	goto L61
L70:
	;
	if v215 != 0 {
		goto L61
	} else {
		goto L71
	}
L71:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v218 = int32(0)
	if l2&int32(4) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v223 = int32(14)
	goto L74
L73:
	;
	v223 = int32(21)
	goto L74
L74:
	;
	v225 = F_errstart(m, v223, int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	if v217 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v225 == int32(0) {
		v301 = v218
		goto L15
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v225 == int32(0) {
		v301 = v218
		goto L15
	} else {
		goto L82
	}
L79:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v232 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v232
	F_errmsg(m, int32(_a_F_RangeVarGetRelidExtended_3), v19+int32(48))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	v293 = v218
	v296 = int32(601)
	goto L16
L82:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v245
	F_errmsg(m, int32(_a_F_RangeVarGetRelidExtended_4), v19+int32(32))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	v293 = v218
	v296 = int32(606)
	goto L16
L85:
	;
	v257 = v185
	goto L52
L86:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v261 = int32(0)
	if v55 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v264 = int32(14)
	goto L89
L88:
	;
	v264 = int32(21)
	goto L89
L89:
	;
	v266 = F_errstart(m, v264, int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	if v260 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v266 == int32(0) {
		v301 = v261
		goto L15
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	if v266 == int32(0) {
		v301 = v261
		goto L15
	} else {
		goto L97
	}
L94:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	v273 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v273
	F_errmsg(m, int32(_a_F_RangeVarGetRelidExtended_5), v19+int32(16))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	v293 = v261
	v296 = int32(634)
	goto L16
L97:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v286
	F_errmsg(m, int32(_a_F_RangeVarGetRelidExtended_6), v19)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	v293 = v261
	v296 = int32(639)
	goto L16
L100:
	;
	v301 = v293
	goto L15
L101:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	v315 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v316
	*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v315
	F_errmsg(m, int32(_a_F_RangeVarGetRelidExtended_7), v19-int32(-64))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_RangeVarGetRelidExtended_0), int32(464), int32(_a_F_RangeVarGetRelidExtended_1))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addRangeClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 float64
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 float64
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v24 = l0
	goto L12
L2:
	;
	if v6 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v13 = int32(0)
	if v6 == v13 {
		v21 = v13
		v22 = l3
		goto L1
	} else {
		goto L8
	}
L5:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v10 = v8
	goto L7
L6:
	;
	v10 = int32(0)
	goto L7
L7:
	;
	v21 = v10
	v22 = l3 ^ int32(1)
	goto L1
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v16 < int32(2) {
		v21 = v13
		v22 = l3
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v21 = v20
	v22 = l3
	goto L1
L10:
	;
	return
L11:
	;
	v59 = F_palloc(m, int32(32))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L29
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L13:
	;
	if v22 != 0 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v32 = F_equal(m, v21, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	if v32 == int32(0) {
		v24 = v28
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
	if v36 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)))
	if v47 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v28)+16)) = l4
	v40 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)) = uint8(v40)
	return
L22:
	;
	goto L23
L23:
	;
	v42 = *(*float64)(unsafe.Add(mBase, uint32(v28)+16))
	if base.F64_gt(v42, l4) == int32(0) {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v28)+16)) = l4
	return
L25:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v28)+24)) = l4
	v51 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)) = uint8(v51)
	return
L26:
	;
	goto L27
L27:
	;
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v28)+24))
	if base.F64_gt(v53, l4) == int32(0) {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v28)+24)) = l4
	return
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v21
	if v22 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v64 = int32(16)
	goto L32
L31:
	;
	v64 = int32(24)
	goto L32
L32:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v59+v64))) = l4
	v68 = v22 ^ int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+9)) = uint8(v68)
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+8)) = uint8(v22)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v71
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v59
	goto L10
}
func F_addRangeTableEntryForJoin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v11 = l10
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v19 = F_palloc0(m, int32(136))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(101)
		if l5 != 0 {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
			if int32(_a_F_addRangeTableEntryForJoin_0) <= v25 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(_a_F_addRangeTableEntryForJoin_1)
						F_errmsg(m, int32(_a_F_addRangeTableEntryForJoin_2), v16+int32(16))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_addRangeTableEntryForJoin_3), int32(2260), int32(_a_F_addRangeTableEntryForJoin_4))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = l8
				*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = l7
				*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = l6
				*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = l5
				*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v19)+12)) = int64(2)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l9
				if l9 != 0 {
					v39 = F_copyObjectImpl(m, l9)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v45 = v39
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
						if v47 != 0 {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
							v49 = v48
						} else {
							v49 = int32(0)
						}
						if l1 != 0 {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v52 = v50
						} else {
							v52 = int32(0)
						}
						if v49 < v52 {
							v54 = F_list_copy_tail(m, l1, v49)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v56 = F_list_concat(m, v47, v54)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v56
									if l1 != 0 {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v61 = v59
									} else {
										v61 = int32(0)
									}
									if v61 < v49 {
										v63 = int32(0)
										F_errstart_cold(m, int32(21), v63)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(_a_F_addRangeTableEntryForJoin_5))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
												if l1 != 0 {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
													v73 = v72
												} else {
													v73 = v63
												}
												*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v49
												*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
												*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
												F_errmsg(m, int32(_a_F_addRangeTableEntryForJoin_6), v16)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_addRangeTableEntryForJoin_3), int32(2285), int32(_a_F_addRangeTableEntryForJoin_4))
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
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+125)) = uint8(v11)
										v86 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+124)) = uint8(v86)
										*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v45
										v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v91 = F_lappend(m, v90, v19)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91
											v95 = F_palloc(m, int32(28))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v19
												*(*int32)(unsafe.Add(mBase, uint32(v95))) = v97
												v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v102 != 0 {
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
													v104 = v103
												} else {
													v104 = v86
												}
												*(*int64)(unsafe.Add(mBase, uint32(v95)+20)) = int64(16777473)
												*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v104
												m.G0 = v16 + int32(32)
												return v95
											}
										}
									}
								}
							}
						} else {
							if l1 != 0 {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v61 = v59
							} else {
								v61 = int32(0)
							}
							if v61 < v49 {
								v63 = int32(0)
								F_errstart_cold(m, int32(21), v63)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(_a_F_addRangeTableEntryForJoin_5))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
										if l1 != 0 {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v73 = v72
										} else {
											v73 = v63
										}
										*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v49
										*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
										*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
										F_errmsg(m, int32(_a_F_addRangeTableEntryForJoin_6), v16)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_addRangeTableEntryForJoin_3), int32(2285), int32(_a_F_addRangeTableEntryForJoin_4))
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
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+125)) = uint8(v11)
								v86 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+124)) = uint8(v86)
								*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v45
								v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v91 = F_lappend(m, v90, v19)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91
									v95 = F_palloc(m, int32(28))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v19
										*(*int32)(unsafe.Add(mBase, uint32(v95))) = v97
										v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v102 != 0 {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
											v104 = v103
										} else {
											v104 = v86
										}
										*(*int64)(unsafe.Add(mBase, uint32(v95)+20)) = int64(16777473)
										*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v104
										m.G0 = v16 + int32(32)
										return v95
									}
								}
							}
						}
					}
				} else {
					v43 = F_makeAlias(m, int32(_a_F_addRangeTableEntryForJoin_7), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = v43
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
						if v47 != 0 {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
							v49 = v48
						} else {
							v49 = int32(0)
						}
						if l1 != 0 {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v52 = v50
						} else {
							v52 = int32(0)
						}
						if v49 < v52 {
							v54 = F_list_copy_tail(m, l1, v49)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v56 = F_list_concat(m, v47, v54)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v56
									if l1 != 0 {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v61 = v59
									} else {
										v61 = int32(0)
									}
									if v61 < v49 {
										v63 = int32(0)
										F_errstart_cold(m, int32(21), v63)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(_a_F_addRangeTableEntryForJoin_5))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
												if l1 != 0 {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
													v73 = v72
												} else {
													v73 = v63
												}
												*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v49
												*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
												*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
												F_errmsg(m, int32(_a_F_addRangeTableEntryForJoin_6), v16)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_addRangeTableEntryForJoin_3), int32(2285), int32(_a_F_addRangeTableEntryForJoin_4))
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
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+125)) = uint8(v11)
										v86 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+124)) = uint8(v86)
										*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v45
										v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v91 = F_lappend(m, v90, v19)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91
											v95 = F_palloc(m, int32(28))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v19
												*(*int32)(unsafe.Add(mBase, uint32(v95))) = v97
												v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												if v102 != 0 {
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
													v104 = v103
												} else {
													v104 = v86
												}
												*(*int64)(unsafe.Add(mBase, uint32(v95)+20)) = int64(16777473)
												*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v104
												m.G0 = v16 + int32(32)
												return v95
											}
										}
									}
								}
							}
						} else {
							if l1 != 0 {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v61 = v59
							} else {
								v61 = int32(0)
							}
							if v61 < v49 {
								v63 = int32(0)
								F_errstart_cold(m, int32(21), v63)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(_a_F_addRangeTableEntryForJoin_5))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
										if l1 != 0 {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v73 = v72
										} else {
											v73 = v63
										}
										*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v49
										*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
										*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
										F_errmsg(m, int32(_a_F_addRangeTableEntryForJoin_6), v16)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_addRangeTableEntryForJoin_3), int32(2285), int32(_a_F_addRangeTableEntryForJoin_4))
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
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+125)) = uint8(v11)
								v86 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+124)) = uint8(v86)
								*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v45
								v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v91 = F_lappend(m, v90, v19)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91
									v95 = F_palloc(m, int32(28))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v19
										*(*int32)(unsafe.Add(mBase, uint32(v95))) = v97
										v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v102 != 0 {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
											v104 = v103
										} else {
											v104 = v86
										}
										*(*int64)(unsafe.Add(mBase, uint32(v95)+20)) = int64(16777473)
										*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v104
										m.G0 = v16 + int32(32)
										return v95
									}
								}
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = l8
			*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = l7
			*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = l6
			*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v19)+12)) = int64(2)
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l9
			if l9 != 0 {
				v39 = F_copyObjectImpl(m, l9)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v45 = v39
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					if v47 != 0 {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
						v49 = v48
					} else {
						v49 = int32(0)
					}
					if l1 != 0 {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v52 = v50
					} else {
						v52 = int32(0)
					}
					if v49 < v52 {
						v54 = F_list_copy_tail(m, l1, v49)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v56 = F_list_concat(m, v47, v54)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v56
								if l1 != 0 {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v61 = v59
								} else {
									v61 = int32(0)
								}
								if v61 < v49 {
									v63 = int32(0)
									F_errstart_cold(m, int32(21), v63)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(_a_F_addRangeTableEntryForJoin_5))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
											if l1 != 0 {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												v73 = v72
											} else {
												v73 = v63
											}
											*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v49
											*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
											*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
											F_errmsg(m, int32(_a_F_addRangeTableEntryForJoin_6), v16)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_addRangeTableEntryForJoin_3), int32(2285), int32(_a_F_addRangeTableEntryForJoin_4))
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
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+125)) = uint8(v11)
									v86 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+124)) = uint8(v86)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v45
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v91 = F_lappend(m, v90, v19)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91
										v95 = F_palloc(m, int32(28))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int32(0)
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v19
											*(*int32)(unsafe.Add(mBase, uint32(v95))) = v97
											v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v102 != 0 {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
												v104 = v103
											} else {
												v104 = v86
											}
											*(*int64)(unsafe.Add(mBase, uint32(v95)+20)) = int64(16777473)
											*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v104
											m.G0 = v16 + int32(32)
											return v95
										}
									}
								}
							}
						}
					} else {
						if l1 != 0 {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v61 = v59
						} else {
							v61 = int32(0)
						}
						if v61 < v49 {
							v63 = int32(0)
							F_errstart_cold(m, int32(21), v63)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(_a_F_addRangeTableEntryForJoin_5))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
									if l1 != 0 {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v73 = v72
									} else {
										v73 = v63
									}
									*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v49
									*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
									*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
									F_errmsg(m, int32(_a_F_addRangeTableEntryForJoin_6), v16)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_addRangeTableEntryForJoin_3), int32(2285), int32(_a_F_addRangeTableEntryForJoin_4))
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
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v19)+125)) = uint8(v11)
							v86 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v19)+124)) = uint8(v86)
							*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v45
							v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v91 = F_lappend(m, v90, v19)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91
								v95 = F_palloc(m, int32(28))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v19
									*(*int32)(unsafe.Add(mBase, uint32(v95))) = v97
									v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v102 != 0 {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
										v104 = v103
									} else {
										v104 = v86
									}
									*(*int64)(unsafe.Add(mBase, uint32(v95)+20)) = int64(16777473)
									*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v104
									m.G0 = v16 + int32(32)
									return v95
								}
							}
						}
					}
				}
			} else {
				v43 = F_makeAlias(m, int32(_a_F_addRangeTableEntryForJoin_7), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = v43
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					if v47 != 0 {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
						v49 = v48
					} else {
						v49 = int32(0)
					}
					if l1 != 0 {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v52 = v50
					} else {
						v52 = int32(0)
					}
					if v49 < v52 {
						v54 = F_list_copy_tail(m, l1, v49)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v56 = F_list_concat(m, v47, v54)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v56
								if l1 != 0 {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v61 = v59
								} else {
									v61 = int32(0)
								}
								if v61 < v49 {
									v63 = int32(0)
									F_errstart_cold(m, int32(21), v63)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(_a_F_addRangeTableEntryForJoin_5))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
											if l1 != 0 {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												v73 = v72
											} else {
												v73 = v63
											}
											*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v49
											*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
											*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
											F_errmsg(m, int32(_a_F_addRangeTableEntryForJoin_6), v16)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_addRangeTableEntryForJoin_3), int32(2285), int32(_a_F_addRangeTableEntryForJoin_4))
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
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+125)) = uint8(v11)
									v86 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+124)) = uint8(v86)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v45
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v91 = F_lappend(m, v90, v19)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91
										v95 = F_palloc(m, int32(28))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int32(0)
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v19
											*(*int32)(unsafe.Add(mBase, uint32(v95))) = v97
											v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v102 != 0 {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
												v104 = v103
											} else {
												v104 = v86
											}
											*(*int64)(unsafe.Add(mBase, uint32(v95)+20)) = int64(16777473)
											*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v104
											m.G0 = v16 + int32(32)
											return v95
										}
									}
								}
							}
						}
					} else {
						if l1 != 0 {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v61 = v59
						} else {
							v61 = int32(0)
						}
						if v61 < v49 {
							v63 = int32(0)
							F_errstart_cold(m, int32(21), v63)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(_a_F_addRangeTableEntryForJoin_5))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
									if l1 != 0 {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v73 = v72
									} else {
										v73 = v63
									}
									*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v49
									*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
									*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
									F_errmsg(m, int32(_a_F_addRangeTableEntryForJoin_6), v16)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_addRangeTableEntryForJoin_3), int32(2285), int32(_a_F_addRangeTableEntryForJoin_4))
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
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v19)+125)) = uint8(v11)
							v86 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v19)+124)) = uint8(v86)
							*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v45
							v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v91 = F_lappend(m, v90, v19)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91
								v95 = F_palloc(m, int32(28))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v19
									*(*int32)(unsafe.Add(mBase, uint32(v95))) = v97
									v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v102 != 0 {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
										v104 = v103
									} else {
										v104 = v86
									}
									*(*int64)(unsafe.Add(mBase, uint32(v95)+20)) = int64(16777473)
									*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v104
									m.G0 = v16 + int32(32)
									return v95
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_addRangeTableEntryForRelation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	v5 = l4
	v6 = l5
	v12 = F_palloc0(m, int32(136))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(101)
		if l3 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			v22 = v18
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
			v22 = v19 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l3
		v24 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v24
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)) = uint8(v5)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v26
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+119)))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)) = uint8(v30)
		v34 = F_makeAlias(m, v22, v24)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v34
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
			F_buildRelationAliases(m, v37, l3, v34)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v12)+125)) = uint8(v6)
				v41 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v12)+124)) = uint8(v41)
				v44 = F_palloc0(m, int32(40))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(102)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v48
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)))
					*(*uint8)(unsafe.Add(mBase, uint32(v44)+8)) = uint8(v50)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v53 = F_lappend(m, v52, v44)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v53
						if v53 != 0 {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
							v58 = v56
						} else {
							v58 = int32(0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v58
						*(*int64)(unsafe.Add(mBase, uint32(v44)+16)) = int64(2)
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v63 = F_lappend(m, v62, v12)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v63
							if v63 != 0 {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
								v67 = v66
							} else {
								v67 = int32(0)
							}
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
							v72 = F_palloc0(m, v69<<(uint(int32(5))%32))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								if int32(0) < v69 {
									v78 = int32(0)
									for {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
										v93 = v68 + v87<<(uint(int32(4))%32) + v78*int32(100)
										v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+111)))
										if v94 == int32(0) {
											v99 = v72 + v78<<(uint(int32(5))%32)
											v101 = v78 + int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v99)+4)) = uint16(v101)
											*(*int32)(unsafe.Add(mBase, uint32(v99))) = v67
											v105 = v93 + int32(20)
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+68))
											*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v106
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)+76))
											*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v108
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v105)+96))
											*(*uint16)(unsafe.Add(mBase, uint32(v99)+28)) = uint16(v101)
											*(*int32)(unsafe.Add(mBase, uint32(v99)+24)) = v67
											*(*int32)(unsafe.Add(mBase, uint32(v99)+16)) = v110
										} else {
										}
										v118 = v78 + int32(1)
										if v118 != v69 {
											v78 = v118
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								v131 = F_palloc(m, int32(28))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v131)+20)) = int64(16777473)
									*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v72
									*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = v44
									*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v67
									*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = v12
									*(*int32)(unsafe.Add(mBase, uint32(v131))) = v133
									return v131
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_addRangeTableEntryForTableFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v4 = l3
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = F_palloc0(m, int32(136))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(101)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		if v21 != 0 {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			if int32(1665) <= v22 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(17039621))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(1664)
						F_errmsg(m, int32(_a_F_addRangeTableEntryForTableFunc_0), v12+int32(16))
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return int32(0)
						} else {
							v129 = F_exprLocation(m, l1)
							mBase = m.M
							F_parser_errposition(m, l0, v129)
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_addRangeTableEntryForTableFunc_1), int32(2086), int32(_a_F_addRangeTableEntryForTableFunc_2))
								mBase = m.M
								v136 = m.ExcPending
								if v136 != 0 {
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
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v15)+12)) = int64(4)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v30
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = v32
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v34
				if l2 != 0 {
					v37 = F_copyObjectImpl(m, l2)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v48 = v37
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
						if v50 != 0 {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
							v52 = v51
						} else {
							v52 = int32(0)
						}
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						if v53 != 0 {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
							v56 = v54
						} else {
							v56 = int32(0)
						}
						if v52 < v56 {
							v58 = F_list_copy_tail(m, v53, v52)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = F_list_concat(m, v50, v58)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v60
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v64 = v63
									if v64 != 0 {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
										v67 = v65
									} else {
										v67 = int32(0)
									}
									if v67 < v52 {
										v69 = int32(0)
										F_errstart_cold(m, int32(21), v69)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(_a_F_addRangeTableEntryForTableFunc_3))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
												if v78 != 0 {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
													v80 = v79
												} else {
													v80 = v69
												}
												*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v52
												*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v80
												if v77 != 0 {
													v85 = int32(_a_F_addRangeTableEntryForTableFunc_4)
												} else {
													v85 = int32(_a_F_addRangeTableEntryForTableFunc_5)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
												F_errmsg(m, int32(_a_F_addRangeTableEntryForTableFunc_6), v12)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_addRangeTableEntryForTableFunc_1), int32(2115), int32(_a_F_addRangeTableEntryForTableFunc_2))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v95 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+125)) = uint8(v95)
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+124)) = uint8(v4)
										*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v48
										v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v100 = F_lappend(m, v99, v15)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v100
											if v100 != 0 {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
												v105 = v103
											} else {
												v105 = int32(0)
											}
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+104))
											v109 = F_buildNSItemFromLists(m, v15, v105, v106, v107, v108)
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												m.G0 = v12 + int32(32)
												return v109
											}
										}
									}
								}
							}
						} else {
							v64 = v53
							if v64 != 0 {
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
								v67 = v65
							} else {
								v67 = int32(0)
							}
							if v67 < v52 {
								v69 = int32(0)
								F_errstart_cold(m, int32(21), v69)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(_a_F_addRangeTableEntryForTableFunc_3))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										if v78 != 0 {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
											v80 = v79
										} else {
											v80 = v69
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v52
										*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v80
										if v77 != 0 {
											v85 = int32(_a_F_addRangeTableEntryForTableFunc_4)
										} else {
											v85 = int32(_a_F_addRangeTableEntryForTableFunc_5)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
										F_errmsg(m, int32(_a_F_addRangeTableEntryForTableFunc_6), v12)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_addRangeTableEntryForTableFunc_1), int32(2115), int32(_a_F_addRangeTableEntryForTableFunc_2))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v95 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+125)) = uint8(v95)
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+124)) = uint8(v4)
								*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v48
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v100 = F_lappend(m, v99, v15)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v100
									if v100 != 0 {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
										v105 = v103
									} else {
										v105 = int32(0)
									}
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+104))
									v109 = F_buildNSItemFromLists(m, v15, v105, v106, v107, v108)
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(32)
										return v109
									}
								}
							}
						}
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v41 != 0 {
						v42 = int32(_a_F_addRangeTableEntryForTableFunc_7)
					} else {
						v42 = int32(_a_F_addRangeTableEntryForTableFunc_8)
					}
					v43 = F_pstrdup(m, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v46 = F_makeAlias(m, v43, int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = v46
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
							if v50 != 0 {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
								v52 = v51
							} else {
								v52 = int32(0)
							}
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							if v53 != 0 {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
								v56 = v54
							} else {
								v56 = int32(0)
							}
							if v52 < v56 {
								v58 = F_list_copy_tail(m, v53, v52)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v60 = F_list_concat(m, v50, v58)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v60
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										v64 = v63
										if v64 != 0 {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
											v67 = v65
										} else {
											v67 = int32(0)
										}
										if v67 < v52 {
											v69 = int32(0)
											F_errstart_cold(m, int32(21), v69)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(_a_F_addRangeTableEntryForTableFunc_3))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
													v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
													if v78 != 0 {
														v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
														v80 = v79
													} else {
														v80 = v69
													}
													*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v52
													*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v80
													if v77 != 0 {
														v85 = int32(_a_F_addRangeTableEntryForTableFunc_4)
													} else {
														v85 = int32(_a_F_addRangeTableEntryForTableFunc_5)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
													F_errmsg(m, int32(_a_F_addRangeTableEntryForTableFunc_6), v12)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_addRangeTableEntryForTableFunc_1), int32(2115), int32(_a_F_addRangeTableEntryForTableFunc_2))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v95 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v15)+125)) = uint8(v95)
											*(*uint8)(unsafe.Add(mBase, uint32(v15)+124)) = uint8(v4)
											*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v48
											v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v100 = F_lappend(m, v99, v15)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v100
												if v100 != 0 {
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
													v105 = v103
												} else {
													v105 = int32(0)
												}
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
												v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+104))
												v109 = F_buildNSItemFromLists(m, v15, v105, v106, v107, v108)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													m.G0 = v12 + int32(32)
													return v109
												}
											}
										}
									}
								}
							} else {
								v64 = v53
								if v64 != 0 {
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
									v67 = v65
								} else {
									v67 = int32(0)
								}
								if v67 < v52 {
									v69 = int32(0)
									F_errstart_cold(m, int32(21), v69)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(_a_F_addRangeTableEntryForTableFunc_3))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											if v78 != 0 {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
												v80 = v79
											} else {
												v80 = v69
											}
											*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v52
											*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v80
											if v77 != 0 {
												v85 = int32(_a_F_addRangeTableEntryForTableFunc_4)
											} else {
												v85 = int32(_a_F_addRangeTableEntryForTableFunc_5)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
											F_errmsg(m, int32(_a_F_addRangeTableEntryForTableFunc_6), v12)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_addRangeTableEntryForTableFunc_1), int32(2115), int32(_a_F_addRangeTableEntryForTableFunc_2))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v95 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+125)) = uint8(v95)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+124)) = uint8(v4)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v48
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v100 = F_lappend(m, v99, v15)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v100
										if v100 != 0 {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
											v105 = v103
										} else {
											v105 = int32(0)
										}
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+104))
										v109 = F_buildNSItemFromLists(m, v15, v105, v106, v107, v108)
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											m.G0 = v12 + int32(32)
											return v109
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v15)+12)) = int64(4)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = v32
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v34
			if l2 != 0 {
				v37 = F_copyObjectImpl(m, l2)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v48 = v37
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
					if v50 != 0 {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
						v52 = v51
					} else {
						v52 = int32(0)
					}
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					if v53 != 0 {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
						v56 = v54
					} else {
						v56 = int32(0)
					}
					if v52 < v56 {
						v58 = F_list_copy_tail(m, v53, v52)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = F_list_concat(m, v50, v58)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v60
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v64 = v63
								if v64 != 0 {
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
									v67 = v65
								} else {
									v67 = int32(0)
								}
								if v67 < v52 {
									v69 = int32(0)
									F_errstart_cold(m, int32(21), v69)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(_a_F_addRangeTableEntryForTableFunc_3))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											if v78 != 0 {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
												v80 = v79
											} else {
												v80 = v69
											}
											*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v52
											*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v80
											if v77 != 0 {
												v85 = int32(_a_F_addRangeTableEntryForTableFunc_4)
											} else {
												v85 = int32(_a_F_addRangeTableEntryForTableFunc_5)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
											F_errmsg(m, int32(_a_F_addRangeTableEntryForTableFunc_6), v12)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_addRangeTableEntryForTableFunc_1), int32(2115), int32(_a_F_addRangeTableEntryForTableFunc_2))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v95 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+125)) = uint8(v95)
									*(*uint8)(unsafe.Add(mBase, uint32(v15)+124)) = uint8(v4)
									*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v48
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v100 = F_lappend(m, v99, v15)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v100
										if v100 != 0 {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
											v105 = v103
										} else {
											v105 = int32(0)
										}
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+104))
										v109 = F_buildNSItemFromLists(m, v15, v105, v106, v107, v108)
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											m.G0 = v12 + int32(32)
											return v109
										}
									}
								}
							}
						}
					} else {
						v64 = v53
						if v64 != 0 {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
							v67 = v65
						} else {
							v67 = int32(0)
						}
						if v67 < v52 {
							v69 = int32(0)
							F_errstart_cold(m, int32(21), v69)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(_a_F_addRangeTableEntryForTableFunc_3))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									if v78 != 0 {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
										v80 = v79
									} else {
										v80 = v69
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v52
									*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v80
									if v77 != 0 {
										v85 = int32(_a_F_addRangeTableEntryForTableFunc_4)
									} else {
										v85 = int32(_a_F_addRangeTableEntryForTableFunc_5)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
									F_errmsg(m, int32(_a_F_addRangeTableEntryForTableFunc_6), v12)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_addRangeTableEntryForTableFunc_1), int32(2115), int32(_a_F_addRangeTableEntryForTableFunc_2))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v95 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+125)) = uint8(v95)
							*(*uint8)(unsafe.Add(mBase, uint32(v15)+124)) = uint8(v4)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v48
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v100 = F_lappend(m, v99, v15)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v100
								if v100 != 0 {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
									v105 = v103
								} else {
									v105 = int32(0)
								}
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+104))
								v109 = F_buildNSItemFromLists(m, v15, v105, v106, v107, v108)
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return int32(0)
								} else {
									m.G0 = v12 + int32(32)
									return v109
								}
							}
						}
					}
				}
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v41 != 0 {
					v42 = int32(_a_F_addRangeTableEntryForTableFunc_7)
				} else {
					v42 = int32(_a_F_addRangeTableEntryForTableFunc_8)
				}
				v43 = F_pstrdup(m, v42)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v46 = F_makeAlias(m, v43, int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v48 = v46
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
						if v50 != 0 {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
							v52 = v51
						} else {
							v52 = int32(0)
						}
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						if v53 != 0 {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
							v56 = v54
						} else {
							v56 = int32(0)
						}
						if v52 < v56 {
							v58 = F_list_copy_tail(m, v53, v52)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = F_list_concat(m, v50, v58)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v60
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									v64 = v63
									if v64 != 0 {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
										v67 = v65
									} else {
										v67 = int32(0)
									}
									if v67 < v52 {
										v69 = int32(0)
										F_errstart_cold(m, int32(21), v69)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(_a_F_addRangeTableEntryForTableFunc_3))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
												if v78 != 0 {
													v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
													v80 = v79
												} else {
													v80 = v69
												}
												*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v52
												*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v80
												if v77 != 0 {
													v85 = int32(_a_F_addRangeTableEntryForTableFunc_4)
												} else {
													v85 = int32(_a_F_addRangeTableEntryForTableFunc_5)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
												F_errmsg(m, int32(_a_F_addRangeTableEntryForTableFunc_6), v12)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_addRangeTableEntryForTableFunc_1), int32(2115), int32(_a_F_addRangeTableEntryForTableFunc_2))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v95 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+125)) = uint8(v95)
										*(*uint8)(unsafe.Add(mBase, uint32(v15)+124)) = uint8(v4)
										*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v48
										v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v100 = F_lappend(m, v99, v15)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v100
											if v100 != 0 {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
												v105 = v103
											} else {
												v105 = int32(0)
											}
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+104))
											v109 = F_buildNSItemFromLists(m, v15, v105, v106, v107, v108)
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												m.G0 = v12 + int32(32)
												return v109
											}
										}
									}
								}
							}
						} else {
							v64 = v53
							if v64 != 0 {
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
								v67 = v65
							} else {
								v67 = int32(0)
							}
							if v67 < v52 {
								v69 = int32(0)
								F_errstart_cold(m, int32(21), v69)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(_a_F_addRangeTableEntryForTableFunc_3))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
										if v78 != 0 {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
											v80 = v79
										} else {
											v80 = v69
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v52
										*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v80
										if v77 != 0 {
											v85 = int32(_a_F_addRangeTableEntryForTableFunc_4)
										} else {
											v85 = int32(_a_F_addRangeTableEntryForTableFunc_5)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
										F_errmsg(m, int32(_a_F_addRangeTableEntryForTableFunc_6), v12)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_addRangeTableEntryForTableFunc_1), int32(2115), int32(_a_F_addRangeTableEntryForTableFunc_2))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v95 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+125)) = uint8(v95)
								*(*uint8)(unsafe.Add(mBase, uint32(v15)+124)) = uint8(v4)
								*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v48
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v100 = F_lappend(m, v99, v15)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v100
									if v100 != 0 {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
										v105 = v103
									} else {
										v105 = int32(0)
									}
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v15)+104))
									v109 = F_buildNSItemFromLists(m, v15, v105, v106, v107, v108)
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(32)
										return v109
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
func F_compute_range_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 float64
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v150 int64
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 float64
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 float64
	_ = v169
	var v170 float64
	_ = v170
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 float64
	_ = v191
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 float64
	_ = v201
	var v208 float32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v375 float64
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v457 int32
	_ = v457
	v5 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(32)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+13)))
	if v31 == int32(109) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+296))
	v35 = v34
	goto L3
L2:
	;
	v35 = v30
	goto L3
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+272))
	v38 = l2 << (uint(int32(3)) % 32)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v41 = F_palloc(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v43 = F_palloc(m, v38)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if l2 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	m.G0 = v27 + int32(32)
	return
L9:
	;
	v59 = int32(0)
	v60 = v5
	v61 = v5
	v63 = v5
	v64 = v5
	v73 = float64(0)
	goto L10
L10:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	if int32(0) < v190 {
		goto L47
	} else {
		goto L48
	}
L12:
	;
	v81 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v59, v27+int32(31))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+31)))
	if v83 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v194 = v59 + int32(1)
	if v194 != l2 {
		v59 = v194
		v60 = v186
		v61 = v187
		v63 = v189
		v64 = v190
		v73 = v191
		goto L10
	} else {
		goto L46
	}
L15:
	;
	v186 = v60
	v187 = v61 + int32(1)
	v189 = v63
	v190 = v64
	v191 = v73
	goto L14
L16:
	;
	goto L17
L17:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v88 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v114 = F_pg_detoast_datum(m, v81)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L27
	}
L19:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if base.Ui32((v92-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v112 = int32(6)
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v104 = int32(1)
	if v88&v104 != 0 {
		v112 = int32(base.Ui32(v88) >> (uint(v104) % 32))
		goto L18
	} else {
		goto L26
	}
L22:
	;
	v99 = int32(18)
	if v92 == v99 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v103 = v99
	goto L25
L24:
	;
	v103 = int32(2)
	goto L25
L25:
	;
	v112 = v103
	goto L18
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v112 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
	goto L18
L27:
	;
	if base.B2i32(v31 != int32(109)) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v186 = v178
	v187 = v61
	v189 = v180
	v190 = v64 + int32(1)
	v191 = base.F64_add(v73, base.F64_convert_i32_u(v112))
	goto L14
L29:
	;
	v178 = v60
	v180 = v63 + int32(1)
	goto L28
L30:
	;
	v145 = v60 << (uint(int32(3)) % 32)
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v39+v145))) = v147
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v27)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v145+v41))) = v150
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+20)))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+12)))
	v158 = (v154 | v155) & int32(1)
	if v158 != 0 {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	if v118 == int32(0) {
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_range_deserialize(m, v35, v114, v27+int32(16), v27+int32(8), v27+int32(30))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L37
	}
L34:
	;
	F_multirange_get_bounds(m, v35, v114, int32(0), v27+int32(16), v27)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	F_multirange_get_bounds(m, v35, v114, v126-int32(1), v27, v27+int32(8))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v133 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+30)) = uint8(v133)
	goto L30
L37:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+30)))
	if v143 != 0 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	goto L30
L39:
	;
	v159 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L41
L40:
	;
	v159 = float64(1)
	goto L41
L41:
	;
	if v158|base.B2i32(v36 == int32(0)) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v170 = v159
	goto L44
L43:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v35)+208))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v167 = F_FunctionCall2Coll(m, v35+int32(268), v164, v165, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L45
	}
L44:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v145+v43))) = v170
	v178 = v60 + int32(1)
	v180 = v63
	goto L28
L45:
	;
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v167)))
	v170 = v169
	goto L44
L46:
	;
	goto L11
L47:
	;
	v199 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v199)
	v201 = base.F64_convert_i32_u(v190)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = base.I32_trunc_sat_f64_s(base.F64_div(v191, v201))
	v208 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v187), base.F64_convert_i32_s(l2)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v208
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = base.F32_neg(base.F32_sub(float32(1), v208))
	v215 = int32(_a_F_compute_range_stats_0)
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_compute_range_stats[0]))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_range_stats[0])) = v218
	if int32(2) <= v186 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	goto L49
L49:
	;
	if v187 <= int32(0) {
		goto L8
	} else {
		goto L82
	}
L50:
	;
	v419 = l0 + v404<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v419)+164)) = v401
	v421 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v419)+84)) = v421
	*(*int32)(unsafe.Add(mBase, uint32(v419)+64)) = int32(672)
	*(*int32)(unsafe.Add(mBase, uint32(v419)+184)) = int32(701)
	*(*int32)(unsafe.Add(mBase, uint32(v419)+144)) = v397
	v430 = l0 + v404<<(uint(int32(1))%32)
	v431 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v430)+204)) = uint16(v431)
	v433 = l0 + v404
	v434 = int32(100)
	*(*uint8)(unsafe.Add(mBase, uint32(v433)+219)) = uint8(v434)
	*(*uint8)(unsafe.Add(mBase, uint32(v433)+214)) = uint8(v421)
	v439 = F_palloc(m, int32(4))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L81
	}
L51:
	;
	F_qsort_interruptible(m, v39, v186, int32(8), int32(1475), v35)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v390 = F_palloc(m, int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L80
	}
L54:
	;
	F_qsort_interruptible(m, v41, v186, int32(8), int32(1475), v35)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v230 = int32(1)
	v232 = v186 - v230
	if v29 < v186 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v236 = v29 + v230
	goto L58
L57:
	;
	v236 = v186
	goto L58
L58:
	;
	v238 = v236 - int32(1)
	v239 = base.I32_div_s(v232, v238)
	v241 = v232 - v239*v238
	v243 = v236 << (uint(int32(2)) % 32)
	v244 = F_palloc(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v246 = int32(0)
	v247 = base.B2i32(v236 <= v246)
	if v247 == v246 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v250 = int32(0)
	v254 = v250
	v255 = v250
	v260 = v250
	goto L63
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v244
	v324 = int32(7)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v324)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v236
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v327
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+204)) = uint16(v329)
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+214)) = uint8(v331)
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+11)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+219)) = uint8(v333)
	v335 = int32(0)
	F_qsort_interruptible(m, v43, v186, int32(8), int32(1476), v335)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L70
	}
L63:
	;
	v281 = v255 << (uint(int32(3)) % 32)
	v284 = int32(0)
	v286 = F_range_serialize(m, v35, v39+v281, v41+v281, v284, v284)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L65
	}
L64:
	;
	goto L62
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244+v260<<(uint(int32(2))%32)))) = v286
	v289 = v254 + v241
	v290 = base.B2i32(v238 <= v289)
	if v238 <= v289 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v294 = v238
	goto L68
L67:
	;
	v294 = int32(0)
	goto L68
L68:
	;
	v297 = v260 + int32(1)
	if v297 != v236 {
		v254 = v289 - v294
		v255 = v290 + (v255 + v239)
		v260 = v297
		goto L63
	} else {
		goto L69
	}
L69:
	;
	goto L64
L70:
	;
	v341 = F_palloc(m, v243)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	if v236 <= v246 {
		v397 = v236
		v401 = v341
		v404 = v230
		goto L50
	} else {
		goto L72
	}
L72:
	;
	v343 = int32(0)
	v347 = v343
	v350 = v343
	v352 = v335
	goto L73
L73:
	;
	v375 = *(*float64)(unsafe.Add(mBase, uint32(v43+v347<<(uint(int32(3))%32))))
	v376 = F_Float8GetDatum(m, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L4
	} else {
		goto L75
	}
L74:
	;
	v397 = v236
	v401 = v341
	v404 = v230
	goto L50
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v341+v352<<(uint(int32(2))%32)))) = v376
	v379 = v350 + v241
	v380 = base.B2i32(v238 <= v379)
	if v238 <= v379 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v384 = v238
	goto L78
L77:
	;
	v384 = int32(0)
	goto L78
L78:
	;
	v387 = v352 + int32(1)
	if v387 != v236 {
		v347 = v380 + (v347 + v239)
		v350 = v379 - v384
		v352 = v387
		goto L73
	} else {
		goto L79
	}
L79:
	;
	goto L74
L80:
	;
	v397 = int32(0)
	v401 = v390
	v404 = int32(0)
	goto L50
L81:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v439))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v189), v201))
	*(*int32)(unsafe.Add(mBase, uint32(v419)+104)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v419)+124)) = v439
	v447 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v430)+52)) = uint16(v447)
	*(*int32)(unsafe.Add(mBase, _c_F_compute_range_stats[0])) = v216
	goto L8
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(1065353216)
	v457 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v457)
	goto L8
}
func F_makeRangeVarFromQualifiedName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v92 int32
	_ = v92
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l1 == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L20
	}
L2:
	;
	v15 = int32(0)
	v17 = F_makeRangeVar(m, v15, v15, l2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v21 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	goto L1
L7:
	;
	v54 = int32(0)
	v56 = F_makeRangeVar(m, v54, v54, l2)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L16
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v29 = v5
	goto L9
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24+v29<<(uint(int32(2))%32))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v37 == int32(468) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_scanner_yyerror(m, int32(_a_F_makeRangeVarFromQualifiedName_0), l3)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L15
	}
L11:
	;
	v41 = v29 + int32(1)
	if v41 != v21 {
		v29 = v41
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L7
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v58 - int32(1) {
	case 0:
		goto L18
	case 1:
		goto L19
	default:
		goto L1
	}
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = v75
	m.G0 = v11 + int32(16)
	return v56
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v73 = v72
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = l0
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v73 = v66 + int32(4)
	goto L17
L20:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v96 = F_makeString(m, l0)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v98 = F_lcons(m, v96, l1)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v100 = F_NameListToString(m, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v100
	F_errmsg(m, int32(_a_F_makeRangeVarFromQualifiedName_1), v11)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	F_scanner_errposition(m, l2, l3)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_makeRangeVarFromQualifiedName_2), int32(_a_F_makeRangeVarFromQualifiedName_3), int32(_a_F_makeRangeVarFromQualifiedName_4))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_range_contained_by_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_range_contains_internal(m, l0, l2, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_range_contains(m *base.Module, l0 int32) int32 {
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
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
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
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = F_range_contains_internal(m, v32, v12, v17)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v33
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_contains_0), v9)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_contains_1), int32(1776), int32(_a_F_range_contains_2))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
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
							v33 = F_range_contains_internal(m, v32, v12, v17)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v33
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_contains_0), v9)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_contains_1), int32(1776), int32(_a_F_range_contains_2))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
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
						v33 = F_range_contains_internal(m, v32, v12, v17)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v33
						}
					}
				}
			}
		}
	}
}
func F_range_contains_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v10 == v11 {
		F_range_deserialize(m, l0, l1, v8+int32(40), v8+int32(32), v8+int32(31))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_range_deserialize(m, l0, l2, v8+int32(20), v8+int32(12), v8+int32(11))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)))
				if v31 != 0 {
					v161 = int32(1)
					m.G0 = v8 + int32(48)
					return v161
				} else {
					v33 = int32(0)
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
					if v34 != 0 {
						v161 = v33
						m.G0 = v8 + int32(48)
						return v161
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+44)))
						if v36 == int32(1) {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+46)))
							if v35&int32(1) != 0 {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
								if v39&int32(1)|base.B2i32(v39 == v44) != 0 {
									v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
									v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
									if v93 == int32(1) {
										v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
										if v92&int32(1) == int32(0) {
											v150 = int32(1)
											if v96&v150 == int32(0) {
												v161 = v150
											} else {
												v161 = int32(0)
											}
										} else {
											v101 = int32(1)
											v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
											if v102 == v96 {
												v161 = v101
											} else {
												if v96&int32(1) != 0 {
													v161 = int32(0)
												} else {
													v161 = v101
												}
											}
										}
										m.G0 = v8 + int32(48)
										return v161
									} else {
										if v92&int32(1) != 0 {
											v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
											if v108 == int32(0) {
												v161 = int32(0)
											} else {
												v161 = int32(1)
											}
											m.G0 = v8 + int32(48)
											return v161
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
											v116 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											v117 = F_FunctionCall2Coll(m, l0+int32(212), v114, v115, v116)
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												if v117 == int32(0) {
													v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
													v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
													if v122 == int32(0) {
														v125 = int32(1)
														v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
														if v121&v125 == int32(0) {
															v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v131 == v126 {
																v161 = v125
															} else {
																if v126&int32(1) == int32(0) {
																	v161 = int32(0)
																} else {
																	v161 = v125
																}
															}
														} else {
															if v126&int32(1) == int32(0) {
																v161 = int32(0)
															} else {
																v161 = v125
															}
														}
													} else {
														v141 = int32(1)
														if v121&v141 != 0 {
															v161 = v141
														} else {
															v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v144&int32(1) != 0 {
																v161 = int32(0)
															} else {
																v161 = v141
															}
														}
													}
												} else {
													if v117 < int32(0) {
														v161 = int32(0)
													} else {
														v161 = int32(1)
													}
												}
												m.G0 = v8 + int32(48)
												return v161
											}
										}
									}
								} else {
									v161 = v33
									m.G0 = v8 + int32(48)
									return v161
								}
							} else {
								if v39&int32(1) != 0 {
									v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
									v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
									if v93 == int32(1) {
										v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
										if v92&int32(1) == int32(0) {
											v150 = int32(1)
											if v96&v150 == int32(0) {
												v161 = v150
											} else {
												v161 = int32(0)
											}
										} else {
											v101 = int32(1)
											v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
											if v102 == v96 {
												v161 = v101
											} else {
												if v96&int32(1) != 0 {
													v161 = int32(0)
												} else {
													v161 = v101
												}
											}
										}
										m.G0 = v8 + int32(48)
										return v161
									} else {
										if v92&int32(1) != 0 {
											v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
											if v108 == int32(0) {
												v161 = int32(0)
											} else {
												v161 = int32(1)
											}
											m.G0 = v8 + int32(48)
											return v161
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
											v116 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											v117 = F_FunctionCall2Coll(m, l0+int32(212), v114, v115, v116)
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												if v117 == int32(0) {
													v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
													v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
													if v122 == int32(0) {
														v125 = int32(1)
														v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
														if v121&v125 == int32(0) {
															v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v131 == v126 {
																v161 = v125
															} else {
																if v126&int32(1) == int32(0) {
																	v161 = int32(0)
																} else {
																	v161 = v125
																}
															}
														} else {
															if v126&int32(1) == int32(0) {
																v161 = int32(0)
															} else {
																v161 = v125
															}
														}
													} else {
														v141 = int32(1)
														if v121&v141 != 0 {
															v161 = v141
														} else {
															v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v144&int32(1) != 0 {
																v161 = int32(0)
															} else {
																v161 = v141
															}
														}
													}
												} else {
													if v117 < int32(0) {
														v161 = int32(0)
													} else {
														v161 = int32(1)
													}
												}
												m.G0 = v8 + int32(48)
												return v161
											}
										}
									}
								} else {
									v161 = v33
									m.G0 = v8 + int32(48)
									return v161
								}
							}
						} else {
							if v35&int32(1) != 0 {
								v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
								if v51 == int32(0) {
									v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
									v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
									if v93 == int32(1) {
										v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
										if v92&int32(1) == int32(0) {
											v150 = int32(1)
											if v96&v150 == int32(0) {
												v161 = v150
											} else {
												v161 = int32(0)
											}
										} else {
											v101 = int32(1)
											v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
											if v102 == v96 {
												v161 = v101
											} else {
												if v96&int32(1) != 0 {
													v161 = int32(0)
												} else {
													v161 = v101
												}
											}
										}
										m.G0 = v8 + int32(48)
										return v161
									} else {
										if v92&int32(1) != 0 {
											v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
											if v108 == int32(0) {
												v161 = int32(0)
											} else {
												v161 = int32(1)
											}
											m.G0 = v8 + int32(48)
											return v161
										} else {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
											v116 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											v117 = F_FunctionCall2Coll(m, l0+int32(212), v114, v115, v116)
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												if v117 == int32(0) {
													v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
													v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
													if v122 == int32(0) {
														v125 = int32(1)
														v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
														if v121&v125 == int32(0) {
															v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v131 == v126 {
																v161 = v125
															} else {
																if v126&int32(1) == int32(0) {
																	v161 = int32(0)
																} else {
																	v161 = v125
																}
															}
														} else {
															if v126&int32(1) == int32(0) {
																v161 = int32(0)
															} else {
																v161 = v125
															}
														}
													} else {
														v141 = int32(1)
														if v121&v141 != 0 {
															v161 = v141
														} else {
															v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v144&int32(1) != 0 {
																v161 = int32(0)
															} else {
																v161 = v141
															}
														}
													}
												} else {
													if v117 < int32(0) {
														v161 = int32(0)
													} else {
														v161 = int32(1)
													}
												}
												m.G0 = v8 + int32(48)
												return v161
											}
										}
									}
								} else {
									v161 = v33
									m.G0 = v8 + int32(48)
									return v161
								}
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
								v59 = F_FunctionCall2Coll(m, l0+int32(212), v56, v57, v58)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									if v59 == int32(0) {
										v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+25)))
										v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+45)))
										if v64 == int32(0) {
											v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+46)))
											if v63&int32(1) == int32(0) {
												v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
												if base.B2i32(v67&int32(1) == int32(0))|base.B2i32(v67 == v76) != 0 {
													v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
													if v93 == int32(1) {
														v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
														if v92&int32(1) == int32(0) {
															v150 = int32(1)
															if v96&v150 == int32(0) {
																v161 = v150
															} else {
																v161 = int32(0)
															}
														} else {
															v101 = int32(1)
															v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v102 == v96 {
																v161 = v101
															} else {
																if v96&int32(1) != 0 {
																	v161 = int32(0)
																} else {
																	v161 = v101
																}
															}
														}
														m.G0 = v8 + int32(48)
														return v161
													} else {
														if v92&int32(1) != 0 {
															v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v108 == int32(0) {
																v161 = int32(0)
															} else {
																v161 = int32(1)
															}
															m.G0 = v8 + int32(48)
															return v161
														} else {
															v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
															v117 = F_FunctionCall2Coll(m, l0+int32(212), v114, v115, v116)
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return int32(0)
															} else {
																if v117 == int32(0) {
																	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
																	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
																	if v122 == int32(0) {
																		v125 = int32(1)
																		v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
																		if v121&v125 == int32(0) {
																			v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																			if v131 == v126 {
																				v161 = v125
																			} else {
																				if v126&int32(1) == int32(0) {
																					v161 = int32(0)
																				} else {
																					v161 = v125
																				}
																			}
																		} else {
																			if v126&int32(1) == int32(0) {
																				v161 = int32(0)
																			} else {
																				v161 = v125
																			}
																		}
																	} else {
																		v141 = int32(1)
																		if v121&v141 != 0 {
																			v161 = v141
																		} else {
																			v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																			if v144&int32(1) != 0 {
																				v161 = int32(0)
																			} else {
																				v161 = v141
																			}
																		}
																	}
																} else {
																	if v117 < int32(0) {
																		v161 = int32(0)
																	} else {
																		v161 = int32(1)
																	}
																}
																m.G0 = v8 + int32(48)
																return v161
															}
														}
													}
												} else {
													v161 = v33
													m.G0 = v8 + int32(48)
													return v161
												}
											} else {
												if v67&int32(1) == int32(0) {
													v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
													if v93 == int32(1) {
														v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
														if v92&int32(1) == int32(0) {
															v150 = int32(1)
															if v96&v150 == int32(0) {
																v161 = v150
															} else {
																v161 = int32(0)
															}
														} else {
															v101 = int32(1)
															v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v102 == v96 {
																v161 = v101
															} else {
																if v96&int32(1) != 0 {
																	v161 = int32(0)
																} else {
																	v161 = v101
																}
															}
														}
														m.G0 = v8 + int32(48)
														return v161
													} else {
														if v92&int32(1) != 0 {
															v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v108 == int32(0) {
																v161 = int32(0)
															} else {
																v161 = int32(1)
															}
															m.G0 = v8 + int32(48)
															return v161
														} else {
															v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
															v117 = F_FunctionCall2Coll(m, l0+int32(212), v114, v115, v116)
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return int32(0)
															} else {
																if v117 == int32(0) {
																	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
																	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
																	if v122 == int32(0) {
																		v125 = int32(1)
																		v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
																		if v121&v125 == int32(0) {
																			v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																			if v131 == v126 {
																				v161 = v125
																			} else {
																				if v126&int32(1) == int32(0) {
																					v161 = int32(0)
																				} else {
																					v161 = v125
																				}
																			}
																		} else {
																			if v126&int32(1) == int32(0) {
																				v161 = int32(0)
																			} else {
																				v161 = v125
																			}
																		}
																	} else {
																		v141 = int32(1)
																		if v121&v141 != 0 {
																			v161 = v141
																		} else {
																			v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																			if v144&int32(1) != 0 {
																				v161 = int32(0)
																			} else {
																				v161 = v141
																			}
																		}
																	}
																} else {
																	if v117 < int32(0) {
																		v161 = int32(0)
																	} else {
																		v161 = int32(1)
																	}
																}
																m.G0 = v8 + int32(48)
																return v161
															}
														}
													}
												} else {
													v161 = v33
													m.G0 = v8 + int32(48)
													return v161
												}
											}
										} else {
											if v63&int32(1) != 0 {
												v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
												v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
												if v93 == int32(1) {
													v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
													if v92&int32(1) == int32(0) {
														v150 = int32(1)
														if v96&v150 == int32(0) {
															v161 = v150
														} else {
															v161 = int32(0)
														}
													} else {
														v101 = int32(1)
														v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
														if v102 == v96 {
															v161 = v101
														} else {
															if v96&int32(1) != 0 {
																v161 = int32(0)
															} else {
																v161 = v101
															}
														}
													}
													m.G0 = v8 + int32(48)
													return v161
												} else {
													if v92&int32(1) != 0 {
														v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
														if v108 == int32(0) {
															v161 = int32(0)
														} else {
															v161 = int32(1)
														}
														m.G0 = v8 + int32(48)
														return v161
													} else {
														v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
														v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
														v117 = F_FunctionCall2Coll(m, l0+int32(212), v114, v115, v116)
														mBase = m.M
														v118 = m.ExcPending
														if v118 != 0 {
															return int32(0)
														} else {
															if v117 == int32(0) {
																v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
																v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
																if v122 == int32(0) {
																	v125 = int32(1)
																	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
																	if v121&v125 == int32(0) {
																		v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																		if v131 == v126 {
																			v161 = v125
																		} else {
																			if v126&int32(1) == int32(0) {
																				v161 = int32(0)
																			} else {
																				v161 = v125
																			}
																		}
																	} else {
																		if v126&int32(1) == int32(0) {
																			v161 = int32(0)
																		} else {
																			v161 = v125
																		}
																	}
																} else {
																	v141 = int32(1)
																	if v121&v141 != 0 {
																		v161 = v141
																	} else {
																		v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																		if v144&int32(1) != 0 {
																			v161 = int32(0)
																		} else {
																			v161 = v141
																		}
																	}
																}
															} else {
																if v117 < int32(0) {
																	v161 = int32(0)
																} else {
																	v161 = int32(1)
																}
															}
															m.G0 = v8 + int32(48)
															return v161
														}
													}
												}
											} else {
												v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
												if v85&int32(1) != 0 {
													v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
													if v93 == int32(1) {
														v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
														if v92&int32(1) == int32(0) {
															v150 = int32(1)
															if v96&v150 == int32(0) {
																v161 = v150
															} else {
																v161 = int32(0)
															}
														} else {
															v101 = int32(1)
															v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v102 == v96 {
																v161 = v101
															} else {
																if v96&int32(1) != 0 {
																	v161 = int32(0)
																} else {
																	v161 = v101
																}
															}
														}
														m.G0 = v8 + int32(48)
														return v161
													} else {
														if v92&int32(1) != 0 {
															v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v108 == int32(0) {
																v161 = int32(0)
															} else {
																v161 = int32(1)
															}
															m.G0 = v8 + int32(48)
															return v161
														} else {
															v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
															v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
															v117 = F_FunctionCall2Coll(m, l0+int32(212), v114, v115, v116)
															mBase = m.M
															v118 = m.ExcPending
															if v118 != 0 {
																return int32(0)
															} else {
																if v117 == int32(0) {
																	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
																	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
																	if v122 == int32(0) {
																		v125 = int32(1)
																		v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
																		if v121&v125 == int32(0) {
																			v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																			if v131 == v126 {
																				v161 = v125
																			} else {
																				if v126&int32(1) == int32(0) {
																					v161 = int32(0)
																				} else {
																					v161 = v125
																				}
																			}
																		} else {
																			if v126&int32(1) == int32(0) {
																				v161 = int32(0)
																			} else {
																				v161 = v125
																			}
																		}
																	} else {
																		v141 = int32(1)
																		if v121&v141 != 0 {
																			v161 = v141
																		} else {
																			v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																			if v144&int32(1) != 0 {
																				v161 = int32(0)
																			} else {
																				v161 = v141
																			}
																		}
																	}
																} else {
																	if v117 < int32(0) {
																		v161 = int32(0)
																	} else {
																		v161 = int32(1)
																	}
																}
																m.G0 = v8 + int32(48)
																return v161
															}
														}
													}
												} else {
													v161 = v33
													m.G0 = v8 + int32(48)
													return v161
												}
											}
										}
									} else {
										if int32(0) < v59 {
											v161 = v33
											m.G0 = v8 + int32(48)
											return v161
										} else {
											v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
											v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
											if v93 == int32(1) {
												v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
												if v92&int32(1) == int32(0) {
													v150 = int32(1)
													if v96&v150 == int32(0) {
														v161 = v150
													} else {
														v161 = int32(0)
													}
												} else {
													v101 = int32(1)
													v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
													if v102 == v96 {
														v161 = v101
													} else {
														if v96&int32(1) != 0 {
															v161 = int32(0)
														} else {
															v161 = v101
														}
													}
												}
												m.G0 = v8 + int32(48)
												return v161
											} else {
												if v92&int32(1) != 0 {
													v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
													if v108 == int32(0) {
														v161 = int32(0)
													} else {
														v161 = int32(1)
													}
													m.G0 = v8 + int32(48)
													return v161
												} else {
													v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
													v115 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
													v117 = F_FunctionCall2Coll(m, l0+int32(212), v114, v115, v116)
													mBase = m.M
													v118 = m.ExcPending
													if v118 != 0 {
														return int32(0)
													} else {
														if v117 == int32(0) {
															v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
															v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
															if v122 == int32(0) {
																v125 = int32(1)
																v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
																if v121&v125 == int32(0) {
																	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																	if v131 == v126 {
																		v161 = v125
																	} else {
																		if v126&int32(1) == int32(0) {
																			v161 = int32(0)
																		} else {
																			v161 = v125
																		}
																	}
																} else {
																	if v126&int32(1) == int32(0) {
																		v161 = int32(0)
																	} else {
																		v161 = v125
																	}
																}
															} else {
																v141 = int32(1)
																if v121&v141 != 0 {
																	v161 = v141
																} else {
																	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																	if v144&int32(1) != 0 {
																		v161 = int32(0)
																	} else {
																		v161 = v141
																	}
																}
															}
														} else {
															if v117 < int32(0) {
																v161 = int32(0)
															} else {
																v161 = int32(1)
															}
														}
														m.G0 = v8 + int32(48)
														return v161
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
		v170 = m.ExcPending
		if v170 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_range_contains_internal_0), int32(0))
			mBase = m.M
			v174 = m.ExcPending
			if v174 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_contains_internal_1), int32(2661), int32(_a_F_range_contains_internal_2))
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
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
func F_range_deduplicate_values(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v12 != v13 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v24 = l0 + v19<<(uint(int32(3))%32) + int32(36)
	F_qsort_arg(m, v24, v12, int32(4), int32(21), v10+int32(8))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(16)
	return
L4:
	;
	return
L5:
	;
	v31 = int32(1)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if int32(2) <= v32 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v38 = v31
	v39 = int32(1)
	goto L9
L7:
	;
	v78 = v31
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v78
	goto L3
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v47 = v24 + v39<<(uint(int32(2))%32)
	v49 = v47 - int32(4)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v52 = F_FunctionCall2Coll(m, v43, v44, v50, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	v78 = v71
	goto L8
L11:
	;
	v73 = v39 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v73 < v74 {
		v38 = v71
		v39 = v73
		goto L9
	} else {
		goto L18
	}
L12:
	;
	if v52 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v60 = F_FunctionCall2Coll(m, v56, v57, v58, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v24+v38<<(uint(int32(2))%32)))) = v67
	v71 = v38 + int32(1)
	goto L11
L16:
	;
	if v60 == int32(0) {
		v71 = v38
		goto L11
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L10
}
func F_range_gist_consistent_int_multirange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	switch l1 - int32(1) {
	case 0:
		v12 = int32(0)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v19 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v13)>>(uint(int32(2))%32))-int32(1)))))
		if v19&int32(1) != 0 {
			v146 = v12
			m.G0 = v8 + int32(16)
			return v146
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			if v22 == int32(0) {
				v146 = v12
				m.G0 = v8 + int32(16)
				return v146
			} else {
				v25 = F_range_overright_multirange_internal(m, l0, l2, l3)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v146 = v25 ^ int32(1)
					m.G0 = v8 + int32(16)
					return v146
				}
			}
		}
	case 1:
		v31 = int32(0)
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v32)>>(uint(int32(2))%32))-int32(1)))))
		if v38&int32(1) != 0 {
			v146 = v31
			m.G0 = v8 + int32(16)
			return v146
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			if v41 == int32(0) {
				v146 = v31
				m.G0 = v8 + int32(16)
				return v146
			} else {
				v44 = F_range_after_multirange_internal(m, l0, l2, l3)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v146 = v44 ^ int32(1)
					m.G0 = v8 + int32(16)
					return v146
				}
			}
		}
	case 2:
		v144 = F_range_overlaps_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v145 = m.ExcPending
		if v145 != 0 {
			return int32(0)
		} else {
			v146 = v144
			m.G0 = v8 + int32(16)
			return v146
		}
	case 3:
		v48 = int32(0)
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v49)>>(uint(int32(2))%32))-int32(1)))))
		if v55&int32(1) != 0 {
			v146 = v48
			m.G0 = v8 + int32(16)
			return v146
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			if v58 == int32(0) {
				v146 = v48
				m.G0 = v8 + int32(16)
				return v146
			} else {
				v61 = F_range_before_multirange_internal(m, l0, l2, l3)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v146 = v61 ^ int32(1)
					m.G0 = v8 + int32(16)
					return v146
				}
			}
		}
	case 4:
		v65 = int32(0)
		v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v72 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v66)>>(uint(int32(2))%32))-int32(1)))))
		if v72&int32(1) != 0 {
			v146 = v65
			m.G0 = v8 + int32(16)
			return v146
		} else {
			v75 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			if v75 == int32(0) {
				v146 = v65
				m.G0 = v8 + int32(16)
				return v146
			} else {
				v78 = F_range_overleft_multirange_internal(m, l0, l2, l3)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					v146 = v78 ^ int32(1)
					m.G0 = v8 + int32(16)
					return v146
				}
			}
		}
	case 5:
		v82 = int32(0)
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v83)>>(uint(int32(2))%32))-int32(1)))))
		if v89&int32(1) != 0 {
			v146 = v82
			m.G0 = v8 + int32(16)
			return v146
		} else {
			v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			if v92 == int32(0) {
				v146 = v82
				m.G0 = v8 + int32(16)
				return v146
			} else {
				v95 = F_range_adjacent_multirange_internal(m, l0, l2, l3)
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return int32(0)
				} else {
					if v95 == int32(0) {
						v144 = F_range_overlaps_multirange_internal(m, l0, l2, l3)
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return int32(0)
						} else {
							v146 = v144
							m.G0 = v8 + int32(16)
							return v146
						}
					} else {
						v146 = int32(1)
						m.G0 = v8 + int32(16)
						return v146
					}
				}
			}
		}
	case 6:
		v100 = F_range_contains_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v101 = m.ExcPending
		if v101 != 0 {
			return int32(0)
		} else {
			v146 = v100
			m.G0 = v8 + int32(16)
			return v146
		}
	case 7:
		v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v108 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v102)>>(uint(int32(2))%32))-int32(1)))))
		if v108&int32(-127) == int32(0) {
			v144 = F_range_overlaps_multirange_internal(m, l0, l2, l3)
			mBase = m.M
			v145 = m.ExcPending
			if v145 != 0 {
				return int32(0)
			} else {
				v146 = v144
				m.G0 = v8 + int32(16)
				return v146
			}
		} else {
			v146 = int32(1)
			m.G0 = v8 + int32(16)
			return v146
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v133 = m.ExcPending
		if v133 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
			F_errmsg_internal(m, int32(_a_F_range_gist_consistent_int_multirange_0), v8)
			mBase = m.M
			v137 = m.ExcPending
			if v137 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_gist_consistent_int_multirange_1), int32(1030), int32(_a_F_range_gist_consistent_int_multirange_2))
				mBase = m.M
				v142 = m.ExcPending
				if v142 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 17:
		v114 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
		if v114 == int32(0) {
			v117 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v123 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v117)>>(uint(int32(2))%32))-int32(1)))))
			v146 = base.B2i32(v123&int32(-127) != int32(0))
			m.G0 = v8 + int32(16)
			return v146
		} else {
			v128 = F_range_contains_multirange_internal(m, l0, l2, l3)
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return int32(0)
			} else {
				v146 = v128
				m.G0 = v8 + int32(16)
				return v146
			}
		}
	}
}
func F_range_gist_consistent_leaf_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	switch l1 - int32(1) {
	case 0:
		v43 = F_range_before_internal(m, l0, l2, l3)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v45 = v43
			m.G0 = v8 + int32(16)
			return v45
		}
	case 1:
		v12 = F_range_overleft_internal(m, l0, l2, l3)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v45 = v12
			m.G0 = v8 + int32(16)
			return v45
		}
	case 2:
		v16 = F_range_overlaps_internal(m, l0, l2, l3)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v45 = v16
			m.G0 = v8 + int32(16)
			return v45
		}
	case 3:
		v18 = F_range_overright_internal(m, l0, l2, l3)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v45 = v18
			m.G0 = v8 + int32(16)
			return v45
		}
	case 4:
		v20 = F_range_after_internal(m, l0, l2, l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v45 = v20
			m.G0 = v8 + int32(16)
			return v45
		}
	case 5:
		v22 = F_range_adjacent_internal(m, l0, l2, l3)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v45 = v22
			m.G0 = v8 + int32(16)
			return v45
		}
	case 6:
		v24 = F_range_contains_internal(m, l0, l2, l3)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v45 = v24
			m.G0 = v8 + int32(16)
			return v45
		}
	case 7:
		v26 = F_range_contained_by_internal(m, l0, l2, l3)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v45 = v26
			m.G0 = v8 + int32(16)
			return v45
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
			F_errmsg_internal(m, int32(_a_F_range_gist_consistent_leaf_range_0), v8)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_gist_consistent_leaf_range_1), int32(1084), int32(_a_F_range_gist_consistent_leaf_range_2))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 17:
		v28 = F_range_eq_internal(m, l0, l2, l3)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v45 = v28
			m.G0 = v8 + int32(16)
			return v45
		}
	}
}
func F_range_gist_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v13 = F_range_get_typcache(m, l0, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if int32(2) <= v15 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = int32(1)
	v23 = v8
	goto L7
L5:
	;
	v40 = v8
	goto L6
L6:
	;
	return v40
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v6+int32(4)+v21<<(uint(int32(4))%32))))
	v30 = F_pg_detoast_datum(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v40 = v32
	goto L6
L9:
	;
	v32 = F_range_super_union(m, v13, v23, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v35 = v21 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v35 < v36 {
		v21 = v35
		v23 = v32
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
}
func F_range_intersect(m *base.Module, l0 int32) int32 {
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
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
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			if v19 == v20 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
				if v23 != 0 {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					if v24 == v19 {
						v34 = v23
						v35 = F_range_intersect_internal(m, v34, v12, v17)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v35
						}
					} else {
						v27 = F_lookup_type_cache(m, v19, int32(2048))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+200))
							if v29 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
									F_errmsg_internal(m, int32(_a_F_range_intersect_0), v9)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_range_intersect_1), int32(1776), int32(_a_F_range_intersect_2))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v27
								v34 = v27
								v35 = F_range_intersect_internal(m, v34, v12, v17)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									m.G0 = v9 + int32(16)
									return v35
								}
							}
						}
					}
				} else {
					v27 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+200))
						if v29 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_intersect_0), v9)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_intersect_1), int32(1776), int32(_a_F_range_intersect_2))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v27
							v34 = v27
							v35 = F_range_intersect_internal(m, v34, v12, v17)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v35
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_range_intersect_3), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_range_intersect_1), int32(1137), int32(_a_F_range_intersect_4))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
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
	}
}
func F_range_intersect_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v187 int32
	_ = v187
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	F_range_deserialize(m, l0, l1, v5+int32(-24), v5+int32(-40), v5+int32(-49))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		F_range_deserialize(m, l0, l2, v5+int32(-32), v5+int32(-48), v5+int32(-50))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
			if v27 != 0 {
				v33 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+62)) = uint8(v33)
				v35 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v7)+60)) = uint16(v35)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v35
				*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v35
				*(*int32)(unsafe.Add(mBase, uint32(v7)+51)) = v35
				v49 = F_make_range(m, l0, v5+int32(-8), v5+int32(-16), v33, v35)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v210 = v49
					m.G0 = v7 - int32(-64)
					return v210
				}
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
				if v28&int32(1) != 0 {
					v33 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+62)) = uint8(v33)
					v35 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v7)+60)) = uint16(v35)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v7)+51)) = v35
					v49 = F_make_range(m, l0, v5+int32(-8), v5+int32(-16), v33, v35)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v210 = v49
						m.G0 = v7 - int32(-64)
						return v210
					}
				} else {
					v31 = F_range_overlaps_internal(m, l0, l1, l2)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 != 0 {
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+36)))
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+44)))
							if v52 == int32(1) {
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+46)))
								if v51&int32(1) != 0 {
									v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
									if base.B2i32(v55&int32(1) == int32(0))|base.B2i32(v62 == v55) != 0 {
										v126 = v5 + int32(-24)
									} else {
										v126 = v5 + int32(-32)
									}
								} else {
									if v55&int32(1) == int32(0) {
										v126 = v5 + int32(-24)
									} else {
										v126 = v5 + int32(-32)
									}
								}
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
								v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)))
								if v128 == int32(1) {
									v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
									if v127&int32(1) != 0 {
										v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
										if v131&int32(1)|base.B2i32(v131 == v136) != 0 {
											v202 = v5 + int32(-40)
										} else {
											v202 = v5 + int32(-48)
										}
									} else {
										if v131&int32(1) != 0 {
											v202 = v5 + int32(-40)
										} else {
											v202 = v5 + int32(-48)
										}
									}
									v203 = int32(0)
									v205 = F_make_range(m, l0, v126, v202, v203, v203)
									mBase = m.M
									v206 = m.ExcPending
									if v206 != 0 {
										return int32(0)
									} else {
										v210 = v205
										m.G0 = v7 - int32(-64)
										return v210
									}
								} else {
									if v127&int32(1) != 0 {
										v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
										if v147 == int32(0) {
											v202 = v5 + int32(-40)
										} else {
											v202 = v5 + int32(-48)
										}
										v203 = int32(0)
										v205 = F_make_range(m, l0, v126, v202, v203, v203)
										mBase = m.M
										v206 = m.ExcPending
										if v206 != 0 {
											return int32(0)
										} else {
											v210 = v205
											m.G0 = v7 - int32(-64)
											return v210
										}
									} else {
										v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
										v155 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
										v157 = F_FunctionCall2Coll(m, l0+int32(212), v154, v155, v156)
										mBase = m.M
										v158 = m.ExcPending
										if v158 != 0 {
											return int32(0)
										} else {
											if v157 == int32(0) {
												v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)))
												v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+29)))
												if v162 == int32(0) {
													v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
													if v161&int32(1) == int32(0) {
														v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
														if base.B2i32(v165&int32(1) == int32(0))|base.B2i32(v165 == v174) != 0 {
															v202 = v5 + int32(-40)
														} else {
															v202 = v5 + int32(-48)
														}
													} else {
														if v165&int32(1) == int32(0) {
															v202 = v5 + int32(-40)
														} else {
															v202 = v5 + int32(-48)
														}
													}
												} else {
													if v161&int32(1) != 0 {
														v202 = v5 + int32(-40)
													} else {
														v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
														if v187&int32(1) != 0 {
															v202 = v5 + int32(-40)
														} else {
															v202 = v5 + int32(-48)
														}
													}
												}
											} else {
												if v157 <= int32(0) {
													v202 = v5 + int32(-40)
												} else {
													v202 = v5 + int32(-48)
												}
											}
											v203 = int32(0)
											v205 = F_make_range(m, l0, v126, v202, v203, v203)
											mBase = m.M
											v206 = m.ExcPending
											if v206 != 0 {
												return int32(0)
											} else {
												v210 = v205
												m.G0 = v7 - int32(-64)
												return v210
											}
										}
									}
								}
							} else {
								if v51&int32(1) != 0 {
									v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
									if v75 != 0 {
										v126 = v5 + int32(-24)
									} else {
										v126 = v5 + int32(-32)
									}
									v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
									v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)))
									if v128 == int32(1) {
										v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
										if v127&int32(1) != 0 {
											v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
											if v131&int32(1)|base.B2i32(v131 == v136) != 0 {
												v202 = v5 + int32(-40)
											} else {
												v202 = v5 + int32(-48)
											}
										} else {
											if v131&int32(1) != 0 {
												v202 = v5 + int32(-40)
											} else {
												v202 = v5 + int32(-48)
											}
										}
										v203 = int32(0)
										v205 = F_make_range(m, l0, v126, v202, v203, v203)
										mBase = m.M
										v206 = m.ExcPending
										if v206 != 0 {
											return int32(0)
										} else {
											v210 = v205
											m.G0 = v7 - int32(-64)
											return v210
										}
									} else {
										if v127&int32(1) != 0 {
											v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
											if v147 == int32(0) {
												v202 = v5 + int32(-40)
											} else {
												v202 = v5 + int32(-48)
											}
											v203 = int32(0)
											v205 = F_make_range(m, l0, v126, v202, v203, v203)
											mBase = m.M
											v206 = m.ExcPending
											if v206 != 0 {
												return int32(0)
											} else {
												v210 = v205
												m.G0 = v7 - int32(-64)
												return v210
											}
										} else {
											v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v155 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
											v156 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
											v157 = F_FunctionCall2Coll(m, l0+int32(212), v154, v155, v156)
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return int32(0)
											} else {
												if v157 == int32(0) {
													v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)))
													v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+29)))
													if v162 == int32(0) {
														v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
														if v161&int32(1) == int32(0) {
															v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
															if base.B2i32(v165&int32(1) == int32(0))|base.B2i32(v165 == v174) != 0 {
																v202 = v5 + int32(-40)
															} else {
																v202 = v5 + int32(-48)
															}
														} else {
															if v165&int32(1) == int32(0) {
																v202 = v5 + int32(-40)
															} else {
																v202 = v5 + int32(-48)
															}
														}
													} else {
														if v161&int32(1) != 0 {
															v202 = v5 + int32(-40)
														} else {
															v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
															if v187&int32(1) != 0 {
																v202 = v5 + int32(-40)
															} else {
																v202 = v5 + int32(-48)
															}
														}
													}
												} else {
													if v157 <= int32(0) {
														v202 = v5 + int32(-40)
													} else {
														v202 = v5 + int32(-48)
													}
												}
												v203 = int32(0)
												v205 = F_make_range(m, l0, v126, v202, v203, v203)
												mBase = m.M
												v206 = m.ExcPending
												if v206 != 0 {
													return int32(0)
												} else {
													v210 = v205
													m.G0 = v7 - int32(-64)
													return v210
												}
											}
										}
									}
								} else {
									v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
									v83 = F_FunctionCall2Coll(m, l0+int32(212), v80, v81, v82)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										if v83 == int32(0) {
											v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+37)))
											v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+45)))
											if v88 == int32(0) {
												v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+46)))
												if v87&int32(1) == int32(0) {
													v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
													if v91&int32(1)|base.B2i32(v98 == v91) != 0 {
														v126 = v5 + int32(-24)
													} else {
														v126 = v5 + int32(-32)
													}
												} else {
													if v91&int32(1) != 0 {
														v126 = v5 + int32(-24)
													} else {
														v126 = v5 + int32(-32)
													}
												}
											} else {
												if v87&int32(1) != 0 {
													v126 = v5 + int32(-24)
												} else {
													v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
													if v109&int32(1) == int32(0) {
														v126 = v5 + int32(-24)
													} else {
														v126 = v5 + int32(-32)
													}
												}
											}
										} else {
											if int32(0) <= v83 {
												v126 = v5 + int32(-24)
											} else {
												v126 = v5 + int32(-32)
											}
										}
										v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
										v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)))
										if v128 == int32(1) {
											v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
											if v127&int32(1) != 0 {
												v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
												if v131&int32(1)|base.B2i32(v131 == v136) != 0 {
													v202 = v5 + int32(-40)
												} else {
													v202 = v5 + int32(-48)
												}
											} else {
												if v131&int32(1) != 0 {
													v202 = v5 + int32(-40)
												} else {
													v202 = v5 + int32(-48)
												}
											}
											v203 = int32(0)
											v205 = F_make_range(m, l0, v126, v202, v203, v203)
											mBase = m.M
											v206 = m.ExcPending
											if v206 != 0 {
												return int32(0)
											} else {
												v210 = v205
												m.G0 = v7 - int32(-64)
												return v210
											}
										} else {
											if v127&int32(1) != 0 {
												v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
												if v147 == int32(0) {
													v202 = v5 + int32(-40)
												} else {
													v202 = v5 + int32(-48)
												}
												v203 = int32(0)
												v205 = F_make_range(m, l0, v126, v202, v203, v203)
												mBase = m.M
												v206 = m.ExcPending
												if v206 != 0 {
													return int32(0)
												} else {
													v210 = v205
													m.G0 = v7 - int32(-64)
													return v210
												}
											} else {
												v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
												v155 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
												v156 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
												v157 = F_FunctionCall2Coll(m, l0+int32(212), v154, v155, v156)
												mBase = m.M
												v158 = m.ExcPending
												if v158 != 0 {
													return int32(0)
												} else {
													if v157 == int32(0) {
														v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)))
														v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+29)))
														if v162 == int32(0) {
															v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
															if v161&int32(1) == int32(0) {
																v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
																if base.B2i32(v165&int32(1) == int32(0))|base.B2i32(v165 == v174) != 0 {
																	v202 = v5 + int32(-40)
																} else {
																	v202 = v5 + int32(-48)
																}
															} else {
																if v165&int32(1) == int32(0) {
																	v202 = v5 + int32(-40)
																} else {
																	v202 = v5 + int32(-48)
																}
															}
														} else {
															if v161&int32(1) != 0 {
																v202 = v5 + int32(-40)
															} else {
																v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
																if v187&int32(1) != 0 {
																	v202 = v5 + int32(-40)
																} else {
																	v202 = v5 + int32(-48)
																}
															}
														}
													} else {
														if v157 <= int32(0) {
															v202 = v5 + int32(-40)
														} else {
															v202 = v5 + int32(-48)
														}
													}
													v203 = int32(0)
													v205 = F_make_range(m, l0, v126, v202, v203, v203)
													mBase = m.M
													v206 = m.ExcPending
													if v206 != 0 {
														return int32(0)
													} else {
														v210 = v205
														m.G0 = v7 - int32(-64)
														return v210
													}
												}
											}
										}
									}
								}
							}
						} else {
							v33 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+62)) = uint8(v33)
							v35 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v7)+60)) = uint16(v35)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v7)+51)) = v35
							v49 = F_make_range(m, l0, v5+int32(-8), v5+int32(-16), v33, v35)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v210 = v49
								m.G0 = v7 - int32(-64)
								return v210
							}
						}
					}
				}
			}
		}
	}
}
func F_range_le(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_range_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v2 <= int32(0))
	}
}
func F_range_overlaps_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v13)>>(uint(int32(2))%32))-int32(1)))))
	goto L2
L1:
	;
	m.G0 = v11 + int32(48)
	return v78
L2:
	;
	if v19&int32(1) != 0 {
		v78 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v22 == int32(0) {
		v78 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = v11 + int32(16)
	v28 = v26 | int32(8)
	F_range_deserialize(m, l0, l1, v26, v28, v11+int32(15))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v35 == int32(0) {
		v78 = v4
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v39 = v35
	v42 = v4
	goto L9
L8:
	;
	v78 = int32(1)
	goto L1
L9:
	;
	v48 = int32(base.Ui32(v39+v42) >> (uint(int32(1)) % 32))
	v50 = v11 + int32(40)
	F_multirange_get_bounds(m, l0, l2, v48, v50, v11+int32(32))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	v78 = int32(0)
	goto L1
L11:
	;
	v55 = F_range_cmp_bounds(m, l0, v28, v50)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	if base.Ui32(v70) < base.Ui32(v69) {
		v39 = v69
		v42 = v70
		goto L9
	} else {
		goto L19
	}
L13:
	;
	if v55 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v69 = v48
	v70 = v42
	goto L12
L15:
	;
	goto L16
L16:
	;
	v63 = F_range_cmp_bounds(m, l0, v11+int32(16), v11+int32(32))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v63 <= int32(0) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v69 = v39
	v70 = v48 + int32(1)
	goto L12
L19:
	;
	goto L10
}
func F_range_overleft_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v9 == v10 {
		F_range_deserialize(m, l0, l1, v7+int32(40), v7+int32(24), v7+int32(15))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_range_deserialize(m, l0, l2, v7+int32(32), v7+int32(16), v7+int32(14))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(0)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
				if v31 != 0 {
					v103 = v30
					m.G0 = v7 + int32(48)
					return v103
				} else {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
					if v32&int32(1) != 0 {
						v103 = v30
						m.G0 = v7 + int32(48)
						return v103
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)))
						if v36 == int32(1) {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
							if v35&int32(1) != 0 {
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
								if v42 == v39 {
									v97 = int32(0)
								} else {
									v46 = int32(1)
									if v39&v46 != 0 {
										v49 = int32(-1)
									} else {
										v49 = v46
									}
									v97 = v49
								}
							} else {
								v51 = int32(1)
								if v39&v51 != 0 {
									v54 = int32(-1)
								} else {
									v54 = v51
								}
								v97 = v54
							}
							v103 = base.B2i32(v97 <= int32(0))
							m.G0 = v7 + int32(48)
							return v103
						} else {
							if v35&int32(1) != 0 {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
								if v59 != 0 {
									v60 = int32(1)
								} else {
									v60 = int32(-1)
								}
								v97 = v60
								v103 = base.B2i32(v97 <= int32(0))
								m.G0 = v7 + int32(48)
								return v103
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
								v66 = F_FunctionCall2Coll(m, l0+int32(212), v63, v64, v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									if v66 != 0 {
										v97 = v66
									} else {
										v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)))
										v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+29)))
										if v69 == int32(0) {
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
											if v68&int32(1) == int32(0) {
												v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
												if v77 == v72 {
													v97 = int32(0)
												} else {
													v80 = int32(1)
													if v72&v80 != 0 {
														v84 = v80
													} else {
														v84 = int32(-1)
													}
													v97 = v84
												}
											} else {
												v85 = int32(1)
												if v72&v85 != 0 {
													v89 = v85
												} else {
													v89 = int32(-1)
												}
												v97 = v89
											}
										} else {
											if v68&int32(1) != 0 {
												v97 = int32(0)
											} else {
												v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
												if v95 != 0 {
													v96 = int32(-1)
												} else {
													v96 = int32(1)
												}
												v97 = v96
											}
										}
									}
									v103 = base.B2i32(v97 <= int32(0))
									m.G0 = v7 + int32(48)
									return v103
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
		v111 = m.ExcPending
		if v111 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_range_overleft_internal_0), int32(0))
			mBase = m.M
			v115 = m.ExcPending
			if v115 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_overleft_internal_1), int32(900), int32(_a_F_range_overleft_internal_2))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
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
func F_range_union(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13987(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_range_union_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
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
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v11 == v12 {
		F_range_deserialize(m, l0, l1, v9+int32(40), v9+int32(24), v9+int32(15))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			F_range_deserialize(m, l0, l2, v9+int32(32), v9+int32(16), v9+int32(14))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				if v32 != 0 {
					v228 = l2
					m.G0 = v9 + int32(48)
					return v228
				} else {
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)))
					if v33 != 0 {
						v228 = l1
						m.G0 = v9 + int32(48)
						return v228
					} else {
						if l3 == int32(0) {
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)))
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)))
							if v43 == int32(1) {
								v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+46)))
								if v42&int32(1) != 0 {
									v50 = v9 + int32(32)
									v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
									if v51 == v46 {
										v143 = v50
									} else {
										if v46&int32(1) != 0 {
											v143 = v9 + int32(40)
										} else {
											v143 = v50
										}
									}
								} else {
									if v46&int32(1) != 0 {
										v143 = v9 + int32(40)
									} else {
										v143 = v9 + int32(32)
									}
								}
								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
								v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
								if v146 == int32(1) {
									v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
									if v145&int32(1) != 0 {
										v153 = v9 + int32(16)
										v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
										if v154 == v149 {
											v220 = v153
										} else {
											if v149&int32(1) == int32(0) {
												v220 = v9 + int32(24)
											} else {
												v220 = v153
											}
										}
									} else {
										if v149&int32(1) == int32(0) {
											v220 = v9 + int32(24)
										} else {
											v220 = v9 + int32(16)
										}
									}
									v223 = int32(0)
									v225 = F_make_range(m, l0, v143, v220, v223, v223)
									mBase = m.M
									v226 = m.ExcPending
									if v226 != 0 {
										return int32(0)
									} else {
										v228 = v225
										m.G0 = v9 + int32(48)
										return v228
									}
								} else {
									if v145&int32(1) != 0 {
										v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
										if v168 != 0 {
											v220 = v9 + int32(24)
										} else {
											v220 = v9 + int32(16)
										}
										v223 = int32(0)
										v225 = F_make_range(m, l0, v143, v220, v223, v223)
										mBase = m.M
										v226 = m.ExcPending
										if v226 != 0 {
											return int32(0)
										} else {
											v228 = v225
											m.G0 = v9 + int32(48)
											return v228
										}
									} else {
										v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
										v174 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
										v175 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
										v176 = F_FunctionCall2Coll(m, l0+int32(212), v173, v174, v175)
										mBase = m.M
										v177 = m.ExcPending
										if v177 != 0 {
											return int32(0)
										} else {
											if v176 == int32(0) {
												v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
												v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
												if v181 == int32(0) {
													v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
													if v180&int32(1) == int32(0) {
														v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
														if base.B2i32(v184&int32(1) == int32(0))|base.B2i32(v184 == v195) != 0 {
															v220 = v9 + int32(16)
														} else {
															v220 = v9 + int32(24)
														}
													} else {
														if v184&int32(1) != 0 {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
													}
												} else {
													v203 = v9 + int32(16)
													if v180&int32(1) != 0 {
														v220 = v203
													} else {
														v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
														if v206&int32(1) == int32(0) {
															v220 = v9 + int32(24)
														} else {
															v220 = v203
														}
													}
												}
											} else {
												if int32(0) < v176 {
													v220 = v9 + int32(24)
												} else {
													v220 = v9 + int32(16)
												}
											}
											v223 = int32(0)
											v225 = F_make_range(m, l0, v143, v220, v223, v223)
											mBase = m.M
											v226 = m.ExcPending
											if v226 != 0 {
												return int32(0)
											} else {
												v228 = v225
												m.G0 = v9 + int32(48)
												return v228
											}
										}
									}
								}
							} else {
								if v42&int32(1) != 0 {
									v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
									if v61 == int32(0) {
										v143 = v9 + int32(40)
									} else {
										v143 = v9 + int32(32)
									}
									v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
									v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
									if v146 == int32(1) {
										v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
										if v145&int32(1) != 0 {
											v153 = v9 + int32(16)
											v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
											if v154 == v149 {
												v220 = v153
											} else {
												if v149&int32(1) == int32(0) {
													v220 = v9 + int32(24)
												} else {
													v220 = v153
												}
											}
										} else {
											if v149&int32(1) == int32(0) {
												v220 = v9 + int32(24)
											} else {
												v220 = v9 + int32(16)
											}
										}
										v223 = int32(0)
										v225 = F_make_range(m, l0, v143, v220, v223, v223)
										mBase = m.M
										v226 = m.ExcPending
										if v226 != 0 {
											return int32(0)
										} else {
											v228 = v225
											m.G0 = v9 + int32(48)
											return v228
										}
									} else {
										if v145&int32(1) != 0 {
											v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
											if v168 != 0 {
												v220 = v9 + int32(24)
											} else {
												v220 = v9 + int32(16)
											}
											v223 = int32(0)
											v225 = F_make_range(m, l0, v143, v220, v223, v223)
											mBase = m.M
											v226 = m.ExcPending
											if v226 != 0 {
												return int32(0)
											} else {
												v228 = v225
												m.G0 = v9 + int32(48)
												return v228
											}
										} else {
											v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v174 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
											v175 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
											v176 = F_FunctionCall2Coll(m, l0+int32(212), v173, v174, v175)
											mBase = m.M
											v177 = m.ExcPending
											if v177 != 0 {
												return int32(0)
											} else {
												if v176 == int32(0) {
													v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
													v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
													if v181 == int32(0) {
														v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
														if v180&int32(1) == int32(0) {
															v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
															if base.B2i32(v184&int32(1) == int32(0))|base.B2i32(v184 == v195) != 0 {
																v220 = v9 + int32(16)
															} else {
																v220 = v9 + int32(24)
															}
														} else {
															if v184&int32(1) != 0 {
																v220 = v9 + int32(24)
															} else {
																v220 = v9 + int32(16)
															}
														}
													} else {
														v203 = v9 + int32(16)
														if v180&int32(1) != 0 {
															v220 = v203
														} else {
															v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
															if v206&int32(1) == int32(0) {
																v220 = v9 + int32(24)
															} else {
																v220 = v203
															}
														}
													}
												} else {
													if int32(0) < v176 {
														v220 = v9 + int32(24)
													} else {
														v220 = v9 + int32(16)
													}
												}
												v223 = int32(0)
												v225 = F_make_range(m, l0, v143, v220, v223, v223)
												mBase = m.M
												v226 = m.ExcPending
												if v226 != 0 {
													return int32(0)
												} else {
													v228 = v225
													m.G0 = v9 + int32(48)
													return v228
												}
											}
										}
									}
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
									v71 = F_FunctionCall2Coll(m, l0+int32(212), v68, v69, v70)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										if v71 == int32(0) {
											v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+37)))
											v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+45)))
											if v76 == int32(0) {
												v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+46)))
												if v75&int32(1) == int32(0) {
													v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
													if v79&int32(1)|base.B2i32(v88 == v79) != 0 {
														v143 = v9 + int32(32)
													} else {
														v143 = v9 + int32(40)
													}
												} else {
													if v79&int32(1) == int32(0) {
														v143 = v9 + int32(40)
													} else {
														v143 = v9 + int32(32)
													}
												}
											} else {
												v98 = v9 + int32(32)
												if v75&int32(1) != 0 {
													v143 = v98
												} else {
													v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
													if v101&int32(1) != 0 {
														v143 = v9 + int32(40)
													} else {
														v143 = v98
													}
												}
											}
										} else {
											if v71 < int32(0) {
												v143 = v9 + int32(40)
											} else {
												v143 = v9 + int32(32)
											}
										}
										v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
										v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
										if v146 == int32(1) {
											v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
											if v145&int32(1) != 0 {
												v153 = v9 + int32(16)
												v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
												if v154 == v149 {
													v220 = v153
												} else {
													if v149&int32(1) == int32(0) {
														v220 = v9 + int32(24)
													} else {
														v220 = v153
													}
												}
											} else {
												if v149&int32(1) == int32(0) {
													v220 = v9 + int32(24)
												} else {
													v220 = v9 + int32(16)
												}
											}
											v223 = int32(0)
											v225 = F_make_range(m, l0, v143, v220, v223, v223)
											mBase = m.M
											v226 = m.ExcPending
											if v226 != 0 {
												return int32(0)
											} else {
												v228 = v225
												m.G0 = v9 + int32(48)
												return v228
											}
										} else {
											if v145&int32(1) != 0 {
												v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
												if v168 != 0 {
													v220 = v9 + int32(24)
												} else {
													v220 = v9 + int32(16)
												}
												v223 = int32(0)
												v225 = F_make_range(m, l0, v143, v220, v223, v223)
												mBase = m.M
												v226 = m.ExcPending
												if v226 != 0 {
													return int32(0)
												} else {
													v228 = v225
													m.G0 = v9 + int32(48)
													return v228
												}
											} else {
												v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
												v174 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
												v175 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
												v176 = F_FunctionCall2Coll(m, l0+int32(212), v173, v174, v175)
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
													return int32(0)
												} else {
													if v176 == int32(0) {
														v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
														v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
														if v181 == int32(0) {
															v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
															if v180&int32(1) == int32(0) {
																v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																if base.B2i32(v184&int32(1) == int32(0))|base.B2i32(v184 == v195) != 0 {
																	v220 = v9 + int32(16)
																} else {
																	v220 = v9 + int32(24)
																}
															} else {
																if v184&int32(1) != 0 {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v9 + int32(16)
																}
															}
														} else {
															v203 = v9 + int32(16)
															if v180&int32(1) != 0 {
																v220 = v203
															} else {
																v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																if v206&int32(1) == int32(0) {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v203
																}
															}
														}
													} else {
														if int32(0) < v176 {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
													}
													v223 = int32(0)
													v225 = F_make_range(m, l0, v143, v220, v223, v223)
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return int32(0)
													} else {
														v228 = v225
														m.G0 = v9 + int32(48)
														return v228
													}
												}
											}
										}
									}
								}
							}
						} else {
							v36 = F_range_overlaps_internal(m, l0, l1, l2)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								if v36 != 0 {
									v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)))
									v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)))
									if v43 == int32(1) {
										v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+46)))
										if v42&int32(1) != 0 {
											v50 = v9 + int32(32)
											v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
											if v51 == v46 {
												v143 = v50
											} else {
												if v46&int32(1) != 0 {
													v143 = v9 + int32(40)
												} else {
													v143 = v50
												}
											}
										} else {
											if v46&int32(1) != 0 {
												v143 = v9 + int32(40)
											} else {
												v143 = v9 + int32(32)
											}
										}
										v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
										v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
										if v146 == int32(1) {
											v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
											if v145&int32(1) != 0 {
												v153 = v9 + int32(16)
												v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
												if v154 == v149 {
													v220 = v153
												} else {
													if v149&int32(1) == int32(0) {
														v220 = v9 + int32(24)
													} else {
														v220 = v153
													}
												}
											} else {
												if v149&int32(1) == int32(0) {
													v220 = v9 + int32(24)
												} else {
													v220 = v9 + int32(16)
												}
											}
											v223 = int32(0)
											v225 = F_make_range(m, l0, v143, v220, v223, v223)
											mBase = m.M
											v226 = m.ExcPending
											if v226 != 0 {
												return int32(0)
											} else {
												v228 = v225
												m.G0 = v9 + int32(48)
												return v228
											}
										} else {
											if v145&int32(1) != 0 {
												v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
												if v168 != 0 {
													v220 = v9 + int32(24)
												} else {
													v220 = v9 + int32(16)
												}
												v223 = int32(0)
												v225 = F_make_range(m, l0, v143, v220, v223, v223)
												mBase = m.M
												v226 = m.ExcPending
												if v226 != 0 {
													return int32(0)
												} else {
													v228 = v225
													m.G0 = v9 + int32(48)
													return v228
												}
											} else {
												v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
												v174 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
												v175 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
												v176 = F_FunctionCall2Coll(m, l0+int32(212), v173, v174, v175)
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
													return int32(0)
												} else {
													if v176 == int32(0) {
														v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
														v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
														if v181 == int32(0) {
															v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
															if v180&int32(1) == int32(0) {
																v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																if base.B2i32(v184&int32(1) == int32(0))|base.B2i32(v184 == v195) != 0 {
																	v220 = v9 + int32(16)
																} else {
																	v220 = v9 + int32(24)
																}
															} else {
																if v184&int32(1) != 0 {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v9 + int32(16)
																}
															}
														} else {
															v203 = v9 + int32(16)
															if v180&int32(1) != 0 {
																v220 = v203
															} else {
																v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																if v206&int32(1) == int32(0) {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v203
																}
															}
														}
													} else {
														if int32(0) < v176 {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
													}
													v223 = int32(0)
													v225 = F_make_range(m, l0, v143, v220, v223, v223)
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return int32(0)
													} else {
														v228 = v225
														m.G0 = v9 + int32(48)
														return v228
													}
												}
											}
										}
									} else {
										if v42&int32(1) != 0 {
											v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
											if v61 == int32(0) {
												v143 = v9 + int32(40)
											} else {
												v143 = v9 + int32(32)
											}
											v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
											v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
											if v146 == int32(1) {
												v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
												if v145&int32(1) != 0 {
													v153 = v9 + int32(16)
													v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
													if v154 == v149 {
														v220 = v153
													} else {
														if v149&int32(1) == int32(0) {
															v220 = v9 + int32(24)
														} else {
															v220 = v153
														}
													}
												} else {
													if v149&int32(1) == int32(0) {
														v220 = v9 + int32(24)
													} else {
														v220 = v9 + int32(16)
													}
												}
												v223 = int32(0)
												v225 = F_make_range(m, l0, v143, v220, v223, v223)
												mBase = m.M
												v226 = m.ExcPending
												if v226 != 0 {
													return int32(0)
												} else {
													v228 = v225
													m.G0 = v9 + int32(48)
													return v228
												}
											} else {
												if v145&int32(1) != 0 {
													v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
													if v168 != 0 {
														v220 = v9 + int32(24)
													} else {
														v220 = v9 + int32(16)
													}
													v223 = int32(0)
													v225 = F_make_range(m, l0, v143, v220, v223, v223)
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return int32(0)
													} else {
														v228 = v225
														m.G0 = v9 + int32(48)
														return v228
													}
												} else {
													v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
													v174 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
													v175 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
													v176 = F_FunctionCall2Coll(m, l0+int32(212), v173, v174, v175)
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
														return int32(0)
													} else {
														if v176 == int32(0) {
															v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
															v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
															if v181 == int32(0) {
																v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
																if v180&int32(1) == int32(0) {
																	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																	if base.B2i32(v184&int32(1) == int32(0))|base.B2i32(v184 == v195) != 0 {
																		v220 = v9 + int32(16)
																	} else {
																		v220 = v9 + int32(24)
																	}
																} else {
																	if v184&int32(1) != 0 {
																		v220 = v9 + int32(24)
																	} else {
																		v220 = v9 + int32(16)
																	}
																}
															} else {
																v203 = v9 + int32(16)
																if v180&int32(1) != 0 {
																	v220 = v203
																} else {
																	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																	if v206&int32(1) == int32(0) {
																		v220 = v9 + int32(24)
																	} else {
																		v220 = v203
																	}
																}
															}
														} else {
															if int32(0) < v176 {
																v220 = v9 + int32(24)
															} else {
																v220 = v9 + int32(16)
															}
														}
														v223 = int32(0)
														v225 = F_make_range(m, l0, v143, v220, v223, v223)
														mBase = m.M
														v226 = m.ExcPending
														if v226 != 0 {
															return int32(0)
														} else {
															v228 = v225
															m.G0 = v9 + int32(48)
															return v228
														}
													}
												}
											}
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
											v71 = F_FunctionCall2Coll(m, l0+int32(212), v68, v69, v70)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int32(0)
											} else {
												if v71 == int32(0) {
													v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+37)))
													v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+45)))
													if v76 == int32(0) {
														v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+46)))
														if v75&int32(1) == int32(0) {
															v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
															if v79&int32(1)|base.B2i32(v88 == v79) != 0 {
																v143 = v9 + int32(32)
															} else {
																v143 = v9 + int32(40)
															}
														} else {
															if v79&int32(1) == int32(0) {
																v143 = v9 + int32(40)
															} else {
																v143 = v9 + int32(32)
															}
														}
													} else {
														v98 = v9 + int32(32)
														if v75&int32(1) != 0 {
															v143 = v98
														} else {
															v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
															if v101&int32(1) != 0 {
																v143 = v9 + int32(40)
															} else {
																v143 = v98
															}
														}
													}
												} else {
													if v71 < int32(0) {
														v143 = v9 + int32(40)
													} else {
														v143 = v9 + int32(32)
													}
												}
												v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
												v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
												if v146 == int32(1) {
													v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
													if v145&int32(1) != 0 {
														v153 = v9 + int32(16)
														v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
														if v154 == v149 {
															v220 = v153
														} else {
															if v149&int32(1) == int32(0) {
																v220 = v9 + int32(24)
															} else {
																v220 = v153
															}
														}
													} else {
														if v149&int32(1) == int32(0) {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
													}
													v223 = int32(0)
													v225 = F_make_range(m, l0, v143, v220, v223, v223)
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return int32(0)
													} else {
														v228 = v225
														m.G0 = v9 + int32(48)
														return v228
													}
												} else {
													if v145&int32(1) != 0 {
														v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
														if v168 != 0 {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
														v223 = int32(0)
														v225 = F_make_range(m, l0, v143, v220, v223, v223)
														mBase = m.M
														v226 = m.ExcPending
														if v226 != 0 {
															return int32(0)
														} else {
															v228 = v225
															m.G0 = v9 + int32(48)
															return v228
														}
													} else {
														v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
														v174 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
														v175 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
														v176 = F_FunctionCall2Coll(m, l0+int32(212), v173, v174, v175)
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return int32(0)
														} else {
															if v176 == int32(0) {
																v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
																v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
																if v181 == int32(0) {
																	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
																	if v180&int32(1) == int32(0) {
																		v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																		if base.B2i32(v184&int32(1) == int32(0))|base.B2i32(v184 == v195) != 0 {
																			v220 = v9 + int32(16)
																		} else {
																			v220 = v9 + int32(24)
																		}
																	} else {
																		if v184&int32(1) != 0 {
																			v220 = v9 + int32(24)
																		} else {
																			v220 = v9 + int32(16)
																		}
																	}
																} else {
																	v203 = v9 + int32(16)
																	if v180&int32(1) != 0 {
																		v220 = v203
																	} else {
																		v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																		if v206&int32(1) == int32(0) {
																			v220 = v9 + int32(24)
																		} else {
																			v220 = v203
																		}
																	}
																}
															} else {
																if int32(0) < v176 {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v9 + int32(16)
																}
															}
															v223 = int32(0)
															v225 = F_make_range(m, l0, v143, v220, v223, v223)
															mBase = m.M
															v226 = m.ExcPending
															if v226 != 0 {
																return int32(0)
															} else {
																v228 = v225
																m.G0 = v9 + int32(48)
																return v228
															}
														}
													}
												}
											}
										}
									}
								} else {
									v38 = F_range_adjacent_internal(m, l0, l1, l2)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										if v38 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(130))
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_range_union_internal_0), int32(0))
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_range_union_internal_1), int32(1084), int32(_a_F_range_union_internal_2))
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+36)))
											v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+44)))
											if v43 == int32(1) {
												v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+46)))
												if v42&int32(1) != 0 {
													v50 = v9 + int32(32)
													v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
													if v51 == v46 {
														v143 = v50
													} else {
														if v46&int32(1) != 0 {
															v143 = v9 + int32(40)
														} else {
															v143 = v50
														}
													}
												} else {
													if v46&int32(1) != 0 {
														v143 = v9 + int32(40)
													} else {
														v143 = v9 + int32(32)
													}
												}
												v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
												v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
												if v146 == int32(1) {
													v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
													if v145&int32(1) != 0 {
														v153 = v9 + int32(16)
														v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
														if v154 == v149 {
															v220 = v153
														} else {
															if v149&int32(1) == int32(0) {
																v220 = v9 + int32(24)
															} else {
																v220 = v153
															}
														}
													} else {
														if v149&int32(1) == int32(0) {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
													}
													v223 = int32(0)
													v225 = F_make_range(m, l0, v143, v220, v223, v223)
													mBase = m.M
													v226 = m.ExcPending
													if v226 != 0 {
														return int32(0)
													} else {
														v228 = v225
														m.G0 = v9 + int32(48)
														return v228
													}
												} else {
													if v145&int32(1) != 0 {
														v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
														if v168 != 0 {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
														v223 = int32(0)
														v225 = F_make_range(m, l0, v143, v220, v223, v223)
														mBase = m.M
														v226 = m.ExcPending
														if v226 != 0 {
															return int32(0)
														} else {
															v228 = v225
															m.G0 = v9 + int32(48)
															return v228
														}
													} else {
														v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
														v174 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
														v175 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
														v176 = F_FunctionCall2Coll(m, l0+int32(212), v173, v174, v175)
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return int32(0)
														} else {
															if v176 == int32(0) {
																v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
																v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
																if v181 == int32(0) {
																	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
																	if v180&int32(1) == int32(0) {
																		v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																		if base.B2i32(v184&int32(1) == int32(0))|base.B2i32(v184 == v195) != 0 {
																			v220 = v9 + int32(16)
																		} else {
																			v220 = v9 + int32(24)
																		}
																	} else {
																		if v184&int32(1) != 0 {
																			v220 = v9 + int32(24)
																		} else {
																			v220 = v9 + int32(16)
																		}
																	}
																} else {
																	v203 = v9 + int32(16)
																	if v180&int32(1) != 0 {
																		v220 = v203
																	} else {
																		v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																		if v206&int32(1) == int32(0) {
																			v220 = v9 + int32(24)
																		} else {
																			v220 = v203
																		}
																	}
																}
															} else {
																if int32(0) < v176 {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v9 + int32(16)
																}
															}
															v223 = int32(0)
															v225 = F_make_range(m, l0, v143, v220, v223, v223)
															mBase = m.M
															v226 = m.ExcPending
															if v226 != 0 {
																return int32(0)
															} else {
																v228 = v225
																m.G0 = v9 + int32(48)
																return v228
															}
														}
													}
												}
											} else {
												if v42&int32(1) != 0 {
													v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
													if v61 == int32(0) {
														v143 = v9 + int32(40)
													} else {
														v143 = v9 + int32(32)
													}
													v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
													v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
													if v146 == int32(1) {
														v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
														if v145&int32(1) != 0 {
															v153 = v9 + int32(16)
															v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
															if v154 == v149 {
																v220 = v153
															} else {
																if v149&int32(1) == int32(0) {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v153
																}
															}
														} else {
															if v149&int32(1) == int32(0) {
																v220 = v9 + int32(24)
															} else {
																v220 = v9 + int32(16)
															}
														}
														v223 = int32(0)
														v225 = F_make_range(m, l0, v143, v220, v223, v223)
														mBase = m.M
														v226 = m.ExcPending
														if v226 != 0 {
															return int32(0)
														} else {
															v228 = v225
															m.G0 = v9 + int32(48)
															return v228
														}
													} else {
														if v145&int32(1) != 0 {
															v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
															if v168 != 0 {
																v220 = v9 + int32(24)
															} else {
																v220 = v9 + int32(16)
															}
															v223 = int32(0)
															v225 = F_make_range(m, l0, v143, v220, v223, v223)
															mBase = m.M
															v226 = m.ExcPending
															if v226 != 0 {
																return int32(0)
															} else {
																v228 = v225
																m.G0 = v9 + int32(48)
																return v228
															}
														} else {
															v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
															v174 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
															v175 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
															v176 = F_FunctionCall2Coll(m, l0+int32(212), v173, v174, v175)
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int32(0)
															} else {
																if v176 == int32(0) {
																	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
																	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
																	if v181 == int32(0) {
																		v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
																		if v180&int32(1) == int32(0) {
																			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																			if base.B2i32(v184&int32(1) == int32(0))|base.B2i32(v184 == v195) != 0 {
																				v220 = v9 + int32(16)
																			} else {
																				v220 = v9 + int32(24)
																			}
																		} else {
																			if v184&int32(1) != 0 {
																				v220 = v9 + int32(24)
																			} else {
																				v220 = v9 + int32(16)
																			}
																		}
																	} else {
																		v203 = v9 + int32(16)
																		if v180&int32(1) != 0 {
																			v220 = v203
																		} else {
																			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																			if v206&int32(1) == int32(0) {
																				v220 = v9 + int32(24)
																			} else {
																				v220 = v203
																			}
																		}
																	}
																} else {
																	if int32(0) < v176 {
																		v220 = v9 + int32(24)
																	} else {
																		v220 = v9 + int32(16)
																	}
																}
																v223 = int32(0)
																v225 = F_make_range(m, l0, v143, v220, v223, v223)
																mBase = m.M
																v226 = m.ExcPending
																if v226 != 0 {
																	return int32(0)
																} else {
																	v228 = v225
																	m.G0 = v9 + int32(48)
																	return v228
																}
															}
														}
													}
												} else {
													v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
													v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
													v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
													v71 = F_FunctionCall2Coll(m, l0+int32(212), v68, v69, v70)
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return int32(0)
													} else {
														if v71 == int32(0) {
															v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+37)))
															v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+45)))
															if v76 == int32(0) {
																v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+46)))
																if v75&int32(1) == int32(0) {
																	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
																	if v79&int32(1)|base.B2i32(v88 == v79) != 0 {
																		v143 = v9 + int32(32)
																	} else {
																		v143 = v9 + int32(40)
																	}
																} else {
																	if v79&int32(1) == int32(0) {
																		v143 = v9 + int32(40)
																	} else {
																		v143 = v9 + int32(32)
																	}
																}
															} else {
																v98 = v9 + int32(32)
																if v75&int32(1) != 0 {
																	v143 = v98
																} else {
																	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
																	if v101&int32(1) != 0 {
																		v143 = v9 + int32(40)
																	} else {
																		v143 = v98
																	}
																}
															}
														} else {
															if v71 < int32(0) {
																v143 = v9 + int32(40)
															} else {
																v143 = v9 + int32(32)
															}
														}
														v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
														v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
														if v146 == int32(1) {
															v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
															if v145&int32(1) != 0 {
																v153 = v9 + int32(16)
																v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																if v154 == v149 {
																	v220 = v153
																} else {
																	if v149&int32(1) == int32(0) {
																		v220 = v9 + int32(24)
																	} else {
																		v220 = v153
																	}
																}
															} else {
																if v149&int32(1) == int32(0) {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v9 + int32(16)
																}
															}
															v223 = int32(0)
															v225 = F_make_range(m, l0, v143, v220, v223, v223)
															mBase = m.M
															v226 = m.ExcPending
															if v226 != 0 {
																return int32(0)
															} else {
																v228 = v225
																m.G0 = v9 + int32(48)
																return v228
															}
														} else {
															if v145&int32(1) != 0 {
																v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																if v168 != 0 {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v9 + int32(16)
																}
																v223 = int32(0)
																v225 = F_make_range(m, l0, v143, v220, v223, v223)
																mBase = m.M
																v226 = m.ExcPending
																if v226 != 0 {
																	return int32(0)
																} else {
																	v228 = v225
																	m.G0 = v9 + int32(48)
																	return v228
																}
															} else {
																v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
																v174 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
																v175 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
																v176 = F_FunctionCall2Coll(m, l0+int32(212), v173, v174, v175)
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int32(0)
																} else {
																	if v176 == int32(0) {
																		v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
																		v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
																		if v181 == int32(0) {
																			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
																			if v180&int32(1) == int32(0) {
																				v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																				if base.B2i32(v184&int32(1) == int32(0))|base.B2i32(v184 == v195) != 0 {
																					v220 = v9 + int32(16)
																				} else {
																					v220 = v9 + int32(24)
																				}
																			} else {
																				if v184&int32(1) != 0 {
																					v220 = v9 + int32(24)
																				} else {
																					v220 = v9 + int32(16)
																				}
																			}
																		} else {
																			v203 = v9 + int32(16)
																			if v180&int32(1) != 0 {
																				v220 = v203
																			} else {
																				v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																				if v206&int32(1) == int32(0) {
																					v220 = v9 + int32(24)
																				} else {
																					v220 = v203
																				}
																			}
																		}
																	} else {
																		if int32(0) < v176 {
																			v220 = v9 + int32(24)
																		} else {
																			v220 = v9 + int32(16)
																		}
																	}
																	v223 = int32(0)
																	v225 = F_make_range(m, l0, v143, v220, v223, v223)
																	mBase = m.M
																	v226 = m.ExcPending
																	if v226 != 0 {
																		return int32(0)
																	} else {
																		v228 = v225
																		m.G0 = v9 + int32(48)
																		return v228
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
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v111 = m.ExcPending
		if v111 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_range_union_internal_3), int32(0))
			mBase = m.M
			v115 = m.ExcPending
			if v115 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_union_internal_1), int32(1068), int32(_a_F_range_union_internal_2))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
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
func F_range_upper_inc(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13985(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
