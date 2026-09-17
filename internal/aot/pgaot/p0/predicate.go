package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PredicateLockAcquire(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
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
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
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
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v211 int64
	_ = v211
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
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
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[0]))
	v20 = F_hash_search(m, v17, l0, v2, v2)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return
L2:
	;
	return
L3:
	;
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+16)))
	if v22 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v25
	goto L8
L7:
	;
	goto L6
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v38 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[1]))
	v67 = F_get_hash_value(m, v66, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L19
	}
L10:
	;
	goto L9
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v47
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[0]))
	v58 = F_hash_search(m, v53, v14+int32(32), v48, v48)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L16
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v41 == int32(-1) {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v46 = v45
	goto L11
L15:
	;
	v46 = int32(-1)
	goto L11
L16:
	;
	if v58 == int32(0) {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+16)))
	if v62 != int32(1) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L1
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[0]))
	v74 = F_hash_search_with_hash_value(m, v70, l0, v67, int32(1), v14+int32(14))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v76 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+16)) = uint8(v76)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)))
	if v78 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = int32(0)
	goto L23
L22:
	;
	goto L23
L23:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[2]))
	F_CreatePredicateLock(m, l0, v67, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v87
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v95 = v91
	v97 = v2
	goto L25
L25:
	;
	if v95 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	if v97 != 0 {
		goto L45
	} else {
		goto L46
	}
L27:
	;
	goto L26
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = int32(0)
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[0]))
	v123 = F_hash_search(m, v117, v14+int32(32), int32(1), v14+int32(15))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L2
	} else {
		goto L33
	}
L29:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v105 == int32(-1) {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v110 = v109
	goto L28
L32:
	;
	v110 = int32(-1)
	goto L28
L33:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	if v125 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v139 != 0 {
		v156 = int32(0)
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v128 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v123)+20)) = v128
	v131 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+16)) = uint8(v131)
	v137 = v128
	goto L34
L36:
	;
	goto L37
L37:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v123)+20))
	v135 = v133 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v123)+20)) = v135
	v137 = v135
	goto L34
L38:
	;
	if v137 <= v156 {
		v95 = v139
		goto L25
	} else {
		goto L44
	}
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v140 == int32(-1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[3]))
	if int32(0) <= v144 {
		v156 = v144
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[4]))
	v156 = v155
	goto L38
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[5]))
	v151 = base.I32_div_s(v148, int32(0)-v144)
	v156 = v151 - int32(1)
	goto L38
L44:
	;
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v158
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v160
	v95 = v139
	v97 = int32(1)
	goto L25
L45:
	;
	F_PredicateLockAcquire(m, v14+int32(16))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L2
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v167 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	goto L1
L49:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[6]))
	v173 = F_LWLockAcquire(m, v169+int32(3840), int32(1))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[2]))
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[7]))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+72))
	if v180 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v183&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v183 = int32(1)
	goto L54
L53:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+76)))
	v183 = v182
	goto L54
L54:
	;
	goto L51
L55:
	;
	v189 = F_LWLockAcquire(m, v176+int32(72), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v176)+52))
	if v191 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[7]))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+72))
	if v322 != 0 {
		goto L86
	} else {
		goto L87
	}
L60:
	;
	v195 = v176 + int32(48)
	if v191 == v195 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v199 = v191
	goto L62
L62:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v199-int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v211
	v213 = base.I32_wrap_i64(v211)
	v214 = *(*int64)(unsafe.Add(mBase, uint32(v213)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v214
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v213)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v218 != v219 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L59
L64:
	;
	if v208 != v195 {
		v199 = v208
		goto L62
	} else {
		goto L84
	}
L65:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if v221 != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+44)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v224 == v225 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v227 = v222
	goto L69
L68:
	;
	v227 = int32(0)
	goto L69
L69:
	;
	v230 = int32(-1)
	if base.B2i32(v227 == int32(0))&(base.B2i32(v225 == v230)|base.B2i32(v224 != v230)) != 0 {
		goto L64
	} else {
		goto L70
	}
L70:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v236 != v237 {
		goto L64
	} else {
		goto L71
	}
L71:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[1]))
	v243 = F_get_hash_value(m, v240, v14+int32(32))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[6]))
	v253 = v246 + v243&int32(15)<<(uint(int32(7))%32) + int32(_a_F_PredicateLockAcquire_0)
	v255 = F_LWLockAcquire(m, v253, int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v257)+4)) = v258
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v260
	v263 = v199 - int32(8)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v265 = int32(4)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v199-v265)))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = v269
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[8]))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v281 = F_hash_search_with_hash_value(m, v272, v14+int32(16), v243^v275<<(uint(v265)%32), int32(2), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v213)+20))
	if v283 != v213+int32(16) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v288 = v283
	goto L77
L76:
	;
	v288 = int32(0)
	goto L77
L77:
	;
	if v288 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[1]))
	v295 = F_hash_search_with_hash_value(m, v292, v213, v243, int32(2), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L2
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_LWLockRelease(m, v253)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	F_DecrementParentLocks(m, v14+int32(32))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	goto L64
L84:
	;
	goto L63
L85:
	;
	if v325&int32(1) != 0 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v325 = int32(1)
	goto L88
L87:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+76)))
	v325 = v324
	goto L88
L88:
	;
	goto L85
L89:
	;
	F_LWLockRelease(m, v176+int32(72))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L2
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockAcquire[6]))
	F_LWLockRelease(m, v333+int32(3840))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L2
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	goto L1
}
func F_PredicateLockRelation(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_PredicateLockRelation[0]))
	if v10 == int32(0) {
		m.G0 = v7 + int32(16)
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		switch v13 {
		case 0, 5:
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+108)))
			if v14&int32(128) != 0 {
				F_ReleasePredicateLocks(m, int32(0), int32(1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if base.Ui32(v21) < base.Ui32(int32(_a_F_PredicateLockRelation_0)) {
					m.G0 = v7 + int32(16)
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+118)))
					if v25 == int32(116) {
						m.G0 = v7 + int32(16)
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(4294967295)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v21
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v28
						F_PredicateLockAcquire(m, v7)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		default:
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_predicate_implied_by(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	if l0 == int32(0) {
		return int32(1)
	} else {
		if l1 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v12 == int32(1) {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v17 = v16
			} else {
				v17 = l0
			}
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v18 == int32(1) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v23 = v22
			} else {
				v23 = l1
			}
			v24 = F_predicate_implied_by_recurse(m, v23, v17, l2)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				return v24
			}
		}
	}
}
func F_predicate_implied_by_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
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
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
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
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(318) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = v15
	goto L3
L2:
	;
	v16 = l0
	goto L3
L3:
	;
	v19 = F_predicate_classify(m, l1, v10+int32(8))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v25 = F_predicate_classify(m, v16, v10+int32(28))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L10
	}
L6:
	;
	m.G0 = v10 + int32(48)
	return v355
L7:
	;
	switch v19 - int32(1) {
	case 0:
		goto L88
	case 1:
		goto L87
	default:
		goto L86
	}
L8:
	;
	v146 = v10 + int32(28)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	m.T0[v147].(func(*base.Module, int32, int32))(m, v16, v146)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L54
	}
L9:
	;
	switch v19 - int32(1) {
	case 0:
		goto L13
	case 1:
		goto L12
	default:
		goto L11
	}
L10:
	;
	switch v25 - int32(1) {
	case 0:
		goto L9
	case 1:
		goto L8
	default:
		goto L7
	}
L11:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	m.T0[v119].(func(*base.Module, int32, int32))(m, v16, v10+int32(28))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L44
	}
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	m.T0[v59].(func(*base.Module, int32, int32))(m, l1, v10+int32(8))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L24
	}
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	m.T0[v33].(func(*base.Module, int32, int32))(m, l1, v10+int32(8))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L15
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v46 = m.T0[v45].(func(*base.Module, int32) int32)(m, v10+int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v54].(func(*base.Module, int32))(m, v10+int32(8))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L23
	}
L17:
	;
	if v46 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v48 = F_predicate_implied_by_recurse(m, v16, v46, l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L16
L21:
	;
	if v48 != 0 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v355 = base.B2i32(v46 == int32(0))
	goto L6
L24:
	;
	goto L26
L25:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v86].(func(*base.Module, int32))(m, v10+int32(8))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L33
	}
L26:
	;
	v70 = v10 + int32(8)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v72 = m.T0[v71].(func(*base.Module, int32) int32)(m, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L28
	}
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v80].(func(*base.Module, int32))(m, v70)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L32
	}
L28:
	;
	if v72 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v76 = F_predicate_implied_by_recurse(m, v16, v72, l2)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	if v76 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	v355 = int32(1)
	goto L6
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	m.T0[v91].(func(*base.Module, int32, int32))(m, v16, v10+int32(28))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L35
L35:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v104 = m.T0[v103].(func(*base.Module, int32) int32)(m, v10+int32(28))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L37
	}
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	m.T0[v114].(func(*base.Module, int32))(m, v10+int32(28))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L43
	}
L37:
	;
	if v104 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v106 = F_predicate_implied_by_recurse(m, v104, l1, l2)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	goto L36
L41:
	;
	if v106 == int32(0) {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v355 = base.B2i32(v104 != int32(0))
	goto L6
L44:
	;
	goto L45
L45:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v132 = m.T0[v131].(func(*base.Module, int32) int32)(m, v10+int32(28))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	m.T0[v142].(func(*base.Module, int32))(m, v10+int32(28))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L53
	}
L47:
	;
	if v132 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v134 = F_predicate_implied_by_recurse(m, v132, l1, l2)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L46
L51:
	;
	if v134 == int32(0) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v355 = base.B2i32(v132 != int32(0))
	goto L6
L54:
	;
	if v19 == int32(2) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v152 = int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v154 = m.T0[v153].(func(*base.Module, int32) int32)(m, v146)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	goto L77
L58:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	m.T0[v211].(func(*base.Module, int32))(m, v10+int32(28))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L76
	}
L59:
	;
	if v154 == int32(0) {
		v206 = v152
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v163 = v154
	goto L61
L61:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	m.T0[v167].(func(*base.Module, int32, int32))(m, l1, v10+int32(8))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L63
	}
L62:
	;
	v206 = v152
	goto L58
L63:
	;
	goto L64
L64:
	;
	v178 = v10 + int32(8)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v180 = m.T0[v179].(func(*base.Module, int32) int32)(m, v178)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L66
	}
L65:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v194].(func(*base.Module, int32))(m, v10+int32(8))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L73
	}
L66:
	;
	if v180 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v184].(func(*base.Module, int32))(m, v178)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v188 = F_predicate_implied_by_recurse(m, v163, v180, l2)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L71
	}
L70:
	;
	v206 = int32(0)
	goto L58
L71:
	;
	if v188 == int32(0) {
		goto L64
	} else {
		goto L72
	}
L72:
	;
	goto L65
L73:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v200 = m.T0[v199].(func(*base.Module, int32) int32)(m, v10+int32(28))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	if v200 != 0 {
		v163 = v200
		goto L61
	} else {
		goto L75
	}
L75:
	;
	goto L62
L76:
	;
	v355 = v206
	goto L6
L77:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v224 = m.T0[v223].(func(*base.Module, int32) int32)(m, v10+int32(28))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L79
	}
L78:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	m.T0[v232].(func(*base.Module, int32))(m, v10+int32(28))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L85
	}
L79:
	;
	if v224 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v226 = F_predicate_implied_by_recurse(m, v224, l1, l2)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	goto L78
L83:
	;
	if v226 != 0 {
		goto L77
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v355 = base.B2i32(v224 == int32(0))
	goto L6
L86:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_predicate_implied_by_recurse[0]))
	if v292 != 0 {
		goto L109
	} else {
		goto L110
	}
L87:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	m.T0[v265].(func(*base.Module, int32, int32))(m, l1, v10+int32(8))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L99
	}
L88:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	m.T0[v239].(func(*base.Module, int32, int32))(m, l1, v10+int32(8))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	goto L90
L90:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v252 = m.T0[v251].(func(*base.Module, int32) int32)(m, v10+int32(8))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L92
	}
L91:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v260].(func(*base.Module, int32))(m, v10+int32(8))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L98
	}
L92:
	;
	if v252 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v254 = F_predicate_implied_by_recurse(m, v16, v252, l2)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	goto L91
L96:
	;
	if v254 != 0 {
		goto L90
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v355 = base.B2i32(v252 == int32(0))
	goto L6
L99:
	;
	goto L100
L100:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v278 = m.T0[v277].(func(*base.Module, int32) int32)(m, v10+int32(8))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L102
	}
L101:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v288].(func(*base.Module, int32))(m, v10+int32(8))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L108
	}
L102:
	;
	if v278 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v280 = F_predicate_implied_by_recurse(m, v16, v278, l2)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	goto L101
L106:
	;
	if v280 == int32(0) {
		goto L100
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v355 = base.B2i32(v278 != int32(0))
	goto L6
L109:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v295 = int32(1)
	v296 = F_equal(m, l1, v16)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	if v296 != 0 {
		v355 = v295
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v298 != int32(17) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if base.B2i32(v335 != int32(52))|l2 != 0 {
		goto L132
	} else {
		goto L133
	}
L116:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v335 = v334
	goto L115
L117:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v301 != int32(91) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+12))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	if v306 == int32(0) {
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	if v309 != int32(7) {
		goto L116
	} else {
		goto L120
	}
L120:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+24)))
	if v312 != 0 {
		goto L116
	} else {
		goto L121
	}
L121:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v306)+20))
	if v314 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v315 = F_equal(m, l1, v313)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if l1 == int32(0) {
		goto L116
	} else {
		goto L127
	}
L125:
	;
	if v315 == int32(0) {
		goto L116
	} else {
		goto L126
	}
L126:
	;
	v355 = v295
	goto L6
L127:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v321 != int32(21) {
		v335 = v321
		goto L115
	} else {
		goto L128
	}
L128:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v324 != int32(2) {
		goto L116
	} else {
		goto L129
	}
L129:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	v330 = F_equal(m, v329, v313)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	if v330 != 0 {
		v355 = v295
		goto L6
	} else {
		goto L131
	}
L131:
	;
	goto L116
L132:
	;
	v349 = F_operator_predicate_proof(m, l1, v16, int32(0), l2)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L4
	} else {
		goto L138
	}
L133:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v340 != int32(1) {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v343 != 0 {
		goto L132
	} else {
		goto L135
	}
L135:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v346 = F_clause_is_strict_for(m, v16, v344, int32(1))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	if v346 != 0 {
		v355 = v295
		goto L6
	} else {
		goto L137
	}
L137:
	;
	goto L132
L138:
	;
	v355 = v349
	goto L6
}
