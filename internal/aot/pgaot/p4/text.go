package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_replace_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	v10 = m.G0
	v12 = v10 - int32(1088)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
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
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = F_pg_detoast_datum_packed(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v25 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	m.G0 = v12 + int32(1088)
	return v218
L6:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v57 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L7:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v31 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v42 = int32(1)
	if v25&v42 != 0 {
		v54 = int32(base.Ui32(v25)>>(uint(v42)%32)) - v42
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v34 = int32(16)
	goto L12
L11:
	;
	v34 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v41 = int32(4)
	goto L15
L14:
	;
	v41 = v34
	goto L15
L15:
	;
	v54 = v41
	goto L6
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	if base.B2i32(v54 <= int32(0))|base.B2i32(v86 <= int32(0)) != 0 {
		v218 = v15
		goto L5
	} else {
		goto L28
	}
L18:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v63 == int32(18) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v74 = int32(1)
	if v57&v74 != 0 {
		v86 = int32(base.Ui32(v57)>>(uint(v74)%32)) - v74
		goto L17
	} else {
		goto L27
	}
L21:
	;
	v66 = int32(16)
	goto L23
L22:
	;
	v66 = int32(0)
	goto L23
L23:
	;
	if base.Ui32((v63-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v73 = int32(4)
	goto L26
L25:
	;
	v73 = v66
	goto L26
L26:
	;
	v86 = v73
	goto L17
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v86 = int32(base.Ui32(v80)>>(uint(int32(2))%32)) - int32(4)
	goto L17
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v92 = v12 + int32(16)
	F_text_position_setup(m, v15, v20, v90, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v95 = F_text_position_next(m, v92)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v95 == int32(0) {
		v218 = v15
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v99 = int32(1)
	v100 = v15 + v99
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v103&v99 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v106 = v100
	goto L34
L33:
	;
	v106 = v15 + int32(4)
	goto L34
L34:
	;
	v110 = v23 + int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1068))
	F_initStringInfo(m, v12)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v114 = v106
	v115 = v111
	goto L36
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_replace_text[0]))
	if v124 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v170 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L38:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_appendBinaryStringInfo(m, v12, v114, v115-v114)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v132 = v130 & int32(1)
	if v132 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v133 = v110
	goto L45
L44:
	;
	v133 = v23 + int32(4)
	goto L45
L45:
	;
	if v130 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	F_appendBinaryStringInfo(m, v12, v133, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L57
	}
L47:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v139 == int32(18) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v150 = int32(1)
	if v132 != 0 {
		v160 = int32(base.Ui32(v130)>>(uint(v150)%32)) - v150
		goto L46
	} else {
		goto L56
	}
L50:
	;
	v142 = int32(16)
	goto L52
L51:
	;
	v142 = int32(0)
	goto L52
L52:
	;
	if base.Ui32((v139-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v149 = int32(4)
	goto L55
L54:
	;
	v149 = v142
	goto L55
L55:
	;
	v160 = v149
	goto L46
L56:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v160 = int32(base.Ui32(v154)>>(uint(int32(2))%32)) - int32(4)
	goto L46
L57:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1072))
	v164 = v115 + v163
	v167 = F_text_position_next(m, v12+int32(16))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v167 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1068))
	v114 = v164
	v115 = v169
	goto L36
L60:
	;
	goto L61
L61:
	;
	goto L37
L62:
	;
	F_appendBinaryStringInfo(m, v12, v164, v195+v15-v164)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L73
	}
L63:
	;
	v174 = int32(18)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	if v176 == v174 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v187 = int32(1)
	if v170&v187 != 0 {
		v195 = int32(base.Ui32(v170) >> (uint(v187) % 32))
		goto L62
	} else {
		goto L72
	}
L66:
	;
	v179 = v174
	goto L68
L67:
	;
	v179 = int32(2)
	goto L68
L68:
	;
	if base.Ui32((v176-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v186 = int32(6)
	goto L71
L70:
	;
	v186 = v179
	goto L71
L71:
	;
	v195 = v186
	goto L62
L72:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v195 = int32(base.Ui32(v191) >> (uint(int32(2)) % 32))
	goto L62
L73:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v203 = v201 + int32(4)
	v204 = F_palloc(m, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v203 << (uint(int32(2)) % 32)
	if v201 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	base.MemoryCopy(m, v204+int32(4), v200, v201)
	goto L77
L76:
	;
	goto L77
L77:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	F_pfree(m, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v218 = v204
	goto L5
}
func F_text_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = v10 + int32(1)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v25 = v23 & int32(1)
			if v25 != 0 {
				v26 = v15
			} else {
				v26 = v10 + int32(4)
			}
			if v23 == int32(1) {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v32 == int32(18) {
					v35 = int32(16)
				} else {
					v35 = int32(0)
				}
				if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v42 = int32(4)
				} else {
					v42 = v35
				}
				v53 = v42
			} else {
				v43 = int32(1)
				if v25 != 0 {
					v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v54 = int32(1)
			v55 = v17 + v54
			if v19&v54 != 0 {
				v60 = v55
			} else {
				v60 = v17 + int32(4)
			}
			if v19 == int32(1) {
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
				if v66 == int32(18) {
					v69 = int32(16)
				} else {
					v69 = int32(0)
				}
				if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v76 = int32(4)
				} else {
					v76 = v69
				}
				v89 = v76
			} else {
				v77 = int32(1)
				if v19&v77 != 0 {
					v89 = int32(base.Ui32(v19)>>(uint(v77)%32)) - v77
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v89 = int32(base.Ui32(v83)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v90 = F_varstr_cmp(m, v26, v53, v60, v89, v20)
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v92 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v96 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v90) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v90) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v96 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v90) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v90) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_text_reverse(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = int32(1)
	v12 = v7 + v11
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v17 = v15 & v11
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v18 = v12
	goto L5
L4:
	;
	v18 = v7 + int32(4)
	goto L5
L5:
	;
	if v15 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v47 = v45 + int32(4)
	v48 = F_palloc(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v24 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v35 = int32(1)
	if v17 != 0 {
		v45 = int32(base.Ui32(v15)>>(uint(v35)%32)) - v35
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v27 = int32(16)
	goto L12
L11:
	;
	v27 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v34 = int32(4)
	goto L15
L14:
	;
	v34 = v27
	goto L15
L15:
	;
	v45 = v34
	goto L6
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v47 << (uint(int32(2)) % 32)
	v53 = v47 + v48
	v54 = v18 + v45
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_text_reverse[0]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57*int32(28))+uint32(_c_F_text_reverse[1])))
	goto L19
L18:
	;
	return v48
L19:
	;
	if v62 <= int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v45 <= int32(0) {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v45 <= int32(0) {
		goto L18
	} else {
		goto L27
	}
L23:
	;
	v67 = v18
	v68 = v53
	goto L24
L24:
	;
	v72 = int32(1)
	v73 = v68 - v72
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v74)
	v77 = v67 + v72
	if base.Ui32(v77) < base.Ui32(v54) {
		v67 = v77
		v68 = v73
		goto L24
	} else {
		goto L26
	}
L25:
	;
	goto L18
L26:
	;
	goto L25
L27:
	;
	v81 = v18
	v82 = v53
	goto L28
L28:
	;
	v86 = F_pg_mblen_range(m, v81, v54)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L18
L30:
	;
	v88 = v82 - v86
	if v86 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	base.MemoryCopy(m, v88, v81, v86)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v90 = v81 + v86
	if base.Ui32(v90) < base.Ui32(v54) {
		v81 = v90
		v82 = v88
		goto L28
	} else {
		goto L34
	}
L34:
	;
	goto L29
}
func F_text_smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v9 + int32(1)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v24 = v22 & int32(1)
			if v24 != 0 {
				v25 = v14
			} else {
				v25 = v9 + int32(4)
			}
			if v22 == int32(1) {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v31 == int32(18) {
					v34 = int32(16)
				} else {
					v34 = int32(0)
				}
				if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v41 = int32(4)
				} else {
					v41 = v34
				}
				v52 = v41
			} else {
				v42 = int32(1)
				if v24 != 0 {
					v52 = int32(base.Ui32(v22)>>(uint(v42)%32)) - v42
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v53 = int32(1)
			v54 = v16 + v53
			if v18&v53 != 0 {
				v59 = v54
			} else {
				v59 = v16 + int32(4)
			}
			if v18 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				if v65 == int32(18) {
					v68 = int32(16)
				} else {
					v68 = int32(0)
				}
				if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v75 = int32(4)
				} else {
					v75 = v68
				}
				v88 = v75
			} else {
				v76 = int32(1)
				if v18&v76 != 0 {
					v88 = int32(base.Ui32(v18)>>(uint(v76)%32)) - v76
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v89 = F_varstr_cmp(m, v25, v52, v59, v88, v19)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int32(0)
			} else {
				if v89 < int32(0) {
					v93 = v9
				} else {
					v93 = v16
				}
				return v93
			}
		}
	}
}
func F_text_to_cstring(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	v5 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
		if v9 == int32(1) {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
			if v15 == int32(18) {
				v18 = int32(16)
			} else {
				v18 = int32(0)
			}
			if base.Ui32((v15-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v25 = int32(4)
			} else {
				v25 = v18
			}
			v38 = v25
		} else {
			v26 = int32(1)
			if v9&v26 != 0 {
				v38 = int32(base.Ui32(v9)>>(uint(v26)%32)) - v26
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v38 = int32(base.Ui32(v32)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v41 = F_palloc(m, v38+int32(1))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			if v38 != 0 {
				v43 = int32(1)
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
				if v45&v43 != 0 {
					v48 = v43
				} else {
					v48 = int32(4)
				}
				base.MemoryCopy(m, v41, v5+v48, v38)
			} else {
			}
			v52 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v38+v41))) = uint8(v52)
			if l0 != v5 {
				F_pfree(m, v5)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					return v41
				}
			} else {
				return v41
			}
		}
	}
}
func F_text_to_stavalues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v14 = *(*int64)(unsafe.Add(mBase, _c_F_text_to_stavalues[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_text_to_stavalues[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v17
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(v19)
	v22 = F_text_to_cstring(m, l2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v11)+36)) = int64(0)
		v28 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)) = uint8(v28)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+68)) = uint8(v28)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l4
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)) = uint8(v28)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+52)) = uint8(v28)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v22
		v39 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+46)) = uint16(v39)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l1
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v11+int32(28))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			F_pfree(m, v22)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
				if v52 == int32(1) {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(19)
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
					F_ThrowErrorData(m, v58)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						v61 = int32(0)
						v89 = v61
						v90 = v61
						*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v90)
						m.G0 = v11 + int32(112)
						return v89
					}
				} else {
					v63 = F_pg_detoast_datum(m, v48)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v65 = F_array_contains_nulls(m, v63)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							if v65 == int32(0) {
								v89 = v48
								v90 = v19
								*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v90)
								m.G0 = v11 + int32(112)
								return v89
							} else {
								v69 = int32(0)
								v73 = F_errstart(m, int32(19), v69)
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									if v73 == int32(0) {
										v89 = v69
										v90 = v69
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v90)
										m.G0 = v11 + int32(112)
										return v89
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
											F_errmsg(m, int32(_a_F_text_to_stavalues_0), v11)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_text_to_stavalues_1), int32(738), int32(_a_F_text_to_stavalues_2))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return int32(0)
												} else {
													v89 = v69
													v90 = v69
													*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v90)
													m.G0 = v11 + int32(112)
													return v89
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
