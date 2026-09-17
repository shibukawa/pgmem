package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_create_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int64
	_ = v24
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if base.Ui32(int32(4095)) < base.Ui32(l1) {
		base.MemoryFill(m, l0+int32(20), int32(0), int32(1472))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = l3 ^ int32(216163848)
		v24 = int64(-1)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(l0)+1452)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+1444)) = l6
		*(*int32)(unsafe.Add(mBase, uint32(l0)+1440)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0)+1448)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l3
		*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v24
		*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v24
		*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v24
		*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v24
		*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = v24
		*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v24
		*(*int64)(unsafe.Add(mBase, uint32(l0)+216)) = v24
		v53 = int32(base.Ui32(l1)>>(uint(int32(10))%32)) & int32(_a_F_create_internal_0)
		v55 = v53 + int32(2048)
		v57 = v55 & int32(4092)
		if v57 != 0 {
			v61 = v53 - v57 + int32(_a_F_create_internal_1)
		} else {
			v61 = v55
		}
		v64 = int32(base.Ui32(l1-v61) >> (uint(int32(12)) % 32))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64
		*(*int32)(unsafe.Add(mBase, uint32(l0)+1472)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(l0)+1460)) = int32(1)
		v70 = F_palloc(m, int32(656))
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v70))) = l0
			v76 = *(*int32)(unsafe.Add(mBase, _c_F_create_internal[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v76
			base.MemoryFill(m, v70+int32(8), int32(0), int32(648))
			v84 = l0 + int32(1476)
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1472))
			*(*uint16)(unsafe.Add(mBase, uint32(v84))) = uint16(v85)
			*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = int32(1073741824)
			*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = int64(-1)
			v101 = int32(0)
			for {
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
				v107 = v102 + v101<<(uint(int32(5))%32) + int32(224)
				v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1472))
				*(*uint16)(unsafe.Add(mBase, uint32(v107))) = uint16(v108)
				*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = int32(1073741824)
				*(*int64)(unsafe.Add(mBase, uint32(v107)+8)) = int64(-1)
				v115 = v101 + int32(1)
				if v115 != int32(38) {
					v101 = v115
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = l0 + int32(2048)
			v125 = l0 + int32(1496)
			*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v125
			v127 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v125)+4)) = v127
			*(*int64)(unsafe.Add(mBase, uint32(v125)+12)) = v127
			*(*int64)(unsafe.Add(mBase, uint32(v125)+20)) = v127
			v133 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v125)+28)) = v133
			v135 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v125)+32)) = uint8(v135)
			if v125 != 0 {
				v141 = v125 - l0 + v135
			} else {
				v141 = v133
			}
			*(*int32)(unsafe.Add(mBase, uint32(v125))) = v141
			base.MemoryFill(m, l0+int32(1532), int32(0), int32(516))
			if v64 == int32(0) {
				v150 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0+int32(160)))) = v150
				v173 = v150
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v174)+20)) = v173
				m.G0 = v12 + int32(16)
				return v70
			} else {
				v153 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
				F_FreePageManagerPut(m, v153, int32(base.Ui32(v61)>>(uint(int32(12))%32)), v64)
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return int32(0)
				} else {
					v158 = int32(14)
					v161 = base.I32_clz(v64) ^ int32(31)
					if base.Ui32(v158) <= base.Ui32(v161) {
						v164 = v158
					} else {
						v164 = v161
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0+v164<<(uint(int32(2))%32))+164)) = int32(0)
					v173 = v164 + int32(1)
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v174)+20)) = v173
					m.G0 = v12 + int32(16)
					return v70
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v183 = m.ExcPending
		if v183 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_create_internal_2)
			F_errmsg_internal(m, int32(_a_F_create_internal_3), v12)
			mBase = m.M
			v189 = m.ExcPending
			if v189 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_create_internal_4), int32(1240), int32(_a_F_create_internal_5))
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
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
