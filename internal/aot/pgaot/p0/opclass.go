package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OpclassIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = Fn13856(m, l0, l1, int32(13), int32(_a_F_OpclassIsVisibleExt_0), int32(2181), int32(_a_F_OpclassIsVisibleExt_1), int32(14))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_get_opclass_name(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v32 int32
	_ = v32
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(14), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v16 = v14 + v15
			if l1 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				v18 = F_GetDefaultOpClass(m, l1, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					if v18 == l0 {
						F_ReleaseCatCache(m, v12)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							m.G0 = v9 + int32(48)
							return
						}
					} else {
						v22 = v16 + int32(8)
						v23 = F_OpclassIsVisible(m, l0)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							if v23 != 0 {
								v25 = F_quote_identifier(m, v22)
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v25
									F_appendStringInfo(m, l2, int32(_a_F_get_opclass_name_0), v9+int32(16))
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v12)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											m.G0 = v9 + int32(48)
											return
										}
									}
								}
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
								v34 = F_get_namespace_name_or_temp(m, v33)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return
								} else {
									v36 = F_quote_identifier(m, v34)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										v38 = F_quote_identifier(m, v22)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v38
											*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v36
											F_appendStringInfo(m, l2, int32(_a_F_get_opclass_name_1), v9+int32(32))
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
												return
											} else {
												F_ReleaseCatCache(m, v12)
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return
												} else {
													m.G0 = v9 + int32(48)
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
			} else {
				v22 = v16 + int32(8)
				v23 = F_OpclassIsVisible(m, l0)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					if v23 != 0 {
						v25 = F_quote_identifier(m, v22)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v25
							F_appendStringInfo(m, l2, int32(_a_F_get_opclass_name_0), v9+int32(16))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								F_ReleaseCatCache(m, v12)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									m.G0 = v9 + int32(48)
									return
								}
							}
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
						v34 = F_get_namespace_name_or_temp(m, v33)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							v36 = F_quote_identifier(m, v34)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v38 = F_quote_identifier(m, v22)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v38
									*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v36
									F_appendStringInfo(m, l2, int32(_a_F_get_opclass_name_1), v9+int32(32))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v12)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											m.G0 = v9 + int32(48)
											return
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
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_get_opclass_name_2), v9)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_opclass_name_3), int32(_a_F_get_opclass_name_4), int32(_a_F_get_opclass_name_5))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
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
