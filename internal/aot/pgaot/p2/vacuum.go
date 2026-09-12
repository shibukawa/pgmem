package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vacuum_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	switch v10 - int32(1) {
	case 0:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		if v13 != int32(-1) {
			v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
				if base.Ui32((v16-int32(1))&int32(65535)) <= base.Ui32(int32(2047)) {
					v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
					v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v28
					*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v27
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v20
					F_errcontext_msg(m, int32(667588), v8+int32(16))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						m.G0 = v8 + int32(144)
						return
					}
				} else {
					v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = v37
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v20
					F_errcontext_msg(m, int32(667697), v8+int32(32))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						m.G0 = v8 + int32(144)
						return
					}
				}
			}
		} else {
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
				*(*int64)(unsafe.Add(mBase, uint32(v8))) = v48
				F_errcontext_msg(m, int32(667523), v8)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					m.G0 = v8 + int32(144)
					return
				}
			}
		}
	case 1:
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v97 = m.ExcPending
		if v97 != 0 {
			return
		} else {
			v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+104)) = v100
			*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = v99
			*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v98
			F_errcontext_msg(m, int32(667835), v8+int32(96))
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return
			} else {
				m.G0 = v8 + int32(144)
				return
			}
		}
	case 2:
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		if v53 != int32(-1) {
			v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
				if base.Ui32((v56-int32(1))&int32(65535)) <= base.Ui32(int32(2047)) {
					v67 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
					v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v68
					*(*int64)(unsafe.Add(mBase, uint32(v8)+72)) = v67
					*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v60
					F_errcontext_msg(m, int32(667642), v8-int32(-64))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						m.G0 = v8 + int32(144)
						return
					}
				} else {
					v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+84)) = v77
					*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v60
					F_errcontext_msg(m, int32(667741), v8+int32(80))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						m.G0 = v8 + int32(144)
						return
					}
				}
			}
		} else {
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return
			} else {
				v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
				*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v88
				F_errcontext_msg(m, int32(667555), v8+int32(48))
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return
				} else {
					m.G0 = v8 + int32(144)
					return
				}
			}
		}
	case 3:
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v111 = m.ExcPending
		if v111 != 0 {
			return
		} else {
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+120)) = v114
			*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = v113
			*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = v112
			F_errcontext_msg(m, int32(667786), v8+int32(112))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return
			} else {
				m.G0 = v8 + int32(144)
				return
			}
		}
	case 4:
		v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		if v123 == int32(-1) {
			m.G0 = v8 + int32(144)
			return
		} else {
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v128 = m.ExcPending
			if v128 != 0 {
				return
			} else {
				v129 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
				v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+136)) = v130
				*(*int64)(unsafe.Add(mBase, uint32(v8)+128)) = v129
				F_errcontext_msg(m, int32(151508), v8+int32(128))
				mBase = m.M
				v137 = m.ExcPending
				if v137 != 0 {
					return
				} else {
					m.G0 = v8 + int32(144)
					return
				}
			}
		}
	default:
		m.G0 = v8 + int32(144)
		return
	}
}
