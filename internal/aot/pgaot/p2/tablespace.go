package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_tablespace_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v13 = F_table_open(m, int32(1213), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = v7 + int32(-48)
		F_ScanKeyInit(m, v18, int32(2), int32(3), int32(62), l0)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v25 = F_table_beginscan_catalog(m, v13, int32(1), v18)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = F_heap_getnext(m, v25)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 != 0 {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30)))
						v35 = v32
					} else {
						v35 = int32(0)
					}
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+188))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
					m.T0[v38].(func(*base.Module, int32))(m, v25)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_relation_close(m, v13, int32(1))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							if l1|v35 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
										F_errmsg(m, int32(_a_F_get_tablespace_oid_0), v9)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_tablespace_oid_1), int32(1461), int32(_a_F_get_tablespace_oid_2))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
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
								m.G0 = v9 - int32(-64)
								return v35
							}
						}
					}
				}
			}
		}
	}
}
func F_get_tablespace_page_costs(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v21 int32
	_ = v21
	var v22 float64
	_ = v22
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	v6 = F_get_tablespace(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		if l1 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			if v8 != 0 {
				v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
				if base.F64_lt(v9, float64(0)) == int32(0) {
					v17 = v9
				} else {
					v16 = *(*float64)(unsafe.Add(mBase, _c_F_get_tablespace_page_costs[0]))
					v17 = v16
				}
			} else {
				v16 = *(*float64)(unsafe.Add(mBase, _c_F_get_tablespace_page_costs[0]))
				v17 = v16
			}
			*(*float64)(unsafe.Add(mBase, uint32(l1))) = v17
		} else {
		}
		if l2 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			if v21 != 0 {
				v22 = *(*float64)(unsafe.Add(mBase, uint32(v21)+16))
				if base.F64_lt(v22, float64(0)) == int32(0) {
					v30 = v22
				} else {
					v29 = *(*float64)(unsafe.Add(mBase, _c_F_get_tablespace_page_costs[1]))
					v30 = v29
				}
			} else {
				v29 = *(*float64)(unsafe.Add(mBase, _c_F_get_tablespace_page_costs[1]))
				v30 = v29
			}
			*(*float64)(unsafe.Add(mBase, uint32(l2))) = v30
		} else {
		}
		return
	}
}
