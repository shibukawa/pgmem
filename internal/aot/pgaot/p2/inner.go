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
func F_inner_product(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 float32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v53 float32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 float32
	_ = v60
	var v62 float32
	_ = v62
	var v65 int32
	_ = v65
	var v67 float32
	_ = v67
	var v69 float32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 float32
	_ = v74
	var v76 float32
	_ = v76
	var v79 float32
	_ = v79
	var v81 float32
	_ = v81
	var v86 float32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v107 float32
	_ = v107
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v122 float32
	_ = v122
	var v125 int32
	_ = v125
	var v127 float32
	_ = v127
	var v129 float32
	_ = v129
	var v131 float32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v148 float32
	_ = v148
	var v165 int64
	_ = v165
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	v2 = int32(0)
	v12 = float32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int64(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v24 = F_pg_detoast_datum(m, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
			if v26 == v27 {
				v29 = base.I32_extend16_s(v26)
				if v29 <= int32(0) {
					v165 = int64(0)
				} else {
					v33 = int32(8)
					v34 = v24 + v33
					v36 = v19 + v33
					v37 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v26) {
						v42 = v37
						v51 = v2
						v53 = v12
						for {
							v56 = v42 << (uint(int32(2)) % 32)
							v58 = v56 | int32(12)
							v60 = *(*float32)(unsafe.Add(mBase, uint32(v36+v58)))
							v62 = *(*float32)(unsafe.Add(mBase, uint32(v34+v58)))
							v65 = v56 | int32(8)
							v67 = *(*float32)(unsafe.Add(mBase, uint32(v36+v65)))
							v69 = *(*float32)(unsafe.Add(mBase, uint32(v34+v65)))
							v71 = int32(4)
							v72 = v56 | v71
							v74 = *(*float32)(unsafe.Add(mBase, uint32(v36+v72)))
							v76 = *(*float32)(unsafe.Add(mBase, uint32(v34+v72)))
							v79 = *(*float32)(unsafe.Add(mBase, uint32(v36+v56)))
							v81 = *(*float32)(unsafe.Add(mBase, uint32(v34+v56)))
							v86 = base.F32_add(base.F32_mul(v60, v62), base.F32_add(base.F32_mul(v67, v69), base.F32_add(base.F32_mul(v74, v76), base.F32_add(base.F32_mul(v79, v81), v53))))
							v88 = v42 + v71
							v90 = v51 + v71
							if v90 != v29&int32(_a_F_inner_product_0) {
								v42 = v88
								v51 = v90
								v53 = v86
								continue
							} else {
								break
							}
							break
						}
						if v26&int32(3) == int32(0) {
							v148 = v86
						} else {
							v96 = v88
							v107 = v86
							v111 = v96
							v121 = v2
							v122 = v107
							for {
								v125 = v111 << (uint(int32(2)) % 32)
								v127 = *(*float32)(unsafe.Add(mBase, uint32(v36+v125)))
								v129 = *(*float32)(unsafe.Add(mBase, uint32(v34+v125)))
								v131 = base.F32_add(base.F32_mul(v127, v129), v122)
								v132 = int32(1)
								v135 = v121 + v132
								if v135 != v29&int32(3) {
									v111 = v111 + v132
									v121 = v135
									v122 = v131
									continue
								} else {
									break
								}
								break
							}
							v148 = v131
						}
					} else {
						v96 = v37
						v107 = v12
						v111 = v96
						v121 = v2
						v122 = v107
						for {
							v125 = v111 << (uint(int32(2)) % 32)
							v127 = *(*float32)(unsafe.Add(mBase, uint32(v36+v125)))
							v129 = *(*float32)(unsafe.Add(mBase, uint32(v34+v125)))
							v131 = base.F32_add(base.F32_mul(v127, v129), v122)
							v132 = int32(1)
							v135 = v121 + v132
							if v135 != v29&int32(3) {
								v111 = v111 + v132
								v121 = v135
								v122 = v131
								continue
							} else {
								break
							}
							break
						}
						v148 = v131
					}
					v165 = base.I64_reinterpret_f64(base.F64_promote_f32(v148))
				}
				m.G0 = v16 + int32(16)
				return v165
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v176 = m.ExcPending
					if v176 != 0 {
						return int64(0)
					} else {
						v177 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+4)))
						v178 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v178
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v177
						F_errmsg(m, int32(_a_F_inner_product_1), v16)
						mBase = m.M
						v183 = m.ExcPending
						if v183 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_inner_product_2), int32(76), int32(_a_F_inner_product_3))
							mBase = m.M
							v188 = m.ExcPending
							if v188 != 0 {
								return int64(0)
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
