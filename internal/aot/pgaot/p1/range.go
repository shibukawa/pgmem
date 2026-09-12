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
										F_errmsg(m, int32(28442), v9)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											F_errfinish(m, int32(494319), int32(21618), int32(29354))
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
											F_errmsg(m, int32(28442), v9)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												F_errfinish(m, int32(494319), int32(21618), int32(29354))
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
												F_errmsg(m, int32(28442), v9)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return
												} else {
													F_errfinish(m, int32(494319), int32(21618), int32(29354))
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
									F_errmsg(m, int32(28442), v9)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										F_errfinish(m, int32(494319), int32(21618), int32(29354))
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
										F_errmsg(m, int32(28442), v9)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											F_errfinish(m, int32(494319), int32(21618), int32(29354))
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
											F_errmsg(m, int32(28442), v9)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												F_errfinish(m, int32(494319), int32(21618), int32(29354))
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
	var v27 int32
	_ = v27
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
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
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
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
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L71
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[158]))
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
	v50 = *(*int64)(unsafe.Add(mBase, _consts[315]))
	v55 = int32(0)
	v56 = v4
	v61 = v4
	v62 = v50
	goto L16
L5:
	;
	return int32(0)
L6:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v27 == int32(0) {
		v46 = v26
		v47 = v27
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v47-v46 != 0 {
		goto L1
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v26 != v27 {
		v46 = v26
		v47 = v27
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v31 = v17
	v32 = v20
	goto L11
L11:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	if v36 == int32(0) {
		v46 = v35
		v47 = v36
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v46 = v35
	v47 = v36
	goto L8
L13:
	;
	v39 = int32(1)
	if v35 == v36 {
		v31 = v31 + v39
		v32 = v32 + v39
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L4
L16:
	;
	v64 = F_RangeVarGetCreationNamespace(m, l0)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	F_RangeVarAdjustRelationPersistence(m, l0, v143)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L67
	}
L18:
	;
	if l2 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v68 = F_get_relname_relid(m, v67, v64)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L22
	}
L20:
	;
	v70 = int32(0)
	goto L21
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v72 == int32(0) {
		v143 = v64
		v144 = v70
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v70 = v68
	goto L21
L23:
	;
	goto L17
L24:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v79 = F_object_aclcheck(m, int32(2615), v64, v77, int64(512))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	if v79 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v82 = F_get_namespace_name(m, v64)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v61 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	F_aclcheck_error(m, v79, int32(36), v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	if v64 != v55 {
		goto L44
	} else {
		goto L45
	}
L32:
	;
	if v70 != v56 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if v64 != v55 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	if v64 != v55 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v143 = v55
	v144 = v56
	goto L23
L36:
	;
	F_UnlockDatabaseObject(m, int32(2615), v55, int32(1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if l1 == int32(0) {
		goto L31
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	if v56 == int32(0) {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	if v70 == v56 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	F_UnlockRelationOid(m, v56, l1)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	goto L31
L44:
	;
	F_LockDatabaseObject(m, int32(2615), v64, int32(1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if l1 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	v141 = *(*int64)(unsafe.Add(mBase, _consts[315]))
	if v62 != v141 {
		v55 = v64
		v56 = v70
		v61 = int32(1)
		v62 = v141
		goto L16
	} else {
		goto L66
	}
L49:
	;
	if v70 == int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v114 = F_object_ownercheck(m, int32(1259), v70, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	if v114 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v119 = F_get_rel_relkind(m, v70)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v70 == v56 {
		goto L48
	} else {
		goto L64
	}
L55:
	;
	switch v119 - int32(73) {
	case 0, 32:
		goto L62
	default:
		v130 = int32(41)
		goto L57
	case 10:
		goto L61
	case 29:
		goto L58
	case 36:
		goto L59
	case 45:
		goto L60
	}
L56:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_aclcheck_error(m, int32(2), v132, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L63
	}
L57:
	;
	v132 = v130
	goto L56
L58:
	;
	v130 = int32(18)
	goto L57
L59:
	;
	v132 = int32(23)
	goto L56
L60:
	;
	v132 = int32(51)
	goto L56
L61:
	;
	v132 = int32(37)
	goto L56
L62:
	;
	v132 = int32(20)
	goto L56
L63:
	;
	goto L54
L64:
	;
	F_LockRelationOid(m, v70, l1)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	goto L48
L66:
	;
	v143 = v64
	v144 = v70
	goto L23
L67:
	;
	if l2 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v144
	goto L70
L69:
	;
	goto L70
L70:
	;
	m.G0 = v15 + int32(16)
	return v143
L71:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	v162 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v163
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v162
	F_errmsg(m, int32(690582), v15)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(499353), int32(760), int32(418409))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RangeVarGetRelidExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
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
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int64
	_ = v317
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
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
	v313 = m.ExcPending
	if v313 != 0 {
		goto L5
	} else {
		goto L102
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[158]))
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
	v54 = l2 & int32(1)
	v58 = *(*int64)(unsafe.Add(mBase, _consts[315]))
	v66 = int32(0)
	v74 = int32(0)
	v75 = v58
	goto L18
L5:
	;
	return int32(0)
L6:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v31 == int32(0) {
		v50 = v30
		v51 = v31
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v51-v50 != 0 {
		goto L1
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v30 != v31 {
		v50 = v30
		v51 = v31
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v35 = v21
	v36 = v24
	goto L11
L11:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v40 == int32(0) {
		v50 = v39
		v51 = v40
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v50 = v39
	v51 = v40
	goto L8
L13:
	;
	v43 = int32(1)
	if v39 == v40 {
		v35 = v35 + v43
		v36 = v36 + v43
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L4
L16:
	;
	m.G0 = v19 + int32(80)
	return v301
L17:
	;
	F_errfinish(m, int32(499353), v297, int32(461350))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L5
	} else {
		goto L101
	}
L18:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v77 == int32(116) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	if v257 != 0 {
		v301 = v257
		goto L16
	} else {
		goto L87
	}
L20:
	;
	if l3 != 0 {
		goto L49
	} else {
		goto L50
	}
L21:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L39
	}
L22:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v122 = F_get_relname_relid(m, v121, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L38
	}
L23:
	;
	v80 = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	if v82 == v80 {
		v185 = v80
		goto L20
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v109 == int32(0) {
		goto L21
	} else {
		goto L34
	}
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v85 == int32(0) {
		v120 = v82
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v88 = F_LookupExplicitNamespace(m, v85, v54)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	if v88 == v91 {
		v120 = v88
		goto L22
	} else {
		goto L29
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errmsg(m, int32(381773), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(499353), int32(519), int32(461350))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v112 = F_LookupExplicitNamespace(m, v109, v54)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	if v54 == int32(0) {
		v120 = v112
		goto L22
	} else {
		goto L36
	}
L36:
	;
	v116 = int32(0)
	if v112 == v116 {
		v185 = v116
		goto L20
	} else {
		goto L37
	}
L37:
	;
	v120 = v112
	goto L22
L38:
	;
	v185 = v122
	goto L20
L39:
	;
	v127 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	if v129 == v127 {
		v185 = v127
		goto L20
	} else {
		goto L40
	}
L40:
	;
	v132 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v132 < v133 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v143 = v132
	goto L44
L42:
	;
	goto L43
L43:
	;
	v185 = int32(0)
	goto L20
L44:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152+v143<<(uint(int32(2))%32))))
	v157 = F_get_relname_relid(m, v124, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	if v157 != 0 {
		v185 = v157
		goto L20
	} else {
		goto L47
	}
L47:
	;
	v160 = v143 + int32(1)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v160 < v161 {
		v143 = v160
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	m.T0[l3].(func(*base.Module, int32, int32, int32, int32))(m, l0, v185, v66, l4)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L5
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if l1 == int32(0) {
		v257 = v185
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	goto L19
L54:
	;
	if v74 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if v185 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L56:
	;
	if v185 == v66 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v257 = v66
	goto L53
L58:
	;
	goto L59
L59:
	;
	if v66 == int32(0) {
		goto L55
	} else {
		goto L60
	}
L60:
	;
	F_UnlockRelationOid(m, v66, l1)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	goto L55
L62:
	;
	v255 = *(*int64)(unsafe.Add(mBase, _consts[315]))
	if v75 != v255 {
		v66 = v185
		v74 = int32(1)
		v75 = v255
		goto L18
	} else {
		goto L86
	}
L63:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L5
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if l2&int32(6) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L62
L67:
	;
	F_LockRelationOid(m, v185, l1)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L5
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v215 = F_ConditionalLockRelationOid(m, v185, l1)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L71
	}
L70:
	;
	goto L62
L71:
	;
	if v215 != 0 {
		goto L62
	} else {
		goto L72
	}
L72:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v218 = int32(0)
	if l2&int32(4) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v223 = int32(14)
	goto L75
L74:
	;
	v223 = int32(21)
	goto L75
L75:
	;
	v225 = F_errstart(m, v223, int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	if v217 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if v225 == int32(0) {
		v301 = v218
		goto L16
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v225 == int32(0) {
		v301 = v218
		goto L16
	} else {
		goto L83
	}
L80:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	v232 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v232
	F_errmsg(m, int32(691060), v19+int32(48))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	v292 = v218
	v297 = int32(601)
	goto L17
L83:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v245
	F_errmsg(m, int32(705279), v19+int32(32))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v292 = v218
	v297 = int32(606)
	goto L17
L86:
	;
	v257 = v185
	goto L53
L87:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v261 = int32(0)
	if v54 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v264 = int32(14)
	goto L90
L89:
	;
	v264 = int32(21)
	goto L90
L90:
	;
	v266 = F_errstart(m, v264, int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	if v260 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v266 == int32(0) {
		v301 = v261
		goto L16
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if v266 == int32(0) {
		v301 = v261
		goto L16
	} else {
		goto L98
	}
L95:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	v273 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v273
	F_errmsg(m, int32(70370), v19+int32(16))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	v292 = v261
	v297 = int32(634)
	goto L17
L98:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v286
	F_errmsg(m, int32(71601), v19)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	v292 = v261
	v297 = int32(639)
	goto L17
L101:
	;
	v301 = v292
	goto L16
L102:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	v317 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v318
	*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v317
	F_errmsg(m, int32(690582), v19-int32(-64))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(499353), int32(464), int32(461350))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
			if int32(32768) <= v25 {
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
						*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(32767)
						F_errmsg(m, int32(148618), v16+int32(16))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496161), int32(2260), int32(275936))
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
						v46 = int32(0)
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
						if v48 != 0 {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							v50 = v49
						} else {
							v50 = v46
						}
						if l1 != 0 {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v52 = v51
						} else {
							v52 = v46
						}
						if v50 < v52 {
							v54 = F_list_copy_tail(m, l1, v50)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v56 = F_list_concat(m, v48, v54)
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
									if v61 < v50 {
										v63 = int32(0)
										F_errstart_cold(m, int32(21), v63)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(393348))
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
												*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v50
												*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
												*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
												F_errmsg(m, int32(457055), v16)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496161), int32(2285), int32(275936))
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
							if v61 < v50 {
								v63 = int32(0)
								F_errstart_cold(m, int32(21), v63)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(393348))
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
										*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v50
										*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
										*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
										F_errmsg(m, int32(457055), v16)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496161), int32(2285), int32(275936))
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
					v43 = F_makeAlias(m, int32(275687), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = v43
						v46 = int32(0)
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
						if v48 != 0 {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
							v50 = v49
						} else {
							v50 = v46
						}
						if l1 != 0 {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v52 = v51
						} else {
							v52 = v46
						}
						if v50 < v52 {
							v54 = F_list_copy_tail(m, l1, v50)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v56 = F_list_concat(m, v48, v54)
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
									if v61 < v50 {
										v63 = int32(0)
										F_errstart_cold(m, int32(21), v63)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(393348))
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
												*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v50
												*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
												*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
												F_errmsg(m, int32(457055), v16)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496161), int32(2285), int32(275936))
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
							if v61 < v50 {
								v63 = int32(0)
								F_errstart_cold(m, int32(21), v63)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(393348))
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
										*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v50
										*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
										*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
										F_errmsg(m, int32(457055), v16)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496161), int32(2285), int32(275936))
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
					v46 = int32(0)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					if v48 != 0 {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
						v50 = v49
					} else {
						v50 = v46
					}
					if l1 != 0 {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v52 = v51
					} else {
						v52 = v46
					}
					if v50 < v52 {
						v54 = F_list_copy_tail(m, l1, v50)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v56 = F_list_concat(m, v48, v54)
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
								if v61 < v50 {
									v63 = int32(0)
									F_errstart_cold(m, int32(21), v63)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(393348))
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
											*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v50
											*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
											*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
											F_errmsg(m, int32(457055), v16)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496161), int32(2285), int32(275936))
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
						if v61 < v50 {
							v63 = int32(0)
							F_errstart_cold(m, int32(21), v63)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(393348))
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
									*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v50
									*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
									*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
									F_errmsg(m, int32(457055), v16)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(496161), int32(2285), int32(275936))
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
				v43 = F_makeAlias(m, int32(275687), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = v43
					v46 = int32(0)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					if v48 != 0 {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
						v50 = v49
					} else {
						v50 = v46
					}
					if l1 != 0 {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						v52 = v51
					} else {
						v52 = v46
					}
					if v50 < v52 {
						v54 = F_list_copy_tail(m, l1, v50)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v56 = F_list_concat(m, v48, v54)
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
								if v61 < v50 {
									v63 = int32(0)
									F_errstart_cold(m, int32(21), v63)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(393348))
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
											*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v50
											*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
											*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
											F_errmsg(m, int32(457055), v16)
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496161), int32(2285), int32(275936))
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
						if v61 < v50 {
							v63 = int32(0)
							F_errstart_cold(m, int32(21), v63)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(393348))
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
									*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v50
									*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v73
									*(*int32)(unsafe.Add(mBase, uint32(v16))) = v71
									F_errmsg(m, int32(457055), v16)
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(496161), int32(2285), int32(275936))
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	v5 = l4
	v6 = l5
	v13 = F_palloc0(m, int32(136))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(101)
		if l3 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			v23 = v19
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
			v23 = v20 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l3
		v25 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v25
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+20)) = uint8(v5)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v27
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+119)))
		*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = l2
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+21)) = uint8(v31)
		v35 = F_makeAlias(m, v23, v25)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v35
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
			F_buildRelationAliases(m, v38, l3, v35)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+125)) = uint8(v6)
				v42 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+124)) = uint8(v42)
				v45 = F_palloc0(m, int32(40))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v45))) = int32(102)
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v49
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+20)))
					*(*uint8)(unsafe.Add(mBase, uint32(v45)+8)) = uint8(v51)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v54 = F_lappend(m, v53, v45)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v54
						if v54 != 0 {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
							v59 = v58
						} else {
							v59 = int32(0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v59
						*(*int64)(unsafe.Add(mBase, uint32(v45)+16)) = int64(2)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v64 = F_lappend(m, v63, v13)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v64
							if v64 != 0 {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
								v68 = v67
							} else {
								v68 = int32(0)
							}
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
							v73 = F_palloc0(m, v70<<(uint(int32(5))%32))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if int32(0) < v70 {
									v81 = int32(0)
									for {
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
										v97 = v69 + int32(20) + v91<<(uint(int32(4))%32) + v81*int32(100)
										v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+91)))
										if v98 == int32(0) {
											v103 = v73 + v81<<(uint(int32(5))%32)
											v105 = v81 + int32(1)
											*(*uint16)(unsafe.Add(mBase, uint32(v103)+4)) = uint16(v105)
											*(*int32)(unsafe.Add(mBase, uint32(v103))) = v68
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v97)+68))
											*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v108
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v97)+76))
											*(*int32)(unsafe.Add(mBase, uint32(v103)+12)) = v110
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v97)+96))
											*(*uint16)(unsafe.Add(mBase, uint32(v103)+28)) = uint16(v105)
											*(*int32)(unsafe.Add(mBase, uint32(v103)+24)) = v68
											*(*int32)(unsafe.Add(mBase, uint32(v103)+16)) = v112
										} else {
										}
										v120 = v81 + int32(1)
										if v120 != v70 {
											v81 = v120
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								v134 = F_palloc(m, int32(28))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return int32(0)
								} else {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v134)+20)) = int64(16777473)
									*(*int32)(unsafe.Add(mBase, uint32(v134)+16)) = v73
									*(*int32)(unsafe.Add(mBase, uint32(v134)+12)) = v45
									*(*int32)(unsafe.Add(mBase, uint32(v134)+8)) = v68
									*(*int32)(unsafe.Add(mBase, uint32(v134)+4)) = v13
									*(*int32)(unsafe.Add(mBase, uint32(v134))) = v136
									return v134
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
						F_errmsg(m, int32(148570), v12+int32(16))
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
								F_errfinish(m, int32(496161), int32(2086), int32(489936))
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
											F_errcode(m, int32(393348))
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
													v85 = int32(540235)
												} else {
													v85 = int32(540500)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
												F_errmsg(m, int32(456927), v12)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496161), int32(2115), int32(489936))
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
									F_errcode(m, int32(393348))
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
											v85 = int32(540235)
										} else {
											v85 = int32(540500)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
										F_errmsg(m, int32(456927), v12)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496161), int32(2115), int32(489936))
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
						v42 = int32(392016)
					} else {
						v42 = int32(391475)
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
												F_errcode(m, int32(393348))
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
														v85 = int32(540235)
													} else {
														v85 = int32(540500)
													}
													*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
													F_errmsg(m, int32(456927), v12)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(496161), int32(2115), int32(489936))
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
										F_errcode(m, int32(393348))
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
												v85 = int32(540235)
											} else {
												v85 = int32(540500)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
											F_errmsg(m, int32(456927), v12)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496161), int32(2115), int32(489936))
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
										F_errcode(m, int32(393348))
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
												v85 = int32(540235)
											} else {
												v85 = int32(540500)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
											F_errmsg(m, int32(456927), v12)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496161), int32(2115), int32(489936))
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
								F_errcode(m, int32(393348))
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
										v85 = int32(540235)
									} else {
										v85 = int32(540500)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
									F_errmsg(m, int32(456927), v12)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(496161), int32(2115), int32(489936))
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
					v42 = int32(392016)
				} else {
					v42 = int32(391475)
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
											F_errcode(m, int32(393348))
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
													v85 = int32(540235)
												} else {
													v85 = int32(540500)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
												F_errmsg(m, int32(456927), v12)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(496161), int32(2115), int32(489936))
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
									F_errcode(m, int32(393348))
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
											v85 = int32(540235)
										} else {
											v85 = int32(540500)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v85
										F_errmsg(m, int32(456927), v12)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496161), int32(2115), int32(489936))
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
	var v56 float64
	_ = v56
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
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int64
	_ = v152
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 float64
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 float64
	_ = v170
	var v172 float64
	_ = v172
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 float64
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 float64
	_ = v204
	var v205 float64
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 float32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
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
	var v384 float64
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v468 int32
	_ = v468
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
	v56 = float64(0)
	v58 = int32(0)
	v61 = v5
	v63 = v5
	v64 = v5
	v65 = v5
	goto L10
L10:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	if int32(0) < v193 {
		goto L48
	} else {
		goto L49
	}
L12:
	;
	v82 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v58, v27+int32(31))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+31)))
	if v84 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v198 = v58 + int32(1)
	if v198 != l2 {
		v56 = v187
		v58 = v198
		v61 = v189
		v63 = v191
		v64 = v192
		v65 = v193
		goto L10
	} else {
		goto L47
	}
L15:
	;
	v187 = v56
	v189 = v61
	v191 = v63
	v192 = v64 + int32(1)
	v193 = v65
	goto L14
L16:
	;
	goto L17
L17:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v89 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v117 = F_pg_detoast_datum(m, v82)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L27
	}
L19:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if base.Ui32((v93-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v115 = int32(6)
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v107 = int32(1)
	if v89&v107 != 0 {
		v115 = int32(base.Ui32(v89) >> (uint(v107) % 32))
		goto L18
	} else {
		goto L26
	}
L22:
	;
	v100 = int32(18)
	if v93&int32(255) == v100 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v106 = v100
	goto L25
L24:
	;
	v106 = int32(2)
	goto L25
L25:
	;
	v115 = v106
	goto L18
L26:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v115 = int32(base.Ui32(v111) >> (uint(int32(2)) % 32))
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
	v187 = base.F64_add(v56, base.F64_convert_i32_u(v115))
	v189 = v180
	v191 = v182
	v192 = v64
	v193 = v65 + int32(1)
	goto L14
L29:
	;
	v180 = v61
	v182 = v63 + int32(1)
	goto L28
L30:
	;
	v150 = v61 << (uint(int32(3)) % 32)
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v27)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v39+v150))) = v152
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v27)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v150+v41))) = v155
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+20)))
	if v158 != 0 {
		v172 = math.Float64frombits(uint64(0x7ff0000000000000))
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	if v121 == int32(0) {
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_range_deserialize(m, v35, v117, v27+int32(16), v27+int32(8), v27+int32(30))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L37
	}
L34:
	;
	F_multirange_get_bounds(m, v35, v117, int32(0), v27+int32(16), v27)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	F_multirange_get_bounds(m, v35, v117, v129-int32(1), v27, v27+int32(8))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+30)) = uint8(v136)
	goto L30
L37:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+30)))
	if v146&int32(1) != 0 {
		goto L29
	} else {
		goto L38
	}
L38:
	;
	goto L30
L39:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v150+v43))) = v172
	v180 = v61 + int32(1)
	v182 = v63
	goto L28
L40:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+12)))
	if v161 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v162 = math.Float64frombits(uint64(0x7ff0000000000000))
	goto L43
L42:
	;
	v162 = float64(1)
	goto L43
L43:
	;
	if v161 != 0 {
		v172 = v162
		goto L39
	} else {
		goto L44
	}
L44:
	;
	if v36 == int32(0) {
		v172 = v162
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v35)+208))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v168 = F_FunctionCall2Coll(m, v35+int32(268), v165, v166, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v170 = *(*float64)(unsafe.Add(mBase, uint32(v168)))
	v172 = v170
	goto L39
L47:
	;
	goto L11
L48:
	;
	v202 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v202)
	v204 = base.F64_convert_i32_u(v193)
	v205 = base.F64_div(v187, v204)
	if base.F64_lt(base.F64_abs(v205), float64(2.147483648e+09)) != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	goto L50
L50:
	;
	if v192 <= int32(0) {
		goto L8
	} else {
		goto L87
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v211
	v217 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v192), base.F64_convert_i32_s(l2)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v217
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = base.F32_neg(base.F32_sub(float32(1), v217))
	v224 = int32(4515248)
	v225 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v227
	if int32(2) <= v189 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v209 = base.I32_trunc_f64_s(v205)
	v211 = v209
	goto L51
L53:
	;
	goto L54
L54:
	;
	v211 = int32(-2147483648)
	goto L51
L55:
	;
	v428 = l0 + v412<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v428)+164)) = v410
	v430 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v428)+84)) = v430
	*(*int32)(unsafe.Add(mBase, uint32(v428-int32(-64)))) = int32(672)
	*(*int32)(unsafe.Add(mBase, uint32(v428)+184)) = int32(701)
	*(*int32)(unsafe.Add(mBase, uint32(v428)+144)) = v411
	v441 = l0 + v412<<(uint(int32(1))%32)
	v442 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+204)) = uint16(v442)
	v444 = l0 + v412
	v445 = int32(100)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+219)) = uint8(v445)
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+214)) = uint8(v430)
	v450 = F_palloc(m, int32(4))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L86
	}
L56:
	;
	F_qsort_interruptible(m, v39, v189, int32(8), int32(1491), v35)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v399 = F_palloc(m, int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L85
	}
L59:
	;
	F_qsort_interruptible(m, v41, v189, int32(8), int32(1491), v35)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v239 = int32(1)
	v241 = v189 - v239
	if v29 < v189 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v245 = v29 + v239
	goto L63
L62:
	;
	v245 = v189
	goto L63
L63:
	;
	v247 = v245 - int32(1)
	v248 = base.I32_div_s(v241, v247)
	v250 = v241 - v248*v247
	v252 = v245 << (uint(int32(2)) % 32)
	v253 = F_palloc(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v255 = int32(0)
	v256 = base.B2i32(v245 <= v255)
	if v256 == v255 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v259 = int32(0)
	v263 = v259
	v264 = v259
	v266 = v259
	goto L68
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v253
	v333 = int32(7)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v333)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v245
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v336
	v338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+204)) = uint16(v338)
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+214)) = uint8(v340)
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+11)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+219)) = uint8(v342)
	v344 = int32(0)
	F_qsort_interruptible(m, v43, v189, int32(8), int32(1492), v344)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L4
	} else {
		goto L75
	}
L68:
	;
	v290 = v264 << (uint(int32(3)) % 32)
	v293 = int32(0)
	v295 = F_range_serialize(m, v35, v39+v290, v290+v41, v293, v293)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L70
	}
L69:
	;
	goto L67
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253+v266<<(uint(int32(2))%32)))) = v295
	v298 = v263 + v250
	v299 = base.B2i32(v247 <= v298)
	if v247 <= v298 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v303 = v247
	goto L73
L72:
	;
	v303 = int32(0)
	goto L73
L73:
	;
	v306 = v266 + int32(1)
	if v306 != v245 {
		v263 = v298 - v303
		v264 = v299 + (v264 + v248)
		v266 = v306
		goto L68
	} else {
		goto L74
	}
L74:
	;
	goto L69
L75:
	;
	v350 = F_palloc(m, v252)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	if v245 <= v255 {
		v410 = v350
		v411 = v245
		v412 = v239
		goto L55
	} else {
		goto L77
	}
L77:
	;
	v352 = int32(0)
	v356 = v352
	v358 = v344
	v360 = v352
	goto L78
L78:
	;
	v384 = *(*float64)(unsafe.Add(mBase, uint32(v43+v356<<(uint(int32(3))%32))))
	v385 = F_Float8GetDatum(m, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L80
	}
L79:
	;
	v410 = v350
	v411 = v245
	v412 = v239
	goto L55
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v350+v358<<(uint(int32(2))%32)))) = v385
	v388 = v360 + v250
	v389 = base.B2i32(v247 <= v388)
	if v247 <= v388 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v393 = v247
	goto L83
L82:
	;
	v393 = int32(0)
	goto L83
L83:
	;
	v396 = v358 + int32(1)
	if v396 != v245 {
		v356 = v389 + (v356 + v248)
		v358 = v396
		v360 = v388 - v393
		goto L78
	} else {
		goto L84
	}
L84:
	;
	goto L79
L85:
	;
	v410 = v399
	v411 = int32(0)
	v412 = int32(0)
	goto L55
L86:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v450))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v191), v204))
	*(*int32)(unsafe.Add(mBase, uint32(v428)+104)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v428)+124)) = v450
	v458 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v441)+52)) = uint16(v458)
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v225
	goto L8
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = int64(1065353216)
	v468 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v468)
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
	F_scanner_yyerror(m, int32(212048), l3)
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
	F_errmsg(m, int32(204655), v11)
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
	F_errfinish(m, int32(26978), int32(19397), int32(382154))
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
								F_errmsg_internal(m, int32(370267), v9)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493991), int32(1776), int32(398753))
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
							F_errmsg_internal(m, int32(370267), v9)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493991), int32(1776), int32(398753))
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
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
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
					v168 = int32(1)
					m.G0 = v8 + int32(48)
					return v168
				} else {
					v32 = int32(0)
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
					if v33 != 0 {
						v168 = v32
						m.G0 = v8 + int32(48)
						return v168
					} else {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+44)))
						if v35 == int32(1) {
							v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+46)))
							if v34&int32(1) != 0 {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
								if v38 == v41 {
									v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
									v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
									if v88 == int32(1) {
										v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
										if v87&int32(1) != 0 {
											v94 = int32(1)
											v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
											if v95 == v91&int32(255) {
												v168 = v94
											} else {
												if v91&int32(1) != 0 {
													v168 = int32(0)
												} else {
													v168 = v94
												}
											}
										} else {
											v101 = int32(1)
											if v91&v101 != 0 {
												v168 = int32(0)
											} else {
												v168 = v101
											}
										}
										m.G0 = v8 + int32(48)
										return v168
									} else {
										if v87&int32(1) != 0 {
											v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
											if v106 == int32(0) {
												v168 = int32(0)
											} else {
												v168 = int32(1)
											}
											m.G0 = v8 + int32(48)
											return v168
										} else {
											v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											v114 = F_FunctionCall2Coll(m, l0+int32(212), v111, v112, v113)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												if v114 == int32(0) {
													v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
													v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
													if v119 == int32(0) {
														v122 = int32(1)
														v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
														if v118&v122 == int32(0) {
															v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v128 == v123&int32(255) {
																v168 = v122
															} else {
																if v123&int32(1) != 0 {
																	v168 = v122
																} else {
																	v168 = int32(0)
																}
															}
														} else {
															if v123&int32(1) == int32(0) {
																v168 = int32(0)
															} else {
																v168 = v122
															}
														}
													} else {
														v138 = int32(1)
														if v118&v138 != 0 {
															v168 = v138
														} else {
															v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v141 != 0 {
																v168 = int32(0)
															} else {
																v168 = v138
															}
														}
													}
												} else {
													if v114 < int32(0) {
														v168 = int32(0)
													} else {
														v168 = int32(1)
													}
												}
												m.G0 = v8 + int32(48)
												return v168
											}
										}
									}
								} else {
									if v38&int32(1) != 0 {
										v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
										v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
										if v88 == int32(1) {
											v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
											if v87&int32(1) != 0 {
												v94 = int32(1)
												v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
												if v95 == v91&int32(255) {
													v168 = v94
												} else {
													if v91&int32(1) != 0 {
														v168 = int32(0)
													} else {
														v168 = v94
													}
												}
											} else {
												v101 = int32(1)
												if v91&v101 != 0 {
													v168 = int32(0)
												} else {
													v168 = v101
												}
											}
											m.G0 = v8 + int32(48)
											return v168
										} else {
											if v87&int32(1) != 0 {
												v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
												if v106 == int32(0) {
													v168 = int32(0)
												} else {
													v168 = int32(1)
												}
												m.G0 = v8 + int32(48)
												return v168
											} else {
												v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
												v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
												v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
												v114 = F_FunctionCall2Coll(m, l0+int32(212), v111, v112, v113)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													if v114 == int32(0) {
														v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
														v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
														if v119 == int32(0) {
															v122 = int32(1)
															v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
															if v118&v122 == int32(0) {
																v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																if v128 == v123&int32(255) {
																	v168 = v122
																} else {
																	if v123&int32(1) != 0 {
																		v168 = v122
																	} else {
																		v168 = int32(0)
																	}
																}
															} else {
																if v123&int32(1) == int32(0) {
																	v168 = int32(0)
																} else {
																	v168 = v122
																}
															}
														} else {
															v138 = int32(1)
															if v118&v138 != 0 {
																v168 = v138
															} else {
																v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																if v141 != 0 {
																	v168 = int32(0)
																} else {
																	v168 = v138
																}
															}
														}
													} else {
														if v114 < int32(0) {
															v168 = int32(0)
														} else {
															v168 = int32(1)
														}
													}
													m.G0 = v8 + int32(48)
													return v168
												}
											}
										}
									} else {
										v168 = v32
										m.G0 = v8 + int32(48)
										return v168
									}
								}
							} else {
								if v38&int32(1) != 0 {
									v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
									v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
									if v88 == int32(1) {
										v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
										if v87&int32(1) != 0 {
											v94 = int32(1)
											v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
											if v95 == v91&int32(255) {
												v168 = v94
											} else {
												if v91&int32(1) != 0 {
													v168 = int32(0)
												} else {
													v168 = v94
												}
											}
										} else {
											v101 = int32(1)
											if v91&v101 != 0 {
												v168 = int32(0)
											} else {
												v168 = v101
											}
										}
										m.G0 = v8 + int32(48)
										return v168
									} else {
										if v87&int32(1) != 0 {
											v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
											if v106 == int32(0) {
												v168 = int32(0)
											} else {
												v168 = int32(1)
											}
											m.G0 = v8 + int32(48)
											return v168
										} else {
											v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											v114 = F_FunctionCall2Coll(m, l0+int32(212), v111, v112, v113)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												if v114 == int32(0) {
													v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
													v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
													if v119 == int32(0) {
														v122 = int32(1)
														v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
														if v118&v122 == int32(0) {
															v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v128 == v123&int32(255) {
																v168 = v122
															} else {
																if v123&int32(1) != 0 {
																	v168 = v122
																} else {
																	v168 = int32(0)
																}
															}
														} else {
															if v123&int32(1) == int32(0) {
																v168 = int32(0)
															} else {
																v168 = v122
															}
														}
													} else {
														v138 = int32(1)
														if v118&v138 != 0 {
															v168 = v138
														} else {
															v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v141 != 0 {
																v168 = int32(0)
															} else {
																v168 = v138
															}
														}
													}
												} else {
													if v114 < int32(0) {
														v168 = int32(0)
													} else {
														v168 = int32(1)
													}
												}
												m.G0 = v8 + int32(48)
												return v168
											}
										}
									}
								} else {
									v168 = v32
									m.G0 = v8 + int32(48)
									return v168
								}
							}
						} else {
							if v34&int32(1) != 0 {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
								if v49 == int32(0) {
									v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
									v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
									if v88 == int32(1) {
										v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
										if v87&int32(1) != 0 {
											v94 = int32(1)
											v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
											if v95 == v91&int32(255) {
												v168 = v94
											} else {
												if v91&int32(1) != 0 {
													v168 = int32(0)
												} else {
													v168 = v94
												}
											}
										} else {
											v101 = int32(1)
											if v91&v101 != 0 {
												v168 = int32(0)
											} else {
												v168 = v101
											}
										}
										m.G0 = v8 + int32(48)
										return v168
									} else {
										if v87&int32(1) != 0 {
											v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
											if v106 == int32(0) {
												v168 = int32(0)
											} else {
												v168 = int32(1)
											}
											m.G0 = v8 + int32(48)
											return v168
										} else {
											v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
											v114 = F_FunctionCall2Coll(m, l0+int32(212), v111, v112, v113)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												if v114 == int32(0) {
													v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
													v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
													if v119 == int32(0) {
														v122 = int32(1)
														v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
														if v118&v122 == int32(0) {
															v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v128 == v123&int32(255) {
																v168 = v122
															} else {
																if v123&int32(1) != 0 {
																	v168 = v122
																} else {
																	v168 = int32(0)
																}
															}
														} else {
															if v123&int32(1) == int32(0) {
																v168 = int32(0)
															} else {
																v168 = v122
															}
														}
													} else {
														v138 = int32(1)
														if v118&v138 != 0 {
															v168 = v138
														} else {
															v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v141 != 0 {
																v168 = int32(0)
															} else {
																v168 = v138
															}
														}
													}
												} else {
													if v114 < int32(0) {
														v168 = int32(0)
													} else {
														v168 = int32(1)
													}
												}
												m.G0 = v8 + int32(48)
												return v168
											}
										}
									}
								} else {
									v168 = v32
									m.G0 = v8 + int32(48)
									return v168
								}
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
								v57 = F_FunctionCall2Coll(m, l0+int32(212), v54, v55, v56)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									if v57 == int32(0) {
										v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+25)))
										v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+45)))
										if v62 == int32(0) {
											v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+46)))
											if v61&int32(1) == int32(0) {
												v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
												if v65 == v70 {
													v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
													v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
													if v88 == int32(1) {
														v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
														if v87&int32(1) != 0 {
															v94 = int32(1)
															v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v95 == v91&int32(255) {
																v168 = v94
															} else {
																if v91&int32(1) != 0 {
																	v168 = int32(0)
																} else {
																	v168 = v94
																}
															}
														} else {
															v101 = int32(1)
															if v91&v101 != 0 {
																v168 = int32(0)
															} else {
																v168 = v101
															}
														}
														m.G0 = v8 + int32(48)
														return v168
													} else {
														if v87&int32(1) != 0 {
															v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v106 == int32(0) {
																v168 = int32(0)
															} else {
																v168 = int32(1)
															}
															m.G0 = v8 + int32(48)
															return v168
														} else {
															v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
															v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
															v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
															v114 = F_FunctionCall2Coll(m, l0+int32(212), v111, v112, v113)
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return int32(0)
															} else {
																if v114 == int32(0) {
																	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
																	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
																	if v119 == int32(0) {
																		v122 = int32(1)
																		v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
																		if v118&v122 == int32(0) {
																			v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																			if v128 == v123&int32(255) {
																				v168 = v122
																			} else {
																				if v123&int32(1) != 0 {
																					v168 = v122
																				} else {
																					v168 = int32(0)
																				}
																			}
																		} else {
																			if v123&int32(1) == int32(0) {
																				v168 = int32(0)
																			} else {
																				v168 = v122
																			}
																		}
																	} else {
																		v138 = int32(1)
																		if v118&v138 != 0 {
																			v168 = v138
																		} else {
																			v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																			if v141 != 0 {
																				v168 = int32(0)
																			} else {
																				v168 = v138
																			}
																		}
																	}
																} else {
																	if v114 < int32(0) {
																		v168 = int32(0)
																	} else {
																		v168 = int32(1)
																	}
																}
																m.G0 = v8 + int32(48)
																return v168
															}
														}
													}
												} else {
													if v65&int32(1) == int32(0) {
														v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
														v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
														if v88 == int32(1) {
															v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
															if v87&int32(1) != 0 {
																v94 = int32(1)
																v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																if v95 == v91&int32(255) {
																	v168 = v94
																} else {
																	if v91&int32(1) != 0 {
																		v168 = int32(0)
																	} else {
																		v168 = v94
																	}
																}
															} else {
																v101 = int32(1)
																if v91&v101 != 0 {
																	v168 = int32(0)
																} else {
																	v168 = v101
																}
															}
															m.G0 = v8 + int32(48)
															return v168
														} else {
															if v87&int32(1) != 0 {
																v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																if v106 == int32(0) {
																	v168 = int32(0)
																} else {
																	v168 = int32(1)
																}
																m.G0 = v8 + int32(48)
																return v168
															} else {
																v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
																v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
																v114 = F_FunctionCall2Coll(m, l0+int32(212), v111, v112, v113)
																mBase = m.M
																v115 = m.ExcPending
																if v115 != 0 {
																	return int32(0)
																} else {
																	if v114 == int32(0) {
																		v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
																		v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
																		if v119 == int32(0) {
																			v122 = int32(1)
																			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
																			if v118&v122 == int32(0) {
																				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																				if v128 == v123&int32(255) {
																					v168 = v122
																				} else {
																					if v123&int32(1) != 0 {
																						v168 = v122
																					} else {
																						v168 = int32(0)
																					}
																				}
																			} else {
																				if v123&int32(1) == int32(0) {
																					v168 = int32(0)
																				} else {
																					v168 = v122
																				}
																			}
																		} else {
																			v138 = int32(1)
																			if v118&v138 != 0 {
																				v168 = v138
																			} else {
																				v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																				if v141 != 0 {
																					v168 = int32(0)
																				} else {
																					v168 = v138
																				}
																			}
																		}
																	} else {
																		if v114 < int32(0) {
																			v168 = int32(0)
																		} else {
																			v168 = int32(1)
																		}
																	}
																	m.G0 = v8 + int32(48)
																	return v168
																}
															}
														}
													} else {
														v168 = v32
														m.G0 = v8 + int32(48)
														return v168
													}
												}
											} else {
												if v65&int32(1) == int32(0) {
													v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
													v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
													if v88 == int32(1) {
														v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
														if v87&int32(1) != 0 {
															v94 = int32(1)
															v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v95 == v91&int32(255) {
																v168 = v94
															} else {
																if v91&int32(1) != 0 {
																	v168 = int32(0)
																} else {
																	v168 = v94
																}
															}
														} else {
															v101 = int32(1)
															if v91&v101 != 0 {
																v168 = int32(0)
															} else {
																v168 = v101
															}
														}
														m.G0 = v8 + int32(48)
														return v168
													} else {
														if v87&int32(1) != 0 {
															v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v106 == int32(0) {
																v168 = int32(0)
															} else {
																v168 = int32(1)
															}
															m.G0 = v8 + int32(48)
															return v168
														} else {
															v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
															v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
															v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
															v114 = F_FunctionCall2Coll(m, l0+int32(212), v111, v112, v113)
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return int32(0)
															} else {
																if v114 == int32(0) {
																	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
																	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
																	if v119 == int32(0) {
																		v122 = int32(1)
																		v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
																		if v118&v122 == int32(0) {
																			v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																			if v128 == v123&int32(255) {
																				v168 = v122
																			} else {
																				if v123&int32(1) != 0 {
																					v168 = v122
																				} else {
																					v168 = int32(0)
																				}
																			}
																		} else {
																			if v123&int32(1) == int32(0) {
																				v168 = int32(0)
																			} else {
																				v168 = v122
																			}
																		}
																	} else {
																		v138 = int32(1)
																		if v118&v138 != 0 {
																			v168 = v138
																		} else {
																			v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																			if v141 != 0 {
																				v168 = int32(0)
																			} else {
																				v168 = v138
																			}
																		}
																	}
																} else {
																	if v114 < int32(0) {
																		v168 = int32(0)
																	} else {
																		v168 = int32(1)
																	}
																}
																m.G0 = v8 + int32(48)
																return v168
															}
														}
													}
												} else {
													v168 = v32
													m.G0 = v8 + int32(48)
													return v168
												}
											}
										} else {
											if v61&int32(1) != 0 {
												v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
												v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
												if v88 == int32(1) {
													v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
													if v87&int32(1) != 0 {
														v94 = int32(1)
														v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
														if v95 == v91&int32(255) {
															v168 = v94
														} else {
															if v91&int32(1) != 0 {
																v168 = int32(0)
															} else {
																v168 = v94
															}
														}
													} else {
														v101 = int32(1)
														if v91&v101 != 0 {
															v168 = int32(0)
														} else {
															v168 = v101
														}
													}
													m.G0 = v8 + int32(48)
													return v168
												} else {
													if v87&int32(1) != 0 {
														v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
														if v106 == int32(0) {
															v168 = int32(0)
														} else {
															v168 = int32(1)
														}
														m.G0 = v8 + int32(48)
														return v168
													} else {
														v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
														v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
														v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
														v114 = F_FunctionCall2Coll(m, l0+int32(212), v111, v112, v113)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															if v114 == int32(0) {
																v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
																v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
																if v119 == int32(0) {
																	v122 = int32(1)
																	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
																	if v118&v122 == int32(0) {
																		v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																		if v128 == v123&int32(255) {
																			v168 = v122
																		} else {
																			if v123&int32(1) != 0 {
																				v168 = v122
																			} else {
																				v168 = int32(0)
																			}
																		}
																	} else {
																		if v123&int32(1) == int32(0) {
																			v168 = int32(0)
																		} else {
																			v168 = v122
																		}
																	}
																} else {
																	v138 = int32(1)
																	if v118&v138 != 0 {
																		v168 = v138
																	} else {
																		v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																		if v141 != 0 {
																			v168 = int32(0)
																		} else {
																			v168 = v138
																		}
																	}
																}
															} else {
																if v114 < int32(0) {
																	v168 = int32(0)
																} else {
																	v168 = int32(1)
																}
															}
															m.G0 = v8 + int32(48)
															return v168
														}
													}
												}
											} else {
												v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
												if v82 != 0 {
													v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
													v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
													if v88 == int32(1) {
														v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
														if v87&int32(1) != 0 {
															v94 = int32(1)
															v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v95 == v91&int32(255) {
																v168 = v94
															} else {
																if v91&int32(1) != 0 {
																	v168 = int32(0)
																} else {
																	v168 = v94
																}
															}
														} else {
															v101 = int32(1)
															if v91&v101 != 0 {
																v168 = int32(0)
															} else {
																v168 = v101
															}
														}
														m.G0 = v8 + int32(48)
														return v168
													} else {
														if v87&int32(1) != 0 {
															v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
															if v106 == int32(0) {
																v168 = int32(0)
															} else {
																v168 = int32(1)
															}
															m.G0 = v8 + int32(48)
															return v168
														} else {
															v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
															v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
															v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
															v114 = F_FunctionCall2Coll(m, l0+int32(212), v111, v112, v113)
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return int32(0)
															} else {
																if v114 == int32(0) {
																	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
																	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
																	if v119 == int32(0) {
																		v122 = int32(1)
																		v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
																		if v118&v122 == int32(0) {
																			v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																			if v128 == v123&int32(255) {
																				v168 = v122
																			} else {
																				if v123&int32(1) != 0 {
																					v168 = v122
																				} else {
																					v168 = int32(0)
																				}
																			}
																		} else {
																			if v123&int32(1) == int32(0) {
																				v168 = int32(0)
																			} else {
																				v168 = v122
																			}
																		}
																	} else {
																		v138 = int32(1)
																		if v118&v138 != 0 {
																			v168 = v138
																		} else {
																			v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																			if v141 != 0 {
																				v168 = int32(0)
																			} else {
																				v168 = v138
																			}
																		}
																	}
																} else {
																	if v114 < int32(0) {
																		v168 = int32(0)
																	} else {
																		v168 = int32(1)
																	}
																}
																m.G0 = v8 + int32(48)
																return v168
															}
														}
													}
												} else {
													v168 = v32
													m.G0 = v8 + int32(48)
													return v168
												}
											}
										}
									} else {
										if int32(0) < v57 {
											v168 = v32
											m.G0 = v8 + int32(48)
											return v168
										} else {
											v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
											v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+36)))
											if v88 == int32(1) {
												v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
												if v87&int32(1) != 0 {
													v94 = int32(1)
													v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
													if v95 == v91&int32(255) {
														v168 = v94
													} else {
														if v91&int32(1) != 0 {
															v168 = int32(0)
														} else {
															v168 = v94
														}
													}
												} else {
													v101 = int32(1)
													if v91&v101 != 0 {
														v168 = int32(0)
													} else {
														v168 = v101
													}
												}
												m.G0 = v8 + int32(48)
												return v168
											} else {
												if v87&int32(1) != 0 {
													v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
													if v106 == int32(0) {
														v168 = int32(0)
													} else {
														v168 = int32(1)
													}
													m.G0 = v8 + int32(48)
													return v168
												} else {
													v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
													v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
													v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
													v114 = F_FunctionCall2Coll(m, l0+int32(212), v111, v112, v113)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														if v114 == int32(0) {
															v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+17)))
															v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+37)))
															if v119 == int32(0) {
																v122 = int32(1)
																v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+38)))
																if v118&v122 == int32(0) {
																	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																	if v128 == v123&int32(255) {
																		v168 = v122
																	} else {
																		if v123&int32(1) != 0 {
																			v168 = v122
																		} else {
																			v168 = int32(0)
																		}
																	}
																} else {
																	if v123&int32(1) == int32(0) {
																		v168 = int32(0)
																	} else {
																		v168 = v122
																	}
																}
															} else {
																v138 = int32(1)
																if v118&v138 != 0 {
																	v168 = v138
																} else {
																	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+18)))
																	if v141 != 0 {
																		v168 = int32(0)
																	} else {
																		v168 = v138
																	}
																}
															}
														} else {
															if v114 < int32(0) {
																v168 = int32(0)
															} else {
																v168 = int32(1)
															}
														}
														m.G0 = v8 + int32(48)
														return v168
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
		v152 = m.ExcPending
		if v152 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(325375), int32(0))
			mBase = m.M
			v156 = m.ExcPending
			if v156 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493991), int32(2661), int32(310957))
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
	var v65 int32
	_ = v65
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
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v14 != v15 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v19
	v22 = l0 + int32(36)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_qsort_arg(m, v22+v23<<(uint(int32(3))%32), v14, int32(4), int32(21), v12+int32(8))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v12 + int32(16)
	return
L4:
	;
	return
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v34 < int32(2) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v91
	goto L3
L7:
	;
	v91 = int32(1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v38 = int32(1)
	v39 = v23 << (uint(v38) % 32)
	v45 = v38
	v46 = int32(1)
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v55 = int32(2)
	v57 = v22 + (v46+(v39-v38))<<(uint(v55)%32)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v62 = v22 + (v46+v39)<<(uint(v55)%32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v64 = F_FunctionCall2Coll(m, v52, v53, v58, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	v91 = v84
	goto L6
L12:
	;
	v86 = v46 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v86 < v87 {
		v45 = v84
		v46 = v86
		goto L10
	} else {
		goto L19
	}
L13:
	;
	if v64 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v72 = F_FunctionCall2Coll(m, v68, v69, v70, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v22+(v45+v39)<<(uint(int32(2))%32)))) = v80
	v84 = v45 + int32(1)
	goto L12
L17:
	;
	if v72 == int32(0) {
		v84 = v45
		goto L12
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L11
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
			F_errmsg_internal(m, int32(480468), v8)
			mBase = m.M
			v137 = m.ExcPending
			if v137 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(492485), int32(1030), int32(400622))
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
			F_errmsg_internal(m, int32(480468), v8)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(492485), int32(1084), int32(401492))
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
									F_errmsg_internal(m, int32(370267), v9)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(493991), int32(1776), int32(398753))
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
								F_errmsg_internal(m, int32(370267), v9)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493991), int32(1776), int32(398753))
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
					F_errmsg_internal(m, int32(325375), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493991), int32(1137), int32(109884))
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v109 int32
	_ = v109
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
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
				v31 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+62)) = uint8(v31)
				v33 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v7)+60)) = uint16(v33)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v33
				*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v33
				*(*int32)(unsafe.Add(mBase, uint32(v7)+51)) = v33
				v47 = F_make_range(m, l0, v5+int32(-8), v5+int32(-16), v31, v33)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v204 = v47
					m.G0 = v7 - int32(-64)
					return v204
				}
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
				if v28 != 0 {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+62)) = uint8(v31)
					v33 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v7)+60)) = uint16(v33)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(v7)+51)) = v33
					v47 = F_make_range(m, l0, v5+int32(-8), v5+int32(-16), v31, v33)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v204 = v47
						m.G0 = v7 - int32(-64)
						return v204
					}
				} else {
					v29 = F_range_overlaps_internal(m, l0, l1, l2)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						if v29 != 0 {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+36)))
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+44)))
							if v50 == int32(1) {
								v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+46)))
								if v49&int32(1) != 0 {
									v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
									if v56 == v53&int32(255) {
										v124 = v5 + int32(-24)
									} else {
										if v53&int32(1) == int32(0) {
											v124 = v5 + int32(-24)
										} else {
											v124 = v5 + int32(-32)
										}
									}
								} else {
									if v53&int32(1) == int32(0) {
										v124 = v5 + int32(-24)
									} else {
										v124 = v5 + int32(-32)
									}
								}
								v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
								v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)))
								if v126 == int32(1) {
									v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
									if v125&int32(1) != 0 {
										v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
										if v129 == v132 {
											v196 = v5 + int32(-40)
										} else {
											if v129&int32(1) != 0 {
												v196 = v5 + int32(-40)
											} else {
												v196 = v5 + int32(-48)
											}
										}
									} else {
										if v129&int32(1) != 0 {
											v196 = v5 + int32(-40)
										} else {
											v196 = v5 + int32(-48)
										}
									}
									v197 = int32(0)
									v199 = F_make_range(m, l0, v124, v196, v197, v197)
									mBase = m.M
									v200 = m.ExcPending
									if v200 != 0 {
										return int32(0)
									} else {
										v204 = v199
										m.G0 = v7 - int32(-64)
										return v204
									}
								} else {
									if v125&int32(1) != 0 {
										v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
										if v144 == int32(0) {
											v196 = v5 + int32(-40)
										} else {
											v196 = v5 + int32(-48)
										}
										v197 = int32(0)
										v199 = F_make_range(m, l0, v124, v196, v197, v197)
										mBase = m.M
										v200 = m.ExcPending
										if v200 != 0 {
											return int32(0)
										} else {
											v204 = v199
											m.G0 = v7 - int32(-64)
											return v204
										}
									} else {
										v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
										v152 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
										v153 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
										v154 = F_FunctionCall2Coll(m, l0+int32(212), v151, v152, v153)
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return int32(0)
										} else {
											if v154 == int32(0) {
												v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)))
												v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+29)))
												if v159 == int32(0) {
													v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
													if v158&int32(1) == int32(0) {
														v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
														if v162 == v167 {
															v196 = v5 + int32(-40)
														} else {
															if v162&int32(1) == int32(0) {
																v196 = v5 + int32(-40)
															} else {
																v196 = v5 + int32(-48)
															}
														}
													} else {
														if v162&int32(1) == int32(0) {
															v196 = v5 + int32(-40)
														} else {
															v196 = v5 + int32(-48)
														}
													}
												} else {
													if v158&int32(1) != 0 {
														v196 = v5 + int32(-40)
													} else {
														v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
														if v183 != 0 {
															v196 = v5 + int32(-40)
														} else {
															v196 = v5 + int32(-48)
														}
													}
												}
											} else {
												if v154 <= int32(0) {
													v196 = v5 + int32(-40)
												} else {
													v196 = v5 + int32(-48)
												}
											}
											v197 = int32(0)
											v199 = F_make_range(m, l0, v124, v196, v197, v197)
											mBase = m.M
											v200 = m.ExcPending
											if v200 != 0 {
												return int32(0)
											} else {
												v204 = v199
												m.G0 = v7 - int32(-64)
												return v204
											}
										}
									}
								}
							} else {
								if v49&int32(1) != 0 {
									v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
									if v74 != 0 {
										v124 = v5 + int32(-24)
									} else {
										v124 = v5 + int32(-32)
									}
									v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
									v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)))
									if v126 == int32(1) {
										v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
										if v125&int32(1) != 0 {
											v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
											if v129 == v132 {
												v196 = v5 + int32(-40)
											} else {
												if v129&int32(1) != 0 {
													v196 = v5 + int32(-40)
												} else {
													v196 = v5 + int32(-48)
												}
											}
										} else {
											if v129&int32(1) != 0 {
												v196 = v5 + int32(-40)
											} else {
												v196 = v5 + int32(-48)
											}
										}
										v197 = int32(0)
										v199 = F_make_range(m, l0, v124, v196, v197, v197)
										mBase = m.M
										v200 = m.ExcPending
										if v200 != 0 {
											return int32(0)
										} else {
											v204 = v199
											m.G0 = v7 - int32(-64)
											return v204
										}
									} else {
										if v125&int32(1) != 0 {
											v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
											if v144 == int32(0) {
												v196 = v5 + int32(-40)
											} else {
												v196 = v5 + int32(-48)
											}
											v197 = int32(0)
											v199 = F_make_range(m, l0, v124, v196, v197, v197)
											mBase = m.M
											v200 = m.ExcPending
											if v200 != 0 {
												return int32(0)
											} else {
												v204 = v199
												m.G0 = v7 - int32(-64)
												return v204
											}
										} else {
											v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v152 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
											v153 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
											v154 = F_FunctionCall2Coll(m, l0+int32(212), v151, v152, v153)
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int32(0)
											} else {
												if v154 == int32(0) {
													v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)))
													v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+29)))
													if v159 == int32(0) {
														v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
														if v158&int32(1) == int32(0) {
															v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
															if v162 == v167 {
																v196 = v5 + int32(-40)
															} else {
																if v162&int32(1) == int32(0) {
																	v196 = v5 + int32(-40)
																} else {
																	v196 = v5 + int32(-48)
																}
															}
														} else {
															if v162&int32(1) == int32(0) {
																v196 = v5 + int32(-40)
															} else {
																v196 = v5 + int32(-48)
															}
														}
													} else {
														if v158&int32(1) != 0 {
															v196 = v5 + int32(-40)
														} else {
															v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
															if v183 != 0 {
																v196 = v5 + int32(-40)
															} else {
																v196 = v5 + int32(-48)
															}
														}
													}
												} else {
													if v154 <= int32(0) {
														v196 = v5 + int32(-40)
													} else {
														v196 = v5 + int32(-48)
													}
												}
												v197 = int32(0)
												v199 = F_make_range(m, l0, v124, v196, v197, v197)
												mBase = m.M
												v200 = m.ExcPending
												if v200 != 0 {
													return int32(0)
												} else {
													v204 = v199
													m.G0 = v7 - int32(-64)
													return v204
												}
											}
										}
									}
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
									v82 = F_FunctionCall2Coll(m, l0+int32(212), v79, v80, v81)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										if v82 == int32(0) {
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+37)))
											v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+45)))
											if v87 == int32(0) {
												v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+46)))
												if v86&int32(1) == int32(0) {
													v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
													if v95 == v90&int32(255) {
														v124 = v5 + int32(-24)
													} else {
														if v90&int32(1) != 0 {
															v124 = v5 + int32(-24)
														} else {
															v124 = v5 + int32(-32)
														}
													}
												} else {
													if v90&int32(1) != 0 {
														v124 = v5 + int32(-24)
													} else {
														v124 = v5 + int32(-32)
													}
												}
											} else {
												if v86&int32(1) != 0 {
													v124 = v5 + int32(-24)
												} else {
													v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
													if v109 == int32(0) {
														v124 = v5 + int32(-24)
													} else {
														v124 = v5 + int32(-32)
													}
												}
											}
										} else {
											if int32(0) <= v82 {
												v124 = v5 + int32(-24)
											} else {
												v124 = v5 + int32(-32)
											}
										}
										v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
										v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)))
										if v126 == int32(1) {
											v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
											if v125&int32(1) != 0 {
												v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
												if v129 == v132 {
													v196 = v5 + int32(-40)
												} else {
													if v129&int32(1) != 0 {
														v196 = v5 + int32(-40)
													} else {
														v196 = v5 + int32(-48)
													}
												}
											} else {
												if v129&int32(1) != 0 {
													v196 = v5 + int32(-40)
												} else {
													v196 = v5 + int32(-48)
												}
											}
											v197 = int32(0)
											v199 = F_make_range(m, l0, v124, v196, v197, v197)
											mBase = m.M
											v200 = m.ExcPending
											if v200 != 0 {
												return int32(0)
											} else {
												v204 = v199
												m.G0 = v7 - int32(-64)
												return v204
											}
										} else {
											if v125&int32(1) != 0 {
												v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
												if v144 == int32(0) {
													v196 = v5 + int32(-40)
												} else {
													v196 = v5 + int32(-48)
												}
												v197 = int32(0)
												v199 = F_make_range(m, l0, v124, v196, v197, v197)
												mBase = m.M
												v200 = m.ExcPending
												if v200 != 0 {
													return int32(0)
												} else {
													v204 = v199
													m.G0 = v7 - int32(-64)
													return v204
												}
											} else {
												v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
												v152 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
												v153 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
												v154 = F_FunctionCall2Coll(m, l0+int32(212), v151, v152, v153)
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return int32(0)
												} else {
													if v154 == int32(0) {
														v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)))
														v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+29)))
														if v159 == int32(0) {
															v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
															if v158&int32(1) == int32(0) {
																v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
																if v162 == v167 {
																	v196 = v5 + int32(-40)
																} else {
																	if v162&int32(1) == int32(0) {
																		v196 = v5 + int32(-40)
																	} else {
																		v196 = v5 + int32(-48)
																	}
																}
															} else {
																if v162&int32(1) == int32(0) {
																	v196 = v5 + int32(-40)
																} else {
																	v196 = v5 + int32(-48)
																}
															}
														} else {
															if v158&int32(1) != 0 {
																v196 = v5 + int32(-40)
															} else {
																v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
																if v183 != 0 {
																	v196 = v5 + int32(-40)
																} else {
																	v196 = v5 + int32(-48)
																}
															}
														}
													} else {
														if v154 <= int32(0) {
															v196 = v5 + int32(-40)
														} else {
															v196 = v5 + int32(-48)
														}
													}
													v197 = int32(0)
													v199 = F_make_range(m, l0, v124, v196, v197, v197)
													mBase = m.M
													v200 = m.ExcPending
													if v200 != 0 {
														return int32(0)
													} else {
														v204 = v199
														m.G0 = v7 - int32(-64)
														return v204
													}
												}
											}
										}
									}
								}
							}
						} else {
							v31 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v7)+62)) = uint8(v31)
							v33 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v7)+60)) = uint16(v33)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v33
							*(*int32)(unsafe.Add(mBase, uint32(v7)+51)) = v33
							v47 = F_make_range(m, l0, v5+int32(-8), v5+int32(-16), v31, v33)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v204 = v47
								m.G0 = v7 - int32(-64)
								return v204
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v12)>>(uint(int32(2))%32))-int32(1)))))
	goto L2
L1:
	;
	m.G0 = v10 + int32(48)
	return v80
L2:
	;
	if v18&int32(1) != 0 {
		v80 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v21 == int32(0) {
		v80 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = v10 + int32(16)
	v29 = v25 | int32(8)
	F_range_deserialize(m, l0, l1, v25, v29, v10+int32(15))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v36 == int32(0) {
		v80 = v4
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v40 = v36
	v43 = v4
	goto L9
L8:
	;
	v80 = int32(1)
	goto L1
L9:
	;
	v48 = int32(base.Ui32(v40+v43) >> (uint(int32(1)) % 32))
	F_multirange_get_bounds(m, l0, l2, v48, v10+int32(40), v10+int32(32))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	v80 = int32(0)
	goto L1
L11:
	;
	v57 = F_range_cmp_bounds(m, l0, v29, v10+int32(40))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	if base.Ui32(v72) < base.Ui32(v71) {
		v40 = v71
		v43 = v72
		goto L9
	} else {
		goto L19
	}
L13:
	;
	if v57 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = v48
	v72 = v43
	goto L12
L15:
	;
	goto L16
L16:
	;
	v65 = F_range_cmp_bounds(m, l0, v10+int32(16), v10+int32(32))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v65 <= int32(0) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v71 = v40
	v72 = v48 + int32(1)
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
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
					v104 = v30
					m.G0 = v7 + int32(48)
					return v104
				} else {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
					if v32 != 0 {
						v104 = v30
						m.G0 = v7 + int32(48)
						return v104
					} else {
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)))
						if v34 == int32(1) {
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
							if v33&int32(1) != 0 {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
								if v40 == v37&int32(255) {
									v99 = int32(0)
								} else {
									v46 = int32(1)
									if v37&v46 != 0 {
										v49 = int32(-1)
									} else {
										v49 = v46
									}
									v99 = v49
								}
							} else {
								v51 = int32(1)
								if v37&v51 != 0 {
									v54 = int32(-1)
								} else {
									v54 = v51
								}
								v99 = v54
							}
							v104 = base.B2i32(v99 <= int32(0))
							m.G0 = v7 + int32(48)
							return v104
						} else {
							if v33&int32(1) != 0 {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
								if v59 != 0 {
									v60 = int32(1)
								} else {
									v60 = int32(-1)
								}
								v99 = v60
								v104 = base.B2i32(v99 <= int32(0))
								m.G0 = v7 + int32(48)
								return v104
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
										v99 = v66
									} else {
										v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)))
										v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+29)))
										if v69 == int32(0) {
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
											if v68&int32(1) == int32(0) {
												v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
												if v77 == v72&int32(255) {
													v99 = int32(0)
												} else {
													v82 = int32(1)
													if v72&v82 != 0 {
														v86 = v82
													} else {
														v86 = int32(-1)
													}
													v99 = v86
												}
											} else {
												v87 = int32(1)
												if v72&v87 != 0 {
													v91 = v87
												} else {
													v91 = int32(-1)
												}
												v99 = v91
											}
										} else {
											if v68&int32(1) != 0 {
												v99 = int32(0)
											} else {
												v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
												if v97 != 0 {
													v98 = int32(-1)
												} else {
													v98 = int32(1)
												}
												v99 = v98
											}
										}
									}
									v104 = base.B2i32(v99 <= int32(0))
									m.G0 = v7 + int32(48)
									return v104
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
		v112 = m.ExcPending
		if v112 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(325375), int32(0))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493991), int32(900), int32(310888))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
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
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v34 = F_range_union_internal(m, v32, v12, v17, int32(1))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
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
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(370267), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493991), int32(1776), int32(398753))
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
							v34 = F_range_union_internal(m, v32, v12, v17, int32(1))
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
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(370267), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493991), int32(1776), int32(398753))
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
						v34 = F_range_union_internal(m, v32, v12, v17, int32(1))
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
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v171 int32
	_ = v171
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
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
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
									if v51 == v46&int32(255) {
										v144 = v50
									} else {
										if v46&int32(1) != 0 {
											v144 = v9 + int32(40)
										} else {
											v144 = v50
										}
									}
								} else {
									if v46&int32(1) != 0 {
										v144 = v9 + int32(40)
									} else {
										v144 = v9 + int32(32)
									}
								}
								v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
								if v147 == int32(1) {
									v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
									if v146&int32(1) != 0 {
										v154 = v9 + int32(16)
										v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
										if v155 == v150&int32(255) {
											v220 = v154
										} else {
											if v150&int32(1) == int32(0) {
												v220 = v9 + int32(24)
											} else {
												v220 = v154
											}
										}
									} else {
										if v150&int32(1) == int32(0) {
											v220 = v9 + int32(24)
										} else {
											v220 = v9 + int32(16)
										}
									}
									v223 = int32(0)
									v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
									if v146&int32(1) != 0 {
										v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
										if v171 != 0 {
											v220 = v9 + int32(24)
										} else {
											v220 = v9 + int32(16)
										}
										v223 = int32(0)
										v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
										v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
										v177 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
										v178 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
										v179 = F_FunctionCall2Coll(m, l0+int32(212), v176, v177, v178)
										mBase = m.M
										v180 = m.ExcPending
										if v180 != 0 {
											return int32(0)
										} else {
											if v179 == int32(0) {
												v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
												v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
												if v184 == int32(0) {
													v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
													if v183&int32(1) == int32(0) {
														v193 = v9 + int32(16)
														v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
														if v187 == v194 {
															v220 = v193
														} else {
															if v187&int32(1) == int32(0) {
																v220 = v193
															} else {
																v220 = v9 + int32(24)
															}
														}
													} else {
														if v187&int32(1) != 0 {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
													}
												} else {
													v205 = v9 + int32(16)
													if v183&int32(1) != 0 {
														v220 = v205
													} else {
														v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
														if v208 == int32(0) {
															v220 = v9 + int32(24)
														} else {
															v220 = v205
														}
													}
												}
											} else {
												if int32(0) < v179 {
													v220 = v9 + int32(24)
												} else {
													v220 = v9 + int32(16)
												}
											}
											v223 = int32(0)
											v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
									v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
									if v63 == int32(0) {
										v144 = v9 + int32(40)
									} else {
										v144 = v9 + int32(32)
									}
									v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
									v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
									if v147 == int32(1) {
										v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
										if v146&int32(1) != 0 {
											v154 = v9 + int32(16)
											v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
											if v155 == v150&int32(255) {
												v220 = v154
											} else {
												if v150&int32(1) == int32(0) {
													v220 = v9 + int32(24)
												} else {
													v220 = v154
												}
											}
										} else {
											if v150&int32(1) == int32(0) {
												v220 = v9 + int32(24)
											} else {
												v220 = v9 + int32(16)
											}
										}
										v223 = int32(0)
										v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
										if v146&int32(1) != 0 {
											v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
											if v171 != 0 {
												v220 = v9 + int32(24)
											} else {
												v220 = v9 + int32(16)
											}
											v223 = int32(0)
											v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
											v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v177 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
											v178 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
											v179 = F_FunctionCall2Coll(m, l0+int32(212), v176, v177, v178)
											mBase = m.M
											v180 = m.ExcPending
											if v180 != 0 {
												return int32(0)
											} else {
												if v179 == int32(0) {
													v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
													v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
													if v184 == int32(0) {
														v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
														if v183&int32(1) == int32(0) {
															v193 = v9 + int32(16)
															v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
															if v187 == v194 {
																v220 = v193
															} else {
																if v187&int32(1) == int32(0) {
																	v220 = v193
																} else {
																	v220 = v9 + int32(24)
																}
															}
														} else {
															if v187&int32(1) != 0 {
																v220 = v9 + int32(24)
															} else {
																v220 = v9 + int32(16)
															}
														}
													} else {
														v205 = v9 + int32(16)
														if v183&int32(1) != 0 {
															v220 = v205
														} else {
															v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
															if v208 == int32(0) {
																v220 = v9 + int32(24)
															} else {
																v220 = v205
															}
														}
													}
												} else {
													if int32(0) < v179 {
														v220 = v9 + int32(24)
													} else {
														v220 = v9 + int32(16)
													}
												}
												v223 = int32(0)
												v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
									v73 = F_FunctionCall2Coll(m, l0+int32(212), v70, v71, v72)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										if v73 == int32(0) {
											v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+37)))
											v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+45)))
											if v78 == int32(0) {
												v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+46)))
												if v77&int32(1) == int32(0) {
													v87 = v9 + int32(32)
													v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
													if v88 == v81&int32(255) {
														v144 = v87
													} else {
														if v81&int32(1) != 0 {
															v144 = v87
														} else {
															v144 = v9 + int32(40)
														}
													}
												} else {
													if v81&int32(1) == int32(0) {
														v144 = v9 + int32(40)
													} else {
														v144 = v9 + int32(32)
													}
												}
											} else {
												v101 = v9 + int32(32)
												if v77&int32(1) != 0 {
													v144 = v101
												} else {
													v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
													if v104 != 0 {
														v144 = v9 + int32(40)
													} else {
														v144 = v101
													}
												}
											}
										} else {
											if v73 < int32(0) {
												v144 = v9 + int32(40)
											} else {
												v144 = v9 + int32(32)
											}
										}
										v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
										v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
										if v147 == int32(1) {
											v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
											if v146&int32(1) != 0 {
												v154 = v9 + int32(16)
												v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
												if v155 == v150&int32(255) {
													v220 = v154
												} else {
													if v150&int32(1) == int32(0) {
														v220 = v9 + int32(24)
													} else {
														v220 = v154
													}
												}
											} else {
												if v150&int32(1) == int32(0) {
													v220 = v9 + int32(24)
												} else {
													v220 = v9 + int32(16)
												}
											}
											v223 = int32(0)
											v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
											if v146&int32(1) != 0 {
												v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
												if v171 != 0 {
													v220 = v9 + int32(24)
												} else {
													v220 = v9 + int32(16)
												}
												v223 = int32(0)
												v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
												v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
												v177 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
												v178 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
												v179 = F_FunctionCall2Coll(m, l0+int32(212), v176, v177, v178)
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													if v179 == int32(0) {
														v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
														v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
														if v184 == int32(0) {
															v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
															if v183&int32(1) == int32(0) {
																v193 = v9 + int32(16)
																v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																if v187 == v194 {
																	v220 = v193
																} else {
																	if v187&int32(1) == int32(0) {
																		v220 = v193
																	} else {
																		v220 = v9 + int32(24)
																	}
																}
															} else {
																if v187&int32(1) != 0 {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v9 + int32(16)
																}
															}
														} else {
															v205 = v9 + int32(16)
															if v183&int32(1) != 0 {
																v220 = v205
															} else {
																v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																if v208 == int32(0) {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v205
																}
															}
														}
													} else {
														if int32(0) < v179 {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
													}
													v223 = int32(0)
													v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
											if v51 == v46&int32(255) {
												v144 = v50
											} else {
												if v46&int32(1) != 0 {
													v144 = v9 + int32(40)
												} else {
													v144 = v50
												}
											}
										} else {
											if v46&int32(1) != 0 {
												v144 = v9 + int32(40)
											} else {
												v144 = v9 + int32(32)
											}
										}
										v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
										v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
										if v147 == int32(1) {
											v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
											if v146&int32(1) != 0 {
												v154 = v9 + int32(16)
												v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
												if v155 == v150&int32(255) {
													v220 = v154
												} else {
													if v150&int32(1) == int32(0) {
														v220 = v9 + int32(24)
													} else {
														v220 = v154
													}
												}
											} else {
												if v150&int32(1) == int32(0) {
													v220 = v9 + int32(24)
												} else {
													v220 = v9 + int32(16)
												}
											}
											v223 = int32(0)
											v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
											if v146&int32(1) != 0 {
												v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
												if v171 != 0 {
													v220 = v9 + int32(24)
												} else {
													v220 = v9 + int32(16)
												}
												v223 = int32(0)
												v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
												v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
												v177 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
												v178 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
												v179 = F_FunctionCall2Coll(m, l0+int32(212), v176, v177, v178)
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													if v179 == int32(0) {
														v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
														v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
														if v184 == int32(0) {
															v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
															if v183&int32(1) == int32(0) {
																v193 = v9 + int32(16)
																v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																if v187 == v194 {
																	v220 = v193
																} else {
																	if v187&int32(1) == int32(0) {
																		v220 = v193
																	} else {
																		v220 = v9 + int32(24)
																	}
																}
															} else {
																if v187&int32(1) != 0 {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v9 + int32(16)
																}
															}
														} else {
															v205 = v9 + int32(16)
															if v183&int32(1) != 0 {
																v220 = v205
															} else {
																v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																if v208 == int32(0) {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v205
																}
															}
														}
													} else {
														if int32(0) < v179 {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
													}
													v223 = int32(0)
													v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
											v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
											if v63 == int32(0) {
												v144 = v9 + int32(40)
											} else {
												v144 = v9 + int32(32)
											}
											v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
											v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
											if v147 == int32(1) {
												v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
												if v146&int32(1) != 0 {
													v154 = v9 + int32(16)
													v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
													if v155 == v150&int32(255) {
														v220 = v154
													} else {
														if v150&int32(1) == int32(0) {
															v220 = v9 + int32(24)
														} else {
															v220 = v154
														}
													}
												} else {
													if v150&int32(1) == int32(0) {
														v220 = v9 + int32(24)
													} else {
														v220 = v9 + int32(16)
													}
												}
												v223 = int32(0)
												v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
												if v146&int32(1) != 0 {
													v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
													if v171 != 0 {
														v220 = v9 + int32(24)
													} else {
														v220 = v9 + int32(16)
													}
													v223 = int32(0)
													v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
													v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
													v177 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
													v178 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
													v179 = F_FunctionCall2Coll(m, l0+int32(212), v176, v177, v178)
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return int32(0)
													} else {
														if v179 == int32(0) {
															v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
															v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
															if v184 == int32(0) {
																v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
																if v183&int32(1) == int32(0) {
																	v193 = v9 + int32(16)
																	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																	if v187 == v194 {
																		v220 = v193
																	} else {
																		if v187&int32(1) == int32(0) {
																			v220 = v193
																		} else {
																			v220 = v9 + int32(24)
																		}
																	}
																} else {
																	if v187&int32(1) != 0 {
																		v220 = v9 + int32(24)
																	} else {
																		v220 = v9 + int32(16)
																	}
																}
															} else {
																v205 = v9 + int32(16)
																if v183&int32(1) != 0 {
																	v220 = v205
																} else {
																	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																	if v208 == int32(0) {
																		v220 = v9 + int32(24)
																	} else {
																		v220 = v205
																	}
																}
															}
														} else {
															if int32(0) < v179 {
																v220 = v9 + int32(24)
															} else {
																v220 = v9 + int32(16)
															}
														}
														v223 = int32(0)
														v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
											v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
											v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
											v73 = F_FunctionCall2Coll(m, l0+int32(212), v70, v71, v72)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												if v73 == int32(0) {
													v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+37)))
													v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+45)))
													if v78 == int32(0) {
														v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+46)))
														if v77&int32(1) == int32(0) {
															v87 = v9 + int32(32)
															v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
															if v88 == v81&int32(255) {
																v144 = v87
															} else {
																if v81&int32(1) != 0 {
																	v144 = v87
																} else {
																	v144 = v9 + int32(40)
																}
															}
														} else {
															if v81&int32(1) == int32(0) {
																v144 = v9 + int32(40)
															} else {
																v144 = v9 + int32(32)
															}
														}
													} else {
														v101 = v9 + int32(32)
														if v77&int32(1) != 0 {
															v144 = v101
														} else {
															v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
															if v104 != 0 {
																v144 = v9 + int32(40)
															} else {
																v144 = v101
															}
														}
													}
												} else {
													if v73 < int32(0) {
														v144 = v9 + int32(40)
													} else {
														v144 = v9 + int32(32)
													}
												}
												v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
												v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
												if v147 == int32(1) {
													v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
													if v146&int32(1) != 0 {
														v154 = v9 + int32(16)
														v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
														if v155 == v150&int32(255) {
															v220 = v154
														} else {
															if v150&int32(1) == int32(0) {
																v220 = v9 + int32(24)
															} else {
																v220 = v154
															}
														}
													} else {
														if v150&int32(1) == int32(0) {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
													}
													v223 = int32(0)
													v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
													if v146&int32(1) != 0 {
														v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
														if v171 != 0 {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
														v223 = int32(0)
														v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
														v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
														v177 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
														v178 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
														v179 = F_FunctionCall2Coll(m, l0+int32(212), v176, v177, v178)
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
															return int32(0)
														} else {
															if v179 == int32(0) {
																v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
																v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
																if v184 == int32(0) {
																	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
																	if v183&int32(1) == int32(0) {
																		v193 = v9 + int32(16)
																		v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																		if v187 == v194 {
																			v220 = v193
																		} else {
																			if v187&int32(1) == int32(0) {
																				v220 = v193
																			} else {
																				v220 = v9 + int32(24)
																			}
																		}
																	} else {
																		if v187&int32(1) != 0 {
																			v220 = v9 + int32(24)
																		} else {
																			v220 = v9 + int32(16)
																		}
																	}
																} else {
																	v205 = v9 + int32(16)
																	if v183&int32(1) != 0 {
																		v220 = v205
																	} else {
																		v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																		if v208 == int32(0) {
																			v220 = v9 + int32(24)
																		} else {
																			v220 = v205
																		}
																	}
																}
															} else {
																if int32(0) < v179 {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v9 + int32(16)
																}
															}
															v223 = int32(0)
															v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(130))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(114697), int32(0))
													mBase = m.M
													v132 = m.ExcPending
													if v132 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(493991), int32(1084), int32(311249))
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
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
													if v51 == v46&int32(255) {
														v144 = v50
													} else {
														if v46&int32(1) != 0 {
															v144 = v9 + int32(40)
														} else {
															v144 = v50
														}
													}
												} else {
													if v46&int32(1) != 0 {
														v144 = v9 + int32(40)
													} else {
														v144 = v9 + int32(32)
													}
												}
												v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
												v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
												if v147 == int32(1) {
													v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
													if v146&int32(1) != 0 {
														v154 = v9 + int32(16)
														v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
														if v155 == v150&int32(255) {
															v220 = v154
														} else {
															if v150&int32(1) == int32(0) {
																v220 = v9 + int32(24)
															} else {
																v220 = v154
															}
														}
													} else {
														if v150&int32(1) == int32(0) {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
													}
													v223 = int32(0)
													v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
													if v146&int32(1) != 0 {
														v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
														if v171 != 0 {
															v220 = v9 + int32(24)
														} else {
															v220 = v9 + int32(16)
														}
														v223 = int32(0)
														v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
														v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
														v177 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
														v178 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
														v179 = F_FunctionCall2Coll(m, l0+int32(212), v176, v177, v178)
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
															return int32(0)
														} else {
															if v179 == int32(0) {
																v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
																v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
																if v184 == int32(0) {
																	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
																	if v183&int32(1) == int32(0) {
																		v193 = v9 + int32(16)
																		v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																		if v187 == v194 {
																			v220 = v193
																		} else {
																			if v187&int32(1) == int32(0) {
																				v220 = v193
																			} else {
																				v220 = v9 + int32(24)
																			}
																		}
																	} else {
																		if v187&int32(1) != 0 {
																			v220 = v9 + int32(24)
																		} else {
																			v220 = v9 + int32(16)
																		}
																	}
																} else {
																	v205 = v9 + int32(16)
																	if v183&int32(1) != 0 {
																		v220 = v205
																	} else {
																		v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																		if v208 == int32(0) {
																			v220 = v9 + int32(24)
																		} else {
																			v220 = v205
																		}
																	}
																}
															} else {
																if int32(0) < v179 {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v9 + int32(16)
																}
															}
															v223 = int32(0)
															v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
													v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
													if v63 == int32(0) {
														v144 = v9 + int32(40)
													} else {
														v144 = v9 + int32(32)
													}
													v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
													v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
													if v147 == int32(1) {
														v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
														if v146&int32(1) != 0 {
															v154 = v9 + int32(16)
															v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
															if v155 == v150&int32(255) {
																v220 = v154
															} else {
																if v150&int32(1) == int32(0) {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v154
																}
															}
														} else {
															if v150&int32(1) == int32(0) {
																v220 = v9 + int32(24)
															} else {
																v220 = v9 + int32(16)
															}
														}
														v223 = int32(0)
														v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
														if v146&int32(1) != 0 {
															v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
															if v171 != 0 {
																v220 = v9 + int32(24)
															} else {
																v220 = v9 + int32(16)
															}
															v223 = int32(0)
															v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
															v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
															v177 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
															v178 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
															v179 = F_FunctionCall2Coll(m, l0+int32(212), v176, v177, v178)
															mBase = m.M
															v180 = m.ExcPending
															if v180 != 0 {
																return int32(0)
															} else {
																if v179 == int32(0) {
																	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
																	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
																	if v184 == int32(0) {
																		v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
																		if v183&int32(1) == int32(0) {
																			v193 = v9 + int32(16)
																			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																			if v187 == v194 {
																				v220 = v193
																			} else {
																				if v187&int32(1) == int32(0) {
																					v220 = v193
																				} else {
																					v220 = v9 + int32(24)
																				}
																			}
																		} else {
																			if v187&int32(1) != 0 {
																				v220 = v9 + int32(24)
																			} else {
																				v220 = v9 + int32(16)
																			}
																		}
																	} else {
																		v205 = v9 + int32(16)
																		if v183&int32(1) != 0 {
																			v220 = v205
																		} else {
																			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																			if v208 == int32(0) {
																				v220 = v9 + int32(24)
																			} else {
																				v220 = v205
																			}
																		}
																	}
																} else {
																	if int32(0) < v179 {
																		v220 = v9 + int32(24)
																	} else {
																		v220 = v9 + int32(16)
																	}
																}
																v223 = int32(0)
																v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
													v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
													v71 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
													v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
													v73 = F_FunctionCall2Coll(m, l0+int32(212), v70, v71, v72)
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int32(0)
													} else {
														if v73 == int32(0) {
															v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+37)))
															v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+45)))
															if v78 == int32(0) {
																v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+46)))
																if v77&int32(1) == int32(0) {
																	v87 = v9 + int32(32)
																	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
																	if v88 == v81&int32(255) {
																		v144 = v87
																	} else {
																		if v81&int32(1) != 0 {
																			v144 = v87
																		} else {
																			v144 = v9 + int32(40)
																		}
																	}
																} else {
																	if v81&int32(1) == int32(0) {
																		v144 = v9 + int32(40)
																	} else {
																		v144 = v9 + int32(32)
																	}
																}
															} else {
																v101 = v9 + int32(32)
																if v77&int32(1) != 0 {
																	v144 = v101
																} else {
																	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
																	if v104 != 0 {
																		v144 = v9 + int32(40)
																	} else {
																		v144 = v101
																	}
																}
															}
														} else {
															if v73 < int32(0) {
																v144 = v9 + int32(40)
															} else {
																v144 = v9 + int32(32)
															}
														}
														v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
														v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
														if v147 == int32(1) {
															v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
															if v146&int32(1) != 0 {
																v154 = v9 + int32(16)
																v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																if v155 == v150&int32(255) {
																	v220 = v154
																} else {
																	if v150&int32(1) == int32(0) {
																		v220 = v9 + int32(24)
																	} else {
																		v220 = v154
																	}
																}
															} else {
																if v150&int32(1) == int32(0) {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v9 + int32(16)
																}
															}
															v223 = int32(0)
															v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
															if v146&int32(1) != 0 {
																v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																if v171 != 0 {
																	v220 = v9 + int32(24)
																} else {
																	v220 = v9 + int32(16)
																}
																v223 = int32(0)
																v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
																v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
																v177 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
																v178 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
																v179 = F_FunctionCall2Coll(m, l0+int32(212), v176, v177, v178)
																mBase = m.M
																v180 = m.ExcPending
																if v180 != 0 {
																	return int32(0)
																} else {
																	if v179 == int32(0) {
																		v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
																		v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
																		if v184 == int32(0) {
																			v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+30)))
																			if v183&int32(1) == int32(0) {
																				v193 = v9 + int32(16)
																				v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																				if v187 == v194 {
																					v220 = v193
																				} else {
																					if v187&int32(1) == int32(0) {
																						v220 = v193
																					} else {
																						v220 = v9 + int32(24)
																					}
																				}
																			} else {
																				if v187&int32(1) != 0 {
																					v220 = v9 + int32(24)
																				} else {
																					v220 = v9 + int32(16)
																				}
																			}
																		} else {
																			v205 = v9 + int32(16)
																			if v183&int32(1) != 0 {
																				v220 = v205
																			} else {
																				v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
																				if v208 == int32(0) {
																					v220 = v9 + int32(24)
																				} else {
																					v220 = v205
																				}
																			}
																		}
																	} else {
																		if int32(0) < v179 {
																			v220 = v9 + int32(24)
																		} else {
																			v220 = v9 + int32(16)
																		}
																	}
																	v223 = int32(0)
																	v225 = F_make_range(m, l0, v144, v220, v223, v223)
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
		v112 = m.ExcPending
		if v112 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(325375), int32(0))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493991), int32(1068), int32(311249))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
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
	var v8 int32
	_ = v8
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
		v8 = int32(2)
		v11 = int32(1)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3+int32(base.Ui32(v7)>>(uint(v8)%32))-v11))))
		return int32(base.Ui32(v13)>>(uint(v8)%32)) & v11
	}
}
