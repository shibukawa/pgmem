package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int64
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
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
	return v334
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
	v334 = v6
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
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_vacuum_page[0]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+(v25^int32(-1))<<(uint(int32(2))%32))))
	v52 = v44
	goto L7
L9:
	;
	goto L10
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_vacuum_page[1]))
	v52 = v46 + v25<<(uint(int32(13))%32) + int32(-8192)
	goto L7
L11:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+28)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = int32(0)
	F_ReleaseBuffer(m, v25)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L2
	} else {
		goto L83
	}
L12:
	;
	v56 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+31)) = uint8(v56)
	v58 = int32(4069)
	v59 = base.I32_div_u_s(l2, v58)
	v60 = int32(1)
	v61 = l3 - v60
	v63 = base.I32_div_u_s(v61, v58)
	if v53 != v60 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v141 < v128 {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v68 = int32(16556761)
	v69 = base.I32_div_u_s(l2, v68)
	v71 = base.I32_div_u_s(v61, v68)
	v78 = v69
	v79 = v6
	v81 = v71
	goto L17
L15:
	;
	v102 = v59
	v105 = v63
	goto L16
L16:
	;
	v113 = int32(4069)
	v114 = base.I32_div_u_s(v105, v113)
	v119 = base.I32_div_u_s(v102, v113)
	v128 = v119
	v132 = v105 - v114*v113
	v135 = v102 - v119*v113
	v138 = v114
	goto L13
L17:
	;
	v89 = int32(0)
	v92 = v79 + int32(2)
	if v92 != v53&int32(2147483646) {
		v78 = v89
		v79 = v92
		v81 = v89
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v94 = int32(0)
	if v53&int32(1) == v94 {
		v128 = v94
		v132 = v81
		v135 = v78
		v138 = v94
		goto L13
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v102 = v89
	v105 = v89
	goto L16
L21:
	;
	v143 = int32(4069)
	goto L23
L22:
	;
	v143 = int32(0)
	goto L23
L23:
	;
	if v141 == v128 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v145 = v135
	goto L26
L25:
	;
	v145 = v143
	goto L26
L26:
	;
	if v141 < v138 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v149 = int32(4068)
	goto L29
L28:
	;
	v149 = int32(-1)
	goto L29
L29:
	;
	if v141 == v138 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v151 = v132
	goto L32
L31:
	;
	v151 = v149
	goto L32
L32:
	;
	if v151 < v145 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	v162 = v145
	goto L34
L34:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_fsm_vacuum_page[2]))
	if v173 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L11
L36:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v176 = int32(0)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+31)))
	if v177 == v176 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v53 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v141*int32(4069) + v162
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v183
	v187 = F_fsm_vacuum_page(m, l0, v18, l2, l3, v18+int32(31))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L43
	}
L41:
	;
	v189 = v176
	goto L42
L42:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v162)+uint32(_c_F_fsm_vacuum_page[3]))))
	goto L44
L43:
	;
	v189 = v187
	goto L42
L44:
	;
	if v193 != v189 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_LockBuffer(m, v25, int32(2))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L2
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v162 != v151 {
		v162 = v162 + int32(1)
		goto L34
	} else {
		goto L82
	}
L48:
	;
	v203 = v52 + int32(28)
	v205 = v162 + int32(4095)
	v206 = v203 + v205
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v207 != v189 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	F_MarkBufferDirtyHint(m, v25, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L2
	} else {
		goto L80
	}
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v206))) = uint8(v189)
	v216 = v205
	goto L53
L51:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if base.Ui32(v209) < base.Ui32(v189) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	v220 = int32(1)
	v221 = v216 - v220
	v222 = int32(2)
	v223 = base.I32_div_s(v221, v222)
	v225 = v223 << (uint(v220) % 32)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203+v225)+1)))
	v229 = v225 + v222
	if base.Ui32(v229) <= base.Ui32(int32(_a_F_fsm_vacuum_page_0)) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if base.Ui32(v248) < base.Ui32(v189) {
		goto L65
	} else {
		goto L66
	}
L55:
	;
	v233 = v227 & int32(255)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203+v229))))
	if base.Ui32(v235) < base.Ui32(v233) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v238 = v227
	goto L57
L57:
	;
	v240 = v223 + v203
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	if v241 != v238&int32(255) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v237 = v233
	goto L60
L59:
	;
	v237 = v235
	goto L60
L60:
	;
	v238 = v237
	goto L57
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v240))) = uint8(v238)
	if int32(1) < v221 {
		v216 = v223
		goto L53
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L54
L64:
	;
	goto L63
L65:
	;
	v254 = int32(4094)
	goto L68
L66:
	;
	goto L67
L67:
	;
	goto L49
L68:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v254) {
		v275 = int32(0)
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L67
L70:
	;
	v276 = v254 + v203
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	if v277 != v275&int32(255) {
		goto L76
	} else {
		goto L77
	}
L71:
	;
	v264 = v254 << (uint(int32(1)) % 32)
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+int32(29)+v264))))
	if v254 == int32(4081) {
		v275 = v266
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264+v203)+2)))
	if base.Ui32(v270) < base.Ui32(v266) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v272 = v266
	goto L75
L74:
	;
	v272 = v270
	goto L75
L75:
	;
	v275 = v272
	goto L70
L76:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v276))) = uint8(v275)
	goto L78
L77:
	;
	goto L78
L78:
	;
	if v254 != 0 {
		v254 = v254 - int32(1)
		goto L68
	} else {
		goto L79
	}
L79:
	;
	goto L69
L80:
	;
	F_LockBuffer(m, v25, int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	goto L47
L82:
	;
	goto L35
L83:
	;
	v334 = v324
	goto L1
}
