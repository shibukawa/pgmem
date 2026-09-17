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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v43 int64
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
	var v114 int32
	_ = v114
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int64
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int64
	_ = v160
	var v165 int64
	_ = v165
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
			v25 = l1
			v26 = l2
			v27 = l3
			v28 = int32(0)
			for {
				v33 = int32(8)
				v34 = v24 - v33
				v36 = v26 + v33
				v38 = v25 + v33
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
				v43 = base.I64_popcnt(v39^v40) + v27
				v45 = v28 + int32(1)
				if v45 != (int32(base.Ui32(v13)>>(uint(v18)%32))+int32(1))&v18 {
					v24 = v34
					v25 = v38
					v26 = v36
					v27 = v43
					v28 = v45
					continue
				} else {
					break
				}
				break
			}
			v47 = v34
			v48 = v38
			v49 = v36
			v50 = v43
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
		v165 = v98
	} else {
		v106 = int32(0)
		if v95 != int32(1) {
			v114 = int32(0)
			v117 = v98
			v118 = v106
			for {
				v124 = v118 | int32(1)
				v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v124))))
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v124))))
				v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v126^v128)+uint32(_c_F_BitHammingDistanceDefault[0]))))
				v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v118))))
				v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v118))))
				v136 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v132^v134)+uint32(_c_F_BitHammingDistanceDefault[0]))))
				v138 = v130 + (v117 + v136)
				v139 = int32(2)
				v140 = v118 + v139
				v142 = v114 + v139
				if v142 != v95&int32(-2) {
					v114 = v142
					v117 = v138
					v118 = v140
					continue
				} else {
					break
				}
				break
			}
			if v95&int32(1) == int32(0) {
				v165 = v138
			} else {
				v149 = v138
				v150 = v140
				v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v150))))
				v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v150))))
				v160 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v156^v158)+uint32(_c_F_BitHammingDistanceDefault[0]))))
				v165 = v149 + v160
			}
		} else {
			v149 = v98
			v150 = v106
			v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v150))))
			v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v150))))
			v160 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v156^v158)+uint32(_c_F_BitHammingDistanceDefault[0]))))
			v165 = v149 + v160
		}
	}
	return v165
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
