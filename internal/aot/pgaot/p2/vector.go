package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_VectorArrayFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_pfree(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_VectorArrayInit(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v4 = int32(0)
	if l2 != 0 {
		v11 = base.B2i32(l0 <= v4) | base.B2i32(l1 <= v4)
	} else {
		v11 = int32(1)
	}
	if v11 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_VectorArrayInit_0), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_VectorArrayInit_1), int32(29), int32(_a_F_VectorArrayInit_2))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v28 = F_palloc(m, int32(20))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(0)
			v37 = (l2 + int32(7)) & int32(-8)
			*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v37
			v39 = F_mul_size(m, l0, v37)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v42 = F_palloc_extended(m, v39, int32(5))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v42
					return v28
				}
			}
		}
	}
}
func F_vector_cmp_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v35 float32
	_ = v35
	var v37 float32
	_ = v37
	var v45 int32
	_ = v45
	v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v13 = base.B2i32(v11 < v12)
	if v11 < v12 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v11 < v12 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v14 = v11
	goto L4
L3:
	;
	v14 = v12
	goto L4
L4:
	;
	if v14 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v17 = int32(8)
	v23 = int32(0)
	goto L6
L6:
	;
	v33 = v23 << (uint(int32(2)) % 32)
	v35 = *(*float32)(unsafe.Add(mBase, uint32(l0+v17+v33)))
	v37 = *(*float32)(unsafe.Add(mBase, uint32(l1+v17+v33)))
	if base.F32_lt(v35, v37) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(1)
L8:
	;
	return int32(-1)
L9:
	;
	goto L10
L10:
	;
	if base.F32_gt(v35, v37) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = v23 + int32(1)
	if v45 == v14 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L7
L14:
	;
	v23 = v45
	goto L6
L15:
	;
	return int32(-1)
L16:
	;
	goto L17
L17:
	;
	return base.B2i32(v12 < v11)
}
func F_vector_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v45 float32
	_ = v45
	var v47 float32
	_ = v47
	var v52 int32
	_ = v52
	var v74 int32
	_ = v74
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v22 = base.B2i32(v20 < v21)
	if v20 < v21 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v74
L5:
	;
	v23 = v20
	goto L7
L6:
	;
	v23 = v21
	goto L7
L7:
	;
	if int32(0) < v23 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v26 = int32(8)
	v31 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	if v20 < v21 {
		v74 = v2
		goto L4
	} else {
		goto L15
	}
L11:
	;
	v43 = v31 << (uint(int32(2)) % 32)
	v45 = *(*float32)(unsafe.Add(mBase, uint32(v13+v26+v43)))
	v47 = *(*float32)(unsafe.Add(mBase, uint32(v18+v26+v43)))
	if base.F32_gt(v45, v47)|base.F32_lt(v45, v47) != 0 {
		v74 = v2
		goto L4
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v52 = v31 + int32(1)
	if v52 != v23 {
		v31 = v52
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v74 = base.B2i32(v20 <= v21)
	goto L4
}
func F_vector_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v39 float32
	_ = v39
	var v41 float32
	_ = v41
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+4)))
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+4)))
	v19 = base.B2i32(v17 < v18)
	if v17 < v18 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v54
L5:
	;
	v20 = v17
	goto L7
L6:
	;
	v20 = v18
	goto L7
L7:
	;
	if v20 <= int32(0) {
		v54 = v19
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v23 = int32(8)
	v28 = int32(0)
	goto L9
L9:
	;
	v37 = v28 << (uint(int32(2)) % 32)
	v39 = *(*float32)(unsafe.Add(mBase, uint32(v10+v23+v37)))
	v41 = *(*float32)(unsafe.Add(mBase, uint32(v15+v23+v37)))
	if base.F32_lt(v39, v41) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v54 = int32(0)
	goto L4
L11:
	;
	return int32(1)
L12:
	;
	goto L13
L13:
	;
	if base.F32_gt(v39, v41) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v49 = v28 + int32(1)
	if v49 == v20 {
		v54 = v19
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L10
L17:
	;
	v28 = v49
	goto L9
}
func F_vector_mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 float32
	_ = v68
	var v70 float32
	_ = v70
	var v74 int32
	_ = v74
	var v77 float32
	_ = v77
	var v79 float32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v106 float32
	_ = v106
	var v108 float32
	_ = v108
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v137 float32
	_ = v137
	var v144 float32
	_ = v144
	var v148 float32
	_ = v148
	var v152 int32
	_ = v152
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
	if v24 == v25 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L36
	}
L5:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L35
	}
L6:
	;
	v30 = F_mul_size(m, int32(4), base.I32_extend16_s(v24))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L31
	}
L9:
	;
	v32 = F_add_size(m, int32(8), v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v34 = F_palloc0(m, v32)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v34)+4)) = uint16(v24)
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v32 << (uint(int32(2)) % 32)
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	if int32(0) < v40 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = int32(8)
	v44 = v17 + v43
	v46 = v22 + v43
	v48 = v34 + v43
	if v40 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	m.G0 = v14 + int32(16)
	return v34
L15:
	;
	v124 = int32(0)
	goto L23
L16:
	;
	v54 = v2
	v62 = v2
	goto L19
L17:
	;
	v92 = v2
	goto L18
L18:
	;
	v103 = v92 << (uint(int32(2)) % 32)
	v106 = *(*float32)(unsafe.Add(mBase, uint32(v103+v46)))
	v108 = *(*float32)(unsafe.Add(mBase, uint32(v103+v44)))
	*(*float32)(unsafe.Add(mBase, uint32(v48+v103))) = base.F32_mul(v106, v108)
	goto L15
L19:
	;
	v64 = int32(2)
	v65 = v54 << (uint(v64) % 32)
	v68 = *(*float32)(unsafe.Add(mBase, uint32(v65+v46)))
	v70 = *(*float32)(unsafe.Add(mBase, uint32(v44+v65)))
	*(*float32)(unsafe.Add(mBase, uint32(v48+v65))) = base.F32_mul(v68, v70)
	v74 = v65 | int32(4)
	v77 = *(*float32)(unsafe.Add(mBase, uint32(v74+v46)))
	v79 = *(*float32)(unsafe.Add(mBase, uint32(v44+v74)))
	*(*float32)(unsafe.Add(mBase, uint32(v48+v74))) = base.F32_mul(v77, v79)
	v83 = v54 + v64
	v85 = v62 + v64
	if v85 != v40&int32(_a_F_vector_mul_0) {
		v54 = v83
		v62 = v85
		goto L19
	} else {
		goto L21
	}
L20:
	;
	if v40&int32(1) == int32(0) {
		goto L15
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v92 = v83
	goto L18
L23:
	;
	v135 = v124 << (uint(int32(2)) % 32)
	v137 = *(*float32)(unsafe.Add(mBase, uint32(v48+v135)))
	if base.F32_eq(base.F32_abs(v137), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L5
	} else {
		goto L25
	}
L24:
	;
	goto L14
L25:
	;
	if base.F32_ne(v137, float32(0)) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v152 = v124 + int32(1)
	if v152 != v40 {
		v124 = v152
		goto L23
	} else {
		goto L30
	}
L27:
	;
	v144 = *(*float32)(unsafe.Add(mBase, uint32(v44+v135)))
	if base.F32_eq(v144, float32(0)) != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v148 = *(*float32)(unsafe.Add(mBase, uint32(v135+v46)))
	if base.F32_ne(v148, float32(0)) != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	goto L24
L31:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	v177 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v176
	F_errmsg(m, int32(_a_F_vector_mul_1), v14)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_vector_mul_2), int32(76), int32(_a_F_vector_mul_3))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vector_negative_inner_product(m *base.Module, l0 int32) int32 {
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v51 float32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 float32
	_ = v57
	var v59 float32
	_ = v59
	var v62 int32
	_ = v62
	var v64 float32
	_ = v64
	var v66 float32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 float32
	_ = v71
	var v73 float32
	_ = v73
	var v76 float32
	_ = v76
	var v78 float32
	_ = v78
	var v83 float32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v104 float32
	_ = v104
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
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
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
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
					v143 = v12
				} else {
					v31 = int32(8)
					v32 = v23 + v31
					v34 = v18 + v31
					v35 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v25) {
						v40 = v35
						v49 = v2
						v51 = v12
						for {
							v53 = v40 << (uint(int32(2)) % 32)
							v55 = v53 | int32(12)
							v57 = *(*float32)(unsafe.Add(mBase, uint32(v34+v55)))
							v59 = *(*float32)(unsafe.Add(mBase, uint32(v32+v55)))
							v62 = v53 | int32(8)
							v64 = *(*float32)(unsafe.Add(mBase, uint32(v34+v62)))
							v66 = *(*float32)(unsafe.Add(mBase, uint32(v32+v62)))
							v68 = int32(4)
							v69 = v53 | v68
							v71 = *(*float32)(unsafe.Add(mBase, uint32(v34+v69)))
							v73 = *(*float32)(unsafe.Add(mBase, uint32(v32+v69)))
							v76 = *(*float32)(unsafe.Add(mBase, uint32(v34+v53)))
							v78 = *(*float32)(unsafe.Add(mBase, uint32(v32+v53)))
							v83 = base.F32_add(base.F32_mul(v57, v59), base.F32_add(base.F32_mul(v64, v66), base.F32_add(base.F32_mul(v71, v73), base.F32_add(base.F32_mul(v76, v78), v51))))
							v85 = v40 + v68
							v87 = v49 + v68
							if v87 != v28&int32(_a_F_vector_negative_inner_product_0) {
								v40 = v85
								v49 = v87
								v51 = v83
								continue
							} else {
								break
							}
							break
						}
						if v25&int32(3) == int32(0) {
							v143 = v83
						} else {
							v93 = v85
							v104 = v83
							v107 = v93
							v117 = v2
							v118 = v104
							for {
								v120 = v107 << (uint(int32(2)) % 32)
								v122 = *(*float32)(unsafe.Add(mBase, uint32(v34+v120)))
								v124 = *(*float32)(unsafe.Add(mBase, uint32(v32+v120)))
								v126 = base.F32_add(base.F32_mul(v122, v124), v118)
								v127 = int32(1)
								v130 = v117 + v127
								if v130 != v28&int32(3) {
									v107 = v107 + v127
									v117 = v130
									v118 = v126
									continue
								} else {
									break
								}
								break
							}
							v143 = v126
						}
					} else {
						v93 = v35
						v104 = v12
						v107 = v93
						v117 = v2
						v118 = v104
						for {
							v120 = v107 << (uint(int32(2)) % 32)
							v122 = *(*float32)(unsafe.Add(mBase, uint32(v34+v120)))
							v124 = *(*float32)(unsafe.Add(mBase, uint32(v32+v120)))
							v126 = base.F32_add(base.F32_mul(v122, v124), v118)
							v127 = int32(1)
							v130 = v117 + v127
							if v130 != v28&int32(3) {
								v107 = v107 + v127
								v117 = v130
								v118 = v126
								continue
							} else {
								break
							}
							break
						}
						v143 = v126
					}
				}
				v146 = F_Float8GetDatum(m, base.F64_promote_f32(base.F32_neg(v143)))
				mBase = m.M
				v147 = m.ExcPending
				if v147 != 0 {
					return int32(0)
				} else {
					m.G0 = v15 + int32(16)
					return v146
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v155 = m.ExcPending
				if v155 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v158 = m.ExcPending
					if v158 != 0 {
						return int32(0)
					} else {
						v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
						v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v160
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v159
						F_errmsg(m, int32(_a_F_vector_negative_inner_product_1), v15)
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_vector_negative_inner_product_2), int32(76), int32(_a_F_vector_negative_inner_product_3))
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
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
