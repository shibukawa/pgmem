package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReplicationSlotDropAtPubNode(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
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
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v21 = v4
	v22 = v4
	v23 = v4
	v24 = v4
	v25 = v4
	v26 = int32(-1)
	v27 = v14
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v26 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v262 = int32(m.ExcTag)
	v263 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v262 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L7:
	;
	v31 = v27 - int32(16)
	m.G0 = v31
	v34 = v31 - int32(160)
	m.G0 = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v31
	F_load_file(m, int32(212219), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		v261 = v34
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v76 = v21
	v77 = v22
	v78 = v23
	v79 = v24
	v80 = v25
	v81 = v27
	goto L9
L9:
	;
	if v80 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v31
	F_initStringInfo(m, v31)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		v261 = v34
		goto L6
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v31
	v54 = F_quote_identifier(m, l1)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		v261 = v34
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v54
	F_appendStringInfo(m, v31, int32(513013), v14+int32(48))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		v261 = v34
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	v70 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v14 + int32(60)
	goto L17
L15:
	;
	v76 = v31
	v77 = v34
	v78 = v70
	v79 = v68
	v80 = int32(0)
	v81 = v34
	goto L9
L17:
	;
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v77
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v88 = *(*int32)(unsafe.Add(mBase, _consts[484]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	v94 = int32(0)
	v96 = m.T0[v89].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v86, v94, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L26
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v79
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v78
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	F_pfree(m, v242)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L54
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L50
	}
L22:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	if v159 != 0 {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	F_errfinish(m, int32(486731), v149, int32(407978))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L35
	}
L24:
	;
	if l2 == int32(0) {
		goto L21
	} else {
		goto L30
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	v105 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L27
	}
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	switch v98 {
	case 0:
		goto L24
	case 1:
		goto L25
	default:
		goto L21
	}
L27:
	;
	if v105 == int32(0) {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l1
	F_errmsg(m, int32(220187), v14+int32(16))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v149 = int32(1937)
	goto L23
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v122 != int32(67137668) {
		goto L21
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	v131 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L32
	}
L32:
	;
	if v131 == int32(0) {
		goto L22
	} else {
		goto L33
	}
L33:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l1
	F_errmsg(m, int32(198622), v14+int32(32))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v149 = int32(1946)
	goto L23
L35:
	;
	goto L22
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	F_pfree(m, v159)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	if v166 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	F_tuplestore_end(m, v166)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	if v173 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	F_FreeTupleDesc(m, v173)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	F_pfree(m, v96)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v79
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v78
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	F_pfree(m, v190)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v79
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v78
	m.G0 = v14 + int32(80)
	return
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	F_errcode(m, int32(100663808))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L51
	}
L51:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
	F_errmsg(m, int32(198622), v14)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	F_errfinish(m, int32(486731), int32(1954), int32(407978))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L53
	}
L53:
	;
	goto L3
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v76
	F_pg_re_throw(m)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		v261 = v81
		goto L6
	} else {
		goto L55
	}
L55:
	;
	goto L5
L56:
	;
	v267 = int32(v263)
	m.G0 = v261
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v14+int32(60) == v274 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	m.ExcPending = 1
	goto L65
L58:
	;
	if v277 != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	v277 = v276
	goto L61
L60:
	;
	v277 = int32(0)
	goto L61
L61:
	;
	goto L58
L62:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v21 = v278
	v22 = v279
	v23 = v280
	v24 = v281
	v25 = v269
	v26 = v277
	v27 = v261
	goto L1
L63:
	;
	goto L64
L64:
	;
	F___wasm_longjmp(m, v270, v269)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	return
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReplicationSlotRelease(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[698])))
	if v15 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = F_pstrdup(m, v12+int32(24))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v26 = int32(0)
	v27 = int32(674157)
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	if v28 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = int32(674237)
	goto L8
L7:
	;
	v25 = int32(674157)
	goto L8
L8:
	;
	v26 = v20
	v27 = v25
	goto L3
L9:
	;
	v31 = int32(4380996)
	v32 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	*(*int32)(unsafe.Add(mBase, _consts[664])) = int32(0)
	F_ReplicationSlotDropPtr(m, v32)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v62 = m.G0
	v63 = int32(16)
	v64 = v62 - v63
	m.G0 = v64
	F___gettimeofday(m, v64)
	mBase = m.M
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v68 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+8)))
	m.G0 = v64 + v63
	v76 = v68 + v67*int64(1000000) - int64(946684800000000)
	goto L21
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v40 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(1)
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_s_lock(m, v12, int32(485480), int32(750), int32(355975))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v51
	F_ReplicationSlotsComputeRequiredXmin(m, v51)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	goto L13
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	if v80 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v112 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[664])) = v112
	v115 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v119 = F_LWLockAcquire(m, v115+int32(512), v112)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L41
	}
L23:
	;
	if v77 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if v77 != 0 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	F_s_lock(m, v12, int32(485480), int32(768), int32(355975))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v88 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v12)+112))
	if v90 == v88 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+272)) = v76
	goto L32
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
	F_ConditionVariableBroadcast(m, v12+int32(224))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L22
L34:
	;
	F_s_lock(m, v12, int32(321643), int32(251), int32(409705))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v12)+112))
	if v105 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+272)) = v76
	goto L40
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
	goto L22
L41:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+124)))
	v125 = v123 & int32(-17)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+124)) = uint8(v125)
	v128 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v122)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v129+v130))) = uint8(v125)
	v134 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v134+int32(512))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, _consts[698])))
	if v140 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, _consts[699])))
	if v146 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	m.G0 = v9 + int32(16)
	return
L46:
	;
	v147 = int32(15)
	goto L48
L47:
	;
	v147 = int32(14)
	goto L48
L48:
	;
	v149 = F_errstart(m, v147, int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v149 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v26
	F_errmsg(m, v27, v9)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_pfree(m, v26)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	F_errfinish(m, int32(485480), int32(792), int32(355975))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L45
}
func F_copy_replication_slot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int64
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int64
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int64
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int64
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(688)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_get_call_result_type(m, l0, v3, v15+int32(96))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L7
	} else {
		goto L148
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L7
	} else {
		goto L143
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L7
	} else {
		goto L139
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L7
	} else {
		goto L135
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L7
	} else {
		goto L131
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L7
	} else {
		goto L124
	}
L7:
	;
	return int32(0)
L8:
	;
	if v22 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_CheckSlotPermissions(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L7
	} else {
		goto L121
	}
L12:
	;
	if l1 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v39 = F_LWLockAcquire(m, v35+int32(4736), int32(1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L19
	}
L14:
	;
	F_CheckLogicalDecodingRequirements(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_CheckSlotRequirements(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	goto L13
L19:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	if int32(0) < v42 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if l1 != 0 {
		goto L62
	} else {
		goto L63
	}
L21:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v172 = base.B2i32(v170 != int32(0))
	if v125 == int32(3) {
		v176 = v124
		v177 = v172
		goto L20
	} else {
		goto L60
	}
L22:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	v50 = v3
	goto L25
L23:
	;
	goto L24
L24:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v147+int32(4736))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L55
	}
L25:
	;
	v61 = v46 + v50*int32(288)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+4)))
	if v62 != int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	v132 = v50 + int32(1)
	if v132 != v42 {
		v50 = v132
		goto L25
	} else {
		goto L54
	}
L28:
	;
	v66 = v61 + int32(24)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v70 == int32(0) {
		v89 = v69
		v90 = v70
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v90-v89 != 0 {
		goto L27
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	if v69 != v70 {
		v89 = v69
		v90 = v70
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v74 = v66
	v75 = v18
	goto L33
L33:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	if v79 == int32(0) {
		v89 = v78
		v90 = v79
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v89 = v78
	v90 = v79
	goto L30
L35:
	;
	v82 = int32(1)
	if v78 == v79 {
		v74 = v74 + v82
		v75 = v75 + v82
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(1)
	if v92 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_s_lock(m, v61, int32(487009), int32(651), int32(83902))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	goto L43
L41:
	;
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v108+int32(4736))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L7
	} else {
		goto L46
	}
L43:
	;
	v103 = F__emscripten_memcpy_bulkmem(m, v15+int32(400), v61, int32(288))
	mBase = m.M
	goto L45
L45:
	;
	goto L42
L46:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v15)+488))
	if l1 != base.B2i32(v113 != int32(0)) {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v15)+504))
	if v117 == int64(0) {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v15)+512))
	if v120 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if l1 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v124 = v15 + int32(537)
	goto L52
L51:
	;
	v124 = int32(0)
	goto L52
L52:
	;
	v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(3) <= v125 {
		goto L21
	} else {
		goto L53
	}
L53:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v15)+492))
	v176 = v124
	v177 = base.B2i32(v128 == int32(2))
	goto L20
L54:
	;
	goto L26
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v18
	F_errmsg(m, int32(69824), v15+int32(80))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(487009), int32(664), int32(83902))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v176 = v175
	v177 = v172
	goto L20
L61:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(1)
	if v219 != 0 {
		goto L77
	} else {
		goto L78
	}
L62:
	;
	v178 = int32(1)
	if v177 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	v203 = int32(0)
	if v177 != 0 {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v181 = int32(2)
	goto L67
L66:
	;
	v181 = v178
	goto L67
L67:
	;
	v182 = int32(0)
	F_ReplicationSlotCreate(m, v17, v178, v181, v182, v182, v182)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = int32(396)
	v193 = int32(0)
	v199 = F_CreateInitDecodingContext(m, v176, v193, v117, v15+int32(112), v193, v193, v193)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	F_FreeDecodingContext(m, v199)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	goto L61
L71:
	;
	v206 = int32(2)
	goto L73
L72:
	;
	v206 = v203
	goto L73
L73:
	;
	v207 = int32(0)
	F_ReplicationSlotCreate(m, v17, v203, v206, v207, v207, v207)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	*(*int64)(unsafe.Add(mBase, uint32(v213)+104)) = v117
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L7
	} else {
		goto L76
	}
L76:
	;
	goto L61
L77:
	;
	F_s_lock(m, v61, int32(487009), int32(753), int32(83902))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L7
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	goto L82
L80:
	;
	goto L79
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(0)
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v15)+216))
	if base.Ui64(v234) < base.Ui64(v117) {
		goto L3
	} else {
		goto L85
	}
L82:
	;
	v230 = F__emscripten_memcpy_bulkmem(m, v15+int32(112), v61, int32(288))
	mBase = m.M
	goto L84
L84:
	;
	goto L81
L85:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v15)+200))
	v237 = int32(0)
	if base.B2i32(v236 == v237) == base.B2i32(v113 != v237) {
		goto L3
	} else {
		goto L86
	}
L86:
	;
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v15)+232))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v15)+212))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v15)+208))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v15)+132))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
	v248 = v15 + int32(136)
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if v252 == int32(0) {
		v271 = v251
		v272 = v252
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v272-v271 != 0 {
		goto L3
	} else {
		goto L95
	}
L88:
	;
	goto L87
L89:
	;
	if v251 != v252 {
		v271 = v251
		v272 = v252
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v256 = v248
	v257 = v18
	goto L91
L91:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+1)))
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+1)))
	if v261 == int32(0) {
		v271 = v260
		v272 = v261
		goto L88
	} else {
		goto L93
	}
L92:
	;
	v271 = v260
	v272 = v261
	goto L88
L93:
	;
	v264 = int32(1)
	if v260 == v261 {
		v256 = v256 + v264
		v257 = v257 + v264
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	if v242 == int64(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v277 = v113
	goto L98
L97:
	;
	v277 = int32(0)
	goto L98
L98:
	;
	if v277 != 0 {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v15)+224))
	if v278 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = int32(1)
	if v281 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	F_s_lock(m, v285, int32(487009), int32(810), int32(83902))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L7
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	*(*int64)(unsafe.Add(mBase, uint32(v292)+120)) = v242
	*(*int64)(unsafe.Add(mBase, uint32(v292)+104)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v292)+100)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v292)+96)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v292)+20)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v292)+16)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v292))) = int32(0)
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L7
	} else {
		goto L105
	}
L104:
	;
	goto L103
L105:
	;
	F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	if (l1^int32(-1)|v177)&int32(1) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F_ReplicationSlotPersist(m)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L7
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v17
	v320 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+102)) = uint8(v320)
	v324 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v324)+120))
	if v325 == int64(0) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L111
L113:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+103)) = uint8(v332)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	v339 = F_heap_form_tuple(m, v334, v15+int32(104), v15+int32(102))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L7
	} else {
		goto L118
	}
L114:
	;
	v332 = int32(1)
	goto L113
L115:
	;
	goto L116
L116:
	;
	v329 = F_Int64GetDatum(m, v325)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L7
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+108)) = v329
	v332 = v320
	goto L113
L118:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v339)+16))
	v342 = F_HeapTupleHeaderGetDatum(m, v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	m.G0 = v15 + int32(688)
	return v342
L121:
	;
	F_errmsg_internal(m, int32(361809), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(487009), int32(622), int32(83902))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v18
	if v113 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v373 = int32(84593)
	goto L128
L127:
	;
	v373 = int32(84413)
	goto L128
L128:
	;
	F_errmsg(m, v373, v15)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L7
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(487009), int32(679), int32(83902))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L7
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(525680), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(487009), int32(685), int32(83902))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L7
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L7
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v18
	F_errmsg(m, int32(674512), v15-int32(-64))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(487009), int32(692), int32(83902))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L7
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v18
	F_errmsg(m, int32(674008), v15+int32(16))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	F_errdetail(m, int32(592189), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L7
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(487009), int32(785), int32(83902))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L7
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L7
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v18
	F_errmsg(m, int32(674315), v15+int32(32))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L7
	} else {
		goto L145
	}
L145:
	;
	F_errhint(m, int32(622613), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L7
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(487009), int32(793), int32(83902))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L7
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v18
	F_errmsg(m, int32(673974), v15+int32(48))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L7
	} else {
		goto L149
	}
L149:
	;
	F_errdetail(m, int32(592270), int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L7
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(487009), int32(807), int32(83902))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L7
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_replication_yyensure_buffer_stack(m *base.Module, l0 int32) {
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
				F_yy_fatal_error_3(m, int32(660297))
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
					F_yy_fatal_error_3(m, int32(660297))
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
func F_replication_yyerror(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_errcode(m, int32(16801924))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
			F_errmsg_internal(m, int32(203498), v5)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(309790), int32(265), int32(207467))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
