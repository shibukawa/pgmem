package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vac_bulkdel_one_index(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_index_bulk_delete(m, l0, l1, int32(583), l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v17 = F_errstart(m, v15, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
				v21 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v21
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v20 + int32(4)
				F_errmsg(m, int32(147541), v8)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(498520), int32(2661), int32(28362))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v11
					}
				}
			} else {
				m.G0 = v8 + int32(16)
				return v11
			}
		}
	}
}
func F_vac_estimate_reltuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) float64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 float32
	_ = v10
	var v11 float64
	_ = v11
	var v12 int32
	_ = v12
	var v28 float64
	_ = v28
	var v32 int32
	_ = v32
	var v53 float64
	_ = v53
	if base.Ui32(l2) < base.Ui32(l1) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v10 = *(*float32)(unsafe.Add(mBase, uint32(v9)+100))
		v11 = base.F64_promote_f32(v10)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+96))
		if l1 == v12 {
			if base.Ui32(l2) < base.Ui32(int32(2)) {
				return v11
			} else {
				if base.F64_lt(base.F64_convert_i32_u(l2), base.F64_mul(base.F64_convert_i32_u(l1), float64(0.02))) == int32(0) {
					v28 = base.F64_convert_i32_u(l1)
					if v12 != 0 {
						v32 = base.F32_lt(v10, float32(0))
					} else {
						v32 = int32(1)
					}
					if v32 != 0 {
						return base.F64_floor(base.F64_add(base.F64_mul(base.F64_div(l3, base.F64_convert_i32_u(l2)), v28), float64(0.5)))
					} else {
						v53 = base.F64_floor(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(v11, base.F64_convert_i32_u(v12)), base.F64_sub(v28, base.F64_convert_i32_u(l2))), l3), float64(0.5)))
						return v53
					}
				} else {
					return v11
				}
			}
		} else {
			if base.Ui32(int32(2)) <= base.Ui32(l2) {
				v28 = base.F64_convert_i32_u(l1)
				if v12 != 0 {
					v32 = base.F32_lt(v10, float32(0))
				} else {
					v32 = int32(1)
				}
				if v32 != 0 {
					return base.F64_floor(base.F64_add(base.F64_mul(base.F64_div(l3, base.F64_convert_i32_u(l2)), v28), float64(0.5)))
				} else {
					v53 = base.F64_floor(base.F64_add(base.F64_add(base.F64_mul(base.F64_div(v11, base.F64_convert_i32_u(v12)), base.F64_sub(v28, base.F64_convert_i32_u(l2))), l3), float64(0.5)))
					return v53
				}
			} else {
				return v11
			}
		}
	} else {
		v53 = l3
		return v53
	}
}
func F_vac_open_indexes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v5 = int32(0)
	v8 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v63
	F_list_free(m, v8)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L22
	}
L2:
	;
	v27 = v10
	v31 = int32(0)
	goto L13
L3:
	;
	return
L4:
	;
	if v8 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v10 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v10 < v11 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	v63 = v5
	goto L1
L8:
	;
	v16 = F_palloc(m, v11<<(uint(int32(2))%32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	v18 = v5
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if int32(0) < v20 {
		goto L2
	} else {
		goto L12
	}
L11:
	;
	v18 = v16
	goto L10
L12:
	;
	v63 = int32(0)
	goto L1
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v27<<(uint(int32(2))%32))))
	v39 = F_index_open(m, v38, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L16
	}
L14:
	;
	v63 = v54
	goto L1
L15:
	;
	v56 = v27 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v56 < v57 {
		v27 = v56
		v31 = v54
		goto L13
	} else {
		goto L21
	}
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+192))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+20)))
	if v42 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v45+v31<<(uint(int32(2))%32)))) = v39
	v54 = v31 + int32(1)
	goto L15
L18:
	;
	goto L19
L19:
	;
	F_relation_close(m, v39, l1)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v54 = v31
	goto L15
L21:
	;
	goto L14
L22:
	;
	return
}
