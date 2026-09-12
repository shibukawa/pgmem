package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReorderBufferQueueChange(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v261 int64
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	v11 = F_ReorderBufferTXNByXid(m, l0, l1, int32(0), l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	return
L3:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v13&int32(8) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_ReorderBufferFreeChange(m, l0, l3, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.Ui32(int32(11)) < base.Ui32(v19) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = l2
	v38 = l3 + int32(52)
	v40 = v11 + int32(128)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	if v41 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	if int32(1)<<(uint(v19)%32)&int32(2319) == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v29 = v28
	goto L13
L12:
	;
	v29 = v11
	goto L13
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v30 | int32(256)
	goto L8
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v40
	goto L16
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+56)) = v40
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v38
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v11)+112))
	v52 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+112)) = v51 + v52
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v11)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+120)) = v55 + v52
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	switch v62 {
	case 0, 1, 2, 8:
		goto L24
	case 3:
		goto L23
	case 4:
		goto L22
	case 5:
		goto L21
	default:
		v101 = int32(64)
		goto L19
	case 11:
		goto L20
	}
L17:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+144)))
	if v138 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L18:
	;
	if v106 == int32(0) {
		goto L17
	} else {
		goto L29
	}
L19:
	;
	v106 = v101
	goto L18
L20:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v101 = v95<<(uint(int32(2))%32) - int32(-64)
	goto L19
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v106 = (v88+v89)<<(uint(int32(2))%32) + int32(136)
	goto L18
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v106 = v82<<(uint(int32(4))%32) - int32(-64)
	goto L18
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v77 = F_strlen(m, v76)
	mBase = m.M
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v106 = v77 + v78 + int32(73)
	goto L18
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	if v64 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v69 = v65 + int32(84)
	goto L27
L26:
	;
	v69 = int32(64)
	goto L27
L27:
	;
	if v63 == int32(0) {
		v101 = v69
		goto L19
	} else {
		goto L28
	}
L28:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v106 = v69 + v72 + int32(20)
	goto L18
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v109 == int32(7) {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+216)) = v113 + v106
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+40))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v117 + v106
	if v116 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v120 = v116
	goto L33
L32:
	;
	v120 = v112
	goto L33
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+220)) = v121 + v106
	if v113 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_remove(m, v124, v112+int32(204))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v129, v112+int32(204))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L2
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	goto L17
L39:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _consts[837]))
	if v218 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v141 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v142 = v141
	goto L43
L42:
	;
	v142 = v11
	goto L43
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if l4 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v169 == int32(8) {
		goto L53
	} else {
		goto L54
	}
L45:
	;
	v165 = v143 | int32(32)
	goto L47
L46:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v143&int32(32) == int32(0) {
		v168 = v143
		v169 = v146
		goto L44
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v168 = v165
	v169 = v167
	goto L44
L48:
	;
	if base.Ui32(int32(8)) < base.Ui32(v146) {
		v168 = v143
		v169 = v146
		goto L44
	} else {
		goto L49
	}
L49:
	;
	if int32(1)<<(uint(v146)%32)&int32(259) == int32(0) {
		v168 = v143
		v169 = v146
		goto L44
	} else {
		goto L50
	}
L50:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+32)))
	if v159 != int32(1) {
		v168 = v143
		v169 = v146
		goto L44
	} else {
		goto L51
	}
L51:
	;
	v165 = v143 & int32(-33)
	goto L47
L52:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v188 < int32(2) {
		goto L39
	} else {
		goto L58
	}
L53:
	;
	v184 = v168 | int32(32)
	goto L55
L54:
	;
	if base.Ui32(int32(1)) < base.Ui32(v169-int32(9)) {
		goto L52
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v184
	goto L52
L56:
	;
	if v168&int32(32) == int32(0) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	v184 = v168 & int32(-33)
	goto L55
L58:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+144)))
	if v192 != int32(1) {
		goto L39
	} else {
		goto L59
	}
L59:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v195)+32))
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v187)+16))
	goto L60
L60:
	;
	if base.Ui64(v196) < base.Ui64(v197) {
		goto L39
	} else {
		goto L61
	}
L61:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v199&int32(32) != 0 {
		goto L39
	} else {
		goto L62
	}
L62:
	;
	if v199&int32(256) == int32(0) {
		goto L39
	} else {
		goto L63
	}
L63:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v206&int32(4) == int32(0) {
		goto L39
	} else {
		goto L64
	}
L64:
	;
	F_ReorderBufferStreamTXN(m, l0, v142)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	goto L39
L66:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v223 = *(*int32)(unsafe.Add(mBase, _consts[838]))
	if base.Ui32(v221) < base.Ui32(v223<<(uint(int32(10))%32)) {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v228 = l0 + int32(12)
	goto L70
L69:
	;
	goto L68
L70:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v240 = *(*int32)(unsafe.Add(mBase, _consts[838]))
	if base.Ui32(v238) < base.Ui32(v240<<(uint(int32(10))%32)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if v238 == int32(0) {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+16))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v252 < int32(2) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _consts[837]))
	if v247 != int32(1) {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+8))
	v314 = v312 - int32(204)
	v315 = F_ReorderBufferCheckAndTruncateAbortedTXN(m, l0, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L2
	} else {
		goto L103
	}
L78:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+144)))
	if v256 != int32(1) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v250)+8))
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)+32))
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v251)+16))
	goto L80
L80:
	;
	if base.Ui64(v260) < base.Ui64(v261) {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v263 == int32(0) {
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v266 = int32(0)
	if v263 == v228 {
		goto L77
	} else {
		goto L83
	}
L83:
	;
	v270 = v263
	v275 = v266
	v276 = v266
	goto L84
L84:
	;
	v279 = v270 - int32(96)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	if v280&int32(2336) != int32(256) {
		v292 = v275
		v293 = v276
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v292 == int32(0) {
		goto L77
	} else {
		goto L99
	}
L86:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v294 != v228 {
		v270 = v294
		v275 = v292
		v276 = v293
		goto L84
	} else {
		goto L98
	}
L87:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v270)+124))
	if base.Ui32(v286) <= base.Ui32(v276) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v288 = v275
	goto L90
L89:
	;
	v288 = int32(0)
	goto L90
L90:
	;
	if v288 != 0 {
		v292 = v275
		v293 = v276
		goto L86
	} else {
		goto L91
	}
L91:
	;
	if v286 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v289 = v279
	goto L94
L93:
	;
	v289 = v275
	goto L94
L94:
	;
	if v286 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v290 = v286
	goto L97
L96:
	;
	v290 = v276
	goto L97
L97:
	;
	v292 = v289
	v293 = v290
	goto L86
L98:
	;
	goto L85
L99:
	;
	v298 = F_ReorderBufferCheckAndTruncateAbortedTXN(m, l0, v292)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	if v298 != 0 {
		goto L70
	} else {
		goto L101
	}
L101:
	;
	F_ReorderBufferStreamTXN(m, l0, v292)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	goto L70
L103:
	;
	if v315 != 0 {
		goto L70
	} else {
		goto L104
	}
L104:
	;
	F_ReorderBufferSerializeTXN(m, l0, v314)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	goto L70
}
