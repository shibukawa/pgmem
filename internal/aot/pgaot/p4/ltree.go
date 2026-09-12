package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltree_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 == v2 {
		v30 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v30&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if v17 == int32(0) {
		v30 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v20 != int32(7) {
		v30 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v23 != int32(17) {
		v30 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+24)))
	v30 = v26 ^ int32(1)
	goto L2
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_get_fn_opclass_options(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v39 = int32(28)
	goto L9
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)))
	if v41 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int32(0)
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v39 = v38
	goto L9
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L55
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L51
	}
L14:
	;
	v201 = F_palloc(m, int32(16))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L10
	} else {
		goto L50
	}
L15:
	;
	v44 = F_pg_detoast_datum(m, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+4)))
	if v148&int32(2) != 0 {
		goto L38
	} else {
		goto L39
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v49 = F_ArrayGetNItems(m, v46, v44+int32(16))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if int32(2) <= v51 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v55 = F_array_contains_nulls(m, v44)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	if v55 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v57 = int32(0)
	v61 = F_ltree_gist_alloc(m, v57, v57, v39, v57, v57)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	if v49 <= int32(0) {
		v197 = v61
		goto L14
	} else {
		goto L24
	}
L24:
	;
	if v54 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v71 = v54
	goto L27
L26:
	;
	v71 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L27
L27:
	;
	v81 = v44 + v71
	v82 = v49
	goto L28
L28:
	;
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+4)))
	if v87 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v197 = v61
	goto L14
L30:
	;
	v90 = v81 + int32(8)
	v91 = v87
	goto L33
L31:
	;
	goto L32
L32:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v144 = int32(1)
	if v144 < v82 {
		v81 = v81 + (int32(base.Ui32(v136)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v82 = v82 - v144
		goto L28
	} else {
		goto L37
	}
L33:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90))))
	v103 = F_ltree_crc32_sz(m, v90+int32(2), v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	v105 = base.I32_rem_u_s(v103, v39<<(uint(int32(3))%32))
	v108 = v61 + int32(8) + int32(base.Ui32(v105)>>(uint(int32(3))%32))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v110 = int32(1)
	v114 = v109 | v110<<(uint(v105&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v114)
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90))))
	if base.Ui32(v110) < base.Ui32(v91) {
		v90 = v90 + (v116+int32(9))&int32(131064)
		v91 = v91 - v110
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	goto L29
L38:
	;
	return v11
L39:
	;
	goto L40
L40:
	;
	v153 = v40 + int32(8)
	v154 = int32(0)
	if v39 <= v154 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v186 = int32(0)
	v188 = F_ltree_gist_alloc(m, int32(1), v153, v39, v186, v186)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L10
	} else {
		goto L49
	}
L42:
	;
	v157 = v154
	goto L43
L43:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+v153))))
	if v168 == int32(255) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	return v11
L45:
	;
	v172 = v157 + int32(1)
	if v39 != v172 {
		v157 = v172
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L41
L49:
	;
	v197 = v188
	goto L14
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v197
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v201)+8)) = v206
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)))
	v209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+14)) = uint8(v209)
	*(*uint16)(unsafe.Add(mBase, uint32(v201)+12)) = uint16(v208)
	return v201
L51:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(325635), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(511887), int32(67), int32(134565))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(160856), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(511887), int32(71), int32(134565))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__ltree_isparent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = F_array_iterator(m, v6, int32(5638), v12, int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v17 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v21 != v12 {
							F_pfree(m, v12)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int32(0)
							} else {
								return v15
							}
						} else {
							return v15
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v21 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							return v15
						}
					} else {
						return v15
					}
				}
			}
		}
	}
}
func F__ltree_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 == v2 {
		v27 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v27&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v14 == int32(0) {
		v27 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v17 != int32(7) {
		v27 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v20 != int32(17) {
		v27 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
	v27 = v23 ^ int32(1)
	goto L2
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = F_get_fn_opclass_options(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v36 = int32(28)
	goto L9
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v38 = int32(2)
	v39 = v37 & v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
	if v40&v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	return int32(0)
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v36 = v35
	goto L9
L12:
	;
	return v6
L13:
	;
	v78 = int32(base.Ui32(v39) >> (uint(int32(1)) % 32))
	goto L15
L14:
	;
	if v39 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v78)
	goto L12
L16:
	;
	v78 = int32(0)
	goto L15
L17:
	;
	v45 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v45)
	v47 = int32(0)
	if v36 <= v47 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v50 = int32(8)
	v54 = v47
	goto L19
L19:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+(v8+v50)))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+(v7+v50)))))
	if v60 != v62 {
		goto L16
	} else {
		goto L21
	}
L20:
	;
	goto L12
L21:
	;
	v65 = v54 + int32(1)
	if v36 != v65 {
		v54 = v65
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
}
func F_ltree_crc32_sz(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[1562]))
	if v16 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = F_pg_newlocale_from_collation(m, int32(100))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v26 = v16
	goto L3
L3:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+3)))
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1562])) = v21
	v26 = v21
	goto L3
L6:
	;
	m.G0 = v13 + int32(16)
	return v176 ^ int32(-1)
L7:
	;
	v77 = l0
	v78 = l1
	v79 = int32(-1)
	goto L20
L8:
	;
	if int32(0) < l1 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v33 = int32(-1)
	if l1 <= int32(0) {
		v176 = v33
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v176 = int32(-1)
	goto L6
L12:
	;
	v36 = l0
	v37 = l1
	v38 = v33
	goto L13
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if base.Ui32((v46-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v176 = v66
	goto L6
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32((v55^v38)&int32(255)<<(uint(int32(2))%32))+uint32(_consts[1346])))
	v66 = v63 ^ int32(base.Ui32(v38)>>(uint(int32(8))%32))
	v67 = int32(1)
	if base.Ui32(v67) < base.Ui32(v37) {
		v36 = v36 + v67
		v37 = v37 - v67
		v38 = v66
		goto L13
	} else {
		goto L19
	}
L16:
	;
	v55 = v46 | int32(32)
	goto L18
L17:
	;
	v55 = v46
	goto L18
L18:
	;
	goto L15
L19:
	;
	goto L14
L20:
	;
	v91 = F_pg_mblen_range(m, v77, l0+l1)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v176 = v162
	goto L6
L22:
	;
	v171 = v78 - v91
	if int32(0) < v171 {
		v77 = v77 + v91
		v78 = v171
		v79 = v162
		goto L20
	} else {
		goto L33
	}
L23:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[1562]))
	v96 = F_pg_strfold(m, v13+int32(4), int32(12), v77, v91, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	if v96 == int32(0) {
		v162 = v79
		goto L22
	} else {
		goto L25
	}
L25:
	;
	if v96&int32(1) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32((v79^v104)&int32(255)<<(uint(int32(2))%32))+uint32(_consts[1346])))
	v118 = v112 ^ int32(base.Ui32(v79)>>(uint(int32(8))%32))
	v119 = v13 + int32(5)
	v120 = v96 - int32(1)
	goto L28
L27:
	;
	v118 = v79
	v119 = v13 + int32(4)
	v120 = v96
	goto L28
L28:
	;
	if v96 == int32(1) {
		v162 = v118
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v125 = v118
	v126 = v119
	v130 = v120
	goto L30
L30:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v136 = int32(255)
	v138 = int32(2)
	v141 = *(*int32)(unsafe.Add(mBase, uint32((v134^v125)&v136<<(uint(v138)%32))+uint32(_consts[1346])))
	v142 = int32(8)
	v144 = v141 ^ int32(base.Ui32(v125)>>(uint(v142)%32))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32((v144^v145)&v136<<(uint(v138)%32))+uint32(_consts[1346])))
	v155 = v152 ^ int32(base.Ui32(v144)>>(uint(v142)%32))
	v159 = v130 - v138
	if v159 != 0 {
		v125 = v155
		v126 = v126 + v138
		v130 = v159
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v162 = v155
	goto L22
L32:
	;
	goto L31
L33:
	;
	goto L21
}
func F_ltree_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	if v23 == int32(0) {
		v145 = v23
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v173 != v15 {
		goto L37
	} else {
		goto L38
	}
L5:
	;
	v171 = (v145 + int32(1)) * (v23 - v22) * int32(10)
	goto L4
L6:
	;
	if v22 == int32(0) {
		v145 = v23
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v28 = int32(8)
	v38 = v15 + v28
	v39 = v20 + v28
	v40 = v23
	v45 = v22
	goto L8
L8:
	;
	v46 = int32(2)
	v47 = v38 + v46
	v49 = v39 + v46
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39))))
	if base.Ui32(v50) < base.Ui32(v51) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v145 = v124
	goto L5
L10:
	;
	v53 = v50
	goto L12
L11:
	;
	v53 = v51
	goto L12
L12:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v53) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	if v115 != 0 {
		v171 = int32(1)
		goto L4
	} else {
		goto L31
	}
L14:
	;
	v115 = int32(0)
	goto L13
L15:
	;
	v89 = v84
	v90 = v85
	v91 = v86
	goto L25
L16:
	;
	if (v47|v49)&int32(3) != 0 {
		v84 = v47
		v85 = v49
		v86 = v53
		goto L15
	} else {
		goto L19
	}
L17:
	;
	v77 = v47
	v78 = v49
	v79 = v53
	goto L18
L18:
	;
	if v79 == int32(0) {
		goto L14
	} else {
		goto L24
	}
L19:
	;
	v61 = v47
	v62 = v49
	v63 = v53
	goto L20
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v66 != v67 {
		v84 = v61
		v85 = v62
		v86 = v63
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v77 = v72
	v78 = v70
	v79 = v74
	goto L18
L22:
	;
	v69 = int32(4)
	v70 = v62 + v69
	v72 = v61 + v69
	v74 = v63 - v69
	if base.Ui32(int32(3)) < base.Ui32(v74) {
		v61 = v72
		v62 = v70
		v63 = v74
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v84 = v77
	v85 = v78
	v86 = v79
	goto L15
L25:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v94 == v95 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v115 = v94 - v95
	goto L13
L27:
	;
	v97 = int32(1)
	v102 = v91 - v97
	if v102 != 0 {
		v89 = v89 + v97
		v90 = v90 + v97
		v91 = v102
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L14
L31:
	;
	if v50 != v51 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v117 = int32(10)
	v171 = (v40*v117 + v117) * (v50 - v51)
	goto L4
L33:
	;
	goto L34
L34:
	;
	v124 = v40 - int32(1)
	if v40 < int32(2) {
		v145 = v124
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v127 = int32(9)
	v129 = int32(131064)
	v137 = int32(1)
	if v137 < v45 {
		v38 = v38 + (v50+v127)&v129
		v39 = v39 + (v51+v127)&v129
		v40 = v124
		v45 = v45 - v137
		goto L8
	} else {
		goto L36
	}
L36:
	;
	goto L9
L37:
	;
	F_pfree(m, v15)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v177 != v20 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	F_pfree(m, v20)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	return base.B2i32(v171 != int32(0))
L44:
	;
	goto L43
}
func F_ltree_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == v2 {
		v31 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v31&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v18 == int32(0) {
		v31 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v21 != int32(7) {
		v31 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v24 != int32(17) {
		v31 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
	v31 = v27 ^ int32(1)
	goto L2
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = F_get_fn_opclass_options(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v40 = int32(8)
	goto L9
L9:
	;
	v42 = v12 + int32(8)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v44&int32(3) != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return int32(0)
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v40 = v39
	goto L9
L12:
	;
	v47 = int32(0)
	goto L14
L13:
	;
	v47 = v40
	goto L14
L14:
	;
	v48 = v42 + v47
	v50 = v10 + int32(8)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v52&int32(3) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v55 = int32(0)
	goto L17
L16:
	;
	v55 = v40
	goto L17
L17:
	;
	v56 = v50 + v55
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+4)))
	if v65 == int32(0) {
		v129 = v65
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v152&int32(1) != 0 {
		v166 = v50
		goto L37
	} else {
		goto L38
	}
L19:
	;
	v151 = (v129 + int32(1)) * (v65 - v64) * int32(10)
	goto L18
L20:
	;
	if v64 == int32(0) {
		v129 = v65
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v70 = int32(8)
	v74 = v48 + v70
	v75 = v56 + v70
	v78 = v65
	v82 = v64
	goto L22
L22:
	;
	v83 = int32(2)
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74))))
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75))))
	if base.Ui32(v87) < base.Ui32(v88) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v129 = v110
	goto L19
L24:
	;
	v110 = v78 - int32(1)
	if v78 < int32(2) {
		v129 = v110
		goto L19
	} else {
		goto L35
	}
L25:
	;
	v90 = v87
	goto L27
L26:
	;
	v90 = v88
	goto L27
L27:
	;
	v91 = F_memcmp(m, v74+v83, v75+v83, v90)
	mBase = m.M
	if v91 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v87 == v88 {
		goto L24
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v91 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v95 = int32(10)
	v151 = (v78*v95 + v95) * (v87 - v88)
	goto L18
L32:
	;
	v107 = int32(-10)
	goto L34
L33:
	;
	v107 = int32(10)
	goto L34
L34:
	;
	v151 = (v78 + int32(1)) * v107
	goto L18
L35:
	;
	v113 = int32(9)
	v115 = int32(131064)
	v123 = int32(1)
	if v123 < v82 {
		v74 = v74 + (v87+v113)&v115
		v75 = v75 + (v88+v113)&v115
		v78 = v110
		v82 = v82 - v123
		goto L22
	} else {
		goto L36
	}
L36:
	;
	goto L23
L37:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v167&int32(1) != 0 {
		v181 = v42
		goto L43
	} else {
		goto L44
	}
L38:
	;
	if v152&int32(2) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v158 = int32(0)
	goto L41
L40:
	;
	v158 = v40
	goto L41
L41:
	;
	v159 = v50 + v158
	if v152&int32(4) != 0 {
		v166 = v159
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v166 = v159 + int32(base.Ui32(v162)>>(uint(int32(2))%32))
	goto L37
L43:
	;
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+4)))
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	if v190 == int32(0) {
		v254 = v190
		goto L50
	} else {
		goto L51
	}
L44:
	;
	if v167&int32(2) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v173 = int32(0)
	goto L47
L46:
	;
	v173 = v40
	goto L47
L47:
	;
	v174 = v42 + v173
	if v167&int32(4) != 0 {
		v181 = v174
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v181 = v174 + int32(base.Ui32(v177)>>(uint(int32(2))%32))
	goto L43
L49:
	;
	v277 = int32(0)
	if v277 < v276 {
		goto L68
	} else {
		goto L69
	}
L50:
	;
	v276 = (v254 + int32(1)) * (v190 - v189) * int32(10)
	goto L49
L51:
	;
	if v189 == int32(0) {
		v254 = v190
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v195 = int32(8)
	v199 = v166 + v195
	v200 = v181 + v195
	v203 = v190
	v207 = v189
	goto L53
L53:
	;
	v208 = int32(2)
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199))))
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200))))
	if base.Ui32(v212) < base.Ui32(v213) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v254 = v235
	goto L50
L55:
	;
	v235 = v203 - int32(1)
	if v203 < int32(2) {
		v254 = v235
		goto L50
	} else {
		goto L66
	}
L56:
	;
	v215 = v212
	goto L58
L57:
	;
	v215 = v213
	goto L58
L58:
	;
	v216 = F_memcmp(m, v199+v208, v200+v208, v215)
	mBase = m.M
	if v216 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	if v212 == v213 {
		goto L55
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v216 < int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v220 = int32(10)
	v276 = (v203*v220 + v220) * (v212 - v213)
	goto L49
L63:
	;
	v232 = int32(-10)
	goto L65
L64:
	;
	v232 = int32(10)
	goto L65
L65:
	;
	v276 = (v203 + int32(1)) * v232
	goto L49
L66:
	;
	v238 = int32(9)
	v240 = int32(131064)
	v248 = int32(1)
	if v248 < v207 {
		v199 = v199 + (v212+v238)&v240
		v200 = v200 + (v213+v238)&v240
		v203 = v235
		v207 = v207 - v248
		goto L53
	} else {
		goto L67
	}
L67:
	;
	goto L54
L68:
	;
	v280 = v276
	goto L70
L69:
	;
	v280 = v277
	goto L70
L70:
	;
	v281 = int32(0)
	if v281 < v151 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v284 = v151
	goto L73
L72:
	;
	v284 = v281
	goto L73
L73:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v8))) = base.F32_convert_i32_u(v280 + v284)
	return v8
}
func F_ltree_prefix_eq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	if base.Ui32(l1) <= base.Ui32(l3) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v53 = int32(0)
	goto L3
L3:
	;
	return v53
L4:
	;
	v53 = base.B2i32(v50 == int32(0))
	goto L3
L5:
	;
	v50 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v12 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v13 = l0
	v14 = l2
	v15 = l1
	v16 = v12
	goto L12
L9:
	;
	v38 = l2
	v42 = int32(0)
	goto L10
L10:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v50 = v42 - v43
	goto L4
L11:
	;
	v38 = v33
	v42 = v35
	goto L10
L12:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v16 != v18 {
		v33 = v14
		v35 = v16
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v33 = v27
	v35 = int32(0)
	goto L11
L14:
	;
	if v18 == int32(0) {
		v33 = v14
		v35 = v16
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v23 = v15 - int32(1)
	if v23 == int32(0) {
		v33 = v14
		v35 = v16
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v26 = int32(1)
	v27 = v14 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v28 != 0 {
		v13 = v13 + v26
		v14 = v27
		v15 = v23
		v16 = v28
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
}
func F_parse_ltree(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = int32(1)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 == v3 {
		v70 = v15
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return v307
L2:
	;
	v80 = F_palloc(m, v70<<(uint(int32(4))%32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L15
	}
L3:
	;
	v22 = l0
	v23 = v3
	goto L4
L4:
	;
	v30 = F_pg_mblen_cstr(m, v22)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v41 = v37 + int32(1)
	if base.Ui32(v37) < base.Ui32(int32(65535)) {
		v70 = v41
		goto L2
	} else {
		goto L9
	}
L6:
	;
	return int32(0)
L7:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v37 = v23 + base.B2i32(v34 == int32(46))
	v38 = v22 + v30
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v39 != 0 {
		v22 = v38
		v23 = v37
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v44 = F_errsave_start(m, l1)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	if v44 == int32(0) {
		v307 = v3
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v41
	F_errmsg(m, int32(705525), v13+int32(32))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errsave_finish(m, l1, int32(516137), int32(67), int32(426376))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v307 = v3
	goto L1
L15:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v82 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v246 = v243 + int32(8)
	v247 = F_palloc0(m, v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L6
	} else {
		goto L65
	}
L17:
	;
	v235 = v80
	v243 = v3
	goto L16
L18:
	;
	goto L19
L19:
	;
	v86 = l0
	v88 = v80
	v89 = int32(0)
	v93 = v15
	v94 = v3
	goto L20
L20:
	;
	v96 = F_pg_mblen_cstr(m, v86)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L22
	}
L21:
	;
	if v183 != 0 {
		goto L53
	} else {
		goto L54
	}
L22:
	;
	if v89 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	v187 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v182)+12)) = v186 + v187
	v191 = v93 + v187
	v192 = v86 + v96
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v193 != 0 {
		v86 = v192
		v88 = v182
		v89 = v183
		v93 = v191
		v94 = v185
		goto L20
	} else {
		goto L52
	}
L24:
	;
	v100 = F_t_isalnum_cstr(m, v86)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L6
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v133 == int32(46) {
		goto L38
	} else {
		goto L39
	}
L27:
	;
	v112 = int32(0)
	v113 = F_errsave_start(m, l1)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L33
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v86
	v182 = v88
	v183 = int32(1)
	v185 = v94
	goto L23
L29:
	;
	if v100 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v102 == int32(95) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	if v102 != int32(45) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	if v113 == int32(0) {
		v307 = v112
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v93
	F_errmsg(m, int32(489052), v13)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	F_errsave_finish(m, l1, int32(516137), int32(84), int32(426376))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v307 = v112
	goto L1
L38:
	;
	v136 = int32(0)
	v138 = F_finish_nodeitem(m, v88, v86, v136, v93, l1)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v151 = int32(1)
	v152 = F_t_isalnum_cstr(m, v86)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	if v138 == int32(0) {
		v307 = v136
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v182 = v88 + int32(16)
	v183 = int32(0)
	v185 = (v142+int32(9))&int32(-8) + v94
	goto L23
L43:
	;
	if v152 != 0 {
		v182 = v88
		v183 = v151
		v185 = v94
		goto L23
	} else {
		goto L44
	}
L44:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v154 == int32(45) {
		v182 = v88
		v183 = v151
		v185 = v94
		goto L23
	} else {
		goto L45
	}
L45:
	;
	if v154 == int32(95) {
		v182 = v88
		v183 = v151
		v185 = v94
		goto L23
	} else {
		goto L46
	}
L46:
	;
	v159 = int32(0)
	v160 = F_errsave_start(m, l1)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	if v160 == int32(0) {
		v307 = v159
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v93
	F_errmsg(m, int32(489052), v13+int32(16))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	F_errsave_finish(m, l1, int32(516137), int32(96), int32(426376))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	v307 = v159
	goto L1
L52:
	;
	goto L21
L53:
	;
	v194 = int32(0)
	v196 = F_finish_nodeitem(m, v182, v192, v194, v191, l1)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L6
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v80 == v182 {
		v235 = v80
		v243 = v185
		goto L16
	} else {
		goto L58
	}
L56:
	;
	if v196 == int32(0) {
		v307 = v194
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v235 = v182 + int32(16)
	v243 = (v202+int32(9))&int32(-8) + v185
	goto L16
L58:
	;
	v209 = int32(0)
	v210 = F_errsave_start(m, l1)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	if v210 == int32(0) {
		v307 = v209
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(221474), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	F_errdetail(m, int32(600971), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	F_errsave_finish(m, l1, int32(516137), int32(118), int32(426376))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v307 = v209
	goto L1
L65:
	;
	v249 = v235 - v80
	v251 = int32(base.Ui32(v249) >> (uint(int32(4)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v247)+4)) = uint16(v251)
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v246 << (uint(int32(2)) % 32)
	if v249&int32(1048560) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v260 = v80
	v262 = v247 + int32(8)
	goto L69
L67:
	;
	goto L68
L68:
	;
	F_pfree(m, v80)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L6
	} else {
		goto L76
	}
L69:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	*(*uint16)(unsafe.Add(mBase, uint32(v262))) = uint16(v270)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	if v270 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L68
L71:
	;
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+4)))
	v286 = v260 + int32(16)
	if (v286-v80)>>(uint(int32(4))%32) < v284 {
		v260 = v286
		v262 = v262 + (v270&int32(65535)+int32(9))&int32(131064)
		goto L69
	} else {
		goto L75
	}
L72:
	;
	v275 = F__emscripten_memcpy_bulkmem(m, v262+int32(2), v274, v270)
	mBase = m.M
	goto L74
L73:
	;
	goto L74
L74:
	;
	goto L71
L75:
	;
	goto L70
L76:
	;
	v307 = v247
	goto L1
}
