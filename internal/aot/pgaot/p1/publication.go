package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_getPublicationSchemaInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v14 = F_SearchSysCache1(m, int32(49), v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			if l1 != 0 {
				v76 = int32(0)
				m.G0 = v10 + int32(32)
				return v76
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v25
					F_errmsg_internal(m, int32(_a_F_getPublicationSchemaInfo_0), v10)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_getPublicationSchemaInfo_1), int32(2926), int32(_a_F_getPublicationSchemaInfo_2))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
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
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
			v37 = v35 + v36
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
			v39 = F_get_publication_name(m, v38, l1)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v39
				if v39 == int32(0) {
					F_ReleaseCatCache(m, v14)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						v76 = base.B2i32(v39 != int32(0))
						m.G0 = v10 + int32(32)
						return v76
					}
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
					v45 = F_get_namespace_name(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v45
						if v45 != 0 {
							F_ReleaseCatCache(m, v14)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v76 = base.B2i32(v39 != int32(0))
								m.G0 = v10 + int32(32)
								return v76
							}
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							F_pfree(m, v49)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v14)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									if l1 != 0 {
										v76 = int32(0)
										m.G0 = v10 + int32(32)
										return v76
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v48
											F_errmsg_internal(m, int32(_a_F_getPublicationSchemaInfo_3), v10+int32(16))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_getPublicationSchemaInfo_1), int32(2947), int32(_a_F_getPublicationSchemaInfo_2))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
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
func F_get_publication_name(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = Fn14293(m, l0, l1, int32(4), int32(_a_F_get_publication_name_0), int32(3990), int32(_a_F_get_publication_name_1), int32(51))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
