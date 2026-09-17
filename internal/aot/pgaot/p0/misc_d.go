package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_DeconstructFkConstraintRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v22 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(21))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L8
	} else {
		goto L107
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L8
	} else {
		goto L104
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L101
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L8
	} else {
		goto L98
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L8
	} else {
		goto L95
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L8
	} else {
		goto L92
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L8
	} else {
		goto L89
	}
L8:
	;
	return
L9:
	;
	v24 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v26 != int32(1) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v29 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v30 != int32(21) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if base.Ui32(v33-int32(33)) <= base.Ui32(int32(-33)) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v39 = v33 << (uint(int32(1)) % 32)
	v40 = int32(0)
	v41 = base.B2i32(v39 == v40)
	if v41 == v40 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	base.MemoryCopy(m, l2, v24+int32(24), v39)
	goto L17
L16:
	;
	goto L17
L17:
	;
	if v24 != v22 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_pfree(m, v24)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v52 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(22))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v54 = F_pg_detoast_datum(m, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v56 != int32(1) {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v59 != v33 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	if v61 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v62 != int32(21) {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	if v41 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	base.MemoryCopy(m, l3, v54+int32(24), v39)
	goto L30
L29:
	;
	goto L30
L30:
	;
	if v54 != v52 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_pfree(m, v54)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if l4 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	if l5 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L36:
	;
	v77 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(23))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v79 = F_pg_detoast_datum(m, v77)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v81 != int32(1) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	if v84 != v33 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if v86 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	if v87 != int32(26) {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v91 = v33 << (uint(int32(2)) % 32)
	if v91 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	base.MemoryCopy(m, l4, v79+int32(24), v91)
	goto L45
L44:
	;
	goto L45
L45:
	;
	if v79 == v77 {
		goto L35
	} else {
		goto L46
	}
L46:
	;
	F_pfree(m, v79)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	goto L35
L48:
	;
	if l6 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L49:
	;
	v105 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(24))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v107 = F_pg_detoast_datum(m, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v109 != int32(1) {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v107)+16))
	if v112 != v33 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	if v114 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	if v115 != int32(26) {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v119 = v33 << (uint(int32(2)) % 32)
	if v119 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	base.MemoryCopy(m, l5, v107+int32(24), v119)
	goto L58
L57:
	;
	goto L58
L58:
	;
	if v107 == v105 {
		goto L48
	} else {
		goto L59
	}
L59:
	;
	F_pfree(m, v107)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	goto L48
L61:
	;
	if l8 != 0 {
		goto L74
	} else {
		goto L75
	}
L62:
	;
	v133 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(25))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	v135 = F_pg_detoast_datum(m, v133)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	if v137 != int32(1) {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
	if v140 != v33 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	if v142 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	if v143 != int32(26) {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v147 = v33 << (uint(int32(2)) % 32)
	if v147 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	base.MemoryCopy(m, l6, v135+int32(24), v147)
	goto L71
L70:
	;
	goto L71
L71:
	;
	if v135 == v133 {
		goto L61
	} else {
		goto L72
	}
L72:
	;
	F_pfree(m, v135)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	goto L61
L74:
	;
	v162 = F_SysCacheGetAttr(m, int32(19), l0, int32(26), v18+int32(15))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
	m.G0 = v18 + int32(16)
	return
L77:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
	if v164 != 0 {
		v185 = int32(0)
		goto L78
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v185
	goto L76
L79:
	;
	v165 = F_pg_detoast_datum(m, v162)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v167 != int32(1) {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	if v170 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	if v171 != int32(21) {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v176 = v174 << (uint(int32(1)) % 32)
	if v176 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	base.MemoryCopy(m, l8, v165+int32(24), v176)
	goto L86
L85:
	;
	goto L86
L86:
	;
	if v165 == v162 {
		v185 = v174
		goto L78
	} else {
		goto L87
	}
L87:
	;
	F_pfree(m, v165)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	v185 = v174
	goto L78
L89:
	;
	F_errmsg_internal(m, int32(_a_F_DeconstructFkConstraintRow_0), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1557), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L8
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
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v33
	F_errmsg_internal(m, int32(_a_F_DeconstructFkConstraintRow_3), v18)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1560), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
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
	F_errmsg_internal(m, int32(_a_F_DeconstructFkConstraintRow_4), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1572), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L8
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
	F_errmsg_internal(m, int32(_a_F_DeconstructFkConstraintRow_8), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1587), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errmsg_internal(m, int32(_a_F_DeconstructFkConstraintRow_7), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1602), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	F_errmsg_internal(m, int32(_a_F_DeconstructFkConstraintRow_6), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1617), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_errmsg_internal(m, int32(_a_F_DeconstructFkConstraintRow_5), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1639), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_DetermineTimeZoneAbbrevOffsetTS(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
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
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v186 int64
	_ = v186
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int64
	_ = v423
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v457 int64
	_ = v457
	var v458 int64
	_ = v458
	var v460 int32
	_ = v460
	var v462 int64
	_ = v462
	var v467 int32
	_ = v467
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v498 int32
	_ = v498
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	v8 = m.G0
	v10 = v8 - int32(288)
	m.G0 = v10
	v13 = base.I64_div_s(l0, int64(1000000))
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+280)) = v13 + int64(946684800)
	v18 = v10 + int32(16)
	goto L5
L2:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+16)))
	if v138 != 0 {
		goto L33
	} else {
		goto L34
	}
L3:
	;
	v135 = F_strlen(m, v124)
	mBase = m.M
	goto L2
L5:
	;
	goto L6
L6:
	;
	v25 = int32(255)
	if (v18^l1)&int32(3) != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v128)
	goto L3
L8:
	;
	v109 = v104
	v110 = v105
	v111 = v106
	goto L29
L9:
	;
	if v99 == int32(0) {
		v124 = v97
		v125 = v98
		goto L7
	} else {
		goto L28
	}
L10:
	;
	v97 = l1
	v98 = v18
	v99 = v25
	goto L9
L11:
	;
	goto L12
L12:
	;
	v29 = int32(0)
	if base.B2i32(l1&int32(3) == v29)|int32(0) == v29 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v65 == int32(0) {
		v124 = v62
		v125 = v63
		goto L7
	} else {
		goto L22
	}
L14:
	;
	v41 = l1
	v42 = v18
	v43 = v25
	goto L17
L15:
	;
	goto L16
L16:
	;
	v62 = l1
	v63 = v18
	v64 = v25
	v65 = int32(1)
	goto L13
L17:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v45)
	if v45 == int32(0) {
		v104 = v41
		v105 = v42
		v106 = v43
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v62 = v56
	v63 = v50
	v64 = v52
	v65 = v54
	goto L13
L19:
	;
	v49 = int32(1)
	v50 = v42 + v49
	v52 = v43 - v49
	v53 = int32(0)
	v54 = base.B2i32(v52 != v53)
	v56 = v41 + v49
	if v56&int32(3) == v53 {
		v62 = v56
		v63 = v50
		v64 = v52
		v65 = v54
		goto L13
	} else {
		goto L20
	}
L20:
	;
	if v52 != 0 {
		v41 = v56
		v42 = v50
		v43 = v52
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if base.B2i32(v68 == int32(0))|base.B2i32(base.Ui32(v64) < base.Ui32(int32(4))) != 0 {
		v97 = v62
		v98 = v63
		v99 = v64
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v75 = v62
	v76 = v63
	v77 = v64
	goto L24
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v83 = int32(-2139062144)
	if (int32(16843008)-v80|v80)&v83 != v83 {
		v104 = v75
		v105 = v76
		v106 = v77
		goto L8
	} else {
		goto L26
	}
L25:
	;
	v97 = v91
	v98 = v89
	v99 = v93
	goto L9
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v80
	v88 = int32(4)
	v89 = v76 + v88
	v91 = v75 + v88
	v93 = v77 - v88
	if base.Ui32(int32(3)) < base.Ui32(v93) {
		v75 = v91
		v76 = v89
		v77 = v93
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v104 = v97
	v105 = v98
	v106 = v99
	goto L8
L29:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	*(*uint8)(unsafe.Add(mBase, uint32(v110))) = uint8(v113)
	if v113 == int32(0) {
		v124 = v109
		v125 = v110
		goto L7
	} else {
		goto L31
	}
L30:
	;
	v124 = v120
	v125 = v118
	goto L7
L31:
	;
	v117 = int32(1)
	v118 = v110 + v117
	v120 = v109 + v117
	v122 = v111 - v117
	if v122 != 0 {
		v109 = v120
		v110 = v118
		v111 = v122
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v140 = v18
	v144 = v138
	goto L36
L34:
	;
	goto L35
L35:
	;
	v174 = int32(0)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l2)+268))
	if v181 <= v174 {
		v334 = v174
		goto L46
	} else {
		goto L47
	}
L36:
	;
	if base.Ui32((v144-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L35
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v156)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)))
	if v158 != 0 {
		v140 = v140 + int32(1)
		v144 = v158
		goto L36
	} else {
		goto L42
	}
L39:
	;
	v154 = v144 - int32(32)
	goto L41
L40:
	;
	v154 = v144
	goto L41
L41:
	;
	v156 = v154 & int32(255)
	goto L38
L42:
	;
	goto L37
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L83
	} else {
		goto L120
	}
L44:
	;
	m.G0 = v10 + int32(288)
	return v517
L45:
	;
	if v334 != 0 {
		goto L80
	} else {
		goto L81
	}
L46:
	;
	goto L45
L47:
	;
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v10+int32(280))))
	v189 = int32(0)
	goto L48
L48:
	;
	v200 = v189 + (l2 + int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_0))
	v201 = F_strcmp(m, v10+int32(16), v200)
	mBase = m.M
	if v201 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l2)+260))
	if v207 <= int32(0) {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v202 = F_strlen(m, v200)
	mBase = m.M
	v205 = v202 + v189 + int32(1)
	if v205 < v181 {
		v189 = v205
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	v334 = v174
	goto L46
L54:
	;
	v252 = l2 + int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_1)
	v254 = l2 + int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_2)
	v255 = v244
	goto L68
L55:
	;
	v244 = int32(0)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v214 = v207
	v219 = int32(0)
	goto L58
L58:
	;
	v227 = int32(1)
	v228 = (v214 + v219) >> (uint(v227) % 32)
	v234 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(280)+v228<<(uint(int32(3))%32))))
	v235 = base.B2i32(v186 < v234)
	if v186 < v234 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v244 = v236
	goto L54
L60:
	;
	v236 = v219
	goto L62
L61:
	;
	v236 = v228 + v227
	goto L62
L62:
	;
	if v186 < v234 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v237 = v228
	goto L65
L64:
	;
	v237 = v214
	goto L65
L65:
	;
	if v236 < v237 {
		v214 = v237
		v219 = v236
		goto L58
	} else {
		goto L66
	}
L66:
	;
	goto L59
L67:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v313)))
	*(*int32)(unsafe.Add(mBase, uint32(v10+int32(12)))) = v319
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v321
	v334 = int32(1)
	goto L46
L68:
	;
	if int32(0) < v255 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l2)+uint32(_c_F_DetermineTimeZoneAbbrevOffsetTS[0])))
	v281 = v254 + v278<<(uint(int32(4))%32)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+8))
	if v282 == v189 {
		v313 = v281
		goto L67
	} else {
		goto L74
	}
L70:
	;
	v270 = v255 - int32(1)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252+v270))))
	v275 = v254 + v272<<(uint(int32(4))%32)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	if v276 != v189 {
		v255 = v270
		goto L68
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	v313 = v275
	goto L67
L74:
	;
	if v207 <= v244 {
		v334 = v174
		goto L46
	} else {
		goto L75
	}
L75:
	;
	v290 = v244
	goto L76
L76:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290+v252))))
	v301 = v254 + v298<<(uint(int32(4))%32)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
	if v302 == v189 {
		v313 = v301
		goto L67
	} else {
		goto L78
	}
L77:
	;
	v334 = v174
	goto L46
L78:
	;
	v305 = v290 + int32(1)
	if v207 != v305 {
		v290 = v305
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v517 = int32(0) - v337
	goto L44
L81:
	;
	goto L82
L82:
	;
	v342 = v10 + int32(16)
	v346 = F_timestamp2tm(m, l0, v10+int32(12), v342, v10+int32(8), int32(0), l2)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	return int32(0)
L84:
	;
	if v346 != 0 {
		goto L43
	} else {
		goto L85
	}
L85:
	;
	v351 = v10 + int32(280)
	v359 = m.G0
	v361 = v359 - int32(32)
	m.G0 = v361
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v342)+20))
	if v363 <= int32(-4713) {
		goto L90
	} else {
		goto L91
	}
L86:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v515
	v517 = v511
	goto L44
L87:
	;
	m.G0 = v361 + int32(32)
	goto L86
L88:
	;
	v498 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v342)+32)) = v498
	*(*int64)(unsafe.Add(mBase, uint32(v351))) = int64(0)
	v511 = v498
	goto L87
L89:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	v383 = int32(60)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v342)+12))
	v394 = base.B2i32(int32(2) < v379)
	if int32(2) < v379 {
		goto L100
	} else {
		goto L101
	}
L90:
	;
	if v363 != int32(-4713) {
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	if v363 <= int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_3) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v342)+16))
	if int32(10) < v368 {
		v379 = v368
		goto L89
	} else {
		goto L94
	}
L94:
	;
	goto L88
L95:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v342)+16))
	v379 = v373
	goto L89
L96:
	;
	goto L97
L97:
	;
	if v363 != int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_4) {
		goto L88
	} else {
		goto L98
	}
L98:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v342)+16))
	if int32(5) < v376 {
		goto L88
	} else {
		goto L99
	}
L99:
	;
	v379 = v376
	goto L89
L100:
	;
	v395 = int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_5)
	goto L102
L101:
	;
	v395 = int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_6)
	goto L102
L102:
	;
	v396 = v395 + v363
	v404 = base.I32_div_u_s(v396, int32(100))
	v407 = base.I32_div_u_s(v396, int32(400))
	if int32(2) < v379 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v411 = int32(1)
	goto L105
L104:
	;
	v411 = int32(13)
	goto L105
L105:
	;
	v416 = base.I32_div_s((v411+v379)*int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_7), int32(256))
	v419 = v390 + v396*int32(365) + int32(base.Ui32(v396)>>(uint(int32(2))%32)) - v404 + v407 + v416 - int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_8)
	v423 = base.I64_extend_i32_s(v380+(v381+v382*v383)*v383) + base.I64_extend_i32_s(v419)*int64(86400)
	if base.B2i32(int32(0) < v419)&base.B2i32(v423 < int64(0)) != 0 {
		goto L88
	} else {
		goto L106
	}
L106:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v361)+24)) = v423 - int64(86400)
	v442 = F_pg_next_dst_boundary(m, v361+int32(24), v361+int32(12), v361+int32(4), v361+int32(16), v361+int32(8), v361, l2)
	mBase = m.M
	if v442 < int32(0) {
		goto L88
	} else {
		goto L107
	}
L107:
	;
	if v442 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+32)) = v447
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v351))) = v423 - base.I64_extend_i32_s(v449)
	v511 = int32(0) - v449
	goto L87
L109:
	;
	goto L110
L110:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	v457 = v423 - base.I64_extend_i32_s(v455)
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v361)+16))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	v462 = v423 - base.I64_extend_i32_s(v460)
	if base.B2i32(v458 <= v457)|base.B2i32(v458 <= v462) == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+32)) = v467
	*(*int64)(unsafe.Add(mBase, uint32(v351))) = v457
	v511 = int32(0) - v455
	goto L87
L112:
	;
	goto L113
L113:
	;
	if base.B2i32(v462 < v458)|base.B2i32(v457 <= v458) == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+32)) = v477
	*(*int64)(unsafe.Add(mBase, uint32(v351))) = v462
	v511 = int32(0) - v460
	goto L87
L115:
	;
	goto L116
L116:
	;
	if v455 < v460 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+32)) = v483
	*(*int64)(unsafe.Add(mBase, uint32(v351))) = v457
	v511 = int32(0) - v455
	goto L87
L118:
	;
	goto L119
L119:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+32)) = v488
	*(*int64)(unsafe.Add(mBase, uint32(v351))) = v462
	v511 = int32(0) - v460
	goto L87
L120:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L83
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_9), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L83
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_10), int32(1826), int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_11))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L83
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_datan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v18 int64
	_ = v18
	var v23 int32
	_ = v23
	var v33 float64
	_ = v33
	var v39 float64
	_ = v39
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v87 float64
	_ = v87
	var v104 float64
	_ = v104
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v115 float64
	_ = v115
	var v118 float64
	_ = v118
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v133 float64
	_ = v133
	var v140 int32
	_ = v140
	var v141 float64
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v6)&int64(9223372036854775807)) {
		v141 = math.Float64frombits(uint64(0x7ff8000000000000))
		v142 = F_Float8GetDatum(m, v141)
		mBase = m.M
		v143 = m.ExcPending
		if v143 != 0 {
			return int32(0)
		} else {
			return v142
		}
	} else {
		v18 = base.I64_reinterpret_f64(v6)
		v23 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1141899264)) <= base.Ui32(v23) {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v6)&int64(9223372036854775807)) {
				v33 = v6
			} else {
				v33 = base.F64_copysign(float64(1.5707963267948966), v6)
			}
			v133 = v33
		} else {
			if base.Ui32(v23) <= base.Ui32(int32(1071382527)) {
				if base.Ui32(int32(1044381696)) <= base.Ui32(v23) {
					v70 = v6
					v71 = int32(-1)
					v72 = base.F64_mul(v70, v70)
					v73 = base.F64_mul(v72, v72)
					v87 = base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
					v104 = base.F64_mul(v72, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
					if base.Ui32(v23) <= base.Ui32(int32(1071382527)) {
						v133 = base.F64_sub(v70, base.F64_mul(v70, base.F64_add(v87, v104)))
					} else {
						v111 = v71 << (uint(int32(3)) % 32)
						v112 = *(*float64)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_datan[0])))
						v115 = *(*float64)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_datan[1])))
						v118 = base.F64_sub(v112, base.F64_sub(base.F64_sub(base.F64_mul(v70, base.F64_add(v87, v104)), v115), v70))
						if v18 < int64(0) {
							v122 = base.F64_neg(v118)
						} else {
							v122 = v118
						}
						v123 = v122
						v133 = v123
					}
				} else {
					v123 = v6
					v133 = v123
				}
			} else {
				v39 = base.F64_abs(v6)
				if base.Ui32(v23) <= base.Ui32(int32(1072889855)) {
					if base.Ui32(v23) <= base.Ui32(int32(1072037887)) {
						v70 = base.F64_div(base.F64_add(base.F64_add(v39, v39), float64(-1)), base.F64_add(v39, float64(2)))
						v71 = int32(0)
					} else {
						v70 = base.F64_div(base.F64_add(v39, float64(-1)), base.F64_add(v39, float64(1)))
						v71 = int32(1)
					}
				} else {
					if base.Ui32(v23) <= base.Ui32(int32(1073971199)) {
						v70 = base.F64_div(base.F64_add(v39, float64(-1.5)), base.F64_add(base.F64_mul(v39, float64(1.5)), float64(1)))
						v71 = int32(2)
					} else {
						v70 = base.F64_div(float64(-1), v39)
						v71 = int32(3)
					}
				}
				v72 = base.F64_mul(v70, v70)
				v73 = base.F64_mul(v72, v72)
				v87 = base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
				v104 = base.F64_mul(v72, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, base.F64_add(base.F64_mul(v73, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
				if base.Ui32(v23) <= base.Ui32(int32(1071382527)) {
					v133 = base.F64_sub(v70, base.F64_mul(v70, base.F64_add(v87, v104)))
				} else {
					v111 = v71 << (uint(int32(3)) % 32)
					v112 = *(*float64)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_datan[0])))
					v115 = *(*float64)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_datan[1])))
					v118 = base.F64_sub(v112, base.F64_sub(base.F64_sub(base.F64_mul(v70, base.F64_add(v87, v104)), v115), v70))
					if v18 < int64(0) {
						v122 = base.F64_neg(v118)
					} else {
						v122 = v118
					}
					v123 = v122
					v133 = v123
				}
			}
		}
		if base.F64_ne(base.F64_abs(v133), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v141 = v133
			v142 = F_Float8GetDatum(m, v141)
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return int32(0)
			} else {
				return v142
			}
		} else {
			F_float_overflow_error(m)
			mBase = m.M
			v140 = m.ExcPending
			if v140 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_datan2d(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 float64
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v18 int32
	_ = v18
	var v19 float64
	_ = v19
	var v26 int32
	_ = v26
	var v37 int64
	_ = v37
	var v50 int64
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 float64
	_ = v60
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v96 float64
	_ = v96
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v138 float64
	_ = v138
	var v142 float64
	_ = v142
	var v145 float64
	_ = v145
	var v152 int32
	_ = v152
	var v153 float64
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = math.Float64frombits(uint64(0x7ff8000000000000))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
		v153 = v10
		v155 = F_Float8GetDatum(m, v153)
		mBase = m.M
		v156 = m.ExcPending
		if v156 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v155
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v19)&int64(9223372036854775807)) {
			v153 = v10
			v155 = F_Float8GetDatum(m, v153)
			mBase = m.M
			v156 = m.ExcPending
			if v156 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return v155
			}
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_datan2d[0])))
			if v26 == int32(0) {
				F_init_degree_constants(m)
				mBase = m.M
			} else {
			}
			v37 = int64(9223372036854775807)
			if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v12)&v37) < base.Ui64(int64(9218868437227405313)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v19)&v37) <= base.Ui64(int64(9218868437227405312))) == int32(0) {
				v138 = base.F64_add(v12, v19)
			} else {
				v50 = base.I64_reinterpret_f64(v19)
				v53 = base.I32_wrap_i64(int64(base.Ui64(v50) >> (uint(int64(32)) % 64)))
				v56 = base.I32_wrap_i64(v50)
				if v53-int32(1072693248)|v56 == int32(0) {
					v60 = F_atan(m, v12)
					mBase = m.M
					v138 = v60
				} else {
					v64 = int32(base.Ui32(v53)>>(uint(int32(30))%32)) & int32(2)
					v65 = base.I64_reinterpret_f64(v12)
					v69 = v64 | base.I32_wrap_i64(int64(base.Ui64(v65)>>(uint(int64(63))%64)))
					v74 = base.I32_wrap_i64(int64(base.Ui64(v65)>>(uint(int64(32))%64))) & int32(2147483647)
					if v74|base.I32_wrap_i64(v65) == int32(0) {
						switch v69 - int32(2) {
						case 0:
							v138 = float64(3.141592653589793)
						case 1:
							v138 = float64(-3.141592653589793)
						default:
							v129 = v12
							v138 = v129
						}
					} else {
						v84 = v53 & int32(2147483647)
						if v84|v56 == int32(0) {
							v138 = base.F64_copysign(float64(1.5707963267948966), v12)
						} else {
							if v84 == int32(2146435072) {
								if v74 != int32(2146435072) {
									v128 = *(*float64)(unsafe.Add(mBase, uint32(v69<<(uint(int32(3))%32))+uint32(_c_F_datan2d[1])))
									v129 = v128
									v138 = v129
								} else {
									v96 = *(*float64)(unsafe.Add(mBase, uint32(v69<<(uint(int32(3))%32))+uint32(_c_F_datan2d[2])))
									v138 = v96
								}
							} else {
								if base.B2i32(v74 != int32(2146435072))&base.B2i32(base.Ui32(v74) <= base.Ui32(v84+int32(67108864))) == int32(0) {
									v138 = base.F64_copysign(float64(1.5707963267948966), v12)
								} else {
									if v64 != 0 {
										if base.Ui32(v74+int32(67108864)) < base.Ui32(v84) {
											v114 = float64(0)
										} else {
											v113 = F_atan(m, base.F64_abs(base.F64_div(v12, v19)))
											mBase = m.M
											v114 = v113
										}
									} else {
										v113 = F_atan(m, base.F64_abs(base.F64_div(v12, v19)))
										mBase = m.M
										v114 = v113
									}
									switch v69 - int32(1) {
									case 0:
										v138 = base.F64_neg(v114)
									case 1:
										v138 = base.F64_sub(float64(3.141592653589793), base.F64_add(v114, float64(-1.2246467991473532e-16)))
									case 2:
										v138 = base.F64_add(base.F64_add(v114, float64(-1.2246467991473532e-16)), float64(-3.141592653589793))
									default:
										v129 = v114
										v138 = v129
									}
								}
							}
						}
					}
				}
			}
			*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v138
			v142 = *(*float64)(unsafe.Add(mBase, _c_F_datan2d[3]))
			v145 = base.F64_mul(base.F64_div(v138, v142), float64(45))
			if base.F64_ne(base.F64_abs(v145), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v153 = v145
				v155 = F_Float8GetDatum(m, v153)
				mBase = m.M
				v156 = m.ExcPending
				if v156 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v155
				}
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v152 = m.ExcPending
				if v152 != 0 {
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
func F_date2isoyear(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	v9 = base.B2i32(int32(2) < l1)
	if int32(2) < l1 {
		v10 = int32(_a_F_date2isoyear_0)
	} else {
		v10 = int32(_a_F_date2isoyear_1)
	}
	v11 = v10 + l0
	v16 = base.I32_div_s(v11, int32(4))
	v19 = base.I32_div_s(v11, int32(-100))
	v22 = base.I32_div_s(v11, int32(400))
	if int32(2) < l1 {
		v26 = int32(1)
	} else {
		v26 = int32(13)
	}
	v31 = base.I32_div_s((v26+l1)*int32(_a_F_date2isoyear_2), int32(256))
	v34 = l2 + v11*int32(365) + v16 + v19 + v22 + v31 - int32(_a_F_date2isoyear_3)
	v43 = int32(_a_F_date2isoyear_1) + l0
	v48 = base.I32_div_s(v43, int32(4))
	v51 = base.I32_div_s(v43, int32(-100))
	v54 = base.I32_div_s(v43, int32(400))
	v63 = base.I32_div_s(int32(_a_F_date2isoyear_4), int32(256))
	v66 = int32(4) + v43*int32(365) + v48 + v51 + v54 + v63 - int32(_a_F_date2isoyear_3)
	v67 = int32(1)
	v71 = int32(7)
	v72 = base.I32_rem_s(v66-v67+v67, v71)
	if v72 < int32(0) {
		v77 = v72 + v71
	} else {
		v77 = v72
	}
	if v34 < v66-v77 {
		v81 = l0 - int32(1)
		v90 = int32(_a_F_date2isoyear_1) + v81
		v95 = base.I32_div_s(v90, int32(4))
		v98 = base.I32_div_s(v90, int32(-100))
		v101 = base.I32_div_s(v90, int32(400))
		v110 = base.I32_div_s(int32(_a_F_date2isoyear_4), int32(256))
		v113 = int32(4) + v90*int32(365) + v95 + v98 + v101 + v110 - int32(_a_F_date2isoyear_3)
		v114 = int32(1)
		v118 = int32(7)
		v119 = base.I32_rem_s(v113-v114+v114, v118)
		if v119 < int32(0) {
			v124 = v119 + v118
		} else {
			v124 = v119
		}
		v125 = v81
		v126 = v113
		v127 = v124
	} else {
		v125 = l0
		v126 = v66
		v127 = v77
	}
	if int32(357) <= v127+v34-v126 {
		v142 = v125 + int32(_a_F_date2isoyear_0)
		v147 = base.I32_div_s(v142, int32(4))
		v150 = base.I32_div_s(v142, int32(-100))
		v153 = base.I32_div_s(v142, int32(400))
		v162 = base.I32_div_s(int32(_a_F_date2isoyear_4), int32(256))
		v165 = int32(4) + v142*int32(365) + v147 + v150 + v153 + v162 - int32(_a_F_date2isoyear_3)
		v166 = int32(1)
		v170 = int32(7)
		v171 = base.I32_rem_s(v165-v166+v166, v170)
		if v171 < int32(0) {
			v176 = v171 + v170
		} else {
			v176 = v171
		}
		if v34 < v165-v176 {
			v179 = v125
		} else {
			v179 = v125 + int32(1)
		}
		v182 = v179
	} else {
		v182 = v125
	}
	return v182
}
func F_date2j(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v8 = base.B2i32(int32(2) < l1)
	if int32(2) < l1 {
		v9 = int32(_a_F_date2j_0)
	} else {
		v9 = int32(_a_F_date2j_1)
	}
	v10 = v9 + l0
	v15 = base.I32_div_s(v10, int32(4))
	v18 = base.I32_div_s(v10, int32(-100))
	v21 = base.I32_div_s(v10, int32(400))
	if int32(2) < l1 {
		v25 = int32(1)
	} else {
		v25 = int32(13)
	}
	v30 = base.I32_div_s((v25+l1)*int32(_a_F_date2j_2), int32(256))
	return l2 + v10*int32(365) + v15 + v18 + v21 + v30 - int32(_a_F_date2j_3)
}
func F_db_dir_size(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v92 int64
	_ = v92
	v5 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(2176)
	m.G0 = v8
	v10 = F_AllocateDir(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v14 = F_ReadDir(m, v10, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v92 = v5
	goto L5
L5:
	;
	m.G0 = v8 + int32(2176)
	return v92
L6:
	;
	if v14 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v18 = v14
	v20 = v5
	goto L10
L8:
	;
	v85 = v5
	goto L9
L9:
	;
	F_FreeDir(m, v10)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L34
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_db_dir_size[0]))
	if v22 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v85 = v78
	goto L9
L12:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+19)))
	if v25 != int32(46) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L14
L16:
	;
	v79 = F_ReadDir(m, v10, l0)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L32
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v18 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	v42 = v8 + int32(128)
	v47 = F_pg_snprintf(m, v42, int32(2048), int32(_a_F_db_dir_size_0), v8+int32(16))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L22
	}
L18:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)))
	if v28 == int32(0) {
		v78 = v20
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+20)))
	if v31 != int32(46) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+21)))
	if v34 == int32(0) {
		v78 = v20
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v53 = F___fstatat(m, int32(-100), v42, v8+int32(32), int32(0))
	mBase = m.M
	goto L23
L23:
	;
	if v53 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_db_dir_size[1]))
	if v57 == int32(44) {
		v78 = v20
		goto L16
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v8)+56))
	v78 = v75 + v20
	goto L16
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v42
	F_errmsg(m, int32(_a_F_db_dir_size_1), v8)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_db_dir_size_2), int32(105), int32(_a_F_db_dir_size_3))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	if v79 != 0 {
		v18 = v79
		v20 = v78
		goto L10
	} else {
		goto L33
	}
L33:
	;
	goto L11
L34:
	;
	v92 = v85
	goto L5
}
func F_db_encoding_convert(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = F_strlen(m, v5)
	mBase = m.M
	v7 = F_pg_any_to_server(m, v5, v6, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v7 != v9 {
			v13 = F_strlen(m, v7)
			mBase = m.M
			v15 = v13 + int32(1)
			v16 = F_emscripten_builtin_malloc(m, v15)
			mBase = m.M
			if v16 == int32(0) {
				v21 = int32(0)
			} else {
				v20 = F___memcpy(m, v16, v7, v15)
				mBase = m.M
				v21 = v20
			}
			if v21 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_errcode(m, int32(_a_F_db_encoding_convert_0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_db_encoding_convert_1), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_db_encoding_convert_2), int32(517), int32(_a_F_db_encoding_convert_3))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				F_emscripten_builtin_free(m, v9)
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21
				F_pfree(m, v7)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			return
		}
	}
}
func F_dccref_deletion_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v3 == int32(0) {
		return
	} else {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v6
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		v12 = v10 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = v12
		if v6 < v12 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
			F_MemoryContextDelete(m, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_dcot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v35 float64
	_ = v35
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		*(*int32)(unsafe.Add(mBase, _c_F_dcot[0])) = int32(0)
		if base.F64_eq(base.F64_abs(v4), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_dcot_0), int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_dcot_1), int32(1925), int32(_a_F_dcot_2))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v19 = m.G0
			v21 = v19 - int32(16)
			m.G0 = v21
			v28 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v4))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(v28) <= base.Ui32(int32(1072243195)) {
				if base.Ui32(v28) < base.Ui32(int32(1044381696)) {
					v45 = v4
				} else {
					v35 = F___tan(m, v4, float64(0), int32(0))
					mBase = m.M
					v45 = v35
				}
			} else {
				if base.Ui32(int32(2146435072)) <= base.Ui32(v28) {
					v45 = base.F64_sub(v4, v4)
				} else {
					v39 = F___rem_pio2(m, v4, v21)
					mBase = m.M
					v40 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
					v41 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
					v44 = F___tan(m, v40, v41, v39&int32(1))
					mBase = m.M
					v45 = v44
				}
			}
			m.G0 = v21 + int32(16)
			v52 = base.F64_div(float64(1), v45)
			v53 = F_Float8GetDatum(m, v52)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				return v53
			}
		}
	} else {
		v52 = math.Float64frombits(uint64(0x7ff8000000000000))
		v53 = F_Float8GetDatum(m, v52)
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			return v53
		}
	}
}
func F_delvacuum_desc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v128 int32
	_ = v128
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	F_appendStringInfoString(m, l0, int32(_a_F_delvacuum_desc_0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_array_desc(m, l0, l1, int32(2), l2, int32(243), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_delvacuum_desc_1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l3 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = int32(1)
	v30 = l1 + l2<<(uint(v26)%32)
	v35 = v30 + l3<<(uint(v26)%32)
	v39 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	F_appendStringInfoChar(m, l0, int32(93))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L28
	}
L8:
	;
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v39<<(uint(int32(1))%32)))))
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v47
	F_appendStringInfo(m, l0, int32(_a_F_delvacuum_desc_2), v13+int32(16))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	if v56 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v62 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_delvacuum_desc_3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L22
	}
L14:
	;
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(2)+v62<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v73
	F_appendStringInfo(m, l0, int32(_a_F_delvacuum_desc_4), v13)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	if v62 < v78-int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_delvacuum_desc_5))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	v86 = v78
	goto L19
L19:
	;
	v88 = v62 + int32(1)
	if base.Ui32(v88) < base.Ui32(v86) {
		v62 = v88
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	v86 = v85
	goto L19
L21:
	;
	goto L15
L22:
	;
	if v39 < l3-v26 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_delvacuum_desc_5))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35))))
	v108 = int32(1)
	v114 = v39 + v108
	if v114 != l3 {
		v35 = v35 + v107<<(uint(v108)%32) + int32(2)
		v39 = v114
		goto L8
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	goto L9
L28:
	;
	m.G0 = v13 + int32(32)
	return
}
func F_desc_recompress_leaf(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v13
	F_appendStringInfo(m, l0, int32(_a_F_desc_recompress_leaf_0), v11+int32(80))
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
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if v20 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v11 + int32(96)
	return
L4:
	;
	v28 = l1 + int32(2)
	v32 = int32(0)
	goto L5
L5:
	;
	v33 = int32(2)
	v34 = v28 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v36&int32(254) == v33 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L3
L7:
	;
	v95 = v32 + int32(1)
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	if base.Ui32(v95) < base.Ui32(v96) {
		v28 = v93
		v32 = v95
		goto L5
	} else {
		goto L23
	}
L8:
	;
	switch v36 - int32(1) {
	case 0:
		goto L15
	case 1:
		goto L18
	case 2:
		goto L17
	default:
		goto L16
	}
L9:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+8)))
	v66 = v34 + ((v41+int32(1))&int32(_a_F_desc_recompress_leaf_1)+int32(9))&int32(_a_F_desc_recompress_leaf_2)
	goto L8
L10:
	;
	goto L11
L11:
	;
	if v36 != int32(4) {
		v66 = v34
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v35
	F_appendStringInfo(m, l0, int32(_a_F_desc_recompress_leaf_3), v11-int32(-64))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v93 = v28 + v53*int32(6) + int32(4)
	goto L7
L14:
	;
	v93 = v66
	goto L7
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v35
	F_appendStringInfo(m, l0, int32(_a_F_desc_recompress_leaf_4), v11+int32(32))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L22
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v35
	F_appendStringInfo(m, l0, int32(_a_F_desc_recompress_leaf_5), v11)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L21
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v35
	F_appendStringInfo(m, l0, int32(_a_F_desc_recompress_leaf_6), v11+int32(16))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v35
	F_appendStringInfo(m, l0, int32(_a_F_desc_recompress_leaf_7), v11+int32(48))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L14
L20:
	;
	goto L14
L21:
	;
	goto L3
L22:
	;
	goto L14
L23:
	;
	goto L6
}
func F_detoast_attr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == int32(1) {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v15 = v13 - int32(1)
		if v15 != 0 {
			if v15 != int32(17) {
				if v13&int32(254) != int32(2) {
					v147 = int32(1)
					v150 = int32(base.Ui32(v147) >> (uint(int32(1)) % 32))
					v152 = v150 + int32(3)
					v153 = F_palloc(m, v152)
					mBase = m.M
					v154 = m.ExcPending
					if v154 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v153))) = v152 << (uint(int32(2)) % 32)
						v159 = v150 - int32(1)
						if v159 == int32(0) {
							v168 = v153
						} else {
							base.MemoryCopy(m, v153+int32(4), l0+int32(1), v159)
							v168 = v153
						}
						m.G0 = v8 + int32(32)
						return v168
					}
				} else {
					v117 = F_detoast_external_attr(m, l0)
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						v168 = v117
						m.G0 = v8 + int32(32)
						return v168
					}
				}
			} else {
				v18 = F_toast_fetch_datum(m, l0)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
					if v22&int32(3) != int32(2) {
						v168 = v18
						m.G0 = v8 + int32(32)
						return v168
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
						v29 = int32(base.Ui32(v27) >> (uint(int32(30)) % 32))
						switch v29 {
						case 0:
							v47 = F_pglz_decompress_datum(m, v18)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v18)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v168 = v47
									m.G0 = v8 + int32(32)
									return v168
								}
							}
						case 1:
							v30 = F_lz4_decompress_datum(m)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v18)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									v168 = v30
									m.G0 = v8 + int32(32)
									return v168
								}
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v29
								F_errmsg_internal(m, int32(_a_F_detoast_attr_0), v8)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_detoast_attr_1), int32(489), int32(_a_F_detoast_attr_2))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
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
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
			v52 = F_detoast_attr(m, v51)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				if v51 != v52 {
					v168 = v52
					m.G0 = v8 + int32(32)
					return v168
				} else {
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
					if v55 == int32(1) {
						v59 = int32(18)
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
						if v61 == v59 {
							v64 = v59
						} else {
							v64 = int32(2)
						}
						if base.Ui32((v61-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v71 = int32(6)
						} else {
							v71 = v64
						}
						v80 = v71
					} else {
						v72 = int32(1)
						if v55&v72 != 0 {
							v80 = int32(base.Ui32(v55) >> (uint(v72) % 32))
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							v80 = int32(base.Ui32(v76) >> (uint(int32(2)) % 32))
						}
					}
					v81 = F_palloc(m, v80)
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
						if v83 == int32(1) {
							v87 = int32(18)
							v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
							if v89 == v87 {
								v92 = v87
							} else {
								v92 = int32(2)
							}
							if base.Ui32((v89-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v99 = int32(6)
							} else {
								v99 = v92
							}
							v108 = v99
						} else {
							v100 = int32(1)
							if v83&v100 != 0 {
								v108 = int32(base.Ui32(v83) >> (uint(v100) % 32))
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
								v108 = int32(base.Ui32(v104) >> (uint(int32(2)) % 32))
							}
						}
						if v108 == int32(0) {
							v168 = v81
						} else {
							base.MemoryCopy(m, v81, v52, v108)
							v168 = v81
						}
						m.G0 = v8 + int32(32)
						return v168
					}
				}
			}
		}
	} else {
		if v10&int32(3) == int32(2) {
			v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v125 = int32(base.Ui32(v123) >> (uint(int32(30)) % 32))
			switch v125 {
			case 0:
				v126 = F_pglz_decompress_datum(m, l0)
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return int32(0)
				} else {
					v168 = v126
					m.G0 = v8 + int32(32)
					return v168
				}
			case 1:
				v128 = F_lz4_decompress_datum(m)
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return int32(0)
				} else {
					v168 = v128
					m.G0 = v8 + int32(32)
					return v168
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v125
					F_errmsg_internal(m, int32(_a_F_detoast_attr_0), v8+int32(16))
					mBase = m.M
					v139 = m.ExcPending
					if v139 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_detoast_attr_1), int32(489), int32(_a_F_detoast_attr_2))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			if v10&int32(1) != 0 {
				v147 = v10
				v150 = int32(base.Ui32(v147) >> (uint(int32(1)) % 32))
				v152 = v150 + int32(3)
				v153 = F_palloc(m, v152)
				mBase = m.M
				v154 = m.ExcPending
				if v154 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v153))) = v152 << (uint(int32(2)) % 32)
					v159 = v150 - int32(1)
					if v159 == int32(0) {
						v168 = v153
					} else {
						base.MemoryCopy(m, v153+int32(4), l0+int32(1), v159)
						v168 = v153
					}
					m.G0 = v8 + int32(32)
					return v168
				}
			} else {
				v168 = l0
				m.G0 = v8 + int32(32)
				return v168
			}
		}
	}
}
func F_dintdict_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_palloc0(m, int32(8))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)) = uint16(v18)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(6)
	if v11 == v18 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L63
	}
L4:
	;
	m.G0 = v9 + int32(16)
	return v14
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v24 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v27 = int32(0)
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v27<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v39 = int32(_a_F_dintdict_init_0)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dintdict_init[0])))
	if base.B2i32(v42 == int32(0))|base.B2i32(v42 != v45) != 0 {
		v63 = v42
		v64 = v45
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L4
L9:
	;
	v201 = v27 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v201 < v202 {
		v27 = v201
		goto L7
	} else {
		goto L62
	}
L10:
	;
	if v63-v64 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	v48 = v38
	v49 = v39
	goto L13
L13:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v53 == int32(0) {
		v63 = v53
		v64 = v52
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v63 = v53
	v64 = v52
	goto L11
L15:
	;
	v56 = int32(1)
	if v53 == v52 {
		v48 = v48 + v56
		v49 = v49 + v56
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v68 = F_defGetString(m, v37)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v137 = int32(_a_F_dintdict_init_1)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dintdict_init[1])))
	if base.B2i32(v140 == int32(0))|base.B2i32(v140 != v143) != 0 {
		v161 = v140
		v162 = v143
		goto L43
	} else {
		goto L44
	}
L20:
	;
	v73 = v68
	goto L22
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v117
	if int32(0) < v117 {
		goto L9
	} else {
		goto L37
	}
L22:
	;
	v78 = v73 + int32(1)
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73))))
	v80 = F___isspace(m, v79)
	mBase = m.M
	if v80 != 0 {
		v73 = v78
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v81 = int32(1)
	switch v79&int32(255) - int32(43) {
	case 0:
		v87 = v81
		goto L26
	default:
		v89 = v79
		v90 = v73
		v91 = v81
		goto L25
	case 2:
		goto L27
	}
L24:
	;
	goto L23
L25:
	;
	v92 = int32(0)
	v94 = v89 - int32(48)
	if base.Ui32(v94) <= base.Ui32(int32(9)) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v88 = int32(*(*int8)(unsafe.Add(mBase, uint32(v78))))
	v89 = v88
	v90 = v78
	v91 = v87
	goto L25
L27:
	;
	v87 = int32(0)
	goto L26
L28:
	;
	v97 = v92
	v98 = v94
	v99 = v90
	goto L31
L29:
	;
	v111 = v92
	goto L30
L30:
	;
	if v91 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v101 = int32(10)
	v103 = v97*v101 - v98
	v104 = int32(*(*int8)(unsafe.Add(mBase, uint32(v99)+1)))
	v108 = v104 - int32(48)
	if base.Ui32(v108) < base.Ui32(v101) {
		v97 = v103
		v98 = v108
		v99 = v99 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v111 = v103
	goto L30
L33:
	;
	goto L32
L34:
	;
	v117 = int32(0) - v111
	goto L36
L35:
	;
	v117 = v111
	goto L36
L36:
	;
	goto L21
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(_a_F_dintdict_init_2), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_dintdict_init_3), int32(57), int32(_a_F_dintdict_init_4))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	if v161-v162 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	goto L42
L44:
	;
	v146 = v38
	v147 = v137
	goto L45
L45:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
	if v151 == int32(0) {
		v161 = v151
		v162 = v150
		goto L43
	} else {
		goto L47
	}
L46:
	;
	v161 = v151
	v162 = v150
	goto L43
L47:
	;
	v154 = int32(1)
	if v151 == v150 {
		v146 = v146 + v154
		v147 = v147 + v154
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v166 = F_defGetBoolean(m, v37)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v169 = int32(_a_F_dintdict_init_5)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dintdict_init[2])))
	if base.B2i32(v172 == int32(0))|base.B2i32(v172 != v175) != 0 {
		v193 = v172
		v194 = v175
		goto L54
	} else {
		goto L55
	}
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)) = uint8(v166)
	goto L9
L53:
	;
	if v193-v194 != 0 {
		goto L3
	} else {
		goto L60
	}
L54:
	;
	goto L53
L55:
	;
	v178 = v38
	v179 = v169
	goto L56
L56:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	if v183 == int32(0) {
		v193 = v183
		v194 = v182
		goto L54
	} else {
		goto L58
	}
L57:
	;
	v193 = v183
	v194 = v182
	goto L54
L58:
	;
	v186 = int32(1)
	if v183 == v182 {
		v178 = v178 + v186
		v179 = v179 + v186
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v196 = F_defGetBoolean(m, v37)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)) = uint8(v196)
	goto L9
L62:
	;
	goto L8
L63:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v221
	F_errmsg(m, int32(_a_F_dintdict_init_6), v9)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_dintdict_init_3), int32(72), int32(_a_F_dintdict_init_4))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_directory_is_empty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v2 = int32(0)
	v5 = F_AllocateDir(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_FreeDir(m, v5)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L18
	}
L2:
	;
	return int32(0)
L3:
	;
	v9 = F_ReadDir(m, v5, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v12 = v9
	goto L8
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	goto L1
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+19)))
	if v15 != int32(46) {
		v33 = v2
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)))
	if v18 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)))
	if v19 != int32(46) {
		v33 = v2
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v23 = F_ReadDir(m, v5, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L16
	}
L14:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)))
	if v22 != 0 {
		v33 = v2
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	if v23 != 0 {
		v12 = v23
		goto L8
	} else {
		goto L17
	}
L17:
	;
	goto L9
L18:
	;
	return v33
}
func F_do_putc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	v1 = l0
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if int32(0) <= v6 {
		if v6 == int32(0) {
			v31 = l1 + int32(76)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
			if v32 != 0 {
				v34 = v32
			} else {
				v34 = int32(1073741823)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v31))) = v34
			v37 = v1 & int32(255)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
			if v37 == v38 {
				F___overflow(m, l1, v37)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
					return
				}
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				if v40 == v41 {
					F___overflow(m, l1, v37)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v40 + int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v1)
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
					return
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_do_putc[0]))
			if v12 != v6&int32(1073741823) {
				v31 = l1 + int32(76)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				if v32 != 0 {
					v34 = v32
				} else {
					v34 = int32(1073741823)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v31))) = v34
				v37 = v1 & int32(255)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
				if v37 == v38 {
					F___overflow(m, l1, v37)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
						return
					}
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					if v40 == v41 {
						F___overflow(m, l1, v37)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v40 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v1)
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
						return
					}
				}
			} else {
				v17 = v1 & int32(255)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
				if v17 == v18 {
					F___overflow(m, l1, v17)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						return
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					if v20 == v21 {
						F___overflow(m, l1, v17)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v20 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v1)
						return
					}
				}
			}
		}
	} else {
		v17 = v1 & int32(255)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
		if v17 == v18 {
			F___overflow(m, l1, v17)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				return
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			if v20 == v21 {
				F___overflow(m, l1, v17)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v20 + int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v1)
				return
			}
		}
	}
}
func F_do_tzset(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_do_tzset[0]))
	if v2 == int32(0) {
		v7 = int32(_a_F_do_tzset_0)
		v8 = int32(_a_F_do_tzset_1)
		m.Env.X_tzset_js(m, int32(_a_F_do_tzset_2), int32(_a_F_do_tzset_3), v7, v8)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, _c_F_do_tzset[1])) = v8
		*(*int32)(unsafe.Add(mBase, _c_F_do_tzset[0])) = v7
	} else {
	}
	return
}
func F_dofindsubquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
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
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v309 int32
	_ = v309
	v5 = int32(0)
	F_check_stack_depth(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_dofindsubquery[0]))
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v18&v19 != v18 {
		v226 = l0
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+4)))
	if v235&int32(2) != 0 {
		v289 = v226
		goto L72
	} else {
		goto L73
	}
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v23 != v25 {
		v226 = l0
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v27&int32(2) != 0 {
		v226 = l0
		goto L7
	} else {
		goto L10
	}
L10:
	;
	if v23 == int32(2) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v215 != 0 {
		v226 = v215
		goto L7
	} else {
		goto L70
	}
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v32 != v33 {
		v226 = l0
		goto L7
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v194 != v195 {
		v226 = l0
		goto L7
	} else {
		goto L61
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v35 == v36 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v38 = F_QTNEq(m, l0, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if base.B2i32(v35 <= v36)|base.B2i32(v36 <= int32(0)) != 0 {
		v226 = l0
		goto L7
	} else {
		goto L27
	}
L19:
	;
	if v38 == int32(0) {
		v226 = l0
		goto L7
	} else {
		goto L20
	}
L20:
	;
	F_QTNFree(m, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if l2 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v54)
	v215 = v53
	goto L11
L23:
	;
	v53 = int32(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v47 = F_QTNCopy(m, l2)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v49 | int32(2)
	v53 = v47
	goto L22
L27:
	;
	v60 = F_palloc0(m, v35)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v62 <= int32(0) {
		v115 = v5
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v116 == v115 {
		goto L41
	} else {
		goto L42
	}
L30:
	;
	v65 = int32(0)
	v71 = v65
	v73 = v65
	v75 = v5
	goto L31
L31:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v76 <= v73 {
		v115 = v75
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v115 = v102
	goto L29
L33:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v79 = int32(2)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v71<<(uint(v79)%32))))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83+v73<<(uint(v79)%32))))
	v88 = F_QTNodeCompare(m, v82, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v104 = v71 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v104 < v105 {
		v71 = v104
		v73 = v101
		v75 = v102
		goto L31
	} else {
		goto L40
	}
L35:
	;
	if v88 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v93 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v71+v60))) = uint8(v93)
	v101 = v73 + v93
	v102 = v75 + v93
	goto L34
L37:
	;
	goto L38
L38:
	;
	if int32(0) <= v88 {
		v115 = v75
		goto L29
	} else {
		goto L39
	}
L39:
	;
	v101 = v73
	v102 = v75
	goto L34
L40:
	;
	goto L32
L41:
	;
	v118 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v118 < v119 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	F_pfree(m, v60)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L60
	}
L44:
	;
	v127 = int32(0)
	v128 = v118
	goto L47
L45:
	;
	v159 = v118
	goto L46
L46:
	;
	if l2 != 0 {
		goto L55
	} else {
		goto L56
	}
L47:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v127<<(uint(int32(2))%32))))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127+v60))))
	if v138 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v159 = v149
	goto L46
L49:
	;
	v151 = v127 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v151 < v152 {
		v127 = v151
		v128 = v149
		goto L47
	} else {
		goto L54
	}
L50:
	;
	F_QTNFree(m, v136)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132+v128<<(uint(int32(2))%32)))) = v136
	v149 = v128 + int32(1)
	goto L49
L53:
	;
	v149 = v128
	goto L49
L54:
	;
	goto L48
L55:
	;
	v163 = F_QTNCopy(m, l2)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	v177 = v159
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v177
	F_QTNSort(m, l0)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v166 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v165 | v166
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v169+v159<<(uint(v166)%32)))) = v163
	v177 = v159 + int32(1)
	goto L57
L59:
	;
	v181 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v181)
	goto L43
L60:
	;
	v215 = l0
	goto L11
L61:
	;
	v197 = F_QTNEq(m, l0, l1)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if v197 == int32(0) {
		v226 = l0
		goto L7
	} else {
		goto L63
	}
L63:
	;
	F_QTNFree(m, l0)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if l2 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v213)
	v215 = v212
	goto L11
L66:
	;
	v212 = int32(0)
	goto L65
L67:
	;
	goto L68
L68:
	;
	v206 = F_QTNCopy(m, l2)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v208 | int32(2)
	v212 = v206
	goto L65
L70:
	;
	return int32(0)
L71:
	;
	F_QTNFree(m, v226)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L85
	}
L72:
	;
	return v289
L73:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v239 != int32(2) {
		v289 = v226
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
	if v242 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v226)+8)) = int32(0)
	goto L71
L76:
	;
	goto L77
L77:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
	v248 = int32(0)
	v254 = v248
	v255 = v247
	v256 = v248
	goto L78
L78:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v255+v256<<(uint(int32(2))%32))))
	v263 = F_dofindsubquery(m, v262, l1, l2, l3)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v226)+8)) = v275
	switch v275 {
	case 0:
		goto L71
	case 1:
		goto L82
	default:
		v289 = v226
		goto L72
	}
L80:
	;
	v266 = v254 << (uint(int32(2)) % 32)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v266+v267))) = v263
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v270+v266)))
	v275 = v254 + base.B2i32(v272 != int32(0))
	v277 = v256 + int32(1)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v226)+8))
	if v277 < v278 {
		v254 = v275
		v255 = v270
		v256 = v277
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+1)))
	if v282 == int32(1) {
		v289 = v226
		goto L72
	} else {
		goto L83
	}
L83:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	F_pfree(m, v226)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v289 = v286
	goto L72
L85:
	;
	return int32(0)
}
func F_dopr_outchmulti(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	v1 = l0
	if l1 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if int32(0) < l1 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v66 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.B2i32(v65 == v66)|base.B2i32(base.Ui32(v68) < base.Ui32(v65)) == v66 {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v13 = l1
	v16 = v11
	goto L7
L5:
	;
	goto L6
L6:
	;
	return
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v18 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L6
L9:
	;
	if int32(0) < v53 {
		v13 = v53
		v16 = v55
		goto L7
	} else {
		goto L30
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v30 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	v19 = v18 - v16
	if v19 <= int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v22 = v13
	goto L13
L13:
	;
	if base.Ui32(v22) < base.Ui32(v13) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v22 = v19
	goto L13
L15:
	;
	v24 = v22
	goto L17
L16:
	;
	v24 = v13
	goto L17
L17:
	;
	if v24 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	base.MemoryFill(m, v16, v1, v24)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v27 = v26 + v24
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v27
	v53 = v13 - v24
	v55 = v27
	goto L9
L21:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v33 + v13
	return
L22:
	;
	goto L23
L23:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v36 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v51
	v53 = v13
	v55 = v51
	goto L9
L25:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v16 == v37 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v40 = v16 - v37
	v41 = F_fwrite(m, v37, int32(1), v40, v30)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return
L28:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v41 + v43
	if v40 == v41 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v47 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v47)
	goto L24
L30:
	;
	goto L8
L31:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v73 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v96 = v68
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v96 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v96))) = uint8(v1)
	return
L34:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v76 + int32(1)
	return
L35:
	;
	goto L36
L36:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v80 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v96 = v95
	goto L33
L38:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v68 == v81 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v84 = v68 - v81
	v85 = F_fwrite(m, v81, int32(1), v84, v73)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L27
	} else {
		goto L40
	}
L40:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v85 + v87
	if v84 == v85 {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v91 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v91)
	goto L37
}
func F_dostr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	if l1 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if int32(0) < l1 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v69 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.B2i32(v68 == v69)|base.B2i32(base.Ui32(v71) < base.Ui32(v68)) == v69 {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v12 = l0
	v13 = l1
	v16 = v11
	goto L7
L5:
	;
	goto L6
L6:
	;
	return
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v18 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L6
L9:
	;
	if int32(0) < v55 {
		v12 = v54
		v13 = v55
		v16 = v57
		goto L7
	} else {
		goto L30
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v31 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	v19 = v18 - v16
	if v19 <= int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v22 = v13
	goto L13
L13:
	;
	if base.Ui32(v22) < base.Ui32(v13) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v22 = v19
	goto L13
L15:
	;
	v24 = v22
	goto L17
L16:
	;
	v24 = v13
	goto L17
L17:
	;
	if v24 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	base.MemoryCopy(m, v16, v12, v24)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v27 = v26 + v24
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v27
	v54 = v12 + v24
	v55 = v13 - v24
	v57 = v27
	goto L9
L21:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v34 + v13
	return
L22:
	;
	goto L23
L23:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v37 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v52
	v54 = v12
	v55 = v13
	v57 = v52
	goto L9
L25:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v16 == v38 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v41 = v16 - v38
	v42 = F_fwrite(m, v38, int32(1), v41, v31)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return
L28:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v42 + v44
	if v41 == v42 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v48 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v48)
	goto L24
L30:
	;
	goto L8
L31:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v76 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v100 = v71
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v100 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v67)
	return
L34:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v79 + int32(1)
	return
L35:
	;
	goto L36
L36:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v83 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v100 = v98
	goto L33
L38:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v71 == v84 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v87 = v71 - v84
	v88 = F_fwrite(m, v84, int32(1), v87, v76)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L27
	} else {
		goto L40
	}
L40:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v88 + v90
	if v87 == v88 {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v94 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v94)
	goto L37
}
func F_dpow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v18 float64
	_ = v18
	var v23 float64
	_ = v23
	var v26 float64
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 float64
	_ = v35
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 float64
	_ = v42
	var v57 float64
	_ = v57
	var v60 float64
	_ = v60
	var v66 float64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v85 float64
	_ = v85
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 float64
	_ = v95
	var v101 float64
	_ = v101
	var v104 float64
	_ = v104
	var v124 float64
	_ = v124
	var v126 int32
	_ = v126
	var v127 float64
	_ = v127
	var v130 float64
	_ = v130
	var v133 float64
	_ = v133
	var v136 float64
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v11 = base.I64_reinterpret_f64(v8) & int64(9223372036854775807)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	v14 = base.F64_abs(v13)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v14)) {
		v18 = math.Float64frombits(uint64(0x7ff8000000000000))
		if base.F64_ne(v8, float64(0)) != 0 {
			v23 = v18
		} else {
			v23 = float64(1)
		}
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(v11) {
			v26 = v18
		} else {
			v26 = v23
		}
		v27 = F_Float8GetDatum(m, v26)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			return v27
		}
	} else {
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v11) {
			v35 = float64(1)
			if base.F64_ne(v13, v35) != 0 {
				v38 = math.Float64frombits(uint64(0x7ff8000000000000))
			} else {
				v38 = v35
			}
			v39 = F_Float8GetDatum(m, v38)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				return v39
			}
		} else {
			v42 = float64(0)
			if base.F64_eq(v13, v42)&base.F64_lt(v8, v42) == int32(0) {
				if base.F64_ne(base.F64_floor(v8), v8)&base.F64_lt(v13, float64(0)) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v159 = m.ExcPending
					if v159 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(369361026))
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_dpow_0), int32(0))
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_dpow_1), int32(1526), int32(_a_F_dpow_2))
								mBase = m.M
								v171 = m.ExcPending
								if v171 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if base.F64_eq(base.F64_abs(v8), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v57 = float64(1)
						if base.F64_eq(v14, v57) != 0 {
							v136 = v57
							v137 = F_Float8GetDatum(m, v136)
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return int32(0)
							} else {
								return v137
							}
						} else {
							v60 = float64(0)
							if base.F64_gt(v8, v60) != 0 {
								if base.F64_gt(v14, float64(1)) != 0 {
									v66 = v8
								} else {
									v66 = float64(0)
								}
								v67 = F_Float8GetDatum(m, v66)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									return v67
								}
							} else {
								if base.F64_gt(v14, float64(1)) != 0 {
									v136 = v60
									v137 = F_Float8GetDatum(m, v136)
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int32(0)
									} else {
										return v137
									}
								} else {
									v73 = F_Float8GetDatum(m, base.F64_neg(v8))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										return v73
									}
								}
							}
						}
					} else {
						if base.F64_eq(v14, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							if base.F64_eq(v8, float64(0)) != 0 {
								v136 = float64(1)
								v137 = F_Float8GetDatum(m, v136)
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return int32(0)
								} else {
									return v137
								}
							} else {
								if base.F64_gt(v13, float64(0)) == int32(0) {
									v124 = base.F64_mul(v8, float64(0.5))
									v126 = base.F64_ne(base.F64_floor(v124), v124)
									if v126 != 0 {
										v127 = v13
									} else {
										v127 = base.F64_neg(v13)
									}
									if v126 != 0 {
										v130 = math.Float64frombits(uint64(0x8000000000000000))
									} else {
										v130 = float64(0)
									}
									if base.F64_gt(v8, float64(0)) != 0 {
										v133 = v127
									} else {
										v133 = v130
									}
									v136 = v133
									v137 = F_Float8GetDatum(m, v136)
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int32(0)
									} else {
										return v137
									}
								} else {
									v85 = float64(0)
									if base.F64_gt(v8, v85) != 0 {
										v88 = v13
									} else {
										v88 = v85
									}
									v89 = F_Float8GetDatum(m, v88)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int32(0)
									} else {
										return v89
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_dpow[0])) = int32(0)
							v95 = F_pow(m, v13, v8)
							mBase = m.M
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v95)&int64(9223372036854775807)) {
								v101 = float64(0)
								if base.F64_eq(v13, v101) != 0 {
									v136 = v101
									v137 = F_Float8GetDatum(m, v136)
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int32(0)
									} else {
										return v137
									}
								} else {
									v104 = float64(1)
									if base.F64_eq(v14, v104) != 0 {
										v136 = v104
										v137 = F_Float8GetDatum(m, v136)
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return int32(0)
										} else {
											return v137
										}
									} else {
										if base.F64_ge(v8, float64(0)) != 0 {
											if base.F64_gt(v14, float64(1)) == int32(0) {
												F_float_underflow_error(m)
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												F_float_overflow_error(m)
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										} else {
											if base.F64_lt(v14, float64(1)) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												F_float_underflow_error(m)
												mBase = m.M
												v174 = m.ExcPending
												if v174 != 0 {
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
							} else {
								if base.F64_eq(base.F64_abs(v95), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if base.F64_ne(v95, float64(0)) != 0 {
										v136 = v95
										v137 = F_Float8GetDatum(m, v136)
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return int32(0)
										} else {
											return v137
										}
									} else {
										if base.F64_ne(v13, float64(0)) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v174 = m.ExcPending
											if v174 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v136 = v95
											v137 = F_Float8GetDatum(m, v136)
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												return v137
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v143 = m.ExcPending
				if v143 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(369361026))
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_dpow_3), int32(0))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_dpow_1), int32(1522), int32(_a_F_dpow_2))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
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
func F_drop_parent_dependency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
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
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v9 = m.G0
	v11 = v9 - int32(144)
	m.G0 = v11
	v15 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v11, int32(1), int32(3), int32(184), int32(1259))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v11+int32(48), int32(2), int32(3), int32(184), l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = int32(3)
	F_ScanKeyInit(m, v11+int32(96), v32, v32, int32(65), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v42 = F_systable_beginscan(m, v15, int32(2673), int32(1), int32(0), int32(3), v11)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v44 = F_systable_getnext(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v44 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = v44
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_systable_endscan(m, v42)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L21
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+22)))
	v56 = v54 + v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v57 != l1 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v68 = F_systable_getnext(m, v42)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L19
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	if v59 != l2 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	if v61 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v62 = int32(*(*int8)(unsafe.Add(mBase, uint32(v56)+24)))
	if l3 != v62 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	F_simple_heap_delete(m, v15, v46+int32(4))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	if v68 != 0 {
		v46 = v68
		goto L11
	} else {
		goto L20
	}
L20:
	;
	goto L12
L21:
	;
	F_relation_close(m, v15, int32(3))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	m.G0 = v11 + int32(144)
	return
}
func F_dropconstraint_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	v18 = m.G0
	v20 = v18 - int32(272)
	m.G0 = v20
	F_check_stack_depth(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l5 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ATSimplePermissions(m, int32(22), l1, int32(289))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v30 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
	v34 = v32 + v33
	v36 = v34 + int32(4)
	if l5 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L154
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L151
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L147
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L142
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L139
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L134
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L130
	}
L15:
	;
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+104)))
	if int32(0) < v39 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+72)))
	if v42 == int32(110) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v47 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v210 = v42
	v220 = int32(0)
	goto L21
L21:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+106)))
	if v210&int32(255) != int32(102) {
		goto L61
	} else {
		goto L62
	}
L22:
	;
	v49 = F_extractNotNullColumn(m, l2)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v53 = F_get_attname(m, v51, v49, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v56 = F_RelationGetIndexAttrBitmap(m, l1, int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	v179 = F_RelationGetIndexAttrBitmap(m, l1, int32(2))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L50
	}
L26:
	;
	F_relation_close(m, v71, int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L49
	}
L27:
	;
	if v56 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+119)))
	if v61 != int32(112) {
		goto L25
	} else {
		goto L31
	}
L29:
	;
	v113 = v56
	goto L30
L30:
	;
	v130 = F_bms_is_member(m, v49+int32(7), v113)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L42
	}
L31:
	;
	v65 = F_RelationGetPrimaryKeyIndex(m, l1, int32(1))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v65 == int32(0) {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v71 = F_relation_open(m, v65, int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+192))
	v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+10)))
	if v74 <= int32(0) {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v80 = int32(0)
	v83 = int32(0)
	v91 = v73
	goto L36
L36:
	;
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91+v83<<(uint(int32(1))%32))+48)))
	v99 = F_bms_add_member(m, v80, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	F_relation_close(m, v71, int32(1))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v102 = v83 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v71)+192))
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+10)))
	if v102 < v104 {
		v80 = v99
		v83 = v102
		v91 = v103
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	if v99 == int32(0) {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v113 = v99
	goto L30
L42:
	;
	if v130 == int32(0) {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v143 = F_get_attname(m, v141, v49, int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v143
	F_errmsg(m, int32(_a_F_dropconstraint_internal_0), v20+int32(96))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_2), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	goto L25
L50:
	;
	v181 = F_bms_is_member(m, v49+int32(7), v179)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v181 != 0 {
		goto L13
	} else {
		goto L52
	}
L52:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v184 = F_SearchSysCacheCopyAttNum(m, v183, v49)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v184 == int32(0) {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+22)))
	v190 = v188 + v189
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+89)))
	if v191 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+86)))
	if v192 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+86)) = uint8(v195)
	F_CatalogTupleUpdate(m, v47, v184+int32(4), v184)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	F_relation_close(m, v47, int32(3))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+72)))
	v210 = v204
	v220 = v53
	goto L21
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v250 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v249
	F_performDeletion(m, l0, l3, v250)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L71
	}
L62:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v34)+96))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v227 == v228 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v231 = F_table_open(m, v227, int32(8))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v231)+48))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+118)))
	if v234 == int32(116) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+24)))
	if v237 == int32(0) {
		goto L8
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_CheckTableNotInUse(m, v231, int32(_a_F_dropconstraint_internal_4))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	F_relation_close(m, v231, int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L61
L71:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+72)))
	switch v256 - int32(99) {
	case 0, 11:
		goto L73
	default:
		goto L74
	}
L72:
	;
	F_relation_close(m, v30, int32(3))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L129
	}
L73:
	;
	if v222&int32(1) != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+119)))
	if v260 != int32(112) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	goto L72
L77:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v266 = F_find_inheritance_children(m, v265, l6)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v266 == int32(0) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v270 <= int32(0) {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v280 = int32(0)
	goto L81
L81:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v295+v280<<(uint(int32(2))%32))))
	v301 = F_table_open(m, v299, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L83
	}
L82:
	;
	goto L76
L83:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v301)+48))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+118)))
	if v304 == int32(116) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+24)))
	if v307 == int32(0) {
		goto L8
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	F_CheckTableNotInUse(m, v301, int32(_a_F_dropconstraint_internal_4))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+72)))
	if v313 == int32(110) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v367)+16))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+22)))
	v371 = v369 + v370
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+72)))
	switch v372 - int32(99) {
	case 0, 11:
		goto L106
	default:
		goto L107
	}
L90:
	;
	v316 = F_findNotNullConstraint(m, v299, v220)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v336 = v20 + int32(128)
	F_ScanKeyInit(m, v336, int32(9), int32(3), int32(184), v299)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L98
	}
L93:
	;
	if v316 != 0 {
		v367 = v316
		goto L89
	} else {
		goto L94
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v301)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v220
	F_errmsg_internal(m, int32(_a_F_dropconstraint_internal_5), v20+int32(32))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_6), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
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
	F_ScanKeyInit(m, v20+int32(176), int32(10), int32(3), int32(184), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_ScanKeyInit(m, v20+int32(224), int32(2), int32(3), int32(62), v36)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v357 = F_systable_beginscan(m, v30, int32(2665), int32(1), int32(0), int32(3), v336)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v359 = F_systable_getnext(m, v357)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	if v359 == int32(0) {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v363 = F_heap_copytuple(m, v359)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_systable_endscan(m, v357)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v367 = v363
	goto L89
L106:
	;
	v388 = int32(*(*int16)(unsafe.Add(mBase, uint32(v371)+104)))
	if v388 <= int32(0) {
		goto L9
	} else {
		goto L111
	}
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errmsg_internal(m, int32(_a_F_dropconstraint_internal_7), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_8), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	if l4 != 0 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	F_pfree(m, v367)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L126
	}
L113:
	;
	F_CatalogTupleUpdate(m, v30, v367+int32(4), v367)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L124
	}
L114:
	;
	if v388 != int32(1) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	v404 = v388 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+104)) = uint16(v404)
	if v404&int32(_a_F_dropconstraint_internal_9) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v401 = v388 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+104)) = uint16(v401)
	goto L113
L118:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+103)))
	if v393 != 0 {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v396 = int32(1)
	F_dropconstraint_internal(m, v20+int32(128), v301, v367, l3, v396, v396, l6)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L112
L121:
	;
	v410 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v371)+103)) = uint8(v410)
	goto L123
L122:
	;
	goto L123
L123:
	;
	goto L113
L124:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L112
L126:
	;
	F_relation_close(m, v301, int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v426 = v280 + int32(1)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v426 < v427 {
		v280 = v426
		goto L81
	} else {
		goto L128
	}
L128:
	;
	goto L82
L129:
	;
	m.G0 = v20 + int32(272)
	return
L130:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v476 + int32(4)
	F_errmsg(m, int32(_a_F_dropconstraint_internal_10), v20+int32(112))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_11), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v500 = F_get_attname(m, v498, v49, int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v500
	F_errmsg(m, int32(_a_F_dropconstraint_internal_12), v20)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_13), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v49
	F_errmsg_internal(m, int32(_a_F_dropconstraint_internal_14), v20+int32(16))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_15), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v537 = F_get_attname(m, v535, v49, int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v539 + int32(4)
	F_errmsg(m, int32(_a_F_dropconstraint_internal_16), v20+int32(80))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_17), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v301)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v561 + int32(4)
	F_errmsg(m, int32(_a_F_dropconstraint_internal_18), v20-int32(-64))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_19), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v371 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v299
	F_errmsg_internal(m, int32(_a_F_dropconstraint_internal_20), v20+int32(48))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_21), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errmsg(m, int32(_a_F_dropconstraint_internal_22), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_23), int32(_a_F_dropconstraint_internal_24))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dsimple_lexize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = F_str_tolower(m, v4, v5, int32(100))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		if v11 != 0 {
			v12 = F_searchstoplist(m, v3, v7)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				if v12 == int32(0) {
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+8)))
					if v22 == int32(1) {
						v26 = F_palloc0(m, int32(16))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v7
							return v26
						}
					} else {
						F_pfree(m, v7)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				} else {
					F_pfree(m, v7)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int32(0)
					} else {
						v19 = F_palloc0(m, int32(16))
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return int32(0)
						} else {
							return v19
						}
					}
				}
			}
		} else {
			F_pfree(m, v7)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v19 = F_palloc0(m, int32(16))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v19
				}
			}
		}
	}
}
func F_dsind(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 float64
	_ = v42
	var v45 int64
	_ = v45
	var v52 float64
	_ = v52
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v141 int64
	_ = v141
	var v144 int32
	_ = v144
	var v159 int64
	_ = v159
	var v167 float64
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 float64
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 float64
	_ = v181
	var v185 float64
	_ = v185
	var v189 float64
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v209 float64
	_ = v209
	var v213 int32
	_ = v213
	var v214 float64
	_ = v214
	var v215 float64
	_ = v215
	var v221 float64
	_ = v221
	var v222 float64
	_ = v222
	var v224 float64
	_ = v224
	var v226 float64
	_ = v226
	var v228 float64
	_ = v228
	var v237 float64
	_ = v237
	var v245 float64
	_ = v245
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v265 float64
	_ = v265
	var v269 int32
	_ = v269
	var v270 float64
	_ = v270
	var v271 float64
	_ = v271
	var v276 float64
	_ = v276
	var v278 float64
	_ = v278
	var v280 float64
	_ = v280
	var v283 float64
	_ = v283
	var v287 float64
	_ = v287
	var v291 float64
	_ = v291
	var v295 float64
	_ = v295
	var v301 float64
	_ = v301
	var v302 float64
	_ = v302
	var v307 float64
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L95
	} else {
		goto L101
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L95
	} else {
		goto L97
	}
L3:
	;
	if base.F64_eq(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	v307 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L5
L5:
	;
	v310 = F_Float8GetDatum(m, v307)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L95
	} else {
		goto L96
	}
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsind[0])))
	if v22 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_init_degree_constants(m)
	mBase = m.M
	goto L9
L8:
	;
	goto L9
L9:
	;
	v33 = base.I64_reinterpret_f64(v12)
	v37 = int32(2047)
	v38 = base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(52))%64))) & v37
	if v38 == v37 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v169 = base.F64_lt(v167, float64(0))
	if v169 != 0 {
		goto L51
	} else {
		goto L52
	}
L11:
	;
	v42 = base.F64_mul(v12, float64(360))
	v167 = base.F64_div(v42, v42)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v45 = v33 << (uint(int64(1)) % 64)
	if base.Ui64(v45) <= base.Ui64(int64(-9156662467374350336)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v45 == int64(-9156662467374350336) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if v38 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v52 = base.F64_mul(v12, float64(0))
	goto L19
L18:
	;
	v52 = v12
	goto L19
L19:
	;
	v167 = v52
	goto L10
L20:
	;
	if int32(1031) < v88 {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	v55 = int32(0)
	v57 = v33 << (uint(int64(12)) % 64)
	if int64(0) <= v57 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v88 = v38
	v89 = v33&int64(4503599627370495) | int64(4503599627370496)
	goto L20
L24:
	;
	v61 = v57
	v64 = v55
	goto L27
L25:
	;
	v75 = v55
	goto L26
L26:
	;
	v88 = v75
	v89 = v33 << (uint(base.I64_extend_i32_u(int32(1)-v75)) % 64)
	goto L20
L27:
	;
	v66 = v64 - int32(1)
	v68 = v61 << (uint(int64(1)) % 64)
	if int64(0) <= v68 {
		v61 = v68
		v64 = v66
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v75 = v66
	goto L26
L29:
	;
	goto L28
L30:
	;
	v93 = v89
	v96 = v88
	goto L33
L31:
	;
	v114 = v89
	v117 = v88
	goto L32
L32:
	;
	v119 = v114 - int64(6333186975989760)
	if v119 < int64(0) {
		v126 = v114
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v98 = v93 - int64(6333186975989760)
	if v98 < int64(0) {
		v105 = v93
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v114 = v107
	v117 = int32(1031)
	goto L32
L35:
	;
	v107 = v105 << (uint(int64(1)) % 64)
	v109 = v96 - int32(1)
	if int32(1031) < v109 {
		v93 = v107
		v96 = v109
		goto L33
	} else {
		goto L38
	}
L36:
	;
	if v98 != int64(0) {
		v105 = v98
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v167 = base.F64_mul(v12, float64(0))
	goto L10
L38:
	;
	goto L34
L39:
	;
	if base.Ui64(v126) <= base.Ui64(int64(4503599627370495)) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if v119 != int64(0) {
		v126 = v119
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v167 = base.F64_mul(v12, float64(0))
	goto L10
L42:
	;
	v130 = v126
	v133 = v117
	goto L45
L43:
	;
	v141 = v126
	v144 = v117
	goto L44
L44:
	;
	if int32(0) < v144 {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v135 = v133 - int32(1)
	v137 = v130 << (uint(int64(1)) % 64)
	if base.Ui64(v130) < base.Ui64(int64(2251799813685248)) {
		v130 = v137
		v133 = v135
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v141 = v137
	v144 = v135
	goto L44
L47:
	;
	goto L46
L48:
	;
	v159 = v141 - int64(4503599627370496) | base.I64_extend_i32_u(v144)<<(uint(int64(52))%64)
	goto L50
L49:
	;
	v159 = int64(base.Ui64(v141) >> (uint(base.I64_extend_i32_u(int32(1)-v144)) % 64))
	goto L50
L50:
	;
	v167 = base.F64_reinterpret_i64(v33&int64(-9223372036854775807-1) | v159)
	goto L10
L51:
	;
	v170 = int32(-1)
	goto L53
L52:
	;
	v170 = int32(1)
	goto L53
L53:
	;
	if v169 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v173 = base.F64_neg(v167)
	goto L56
L55:
	;
	v173 = v167
	goto L56
L56:
	;
	v175 = base.F64_gt(v173, float64(180))
	if v175 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v176 = int32(0) - v170
	goto L59
L58:
	;
	v176 = v170
	goto L59
L59:
	;
	if v175 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v302 = base.F64_mul(v301, base.F64_convert_i32_s(v176))
	if base.F64_eq(base.F64_abs(v302), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L94
	}
L61:
	;
	v181 = base.F64_sub(float64(360), v173)
	goto L63
L62:
	;
	v181 = v173
	goto L63
L63:
	;
	if base.F64_gt(v181, float64(90)) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v185 = base.F64_sub(float64(180), v181)
	goto L66
L65:
	;
	v185 = v181
	goto L66
L66:
	;
	if base.F64_le(v185, float64(30)) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v189 = base.F64_mul(v185, float64(0.017453292519943295))
	v193 = m.G0
	v195 = v193 - int32(16)
	m.G0 = v195
	v202 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v189))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v202) <= base.Ui32(int32(1072243195)) {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	goto L69
L69:
	;
	v245 = base.F64_mul(base.F64_sub(float64(90), v185), float64(0.017453292519943295))
	v249 = m.G0
	v251 = v249 - int32(16)
	m.G0 = v251
	v258 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v245))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v258) <= base.Ui32(int32(1072243195)) {
		goto L85
	} else {
		goto L86
	}
L70:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v228
	v237 = *(*float64)(unsafe.Add(mBase, _c_F_dsind[1]))
	v301 = base.F64_mul(base.F64_div(v228, v237), float64(0.5))
	goto L60
L71:
	;
	m.G0 = v195 + int32(16)
	goto L70
L72:
	;
	if base.Ui32(v202) < base.Ui32(int32(1045430272)) {
		v228 = v189
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v202) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v209 = F___sin(m, v189, float64(0), int32(0))
	mBase = m.M
	v228 = v209
	goto L71
L76:
	;
	v228 = base.F64_sub(v189, v189)
	goto L71
L77:
	;
	goto L78
L78:
	;
	v213 = F___rem_pio2(m, v189, v195)
	mBase = m.M
	v214 = *(*float64)(unsafe.Add(mBase, uint32(v195)+8))
	v215 = *(*float64)(unsafe.Add(mBase, uint32(v195)))
	switch v213&int32(3) - int32(1) {
	case 0:
		goto L81
	case 1:
		goto L80
	case 2:
		goto L79
	default:
		goto L82
	}
L79:
	;
	v226 = F___cos(m, v215, v214)
	mBase = m.M
	v228 = base.F64_neg(v226)
	goto L71
L80:
	;
	v224 = F___sin(m, v215, v214, int32(1))
	mBase = m.M
	v228 = base.F64_neg(v224)
	goto L71
L81:
	;
	v222 = F___cos(m, v215, v214)
	mBase = m.M
	v228 = v222
	goto L71
L82:
	;
	v221 = F___sin(m, v215, v214, int32(1))
	mBase = m.M
	v228 = v221
	goto L71
L83:
	;
	v291 = base.F64_sub(float64(1), v287)
	*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v291
	v295 = *(*float64)(unsafe.Add(mBase, _c_F_dsind[2]))
	v301 = base.F64_add(base.F64_mul(base.F64_div(v291, v295), float64(-0.5)), float64(1))
	goto L60
L84:
	;
	m.G0 = v251 + int32(16)
	goto L83
L85:
	;
	if base.Ui32(v258) < base.Ui32(int32(1044816030)) {
		v287 = float64(1)
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v258) {
		v287 = base.F64_sub(v245, v245)
		goto L84
	} else {
		goto L89
	}
L88:
	;
	v265 = F___cos(m, v245, float64(0))
	mBase = m.M
	v287 = v265
	goto L84
L89:
	;
	v269 = F___rem_pio2(m, v245, v251)
	mBase = m.M
	v270 = *(*float64)(unsafe.Add(mBase, uint32(v251)+8))
	v271 = *(*float64)(unsafe.Add(mBase, uint32(v251)))
	switch v269&int32(3) - int32(1) {
	case 0:
		goto L92
	case 1:
		goto L91
	case 2:
		goto L90
	default:
		goto L93
	}
L90:
	;
	v283 = F___sin(m, v271, v270, int32(1))
	mBase = m.M
	v287 = v283
	goto L84
L91:
	;
	v280 = F___cos(m, v271, v270)
	mBase = m.M
	v287 = base.F64_neg(v280)
	goto L84
L92:
	;
	v278 = F___sin(m, v271, v270, int32(1))
	mBase = m.M
	v287 = base.F64_neg(v278)
	goto L84
L93:
	;
	v276 = F___cos(m, v271, v270)
	mBase = m.M
	v287 = v276
	goto L84
L94:
	;
	v307 = v302
	goto L5
L95:
	;
	return int32(0)
L96:
	;
	m.G0 = v8 + int32(16)
	return v310
L97:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_dsind_0), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L95
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_dsind_1), int32(2455), int32(_a_F_dsind_2))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L95
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dsqrt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v10 float64
	_ = v10
	var v12 float64
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	if base.F64_lt(v5, float64(0)) == int32(0) {
		v10 = base.F64_sqrt(v5)
		v12 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_eq(base.F64_abs(v10), v12)&base.F64_ne(base.F64_abs(v5), v12) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v18 = F_Float8GetDatum(m, v10)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return v18
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(369361026))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_dsqrt_0), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_dsqrt_1), int32(1454), int32(_a_F_dsqrt_2))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
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
func F_dsynonym_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v2
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+30)) = uint16(v2)
	if v13 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L21
	} else {
		goto L81
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L21
	} else {
		goto L77
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(0) < v20 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L21
	} else {
		goto L73
	}
L5:
	;
	v24 = int32(0)
	v29 = v2
	v31 = v2
	goto L8
L6:
	;
	v111 = v2
	v113 = v2
	goto L7
L7:
	;
	if v111 == int32(0) {
		goto L2
	} else {
		goto L33
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v24<<(uint(int32(2))%32))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v38 = int32(_a_F_dsynonym_init_0)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsynonym_init[0])))
	if base.B2i32(v41 == int32(0))|base.B2i32(v41 != v44) != 0 {
		v62 = v41
		v63 = v44
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v111 = v100
	v113 = v101
	goto L7
L10:
	;
	v103 = v24 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v103 < v104 {
		v24 = v103
		v29 = v100
		v31 = v101
		goto L8
	} else {
		goto L32
	}
L11:
	;
	if v62-v63 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	v47 = v37
	v48 = v38
	goto L14
L14:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v52 == int32(0) {
		v62 = v52
		v63 = v51
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v62 = v52
	v63 = v51
	goto L12
L16:
	;
	v55 = int32(1)
	if v52 == v51 {
		v47 = v47 + v55
		v48 = v48 + v55
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v67 = F_defGetString(m, v36)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v71 = int32(_a_F_dsynonym_init_1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsynonym_init[1])))
	if base.B2i32(v74 == int32(0))|base.B2i32(v74 != v77) != 0 {
		v95 = v74
		v96 = v77
		goto L24
	} else {
		goto L25
	}
L21:
	;
	return int32(0)
L22:
	;
	v100 = v67
	v101 = v31
	goto L10
L23:
	;
	if v95-v96 != 0 {
		goto L4
	} else {
		goto L30
	}
L24:
	;
	goto L23
L25:
	;
	v80 = v37
	v81 = v71
	goto L26
L26:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	if v85 == int32(0) {
		v95 = v85
		v96 = v84
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v95 = v85
	v96 = v84
	goto L24
L28:
	;
	v88 = int32(1)
	if v85 == v84 {
		v80 = v80 + v88
		v81 = v81 + v88
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v98 = F_defGetBoolean(m, v36)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v100 = v29
	v101 = v98
	goto L10
L32:
	;
	goto L9
L33:
	;
	v117 = v11 + int32(36)
	v119 = F_get_tsearch_config_filename(m, v111, int32(_a_F_dsynonym_init_2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L21
	} else {
		goto L34
	}
L34:
	;
	v121 = F_tsearch_readline_begin(m, v117, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	if v121 == int32(0) {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v126 = F_palloc0(m, int32(12))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	v128 = F_tsearch_readline(m, v117)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L21
	} else {
		goto L39
	}
L38:
	;
	F_tsearch_readline_end(m, v11+int32(36))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L21
	} else {
		goto L71
	}
L39:
	;
	if v128 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v239 = int32(0)
	goto L38
L41:
	;
	goto L42
L42:
	;
	v134 = v128
	v137 = int32(0)
	goto L43
L43:
	;
	v143 = v11 + int32(32)
	v145 = F_findwrd(m, v134, v143, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L21
	} else {
		goto L46
	}
L44:
	;
	v239 = v226
	goto L38
L45:
	;
	F_pfree(m, v134)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L21
	} else {
		goto L68
	}
L46:
	;
	if v145 == int32(0) {
		v226 = v137
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v150 == int32(0) {
		v226 = v137
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v153 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v153)
	v159 = F_findwrd(m, v149+int32(1), v143, v11+int32(30))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	if v159 == int32(0) {
		v226 = v137
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v164 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v163))) = uint8(v164)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v166 <= v137 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v166 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L53
L53:
	;
	if v113&int32(1) != 0 {
		goto L61
	} else {
		goto L62
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v183
	goto L53
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = int32(64)
	v173 = F_palloc(m, int32(1024))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L21
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v166 << (uint(int32(1)) % 32)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v181 = F_repalloc(m, v178, v166<<(uint(int32(5))%32))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L21
	} else {
		goto L59
	}
L58:
	;
	v183 = v173
	goto L54
L59:
	;
	v183 = v181
	goto L54
L60:
	;
	v212 = v137 << (uint(int32(4)) % 32)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v212+v213)+4)) = v210
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v218 = F_strlen(m, v159)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v216+v212)+8)) = v218
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+30)))
	*(*uint16)(unsafe.Add(mBase, uint32(v220+v212)+12)) = uint16(v222)
	v226 = v137 + int32(1)
	goto L45
L61:
	;
	v187 = F_pstrdup(m, v145)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L21
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v196 = F_strlen(m, v145)
	mBase = m.M
	v198 = F_str_tolower(m, v145, v196, int32(100))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L21
	} else {
		goto L66
	}
L64:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v189+v137<<(uint(int32(4))%32)))) = v187
	v194 = F_pstrdup(m, v159)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L21
	} else {
		goto L65
	}
L65:
	;
	v210 = v194
	goto L60
L66:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v200+v137<<(uint(int32(4))%32)))) = v198
	v205 = F_strlen(m, v159)
	mBase = m.M
	v207 = F_str_tolower(m, v159, v205, int32(100))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L21
	} else {
		goto L67
	}
L67:
	;
	v210 = v207
	goto L60
L68:
	;
	v234 = F_tsearch_readline(m, v11+int32(36))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L21
	} else {
		goto L69
	}
L69:
	;
	if v234 != 0 {
		v134 = v234
		v137 = v226
		goto L43
	} else {
		goto L70
	}
L70:
	;
	goto L44
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v239
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	F_pg_qsort(m, v249, v239, int32(16), int32(1148))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L21
	} else {
		goto L72
	}
L72:
	;
	v255 = v113 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+8)) = uint8(v255)
	m.G0 = v11 + int32(80)
	return v126
L73:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L21
	} else {
		goto L74
	}
L74:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v268
	F_errmsg(m, int32(_a_F_dsynonym_init_3), v11+int32(16))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L21
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_dsynonym_init_4), int32(121), int32(_a_F_dsynonym_init_5))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L21
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L21
	} else {
		goto L78
	}
L78:
	;
	F_errmsg(m, int32(_a_F_dsynonym_init_6), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L21
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_dsynonym_init_4), int32(127), int32(_a_F_dsynonym_init_5))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L21
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
	F_errcode(m, int32(22))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L21
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v119
	F_errmsg(m, int32(_a_F_dsynonym_init_7), v11)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L21
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_dsynonym_init_4), int32(135), int32(_a_F_dsynonym_init_5))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L21
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dtrunc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v11 float64
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	if base.F64_ge(v4, float64(0)) != 0 {
		v11 = base.F64_floor(v4)
	} else {
		v11 = base.F64_neg(base.F64_floor(base.F64_neg(v4)))
	}
	v12 = F_Float8GetDatum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
