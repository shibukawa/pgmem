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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
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
		goto L14
	case 1:
		goto L15
	default:
		goto L11
	case 10:
		goto L12
	case 13:
		goto L16
	case 15:
		goto L13
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l0
	F_errmsg(m, int32(_a_F_CheckAttributeType_0), v15+int32(32))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L100
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L95
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L90
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L86
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = l0
	F_errmsg(m, int32(_a_F_CheckAttributeType_1), v15-int32(-64))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L84
	}
L9:
	;
	if l2 != 0 {
		goto L73
	} else {
		goto L74
	}
L10:
	;
	if base.Ui32(l1) < base.Ui32(int32(_a_F_CheckAttributeType_2)) {
		goto L9
	} else {
		goto L71
	}
L11:
	;
	v197 = F_get_element_type(m, l1)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L68
	}
L12:
	;
	v192 = F_get_multirange_range(m, l1)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L66
	}
L13:
	;
	v175 = F_get_range_subtype(m, l1)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L59
	}
L14:
	;
	v71 = int32(0)
	if l3 == v71 {
		goto L31
	} else {
		goto L32
	}
L15:
	;
	if l4&int32(8) != 0 {
		goto L7
	} else {
		goto L27
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
	v33 = base.B2i32(l1 != int32(2249)) & base.B2i32(l1 != int32(2287))
	goto L19
L18:
	;
	v33 = int32(1)
	goto L19
L19:
	;
	if base.B2i32(v33 == int32(0))|l4&int32(1)&base.B2i32(l1 == int32(2277)) != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l0
	F_errmsg(m, int32(_a_F_CheckAttributeType_3), v15+int32(48))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(586), int32(_a_F_CheckAttributeType_5))
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
		v103 = v71
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v109 = v103
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
		v103 = v95
		goto L34
	} else {
		goto L41
	}
L40:
	;
	v103 = v95
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
	v127 = v118
	v130 = int32(0)
	goto L50
L48:
	;
	goto L49
L49:
	;
	F_relation_close(m, v115, int32(1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L57
	}
L50:
	;
	v141 = v117 + v127<<(uint(int32(3))%32) + v130*int32(100)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+119)))
	if v142 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	v146 = v141 + int32(28)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)+68))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146)+96))
	F_CheckAttributeType(m, v141+int32(32), v149, v150, v110, l4&int32(-5))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v154 = v127
	goto L54
L54:
	;
	v156 = v130 + int32(1)
	if v156 < v154 {
		v127 = v154
		v130 = v156
		goto L50
	} else {
		goto L56
	}
L55:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v154 = v153
	goto L54
L56:
	;
	goto L51
L57:
	;
	v173 = F_list_delete_last(m, v110)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L10
L59:
	;
	v179 = F_SearchSysCache1(m, int32(55), base.I64_extend_i32_u(l1))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v179 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+22)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v181+v182)+12))
	F_ReleaseCatCache(m, v179)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	v189 = int32(0)
	goto L63
L63:
	;
	F_CheckAttributeType(m, l0, v175, v189, l3, l4)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	v189 = v184
	goto L63
L65:
	;
	goto L10
L66:
	;
	F_CheckAttributeType(m, l0, v192, int32(0), l3, l4)
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
	if v197 == int32(0) {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	F_CheckAttributeType(m, l0, v197, l2, l3, l4)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L10
L71:
	;
	if l4&int32(8) != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	goto L9
L73:
	;
	m.G0 = v15 + int32(112)
	return
L74:
	;
	v231 = F_type_is_collatable(m, l1)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v231 == int32(0) {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v242 = F_format_type_be(m, l1)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if l4&int32(4) != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
	F_errmsg(m, int32(_a_F_CheckAttributeType_6), v15+int32(16))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errhint(m, int32(_a_F_CheckAttributeType_7), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(709), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
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
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(581), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l0
	F_errmsg(m, int32(_a_F_CheckAttributeType_8), v15+int32(80))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(600), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v302 = F_format_type_be(m, l1)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v302
	F_errmsg(m, int32(_a_F_CheckAttributeType_9), v15+int32(96))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(628), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg(m, int32(_a_F_CheckAttributeType_10), v15)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v328 = F_errdetail(m, int32(_a_F_CheckAttributeType_11), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(689), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errhint(m, int32(_a_F_CheckAttributeType_7), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeType_4), int32(703), int32(_a_F_CheckAttributeType_5))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 + int32(-64)
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
	F_ScanKeyInit(m, v7, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(1)
	v23 = F_systable_beginscan(m, v11, int32(2659), v20, int32(0), v20, v7)
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = F_systable_getnext(m, v23)
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = v25
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_systable_endscan(m, v23)
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	F_simple_heap_delete(m, v11, v27+int32(4))
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v35 = F_systable_getnext(m, v23)
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v35 != 0 {
		v27 = v35
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
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	m.G0 = v7 - int32(-64)
	return
}
