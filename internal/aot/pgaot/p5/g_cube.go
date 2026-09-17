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
	var v11 float64
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 float64
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v93 float64
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 float64
	_ = v102
	var v103 float64
	_ = v103
	var v117 float64
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 float64
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 float64
	_ = v160
	var v161 float64
	_ = v161
	var v168 float64
	_ = v168
	var v169 float64
	_ = v169
	var v172 float64
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v191 float64
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 float64
	_ = v199
	var v200 float64
	_ = v200
	var v215 float64
	_ = v215
	var v218 float32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	v11 = float64(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		v23 = F_pg_detoast_datum(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = F_cube_union_v0(m, v18, v23)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v25 == int32(0) {
					v117 = v11
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					if v29 <= int32(0) {
						v117 = v11
					} else {
						v33 = v25 + int32(8)
						if v29 == int32(1) {
							v83 = int32(0)
							v93 = float64(1)
							v96 = int32(3)
							v98 = v33 + v83<<(uint(v96)%32)
							v102 = *(*float64)(unsafe.Add(mBase, uint32(v98+v29<<(uint(v96)%32))))
							v103 = *(*float64)(unsafe.Add(mBase, uint32(v98)))
							v117 = base.F64_mul(v93, base.F64_abs(base.F64_sub(v102, v103)))
						} else {
							v44 = int32(0)
							v48 = int32(0)
							v54 = float64(1)
							for {
								v57 = int32(3)
								v59 = v33 + v44<<(uint(v57)%32)
								v61 = v29 << (uint(v57) % 32)
								v63 = *(*float64)(unsafe.Add(mBase, uint32(v59+v61)))
								v64 = *(*float64)(unsafe.Add(mBase, uint32(v59)))
								v71 = *(*float64)(unsafe.Add(mBase, uint32(v59+int32(8)+v61)))
								v72 = *(*float64)(unsafe.Add(mBase, uint32(v59)+8))
								v75 = base.F64_mul(base.F64_mul(v54, base.F64_abs(base.F64_sub(v63, v64))), base.F64_abs(base.F64_sub(v71, v72)))
								v76 = int32(2)
								v77 = v44 + v76
								v79 = v48 + v76
								if v79 != v29&int32(2147483646) {
									v44 = v77
									v48 = v79
									v54 = v75
									continue
								} else {
									break
								}
								break
							}
							if v29&int32(1) == int32(0) {
								v117 = v75
							} else {
								v83 = v77
								v93 = v75
								v96 = int32(3)
								v98 = v33 + v83<<(uint(v96)%32)
								v102 = *(*float64)(unsafe.Add(mBase, uint32(v98+v29<<(uint(v96)%32))))
								v103 = *(*float64)(unsafe.Add(mBase, uint32(v98)))
								v117 = base.F64_mul(v93, base.F64_abs(base.F64_sub(v102, v103)))
							}
						}
					}
				}
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v121 = F_pg_detoast_datum(m, v120)
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					if v121 == int32(0) {
						v215 = v11
					} else {
						v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
						if v125 <= int32(0) {
							v215 = v11
						} else {
							v129 = v121 + int32(8)
							if v125 == int32(1) {
								v180 = int32(0)
								v191 = float64(1)
								v193 = int32(3)
								v195 = v129 + v180<<(uint(v193)%32)
								v199 = *(*float64)(unsafe.Add(mBase, uint32(v195+v125<<(uint(v193)%32))))
								v200 = *(*float64)(unsafe.Add(mBase, uint32(v195)))
								v215 = base.F64_mul(v191, base.F64_abs(base.F64_sub(v199, v200)))
							} else {
								v139 = int32(0)
								v141 = v139
								v145 = v139
								v152 = float64(1)
								for {
									v154 = int32(3)
									v156 = v129 + v141<<(uint(v154)%32)
									v158 = v125 << (uint(v154) % 32)
									v160 = *(*float64)(unsafe.Add(mBase, uint32(v156+v158)))
									v161 = *(*float64)(unsafe.Add(mBase, uint32(v156)))
									v168 = *(*float64)(unsafe.Add(mBase, uint32(v156+int32(8)+v158)))
									v169 = *(*float64)(unsafe.Add(mBase, uint32(v156)+8))
									v172 = base.F64_mul(base.F64_mul(v152, base.F64_abs(base.F64_sub(v160, v161))), base.F64_abs(base.F64_sub(v168, v169)))
									v173 = int32(2)
									v174 = v141 + v173
									v176 = v145 + v173
									if v176 != v125&int32(2147483646) {
										v141 = v174
										v145 = v176
										v152 = v172
										continue
									} else {
										break
									}
									break
								}
								if v125&int32(1) == int32(0) {
									v215 = v172
								} else {
									v180 = v174
									v191 = v172
									v193 = int32(3)
									v195 = v129 + v180<<(uint(v193)%32)
									v199 = *(*float64)(unsafe.Add(mBase, uint32(v195+v125<<(uint(v193)%32))))
									v200 = *(*float64)(unsafe.Add(mBase, uint32(v195)))
									v215 = base.F64_mul(v191, base.F64_abs(base.F64_sub(v199, v200)))
								}
							}
						}
					}
					v218 = base.F32_demote_f64(base.F64_sub(v117, v215))
					*(*float32)(unsafe.Add(mBase, uint32(v14))) = v218
					v221 = F_Float8GetDatum(m, base.F64_promote_f32(v218))
					mBase = m.M
					v222 = m.ExcPending
					if v222 != 0 {
						return int32(0)
					} else {
						return v221
					}
				}
			}
		}
	}
}
