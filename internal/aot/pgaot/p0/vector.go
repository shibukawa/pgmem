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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_DirectFunctionCall1Coll(m, int32(7282), int32(0), l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v14 = F_errstart(m, int32(17), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if v14 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v10
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(208346), v6)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errfinish(m, int32(517306), int32(336), int32(218762))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						F_pfree(m, v10)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					}
				}
			} else {
				F_pfree(m, v10)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					m.G0 = v6 + int32(16)
					return
				}
			}
		}
	}
}
func F_vector_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v43 float32
	_ = v43
	var v45 float32
	_ = v45
	var v53 int32
	_ = v53
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	v21 = base.B2i32(v19 < v20)
	if v19 < v20 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v19 < v20 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v22 = v19
	goto L7
L6:
	;
	v22 = v20
	goto L7
L7:
	;
	if v22 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v25 = int32(8)
	v30 = int32(0)
	goto L9
L9:
	;
	v41 = v30 << (uint(int32(2)) % 32)
	v43 = *(*float32)(unsafe.Add(mBase, uint32(v12+v25+v41)))
	v45 = *(*float32)(unsafe.Add(mBase, uint32(v41+(v17+v25))))
	if base.F32_lt(v43, v45) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	return int32(1)
L11:
	;
	return int32(-1)
L12:
	;
	goto L13
L13:
	;
	if base.F32_gt(v43, v45) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v53 = v30 + int32(1)
	if v53 == v22 {
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
	v30 = v53
	goto L9
L18:
	;
	return int32(-1)
L19:
	;
	goto L20
L20:
	;
	return base.B2i32(v20 < v19)
}
func F_vector_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v43 float32
	_ = v43
	var v45 float32
	_ = v45
	var v51 int32
	_ = v51
	var v70 int32
	_ = v70
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	if v19 < v20 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v70
L5:
	;
	v70 = base.B2i32(v20 <= v19)
	goto L4
L6:
	;
	v22 = v19
	goto L8
L7:
	;
	v22 = v20
	goto L8
L8:
	;
	if v22 <= int32(0) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v25 = int32(8)
	v30 = int32(0)
	goto L10
L10:
	;
	v41 = v30 << (uint(int32(2)) % 32)
	v43 = *(*float32)(unsafe.Add(mBase, uint32(v12+v25+v41)))
	v45 = *(*float32)(unsafe.Add(mBase, uint32(v41+(v17+v25))))
	if base.F32_lt(v43, v45) != 0 {
		v70 = int32(0)
		goto L4
	} else {
		goto L12
	}
L11:
	;
	return int32(1)
L12:
	;
	if base.F32_gt(v43, v45) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v51 = v30 + int32(1)
	if v51 == v22 {
		goto L5
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L11
L16:
	;
	v30 = v51
	goto L10
}
func F_vector_send(m *base.Module, l0 int32) int32 {
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
	return v73
}
func F_vector_to_float4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
		v17 = F_mul_size(m, int32(4), v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_palloc(m, v17)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
				if v21 <= int32(0) {
				} else {
					v25 = v12 + int32(8)
					v26 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v21) {
						v31 = v26
						v38 = v2
						for {
							v41 = v31 << (uint(int32(2)) % 32)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v41+v25)))
							*(*int32)(unsafe.Add(mBase, uint32(v19+v41))) = v44
							v46 = int32(4)
							v47 = v41 | v46
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v25+v47)))
							*(*int32)(unsafe.Add(mBase, uint32(v19+v47))) = v50
							v53 = v41 | int32(8)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v25+v53)))
							*(*int32)(unsafe.Add(mBase, uint32(v19+v53))) = v56
							v59 = v41 | int32(12)
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v59+v25)))
							*(*int32)(unsafe.Add(mBase, uint32(v19+v59))) = v62
							v65 = v31 + v46
							v67 = v38 + v46
							if v67 != v21&int32(32764) {
								v31 = v65
								v38 = v67
								continue
							} else {
								break
							}
							break
						}
						v69 = v65
					} else {
						v69 = v26
					}
					v79 = v21 & int32(3)
					if v79 == int32(0) {
					} else {
						v82 = v69
						v88 = v2
						for {
							v92 = v82 << (uint(int32(2)) % 32)
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v92+v25)))
							*(*int32)(unsafe.Add(mBase, uint32(v19+v92))) = v95
							v97 = int32(1)
							v100 = v88 + v97
							if v100 != v79 {
								v82 = v82 + v97
								v88 = v100
								continue
							} else {
								break
							}
							break
						}
					}
				}
				v115 = F_construct_array(m, v19, v21, int32(700), int32(4), int32(1), int32(105))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v19)
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						return v115
					}
				}
			}
		}
	}
}
