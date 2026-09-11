package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_interval_finite(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
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
	var v27 int32
	_ = v27
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	if v6 != int32(2147483647) {
		if v6 != int32(-2147483648) {
			v27 = v4
			return v27
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			if v11 == int32(-2147483648) {
				v15 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
				if v15 == int64(-9223372036854775807-1) {
					v27 = int32(0)
					return v27
				} else {
					return int32(1)
				}
			} else {
				return int32(1)
			}
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		if v21 != int32(2147483647) {
			v27 = v4
		} else {
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
			v27 = base.B2i32(v24 != int64(9223372036854775807))
		}
		return v27
	}
}
func F_interval_larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v53 int64
	_ = v53
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	var v99 int64
	_ = v99
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v133 int64
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = v14 + int32(16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+12)))
	v22 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+8)))
	v23 = v19*int64(30) + v22
	v32 = int64(32)
	v33 = int64(20)
	v35 = int64(base.Ui64(v23) >> (uint(v32) % 64))
	v38 = int64(4294967295)
	v39 = int64(500654080)
	v41 = v23 & v38
	v42 = v39 * v41
	v46 = int64(base.Ui64(v42)>>(uint(v32)%64)) + v39*v35
	v53 = v41*v33 + v46&v38
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v23*int64(0) + v23>>(uint(int64(63))%64)*int64(86400000000) + v33*v35 + int64(base.Ui64(v46)>>(uint(v32)%64)) + int64(base.Ui64(v53)>>(uint(v32)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v42&v38 | v53<<(uint(v32)%64)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+12)))
	v68 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+8)))
	v69 = v65*int64(30) + v68
	v78 = int64(32)
	v79 = int64(20)
	v81 = int64(base.Ui64(v69) >> (uint(v78) % 64))
	v84 = int64(4294967295)
	v85 = int64(500654080)
	v87 = v69 & v84
	v88 = v85 * v87
	v92 = int64(base.Ui64(v88)>>(uint(v78)%64)) + v85*v81
	v99 = v87*v79 + v92&v84
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v69*int64(0) + v69>>(uint(int64(63))%64)*int64(86400000000) + v79*v81 + int64(base.Ui64(v92)>>(uint(v78)%64)) + int64(base.Ui64(v99)>>(uint(v78)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v88&v84 | v99<<(uint(v78)%64)
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	m.G0 = v14 + int32(32)
	v119 = v110 + v112
	v120 = v113 + v115
	v124 = int64(63)
	v127 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v112))) + (v111 + v110>>(uint(v124)%64))
	v133 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v115))) + (v114 + v113>>(uint(v124)%64))
	if v127 == v133 {
		v136 = base.B2i32(base.Ui64(v120) < base.Ui64(v119))
	} else {
		v136 = base.B2i32(v133 < v127)
	}
	if v136 != 0 {
		v137 = v18
	} else {
		v137 = v64
	}
	return v137
}
