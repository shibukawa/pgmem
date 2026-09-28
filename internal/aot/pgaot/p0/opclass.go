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
	v8 = Fn14219(m, l0, l1, int32(13), int32(_a_F_OpclassIsVisibleExt_0), int32(2250), int32(_a_F_OpclassIsVisibleExt_1), int32(14))
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
	var v20 int32
	_ = v20
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
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(l0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v17 = v15 + v16
			if l1 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				v19 = F_GetDefaultOpClass(m, l1, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					if v19 == l0 {
						F_ReleaseCatCache(m, v13)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							m.G0 = v9 + int32(48)
							return
						}
					} else {
						v23 = v17 + int32(8)
						v24 = F_OpclassIsVisible(m, l0)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							if v24 != 0 {
								v26 = F_quote_identifier(m, v23)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v26
									F_appendStringInfo(m, l2, int32(_a_F_get_opclass_name_0), v9+int32(16))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v13)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return
										} else {
											m.G0 = v9 + int32(48)
											return
										}
									}
								}
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
								v35 = F_get_namespace_name_or_temp(m, v34)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									v37 = F_quote_identifier(m, v35)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										v39 = F_quote_identifier(m, v23)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v39
											*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v37
											F_appendStringInfo(m, l2, int32(_a_F_get_opclass_name_1), v9+int32(32))
											mBase = m.M
											v47 = m.ExcPending
											if v47 != 0 {
												return
											} else {
												F_ReleaseCatCache(m, v13)
												mBase = m.M
												v51 = m.ExcPending
												if v51 != 0 {
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
				v23 = v17 + int32(8)
				v24 = F_OpclassIsVisible(m, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					if v24 != 0 {
						v26 = F_quote_identifier(m, v23)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v26
							F_appendStringInfo(m, l2, int32(_a_F_get_opclass_name_0), v9+int32(16))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								F_ReleaseCatCache(m, v13)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									m.G0 = v9 + int32(48)
									return
								}
							}
						}
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
						v35 = F_get_namespace_name_or_temp(m, v34)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							v37 = F_quote_identifier(m, v35)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v39 = F_quote_identifier(m, v23)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v39
									*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v37
									F_appendStringInfo(m, l2, int32(_a_F_get_opclass_name_1), v9+int32(32))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										F_ReleaseCatCache(m, v13)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
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
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_get_opclass_name_2), v9)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_opclass_name_3), int32(_a_F_get_opclass_name_4), int32(_a_F_get_opclass_name_5))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
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
