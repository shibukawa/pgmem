package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckAttributeNamesTypes(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v13) < base.Ui32(int32(1601)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L34
	} else {
		goto L72
	}
L2:
	;
	v17 = l1 - int32(99)
	if base.B2i32(v17 == int32(0))|base.B2i32(v17 == int32(19)) != 0 {
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
	v234 = m.ExcPending
	if v234 != 0 {
		goto L34
	} else {
		goto L68
	}
L5:
	;
	m.G0 = v11 + int32(48)
	return
L6:
	;
	if int32(2) <= v13 {
		goto L40
	} else {
		goto L41
	}
L7:
	;
	if v13 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v30 = int32(0)
	goto L9
L9:
	;
	v41 = l0 + v13<<(uint(int32(4))%32) + v30*int32(100) + int32(24)
	v43 = F_strcmp(m, int32(_a_F_CheckAttributeNamesTypes_0), v41)
	mBase = m.M
	if v43 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L34
	} else {
		goto L35
	}
L11:
	;
	if v72 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L12:
	;
	v72 = int32(_a_F_CheckAttributeNamesTypes_1)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v48 = F_strcmp(m, int32(_a_F_CheckAttributeNamesTypes_2), v41)
	mBase = m.M
	if v48 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v72 = int32(_a_F_CheckAttributeNamesTypes_3)
	goto L11
L16:
	;
	goto L17
L17:
	;
	v53 = F_strcmp(m, int32(_a_F_CheckAttributeNamesTypes_4), v41)
	mBase = m.M
	if v53 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v72 = int32(_a_F_CheckAttributeNamesTypes_5)
	goto L11
L19:
	;
	goto L20
L20:
	;
	v58 = F_strcmp(m, int32(_a_F_CheckAttributeNamesTypes_6), v41)
	mBase = m.M
	if v58 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v72 = int32(_a_F_CheckAttributeNamesTypes_7)
	goto L11
L22:
	;
	goto L23
L23:
	;
	v63 = F_strcmp(m, int32(_a_F_CheckAttributeNamesTypes_8), v41)
	mBase = m.M
	if v63 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v72 = int32(_a_F_CheckAttributeNamesTypes_9)
	goto L11
L25:
	;
	goto L26
L26:
	;
	v70 = F_strcmp(m, int32(_a_F_CheckAttributeNamesTypes_10), v41)
	mBase = m.M
	if v70 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v71 = int32(0)
	goto L29
L28:
	;
	v71 = int32(_a_F_CheckAttributeNamesTypes_11)
	goto L29
L29:
	;
	v72 = v71
	goto L11
L30:
	;
	v76 = v30 + int32(1)
	if v13 != v76 {
		v30 = v76
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L10
L33:
	;
	goto L6
L34:
	;
	return
L35:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v41
	F_errmsg(m, int32(_a_F_CheckAttributeNamesTypes_12), v11+int32(32))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeNamesTypes_13), int32(482), int32(_a_F_CheckAttributeNamesTypes_14))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v183 = int32(1)
	if v13 <= v183 {
		goto L58
	} else {
		goto L59
	}
L40:
	;
	v110 = l0 + v13<<(uint(int32(4))%32) + int32(20)
	v116 = int32(1)
	goto L43
L41:
	;
	goto L42
L42:
	;
	if v13 == int32(0) {
		goto L5
	} else {
		goto L57
	}
L43:
	;
	v124 = v110 + v116*int32(100) + int32(4)
	v127 = int32(0)
	goto L45
L44:
	;
	goto L39
L45:
	;
	v138 = v110 + v127*int32(100) + int32(4)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if base.B2i32(v141 == int32(0))|base.B2i32(v141 != v144) != 0 {
		v162 = v141
		v163 = v144
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v171 = v116 + int32(1)
	if v171 != v13 {
		v116 = v171
		goto L43
	} else {
		goto L56
	}
L47:
	;
	if v162-v163 == int32(0) {
		goto L1
	} else {
		goto L54
	}
L48:
	;
	goto L47
L49:
	;
	v147 = v138
	v148 = v124
	goto L50
L50:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	if v152 == int32(0) {
		v162 = v152
		v163 = v151
		goto L48
	} else {
		goto L52
	}
L51:
	;
	v162 = v152
	v163 = v151
	goto L48
L52:
	;
	v155 = int32(1)
	if v152 == v151 {
		v147 = v147 + v155
		v148 = v148 + v155
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v168 = v127 + int32(1)
	if v168 != v116 {
		v127 = v168
		goto L45
	} else {
		goto L55
	}
L55:
	;
	goto L46
L56:
	;
	goto L44
L57:
	;
	goto L39
L58:
	;
	v186 = v183
	goto L60
L59:
	;
	v186 = v13
	goto L60
L60:
	;
	v192 = int32(0)
	goto L61
L61:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v202 = l0 + v196<<(uint(int32(4))%32) + v192*int32(100)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v202)+88))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202)+116))
	v207 = int32(0)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+110)))
	if v210 == int32(118) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L5
L63:
	;
	v213 = int32(8)
	goto L65
L64:
	;
	v213 = v207
	goto L65
L65:
	;
	F_CheckAttributeType(m, v202+int32(24), v205, v206, v207, v213|l2)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L34
	} else {
		goto L66
	}
L66:
	;
	v218 = v192 + int32(1)
	if v218 != v186 {
		v192 = v218
		goto L61
	} else {
		goto L67
	}
L67:
	;
	goto L62
L68:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L34
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1600)
	F_errmsg(m, int32(_a_F_CheckAttributeNamesTypes_15), v11)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L34
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeNamesTypes_13), int32(464), int32(_a_F_CheckAttributeNamesTypes_14))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L34
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errcode(m, int32(16806020))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L34
	} else {
		goto L73
	}
L73:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0 + v255<<(uint(int32(4))%32) + v127*int32(100) + int32(24)
	F_errmsg(m, int32(_a_F_CheckAttributeNamesTypes_16), v11+int32(16))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L34
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_CheckAttributeNamesTypes_13), int32(498), int32(_a_F_CheckAttributeNamesTypes_14))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L34
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
