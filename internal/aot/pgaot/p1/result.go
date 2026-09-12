package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_wait_result_is_signal(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v25 int32
	_ = v25
	v5 = int32(1)
	v7 = l0 & int32(127)
	if l1 == v7 {
		if base.Ui32(l0&int32(65535)-int32(1)) < base.Ui32(int32(255)) {
			v25 = v5
		} else {
			if v7 == int32(0) {
				if l1+int32(128) == int32(base.Ui32(l0)>>(uint(int32(8))%32))&int32(255) {
					v25 = v5
				} else {
					v25 = int32(0)
				}
			} else {
				v25 = int32(0)
			}
		}
	} else {
		if v7 == int32(0) {
			if l1+int32(128) == int32(base.Ui32(l0)>>(uint(int32(8))%32))&int32(255) {
				v25 = v5
			} else {
				v25 = int32(0)
			}
		} else {
			v25 = int32(0)
		}
	}
	return v25
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	v4 = m.G0
	v6 = v4 - int32(560)
	m.G0 = v6
	if l0 == int32(-1) {
		v15 = F_pg_snprintf(m, v6+int32(48), int32(512), int32(299869), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v101 = F_pstrdup(m, v6+int32(48))
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(560)
				return v101
			}
		}
	} else {
		v20 = l0 & int32(127)
		if v20 == int32(0) {
			v26 = int32(base.Ui32(l0)>>(uint(int32(8))%32)) & int32(255)
			switch v26 - int32(126) {
			case 0:
				v34 = F_pg_snprintf(m, v6+int32(48), int32(512), int32(391530), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v101 = F_pstrdup(m, v6+int32(48))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(560)
						return v101
					}
				}
			case 1:
				v41 = F_pg_snprintf(m, v6+int32(48), int32(512), int32(423872), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v101 = F_pstrdup(m, v6+int32(48))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(560)
						return v101
					}
				}
			default:
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v26
				v48 = F_pg_snprintf(m, v6+int32(48), int32(512), int32(477609), v6)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v101 = F_pstrdup(m, v6+int32(48))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(560)
						return v101
					}
				}
			}
		} else {
			if base.Ui32(l0&int32(65535)-int32(1)) <= base.Ui32(int32(254)) {
				v58 = int32(4103920)
				if base.Ui32(v20-int32(65)) < base.Ui32(int32(-64)) {
					v75 = v58
				} else {
					v64 = v20
					v65 = v58
					for {
						v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
						v69 = v65 + int32(1)
						if v67 != 0 {
							v65 = v69
							continue
						} else {
						}
						v71 = v64 - int32(1)
						if v71 != 0 {
							v64 = v71
							v65 = v69
							continue
						} else {
							break
						}
						break
					}
					v75 = v69
				}
				if v75 != 0 {
					v77 = v75
				} else {
					v77 = int32(313597)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v77
				*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v20
				v86 = F_pg_snprintf(m, v6+int32(48), int32(512), int32(204306), v6+int32(16))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					v101 = F_pstrdup(m, v6+int32(48))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(560)
						return v101
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
				v95 = F_pg_snprintf(m, v6+int32(48), int32(512), int32(468842), v6+int32(32))
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return int32(0)
				} else {
					v101 = F_pstrdup(m, v6+int32(48))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(560)
						return v101
					}
				}
			}
		}
	}
}
