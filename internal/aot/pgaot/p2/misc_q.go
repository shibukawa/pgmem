package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
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
				v18 = F_palloc0_mul(m, int32(4), int32(2))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v18
					v23 = F_QT2QTN(m, l0+int32(12), l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v25))) = v23
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v29
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
						if v31 == int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(1)
							return v6
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(2)
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v42 = F_QT2QTN(m, l0+v38*int32(12), l1)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v42
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v46 | v49
								return v6
							}
						}
					}
				}
			} else {
				if l1 == int32(0) {
					return v6
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1 + int32(base.Ui32(v55)>>(uint(int32(12))%32))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(1) << (uint(v61) % 32)
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
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_store[0]))
	v17 = base.AtomicRmwXchg32(m, v14, int32(140), int32(1))
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, v14+int32(140), int32(_a_F_qtext_store_0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_store[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+148))
	v28 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+148)) = v27 + v28
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+144))
	v32 = l1 + v31
	*(*int32)(unsafe.Add(mBase, uint32(v26)+144)) = v32 + v28
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v36
	goto L8
L7:
	;
	goto L8
L8:
	;
	v38 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v26)+140)), uint32(v38))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31
	if base.Ui32(v31^int32(2147483647)) <= base.Ui32(l1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_store[0]))
	v96 = base.AtomicRmwXchg32(m, v93, int32(140), int32(1))
	if v96 != 0 {
		goto L30
	} else {
		goto L31
	}
L10:
	;
	v71 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L19
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_qtext_store[1])) = int32(22)
	v68 = int32(-1)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v51 = F_OpenTransientFile(m, int32(_a_F_qtext_store_1), int32(66))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	if v51 < int32(0) {
		v68 = v51
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v56 = F_pwrite(m, v51, l0, l1, base.I64_extend_i32_u(v31))
	mBase = m.M
	if v56 != l1 {
		v68 = v51
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v58 = int32(1)
	v62 = F_pwrite(m, v51, int32(_a_F_qtext_store_5), v58, base.I64_extend_i32_u(v32))
	mBase = m.M
	if v62 != v58 {
		v68 = v51
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v65 = F_CloseTransientFile(m, v51)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v90 = v58
	goto L9
L19:
	;
	if v71 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if int32(0) <= v68 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_qtext_store_1)
	F_errmsg(m, int32(_a_F_qtext_store_2), v11)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_qtext_store_3), int32(2285), int32(_a_F_qtext_store_4))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v87 = F_CloseTransientFile(m, v68)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v90 = int32(0)
	goto L9
L29:
	;
	goto L28
L30:
	;
	F_s_lock(m, v93+int32(140), int32(_a_F_qtext_store_0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_qtext_store[0]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+148)) = v104 - int32(1)
	v108 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v103)+140)), uint32(v108))
	m.G0 = v11 + int32(16)
	return v90
L33:
	;
	goto L32
}
func F_quote_literal(m *base.Module, l0 int32) int64 {
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	return base.I64_extend_i32_u(v50)
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
