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
	return base.B2i32(v119^v120|(base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v112)))+(v111+v110>>(uint(v124)%64))^(base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v115)))+(v114+v113>>(uint(v124)%64)))) == int64(0))
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
	v22 = F_DirectFunctionCall1Coll(m, int32(1285), int32(0), v5+int32(8))
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
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
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
						F_finite_interval_mi(m, v6, v5, v8)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v62 != int32(-2147483648) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v65 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v65 != int64(-9223372036854775807-1) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
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
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
					if v73 != int32(2147483647) {
						F_finite_interval_mi(m, v6, v5, v8)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
						if v76 != int64(9223372036854775807) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
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
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					if v57 != int32(2147483647) {
						if v57 != int32(-2147483648) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
							if v62 != int32(-2147483648) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v65 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								if v65 != int64(-9223372036854775807-1) {
									F_finite_interval_mi(m, v6, v5, v8)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
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
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v73 != int32(2147483647) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v76 != int64(9223372036854775807) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
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
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
						if v57 != int32(2147483647) {
							if v57 != int32(-2147483648) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
								if v62 != int32(-2147483648) {
									F_finite_interval_mi(m, v6, v5, v8)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										return v8
									}
								} else {
									v65 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
									if v65 != int64(-9223372036854775807-1) {
										F_finite_interval_mi(m, v6, v5, v8)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
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
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
							if v73 != int32(2147483647) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								if v76 != int64(9223372036854775807) {
									F_finite_interval_mi(m, v6, v5, v8)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
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
									v90 = m.ExcPending
									if v90 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(398189), int32(0))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(490442), int32(3596), int32(316518))
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
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
						F_finite_interval_mi(m, v6, v5, v8)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v62 != int32(-2147483648) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v65 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v65 != int64(-9223372036854775807-1) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
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
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
					if v73 != int32(2147483647) {
						F_finite_interval_mi(m, v6, v5, v8)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							return v8
						}
					} else {
						v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
						if v76 != int64(9223372036854775807) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
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
				v40 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
				if v40 != int64(9223372036854775807) {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					if v57 != int32(2147483647) {
						if v57 != int32(-2147483648) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
							if v62 != int32(-2147483648) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									return v8
								}
							} else {
								v65 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								if v65 != int64(-9223372036854775807-1) {
									F_finite_interval_mi(m, v6, v5, v8)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
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
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v73 != int32(2147483647) {
							F_finite_interval_mi(m, v6, v5, v8)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								return v8
							}
						} else {
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v76 != int64(9223372036854775807) {
								F_finite_interval_mi(m, v6, v5, v8)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
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
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					if v43 != int32(2147483647) {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
						return v8
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						if v46 != int32(2147483647) {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(9223372034707292159)
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(9223372036854775807)
							return v8
						} else {
							v49 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							if v49 == int64(9223372036854775807) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(398189), int32(0))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(490442), int32(3605), int32(316518))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
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
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int64
	_ = v169
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v196 int64
	_ = v196
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v1002 int32
	_ = v1002
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1182 int32
	_ = v1182
	var v1188 int32
	_ = v1188
	var v1191 int64
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1270 int32
	_ = v1270
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1411 int32
	_ = v1411
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1427 int32
	_ = v1427
	var v1432 int32
	_ = v1432
	var v1436 int32
	_ = v1436
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1462 int32
	_ = v1462
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1481 int64
	_ = v1481
	var v1485 int64
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1497 int64
	_ = v1497
	var v1505 int32
	_ = v1505
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1519 int32
	_ = v1519
	var v1524 int32
	_ = v1524
	var v1528 int32
	_ = v1528
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1545 int32
	_ = v1545
	var v1548 int32
	_ = v1548
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1581 int64
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1590 int64
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1595 int64
	_ = v1595
	var v1603 int32
	_ = v1603
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1617 int32
	_ = v1617
	var v1622 int32
	_ = v1622
	var v1626 int32
	_ = v1626
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1681 int64
	_ = v1681
	var v1686 int64
	_ = v1686
	var v1690 int64
	_ = v1690
	var v1695 int64
	_ = v1695
	var v1701 int64
	_ = v1701
	var v1706 int64
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1709 int64
	_ = v1709
	var v1717 int32
	_ = v1717
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1740 int32
	_ = v1740
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1751 int32
	_ = v1751
	var v1757 int32
	_ = v1757
	var v1760 int32
	_ = v1760
	var v1766 int32
	_ = v1766
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1789 int64
	_ = v1789
	var v1791 int64
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1796 int64
	_ = v1796
	var v1804 int32
	_ = v1804
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1818 int32
	_ = v1818
	var v1823 int32
	_ = v1823
	var v1827 int32
	_ = v1827
	var v1832 int32
	_ = v1832
	var v1834 int32
	_ = v1834
	var v1838 int32
	_ = v1838
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1914 int32
	_ = v1914
	var v1919 int32
	_ = v1919
	var v1923 int32
	_ = v1923
	var v1926 int32
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1953 int32
	_ = v1953
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1967 int32
	_ = v1967
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1981 int32
	_ = v1981
	var v1983 int32
	_ = v1983
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1995 int32
	_ = v1995
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2009 int32
	_ = v2009
	var v2011 int32
	_ = v2011
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2030 int32
	_ = v2030
	var v2032 int32
	_ = v2032
	var v2037 int32
	_ = v2037
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2049 int32
	_ = v2049
	var v2054 int32
	_ = v2054
	var v2058 int32
	_ = v2058
	var v2063 int32
	_ = v2063
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2075 int32
	_ = v2075
	var v2078 int32
	_ = v2078
	var v2084 int32
	_ = v2084
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2110 int32
	_ = v2110
	var v2121 int32
	_ = v2121
	var v2126 int32
	_ = v2126
	var v2130 int32
	_ = v2130
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2141 int32
	_ = v2141
	var v2147 int32
	_ = v2147
	var v2150 int32
	_ = v2150
	var v2156 int32
	_ = v2156
	var v2160 int32
	_ = v2160
	var v2162 int32
	_ = v2162
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2176 int32
	_ = v2176
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2184 int64
	_ = v2184
	var v2185 int64
	_ = v2185
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2208 int32
	_ = v2208
	var v2213 int32
	_ = v2213
	var v2217 int32
	_ = v2217
	var v2222 int32
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2228 int32
	_ = v2228
	var v2234 int32
	_ = v2234
	var v2237 int32
	_ = v2237
	var v2243 int32
	_ = v2243
	var v2247 int32
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2257 int32
	_ = v2257
	var v2264 int32
	_ = v2264
	var v2268 int32
	_ = v2268
	var v2273 int32
	_ = v2273
	var v2276 int32
	_ = v2276
	var v2278 int32
	_ = v2278
	var v2280 int32
	_ = v2280
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2291 int32
	_ = v2291
	var v2293 int32
	_ = v2293
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2303 int32
	_ = v2303
	var v2305 int32
	_ = v2305
	var v2307 int32
	_ = v2307
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2317 int32
	_ = v2317
	var v2319 int32
	_ = v2319
	var v2321 int32
	_ = v2321
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2331 int32
	_ = v2331
	var v2333 int32
	_ = v2333
	var v2335 int32
	_ = v2335
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2352 int32
	_ = v2352
	var v2359 int32
	_ = v2359
	var v2361 int32
	_ = v2361
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2393 int32
	_ = v2393
	var v2398 int32
	_ = v2398
	var v2402 int32
	_ = v2402
	var v2407 int32
	_ = v2407
	var v2409 int32
	_ = v2409
	var v2413 int32
	_ = v2413
	var v2419 int32
	_ = v2419
	var v2422 int32
	_ = v2422
	var v2428 int32
	_ = v2428
	var v2432 int32
	_ = v2432
	var v2434 int32
	_ = v2434
	var v2442 int32
	_ = v2442
	var v2449 int32
	_ = v2449
	var v2453 int32
	_ = v2453
	var v2458 int32
	_ = v2458
	var v2461 int32
	_ = v2461
	var v2463 int32
	_ = v2463
	var v2465 int32
	_ = v2465
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2476 int32
	_ = v2476
	var v2478 int32
	_ = v2478
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2488 int32
	_ = v2488
	var v2490 int32
	_ = v2490
	var v2492 int32
	_ = v2492
	var v2495 int32
	_ = v2495
	var v2496 int32
	_ = v2496
	var v2502 int32
	_ = v2502
	var v2504 int32
	_ = v2504
	var v2506 int32
	_ = v2506
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2516 int32
	_ = v2516
	var v2518 int32
	_ = v2518
	var v2520 int32
	_ = v2520
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2530 int32
	_ = v2530
	var v2532 int32
	_ = v2532
	var v2534 int32
	_ = v2534
	var v2537 int32
	_ = v2537
	var v2544 int32
	_ = v2544
	var v2546 int32
	_ = v2546
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	v18 = m.G0
	v20 = v18 - int32(176)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v23 != int32(-2147483648) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v2598 = F_pstrdup(m, v20)
	mBase = m.M
	v2599 = m.ExcPending
	if v2599 != 0 {
		goto L28
	} else {
		goto L769
	}
L2:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v55
	v58 = int32(12)
	v59 = base.I32_div_s(v23, v58)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+168)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v20)+164)) = v23 - v59*v58
	v66 = base.I64_div_s(v56, int64(3600000000))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+152)) = v66
	v70 = v56 + v66*int64(-3600000000)
	v72 = base.I64_div_s(v70, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+144)) = uint32(v72)
	v76 = v72*int64(-60000000) + v70
	v78 = base.I64_div_s(v76, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+140)) = uint32(v78)
	v82 = v78*int64(4293967296) + v76
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+136)) = uint32(v82)
	v85 = *(*int32)(unsafe.Add(mBase, _consts[854]))
	v86 = m.G0
	v88 = v86 - int32(416)
	m.G0 = v88
	v91 = v20 + int32(136)
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
		goto L17
	default:
		goto L16
	case 2:
		goto L19
	case 3:
		goto L18
	}
L3:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _consts[855])))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v50)
	v53 = *(*int64)(unsafe.Add(mBase, _consts[856]))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v53
	goto L1
L4:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, _consts[857])))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+8)) = uint16(v44)
	v47 = *(*int64)(unsafe.Add(mBase, _consts[858]))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v47
	goto L1
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v36 != int32(2147483647) {
		v55 = v36
		goto L2
	} else {
		goto L12
	}
L6:
	;
	if v23 == int32(2147483647) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v29 != int32(-2147483648) {
		v55 = v29
		goto L2
	} else {
		goto L10
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v55 = v28
	goto L2
L10:
	;
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	if v32 == int64(-9223372036854775807-1) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v55 = int32(-2147483648)
	goto L2
L12:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	if v39 == int64(9223372036854775807) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v55 = int32(2147483647)
	goto L2
L14:
	;
	m.G0 = v88 + int32(416)
	goto L1
L15:
	;
	if v100 != v102 {
		goto L656
	} else {
		goto L657
	}
L16:
	;
	v1392 = int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v20))) = uint16(v1392)
	v1395 = v20 + int32(1)
	if v97 == int32(0) {
		goto L423
	} else {
		goto L424
	}
L17:
	;
	if v97 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L18:
	;
	if v97 != 0 {
		goto L96
	} else {
		goto L97
	}
L19:
	;
	v100 = v96 | v97
	v102 = int32(0)
	v104 = int64(0)
	v115 = base.B2i32(v102 <= v100|v98) & base.B2i32(v104 <= v95) & base.B2i32(v102 <= v94) & base.B2i32(v102 <= v93) & base.B2i32(v102 <= v92)
	v135 = base.B2i32(v97 <= v102) & base.B2i32(v96 <= v102) & base.B2i32(v98 <= v102) & base.B2i32(v95 <= v104) & base.B2i32(v94 <= v102) & base.B2i32(v93 <= v102) & base.B2i32(v92 <= v102)
	v138 = base.B2i32(v100 != v102)
	v156 = (v115 | v135) & base.B2i32(v138&(base.B2i32(v98 != v102)|base.B2i32(v95 != v104)|base.B2i32(v94 != v102)|base.B2i32(v93 != v102)|base.B2i32(v92 != v102)) == v102)
	if v115 == v102 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v186 = v94 >> (uint(int32(31)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+200)) = v94 ^ v186 - v186
	v190 = int64(63)
	v191 = v95 >> (uint(v190) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+192)) = v95 ^ v191 - v191
	v196 = v99 >> (uint(v190) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+176)) = v99 ^ v196 - v196
	v200 = int32(45)
	if v94|(v93|v92) < int32(0) {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	if v156 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v135 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v161 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v161)
	v163 = int32(0)
	v169 = int64(0)
	v2178 = v163 - v96
	v2179 = v20 + int32(1)
	v2180 = v163 - v97
	v2181 = v163 - v94
	v2182 = v163 - v93
	v2183 = v163 - v92
	v2184 = v169 - v95
	v2185 = v169 - v99
	goto L15
L25:
	;
	v181 = F_pg_sprintf(m, v20, int32(549266), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if v156 != 0 {
		v2178 = v96
		v2179 = v20
		v2180 = v97
		v2181 = v94
		v2182 = v93
		v2183 = v92
		v2184 = v95
		v2185 = v99
		goto L15
	} else {
		goto L30
	}
L28:
	;
	return int32(0)
L29:
	;
	goto L14
L30:
	;
	goto L20
L31:
	;
	v207 = v200
	goto L33
L32:
	;
	v207 = int32(43)
	goto L33
L33:
	;
	if v95 < int64(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v210 = v200
	goto L36
L35:
	;
	v210 = v207
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+184)) = v210
	if v98 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v216 = int32(45)
	goto L39
L38:
	;
	v216 = int32(43)
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+172)) = v216
	v218 = int32(31)
	v219 = v96 >> (uint(v218) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+168)) = v96 ^ v219 - v219
	v224 = v97 >> (uint(v218) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+164)) = v97 ^ v224 - v224
	if v100 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v232 = int32(45)
	goto L42
L41:
	;
	v232 = int32(43)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+160)) = v232
	v237 = F_pg_sprintf(m, v20, int32(541290), v88+int32(160))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L28
	} else {
		goto L43
	}
L43:
	;
	if v20&int32(3) == int32(0) {
		v262 = v20
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v302 = v93 >> (uint(int32(31)) % 32)
	goto L63
L45:
	;
	v295 = v287 - v20
	goto L44
L46:
	;
	v266 = v262
	goto L55
L47:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v246 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v295 = int32(0)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v251 = v20
	goto L51
L51:
	;
	v255 = v251 + int32(1)
	if v255&int32(3) == int32(0) {
		v262 = v255
		goto L46
	} else {
		goto L53
	}
L52:
	;
	v287 = v255
	goto L45
L53:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if v260 != 0 {
		v251 = v255
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v275 = int32(-2139062144)
	if (int32(16843008)-v272|v272)&v275 == v275 {
		v266 = v266 + int32(4)
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v281 = v266
	goto L58
L57:
	;
	goto L56
L58:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v285 != 0 {
		v281 = v281 + int32(1)
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v287 = v281
	goto L45
L60:
	;
	goto L59
L61:
	;
	v415 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v414))) = uint8(v415)
	goto L14
L62:
	;
	if v92 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v306 = F_pg_ultostr_zeropad(m, v295+v20, v93^v302-v302, int32(2))
	mBase = m.M
	goto L62
L66:
	;
	v414 = v306
	goto L61
L67:
	;
	goto L68
L68:
	;
	v311 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v306))) = uint8(v311)
	v314 = v92 >> (uint(int32(31)) % 32)
	v316 = v92 ^ v314 - v314
	v318 = base.I32_div_s(v316, int32(10))
	v321 = v318*int32(-10) + v316
	if v321 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v331 = base.I32_div_s(v316, int32(100))
	v334 = v331*int32(-10) + v318
	v335 = v321 | v334
	if v335 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v323 = v321 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v306)+6)) = uint8(v323)
	v329 = v306 + int32(7)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v329 = v306 + int32(6)
	goto L69
L73:
	;
	v345 = base.I32_div_s(v316, int32(1000))
	v348 = v345*int32(-10) + v331
	v349 = v335 | v348
	if v349 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v343 = v306 + int32(5)
	goto L73
L75:
	;
	goto L76
L76:
	;
	v341 = v334 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v306)+5)) = uint8(v341)
	v343 = v329
	goto L73
L77:
	;
	v359 = base.I32_div_s(v316, int32(10000))
	v362 = v359*int32(-10) + v345
	v363 = v349 | v362
	if v363 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	v357 = v306 + int32(4)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v355 = v348 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v306)+4)) = uint8(v355)
	v357 = v343
	goto L77
L81:
	;
	v373 = base.I32_div_s(v316, int32(100000))
	v376 = v373*int32(-10) + v359
	v377 = v363 | v376
	if v377 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L82:
	;
	v371 = v306 + int32(3)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v369 = v362 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v306)+3)) = uint8(v369)
	v371 = v357
	goto L81
L85:
	;
	v387 = base.I32_div_s(v316, int32(1000000))
	v390 = v387*int32(-10) + v373
	if v377|v390 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L86:
	;
	v385 = v306 + int32(2)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v383 = v376 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v306)+2)) = uint8(v383)
	v385 = v371
	goto L85
L89:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v373+int32(9)) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v399 = v306 + int32(1)
	goto L89
L91:
	;
	goto L92
L92:
	;
	v397 = v390 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v306)+1)) = uint8(v397)
	v399 = v385
	goto L89
L93:
	;
	v406 = F_pg_ultostr(m, v306+int32(1), v316)
	mBase = m.M
	v407 = v406
	goto L95
L94:
	;
	v407 = v399
	goto L95
L95:
	;
	v414 = v407
	goto L61
L96:
	;
	v423 = int32(80)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v423)
	v426 = v20 + int32(1)
	if v97 != 0 {
		goto L105
	} else {
		goto L106
	}
L97:
	;
	if v96 != 0 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	if v98 != 0 {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	if v95 != int64(0) {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	if v94 != 0 {
		goto L96
	} else {
		goto L101
	}
L101:
	;
	if v93 != 0 {
		goto L96
	} else {
		goto L102
	}
L102:
	;
	if v92 != 0 {
		goto L96
	} else {
		goto L103
	}
L103:
	;
	v421 = F_pg_sprintf(m, v20, int32(518485), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L28
	} else {
		goto L104
	}
L104:
	;
	goto L14
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+280)) = int32(89)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+272)) = base.I64_extend_i32_s(v97)
	v434 = F_pg_sprintf(m, v426, int32(494762), v88+int32(272))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L28
	} else {
		goto L108
	}
L106:
	;
	v494 = v426
	goto L107
L107:
	;
	if v96 != 0 {
		goto L126
	} else {
		goto L127
	}
L108:
	;
	if v426&int32(3) == int32(0) {
		v459 = v426
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v494 = v492 + v426
	goto L107
L110:
	;
	v492 = v484 - v426
	goto L109
L111:
	;
	v463 = v459
	goto L120
L112:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	if v443 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v492 = int32(0)
	goto L109
L114:
	;
	goto L115
L115:
	;
	v448 = v426
	goto L116
L116:
	;
	v452 = v448 + int32(1)
	if v452&int32(3) == int32(0) {
		v459 = v452
		goto L111
	} else {
		goto L118
	}
L117:
	;
	v484 = v452
	goto L110
L118:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
	if v457 != 0 {
		v448 = v452
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v463)))
	v472 = int32(-2139062144)
	if (int32(16843008)-v469|v469)&v472 == v472 {
		v463 = v463 + int32(4)
		goto L120
	} else {
		goto L122
	}
L121:
	;
	v478 = v463
	goto L123
L122:
	;
	goto L121
L123:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	if v482 != 0 {
		v478 = v478 + int32(1)
		goto L123
	} else {
		goto L125
	}
L124:
	;
	v484 = v478
	goto L110
L125:
	;
	goto L124
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+264)) = int32(77)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+256)) = base.I64_extend_i32_s(v96)
	v502 = F_pg_sprintf(m, v494, int32(494762), v88+int32(256))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L28
	} else {
		goto L129
	}
L127:
	;
	v562 = v494
	goto L128
L128:
	;
	if v98 != 0 {
		goto L147
	} else {
		goto L148
	}
L129:
	;
	if v494&int32(3) == int32(0) {
		v527 = v494
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v562 = v560 + v494
	goto L128
L131:
	;
	v560 = v552 - v494
	goto L130
L132:
	;
	v531 = v527
	goto L141
L133:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494))))
	if v511 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v560 = int32(0)
	goto L130
L135:
	;
	goto L136
L136:
	;
	v516 = v494
	goto L137
L137:
	;
	v520 = v516 + int32(1)
	if v520&int32(3) == int32(0) {
		v527 = v520
		goto L132
	} else {
		goto L139
	}
L138:
	;
	v552 = v520
	goto L131
L139:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	if v525 != 0 {
		v516 = v520
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	v540 = int32(-2139062144)
	if (int32(16843008)-v537|v537)&v540 == v540 {
		v531 = v531 + int32(4)
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v546 = v531
	goto L144
L143:
	;
	goto L142
L144:
	;
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546))))
	if v550 != 0 {
		v546 = v546 + int32(1)
		goto L144
	} else {
		goto L146
	}
L145:
	;
	v552 = v546
	goto L131
L146:
	;
	goto L145
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+248)) = int32(68)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+240)) = v99
	v569 = F_pg_sprintf(m, v562, int32(494762), v88+int32(240))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L28
	} else {
		goto L150
	}
L148:
	;
	v629 = v562
	goto L149
L149:
	;
	if v95 != int64(0) {
		goto L169
	} else {
		goto L170
	}
L150:
	;
	if v562&int32(3) == int32(0) {
		v594 = v562
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v629 = v627 + v562
	goto L149
L152:
	;
	v627 = v619 - v562
	goto L151
L153:
	;
	v598 = v594
	goto L162
L154:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562))))
	if v578 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v627 = int32(0)
	goto L151
L156:
	;
	goto L157
L157:
	;
	v583 = v562
	goto L158
L158:
	;
	v587 = v583 + int32(1)
	if v587&int32(3) == int32(0) {
		v594 = v587
		goto L153
	} else {
		goto L160
	}
L159:
	;
	v619 = v587
	goto L152
L160:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	if v592 != 0 {
		v583 = v587
		goto L158
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v598)))
	v607 = int32(-2139062144)
	if (int32(16843008)-v604|v604)&v607 == v607 {
		v598 = v598 + int32(4)
		goto L162
	} else {
		goto L164
	}
L163:
	;
	v613 = v598
	goto L165
L164:
	;
	goto L163
L165:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	if v617 != 0 {
		v613 = v613 + int32(1)
		goto L165
	} else {
		goto L167
	}
L166:
	;
	v619 = v613
	goto L152
L167:
	;
	goto L166
L168:
	;
	if v93|v92 == int32(0) {
		goto L14
	} else {
		goto L214
	}
L169:
	;
	v634 = int32(84)
	*(*uint8)(unsafe.Add(mBase, uint32(v629))) = uint8(v634)
	v637 = v629 + int32(1)
	if v95 != int64(0) {
		goto L174
	} else {
		goto L175
	}
L170:
	;
	if v94 != 0 {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	if v93 != 0 {
		goto L169
	} else {
		goto L172
	}
L172:
	;
	if v92 == int32(0) {
		v776 = v629
		goto L168
	} else {
		goto L173
	}
L173:
	;
	goto L169
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+232)) = int32(72)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+224)) = v95
	v646 = F_pg_sprintf(m, v637, int32(494762), v88+int32(224))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L28
	} else {
		goto L177
	}
L175:
	;
	v706 = v637
	goto L176
L176:
	;
	if v94 == int32(0) {
		v776 = v706
		goto L168
	} else {
		goto L195
	}
L177:
	;
	if v637&int32(3) == int32(0) {
		v671 = v637
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v706 = v704 + v637
	goto L176
L179:
	;
	v704 = v696 - v637
	goto L178
L180:
	;
	v675 = v671
	goto L189
L181:
	;
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637))))
	if v655 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v704 = int32(0)
	goto L178
L183:
	;
	goto L184
L184:
	;
	v660 = v637
	goto L185
L185:
	;
	v664 = v660 + int32(1)
	if v664&int32(3) == int32(0) {
		v671 = v664
		goto L180
	} else {
		goto L187
	}
L186:
	;
	v696 = v664
	goto L179
L187:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	if v669 != 0 {
		v660 = v664
		goto L185
	} else {
		goto L188
	}
L188:
	;
	goto L186
L189:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v675)))
	v684 = int32(-2139062144)
	if (int32(16843008)-v681|v681)&v684 == v684 {
		v675 = v675 + int32(4)
		goto L189
	} else {
		goto L191
	}
L190:
	;
	v690 = v675
	goto L192
L191:
	;
	goto L190
L192:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690))))
	if v694 != 0 {
		v690 = v690 + int32(1)
		goto L192
	} else {
		goto L194
	}
L193:
	;
	v696 = v690
	goto L179
L194:
	;
	goto L193
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+216)) = int32(77)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+208)) = base.I64_extend_i32_s(v94)
	v716 = F_pg_sprintf(m, v706, int32(494762), v88+int32(208))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L28
	} else {
		goto L196
	}
L196:
	;
	if v706&int32(3) == int32(0) {
		v741 = v706
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v776 = v774 + v706
	goto L168
L198:
	;
	v774 = v766 - v706
	goto L197
L199:
	;
	v745 = v741
	goto L208
L200:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706))))
	if v725 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v774 = int32(0)
	goto L197
L202:
	;
	goto L203
L203:
	;
	v730 = v706
	goto L204
L204:
	;
	v734 = v730 + int32(1)
	if v734&int32(3) == int32(0) {
		v741 = v734
		goto L199
	} else {
		goto L206
	}
L205:
	;
	v766 = v734
	goto L198
L206:
	;
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734))))
	if v739 != 0 {
		v730 = v734
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v745)))
	v754 = int32(-2139062144)
	if (int32(16843008)-v751|v751)&v754 == v754 {
		v745 = v745 + int32(4)
		goto L208
	} else {
		goto L210
	}
L209:
	;
	v760 = v745
	goto L211
L210:
	;
	goto L209
L211:
	;
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760))))
	if v764 != 0 {
		v760 = v760 + int32(1)
		goto L211
	} else {
		goto L213
	}
L212:
	;
	v766 = v760
	goto L198
L213:
	;
	goto L212
L214:
	;
	if v93|v92 < int32(0) {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v783 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v776))) = uint8(v783)
	v787 = v776 + int32(1)
	goto L217
L216:
	;
	v787 = v776
	goto L217
L217:
	;
	v793 = v93 >> (uint(int32(31)) % 32)
	goto L221
L218:
	;
	v906 = int32(83)
	*(*uint16)(unsafe.Add(mBase, uint32(v905))) = uint16(v906)
	goto L14
L219:
	;
	if v92 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	goto L222
L222:
	;
	v798 = F_pg_ultostr(m, v787, v93^v793-v793)
	mBase = m.M
	goto L219
L223:
	;
	v905 = v798
	goto L218
L224:
	;
	goto L225
L225:
	;
	v802 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v798))) = uint8(v802)
	v805 = v92 >> (uint(int32(31)) % 32)
	v807 = v92 ^ v805 - v805
	v809 = base.I32_div_s(v807, int32(10))
	v812 = v809*int32(-10) + v807
	if v812 != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v822 = base.I32_div_s(v807, int32(100))
	v825 = v822*int32(-10) + v809
	v826 = v812 | v825
	if v826 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L227:
	;
	v814 = v812 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v798)+6)) = uint8(v814)
	v820 = v798 + int32(7)
	goto L226
L228:
	;
	goto L229
L229:
	;
	v820 = v798 + int32(6)
	goto L226
L230:
	;
	v836 = base.I32_div_s(v807, int32(1000))
	v839 = v836*int32(-10) + v822
	v840 = v826 | v839
	if v840 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L231:
	;
	v834 = v798 + int32(5)
	goto L230
L232:
	;
	goto L233
L233:
	;
	v832 = v825 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v798)+5)) = uint8(v832)
	v834 = v820
	goto L230
L234:
	;
	v850 = base.I32_div_s(v807, int32(10000))
	v853 = v850*int32(-10) + v836
	v854 = v840 | v853
	if v854 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L235:
	;
	v848 = v798 + int32(4)
	goto L234
L236:
	;
	goto L237
L237:
	;
	v846 = v839 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v798)+4)) = uint8(v846)
	v848 = v834
	goto L234
L238:
	;
	v864 = base.I32_div_s(v807, int32(100000))
	v867 = v864*int32(-10) + v850
	v868 = v854 | v867
	if v868 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L239:
	;
	v862 = v798 + int32(3)
	goto L238
L240:
	;
	goto L241
L241:
	;
	v860 = v853 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v798)+3)) = uint8(v860)
	v862 = v848
	goto L238
L242:
	;
	v878 = base.I32_div_s(v807, int32(1000000))
	v881 = v878*int32(-10) + v864
	if v868|v881 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L243:
	;
	v876 = v798 + int32(2)
	goto L242
L244:
	;
	goto L245
L245:
	;
	v874 = v867 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v798)+2)) = uint8(v874)
	v876 = v862
	goto L242
L246:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v864+int32(9)) {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	v890 = v798 + int32(1)
	goto L246
L248:
	;
	goto L249
L249:
	;
	v888 = v881 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v798)+1)) = uint8(v888)
	v890 = v876
	goto L246
L250:
	;
	v897 = F_pg_ultostr(m, v798+int32(1), v807)
	mBase = m.M
	v898 = v897
	goto L252
L251:
	;
	v898 = v890
	goto L252
L252:
	;
	v905 = v898
	goto L218
L253:
	;
	if v96 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L254:
	;
	v990 = v20
	v991 = int32(0)
	goto L253
L255:
	;
	goto L256
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+400)) = int32(228166)
	if v97 == int32(1) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v917 = int32(735586)
	goto L259
L258:
	;
	v917 = int32(204454)
	goto L259
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+404)) = v917
	v919 = int32(735586)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+388)) = v919
	*(*int32)(unsafe.Add(mBase, uint32(v88)+384)) = v919
	*(*int64)(unsafe.Add(mBase, uint32(v88)+392)) = base.I64_extend_i32_s(v97)
	v928 = F_pg_sprintf(m, v20, int32(174040), v88+int32(384))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L28
	} else {
		goto L260
	}
L260:
	;
	if v20&int32(3) == int32(0) {
		v955 = v20
		goto L263
	} else {
		goto L264
	}
L261:
	;
	v990 = v988 + v20
	v991 = int32(base.Ui32(v97) >> (uint(int32(31)) % 32))
	goto L253
L262:
	;
	v988 = v980 - v20
	goto L261
L263:
	;
	v959 = v955
	goto L272
L264:
	;
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v939 == int32(0) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v988 = int32(0)
	goto L261
L266:
	;
	goto L267
L267:
	;
	v944 = v20
	goto L268
L268:
	;
	v948 = v944 + int32(1)
	if v948&int32(3) == int32(0) {
		v955 = v948
		goto L263
	} else {
		goto L270
	}
L269:
	;
	v980 = v948
	goto L262
L270:
	;
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948))))
	if v953 != 0 {
		v944 = v948
		goto L268
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v959)))
	v968 = int32(-2139062144)
	if (int32(16843008)-v965|v965)&v968 == v968 {
		v959 = v959 + int32(4)
		goto L272
	} else {
		goto L274
	}
L273:
	;
	v974 = v959
	goto L275
L274:
	;
	goto L273
L275:
	;
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v974))))
	if v978 != 0 {
		v974 = v974 + int32(1)
		goto L275
	} else {
		goto L277
	}
L276:
	;
	v980 = v974
	goto L262
L277:
	;
	goto L276
L278:
	;
	if v98 != 0 {
		goto L312
	} else {
		goto L313
	}
L279:
	;
	v1084 = v990
	v1085 = base.B2i32(v97 == int32(0))
	v1086 = v991
	goto L278
L280:
	;
	goto L281
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+368)) = int32(243522)
	if v96 == int32(1) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v1002 = int32(735586)
	goto L284
L283:
	;
	v1002 = int32(204454)
	goto L284
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+372)) = v1002
	*(*int64)(unsafe.Add(mBase, uint32(v88)+360)) = base.I64_extend_i32_s(v96)
	if v97 != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1008 = int32(724525)
	goto L287
L286:
	;
	v1008 = int32(735586)
	goto L287
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+352)) = v1008
	v1011 = int32(735586)
	if v991 != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1014 = int32(648405)
	goto L290
L289:
	;
	v1014 = v1011
	goto L290
L290:
	;
	if v96 <= int32(0) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1017 = v1011
	goto L293
L292:
	;
	v1017 = v1014
	goto L293
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+356)) = v1017
	v1022 = F_pg_sprintf(m, v990, int32(174040), v88+int32(352))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L28
	} else {
		goto L294
	}
L294:
	;
	if v990&int32(3) == int32(0) {
		v1049 = v990
		goto L297
	} else {
		goto L298
	}
L295:
	;
	v1084 = v1082 + v990
	v1085 = int32(0)
	v1086 = int32(base.Ui32(v96) >> (uint(int32(31)) % 32))
	goto L278
L296:
	;
	v1082 = v1074 - v990
	goto L295
L297:
	;
	v1053 = v1049
	goto L306
L298:
	;
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
	if v1033 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1082 = int32(0)
	goto L295
L300:
	;
	goto L301
L301:
	;
	v1038 = v990
	goto L302
L302:
	;
	v1042 = v1038 + int32(1)
	if v1042&int32(3) == int32(0) {
		v1049 = v1042
		goto L297
	} else {
		goto L304
	}
L303:
	;
	v1074 = v1042
	goto L296
L304:
	;
	v1047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042))))
	if v1047 != 0 {
		v1038 = v1042
		goto L302
	} else {
		goto L305
	}
L305:
	;
	goto L303
L306:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1053)))
	v1062 = int32(-2139062144)
	if (int32(16843008)-v1059|v1059)&v1062 == v1062 {
		v1053 = v1053 + int32(4)
		goto L306
	} else {
		goto L308
	}
L307:
	;
	v1068 = v1053
	goto L309
L308:
	;
	goto L307
L309:
	;
	v1072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1068))))
	if v1072 != 0 {
		v1068 = v1068 + int32(1)
		goto L309
	} else {
		goto L311
	}
L310:
	;
	v1074 = v1068
	goto L296
L311:
	;
	goto L310
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+336)) = int32(26818)
	if v98 == int32(1) {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	v1174 = v1084
	v1175 = v1085
	v1176 = v1086
	goto L314
L314:
	;
	if v1175 != 0 {
		goto L345
	} else {
		goto L346
	}
L315:
	;
	v1093 = int32(735586)
	goto L317
L316:
	;
	v1093 = int32(204454)
	goto L317
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+340)) = v1093
	*(*int64)(unsafe.Add(mBase, uint32(v88)+328)) = v99
	if v1085 != 0 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1098 = int32(735586)
	goto L320
L319:
	;
	v1098 = int32(724525)
	goto L320
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+320)) = v1098
	v1101 = int32(735586)
	if v1086 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1104 = int32(648405)
	goto L323
L322:
	;
	v1104 = v1101
	goto L323
L323:
	;
	if v98 <= int32(0) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1107 = v1101
	goto L326
L325:
	;
	v1107 = v1104
	goto L326
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+324)) = v1107
	v1112 = F_pg_sprintf(m, v1084, int32(174040), v88+int32(320))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L28
	} else {
		goto L327
	}
L327:
	;
	if v1084&int32(3) == int32(0) {
		v1139 = v1084
		goto L330
	} else {
		goto L331
	}
L328:
	;
	v1174 = v1172 + v1084
	v1175 = int32(0)
	v1176 = int32(base.Ui32(v98) >> (uint(int32(31)) % 32))
	goto L314
L329:
	;
	v1172 = v1164 - v1084
	goto L328
L330:
	;
	v1143 = v1139
	goto L339
L331:
	;
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1084))))
	if v1123 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1172 = int32(0)
	goto L328
L333:
	;
	goto L334
L334:
	;
	v1128 = v1084
	goto L335
L335:
	;
	v1132 = v1128 + int32(1)
	if v1132&int32(3) == int32(0) {
		v1139 = v1132
		goto L330
	} else {
		goto L337
	}
L336:
	;
	v1164 = v1132
	goto L329
L337:
	;
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1132))))
	if v1137 != 0 {
		v1128 = v1132
		goto L335
	} else {
		goto L338
	}
L338:
	;
	goto L336
L339:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1143)))
	v1152 = int32(-2139062144)
	if (int32(16843008)-v1149|v1149)&v1152 == v1152 {
		v1143 = v1143 + int32(4)
		goto L339
	} else {
		goto L341
	}
L340:
	;
	v1158 = v1143
	goto L342
L341:
	;
	goto L340
L342:
	;
	v1162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1158))))
	if v1162 != 0 {
		v1158 = v1158 + int32(1)
		goto L342
	} else {
		goto L344
	}
L343:
	;
	v1164 = v1158
	goto L329
L344:
	;
	goto L343
L345:
	;
	v1182 = v94 >> (uint(int32(31)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+304)) = v94 ^ v1182 - v1182
	if v1175 != 0 {
		goto L351
	} else {
		goto L352
	}
L346:
	;
	if v95 != int64(0) {
		goto L345
	} else {
		goto L347
	}
L347:
	;
	if v94 != 0 {
		goto L345
	} else {
		goto L348
	}
L348:
	;
	if v93 != 0 {
		goto L345
	} else {
		goto L349
	}
L349:
	;
	if v92 == int32(0) {
		goto L14
	} else {
		goto L350
	}
L350:
	;
	goto L345
L351:
	;
	v1188 = int32(735586)
	goto L353
L352:
	;
	v1188 = int32(724525)
	goto L353
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+288)) = v1188
	v1191 = v95 >> (uint(int64(63)) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+296)) = v95 ^ v1191 - v1191
	v1195 = int32(648393)
	if v1176 != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1199 = int32(648405)
	goto L356
L355:
	;
	v1199 = int32(735586)
	goto L356
L356:
	;
	if v94|(v93|v92) < int32(0) {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1204 = v1195
	goto L359
L358:
	;
	v1204 = v1199
	goto L359
L359:
	;
	if v95 < int64(0) {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1207 = v1195
	goto L362
L361:
	;
	v1207 = v1204
	goto L362
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+292)) = v1207
	v1212 = F_pg_sprintf(m, v1174, int32(541273), v88+int32(288))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L28
	} else {
		goto L363
	}
L363:
	;
	if v1174&int32(3) == int32(0) {
		v1237 = v1174
		goto L366
	} else {
		goto L367
	}
L364:
	;
	v1277 = v93 >> (uint(int32(31)) % 32)
	goto L383
L365:
	;
	v1270 = v1262 - v1174
	goto L364
L366:
	;
	v1241 = v1237
	goto L375
L367:
	;
	v1221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1174))))
	if v1221 == int32(0) {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1270 = int32(0)
	goto L364
L369:
	;
	goto L370
L370:
	;
	v1226 = v1174
	goto L371
L371:
	;
	v1230 = v1226 + int32(1)
	if v1230&int32(3) == int32(0) {
		v1237 = v1230
		goto L366
	} else {
		goto L373
	}
L372:
	;
	v1262 = v1230
	goto L365
L373:
	;
	v1235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1230))))
	if v1235 != 0 {
		v1226 = v1230
		goto L371
	} else {
		goto L374
	}
L374:
	;
	goto L372
L375:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1241)))
	v1250 = int32(-2139062144)
	if (int32(16843008)-v1247|v1247)&v1250 == v1250 {
		v1241 = v1241 + int32(4)
		goto L375
	} else {
		goto L377
	}
L376:
	;
	v1256 = v1241
	goto L378
L377:
	;
	goto L376
L378:
	;
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1256))))
	if v1260 != 0 {
		v1256 = v1256 + int32(1)
		goto L378
	} else {
		goto L380
	}
L379:
	;
	v1262 = v1256
	goto L365
L380:
	;
	goto L379
L381:
	;
	v1390 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1389))) = uint8(v1390)
	goto L14
L382:
	;
	if v92 == int32(0) {
		goto L386
	} else {
		goto L387
	}
L383:
	;
	v1281 = F_pg_ultostr_zeropad(m, v1270+v1174, v93^v1277-v1277, int32(2))
	mBase = m.M
	goto L382
L386:
	;
	v1389 = v1281
	goto L381
L387:
	;
	goto L388
L388:
	;
	v1286 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1281))) = uint8(v1286)
	v1289 = v92 >> (uint(int32(31)) % 32)
	v1291 = v92 ^ v1289 - v1289
	v1293 = base.I32_div_s(v1291, int32(10))
	v1296 = v1293*int32(-10) + v1291
	if v1296 != 0 {
		goto L390
	} else {
		goto L391
	}
L389:
	;
	v1306 = base.I32_div_s(v1291, int32(100))
	v1309 = v1306*int32(-10) + v1293
	v1310 = v1296 | v1309
	if v1310 == int32(0) {
		goto L394
	} else {
		goto L395
	}
L390:
	;
	v1298 = v1296 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1281)+6)) = uint8(v1298)
	v1304 = v1281 + int32(7)
	goto L389
L391:
	;
	goto L392
L392:
	;
	v1304 = v1281 + int32(6)
	goto L389
L393:
	;
	v1320 = base.I32_div_s(v1291, int32(1000))
	v1323 = v1320*int32(-10) + v1306
	v1324 = v1310 | v1323
	if v1324 == int32(0) {
		goto L398
	} else {
		goto L399
	}
L394:
	;
	v1318 = v1281 + int32(5)
	goto L393
L395:
	;
	goto L396
L396:
	;
	v1316 = v1309 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1281)+5)) = uint8(v1316)
	v1318 = v1304
	goto L393
L397:
	;
	v1334 = base.I32_div_s(v1291, int32(10000))
	v1337 = v1334*int32(-10) + v1320
	v1338 = v1324 | v1337
	if v1338 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L398:
	;
	v1332 = v1281 + int32(4)
	goto L397
L399:
	;
	goto L400
L400:
	;
	v1330 = v1323 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1281)+4)) = uint8(v1330)
	v1332 = v1318
	goto L397
L401:
	;
	v1348 = base.I32_div_s(v1291, int32(100000))
	v1351 = v1348*int32(-10) + v1334
	v1352 = v1338 | v1351
	if v1352 == int32(0) {
		goto L406
	} else {
		goto L407
	}
L402:
	;
	v1346 = v1281 + int32(3)
	goto L401
L403:
	;
	goto L404
L404:
	;
	v1344 = v1337 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1281)+3)) = uint8(v1344)
	v1346 = v1332
	goto L401
L405:
	;
	v1362 = base.I32_div_s(v1291, int32(1000000))
	v1365 = v1362*int32(-10) + v1348
	if v1352|v1365 == int32(0) {
		goto L410
	} else {
		goto L411
	}
L406:
	;
	v1360 = v1281 + int32(2)
	goto L405
L407:
	;
	goto L408
L408:
	;
	v1358 = v1351 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1281)+2)) = uint8(v1358)
	v1360 = v1346
	goto L405
L409:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v1348+int32(9)) {
		goto L413
	} else {
		goto L414
	}
L410:
	;
	v1374 = v1281 + int32(1)
	goto L409
L411:
	;
	goto L412
L412:
	;
	v1372 = v1365 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1281)+1)) = uint8(v1372)
	v1374 = v1360
	goto L409
L413:
	;
	v1381 = F_pg_ultostr(m, v1281+int32(1), v1291)
	mBase = m.M
	v1382 = v1381
	goto L415
L414:
	;
	v1382 = v1374
	goto L415
L415:
	;
	v1389 = v1382
	goto L381
L416:
	;
	v1681 = int64(0)
	if v95 != v1681 {
		goto L506
	} else {
		goto L507
	}
L417:
	;
	v1675 = v1673
	v1677 = v1670
	v1680 = int32(0)
	goto L416
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+56)) = int32(26818)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+48)) = v1595
	if v1595 == int64(1) {
		goto L479
	} else {
		goto L480
	}
L419:
	;
	if v1586 != 0 {
		goto L476
	} else {
		goto L477
	}
L420:
	;
	v1574 = int32(0)
	if v98 == v1574 {
		v1675 = v1572
		v1677 = v1573
		v1680 = base.B2i32(v97 == v1574)
		goto L416
	} else {
		goto L474
	}
L421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+72)) = int32(243522)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+64)) = v1497
	if v1497 == int64(1) {
		goto L452
	} else {
		goto L453
	}
L422:
	;
	v1486 = int32(31)
	v1489 = v96 >> (uint(v1486) % 32)
	v1493 = v1395
	v1495 = int32(base.Ui32(v96) >> (uint(v1486) % 32))
	v1497 = base.I64_extend_i32_u(v96 ^ v1489 - v1489)
	goto L421
L423:
	;
	if v96 != 0 {
		goto L422
	} else {
		goto L426
	}
L424:
	;
	goto L425
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+88)) = int32(228166)
	v1402 = v97 >> (uint(int32(31)) % 32)
	v1404 = v97 ^ v1402 - v1402
	*(*int64)(unsafe.Add(mBase, uint32(v88)+80)) = base.I64_extend_i32_u(v1404)
	if v1404 == int32(1) {
		goto L427
	} else {
		goto L428
	}
L426:
	;
	v1572 = v1395
	v1573 = int32(0)
	goto L420
L427:
	;
	v1411 = int32(735586)
	goto L429
L428:
	;
	v1411 = int32(204454)
	goto L429
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+92)) = v1411
	v1416 = F_pg_sprintf(m, v1395, int32(174054), v88+int32(80))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L28
	} else {
		goto L430
	}
L430:
	;
	v1419 = int32(base.Ui32(v97) >> (uint(int32(31)) % 32))
	if v1395&int32(3) == int32(0) {
		v1443 = v1395
		goto L433
	} else {
		goto L434
	}
L431:
	;
	v1477 = v1476 + v1395
	if v96 == int32(0) {
		v1572 = v1477
		v1573 = v1419
		goto L420
	} else {
		goto L448
	}
L432:
	;
	v1476 = v1468 - v1395
	goto L431
L433:
	;
	v1447 = v1443
	goto L442
L434:
	;
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1395))))
	if v1427 == int32(0) {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v1476 = int32(0)
	goto L431
L436:
	;
	goto L437
L437:
	;
	v1432 = v1395
	goto L438
L438:
	;
	v1436 = v1432 + int32(1)
	if v1436&int32(3) == int32(0) {
		v1443 = v1436
		goto L433
	} else {
		goto L440
	}
L439:
	;
	v1468 = v1436
	goto L432
L440:
	;
	v1441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1436))))
	if v1441 != 0 {
		v1432 = v1436
		goto L438
	} else {
		goto L441
	}
L441:
	;
	goto L439
L442:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v1447)))
	v1456 = int32(-2139062144)
	if (int32(16843008)-v1453|v1453)&v1456 == v1456 {
		v1447 = v1447 + int32(4)
		goto L442
	} else {
		goto L444
	}
L443:
	;
	v1462 = v1447
	goto L445
L444:
	;
	goto L443
L445:
	;
	v1466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1462))))
	if v1466 != 0 {
		v1462 = v1462 + int32(1)
		goto L445
	} else {
		goto L447
	}
L446:
	;
	v1468 = v1462
	goto L432
L447:
	;
	goto L446
L448:
	;
	v1481 = base.I64_extend_i32_s(v96)
	if v97 < int32(0) {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	v1485 = int64(0) - v1481
	goto L451
L450:
	;
	v1485 = v1481
	goto L451
L451:
	;
	v1493 = v1477
	v1495 = v1419
	v1497 = v1485
	goto L421
L452:
	;
	v1505 = int32(735586)
	goto L454
L453:
	;
	v1505 = int32(204454)
	goto L454
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+76)) = v1505
	v1510 = F_pg_sprintf(m, v1493, int32(174054), v88-int32(-64))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L28
	} else {
		goto L455
	}
L455:
	;
	if v1493&int32(3) == int32(0) {
		v1535 = v1493
		goto L458
	} else {
		goto L459
	}
L456:
	;
	v1569 = v1568 + v1493
	if v98 == int32(0) {
		v1670 = v1495
		v1673 = v1569
		goto L417
	} else {
		goto L473
	}
L457:
	;
	v1568 = v1560 - v1493
	goto L456
L458:
	;
	v1539 = v1535
	goto L467
L459:
	;
	v1519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1493))))
	if v1519 == int32(0) {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v1568 = int32(0)
	goto L456
L461:
	;
	goto L462
L462:
	;
	v1524 = v1493
	goto L463
L463:
	;
	v1528 = v1524 + int32(1)
	if v1528&int32(3) == int32(0) {
		v1535 = v1528
		goto L458
	} else {
		goto L465
	}
L464:
	;
	v1560 = v1528
	goto L457
L465:
	;
	v1533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1528))))
	if v1533 != 0 {
		v1524 = v1528
		goto L463
	} else {
		goto L466
	}
L466:
	;
	goto L464
L467:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1539)))
	v1548 = int32(-2139062144)
	if (int32(16843008)-v1545|v1545)&v1548 == v1548 {
		v1539 = v1539 + int32(4)
		goto L467
	} else {
		goto L469
	}
L468:
	;
	v1554 = v1539
	goto L470
L469:
	;
	goto L468
L470:
	;
	v1558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1554))))
	if v1558 != 0 {
		v1554 = v1554 + int32(1)
		goto L470
	} else {
		goto L472
	}
L471:
	;
	v1560 = v1554
	goto L457
L472:
	;
	goto L471
L473:
	;
	v1584 = v1569
	v1586 = v1495
	goto L419
L474:
	;
	if v97 != 0 {
		v1584 = v1572
		v1586 = v1573
		goto L419
	} else {
		goto L475
	}
L475:
	;
	v1581 = v99 >> (uint(int64(63)) % 64)
	v1591 = v1572
	v1593 = int32(base.Ui32(v98) >> (uint(int32(31)) % 32))
	v1595 = v99 ^ v1581 - v1581
	goto L418
L476:
	;
	v1590 = int64(0) - v99
	goto L478
L477:
	;
	v1590 = v99
	goto L478
L478:
	;
	v1591 = v1584
	v1593 = v1586
	v1595 = v1590
	goto L418
L479:
	;
	v1603 = int32(735586)
	goto L481
L480:
	;
	v1603 = int32(204454)
	goto L481
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+60)) = v1603
	v1608 = F_pg_sprintf(m, v1591, int32(174054), v88+int32(48))
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L28
	} else {
		goto L482
	}
L482:
	;
	if v1591&int32(3) == int32(0) {
		v1633 = v1591
		goto L485
	} else {
		goto L486
	}
L483:
	;
	v1670 = v1593
	v1673 = v1666 + v1591
	goto L417
L484:
	;
	v1666 = v1658 - v1591
	goto L483
L485:
	;
	v1637 = v1633
	goto L494
L486:
	;
	v1617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1591))))
	if v1617 == int32(0) {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v1666 = int32(0)
	goto L483
L488:
	;
	goto L489
L489:
	;
	v1622 = v1591
	goto L490
L490:
	;
	v1626 = v1622 + int32(1)
	if v1626&int32(3) == int32(0) {
		v1633 = v1626
		goto L485
	} else {
		goto L492
	}
L491:
	;
	v1658 = v1626
	goto L484
L492:
	;
	v1631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1626))))
	if v1631 != 0 {
		v1622 = v1626
		goto L490
	} else {
		goto L493
	}
L493:
	;
	goto L491
L494:
	;
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1637)))
	v1646 = int32(-2139062144)
	if (int32(16843008)-v1643|v1643)&v1646 == v1646 {
		v1637 = v1637 + int32(4)
		goto L494
	} else {
		goto L496
	}
L495:
	;
	v1652 = v1637
	goto L497
L496:
	;
	goto L495
L497:
	;
	v1656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1652))))
	if v1656 != 0 {
		v1652 = v1652 + int32(1)
		goto L497
	} else {
		goto L499
	}
L498:
	;
	v1658 = v1652
	goto L484
L499:
	;
	goto L498
L500:
	;
	if v93|v92 != 0 {
		goto L562
	} else {
		goto L563
	}
L501:
	;
	v1875 = v1873
	v1876 = int32(0)
	v1877 = v1870
	goto L500
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = int32(273522)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v1796
	if v1796 == int64(1) {
		goto L540
	} else {
		goto L541
	}
L503:
	;
	if v1786 != 0 {
		goto L537
	} else {
		goto L538
	}
L504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+40)) = int32(204472)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+32)) = v1709
	if v1709 == int64(1) {
		goto L515
	} else {
		goto L516
	}
L505:
	;
	if v1677 != 0 {
		goto L512
	} else {
		goto L513
	}
L506:
	;
	if v1680 == int32(0) {
		goto L505
	} else {
		goto L509
	}
L507:
	;
	goto L508
L508:
	;
	if v94 == int32(0) {
		v1875 = v1675
		v1876 = v1680
		v1877 = v1677
		goto L500
	} else {
		goto L510
	}
L509:
	;
	v1686 = int64(63)
	v1690 = v95 >> (uint(v1686) % 64)
	v1707 = base.I32_wrap_i64(int64(base.Ui64(v95) >> (uint(v1686) % 64)))
	v1709 = v95 ^ v1690 - v1690
	goto L504
L510:
	;
	v1695 = base.I64_extend_i32_s(v94)
	if v1680 == int32(0) {
		v1785 = v1675
		v1786 = v1677
		v1789 = v1695
		goto L503
	} else {
		goto L511
	}
L511:
	;
	v1701 = v1695 >> (uint(int64(63)) % 64)
	v1792 = v1675
	v1793 = int32(base.Ui32(v94) >> (uint(int32(31)) % 32))
	v1796 = v1695 ^ v1701 - v1701
	goto L502
L512:
	;
	v1706 = int64(0) - v95
	goto L514
L513:
	;
	v1706 = v95
	goto L514
L514:
	;
	v1707 = v1677
	v1709 = v1706
	goto L504
L515:
	;
	v1717 = int32(735586)
	goto L517
L516:
	;
	v1717 = int32(204454)
	goto L517
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+44)) = v1717
	v1722 = F_pg_sprintf(m, v1675, int32(174054), v88+int32(32))
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		goto L28
	} else {
		goto L518
	}
L518:
	;
	if v1675&int32(3) == int32(0) {
		v1747 = v1675
		goto L521
	} else {
		goto L522
	}
L519:
	;
	v1781 = v1780 + v1675
	if v94 == int32(0) {
		v1870 = v1707
		v1873 = v1781
		goto L501
	} else {
		goto L536
	}
L520:
	;
	v1780 = v1772 - v1675
	goto L519
L521:
	;
	v1751 = v1747
	goto L530
L522:
	;
	v1731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1675))))
	if v1731 == int32(0) {
		goto L523
	} else {
		goto L524
	}
L523:
	;
	v1780 = int32(0)
	goto L519
L524:
	;
	goto L525
L525:
	;
	v1736 = v1675
	goto L526
L526:
	;
	v1740 = v1736 + int32(1)
	if v1740&int32(3) == int32(0) {
		v1747 = v1740
		goto L521
	} else {
		goto L528
	}
L527:
	;
	v1772 = v1740
	goto L520
L528:
	;
	v1745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1740))))
	if v1745 != 0 {
		v1736 = v1740
		goto L526
	} else {
		goto L529
	}
L529:
	;
	goto L527
L530:
	;
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1751)))
	v1760 = int32(-2139062144)
	if (int32(16843008)-v1757|v1757)&v1760 == v1760 {
		v1751 = v1751 + int32(4)
		goto L530
	} else {
		goto L532
	}
L531:
	;
	v1766 = v1751
	goto L533
L532:
	;
	goto L531
L533:
	;
	v1770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1766))))
	if v1770 != 0 {
		v1766 = v1766 + int32(1)
		goto L533
	} else {
		goto L535
	}
L534:
	;
	v1772 = v1766
	goto L520
L535:
	;
	goto L534
L536:
	;
	v1785 = v1781
	v1786 = v1707
	v1789 = base.I64_extend_i32_s(v94)
	goto L503
L537:
	;
	v1791 = v1681 - v1789
	goto L539
L538:
	;
	v1791 = v1789
	goto L539
L539:
	;
	v1792 = v1785
	v1793 = v1786
	v1796 = v1791
	goto L502
L540:
	;
	v1804 = int32(735586)
	goto L542
L541:
	;
	v1804 = int32(204454)
	goto L542
L542:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = v1804
	v1809 = F_pg_sprintf(m, v1792, int32(174054), v88+int32(16))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L28
	} else {
		goto L543
	}
L543:
	;
	if v1792&int32(3) == int32(0) {
		v1834 = v1792
		goto L546
	} else {
		goto L547
	}
L544:
	;
	v1870 = v1793
	v1873 = v1867 + v1792
	goto L501
L545:
	;
	v1867 = v1859 - v1792
	goto L544
L546:
	;
	v1838 = v1834
	goto L555
L547:
	;
	v1818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1792))))
	if v1818 == int32(0) {
		goto L548
	} else {
		goto L549
	}
L548:
	;
	v1867 = int32(0)
	goto L544
L549:
	;
	goto L550
L550:
	;
	v1823 = v1792
	goto L551
L551:
	;
	v1827 = v1823 + int32(1)
	if v1827&int32(3) == int32(0) {
		v1834 = v1827
		goto L546
	} else {
		goto L553
	}
L552:
	;
	v1859 = v1827
	goto L545
L553:
	;
	v1832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1827))))
	if v1832 != 0 {
		v1823 = v1827
		goto L551
	} else {
		goto L554
	}
L554:
	;
	goto L552
L555:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1838)))
	v1847 = int32(-2139062144)
	if (int32(16843008)-v1844|v1844)&v1847 == v1847 {
		v1838 = v1838 + int32(4)
		goto L555
	} else {
		goto L557
	}
L556:
	;
	v1853 = v1838
	goto L558
L557:
	;
	goto L556
L558:
	;
	v1857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1853))))
	if v1857 != 0 {
		v1853 = v1853 + int32(1)
		goto L558
	} else {
		goto L560
	}
L559:
	;
	v1859 = v1853
	goto L545
L560:
	;
	goto L559
L561:
	;
	if v2110&int32(3) == int32(0) {
		v2137 = v2110
		goto L641
	} else {
		goto L642
	}
L562:
	;
	v1881 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1875))) = uint8(v1881)
	v1883 = int32(1)
	v1885 = v1875 + v1883
	v1886 = int32(0)
	if v1886 <= v93 {
		goto L568
	} else {
		goto L569
	}
L563:
	;
	goto L564
L564:
	;
	if v1876 != 0 {
		goto L618
	} else {
		goto L619
	}
L565:
	;
	v1914 = v93 >> (uint(int32(31)) % 32)
	goto L578
L566:
	;
	v1902 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1875)+1)) = uint8(v1902)
	v1907 = v1900
	v1908 = v1875 + int32(2)
	goto L565
L567:
	;
	v1896 = int32(0)
	if v1877 == v1896 {
		v1907 = v1896
		v1908 = v1885
		goto L565
	} else {
		goto L574
	}
L568:
	;
	if v93 != 0 {
		goto L567
	} else {
		goto L571
	}
L569:
	;
	goto L570
L570:
	;
	if (v1876|v1877)&int32(1) == int32(0) {
		v1900 = v1886
		goto L566
	} else {
		goto L573
	}
L571:
	;
	if int32(0) <= v92 {
		goto L567
	} else {
		goto L572
	}
L572:
	;
	goto L570
L573:
	;
	v1907 = v1883
	v1908 = v1885
	goto L565
L574:
	;
	v1900 = int32(1)
	goto L566
L575:
	;
	v2027 = int32(204454)
	if v92 != 0 {
		goto L610
	} else {
		goto L611
	}
L576:
	;
	if v92 == int32(0) {
		goto L580
	} else {
		goto L581
	}
L578:
	;
	goto L579
L579:
	;
	v1919 = F_pg_ultostr(m, v1908, v93^v1914-v1914)
	mBase = m.M
	goto L576
L580:
	;
	v2026 = v1919
	goto L575
L581:
	;
	goto L582
L582:
	;
	v1923 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1919))) = uint8(v1923)
	v1926 = v92 >> (uint(int32(31)) % 32)
	v1928 = v92 ^ v1926 - v1926
	v1930 = base.I32_div_s(v1928, int32(10))
	v1933 = v1930*int32(-10) + v1928
	if v1933 != 0 {
		goto L584
	} else {
		goto L585
	}
L583:
	;
	v1943 = base.I32_div_s(v1928, int32(100))
	v1946 = v1943*int32(-10) + v1930
	v1947 = v1933 | v1946
	if v1947 == int32(0) {
		goto L588
	} else {
		goto L589
	}
L584:
	;
	v1935 = v1933 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1919)+6)) = uint8(v1935)
	v1941 = v1919 + int32(7)
	goto L583
L585:
	;
	goto L586
L586:
	;
	v1941 = v1919 + int32(6)
	goto L583
L587:
	;
	v1957 = base.I32_div_s(v1928, int32(1000))
	v1960 = v1957*int32(-10) + v1943
	v1961 = v1947 | v1960
	if v1961 == int32(0) {
		goto L592
	} else {
		goto L593
	}
L588:
	;
	v1955 = v1919 + int32(5)
	goto L587
L589:
	;
	goto L590
L590:
	;
	v1953 = v1946 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1919)+5)) = uint8(v1953)
	v1955 = v1941
	goto L587
L591:
	;
	v1971 = base.I32_div_s(v1928, int32(10000))
	v1974 = v1971*int32(-10) + v1957
	v1975 = v1961 | v1974
	if v1975 == int32(0) {
		goto L596
	} else {
		goto L597
	}
L592:
	;
	v1969 = v1919 + int32(4)
	goto L591
L593:
	;
	goto L594
L594:
	;
	v1967 = v1960 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1919)+4)) = uint8(v1967)
	v1969 = v1955
	goto L591
L595:
	;
	v1985 = base.I32_div_s(v1928, int32(100000))
	v1988 = v1985*int32(-10) + v1971
	v1989 = v1975 | v1988
	if v1989 == int32(0) {
		goto L600
	} else {
		goto L601
	}
L596:
	;
	v1983 = v1919 + int32(3)
	goto L595
L597:
	;
	goto L598
L598:
	;
	v1981 = v1974 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1919)+3)) = uint8(v1981)
	v1983 = v1969
	goto L595
L599:
	;
	v1999 = base.I32_div_s(v1928, int32(1000000))
	v2002 = v1999*int32(-10) + v1985
	if v1989|v2002 == int32(0) {
		goto L604
	} else {
		goto L605
	}
L600:
	;
	v1997 = v1919 + int32(2)
	goto L599
L601:
	;
	goto L602
L602:
	;
	v1995 = v1988 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1919)+2)) = uint8(v1995)
	v1997 = v1983
	goto L599
L603:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v1985+int32(9)) {
		goto L607
	} else {
		goto L608
	}
L604:
	;
	v2011 = v1919 + int32(1)
	goto L603
L605:
	;
	goto L606
L606:
	;
	v2009 = v2002 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1919)+1)) = uint8(v2009)
	v2011 = v1997
	goto L603
L607:
	;
	v2018 = F_pg_ultostr(m, v1919+int32(1), v1928)
	mBase = m.M
	v2019 = v2018
	goto L609
L608:
	;
	v2019 = v2011
	goto L609
L609:
	;
	v2026 = v2019
	goto L575
L610:
	;
	v2030 = v2027
	goto L612
L611:
	;
	v2030 = int32(735586)
	goto L612
L612:
	;
	v2032 = v93 >> (uint(int32(31)) % 32)
	if v93^v2032-v2032 != int32(1) {
		goto L613
	} else {
		goto L614
	}
L613:
	;
	v2037 = v2027
	goto L615
L614:
	;
	v2037 = v2030
	goto L615
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v2037
	v2040 = F_pg_sprintf(m, v2026, int32(174170), v88)
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L28
	} else {
		goto L616
	}
L616:
	;
	if v1907 != 0 {
		v2110 = v2026
		goto L561
	} else {
		goto L617
	}
L617:
	;
	goto L14
L618:
	;
	if v1875&int32(3) == int32(0) {
		v2065 = v1875
		goto L623
	} else {
		goto L624
	}
L619:
	;
	goto L620
L620:
	;
	if v1877 == int32(0) {
		goto L14
	} else {
		goto L638
	}
L621:
	;
	v2099 = v2098 + v1875
	v2101 = int32(*(*uint16)(unsafe.Add(mBase, _consts[859])))
	*(*uint16)(unsafe.Add(mBase, uint32(v2099))) = uint16(v2101)
	v2104 = int32(*(*uint8)(unsafe.Add(mBase, _consts[860])))
	*(*uint8)(unsafe.Add(mBase, uint32(v2099)+2)) = uint8(v2104)
	goto L620
L622:
	;
	v2098 = v2090 - v1875
	goto L621
L623:
	;
	v2069 = v2065
	goto L632
L624:
	;
	v2049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1875))))
	if v2049 == int32(0) {
		goto L625
	} else {
		goto L626
	}
L625:
	;
	v2098 = int32(0)
	goto L621
L626:
	;
	goto L627
L627:
	;
	v2054 = v1875
	goto L628
L628:
	;
	v2058 = v2054 + int32(1)
	if v2058&int32(3) == int32(0) {
		v2065 = v2058
		goto L623
	} else {
		goto L630
	}
L629:
	;
	v2090 = v2058
	goto L622
L630:
	;
	v2063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2058))))
	if v2063 != 0 {
		v2054 = v2058
		goto L628
	} else {
		goto L631
	}
L631:
	;
	goto L629
L632:
	;
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v2069)))
	v2078 = int32(-2139062144)
	if (int32(16843008)-v2075|v2075)&v2078 == v2078 {
		v2069 = v2069 + int32(4)
		goto L632
	} else {
		goto L634
	}
L633:
	;
	v2084 = v2069
	goto L635
L634:
	;
	goto L633
L635:
	;
	v2088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2084))))
	if v2088 != 0 {
		v2084 = v2084 + int32(1)
		goto L635
	} else {
		goto L637
	}
L636:
	;
	v2090 = v2084
	goto L622
L637:
	;
	goto L636
L638:
	;
	v2110 = v1875
	goto L561
L639:
	;
	v2171 = v2170 + v2110
	v2173 = *(*int32)(unsafe.Add(mBase, _consts[861]))
	*(*int32)(unsafe.Add(mBase, uint32(v2171))) = v2173
	v2176 = int32(*(*uint8)(unsafe.Add(mBase, _consts[862])))
	*(*uint8)(unsafe.Add(mBase, uint32(v2171)+4)) = uint8(v2176)
	goto L14
L640:
	;
	v2170 = v2162 - v2110
	goto L639
L641:
	;
	v2141 = v2137
	goto L650
L642:
	;
	v2121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2110))))
	if v2121 == int32(0) {
		goto L643
	} else {
		goto L644
	}
L643:
	;
	v2170 = int32(0)
	goto L639
L644:
	;
	goto L645
L645:
	;
	v2126 = v2110
	goto L646
L646:
	;
	v2130 = v2126 + int32(1)
	if v2130&int32(3) == int32(0) {
		v2137 = v2130
		goto L641
	} else {
		goto L648
	}
L647:
	;
	v2162 = v2130
	goto L640
L648:
	;
	v2135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2130))))
	if v2135 != 0 {
		v2126 = v2130
		goto L646
	} else {
		goto L649
	}
L649:
	;
	goto L647
L650:
	;
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v2141)))
	v2150 = int32(-2139062144)
	if (int32(16843008)-v2147|v2147)&v2150 == v2150 {
		v2141 = v2141 + int32(4)
		goto L650
	} else {
		goto L652
	}
L651:
	;
	v2156 = v2141
	goto L653
L652:
	;
	goto L651
L653:
	;
	v2160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2156))))
	if v2160 != 0 {
		v2156 = v2156 + int32(1)
		goto L653
	} else {
		goto L655
	}
L654:
	;
	v2162 = v2156
	goto L640
L655:
	;
	goto L654
L656:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+100)) = v2178
	*(*int32)(unsafe.Add(mBase, uint32(v88)+96)) = v2180
	v2191 = F_pg_sprintf(m, v2179, int32(461877), v88+int32(96))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L28
	} else {
		goto L659
	}
L657:
	;
	goto L658
L658:
	;
	if v98 != 0 {
		goto L660
	} else {
		goto L661
	}
L659:
	;
	goto L14
L660:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+128)) = v2181
	*(*int64)(unsafe.Add(mBase, uint32(v88)+120)) = v2184
	*(*int64)(unsafe.Add(mBase, uint32(v88)+112)) = v2185
	v2199 = F_pg_sprintf(m, v2179, int32(541318), v88+int32(112))
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L28
	} else {
		goto L663
	}
L661:
	;
	goto L662
L662:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+152)) = v2181
	*(*int64)(unsafe.Add(mBase, uint32(v88)+144)) = v2184
	v2384 = F_pg_sprintf(m, v2179, int32(541323), v88+int32(144))
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L28
	} else {
		goto L716
	}
L663:
	;
	if v2179&int32(3) == int32(0) {
		v2224 = v2179
		goto L666
	} else {
		goto L667
	}
L664:
	;
	v2264 = v2182 >> (uint(int32(31)) % 32)
	goto L683
L665:
	;
	v2257 = v2249 - v2179
	goto L664
L666:
	;
	v2228 = v2224
	goto L675
L667:
	;
	v2208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2179))))
	if v2208 == int32(0) {
		goto L668
	} else {
		goto L669
	}
L668:
	;
	v2257 = int32(0)
	goto L664
L669:
	;
	goto L670
L670:
	;
	v2213 = v2179
	goto L671
L671:
	;
	v2217 = v2213 + int32(1)
	if v2217&int32(3) == int32(0) {
		v2224 = v2217
		goto L666
	} else {
		goto L673
	}
L672:
	;
	v2249 = v2217
	goto L665
L673:
	;
	v2222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2217))))
	if v2222 != 0 {
		v2213 = v2217
		goto L671
	} else {
		goto L674
	}
L674:
	;
	goto L672
L675:
	;
	v2234 = *(*int32)(unsafe.Add(mBase, uint32(v2228)))
	v2237 = int32(-2139062144)
	if (int32(16843008)-v2234|v2234)&v2237 == v2237 {
		v2228 = v2228 + int32(4)
		goto L675
	} else {
		goto L677
	}
L676:
	;
	v2243 = v2228
	goto L678
L677:
	;
	goto L676
L678:
	;
	v2247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2243))))
	if v2247 != 0 {
		v2243 = v2243 + int32(1)
		goto L678
	} else {
		goto L680
	}
L679:
	;
	v2249 = v2243
	goto L665
L680:
	;
	goto L679
L681:
	;
	v2377 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2376))) = uint8(v2377)
	goto L14
L682:
	;
	if v2183 == int32(0) {
		goto L686
	} else {
		goto L687
	}
L683:
	;
	v2268 = F_pg_ultostr_zeropad(m, v2257+v2179, v2182^v2264-v2264, int32(2))
	mBase = m.M
	goto L682
L686:
	;
	v2376 = v2268
	goto L681
L687:
	;
	goto L688
L688:
	;
	v2273 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2268))) = uint8(v2273)
	v2276 = v2183 >> (uint(int32(31)) % 32)
	v2278 = v2183 ^ v2276 - v2276
	v2280 = base.I32_div_s(v2278, int32(10))
	v2283 = v2280*int32(-10) + v2278
	if v2283 != 0 {
		goto L690
	} else {
		goto L691
	}
L689:
	;
	v2293 = base.I32_div_s(v2278, int32(100))
	v2296 = v2293*int32(-10) + v2280
	v2297 = v2283 | v2296
	if v2297 == int32(0) {
		goto L694
	} else {
		goto L695
	}
L690:
	;
	v2285 = v2283 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2268)+6)) = uint8(v2285)
	v2291 = v2268 + int32(7)
	goto L689
L691:
	;
	goto L692
L692:
	;
	v2291 = v2268 + int32(6)
	goto L689
L693:
	;
	v2307 = base.I32_div_s(v2278, int32(1000))
	v2310 = v2307*int32(-10) + v2293
	v2311 = v2297 | v2310
	if v2311 == int32(0) {
		goto L698
	} else {
		goto L699
	}
L694:
	;
	v2305 = v2268 + int32(5)
	goto L693
L695:
	;
	goto L696
L696:
	;
	v2303 = v2296 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2268)+5)) = uint8(v2303)
	v2305 = v2291
	goto L693
L697:
	;
	v2321 = base.I32_div_s(v2278, int32(10000))
	v2324 = v2321*int32(-10) + v2307
	v2325 = v2311 | v2324
	if v2325 == int32(0) {
		goto L702
	} else {
		goto L703
	}
L698:
	;
	v2319 = v2268 + int32(4)
	goto L697
L699:
	;
	goto L700
L700:
	;
	v2317 = v2310 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2268)+4)) = uint8(v2317)
	v2319 = v2305
	goto L697
L701:
	;
	v2335 = base.I32_div_s(v2278, int32(100000))
	v2338 = v2335*int32(-10) + v2321
	v2339 = v2325 | v2338
	if v2339 == int32(0) {
		goto L706
	} else {
		goto L707
	}
L702:
	;
	v2333 = v2268 + int32(3)
	goto L701
L703:
	;
	goto L704
L704:
	;
	v2331 = v2324 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2268)+3)) = uint8(v2331)
	v2333 = v2319
	goto L701
L705:
	;
	v2349 = base.I32_div_s(v2278, int32(1000000))
	v2352 = v2349*int32(-10) + v2335
	if v2339|v2352 == int32(0) {
		goto L710
	} else {
		goto L711
	}
L706:
	;
	v2347 = v2268 + int32(2)
	goto L705
L707:
	;
	goto L708
L708:
	;
	v2345 = v2338 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2268)+2)) = uint8(v2345)
	v2347 = v2333
	goto L705
L709:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v2335+int32(9)) {
		goto L713
	} else {
		goto L714
	}
L710:
	;
	v2361 = v2268 + int32(1)
	goto L709
L711:
	;
	goto L712
L712:
	;
	v2359 = v2352 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2268)+1)) = uint8(v2359)
	v2361 = v2347
	goto L709
L713:
	;
	v2368 = F_pg_ultostr(m, v2268+int32(1), v2278)
	mBase = m.M
	v2369 = v2368
	goto L715
L714:
	;
	v2369 = v2361
	goto L715
L715:
	;
	v2376 = v2369
	goto L681
L716:
	;
	if v2179&int32(3) == int32(0) {
		v2409 = v2179
		goto L719
	} else {
		goto L720
	}
L717:
	;
	v2449 = v2182 >> (uint(int32(31)) % 32)
	goto L736
L718:
	;
	v2442 = v2434 - v2179
	goto L717
L719:
	;
	v2413 = v2409
	goto L728
L720:
	;
	v2393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2179))))
	if v2393 == int32(0) {
		goto L721
	} else {
		goto L722
	}
L721:
	;
	v2442 = int32(0)
	goto L717
L722:
	;
	goto L723
L723:
	;
	v2398 = v2179
	goto L724
L724:
	;
	v2402 = v2398 + int32(1)
	if v2402&int32(3) == int32(0) {
		v2409 = v2402
		goto L719
	} else {
		goto L726
	}
L725:
	;
	v2434 = v2402
	goto L718
L726:
	;
	v2407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2402))))
	if v2407 != 0 {
		v2398 = v2402
		goto L724
	} else {
		goto L727
	}
L727:
	;
	goto L725
L728:
	;
	v2419 = *(*int32)(unsafe.Add(mBase, uint32(v2413)))
	v2422 = int32(-2139062144)
	if (int32(16843008)-v2419|v2419)&v2422 == v2422 {
		v2413 = v2413 + int32(4)
		goto L728
	} else {
		goto L730
	}
L729:
	;
	v2428 = v2413
	goto L731
L730:
	;
	goto L729
L731:
	;
	v2432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2428))))
	if v2432 != 0 {
		v2428 = v2428 + int32(1)
		goto L731
	} else {
		goto L733
	}
L732:
	;
	v2434 = v2428
	goto L718
L733:
	;
	goto L732
L734:
	;
	v2562 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2561))) = uint8(v2562)
	goto L14
L735:
	;
	if v2183 == int32(0) {
		goto L739
	} else {
		goto L740
	}
L736:
	;
	v2453 = F_pg_ultostr_zeropad(m, v2442+v2179, v2182^v2449-v2449, int32(2))
	mBase = m.M
	goto L735
L739:
	;
	v2561 = v2453
	goto L734
L740:
	;
	goto L741
L741:
	;
	v2458 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v2453))) = uint8(v2458)
	v2461 = v2183 >> (uint(int32(31)) % 32)
	v2463 = v2183 ^ v2461 - v2461
	v2465 = base.I32_div_s(v2463, int32(10))
	v2468 = v2465*int32(-10) + v2463
	if v2468 != 0 {
		goto L743
	} else {
		goto L744
	}
L742:
	;
	v2478 = base.I32_div_s(v2463, int32(100))
	v2481 = v2478*int32(-10) + v2465
	v2482 = v2468 | v2481
	if v2482 == int32(0) {
		goto L747
	} else {
		goto L748
	}
L743:
	;
	v2470 = v2468 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2453)+6)) = uint8(v2470)
	v2476 = v2453 + int32(7)
	goto L742
L744:
	;
	goto L745
L745:
	;
	v2476 = v2453 + int32(6)
	goto L742
L746:
	;
	v2492 = base.I32_div_s(v2463, int32(1000))
	v2495 = v2492*int32(-10) + v2478
	v2496 = v2482 | v2495
	if v2496 == int32(0) {
		goto L751
	} else {
		goto L752
	}
L747:
	;
	v2490 = v2453 + int32(5)
	goto L746
L748:
	;
	goto L749
L749:
	;
	v2488 = v2481 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2453)+5)) = uint8(v2488)
	v2490 = v2476
	goto L746
L750:
	;
	v2506 = base.I32_div_s(v2463, int32(10000))
	v2509 = v2506*int32(-10) + v2492
	v2510 = v2496 | v2509
	if v2510 == int32(0) {
		goto L755
	} else {
		goto L756
	}
L751:
	;
	v2504 = v2453 + int32(4)
	goto L750
L752:
	;
	goto L753
L753:
	;
	v2502 = v2495 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2453)+4)) = uint8(v2502)
	v2504 = v2490
	goto L750
L754:
	;
	v2520 = base.I32_div_s(v2463, int32(100000))
	v2523 = v2520*int32(-10) + v2506
	v2524 = v2510 | v2523
	if v2524 == int32(0) {
		goto L759
	} else {
		goto L760
	}
L755:
	;
	v2518 = v2453 + int32(3)
	goto L754
L756:
	;
	goto L757
L757:
	;
	v2516 = v2509 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2453)+3)) = uint8(v2516)
	v2518 = v2504
	goto L754
L758:
	;
	v2534 = base.I32_div_s(v2463, int32(1000000))
	v2537 = v2534*int32(-10) + v2520
	if v2524|v2537 == int32(0) {
		goto L763
	} else {
		goto L764
	}
L759:
	;
	v2532 = v2453 + int32(2)
	goto L758
L760:
	;
	goto L761
L761:
	;
	v2530 = v2523 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2453)+2)) = uint8(v2530)
	v2532 = v2518
	goto L758
L762:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v2520+int32(9)) {
		goto L766
	} else {
		goto L767
	}
L763:
	;
	v2546 = v2453 + int32(1)
	goto L762
L764:
	;
	goto L765
L765:
	;
	v2544 = v2537 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2453)+1)) = uint8(v2544)
	v2546 = v2532
	goto L762
L766:
	;
	v2553 = F_pg_ultostr(m, v2453+int32(1), v2463)
	mBase = m.M
	v2554 = v2553
	goto L768
L767:
	;
	v2554 = v2546
	goto L768
L768:
	;
	v2561 = v2554
	goto L734
L769:
	;
	m.G0 = v20 + int32(176)
	return v2598
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
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v74
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v76
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
							v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v74
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v76
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
									v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
									*(*int64)(unsafe.Add(mBase, uint32(v8))) = v74
									v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v76
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
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v74
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v76
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
										v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
										*(*int64)(unsafe.Add(mBase, uint32(v8))) = v74
										v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v76
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
									v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
									*(*int64)(unsafe.Add(mBase, uint32(v8))) = v74
									v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v76
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
											F_errmsg(m, int32(398189), int32(0))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(490442), int32(3540), int32(298090))
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
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v74
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v76
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
							v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v74
							v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v76
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
									v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
									*(*int64)(unsafe.Add(mBase, uint32(v8))) = v74
									v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v76
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
								v74 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v74
								v76 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v76
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
										F_errmsg(m, int32(398189), int32(0))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(490442), int32(3549), int32(298090))
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
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v68 int64
	_ = v68
	var v71 float64
	_ = v71
	var v75 float64
	_ = v75
	var v80 float64
	_ = v80
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
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
	if base.F64_eq(v14, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
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
				F_errmsg(m, int32(398189), int32(0))
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(490442), int32(1578), int32(306300))
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
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v14)) {
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
					F_errmsg(m, int32(398189), int32(0))
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(490442), int32(1578), int32(306300))
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
			v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+60)))
			v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+52)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v25 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
			v27 = F_palloc(m, int32(16))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v32 = v25 * int64(12)
				v33 = base.I32_wrap_i64(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v33
				if base.I32_wrap_i64(int64(base.Ui64(v32)>>(uint(int64(32))%64))) != v33>>(uint(int32(31))%32) {
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
							F_errmsg(m, int32(398189), int32(0))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(490442), int32(1578), int32(306300))
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
					v41 = v33 + v24
					*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v41
					if base.B2i32(v24 < int32(0))^base.B2i32(v41 < v33) != 0 {
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
								F_errmsg(m, int32(398189), int32(0))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(490442), int32(1578), int32(306300))
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
						v49 = base.I64_extend_i32_s(v23) * int64(7)
						v50 = base.I32_wrap_i64(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v50
						if base.I32_wrap_i64(int64(base.Ui64(v49)>>(uint(int64(32))%64))) != v50>>(uint(int32(31))%32) {
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
									F_errmsg(m, int32(398189), int32(0))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490442), int32(1578), int32(306300))
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
							v58 = v50 + v22
							*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v58
							if base.B2i32(v22 < int32(0))^base.B2i32(v58 < v50) != 0 {
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
										F_errmsg(m, int32(398189), int32(0))
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(490442), int32(1578), int32(306300))
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
								v68 = v20*int64(60000000) + v21*int64(3600000000)
								*(*int64)(unsafe.Add(mBase, uint32(v27))) = v68
								v71 = base.F64_mul(v13, float64(1e+06))
								if base.F64_eq(base.F64_abs(v71), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
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
									v75 = float64(0)
									if base.F64_eq(v71, v75)&base.F64_ne(v13, v75) != 0 {
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
										v80 = base.F64_nearest(v71)
										if base.F64_ge(v80, float64(-9.223372036854776e+18)) == int32(0) {
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
													F_errmsg(m, int32(398189), int32(0))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(490442), int32(1578), int32(306300))
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
											if base.F64_lt(v80, float64(9.223372036854776e+18)) == int32(0) {
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
														F_errmsg(m, int32(398189), int32(0))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(490442), int32(1578), int32(306300))
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
												if base.F64_lt(base.F64_abs(v80), float64(9.223372036854776e+18)) != 0 {
													v92 = base.I64_trunc_f64_s(v80)
													v94 = v92
												} else {
													v94 = int64(-9223372036854775807 - 1)
												}
												v95 = v94 + v68
												*(*int64)(unsafe.Add(mBase, uint32(v27))) = v95
												if base.B2i32(v94 < int64(0))^base.B2i32(v95 < v68) != 0 {
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
															F_errmsg(m, int32(398189), int32(0))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(490442), int32(1578), int32(306300))
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
													if v41 != int32(2147483647) {
														if v41 != int32(-2147483648) {
															return v27
														} else {
															if v58 != int32(-2147483648) {
																return v27
															} else {
																if v95 == int64(-9223372036854775807-1) {
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
																			F_errmsg(m, int32(398189), int32(0))
																			mBase = m.M
																			v133 = m.ExcPending
																			if v133 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(490442), int32(1578), int32(306300))
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
																	return v27
																}
															}
														}
													} else {
														if v58 != int32(2147483647) {
															return v27
														} else {
															if v95 != int64(9223372036854775807) {
																return v27
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
																		F_errmsg(m, int32(398189), int32(0))
																		mBase = m.M
																		v133 = m.ExcPending
																		if v133 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(490442), int32(1578), int32(306300))
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
		}
	}
}
