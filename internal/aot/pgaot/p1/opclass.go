package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_opclass_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	F_DeconstructQualifiedName(m, l1, v8+int32(28), v8+int32(24))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
		if v18 != 0 {
			v20 = F_LookupExplicitNamespace(m, v18, l2)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if v20 == int32(0) {
					v38 = int32(0)
					if l2|v38 == int32(0) {
						v43 = F_SearchSysCache1(m, int32(2), l0)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								if v43 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
									F_errmsg_internal(m, int32(_a_F_get_opclass_oid_0), v8)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(202), int32(_a_F_get_opclass_oid_2))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										v54 = F_NameListToString(m, l1)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
											v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v54
											*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v56 + v57 + int32(4)
											F_errmsg(m, int32(_a_F_get_opclass_oid_3), v8+int32(16))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(207), int32(_a_F_get_opclass_oid_2))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
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
					} else {
						if v38 == int32(0) {
							v83 = int32(0)
							m.G0 = v8 + int32(32)
							return v83
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+22)))
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v76+v77)))
							F_ReleaseCatCache(m, v38)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v83 = v79
								m.G0 = v8 + int32(32)
								return v83
							}
						}
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
					v26 = F_SearchSysCache3(m, int32(13), l0, v25, v20)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v38 = v26
						if l2|v38 == int32(0) {
							v43 = F_SearchSysCache1(m, int32(2), l0)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									if v43 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
										F_errmsg_internal(m, int32(_a_F_get_opclass_oid_0), v8)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(202), int32(_a_F_get_opclass_oid_2))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										F_errcode(m, int32(67137668))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											v54 = F_NameListToString(m, l1)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return int32(0)
											} else {
												v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
												v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
												*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v54
												*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v56 + v57 + int32(4)
												F_errmsg(m, int32(_a_F_get_opclass_oid_3), v8+int32(16))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(207), int32(_a_F_get_opclass_oid_2))
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
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
						} else {
							if v38 == int32(0) {
								v83 = int32(0)
								m.G0 = v8 + int32(32)
								return v83
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+22)))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v76+v77)))
								F_ReleaseCatCache(m, v38)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v83 = v79
									m.G0 = v8 + int32(32)
									return v83
								}
							}
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
			v30 = F_OpclassnameGetOpcid(m, l0, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				if v30 == int32(0) {
					v38 = int32(0)
					if l2|v38 == int32(0) {
						v43 = F_SearchSysCache1(m, int32(2), l0)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								if v43 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
									F_errmsg_internal(m, int32(_a_F_get_opclass_oid_0), v8)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(202), int32(_a_F_get_opclass_oid_2))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										v54 = F_NameListToString(m, l1)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
											v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v54
											*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v56 + v57 + int32(4)
											F_errmsg(m, int32(_a_F_get_opclass_oid_3), v8+int32(16))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(207), int32(_a_F_get_opclass_oid_2))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
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
					} else {
						if v38 == int32(0) {
							v83 = int32(0)
							m.G0 = v8 + int32(32)
							return v83
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+22)))
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v76+v77)))
							F_ReleaseCatCache(m, v38)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v83 = v79
								m.G0 = v8 + int32(32)
								return v83
							}
						}
					}
				} else {
					v35 = F_SearchSysCache1(m, int32(14), v30)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v38 = v35
						if l2|v38 == int32(0) {
							v43 = F_SearchSysCache1(m, int32(2), l0)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									if v43 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
										F_errmsg_internal(m, int32(_a_F_get_opclass_oid_0), v8)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(202), int32(_a_F_get_opclass_oid_2))
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										F_errcode(m, int32(67137668))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											v54 = F_NameListToString(m, l1)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return int32(0)
											} else {
												v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
												v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
												*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v54
												*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v56 + v57 + int32(4)
												F_errmsg(m, int32(_a_F_get_opclass_oid_3), v8+int32(16))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(207), int32(_a_F_get_opclass_oid_2))
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
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
						} else {
							if v38 == int32(0) {
								v83 = int32(0)
								m.G0 = v8 + int32(32)
								return v83
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+22)))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v76+v77)))
								F_ReleaseCatCache(m, v38)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v83 = v79
									m.G0 = v8 + int32(32)
									return v83
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_opclass_for_family_datatype(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v54 int32
	_ = v54
	v4 = int32(0)
	v12 = F_SearchSysCacheList(m, int32(13), int32(1), l0, v4, v4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if int32(0) < v16 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = v4
	goto L6
L4:
	;
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v12)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(48)+v25<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+56))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
	v34 = v32 + v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+80))
	if v35 != l1 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v44 = v25 + int32(1)
	if v44 != v16 {
		v25 = v44
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+84))
	if v37 != l2 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	F_ReleaseCatCacheList(m, v12)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	return v39
L12:
	;
	goto L7
L13:
	;
	return int32(0)
}
