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
	var v23 int32
	_ = v23
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v13 = F_table_open(m, int32(1213), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v7+int32(-48), int32(2), int32(3), int32(62), l0)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v27 = F_table_beginscan_catalog(m, v13, int32(1), v7+int32(-48))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = F_heap_getnext(m, v27)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v29 != 0 {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v31+v32)))
						v37 = v34
					} else {
						v37 = int32(0)
					}
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+188))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
					m.T0[v40].(func(*base.Module, int32))(m, v27)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_sequence_close(m, v13, int32(1))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							if l1 != 0 {
								m.G0 = v9 - int32(-64)
								return v37
							} else {
								if v37 != 0 {
									m.G0 = v9 - int32(-64)
									return v37
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67137668))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
											F_errmsg(m, int32(72110), v9)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(494516), int32(1461), int32(430695))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
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
					v16 = *(*float64)(unsafe.Add(mBase, _consts[484]))
					v17 = v16
				}
			} else {
				v16 = *(*float64)(unsafe.Add(mBase, _consts[484]))
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
					v29 = *(*float64)(unsafe.Add(mBase, _consts[483]))
					v30 = v29
				}
			} else {
				v29 = *(*float64)(unsafe.Add(mBase, _consts[483]))
				v30 = v29
			}
			*(*float64)(unsafe.Add(mBase, uint32(l2))) = v30
		} else {
		}
		return
	}
}
