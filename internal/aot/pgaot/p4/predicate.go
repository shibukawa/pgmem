package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_predicate_classify(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v6 - int32(1) {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(881)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(882)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(883)
		return int32(1)
	default:
		v84 = v3
		return v84
	case 19:
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
		if v36 == int32(0) {
			v84 = v3
			return v84
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
			if v39 != int32(35) {
				if v39 != int32(7) {
					v84 = v3
					return v84
				} else {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+24)))
					if v44 != 0 {
						v84 = v3
						return v84
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
						v46 = F_pg_detoast_datum(m, v45)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
							v53 = F_ArrayGetNItems(m, v50, v46+int32(16))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								if int32(100) < v53 {
									v84 = v3
									return v84
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(885)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(886)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(887)
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
									if v65 != 0 {
										v66 = int32(2)
									} else {
										v66 = int32(1)
									}
									return v66
								}
							}
						}
					}
				}
			} else {
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+20)))
				if v68 != 0 {
					v84 = v3
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
					if v69 != 0 {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
						if int32(100) < v70 {
							v84 = v3
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(888)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(889)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(890)
							v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
							if v81 != 0 {
								v82 = int32(2)
							} else {
								v82 = int32(1)
							}
							v84 = v82
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(888)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(889)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(890)
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v81 != 0 {
							v82 = int32(2)
						} else {
							v82 = int32(1)
						}
						v84 = v82
					}
				}
				return v84
			}
		}
	case 20:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		switch v17 {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(881)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(882)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(884)
			return int32(1)
		case 1:
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(881)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(882)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(884)
			return int32(2)
		default:
			v84 = v3
			return v84
		}
	}
}
