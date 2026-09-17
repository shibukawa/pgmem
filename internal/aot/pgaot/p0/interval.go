package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_interval_eq(m *base.Module, l0 int32) int32 {
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
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	m.G0 = v14 + int32(32)
	v119 = v113 + v115
	v120 = v110 + v112
	v124 = int64(63)
	return base.B2i32(v119^v120|(base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v115)))+(v114+v113>>(uint(v124)%64))^(base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v112)))+(v111+v110>>(uint(v124)%64)))) == int64(0))
}
func F_interval_ge(m *base.Module, l0 int32) int32 {
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
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	m.G0 = v14 + int32(32)
	v119 = v113 + v115
	v120 = v110 + v112
	v124 = int64(63)
	v127 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v115))) + (v114 + v113>>(uint(v124)%64))
	v133 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v112))) + (v111 + v110>>(uint(v124)%64))
	if v133 == v127 {
		v136 = base.B2i32(base.Ui64(v120) <= base.Ui64(v119))
	} else {
		v136 = base.B2i32(v133 <= v127)
	}
	return v136
}
func F_interval_hash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v15 int64
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7)+12)))
	v11 = int64(*(*int32)(unsafe.Add(mBase, uint32(v7)+8)))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = (v8*int64(30)+v11)*int64(86400000000) + v15
	v22 = F_DirectFunctionCall1Coll(m, int32(1269), int32(0), v5+int32(8))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(16)
		return v22
	}
}
func F_interval_lt(m *base.Module, l0 int32) int32 {
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
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	m.G0 = v14 + int32(32)
	v119 = v113 + v115
	v120 = v110 + v112
	v124 = int64(63)
	v127 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v115))) + (v114 + v113>>(uint(v124)%64))
	v133 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v112))) + (v111 + v110>>(uint(v124)%64))
	if v133 == v127 {
		v136 = base.B2i32(base.Ui64(v119) < base.Ui64(v120))
	} else {
		v136 = base.B2i32(v127 < v133)
	}
	return v136
}
func F_interval_mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		if v12 != int32(2147483647) {
			if v12 != int32(-2147483648) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				if v47 != int32(2147483647) {
					if v47 != int32(-2147483648) {
						F_finite_interval_mi(m, v6, v5, v8)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v52 != int32(-2147483648) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v55 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v55 != int64(-9223372036854775807-1) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
								return v8
							}
						}
					}
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
					if v58 != int32(2147483647) {
						F_finite_interval_mi(m, v6, v5, v8)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						v61 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
						if v61 != int64(9223372036854775807) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
							return v8
						}
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				if v17 != int32(-2147483648) {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					if v47 != int32(2147483647) {
						if v47 != int32(-2147483648) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
							if v52 != int32(-2147483648) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v55 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								if v55 != int64(-9223372036854775807-1) {
									F_finite_interval_mi(m, v6, v5, v8)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										return v8
									}
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
									*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
									return v8
								}
							}
						}
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v58 != int32(2147483647) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v61 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v61 != int64(9223372036854775807) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
								return v8
							}
						}
					}
				} else {
					v20 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
					if v20 != int64(-9223372036854775807-1) {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
						if v47 != int32(2147483647) {
							if v47 != int32(-2147483648) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
								if v52 != int32(-2147483648) {
									F_finite_interval_mi(m, v6, v5, v8)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										return v8
									}
								} else {
									v55 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
									if v55 != int64(-9223372036854775807-1) {
										F_finite_interval_mi(m, v6, v5, v8)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											return v8
										}
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
										*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
										return v8
									}
								}
							}
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
							if v58 != int32(2147483647) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v61 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								if v61 != int64(9223372036854775807) {
									F_finite_interval_mi(m, v6, v5, v8)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										return v8
									}
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
									*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
									return v8
								}
							}
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
						if v23 != int32(-2147483648) {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
							return v8
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
							if v26 != int32(-2147483648) {
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
								return v8
							} else {
								v29 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								if v29 == int64(-9223372036854775807-1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_interval_mi_0), int32(0))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_interval_mi_1), int32(3596), int32(_a_F_interval_mi_2))
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
									*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
									return v8
								}
							}
						}
					}
				}
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if v32 != int32(2147483647) {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				if v47 != int32(2147483647) {
					if v47 != int32(-2147483648) {
						F_finite_interval_mi(m, v6, v5, v8)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v52 != int32(-2147483648) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v55 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v55 != int64(-9223372036854775807-1) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
								return v8
							}
						}
					}
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
					if v58 != int32(2147483647) {
						F_finite_interval_mi(m, v6, v5, v8)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						v61 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
						if v61 != int64(9223372036854775807) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
							return v8
						}
					}
				}
			} else {
				v35 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
				if v35 != int64(9223372036854775807) {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					if v47 != int32(2147483647) {
						if v47 != int32(-2147483648) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
							if v52 != int32(-2147483648) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v55 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								if v55 != int64(-9223372036854775807-1) {
									F_finite_interval_mi(m, v6, v5, v8)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										return v8
									}
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
									*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
									return v8
								}
							}
						}
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v58 != int32(2147483647) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v61 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v61 != int64(9223372036854775807) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
								return v8
							}
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					if v38 != int32(2147483647) {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
						return v8
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v41 != int32(2147483647) {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
							return v8
						} else {
							v44 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v44 == int64(9223372036854775807) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_interval_mi_0), int32(0))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_interval_mi_1), int32(3605), int32(_a_F_interval_mi_2))
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
								return v8
							}
						}
					}
				}
			}
		}
	}
}
func F_interval_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int64
	_ = v66
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v115 int32
	_ = v115
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int64
	_ = v166
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v206 int64
	_ = v206
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v961 int32
	_ = v961
	var v967 int32
	_ = v967
	var v970 int64
	_ = v970
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1146 int64
	_ = v1146
	var v1150 int64
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1163 int64
	_ = v1163
	var v1171 int32
	_ = v1171
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1192 int64
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1202 int64
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1208 int64
	_ = v1208
	var v1216 int32
	_ = v1216
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int64
	_ = v1238
	var v1243 int64
	_ = v1243
	var v1247 int64
	_ = v1247
	var v1252 int64
	_ = v1252
	var v1258 int64
	_ = v1258
	var v1263 int64
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int64
	_ = v1266
	var v1274 int32
	_ = v1274
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1290 int64
	_ = v1290
	var v1292 int64
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1297 int64
	_ = v1297
	var v1305 int32
	_ = v1305
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1338 int32
	_ = v1338
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1363 int32
	_ = v1363
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	v17 = m.G0
	v19 = v17 - int32(176)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v22 != int32(-2147483648) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v1544 = F_pstrdup(m, v19)
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L25
	} else {
		goto L431
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v55
	v58 = int32(12)
	v59 = base.I32_div_s(v22, v58)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+168)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v19)+164)) = v22 - v59*v58
	v66 = base.I64_div_s(v56, int64(3600000000))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+152)) = v66
	v70 = v66*int64(-3600000000) + v56
	v72 = base.I64_div_s(v70, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+144)) = uint32(v72)
	v76 = v72*int64(-60000000) + v70
	v78 = base.I64_div_s(v76, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+140)) = uint32(v78)
	v82 = v78*int64(4293967296) + v76
	*(*uint32)(unsafe.Add(mBase, uint32(v19)+136)) = uint32(v82)
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_interval_out[0]))
	v86 = m.G0
	v88 = v86 - int32(416)
	m.G0 = v88
	v91 = v19 + int32(136)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v91)+16))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)+28))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+32))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v91)+24))
	v99 = base.I64_extend_i32_s(v98)
	switch v85 {
	case 0:
		goto L16
	default:
		goto L15
	case 2:
		goto L18
	case 3:
		goto L17
	}
L3:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_out[5])))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+8)) = uint8(v50)
	v53 = *(*int64)(unsafe.Add(mBase, _c_F_interval_out[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v53
	goto L1
L4:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_interval_out[7])))
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)) = uint16(v44)
	v47 = *(*int64)(unsafe.Add(mBase, _c_F_interval_out[8]))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v47
	goto L1
L5:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v37 != int32(2147483647) {
		v55 = v37
		v56 = v36
		goto L2
	} else {
		goto L12
	}
L6:
	;
	if v22 == int32(2147483647) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v30 != int32(-2147483648) {
		v55 = v30
		v56 = v29
		goto L2
	} else {
		goto L10
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v55 = v27
	v56 = v28
	goto L2
L10:
	;
	if v29 == int64(-9223372036854775807-1) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v55 = int32(-2147483648)
	v56 = v29
	goto L2
L12:
	;
	if v36 == int64(9223372036854775807) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v55 = int32(2147483647)
	v56 = v36
	goto L2
L14:
	;
	m.G0 = v88 + int32(416)
	goto L1
L15:
	;
	v1113 = int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v19))) = uint16(v1113)
	v1116 = v19 + int32(1)
	if v97 == int32(0) {
		goto L317
	} else {
		goto L318
	}
L16:
	;
	if v97 == int32(0) {
		goto L221
	} else {
		goto L222
	}
L17:
	;
	if v97|v96|(v98|base.B2i32(v95 != int64(0)))|(v92|(v94|v93)) == int32(0) {
		goto L157
	} else {
		goto L158
	}
L18:
	;
	v100 = v97 | v96
	v102 = int32(0)
	v104 = int64(0)
	v115 = base.B2i32(v100|v98 < v102) | base.B2i32(v95 < v104) | base.B2i32(v94 < v102) | base.B2i32(v93 < v102) | base.B2i32(v92 < v102)
	v135 = base.B2i32(v102 < v97) | base.B2i32(v102 < v96) | base.B2i32(v102 < v98) | base.B2i32(v104 < v95) | base.B2i32(v102 < v94) | base.B2i32(v102 < v93) | base.B2i32(v102 < v92)
	v138 = base.B2i32(v100 != v102)
	v154 = v115&v135 | v138&(base.B2i32(v98 != v102)|base.B2i32(v95 != v104)|base.B2i32(v94 != v102)|base.B2i32(v93 != v102)|base.B2i32(v92 != v102))
	if v154|base.B2i32(v115 == v102) == v102 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v160 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v160)
	v162 = int32(0)
	v166 = int64(0)
	v178 = v162 - v94
	v179 = v19 + int32(1)
	v180 = v162 - v97
	v181 = v162 - v96
	v182 = v162 - v93
	v183 = v162 - v92
	v184 = v166 - v99
	v185 = v166 - v95
	goto L21
L20:
	;
	v178 = v94
	v179 = v19
	v180 = v97
	v181 = v96
	v182 = v93
	v183 = v92
	v184 = v99
	v185 = v95
	goto L21
L21:
	;
	if v115|v135 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v191 = F_pg_sprintf(m, v179, int32(_a_F_interval_out_17), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v154 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	return int32(0)
L26:
	;
	goto L14
L27:
	;
	v196 = v178 >> (uint(int32(31)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+136)) = v178 ^ v196 - v196
	v200 = int64(63)
	v201 = v185 >> (uint(v200) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+128)) = v185 ^ v201 - v201
	v206 = v184 >> (uint(v200) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+112)) = v184 ^ v206 - v206
	v210 = int32(45)
	if v178|(v182|v183) < int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	if v100 != v102 {
		goto L78
	} else {
		goto L79
	}
L30:
	;
	v217 = v210
	goto L32
L31:
	;
	v217 = int32(43)
	goto L32
L32:
	;
	if v185 < int64(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v220 = v210
	goto L35
L34:
	;
	v220 = v217
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+120)) = v220
	if v184 < int64(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v226 = int32(45)
	goto L38
L37:
	;
	v226 = int32(43)
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+108)) = v226
	v228 = int32(31)
	v229 = v181 >> (uint(v228) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+104)) = v181 ^ v229 - v229
	v234 = v180 >> (uint(v228) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+100)) = v180 ^ v234 - v234
	if v180|v181 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v243 = int32(45)
	goto L41
L40:
	;
	v243 = int32(43)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+96)) = v243
	v248 = F_pg_sprintf(m, v179, int32(_a_F_interval_out_18), v88+int32(96))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L25
	} else {
		goto L42
	}
L42:
	;
	v250 = F_strlen(m, v179)
	mBase = m.M
	v256 = v182 >> (uint(int32(31)) % 32)
	goto L45
L43:
	;
	v368 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v367))) = uint8(v368)
	goto L14
L44:
	;
	if v183 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v260 = F_pg_ultostr_zeropad(m, v250+v179, v182^v256-v256, int32(2))
	mBase = m.M
	goto L44
L48:
	;
	v367 = v260
	goto L43
L49:
	;
	goto L50
L50:
	;
	v265 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v260))) = uint8(v265)
	v268 = v183 >> (uint(int32(31)) % 32)
	v270 = v183 ^ v268 - v268
	v272 = base.I32_div_s(v270, int32(10))
	v275 = v272*int32(-10) + v270
	if v275 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v285 = base.I32_div_s(v270, int32(100))
	v288 = v285*int32(-10) + v272
	v289 = v275 | v288
	if v289 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v277 = v275 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v260)+6)) = uint8(v277)
	v283 = v260 + int32(7)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v283 = v260 + int32(6)
	goto L51
L55:
	;
	v299 = base.I32_div_s(v270, int32(1000))
	v302 = v285 + v299*int32(-10)
	v303 = v289 | v302
	if v303 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v297 = v260 + int32(5)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v295 = v288 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v260)+5)) = uint8(v295)
	v297 = v283
	goto L55
L59:
	;
	v313 = base.I32_div_s(v270, int32(_a_F_interval_out_8))
	v316 = v299 + v313*int32(-10)
	v317 = v303 | v316
	if v317 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L60:
	;
	v311 = v260 + int32(4)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v309 = v302 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v260)+4)) = uint8(v309)
	v311 = v297
	goto L59
L63:
	;
	v327 = base.I32_div_s(v270, int32(_a_F_interval_out_9))
	v330 = v313 + v327*int32(-10)
	v331 = v317 | v330
	if v331 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v325 = v260 + int32(3)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v323 = v316 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v260)+3)) = uint8(v323)
	v325 = v311
	goto L63
L67:
	;
	v341 = base.I32_div_s(v270, int32(_a_F_interval_out_10))
	v344 = v341*int32(-10) + v327
	if v331|v344 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v339 = v260 + int32(2)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v337 = v330 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v260)+2)) = uint8(v337)
	v339 = v325
	goto L67
L71:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v327+int32(9)) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v353 = v260 + int32(1)
	goto L71
L73:
	;
	goto L74
L74:
	;
	v351 = v344 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)) = uint8(v351)
	v353 = v339
	goto L71
L75:
	;
	v360 = F_pg_ultostr(m, v260+int32(1), v270)
	mBase = m.M
	v361 = v360
	goto L77
L76:
	;
	v361 = v353
	goto L77
L77:
	;
	v367 = v361
	goto L43
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+148)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v88)+144)) = v180
	v375 = F_pg_sprintf(m, v179, int32(_a_F_interval_out_19), v88+int32(144))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L25
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v98 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L14
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+176)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v88)+168)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v88)+160)) = v184
	v383 = F_pg_sprintf(m, v179, int32(_a_F_interval_out_20), v88+int32(160))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L25
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+200)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v88)+192)) = v185
	v510 = F_pg_sprintf(m, v179, int32(_a_F_interval_out_21), v88+int32(192))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L25
	} else {
		goto L121
	}
L85:
	;
	v385 = F_strlen(m, v179)
	mBase = m.M
	v391 = v182 >> (uint(int32(31)) % 32)
	goto L88
L86:
	;
	v503 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v502))) = uint8(v503)
	goto L14
L87:
	;
	if v183 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v395 = F_pg_ultostr_zeropad(m, v385+v179, v182^v391-v391, int32(2))
	mBase = m.M
	goto L87
L91:
	;
	v502 = v395
	goto L86
L92:
	;
	goto L93
L93:
	;
	v400 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v395))) = uint8(v400)
	v403 = v183 >> (uint(int32(31)) % 32)
	v405 = v183 ^ v403 - v403
	v407 = base.I32_div_s(v405, int32(10))
	v410 = v407*int32(-10) + v405
	if v410 != 0 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v420 = base.I32_div_s(v405, int32(100))
	v423 = v420*int32(-10) + v407
	v424 = v410 | v423
	if v424 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	v412 = v410 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+6)) = uint8(v412)
	v418 = v395 + int32(7)
	goto L94
L96:
	;
	goto L97
L97:
	;
	v418 = v395 + int32(6)
	goto L94
L98:
	;
	v434 = base.I32_div_s(v405, int32(1000))
	v437 = v420 + v434*int32(-10)
	v438 = v424 | v437
	if v438 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	v432 = v395 + int32(5)
	goto L98
L100:
	;
	goto L101
L101:
	;
	v430 = v423 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+5)) = uint8(v430)
	v432 = v418
	goto L98
L102:
	;
	v448 = base.I32_div_s(v405, int32(_a_F_interval_out_8))
	v451 = v434 + v448*int32(-10)
	v452 = v438 | v451
	if v452 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L103:
	;
	v446 = v395 + int32(4)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v444 = v437 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+4)) = uint8(v444)
	v446 = v432
	goto L102
L106:
	;
	v462 = base.I32_div_s(v405, int32(_a_F_interval_out_9))
	v465 = v448 + v462*int32(-10)
	v466 = v452 | v465
	if v466 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	v460 = v395 + int32(3)
	goto L106
L108:
	;
	goto L109
L109:
	;
	v458 = v451 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+3)) = uint8(v458)
	v460 = v446
	goto L106
L110:
	;
	v476 = base.I32_div_s(v405, int32(_a_F_interval_out_10))
	v479 = v476*int32(-10) + v462
	if v466|v479 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	v474 = v395 + int32(2)
	goto L110
L112:
	;
	goto L113
L113:
	;
	v472 = v465 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+2)) = uint8(v472)
	v474 = v460
	goto L110
L114:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v462+int32(9)) {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v488 = v395 + int32(1)
	goto L114
L116:
	;
	goto L117
L117:
	;
	v486 = v479 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+1)) = uint8(v486)
	v488 = v474
	goto L114
L118:
	;
	v495 = F_pg_ultostr(m, v395+int32(1), v405)
	mBase = m.M
	v496 = v495
	goto L120
L119:
	;
	v496 = v488
	goto L120
L120:
	;
	v502 = v496
	goto L86
L121:
	;
	v512 = F_strlen(m, v179)
	mBase = m.M
	v518 = v182 >> (uint(int32(31)) % 32)
	goto L124
L122:
	;
	v630 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v629))) = uint8(v630)
	goto L14
L123:
	;
	if v183 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L124:
	;
	v522 = F_pg_ultostr_zeropad(m, v512+v179, v182^v518-v518, int32(2))
	mBase = m.M
	goto L123
L127:
	;
	v629 = v522
	goto L122
L128:
	;
	goto L129
L129:
	;
	v527 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v522))) = uint8(v527)
	v530 = v183 >> (uint(int32(31)) % 32)
	v532 = v183 ^ v530 - v530
	v534 = base.I32_div_s(v532, int32(10))
	v537 = v534*int32(-10) + v532
	if v537 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v547 = base.I32_div_s(v532, int32(100))
	v550 = v547*int32(-10) + v534
	v551 = v537 | v550
	if v551 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L131:
	;
	v539 = v537 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v522)+6)) = uint8(v539)
	v545 = v522 + int32(7)
	goto L130
L132:
	;
	goto L133
L133:
	;
	v545 = v522 + int32(6)
	goto L130
L134:
	;
	v561 = base.I32_div_s(v532, int32(1000))
	v564 = v547 + v561*int32(-10)
	v565 = v551 | v564
	if v565 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	v559 = v522 + int32(5)
	goto L134
L136:
	;
	goto L137
L137:
	;
	v557 = v550 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v522)+5)) = uint8(v557)
	v559 = v545
	goto L134
L138:
	;
	v575 = base.I32_div_s(v532, int32(_a_F_interval_out_8))
	v578 = v561 + v575*int32(-10)
	v579 = v565 | v578
	if v579 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L139:
	;
	v573 = v522 + int32(4)
	goto L138
L140:
	;
	goto L141
L141:
	;
	v571 = v564 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v522)+4)) = uint8(v571)
	v573 = v559
	goto L138
L142:
	;
	v589 = base.I32_div_s(v532, int32(_a_F_interval_out_9))
	v592 = v575 + v589*int32(-10)
	v593 = v579 | v592
	if v593 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L143:
	;
	v587 = v522 + int32(3)
	goto L142
L144:
	;
	goto L145
L145:
	;
	v585 = v578 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v522)+3)) = uint8(v585)
	v587 = v573
	goto L142
L146:
	;
	v603 = base.I32_div_s(v532, int32(_a_F_interval_out_10))
	v606 = v603*int32(-10) + v589
	if v593|v606 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L147:
	;
	v601 = v522 + int32(2)
	goto L146
L148:
	;
	goto L149
L149:
	;
	v599 = v592 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v522)+2)) = uint8(v599)
	v601 = v587
	goto L146
L150:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v589+int32(9)) {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	v615 = v522 + int32(1)
	goto L150
L152:
	;
	goto L153
L153:
	;
	v613 = v606 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v522)+1)) = uint8(v613)
	v615 = v601
	goto L150
L154:
	;
	v622 = F_pg_ultostr(m, v522+int32(1), v532)
	mBase = m.M
	v623 = v622
	goto L156
L155:
	;
	v623 = v615
	goto L156
L156:
	;
	v629 = v623
	goto L122
L157:
	;
	v644 = F_pg_sprintf(m, v19, int32(_a_F_interval_out_22), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L25
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v646 = int32(80)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v646)
	v649 = v19 + int32(1)
	if v97 != 0 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	goto L14
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+280)) = int32(89)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+272)) = base.I64_extend_i32_s(v97)
	v657 = F_pg_sprintf(m, v649, int32(_a_F_interval_out_23), v88+int32(272))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L25
	} else {
		goto L164
	}
L162:
	;
	v661 = v649
	goto L163
L163:
	;
	if v96 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v659 = F_strlen(m, v649)
	mBase = m.M
	v661 = v659 + v649
	goto L163
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+264)) = int32(77)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+256)) = base.I64_extend_i32_s(v96)
	v669 = F_pg_sprintf(m, v661, int32(_a_F_interval_out_23), v88+int32(256))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L25
	} else {
		goto L168
	}
L166:
	;
	v673 = v661
	goto L167
L167:
	;
	if v98 != 0 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v671 = F_strlen(m, v661)
	mBase = m.M
	v673 = v671 + v661
	goto L167
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+248)) = int32(68)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+240)) = v99
	v680 = F_pg_sprintf(m, v673, int32(_a_F_interval_out_23), v88+int32(240))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L25
	} else {
		goto L172
	}
L170:
	;
	v684 = v673
	goto L171
L171:
	;
	if v93|(base.B2i32(v95 != int64(0))|v94)|v92 == int32(0) {
		v722 = v684
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v682 = F_strlen(m, v673)
	mBase = m.M
	v684 = v682 + v673
	goto L171
L173:
	;
	v723 = v93 | v92
	if v723 == int32(0) {
		goto L14
	} else {
		goto L181
	}
L174:
	;
	v692 = int32(84)
	*(*uint8)(unsafe.Add(mBase, uint32(v684))) = uint8(v692)
	v695 = v684 + int32(1)
	if v95 != int64(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+232)) = int32(72)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+224)) = v95
	v704 = F_pg_sprintf(m, v695, int32(_a_F_interval_out_23), v88+int32(224))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L25
	} else {
		goto L178
	}
L176:
	;
	v708 = v695
	goto L177
L177:
	;
	if v94 == int32(0) {
		v722 = v708
		goto L173
	} else {
		goto L179
	}
L178:
	;
	v706 = F_strlen(m, v695)
	mBase = m.M
	v708 = v706 + v695
	goto L177
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+216)) = int32(77)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+208)) = base.I64_extend_i32_s(v94)
	v718 = F_pg_sprintf(m, v708, int32(_a_F_interval_out_23), v88+int32(208))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L25
	} else {
		goto L180
	}
L180:
	;
	v720 = F_strlen(m, v708)
	mBase = m.M
	v722 = v720 + v708
	goto L173
L181:
	;
	if v723 < int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v728 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v722))) = uint8(v728)
	v732 = v722 + int32(1)
	goto L184
L183:
	;
	v732 = v722
	goto L184
L184:
	;
	v737 = v93 >> (uint(int32(31)) % 32)
	goto L188
L185:
	;
	v849 = int32(83)
	*(*uint16)(unsafe.Add(mBase, uint32(v848))) = uint16(v849)
	goto L14
L186:
	;
	if v92 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L189
L189:
	;
	v742 = F_pg_ultostr(m, v732, v93^v737-v737)
	mBase = m.M
	goto L186
L190:
	;
	v848 = v742
	goto L185
L191:
	;
	goto L192
L192:
	;
	v746 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v742))) = uint8(v746)
	v749 = v92 >> (uint(int32(31)) % 32)
	v751 = v92 ^ v749 - v749
	v753 = base.I32_div_s(v751, int32(10))
	v756 = v753*int32(-10) + v751
	if v756 != 0 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v766 = base.I32_div_s(v751, int32(100))
	v769 = v766*int32(-10) + v753
	v770 = v756 | v769
	if v770 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L194:
	;
	v758 = v756 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v742)+6)) = uint8(v758)
	v764 = v742 + int32(7)
	goto L193
L195:
	;
	goto L196
L196:
	;
	v764 = v742 + int32(6)
	goto L193
L197:
	;
	v780 = base.I32_div_s(v751, int32(1000))
	v783 = v766 + v780*int32(-10)
	v784 = v770 | v783
	if v784 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L198:
	;
	v778 = v742 + int32(5)
	goto L197
L199:
	;
	goto L200
L200:
	;
	v776 = v769 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v742)+5)) = uint8(v776)
	v778 = v764
	goto L197
L201:
	;
	v794 = base.I32_div_s(v751, int32(_a_F_interval_out_8))
	v797 = v780 + v794*int32(-10)
	v798 = v784 | v797
	if v798 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L202:
	;
	v792 = v742 + int32(4)
	goto L201
L203:
	;
	goto L204
L204:
	;
	v790 = v783 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v742)+4)) = uint8(v790)
	v792 = v778
	goto L201
L205:
	;
	v808 = base.I32_div_s(v751, int32(_a_F_interval_out_9))
	v811 = v794 + v808*int32(-10)
	v812 = v798 | v811
	if v812 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L206:
	;
	v806 = v742 + int32(3)
	goto L205
L207:
	;
	goto L208
L208:
	;
	v804 = v797 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v742)+3)) = uint8(v804)
	v806 = v792
	goto L205
L209:
	;
	v822 = base.I32_div_s(v751, int32(_a_F_interval_out_10))
	v825 = v822*int32(-10) + v808
	if v812|v825 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L210:
	;
	v820 = v742 + int32(2)
	goto L209
L211:
	;
	goto L212
L212:
	;
	v818 = v811 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v742)+2)) = uint8(v818)
	v820 = v806
	goto L209
L213:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v808+int32(9)) {
		goto L217
	} else {
		goto L218
	}
L214:
	;
	v834 = v742 + int32(1)
	goto L213
L215:
	;
	goto L216
L216:
	;
	v832 = v825 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v742)+1)) = uint8(v832)
	v834 = v820
	goto L213
L217:
	;
	v841 = F_pg_ultostr(m, v742+int32(1), v751)
	mBase = m.M
	v842 = v841
	goto L219
L218:
	;
	v842 = v834
	goto L219
L219:
	;
	v848 = v842
	goto L185
L220:
	;
	if v96 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L221:
	;
	v877 = v19
	v878 = int32(0)
	goto L220
L222:
	;
	goto L223
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+400)) = int32(_a_F_interval_out_12)
	if v97 == int32(1) {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v860 = int32(_a_F_interval_out_1)
	goto L226
L225:
	;
	v860 = int32(_a_F_interval_out_2)
	goto L226
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+404)) = v860
	v862 = int32(_a_F_interval_out_1)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+388)) = v862
	*(*int32)(unsafe.Add(mBase, uint32(v88)+384)) = v862
	*(*int64)(unsafe.Add(mBase, uint32(v88)+392)) = base.I64_extend_i32_s(v97)
	v871 = F_pg_sprintf(m, v19, int32(_a_F_interval_out_5), v88+int32(384))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L25
	} else {
		goto L227
	}
L227:
	;
	v875 = F_strlen(m, v19)
	mBase = m.M
	v877 = v875 + v19
	v878 = int32(base.Ui32(v97) >> (uint(int32(31)) % 32))
	goto L220
L228:
	;
	if v98 != 0 {
		goto L245
	} else {
		goto L246
	}
L229:
	;
	v915 = v877
	v916 = base.B2i32(v97 == int32(0))
	v917 = v878
	goto L228
L230:
	;
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+368)) = int32(_a_F_interval_out_11)
	if v96 == int32(1) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v889 = int32(_a_F_interval_out_1)
	goto L234
L233:
	;
	v889 = int32(_a_F_interval_out_2)
	goto L234
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+372)) = v889
	*(*int64)(unsafe.Add(mBase, uint32(v88)+360)) = base.I64_extend_i32_s(v96)
	if v97 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v895 = int32(_a_F_interval_out_3)
	goto L237
L236:
	;
	v895 = int32(_a_F_interval_out_1)
	goto L237
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+352)) = v895
	v898 = int32(_a_F_interval_out_1)
	if v878 != 0 {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v901 = int32(_a_F_interval_out_4)
	goto L240
L239:
	;
	v901 = v898
	goto L240
L240:
	;
	if v96 <= int32(0) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v904 = v898
	goto L243
L242:
	;
	v904 = v901
	goto L243
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+356)) = v904
	v909 = F_pg_sprintf(m, v877, int32(_a_F_interval_out_5), v88+int32(352))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L25
	} else {
		goto L244
	}
L244:
	;
	v913 = F_strlen(m, v877)
	mBase = m.M
	v915 = v913 + v877
	v916 = int32(0)
	v917 = int32(base.Ui32(v96) >> (uint(int32(31)) % 32))
	goto L228
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+336)) = int32(_a_F_interval_out_0)
	if v98 == int32(1) {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	v949 = v915
	v950 = v916
	v951 = v917
	goto L247
L247:
	;
	if v950|base.B2i32(v95 != int64(0))|(v94|v93)|v92 == int32(0) {
		goto L14
	} else {
		goto L261
	}
L248:
	;
	v924 = int32(_a_F_interval_out_1)
	goto L250
L249:
	;
	v924 = int32(_a_F_interval_out_2)
	goto L250
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+340)) = v924
	*(*int64)(unsafe.Add(mBase, uint32(v88)+328)) = v99
	if v916 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v929 = int32(_a_F_interval_out_1)
	goto L253
L252:
	;
	v929 = int32(_a_F_interval_out_3)
	goto L253
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+320)) = v929
	v932 = int32(_a_F_interval_out_1)
	if v917 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v935 = int32(_a_F_interval_out_4)
	goto L256
L255:
	;
	v935 = v932
	goto L256
L256:
	;
	if v98 <= int32(0) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v938 = v932
	goto L259
L258:
	;
	v938 = v935
	goto L259
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+324)) = v938
	v943 = F_pg_sprintf(m, v915, int32(_a_F_interval_out_5), v88+int32(320))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L25
	} else {
		goto L260
	}
L260:
	;
	v947 = F_strlen(m, v915)
	mBase = m.M
	v949 = v947 + v915
	v950 = int32(0)
	v951 = int32(base.Ui32(v98) >> (uint(int32(31)) % 32))
	goto L247
L261:
	;
	v961 = v94 >> (uint(int32(31)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+304)) = v94 ^ v961 - v961
	if v950 != 0 {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v967 = int32(_a_F_interval_out_1)
	goto L264
L263:
	;
	v967 = int32(_a_F_interval_out_3)
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+288)) = v967
	v970 = v95 >> (uint(int64(63)) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+296)) = v95 ^ v970 - v970
	v974 = int32(_a_F_interval_out_6)
	if v951 != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v978 = int32(_a_F_interval_out_4)
	goto L267
L266:
	;
	v978 = int32(_a_F_interval_out_1)
	goto L267
L267:
	;
	if v94|(v93|v92) < int32(0) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v983 = v974
	goto L270
L269:
	;
	v983 = v978
	goto L270
L270:
	;
	if v95 < int64(0) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v986 = v974
	goto L273
L272:
	;
	v986 = v983
	goto L273
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+292)) = v986
	v991 = F_pg_sprintf(m, v949, int32(_a_F_interval_out_7), v88+int32(288))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L25
	} else {
		goto L274
	}
L274:
	;
	v993 = F_strlen(m, v949)
	mBase = m.M
	v999 = v93 >> (uint(int32(31)) % 32)
	goto L277
L275:
	;
	v1111 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1110))) = uint8(v1111)
	goto L14
L276:
	;
	if v92 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L277:
	;
	v1003 = F_pg_ultostr_zeropad(m, v993+v949, v93^v999-v999, int32(2))
	mBase = m.M
	goto L276
L280:
	;
	v1110 = v1003
	goto L275
L281:
	;
	goto L282
L282:
	;
	v1008 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1003))) = uint8(v1008)
	v1011 = v92 >> (uint(int32(31)) % 32)
	v1013 = v92 ^ v1011 - v1011
	v1015 = base.I32_div_s(v1013, int32(10))
	v1018 = v1015*int32(-10) + v1013
	if v1018 != 0 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	v1028 = base.I32_div_s(v1013, int32(100))
	v1031 = v1028*int32(-10) + v1015
	v1032 = v1018 | v1031
	if v1032 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L284:
	;
	v1020 = v1018 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1003)+6)) = uint8(v1020)
	v1026 = v1003 + int32(7)
	goto L283
L285:
	;
	goto L286
L286:
	;
	v1026 = v1003 + int32(6)
	goto L283
L287:
	;
	v1042 = base.I32_div_s(v1013, int32(1000))
	v1045 = v1028 + v1042*int32(-10)
	v1046 = v1032 | v1045
	if v1046 == int32(0) {
		goto L292
	} else {
		goto L293
	}
L288:
	;
	v1040 = v1003 + int32(5)
	goto L287
L289:
	;
	goto L290
L290:
	;
	v1038 = v1031 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1003)+5)) = uint8(v1038)
	v1040 = v1026
	goto L287
L291:
	;
	v1056 = base.I32_div_s(v1013, int32(_a_F_interval_out_8))
	v1059 = v1042 + v1056*int32(-10)
	v1060 = v1046 | v1059
	if v1060 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L292:
	;
	v1054 = v1003 + int32(4)
	goto L291
L293:
	;
	goto L294
L294:
	;
	v1052 = v1045 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1003)+4)) = uint8(v1052)
	v1054 = v1040
	goto L291
L295:
	;
	v1070 = base.I32_div_s(v1013, int32(_a_F_interval_out_9))
	v1073 = v1056 + v1070*int32(-10)
	v1074 = v1060 | v1073
	if v1074 == int32(0) {
		goto L300
	} else {
		goto L301
	}
L296:
	;
	v1068 = v1003 + int32(3)
	goto L295
L297:
	;
	goto L298
L298:
	;
	v1066 = v1059 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1003)+3)) = uint8(v1066)
	v1068 = v1054
	goto L295
L299:
	;
	v1084 = base.I32_div_s(v1013, int32(_a_F_interval_out_10))
	v1087 = v1084*int32(-10) + v1070
	if v1074|v1087 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L300:
	;
	v1082 = v1003 + int32(2)
	goto L299
L301:
	;
	goto L302
L302:
	;
	v1080 = v1073 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1003)+2)) = uint8(v1080)
	v1082 = v1068
	goto L299
L303:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v1070+int32(9)) {
		goto L307
	} else {
		goto L308
	}
L304:
	;
	v1096 = v1003 + int32(1)
	goto L303
L305:
	;
	goto L306
L306:
	;
	v1094 = v1087 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1003)+1)) = uint8(v1094)
	v1096 = v1082
	goto L303
L307:
	;
	v1103 = F_pg_ultostr(m, v1003+int32(1), v1013)
	mBase = m.M
	v1104 = v1103
	goto L309
L308:
	;
	v1104 = v1096
	goto L309
L309:
	;
	v1110 = v1104
	goto L275
L310:
	;
	v1238 = int64(0)
	if v95 != v1238 {
		goto L349
	} else {
		goto L350
	}
L311:
	;
	v1232 = v1230
	v1234 = v1227
	v1237 = int32(0)
	goto L310
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+56)) = int32(_a_F_interval_out_0)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+48)) = v1208
	if v1208 == int64(1) {
		goto L339
	} else {
		goto L340
	}
L313:
	;
	if v1197 != 0 {
		goto L336
	} else {
		goto L337
	}
L314:
	;
	v1185 = int32(0)
	if v98 == v1185 {
		v1232 = v1182
		v1234 = v1183
		v1237 = base.B2i32(v97 == v1185)
		goto L310
	} else {
		goto L334
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+72)) = int32(_a_F_interval_out_11)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+64)) = v1163
	if v1163 == int64(1) {
		goto L329
	} else {
		goto L330
	}
L316:
	;
	v1151 = int32(31)
	v1154 = v96 >> (uint(v1151) % 32)
	v1158 = v1116
	v1160 = int32(base.Ui32(v96) >> (uint(v1151) % 32))
	v1163 = base.I64_extend_i32_u(v96 ^ v1154 - v1154)
	goto L315
L317:
	;
	if v96 != 0 {
		goto L316
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+88)) = int32(_a_F_interval_out_12)
	v1123 = v97 >> (uint(int32(31)) % 32)
	v1125 = v97 ^ v1123 - v1123
	*(*int64)(unsafe.Add(mBase, uint32(v88)+80)) = base.I64_extend_i32_u(v1125)
	if v1125 == int32(1) {
		goto L321
	} else {
		goto L322
	}
L320:
	;
	v1182 = v1116
	v1183 = int32(0)
	goto L314
L321:
	;
	v1132 = int32(_a_F_interval_out_1)
	goto L323
L322:
	;
	v1132 = int32(_a_F_interval_out_2)
	goto L323
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+92)) = v1132
	v1137 = F_pg_sprintf(m, v1116, int32(_a_F_interval_out_13), v88+int32(80))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L25
	} else {
		goto L324
	}
L324:
	;
	v1140 = int32(base.Ui32(v97) >> (uint(int32(31)) % 32))
	v1141 = F_strlen(m, v1116)
	mBase = m.M
	v1142 = v1141 + v1116
	if v96 == int32(0) {
		v1182 = v1142
		v1183 = v1140
		goto L314
	} else {
		goto L325
	}
L325:
	;
	v1146 = base.I64_extend_i32_s(v96)
	if v97 < int32(0) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v1150 = int64(0) - v1146
	goto L328
L327:
	;
	v1150 = v1146
	goto L328
L328:
	;
	v1158 = v1142
	v1160 = v1140
	v1163 = v1150
	goto L315
L329:
	;
	v1171 = int32(_a_F_interval_out_1)
	goto L331
L330:
	;
	v1171 = int32(_a_F_interval_out_2)
	goto L331
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+76)) = v1171
	v1176 = F_pg_sprintf(m, v1158, int32(_a_F_interval_out_13), v88-int32(-64))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L25
	} else {
		goto L332
	}
L332:
	;
	v1178 = F_strlen(m, v1158)
	mBase = m.M
	v1179 = v1178 + v1158
	if v98 == int32(0) {
		v1227 = v1160
		v1230 = v1179
		goto L311
	} else {
		goto L333
	}
L333:
	;
	v1195 = v1179
	v1197 = v1160
	goto L313
L334:
	;
	if v97 != 0 {
		v1195 = v1182
		v1197 = v1183
		goto L313
	} else {
		goto L335
	}
L335:
	;
	v1192 = v99 >> (uint(int64(63)) % 64)
	v1203 = v1182
	v1205 = int32(base.Ui32(v98) >> (uint(int32(31)) % 32))
	v1208 = v99 ^ v1192 - v1192
	goto L312
L336:
	;
	v1202 = int64(0) - v99
	goto L338
L337:
	;
	v1202 = v99
	goto L338
L338:
	;
	v1203 = v1195
	v1205 = v1197
	v1208 = v1202
	goto L312
L339:
	;
	v1216 = int32(_a_F_interval_out_1)
	goto L341
L340:
	;
	v1216 = int32(_a_F_interval_out_2)
	goto L341
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+60)) = v1216
	v1221 = F_pg_sprintf(m, v1203, int32(_a_F_interval_out_13), v88+int32(48))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L25
	} else {
		goto L342
	}
L342:
	;
	v1223 = F_strlen(m, v1203)
	mBase = m.M
	v1227 = v1205
	v1230 = v1223 + v1203
	goto L311
L343:
	;
	if v93|v92 != 0 {
		goto L371
	} else {
		goto L372
	}
L344:
	;
	v1320 = v1318
	v1321 = v1315
	v1322 = int32(0)
	goto L343
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = int32(_a_F_interval_out_16)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v1297
	if v1297 == int64(1) {
		goto L366
	} else {
		goto L367
	}
L346:
	;
	if v1287 != 0 {
		goto L363
	} else {
		goto L364
	}
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+40)) = int32(_a_F_interval_out_14)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+32)) = v1266
	if v1266 == int64(1) {
		goto L358
	} else {
		goto L359
	}
L348:
	;
	if v1234 != 0 {
		goto L355
	} else {
		goto L356
	}
L349:
	;
	if v1237 == int32(0) {
		goto L348
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	if v94 == int32(0) {
		v1320 = v1232
		v1321 = v1234
		v1322 = v1237
		goto L343
	} else {
		goto L353
	}
L352:
	;
	v1243 = int64(63)
	v1247 = v95 >> (uint(v1243) % 64)
	v1264 = base.I32_wrap_i64(int64(base.Ui64(v95) >> (uint(v1243) % 64)))
	v1266 = v95 ^ v1247 - v1247
	goto L347
L353:
	;
	v1252 = base.I64_extend_i32_s(v94)
	if v1237 == int32(0) {
		v1286 = v1232
		v1287 = v1234
		v1290 = v1252
		goto L346
	} else {
		goto L354
	}
L354:
	;
	v1258 = v1252 >> (uint(int64(63)) % 64)
	v1293 = v1232
	v1294 = int32(base.Ui32(v94) >> (uint(int32(31)) % 32))
	v1297 = v1252 ^ v1258 - v1258
	goto L345
L355:
	;
	v1263 = int64(0) - v95
	goto L357
L356:
	;
	v1263 = v95
	goto L357
L357:
	;
	v1264 = v1234
	v1266 = v1263
	goto L347
L358:
	;
	v1274 = int32(_a_F_interval_out_1)
	goto L360
L359:
	;
	v1274 = int32(_a_F_interval_out_2)
	goto L360
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+44)) = v1274
	v1279 = F_pg_sprintf(m, v1232, int32(_a_F_interval_out_13), v88+int32(32))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L25
	} else {
		goto L361
	}
L361:
	;
	v1281 = F_strlen(m, v1232)
	mBase = m.M
	v1282 = v1281 + v1232
	if v94 == int32(0) {
		v1315 = v1264
		v1318 = v1282
		goto L344
	} else {
		goto L362
	}
L362:
	;
	v1286 = v1282
	v1287 = v1264
	v1290 = base.I64_extend_i32_s(v94)
	goto L346
L363:
	;
	v1292 = v1238 - v1290
	goto L365
L364:
	;
	v1292 = v1290
	goto L365
L365:
	;
	v1293 = v1286
	v1294 = v1287
	v1297 = v1292
	goto L345
L366:
	;
	v1305 = int32(_a_F_interval_out_1)
	goto L368
L367:
	;
	v1305 = int32(_a_F_interval_out_2)
	goto L368
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = v1305
	v1310 = F_pg_sprintf(m, v1293, int32(_a_F_interval_out_13), v88+int32(16))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L25
	} else {
		goto L369
	}
L369:
	;
	v1312 = F_strlen(m, v1293)
	mBase = m.M
	v1315 = v1294
	v1318 = v1312 + v1293
	goto L344
L370:
	;
	v1505 = F_strlen(m, v1502)
	mBase = m.M
	v1506 = v1505 + v1502
	v1508 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_out[1])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1506)+4)) = uint8(v1508)
	v1511 = *(*int32)(unsafe.Add(mBase, _c_F_interval_out[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v1506))) = v1511
	goto L14
L371:
	;
	v1326 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1320))) = uint8(v1326)
	v1328 = int32(1)
	v1330 = v1320 + v1328
	v1331 = int32(0)
	if v93|base.B2i32(v1331 <= v92) != 0 {
		goto L376
	} else {
		goto L377
	}
L372:
	;
	goto L373
L373:
	;
	if v1322 != 0 {
		goto L427
	} else {
		goto L428
	}
L374:
	;
	v1363 = v93 >> (uint(int32(31)) % 32)
	goto L387
L375:
	;
	v1352 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1320)+1)) = uint8(v1352)
	v1357 = v1350
	v1358 = v1320 + int32(2)
	goto L374
L376:
	;
	v1338 = base.B2i32(v1331 <= v93)
	goto L378
L377:
	;
	v1338 = v1331
	goto L378
L378:
	;
	if v1338 == int32(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	if (v1321|v1322)&int32(1) == int32(0) {
		v1350 = v1331
		goto L375
	} else {
		goto L382
	}
L380:
	;
	goto L381
L381:
	;
	v1346 = int32(0)
	if v1321 == v1346 {
		v1357 = v1346
		v1358 = v1330
		goto L374
	} else {
		goto L383
	}
L382:
	;
	v1357 = v1328
	v1358 = v1330
	goto L374
L383:
	;
	v1350 = int32(1)
	goto L375
L384:
	;
	v1475 = int32(_a_F_interval_out_2)
	if v92 != 0 {
		goto L419
	} else {
		goto L420
	}
L385:
	;
	if v92 == int32(0) {
		goto L389
	} else {
		goto L390
	}
L387:
	;
	goto L388
L388:
	;
	v1368 = F_pg_ultostr(m, v1358, v93^v1363-v1363)
	mBase = m.M
	goto L385
L389:
	;
	v1474 = v1368
	goto L384
L390:
	;
	goto L391
L391:
	;
	v1372 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1368))) = uint8(v1372)
	v1375 = v92 >> (uint(int32(31)) % 32)
	v1377 = v92 ^ v1375 - v1375
	v1379 = base.I32_div_s(v1377, int32(10))
	v1382 = v1379*int32(-10) + v1377
	if v1382 != 0 {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	v1392 = base.I32_div_s(v1377, int32(100))
	v1395 = v1392*int32(-10) + v1379
	v1396 = v1382 | v1395
	if v1396 == int32(0) {
		goto L397
	} else {
		goto L398
	}
L393:
	;
	v1384 = v1382 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1368)+6)) = uint8(v1384)
	v1390 = v1368 + int32(7)
	goto L392
L394:
	;
	goto L395
L395:
	;
	v1390 = v1368 + int32(6)
	goto L392
L396:
	;
	v1406 = base.I32_div_s(v1377, int32(1000))
	v1409 = v1392 + v1406*int32(-10)
	v1410 = v1396 | v1409
	if v1410 == int32(0) {
		goto L401
	} else {
		goto L402
	}
L397:
	;
	v1404 = v1368 + int32(5)
	goto L396
L398:
	;
	goto L399
L399:
	;
	v1402 = v1395 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1368)+5)) = uint8(v1402)
	v1404 = v1390
	goto L396
L400:
	;
	v1420 = base.I32_div_s(v1377, int32(_a_F_interval_out_8))
	v1423 = v1406 + v1420*int32(-10)
	v1424 = v1410 | v1423
	if v1424 == int32(0) {
		goto L405
	} else {
		goto L406
	}
L401:
	;
	v1418 = v1368 + int32(4)
	goto L400
L402:
	;
	goto L403
L403:
	;
	v1416 = v1409 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1368)+4)) = uint8(v1416)
	v1418 = v1404
	goto L400
L404:
	;
	v1434 = base.I32_div_s(v1377, int32(_a_F_interval_out_9))
	v1437 = v1420 + v1434*int32(-10)
	v1438 = v1424 | v1437
	if v1438 == int32(0) {
		goto L409
	} else {
		goto L410
	}
L405:
	;
	v1432 = v1368 + int32(3)
	goto L404
L406:
	;
	goto L407
L407:
	;
	v1430 = v1423 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1368)+3)) = uint8(v1430)
	v1432 = v1418
	goto L404
L408:
	;
	v1448 = base.I32_div_s(v1377, int32(_a_F_interval_out_10))
	v1451 = v1448*int32(-10) + v1434
	if v1438|v1451 == int32(0) {
		goto L413
	} else {
		goto L414
	}
L409:
	;
	v1446 = v1368 + int32(2)
	goto L408
L410:
	;
	goto L411
L411:
	;
	v1444 = v1437 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1368)+2)) = uint8(v1444)
	v1446 = v1432
	goto L408
L412:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v1434+int32(9)) {
		goto L416
	} else {
		goto L417
	}
L413:
	;
	v1460 = v1368 + int32(1)
	goto L412
L414:
	;
	goto L415
L415:
	;
	v1458 = v1451 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1368)+1)) = uint8(v1458)
	v1460 = v1446
	goto L412
L416:
	;
	v1467 = F_pg_ultostr(m, v1368+int32(1), v1377)
	mBase = m.M
	v1468 = v1467
	goto L418
L417:
	;
	v1468 = v1460
	goto L418
L418:
	;
	v1474 = v1468
	goto L384
L419:
	;
	v1478 = v1475
	goto L421
L420:
	;
	v1478 = int32(_a_F_interval_out_1)
	goto L421
L421:
	;
	v1480 = v93 >> (uint(int32(31)) % 32)
	if v93^v1480-v1480 != int32(1) {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v1485 = v1475
	goto L424
L423:
	;
	v1485 = v1478
	goto L424
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v1485
	v1488 = F_pg_sprintf(m, v1474, int32(_a_F_interval_out_15), v88)
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L25
	} else {
		goto L425
	}
L425:
	;
	if v1357 != 0 {
		v1502 = v1474
		goto L370
	} else {
		goto L426
	}
L426:
	;
	goto L14
L427:
	;
	v1490 = F_strlen(m, v1320)
	mBase = m.M
	v1491 = v1490 + v1320
	v1493 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_interval_out[3])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1491)+2)) = uint8(v1493)
	v1496 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_interval_out[4])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1491))) = uint16(v1496)
	goto L429
L428:
	;
	goto L429
L429:
	;
	if v1321 == int32(0) {
		goto L14
	} else {
		goto L430
	}
L430:
	;
	v1502 = v1320
	goto L370
L431:
	;
	m.G0 = v19 + int32(176)
	return v1544
}
func F_interval_pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		if v12 != int32(2147483647) {
			if v12 != int32(-2147483648) {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				if v57 != int32(2147483647) {
					if v57 != int32(-2147483648) {
						F_finite_interval_pl(m, v6, v5, v8)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v62 != int32(-2147483648) {
							F_finite_interval_pl(m, v6, v5, v8)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v65 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v65 == int64(-9223372036854775807-1) {
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v74
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v76
								return v8
							} else {
								F_finite_interval_pl(m, v6, v5, v8)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									return v8
								}
							}
						}
					}
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
					if v68 != int32(2147483647) {
						F_finite_interval_pl(m, v6, v5, v8)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						v71 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
						if v71 != int64(9223372036854775807) {
							F_finite_interval_pl(m, v6, v5, v8)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v74
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v76
							return v8
						}
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				if v17 != int32(-2147483648) {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					if v57 != int32(2147483647) {
						if v57 != int32(-2147483648) {
							F_finite_interval_pl(m, v6, v5, v8)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
							if v62 != int32(-2147483648) {
								F_finite_interval_pl(m, v6, v5, v8)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v65 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								if v65 == int64(-9223372036854775807-1) {
									v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v74
									v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
									*(*int64)(unsafe.Add(mBase, uint32(v8))) = v76
									return v8
								} else {
									F_finite_interval_pl(m, v6, v5, v8)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										return v8
									}
								}
							}
						}
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v68 != int32(2147483647) {
							F_finite_interval_pl(m, v6, v5, v8)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v71 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v71 != int64(9223372036854775807) {
								F_finite_interval_pl(m, v6, v5, v8)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v74
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v76
								return v8
							}
						}
					}
				} else {
					v20 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
					if v20 != int64(-9223372036854775807-1) {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
						if v57 != int32(2147483647) {
							if v57 != int32(-2147483648) {
								F_finite_interval_pl(m, v6, v5, v8)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
								if v62 != int32(-2147483648) {
									F_finite_interval_pl(m, v6, v5, v8)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										return v8
									}
								} else {
									v65 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
									if v65 == int64(-9223372036854775807-1) {
										v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v74
										v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
										*(*int64)(unsafe.Add(mBase, uint32(v8))) = v76
										return v8
									} else {
										F_finite_interval_pl(m, v6, v5, v8)
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											return v8
										}
									}
								}
							}
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
							if v68 != int32(2147483647) {
								F_finite_interval_pl(m, v6, v5, v8)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v71 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								if v71 != int64(9223372036854775807) {
									F_finite_interval_pl(m, v6, v5, v8)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										return v8
									}
								} else {
									v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v74
									v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
									*(*int64)(unsafe.Add(mBase, uint32(v8))) = v76
									return v8
								}
							}
						}
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
						if v23 != int32(2147483647) {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
							return v8
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
							if v26 != int32(2147483647) {
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
								return v8
							} else {
								v29 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								if v29 == int64(9223372036854775807) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_interval_pl_0), int32(0))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_interval_pl_1), int32(3540), int32(_a_F_interval_pl_2))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(-9223372034707292160)
									*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(-9223372036854775807 - 1)
									return v8
								}
							}
						}
					}
				}
			}
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if v37 != int32(2147483647) {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				if v57 != int32(2147483647) {
					if v57 != int32(-2147483648) {
						F_finite_interval_pl(m, v6, v5, v8)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v62 != int32(-2147483648) {
							F_finite_interval_pl(m, v6, v5, v8)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v65 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v65 == int64(-9223372036854775807-1) {
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v74
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v76
								return v8
							} else {
								F_finite_interval_pl(m, v6, v5, v8)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									return v8
								}
							}
						}
					}
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
					if v68 != int32(2147483647) {
						F_finite_interval_pl(m, v6, v5, v8)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						v71 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
						if v71 != int64(9223372036854775807) {
							F_finite_interval_pl(m, v6, v5, v8)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v74
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v76
							return v8
						}
					}
				}
			} else {
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
				if v40 != int64(9223372036854775807) {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					if v57 != int32(2147483647) {
						if v57 != int32(-2147483648) {
							F_finite_interval_pl(m, v6, v5, v8)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
							if v62 != int32(-2147483648) {
								F_finite_interval_pl(m, v6, v5, v8)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v65 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								if v65 == int64(-9223372036854775807-1) {
									v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v74
									v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
									*(*int64)(unsafe.Add(mBase, uint32(v8))) = v76
									return v8
								} else {
									F_finite_interval_pl(m, v6, v5, v8)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										return v8
									}
								}
							}
						}
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v68 != int32(2147483647) {
							F_finite_interval_pl(m, v6, v5, v8)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v71 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v71 != int64(9223372036854775807) {
								F_finite_interval_pl(m, v6, v5, v8)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v74
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v76
								return v8
							}
						}
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					if v43 != int32(-2147483648) {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
						return v8
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v46 != int32(-2147483648) {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
							return v8
						} else {
							v49 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v49 == int64(-9223372036854775807-1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_interval_pl_0), int32(0))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_interval_pl_1), int32(3549), int32(_a_F_interval_pl_2))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
								return v8
							}
						}
					}
				}
			}
		}
	}
}
func F_interval_smaller(m *base.Module, l0 int32) int32 {
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
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	m.G0 = v14 + int32(32)
	v119 = v113 + v115
	v120 = v110 + v112
	v124 = int64(63)
	v127 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v115))) + (v114 + v113>>(uint(v124)%64))
	v133 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v112))) + (v111 + v110>>(uint(v124)%64))
	if v133 == v127 {
		v136 = base.B2i32(base.Ui64(v119) < base.Ui64(v120))
	} else {
		v136 = base.B2i32(v127 < v133)
	}
	if v136 != 0 {
		v137 = v18
	} else {
		v137 = v64
	}
	return v137
}
func F_make_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v69 int64
	_ = v69
	var v72 float64
	_ = v72
	var v76 float64
	_ = v76
	var v81 float64
	_ = v81
	var v84 int32
	_ = v84
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v100 int32
	_ = v100
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	v14 = base.F64_abs(v13)
	if base.F64_eq(v14, math.Float64frombits(uint64(0x7ff0000000000000)))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v14))) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v126 = m.ExcPending
		if v126 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v129 = m.ExcPending
			if v129 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_make_interval_0), int32(0))
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_make_interval_1), int32(1578), int32(_a_F_make_interval_2))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+60)))
		v22 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+52)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v26 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
		v28 = F_palloc(m, int32(16))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v33 = v26 * int64(12)
			v34 = base.I32_wrap_i64(v33)
			*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v34
			if base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(32))%64))) != v34>>(uint(int32(31))%32) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_make_interval_0), int32(0))
						mBase = m.M
						v133 = m.ExcPending
						if v133 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_make_interval_1), int32(1578), int32(_a_F_make_interval_2))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v42 = v34 + v25
				*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v42
				if base.B2i32(v25 < int32(0))^base.B2i32(v42 < v34) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_make_interval_0), int32(0))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_make_interval_1), int32(1578), int32(_a_F_make_interval_2))
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v50 = base.I64_extend_i32_s(v24) * int64(7)
					v51 = base.I32_wrap_i64(v50)
					*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v51
					if base.I32_wrap_i64(int64(base.Ui64(v50)>>(uint(int64(32))%64))) != v51>>(uint(int32(31))%32) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_make_interval_0), int32(0))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_make_interval_1), int32(1578), int32(_a_F_make_interval_2))
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v59 = v51 + v23
						*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v59
						if base.B2i32(v23 < int32(0))^base.B2i32(v59 < v51) != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_make_interval_0), int32(0))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_make_interval_1), int32(1578), int32(_a_F_make_interval_2))
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v69 = v21*int64(60000000) + v22*int64(3600000000)
							*(*int64)(unsafe.Add(mBase, uint32(v28))) = v69
							v72 = base.F64_mul(v13, float64(1e+06))
							if base.F64_eq(base.F64_abs(v72), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v76 = float64(0)
								if base.F64_eq(v72, v76)&base.F64_ne(v13, v76) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v81 = base.F64_nearest(v72)
									v84 = int32(0)
									if base.B2i32(base.F64_ge(v81, float64(-9.223372036854776e+18)) == v84)|base.B2i32(base.F64_lt(v81, float64(9.223372036854776e+18)) == v84) != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_make_interval_0), int32(0))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_make_interval_1), int32(1578), int32(_a_F_make_interval_2))
													mBase = m.M
													v138 = m.ExcPending
													if v138 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v91 = base.I64_trunc_sat_f64_s(v81)
										v92 = v69 + v91
										*(*int64)(unsafe.Add(mBase, uint32(v28))) = v92
										if base.B2i32(v91 < int64(0))^base.B2i32(v92 < v69) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_make_interval_0), int32(0))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_make_interval_1), int32(1578), int32(_a_F_make_interval_2))
														mBase = m.M
														v138 = m.ExcPending
														if v138 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											if v42 != int32(2147483647) {
												v100 = int32(-2147483648)
												if base.B2i32(v42 != v100)|base.B2i32(v59 != v100) != 0 {
													return v28
												} else {
													if v92 == int64(-9223372036854775807-1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(134217858))
															mBase = m.M
															v129 = m.ExcPending
															if v129 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_make_interval_0), int32(0))
																mBase = m.M
																v133 = m.ExcPending
																if v133 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_make_interval_1), int32(1578), int32(_a_F_make_interval_2))
																	mBase = m.M
																	v138 = m.ExcPending
																	if v138 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														return v28
													}
												}
											} else {
												if base.B2i32(v59 != int32(2147483647))|base.B2i32(v92 != int64(9223372036854775807)) != 0 {
													return v28
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v129 = m.ExcPending
														if v129 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_make_interval_0), int32(0))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_make_interval_1), int32(1578), int32(_a_F_make_interval_2))
																mBase = m.M
																v138 = m.ExcPending
																if v138 != 0 {
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
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
