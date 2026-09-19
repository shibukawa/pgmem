package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitializeShmemGUCs(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v5 = m.G0
	v7 = v5 - int32(112)
	m.G0 = v7
	v11 = F_CalculateShmemSize(m, v7+int32(40))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v14 = F_add_size(m, v11, int32(_a_F_InitializeShmemGUCs_0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(base.Ui32(v14) >> (uint(int32(20)) % 32))
			v20 = v7 + int32(48)
			v24 = F_pg_sprintf(m, v20, int32(_a_F_InitializeShmemGUCs_1), v7+int32(32))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_2), v20, int32(0), int32(1))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeShmemGUCs[0]))
					if v32 != 0 {
						v36 = v32 << (uint(int32(10)) % 32)
					} else {
						v36 = int32(_a_F_InitializeShmemGUCs_3)
					}
					if v36 != 0 {
					} else {
					}
					v49 = v7 + int32(44)
					if v49 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v49))) = v36
					} else {
					}
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
					if v51 != 0 {
						v52 = base.I32_div_u_s(v11, v51)
						v54 = F_add_size(m, v52, int32(1))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v54
							v58 = v7 + int32(48)
							v62 = F_pg_sprintf(m, v58, int32(_a_F_InitializeShmemGUCs_1), v7+int32(16))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_4), v58, int32(0), int32(1))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v70
									v73 = v7 + int32(48)
									v75 = F_pg_sprintf(m, v73, int32(_a_F_InitializeShmemGUCs_5), v7)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_6), v73, int32(0), int32(1))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											m.G0 = v7 + int32(112)
											return
										}
									}
								}
							}
						}
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v70
						v73 = v7 + int32(48)
						v75 = F_pg_sprintf(m, v73, int32(_a_F_InitializeShmemGUCs_5), v7)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							F_SetConfigOption(m, int32(_a_F_InitializeShmemGUCs_6), v73, int32(0), int32(1))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								m.G0 = v7 + int32(112)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_ShmemInitHash(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v36 int32
	_ = v36
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v19 = int32(1073741823)
	if v19 <= l2 {
		v22 = v19
	} else {
		v22 = l2
	}
	v36 = int32(256)
	for {
		if v36 < int32(1)<<(uint(v6-base.I32_clz(int32(base.Ui32(int32(-1)<<(uint(v6-base.I32_clz(v22-int32(1)))%32)^int32(-1))>>(uint(int32(8))%32))))%32) {
			v36 = v36 << (uint(int32(1)) % 32)
			continue
		} else {
			break
		}
		break
	}
	*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = int32(1108)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v36
	v55 = F_ShmemInitStruct(m, l0, v36<<(uint(int32(2))%32)+int32(432), v11+int32(15))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		return int32(0)
	} else {
		v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
		*(*int32)(unsafe.Add(mBase, uint32(l3)+44)) = v55
		if v59 != 0 {
			v63 = l4 | int32(_a_F_ShmemInitHash_0)
		} else {
			v63 = l4 | int32(2564)
		}
		v64 = F_hash_create(m, l0, l1, l3, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			m.G0 = v11 + int32(16)
			return v64
		}
	}
}
