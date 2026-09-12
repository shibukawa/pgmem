package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckTableForSerializableConflictIn(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
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
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[849]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v17 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v24) < base.Ui32(int32(12000)) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+118)))
	if v28 == int32(116) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v32 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[850])) = uint8(v32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v40 = F_LWLockAcquire(m, v36+int32(3840), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v47 = F_LWLockAcquire(m, v43+int32(25344), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v54 = F_LWLockAcquire(m, v50+int32(25472), int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v61 = F_LWLockAcquire(m, v57+int32(25600), int32(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v68 = F_LWLockAcquire(m, v64+int32(25728), int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v75 = F_LWLockAcquire(m, v71+int32(25856), int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v82 = F_LWLockAcquire(m, v78+int32(25984), int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v89 = F_LWLockAcquire(m, v85+int32(26112), int32(1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v96 = F_LWLockAcquire(m, v92+int32(26240), int32(1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v103 = F_LWLockAcquire(m, v99+int32(26368), int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v110 = F_LWLockAcquire(m, v106+int32(26496), int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v117 = F_LWLockAcquire(m, v113+int32(26624), int32(1))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v124 = F_LWLockAcquire(m, v120+int32(26752), int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v131 = F_LWLockAcquire(m, v127+int32(26880), int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v138 = F_LWLockAcquire(m, v134+int32(27008), int32(1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v145 = F_LWLockAcquire(m, v141+int32(27136), int32(1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v152 = F_LWLockAcquire(m, v148+int32(27264), int32(1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v159 = F_LWLockAcquire(m, v155+int32(3584), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[851]))
	F_hash_seq_init(m, v13+int32(12), v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v169 = F_hash_seq_search(m, v13+int32(12))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	if v169 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v172 = v169
	goto L30
L28:
	;
	goto L29
L29:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v290+int32(3584))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L6
	} else {
		goto L56
	}
L30:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v181 != v24 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v277 = F_hash_seq_search(m, v13+int32(12))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L6
	} else {
		goto L54
	}
L33:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v183 != v34 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v172)+20))
	if v185 == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v189 = v172 + int32(16)
	if v185 == v189 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	v193 = v185
	v194 = v192
	goto L37
L37:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v193-int32(4))))
	if v206 == v194 {
		v255 = v194
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L32
L39:
	;
	if v205 != v189 {
		v193 = v205
		v194 = v255
		goto L37
	} else {
		goto L53
	}
L40:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+108)))
	if v208&int32(8) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_FlagRWConflict(m, v206, v194)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L6
	} else {
		goto L52
	}
L42:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+108)))
	if v211&int32(8) != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v206)+36))
	if v214 == int32(0) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v218 = v206 + int32(32)
	if v214 == v218 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v194)+44))
	if v220 == int32(0) {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	if v220 == v194+int32(40) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v230 = v214
	goto L48
L48:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v230)+20))
	if v236 == v194 {
		v255 = v194
		goto L39
	} else {
		goto L50
	}
L49:
	;
	goto L41
L50:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	if v238 != v218 {
		v230 = v238
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	v255 = v253
	goto L39
L53:
	;
	goto L38
L54:
	;
	if v277 != 0 {
		v172 = v277
		goto L30
	} else {
		goto L55
	}
L55:
	;
	goto L31
L56:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v296+int32(27264))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v302+int32(27136))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v308+int32(27008))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v314+int32(26880))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v320+int32(26752))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v326+int32(26624))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v332+int32(26496))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v338+int32(26368))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v344+int32(26240))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v350+int32(26112))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v356+int32(25984))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v362+int32(25856))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v368+int32(25728))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v374+int32(25600))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L6
	} else {
		goto L70
	}
L70:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v380+int32(25472))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	v386 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v386+int32(25344))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v392+int32(3840))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	goto L1
}
func F_get_table_am_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_get_am_type_oid(m, l0, int32(116), l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_has_table_privilege_id(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v16)
		v19 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v21 = F_convert_any_priv_string(m, v12, int32(1615424))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v25 = F_pg_class_aclcheck_ext(m, v10, v19, v21, v8+int32(15))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
				if v27 == int32(1) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v34 = int32(0)
				} else {
					v34 = base.B2i32(v25 == int32(0))
				}
				m.G0 = v8 + int32(16)
				return v34
			}
		}
	}
}
func F_has_table_privilege_id_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v21 = F_convert_any_priv_string(m, v14, int32(1615424))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v25 = F_pg_class_aclcheck_ext(m, v11, v12, v21, v9+int32(15))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				if v27 == int32(1) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v34 = int32(0)
				} else {
					v34 = base.B2i32(v25 == int32(0))
				}
				m.G0 = v9 + int32(16)
				return v34
			}
		}
	}
}
func F_has_table_privilege_id_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_textToQualifiedNameList(m, v6)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = F_makeRangeVarFromNameList(m, v13)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					v17 = int32(0)
					v21 = F_RangeVarGetRelidExtended(m, v15, v17, v17, v17, v17)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v24 = F_convert_any_priv_string(m, v11, int32(1615424))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v26 = F_pg_class_aclcheck(m, v21, v4, v24)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v26 == int32(0))
							}
						}
					}
				}
			}
		}
	}
}
func F_has_table_privilege_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v14 = F_textToQualifiedNameList(m, v5)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = F_makeRangeVarFromNameList(m, v14)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = int32(0)
					v22 = F_RangeVarGetRelidExtended(m, v16, v18, v18, v18, v18)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = F_convert_any_priv_string(m, v10, int32(1615424))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = F_pg_class_aclcheck(m, v22, v13, v25)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v27 == int32(0))
							}
						}
					}
				}
			}
		}
	}
}
func F_table_beginscan_catalog(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v5 = F_GetCatalogSnapshot(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_RegisterSnapshot(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			v15 = m.T0[v14].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v9, l1, l2, int32(0), int32(961))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_table_open(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_relation_open(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+119)))
		switch v13 - int32(99) {
		case 0, 6:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v25 + int32(4)
					F_errmsg(m, int32(666839), v6)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32)+119)))
						F_errdetail_relkind_not_supported(m, v33)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(477337), int32(147), int32(406537))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
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
		case 1, 2, 3, 4, 5:
			m.G0 = v6 + int32(16)
			return v8
		default:
			if v13 != int32(73) {
				m.G0 = v6 + int32(16)
				return v8
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v25 + int32(4)
						F_errmsg(m, int32(666839), v6)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
							v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32)+119)))
							F_errdetail_relkind_not_supported(m, v33)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(477337), int32(147), int32(406537))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
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
		}
	}
}
