package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltree_risparent(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14226(m, l0, int32(_a_F__ltree_risparent_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_ltree_addltree(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = F_ltree_concat(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v13)
							}
						} else {
							return base.I64_extend_i32_u(v13)
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v13)
						}
					} else {
						return base.I64_extend_i32_u(v13)
					}
				}
			}
		}
	}
}
func F_ltree_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v139 int32
	_ = v139
	v3 = int32(0)
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if base.B2i32(v10 == v3)|base.B2i32(v13 == v3) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v139
L2:
	;
	v139 = v10 - v13
	goto L1
L3:
	;
	v17 = int32(8)
	v21 = l0 + v17
	v22 = l1 + v17
	v25 = v13
	v26 = v10
	goto L4
L4:
	;
	v30 = int32(2)
	v31 = v21 + v30
	v33 = v22 + v30
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22))))
	if base.Ui32(v34) < base.Ui32(v35) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L2
L6:
	;
	v37 = v34
	goto L8
L7:
	;
	v37 = v35
	goto L8
L8:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v37) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	if v99 != 0 {
		v139 = v99
		goto L1
	} else {
		goto L27
	}
L10:
	;
	v99 = int32(0)
	goto L9
L11:
	;
	v73 = v68
	v74 = v69
	v75 = v70
	goto L21
L12:
	;
	if (v31|v33)&int32(3) != 0 {
		v68 = v31
		v69 = v33
		v70 = v37
		goto L11
	} else {
		goto L15
	}
L13:
	;
	v61 = v31
	v62 = v33
	v63 = v37
	goto L14
L14:
	;
	if v63 == int32(0) {
		goto L10
	} else {
		goto L20
	}
L15:
	;
	v45 = v31
	v46 = v33
	v47 = v37
	goto L16
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v50 != v51 {
		v68 = v45
		v69 = v46
		v70 = v47
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v61 = v56
	v62 = v54
	v63 = v58
	goto L14
L18:
	;
	v53 = int32(4)
	v54 = v46 + v53
	v56 = v45 + v53
	v58 = v47 - v53
	if base.Ui32(int32(3)) < base.Ui32(v58) {
		v45 = v56
		v46 = v54
		v47 = v58
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v68 = v61
	v69 = v62
	v70 = v63
	goto L11
L21:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v78 == v79 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v99 = v78 - v79
	goto L9
L23:
	;
	v81 = int32(1)
	v86 = v75 - v81
	if v86 != 0 {
		v73 = v73 + v81
		v74 = v74 + v81
		v75 = v86
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	goto L10
L27:
	;
	if v34 != v35 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	return v34 - v35
L29:
	;
	goto L30
L30:
	;
	if v26 < int32(2) {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v105 = int32(1)
	v107 = int32(9)
	v109 = int32(_a_F_ltree_compare_0)
	if v105 < v25 {
		v21 = v21 + (v34+v107)&v109
		v22 = v22 + (v35+v107)&v109
		v25 = v25 - v105
		v26 = v26 - v105
		goto L4
	} else {
		goto L32
	}
L32:
	;
	goto L5
}
func F_ltree_compare_distance(m *base.Module, l0 int32, l1 int32) float32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v118 float64
	_ = v118
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	v3 = int32(0)
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if base.B2i32(v11 == v3)|base.B2i32(v14 == v3) != 0 {
		v150 = v11
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.F32_demote_f64(base.F64_mul(base.F64_mul(base.F64_promote_f32(base.F32_convert_i32_s(v11-v14)), float64(10)), base.F64_convert_i32_u(v150+int32(1))))
L2:
	;
	v18 = int32(8)
	v22 = l0 + v18
	v23 = l1 + v18
	v26 = v11
	v29 = v14
	goto L3
L3:
	;
	v32 = int32(2)
	v33 = v22 + v32
	v35 = v23 + v32
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22))))
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	if base.Ui32(v36) < base.Ui32(v37) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v150 = v130
	goto L1
L5:
	;
	v130 = v26 - int32(1)
	if v26 < int32(2) {
		v150 = v130
		goto L1
	} else {
		goto L34
	}
L6:
	;
	v39 = v36
	goto L8
L7:
	;
	v39 = v37
	goto L8
L8:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v39) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	if v101 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	v101 = int32(0)
	goto L9
L11:
	;
	v75 = v70
	v76 = v71
	v77 = v72
	goto L21
L12:
	;
	if (v33|v35)&int32(3) != 0 {
		v70 = v33
		v71 = v35
		v72 = v39
		goto L11
	} else {
		goto L15
	}
L13:
	;
	v63 = v33
	v64 = v35
	v65 = v39
	goto L14
L14:
	;
	if v65 == int32(0) {
		goto L10
	} else {
		goto L20
	}
L15:
	;
	v47 = v33
	v48 = v35
	v49 = v39
	goto L16
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v52 != v53 {
		v70 = v47
		v71 = v48
		v72 = v49
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v63 = v58
	v64 = v56
	v65 = v60
	goto L14
L18:
	;
	v55 = int32(4)
	v56 = v48 + v55
	v58 = v47 + v55
	v60 = v49 - v55
	if base.Ui32(int32(3)) < base.Ui32(v60) {
		v47 = v58
		v48 = v56
		v49 = v60
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v70 = v63
	v71 = v64
	v72 = v65
	goto L11
L21:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v80 == v81 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v101 = v80 - v81
	goto L9
L23:
	;
	v83 = int32(1)
	v88 = v77 - v83
	if v88 != 0 {
		v75 = v75 + v83
		v76 = v76 + v83
		v77 = v88
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	goto L10
L27:
	;
	if v36 == v37 {
		goto L5
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v118 = base.F64_convert_i32_u(v26 + int32(1))
	if v101 < int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	return base.F32_demote_f64(base.F64_mul(base.F64_mul(base.F64_promote_f32(base.F32_convert_i32_s(v36-v37)), float64(10)), base.F64_convert_i32_u(v26+int32(1))))
L31:
	;
	return base.F32_demote_f64(base.F64_mul(v118, float64(-10)))
L32:
	;
	goto L33
L33:
	;
	return base.F32_demote_f64(base.F64_mul(v118, float64(10)))
L34:
	;
	v133 = int32(9)
	v135 = int32(_a_F_ltree_compare_distance_0)
	v143 = int32(1)
	if v143 < v29 {
		v22 = v22 + (v36+v133)&v135
		v23 = v23 + (v37+v133)&v135
		v26 = v130
		v29 = v29 - v143
		goto L3
	} else {
		goto L35
	}
L35:
	;
	goto L4
}
func F_ltree_decompress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = base.I64_extend_i32_u(v6)
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
		if v10 == v11 {
			return base.I64_extend_i32_u(v4)
		} else {
			v16 = F_palloc(m, int32(24))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v16))) = v10
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v19
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v21
				v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+16)))
				v24 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v16)+18)) = uint8(v24)
				*(*uint16)(unsafe.Add(mBase, uint32(v16)+16)) = uint16(v23)
				return base.I64_extend_i32_u(v16)
			}
		}
	}
}
func F_ltree_gist_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_ltree_gist_in_0), int32(26), int32(_a_F_ltree_gist_in_1), int32(_a_F_ltree_gist_in_2), int32(_a_F_ltree_gist_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_ltree_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	v23 = int32(0)
	if base.B2i32(v22 == v23)|base.B2i32(v21 == v23) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v159 != v14 {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v157 = v22 - v21
	goto L4
L6:
	;
	v28 = int32(8)
	v35 = v14 + v28
	v36 = v19 + v28
	v40 = v21
	v41 = v22
	goto L7
L7:
	;
	v44 = int32(2)
	v45 = v35 + v44
	v47 = v36 + v44
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	if base.Ui32(v48) < base.Ui32(v49) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	v51 = v48
	goto L11
L10:
	;
	v51 = v49
	goto L11
L11:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v51) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v113 != 0 {
		v157 = v113
		goto L4
	} else {
		goto L30
	}
L13:
	;
	v113 = int32(0)
	goto L12
L14:
	;
	v87 = v82
	v88 = v83
	v89 = v84
	goto L24
L15:
	;
	if (v45|v47)&int32(3) != 0 {
		v82 = v45
		v83 = v47
		v84 = v51
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v75 = v45
	v76 = v47
	v77 = v51
	goto L17
L17:
	;
	if v77 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v59 = v45
	v60 = v47
	v61 = v51
	goto L19
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 != v65 {
		v82 = v59
		v83 = v60
		v84 = v61
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v75 = v70
	v76 = v68
	v77 = v72
	goto L17
L21:
	;
	v67 = int32(4)
	v68 = v60 + v67
	v70 = v59 + v67
	v72 = v61 - v67
	if base.Ui32(int32(3)) < base.Ui32(v72) {
		v59 = v70
		v60 = v68
		v61 = v72
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v82 = v75
	v83 = v76
	v84 = v77
	goto L14
L24:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v92 == v93 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v113 = v92 - v93
	goto L12
L26:
	;
	v95 = int32(1)
	v100 = v89 - v95
	if v100 != 0 {
		v87 = v87 + v95
		v88 = v88 + v95
		v89 = v100
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L13
L30:
	;
	if v48 != v49 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v157 = v48 - v49
	goto L4
L32:
	;
	goto L33
L33:
	;
	if v41 < int32(2) {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v118 = int32(1)
	v120 = int32(9)
	v122 = int32(_a_F_ltree_le_0)
	if v118 < v40 {
		v35 = v35 + (v48+v120)&v122
		v36 = v36 + (v49+v120)&v122
		v40 = v40 - v118
		v41 = v41 - v118
		goto L7
	} else {
		goto L35
	}
L35:
	;
	goto L8
L36:
	;
	F_pfree(m, v14)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v163 != v19 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_pfree(m, v19)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	return base.I64_extend_i32_u(base.B2i32(v157 <= int32(0)))
L43:
	;
	goto L42
}
