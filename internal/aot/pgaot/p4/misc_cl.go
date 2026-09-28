package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ClosePipeStream(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ClosePipeStream[0]))
	v8 = v6 - int32(1)
	if int32(0) <= v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ClosePipeStream[1]))
	v14 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	v40 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L12
	}
L4:
	;
	v19 = v12 + v14*int32(12)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v20 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if int32(0) < v14 {
		v14 = v14 - int32(1)
		goto L4
	} else {
		goto L11
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v23 != l0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v25 = F_FreeDesc(m, v19)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	return v25
L11:
	;
	goto L5
L12:
	;
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_errmsg_internal(m, int32(_a_F_ClosePipeStream_0), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v51 = F_pgl_pclose(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	F_errfinish(m, int32(_a_F_ClosePipeStream_1), int32(3060), int32(_a_F_ClosePipeStream_2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	return v51
}
func F_clause_is_strict_for(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	v4 = int32(0)
	if base.B2i32(l0 == v4)|base.B2i32(l1 == v4) != 0 {
		v164 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v164 & int32(1)
L2:
	;
	v12 = l0
	v13 = l1
	v14 = l2
	v16 = v4
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v18 == int32(27) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v164 = v159
	goto L1
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v22 = v21
	goto L7
L6:
	;
	v22 = v13
	goto L7
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v23 != int32(27) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = v12
	goto L10
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v27 = v26
	goto L10
L10:
	;
	v28 = F_equal(m, v27, v22)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	if v28 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v164 = int32(1)
	goto L1
L14:
	;
	goto L15
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v33 == int32(17) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v37 = F_op_strict(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L19
	}
L17:
	;
	v65 = v33
	goto L18
L18:
	;
	if v65 == int32(15) {
		goto L35
	} else {
		goto L36
	}
L19:
	;
	if v37 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v39 == int32(0) {
		v164 = v16
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v65 = v64
	goto L18
L23:
	;
	v42 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v43 <= v42 {
		v164 = v16
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v46 = v42
	goto L25
L25:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v46<<(uint(int32(2))%32))))
	v58 = F_clause_is_strict_for(m, v56, v22, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L27
	}
L26:
	;
	v164 = v58
	goto L1
L27:
	;
	if v58 != 0 {
		v164 = v58
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v61 = v46 + int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v61 < v62 {
		v46 = v61
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v155 = int32(0)
	if v151 == v155 {
		goto L68
	} else {
		goto L69
	}
L31:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+32)))
	v164 = v150
	goto L1
L32:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v106 = F_clause_is_strict_for(m, v104, v22, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L11
	} else {
		goto L47
	}
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v151 = v100
	goto L30
L34:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v75 == int32(0) {
		v164 = v16
		goto L1
	} else {
		goto L40
	}
L35:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v69 = F_func_strict(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L11
	} else {
		goto L38
	}
L36:
	;
	v72 = v65
	goto L37
L37:
	;
	switch v72 - int32(7) {
	case 0:
		goto L31
	default:
		v164 = v16
		goto L1
	case 13:
		goto L32
	case 21, 22, 23, 48:
		goto L33
	}
L38:
	;
	if v69 != 0 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v72 = v71
	goto L37
L40:
	;
	v78 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v79 <= v78 {
		v164 = v16
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v82 = v78
	goto L42
L42:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v82<<(uint(int32(2))%32))))
	v94 = F_clause_is_strict_for(m, v92, v22, int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L11
	} else {
		goto L44
	}
L43:
	;
	v164 = v94
	goto L1
L44:
	;
	if v94 != 0 {
		v164 = v94
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v97 = v82 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v97 < v98 {
		v82 = v97
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	if v106 == int32(0) {
		v151 = v103
		goto L30
	} else {
		goto L48
	}
L48:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v111 = F_op_strict(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L11
	} else {
		goto L49
	}
L49:
	;
	if v111 == int32(0) {
		v151 = v103
		goto L30
	} else {
		goto L50
	}
L50:
	;
	if v14&int32(1) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v103 == int32(0) {
		v164 = v16
		goto L1
	} else {
		goto L54
	}
L52:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+20)))
	if v119 == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v164 = int32(1)
	goto L1
L54:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v125 != int32(35) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v146 <= int32(0) {
		v151 = v103
		goto L30
	} else {
		goto L67
	}
L56:
	;
	if v125 != int32(7) {
		v151 = v103
		goto L30
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)))
	if v140 != 0 {
		v151 = v103
		goto L30
	} else {
		goto L65
	}
L59:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+32)))
	if v130 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v164 = int32(1)
	goto L1
L61:
	;
	goto L62
L62:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v103)+24))
	v133 = F_pg_detoast_datum(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v138 = F_ArrayGetNItemsSafe(m, v135, v133+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v146 = v138
	goto L55
L65:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v103)+16))
	if v141 == int32(0) {
		v151 = v103
		goto L30
	} else {
		goto L66
	}
L66:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v146 = v144
	goto L55
L67:
	;
	v164 = int32(1)
	goto L1
L68:
	;
	v164 = int32(0)
	goto L1
L69:
	;
	goto L70
L70:
	;
	v159 = int32(0)
	if v22 != 0 {
		v12 = v151
		v13 = v22
		v14 = v155
		v16 = v159
		goto L3
	} else {
		goto L71
	}
L71:
	;
	goto L4
}
func F_clean_stopword_intree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	F_check_stack_depth(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v17
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v17
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
		switch v22 - int32(1) {
		case 0:
			v123 = l0
			m.G0 = v11 + int32(16)
			return v123
		default:
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
			if v28 == int32(1) {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v32 = F_clean_stopword_intree(m, v31, l1, l2)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32
					if v32 != 0 {
						v123 = l0
						m.G0 = v11 + int32(16)
						return v123
					} else {
						F_freetree(m, l0)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							v123 = int32(0)
							m.G0 = v11 + int32(16)
							return v123
						}
					}
				}
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v40 = F_clean_stopword_intree(m, v35, v11+int32(12), v11+int32(8))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v40
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v46 = F_clean_stopword_intree(m, v43, v11+int32(4), v11)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
						if v50 == int32(4) {
							v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+2)))
							v54 = v53
						} else {
							v54 = int32(0)
						}
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						if v55 == int32(0) {
							if v46 == int32(0) {
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
								if v50 == int32(4) {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v70 = v63 + (v60 + v54)
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									if v66 < v60 {
										v68 = v60
									} else {
										v68 = v66
									}
									v70 = v68
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v70
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v70
								F_freetree(m, l0)
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									v123 = int32(0)
									m.G0 = v11 + int32(16)
									return v123
								}
							} else {
								if v50 == int32(4) {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
									v80 = v75 + (v76 + v54)
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v80 = v79
								}
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v80
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v82
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								F_pfree(m, l0)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									v123 = v84
									m.G0 = v11 + int32(16)
									return v123
								}
							}
						} else {
							if v46 == int32(0) {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v91
								if v50 == int32(4) {
									v97 = v89 + (v90 + v54)
								} else {
									v97 = v90
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v97
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								F_pfree(m, l0)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									v123 = v99
									m.G0 = v11 + int32(16)
									return v123
								}
							} else {
								if v50 != int32(4) {
									v123 = l0
								} else {
									v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+2)))
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
									v108 = v104 + (v105 + v106)
									*(*uint16)(unsafe.Add(mBase, uint32(v49)+2)) = uint16(v108)
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v110
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v112
									v123 = l0
								}
								m.G0 = v11 + int32(16)
								return v123
							}
						}
					}
				}
			}
		case 2:
			F_pfree(m, l0)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v123 = int32(0)
				m.G0 = v11 + int32(16)
				return v123
			}
		}
	}
}
func F_clogsyncfiletag(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_SlruSyncFileTag(m, int32(_a_F_clogsyncfiletag_0), l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
