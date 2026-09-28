package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_DeconstructFkConstraintRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
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
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v23 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(21))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L107
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L8
	} else {
		goto L104
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L101
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L98
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L8
	} else {
		goto L95
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L92
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L8
	} else {
		goto L89
	}
L8:
	;
	return
L9:
	;
	v25 = base.I32_wrap_i64(v23)
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v28 != int32(1) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v31 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v32 != int32(21) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if base.Ui32(v35-int32(33)) <= base.Ui32(int32(-33)) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v41 = v35 << (uint(int32(1)) % 32)
	v42 = int32(0)
	v43 = base.B2i32(v41 == v42)
	if v43 == v42 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	base.MemoryCopy(m, l2, v26+int32(24), v41)
	goto L17
L16:
	;
	goto L17
L17:
	;
	if v26 != v25 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_pfree(m, v26)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v54 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(22))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v56 = base.I32_wrap_i64(v54)
	v57 = F_pg_detoast_datum(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v59 != int32(1) {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v62 != v35 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v64 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	if v65 != int32(21) {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	if v43 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	base.MemoryCopy(m, l3, v57+int32(24), v41)
	goto L30
L29:
	;
	goto L30
L30:
	;
	if v57 != v56 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_pfree(m, v57)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
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
	v80 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(23))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v82 = base.I32_wrap_i64(v80)
	v83 = F_pg_detoast_datum(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v85 != int32(1) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	if v88 != v35 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v90 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	if v91 != int32(26) {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v95 = v35 << (uint(int32(2)) % 32)
	if v95 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	base.MemoryCopy(m, l4, v83+int32(24), v95)
	goto L45
L44:
	;
	goto L45
L45:
	;
	if v83 == v82 {
		goto L35
	} else {
		goto L46
	}
L46:
	;
	F_pfree(m, v83)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
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
	v109 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(24))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v111 = base.I32_wrap_i64(v109)
	v112 = F_pg_detoast_datum(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v114 != int32(1) {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if v117 != v35 {
		goto L3
	} else {
		goto L53
	}
L53:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	if v119 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	if v120 != int32(26) {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v124 = v35 << (uint(int32(2)) % 32)
	if v124 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	base.MemoryCopy(m, l5, v112+int32(24), v124)
	goto L58
L57:
	;
	goto L58
L58:
	;
	if v112 == v111 {
		goto L48
	} else {
		goto L59
	}
L59:
	;
	F_pfree(m, v112)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
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
	v138 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(25))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	v140 = base.I32_wrap_i64(v138)
	v141 = F_pg_detoast_datum(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v143 != int32(1) {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	if v146 != v35 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	if v148 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	if v149 != int32(26) {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v153 = v35 << (uint(int32(2)) % 32)
	if v153 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	base.MemoryCopy(m, l6, v141+int32(24), v153)
	goto L71
L70:
	;
	goto L71
L71:
	;
	if v141 == v140 {
		goto L61
	} else {
		goto L72
	}
L72:
	;
	F_pfree(m, v141)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	goto L61
L74:
	;
	v168 = F_SysCacheGetAttr(m, int32(19), l0, int32(26), v19+int32(15))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L8
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
	m.G0 = v19 + int32(16)
	return
L77:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+15)))
	if v170 != 0 {
		v193 = int32(0)
		goto L78
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v193
	goto L76
L79:
	;
	v171 = base.I32_wrap_i64(v168)
	v172 = F_pg_detoast_datum(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v174 != int32(1) {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172)+8))
	if v177 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	if v178 != int32(21) {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v172)+16))
	v183 = v181 << (uint(int32(1)) % 32)
	if v183 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	base.MemoryCopy(m, l8, v172+int32(24), v183)
	goto L86
L85:
	;
	goto L86
L86:
	;
	if v172 == v171 {
		v193 = v181
		goto L78
	} else {
		goto L87
	}
L87:
	;
	F_pfree(m, v172)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	v193 = v181
	goto L78
L89:
	;
	F_errmsg_internal(m, int32(_a_F_DeconstructFkConstraintRow_0), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1562), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v35
	F_errmsg_internal(m, int32(_a_F_DeconstructFkConstraintRow_3), v19)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1565), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
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
	v237 = m.ExcPending
	if v237 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1577), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
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
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1592), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
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
	v263 = m.ExcPending
	if v263 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1607), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
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
	v276 = m.ExcPending
	if v276 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1622), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
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
	v289 = m.ExcPending
	if v289 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_DeconstructFkConstraintRow_1), int32(1644), int32(_a_F_DeconstructFkConstraintRow_2))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
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
	F_errfinish(m, int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_10), int32(1833), int32(_a_F_DetermineTimeZoneAbbrevOffsetTS_11))
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
func F_datan(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v9 float64
	_ = v9
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v31 float64
	_ = v31
	var v37 float64
	_ = v37
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v85 float64
	_ = v85
	var v102 float64
	_ = v102
	var v109 int32
	_ = v109
	var v110 float64
	_ = v110
	var v113 float64
	_ = v113
	var v116 float64
	_ = v116
	var v120 float64
	_ = v120
	var v121 float64
	_ = v121
	var v131 float64
	_ = v131
	var v138 int64
	_ = v138
	var v143 int32
	_ = v143
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v4&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v9 = base.F64_reinterpret_i64(v4)
		v16 = base.I64_reinterpret_f64(v9)
		v21 = base.I32_wrap_i64(int64(base.Ui64(v16)>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(int32(1141899264)) <= base.Ui32(v21) {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v9)&int64(9223372036854775807)) {
				v31 = v9
			} else {
				v31 = base.F64_copysign(float64(1.5707963267948966), v9)
			}
			v131 = v31
		} else {
			if base.Ui32(v21) <= base.Ui32(int32(1071382527)) {
				if base.Ui32(int32(1044381696)) <= base.Ui32(v21) {
					v68 = v9
					v69 = int32(-1)
					v70 = base.F64_mul(v68, v68)
					v71 = base.F64_mul(v70, v70)
					v85 = base.F64_mul(v71, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
					v102 = base.F64_mul(v70, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
					if base.Ui32(v21) <= base.Ui32(int32(1071382527)) {
						v131 = base.F64_sub(v68, base.F64_mul(v68, base.F64_add(v85, v102)))
					} else {
						v109 = v69 << (uint(int32(3)) % 32)
						v110 = *(*float64)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_datan[0])))
						v113 = *(*float64)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_datan[1])))
						v116 = base.F64_sub(v110, base.F64_sub(base.F64_sub(base.F64_mul(v68, base.F64_add(v85, v102)), v113), v68))
						if v16 < int64(0) {
							v120 = base.F64_neg(v116)
						} else {
							v120 = v116
						}
						v121 = v120
						v131 = v121
					}
				} else {
					v121 = v9
					v131 = v121
				}
			} else {
				v37 = base.F64_abs(v9)
				if base.Ui32(v21) <= base.Ui32(int32(1072889855)) {
					if base.Ui32(v21) <= base.Ui32(int32(1072037887)) {
						v68 = base.F64_div(base.F64_add(base.F64_add(v37, v37), float64(-1)), base.F64_add(v37, float64(2)))
						v69 = int32(0)
					} else {
						v68 = base.F64_div(base.F64_add(v37, float64(-1)), base.F64_add(v37, float64(1)))
						v69 = int32(1)
					}
				} else {
					if base.Ui32(v21) <= base.Ui32(int32(1073971199)) {
						v68 = base.F64_div(base.F64_add(v37, float64(-1.5)), base.F64_add(base.F64_mul(v37, float64(1.5)), float64(1)))
						v69 = int32(2)
					} else {
						v68 = base.F64_div(float64(-1), v37)
						v69 = int32(3)
					}
				}
				v70 = base.F64_mul(v68, v68)
				v71 = base.F64_mul(v70, v70)
				v85 = base.F64_mul(v71, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
				v102 = base.F64_mul(v70, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, base.F64_add(base.F64_mul(v71, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
				if base.Ui32(v21) <= base.Ui32(int32(1071382527)) {
					v131 = base.F64_sub(v68, base.F64_mul(v68, base.F64_add(v85, v102)))
				} else {
					v109 = v69 << (uint(int32(3)) % 32)
					v110 = *(*float64)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_datan[0])))
					v113 = *(*float64)(unsafe.Add(mBase, uint32(v109)+uint32(_c_F_datan[1])))
					v116 = base.F64_sub(v110, base.F64_sub(base.F64_sub(base.F64_mul(v68, base.F64_add(v85, v102)), v113), v68))
					if v16 < int64(0) {
						v120 = base.F64_neg(v116)
					} else {
						v120 = v116
					}
					v121 = v120
					v131 = v121
				}
			}
		}
		if base.F64_eq(base.F64_abs(v131), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return int64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v138 = base.I64_reinterpret_f64(v131)
			return v138
		}
	} else {
		v138 = int64(9221120237041090560)
		return v138
	}
}
func F_datan2d(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v17 int64
	_ = v17
	var v23 int32
	_ = v23
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v36 int64
	_ = v36
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 float64
	_ = v59
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v95 float64
	_ = v95
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v127 float64
	_ = v127
	var v128 float64
	_ = v128
	var v137 float64
	_ = v137
	var v141 float64
	_ = v141
	var v144 float64
	_ = v144
	var v149 int64
	_ = v149
	var v159 int32
	_ = v159
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int64(9221120237041090560)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v12&int64(9223372036854775807)) {
		v149 = v11
		m.G0 = v9 + int32(16)
		return v149
	} else {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(v17&int64(9223372036854775807)) {
			v149 = v11
			m.G0 = v9 + int32(16)
			return v149
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_datan2d[0])))
			if v23 == int32(0) {
				F_init_degree_constants(m)
				mBase = m.M
			} else {
			}
			v27 = base.F64_reinterpret_i64(v12)
			v28 = base.F64_reinterpret_i64(v17)
			v36 = int64(9223372036854775807)
			if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v27)&v36) < base.Ui64(int64(9218868437227405313)))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v28)&v36) <= base.Ui64(int64(9218868437227405312))) == int32(0) {
				v137 = base.F64_add(v27, v28)
			} else {
				v49 = base.I64_reinterpret_f64(v28)
				v52 = base.I32_wrap_i64(int64(base.Ui64(v49) >> (uint(int64(32)) % 64)))
				v55 = base.I32_wrap_i64(v49)
				if v52-int32(1072693248)|v55 == int32(0) {
					v59 = F_atan(m, v27)
					mBase = m.M
					v137 = v59
				} else {
					v63 = int32(base.Ui32(v52)>>(uint(int32(30))%32)) & int32(2)
					v64 = base.I64_reinterpret_f64(v27)
					v68 = v63 | base.I32_wrap_i64(int64(base.Ui64(v64)>>(uint(int64(63))%64)))
					v73 = base.I32_wrap_i64(int64(base.Ui64(v64)>>(uint(int64(32))%64))) & int32(2147483647)
					if v73|base.I32_wrap_i64(v64) == int32(0) {
						switch v68 - int32(2) {
						case 0:
							v137 = float64(3.141592653589793)
						case 1:
							v137 = float64(-3.141592653589793)
						default:
							v128 = v27
							v137 = v128
						}
					} else {
						v83 = v52 & int32(2147483647)
						if v83|v55 == int32(0) {
							v137 = base.F64_copysign(float64(1.5707963267948966), v27)
						} else {
							if v83 == int32(2146435072) {
								if v73 != int32(2146435072) {
									v127 = *(*float64)(unsafe.Add(mBase, uint32(v68<<(uint(int32(3))%32))+uint32(_c_F_datan2d[1])))
									v128 = v127
									v137 = v128
								} else {
									v95 = *(*float64)(unsafe.Add(mBase, uint32(v68<<(uint(int32(3))%32))+uint32(_c_F_datan2d[2])))
									v137 = v95
								}
							} else {
								if base.B2i32(v73 != int32(2146435072))&base.B2i32(base.Ui32(v73) <= base.Ui32(v83+int32(67108864))) == int32(0) {
									v137 = base.F64_copysign(float64(1.5707963267948966), v27)
								} else {
									if v63 != 0 {
										if base.Ui32(v73+int32(67108864)) < base.Ui32(v83) {
											v113 = float64(0)
										} else {
											v112 = F_atan(m, base.F64_abs(base.F64_div(v27, v28)))
											mBase = m.M
											v113 = v112
										}
									} else {
										v112 = F_atan(m, base.F64_abs(base.F64_div(v27, v28)))
										mBase = m.M
										v113 = v112
									}
									switch v68 - int32(1) {
									case 0:
										v137 = base.F64_neg(v113)
									case 1:
										v137 = base.F64_sub(float64(3.141592653589793), base.F64_add(v113, float64(-1.2246467991473532e-16)))
									case 2:
										v137 = base.F64_add(base.F64_add(v113, float64(-1.2246467991473532e-16)), float64(-3.141592653589793))
									default:
										v128 = v113
										v137 = v128
									}
								}
							}
						}
					}
				}
			}
			*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v137
			v141 = *(*float64)(unsafe.Add(mBase, _c_F_datan2d[3]))
			v144 = base.F64_mul(base.F64_div(v137, v141), float64(45))
			if base.F64_eq(base.F64_abs(v144), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v159 = m.ExcPending
				if v159 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v149 = base.I64_reinterpret_f64(v144)
				m.G0 = v9 + int32(16)
				return v149
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
	if int32(357) <= v127+(v34-v126) {
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
							F_errfinish(m, int32(_a_F_db_encoding_convert_2), int32(491), int32(_a_F_db_encoding_convert_3))
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
func F_dcot(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v12 float64
	_ = v12
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
	var v54 int64
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v4&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		*(*int32)(unsafe.Add(mBase, _c_F_dcot[0])) = int32(0)
		v12 = base.F64_reinterpret_i64(v4)
		if base.F64_eq(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_dcot_0), int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_dcot_1), int32(1966), int32(_a_F_dcot_2))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int64(0)
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
			v28 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v12))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(v28) <= base.Ui32(int32(1072243195)) {
				if base.Ui32(v28) < base.Ui32(int32(1044381696)) {
					v45 = v12
				} else {
					v35 = F___tan(m, v12, float64(0), int32(0))
					mBase = m.M
					v45 = v35
				}
			} else {
				if base.Ui32(int32(2146435072)) <= base.Ui32(v28) {
					v45 = base.F64_sub(v12, v12)
				} else {
					v39 = F___rem_pio2(m, v12, v21)
					mBase = m.M
					v40 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
					v41 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
					v44 = F___tan(m, v40, v41, v39&int32(1))
					mBase = m.M
					v45 = v44
				}
			}
			m.G0 = v21 + int32(16)
			v54 = base.I64_reinterpret_f64(base.F64_div(float64(1), v45))
			return v54
		}
	} else {
		v54 = int64(9221120237041090560)
		return v54
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
	F_array_desc(m, l0, l1, int32(2), l2, int32(251), int32(0))
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
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
					v149 = int32(base.Ui32(v10) >> (uint(int32(1)) % 32))
					v151 = v149 + int32(3)
					v152 = F_palloc(m, v151)
					mBase = m.M
					v153 = m.ExcPending
					if v153 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v152))) = v151 << (uint(int32(2)) % 32)
						v158 = v149 - int32(1)
						if v158 == int32(0) {
							v167 = v152
						} else {
							base.MemoryCopy(m, v152+int32(4), l0+int32(1), v158)
							v167 = v152
						}
						m.G0 = v8 + int32(32)
						return v167
					}
				} else {
					v116 = F_detoast_external_attr(m, l0)
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return int32(0)
					} else {
						v167 = v116
						m.G0 = v8 + int32(32)
						return v167
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
						v167 = v18
						m.G0 = v8 + int32(32)
						return v167
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
									v167 = v47
									m.G0 = v8 + int32(32)
									return v167
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
									v167 = v30
									m.G0 = v8 + int32(32)
									return v167
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
					v167 = v52
					m.G0 = v8 + int32(32)
					return v167
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
							v167 = v81
						} else {
							base.MemoryCopy(m, v81, v52, v108)
							v167 = v81
						}
						m.G0 = v8 + int32(32)
						return v167
					}
				}
			}
		}
	} else {
		if v10&int32(3) == int32(2) {
			v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v124 = int32(base.Ui32(v122) >> (uint(int32(30)) % 32))
			switch v124 {
			case 0:
				v125 = F_pglz_decompress_datum(m, l0)
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					v167 = v125
					m.G0 = v8 + int32(32)
					return v167
				}
			case 1:
				v127 = F_lz4_decompress_datum(m)
				mBase = m.M
				v128 = m.ExcPending
				if v128 != 0 {
					return int32(0)
				} else {
					v167 = v127
					m.G0 = v8 + int32(32)
					return v167
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v124
					F_errmsg_internal(m, int32(_a_F_detoast_attr_0), v8+int32(16))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_detoast_attr_1), int32(489), int32(_a_F_detoast_attr_2))
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
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
				v149 = int32(base.Ui32(v10) >> (uint(int32(1)) % 32))
				v151 = v149 + int32(3)
				v152 = F_palloc(m, v151)
				mBase = m.M
				v153 = m.ExcPending
				if v153 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v152))) = v151 << (uint(int32(2)) % 32)
					v158 = v149 - int32(1)
					if v158 == int32(0) {
						v167 = v152
					} else {
						base.MemoryCopy(m, v152+int32(4), l0+int32(1), v158)
						v167 = v152
					}
					m.G0 = v8 + int32(32)
					return v167
				}
			} else {
				v167 = l0
				m.G0 = v8 + int32(32)
				return v167
			}
		}
	}
}
func F_dintdict_init(m *base.Module, l0 int32) int64 {
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
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L63
	}
L4:
	;
	m.G0 = v9 + int32(16)
	return base.I64_extend_i32_u(v14)
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
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v222
	F_errmsg(m, int32(_a_F_dintdict_init_6), v9)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_dintdict_init_3), int32(72), int32(_a_F_dintdict_init_4))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
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
func F_dpow(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 float64
	_ = v11
	var v12 int64
	_ = v12
	var v17 int64
	_ = v17
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v27 float64
	_ = v27
	var v35 float64
	_ = v35
	var v45 float64
	_ = v45
	var v49 float64
	_ = v49
	var v52 float64
	_ = v52
	var v58 float64
	_ = v58
	var v71 float64
	_ = v71
	var v74 float64
	_ = v74
	var v78 float64
	_ = v78
	var v84 float64
	_ = v84
	var v87 float64
	_ = v87
	var v107 float64
	_ = v107
	var v109 int32
	_ = v109
	var v110 float64
	_ = v110
	var v113 float64
	_ = v113
	var v116 float64
	_ = v116
	var v118 float64
	_ = v118
	var v123 int64
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = int64(9223372036854775807)
	v10 = v8 & v9
	v11 = base.F64_reinterpret_i64(v8)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v12&v9) {
		v17 = int64(9221120237041090560)
		if base.F64_ne(v11, float64(0)) != 0 {
			v22 = v17
		} else {
			v22 = int64(4607182418800017408)
		}
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(v10) {
			v25 = v17
		} else {
			v25 = v22
		}
		return v25
	} else {
		v27 = base.F64_reinterpret_i64(v12)
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v10) {
			if base.F64_ne(v27, float64(1)) != 0 {
				v123 = int64(9221120237041090560)
				return v123
			} else {
				return int64(4607182418800017408)
			}
		} else {
			v35 = float64(0)
			if base.F64_eq(v27, v35)&base.F64_lt(v11, v35) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v130 = m.ExcPending
				if v130 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(369361026))
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_dpow_0), int32(0))
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_dpow_1), int32(1563), int32(_a_F_dpow_2))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if base.F64_ne(base.F64_floor(v11), v11)&base.F64_lt(v27, float64(0)) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(369361026))
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_dpow_3), int32(0))
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_dpow_1), int32(1567), int32(_a_F_dpow_2))
								mBase = m.M
								v158 = m.ExcPending
								if v158 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v45 = base.F64_abs(v27)
					if base.F64_eq(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v49 = float64(1)
						if base.F64_eq(v45, v49) != 0 {
							v118 = v49
						} else {
							v52 = float64(0)
							if base.F64_gt(v11, v52) != 0 {
								if base.F64_gt(v45, float64(1)) != 0 {
									v58 = v11
								} else {
									v58 = float64(0)
								}
								v118 = v58
							} else {
								if base.F64_gt(v45, float64(1)) != 0 {
									v118 = v52
								} else {
									v118 = base.F64_neg(v11)
								}
							}
						}
						v123 = base.I64_reinterpret_f64(v118)
						return v123
					} else {
						if base.F64_eq(v45, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							if base.F64_eq(v11, float64(0)) != 0 {
								v118 = float64(1)
							} else {
								if base.F64_gt(v27, float64(0)) == int32(0) {
									v107 = base.F64_mul(v11, float64(0.5))
									v109 = base.F64_ne(base.F64_floor(v107), v107)
									if v109 != 0 {
										v110 = v27
									} else {
										v110 = base.F64_neg(v27)
									}
									if v109 != 0 {
										v113 = math.Float64frombits(uint64(0x8000000000000000))
									} else {
										v113 = float64(0)
									}
									if base.F64_gt(v11, float64(0)) != 0 {
										v116 = v110
									} else {
										v116 = v113
									}
									v118 = v116
								} else {
									v71 = float64(0)
									if base.F64_gt(v11, v71) != 0 {
										v74 = v27
									} else {
										v74 = v71
									}
									v118 = v74
								}
							}
							v123 = base.I64_reinterpret_f64(v118)
							return v123
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_dpow[0])) = int32(0)
							v78 = F_pow(m, v27, v11)
							mBase = m.M
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v78)&int64(9223372036854775807)) {
								v84 = float64(0)
								if base.F64_eq(v27, v84) != 0 {
									v118 = v84
									v123 = base.I64_reinterpret_f64(v118)
									return v123
								} else {
									v87 = float64(1)
									if base.F64_eq(v45, v87) != 0 {
										v118 = v87
										v123 = base.I64_reinterpret_f64(v118)
										return v123
									} else {
										if base.F64_ge(v11, float64(0)) != 0 {
											if base.F64_gt(v45, float64(1)) == int32(0) {
												F_float_underflow_error(m)
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												F_float_overflow_error(m)
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										} else {
											if base.F64_lt(v45, float64(1)) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												F_float_underflow_error(m)
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return int64(0)
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
								if base.F64_eq(base.F64_abs(v78), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if base.F64_ne(v78, float64(0)) != 0 {
										v118 = v78
										v123 = base.I64_reinterpret_f64(v118)
										return v123
									} else {
										if base.F64_ne(v27, float64(0)) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v118 = v78
											v123 = base.I64_reinterpret_f64(v118)
											return v123
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	v9 = m.G0
	v11 = v9 - int32(176)
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
	F_ScanKeyInit(m, v11, int32(1), int32(3), int32(184), int64(1259))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v11+int32(56), int32(2), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = int32(3)
	F_ScanKeyInit(m, v11+int32(112), v33, v33, int32(65), int64(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v43 = F_systable_beginscan(m, v15, int32(2673), int32(1), int32(0), int32(3), v11)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v45 = F_systable_getnext(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v45 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v47 = v45
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_systable_endscan(m, v43)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L21
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+22)))
	v57 = v55 + v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	if v58 != l1 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v69 = F_systable_getnext(m, v43)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L19
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v60 != l2 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v62 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57)+24)))
	if l3 != v63 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	F_simple_heap_delete(m, v15, v47+int32(4))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	if v69 != 0 {
		v47 = v69
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
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	m.G0 = v11 + int32(176)
	return
}
func F_dropconstraint_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v286 int32
	_ = v286
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	v19 = m.G0
	v21 = v19 - int32(304)
	m.G0 = v21
	F_check_stack_depth(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
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
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v31 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v35 = v33 + v34
	v37 = v35 + int32(4)
	if l5 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L154
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L151
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L147
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L142
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L139
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L134
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L130
	}
L15:
	;
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+104)))
	if int32(0) < v40 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+72)))
	if v43 == int32(110) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v48 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v214 = v43
	v224 = int32(0)
	goto L21
L21:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+106)))
	if v214&int32(255) != int32(102) {
		goto L61
	} else {
		goto L62
	}
L22:
	;
	v50 = F_extractNotNullColumn(m, l2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v54 = F_get_attname(m, v52, v50, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v57 = F_RelationGetIndexAttrBitmap(m, l1, int32(1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	v183 = F_RelationGetIndexAttrBitmap(m, l1, int32(2))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L50
	}
L26:
	;
	F_relation_close(m, v72, int32(1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L49
	}
L27:
	;
	if v57 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+119)))
	if v62 != int32(112) {
		goto L25
	} else {
		goto L31
	}
L29:
	;
	v115 = v57
	goto L30
L30:
	;
	v133 = F_bms_is_member(m, v50+int32(7), v115)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L42
	}
L31:
	;
	v66 = F_RelationGetPrimaryKeyIndex(m, l1, int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v66 == int32(0) {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v72 = F_relation_open(m, v66, int32(1))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+192))
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v74)+10)))
	if v75 <= int32(0) {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	v81 = int32(0)
	v84 = int32(0)
	v92 = v74
	goto L36
L36:
	;
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v92+v84<<(uint(int32(1))%32))+48)))
	v101 = F_bms_add_member(m, v81, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	F_relation_close(m, v72, int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v104 = v84 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v72)+192))
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+10)))
	if v104 < v106 {
		v81 = v101
		v84 = v104
		v92 = v105
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	if v101 == int32(0) {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v115 = v101
	goto L30
L42:
	;
	if v133 == int32(0) {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v146 = F_get_attname(m, v144, v50, int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v146
	F_errmsg(m, int32(_a_F_dropconstraint_internal_0), v21+int32(96))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_2), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
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
	v185 = F_bms_is_member(m, v50+int32(7), v183)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v185 != 0 {
		goto L13
	} else {
		goto L52
	}
L52:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v188 = F_SearchSysCacheCopyAttNum(m, v187, v50)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v188 == int32(0) {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+22)))
	v194 = v192 + v193
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+89)))
	if v195 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+86)))
	if v196 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v199 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v194)+86)) = uint8(v199)
	F_CatalogTupleUpdate(m, v48, v188+int32(4), v188)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	F_relation_close(m, v48, int32(3))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+72)))
	v214 = v208
	v224 = v54
	goto L21
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v255 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v254
	F_performDeletion(m, l0, l3, v255)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L71
	}
L62:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v35)+96))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v232 == v233 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v236 = F_table_open(m, v232, int32(8))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v236)+48))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+118)))
	if v239 == int32(116) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+24)))
	if v242 == int32(0) {
		goto L8
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_CheckTableNotInUse(m, v236, int32(_a_F_dropconstraint_internal_4))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	F_relation_close(m, v236, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L61
L71:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+72)))
	switch v261 - int32(99) {
	case 0, 11:
		goto L73
	default:
		goto L74
	}
L72:
	;
	F_relation_close(m, v31, int32(3))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L129
	}
L73:
	;
	if v227&int32(1) != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+119)))
	if v265 != int32(112) {
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
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v271 = F_find_inheritance_children(m, v270, l6)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v271 == int32(0) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	if v275 <= int32(0) {
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v286 = int32(0)
	goto L81
L81:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v302+v286<<(uint(int32(2))%32))))
	v308 = F_table_open(m, v306, int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L83
	}
L82:
	;
	goto L76
L83:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v308)+48))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+118)))
	if v311 == int32(116) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+24)))
	if v314 == int32(0) {
		goto L8
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	F_CheckTableNotInUse(m, v308, int32(_a_F_dropconstraint_internal_4))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+72)))
	if v320 == int32(110) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v375)+16))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+22)))
	v379 = v377 + v378
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+72)))
	switch v380 - int32(99) {
	case 0, 11:
		goto L106
	default:
		goto L107
	}
L90:
	;
	v323 = F_findNotNullConstraint(m, v306, v224)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v343 = v21 + int32(128)
	F_ScanKeyInit(m, v343, int32(9), int32(3), int32(184), base.I64_extend_i32_u(v306))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L98
	}
L93:
	;
	if v323 != 0 {
		v375 = v323
		goto L89
	} else {
		goto L94
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v308)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v224
	F_errmsg_internal(m, int32(_a_F_dropconstraint_internal_5), v21+int32(32))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_6), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
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
	F_ScanKeyInit(m, v21+int32(184), int32(10), int32(3), int32(184), int64(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_ScanKeyInit(m, v21+int32(240), int32(2), int32(3), int32(62), base.I64_extend_i32_u(v37))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v365 = F_systable_beginscan(m, v31, int32(2665), int32(1), int32(0), int32(3), v343)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v367 = F_systable_getnext(m, v365)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	if v367 == int32(0) {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v371 = F_heap_copytuple(m, v367)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_systable_endscan(m, v365)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v375 = v371
	goto L89
L106:
	;
	v396 = int32(*(*int16)(unsafe.Add(mBase, uint32(v379)+104)))
	if v396 <= int32(0) {
		goto L9
	} else {
		goto L111
	}
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errmsg_internal(m, int32(_a_F_dropconstraint_internal_7), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_8), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
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
	F_pfree(m, v375)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L126
	}
L113:
	;
	F_CatalogTupleUpdate(m, v31, v375+int32(4), v375)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L124
	}
L114:
	;
	if v396 != int32(1) {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	goto L116
L116:
	;
	v412 = v396 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v379)+104)) = uint16(v412)
	if v412&int32(_a_F_dropconstraint_internal_9) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L117:
	;
	v409 = v396 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v379)+104)) = uint16(v409)
	goto L113
L118:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+103)))
	if v401 != 0 {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v404 = int32(1)
	F_dropconstraint_internal(m, v21+int32(128), v308, v375, l3, v404, v404, l6)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L112
L121:
	;
	v418 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v379)+103)) = uint8(v418)
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
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L112
L126:
	;
	F_relation_close(m, v308, int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v434 = v286 + int32(1)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	if v434 < v435 {
		v286 = v434
		goto L81
	} else {
		goto L128
	}
L128:
	;
	goto L82
L129:
	;
	m.G0 = v21 + int32(304)
	return
L130:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = v486 + int32(4)
	F_errmsg(m, int32(_a_F_dropconstraint_internal_10), v21+int32(112))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_11), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
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
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v510 = F_get_attname(m, v508, v50, int32(0))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v510
	F_errmsg(m, int32(_a_F_dropconstraint_internal_12), v21)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_13), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
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
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v525
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v50
	F_errmsg_internal(m, int32(_a_F_dropconstraint_internal_14), v21+int32(16))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_15), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
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
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v547 = F_get_attname(m, v545, v50, int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v547
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v549 + int32(4)
	F_errmsg(m, int32(_a_F_dropconstraint_internal_16), v21+int32(80))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_17), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
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
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v308)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v571 + int32(4)
	F_errmsg(m, int32(_a_F_dropconstraint_internal_18), v21-int32(-64))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_19), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v379 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v306
	F_errmsg_internal(m, int32(_a_F_dropconstraint_internal_20), v21+int32(48))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_21), int32(_a_F_dropconstraint_internal_3))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
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
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errmsg(m, int32(_a_F_dropconstraint_internal_22), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_dropconstraint_internal_1), int32(_a_F_dropconstraint_internal_23), int32(_a_F_dropconstraint_internal_24))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
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
func F_dsimple_lexize(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v7 = F_str_tolower(m, v4, v5, int32(100))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		if v11 != 0 {
			v12 = F_searchstoplist(m, v3, v7)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				if v12 == int32(0) {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+8)))
					if v24 == int32(1) {
						v29 = F_palloc0_mul(m, int32(8), int32(2))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v7
							return base.I64_extend_i32_u(v29)
						}
					} else {
						F_pfree(m, v7)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							return int64(0)
						}
					}
				} else {
					F_pfree(m, v7)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int64(0)
					} else {
						v20 = F_palloc0_mul(m, int32(8), int32(2))
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v20)
						}
					}
				}
			}
		} else {
			F_pfree(m, v7)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				v20 = F_palloc0_mul(m, int32(8), int32(2))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v20)
				}
			}
		}
	}
}
func F_dsind(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v18 float64
	_ = v18
	var v23 int32
	_ = v23
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 float64
	_ = v43
	var v46 int64
	_ = v46
	var v53 float64
	_ = v53
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v145 int32
	_ = v145
	var v160 int64
	_ = v160
	var v168 float64
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 float64
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 float64
	_ = v182
	var v186 float64
	_ = v186
	var v190 float64
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v210 float64
	_ = v210
	var v214 int32
	_ = v214
	var v215 float64
	_ = v215
	var v216 float64
	_ = v216
	var v222 float64
	_ = v222
	var v223 float64
	_ = v223
	var v225 float64
	_ = v225
	var v227 float64
	_ = v227
	var v229 float64
	_ = v229
	var v238 float64
	_ = v238
	var v246 float64
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v266 float64
	_ = v266
	var v270 int32
	_ = v270
	var v271 float64
	_ = v271
	var v272 float64
	_ = v272
	var v277 float64
	_ = v277
	var v279 float64
	_ = v279
	var v281 float64
	_ = v281
	var v284 float64
	_ = v284
	var v288 float64
	_ = v288
	var v292 float64
	_ = v292
	var v296 float64
	_ = v296
	var v302 float64
	_ = v302
	var v303 float64
	_ = v303
	var v312 int64
	_ = v312
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v13&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L95
	} else {
		goto L100
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L95
	} else {
		goto L96
	}
L3:
	;
	v18 = base.F64_reinterpret_i64(v13)
	if base.F64_eq(base.F64_abs(v18), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	v312 = int64(9221120237041090560)
	goto L5
L5:
	;
	m.G0 = v10 + int32(16)
	return v312
L6:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsind[0])))
	if v23 == int32(0) {
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
	v34 = base.I64_reinterpret_f64(v18)
	v38 = int32(2047)
	v39 = base.I32_wrap_i64(int64(base.Ui64(v34)>>(uint(int64(52))%64))) & v38
	if v39 == v38 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v170 = base.F64_lt(v168, float64(0))
	if v170 != 0 {
		goto L51
	} else {
		goto L52
	}
L11:
	;
	v43 = base.F64_mul(v18, float64(360))
	v168 = base.F64_div(v43, v43)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v46 = v34 << (uint(int64(1)) % 64)
	if base.Ui64(v46) <= base.Ui64(int64(-9156662467374350336)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v46 == int64(-9156662467374350336) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if v39 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v53 = base.F64_mul(v18, float64(0))
	goto L19
L18:
	;
	v53 = v18
	goto L19
L19:
	;
	v168 = v53
	goto L10
L20:
	;
	if int32(1031) < v89 {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	v56 = int32(0)
	v58 = v34 << (uint(int64(12)) % 64)
	if int64(0) <= v58 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v89 = v39
	v90 = v34&int64(4503599627370495) | int64(4503599627370496)
	goto L20
L24:
	;
	v62 = v58
	v65 = v56
	goto L27
L25:
	;
	v76 = v56
	goto L26
L26:
	;
	v89 = v76
	v90 = v34 << (uint(base.I64_extend_i32_u(int32(1)-v76)) % 64)
	goto L20
L27:
	;
	v67 = v65 - int32(1)
	v69 = v62 << (uint(int64(1)) % 64)
	if int64(0) <= v69 {
		v62 = v69
		v65 = v67
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v76 = v67
	goto L26
L29:
	;
	goto L28
L30:
	;
	v94 = v90
	v97 = v89
	goto L33
L31:
	;
	v115 = v90
	v118 = v89
	goto L32
L32:
	;
	v120 = v115 - int64(6333186975989760)
	if v120 < int64(0) {
		v127 = v115
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v99 = v94 - int64(6333186975989760)
	if v99 < int64(0) {
		v106 = v94
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v115 = v108
	v118 = int32(1031)
	goto L32
L35:
	;
	v108 = v106 << (uint(int64(1)) % 64)
	v110 = v97 - int32(1)
	if int32(1031) < v110 {
		v94 = v108
		v97 = v110
		goto L33
	} else {
		goto L38
	}
L36:
	;
	if v99 != int64(0) {
		v106 = v99
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v168 = base.F64_mul(v18, float64(0))
	goto L10
L38:
	;
	goto L34
L39:
	;
	if base.Ui64(v127) <= base.Ui64(int64(4503599627370495)) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if v120 != int64(0) {
		v127 = v120
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v168 = base.F64_mul(v18, float64(0))
	goto L10
L42:
	;
	v131 = v127
	v134 = v118
	goto L45
L43:
	;
	v142 = v127
	v145 = v118
	goto L44
L44:
	;
	if int32(0) < v145 {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v136 = v134 - int32(1)
	v138 = v131 << (uint(int64(1)) % 64)
	if base.Ui64(v131) < base.Ui64(int64(2251799813685248)) {
		v131 = v138
		v134 = v136
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v142 = v138
	v145 = v136
	goto L44
L47:
	;
	goto L46
L48:
	;
	v160 = v142 - int64(4503599627370496) | base.I64_extend_i32_u(v145)<<(uint(int64(52))%64)
	goto L50
L49:
	;
	v160 = int64(base.Ui64(v142) >> (uint(base.I64_extend_i32_u(int32(1)-v145)) % 64))
	goto L50
L50:
	;
	v168 = base.F64_reinterpret_i64(v34&int64(-9223372036854775807-1) | v160)
	goto L10
L51:
	;
	v171 = int32(-1)
	goto L53
L52:
	;
	v171 = int32(1)
	goto L53
L53:
	;
	if v170 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v174 = base.F64_neg(v168)
	goto L56
L55:
	;
	v174 = v168
	goto L56
L56:
	;
	v176 = base.F64_gt(v174, float64(180))
	if v176 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v177 = int32(0) - v171
	goto L59
L58:
	;
	v177 = v171
	goto L59
L59:
	;
	if v176 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v303 = base.F64_mul(v302, base.F64_convert_i32_s(v177))
	if base.F64_eq(base.F64_abs(v303), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L94
	}
L61:
	;
	v182 = base.F64_sub(float64(360), v174)
	goto L63
L62:
	;
	v182 = v174
	goto L63
L63:
	;
	if base.F64_gt(v182, float64(90)) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v186 = base.F64_sub(float64(180), v182)
	goto L66
L65:
	;
	v186 = v182
	goto L66
L66:
	;
	if base.F64_le(v186, float64(30)) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v190 = base.F64_mul(v186, float64(0.017453292519943295))
	v194 = m.G0
	v196 = v194 - int32(16)
	m.G0 = v196
	v203 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v190))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v203) <= base.Ui32(int32(1072243195)) {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	goto L69
L69:
	;
	v246 = base.F64_mul(base.F64_sub(float64(90), v186), float64(0.017453292519943295))
	v250 = m.G0
	v252 = v250 - int32(16)
	m.G0 = v252
	v259 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v246))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v259) <= base.Ui32(int32(1072243195)) {
		goto L85
	} else {
		goto L86
	}
L70:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v229
	v238 = *(*float64)(unsafe.Add(mBase, _c_F_dsind[1]))
	v302 = base.F64_mul(base.F64_div(v229, v238), float64(0.5))
	goto L60
L71:
	;
	m.G0 = v196 + int32(16)
	goto L70
L72:
	;
	if base.Ui32(v203) < base.Ui32(int32(1045430272)) {
		v229 = v190
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v203) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v210 = F___sin(m, v190, float64(0), int32(0))
	mBase = m.M
	v229 = v210
	goto L71
L76:
	;
	v229 = base.F64_sub(v190, v190)
	goto L71
L77:
	;
	goto L78
L78:
	;
	v214 = F___rem_pio2(m, v190, v196)
	mBase = m.M
	v215 = *(*float64)(unsafe.Add(mBase, uint32(v196)+8))
	v216 = *(*float64)(unsafe.Add(mBase, uint32(v196)))
	switch v214&int32(3) - int32(1) {
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
	v227 = F___cos(m, v216, v215)
	mBase = m.M
	v229 = base.F64_neg(v227)
	goto L71
L80:
	;
	v225 = F___sin(m, v216, v215, int32(1))
	mBase = m.M
	v229 = base.F64_neg(v225)
	goto L71
L81:
	;
	v223 = F___cos(m, v216, v215)
	mBase = m.M
	v229 = v223
	goto L71
L82:
	;
	v222 = F___sin(m, v216, v215, int32(1))
	mBase = m.M
	v229 = v222
	goto L71
L83:
	;
	v292 = base.F64_sub(float64(1), v288)
	*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v292
	v296 = *(*float64)(unsafe.Add(mBase, _c_F_dsind[2]))
	v302 = base.F64_add(base.F64_mul(base.F64_div(v292, v296), float64(-0.5)), float64(1))
	goto L60
L84:
	;
	m.G0 = v252 + int32(16)
	goto L83
L85:
	;
	if base.Ui32(v259) < base.Ui32(int32(1044816030)) {
		v288 = float64(1)
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v259) {
		v288 = base.F64_sub(v246, v246)
		goto L84
	} else {
		goto L89
	}
L88:
	;
	v266 = F___cos(m, v246, float64(0))
	mBase = m.M
	v288 = v266
	goto L84
L89:
	;
	v270 = F___rem_pio2(m, v246, v252)
	mBase = m.M
	v271 = *(*float64)(unsafe.Add(mBase, uint32(v252)+8))
	v272 = *(*float64)(unsafe.Add(mBase, uint32(v252)))
	switch v270&int32(3) - int32(1) {
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
	v284 = F___sin(m, v272, v271, int32(1))
	mBase = m.M
	v288 = v284
	goto L84
L91:
	;
	v281 = F___cos(m, v272, v271)
	mBase = m.M
	v288 = base.F64_neg(v281)
	goto L84
L92:
	;
	v279 = F___sin(m, v272, v271, int32(1))
	mBase = m.M
	v288 = base.F64_neg(v279)
	goto L84
L93:
	;
	v277 = F___cos(m, v272, v271)
	mBase = m.M
	v288 = v277
	goto L84
L94:
	;
	v312 = base.I64_reinterpret_f64(v303)
	goto L5
L95:
	;
	return int64(0)
L96:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(_a_F_dsind_0), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_dsind_1), int32(2496), int32(_a_F_dsind_2))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L95
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
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dsqrt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.F64_lt(v4, float64(0)) == int32(0) {
		v9 = base.F64_sqrt(v4)
		v11 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_eq(base.F64_abs(v9), v11)&base.F64_ne(base.F64_abs(v4), v11) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return base.I64_reinterpret_f64(v9)
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(369361026))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_dsqrt_0), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_dsqrt_1), int32(1495), int32(_a_F_dsqrt_2))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
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
func F_dsynonym_init(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v2
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+30)) = uint16(v2)
	if v14 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L21
	} else {
		goto L82
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L21
	} else {
		goto L78
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if int32(0) < v21 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L21
	} else {
		goto L74
	}
L5:
	;
	v25 = int32(0)
	v29 = v2
	v32 = v2
	goto L8
L6:
	;
	v112 = v2
	v115 = v2
	goto L7
L7:
	;
	if v112 == int32(0) {
		goto L2
	} else {
		goto L33
	}
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v25<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v40 = int32(_a_F_dsynonym_init_0)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsynonym_init[0])))
	if base.B2i32(v43 == int32(0))|base.B2i32(v43 != v46) != 0 {
		v64 = v43
		v65 = v46
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v112 = v102
	v115 = v103
	goto L7
L10:
	;
	v105 = v25 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v105 < v106 {
		v25 = v105
		v29 = v102
		v32 = v103
		goto L8
	} else {
		goto L32
	}
L11:
	;
	if v64-v65 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	v49 = v39
	v50 = v40
	goto L14
L14:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v54 == int32(0) {
		v64 = v54
		v65 = v53
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v64 = v54
	v65 = v53
	goto L12
L16:
	;
	v57 = int32(1)
	if v54 == v53 {
		v49 = v49 + v57
		v50 = v50 + v57
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v69 = F_defGetString(m, v38)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v73 = int32(_a_F_dsynonym_init_1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsynonym_init[1])))
	if base.B2i32(v76 == int32(0))|base.B2i32(v76 != v79) != 0 {
		v97 = v76
		v98 = v79
		goto L24
	} else {
		goto L25
	}
L21:
	;
	return int64(0)
L22:
	;
	v102 = v69
	v103 = v32
	goto L10
L23:
	;
	if v97-v98 != 0 {
		goto L4
	} else {
		goto L30
	}
L24:
	;
	goto L23
L25:
	;
	v82 = v39
	v83 = v73
	goto L26
L26:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v87 == int32(0) {
		v97 = v87
		v98 = v86
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v97 = v87
	v98 = v86
	goto L24
L28:
	;
	v90 = int32(1)
	if v87 == v86 {
		v82 = v82 + v90
		v83 = v83 + v90
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v100 = F_defGetBoolean(m, v38)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v102 = v29
	v103 = v100
	goto L10
L32:
	;
	goto L9
L33:
	;
	v120 = v12 + int32(36)
	v122 = F_get_tsearch_config_filename(m, v112, int32(_a_F_dsynonym_init_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L21
	} else {
		goto L34
	}
L34:
	;
	v124 = F_tsearch_readline_begin(m, v120, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	if v124 == int32(0) {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v129 = F_palloc0(m, int32(12))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	v131 = F_tsearch_readline(m, v120)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L21
	} else {
		goto L39
	}
L38:
	;
	F_tsearch_readline_end(m, v12+int32(36))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L21
	} else {
		goto L71
	}
L39:
	;
	if v131 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v241 = int32(0)
	goto L38
L41:
	;
	goto L42
L42:
	;
	v137 = v131
	v140 = int32(0)
	goto L43
L43:
	;
	v147 = v12 + int32(32)
	v149 = F_findwrd(m, v137, v147, int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L21
	} else {
		goto L46
	}
L44:
	;
	v241 = v228
	goto L38
L45:
	;
	F_pfree(m, v137)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L21
	} else {
		goto L68
	}
L46:
	;
	if v149 == int32(0) {
		v228 = v140
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	if v154 == int32(0) {
		v228 = v140
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v157 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v153))) = uint8(v157)
	v163 = F_findwrd(m, v153+int32(1), v147, v12+int32(30))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	if v163 == int32(0) {
		v228 = v140
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v168)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if v170 <= v140 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v170 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L53
L53:
	;
	if v115&int32(1) != 0 {
		goto L61
	} else {
		goto L62
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v188
	goto L53
L55:
	;
	v174 = int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v174
	v178 = F_palloc_mul(m, int32(12), v174)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L21
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v181 = v170 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v185 = F_repalloc_mul(m, v183, int32(12), v181)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L21
	} else {
		goto L59
	}
L58:
	;
	v188 = v178
	goto L54
L59:
	;
	v188 = v185
	goto L54
L60:
	;
	v218 = v140 * int32(12)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v218+v219)+4)) = v216
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+30)))
	*(*uint16)(unsafe.Add(mBase, uint32(v222+v218)+8)) = uint16(v224)
	v228 = v140 + int32(1)
	goto L45
L61:
	;
	v193 = F_pstrdup(m, v149)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L21
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v202 = F_strlen(m, v149)
	mBase = m.M
	v204 = F_str_tolower(m, v149, v202, int32(100))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L21
	} else {
		goto L66
	}
L64:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v195+v140*int32(12)))) = v193
	v200 = F_pstrdup(m, v163)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L21
	} else {
		goto L65
	}
L65:
	;
	v216 = v200
	goto L60
L66:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v206+v140*int32(12)))) = v204
	v211 = F_strlen(m, v163)
	mBase = m.M
	v213 = F_str_tolower(m, v163, v211, int32(100))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L21
	} else {
		goto L67
	}
L67:
	;
	v216 = v213
	goto L60
L68:
	;
	v236 = F_tsearch_readline(m, v12+int32(36))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L21
	} else {
		goto L69
	}
L69:
	;
	if v236 != 0 {
		v137 = v236
		v140 = v228
		goto L43
	} else {
		goto L70
	}
L70:
	;
	goto L44
L71:
	;
	F_pfree(m, v122)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L21
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v241
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	F_pg_qsort(m, v254, v241, int32(12), int32(1263))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L21
	} else {
		goto L73
	}
L73:
	;
	v260 = v115 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v129)+8)) = uint8(v260)
	m.G0 = v12 + int32(80)
	return base.I64_extend_i32_u(v129)
L74:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L21
	} else {
		goto L75
	}
L75:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v274
	F_errmsg(m, int32(_a_F_dsynonym_init_3), v12+int32(16))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L21
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_dsynonym_init_4), int32(120), int32(_a_F_dsynonym_init_5))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L21
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L21
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(_a_F_dsynonym_init_6), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L21
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_dsynonym_init_4), int32(126), int32(_a_F_dsynonym_init_5))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L21
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
	F_errcode(m, int32(22))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L21
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v122
	F_errmsg(m, int32(_a_F_dsynonym_init_7), v12)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L21
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_dsynonym_init_4), int32(134), int32(_a_F_dsynonym_init_5))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L21
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dtrunc(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v10 float64
	_ = v10
	v3 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.F64_ge(v3, float64(0)) != 0 {
		v10 = base.F64_floor(v3)
	} else {
		v10 = base.F64_neg(base.F64_floor(base.F64_neg(v3)))
	}
	return base.I64_reinterpret_f64(v10)
}
func F_dutch_porter_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v524 int32
	_ = v524
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v860 int32
	_ = v860
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = v6
	goto L2
L1:
	;
	return v1052
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v13 <= v8 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	if v72 == v6 {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	goto L3
L5:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = v80
	goto L2
L6:
	;
	if v72 <= v70 {
		goto L4
	} else {
		goto L30
	}
L7:
	;
	v31 = F_find_among(m, l0, int32(_a_F_dutch_porter_ISO_8859_1_stem_0), int32(11), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
	v70 = v8
	v72 = v13
	goto L6
L9:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v8))))
	v18 = int32(224)
	if v17&v18 != v18 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if int32(1)<<(uint(v17)%32)&int32(340306450) != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	return int32(0)
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v35
	switch v31 - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	case 5:
		goto L14
	default:
		goto L5
	}
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v70 = v35
	v72 = v69
	goto L6
L15:
	;
	v65 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_ISO_8859_1_stem_1))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L28
	}
L16:
	;
	v59 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_ISO_8859_1_stem_2))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L26
	}
L17:
	;
	v53 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_ISO_8859_1_stem_3))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L24
	}
L18:
	;
	v47 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_ISO_8859_1_stem_4))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L22
	}
L19:
	;
	v41 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_ISO_8859_1_stem_5))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if int32(0) <= v41 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v1052 = v41
	goto L1
L22:
	;
	if int32(0) <= v47 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v1052 = v47
	goto L1
L24:
	;
	if int32(0) <= v53 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v1052 = v53
	goto L1
L26:
	;
	if int32(0) <= v59 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v1052 = v59
	goto L1
L28:
	;
	if int32(0) <= v65 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v1052 = v65
	goto L1
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v70 + int32(1)
	goto L5
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v109 < v108 {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v6))))
	if v86 != int32(121) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v89 = int32(1)
	v90 = v6 + v89
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v90
	v95 = F_slice_from_s(m, l0, v89, int32(_a_F_dutch_porter_ISO_8859_1_stem_6))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	if v95 < int32(0) {
		v1052 = v95
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	if int32(0) <= v149 {
		goto L51
	} else {
		goto L52
	}
L37:
	;
	v111 = v108
	goto L39
L38:
	;
	v111 = v109
	goto L39
L39:
	;
	v118 = v108
	goto L41
L40:
	;
	v149 = v129
	goto L36
L41:
	;
	if v118 == v111 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v149 = int32(-1)
	goto L36
L44:
	;
	goto L45
L45:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122+v118))))
	if int32(232) < v124 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v141 = v118 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v141
	v118 = v141
	goto L41
L47:
	;
	v126 = v124 - int32(97)
	if v126 < int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v129 = int32(1)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v126)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v133)>>(uint(v126&int32(7))%32))&v129 != 0 {
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L46
L51:
	;
	v153 = v149
	goto L54
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v306
	v310 = v6 + int32(3)
	if v306 < v310 {
		goto L102
	} else {
		goto L103
	}
L54:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v158 = v157 + v153
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v158
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v158 == v161 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v257 < v256 {
		goto L87
	} else {
		goto L88
	}
L57:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v158))))
	v167 = v165 - int32(105)
	if v167 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v237 = int32(1)
	v238 = v158 + v237
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v238
	v243 = F_slice_from_s(m, l0, v237, int32(_a_F_dutch_porter_ISO_8859_1_stem_7))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L12
	} else {
		goto L84
	}
L59:
	;
	if v167 == int32(16) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v171 = v158 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v171
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v183 < v171 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	goto L58
L63:
	;
	goto L56
L65:
	;
	if v226 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L66:
	;
	v185 = v171
	goto L68
L67:
	;
	v185 = v183
	goto L68
L68:
	;
	goto L70
L69:
	;
	v226 = v221
	goto L65
L70:
	;
	if v171 == v185 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v221 = int32(0)
	goto L69
L72:
	;
	v226 = int32(-1)
	goto L65
L73:
	;
	goto L74
L74:
	;
	v197 = int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v171))))
	if int32(232) < v200 {
		v221 = v197
		goto L69
	} else {
		goto L75
	}
L75:
	;
	v202 = v200 - int32(97)
	if v202 < int32(0) {
		v221 = v197
		goto L69
	} else {
		goto L76
	}
L76:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v202)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v208)>>(uint(v202&int32(7))%32))&int32(1) == int32(0) {
		v221 = v197
		goto L69
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v158 + int32(2)
	goto L78
L78:
	;
	goto L71
L79:
	;
	v231 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_ISO_8859_1_stem_8))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L12
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v171
	goto L56
L82:
	;
	if v231 < int32(0) {
		v1052 = v231
		goto L1
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	if v243 < int32(0) {
		v1052 = v243
		goto L1
	} else {
		goto L85
	}
L85:
	;
	goto L56
L86:
	;
	if int32(0) <= v297 {
		v153 = v297
		goto L54
	} else {
		goto L101
	}
L87:
	;
	v259 = v256
	goto L89
L88:
	;
	v259 = v257
	goto L89
L89:
	;
	v266 = v256
	goto L91
L90:
	;
	v297 = v277
	goto L86
L91:
	;
	if v266 == v259 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v297 = int32(-1)
	goto L86
L94:
	;
	goto L95
L95:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270+v266))))
	if int32(232) < v272 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v289 = v266 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v289
	v266 = v289
	goto L91
L97:
	;
	v274 = v272 - int32(97)
	if v274 < int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v277 = int32(1)
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v274)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v281)>>(uint(v274&int32(7))%32))&v277 != 0 {
		goto L90
	} else {
		goto L99
	}
L99:
	;
	goto L96
L101:
	;
	goto L55
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v541
	if v541 <= v6 {
		v644 = v310
		goto L169
	} else {
		goto L170
	}
L103:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v320 < v319 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v360 < int32(0) {
		goto L102
	} else {
		goto L119
	}
L105:
	;
	v322 = v319
	goto L107
L106:
	;
	v322 = v320
	goto L107
L107:
	;
	v329 = v319
	goto L109
L108:
	;
	v360 = v340
	goto L104
L109:
	;
	if v329 == v322 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v360 = int32(-1)
	goto L104
L112:
	;
	goto L113
L113:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333+v329))))
	if int32(232) < v335 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v352 = v329 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v352
	v329 = v352
	goto L109
L115:
	;
	v337 = v335 - int32(97)
	if v337 < int32(0) {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v340 = int32(1)
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v337)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v344)>>(uint(v337&int32(7))%32))&v340 != 0 {
		goto L108
	} else {
		goto L117
	}
L117:
	;
	goto L114
L119:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v364 = v363 + v360
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v364
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v375 < v364 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	if v418 < int32(0) {
		goto L102
	} else {
		goto L134
	}
L121:
	;
	v377 = v364
	goto L123
L122:
	;
	v377 = v375
	goto L123
L123:
	;
	v383 = v364
	goto L125
L124:
	;
	v418 = int32(1)
	goto L120
L125:
	;
	if v383 == v377 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v418 = int32(-1)
	goto L120
L128:
	;
	goto L129
L129:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390+v383))))
	if int32(232) < v392 {
		goto L124
	} else {
		goto L130
	}
L130:
	;
	v394 = v392 - int32(97)
	if v394 < int32(0) {
		goto L124
	} else {
		goto L131
	}
L131:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v394)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v400)>>(uint(v394&int32(7))%32))&int32(1) == int32(0) {
		goto L124
	} else {
		goto L132
	}
L132:
	;
	v409 = v383 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v409
	v383 = v409
	goto L125
L134:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v422 = v421 + v418
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v422
	if v310 < v422 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v425 = v422
	goto L137
L136:
	;
	v425 = v310
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v425
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v435 < v434 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if v475 < int32(0) {
		goto L102
	} else {
		goto L153
	}
L139:
	;
	v437 = v434
	goto L141
L140:
	;
	v437 = v435
	goto L141
L141:
	;
	v444 = v434
	goto L143
L142:
	;
	v475 = v455
	goto L138
L143:
	;
	if v444 == v437 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v475 = int32(-1)
	goto L138
L146:
	;
	goto L147
L147:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448+v444))))
	if int32(232) < v450 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v467 = v444 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v467
	v444 = v467
	goto L143
L149:
	;
	v452 = v450 - int32(97)
	if v452 < int32(0) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v455 = int32(1)
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v452)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v459)>>(uint(v452&int32(7))%32))&v455 != 0 {
		goto L142
	} else {
		goto L151
	}
L151:
	;
	goto L148
L153:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v479 = v478 + v475
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v479
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v490 < v479 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	if v533 < int32(0) {
		goto L102
	} else {
		goto L168
	}
L155:
	;
	v492 = v479
	goto L157
L156:
	;
	v492 = v490
	goto L157
L157:
	;
	v498 = v479
	goto L159
L158:
	;
	v533 = int32(1)
	goto L154
L159:
	;
	if v498 == v492 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v533 = int32(-1)
	goto L154
L162:
	;
	goto L163
L163:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505+v498))))
	if int32(232) < v507 {
		goto L158
	} else {
		goto L164
	}
L164:
	;
	v509 = v507 - int32(97)
	if v509 < int32(0) {
		goto L158
	} else {
		goto L165
	}
L165:
	;
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v509)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v515)>>(uint(v509&int32(7))%32))&int32(1) == int32(0) {
		goto L158
	} else {
		goto L166
	}
L166:
	;
	v524 = v498 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v524
	v498 = v524
	goto L159
L168:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v536 + v533
	goto L102
L169:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v645
	v647 = F_r_e_ending_1(m, l0)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L12
	} else {
		goto L201
	}
L170:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v547 = int32(1)
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545+v541-v547))))
	if base.B2i32(v549&int32(224) != int32(96))|base.B2i32(v547<<(uint(v549)%32)&int32(_a_F_dutch_porter_ISO_8859_1_stem_9) == int32(0)) != 0 {
		v644 = v310
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v564 = F_find_among_b(m, l0, int32(_a_F_dutch_porter_ISO_8859_1_stem_10), int32(5), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L12
	} else {
		goto L172
	}
L172:
	;
	if v564 == int32(0) {
		v644 = v310
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v568
	switch v564 - int32(1) {
	case 0:
		goto L176
	case 1:
		goto L175
	case 2:
		goto L174
	default:
		v644 = v310
		goto L169
	}
L174:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v568 < v586 {
		goto L184
	} else {
		goto L185
	}
L175:
	;
	v582 = F_r_en_ending_1(m, l0)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L12
	} else {
		goto L182
	}
L176:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v568 < v572 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v644 = int32(0)
	goto L169
L178:
	;
	goto L179
L179:
	;
	v577 = F_slice_from_s(m, l0, int32(4), int32(_a_F_dutch_porter_ISO_8859_1_stem_11))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L12
	} else {
		goto L180
	}
L180:
	;
	if v577 < int32(0) {
		v1052 = v577
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v644 = int32(1)
	goto L169
L182:
	;
	if int32(0) <= v582 {
		v644 = v582
		goto L169
	} else {
		goto L183
	}
L183:
	;
	v1052 = v582
	goto L1
L184:
	;
	v644 = int32(0)
	goto L169
L185:
	;
	goto L186
L186:
	;
	v589 = int32(1)
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L189
L187:
	;
	if v638 != 0 {
		v644 = v589
		goto L169
	} else {
		goto L199
	}
L188:
	;
	v638 = v635
	goto L187
L189:
	;
	if v597 <= v598 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v635 = int32(0)
	goto L188
L191:
	;
	v638 = int32(-1)
	goto L187
L192:
	;
	goto L193
L193:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609+v597-int32(1)))))
	if int32(232) < v613 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v597 - int32(1)
	goto L198
L195:
	;
	v615 = v613 - int32(97)
	if v615 < int32(0) {
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v618 = int32(1)
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v615)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v622)>>(uint(v615&int32(7))%32))&v618 != 0 {
		v635 = v618
		goto L188
	} else {
		goto L197
	}
L197:
	;
	goto L194
L198:
	;
	goto L190
L199:
	;
	v639 = F_slice_del(m, l0)
	mBase = m.M
	if v639 < int32(0) {
		v1052 = v639
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v644 = v589
	goto L169
L201:
	;
	if v647 < int32(0) {
		v1052 = v647
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v651
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v651
	v654 = int32(4)
	v656 = int32(0)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v651-v659 < v654 {
		v669 = v656
		goto L210
	} else {
		goto L211
	}
L203:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v995
	v998 = v995
	goto L310
L204:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v850
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L280
L205:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v747 < v839 {
		goto L204
	} else {
		goto L275
	}
L206:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v747 < v834 {
		goto L204
	} else {
		goto L273
	}
L207:
	;
	if int32(0) <= v829 {
		goto L203
	} else {
		goto L272
	}
L208:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v719
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v719
	v723 = v719 - int32(1)
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v723 <= v724 {
		goto L204
	} else {
		goto L233
	}
L209:
	;
	if v669 == int32(0) {
		v718 = v644
		goto L208
	} else {
		goto L213
	}
L210:
	;
	goto L209
L211:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v665 = F_memcmp(m, v662+v651-v654, int32(_a_F_dutch_porter_ISO_8859_1_stem_12), v654)
	mBase = m.M
	if v665 != 0 {
		v669 = v656
		goto L210
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v651 - v654
	v669 = int32(1)
	goto L210
L213:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v672
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v672 < v674 {
		v718 = v644
		goto L208
	} else {
		goto L214
	}
L214:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v676 < v672 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678+v672-int32(1)))))
	if v682 == int32(99) {
		v718 = v644
		goto L208
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v685 = F_slice_del(m, l0)
	mBase = m.M
	if v685 < int32(0) {
		v1052 = v685
		goto L1
	} else {
		goto L219
	}
L218:
	;
	goto L217
L219:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v688
	v690 = int32(2)
	v692 = int32(0)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v688-v695 < v690 {
		v705 = v692
		goto L221
	} else {
		goto L222
	}
L220:
	;
	if v705 == int32(0) {
		v718 = v644
		goto L208
	} else {
		goto L224
	}
L221:
	;
	goto L220
L222:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v701 = F_memcmp(m, v698+v688-v690, int32(_a_F_dutch_porter_ISO_8859_1_stem_13), v690)
	mBase = m.M
	if v701 != 0 {
		v705 = v692
		goto L221
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v688 - v690
	v705 = int32(1)
	goto L221
L224:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v708
	v710 = F_r_en_ending_1(m, l0)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L12
	} else {
		goto L225
	}
L225:
	;
	v713 = base.B2i32(v710 < int32(0))
	if v710 < int32(0) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v714 = v710
	goto L228
L227:
	;
	v714 = v644
	goto L228
L228:
	;
	if v710 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v715 = v714
	goto L231
L230:
	;
	v715 = v644
	goto L231
L231:
	;
	if v710 < int32(0) {
		v829 = v715
		goto L207
	} else {
		goto L232
	}
L232:
	;
	v718 = v715
	goto L208
L233:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726+v723))))
	if base.B2i32(v728&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v728)%32)&int32(_a_F_dutch_porter_ISO_8859_1_stem_14) == int32(0)) != 0 {
		goto L204
	} else {
		goto L234
	}
L234:
	;
	v743 = F_find_among_b(m, l0, int32(_a_F_dutch_porter_ISO_8859_1_stem_15), int32(6), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L12
	} else {
		goto L235
	}
L235:
	;
	if v743 == int32(0) {
		goto L204
	} else {
		goto L236
	}
L236:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v747
	switch v743 - int32(1) {
	case 0:
		goto L239
	case 1:
		goto L238
	case 2:
		goto L237
	case 3:
		goto L206
	case 4:
		goto L205
	default:
		goto L204
	}
L237:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v747 < v816 {
		goto L204
	} else {
		goto L262
	}
L238:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v747 < v802 {
		goto L204
	} else {
		goto L256
	}
L239:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v747 < v751 {
		goto L204
	} else {
		goto L240
	}
L240:
	;
	v753 = F_slice_del(m, l0)
	mBase = m.M
	if v753 < int32(0) {
		v1052 = v753
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v756
	v758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v759 = int32(2)
	v761 = int32(0)
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v756-v764 < v759 {
		v774 = v761
		goto L244
	} else {
		goto L245
	}
L242:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v794 + (v756 - v758)
	v798 = F_r_undouble_1(m, l0)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L12
	} else {
		goto L254
	}
L243:
	;
	if v774 == int32(0) {
		goto L242
	} else {
		goto L247
	}
L244:
	;
	goto L243
L245:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v770 = F_memcmp(m, v767+v756-v759, int32(_a_F_dutch_porter_ISO_8859_1_stem_16), v759)
	mBase = m.M
	if v770 != 0 {
		v774 = v761
		goto L244
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v756 - v759
	v774 = int32(1)
	goto L244
L247:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v777
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v777 < v779 {
		goto L242
	} else {
		goto L248
	}
L248:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v781 < v777 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783+v777-int32(1)))))
	if v787 == int32(101) {
		goto L242
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v790 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v790 {
		goto L204
	} else {
		goto L253
	}
L252:
	;
	goto L251
L253:
	;
	v1052 = v790
	goto L1
L254:
	;
	if int32(0) <= v798 {
		goto L204
	} else {
		goto L255
	}
L255:
	;
	v1052 = v798
	goto L1
L256:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v804 < v747 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806+v747-int32(1)))))
	if v810 == int32(101) {
		goto L204
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v813 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v813 {
		goto L204
	} else {
		goto L261
	}
L260:
	;
	goto L259
L261:
	;
	v1052 = v813
	goto L1
L262:
	;
	v818 = F_slice_del(m, l0)
	mBase = m.M
	if v818 < int32(0) {
		v1052 = v818
		goto L1
	} else {
		goto L263
	}
L263:
	;
	v821 = F_r_e_ending_1(m, l0)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L12
	} else {
		goto L264
	}
L264:
	;
	if int32(0) <= v821 {
		goto L204
	} else {
		goto L265
	}
L265:
	;
	if v821 < int32(0) {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v827 = v821
	goto L268
L267:
	;
	v827 = v718
	goto L268
L268:
	;
	if v821 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v828 = v827
	goto L271
L270:
	;
	v828 = v718
	goto L271
L271:
	;
	v829 = v828
	goto L207
L272:
	;
	v1052 = v829
	goto L1
L273:
	;
	v836 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v836 {
		goto L204
	} else {
		goto L274
	}
L274:
	;
	v1052 = v836
	goto L1
L275:
	;
	v841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v841 == int32(0) {
		goto L204
	} else {
		goto L276
	}
L276:
	;
	v844 = F_slice_del(m, l0)
	mBase = m.M
	if v844 < int32(0) {
		v1052 = v844
		goto L1
	} else {
		goto L277
	}
L277:
	;
	goto L204
L278:
	;
	if v900 != 0 {
		goto L203
	} else {
		goto L290
	}
L279:
	;
	v900 = v897
	goto L278
L280:
	;
	if v850 <= v860 {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	v897 = int32(0)
	goto L279
L282:
	;
	v900 = int32(-1)
	goto L278
L283:
	;
	goto L284
L284:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871+v850-int32(1)))))
	if int32(232) < v875 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v850 - int32(1)
	goto L289
L286:
	;
	v877 = v875 - int32(73)
	if v877 < int32(0) {
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v880 = int32(1)
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v877)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v884)>>(uint(v877&int32(7))%32))&v880 != 0 {
		v897 = v880
		goto L279
	} else {
		goto L288
	}
L288:
	;
	goto L285
L289:
	;
	goto L281
L290:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v903 = v901 - int32(1)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v903 <= v904 {
		goto L203
	} else {
		goto L291
	}
L291:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906+v903))))
	if base.B2i32(v908&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v908)%32)&int32(_a_F_dutch_porter_ISO_8859_1_stem_17) == int32(0)) != 0 {
		goto L203
	} else {
		goto L292
	}
L292:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v924 = F_find_among_b(m, l0, int32(_a_F_dutch_porter_ISO_8859_1_stem_18), int32(4), int32(0))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L12
	} else {
		goto L293
	}
L293:
	;
	if v924 == int32(0) {
		goto L203
	} else {
		goto L294
	}
L294:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L297
L295:
	;
	if v976 != 0 {
		goto L203
	} else {
		goto L307
	}
L296:
	;
	v976 = v973
	goto L295
L297:
	;
	if v935 <= v936 {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v973 = int32(0)
	goto L296
L299:
	;
	v976 = int32(-1)
	goto L295
L300:
	;
	goto L301
L301:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947+v935-int32(1)))))
	if int32(232) < v951 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v935 - int32(1)
	goto L306
L303:
	;
	v953 = v951 - int32(97)
	if v953 < int32(0) {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v956 = int32(1)
	v960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v953)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v960)>>(uint(v953&int32(7))%32))&v956 != 0 {
		v973 = v956
		goto L296
	} else {
		goto L305
	}
L305:
	;
	goto L302
L306:
	;
	goto L298
L307:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v979 = v977 + (v901 - v920)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v979
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v979
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v979 <= v982 {
		goto L203
	} else {
		goto L308
	}
L308:
	;
	v985 = v979 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v985
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v985
	v988 = F_slice_del(m, l0)
	mBase = m.M
	if v988 < int32(0) {
		v1052 = v988
		goto L1
	} else {
		goto L309
	}
L309:
	;
	goto L203
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v998
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1003 <= v998 {
		goto L315
	} else {
		goto L316
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v995
	v1052 = int32(1)
	goto L1
L312:
	;
	goto L311
L313:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v998 = v1048
	goto L310
L314:
	;
	if v1039 <= v1038 {
		goto L312
	} else {
		goto L329
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v998
	v1038 = v998
	v1039 = v1003
	goto L314
L316:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005+v998))))
	v1009 = v1007 - int32(73)
	if v1009 != int32(16) {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1013 = v1009
	goto L319
L318:
	;
	v1013 = int32(0)
	goto L319
L319:
	;
	if v1013 != 0 {
		goto L315
	} else {
		goto L320
	}
L320:
	;
	v1017 = F_find_among(m, l0, int32(_a_F_dutch_porter_ISO_8859_1_stem_19), int32(3), int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L12
	} else {
		goto L321
	}
L321:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1019
	switch v1017 - int32(1) {
	case 0:
		goto L323
	case 1:
		goto L322
	case 2:
		goto L324
	default:
		goto L313
	}
L322:
	;
	v1032 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_ISO_8859_1_stem_20))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L12
	} else {
		goto L327
	}
L323:
	;
	v1026 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_ISO_8859_1_stem_21))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L12
	} else {
		goto L325
	}
L324:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1038 = v1019
	v1039 = v1023
	goto L314
L325:
	;
	if int32(0) <= v1026 {
		goto L313
	} else {
		goto L326
	}
L326:
	;
	v1052 = v1026
	goto L1
L327:
	;
	if int32(0) <= v1032 {
		goto L313
	} else {
		goto L328
	}
L328:
	;
	v1052 = v1032
	goto L1
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1038 + int32(1)
	goto L313
}
func F_dutch_porter_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v228 int32
	_ = v228
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v369 int32
	_ = v369
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v510 int32
	_ = v510
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v618 int32
	_ = v618
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v694 int32
	_ = v694
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v727 int32
	_ = v727
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v765 int32
	_ = v765
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v816 int32
	_ = v816
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v850 int32
	_ = v850
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v891 int32
	_ = v891
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v942 int32
	_ = v942
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v968 int32
	_ = v968
	var v975 int32
	_ = v975
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1013 int32
	_ = v1013
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1064 int32
	_ = v1064
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1098 int32
	_ = v1098
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1260 int32
	_ = v1260
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1272 int32
	_ = v1272
	var v1289 int32
	_ = v1289
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1508 int32
	_ = v1508
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1609 int32
	_ = v1609
	var v1615 int32
	_ = v1615
	var v1632 int32
	_ = v1632
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1728 int32
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1750 int32
	_ = v1750
	var v1752 int32
	_ = v1752
	var v1760 int32
	_ = v1760
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1772 int32
	_ = v1772
	var v1789 int32
	_ = v1789
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1825 int32
	_ = v1825
	var v1830 int32
	_ = v1830
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1855 int32
	_ = v1855
	var v1860 int32
	_ = v1860
	var v1867 int32
	_ = v1867
	var v1870 int32
	_ = v1870
	var v1875 int32
	_ = v1875
	var v1877 int32
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1885 int32
	_ = v1885
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1936 int32
	_ = v1936
	var v1939 int32
	_ = v1939
	var v1943 int32
	_ = v1943
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1965 int32
	_ = v1965
	var v1972 int32
	_ = v1972
	var v1976 int32
	_ = v1976
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = v6
	goto L2
L1:
	;
	return v1976
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v8
	v14 = v8 + int32(1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 <= v14 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v6 == v137 {
		goto L51
	} else {
		goto L52
	}
L4:
	;
	goto L3
L5:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = v134
	goto L2
L6:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L32
L7:
	;
	v33 = F_find_among(m, l0, int32(_a_F_dutch_porter_UTF_8_stem_0), int32(11), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
	v72 = v8
	v74 = v15
	goto L6
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v14))))
	if v19&int32(224) != int32(160) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if int32(1)<<(uint(v19)%32)&int32(340306450) != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	return int32(0)
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
	switch v33 - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	case 5:
		goto L14
	default:
		goto L5
	}
L14:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v72 = v37
	v74 = v71
	goto L6
L15:
	;
	v67 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_UTF_8_stem_1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L28
	}
L16:
	;
	v61 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_UTF_8_stem_2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L12
	} else {
		goto L26
	}
L17:
	;
	v55 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_UTF_8_stem_3))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L24
	}
L18:
	;
	v49 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_UTF_8_stem_4))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L12
	} else {
		goto L22
	}
L19:
	;
	v43 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_UTF_8_stem_5))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if int32(0) <= v43 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v1976 = v43
	goto L1
L22:
	;
	if int32(0) <= v49 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v1976 = v49
	goto L1
L24:
	;
	if int32(0) <= v55 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v1976 = v55
	goto L1
L26:
	;
	if int32(0) <= v61 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v1976 = v61
	goto L1
L28:
	;
	if int32(0) <= v67 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v1976 = v67
	goto L1
L30:
	;
	if v127 < int32(0) {
		goto L4
	} else {
		goto L50
	}
L32:
	;
	goto L33
L33:
	;
	goto L34
L34:
	;
	v82 = v72
	v84 = int32(1)
	goto L37
L36:
	;
	v127 = v112
	goto L30
L37:
	;
	if v74 <= v82 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L36
L39:
	;
	v127 = int32(-1)
	goto L30
L40:
	;
	goto L41
L41:
	;
	v89 = v82 + int32(1)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v82))))
	if base.Ui32(v91) < base.Ui32(int32(192)) {
		v112 = v89
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v113 = int32(1)
	if v113 < v84 {
		v82 = v112
		v84 = v84 - v113
		goto L37
	} else {
		goto L49
	}
L43:
	;
	if v74 <= v89 {
		v112 = v89
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v98 = v89
	goto L45
L45:
	;
	v101 = int32(*(*int8)(unsafe.Add(mBase, uint32(v75+v98))))
	if int32(-65) < v101 {
		v112 = v98
		goto L42
	} else {
		goto L47
	}
L46:
	;
	v112 = v74
	goto L42
L47:
	;
	v105 = v98 + int32(1)
	if v105 != v74 {
		v98 = v105
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	goto L38
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v127
	goto L5
L51:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v177 = v167
	goto L58
L52:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139+v6))))
	if v141 != int32(121) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v144 = int32(1)
	v145 = v6 + v144
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v145
	v150 = F_slice_from_s(m, l0, v144, int32(_a_F_dutch_porter_UTF_8_stem_6))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	if v150 < int32(0) {
		v1976 = v150
		goto L1
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	if int32(0) <= v272 {
		goto L81
	} else {
		goto L82
	}
L57:
	;
	v272 = v244
	goto L56
L58:
	;
	if v168 <= v177 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v272 = int32(-1)
	goto L56
L61:
	;
	goto L62
L62:
	;
	v184 = int32(1)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v169))))
	if base.Ui32(v186) < base.Ui32(int32(192)) {
		v243 = v186
		v244 = v184
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if int32(232) < v243 {
		goto L76
	} else {
		goto L77
	}
L64:
	;
	v190 = v177 + int32(1)
	if v190 == v168 {
		v243 = v186
		v244 = v184
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190+v169))))
	v195 = v193 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v186) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199+v169))))
	v211 = v209 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v186) {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v199 = v177 + int32(2)
	if v199 != v168 {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v243 = v186<<(uint(int32(6))%32)&int32(1984) | v195
	v244 = int32(2)
	goto L63
L70:
	;
	goto L69
L71:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v215))))
	v243 = v228&int32(63) | (v186<<(uint(int32(18))%32)&int32(_a_F_dutch_porter_UTF_8_stem_7) | v195<<(uint(int32(12))%32) | v211<<(uint(int32(6))%32))
	v244 = int32(4)
	goto L63
L72:
	;
	v215 = v177 + int32(3)
	if v215 != v168 {
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v243 = v186<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_8) | v195<<(uint(int32(6))%32) | v211
	v244 = int32(3)
	goto L63
L75:
	;
	goto L74
L76:
	;
	v261 = v244 + v177
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v261
	v177 = v261
	goto L58
L77:
	;
	v248 = v243 - int32(97)
	if v248 < int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v248)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v254)>>(uint(v248&int32(7))%32))&int32(1) != 0 {
		goto L57
	} else {
		goto L79
	}
L79:
	;
	goto L76
L81:
	;
	v276 = v272
	v278 = v74
	goto L84
L82:
	;
	v560 = v74
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v563
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v563
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L155
L84:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v281 = v280 + v276
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v281
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v281
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v281 == v284 {
		v437 = v278
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v560 = v437
	goto L83
L86:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v459 = v449
	goto L128
L87:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+v281))))
	v290 = v288 - int32(105)
	if v290 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v425 = int32(1)
	v426 = v281 + v425
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v426
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v426
	v431 = F_slice_from_s(m, l0, v425, int32(_a_F_dutch_porter_UTF_8_stem_9))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L12
	} else {
		goto L124
	}
L89:
	;
	if v290 == int32(16) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v294 = v281 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v294
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L97
L92:
	;
	goto L88
L93:
	;
	v437 = v278
	goto L86
L95:
	;
	if v414 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L96:
	;
	v414 = v407
	goto L95
L97:
	;
	if v309 <= v294 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v407 = int32(0)
	goto L96
L99:
	;
	v414 = int32(-1)
	goto L95
L100:
	;
	goto L101
L101:
	;
	v325 = int32(1)
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v310))))
	if base.Ui32(v327) < base.Ui32(int32(192)) {
		v384 = v327
		v385 = v325
		goto L102
	} else {
		goto L103
	}
L102:
	;
	if int32(232) < v384 {
		v407 = v385
		goto L96
	} else {
		goto L115
	}
L103:
	;
	v331 = v281 + int32(2)
	if v331 == v309 {
		v384 = v327
		v385 = v325
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+v310))))
	v336 = v334 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v327) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340+v310))))
	v352 = v350 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v327) {
		goto L111
	} else {
		goto L112
	}
L106:
	;
	v340 = v281 + int32(3)
	if v340 != v309 {
		goto L105
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v384 = v327<<(uint(int32(6))%32)&int32(1984) | v336
	v385 = int32(2)
	goto L102
L109:
	;
	goto L108
L110:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310+v356))))
	v384 = v369&int32(63) | (v327<<(uint(int32(18))%32)&int32(_a_F_dutch_porter_UTF_8_stem_7) | v336<<(uint(int32(12))%32) | v352<<(uint(int32(6))%32))
	v385 = int32(4)
	goto L102
L111:
	;
	v356 = v281 + int32(4)
	if v356 != v309 {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v384 = v327<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_8) | v336<<(uint(int32(6))%32) | v352
	v385 = int32(3)
	goto L102
L114:
	;
	goto L113
L115:
	;
	v389 = v384 - int32(97)
	if v389 < int32(0) {
		v407 = v385
		goto L96
	} else {
		goto L116
	}
L116:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v389)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v395)>>(uint(v389&int32(7))%32))&int32(1) == int32(0) {
		v407 = v385
		goto L96
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 + v294
	goto L118
L118:
	;
	goto L98
L119:
	;
	v419 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_UTF_8_stem_10))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L12
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v294
	v437 = v294
	goto L86
L122:
	;
	if v419 < int32(0) {
		v1976 = v419
		goto L1
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	if v431 < int32(0) {
		v1976 = v431
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v437 = v278
	goto L86
L126:
	;
	if int32(0) <= v554 {
		v276 = v554
		v278 = v437
		goto L84
	} else {
		goto L151
	}
L127:
	;
	v554 = v526
	goto L126
L128:
	;
	if v450 <= v459 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v554 = int32(-1)
	goto L126
L131:
	;
	goto L132
L132:
	;
	v466 = int32(1)
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459+v451))))
	if base.Ui32(v468) < base.Ui32(int32(192)) {
		v525 = v468
		v526 = v466
		goto L133
	} else {
		goto L134
	}
L133:
	;
	if int32(232) < v525 {
		goto L146
	} else {
		goto L147
	}
L134:
	;
	v472 = v459 + int32(1)
	if v472 == v450 {
		v525 = v468
		v526 = v466
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472+v451))))
	v477 = v475 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v468) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481+v451))))
	v493 = v491 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v468) {
		goto L142
	} else {
		goto L143
	}
L137:
	;
	v481 = v459 + int32(2)
	if v481 != v450 {
		goto L136
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v525 = v468<<(uint(int32(6))%32)&int32(1984) | v477
	v526 = int32(2)
	goto L133
L140:
	;
	goto L139
L141:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451+v497))))
	v525 = v510&int32(63) | (v468<<(uint(int32(18))%32)&int32(_a_F_dutch_porter_UTF_8_stem_7) | v477<<(uint(int32(12))%32) | v493<<(uint(int32(6))%32))
	v526 = int32(4)
	goto L133
L142:
	;
	v497 = v459 + int32(3)
	if v497 != v450 {
		goto L141
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v525 = v468<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_8) | v477<<(uint(int32(6))%32) | v493
	v526 = int32(3)
	goto L133
L145:
	;
	goto L144
L146:
	;
	v543 = v526 + v459
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v543
	v459 = v543
	goto L128
L147:
	;
	v530 = v525 - int32(97)
	if v530 < int32(0) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v530)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v536)>>(uint(v530&int32(7))%32))&int32(1) != 0 {
		goto L127
	} else {
		goto L149
	}
L149:
	;
	goto L146
L151:
	;
	goto L85
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1118
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1118
	if v1118 <= v6 {
		v1302 = v1116
		goto L279
	} else {
		goto L280
	}
L153:
	;
	if v618 < int32(0) {
		v1116 = v560
		goto L152
	} else {
		goto L173
	}
L155:
	;
	goto L156
L156:
	;
	goto L157
L157:
	;
	v573 = v6
	v575 = int32(3)
	goto L160
L159:
	;
	v618 = v603
	goto L153
L160:
	;
	if v563 <= v573 {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	goto L159
L162:
	;
	v618 = int32(-1)
	goto L153
L163:
	;
	goto L164
L164:
	;
	v580 = v573 + int32(1)
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566+v573))))
	if base.Ui32(v582) < base.Ui32(int32(192)) {
		v603 = v580
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v604 = int32(1)
	if v604 < v575 {
		v573 = v603
		v575 = v575 - v604
		goto L160
	} else {
		goto L172
	}
L166:
	;
	if v563 <= v580 {
		v603 = v580
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v589 = v580
	goto L168
L168:
	;
	v592 = int32(*(*int8)(unsafe.Add(mBase, uint32(v566+v589))))
	if int32(-65) < v592 {
		v603 = v589
		goto L165
	} else {
		goto L170
	}
L169:
	;
	v603 = v563
	goto L165
L170:
	;
	v596 = v589 + int32(1)
	if v596 != v563 {
		v589 = v596
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	goto L161
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v643 = v6
	goto L176
L174:
	;
	if v738 < int32(0) {
		v1116 = v738
		goto L152
	} else {
		goto L199
	}
L175:
	;
	v738 = v710
	goto L174
L176:
	;
	if v634 <= v643 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v738 = int32(-1)
	goto L174
L179:
	;
	goto L180
L180:
	;
	v650 = int32(1)
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643+v635))))
	if base.Ui32(v652) < base.Ui32(int32(192)) {
		v709 = v652
		v710 = v650
		goto L181
	} else {
		goto L182
	}
L181:
	;
	if int32(232) < v709 {
		goto L194
	} else {
		goto L195
	}
L182:
	;
	v656 = v643 + int32(1)
	if v656 == v634 {
		v709 = v652
		v710 = v650
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656+v635))))
	v661 = v659 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v652) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665+v635))))
	v677 = v675 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v652) {
		goto L190
	} else {
		goto L191
	}
L185:
	;
	v665 = v643 + int32(2)
	if v665 != v634 {
		goto L184
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v709 = v652<<(uint(int32(6))%32)&int32(1984) | v661
	v710 = int32(2)
	goto L181
L188:
	;
	goto L187
L189:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635+v681))))
	v709 = v694&int32(63) | (v652<<(uint(int32(18))%32)&int32(_a_F_dutch_porter_UTF_8_stem_7) | v661<<(uint(int32(12))%32) | v677<<(uint(int32(6))%32))
	v710 = int32(4)
	goto L181
L190:
	;
	v681 = v643 + int32(3)
	if v681 != v634 {
		goto L189
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v709 = v652<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_8) | v661<<(uint(int32(6))%32) | v677
	v710 = int32(3)
	goto L181
L193:
	;
	goto L192
L194:
	;
	v727 = v710 + v643
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v727
	v643 = v727
	goto L176
L195:
	;
	v714 = v709 - int32(97)
	if v714 < int32(0) {
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v714)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v720)>>(uint(v714&int32(7))%32))&int32(1) != 0 {
		goto L175
	} else {
		goto L197
	}
L197:
	;
	goto L194
L199:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v742 = v741 + v738
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v742
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v765 = v742
	goto L202
L200:
	;
	if v861 < int32(0) {
		v1116 = v861
		goto L152
	} else {
		goto L224
	}
L201:
	;
	v861 = v832
	goto L200
L202:
	;
	if v756 <= v765 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v861 = int32(-1)
	goto L200
L205:
	;
	goto L206
L206:
	;
	v772 = int32(1)
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765+v757))))
	if base.Ui32(v774) < base.Ui32(int32(192)) {
		v831 = v774
		v832 = v772
		goto L207
	} else {
		goto L208
	}
L207:
	;
	if int32(232) < v831 {
		goto L201
	} else {
		goto L220
	}
L208:
	;
	v778 = v765 + int32(1)
	if v778 == v756 {
		v831 = v774
		v832 = v772
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778+v757))))
	v783 = v781 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v774) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v787+v757))))
	v799 = v797 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v774) {
		goto L216
	} else {
		goto L217
	}
L211:
	;
	v787 = v765 + int32(2)
	if v787 != v756 {
		goto L210
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v831 = v774<<(uint(int32(6))%32)&int32(1984) | v783
	v832 = int32(2)
	goto L207
L214:
	;
	goto L213
L215:
	;
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757+v803))))
	v831 = v816&int32(63) | (v774<<(uint(int32(18))%32)&int32(_a_F_dutch_porter_UTF_8_stem_7) | v783<<(uint(int32(12))%32) | v799<<(uint(int32(6))%32))
	v832 = int32(4)
	goto L207
L216:
	;
	v803 = v765 + int32(3)
	if v803 != v756 {
		goto L215
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v831 = v774<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_8) | v783<<(uint(int32(6))%32) | v799
	v832 = int32(3)
	goto L207
L219:
	;
	goto L218
L220:
	;
	v836 = v831 - int32(97)
	if v836 < int32(0) {
		goto L201
	} else {
		goto L221
	}
L221:
	;
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v836)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v842)>>(uint(v836&int32(7))%32))&int32(1) == int32(0) {
		goto L201
	} else {
		goto L222
	}
L222:
	;
	v850 = v832 + v765
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v850
	v765 = v850
	goto L202
L224:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v865 = v864 + v861
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v865
	if v618 < v865 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v868 = v865
	goto L227
L226:
	;
	v868 = v618
	goto L227
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v868
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v891 = v881
	goto L230
L228:
	;
	if v986 < int32(0) {
		v1116 = v865
		goto L152
	} else {
		goto L253
	}
L229:
	;
	v986 = v958
	goto L228
L230:
	;
	if v882 <= v891 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v986 = int32(-1)
	goto L228
L233:
	;
	goto L234
L234:
	;
	v898 = int32(1)
	v900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891+v883))))
	if base.Ui32(v900) < base.Ui32(int32(192)) {
		v957 = v900
		v958 = v898
		goto L235
	} else {
		goto L236
	}
L235:
	;
	if int32(232) < v957 {
		goto L248
	} else {
		goto L249
	}
L236:
	;
	v904 = v891 + int32(1)
	if v904 == v882 {
		v957 = v900
		v958 = v898
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904+v883))))
	v909 = v907 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v900) {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v913+v883))))
	v925 = v923 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v900) {
		goto L244
	} else {
		goto L245
	}
L239:
	;
	v913 = v891 + int32(2)
	if v913 != v882 {
		goto L238
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v957 = v900<<(uint(int32(6))%32)&int32(1984) | v909
	v958 = int32(2)
	goto L235
L242:
	;
	goto L241
L243:
	;
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v883+v929))))
	v957 = v942&int32(63) | (v900<<(uint(int32(18))%32)&int32(_a_F_dutch_porter_UTF_8_stem_7) | v909<<(uint(int32(12))%32) | v925<<(uint(int32(6))%32))
	v958 = int32(4)
	goto L235
L244:
	;
	v929 = v891 + int32(3)
	if v929 != v882 {
		goto L243
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v957 = v900<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_8) | v909<<(uint(int32(6))%32) | v925
	v958 = int32(3)
	goto L235
L247:
	;
	goto L246
L248:
	;
	v975 = v958 + v891
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v975
	v891 = v975
	goto L230
L249:
	;
	v962 = v957 - int32(97)
	if v962 < int32(0) {
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v962)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v968)>>(uint(v962&int32(7))%32))&int32(1) != 0 {
		goto L229
	} else {
		goto L251
	}
L251:
	;
	goto L248
L253:
	;
	v989 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v990 = v989 + v986
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v990
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1013 = v990
	goto L256
L254:
	;
	if v1109 < int32(0) {
		v1116 = v865
		goto L152
	} else {
		goto L278
	}
L255:
	;
	v1109 = v1080
	goto L254
L256:
	;
	if v1004 <= v1013 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1109 = int32(-1)
	goto L254
L259:
	;
	goto L260
L260:
	;
	v1020 = int32(1)
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1013+v1005))))
	if base.Ui32(v1022) < base.Ui32(int32(192)) {
		v1079 = v1022
		v1080 = v1020
		goto L261
	} else {
		goto L262
	}
L261:
	;
	if int32(232) < v1079 {
		goto L255
	} else {
		goto L274
	}
L262:
	;
	v1026 = v1013 + int32(1)
	if v1026 == v1004 {
		v1079 = v1022
		v1080 = v1020
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1026+v1005))))
	v1031 = v1029 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v1022) {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	v1045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035+v1005))))
	v1047 = v1045 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v1022) {
		goto L270
	} else {
		goto L271
	}
L265:
	;
	v1035 = v1013 + int32(2)
	if v1035 != v1004 {
		goto L264
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v1079 = v1022<<(uint(int32(6))%32)&int32(1984) | v1031
	v1080 = int32(2)
	goto L261
L268:
	;
	goto L267
L269:
	;
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005+v1051))))
	v1079 = v1064&int32(63) | (v1022<<(uint(int32(18))%32)&int32(_a_F_dutch_porter_UTF_8_stem_7) | v1031<<(uint(int32(12))%32) | v1047<<(uint(int32(6))%32))
	v1080 = int32(4)
	goto L261
L270:
	;
	v1051 = v1013 + int32(3)
	if v1051 != v1004 {
		goto L269
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v1079 = v1022<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_8) | v1031<<(uint(int32(6))%32) | v1047
	v1080 = int32(3)
	goto L261
L273:
	;
	goto L272
L274:
	;
	v1084 = v1079 - int32(97)
	if v1084 < int32(0) {
		goto L255
	} else {
		goto L275
	}
L275:
	;
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1084)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v1090)>>(uint(v1084&int32(7))%32))&int32(1) == int32(0) {
		goto L255
	} else {
		goto L276
	}
L276:
	;
	v1098 = v1080 + v1013
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1098
	v1013 = v1098
	goto L256
L278:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v1112 + v1109
	v1116 = v865
	goto L152
L279:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1303
	v1305 = F_r_e_ending_2(m, l0)
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L12
	} else {
		goto L317
	}
L280:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1124 = int32(1)
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1122+v1118-v1124))))
	if base.B2i32(v1126&int32(224) != int32(96))|base.B2i32(v1124<<(uint(v1126)%32)&int32(_a_F_dutch_porter_UTF_8_stem_11) == int32(0)) != 0 {
		v1302 = v1116
		goto L279
	} else {
		goto L281
	}
L281:
	;
	v1141 = F_find_among_b(m, l0, int32(_a_F_dutch_porter_UTF_8_stem_12), int32(5), int32(0))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L12
	} else {
		goto L282
	}
L282:
	;
	if v1141 == int32(0) {
		v1302 = v1116
		goto L279
	} else {
		goto L283
	}
L283:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1145
	switch v1141 - int32(1) {
	case 0:
		goto L286
	case 1:
		goto L285
	case 2:
		goto L284
	default:
		v1302 = v1116
		goto L279
	}
L284:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1145 < v1163 {
		goto L294
	} else {
		goto L295
	}
L285:
	;
	v1159 = F_r_en_ending_2(m, l0)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L12
	} else {
		goto L292
	}
L286:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1145 < v1149 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1302 = int32(0)
	goto L279
L288:
	;
	goto L289
L289:
	;
	v1154 = F_slice_from_s(m, l0, int32(4), int32(_a_F_dutch_porter_UTF_8_stem_13))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L12
	} else {
		goto L290
	}
L290:
	;
	if v1154 < int32(0) {
		v1976 = v1154
		goto L1
	} else {
		goto L291
	}
L291:
	;
	v1302 = int32(1)
	goto L279
L292:
	;
	if int32(0) <= v1159 {
		v1302 = v1159
		goto L279
	} else {
		goto L293
	}
L293:
	;
	v1976 = v1159
	goto L1
L294:
	;
	v1302 = int32(0)
	goto L279
L295:
	;
	goto L296
L296:
	;
	v1166 = int32(1)
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L299
L297:
	;
	if v1296 != 0 {
		v1302 = v1166
		goto L279
	} else {
		goto L315
	}
L298:
	;
	v1296 = v1289
	goto L297
L299:
	;
	if v1179 <= v1180 {
		v1289 = int32(-1)
		goto L298
	} else {
		goto L301
	}
L300:
	;
	v1289 = int32(0)
	goto L298
L301:
	;
	v1197 = int32(1)
	v1198 = v1179 - v1197
	v1200 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1181+v1198))))
	v1202 = v1200 & int32(255)
	if base.B2i32(v1198 == v1180)|base.B2i32(int32(0) <= v1200) != 0 {
		v1260 = v1202
		v1264 = v1197
		goto L302
	} else {
		goto L303
	}
L302:
	;
	if int32(232) < v1260 {
		goto L310
	} else {
		goto L311
	}
L303:
	;
	v1209 = v1202 & int32(63)
	v1211 = v1179 - int32(2)
	v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1181+v1211))))
	v1215 = v1213 << (uint(int32(6)) % 32)
	if base.B2i32(v1211 != v1180)&base.B2i32(base.Ui32(v1213) < base.Ui32(int32(192))) == int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1260 = v1215&int32(1984) | v1209
	v1264 = int32(2)
	goto L302
L305:
	;
	goto L306
L306:
	;
	v1228 = v1215&int32(4032) | v1209
	v1230 = v1179 - int32(3)
	v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1181+v1230))))
	if base.B2i32(v1230 != v1180)&base.B2i32(base.Ui32(v1232) < base.Ui32(int32(224))) == int32(0) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1260 = v1232<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_8) | v1228
	v1264 = int32(3)
	goto L302
L308:
	;
	goto L309
L309:
	;
	v1250 = int32(4)
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179+v1181-v1250))))
	v1260 = v1232<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_14) | v1252&int32(7)<<(uint(int32(18))%32) | v1228
	v1264 = v1250
	goto L302
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1179 - v1264
	goto L314
L311:
	;
	v1266 = v1260 - int32(97)
	if v1266 < int32(0) {
		goto L310
	} else {
		goto L312
	}
L312:
	;
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1266)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_UTF_8_stem[1]))))
	if int32(base.Ui32(v1272)>>(uint(v1266&int32(7))%32))&int32(1) == int32(0) {
		goto L310
	} else {
		goto L313
	}
L313:
	;
	v1296 = v1264
	goto L297
L314:
	;
	goto L300
L315:
	;
	v1297 = F_slice_del(m, l0)
	mBase = m.M
	if v1297 < int32(0) {
		v1976 = v1297
		goto L1
	} else {
		goto L316
	}
L316:
	;
	v1302 = v1166
	goto L279
L317:
	;
	if v1305 < int32(0) {
		v1976 = v1305
		goto L1
	} else {
		goto L318
	}
L318:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1309
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1309
	v1312 = int32(4)
	v1314 = int32(0)
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1309-v1317 < v1312 {
		v1327 = v1314
		goto L326
	} else {
		goto L327
	}
L319:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1867
	v1870 = v1867
	goto L457
L320:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1508
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L396
L321:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1405 < v1497 {
		goto L320
	} else {
		goto L391
	}
L322:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1405 < v1492 {
		goto L320
	} else {
		goto L389
	}
L323:
	;
	if int32(0) <= v1487 {
		goto L319
	} else {
		goto L388
	}
L324:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1377
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1377
	v1381 = v1377 - int32(1)
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1381 <= v1382 {
		goto L320
	} else {
		goto L349
	}
L325:
	;
	if v1327 == int32(0) {
		v1376 = v1302
		goto L324
	} else {
		goto L329
	}
L326:
	;
	goto L325
L327:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1323 = F_memcmp(m, v1320+v1309-v1312, int32(_a_F_dutch_porter_UTF_8_stem_15), v1312)
	mBase = m.M
	if v1323 != 0 {
		v1327 = v1314
		goto L326
	} else {
		goto L328
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1309 - v1312
	v1327 = int32(1)
	goto L326
L329:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1330
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1330 < v1332 {
		v1376 = v1302
		goto L324
	} else {
		goto L330
	}
L330:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1334 < v1330 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1336+v1330-int32(1)))))
	if v1340 == int32(99) {
		v1376 = v1302
		goto L324
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	v1343 = F_slice_del(m, l0)
	mBase = m.M
	if v1343 < int32(0) {
		v1976 = v1343
		goto L1
	} else {
		goto L335
	}
L334:
	;
	goto L333
L335:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1346
	v1348 = int32(2)
	v1350 = int32(0)
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1346-v1353 < v1348 {
		v1363 = v1350
		goto L337
	} else {
		goto L338
	}
L336:
	;
	if v1363 == int32(0) {
		v1376 = v1302
		goto L324
	} else {
		goto L340
	}
L337:
	;
	goto L336
L338:
	;
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1359 = F_memcmp(m, v1356+v1346-v1348, int32(_a_F_dutch_porter_UTF_8_stem_16), v1348)
	mBase = m.M
	if v1359 != 0 {
		v1363 = v1350
		goto L337
	} else {
		goto L339
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1346 - v1348
	v1363 = int32(1)
	goto L337
L340:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1366
	v1368 = F_r_en_ending_2(m, l0)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L12
	} else {
		goto L341
	}
L341:
	;
	v1371 = base.B2i32(v1368 < int32(0))
	if v1368 < int32(0) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1372 = v1368
	goto L344
L343:
	;
	v1372 = v1302
	goto L344
L344:
	;
	if v1368 != 0 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1373 = v1372
	goto L347
L346:
	;
	v1373 = v1302
	goto L347
L347:
	;
	if v1368 < int32(0) {
		v1487 = v1373
		goto L323
	} else {
		goto L348
	}
L348:
	;
	v1376 = v1373
	goto L324
L349:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1384+v1381))))
	if base.B2i32(v1386&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1386)%32)&int32(_a_F_dutch_porter_UTF_8_stem_17) == int32(0)) != 0 {
		goto L320
	} else {
		goto L350
	}
L350:
	;
	v1401 = F_find_among_b(m, l0, int32(_a_F_dutch_porter_UTF_8_stem_18), int32(6), int32(0))
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L12
	} else {
		goto L351
	}
L351:
	;
	if v1401 == int32(0) {
		goto L320
	} else {
		goto L352
	}
L352:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1405
	switch v1401 - int32(1) {
	case 0:
		goto L355
	case 1:
		goto L354
	case 2:
		goto L353
	case 3:
		goto L322
	case 4:
		goto L321
	default:
		goto L320
	}
L353:
	;
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1405 < v1474 {
		goto L320
	} else {
		goto L378
	}
L354:
	;
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1405 < v1460 {
		goto L320
	} else {
		goto L372
	}
L355:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1405 < v1409 {
		goto L320
	} else {
		goto L356
	}
L356:
	;
	v1411 = F_slice_del(m, l0)
	mBase = m.M
	if v1411 < int32(0) {
		v1976 = v1411
		goto L1
	} else {
		goto L357
	}
L357:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1414
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1417 = int32(2)
	v1419 = int32(0)
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1414-v1422 < v1417 {
		v1432 = v1419
		goto L360
	} else {
		goto L361
	}
L358:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1452 + (v1414 - v1416)
	v1456 = F_r_undouble_3(m, l0)
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L12
	} else {
		goto L370
	}
L359:
	;
	if v1432 == int32(0) {
		goto L358
	} else {
		goto L363
	}
L360:
	;
	goto L359
L361:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1428 = F_memcmp(m, v1425+v1414-v1417, int32(_a_F_dutch_porter_UTF_8_stem_19), v1417)
	mBase = m.M
	if v1428 != 0 {
		v1432 = v1419
		goto L360
	} else {
		goto L362
	}
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1414 - v1417
	v1432 = int32(1)
	goto L360
L363:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1435
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1435 < v1437 {
		goto L358
	} else {
		goto L364
	}
L364:
	;
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1439 < v1435 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1441+v1435-int32(1)))))
	if v1445 == int32(101) {
		goto L358
	} else {
		goto L368
	}
L366:
	;
	goto L367
L367:
	;
	v1448 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1448 {
		goto L320
	} else {
		goto L369
	}
L368:
	;
	goto L367
L369:
	;
	v1976 = v1448
	goto L1
L370:
	;
	if int32(0) <= v1456 {
		goto L320
	} else {
		goto L371
	}
L371:
	;
	v1976 = v1456
	goto L1
L372:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1462 < v1405 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1464+v1405-int32(1)))))
	if v1468 == int32(101) {
		goto L320
	} else {
		goto L376
	}
L374:
	;
	goto L375
L375:
	;
	v1471 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1471 {
		goto L320
	} else {
		goto L377
	}
L376:
	;
	goto L375
L377:
	;
	v1976 = v1471
	goto L1
L378:
	;
	v1476 = F_slice_del(m, l0)
	mBase = m.M
	if v1476 < int32(0) {
		v1976 = v1476
		goto L1
	} else {
		goto L379
	}
L379:
	;
	v1479 = F_r_e_ending_2(m, l0)
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L12
	} else {
		goto L380
	}
L380:
	;
	if int32(0) <= v1479 {
		goto L320
	} else {
		goto L381
	}
L381:
	;
	if v1479 < int32(0) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v1485 = v1479
	goto L384
L383:
	;
	v1485 = v1376
	goto L384
L384:
	;
	if v1479 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1486 = v1485
	goto L387
L386:
	;
	v1486 = v1376
	goto L387
L387:
	;
	v1487 = v1486
	goto L323
L388:
	;
	v1976 = v1487
	goto L1
L389:
	;
	v1494 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1494 {
		goto L320
	} else {
		goto L390
	}
L390:
	;
	v1976 = v1494
	goto L1
L391:
	;
	v1499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v1499 == int32(0) {
		goto L320
	} else {
		goto L392
	}
L392:
	;
	v1502 = F_slice_del(m, l0)
	mBase = m.M
	if v1502 < int32(0) {
		v1976 = v1502
		goto L1
	} else {
		goto L393
	}
L393:
	;
	goto L320
L394:
	;
	if v1639 != 0 {
		goto L319
	} else {
		goto L412
	}
L395:
	;
	v1639 = v1632
	goto L394
L396:
	;
	if v1508 <= v1523 {
		v1632 = int32(-1)
		goto L395
	} else {
		goto L398
	}
L397:
	;
	v1632 = int32(0)
	goto L395
L398:
	;
	v1540 = int32(1)
	v1541 = v1508 - v1540
	v1543 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1524+v1541))))
	v1545 = v1543 & int32(255)
	if base.B2i32(v1541 == v1523)|base.B2i32(int32(0) <= v1543) != 0 {
		v1603 = v1545
		v1607 = v1540
		goto L399
	} else {
		goto L400
	}
L399:
	;
	if int32(232) < v1603 {
		goto L407
	} else {
		goto L408
	}
L400:
	;
	v1552 = v1545 & int32(63)
	v1554 = v1508 - int32(2)
	v1556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1524+v1554))))
	v1558 = v1556 << (uint(int32(6)) % 32)
	if base.B2i32(v1554 != v1523)&base.B2i32(base.Ui32(v1556) < base.Ui32(int32(192))) == int32(0) {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v1603 = v1558&int32(1984) | v1552
	v1607 = int32(2)
	goto L399
L402:
	;
	goto L403
L403:
	;
	v1571 = v1558&int32(4032) | v1552
	v1573 = v1508 - int32(3)
	v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1524+v1573))))
	if base.B2i32(v1573 != v1523)&base.B2i32(base.Ui32(v1575) < base.Ui32(int32(224))) == int32(0) {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1603 = v1575<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_8) | v1571
	v1607 = int32(3)
	goto L399
L405:
	;
	goto L406
L406:
	;
	v1593 = int32(4)
	v1595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1508+v1524-v1593))))
	v1603 = v1575<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_14) | v1595&int32(7)<<(uint(int32(18))%32) | v1571
	v1607 = v1593
	goto L399
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1508 - v1607
	goto L411
L408:
	;
	v1609 = v1603 - int32(73)
	if v1609 < int32(0) {
		goto L407
	} else {
		goto L409
	}
L409:
	;
	v1615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1609)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_UTF_8_stem[2]))))
	if int32(base.Ui32(v1615)>>(uint(v1609&int32(7))%32))&int32(1) == int32(0) {
		goto L407
	} else {
		goto L410
	}
L410:
	;
	v1639 = v1607
	goto L394
L411:
	;
	goto L397
L412:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1642 = v1640 - int32(1)
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1642 <= v1643 {
		goto L319
	} else {
		goto L413
	}
L413:
	;
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1645+v1642))))
	if base.B2i32(v1647&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1647)%32)&int32(_a_F_dutch_porter_UTF_8_stem_20) == int32(0)) != 0 {
		goto L319
	} else {
		goto L414
	}
L414:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1663 = F_find_among_b(m, l0, int32(_a_F_dutch_porter_UTF_8_stem_21), int32(4), int32(0))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L12
	} else {
		goto L415
	}
L415:
	;
	if v1663 == int32(0) {
		goto L319
	} else {
		goto L416
	}
L416:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L419
L417:
	;
	if v1796 != 0 {
		goto L319
	} else {
		goto L435
	}
L418:
	;
	v1796 = v1789
	goto L417
L419:
	;
	if v1679 <= v1680 {
		v1789 = int32(-1)
		goto L418
	} else {
		goto L421
	}
L420:
	;
	v1789 = int32(0)
	goto L418
L421:
	;
	v1697 = int32(1)
	v1698 = v1679 - v1697
	v1700 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1681+v1698))))
	v1702 = v1700 & int32(255)
	if base.B2i32(v1698 == v1680)|base.B2i32(int32(0) <= v1700) != 0 {
		v1760 = v1702
		v1764 = v1697
		goto L422
	} else {
		goto L423
	}
L422:
	;
	if int32(232) < v1760 {
		goto L430
	} else {
		goto L431
	}
L423:
	;
	v1709 = v1702 & int32(63)
	v1711 = v1679 - int32(2)
	v1713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1681+v1711))))
	v1715 = v1713 << (uint(int32(6)) % 32)
	if base.B2i32(v1711 != v1680)&base.B2i32(base.Ui32(v1713) < base.Ui32(int32(192))) == int32(0) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1760 = v1715&int32(1984) | v1709
	v1764 = int32(2)
	goto L422
L425:
	;
	goto L426
L426:
	;
	v1728 = v1715&int32(4032) | v1709
	v1730 = v1679 - int32(3)
	v1732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1681+v1730))))
	if base.B2i32(v1730 != v1680)&base.B2i32(base.Ui32(v1732) < base.Ui32(int32(224))) == int32(0) {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1760 = v1732<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_8) | v1728
	v1764 = int32(3)
	goto L422
L428:
	;
	goto L429
L429:
	;
	v1750 = int32(4)
	v1752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1679+v1681-v1750))))
	v1760 = v1732<<(uint(int32(12))%32)&int32(_a_F_dutch_porter_UTF_8_stem_14) | v1752&int32(7)<<(uint(int32(18))%32) | v1728
	v1764 = v1750
	goto L422
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1679 - v1764
	goto L434
L431:
	;
	v1766 = v1760 - int32(97)
	if v1766 < int32(0) {
		goto L430
	} else {
		goto L432
	}
L432:
	;
	v1772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1766)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v1772)>>(uint(v1766&int32(7))%32))&int32(1) == int32(0) {
		goto L430
	} else {
		goto L433
	}
L433:
	;
	v1796 = v1764
	goto L417
L434:
	;
	goto L420
L435:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1799 = v1797 + (v1640 - v1659)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1799
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1799
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L438
L436:
	;
	if v1855 < int32(0) {
		goto L319
	} else {
		goto L455
	}
L438:
	;
	goto L439
L439:
	;
	goto L440
L440:
	;
	v1810 = v1799
	v1812 = int32(1)
	goto L443
L442:
	;
	v1855 = v1837
	goto L436
L443:
	;
	if v1810 <= v1803 {
		goto L445
	} else {
		goto L446
	}
L444:
	;
	goto L442
L445:
	;
	v1855 = int32(-1)
	goto L436
L446:
	;
	goto L447
L447:
	;
	v1817 = v1810 - int32(1)
	v1819 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1802+v1817))))
	if base.B2i32(int32(0) <= v1819)|base.B2i32(v1817 <= v1803) != 0 {
		v1837 = v1817
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v1841 = int32(1)
	if v1841 < v1812 {
		v1810 = v1837
		v1812 = v1812 - v1841
		goto L443
	} else {
		goto L454
	}
L449:
	;
	v1825 = v1817
	goto L450
L450:
	;
	v1830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1802+v1825))))
	if base.Ui32(int32(191)) < base.Ui32(v1830) {
		v1837 = v1825
		goto L448
	} else {
		goto L452
	}
L451:
	;
	v1837 = v1803
	goto L448
L452:
	;
	v1834 = v1825 - int32(1)
	if v1803 < v1834 {
		v1825 = v1834
		goto L450
	} else {
		goto L453
	}
L453:
	;
	goto L451
L454:
	;
	goto L444
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1855
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1855
	v1860 = F_slice_del(m, l0)
	mBase = m.M
	if v1860 < int32(0) {
		v1976 = v1860
		goto L1
	} else {
		goto L456
	}
L456:
	;
	goto L319
L457:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1870
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1875 <= v1870 {
		goto L462
	} else {
		goto L463
	}
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1867
	v1976 = int32(1)
	goto L1
L459:
	;
	goto L458
L460:
	;
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1870 = v1972
	goto L457
L461:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L478
L462:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1870
	v1910 = v1870
	v1911 = v1875
	goto L461
L463:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1877+v1870))))
	v1881 = v1879 - int32(73)
	if v1881 != int32(16) {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v1885 = v1881
	goto L466
L465:
	;
	v1885 = int32(0)
	goto L466
L466:
	;
	if v1885 != 0 {
		goto L462
	} else {
		goto L467
	}
L467:
	;
	v1889 = F_find_among(m, l0, int32(_a_F_dutch_porter_UTF_8_stem_22), int32(3), int32(0))
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L12
	} else {
		goto L468
	}
L468:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1891
	switch v1889 - int32(1) {
	case 0:
		goto L470
	case 1:
		goto L469
	case 2:
		goto L471
	default:
		goto L460
	}
L469:
	;
	v1904 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_UTF_8_stem_23))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L12
	} else {
		goto L474
	}
L470:
	;
	v1898 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_porter_UTF_8_stem_24))
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L12
	} else {
		goto L472
	}
L471:
	;
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1910 = v1891
	v1911 = v1895
	goto L461
L472:
	;
	if int32(0) <= v1898 {
		goto L460
	} else {
		goto L473
	}
L473:
	;
	v1976 = v1898
	goto L1
L474:
	;
	if int32(0) <= v1904 {
		goto L460
	} else {
		goto L475
	}
L475:
	;
	v1976 = v1904
	goto L1
L476:
	;
	if v1965 < int32(0) {
		goto L459
	} else {
		goto L496
	}
L478:
	;
	goto L479
L479:
	;
	goto L480
L480:
	;
	v1920 = v1910
	v1922 = int32(1)
	goto L483
L482:
	;
	v1965 = v1950
	goto L476
L483:
	;
	if v1911 <= v1920 {
		goto L485
	} else {
		goto L486
	}
L484:
	;
	goto L482
L485:
	;
	v1965 = int32(-1)
	goto L476
L486:
	;
	goto L487
L487:
	;
	v1927 = v1920 + int32(1)
	v1929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1913+v1920))))
	if base.Ui32(v1929) < base.Ui32(int32(192)) {
		v1950 = v1927
		goto L488
	} else {
		goto L489
	}
L488:
	;
	v1951 = int32(1)
	if v1951 < v1922 {
		v1920 = v1950
		v1922 = v1922 - v1951
		goto L483
	} else {
		goto L495
	}
L489:
	;
	if v1911 <= v1927 {
		v1950 = v1927
		goto L488
	} else {
		goto L490
	}
L490:
	;
	v1936 = v1927
	goto L491
L491:
	;
	v1939 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1913+v1936))))
	if int32(-65) < v1939 {
		v1950 = v1936
		goto L488
	} else {
		goto L493
	}
L492:
	;
	v1950 = v1911
	goto L488
L493:
	;
	v1943 = v1936 + int32(1)
	if v1943 != v1911 {
		v1936 = v1943
		goto L491
	} else {
		goto L494
	}
L494:
	;
	goto L492
L495:
	;
	goto L484
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1965
	goto L460
}
