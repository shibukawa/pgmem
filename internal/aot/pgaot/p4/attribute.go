package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckAttributeType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v33 int32
	_ = v33
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
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
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	v17 = F_get_typtype(m, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	switch v17&int32(255) - int32(99) {
	case 0:
		goto L13
	case 1:
		goto L14
	default:
		goto L11
	case 13:
		goto L15
	case 15:
		goto L12
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l0
	F_errmsg(m, int32(_a_F_CheckAttributeType_0), v15+int32(32))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L97
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L92
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L87
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L83
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = l0
	F_errmsg(m, int32(_a_F_CheckAttributeType_1), v15-int32(-64))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L81
	}
L9:
	;
	if l2 != 0 {
		goto L70
	} else {
		goto L71
	}
L10:
	;
	if base.Ui32(l1) < base.Ui32(int32(_a_F_CheckAttributeType_2)) {
		goto L9
	} else {
		goto L68
	}
L11:
	;
	v191 = F_get_element_type(m, l1)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L65
	}
L12:
	;
	v175 = F_get_range_subtype(m, l1)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L58
	}
L13:
	;
	v71 = int32(0)
	if l3 == v71 {
		goto L30
	} else {
		goto L31
	}
L14:
	;
	if l4&int32(8) != 0 {
		goto L7
	} else {
		goto L26
	}
L15:
	;
	if l4&int32(2) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v33 = base.B2i32(l1 != int32(2249)) & base.B2i32(l1 != int32(2287))
	goto L18
L17:
	;
	v33 = int32(1)
	goto L18
L18:
	;
	if base.B2i32(v33 == int32(0))|l4&int32(1)&base.B2i32(l1 == int32(2277)) != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v49 = F_format_type_be(m, l1)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if l4&int32(4) != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l0
	F_errmsg(m, int32(_a_F_CheckAttributeType_3), v15+int32(48))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(581), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v67 = F_getBaseType(m, l1)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_CheckAttributeType(m, l0, v67, l2, l3, l4)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L10
L29:
	;
	if v109 != 0 {
		goto L6
	} else {
		goto L42
	}
L30:
	;
	v109 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v77 <= int32(0) {
		v103 = v71
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v109 = v103
	goto L29
L34:
	;
	v80 = int32(0)
	if v80 < v77 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v83 = v77
	goto L37
L36:
	;
	v83 = v80
	goto L37
L37:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v86 = int32(0)
	goto L38
L38:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84+v86<<(uint(int32(2))%32))))
	v95 = base.B2i32(v94 == l1)
	if v94 == l1 {
		v103 = v95
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v103 = v95
	goto L33
L40:
	;
	v97 = v86 + int32(1)
	if v97 != v83 {
		v86 = v97
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v110 = F_lappend_oid(m, l3, l1)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v112 = F_get_typ_typrelid(m, l1)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v115 = F_relation_open(m, v112, int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+52))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if int32(0) < v118 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v127 = v118
	v130 = int32(0)
	goto L49
L47:
	;
	goto L48
L48:
	;
	F_relation_close(m, v115, int32(1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L56
	}
L49:
	;
	v141 = v117 + v127<<(uint(int32(4))%32) + v130*int32(100)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+111)))
	if v142 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L48
L51:
	;
	v146 = v141 + int32(20)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)+68))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146)+96))
	F_CheckAttributeType(m, v141+int32(24), v149, v150, v110, l4&int32(-5))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	v154 = v127
	goto L53
L53:
	;
	v156 = v130 + int32(1)
	if v156 < v154 {
		v127 = v154
		v130 = v156
		goto L49
	} else {
		goto L55
	}
L54:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v154 = v153
	goto L53
L55:
	;
	goto L50
L56:
	;
	v173 = F_list_delete_last(m, v110)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	goto L10
L58:
	;
	v178 = F_SearchSysCache1(m, int32(55), l1)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	if v178 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v178)+16))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+22)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v180+v181)+12))
	F_ReleaseCatCache(m, v178)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	v188 = int32(0)
	goto L62
L62:
	;
	F_CheckAttributeType(m, l0, v175, v188, l3, l4)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	v188 = v183
	goto L62
L64:
	;
	goto L10
L65:
	;
	if v191 == int32(0) {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	F_CheckAttributeType(m, l0, v191, l2, l3, l4)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	goto L10
L68:
	;
	if l4&int32(8) != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	goto L9
L70:
	;
	m.G0 = v15 + int32(112)
	return
L71:
	;
	v225 = F_type_is_collatable(m, l1)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v225 == int32(0) {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v236 = F_format_type_be(m, l1)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if l4&int32(4) != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
	F_errmsg(m, int32(_a_F_CheckAttributeType_6), v15+int32(16))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errhint(m, int32(_a_F_CheckAttributeType_7), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(694), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(576), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l0
	F_errmsg(m, int32(_a_F_CheckAttributeType_8), v15+int32(80))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(595), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v296 = F_format_type_be(m, l1)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v296
	F_errmsg(m, int32(_a_F_CheckAttributeType_9), v15+int32(96))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(623), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg(m, int32(_a_F_CheckAttributeType_10), v15)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errdetail(m, int32(_a_F_CheckAttributeType_11), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(674), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errhint(m, int32(_a_F_CheckAttributeType_7), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(688), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_DeleteAttributeTuples(m *base.Module, l0 int32) {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v11 = F_table_open(m, int32(1249), int32(3))
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v7, int32(1), int32(3), int32(184), l0)
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(1)
	v22 = F_systable_beginscan(m, v11, int32(2659), v19, int32(0), v19, v7)
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = F_systable_getnext(m, v22)
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = v24
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_systable_endscan(m, v22)
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	F_simple_heap_delete(m, v11, v26+int32(4))
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v34 = F_systable_getnext(m, v22)
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v34 != 0 {
		v26 = v34
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	F_relation_close(m, v11, int32(3))
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	m.G0 = v7 + int32(48)
	return
}
