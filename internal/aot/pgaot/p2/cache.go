package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResetPlanCache(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1164]))
	if v5 == int32(0) {
	} else {
		if v5 == int32(4126848) {
		} else {
			v10 = v5
			for {
				v14 = v10 - int32(5)
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v15 != int32(1) {
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v10-int32(96))))
					if v20 != 0 {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						switch v24 - int32(137) {
						case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
							v28 = int32(1)
						default:
							v28 = int32(0)
						}
						if v28 != 0 {
							v56 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v56)
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v10-int32(12))))
							if v60 == v56 {
							} else {
								v63 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v60)+10)) = uint8(v63)
							}
						} else {
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v10-int32(92))))
						if v31 == int32(0) {
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
							if v35 != int32(6) {
								v50 = int32(1)
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
								v42 = v40 - int32(201)
								if base.Ui32(int32(41)) < base.Ui32(v42) {
									v50 = int32(0)
								} else {
									v50 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v42)) % 64)))
								}
							}
							if v50&int32(1) == int32(0) {
							} else {
								v56 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v56)
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v10-int32(12))))
								if v60 == v56 {
								} else {
									v63 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v60)+10)) = uint8(v63)
								}
							}
						}
					}
				}
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				if v67 != int32(4126848) {
					v10 = v67
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v74 = *(*int32)(unsafe.Add(mBase, _consts[1165]))
	if v74 == int32(0) {
	} else {
		if v74 == int32(4126856) {
		} else {
			v79 = v74
			for {
				v84 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v79-int32(16)))) = uint8(v84)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
				if v86 != int32(4126856) {
					v79 = v86
					continue
				} else {
					break
				}
				break
			}
		}
	}
	return
}
