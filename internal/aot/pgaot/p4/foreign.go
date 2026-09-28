package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ForeignServerConnectionString(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v12 = F_GetForeignDataWrapperExtended(m, v10, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
		if v16 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v26
					F_errmsg(m, int32(_a_F_ForeignServerConnectionString_0), v8)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v33 = F_errdetail(m, int32(_a_F_ForeignServerConnectionString_1), int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ForeignServerConnectionString_2), int32(214), int32(_a_F_ForeignServerConnectionString_3))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
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
		} else {
			v41 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1))))
			v43 = F_OidFunctionCall3Coll(m, v16, base.I64_extend_i32_u(l0), v41, int64(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v46 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v43))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v48 = F_text_to_cstring(m, v46)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v48
					}
				}
			}
		}
	}
}
func F_GetForeignDataWrapperExtended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = F_SearchSysCache1(m, int32(30), base.I64_extend_i32_u(l0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			if l1&int32(1) != 0 {
				v71 = int32(0)
				m.G0 = v10 + int32(16)
				return v71
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
					F_errmsg_internal(m, int32(_a_F_GetForeignDataWrapperExtended_0), v10)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_GetForeignDataWrapperExtended_1), int32(64), int32(_a_F_GetForeignDataWrapperExtended_2))
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
			v38 = F_palloc(m, int32(28))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = l0
				v41 = v35 + v36
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
				*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v42
				v46 = F_pstrdup(m, v41+int32(4))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v46
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
					*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v49
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v41)+76))
					*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v51
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v41)+80))
					*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v53
					v59 = F_SysCacheGetAttr(m, int32(30), v14, int32(8), v10+int32(15))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v61 != 0 {
							v65 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v65
							F_ReleaseCatCache(m, v14)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v71 = v38
								m.G0 = v10 + int32(16)
								return v71
							}
						} else {
							v63 = F_untransformRelOptions(m, v59)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								v65 = v63
								*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v65
								F_ReleaseCatCache(m, v14)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v71 = v38
									m.G0 = v10 + int32(16)
									return v71
								}
							}
						}
					}
				}
			}
		}
	}
}
