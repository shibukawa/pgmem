package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_currval_oid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_init_sequence(m, v8, v6+int32(28), v6+int32(24))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_currval_oid[0]))
		v22 = F_pg_class_aclcheck(m, v18, v20, int64(258))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int64(0)
		} else {
			if v22 == int32(0) {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
				if v26 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int64(0)
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v68 + int32(4)
							F_errmsg(m, int32(_a_F_currval_oid_0), v6)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_currval_oid_1), int32(888), int32(_a_F_currval_oid_2))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v29 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
					F_relation_close(m, v30, int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						m.G0 = v6 + int32(32)
						return v29
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int64(0)
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+24))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v46 + int32(4)
						F_errmsg(m, int32(_a_F_currval_oid_3), v6+int32(16))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_currval_oid_1), int32(882), int32(_a_F_currval_oid_2))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int64(0)
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
