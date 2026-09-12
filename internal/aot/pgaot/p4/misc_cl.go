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
	F_errfinish(m, int32(_a_F_ClosePipeStream_1), int32(3076), int32(_a_F_ClosePipeStream_2))
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	v4 = int32(0)
	if l0 == v4 {
		v163 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v163 & int32(1)
L2:
	;
	if l1 == int32(0) {
		v163 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = l0
	v12 = l1
	v13 = l2
	v15 = v4
	goto L4
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v17 == int32(27) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v163 = v158
	goto L1
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v21 = v20
	goto L8
L7:
	;
	v21 = v12
	goto L8
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v22 != int32(27) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = v11
	goto L11
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v26 = v25
	goto L11
L11:
	;
	v27 = F_equal(m, v26, v21)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v27 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v163 = int32(1)
	goto L1
L15:
	;
	goto L16
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v32 == int32(17) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v36 = F_op_strict(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L20
	}
L18:
	;
	v64 = v32
	goto L19
L19:
	;
	if v64 == int32(15) {
		goto L40
	} else {
		goto L41
	}
L20:
	;
	if v36 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	if v38 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v64 = v63
	goto L19
L24:
	;
	v163 = v15
	goto L1
L25:
	;
	goto L26
L26:
	;
	v41 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v42 <= v41 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v163 = v15
	goto L1
L28:
	;
	goto L29
L29:
	;
	v48 = v41
	goto L30
L30:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v48<<(uint(int32(2))%32))))
	v57 = F_clause_is_strict_for(m, v55, v21, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L32
	}
L31:
	;
	v163 = v57
	goto L1
L32:
	;
	if v57 != 0 {
		v163 = v57
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v60 = v48 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v60 < v61 {
		v48 = v60
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v154 = int32(0)
	if v150 == v154 {
		goto L77
	} else {
		goto L78
	}
L36:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+24)))
	v163 = v149
	goto L1
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v105 = F_clause_is_strict_for(m, v103, v21, int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L12
	} else {
		goto L56
	}
L38:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v150 = v99
	goto L35
L39:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	if v74 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v68 = F_func_strict(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L43
	}
L41:
	;
	v71 = v64
	goto L42
L42:
	;
	switch v71 - int32(7) {
	case 0:
		goto L36
	default:
		v163 = v15
		goto L1
	case 13:
		goto L37
	case 21, 22, 23, 48:
		goto L38
	}
L43:
	;
	if v68 != 0 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v71 = v70
	goto L42
L45:
	;
	v163 = v15
	goto L1
L46:
	;
	goto L47
L47:
	;
	v77 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v78 <= v77 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v163 = v15
	goto L1
L49:
	;
	goto L50
L50:
	;
	v84 = v77
	goto L51
L51:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v84<<(uint(int32(2))%32))))
	v93 = F_clause_is_strict_for(m, v91, v21, int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L53
	}
L52:
	;
	v163 = v93
	goto L1
L53:
	;
	if v93 != 0 {
		v163 = v93
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v96 = v84 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v96 < v97 {
		v84 = v96
		goto L51
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	if v105 == int32(0) {
		v150 = v102
		goto L35
	} else {
		goto L57
	}
L57:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v110 = F_op_strict(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	if v110 == int32(0) {
		v150 = v102
		goto L35
	} else {
		goto L59
	}
L59:
	;
	if v13&int32(1) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v102 == int32(0) {
		v163 = v15
		goto L1
	} else {
		goto L63
	}
L61:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+20)))
	if v118 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v163 = int32(1)
	goto L1
L63:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v124 != int32(35) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v145 <= int32(0) {
		v150 = v102
		goto L35
	} else {
		goto L76
	}
L65:
	;
	if v124 != int32(7) {
		v150 = v102
		goto L35
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+20)))
	if v139 != 0 {
		v150 = v102
		goto L35
	} else {
		goto L74
	}
L68:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+24)))
	if v129 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v163 = int32(1)
	goto L1
L70:
	;
	goto L71
L71:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v102)+20))
	v132 = F_pg_detoast_datum(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L72
	}
L72:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v137 = F_ArrayGetNItems(m, v134, v132+int32(16))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v145 = v137
	goto L64
L74:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	if v140 == int32(0) {
		v150 = v102
		goto L35
	} else {
		goto L75
	}
L75:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v145 = v143
	goto L64
L76:
	;
	v163 = int32(1)
	goto L1
L77:
	;
	v163 = int32(0)
	goto L1
L78:
	;
	goto L79
L79:
	;
	v158 = int32(0)
	if v21 != 0 {
		v11 = v150
		v12 = v21
		v13 = v154
		v15 = v158
		goto L4
	} else {
		goto L80
	}
L80:
	;
	goto L5
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
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
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
			v122 = l0
			m.G0 = v11 + int32(16)
			return v122
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
						v122 = l0
						m.G0 = v11 + int32(16)
						return v122
					} else {
						F_freetree(m, l0)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							v122 = int32(0)
							m.G0 = v11 + int32(16)
							return v122
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
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									v122 = int32(0)
									m.G0 = v11 + int32(16)
									return v122
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
									v122 = v84
									m.G0 = v11 + int32(16)
									return v122
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
									v122 = v99
									m.G0 = v11 + int32(16)
									return v122
								}
							} else {
								if v50 != int32(4) {
									v122 = l0
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
									v122 = l0
								}
								m.G0 = v11 + int32(16)
								return v122
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
				v122 = int32(0)
				m.G0 = v11 + int32(16)
				return v122
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
