package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PrintVector(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v11 = F_DirectFunctionCall1Coll(m, int32(_a_F_PrintVector_0), int32(0), base.I64_extend_i32_u(l1))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = base.I32_wrap_i64(v11)
		v16 = F_errstart(m, int32(17), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v13
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_PrintVector_1), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PrintVector_2), int32(336), int32(_a_F_PrintVector_3))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_pfree(m, v13)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					}
				}
			} else {
				F_pfree(m, v13)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	}
}
func F_vector_cmp(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14402(m, l0, int64(-1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_vector_ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v41 float32
	_ = v41
	var v43 float32
	_ = v43
	var v51 int32
	_ = v51
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+4)))
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
	if v18 < v19 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return base.I64_extend_i32_u(base.B2i32(v19 <= v18))
L5:
	;
	v21 = v18
	goto L7
L6:
	;
	v21 = v19
	goto L7
L7:
	;
	if v21 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v24 = int32(8)
	v29 = int32(0)
	goto L9
L9:
	;
	v39 = v29 << (uint(int32(2)) % 32)
	v41 = *(*float32)(unsafe.Add(mBase, uint32(v11+v24+v39)))
	v43 = *(*float32)(unsafe.Add(mBase, uint32(v16+v24+v39)))
	if base.F32_lt(v41, v43) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	return int64(1)
L11:
	;
	return int64(0)
L12:
	;
	goto L13
L13:
	;
	if base.F32_gt(v41, v43) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v51 = v29 + int32(1)
	if v51 == v21 {
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
	v29 = v51
	goto L9
}
func F_vector_send(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v61 float32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	F_pq_begintypsend(m, v7)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
	F_enlargeStringInfo(m, v7, int32(2))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v23 = int32(8)
	v27 = v16<<(uint(v23)%32) | int32(base.Ui32(v16)>>(uint(v23)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v20+v21))) = uint16(v27)
	v29 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v20 + v29
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+6)))
	F_enlargeStringInfo(m, v7, v29)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v39 = int32(8)
	v43 = v32<<(uint(v39)%32) | int32(base.Ui32(v32)>>(uint(v39)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v36+v37))) = uint16(v43)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v36 + int32(2)
	v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+4)))
	if int32(0) < v48 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v54 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v74 << (uint(int32(2)) % 32)
	goto L13
L9:
	;
	v61 = *(*float32)(unsafe.Add(mBase, uint32(v10+int32(8)+v54<<(uint(int32(2))%32))))
	F_pq_sendfloat4(m, v7, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v65 = v54 + int32(1)
	v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+4)))
	if v65 < v66 {
		v54 = v65
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	m.G0 = v7 + int32(16)
	return base.I64_extend_i32_u(v73)
}
func F_vector_to_float4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v11 int32
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+4)))
		v16 = F_palloc_mul(m, int32(8), v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+4)))
			if v18 <= int32(0) {
			} else {
				v22 = v11 + int32(8)
				v23 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v18) {
					v28 = v23
					v34 = v2
					for {
						v36 = int32(3)
						v39 = int32(2)
						v42 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22+v28<<(uint(v39)%32)))))
						*(*int64)(unsafe.Add(mBase, uint32(v16+v28<<(uint(v36)%32)))) = v42
						v45 = v28 | int32(1)
						v52 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22+v45<<(uint(v39)%32)))))
						*(*int64)(unsafe.Add(mBase, uint32(v16+v45<<(uint(v36)%32)))) = v52
						v55 = v28 | v39
						v62 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22+v55<<(uint(v39)%32)))))
						*(*int64)(unsafe.Add(mBase, uint32(v16+v55<<(uint(v36)%32)))) = v62
						v65 = v28 | v36
						v72 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22+v65<<(uint(v39)%32)))))
						*(*int64)(unsafe.Add(mBase, uint32(v16+v65<<(uint(v36)%32)))) = v72
						v74 = int32(4)
						v75 = v28 + v74
						v77 = v34 + v74
						if v77 != v18&int32(_a_F_vector_to_float4_0) {
							v28 = v75
							v34 = v77
							continue
						} else {
							break
						}
						break
					}
					if v18&int32(3) == int32(0) {
					} else {
						v83 = v75
						v93 = v83
						v100 = v2
						for {
							v107 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22+v93<<(uint(int32(2))%32)))))
							*(*int64)(unsafe.Add(mBase, uint32(v16+v93<<(uint(int32(3))%32)))) = v107
							v109 = int32(1)
							v112 = v100 + v109
							if v112 != v18&int32(3) {
								v93 = v93 + v109
								v100 = v112
								continue
							} else {
								break
							}
							break
						}
					}
				} else {
					v83 = v23
					v93 = v83
					v100 = v2
					for {
						v107 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22+v93<<(uint(int32(2))%32)))))
						*(*int64)(unsafe.Add(mBase, uint32(v16+v93<<(uint(int32(3))%32)))) = v107
						v109 = int32(1)
						v112 = v100 + v109
						if v112 != v18&int32(3) {
							v93 = v93 + v109
							v100 = v112
							continue
						} else {
							break
						}
						break
					}
				}
			}
			v126 = F_construct_array(m, v16, v18, int32(700), int32(4), int32(1), int32(105))
			mBase = m.M
			v127 = m.ExcPending
			if v127 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v16)
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v126)
				}
			}
		}
	}
}
