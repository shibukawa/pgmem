package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CommentObject(m *base.Module, l0 int32, l1 int32) {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = *(*int64)(unsafe.Add(mBase, _c_F_CommentObject[0]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v14
	v17 = l0 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_CommentObject[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v21 == int32(9) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return
L2:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	if v236 == int32(0) {
		goto L1
	} else {
		goto L58
	}
L3:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_CreateComments(m, v223, v224, v225, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L10
	} else {
		goto L57
	}
L4:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v118 = int32(0)
	v120 = m.G0
	v122 = v120 - int32(128)
	m.G0 = v122
	v124 = int32(1)
	if v117 == v118 {
		v145 = v118
		v146 = v124
		goto L29
	} else {
		goto L30
	}
L5:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+48))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+119)))
	v79 = v77 - int32(99)
	if int32(1)<<(uint(v79)%32)&int32(_a_F_CommentObject_3) != 0 {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v59 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L15
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v27 = F_get_database_oid(m, v25, int32(1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v33 = v21
	goto L9
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_get_object_address(m, l0, v33, v34, v11+int32(44), int32(4), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L13
	}
L10:
	;
	return
L11:
	;
	if v27 == int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v33 = v31
	goto L9
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_CommentObject[2]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v45
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v47
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	F_check_object_ownership(m, v42, v44, v11+int32(32), v43, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v54 - int32(6) {
	case 0:
		goto L5
	default:
		goto L3
	case 3, 27, 36:
		goto L4
	}
L15:
	;
	if v59 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v25
	F_errmsg(m, int32(_a_F_CommentObject_0), v11)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_CommentObject_1), int32(61), int32(_a_F_CommentObject_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	v87 = base.B2i32(base.Ui32(v79) <= base.Ui32(int32(19)))
	goto L22
L21:
	;
	v87 = int32(0)
	goto L22
L22:
	;
	if v87 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v96 + int32(4)
	F_errmsg(m, int32(_a_F_CommentObject_4), v11+int32(16))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+48))
	v107 = int32(*(*int8)(unsafe.Add(mBase, uint32(v106)+119)))
	F_errdetail_relkind_not_supported(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_CommentObject_1), int32(103), int32(_a_F_CommentObject_2))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_ScanKeyInit(m, v122+int32(32), int32(1), int32(3), int32(184), v115)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L10
	} else {
		goto L33
	}
L30:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v127 == int32(0) {
		v145 = v118
		v146 = v124
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+18)) = uint8(v130)
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+14)) = uint8(v133)
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+16)) = uint16(v130)
	v138 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+12)) = uint16(v138)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+24)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v122)+20)) = v115
	v142 = F_cstring_to_text(m, v117)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+28)) = v142
	v145 = v133
	v146 = v130
	goto L29
L33:
	;
	F_ScanKeyInit(m, v122+int32(80), int32(2), int32(3), int32(184), v116)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v163 = F_table_open(m, int32(2396), int32(3))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L10
	} else {
		goto L36
	}
L35:
	;
	F_systable_endscan(m, v171)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L46
	}
L36:
	;
	v171 = F_systable_beginscan(m, v163, int32(2397), int32(1), int32(0), int32(2), v122+int32(32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	v173 = F_systable_getnext(m, v171)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	if v173 == int32(0) {
		v194 = v118
		goto L35
	} else {
		goto L39
	}
L39:
	;
	if v146 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_CatalogTupleDelete(m, v163, v173+int32(4))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L10
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v163)+52))
	v190 = F_heap_modify_tuple(m, v173, v183, v122+int32(20), v122+int32(16), v122+int32(12))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L10
	} else {
		goto L44
	}
L43:
	;
	v194 = v118
	goto L35
L44:
	;
	F_CatalogTupleUpdate(m, v163, v173+int32(4), v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	v194 = v190
	goto L35
L46:
	;
	v197 = int32(0)
	if base.B2i32(v145 == v197)|base.B2i32(v194 != v197) == v197 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v163)+52))
	v209 = F_heap_form_tuple(m, v204, v122+int32(20), v122+int32(16))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L50
	}
L48:
	;
	v213 = v194
	goto L49
L49:
	;
	if v213 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	F_CatalogTupleInsert(m, v163, v209)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	v213 = v209
	goto L49
L52:
	;
	F_pfree(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L10
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	F_sequence_close(m, v163, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L10
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	m.G0 = v122 + int32(128)
	goto L2
L57:
	;
	goto L2
L58:
	;
	F_relation_close(m, v236, int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	goto L1
}
func F_CompareFurthestCandidates(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)+20))
	if base.F64_lt(v7, v8) != 0 {
		v11 = int32(-1)
	} else {
		v11 = base.F64_gt(v7, v8)
	}
	return v11
}
func F_CompareLists(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_lt(v9, v10) != 0 {
		v12 = int32(-1)
	} else {
		v12 = int32(0)
	}
	if base.F64_gt(v9, v10) != 0 {
		v14 = int32(1)
	} else {
		v14 = v12
	}
	return v14
}
func F_CompareNearestDiscardedCandidates(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)+20))
	if base.F64_gt(v9, v10) != 0 {
		v12 = int32(-1)
	} else {
		v12 = int32(0)
	}
	if base.F64_lt(v9, v10) != 0 {
		v14 = int32(1)
	} else {
		v14 = v12
	}
	return v14
}
func F_CompleteCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v193 int32
	_ = v193
	v9 = l8
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v63
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v66 != 0 {
		goto L28
	} else {
		goto L29
	}
L2:
	;
	v62 = l1
	v63 = v14
	goto L1
L3:
	;
	goto L4
L4:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v20 != v15 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v57 = F_AllocSetContextCreateInternal(m, v15, int32(_a_F_CompleteCachedPlan_0), int32(0), int32(1024), int32(_a_F_CompleteCachedPlan_1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L25
	} else {
		goto L26
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0])) = l2
	v62 = l1
	v63 = l2
	goto L1
L9:
	;
	if v20 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	if v15 != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v25 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v24 == int32(0) {
		goto L12
	} else {
		goto L18
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v24
	goto L14
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v24
	goto L14
L18:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v30
	goto L12
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v15
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v37
	if v37 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = int32(0)
	goto L11
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = l2
	goto L24
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l2
	goto L8
L25:
	;
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0])) = v57
	v60 = F_copyObjectImpl(m, l1)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v62 = v60
	v63 = v57
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0])) = v15
	if int32(0) < l4 {
		goto L47
	} else {
		goto L48
	}
L29:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v67 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	F_extract_query_dependencies(m, v62, l0-int32(-64), l0+int32(68), l0+int32(85))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L25
	} else {
		goto L44
	}
L31:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	switch v71 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v75 = int32(1)
		goto L35
	default:
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v76 == int32(0) {
		goto L28
	} else {
		goto L38
	}
L34:
	;
	if v75 != 0 {
		goto L30
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v75 = int32(0)
	goto L35
L37:
	;
	goto L28
L38:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v80 != int32(6) {
		v95 = int32(1)
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v95&int32(1) == int32(0) {
		goto L28
	} else {
		goto L43
	}
L40:
	;
	goto L39
L41:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v76)+28))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v87 = v85 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v87) {
		v95 = int32(0)
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v95 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v87)) % 64)))
	goto L40
L43:
	;
	goto L30
L44:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v110
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[2])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v113)
	v115 = F_GetSearchPathMatcher(m, v63)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v115
	goto L28
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v9)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l4
	v140 = F_ChoosePortalStrategy(m, v62)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L25
	} else {
		goto L59
	}
L47:
	;
	v124 = l4 << (uint(int32(2)) % 32)
	v125 = F_palloc(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L25
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	goto L46
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v125
	if v124 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L46
L52:
	;
	v128 = F__emscripten_memcpy_bulkmem(m, v125, l3, v124)
	mBase = m.M
	goto L54
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v180
	*(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0])) = v14
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v193)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)) = uint8(v193)
	return
L56:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+28))
	v176 = F_UtilityTupleDescriptor(m, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L25
	} else {
		goto L65
	}
L57:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v150 = int32(0)
	goto L61
L58:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+76))
	v145 = F_ExecCleanTypeFromTL(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L25
	} else {
		goto L60
	}
L59:
	;
	switch v140 {
	case 0, 2:
		goto L58
	case 1:
		goto L57
	case 3:
		goto L56
	default:
		v180 = int32(0)
		goto L55
	}
L60:
	;
	v180 = v145
	goto L55
L61:
	;
	v163 = int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v150<<(uint(int32(2))%32)+v147)))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+24)))
	if v167 != v163 {
		v150 = v150 + v163
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166)+96))
	v171 = F_ExecCleanTypeFromTL(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L25
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v180 = v171
	goto L55
L65:
	;
	v180 = v176
	goto L55
}
func F_ConditionVariableSleep(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = F_ConditionVariableTimedSleep(m, l0, int32(-1), l1)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_CopyLoadRawBuf(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int64
	_ = v188
	var v190 int64
	_ = v190
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v9 = v7 - v8
	if v8 <= int32(0) {
		v163 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+332)) = v163
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v167 == v168 {
		goto L50
	} else {
		goto L51
	}
L2:
	;
	if v9 <= int32(0) {
		v163 = v9
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v15 = v14 + v8
	if v14 == v15 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v163 = v160 - v161
	goto L1
L5:
	;
	goto L4
L6:
	;
	v19 = v14 + v9
	if base.Ui32(v15-v19) <= base.Ui32(int32(0)-v9<<(uint(int32(1))%32)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = F___memcpy(m, v14, v15, v9)
	mBase = m.M
	goto L4
L8:
	;
	goto L9
L9:
	;
	v29 = (v14 ^ v15) & int32(3)
	if base.Ui32(v14) < base.Ui32(v15) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if v131 == int32(0) {
		goto L5
	} else {
		goto L46
	}
L11:
	;
	if base.Ui32(v109) <= base.Ui32(int32(3)) {
		v130 = v108
		v131 = v109
		v132 = v110
		goto L10
	} else {
		goto L42
	}
L12:
	;
	if v29 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v29 != 0 {
		v91 = v9
		goto L25
	} else {
		goto L26
	}
L15:
	;
	v130 = v15
	v131 = v9
	v132 = v14
	goto L10
L16:
	;
	goto L17
L17:
	;
	if v14&int32(3) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v108 = v15
	v109 = v9
	v110 = v14
	goto L11
L19:
	;
	goto L20
L20:
	;
	v36 = v15
	v37 = v9
	v38 = v14
	goto L21
L21:
	;
	if v37 == int32(0) {
		goto L5
	} else {
		goto L23
	}
L22:
	;
	v108 = v45
	v109 = v47
	v110 = v49
	goto L11
L23:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v42)
	v44 = int32(1)
	v45 = v36 + v44
	v47 = v37 - v44
	v49 = v38 + v44
	if v49&int32(3) != 0 {
		v36 = v45
		v37 = v47
		v38 = v49
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	if v91 == int32(0) {
		goto L5
	} else {
		goto L38
	}
L26:
	;
	if v19&int32(3) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v56 = v9
	goto L30
L28:
	;
	v71 = v9
	goto L29
L29:
	;
	if base.Ui32(v71) <= base.Ui32(int32(3)) {
		v91 = v71
		goto L25
	} else {
		goto L34
	}
L30:
	;
	if v56 == int32(0) {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	v71 = v62
	goto L29
L32:
	;
	v62 = v56 - int32(1)
	v63 = v14 + v62
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v62))))
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v65)
	if v63&int32(3) != 0 {
		v56 = v62
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v78 = v71
	goto L35
L35:
	;
	v82 = v78 - int32(4)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v15+v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v14+v82))) = v85
	if base.Ui32(int32(3)) < base.Ui32(v82) {
		v78 = v82
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v91 = v82
	goto L25
L37:
	;
	goto L36
L38:
	;
	v98 = v91
	goto L39
L39:
	;
	v102 = v98 - int32(1)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v102))))
	*(*uint8)(unsafe.Add(mBase, uint32(v14+v102))) = uint8(v105)
	if v102 != 0 {
		v98 = v102
		goto L39
	} else {
		goto L41
	}
L40:
	;
	goto L5
L41:
	;
	goto L40
L42:
	;
	v115 = v108
	v116 = v109
	v117 = v110
	goto L43
L43:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v119
	v121 = int32(4)
	v122 = v115 + v121
	v124 = v117 + v121
	v126 = v116 - v121
	if base.Ui32(int32(3)) < base.Ui32(v126) {
		v115 = v122
		v116 = v126
		v117 = v124
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v130 = v122
	v131 = v126
	v132 = v124
	goto L10
L45:
	;
	goto L44
L46:
	;
	v137 = v130
	v138 = v131
	v139 = v132
	goto L47
L47:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v141)
	v143 = int32(1)
	v148 = v138 - v143
	if v148 != 0 {
		v137 = v137 + v143
		v138 = v148
		v139 = v139 + v143
		goto L47
	} else {
		goto L49
	}
L48:
	;
	goto L5
L49:
	;
	goto L48
L50:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v173 - v170
	goto L52
L51:
	;
	goto L52
L52:
	;
	v180 = F_CopyGetData(m, l0, v163+v167, int32(_a_F_CopyLoadRawBuf_0)-v163)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	return
L54:
	;
	v182 = v180 + v9
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v185 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v182+v183))) = uint8(v185)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+332)) = v182
	v188 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
	v190 = v188 + base.I64_extend_i32_s(v180)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+344)) = v190
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[0]))
	if v195 == v185 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v180 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	goto L55
L57:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[1])))
	if v199 != int32(1) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v202 = int32(_a_F_CopyLoadRawBuf_1)
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[2]))
	v205 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[2])) = v204 + v205
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v208 + v205
	*(*int64)(unsafe.Add(mBase, uint32(v195+int32(0))+232)) = v190
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v216 + v205
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_CopyLoadRawBuf[2])) = v222 - v205
	goto L56
L59:
	;
	v228 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)) = uint8(v228)
	goto L61
L60:
	;
	goto L61
L61:
	;
	return
}
func F_CopyReadAttributesCSV(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v21 <= v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L7
	} else {
		goto L102
	}
L2:
	;
	m.G0 = v19 + int32(16)
	return v411
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	if v24 == int32(0) {
		v411 = v2
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v52 = l0 + int32(264)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v54
	goto L12
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesCSV_0), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesCSV_1), int32(1837), int32(_a_F_CopyReadAttributesCSV_2))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v61 <= v60 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_enlargeStringInfo(m, v52, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	v66 = v60
	goto L15
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	v68 = v66 + v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v76 = v69
	v77 = v2
	v78 = v67
	goto L18
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+292))
	v66 = v65
	goto L15
L17:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+268)) = v390 - v403
	v411 = v392
	goto L2
L18:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	if v89 <= v77 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+52))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L97
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v89 << (uint(int32(1)) % 32)
	v96 = F_repalloc(m, v88, v89<<(uint(int32(3))%32))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	v99 = v88
	goto L22
L22:
	;
	v101 = v77 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v99+v101))) = v76
	if base.Ui32(v78) < base.Ui32(v68) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v96
	v99 = v96
	goto L22
L24:
	;
	goto L19
L25:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v350+v101))) = int32(0)
	v355 = v77 + int32(1)
	if v206 != 0 {
		v76 = v197
		v77 = v355
		v78 = v198
		goto L18
	} else {
		goto L96
	}
L26:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v277 != 0 {
		goto L73
	} else {
		goto L74
	}
L27:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v212 == v203 {
		goto L54
	} else {
		goto L55
	}
L28:
	;
	v107 = v78
	v109 = v76
	v114 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v76))) = uint8(v190)
	v197 = v76 + int32(1)
	v198 = v78
	v199 = v76
	v203 = v190
	v206 = v190
	goto L27
L31:
	;
	v123 = v107 + int32(1)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	v125 = base.B2i32(v124 == v50&int32(255))
	if v125 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v184)
	v263 = v163
	v264 = v149
	v265 = v149 + int32(1)
	v268 = v163 - v78
	v271 = v184
	goto L26
L33:
	;
	v148 = v123
	v149 = v109
	goto L43
L34:
	;
	if v48 == v124 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v134 = v107
	v135 = v109
	goto L36
L36:
	;
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v137)
	v139 = v134 - v78
	v140 = int32(1)
	v141 = v135 + v140
	if v114&v140 == v137 {
		v197 = v141
		v198 = v123
		v199 = v135
		v203 = v139
		v206 = v125
		goto L27
	} else {
		goto L42
	}
L37:
	;
	if base.Ui32(v123) < base.Ui32(v68) {
		goto L33
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v124)
	v132 = v109 + int32(1)
	if base.Ui32(v123) < base.Ui32(v68) {
		v107 = v123
		v109 = v132
		goto L31
	} else {
		goto L41
	}
L40:
	;
	goto L1
L41:
	;
	v134 = v123
	v135 = v132
	goto L36
L42:
	;
	v263 = v123
	v264 = v135
	v265 = v141
	v268 = v139
	v271 = v125
	goto L26
L43:
	;
	v163 = v148 + int32(1)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v164 != v46 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if base.Ui32(v163) < base.Ui32(v68) {
		v107 = v163
		v109 = v149
		v114 = int32(1)
		goto L31
	} else {
		goto L53
	}
L45:
	;
	goto L44
L46:
	;
	if base.Ui32(v178) < base.Ui32(v68) {
		v148 = v178
		v149 = v149 + int32(1)
		goto L43
	} else {
		goto L52
	}
L47:
	;
	if v164 == v48 {
		goto L45
	} else {
		goto L51
	}
L48:
	;
	if base.Ui32(v68) <= base.Ui32(v163) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if base.B2i32(v46 != v167)&base.B2i32(v48 != v167) != 0 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v167)
	v178 = v148 + int32(2)
	goto L46
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v164)
	v178 = v163
	goto L46
L52:
	;
	goto L1
L53:
	;
	goto L32
L54:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v203 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	goto L56
L56:
	;
	v263 = v198
	v264 = v199
	v265 = v197
	v268 = v203
	v271 = v206
	goto L26
L57:
	;
	if v258 == int32(0) {
		goto L25
	} else {
		goto L71
	}
L58:
	;
	v258 = int32(0)
	goto L57
L59:
	;
	goto L60
L60:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v220 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v221 = v78
	v222 = v214
	v223 = v203
	v224 = v220
	goto L65
L62:
	;
	v246 = v214
	v250 = int32(0)
	goto L63
L63:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	v258 = v250 - v251
	goto L57
L64:
	;
	v246 = v241
	v250 = v243
	goto L63
L65:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	if v224 != v226 {
		v241 = v222
		v243 = v224
		goto L64
	} else {
		goto L67
	}
L66:
	;
	v241 = v235
	v243 = int32(0)
	goto L64
L67:
	;
	if v226 == int32(0) {
		v241 = v222
		v243 = v224
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v231 = v223 - int32(1)
	if v231 == int32(0) {
		v241 = v222
		v243 = v224
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v234 = int32(1)
	v235 = v222 + v234
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	if v236 != 0 {
		v221 = v221 + v234
		v222 = v235
		v223 = v231
		v224 = v236
		goto L65
	} else {
		goto L70
	}
L70:
	;
	goto L66
L71:
	;
	goto L56
L72:
	;
	v349 = v77 + int32(1)
	if v271 != 0 {
		v76 = v265
		v77 = v349
		v78 = v263
		goto L18
	} else {
		goto L95
	}
L73:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v280 = v278
	goto L75
L74:
	;
	v280 = int32(0)
	goto L75
L75:
	;
	if v280 <= v77 {
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v282 == int32(0) {
		goto L72
	} else {
		goto L77
	}
L77:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v268 != v285 {
		goto L72
	} else {
		goto L78
	}
L78:
	;
	if v268 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v330 != 0 {
		goto L72
	} else {
		goto L93
	}
L80:
	;
	v330 = int32(0)
	goto L79
L81:
	;
	goto L82
L82:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v292 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v293 = v78
	v294 = v282
	v295 = v268
	v296 = v292
	goto L87
L84:
	;
	v318 = v282
	v322 = int32(0)
	goto L85
L85:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	v330 = v322 - v323
	goto L79
L86:
	;
	v318 = v313
	v322 = v315
	goto L85
L87:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v296 != v298 {
		v313 = v294
		v315 = v296
		goto L86
	} else {
		goto L89
	}
L88:
	;
	v313 = v307
	v315 = int32(0)
	goto L86
L89:
	;
	if v298 == int32(0) {
		v313 = v294
		v315 = v296
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v303 = v295 - int32(1)
	if v303 == int32(0) {
		v313 = v294
		v315 = v296
		goto L86
	} else {
		goto L91
	}
L91:
	;
	v306 = int32(1)
	v307 = v294 + v306
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+1)))
	if v308 != 0 {
		v293 = v293 + v306
		v294 = v307
		v295 = v303
		v296 = v308
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v277)+12))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v332+v101)))
	v336 = v334 - int32(1)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v331+v336<<(uint(int32(2))%32))))
	if v340 == int32(0) {
		goto L24
	} else {
		goto L94
	}
L94:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v345 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v343+v336))) = uint8(v345)
	goto L72
L95:
	;
	v390 = v264
	v392 = v349
	goto L17
L96:
	;
	v390 = v199
	v392 = v355
	goto L17
L97:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesCSV_3), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v357 + v358<<(uint(int32(4))%32) + v336*int32(100) + int32(24)
	F_errdetail(m, int32(_a_F_CopyReadAttributesCSV_4), v19)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesCSV_1), int32(1990), int32(_a_F_CopyReadAttributesCSV_2))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesCSV_5), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesCSV_1), int32(1921), int32(_a_F_CopyReadAttributesCSV_2))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_collect_visibility_data(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v3
	v18 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = F_relation_open(m, l0, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L52
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
	v28 = v26 - int32(109)
	if base.Ui32(int32(7)) < base.Ui32(v28) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if int32(1)<<(uint(v28)%32)&int32(161) == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v38 = F_RelationGetNumberOfBlocksInFork(m, v23, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v42 = F_palloc0(m, v38+int32(8))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(0)
	if l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v38
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v48
	v55 = F_read_stream_begin_relation(m, int32(12), v18, v23, int32(120), v13+int32(4), v48)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v57 = v3
	goto L11
L11:
	;
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v57 = v55
	goto L11
L13:
	;
	v59 = v42 + int32(8)
	v61 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	if l1 != 0 {
		goto L43
	} else {
		goto L44
	}
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_collect_visibility_data[0]))
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v77 = F_visibilitymap_get_status(m, v23, v61, v13+int32(12))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	if v77&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v81 = v61 + v59
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v84 = v82 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v84)
	goto L25
L24:
	;
	goto L25
L25:
	;
	if v77&int32(2) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v89 = v61 + v59
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v92 = v90 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v92)
	goto L28
L27:
	;
	goto L28
L28:
	;
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v96 = F_read_stream_next_buffer(m, v57, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v133 = v61 + int32(1)
	if v133 != v38 {
		v61 = v133
		goto L16
	} else {
		goto L42
	}
L32:
	;
	F_LockBuffer(m, v96, int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v96 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+10)))
	if v119&int32(4) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_collect_visibility_data[1]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104+(v96^int32(-1))<<(uint(int32(2))%32))))
	v118 = v110
	goto L34
L36:
	;
	goto L37
L37:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_collect_visibility_data[2]))
	v118 = v112 + v96<<(uint(int32(13))%32) + int32(-8192)
	goto L34
L38:
	;
	v122 = v61 + v59
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v125 = v123 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v122))) = uint8(v125)
	goto L40
L39:
	;
	goto L40
L40:
	;
	F_UnlockReleaseBuffer(m, v96)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L31
L42:
	;
	goto L17
L43:
	;
	F_read_stream_end(m, v57)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v147 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_ReleaseBuffer(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_relation_close(m, v23, int32(1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	m.G0 = v13 + int32(16)
	return v42
L52:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v164 + int32(4)
	F_errmsg(m, int32(_a_F_collect_visibility_data_0), v13)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v174 = int32(*(*int8)(unsafe.Add(mBase, uint32(v173)+119)))
	F_errdetail_relkind_not_supported(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_collect_visibility_data_1), int32(951), int32(_a_F_collect_visibility_data_2))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_compare3(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	v6 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v8) < base.Ui32(v7) {
		v24 = v6
	} else {
		v10 = base.B2i32(v7 != v8)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if base.B2i32(v10 == int32(0))&base.B2i32(base.Ui32(v14) < base.Ui32(v13)) != 0 {
			v24 = v6
		} else {
			v24 = int32(0) - (v10 | base.B2i32(v13 != v14))
		}
	}
	return v24
}
func F_compare_fractional_path_costs(m *base.Module, l0 int32, l1 int32, l2 float64) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v45 float64
	_ = v45
	var v51 int32
	_ = v51
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v8 != v9 {
		if v8 < v9 {
			v14 = int32(-1)
		} else {
			v14 = int32(1)
		}
		return v14
	} else {
		if base.F64_le(l2, float64(0))|base.F64_ge(l2, float64(1)) != 0 {
			v21 = int32(-1)
			v22 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
			v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			if base.F64_lt(v22, v23) != 0 {
				v51 = v21
				return v51
			} else {
				if base.F64_gt(v22, v23) != 0 {
					return int32(1)
				} else {
					v28 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
					v29 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
					if base.F64_lt(v28, v29) != 0 {
						v51 = v21
						return v51
					} else {
						if base.F64_gt(v28, v29) != 0 {
							v51 = int32(1)
							return v51
						} else {
							return int32(0)
						}
					}
				}
			}
		} else {
			v36 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
			v37 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
			v40 = base.F64_add(base.F64_mul(l2, base.F64_sub(v36, v37)), v37)
			v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			v42 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
			v45 = base.F64_add(base.F64_mul(l2, base.F64_sub(v41, v42)), v42)
			if base.F64_lt(v40, v45) != 0 {
				v51 = int32(-1)
			} else {
				v51 = base.F64_lt(v45, v40)
			}
			return v51
		}
	}
}
func F_compare_subnode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	if l2 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	v15 = l1 + l2
	v17 = l0 + int32(2)
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v19 = v17 + v18
	v21 = l1
	goto L3
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v32 != int32(95) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L1
L5:
	;
	if v18 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v49 = v21
	goto L14
L7:
	;
	if base.Ui32(v21) < base.Ui32(v15) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v36 = F_pg_mblen_range(m, v21, v15)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v68 = v21
	goto L5
L11:
	;
	return int32(0)
L12:
	;
	v40 = v36 + v21
	if base.Ui32(v40) < base.Ui32(v15) {
		v21 = v40
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	v54 = F_pg_mblen_range(m, v49, v15)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v68 = v56
	goto L5
L16:
	;
	v56 = v54 + v49
	if base.Ui32(v15) <= base.Ui32(v56) {
		v68 = v56
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v58 != int32(95) {
		v49 = v56
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	return int32(0)
L20:
	;
	goto L21
L21:
	;
	v77 = v68 - v21
	v83 = v17
	goto L23
L22:
	;
	if base.Ui32(v68) < base.Ui32(v15) {
		v21 = v21 + v77
		goto L3
	} else {
		goto L46
	}
L23:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v90 != int32(95) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	return int32(0)
L25:
	;
	v135 = int32(0)
	if base.B2i32(l4&base.B2i32(v77 < v129) == v135)&base.B2i32(v129 != v77) == v135 {
		goto L40
	} else {
		goto L41
	}
L26:
	;
	v103 = v83
	goto L33
L27:
	;
	if base.Ui32(v83) < base.Ui32(v19) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v95 = F_pg_mblen_range(m, v83, v19)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L11
	} else {
		goto L31
	}
L30:
	;
	v129 = int32(0)
	goto L25
L31:
	;
	v97 = v95 + v83
	if base.Ui32(v97) < base.Ui32(v19) {
		v83 = v97
		goto L23
	} else {
		goto L32
	}
L32:
	;
	return int32(0)
L33:
	;
	v113 = F_pg_mblen_range(m, v103, v19)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L11
	} else {
		goto L35
	}
L34:
	;
	v129 = v115 - v83
	goto L25
L35:
	;
	v115 = v113 + v103
	if base.Ui32(v115) < base.Ui32(v19) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v117 != int32(95) {
		v103 = v115
		goto L33
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L34
L39:
	;
	goto L38
L40:
	;
	v141 = m.T0[l3].(func(*base.Module, int32, int32, int32, int32) int32)(m, v21, v77, v83, v129)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L11
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v143 = v83 + v129
	if base.Ui32(v143) < base.Ui32(v19) {
		v83 = v143
		goto L23
	} else {
		goto L45
	}
L43:
	;
	if v141 != 0 {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	goto L24
L46:
	;
	goto L4
}
func F_comparecost_1(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	return base.B2i32(v4 < v3) - base.B2i32(v3 < v4)
}
func F_compute_new_xmax_infomask(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v383 int32
	_ = v383
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v746 int32
	_ = v746
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v903 int32
	_ = v903
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1044 int32
	_ = v1044
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	v10 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	if l1&int32(2048) != 0 {
		v1027 = l4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l7))) = uint16(v1068)
	*(*uint16)(unsafe.Add(mBase, uint32(l8))) = uint16(v1064)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1066
	m.G0 = v19 + int32(80)
	return
L2:
	;
	if l5 != 0 {
		goto L321
	} else {
		goto L322
	}
L3:
	;
	if l1&int32(_a_F_compute_new_xmax_infomask_0) != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L12
	} else {
		goto L315
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L12
	} else {
		goto L309
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L12
	} else {
		goto L303
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L12
	} else {
		goto L297
	}
L8:
	;
	if l1&int32(_a_F_compute_new_xmax_infomask_1) == int32(_a_F_compute_new_xmax_infomask_2) {
		v1027 = l4
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v432 = l2 & int32(_a_F_compute_new_xmax_infomask_3)
	if v432 != 0 {
		goto L115
	} else {
		goto L116
	}
L11:
	;
	v37 = int32(base.Ui32(l1&int32(128))>>(uint(int32(7))%32)) | base.B2i32(l1&int32(_a_F_compute_new_xmax_infomask_4) == int32(64))
	v38 = F_MultiXactIdIsRunning(m, l0, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	if v38 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v37 != 0 {
		v1027 = l4
		goto L2
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if l5 != 0 {
		goto L31
	} else {
		goto L32
	}
L17:
	;
	v42 = int32(0)
	v46 = F_GetMultiXactIdMembers(m, l0, v19+int32(76), v42)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	if int32(0) < v46 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	v52 = v42
	goto L24
L20:
	;
	v83 = v42
	goto L21
L21:
	;
	v98 = F_TransactionIdDidCommit(m, v83)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L12
	} else {
		goto L29
	}
L22:
	;
	F_pfree(m, v50)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L12
	} else {
		goto L28
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v79 = v77
	goto L22
L24:
	;
	v69 = v50 + v52<<(uint(int32(3))%32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v70) {
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v79 = int32(0)
	goto L22
L26:
	;
	v74 = v52 + int32(1)
	if v74 != v46 {
		v52 = v74
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v83 = v79
	goto L21
L29:
	;
	if v98 == int32(0) {
		v1027 = l4
		goto L2
	} else {
		goto L30
	}
L30:
	;
	goto L16
L31:
	;
	v122 = int32(8)
	goto L33
L32:
	;
	v122 = int32(4)
	goto L33
L33:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v122)+uint32(_c_F_compute_new_xmax_infomask[0])))
	if v126 == int32(-1) {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v129 = int32(0)
	v131 = m.G0
	v133 = v131 - int32(16)
	m.G0 = v133
	v138 = F_GetMultiXactIdMembers(m, l0, v133+int32(12), v129)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L12
	} else {
		goto L38
	}
L35:
	;
	m.G0 = v133 + int32(16)
	v293 = F_GetMultiXactIdMembers(m, v271, v19+int32(76), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L12
	} else {
		goto L70
	}
L36:
	;
	v260 = v252 + v242<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v260)+4)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = l3
	v265 = F_MultiXactIdCreateFromMembers(m, v242+int32(1), v252)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L12
	} else {
		goto L67
	}
L37:
	;
	v240 = F_palloc(m, int32(8))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L12
	} else {
		goto L66
	}
L38:
	;
	if int32(0) <= v138 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	if v138 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+8)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = l3
	v236 = F_MultiXactIdCreateFromMembers(m, int32(1), v133+int32(4))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L12
	} else {
		goto L65
	}
L42:
	;
	v149 = v129
	goto L43
L43:
	;
	v163 = v142 + v149<<(uint(int32(3))%32)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v164 != l3 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v173 = int32(1)
	if v138 <= v173 {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	v171 = v149 + int32(1)
	if v171 != v138 {
		v149 = v171
		goto L43
	} else {
		goto L49
	}
L46:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v166 != v126 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	F_pfree(m, v142)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	v271 = l0
	goto L35
L49:
	;
	goto L44
L50:
	;
	v176 = v173
	goto L52
L51:
	;
	v176 = v138
	goto L52
L52:
	;
	v182 = F_palloc(m, v138<<(uint(int32(3))%32)+int32(8))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	v185 = int32(0)
	v189 = int32(0)
	goto L54
L54:
	;
	v203 = v142 + v189<<(uint(int32(3))%32)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v205 = F_TransactionIdIsInProgress(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L12
	} else {
		goto L57
	}
L55:
	;
	v242 = v226
	v252 = v182
	goto L36
L56:
	;
	v229 = v189 + int32(1)
	if v229 != v176 {
		v185 = v226
		v189 = v229
		goto L54
	} else {
		goto L64
	}
L57:
	;
	if v205 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	if base.Ui32(v209) < base.Ui32(int32(4)) {
		v226 = v185
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v219 = v182 + v185<<(uint(int32(3))%32)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v219)+4)) = v222
	v226 = v185 + int32(1)
	goto L56
L61:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v213 = F_TransactionIdDidCommit(m, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	if v213 == int32(0) {
		v226 = v185
		goto L56
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	goto L55
L65:
	;
	v271 = v236
	goto L35
L66:
	;
	v242 = int32(0)
	v252 = v240
	goto L36
L67:
	;
	F_pfree(m, v142)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v252)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	v271 = v265
	goto L35
L70:
	;
	if v293 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v1064 = v129
	v1066 = v271
	v1068 = int32(_a_F_compute_new_xmax_infomask_5)
	goto L1
L72:
	;
	goto L73
L73:
	;
	v298 = int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v293 == v298 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v293&v298 == int32(0) {
		v408 = v371
		v410 = v375
		v411 = v383
		goto L95
	} else {
		goto L96
	}
L75:
	;
	v303 = int32(0)
	v371 = v129
	v374 = v303
	v375 = v303
	v383 = v10
	goto L74
L76:
	;
	goto L77
L77:
	;
	v307 = int32(0)
	v310 = v307
	v311 = v129
	v314 = v307
	v315 = v307
	v323 = v10
	goto L78
L78:
	;
	v328 = v300 + v314<<(uint(int32(3))%32)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v329<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v315) < base.Ui32(v334) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v371 = v362
	v374 = v366
	v375 = v364
	v383 = v363
	goto L74
L80:
	;
	v336 = v334
	goto L82
L81:
	;
	v336 = v315
	goto L82
L82:
	;
	switch v329 - int32(3) {
	case 0:
		goto L86
	case 1:
		v343 = v311
		goto L84
	case 2:
		goto L85
	default:
		v345 = v311
		v346 = v323
		goto L83
	}
L83:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v347<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	switch v347 - int32(3) {
	case 0:
		goto L90
	case 1:
		v360 = v345
		goto L88
	case 2:
		goto L89
	default:
		v362 = v345
		v363 = v346
		goto L87
	}
L84:
	;
	v345 = v343
	v346 = int32(1)
	goto L83
L85:
	;
	v343 = v311 | int32(_a_F_compute_new_xmax_infomask_3)
	goto L84
L86:
	;
	v345 = v311 | int32(_a_F_compute_new_xmax_infomask_3)
	v346 = v323
	goto L83
L87:
	;
	if base.Ui32(v336) < base.Ui32(v352) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v362 = v360
	v363 = int32(1)
	goto L87
L89:
	;
	v360 = v345 | int32(_a_F_compute_new_xmax_infomask_3)
	goto L88
L90:
	;
	v362 = v345 | int32(_a_F_compute_new_xmax_infomask_3)
	v363 = v346
	goto L87
L91:
	;
	v364 = v352
	goto L93
L92:
	;
	v364 = v336
	goto L93
L93:
	;
	v365 = int32(2)
	v366 = v314 + v365
	v368 = v310 + v365
	if v368 != v293&int32(2147483646) {
		v310 = v368
		v311 = v362
		v314 = v366
		v315 = v364
		v323 = v363
		goto L78
	} else {
		goto L94
	}
L94:
	;
	goto L79
L95:
	;
	F_pfree(m, v300)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L12
	} else {
		goto L103
	}
L96:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v300+v374<<(uint(int32(3))%32))+4))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v391<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v375) < base.Ui32(v396) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v398 = v396
	goto L99
L98:
	;
	v398 = v375
	goto L99
L99:
	;
	switch v391 - int32(3) {
	case 0:
		goto L102
	case 1:
		v405 = v371
		goto L100
	case 2:
		goto L101
	default:
		v408 = v371
		v410 = v398
		v411 = v383
		goto L95
	}
L100:
	;
	v408 = v405
	v410 = v398
	v411 = int32(1)
	goto L95
L101:
	;
	v405 = v371 | int32(_a_F_compute_new_xmax_infomask_3)
	goto L100
L102:
	;
	v408 = v371 | int32(_a_F_compute_new_xmax_infomask_3)
	v410 = v398
	v411 = v383
	goto L95
L103:
	;
	if v410&int32(-2) == int32(2) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if v411 != 0 {
		v1064 = v408
		v1066 = v271
		v1068 = int32(_a_F_compute_new_xmax_infomask_6)
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v410 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v1064 = v408
	v1066 = v271
	v1068 = int32(_a_F_compute_new_xmax_infomask_7)
	goto L1
L108:
	;
	v423 = int32(_a_F_compute_new_xmax_infomask_0)
	goto L110
L109:
	;
	v423 = int32(_a_F_compute_new_xmax_infomask_8)
	goto L110
L110:
	;
	if v410 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v426 = int32(_a_F_compute_new_xmax_infomask_4)
	goto L113
L112:
	;
	v426 = v423
	goto L113
L113:
	;
	if v411 != 0 {
		v1064 = v408
		v1066 = v271
		v1068 = v426
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v1064 = v408
	v1066 = v271
	v1068 = v426 | int32(128)
	goto L1
L115:
	;
	v433 = int32(5)
	goto L117
L116:
	;
	v433 = int32(4)
	goto L117
L117:
	;
	if l1&int32(1024) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	if l5 != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v597 = int32(base.Ui32(l1&int32(128))>>(uint(int32(7))%32)) | base.B2i32(l1&int32(80) == int32(64))
	v598 = F_TransactionIdIsInProgress(m, l0)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L12
	} else {
		goto L171
	}
L121:
	;
	v440 = int32(8)
	goto L123
L122:
	;
	v440 = int32(4)
	goto L123
L123:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v440)+uint32(_c_F_compute_new_xmax_infomask[0])))
	if v444 == int32(-1) {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	v447 = int32(0)
	v448 = F_MultiXactIdCreate(m, l0, v433, l3, v444)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L12
	} else {
		goto L125
	}
L125:
	;
	v453 = F_GetMultiXactIdMembers(m, v448, v19+int32(76), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L12
	} else {
		goto L126
	}
L126:
	;
	if v453 <= int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v1064 = v447
	v1066 = v448
	v1068 = int32(_a_F_compute_new_xmax_infomask_5)
	goto L1
L128:
	;
	goto L129
L129:
	;
	v458 = int32(1)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v453 == v458 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v453&v458 == int32(0) {
		v568 = v531
		v570 = v535
		v571 = v543
		goto L151
	} else {
		goto L152
	}
L131:
	;
	v463 = int32(0)
	v531 = v447
	v534 = v463
	v535 = v463
	v543 = v10
	goto L130
L132:
	;
	goto L133
L133:
	;
	v467 = int32(0)
	v470 = v467
	v471 = v447
	v474 = v467
	v475 = v467
	v483 = v10
	goto L134
L134:
	;
	v488 = v460 + v474<<(uint(int32(3))%32)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v489<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v475) < base.Ui32(v494) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v531 = v522
	v534 = v526
	v535 = v524
	v543 = v523
	goto L130
L136:
	;
	v496 = v494
	goto L138
L137:
	;
	v496 = v475
	goto L138
L138:
	;
	switch v489 - int32(3) {
	case 0:
		goto L142
	case 1:
		v503 = v471
		goto L140
	case 2:
		goto L141
	default:
		v505 = v471
		v506 = v483
		goto L139
	}
L139:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v507<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	switch v507 - int32(3) {
	case 0:
		goto L146
	case 1:
		v520 = v505
		goto L144
	case 2:
		goto L145
	default:
		v522 = v505
		v523 = v506
		goto L143
	}
L140:
	;
	v505 = v503
	v506 = int32(1)
	goto L139
L141:
	;
	v503 = v471 | int32(_a_F_compute_new_xmax_infomask_3)
	goto L140
L142:
	;
	v505 = v471 | int32(_a_F_compute_new_xmax_infomask_3)
	v506 = v483
	goto L139
L143:
	;
	if base.Ui32(v496) < base.Ui32(v512) {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	v522 = v520
	v523 = int32(1)
	goto L143
L145:
	;
	v520 = v505 | int32(_a_F_compute_new_xmax_infomask_3)
	goto L144
L146:
	;
	v522 = v505 | int32(_a_F_compute_new_xmax_infomask_3)
	v523 = v506
	goto L143
L147:
	;
	v524 = v512
	goto L149
L148:
	;
	v524 = v496
	goto L149
L149:
	;
	v525 = int32(2)
	v526 = v474 + v525
	v528 = v470 + v525
	if v528 != v453&int32(2147483646) {
		v470 = v528
		v471 = v522
		v474 = v526
		v475 = v524
		v483 = v523
		goto L134
	} else {
		goto L150
	}
L150:
	;
	goto L135
L151:
	;
	F_pfree(m, v460)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L12
	} else {
		goto L159
	}
L152:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v460+v534<<(uint(int32(3))%32))+4))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v551<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v535) < base.Ui32(v556) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v558 = v556
	goto L155
L154:
	;
	v558 = v535
	goto L155
L155:
	;
	switch v551 - int32(3) {
	case 0:
		goto L158
	case 1:
		v565 = v531
		goto L156
	case 2:
		goto L157
	default:
		v568 = v531
		v570 = v558
		v571 = v543
		goto L151
	}
L156:
	;
	v568 = v565
	v570 = v558
	v571 = int32(1)
	goto L151
L157:
	;
	v565 = v531 | int32(_a_F_compute_new_xmax_infomask_3)
	goto L156
L158:
	;
	v568 = v531 | int32(_a_F_compute_new_xmax_infomask_3)
	v570 = v558
	v571 = v543
	goto L151
L159:
	;
	if v570&int32(-2) == int32(2) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	if v571 != 0 {
		v1064 = v568
		v1066 = v448
		v1068 = int32(_a_F_compute_new_xmax_infomask_6)
		goto L1
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if v570 != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v1064 = v568
	v1066 = v448
	v1068 = int32(_a_F_compute_new_xmax_infomask_7)
	goto L1
L164:
	;
	v583 = int32(_a_F_compute_new_xmax_infomask_0)
	goto L166
L165:
	;
	v583 = int32(_a_F_compute_new_xmax_infomask_8)
	goto L166
L166:
	;
	if v570 == int32(1) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v586 = int32(_a_F_compute_new_xmax_infomask_4)
	goto L169
L168:
	;
	v586 = v583
	goto L169
L169:
	;
	if v571 != 0 {
		v1064 = v568
		v1066 = v448
		v1068 = v586
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v1064 = v568
	v1066 = v448
	v1068 = v586 | int32(128)
	goto L1
L171:
	;
	if v598 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	if v597 == int32(0) {
		v630 = v433
		goto L175
	} else {
		goto L176
	}
L173:
	;
	goto L174
L174:
	;
	if v597 != 0 {
		v1027 = l4
		goto L2
	} else {
		goto L244
	}
L175:
	;
	if l0 == l3 {
		goto L188
	} else {
		goto L189
	}
L176:
	;
	switch int32(base.Ui32(l1)>>(uint(int32(4))%32))&int32(5) - int32(1) {
	case 0:
		v630 = int32(0)
		goto L175
	case 1, 2:
		goto L179
	case 3:
		goto L180
	case 4:
		goto L177
	default:
		goto L178
	}
L177:
	;
	v630 = int32(1)
	goto L175
L178:
	;
	v614 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L12
	} else {
		goto L184
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	if v432 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v611 = int32(3)
	goto L183
L182:
	;
	v611 = int32(2)
	goto L183
L183:
	;
	v630 = v611
	goto L175
L184:
	;
	if v614 == int32(0) {
		v1027 = l4
		goto L2
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_9), v19+int32(16))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L12
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_compute_new_xmax_infomask_10), int32(_a_F_compute_new_xmax_infomask_11), int32(_a_F_compute_new_xmax_infomask_12))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L12
	} else {
		goto L187
	}
L187:
	;
	v1027 = l4
	goto L2
L188:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v630<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v636) < base.Ui32(l4) {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	goto L190
L190:
	;
	if l5 != 0 {
		goto L194
	} else {
		goto L195
	}
L191:
	;
	v638 = l4
	goto L193
L192:
	;
	v638 = v636
	goto L193
L193:
	;
	v1027 = v638
	goto L2
L194:
	;
	v643 = int32(8)
	goto L196
L195:
	;
	v643 = int32(4)
	goto L196
L196:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v643)+uint32(_c_F_compute_new_xmax_infomask[0])))
	if v647 == int32(-1) {
		goto L5
	} else {
		goto L197
	}
L197:
	;
	v650 = int32(0)
	v651 = F_MultiXactIdCreate(m, l0, v630, l3, v647)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L12
	} else {
		goto L198
	}
L198:
	;
	v656 = F_GetMultiXactIdMembers(m, v651, v19+int32(76), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L12
	} else {
		goto L199
	}
L199:
	;
	if v656 <= int32(0) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1064 = v650
	v1066 = v651
	v1068 = int32(_a_F_compute_new_xmax_infomask_5)
	goto L1
L201:
	;
	goto L202
L202:
	;
	v661 = int32(1)
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v656 == v661 {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	if v656&v661 == int32(0) {
		v771 = v734
		v773 = v738
		v774 = v746
		goto L224
	} else {
		goto L225
	}
L204:
	;
	v666 = int32(0)
	v734 = v650
	v737 = v666
	v738 = v666
	v746 = v10
	goto L203
L205:
	;
	goto L206
L206:
	;
	v670 = int32(0)
	v673 = v670
	v674 = v650
	v677 = v670
	v678 = v670
	v686 = v10
	goto L207
L207:
	;
	v691 = v663 + v677<<(uint(int32(3))%32)
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v692<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v678) < base.Ui32(v697) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v734 = v725
	v737 = v729
	v738 = v727
	v746 = v726
	goto L203
L209:
	;
	v699 = v697
	goto L211
L210:
	;
	v699 = v678
	goto L211
L211:
	;
	switch v692 - int32(3) {
	case 0:
		goto L215
	case 1:
		v706 = v674
		goto L213
	case 2:
		goto L214
	default:
		v708 = v674
		v709 = v686
		goto L212
	}
L212:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v691)+12))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v710<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	switch v710 - int32(3) {
	case 0:
		goto L219
	case 1:
		v723 = v708
		goto L217
	case 2:
		goto L218
	default:
		v725 = v708
		v726 = v709
		goto L216
	}
L213:
	;
	v708 = v706
	v709 = int32(1)
	goto L212
L214:
	;
	v706 = v674 | int32(_a_F_compute_new_xmax_infomask_3)
	goto L213
L215:
	;
	v708 = v674 | int32(_a_F_compute_new_xmax_infomask_3)
	v709 = v686
	goto L212
L216:
	;
	if base.Ui32(v699) < base.Ui32(v715) {
		goto L220
	} else {
		goto L221
	}
L217:
	;
	v725 = v723
	v726 = int32(1)
	goto L216
L218:
	;
	v723 = v708 | int32(_a_F_compute_new_xmax_infomask_3)
	goto L217
L219:
	;
	v725 = v708 | int32(_a_F_compute_new_xmax_infomask_3)
	v726 = v709
	goto L216
L220:
	;
	v727 = v715
	goto L222
L221:
	;
	v727 = v699
	goto L222
L222:
	;
	v728 = int32(2)
	v729 = v677 + v728
	v731 = v673 + v728
	if v731 != v656&int32(2147483646) {
		v673 = v731
		v674 = v725
		v677 = v729
		v678 = v727
		v686 = v726
		goto L207
	} else {
		goto L223
	}
L223:
	;
	goto L208
L224:
	;
	F_pfree(m, v663)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L12
	} else {
		goto L232
	}
L225:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v663+v737<<(uint(int32(3))%32))+4))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v754<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v738) < base.Ui32(v759) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v761 = v759
	goto L228
L227:
	;
	v761 = v738
	goto L228
L228:
	;
	switch v754 - int32(3) {
	case 0:
		goto L231
	case 1:
		v768 = v734
		goto L229
	case 2:
		goto L230
	default:
		v771 = v734
		v773 = v761
		v774 = v746
		goto L224
	}
L229:
	;
	v771 = v768
	v773 = v761
	v774 = int32(1)
	goto L224
L230:
	;
	v768 = v734 | int32(_a_F_compute_new_xmax_infomask_3)
	goto L229
L231:
	;
	v771 = v734 | int32(_a_F_compute_new_xmax_infomask_3)
	v773 = v761
	v774 = v746
	goto L224
L232:
	;
	if v773&int32(-2) == int32(2) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if v774 != 0 {
		v1064 = v771
		v1066 = v651
		v1068 = int32(_a_F_compute_new_xmax_infomask_6)
		goto L1
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	if v773 != 0 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v1064 = v771
	v1066 = v651
	v1068 = int32(_a_F_compute_new_xmax_infomask_7)
	goto L1
L237:
	;
	v786 = int32(_a_F_compute_new_xmax_infomask_0)
	goto L239
L238:
	;
	v786 = int32(_a_F_compute_new_xmax_infomask_8)
	goto L239
L239:
	;
	if v773 == int32(1) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v789 = int32(_a_F_compute_new_xmax_infomask_4)
	goto L242
L241:
	;
	v789 = v786
	goto L242
L242:
	;
	if v774 != 0 {
		v1064 = v771
		v1066 = v651
		v1068 = v789
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v1064 = v771
	v1066 = v651
	v1068 = v789 | int32(128)
	goto L1
L244:
	;
	v792 = F_TransactionIdDidCommit(m, l0)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L12
	} else {
		goto L245
	}
L245:
	;
	if v792 == int32(0) {
		v1027 = l4
		goto L2
	} else {
		goto L246
	}
L246:
	;
	if l5 != 0 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v800 = int32(8)
	goto L249
L248:
	;
	v800 = int32(4)
	goto L249
L249:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v800)+uint32(_c_F_compute_new_xmax_infomask[0])))
	if v804 == int32(-1) {
		goto L4
	} else {
		goto L250
	}
L250:
	;
	v807 = int32(0)
	v808 = F_MultiXactIdCreate(m, l0, v433, l3, v804)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L12
	} else {
		goto L251
	}
L251:
	;
	v813 = F_GetMultiXactIdMembers(m, v808, v19+int32(76), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L12
	} else {
		goto L252
	}
L252:
	;
	if v813 <= int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1064 = v807
	v1066 = v808
	v1068 = int32(_a_F_compute_new_xmax_infomask_5)
	goto L1
L254:
	;
	goto L255
L255:
	;
	v818 = int32(1)
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v813 == v818 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v813&v818 == int32(0) {
		v928 = v891
		v930 = v895
		v931 = v903
		goto L277
	} else {
		goto L278
	}
L257:
	;
	v823 = int32(0)
	v891 = v807
	v894 = v823
	v895 = v823
	v903 = v10
	goto L256
L258:
	;
	goto L259
L259:
	;
	v827 = int32(0)
	v830 = v827
	v831 = v807
	v834 = v827
	v835 = v827
	v843 = v10
	goto L260
L260:
	;
	v848 = v820 + v834<<(uint(int32(3))%32)
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v848)+4))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v849<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v835) < base.Ui32(v854) {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	v891 = v882
	v894 = v886
	v895 = v884
	v903 = v883
	goto L256
L262:
	;
	v856 = v854
	goto L264
L263:
	;
	v856 = v835
	goto L264
L264:
	;
	switch v849 - int32(3) {
	case 0:
		goto L268
	case 1:
		v863 = v831
		goto L266
	case 2:
		goto L267
	default:
		v865 = v831
		v866 = v843
		goto L265
	}
L265:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v848)+12))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v867<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	switch v867 - int32(3) {
	case 0:
		goto L272
	case 1:
		v880 = v865
		goto L270
	case 2:
		goto L271
	default:
		v882 = v865
		v883 = v866
		goto L269
	}
L266:
	;
	v865 = v863
	v866 = int32(1)
	goto L265
L267:
	;
	v863 = v831 | int32(_a_F_compute_new_xmax_infomask_3)
	goto L266
L268:
	;
	v865 = v831 | int32(_a_F_compute_new_xmax_infomask_3)
	v866 = v843
	goto L265
L269:
	;
	if base.Ui32(v856) < base.Ui32(v872) {
		goto L273
	} else {
		goto L274
	}
L270:
	;
	v882 = v880
	v883 = int32(1)
	goto L269
L271:
	;
	v880 = v865 | int32(_a_F_compute_new_xmax_infomask_3)
	goto L270
L272:
	;
	v882 = v865 | int32(_a_F_compute_new_xmax_infomask_3)
	v883 = v866
	goto L269
L273:
	;
	v884 = v872
	goto L275
L274:
	;
	v884 = v856
	goto L275
L275:
	;
	v885 = int32(2)
	v886 = v834 + v885
	v888 = v830 + v885
	if v888 != v813&int32(2147483646) {
		v830 = v888
		v831 = v882
		v834 = v886
		v835 = v884
		v843 = v883
		goto L260
	} else {
		goto L276
	}
L276:
	;
	goto L261
L277:
	;
	F_pfree(m, v820)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L12
	} else {
		goto L285
	}
L278:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v820+v894<<(uint(int32(3))%32))+4))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v911<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v895) < base.Ui32(v916) {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v918 = v916
	goto L281
L280:
	;
	v918 = v895
	goto L281
L281:
	;
	switch v911 - int32(3) {
	case 0:
		goto L284
	case 1:
		v925 = v891
		goto L282
	case 2:
		goto L283
	default:
		v928 = v891
		v930 = v918
		v931 = v903
		goto L277
	}
L282:
	;
	v928 = v925
	v930 = v918
	v931 = int32(1)
	goto L277
L283:
	;
	v925 = v891 | int32(_a_F_compute_new_xmax_infomask_3)
	goto L282
L284:
	;
	v928 = v891 | int32(_a_F_compute_new_xmax_infomask_3)
	v930 = v918
	v931 = v903
	goto L277
L285:
	;
	if v930&int32(-2) == int32(2) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	if v931 != 0 {
		v1064 = v928
		v1066 = v808
		v1068 = int32(_a_F_compute_new_xmax_infomask_6)
		goto L1
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	if v930 != 0 {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	v1064 = v928
	v1066 = v808
	v1068 = int32(_a_F_compute_new_xmax_infomask_7)
	goto L1
L290:
	;
	v943 = int32(_a_F_compute_new_xmax_infomask_0)
	goto L292
L291:
	;
	v943 = int32(_a_F_compute_new_xmax_infomask_8)
	goto L292
L292:
	;
	if v930 == int32(1) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v946 = int32(_a_F_compute_new_xmax_infomask_4)
	goto L295
L294:
	;
	v946 = v943
	goto L295
L295:
	;
	if v931 != 0 {
		v1064 = v928
		v1066 = v808
		v1068 = v946
		goto L1
	} else {
		goto L296
	}
L296:
	;
	v1064 = v928
	v1066 = v808
	v1068 = v946 | int32(128)
	goto L1
L297:
	;
	if l5 != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v955 = int32(_a_F_compute_new_xmax_infomask_13)
	goto L300
L299:
	;
	v955 = int32(_a_F_compute_new_xmax_infomask_14)
	goto L300
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v955
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = l4
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_15), v19-int32(-64))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L12
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(_a_F_compute_new_xmax_infomask_10), int32(_a_F_compute_new_xmax_infomask_16), int32(_a_F_compute_new_xmax_infomask_17))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L12
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L303:
	;
	if l5 != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v974 = int32(_a_F_compute_new_xmax_infomask_13)
	goto L306
L305:
	;
	v974 = int32(_a_F_compute_new_xmax_infomask_14)
	goto L306
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v974
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = l4
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_15), v19+int32(48))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L12
	} else {
		goto L307
	}
L307:
	;
	F_errfinish(m, int32(_a_F_compute_new_xmax_infomask_10), int32(_a_F_compute_new_xmax_infomask_16), int32(_a_F_compute_new_xmax_infomask_17))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L12
	} else {
		goto L308
	}
L308:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L309:
	;
	if l5 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v993 = int32(_a_F_compute_new_xmax_infomask_13)
	goto L312
L311:
	;
	v993 = int32(_a_F_compute_new_xmax_infomask_14)
	goto L312
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l4
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_15), v19)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L12
	} else {
		goto L313
	}
L313:
	;
	F_errfinish(m, int32(_a_F_compute_new_xmax_infomask_10), int32(_a_F_compute_new_xmax_infomask_16), int32(_a_F_compute_new_xmax_infomask_17))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L12
	} else {
		goto L314
	}
L314:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L315:
	;
	if l5 != 0 {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1010 = int32(_a_F_compute_new_xmax_infomask_13)
	goto L318
L317:
	;
	v1010 = int32(_a_F_compute_new_xmax_infomask_14)
	goto L318
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v1010
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l4
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_15), v19+int32(32))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L12
	} else {
		goto L319
	}
L319:
	;
	F_errfinish(m, int32(_a_F_compute_new_xmax_infomask_10), int32(_a_F_compute_new_xmax_infomask_16), int32(_a_F_compute_new_xmax_infomask_17))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L12
	} else {
		goto L320
	}
L320:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L321:
	;
	v1064 = base.B2i32(v1027 == int32(3)) << (uint(int32(13)) % 32)
	v1066 = l3
	v1068 = int32(0)
	goto L1
L322:
	;
	goto L323
L323:
	;
	v1044 = int32(0)
	switch v1027 {
	case 0:
		v1064 = v1044
		v1066 = l3
		v1068 = int32(144)
		goto L1
	case 1:
		goto L327
	case 2:
		goto L326
	case 3:
		goto L325
	default:
		goto L324
	}
L324:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L12
	} else {
		goto L328
	}
L325:
	;
	v1064 = int32(_a_F_compute_new_xmax_infomask_3)
	v1066 = l3
	v1068 = int32(192)
	goto L1
L326:
	;
	v1064 = v1044
	v1066 = l3
	v1068 = int32(192)
	goto L1
L327:
	;
	v1064 = v1044
	v1066 = l3
	v1068 = int32(208)
	goto L1
L328:
	;
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_18), int32(0))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L12
	} else {
		goto L329
	}
L329:
	;
	F_errfinish(m, int32(_a_F_compute_new_xmax_infomask_10), int32(_a_F_compute_new_xmax_infomask_19), int32(_a_F_compute_new_xmax_infomask_12))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L12
	} else {
		goto L330
	}
L330:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_concat_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v14 == v4 {
		v26 = v4
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L12
	} else {
		goto L51
	}
L2:
	;
	m.G0 = v12 + int32(32)
	return v181
L3:
	;
	if v26&int32(1) != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L3
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v18 == int32(0) {
		v26 = v4
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v21 != int32(15) {
		v26 = v4
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+13)))
	v26 = v24
	goto L4
L8:
	;
	v31 = l2 + l1<<(uint(int32(3))%32)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+24)))
	if v32 != 0 {
		v181 = v4
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_initStringInfo(m, v12+int32(8))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L15
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	v34 = F_pg_detoast_datum(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v39 = F_array_to_text_internal(m, l2, v34, l0, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v181 = v39
	goto L2
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v46 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+18)))
	v53 = F_MemoryContextAlloc(m, v49, v50*int32(28))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	v106 = v46
	goto L18
L18:
	;
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+18)))
	if l1 < v109 {
		goto L30
	} else {
		goto L31
	}
L19:
	;
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+18)))
	if l1 < v55 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v60 = l1
	goto L23
L21:
	;
	goto L22
L22:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+16)) = v53
	v106 = v53
	goto L18
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v67 = F_get_fn_expr_argtype(m, v66, v60)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	if v67 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_getTypeOutputInfo(m, v67, v12+int32(28), v12+int32(27))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	F_fmgr_info_cxt(m, v77, v53+v60*int32(28), v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v86 = v60 + int32(1)
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+18)))
	if v86 < v87 {
		v60 = v86
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	v115 = l1
	v117 = v109
	v121 = int32(1)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v165 = v163 + int32(4)
	v166 = F_palloc(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L12
	} else {
		goto L45
	}
L33:
	;
	v125 = l2 + int32(20) + v115<<(uint(int32(3))%32)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+4)))
	if v126 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v121 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v147 = v117
	v148 = v121
	goto L37
L37:
	;
	v150 = v115 + int32(1)
	if v150 < base.I32_extend16_s(v147) {
		v115 = v150
		v117 = v147
		v121 = v148
		goto L33
	} else {
		goto L44
	}
L38:
	;
	F_appendStringInfoString(m, v12+int32(8), l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L12
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v141 = F_OutputFunctionCall(m, v106+v115*int32(28), v129)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L12
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	F_appendStringInfoString(m, v12+int32(8), v141)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)))
	v147 = v146
	v148 = int32(0)
	goto L37
L44:
	;
	goto L34
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v165 << (uint(int32(2)) % 32)
	if v163 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	F_pfree(m, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L12
	} else {
		goto L50
	}
L47:
	;
	v173 = F__emscripten_memcpy_bulkmem(m, v166+int32(4), v162, v163)
	mBase = m.M
	goto L49
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	v181 = v166
	goto L2
L51:
	;
	F_errmsg_internal(m, int32(_a_F_concat_internal_0), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_concat_internal_1), int32(_a_F_concat_internal_2), int32(_a_F_concat_internal_3))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_convert_saop_to_hashed_saop_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == v3 {
		v94 = v3
		m.G0 = v9 + int32(16)
		return v94
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v13 != int32(20) {
			v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v94 = v90
				m.G0 = v9 + int32(16)
				return v94
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			if v18 == int32(0) {
				v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					v94 = v90
					m.G0 = v9 + int32(16)
					return v94
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				if v21 != int32(7) {
					v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						v94 = v90
						m.G0 = v9 + int32(16)
						return v94
					}
				} else {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
					if v24 != 0 {
						v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							v94 = v90
							m.G0 = v9 + int32(16)
							return v94
						}
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v26 == int32(1) {
							v33 = F_get_op_hash_functions(m, v25, v9+int32(12), v9+int32(8))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v33 == int32(0) {
									v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										v94 = v90
										m.G0 = v9 + int32(16)
										return v94
									}
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
									v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
									if v39 != v40 {
										v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											v94 = v90
											m.G0 = v9 + int32(16)
											return v94
										}
									} else {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
										v46 = F_ArrayGetNItems(m, v43, v42+int32(16))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int32(0)
										} else {
											if v46 < int32(9) {
												v94 = v3
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												v82 = int32(12)
												v83 = v51
												*(*int32)(unsafe.Add(mBase, uint32(l0+v82))) = v83
												v94 = v3
											}
											m.G0 = v9 + int32(16)
											return v94
										}
									}
								}
							}
						} else {
							v52 = F_get_negator(m, v25)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v52 == int32(0) {
									v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										v94 = v90
										m.G0 = v9 + int32(16)
										return v94
									}
								} else {
									v60 = F_get_op_hash_functions(m, v52, v9+int32(12), v9+int32(8))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										if v60 == int32(0) {
											v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												v94 = v90
												m.G0 = v9 + int32(16)
												return v94
											}
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
											v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
											if v64 != v65 {
												v90 = F_expression_tree_walker_impl(m, l0, int32(874), int32(0))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int32(0)
												} else {
													v94 = v90
													m.G0 = v9 + int32(16)
													return v94
												}
											} else {
												v67 = int32(16)
												v68 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
												v72 = F_ArrayGetNItems(m, v69, v68+v67)
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return int32(0)
												} else {
													if v72 < int32(9) {
														v94 = v3
														m.G0 = v9 + int32(16)
														return v94
													} else {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v76
														v78 = F_get_opcode(m, v52)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															v82 = v67
															v83 = v78
															*(*int32)(unsafe.Add(mBase, uint32(l0+v82))) = v83
															v94 = v3
															m.G0 = v9 + int32(16)
															return v94
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
	}
}
func F_convert_testexpr_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		v40 = int32(0)
		m.G0 = v7 + int32(16)
		return v40
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v12 - int32(8) {
		case 0:
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v15 != int32(2) {
				v38 = F_expression_tree_mutator_impl(m, l0, int32(845), l1)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = v38
					m.G0 = v7 + int32(16)
					return v40
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v18 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v52
						F_errmsg_internal(m, int32(_a_F_convert_testexpr_mutator_0), v7)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_convert_testexpr_mutator_1), int32(669), int32(_a_F_convert_testexpr_mutator_2))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v21 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v52
							F_errmsg_internal(m, int32(_a_F_convert_testexpr_mutator_0), v7)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_convert_testexpr_mutator_1), int32(669), int32(_a_F_convert_testexpr_mutator_2))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
						if v24 < v18 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v52
								F_errmsg_internal(m, int32(_a_F_convert_testexpr_mutator_0), v7)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_convert_testexpr_mutator_1), int32(669), int32(_a_F_convert_testexpr_mutator_2))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+v18<<(uint(int32(2))%32)-int32(4))))
							v33 = F_copyObjectImpl(m, v32)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v40 = v33
								m.G0 = v7 + int32(16)
								return v40
							}
						}
					}
				}
			}
		default:
			v38 = F_expression_tree_mutator_impl(m, l0, int32(845), l1)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = v38
				m.G0 = v7 + int32(16)
				return v40
			}
		case 14:
			v40 = l0
			m.G0 = v7 + int32(16)
			return v40
		}
	}
}
func F_core_yyensure_buffer_stack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		v8 = F_palloc(m, int32(4))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
			if v8 == int32(0) {
				F_yy_fatal_error_2(m, int32(_a_F_core_yyensure_buffer_stack_0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967296)
				return
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(v18-int32(1)) <= base.Ui32(v17) {
			v23 = v18 + int32(8)
			v26 = F_repalloc(m, v4, v23<<(uint(int32(2))%32))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
				if v26 == int32(0) {
					F_yy_fatal_error_2(m, int32(_a_F_core_yyensure_buffer_stack_0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v26 + v31<<(uint(int32(2))%32)
					v35 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					return
				}
			}
		} else {
			return
		}
	}
}
func F_cost_append(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 float64
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 float64
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 float64
	_ = v180
	var v181 float64
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 float64
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v198 float64
	_ = v198
	var v201 float64
	_ = v201
	var v203 float64
	_ = v203
	var v206 float64
	_ = v206
	var v213 float64
	_ = v213
	var v215 float64
	_ = v215
	var v219 int32
	_ = v219
	var v220 float64
	_ = v220
	var v223 int64
	_ = v223
	var v224 float64
	_ = v224
	var v226 float64
	_ = v226
	var v227 int32
	_ = v227
	var v230 float64
	_ = v230
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	var v238 float64
	_ = v238
	var v240 float64
	_ = v240
	var v241 float64
	_ = v241
	var v244 float64
	_ = v244
	var v247 float64
	_ = v247
	var v251 float64
	_ = v251
	var v258 float64
	_ = v258
	var v259 float64
	_ = v259
	var v262 float64
	_ = v262
	var v266 float64
	_ = v266
	var v276 float64
	_ = v276
	var v279 float64
	_ = v279
	var v283 int32
	_ = v283
	var v288 float64
	_ = v288
	var v289 float64
	_ = v289
	var v291 float64
	_ = v291
	var v296 int32
	_ = v296
	var v299 float64
	_ = v299
	var v300 float64
	_ = v300
	var v301 float64
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 float64
	_ = v307
	var v308 float64
	_ = v308
	var v311 float64
	_ = v311
	var v312 float64
	_ = v312
	var v313 float64
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 float64
	_ = v320
	var v322 int32
	_ = v322
	var v328 float64
	_ = v328
	var v332 float64
	_ = v332
	var v335 float64
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 float64
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 float64
	_ = v349
	var v350 float64
	_ = v350
	var v354 float64
	_ = v354
	var v358 float64
	_ = v358
	var v361 float64
	_ = v361
	var v364 float64
	_ = v364
	var v365 float64
	_ = v365
	var v367 float64
	_ = v367
	var v369 float64
	_ = v369
	var v371 float64
	_ = v371
	var v374 float64
	_ = v374
	var v376 float64
	_ = v376
	var v378 float64
	_ = v378
	var v379 int32
	_ = v379
	var v381 float64
	_ = v381
	var v389 float64
	_ = v389
	var v393 float64
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v418 float64
	_ = v418
	var v420 float64
	_ = v420
	var v421 float64
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 float64
	_ = v429
	var v431 float64
	_ = v431
	var v434 float64
	_ = v434
	var v436 float64
	_ = v436
	var v438 float64
	_ = v438
	var v440 int32
	_ = v440
	var v441 float64
	_ = v441
	var v442 float64
	_ = v442
	var v446 float64
	_ = v446
	var v450 float64
	_ = v450
	var v453 float64
	_ = v453
	var v456 float64
	_ = v456
	var v458 float64
	_ = v458
	var v459 float64
	_ = v459
	var v461 float64
	_ = v461
	var v462 float64
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 float64
	_ = v467
	var v475 float64
	_ = v475
	var v479 float64
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v495 int32
	_ = v495
	var v504 float64
	_ = v504
	var v506 float64
	_ = v506
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v551 int32
	_ = v551
	var v552 float64
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 float64
	_ = v660
	var v661 float64
	_ = v661
	var v667 int32
	_ = v667
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v706 float64
	_ = v706
	var v710 float64
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 float64
	_ = v716
	var v720 float64
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 float64
	_ = v726
	var v730 float64
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 float64
	_ = v736
	var v740 float64
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v798 int32
	_ = v798
	var v801 float64
	_ = v801
	var v805 float64
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v939 float64
	_ = v939
	var v943 float64
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 float64
	_ = v949
	var v953 float64
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 float64
	_ = v959
	var v963 float64
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 float64
	_ = v969
	var v973 float64
	_ = v973
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1031 int32
	_ = v1031
	var v1034 float64
	_ = v1034
	var v1038 float64
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1049 int32
	_ = v1049
	var v1070 float64
	_ = v1070
	var v1074 float64
	_ = v1074
	var v1075 float64
	_ = v1075
	var v1095 float64
	_ = v1095
	var v1096 float64
	_ = v1096
	var v1102 float64
	_ = v1102
	v2 = int32(0)
	v19 = float64(0)
	v25 = m.G0
	v27 = v25 - int32(96)
	m.G0 = v27
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v29
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = l0 + int32(48)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v40 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v27 + int32(96)
	return
L4:
	;
	v1102 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[0]))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(base.F64_mul(base.F64_mul(v1102, float64(0.5)), v1096), v1095)
	goto L3
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v43 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v320 = base.F64_convert_i32_s(v319)
	v322 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_append[1])))
	if v322 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L8:
	;
	v99 = v2
	goto L19
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if int32(0) < v44 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v49 = *(*float64)(unsafe.Add(mBase, uint32(v48)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v39))) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v51 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v1095 = v19
	v1096 = v19
	goto L4
L13:
	;
	v1095 = v19
	v1096 = v19
	goto L4
L14:
	;
	goto L15
L15:
	;
	v56 = v2
	v58 = int32(0)
	v73 = v19
	v74 = v19
	goto L16
L16:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v56<<(uint(int32(2))%32))))
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v83)+32))
	v85 = base.F64_add(v84, v74)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+40))
	v88 = v58 + v87
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v88
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v83)+56))
	v91 = base.F64_add(v90, v73)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v91
	v94 = v56 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v94 < v95 {
		v56 = v94
		v58 = v88
		v73 = v91
		v74 = v85
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v1095 = v91
	v1096 = v85
	goto L4
L18:
	;
	goto L17
L19:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v99<<(uint(int32(2))%32))))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+64))
	if v43 == v126 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v1095 = v313
	v1096 = v301
	goto L4
L21:
	;
	v300 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v301 = base.F64_add(v299, v300)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v301
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v296)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v303 + v304
	v307 = *(*float64)(unsafe.Add(mBase, uint32(v296)+48))
	v308 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = base.F64_add(v307, v308)
	v311 = *(*float64)(unsafe.Add(mBase, uint32(v296)+56))
	v312 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v313 = base.F64_add(v311, v312)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v313
	v316 = v99 + int32(1)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v316 < v317 {
		v99 = v316
		goto L19
	} else {
		goto L63
	}
L22:
	;
	if v179 != 0 {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	v179 = int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v135 = int32(0)
	goto L27
L26:
	;
	v179 = v171
	goto L22
L27:
	;
	v139 = int32(0)
	if v43 == v139 {
		v149 = v139
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v171 = int32(0)
	goto L26
L29:
	;
	if v126 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v143 <= v135 {
		v149 = int32(0)
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v149 = v145 + v135<<(uint(int32(2))%32)
	goto L29
L32:
	;
	v155 = base.B2i32(v149 == int32(0))
	if v149 == int32(0) {
		v171 = v155
		goto L26
	} else {
		goto L37
	}
L33:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v135 < v150 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v179 = base.B2i32(v149 == int32(0))
	goto L22
L36:
	;
	goto L35
L37:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v161 = v158 + v135<<(uint(int32(2))%32)
	if v161 == int32(0) {
		v171 = v155
		goto L26
	} else {
		goto L38
	}
L38:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v166 == v167 {
		v135 = v135 + int32(1)
		goto L27
	} else {
		goto L39
	}
L39:
	;
	goto L28
L40:
	;
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v125)+32))
	v296 = v125
	v299 = v180
	goto L21
L41:
	;
	goto L42
L42:
	;
	v181 = *(*float64)(unsafe.Add(mBase, uint32(v125)+56))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v125)+40))
	v184 = v27 + int32(88)
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v125)+32))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+32))
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_cost_append[2]))
	v193 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
	v198 = float64(2)
	if base.F64_lt(v187, v198) != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_append[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v182 + (v283 ^ int32(1))
	v288 = *(*float64)(unsafe.Add(mBase, uint32(v27)+88))
	v289 = base.F64_add(v181, v288)
	*(*float64)(unsafe.Add(mBase, uint32(v27)+56)) = v289
	v291 = *(*float64)(unsafe.Add(mBase, uint32(v27)+80))
	*(*float64)(unsafe.Add(mBase, uint32(v27)+64)) = base.F64_add(v289, v291)
	v296 = v27 + int32(8)
	v299 = v187
	goto L21
L44:
	;
	v201 = v198
	goto L46
L45:
	;
	v201 = v187
	goto L46
L46:
	;
	v203 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[4]))
	v206 = base.F64_mul(v201, base.F64_add(base.F64_add(v203, v203), float64(0)))
	v213 = base.F64_convert_i32_u((v189+int32(7))&int32(-8) + int32(24))
	v215 = base.F64_mul(v187, v213)
	v219 = base.F64_lt(v193, v201) & base.F64_gt(v193, float64(0))
	if v219 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v184))) = v276
	v279 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[4]))
	*(*float64)(unsafe.Add(mBase, uint32(v27+int32(80)))) = base.F64_mul(v201, v279)
	goto L43
L48:
	;
	v220 = base.F64_mul(v193, v213)
	goto L50
L49:
	;
	v220 = v215
	goto L50
L50:
	;
	v223 = base.I64_extend_i32_s(v192) << (uint(int64(10)) % 64)
	v224 = base.F64_convert_i64_s(v223)
	if base.F64_gt(v220, v224) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v226 = F_log(m, v201)
	mBase = m.M
	v227 = F_tuplesort_merge_order(m, v223)
	mBase = m.M
	v230 = base.F64_mul(base.F64_div(v226, float64(0.693147180559945)), v206)
	*(*float64)(unsafe.Add(mBase, uint32(v184))) = v230
	v235 = base.F64_ceil(base.F64_mul(v215, float64(0.0001220703125)))
	v237 = base.F64_div(v215, v224)
	v238 = base.F64_convert_i32_s(v227)
	if base.F64_gt(v237, v238) != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	if v219 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v240 = F_log(m, v237)
	mBase = m.M
	v241 = F_log(m, v238)
	mBase = m.M
	v244 = base.F64_ceil(base.F64_div(v240, v241))
	goto L56
L55:
	;
	v244 = float64(1)
	goto L56
L56:
	;
	v247 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[5]))
	v251 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[6]))
	v276 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v235, v235), v244), base.F64_add(base.F64_mul(v247, float64(0.75)), base.F64_mul(v251, float64(0.25)))), v230)
	goto L47
L57:
	;
	v258 = v193
	goto L59
L58:
	;
	v258 = v201
	goto L59
L59:
	;
	v259 = base.F64_add(v258, v258)
	if base.F64_gt(v215, v224)|base.F64_gt(v201, v259) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v262 = F_log(m, v259)
	mBase = m.M
	v276 = base.F64_mul(base.F64_div(v262, float64(0.693147180559945)), v206)
	goto L47
L61:
	;
	goto L62
L62:
	;
	v266 = F_log(m, v201)
	mBase = m.M
	v276 = base.F64_mul(base.F64_div(v266, float64(0.693147180559945)), v206)
	goto L47
L63:
	;
	goto L20
L64:
	;
	v328 = base.F64_add(base.F64_mul(v320, float64(-0.3)), float64(1))
	if base.F64_gt(v328, float64(0)) != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v335 = v320
	goto L66
L66:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v336 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	v332 = v328
	goto L69
L68:
	;
	v332 = math.Float64frombits(uint64(0x8000000000000000))
	goto L69
L69:
	;
	v335 = base.F64_add(v332, v320)
	goto L66
L70:
	;
	if v495 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L71:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v495 = v339
	v504 = float64(0)
	v506 = v19
	goto L70
L72:
	;
	goto L73
L73:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v342)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v343
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v345 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v342)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v379
	v381 = float64(1e+100)
	if base.F64_gt(v376, v381) != 0 {
		v393 = v381
		goto L84
	} else {
		goto L85
	}
L75:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v342)+24))
	v349 = base.F64_convert_i32_s(v348)
	v350 = *(*float64)(unsafe.Add(mBase, uint32(v342)+32))
	if v322 != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	v371 = *(*float64)(unsafe.Add(mBase, uint32(v342)+32))
	v374 = base.F64_add(base.F64_div(v371, v335), float64(0))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v374
	v376 = v374
	v378 = v19
	goto L74
L78:
	;
	v354 = base.F64_add(base.F64_mul(v349, float64(-0.3)), float64(1))
	if base.F64_gt(v354, float64(0)) != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v361 = v349
	goto L80
L80:
	;
	v364 = float64(0)
	v365 = base.F64_add(base.F64_mul(v350, base.F64_div(v361, v335)), v364)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v365
	v367 = *(*float64)(unsafe.Add(mBase, uint32(v342)+56))
	v369 = base.F64_add(v367, v364)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v369
	v376 = v365
	v378 = v369
	goto L74
L81:
	;
	v358 = v354
	goto L83
L82:
	;
	v358 = math.Float64frombits(uint64(0x8000000000000000))
	goto L83
L83:
	;
	v361 = base.F64_add(v358, v349)
	goto L80
L84:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v393
	v395 = int32(1)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v396 <= v395 {
		v495 = v345
		v504 = v393
		v506 = v378
		goto L70
	} else {
		goto L88
	}
L85:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v376)&int64(9223372036854775807)) {
		v393 = v381
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v389 = float64(1)
	if base.F64_le(v376, v389) != 0 {
		v393 = v389
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v393 = base.F64_nearest(v376)
	goto L84
L88:
	;
	v400 = v395
	v402 = v379
	v418 = v393
	v420 = v378
	v421 = v343
	goto L89
L89:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v423+v400<<(uint(int32(2))%32))))
	if v400 < v319 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v495 = v345
	v504 = v479
	v506 = v462
	goto L70
L91:
	;
	v429 = *(*float64)(unsafe.Add(mBase, uint32(v427)+48))
	if base.F64_gt(v429, v421) != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v434 = v421
	goto L93
L93:
	;
	if v400 < v345 {
		goto L98
	} else {
		goto L99
	}
L94:
	;
	v431 = v421
	goto L96
L95:
	;
	v431 = v429
	goto L96
L96:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v39))) = v431
	v434 = v431
	goto L93
L97:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v427)+40))
	v465 = v402 + v464
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v465
	v467 = float64(1e+100)
	if base.F64_gt(v461, v467) != 0 {
		v479 = v467
		goto L107
	} else {
		goto L108
	}
L98:
	;
	v436 = *(*float64)(unsafe.Add(mBase, uint32(v427)+32))
	v438 = base.F64_add(v418, base.F64_div(v436, v335))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v438
	v461 = v438
	v462 = v420
	goto L97
L99:
	;
	goto L100
L100:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v427)+24))
	v441 = base.F64_convert_i32_s(v440)
	v442 = *(*float64)(unsafe.Add(mBase, uint32(v427)+32))
	if v322 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v446 = base.F64_add(base.F64_mul(v441, float64(-0.3)), float64(1))
	if base.F64_gt(v446, float64(0)) != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v453 = v441
	goto L103
L103:
	;
	v456 = base.F64_add(base.F64_mul(v442, base.F64_div(v453, v335)), v418)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v456
	v458 = *(*float64)(unsafe.Add(mBase, uint32(v427)+56))
	v459 = base.F64_add(v458, v420)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v459
	v461 = v456
	v462 = v459
	goto L97
L104:
	;
	v450 = v446
	goto L106
L105:
	;
	v450 = math.Float64frombits(uint64(0x8000000000000000))
	goto L106
L106:
	;
	v453 = base.F64_add(v450, v441)
	goto L103
L107:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v479
	v482 = v400 + int32(1)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v482 < v483 {
		v400 = v482
		v402 = v465
		v418 = v479
		v420 = v462
		v421 = v434
		goto L89
	} else {
		goto L111
	}
L108:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v461)&int64(9223372036854775807)) {
		v479 = v467
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v475 = float64(1)
	if base.F64_le(v461, v475) != 0 {
		v479 = v475
		goto L107
	} else {
		goto L110
	}
L110:
	;
	v479 = base.F64_nearest(v461)
	goto L107
L111:
	;
	goto L90
L112:
	;
	v1095 = base.F64_add(v506, float64(0))
	v1096 = v504
	goto L4
L113:
	;
	goto L114
L114:
	;
	if v319 < v495 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v514 = v319
	goto L117
L116:
	;
	v514 = v495
	goto L117
L117:
	;
	v517 = F_palloc(m, v514<<(uint(int32(3))%32))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	return
L119:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v519 <= int32(0) {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v1070 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v1074 = *(*float64)(unsafe.Add(mBase, uint32(v517+v1049<<(uint(int32(3))%32))))
	v1075 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v1095 = base.F64_add(v1074, v1075)
	v1096 = v1070
	goto L4
L121:
	;
	v894 = int32(3)
	v895 = v514 & v894
	v896 = int32(0)
	if base.Ui32(v894) <= base.Ui32(v514-int32(1)) {
		goto L170
	} else {
		goto L171
	}
L122:
	;
	v614 = int32(2)
	v617 = v514 << (uint(v614) % 32) >> (uint(v614) % 32)
	if v568 <= v617 {
		goto L134
	} else {
		goto L135
	}
L123:
	;
	v611 = int32(0)
	if v514 <= v611 {
		v1049 = v611
		goto L120
	} else {
		goto L133
	}
L124:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v514 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v526 = int32(0)
	v527 = v522
	goto L128
L126:
	;
	v566 = v522
	v568 = v519
	goto L127
L127:
	;
	if v566 != 0 {
		goto L122
	} else {
		goto L132
	}
L128:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	v552 = *(*float64)(unsafe.Add(mBase, uint32(v551)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v517+v526<<(uint(int32(3))%32)))) = v552
	v555 = v526 + int32(1)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v556 <= v555 {
		goto L123
	} else {
		goto L130
	}
L129:
	;
	v566 = v561
	v568 = v556
	goto L127
L130:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v561 = v558 + v555<<(uint(int32(2))%32)
	if v555 != v514 {
		v526 = v555
		v527 = v561
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	goto L123
L133:
	;
	goto L121
L134:
	;
	if int32(0) < v514 {
		goto L121
	} else {
		goto L169
	}
L135:
	;
	v622 = v514 & int32(3)
	v630 = v514 - int32(1)
	v640 = v617
	v641 = v514
	goto L136
L136:
	;
	if v495 == v641 {
		goto L134
	} else {
		goto L138
	}
L137:
	;
	goto L134
L138:
	;
	v654 = v517 + v630<<(uint(int32(3))%32)
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v655+v640<<(uint(int32(2))%32))))
	v660 = *(*float64)(unsafe.Add(mBase, uint32(v659)+56))
	v661 = *(*float64)(unsafe.Add(mBase, uint32(v654)))
	*(*float64)(unsafe.Add(mBase, uint32(v654))) = base.F64_add(v660, v661)
	if v514 <= int32(0) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v837 = int32(1)
	v840 = v640 + v837
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v840 < v841 {
		v630 = v816
		v640 = v840
		v641 = v641 + v837
		goto L136
	} else {
		goto L168
	}
L140:
	;
	v816 = int32(0)
	goto L139
L141:
	;
	goto L142
L142:
	;
	v667 = int32(0)
	if base.B2i32(base.Ui32(v514) < base.Ui32(int32(4))) == v667 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v675 = v667
	v676 = v667
	v681 = v667
	goto L146
L144:
	;
	v750 = v667
	v751 = v667
	goto L145
L145:
	;
	if v622 == int32(0) {
		v816 = v751
		goto L139
	} else {
		goto L161
	}
L146:
	;
	v697 = int32(3)
	v698 = v675 | v697
	v700 = v675 | int32(2)
	v702 = v675 | int32(1)
	v706 = *(*float64)(unsafe.Add(mBase, uint32(v517+v675<<(uint(v697)%32))))
	v710 = *(*float64)(unsafe.Add(mBase, uint32(v517+v676<<(uint(v697)%32))))
	if base.F64_lt(v706, v710) != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v750 = v744
	v751 = v742
	goto L145
L148:
	;
	v712 = v675
	goto L150
L149:
	;
	v712 = v676
	goto L150
L150:
	;
	v713 = int32(3)
	v716 = *(*float64)(unsafe.Add(mBase, uint32(v517+v702<<(uint(v713)%32))))
	v720 = *(*float64)(unsafe.Add(mBase, uint32(v517+v712<<(uint(v713)%32))))
	if base.F64_lt(v716, v720) != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v722 = v702
	goto L153
L152:
	;
	v722 = v712
	goto L153
L153:
	;
	v723 = int32(3)
	v726 = *(*float64)(unsafe.Add(mBase, uint32(v517+v700<<(uint(v723)%32))))
	v730 = *(*float64)(unsafe.Add(mBase, uint32(v517+v722<<(uint(v723)%32))))
	if base.F64_lt(v726, v730) != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v732 = v700
	goto L156
L155:
	;
	v732 = v722
	goto L156
L156:
	;
	v733 = int32(3)
	v736 = *(*float64)(unsafe.Add(mBase, uint32(v517+v698<<(uint(v733)%32))))
	v740 = *(*float64)(unsafe.Add(mBase, uint32(v517+v732<<(uint(v733)%32))))
	if base.F64_lt(v736, v740) != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v742 = v698
	goto L159
L158:
	;
	v742 = v732
	goto L159
L159:
	;
	v743 = int32(4)
	v744 = v675 + v743
	v746 = v681 + v743
	if v746 != v514&int32(2147483644) {
		v675 = v744
		v676 = v742
		v681 = v746
		goto L146
	} else {
		goto L160
	}
L160:
	;
	goto L147
L161:
	;
	v776 = v750
	v777 = v751
	v779 = v667
	goto L162
L162:
	;
	v798 = int32(3)
	v801 = *(*float64)(unsafe.Add(mBase, uint32(v517+v776<<(uint(v798)%32))))
	v805 = *(*float64)(unsafe.Add(mBase, uint32(v517+v777<<(uint(v798)%32))))
	if base.F64_lt(v801, v805) != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v816 = v807
	goto L139
L164:
	;
	v807 = v776
	goto L166
L165:
	;
	v807 = v777
	goto L166
L166:
	;
	v808 = int32(1)
	v811 = v779 + v808
	if v811 != v622 {
		v776 = v776 + v808
		v777 = v807
		v779 = v811
		goto L162
	} else {
		goto L167
	}
L167:
	;
	goto L163
L168:
	;
	goto L137
L169:
	;
	v1049 = int32(0)
	goto L120
L170:
	;
	v908 = v896
	v909 = v896
	v911 = int32(0)
	goto L173
L171:
	;
	v983 = v896
	v984 = v896
	goto L172
L172:
	;
	if v895 == int32(0) {
		v1049 = v984
		goto L120
	} else {
		goto L188
	}
L173:
	;
	v930 = int32(3)
	v931 = v908 | v930
	v933 = v908 | int32(2)
	v935 = v908 | int32(1)
	v939 = *(*float64)(unsafe.Add(mBase, uint32(v517+v908<<(uint(v930)%32))))
	v943 = *(*float64)(unsafe.Add(mBase, uint32(v517+v909<<(uint(v930)%32))))
	if base.F64_gt(v939, v943) != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v983 = v977
	v984 = v975
	goto L172
L175:
	;
	v945 = v908
	goto L177
L176:
	;
	v945 = v909
	goto L177
L177:
	;
	v946 = int32(3)
	v949 = *(*float64)(unsafe.Add(mBase, uint32(v517+v935<<(uint(v946)%32))))
	v953 = *(*float64)(unsafe.Add(mBase, uint32(v517+v945<<(uint(v946)%32))))
	if base.F64_gt(v949, v953) != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v955 = v935
	goto L180
L179:
	;
	v955 = v945
	goto L180
L180:
	;
	v956 = int32(3)
	v959 = *(*float64)(unsafe.Add(mBase, uint32(v517+v933<<(uint(v956)%32))))
	v963 = *(*float64)(unsafe.Add(mBase, uint32(v517+v955<<(uint(v956)%32))))
	if base.F64_gt(v959, v963) != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v965 = v933
	goto L183
L182:
	;
	v965 = v955
	goto L183
L183:
	;
	v966 = int32(3)
	v969 = *(*float64)(unsafe.Add(mBase, uint32(v517+v931<<(uint(v966)%32))))
	v973 = *(*float64)(unsafe.Add(mBase, uint32(v517+v965<<(uint(v966)%32))))
	if base.F64_gt(v969, v973) != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v975 = v931
	goto L186
L185:
	;
	v975 = v965
	goto L186
L186:
	;
	v976 = int32(4)
	v977 = v908 + v976
	v979 = v911 + v976
	if v979 != v514&int32(-4) {
		v908 = v977
		v909 = v975
		v911 = v979
		goto L173
	} else {
		goto L187
	}
L187:
	;
	goto L174
L188:
	;
	v1009 = v983
	v1010 = v984
	v1011 = v896
	goto L189
L189:
	;
	v1031 = int32(3)
	v1034 = *(*float64)(unsafe.Add(mBase, uint32(v517+v1009<<(uint(v1031)%32))))
	v1038 = *(*float64)(unsafe.Add(mBase, uint32(v517+v1010<<(uint(v1031)%32))))
	if base.F64_gt(v1034, v1038) != 0 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v1049 = v1040
	goto L120
L191:
	;
	v1040 = v1009
	goto L193
L192:
	;
	v1040 = v1010
	goto L193
L193:
	;
	v1041 = int32(1)
	v1044 = v1011 + v1041
	if v1044 != v895 {
		v1009 = v1009 + v1041
		v1010 = v1040
		v1011 = v1044
		goto L189
	} else {
		goto L194
	}
L194:
	;
	goto L190
}
func F_countVariablesFromJsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	if l0 == int32(0) {
		return base.B2i32(l0 != int32(0))
	} else {
		v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
		if v4&int32(32) != 0 {
			return base.B2i32(l0 != int32(0))
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_countVariablesFromJsonb_0), int32(0))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						F_errdetail(m, int32(_a_F_countVariablesFromJsonb_1), int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_countVariablesFromJsonb_2), int32(3211), int32(_a_F_countVariablesFromJsonb_3))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_countitem_compare_count(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	return base.B2i32(v7 < v5) - base.B2i32(v5 < v7)
}
