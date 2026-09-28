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
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
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
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
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
					v41 = int32(0)
					if l2|v41 == int32(0) {
						v47 = F_SearchSysCache1(m, int32(2), base.I64_extend_i32_u(l0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								if v47 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
									F_errmsg_internal(m, int32(_a_F_get_opclass_oid_0), v8)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(202), int32(_a_F_get_opclass_oid_2))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
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
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v58 = F_NameListToString(m, l1)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
											v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v58
											*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v60 + v61 + int32(4)
											F_errmsg(m, int32(_a_F_get_opclass_oid_3), v8+int32(16))
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(207), int32(_a_F_get_opclass_oid_2))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
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
						if v41 == int32(0) {
							v87 = int32(0)
							m.G0 = v8 + int32(32)
							return v87
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+22)))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v80+v81)))
							F_ReleaseCatCache(m, v41)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								v87 = v83
								m.G0 = v8 + int32(32)
								return v87
							}
						}
					}
				} else {
					v26 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+24)))
					v28 = F_SearchSysCache3(m, int32(13), base.I64_extend_i32_u(l0), v26, base.I64_extend_i32_u(v20))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v41 = v28
						if l2|v41 == int32(0) {
							v47 = F_SearchSysCache1(m, int32(2), base.I64_extend_i32_u(l0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									if v47 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
										F_errmsg_internal(m, int32(_a_F_get_opclass_oid_0), v8)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(202), int32(_a_F_get_opclass_oid_2))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
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
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = F_NameListToString(m, l1)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
												v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
												*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v58
												*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v60 + v61 + int32(4)
												F_errmsg(m, int32(_a_F_get_opclass_oid_3), v8+int32(16))
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(207), int32(_a_F_get_opclass_oid_2))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
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
							if v41 == int32(0) {
								v87 = int32(0)
								m.G0 = v8 + int32(32)
								return v87
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+22)))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v80+v81)))
								F_ReleaseCatCache(m, v41)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									v87 = v83
									m.G0 = v8 + int32(32)
									return v87
								}
							}
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
			v32 = F_OpclassnameGetOpcid(m, l0, v31)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				if v32 == int32(0) {
					v41 = int32(0)
					if l2|v41 == int32(0) {
						v47 = F_SearchSysCache1(m, int32(2), base.I64_extend_i32_u(l0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								if v47 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
									F_errmsg_internal(m, int32(_a_F_get_opclass_oid_0), v8)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(202), int32(_a_F_get_opclass_oid_2))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
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
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v58 = F_NameListToString(m, l1)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v60 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
											v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v58
											*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v60 + v61 + int32(4)
											F_errmsg(m, int32(_a_F_get_opclass_oid_3), v8+int32(16))
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(207), int32(_a_F_get_opclass_oid_2))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
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
						if v41 == int32(0) {
							v87 = int32(0)
							m.G0 = v8 + int32(32)
							return v87
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+22)))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v80+v81)))
							F_ReleaseCatCache(m, v41)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								v87 = v83
								m.G0 = v8 + int32(32)
								return v87
							}
						}
					}
				} else {
					v38 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(v32))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v41 = v38
						if l2|v41 == int32(0) {
							v47 = F_SearchSysCache1(m, int32(2), base.I64_extend_i32_u(l0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									if v47 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
										F_errmsg_internal(m, int32(_a_F_get_opclass_oid_0), v8)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(202), int32(_a_F_get_opclass_oid_2))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
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
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = F_NameListToString(m, l1)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return int32(0)
											} else {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
												v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
												*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v58
												*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v60 + v61 + int32(4)
												F_errmsg(m, int32(_a_F_get_opclass_oid_3), v8+int32(16))
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_get_opclass_oid_1), int32(207), int32(_a_F_get_opclass_oid_2))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
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
							if v41 == int32(0) {
								v87 = int32(0)
								m.G0 = v8 + int32(32)
								return v87
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+22)))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v80+v81)))
								F_ReleaseCatCache(m, v41)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									v87 = v83
									m.G0 = v8 + int32(32)
									return v87
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
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	v11 = int64(0)
	v13 = F_SearchSysCacheList(m, int32(13), int32(1), base.I64_extend_i32_u(l0), v11, v11)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	if int32(0) < v17 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v13)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(-64)+v23<<(uint(int32(2))%32))))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
	v36 = v34 + v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
	if v37 != l1 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v46 = v23 + int32(1)
	if v46 != v17 {
		v23 = v46
		goto L6
	} else {
		goto L12
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
	if v39 != l2 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	F_ReleaseCatCacheList(m, v13)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	return v41
L12:
	;
	goto L7
L13:
	;
	return int32(0)
}
