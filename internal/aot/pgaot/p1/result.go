package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_wait_result_is_signal(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	v7 = l0 & int32(127)
	v13 = int32(255)
	return base.B2i32(l1 == v7)&base.B2i32(base.Ui32(l0&int32(_a_F_wait_result_is_signal_0)-int32(1)) < base.Ui32(v13)) | base.B2i32(v7 == int32(0))&base.B2i32(l1+int32(128) == int32(base.Ui32(l0)>>(uint(int32(8))%32))&v13)
}
func F_wait_result_to_str(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	v4 = m.G0
	v6 = v4 - int32(560)
	m.G0 = v6
	if l0 == int32(-1) {
		v15 = F_pg_snprintf(m, v6+int32(48), int32(512), int32(_a_F_wait_result_to_str_0), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v99 = F_pstrdup(m, v6+int32(48))
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(560)
				return v99
			}
		}
	} else {
		v20 = l0 & int32(127)
		if v20 == int32(0) {
			v26 = int32(base.Ui32(l0)>>(uint(int32(8))%32)) & int32(255)
			switch v26 - int32(126) {
			case 0:
				v34 = F_pg_snprintf(m, v6+int32(48), int32(512), int32(_a_F_wait_result_to_str_1), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v99 = F_pstrdup(m, v6+int32(48))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(560)
						return v99
					}
				}
			case 1:
				v41 = F_pg_snprintf(m, v6+int32(48), int32(512), int32(_a_F_wait_result_to_str_2), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v99 = F_pstrdup(m, v6+int32(48))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(560)
						return v99
					}
				}
			default:
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v26
				v48 = F_pg_snprintf(m, v6+int32(48), int32(512), int32(_a_F_wait_result_to_str_3), v6)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v99 = F_pstrdup(m, v6+int32(48))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(560)
						return v99
					}
				}
			}
		} else {
			if base.Ui32(l0&int32(_a_F_wait_result_to_str_4)-int32(1)) <= base.Ui32(int32(254)) {
				v58 = int32(_a_F_wait_result_to_str_5)
				if base.Ui32(int32(-64)) <= base.Ui32(v20-int32(65)) {
					v63 = v20
					v64 = v58
					for {
						v67 = v64 + int32(1)
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
						if v68 != 0 {
							v64 = v67
							continue
						} else {
						}
						v70 = v63 - int32(1)
						if v70 != 0 {
							v63 = v70
							v64 = v67
							continue
						} else {
							break
						}
						break
					}
					v72 = v67
				} else {
					v72 = v58
				}
				if v72 != 0 {
					v75 = v72
				} else {
					v75 = int32(_a_F_wait_result_to_str_6)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v75
				*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v20
				v84 = F_pg_snprintf(m, v6+int32(48), int32(512), int32(_a_F_wait_result_to_str_7), v6+int32(16))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					v99 = F_pstrdup(m, v6+int32(48))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(560)
						return v99
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
				v93 = F_pg_snprintf(m, v6+int32(48), int32(512), int32(_a_F_wait_result_to_str_8), v6+int32(32))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return int32(0)
				} else {
					v99 = F_pstrdup(m, v6+int32(48))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(560)
						return v99
					}
				}
			}
		}
	}
}
