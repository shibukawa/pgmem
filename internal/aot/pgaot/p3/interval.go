package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_interval_finite(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v28 int64
	_ = v28
	v4 = int64(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	if v6 != int32(2147483647) {
		if v6 != int32(-2147483648) {
			v28 = v4
			return v28
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			if v11 == int32(-2147483648) {
				v15 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
				if v15 == int64(-9223372036854775807-1) {
					v28 = int64(0)
					return v28
				} else {
					return int64(1)
				}
			} else {
				return int64(1)
			}
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		if v21 != int32(2147483647) {
			v28 = v4
		} else {
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
			v28 = base.I64_extend_i32_u(base.B2i32(v24 != int64(9223372036854775807)))
		}
		return v28
	}
}
func F_interval_larger(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v47 int64
	_ = v47
	var v54 int64
	_ = v54
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v101 int64
	_ = v101
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v135 int64
	_ = v135
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = v14 + int32(16)
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = base.I32_wrap_i64(v18)
	v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v19)+12)))
	v23 = int64(*(*int32)(unsafe.Add(mBase, uint32(v19)+8)))
	v24 = v20*int64(30) + v23
	v33 = int64(32)
	v34 = int64(20)
	v36 = int64(base.Ui64(v24) >> (uint(v33) % 64))
	v39 = int64(4294967295)
	v40 = int64(500654080)
	v42 = v24 & v39
	v43 = v40 * v42
	v47 = int64(base.Ui64(v43)>>(uint(v33)%64)) + v40*v36
	v54 = v42*v34 + v47&v39
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v24*int64(0) + v24>>(uint(int64(63))%64)*int64(86400000000) + v34*v36 + int64(base.Ui64(v47)>>(uint(v33)%64)) + int64(base.Ui64(v54)>>(uint(v33)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v43&v39 | v54<<(uint(v33)%64)
	v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v66 = base.I32_wrap_i64(v65)
	v67 = int64(*(*int32)(unsafe.Add(mBase, uint32(v66)+12)))
	v70 = int64(*(*int32)(unsafe.Add(mBase, uint32(v66)+8)))
	v71 = v67*int64(30) + v70
	v80 = int64(32)
	v81 = int64(20)
	v83 = int64(base.Ui64(v71) >> (uint(v80) % 64))
	v86 = int64(4294967295)
	v87 = int64(500654080)
	v89 = v71 & v86
	v90 = v87 * v89
	v94 = int64(base.Ui64(v90)>>(uint(v80)%64)) + v87*v83
	v101 = v89*v81 + v94&v86
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v71*int64(0) + v71>>(uint(int64(63))%64)*int64(86400000000) + v81*v83 + int64(base.Ui64(v94)>>(uint(v80)%64)) + int64(base.Ui64(v101)>>(uint(v80)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v90&v86 | v101<<(uint(v80)%64)
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	m.G0 = v14 + int32(32)
	v121 = v115 + v117
	v122 = v112 + v114
	v126 = int64(63)
	v129 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v121) < base.Ui64(v117))) + (v116 + v115>>(uint(v126)%64))
	v135 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v122) < base.Ui64(v114))) + (v113 + v112>>(uint(v126)%64))
	if v135 == v129 {
		v138 = base.B2i32(base.Ui64(v122) < base.Ui64(v121))
	} else {
		v138 = base.B2i32(v135 < v129)
	}
	if v138 != 0 {
		v139 = v18
	} else {
		v139 = v65
	}
	return v139 & int64(4294967295)
}
