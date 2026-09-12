package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EventTriggerCollectAlterOpFam(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v9 = *(*int32)(unsafe.Add(mBase, _consts[287]))
	if v9 == int32(0) {
		return
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
		if v12 != 0 {
			return
		} else {
			v13 = int32(4442992)
			v14 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			*(*int32)(unsafe.Add(mBase, _consts[9])) = v16
			v19 = F_palloc(m, int32(40))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(3)
				v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[258])))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(2753)
				*(*uint8)(unsafe.Add(mBase, uint32(v19)+4)) = uint8(v24)
				v33 = F_copyObjectImpl(m, l0)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v33
					v37 = *(*int32)(unsafe.Add(mBase, _consts[287]))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
					v39 = F_lappend(m, v38, v19)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, _consts[287]))
						*(*int32)(unsafe.Add(mBase, uint32(v42)+28)) = v39
						*(*int32)(unsafe.Add(mBase, _consts[9])) = v14
						return
					}
				}
			}
		}
	}
}
func F_EventTriggerCollectAlterTSConfig(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v9 = *(*int32)(unsafe.Add(mBase, _consts[287]))
	if v9 == int32(0) {
		return
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
		if v12 != 0 {
			return
		} else {
			v13 = int32(4442992)
			v14 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			*(*int32)(unsafe.Add(mBase, _consts[9])) = v16
			v19 = F_palloc0(m, int32(40))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(6)
				v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[258])))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(3602)
				*(*uint8)(unsafe.Add(mBase, uint32(v19)+4)) = uint8(v24)
				v32 = l3 << (uint(int32(2)) % 32)
				v33 = F_palloc(m, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v33
					if v32 != 0 {
						v36 = F__emscripten_memcpy_bulkmem(m, v33, l2, v32)
						mBase = m.M
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = l3
					v39 = F_copyObjectImpl(m, l0)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v39
						v43 = *(*int32)(unsafe.Add(mBase, _consts[287]))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
						v45 = F_lappend(m, v44, v19)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, _consts[287]))
							*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = v45
							*(*int32)(unsafe.Add(mBase, _consts[9])) = v14
							return
						}
					}
				}
			}
		}
	}
}
