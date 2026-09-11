package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dshash_create(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_palloc(m, int32(44))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v20 = F_dsa_allocate_extended(m, l0, int32(2580), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
			v23 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = v23
			v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v14)+12)) = v25
			v27 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v14)+20)) = v27
			*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l2
			v30 = F_dsa_get_address(m, l0, v20)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v30))) = v20
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(1979673120)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v37)+2568)) = v38
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+2568))
				v46 = int32(0)
				for {
					v55 = v40 + int32(8) + v46*int32(20)
					*(*uint16)(unsafe.Add(mBase, uint32(v55))) = uint16(v43)
					*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = int32(1073741824)
					*(*int64)(unsafe.Add(mBase, uint32(v55)+8)) = int64(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = int32(0)
					v64 = v46 + int32(1)
					if v64 != int32(128) {
						v46 = v64
						continue
					} else {
						break
					}
					break
				}
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v67)+2572)) = int32(7)
				v72 = F_dsa_allocate_extended(m, l0, int32(512), int32(6))
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v74)+2576)) = v72
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+2576))
					if v77 == int32(0) {
						F_dsa_free(m, l0, v20)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(_a_F_dshash_create_0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_dshash_create_1), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(512)
										F_errdetail(m, int32(_a_F_dshash_create_2), v11)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_dshash_create_3), int32(255), int32(_a_F_dshash_create_4))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
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
					} else {
						v103 = F_dsa_get_address(m, l0, v77)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v103
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+2572))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v107
							m.G0 = v11 + int32(16)
							return v14
						}
					}
				}
			}
		}
	}
}
func F_dshash_seq_init(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	v3 = l2
	v4 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v4
	return
}
