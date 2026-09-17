package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BitItemSize(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = int32(8)
	v5 = base.I32_div_s(l0+int32(7), v4)
	return v5 + v4
}
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
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
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
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v37 = int32(base.Ui32(v33)>>(uint(v30)%32)) - int32(8)
			if v37 != 0 {
				v38 = int32(8)
				base.MemoryCopy(m, v25+v38, l0+v38, v37)
			} else {
			}
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v45 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
			v49 = v45<<(uint(int32(3))%32) + int32(-64)
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v49 == v50 {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v56 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(8)
				if v56 == int32(0) {
				} else {
					base.MemoryCopy(m, v25+v45, l1+int32(8), v56)
				}
			} else {
				if v15 <= int32(0) {
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if base.Ui32(v65) < base.Ui32(int32(36)) {
					} else {
						v68 = int32(8)
						v69 = v49 - v50
						v76 = v25 + v45 - int32(1)
						v79 = l1 + v68
						for {
							v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
							v87 = v84 | int32(base.Ui32(v85)>>(uint(v68-v69)%32))
							*(*uint8)(unsafe.Add(mBase, uint32(v76))) = uint8(v87)
							v90 = v76 + int32(1)
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
							if base.Ui32(v90) < base.Ui32(v25+int32(base.Ui32(v91)>>(uint(int32(2))%32))) {
								v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
								v97 = v96 << (uint(v69) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v97)
							} else {
							}
							v100 = v79 + int32(1)
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if base.Ui32(v100) < base.Ui32(l1+int32(base.Ui32(v101)>>(uint(int32(2))%32))) {
								v76 = v90
								v79 = v100
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
		v121 = m.ExcPending
		if v121 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(2147483640)
				F_errmsg(m, int32(_a_F_bit_catenate_0), v11)
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_bit_catenate_1), int32(995), int32(_a_F_bit_catenate_2))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
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
