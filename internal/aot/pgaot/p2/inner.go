package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inner_int_contains(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = F_ArrayGetNItems(m, v10, l0+int32(16))
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
	v20 = F_ArrayGetNItems(m, v17, l1+int32(16))
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
	if v13 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return base.B2i32(v20 == int32(0))
L11:
	;
	goto L12
L12:
	;
	if v20 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return base.B2i32(v20 == int32(0))
L14:
	;
	goto L15
L15:
	;
	v56 = int32(0)
	v59 = v56
	v60 = v56
	v63 = int32(0)
	goto L16
L16:
	;
	v67 = int32(2)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0+v32+v63<<(uint(v67)%32))))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1+v43+v59<<(uint(v67)%32))))
	if v74 <= v70 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	return base.B2i32(v88 == v20)
L18:
	;
	goto L17
L19:
	;
	if v70 != v74 {
		v88 = v60
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v81 = v59
	v82 = v60
	goto L21
L21:
	;
	v84 = v63 + int32(1)
	if v13 <= v84 {
		v88 = v82
		goto L18
	} else {
		goto L23
	}
L22:
	;
	v77 = int32(1)
	v81 = v59 + v77
	v82 = v60 + v77
	goto L21
L23:
	;
	if v81 < v20 {
		v59 = v81
		v60 = v82
		v63 = v84
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v88 = v82
	goto L18
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
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 float32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 float32
	_ = v59
	var v61 float32
	_ = v61
	var v64 int32
	_ = v64
	var v66 float32
	_ = v66
	var v68 float32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 float32
	_ = v73
	var v75 float32
	_ = v75
	var v78 float32
	_ = v78
	var v80 float32
	_ = v80
	var v85 float32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v102 float32
	_ = v102
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v118 float32
	_ = v118
	var v120 int32
	_ = v120
	var v122 float32
	_ = v122
	var v124 float32
	_ = v124
	var v126 float32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v143 float32
	_ = v143
	var v157 float64
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
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
					v157 = float64(0)
				} else {
					v32 = int32(8)
					v33 = v23 + v32
					v35 = v18 + v32
					if base.Ui32(v25) < base.Ui32(int32(4)) {
						v91 = int32(0)
						v102 = v12
					} else {
						v42 = int32(0)
						v48 = v2
						v53 = v12
						for {
							v55 = v42 << (uint(int32(2)) % 32)
							v57 = v55 | int32(12)
							v59 = *(*float32)(unsafe.Add(mBase, uint32(v35+v57)))
							v61 = *(*float32)(unsafe.Add(mBase, uint32(v33+v57)))
							v64 = v55 | int32(8)
							v66 = *(*float32)(unsafe.Add(mBase, uint32(v35+v64)))
							v68 = *(*float32)(unsafe.Add(mBase, uint32(v33+v64)))
							v70 = int32(4)
							v71 = v55 | v70
							v73 = *(*float32)(unsafe.Add(mBase, uint32(v35+v71)))
							v75 = *(*float32)(unsafe.Add(mBase, uint32(v33+v71)))
							v78 = *(*float32)(unsafe.Add(mBase, uint32(v35+v55)))
							v80 = *(*float32)(unsafe.Add(mBase, uint32(v33+v55)))
							v85 = base.F32_add(base.F32_mul(v59, v61), base.F32_add(base.F32_mul(v66, v68), base.F32_add(base.F32_mul(v73, v75), base.F32_add(base.F32_mul(v78, v80), v53))))
							v87 = v42 + v70
							v89 = v48 + v70
							if v89 != v28&int32(32764) {
								v42 = v87
								v48 = v89
								v53 = v85
								continue
							} else {
								break
							}
							break
						}
						v91 = v87
						v102 = v85
					}
					if v25&int32(3) != 0 {
						v107 = v91
						v116 = v2
						v118 = v102
						for {
							v120 = v107 << (uint(int32(2)) % 32)
							v122 = *(*float32)(unsafe.Add(mBase, uint32(v35+v120)))
							v124 = *(*float32)(unsafe.Add(mBase, uint32(v33+v120)))
							v126 = base.F32_add(base.F32_mul(v122, v124), v118)
							v127 = int32(1)
							v130 = v116 + v127
							if v130 != v28&int32(3) {
								v107 = v107 + v127
								v116 = v130
								v118 = v126
								continue
							} else {
								break
							}
							break
						}
						v143 = v126
					} else {
						v143 = v102
					}
					v157 = base.F64_promote_f32(v143)
				}
				v158 = F_Float8GetDatum(m, v157)
				mBase = m.M
				v159 = m.ExcPending
				if v159 != 0 {
					return int32(0)
				} else {
					m.G0 = v15 + int32(16)
					return v158
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v170 = m.ExcPending
					if v170 != 0 {
						return int32(0)
					} else {
						v171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
						v172 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v172
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v171
						F_errmsg(m, int32(499600), v15)
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(518027), int32(76), int32(159834))
							mBase = m.M
							v182 = m.ExcPending
							if v182 != 0 {
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
