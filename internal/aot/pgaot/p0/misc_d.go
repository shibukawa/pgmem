package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_DeconstructFkConstraintRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v21 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(21))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L8
	} else {
		goto L113
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L8
	} else {
		goto L110
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L8
	} else {
		goto L107
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L104
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L101
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L8
	} else {
		goto L98
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L8
	} else {
		goto L95
	}
L8:
	;
	return
L9:
	;
	v23 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v25 != int32(1) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v28 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v29 != int32(21) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if base.Ui32(v32-int32(33)) <= base.Ui32(int32(-33)) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v40 = v32 << (uint(int32(1)) % 32)
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v23 != v21 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v41 = F__emscripten_memcpy_bulkmem(m, l2, v23+int32(24), v40)
	mBase = m.M
	goto L18
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	F_pfree(m, v23)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v48 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(22))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v50 = F_pg_detoast_datum(m, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v52 != int32(1) {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	if v55 != v32 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	if v57 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v58 != int32(21) {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	if v40 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v48 != v50 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v63 = F__emscripten_memcpy_bulkmem(m, l3, v50+int32(24), v40)
	mBase = m.M
	goto L32
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	F_pfree(m, v50)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if l4 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	if l5 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L38:
	;
	v72 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(23))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v74 = F_pg_detoast_datum(m, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v76 != int32(1) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	if v79 != v32 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	if v81 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	if v82 != int32(26) {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v88 = v32 << (uint(int32(2)) % 32)
	if v88 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v72 == v74 {
		goto L37
	} else {
		goto L49
	}
L46:
	;
	v89 = F__emscripten_memcpy_bulkmem(m, l4, v74+int32(24), v88)
	mBase = m.M
	goto L48
L47:
	;
	goto L48
L48:
	;
	goto L45
L49:
	;
	F_pfree(m, v74)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	goto L37
L51:
	;
	if l6 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L52:
	;
	v100 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(24))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v102 = F_pg_detoast_datum(m, v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if v104 != int32(1) {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v102)+16))
	if v107 != v32 {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	if v109 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	if v110 != int32(26) {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v116 = v32 << (uint(int32(2)) % 32)
	if v116 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v100 == v102 {
		goto L51
	} else {
		goto L63
	}
L60:
	;
	v117 = F__emscripten_memcpy_bulkmem(m, l5, v102+int32(24), v116)
	mBase = m.M
	goto L62
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	F_pfree(m, v102)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	goto L51
L65:
	;
	if l8 != 0 {
		goto L79
	} else {
		goto L80
	}
L66:
	;
	v128 = F_SysCacheGetAttrNotNull(m, int32(19), l0, int32(25))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	v130 = F_pg_detoast_datum(m, v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v132 != int32(1) {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	if v135 != v32 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	if v137 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	if v138 != int32(26) {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v144 = v32 << (uint(int32(2)) % 32)
	if v144 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v128 == v130 {
		goto L65
	} else {
		goto L77
	}
L74:
	;
	v145 = F__emscripten_memcpy_bulkmem(m, l6, v130+int32(24), v144)
	mBase = m.M
	goto L76
L75:
	;
	goto L76
L76:
	;
	goto L73
L77:
	;
	F_pfree(m, v130)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	goto L65
L79:
	;
	v157 = F_SysCacheGetAttr(m, int32(19), l0, int32(26), v17+int32(15))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L8
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v32
	m.G0 = v17 + int32(16)
	return
L82:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
	if v159 != 0 {
		v179 = int32(0)
		goto L83
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v179
	goto L81
L84:
	;
	v160 = F_pg_detoast_datum(m, v157)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v162 != int32(1) {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	if v165 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	if v166 != int32(21) {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v160)+16))
	v173 = v171 << (uint(int32(1)) % 32)
	if v173 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if v157 == v160 {
		v179 = v171
		goto L83
	} else {
		goto L93
	}
L90:
	;
	v174 = F__emscripten_memcpy_bulkmem(m, l8, v160+int32(24), v173)
	mBase = m.M
	goto L92
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	F_pfree(m, v160)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	v179 = v171
	goto L83
L95:
	;
	F_errmsg_internal(m, int32(25214), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(512484), int32(1557), int32(33041))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v32
	F_errmsg_internal(m, int32(156677), v17)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(512484), int32(1560), int32(33041))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
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
	F_errmsg_internal(m, int32(25249), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(512484), int32(1572), int32(33041))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
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
	F_errmsg_internal(m, int32(26304), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(512484), int32(1587), int32(33041))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
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
	F_errmsg_internal(m, int32(26271), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(512484), int32(1602), int32(33041))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errmsg_internal(m, int32(26337), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(512484), int32(1617), int32(33041))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L8
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errmsg_internal(m, int32(25285), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(512484), int32(1639), int32(33041))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L8
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_DetermineTimeZoneAbbrevOffsetTS(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int64
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	v7 = m.G0
	v9 = v7 - int32(288)
	m.G0 = v9
	v12 = base.I64_div_s(l0, int64(1000000))
	goto L1
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+280)) = v12 + int64(946684800)
	v17 = v9 + int32(16)
	goto L5
L2:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
	if v133 != 0 {
		goto L34
	} else {
		goto L35
	}
L3:
	;
	v130 = F_strlen(m, v119)
	mBase = m.M
	goto L2
L5:
	;
	goto L6
L6:
	;
	v24 = int32(255)
	if (v17^l1)&int32(3) != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v123 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v123)
	goto L3
L8:
	;
	v104 = v99
	v105 = v100
	v106 = v101
	goto L30
L9:
	;
	if v94 == int32(0) {
		v119 = v92
		v120 = v93
		goto L7
	} else {
		goto L29
	}
L10:
	;
	v92 = l1
	v93 = v17
	v94 = v24
	goto L9
L11:
	;
	goto L12
L12:
	;
	if l1&int32(3) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v61 == int32(0) {
		v119 = v58
		v120 = v59
		goto L7
	} else {
		goto L22
	}
L14:
	;
	v58 = l1
	v59 = v17
	v60 = v24
	v61 = int32(1)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v37 = l1
	v38 = v17
	v39 = v24
	goto L17
L17:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v41)
	if v41 == int32(0) {
		v99 = v37
		v100 = v38
		v101 = v39
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v58 = v52
	v59 = v46
	v60 = v48
	v61 = v50
	goto L13
L19:
	;
	v45 = int32(1)
	v46 = v38 + v45
	v48 = v39 - v45
	v49 = int32(0)
	v50 = base.B2i32(v48 != v49)
	v52 = v37 + v45
	if v52&int32(3) == v49 {
		v58 = v52
		v59 = v46
		v60 = v48
		v61 = v50
		goto L13
	} else {
		goto L20
	}
L20:
	;
	if v48 != 0 {
		v37 = v52
		v38 = v46
		v39 = v48
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v64 == int32(0) {
		v92 = v58
		v93 = v59
		v94 = v60
		goto L9
	} else {
		goto L23
	}
L23:
	;
	if base.Ui32(v60) < base.Ui32(int32(4)) {
		v92 = v58
		v93 = v59
		v94 = v60
		goto L9
	} else {
		goto L24
	}
L24:
	;
	v70 = v58
	v71 = v59
	v72 = v60
	goto L25
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v78 = int32(-2139062144)
	if (int32(16843008)-v75|v75)&v78 != v78 {
		v99 = v70
		v100 = v71
		v101 = v72
		goto L8
	} else {
		goto L27
	}
L26:
	;
	v92 = v86
	v93 = v84
	v94 = v88
	goto L9
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v75
	v83 = int32(4)
	v84 = v71 + v83
	v86 = v70 + v83
	v88 = v72 - v83
	if base.Ui32(int32(3)) < base.Ui32(v88) {
		v70 = v86
		v71 = v84
		v72 = v88
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v99 = v92
	v100 = v93
	v101 = v94
	goto L8
L30:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v108)
	if v108 == int32(0) {
		v119 = v104
		v120 = v105
		goto L7
	} else {
		goto L32
	}
L31:
	;
	v119 = v115
	v120 = v113
	goto L7
L32:
	;
	v112 = int32(1)
	v113 = v105 + v112
	v115 = v104 + v112
	v117 = v106 - v112
	if v117 != 0 {
		v104 = v115
		v105 = v113
		v106 = v117
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v137 = v9 + int32(16)
	v141 = v133
	goto L37
L35:
	;
	goto L36
L36:
	;
	v169 = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l2)+268))
	if v176 <= v169 {
		v339 = v169
		goto L47
	} else {
		goto L48
	}
L37:
	;
	if base.Ui32((v141-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L36
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v137))) = uint8(v152)
	v155 = v137 + int32(1)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v156 != 0 {
		v137 = v155
		v141 = v156
		goto L37
	} else {
		goto L43
	}
L40:
	;
	v150 = v141 - int32(32)
	goto L42
L41:
	;
	v150 = v141
	goto L42
L42:
	;
	v152 = v150 & int32(255)
	goto L39
L43:
	;
	goto L38
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L87
	} else {
		goto L90
	}
L45:
	;
	m.G0 = v9 + int32(288)
	return v363
L46:
	;
	if v339 != 0 {
		goto L84
	} else {
		goto L85
	}
L47:
	;
	goto L46
L48:
	;
	v180 = l2 + int32(22376)
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v9+int32(280))))
	v189 = v169
	goto L49
L49:
	;
	v195 = F_strcmp(m, v9+int32(16), v180+v189)
	mBase = m.M
	if v195 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l2)+260))
	if v213 <= int32(0) {
		goto L59
	} else {
		goto L60
	}
L51:
	;
	v197 = v189
	goto L54
L52:
	;
	goto L53
L53:
	;
	goto L50
L54:
	;
	v210 = v197 + int32(1)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197+v180))))
	if v211 != 0 {
		v197 = v210
		goto L54
	} else {
		goto L56
	}
L55:
	;
	if v210 < v176 {
		v189 = v210
		goto L49
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v339 = v169
	goto L47
L58:
	;
	v258 = l2 + int32(16280)
	v260 = l2 + int32(18280)
	v267 = v246
	goto L72
L59:
	;
	v246 = int32(0)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v221 = int32(0)
	v226 = v213
	goto L62
L62:
	;
	v233 = int32(1)
	v234 = (v221 + v226) >> (uint(v233) % 32)
	v240 = *(*int64)(unsafe.Add(mBase, uint32(l2+int32(280)+v234<<(uint(int32(3))%32))))
	v241 = base.B2i32(v181 < v240)
	if v181 < v240 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v246 = v242
	goto L58
L64:
	;
	v242 = v221
	goto L66
L65:
	;
	v242 = v234 + v233
	goto L66
L66:
	;
	if v181 < v240 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v243 = v234
	goto L69
L68:
	;
	v243 = v226
	goto L69
L69:
	;
	if v242 < v243 {
		v221 = v242
		v226 = v243
		goto L62
	} else {
		goto L70
	}
L70:
	;
	goto L63
L71:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v318)))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(12)))) = v325
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v327
	v339 = int32(1)
	goto L47
L72:
	;
	if int32(0) < v267 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l2)+uint32(_consts[786])))
	v287 = v260 + v284<<(uint(int32(4))%32)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	if v288 == v189 {
		v318 = v287
		goto L71
	} else {
		goto L78
	}
L74:
	;
	v276 = v267 - int32(1)
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258+v276))))
	v281 = v260 + v278<<(uint(int32(4))%32)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+8))
	if v282 != v189 {
		v267 = v276
		goto L72
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	goto L73
L77:
	;
	v318 = v281
	goto L71
L78:
	;
	if v213 <= v246 {
		v339 = v169
		goto L47
	} else {
		goto L79
	}
L79:
	;
	v292 = v246
	goto L80
L80:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292+v258))))
	v307 = v260 + v304<<(uint(int32(4))%32)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+8))
	if v308 == v189 {
		v318 = v307
		goto L71
	} else {
		goto L82
	}
L81:
	;
	v339 = v169
	goto L47
L82:
	;
	v311 = v292 + int32(1)
	if v213 != v311 {
		v292 = v311
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v363 = int32(0) - v343
	goto L45
L85:
	;
	goto L86
L86:
	;
	v352 = F_timestamp2tm(m, l0, v9+int32(12), v9+int32(16), v9+int32(8), int32(0), l2)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	return int32(0)
L88:
	;
	if v352 != 0 {
		goto L44
	} else {
		goto L89
	}
L89:
	;
	v360 = F_DetermineTimeZoneOffsetInternal(m, v9+int32(16), l2, v9+int32(280))
	mBase = m.M
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v361
	v363 = v360
	goto L45
L90:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L87
	} else {
		goto L91
	}
L91:
	;
	F_errmsg(m, int32(418494), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(519327), int32(1826), int32(543732))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L87
	} else {
		goto L93
	}
L93:
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
	var v114 float64
	_ = v114
	var v119 float64
	_ = v119
	var v122 float64
	_ = v122
	var v126 float64
	_ = v126
	var v127 float64
	_ = v127
	var v139 float64
	_ = v139
	var v146 int32
	_ = v146
	var v147 float64
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v6)&int64(9223372036854775807)) {
		v147 = math.Float64frombits(uint64(0x7ff8000000000000))
		v148 = F_Float8GetDatum(m, v147)
		mBase = m.M
		v149 = m.ExcPending
		if v149 != 0 {
			return int32(0)
		} else {
			return v148
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
			v139 = v33
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
						v139 = base.F64_sub(v70, base.F64_mul(v70, base.F64_add(v87, v104)))
					} else {
						v111 = v71 << (uint(int32(3)) % 32)
						v114 = *(*float64)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[791])))
						v119 = *(*float64)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[792])))
						v122 = base.F64_sub(v114, base.F64_sub(base.F64_sub(base.F64_mul(v70, base.F64_add(v87, v104)), v119), v70))
						if v18 < int64(0) {
							v126 = base.F64_neg(v122)
						} else {
							v126 = v122
						}
						v127 = v126
						v139 = v127
					}
				} else {
					v127 = v6
					v139 = v127
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
					v139 = base.F64_sub(v70, base.F64_mul(v70, base.F64_add(v87, v104)))
				} else {
					v111 = v71 << (uint(int32(3)) % 32)
					v114 = *(*float64)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[791])))
					v119 = *(*float64)(unsafe.Add(mBase, uint32(v111)+uint32(_consts[792])))
					v122 = base.F64_sub(v114, base.F64_sub(base.F64_sub(base.F64_mul(v70, base.F64_add(v87, v104)), v119), v70))
					if v18 < int64(0) {
						v126 = base.F64_neg(v122)
					} else {
						v126 = v122
					}
					v127 = v126
					v139 = v127
				}
			}
		}
		if base.F64_ne(base.F64_abs(v139), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v147 = v139
			v148 = F_Float8GetDatum(m, v147)
			mBase = m.M
			v149 = m.ExcPending
			if v149 != 0 {
				return int32(0)
			} else {
				return v148
			}
		} else {
			F_float_overflow_error(m)
			mBase = m.M
			v146 = m.ExcPending
			if v146 != 0 {
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
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 float64
	_ = v57
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v95 float64
	_ = v95
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v129 float64
	_ = v129
	var v130 float64
	_ = v130
	var v139 float64
	_ = v139
	var v143 float64
	_ = v143
	var v146 float64
	_ = v146
	var v153 int32
	_ = v153
	var v154 float64
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = math.Float64frombits(uint64(0x7ff8000000000000))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) {
		v154 = v10
		v156 = F_Float8GetDatum(m, v154)
		mBase = m.M
		v157 = m.ExcPending
		if v157 != 0 {
			return int32(0)
		} else {
			m.G0 = v8 + int32(16)
			return v156
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v19)&int64(9223372036854775807)) {
			v154 = v10
			v156 = F_Float8GetDatum(m, v154)
			mBase = m.M
			v157 = m.ExcPending
			if v157 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return v156
			}
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, _consts[793])))
			if v26 == int32(0) {
				F_init_degree_constants(m)
				mBase = m.M
			} else {
			}
			if base.Ui64(base.I64_reinterpret_f64(v19)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
				if base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
					v47 = base.I64_reinterpret_f64(v19)
					v50 = base.I32_wrap_i64(int64(base.Ui64(v47) >> (uint(int64(32)) % 64)))
					v53 = base.I32_wrap_i64(v47)
					if v50-int32(1072693248)|v53 == int32(0) {
						v57 = F_atan(m, v12)
						mBase = m.M
						v139 = v57
					} else {
						v61 = int32(base.Ui32(v50)>>(uint(int32(30))%32)) & int32(2)
						v62 = base.I64_reinterpret_f64(v12)
						v66 = v61 | base.I32_wrap_i64(int64(base.Ui64(v62)>>(uint(int64(63))%64)))
						v71 = base.I32_wrap_i64(int64(base.Ui64(v62)>>(uint(int64(32))%64))) & int32(2147483647)
						if v71|base.I32_wrap_i64(v62) == int32(0) {
							switch v66 - int32(2) {
							case 0:
								v139 = float64(3.141592653589793)
							case 1:
								v139 = float64(-3.141592653589793)
							default:
								v130 = v12
								v139 = v130
							}
						} else {
							v81 = v50 & int32(2147483647)
							if v81|v53 == int32(0) {
								v139 = base.F64_copysign(float64(1.5707963267948966), v12)
							} else {
								if v81 == int32(2146435072) {
									if v71 != int32(2146435072) {
										v129 = *(*float64)(unsafe.Add(mBase, uint32(v66<<(uint(int32(3))%32))+uint32(_consts[794])))
										v130 = v129
										v139 = v130
									} else {
										v95 = *(*float64)(unsafe.Add(mBase, uint32(v66<<(uint(int32(3))%32))+uint32(_consts[795])))
										v139 = v95
									}
								} else {
									if base.B2i32(v71 != int32(2146435072))&base.B2i32(base.Ui32(v71) <= base.Ui32(v81+int32(67108864))) == int32(0) {
										v139 = base.F64_copysign(float64(1.5707963267948966), v12)
									} else {
										if v61 != 0 {
											if base.Ui32(v71+int32(67108864)) < base.Ui32(v81) {
												v113 = float64(0)
											} else {
												v112 = F_atan(m, base.F64_abs(base.F64_div(v12, v19)))
												mBase = m.M
												v113 = v112
											}
										} else {
											v112 = F_atan(m, base.F64_abs(base.F64_div(v12, v19)))
											mBase = m.M
											v113 = v112
										}
										switch v66 - int32(1) {
										case 0:
											v139 = base.F64_neg(v113)
										case 1:
											v139 = base.F64_sub(float64(3.141592653589793), base.F64_add(v113, float64(-1.2246467991473532e-16)))
										case 2:
											v139 = base.F64_add(base.F64_add(v113, float64(-1.2246467991473532e-16)), float64(-3.141592653589793))
										default:
											v130 = v113
											v139 = v130
										}
									}
								}
							}
						}
					}
				} else {
					v139 = base.F64_add(v12, v19)
				}
			} else {
				v139 = base.F64_add(v12, v19)
			}
			*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v139
			v143 = *(*float64)(unsafe.Add(mBase, _consts[796]))
			v146 = base.F64_mul(base.F64_div(v139, v143), float64(45))
			if base.F64_ne(base.F64_abs(v146), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v154 = v146
				v156 = F_Float8GetDatum(m, v154)
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v156
				}
			} else {
				F_float_overflow_error(m)
				mBase = m.M
				v153 = m.ExcPending
				if v153 != 0 {
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
		v10 = int32(4800)
	} else {
		v10 = int32(4799)
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
	v31 = base.I32_div_s((v26+l1)*int32(7834), int32(256))
	v34 = l2 + v11*int32(365) + v16 + v19 + v22 + v31 - int32(32167)
	v43 = int32(4799) + l0
	v48 = base.I32_div_s(v43, int32(4))
	v51 = base.I32_div_s(v43, int32(-100))
	v54 = base.I32_div_s(v43, int32(400))
	v63 = base.I32_div_s(int32(109676), int32(256))
	v66 = int32(4) + v43*int32(365) + v48 + v51 + v54 + v63 - int32(32167)
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
		v90 = int32(4799) + v81
		v95 = base.I32_div_s(v90, int32(4))
		v98 = base.I32_div_s(v90, int32(-100))
		v101 = base.I32_div_s(v90, int32(400))
		v110 = base.I32_div_s(int32(109676), int32(256))
		v113 = int32(4) + v90*int32(365) + v95 + v98 + v101 + v110 - int32(32167)
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
	if int32(357) <= v34+v127-v126 {
		v142 = v125 + int32(4800)
		v147 = base.I32_div_s(v142, int32(4))
		v150 = base.I32_div_s(v142, int32(-100))
		v153 = base.I32_div_s(v142, int32(400))
		v162 = base.I32_div_s(int32(109676), int32(256))
		v165 = int32(4) + v142*int32(365) + v147 + v150 + v153 + v162 - int32(32167)
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
		v9 = int32(4800)
	} else {
		v9 = int32(4799)
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
	v30 = base.I32_div_s((v25+l1)*int32(7834), int32(256))
	return l2 + v10*int32(365) + v15 + v18 + v21 + v30 - int32(32167)
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v95 int64
	_ = v95
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
	v95 = v5
	goto L5
L5:
	;
	m.G0 = v8 + int32(2176)
	return v95
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
	v88 = v5
	goto L9
L9:
	;
	F_FreeDir(m, v10)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L34
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v22 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v88 = v81
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
	v82 = F_ReadDir(m, v10, l0)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L32
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v18 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	v47 = F_pg_snprintf(m, v8+int32(128), int32(2048), int32(186013), v8+int32(16))
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
		v81 = v20
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
		v81 = v20
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L17
L22:
	;
	v55 = F___fstatat(m, int32(-100), v8+int32(128), v8+int32(32), int32(0))
	mBase = m.M
	goto L23
L23:
	;
	if v55 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v59 == int32(44) {
		v81 = v20
		goto L16
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v8)+56))
	v81 = v79 + v20
	goto L16
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(128)
	F_errmsg(m, int32(309580), v8)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(518815), int32(105), int32(355042))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
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
	if v82 != 0 {
		v18 = v82
		v20 = v81
		goto L10
	} else {
		goto L33
	}
L33:
	;
	goto L11
L34:
	;
	v95 = v88
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
					F_errcode(m, int32(8389))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_errmsg(m, int32(14012), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							F_errfinish(m, int32(519586), int32(517), int32(87031))
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
		*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(0)
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
					F_errmsg(m, int32(418208), int32(0))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(512946), int32(1925), int32(93088))
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
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
	F_appendStringInfoString(m, l0, int32(568949))
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
	F_appendStringInfoString(m, l0, int32(529112))
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
	v38 = v30 + l3<<(uint(v26)%32)
	v40 = int32(0)
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
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v40<<(uint(int32(1))%32)))))
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v47
	F_appendStringInfo(m, l0, int32(529074), v13+int32(16))
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
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
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
	F_appendStringInfoString(m, l0, int32(6890))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L22
	}
L14:
	;
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38+int32(2)+v62<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v73
	F_appendStringInfo(m, l0, int32(64331), v13)
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
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if v62 < v78-int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_appendStringInfoString(m, l0, int32(774448))
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
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	v86 = v85
	goto L19
L21:
	;
	goto L15
L22:
	;
	if v40 < l3-v26 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_appendStringInfoString(m, l0, int32(774448))
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
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	v108 = int32(1)
	v114 = v40 + v108
	if v114 != l3 {
		v38 = v38 + v107<<(uint(v108)%32) + int32(2)
		v40 = v114
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
	F_appendStringInfo(m, l0, int32(568687), v11+int32(80))
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
	v66 = v34 + ((v41+int32(1))&int32(131070)+int32(9))&int32(262142)
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
	F_appendStringInfo(m, l0, int32(698728), v11-int32(-64))
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
	F_appendStringInfo(m, l0, int32(702338), v11+int32(32))
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
	F_appendStringInfo(m, l0, int32(567411), v11)
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
	F_appendStringInfo(m, l0, int32(702696), v11+int32(16))
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
	F_appendStringInfo(m, l0, int32(697458), v11+int32(48))
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
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
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == int32(1) {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		switch v13 - int32(1) {
		case 0:
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
			v50 = F_detoast_attr(m, v49)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				if v49 != v50 {
					v170 = v50
					m.G0 = v8 + int32(32)
					return v170
				} else {
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
					if v53 == int32(1) {
						v56 = int32(6)
						v58 = int32(18)
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
						if v60 == v58 {
							v63 = v58
						} else {
							v63 = int32(2)
						}
						if v60&int32(254) == int32(2) {
							v68 = v56
						} else {
							v68 = v63
						}
						if v60 == int32(1) {
							v71 = v56
						} else {
							v71 = v68
						}
						v80 = v71
					} else {
						v72 = int32(1)
						if v53&v72 != 0 {
							v80 = int32(base.Ui32(v53) >> (uint(v72) % 32))
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
							v80 = int32(base.Ui32(v76) >> (uint(int32(2)) % 32))
						}
					}
					v81 = F_palloc(m, v80)
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
						if v83 == int32(1) {
							v86 = int32(6)
							v88 = int32(18)
							v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
							if v90 == v88 {
								v93 = v88
							} else {
								v93 = int32(2)
							}
							if v90&int32(254) == int32(2) {
								v98 = v86
							} else {
								v98 = v93
							}
							if v90 == int32(1) {
								v101 = v86
							} else {
								v101 = v98
							}
							if v101 != 0 {
								v102 = F__emscripten_memcpy_bulkmem(m, v81, v50, v101)
								mBase = m.M
							} else {
							}
							v170 = v81
						} else {
							if v83&int32(1) != 0 {
								v107 = int32(base.Ui32(v83) >> (uint(int32(1)) % 32))
								if v107 != 0 {
									v108 = F__emscripten_memcpy_bulkmem(m, v81, v50, v107)
									mBase = m.M
								} else {
								}
								v170 = v81
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
								v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
								if v112 != 0 {
									v113 = F__emscripten_memcpy_bulkmem(m, v81, v50, v112)
									mBase = m.M
								} else {
								}
								v170 = v81
							}
						}
						m.G0 = v8 + int32(32)
						return v170
					}
				}
			}
		default:
			if v13&int32(254) != int32(2) {
				v150 = int32(1)
				v153 = int32(base.Ui32(v150) >> (uint(int32(1)) % 32))
				v155 = v153 + int32(3)
				v156 = F_palloc(m, v155)
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v156))) = v155 << (uint(int32(2)) % 32)
					v163 = int32(1)
					v166 = v153 - v163
					if v166 != 0 {
						v167 = F__emscripten_memcpy_bulkmem(m, v156+int32(4), l0+v163, v166)
						mBase = m.M
					} else {
					}
					v170 = v156
					m.G0 = v8 + int32(32)
					return v170
				}
			} else {
				v120 = F_detoast_external_attr(m, l0)
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int32(0)
				} else {
					v170 = v120
					m.G0 = v8 + int32(32)
					return v170
				}
			}
		case 17:
			v16 = F_toast_fetch_datum(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				if v20&int32(3) != int32(2) {
					v170 = v16
					m.G0 = v8 + int32(32)
					return v170
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
					v27 = int32(base.Ui32(v25) >> (uint(int32(30)) % 32))
					switch v27 {
					case 0:
						v45 = F_pglz_decompress_datum(m, v16)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v16)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v170 = v45
								m.G0 = v8 + int32(32)
								return v170
							}
						}
					case 1:
						v28 = F_lz4_decompress_datum(m)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v16)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v170 = v28
								m.G0 = v8 + int32(32)
								return v170
							}
						}
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v27
							F_errmsg_internal(m, int32(497468), v8)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(511955), int32(489), int32(297675))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
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
		if v10&int32(3) == int32(2) {
			v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v128 = int32(base.Ui32(v126) >> (uint(int32(30)) % 32))
			switch v128 {
			case 0:
				v129 = F_pglz_decompress_datum(m, l0)
				mBase = m.M
				v130 = m.ExcPending
				if v130 != 0 {
					return int32(0)
				} else {
					v170 = v129
					m.G0 = v8 + int32(32)
					return v170
				}
			case 1:
				v131 = F_lz4_decompress_datum(m)
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return int32(0)
				} else {
					v170 = v131
					m.G0 = v8 + int32(32)
					return v170
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v128
					F_errmsg_internal(m, int32(497468), v8+int32(16))
					mBase = m.M
					v142 = m.ExcPending
					if v142 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(511955), int32(489), int32(297675))
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
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
				v150 = v10
				v153 = int32(base.Ui32(v150) >> (uint(int32(1)) % 32))
				v155 = v153 + int32(3)
				v156 = F_palloc(m, v155)
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v156))) = v155 << (uint(int32(2)) % 32)
					v163 = int32(1)
					v166 = v153 - v163
					if v166 != 0 {
						v167 = F__emscripten_memcpy_bulkmem(m, v156+int32(4), l0+v163, v166)
						mBase = m.M
					} else {
					}
					v170 = v156
					m.G0 = v8 + int32(32)
					return v170
				}
			} else {
				v170 = l0
				m.G0 = v8 + int32(32)
				return v170
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
	var v43 int32
	_ = v43
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
	var v68 int32
	_ = v68
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
	var v86 int32
	_ = v86
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
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
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
	var v203 int32
	_ = v203
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
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
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L66
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
	v39 = int32(293098)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1124])))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v43 == int32(0) {
		v62 = v42
		v63 = v43
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L4
L9:
	;
	v202 = v27 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v202 < v203 {
		v27 = v202
		goto L7
	} else {
		goto L65
	}
L10:
	;
	if v63-v62 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	goto L10
L12:
	;
	if v42 != v43 {
		v62 = v42
		v63 = v43
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v47 = v38
	v48 = v39
	goto L14
L14:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	if v52 == int32(0) {
		v62 = v51
		v63 = v52
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v62 = v51
	v63 = v52
	goto L11
L16:
	;
	v55 = int32(1)
	if v51 == v52 {
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
	v67 = F_defGetString(m, v37)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v140 = int32(341589)
	v143 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1125])))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v144 == int32(0) {
		v163 = v143
		v164 = v144
		goto L44
	} else {
		goto L45
	}
L21:
	;
	v72 = v67
	goto L23
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v116
	if int32(0) < v116 {
		goto L9
	} else {
		goto L38
	}
L23:
	;
	v77 = v72 + int32(1)
	v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v72))))
	v79 = F___isspace(m, v78)
	mBase = m.M
	if v79 != 0 {
		v72 = v77
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v80 = int32(1)
	switch v78&int32(255) - int32(43) {
	case 0:
		v86 = v80
		goto L27
	default:
		v88 = v78
		v89 = v72
		v90 = v80
		goto L26
	case 2:
		goto L28
	}
L25:
	;
	goto L24
L26:
	;
	v91 = int32(0)
	v93 = v88 - int32(48)
	if base.Ui32(v93) <= base.Ui32(int32(9)) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(v77))))
	v88 = v87
	v89 = v77
	v90 = v86
	goto L26
L28:
	;
	v86 = int32(0)
	goto L27
L29:
	;
	v96 = v91
	v97 = v93
	v98 = v89
	goto L32
L30:
	;
	v110 = v91
	goto L31
L31:
	;
	if v90 != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v100 = int32(10)
	v102 = v96*v100 - v97
	v103 = int32(*(*int8)(unsafe.Add(mBase, uint32(v98)+1)))
	v107 = v103 - int32(48)
	if base.Ui32(v107) < base.Ui32(v100) {
		v96 = v102
		v97 = v107
		v98 = v98 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v110 = v102
	goto L31
L34:
	;
	goto L33
L35:
	;
	v116 = int32(0) - v110
	goto L37
L36:
	;
	v116 = v110
	goto L37
L37:
	;
	goto L22
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(584707), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(512500), int32(57), int32(105921))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	if v164-v163 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L44:
	;
	goto L43
L45:
	;
	if v143 != v144 {
		v163 = v143
		v164 = v144
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v148 = v38
	v149 = v140
	goto L47
L47:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	if v153 == int32(0) {
		v163 = v152
		v164 = v153
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v163 = v152
	v164 = v153
	goto L44
L49:
	;
	v156 = int32(1)
	if v152 == v153 {
		v148 = v148 + v156
		v149 = v149 + v156
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v168 = F_defGetBoolean(m, v37)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v171 = int32(321161)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1126])))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v175 == int32(0) {
		v194 = v174
		v195 = v175
		goto L56
	} else {
		goto L57
	}
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)) = uint8(v168)
	goto L9
L55:
	;
	if v195-v194 != 0 {
		goto L3
	} else {
		goto L63
	}
L56:
	;
	goto L55
L57:
	;
	if v174 != v175 {
		v194 = v174
		v195 = v175
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v179 = v38
	v180 = v171
	goto L59
L59:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	if v184 == int32(0) {
		v194 = v183
		v195 = v184
		goto L56
	} else {
		goto L61
	}
L60:
	;
	v194 = v183
	v195 = v184
	goto L56
L61:
	;
	v187 = int32(1)
	if v183 == v184 {
		v179 = v179 + v187
		v180 = v180 + v187
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v197 = F_defGetBoolean(m, v37)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)) = uint8(v197)
	goto L9
L65:
	;
	goto L8
L66:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v222
	F_errmsg(m, int32(753582), v9)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(512500), int32(72), int32(105921))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
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
func F_do_tzset(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1155])))
	if v2&int32(1) == int32(0) {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1155])))
		if v8&int32(1) == int32(0) {
			v15 = int32(4713840)
			v16 = int32(4713872)
			m.Env.X_tzset_js(m, int32(4713792), int32(4713796), v15, v16)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, _consts[1156])) = v16
			*(*int32)(unsafe.Add(mBase, _consts[1157])) = v15
			v25 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[1155])) = uint8(v25)
		} else {
		}
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
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
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v308 int32
	_ = v308
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
		v225 = l0
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+4)))
	if v234&int32(2) != 0 {
		v288 = v225
		goto L75
	} else {
		goto L76
	}
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v23 != v25 {
		v225 = l0
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v27&int32(2) != 0 {
		v225 = l0
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
	if v214 != 0 {
		v225 = v214
		goto L7
	} else {
		goto L73
	}
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v32 != v33 {
		v225 = l0
		goto L7
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v193 != v194 {
		v225 = l0
		goto L7
	} else {
		goto L64
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
	if v35 <= v36 {
		v225 = l0
		goto L7
	} else {
		goto L27
	}
L19:
	;
	if v38 == int32(0) {
		v225 = l0
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
	v214 = v53
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
	if v36 <= int32(0) {
		v225 = l0
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v59 = F_palloc0(m, v35)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v61 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v115 == v113 {
		goto L44
	} else {
		goto L45
	}
L31:
	;
	v113 = v5
	goto L30
L32:
	;
	goto L33
L33:
	;
	v64 = int32(0)
	v70 = v64
	v71 = v64
	v73 = v5
	goto L34
L34:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v75 <= v71 {
		v113 = v73
		goto L30
	} else {
		goto L36
	}
L35:
	;
	v113 = v101
	goto L30
L36:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v78 = int32(2)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77+v70<<(uint(v78)%32))))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v71<<(uint(v78)%32))))
	v87 = F_QTNodeCompare(m, v81, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v103 = v70 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v103 < v104 {
		v70 = v103
		v71 = v100
		v73 = v101
		goto L34
	} else {
		goto L43
	}
L38:
	;
	if v87 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v92 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v70+v59))) = uint8(v92)
	v100 = v71 + v92
	v101 = v73 + v92
	goto L37
L40:
	;
	goto L41
L41:
	;
	if int32(0) <= v87 {
		v113 = v73
		goto L30
	} else {
		goto L42
	}
L42:
	;
	v100 = v71
	v101 = v73
	goto L37
L43:
	;
	goto L35
L44:
	;
	v117 = int32(0)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v117 < v118 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	F_pfree(m, v59)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L63
	}
L47:
	;
	v126 = int32(0)
	v128 = v117
	goto L50
L48:
	;
	v159 = v117
	goto L49
L49:
	;
	if l2 != 0 {
		goto L58
	} else {
		goto L59
	}
L50:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131+v126<<(uint(int32(2))%32))))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126+v59))))
	if v137 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v159 = v148
	goto L49
L52:
	;
	v150 = v126 + int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v150 < v151 {
		v126 = v150
		v128 = v148
		goto L50
	} else {
		goto L57
	}
L53:
	;
	F_QTNFree(m, v135)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131+v128<<(uint(int32(2))%32)))) = v135
	v148 = v128 + int32(1)
	goto L52
L56:
	;
	v148 = v128
	goto L52
L57:
	;
	goto L51
L58:
	;
	v162 = F_QTNCopy(m, l2)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	v176 = v159
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v176
	F_QTNSort(m, l0)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v165 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v164 | v165
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v168+v159<<(uint(v165)%32)))) = v162
	v176 = v159 + int32(1)
	goto L60
L62:
	;
	v180 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v180)
	goto L46
L63:
	;
	v214 = l0
	goto L11
L64:
	;
	v196 = F_QTNEq(m, l0, l1)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v196 == int32(0) {
		v225 = l0
		goto L7
	} else {
		goto L66
	}
L66:
	;
	F_QTNFree(m, l0)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if l2 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v212 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v212)
	v214 = v211
	goto L11
L69:
	;
	v211 = int32(0)
	goto L68
L70:
	;
	goto L71
L71:
	;
	v205 = F_QTNCopy(m, l2)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v205)+4)) = v207 | int32(2)
	v211 = v205
	goto L68
L73:
	;
	return int32(0)
L74:
	;
	F_QTNFree(m, v225)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L88
	}
L75:
	;
	return v288
L76:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	if v238 != int32(2) {
		v288 = v225
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	if v241 <= int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+8)) = int32(0)
	goto L74
L79:
	;
	goto L80
L80:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v225)+20))
	v247 = int32(0)
	v253 = v247
	v254 = v247
	v255 = v246
	goto L81
L81:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v255+v254<<(uint(int32(2))%32))))
	v262 = F_dofindsubquery(m, v261, l1, l2, l3)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L83
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+8)) = v274
	switch v274 {
	case 0:
		goto L74
	case 1:
		goto L85
	default:
		v288 = v225
		goto L75
	}
L83:
	;
	v265 = v253 << (uint(int32(2)) % 32)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v225)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v265+v266))) = v262
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v225)+20))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v269+v265)))
	v274 = v253 + base.B2i32(v271 != int32(0))
	v276 = v254 + int32(1)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	if v276 < v277 {
		v253 = v274
		v254 = v276
		v255 = v269
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)))
	if v281 == int32(1) {
		v288 = v225
		goto L75
	} else {
		goto L86
	}
L86:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v225)+20))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	F_pfree(m, v225)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v288 = v285
	goto L75
L88:
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
	var v28 int32
	_ = v28
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
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v67 == int32(0) {
		v95 = v66
		goto L29
	} else {
		goto L30
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
	if int32(0) < v54 {
		v13 = v54
		v16 = v56
		goto L7
	} else {
		goto L28
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v31 == int32(0) {
		goto L19
	} else {
		goto L20
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
	v26 = F__emscripten_memset_bulkmem(m, v16, base.I32_extend8_s(v1), v24)
	mBase = m.M
	goto L18
L18:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v28 = v27 + v24
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v28
	v54 = v13 - v24
	v56 = v28
	goto L9
L19:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v34 + v13
	return
L20:
	;
	goto L21
L21:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v37 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v52
	v54 = v13
	v56 = v52
	goto L9
L23:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v16 == v38 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v41 = v16 - v38
	v42 = F_fwrite(m, v38, int32(1), v41, v31)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return
L26:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v42 + v44
	if v42 == v41 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v48 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v48)
	goto L22
L28:
	;
	goto L8
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v95 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v1)
	return
L30:
	;
	if base.Ui32(v66) < base.Ui32(v67) {
		v95 = v66
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v71 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v74 + int32(1)
	return
L33:
	;
	goto L34
L34:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v78 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v95 = v94
	goto L29
L36:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v66 == v79 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v82 = v66 - v79
	v83 = F_fwrite(m, v79, int32(1), v82, v71)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v83 + v85
	if v82 == v83 {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v89 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v89)
	goto L35
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
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
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
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
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v212 == int32(0) {
		v241 = v210
		goto L74
	} else {
		goto L75
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v12 = l0
	v13 = l1
	v15 = v11
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
	if int32(0) < v198 {
		v12 = v197
		v13 = v198
		v15 = v199
		goto L7
	} else {
		goto L73
	}
L10:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v174 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L11:
	;
	v19 = v18 - v15
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
	if v15 == v12 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v170 = v169 + v24
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v170
	v197 = v12 + v24
	v198 = v13 - v24
	v199 = v170
	goto L9
L19:
	;
	goto L18
L20:
	;
	v28 = v15 + v24
	if base.Ui32(v12-v28) <= base.Ui32(int32(0)-v24<<(uint(int32(1))%32)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v35 = F___memcpy(m, v15, v12, v24)
	mBase = m.M
	goto L18
L22:
	;
	goto L23
L23:
	;
	v38 = (v15 ^ v12) & int32(3)
	if base.Ui32(v15) < base.Ui32(v12) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	if v140 == int32(0) {
		goto L19
	} else {
		goto L60
	}
L25:
	;
	if base.Ui32(v118) <= base.Ui32(int32(3)) {
		v139 = v117
		v140 = v118
		v141 = v119
		goto L24
	} else {
		goto L56
	}
L26:
	;
	if v38 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v38 != 0 {
		v100 = v24
		goto L39
	} else {
		goto L40
	}
L29:
	;
	v139 = v12
	v140 = v24
	v141 = v15
	goto L24
L30:
	;
	goto L31
L31:
	;
	if v15&int32(3) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v117 = v12
	v118 = v24
	v119 = v15
	goto L25
L33:
	;
	goto L34
L34:
	;
	v45 = v12
	v46 = v24
	v47 = v15
	goto L35
L35:
	;
	if v46 == int32(0) {
		goto L19
	} else {
		goto L37
	}
L36:
	;
	v117 = v54
	v118 = v56
	v119 = v58
	goto L25
L37:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v51)
	v53 = int32(1)
	v54 = v45 + v53
	v56 = v46 - v53
	v58 = v47 + v53
	if v58&int32(3) != 0 {
		v45 = v54
		v46 = v56
		v47 = v58
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	if v100 == int32(0) {
		goto L19
	} else {
		goto L52
	}
L40:
	;
	if v28&int32(3) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v65 = v24
	goto L44
L42:
	;
	v80 = v24
	goto L43
L43:
	;
	if base.Ui32(v80) <= base.Ui32(int32(3)) {
		v100 = v80
		goto L39
	} else {
		goto L48
	}
L44:
	;
	if v65 == int32(0) {
		goto L19
	} else {
		goto L46
	}
L45:
	;
	v80 = v71
	goto L43
L46:
	;
	v71 = v65 - int32(1)
	v72 = v15 + v71
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v71))))
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v74)
	if v72&int32(3) != 0 {
		v65 = v71
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v87 = v80
	goto L49
L49:
	;
	v91 = v87 - int32(4)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v12+v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v15+v91))) = v94
	if base.Ui32(int32(3)) < base.Ui32(v91) {
		v87 = v91
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v100 = v91
	goto L39
L51:
	;
	goto L50
L52:
	;
	v107 = v100
	goto L53
L53:
	;
	v111 = v107 - int32(1)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v111))))
	*(*uint8)(unsafe.Add(mBase, uint32(v15+v111))) = uint8(v114)
	if v111 != 0 {
		v107 = v111
		goto L53
	} else {
		goto L55
	}
L54:
	;
	goto L19
L55:
	;
	goto L54
L56:
	;
	v124 = v117
	v125 = v118
	v126 = v119
	goto L57
L57:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v128
	v130 = int32(4)
	v131 = v124 + v130
	v133 = v126 + v130
	v135 = v125 - v130
	if base.Ui32(int32(3)) < base.Ui32(v135) {
		v124 = v131
		v125 = v135
		v126 = v133
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v139 = v131
	v140 = v135
	v141 = v133
	goto L24
L59:
	;
	goto L58
L60:
	;
	v146 = v139
	v147 = v140
	v148 = v141
	goto L61
L61:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	*(*uint8)(unsafe.Add(mBase, uint32(v148))) = uint8(v150)
	v152 = int32(1)
	v157 = v147 - v152
	if v157 != 0 {
		v146 = v146 + v152
		v147 = v157
		v148 = v148 + v152
		goto L61
	} else {
		goto L63
	}
L62:
	;
	goto L19
L63:
	;
	goto L62
L64:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v177 + v13
	return
L65:
	;
	goto L66
L66:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v180 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v195
	v197 = v12
	v198 = v13
	v199 = v195
	goto L9
L68:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v15 == v181 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v184 = v15 - v181
	v185 = F_fwrite(m, v181, int32(1), v184, v174)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	return
L71:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v185 + v187
	if v184 == v185 {
		goto L67
	} else {
		goto L72
	}
L72:
	;
	v191 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v191)
	goto L67
L73:
	;
	goto L8
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v241 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v241))) = uint8(v211)
	return
L75:
	;
	if base.Ui32(v210) < base.Ui32(v212) {
		v241 = v210
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v216 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v219 + int32(1)
	return
L78:
	;
	goto L79
L79:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v223 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v241 = v239
	goto L74
L81:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v210 == v224 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v227 = v210 - v224
	v228 = F_fwrite(m, v224, int32(1), v227, v216)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L70
	} else {
		goto L83
	}
L83:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v228 + v230
	if v227 == v228 {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v234 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v234)
	goto L80
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
							F_errmsg(m, int32(103808), int32(0))
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(512946), int32(1526), int32(32066))
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
							*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(0)
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
						F_errmsg(m, int32(470212), int32(0))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(512946), int32(1522), int32(32066))
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
	F_CatalogTupleDelete(m, v15, v46+int32(4))
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
	F_sequence_close(m, v15, int32(3))
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
	var v87 int32
	_ = v87
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
	var v221 int32
	_ = v221
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
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v283 int32
	_ = v283
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v457 int32
	_ = v457
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
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
		goto L16
	} else {
		goto L17
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L162
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L158
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L154
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L150
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L145
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L142
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L137
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L133
	}
L16:
	;
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+104)))
	if int32(0) < v39 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+72)))
	if v42 == int32(110) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v47 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v210 = v42
	v221 = int32(0)
	goto L22
L22:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+106)))
	if v210&int32(255) != int32(102) {
		goto L62
	} else {
		goto L63
	}
L23:
	;
	v49 = F_extractNotNullColumn(m, l2)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v53 = F_get_attname(m, v51, v49, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v56 = F_RelationGetIndexAttrBitmap(m, l1, int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v179 = F_RelationGetIndexAttrBitmap(m, l1, int32(2))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L51
	}
L27:
	;
	F_relation_close(m, v71, int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L50
	}
L28:
	;
	if v56 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+119)))
	if v61 != int32(112) {
		goto L26
	} else {
		goto L32
	}
L30:
	;
	v113 = v56
	goto L31
L31:
	;
	v130 = F_bms_is_member(m, v49+int32(7), v113)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L43
	}
L32:
	;
	v65 = F_RelationGetPrimaryKeyIndex(m, l1, int32(1))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v65 == int32(0) {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v71 = F_relation_open(m, v65, int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)+192))
	v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+10)))
	if v74 <= int32(0) {
		goto L27
	} else {
		goto L36
	}
L36:
	;
	v80 = int32(0)
	v83 = int32(0)
	v87 = v73
	goto L37
L37:
	;
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87+v83<<(uint(int32(1))%32))+48)))
	v99 = F_bms_add_member(m, v80, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	F_relation_close(m, v71, int32(1))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v102 = v83 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v71)+192))
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+10)))
	if v102 < v104 {
		v80 = v99
		v83 = v102
		v87 = v103
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	if v99 == int32(0) {
		goto L26
	} else {
		goto L42
	}
L42:
	;
	v113 = v99
	goto L31
L43:
	;
	if v130 == int32(0) {
		goto L26
	} else {
		goto L44
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v143 = F_get_attname(m, v141, v49, int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v143
	F_errmsg(m, int32(21649), v20+int32(96))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(514330), int32(14159), int32(323288))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	goto L26
L51:
	;
	v181 = F_bms_is_member(m, v49+int32(7), v179)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v181 != 0 {
		goto L14
	} else {
		goto L53
	}
L53:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v184 = F_SearchSysCacheCopyAttNum(m, v183, v49)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v184 == int32(0) {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+22)))
	v190 = v188 + v189
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+89)))
	if v191 != 0 {
		goto L12
	} else {
		goto L56
	}
L56:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+86)))
	if v192 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+86)) = uint8(v195)
	F_CatalogTupleUpdate(m, v47, v184+int32(4), v184)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_sequence_close(m, v47, int32(3))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+72)))
	v210 = v204
	v221 = v53
	goto L22
L62:
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
		goto L72
	}
L63:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v34)+96))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v227 == v228 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v231 = F_table_open(m, v227, int32(8))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v231)+48))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+118)))
	if v234 == int32(116) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+24)))
	if v237 == int32(0) {
		goto L11
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	F_CheckTableNotInUse(m, v231, int32(562014))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	F_sequence_close(m, v231, int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L62
L72:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+72)))
	switch v256 - int32(99) {
	case 0, 11:
		goto L74
	default:
		goto L75
	}
L73:
	;
	m.G0 = v20 + int32(272)
	return
L74:
	;
	if v222&int32(1) != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+119)))
	if v260 != int32(112) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	F_sequence_close(m, v30, int32(3))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L73
L78:
	;
	F_sequence_close(m, v30, int32(3))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L132
	}
L79:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v269 = F_find_inheritance_children(m, v268, l6)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v269 == int32(0) {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v273 <= int32(0) {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	v283 = int32(0)
	goto L83
L83:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v298+v283<<(uint(int32(2))%32))))
	v304 = F_table_open(m, v302, int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	goto L78
L85:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v304)+48))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+118)))
	if v307 == int32(116) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+24)))
	if v310 == int32(0) {
		goto L10
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	F_CheckTableNotInUse(m, v304, int32(562014))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+72)))
	if v316 == int32(110) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+22)))
	v376 = v374 + v375
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+72)))
	switch v377 - int32(99) {
	case 0, 11:
		goto L108
	default:
		goto L109
	}
L92:
	;
	v319 = F_findNotNullConstraint(m, v302, v221)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	F_ScanKeyInit(m, v20+int32(128), int32(9), int32(3), int32(184), v302)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L100
	}
L95:
	;
	if v319 != 0 {
		v372 = v319
		goto L91
	} else {
		goto L96
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v304)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v221
	F_errmsg_internal(m, int32(49742), v20+int32(32))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(514330), int32(14260), int32(323288))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
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
	F_ScanKeyInit(m, v20+int32(176), int32(10), int32(3), int32(184), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_ScanKeyInit(m, v20+int32(224), int32(2), int32(3), int32(62), v36)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v362 = F_systable_beginscan(m, v30, int32(2665), int32(1), int32(0), int32(3), v20+int32(128))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v364 = F_systable_getnext(m, v362)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	if v364 == int32(0) {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	v368 = F_heap_copytuple(m, v364)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_systable_endscan(m, v362)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v372 = v368
	goto L91
L108:
	;
	v393 = int32(*(*int16)(unsafe.Add(mBase, uint32(v376)+104)))
	if v393 <= int32(0) {
		goto L8
	} else {
		goto L113
	}
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errmsg_internal(m, int32(95693), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(514330), int32(14298), int32(323288))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	if l4 != 0 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	F_pfree(m, v372)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L129
	}
L115:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L128
	}
L116:
	;
	if v393 != int32(1) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	v413 = v393 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v376)+104)) = uint16(v413)
	if v413&int32(65535) == int32(0) {
		goto L124
	} else {
		goto L125
	}
L119:
	;
	v406 = v393 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v376)+104)) = uint16(v406)
	F_CatalogTupleUpdate(m, v30, v372+int32(4), v372)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L123
	}
L120:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+103)))
	if v398 != 0 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v401 = int32(1)
	F_dropconstraint_internal(m, v20+int32(128), v304, v372, l3, v401, v401, l6)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	goto L114
L123:
	;
	goto L115
L124:
	;
	v419 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v376)+103)) = uint8(v419)
	goto L126
L125:
	;
	goto L126
L126:
	;
	F_CatalogTupleUpdate(m, v30, v372+int32(4), v372)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L115
L128:
	;
	goto L114
L129:
	;
	F_sequence_close(m, v304, int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v435 = v283 + int32(1)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v435 < v436 {
		v283 = v435
		goto L83
	} else {
		goto L131
	}
L131:
	;
	goto L84
L132:
	;
	goto L73
L133:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v485 + int32(4)
	F_errmsg(m, int32(736170), v20+int32(112))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(514330), int32(14107), int32(323288))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v509 = F_get_attname(m, v507, v49, int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v509
	F_errmsg(m, int32(10484), v20)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(514330), int32(14167), int32(323288))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
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
	v524 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v49
	F_errmsg_internal(m, int32(49690), v20+int32(16))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(514330), int32(14173), int32(323288))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v546 = F_get_attname(m, v544, v49, int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v546
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v548 + int32(4)
	F_errmsg(m, int32(284396), v20+int32(80))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(514330), int32(14181), int32(323288))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errmsg(m, int32(151493), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(514330), int32(4460), int32(426109))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
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
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errmsg(m, int32(151493), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(514330), int32(4460), int32(426109))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v304)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v602 + int32(4)
	F_errmsg(m, int32(76951), v20-int32(-64))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(514330), int32(14288), int32(323288))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v376 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v302
	F_errmsg_internal(m, int32(727117), v20+int32(48))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(514330), int32(14302), int32(323288))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
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
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v74 int32
	_ = v74
	var v88 int32
	_ = v88
	var v93 int64
	_ = v93
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v103 int64
	_ = v103
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v125 int64
	_ = v125
	var v132 int64
	_ = v132
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v142 int32
	_ = v142
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
	var v151 int64
	_ = v151
	var v165 int64
	_ = v165
	var v176 float64
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 float64
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 float64
	_ = v190
	var v194 float64
	_ = v194
	var v198 float64
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v218 float64
	_ = v218
	var v222 int32
	_ = v222
	var v223 float64
	_ = v223
	var v224 float64
	_ = v224
	var v230 float64
	_ = v230
	var v231 float64
	_ = v231
	var v233 float64
	_ = v233
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	var v246 float64
	_ = v246
	var v254 float64
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v274 float64
	_ = v274
	var v278 int32
	_ = v278
	var v279 float64
	_ = v279
	var v280 float64
	_ = v280
	var v285 float64
	_ = v285
	var v287 float64
	_ = v287
	var v289 float64
	_ = v289
	var v292 float64
	_ = v292
	var v296 float64
	_ = v296
	var v300 float64
	_ = v300
	var v304 float64
	_ = v304
	var v310 float64
	_ = v310
	var v311 float64
	_ = v311
	var v316 float64
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
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
	v344 = m.ExcPending
	if v344 != 0 {
		goto L95
	} else {
		goto L101
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
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
	v316 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L5
L5:
	;
	v319 = F_Float8GetDatum(m, v316)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L95
	} else {
		goto L96
	}
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _consts[793])))
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
	v34 = base.I64_reinterpret_f64(v12)
	v38 = int32(2047)
	v39 = base.I32_wrap_i64(int64(base.Ui64(v34)>>(uint(int64(52))%64))) & v38
	if v39 != v38 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v178 = base.F64_lt(v176, float64(0))
	if v178 != 0 {
		goto L51
	} else {
		goto L52
	}
L11:
	;
	v46 = v34 << (uint(int64(1)) % 64)
	if base.Ui64(v46) <= base.Ui64(int64(-9156662467374350336)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v43 = base.F64_mul(v12, float64(360))
	v176 = base.F64_div(v43, v43)
	goto L10
L13:
	;
	if v46 == int64(-9156662467374350336) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if v39 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v53 = base.F64_mul(v12, float64(0))
	goto L18
L17:
	;
	v53 = v12
	goto L18
L18:
	;
	v176 = v53
	goto L10
L19:
	;
	if int32(1031) < v88 {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	v56 = int32(0)
	v58 = v34 << (uint(int64(12)) % 64)
	if int64(0) <= v58 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v88 = v39
	v93 = v34&int64(4503599627370495) | int64(4503599627370496)
	goto L19
L23:
	;
	v62 = v56
	v64 = v58
	goto L26
L24:
	;
	v74 = v56
	goto L25
L25:
	;
	v88 = v74
	v93 = v34 << (uint(base.I64_extend_i32_u(int32(1)-v74)) % 64)
	goto L19
L26:
	;
	v68 = v62 - int32(1)
	v70 = v64 << (uint(int64(1)) % 64)
	if int64(0) <= v70 {
		v62 = v68
		v64 = v70
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v74 = v68
	goto L25
L28:
	;
	goto L27
L29:
	;
	v97 = v88
	v99 = v93
	goto L32
L30:
	;
	v119 = v88
	v121 = v93
	goto L31
L31:
	;
	v125 = v121 - int64(6333186975989760)
	if v125 < int64(0) {
		v132 = v121
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v103 = v99 - int64(6333186975989760)
	if v103 < int64(0) {
		v110 = v99
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v119 = int32(1031)
	v121 = v112
	goto L31
L34:
	;
	v112 = v110 << (uint(int64(1)) % 64)
	v114 = v97 - int32(1)
	if int32(1031) < v114 {
		v97 = v114
		v99 = v112
		goto L32
	} else {
		goto L37
	}
L35:
	;
	if v103 != int64(0) {
		v110 = v103
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v176 = base.F64_mul(v12, float64(0))
	goto L10
L37:
	;
	goto L33
L38:
	;
	if base.Ui64(int64(4503599627370495)) < base.Ui64(v132) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	if v125 != int64(0) {
		v132 = v125
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v176 = base.F64_mul(v12, float64(0))
	goto L10
L41:
	;
	if int32(0) < v148 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v148 = v119
	v151 = v132
	goto L41
L43:
	;
	goto L44
L44:
	;
	v136 = v119
	v138 = v132
	goto L45
L45:
	;
	v142 = v136 - int32(1)
	v146 = v138 << (uint(int64(1)) % 64)
	if base.Ui64(v138) < base.Ui64(int64(2251799813685248)) {
		v136 = v142
		v138 = v146
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v148 = v142
	v151 = v146
	goto L41
L47:
	;
	goto L46
L48:
	;
	v165 = v151 - int64(4503599627370496) | base.I64_extend_i32_u(v148)<<(uint(int64(52))%64)
	goto L50
L49:
	;
	v165 = int64(base.Ui64(v151) >> (uint(base.I64_extend_i32_u(int32(1)-v148)) % 64))
	goto L50
L50:
	;
	v176 = base.F64_reinterpret_i64(v165 | v34&int64(-9223372036854775807-1))
	goto L10
L51:
	;
	v179 = int32(-1)
	goto L53
L52:
	;
	v179 = int32(1)
	goto L53
L53:
	;
	if v178 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v182 = base.F64_neg(v176)
	goto L56
L55:
	;
	v182 = v176
	goto L56
L56:
	;
	v184 = base.F64_gt(v182, float64(180))
	if v184 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v185 = int32(0) - v179
	goto L59
L58:
	;
	v185 = v179
	goto L59
L59:
	;
	if v184 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v311 = base.F64_mul(v310, base.F64_convert_i32_s(v185))
	if base.F64_eq(base.F64_abs(v311), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L94
	}
L61:
	;
	v190 = base.F64_sub(float64(360), v182)
	goto L63
L62:
	;
	v190 = v182
	goto L63
L63:
	;
	if base.F64_gt(v190, float64(90)) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v194 = base.F64_sub(float64(180), v190)
	goto L66
L65:
	;
	v194 = v190
	goto L66
L66:
	;
	if base.F64_le(v194, float64(30)) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v198 = base.F64_mul(v194, float64(0.017453292519943295))
	v202 = m.G0
	v204 = v202 - int32(16)
	m.G0 = v204
	v211 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v198))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v211) <= base.Ui32(int32(1072243195)) {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	goto L69
L69:
	;
	v254 = base.F64_mul(base.F64_sub(float64(90), v194), float64(0.017453292519943295))
	v258 = m.G0
	v260 = v258 - int32(16)
	m.G0 = v260
	v267 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v254))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v267) <= base.Ui32(int32(1072243195)) {
		goto L85
	} else {
		goto L86
	}
L70:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v237
	v246 = *(*float64)(unsafe.Add(mBase, _consts[797]))
	v310 = base.F64_mul(base.F64_div(v237, v246), float64(0.5))
	goto L60
L71:
	;
	m.G0 = v204 + int32(16)
	goto L70
L72:
	;
	if base.Ui32(v211) < base.Ui32(int32(1045430272)) {
		v237 = v198
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v211) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v218 = F___sin(m, v198, float64(0), int32(0))
	mBase = m.M
	v237 = v218
	goto L71
L76:
	;
	v237 = base.F64_sub(v198, v198)
	goto L71
L77:
	;
	goto L78
L78:
	;
	v222 = F___rem_pio2(m, v198, v204)
	mBase = m.M
	v223 = *(*float64)(unsafe.Add(mBase, uint32(v204)+8))
	v224 = *(*float64)(unsafe.Add(mBase, uint32(v204)))
	switch v222&int32(3) - int32(1) {
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
	v235 = F___cos(m, v224, v223)
	mBase = m.M
	v237 = base.F64_neg(v235)
	goto L71
L80:
	;
	v233 = F___sin(m, v224, v223, int32(1))
	mBase = m.M
	v237 = base.F64_neg(v233)
	goto L71
L81:
	;
	v231 = F___cos(m, v224, v223)
	mBase = m.M
	v237 = v231
	goto L71
L82:
	;
	v230 = F___sin(m, v224, v223, int32(1))
	mBase = m.M
	v237 = v230
	goto L71
L83:
	;
	v300 = base.F64_sub(float64(1), v296)
	*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v300
	v304 = *(*float64)(unsafe.Add(mBase, _consts[798]))
	v310 = base.F64_add(base.F64_mul(base.F64_div(v300, v304), float64(-0.5)), float64(1))
	goto L60
L84:
	;
	m.G0 = v260 + int32(16)
	goto L83
L85:
	;
	if base.Ui32(v267) < base.Ui32(int32(1044816030)) {
		v296 = float64(1)
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v267) {
		v296 = base.F64_sub(v254, v254)
		goto L84
	} else {
		goto L89
	}
L88:
	;
	v274 = F___cos(m, v254, float64(0))
	mBase = m.M
	v296 = v274
	goto L84
L89:
	;
	v278 = F___rem_pio2(m, v254, v260)
	mBase = m.M
	v279 = *(*float64)(unsafe.Add(mBase, uint32(v260)+8))
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v260)))
	switch v278&int32(3) - int32(1) {
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
	v292 = F___sin(m, v280, v279, int32(1))
	mBase = m.M
	v296 = v292
	goto L84
L91:
	;
	v289 = F___cos(m, v280, v279)
	mBase = m.M
	v296 = base.F64_neg(v289)
	goto L84
L92:
	;
	v287 = F___sin(m, v280, v279, int32(1))
	mBase = m.M
	v296 = base.F64_neg(v287)
	goto L84
L93:
	;
	v285 = F___cos(m, v280, v279)
	mBase = m.M
	v296 = v285
	goto L84
L94:
	;
	v316 = v311
	goto L5
L95:
	;
	return int32(0)
L96:
	;
	m.G0 = v8 + int32(16)
	return v319
L97:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L95
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(418208), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L95
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(512946), int32(2455), int32(442200))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
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
				F_errmsg(m, int32(238073), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(512946), int32(1454), int32(83719))
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
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v2
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+30)) = uint16(v2)
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v156 = F_get_tsearch_config_filename(m, v124, int32(253824))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L24
	} else {
		goto L46
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v18 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L24
	} else {
		goto L42
	}
L5:
	;
	if v124 != 0 {
		goto L1
	} else {
		goto L41
	}
L6:
	;
	v124 = v2
	v128 = v2
	goto L5
L7:
	;
	goto L8
L8:
	;
	v22 = int32(0)
	v25 = v2
	v29 = v2
	goto L9
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v22<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v36 = int32(158425)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[677])))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v40 == int32(0) {
		v59 = v39
		v60 = v40
		goto L14
	} else {
		goto L15
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L24
	} else {
		goto L37
	}
L11:
	;
	goto L10
L12:
	;
	v99 = v22 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v99 < v100 {
		v22 = v99
		v25 = v96
		v29 = v97
		goto L9
	} else {
		goto L36
	}
L13:
	;
	if v60-v59 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	goto L13
L15:
	;
	if v39 != v40 {
		v59 = v39
		v60 = v40
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v44 = v35
	v45 = v36
	goto L17
L17:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v49 == int32(0) {
		v59 = v48
		v60 = v49
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v59 = v48
	v60 = v49
	goto L14
L19:
	;
	v52 = int32(1)
	if v48 == v49 {
		v44 = v44 + v52
		v45 = v45 + v52
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v64 = F_defGetString(m, v34)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v68 = int32(357308)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[678])))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v72 == int32(0) {
		v91 = v71
		v92 = v72
		goto L27
	} else {
		goto L28
	}
L24:
	;
	return int32(0)
L25:
	;
	v96 = v64
	v97 = v29
	goto L12
L26:
	;
	if v92-v91 != 0 {
		goto L11
	} else {
		goto L34
	}
L27:
	;
	goto L26
L28:
	;
	if v71 != v72 {
		v91 = v71
		v92 = v72
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v76 = v35
	v77 = v68
	goto L30
L30:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v81 == int32(0) {
		v91 = v80
		v92 = v81
		goto L27
	} else {
		goto L32
	}
L31:
	;
	v91 = v80
	v92 = v81
	goto L27
L32:
	;
	v84 = int32(1)
	if v80 == v81 {
		v76 = v76 + v84
		v77 = v77 + v84
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v94 = F_defGetBoolean(m, v34)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L24
	} else {
		goto L35
	}
L35:
	;
	v96 = v25
	v97 = v94
	goto L12
L36:
	;
	v124 = v96
	v128 = v97
	goto L5
L37:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v109
	F_errmsg(m, int32(753775), v11+int32(16))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L24
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(517280), int32(121), int32(106150))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L24
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	goto L4
L42:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L24
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(225939), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L24
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(517280), int32(127), int32(106150))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L24
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	v158 = F_tsearch_readline_begin(m, v11+int32(36), v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L24
	} else {
		goto L47
	}
L47:
	;
	if v158 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v161 = F_palloc0(m, int32(12))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L24
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L24
	} else {
		goto L87
	}
L51:
	;
	v165 = F_tsearch_readline(m, v11+int32(36))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L24
	} else {
		goto L53
	}
L52:
	;
	F_tsearch_readline_end(m, v11+int32(36))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L24
	} else {
		goto L85
	}
L53:
	;
	if v165 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v280 = int32(0)
	goto L52
L55:
	;
	goto L56
L56:
	;
	v171 = v165
	v176 = int32(0)
	goto L57
L57:
	;
	v182 = F_findwrd(m, v171, v11+int32(32), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L24
	} else {
		goto L60
	}
L58:
	;
	v280 = v267
	goto L52
L59:
	;
	F_pfree(m, v171)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L24
	} else {
		goto L82
	}
L60:
	;
	if v182 == int32(0) {
		v267 = v176
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v187 == int32(0) {
		v267 = v176
		goto L59
	} else {
		goto L62
	}
L62:
	;
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v190)
	v198 = F_findwrd(m, v186+int32(1), v11+int32(32), v11+int32(30))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L24
	} else {
		goto L63
	}
L63:
	;
	if v198 == int32(0) {
		v267 = v176
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v203 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v202))) = uint8(v203)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v205 <= v176 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if v205 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	goto L67
L67:
	;
	if v128&int32(1) != 0 {
		goto L75
	} else {
		goto L76
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+4)) = v222
	goto L67
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = int32(64)
	v212 = F_palloc(m, int32(1024))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L24
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = v205 << (uint(int32(1)) % 32)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v220 = F_repalloc(m, v217, v205<<(uint(int32(5))%32))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L24
	} else {
		goto L73
	}
L72:
	;
	v222 = v212
	goto L68
L73:
	;
	v222 = v220
	goto L68
L74:
	;
	v251 = v176 << (uint(int32(4)) % 32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v251+v252)+4)) = v249
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v257 = F_strlen(m, v198)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v255+v251)+8)) = v257
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+30)))
	*(*uint16)(unsafe.Add(mBase, uint32(v259+v251)+12)) = uint16(v261)
	v267 = v176 + int32(1)
	goto L59
L75:
	;
	v226 = F_pstrdup(m, v182)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L24
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v235 = F_strlen(m, v182)
	mBase = m.M
	v237 = F_str_tolower(m, v182, v235, int32(100))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L24
	} else {
		goto L80
	}
L78:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v228+v176<<(uint(int32(4))%32)))) = v226
	v233 = F_pstrdup(m, v198)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L24
	} else {
		goto L79
	}
L79:
	;
	v249 = v233
	goto L74
L80:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v239+v176<<(uint(int32(4))%32)))) = v237
	v244 = F_strlen(m, v198)
	mBase = m.M
	v246 = F_str_tolower(m, v198, v244, int32(100))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L24
	} else {
		goto L81
	}
L81:
	;
	v249 = v246
	goto L74
L82:
	;
	v273 = F_tsearch_readline(m, v11+int32(36))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L24
	} else {
		goto L83
	}
L83:
	;
	if v273 != 0 {
		v171 = v273
		v176 = v267
		goto L57
	} else {
		goto L84
	}
L84:
	;
	goto L58
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = v280
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	F_pg_qsort(m, v288, v280, int32(16), int32(1164))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L24
	} else {
		goto L86
	}
L86:
	;
	v294 = v128 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v161)+8)) = uint8(v294)
	m.G0 = v11 + int32(80)
	return v161
L87:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L24
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v156
	F_errmsg(m, int32(310507), v11)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L24
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(517280), int32(135), int32(106150))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L24
	} else {
		goto L90
	}
L90:
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
