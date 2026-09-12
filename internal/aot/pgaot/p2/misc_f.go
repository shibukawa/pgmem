package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_FigureColnameInternal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
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
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v375 int32
	_ = v375
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if l0 == v3 {
		v375 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return v375
L2:
	;
	v21 = l0
	goto L5
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v355
	v375 = v358
	goto L1
L4:
	;
	if v339 == int32(0) {
		v375 = v3
		goto L1
	} else {
		goto L93
	}
L5:
	;
	v35 = int32(2)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	switch v37 - int32(10) {
	case 0:
		goto L15
	default:
		v375 = v3
		goto L1
	case 3:
		goto L11
	case 12:
		goto L29
	case 22:
		goto L28
	case 26:
		goto L27
	case 28:
		goto L26
	case 29:
		goto L25
	case 30:
		goto L14
	case 31:
		goto L13
	case 59:
		goto L34
	case 61:
		goto L31
	case 63:
		goto L30
	case 64:
		goto L8
	case 66:
		goto L32
	case 69:
		goto L33
	case 70:
		v355 = int32(26118)
		v358 = v35
		goto L3
	case 85:
		goto L24
	case 112:
		goto L16
	case 117:
		goto L23
	case 118:
		goto L22
	case 119:
		goto L21
	case 120:
		goto L20
	case 121, 122:
		goto L19
	case 124:
		goto L18
	case 125:
		goto L17
	}
L6:
	;
	if v50&v51 == int32(0) {
		v339 = v314
		goto L4
	} else {
		goto L91
	}
L7:
	;
	goto L6
L8:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v311 != 0 {
		v21 = v311
		goto L5
	} else {
		goto L90
	}
L9:
	;
	if v285 != 0 {
		v355 = v285
		v358 = v35
		goto L3
	} else {
		goto L89
	}
L10:
	;
	if v105&v106 == int32(0) {
		v285 = v260
		goto L9
	} else {
		goto L87
	}
L11:
	;
	v355 = int32(255785)
	v358 = v35
	goto L3
L12:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v218<<(uint(int32(2))%32))+uint32(_consts[299])))
	v355 = v256
	v358 = v35
	goto L3
L13:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if base.Ui32(int32(7)) <= base.Ui32(v244) {
		v375 = v3
		goto L1
	} else {
		goto L86
	}
L14:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if base.Ui32(int32(15)) <= base.Ui32(v236) {
		v375 = v3
		goto L1
	} else {
		goto L85
	}
L15:
	;
	v355 = int32(327627)
	v358 = v35
	goto L3
L16:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if base.Ui32(v218) < base.Ui32(int32(3)) {
		goto L12
	} else {
		goto L81
	}
L17:
	;
	v355 = int32(333952)
	v358 = v35
	goto L3
L18:
	;
	v355 = int32(333966)
	v358 = v35
	goto L3
L19:
	;
	v355 = int32(24249)
	v358 = v35
	goto L3
L20:
	;
	v355 = int32(109321)
	v358 = v35
	goto L3
L21:
	;
	v355 = int32(338322)
	v358 = v35
	goto L3
L22:
	;
	v355 = int32(227286)
	v358 = v35
	goto L3
L23:
	;
	v355 = int32(242944)
	v358 = v35
	goto L3
L24:
	;
	v355 = int32(337972)
	v358 = v35
	goto L3
L25:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	switch v208 {
	case 0:
		v355 = int32(76763)
		v358 = v35
		goto L3
	case 1:
		goto L80
	default:
		v375 = v3
		goto L1
	}
L26:
	;
	v355 = int32(410021)
	v358 = v35
	goto L3
L27:
	;
	v355 = int32(30707)
	v358 = v35
	goto L3
L28:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v199 = F_FigureColnameInternal(m, v198, l1)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L70
	} else {
		goto L78
	}
L29:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	switch v187 {
	case 0:
		v355 = int32(116292)
		v358 = v35
		goto L3
	default:
		v375 = v3
		goto L1
	case 4:
		goto L74
	case 6:
		goto L75
	}
L30:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v166 = F_FigureColnameInternal(m, v165, l1)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L70
	} else {
		goto L71
	}
L31:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v160 != int32(5) {
		v375 = v3
		goto L1
	} else {
		goto L69
	}
L32:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v151+v152<<(uint(int32(2))%32)-int32(4))))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v355 = v159
	v358 = v35
	goto L3
L33:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v95 == int32(0) {
		goto L8
	} else {
		goto L52
	}
L34:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v40 == int32(0) {
		v375 = v3
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v43 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v44 <= v43 {
		v339 = v43
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v47 = int32(0)
	if v47 < v44 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v50 = v44
	goto L39
L38:
	;
	v50 = v47
	goto L39
L39:
	;
	v51 = int32(1)
	if v44 == v51 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v314 = v43
	v315 = int32(0)
	goto L7
L41:
	;
	goto L42
L42:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v59 = int32(0)
	v63 = v43
	v64 = v59
	v65 = v59
	goto L43
L43:
	;
	v77 = v58 + v64<<(uint(int32(2))%32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 == int32(468) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v314 = v89
	v315 = v91
	goto L7
L45:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v83 = v82
	goto L47
L46:
	;
	v83 = v63
	goto L47
L47:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v85 == int32(468) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v89 = v88
	goto L50
L49:
	;
	v89 = v83
	goto L50
L50:
	;
	v90 = int32(2)
	v91 = v64 + v90
	v93 = v65 + v90
	if v50&int32(2147483646) != v93 {
		v63 = v89
		v64 = v91
		v65 = v93
		goto L43
	} else {
		goto L51
	}
L51:
	;
	goto L44
L52:
	;
	v98 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v99 <= v98 {
		v285 = v98
		goto L9
	} else {
		goto L53
	}
L53:
	;
	v102 = int32(0)
	if v102 < v99 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v105 = v99
	goto L56
L55:
	;
	v105 = v102
	goto L56
L56:
	;
	v106 = int32(1)
	if v99 == v106 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v260 = v98
	v261 = int32(0)
	goto L10
L58:
	;
	goto L59
L59:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v114 = int32(0)
	v118 = v98
	v119 = v114
	v120 = v114
	goto L60
L60:
	;
	v132 = v113 + v119<<(uint(int32(2))%32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v134 == int32(468) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v260 = v144
	v261 = v146
	goto L10
L62:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v138 = v137
	goto L64
L63:
	;
	v138 = v118
	goto L64
L64:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v140 == int32(468) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v144 = v143
	goto L67
L66:
	;
	v144 = v138
	goto L67
L67:
	;
	v145 = int32(2)
	v146 = v119 + v145
	v148 = v120 + v145
	if v105&int32(2147483646) != v148 {
		v118 = v144
		v119 = v146
		v120 = v148
		goto L60
	} else {
		goto L68
	}
L68:
	;
	goto L61
L69:
	;
	v355 = int32(335211)
	v358 = v35
	goto L3
L70:
	;
	return int32(0)
L71:
	;
	if base.Ui32(int32(1)) < base.Ui32(v166) {
		v375 = int32(2)
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v172 == int32(0) {
		v375 = v166
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v176+v177<<(uint(int32(2))%32)-int32(4))))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v355 = v184
	v358 = int32(1)
	goto L3
L74:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v190 != int32(67) {
		v375 = v3
		goto L1
	} else {
		goto L76
	}
L75:
	;
	v355 = int32(26118)
	v358 = v35
	goto L3
L76:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189)+76))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	if v196 != 0 {
		v355 = v196
		v358 = v35
		goto L3
	} else {
		goto L77
	}
L77:
	;
	v375 = v3
	goto L1
L78:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v199) {
		v375 = int32(2)
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v355 = int32(357600)
	v358 = int32(1)
	goto L3
L80:
	;
	v355 = int32(77499)
	v358 = v35
	goto L3
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L70
	} else {
		goto L82
	}
L82:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v225
	F_errmsg_internal(m, int32(477421), v17)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L70
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(487944), int32(2034), int32(309555))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L70
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v236<<(uint(int32(2))%32))+uint32(_consts[300])))
	v355 = v243
	v358 = v35
	goto L3
L86:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v244<<(uint(int32(2))%32))+uint32(_consts[301])))
	v355 = v251
	v358 = v35
	goto L3
L87:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v274+v261<<(uint(int32(2))%32))))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	if v279 != int32(468) {
		v285 = v260
		goto L9
	} else {
		goto L88
	}
L88:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	v285 = v282
	goto L9
L89:
	;
	goto L8
L90:
	;
	v375 = v3
	goto L1
L91:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v328+v315<<(uint(int32(2))%32))))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	if v333 != int32(468) {
		v339 = v314
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	v339 = v336
	goto L4
L93:
	;
	v355 = v339
	v358 = v35
	goto L3
}
func F_ForgetPortalSnapshots(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v12 = *(*int32)(unsafe.Add(mBase, _consts[1238]))
	F_hash_seq_init(m, v7+int32(12), v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = F_hash_seq_search(m, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = v1
	v22 = v17
	goto L7
L5:
	;
	v35 = v1
	goto L6
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	goto L14
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+100))
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v35 = v29
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = int32(0)
	v29 = v20 + int32(1)
	goto L11
L10:
	;
	v29 = v20
	goto L11
L11:
	;
	v32 = F_hash_seq_search(m, v7+int32(12))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v32 != 0 {
		v20 = v29
		v22 = v32
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L8
L14:
	;
	if v39 != int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v44 = v1
	goto L18
L16:
	;
	v56 = v1
	goto L17
L17:
	;
	if v35 != v56 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v47 = v44 + int32(1)
	F_PopActiveSnapshot(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v56 = v47
	goto L17
L20:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	goto L21
L21:
	;
	if v51 != int32(0) {
		v44 = v47
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	m.G0 = v7 + int32(32)
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v35
	F_errmsg_internal(m, int32(657419), v7)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(491713), int32(1292), int32(117917))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_FreeBulkInsertState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 != 0 {
		F_ReleaseBuffer(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_bms_free(m, v6)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_bms_free(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_FreeSpaceMapPrepareTruncateRel(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v182 int32
	_ = v182
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
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v235 int32
	_ = v235
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int64
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	v3 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v22 == v3 {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v26
		v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v28
		v32 = F_smgropen(m, v20+int32(24), v25)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v32
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
			if v38 != 0 {
				v46 = v38
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+76))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+80))
				*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v40
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v40))) = v42
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
				v46 = v44
			}
			*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v46 + int32(1)
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v52 = v50
			v53 = int32(-1)
			v55 = F_smgrexists(m, v52, int32(1))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				if v55 == int32(0) {
					v371 = v53
					m.G0 = v20 + int32(48)
					return v371
				} else {
					v59 = int32(4069)
					v60 = base.I32_div_u_s(l1, v59)
					*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = base.I64_extend_i32_u(v60) << (uint(int64(32)) % 64)
					v67 = l1 - v60*v59
					if v67 != 0 {
						v68 = *(*int64)(unsafe.Add(mBase, uint32(v20)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v68
						v73 = F_fsm_readbuf(m, l0, v20+int32(16), int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							if v73 == int32(0) {
								v371 = v53
								m.G0 = v20 + int32(48)
								return v371
							} else {
								F_LockBuffer(m, v73, int32(2))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									v80 = int32(4470804)
									v82 = *(*int32)(unsafe.Add(mBase, _consts[7]))
									*(*int32)(unsafe.Add(mBase, _consts[7])) = v82 + int32(1)
									if v73 < int32(0) {
										v89 = *(*int32)(unsafe.Add(mBase, _consts[5]))
										v95 = *(*int32)(unsafe.Add(mBase, uint32(v89+(v73^int32(-1))<<(uint(int32(2))%32))))
										v103 = v95
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, _consts[6]))
										v103 = v97 + v73<<(uint(int32(13))%32) + int32(-8192)
									}
									v105 = v103 + int32(28)
									v108 = v67 + v105 + int32(4095)
									if base.Ui32(v103-int32(-8192)) <= base.Ui32(v108) {
									} else {
										v113 = int32(4069) - v67
										v114 = int32(3)
										v115 = v113 & v114
										if base.Ui32(v67-int32(4066)) < base.Ui32(v114) {
											v165 = v108
											v166 = int32(0)
										} else {
											v123 = int32(0)
											v127 = v108
											v128 = v123
											v131 = v123
											for {
												v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
												v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
												v144 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v127))) = uint16(v144)
												v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
												v148 = v127 + int32(2)
												v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
												*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v144)
												v157 = base.B2i32(v142|(v146|v149)|v143 != v144) | v128
												v158 = int32(4)
												v159 = v127 + v158
												v161 = v131 + v158
												if v161 != v113&int32(-4) {
													v127 = v159
													v128 = v157
													v131 = v161
													continue
												} else {
													break
												}
												break
											}
											v165 = v159
											v166 = v157
										}
										if v115 != 0 {
											v182 = v165
											v183 = v166
											v185 = v3
											for {
												v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
												v198 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v198)
												v200 = int32(1)
												v204 = base.B2i32(v197 != v198) | v183
												v206 = v185 + v200
												if v206 != v115 {
													v182 = v182 + v200
													v183 = v204
													v185 = v206
													continue
												} else {
													break
												}
												break
											}
											v211 = v204
										} else {
											v211 = v166
										}
										if v211&int32(1) == int32(0) {
										} else {
											v235 = int32(4094)
											for {
												if base.Ui32(int32(4081)) < base.Ui32(v235) {
													v266 = int32(0)
												} else {
													v251 = v235 << (uint(int32(1)) % 32)
													v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+v251)+1)))
													if v235 == int32(4081) {
														v266 = v253
													} else {
														v257 = v253 & int32(255)
														v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+(v251+int32(2))))))
														if base.Ui32(v261) < base.Ui32(v257) {
															v263 = v257
														} else {
															v263 = v261
														}
														v266 = v263
													}
												}
												v267 = v235 + v105
												v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
												if v268 != v266&int32(255) {
													*(*uint8)(unsafe.Add(mBase, uint32(v267))) = uint8(v266)
												} else {
												}
												if v235 != 0 {
													v235 = v235 - int32(1)
													continue
												} else {
													break
												}
												break
											}
										}
									}
									F_MarkBufferDirty(m, v73)
									mBase = m.M
									v293 = m.ExcPending
									if v293 != 0 {
										return int32(0)
									} else {
										v295 = int32(*(*uint8)(unsafe.Add(mBase, _consts[158])))
										if v295 != 0 {
											v320 = int32(4470804)
											v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
											*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
											F_UnlockReleaseBuffer(m, v73)
											mBase = m.M
											v327 = m.ExcPending
											if v327 != 0 {
												return int32(0)
											} else {
												v329 = base.I32_div_u_s(l1, int32(16556761))
												v371 = v60 + v329 + int32(3)
												m.G0 = v20 + int32(48)
												return v371
											}
										} else {
											v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+118)))
											if v297 != int32(112) {
												v320 = int32(4470804)
												v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
												*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
												F_UnlockReleaseBuffer(m, v73)
												mBase = m.M
												v327 = m.ExcPending
												if v327 != 0 {
													return int32(0)
												} else {
													v329 = base.I32_div_u_s(l1, int32(16556761))
													v371 = v60 + v329 + int32(3)
													m.G0 = v20 + int32(48)
													return v371
												}
											} else {
												v301 = *(*int32)(unsafe.Add(mBase, _consts[10]))
												if v301 <= int32(0) {
													v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v304 != 0 {
														v320 = int32(4470804)
														v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
														*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
														F_UnlockReleaseBuffer(m, v73)
														mBase = m.M
														v327 = m.ExcPending
														if v327 != 0 {
															return int32(0)
														} else {
															v329 = base.I32_div_u_s(l1, int32(16556761))
															v371 = v60 + v329 + int32(3)
															m.G0 = v20 + int32(48)
															return v371
														}
													} else {
														v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														if v305 != 0 {
															v320 = int32(4470804)
															v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
															*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
															F_UnlockReleaseBuffer(m, v73)
															mBase = m.M
															v327 = m.ExcPending
															if v327 != 0 {
																return int32(0)
															} else {
																v329 = base.I32_div_u_s(l1, int32(16556761))
																v371 = v60 + v329 + int32(3)
																m.G0 = v20 + int32(48)
																return v371
															}
														} else {
															v307 = *(*int32)(unsafe.Add(mBase, _consts[199]))
															v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+252))
															if base.B2i32(v308 != int32(0)) == int32(0) {
																v314 = int32(*(*uint8)(unsafe.Add(mBase, _consts[761])))
																if v314 != int32(1) {
																	v320 = int32(4470804)
																	v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																	*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
																	F_UnlockReleaseBuffer(m, v73)
																	mBase = m.M
																	v327 = m.ExcPending
																	if v327 != 0 {
																		return int32(0)
																	} else {
																		v329 = base.I32_div_u_s(l1, int32(16556761))
																		v371 = v60 + v329 + int32(3)
																		m.G0 = v20 + int32(48)
																		return v371
																	}
																} else {
																	F_log_newpage_buffer(m, v73, int32(0))
																	mBase = m.M
																	v319 = m.ExcPending
																	if v319 != 0 {
																		return int32(0)
																	} else {
																		v320 = int32(4470804)
																		v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																		*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
																		F_UnlockReleaseBuffer(m, v73)
																		mBase = m.M
																		v327 = m.ExcPending
																		if v327 != 0 {
																			return int32(0)
																		} else {
																			v329 = base.I32_div_u_s(l1, int32(16556761))
																			v371 = v60 + v329 + int32(3)
																			m.G0 = v20 + int32(48)
																			return v371
																		}
																	}
																}
															} else {
																F_log_newpage_buffer(m, v73, int32(0))
																mBase = m.M
																v319 = m.ExcPending
																if v319 != 0 {
																	return int32(0)
																} else {
																	v320 = int32(4470804)
																	v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																	*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
																	F_UnlockReleaseBuffer(m, v73)
																	mBase = m.M
																	v327 = m.ExcPending
																	if v327 != 0 {
																		return int32(0)
																	} else {
																		v329 = base.I32_div_u_s(l1, int32(16556761))
																		v371 = v60 + v329 + int32(3)
																		m.G0 = v20 + int32(48)
																		return v371
																	}
																}
															}
														}
													}
												} else {
													v307 = *(*int32)(unsafe.Add(mBase, _consts[199]))
													v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+252))
													if base.B2i32(v308 != int32(0)) == int32(0) {
														v314 = int32(*(*uint8)(unsafe.Add(mBase, _consts[761])))
														if v314 != int32(1) {
															v320 = int32(4470804)
															v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
															*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
															F_UnlockReleaseBuffer(m, v73)
															mBase = m.M
															v327 = m.ExcPending
															if v327 != 0 {
																return int32(0)
															} else {
																v329 = base.I32_div_u_s(l1, int32(16556761))
																v371 = v60 + v329 + int32(3)
																m.G0 = v20 + int32(48)
																return v371
															}
														} else {
															F_log_newpage_buffer(m, v73, int32(0))
															mBase = m.M
															v319 = m.ExcPending
															if v319 != 0 {
																return int32(0)
															} else {
																v320 = int32(4470804)
																v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
																F_UnlockReleaseBuffer(m, v73)
																mBase = m.M
																v327 = m.ExcPending
																if v327 != 0 {
																	return int32(0)
																} else {
																	v329 = base.I32_div_u_s(l1, int32(16556761))
																	v371 = v60 + v329 + int32(3)
																	m.G0 = v20 + int32(48)
																	return v371
																}
															}
														}
													} else {
														F_log_newpage_buffer(m, v73, int32(0))
														mBase = m.M
														v319 = m.ExcPending
														if v319 != 0 {
															return int32(0)
														} else {
															v320 = int32(4470804)
															v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
															*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
															F_UnlockReleaseBuffer(m, v73)
															mBase = m.M
															v327 = m.ExcPending
															if v327 != 0 {
																return int32(0)
															} else {
																v329 = base.I32_div_u_s(l1, int32(16556761))
																v371 = v60 + v329 + int32(3)
																m.G0 = v20 + int32(48)
																return v371
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
						v335 = base.I32_div_u_s(l1, int32(16556761))
						v338 = v60 + v335 + int32(2)
						v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v339 != 0 {
							v363 = v339
							v365 = F_smgrnblocks(m, v363, int32(1))
							mBase = m.M
							v366 = m.ExcPending
							if v366 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v365) <= base.Ui32(v338) {
									v368 = int32(-1)
								} else {
									v368 = v338
								}
								v371 = v368
								m.G0 = v20 + int32(48)
								return v371
							}
						} else {
							v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v341
							v343 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v20))) = v343
							v345 = F_smgropen(m, v20, v340)
							mBase = m.M
							v346 = m.ExcPending
							if v346 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v345
								v349 = *(*int32)(unsafe.Add(mBase, uint32(v345)+72))
								if v349 != 0 {
									v357 = v349
								} else {
									v350 = *(*int32)(unsafe.Add(mBase, uint32(v345)+76))
									v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)+80))
									*(*int32)(unsafe.Add(mBase, uint32(v350)+4)) = v351
									v353 = *(*int32)(unsafe.Add(mBase, uint32(v345)+76))
									*(*int32)(unsafe.Add(mBase, uint32(v351))) = v353
									v355 = *(*int32)(unsafe.Add(mBase, uint32(v345)+72))
									v357 = v355
								}
								*(*int32)(unsafe.Add(mBase, uint32(v345)+72)) = v357 + int32(1)
								v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v363 = v361
								v365 = F_smgrnblocks(m, v363, int32(1))
								mBase = m.M
								v366 = m.ExcPending
								if v366 != 0 {
									return int32(0)
								} else {
									if base.Ui32(v365) <= base.Ui32(v338) {
										v368 = int32(-1)
									} else {
										v368 = v338
									}
									v371 = v368
									m.G0 = v20 + int32(48)
									return v371
								}
							}
						}
					}
				}
			}
		}
	} else {
		v52 = v22
		v53 = int32(-1)
		v55 = F_smgrexists(m, v52, int32(1))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			if v55 == int32(0) {
				v371 = v53
				m.G0 = v20 + int32(48)
				return v371
			} else {
				v59 = int32(4069)
				v60 = base.I32_div_u_s(l1, v59)
				*(*int64)(unsafe.Add(mBase, uint32(v20)+40)) = base.I64_extend_i32_u(v60) << (uint(int64(32)) % 64)
				v67 = l1 - v60*v59
				if v67 != 0 {
					v68 = *(*int64)(unsafe.Add(mBase, uint32(v20)+40))
					*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v68
					v73 = F_fsm_readbuf(m, l0, v20+int32(16), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						if v73 == int32(0) {
							v371 = v53
							m.G0 = v20 + int32(48)
							return v371
						} else {
							F_LockBuffer(m, v73, int32(2))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								v80 = int32(4470804)
								v82 = *(*int32)(unsafe.Add(mBase, _consts[7]))
								*(*int32)(unsafe.Add(mBase, _consts[7])) = v82 + int32(1)
								if v73 < int32(0) {
									v89 = *(*int32)(unsafe.Add(mBase, _consts[5]))
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v89+(v73^int32(-1))<<(uint(int32(2))%32))))
									v103 = v95
								} else {
									v97 = *(*int32)(unsafe.Add(mBase, _consts[6]))
									v103 = v97 + v73<<(uint(int32(13))%32) + int32(-8192)
								}
								v105 = v103 + int32(28)
								v108 = v67 + v105 + int32(4095)
								if base.Ui32(v103-int32(-8192)) <= base.Ui32(v108) {
								} else {
									v113 = int32(4069) - v67
									v114 = int32(3)
									v115 = v113 & v114
									if base.Ui32(v67-int32(4066)) < base.Ui32(v114) {
										v165 = v108
										v166 = int32(0)
									} else {
										v123 = int32(0)
										v127 = v108
										v128 = v123
										v131 = v123
										for {
											v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
											v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
											v144 = int32(0)
											*(*uint16)(unsafe.Add(mBase, uint32(v127))) = uint16(v144)
											v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+3)))
											v148 = v127 + int32(2)
											v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
											*(*uint16)(unsafe.Add(mBase, uint32(v148))) = uint16(v144)
											v157 = base.B2i32(v142|(v146|v149)|v143 != v144) | v128
											v158 = int32(4)
											v159 = v127 + v158
											v161 = v131 + v158
											if v161 != v113&int32(-4) {
												v127 = v159
												v128 = v157
												v131 = v161
												continue
											} else {
												break
											}
											break
										}
										v165 = v159
										v166 = v157
									}
									if v115 != 0 {
										v182 = v165
										v183 = v166
										v185 = v3
										for {
											v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
											v198 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v198)
											v200 = int32(1)
											v204 = base.B2i32(v197 != v198) | v183
											v206 = v185 + v200
											if v206 != v115 {
												v182 = v182 + v200
												v183 = v204
												v185 = v206
												continue
											} else {
												break
											}
											break
										}
										v211 = v204
									} else {
										v211 = v166
									}
									if v211&int32(1) == int32(0) {
									} else {
										v235 = int32(4094)
										for {
											if base.Ui32(int32(4081)) < base.Ui32(v235) {
												v266 = int32(0)
											} else {
												v251 = v235 << (uint(int32(1)) % 32)
												v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+v251)+1)))
												if v235 == int32(4081) {
													v266 = v253
												} else {
													v257 = v253 & int32(255)
													v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+(v251+int32(2))))))
													if base.Ui32(v261) < base.Ui32(v257) {
														v263 = v257
													} else {
														v263 = v261
													}
													v266 = v263
												}
											}
											v267 = v235 + v105
											v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
											if v268 != v266&int32(255) {
												*(*uint8)(unsafe.Add(mBase, uint32(v267))) = uint8(v266)
											} else {
											}
											if v235 != 0 {
												v235 = v235 - int32(1)
												continue
											} else {
												break
											}
											break
										}
									}
								}
								F_MarkBufferDirty(m, v73)
								mBase = m.M
								v293 = m.ExcPending
								if v293 != 0 {
									return int32(0)
								} else {
									v295 = int32(*(*uint8)(unsafe.Add(mBase, _consts[158])))
									if v295 != 0 {
										v320 = int32(4470804)
										v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
										*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
										F_UnlockReleaseBuffer(m, v73)
										mBase = m.M
										v327 = m.ExcPending
										if v327 != 0 {
											return int32(0)
										} else {
											v329 = base.I32_div_u_s(l1, int32(16556761))
											v371 = v60 + v329 + int32(3)
											m.G0 = v20 + int32(48)
											return v371
										}
									} else {
										v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+118)))
										if v297 != int32(112) {
											v320 = int32(4470804)
											v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
											*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
											F_UnlockReleaseBuffer(m, v73)
											mBase = m.M
											v327 = m.ExcPending
											if v327 != 0 {
												return int32(0)
											} else {
												v329 = base.I32_div_u_s(l1, int32(16556761))
												v371 = v60 + v329 + int32(3)
												m.G0 = v20 + int32(48)
												return v371
											}
										} else {
											v301 = *(*int32)(unsafe.Add(mBase, _consts[10]))
											if v301 <= int32(0) {
												v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v304 != 0 {
													v320 = int32(4470804)
													v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
													*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
													F_UnlockReleaseBuffer(m, v73)
													mBase = m.M
													v327 = m.ExcPending
													if v327 != 0 {
														return int32(0)
													} else {
														v329 = base.I32_div_u_s(l1, int32(16556761))
														v371 = v60 + v329 + int32(3)
														m.G0 = v20 + int32(48)
														return v371
													}
												} else {
													v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													if v305 != 0 {
														v320 = int32(4470804)
														v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
														*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
														F_UnlockReleaseBuffer(m, v73)
														mBase = m.M
														v327 = m.ExcPending
														if v327 != 0 {
															return int32(0)
														} else {
															v329 = base.I32_div_u_s(l1, int32(16556761))
															v371 = v60 + v329 + int32(3)
															m.G0 = v20 + int32(48)
															return v371
														}
													} else {
														v307 = *(*int32)(unsafe.Add(mBase, _consts[199]))
														v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+252))
														if base.B2i32(v308 != int32(0)) == int32(0) {
															v314 = int32(*(*uint8)(unsafe.Add(mBase, _consts[761])))
															if v314 != int32(1) {
																v320 = int32(4470804)
																v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
																F_UnlockReleaseBuffer(m, v73)
																mBase = m.M
																v327 = m.ExcPending
																if v327 != 0 {
																	return int32(0)
																} else {
																	v329 = base.I32_div_u_s(l1, int32(16556761))
																	v371 = v60 + v329 + int32(3)
																	m.G0 = v20 + int32(48)
																	return v371
																}
															} else {
																F_log_newpage_buffer(m, v73, int32(0))
																mBase = m.M
																v319 = m.ExcPending
																if v319 != 0 {
																	return int32(0)
																} else {
																	v320 = int32(4470804)
																	v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																	*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
																	F_UnlockReleaseBuffer(m, v73)
																	mBase = m.M
																	v327 = m.ExcPending
																	if v327 != 0 {
																		return int32(0)
																	} else {
																		v329 = base.I32_div_u_s(l1, int32(16556761))
																		v371 = v60 + v329 + int32(3)
																		m.G0 = v20 + int32(48)
																		return v371
																	}
																}
															}
														} else {
															F_log_newpage_buffer(m, v73, int32(0))
															mBase = m.M
															v319 = m.ExcPending
															if v319 != 0 {
																return int32(0)
															} else {
																v320 = int32(4470804)
																v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
																*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
																F_UnlockReleaseBuffer(m, v73)
																mBase = m.M
																v327 = m.ExcPending
																if v327 != 0 {
																	return int32(0)
																} else {
																	v329 = base.I32_div_u_s(l1, int32(16556761))
																	v371 = v60 + v329 + int32(3)
																	m.G0 = v20 + int32(48)
																	return v371
																}
															}
														}
													}
												}
											} else {
												v307 = *(*int32)(unsafe.Add(mBase, _consts[199]))
												v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+252))
												if base.B2i32(v308 != int32(0)) == int32(0) {
													v314 = int32(*(*uint8)(unsafe.Add(mBase, _consts[761])))
													if v314 != int32(1) {
														v320 = int32(4470804)
														v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
														*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
														F_UnlockReleaseBuffer(m, v73)
														mBase = m.M
														v327 = m.ExcPending
														if v327 != 0 {
															return int32(0)
														} else {
															v329 = base.I32_div_u_s(l1, int32(16556761))
															v371 = v60 + v329 + int32(3)
															m.G0 = v20 + int32(48)
															return v371
														}
													} else {
														F_log_newpage_buffer(m, v73, int32(0))
														mBase = m.M
														v319 = m.ExcPending
														if v319 != 0 {
															return int32(0)
														} else {
															v320 = int32(4470804)
															v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
															*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
															F_UnlockReleaseBuffer(m, v73)
															mBase = m.M
															v327 = m.ExcPending
															if v327 != 0 {
																return int32(0)
															} else {
																v329 = base.I32_div_u_s(l1, int32(16556761))
																v371 = v60 + v329 + int32(3)
																m.G0 = v20 + int32(48)
																return v371
															}
														}
													}
												} else {
													F_log_newpage_buffer(m, v73, int32(0))
													mBase = m.M
													v319 = m.ExcPending
													if v319 != 0 {
														return int32(0)
													} else {
														v320 = int32(4470804)
														v322 = *(*int32)(unsafe.Add(mBase, _consts[7]))
														*(*int32)(unsafe.Add(mBase, _consts[7])) = v322 - int32(1)
														F_UnlockReleaseBuffer(m, v73)
														mBase = m.M
														v327 = m.ExcPending
														if v327 != 0 {
															return int32(0)
														} else {
															v329 = base.I32_div_u_s(l1, int32(16556761))
															v371 = v60 + v329 + int32(3)
															m.G0 = v20 + int32(48)
															return v371
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
					v335 = base.I32_div_u_s(l1, int32(16556761))
					v338 = v60 + v335 + int32(2)
					v339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v339 != 0 {
						v363 = v339
						v365 = F_smgrnblocks(m, v363, int32(1))
						mBase = m.M
						v366 = m.ExcPending
						if v366 != 0 {
							return int32(0)
						} else {
							if base.Ui32(v365) <= base.Ui32(v338) {
								v368 = int32(-1)
							} else {
								v368 = v338
							}
							v371 = v368
							m.G0 = v20 + int32(48)
							return v371
						}
					} else {
						v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v341
						v343 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
						*(*int64)(unsafe.Add(mBase, uint32(v20))) = v343
						v345 = F_smgropen(m, v20, v340)
						mBase = m.M
						v346 = m.ExcPending
						if v346 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v345
							v349 = *(*int32)(unsafe.Add(mBase, uint32(v345)+72))
							if v349 != 0 {
								v357 = v349
							} else {
								v350 = *(*int32)(unsafe.Add(mBase, uint32(v345)+76))
								v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)+80))
								*(*int32)(unsafe.Add(mBase, uint32(v350)+4)) = v351
								v353 = *(*int32)(unsafe.Add(mBase, uint32(v345)+76))
								*(*int32)(unsafe.Add(mBase, uint32(v351))) = v353
								v355 = *(*int32)(unsafe.Add(mBase, uint32(v345)+72))
								v357 = v355
							}
							*(*int32)(unsafe.Add(mBase, uint32(v345)+72)) = v357 + int32(1)
							v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v363 = v361
							v365 = F_smgrnblocks(m, v363, int32(1))
							mBase = m.M
							v366 = m.ExcPending
							if v366 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v365) <= base.Ui32(v338) {
									v368 = int32(-1)
								} else {
									v368 = v338
								}
								v371 = v368
								m.G0 = v20 + int32(48)
								return v371
							}
						}
					}
				}
			}
		}
	}
}
func F___ftello(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v2 < int32(0) {
		v5 = F___ftello_unlocked(m, l0)
		mBase = m.M
		return v5
	} else {
		v7 = F___ftello_unlocked(m, l0)
		mBase = m.M
		return v7
	}
}
func F___funcs_on_exit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1420]))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, _consts[1421]))
	v8 = v5
	v10 = v7
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v13 = v10 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1421])) = v13
	if int32(0) < v10 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v18 = v13
	goto L9
L7:
	;
	v42 = v8
	goto L8
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v44 = int32(32)
	*(*int32)(unsafe.Add(mBase, _consts[1421])) = v44
	*(*int32)(unsafe.Add(mBase, _consts[1420])) = v43
	if v43 != 0 {
		v8 = v43
		v10 = v44
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[1420]))
	v24 = v21 + v18<<(uint(int32(2))%32)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+132))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	m.T0[v26].(func(*base.Module, int32))(m, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[1420]))
	v42 = v38
	goto L8
L11:
	;
	return
L12:
	;
	v29 = int32(4652800)
	v31 = *(*int32)(unsafe.Add(mBase, _consts[1421]))
	v33 = v31 - int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1421])) = v33
	if int32(0) < v31 {
		v18 = v33
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L5
}
func F_fastgetattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	v14 = int32(1)
	v15 = l1 - v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+20)))
	if v17&v14 == v5 {
		v26 = l2 + v15<<(uint(int32(4))%32) + int32(20)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		if v27 < int32(0) {
			v73 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				v79 = v73
				m.G0 = v10 + int32(16)
				return v79
			}
		} else {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v32 = v16 + v30 + v27
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+6)))
			if v33 != int32(1) {
				v79 = v32
				m.G0 = v10 + int32(16)
				return v79
			} else {
				v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
				switch v36&int32(65535) - int32(1) {
				case 0:
					v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32))))
					v79 = v41
					m.G0 = v10 + int32(16)
					return v79
				case 1:
					v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32))))
					v79 = v42
					m.G0 = v10 + int32(16)
					return v79
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v36
						F_errmsg_internal(m, int32(477953), v10)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(323177), int32(70), int32(67251))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
					v79 = v43
					m.G0 = v10 + int32(16)
					return v79
				}
			}
		}
	} else {
		v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(base.Ui32(v15)>>(uint(int32(3))%32)))+23)))
		if int32(base.Ui32(v62)>>(uint(v15&int32(7))%32))&int32(1) != 0 {
			v73 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				v79 = v73
				m.G0 = v10 + int32(16)
				return v79
			}
		} else {
			v68 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v68)
			v79 = int32(0)
			m.G0 = v10 + int32(16)
			return v79
		}
	}
}
func F_ferror(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v2 < int32(0) {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v7 = v5
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v7 = v6
	}
	return int32(base.Ui32(v7)>>(uint(int32(5))%32)) & int32(1)
}
func F_fetch_input_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	if v4 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[1]))
		if v6 != 0 {
			F_ProcessInterrupts(m)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
				v12 = v11
				v14 = int32(0)
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
				v17 = F_tuplesort_gettupleslot(m, v12, int32(1), v14, v15, v14)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v17 == int32(0) {
						v44 = int32(0)
						return v44
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
						v30 = v21
						if v30 == int32(0) {
							return int32(0)
						} else {
							v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
							if v35&int32(2) != 0 {
								v44 = v30
								return v44
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
								if v38 == int32(0) {
									v44 = v30
									return v44
								} else {
									F_tuplesort_puttupleslot(m, v38, v30)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v44 = v30
										return v44
									}
								}
							}
						}
					}
				}
			}
		} else {
			v12 = v4
			v14 = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
			v17 = F_tuplesort_gettupleslot(m, v12, int32(1), v14, v15, v14)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v44 = int32(0)
					return v44
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
					v30 = v21
					if v30 == int32(0) {
						return int32(0)
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
						if v35&int32(2) != 0 {
							v44 = v30
							return v44
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
							if v38 == int32(0) {
								v44 = v30
								return v44
							} else {
								F_tuplesort_puttupleslot(m, v38, v30)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v44 = v30
									return v44
								}
							}
						}
					}
				}
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
		if v23 != 0 {
			F_ExecReScan(m, v22)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
				v27 = m.T0[v26].(func(*base.Module, int32) int32)(m, v22)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v30 = v27
					if v30 == int32(0) {
						return int32(0)
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
						if v35&int32(2) != 0 {
							v44 = v30
							return v44
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
							if v38 == int32(0) {
								v44 = v30
								return v44
							} else {
								F_tuplesort_puttupleslot(m, v38, v30)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v44 = v30
									return v44
								}
							}
						}
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
			v27 = m.T0[v26].(func(*base.Module, int32) int32)(m, v22)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v30 = v27
				if v30 == int32(0) {
					return int32(0)
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
					if v35&int32(2) != 0 {
						v44 = v30
						return v44
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
						if v38 == int32(0) {
							v44 = v30
							return v44
						} else {
							F_tuplesort_puttupleslot(m, v38, v30)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v44 = v30
								return v44
							}
						}
					}
				}
			}
		}
	}
}
func F_fillQT(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	F_check_stack_depth(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v8 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return
L4:
	;
	v12 = l1
	v14 = v7
	goto L7
L5:
	;
	v43 = l1
	v45 = v7
	goto L6
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v56 = v54 & int32(4095)
	if v56 != 0 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18 + int32(12)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	F_fillQT(m, l0, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v43 = v35
	v45 = v38
	goto L6
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v26 != int32(2) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = base.I32_div_s(v29-v15, int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v39 != int32(1) {
		v12 = v35
		v14 = v38
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v61 = int32(4095)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v66 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v60&v61 | (v63-v64)<<(uint(v66)%32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v76 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v70+v72&v61))) = uint8(v76)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v80 + v66
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v84 + v79&v61 + int32(1)
	goto L3
L14:
	;
	v57 = F__emscripten_memcpy_bulkmem(m, v51, v52, v56)
	mBase = m.M
	goto L16
L15:
	;
	goto L16
L16:
	;
	goto L13
}
func F_finalize_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
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
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
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
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
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
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int64
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v883 int32
	_ = v883
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L16
	} else {
		goto L245
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L16
	} else {
		goto L242
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L16
	} else {
		goto L239
	}
L4:
	;
	v18 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = l0
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v21 == v18 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v883 = int32(0)
	goto L6
L6:
	;
	m.G0 = v16 + int32(48)
	return v883
L7:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v142 = F_finalize_primnode(m, v139, v16+int32(40))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L16
	} else {
		goto L30
	}
L8:
	;
	v129 = l3
	v137 = v6
	v138 = v6
	goto L7
L9:
	;
	goto L10
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if int32(0) < v24 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = v6
	v38 = v6
	v39 = v6
	goto L14
L12:
	;
	v119 = v6
	v120 = v6
	goto L13
L13:
	;
	if v119 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v44 = int32(2)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v34<<(uint(v44)%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42+v48<<(uint(v44)%32)-int32(4))))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+64))
	v56 = F_bms_add_members(m, v39, v55)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v119 = v102
	v120 = v56
	goto L13
L16:
	;
	return int32(0)
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
	if v60 == int32(0) {
		v102 = v38
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v105 = v34 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v105 < v106 {
		v34 = v105
		v38 = v102
		v39 = v56
		goto L14
	} else {
		goto L25
	}
L19:
	;
	v63 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v64 <= v63 {
		v102 = v38
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v75 = v63
	v78 = v38
	goto L21
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v75<<(uint(int32(2))%32))))
	v85 = F_bms_add_member(m, v78, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v102 = v85
	goto L18
L23:
	;
	v88 = v75 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v88 < v89 {
		v75 = v88
		v78 = v85
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L15
L26:
	;
	v129 = l3
	v137 = int32(0)
	v138 = v120
	goto L7
L27:
	;
	goto L28
L28:
	;
	v124 = F_bms_union(m, l3, v119)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L16
	} else {
		goto L29
	}
L29:
	;
	v129 = v124
	v137 = v119
	v138 = v120
	goto L7
L30:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v147 = F_finalize_primnode(m, v144, v16+int32(40))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)))
	if v149 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if l2 < int32(0) {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v158 = int32(-1)
	v159 = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v160 - int32(331) {
	case 0:
		goto L38
	case 1, 29, 31, 32, 33, 36, 40:
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	case 2:
		goto L56
	case 3:
		goto L55
	case 4:
		goto L54
	case 5:
		goto L46
	case 6:
		goto L53
	case 7:
		goto L52
	case 8:
		goto L73
	case 9:
		goto L72
	case 10:
		goto L71
	case 11:
		goto L70
	case 12:
		goto L69
	case 13:
		goto L68
	case 14:
		goto L67
	case 15:
		goto L66
	case 16:
		goto L65
	case 17:
		goto L64
	case 18:
		goto L62
	case 19:
		goto L63
	case 20:
		goto L61
	case 21:
		goto L59
	case 22:
		goto L60
	case 23:
		goto L58
	case 24:
		goto L57
	case 25:
		goto L51
	default:
		goto L39
	case 27:
		goto L50
	case 28:
		goto L49
	case 30:
		goto L40
	case 34:
		goto L44
	case 35:
		goto L43
	case 37:
		goto L42
	case 38:
		goto L41
	case 39:
		goto L48
	case 41:
		goto L45
	case 42:
		goto L47
	}
L35:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v155 = F_bms_add_member(m, v154, l2)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v155
	goto L34
L37:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v770 = F_finalize_plan(m, l0, v769, v758, v759, v760)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L16
	} else {
		goto L204
	}
L38:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v754 = F_finalize_primnode(m, v751, v16+int32(40))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L16
	} else {
		goto L203
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L16
	} else {
		goto L200
	}
L40:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v735 = F_finalize_primnode(m, v732, v16+int32(40))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L16
	} else {
		goto L199
	}
L41:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v725 < int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v725
		goto L37
	} else {
		goto L196
	}
L42:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v718 < int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v718
		goto L37
	} else {
		goto L193
	}
L43:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v711 = F_finalize_primnode(m, v708, v16+int32(40))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L16
	} else {
		goto L191
	}
L44:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v690 != int32(2) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L188
	}
L45:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v682 = F_bms_copy(m, v129)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L16
	} else {
		goto L184
	}
L46:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v677 = F_bms_copy(m, v129)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L16
	} else {
		goto L182
	}
L47:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v669 = F_finalize_primnode(m, v666, v16+int32(40))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L16
	} else {
		goto L180
	}
L48:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v664 = F_finalize_primnode(m, v661, v16+int32(40))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L16
	} else {
		goto L179
	}
L49:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v654 = F_finalize_primnode(m, v651, v16+int32(40))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L16
	} else {
		goto L177
	}
L50:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v644 = F_finalize_primnode(m, v641, v16+int32(40))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L16
	} else {
		goto L175
	}
L51:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v607 = F_finalize_primnode(m, v604, v16+int32(40))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L16
	} else {
		goto L168
	}
L52:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v569 == int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L161
	}
L53:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v534 == int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L154
	}
L54:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v499 == int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L147
	}
L55:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v464 == int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L140
	}
L56:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	v441 = F_bms_copy(m, v129)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L16
	} else {
		goto L133
	}
L57:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v400 = F_finalize_primnode(m, v397, v16+int32(40))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L16
	} else {
		goto L124
	}
L58:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v386 = F_finalize_primnode(m, v383, v16+int32(40))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L16
	} else {
		goto L121
	}
L59:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v380 = F_bms_add_members(m, v379, l4)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L16
	} else {
		goto L120
	}
L60:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v373 = F_bms_add_member(m, v371, v372)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L16
	} else {
		goto L118
	}
L61:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v347 <= int32(0) {
		goto L2
	} else {
		goto L113
	}
L62:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v341 = F_finalize_primnode(m, v338, v16+int32(40))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L16
	} else {
		goto L111
	}
L63:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v332 = F_finalize_primnode(m, v329, v16+int32(40))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L16
	} else {
		goto L109
	}
L64:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v268 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L65:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v242 = F_find_base_rel(m, l0, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L16
	} else {
		goto L91
	}
L66:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v235 = F_finalize_primnode(m, v232, v16+int32(40))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L16
	} else {
		goto L89
	}
L67:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v226 = F_finalize_primnode(m, v223, v16+int32(40))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L16
	} else {
		goto L87
	}
L68:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v217 = F_finalize_primnode(m, v214, v16+int32(40))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L16
	} else {
		goto L85
	}
L69:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v212 = F_finalize_primnode(m, v209, v16+int32(40))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L16
	} else {
		goto L84
	}
L70:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v193 = F_finalize_primnode(m, v190, v16+int32(40))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L16
	} else {
		goto L80
	}
L71:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v179 = F_finalize_primnode(m, v176, v16+int32(40))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L16
	} else {
		goto L77
	}
L72:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v170 = F_finalize_primnode(m, v167, v16+int32(40))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L16
	} else {
		goto L75
	}
L73:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v164 = F_bms_add_members(m, v163, l4)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v164
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L75:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v173 = F_bms_add_members(m, v172, l4)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L16
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v173
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L77:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	v184 = F_finalize_primnode(m, v181, v16+int32(40))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L16
	} else {
		goto L78
	}
L78:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v187 = F_bms_add_members(m, v186, l4)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v187
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L80:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v198 = F_finalize_primnode(m, v195, v16+int32(40))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L16
	} else {
		goto L81
	}
L81:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	v203 = F_finalize_primnode(m, v200, v16+int32(40))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L16
	} else {
		goto L82
	}
L82:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v206 = F_bms_add_members(m, v205, l4)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L16
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v206
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L84:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L85:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v220 = F_bms_add_members(m, v219, l4)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L16
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v220
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L87:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v229 = F_bms_add_members(m, v228, l4)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L16
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v229
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L89:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v238 = F_bms_add_members(m, v237, l4)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L16
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v238
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L91:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v242)+140))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+24))
	if int32(0) <= l2 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v248 = F_bms_copy(m, v245)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L16
	} else {
		goto L95
	}
L93:
	;
	v253 = v245
	v254 = v244
	goto L94
L94:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v257 = F_finalize_plan(m, v254, v255, l2, v253, int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L16
	} else {
		goto L97
	}
L95:
	;
	v250 = F_bms_add_member(m, v248, l2)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L16
	} else {
		goto L96
	}
L96:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v242)+140))
	v253 = v250
	v254 = v252
	goto L94
L97:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+64))
	v262 = F_bms_add_members(m, v259, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L16
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v262
	v265 = F_bms_add_members(m, v262, l4)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L16
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v265
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L100:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v326 = F_bms_add_members(m, v325, l4)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L16
	} else {
		goto L108
	}
L101:
	;
	v271 = int32(0)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v272 <= v271 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v281 = v271
	goto L103
L103:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288+v281<<(uint(int32(2))%32))))
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = int32(0)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	v300 = F_finalize_primnode(m, v297, v16+int32(32))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L16
	} else {
		goto L105
	}
L104:
	;
	goto L100
L105:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v292)+28)) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v305 = F_bms_add_members(m, v304, v302)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L16
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v305
	v309 = v281 + int32(1)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v309 < v310 {
		v281 = v309
		goto L103
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v326
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L109:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v335 = F_bms_add_members(m, v334, l4)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L16
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v335
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L111:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v344 = F_bms_add_members(m, v343, l4)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L16
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v344
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L113:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)+8))
	if v351 == int32(0) {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	if v354 < v347 {
		goto L2
	} else {
		goto L115
	}
L115:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v357+v347<<(uint(int32(2))%32)-int32(4))))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+64))
	v365 = F_bms_add_members(m, v356, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L16
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v365
	v368 = F_bms_add_members(m, v365, l4)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L16
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v368
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v373
	v376 = F_bms_add_members(m, v373, l4)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L16
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v376
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v380
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L121:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v391 = F_finalize_primnode(m, v388, v16+int32(40))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L16
	} else {
		goto L122
	}
L122:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v394 = F_bms_add_members(m, v393, l4)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L16
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v394
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L124:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v403 = F_bms_add_members(m, v402, l4)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L16
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v403
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v406 == int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L126
	}
L126:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v409 <= int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L127
	}
L127:
	;
	v419 = int32(0)
	v423 = v403
	goto L128
L128:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v406)+12))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v426+v419<<(uint(int32(2))%32))))
	v431 = F_finalize_plan(m, l0, v430, l2, v129, l4)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L16
	} else {
		goto L130
	}
L129:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L130:
	;
	v433 = F_bms_add_members(m, v423, v431)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L16
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v433
	v437 = v419 + int32(1)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v437 < v438 {
		v419 = v437
		v423 = v433
		goto L128
	} else {
		goto L132
	}
L132:
	;
	goto L129
L133:
	;
	v443 = F_bms_add_member(m, v441, v440)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L16
	} else {
		goto L134
	}
L134:
	;
	v445 = F_bms_copy(m, l4)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L16
	} else {
		goto L135
	}
L135:
	;
	v447 = F_bms_add_member(m, v445, v440)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L16
	} else {
		goto L136
	}
L136:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v452 = F_finalize_primnode(m, v449, v16+int32(40))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L16
	} else {
		goto L137
	}
L137:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v457 = F_finalize_primnode(m, v454, v16+int32(40))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L16
	} else {
		goto L138
	}
L138:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	v462 = F_finalize_primnode(m, v459, v16+int32(40))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L16
	} else {
		goto L139
	}
L139:
	;
	v758 = l2
	v759 = v443
	v760 = v447
	v764 = v159
	v765 = v440
	goto L37
L140:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	if v467 <= int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L141
	}
L141:
	;
	v477 = int32(0)
	goto L142
L142:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v485+v477<<(uint(int32(2))%32))))
	v490 = F_finalize_plan(m, l0, v489, l2, v129, l4)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L16
	} else {
		goto L144
	}
L143:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L144:
	;
	v492 = F_bms_add_members(m, v484, v490)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L16
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v492
	v496 = v477 + int32(1)
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	if v496 < v497 {
		v477 = v496
		goto L142
	} else {
		goto L146
	}
L146:
	;
	goto L143
L147:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v502 <= int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L148
	}
L148:
	;
	v512 = int32(0)
	goto L149
L149:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v499)+12))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v520+v512<<(uint(int32(2))%32))))
	v525 = F_finalize_plan(m, l0, v524, l2, v129, l4)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L16
	} else {
		goto L151
	}
L150:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L151:
	;
	v527 = F_bms_add_members(m, v519, v525)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L16
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v527
	v531 = v512 + int32(1)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v531 < v532 {
		v512 = v531
		goto L149
	} else {
		goto L153
	}
L153:
	;
	goto L150
L154:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	if v537 <= int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L155
	}
L155:
	;
	v547 = int32(0)
	goto L156
L156:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v534)+12))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v555+v547<<(uint(int32(2))%32))))
	v560 = F_finalize_plan(m, l0, v559, l2, v129, l4)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L16
	} else {
		goto L158
	}
L157:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L158:
	;
	v562 = F_bms_add_members(m, v554, v560)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L16
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v562
	v566 = v547 + int32(1)
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	if v566 < v567 {
		v547 = v566
		goto L156
	} else {
		goto L160
	}
L160:
	;
	goto L157
L161:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v569)+4))
	if v572 <= int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L162
	}
L162:
	;
	v582 = int32(0)
	goto L163
L163:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v569)+12))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v590+v582<<(uint(int32(2))%32))))
	v595 = F_finalize_plan(m, l0, v594, l2, v129, l4)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L16
	} else {
		goto L165
	}
L164:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L165:
	;
	v597 = F_bms_add_members(m, v589, v595)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L16
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v597
	v601 = v582 + int32(1)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v569)+4))
	if v601 < v602 {
		v582 = v601
		goto L163
	} else {
		goto L167
	}
L167:
	;
	goto L164
L168:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v609 == int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L169
	}
L169:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v609)+4))
	if v612 <= int32(0) {
		v758 = l2
		v759 = v129
		v760 = l4
		v764 = v159
		v765 = v158
		goto L37
	} else {
		goto L170
	}
L170:
	;
	v622 = int32(0)
	v624 = v159
	goto L171
L171:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v609)+12))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v629+v622<<(uint(int32(2))%32))))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	v635 = F_bms_add_member(m, v624, v634)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L16
	} else {
		goto L173
	}
L172:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v635
	v765 = v158
	goto L37
L173:
	;
	v638 = v622 + int32(1)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v609)+4))
	if v638 < v639 {
		v622 = v638
		v624 = v635
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	v649 = F_finalize_primnode(m, v646, v16+int32(40))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L16
	} else {
		goto L176
	}
L176:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L177:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v659 = F_finalize_primnode(m, v656, v16+int32(40))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L16
	} else {
		goto L178
	}
L178:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L179:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L180:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v674 = F_finalize_primnode(m, v671, v16+int32(40))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L16
	} else {
		goto L181
	}
L181:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L182:
	;
	v679 = F_bms_add_member(m, v677, v676)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L16
	} else {
		goto L183
	}
L183:
	;
	v758 = l2
	v759 = v679
	v760 = l4
	v764 = v159
	v765 = v676
	goto L37
L184:
	;
	v684 = F_bms_add_member(m, v682, v681)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L16
	} else {
		goto L185
	}
L185:
	;
	v686 = F_bms_copy(m, l4)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L16
	} else {
		goto L186
	}
L186:
	;
	v688 = F_bms_add_member(m, v686, v681)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L16
	} else {
		goto L187
	}
L187:
	;
	v758 = l2
	v759 = v684
	v760 = v688
	v764 = v159
	v765 = v681
	goto L37
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l0
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v699 = F_finalize_agg_primnode(m, v696, v16+int32(32))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L16
	} else {
		goto L189
	}
L189:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v704 = F_finalize_agg_primnode(m, v701, v16+int32(32))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L16
	} else {
		goto L190
	}
L190:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v706
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L191:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
	v716 = F_finalize_primnode(m, v713, v16+int32(40))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L16
	} else {
		goto L192
	}
L192:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L193:
	;
	v721 = F_bms_copy(m, v129)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L16
	} else {
		goto L194
	}
L194:
	;
	v723 = F_bms_add_member(m, v721, v718)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L16
	} else {
		goto L195
	}
L195:
	;
	v758 = v718
	v759 = v723
	v760 = l4
	v764 = v159
	v765 = v718
	goto L37
L196:
	;
	v728 = F_bms_copy(m, v129)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L16
	} else {
		goto L197
	}
L197:
	;
	v730 = F_bms_add_member(m, v728, v725)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L16
	} else {
		goto L198
	}
L198:
	;
	v758 = v725
	v759 = v730
	v760 = l4
	v764 = v159
	v765 = v725
	goto L37
L199:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L200:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v741
	F_errmsg_internal(m, int32(480638), v16)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L16
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(488004), int32(2929), int32(280126))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L16
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	v758 = l2
	v759 = v129
	v760 = l4
	v764 = v159
	v765 = v158
	goto L37
L204:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v773 = F_bms_add_members(m, v772, v770)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L16
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v773
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v764 != 0 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v789 = F_bms_add_members(m, v788, v787)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L16
	} else {
		goto L215
	}
L207:
	;
	v777 = F_bms_union(m, v764, v759)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L16
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v785 = F_finalize_plan(m, l0, v776, v758, v759, v760)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L16
	} else {
		goto L214
	}
L210:
	;
	v779 = F_finalize_plan(m, l0, v776, v758, v777, v760)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L16
	} else {
		goto L211
	}
L211:
	;
	v781 = F_bms_difference(m, v779, v764)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L16
	} else {
		goto L212
	}
L212:
	;
	F_bms_free(m, v764)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L16
	} else {
		goto L213
	}
L213:
	;
	v787 = v781
	goto L206
L214:
	;
	v787 = v785
	goto L206
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v789
	if int32(0) <= v765 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v794 = F_bms_del_member(m, v789, v765)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L16
	} else {
		goto L219
	}
L217:
	;
	v797 = v789
	goto L218
L218:
	;
	v798 = int32(0)
	if v797 == v798 {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v794
	v797 = v794
	goto L218
L220:
	;
	if v851 == int32(0) {
		goto L1
	} else {
		goto L234
	}
L221:
	;
	v851 = int32(1)
	goto L220
L222:
	;
	goto L223
L223:
	;
	if v759 == int32(0) {
		v842 = v798
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v851 = v842
	goto L220
L225:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v797)+4))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v759)+4))
	if v808 < v807 {
		v842 = v798
		goto L224
	} else {
		goto L226
	}
L226:
	;
	v810 = int32(1)
	if v807 <= v810 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v813 = v810
	goto L229
L228:
	;
	v813 = v807
	goto L229
L229:
	;
	v814 = int32(8)
	v819 = int32(0)
	goto L230
L230:
	;
	v826 = v819 << (uint(int32(2)) % 32)
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v797+v814+v826)))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v826+(v759+v814))))
	v833 = v828 & (v830 ^ int32(-1))
	v835 = base.B2i32(v833 == int32(0))
	if v833 != 0 {
		v842 = v835
		goto L224
	} else {
		goto L232
	}
L231:
	;
	v842 = v835
	goto L224
L232:
	;
	v837 = v819 + int32(1)
	if v837 != v813 {
		v819 = v837
		goto L230
	} else {
		goto L233
	}
L233:
	;
	goto L231
L234:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v855 = F_bms_union(m, v854, v138)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L16
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v855
	v858 = F_bms_add_members(m, v855, v137)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L16
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v858
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v862 = F_bms_union(m, v861, v138)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L16
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v862
	v865 = F_bms_del_members(m, v862, v137)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L16
	} else {
		goto L238
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v865
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v883 = v868
	goto L6
L239:
	;
	F_errmsg_internal(m, int32(220991), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L16
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(488004), int32(2468), int32(280126))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L16
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v347
	F_errmsg_internal(m, int32(474507), v16+int32(16))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L16
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(488004), int32(2636), int32(280126))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L16
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L245:
	;
	F_errmsg_internal(m, int32(392257), int32(0))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L16
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(488004), int32(2978), int32(280126))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L16
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_find_among(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	v4 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = l2
	v25 = v4
	v27 = v4
	v28 = v4
	v30 = v4
	goto L1
L1:
	;
	v40 = v24
	v43 = v27
	v44 = v28
	v46 = v30
	goto L3
L2:
	;
	v145 = v117
	goto L28
L3:
	;
	if v44 < v46 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v134 = base.B2i32(v114 != v117) & base.B2i32(v117 <= int32(0))
	if v134 != 0 {
		goto L20
	} else {
		goto L21
	}
L5:
	;
	if int32(1) < v114-v117 {
		v40 = v114
		v43 = v117
		v44 = v118
		v46 = v120
		goto L3
	} else {
		goto L19
	}
L6:
	;
	v114 = v40
	v117 = v59
	v118 = v99
	v120 = v46
	goto L5
L7:
	;
	v55 = v44
	goto L9
L8:
	;
	v55 = v46
	goto L9
L9:
	;
	v59 = (v40-v43)>>(uint(int32(1))%32) + v43
	v62 = l1 + v59*int32(20)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 <= v55 {
		v99 = v55
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v68 = v55
	goto L13
L11:
	;
	v114 = v59
	v117 = v43
	v118 = v44
	v120 = v95
	goto L5
L12:
	;
	if int32(0) <= v88 {
		v99 = v68
		goto L6
	} else {
		goto L18
	}
L13:
	;
	if v17 == v68+v18 {
		v95 = v17 - v18
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v114 = v40
	v117 = v59
	v118 = v63
	v120 = v46
	goto L5
L15:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+(v20+v18)))))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v68))))
	v88 = v84 - v87
	if v88 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v90 = v68 + int32(1)
	if v90 != v63 {
		v68 = v90
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v95 = v68
	goto L11
L19:
	;
	goto L4
L20:
	;
	if v134 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	goto L2
L23:
	;
	v136 = int32(1)
	goto L25
L24:
	;
	v136 = v25
	goto L25
L25:
	;
	if v25 == int32(0) {
		v24 = v114
		v25 = v136
		v27 = v117
		v28 = v118
		v30 = v120
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	return v180
L28:
	;
	v158 = l1 + v145*int32(20)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v159 <= v118 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	v180 = v179
	goto L27
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v159 + v18
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	if v163 == int32(0) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v174 = int32(0)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	if v174 <= v175 {
		v145 = v175
		goto L28
	} else {
		goto L38
	}
L34:
	;
	v166 = m.T0[v163].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return int32(0)
L36:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170 + v18
	if v166 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v180 = v174
	goto L27
}
func F_find_composite_type_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	v12 = m.G0
	v14 = v12 - int32(160)
	m.G0 = v14
	F_check_stack_depth(m)
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
	v20 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v14-int32(-64), int32(4), int32(3), int32(184), int32(1247))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_ScanKeyInit(m, v14+int32(112), int32(5), int32(3), int32(184), l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v43 = F_systable_beginscan(m, v20, int32(2674), int32(1), int32(0), int32(2), v14-int32(-64))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L55
	}
L7:
	;
	v45 = F_systable_getnext(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v45 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v50 = v45
	goto L12
L10:
	;
	goto L11
L11:
	;
	F_systable_endscan(m, v43)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L53
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
	v60 = v58 + v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	switch v61 - int32(1247) {
	case 0:
		goto L16
	default:
		goto L14
	case 12:
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v218 = F_systable_getnext(m, v43)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L51
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v69 = F_relation_open(m, v67, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	F_find_composite_type_dependencies(m, v64, l1, l2)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+52))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v73 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	F_relation_close(m, v69, int32(1))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L50
	}
L20:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+119)))
	switch v129 - int32(73) {
	case 0, 10, 32, 36, 39, 41, 43:
		goto L35
	default:
		goto L34
	}
L21:
	;
	if v72 <= int32(0) {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	if v72 < v73 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v122 = v71 + v72<<(uint(int32(4))%32) + v73*int32(100) - int32(80)
	goto L20
L24:
	;
	v96 = int32(1)
	goto L25
L25:
	;
	v106 = v71 + v72<<(uint(int32(4))%32) - int32(80) + v96*int32(100)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+68))
	if l0 == v107 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	if v106 == int32(0) {
		goto L19
	} else {
		goto L33
	}
L27:
	;
	goto L26
L28:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+91)))
	if v109 != int32(1) {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v113 = v96 + int32(1)
	if v113 <= v72 {
		v96 = v113
		goto L25
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	goto L19
L33:
	;
	v122 = v106
	goto L20
L34:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v128)+72))
	if v188 == int32(0) {
		goto L19
	} else {
		goto L48
	}
L35:
	;
	if l2 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+119)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v141 = int32(4)
	v142 = v122 + v141
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
	v145 = v143 + v141
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v148 = v146 + v141
	switch v133 - int32(99) {
	case 0:
		goto L41
	default:
		goto L39
	case 3:
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v148
	F_errmsg(m, int32(363077), v14)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L46
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v148
	F_errmsg(m, int32(363142), v14+int32(32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L44
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v148
	F_errmsg(m, int32(102774), v14+int32(16))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(489058), int32(7051), int32(166911))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errfinish(m, int32(489058), int32(7058), int32(166911))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
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
	F_errfinish(m, int32(489058), int32(7065), int32(166911))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_find_composite_type_dependencies(m, v188, l1, l2)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L19
L50:
	;
	goto L14
L51:
	;
	if v218 != 0 {
		v50 = v218
		goto L12
	} else {
		goto L52
	}
L52:
	;
	goto L13
L53:
	;
	F_relation_close(m, v20, int32(1))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	m.G0 = v14 + int32(160)
	return
L55:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
	v247 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v122 + v247
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v246 + v247
	F_errmsg(m, int32(102774), v14+int32(48))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(489058), int32(7044), int32(166911))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_find_computable_ec_member(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	v14 = F_pull_var_clause(m, l2, int32(85))
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v18 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = l3
	goto L8
L7:
	;
	v25 = int32(0)
	goto L8
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v30 = int32(-1)
	v36 = v18
	v37 = v26
	goto L9
L9:
	;
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	if v137 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v127 = v30
	v131 = v40
	v133 = v36
	v134 = v37
	goto L11
L13:
	;
	goto L14
L14:
	;
	v43 = v30
	goto L15
L15:
	;
	if v25 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v127 = v108
	v131 = v124
	v133 = v121
	v134 = v124
	goto L11
L17:
	;
	if v108 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L18:
	;
	v108 = base.I32_ctz(v94) | v95<<(uint(int32(5))%32)
	goto L17
L19:
	;
	v108 = int32(-2)
	goto L17
L20:
	;
	v59 = v43 + int32(1)
	v61 = base.I32_div_s(v59, int32(32))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v62 <= v61 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v65 = v25 + int32(8)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v61<<(uint(int32(2))%32))))
	v72 = v69 & (int32(-1) << (uint(v59) % 32))
	if v72 != 0 {
		v94 = v72
		v95 = v61
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v74 = v61 + int32(1)
	if v74 == v62 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v77 = v74
	goto L24
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v65+v77<<(uint(int32(2))%32))))
	if v84 != 0 {
		v94 = v84
		v95 = v77
		goto L18
	} else {
		goto L26
	}
L25:
	;
	goto L19
L26:
	;
	v86 = v77 + int32(1)
	if v86 != v62 {
		v77 = v86
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	return int32(0)
L29:
	;
	goto L30
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v113 <= v108 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	goto L33
L33:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117+v108<<(uint(int32(2))%32))))
	if v121 == int32(0) {
		v43 = v108
		goto L15
	} else {
		goto L34
	}
L34:
	;
	goto L16
L35:
	;
	return int32(0)
L36:
	;
	goto L37
L37:
	;
	v143 = v134 + int32(4)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if base.Ui32(v143) < base.Ui32(v131+v145<<(uint(int32(2))%32)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v150 = v143
	goto L40
L39:
	;
	v150 = int32(0)
	goto L40
L40:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+12)))
	if v151 != 0 {
		v30 = v127
		v36 = v133
		v37 = v150
		goto L9
	} else {
		goto L41
	}
L41:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+13)))
	if v152 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v156 = int32(0)
	if v155 == v156 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L44
L44:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v214 = F_pull_var_clause(m, v212, int32(21))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L62
	}
L45:
	;
	if v209 == int32(0) {
		v30 = v127
		v36 = v133
		v37 = v150
		goto L9
	} else {
		goto L59
	}
L46:
	;
	v209 = int32(1)
	goto L45
L47:
	;
	goto L48
L48:
	;
	if l3 == int32(0) {
		v200 = v156
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v209 = v200
	goto L45
L50:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v166 < v165 {
		v200 = v156
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v168 = int32(1)
	if v165 <= v168 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v171 = v168
	goto L54
L53:
	;
	v171 = v165
	goto L54
L54:
	;
	v172 = int32(8)
	v177 = int32(0)
	goto L55
L55:
	;
	v184 = v177 << (uint(int32(2)) % 32)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v155+v172+v184)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184+(l3+v172))))
	v191 = v186 & (v188 ^ int32(-1))
	v193 = base.B2i32(v191 == int32(0))
	if v191 != 0 {
		v200 = v193
		goto L49
	} else {
		goto L57
	}
L56:
	;
	v200 = v193
	goto L49
L57:
	;
	v195 = v177 + int32(1)
	if v195 != v171 {
		v177 = v195
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	goto L44
L60:
	;
	F_list_free(m, v214)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L76
	}
L61:
	;
	F_list_free(m, v214)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L70
	}
L62:
	;
	if v214 == int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v218 = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v219 <= v218 {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	v228 = v218
	goto L65
L65:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234+v228<<(uint(int32(2))%32))))
	v239 = F_list_member(m, v14, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L61
L67:
	;
	if v239 == int32(0) {
		goto L60
	} else {
		goto L68
	}
L68:
	;
	v244 = v228 + int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v244 < v245 {
		v228 = v244
		goto L65
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	if l4 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v262 = F_is_parallel_safe(m, l0, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	return v137
L74:
	;
	if v262 == int32(0) {
		v30 = v127
		v36 = v133
		v37 = v150
		goto L9
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v30 = v127
	v36 = v133
	v37 = v150
	goto L9
}
func F_find_dependent_phvs_in_jointree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v260
L2:
	;
	v12 = F_bms_make_singleton(m, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v260 = int32(0)
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v16 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	if l1 == v16 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v104 = int32(0)
	v107 = F_get_relids_in_jointree(m, l1, v104, v104)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L34
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v21 == int32(319) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v97 = F_expression_tree_walker_impl(m, l1, int32(853), v8+int32(8))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L32
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v24 != 0 {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v78 = v21
	goto L12
L12:
	;
	if v78 != int32(67) {
		goto L9
	} else {
		goto L29
	}
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v26 = int32(0)
	v33 = base.B2i32(v12|v25 == v26)
	if v12 == v26 {
		v72 = v33
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v72 != 0 {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	goto L14
L16:
	;
	if v25 == int32(0) {
		v72 = v33
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v39 != v40 {
		v72 = int32(0)
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v42 = int32(1)
	if v39 <= v42 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v45 = v42
	goto L21
L20:
	;
	v45 = v39
	goto L21
L21:
	;
	v46 = int32(8)
	v51 = int32(0)
	goto L22
L22:
	;
	v59 = v51 << (uint(int32(2)) % 32)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v12+v46+v59)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+(v25+v46))))
	v64 = base.B2i32(v61 == v63)
	if v63 != v61 {
		v72 = v64
		goto L15
	} else {
		goto L24
	}
L23:
	;
	v72 = v64
	goto L15
L24:
	;
	v67 = v51 + int32(1)
	if v67 != v45 {
		v51 = v67
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v260 = int32(1)
	goto L1
L27:
	;
	goto L28
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v78 = v77
	goto L12
L29:
	;
	v81 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v81
	v88 = F_query_tree_walker_impl(m, l1, int32(853), v8+int32(8), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v90 - int32(1)
	if v88 != 0 {
		v260 = v81
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L7
L32:
	;
	if v97 == int32(0) {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v260 = int32(1)
	goto L1
L34:
	;
	if v107 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if v165 < int32(0) {
		v260 = v104
		goto L1
	} else {
		goto L46
	}
L36:
	;
	v165 = base.I32_ctz(v151) | v152<<(uint(int32(5))%32)
	goto L35
L37:
	;
	v165 = int32(-2)
	goto L35
L38:
	;
	v118 = base.I32_div_s(int32(0), int32(32))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v119 <= v118 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v122 = v107 + int32(8)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v118<<(uint(int32(2))%32))))
	v129 = v126 & int32(-1)
	if v129 != 0 {
		v151 = v129
		v152 = v118
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v131 = v118 + int32(1)
	if v131 == v119 {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v134 = v131
	goto L42
L42:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v122+v134<<(uint(int32(2))%32))))
	if v141 != 0 {
		v151 = v141
		v152 = v134
		goto L36
	} else {
		goto L44
	}
L43:
	;
	goto L37
L44:
	;
	v143 = v134 + int32(1)
	if v143 != v119 {
		v134 = v143
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v169 = v165
	goto L47
L47:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+52))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v175+v169<<(uint(int32(2))%32)-int32(4))))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+124)))
	if v182 != int32(1) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L4
L49:
	;
	if v107 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v189 = F_range_table_entry_walker_impl(m, v181, int32(853), v8+int32(8), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	if v189 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v260 = int32(1)
	goto L1
L53:
	;
	if int32(0) <= v249 {
		v169 = v249
		goto L47
	} else {
		goto L64
	}
L54:
	;
	v249 = base.I32_ctz(v235) | v236<<(uint(int32(5))%32)
	goto L53
L55:
	;
	v249 = int32(-2)
	goto L53
L56:
	;
	v200 = v169 + int32(1)
	v202 = base.I32_div_s(v200, int32(32))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v203 <= v202 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v206 = v107 + int32(8)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206+v202<<(uint(int32(2))%32))))
	v213 = v210 & (int32(-1) << (uint(v200) % 32))
	if v213 != 0 {
		v235 = v213
		v236 = v202
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v215 = v202 + int32(1)
	if v215 == v203 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v218 = v215
	goto L60
L60:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v206+v218<<(uint(int32(2))%32))))
	if v225 != 0 {
		v235 = v225
		v236 = v218
		goto L54
	} else {
		goto L62
	}
L61:
	;
	goto L55
L62:
	;
	v227 = v218 + int32(1)
	if v227 != v203 {
		v218 = v227
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L48
}
func F_find_dependent_phvs_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 == int32(319) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v88 = F_expression_tree_walker_impl(m, l0, int32(853), l1)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L25
	} else {
		goto L27
	}
L5:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 != v12 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v69 = v8
	goto L7
L7:
	;
	if v69 != int32(67) {
		goto L4
	} else {
		goto L24
	}
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = int32(0)
	v23 = base.B2i32(v14|v15 == v16)
	if v14 == v16 {
		v62 = v23
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v62 != 0 {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	goto L9
L11:
	;
	if v15 == int32(0) {
		v62 = v23
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v29 != v30 {
		v62 = int32(0)
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v32 = int32(1)
	if v29 <= v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v35 = v32
	goto L16
L15:
	;
	v35 = v29
	goto L16
L16:
	;
	v36 = int32(8)
	v41 = int32(0)
	goto L17
L17:
	;
	v49 = v41 << (uint(int32(2)) % 32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v14+v36+v49)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v15+v36))))
	v54 = base.B2i32(v51 == v53)
	if v53 != v51 {
		v62 = v54
		goto L10
	} else {
		goto L19
	}
L18:
	;
	v62 = v54
	goto L10
L19:
	;
	v57 = v41 + int32(1)
	if v57 != v35 {
		v41 = v57
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	return int32(1)
L22:
	;
	goto L23
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v69 = v68
	goto L7
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v72 + int32(1)
	v78 = F_query_tree_walker_impl(m, l0, int32(853), l1, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v82 - int32(1)
	return v78
L27:
	;
	return v88
}
func F_find_my_exec(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
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
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	goto L4
L1:
	;
	v128 = l1
	goto L36
L2:
	;
	v124 = F_strlen(m, v113)
	mBase = m.M
	goto L1
L4:
	;
	goto L5
L5:
	;
	v18 = int32(1023)
	if (l1^l0)&int32(3) != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v117)
	goto L2
L7:
	;
	v98 = v93
	v99 = v94
	v100 = v95
	goto L29
L8:
	;
	if v88 == int32(0) {
		v113 = v86
		v114 = v87
		goto L6
	} else {
		goto L28
	}
L9:
	;
	v86 = l0
	v87 = l1
	v88 = v18
	goto L8
L10:
	;
	goto L11
L11:
	;
	if l0&int32(3) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v55 == int32(0) {
		v113 = v52
		v114 = v53
		goto L6
	} else {
		goto L21
	}
L13:
	;
	v52 = l0
	v53 = l1
	v54 = v18
	v55 = int32(1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v31 = l0
	v32 = l1
	v33 = v18
	goto L16
L16:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v35)
	if v35 == int32(0) {
		v93 = v31
		v94 = v32
		v95 = v33
		goto L7
	} else {
		goto L18
	}
L17:
	;
	v52 = v46
	v53 = v40
	v54 = v42
	v55 = v44
	goto L12
L18:
	;
	v39 = int32(1)
	v40 = v32 + v39
	v42 = v33 - v39
	v43 = int32(0)
	v44 = base.B2i32(v42 != v43)
	v46 = v31 + v39
	if v46&int32(3) == v43 {
		v52 = v46
		v53 = v40
		v54 = v42
		v55 = v44
		goto L12
	} else {
		goto L19
	}
L19:
	;
	if v42 != 0 {
		v31 = v46
		v32 = v40
		v33 = v42
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v58 == int32(0) {
		v86 = v52
		v87 = v53
		v88 = v54
		goto L8
	} else {
		goto L22
	}
L22:
	;
	if base.Ui32(v54) < base.Ui32(int32(4)) {
		v86 = v52
		v87 = v53
		v88 = v54
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v64 = v52
	v65 = v53
	v66 = v54
	goto L24
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v72 = int32(-2139062144)
	if (int32(16843008)-v69|v69)&v72 != v72 {
		v93 = v64
		v94 = v65
		v95 = v66
		goto L7
	} else {
		goto L26
	}
L25:
	;
	v86 = v80
	v87 = v78
	v88 = v82
	goto L8
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v69
	v77 = int32(4)
	v78 = v65 + v77
	v80 = v64 + v77
	v82 = v66 - v77
	if base.Ui32(int32(3)) < base.Ui32(v82) {
		v64 = v80
		v65 = v78
		v66 = v82
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v93 = v86
	v94 = v87
	v95 = v88
	goto L7
L29:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v102)
	if v102 == int32(0) {
		v113 = v98
		v114 = v99
		goto L6
	} else {
		goto L31
	}
L30:
	;
	v113 = v109
	v114 = v107
	goto L6
L31:
	;
	v106 = int32(1)
	v107 = v99 + v106
	v109 = v98 + v106
	v111 = v100 - v106
	if v111 != 0 {
		v98 = v109
		v99 = v107
		v100 = v111
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	m.G0 = v10 + int32(144)
	return v526
L34:
	;
	v503 = int32(-1)
	v506 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L56
	} else {
		goto L170
	}
L35:
	;
	if v138 != 0 {
		goto L43
	} else {
		goto L44
	}
L36:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v130 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L35
L38:
	;
	goto L37
L39:
	;
	v138 = int32(0)
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v130 == int32(47) {
		v138 = v128
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v128 = v128 + int32(1)
	goto L36
L43:
	;
	v143 = F___fstatat(m, int32(-100), l1, v10+int32(48), int32(0))
	mBase = m.M
	goto L46
L44:
	;
	goto L45
L45:
	;
	v166 = int32(529944)
	v167 = int32(0)
	v172 = F___strchrnul(m, v166, int32(61))
	mBase = m.M
	if v166 == v172 {
		goto L61
	} else {
		goto L62
	}
L46:
	;
	if v143 < int32(0) {
		goto L34
	} else {
		goto L47
	}
L47:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v148 = v146 & int32(61440)
	if v148 != int32(32768) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v148 == int32(16384) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v159 = F_access(m, l1, int32(4))
	mBase = m.M
	v161 = F_access(m, l1, int32(1))
	mBase = m.M
	if v161 != 0 {
		goto L34
	} else {
		goto L54
	}
L51:
	;
	v156 = int32(31)
	goto L53
L52:
	;
	v156 = int32(63)
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, _consts[155])) = v156
	goto L34
L54:
	;
	if v159 != 0 {
		goto L34
	} else {
		goto L55
	}
L55:
	;
	v162 = F_normalize_exec_path(m, l1)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	return int32(0)
L57:
	;
	v526 = v162
	goto L33
L58:
	;
	v500 = F_normalize_exec_path(m, l1)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L56
	} else {
		goto L169
	}
L59:
	;
	v481 = int32(-1)
	v484 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L56
	} else {
		goto L164
	}
L60:
	;
	if v214 == int32(0) {
		goto L59
	} else {
		goto L76
	}
L61:
	;
	v214 = int32(0)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v175 = v172 - v166
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[1273]))))
	if v177 != 0 {
		v207 = v167
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v214 = v207
	goto L60
L65:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _consts[1229]))
	if v179 == int32(0) {
		v207 = v167
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	if v182 == int32(0) {
		v207 = v167
		goto L64
	} else {
		goto L67
	}
L67:
	;
	v186 = v179
	v187 = v182
	goto L68
L68:
	;
	v190 = F_strncmp(m, v166, v187, v175)
	mBase = m.M
	if v190 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v207 = v194 + int32(1)
	goto L64
L70:
	;
	goto L69
L71:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v194 = v193 + v175
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v195 == int32(61) {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v199 != 0 {
		v186 = v186 + int32(4)
		v187 = v199
		goto L68
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	v207 = v167
	goto L64
L76:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	if v217 == int32(0) {
		goto L59
	} else {
		goto L77
	}
L77:
	;
	v222 = v3
	v224 = v3
	goto L78
L78:
	;
	if v222 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L59
L80:
	;
	v229 = v224 + int32(1)
	goto L82
L81:
	;
	v229 = v214
	goto L82
L82:
	;
	v231 = v229
	goto L84
L83:
	;
	if v241 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L84:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v233 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L83
L86:
	;
	goto L85
L87:
	;
	v241 = int32(0)
	goto L86
L88:
	;
	goto L89
L89:
	;
	if v233 == int32(58) {
		v241 = v231
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v231 = v231 + int32(1)
	goto L84
L91:
	;
	if v229&int32(3) == int32(0) {
		v267 = v229
		goto L96
	} else {
		goto L97
	}
L92:
	;
	v302 = v241
	goto L93
L93:
	;
	v303 = int32(1024)
	v306 = v302 - v229 + int32(1)
	if v303 <= v306 {
		goto L111
	} else {
		goto L112
	}
L94:
	;
	v302 = v300 + v229
	goto L93
L95:
	;
	v300 = v292 - v229
	goto L94
L96:
	;
	v271 = v267
	goto L105
L97:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	if v251 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v300 = int32(0)
	goto L94
L99:
	;
	goto L100
L100:
	;
	v256 = v229
	goto L101
L101:
	;
	v260 = v256 + int32(1)
	if v260&int32(3) == int32(0) {
		v267 = v260
		goto L96
	} else {
		goto L103
	}
L102:
	;
	v292 = v260
	goto L95
L103:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if v265 != 0 {
		v256 = v260
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v280 = int32(-2139062144)
	if (int32(16843008)-v277|v277)&v280 == v280 {
		v271 = v271 + int32(4)
		goto L105
	} else {
		goto L107
	}
L106:
	;
	v286 = v271
	goto L108
L107:
	;
	goto L106
L108:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v290 != 0 {
		v286 = v286 + int32(1)
		goto L108
	} else {
		goto L110
	}
L109:
	;
	v292 = v286
	goto L95
L110:
	;
	goto L109
L111:
	;
	v309 = v303
	goto L113
L112:
	;
	v309 = v306
	goto L113
L113:
	;
	if v309 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	F_join_path_components(m, l1, l1, l0)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L56
	} else {
		goto L146
	}
L115:
	;
	v421 = F_strlen(m, v417)
	mBase = m.M
	goto L114
L116:
	;
	v417 = v229
	goto L115
L117:
	;
	goto L118
L118:
	;
	v315 = v309 - int32(1)
	if (l1^v229)&int32(3) != 0 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v414 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v411))) = uint8(v414)
	v417 = v410
	goto L115
L120:
	;
	v395 = v390
	v396 = v391
	v397 = v392
	goto L142
L121:
	;
	if v385 == int32(0) {
		v410 = v383
		v411 = v384
		goto L119
	} else {
		goto L141
	}
L122:
	;
	v383 = v229
	v384 = l1
	v385 = v315
	goto L121
L123:
	;
	goto L124
L124:
	;
	v319 = int32(0)
	if v229&int32(3) == v319 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v352 == int32(0) {
		v410 = v349
		v411 = v350
		goto L119
	} else {
		goto L134
	}
L126:
	;
	v349 = v229
	v350 = l1
	v351 = v315
	v352 = base.B2i32(v315 != v319)
	goto L125
L127:
	;
	if v315 == int32(0) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v328 = v229
	v329 = l1
	v330 = v315
	goto L129
L129:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	*(*uint8)(unsafe.Add(mBase, uint32(v329))) = uint8(v332)
	if v332 == int32(0) {
		v390 = v328
		v391 = v329
		v392 = v330
		goto L120
	} else {
		goto L131
	}
L130:
	;
	v349 = v343
	v350 = v337
	v351 = v339
	v352 = v341
	goto L125
L131:
	;
	v336 = int32(1)
	v337 = v329 + v336
	v339 = v330 - v336
	v340 = int32(0)
	v341 = base.B2i32(v339 != v340)
	v343 = v328 + v336
	if v343&int32(3) == v340 {
		v349 = v343
		v350 = v337
		v351 = v339
		v352 = v341
		goto L125
	} else {
		goto L132
	}
L132:
	;
	if v339 != 0 {
		v328 = v343
		v329 = v337
		v330 = v339
		goto L129
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	if v355 == int32(0) {
		v383 = v349
		v384 = v350
		v385 = v351
		goto L121
	} else {
		goto L135
	}
L135:
	;
	if base.Ui32(v351) < base.Ui32(int32(4)) {
		v383 = v349
		v384 = v350
		v385 = v351
		goto L121
	} else {
		goto L136
	}
L136:
	;
	v361 = v349
	v362 = v350
	v363 = v351
	goto L137
L137:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	v369 = int32(-2139062144)
	if (int32(16843008)-v366|v366)&v369 != v369 {
		v390 = v361
		v391 = v362
		v392 = v363
		goto L120
	} else {
		goto L139
	}
L138:
	;
	v383 = v377
	v384 = v375
	v385 = v379
	goto L121
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v366
	v374 = int32(4)
	v375 = v362 + v374
	v377 = v361 + v374
	v379 = v363 - v374
	if base.Ui32(int32(3)) < base.Ui32(v379) {
		v361 = v377
		v362 = v375
		v363 = v379
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	v390 = v383
	v391 = v384
	v392 = v385
	goto L120
L142:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	*(*uint8)(unsafe.Add(mBase, uint32(v396))) = uint8(v399)
	if v399 == int32(0) {
		v410 = v395
		v411 = v396
		goto L119
	} else {
		goto L144
	}
L143:
	;
	v410 = v406
	v411 = v404
	goto L119
L144:
	;
	v403 = int32(1)
	v404 = v396 + v403
	v406 = v395 + v403
	v408 = v397 - v403
	if v408 != 0 {
		v395 = v406
		v396 = v404
		v397 = v408
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	F_canonicalize_path_enc(m, l1)
	mBase = m.M
	v431 = F___fstatat(m, int32(-100), l1, v10+int32(48), int32(0))
	mBase = m.M
	goto L148
L147:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if v473 != 0 {
		v222 = v229
		v224 = v302
		goto L78
	} else {
		goto L163
	}
L148:
	;
	if v431 < int32(0) {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v436 = v434 & int32(61440)
	if v436 != int32(32768) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if v436 == int32(16384) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	v447 = F_access(m, l1, int32(4))
	mBase = m.M
	v449 = F_access(m, l1, int32(1))
	mBase = m.M
	if v449 != 0 {
		goto L147
	} else {
		goto L156
	}
L153:
	;
	v444 = int32(31)
	goto L155
L154:
	;
	v444 = int32(63)
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, _consts[155])) = v444
	goto L147
L156:
	;
	if v447 == int32(0) {
		goto L58
	} else {
		goto L157
	}
L157:
	;
	v454 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L56
	} else {
		goto L158
	}
L158:
	;
	if v454 == int32(0) {
		goto L147
	} else {
		goto L159
	}
L159:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L56
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
	F_errmsg_internal(m, int32(293631), v10+int32(16))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L56
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(494549), int32(218), int32(486127))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L56
	} else {
		goto L162
	}
L162:
	;
	goto L147
L163:
	;
	goto L79
L164:
	;
	if v484 == int32(0) {
		v526 = v481
		goto L33
	} else {
		goto L165
	}
L165:
	;
	F_errcode(m, int32(16908805))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L56
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	F_errmsg_internal(m, int32(344924), v10)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L56
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(494549), int32(225), int32(486127))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L56
	} else {
		goto L168
	}
L168:
	;
	v526 = v481
	goto L33
L169:
	;
	v526 = v500
	goto L33
L170:
	;
	if v506 == int32(0) {
		v526 = v503
		goto L33
	} else {
		goto L171
	}
L171:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L56
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l1
	F_errmsg_internal(m, int32(293607), v10+int32(32))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L56
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(494549), int32(174), int32(486127))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L56
	} else {
		goto L174
	}
L174:
	;
	v526 = v503
	goto L33
}
func F_findconstraintloop(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var __phi34 int32
	_ = __phi34
	var v35 int32
	_ = v35
	var __phi35 int32
	_ = __phi35
	var v38 int32
	_ = v38
	var __phi38 int32
	_ = __phi38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v260 int64
	_ = v260
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v337 int32
	_ = v337
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = m.T0[v14].(func(*base.Module) int32)(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = int32(101)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v29 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v25 = v23
	goto L8
L7:
	;
	v25 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v25
	return int32(1)
L9:
	;
	return v337
L10:
	;
	if l1 == v29 {
		v337 = int32(0)
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v289 != 0 {
		goto L114
	} else {
		goto L115
	}
L13:
	;
	__phi34 = v29
	__phi35 = int32(0)
	__phi38 = l1
	v34 = __phi34
	v35 = __phi35
	v38 = __phi38
	goto L14
L14:
	;
	if v35 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	__phi34 = v288
	__phi35 = v281
	__phi38 = v34
	v34 = __phi34
	v35 = __phi35
	v38 = __phi38
	goto L14
L17:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v109 != 0 {
		goto L43
	} else {
		goto L44
	}
L18:
	;
	if v84 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L19:
	;
	v92 = int32(0)
	if l1 != v34 {
		v281 = v92
		goto L16
	} else {
		goto L39
	}
L20:
	;
	v44 = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	if v46 == v44 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	v84 = v35
	goto L22
L22:
	;
	if l1 == v34 {
		goto L18
	} else {
		goto L38
	}
L23:
	;
	v51 = v46
	v56 = v44
	v57 = v44
	goto L24
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	if v34 != v59 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if v74 <= int32(1) {
		goto L35
	} else {
		goto L36
	}
L26:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v76 != 0 {
		v51 = v76
		v56 = v73
		v57 = v74
		goto L24
	} else {
		goto L34
	}
L27:
	;
	v71 = v56
	v73 = v56
	v74 = v57
	goto L26
L28:
	;
	goto L29
L29:
	;
	v61 = int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	switch v62 - int32(76) {
	case 0, 18, 21, 38:
		v68 = v51
		v69 = v61
		goto L30
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L31
	default:
		goto L32
	}
L30:
	;
	v71 = v68
	v73 = v68
	v74 = v69 + v57
	goto L26
L31:
	;
	v68 = v56
	v69 = int32(0)
	goto L30
L32:
	;
	if v62 == int32(36) {
		v68 = v51
		v69 = v61
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L25
L35:
	;
	v80 = v71
	goto L37
L36:
	;
	v80 = int32(0)
	goto L37
L37:
	;
	v84 = v80
	goto L22
L38:
	;
	v281 = v84
	goto L16
L39:
	;
	v100 = l1
	v102 = v92
	v104 = v29
	goto L17
L40:
	;
	v100 = l1
	v102 = int32(0)
	v104 = v29
	goto L17
L41:
	;
	goto L42
L42:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	v100 = v98
	v102 = v84
	v104 = v97
	goto L17
L43:
	;
	v112 = v109
	goto L46
L44:
	;
	goto L45
L45:
	;
	v133 = F_newstate(m, l0)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+24)) = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v112)+28))
	if v122 != 0 {
		v112 = v122
		goto L46
	} else {
		goto L48
	}
L47:
	;
	goto L45
L48:
	;
	goto L47
L49:
	;
	if v133 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	return int32(1)
L51:
	;
	goto L52
L52:
	;
	v139 = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_clonesuccessorstates(m, l0, v104, v133, v100, v102, v139, v139, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	if v145 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	return int32(1)
L55:
	;
	goto L56
L56:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	if v148 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v168 = v133
	goto L59
L58:
	;
	v149 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+4)) = uint8(v149)
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = int32(-1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v133)+32))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v133)+28))
	if v154 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	if v169 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L60:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v133)+28))
	if v153 != 0 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+32)) = v153
	goto L60
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v153
	goto L60
L64:
	;
	v160 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v133)+32)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+28)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v133
	v168 = v160
	goto L59
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+28)) = v157
	goto L64
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v157
	goto L64
L68:
	;
	return int32(1)
L69:
	;
	goto L70
L70:
	;
	v176 = v169
	goto L71
L71:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)+16))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	if v185 != v104 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	return int32(1)
L73:
	;
	if v184 != 0 {
		v176 = v184
		goto L71
	} else {
		goto L113
	}
L74:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	switch v187 - int32(76) {
	case 0, 18, 21, 38:
		goto L75
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L73
	default:
		goto L76
	}
L75:
	;
	if v168 != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	if v187 != int32(36) {
		goto L73
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	F_cparc(m, l0, v176, v100, v168)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
	v201 = int32(*(*int16)(unsafe.Add(mBase, uint32(v176)+4)))
	if v201 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L80
L82:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	if v270 != 0 {
		goto L109
	} else {
		goto L110
	}
L83:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v176)+16))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v176)+20))
	if v235 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L84:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v206 = v204 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v206) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	if int32(1)<<(uint(v206)%32)&int32(163841) == int32(0) {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v215 != 0 {
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v176)+36))
	if v216 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v228 != 0 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v176)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v220+v201*int32(24))+12)) = v224
	v228 = v224
	goto L88
L90:
	;
	goto L91
L91:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v176)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+32)) = v226
	v228 = v226
	goto L88
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228)+36)) = v216
	goto L94
L93:
	;
	goto L94
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v176)+32)) = int64(0)
	goto L83
L95:
	;
	if v234 != 0 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200)+20)) = v234
	goto L95
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235)+16)) = v234
	goto L95
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234)+20)) = v235
	goto L101
L100:
	;
	goto L101
L101:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v241 - int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v176)+24))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v176)+28))
	if v246 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v252 = v176 + int32(8)
	if v245 != 0 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+16)) = v245
	goto L102
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = v245
	goto L102
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+28)) = v246
	goto L108
L107:
	;
	goto L108
L108:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+8)) = v254 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = int32(0)
	v260 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v252)+16)) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v252)+8)) = v260
	*(*int64)(unsafe.Add(mBase, uint32(v252))) = v260
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+16)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v176
	goto L82
L109:
	;
	return int32(1)
L110:
	;
	goto L111
L111:
	;
	if v184 != 0 {
		v176 = v184
		goto L71
	} else {
		goto L112
	}
L112:
	;
	return int32(1)
L113:
	;
	goto L72
L114:
	;
	v292 = v289
	goto L117
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l1
	v337 = int32(0)
	goto L9
L117:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	switch v300 - int32(76) {
	case 0, 18, 21, 38:
		goto L120
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L119
	default:
		goto L121
	}
L118:
	;
	goto L116
L119:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v292)+16))
	if v314 != 0 {
		v292 = v314
		goto L117
	} else {
		goto L125
	}
L120:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v305
	v307 = F_findconstraintloop(m, l0, v305)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	if v300 != int32(36) {
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	if v307 == int32(0) {
		goto L119
	} else {
		goto L124
	}
L124:
	;
	return int32(1)
L125:
	;
	goto L118
}
func F_findoprnd_1(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	goto L1
L1:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = l0 + v13*int32(12)
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16))))
	switch v17 - int32(2) {
	case 0, 4:
		goto L6
	default:
		goto L5
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v13 + int32(1)
	if v28 == int32(33) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v23 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v13*int32(12))+2)) = uint16(v23)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v13 + int32(1)
	return
L7:
	;
	v34 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+2)) = uint16(v34)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_findoprnd_1(m, l0, l1)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v39 = v38 - v13
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+2)) = uint16(v39)
	goto L1
}
func F_findwrd(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var __phi26 int32
	_ = __phi26
	var v27 int32
	_ = v27
	var __phi27 int32
	_ = __phi27
	var v28 int32
	_ = v28
	var __phi28 int32
	_ = __phi28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v4 = int32(0)
	v8 = l0
	goto L1
L1:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if base.Ui32(v15-int32(9)) < base.Ui32(int32(5)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v60 = F_pg_mblen_cstr(m, v8)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L13
	} else {
		goto L24
	}
L4:
	;
	if v15 == int32(32) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	__phi26 = v8
	__phi27 = v15
	__phi28 = v8
	v26 = __phi26
	v27 = __phi27
	v28 = __phi28
	goto L9
L7:
	;
	v54 = v4
	v55 = v4
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v55
	return v54
L9:
	;
	switch v27 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L11
	default:
		goto L12
	}
L10:
	;
	if v28-v26 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	v29 = F_pg_mblen_cstr(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v33 = v29 + v28
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	__phi26 = v28
	__phi27 = v34
	__phi28 = v33
	v26 = __phi26
	v27 = __phi27
	v28 = __phi28
	goto L9
L15:
	;
	v54 = v8
	v55 = v50
	goto L8
L16:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v48)
	v50 = v47
	goto L15
L17:
	;
	if l2 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	if l2 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v40 != int32(42) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v47 = v26
	v48 = int32(2)
	goto L16
L21:
	;
	v50 = v28
	goto L15
L22:
	;
	goto L23
L23:
	;
	v47 = v28
	v48 = int32(0)
	goto L16
L24:
	;
	v8 = v60 + v8
	goto L1
}
func F_finite_interval_mi(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v13 = v11 - v12
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v13
	if base.B2i32(int32(0) < v12)^base.B2i32(v13 < v11) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				F_errmsg(m, int32(398189), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_errfinish(m, int32(490442), int32(3573), int32(316511))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
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
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v21 = v19 - v20
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v21
		if base.B2i32(v21 < v19)^base.B2i32(int32(0) < v20) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					F_errmsg(m, int32(398189), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						F_errfinish(m, int32(490442), int32(3573), int32(316511))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
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
			v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v29 = v27 - v28
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v29
			if base.B2i32(v29 < v27)^base.B2i32(int64(0) < v28) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						F_errmsg(m, int32(398189), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							F_errfinish(m, int32(490442), int32(3573), int32(316511))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
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
				if v13 != int32(2147483647) {
					if v13 != int32(-2147483648) {
						return
					} else {
						if v21 != int32(-2147483648) {
							return
						} else {
							if v29 == int64(-9223372036854775807-1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										F_errmsg(m, int32(398189), int32(0))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											F_errfinish(m, int32(490442), int32(3573), int32(316511))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
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
								return
							}
						}
					}
				} else {
					if v21 != int32(2147483647) {
						return
					} else {
						if v29 != int64(9223372036854775807) {
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_errmsg(m, int32(398189), int32(0))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_errfinish(m, int32(490442), int32(3573), int32(316511))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
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
					}
				}
			}
		}
	}
}
func F_finite_interval_pl(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v13 = v11 + v12
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v13
	if base.B2i32(v12 < int32(0))^base.B2i32(v13 < v11) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				F_errmsg(m, int32(398189), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					F_errfinish(m, int32(490442), int32(3517), int32(298083))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
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
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v21 = v19 + v20
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v21
		if base.B2i32(v20 < int32(0))^base.B2i32(v21 < v19) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					F_errmsg(m, int32(398189), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						F_errfinish(m, int32(490442), int32(3517), int32(298083))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
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
			v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v29 = v27 + v28
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v29
			if base.B2i32(v28 < int64(0))^base.B2i32(v29 < v27) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						F_errmsg(m, int32(398189), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							F_errfinish(m, int32(490442), int32(3517), int32(298083))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
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
				if v13 != int32(2147483647) {
					if v13 != int32(-2147483648) {
						return
					} else {
						if v21 != int32(-2147483648) {
							return
						} else {
							if v29 == int64(-9223372036854775807-1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										F_errmsg(m, int32(398189), int32(0))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											F_errfinish(m, int32(490442), int32(3517), int32(298083))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
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
								return
							}
						}
					}
				} else {
					if v21 != int32(2147483647) {
						return
					} else {
						if v29 != int64(9223372036854775807) {
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_errmsg(m, int32(398189), int32(0))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										F_errfinish(m, int32(490442), int32(3517), int32(298083))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
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
					}
				}
			}
		}
	}
}
func F_first_dir_separator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	v3 = l0
	goto L1
L1:
	;
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v5 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return v13
L3:
	;
	goto L2
L4:
	;
	v13 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	if v5 == int32(47) {
		v13 = v3
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v3 = v3 + int32(1)
	goto L1
}
func F_first_path_var_separator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	v3 = l0
	goto L1
L1:
	;
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v5 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return v13
L3:
	;
	goto L2
L4:
	;
	v13 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	if v5 == int32(58) {
		v13 = v3
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v3 = v3 + int32(1)
	goto L1
}
func F_fix_windowagg_condition_expr_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8 == int32(11) {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v14 = F_tlist_member(m, l0, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					v18 = F_makeVarFromTargetEntry(m, v11, v14)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v20 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v18)+40)) = uint16(v20)
						*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v20
						return v18
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(116416), int32(0))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(488669), int32(3458), int32(207148))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
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
			v39 = F_expression_tree_mutator_impl(m, l0, int32(841), l1)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				return v39
			}
		}
	}
}
func F_flatten_reloptions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v19 = F_SysCacheGetAttr(m, int32(57), v10, int32(33), v7+int32(31))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)))
				if v21 == int32(0) {
					F_initStringInfo(m, v7+int32(12))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_get_reloptions(m, v7+int32(12), v19)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							v33 = v32
							F_ReleaseCatCache(m, v10)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(32)
								return v33
							}
						}
					}
				} else {
					v33 = int32(0)
					F_ReleaseCatCache(m, v10)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(32)
						return v33
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(46015), v7)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(488513), int32(13652), int32(135660))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
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
func F_float4_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 float32
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v36 float64
	_ = v36
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v44 float64
	_ = v44
	var v48 float64
	_ = v48
	var v51 float64
	_ = v51
	var v64 int32
	_ = v64
	var v65 float64
	_ = v65
	var v76 float64
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v108 int32
	_ = v108
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		if v20 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v135 = m.ExcPending
			if v135 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(284459)
				F_errmsg_internal(m, int32(26050), v13)
				mBase = m.M
				v142 = m.ExcPending
				if v142 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(488100), int32(2938), int32(24590))
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
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			if v23 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v135 = m.ExcPending
				if v135 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(284459)
					F_errmsg_internal(m, int32(26050), v13)
					mBase = m.M
					v142 = m.ExcPending
					if v142 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(488100), int32(2938), int32(24590))
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
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
				if v26 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(284459)
						F_errmsg_internal(m, int32(26050), v13)
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(488100), int32(2938), int32(24590))
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
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
					if v27 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(284459)
							F_errmsg_internal(m, int32(26050), v13)
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(488100), int32(2938), int32(24590))
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
					} else {
						v30 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
						v31 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
						v32 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
						v33 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
						*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v33
						v36 = base.F64_add(v32, float64(1))
						*(*float64)(unsafe.Add(mBase, uint32(v13)+40)) = v36
						v38 = base.F64_promote_f32(v30)
						v39 = base.F64_add(v31, v38)
						*(*float64)(unsafe.Add(mBase, uint32(v13)+32)) = v39
						if base.F64_gt(v32, float64(0)) != 0 {
							v44 = base.F64_sub(base.F64_mul(v38, v36), v39)
							v48 = base.F64_add(v33, base.F64_div(base.F64_mul(v44, v44), base.F64_mul(v36, v32)))
							*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v48
							v51 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v39), v51)&base.F64_ne(base.F64_abs(v48), v51) != 0 {
								v76 = v48
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v80 == int32(0) {
									v108 = int32(0)
								} else {
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
									switch v83 - int32(429) {
									case 0:
										v108 = int32(1)
									case 1:
										v108 = int32(2)
									default:
										v108 = int32(0)
									}
								}
								if v108 != 0 {
									*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v76
									*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v39
									*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v36
									v127 = v16
									m.G0 = v13 + int32(48)
									return v127
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v13 + int32(24)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v13 + int32(32)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v13 + int32(40)
									v125 = F_construct_array_builtin(m, v13+int32(12), int32(3), int32(701))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										v127 = v125
										m.G0 = v13 + int32(48)
										return v127
									}
								}
							} else {
								if base.F64_eq(base.F64_abs(v31), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(9221120237041090560)
									v76 = math.Float64frombits(uint64(0x7ff8000000000000))
									v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v80 == int32(0) {
										v108 = int32(0)
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
										switch v83 - int32(429) {
										case 0:
											v108 = int32(1)
										case 1:
											v108 = int32(2)
										default:
											v108 = int32(0)
										}
									}
									if v108 != 0 {
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v76
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v39
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v36
										v127 = v16
										m.G0 = v13 + int32(48)
										return v127
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v13 + int32(24)
										*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v13 + int32(32)
										*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v13 + int32(40)
										v125 = F_construct_array_builtin(m, v13+int32(12), int32(3), int32(701))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int32(0)
										} else {
											v127 = v125
											m.G0 = v13 + int32(48)
											return v127
										}
									}
								} else {
									if base.F64_eq(base.F64_abs(v38), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(9221120237041090560)
										v76 = math.Float64frombits(uint64(0x7ff8000000000000))
										v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v80 == int32(0) {
											v108 = int32(0)
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
											switch v83 - int32(429) {
											case 0:
												v108 = int32(1)
											case 1:
												v108 = int32(2)
											default:
												v108 = int32(0)
											}
										}
										if v108 != 0 {
											*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v76
											*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v39
											*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v36
											v127 = v16
											m.G0 = v13 + int32(48)
											return v127
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v13 + int32(24)
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v13 + int32(32)
											*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v13 + int32(40)
											v125 = F_construct_array_builtin(m, v13+int32(12), int32(3), int32(701))
											mBase = m.M
											v126 = m.ExcPending
											if v126 != 0 {
												return int32(0)
											} else {
												v127 = v125
												m.G0 = v13 + int32(48)
												return v127
											}
										}
									} else {
										F_float_overflow_error(m)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
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
							v65 = base.F64_abs(v38)
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v65)) {
								*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(9221120237041090560)
								v76 = math.Float64frombits(uint64(0x7ff8000000000000))
							} else {
								if base.F64_ne(v65, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v76 = v33
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(9221120237041090560)
									v76 = math.Float64frombits(uint64(0x7ff8000000000000))
								}
							}
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v80 == int32(0) {
								v108 = int32(0)
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
								switch v83 - int32(429) {
								case 0:
									v108 = int32(1)
								case 1:
									v108 = int32(2)
								default:
									v108 = int32(0)
								}
							}
							if v108 != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v76
								*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v39
								*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v36
								v127 = v16
								m.G0 = v13 + int32(48)
								return v127
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v13 + int32(24)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v13 + int32(32)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v13 + int32(40)
								v125 = F_construct_array_builtin(m, v13+int32(12), int32(3), int32(701))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									v127 = v125
									m.G0 = v13 + int32(48)
									return v127
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_float4eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float32
	_ = v9
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = int32(2147483647)
	v8 = base.I32_reinterpret_f32(v5) & v7
	v9 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v9)&v7) {
		return base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v8))
	} else {
		return base.B2i32(base.Ui32(v8) < base.Ui32(int32(2139095041))) & base.F32_eq(v5, v9)
	}
}
func F_float4recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v5-v6 <= int32(3) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(400218), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(488108), int32(533), int32(157248))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v28+v6)))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v6 + int32(4)
		v34 = int32(24)
		v36 = int32(65280)
		v38 = int32(8)
		return v30<<(uint(v34)%32) | v30&v36<<(uint(v38)%32) | (int32(base.Ui32(v30)>>(uint(v38)%32))&v36 | int32(base.Ui32(v30)>>(uint(v34)%32)))
	}
}
func F_float84ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v13 float32
	_ = v13
	var v14 float64
	_ = v14
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.Ui64(base.I64_reinterpret_f64(v7)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v13 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = base.F64_promote_f32(v13)
		v23 = base.F64_ge(v7, v14) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v14)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	} else {
		v23 = int32(1)
	}
	return v23
}
func F_float8abs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v5 = F_Float8GetDatum(m, base.F64_abs(v3))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_float8out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v7 = F_palloc(m, int32(32))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[1072]))
		if int32(0) < v12 {
			F_double_to_shortest_decimal_buf(m, v5, v7)
			mBase = m.M
			return v7
		} else {
			F_pg_strfromd(m, v7, v12+int32(15), v5)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v7
			}
		}
	}
}
func F_float8recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pq_getmsgfloat8(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_Float8GetDatum(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_float_compare_desc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float32
	_ = v6
	var v7 float32
	_ = v7
	var v10 int32
	_ = v10
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	if base.F32_gt(v6, v7) != 0 {
		v10 = int32(-1)
	} else {
		v10 = base.F32_lt(v6, v7)
	}
	return v10
}
func F_fmodl(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v126 int32
	_ = v126
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v135 int32
	_ = v135
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v193 int32
	_ = v193
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v248 int32
	_ = v248
	var v254 int64
	_ = v254
	var v260 int64
	_ = v260
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v271 int64
	_ = v271
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v281 int64
	_ = v281
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v288 int64
	_ = v288
	var v294 int64
	_ = v294
	var v296 int32
	_ = v296
	var v300 int64
	_ = v300
	var v305 int64
	_ = v305
	var v308 int64
	_ = v308
	var v314 int64
	_ = v314
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v334 int32
	_ = v334
	var v340 int64
	_ = v340
	var v342 int32
	_ = v342
	var v346 int64
	_ = v346
	var v351 int64
	_ = v351
	var v354 int64
	_ = v354
	var v360 int64
	_ = v360
	var v363 int64
	_ = v363
	var v364 int64
	_ = v364
	var v365 int64
	_ = v365
	var v366 int64
	_ = v366
	var v373 int64
	_ = v373
	var v375 int32
	_ = v375
	var v378 int64
	_ = v378
	var v384 int32
	_ = v384
	var v385 int64
	_ = v385
	var v386 int64
	_ = v386
	var v389 int64
	_ = v389
	var v396 int64
	_ = v396
	var v398 int32
	_ = v398
	var v401 int64
	_ = v401
	var v405 int32
	_ = v405
	var v422 int64
	_ = v422
	var v423 int64
	_ = v423
	var v433 int64
	_ = v433
	var v435 int64
	_ = v435
	v10 = int64(0)
	v13 = m.G0
	v15 = v13 - int32(128)
	m.G0 = v15
	v26 = l4 & int64(9223372036854775807)
	v27 = int64(9223090561878065152)
	if v26 == v27 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v435
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v433
	m.G0 = v15 + int32(128)
	return
L2:
	;
	v119 = int64(9223372036854775807)
	v120 = l2 & v119
	v122 = l4 & v119
	v126 = int32(1)
	v130 = v120 & v119
	v131 = int64(9223090561878065152)
	if v130 == v131 {
		goto L46
	} else {
		goto L47
	}
L3:
	;
	F___multf3(m, v15+int32(16), l1, l2, l3, l4)
	mBase = m.M
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	F___divtf3(m, v15, v114, v115, v114, v115)
	mBase = m.M
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v433 = v117
	v435 = v118
	goto L1
L4:
	;
	if v74 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L5:
	;
	v74 = v70
	goto L4
L6:
	;
	v31 = base.B2i32(l3 != v10)
	goto L8
L7:
	;
	v31 = base.B2i32(base.Ui64(v27) < base.Ui64(v26))
	goto L8
L8:
	;
	if v31 != 0 {
		v70 = int32(1)
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L11
L11:
	;
	goto L12
L12:
	;
	goto L13
L13:
	;
	if l3|v10|(v26|int64(0)) == int64(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v74 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	if int64(0) <= l4&v10 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if l4 == v10 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	if l4 == v10 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v53 = base.B2i32(base.Ui64(l3) < base.Ui64(v10))
	goto L22
L21:
	;
	v53 = base.B2i32(l4 < v10)
	goto L22
L22:
	;
	if v53 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v74 = int32(-1)
	goto L4
L24:
	;
	goto L25
L25:
	;
	v74 = base.B2i32(l3^v10|(l4^v10) != int64(0))
	goto L4
L26:
	;
	v63 = base.B2i32(base.Ui64(v10) < base.Ui64(l3))
	goto L28
L27:
	;
	v63 = base.B2i32(v10 < l4)
	goto L28
L28:
	;
	if v63 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v74 = int32(-1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v70 = base.B2i32(l3^v10|(l4^v10) != int64(0))
	goto L5
L32:
	;
	v78 = l4 & int64(281474976710655)
	v82 = int32(32767)
	v83 = base.I32_wrap_i64(int64(base.Ui64(l4)>>(uint(int64(48))%64))) & v82
	if v83 != v82 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if v97 == int32(0) {
		goto L3
	} else {
		goto L42
	}
L34:
	;
	v97 = v96
	goto L33
L35:
	;
	if v83 != 0 {
		v96 = int32(4)
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v96 = base.B2i32(l3|v78 == int64(0))
	goto L34
L38:
	;
	if l3|v78 == int64(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v92 = int32(2)
	goto L41
L40:
	;
	v92 = int32(3)
	goto L41
L41:
	;
	v97 = v92
	goto L33
L42:
	;
	v102 = base.I32_wrap_i64(int64(base.Ui64(l2) >> (uint(int64(48)) % 64)))
	v103 = int32(32767)
	v104 = v102 & v103
	if v104 != v103 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	goto L3
L44:
	;
	if v178 <= int32(0) {
		goto L72
	} else {
		goto L73
	}
L45:
	;
	v178 = v174
	goto L44
L46:
	;
	v135 = base.B2i32(l1 != int64(0))
	goto L48
L47:
	;
	v135 = base.B2i32(base.Ui64(v131) < base.Ui64(v130))
	goto L48
L48:
	;
	if v135 != 0 {
		v174 = v126
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v139 = v122 & int64(9223372036854775807)
	v140 = int64(9223090561878065152)
	if v139 == v140 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v144 = base.B2i32(l3 != int64(0))
	goto L52
L51:
	;
	v144 = base.B2i32(base.Ui64(v140) < base.Ui64(v139))
	goto L52
L52:
	;
	if v144 != 0 {
		v174 = v126
		goto L45
	} else {
		goto L53
	}
L53:
	;
	if l1|l3|(v130|v139) == int64(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v178 = int32(0)
	goto L44
L55:
	;
	goto L56
L56:
	;
	if int64(0) <= v120&v122 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v120 == v122 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	if v120 == v122 {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	v157 = base.B2i32(base.Ui64(l1) < base.Ui64(l3))
	goto L62
L61:
	;
	v157 = base.B2i32(v120 < v122)
	goto L62
L62:
	;
	if v157 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v178 = int32(-1)
	goto L44
L64:
	;
	goto L65
L65:
	;
	v178 = base.B2i32(l1^l3|(v120^v122) != int64(0))
	goto L44
L66:
	;
	v167 = base.B2i32(base.Ui64(l3) < base.Ui64(l1))
	goto L68
L67:
	;
	v167 = base.B2i32(v122 < v120)
	goto L68
L68:
	;
	if v167 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v178 = int32(-1)
	goto L44
L70:
	;
	goto L71
L71:
	;
	v174 = base.B2i32(l1^l3|(v120^v122) != int64(0))
	goto L45
L72:
	;
	v184 = int32(1)
	v188 = v120 & int64(9223372036854775807)
	v189 = int64(9223090561878065152)
	if v188 == v189 {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	goto L74
L74:
	;
	v248 = base.I32_wrap_i64(int64(base.Ui64(l4)>>(uint(int64(48))%64))) & int32(32767)
	if v104 != 0 {
		goto L106
	} else {
		goto L107
	}
L75:
	;
	if v236 != 0 {
		goto L103
	} else {
		goto L104
	}
L76:
	;
	v236 = v232
	goto L75
L77:
	;
	v193 = base.B2i32(l1 != int64(0))
	goto L79
L78:
	;
	v193 = base.B2i32(base.Ui64(v189) < base.Ui64(v188))
	goto L79
L79:
	;
	if v193 != 0 {
		v232 = v184
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v197 = v122 & int64(9223372036854775807)
	v198 = int64(9223090561878065152)
	if v197 == v198 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v202 = base.B2i32(l3 != int64(0))
	goto L83
L82:
	;
	v202 = base.B2i32(base.Ui64(v198) < base.Ui64(v197))
	goto L83
L83:
	;
	if v202 != 0 {
		v232 = v184
		goto L76
	} else {
		goto L84
	}
L84:
	;
	if l1|l3|(v188|v197) == int64(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v236 = int32(0)
	goto L75
L86:
	;
	goto L87
L87:
	;
	if int64(0) <= v120&v122 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v120 == v122 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if v120 == v122 {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	v215 = base.B2i32(base.Ui64(l1) < base.Ui64(l3))
	goto L93
L92:
	;
	v215 = base.B2i32(v120 < v122)
	goto L93
L93:
	;
	if v215 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v236 = int32(-1)
	goto L75
L95:
	;
	goto L96
L96:
	;
	v236 = base.B2i32(l1^l3|(v120^v122) != int64(0))
	goto L75
L97:
	;
	v225 = base.B2i32(base.Ui64(l3) < base.Ui64(l1))
	goto L99
L98:
	;
	v225 = base.B2i32(v122 < v120)
	goto L99
L99:
	;
	if v225 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v236 = int32(-1)
	goto L75
L101:
	;
	goto L102
L102:
	;
	v232 = base.B2i32(l1^l3|(v120^v122) != int64(0))
	goto L76
L103:
	;
	v433 = l2
	v435 = l1
	goto L1
L104:
	;
	goto L105
L105:
	;
	v239 = int64(0)
	F___multf3(m, v15+int32(112), l1, l2, v239, v239)
	mBase = m.M
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v15)+120))
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	v433 = v242
	v435 = v243
	goto L1
L106:
	;
	v261 = v104
	v262 = v120
	v263 = l1
	goto L108
L107:
	;
	F___multf3(m, v15+int32(96), l1, v120, int64(0), int64(4645181540655955968))
	mBase = m.M
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v15)+104))
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v15)+96))
	v261 = base.I32_wrap_i64(int64(base.Ui64(v254)>>(uint(int64(48))%64))) - int32(120)
	v262 = v254
	v263 = v260
	goto L108
L108:
	;
	if v248 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F___multf3(m, v15+int32(80), l3, v122, int64(0), int64(4645181540655955968))
	mBase = m.M
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v15)+80))
	v278 = v277
	v279 = base.I32_wrap_i64(int64(base.Ui64(v271)>>(uint(int64(48))%64))) - int32(120)
	v280 = v271
	goto L111
L110:
	;
	v278 = l3
	v279 = v248
	v280 = v122
	goto L111
L111:
	;
	v281 = int64(281474976710655)
	v283 = int64(281474976710656)
	v284 = v280&v281 | v283
	v288 = v262&v281 | v283
	if v279 < v261 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v294 = v263
	v296 = v261
	v300 = v288
	goto L115
L113:
	;
	v340 = v263
	v342 = v261
	v346 = v288
	goto L114
L114:
	;
	v351 = v346 - v284 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v340) < base.Ui64(v278)))
	if v351 < int64(0) {
		goto L126
	} else {
		goto L127
	}
L115:
	;
	v305 = v300 - v284 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v294) < base.Ui64(v278)))
	if int64(0) <= v305 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v340 = v332
	v342 = v279
	v346 = v330
	goto L114
L117:
	;
	v332 = v329 << (uint(int64(1)) % 64)
	v334 = v296 - int32(1)
	if v279 < v334 {
		v294 = v332
		v296 = v334
		v300 = v330
		goto L115
	} else {
		goto L124
	}
L118:
	;
	v308 = v294 - v278
	if v305|v308 == int64(0) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v329 = v294
	v330 = v300<<(uint(int64(1))%64) | int64(base.Ui64(v294)>>(uint(int64(63))%64))
	goto L117
L121:
	;
	v314 = int64(0)
	F___multf3(m, v15+int32(32), l1, l2, v314, v314)
	mBase = m.M
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v15)+40))
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	v433 = v317
	v435 = v318
	goto L1
L122:
	;
	goto L123
L123:
	;
	v329 = v308
	v330 = v305<<(uint(int64(1))%64) | int64(base.Ui64(v308)>>(uint(int64(63))%64))
	goto L117
L124:
	;
	goto L116
L125:
	;
	if base.Ui64(v366) <= base.Ui64(int64(281474976710655)) {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	v365 = v340
	v366 = v346
	goto L125
L127:
	;
	goto L128
L128:
	;
	v354 = v340 - v278
	if v351|v354 != int64(0) {
		v365 = v354
		v366 = v351
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v360 = int64(0)
	F___multf3(m, v15+int32(48), l1, l2, v360, v360)
	mBase = m.M
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v15)+56))
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	v433 = v363
	v435 = v364
	goto L1
L130:
	;
	v373 = v365
	v375 = v342
	v378 = v366
	goto L133
L131:
	;
	v396 = v365
	v398 = v342
	v401 = v366
	goto L132
L132:
	;
	v405 = v102 & int32(32768)
	if v398 <= int32(0) {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v384 = v375 - int32(1)
	v385 = int64(1)
	v386 = v373 << (uint(v385) % 64)
	v389 = int64(base.Ui64(v373)>>(uint(int64(63))%64)) | v378<<(uint(v385)%64)
	if base.Ui64(v389) < base.Ui64(int64(281474976710656)) {
		v373 = v386
		v375 = v384
		v378 = v389
		goto L133
	} else {
		goto L135
	}
L134:
	;
	v396 = v386
	v398 = v384
	v401 = v389
	goto L132
L135:
	;
	goto L134
L136:
	;
	F___multf3(m, v15-int32(-64), v396, v401&int64(281474976710655)|base.I64_extend_i32_u(v398+int32(120)|v405)<<(uint(int64(48))%64), int64(0), int64(4577627546245398528))
	mBase = m.M
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v15)+72))
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v15)+64))
	v433 = v422
	v435 = v423
	goto L1
L137:
	;
	goto L138
L138:
	;
	v433 = v401&int64(281474976710655) | base.I64_extend_i32_u(v398|v405)<<(uint(int64(48))%64)
	v435 = v396
	goto L1
}
func F_fmtint(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v64 int32
	_ = v64
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int64
	_ = v87
	var v96 int32
	_ = v96
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int64
	_ = v118
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	v10 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	v21 = int32(335886)
	switch l1 - int32(88) {
	case 0:
		v39 = int32(531155)
		v40 = int32(1)
		v41 = int32(0)
		v42 = v41
		v43 = v39
		v44 = v40
		v45 = v10
		if l6 != 0 {
			if v42 == int32(0) {
				if v44 != 0 {
					v87 = l0
					v96 = v10
					for {
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
						v111 = v96 + int32(1)
						if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
							v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
							v96 = v111
							continue
						} else {
							break
						}
						break
					}
					v156 = v111
					v161 = v45
				} else {
					v118 = l0
					v127 = v10
					for {
						v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
						v142 = v127 + int32(1)
						if base.Ui64(int64(7)) < base.Ui64(v118) {
							v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
							v127 = v142
							continue
						} else {
							break
						}
						break
					}
					v156 = v142
					v161 = v45
				}
			} else {
				v50 = l0
				v52 = v43
				v54 = v45
				v55 = v50
				v64 = v10
				for {
					v72 = int64(10)
					v73 = base.I64_div_u_s(v55, v72)
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
					*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
					v82 = v64 + int32(1)
					if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
						v55 = v73
						v64 = v82
						continue
					} else {
						break
					}
					break
				}
				v156 = v82
				v161 = v54
			}
		} else {
			if l0 != int64(0) {
				if v42 == int32(0) {
					if v44 != 0 {
						v87 = l0
						v96 = v10
						for {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
							v111 = v96 + int32(1)
							if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
								v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
								v96 = v111
								continue
							} else {
								break
							}
							break
						}
						v156 = v111
						v161 = v45
					} else {
						v118 = l0
						v127 = v10
						for {
							v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
							v142 = v127 + int32(1)
							if base.Ui64(int64(7)) < base.Ui64(v118) {
								v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
								v127 = v142
								continue
							} else {
								break
							}
							break
						}
						v156 = v142
						v161 = v45
					}
				} else {
					v50 = l0
					v52 = v43
					v54 = v45
					v55 = v50
					v64 = v10
					for {
						v72 = int64(10)
						v73 = base.I64_div_u_s(v55, v72)
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
						v82 = v64 + int32(1)
						if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
							v55 = v73
							v64 = v82
							continue
						} else {
							break
						}
						break
					}
					v156 = v82
					v161 = v54
				}
			} else {
				if l7 != 0 {
					v156 = v10
					v161 = v45
				} else {
					if v42 == int32(0) {
						if v44 != 0 {
							v87 = l0
							v96 = v10
							for {
								v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
								v111 = v96 + int32(1)
								if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
									v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
									v96 = v111
									continue
								} else {
									break
								}
								break
							}
							v156 = v111
							v161 = v45
						} else {
							v118 = l0
							v127 = v10
							for {
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
								v142 = v127 + int32(1)
								if base.Ui64(int64(7)) < base.Ui64(v118) {
									v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
									v127 = v142
									continue
								} else {
									break
								}
								break
							}
							v156 = v142
							v161 = v45
						}
					} else {
						v50 = l0
						v52 = v43
						v54 = v45
						v55 = v50
						v64 = v10
						for {
							v72 = int64(10)
							v73 = base.I64_div_u_s(v55, v72)
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
							v82 = v64 + int32(1)
							if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
								v55 = v73
								v64 = v82
								continue
							} else {
								break
							}
							break
						}
						v156 = v82
						v161 = v54
					}
				}
			}
		}
		v163 = int32(0)
		v164 = l6 - v156
		if v163 < v164 {
			v168 = v164
		} else {
			v168 = v163
		}
		v170 = l4 - (v156 + v168)
		v171 = int32(0)
		if v171 < v170 {
			v174 = v170
		} else {
			v174 = v171
		}
		if l3 != 0 {
			v176 = v163 - v174
		} else {
			v176 = v174
		}
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v176
		F_leading_pad(m, l5, v161, v19+int32(12), l8)
		mBase = m.M
		v181 = m.ExcPending
		if v181 != 0 {
			return
		} else {
			v182 = int32(0)
			if v182 < v164 {
				F_dopr_outchmulti(m, int32(48), v168, l8)
				mBase = m.M
				v188 = m.ExcPending
				if v188 != 0 {
					return
				} else {
					F_dostr(m, v182-v156+v19+int32(80), v156, l8)
					mBase = m.M
					v193 = m.ExcPending
					if v193 != 0 {
						return
					} else {
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if int32(0) <= v194 {
							m.G0 = v19 + int32(80)
							return
						} else {
							F_dopr_outchmulti(m, int32(32), int32(0)-v194, l8)
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
								return
							} else {
								m.G0 = v19 + int32(80)
								return
							}
						}
					}
				}
			} else {
				F_dostr(m, v182-v156+v19+int32(80), v156, l8)
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
					return
				} else {
					v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					if int32(0) <= v194 {
						m.G0 = v19 + int32(80)
						return
					} else {
						F_dopr_outchmulti(m, int32(32), int32(0)-v194, l8)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							m.G0 = v19 + int32(80)
							return
						}
					}
				}
			}
		}
	default:
		m.G0 = v19 + int32(80)
		return
	case 12, 17:
		if int64(0) <= l0 {
			if l2 != 0 {
				v33 = int32(43)
			} else {
				v33 = int32(0)
			}
			v42 = int32(1)
			v43 = v21
			v44 = v10
			v45 = v33
			if l6 != 0 {
				if v42 == int32(0) {
					if v44 != 0 {
						v87 = l0
						v96 = v10
						for {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
							v111 = v96 + int32(1)
							if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
								v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
								v96 = v111
								continue
							} else {
								break
							}
							break
						}
						v156 = v111
						v161 = v45
					} else {
						v118 = l0
						v127 = v10
						for {
							v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
							v142 = v127 + int32(1)
							if base.Ui64(int64(7)) < base.Ui64(v118) {
								v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
								v127 = v142
								continue
							} else {
								break
							}
							break
						}
						v156 = v142
						v161 = v45
					}
				} else {
					v50 = l0
					v52 = v43
					v54 = v45
					v55 = v50
					v64 = v10
					for {
						v72 = int64(10)
						v73 = base.I64_div_u_s(v55, v72)
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
						v82 = v64 + int32(1)
						if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
							v55 = v73
							v64 = v82
							continue
						} else {
							break
						}
						break
					}
					v156 = v82
					v161 = v54
				}
			} else {
				if l0 != int64(0) {
					if v42 == int32(0) {
						if v44 != 0 {
							v87 = l0
							v96 = v10
							for {
								v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
								v111 = v96 + int32(1)
								if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
									v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
									v96 = v111
									continue
								} else {
									break
								}
								break
							}
							v156 = v111
							v161 = v45
						} else {
							v118 = l0
							v127 = v10
							for {
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
								v142 = v127 + int32(1)
								if base.Ui64(int64(7)) < base.Ui64(v118) {
									v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
									v127 = v142
									continue
								} else {
									break
								}
								break
							}
							v156 = v142
							v161 = v45
						}
					} else {
						v50 = l0
						v52 = v43
						v54 = v45
						v55 = v50
						v64 = v10
						for {
							v72 = int64(10)
							v73 = base.I64_div_u_s(v55, v72)
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
							v82 = v64 + int32(1)
							if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
								v55 = v73
								v64 = v82
								continue
							} else {
								break
							}
							break
						}
						v156 = v82
						v161 = v54
					}
				} else {
					if l7 != 0 {
						v156 = v10
						v161 = v45
					} else {
						if v42 == int32(0) {
							if v44 != 0 {
								v87 = l0
								v96 = v10
								for {
									v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
									*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
									v111 = v96 + int32(1)
									if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
										v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
										v96 = v111
										continue
									} else {
										break
									}
									break
								}
								v156 = v111
								v161 = v45
							} else {
								v118 = l0
								v127 = v10
								for {
									v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
									*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
									v142 = v127 + int32(1)
									if base.Ui64(int64(7)) < base.Ui64(v118) {
										v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
										v127 = v142
										continue
									} else {
										break
									}
									break
								}
								v156 = v142
								v161 = v45
							}
						} else {
							v50 = l0
							v52 = v43
							v54 = v45
							v55 = v50
							v64 = v10
							for {
								v72 = int64(10)
								v73 = base.I64_div_u_s(v55, v72)
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
								v82 = v64 + int32(1)
								if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
									v55 = v73
									v64 = v82
									continue
								} else {
									break
								}
								break
							}
							v156 = v82
							v161 = v54
						}
					}
				}
			}
		} else {
			v50 = int64(0) - l0
			v52 = v21
			v54 = int32(45)
			v55 = v50
			v64 = v10
			for {
				v72 = int64(10)
				v73 = base.I64_div_u_s(v55, v72)
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
				*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
				v82 = v64 + int32(1)
				if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
					v55 = v73
					v64 = v82
					continue
				} else {
					break
				}
				break
			}
			v156 = v82
			v161 = v54
		}
		v163 = int32(0)
		v164 = l6 - v156
		if v163 < v164 {
			v168 = v164
		} else {
			v168 = v163
		}
		v170 = l4 - (v156 + v168)
		v171 = int32(0)
		if v171 < v170 {
			v174 = v170
		} else {
			v174 = v171
		}
		if l3 != 0 {
			v176 = v163 - v174
		} else {
			v176 = v174
		}
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v176
		F_leading_pad(m, l5, v161, v19+int32(12), l8)
		mBase = m.M
		v181 = m.ExcPending
		if v181 != 0 {
			return
		} else {
			v182 = int32(0)
			if v182 < v164 {
				F_dopr_outchmulti(m, int32(48), v168, l8)
				mBase = m.M
				v188 = m.ExcPending
				if v188 != 0 {
					return
				} else {
					F_dostr(m, v182-v156+v19+int32(80), v156, l8)
					mBase = m.M
					v193 = m.ExcPending
					if v193 != 0 {
						return
					} else {
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if int32(0) <= v194 {
							m.G0 = v19 + int32(80)
							return
						} else {
							F_dopr_outchmulti(m, int32(32), int32(0)-v194, l8)
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
								return
							} else {
								m.G0 = v19 + int32(80)
								return
							}
						}
					}
				}
			} else {
				F_dostr(m, v182-v156+v19+int32(80), v156, l8)
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
					return
				} else {
					v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					if int32(0) <= v194 {
						m.G0 = v19 + int32(80)
						return
					} else {
						F_dopr_outchmulti(m, int32(32), int32(0)-v194, l8)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							m.G0 = v19 + int32(80)
							return
						}
					}
				}
			}
		}
	case 23:
		v42 = v10
		v43 = v21
		v44 = v10
		v45 = v10
		if l6 != 0 {
			if v42 == int32(0) {
				if v44 != 0 {
					v87 = l0
					v96 = v10
					for {
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
						v111 = v96 + int32(1)
						if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
							v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
							v96 = v111
							continue
						} else {
							break
						}
						break
					}
					v156 = v111
					v161 = v45
				} else {
					v118 = l0
					v127 = v10
					for {
						v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
						v142 = v127 + int32(1)
						if base.Ui64(int64(7)) < base.Ui64(v118) {
							v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
							v127 = v142
							continue
						} else {
							break
						}
						break
					}
					v156 = v142
					v161 = v45
				}
			} else {
				v50 = l0
				v52 = v43
				v54 = v45
				v55 = v50
				v64 = v10
				for {
					v72 = int64(10)
					v73 = base.I64_div_u_s(v55, v72)
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
					*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
					v82 = v64 + int32(1)
					if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
						v55 = v73
						v64 = v82
						continue
					} else {
						break
					}
					break
				}
				v156 = v82
				v161 = v54
			}
		} else {
			if l0 != int64(0) {
				if v42 == int32(0) {
					if v44 != 0 {
						v87 = l0
						v96 = v10
						for {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
							v111 = v96 + int32(1)
							if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
								v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
								v96 = v111
								continue
							} else {
								break
							}
							break
						}
						v156 = v111
						v161 = v45
					} else {
						v118 = l0
						v127 = v10
						for {
							v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
							v142 = v127 + int32(1)
							if base.Ui64(int64(7)) < base.Ui64(v118) {
								v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
								v127 = v142
								continue
							} else {
								break
							}
							break
						}
						v156 = v142
						v161 = v45
					}
				} else {
					v50 = l0
					v52 = v43
					v54 = v45
					v55 = v50
					v64 = v10
					for {
						v72 = int64(10)
						v73 = base.I64_div_u_s(v55, v72)
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
						v82 = v64 + int32(1)
						if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
							v55 = v73
							v64 = v82
							continue
						} else {
							break
						}
						break
					}
					v156 = v82
					v161 = v54
				}
			} else {
				if l7 != 0 {
					v156 = v10
					v161 = v45
				} else {
					if v42 == int32(0) {
						if v44 != 0 {
							v87 = l0
							v96 = v10
							for {
								v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
								v111 = v96 + int32(1)
								if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
									v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
									v96 = v111
									continue
								} else {
									break
								}
								break
							}
							v156 = v111
							v161 = v45
						} else {
							v118 = l0
							v127 = v10
							for {
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
								v142 = v127 + int32(1)
								if base.Ui64(int64(7)) < base.Ui64(v118) {
									v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
									v127 = v142
									continue
								} else {
									break
								}
								break
							}
							v156 = v142
							v161 = v45
						}
					} else {
						v50 = l0
						v52 = v43
						v54 = v45
						v55 = v50
						v64 = v10
						for {
							v72 = int64(10)
							v73 = base.I64_div_u_s(v55, v72)
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
							v82 = v64 + int32(1)
							if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
								v55 = v73
								v64 = v82
								continue
							} else {
								break
							}
							break
						}
						v156 = v82
						v161 = v54
					}
				}
			}
		}
		v163 = int32(0)
		v164 = l6 - v156
		if v163 < v164 {
			v168 = v164
		} else {
			v168 = v163
		}
		v170 = l4 - (v156 + v168)
		v171 = int32(0)
		if v171 < v170 {
			v174 = v170
		} else {
			v174 = v171
		}
		if l3 != 0 {
			v176 = v163 - v174
		} else {
			v176 = v174
		}
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v176
		F_leading_pad(m, l5, v161, v19+int32(12), l8)
		mBase = m.M
		v181 = m.ExcPending
		if v181 != 0 {
			return
		} else {
			v182 = int32(0)
			if v182 < v164 {
				F_dopr_outchmulti(m, int32(48), v168, l8)
				mBase = m.M
				v188 = m.ExcPending
				if v188 != 0 {
					return
				} else {
					F_dostr(m, v182-v156+v19+int32(80), v156, l8)
					mBase = m.M
					v193 = m.ExcPending
					if v193 != 0 {
						return
					} else {
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if int32(0) <= v194 {
							m.G0 = v19 + int32(80)
							return
						} else {
							F_dopr_outchmulti(m, int32(32), int32(0)-v194, l8)
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
								return
							} else {
								m.G0 = v19 + int32(80)
								return
							}
						}
					}
				}
			} else {
				F_dostr(m, v182-v156+v19+int32(80), v156, l8)
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
					return
				} else {
					v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					if int32(0) <= v194 {
						m.G0 = v19 + int32(80)
						return
					} else {
						F_dopr_outchmulti(m, int32(32), int32(0)-v194, l8)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							m.G0 = v19 + int32(80)
							return
						}
					}
				}
			}
		}
	case 29:
		v39 = v21
		v40 = v10
		v41 = int32(1)
		v42 = v41
		v43 = v39
		v44 = v40
		v45 = v10
		if l6 != 0 {
			if v42 == int32(0) {
				if v44 != 0 {
					v87 = l0
					v96 = v10
					for {
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
						v111 = v96 + int32(1)
						if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
							v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
							v96 = v111
							continue
						} else {
							break
						}
						break
					}
					v156 = v111
					v161 = v45
				} else {
					v118 = l0
					v127 = v10
					for {
						v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
						v142 = v127 + int32(1)
						if base.Ui64(int64(7)) < base.Ui64(v118) {
							v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
							v127 = v142
							continue
						} else {
							break
						}
						break
					}
					v156 = v142
					v161 = v45
				}
			} else {
				v50 = l0
				v52 = v43
				v54 = v45
				v55 = v50
				v64 = v10
				for {
					v72 = int64(10)
					v73 = base.I64_div_u_s(v55, v72)
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
					*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
					v82 = v64 + int32(1)
					if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
						v55 = v73
						v64 = v82
						continue
					} else {
						break
					}
					break
				}
				v156 = v82
				v161 = v54
			}
		} else {
			if l0 != int64(0) {
				if v42 == int32(0) {
					if v44 != 0 {
						v87 = l0
						v96 = v10
						for {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
							v111 = v96 + int32(1)
							if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
								v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
								v96 = v111
								continue
							} else {
								break
							}
							break
						}
						v156 = v111
						v161 = v45
					} else {
						v118 = l0
						v127 = v10
						for {
							v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
							v142 = v127 + int32(1)
							if base.Ui64(int64(7)) < base.Ui64(v118) {
								v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
								v127 = v142
								continue
							} else {
								break
							}
							break
						}
						v156 = v142
						v161 = v45
					}
				} else {
					v50 = l0
					v52 = v43
					v54 = v45
					v55 = v50
					v64 = v10
					for {
						v72 = int64(10)
						v73 = base.I64_div_u_s(v55, v72)
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
						v82 = v64 + int32(1)
						if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
							v55 = v73
							v64 = v82
							continue
						} else {
							break
						}
						break
					}
					v156 = v82
					v161 = v54
				}
			} else {
				if l7 != 0 {
					v156 = v10
					v161 = v45
				} else {
					if v42 == int32(0) {
						if v44 != 0 {
							v87 = l0
							v96 = v10
							for {
								v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
								v111 = v96 + int32(1)
								if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
									v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
									v96 = v111
									continue
								} else {
									break
								}
								break
							}
							v156 = v111
							v161 = v45
						} else {
							v118 = l0
							v127 = v10
							for {
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
								v142 = v127 + int32(1)
								if base.Ui64(int64(7)) < base.Ui64(v118) {
									v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
									v127 = v142
									continue
								} else {
									break
								}
								break
							}
							v156 = v142
							v161 = v45
						}
					} else {
						v50 = l0
						v52 = v43
						v54 = v45
						v55 = v50
						v64 = v10
						for {
							v72 = int64(10)
							v73 = base.I64_div_u_s(v55, v72)
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
							v82 = v64 + int32(1)
							if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
								v55 = v73
								v64 = v82
								continue
							} else {
								break
							}
							break
						}
						v156 = v82
						v161 = v54
					}
				}
			}
		}
		v163 = int32(0)
		v164 = l6 - v156
		if v163 < v164 {
			v168 = v164
		} else {
			v168 = v163
		}
		v170 = l4 - (v156 + v168)
		v171 = int32(0)
		if v171 < v170 {
			v174 = v170
		} else {
			v174 = v171
		}
		if l3 != 0 {
			v176 = v163 - v174
		} else {
			v176 = v174
		}
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v176
		F_leading_pad(m, l5, v161, v19+int32(12), l8)
		mBase = m.M
		v181 = m.ExcPending
		if v181 != 0 {
			return
		} else {
			v182 = int32(0)
			if v182 < v164 {
				F_dopr_outchmulti(m, int32(48), v168, l8)
				mBase = m.M
				v188 = m.ExcPending
				if v188 != 0 {
					return
				} else {
					F_dostr(m, v182-v156+v19+int32(80), v156, l8)
					mBase = m.M
					v193 = m.ExcPending
					if v193 != 0 {
						return
					} else {
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if int32(0) <= v194 {
							m.G0 = v19 + int32(80)
							return
						} else {
							F_dopr_outchmulti(m, int32(32), int32(0)-v194, l8)
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
								return
							} else {
								m.G0 = v19 + int32(80)
								return
							}
						}
					}
				}
			} else {
				F_dostr(m, v182-v156+v19+int32(80), v156, l8)
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
					return
				} else {
					v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					if int32(0) <= v194 {
						m.G0 = v19 + int32(80)
						return
					} else {
						F_dopr_outchmulti(m, int32(32), int32(0)-v194, l8)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							m.G0 = v19 + int32(80)
							return
						}
					}
				}
			}
		}
	case 32:
		v39 = v21
		v40 = int32(1)
		v41 = int32(0)
		v42 = v41
		v43 = v39
		v44 = v40
		v45 = v10
		if l6 != 0 {
			if v42 == int32(0) {
				if v44 != 0 {
					v87 = l0
					v96 = v10
					for {
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
						v111 = v96 + int32(1)
						if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
							v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
							v96 = v111
							continue
						} else {
							break
						}
						break
					}
					v156 = v111
					v161 = v45
				} else {
					v118 = l0
					v127 = v10
					for {
						v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
						v142 = v127 + int32(1)
						if base.Ui64(int64(7)) < base.Ui64(v118) {
							v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
							v127 = v142
							continue
						} else {
							break
						}
						break
					}
					v156 = v142
					v161 = v45
				}
			} else {
				v50 = l0
				v52 = v43
				v54 = v45
				v55 = v50
				v64 = v10
				for {
					v72 = int64(10)
					v73 = base.I64_div_u_s(v55, v72)
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
					*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
					v82 = v64 + int32(1)
					if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
						v55 = v73
						v64 = v82
						continue
					} else {
						break
					}
					break
				}
				v156 = v82
				v161 = v54
			}
		} else {
			if l0 != int64(0) {
				if v42 == int32(0) {
					if v44 != 0 {
						v87 = l0
						v96 = v10
						for {
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
							v111 = v96 + int32(1)
							if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
								v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
								v96 = v111
								continue
							} else {
								break
							}
							break
						}
						v156 = v111
						v161 = v45
					} else {
						v118 = l0
						v127 = v10
						for {
							v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
							v142 = v127 + int32(1)
							if base.Ui64(int64(7)) < base.Ui64(v118) {
								v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
								v127 = v142
								continue
							} else {
								break
							}
							break
						}
						v156 = v142
						v161 = v45
					}
				} else {
					v50 = l0
					v52 = v43
					v54 = v45
					v55 = v50
					v64 = v10
					for {
						v72 = int64(10)
						v73 = base.I64_div_u_s(v55, v72)
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
						v82 = v64 + int32(1)
						if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
							v55 = v73
							v64 = v82
							continue
						} else {
							break
						}
						break
					}
					v156 = v82
					v161 = v54
				}
			} else {
				if l7 != 0 {
					v156 = v10
					v161 = v45
				} else {
					if v42 == int32(0) {
						if v44 != 0 {
							v87 = l0
							v96 = v10
							for {
								v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v87)&int32(15)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v19-v96)+79)) = uint8(v108)
								v111 = v96 + int32(1)
								if base.B2i32(base.Ui64(v87) < base.Ui64(int64(16))) == int32(0) {
									v87 = int64(base.Ui64(v87) >> (uint(int64(4)) % 64))
									v96 = v111
									continue
								} else {
									break
								}
								break
							}
							v156 = v111
							v161 = v45
						} else {
							v118 = l0
							v127 = v10
							for {
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+base.I32_wrap_i64(v118)&int32(7)))))
								*(*uint8)(unsafe.Add(mBase, uint32(v19-v127)+79)) = uint8(v139)
								v142 = v127 + int32(1)
								if base.Ui64(int64(7)) < base.Ui64(v118) {
									v118 = int64(base.Ui64(v118) >> (uint(int64(3)) % 64))
									v127 = v142
									continue
								} else {
									break
								}
								break
							}
							v156 = v142
							v161 = v45
						}
					} else {
						v50 = l0
						v52 = v43
						v54 = v45
						v55 = v50
						v64 = v10
						for {
							v72 = int64(10)
							v73 = base.I64_div_u_s(v55, v72)
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+base.I32_wrap_i64(v55-v73*v72)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v64)+79)) = uint8(v79)
							v82 = v64 + int32(1)
							if base.B2i32(base.Ui64(v55) < base.Ui64(v72)) == int32(0) {
								v55 = v73
								v64 = v82
								continue
							} else {
								break
							}
							break
						}
						v156 = v82
						v161 = v54
					}
				}
			}
		}
		v163 = int32(0)
		v164 = l6 - v156
		if v163 < v164 {
			v168 = v164
		} else {
			v168 = v163
		}
		v170 = l4 - (v156 + v168)
		v171 = int32(0)
		if v171 < v170 {
			v174 = v170
		} else {
			v174 = v171
		}
		if l3 != 0 {
			v176 = v163 - v174
		} else {
			v176 = v174
		}
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v176
		F_leading_pad(m, l5, v161, v19+int32(12), l8)
		mBase = m.M
		v181 = m.ExcPending
		if v181 != 0 {
			return
		} else {
			v182 = int32(0)
			if v182 < v164 {
				F_dopr_outchmulti(m, int32(48), v168, l8)
				mBase = m.M
				v188 = m.ExcPending
				if v188 != 0 {
					return
				} else {
					F_dostr(m, v182-v156+v19+int32(80), v156, l8)
					mBase = m.M
					v193 = m.ExcPending
					if v193 != 0 {
						return
					} else {
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if int32(0) <= v194 {
							m.G0 = v19 + int32(80)
							return
						} else {
							F_dopr_outchmulti(m, int32(32), int32(0)-v194, l8)
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
								return
							} else {
								m.G0 = v19 + int32(80)
								return
							}
						}
					}
				}
			} else {
				F_dostr(m, v182-v156+v19+int32(80), v156, l8)
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
					return
				} else {
					v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					if int32(0) <= v194 {
						m.G0 = v19 + int32(80)
						return
					} else {
						F_dopr_outchmulti(m, int32(32), int32(0)-v194, l8)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							m.G0 = v19 + int32(80)
							return
						}
					}
				}
			}
		}
	}
}
func F_fputs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	v4 = int32(0)
	if l0&int32(3) == v4 {
		v29 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v63 = F_fwrite(m, l0, int32(1), v62, l1)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v62 = v54 - l0
	goto L1
L3:
	;
	v33 = v29
	goto L12
L4:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v13 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v62 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v18 = l0
	goto L8
L8:
	;
	v22 = v18 + int32(1)
	if v22&int32(3) == int32(0) {
		v29 = v22
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v54 = v22
	goto L2
L10:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v27 != 0 {
		v18 = v22
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v42 = int32(-2139062144)
	if (int32(16843008)-v39|v39)&v42 == v42 {
		v33 = v33 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v48 = v33
	goto L15
L14:
	;
	goto L13
L15:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v52 != 0 {
		v48 = v48 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v54 = v48
	goto L2
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	if v63 != v62 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v68 = int32(-1)
	goto L22
L21:
	;
	v68 = v4
	goto L22
L22:
	;
	return v68
}
func F_freearc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int64
	_ = v69
	var v75 int32
	_ = v75
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v10 < int32(0) {
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v15 = v13 - int32(97)
		if base.Ui32(int32(17)) < base.Ui32(v15) {
		} else {
			if int32(1)<<(uint(v15)%32)&int32(163841) == int32(0) {
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				if v24 != 0 {
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					if v25 == int32(0) {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						*(*int32)(unsafe.Add(mBase, uint32(v29+v10*int32(24))+12)) = v33
						v37 = v33
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v35
						v37 = v35
					}
					if v37 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v37)+36)) = v25
					} else {
					}
					*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = int64(0)
				}
			}
		}
	}
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v44 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v43
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v43
	}
	if v43 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v44
	} else {
	}
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v50 - int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v55 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v54
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v54
	}
	v61 = l1 + int32(8)
	if v54 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v54)+28)) = v55
	} else {
	}
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v63 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v69 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v61)+16)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v61)+8)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v61))) = v69
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	return
}
func F_freestate_cluster(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+152))
		if v6 != 0 {
			v9 = v6
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			F_ExecDropSingleTupleTableSlot(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
				F_FreeExecutorState(m, v13)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v7 = F_MakePerTupleExprContext(m, v5)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				v9 = v7
				v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				F_ExecDropSingleTupleTableSlot(m, v10)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
					F_FreeExecutorState(m, v13)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		return
	}
}
func F_fsm_vacuum_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int64
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v20
	v25 = F_fsm_readbuf(m, l0, v18+int32(8), v6)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(32)
	return v317
L2:
	;
	return int32(0)
L3:
	;
	if v25 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v31)
	v317 = v6
	goto L1
L5:
	;
	goto L6
L6:
	;
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v33)
	if v25 < v33 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v53 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+(v25^int32(-1))<<(uint(int32(2))%32))))
	v52 = v44
	goto L7
L9:
	;
	goto L10
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v52 = v46 + v25<<(uint(int32(13))%32) + int32(-8192)
	goto L7
L11:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+28)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = int32(0)
	F_ReleaseBuffer(m, v25)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L2
	} else {
		goto L84
	}
L12:
	;
	v56 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+31)) = uint8(v56)
	v59 = int32(4069)
	v60 = base.I32_div_u_s(l2, v59)
	v61 = int32(1)
	v62 = l3 - v61
	v64 = base.I32_div_u_s(v62, v59)
	if v53 != v61 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v67 = int32(16556761)
	v68 = base.I32_div_u_s(l2, v67)
	v70 = base.I32_div_u_s(v62, v67)
	v77 = v56
	v78 = v68
	v80 = v70
	goto L16
L14:
	;
	v99 = v60
	v106 = v6
	v107 = v64
	v108 = v6
	goto L15
L15:
	;
	v109 = int32(0)
	if v53&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v88 = int32(0)
	v92 = v77 + int32(2)
	if v92 != v53&int32(2147483646) {
		v77 = v92
		v78 = v88
		v80 = v88
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v99 = v88
	v106 = v80
	v107 = v88
	v108 = v78
	goto L15
L18:
	;
	goto L17
L19:
	;
	v113 = int32(4069)
	v114 = base.I32_div_u_s(v99, v113)
	v116 = base.I32_div_u_s(v107, v113)
	v117 = v116
	v118 = v114
	v119 = v107
	v120 = v99
	goto L21
L20:
	;
	v117 = v109
	v118 = v109
	v119 = v106
	v120 = v108
	goto L21
L21:
	;
	v121 = int32(4069)
	v122 = base.I32_rem_u_s(v120, v121)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v125 < v118 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v127 = v121
	goto L24
L23:
	;
	v127 = int32(0)
	goto L24
L24:
	;
	if v125 == v118 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v129 = v122
	goto L27
L26:
	;
	v129 = v127
	goto L27
L27:
	;
	v131 = base.I32_rem_u_s(v119, int32(4069))
	if v125 < v117 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v135 = int32(4068)
	goto L30
L29:
	;
	v135 = int32(-1)
	goto L30
L30:
	;
	if v125 == v117 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v137 = v131
	goto L33
L32:
	;
	v137 = v135
	goto L33
L33:
	;
	if v137 < v129 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	v148 = v129
	goto L35
L35:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v159 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L11
L37:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v162 = int32(0)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+31)))
	if v163 == v162 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v53 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v148 + v125*int32(4069)
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v169
	v173 = F_fsm_vacuum_page(m, l0, v18, l2, l3, v18+int32(31))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L44
	}
L42:
	;
	v175 = v162
	goto L43
L43:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v148)+uint32(_consts[95]))))
	goto L45
L44:
	;
	v175 = v173
	goto L43
L45:
	;
	if v179 != v175 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_LockBuffer(m, v25, int32(2))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v148 != v137 {
		v148 = v148 + int32(1)
		goto L35
	} else {
		goto L83
	}
L49:
	;
	v188 = v52 + int32(28)
	v190 = v148 + int32(4095)
	v191 = v188 + v190
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	if v192 != v175 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	F_MarkBufferDirtyHint(m, v25, int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L2
	} else {
		goto L81
	}
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v175)
	v202 = v190
	goto L54
L52:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if base.Ui32(v194) < base.Ui32(v175) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	v204 = int32(1)
	v205 = v202 - v204
	v206 = int32(2)
	v207 = base.I32_div_s(v205, v206)
	v209 = v207 << (uint(v204) % 32)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v209)+1)))
	v213 = v209 + v206
	if base.Ui32(v213) <= base.Ui32(int32(8163)) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if base.Ui32(v232) < base.Ui32(v175) {
		goto L66
	} else {
		goto L67
	}
L56:
	;
	v217 = v211 & int32(255)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v213))))
	if base.Ui32(v219) < base.Ui32(v217) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v222 = v211
	goto L58
L58:
	;
	v224 = v188 + v207
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v225 != v222&int32(255) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v221 = v217
	goto L61
L60:
	;
	v221 = v219
	goto L61
L61:
	;
	v222 = v221
	goto L58
L62:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v222)
	if int32(1) < v205 {
		v202 = v207
		goto L54
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	goto L55
L65:
	;
	goto L64
L66:
	;
	v238 = int32(4094)
	goto L69
L67:
	;
	goto L68
L68:
	;
	goto L50
L69:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v238) {
		v260 = int32(0)
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L68
L71:
	;
	v261 = v188 + v238
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	if v262 != v260&int32(255) {
		goto L77
	} else {
		goto L78
	}
L72:
	;
	v245 = v238 << (uint(int32(1)) % 32)
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v245)+1)))
	if v238 == int32(4081) {
		v260 = v247
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v251 = v247 & int32(255)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+(v245+int32(2))))))
	if base.Ui32(v255) < base.Ui32(v251) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v257 = v251
	goto L76
L75:
	;
	v257 = v255
	goto L76
L76:
	;
	v260 = v257
	goto L71
L77:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v260)
	goto L79
L78:
	;
	goto L79
L79:
	;
	if v238 != 0 {
		v238 = v238 - int32(1)
		goto L69
	} else {
		goto L80
	}
L80:
	;
	goto L70
L81:
	;
	F_LockBuffer(m, v25, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	goto L48
L83:
	;
	goto L36
L84:
	;
	v317 = v307
	goto L1
}
func F_ftoi2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float32
	_ = v3
	var v4 float32
	_ = v4
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v3 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = base.F32_nearest(v3)
	if base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v4)&int32(2147483647)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(397742), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(488100), int32(1328), int32(547018))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
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
		if base.F32_ge(v4, float32(-32768)) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(397742), int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(488100), int32(1328), int32(547018))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
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
			if base.F32_lt(v4, float32(32768)) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(397742), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(488100), int32(1328), int32(547018))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
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
				if base.F32_lt(base.F32_abs(v4), float32(2.1474836e+09)) != 0 {
					v21 = base.I32_trunc_f32_s(v4)
					return v21
				} else {
					return int32(-2147483648)
				}
			}
		}
	}
}
func F_ftoi4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float32
	_ = v3
	var v4 float32
	_ = v4
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v3 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = base.F32_nearest(v3)
	if base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v4)&int32(2147483647)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(398054), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(488100), int32(1303), int32(545509))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
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
		if base.F32_ge(v4, float32(-2.1474836e+09)) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(398054), int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(488100), int32(1303), int32(545509))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
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
			if base.F32_lt(v4, float32(2.1474836e+09)) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(398054), int32(0))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(488100), int32(1303), int32(545509))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
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
				if base.F32_lt(base.F32_abs(v4), float32(2.1474836e+09)) != 0 {
					v21 = base.I32_trunc_f32_s(v4)
					return v21
				} else {
					return int32(-2147483648)
				}
			}
		}
	}
}
func F_ftoi8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float32
	_ = v3
	var v4 float32
	_ = v4
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v3 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = base.F32_nearest(v3)
	if base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v4)&int32(2147483647)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(397764), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494690), int32(1347), int32(542829))
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
	} else {
		if base.F32_ge(v4, float32(-9.223372e+18)) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(397764), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(494690), int32(1347), int32(542829))
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
		} else {
			if base.F32_lt(v4, float32(9.223372e+18)) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(397764), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494690), int32(1347), int32(542829))
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
			} else {
				if base.F32_lt(base.F32_abs(v4), float32(9.223372e+18)) != 0 {
					v21 = base.I64_trunc_f32_s(v4)
					v22 = F_Int64GetDatum(m, v21)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						return v22
					}
				} else {
					v28 = F_Int64GetDatum(m, int64(-9223372036854775807-1))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						return v28
					}
				}
			}
		}
	}
}
