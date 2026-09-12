package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetAttributeStorage(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v12 = l1
	v13 = int32(273705)
	goto L5
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L83
	} else {
		goto L92
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L83
	} else {
		goto L88
	}
L3:
	;
	m.G0 = v7 + int32(32)
	return v244
L4:
	;
	if v50 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v16 == v17 {
		v39 = v16
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v50 = int32(0)
	goto L4
L7:
	;
	v41 = int32(1)
	if v39 != 0 {
		v12 = v12 + v41
		v13 = v13 + v41
		goto L5
	} else {
		goto L16
	}
L8:
	;
	if base.Ui32((v16-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = v16 | int32(32)
	goto L11
L10:
	;
	v27 = v16
	goto L11
L11:
	;
	if base.Ui32((v17-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v36 = v17 | int32(32)
	goto L14
L13:
	;
	v36 = v17
	goto L14
L14:
	;
	if v27 == v36 {
		v39 = v27
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v50 = v27 - v36
	goto L4
L16:
	;
	goto L6
L17:
	;
	v244 = int32(112)
	goto L3
L18:
	;
	goto L19
L19:
	;
	v57 = l1
	v58 = int32(305248)
	goto L22
L20:
	;
	v240 = F_get_typstorage(m, l0)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L83
	} else {
		goto L86
	}
L21:
	;
	if v95 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v61 == v62 {
		v84 = v61
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v95 = int32(0)
	goto L21
L24:
	;
	v86 = int32(1)
	if v84 != 0 {
		v57 = v57 + v86
		v58 = v58 + v86
		goto L22
	} else {
		goto L33
	}
L25:
	;
	if base.Ui32((v61-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v72 = v61 | int32(32)
	goto L28
L27:
	;
	v72 = v61
	goto L28
L28:
	;
	if base.Ui32((v62-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v81 = v62 | int32(32)
	goto L31
L30:
	;
	v81 = v62
	goto L31
L31:
	;
	if v72 == v81 {
		v84 = v72
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v95 = v72 - v81
	goto L21
L33:
	;
	goto L23
L34:
	;
	v239 = int32(101)
	goto L20
L35:
	;
	goto L36
L36:
	;
	v102 = l1
	v103 = int32(453484)
	goto L38
L37:
	;
	if v140 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v106 == v107 {
		v129 = v106
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v140 = int32(0)
	goto L37
L40:
	;
	v131 = int32(1)
	if v129 != 0 {
		v102 = v102 + v131
		v103 = v103 + v131
		goto L38
	} else {
		goto L49
	}
L41:
	;
	if base.Ui32((v106-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v117 = v106 | int32(32)
	goto L44
L43:
	;
	v117 = v106
	goto L44
L44:
	;
	if base.Ui32((v107-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v126 = v107 | int32(32)
	goto L47
L46:
	;
	v126 = v107
	goto L47
L47:
	;
	if v117 == v126 {
		v129 = v117
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v140 = v117 - v126
	goto L37
L49:
	;
	goto L39
L50:
	;
	v239 = int32(120)
	goto L20
L51:
	;
	goto L52
L52:
	;
	v147 = l1
	v148 = int32(273700)
	goto L54
L53:
	;
	if v185 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v151 == v152 {
		v174 = v151
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v185 = int32(0)
	goto L53
L56:
	;
	v176 = int32(1)
	if v174 != 0 {
		v147 = v147 + v176
		v148 = v148 + v176
		goto L54
	} else {
		goto L65
	}
L57:
	;
	if base.Ui32((v151-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v162 = v151 | int32(32)
	goto L60
L59:
	;
	v162 = v151
	goto L60
L60:
	;
	if base.Ui32((v152-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v171 = v152 | int32(32)
	goto L63
L62:
	;
	v171 = v152
	goto L63
L63:
	;
	if v162 == v171 {
		v174 = v162
		goto L56
	} else {
		goto L64
	}
L64:
	;
	v185 = v162 - v171
	goto L53
L65:
	;
	goto L55
L66:
	;
	v239 = int32(109)
	goto L20
L67:
	;
	goto L68
L68:
	;
	v192 = l1
	v193 = int32(96979)
	goto L70
L69:
	;
	if v230 != 0 {
		goto L2
	} else {
		goto L82
	}
L70:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	if v196 == v197 {
		v219 = v196
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v230 = int32(0)
	goto L69
L72:
	;
	v221 = int32(1)
	if v219 != 0 {
		v192 = v192 + v221
		v193 = v193 + v221
		goto L70
	} else {
		goto L81
	}
L73:
	;
	if base.Ui32((v196-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v207 = v196 | int32(32)
	goto L76
L75:
	;
	v207 = v196
	goto L76
L76:
	;
	if base.Ui32((v197-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v216 = v197 | int32(32)
	goto L79
L78:
	;
	v216 = v197
	goto L79
L79:
	;
	if v207 == v216 {
		v219 = v207
		goto L72
	} else {
		goto L80
	}
L80:
	;
	v230 = v207 - v216
	goto L69
L81:
	;
	goto L71
L82:
	;
	v232 = F_get_typstorage(m, l0)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	return int32(0)
L84:
	;
	if v232 == int32(112) {
		v244 = int32(112)
		goto L3
	} else {
		goto L85
	}
L85:
	;
	v239 = v232
	goto L20
L86:
	;
	if v240 == int32(112) {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v244 = v239
	goto L3
L88:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L83
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
	F_errmsg(m, int32(688540), v7+int32(16))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L83
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(485619), int32(22100), int32(399808))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L83
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
	v274 = m.ExcPending
	if v274 != 0 {
		goto L83
	} else {
		goto L93
	}
L93:
	;
	v275 = F_format_type_be(m, l0)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L83
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v275
	F_errmsg(m, int32(521108), v7)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L83
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(485619), int32(22110), int32(399808))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L83
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
