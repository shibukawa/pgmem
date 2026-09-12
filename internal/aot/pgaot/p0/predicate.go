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
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v221 int64
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
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
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[132]))
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
	v66 = *(*int32)(unsafe.Add(mBase, _consts[636]))
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
	v53 = *(*int32)(unsafe.Add(mBase, _consts[132]))
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
	v70 = *(*int32)(unsafe.Add(mBase, _consts[132]))
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
	v84 = *(*int32)(unsafe.Add(mBase, _consts[130]))
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
	v88 = v14 + int32(40)
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v88))) = v89
	v91 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v97 = v93
	v100 = v2
	goto L25
L25:
	;
	if v97 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	if v100&int32(1) != 0 {
		goto L45
	} else {
		goto L46
	}
L27:
	;
	goto L26
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	v125 = F_hash_search(m, v119, v14+int32(32), int32(1), v14+int32(15))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L33
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v107 == int32(-1) {
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v112 = v111
	goto L28
L32:
	;
	v112 = int32(-1)
	goto L28
L33:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	if v127 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v141 != 0 {
		v158 = int32(0)
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v130 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v125)+20)) = v130
	v133 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v125)+16)) = uint8(v133)
	v139 = v130
	goto L34
L36:
	;
	goto L37
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v125)+20))
	v137 = v135 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v125)+20)) = v137
	v139 = v137
	goto L34
L38:
	;
	if v139 <= v158 {
		v97 = v141
		goto L25
	} else {
		goto L44
	}
L39:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v142 == int32(-1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _consts[637]))
	if int32(0) <= v146 {
		v158 = v146
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[638]))
	v158 = v157
	goto L38
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	v153 = base.I32_div_s(v150, int32(0)-v146)
	v158 = v153 - int32(1)
	goto L38
L44:
	;
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v160
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v162
	v97 = v141
	v100 = int32(1)
	goto L25
L45:
	;
	F_PredicateLockAcquire(m, v14+int32(16))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v171 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	goto L1
L49:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v177 = F_LWLockAcquire(m, v173+int32(3840), int32(1))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v185 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+72))
	if v186 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v188&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v188 = int32(1)
	goto L54
L53:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+76)))
	v188 = v187
	goto L54
L54:
	;
	goto L51
L55:
	;
	v194 = F_LWLockAcquire(m, v180+int32(72), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v180)+52))
	if v196 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+72))
	if v327 != 0 {
		goto L90
	} else {
		goto L91
	}
L60:
	;
	v200 = v180 + int32(48)
	if v196 == v200 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v204 = v196
	goto L62
L62:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v204-int32(16))))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v216
	v218 = base.I32_wrap_i64(v216)
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v218)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v219
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v218)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v221
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v223 != v224 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L59
L64:
	;
	if v213 != v200 {
		v204 = v213
		goto L62
	} else {
		goto L88
	}
L65:
	;
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if v226 != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+44)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v229 == v230 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v232 = v227
	goto L69
L68:
	;
	v232 = int32(0)
	goto L69
L69:
	;
	if v232 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if v229 != int32(-1) {
		goto L64
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v239 != v240 {
		goto L64
	} else {
		goto L75
	}
L73:
	;
	if v230 == int32(-1) {
		goto L64
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[636]))
	v246 = F_get_hash_value(m, v243, v14+int32(32))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v256 = v249 + v246&int32(15)<<(uint(int32(7))%32) + int32(25344)
	v258 = F_LWLockAcquire(m, v256, int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v260)+4)) = v261
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v263
	v266 = v204 - int32(8)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v268 = int32(4)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v204-v268)))
	*(*int32)(unsafe.Add(mBase, uint32(v267)+4)) = v270
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = v272
	v275 = *(*int32)(unsafe.Add(mBase, _consts[639]))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v284 = F_hash_search_with_hash_value(m, v275, v14+int32(16), v246^v278<<(uint(v268)%32), int32(2), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v218)+20))
	if v286 != v218+int32(16) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v291 = v286
	goto L81
L80:
	;
	v291 = int32(0)
	goto L81
L81:
	;
	if v291 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _consts[636]))
	v298 = F_hash_search_with_hash_value(m, v295, v218, v246, int32(2), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	F_LWLockRelease(m, v256)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L2
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	F_DecrementParentLocks(m, v14+int32(32))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	goto L64
L88:
	;
	goto L63
L89:
	;
	if v329&int32(1) != 0 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v329 = int32(1)
	goto L92
L91:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+76)))
	v329 = v328
	goto L92
L92:
	;
	goto L89
L93:
	;
	F_LWLockRelease(m, v180+int32(72))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L2
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v337+int32(3840))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L2
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	goto L1
}
func F_PredicateLockPageCombine(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v5 int32
	_ = v5
	F_PredicateLockPageSplit(m, l0, l1, l2)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
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
	v10 = *(*int32)(unsafe.Add(mBase, _consts[130]))
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
				if base.Ui32(v21) < base.Ui32(int32(12000)) {
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
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
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
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
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
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
	return v361
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
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	m.T0[v149].(func(*base.Module, int32, int32))(m, v16, v10+int32(28))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
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
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	m.T0[v121].(func(*base.Module, int32, int32))(m, v16, v10+int32(28))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
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
	v361 = base.B2i32(v46 == int32(0))
	goto L6
L24:
	;
	goto L26
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v88].(func(*base.Module, int32))(m, v10+int32(8))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L33
	}
L26:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v72 = m.T0[v71].(func(*base.Module, int32) int32)(m, v10+int32(8))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L28
	}
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v82].(func(*base.Module, int32))(m, v10+int32(8))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
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
	v361 = int32(1)
	goto L6
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	m.T0[v93].(func(*base.Module, int32, int32))(m, v16, v10+int32(28))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	goto L35
L35:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v106 = m.T0[v105].(func(*base.Module, int32) int32)(m, v10+int32(28))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L37
	}
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	m.T0[v116].(func(*base.Module, int32))(m, v10+int32(28))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L43
	}
L37:
	;
	if v106 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v108 = F_predicate_implied_by_recurse(m, v106, l1, l2)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
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
	if v108 == int32(0) {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v361 = base.B2i32(v106 != int32(0))
	goto L6
L44:
	;
	goto L45
L45:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v134 = m.T0[v133].(func(*base.Module, int32) int32)(m, v10+int32(28))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	m.T0[v144].(func(*base.Module, int32))(m, v10+int32(28))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L53
	}
L47:
	;
	if v134 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v136 = F_predicate_implied_by_recurse(m, v134, l1, l2)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
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
	if v136 == int32(0) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v361 = base.B2i32(v134 != int32(0))
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
	v154 = int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v158 = m.T0[v157].(func(*base.Module, int32) int32)(m, v10+int32(28))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
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
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	m.T0[v217].(func(*base.Module, int32))(m, v10+int32(28))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L76
	}
L59:
	;
	if v158 == int32(0) {
		v213 = v154
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v166 = v158
	goto L61
L61:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	m.T0[v171].(func(*base.Module, int32, int32))(m, l1, v10+int32(8))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L63
	}
L62:
	;
	v213 = v154
	goto L58
L63:
	;
	goto L64
L64:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v184 = m.T0[v183].(func(*base.Module, int32) int32)(m, v10+int32(8))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L66
	}
L65:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v200].(func(*base.Module, int32))(m, v10+int32(8))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L73
	}
L66:
	;
	if v184 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v190].(func(*base.Module, int32))(m, v10+int32(8))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v194 = F_predicate_implied_by_recurse(m, v166, v184, l2)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L71
	}
L70:
	;
	v213 = int32(0)
	goto L58
L71:
	;
	if v194 == int32(0) {
		goto L64
	} else {
		goto L72
	}
L72:
	;
	goto L65
L73:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v206 = m.T0[v205].(func(*base.Module, int32) int32)(m, v10+int32(28))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	if v206 != 0 {
		v166 = v206
		goto L61
	} else {
		goto L75
	}
L75:
	;
	goto L62
L76:
	;
	v361 = v213
	goto L6
L77:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v230 = m.T0[v229].(func(*base.Module, int32) int32)(m, v10+int32(28))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L79
	}
L78:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	m.T0[v238].(func(*base.Module, int32))(m, v10+int32(28))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L85
	}
L79:
	;
	if v230 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v232 = F_predicate_implied_by_recurse(m, v230, l1, l2)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
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
	if v232 != 0 {
		goto L77
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v361 = base.B2i32(v230 == int32(0))
	goto L6
L86:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v298 != 0 {
		goto L109
	} else {
		goto L110
	}
L87:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	m.T0[v271].(func(*base.Module, int32, int32))(m, l1, v10+int32(8))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L99
	}
L88:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	m.T0[v245].(func(*base.Module, int32, int32))(m, l1, v10+int32(8))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	goto L90
L90:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v258 = m.T0[v257].(func(*base.Module, int32) int32)(m, v10+int32(8))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L92
	}
L91:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v266].(func(*base.Module, int32))(m, v10+int32(8))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L98
	}
L92:
	;
	if v258 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v260 = F_predicate_implied_by_recurse(m, v16, v258, l2)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
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
	if v260 != 0 {
		goto L90
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v361 = base.B2i32(v258 == int32(0))
	goto L6
L99:
	;
	goto L100
L100:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v284 = m.T0[v283].(func(*base.Module, int32) int32)(m, v10+int32(8))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L102
	}
L101:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	m.T0[v294].(func(*base.Module, int32))(m, v10+int32(8))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L108
	}
L102:
	;
	if v284 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v286 = F_predicate_implied_by_recurse(m, v16, v284, l2)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
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
	if v286 == int32(0) {
		goto L100
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v361 = base.B2i32(v284 != int32(0))
	goto L6
L109:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v301 = int32(1)
	v302 = F_equal(m, l1, v16)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	if v302 != 0 {
		v361 = v301
		goto L6
	} else {
		goto L114
	}
L114:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v304 != int32(17) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v341 != int32(52) {
		goto L132
	} else {
		goto L133
	}
L116:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v341 = v340
	goto L115
L117:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v307 != int32(91) {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	if v312 == int32(0) {
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	if v315 != int32(7) {
		goto L116
	} else {
		goto L120
	}
L120:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+24)))
	if v318 != 0 {
		goto L116
	} else {
		goto L121
	}
L121:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v312)+20))
	if v320 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v321 = F_equal(m, l1, v319)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
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
	if v321 == int32(0) {
		goto L116
	} else {
		goto L126
	}
L126:
	;
	v361 = v301
	goto L6
L127:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v327 != int32(21) {
		v341 = v327
		goto L115
	} else {
		goto L128
	}
L128:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v330 != int32(2) {
		goto L116
	} else {
		goto L129
	}
L129:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+12))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	v336 = F_equal(m, v335, v319)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	if v336 != 0 {
		v361 = v301
		goto L6
	} else {
		goto L131
	}
L131:
	;
	goto L116
L132:
	;
	v354 = F_operator_predicate_proof(m, l1, v16, int32(0), l2)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L139
	}
L133:
	;
	if l2 != 0 {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v345 != int32(1) {
		goto L132
	} else {
		goto L135
	}
L135:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v348 != 0 {
		goto L132
	} else {
		goto L136
	}
L136:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v351 = F_clause_is_strict_for(m, v16, v349, int32(1))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	if v351 != 0 {
		v361 = v301
		goto L6
	} else {
		goto L138
	}
L138:
	;
	goto L132
L139:
	;
	v361 = v354
	goto L6
}
