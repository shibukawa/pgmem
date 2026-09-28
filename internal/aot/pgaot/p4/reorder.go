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
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int64
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int64
	_ = v272
	var v273 int64
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	v11 = F_ReorderBufferTXNByXid(m, l0, l1, int32(0), l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v13&int32(8) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ReorderBufferFreeChange(m, l0, l3, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v26 = int32(0)
	if base.B2i32(base.Ui32(int32(11)) < base.Ui32(v19))|base.B2i32(int32(1)<<(uint(v19)%32)&int32(2319) == v26) == v26 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = l2
	v41 = v11 + int32(128)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v32 = v31
	goto L12
L11:
	;
	v32 = v11
	goto L12
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v33 | int32(256)
	goto L9
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v41
	goto L15
L14:
	;
	goto L15
L15:
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
		goto L23
	case 3:
		goto L22
	case 4:
		goto L21
	case 5:
		goto L20
	default:
		v104 = int32(64)
		goto L18
	case 11:
		goto L19
	}
L16:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+144)))
	if v141 != int32(1) {
		goto L38
	} else {
		goto L39
	}
L17:
	;
	if v109 == int32(0) {
		goto L16
	} else {
		goto L28
	}
L18:
	;
	v109 = v104
	goto L17
L19:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v104 = v98<<(uint(int32(2))%32) - int32(-64)
	goto L18
L20:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+24))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v109 = (v91+v92)<<(uint(int32(2))%32) + int32(136)
	goto L17
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v109 = v85<<(uint(int32(4))%32) - int32(-64)
	goto L17
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v80 = F_strlen(m, v79)
	mBase = m.M
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v109 = v80 + v81 + int32(73)
	goto L17
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	if v67 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v72 = v68 + int32(84)
	goto L26
L25:
	;
	v72 = int32(64)
	goto L26
L26:
	;
	if v66 == int32(0) {
		v104 = v72
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v109 = v72 + v75 + int32(20)
	goto L17
L28:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v112 == int32(7) {
		goto L16
	} else {
		goto L29
	}
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+216)) = v116 + v109
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)+40))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v120 + v109
	if v119 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v123 = v119
	goto L32
L31:
	;
	v123 = v115
	goto L32
L32:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+220)) = v124 + v109
	if v116 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_remove(m, v127, v115+int32(204))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v132, v115+int32(204))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	goto L16
L38:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueChange[0]))
	if base.Ui32(v226<<(uint(int32(10))%32)) <= base.Ui32(v224) {
		goto L63
	} else {
		goto L64
	}
L39:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	if v144 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v145 = v144
	goto L42
L41:
	;
	v145 = v11
	goto L42
L42:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	if l4 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v173 == int32(8) {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	v170 = v146 | int32(32)
	goto L46
L45:
	;
	v151 = int32(0)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if base.B2i32(v146&int32(32) == v151)|base.B2i32(base.Ui32(int32(8)) < base.Ui32(v153))|base.B2i32(int32(1)<<(uint(v153)%32)&int32(259) == v151) != 0 {
		v173 = v153
		v174 = v146
		goto L43
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v173 = v172
	v174 = v170
	goto L43
L47:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+32)))
	if v164 != int32(1) {
		v173 = v153
		v174 = v146
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v170 = v146 & int32(-33)
	goto L46
L49:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+16))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if v194 < int32(2) {
		goto L38
	} else {
		goto L54
	}
L50:
	;
	v190 = v174 | int32(32)
	goto L52
L51:
	;
	if base.B2i32(v174&int32(32) == int32(0))|base.B2i32(base.Ui32(int32(1)) < base.Ui32(v173-int32(9))) != 0 {
		goto L49
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v190
	goto L49
L53:
	;
	v190 = v174 & int32(-33)
	goto L52
L54:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+144)))
	if v198 != int32(1) {
		goto L38
	} else {
		goto L55
	}
L55:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v201)+32))
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v193)+16))
	goto L56
L56:
	;
	if base.Ui64(v202) < base.Ui64(v203) {
		goto L38
	} else {
		goto L57
	}
L57:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	if v205&int32(32)|base.B2i32(v205&int32(256) == int32(0)) != 0 {
		goto L38
	} else {
		goto L58
	}
L58:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v213&int32(4) == int32(0) {
		goto L38
	} else {
		goto L59
	}
L59:
	;
	F_ReorderBufferStreamTXN(m, l0, v145)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L38
L61:
	;
	return
L62:
	;
	v239 = l0 + int32(12)
	v244 = int32(1)
	goto L67
L63:
	;
	v230 = *(*int64)(unsafe.Add(mBase, uint32(l0)+208))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+208)) = v230 + int64(1)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueChange[1]))
	if v235 == int32(0) {
		goto L61
	} else {
		goto L66
	}
L66:
	;
	goto L62
L67:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueChange[0]))
	if base.Ui32(v250) < base.Ui32(v252<<(uint(int32(10))%32)) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	if v244 == int32(0) {
		goto L61
	} else {
		goto L104
	}
L69:
	;
	goto L68
L70:
	;
	if v250 == int32(0) {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+16))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	if v264 < int32(2) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_ReorderBufferQueueChange[1]))
	if v259 != int32(1) {
		goto L69
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+8))
	v327 = v325 - int32(204)
	v328 = F_ReorderBufferCheckAndTruncateAbortedTXN(m, l0, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L101
	}
L76:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267)+144)))
	if v268 != int32(1) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v271)+32))
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v263)+16))
	goto L78
L78:
	;
	if base.Ui64(v272) < base.Ui64(v273) {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v275 == int32(0) {
		goto L75
	} else {
		goto L80
	}
L80:
	;
	v278 = int32(0)
	if v275 == v239 {
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v282 = v275
	v285 = v278
	v286 = v278
	goto L82
L82:
	;
	v291 = v282 - int32(96)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if v292&int32(2336) != int32(256) {
		v303 = v285
		v304 = v286
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v303 == int32(0) {
		goto L75
	} else {
		goto L97
	}
L84:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v306 != v239 {
		v282 = v306
		v285 = v303
		v286 = v304
		goto L82
	} else {
		goto L96
	}
L85:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v282)+124))
	if base.Ui32(v298) <= base.Ui32(v286) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v300 = v285
	goto L88
L87:
	;
	v300 = int32(0)
	goto L88
L88:
	;
	if v300 != 0 {
		v303 = v285
		v304 = v286
		goto L84
	} else {
		goto L89
	}
L89:
	;
	if v298 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v301 = v291
	goto L92
L91:
	;
	v301 = v285
	goto L92
L92:
	;
	if v298 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v302 = v298
	goto L95
L94:
	;
	v302 = v286
	goto L95
L95:
	;
	v303 = v301
	v304 = v302
	goto L84
L96:
	;
	goto L83
L97:
	;
	v310 = F_ReorderBufferCheckAndTruncateAbortedTXN(m, l0, v303)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	if v310 != 0 {
		goto L67
	} else {
		goto L99
	}
L99:
	;
	F_ReorderBufferStreamTXN(m, l0, v303)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v244 = int32(0)
	goto L67
L101:
	;
	if v328 != 0 {
		goto L67
	} else {
		goto L102
	}
L102:
	;
	F_ReorderBufferSerializeTXN(m, l0, v327)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v244 = int32(0)
	goto L67
L104:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	F_UpdateDecodingStats(m, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	goto L61
}
