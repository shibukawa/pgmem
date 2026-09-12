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
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(1)
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	F_s_lock(m, v19+int32(20), int32(518750), int32(2237), int32(383653))
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
	v30 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
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
	v125 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+20)) = int32(0)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+28)) = v128 - int32(1)
	m.G0 = v11 + int32(16)
	return v121
L10:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	F_s_lock(m, v114+int32(20), int32(518750), v112, int32(383653))
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
	*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(22)
	v79 = int32(-1)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v54 = F_OpenTransientFile(m, int32(119677), int32(66))
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
	v65 = F_pwrite(m, v54, int32(4120644), v61, base.I64_extend_i32_u(v36))
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
	v71 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
	if v72 == int32(0) {
		v121 = v61
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v109 = v61
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
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(119677)
	F_errmsg(m, int32(314307), v11)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(518750), int32(2282), int32(383653))
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
	v101 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = int32(1)
	v105 = int32(0)
	if v102 == v105 {
		v121 = v105
		goto L9
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v109 = v105
	v112 = int32(2288)
	goto L10
L33:
	;
	v121 = v109
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
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
	v51 = F_palloc(m, v46<<(uint(int32(1))%32)+int32(7))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
	v19 = int32(4)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v21&int32(254) == int32(2) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v34 = int32(1)
	if v16&v34 != 0 {
		v46 = int32(base.Ui32(v16)>>(uint(v34)%32)) - v34
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v30 = v19
	goto L9
L8:
	;
	v30 = base.B2i32(v21 == int32(18)) << (uint(v19) % 32)
	goto L9
L9:
	;
	if v21 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v33 = v19
	goto L12
L11:
	;
	v33 = v30
	goto L12
L12:
	;
	v46 = v33
	goto L1
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L14:
	;
	v53 = int32(4)
	v54 = v51 + v53
	v55 = int32(1)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v57&v55 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v94 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v94)
	v97 = v85 + int32(1)
	if v46 == int32(0) {
		v227 = v85
		v228 = v97
		goto L25
	} else {
		goto L26
	}
L16:
	;
	v60 = v55
	goto L18
L17:
	;
	v60 = v53
	goto L18
L18:
	;
	v61 = v12 + v60
	v62 = v61 + v46
	if base.Ui32(v62) <= base.Ui32(v61) {
		v85 = v54
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v65 = v61
	goto L21
L20:
	;
	v80 = int32(69)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)) = uint8(v80)
	v85 = v51 + int32(5)
	goto L15
L21:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v74 == int32(92) {
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v85 = v54
	goto L15
L23:
	;
	v78 = v65 + int32(1)
	if v78 != v62 {
		v65 = v78
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v236 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v228))) = uint8(v236)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = (v227-v54)<<(uint(int32(2))%32) + int32(24)
	return v51
L26:
	;
	v101 = v46 & int32(3)
	if v101 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if base.Ui32(v46) < base.Ui32(int32(4)) {
		v227 = v136
		v228 = v137
		goto L25
	} else {
		goto L38
	}
L28:
	;
	v135 = v61
	v136 = v85
	v137 = v97
	v141 = v46
	goto L27
L29:
	;
	goto L30
L30:
	;
	v104 = v61
	v105 = v85
	v106 = v97
	v109 = int32(0)
	v110 = v46
	goto L31
L31:
	;
	v115 = v110 - int32(1)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v116 == int32(92) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v135 = v131
	v136 = v126
	v137 = v129
	v141 = v115
	goto L27
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v125)
	v128 = int32(1)
	v129 = v126 + v128
	v131 = v104 + v128
	v133 = v109 + v128
	if v133 != v101 {
		v104 = v131
		v105 = v126
		v106 = v129
		v109 = v133
		v110 = v115
		goto L31
	} else {
		goto L37
	}
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v116)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v125 = v122
	v126 = v105 + int32(2)
	goto L33
L35:
	;
	if v116 == int32(39) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v125 = v116
	v126 = v106
	goto L33
L37:
	;
	goto L32
L38:
	;
	v147 = v135
	v148 = v136
	v149 = v137
	v153 = v141
	goto L39
L39:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if base.B2i32(v157 != int32(92))&base.B2i32(v157 != int32(39)) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v227 = v218
	v228 = v221
	goto L25
L41:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v157)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v169 = v148 + int32(2)
	v170 = v166
	goto L43
L42:
	;
	v169 = v149
	v170 = v157
	goto L43
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v170)
	v173 = v147 + int32(1)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	if v174 == int32(92) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v185)
	v189 = v147 + int32(2)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v190 == int32(92) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)) = uint8(v174)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v185 = v182
	v186 = v169 + int32(2)
	goto L44
L46:
	;
	if v174 == int32(39) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v185 = v174
	v186 = v169 + int32(1)
	goto L44
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v202))) = uint8(v201)
	v205 = v147 + int32(3)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v206 == int32(92) {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v186)+1)) = uint8(v190)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v201 = v198
	v202 = v186 + int32(2)
	goto L48
L50:
	;
	if v190 == int32(39) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v201 = v190
	v202 = v186 + int32(1)
	goto L48
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v218))) = uint8(v217)
	v221 = v218 + int32(1)
	v222 = int32(4)
	v225 = v153 - v222
	if v225 != 0 {
		v147 = v147 + v222
		v148 = v218
		v149 = v221
		v153 = v225
		goto L39
	} else {
		goto L56
	}
L53:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+1)) = uint8(v206)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v217 = v214
	v218 = v202 + int32(2)
	goto L52
L54:
	;
	if v206 == int32(39) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v217 = v206
	v218 = v202 + int32(1)
	goto L52
L56:
	;
	goto L40
}
