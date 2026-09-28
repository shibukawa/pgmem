package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
	"unsafe"
)

func F_ConditionVariableBroadcast(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
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
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[0]))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[1]))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = base.AtomicRmwXchg32(m, v11, int32(0), int32(1))
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v70 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v70 != 0 {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	F_s_lock(m, v11, int32(_a_F_ConditionVariableBroadcast_0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[0]))
	v25 = v20 + v22*int32(768)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+360))
	if v26 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	v58 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v11))), uint32(v58))
	*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[1])) = v58
	goto L3
L10:
	;
	if v43 == int32(-1) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v26*int32(768))+356)) = v37
	v42 = v26
	v43 = v37
	goto L10
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+356))
	if v29 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+356))
	if v26 != int32(-1) {
		v37 = v32
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v37 = v29
	goto L11
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v32
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+360))
	v42 = v36
	v43 = v32
	goto L10
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+356)) = int64(0)
	goto L9
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v42
	goto L17
L19:
	;
	goto L20
L20:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v49+v43*int32(768))+360)) = v42
	goto L17
L21:
	;
	F_s_lock(m, l0, int32(_a_F_ConditionVariableBroadcast_0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v74 != int32(-1) {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	goto L23
L25:
	;
	return
L26:
	;
	v313 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v313))
	v317 = v82 + int32(316)
	v321 = base.AtomicRmwOr32(m, v313, int32(_a_F_ConditionVariableBroadcast_1), v313)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	if v322 != 0 {
		goto L97
	} else {
		goto L98
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
	v135 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v135))
	v139 = v82 + int32(316)
	v143 = base.AtomicRmwOr32(m, v135, int32(_a_F_ConditionVariableBroadcast_1), v135)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v144 != 0 {
		goto L43
	} else {
		goto L44
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+360)) = v115
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v127+v115*int32(768))+356)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v114)+356)) = int32(-1)
	goto L27
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v82 = v79 + v74*int32(768)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+356))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+360))
	if v84 == int32(-1) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v121 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v121))
	return
L32:
	;
	if v83 == int32(-1) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v83
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v82)+360))
	v93 = v88
	goto L32
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79+v84*int32(768))+356)) = v83
	v93 = v84
	goto L32
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v82)+356)) = int64(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v106 == int32(-1) {
		goto L26
	} else {
		goto L40
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v93
	goto L36
L38:
	;
	goto L39
L39:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	*(*int32)(unsafe.Add(mBase, uint32(v99+v83*int32(768))+360)) = v93
	goto L36
L40:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v114 = v111 + v9*int32(768)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v115 != int32(-1) {
		goto L28
	} else {
		goto L41
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v114)+356)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	goto L27
L42:
	;
	goto L56
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = int32(1)
	v147 = int32(0)
	v150 = base.AtomicRmwOr32(m, v147, int32(_a_F_ConditionVariableBroadcast_1), v147)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v151 == v147 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	if v154 == int32(0) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[3]))
	if v158 == v154 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v160 = m.G0
	v162 = v160 - int32(16)
	m.G0 = v162
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[4]))
	if v165 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v188 = F_pgmem_kill(m, v154, int32(23))
	mBase = m.M
	goto L43
L50:
	;
	m.G0 = v162 + int32(16)
	goto L42
L51:
	;
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v162)+15)) = uint8(v168)
	goto L52
L52:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[5]))
	v176 = F_write(m, v172, v162+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v176 {
		goto L50
	} else {
		goto L54
	}
L53:
	;
	goto L50
L54:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[6]))
	if v180 == int32(27) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v202 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v202 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L25
L58:
	;
	F_s_lock(m, l0, int32(_a_F_ConditionVariableBroadcast_0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v206 == int32(-1) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L60
L62:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v246 = v245 + v9*int32(768)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+360))
	if v247 != 0 {
		goto L75
	} else {
		goto L76
	}
L63:
	;
	v240 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v215 = v212 + v206*int32(768)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+356))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215)+360))
	if v217 == int32(-1) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v216 == int32(-1) {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v216
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v215)+360))
	v226 = v221
	goto L66
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212+v217*int32(768))+356)) = v216
	v226 = v217
	goto L66
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v215)+356)) = int64(0)
	v240 = v215
	goto L62
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v226
	goto L70
L72:
	;
	goto L73
L73:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[2]))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	*(*int32)(unsafe.Add(mBase, uint32(v232+v216*int32(768))+360)) = v226
	goto L70
L74:
	;
	v252 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v252))
	if v240 == v252 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v251 = int32(1)
	goto L74
L76:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v246)+356))
	if v248 != 0 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v251 = int32(0)
	goto L74
L78:
	;
	if v251 != 0 {
		goto L56
	} else {
		goto L95
	}
L79:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[7]))
	if v240 == v258 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v261 = v240 + int32(316)
	v262 = int32(0)
	v265 = base.AtomicRmwOr32(m, v262, int32(_a_F_ConditionVariableBroadcast_1), v262)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	if v266 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L78
L82:
	;
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = int32(1)
	v269 = int32(0)
	v272 = base.AtomicRmwOr32(m, v269, int32(_a_F_ConditionVariableBroadcast_1), v269)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v273 == v269 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	if v276 == int32(0) {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[3]))
	if v280 == v276 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v282 = m.G0
	v284 = v282 - int32(16)
	m.G0 = v284
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[4]))
	if v287 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v310 = F_pgmem_kill(m, v276, int32(23))
	mBase = m.M
	goto L82
L89:
	;
	m.G0 = v284 + int32(16)
	goto L81
L90:
	;
	v290 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v284)+15)) = uint8(v290)
	goto L91
L91:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[5]))
	v298 = F_write(m, v294, v284+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v298 {
		goto L89
	} else {
		goto L93
	}
L92:
	;
	goto L89
L93:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[6]))
	if v302 == int32(27) {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	goto L57
L96:
	;
	goto L25
L97:
	;
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317))) = int32(1)
	v325 = int32(0)
	v328 = base.AtomicRmwOr32(m, v325, int32(_a_F_ConditionVariableBroadcast_1), v325)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
	if v329 == v325 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v317)+12))
	if v332 == int32(0) {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v336 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[3]))
	if v336 == v332 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v338 = m.G0
	v340 = v338 - int32(16)
	m.G0 = v340
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[4]))
	if v343 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	v366 = F_pgmem_kill(m, v332, int32(23))
	mBase = m.M
	goto L97
L104:
	;
	m.G0 = v340 + int32(16)
	goto L96
L105:
	;
	v346 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v340)+15)) = uint8(v346)
	goto L106
L106:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[5]))
	v354 = F_write(m, v350, v340+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v354 {
		goto L104
	} else {
		goto L108
	}
L107:
	;
	goto L104
L108:
	;
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariableBroadcast[6]))
	if v358 == int32(27) {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
}
func F_ConditionVariablePrepareToSleep(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[0]))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[1]))
	if v11 != 0 {
		v14 = base.AtomicRmwXchg32(m, v11, int32(0), int32(1))
		if v14 != 0 {
			F_s_lock(m, v11, int32(_a_F_ConditionVariablePrepareToSleep_0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[0]))
				v25 = v20 + v22*int32(768)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+360))
				if v26 == int32(0) {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+356))
					if v29 == int32(0) {
					} else {
						v37 = v29
						*(*int32)(unsafe.Add(mBase, uint32(v20+v26*int32(768))+356)) = v37
						v42 = v26
						v43 = v37
						if v43 == int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v42
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
							*(*int32)(unsafe.Add(mBase, uint32(v49+v43*int32(768))+360)) = v42
						}
						*(*int64)(unsafe.Add(mBase, uint32(v25)+356)) = int64(0)
					}
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+356))
					if v26 != int32(-1) {
						v37 = v32
						*(*int32)(unsafe.Add(mBase, uint32(v20+v26*int32(768))+356)) = v37
						v42 = v26
						v43 = v37
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v32
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+360))
						v42 = v36
						v43 = v32
					}
					if v43 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v42
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
						*(*int32)(unsafe.Add(mBase, uint32(v49+v43*int32(768))+360)) = v42
					}
					*(*int64)(unsafe.Add(mBase, uint32(v25)+356)) = int64(0)
				}
				v58 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v11))), uint32(v58))
				*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[1])) = l0
				v69 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
				if v69 != 0 {
					F_s_lock(m, l0, int32(_a_F_ConditionVariablePrepareToSleep_0))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
						v78 = v75 + v9*int32(768)
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v79 == int32(-1) {
							*(*int64)(unsafe.Add(mBase, uint32(v78)+356)) = int64(-1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v78)+360)) = v79
							v87 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
							*(*int32)(unsafe.Add(mBase, uint32(v88+v79*int32(768))+356)) = v9
							*(*int32)(unsafe.Add(mBase, uint32(v78)+356)) = int32(-1)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
						v96 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v96))
						return
					}
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
					v78 = v75 + v9*int32(768)
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v79 == int32(-1) {
						*(*int64)(unsafe.Add(mBase, uint32(v78)+356)) = int64(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v78)+360)) = v79
						v87 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
						*(*int32)(unsafe.Add(mBase, uint32(v88+v79*int32(768))+356)) = v9
						*(*int32)(unsafe.Add(mBase, uint32(v78)+356)) = int32(-1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
					v96 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v96))
					return
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[0]))
			v25 = v20 + v22*int32(768)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+360))
			if v26 == int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+356))
				if v29 == int32(0) {
				} else {
					v37 = v29
					*(*int32)(unsafe.Add(mBase, uint32(v20+v26*int32(768))+356)) = v37
					v42 = v26
					v43 = v37
					if v43 == int32(-1) {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v42
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
						*(*int32)(unsafe.Add(mBase, uint32(v49+v43*int32(768))+360)) = v42
					}
					*(*int64)(unsafe.Add(mBase, uint32(v25)+356)) = int64(0)
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+356))
				if v26 != int32(-1) {
					v37 = v32
					*(*int32)(unsafe.Add(mBase, uint32(v20+v26*int32(768))+356)) = v37
					v42 = v26
					v43 = v37
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v32
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+360))
					v42 = v36
					v43 = v32
				}
				if v43 == int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v42
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
					*(*int32)(unsafe.Add(mBase, uint32(v49+v43*int32(768))+360)) = v42
				}
				*(*int64)(unsafe.Add(mBase, uint32(v25)+356)) = int64(0)
			}
			v58 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v11))), uint32(v58))
			*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[1])) = l0
			v69 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
			if v69 != 0 {
				F_s_lock(m, l0, int32(_a_F_ConditionVariablePrepareToSleep_0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
					v78 = v75 + v9*int32(768)
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					if v79 == int32(-1) {
						*(*int64)(unsafe.Add(mBase, uint32(v78)+356)) = int64(-1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v78)+360)) = v79
						v87 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
						*(*int32)(unsafe.Add(mBase, uint32(v88+v79*int32(768))+356)) = v9
						*(*int32)(unsafe.Add(mBase, uint32(v78)+356)) = int32(-1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
					v96 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v96))
					return
				}
			} else {
				v74 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
				v78 = v75 + v9*int32(768)
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v79 == int32(-1) {
					*(*int64)(unsafe.Add(mBase, uint32(v78)+356)) = int64(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v78)+360)) = v79
					v87 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
					*(*int32)(unsafe.Add(mBase, uint32(v88+v79*int32(768))+356)) = v9
					*(*int32)(unsafe.Add(mBase, uint32(v78)+356)) = int32(-1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
				v96 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v96))
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[1])) = l0
		v69 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
		if v69 != 0 {
			F_s_lock(m, l0, int32(_a_F_ConditionVariablePrepareToSleep_0))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				v74 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
				v78 = v75 + v9*int32(768)
				v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v79 == int32(-1) {
					*(*int64)(unsafe.Add(mBase, uint32(v78)+356)) = int64(-1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v78)+360)) = v79
					v87 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
					v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
					*(*int32)(unsafe.Add(mBase, uint32(v88+v79*int32(768))+356)) = v9
					*(*int32)(unsafe.Add(mBase, uint32(v78)+356)) = int32(-1)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
				v96 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v96))
				return
			}
		} else {
			v74 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
			v78 = v75 + v9*int32(768)
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v79 == int32(-1) {
				*(*int64)(unsafe.Add(mBase, uint32(v78)+356)) = int64(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v78)+360)) = v79
				v87 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionVariablePrepareToSleep[2]))
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
				*(*int32)(unsafe.Add(mBase, uint32(v88+v79*int32(768))+356)) = v9
				*(*int32)(unsafe.Add(mBase, uint32(v78)+356)) = int32(-1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
			v96 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v96))
			return
		}
	}
}
func F_ConditionalLockBufferForCleanup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
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
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v126 int64
	_ = v126
	var v135 int64
	_ = v135
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	var v154 int64
	_ = v154
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v188 int64
	_ = v188
	var v190 int64
	_ = v190
	var v200 int64
	_ = v200
	var v203 int64
	_ = v203
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l0 < v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return v220
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[0]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+(l0^int32(-1))<<(uint(int32(2))%32))))
	v220 = base.B2i32(v20 == int32(1))
	goto L1
L3:
	;
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[1]))
	if v24 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	if v42 != int32(1) {
		v220 = v2
		goto L1
	} else {
		goto L13
	}
L6:
	;
	v28 = v24 << (uint(int32(4)) % 32)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_ConditionalLockBufferForCleanup[2])))
	if v31 == l0 {
		v41 = v28 + int32(_a_F_ConditionalLockBufferForCleanup_0)
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v35 = F_GetPrivateRefCountEntrySlow(m, l0, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	return int32(0)
L11:
	;
	if v35 == int32(0) {
		v220 = v2
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v41 = v35
	goto L5
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[3]))
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[1]))
	if v48 != int32(-1) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v52 = v48 << (uint(int32(4)) % 32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+uint32(_c_F_ConditionalLockBufferForCleanup[2])))
	if v55 == l0 {
		v61 = v52 + int32(_a_F_ConditionalLockBufferForCleanup_0)
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v59 = F_GetPrivateRefCountEntrySlow(m, l0, int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v61 = v59
	goto L14
L20:
	;
	v220 = int32(0)
	goto L1
L21:
	;
	goto L22
L22:
	;
	v64 = int32(_a_F_ConditionalLockBufferForCleanup_1)
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[4])) = v66 + int32(1)
	v74 = v46 + l0*int32(56) - int32(32)
	v75 = int64(0)
	v78 = base.AtomicRmwCmpxchg64(m, v74, int32(0), v75, v75)
	v84 = v78
	goto L23
L23:
	;
	v86 = int64(0)
	v90 = base.B2i32(v84&int64(18014381329612800) == v86)
	if v84&int64(18014381329612800) == v86 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v90 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v91 = int64(9007199254740992)
	goto L27
L26:
	;
	v91 = v86
	goto L27
L27:
	;
	v94 = base.AtomicRmwCmpxchg64(m, v74, int32(0), v84, v91|v84)
	if v84 != v94 {
		v84 = v94
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	v99 = int32(_a_F_ConditionalLockBufferForCleanup_1)
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[4])) = v101 - int32(1)
	v220 = int32(0)
	goto L1
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+12)) = int32(3)
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[3]))
	v113 = v108 + l0*int32(56) - int32(32)
	v114 = int64(4194304)
	v116 = base.AtomicRmwOr64(m, v113, int32(0), v114)
	if v116&v114 != int64(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v126 = v116
	goto L35
L33:
	;
	v200 = v116
	goto L34
L34:
	;
	v203 = base.AtomicRmwSub64(m, v113, int32(0), int64(4194304))
	if v200&int64(262143) == int64(1) {
		goto L56
	} else {
		goto L57
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = int32(_a_F_ConditionalLockBufferForCleanup_2)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(_a_F_ConditionalLockBufferForCleanup_3)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_ConditionalLockBufferForCleanup_4)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(0)
	v135 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v135
	if v126&int64(4194304) != v135 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v200 = v190
	goto L34
L37:
	;
	goto L40
L38:
	;
	goto L39
L39:
	;
	v168 = int32(_a_F_ConditionalLockBufferForCleanup_5)
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[5]))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(8))+8))
	if v171 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	F_perform_spin_delay(m, v9+int32(8))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L10
	} else {
		goto L42
	}
L41:
	;
	goto L39
L42:
	;
	v151 = int64(0)
	v154 = base.AtomicRmwCmpxchg64(m, v113, int32(0), v151, v151)
	if v154&int64(4194304) != v151 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v188 = int64(4194304)
	v190 = base.AtomicRmwOr64(m, v113, int32(0), v188)
	if v190&v188 != int64(0) {
		v126 = v190
		goto L35
	} else {
		goto L55
	}
L45:
	;
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[5])) = v186
	goto L45
L47:
	;
	if int32(999) < v169 {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v169 < int32(11) {
		goto L45
	} else {
		goto L54
	}
L50:
	;
	v176 = int32(900)
	if v176 <= v169 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v179 = v176
	goto L53
L52:
	;
	v179 = v169
	goto L53
L53:
	;
	v186 = v179 + int32(100)
	goto L46
L54:
	;
	v186 = v169 - int32(1)
	goto L46
L55:
	;
	goto L36
L56:
	;
	v220 = int32(1)
	goto L1
L57:
	;
	goto L58
L58:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_ConditionalLockBufferForCleanup[3]))
	v212 = int32(56)
	F_BufferLockUnlock(m, l0, v211+l0*v212-v212)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	v220 = int32(0)
	goto L1
}
func F_CountChildren(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CountChildren[0]))
	if base.B2i32(v12 == v2)|base.B2i32(v12 == int32(_a_F_CountChildren_0)) == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = v12
	v33 = v2
	goto L4
L2:
	;
	v103 = v2
	goto L3
L3:
	;
	m.G0 = v9 + int32(16)
	return v103
L4:
	;
	if int32(base.Ui32(l0&int32(64))>>(uint(int32(6))%32))^int32(base.Ui32(l0&int32(2))>>(uint(int32(1))%32)) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v103 = v95
	goto L3
L6:
	;
	v59 = v30 - int32(12)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if int32(base.Ui32(l0)>>(uint(v60)%32))&int32(1) != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v38 = v30 - int32(12)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v39 != int32(1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30-int32(16))))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_CountChildren[1]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v44<<(uint(int32(2))%32))+48))
	goto L9
L9:
	;
	if base.B2i32(v50 == int32(3)) == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = int32(6)
	goto L6
L11:
	;
	v66 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v95 = v33
	goto L13
L13:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v96 != int32(_a_F_CountChildren_0) {
		v30 = v96
		v33 = v95
		goto L4
	} else {
		goto L25
	}
L14:
	;
	return int32(0)
L15:
	;
	if v66 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if base.Ui32(v70) <= base.Ui32(int32(17)) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	v95 = v33 + int32(1)
	goto L13
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v30-int32(20))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v77
	F_errmsg_internal(m, int32(_a_F_CountChildren_1), v9)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L14
	} else {
		goto L23
	}
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_c_F_CountChildren[2])))
	v77 = v75
	goto L22
L21:
	;
	v77 = int32(_a_F_CountChildren_2)
	goto L22
L22:
	;
	goto L19
L23:
	;
	F_errfinish(m, int32(_a_F_CountChildren_3), int32(4000), int32(_a_F_CountChildren_4))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	goto L5
}
func F_CountDBBackends(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CountDBBackends[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_CountDBBackends[1]))
	v18 = F_LWLockAcquire(m, v14+int32(512), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v22 <= int32(0) {
			v108 = v2
		} else {
			v26 = v12 + int32(36)
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_CountDBBackends[2]))
			v29 = int32(0)
			if v22 != int32(1) {
				v38 = v29
				v39 = int32(0)
				v40 = v2
				for {
					v49 = v26 + v38<<(uint(int32(2))%32)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
					v53 = v28 + v50*int32(768)
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
					if v54 == int32(0) {
						v61 = v40
					} else {
						if l0 != 0 {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
							if v57 != l0 {
								v61 = v40
							} else {
								v61 = v40 + int32(1)
							}
						} else {
							v61 = v40 + int32(1)
						}
					}
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
					v65 = v28 + v62*int32(768)
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
					if v66 == int32(0) {
						v73 = v61
					} else {
						if l0 != 0 {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
							if v69 != l0 {
								v73 = v61
							} else {
								v73 = v61 + int32(1)
							}
						} else {
							v73 = v61 + int32(1)
						}
					}
					v74 = int32(2)
					v75 = v38 + v74
					v77 = v39 + v74
					if v77 != v22&int32(2147483646) {
						v38 = v75
						v39 = v77
						v40 = v73
						continue
					} else {
						break
					}
					break
				}
				if v22&int32(1) == int32(0) {
					v108 = v73
				} else {
					v82 = v75
					v84 = v73
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v26+v82<<(uint(int32(2))%32))))
					v97 = v28 + v94*int32(768)
					v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
					if v98 == int32(0) {
						v108 = v84
					} else {
						if l0 != 0 {
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
							if v101 != l0 {
								v108 = v84
							} else {
								v108 = v84 + int32(1)
							}
						} else {
							v108 = v84 + int32(1)
						}
					}
				}
			} else {
				v82 = v29
				v84 = v2
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v26+v82<<(uint(int32(2))%32))))
				v97 = v28 + v94*int32(768)
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
				if v98 == int32(0) {
					v108 = v84
				} else {
					if l0 != 0 {
						v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
						if v101 != l0 {
							v108 = v84
						} else {
							v108 = v84 + int32(1)
						}
					} else {
						v108 = v84 + int32(1)
					}
				}
			}
		}
		v116 = *(*int32)(unsafe.Add(mBase, _c_F_CountDBBackends[1]))
		F_LWLockRelease(m, v116+int32(512))
		mBase = m.M
		v120 = m.ExcPending
		if v120 != 0 {
			return int32(0)
		} else {
			return v108
		}
	}
}
func F_codepoint_range_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v4) <= base.Ui32(v3) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v9 = base.B2i32(base.Ui32(v6) < base.Ui32(v3))
	} else {
		v9 = int32(-1)
	}
	return v9
}
func F_colname_is_unique(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
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
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v8 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v253
L2:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v190 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v11 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v177 = int32(0)
	v179 = F_hash_search(m, v8, l0, v177, v177)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L56
	} else {
		goto L57
	}
L6:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v19 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if int32(0) < v67 {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14+v19<<(uint(int32(2))%32))))
	if v26 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v58 = v19 + int32(1)
	if v58 != v11 {
		v19 = v58
		goto L9
	} else {
		goto L21
	}
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v31 == int32(0))|base.B2i32(v31 != v34) != 0 {
		v52 = v31
		v53 = v34
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v52-v53 != 0 {
		goto L11
	} else {
		goto L20
	}
L14:
	;
	goto L13
L15:
	;
	v37 = v26
	v38 = l0
	goto L16
L16:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v42 == int32(0) {
		v52 = v42
		v53 = v41
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v52 = v42
	v53 = v41
	goto L14
L18:
	;
	v45 = int32(1)
	if v42 == v41 {
		v37 = v37 + v45
		v38 = v38 + v45
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	return int32(0)
L21:
	;
	goto L10
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v75 = int32(0)
	goto L25
L23:
	;
	goto L24
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v123 == int32(0) {
		goto L2
	} else {
		goto L38
	}
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v70+v75<<(uint(int32(2))%32))))
	if v82 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	v114 = v75 + int32(1)
	if v114 != v67 {
		v75 = v114
		goto L25
	} else {
		goto L37
	}
L28:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v87 == int32(0))|base.B2i32(v87 != v90) != 0 {
		v108 = v87
		v109 = v90
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v108-v109 != 0 {
		goto L27
	} else {
		goto L36
	}
L30:
	;
	goto L29
L31:
	;
	v93 = v82
	v94 = l0
	goto L32
L32:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
	if v98 == int32(0) {
		v108 = v98
		v109 = v97
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v108 = v98
	v109 = v97
	goto L30
L34:
	;
	v101 = int32(1)
	if v98 == v97 {
		v93 = v93 + v101
		v94 = v94 + v101
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	return int32(0)
L37:
	;
	goto L26
L38:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v126 <= int32(0) {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v129 = int32(0)
	if v129 < v126 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v133 = v126
	goto L42
L41:
	;
	v133 = v129
	goto L42
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v138 = v129
	goto L43
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v134+v138<<(uint(int32(2))%32))))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v148 == int32(0))|base.B2i32(v148 != v151) != 0 {
		v169 = v148
		v170 = v151
		goto L46
	} else {
		goto L47
	}
L44:
	;
	return int32(0)
L45:
	;
	if v169-v170 != 0 {
		goto L52
	} else {
		goto L53
	}
L46:
	;
	goto L45
L47:
	;
	v154 = v145
	v155 = l0
	goto L48
L48:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	if v159 == int32(0) {
		v169 = v159
		v170 = v158
		goto L46
	} else {
		goto L50
	}
L49:
	;
	v169 = v159
	v170 = v158
	goto L46
L50:
	;
	v162 = int32(1)
	if v159 == v158 {
		v154 = v154 + v162
		v155 = v155 + v162
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v173 = v138 + int32(1)
	if v133 != v173 {
		v138 = v173
		goto L43
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L44
L55:
	;
	goto L2
L56:
	;
	return int32(0)
L57:
	;
	if v179 != 0 {
		v253 = v4
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L2
L59:
	;
	return int32(1)
L60:
	;
	goto L61
L61:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v196 <= int32(0) {
		v253 = int32(1)
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v199 = int32(0)
	if v199 < v196 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v202 = v196
	goto L65
L64:
	;
	v202 = v199
	goto L65
L65:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	v208 = int32(0)
	goto L66
L66:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v203+v208<<(uint(int32(2))%32))))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v218 == int32(0))|base.B2i32(v218 != v221) != 0 {
		v239 = v218
		v240 = v221
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v253 = v243
	goto L1
L68:
	;
	v242 = int32(0)
	v243 = base.B2i32(v241 != v242)
	if v241 == v242 {
		v253 = v243
		goto L1
	} else {
		goto L75
	}
L69:
	;
	v241 = v239 - v240
	goto L68
L70:
	;
	v224 = v215
	v225 = l0
	goto L71
L71:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+1)))
	if v229 == int32(0) {
		v239 = v229
		v240 = v228
		goto L69
	} else {
		goto L73
	}
L72:
	;
	v239 = v229
	v240 = v228
	goto L69
L73:
	;
	v232 = int32(1)
	if v229 == v228 {
		v224 = v224 + v232
		v225 = v225 + v232
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v247 = v208 + int32(1)
	if v247 != v202 {
		v208 = v247
		goto L66
	} else {
		goto L76
	}
L76:
	;
	goto L67
}
func F_compact_trigram(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	v9 = int32(255)
	switch l2 {
	case 0:
		v61 = l2
		*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v61)
		v69 = int32(base.Ui32(v61) >> (uint(int32(16)) % 32))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v69)
		return
	default:
		v20 = l1
		v21 = l2
		v23 = v9
		v24 = v9
		v25 = v9
		v26 = v9
		for {
			v27 = int32(8)
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			v35 = *(*int32)(unsafe.Add(mBase, uint32((v29^v23)<<(uint(int32(2))%32))+uint32(_c_F_compact_trigram[0])))
			v38 = int32(16)
			v43 = int32(24)
			v46 = v35 ^ (v24<<(uint(v27)%32)&int32(_a_F_compact_trigram_0) | v25<<(uint(v38)%32)&int32(16711680) | v26<<(uint(v43)%32))
			v53 = int32(1)
			v56 = v21 - v53
			if v56 != 0 {
				v20 = v20 + v53
				v21 = v56
				v23 = int32(base.Ui32(v46) >> (uint(v43) % 32))
				v24 = v35
				v25 = int32(base.Ui32(v46) >> (uint(v27) % 32))
				v26 = int32(base.Ui32(v46) >> (uint(v38) % 32))
				continue
			} else {
				break
			}
			break
		}
		v61 = v46 ^ int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v61)
		v69 = int32(base.Ui32(v61) >> (uint(int32(16)) % 32))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v69)
		return
	case 3:
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v13)
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v15)
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v17)
		return
	}
}
func F_compare_lexeme_textfreq(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v8 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v37 < v6 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v14 == int32(18) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v25 = int32(1)
	if v8&v25 != 0 {
		v37 = int32(base.Ui32(v8)>>(uint(v25)%32)) - v25
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v17 = int32(16)
	goto L7
L6:
	;
	v17 = int32(0)
	goto L7
L7:
	;
	if base.Ui32((v14-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v24 = int32(4)
	goto L10
L9:
	;
	v24 = v17
	goto L10
L10:
	;
	v37 = v24
	goto L1
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v37 = int32(base.Ui32(v31)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	return int32(1)
L13:
	;
	goto L14
L14:
	;
	if v6 < v37 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(-1)
L16:
	;
	goto L17
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v45 = int32(1)
	if v8&v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v49 = v45
	goto L20
L19:
	;
	v49 = int32(4)
	goto L20
L20:
	;
	v50 = v7 + v49
	if v6 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	return v95
L22:
	;
	v95 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v56 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v57 = v44
	v58 = v50
	v59 = v6
	v60 = v56
	goto L29
L26:
	;
	v83 = v50
	v87 = int32(0)
	goto L27
L27:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v95 = v87 - v88
	goto L21
L28:
	;
	v83 = v78
	v87 = v80
	goto L27
L29:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if base.B2i32(v60 != v62)|base.B2i32(v62 == int32(0)) != 0 {
		v78 = v58
		v80 = v60
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v78 = v72
	v80 = int32(0)
	goto L28
L31:
	;
	v68 = v59 - int32(1)
	if v68 == int32(0) {
		v78 = v58
		v80 = v60
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v71 = int32(1)
	v72 = v58 + v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	if v73 != 0 {
		v57 = v57 + v71
		v58 = v72
		v59 = v68
		v60 = v73
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
}
func F_compare_scalars_simple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v9 = m.T0[v8].(func(*base.Module, int64, int64, int32) int32)(m, v6, v7, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 < int32(0) {
			v16 = int32(1)
		} else {
			v16 = int32(0) - v9
		}
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
		if v17 != 0 {
			v18 = v16
		} else {
			v18 = v9
		}
		return v18
	}
}
func F_compute_bucket(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v20 = base.I32_extend16_s(v19)
	v22 = base.B2i32(int32(0) <= v20)
	if int32(0) <= v20 {
		v23 = int32(-8)
	} else {
		v23 = int32(-6)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(base.Ui32(int32(base.Ui32(v14)>>(uint(int32(2))%32))+v23) >> (uint(int32(1)) % 32))
	if int32(0) <= v20 {
		v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
		v38 = v28
	} else {
		v38 = v19<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v19&int32(63)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v38
	v40 = int32(_a_F_compute_bucket_0)
	v41 = v19 & v40
	if v41 != v40 {
		if v41 != int32(_a_F_compute_bucket_1) {
			v52 = v41
		} else {
			v52 = v19 << (uint(int32(1)) % 32) & int32(_a_F_compute_bucket_2)
		}
	} else {
		v52 = v19 & int32(_a_F_compute_bucket_3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v52
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v54
	v63 = base.B2i32(v20 < v54)
	if v20 < v54 {
		v64 = int32(base.Ui32(v19)>>(uint(int32(7))%32)) & int32(63)
	} else {
		v64 = v19 & int32(_a_F_compute_bucket_4)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v64
	if v20 < v54 {
		v68 = int32(6)
	} else {
		v68 = int32(8)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = l1 + v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v77 = base.I32_extend16_s(v76)
	v79 = base.B2i32(int32(0) <= v77)
	if int32(0) <= v77 {
		v80 = int32(-8)
	} else {
		v80 = int32(-6)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(base.Ui32(int32(base.Ui32(v71)>>(uint(int32(2))%32))+v80) >> (uint(int32(1)) % 32))
	if int32(0) <= v77 {
		v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
		v95 = v85
	} else {
		v95 = v76<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v76&int32(63)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v95
	v97 = int32(_a_F_compute_bucket_0)
	v98 = v76 & v97
	if v98 != v97 {
		if v98 != int32(_a_F_compute_bucket_1) {
			v109 = v98
		} else {
			v109 = v76 << (uint(int32(1)) % 32) & int32(_a_F_compute_bucket_2)
		}
	} else {
		v109 = v76 & int32(_a_F_compute_bucket_3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v109
	v111 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v111
	v120 = base.B2i32(v77 < v111)
	if v77 < v111 {
		v121 = int32(base.Ui32(v76)>>(uint(int32(7))%32)) & int32(63)
	} else {
		v121 = v76 & int32(_a_F_compute_bucket_4)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v121
	if v77 < v111 {
		v125 = int32(6)
	} else {
		v125 = int32(8)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = l2 + v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v134 = base.I32_extend16_s(v133)
	v136 = base.B2i32(int32(0) <= v134)
	if int32(0) <= v134 {
		v137 = int32(-8)
	} else {
		v137 = int32(-6)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(base.Ui32(int32(base.Ui32(v128)>>(uint(int32(2))%32))+v137) >> (uint(int32(1)) % 32))
	if int32(0) <= v134 {
		v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
		v152 = v142
	} else {
		v152 = v133<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v133&int32(63)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v152
	v154 = int32(_a_F_compute_bucket_0)
	v155 = v133 & v154
	if v155 != v154 {
		if v155 != int32(_a_F_compute_bucket_1) {
			v166 = v155
		} else {
			v166 = v133 << (uint(int32(1)) % 32) & int32(_a_F_compute_bucket_2)
		}
	} else {
		v166 = v133 & int32(_a_F_compute_bucket_3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v166
	v168 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v168
	v173 = base.B2i32(v134 < v168)
	if v134 < v168 {
		v174 = int32(6)
	} else {
		v174 = int32(8)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l0 + v174
	if v134 < v168 {
		v183 = int32(base.Ui32(v133)>>(uint(int32(7))%32)) & int32(63)
	} else {
		v183 = v133 & int32(_a_F_compute_bucket_4)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v183
	v186 = v12 + int32(8)
	v188 = v12 + int32(56)
	F_sub_var(m, v186, v188, v186)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		return
	} else {
		v192 = v12 + int32(32)
		F_sub_var(m, v192, v188, v192)
		mBase = m.M
		v194 = m.ExcPending
		if v194 != 0 {
			return
		} else {
			v195 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			v196 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			F_mul_var(m, v186, l3, v186, v195+v196)
			mBase = m.M
			v199 = m.ExcPending
			if v199 != 0 {
				return
			} else {
				v200 = int32(0)
				F_div_var(m, v186, v192, l4, v200, v200, int32(1))
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return
				} else {
					F_add_var(m, l4, int32(_a_F_compute_bucket_5), l4)
					mBase = m.M
					v207 = m.ExcPending
					if v207 != 0 {
						return
					} else {
						v208 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
						if v208 != 0 {
							F_pfree(m, v208)
							mBase = m.M
							v210 = m.ExcPending
							if v210 != 0 {
								return
							} else {
								v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
								if v211 != 0 {
									F_pfree(m, v211)
									mBase = m.M
									v213 = m.ExcPending
									if v213 != 0 {
										return
									} else {
										m.G0 = v12 + int32(80)
										return
									}
								} else {
									m.G0 = v12 + int32(80)
									return
								}
							}
						} else {
							v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
							if v211 != 0 {
								F_pfree(m, v211)
								mBase = m.M
								v213 = m.ExcPending
								if v213 != 0 {
									return
								} else {
									m.G0 = v12 + int32(80)
									return
								}
							} else {
								m.G0 = v12 + int32(80)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_compute_distinct_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 float64
	_ = v87
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 float64
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v156 float64
	_ = v156
	var v157 int64
	_ = v157
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int64
	_ = v294
	var v296 int32
	_ = v296
	var v297 int64
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int64
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int64
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int64
	_ = v374
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v414 int32
	_ = v414
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v440 float64
	_ = v440
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 float64
	_ = v455
	var v457 float32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 float64
	_ = v565
	var v569 float64
	_ = v569
	var v576 float64
	_ = v576
	var v577 float64
	_ = v577
	var v579 float64
	_ = v579
	var v585 float64
	_ = v585
	var v586 float64
	_ = v586
	var v589 float64
	_ = v589
	var v591 float64
	_ = v591
	var v624 float32
	_ = v624
	var v626 float64
	_ = v626
	var v632 float32
	_ = v632
	var v634 float32
	_ = v634
	var v637 int32
	_ = v637
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v742 int32
	_ = v742
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v835 float32
	_ = v835
	var v836 float64
	_ = v836
	var v837 float32
	_ = v837
	var v839 float64
	_ = v839
	var v845 int32
	_ = v845
	var v851 float64
	_ = v851
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v877 int32
	_ = v877
	var v880 float64
	_ = v880
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v905 float64
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v917 int32
	_ = v917
	var v920 float64
	_ = v920
	var v935 int32
	_ = v935
	var v938 float64
	_ = v938
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v954 float64
	_ = v954
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v967 float64
	_ = v967
	var v978 float64
	_ = v978
	var v985 int32
	_ = v985
	var v991 float64
	_ = v991
	var v1002 float64
	_ = v1002
	var v1003 float64
	_ = v1003
	var v1007 float64
	_ = v1007
	var v1010 float64
	_ = v1010
	var v1013 float64
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1017 float64
	_ = v1017
	var v1021 float64
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1027 float64
	_ = v1027
	var v1029 float64
	_ = v1029
	var v1035 float64
	_ = v1035
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1073 int32
	_ = v1073
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1145 int32
	_ = v1145
	var v1146 int64
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int64
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	v5 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(32)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+78)))
	if v34 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+76)))
	v40 = int32(_a_F_compute_distinct_stats_0)
	v45 = base.B2i32(v37 < int32(0))
	v46 = base.B2i32(v37&v40 == v40)
	goto L3
L2:
	;
	v45 = v5
	v46 = v5
	goto L3
L3:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v48 = int32(10)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = v49 << (uint(int32(1)) % 32)
	if v51 <= v48 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v54 = v48
	goto L6
L5:
	;
	v54 = v51
	goto L6
L6:
	;
	v57 = F_palloc(m, v54<<(uint(int32(4))%32))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	F_fmgr_info(m, v59, v31+int32(4))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if l2 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	m.G0 = v31 + int32(32)
	return
L11:
	;
	v71 = v5
	v76 = v5
	v77 = v5
	v83 = v5
	v86 = v5
	v87 = float64(0)
	goto L12
L12:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L7
	} else {
		goto L14
	}
L13:
	;
	if int32(0) < v429 {
		goto L76
	} else {
		goto L77
	}
L14:
	;
	v99 = m.T0[l1].(func(*base.Module, int32, int32, int32) int64)(m, l0, v86, v31+int32(3))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+3)))
	if v101 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v448 = v86 + int32(1)
	if v448 != l2 {
		v71 = v424
		v76 = v429
		v77 = v430
		v83 = v436
		v86 = v448
		v87 = v440
		goto L12
	} else {
		goto L73
	}
L17:
	;
	v424 = v71
	v429 = v76
	v430 = v77 + int32(1)
	v436 = v83
	v440 = v87
	goto L16
L18:
	;
	goto L19
L19:
	;
	v107 = v76 + int32(1)
	if v46 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v158 = int32(0)
	if v158 < v71 {
		goto L45
	} else {
		goto L46
	}
L21:
	;
	v108 = base.I32_wrap_i64(v99)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v109 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	if v45 == int32(0) {
		v156 = v87
		v157 = v99
		goto L20
	} else {
		goto L40
	}
L24:
	;
	v136 = base.F64_add(v87, base.F64_convert_i32_u(v134))
	v137 = F_toast_raw_datum_size(m, v99)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L7
	} else {
		goto L35
	}
L25:
	;
	v113 = int32(18)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	if v115 == v113 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v126 = int32(1)
	if v109&v126 != 0 {
		v134 = int32(base.Ui32(v109) >> (uint(v126) % 32))
		goto L24
	} else {
		goto L34
	}
L28:
	;
	v118 = v113
	goto L30
L29:
	;
	v118 = int32(2)
	goto L30
L30:
	;
	if base.Ui32((v115-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v125 = int32(6)
	goto L33
L32:
	;
	v125 = v118
	goto L33
L33:
	;
	v134 = v125
	goto L24
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v134 = int32(base.Ui32(v130) >> (uint(int32(2)) % 32))
	goto L24
L35:
	;
	if base.Ui32(int32(1025)) <= base.Ui32(v137) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v424 = v71
	v429 = v107
	v430 = v77
	v436 = v83 + int32(1)
	v440 = v136
	goto L16
L37:
	;
	goto L38
L38:
	;
	v143 = F_pg_detoast_datum(m, v108)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v156 = v136
	v157 = base.I64_extend_i32_u(v143)
	goto L20
L40:
	;
	v149 = F_strlen(m, base.I32_wrap_i64(v99))
	mBase = m.M
	v156 = base.F64_add(v87, base.F64_convert_i32_u(v149+int32(1)))
	v157 = v99
	goto L20
L41:
	;
	if v216 < v238 {
		goto L70
	} else {
		goto L71
	}
L42:
	;
	if v216 == v237+v71-int32(2) {
		goto L41
	} else {
		goto L66
	}
L43:
	;
	v305 = int32(4)
	v307 = v57 + v240<<(uint(v305)%32)
	v310 = v57 + v238<<(uint(v305)%32)
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v310-int32(32))))
	*(*int64)(unsafe.Add(mBase, uint32(v307))) = v313
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v310-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v307)+8)) = v317
	v321 = v238 - int32(2)
	v323 = v240
	goto L42
L44:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+8)) = v251 + int32(1)
	if v165 == int32(0) {
		v424 = v71
		v429 = v107
		v430 = v77
		v436 = v83
		v440 = v156
		goto L16
	} else {
		goto L61
	}
L45:
	;
	v165 = v158
	v168 = v71
	goto L48
L46:
	;
	v216 = v71
	goto L47
L47:
	;
	v237 = base.B2i32(v71 < v54)
	v238 = v71 + v237
	v240 = v238 - int32(1)
	if v240 <= v216 {
		goto L41
	} else {
		goto L59
	}
L48:
	;
	v189 = int32(4)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v194 = v57 + v165<<(uint(v189)%32)
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v194)))
	v196 = F_FunctionCall2Coll(m, v31+v189, v191, v157, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L7
	} else {
		goto L50
	}
L49:
	;
	v216 = v205
	goto L47
L50:
	;
	if v196 != int64(0) {
		goto L44
	} else {
		goto L51
	}
L51:
	;
	if v165 < v168 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	if v201 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v205 = v168
	goto L54
L54:
	;
	v207 = v165 + int32(1)
	if v207 != v71 {
		v165 = v207
		v168 = v205
		goto L48
	} else {
		goto L58
	}
L55:
	;
	v204 = v165
	goto L57
L56:
	;
	v204 = v168
	goto L57
L57:
	;
	v205 = v204
	goto L54
L58:
	;
	goto L49
L59:
	;
	if (v71-v237+v216)&int32(1) == int32(0) {
		goto L43
	} else {
		goto L60
	}
L60:
	;
	v321 = v240
	v323 = v238
	goto L42
L61:
	;
	v261 = v165
	goto L62
L62:
	;
	v287 = v57 + v261<<(uint(int32(4))%32)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	v290 = v287 - int32(8)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	if v288 <= v291 {
		v424 = v71
		v429 = v107
		v430 = v77
		v436 = v83
		v440 = v156
		goto L16
	} else {
		goto L64
	}
L63:
	;
	v424 = v71
	v429 = v107
	v430 = v77
	v436 = v83
	v440 = v156
	goto L16
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+8)) = v291
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v287)))
	v296 = v287 - int32(16)
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v296)))
	*(*int64)(unsafe.Add(mBase, uint32(v287))) = v297
	*(*int64)(unsafe.Add(mBase, uint32(v296))) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = v288
	v301 = int32(1)
	if v301 < v261 {
		v261 = v261 - v301
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v329 = v321
	v331 = v323
	goto L67
L67:
	;
	v353 = int32(4)
	v355 = v57 + v329<<(uint(v353)%32)
	v358 = v57 + v331<<(uint(v353)%32)
	v359 = int32(32)
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v358-v359)))
	*(*int64)(unsafe.Add(mBase, uint32(v355))) = v361
	v363 = int32(24)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v358-v363)))
	*(*int32)(unsafe.Add(mBase, uint32(v355)+8)) = v365
	v368 = v329 - int32(1)
	v371 = v57 + v368<<(uint(v353)%32)
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v355-v359)))
	*(*int64)(unsafe.Add(mBase, uint32(v371))) = v374
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v355-v363)))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+8)) = v378
	v381 = v329 - int32(2)
	if v216 < v381 {
		v329 = v381
		v331 = v368
		goto L67
	} else {
		goto L69
	}
L68:
	;
	goto L41
L69:
	;
	goto L68
L70:
	;
	v414 = v57 + v216<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v414)+8)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v414))) = v157
	goto L72
L71:
	;
	goto L72
L72:
	;
	v424 = v238
	v429 = v107
	v430 = v77
	v436 = v83
	v440 = v156
	goto L16
L73:
	;
	goto L13
L74:
	;
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v624
	v626 = base.F64_promote_f32(v624)
	if base.F64_gt(v626, base.F64_mul(l3, float64(0.1))) != 0 {
		goto L106
	} else {
		goto L107
	}
L75:
	;
	if base.B2i32(v54 <= v424)|v436|base.B2i32(v554 != v424) == int32(0) {
		v624 = base.F32_convert_i32_s(v424)
		goto L74
	} else {
		goto L95
	}
L76:
	;
	v452 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v452)
	v455 = base.F64_convert_i32_s(l2)
	v457 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v430), v455))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v457
	if v45 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	if v430 <= int32(0) {
		goto L10
	} else {
		goto L91
	}
L79:
	;
	v464 = base.I32_trunc_sat_f64_s(base.F64_div(v440, base.F64_convert_i32_u(v429)))
	goto L81
L80:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v463 = int32(*(*int16)(unsafe.Add(mBase, uint32(v462)+76)))
	v464 = v463
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v464
	if int32(0) < v424 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v468 = int32(0)
	v474 = v468
	v477 = v468
	goto L86
L83:
	;
	goto L84
L84:
	;
	v624 = base.F32_neg(base.F32_sub(float32(1), v457))
	goto L74
L85:
	;
	if v474 != 0 {
		v554 = v474
		v555 = v477
		goto L75
	} else {
		goto L90
	}
L86:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v57+v474<<(uint(int32(4))%32))+8))
	if v501 == int32(1) {
		goto L85
	} else {
		goto L88
	}
L87:
	;
	v554 = v424
	v555 = v504
	goto L75
L88:
	;
	v504 = v501 + v477
	v506 = v474 + int32(1)
	if v506 != v424 {
		v474 = v506
		v477 = v504
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	goto L84
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v543 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v543)
	v545 = int32(0)
	if v45 == v545 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v549 = int32(*(*int16)(unsafe.Add(mBase, uint32(v548)+76)))
	v550 = v549
	goto L94
L93:
	;
	v550 = v545
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v550
	goto L10
L95:
	;
	v563 = v429 - v555
	v564 = v554 + v563
	v565 = float64(0)
	v569 = base.F64_mul(l3, base.F64_sub(float64(1), base.F64_promote_f32(v457)))
	if base.F64_gt(v569, v565) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if base.F64_lt(v585, v586) != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v585 = v565
	v586 = base.F64_convert_i32_s(v564)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v576 = base.F64_convert_i32_s(l2 - v430)
	v577 = base.F64_convert_i32_s(v564)
	v579 = base.F64_convert_i32_s(v563)
	v585 = base.F64_div(base.F64_mul(v576, v577), base.F64_add(base.F64_sub(v576, v579), base.F64_div(base.F64_mul(v576, v579), v569)))
	v586 = v577
	goto L96
L100:
	;
	v589 = v586
	goto L102
L101:
	;
	v589 = v585
	goto L102
L102:
	;
	if base.F64_gt(v589, v569) != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v591 = v569
	goto L105
L104:
	;
	v591 = v589
	goto L105
L105:
	;
	v624 = base.F32_demote_f64(base.F64_floor(base.F64_add(v591, float64(0.5))))
	goto L74
L106:
	;
	v632 = base.F32_demote_f64(base.F64_div(base.F64_neg(v626), l3))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v632
	v634 = v632
	goto L108
L107:
	;
	v634 = v624
	goto L108
L108:
	;
	v637 = int32(0)
	if base.B2i32(base.B2i32(base.F32_gt(v634, float32(0)) == v637)|(base.B2i32(v54 <= v424)|v436) == v637)&base.B2i32(v424 <= v49) == v637 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if v49 < v424 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v1073 = v424
	goto L111
L111:
	;
	if v1073 <= int32(0) {
		goto L10
	} else {
		goto L161
	}
L112:
	;
	v649 = v49
	goto L114
L113:
	;
	v649 = v424
	goto L114
L114:
	;
	if v649 <= int32(0) {
		goto L10
	} else {
		goto L115
	}
L115:
	;
	v653 = v649 & int32(3)
	v657 = F_palloc(m, v649<<(uint(int32(2))%32))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	v659 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v649) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v835 = *(*float32)(unsafe.Add(mBase, uint32(l0)+48))
	v836 = base.F64_promote_f32(v835)
	v837 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v839 = float64(0)
	v845 = int32(0)
	v851 = base.F64_convert_i32_s(l2)
	if base.F64_eq(l3, v851)|base.F64_le(l3, float64(1)) != 0 {
		v1051 = v649
		goto L129
	} else {
		goto L130
	}
L118:
	;
	v669 = v659
	v671 = int32(0)
	goto L121
L119:
	;
	v742 = v659
	goto L120
L120:
	;
	v770 = v742
	v771 = int32(0)
	goto L125
L121:
	;
	v693 = int32(2)
	v696 = int32(4)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v57+v669<<(uint(v696)%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v657+v669<<(uint(v693)%32)))) = v699
	v702 = v669 | int32(1)
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v57+v702<<(uint(v696)%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v657+v702<<(uint(v693)%32)))) = v709
	v712 = v669 | v693
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v57+v712<<(uint(v696)%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v657+v712<<(uint(v693)%32)))) = v719
	v722 = v669 | int32(3)
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v57+v722<<(uint(v696)%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v657+v722<<(uint(v693)%32)))) = v729
	v732 = v669 + v696
	v734 = v671 + v696
	if v734 != v649&int32(2147483644) {
		v669 = v732
		v671 = v734
		goto L121
	} else {
		goto L123
	}
L122:
	;
	if v653 == int32(0) {
		goto L117
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v742 = v732
	goto L120
L125:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v57+v770<<(uint(int32(4))%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v657+v770<<(uint(int32(2))%32)))) = v800
	v802 = int32(1)
	v805 = v771 + v802
	if v805 != v653 {
		v770 = v770 + v802
		v771 = v805
		goto L125
	} else {
		goto L127
	}
L126:
	;
	goto L117
L127:
	;
	goto L126
L128:
	;
	v1073 = v1051
	goto L111
L129:
	;
	goto L128
L130:
	;
	if base.Ui32(v649) < base.Ui32(int32(2)) {
		v967 = v839
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if base.F64_lt(v836, float64(0)) != 0 {
		goto L143
	} else {
		goto L144
	}
L132:
	;
	v863 = v649 - int32(1)
	v864 = int32(3)
	v865 = v863 & v864
	v866 = int32(0)
	if base.Ui32(v864) <= base.Ui32(v649-int32(2)) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v877 = v866
	v880 = v839
	v889 = v845
	goto L136
L134:
	;
	v917 = v866
	v920 = v839
	goto L135
L135:
	;
	v935 = v917
	v938 = v920
	v948 = v845
	goto L140
L136:
	;
	v893 = v657 + v877<<(uint(int32(2))%32)
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v893)))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v893)+4))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v893)+8))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v893)+12))
	v905 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v880, base.F64_convert_i32_s(v894)), base.F64_convert_i32_s(v897)), base.F64_convert_i32_s(v900)), base.F64_convert_i32_s(v903))
	v906 = int32(4)
	v907 = v877 + v906
	v909 = v889 + v906
	if v909 != v863&int32(-4) {
		v877 = v907
		v880 = v905
		v889 = v909
		goto L136
	} else {
		goto L138
	}
L137:
	;
	if v865 == int32(0) {
		v967 = v905
		goto L131
	} else {
		goto L139
	}
L138:
	;
	goto L137
L139:
	;
	v917 = v907
	v920 = v905
	goto L135
L140:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v657+v935<<(uint(int32(2))%32))))
	v954 = base.F64_add(v938, base.F64_convert_i32_s(v952))
	v955 = int32(1)
	v958 = v948 + v955
	if v958 != v865 {
		v935 = v935 + v955
		v938 = v954
		v948 = v958
		goto L140
	} else {
		goto L142
	}
L141:
	;
	v967 = v954
	goto L131
L142:
	;
	goto L141
L143:
	;
	v978 = base.F64_mul(l3, base.F64_neg(v836))
	goto L145
L144:
	;
	v978 = v836
	goto L145
L145:
	;
	v985 = v649
	v991 = v967
	goto L146
L146:
	;
	v1002 = float64(1)
	v1003 = float64(0)
	v1007 = base.F64_sub(base.F64_sub(v1002, base.F64_div(v991, v851)), base.F64_promote_f32(v837))
	if base.F64_lt(v1007, v1003) != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v1051 = int32(0)
	goto L129
L148:
	;
	v1010 = v1003
	goto L150
L149:
	;
	v1010 = v1007
	goto L150
L150:
	;
	if base.F64_gt(v1010, float64(1)) != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v1013 = v1002
	goto L153
L152:
	;
	v1013 = v1010
	goto L153
L153:
	;
	v1015 = v985 - int32(1)
	v1017 = base.F64_sub(v978, base.F64_convert_i32_u(v1015))
	if base.F64_gt(v1017, float64(1)) != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v1021 = base.F64_div(v1013, v1017)
	goto L156
L155:
	;
	v1021 = v1013
	goto L156
L156:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v657+v1015<<(uint(int32(2))%32))))
	v1027 = base.F64_convert_i32_s(v1026)
	v1029 = base.F64_div(base.F64_mul(l3, v1027), v851)
	v1035 = base.F64_sqrt(base.F64_div(base.F64_mul(base.F64_sub(l3, v851), base.F64_mul(base.F64_mul(v1029, v851), base.F64_sub(l3, v1029))), base.F64_mul(base.F64_mul(l3, l3), base.F64_add(l3, float64(-1)))))
	if base.F64_lt(base.F64_add(base.F64_add(base.F64_mul(v1021, v851), base.F64_add(v1035, v1035)), float64(0.5)), v1027) != 0 {
		v1051 = v985
		goto L129
	} else {
		goto L157
	}
L157:
	;
	if v1015 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v657+v985<<(uint(int32(2))%32)-int32(8))))
	v985 = v1015
	v991 = base.F64_sub(v991, base.F64_convert_i32_s(v1046))
	goto L146
L159:
	;
	goto L160
L160:
	;
	goto L147
L161:
	;
	v1099 = int32(_a_F_compute_distinct_stats_1)
	v1100 = *(*int32)(unsafe.Add(mBase, _c_F_compute_distinct_stats[0]))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_distinct_stats[0])) = v1102
	v1106 = F_palloc(m, v1073<<(uint(int32(3))%32))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L7
	} else {
		goto L162
	}
L162:
	;
	v1110 = F_palloc(m, v1073<<(uint(int32(2))%32))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L7
	} else {
		goto L163
	}
L163:
	;
	v1116 = int32(0)
	goto L164
L164:
	;
	v1145 = v57 + v1116<<(uint(int32(4))%32)
	v1146 = *(*int64)(unsafe.Add(mBase, uint32(v1145)))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1147)+78)))
	v1149 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1147)+76)))
	v1150 = F_datumCopy(m, v1146, v1148, v1149)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L7
	} else {
		goto L166
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_distinct_stats[0])) = v1100
	v1166 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v1166)
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1168
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v1110
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1171
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v1106
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v1073
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v1073
	goto L10
L166:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1106+v1116<<(uint(int32(3))%32)))) = v1150
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+8))
	*(*float32)(unsafe.Add(mBase, uint32(v1110+v1116<<(uint(int32(2))%32)))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v1156), v455))
	v1162 = v1116 + int32(1)
	if v1162 != v1073 {
		v1116 = v1162
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
}
func F_compute_scalar_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v25 float64
	_ = v25
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 float64
	_ = v110
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 float64
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v180 float64
	_ = v180
	var v181 int64
	_ = v181
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 float64
	_ = v199
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v244 float64
	_ = v244
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var __phi285 int32
	_ = __phi285
	var v291 int32
	_ = v291
	var __phi291 int32
	_ = __phi291
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v370 int32
	_ = v370
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v439 int32
	_ = v439
	var v440 float64
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 float64
	_ = v447
	var v449 float32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v468 int32
	_ = v468
	var v469 float64
	_ = v469
	var v473 float64
	_ = v473
	var v480 float64
	_ = v480
	var v481 float64
	_ = v481
	var v485 float64
	_ = v485
	var v491 float64
	_ = v491
	var v492 float64
	_ = v492
	var v495 float64
	_ = v495
	var v497 float64
	_ = v497
	var v507 float32
	_ = v507
	var v509 float64
	_ = v509
	var v515 float32
	_ = v515
	var v517 float32
	_ = v517
	var v520 int32
	_ = v520
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v628 int32
	_ = v628
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v733 float32
	_ = v733
	var v734 float64
	_ = v734
	var v735 float32
	_ = v735
	var v737 float64
	_ = v737
	var v743 int32
	_ = v743
	var v749 float64
	_ = v749
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v775 int32
	_ = v775
	var v778 float64
	_ = v778
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v803 float64
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v815 int32
	_ = v815
	var v818 float64
	_ = v818
	var v833 int32
	_ = v833
	var v836 float64
	_ = v836
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v852 float64
	_ = v852
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v865 float64
	_ = v865
	var v876 float64
	_ = v876
	var v883 int32
	_ = v883
	var v889 float64
	_ = v889
	var v900 float64
	_ = v900
	var v901 float64
	_ = v901
	var v905 float64
	_ = v905
	var v908 float64
	_ = v908
	var v911 float64
	_ = v911
	var v913 int32
	_ = v913
	var v915 float64
	_ = v915
	var v919 float64
	_ = v919
	var v924 int32
	_ = v924
	var v925 float64
	_ = v925
	var v927 float64
	_ = v927
	var v933 float64
	_ = v933
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v971 int32
	_ = v971
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1019 int32
	_ = v1019
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1055 int64
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int64
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1091 int32
	_ = v1091
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1202 int32
	_ = v1202
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1289 int64
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int64
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1326 int32
	_ = v1326
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1368 float64
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1373 float64
	_ = v1373
	var v1375 float64
	_ = v1375
	var v1377 float64
	_ = v1377
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1409 int32
	_ = v1409
	var v1414 float32
	_ = v1414
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1431 int32
	_ = v1431
	var v1433 int32
	_ = v1433
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	v5 = int32(0)
	v25 = float64(0)
	v33 = m.G0
	v35 = v33 - int32(48)
	m.G0 = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+78)))
	if v38 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+76)))
	v44 = int32(_a_F_compute_scalar_stats_0)
	v48 = base.B2i32(v41&v44 == v44)
	v49 = base.B2i32(v41 < int32(0))
	goto L3
L2:
	;
	v48 = v5
	v49 = v5
	goto L3
L3:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v54 = F_palloc(m, l2<<(uint(int32(4))%32))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v58 = F_palloc(m, l2<<(uint(int32(2))%32))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v62 = F_palloc(m, v50<<(uint(int32(3))%32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v64 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+44)) = v64
	v66 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+36)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v35)+28)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v35)+20)) = v66
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v75
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)) = uint8(v64)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	F_PrepareSortSupportFromOrderingOp(m, v79, v35+int32(12))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if l2 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	m.G0 = v35 + int32(48)
	return
L10:
	;
	v90 = v5
	v95 = v5
	v97 = v5
	v103 = v5
	v107 = v5
	v110 = v25
	goto L11
L11:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	if int32(0) < v196 {
		goto L41
	} else {
		goto L42
	}
L13:
	;
	v123 = m.T0[l1].(func(*base.Module, int32, int32, int32) int64)(m, l0, v90, v35+int32(4))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+4)))
	if v125 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v202 = v90 + int32(1)
	if v202 != l2 {
		v90 = v202
		v95 = v195
		v97 = v196
		v103 = v197
		v107 = v198
		v110 = v199
		goto L11
	} else {
		goto L40
	}
L16:
	;
	v195 = v95 + int32(1)
	v196 = v97
	v197 = v103
	v198 = v107
	v199 = v110
	goto L15
L17:
	;
	goto L18
L18:
	;
	v131 = v107 + int32(1)
	if v48 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v184 = v54 + v97<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v184)+8)) = v97
	*(*int64)(unsafe.Add(mBase, uint32(v184))) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v58+v97<<(uint(int32(2))%32)))) = v97
	v195 = v95
	v196 = v97 + int32(1)
	v197 = v103
	v198 = v131
	v199 = v180
	goto L15
L20:
	;
	v132 = base.I32_wrap_i64(v123)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v133 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L22
L22:
	;
	if v49 == int32(0) {
		v180 = v110
		v181 = v123
		goto L19
	} else {
		goto L39
	}
L23:
	;
	v160 = base.F64_add(v110, base.F64_convert_i32_u(v158))
	v161 = F_toast_raw_datum_size(m, v123)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L34
	}
L24:
	;
	v137 = int32(18)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v139 == v137 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v150 = int32(1)
	if v133&v150 != 0 {
		v158 = int32(base.Ui32(v133) >> (uint(v150) % 32))
		goto L23
	} else {
		goto L33
	}
L27:
	;
	v142 = v137
	goto L29
L28:
	;
	v142 = int32(2)
	goto L29
L29:
	;
	if base.Ui32((v139-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v149 = int32(6)
	goto L32
L31:
	;
	v149 = v142
	goto L32
L32:
	;
	v158 = v149
	goto L23
L33:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v158 = int32(base.Ui32(v154) >> (uint(int32(2)) % 32))
	goto L23
L34:
	;
	if base.Ui32(int32(1025)) <= base.Ui32(v161) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v195 = v95
	v196 = v97
	v197 = v103 + int32(1)
	v198 = v131
	v199 = v160
	goto L15
L36:
	;
	goto L37
L37:
	;
	v167 = F_pg_detoast_datum(m, v132)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v180 = v160
	v181 = base.I64_extend_i32_u(v167)
	goto L19
L39:
	;
	v173 = F_strlen(m, base.I32_wrap_i64(v123))
	mBase = m.M
	v180 = base.F64_add(v110, base.F64_convert_i32_u(v173+int32(1)))
	v181 = v123
	goto L19
L40:
	;
	goto L12
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v35 + int32(12)
	F_qsort_interruptible(m, v54, v196, int32(16), int32(543), v35+int32(4))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if int32(0) < v198 {
		goto L177
	} else {
		goto L178
	}
L44:
	;
	v222 = int32(0)
	v229 = v5
	v230 = v5
	v232 = v5
	v235 = v5
	v244 = v25
	goto L45
L45:
	;
	v250 = v229 + int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v54+v222<<(uint(int32(4))%32))+8))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v58+v255<<(uint(int32(2))%32))))
	if v255 != v261 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v444 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v444)
	v447 = base.F64_convert_i32_s(l2)
	v449 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v195), v447))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v449
	if v49 != 0 {
		goto L63
	} else {
		goto L64
	}
L47:
	;
	v420 = v230
	v422 = v232
	v425 = v235
	v439 = v250
	goto L49
L48:
	;
	if v250 < int32(2) {
		v389 = v232
		v392 = v235
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v440 = base.F64_add(base.F64_mul(base.F64_convert_i32_u(v222), base.F64_convert_i32_s(v255)), v244)
	v442 = v222 + int32(1)
	if v442 != v196 {
		v222 = v442
		v229 = v439
		v230 = v420
		v232 = v422
		v235 = v425
		v244 = v440
		goto L45
	} else {
		goto L62
	}
L50:
	;
	v420 = v230 + int32(1)
	v422 = v389
	v425 = v392
	v439 = int32(0)
	goto L49
L51:
	;
	v268 = v235 + int32(1)
	v269 = base.B2i32(v232 < v50)
	if v269 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v62+v232<<(uint(int32(3))%32)-int32(8))))
	if v250 <= v277 {
		v389 = v232
		v392 = v268
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v279 = v269 + v232
	v281 = v279 - int32(1)
	if v281 <= int32(0) {
		v337 = v281
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v370 = v62 + v337<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v370)+4)) = v222 - v229
	*(*int32)(unsafe.Add(mBase, uint32(v370))) = v250
	v389 = v279
	v392 = v268
	goto L50
L57:
	;
	__phi285 = v281
	__phi291 = v279
	v285 = __phi285
	v291 = __phi291
	goto L58
L58:
	;
	v318 = v62 + v291<<(uint(int32(3))%32)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v318-int32(16))))
	if v250 <= v321 {
		v337 = v285
		goto L56
	} else {
		goto L60
	}
L59:
	;
	v337 = int32(0)
	goto L56
L60:
	;
	v325 = v62 + v285<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v321
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v318-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v325)+4)) = v329
	v331 = int32(1)
	if v331 < v285 {
		__phi285 = v285 - v331
		__phi291 = v285
		v285 = __phi285
		v291 = __phi291
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	goto L46
L63:
	;
	v456 = base.I32_trunc_sat_f64_s(base.F64_div(v199, base.F64_convert_i32_s(v198)))
	goto L65
L64:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v455 = int32(*(*int16)(unsafe.Add(mBase, uint32(v454)+76)))
	v456 = v455
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v456
	if v425 == int32(0) {
		v507 = base.F32_neg(base.F32_sub(float32(1), v449))
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v507
	v509 = base.F64_promote_f32(v507)
	if base.F64_gt(v509, base.F64_mul(l3, float64(0.1))) != 0 {
		goto L79
	} else {
		goto L80
	}
L67:
	;
	if v197|base.B2i32(v420 != v425) == int32(0) {
		v507 = base.F32_convert_i32_s(v425)
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v468 = v420 + v197
	v469 = float64(0)
	v473 = base.F64_mul(l3, base.F64_sub(float64(1), base.F64_promote_f32(v449)))
	if base.F64_gt(v473, v469) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if base.F64_lt(v491, v492) != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v491 = v469
	v492 = base.F64_convert_i32_s(v468)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v480 = base.F64_convert_i32_s(l2 - v195)
	v481 = base.F64_convert_i32_s(v468)
	v485 = base.F64_convert_i32_s(v197 - v425 + v420)
	v491 = base.F64_div(base.F64_mul(v480, v481), base.F64_add(base.F64_sub(v480, v485), base.F64_div(base.F64_mul(v480, v485), v473)))
	v492 = v481
	goto L69
L73:
	;
	v495 = v492
	goto L75
L74:
	;
	v495 = v491
	goto L75
L75:
	;
	if base.F64_gt(v495, v473) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v497 = v473
	goto L78
L77:
	;
	v497 = v495
	goto L78
L78:
	;
	v507 = base.F32_demote_f64(base.F64_floor(base.F64_add(v497, float64(0.5))))
	goto L66
L79:
	;
	v515 = base.F32_demote_f64(base.F64_div(base.F64_neg(v509), l3))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = v515
	v517 = v515
	goto L81
L80:
	;
	v517 = v507
	goto L81
L81:
	;
	v520 = int32(0)
	if base.B2i32(base.F32_gt(v517, float32(0)) == v520)|(base.B2i32(v420 != v422)|v197) == v520 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v1121 = v420 - v1091
	if v50 < v1121 {
		goto L144
	} else {
		goto L145
	}
L83:
	;
	v998 = int32(0)
	if v971 <= v998 {
		v1091 = v971
		v1118 = v998
		goto L82
	} else {
		goto L137
	}
L84:
	;
	if v420 <= v50 {
		v971 = v420
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if v50 < v422 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L86
L88:
	;
	v531 = v50
	goto L90
L89:
	;
	v531 = v422
	goto L90
L90:
	;
	if v531 <= int32(0) {
		v1091 = v531
		v1118 = int32(0)
		goto L82
	} else {
		goto L91
	}
L91:
	;
	v535 = v531 & int32(3)
	v539 = F_palloc(m, v531<<(uint(int32(2))%32))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v541 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v531) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v733 = *(*float32)(unsafe.Add(mBase, uint32(l0)+48))
	v734 = base.F64_promote_f32(v733)
	v735 = *(*float32)(unsafe.Add(mBase, uint32(l0)+40))
	v737 = float64(0)
	v743 = int32(0)
	v749 = base.F64_convert_i32_s(l2)
	if base.F64_eq(l3, v749)|base.F64_le(l3, float64(1)) != 0 {
		v949 = v531
		goto L105
	} else {
		goto L106
	}
L94:
	;
	v551 = v541
	v554 = int32(0)
	goto L97
L95:
	;
	v628 = v541
	goto L96
L96:
	;
	v660 = v628
	v662 = int32(0)
	goto L101
L97:
	;
	v579 = int32(2)
	v582 = int32(3)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v62+v551<<(uint(v582)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v539+v551<<(uint(v579)%32)))) = v585
	v588 = v551 | int32(1)
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v62+v588<<(uint(v582)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v539+v588<<(uint(v579)%32)))) = v595
	v598 = v551 | v579
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v62+v598<<(uint(v582)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v539+v598<<(uint(v579)%32)))) = v605
	v608 = v551 | v582
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v62+v608<<(uint(v582)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v539+v608<<(uint(v579)%32)))) = v615
	v617 = int32(4)
	v618 = v551 + v617
	v620 = v554 + v617
	if v620 != v531&int32(2147483644) {
		v551 = v618
		v554 = v620
		goto L97
	} else {
		goto L99
	}
L98:
	;
	if v535 == int32(0) {
		goto L93
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	v628 = v618
	goto L96
L101:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v62+v660<<(uint(int32(3))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v539+v660<<(uint(int32(2))%32)))) = v694
	v696 = int32(1)
	v699 = v662 + v696
	if v699 != v535 {
		v660 = v660 + v696
		v662 = v699
		goto L101
	} else {
		goto L103
	}
L102:
	;
	goto L93
L103:
	;
	goto L102
L104:
	;
	v971 = v949
	goto L83
L105:
	;
	goto L104
L106:
	;
	if base.Ui32(v531) < base.Ui32(int32(2)) {
		v865 = v737
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if base.F64_lt(v734, float64(0)) != 0 {
		goto L119
	} else {
		goto L120
	}
L108:
	;
	v761 = v531 - int32(1)
	v762 = int32(3)
	v763 = v761 & v762
	v764 = int32(0)
	if base.Ui32(v762) <= base.Ui32(v531-int32(2)) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v775 = v764
	v778 = v737
	v787 = v743
	goto L112
L110:
	;
	v815 = v764
	v818 = v737
	goto L111
L111:
	;
	v833 = v815
	v836 = v818
	v846 = v743
	goto L116
L112:
	;
	v791 = v539 + v775<<(uint(int32(2))%32)
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v791)))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v791)+4))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v791)+8))
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v791)+12))
	v803 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v778, base.F64_convert_i32_s(v792)), base.F64_convert_i32_s(v795)), base.F64_convert_i32_s(v798)), base.F64_convert_i32_s(v801))
	v804 = int32(4)
	v805 = v775 + v804
	v807 = v787 + v804
	if v807 != v761&int32(-4) {
		v775 = v805
		v778 = v803
		v787 = v807
		goto L112
	} else {
		goto L114
	}
L113:
	;
	if v763 == int32(0) {
		v865 = v803
		goto L107
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	v815 = v805
	v818 = v803
	goto L111
L116:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v539+v833<<(uint(int32(2))%32))))
	v852 = base.F64_add(v836, base.F64_convert_i32_s(v850))
	v853 = int32(1)
	v856 = v846 + v853
	if v856 != v763 {
		v833 = v833 + v853
		v836 = v852
		v846 = v856
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v865 = v852
	goto L107
L118:
	;
	goto L117
L119:
	;
	v876 = base.F64_mul(l3, base.F64_neg(v734))
	goto L121
L120:
	;
	v876 = v734
	goto L121
L121:
	;
	v883 = v531
	v889 = v865
	goto L122
L122:
	;
	v900 = float64(1)
	v901 = float64(0)
	v905 = base.F64_sub(base.F64_sub(v900, base.F64_div(v889, v749)), base.F64_promote_f32(v735))
	if base.F64_lt(v905, v901) != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v949 = int32(0)
	goto L105
L124:
	;
	v908 = v901
	goto L126
L125:
	;
	v908 = v905
	goto L126
L126:
	;
	if base.F64_gt(v908, float64(1)) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v911 = v900
	goto L129
L128:
	;
	v911 = v908
	goto L129
L129:
	;
	v913 = v883 - int32(1)
	v915 = base.F64_sub(v876, base.F64_convert_i32_u(v913))
	if base.F64_gt(v915, float64(1)) != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v919 = base.F64_div(v911, v915)
	goto L132
L131:
	;
	v919 = v911
	goto L132
L132:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v539+v913<<(uint(int32(2))%32))))
	v925 = base.F64_convert_i32_s(v924)
	v927 = base.F64_div(base.F64_mul(l3, v925), v749)
	v933 = base.F64_sqrt(base.F64_div(base.F64_mul(base.F64_sub(l3, v749), base.F64_mul(base.F64_mul(v927, v749), base.F64_sub(l3, v927))), base.F64_mul(base.F64_mul(l3, l3), base.F64_add(l3, float64(-1)))))
	if base.F64_lt(base.F64_add(base.F64_add(base.F64_mul(v919, v749), base.F64_add(v933, v933)), float64(0.5)), v925) != 0 {
		v949 = v883
		goto L105
	} else {
		goto L133
	}
L133:
	;
	if v913 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v539+v883<<(uint(int32(2))%32)-int32(8))))
	v883 = v913
	v889 = base.F64_sub(v889, base.F64_convert_i32_s(v944))
	goto L122
L135:
	;
	goto L136
L136:
	;
	goto L123
L137:
	;
	v1002 = int32(_a_F_compute_scalar_stats_1)
	v1003 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1005
	v1009 = F_palloc(m, v971<<(uint(int32(3))%32))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	v1013 = F_palloc(m, v971<<(uint(int32(2))%32))
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	v1019 = int32(0)
	goto L140
L140:
	;
	v1048 = v1019 << (uint(int32(3)) % 32)
	v1050 = v1048 + v62
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1050)+4))
	v1055 = *(*int64)(unsafe.Add(mBase, uint32(v54+v1051<<(uint(int32(4))%32))))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056)+78)))
	v1058 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1056)+76)))
	v1059 = F_datumCopy(m, v1055, v1057, v1058)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L4
	} else {
		goto L142
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1003
	v1075 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v1075)
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v1077
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v1013
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v1080
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v1009
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v971
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v971
	v1091 = v971
	v1118 = v1075
	goto L82
L142:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1009+v1048))) = v1059
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1050)))
	*(*float32)(unsafe.Add(mBase, uint32(v1013+v1019<<(uint(int32(2))%32)))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v1065), v447))
	v1071 = v1019 + int32(1)
	if v1071 != v971 {
		v1019 = v1071
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v1123 = v50 + int32(1)
	goto L146
L145:
	;
	v1123 = v1121
	goto L146
L146:
	;
	if int32(2) <= v1123 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v1126 = int32(0)
	F_qsort_interruptible(m, v62, v1091, int32(8), int32(544), v1126)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L4
	} else {
		goto L150
	}
L148:
	;
	v1326 = v1118
	goto L149
L149:
	;
	if v196 == int32(1) {
		goto L9
	} else {
		goto L175
	}
L150:
	;
	if v1118 != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v1132 = int32(0)
	v1138 = v1126
	v1140 = v1132
	v1141 = v1132
	goto L154
L152:
	;
	v1202 = v196
	goto L153
L153:
	;
	v1228 = int32(_a_F_compute_scalar_stats_1)
	v1229 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1231
	v1233 = int32(1)
	v1234 = v1202 - v1233
	v1236 = v1123 - v1233
	v1237 = base.I32_div_s(v1234, v1236)
	if v1123 <= v1233 {
		goto L164
	} else {
		goto L165
	}
L154:
	;
	if v1091 <= v1141 {
		v1176 = v196
		goto L157
	} else {
		goto L158
	}
L155:
	;
	v1202 = v1190
	goto L153
L156:
	;
	if v1194 < v196 {
		v1138 = v1194
		v1140 = v1190
		v1141 = v1191
		goto L154
	} else {
		goto L163
	}
L157:
	;
	v1178 = v1176 - v1138
	v1180 = v1178 << (uint(int32(4)) % 32)
	if v1180 != 0 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	v1169 = v62 + v1141<<(uint(int32(3))%32)
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1169)+4))
	if v1138 < v1170 {
		v1176 = v1170
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1169)))
	v1190 = v1140
	v1191 = v1141 + int32(1)
	v1194 = v1174 + v1170
	goto L156
L160:
	;
	v1181 = int32(4)
	base.MemoryCopy(m, v54+v1140<<(uint(v1181)%32), v54+v1138<<(uint(v1181)%32), v1180)
	goto L162
L161:
	;
	goto L162
L162:
	;
	v1190 = v1140 + v1178
	v1191 = v1141
	v1194 = v1176
	goto L156
L163:
	;
	goto L155
L164:
	;
	v1243 = v1233
	goto L166
L165:
	;
	v1243 = v1123
	goto L166
L166:
	;
	v1246 = F_palloc(m, v1123<<(uint(int32(3))%32))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	v1248 = int32(0)
	v1255 = v1248
	v1256 = v1248
	v1257 = v1248
	goto L168
L168:
	;
	v1289 = *(*int64)(unsafe.Add(mBase, uint32(v54+v1257<<(uint(int32(4))%32))))
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1290)+78)))
	v1292 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1290)+76)))
	v1293 = F_datumCopy(m, v1289, v1291, v1292)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L4
	} else {
		goto L170
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1229
	v1308 = int32(1)
	v1311 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v1118<<(uint(v1308)%32))+52)) = uint16(v1311)
	v1315 = l0 + v1118<<(uint(v1311)%32)
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1315)+64)) = v1316
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1315)+164)) = v1246
	*(*int32)(unsafe.Add(mBase, uint32(v1315)+84)) = v1318
	*(*int32)(unsafe.Add(mBase, uint32(v1315)+144)) = v1123
	v1326 = v1118 + v1308
	goto L149
L170:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1246+v1255<<(uint(int32(3))%32)))) = v1293
	v1296 = v1256 + (v1234 - v1237*v1236)
	v1297 = base.B2i32(v1236 <= v1296)
	if v1236 <= v1296 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v1301 = v1236
	goto L173
L172:
	;
	v1301 = int32(0)
	goto L173
L173:
	;
	v1304 = v1255 + int32(1)
	if v1304 != v1243 {
		v1255 = v1304
		v1256 = v1296 - v1301
		v1257 = v1297 + (v1257 + v1237)
		goto L168
	} else {
		goto L174
	}
L174:
	;
	goto L169
L175:
	;
	v1358 = int32(_a_F_compute_scalar_stats_1)
	v1359 = *(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0]))
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1361
	v1364 = F_palloc(m, int32(4))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_compute_scalar_stats[0])) = v1359
	v1368 = base.F64_convert_i32_u(v196)
	v1370 = int32(1)
	v1373 = base.F64_mul(v1368, base.F64_convert_i32_u(v196-v1370))
	v1375 = base.F64_mul(v1373, float64(0.5))
	v1377 = base.F64_mul(v1375, base.F64_neg(v1375))
	*(*float32)(unsafe.Add(mBase, uint32(v1364))) = base.F32_demote_f64(base.F64_div(base.F64_add(base.F64_mul(v1368, v440), v1377), base.F64_add(base.F64_mul(v1368, base.F64_div(base.F64_mul(v1373, base.F64_convert_i32_s(v196<<(uint(v1370)%32)-v1370)), float64(6))), v1377)))
	v1395 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v1326<<(uint(v1370)%32))+52)) = uint16(v1395)
	v1399 = l0 + v1326<<(uint(int32(2))%32)
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1399)+64)) = v1400
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1399)+124)) = v1364
	*(*int32)(unsafe.Add(mBase, uint32(v1399)+84)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(v1399)+104)) = v1370
	goto L9
L177:
	;
	v1409 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v1409)
	v1414 = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v195), base.F64_convert_i32_s(l2)))
	*(*float32)(unsafe.Add(mBase, uint32(l0)+40)) = v1414
	if v49 != 0 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	goto L179
L179:
	;
	if v195 <= int32(0) {
		goto L9
	} else {
		goto L183
	}
L180:
	;
	v1421 = base.I32_trunc_sat_f64_s(base.F64_div(v199, base.F64_convert_i32_u(v198)))
	goto L182
L181:
	;
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1420 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1419)+76)))
	v1421 = v1420
	goto L182
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1421
	*(*float32)(unsafe.Add(mBase, uint32(l0)+48)) = base.F32_neg(base.F32_sub(float32(1), v1414))
	goto L9
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(1065353216)
	v1431 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v1431)
	v1433 = int32(0)
	if v49 == v1433 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1437 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1436)+76)))
	v1438 = v1437
	goto L186
L185:
	;
	v1438 = v1433
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1438
	goto L9
}
func F_connectby_text_serial(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int64(0)
	} else {
		v26 = F_text_to_cstring(m, v22)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v29 = F_pg_detoast_datum_packed(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v31 = F_text_to_cstring(m, v29)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v34 = F_pg_detoast_datum_packed(m, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						v36 = F_text_to_cstring(m, v34)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int64(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							v39 = F_pg_detoast_datum_packed(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								v41 = F_text_to_cstring(m, v39)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
									v44 = F_pg_detoast_datum_packed(m, v43)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int64(0)
									} else {
										v46 = F_text_to_cstring(m, v44)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int64(0)
										} else {
											v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v48 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v138 = m.ExcPending
													if v138 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_connectby_text_serial_0), int32(0))
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_connectby_text_serial_1), int32(1078), int32(_a_F_connectby_text_serial_2))
															mBase = m.M
															v147 = m.ExcPending
															if v147 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
												if v51 != int32(389) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v135 = m.ExcPending
													if v135 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v138 = m.ExcPending
														if v138 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_connectby_text_serial_0), int32(0))
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_connectby_text_serial_1), int32(1078), int32(_a_F_connectby_text_serial_2))
																mBase = m.M
																v147 = m.ExcPending
																if v147 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+12)))
													if v54&int32(2) == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(1088))
															mBase = m.M
															v154 = m.ExcPending
															if v154 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_connectby_text_serial_3), int32(0))
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_connectby_text_serial_1), int32(1083), int32(_a_F_connectby_text_serial_2))
																	mBase = m.M
																	v163 = m.ExcPending
																	if v163 != 0 {
																		return int64(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														v59 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
														if v59 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(1088))
																mBase = m.M
																v154 = m.ExcPending
																if v154 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_connectby_text_serial_3), int32(0))
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_connectby_text_serial_1), int32(1083), int32(_a_F_connectby_text_serial_2))
																		mBase = m.M
																		v163 = m.ExcPending
																		if v163 != 0 {
																			return int64(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															}
														} else {
															v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
															v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
															if v63 == int32(7) {
																v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
																v67 = F_pg_detoast_datum_packed(m, v66)
																mBase = m.M
																v68 = m.ExcPending
																if v68 != 0 {
																	return int64(0)
																} else {
																	v69 = F_text_to_cstring(m, v67)
																	mBase = m.M
																	v70 = m.ExcPending
																	if v70 != 0 {
																		return int64(0)
																	} else {
																		v74 = v69
																		v75 = int32(_a_F_connectby_text_serial_4)
																		v76 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0]))
																		v78 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
																		v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
																		*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v79
																		v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
																		v82 = F_CreateTupleDescCopy(m, v81)
																		mBase = m.M
																		v83 = m.ExcPending
																		if v83 != 0 {
																			return int64(0)
																		} else {
																			v85 = base.B2i32(v63 == int32(7))
																			F_validateConnectbyTupleDesc(m, v82, v85, int32(1))
																			mBase = m.M
																			v88 = m.ExcPending
																			if v88 != 0 {
																				return int64(0)
																			} else {
																				v89 = F_TupleDescGetAttInMetadata(m, v82)
																				mBase = m.M
																				v90 = m.ExcPending
																				if v90 != 0 {
																					return int64(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = int32(2)
																					v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
																					*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(1)
																					F_SPI_connect_ext(m, int32(0))
																					mBase = m.M
																					v98 = m.ExcPending
																					if v98 != 0 {
																						return int64(0)
																					} else {
																						v99 = int32(_a_F_connectby_text_serial_4)
																						v100 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0]))
																						*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v79
																						v109 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[1]))
																						v110 = F_tuplestore_begin_heap(m, int32(base.Ui32(v93&int32(4))>>(uint(int32(2))%32)), int32(0), v109)
																						mBase = m.M
																						v111 = m.ExcPending
																						if v111 != 0 {
																							return int64(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v100
																							F_build_tuplestore_recursively(m, v31, v36, v26, v41, v74, v46, v46, int32(0), v19+int32(12), base.I32_wrap_i64(v62), v85, int32(1), v89, v110)
																							mBase = m.M
																							v120 = m.ExcPending
																							if v120 != 0 {
																								return int64(0)
																							} else {
																								v121 = F_SPI_finish(m)
																								mBase = m.M
																								v122 = m.ExcPending
																								if v122 != 0 {
																									return int64(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = v82
																									*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v110
																									*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v76
																									m.G0 = v19 + int32(16)
																									return int64(0)
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
																v72 = F_pstrdup(m, int32(_a_F_connectby_text_serial_5))
																mBase = m.M
																v73 = m.ExcPending
																if v73 != 0 {
																	return int64(0)
																} else {
																	v74 = v72
																	v75 = int32(_a_F_connectby_text_serial_4)
																	v76 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0]))
																	v78 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
																	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
																	*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v79
																	v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
																	v82 = F_CreateTupleDescCopy(m, v81)
																	mBase = m.M
																	v83 = m.ExcPending
																	if v83 != 0 {
																		return int64(0)
																	} else {
																		v85 = base.B2i32(v63 == int32(7))
																		F_validateConnectbyTupleDesc(m, v82, v85, int32(1))
																		mBase = m.M
																		v88 = m.ExcPending
																		if v88 != 0 {
																			return int64(0)
																		} else {
																			v89 = F_TupleDescGetAttInMetadata(m, v82)
																			mBase = m.M
																			v90 = m.ExcPending
																			if v90 != 0 {
																				return int64(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = int32(2)
																				v93 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
																				*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(1)
																				F_SPI_connect_ext(m, int32(0))
																				mBase = m.M
																				v98 = m.ExcPending
																				if v98 != 0 {
																					return int64(0)
																				} else {
																					v99 = int32(_a_F_connectby_text_serial_4)
																					v100 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0]))
																					*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v79
																					v109 = *(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[1]))
																					v110 = F_tuplestore_begin_heap(m, int32(base.Ui32(v93&int32(4))>>(uint(int32(2))%32)), int32(0), v109)
																					mBase = m.M
																					v111 = m.ExcPending
																					if v111 != 0 {
																						return int64(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v100
																						F_build_tuplestore_recursively(m, v31, v36, v26, v41, v74, v46, v46, int32(0), v19+int32(12), base.I32_wrap_i64(v62), v85, int32(1), v89, v110)
																						mBase = m.M
																						v120 = m.ExcPending
																						if v120 != 0 {
																							return int64(0)
																						} else {
																							v121 = F_SPI_finish(m)
																							mBase = m.M
																							v122 = m.ExcPending
																							if v122 != 0 {
																								return int64(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = v82
																								*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v110
																								*(*int32)(unsafe.Add(mBase, _c_F_connectby_text_serial[0])) = v76
																								m.G0 = v19 + int32(16)
																								return int64(0)
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
				}
			}
		}
	}
}
func F_contains_multiexpr_param(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 == int32(8) {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			return base.B2i32(v10 == int32(3))
		} else {
			v15 = F_expression_tree_walker_impl(m, l0, int32(1134), l1)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_contsel(m *base.Module, l0 int32) int64 {
	return int64(4562254508917369340)
}
func F_convert_case(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v374 int32
	_ = v374
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v471 int32
	_ = v471
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v579 int32
	_ = v579
	var v602 int32
	_ = v602
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v646 int32
	_ = v646
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v681 int32
	_ = v681
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v703 int32
	_ = v703
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v781 int32
	_ = v781
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v820 int32
	_ = v820
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v889 int32
	_ = v889
	var v896 int32
	_ = v896
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1016 int32
	_ = v1016
	var v1023 int32
	_ = v1023
	var v1030 int32
	_ = v1030
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1091 int32
	_ = v1091
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1135 int32
	_ = v1135
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1154 int32
	_ = v1154
	var v1161 int32
	_ = v1161
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1193 int32
	_ = v1193
	var v1219 int32
	_ = v1219
	var v1237 int32
	_ = v1237
	v9 = int32(0)
	if l4 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = m.T0[l6].(func(*base.Module, int32) int32)(m, l7)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v32 = v9
	goto L3
L3:
	;
	if l3 == int32(0) {
		v1219 = v9
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v32 = v28
	goto L3
L6:
	;
	if base.Ui32(v1219) < base.Ui32(l1) {
		goto L309
	} else {
		goto L310
	}
L7:
	;
	if l5 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = int32(1)
	goto L10
L9:
	;
	v37 = int32(2)
	goto L10
L10:
	;
	v47 = v9
	v52 = l4
	v53 = v9
	v54 = v32
	goto L11
L11:
	;
	v63 = l2 + v53
	v64 = int32(*(*int8)(unsafe.Add(mBase, uint32(v63))))
	v66 = v64 & int32(255)
	if v64 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v1219 = v1193
	goto L6
L13:
	;
	if l4 != int32(1) {
		v153 = v52
		v154 = v54
		goto L33
	} else {
		goto L34
	}
L14:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v63))))
	v143 = v134
	v144 = v137 | v139&int32(63)
	v146 = v136
	goto L13
L15:
	;
	if v66&int32(248) != int32(240) {
		goto L30
	} else {
		goto L31
	}
L16:
	;
	if v66&int32(224) != int32(192) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if base.Ui32(l3) <= base.Ui32(v53) {
		v1219 = v47
		goto L6
	} else {
		goto L29
	}
L19:
	;
	v74 = v66 & int32(240)
	if v74 == int32(224) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v100 = v53 + int32(2)
	if base.Ui32(l3) < base.Ui32(v100) {
		v1219 = v47
		goto L6
	} else {
		goto L28
	}
L22:
	;
	v83 = int32(3)
	goto L24
L23:
	;
	if v66&int32(248) != int32(240) {
		v1219 = v47
		goto L6
	} else {
		goto L25
	}
L24:
	;
	v84 = v83 + v53
	if base.Ui32(l3) < base.Ui32(v84) {
		v1219 = v47
		goto L6
	} else {
		goto L26
	}
L25:
	;
	v83 = int32(4)
	goto L24
L26:
	;
	if v74 != int32(224) {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v134 = v83
	v135 = int32(2)
	v136 = v84
	v137 = v66<<(uint(int32(12))%32)&int32(_a_F_convert_case_0) | v93&int32(63)<<(uint(int32(6))%32)
	goto L14
L28:
	;
	v134 = int32(2)
	v135 = int32(1)
	v136 = v100
	v137 = v66 << (uint(int32(6)) % 32) & int32(1984)
	goto L14
L29:
	;
	v109 = int32(1)
	v143 = v109
	v144 = v66
	v146 = v53 + v109
	goto L13
L30:
	;
	v143 = v83
	v144 = int32(-1)
	v146 = v84
	goto L13
L31:
	;
	goto L32
L32:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v123 = int32(63)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
	v134 = v83
	v135 = int32(3)
	v136 = v84
	v137 = v66<<(uint(int32(18))%32)&int32(_a_F_convert_case_1) | v122&v123<<(uint(int32(12))%32) | v128&v123<<(uint(int32(6))%32)
	goto L14
L33:
	;
	if base.Ui32(v144) <= base.Ui32(int32(127)) {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	if v53 != v54 {
		v153 = int32(0)
		v154 = v54
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v151 = m.T0[l6].(func(*base.Module, int32) int32)(m, l7)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v153 = v37
	v154 = v151
	goto L33
L37:
	;
	if base.Ui32(v146) < base.Ui32(l3) {
		v47 = v1193
		v52 = v153
		v53 = v146
		v54 = v154
		goto L11
	} else {
		goto L308
	}
L38:
	;
	v1178 = v143 + v47
	if base.Ui32(l1) < base.Ui32(v1178) {
		goto L302
	} else {
		goto L303
	}
L39:
	;
	v1071 = int32(0)
	v1072 = v47
	goto L278
L40:
	;
	v974 = l0 + v47
	if base.Ui32(v955) <= base.Ui32(int32(2047)) {
		goto L272
	} else {
		goto L273
	}
L41:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v954)))
	if base.Ui32(int32(128)) <= base.Ui32(v955) {
		goto L259
	} else {
		goto L260
	}
L42:
	;
	v157 = int32(2)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153<<(uint(v157)%32))+uint32(_c_F_convert_case[0])))
	v954 = v159 + v144<<(uint(v157)%32) + int32(4)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v165 = int32(0)
	if base.Ui32(v144) <= base.Ui32(int32(1415)) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	if v297 == int32(0) {
		goto L38
	} else {
		goto L102
	}
L46:
	;
	goto L45
L47:
	;
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[1]))))
	v297 = v170
	goto L46
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(v144) <= base.Ui32(int32(_a_F_convert_case_2)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if base.Ui32(v144) <= base.Ui32(int32(_a_F_convert_case_3)) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(v144) < base.Ui32(int32(_a_F_convert_case_4)) {
		v297 = v165
		goto L46
	} else {
		goto L77
	}
L53:
	;
	if base.Ui32(v144-int32(_a_F_convert_case_5)) <= base.Ui32(int32(95)) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(v144) < base.Ui32(int32(_a_F_convert_case_6)) {
		v297 = v165
		goto L46
	} else {
		goto L64
	}
L56:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[2]))))
	v297 = v183
	goto L46
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(v144) < base.Ui32(int32(_a_F_convert_case_7)) {
		v297 = v165
		goto L46
	} else {
		goto L59
	}
L59:
	;
	if base.Ui32(v144) <= base.Ui32(int32(_a_F_convert_case_8)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[3]))))
	v297 = v192
	goto L46
L61:
	;
	goto L62
L62:
	;
	if base.Ui32(v144) < base.Ui32(int32(_a_F_convert_case_9)) {
		v297 = v165
		goto L46
	} else {
		goto L63
	}
L63:
	;
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[4]))))
	v297 = v199
	goto L46
L64:
	;
	if base.Ui32(v144) <= base.Ui32(int32(_a_F_convert_case_10)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	if base.Ui32(v144) <= base.Ui32(int32(_a_F_convert_case_11)) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if base.Ui32(v144) < base.Ui32(int32(_a_F_convert_case_12)) {
		v297 = v165
		goto L46
	} else {
		goto L72
	}
L68:
	;
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[5]))))
	v297 = v210
	goto L46
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(v144) < base.Ui32(int32(_a_F_convert_case_13)) {
		v297 = v165
		goto L46
	} else {
		goto L71
	}
L71:
	;
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[6]))))
	v297 = v217
	goto L46
L72:
	;
	if base.Ui32(v144) <= base.Ui32(int32(_a_F_convert_case_14)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[7]))))
	v297 = v226
	goto L46
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(v144) < base.Ui32(int32(_a_F_convert_case_15)) {
		v297 = v165
		goto L46
	} else {
		goto L76
	}
L76:
	;
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[8]))))
	v297 = v233
	goto L46
L77:
	;
	if base.Ui32(v144) <= base.Ui32(int32(_a_F_convert_case_16)) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if base.Ui32(v144) <= base.Ui32(int32(_a_F_convert_case_17)) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	if base.Ui32(v144) < base.Ui32(int32(_a_F_convert_case_18)) {
		v297 = v165
		goto L46
	} else {
		goto L93
	}
L81:
	;
	if base.Ui32(v144) <= base.Ui32(int32(_a_F_convert_case_19)) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	if base.Ui32(v144) < base.Ui32(int32(_a_F_convert_case_20)) {
		v297 = v165
		goto L46
	} else {
		goto L88
	}
L84:
	;
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[9]))))
	v297 = v246
	goto L46
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(v144) < base.Ui32(int32(_a_F_convert_case_21)) {
		v297 = v165
		goto L46
	} else {
		goto L87
	}
L87:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[10]))))
	v297 = v253
	goto L46
L88:
	;
	if base.Ui32(v144) <= base.Ui32(int32(_a_F_convert_case_22)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[11]))))
	v297 = v262
	goto L46
L90:
	;
	goto L91
L91:
	;
	if base.Ui32(v144) < base.Ui32(int32(_a_F_convert_case_23)) {
		v297 = v165
		goto L46
	} else {
		goto L92
	}
L92:
	;
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[12]))))
	v297 = v269
	goto L46
L93:
	;
	if base.Ui32(v144) <= base.Ui32(int32(_a_F_convert_case_24)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	if base.Ui32(v144) <= base.Ui32(int32(_a_F_convert_case_25)) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	if base.Ui32(int32(67)) < base.Ui32(v144-int32(_a_F_convert_case_26)) {
		v297 = v165
		goto L46
	} else {
		goto L101
	}
L97:
	;
	v280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[13]))))
	v297 = v280
	goto L46
L98:
	;
	goto L99
L99:
	;
	if base.Ui32(v144) < base.Ui32(int32(_a_F_convert_case_27)) {
		v297 = v165
		goto L46
	} else {
		goto L100
	}
L100:
	;
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[14]))))
	v297 = v287
	goto L46
L101:
	;
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144<<(uint(int32(1))%32))+uint32(_c_F_convert_case[15]))))
	v297 = v296
	goto L46
L102:
	;
	if l5 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v923 = int32(2)
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v153<<(uint(v923)%32))+uint32(_c_F_convert_case[0])))
	v954 = v925 + v297<<(uint(v923)%32)
	goto L41
L104:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+uint32(_c_F_convert_case[16]))))
	if v302 == int32(0) {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v306 = v302 * int32(52)
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v306)+uint32(_c_F_convert_case[17]))))
	switch v307 {
	case 0:
		goto L39
	case 1:
		goto L106
	default:
		goto L103
	}
L106:
	;
	v308 = int32(0)
	if v53 == v308 {
		v602 = v308
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v610 = int32(*(*int8)(unsafe.Add(mBase, uint32(v63))))
	if int32(0) <= v610 {
		v633 = int32(1)
		goto L184
	} else {
		goto L185
	}
L108:
	;
	v319 = v53
	goto L109
L109:
	;
	v337 = v319 - int32(1)
	v338 = l2 + v337
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	v340 = base.I32_extend8_s(v339)
	if v340 <= int32(-65) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v602 = v308
	goto L107
L111:
	;
	if v337 != 0 {
		v319 = v337
		goto L109
	} else {
		goto L183
	}
L112:
	;
	if v340 < int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	if base.Ui32(int32(127)) < base.Ui32(v422) {
		goto L134
	} else {
		goto L135
	}
L114:
	;
	if v339&int32(224) != int32(192) {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	goto L116
L116:
	;
	if base.Ui32(l3) < base.Ui32(v319) {
		goto L103
	} else {
		goto L132
	}
L117:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414+v338))))
	v422 = v416&int32(63) | v410
	goto L113
L118:
	;
	if v353 != int32(240) {
		v422 = int32(-1)
		goto L113
	} else {
		goto L131
	}
L119:
	;
	v353 = v339 & int32(248)
	if v353 == int32(240) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	if base.Ui32(l3) < base.Ui32(v319+int32(1)) {
		goto L103
	} else {
		goto L130
	}
L122:
	;
	v356 = int32(4)
	goto L124
L123:
	;
	v356 = int32(-1)
	goto L124
L124:
	;
	v360 = base.B2i32(v339&int32(240) == int32(224))
	if v339&int32(240) == int32(224) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v361 = int32(3)
	goto L127
L126:
	;
	v361 = v356
	goto L127
L127:
	;
	if base.B2i32(v361 < int32(0))|base.B2i32(base.Ui32(l3) < base.Ui32(v337+v361)) != 0 {
		goto L103
	} else {
		goto L128
	}
L128:
	;
	if v360 == int32(0) {
		goto L118
	} else {
		goto L129
	}
L129:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v319))))
	v410 = v339<<(uint(int32(12))%32)&int32(_a_F_convert_case_0) | v374&int32(63)<<(uint(int32(6))%32)
	v414 = int32(2)
	goto L117
L130:
	;
	v410 = v339 << (uint(int32(6)) % 32) & int32(1984)
	v414 = int32(1)
	goto L117
L131:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v319))))
	v398 = int32(63)
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+2)))
	v410 = v339<<(uint(int32(18))%32)&int32(_a_F_convert_case_1) | v397&v398<<(uint(int32(12))%32) | v403&v398<<(uint(int32(6))%32)
	v414 = int32(3)
	goto L117
L132:
	;
	v422 = v339
	goto L113
L133:
	;
	if v471 != 0 {
		goto L111
	} else {
		goto L147
	}
L134:
	;
	v433 = int32(517)
	v434 = int32(0)
	goto L137
L135:
	;
	goto L136
L136:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422<<(uint(int32(1))%32))+uint32(_c_F_convert_case[18]))))
	v471 = int32(base.Ui32(v461&int32(16)) >> (uint(int32(4)) % 32))
	goto L133
L137:
	;
	v439 = base.I32_div_s(v433+v434, int32(2))
	v441 = v439 << (uint(int32(3)) % 32)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v441)+uint32(_c_F_convert_case[19])))
	if base.Ui32(v444) < base.Ui32(v422) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v471 = int32(0)
	goto L133
L139:
	;
	if v456 <= v455 {
		v433 = v455
		v434 = v456
		goto L137
	} else {
		goto L146
	}
L140:
	;
	v455 = v433
	v456 = v439 + int32(1)
	goto L139
L141:
	;
	goto L142
L142:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v441)+uint32(_c_F_convert_case[20])))
	if base.Ui32(v450) <= base.Ui32(v422) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v471 = int32(1)
	goto L133
L144:
	;
	goto L145
L145:
	;
	v455 = v439 - int32(1)
	v456 = v434
	goto L139
L146:
	;
	goto L138
L147:
	;
	if base.Ui32(int32(127)) < base.Ui32(v422) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v602 = v579
	goto L107
L149:
	;
	v520 = int32(691)
	v521 = int32(0)
	goto L163
L150:
	;
	v480 = int32(3408)
	v481 = int32(0)
	goto L153
L151:
	;
	goto L152
L152:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422<<(uint(int32(1))%32))+uint32(_c_F_convert_case[18]))))
	v579 = int32(base.Ui32(v510&int32(8)) >> (uint(int32(3)) % 32))
	goto L148
L153:
	;
	v486 = base.I32_div_s(v480+v481, int32(2))
	v488 = v486 * int32(12)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v488)+uint32(_c_F_convert_case[21])))
	if base.Ui32(v491) < base.Ui32(v422) {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+uint32(_c_F_convert_case[22]))))
	if v504 != int32(3) {
		goto L149
	} else {
		goto L162
	}
L155:
	;
	goto L154
L156:
	;
	if v502 <= v501 {
		v480 = v501
		v481 = v502
		goto L153
	} else {
		goto L161
	}
L157:
	;
	v501 = v480
	v502 = v486 + int32(1)
	goto L156
L158:
	;
	goto L159
L159:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v488)+uint32(_c_F_convert_case[23])))
	if base.Ui32(v497) <= base.Ui32(v422) {
		goto L155
	} else {
		goto L160
	}
L160:
	;
	v501 = v486 - int32(1)
	v502 = v481
	goto L156
L161:
	;
	goto L149
L162:
	;
	v579 = int32(1)
	goto L148
L163:
	;
	v526 = base.I32_div_s(v520+v521, int32(2))
	v528 = v526 << (uint(int32(3)) % 32)
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v528)+uint32(_c_F_convert_case[24])))
	if base.Ui32(v531) < base.Ui32(v422) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v548 = int32(659)
	v549 = int32(0)
	goto L173
L165:
	;
	if v543 <= v542 {
		v520 = v542
		v521 = v543
		goto L163
	} else {
		goto L172
	}
L166:
	;
	v542 = v520
	v543 = v526 + int32(1)
	goto L165
L167:
	;
	goto L168
L168:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v528)+uint32(_c_F_convert_case[25])))
	if base.Ui32(v537) <= base.Ui32(v422) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v579 = int32(1)
	goto L148
L170:
	;
	goto L171
L171:
	;
	v542 = v526 - int32(1)
	v543 = v521
	goto L165
L172:
	;
	goto L164
L173:
	;
	v554 = base.I32_div_s(v548+v549, int32(2))
	v556 = v554 << (uint(int32(3)) % 32)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v556)+uint32(_c_F_convert_case[26])))
	if base.Ui32(v559) < base.Ui32(v422) {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v579 = int32(0)
	goto L148
L175:
	;
	if v571 <= v570 {
		v548 = v570
		v549 = v571
		goto L173
	} else {
		goto L182
	}
L176:
	;
	v570 = v548
	v571 = v554 + int32(1)
	goto L175
L177:
	;
	goto L178
L178:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v556)+uint32(_c_F_convert_case[27])))
	if base.Ui32(v565) <= base.Ui32(v422) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v579 = int32(1)
	goto L148
L180:
	;
	goto L181
L181:
	;
	v570 = v554 - int32(1)
	v571 = v549
	goto L175
L182:
	;
	goto L174
L183:
	;
	goto L110
L184:
	;
	v646 = v633 + v53
	goto L192
L185:
	;
	v615 = v610 & int32(255)
	if v615&int32(224) == int32(192) {
		v633 = int32(2)
		goto L184
	} else {
		goto L186
	}
L186:
	;
	if v615&int32(240) == int32(224) {
		v633 = int32(3)
		goto L184
	} else {
		goto L187
	}
L187:
	;
	if v615&int32(248) == int32(240) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v631 = int32(4)
	goto L190
L189:
	;
	v631 = int32(-1)
	goto L190
L190:
	;
	v633 = v631
	goto L184
L191:
	;
	if v896&v602 != 0 {
		goto L39
	} else {
		goto L258
	}
L192:
	;
	if base.Ui32(l3) <= base.Ui32(v646) {
		v896 = int32(1)
		goto L191
	} else {
		goto L194
	}
L193:
	;
	if base.Ui32(int32(127)) < base.Ui32(v732) {
		goto L225
	} else {
		goto L226
	}
L194:
	;
	v662 = l2 + v646
	v663 = int32(*(*int8)(unsafe.Add(mBase, uint32(v662))))
	v665 = v663 & int32(255)
	if v663 < int32(0) {
		goto L199
	} else {
		goto L200
	}
L195:
	;
	if base.Ui32(int32(127)) < base.Ui32(v732) {
		goto L209
	} else {
		goto L210
	}
L196:
	;
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662+v725))))
	v732 = v726 | v728&int32(63)
	v733 = v724
	goto L195
L197:
	;
	v711 = v646 + int32(3)
	if base.Ui32(l3) < base.Ui32(v711) {
		goto L103
	} else {
		goto L207
	}
L198:
	;
	v703 = v646 + int32(2)
	if base.Ui32(l3) < base.Ui32(v703) {
		goto L103
	} else {
		goto L206
	}
L199:
	;
	if v665&int32(224) == int32(192) {
		goto L198
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v732 = v665
	v733 = v646 + int32(1)
	goto L195
L202:
	;
	if v665&int32(240) == int32(224) {
		goto L197
	} else {
		goto L203
	}
L203:
	;
	if v665&int32(248) != int32(240) {
		goto L103
	} else {
		goto L204
	}
L204:
	;
	v681 = v646 + int32(4)
	if base.Ui32(l3) < base.Ui32(v681) {
		goto L103
	} else {
		goto L205
	}
L205:
	;
	v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+1)))
	v689 = int32(63)
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+2)))
	v724 = v681
	v725 = int32(3)
	v726 = v665<<(uint(int32(18))%32)&int32(_a_F_convert_case_1) | v688&v689<<(uint(int32(12))%32) | v694&v689<<(uint(int32(6))%32)
	goto L196
L206:
	;
	v724 = v703
	v725 = int32(1)
	v726 = v665 << (uint(int32(6)) % 32) & int32(1984)
	goto L196
L207:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662)+1)))
	v724 = v711
	v725 = int32(2)
	v726 = v665<<(uint(int32(12))%32)&int32(_a_F_convert_case_0) | v718&int32(63)<<(uint(int32(6))%32)
	goto L196
L208:
	;
	if v781 != 0 {
		v646 = v733
		goto L192
	} else {
		goto L222
	}
L209:
	;
	v743 = int32(517)
	v744 = int32(0)
	goto L212
L210:
	;
	goto L211
L211:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732<<(uint(int32(1))%32))+uint32(_c_F_convert_case[18]))))
	v781 = int32(base.Ui32(v771&int32(16)) >> (uint(int32(4)) % 32))
	goto L208
L212:
	;
	v749 = base.I32_div_s(v743+v744, int32(2))
	v751 = v749 << (uint(int32(3)) % 32)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v751)+uint32(_c_F_convert_case[19])))
	if base.Ui32(v754) < base.Ui32(v732) {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	v781 = int32(0)
	goto L208
L214:
	;
	if v766 <= v765 {
		v743 = v765
		v744 = v766
		goto L212
	} else {
		goto L221
	}
L215:
	;
	v765 = v743
	v766 = v749 + int32(1)
	goto L214
L216:
	;
	goto L217
L217:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v751)+uint32(_c_F_convert_case[20])))
	if base.Ui32(v760) <= base.Ui32(v732) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v781 = int32(1)
	goto L208
L219:
	;
	goto L220
L220:
	;
	v765 = v749 - int32(1)
	v766 = v744
	goto L214
L221:
	;
	goto L213
L222:
	;
	goto L193
L223:
	;
	v896 = v889 ^ int32(1)
	goto L191
L224:
	;
	v830 = int32(691)
	v831 = int32(0)
	goto L238
L225:
	;
	v790 = int32(3408)
	v791 = int32(0)
	goto L228
L226:
	;
	goto L227
L227:
	;
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v732<<(uint(int32(1))%32))+uint32(_c_F_convert_case[18]))))
	v889 = int32(base.Ui32(v820&int32(8)) >> (uint(int32(3)) % 32))
	goto L223
L228:
	;
	v796 = base.I32_div_s(v790+v791, int32(2))
	v798 = v796 * int32(12)
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v798)+uint32(_c_F_convert_case[21])))
	if base.Ui32(v801) < base.Ui32(v732) {
		goto L232
	} else {
		goto L233
	}
L229:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798)+uint32(_c_F_convert_case[22]))))
	if v814 != int32(3) {
		goto L224
	} else {
		goto L237
	}
L230:
	;
	goto L229
L231:
	;
	if v812 <= v811 {
		v790 = v811
		v791 = v812
		goto L228
	} else {
		goto L236
	}
L232:
	;
	v811 = v790
	v812 = v796 + int32(1)
	goto L231
L233:
	;
	goto L234
L234:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v798)+uint32(_c_F_convert_case[23])))
	if base.Ui32(v807) <= base.Ui32(v732) {
		goto L230
	} else {
		goto L235
	}
L235:
	;
	v811 = v796 - int32(1)
	v812 = v791
	goto L231
L236:
	;
	goto L224
L237:
	;
	v889 = int32(1)
	goto L223
L238:
	;
	v836 = base.I32_div_s(v830+v831, int32(2))
	v838 = v836 << (uint(int32(3)) % 32)
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v838)+uint32(_c_F_convert_case[24])))
	if base.Ui32(v841) < base.Ui32(v732) {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	v858 = int32(659)
	v859 = int32(0)
	goto L248
L240:
	;
	if v853 <= v852 {
		v830 = v852
		v831 = v853
		goto L238
	} else {
		goto L247
	}
L241:
	;
	v852 = v830
	v853 = v836 + int32(1)
	goto L240
L242:
	;
	goto L243
L243:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v838)+uint32(_c_F_convert_case[25])))
	if base.Ui32(v847) <= base.Ui32(v732) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v889 = int32(1)
	goto L223
L245:
	;
	goto L246
L246:
	;
	v852 = v836 - int32(1)
	v853 = v831
	goto L240
L247:
	;
	goto L239
L248:
	;
	v864 = base.I32_div_s(v858+v859, int32(2))
	v866 = v864 << (uint(int32(3)) % 32)
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v866)+uint32(_c_F_convert_case[26])))
	if base.Ui32(v869) < base.Ui32(v732) {
		goto L251
	} else {
		goto L252
	}
L249:
	;
	v889 = int32(0)
	goto L223
L250:
	;
	if v881 <= v880 {
		v858 = v880
		v859 = v881
		goto L248
	} else {
		goto L257
	}
L251:
	;
	v880 = v858
	v881 = v864 + int32(1)
	goto L250
L252:
	;
	goto L253
L253:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v866)+uint32(_c_F_convert_case[27])))
	if base.Ui32(v875) <= base.Ui32(v732) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v889 = int32(1)
	goto L223
L255:
	;
	goto L256
L256:
	;
	v880 = v864 - int32(1)
	v881 = v859
	goto L250
L257:
	;
	goto L249
L258:
	;
	goto L103
L259:
	;
	if base.Ui32(v955) < base.Ui32(int32(_a_F_convert_case_28)) {
		goto L262
	} else {
		goto L263
	}
L260:
	;
	goto L261
L261:
	;
	v970 = v47 + int32(1)
	if base.Ui32(l1) < base.Ui32(v970) {
		goto L269
	} else {
		goto L270
	}
L262:
	;
	v963 = int32(3)
	goto L264
L263:
	;
	v963 = int32(4)
	goto L264
L264:
	;
	if base.Ui32(v955) < base.Ui32(int32(2048)) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v966 = int32(2)
	goto L267
L266:
	;
	v966 = v963
	goto L267
L267:
	;
	v967 = v966 + v47
	if base.Ui32(v967) <= base.Ui32(l1) {
		goto L40
	} else {
		goto L268
	}
L268:
	;
	v1193 = v967
	goto L37
L269:
	;
	v1193 = v970
	goto L37
L270:
	;
	goto L271
L271:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v47))) = uint8(v955)
	v1193 = v970
	goto L37
L272:
	;
	v980 = v955&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v974)+1)) = uint8(v980)
	v985 = int32(base.Ui32(v955)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v974))) = uint8(v985)
	v1193 = v967
	goto L37
L273:
	;
	goto L274
L274:
	;
	if base.Ui32(v955) <= base.Ui32(int32(_a_F_convert_case_29)) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v989 = int32(63)
	v991 = int32(128)
	v992 = v955&v989 | v991
	*(*uint8)(unsafe.Add(mBase, uint32(v974)+2)) = uint8(v992)
	v997 = int32(base.Ui32(v955)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v974))) = uint8(v997)
	v1004 = int32(base.Ui32(v955)>>(uint(int32(6))%32))&v989 | v991
	*(*uint8)(unsafe.Add(mBase, uint32(v974)+1)) = uint8(v1004)
	v1193 = v967
	goto L37
L276:
	;
	goto L277
L277:
	;
	v1006 = int32(63)
	v1008 = int32(128)
	v1009 = v955&v1006 | v1008
	*(*uint8)(unsafe.Add(mBase, uint32(v974)+3)) = uint8(v1009)
	v1016 = int32(base.Ui32(v955)>>(uint(int32(6))%32))&v1006 | v1008
	*(*uint8)(unsafe.Add(mBase, uint32(v974)+2)) = uint8(v1016)
	v1023 = int32(base.Ui32(v955)>>(uint(int32(12))%32))&v1006 | v1008
	*(*uint8)(unsafe.Add(mBase, uint32(v974)+1)) = uint8(v1023)
	v1030 = int32(base.Ui32(v955)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v974))) = uint8(v1030)
	v1193 = v967
	goto L37
L278:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v153*int32(12)+v306+int32(_a_F_convert_case_30)+v1071<<(uint(int32(2))%32))))
	if v1091 == int32(0) {
		v1193 = v1072
		goto L37
	} else {
		goto L280
	}
L279:
	;
	v1193 = v1173
	goto L37
L280:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v1091) {
		goto L284
	} else {
		goto L285
	}
L281:
	;
	v1175 = v1071 + int32(1)
	if v1175 != int32(3) {
		v1071 = v1175
		v1072 = v1173
		goto L278
	} else {
		goto L301
	}
L282:
	;
	v1173 = v1105
	goto L281
L283:
	;
	if base.Ui32(v1091) <= base.Ui32(int32(_a_F_convert_case_29)) {
		goto L298
	} else {
		goto L299
	}
L284:
	;
	if base.Ui32(v1091) < base.Ui32(int32(_a_F_convert_case_28)) {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	goto L286
L286:
	;
	v1121 = v1072 + int32(1)
	if base.Ui32(l1) < base.Ui32(v1121) {
		goto L295
	} else {
		goto L296
	}
L287:
	;
	v1101 = int32(3)
	goto L289
L288:
	;
	v1101 = int32(4)
	goto L289
L289:
	;
	if base.Ui32(v1091) < base.Ui32(int32(2048)) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1104 = int32(2)
	goto L292
L291:
	;
	v1104 = v1101
	goto L292
L292:
	;
	v1105 = v1104 + v1072
	if base.Ui32(l1) < base.Ui32(v1105) {
		goto L282
	} else {
		goto L293
	}
L293:
	;
	v1107 = l0 + v1072
	if base.Ui32(int32(2047)) < base.Ui32(v1091) {
		goto L283
	} else {
		goto L294
	}
L294:
	;
	v1113 = v1091&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v1107)+1)) = uint8(v1113)
	v1118 = int32(base.Ui32(v1091)>>(uint(int32(6))%32)) | int32(192)
	*(*uint8)(unsafe.Add(mBase, uint32(v1107))) = uint8(v1118)
	goto L282
L295:
	;
	v1173 = v1121
	goto L281
L296:
	;
	goto L297
L297:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v1072))) = uint8(v1091)
	v1173 = v1121
	goto L281
L298:
	;
	v1127 = int32(63)
	v1129 = int32(128)
	v1130 = v1091&v1127 | v1129
	*(*uint8)(unsafe.Add(mBase, uint32(v1107)+2)) = uint8(v1130)
	v1135 = int32(base.Ui32(v1091)>>(uint(int32(12))%32)) | int32(224)
	*(*uint8)(unsafe.Add(mBase, uint32(v1107))) = uint8(v1135)
	v1142 = int32(base.Ui32(v1091)>>(uint(int32(6))%32))&v1127 | v1129
	*(*uint8)(unsafe.Add(mBase, uint32(v1107)+1)) = uint8(v1142)
	goto L282
L299:
	;
	goto L300
L300:
	;
	v1144 = int32(63)
	v1146 = int32(128)
	v1147 = v1091&v1144 | v1146
	*(*uint8)(unsafe.Add(mBase, uint32(v1107)+3)) = uint8(v1147)
	v1154 = int32(base.Ui32(v1091)>>(uint(int32(6))%32))&v1144 | v1146
	*(*uint8)(unsafe.Add(mBase, uint32(v1107)+2)) = uint8(v1154)
	v1161 = int32(base.Ui32(v1091)>>(uint(int32(12))%32))&v1144 | v1146
	*(*uint8)(unsafe.Add(mBase, uint32(v1107)+1)) = uint8(v1161)
	v1168 = int32(base.Ui32(v1091)>>(uint(int32(18))%32))&int32(7) | int32(240)
	*(*uint8)(unsafe.Add(mBase, uint32(v1107))) = uint8(v1168)
	goto L282
L301:
	;
	goto L279
L302:
	;
	v1193 = v1178
	goto L37
L303:
	;
	goto L304
L304:
	;
	if v143 == int32(0) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1193 = v1178
	goto L37
L306:
	;
	goto L307
L307:
	;
	base.MemoryCopy(m, l0+v47, v63, v143)
	v1193 = v1178
	goto L37
L308:
	;
	goto L12
L309:
	;
	v1237 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v1219))) = uint8(v1237)
	goto L311
L310:
	;
	goto L311
L311:
	;
	return v1219
}
func F_copy_dest_receive(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+200))
	F_MemoryContextReset(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(_a_F_copy_dest_receive_0)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[0]))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+200))
		*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[0])) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
		v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
		if v20 < v19 {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
			m.T0[v23].(func(*base.Module, int32, int32))(m, l0, v19)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
				m.T0[v27].(func(*base.Module, int32, int32))(m, v7, l0)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[0])) = v14
					v32 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
					v34 = v32 + int64(1)
					*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v34
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[1]))
					if v39 == int32(0) {
					} else {
						v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_copy_dest_receive[2])))
						if v43&int32(1) == int32(0) {
						} else {
							v48 = int32(_a_F_copy_dest_receive_1)
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3]))
							v51 = int32(1)
							*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3])) = v50 + v51
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
							*(*int32)(unsafe.Add(mBase, uint32(v39))) = v54 + v51
							v58 = int32(0)
							v60 = int32(_a_F_copy_dest_receive_2)
							v61 = base.AtomicRmwOr32(m, v58, v60, v58)
							*(*int64)(unsafe.Add(mBase, uint32(v39+int32(16))+232)) = v34
							v69 = base.AtomicRmwOr32(m, v58, v60, v58)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
							*(*int32)(unsafe.Add(mBase, uint32(v39))) = v70 + v51
							v76 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3]))
							*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3])) = v76 - v51
						}
					}
					return int32(1)
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
			m.T0[v27].(func(*base.Module, int32, int32))(m, v7, l0)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[0])) = v14
				v32 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v34 = v32 + int64(1)
				*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v34
				v39 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[1]))
				if v39 == int32(0) {
				} else {
					v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_copy_dest_receive[2])))
					if v43&int32(1) == int32(0) {
					} else {
						v48 = int32(_a_F_copy_dest_receive_1)
						v50 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3]))
						v51 = int32(1)
						*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3])) = v50 + v51
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
						*(*int32)(unsafe.Add(mBase, uint32(v39))) = v54 + v51
						v58 = int32(0)
						v60 = int32(_a_F_copy_dest_receive_2)
						v61 = base.AtomicRmwOr32(m, v58, v60, v58)
						*(*int64)(unsafe.Add(mBase, uint32(v39+int32(16))+232)) = v34
						v69 = base.AtomicRmwOr32(m, v58, v60, v58)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
						*(*int32)(unsafe.Add(mBase, uint32(v39))) = v70 + v51
						v76 = *(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_copy_dest_receive[3])) = v76 - v51
					}
				}
				return int32(1)
			}
		}
	}
}
func F_copy_intArrayType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = F_ArrayGetNItemsSafe(m, v6, l0+int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if int32(0) < v9 {
			v18 = v9<<(uint(int32(2))%32) + int32(24)
			v19 = F_palloc0(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v9
				*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(23)
				*(*int64)(unsafe.Add(mBase, uint32(v19)+4)) = int64(1)
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v18 << (uint(int32(2)) % 32)
				v35 = v19
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v44 = v35
				v45 = (v37<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v46 == int32(0) {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v56 = (v49<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				} else {
					v56 = v46
				}
				v58 = v9 << (uint(int32(2)) % 32)
				if v58 != 0 {
					base.MemoryCopy(m, v44+v45, l0+v56, v58)
				} else {
				}
				return v44
			}
		} else {
			v32 = F_construct_empty_array(m, int32(23))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
				if v34 != 0 {
					v44 = v32
					v45 = v34
				} else {
					v35 = v32
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
					v44 = v35
					v45 = (v37<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				}
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v46 == int32(0) {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v56 = (v49<<(uint(int32(3))%32) + int32(23)) & int32(-8)
				} else {
					v56 = v46
				}
				v58 = v9 << (uint(int32(2)) % 32)
				if v58 != 0 {
					base.MemoryCopy(m, v44+v45, l0+v56, v58)
				} else {
				}
				return v44
			}
		}
	}
}
func F_copy_pathtarget(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v40 int32
	_ = v40
	v6 = F_palloc0(m, int32(40))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(280)
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v12
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v16
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v18
		v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v20
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v23 = F_list_copy(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v23
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v26 == int32(0) {
				return v6
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v29 != 0 {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					v34 = v30 << (uint(int32(2)) % 32)
				} else {
					v34 = int32(0)
				}
				v35 = F_palloc(m, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v35
					if v34 == int32(0) {
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						base.MemoryCopy(m, v35, v40, v34)
					}
					return v6
				}
			}
		}
	}
}
func F_copysignl(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64) {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = l1
	v10 = int64(48)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = l2&int64(281474976710655) | base.I64_extend_i32_u(base.I32_wrap_i64(int64(base.Ui64(l2&int64(9223090561878065152))>>(uint(v10)%64)))|base.I32_wrap_i64(int64(base.Ui64(l3)>>(uint(v10)%64)))&int32(_a_F_copysignl_0))<<(uint(v10)%64)
	return
}
func F_cost_resultscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 float64
	_ = v16
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v65 float64
	_ = v65
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v85 float64
	_ = v85
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v93 int64
	_ = v93
	var v95 float64
	_ = v95
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	v7 = float64(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	if l3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(l2)+128))
	v92 = *(*float64)(unsafe.Add(mBase, _c_F_cost_resultscan[0]))
	v93 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	v95 = base.F64_add(v89, float64(0))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v95
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v99 != 0 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v16 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v19 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v19
	if v18 == int32(0) {
		v65 = v7
		v69 = float64(0)
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v74
	v76 = *(*float64)(unsafe.Add(mBase, uint32(l2)+216))
	v77 = *(*float64)(unsafe.Add(mBase, uint32(l2)+208))
	v85 = v76
	v89 = v77
	goto L1
L5:
	;
	v70 = *(*float64)(unsafe.Add(mBase, uint32(l2)+216))
	v72 = *(*float64)(unsafe.Add(mBase, uint32(l2)+208))
	v85 = base.F64_add(v65, v70)
	v89 = base.F64_add(v69, v72)
	goto L1
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v28 <= int32(0) {
		v65 = v7
		v69 = float64(0)
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v35 = int32(0)
	goto L8
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v35<<(uint(int32(2))%32))))
	v50 = F_cost_qual_eval_walker(m, v47, v14+int32(8))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v65 = v56
	v69 = v57
	goto L5
L10:
	;
	return
L11:
	;
	v53 = v35 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v53 < v54 {
		v35 = v53
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v100 = int64(-1)
	goto L15
L14:
	;
	v100 = int64(-262145)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = base.B2i32(v93|v100 != int64(-1))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v95, base.F64_add(base.F64_mul(v90, base.F64_add(v85, v92)), float64(0)))
	m.G0 = v14 + int32(32)
	return
}
func F_cost_seqscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 float64
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v84 float64
	_ = v84
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v109 float64
	_ = v109
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v120 float64
	_ = v120
	var v122 float64
	_ = v122
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v126 int32
	_ = v126
	var v130 float64
	_ = v130
	var v132 int32
	_ = v132
	var v138 float64
	_ = v138
	var v142 float64
	_ = v142
	var v145 float64
	_ = v145
	var v147 float64
	_ = v147
	var v148 float64
	_ = v148
	var v157 float64
	_ = v157
	var v161 float64
	_ = v161
	var v164 float64
	_ = v164
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v171 float64
	_ = v171
	v8 = float64(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = l3 + int32(8)
	goto L3
L2:
	;
	v24 = l2 + int32(16)
	goto L3
L3:
	;
	v25 = *(*float64)(unsafe.Add(mBase, uint32(v24)))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	F_get_tablespace_page_costs(m, v27, int32(0), v18)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+124))
	v32 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
	if l3 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v116 = *(*float64)(unsafe.Add(mBase, uint32(v115)+24))
	v117 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v120 = *(*float64)(unsafe.Add(mBase, _c_F_cost_seqscan[0]))
	v122 = *(*float64)(unsafe.Add(mBase, uint32(l2)+128))
	v124 = base.F64_add(base.F64_mul(v116, v117), base.F64_mul(base.F64_add(v114, v120), v122))
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v115)+16))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v126 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v34 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v34
	if v33 == int32(0) {
		v84 = v8
		v92 = float64(0)
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v97 = *(*float64)(unsafe.Add(mBase, uint32(l2)+208))
	v98 = *(*float64)(unsafe.Add(mBase, uint32(l2)+216))
	v109 = v97
	v114 = v98
	goto L6
L10:
	;
	v93 = *(*float64)(unsafe.Add(mBase, uint32(l2)+208))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(l2)+216))
	v109 = base.F64_add(v92, v93)
	v114 = base.F64_add(v84, v95)
	goto L6
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v43 <= int32(0) {
		v84 = v8
		v92 = float64(0)
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v50 = int32(0)
	goto L13
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v50<<(uint(int32(2))%32))))
	v69 = F_cost_qual_eval_walker(m, v66, v18+int32(8))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	v75 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
	v76 = *(*float64)(unsafe.Add(mBase, uint32(v18)+16))
	v84 = v75
	v92 = v76
	goto L10
L15:
	;
	v72 = v50 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v72 < v73 {
		v50 = v72
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v164 = v124
	v167 = int64(-262146)
	goto L19
L18:
	;
	v130 = base.F64_convert_i32_u(v126)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_seqscan[1])))
	if v132 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v168 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
	v171 = base.F64_add(base.F64_add(v109, float64(0)), v125)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = base.B2i32(v167|v168 != int64(-1))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(base.F64_mul(v32, base.F64_convert_i32_u(v31)), base.F64_add(v171, v164))
	m.G0 = v18 + int32(32)
	return
L20:
	;
	v138 = base.F64_add(base.F64_mul(v130, float64(-0.3)), float64(1))
	if base.F64_gt(v138, float64(0)) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v145 = v130
	goto L22
L22:
	;
	v147 = float64(1e+100)
	v148 = base.F64_div(v117, v145)
	if base.F64_gt(v148, v147)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v148)&int64(9223372036854775807))) != 0 {
		v161 = v147
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v142 = v138
	goto L25
L24:
	;
	v142 = math.Float64frombits(uint64(0x8000000000000000))
	goto L25
L25:
	;
	v145 = base.F64_add(v142, v130)
	goto L22
L26:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v161
	v164 = base.F64_div(v124, v145)
	v167 = int64(-2)
	goto L19
L27:
	;
	v157 = float64(1)
	if base.F64_le(v148, v157) != 0 {
		v161 = v157
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v161 = base.F64_nearest(v148)
	goto L26
}
func F_count_nulls(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v234 int32
	_ = v234
	v4 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 == v4 {
		v27 = v4
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
		if v19 == int32(0) {
			v27 = v4
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v22 != int32(15) {
				v27 = v4
			} else {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+13)))
				v27 = v25
			}
		}
	}
	if v27&int32(1) == int32(0) {
		v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v32 <= int32(0) {
			v207 = v32
			v209 = v4
		} else {
			if base.Ui32(int32(4)) <= base.Ui32(v32) {
				v42 = v4
				v45 = v4
				v48 = v4
				for {
					v53 = int32(4)
					v55 = l0 + v42<<(uint(v53)%32)
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+32)))
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+48)))
					v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+64)))
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+80)))
					v63 = v45 + v56 + v58 + v60 + v62
					v65 = v42 + v53
					v67 = v48 + v53
					if v67 != v32&int32(_a_F_count_nulls_0) {
						v42 = v65
						v45 = v63
						v48 = v67
						continue
					} else {
						break
					}
					break
				}
				if v32&int32(3) == int32(0) {
					v207 = v32
					v209 = v63
				} else {
					v76 = v65
					v79 = v63
					v92 = v76
					v95 = v79
					v97 = v4
					for {
						v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v92<<(uint(int32(4))%32))+32)))
						v107 = v95 + v106
						v108 = int32(1)
						v111 = v97 + v108
						if v111 != v32&int32(3) {
							v92 = v92 + v108
							v95 = v107
							v97 = v111
							continue
						} else {
							break
						}
						break
					}
					v207 = v32
					v209 = v107
				}
			} else {
				v76 = v4
				v79 = v4
				v92 = v76
				v95 = v79
				v97 = v4
				for {
					v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v92<<(uint(int32(4))%32))+32)))
					v107 = v95 + v106
					v108 = int32(1)
					v111 = v97 + v108
					if v111 != v32&int32(3) {
						v92 = v92 + v108
						v95 = v107
						v97 = v111
						continue
					} else {
						break
					}
					break
				}
				v207 = v32
				v209 = v107
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v207
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v209
		v234 = int32(1)
		return v234
	} else {
		v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v114 != 0 {
			v234 = int32(0)
			return v234
		} else {
			v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v116 = F_pg_detoast_datum(m, v115)
			mBase = m.M
			v119 = m.ExcPending
			if v119 != 0 {
				return int32(0)
			} else {
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
				v122 = v116 + int32(16)
				v123 = F_ArrayGetNItemsSafe(m, v120, v122)
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
					if v125 == int32(0) {
						v207 = v123
						v209 = v4
					} else {
						v128 = int32(1)
						if v123 <= int32(0) {
							v207 = v123
							v209 = v4
						} else {
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
							v134 = v122 + v131<<(uint(int32(3))%32)
							if v123 != int32(1) {
								v141 = v134
								v144 = v128
								v147 = v4
								v149 = v4
								for {
									v155 = int32(1)
									v158 = v144 << (uint(v155) % 32)
									v160 = base.B2i32(v158 == int32(256))
									if v158 == int32(256) {
										v161 = v155
									} else {
										v161 = v158
									}
									v163 = v161 << (uint(int32(1)) % 32)
									v165 = base.B2i32(v163 == int32(256))
									if v163 == int32(256) {
										v166 = v155
									} else {
										v166 = v163
									}
									v167 = v160 + v141
									v168 = v165 + v167
									v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
									v171 = int32(0)
									v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
									v178 = base.B2i32(v161&v169 == v171) + (v147 + base.B2i32(v144&v173 == v171))
									v180 = v149 + int32(2)
									if v180 != v123&int32(2147483646) {
										v141 = v168
										v144 = v166
										v147 = v178
										v149 = v180
										continue
									} else {
										break
									}
									break
								}
								if v123&int32(1) == int32(0) {
									v207 = v123
									v209 = v178
								} else {
									v184 = v168
									v187 = v166
									v190 = v178
									v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
									v207 = v123
									v209 = v190 + base.B2i32(v187&v198 == int32(0))
								}
							} else {
								v184 = v134
								v187 = v128
								v190 = v4
								v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
								v207 = v123
								v209 = v190 + base.B2i32(v187&v198 == int32(0))
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v207
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v209
					v234 = int32(1)
					return v234
				}
			}
		}
	}
}
