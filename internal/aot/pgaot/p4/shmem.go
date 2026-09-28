package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_InitShmemIndexEntry(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int64
	_ = v179
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemIndexEntry[0]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v25 = F_hash_search(m, v19, v21, int32(3), v16+int32(47))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+47)))
	if v27 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v440 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemIndexEntry[0]))
	v443 = F_hash_search(m, v440, v21, int32(2), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L84
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L80
	}
L5:
	;
	if v25 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L77
	}
L8:
	;
	v32 = int32(128)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	if base.Ui32(v34) <= base.Ui32(v32) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v37 = v32
	goto L11
L10:
	;
	v37 = v34
	goto L11
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemIndexEntry[1]))
	v43 = base.AtomicRmwXchg32(m, v40, int32(4), int32(1))
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_s_lock(m, v40+int32(4), int32(_a_F_InitShmemIndexEntry_0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemIndexEntry[1]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v57 = (v37 + v51 - int32(1)) & (int32(0) - v37)
	v58 = v38 + v57
	if base.Ui32(v57) <= base.Ui32(v58) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L14
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemIndexEntry[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v58
	v70 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v50)+4)), uint32(v70))
	if v68 == v70 {
		goto L3
	} else {
		goto L21
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_InitShmemIndexEntry[3]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if base.Ui32(v58) <= base.Ui32(v62) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v64 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v50)+4)), uint32(v64))
	goto L3
L20:
	;
	goto L19
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+56)) = v58 - v51
	*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = v76
	v80 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+60)) = uint8(v80)
	v82 = v57 + v68
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v84 {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	default:
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v25
	m.G0 = v16 + int32(48)
	return
L23:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v120 = m.G0
	v122 = v120 - int32(96)
	m.G0 = v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)+28))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119)+24))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+60)) = v82
	v129 = base.I32_div_s(v125, int32(16))
	*(*uint16)(unsafe.Add(mBase, uint32(v126)+64)) = uint16(v129)
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = v131
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v119)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = v133
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v119)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+16)) = v135
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v119)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+24)) = v137
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v119)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+32)) = v139
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v119)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+40)) = v141
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v119)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+48)) = v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v119)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+56)) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v126)+52))
	if v147 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v91 = m.G0
	v93 = v91 - int32(16)
	m.G0 = v93
	*(*int32)(unsafe.Add(mBase, uint32(v90)+72)) = v82
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v90)+24))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+80))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+60)) = int32(1208)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v82 + v99
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v90)+64)) = v93 + int32(8)
	v112 = F_hash_create(m, v96, v97, v90+int32(32), v98|int32(_a_F_InitShmemIndexEntry_1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	if v86 == int32(0) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v82
	goto L22
L27:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v90)+84))
	if v114 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v112
	goto L30
L29:
	;
	goto L30
L30:
	;
	m.G0 = v93 + int32(16)
	goto L22
L31:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+16)) = v150
	v153 = v122 + int32(32)
	v158 = F_pg_snprintf(m, v153, int32(64), int32(_a_F_InitShmemIndexEntry_2), v122+int32(16))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v126)+56))
	if v164 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v160 = F_LWLockNewTrancheId(m, v153)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+52)) = v160
	goto L33
L36:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v167
	v170 = v122 + int32(32)
	v173 = F_pg_snprintf(m, v170, int32(64), int32(_a_F_InitShmemIndexEntry_3), v122)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v179 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v82)+48)) = v179
	*(*int64)(unsafe.Add(mBase, uint32(v82)+40)) = v179
	*(*int64)(unsafe.Add(mBase, uint32(v82))) = v179
	*(*int64)(unsafe.Add(mBase, uint32(v82)+56)) = v179
	*(*int64)(unsafe.Add(mBase, uint32(v82)+32)) = v179
	*(*int64)(unsafe.Add(mBase, uint32(v82)+24)) = v179
	*(*int64)(unsafe.Add(mBase, uint32(v82)+16)) = v179
	*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = v179
	*(*int64)(unsafe.Add(mBase, uint32(v82)+48)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v82)+40)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v125
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v201 = F_strcmp(m, int32(_a_F_InitShmemIndexEntry_4), v199)
	mBase = m.M
	if v201 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v175 = F_LWLockNewTrancheId(m, v170)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+56)) = v175
	goto L38
L41:
	;
	v236 = int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v82 - v236
	*(*int32)(unsafe.Add(mBase, uint32(v82)+56)) = v235
	v240 = int32(2)
	v242 = int32(7)
	v244 = int32(-8)
	v245 = (v125<<(uint(v240)%32) + v242) & v244
	v247 = v245 - v236
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v82 + v247
	v250 = v245 + v247
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v82 + v250
	v257 = v250 + (v125+v242)&v244
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v82 + v257
	v261 = v125 << (uint(int32(3)) % 32)
	v262 = v257 + v261
	*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = v82 + v262
	v265 = v245 + v262
	v266 = v82 + v265
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = v266
	v271 = v265 + (v125+v129)<<(uint(v242)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+32)) = v82 + v271
	*(*int32)(unsafe.Add(mBase, uint32(v82)+28)) = v266 + v125<<(uint(v242)%32)
	v284 = v271 + (v129<<(uint(v240)%32)+v242)&v244
	if int32(0) < v124 {
		goto L63
	} else {
		goto L64
	}
L42:
	;
	v235 = int32(0)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v206 = F_strcmp(m, int32(_a_F_InitShmemIndexEntry_5), v199)
	mBase = m.M
	if v206 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v235 = int32(1)
	goto L41
L46:
	;
	goto L47
L47:
	;
	v211 = F_strcmp(m, int32(_a_F_InitShmemIndexEntry_6), v199)
	mBase = m.M
	if v211 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v235 = int32(2)
	goto L41
L49:
	;
	goto L50
L50:
	;
	v216 = F_strcmp(m, int32(_a_F_InitShmemIndexEntry_7), v199)
	mBase = m.M
	if v216 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v235 = int32(3)
	goto L41
L52:
	;
	goto L53
L53:
	;
	v221 = F_strcmp(m, int32(_a_F_InitShmemIndexEntry_8), v199)
	mBase = m.M
	if v221 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v235 = int32(4)
	goto L41
L55:
	;
	goto L56
L56:
	;
	v226 = F_strcmp(m, int32(_a_F_InitShmemIndexEntry_9), v199)
	mBase = m.M
	if v226 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v235 = int32(5)
	goto L41
L58:
	;
	goto L59
L59:
	;
	v233 = F_strcmp(m, int32(_a_F_InitShmemIndexEntry_10), v199)
	mBase = m.M
	if v233 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v234 = int32(7)
	goto L62
L61:
	;
	v234 = int32(6)
	goto L62
L62:
	;
	v235 = v234
	goto L41
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v284 + v82
	v291 = v284 + v124*v261
	goto L65
L64:
	;
	v291 = v284
	goto L65
L65:
	;
	if v125 <= int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	m.G0 = v122 + int32(96)
	goto L22
L67:
	;
	v301 = int32(0)
	v307 = v82 + (v291+int32(31))&int32(-32)
	goto L68
L68:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v126)+52))
	F_LWLockInitialize(m, v313+v301<<(uint(int32(7))%32), v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	if base.Ui32(v125) <= base.Ui32(int32(15)) {
		goto L66
	} else {
		goto L72
	}
L70:
	;
	v321 = v301 << (uint(int32(2)) % 32)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v321+v322))) = v307
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v327 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v325+v321))) = v327
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v329+v301))) = uint8(v327)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v333+v321))) = v327
	v340 = v301 + int32(1)
	if v340 != v125 {
		v301 = v340
		v307 = v307 - int32(-8192)
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v346 = int32(0)
	goto L73
L73:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v126)+56))
	F_LWLockInitialize(m, v358+v346<<(uint(int32(7))%32), v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	goto L66
L75:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v82)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v365+v346<<(uint(int32(2))%32)))) = int32(0)
	v372 = v346 + int32(1)
	if v372 != v129 {
		v346 = v372
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v21
	F_errmsg_internal(m, int32(_a_F_InitShmemIndexEntry_11), v16)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_InitShmemIndexEntry_12), int32(546), int32(_a_F_InitShmemIndexEntry_13))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_errcode(m, int32(_a_F_InitShmemIndexEntry_14))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v21
	F_errmsg(m, int32(_a_F_InitShmemIndexEntry_15), v16+int32(16))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_InitShmemIndexEntry_12), int32(553), int32(_a_F_InitShmemIndexEntry_13))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errcode(m, int32(_a_F_InitShmemIndexEntry_14))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v21
	F_errmsg(m, int32(_a_F_InitShmemIndexEntry_16), v16+int32(32))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_InitShmemIndexEntry_12), int32(571), int32(_a_F_InitShmemIndexEntry_13))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ShmemRequestInternal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v218 = int32(_a_F_ShmemRequestInternal_0)
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemRequestInternal[0]))
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemRequestInternal[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_ShmemRequestInternal[0])) = v222
	v225 = F_palloc(m, int32(12))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L14
	} else {
		goto L58
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L14
	} else {
		goto L55
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L14
	} else {
		goto L52
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L14
	} else {
		goto L49
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L14
	} else {
		goto L46
	}
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ShmemRequestInternal[2])))
	if v15 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L14
	} else {
		goto L43
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v46 = int32(0)
	if base.B2i32(base.B2i32(v45 == v46)|base.B2i32(v45&(v45-int32(1)) == v46) == v46)&base.B2i32(int32(2)<<(uint(base.I32_clz(v45)^int32(31))%32) != v45) != 0 {
		goto L3
	} else {
		goto L20
	}
L10:
	;
	if base.B2i32(v13 == int32(-1))|base.B2i32(int32(0) < v13) != 0 {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v13 == int32(-1) {
		goto L5
	} else {
		goto L18
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = base.I64_rotl(v27, int64(32))
	F_errmsg_internal(m, int32(_a_F_ShmemRequestInternal_1), v8+int32(-32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_ShmemRequestInternal_2), int32(359), int32(_a_F_ShmemRequestInternal_3))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	if v13 <= int32(0) {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L9
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemRequestInternal[3]))
	if v64 != int32(1) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemRequestInternal[4]))
	if v68 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v71 <= int32(0) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v74 = int32(0)
	if v74 < v71 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v78 = v71
	goto L26
L25:
	;
	v78 = v74
	goto L26
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v82 = v74
	goto L27
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v79+v82<<(uint(int32(2))%32))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if base.B2i32(v95 == int32(0))|base.B2i32(v95 != v98) != 0 {
		v116 = v95
		v117 = v98
		goto L30
	} else {
		goto L31
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L14
	} else {
		goto L40
	}
L29:
	;
	if v116-v117 != 0 {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	v101 = v92
	v102 = v12
	goto L32
L32:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v106 == int32(0) {
		v116 = v106
		v117 = v105
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v116 = v106
	v117 = v105
	goto L30
L34:
	;
	v109 = int32(1)
	if v106 == v105 {
		v101 = v101 + v109
		v102 = v102 + v109
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v120 = v82 + int32(1)
	if v78 != v120 {
		v82 = v120
		goto L27
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L28
L39:
	;
	goto L1
L40:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v126
	F_errmsg(m, int32(_a_F_ShmemRequestInternal_4), v10)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_ShmemRequestInternal_2), int32(384), int32(_a_F_ShmemRequestInternal_3))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_errmsg_internal(m, int32(_a_F_ShmemRequestInternal_5), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_ShmemRequestInternal_2), int32(353), int32(_a_F_ShmemRequestInternal_3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L14
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
	F_errmsg_internal(m, int32(_a_F_ShmemRequestInternal_6), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L14
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_ShmemRequestInternal_2), int32(364), int32(_a_F_ShmemRequestInternal_3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L14
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = base.I64_rotl(v166, int64(32))
	F_errmsg_internal(m, int32(_a_F_ShmemRequestInternal_1), v8+int32(-16))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_ShmemRequestInternal_2), int32(367), int32(_a_F_ShmemRequestInternal_3))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v184
	F_errmsg_internal(m, int32(_a_F_ShmemRequestInternal_7), v8+int32(-48))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L14
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_ShmemRequestInternal_2), int32(372), int32(_a_F_ShmemRequestInternal_3))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L14
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errmsg_internal(m, int32(_a_F_ShmemRequestInternal_8), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L14
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_ShmemRequestInternal_2), int32(376), int32(_a_F_ShmemRequestInternal_3))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L14
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
	*(*int32)(unsafe.Add(mBase, uint32(v225)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = l0
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemRequestInternal[4]))
	v233 = F_lappend(m, v232, v225)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L14
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ShmemRequestInternal[0])) = v219
	*(*int32)(unsafe.Add(mBase, _c_F_ShmemRequestInternal[4])) = v233
	m.G0 = v10 - int32(-64)
	return
}
func F_ShmemRequestStructWithOpts(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemRequestStructWithOpts[0]))
	v6 = F_MemoryContextAlloc(m, v4, int32(16))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v10
		F_ShmemRequestInternal(m, v6, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			return
		}
	}
}
