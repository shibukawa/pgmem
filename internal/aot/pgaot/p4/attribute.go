package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckAttributeType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v39 int32
	_ = v39
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
	var v102 int32
	_ = v102
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
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
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
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	v18 = F_get_typtype(m, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
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
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	switch v18&int32(255) - int32(99) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l0
	F_errmsg(m, int32(192624), v16+int32(32))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L98
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L93
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L88
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L84
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l0
	F_errmsg(m, int32(188136), v16-int32(-64))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L82
	}
L9:
	;
	if l2 != 0 {
		goto L71
	} else {
		goto L72
	}
L10:
	;
	if base.Ui32(l1) < base.Ui32(int32(12000)) {
		goto L9
	} else {
		goto L69
	}
L11:
	;
	v193 = F_get_element_type(m, l1)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L66
	}
L12:
	;
	v177 = F_get_range_subtype(m, l1)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L59
	}
L13:
	;
	v71 = int32(0)
	if l3 == v71 {
		goto L31
	} else {
		goto L32
	}
L14:
	;
	if l4&int32(8) != 0 {
		goto L7
	} else {
		goto L27
	}
L15:
	;
	if l4&int32(1)&base.B2i32(l1 == int32(2277)) != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	if l4&int32(2) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v39 = base.B2i32(l1 != int32(2249)) & base.B2i32(l1 != int32(2287))
	goto L19
L18:
	;
	v39 = int32(1)
	goto L19
L19:
	;
	if v39 == int32(0) {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v49 = F_format_type_be(m, l1)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if l4&int32(4) != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l0
	F_errmsg(m, int32(188179), v16+int32(48))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(495872), int32(581), int32(371418))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v67 = F_getBaseType(m, l1)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_CheckAttributeType(m, l0, v67, l2, l3, l4)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L10
L30:
	;
	if v109 != 0 {
		goto L6
	} else {
		goto L43
	}
L31:
	;
	v109 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v77 <= int32(0) {
		v102 = v71
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v109 = v102
	goto L30
L35:
	;
	v80 = int32(0)
	if v80 < v77 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v83 = v77
	goto L38
L37:
	;
	v83 = v80
	goto L38
L38:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v86 = int32(0)
	goto L39
L39:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84+v86<<(uint(int32(2))%32))))
	v95 = base.B2i32(v94 == l1)
	if v94 == l1 {
		v102 = v95
		goto L34
	} else {
		goto L41
	}
L40:
	;
	v102 = v95
	goto L34
L41:
	;
	v97 = v86 + int32(1)
	if v97 != v83 {
		v86 = v97
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v110 = F_lappend_oid(m, l3, l1)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v112 = F_get_typ_typrelid(m, l1)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v115 = F_relation_open(m, v112, int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+52))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if int32(0) < v118 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v129 = v118
	v132 = int32(0)
	goto L50
L48:
	;
	goto L49
L49:
	;
	F_relation_close(m, v115, int32(1))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L57
	}
L50:
	;
	v144 = v117 + int32(20) + v129<<(uint(int32(4))%32) + v132*int32(100)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+91)))
	if v145 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v144)+68))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v144)+96))
	F_CheckAttributeType(m, v144+int32(4), v150, v151, v110, l4&int32(-5))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v155 = v129
	goto L54
L54:
	;
	v157 = v132 + int32(1)
	if v157 < v155 {
		v129 = v155
		v132 = v157
		goto L50
	} else {
		goto L56
	}
L55:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v155 = v154
	goto L54
L56:
	;
	goto L51
L57:
	;
	v175 = F_list_delete_last(m, v110)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L10
L59:
	;
	v180 = F_SearchSysCache1(m, int32(55), l1)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v180 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+22)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v182+v183)+12))
	F_ReleaseCatCache(m, v180)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	v190 = int32(0)
	goto L63
L63:
	;
	F_CheckAttributeType(m, l0, v177, v190, l3, l4)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	v190 = v185
	goto L63
L65:
	;
	goto L10
L66:
	;
	if v193 == int32(0) {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	F_CheckAttributeType(m, l0, v193, l2, l3, l4)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	goto L10
L69:
	;
	if l4&int32(8) != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	goto L9
L71:
	;
	m.G0 = v16 + int32(112)
	return
L72:
	;
	v229 = F_type_is_collatable(m, l1)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v229 == int32(0) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v240 = F_format_type_be(m, l1)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if l4&int32(4) != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l0
	F_errmsg(m, int32(192701), v16+int32(16))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errhint(m, int32(574636), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(495872), int32(694), int32(371418))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errfinish(m, int32(495872), int32(576), int32(371418))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = l0
	F_errmsg(m, int32(368701), v16+int32(80))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(495872), int32(595), int32(371418))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v300 = F_format_type_be(m, l1)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v300
	F_errmsg(m, int32(338656), v16+int32(96))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(495872), int32(623), int32(371418))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	F_errmsg(m, int32(370468), v16)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errdetail(m, int32(647405), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(495872), int32(674), int32(371418))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errhint(m, int32(574636), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(495872), int32(688), int32(371418))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
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
	F_CatalogTupleDelete(m, v11, v26+int32(4))
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
	F_sequence_close(m, v11, int32(3))
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
