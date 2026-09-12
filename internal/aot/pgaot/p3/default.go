package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_default_tablespace(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if base.B2i32(v13 == int32(2)) == int32(0) {
		v66 = v10
		m.G0 = v8 + int32(32)
		return v66
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _consts[130]))
		if v19 == int32(0) {
			v66 = v10
			m.G0 = v8 + int32(32)
			return v66
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
			if v23 == int32(0) {
				v66 = v10
				m.G0 = v8 + int32(32)
				return v66
			} else {
				v27 = F_get_tablespace_oid(m, v22, int32(1))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v27 != 0 {
						v66 = v10
						m.G0 = v8 + int32(32)
						return v66
					} else {
						if l2 == int32(12) {
							v35 = F_errstart(m, int32(18), int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v35 == int32(0) {
									v66 = v10
									m.G0 = v8 + int32(32)
									return v66
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v42
										F_errmsg(m, int32(78617), v8)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(524557), int32(1112), int32(439360))
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return int32(0)
											} else {
												v66 = v10
												m.G0 = v8 + int32(32)
												return v66
											}
										}
									}
								}
							}
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, _consts[43]))
							*(*int32)(unsafe.Add(mBase, _consts[512])) = v54
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v57
							v63 = F_format_elog_string(m, int32(606496), v8+int32(16))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[513])) = v63
								v66 = int32(0)
								m.G0 = v8 + int32(32)
								return v66
							}
						}
					}
				}
			}
		}
	}
}
func F_get_default_oid_from_partdesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v2 = int32(0)
	if l0 == v2 {
		v18 = v2
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v6 == int32(0) {
			v18 = v2
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
			if v9 == int32(-1) {
				v18 = v2
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+v9<<(uint(int32(2))%32))))
				v18 = v16
			}
		}
	}
	return v18
}
func F_update_default_partition_oid(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_table_open(m, int32(3350), int32(3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v16 = F_SearchSysCacheCopy(m, int32(45), l0, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(50064), v8)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_errfinish(m, int32(520777), int32(352), int32(455250))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
				*(*int32)(unsafe.Add(mBase, uint32(v33+v34)+8)) = l1
				F_CatalogTupleUpdate(m, v12, v16+int32(4), v16)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_pfree(m, v16)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_sequence_close(m, v12, int32(3))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
