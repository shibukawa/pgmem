package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_cube_binary_union(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v4 = F_cube_union_v0(m, l0, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(base.Ui32(v8) >> (uint(int32(2)) % 32))
		return v4
	}
}
func F_g_cube_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 float64
	_ = v52
	var v56 int32
	_ = v56
	var v59 float64
	_ = v59
	var v63 float64
	_ = v63
	var v68 int32
	_ = v68
	var v73 float64
	_ = v73
	var v77 float64
	_ = v77
	var v80 float64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 float64
	_ = v95
	var v101 int32
	_ = v101
	var v104 float64
	_ = v104
	var v108 float64
	_ = v108
	var v121 float64
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v155 float64
	_ = v155
	var v158 int32
	_ = v158
	var v161 float64
	_ = v161
	var v165 float64
	_ = v165
	var v170 int32
	_ = v170
	var v175 float64
	_ = v175
	var v179 float64
	_ = v179
	var v182 float64
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v198 float64
	_ = v198
	var v203 int32
	_ = v203
	var v206 float64
	_ = v206
	var v210 float64
	_ = v210
	var v224 float64
	_ = v224
	var v227 float32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	v10 = float64(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v22 = F_pg_detoast_datum(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_cube_union_v0(m, v17, v22)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24 == int32(0) {
					v121 = v10
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					if v28 <= int32(0) {
						v121 = v10
					} else {
						v31 = int32(1)
						v34 = v24 + int32(8)
						if v28 == v31 {
							v86 = int32(0)
							v95 = float64(1)
						} else {
							v43 = int32(0)
							v46 = int32(0)
							v52 = float64(1)
							for {
								v56 = int32(3)
								v59 = *(*float64)(unsafe.Add(mBase, uint32(v34+(v43+v28)<<(uint(v56)%32))))
								v63 = *(*float64)(unsafe.Add(mBase, uint32(v34+v43<<(uint(v56)%32))))
								v68 = v43 | int32(1)
								v73 = *(*float64)(unsafe.Add(mBase, uint32(v34+(v68+v28)<<(uint(v56)%32))))
								v77 = *(*float64)(unsafe.Add(mBase, uint32(v34+v68<<(uint(v56)%32))))
								v80 = base.F64_mul(base.F64_mul(v52, base.F64_abs(base.F64_sub(v59, v63))), base.F64_abs(base.F64_sub(v73, v77)))
								v81 = int32(2)
								v82 = v43 + v81
								v84 = v46 + v81
								if v84 != v28&int32(2147483646) {
									v43 = v82
									v46 = v84
									v52 = v80
									continue
								} else {
									break
								}
								break
							}
							v86 = v82
							v95 = v80
						}
						if v28&v31 == int32(0) {
							v121 = v95
						} else {
							v101 = int32(3)
							v104 = *(*float64)(unsafe.Add(mBase, uint32(v34+(v86+v28)<<(uint(v101)%32))))
							v108 = *(*float64)(unsafe.Add(mBase, uint32(v34+v86<<(uint(v101)%32))))
							v121 = base.F64_mul(v95, base.F64_abs(base.F64_sub(v104, v108)))
						}
					}
				}
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v125 = F_pg_detoast_datum(m, v124)
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					if v125 == int32(0) {
						v224 = v10
					} else {
						v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
						if v129 <= int32(0) {
							v224 = v10
						} else {
							v132 = int32(1)
							v135 = v125 + int32(8)
							if v129 == v132 {
								v188 = int32(0)
								v198 = float64(1)
							} else {
								v143 = int32(0)
								v145 = v143
								v148 = v143
								v155 = float64(1)
								for {
									v158 = int32(3)
									v161 = *(*float64)(unsafe.Add(mBase, uint32(v135+(v145+v129)<<(uint(v158)%32))))
									v165 = *(*float64)(unsafe.Add(mBase, uint32(v135+v145<<(uint(v158)%32))))
									v170 = v145 | int32(1)
									v175 = *(*float64)(unsafe.Add(mBase, uint32(v135+(v170+v129)<<(uint(v158)%32))))
									v179 = *(*float64)(unsafe.Add(mBase, uint32(v135+v170<<(uint(v158)%32))))
									v182 = base.F64_mul(base.F64_mul(v155, base.F64_abs(base.F64_sub(v161, v165))), base.F64_abs(base.F64_sub(v175, v179)))
									v183 = int32(2)
									v184 = v145 + v183
									v186 = v148 + v183
									if v186 != v129&int32(2147483646) {
										v145 = v184
										v148 = v186
										v155 = v182
										continue
									} else {
										break
									}
									break
								}
								v188 = v184
								v198 = v182
							}
							if v129&v132 == int32(0) {
								v224 = v198
							} else {
								v203 = int32(3)
								v206 = *(*float64)(unsafe.Add(mBase, uint32(v135+(v188+v129)<<(uint(v203)%32))))
								v210 = *(*float64)(unsafe.Add(mBase, uint32(v135+v188<<(uint(v203)%32))))
								v224 = base.F64_mul(v198, base.F64_abs(base.F64_sub(v206, v210)))
							}
						}
					}
					v227 = base.F32_demote_f64(base.F64_sub(v121, v224))
					*(*float32)(unsafe.Add(mBase, uint32(v13))) = v227
					v230 = F_Float8GetDatum(m, base.F64_promote_f32(v227))
					mBase = m.M
					v231 = m.ExcPending
					if v231 != 0 {
						return int32(0)
					} else {
						return v230
					}
				}
			}
		}
	}
}
