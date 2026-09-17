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
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[0]))
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
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[1]))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v24) < base.Ui32(int32(_a_F_CheckTableForSerializableConflictIn_0)) {
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
	*(*uint8)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[2])) = uint8(v32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
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
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v47 = F_LWLockAcquire(m, v43+int32(_a_F_CheckTableForSerializableConflictIn_1), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v54 = F_LWLockAcquire(m, v50+int32(_a_F_CheckTableForSerializableConflictIn_2), int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v61 = F_LWLockAcquire(m, v57+int32(_a_F_CheckTableForSerializableConflictIn_3), int32(1))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v68 = F_LWLockAcquire(m, v64+int32(_a_F_CheckTableForSerializableConflictIn_4), int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v75 = F_LWLockAcquire(m, v71+int32(_a_F_CheckTableForSerializableConflictIn_5), int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v82 = F_LWLockAcquire(m, v78+int32(_a_F_CheckTableForSerializableConflictIn_6), int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v89 = F_LWLockAcquire(m, v85+int32(_a_F_CheckTableForSerializableConflictIn_7), int32(1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v96 = F_LWLockAcquire(m, v92+int32(_a_F_CheckTableForSerializableConflictIn_8), int32(1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v103 = F_LWLockAcquire(m, v99+int32(_a_F_CheckTableForSerializableConflictIn_9), int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v110 = F_LWLockAcquire(m, v106+int32(_a_F_CheckTableForSerializableConflictIn_10), int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v117 = F_LWLockAcquire(m, v113+int32(_a_F_CheckTableForSerializableConflictIn_11), int32(1))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v124 = F_LWLockAcquire(m, v120+int32(_a_F_CheckTableForSerializableConflictIn_12), int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v131 = F_LWLockAcquire(m, v127+int32(_a_F_CheckTableForSerializableConflictIn_13), int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v138 = F_LWLockAcquire(m, v134+int32(_a_F_CheckTableForSerializableConflictIn_14), int32(1))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v145 = F_LWLockAcquire(m, v141+int32(_a_F_CheckTableForSerializableConflictIn_15), int32(1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	v152 = F_LWLockAcquire(m, v148+int32(_a_F_CheckTableForSerializableConflictIn_16), int32(1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
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
	v162 = v13 + int32(12)
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[4]))
	F_hash_seq_init(m, v162, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v167 = F_hash_seq_search(m, v162)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	if v167 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v170 = v167
	goto L30
L28:
	;
	goto L29
L29:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v289+int32(3584))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L6
	} else {
		goto L55
	}
L30:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if v179 != v24 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v276 = F_hash_seq_search(m, v13+int32(12))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L6
	} else {
		goto L53
	}
L33:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	if v181 != v34 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v170)+20))
	if v183 == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v187 = v170 + int32(16)
	if v183 == v187 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[1]))
	v191 = v183
	v192 = v190
	goto L37
L37:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v191-int32(4))))
	if v204 == v192 {
		v254 = v192
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L32
L39:
	;
	if v203 != v187 {
		v191 = v203
		v192 = v254
		goto L37
	} else {
		goto L52
	}
L40:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+108)))
	if v206&int32(8) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_FlagRWConflict(m, v204, v192)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L6
	} else {
		goto L51
	}
L42:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+108)))
	if v209&int32(8) != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
	if v212 == int32(0) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v216 = v204 + int32(32)
	if v212 == v216 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v192)+44))
	if base.B2i32(v218 == int32(0))|base.B2i32(v218 == v192+int32(40)) != 0 {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v228 = v212
	goto L47
L47:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v228)+20))
	if v235 == v192 {
		v254 = v192
		goto L39
	} else {
		goto L49
	}
L48:
	;
	goto L41
L49:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v237 != v216 {
		v228 = v237
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[1]))
	v254 = v252
	goto L39
L52:
	;
	goto L38
L53:
	;
	if v276 != 0 {
		v170 = v276
		goto L30
	} else {
		goto L54
	}
L54:
	;
	goto L31
L55:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v295+int32(_a_F_CheckTableForSerializableConflictIn_16))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v301+int32(_a_F_CheckTableForSerializableConflictIn_15))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v307+int32(_a_F_CheckTableForSerializableConflictIn_14))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v313+int32(_a_F_CheckTableForSerializableConflictIn_13))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v319+int32(_a_F_CheckTableForSerializableConflictIn_12))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v325+int32(_a_F_CheckTableForSerializableConflictIn_11))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v331+int32(_a_F_CheckTableForSerializableConflictIn_10))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v337+int32(_a_F_CheckTableForSerializableConflictIn_9))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v343+int32(_a_F_CheckTableForSerializableConflictIn_8))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v349+int32(_a_F_CheckTableForSerializableConflictIn_7))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v355+int32(_a_F_CheckTableForSerializableConflictIn_6))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v361+int32(_a_F_CheckTableForSerializableConflictIn_5))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v367+int32(_a_F_CheckTableForSerializableConflictIn_4))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v373+int32(_a_F_CheckTableForSerializableConflictIn_3))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v379+int32(_a_F_CheckTableForSerializableConflictIn_2))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L6
	} else {
		goto L70
	}
L70:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v385+int32(_a_F_CheckTableForSerializableConflictIn_1))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_CheckTableForSerializableConflictIn[3]))
	F_LWLockRelease(m, v391+int32(3840))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L6
	} else {
		goto L72
	}
L72:
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
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_has_table_privilege_id[0]))
		v21 = F_convert_any_priv_string(m, v12, int32(_a_F_has_table_privilege_id_0))
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
		v21 = F_convert_any_priv_string(m, v14, int32(_a_F_has_table_privilege_id_id_0))
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
						v24 = F_convert_any_priv_string(m, v11, int32(_a_F_has_table_privilege_id_name_0))
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
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_has_table_privilege_name[0]))
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
						v25 = F_convert_any_priv_string(m, v10, int32(_a_F_has_table_privilege_name_0))
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
					F_errmsg(m, int32(_a_F_table_open_0), v6)
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
							F_errfinish(m, int32(_a_F_table_open_1), int32(147), int32(_a_F_table_open_2))
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
						F_errmsg(m, int32(_a_F_table_open_0), v6)
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
								F_errfinish(m, int32(_a_F_table_open_1), int32(147), int32(_a_F_table_open_2))
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
