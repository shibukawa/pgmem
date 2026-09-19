package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EndTransactionBlock(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var __phi16 int32
	_ = __phi16
	var v18 int32
	_ = v18
	var __phi18 int32
	_ = __phi18
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var __phi211 int32
	_ = __phi211
	var v213 int32
	_ = v213
	var __phi213 int32
	_ = __phi213
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	v1 = l0
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_EndTransactionBlock[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	switch v13 {
	case 0, 2, 6, 8, 9, 10, 11, 13, 14, 16, 17, 18, 19:
		goto L9
	case 1:
		goto L11
	case 3:
		goto L5
	case 4:
		goto L6
	case 5:
		goto L10
	case 7:
		goto L7
	case 12:
		goto L3
	case 15:
		goto L12
	default:
		v263 = v12
		v267 = v2
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L22
	} else {
		goto L89
	}
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v263)+77)) = uint8(v1)
	m.G0 = v9 + int32(112)
	return v267
L3:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if v209 != 0 {
		goto L74
	} else {
		goto L75
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L22
	} else {
		goto L69
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(6)
	v263 = v12
	v267 = int32(1)
	goto L2
L6:
	;
	if v1 != 0 {
		goto L4
	} else {
		goto L61
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(8)
	v263 = v12
	v267 = v2
	goto L2
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L22
	} else {
		goto L57
	}
L9:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L22
	} else {
		goto L50
	}
L10:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L22
	} else {
		goto L46
	}
L11:
	;
	if v1 != 0 {
		goto L8
	} else {
		goto L40
	}
L12:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if v14 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L22
	} else {
		goto L33
	}
L14:
	;
	__phi16 = v12
	__phi18 = v14
	v16 = __phi16
	v18 = __phi18
	goto L17
L15:
	;
	goto L16
L16:
	;
	v66 = v12 + int32(24)
	goto L13
L17:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	switch v22 - int32(12) {
	case 0:
		v51 = int32(17)
		goto L19
	default:
		goto L21
	case 3:
		goto L20
	}
L18:
	;
	v55 = v18 + int32(24)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	switch v56 - int32(3) {
	case 0:
		goto L32
	default:
		v66 = v55
		goto L13
	case 4:
		goto L31
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
	if v53 != 0 {
		__phi16 = v18
		__phi18 = v53
		v16 = __phi16
		v18 = __phi18
		goto L17
	} else {
		goto L30
	}
L20:
	;
	v51 = int32(16)
	goto L19
L21:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if base.Ui32(v31) <= base.Ui32(int32(19)) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v38
	F_errmsg_internal(m, int32(_a_F_EndTransactionBlock_0), v9-int32(-64))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L22
	} else {
		goto L28
	}
L25:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31<<(uint(int32(2))%32))+uint32(_c_F_EndTransactionBlock[1])))
	v38 = v36
	goto L27
L26:
	;
	v38 = int32(_a_F_EndTransactionBlock_1)
	goto L27
L27:
	;
	goto L24
L28:
	;
	F_errfinish(m, int32(_a_F_EndTransactionBlock_2), int32(_a_F_EndTransactionBlock_3), int32(_a_F_EndTransactionBlock_4))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	goto L18
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(8)
	v263 = v18
	v267 = v2
	goto L2
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(9)
	v263 = v18
	v267 = v2
	goto L2
L33:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if base.Ui32(v75) <= base.Ui32(int32(19)) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v82
	F_errmsg_internal(m, int32(_a_F_EndTransactionBlock_0), v9+int32(48))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L22
	} else {
		goto L38
	}
L35:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75<<(uint(int32(2))%32))+uint32(_c_F_EndTransactionBlock[1])))
	v82 = v80
	goto L37
L36:
	;
	v82 = int32(_a_F_EndTransactionBlock_1)
	goto L37
L37:
	;
	goto L34
L38:
	;
	F_errfinish(m, int32(_a_F_EndTransactionBlock_2), int32(_a_F_EndTransactionBlock_5), int32(_a_F_EndTransactionBlock_4))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L22
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v94 = int32(1)
	v97 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L22
	} else {
		goto L41
	}
L41:
	;
	if v97 == int32(0) {
		v263 = v12
		v267 = v94
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(16908610))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L22
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_EndTransactionBlock_6), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_EndTransactionBlock_2), int32(_a_F_EndTransactionBlock_7), int32(_a_F_EndTransactionBlock_4))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L22
	} else {
		goto L45
	}
L45:
	;
	v263 = v12
	v267 = v94
	goto L2
L46:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L22
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_EndTransactionBlock_8), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L22
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_EndTransactionBlock_2), int32(_a_F_EndTransactionBlock_9), int32(_a_F_EndTransactionBlock_4))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L22
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
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if base.Ui32(v133) <= base.Ui32(int32(19)) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v140
	F_errmsg_internal(m, int32(_a_F_EndTransactionBlock_0), v9+int32(96))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L22
	} else {
		goto L55
	}
L52:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v133<<(uint(int32(2))%32))+uint32(_c_F_EndTransactionBlock[1])))
	v140 = v138
	goto L54
L53:
	;
	v140 = int32(_a_F_EndTransactionBlock_1)
	goto L54
L54:
	;
	goto L51
L55:
	;
	F_errfinish(m, int32(_a_F_EndTransactionBlock_2), int32(_a_F_EndTransactionBlock_10), int32(_a_F_EndTransactionBlock_4))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L22
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(16908610))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L22
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = int32(_a_F_EndTransactionBlock_11)
	F_errmsg(m, int32(_a_F_EndTransactionBlock_12), v9+int32(80))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L22
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_EndTransactionBlock_2), int32(_a_F_EndTransactionBlock_13), int32(_a_F_EndTransactionBlock_4))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L22
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v175 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L22
	} else {
		goto L62
	}
L62:
	;
	if v175 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_errcode(m, int32(16908610))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L22
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	goto L5
L66:
	;
	F_errmsg(m, int32(_a_F_EndTransactionBlock_6), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L22
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_EndTransactionBlock_2), int32(4075), int32(_a_F_EndTransactionBlock_4))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L22
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	F_errcode(m, int32(16908610))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L22
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_EndTransactionBlock_11)
	F_errmsg(m, int32(_a_F_EndTransactionBlock_12), v9)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L22
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_EndTransactionBlock_2), int32(4071), int32(_a_F_EndTransactionBlock_4))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L22
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L22
	} else {
		goto L82
	}
L74:
	;
	__phi211 = v12
	__phi213 = v209
	v211 = __phi211
	v213 = __phi213
	goto L77
L75:
	;
	goto L76
L76:
	;
	v238 = v12 + int32(24)
	goto L73
L77:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	if v216 != int32(12) {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	v223 = v213 + int32(24)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v213)+24))
	if v224 != int32(3) {
		v238 = v223
		goto L73
	} else {
		goto L81
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+24)) = int32(14)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v213)+80))
	if v221 != 0 {
		__phi211 = v213
		__phi213 = v221
		v211 = __phi211
		v213 = __phi213
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = int32(6)
	v263 = v213
	v267 = int32(1)
	goto L2
L82:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	if base.Ui32(v243) <= base.Ui32(int32(19)) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v250
	F_errmsg_internal(m, int32(_a_F_EndTransactionBlock_0), v9+int32(16))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L22
	} else {
		goto L87
	}
L84:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v243<<(uint(int32(2))%32))+uint32(_c_F_EndTransactionBlock[1])))
	v250 = v248
	goto L86
L85:
	;
	v250 = int32(_a_F_EndTransactionBlock_1)
	goto L86
L86:
	;
	goto L83
L87:
	;
	F_errfinish(m, int32(_a_F_EndTransactionBlock_2), int32(_a_F_EndTransactionBlock_14), int32(_a_F_EndTransactionBlock_4))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L22
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	if base.Ui32(v277) <= base.Ui32(int32(19)) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v284
	F_errmsg_internal(m, int32(_a_F_EndTransactionBlock_0), v9+int32(32))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L22
	} else {
		goto L94
	}
L91:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v277<<(uint(int32(2))%32))+uint32(_c_F_EndTransactionBlock[1])))
	v284 = v282
	goto L93
L92:
	;
	v284 = int32(_a_F_EndTransactionBlock_1)
	goto L93
L93:
	;
	goto L90
L94:
	;
	F_errfinish(m, int32(_a_F_EndTransactionBlock_2), int32(_a_F_EndTransactionBlock_15), int32(_a_F_EndTransactionBlock_4))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L22
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SetTransactionIdLimit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[0]))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[1]))
	v19 = F_LWLockAcquire(m, v15+int32(384), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[2]))
		*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = l1
		v27 = l0 + int32(2147483647)
		if base.Ui32(v27) < base.Ui32(int32(3)) {
			v30 = l0 - int32(2147483646)
		} else {
			v30 = v27
		}
		*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v30
		v32 = l0 + v13
		v33 = int32(3)
		if base.Ui32(v32) < base.Ui32(v33) {
			v37 = v32 + v33
		} else {
			v37 = v32
		}
		*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v37
		*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l0
		v43 = v30 - int32(_a_F_SetTransactionIdLimit_0)
		if base.Ui32(v43) < base.Ui32(int32(3)) {
			v46 = v30 - int32(_a_F_SetTransactionIdLimit_1)
		} else {
			v46 = v43
		}
		*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v46
		v51 = v30 - int32(40000000)
		if base.Ui32(v51) < base.Ui32(int32(3)) {
			v54 = v30 - int32(40000003)
		} else {
			v54 = v51
		}
		*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v54
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
		v58 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[1]))
		F_LWLockRelease(m, v58+int32(384))
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			v65 = F_errstart(m, int32(14), int32(0))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				if v65 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v30
					F_errmsg_internal(m, int32(_a_F_SetTransactionIdLimit_2), v10+int32(32))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_SetTransactionIdLimit_3), int32(456), int32(_a_F_SetTransactionIdLimit_4))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return
						} else {
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v37))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v56)) == int32(0) {
								v90 = base.B2i32(base.Ui32(v37) <= base.Ui32(v56))
							} else {
								v90 = base.B2i32(int32(0) <= v56-v37)
							}
							if v90 == int32(0) {
							} else {
								v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[3])))
								if v94&int32(1) == int32(0) {
								} else {
									v100 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[4])))
									if v100&int32(1) != 0 {
									} else {
										v105 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[3])))
										if v105 == int32(1) {
											v109 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[5]))
											*(*int32)(unsafe.Add(mBase, uint32(v109+int32(16)))) = int32(1)
											v116 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[6]))
											v118 = F_pgmem_kill(m, v116, int32(10))
											mBase = m.M
										} else {
										}
									}
								}
							}
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v54))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v56)) == int32(0) {
								v130 = base.B2i32(base.Ui32(v54) <= base.Ui32(v56))
							} else {
								v130 = base.B2i32(int32(0) <= v56-v54)
							}
							if v130 == int32(0) {
								m.G0 = v10 + int32(48)
								return
							} else {
								v134 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[4])))
								if v134&int32(1) != 0 {
									m.G0 = v10 + int32(48)
									return
								} else {
									v138 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[7]))
									v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
									if base.B2i32(v139 == int32(2)) == int32(0) {
										v166 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return
										} else {
											if v166 == int32(0) {
												m.G0 = v10 + int32(48)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v30 - v56
												F_errmsg(m, int32(_a_F_SetTransactionIdLimit_5), v10)
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return
												} else {
													v178 = int32(501)
													F_errhint(m, int32(_a_F_SetTransactionIdLimit_6), int32(0))
													mBase = m.M
													v182 = m.ExcPending
													if v182 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_SetTransactionIdLimit_3), v178, int32(_a_F_SetTransactionIdLimit_4))
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return
														} else {
															m.G0 = v10 + int32(48)
															return
														}
													}
												}
											}
										}
									} else {
										v144 = F_get_database_name(m, l1)
										mBase = m.M
										v145 = m.ExcPending
										if v145 != 0 {
											return
										} else {
											if v144 == int32(0) {
												v166 = F_errstart(m, int32(19), int32(0))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return
												} else {
													if v166 == int32(0) {
														m.G0 = v10 + int32(48)
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
														*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v30 - v56
														F_errmsg(m, int32(_a_F_SetTransactionIdLimit_5), v10)
														mBase = m.M
														v175 = m.ExcPending
														if v175 != 0 {
															return
														} else {
															v178 = int32(501)
															F_errhint(m, int32(_a_F_SetTransactionIdLimit_6), int32(0))
															mBase = m.M
															v182 = m.ExcPending
															if v182 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_SetTransactionIdLimit_3), v178, int32(_a_F_SetTransactionIdLimit_4))
																mBase = m.M
																v186 = m.ExcPending
																if v186 != 0 {
																	return
																} else {
																	m.G0 = v10 + int32(48)
																	return
																}
															}
														}
													}
												}
											} else {
												v150 = F_errstart(m, int32(19), int32(0))
												mBase = m.M
												v151 = m.ExcPending
												if v151 != 0 {
													return
												} else {
													if v150 == int32(0) {
														m.G0 = v10 + int32(48)
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v144
														*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v30 - v56
														F_errmsg(m, int32(_a_F_SetTransactionIdLimit_7), v10+int32(16))
														mBase = m.M
														v161 = m.ExcPending
														if v161 != 0 {
															return
														} else {
															v178 = int32(494)
															F_errhint(m, int32(_a_F_SetTransactionIdLimit_6), int32(0))
															mBase = m.M
															v182 = m.ExcPending
															if v182 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_SetTransactionIdLimit_3), v178, int32(_a_F_SetTransactionIdLimit_4))
																mBase = m.M
																v186 = m.ExcPending
																if v186 != 0 {
																	return
																} else {
																	m.G0 = v10 + int32(48)
																	return
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
				} else {
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v37))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v56)) == int32(0) {
						v90 = base.B2i32(base.Ui32(v37) <= base.Ui32(v56))
					} else {
						v90 = base.B2i32(int32(0) <= v56-v37)
					}
					if v90 == int32(0) {
					} else {
						v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[3])))
						if v94&int32(1) == int32(0) {
						} else {
							v100 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[4])))
							if v100&int32(1) != 0 {
							} else {
								v105 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[3])))
								if v105 == int32(1) {
									v109 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[5]))
									*(*int32)(unsafe.Add(mBase, uint32(v109+int32(16)))) = int32(1)
									v116 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[6]))
									v118 = F_pgmem_kill(m, v116, int32(10))
									mBase = m.M
								} else {
								}
							}
						}
					}
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v54))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v56)) == int32(0) {
						v130 = base.B2i32(base.Ui32(v54) <= base.Ui32(v56))
					} else {
						v130 = base.B2i32(int32(0) <= v56-v54)
					}
					if v130 == int32(0) {
						m.G0 = v10 + int32(48)
						return
					} else {
						v134 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[4])))
						if v134&int32(1) != 0 {
							m.G0 = v10 + int32(48)
							return
						} else {
							v138 = *(*int32)(unsafe.Add(mBase, _c_F_SetTransactionIdLimit[7]))
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+20))
							if base.B2i32(v139 == int32(2)) == int32(0) {
								v166 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v167 = m.ExcPending
								if v167 != 0 {
									return
								} else {
									if v166 == int32(0) {
										m.G0 = v10 + int32(48)
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v30 - v56
										F_errmsg(m, int32(_a_F_SetTransactionIdLimit_5), v10)
										mBase = m.M
										v175 = m.ExcPending
										if v175 != 0 {
											return
										} else {
											v178 = int32(501)
											F_errhint(m, int32(_a_F_SetTransactionIdLimit_6), int32(0))
											mBase = m.M
											v182 = m.ExcPending
											if v182 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_SetTransactionIdLimit_3), v178, int32(_a_F_SetTransactionIdLimit_4))
												mBase = m.M
												v186 = m.ExcPending
												if v186 != 0 {
													return
												} else {
													m.G0 = v10 + int32(48)
													return
												}
											}
										}
									}
								}
							} else {
								v144 = F_get_database_name(m, l1)
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									if v144 == int32(0) {
										v166 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return
										} else {
											if v166 == int32(0) {
												m.G0 = v10 + int32(48)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v30 - v56
												F_errmsg(m, int32(_a_F_SetTransactionIdLimit_5), v10)
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return
												} else {
													v178 = int32(501)
													F_errhint(m, int32(_a_F_SetTransactionIdLimit_6), int32(0))
													mBase = m.M
													v182 = m.ExcPending
													if v182 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_SetTransactionIdLimit_3), v178, int32(_a_F_SetTransactionIdLimit_4))
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return
														} else {
															m.G0 = v10 + int32(48)
															return
														}
													}
												}
											}
										}
									} else {
										v150 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return
										} else {
											if v150 == int32(0) {
												m.G0 = v10 + int32(48)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v144
												*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v30 - v56
												F_errmsg(m, int32(_a_F_SetTransactionIdLimit_7), v10+int32(16))
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return
												} else {
													v178 = int32(494)
													F_errhint(m, int32(_a_F_SetTransactionIdLimit_6), int32(0))
													mBase = m.M
													v182 = m.ExcPending
													if v182 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_SetTransactionIdLimit_3), v178, int32(_a_F_SetTransactionIdLimit_4))
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return
														} else {
															m.G0 = v10 + int32(48)
															return
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
func F_ShowTransactionStateRec(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v14 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v58 = v12 + int32(96)
	F_initStringInfo(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L13
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_ShowTransactionStateRec[0]))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ShowTransactionStateRec[1]))
	v23 = m.G0
	v26 = v22 - (v23 - int32(1))
	v28 = v26 >> (uint(int32(31)) % 32)
	goto L3
L3:
	;
	if base.B2i32(v20 < v26^v28-v28)&base.B2i32(v22 != int32(0)) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v37 = F_errstart(m, int32(10), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	F_ShowTransactionStateRec(m, l0, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L7
	} else {
		goto L12
	}
L7:
	;
	return
L8:
	;
	if v37 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = l0
	F_errmsg_internal(m, int32(_a_F_ShowTransactionStateRec_0), v12+int32(80))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_ShowTransactionStateRec_1), int32(_a_F_ShowTransactionStateRec_2), int32(_a_F_ShowTransactionStateRec_3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L1
L12:
	;
	goto L1
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v62 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v113 = F_errstart(m, int32(10), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L22
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v66
	F_appendStringInfo(m, v58, int32(_a_F_ShowTransactionStateRec_4), v12-int32(-64))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v73 < int32(2) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v79 = int32(1)
	goto L18
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v79<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v89
	F_appendStringInfo(m, v12+int32(96), int32(_a_F_ShowTransactionStateRec_5), v12+int32(48))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L20
	}
L19:
	;
	goto L14
L20:
	;
	v99 = v79 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v99 < v100 {
		v79 = v99
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if v113 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v116 = int32(_a_F_ShowTransactionStateRec_6)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if base.Ui32(v118) <= base.Ui32(int32(19)) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
	F_pfree(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L40
	}
L26:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118<<(uint(int32(2))%32))+uint32(_c_F_ShowTransactionStateRec[2])))
	v124 = v123
	goto L28
L27:
	;
	v124 = v116
	goto L28
L28:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if base.Ui32(v126) <= base.Ui32(int32(5)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v126<<(uint(int32(2))%32))+uint32(_c_F_ShowTransactionStateRec[3])))
	v132 = v131
	goto L31
L30:
	;
	v132 = v116
	goto L31
L31:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v134 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v132
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+20)) = uint32(v134)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v133
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_ShowTransactionStateRec[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v141
	v146 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ShowTransactionStateRec[5])))
	if v146 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v147 = int32(_a_F_ShowTransactionStateRec_7)
	goto L34
L33:
	;
	v147 = int32(_a_F_ShowTransactionStateRec_8)
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v125
	if v115 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v152 = v115
	goto L37
L36:
	;
	v152 = int32(_a_F_ShowTransactionStateRec_9)
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v124
	F_errmsg_internal(m, int32(_a_F_ShowTransactionStateRec_10), v12)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_ShowTransactionStateRec_1), int32(_a_F_ShowTransactionStateRec_11), int32(_a_F_ShowTransactionStateRec_3))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L25
L40:
	;
	m.G0 = v12 + int32(112)
	return
}
func F_TransactionIdLatest(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v7 = l1 - int32(1)
	if v7 < int32(0) {
		v75 = l0
	} else {
		if l1&int32(1) != 0 {
			v12 = int32(2)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l2+v7<<(uint(v12)%32))))
			if base.B2i32(base.Ui32(v12) < base.Ui32(v15))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
				if base.Ui32(l0) < base.Ui32(v15) {
					v27 = v15
				} else {
					v27 = l0
				}
			} else {
				if int32(0) <= l0-v15 {
					v27 = l0
				} else {
					v27 = v15
				}
			}
			v30 = v27
			v32 = l1 - int32(2)
		} else {
			v30 = l0
			v32 = v7
		}
		if v7 == int32(0) {
			v75 = v30
		} else {
			v35 = v30
			v36 = v32
			for {
				v40 = int32(3)
				v44 = l2 + v36<<(uint(int32(2))%32)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
				if base.B2i32(base.Ui32(v35) < base.Ui32(v40))|base.B2i32(base.Ui32(v45) < base.Ui32(v40)) == int32(0) {
					if v35-v45 < int32(0) {
						v55 = v45
					} else {
						v55 = v35
					}
				} else {
					if base.Ui32(v45) <= base.Ui32(v35) {
						v55 = v35
					} else {
						v55 = v45
					}
				}
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v44-int32(4))))
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v58))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v55)) == int32(0) {
					if base.Ui32(v55) < base.Ui32(v58) {
						v70 = v58
					} else {
						v70 = v55
					}
				} else {
					if int32(0) <= v55-v58 {
						v70 = v55
					} else {
						v70 = v58
					}
				}
				if int32(1) < v36 {
					v35 = v70
					v36 = v36 - int32(2)
					continue
				} else {
					break
				}
				break
			}
			v75 = v70
		}
	}
	return v75
}
func F_TransactionIdPrecedes(m *base.Module, l0 int32, l1 int32) int32 {
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		return base.B2i32(base.Ui32(l0) < base.Ui32(l1))
	} else {
		return int32(base.Ui32(l0-l1) >> (uint(int32(31)) % 32))
	}
}
func F_TransactionTreeSetCommitTsData(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	v5 = l4
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionTreeSetCommitTsData[0]))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
	if v19 == int32(1) {
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
	return
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2+l1<<(uint(int32(2))%32)-int32(4))))
	v30 = v29
	goto L6
L5:
	;
	v30 = l0
	goto L6
L6:
	;
	v36 = l0
	v39 = int32(0)
	goto L7
L7:
	;
	v48 = base.I32_div_u_s(v36, int32(819))
	if l1 <= v39 {
		v84 = v39
		v89 = int32(0)
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionTreeSetCommitTsData[1]))
	v202 = F_LWLockAcquire(m, v198+int32(_a_F_TransactionTreeSetCommitTsData_0), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L15
	} else {
		goto L28
	}
L9:
	;
	v94 = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionTreeSetCommitTsData[2]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	v99 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionTreeSetCommitTsData[3])))
	v100 = base.I32_rem_u_s(v48, v99)
	v103 = v97 + v100<<(uint(int32(7))%32)
	v105 = F_LWLockAcquire(m, v103, v94)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v58 = v39
	goto L11
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2+v58<<(uint(int32(2))%32))))
	v73 = base.I32_div_u_s(v71, int32(819))
	v74 = base.B2i32(v73 != v48)
	if v73 != v48 {
		v84 = v58
		v89 = v74
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v84 = l1
	v89 = v74
	goto L9
L13:
	;
	v76 = v58 + int32(1)
	if v76 != l1 {
		v58 = v76
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	return
L16:
	;
	v109 = F_SimpleLruReadPage(m, int32(_a_F_TransactionTreeSetCommitTsData_1), base.I64_extend_i32_u(v48), int32(1), v36)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v112 = v109 << (uint(int32(2)) % 32)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionTreeSetCommitTsData[2]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112+v115)))
	v123 = v117 + (v36-v48*int32(819))*int32(10)
	*(*uint16)(unsafe.Add(mBase, uint32(v123)+8)) = uint16(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v123))) = l3
	v126 = v84 - v39
	if int32(0) < v126 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v142 = v94
	goto L21
L19:
	;
	goto L20
L20:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionTreeSetCommitTsData[2]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	v187 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v185+v109))) = uint8(v187)
	F_LWLockRelease(m, v103)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L15
	} else {
		goto L24
	}
L21:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionTreeSetCommitTsData[2]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v150+v112)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l2+v39<<(uint(int32(2))%32)+v142<<(uint(int32(2))%32))))
	v158 = base.I32_rem_u_s(v156, int32(819))
	v161 = v152 + v158*int32(10)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+8)) = uint16(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v161))) = l3
	v165 = v142 + int32(1)
	if v165 != v126 {
		v142 = v165
		goto L21
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	goto L22
L24:
	;
	if v89 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l2+v84<<(uint(int32(2))%32))))
	v36 = v196
	v39 = v84 + int32(1)
	goto L7
L26:
	;
	goto L27
L27:
	;
	goto L8
L28:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionTreeSetCommitTsData[0]))
	*(*uint16)(unsafe.Add(mBase, uint32(v205)+16)) = uint16(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v205)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = l0
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionTreeSetCommitTsData[4]))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+44))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v30))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v211)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v223 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v223 = base.B2i32(base.Ui32(v211) < base.Ui32(v30))
	goto L29
L31:
	;
	goto L32
L32:
	;
	v223 = int32(base.Ui32(v211-v30) >> (uint(int32(31)) % 32))
	goto L29
L33:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionTreeSetCommitTsData[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+44)) = v30
	goto L35
L34:
	;
	goto L35
L35:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionTreeSetCommitTsData[1]))
	F_LWLockRelease(m, v228+int32(_a_F_TransactionTreeSetCommitTsData_0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L15
	} else {
		goto L36
	}
L36:
	;
	goto L3
}
