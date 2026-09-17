package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QT2QTN(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	v6 = F_palloc0(m, int32(24))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		F_check_stack_depth(m)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v13 == int32(2) {
				v17 = F_palloc0(m, int32(8))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v17
					v22 = F_QT2QTN(m, l0+int32(12), l1)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v24))) = v22
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v28
						v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
						if v30 == int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(1)
							return v6
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(2)
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v41 = F_QT2QTN(m, l0+v37*int32(12), l1)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v41
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v45 | v48
								return v6
							}
						}
					}
				}
			} else {
				if l1 == int32(0) {
					return v6
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1 + int32(base.Ui32(v54)>>(uint(int32(12))%32))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(1) << (uint(v60) % 32)
					return v6
				}
			}
		}
	}
}
func F_qtext_store(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_store[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(1)
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_store[0]))
	F_s_lock(m, v19+int32(20), int32(_a_F_qtext_store_0), int32(2237), int32(_a_F_qtext_store_1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_store[0]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	v32 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+28)) = v31 + v32
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v36 = l1 + v35
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v36 + v32
	if l3 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v40
	goto L8
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v35
	if base.Ui32(v35^int32(2147483647)) <= base.Ui32(l1) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_store[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+20)) = int32(0)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+28)) = v128 - int32(1)
	m.G0 = v11 + int32(16)
	return v122
L10:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_store[0]))
	F_s_lock(m, v114+int32(20), int32(_a_F_qtext_store_0), v112, int32(_a_F_qtext_store_1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L33
	}
L11:
	;
	v82 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L21
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_qtext_store[1])) = int32(22)
	v79 = int32(-1)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v54 = F_OpenTransientFile(m, int32(_a_F_qtext_store_2), int32(66))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if v54 < int32(0) {
		v79 = v54
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v59 = F_pwrite(m, v54, l0, l1, base.I64_extend_i32_u(v35))
	mBase = m.M
	if v59 != l1 {
		v79 = v54
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v61 = int32(1)
	v65 = F_pwrite(m, v54, int32(_a_F_qtext_store_4), v61, base.I64_extend_i32_u(v36))
	mBase = m.M
	if v65 != v61 {
		v79 = v54
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v68 = F_CloseTransientFile(m, v54)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_store[0]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
	if v72 == int32(0) {
		v122 = v61
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v110 = v61
	v112 = int32(2272)
	goto L10
L21:
	;
	if v82 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if int32(0) <= v79 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_qtext_store_2)
	F_errmsg(m, int32(_a_F_qtext_store_3), v11)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_qtext_store_0), int32(2282), int32(_a_F_qtext_store_1))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v98 = F_CloseTransientFile(m, v79)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_store[0]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = int32(1)
	v105 = int32(0)
	if v102 == v105 {
		v122 = v105
		goto L9
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v110 = v105
	v112 = int32(2288)
	goto L10
L33:
	;
	v122 = v110
	goto L9
}
func F_quote_literal(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v50 = F_palloc(m, v45<<(uint(int32(1))%32)+int32(7))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L14
	}
L2:
	;
	return int32(0)
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v16 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v22 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v33 = int32(1)
	if v16&v33 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v25 = int32(16)
	goto L9
L8:
	;
	v25 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v32 = int32(4)
	goto L12
L11:
	;
	v32 = v25
	goto L12
L12:
	;
	v45 = v32
	goto L1
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L14:
	;
	v53 = v50 + int32(4)
	if v45 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v226 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v218))) = uint8(v226)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = (v217-v53)<<(uint(int32(2))%32) + int32(24)
	return v50
L16:
	;
	v54 = int32(1)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v56&v54 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v212 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+4)) = uint8(v212)
	v217 = v53
	v218 = v50 + int32(5)
	goto L15
L19:
	;
	v59 = v54
	goto L21
L20:
	;
	v59 = int32(4)
	goto L21
L21:
	;
	v60 = v12 + v59
	v63 = v60
	goto L23
L22:
	;
	v84 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v84)
	v87 = v83 + int32(1)
	v89 = v45 & int32(3)
	if v89 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v72 == int32(92) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v83 = v53
	goto L22
L25:
	;
	v75 = int32(69)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+4)) = uint8(v75)
	v83 = v50 + int32(5)
	goto L22
L26:
	;
	goto L27
L27:
	;
	v80 = v63 + int32(1)
	if base.Ui32(v80) < base.Ui32(v60+v45) {
		v63 = v80
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	if base.Ui32(v45) < base.Ui32(int32(4)) {
		v217 = v125
		v218 = v126
		goto L15
	} else {
		goto L39
	}
L30:
	;
	v124 = v60
	v125 = v83
	v126 = v87
	v129 = v45
	goto L29
L31:
	;
	goto L32
L32:
	;
	v92 = v60
	v93 = v83
	v94 = v87
	v97 = v45
	v101 = int32(0)
	goto L33
L33:
	;
	v103 = v97 - int32(1)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if base.B2i32(v104 == int32(92))|base.B2i32(v104 == int32(39)) != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v124 = v120
	v125 = v115
	v126 = v118
	v129 = v103
	goto L29
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v104)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v114 = v111
	v115 = v93 + int32(2)
	goto L37
L36:
	;
	v114 = v104
	v115 = v94
	goto L37
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v114)
	v117 = int32(1)
	v118 = v115 + v117
	v120 = v92 + v117
	v122 = v101 + v117
	if v122 != v89 {
		v92 = v120
		v93 = v115
		v94 = v118
		v97 = v103
		v101 = v122
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v136 = v124
	v137 = v125
	v138 = v126
	v141 = v129
	goto L40
L40:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if base.B2i32(v146 != int32(92))&base.B2i32(v146 != int32(39)) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v217 = v204
	v218 = v207
	goto L15
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v138))) = uint8(v146)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v158 = v137 + int32(2)
	v159 = v155
	goto L44
L43:
	;
	v158 = v138
	v159 = v146
	goto L44
L44:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v159)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	if base.B2i32(v161 == int32(92))|base.B2i32(v161 == int32(39)) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v158)+1)) = uint8(v161)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	v173 = v168
	v174 = v158 + int32(2)
	goto L47
L46:
	;
	v173 = v161
	v174 = v158 + int32(1)
	goto L47
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v173)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+2)))
	if base.B2i32(v176 == int32(92))|base.B2i32(v176 == int32(39)) != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)) = uint8(v176)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+2)))
	v188 = v183
	v189 = v174 + int32(2)
	goto L50
L49:
	;
	v188 = v176
	v189 = v174 + int32(1)
	goto L50
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v189))) = uint8(v188)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+3)))
	if base.B2i32(v191 == int32(92))|base.B2i32(v191 == int32(39)) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)) = uint8(v191)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+3)))
	v203 = v198
	v204 = v189 + int32(2)
	goto L53
L52:
	;
	v203 = v191
	v204 = v189 + int32(1)
	goto L53
L53:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v204))) = uint8(v203)
	v207 = v204 + int32(1)
	v208 = int32(4)
	v211 = v141 - v208
	if v211 != 0 {
		v136 = v136 + v208
		v137 = v204
		v138 = v207
		v141 = v211
		goto L40
	} else {
		goto L54
	}
L54:
	;
	goto L41
}
