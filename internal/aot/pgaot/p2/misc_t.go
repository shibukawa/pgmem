package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TransferPredicateLocksToNewTarget(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v145 int64
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int64
	_ = v318
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v24 = F_get_hash_value(m, v23, l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v30 = F_get_hash_value(m, v29, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[1]))
	v34 = int32(15)
	v35 = v24 & v34
	v36 = int32(7)
	v40 = v30 & v34
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[2]))
	v47 = F_LWLockAcquire(m, v45, int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v63 = int32(_a_F_TransferPredicateLocksToNewTarget_0)
	v64 = v33 + v35<<(uint(v36)%32) + v63
	v66 = v33 + v40<<(uint(v36)%32) + v63
	if base.Ui32(v35) < base.Ui32(v40) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[3]))
	v57 = F_hash_search_with_hash_value(m, v50, int32(_a_F_TransferPredicateLocksToNewTarget_1), v53, int32(2), v20+int32(13))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[2]))
	F_LWLockRelease(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v85 = int32(0)
	v87 = F_hash_search_with_hash_value(m, v84, l0, v24, v85, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L20
	}
L11:
	;
	v70 = F_LWLockAcquire(m, v64, l2^int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v76 = F_LWLockAcquire(m, v66, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v73 = F_LWLockAcquire(m, v66, int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	if base.Ui32(v35) <= base.Ui32(v40) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v81 = F_LWLockAcquire(m, v64, l2^int32(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	if base.Ui32(v35) < base.Ui32(v40) {
		goto L80
	} else {
		goto L81
	}
L20:
	;
	if v87 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v365 = int32(1)
	goto L19
L22:
	;
	goto L23
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v97 = F_hash_search_with_hash_value(m, v93, l1, v30, int32(3), v20+int32(12))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v97 == int32(0) {
		v365 = int32(0)
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+12)))
	if v101 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v105 = v97 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+20)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v105
	goto L28
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v97
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[1]))
	v115 = F_LWLockAcquire(m, v111+int32(3584), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v87)+20))
	v118 = int32(0)
	v121 = v87 + int32(16)
	if base.B2i32(v117 == v118)|base.B2i32(v121 == v117) == v118 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v127 = v97 + int32(16)
	v129 = v117
	goto L33
L31:
	;
	goto L32
L32:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[1]))
	F_LWLockRelease(m, v344+int32(3584))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L71
	}
L33:
	;
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v129)+16))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v148 = v129 - int32(4)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v149
	if l2 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+4)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v156)+4)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v159
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[4]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v172 = F_hash_search_with_hash_value(m, v162, v129-int32(8), v165<<(uint(int32(4))%32)^v24, int32(2), v20+int32(12))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v175 = v149
	goto L37
L37:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[4]))
	v178 = int32(4)
	v186 = F_hash_search_with_hash_value(m, v177, v20+v178, v175<<(uint(v178)%32)^v30, int32(3), v20+int32(12))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v175 = v174
	goto L37
L39:
	;
	if v186 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v190 = int32(0)
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[1]))
	F_LWLockRelease(m, v192+int32(3584))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+12)))
	if v288 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L43:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[1]))
	v202 = F_LWLockAcquire(m, v198+int32(3584), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	v205 = int32(0)
	if base.B2i32(v204 == v205)|base.B2i32(v204 == v127) == v205 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v211 = v204
	goto L48
L46:
	;
	goto L47
L47:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[1]))
	F_LWLockRelease(m, v273+int32(3584))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L52
	}
L48:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+4)) = v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v230))) = v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v234)+4)) = v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	*(*int32)(unsafe.Add(mBase, uint32(v235))) = v237
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[4]))
	v243 = int32(4)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v211-v243)))
	v252 = F_hash_search_with_hash_value(m, v240, v211-int32(8), v245<<(uint(v243)%32)^v30, int32(2), v20+int32(14))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	goto L47
L50:
	;
	if v127 != v228 {
		v211 = v228
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	if v278 != v127 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v281 = v278
	goto L55
L54:
	;
	v281 = int32(0)
	goto L55
L55:
	;
	if v281 != 0 {
		v365 = v190
		goto L19
	} else {
		goto L56
	}
L56:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v286 = F_hash_search_with_hash_value(m, v283, v97, v30, int32(2), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v365 = v190
	goto L19
L58:
	;
	if v146 != v121 {
		v129 = v146
		goto L33
	} else {
		goto L70
	}
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v186)+24)) = v145
	goto L58
L60:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	if v291 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v186)+24))
	if base.Ui64(v145) <= base.Ui64(v318) {
		goto L58
	} else {
		goto L69
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+20)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v127
	goto L65
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v127
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+8)) = v297
	v300 = v186 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v297)+4)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = v300
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v305 = v303 + int32(48)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v303)+52))
	if v306 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+52)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v303)+48)) = v305
	goto L68
L67:
	;
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = v305
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+16)) = v312
	v315 = v186 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v312)+4)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = v315
	goto L59
L69:
	;
	goto L59
L70:
	;
	goto L34
L71:
	;
	v349 = int32(1)
	if l2 == int32(0) {
		v365 = v349
		goto L19
	} else {
		goto L72
	}
L72:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v87)+20))
	if v352 != v121 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v355 = v352
	goto L75
L74:
	;
	v355 = int32(0)
	goto L75
L75:
	;
	if v355 != 0 {
		v365 = v349
		goto L19
	} else {
		goto L76
	}
L76:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v360 = F_hash_search_with_hash_value(m, v357, v87, v24, int32(2), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v365 = v349
	goto L19
L78:
	;
	F_LWLockRelease(m, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L85
	}
L79:
	;
	F_LWLockRelease(m, v381)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L84
	}
L80:
	;
	v381 = v66
	v382 = v64
	goto L79
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(v35) <= base.Ui32(v40) {
		v386 = v66
		goto L78
	} else {
		goto L83
	}
L83:
	;
	v381 = v64
	v382 = v66
	goto L79
L84:
	;
	v386 = v382
	goto L78
L85:
	;
	if l2 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[2]))
	v392 = F_LWLockAcquire(m, v390, int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	m.G0 = v20 + int32(16)
	return v365
L89:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[0]))
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[3]))
	v402 = F_hash_search_with_hash_value(m, v395, int32(_a_F_TransferPredicateLocksToNewTarget_1), v398, int32(1), v20+int32(15))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v405 = *(*int32)(unsafe.Add(mBase, _c_F_TransferPredicateLocksToNewTarget[2]))
	F_LWLockRelease(m, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	goto L88
}
func F_t_isalpha_cstr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = F_pg_mblen_cstr(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = F_pg_mb2wchar_with_len(m, l0, v5+int32(8), v9)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			v16 = F_pg_database_locale(m)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = F_pg_iswalpha(m, v15, v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(16)
					return v18
				}
			}
		}
	}
}
func F_tblspc_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+48)))
	v13 = v11 & int32(240)
	if v13 != 0 {
		if v13 == int32(16) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v24
			F_appendStringInfo(m, l0, int32(_a_F_tblspc_desc_0), v7+int32(16))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				m.G0 = v7 + int32(32)
				return
			}
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v10 + int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v16
		F_appendStringInfo(m, l0, int32(_a_F_tblspc_desc_1), v7)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	}
}
func F_test_lockmode_for_conflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	v6 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v6)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2*int32(12))+uint32(_c_F_test_lockmode_for_conflict[0])))
	if base.Ui32(l1) < base.Ui32(int32(3)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v241
L2:
	;
	if v144 != 0 {
		v241 = int32(2)
		goto L1
	} else {
		goto L42
	}
L3:
	;
	v144 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_test_lockmode_for_conflict[1]))
	if v24 == l1 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v144 = int32(1)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_test_lockmode_for_conflict[2]))
	if v28 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v144 = v134
	goto L2
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_test_lockmode_for_conflict[3]))
	if v32 == int32(0) {
		v134 = v6
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_test_lockmode_for_conflict[4]))
	v104 = int32(0)
	v107 = v28 - int32(1)
	goto L32
L13:
	;
	v37 = v32
	goto L14
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	if v43 == int32(4) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v134 = int32(0)
	goto L9
L16:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v37)+80))
	if v97 != 0 {
		v37 = v97
		goto L14
	} else {
		goto L31
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v46 == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v49 = int32(1)
	if l1 == v46 {
		v134 = v49
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v37)+52))
	v53 = v51 - int32(1)
	if v53 < int32(0) {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v37)+48))
	v59 = int32(0)
	v62 = v53
	goto L21
L21:
	;
	v67 = int32(2)
	v68 = base.I32_div_s(v62-v59, v67)
	v69 = v68 + v59
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v56+v69<<(uint(v67)%32))))
	if v73 == l1 {
		v134 = v49
		goto L9
	} else {
		goto L23
	}
L22:
	;
	goto L16
L23:
	;
	v82 = base.B2i32(v73-l1 < int32(0)) | base.B2i32(base.Ui32(v73) < base.Ui32(int32(3)))
	if v82 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v83 = v69 + int32(1)
	goto L26
L25:
	;
	v83 = v59
	goto L26
L26:
	;
	if v82 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v86 = v62
	goto L29
L28:
	;
	v86 = v69 - int32(1)
	goto L29
L29:
	;
	if v83 <= v86 {
		v59 = v83
		v62 = v86
		goto L21
	} else {
		goto L30
	}
L30:
	;
	goto L22
L31:
	;
	goto L15
L32:
	;
	v112 = int32(2)
	v113 = base.I32_div_s(v107-v104, v112)
	v114 = v113 + v104
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v102+v114<<(uint(v112)%32))))
	v119 = base.B2i32(v118 == l1)
	if v118 == l1 {
		v134 = v119
		goto L9
	} else {
		goto L34
	}
L33:
	;
	v134 = v119
	goto L9
L34:
	;
	v122 = base.B2i32(base.Ui32(v118) < base.Ui32(l1))
	if base.Ui32(v118) < base.Ui32(l1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v123 = v114 + int32(1)
	goto L37
L36:
	;
	v123 = v104
	goto L37
L37:
	;
	if base.Ui32(v118) < base.Ui32(l1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v126 = v107
	goto L40
L39:
	;
	v126 = v114 - int32(1)
	goto L40
L40:
	;
	if v123 <= v126 {
		v104 = v123
		v107 = v126
		goto L32
	} else {
		goto L41
	}
L41:
	;
	goto L33
L42:
	;
	v145 = F_TransactionIdIsInProgress(m, l1)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return int32(0)
L44:
	;
	if v145 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v150 = int32(2)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v150)%32))+uint32(_c_F_test_lockmode_for_conflict[5])))
	v153 = int32(12)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v152*v153)+uint32(_c_F_test_lockmode_for_conflict[6])))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v11<<(uint(v150)%32))+uint32(_c_F_test_lockmode_for_conflict[5])))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160*v153)+uint32(_c_F_test_lockmode_for_conflict[6])))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v157<<(uint(v150)%32))+uint32(_c_F_test_lockmode_for_conflict[7])))
	goto L48
L46:
	;
	goto L47
L47:
	;
	v180 = int32(0)
	v181 = F_TransactionIdDidAbort(m, l1)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L43
	} else {
		goto L50
	}
L48:
	;
	if int32(base.Ui32(v170)>>(uint(v165)%32))&int32(1) == int32(0) {
		v241 = int32(0)
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v176 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v176)
	return int32(0)
L50:
	;
	if v181 != 0 {
		v241 = v180
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v183 = F_TransactionIdDidCommit(m, l1)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L43
	} else {
		goto L52
	}
L52:
	;
	if base.B2i32(v183 == int32(0))|base.B2i32(base.Ui32(l0) < base.Ui32(int32(4))) != 0 {
		v241 = v180
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v190 = int32(2)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(v190)%32))+uint32(_c_F_test_lockmode_for_conflict[5])))
	v193 = int32(12)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192*v193)+uint32(_c_F_test_lockmode_for_conflict[6])))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v11<<(uint(v190)%32))+uint32(_c_F_test_lockmode_for_conflict[5])))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v200*v193)+uint32(_c_F_test_lockmode_for_conflict[6])))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v197<<(uint(v190)%32))+uint32(_c_F_test_lockmode_for_conflict[7])))
	goto L54
L54:
	;
	if int32(base.Ui32(v210)>>(uint(v205)%32))&int32(1) == int32(0) {
		v241 = v180
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v216 = int32(4)
	v219 = l3 + v216
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v222 = v220 + int32(12)
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+2)))
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219))))
	v225 = int32(16)
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+2)))
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222))))
	if v223|v224<<(uint(v225)%32) == v228|v229<<(uint(v225)%32) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	if v239 != 0 {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	goto L56
L58:
	;
	v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+4)))
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+4)))
	if v235 == v236 {
		v239 = int32(1)
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v239 = int32(0)
	goto L57
L61:
	;
	goto L60
L62:
	;
	v240 = v216
	goto L64
L63:
	;
	v240 = int32(3)
	goto L64
L64:
	;
	v241 = v240
	goto L1
}
func F_texticlike(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v11 = F_Generic_Text_IC_like(m, v3, v8, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(base.B2i32(v11 == int32(1)))
			}
		}
	}
}
func F_texticlike_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_like_regex_support(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_texticnlike(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v11 = F_Generic_Text_IC_like(m, v3, v8, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(base.B2i32(v11 != int32(1)))
			}
		}
	}
}
func F_texticregexeq(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14378(m, l0, int32(27))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_textin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_strlen(m, v5)
	mBase = m.M
	v8 = v6 + int32(4)
	v9 = F_palloc(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v8 << (uint(int32(2)) % 32)
		if v6 != 0 {
			base.MemoryCopy(m, v9+int32(4), v5, v6)
		} else {
		}
		return base.I64_extend_i32_u(v9)
	}
}
func F_textltname(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_DirectFunctionCall2Coll(m, int32(1756), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(base.Ui64(v6)>>(uint(int64(31))%64)) & int64(1)
	}
}
func F_textne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = F_pg_newlocale_from_collation(m, v10)
	mBase = m.M
	v14 = m.ExcPending
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L82
	}
L4:
	;
	return int64(0)
L5:
	;
	v15 = int32(1)
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v17 == v15 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return base.I64_extend_i32_u(v214)
L7:
	;
	F_pfree(m, v208)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L81
	}
L8:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v21 = F_toast_raw_datum_size(m, v16)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v119 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v16))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L45
	}
L11:
	;
	v23 = F_toast_raw_datum_size(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v21 != v23 {
		v214 = v15
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v27 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v16))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v30 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v20))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v32 = int32(1)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v34&v32 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v37 = v32
	goto L18
L17:
	;
	v37 = int32(4)
	goto L18
L18:
	;
	v38 = v27 + v37
	v39 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v41&v39 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v44 = v39
	goto L21
L20:
	;
	v44 = int32(4)
	goto L21
L21:
	;
	v45 = v30 + v44
	v46 = int32(4)
	v47 = v21 - v46
	if base.Ui32(v46) <= base.Ui32(v47) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v110 != v27 {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	v109 = int32(0)
	goto L22
L24:
	;
	v83 = v78
	v84 = v79
	v85 = v80
	goto L34
L25:
	;
	if (v38|v45)&int32(3) != 0 {
		v78 = v38
		v79 = v45
		v80 = v47
		goto L24
	} else {
		goto L28
	}
L26:
	;
	v71 = v38
	v72 = v45
	v73 = v47
	goto L27
L27:
	;
	if v73 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v55 = v38
	v56 = v45
	v57 = v47
	goto L29
L29:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v60 != v61 {
		v78 = v55
		v79 = v56
		v80 = v57
		goto L24
	} else {
		goto L31
	}
L30:
	;
	v71 = v66
	v72 = v64
	v73 = v68
	goto L27
L31:
	;
	v63 = int32(4)
	v64 = v56 + v63
	v66 = v55 + v63
	v68 = v57 - v63
	if base.Ui32(int32(3)) < base.Ui32(v68) {
		v55 = v66
		v56 = v64
		v57 = v68
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v78 = v71
	v79 = v72
	v80 = v73
	goto L24
L34:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v88 == v89 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v109 = v88 - v89
	goto L22
L36:
	;
	v91 = int32(1)
	v96 = v85 - v91
	if v96 != 0 {
		v83 = v83 + v91
		v84 = v84 + v91
		v85 = v96
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L35
L39:
	;
	goto L23
L40:
	;
	F_pfree(m, v27)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v115 = base.B2i32(v109 != int32(0))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v30 != v116 {
		v206 = v115
		v208 = v30
		goto L7
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v214 = v115
	goto L6
L45:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v122 = F_pg_detoast_datum_packed(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v125 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v155 = int32(1)
	if v125&v155 != 0 {
		goto L58
	} else {
		goto L59
	}
L48:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	if v131 == int32(18) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v142 = int32(1)
	if v125&v142 != 0 {
		v154 = int32(base.Ui32(v125)>>(uint(v142)%32)) - v142
		goto L47
	} else {
		goto L57
	}
L51:
	;
	v134 = int32(16)
	goto L53
L52:
	;
	v134 = int32(0)
	goto L53
L53:
	;
	if base.Ui32((v131-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v141 = int32(4)
	goto L56
L55:
	;
	v141 = v134
	goto L56
L56:
	;
	v154 = v141
	goto L47
L57:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v154 = int32(base.Ui32(v148)>>(uint(int32(2))%32)) - int32(4)
	goto L47
L58:
	;
	v159 = v155
	goto L60
L59:
	;
	v159 = int32(4)
	goto L60
L60:
	;
	v161 = int32(1)
	if v124&v161 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v165 = v161
	goto L63
L62:
	;
	v165 = int32(4)
	goto L63
L63:
	;
	if v124 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v196 = F_varstr_cmp(m, v119+v159, v154, v122+v165, v195, v10)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L75
	}
L65:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	if v172 == int32(18) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v183 = int32(1)
	if v124&v183 != 0 {
		v195 = int32(base.Ui32(v124)>>(uint(v183)%32)) - v183
		goto L64
	} else {
		goto L74
	}
L68:
	;
	v175 = int32(16)
	goto L70
L69:
	;
	v175 = int32(0)
	goto L70
L70:
	;
	if base.Ui32((v172-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v182 = int32(4)
	goto L73
L72:
	;
	v182 = v175
	goto L73
L73:
	;
	v195 = v182
	goto L64
L74:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v195 = int32(base.Ui32(v189)>>(uint(int32(2))%32)) - int32(4)
	goto L64
L75:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v198 != v119 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_pfree(m, v119)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v203 = base.B2i32(v196 != int32(0))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v122 == v204 {
		v214 = v203
		goto L6
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v206 = v203
	v208 = v122
	goto L7
L81:
	;
	v214 = v206
	goto L6
L82:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_textne_0), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errhint(m, int32(_a_F_textne_1), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_textne_2), int32(1337), int32(_a_F_textne_3))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_throttle(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v7 = v5 + base.I64_extend_i32_u(l1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v9) <= base.Ui64(v7) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = base.I64_div_u_s(v7, v9)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v21 = m.G0
	v22 = int32(16)
	v23 = v21 - v22
	m.G0 = v23
	F_gettimeofday(m, v23)
	mBase = m.M
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	v27 = int64(*(*int32)(unsafe.Add(mBase, uint32(v23)+8)))
	m.G0 = v23 + v22
	goto L7
L5:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v79 = base.I64_rem_u_s(v77, v78)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v79
	v84 = m.G0
	v85 = int32(16)
	v86 = v84 - v85
	m.G0 = v86
	F_gettimeofday(m, v86)
	mBase = m.M
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
	v90 = int64(*(*int32)(unsafe.Add(mBase, uint32(v86)+8)))
	m.G0 = v86 + v85
	goto L21
L6:
	;
	goto L5
L7:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v38 = v36 - (v27 + v26*int64(1000000) - int64(946684800000000)) + v11*v12
	if v38 <= int64(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_throttle[0]))
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v43
	v48 = base.AtomicRmwOr32(m, v43, int32(_a_F_throttle_0), v43)
	goto L9
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_throttle[1]))
	if v50 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_throttle[0]))
	v57 = base.I64_div_u_s(v38, int64(1000))
	v60 = F_WaitLatch(m, v54, int32(41), base.I32_wrap_i64(v57), int32(150994944))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L13
	} else {
		goto L16
	}
L13:
	;
	return
L14:
	;
	goto L12
L15:
	;
	if v60&int32(8) == int32(0) {
		goto L4
	} else {
		goto L20
	}
L16:
	;
	if v60&int32(1) == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_throttle[1]))
	if v67 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	goto L6
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v90 + v89*int64(1000000) - int64(946684800000000)
	goto L3
}
func F_tidle(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+2)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2))))
	v9 = int32(16)
	v11 = v7 | v8<<(uint(v9)%32)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v16 = v12 | v13<<(uint(v9)%32)
	if base.Ui32(v11) < base.Ui32(v16) {
		v27 = int32(-1)
	} else {
		if base.Ui32(v16) < base.Ui32(v11) {
			v27 = int32(1)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+4)))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v21) < base.Ui32(v22) {
				v27 = int32(-1)
			} else {
				v27 = base.B2i32(base.Ui32(v22) < base.Ui32(v21))
			}
		}
	}
	return base.I64_extend_i32_u(base.B2i32(v27 <= int32(0)))
}
func F_timeofday(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v5 = m.G0
	v7 = v5 - int32(576)
	m.G0 = v7
	F_gettimeofday(m, v7+int32(560))
	mBase = m.M
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v7)+560))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+552)) = v12
	v15 = v7 + int32(416)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_timeofday[0]))
	v22 = F_pg_localtime(m, v7+int32(552), v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int64(0)
	} else {
		v26 = F_pg_strftime(m, v15, int32(128), int32(_a_F_timeofday_0), v22)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			v29 = v7 + int32(288)
			v32 = F_pg_strftime(m, v29, int32(128), int32(_a_F_timeofday_1), v22)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int64(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+568))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v15
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v29
				v39 = v7 + int32(16)
				v42 = F_pg_snprintf(m, v39, int32(266), int32(_a_F_timeofday_2), v7)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					v44 = F_cstring_to_text(m, v39)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int64(0)
					} else {
						m.G0 = v7 + int32(576)
						return base.I64_extend_i32_u(v44)
					}
				}
			}
		}
	}
}
func F_timestamptz2timestamp_safe(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v88 int64
	_ = v88
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v111 int64
	_ = v111
	var v118 int64
	_ = v118
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int64
	_ = v146
	var v149 int64
	_ = v149
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int64
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int64
	_ = v186
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	if base.Ui64(l0-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		v186 = l0
		m.G0 = v10 + int32(80)
		return v186
	} else {
		v22 = int32(0)
		v24 = F_timestamp2tm(m, l0, v10+int32(28), v10+int32(36), v10+int32(32), v22, v22)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			if v24 != 0 {
				v31 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
				v32 = F_errsave_start(m, l1)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					if v32 == int32(0) {
						v186 = v31
						m.G0 = v10 + int32(80)
						return v186
					} else {
						v170 = v31
						v172 = int32(_a_F_timestamptz2timestamp_safe_0)
						F_errcode(m, int32(134217858))
						mBase = m.M
						v175 = m.ExcPending
						if v175 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
							mBase = m.M
							v179 = m.ExcPending
							if v179 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
								mBase = m.M
								v183 = m.ExcPending
								if v183 != 0 {
									return int64(0)
								} else {
									v186 = v170
									m.G0 = v10 + int32(80)
									return v186
								}
							}
						}
					}
				}
			} else {
				v37 = int64(*(*int32)(unsafe.Add(mBase, uint32(v10)+32)))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
				if v38 <= int32(-4713) {
					if v38 != int32(-4713) {
						v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
						v163 = F_errsave_start(m, l1)
						mBase = m.M
						v164 = m.ExcPending
						if v164 != 0 {
							return int64(0)
						} else {
							if v163 == int32(0) {
								v186 = v162
								m.G0 = v10 + int32(80)
								return v186
							} else {
								v170 = v162
								v172 = int32(_a_F_timestamptz2timestamp_safe_4)
								F_errcode(m, int32(134217858))
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
										return int64(0)
									} else {
										F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
										mBase = m.M
										v183 = m.ExcPending
										if v183 != 0 {
											return int64(0)
										} else {
											v186 = v170
											m.G0 = v10 + int32(80)
											return v186
										}
									}
								}
							}
						}
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
						if int32(10) < v43 {
							v54 = v43
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
							v60 = base.B2i32(int32(2) < v54)
							if int32(2) < v54 {
								v61 = int32(_a_F_timestamptz2timestamp_safe_5)
							} else {
								v61 = int32(_a_F_timestamptz2timestamp_safe_6)
							}
							v62 = v61 + v38
							v67 = base.I32_div_s(v62, int32(4))
							v70 = base.I32_div_s(v62, int32(-100))
							v73 = base.I32_div_s(v62, int32(400))
							if int32(2) < v54 {
								v77 = int32(1)
							} else {
								v77 = int32(13)
							}
							v82 = base.I32_div_s((v77+v54)*int32(_a_F_timestamptz2timestamp_safe_7), int32(256))
							v88 = base.I64_extend_i32_s(v55 + v62*int32(365) + v67 + v70 + v73 + v82 - int32(_a_F_timestamptz2timestamp_safe_8) - int32(_a_F_timestamptz2timestamp_safe_9))
							v97 = int64(32)
							v98 = int64(20)
							v100 = int64(base.Ui64(v88) >> (uint(v97) % 64))
							v103 = int64(4294967295)
							v104 = int64(500654080)
							v106 = v88 & v103
							v107 = v104 * v106
							v111 = int64(base.Ui64(v107)>>(uint(v97)%64)) + v104*v100
							v118 = v106*v98 + v111&v103
							*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v88*int64(0) + v88>>(uint(int64(63))%64)*int64(86400000000) + v98*v100 + int64(base.Ui64(v111)>>(uint(v97)%64)) + int64(base.Ui64(v118)>>(uint(v97)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v10))) = v107&v103 | v118<<(uint(v97)%64)
							v129 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
							v130 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
							if v129 != v130>>(uint(int64(63))%64) {
								v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
								v163 = F_errsave_start(m, l1)
								mBase = m.M
								v164 = m.ExcPending
								if v164 != 0 {
									return int64(0)
								} else {
									if v163 == int32(0) {
										v186 = v162
										m.G0 = v10 + int32(80)
										return v186
									} else {
										v170 = v162
										v172 = int32(_a_F_timestamptz2timestamp_safe_4)
										F_errcode(m, int32(134217858))
										mBase = m.M
										v175 = m.ExcPending
										if v175 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
											mBase = m.M
											v179 = m.ExcPending
											if v179 != 0 {
												return int64(0)
											} else {
												F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int64(0)
												} else {
													v186 = v170
													m.G0 = v10 + int32(80)
													return v186
												}
											}
										}
									}
								}
							} else {
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
								v137 = int32(60)
								v146 = base.I64_extend_i32_s(v134+(v135+v136*v137)*v137)*int64(1000000) + v37
								v149 = v130 + v146
								if base.B2i32(v146 < int64(0))^base.B2i32(v149 < v130) != 0 {
									v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
									v163 = F_errsave_start(m, l1)
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int64(0)
									} else {
										if v163 == int32(0) {
											v186 = v162
											m.G0 = v10 + int32(80)
											return v186
										} else {
											v170 = v162
											v172 = int32(_a_F_timestamptz2timestamp_safe_4)
											F_errcode(m, int32(134217858))
											mBase = m.M
											v175 = m.ExcPending
											if v175 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
												mBase = m.M
												v179 = m.ExcPending
												if v179 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int64(0)
													} else {
														v186 = v170
														m.G0 = v10 + int32(80)
														return v186
													}
												}
											}
										}
									}
								} else {
									if base.Ui64(int64(9011559254509551615)) < base.Ui64(v149-int64(9223371331200000000)) {
										v186 = v149
										m.G0 = v10 + int32(80)
										return v186
									} else {
										v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
										v163 = F_errsave_start(m, l1)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int64(0)
										} else {
											if v163 == int32(0) {
												v186 = v162
												m.G0 = v10 + int32(80)
												return v186
											} else {
												v170 = v162
												v172 = int32(_a_F_timestamptz2timestamp_safe_4)
												F_errcode(m, int32(134217858))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
													mBase = m.M
													v179 = m.ExcPending
													if v179 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return int64(0)
														} else {
															v186 = v170
															m.G0 = v10 + int32(80)
															return v186
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
							v163 = F_errsave_start(m, l1)
							mBase = m.M
							v164 = m.ExcPending
							if v164 != 0 {
								return int64(0)
							} else {
								if v163 == int32(0) {
									v186 = v162
									m.G0 = v10 + int32(80)
									return v186
								} else {
									v170 = v162
									v172 = int32(_a_F_timestamptz2timestamp_safe_4)
									F_errcode(m, int32(134217858))
									mBase = m.M
									v175 = m.ExcPending
									if v175 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
											return int64(0)
										} else {
											F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
											mBase = m.M
											v183 = m.ExcPending
											if v183 != 0 {
												return int64(0)
											} else {
												v186 = v170
												m.G0 = v10 + int32(80)
												return v186
											}
										}
									}
								}
							}
						}
					}
				} else {
					if v38 <= int32(_a_F_timestamptz2timestamp_safe_10) {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
						v54 = v48
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
						v60 = base.B2i32(int32(2) < v54)
						if int32(2) < v54 {
							v61 = int32(_a_F_timestamptz2timestamp_safe_5)
						} else {
							v61 = int32(_a_F_timestamptz2timestamp_safe_6)
						}
						v62 = v61 + v38
						v67 = base.I32_div_s(v62, int32(4))
						v70 = base.I32_div_s(v62, int32(-100))
						v73 = base.I32_div_s(v62, int32(400))
						if int32(2) < v54 {
							v77 = int32(1)
						} else {
							v77 = int32(13)
						}
						v82 = base.I32_div_s((v77+v54)*int32(_a_F_timestamptz2timestamp_safe_7), int32(256))
						v88 = base.I64_extend_i32_s(v55 + v62*int32(365) + v67 + v70 + v73 + v82 - int32(_a_F_timestamptz2timestamp_safe_8) - int32(_a_F_timestamptz2timestamp_safe_9))
						v97 = int64(32)
						v98 = int64(20)
						v100 = int64(base.Ui64(v88) >> (uint(v97) % 64))
						v103 = int64(4294967295)
						v104 = int64(500654080)
						v106 = v88 & v103
						v107 = v104 * v106
						v111 = int64(base.Ui64(v107)>>(uint(v97)%64)) + v104*v100
						v118 = v106*v98 + v111&v103
						*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v88*int64(0) + v88>>(uint(int64(63))%64)*int64(86400000000) + v98*v100 + int64(base.Ui64(v111)>>(uint(v97)%64)) + int64(base.Ui64(v118)>>(uint(v97)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v10))) = v107&v103 | v118<<(uint(v97)%64)
						v129 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
						if v129 != v130>>(uint(int64(63))%64) {
							v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
							v163 = F_errsave_start(m, l1)
							mBase = m.M
							v164 = m.ExcPending
							if v164 != 0 {
								return int64(0)
							} else {
								if v163 == int32(0) {
									v186 = v162
									m.G0 = v10 + int32(80)
									return v186
								} else {
									v170 = v162
									v172 = int32(_a_F_timestamptz2timestamp_safe_4)
									F_errcode(m, int32(134217858))
									mBase = m.M
									v175 = m.ExcPending
									if v175 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
											return int64(0)
										} else {
											F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
											mBase = m.M
											v183 = m.ExcPending
											if v183 != 0 {
												return int64(0)
											} else {
												v186 = v170
												m.G0 = v10 + int32(80)
												return v186
											}
										}
									}
								}
							}
						} else {
							v134 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
							v135 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
							v137 = int32(60)
							v146 = base.I64_extend_i32_s(v134+(v135+v136*v137)*v137)*int64(1000000) + v37
							v149 = v130 + v146
							if base.B2i32(v146 < int64(0))^base.B2i32(v149 < v130) != 0 {
								v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
								v163 = F_errsave_start(m, l1)
								mBase = m.M
								v164 = m.ExcPending
								if v164 != 0 {
									return int64(0)
								} else {
									if v163 == int32(0) {
										v186 = v162
										m.G0 = v10 + int32(80)
										return v186
									} else {
										v170 = v162
										v172 = int32(_a_F_timestamptz2timestamp_safe_4)
										F_errcode(m, int32(134217858))
										mBase = m.M
										v175 = m.ExcPending
										if v175 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
											mBase = m.M
											v179 = m.ExcPending
											if v179 != 0 {
												return int64(0)
											} else {
												F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int64(0)
												} else {
													v186 = v170
													m.G0 = v10 + int32(80)
													return v186
												}
											}
										}
									}
								}
							} else {
								if base.Ui64(int64(9011559254509551615)) < base.Ui64(v149-int64(9223371331200000000)) {
									v186 = v149
									m.G0 = v10 + int32(80)
									return v186
								} else {
									v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
									v163 = F_errsave_start(m, l1)
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int64(0)
									} else {
										if v163 == int32(0) {
											v186 = v162
											m.G0 = v10 + int32(80)
											return v186
										} else {
											v170 = v162
											v172 = int32(_a_F_timestamptz2timestamp_safe_4)
											F_errcode(m, int32(134217858))
											mBase = m.M
											v175 = m.ExcPending
											if v175 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
												mBase = m.M
												v179 = m.ExcPending
												if v179 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int64(0)
													} else {
														v186 = v170
														m.G0 = v10 + int32(80)
														return v186
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						if v38 != int32(_a_F_timestamptz2timestamp_safe_11) {
							v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
							v163 = F_errsave_start(m, l1)
							mBase = m.M
							v164 = m.ExcPending
							if v164 != 0 {
								return int64(0)
							} else {
								if v163 == int32(0) {
									v186 = v162
									m.G0 = v10 + int32(80)
									return v186
								} else {
									v170 = v162
									v172 = int32(_a_F_timestamptz2timestamp_safe_4)
									F_errcode(m, int32(134217858))
									mBase = m.M
									v175 = m.ExcPending
									if v175 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
											return int64(0)
										} else {
											F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
											mBase = m.M
											v183 = m.ExcPending
											if v183 != 0 {
												return int64(0)
											} else {
												v186 = v170
												m.G0 = v10 + int32(80)
												return v186
											}
										}
									}
								}
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
							if int32(5) < v51 {
								v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
								v163 = F_errsave_start(m, l1)
								mBase = m.M
								v164 = m.ExcPending
								if v164 != 0 {
									return int64(0)
								} else {
									if v163 == int32(0) {
										v186 = v162
										m.G0 = v10 + int32(80)
										return v186
									} else {
										v170 = v162
										v172 = int32(_a_F_timestamptz2timestamp_safe_4)
										F_errcode(m, int32(134217858))
										mBase = m.M
										v175 = m.ExcPending
										if v175 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
											mBase = m.M
											v179 = m.ExcPending
											if v179 != 0 {
												return int64(0)
											} else {
												F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int64(0)
												} else {
													v186 = v170
													m.G0 = v10 + int32(80)
													return v186
												}
											}
										}
									}
								}
							} else {
								v54 = v51
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
								v60 = base.B2i32(int32(2) < v54)
								if int32(2) < v54 {
									v61 = int32(_a_F_timestamptz2timestamp_safe_5)
								} else {
									v61 = int32(_a_F_timestamptz2timestamp_safe_6)
								}
								v62 = v61 + v38
								v67 = base.I32_div_s(v62, int32(4))
								v70 = base.I32_div_s(v62, int32(-100))
								v73 = base.I32_div_s(v62, int32(400))
								if int32(2) < v54 {
									v77 = int32(1)
								} else {
									v77 = int32(13)
								}
								v82 = base.I32_div_s((v77+v54)*int32(_a_F_timestamptz2timestamp_safe_7), int32(256))
								v88 = base.I64_extend_i32_s(v55 + v62*int32(365) + v67 + v70 + v73 + v82 - int32(_a_F_timestamptz2timestamp_safe_8) - int32(_a_F_timestamptz2timestamp_safe_9))
								v97 = int64(32)
								v98 = int64(20)
								v100 = int64(base.Ui64(v88) >> (uint(v97) % 64))
								v103 = int64(4294967295)
								v104 = int64(500654080)
								v106 = v88 & v103
								v107 = v104 * v106
								v111 = int64(base.Ui64(v107)>>(uint(v97)%64)) + v104*v100
								v118 = v106*v98 + v111&v103
								*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v88*int64(0) + v88>>(uint(int64(63))%64)*int64(86400000000) + v98*v100 + int64(base.Ui64(v111)>>(uint(v97)%64)) + int64(base.Ui64(v118)>>(uint(v97)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v10))) = v107&v103 | v118<<(uint(v97)%64)
								v129 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
								v130 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
								if v129 != v130>>(uint(int64(63))%64) {
									v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
									v163 = F_errsave_start(m, l1)
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int64(0)
									} else {
										if v163 == int32(0) {
											v186 = v162
											m.G0 = v10 + int32(80)
											return v186
										} else {
											v170 = v162
											v172 = int32(_a_F_timestamptz2timestamp_safe_4)
											F_errcode(m, int32(134217858))
											mBase = m.M
											v175 = m.ExcPending
											if v175 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
												mBase = m.M
												v179 = m.ExcPending
												if v179 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int64(0)
													} else {
														v186 = v170
														m.G0 = v10 + int32(80)
														return v186
													}
												}
											}
										}
									}
								} else {
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
									v136 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
									v137 = int32(60)
									v146 = base.I64_extend_i32_s(v134+(v135+v136*v137)*v137)*int64(1000000) + v37
									v149 = v130 + v146
									if base.B2i32(v146 < int64(0))^base.B2i32(v149 < v130) != 0 {
										v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
										v163 = F_errsave_start(m, l1)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int64(0)
										} else {
											if v163 == int32(0) {
												v186 = v162
												m.G0 = v10 + int32(80)
												return v186
											} else {
												v170 = v162
												v172 = int32(_a_F_timestamptz2timestamp_safe_4)
												F_errcode(m, int32(134217858))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
													mBase = m.M
													v179 = m.ExcPending
													if v179 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return int64(0)
														} else {
															v186 = v170
															m.G0 = v10 + int32(80)
															return v186
														}
													}
												}
											}
										}
									} else {
										if base.Ui64(int64(9011559254509551615)) < base.Ui64(v149-int64(9223371331200000000)) {
											v186 = v149
											m.G0 = v10 + int32(80)
											return v186
										} else {
											v162 = l0>>(uint(int64(63))%64) ^ int64(9223372036854775807)
											v163 = F_errsave_start(m, l1)
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int64(0)
											} else {
												if v163 == int32(0) {
													v186 = v162
													m.G0 = v10 + int32(80)
													return v186
												} else {
													v170 = v162
													v172 = int32(_a_F_timestamptz2timestamp_safe_4)
													F_errcode(m, int32(134217858))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamptz2timestamp_safe_1), int32(0))
														mBase = m.M
														v179 = m.ExcPending
														if v179 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, l1, int32(_a_F_timestamptz2timestamp_safe_2), v172, int32(_a_F_timestamptz2timestamp_safe_3))
															mBase = m.M
															v183 = m.ExcPending
															if v183 != 0 {
																return int64(0)
															} else {
																v186 = v170
																m.G0 = v10 + int32(80)
																return v186
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
func F_timetztypmodout(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14386(m, l0, int32(_a_F_timetztypmodout_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_tokenize_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = base.I64_rotl(v10, int64(32))
		F_errcontext_msg(m, int32(_a_F_tokenize_error_callback_0), v5)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			m.G0 = v5 + int32(16)
			return
		}
	}
}
func F_tolower(m *base.Module, l0 int32) int32 {
	var v8 int32
	_ = v8
	if base.Ui32(l0-int32(65)) < base.Ui32(int32(26)) {
		v8 = l0 | int32(32)
	} else {
		v8 = l0
	}
	return v8
}
func F_top12(m *base.Module, l0 float64) int32 {
	return base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0)) >> (uint(int64(52)) % 64)))
}
func F_touched_lseg_inside_poly(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v27 float64
	_ = v27
	var v30 int64
	_ = v30
	var v33 float64
	_ = v33
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 float64
	_ = v37
	var v40 int64
	_ = v40
	var v47 float64
	_ = v47
	var v54 int64
	_ = v54
	var v59 float64
	_ = v59
	var v79 float64
	_ = v79
	var v83 int32
	_ = v83
	var v90 float64
	_ = v90
	var v94 int64
	_ = v94
	var v95 float64
	_ = v95
	var v98 int64
	_ = v98
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int64
	_ = v129
	var v134 float64
	_ = v134
	var v139 float64
	_ = v139
	var v141 int32
	_ = v141
	var v143 int64
	_ = v143
	var v144 float64
	_ = v144
	var v152 float64
	_ = v152
	var v153 float64
	_ = v153
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v163 float64
	_ = v163
	var v184 int32
	_ = v184
	var v185 float64
	_ = v185
	var v186 float64
	_ = v186
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v197 float64
	_ = v197
	var v199 int32
	_ = v199
	var v201 int64
	_ = v201
	var v202 float64
	_ = v202
	var v205 int64
	_ = v205
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v19
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v21
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v25
	v27 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v30 = base.I64_reinterpret_f64(v27) & int64(9223372036854775807)
	if base.Ui64(v30) <= base.Ui64(int64(9218868437227405312)) {
		goto L17
	} else {
		goto L18
	}
L1:
	;
	m.G0 = v17 + int32(32)
	return v258
L2:
	;
	v247 = F_lseg_inside_poly(m, l1, v239, l3, l4)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L36
	} else {
		goto L60
	}
L3:
	;
	v233 = F_lseg_contain_point(m, v17, l2)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L36
	} else {
		goto L56
	}
L4:
	;
	v222 = F_lseg_contain_point(m, v17, l2)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L36
	} else {
		goto L54
	}
L5:
	;
	v202 = *(*float64)(unsafe.Add(mBase, uint32(v199)+8))
	v205 = base.I64_reinterpret_f64(v202) & int64(9223372036854775807)
	if base.Ui64(v201) <= base.Ui64(int64(9218868437227405312)) {
		goto L49
	} else {
		goto L50
	}
L6:
	;
	if base.B2i32(v184 == int32(0))|base.F64_ne(v185, v27) != 0 {
		v230 = v188
		goto L3
	} else {
		goto L48
	}
L7:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v152)&int64(9223372036854775807)) {
		v230 = v155
		goto L3
	} else {
		goto L42
	}
L8:
	;
	v144 = *(*float64)(unsafe.Add(mBase, uint32(v141)))
	if base.Ui64(v143) < base.Ui64(int64(9218868437227405313)) {
		v152 = v144
		v153 = v139
		v155 = v141
		v157 = v143
		goto L7
	} else {
		goto L41
	}
L9:
	;
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
	if base.Ui64(v129&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		v230 = v126
		goto L3
	} else {
		goto L40
	}
L10:
	;
	v120 = l2 + int32(16)
	if base.Ui64(v30) <= base.Ui64(int64(9218868437227405312)) {
		v139 = v90
		v141 = v120
		v143 = v94
		goto L8
	} else {
		goto L39
	}
L11:
	;
	v113 = l2 + int32(16)
	v114 = F_lseg_contain_point(m, v17, v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L36
	} else {
		goto L37
	}
L12:
	;
	v95 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v98 = base.I64_reinterpret_f64(v95) & int64(9223372036854775807)
	if base.Ui64(v94) <= base.Ui64(int64(9218868437227405312)) {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	v139 = v37
	v141 = l2 + int32(16)
	v143 = v40
	goto L8
L14:
	;
	if base.F64_ne(v33, v27) != 0 {
		goto L13
	} else {
		goto L28
	}
L15:
	;
	v126 = l2 + int32(16)
	goto L9
L16:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v27, v33)), float64(1e-06)) == int32(0))&base.F64_ne(v33, v27) != 0 {
		goto L13
	} else {
		goto L26
	}
L17:
	;
	v33 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v35 = int64(9223372036854775807)
	v36 = base.I64_reinterpret_f64(v33) & v35
	v37 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = base.I64_reinterpret_f64(v37) & v35
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v40) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui64(v54&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v83 = base.B2i32(base.Ui64(v36) < base.Ui64(int64(9218868437227405313)))
	goto L14
L21:
	;
	goto L22
L22:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v36) {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	v47 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui64(base.I64_reinterpret_f64(v47)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	v83 = int32(1)
	goto L14
L25:
	;
	v59 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v90 = v59
	v94 = base.I64_reinterpret_f64(v59) & int64(9223372036854775807)
	goto L12
L26:
	;
	if base.F64_eq(v37, v47)|base.F64_le(base.F64_abs(base.F64_sub(v37, v47)), float64(1e-06)) != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v79 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	v152 = v79
	v153 = v37
	v155 = l2 + int32(16)
	v157 = v40
	goto L7
L28:
	;
	if v83 != 0 {
		v90 = v37
		v94 = v40
		goto L12
	} else {
		goto L29
	}
L29:
	;
	goto L13
L30:
	;
	if base.F64_ne(v95, v90) != 0 {
		goto L10
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if base.Ui64(v98) < base.Ui64(int64(9218868437227405313)) {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	if base.Ui64(v98) < base.Ui64(int64(9218868437227405313)) {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	goto L10
L35:
	;
	goto L11
L36:
	;
	return int32(0)
L37:
	;
	if v114 != 0 {
		v239 = v113
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v258 = int32(1)
	goto L1
L39:
	;
	v126 = v120
	goto L9
L40:
	;
	v134 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	v197 = v134
	v199 = v126
	v201 = base.I64_reinterpret_f64(v134) & int64(9223372036854775807)
	goto L5
L41:
	;
	v184 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v144)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	v185 = v144
	v186 = v139
	v188 = v141
	v190 = v143
	goto L6
L42:
	;
	v163 = *(*float64)(unsafe.Add(mBase, uint32(l2)+24))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v163)&int64(9223372036854775807)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v184 = int32(1)
	v185 = v152
	v186 = v153
	v188 = v155
	v190 = v157
	goto L6
L44:
	;
	goto L45
L45:
	;
	if base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v27, v152)), float64(1e-06)) == int32(0))&base.F64_ne(v152, v27) != 0 {
		v230 = v155
		goto L3
	} else {
		goto L46
	}
L46:
	;
	if base.F64_eq(v153, v163)|base.F64_le(base.F64_abs(base.F64_sub(v153, v163)), float64(1e-06)) != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v230 = v155
	goto L3
L48:
	;
	v197 = v186
	v199 = v188
	v201 = v190
	goto L5
L49:
	;
	if base.B2i32(base.Ui64(int64(9218868437227405313)) <= base.Ui64(v205))|base.F64_ne(v202, v197) != 0 {
		v230 = v199
		goto L3
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if base.Ui64(v205) < base.Ui64(int64(9218868437227405313)) {
		v230 = v199
		goto L3
	} else {
		goto L53
	}
L52:
	;
	goto L4
L53:
	;
	goto L4
L54:
	;
	if v222 != 0 {
		v239 = l2
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v258 = int32(1)
	goto L1
L56:
	;
	if v233 != 0 {
		v239 = l2
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v235 = F_lseg_contain_point(m, v17, v230)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L36
	} else {
		goto L58
	}
L58:
	;
	if v235 != 0 {
		v239 = v230
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v258 = int32(1)
	goto L1
L60:
	;
	v258 = v247
	goto L1
}
func F_tqueueDestroyReceiver(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3 != 0 {
		F_shm_mq_detach(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v7 = m.ExcPending
			if v7 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
func F_tqueueReceiveSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v40 int32
	_ = v40
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_ExecFetchSlotMinimalTuple(m, l0, v6+int32(15))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v16 = int32(0)
		v18 = F_shm_mq_send(m, v14, v15, v10, v16, v16)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
			if v20 == int32(1) {
				F_pfree(m, v10)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					switch v18 {
					case 0, 2:
						m.G0 = v6 + int32(16)
						return base.B2i32(v18 != int32(2))
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_tqueueReceiveSlot_0), int32(0))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_tqueueReceiveSlot_1), int32(74), int32(_a_F_tqueueReceiveSlot_2))
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
			} else {
				switch v18 {
				case 0, 2:
					m.G0 = v6 + int32(16)
					return base.B2i32(v18 != int32(2))
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_tqueueReceiveSlot_0), int32(0))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_tqueueReceiveSlot_1), int32(74), int32(_a_F_tqueueReceiveSlot_2))
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
func F_treekey_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5 = int32(0)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
	if base.B2i32(v12 == v5)|base.B2i32(v15 == v5) != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v89
L2:
	;
	v89 = v79
	goto L1
L3:
	;
	v79 = v12 - v15
	goto L2
L4:
	;
	v19 = int32(8)
	v23 = v3 + v19
	v24 = v4 + v19
	v27 = v15
	v28 = v12
	goto L5
L5:
	;
	v32 = int32(2)
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24))))
	if base.Ui32(v36) < base.Ui32(v37) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	v39 = v36
	goto L9
L8:
	;
	v39 = v37
	goto L9
L9:
	;
	v40 = F_memcmp(m, v23+v32, v24+v32, v39)
	mBase = m.M
	if v40 != 0 {
		v79 = v40
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v36 != v37 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v89 = v36 - v37
	goto L1
L12:
	;
	goto L13
L13:
	;
	if v28 < int32(2) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v45 = int32(1)
	v47 = int32(9)
	v49 = int32(_a_F_treekey_cmp_0)
	if v45 < v27 {
		v23 = v23 + (v36+v47)&v49
		v24 = v24 + (v37+v47)&v49
		v27 = v27 - v45
		v28 = v28 - v45
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L6
}
func F_trigramsMatchGraph(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v80 int32
	_ = v80
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v195 int32
	_ = v195
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	base.MemoryFill(m, v14, int32(0), v13)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryFill(m, v18, int32(0), v17)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v21 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = v3
	v29 = v3
	goto L10
L8:
	;
	goto L9
L9:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v95 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v95)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v107 = v95
	v112 = v3
	goto L22
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v24+v28<<(uint(int32(2))%32))))
	v41 = v40 + v29
	if v40 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v80 = v28 + int32(1)
	if v80 != v21 {
		v28 = v80
		v29 = v41
		goto L10
	} else {
		goto L20
	}
L13:
	;
	v46 = v29
	goto L14
L14:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v46))))
	if v57 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v65 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v63+v28))) = uint8(v65)
	goto L12
L16:
	;
	v61 = v46 + int32(1)
	if v61 < v41 {
		v46 = v61
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	goto L12
L20:
	;
	goto L11
L21:
	;
	return v195
L22:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v97+v112<<(uint(int32(2))%32))))
	v120 = v100 + v117<<(uint(int32(3))%32)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if int32(0) < v121 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v195 = int32(0)
	goto L21
L24:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	v129 = int32(0)
	v132 = v107
	goto L27
L25:
	;
	v172 = v107
	goto L26
L26:
	;
	v180 = v112 + int32(1)
	if v180 < v172 {
		v107 = v172
		v112 = v180
		goto L22
	} else {
		goto L34
	}
L27:
	;
	v141 = v125 + v129<<(uint(int32(3))%32)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124+v142))))
	if v144 != int32(1) {
		v162 = v132
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v172 = v162
	goto L26
L29:
	;
	v165 = v129 + int32(1)
	if v165 != v121 {
		v129 = v165
		v132 = v162
		goto L27
	} else {
		goto L33
	}
L30:
	;
	v147 = int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v148 == v147 {
		v195 = v147
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v151 = v94 + v148
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v152 != 0 {
		v162 = v132
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v153 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v153)
	*(*int32)(unsafe.Add(mBase, uint32(v97+v132<<(uint(int32(2))%32)))) = v148
	v162 = v132 + v153
	goto L29
L33:
	;
	goto L28
L34:
	;
	goto L23
}
func F_truncate_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	if v15 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(_a_F_truncate_cb_wrapper_0)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v14
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(1058)
		*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v19
		v23 = int32(_a_F_truncate_cb_wrapper_1)
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_truncate_cb_wrapper[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_truncate_cb_wrapper[0])) = v12 + int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v12 + int32(16)
		v33 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+147)) = uint8(v33)
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v35
		v37 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
		v38 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+164)) = uint8(v38)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+152)) = v37
		m.T0[v15].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14, l1, l2, l3, l4)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_truncate_cb_wrapper[0])) = v44
			m.G0 = v12 + int32(32)
			return
		}
	} else {
		m.G0 = v12 + int32(32)
		return
	}
}
func F_try_hashjoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v101 int32
	_ = v101
	var v104 float64
	_ = v104
	var v105 float64
	_ = v105
	var v108 float64
	_ = v108
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v115 float64
	_ = v115
	var v117 float64
	_ = v117
	var v120 float64
	_ = v120
	var v122 float64
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v160 float64
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 float64
	_ = v171
	var v173 float64
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v187 float64
	_ = v187
	var v193 float64
	_ = v193
	var v195 float64
	_ = v195
	var v205 int64
	_ = v205
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 float64
	_ = v219
	var v220 float64
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 float64
	_ = v260
	var v267 float64
	_ = v267
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(96)
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v31 = F_calc_non_nestloop_required_outer(m, l2, l3)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L17
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v19 = v17
	goto L7
L6:
	;
	v19 = int32(0)
	goto L7
L7:
	;
	v20 = F_bms_is_member(m, v15, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	if v20 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v24 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v27 = v25
	goto L13
L12:
	;
	v27 = int32(0)
	goto L13
L13:
	;
	v28 = F_bms_is_member(m, v23, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	if v28 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L4
L16:
	;
	F_bms_free(m, v31)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L8
	} else {
		goto L93
	}
L17:
	;
	if v31 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v34 = int32(0)
	if base.B2i32(v31 == v34)|base.B2i32(v33 == v34) != 0 {
		v79 = v34
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v93 = m.G0
	v95 = v93 - int32(16)
	m.G0 = v95
	v97 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v98 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v100 = *(*float64)(unsafe.Add(mBase, _c_F_try_hashjoin_path[0]))
	if l4 != 0 {
		goto L36
	} else {
		goto L37
	}
L21:
	;
	if v79 == int32(0) {
		goto L16
	} else {
		goto L34
	}
L22:
	;
	goto L21
L23:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v44 < v45 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v47 = v44
	goto L26
L25:
	;
	v47 = v45
	goto L26
L26:
	;
	if v47 <= int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v50 = int32(1)
	goto L29
L28:
	;
	v50 = v47
	goto L29
L29:
	;
	v51 = int32(8)
	v56 = int32(0)
	goto L30
L30:
	;
	v63 = v56 << (uint(int32(2)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v33+v51+v63)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v31+v51+v63)))
	v68 = v65 & v67
	v70 = base.B2i32(v68 != int32(0))
	if v68 != 0 {
		v79 = v70
		goto L22
	} else {
		goto L32
	}
L31:
	;
	v79 = v70
	goto L22
L32:
	;
	v72 = v56 + int32(1)
	if v72 != v50 {
		v56 = v72
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	goto L20
L35:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v219 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
	v220 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
	v221 = int32(0)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v226 == v221 {
		goto L55
	} else {
		goto L56
	}
L36:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v104 = base.F64_convert_i32_s(v101)
	goto L38
L37:
	;
	v104 = float64(0)
	goto L38
L38:
	;
	v105 = base.F64_mul(v100, v104)
	v108 = *(*float64)(unsafe.Add(mBase, _c_F_try_hashjoin_path[1]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v112 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v113 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v115 = float64(0)
	v117 = base.F64_add(base.F64_mul(v105, v98), base.F64_add(base.F64_sub(v112, v113), v115))
	v120 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	v122 = base.F64_add(base.F64_mul(base.F64_add(v105, v108), v97), base.F64_add(base.F64_add(v113, v115), v120))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v125 = *(*int64)(unsafe.Add(mBase, uint32(l6)+40))
	goto L40
L40:
	;
	goto L41
L41:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+32))
	F_ExecChooseHashTableSize(m, v97, v147, int32(1), int32(0), v111, v95, v95+int32(12), v95+int32(8), v95+int32(4))
	mBase = m.M
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	if int32(2) <= v156 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v160 = *(*float64)(unsafe.Add(mBase, _c_F_try_hashjoin_path[2]))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+32))
	v163 = int32(7)
	v165 = int32(-8)
	v167 = int32(24)
	v171 = float64(0.0001220703125)
	v173 = base.F64_ceil(base.F64_mul(base.F64_mul(v98, base.F64_convert_i32_u((v162+v163)&v165+v167)), v171))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+32))
	v187 = base.F64_ceil(base.F64_mul(base.F64_mul(v97, base.F64_convert_i32_u((v176+v163)&v165+v167)), v171))
	v193 = base.F64_add(base.F64_mul(v160, v187), v122)
	v195 = base.F64_add(base.F64_mul(v160, base.F64_add(base.F64_add(v173, v173), v187)), v117)
	goto L50
L49:
	;
	v193 = v122
	v195 = v117
	goto L50
L50:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v195
	*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v193
	*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = base.F64_add(v195, v193)
	if v111 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v205 = int64(-2049)
	goto L53
L52:
	;
	v205 = int64(-264193)
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v123 + v124 + base.B2i32(v205|v125 != int64(-1))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+88)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v211
	m.G0 = v95 + int32(16)
	goto L35
L54:
	;
	if v306 == int32(0) {
		goto L16
	} else {
		goto L90
	}
L55:
	;
	v306 = int32(1)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v230 <= int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v306 = int32(1)
	goto L54
L59:
	;
	goto L60
L60:
	;
	if v31 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v235 = int32(0)
	goto L63
L62:
	;
	v235 = v221
	goto L63
L63:
	;
	if v31 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v238 = int32(25)
	goto L66
L65:
	;
	v238 = int32(24)
	goto L66
L66:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v238))))
	v248 = v221
	goto L67
L67:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251+v248<<(uint(int32(2))%32))))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+40))
	if v218 != v256 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v306 = v294
	goto L54
L69:
	;
	if v240 != 0 {
		goto L77
	} else {
		goto L78
	}
L70:
	;
	if v256 <= v218 {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v260 = *(*float64)(unsafe.Add(mBase, uint32(v255)+56))
	if base.F64_le(v220, base.F64_mul(v260, float64(1.01))) == int32(0) {
		goto L69
	} else {
		goto L74
	}
L73:
	;
	v306 = int32(1)
	goto L54
L74:
	;
	v306 = int32(1)
	goto L54
L75:
	;
	goto L68
L76:
	;
	v288 = int32(1)
	v290 = v248 + v288
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v290 < v291 {
		v248 = v290
		goto L67
	} else {
		goto L89
	}
L77:
	;
	v267 = *(*float64)(unsafe.Add(mBase, uint32(v255)+48))
	if base.F64_gt(v219, base.F64_mul(v267, float64(1.01))) == int32(0) {
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	if v274 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	v277 = int32(0)
	goto L83
L82:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v255)+64))
	v277 = v276
	goto L83
L83:
	;
	v278 = F_compare_pathkeys(m, v235, v277)
	mBase = m.M
	if v278&int32(-3) != 0 {
		goto L76
	} else {
		goto L84
	}
L84:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	if v281 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v284 = v282
	goto L87
L86:
	;
	v284 = int32(0)
	goto L87
L87:
	;
	v285 = F_bms_equal(m, v31, v284)
	mBase = m.M
	if v285 != 0 {
		v294 = int32(0)
		goto L75
	} else {
		goto L88
	}
L88:
	;
	goto L76
L89:
	;
	v294 = v288
	goto L75
L90:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v311 = F_create_hashjoin_path(m, l0, l1, l5, v12, l6, l2, l3, int32(0), v310, v31, l4)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	F_add_path(m, l1, v311)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	goto L1
L93:
	;
	goto L1
}
func F_tsm_bernoulli_handler(m *base.Module, l0 int32) int64 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14392(m, l0, int32(257), int32(289), int32(288), int32(287), int32(0), int32(286), int32(700))
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v9
	}
}
func F_tsmatchjoinsel(m *base.Module, l0 int32) int64 {
	return int64(4572414629676717179)
}
func F_tsmatchsel(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 float64
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v134 int32
	_ = v134
	var v137 float32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int64
	_ = v145
	var v150 float32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int64
	_ = v180
	var v185 float32
	_ = v185
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v215 float32
	_ = v215
	var v216 float64
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v235 float64
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 float64
	_ = v250
	var v251 int32
	_ = v251
	var v267 float64
	_ = v267
	var v270 float32
	_ = v270
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 float64
	_ = v282
	var v283 int32
	_ = v283
	var v299 float64
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 float64
	_ = v305
	var v313 float64
	_ = v313
	var v331 int64
	_ = v331
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	v22 = int64(4572414629676717179)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v32 = F_get_restriction_variable(m, v23, v24, v25, v20+int32(12), v20+int32(8), v20+int32(7))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int64(0)
	} else {
		if v32 == int32(0) {
			v331 = v22
			m.G0 = v20 + int32(80)
			return v331
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
			if v39 != int32(7) {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
				if v42 == int32(0) {
					v331 = v22
					m.G0 = v20 + int32(80)
					return v331
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
					m.T0[v45].(func(*base.Module, int32))(m, v42)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int64(0)
					} else {
						v331 = v22
						m.G0 = v20 + int32(80)
						return v331
					}
				}
			} else {
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+32)))
				if v48 == int32(1) {
					v51 = int64(0)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
					if v52 == int32(0) {
						v331 = v51
						m.G0 = v20 + int32(80)
						return v331
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
						m.T0[v55].(func(*base.Module, int32))(m, v52)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							v331 = v51
							m.G0 = v20 + int32(80)
							return v331
						}
					}
				} else {
					v58 = float64(0.005)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
					if v59 != int32(3615) {
						v299 = v58
						v301 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
						if v301 != 0 {
							v302 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
							m.T0[v302].(func(*base.Module, int32))(m, v301)
							mBase = m.M
							v304 = m.ExcPending
							if v304 != 0 {
								return int64(0)
							} else {
								v305 = float64(0)
								if base.F64_lt(v299, v305) != 0 {
									v313 = v305
								} else {
									if base.F64_gt(v299, float64(1)) == int32(0) {
										v313 = v299
									} else {
										v313 = float64(1)
									}
								}
								v331 = base.I64_reinterpret_f64(v313)
								m.G0 = v20 + int32(80)
								return v331
							}
						} else {
							v305 = float64(0)
							if base.F64_lt(v299, v305) != 0 {
								v313 = v305
							} else {
								if base.F64_gt(v299, float64(1)) == int32(0) {
									v313 = v299
								} else {
									v313 = float64(1)
								}
							}
							v331 = base.I64_reinterpret_f64(v313)
							m.G0 = v20 + int32(80)
							return v331
						}
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
						if v62 != int32(3614) {
							v299 = v58
							v301 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
							if v301 != 0 {
								v302 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
								m.T0[v302].(func(*base.Module, int32))(m, v301)
								mBase = m.M
								v304 = m.ExcPending
								if v304 != 0 {
									return int64(0)
								} else {
									v305 = float64(0)
									if base.F64_lt(v299, v305) != 0 {
										v313 = v305
									} else {
										if base.F64_gt(v299, float64(1)) == int32(0) {
											v313 = v299
										} else {
											v313 = float64(1)
										}
									}
									v331 = base.I64_reinterpret_f64(v313)
									m.G0 = v20 + int32(80)
									return v331
								}
							} else {
								v305 = float64(0)
								if base.F64_lt(v299, v305) != 0 {
									v313 = v305
								} else {
									if base.F64_gt(v299, float64(1)) == int32(0) {
										v313 = v299
									} else {
										v313 = float64(1)
									}
								}
								v331 = base.I64_reinterpret_f64(v313)
								m.G0 = v20 + int32(80)
								return v331
							}
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
							if v66 == int32(0) {
								v299 = float64(0)
								v301 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
								if v301 != 0 {
									v302 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
									m.T0[v302].(func(*base.Module, int32))(m, v301)
									mBase = m.M
									v304 = m.ExcPending
									if v304 != 0 {
										return int64(0)
									} else {
										v305 = float64(0)
										if base.F64_lt(v299, v305) != 0 {
											v313 = v305
										} else {
											if base.F64_gt(v299, float64(1)) == int32(0) {
												v313 = v299
											} else {
												v313 = float64(1)
											}
										}
										v331 = base.I64_reinterpret_f64(v313)
										m.G0 = v20 + int32(80)
										return v331
									}
								} else {
									v305 = float64(0)
									if base.F64_lt(v299, v305) != 0 {
										v313 = v305
									} else {
										if base.F64_gt(v299, float64(1)) == int32(0) {
											v313 = v299
										} else {
											v313 = float64(1)
										}
									}
									v331 = base.I64_reinterpret_f64(v313)
									m.G0 = v20 + int32(80)
									return v331
								}
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
								if v70 != 0 {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
									v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
									v79 = F_get_attstatsslot(m, v20+int32(44), v70, int32(4), int32(0), int32(3))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int64(0)
									} else {
										if v79 != 0 {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
											v82 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
											if v81 != v82+int32(2) {
												v87 = v65 + int32(8)
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
												v92 = int32(0)
												v95 = F_tsquery_opr_selec(m, v87, v87+v88*int32(12), v92, v92, float32(0))
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int64(0)
												} else {
													v235 = v95
													F_free_attstatsslot(m, v20+int32(44))
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int64(0)
													} else {
														v267 = v235
														v270 = *(*float32)(unsafe.Add(mBase, uint32(v71+v72)+8))
														v299 = base.F64_mul(v267, base.F64_sub(float64(1), base.F64_promote_f32(v270)))
														v301 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
														if v301 != 0 {
															v302 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
															m.T0[v302].(func(*base.Module, int32))(m, v301)
															mBase = m.M
															v304 = m.ExcPending
															if v304 != 0 {
																return int64(0)
															} else {
																v305 = float64(0)
																if base.F64_lt(v299, v305) != 0 {
																	v313 = v305
																} else {
																	if base.F64_gt(v299, float64(1)) == int32(0) {
																		v313 = v299
																	} else {
																		v313 = float64(1)
																	}
																}
																v331 = base.I64_reinterpret_f64(v313)
																m.G0 = v20 + int32(80)
																return v331
															}
														} else {
															v305 = float64(0)
															if base.F64_lt(v299, v305) != 0 {
																v313 = v305
															} else {
																if base.F64_gt(v299, float64(1)) == int32(0) {
																	v313 = v299
																} else {
																	v313 = float64(1)
																}
															}
															v331 = base.I64_reinterpret_f64(v313)
															m.G0 = v20 + int32(80)
															return v331
														}
													}
												}
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
												v100 = F_palloc_mul(m, int32(8), v82)
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int64(0)
												} else {
													if v82 <= int32(0) {
													} else {
														v104 = int32(0)
														if v82 != int32(1) {
															v111 = v104
															v125 = int32(0)
															for {
																v128 = int32(3)
																v129 = v111 << (uint(v128) % 32)
																v130 = v100 + v129
																v132 = *(*int64)(unsafe.Add(mBase, uint32(v98+v129)))
																*(*uint32)(unsafe.Add(mBase, uint32(v130))) = uint32(v132)
																v134 = int32(2)
																v137 = *(*float32)(unsafe.Add(mBase, uint32(v97+v111<<(uint(v134)%32))))
																*(*float32)(unsafe.Add(mBase, uint32(v130)+4)) = v137
																v140 = v111 | int32(1)
																v142 = v140 << (uint(v128) % 32)
																v143 = v100 + v142
																v145 = *(*int64)(unsafe.Add(mBase, uint32(v98+v142)))
																*(*uint32)(unsafe.Add(mBase, uint32(v143))) = uint32(v145)
																v150 = *(*float32)(unsafe.Add(mBase, uint32(v97+v140<<(uint(v134)%32))))
																*(*float32)(unsafe.Add(mBase, uint32(v143)+4)) = v150
																v153 = v111 + v134
																v155 = v125 + v134
																if v155 != v82&int32(2147483646) {
																	v111 = v153
																	v125 = v155
																	continue
																} else {
																	break
																}
																break
															}
															if v82&int32(1) == int32(0) {
															} else {
																v159 = v153
																v177 = v159 << (uint(int32(3)) % 32)
																v178 = v100 + v177
																v180 = *(*int64)(unsafe.Add(mBase, uint32(v98+v177)))
																*(*uint32)(unsafe.Add(mBase, uint32(v178))) = uint32(v180)
																v185 = *(*float32)(unsafe.Add(mBase, uint32(v97+v159<<(uint(int32(2))%32))))
																*(*float32)(unsafe.Add(mBase, uint32(v178)+4)) = v185
															}
														} else {
															v159 = v104
															v177 = v159 << (uint(int32(3)) % 32)
															v178 = v100 + v177
															v180 = *(*int64)(unsafe.Add(mBase, uint32(v98+v177)))
															*(*uint32)(unsafe.Add(mBase, uint32(v178))) = uint32(v180)
															v185 = *(*float32)(unsafe.Add(mBase, uint32(v97+v159<<(uint(int32(2))%32))))
															*(*float32)(unsafe.Add(mBase, uint32(v178)+4)) = v185
														}
													}
													v204 = int32(8)
													v205 = v65 + v204
													v206 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
													v215 = *(*float32)(unsafe.Add(mBase, uint32(v97+v81<<(uint(int32(2))%32)-v204)))
													v216 = F_tsquery_opr_selec(m, v205, v205+v206*int32(12), v100, v82, v215)
													mBase = m.M
													v217 = m.ExcPending
													if v217 != 0 {
														return int64(0)
													} else {
														F_pfree(m, v100)
														mBase = m.M
														v219 = m.ExcPending
														if v219 != 0 {
															return int64(0)
														} else {
															v235 = v216
															F_free_attstatsslot(m, v20+int32(44))
															mBase = m.M
															v240 = m.ExcPending
															if v240 != 0 {
																return int64(0)
															} else {
																v267 = v235
																v270 = *(*float32)(unsafe.Add(mBase, uint32(v71+v72)+8))
																v299 = base.F64_mul(v267, base.F64_sub(float64(1), base.F64_promote_f32(v270)))
																v301 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
																if v301 != 0 {
																	v302 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
																	m.T0[v302].(func(*base.Module, int32))(m, v301)
																	mBase = m.M
																	v304 = m.ExcPending
																	if v304 != 0 {
																		return int64(0)
																	} else {
																		v305 = float64(0)
																		if base.F64_lt(v299, v305) != 0 {
																			v313 = v305
																		} else {
																			if base.F64_gt(v299, float64(1)) == int32(0) {
																				v313 = v299
																			} else {
																				v313 = float64(1)
																			}
																		}
																		v331 = base.I64_reinterpret_f64(v313)
																		m.G0 = v20 + int32(80)
																		return v331
																	}
																} else {
																	v305 = float64(0)
																	if base.F64_lt(v299, v305) != 0 {
																		v313 = v305
																	} else {
																		if base.F64_gt(v299, float64(1)) == int32(0) {
																			v313 = v299
																		} else {
																			v313 = float64(1)
																		}
																	}
																	v331 = base.I64_reinterpret_f64(v313)
																	m.G0 = v20 + int32(80)
																	return v331
																}
															}
														}
													}
												}
											}
										} else {
											v242 = v65 + int32(8)
											v243 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
											v247 = int32(0)
											v250 = F_tsquery_opr_selec(m, v242, v242+v243*int32(12), v247, v247, float32(0))
											mBase = m.M
											v251 = m.ExcPending
											if v251 != 0 {
												return int64(0)
											} else {
												v267 = v250
												v270 = *(*float32)(unsafe.Add(mBase, uint32(v71+v72)+8))
												v299 = base.F64_mul(v267, base.F64_sub(float64(1), base.F64_promote_f32(v270)))
												v301 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
												if v301 != 0 {
													v302 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
													m.T0[v302].(func(*base.Module, int32))(m, v301)
													mBase = m.M
													v304 = m.ExcPending
													if v304 != 0 {
														return int64(0)
													} else {
														v305 = float64(0)
														if base.F64_lt(v299, v305) != 0 {
															v313 = v305
														} else {
															if base.F64_gt(v299, float64(1)) == int32(0) {
																v313 = v299
															} else {
																v313 = float64(1)
															}
														}
														v331 = base.I64_reinterpret_f64(v313)
														m.G0 = v20 + int32(80)
														return v331
													}
												} else {
													v305 = float64(0)
													if base.F64_lt(v299, v305) != 0 {
														v313 = v305
													} else {
														if base.F64_gt(v299, float64(1)) == int32(0) {
															v313 = v299
														} else {
															v313 = float64(1)
														}
													}
													v331 = base.I64_reinterpret_f64(v313)
													m.G0 = v20 + int32(80)
													return v331
												}
											}
										}
									}
								} else {
									v275 = v65 + int32(8)
									v279 = int32(0)
									v282 = F_tsquery_opr_selec(m, v275, v275+v66*int32(12), v279, v279, float32(0))
									mBase = m.M
									v283 = m.ExcPending
									if v283 != 0 {
										return int64(0)
									} else {
										v299 = v282
										v301 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
										if v301 != 0 {
											v302 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
											m.T0[v302].(func(*base.Module, int32))(m, v301)
											mBase = m.M
											v304 = m.ExcPending
											if v304 != 0 {
												return int64(0)
											} else {
												v305 = float64(0)
												if base.F64_lt(v299, v305) != 0 {
													v313 = v305
												} else {
													if base.F64_gt(v299, float64(1)) == int32(0) {
														v313 = v299
													} else {
														v313 = float64(1)
													}
												}
												v331 = base.I64_reinterpret_f64(v313)
												m.G0 = v20 + int32(80)
												return v331
											}
										} else {
											v305 = float64(0)
											if base.F64_lt(v299, v305) != 0 {
												v313 = v305
											} else {
												if base.F64_gt(v299, float64(1)) == int32(0) {
													v313 = v299
												} else {
													v313 = float64(1)
												}
											}
											v331 = base.I64_reinterpret_f64(v313)
											m.G0 = v20 + int32(80)
											return v331
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
func F_tsrange_subdiff(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v4 int64
	_ = v4
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_reinterpret_f64(base.F64_div(base.F64_sub(base.F64_convert_i64_s(v2), base.F64_convert_i64_s(v4)), float64(1e+06)))
}
func F_tstoreShutdownReceiver(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v3 != 0 {
		F_pfree(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v8 != 0 {
				F_pfree(m, v8)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if v13 != 0 {
						F_free_conversion_map(m, v13)
						mBase = m.M
						v15 = m.ExcPending
						if v15 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
							v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							if v18 != 0 {
								F_ExecDropSingleTupleTableSlot(m, v18)
								mBase = m.M
								v20 = m.ExcPending
								if v20 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
									return
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v18 != 0 {
							F_ExecDropSingleTupleTableSlot(m, v18)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
							return
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if v13 != 0 {
					F_free_conversion_map(m, v13)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v18 != 0 {
							F_ExecDropSingleTupleTableSlot(m, v18)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					if v18 != 0 {
						F_ExecDropSingleTupleTableSlot(m, v18)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
						return
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v8 != 0 {
			F_pfree(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if v13 != 0 {
					F_free_conversion_map(m, v13)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v18 != 0 {
							F_ExecDropSingleTupleTableSlot(m, v18)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
							return
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					if v18 != 0 {
						F_ExecDropSingleTupleTableSlot(m, v18)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
						return
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = int32(0)
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			if v13 != 0 {
				F_free_conversion_map(m, v13)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					if v18 != 0 {
						F_ExecDropSingleTupleTableSlot(m, v18)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				if v18 != 0 {
					F_ExecDropSingleTupleTableSlot(m, v18)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
					return
				}
			}
		}
	}
}
func F_tuplehash_start_iterate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v6 = int32(-1)
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v7 == int64(0) {
		v29 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(v32)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(0)
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10+v12*int32(12))+4))
	if v20 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v29 = v6
	goto L1
L5:
	;
	v29 = v12
	goto L1
L6:
	;
	goto L7
L7:
	;
	v24 = v12 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v24)) < base.Ui64(v7) {
		v12 = v24
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L4
}
func F_tzload(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v513 int32
	_ = v513
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v717 int32
	_ = v717
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v914 int32
	_ = v914
	var v925 int32
	_ = v925
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1141 int32
	_ = v1141
	var v1146 int32
	_ = v1146
	var v1148 int64
	_ = v1148
	var v1149 int64
	_ = v1149
	var v1150 int64
	_ = v1150
	var v1151 int64
	_ = v1151
	var v1152 int64
	_ = v1152
	var v1156 int64
	_ = v1156
	var v1169 int64
	_ = v1169
	var v1176 int64
	_ = v1176
	var v1181 int64
	_ = v1181
	var v1184 int64
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1194 int64
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1273 int32
	_ = v1273
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1336 int32
	_ = v1336
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1374 int32
	_ = v1374
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1404 int32
	_ = v1404
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1441 int32
	_ = v1441
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1517 int32
	_ = v1517
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1556 int32
	_ = v1556
	var v1574 int32
	_ = v1574
	var v1578 int32
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1612 int64
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1630 int32
	_ = v1630
	var v1635 int32
	_ = v1635
	var v1637 int64
	_ = v1637
	var v1638 int64
	_ = v1638
	var v1639 int64
	_ = v1639
	var v1640 int64
	_ = v1640
	var v1641 int64
	_ = v1641
	var v1645 int64
	_ = v1645
	var v1658 int64
	_ = v1658
	var v1665 int64
	_ = v1665
	var v1670 int64
	_ = v1670
	var v1673 int64
	_ = v1673
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1728 int32
	_ = v1728
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1777 int32
	_ = v1777
	var v1781 int32
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1826 int32
	_ = v1826
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1893 int32
	_ = v1893
	var v1901 int32
	_ = v1901
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1914 int32
	_ = v1914
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1941 int32
	_ = v1941
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1978 int32
	_ = v1978
	var v1981 int32
	_ = v1981
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1995 int32
	_ = v1995
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2008 int32
	_ = v2008
	var v2014 int32
	_ = v2014
	var v2038 int32
	_ = v2038
	var v2039 int32
	_ = v2039
	var v2042 int32
	_ = v2042
	var v2048 int32
	_ = v2048
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2066 int32
	_ = v2066
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2073 int32
	_ = v2073
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2090 int32
	_ = v2090
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2153 int32
	_ = v2153
	var v2158 int32
	_ = v2158
	var v2179 int32
	_ = v2179
	var v2184 int32
	_ = v2184
	var v2189 int32
	_ = v2189
	var v2211 int32
	_ = v2211
	var v2217 int32
	_ = v2217
	var v2242 int32
	_ = v2242
	var v2245 int32
	_ = v2245
	var v2248 int32
	_ = v2248
	var v2251 int32
	_ = v2251
	var v2259 int32
	_ = v2259
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2296 int64
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2302 int32
	_ = v2302
	var v2329 int64
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2360 int32
	_ = v2360
	var v2365 int32
	_ = v2365
	var v2366 int64
	_ = v2366
	var v2368 int64
	_ = v2368
	var v2370 int64
	_ = v2370
	var v2374 int32
	_ = v2374
	var v2380 int32
	_ = v2380
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2446 int64
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2477 int32
	_ = v2477
	var v2482 int32
	_ = v2482
	var v2483 int64
	_ = v2483
	var v2485 int64
	_ = v2485
	var v2487 int64
	_ = v2487
	var v2490 int32
	_ = v2490
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2502 int32
	_ = v2502
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2545 int32
	_ = v2545
	var v2547 int32
	_ = v2547
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2576 int32
	_ = v2576
	var v2578 int32
	_ = v2578
	var v2581 int32
	_ = v2581
	var v2582 int64
	_ = v2582
	var v2584 int64
	_ = v2584
	var v2586 int32
	_ = v2586
	var v2592 int32
	_ = v2592
	var v2593 int64
	_ = v2593
	var v2595 int64
	_ = v2595
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2600 int32
	_ = v2600
	var v2605 int32
	_ = v2605
	var v2632 int32
	_ = v2632
	var v2636 int32
	_ = v2636
	var v2638 int32
	_ = v2638
	var v2641 int32
	_ = v2641
	var v2642 int64
	_ = v2642
	var v2644 int64
	_ = v2644
	var v2674 int32
	_ = v2674
	var v2677 int32
	_ = v2677
	var v2683 int32
	_ = v2683
	var v2686 int32
	_ = v2686
	var v2708 int32
	_ = v2708
	var v2711 int32
	_ = v2711
	var v2713 int32
	_ = v2713
	var v2743 int32
	_ = v2743
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2749 int32
	_ = v2749
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2766 int32
	_ = v2766
	var v2769 int32
	_ = v2769
	var v2772 int32
	_ = v2772
	var v2773 int32
	_ = v2773
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2780 int32
	_ = v2780
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2793 int64
	_ = v2793
	var v2794 int64
	_ = v2794
	var v2798 int32
	_ = v2798
	var v2802 int32
	_ = v2802
	var v2816 int32
	_ = v2816
	var v2819 int32
	_ = v2819
	var v2821 int32
	_ = v2821
	var v2850 int32
	_ = v2850
	var v2852 int32
	_ = v2852
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2858 int32
	_ = v2858
	var v2859 int32
	_ = v2859
	var v2861 int32
	_ = v2861
	var v2862 int32
	_ = v2862
	var v2864 int32
	_ = v2864
	var v2865 int32
	_ = v2865
	var v2867 int32
	_ = v2867
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2873 int32
	_ = v2873
	var v2876 int32
	_ = v2876
	var v2879 int32
	_ = v2879
	var v2880 int32
	_ = v2880
	var v2883 int32
	_ = v2883
	var v2884 int32
	_ = v2884
	var v2887 int32
	_ = v2887
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2897 int64
	_ = v2897
	var v2901 int64
	_ = v2901
	var v2905 int32
	_ = v2905
	var v2940 int32
	_ = v2940
	var v2944 int32
	_ = v2944
	var v2973 int32
	_ = v2973
	var v2975 int32
	_ = v2975
	var v2977 int32
	_ = v2977
	var v2981 int32
	_ = v2981
	var v2984 int32
	_ = v2984
	var v3015 int32
	_ = v3015
	var v3019 int32
	_ = v3019
	var v3048 int32
	_ = v3048
	var v3051 int32
	_ = v3051
	var v3053 int32
	_ = v3053
	var v3086 int32
	_ = v3086
	var v3090 int32
	_ = v3090
	var v3121 int32
	_ = v3121
	var v3174 int32
	_ = v3174
	var v3181 int32
	_ = v3181
	v30 = F_emscripten_builtin_malloc(m, int32(_a_F_tzload_0))
	mBase = m.M
	if v30 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_tzload[0]))
	return v34
L2:
	;
	goto L3
L3:
	;
	v36 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)) = uint16(v36)
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v39 = l0
	goto L6
L5:
	;
	v39 = int32(_a_F_tzload_1)
	goto L6
L6:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v43 = v39 + base.B2i32(v40 == int32(58))
	v44 = m.G0
	v46 = v44 - int32(1056)
	m.G0 = v46
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tzload[1])))
	if v49 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_get_share_path(m, int32(_a_F_tzload_2))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v187 = v46 + int32(32)
	v188 = int32(_a_F_tzload_2)
	goto L46
L10:
	;
	return int32(0)
L11:
	;
	v57 = int32(_a_F_tzload_2)
	v58 = F_strlen(m, v57)
	mBase = m.M
	v60 = v58 + v57
	v61 = int32(_a_F_tzload_3)
	v63 = int32(1024) - v58
	if v63 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v183 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_tzload[1])) = uint8(v183)
	goto L9
L13:
	;
	v179 = F_strlen(m, v175)
	mBase = m.M
	goto L12
L14:
	;
	v175 = v61
	goto L13
L15:
	;
	goto L16
L16:
	;
	v69 = v63 - int32(1)
	if (v60^v61)&int32(3) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v172)
	v175 = v168
	goto L13
L18:
	;
	v153 = v148
	v154 = v149
	v155 = v150
	goto L39
L19:
	;
	if v143 == int32(0) {
		v168 = v141
		v169 = v142
		goto L17
	} else {
		goto L38
	}
L20:
	;
	v141 = v61
	v142 = v60
	v143 = v69
	goto L19
L21:
	;
	goto L22
L22:
	;
	v73 = int32(0)
	if int32(0)|base.B2i32(v69 == v73) == v73 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v109 == int32(0) {
		v168 = v106
		v169 = v107
		goto L17
	} else {
		goto L32
	}
L24:
	;
	v85 = v61
	v86 = v60
	v87 = v69
	goto L27
L25:
	;
	goto L26
L26:
	;
	v106 = v61
	v107 = v60
	v108 = v69
	v109 = base.B2i32(v69 != v73)
	goto L23
L27:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v89)
	if v89 == int32(0) {
		v148 = v85
		v149 = v86
		v150 = v87
		goto L18
	} else {
		goto L29
	}
L28:
	;
	v106 = v100
	v107 = v94
	v108 = v96
	v109 = v98
	goto L23
L29:
	;
	v93 = int32(1)
	v94 = v86 + v93
	v96 = v87 - v93
	v97 = int32(0)
	v98 = base.B2i32(v96 != v97)
	v100 = v85 + v93
	if v100&int32(3) == v97 {
		v106 = v100
		v107 = v94
		v108 = v96
		v109 = v98
		goto L23
	} else {
		goto L30
	}
L30:
	;
	if v96 != 0 {
		v85 = v100
		v86 = v94
		v87 = v96
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if base.B2i32(v112 == int32(0))|base.B2i32(base.Ui32(v108) < base.Ui32(int32(4))) != 0 {
		v141 = v106
		v142 = v107
		v143 = v108
		goto L19
	} else {
		goto L33
	}
L33:
	;
	v119 = v106
	v120 = v107
	v121 = v108
	goto L34
L34:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v127 = int32(-2139062144)
	if (int32(16843008)-v124|v124)&v127 != v127 {
		v148 = v119
		v149 = v120
		v150 = v121
		goto L18
	} else {
		goto L36
	}
L35:
	;
	v141 = v135
	v142 = v133
	v143 = v137
	goto L19
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v124
	v132 = int32(4)
	v133 = v120 + v132
	v135 = v119 + v132
	v137 = v121 - v132
	if base.Ui32(int32(3)) < base.Ui32(v137) {
		v119 = v135
		v120 = v133
		v121 = v137
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v148 = v141
	v149 = v142
	v150 = v143
	goto L18
L39:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v157)
	if v157 == int32(0) {
		v168 = v153
		v169 = v154
		goto L17
	} else {
		goto L41
	}
L40:
	;
	v168 = v164
	v169 = v162
	goto L17
L41:
	;
	v161 = int32(1)
	v162 = v154 + v161
	v164 = v153 + v161
	v166 = v155 - v161
	if v166 != 0 {
		v153 = v164
		v154 = v162
		v155 = v166
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v309 = F_strlen(m, v187)
	mBase = m.M
	v310 = F_strlen(m, v43)
	mBase = m.M
	if base.Ui32(int32(1023)) < base.Ui32(v309+v310+int32(1)) {
		v843 = int32(-1)
		goto L74
	} else {
		goto L75
	}
L44:
	;
	v305 = F_strlen(m, v294)
	mBase = m.M
	goto L43
L46:
	;
	goto L47
L47:
	;
	v195 = int32(1023)
	if (v187^v188)&int32(3) != 0 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v298 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v295))) = uint8(v298)
	goto L44
L49:
	;
	v279 = v274
	v280 = v275
	v281 = v276
	goto L70
L50:
	;
	if v269 == int32(0) {
		v294 = v267
		v295 = v268
		goto L48
	} else {
		goto L69
	}
L51:
	;
	v267 = v188
	v268 = v187
	v269 = v195
	goto L50
L52:
	;
	goto L53
L53:
	;
	goto L56
L54:
	;
	goto L63
L56:
	;
	goto L57
L57:
	;
	goto L54
L63:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tzload[2])))
	if base.B2i32(v238 == int32(0))|int32(0) != 0 {
		v267 = v188
		v268 = v187
		v269 = v195
		goto L50
	} else {
		goto L64
	}
L64:
	;
	v245 = v188
	v246 = v187
	v247 = v195
	goto L65
L65:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v253 = int32(-2139062144)
	if (int32(16843008)-v250|v250)&v253 != v253 {
		v274 = v245
		v275 = v246
		v276 = v247
		goto L49
	} else {
		goto L67
	}
L66:
	;
	v267 = v261
	v268 = v259
	v269 = v263
	goto L50
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246))) = v250
	v258 = int32(4)
	v259 = v246 + v258
	v261 = v245 + v258
	v263 = v247 - v258
	if base.Ui32(int32(3)) < base.Ui32(v263) {
		v245 = v261
		v246 = v259
		v247 = v263
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v274 = v267
	v275 = v268
	v276 = v269
	goto L49
L70:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v283)
	if v283 == int32(0) {
		v294 = v279
		v295 = v280
		goto L48
	} else {
		goto L72
	}
L71:
	;
	v294 = v290
	v295 = v288
	goto L48
L72:
	;
	v287 = int32(1)
	v288 = v280 + v287
	v290 = v279 + v287
	v292 = v281 - v287
	if v292 != 0 {
		v279 = v290
		v280 = v288
		v281 = v292
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	m.G0 = v46 + int32(1056)
	if v843 < int32(0) {
		goto L209
	} else {
		goto L210
	}
L75:
	;
	if l1 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v318 = v309 + v187
	v319 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v318))) = uint8(v319)
	v322 = v318 + int32(1)
	if (v43^v322)&int32(3) != 0 {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	goto L78
L78:
	;
	v413 = v309
	v419 = v43
	goto L101
L79:
	;
	v397 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v397
	v402 = F_open(m, v187, v397, v46+int32(16))
	mBase = m.M
	if v397 <= v402 {
		v843 = v402
		goto L74
	} else {
		goto L100
	}
L80:
	;
	goto L79
L81:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v377))) = uint8(v376)
	if v376&int32(255) == int32(0) {
		goto L80
	} else {
		goto L96
	}
L82:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v375 = v43
	v376 = v328
	v377 = v322
	goto L81
L83:
	;
	goto L84
L84:
	;
	if v43&int32(3) != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v332 = v43
	v334 = v322
	goto L88
L86:
	;
	v346 = v43
	v348 = v322
	goto L87
L87:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v353 = int32(-2139062144)
	if (int32(16843008)-v350|v350)&v353 != v353 {
		v375 = v346
		v376 = v350
		v377 = v348
		goto L81
	} else {
		goto L92
	}
L88:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	*(*uint8)(unsafe.Add(mBase, uint32(v334))) = uint8(v335)
	if v335 == int32(0) {
		goto L80
	} else {
		goto L90
	}
L89:
	;
	v346 = v342
	v348 = v340
	goto L87
L90:
	;
	v339 = int32(1)
	v340 = v334 + v339
	v342 = v332 + v339
	if v342&int32(3) != 0 {
		v332 = v342
		v334 = v340
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v358 = v346
	v359 = v350
	v360 = v348
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360))) = v359
	v362 = int32(4)
	v363 = v360 + v362
	v365 = v358 + v362
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	v370 = int32(-2139062144)
	if (int32(16843008)-v367|v367)&v370 == v370 {
		v358 = v365
		v359 = v367
		v360 = v363
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v375 = v365
	v376 = v367
	v377 = v363
	goto L81
L95:
	;
	goto L94
L96:
	;
	v384 = v375
	v386 = v377
	goto L97
L97:
	;
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v386)+1)) = uint8(v387)
	v389 = int32(1)
	if v387 != 0 {
		v384 = v384 + v389
		v386 = v386 + v389
		goto L97
	} else {
		goto L99
	}
L98:
	;
	goto L80
L99:
	;
	goto L98
L100:
	;
	v405 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v318))) = uint8(v405)
	goto L78
L101:
	;
	v437 = int32(47)
	v438 = F___strchrnul(m, v419, v437)
	mBase = m.M
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	if v440 == v437 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	if l1 != 0 {
		goto L175
	} else {
		goto L176
	}
L103:
	;
	v449 = v46 + int32(32)
	v450 = F_AllocateDir(m, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L10
	} else {
		goto L112
	}
L104:
	;
	if v444 != 0 {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	v444 = v438
	goto L107
L106:
	;
	v444 = int32(0)
	goto L107
L107:
	;
	goto L104
L108:
	;
	v447 = v444 - v419
	goto L103
L109:
	;
	goto L110
L110:
	;
	v446 = F_strlen(m, v419)
	mBase = m.M
	v447 = v446
	goto L103
L111:
	;
	if v456 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L112:
	;
	v453 = F_ReadDirExtended(m, v450, v449, int32(15))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	if v453 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v456 = int32(1023) - v413
	v457 = v413 + v449
	v459 = v457 + int32(1)
	v460 = v453
	goto L117
L115:
	;
	goto L116
L116:
	;
	F_FreeDir(m, v450)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L10
	} else {
		goto L141
	}
L117:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460)+19)))
	if v488 == int32(46) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L116
L119:
	;
	v551 = F_ReadDirExtended(m, v450, v46+int32(32), int32(15))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L10
	} else {
		goto L139
	}
L120:
	;
	v492 = v460 + int32(19)
	v493 = F_strlen(m, v492)
	mBase = m.M
	if v493 != v447 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v497 = v492
	v498 = v419
	v499 = v447
	goto L123
L122:
	;
	if v544 == int32(0) {
		goto L111
	} else {
		goto L138
	}
L123:
	;
	if v499 != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v544 = int32(0)
	goto L122
L125:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498))))
	if v502 == v503 {
		v525 = v502
		goto L128
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	goto L124
L128:
	;
	v527 = int32(1)
	if v525 != 0 {
		v497 = v497 + v527
		v498 = v498 + v527
		v499 = v499 - v527
		goto L123
	} else {
		goto L137
	}
L129:
	;
	if base.Ui32((v502-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v513 = v502 | int32(32)
	goto L132
L131:
	;
	v513 = v502
	goto L132
L132:
	;
	if base.Ui32((v503-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v522 = v503 | int32(32)
	goto L135
L134:
	;
	v522 = v503
	goto L135
L135:
	;
	if v513 == v522 {
		v525 = v513
		goto L128
	} else {
		goto L136
	}
L136:
	;
	v544 = v513 - v522
	goto L122
L137:
	;
	goto L127
L138:
	;
	goto L119
L139:
	;
	if v551 != 0 {
		v460 = v551
		goto L117
	} else {
		goto L140
	}
L140:
	;
	goto L118
L141:
	;
	v843 = int32(-1)
	goto L74
L142:
	;
	F_FreeDir(m, v450)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L10
	} else {
		goto L173
	}
L143:
	;
	v699 = F_strlen(m, v695)
	mBase = m.M
	goto L142
L144:
	;
	v695 = v492
	goto L143
L145:
	;
	goto L146
L146:
	;
	v589 = v456 - int32(1)
	if (v459^v492)&int32(3) != 0 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v692 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v689))) = uint8(v692)
	v695 = v688
	goto L143
L148:
	;
	v673 = v668
	v674 = v669
	v675 = v670
	goto L169
L149:
	;
	if v663 == int32(0) {
		v688 = v661
		v689 = v662
		goto L147
	} else {
		goto L168
	}
L150:
	;
	v661 = v492
	v662 = v459
	v663 = v589
	goto L149
L151:
	;
	goto L152
L152:
	;
	v593 = int32(0)
	if base.B2i32(v492&int32(3) == v593)|base.B2i32(v589 == v593) == v593 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if v629 == int32(0) {
		v688 = v626
		v689 = v627
		goto L147
	} else {
		goto L162
	}
L154:
	;
	v605 = v492
	v606 = v459
	v607 = v589
	goto L157
L155:
	;
	goto L156
L156:
	;
	v626 = v492
	v627 = v459
	v628 = v589
	v629 = base.B2i32(v589 != v593)
	goto L153
L157:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605))))
	*(*uint8)(unsafe.Add(mBase, uint32(v606))) = uint8(v609)
	if v609 == int32(0) {
		v668 = v605
		v669 = v606
		v670 = v607
		goto L148
	} else {
		goto L159
	}
L158:
	;
	v626 = v620
	v627 = v614
	v628 = v616
	v629 = v618
	goto L153
L159:
	;
	v613 = int32(1)
	v614 = v606 + v613
	v616 = v607 - v613
	v617 = int32(0)
	v618 = base.B2i32(v616 != v617)
	v620 = v605 + v613
	if v620&int32(3) == v617 {
		v626 = v620
		v627 = v614
		v628 = v616
		v629 = v618
		goto L153
	} else {
		goto L160
	}
L160:
	;
	if v616 != 0 {
		v605 = v620
		v606 = v614
		v607 = v616
		goto L157
	} else {
		goto L161
	}
L161:
	;
	goto L158
L162:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	if base.B2i32(v632 == int32(0))|base.B2i32(base.Ui32(v628) < base.Ui32(int32(4))) != 0 {
		v661 = v626
		v662 = v627
		v663 = v628
		goto L149
	} else {
		goto L163
	}
L163:
	;
	v639 = v626
	v640 = v627
	v641 = v628
	goto L164
L164:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v639)))
	v647 = int32(-2139062144)
	if (int32(16843008)-v644|v644)&v647 != v647 {
		v668 = v639
		v669 = v640
		v670 = v641
		goto L148
	} else {
		goto L166
	}
L165:
	;
	v661 = v655
	v662 = v653
	v663 = v657
	goto L149
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v640))) = v644
	v652 = int32(4)
	v653 = v640 + v652
	v655 = v639 + v652
	v657 = v641 - v652
	if base.Ui32(int32(3)) < base.Ui32(v657) {
		v639 = v655
		v640 = v653
		v641 = v657
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v668 = v661
	v669 = v662
	v670 = v663
	goto L148
L169:
	;
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673))))
	*(*uint8)(unsafe.Add(mBase, uint32(v674))) = uint8(v677)
	if v677 == int32(0) {
		v688 = v673
		v689 = v674
		goto L147
	} else {
		goto L171
	}
L170:
	;
	v688 = v684
	v689 = v682
	goto L147
L171:
	;
	v681 = int32(1)
	v682 = v674 + v681
	v684 = v673 + v681
	v686 = v675 - v681
	if v686 != 0 {
		v673 = v684
		v674 = v682
		v675 = v686
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v704 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v457))) = uint8(v704)
	v706 = int32(1)
	v707 = v413 + v706
	v711 = F_strlen(m, v707+(v46+int32(32)))
	mBase = m.M
	if v444 != 0 {
		v413 = v711 + v707
		v419 = v444 + v706
		goto L101
	} else {
		goto L174
	}
L174:
	;
	goto L102
L175:
	;
	v717 = v309 + v46 + int32(33)
	goto L181
L176:
	;
	goto L177
L177:
	;
	v837 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v837
	v842 = F_open(m, v46+int32(32), v837, v46)
	mBase = m.M
	v843 = v842
	goto L74
L178:
	;
	goto L177
L179:
	;
	v834 = F_strlen(m, v823)
	mBase = m.M
	goto L178
L181:
	;
	goto L182
L182:
	;
	v724 = int32(255)
	if (l1^v717)&int32(3) != 0 {
		goto L186
	} else {
		goto L187
	}
L183:
	;
	v827 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v824))) = uint8(v827)
	goto L179
L184:
	;
	v808 = v803
	v809 = v804
	v810 = v805
	goto L205
L185:
	;
	if v798 == int32(0) {
		v823 = v796
		v824 = v797
		goto L183
	} else {
		goto L204
	}
L186:
	;
	v796 = v717
	v797 = l1
	v798 = v724
	goto L185
L187:
	;
	goto L188
L188:
	;
	v728 = int32(0)
	if base.B2i32(v717&int32(3) == v728)|int32(0) == v728 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	if v764 == int32(0) {
		v823 = v761
		v824 = v762
		goto L183
	} else {
		goto L198
	}
L190:
	;
	v740 = v717
	v741 = l1
	v742 = v724
	goto L193
L191:
	;
	goto L192
L192:
	;
	v761 = v717
	v762 = l1
	v763 = v724
	v764 = int32(1)
	goto L189
L193:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740))))
	*(*uint8)(unsafe.Add(mBase, uint32(v741))) = uint8(v744)
	if v744 == int32(0) {
		v803 = v740
		v804 = v741
		v805 = v742
		goto L184
	} else {
		goto L195
	}
L194:
	;
	v761 = v755
	v762 = v749
	v763 = v751
	v764 = v753
	goto L189
L195:
	;
	v748 = int32(1)
	v749 = v741 + v748
	v751 = v742 - v748
	v752 = int32(0)
	v753 = base.B2i32(v751 != v752)
	v755 = v740 + v748
	if v755&int32(3) == v752 {
		v761 = v755
		v762 = v749
		v763 = v751
		v764 = v753
		goto L189
	} else {
		goto L196
	}
L196:
	;
	if v751 != 0 {
		v740 = v755
		v741 = v749
		v742 = v751
		goto L193
	} else {
		goto L197
	}
L197:
	;
	goto L194
L198:
	;
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761))))
	if base.B2i32(v767 == int32(0))|base.B2i32(base.Ui32(v763) < base.Ui32(int32(4))) != 0 {
		v796 = v761
		v797 = v762
		v798 = v763
		goto L185
	} else {
		goto L199
	}
L199:
	;
	v774 = v761
	v775 = v762
	v776 = v763
	goto L200
L200:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v774)))
	v782 = int32(-2139062144)
	if (int32(16843008)-v779|v779)&v782 != v782 {
		v803 = v774
		v804 = v775
		v805 = v776
		goto L184
	} else {
		goto L202
	}
L201:
	;
	v796 = v790
	v797 = v788
	v798 = v792
	goto L185
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v775))) = v779
	v787 = int32(4)
	v788 = v775 + v787
	v790 = v774 + v787
	v792 = v776 - v787
	if base.Ui32(int32(3)) < base.Ui32(v792) {
		v774 = v790
		v775 = v788
		v776 = v792
		goto L200
	} else {
		goto L203
	}
L203:
	;
	goto L201
L204:
	;
	v803 = v796
	v804 = v797
	v805 = v798
	goto L184
L205:
	;
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v808))))
	*(*uint8)(unsafe.Add(mBase, uint32(v809))) = uint8(v812)
	if v812 == int32(0) {
		v823 = v808
		v824 = v809
		goto L183
	} else {
		goto L207
	}
L206:
	;
	v823 = v819
	v824 = v817
	goto L183
L207:
	;
	v816 = int32(1)
	v817 = v809 + v816
	v819 = v808 + v816
	v821 = v810 - v816
	if v821 != 0 {
		v808 = v819
		v809 = v817
		v810 = v821
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	F_emscripten_builtin_free(m, v30)
	mBase = m.M
	return int32(44)
L210:
	;
	goto L211
L211:
	;
	v880 = F_read(m, v843, v30, int32(_a_F_tzload_4))
	mBase = m.M
	if v880 <= int32(43) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	if v880 < int32(0) {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	goto L214
L214:
	;
	v892 = F_close(m, v843)
	mBase = m.M
	if int32(0) <= v892 {
		goto L218
	} else {
		goto L219
	}
L215:
	;
	v887 = *(*int32)(unsafe.Add(mBase, _c_F_tzload[0]))
	v888 = v887
	goto L217
L216:
	;
	v888 = int32(28)
	goto L217
L217:
	;
	v889 = F_close(m, v843)
	mBase = m.M
	F_emscripten_builtin_free(m, v30)
	mBase = m.M
	return v888
L218:
	;
	v896 = l2 + int32(16)
	v898 = l2 + int32(_a_F_tzload_5)
	v900 = l2 + int32(_a_F_tzload_6)
	v902 = l2 + int32(_a_F_tzload_7)
	v904 = l2 + int32(24)
	v906 = l2 + int32(_a_F_tzload_8)
	v908 = v30 + int32(44)
	v914 = int32(4)
	v925 = v880
	goto L222
L219:
	;
	goto L220
L220:
	;
	v3181 = *(*int32)(unsafe.Add(mBase, _c_F_tzload[0]))
	F_emscripten_builtin_free(m, v30)
	mBase = m.M
	return v3181
L221:
	;
	F_emscripten_builtin_free(m, v30)
	mBase = m.M
	return v3174
L222:
	;
	v938 = int32(28)
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+31)))
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+30)))
	v941 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+28)))
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+29)))
	v947 = int32(8)
	v953 = v939 | (v940|(v941&int32(127)<<(uint(int32(16))%32)|v946<<(uint(v947)%32)))<<(uint(v947)%32)
	if v941 < int32(0) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	if v1877 < int32(3) {
		goto L361
	} else {
		goto L362
	}
L224:
	;
	v958 = v953 | int32(-2147483648)
	goto L226
L225:
	;
	v958 = v953
	goto L226
L226:
	;
	if base.Ui32(int32(49)) < base.Ui32(v958) {
		v3174 = v938
		goto L221
	} else {
		goto L227
	}
L227:
	;
	v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+39)))
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+38)))
	v963 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+36)))
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+37)))
	v969 = int32(8)
	v975 = v961 | (v962|(v963&int32(127)<<(uint(int32(16))%32)|v968<<(uint(v969)%32)))<<(uint(v969)%32)
	if v963 < int32(0) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v980 = v975 | int32(-2147483648)
	goto L230
L229:
	;
	v980 = v975
	goto L230
L230:
	;
	if base.Ui32(int32(255)) < base.Ui32(v980) {
		v3174 = v938
		goto L221
	} else {
		goto L231
	}
L231:
	;
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+35)))
	v984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+34)))
	v985 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+32)))
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+33)))
	v991 = int32(8)
	v997 = v983 | (v984|(v985&int32(127)<<(uint(int32(16))%32)|v990<<(uint(v991)%32)))<<(uint(v991)%32)
	if v985 < int32(0) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v1002 = v997 | int32(-2147483648)
	goto L234
L233:
	;
	v1002 = v997
	goto L234
L234:
	;
	if base.Ui32(int32(1999)) < base.Ui32(v1002) {
		v3174 = v938
		goto L221
	} else {
		goto L235
	}
L235:
	;
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+43)))
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+42)))
	v1007 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+40)))
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+41)))
	v1013 = int32(8)
	v1019 = v1005 | (v1006|(v1007&int32(127)<<(uint(int32(16))%32)|v1012<<(uint(v1013)%32)))<<(uint(v1013)%32)
	if v1007 < int32(0) {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1024 = v1019 | int32(-2147483648)
	goto L238
L237:
	;
	v1024 = v1019
	goto L238
L238:
	;
	if base.Ui32(int32(49)) < base.Ui32(v1024) {
		v3174 = v938
		goto L221
	} else {
		goto L239
	}
L239:
	;
	v1027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+21)))
	v1028 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+20)))
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+23)))
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+27)))
	v1032 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+26)))
	v1033 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30)+24)))
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+25)))
	v1039 = int32(8)
	v1045 = v1031 | (v1032|(v1033&int32(127)<<(uint(int32(16))%32)|v1038<<(uint(v1039)%32)))<<(uint(v1039)%32)
	if v1033 < int32(0) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1050 = v1045 | int32(-2147483648)
	goto L242
L241:
	;
	v1050 = v1045
	goto L242
L242:
	;
	if v980 != v1050 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1053 = v1050
	goto L245
L244:
	;
	v1053 = int32(0)
	goto L245
L245:
	;
	if v1053 != 0 {
		v3174 = v938
		goto L221
	} else {
		goto L246
	}
L246:
	;
	v1058 = int32(8)
	v1064 = (v1028&int32(127)<<(uint(int32(16))%32)|v1027<<(uint(v1058)%32)|v1029)<<(uint(v1058)%32) | v1030
	if v1028 < int32(0) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1069 = v1064 | int32(-2147483648)
	goto L249
L248:
	;
	v1069 = v1064
	goto L249
L249:
	;
	if v980 != v1069 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1072 = v1069
	goto L252
L251:
	;
	v1072 = int32(0)
	goto L252
L252:
	;
	if v1072 != 0 {
		v3174 = v938
		goto L221
	} else {
		goto L253
	}
L253:
	;
	v1075 = v914 + int32(4)
	if v925 < v1050+v1069+v958*v1075+v1002+v1002*v914+v980*int32(6)+v1024+int32(44) {
		v3174 = v938
		goto L221
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v1024
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v980
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1002
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v958
	v1092 = int32(0)
	if v1002 != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1094 = int32(0)
	v1096 = v908
	v1097 = v1094
	v1099 = v1094
	goto L258
L256:
	;
	v1294 = v908
	v1299 = v980
	v1302 = v1092
	goto L257
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1302
	if v1299 <= int32(0) {
		goto L289
	} else {
		goto L290
	}
L258:
	;
	v1124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1096))))
	v1126 = v1124 & int32(127)
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1096)+2)))
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1096)+1)))
	v1129 = base.I32_extend8_s(v1124)
	if v914 == int32(4) {
		goto L261
	} else {
		goto L262
	}
L259:
	;
	v1216 = int32(0)
	if v1216 < v1214 {
		goto L277
	} else {
		goto L278
	}
L260:
	;
	v1185 = v1099 + v906
	v1186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1185))) = uint8(v1186)
	if v1097 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L261:
	;
	v1132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1096)+3)))
	v1135 = int32(8)
	v1141 = v1132 | (v1126<<(uint(int32(16))%32)|v1128<<(uint(v1135)%32)|v1127)<<(uint(v1135)%32)
	if v1129 < int32(0) {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	goto L263
L263:
	;
	v1148 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1096)+7)))
	v1149 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1096)+6)))
	v1150 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1096)+4)))
	v1151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1096)+3)))
	v1152 = int64(8)
	v1156 = int64(16)
	v1169 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1096)+5)))
	v1176 = v1148 | (v1149|((v1150|(v1151<<(uint(v1152)%64)|(base.I64_extend_i32_u(v1127)|(base.I64_extend_i32_u(v1126)<<(uint(v1156)%64)|base.I64_extend_i32_u(v1128)<<(uint(v1152)%64)))<<(uint(v1156)%64)))<<(uint(v1156)%64)|v1169<<(uint(v1152)%64)))<<(uint(v1152)%64)
	if v1129 < int32(0) {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	v1146 = v1141 | int32(-2147483648)
	goto L266
L265:
	;
	v1146 = v1141
	goto L266
L266:
	;
	v1184 = base.I64_extend_i32_s(v1146)
	goto L260
L267:
	;
	v1181 = v1176 | int64(-9223372036854775807-1)
	goto L269
L268:
	;
	v1181 = v1176
	goto L269
L269:
	;
	v1184 = v1181
	goto L260
L270:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v904+v1203<<(uint(int32(3))%32)))) = v1184
	v1209 = v1096 + v914
	v1210 = int32(1)
	v1213 = v1099 + v1210
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1213 < v1214 {
		v1096 = v1209
		v1097 = v1203 + v1210
		v1099 = v1213
		goto L258
	} else {
		goto L276
	}
L271:
	;
	v1203 = int32(0)
	goto L270
L272:
	;
	goto L273
L273:
	;
	v1194 = *(*int64)(unsafe.Add(mBase, uint32(v896+v1097<<(uint(int32(3))%32))))
	if v1194 < v1184 {
		v1203 = v1097
		goto L270
	} else {
		goto L274
	}
L274:
	;
	if v1184 < v1194 {
		v3174 = v938
		goto L221
	} else {
		goto L275
	}
L275:
	;
	v1197 = int32(1)
	v1199 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1185-v1197))) = uint8(v1199)
	v1203 = v1097 - v1197
	goto L270
L276:
	;
	goto L259
L277:
	;
	v1220 = v1209
	v1221 = v1216
	v1228 = v1216
	v1229 = v1214
	goto L280
L278:
	;
	v1265 = v1209
	v1273 = v1216
	goto L279
L279:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v1294 = v1265
	v1299 = v1293
	v1302 = v1273
	goto L257
L280:
	;
	v1248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1220))))
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v1249 <= v1248 {
		v3174 = v938
		goto L221
	} else {
		goto L282
	}
L281:
	;
	v1265 = v1261
	v1273 = v1258
	goto L279
L282:
	;
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1221+v906))))
	if v1252 != 0 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1228+v906))) = uint8(v1248)
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1258 = v1228 + int32(1)
	v1259 = v1255
	goto L285
L284:
	;
	v1258 = v1228
	v1259 = v1229
	goto L285
L285:
	;
	v1260 = int32(1)
	v1261 = v1220 + v1260
	v1263 = v1221 + v1260
	if v1263 < v1259 {
		v1220 = v1261
		v1221 = v1263
		v1228 = v1258
		v1229 = v1259
		goto L280
	} else {
		goto L286
	}
L286:
	;
	goto L281
L287:
	;
	v1574 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1556+v900))) = uint8(v1574)
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v1578 <= v1574 {
		goto L313
	} else {
		goto L314
	}
L288:
	;
	v1422 = int32(3)
	v1423 = v1404 & v1422
	if base.Ui32(v1404-int32(1)) < base.Ui32(v1422) {
		goto L302
	} else {
		goto L303
	}
L289:
	;
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if int32(0) < v1325 {
		v1394 = v1294
		v1404 = v1325
		goto L288
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v1329 = v1294
	v1336 = v1092
	goto L293
L292:
	;
	v1546 = v1294
	v1556 = int32(0)
	goto L287
L293:
	;
	v1359 = v902 + v1336<<(uint(int32(4))%32)
	v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+3)))
	v1361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+2)))
	v1362 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1329))))
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+1)))
	v1368 = int32(8)
	v1374 = v1360 | (v1361|(v1362&int32(127)<<(uint(int32(16))%32)|v1367<<(uint(v1368)%32)))<<(uint(v1368)%32)
	if v1362 < int32(0) {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	v1394 = v1390
	v1404 = v1385
	goto L288
L295:
	;
	v1379 = v1374 | int32(-2147483648)
	goto L297
L296:
	;
	v1379 = v1374
	goto L297
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1359))) = v1379
	v1381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+4)))
	if base.Ui32(int32(1)) < base.Ui32(v1381) {
		v3174 = v938
		goto L221
	} else {
		goto L298
	}
L298:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1359)+4)) = uint8(v1381)
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+5)))
	if v1385 <= v1386 {
		v3174 = v938
		goto L221
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1359)+8)) = v1386
	v1390 = v1329 + int32(6)
	v1392 = v1336 + int32(1)
	if v1392 != v1299 {
		v1329 = v1390
		v1336 = v1392
		goto L293
	} else {
		goto L300
	}
L300:
	;
	goto L294
L301:
	;
	v1508 = v1480
	v1509 = v1481
	v1517 = int32(0)
	goto L309
L302:
	;
	v1480 = v1394
	v1481 = int32(0)
	goto L301
L303:
	;
	goto L304
L304:
	;
	v1432 = int32(0)
	v1434 = v1394
	v1435 = v1432
	v1441 = v1432
	goto L305
L305:
	;
	v1462 = v1435 + v900
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1434))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1462))) = uint8(v1463)
	v1465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1434)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1462)+1)) = uint8(v1465)
	v1467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1434)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1462)+2)) = uint8(v1467)
	v1469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1434)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1462)+3)) = uint8(v1469)
	v1471 = int32(4)
	v1472 = v1435 + v1471
	v1474 = v1434 + v1471
	v1476 = v1441 + v1471
	if v1476 != v1404&int32(-4) {
		v1434 = v1474
		v1435 = v1472
		v1441 = v1476
		goto L305
	} else {
		goto L307
	}
L306:
	;
	if v1423 == int32(0) {
		v1546 = v1474
		v1556 = v1404
		goto L287
	} else {
		goto L308
	}
L307:
	;
	goto L306
L308:
	;
	v1480 = v1474
	v1481 = v1472
	goto L301
L309:
	;
	v1537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1508))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1509+v900))) = uint8(v1537)
	v1539 = int32(1)
	v1542 = v1508 + v1539
	v1544 = v1517 + v1539
	if v1544 != v1423 {
		v1508 = v1542
		v1509 = v1509 + v1539
		v1517 = v1544
		goto L309
	} else {
		goto L311
	}
L310:
	;
	v1546 = v1542
	v1556 = v1404
	goto L287
L311:
	;
	goto L310
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1728
	if int32(0) < v1299 {
		goto L334
	} else {
		goto L335
	}
L313:
	;
	v1718 = v1546
	v1728 = int32(0)
	goto L312
L314:
	;
	goto L315
L315:
	;
	v1583 = int32(0)
	v1585 = v1546
	v1592 = v1583
	v1594 = v1583
	v1612 = int64(0)
	goto L316
L316:
	;
	v1613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1585))))
	v1615 = v1613 & int32(127)
	v1616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1585)+2)))
	v1617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1585)+1)))
	v1618 = base.I32_extend8_s(v1613)
	if v914 == int32(4) {
		goto L319
	} else {
		goto L320
	}
L317:
	;
	v1718 = v1708
	v1728 = v1578
	goto L312
L318:
	;
	if base.B2i32(v1673 < int64(0))|base.B2i32(v1673-v1612 < int64(2419199)) != 0 {
		v3174 = v938
		goto L221
	} else {
		goto L328
	}
L319:
	;
	v1621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1585)+3)))
	v1624 = int32(8)
	v1630 = v1621 | (v1615<<(uint(int32(16))%32)|v1617<<(uint(v1624)%32)|v1616)<<(uint(v1624)%32)
	if v1618 < int32(0) {
		goto L322
	} else {
		goto L323
	}
L320:
	;
	goto L321
L321:
	;
	v1637 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1585)+7)))
	v1638 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1585)+6)))
	v1639 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1585)+4)))
	v1640 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1585)+3)))
	v1641 = int64(8)
	v1645 = int64(16)
	v1658 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1585)+5)))
	v1665 = v1637 | (v1638|((v1639|(v1640<<(uint(v1641)%64)|(base.I64_extend_i32_u(v1616)|(base.I64_extend_i32_u(v1615)<<(uint(v1645)%64)|base.I64_extend_i32_u(v1617)<<(uint(v1641)%64)))<<(uint(v1645)%64)))<<(uint(v1645)%64)|v1658<<(uint(v1641)%64)))<<(uint(v1641)%64)
	if v1618 < int32(0) {
		goto L325
	} else {
		goto L326
	}
L322:
	;
	v1635 = v1630 | int32(-2147483648)
	goto L324
L323:
	;
	v1635 = v1630
	goto L324
L324:
	;
	v1673 = base.I64_extend_i32_s(v1635)
	goto L318
L325:
	;
	v1670 = v1665 | int64(-9223372036854775807-1)
	goto L327
L326:
	;
	v1670 = v1665
	goto L327
L327:
	;
	v1673 = v1670
	goto L318
L328:
	;
	v1680 = v1585 + v914
	v1681 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1680))))
	v1686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1680)+1)))
	v1687 = int32(8)
	v1690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1680)+2)))
	v1694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1680)+3)))
	v1695 = (v1681&int32(127)<<(uint(int32(16))%32)|v1686<<(uint(v1687)%32)|v1690)<<(uint(v1687)%32) | v1694
	if v1681 < int32(0) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1700 = v1695 | int32(-2147483648)
	goto L331
L330:
	;
	v1700 = v1695
	goto L331
L331:
	;
	v1701 = int32(1)
	if base.B2i32(v1700 != v1592-v1701)&base.B2i32(v1700 != v1592+v1701) != 0 {
		v3174 = v938
		goto L221
	} else {
		goto L332
	}
L332:
	;
	v1708 = v1585 + v1075
	v1711 = v898 + v1594<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1711)+8)) = base.I64_extend_i32_s(v1700)
	*(*int64)(unsafe.Add(mBase, uint32(v1711))) = v1673
	v1716 = v1594 + int32(1)
	if v1716 != v1578 {
		v1585 = v1708
		v1592 = v1700
		v1594 = v1716
		v1612 = v1673
		goto L316
	} else {
		goto L333
	}
L333:
	;
	goto L317
L334:
	;
	v1749 = v1718
	v1752 = v1574
	goto L337
L335:
	;
	v1840 = v1718
	goto L336
L336:
	;
	v1868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
	if v1868 != 0 {
		goto L353
	} else {
		goto L354
	}
L337:
	;
	v1777 = int32(0)
	if v1050 == v1777 {
		goto L340
	} else {
		goto L341
	}
L338:
	;
	v1795 = v1786
	v1796 = v1777
	goto L345
L339:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v902+v1752<<(uint(int32(4))%32))+12)) = uint8(v1787)
	v1793 = v1752 + int32(1)
	if v1793 != v1299 {
		v1749 = v1786
		v1752 = v1793
		goto L337
	} else {
		goto L344
	}
L340:
	;
	v1786 = v1749
	v1787 = int32(0)
	goto L339
L341:
	;
	goto L342
L342:
	;
	v1781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1749))))
	if base.Ui32(int32(1)) < base.Ui32(v1781) {
		v3174 = v938
		goto L221
	} else {
		goto L343
	}
L343:
	;
	v1786 = v1749 + int32(1)
	v1787 = v1781
	goto L339
L344:
	;
	goto L338
L345:
	;
	if v1069 == int32(0) {
		goto L348
	} else {
		goto L349
	}
L346:
	;
	v1840 = v1831
	goto L336
L347:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v902+v1796<<(uint(int32(4))%32))+13)) = uint8(v1832)
	v1838 = v1796 + int32(1)
	if v1838 != v1299 {
		v1795 = v1831
		v1796 = v1838
		goto L345
	} else {
		goto L352
	}
L348:
	;
	v1831 = v1795
	v1832 = int32(0)
	goto L347
L349:
	;
	goto L350
L350:
	;
	v1826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1795))))
	if base.Ui32(int32(1)) < base.Ui32(v1826) {
		v3174 = v938
		goto L221
	} else {
		goto L351
	}
L351:
	;
	v1831 = v1795 + int32(1)
	v1832 = v1826
	goto L347
L352:
	;
	goto L346
L353:
	;
	v1870 = v30 - v1840 + v925
	if v1870 != 0 {
		goto L356
	} else {
		goto L357
	}
L354:
	;
	v1877 = v925
	goto L355
L355:
	;
	goto L223
L356:
	;
	base.MemoryCopy(m, v30, v1840, v1870)
	goto L358
L357:
	;
	goto L358
L358:
	;
	if base.Ui32(v914) < base.Ui32(int32(5)) {
		v914 = v914 << (uint(int32(1)) % 32)
		v925 = v1870
		goto L222
	} else {
		goto L359
	}
L359:
	;
	v1877 = v1870
	goto L355
L360:
	;
	if v2686 < int32(2) {
		goto L451
	} else {
		goto L452
	}
L361:
	;
	v2674 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v2674 == int32(0) {
		v3174 = v938
		goto L221
	} else {
		goto L450
	}
L362:
	;
	v1880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v1880 != int32(10) {
		goto L361
	} else {
		goto L363
	}
L363:
	;
	v1885 = v30 + v1877 - int32(1)
	v1886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1885))))
	if v1886 != int32(10) {
		goto L361
	} else {
		goto L364
	}
L364:
	;
	if int32(256) < v1299+int32(2) {
		v2683 = v1299
		v2686 = v1302
		goto L360
	} else {
		goto L365
	}
L365:
	;
	v1893 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1885))) = uint8(v1893)
	v1901 = F_tzparse(m, v30+int32(1), v30+int32(_a_F_tzload_4), v1893)
	mBase = m.M
	if v1901 == v1893 {
		goto L361
	} else {
		goto L366
	}
L366:
	;
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tzload[3])))
	if int32(0) < v1905 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1914 = v1904
	v1919 = v1893
	v1920 = int32(0)
	goto L370
L368:
	;
	v2184 = v1904
	v2189 = v1893
	goto L369
L369:
	;
	if v1905 != v2189 {
		goto L361
	} else {
		goto L413
	}
L370:
	;
	v1941 = v30 + v1920<<(uint(int32(4))%32)
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1941)+uint32(_c_F_tzload[4])))
	v1945 = v30 + int32(_a_F_tzload_9) + v1944
	v1946 = int32(0)
	if v1946 < v1914 {
		goto L374
	} else {
		goto L375
	}
L371:
	;
	v2184 = v2153
	v2189 = v2158
	goto L369
L372:
	;
	v2179 = v1920 + int32(1)
	if v2179 < v1905 {
		v1914 = v2153
		v1919 = v2158
		v1920 = v2179
		goto L370
	} else {
		goto L412
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1941)+uint32(_c_F_tzload[4]))) = v2123
	v2153 = v2122
	v2158 = v1919 + int32(1)
	goto L372
L374:
	;
	v1950 = v1946
	goto L377
L375:
	;
	v2014 = v1946
	goto L376
L376:
	;
	v2038 = F_strlen(m, v1945)
	mBase = m.M
	v2039 = v2038 + v2014
	if int32(49) < v2039 {
		v2153 = v1914
		v2158 = v1919
		goto L372
	} else {
		goto L390
	}
L377:
	;
	v1978 = v1950 + v900
	v1981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1978))))
	v1984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1945))))
	if base.B2i32(v1981 == int32(0))|base.B2i32(v1981 != v1984) != 0 {
		v2002 = v1981
		v2003 = v1984
		goto L380
	} else {
		goto L381
	}
L378:
	;
	v2014 = v1914
	goto L376
L379:
	;
	if v2002-v2003 == int32(0) {
		goto L386
	} else {
		goto L387
	}
L380:
	;
	goto L379
L381:
	;
	v1987 = v1978
	v1988 = v1945
	goto L382
L382:
	;
	v1991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1988)+1)))
	v1992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1987)+1)))
	if v1992 == int32(0) {
		v2002 = v1992
		v2003 = v1991
		goto L380
	} else {
		goto L384
	}
L383:
	;
	v2002 = v1992
	v2003 = v1991
	goto L380
L384:
	;
	v1995 = int32(1)
	if v1992 == v1991 {
		v1987 = v1987 + v1995
		v1988 = v1988 + v1995
		goto L382
	} else {
		goto L385
	}
L385:
	;
	goto L383
L386:
	;
	v2122 = v1914
	v2123 = v1950
	goto L373
L387:
	;
	goto L388
L388:
	;
	v2008 = v1950 + int32(1)
	if v2008 != v1914 {
		v1950 = v2008
		goto L377
	} else {
		goto L389
	}
L389:
	;
	goto L378
L390:
	;
	v2042 = v2014 + v900
	if (v1945^v2042)&int32(3) != 0 {
		goto L394
	} else {
		goto L395
	}
L391:
	;
	v2122 = v2039 + int32(1)
	v2123 = v2014
	goto L373
L392:
	;
	goto L391
L393:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2097))) = uint8(v2096)
	if v2096&int32(255) == int32(0) {
		goto L392
	} else {
		goto L408
	}
L394:
	;
	v2048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1945))))
	v2095 = v1945
	v2096 = v2048
	v2097 = v2042
	goto L393
L395:
	;
	goto L396
L396:
	;
	if v1945&int32(3) != 0 {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v2052 = v1945
	v2054 = v2042
	goto L400
L398:
	;
	v2066 = v1945
	v2068 = v2042
	goto L399
L399:
	;
	v2070 = *(*int32)(unsafe.Add(mBase, uint32(v2066)))
	v2073 = int32(-2139062144)
	if (int32(16843008)-v2070|v2070)&v2073 != v2073 {
		v2095 = v2066
		v2096 = v2070
		v2097 = v2068
		goto L393
	} else {
		goto L404
	}
L400:
	;
	v2055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2052))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2054))) = uint8(v2055)
	if v2055 == int32(0) {
		goto L392
	} else {
		goto L402
	}
L401:
	;
	v2066 = v2062
	v2068 = v2060
	goto L399
L402:
	;
	v2059 = int32(1)
	v2060 = v2054 + v2059
	v2062 = v2052 + v2059
	if v2062&int32(3) != 0 {
		v2052 = v2062
		v2054 = v2060
		goto L400
	} else {
		goto L403
	}
L403:
	;
	goto L401
L404:
	;
	v2078 = v2066
	v2079 = v2070
	v2080 = v2068
	goto L405
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2080))) = v2079
	v2082 = int32(4)
	v2083 = v2080 + v2082
	v2085 = v2078 + v2082
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v2078)+4))
	v2090 = int32(-2139062144)
	if (int32(16843008)-v2087|v2087)&v2090 == v2090 {
		v2078 = v2085
		v2079 = v2087
		v2080 = v2083
		goto L405
	} else {
		goto L407
	}
L406:
	;
	v2095 = v2085
	v2096 = v2087
	v2097 = v2083
	goto L393
L407:
	;
	goto L406
L408:
	;
	v2104 = v2095
	v2106 = v2097
	goto L409
L409:
	;
	v2107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2104)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2106)+1)) = uint8(v2107)
	v2109 = int32(1)
	if v2107 != 0 {
		v2104 = v2104 + v2109
		v2106 = v2106 + v2109
		goto L409
	} else {
		goto L411
	}
L410:
	;
	goto L392
L411:
	;
	goto L410
L412:
	;
	goto L371
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v2184
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v2211 < int32(2) {
		v2259 = v2211
		goto L414
	} else {
		goto L415
	}
L414:
	;
	v2284 = int32(0)
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_tzload[5])))
	if base.B2i32(v2259 == v2284)|base.B2i32(v2287 <= v2284) != 0 {
		v2380 = v2284
		goto L421
	} else {
		goto L422
	}
L415:
	;
	v2217 = v2211
	goto L416
L416:
	;
	v2242 = v2217 + v906
	v2245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2242-int32(1)))))
	v2248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2242-int32(2)))))
	if v2245 != v2248 {
		v2259 = v2217
		goto L414
	} else {
		goto L418
	}
L417:
	;
	v2259 = int32(1)
	goto L414
L418:
	;
	v2251 = v2217 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2251
	if int32(2) < v2217 {
		v2217 = v2251
		goto L416
	} else {
		goto L419
	}
L419:
	;
	goto L417
L420:
	;
	if v1905 <= int32(0) {
		goto L361
	} else {
		goto L442
	}
L421:
	;
	if v2287 <= v2380 {
		goto L420
	} else {
		goto L432
	}
L422:
	;
	v2296 = *(*int64)(unsafe.Add(mBase, uint32(v896+v2259<<(uint(int32(3))%32))))
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v2302 = v2284
	goto L423
L423:
	;
	v2329 = *(*int64)(unsafe.Add(mBase, uint32(v30+int32(_a_F_tzload_10)+v2302<<(uint(int32(3))%32))))
	v2330 = v2297
	goto L426
L424:
	;
	goto L420
L425:
	;
	if v2296 < v2329+v2370 {
		v2380 = v2302
		goto L421
	} else {
		goto L430
	}
L426:
	;
	v2360 = v2330 - int32(1)
	if v2360 < int32(0) {
		v2370 = int64(0)
		goto L425
	} else {
		goto L428
	}
L427:
	;
	v2368 = *(*int64)(unsafe.Add(mBase, uint32(v2365)+8))
	v2370 = v2368
	goto L425
L428:
	;
	v2365 = v898 + v2360<<(uint(int32(4))%32)
	v2366 = *(*int64)(unsafe.Add(mBase, uint32(v2365)))
	if v2329 < v2366 {
		v2330 = v2360
		goto L426
	} else {
		goto L429
	}
L429:
	;
	goto L427
L430:
	;
	v2374 = v2302 + int32(1)
	if v2374 != v2287 {
		v2302 = v2374
		goto L423
	} else {
		goto L431
	}
L431:
	;
	goto L424
L432:
	;
	v2412 = v2259
	v2413 = v2380
	goto L433
L433:
	;
	if int32(1999) < v2412 {
		goto L420
	} else {
		goto L435
	}
L434:
	;
	goto L420
L435:
	;
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v2440 = int32(3)
	v2446 = *(*int64)(unsafe.Add(mBase, uint32(v30+int32(_a_F_tzload_10)+v2413<<(uint(v2440)%32))))
	v2447 = v2439
	goto L437
L436:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v904+v2412<<(uint(v2440)%32)))) = v2446 + v2487
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v2493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+int32(_a_F_tzload_11)+v2413))))
	v2494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v2495 = v2493 + v2494
	*(*uint8)(unsafe.Add(mBase, uint32(v906+v2490))) = uint8(v2495)
	v2497 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v2498 = int32(1)
	v2499 = v2497 + v2498
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2499
	v2502 = v2413 + v2498
	if v2502 != v2287 {
		v2412 = v2499
		v2413 = v2502
		goto L433
	} else {
		goto L441
	}
L437:
	;
	v2477 = v2447 - int32(1)
	if v2477 < int32(0) {
		v2487 = int64(0)
		goto L436
	} else {
		goto L439
	}
L438:
	;
	v2485 = *(*int64)(unsafe.Add(mBase, uint32(v2482)+8))
	v2487 = v2485
	goto L436
L439:
	;
	v2482 = v898 + v2477<<(uint(int32(4))%32)
	v2483 = *(*int64)(unsafe.Add(mBase, uint32(v2482)))
	if v2446 < v2483 {
		v2447 = v2477
		goto L437
	} else {
		goto L440
	}
L440:
	;
	goto L438
L441:
	;
	goto L434
L442:
	;
	v2535 = v30 + int32(_a_F_tzload_12)
	v2536 = int32(0)
	if v1905 != int32(1) {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	v2545 = v2536
	v2547 = int32(0)
	goto L446
L444:
	;
	v2605 = v2536
	goto L445
L445:
	;
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2632 + int32(1)
	v2636 = int32(4)
	v2638 = v902 + v2632<<(uint(v2636)%32)
	v2641 = v2535 + v2605<<(uint(v2636)%32)
	v2642 = *(*int64)(unsafe.Add(mBase, uint32(v2641)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2638)+8)) = v2642
	v2644 = *(*int64)(unsafe.Add(mBase, uint32(v2641)))
	*(*int64)(unsafe.Add(mBase, uint32(v2638))) = v2644
	goto L361
L446:
	;
	v2572 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2573 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2572 + v2573
	v2576 = int32(4)
	v2578 = v902 + v2572<<(uint(v2576)%32)
	v2581 = v2535 + v2545<<(uint(v2576)%32)
	v2582 = *(*int64)(unsafe.Add(mBase, uint32(v2581)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2578)+8)) = v2582
	v2584 = *(*int64)(unsafe.Add(mBase, uint32(v2581)))
	*(*int64)(unsafe.Add(mBase, uint32(v2578))) = v2584
	v2586 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v2586 + v2573
	v2592 = v902 + v2586<<(uint(v2576)%32)
	v2593 = *(*int64)(unsafe.Add(mBase, uint32(v2581)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v2592)+8)) = v2593
	v2595 = *(*int64)(unsafe.Add(mBase, uint32(v2581)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2592))) = v2595
	v2597 = int32(2)
	v2598 = v2545 + v2597
	v2600 = v2547 + v2597
	if v2600 != v1905&int32(2147483646) {
		v2545 = v2598
		v2547 = v2600
		goto L446
	} else {
		goto L448
	}
L447:
	;
	if v1905&int32(1) == int32(0) {
		goto L361
	} else {
		goto L449
	}
L448:
	;
	goto L447
L449:
	;
	v2605 = v2598
	goto L445
L450:
	;
	v2677 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v2683 = v2674
	v2686 = v2677
	goto L360
L451:
	;
	v2940 = int32(0)
	if v2686 <= v2940 {
		v3121 = v2940
		goto L492
	} else {
		goto L493
	}
L452:
	;
	v2708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+uint32(_c_F_tzload[6]))))
	v2711 = v902 + v2708<<(uint(int32(4))%32)
	v2713 = int32(1)
	goto L453
L453:
	;
	if v2683 <= v2708 {
		goto L456
	} else {
		goto L457
	}
L454:
	;
	v2816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2686+v906-int32(1)))))
	v2819 = v902 + v2816<<(uint(int32(4))%32)
	v2821 = v2686 - int32(2)
	goto L473
L455:
	;
	goto L454
L456:
	;
	v2802 = v2713 + int32(1)
	if v2802 != v2686 {
		v2713 = v2802
		goto L453
	} else {
		goto L472
	}
L457:
	;
	v2743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2713+v906))))
	if v2683 <= v2743 {
		goto L456
	} else {
		goto L458
	}
L458:
	;
	v2747 = v902 + v2743<<(uint(int32(4))%32)
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(v2747)))
	v2749 = *(*int32)(unsafe.Add(mBase, uint32(v2711)))
	if v2748 != v2749 {
		goto L456
	} else {
		goto L459
	}
L459:
	;
	v2751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2747)+4)))
	v2752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2711)+4)))
	if v2751 != v2752 {
		goto L456
	} else {
		goto L460
	}
L460:
	;
	v2754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2747)+12)))
	v2755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2711)+12)))
	if v2754 != v2755 {
		goto L456
	} else {
		goto L461
	}
L461:
	;
	v2757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2747)+13)))
	v2758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2711)+13)))
	if v2757 != v2758 {
		goto L456
	} else {
		goto L462
	}
L462:
	;
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(v2747)+8))
	v2761 = v900 + v2760
	v2762 = *(*int32)(unsafe.Add(mBase, uint32(v2711)+8))
	v2763 = v900 + v2762
	v2766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2761))))
	v2769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2763))))
	if base.B2i32(v2766 == int32(0))|base.B2i32(v2766 != v2769) != 0 {
		v2787 = v2766
		v2788 = v2769
		goto L464
	} else {
		goto L465
	}
L463:
	;
	if v2787-v2788 != 0 {
		goto L456
	} else {
		goto L470
	}
L464:
	;
	goto L463
L465:
	;
	v2772 = v2761
	v2773 = v2763
	goto L466
L466:
	;
	v2776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2773)+1)))
	v2777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2772)+1)))
	if v2777 == int32(0) {
		v2787 = v2777
		v2788 = v2776
		goto L464
	} else {
		goto L468
	}
L467:
	;
	v2787 = v2777
	v2788 = v2776
	goto L464
L468:
	;
	v2780 = int32(1)
	if v2777 == v2776 {
		v2772 = v2772 + v2780
		v2773 = v2773 + v2780
		goto L466
	} else {
		goto L469
	}
L469:
	;
	goto L467
L470:
	;
	v2793 = *(*int64)(unsafe.Add(mBase, uint32(v904+v2713<<(uint(int32(3))%32))))
	v2794 = *(*int64)(unsafe.Add(mBase, uint32(v904)))
	if v2793-v2794 != int64(12622780800) {
		goto L456
	} else {
		goto L471
	}
L471:
	;
	v2798 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v896))) = uint8(v2798)
	goto L455
L472:
	;
	goto L455
L473:
	;
	if v2683 <= v2816 {
		goto L475
	} else {
		goto L476
	}
L474:
	;
	goto L451
L475:
	;
	if int32(0) < v2821 {
		v2821 = v2821 - int32(1)
		goto L473
	} else {
		goto L491
	}
L476:
	;
	v2850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2821+v906))))
	if v2683 <= v2850 {
		goto L475
	} else {
		goto L477
	}
L477:
	;
	v2852 = *(*int32)(unsafe.Add(mBase, uint32(v2819)))
	v2855 = v902 + v2850<<(uint(int32(4))%32)
	v2856 = *(*int32)(unsafe.Add(mBase, uint32(v2855)))
	if v2852 != v2856 {
		goto L475
	} else {
		goto L478
	}
L478:
	;
	v2858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2819)+4)))
	v2859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855)+4)))
	if v2858 != v2859 {
		goto L475
	} else {
		goto L479
	}
L479:
	;
	v2861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2819)+12)))
	v2862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855)+12)))
	if v2861 != v2862 {
		goto L475
	} else {
		goto L480
	}
L480:
	;
	v2864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2819)+13)))
	v2865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2855)+13)))
	if v2864 != v2865 {
		goto L475
	} else {
		goto L481
	}
L481:
	;
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(v2819)+8))
	v2868 = v900 + v2867
	v2869 = *(*int32)(unsafe.Add(mBase, uint32(v2855)+8))
	v2870 = v900 + v2869
	v2873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2868))))
	v2876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2870))))
	if base.B2i32(v2873 == int32(0))|base.B2i32(v2873 != v2876) != 0 {
		v2894 = v2873
		v2895 = v2876
		goto L483
	} else {
		goto L484
	}
L482:
	;
	if v2894-v2895 != 0 {
		goto L475
	} else {
		goto L489
	}
L483:
	;
	goto L482
L484:
	;
	v2879 = v2868
	v2880 = v2870
	goto L485
L485:
	;
	v2883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2880)+1)))
	v2884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2879)+1)))
	if v2884 == int32(0) {
		v2894 = v2884
		v2895 = v2883
		goto L483
	} else {
		goto L487
	}
L486:
	;
	v2894 = v2884
	v2895 = v2883
	goto L483
L487:
	;
	v2887 = int32(1)
	if v2884 == v2883 {
		v2879 = v2879 + v2887
		v2880 = v2880 + v2887
		goto L485
	} else {
		goto L488
	}
L488:
	;
	goto L486
L489:
	;
	v2897 = *(*int64)(unsafe.Add(mBase, uint32(v904+v2686<<(uint(int32(3))%32)-int32(8))))
	v2901 = *(*int64)(unsafe.Add(mBase, uint32(v904+v2821<<(uint(int32(3))%32))))
	if v2897-v2901 != int64(12622780800) {
		goto L475
	} else {
		goto L490
	}
L490:
	;
	v2905 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+17)) = uint8(v2905)
	goto L451
L491:
	;
	goto L474
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+uint32(_c_F_tzload[7]))) = v3121
	v3174 = v2940
	goto L221
L493:
	;
	v2944 = v2940
	goto L495
L494:
	;
	v3121 = int32(0)
	goto L492
L495:
	;
	v2973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2944+v906))))
	if v2973 != 0 {
		goto L497
	} else {
		goto L498
	}
L496:
	;
	v2977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+uint32(_c_F_tzload[6]))))
	v2981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902+v2977<<(uint(int32(4))%32))+4)))
	if v2981 != int32(1) {
		goto L501
	} else {
		goto L502
	}
L497:
	;
	v2975 = v2944 + int32(1)
	if v2686 != v2975 {
		v2944 = v2975
		goto L495
	} else {
		goto L500
	}
L498:
	;
	goto L499
L499:
	;
	goto L496
L500:
	;
	goto L494
L501:
	;
	v3048 = int32(1)
	if v2683 <= v3048 {
		goto L507
	} else {
		goto L508
	}
L502:
	;
	v2984 = v2977
	goto L503
L503:
	;
	if v2984 <= int32(0) {
		goto L501
	} else {
		goto L505
	}
L504:
	;
	v3121 = v3015
	goto L492
L505:
	;
	v3015 = v2984 - int32(1)
	v3019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902+v3015<<(uint(int32(4))%32))+4)))
	if v3019 != 0 {
		v2984 = v3015
		goto L503
	} else {
		goto L506
	}
L506:
	;
	goto L504
L507:
	;
	v3051 = v3048
	goto L509
L508:
	;
	v3051 = v2683
	goto L509
L509:
	;
	v3053 = int32(0)
	goto L510
L510:
	;
	v3086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v3053<<(uint(int32(4))%32))+uint32(_c_F_tzload[8]))))
	if v3086 != int32(1) {
		v3121 = v3053
		goto L492
	} else {
		goto L512
	}
L511:
	;
	goto L494
L512:
	;
	v3090 = v3053 + int32(1)
	if v3090 != v3051 {
		v3053 = v3090
		goto L510
	} else {
		goto L513
	}
L513:
	;
	goto L511
}
