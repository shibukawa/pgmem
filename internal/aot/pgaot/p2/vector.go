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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	if l0 <= int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(25133), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(505240), int32(29), int32(102412))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		if l1 <= int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(25133), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(505240), int32(29), int32(102412))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if l2 != 0 {
				v25 = F_palloc(m, int32(20))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(0)
					v34 = (l2 + int32(7)) & int32(-8)
					*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v34
					v36 = F_mul_size(m, l0, v34)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v39 = F_palloc_extended(m, v36, int32(5))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v39
							return v25
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(25133), int32(0))
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(505240), int32(29), int32(102412))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
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
	v37 = *(*float32)(unsafe.Add(mBase, uint32(v33+(l1+v17))))
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
	var v68 int32
	_ = v68
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
	return v68
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
		v68 = v2
		goto L4
	} else {
		goto L15
	}
L11:
	;
	v43 = v31 << (uint(int32(2)) % 32)
	v45 = *(*float32)(unsafe.Add(mBase, uint32(v13+v26+v43)))
	v47 = *(*float32)(unsafe.Add(mBase, uint32(v43+(v18+v26))))
	if base.F32_gt(v45, v47)|base.F32_lt(v45, v47) != 0 {
		v68 = v2
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
	v68 = base.B2i32(v20 <= v21)
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
	var v56 int32
	_ = v56
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
	return v56
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
		v56 = v19
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
	v41 = *(*float32)(unsafe.Add(mBase, uint32(v37+(v15+v23))))
	if base.F32_lt(v39, v41) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v56 = int32(0)
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
		v56 = v19
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
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 float32
	_ = v69
	var v71 float32
	_ = v71
	var v75 int32
	_ = v75
	var v78 float32
	_ = v78
	var v80 float32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v105 float32
	_ = v105
	var v107 float32
	_ = v107
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v126 float32
	_ = v126
	var v133 float32
	_ = v133
	var v137 float32
	_ = v137
	var v141 int32
	_ = v141
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
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
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L37
	}
L5:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L36
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
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L32
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
	v49 = int32(0)
	if v40 != int32(1) {
		goto L15
	} else {
		goto L16
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
	v56 = v49
	v62 = int32(0)
	goto L18
L16:
	;
	v90 = v49
	goto L17
L17:
	;
	if v40&int32(1) != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v65 = int32(2)
	v66 = v56 << (uint(v65) % 32)
	v69 = *(*float32)(unsafe.Add(mBase, uint32(v66+v46)))
	v71 = *(*float32)(unsafe.Add(mBase, uint32(v66+v44)))
	*(*float32)(unsafe.Add(mBase, uint32(v48+v66))) = base.F32_mul(v69, v71)
	v75 = v66 | int32(4)
	v78 = *(*float32)(unsafe.Add(mBase, uint32(v75+v46)))
	v80 = *(*float32)(unsafe.Add(mBase, uint32(v75+v44)))
	*(*float32)(unsafe.Add(mBase, uint32(v48+v75))) = base.F32_mul(v78, v80)
	v84 = v56 + v65
	v86 = v62 + v65
	if v86 != v40&int32(32766) {
		v56 = v84
		v62 = v86
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v90 = v84
	goto L17
L20:
	;
	goto L19
L21:
	;
	v102 = v90 << (uint(int32(2)) % 32)
	v105 = *(*float32)(unsafe.Add(mBase, uint32(v102+v46)))
	v107 = *(*float32)(unsafe.Add(mBase, uint32(v102+v44)))
	*(*float32)(unsafe.Add(mBase, uint32(v48+v102))) = base.F32_mul(v105, v107)
	goto L23
L22:
	;
	goto L23
L23:
	;
	v113 = int32(0)
	goto L24
L24:
	;
	v124 = v113 << (uint(int32(2)) % 32)
	v126 = *(*float32)(unsafe.Add(mBase, uint32(v48+v124)))
	if base.F32_eq(base.F32_abs(v126), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	goto L14
L26:
	;
	if base.F32_ne(v126, float32(0)) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v141 = v113 + int32(1)
	if v141 != v40 {
		v113 = v141
		goto L24
	} else {
		goto L31
	}
L28:
	;
	v133 = *(*float32)(unsafe.Add(mBase, uint32(v124+v44)))
	if base.F32_eq(v133, float32(0)) != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v137 = *(*float32)(unsafe.Add(mBase, uint32(v124+v46)))
	if base.F32_ne(v137, float32(0)) != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	goto L25
L32:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v165 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	v166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v165
	F_errmsg(m, int32(488212), v14)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(506510), int32(76), int32(154322))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L35
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
L37:
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
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
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
	var v90 int32
	_ = v90
	var v101 float32
	_ = v101
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
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
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
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
					v144 = v12
				} else {
					v31 = int32(8)
					v32 = v23 + v31
					v34 = v18 + v31
					if base.Ui32(v25) < base.Ui32(int32(4)) {
						v90 = int32(0)
						v101 = v12
					} else {
						v41 = int32(0)
						v47 = v2
						v52 = v12
						for {
							v54 = v41 << (uint(int32(2)) % 32)
							v56 = v54 | int32(12)
							v58 = *(*float32)(unsafe.Add(mBase, uint32(v34+v56)))
							v60 = *(*float32)(unsafe.Add(mBase, uint32(v32+v56)))
							v63 = v54 | int32(8)
							v65 = *(*float32)(unsafe.Add(mBase, uint32(v34+v63)))
							v67 = *(*float32)(unsafe.Add(mBase, uint32(v32+v63)))
							v69 = int32(4)
							v70 = v54 | v69
							v72 = *(*float32)(unsafe.Add(mBase, uint32(v34+v70)))
							v74 = *(*float32)(unsafe.Add(mBase, uint32(v32+v70)))
							v77 = *(*float32)(unsafe.Add(mBase, uint32(v34+v54)))
							v79 = *(*float32)(unsafe.Add(mBase, uint32(v32+v54)))
							v84 = base.F32_add(base.F32_mul(v58, v60), base.F32_add(base.F32_mul(v65, v67), base.F32_add(base.F32_mul(v72, v74), base.F32_add(base.F32_mul(v77, v79), v52))))
							v86 = v41 + v69
							v88 = v47 + v69
							if v88 != v28&int32(32764) {
								v41 = v86
								v47 = v88
								v52 = v84
								continue
							} else {
								break
							}
							break
						}
						v90 = v86
						v101 = v84
					}
					if v25&int32(3) == int32(0) {
						v144 = v101
					} else {
						v108 = v90
						v117 = v2
						v119 = v101
						for {
							v121 = v108 << (uint(int32(2)) % 32)
							v123 = *(*float32)(unsafe.Add(mBase, uint32(v34+v121)))
							v125 = *(*float32)(unsafe.Add(mBase, uint32(v32+v121)))
							v127 = base.F32_add(base.F32_mul(v123, v125), v119)
							v128 = int32(1)
							v131 = v117 + v128
							if v131 != v28&int32(3) {
								v108 = v108 + v128
								v117 = v131
								v119 = v127
								continue
							} else {
								break
							}
							break
						}
						v144 = v127
					}
				}
				v147 = F_Float8GetDatum(m, base.F64_promote_f32(base.F32_neg(v144)))
				mBase = m.M
				v148 = m.ExcPending
				if v148 != 0 {
					return int32(0)
				} else {
					m.G0 = v15 + int32(16)
					return v147
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v156 = m.ExcPending
				if v156 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v159 = m.ExcPending
					if v159 != 0 {
						return int32(0)
					} else {
						v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
						v161 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v161
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v160
						F_errmsg(m, int32(488212), v15)
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(506510), int32(76), int32(154322))
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
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
