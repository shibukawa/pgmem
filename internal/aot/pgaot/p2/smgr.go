package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_smgr_bulk_finish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	F_smgr_bulk_flush(m, l0)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
		if v6 != int32(-1) {
			return
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
			if v9 == int32(0) {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_smgrregistersync(m, v5, v12)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					return
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_finish[0]))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+120))
				*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v17 | int32(1)
				v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+408))
				v22 = F_GetRedoRecPtr(m)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					if v21 != v22 {
						v26 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_finish[0]))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+120))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+120)) = v27 & int32(-2)
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v33 = int32(_a_F_smgr_bulk_finish_0)
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_finish[1]))
						*(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_finish[1])) = v35 + int32(1)
						F_mdimmedsync(m, v31, v32)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v41 = int32(_a_F_smgr_bulk_finish_0)
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_finish[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_finish[1])) = v43 - int32(1)
							v49 = F_errstart(m, int32(14), int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								if v49 == int32(0) {
									return
								} else {
									F_errmsg_internal(m, int32(_a_F_smgr_bulk_finish_1), int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_smgr_bulk_finish_2), int32(214), int32(_a_F_smgr_bulk_finish_3))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_smgrregistersync(m, v62, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, _c_F_smgr_bulk_finish[0]))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+120))
							*(*int32)(unsafe.Add(mBase, uint32(v67)+120)) = v68 & int32(-2)
							return
						}
					}
				}
			}
		}
	}
}
func F_smgr_identify(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v6 = l0 & int32(240)
	if v6 == int32(32) {
		v9 = int32(_a_F_smgr_identify_0)
	} else {
		v9 = int32(0)
	}
	if v6 == int32(16) {
		v12 = int32(_a_F_smgr_identify_1)
	} else {
		v12 = v9
	}
	return v12
}
