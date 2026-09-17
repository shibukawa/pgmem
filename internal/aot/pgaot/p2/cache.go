package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResetPlanCache(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v1 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_ResetPlanCache[0]))
	if base.B2i32(v5 == v1)|base.B2i32(v5 == int32(_a_F_ResetPlanCache_0)) == v1 {
		v13 = v5
		for {
			v17 = v13 - int32(5)
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			if v18 != int32(1) {
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(96))))
				if v23 != 0 {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					switch v27 - int32(137) {
					case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
						v31 = int32(1)
					default:
						v31 = int32(0)
					}
					if v31 != 0 {
						v59 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v59)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(12))))
						if v63 == v59 {
						} else {
							v66 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v63)+10)) = uint8(v66)
						}
					} else {
					}
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(92))))
					if v34 == int32(0) {
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
						if v38 != int32(6) {
							v53 = int32(1)
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
							v45 = v43 - int32(201)
							if base.Ui32(int32(41)) < base.Ui32(v45) {
								v53 = int32(0)
							} else {
								v53 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v45)) % 64)))
							}
						}
						if v53&int32(1) == int32(0) {
						} else {
							v59 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v59)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(12))))
							if v63 == v59 {
							} else {
								v66 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v63)+10)) = uint8(v66)
							}
						}
					}
				}
			}
			v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v70 != int32(_a_F_ResetPlanCache_0) {
				v13 = v70
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_ResetPlanCache[1]))
	v78 = int32(0)
	if base.B2i32(v77 == v78)|base.B2i32(v77 == int32(_a_F_ResetPlanCache_1)) == v78 {
		v85 = v77
		for {
			v90 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v85-int32(16)))) = uint8(v90)
			v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
			if v92 != int32(_a_F_ResetPlanCache_1) {
				v85 = v92
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
