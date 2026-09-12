package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BitHammingDistanceDefault(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v86 int64
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int64
	_ = v152
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int64
	_ = v167
	var v172 int64
	_ = v172
	if base.Ui32(l0) < base.Ui32(int32(8)) {
		v95 = l0
		v96 = l1
		v97 = l2
		v98 = l3
	} else {
		v13 = l0 - int32(8)
		v14 = int32(24)
		if v13&v14 != v14 {
			v18 = int32(3)
			v24 = l0
			v27 = l3
			v28 = l1
			v29 = l2
			v30 = int32(0)
			for {
				v33 = int32(8)
				v34 = v24 - v33
				v35 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
				v36 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
				v39 = base.I64_popcnt(v35^v36) + v27
				v41 = v28 + v33
				v43 = v29 + v33
				v45 = v30 + int32(1)
				if v45 != (int32(base.Ui32(v13)>>(uint(v18)%32))+int32(1))&v18 {
					v24 = v34
					v27 = v39
					v28 = v41
					v29 = v43
					v30 = v45
					continue
				} else {
					break
				}
				break
			}
			v47 = v34
			v48 = v41
			v49 = v43
			v50 = v39
		} else {
			v47 = l0
			v48 = l1
			v49 = l2
			v50 = l3
		}
		if base.Ui32(v13) <= base.Ui32(int32(23)) {
			v95 = v47
			v96 = v48
			v97 = v49
			v98 = v50
		} else {
			v58 = v47
			v59 = v48
			v60 = v49
			v61 = v50
			for {
				v67 = *(*int64)(unsafe.Add(mBase, uint32(v60)+24))
				v68 = *(*int64)(unsafe.Add(mBase, uint32(v59)+24))
				v71 = *(*int64)(unsafe.Add(mBase, uint32(v60)+16))
				v72 = *(*int64)(unsafe.Add(mBase, uint32(v59)+16))
				v75 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
				v76 = *(*int64)(unsafe.Add(mBase, uint32(v59)+8))
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
				v80 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
				v86 = base.I64_popcnt(v67^v68) + (base.I64_popcnt(v71^v72) + (base.I64_popcnt(v75^v76) + (base.I64_popcnt(v79^v80) + v61)))
				v87 = int32(32)
				v88 = v60 + v87
				v90 = v59 + v87
				v92 = v58 - v87
				if base.Ui32(int32(7)) < base.Ui32(v92) {
					v58 = v92
					v59 = v90
					v60 = v88
					v61 = v86
					continue
				} else {
					break
				}
				break
			}
			v95 = v92
			v96 = v90
			v97 = v88
			v98 = v86
		}
	}
	if v95 == int32(0) {
		v172 = v98
	} else {
		v106 = int32(1)
		if v95 == v106 {
			v149 = int32(0)
			v152 = v98
		} else {
			v113 = int32(0)
			v115 = v113
			v118 = v98
			v119 = v113
			for {
				v125 = v115 | int32(1)
				v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v125))))
				v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v125))))
				v133 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v127^v129)+uint32(_consts[814]))))
				v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115+v97))))
				v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115+v96))))
				v141 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v135^v137)+uint32(_consts[814]))))
				v143 = v133 + (v118 + v141)
				v144 = int32(2)
				v145 = v115 + v144
				v147 = v119 + v144
				if v147 != v95&int32(-2) {
					v115 = v145
					v118 = v143
					v119 = v147
					continue
				} else {
					break
				}
				break
			}
			v149 = v145
			v152 = v143
		}
		if v95&v106 == int32(0) {
			v172 = v152
		} else {
			v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+v97))))
			v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+v96))))
			v167 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v161^v163)+uint32(_consts[814]))))
			v172 = v152 + v167
		}
	}
	return v172
}
func F_InitBitVector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v6 = int32(8)
	v7 = base.I32_div_s(l0+int32(7), v6)
	v9 = v7 + v6
	v10 = F_palloc0(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v9 << (uint(int32(2)) % 32)
		return v10
	}
}
func F_bit_out(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_varbit_out(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
