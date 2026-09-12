package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_extractNotNullColumn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	v4 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(21))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_pg_detoast_datum(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			if v10 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(23459), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(473998), int32(717), int32(263604))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				if v13 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(23459), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(473998), int32(717), int32(263604))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					if v14 != int32(21) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(23459), int32(0))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(473998), int32(717), int32(263604))
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
						if v17 == int32(1) {
							v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+24)))
							return v33
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(23459), int32(0))
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(473998), int32(717), int32(263604))
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
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
func F_makeNotNullConstraint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_palloc0(m, int32(108))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+17)) = uint8(v13)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v13
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(4294967457)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+104)) = int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)) = uint16(v13)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
		v28 = F_list_make1_impl(m, int32(1), v6+int32(8))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)) = uint8(v30)
			*(*uint16)(unsafe.Add(mBase, uint32(v9)+14)) = uint16(v30)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v28
			m.G0 = v6 + int32(16)
			return v9
		}
	}
}
