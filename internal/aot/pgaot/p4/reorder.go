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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int64
	_ = v267
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
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
	v26 = int32(0)
	if base.B2i32(base.Ui32(int32(11)) < base.Ui32(v19))|base.B2i32(int32(1)<<(uint(v19)%32)&int32(2319) == v26) == v26 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = l2
	v41 = v11 + int32(128)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	if v42 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v32 = v31
	goto L13
L12:
	;
	v32 = v11
	goto L13
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v33 | int32(256)
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v41
	goto L16
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+56)) = v41
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v48
	v51 = l3 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v51
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v11)+112))
	v55 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+112)) = v54 + v55
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v11)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+120)) = v58 + v55
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	switch v65 {
	case 0, 1, 2, 8:
		goto L24
	case 3:
		goto L23
	case 4:
		goto L22
	case 5:
		goto L21
	default:
		v104 = int32(64)
		goto L19
	case 11:
		goto L20
	}
L17:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+144)))
	if v141 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L18:
	;
	if v109 == int32(0) {
		goto L17
	} else {
		goto L29
	}
L19:
	;
	v109 = v104
	goto L18
L20:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v104 = v98<<(uint(int32(2))%32) - int32(-64)
	goto L19
L21:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v109 = (v91+v92)<<(uint(int32(2))%32) + int32(136)
	goto L18
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v109 = v85<<(uint(int32(4))%32) - int32(-64)
	goto L18
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v80 = F_strlen(m, v79)
	mBase = m.M
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v109 = v80 + v81 + int32(73)
	goto L18
L24:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	if v67 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v72 = v68 + int32(84)
	goto L27
L26:
	;
	v72 = int32(64)
	goto L27
L27:
	;
	if v66 == int32(0) {
		v104 = v72
		goto L19
	} else {
		goto L28
	}
L28:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v109 = v72 + v75 + int32(20)
	goto L18
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v112 == int32(7) {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+216)) = v116 + v109
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)+40))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v120 + v109
	if v119 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v123 = v119
	goto L33
L32:
	;
	v123 = v115
	goto L33
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+220)) = v124 + v109
	if v116 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_remove(m, v127, v115+int32(204))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v132, v115+int32(204))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
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
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueChange[0]))
	if v225 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L40:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v144 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v145 = v144
	goto L43
L42:
	;
	v145 = v11
	goto L43
L43:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	if l4 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v173 == int32(8) {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	v170 = v146 | int32(32)
	goto L47
L46:
	;
	v151 = int32(0)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.B2i32(v146&int32(32) == v151)|base.B2i32(base.Ui32(int32(8)) < base.Ui32(v153))|base.B2i32(int32(1)<<(uint(v153)%32)&int32(259) == v151) != 0 {
		v173 = v153
		v174 = v146
		goto L44
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v173 = v172
	v174 = v170
	goto L44
L48:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+32)))
	if v164 != int32(1) {
		v173 = v153
		v174 = v146
		goto L44
	} else {
		goto L49
	}
L49:
	;
	v170 = v146 & int32(-33)
	goto L47
L50:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+16))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if v194 < int32(2) {
		goto L39
	} else {
		goto L55
	}
L51:
	;
	v190 = v174 | int32(32)
	goto L53
L52:
	;
	if base.B2i32(v174&int32(32) == int32(0))|base.B2i32(base.Ui32(int32(1)) < base.Ui32(v173-int32(9))) != 0 {
		goto L50
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v190
	goto L50
L54:
	;
	v190 = v174 & int32(-33)
	goto L53
L55:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+144)))
	if v198 != int32(1) {
		goto L39
	} else {
		goto L56
	}
L56:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v201)+32))
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v193)+16))
	goto L57
L57:
	;
	if base.Ui64(v202) < base.Ui64(v203) {
		goto L39
	} else {
		goto L58
	}
L58:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	if v205&int32(32)|base.B2i32(v205&int32(256) == int32(0)) != 0 {
		goto L39
	} else {
		goto L59
	}
L59:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v213&int32(4) == int32(0) {
		goto L39
	} else {
		goto L60
	}
L60:
	;
	F_ReorderBufferStreamTXN(m, l0, v145)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	goto L39
L62:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueChange[1]))
	if base.Ui32(v228) < base.Ui32(v230<<(uint(int32(10))%32)) {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v235 = l0 + int32(12)
	goto L66
L65:
	;
	goto L64
L66:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueChange[1]))
	if base.Ui32(v245) < base.Ui32(v247<<(uint(int32(10))%32)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if v245 == int32(0) {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+16))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	if v259 < int32(2) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueChange[0]))
	if v254 != int32(1) {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	v321 = v319 - int32(204)
	v322 = F_ReorderBufferCheckAndTruncateAbortedTXN(m, l0, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L2
	} else {
		goto L99
	}
L74:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+144)))
	if v263 != int32(1) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v266)+32))
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v258)+16))
	goto L76
L76:
	;
	if base.Ui64(v267) < base.Ui64(v268) {
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v270 == int32(0) {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v273 = int32(0)
	if v270 == v235 {
		goto L73
	} else {
		goto L79
	}
L79:
	;
	v277 = v270
	v279 = v273
	v280 = v273
	goto L80
L80:
	;
	v286 = v277 - int32(96)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	if v287&int32(2336) != int32(256) {
		v298 = v279
		v299 = v280
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v299 == int32(0) {
		goto L73
	} else {
		goto L95
	}
L82:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	if v301 != v235 {
		v277 = v301
		v279 = v298
		v280 = v299
		goto L80
	} else {
		goto L94
	}
L83:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v277)+124))
	if base.Ui32(v293) <= base.Ui32(v279) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v295 = v280
	goto L86
L85:
	;
	v295 = int32(0)
	goto L86
L86:
	;
	if v295 != 0 {
		v298 = v279
		v299 = v280
		goto L82
	} else {
		goto L87
	}
L87:
	;
	if v293 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v296 = v286
	goto L90
L89:
	;
	v296 = v280
	goto L90
L90:
	;
	if v293 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v297 = v293
	goto L93
L92:
	;
	v297 = v279
	goto L93
L93:
	;
	v298 = v297
	v299 = v296
	goto L82
L94:
	;
	goto L81
L95:
	;
	v305 = F_ReorderBufferCheckAndTruncateAbortedTXN(m, l0, v299)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	if v305 != 0 {
		goto L66
	} else {
		goto L97
	}
L97:
	;
	F_ReorderBufferStreamTXN(m, l0, v299)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	goto L66
L99:
	;
	if v322 != 0 {
		goto L66
	} else {
		goto L100
	}
L100:
	;
	F_ReorderBufferSerializeTXN(m, l0, v321)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	goto L66
}
