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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v16 = F_SearchSysCache1(m, int32(40), base.I64_extend_i32_u(l3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
			v20 = v18 + v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
			v22 = F_get_namespace_name(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_appendStringInfoString(m, l0, l1)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
					if l2 != v26 {
						F_add_cast_to(m, l0, v26)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							v32 = F_quote_identifier(m, v22)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v32
								F_appendStringInfo(m, l0, int32(_a_F_generate_operator_clause_0), v12+int32(32))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									F_appendStringInfoString(m, l0, v20+int32(4))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l4
										F_appendStringInfo(m, l0, int32(_a_F_generate_operator_clause_1), v12+int32(16))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
											if v48 != l5 {
												F_add_cast_to(m, l0, v48)
												mBase = m.M
												v51 = m.ExcPending
												if v51 != 0 {
													return
												} else {
													F_ReleaseCatCache(m, v16)
													mBase = m.M
													v53 = m.ExcPending
													if v53 != 0 {
														return
													} else {
														m.G0 = v12 + int32(48)
														return
													}
												}
											} else {
												F_ReleaseCatCache(m, v16)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
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
						v32 = F_quote_identifier(m, v22)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v32
							F_appendStringInfo(m, l0, int32(_a_F_generate_operator_clause_0), v12+int32(32))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_appendStringInfoString(m, l0, v20+int32(4))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l4
									F_appendStringInfo(m, l0, int32(_a_F_generate_operator_clause_1), v12+int32(16))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return
									} else {
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
										if v48 != l5 {
											F_add_cast_to(m, l0, v48)
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return
											} else {
												F_ReleaseCatCache(m, v16)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return
												} else {
													m.G0 = v12 + int32(48)
													return
												}
											}
										} else {
											F_ReleaseCatCache(m, v16)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
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
			v60 = m.ExcPending
			if v60 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = l3
				F_errmsg_internal(m, int32(_a_F_generate_operator_clause_2), v12)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_generate_operator_clause_3), int32(_a_F_generate_operator_clause_4), int32(_a_F_generate_operator_clause_5))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
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
