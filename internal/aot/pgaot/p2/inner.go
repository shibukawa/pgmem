package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inner_int_contains(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	v3 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = F_ArrayGetNItemsSafe(m, v10, l0+int32(16))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = F_ArrayGetNItemsSafe(m, v17, l1+int32(16))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = (v25<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L6
L5:
	;
	v32 = v22
	goto L6
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v33 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v43 = (v36<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L9
L8:
	;
	v43 = v33
	goto L9
L9:
	;
	v44 = int32(0)
	if base.B2i32(v13 <= v44)|base.B2i32(v20 <= v44) != 0 {
		v88 = v3
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return base.B2i32(v20 == v88)
L11:
	;
	v51 = int32(0)
	v53 = v51
	v54 = v51
	v59 = v3
	goto L12
L12:
	;
	v62 = int32(2)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0+v32+v54<<(uint(v62)%32))))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1+v43+v53<<(uint(v62)%32))))
	if v69 <= v65 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v88 = v77
	goto L10
L14:
	;
	if v65 != v69 {
		v88 = v59
		goto L10
	} else {
		goto L17
	}
L15:
	;
	v76 = v53
	v77 = v59
	goto L16
L16:
	;
	v79 = v54 + int32(1)
	if v13 <= v79 {
		v88 = v77
		goto L10
	} else {
		goto L18
	}
L17:
	;
	v72 = int32(1)
	v76 = v53 + v72
	v77 = v59 + v72
	goto L16
L18:
	;
	if v76 < v20 {
		v53 = v76
		v54 = v79
		v59 = v77
		goto L12
	} else {
		goto L19
	}
L19:
	;
	goto L13
}
func F_inner_product(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 float32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v52 float32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 float32
	_ = v58
	var v60 float32
	_ = v60
	var v63 int32
	_ = v63
	var v65 float32
	_ = v65
	var v67 float32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 float32
	_ = v72
	var v74 float32
	_ = v74
	var v77 float32
	_ = v77
	var v79 float32
	_ = v79
	var v84 float32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v105 float32
	_ = v105
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v119 float32
	_ = v119
	var v121 int32
	_ = v121
	var v123 float32
	_ = v123
	var v125 float32
	_ = v125
	var v127 float32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v144 float32
	_ = v144
	var v158 float64
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	v2 = int32(0)
	v12 = float32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v23 = F_pg_detoast_datum(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
			if v25 == v26 {
				v28 = base.I32_extend16_s(v25)
				if v28 <= int32(0) {
					v158 = float64(0)
				} else {
					v32 = int32(8)
					v33 = v23 + v32
					v35 = v18 + v32
					v36 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v25) {
						v41 = v36
						v50 = v2
						v52 = v12
						for {
							v54 = v41 << (uint(int32(2)) % 32)
							v56 = v54 | int32(12)
							v58 = *(*float32)(unsafe.Add(mBase, uint32(v35+v56)))
							v60 = *(*float32)(unsafe.Add(mBase, uint32(v33+v56)))
							v63 = v54 | int32(8)
							v65 = *(*float32)(unsafe.Add(mBase, uint32(v35+v63)))
							v67 = *(*float32)(unsafe.Add(mBase, uint32(v33+v63)))
							v69 = int32(4)
							v70 = v54 | v69
							v72 = *(*float32)(unsafe.Add(mBase, uint32(v35+v70)))
							v74 = *(*float32)(unsafe.Add(mBase, uint32(v33+v70)))
							v77 = *(*float32)(unsafe.Add(mBase, uint32(v35+v54)))
							v79 = *(*float32)(unsafe.Add(mBase, uint32(v33+v54)))
							v84 = base.F32_add(base.F32_mul(v58, v60), base.F32_add(base.F32_mul(v65, v67), base.F32_add(base.F32_mul(v72, v74), base.F32_add(base.F32_mul(v77, v79), v52))))
							v86 = v41 + v69
							v88 = v50 + v69
							if v88 != v28&int32(_a_F_inner_product_0) {
								v41 = v86
								v50 = v88
								v52 = v84
								continue
							} else {
								break
							}
							break
						}
						if v25&int32(3) == int32(0) {
							v144 = v84
						} else {
							v94 = v86
							v105 = v84
							v108 = v94
							v118 = v2
							v119 = v105
							for {
								v121 = v108 << (uint(int32(2)) % 32)
								v123 = *(*float32)(unsafe.Add(mBase, uint32(v35+v121)))
								v125 = *(*float32)(unsafe.Add(mBase, uint32(v33+v121)))
								v127 = base.F32_add(base.F32_mul(v123, v125), v119)
								v128 = int32(1)
								v131 = v118 + v128
								if v131 != v28&int32(3) {
									v108 = v108 + v128
									v118 = v131
									v119 = v127
									continue
								} else {
									break
								}
								break
							}
							v144 = v127
						}
					} else {
						v94 = v36
						v105 = v12
						v108 = v94
						v118 = v2
						v119 = v105
						for {
							v121 = v108 << (uint(int32(2)) % 32)
							v123 = *(*float32)(unsafe.Add(mBase, uint32(v35+v121)))
							v125 = *(*float32)(unsafe.Add(mBase, uint32(v33+v121)))
							v127 = base.F32_add(base.F32_mul(v123, v125), v119)
							v128 = int32(1)
							v131 = v118 + v128
							if v131 != v28&int32(3) {
								v108 = v108 + v128
								v118 = v131
								v119 = v127
								continue
							} else {
								break
							}
							break
						}
						v144 = v127
					}
					v158 = base.F64_promote_f32(v144)
				}
				v159 = F_Float8GetDatum(m, v158)
				mBase = m.M
				v160 = m.ExcPending
				if v160 != 0 {
					return int32(0)
				} else {
					m.G0 = v15 + int32(16)
					return v159
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v168 = m.ExcPending
				if v168 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						v172 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
						v173 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v173
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v172
						F_errmsg(m, int32(_a_F_inner_product_1), v15)
						mBase = m.M
						v178 = m.ExcPending
						if v178 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_inner_product_2), int32(76), int32(_a_F_inner_product_3))
							mBase = m.M
							v183 = m.ExcPending
							if v183 != 0 {
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
