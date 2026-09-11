package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_build_base_rel_tlists(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v4 = F_pull_var_clause(m, l1, int32(26))
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		if v4 != 0 {
			v7 = F_bms_make_singleton(m, int32(0))
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_add_vars_to_targetlist(m, l0, v4, v7)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					F_list_free(m, v4)
					mBase = m.M
					v12 = m.ExcPending
					if v12 != 0 {
						return
					} else {
						v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+112))
						if v14 == int32(0) {
							return
						} else {
							v18 = F_pull_var_clause(m, v14, int32(18))
							mBase = m.M
							v19 = m.ExcPending
							if v19 != 0 {
								return
							} else {
								if v18 == int32(0) {
									return
								} else {
									v23 = F_bms_make_singleton(m, int32(0))
									mBase = m.M
									v24 = m.ExcPending
									if v24 != 0 {
										return
									} else {
										F_add_vars_to_targetlist(m, l0, v18, v23)
										mBase = m.M
										v26 = m.ExcPending
										if v26 != 0 {
											return
										} else {
											F_list_free(m, v18)
											mBase = m.M
											v28 = m.ExcPending
											if v28 != 0 {
												return
											} else {
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
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+112))
			if v14 == int32(0) {
				return
			} else {
				v18 = F_pull_var_clause(m, v14, int32(18))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					if v18 == int32(0) {
						return
					} else {
						v23 = F_bms_make_singleton(m, int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return
						} else {
							F_add_vars_to_targetlist(m, l0, v18, v23)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								F_list_free(m, v18)
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
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
func F_get_base_element_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v5 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v5 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v9 = v5
	goto L6
L4:
	;
	goto L5
L5:
	;
	return int32(0)
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	v14 = v12 + v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+79)))
	if v15 != int32(100) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	if v18 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+132))
	F_ReleaseCatCache(m, v9)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	F_ReleaseCatCache(m, v9)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	if v19 == int32(6179) {
		v23 = v18
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v23 = int32(0)
	goto L11
L15:
	;
	goto L14
L16:
	;
	return v23
L17:
	;
	v31 = F_SearchSysCache1(m, int32(82), v27)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v31 != 0 {
		v9 = v31
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L7
}
