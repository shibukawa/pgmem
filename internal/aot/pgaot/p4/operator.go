package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_format_operator(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_format_operator_extended(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_generate_operator_clause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = F_SearchSysCache1(m, int32(40), l3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
			v19 = v17 + v18
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
			v21 = F_get_namespace_name(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_appendStringInfoString(m, l0, l1)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)+80))
					if l2 != v27 {
						F_add_cast_to(m, l0, v27)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							v31 = F_quote_identifier(m, v21)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v31
								F_appendStringInfo(m, l0, int32(585985), v12+int32(32))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									F_appendStringInfoString(m, l0, v19+int32(4))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l4
										F_appendStringInfo(m, l0, int32(205684), v12+int32(16))
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
											if v47 != l5 {
												F_add_cast_to(m, l0, v47)
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return
												} else {
													F_ReleaseCatCache(m, v15)
													mBase = m.M
													v52 = m.ExcPending
													if v52 != 0 {
														return
													} else {
														m.G0 = v12 + int32(48)
														return
													}
												}
											} else {
												F_ReleaseCatCache(m, v15)
												mBase = m.M
												v52 = m.ExcPending
												if v52 != 0 {
													return
												} else {
													m.G0 = v12 + int32(48)
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						v31 = F_quote_identifier(m, v21)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v31
							F_appendStringInfo(m, l0, int32(585985), v12+int32(32))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								F_appendStringInfoString(m, l0, v19+int32(4))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l4
									F_appendStringInfo(m, l0, int32(205684), v12+int32(16))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
										if v47 != l5 {
											F_add_cast_to(m, l0, v47)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												F_ReleaseCatCache(m, v15)
												mBase = m.M
												v52 = m.ExcPending
												if v52 != 0 {
													return
												} else {
													m.G0 = v12 + int32(48)
													return
												}
											}
										} else {
											F_ReleaseCatCache(m, v15)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return
											} else {
												m.G0 = v12 + int32(48)
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = l3
				F_errmsg_internal(m, int32(43030), v12)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_errfinish(m, int32(492166), int32(13451), int32(356869))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
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
