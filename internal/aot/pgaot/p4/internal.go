package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_create_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if base.Ui32(int32(4095)) < base.Ui32(l1) {
		v22 = F__emscripten_memset_bulkmem(m, l0+int32(20), base.I32_extend8_s(int32(0)), int32(1472))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = l3 ^ int32(216163848)
		v26 = int64(-1)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v26
		*(*int32)(unsafe.Add(mBase, uint32(l0)+1452)) = int32(-1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+1444)) = l6
		*(*int32)(unsafe.Add(mBase, uint32(l0)+1440)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0)+1448)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v26
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l3
		*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v26
		*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v26
		*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v26
		*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v26
		*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = v26
		*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v26
		*(*int64)(unsafe.Add(mBase, uint32(l0)+216)) = v26
		v55 = int32(base.Ui32(l1)>>(uint(int32(10))%32)) & int32(4194300)
		v57 = v55 + int32(2048)
		v59 = v57 & int32(4092)
		if v59 != 0 {
			v63 = v55 - v59 + int32(6144)
		} else {
			v63 = v57
		}
		v64 = l1 - v63
		v66 = int32(base.Ui32(v64) >> (uint(int32(12)) % 32))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v66
		*(*int32)(unsafe.Add(mBase, uint32(l0)+1472)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(l0)+1460)) = int32(1)
		v72 = F_palloc(m, int32(656))
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v72))) = l0
			v78 = *(*int32)(unsafe.Add(mBase, _consts[182]))
			*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v78
			v85 = F__emscripten_memset_bulkmem(m, v72+int32(8), base.I32_extend8_s(int32(0)), int32(648))
			mBase = m.M
			v87 = l0 + int32(1476)
			v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1472))
			*(*uint16)(unsafe.Add(mBase, uint32(v87))) = uint16(v88)
			*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = int32(1073741824)
			*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = int64(-1)
			v103 = int32(0)
			for {
				v106 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
				v111 = v106 + v103<<(uint(int32(5))%32) + int32(224)
				v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1472))
				*(*uint16)(unsafe.Add(mBase, uint32(v111))) = uint16(v112)
				*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = int32(1073741824)
				*(*int64)(unsafe.Add(mBase, uint32(v111)+8)) = int64(-1)
				v119 = v103 + int32(1)
				if v119 != int32(38) {
					v103 = v119
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(v72)+16)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = l0 + int32(2048)
			v129 = l0 + int32(1496)
			*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = v129
			v131 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v129)+4)) = v131
			v133 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v129)+32)) = uint8(v133)
			*(*int64)(unsafe.Add(mBase, uint32(v129)+12)) = v131
			*(*int64)(unsafe.Add(mBase, uint32(v129)+20)) = v131
			v139 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v129)+28)) = v139
			if v129 != 0 {
				v145 = v129 - l0 + v133
			} else {
				v145 = v139
			}
			*(*int32)(unsafe.Add(mBase, uint32(v129))) = v145
			v151 = F___memset(m, l0+int32(1532), int32(0), int32(516))
			mBase = m.M
			v153 = base.B2i32(base.Ui32(v64) < base.Ui32(int32(4096)))
			if v153 == int32(0) {
				v156 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
				F_FreePageManagerPut(m, v156, int32(base.Ui32(v63)>>(uint(int32(12))%32)), v66)
				mBase = m.M
				v160 = m.ExcPending
				if v160 != 0 {
					return int32(0)
				} else {
					v162 = int32(15)
					v165 = int32(32) - base.I32_clz(v66)
					if base.Ui32(v162) <= base.Ui32(v165) {
						v168 = v162
					} else {
						v168 = v165
					}
					if base.Ui32(v64) < base.Ui32(int32(4096)) {
						v169 = int32(0)
					} else {
						v169 = v168
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0+int32(160)+v169<<(uint(int32(2))%32)))) = int32(0)
					v175 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v175)+20)) = v169
					m.G0 = v13 + int32(16)
					return v72
				}
			} else {
				v162 = int32(15)
				v165 = int32(32) - base.I32_clz(v66)
				if base.Ui32(v162) <= base.Ui32(v165) {
					v168 = v162
				} else {
					v168 = v165
				}
				if base.Ui32(v64) < base.Ui32(int32(4096)) {
					v169 = int32(0)
				} else {
					v169 = v168
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0+int32(160)+v169<<(uint(int32(2))%32)))) = int32(0)
				v175 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v175)+20)) = v169
				m.G0 = v13 + int32(16)
				return v72
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v184 = m.ExcPending
		if v184 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(4096)
			F_errmsg_internal(m, int32(461156), v13)
			mBase = m.M
			v190 = m.ExcPending
			if v190 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(499500), int32(1240), int32(311743))
				mBase = m.M
				v195 = m.ExcPending
				if v195 != 0 {
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
