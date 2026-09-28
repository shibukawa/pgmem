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
	var v24 int32
	_ = v24
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v13 = F_table_open(m, int32(1213), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = v9 + int32(16)
		F_ScanKeyInit(m, v18, int32(2), int32(3), int32(62), base.I64_extend_i32_u(l0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v26 = F_table_beginscan_catalog(m, v13, int32(1), v18)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = F_heap_getnext(m, v26)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					if v28 != 0 {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v30+v31)))
						v36 = v33
					} else {
						v36 = int32(0)
					}
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+188))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
					m.T0[v39].(func(*base.Module, int32))(m, v26)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_relation_close(m, v13, int32(1))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							if l1|v36 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
										F_errmsg(m, int32(_a_F_get_tablespace_oid_0), v9)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_get_tablespace_oid_1), int32(1487), int32(_a_F_get_tablespace_oid_2))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
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
								m.G0 = v9 + int32(80)
								return v36
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
