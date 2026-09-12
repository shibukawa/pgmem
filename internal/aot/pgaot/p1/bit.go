package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bit_catenate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v13 <= int32(2147483640)-v15 {
		v18 = v13 + v15
		v21 = int32(8)
		v22 = base.I32_div_s(v18+int32(7), v21)
		v24 = v22 + v21
		v25 = F_palloc(m, v24)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v18
			v30 = int32(2)
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = v24 << (uint(v30) % 32)
			v33 = int32(8)
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v41 = int32(base.Ui32(v37)>>(uint(v30)%32)) - v33
			if v41 != 0 {
				v42 = F__emscripten_memcpy_bulkmem(m, v25+v33, l0+v33, v41)
				mBase = m.M
			} else {
			}
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v46 = int32(base.Ui32(v44) >> (uint(int32(2)) % 32))
			v50 = v46<<(uint(int32(3))%32) + int32(-64)
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v50 == v51 {
				v54 = int32(8)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v60 = int32(base.Ui32(v56)>>(uint(int32(2))%32)) - v54
				if v60 != 0 {
					v61 = F__emscripten_memcpy_bulkmem(m, v25+v46, l1+v54, v60)
					mBase = m.M
				} else {
				}
			} else {
				if v15 <= int32(0) {
				} else {
					v66 = l1 + int32(8)
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if base.Ui32(l1+int32(base.Ui32(v67)>>(uint(int32(2))%32))) <= base.Ui32(v66) {
					} else {
						v73 = v50 - v51
						v78 = v25 + v46 - int32(1)
						v81 = v66
						for {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
							v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
							v89 = v86 | int32(base.Ui32(v87)>>(uint(int32(8)-v73)%32))
							*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v89)
							v92 = v78 + int32(1)
							v93 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
							if base.Ui32(v92) < base.Ui32(v25+int32(base.Ui32(v93)>>(uint(int32(2))%32))) {
								v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
								v99 = v98 << (uint(v73) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(v99)
							} else {
							}
							v102 = v81 + int32(1)
							v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if base.Ui32(v102) < base.Ui32(l1+int32(base.Ui32(v103)>>(uint(int32(2))%32))) {
								v78 = v92
								v81 = v102
								continue
							} else {
								break
							}
							break
						}
					}
				}
			}
			m.G0 = v11 + int32(16)
			return v25
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v123 = m.ExcPending
		if v123 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v126 = m.ExcPending
			if v126 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(2147483640)
				F_errmsg(m, int32(681509), v11)
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494687), int32(995), int32(355588))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
