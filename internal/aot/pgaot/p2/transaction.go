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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var __phi222 int32
	_ = __phi222
	var v224 int32
	_ = v224
	var __phi224 int32
	_ = __phi224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	v1 = l0
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	switch v13 {
	case 0, 2, 6, 8, 9, 10, 11, 13, 14, 16, 17, 18, 19:
		goto L10
	case 1:
		goto L12
	case 3:
		goto L9
	case 4:
		goto L6
	case 5:
		goto L11
	case 7:
		goto L7
	case 12:
		goto L3
	case 15:
		goto L13
	default:
		v277 = v12
		v281 = v2
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L23
	} else {
		goto L90
	}
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v277)+77)) = uint8(v1)
	m.G0 = v9 + int32(112)
	return v281
L3:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if v220 != 0 {
		goto L75
	} else {
		goto L76
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L23
	} else {
		goto L70
	}
L5:
	;
	v277 = v12
	v281 = int32(1)
	goto L2
L6:
	;
	if v1 != 0 {
		goto L4
	} else {
		goto L62
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(8)
	v277 = v12
	v281 = v2
	goto L2
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L23
	} else {
		goto L58
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(6)
	goto L5
L10:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L23
	} else {
		goto L51
	}
L11:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L23
	} else {
		goto L47
	}
L12:
	;
	if v1 != 0 {
		goto L8
	} else {
		goto L41
	}
L13:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if v14 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L34
	}
L15:
	;
	__phi16 = v12
	__phi18 = v14
	v16 = __phi16
	v18 = __phi18
	goto L18
L16:
	;
	goto L17
L17:
	;
	v69 = v12 + int32(24)
	goto L14
L18:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	switch v22 - int32(12) {
	case 0:
		v54 = int32(17)
		goto L20
	default:
		goto L22
	case 3:
		goto L21
	}
L19:
	;
	v58 = v18 + int32(24)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	switch v59 - int32(3) {
	case 0:
		goto L33
	default:
		v69 = v58
		goto L14
	case 4:
		goto L32
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+80))
	if v56 != 0 {
		__phi16 = v18
		__phi18 = v56
		v16 = __phi16
		v18 = __phi18
		goto L18
	} else {
		goto L31
	}
L21:
	;
	v54 = int32(16)
	goto L20
L22:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if base.Ui32(v31) <= base.Ui32(int32(19)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v41
	F_errmsg_internal(m, int32(198383), v9-int32(-64))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L23
	} else {
		goto L29
	}
L26:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31<<(uint(int32(2))%32))+uint32(_consts[164])))
	v41 = v40
	goto L28
L27:
	;
	v41 = int32(571067)
	goto L28
L28:
	;
	goto L25
L29:
	;
	F_errfinish(m, int32(518151), int32(4124), int32(333339))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	goto L19
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(8)
	v277 = v18
	v281 = v2
	goto L2
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = int32(9)
	v277 = v18
	v281 = v2
	goto L2
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if base.Ui32(v78) <= base.Ui32(int32(19)) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v88
	F_errmsg_internal(m, int32(198383), v9+int32(48))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L23
	} else {
		goto L39
	}
L36:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v78<<(uint(int32(2))%32))+uint32(_consts[164])))
	v88 = v87
	goto L38
L37:
	;
	v88 = int32(571067)
	goto L38
L38:
	;
	goto L35
L39:
	;
	F_errfinish(m, int32(518151), int32(4133), int32(333339))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L23
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
	v100 = int32(1)
	v103 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L23
	} else {
		goto L42
	}
L42:
	;
	if v103 == int32(0) {
		v277 = v12
		v281 = v100
		goto L2
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(16908610))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L23
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(136652), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(518151), int32(4154), int32(333339))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	v277 = v12
	v281 = v100
	goto L2
L47:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L23
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(273351), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L23
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(518151), int32(4165), int32(333339))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L23
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if base.Ui32(v139) <= base.Ui32(int32(19)) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v149
	F_errmsg_internal(m, int32(198383), v9+int32(96))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L23
	} else {
		goto L56
	}
L53:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v139<<(uint(int32(2))%32))+uint32(_consts[164])))
	v149 = v148
	goto L55
L54:
	;
	v149 = int32(571067)
	goto L55
L55:
	;
	goto L52
L56:
	;
	F_errfinish(m, int32(518151), int32(4183), int32(333339))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L23
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(16908610))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = int32(557587)
	F_errmsg(m, int32(163983), v9+int32(80))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(518151), int32(4150), int32(333339))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L23
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v186 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L23
	} else {
		goto L63
	}
L63:
	;
	if v186 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_errcode(m, int32(16908610))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L23
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(6)
	goto L5
L67:
	;
	F_errmsg(m, int32(136652), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L23
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(518151), int32(4075), int32(333339))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L23
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	F_errcode(m, int32(16908610))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L23
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(557587)
	F_errmsg(m, int32(163983), v9)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L23
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(518151), int32(4071), int32(333339))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L23
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L23
	} else {
		goto L83
	}
L75:
	;
	__phi222 = v12
	__phi224 = v220
	v222 = __phi222
	v224 = __phi224
	goto L78
L76:
	;
	goto L77
L77:
	;
	v249 = v12 + int32(24)
	goto L74
L78:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v222)+24))
	if v227 != int32(12) {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v234 = v224 + int32(24)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
	if v235 != int32(3) {
		v249 = v234
		goto L74
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+24)) = int32(14)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v224)+80))
	if v232 != 0 {
		__phi222 = v224
		__phi224 = v232
		v222 = __phi222
		v224 = __phi224
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = int32(6)
	v277 = v224
	v281 = int32(1)
	goto L2
L83:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	if base.Ui32(v254) <= base.Ui32(int32(19)) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v264
	F_errmsg_internal(m, int32(198383), v9+int32(16))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L23
	} else {
		goto L88
	}
L85:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v254<<(uint(int32(2))%32))+uint32(_consts[164])))
	v264 = v263
	goto L87
L86:
	;
	v264 = int32(571067)
	goto L87
L87:
	;
	goto L84
L88:
	;
	F_errfinish(m, int32(518151), int32(4106), int32(333339))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L23
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
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v222)+24))
	if base.Ui32(v291) <= base.Ui32(int32(19)) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v301
	F_errmsg_internal(m, int32(198383), v9+int32(32))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L23
	} else {
		goto L95
	}
L92:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v291<<(uint(int32(2))%32))+uint32(_consts[164])))
	v301 = v300
	goto L94
L93:
	;
	v301 = int32(571067)
	goto L94
L94:
	;
	goto L91
L95:
	;
	F_errfinish(m, int32(518151), int32(4099), int32(333339))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L23
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SetTransactionIdLimit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[160]))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v18 = F_LWLockAcquire(m, v14+int32(384), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, _consts[142]))
		*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = l1
		v26 = l0 + int32(2147483647)
		if base.Ui32(v26) < base.Ui32(int32(3)) {
			v29 = l0 - int32(2147483646)
		} else {
			v29 = v26
		}
		*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v29
		v31 = l0 + v12
		v32 = int32(3)
		if base.Ui32(v31) < base.Ui32(v32) {
			v36 = v31 + v32
		} else {
			v36 = v31
		}
		*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v36
		*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l0
		v42 = v29 - int32(3000000)
		if base.Ui32(v42) < base.Ui32(int32(3)) {
			v45 = v29 - int32(3000003)
		} else {
			v45 = v42
		}
		*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v45
		v50 = v29 - int32(40000000)
		if base.Ui32(v50) < base.Ui32(int32(3)) {
			v53 = v29 - int32(40000003)
		} else {
			v53 = v50
		}
		*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v53
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
		v57 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		F_LWLockRelease(m, v57+int32(384))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			v64 = F_errstart(m, int32(14), int32(0))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				if v64 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v29
					F_errmsg_internal(m, int32(61545), v9+int32(32))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						F_errfinish(m, int32(521008), int32(456), int32(108140))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v36))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v55)) == int32(0) {
								v89 = base.B2i32(base.Ui32(v36) <= base.Ui32(v55))
							} else {
								v89 = base.B2i32(int32(0) <= v55-v36)
							}
							if v89 == int32(0) {
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v53))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v55)) == int32(0) {
									v112 = base.B2i32(base.Ui32(v53) <= base.Ui32(v55))
								} else {
									v112 = base.B2i32(int32(0) <= v55-v53)
								}
								if v112 == int32(0) {
									m.G0 = v9 + int32(48)
									return
								} else {
									v116 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
									if v116 != 0 {
										m.G0 = v9 + int32(48)
										return
									} else {
										v118 = *(*int32)(unsafe.Add(mBase, _consts[65]))
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
										if base.B2i32(v119 == int32(2)) == int32(0) {
											v146 = F_errstart(m, int32(19), int32(0))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return
											} else {
												if v146 == int32(0) {
													m.G0 = v9 + int32(48)
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
													F_errmsg(m, int32(151086), v9)
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return
													} else {
														v158 = int32(501)
														F_errhint(m, int32(613991), int32(0))
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return
														} else {
															F_errfinish(m, int32(521008), v158, int32(108140))
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return
															} else {
																m.G0 = v9 + int32(48)
																return
															}
														}
													}
												}
											}
										} else {
											v124 = F_get_database_name(m, l1)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return
											} else {
												if v124 == int32(0) {
													v146 = F_errstart(m, int32(19), int32(0))
													mBase = m.M
													v147 = m.ExcPending
													if v147 != 0 {
														return
													} else {
														if v146 == int32(0) {
															m.G0 = v9 + int32(48)
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
															F_errmsg(m, int32(151086), v9)
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return
															} else {
																v158 = int32(501)
																F_errhint(m, int32(613991), int32(0))
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(521008), v158, int32(108140))
																	mBase = m.M
																	v166 = m.ExcPending
																	if v166 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(48)
																		return
																	}
																}
															}
														}
													}
												} else {
													v130 = F_errstart(m, int32(19), int32(0))
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return
													} else {
														if v130 == int32(0) {
															m.G0 = v9 + int32(48)
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v124
															*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v29 - v55
															F_errmsg(m, int32(151147), v9+int32(16))
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return
															} else {
																v158 = int32(494)
																F_errhint(m, int32(613991), int32(0))
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(521008), v158, int32(108140))
																	mBase = m.M
																	v166 = m.ExcPending
																	if v166 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(48)
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
							} else {
								v93 = int32(*(*uint8)(unsafe.Add(mBase, _consts[96])))
								if v93 != int32(1) {
									if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v53))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v55)) == int32(0) {
										v112 = base.B2i32(base.Ui32(v53) <= base.Ui32(v55))
									} else {
										v112 = base.B2i32(int32(0) <= v55-v53)
									}
									if v112 == int32(0) {
										m.G0 = v9 + int32(48)
										return
									} else {
										v116 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
										if v116 != 0 {
											m.G0 = v9 + int32(48)
											return
										} else {
											v118 = *(*int32)(unsafe.Add(mBase, _consts[65]))
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
											if base.B2i32(v119 == int32(2)) == int32(0) {
												v146 = F_errstart(m, int32(19), int32(0))
												mBase = m.M
												v147 = m.ExcPending
												if v147 != 0 {
													return
												} else {
													if v146 == int32(0) {
														m.G0 = v9 + int32(48)
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
														*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
														F_errmsg(m, int32(151086), v9)
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return
														} else {
															v158 = int32(501)
															F_errhint(m, int32(613991), int32(0))
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return
															} else {
																F_errfinish(m, int32(521008), v158, int32(108140))
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(48)
																	return
																}
															}
														}
													}
												}
											} else {
												v124 = F_get_database_name(m, l1)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return
												} else {
													if v124 == int32(0) {
														v146 = F_errstart(m, int32(19), int32(0))
														mBase = m.M
														v147 = m.ExcPending
														if v147 != 0 {
															return
														} else {
															if v146 == int32(0) {
																m.G0 = v9 + int32(48)
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
																*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
																F_errmsg(m, int32(151086), v9)
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return
																} else {
																	v158 = int32(501)
																	F_errhint(m, int32(613991), int32(0))
																	mBase = m.M
																	v162 = m.ExcPending
																	if v162 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(521008), v158, int32(108140))
																		mBase = m.M
																		v166 = m.ExcPending
																		if v166 != 0 {
																			return
																		} else {
																			m.G0 = v9 + int32(48)
																			return
																		}
																	}
																}
															}
														}
													} else {
														v130 = F_errstart(m, int32(19), int32(0))
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return
														} else {
															if v130 == int32(0) {
																m.G0 = v9 + int32(48)
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v124
																*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v29 - v55
																F_errmsg(m, int32(151147), v9+int32(16))
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return
																} else {
																	v158 = int32(494)
																	F_errhint(m, int32(613991), int32(0))
																	mBase = m.M
																	v162 = m.ExcPending
																	if v162 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(521008), v158, int32(108140))
																		mBase = m.M
																		v166 = m.ExcPending
																		if v166 != 0 {
																			return
																		} else {
																			m.G0 = v9 + int32(48)
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
								} else {
									v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
									if v97 != 0 {
										if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v53))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v55)) == int32(0) {
											v112 = base.B2i32(base.Ui32(v53) <= base.Ui32(v55))
										} else {
											v112 = base.B2i32(int32(0) <= v55-v53)
										}
										if v112 == int32(0) {
											m.G0 = v9 + int32(48)
											return
										} else {
											v116 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
											if v116 != 0 {
												m.G0 = v9 + int32(48)
												return
											} else {
												v118 = *(*int32)(unsafe.Add(mBase, _consts[65]))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
												if base.B2i32(v119 == int32(2)) == int32(0) {
													v146 = F_errstart(m, int32(19), int32(0))
													mBase = m.M
													v147 = m.ExcPending
													if v147 != 0 {
														return
													} else {
														if v146 == int32(0) {
															m.G0 = v9 + int32(48)
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
															F_errmsg(m, int32(151086), v9)
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return
															} else {
																v158 = int32(501)
																F_errhint(m, int32(613991), int32(0))
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(521008), v158, int32(108140))
																	mBase = m.M
																	v166 = m.ExcPending
																	if v166 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(48)
																		return
																	}
																}
															}
														}
													}
												} else {
													v124 = F_get_database_name(m, l1)
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return
													} else {
														if v124 == int32(0) {
															v146 = F_errstart(m, int32(19), int32(0))
															mBase = m.M
															v147 = m.ExcPending
															if v147 != 0 {
																return
															} else {
																if v146 == int32(0) {
																	m.G0 = v9 + int32(48)
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
																	F_errmsg(m, int32(151086), v9)
																	mBase = m.M
																	v155 = m.ExcPending
																	if v155 != 0 {
																		return
																	} else {
																		v158 = int32(501)
																		F_errhint(m, int32(613991), int32(0))
																		mBase = m.M
																		v162 = m.ExcPending
																		if v162 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(521008), v158, int32(108140))
																			mBase = m.M
																			v166 = m.ExcPending
																			if v166 != 0 {
																				return
																			} else {
																				m.G0 = v9 + int32(48)
																				return
																			}
																		}
																	}
																}
															}
														} else {
															v130 = F_errstart(m, int32(19), int32(0))
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return
															} else {
																if v130 == int32(0) {
																	m.G0 = v9 + int32(48)
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v124
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v29 - v55
																	F_errmsg(m, int32(151147), v9+int32(16))
																	mBase = m.M
																	v141 = m.ExcPending
																	if v141 != 0 {
																		return
																	} else {
																		v158 = int32(494)
																		F_errhint(m, int32(613991), int32(0))
																		mBase = m.M
																		v162 = m.ExcPending
																		if v162 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(521008), v158, int32(108140))
																			mBase = m.M
																			v166 = m.ExcPending
																			if v166 != 0 {
																				return
																			} else {
																				m.G0 = v9 + int32(48)
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
									} else {
										F_SendPostmasterSignal(m, int32(4))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return
										} else {
											if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v53))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v55)) == int32(0) {
												v112 = base.B2i32(base.Ui32(v53) <= base.Ui32(v55))
											} else {
												v112 = base.B2i32(int32(0) <= v55-v53)
											}
											if v112 == int32(0) {
												m.G0 = v9 + int32(48)
												return
											} else {
												v116 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
												if v116 != 0 {
													m.G0 = v9 + int32(48)
													return
												} else {
													v118 = *(*int32)(unsafe.Add(mBase, _consts[65]))
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
													if base.B2i32(v119 == int32(2)) == int32(0) {
														v146 = F_errstart(m, int32(19), int32(0))
														mBase = m.M
														v147 = m.ExcPending
														if v147 != 0 {
															return
														} else {
															if v146 == int32(0) {
																m.G0 = v9 + int32(48)
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
																*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
																F_errmsg(m, int32(151086), v9)
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return
																} else {
																	v158 = int32(501)
																	F_errhint(m, int32(613991), int32(0))
																	mBase = m.M
																	v162 = m.ExcPending
																	if v162 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(521008), v158, int32(108140))
																		mBase = m.M
																		v166 = m.ExcPending
																		if v166 != 0 {
																			return
																		} else {
																			m.G0 = v9 + int32(48)
																			return
																		}
																	}
																}
															}
														}
													} else {
														v124 = F_get_database_name(m, l1)
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return
														} else {
															if v124 == int32(0) {
																v146 = F_errstart(m, int32(19), int32(0))
																mBase = m.M
																v147 = m.ExcPending
																if v147 != 0 {
																	return
																} else {
																	if v146 == int32(0) {
																		m.G0 = v9 + int32(48)
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
																		F_errmsg(m, int32(151086), v9)
																		mBase = m.M
																		v155 = m.ExcPending
																		if v155 != 0 {
																			return
																		} else {
																			v158 = int32(501)
																			F_errhint(m, int32(613991), int32(0))
																			mBase = m.M
																			v162 = m.ExcPending
																			if v162 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(521008), v158, int32(108140))
																				mBase = m.M
																				v166 = m.ExcPending
																				if v166 != 0 {
																					return
																				} else {
																					m.G0 = v9 + int32(48)
																					return
																				}
																			}
																		}
																	}
																}
															} else {
																v130 = F_errstart(m, int32(19), int32(0))
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return
																} else {
																	if v130 == int32(0) {
																		m.G0 = v9 + int32(48)
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v124
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v29 - v55
																		F_errmsg(m, int32(151147), v9+int32(16))
																		mBase = m.M
																		v141 = m.ExcPending
																		if v141 != 0 {
																			return
																		} else {
																			v158 = int32(494)
																			F_errhint(m, int32(613991), int32(0))
																			mBase = m.M
																			v162 = m.ExcPending
																			if v162 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(521008), v158, int32(108140))
																				mBase = m.M
																				v166 = m.ExcPending
																				if v166 != 0 {
																					return
																				} else {
																					m.G0 = v9 + int32(48)
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
					}
				} else {
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v36))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v55)) == int32(0) {
						v89 = base.B2i32(base.Ui32(v36) <= base.Ui32(v55))
					} else {
						v89 = base.B2i32(int32(0) <= v55-v36)
					}
					if v89 == int32(0) {
						if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v53))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v55)) == int32(0) {
							v112 = base.B2i32(base.Ui32(v53) <= base.Ui32(v55))
						} else {
							v112 = base.B2i32(int32(0) <= v55-v53)
						}
						if v112 == int32(0) {
							m.G0 = v9 + int32(48)
							return
						} else {
							v116 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
							if v116 != 0 {
								m.G0 = v9 + int32(48)
								return
							} else {
								v118 = *(*int32)(unsafe.Add(mBase, _consts[65]))
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
								if base.B2i32(v119 == int32(2)) == int32(0) {
									v146 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return
									} else {
										if v146 == int32(0) {
											m.G0 = v9 + int32(48)
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
											F_errmsg(m, int32(151086), v9)
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return
											} else {
												v158 = int32(501)
												F_errhint(m, int32(613991), int32(0))
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return
												} else {
													F_errfinish(m, int32(521008), v158, int32(108140))
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
														return
													} else {
														m.G0 = v9 + int32(48)
														return
													}
												}
											}
										}
									}
								} else {
									v124 = F_get_database_name(m, l1)
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return
									} else {
										if v124 == int32(0) {
											v146 = F_errstart(m, int32(19), int32(0))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return
											} else {
												if v146 == int32(0) {
													m.G0 = v9 + int32(48)
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
													F_errmsg(m, int32(151086), v9)
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return
													} else {
														v158 = int32(501)
														F_errhint(m, int32(613991), int32(0))
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return
														} else {
															F_errfinish(m, int32(521008), v158, int32(108140))
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return
															} else {
																m.G0 = v9 + int32(48)
																return
															}
														}
													}
												}
											}
										} else {
											v130 = F_errstart(m, int32(19), int32(0))
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return
											} else {
												if v130 == int32(0) {
													m.G0 = v9 + int32(48)
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v124
													*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v29 - v55
													F_errmsg(m, int32(151147), v9+int32(16))
													mBase = m.M
													v141 = m.ExcPending
													if v141 != 0 {
														return
													} else {
														v158 = int32(494)
														F_errhint(m, int32(613991), int32(0))
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return
														} else {
															F_errfinish(m, int32(521008), v158, int32(108140))
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return
															} else {
																m.G0 = v9 + int32(48)
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
					} else {
						v93 = int32(*(*uint8)(unsafe.Add(mBase, _consts[96])))
						if v93 != int32(1) {
							if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v53))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v55)) == int32(0) {
								v112 = base.B2i32(base.Ui32(v53) <= base.Ui32(v55))
							} else {
								v112 = base.B2i32(int32(0) <= v55-v53)
							}
							if v112 == int32(0) {
								m.G0 = v9 + int32(48)
								return
							} else {
								v116 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
								if v116 != 0 {
									m.G0 = v9 + int32(48)
									return
								} else {
									v118 = *(*int32)(unsafe.Add(mBase, _consts[65]))
									v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
									if base.B2i32(v119 == int32(2)) == int32(0) {
										v146 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return
										} else {
											if v146 == int32(0) {
												m.G0 = v9 + int32(48)
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
												F_errmsg(m, int32(151086), v9)
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return
												} else {
													v158 = int32(501)
													F_errhint(m, int32(613991), int32(0))
													mBase = m.M
													v162 = m.ExcPending
													if v162 != 0 {
														return
													} else {
														F_errfinish(m, int32(521008), v158, int32(108140))
														mBase = m.M
														v166 = m.ExcPending
														if v166 != 0 {
															return
														} else {
															m.G0 = v9 + int32(48)
															return
														}
													}
												}
											}
										}
									} else {
										v124 = F_get_database_name(m, l1)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return
										} else {
											if v124 == int32(0) {
												v146 = F_errstart(m, int32(19), int32(0))
												mBase = m.M
												v147 = m.ExcPending
												if v147 != 0 {
													return
												} else {
													if v146 == int32(0) {
														m.G0 = v9 + int32(48)
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
														*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
														F_errmsg(m, int32(151086), v9)
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return
														} else {
															v158 = int32(501)
															F_errhint(m, int32(613991), int32(0))
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return
															} else {
																F_errfinish(m, int32(521008), v158, int32(108140))
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(48)
																	return
																}
															}
														}
													}
												}
											} else {
												v130 = F_errstart(m, int32(19), int32(0))
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return
												} else {
													if v130 == int32(0) {
														m.G0 = v9 + int32(48)
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v124
														*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v29 - v55
														F_errmsg(m, int32(151147), v9+int32(16))
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return
														} else {
															v158 = int32(494)
															F_errhint(m, int32(613991), int32(0))
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return
															} else {
																F_errfinish(m, int32(521008), v158, int32(108140))
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(48)
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
						} else {
							v97 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
							if v97 != 0 {
								if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v53))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v55)) == int32(0) {
									v112 = base.B2i32(base.Ui32(v53) <= base.Ui32(v55))
								} else {
									v112 = base.B2i32(int32(0) <= v55-v53)
								}
								if v112 == int32(0) {
									m.G0 = v9 + int32(48)
									return
								} else {
									v116 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
									if v116 != 0 {
										m.G0 = v9 + int32(48)
										return
									} else {
										v118 = *(*int32)(unsafe.Add(mBase, _consts[65]))
										v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
										if base.B2i32(v119 == int32(2)) == int32(0) {
											v146 = F_errstart(m, int32(19), int32(0))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return
											} else {
												if v146 == int32(0) {
													m.G0 = v9 + int32(48)
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
													F_errmsg(m, int32(151086), v9)
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return
													} else {
														v158 = int32(501)
														F_errhint(m, int32(613991), int32(0))
														mBase = m.M
														v162 = m.ExcPending
														if v162 != 0 {
															return
														} else {
															F_errfinish(m, int32(521008), v158, int32(108140))
															mBase = m.M
															v166 = m.ExcPending
															if v166 != 0 {
																return
															} else {
																m.G0 = v9 + int32(48)
																return
															}
														}
													}
												}
											}
										} else {
											v124 = F_get_database_name(m, l1)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return
											} else {
												if v124 == int32(0) {
													v146 = F_errstart(m, int32(19), int32(0))
													mBase = m.M
													v147 = m.ExcPending
													if v147 != 0 {
														return
													} else {
														if v146 == int32(0) {
															m.G0 = v9 + int32(48)
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
															F_errmsg(m, int32(151086), v9)
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return
															} else {
																v158 = int32(501)
																F_errhint(m, int32(613991), int32(0))
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(521008), v158, int32(108140))
																	mBase = m.M
																	v166 = m.ExcPending
																	if v166 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(48)
																		return
																	}
																}
															}
														}
													}
												} else {
													v130 = F_errstart(m, int32(19), int32(0))
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return
													} else {
														if v130 == int32(0) {
															m.G0 = v9 + int32(48)
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v124
															*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v29 - v55
															F_errmsg(m, int32(151147), v9+int32(16))
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return
															} else {
																v158 = int32(494)
																F_errhint(m, int32(613991), int32(0))
																mBase = m.M
																v162 = m.ExcPending
																if v162 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(521008), v158, int32(108140))
																	mBase = m.M
																	v166 = m.ExcPending
																	if v166 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(48)
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
							} else {
								F_SendPostmasterSignal(m, int32(4))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v53))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v55)) == int32(0) {
										v112 = base.B2i32(base.Ui32(v53) <= base.Ui32(v55))
									} else {
										v112 = base.B2i32(int32(0) <= v55-v53)
									}
									if v112 == int32(0) {
										m.G0 = v9 + int32(48)
										return
									} else {
										v116 = int32(*(*uint8)(unsafe.Add(mBase, _consts[161])))
										if v116 != 0 {
											m.G0 = v9 + int32(48)
											return
										} else {
											v118 = *(*int32)(unsafe.Add(mBase, _consts[65]))
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+20))
											if base.B2i32(v119 == int32(2)) == int32(0) {
												v146 = F_errstart(m, int32(19), int32(0))
												mBase = m.M
												v147 = m.ExcPending
												if v147 != 0 {
													return
												} else {
													if v146 == int32(0) {
														m.G0 = v9 + int32(48)
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
														*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
														F_errmsg(m, int32(151086), v9)
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return
														} else {
															v158 = int32(501)
															F_errhint(m, int32(613991), int32(0))
															mBase = m.M
															v162 = m.ExcPending
															if v162 != 0 {
																return
															} else {
																F_errfinish(m, int32(521008), v158, int32(108140))
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(48)
																	return
																}
															}
														}
													}
												}
											} else {
												v124 = F_get_database_name(m, l1)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return
												} else {
													if v124 == int32(0) {
														v146 = F_errstart(m, int32(19), int32(0))
														mBase = m.M
														v147 = m.ExcPending
														if v147 != 0 {
															return
														} else {
															if v146 == int32(0) {
																m.G0 = v9 + int32(48)
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
																*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v29 - v55
																F_errmsg(m, int32(151086), v9)
																mBase = m.M
																v155 = m.ExcPending
																if v155 != 0 {
																	return
																} else {
																	v158 = int32(501)
																	F_errhint(m, int32(613991), int32(0))
																	mBase = m.M
																	v162 = m.ExcPending
																	if v162 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(521008), v158, int32(108140))
																		mBase = m.M
																		v166 = m.ExcPending
																		if v166 != 0 {
																			return
																		} else {
																			m.G0 = v9 + int32(48)
																			return
																		}
																	}
																}
															}
														}
													} else {
														v130 = F_errstart(m, int32(19), int32(0))
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return
														} else {
															if v130 == int32(0) {
																m.G0 = v9 + int32(48)
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v124
																*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v29 - v55
																F_errmsg(m, int32(151147), v9+int32(16))
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return
																} else {
																	v158 = int32(494)
																	F_errhint(m, int32(613991), int32(0))
																	mBase = m.M
																	v162 = m.ExcPending
																	if v162 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(521008), v158, int32(108140))
																		mBase = m.M
																		v166 = m.ExcPending
																		if v166 != 0 {
																			return
																		} else {
																			m.G0 = v9 + int32(48)
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
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
	F_initStringInfo(m, v12+int32(96))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L13
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	v22 = *(*int32)(unsafe.Add(mBase, _consts[163]))
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
	F_errmsg_internal(m, int32(32951), v12+int32(80))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(518151), int32(5674), int32(516214))
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
	v115 = F_errstart(m, int32(10), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L7
	} else {
		goto L22
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v66
	F_appendStringInfo(m, v12+int32(96), int32(63635), v12-int32(-64))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v75 < int32(2) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v81 = int32(1)
	goto L18
L18:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v81<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v91
	F_appendStringInfo(m, v12+int32(96), int32(65006), v12+int32(48))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L20
	}
L19:
	;
	goto L14
L20:
	;
	v101 = v81 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v101 < v102 {
		v81 = v101
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if v115 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v118 = int32(571067)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if base.Ui32(v120) <= base.Ui32(int32(19)) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
	F_pfree(m, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L7
	} else {
		goto L40
	}
L26:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v120<<(uint(int32(2))%32))+uint32(_consts[164])))
	v128 = v127
	goto L28
L27:
	;
	v128 = v118
	goto L28
L28:
	;
	if v117 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v130 = v117
	goto L31
L30:
	;
	v130 = int32(476075)
	goto L31
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if base.Ui32(v132) <= base.Ui32(int32(5)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132<<(uint(int32(2))%32))+uint32(_consts[165])))
	v140 = v139
	goto L34
L33:
	;
	v140 = v118
	goto L34
L34:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v142 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v140
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+20)) = uint32(v142)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v141
	v147 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v149
	v154 = int32(*(*uint8)(unsafe.Add(mBase, _consts[167])))
	if v154 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v155 = int32(711000)
	goto L37
L36:
	;
	v155 = int32(793540)
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v128
	F_errmsg_internal(m, int32(186414), v12)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(518151), int32(5698), int32(516214))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
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
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v8 = l1 - int32(1)
	if v8 < int32(0) {
		v75 = l0
	} else {
		if l1&int32(1) != 0 {
			v13 = int32(2)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l2+v8<<(uint(v13)%32))))
			if base.B2i32(base.Ui32(v13) < base.Ui32(v16))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
				if base.Ui32(l0) < base.Ui32(v16) {
					v28 = v16
				} else {
					v28 = l0
				}
			} else {
				if int32(0) <= l0-v16 {
					v28 = l0
				} else {
					v28 = v16
				}
			}
			v31 = v28
			v33 = l1 - int32(2)
		} else {
			v31 = l0
			v33 = v8
		}
		if v8 == int32(0) {
			v75 = v31
		} else {
			v38 = v31
			v39 = v33
			for {
				v45 = v39 << (uint(int32(2)) % 32)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l2+v45)))
				if base.Ui32(v38) < base.Ui32(int32(3)) {
					if base.Ui32(v47) <= base.Ui32(v38) {
						v56 = v38
					} else {
						v56 = v47
					}
				} else {
					if base.Ui32(v47) < base.Ui32(int32(3)) {
						if base.Ui32(v47) <= base.Ui32(v38) {
							v56 = v38
						} else {
							v56 = v47
						}
					} else {
						if v38-v47 < int32(0) {
							v56 = v47
						} else {
							v56 = v38
						}
					}
				}
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v45+(l2-int32(4)))))
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v58))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v56)) == int32(0) {
					if base.Ui32(v56) < base.Ui32(v58) {
						v70 = v58
					} else {
						v70 = v56
					}
				} else {
					if int32(0) <= v56-v58 {
						v70 = v56
					} else {
						v70 = v58
					}
				}
				if int32(1) < v39 {
					v38 = v70
					v39 = v39 - int32(2)
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
	var v38 int32
	_ = v38
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
	var v141 int32
	_ = v141
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
	v18 = *(*int32)(unsafe.Add(mBase, _consts[139]))
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
	v38 = l0
	v39 = int32(0)
	goto L7
L7:
	;
	v48 = base.I32_div_u_s(v38, int32(819))
	if l1 <= v39 {
		v84 = v39
		v89 = int32(0)
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v202 = F_LWLockAcquire(m, v198+int32(4992), int32(0))
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
	v96 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	v99 = int32(*(*uint16)(unsafe.Add(mBase, _consts[141])))
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
	v109 = F_SimpleLruReadPage(m, int32(4456676), base.I64_extend_i32_u(v48), int32(1), v38)
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
	v114 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112+v115)))
	v123 = v117 + (v38-v48*int32(819))*int32(10)
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
	v141 = v94
	goto L21
L19:
	;
	goto L20
L20:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _consts[140]))
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
	v149 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v150+v112)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l2+v39<<(uint(int32(2))%32)+v141<<(uint(int32(2))%32))))
	v158 = base.I32_rem_u_s(v156, int32(819))
	v161 = v152 + v158*int32(10)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+8)) = uint16(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v161))) = l3
	v165 = v141 + int32(1)
	if v165 != v126 {
		v141 = v165
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
	v38 = v196
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
	v205 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	*(*uint16)(unsafe.Add(mBase, uint32(v205)+16)) = uint16(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v205)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = l0
	v210 = *(*int32)(unsafe.Add(mBase, _consts[142]))
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
	v225 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+44)) = v30
	goto L35
L34:
	;
	goto L35
L35:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v228+int32(4992))
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
