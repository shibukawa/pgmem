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
											F_errmsg(m, int32(402870), int32(0))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496282), int32(3596), int32(320082))
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
										F_errmsg(m, int32(402870), int32(0))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496282), int32(3605), int32(320082))
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
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v684 int32
	_ = v684
	var v687 int64
	_ = v687
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v865 int64
	_ = v865
	var v869 int64
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v881 int64
	_ = v881
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v909 int64
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v918 int64
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v923 int64
	_ = v923
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v953 int64
	_ = v953
	var v958 int64
	_ = v958
	var v962 int64
	_ = v962
	var v967 int64
	_ = v967
	var v973 int64
	_ = v973
	var v978 int64
	_ = v978
	var v979 int32
	_ = v979
	var v981 int64
	_ = v981
	var v989 int32
	_ = v989
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1005 int64
	_ = v1005
	var v1007 int64
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1012 int64
	_ = v1012
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int64
	_ = v1232
	var v1233 int64
	_ = v1233
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1480 int32
	_ = v1480
	var v1482 int32
	_ = v1482
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
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
	v1534 = F_pstrdup(m, v20)
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L28
	} else {
		goto L446
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
	v85 = *(*int32)(unsafe.Add(mBase, _consts[855]))
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
	v50 = int32(*(*uint8)(unsafe.Add(mBase, _consts[860])))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v50)
	v53 = *(*int64)(unsafe.Add(mBase, _consts[861]))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v53
	goto L1
L4:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, _consts[862])))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+8)) = uint16(v44)
	v47 = *(*int64)(unsafe.Add(mBase, _consts[863]))
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
		goto L367
	} else {
		goto L368
	}
L16:
	;
	v832 = int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v20))) = uint16(v832)
	v835 = v20 + int32(1)
	if v97 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L17:
	;
	if v97 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L18:
	;
	if v97 != 0 {
		goto L79
	} else {
		goto L80
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
	v1226 = v163 - v96
	v1227 = v20 + int32(1)
	v1228 = v163 - v97
	v1229 = v163 - v94
	v1230 = v163 - v93
	v1231 = v163 - v92
	v1232 = v169 - v95
	v1233 = v169 - v99
	goto L15
L25:
	;
	v181 = F_pg_sprintf(m, v20, int32(571040), int32(0))
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
		v1226 = v96
		v1227 = v20
		v1228 = v97
		v1229 = v94
		v1230 = v93
		v1231 = v92
		v1232 = v95
		v1233 = v99
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
	v237 = F_pg_sprintf(m, v20, int32(547811), v88+int32(160))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L28
	} else {
		goto L43
	}
L43:
	;
	v239 = F_strlen(m, v20)
	mBase = m.M
	v246 = v93 >> (uint(int32(31)) % 32)
	goto L46
L44:
	;
	v359 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v358))) = uint8(v359)
	goto L14
L45:
	;
	if v92 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v250 = F_pg_ultostr_zeropad(m, v239+v20, v93^v246-v246, int32(2))
	mBase = m.M
	goto L45
L49:
	;
	v358 = v250
	goto L44
L50:
	;
	goto L51
L51:
	;
	v255 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v250))) = uint8(v255)
	v258 = v92 >> (uint(int32(31)) % 32)
	v260 = v92 ^ v258 - v258
	v262 = base.I32_div_s(v260, int32(10))
	v265 = v262*int32(-10) + v260
	if v265 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v275 = base.I32_div_s(v260, int32(100))
	v278 = v275*int32(-10) + v262
	v279 = v265 | v278
	if v279 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v267 = v265 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v250)+6)) = uint8(v267)
	v273 = v250 + int32(7)
	goto L52
L54:
	;
	goto L55
L55:
	;
	v273 = v250 + int32(6)
	goto L52
L56:
	;
	v289 = base.I32_div_s(v260, int32(1000))
	v292 = v289*int32(-10) + v275
	v293 = v279 | v292
	if v293 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v287 = v250 + int32(5)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v285 = v278 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v250)+5)) = uint8(v285)
	v287 = v273
	goto L56
L60:
	;
	v303 = base.I32_div_s(v260, int32(10000))
	v306 = v303*int32(-10) + v289
	v307 = v293 | v306
	if v307 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	v301 = v250 + int32(4)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v299 = v292 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v250)+4)) = uint8(v299)
	v301 = v287
	goto L60
L64:
	;
	v317 = base.I32_div_s(v260, int32(100000))
	v320 = v317*int32(-10) + v303
	v321 = v307 | v320
	if v321 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v315 = v250 + int32(3)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v313 = v306 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v250)+3)) = uint8(v313)
	v315 = v301
	goto L64
L68:
	;
	v331 = base.I32_div_s(v260, int32(1000000))
	v334 = v331*int32(-10) + v317
	if v321|v334 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	v329 = v250 + int32(2)
	goto L68
L70:
	;
	goto L71
L71:
	;
	v327 = v320 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v250)+2)) = uint8(v327)
	v329 = v315
	goto L68
L72:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v317+int32(9)) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v343 = v250 + int32(1)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v341 = v334 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v250)+1)) = uint8(v341)
	v343 = v329
	goto L72
L76:
	;
	v350 = F_pg_ultostr(m, v250+int32(1), v260)
	mBase = m.M
	v351 = v350
	goto L78
L77:
	;
	v351 = v343
	goto L78
L78:
	;
	v358 = v351
	goto L44
L79:
	;
	v367 = int32(80)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v367)
	v370 = v20 + int32(1)
	if v97 != 0 {
		goto L88
	} else {
		goto L89
	}
L80:
	;
	if v96 != 0 {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	if v98 != 0 {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	if v95 != int64(0) {
		goto L79
	} else {
		goto L83
	}
L83:
	;
	if v94 != 0 {
		goto L79
	} else {
		goto L84
	}
L84:
	;
	if v93 != 0 {
		goto L79
	} else {
		goto L85
	}
L85:
	;
	if v92 != 0 {
		goto L79
	} else {
		goto L86
	}
L86:
	;
	v365 = F_pg_sprintf(m, v20, int32(524699), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L28
	} else {
		goto L87
	}
L87:
	;
	goto L14
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+280)) = int32(89)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+272)) = base.I64_extend_i32_s(v97)
	v378 = F_pg_sprintf(m, v370, int32(500700), v88+int32(272))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L28
	} else {
		goto L91
	}
L89:
	;
	v382 = v370
	goto L90
L90:
	;
	if v96 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v380 = F_strlen(m, v370)
	mBase = m.M
	v382 = v380 + v370
	goto L90
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+264)) = int32(77)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+256)) = base.I64_extend_i32_s(v96)
	v390 = F_pg_sprintf(m, v382, int32(500700), v88+int32(256))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L28
	} else {
		goto L95
	}
L93:
	;
	v394 = v382
	goto L94
L94:
	;
	if v98 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v392 = F_strlen(m, v382)
	mBase = m.M
	v394 = v392 + v382
	goto L94
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+248)) = int32(68)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+240)) = v99
	v401 = F_pg_sprintf(m, v394, int32(500700), v88+int32(240))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L28
	} else {
		goto L99
	}
L97:
	;
	v405 = v394
	goto L98
L98:
	;
	if v95 != int64(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v403 = F_strlen(m, v394)
	mBase = m.M
	v405 = v403 + v394
	goto L98
L100:
	;
	if v93|v92 == int32(0) {
		goto L14
	} else {
		goto L112
	}
L101:
	;
	v410 = int32(84)
	*(*uint8)(unsafe.Add(mBase, uint32(v405))) = uint8(v410)
	v413 = v405 + int32(1)
	if v95 != int64(0) {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	if v94 != 0 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	if v93 != 0 {
		goto L101
	} else {
		goto L104
	}
L104:
	;
	if v92 == int32(0) {
		v440 = v405
		goto L100
	} else {
		goto L105
	}
L105:
	;
	goto L101
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+232)) = int32(72)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+224)) = v95
	v422 = F_pg_sprintf(m, v413, int32(500700), v88+int32(224))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L28
	} else {
		goto L109
	}
L107:
	;
	v426 = v413
	goto L108
L108:
	;
	if v94 == int32(0) {
		v440 = v426
		goto L100
	} else {
		goto L110
	}
L109:
	;
	v424 = F_strlen(m, v413)
	mBase = m.M
	v426 = v424 + v413
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+216)) = int32(77)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+208)) = base.I64_extend_i32_s(v94)
	v436 = F_pg_sprintf(m, v426, int32(500700), v88+int32(208))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L28
	} else {
		goto L111
	}
L111:
	;
	v438 = F_strlen(m, v426)
	mBase = m.M
	v440 = v438 + v426
	goto L100
L112:
	;
	if v93|v92 < int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v447 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v447)
	v451 = v440 + int32(1)
	goto L115
L114:
	;
	v451 = v440
	goto L115
L115:
	;
	v457 = v93 >> (uint(int32(31)) % 32)
	goto L119
L116:
	;
	v570 = int32(83)
	*(*uint16)(unsafe.Add(mBase, uint32(v569))) = uint16(v570)
	goto L14
L117:
	;
	if v92 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v462 = F_pg_ultostr(m, v451, v93^v457-v457)
	mBase = m.M
	goto L117
L121:
	;
	v569 = v462
	goto L116
L122:
	;
	goto L123
L123:
	;
	v466 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v462))) = uint8(v466)
	v469 = v92 >> (uint(int32(31)) % 32)
	v471 = v92 ^ v469 - v469
	v473 = base.I32_div_s(v471, int32(10))
	v476 = v473*int32(-10) + v471
	if v476 != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v486 = base.I32_div_s(v471, int32(100))
	v489 = v486*int32(-10) + v473
	v490 = v476 | v489
	if v490 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v478 = v476 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v462)+6)) = uint8(v478)
	v484 = v462 + int32(7)
	goto L124
L126:
	;
	goto L127
L127:
	;
	v484 = v462 + int32(6)
	goto L124
L128:
	;
	v500 = base.I32_div_s(v471, int32(1000))
	v503 = v500*int32(-10) + v486
	v504 = v490 | v503
	if v504 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	v498 = v462 + int32(5)
	goto L128
L130:
	;
	goto L131
L131:
	;
	v496 = v489 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v462)+5)) = uint8(v496)
	v498 = v484
	goto L128
L132:
	;
	v514 = base.I32_div_s(v471, int32(10000))
	v517 = v514*int32(-10) + v500
	v518 = v504 | v517
	if v518 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	v512 = v462 + int32(4)
	goto L132
L134:
	;
	goto L135
L135:
	;
	v510 = v503 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v462)+4)) = uint8(v510)
	v512 = v498
	goto L132
L136:
	;
	v528 = base.I32_div_s(v471, int32(100000))
	v531 = v528*int32(-10) + v514
	v532 = v518 | v531
	if v532 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L137:
	;
	v526 = v462 + int32(3)
	goto L136
L138:
	;
	goto L139
L139:
	;
	v524 = v517 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v462)+3)) = uint8(v524)
	v526 = v512
	goto L136
L140:
	;
	v542 = base.I32_div_s(v471, int32(1000000))
	v545 = v542*int32(-10) + v528
	if v532|v545 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L141:
	;
	v540 = v462 + int32(2)
	goto L140
L142:
	;
	goto L143
L143:
	;
	v538 = v531 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v462)+2)) = uint8(v538)
	v540 = v526
	goto L140
L144:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v528+int32(9)) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v554 = v462 + int32(1)
	goto L144
L146:
	;
	goto L147
L147:
	;
	v552 = v545 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v462)+1)) = uint8(v552)
	v554 = v540
	goto L144
L148:
	;
	v561 = F_pg_ultostr(m, v462+int32(1), v471)
	mBase = m.M
	v562 = v561
	goto L150
L149:
	;
	v562 = v554
	goto L150
L150:
	;
	v569 = v562
	goto L116
L151:
	;
	if v96 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L152:
	;
	v598 = v20
	v599 = int32(0)
	goto L151
L153:
	;
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+400)) = int32(230552)
	if v97 == int32(1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v581 = int32(757756)
	goto L157
L156:
	;
	v581 = int32(206261)
	goto L157
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+404)) = v581
	v583 = int32(757756)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+388)) = v583
	*(*int32)(unsafe.Add(mBase, uint32(v88)+384)) = v583
	*(*int64)(unsafe.Add(mBase, uint32(v88)+392)) = base.I64_extend_i32_s(v97)
	v592 = F_pg_sprintf(m, v20, int32(175773), v88+int32(384))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L28
	} else {
		goto L158
	}
L158:
	;
	v596 = F_strlen(m, v20)
	mBase = m.M
	v598 = v596 + v20
	v599 = int32(base.Ui32(v97) >> (uint(int32(31)) % 32))
	goto L151
L159:
	;
	if v98 != 0 {
		goto L176
	} else {
		goto L177
	}
L160:
	;
	v636 = v598
	v637 = base.B2i32(v97 == int32(0))
	v638 = v599
	goto L159
L161:
	;
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+368)) = int32(246177)
	if v96 == int32(1) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v610 = int32(757756)
	goto L165
L164:
	;
	v610 = int32(206261)
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+372)) = v610
	*(*int64)(unsafe.Add(mBase, uint32(v88)+360)) = base.I64_extend_i32_s(v96)
	if v97 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v616 = int32(746695)
	goto L168
L167:
	;
	v616 = int32(757756)
	goto L168
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+352)) = v616
	v619 = int32(757756)
	if v599 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v622 = int32(670327)
	goto L171
L170:
	;
	v622 = v619
	goto L171
L171:
	;
	if v96 <= int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v625 = v619
	goto L174
L173:
	;
	v625 = v622
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+356)) = v625
	v630 = F_pg_sprintf(m, v598, int32(175773), v88+int32(352))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L28
	} else {
		goto L175
	}
L175:
	;
	v634 = F_strlen(m, v598)
	mBase = m.M
	v636 = v634 + v598
	v637 = int32(0)
	v638 = int32(base.Ui32(v96) >> (uint(int32(31)) % 32))
	goto L159
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+336)) = int32(26930)
	if v98 == int32(1) {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	v670 = v636
	v671 = v637
	v672 = v638
	goto L178
L178:
	;
	if v671 != 0 {
		goto L192
	} else {
		goto L193
	}
L179:
	;
	v645 = int32(757756)
	goto L181
L180:
	;
	v645 = int32(206261)
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+340)) = v645
	*(*int64)(unsafe.Add(mBase, uint32(v88)+328)) = v99
	if v637 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v650 = int32(757756)
	goto L184
L183:
	;
	v650 = int32(746695)
	goto L184
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+320)) = v650
	v653 = int32(757756)
	if v638 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v656 = int32(670327)
	goto L187
L186:
	;
	v656 = v653
	goto L187
L187:
	;
	if v98 <= int32(0) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v659 = v653
	goto L190
L189:
	;
	v659 = v656
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+324)) = v659
	v664 = F_pg_sprintf(m, v636, int32(175773), v88+int32(320))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L28
	} else {
		goto L191
	}
L191:
	;
	v668 = F_strlen(m, v636)
	mBase = m.M
	v670 = v668 + v636
	v671 = int32(0)
	v672 = int32(base.Ui32(v98) >> (uint(int32(31)) % 32))
	goto L178
L192:
	;
	v678 = v94 >> (uint(int32(31)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+304)) = v94 ^ v678 - v678
	if v671 != 0 {
		goto L198
	} else {
		goto L199
	}
L193:
	;
	if v95 != int64(0) {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	if v94 != 0 {
		goto L192
	} else {
		goto L195
	}
L195:
	;
	if v93 != 0 {
		goto L192
	} else {
		goto L196
	}
L196:
	;
	if v92 == int32(0) {
		goto L14
	} else {
		goto L197
	}
L197:
	;
	goto L192
L198:
	;
	v684 = int32(757756)
	goto L200
L199:
	;
	v684 = int32(746695)
	goto L200
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+288)) = v684
	v687 = v95 >> (uint(int64(63)) % 64)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+296)) = v95 ^ v687 - v687
	v691 = int32(670315)
	if v672 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v695 = int32(670327)
	goto L203
L202:
	;
	v695 = int32(757756)
	goto L203
L203:
	;
	if v94|(v93|v92) < int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v700 = v691
	goto L206
L205:
	;
	v700 = v695
	goto L206
L206:
	;
	if v95 < int64(0) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v703 = v691
	goto L209
L208:
	;
	v703 = v700
	goto L209
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+292)) = v703
	v708 = F_pg_sprintf(m, v670, int32(547794), v88+int32(288))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L28
	} else {
		goto L210
	}
L210:
	;
	v710 = F_strlen(m, v670)
	mBase = m.M
	v717 = v93 >> (uint(int32(31)) % 32)
	goto L213
L211:
	;
	v830 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v829))) = uint8(v830)
	goto L14
L212:
	;
	if v92 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L213:
	;
	v721 = F_pg_ultostr_zeropad(m, v710+v670, v93^v717-v717, int32(2))
	mBase = m.M
	goto L212
L216:
	;
	v829 = v721
	goto L211
L217:
	;
	goto L218
L218:
	;
	v726 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v721))) = uint8(v726)
	v729 = v92 >> (uint(int32(31)) % 32)
	v731 = v92 ^ v729 - v729
	v733 = base.I32_div_s(v731, int32(10))
	v736 = v733*int32(-10) + v731
	if v736 != 0 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v746 = base.I32_div_s(v731, int32(100))
	v749 = v746*int32(-10) + v733
	v750 = v736 | v749
	if v750 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L220:
	;
	v738 = v736 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v721)+6)) = uint8(v738)
	v744 = v721 + int32(7)
	goto L219
L221:
	;
	goto L222
L222:
	;
	v744 = v721 + int32(6)
	goto L219
L223:
	;
	v760 = base.I32_div_s(v731, int32(1000))
	v763 = v760*int32(-10) + v746
	v764 = v750 | v763
	if v764 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L224:
	;
	v758 = v721 + int32(5)
	goto L223
L225:
	;
	goto L226
L226:
	;
	v756 = v749 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v721)+5)) = uint8(v756)
	v758 = v744
	goto L223
L227:
	;
	v774 = base.I32_div_s(v731, int32(10000))
	v777 = v774*int32(-10) + v760
	v778 = v764 | v777
	if v778 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L228:
	;
	v772 = v721 + int32(4)
	goto L227
L229:
	;
	goto L230
L230:
	;
	v770 = v763 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v721)+4)) = uint8(v770)
	v772 = v758
	goto L227
L231:
	;
	v788 = base.I32_div_s(v731, int32(100000))
	v791 = v788*int32(-10) + v774
	v792 = v778 | v791
	if v792 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L232:
	;
	v786 = v721 + int32(3)
	goto L231
L233:
	;
	goto L234
L234:
	;
	v784 = v777 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v721)+3)) = uint8(v784)
	v786 = v772
	goto L231
L235:
	;
	v802 = base.I32_div_s(v731, int32(1000000))
	v805 = v802*int32(-10) + v788
	if v792|v805 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L236:
	;
	v800 = v721 + int32(2)
	goto L235
L237:
	;
	goto L238
L238:
	;
	v798 = v791 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v721)+2)) = uint8(v798)
	v800 = v786
	goto L235
L239:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v788+int32(9)) {
		goto L243
	} else {
		goto L244
	}
L240:
	;
	v814 = v721 + int32(1)
	goto L239
L241:
	;
	goto L242
L242:
	;
	v812 = v805 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v721)+1)) = uint8(v812)
	v814 = v800
	goto L239
L243:
	;
	v821 = F_pg_ultostr(m, v721+int32(1), v731)
	mBase = m.M
	v822 = v821
	goto L245
L244:
	;
	v822 = v814
	goto L245
L245:
	;
	v829 = v822
	goto L211
L246:
	;
	v953 = int64(0)
	if v95 != v953 {
		goto L285
	} else {
		goto L286
	}
L247:
	;
	v947 = v945
	v949 = v942
	v952 = int32(0)
	goto L246
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+56)) = int32(26930)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+48)) = v923
	if v923 == int64(1) {
		goto L275
	} else {
		goto L276
	}
L249:
	;
	if v914 != 0 {
		goto L272
	} else {
		goto L273
	}
L250:
	;
	v902 = int32(0)
	if v98 == v902 {
		v947 = v900
		v949 = v901
		v952 = base.B2i32(v97 == v902)
		goto L246
	} else {
		goto L270
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+72)) = int32(246177)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+64)) = v881
	if v881 == int64(1) {
		goto L265
	} else {
		goto L266
	}
L252:
	;
	v870 = int32(31)
	v873 = v96 >> (uint(v870) % 32)
	v877 = v835
	v879 = int32(base.Ui32(v96) >> (uint(v870) % 32))
	v881 = base.I64_extend_i32_u(v96 ^ v873 - v873)
	goto L251
L253:
	;
	if v96 != 0 {
		goto L252
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+88)) = int32(230552)
	v842 = v97 >> (uint(int32(31)) % 32)
	v844 = v97 ^ v842 - v842
	*(*int64)(unsafe.Add(mBase, uint32(v88)+80)) = base.I64_extend_i32_u(v844)
	if v844 == int32(1) {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v900 = v835
	v901 = int32(0)
	goto L250
L257:
	;
	v851 = int32(757756)
	goto L259
L258:
	;
	v851 = int32(206261)
	goto L259
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+92)) = v851
	v856 = F_pg_sprintf(m, v835, int32(175787), v88+int32(80))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L28
	} else {
		goto L260
	}
L260:
	;
	v859 = int32(base.Ui32(v97) >> (uint(int32(31)) % 32))
	v860 = F_strlen(m, v835)
	mBase = m.M
	v861 = v860 + v835
	if v96 == int32(0) {
		v900 = v861
		v901 = v859
		goto L250
	} else {
		goto L261
	}
L261:
	;
	v865 = base.I64_extend_i32_s(v96)
	if v97 < int32(0) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v869 = int64(0) - v865
	goto L264
L263:
	;
	v869 = v865
	goto L264
L264:
	;
	v877 = v861
	v879 = v859
	v881 = v869
	goto L251
L265:
	;
	v889 = int32(757756)
	goto L267
L266:
	;
	v889 = int32(206261)
	goto L267
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+76)) = v889
	v894 = F_pg_sprintf(m, v877, int32(175787), v88-int32(-64))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L28
	} else {
		goto L268
	}
L268:
	;
	v896 = F_strlen(m, v877)
	mBase = m.M
	v897 = v896 + v877
	if v98 == int32(0) {
		v942 = v879
		v945 = v897
		goto L247
	} else {
		goto L269
	}
L269:
	;
	v912 = v897
	v914 = v879
	goto L249
L270:
	;
	if v97 != 0 {
		v912 = v900
		v914 = v901
		goto L249
	} else {
		goto L271
	}
L271:
	;
	v909 = v99 >> (uint(int64(63)) % 64)
	v919 = v900
	v921 = int32(base.Ui32(v98) >> (uint(int32(31)) % 32))
	v923 = v99 ^ v909 - v909
	goto L248
L272:
	;
	v918 = int64(0) - v99
	goto L274
L273:
	;
	v918 = v99
	goto L274
L274:
	;
	v919 = v912
	v921 = v914
	v923 = v918
	goto L248
L275:
	;
	v931 = int32(757756)
	goto L277
L276:
	;
	v931 = int32(206261)
	goto L277
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+60)) = v931
	v936 = F_pg_sprintf(m, v919, int32(175787), v88+int32(48))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L28
	} else {
		goto L278
	}
L278:
	;
	v938 = F_strlen(m, v919)
	mBase = m.M
	v942 = v921
	v945 = v938 + v919
	goto L247
L279:
	;
	if v93|v92 != 0 {
		goto L307
	} else {
		goto L308
	}
L280:
	;
	v1035 = v1033
	v1036 = int32(0)
	v1037 = v1030
	goto L279
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = int32(276549)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v1012
	if v1012 == int64(1) {
		goto L302
	} else {
		goto L303
	}
L282:
	;
	if v1002 != 0 {
		goto L299
	} else {
		goto L300
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+40)) = int32(206279)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+32)) = v981
	if v981 == int64(1) {
		goto L294
	} else {
		goto L295
	}
L284:
	;
	if v949 != 0 {
		goto L291
	} else {
		goto L292
	}
L285:
	;
	if v952 == int32(0) {
		goto L284
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	if v94 == int32(0) {
		v1035 = v947
		v1036 = v952
		v1037 = v949
		goto L279
	} else {
		goto L289
	}
L288:
	;
	v958 = int64(63)
	v962 = v95 >> (uint(v958) % 64)
	v979 = base.I32_wrap_i64(int64(base.Ui64(v95) >> (uint(v958) % 64)))
	v981 = v95 ^ v962 - v962
	goto L283
L289:
	;
	v967 = base.I64_extend_i32_s(v94)
	if v952 == int32(0) {
		v1001 = v947
		v1002 = v949
		v1005 = v967
		goto L282
	} else {
		goto L290
	}
L290:
	;
	v973 = v967 >> (uint(int64(63)) % 64)
	v1008 = v947
	v1009 = int32(base.Ui32(v94) >> (uint(int32(31)) % 32))
	v1012 = v967 ^ v973 - v973
	goto L281
L291:
	;
	v978 = int64(0) - v95
	goto L293
L292:
	;
	v978 = v95
	goto L293
L293:
	;
	v979 = v949
	v981 = v978
	goto L283
L294:
	;
	v989 = int32(757756)
	goto L296
L295:
	;
	v989 = int32(206261)
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+44)) = v989
	v994 = F_pg_sprintf(m, v947, int32(175787), v88+int32(32))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L28
	} else {
		goto L297
	}
L297:
	;
	v996 = F_strlen(m, v947)
	mBase = m.M
	v997 = v996 + v947
	if v94 == int32(0) {
		v1030 = v979
		v1033 = v997
		goto L280
	} else {
		goto L298
	}
L298:
	;
	v1001 = v997
	v1002 = v979
	v1005 = base.I64_extend_i32_s(v94)
	goto L282
L299:
	;
	v1007 = v953 - v1005
	goto L301
L300:
	;
	v1007 = v1005
	goto L301
L301:
	;
	v1008 = v1001
	v1009 = v1002
	v1012 = v1007
	goto L281
L302:
	;
	v1020 = int32(757756)
	goto L304
L303:
	;
	v1020 = int32(206261)
	goto L304
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = v1020
	v1025 = F_pg_sprintf(m, v1008, int32(175787), v88+int32(16))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L28
	} else {
		goto L305
	}
L305:
	;
	v1027 = F_strlen(m, v1008)
	mBase = m.M
	v1030 = v1009
	v1033 = v1027 + v1008
	goto L280
L306:
	;
	v1218 = F_strlen(m, v1214)
	mBase = m.M
	v1219 = v1218 + v1214
	v1221 = *(*int32)(unsafe.Add(mBase, _consts[856]))
	*(*int32)(unsafe.Add(mBase, uint32(v1219))) = v1221
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, _consts[857])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1219)+4)) = uint8(v1224)
	goto L14
L307:
	;
	v1041 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1035))) = uint8(v1041)
	v1043 = int32(1)
	v1045 = v1035 + v1043
	v1046 = int32(0)
	if v1046 <= v93 {
		goto L313
	} else {
		goto L314
	}
L308:
	;
	goto L309
L309:
	;
	if v1036 != 0 {
		goto L363
	} else {
		goto L364
	}
L310:
	;
	v1074 = v93 >> (uint(int32(31)) % 32)
	goto L323
L311:
	;
	v1062 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v1035)+1)) = uint8(v1062)
	v1067 = v1060
	v1068 = v1035 + int32(2)
	goto L310
L312:
	;
	v1056 = int32(0)
	if v1037 == v1056 {
		v1067 = v1056
		v1068 = v1045
		goto L310
	} else {
		goto L319
	}
L313:
	;
	if v93 != 0 {
		goto L312
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	if (v1036|v1037)&int32(1) == int32(0) {
		v1060 = v1046
		goto L311
	} else {
		goto L318
	}
L316:
	;
	if int32(0) <= v92 {
		goto L312
	} else {
		goto L317
	}
L317:
	;
	goto L315
L318:
	;
	v1067 = v1043
	v1068 = v1045
	goto L310
L319:
	;
	v1060 = int32(1)
	goto L311
L320:
	;
	v1187 = int32(206261)
	if v92 != 0 {
		goto L355
	} else {
		goto L356
	}
L321:
	;
	if v92 == int32(0) {
		goto L325
	} else {
		goto L326
	}
L323:
	;
	goto L324
L324:
	;
	v1079 = F_pg_ultostr(m, v1068, v93^v1074-v1074)
	mBase = m.M
	goto L321
L325:
	;
	v1186 = v1079
	goto L320
L326:
	;
	goto L327
L327:
	;
	v1083 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1079))) = uint8(v1083)
	v1086 = v92 >> (uint(int32(31)) % 32)
	v1088 = v92 ^ v1086 - v1086
	v1090 = base.I32_div_s(v1088, int32(10))
	v1093 = v1090*int32(-10) + v1088
	if v1093 != 0 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	v1103 = base.I32_div_s(v1088, int32(100))
	v1106 = v1103*int32(-10) + v1090
	v1107 = v1093 | v1106
	if v1107 == int32(0) {
		goto L333
	} else {
		goto L334
	}
L329:
	;
	v1095 = v1093 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1079)+6)) = uint8(v1095)
	v1101 = v1079 + int32(7)
	goto L328
L330:
	;
	goto L331
L331:
	;
	v1101 = v1079 + int32(6)
	goto L328
L332:
	;
	v1117 = base.I32_div_s(v1088, int32(1000))
	v1120 = v1117*int32(-10) + v1103
	v1121 = v1107 | v1120
	if v1121 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L333:
	;
	v1115 = v1079 + int32(5)
	goto L332
L334:
	;
	goto L335
L335:
	;
	v1113 = v1106 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1079)+5)) = uint8(v1113)
	v1115 = v1101
	goto L332
L336:
	;
	v1131 = base.I32_div_s(v1088, int32(10000))
	v1134 = v1131*int32(-10) + v1117
	v1135 = v1121 | v1134
	if v1135 == int32(0) {
		goto L341
	} else {
		goto L342
	}
L337:
	;
	v1129 = v1079 + int32(4)
	goto L336
L338:
	;
	goto L339
L339:
	;
	v1127 = v1120 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1079)+4)) = uint8(v1127)
	v1129 = v1115
	goto L336
L340:
	;
	v1145 = base.I32_div_s(v1088, int32(100000))
	v1148 = v1145*int32(-10) + v1131
	v1149 = v1135 | v1148
	if v1149 == int32(0) {
		goto L345
	} else {
		goto L346
	}
L341:
	;
	v1143 = v1079 + int32(3)
	goto L340
L342:
	;
	goto L343
L343:
	;
	v1141 = v1134 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1079)+3)) = uint8(v1141)
	v1143 = v1129
	goto L340
L344:
	;
	v1159 = base.I32_div_s(v1088, int32(1000000))
	v1162 = v1159*int32(-10) + v1145
	if v1149|v1162 == int32(0) {
		goto L349
	} else {
		goto L350
	}
L345:
	;
	v1157 = v1079 + int32(2)
	goto L344
L346:
	;
	goto L347
L347:
	;
	v1155 = v1148 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1079)+2)) = uint8(v1155)
	v1157 = v1143
	goto L344
L348:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v1145+int32(9)) {
		goto L352
	} else {
		goto L353
	}
L349:
	;
	v1171 = v1079 + int32(1)
	goto L348
L350:
	;
	goto L351
L351:
	;
	v1169 = v1162 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1079)+1)) = uint8(v1169)
	v1171 = v1157
	goto L348
L352:
	;
	v1178 = F_pg_ultostr(m, v1079+int32(1), v1088)
	mBase = m.M
	v1179 = v1178
	goto L354
L353:
	;
	v1179 = v1171
	goto L354
L354:
	;
	v1186 = v1179
	goto L320
L355:
	;
	v1190 = v1187
	goto L357
L356:
	;
	v1190 = int32(757756)
	goto L357
L357:
	;
	v1192 = v93 >> (uint(int32(31)) % 32)
	if v93^v1192-v1192 != int32(1) {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1197 = v1187
	goto L360
L359:
	;
	v1197 = v1190
	goto L360
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v1197
	v1200 = F_pg_sprintf(m, v1186, int32(175903), v88)
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L28
	} else {
		goto L361
	}
L361:
	;
	if v1067 != 0 {
		v1214 = v1186
		goto L306
	} else {
		goto L362
	}
L362:
	;
	goto L14
L363:
	;
	v1202 = F_strlen(m, v1035)
	mBase = m.M
	v1203 = v1202 + v1035
	v1205 = int32(*(*uint16)(unsafe.Add(mBase, _consts[858])))
	*(*uint16)(unsafe.Add(mBase, uint32(v1203))) = uint16(v1205)
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, _consts[859])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1203)+2)) = uint8(v1208)
	goto L365
L364:
	;
	goto L365
L365:
	;
	if v1037 == int32(0) {
		goto L14
	} else {
		goto L366
	}
L366:
	;
	v1214 = v1035
	goto L306
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+100)) = v1226
	*(*int32)(unsafe.Add(mBase, uint32(v88)+96)) = v1228
	v1239 = F_pg_sprintf(m, v1227, int32(467246), v88+int32(96))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L28
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	if v98 != 0 {
		goto L371
	} else {
		goto L372
	}
L370:
	;
	goto L14
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+128)) = v1229
	*(*int64)(unsafe.Add(mBase, uint32(v88)+120)) = v1232
	*(*int64)(unsafe.Add(mBase, uint32(v88)+112)) = v1233
	v1247 = F_pg_sprintf(m, v1227, int32(547839), v88+int32(112))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L28
	} else {
		goto L374
	}
L372:
	;
	goto L373
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+152)) = v1229
	*(*int64)(unsafe.Add(mBase, uint32(v88)+144)) = v1232
	v1376 = F_pg_sprintf(m, v1227, int32(547844), v88+int32(144))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L28
	} else {
		goto L410
	}
L374:
	;
	v1249 = F_strlen(m, v1227)
	mBase = m.M
	v1256 = v1230 >> (uint(int32(31)) % 32)
	goto L377
L375:
	;
	v1369 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1368))) = uint8(v1369)
	goto L14
L376:
	;
	if v1231 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L377:
	;
	v1260 = F_pg_ultostr_zeropad(m, v1249+v1227, v1230^v1256-v1256, int32(2))
	mBase = m.M
	goto L376
L380:
	;
	v1368 = v1260
	goto L375
L381:
	;
	goto L382
L382:
	;
	v1265 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1260))) = uint8(v1265)
	v1268 = v1231 >> (uint(int32(31)) % 32)
	v1270 = v1231 ^ v1268 - v1268
	v1272 = base.I32_div_s(v1270, int32(10))
	v1275 = v1272*int32(-10) + v1270
	if v1275 != 0 {
		goto L384
	} else {
		goto L385
	}
L383:
	;
	v1285 = base.I32_div_s(v1270, int32(100))
	v1288 = v1285*int32(-10) + v1272
	v1289 = v1275 | v1288
	if v1289 == int32(0) {
		goto L388
	} else {
		goto L389
	}
L384:
	;
	v1277 = v1275 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1260)+6)) = uint8(v1277)
	v1283 = v1260 + int32(7)
	goto L383
L385:
	;
	goto L386
L386:
	;
	v1283 = v1260 + int32(6)
	goto L383
L387:
	;
	v1299 = base.I32_div_s(v1270, int32(1000))
	v1302 = v1299*int32(-10) + v1285
	v1303 = v1289 | v1302
	if v1303 == int32(0) {
		goto L392
	} else {
		goto L393
	}
L388:
	;
	v1297 = v1260 + int32(5)
	goto L387
L389:
	;
	goto L390
L390:
	;
	v1295 = v1288 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1260)+5)) = uint8(v1295)
	v1297 = v1283
	goto L387
L391:
	;
	v1313 = base.I32_div_s(v1270, int32(10000))
	v1316 = v1313*int32(-10) + v1299
	v1317 = v1303 | v1316
	if v1317 == int32(0) {
		goto L396
	} else {
		goto L397
	}
L392:
	;
	v1311 = v1260 + int32(4)
	goto L391
L393:
	;
	goto L394
L394:
	;
	v1309 = v1302 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1260)+4)) = uint8(v1309)
	v1311 = v1297
	goto L391
L395:
	;
	v1327 = base.I32_div_s(v1270, int32(100000))
	v1330 = v1327*int32(-10) + v1313
	v1331 = v1317 | v1330
	if v1331 == int32(0) {
		goto L400
	} else {
		goto L401
	}
L396:
	;
	v1325 = v1260 + int32(3)
	goto L395
L397:
	;
	goto L398
L398:
	;
	v1323 = v1316 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1260)+3)) = uint8(v1323)
	v1325 = v1311
	goto L395
L399:
	;
	v1341 = base.I32_div_s(v1270, int32(1000000))
	v1344 = v1341*int32(-10) + v1327
	if v1331|v1344 == int32(0) {
		goto L404
	} else {
		goto L405
	}
L400:
	;
	v1339 = v1260 + int32(2)
	goto L399
L401:
	;
	goto L402
L402:
	;
	v1337 = v1330 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1260)+2)) = uint8(v1337)
	v1339 = v1325
	goto L399
L403:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v1327+int32(9)) {
		goto L407
	} else {
		goto L408
	}
L404:
	;
	v1353 = v1260 + int32(1)
	goto L403
L405:
	;
	goto L406
L406:
	;
	v1351 = v1344 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1260)+1)) = uint8(v1351)
	v1353 = v1339
	goto L403
L407:
	;
	v1360 = F_pg_ultostr(m, v1260+int32(1), v1270)
	mBase = m.M
	v1361 = v1360
	goto L409
L408:
	;
	v1361 = v1353
	goto L409
L409:
	;
	v1368 = v1361
	goto L375
L410:
	;
	v1378 = F_strlen(m, v1227)
	mBase = m.M
	v1385 = v1230 >> (uint(int32(31)) % 32)
	goto L413
L411:
	;
	v1498 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1497))) = uint8(v1498)
	goto L14
L412:
	;
	if v1231 == int32(0) {
		goto L416
	} else {
		goto L417
	}
L413:
	;
	v1389 = F_pg_ultostr_zeropad(m, v1378+v1227, v1230^v1385-v1385, int32(2))
	mBase = m.M
	goto L412
L416:
	;
	v1497 = v1389
	goto L411
L417:
	;
	goto L418
L418:
	;
	v1394 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1389))) = uint8(v1394)
	v1397 = v1231 >> (uint(int32(31)) % 32)
	v1399 = v1231 ^ v1397 - v1397
	v1401 = base.I32_div_s(v1399, int32(10))
	v1404 = v1401*int32(-10) + v1399
	if v1404 != 0 {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	v1414 = base.I32_div_s(v1399, int32(100))
	v1417 = v1414*int32(-10) + v1401
	v1418 = v1404 | v1417
	if v1418 == int32(0) {
		goto L424
	} else {
		goto L425
	}
L420:
	;
	v1406 = v1404 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1389)+6)) = uint8(v1406)
	v1412 = v1389 + int32(7)
	goto L419
L421:
	;
	goto L422
L422:
	;
	v1412 = v1389 + int32(6)
	goto L419
L423:
	;
	v1428 = base.I32_div_s(v1399, int32(1000))
	v1431 = v1428*int32(-10) + v1414
	v1432 = v1418 | v1431
	if v1432 == int32(0) {
		goto L428
	} else {
		goto L429
	}
L424:
	;
	v1426 = v1389 + int32(5)
	goto L423
L425:
	;
	goto L426
L426:
	;
	v1424 = v1417 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1389)+5)) = uint8(v1424)
	v1426 = v1412
	goto L423
L427:
	;
	v1442 = base.I32_div_s(v1399, int32(10000))
	v1445 = v1442*int32(-10) + v1428
	v1446 = v1432 | v1445
	if v1446 == int32(0) {
		goto L432
	} else {
		goto L433
	}
L428:
	;
	v1440 = v1389 + int32(4)
	goto L427
L429:
	;
	goto L430
L430:
	;
	v1438 = v1431 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1389)+4)) = uint8(v1438)
	v1440 = v1426
	goto L427
L431:
	;
	v1456 = base.I32_div_s(v1399, int32(100000))
	v1459 = v1456*int32(-10) + v1442
	v1460 = v1446 | v1459
	if v1460 == int32(0) {
		goto L436
	} else {
		goto L437
	}
L432:
	;
	v1454 = v1389 + int32(3)
	goto L431
L433:
	;
	goto L434
L434:
	;
	v1452 = v1445 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1389)+3)) = uint8(v1452)
	v1454 = v1440
	goto L431
L435:
	;
	v1470 = base.I32_div_s(v1399, int32(1000000))
	v1473 = v1470*int32(-10) + v1456
	if v1460|v1473 == int32(0) {
		goto L440
	} else {
		goto L441
	}
L436:
	;
	v1468 = v1389 + int32(2)
	goto L435
L437:
	;
	goto L438
L438:
	;
	v1466 = v1459 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1389)+2)) = uint8(v1466)
	v1468 = v1454
	goto L435
L439:
	;
	if base.Ui32(int32(19)) <= base.Ui32(v1456+int32(9)) {
		goto L443
	} else {
		goto L444
	}
L440:
	;
	v1482 = v1389 + int32(1)
	goto L439
L441:
	;
	goto L442
L442:
	;
	v1480 = v1473 + int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v1389)+1)) = uint8(v1480)
	v1482 = v1468
	goto L439
L443:
	;
	v1489 = F_pg_ultostr(m, v1389+int32(1), v1399)
	mBase = m.M
	v1490 = v1489
	goto L445
L444:
	;
	v1490 = v1482
	goto L445
L445:
	;
	v1497 = v1490
	goto L411
L446:
	;
	m.G0 = v20 + int32(176)
	return v1534
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
											F_errmsg(m, int32(402870), int32(0))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(496282), int32(3540), int32(301504))
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
										F_errmsg(m, int32(402870), int32(0))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496282), int32(3549), int32(301504))
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
				F_errmsg(m, int32(402870), int32(0))
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(496282), int32(1578), int32(309721))
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
					F_errmsg(m, int32(402870), int32(0))
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(496282), int32(1578), int32(309721))
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
							F_errmsg(m, int32(402870), int32(0))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(496282), int32(1578), int32(309721))
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
								F_errmsg(m, int32(402870), int32(0))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(496282), int32(1578), int32(309721))
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
									F_errmsg(m, int32(402870), int32(0))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(496282), int32(1578), int32(309721))
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
										F_errmsg(m, int32(402870), int32(0))
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(496282), int32(1578), int32(309721))
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
													F_errmsg(m, int32(402870), int32(0))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(496282), int32(1578), int32(309721))
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
														F_errmsg(m, int32(402870), int32(0))
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(496282), int32(1578), int32(309721))
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
															F_errmsg(m, int32(402870), int32(0))
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(496282), int32(1578), int32(309721))
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
																			F_errmsg(m, int32(402870), int32(0))
																			mBase = m.M
																			v133 = m.ExcPending
																			if v133 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(496282), int32(1578), int32(309721))
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
																		F_errmsg(m, int32(402870), int32(0))
																		mBase = m.M
																		v133 = m.ExcPending
																		if v133 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(496282), int32(1578), int32(309721))
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
