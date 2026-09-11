package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sn_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	if l2 != 0 {
		v6 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v6)
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v13 = v11 + v12
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13-int32(1)))))
		if v16 == int32(123) {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F_escape_json(m, v41, l1)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
				if v48 <= v45+int32(1) {
					F_appendStringInfoChar(m, v44, int32(58))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
					v57 = int32(58)
					*(*uint8)(unsafe.Add(mBase, uint32(v55+v45))) = uint8(v57)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
					v62 = v60 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v62
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					v66 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v64+v62))) = uint8(v66)
					return v66
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			if v19 <= v12+int32(1) {
				F_appendStringInfoChar(m, v10, int32(44))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_escape_json(m, v41, l1)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
						if v48 <= v45+int32(1) {
							F_appendStringInfoChar(m, v44, int32(58))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								return int32(0)
							}
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
							v57 = int32(58)
							*(*uint8)(unsafe.Add(mBase, uint32(v55+v45))) = uint8(v57)
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
							v62 = v60 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v62
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
							v66 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v64+v62))) = uint8(v66)
							return v66
						}
					}
				}
			} else {
				v28 = int32(44)
				*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v28)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
				v33 = v31 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v33
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
				v37 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v35+v33))) = uint8(v37)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_escape_json(m, v41, l1)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
					if v48 <= v45+int32(1) {
						F_appendStringInfoChar(m, v44, int32(58))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
						v57 = int32(58)
						*(*uint8)(unsafe.Add(mBase, uint32(v55+v45))) = uint8(v57)
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
						v62 = v60 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v62
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
						v66 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v64+v62))) = uint8(v66)
						return v66
					}
				}
			}
		}
	}
}
