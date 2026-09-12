package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InstrEndParallelQuery(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	v8 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	v9 = int32(4434696)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v12 = *(*int64)(unsafe.Add(mBase, _consts[103]))
	v13 = *(*int64)(unsafe.Add(mBase, _consts[440]))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v10 + (v12 - v13)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
	v19 = *(*int64)(unsafe.Add(mBase, _consts[104]))
	v20 = *(*int64)(unsafe.Add(mBase, _consts[441]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v17 + (v19 - v20)
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
	v26 = *(*int64)(unsafe.Add(mBase, _consts[105]))
	v27 = *(*int64)(unsafe.Add(mBase, _consts[442]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v24 + (v26 - v27)
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
	v33 = *(*int64)(unsafe.Add(mBase, _consts[106]))
	v34 = *(*int64)(unsafe.Add(mBase, _consts[443]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v31 + (v33 - v34)
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
	v40 = *(*int64)(unsafe.Add(mBase, _consts[107]))
	v41 = *(*int64)(unsafe.Add(mBase, _consts[444]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v38 + (v40 - v41)
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v8)+40))
	v47 = *(*int64)(unsafe.Add(mBase, _consts[108]))
	v48 = *(*int64)(unsafe.Add(mBase, _consts[445]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v45 + (v47 - v48)
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v8)+48))
	v54 = *(*int64)(unsafe.Add(mBase, _consts[109]))
	v55 = *(*int64)(unsafe.Add(mBase, _consts[446]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v52 + (v54 - v55)
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v8)+56))
	v61 = *(*int64)(unsafe.Add(mBase, _consts[110]))
	v62 = *(*int64)(unsafe.Add(mBase, _consts[447]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v59 + (v61 - v62)
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v8)+64))
	v68 = *(*int64)(unsafe.Add(mBase, _consts[111]))
	v69 = *(*int64)(unsafe.Add(mBase, _consts[448]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v66 + (v68 - v69)
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v8)+72))
	v75 = *(*int64)(unsafe.Add(mBase, _consts[112]))
	v76 = *(*int64)(unsafe.Add(mBase, _consts[449]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+72)) = v73 + (v75 - v76)
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
	v82 = *(*int64)(unsafe.Add(mBase, _consts[113]))
	v83 = *(*int64)(unsafe.Add(mBase, _consts[450]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+80)) = v80 + (v82 - v83)
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v8)+88))
	v89 = *(*int64)(unsafe.Add(mBase, _consts[114]))
	v90 = *(*int64)(unsafe.Add(mBase, _consts[451]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+88)) = v87 + (v89 - v90)
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v8)+96))
	v96 = *(*int64)(unsafe.Add(mBase, _consts[115]))
	v97 = *(*int64)(unsafe.Add(mBase, _consts[452]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+96)) = v94 + (v96 - v97)
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v8)+104))
	v103 = *(*int64)(unsafe.Add(mBase, _consts[116]))
	v104 = *(*int64)(unsafe.Add(mBase, _consts[453]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+104)) = v101 + (v103 - v104)
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v8)+112))
	v110 = *(*int64)(unsafe.Add(mBase, _consts[117]))
	v111 = *(*int64)(unsafe.Add(mBase, _consts[454]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+112)) = v108 + (v110 - v111)
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v8)+120))
	v117 = *(*int64)(unsafe.Add(mBase, _consts[118]))
	v118 = *(*int64)(unsafe.Add(mBase, _consts[455]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+120)) = v115 + (v117 - v118)
	v123 = l1 + int32(24)
	v124 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v123))) = v124
	v127 = l1 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v124
	v131 = l1 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v131))) = v124
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v124
	v137 = *(*int64)(unsafe.Add(mBase, _consts[43]))
	v139 = *(*int64)(unsafe.Add(mBase, _consts[44]))
	*(*int64)(unsafe.Add(mBase, uint32(v127))) = v137 - v139
	v143 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v145 = *(*int64)(unsafe.Add(mBase, _consts[48]))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v143 - v145
	v149 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	v151 = *(*int64)(unsafe.Add(mBase, _consts[46]))
	*(*int64)(unsafe.Add(mBase, uint32(v131))) = v149 - v151
	v155 = *(*int64)(unsafe.Add(mBase, _consts[41]))
	v157 = *(*int64)(unsafe.Add(mBase, _consts[42]))
	*(*int64)(unsafe.Add(mBase, uint32(v123))) = v155 - v157
	return
}
func F_InstrStartNode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7 == int32(1) {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		if v10 != int64(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(31751), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					F_errfinish(m, int32(504125), int32(72), int32(422946))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			F___clock_gettime(m, int32(1), v5)
			mBase = m.M
			v15 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5)+8)))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v15 + v16*int64(1000000000)
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if v21 == int32(1) {
				v28 = F__emscripten_memcpy_bulkmem(m, l0+int32(40), int32(4434536), int32(128))
				mBase = m.M
			} else {
			}
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			if v30 == int32(1) {
				v34 = *(*int64)(unsafe.Add(mBase, _consts[41]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v34
				v37 = *(*int64)(unsafe.Add(mBase, _consts[43]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v37
				v40 = *(*int64)(unsafe.Add(mBase, _consts[45]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v40
				v43 = *(*int64)(unsafe.Add(mBase, _consts[47]))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v43
			} else {
			}
			m.G0 = v5 + int32(16)
			return
		}
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v21 == int32(1) {
			v28 = F__emscripten_memcpy_bulkmem(m, l0+int32(40), int32(4434536), int32(128))
			mBase = m.M
		} else {
		}
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
		if v30 == int32(1) {
			v34 = *(*int64)(unsafe.Add(mBase, _consts[41]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v34
			v37 = *(*int64)(unsafe.Add(mBase, _consts[43]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v37
			v40 = *(*int64)(unsafe.Add(mBase, _consts[45]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+176)) = v40
			v43 = *(*int64)(unsafe.Add(mBase, _consts[47]))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v43
		} else {
		}
		m.G0 = v5 + int32(16)
		return
	}
}
