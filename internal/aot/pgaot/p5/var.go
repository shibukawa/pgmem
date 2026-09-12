package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_record_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	if v9 != v5 {
		if v9 == int32(0) {
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
			if v14 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v13
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v13
			}
			if v13 == int32(0) {
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v19
			}
		}
		if v5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = v5
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = v26
			if v26 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v4
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v4
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(0)
		}
	} else {
	}
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v39 != 0 {
		F_DeleteExpandedObject(m, v39+int32(12))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = l2
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = l2
		return
	}
}
func F_makeVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = l1
	v9 = F_palloc0(m, int32(48))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l5
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
		*(*uint16)(unsafe.Add(mBase, uint32(v9)+8)) = uint16(v2)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(6)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(-1)
		*(*uint16)(unsafe.Add(mBase, uint32(v9)+40)) = uint16(v2)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v13
		return v9
	}
}
func F_set_var_from_str(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v181 int64
	_ = v181
	var v193 int64
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v222 int64
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v423 int32
	_ = v423
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	v6 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v20 - int32(43) {
	case 0:
		goto L3
	default:
		v28 = l1
		v29 = v6
		goto L1
	case 2:
		goto L2
	}
L1:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v32 = base.B2i32(v30 == int32(46))
	v33 = v28 + v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if base.Ui32(int32(9)) < base.Ui32((v34-int32(48))&int32(255)) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v28 = l1 + int32(1)
	v29 = int32(16384)
	goto L1
L3:
	;
	v28 = l1 + int32(1)
	v29 = v6
	goto L1
L4:
	;
	m.G0 = v18 + int32(16)
	return v534
L5:
	;
	v514 = int32(0)
	v515 = F_errsave_start(m, l4)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L7
	} else {
		goto L91
	}
L6:
	;
	v41 = F_strlen(m, v33)
	mBase = m.M
	v44 = F_palloc(m, v41+int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v48
	v50 = int32(4)
	v51 = int32(-1)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v52 == v48 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v137 = v129 + v44
	v138 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v137))) = uint16(v138)
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+2)) = uint8(v138)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v142|int32(32) != int32(101) {
		goto L29
	} else {
		goto L30
	}
L10:
	;
	v127 = v33
	v129 = v50
	v130 = v51
	v132 = v6
	goto L9
L11:
	;
	goto L12
L12:
	;
	v56 = v52
	v60 = v33
	v61 = v32
	v62 = v50
	v63 = v51
	v65 = v6
	goto L13
L13:
	;
	v71 = v56 - int32(48)
	if base.Ui32(v71&int32(255)) <= base.Ui32(int32(9)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v127 = v115
	v129 = v117
	v130 = v118
	v132 = v119
	goto L9
L15:
	;
	if v114&int32(255) != 0 {
		v56 = v114
		v60 = v115
		v61 = v116
		v62 = v117
		v63 = v118
		v65 = v119
		goto L13
	} else {
		goto L26
	}
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v44))) = uint8(v71)
	v78 = int32(1)
	v89 = v60 + v78
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v114 = v90
	v115 = v89
	v116 = v61
	v117 = v62 + v78
	v118 = v63 + (v61^int32(-1))&v78
	v119 = v65 + v61&v78
	goto L15
L17:
	;
	goto L18
L18:
	;
	v92 = v56 & int32(255)
	if v92 != int32(95) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v92 != int32(46) {
		v127 = v60
		v129 = v62
		v130 = v63
		v132 = v65
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v106 = v60 + int32(1)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if base.Ui32(int32(9)) < base.Ui32((v107-int32(48))&int32(255)) {
		goto L5
	} else {
		goto L25
	}
L22:
	;
	if v61&int32(1) != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v99 = int32(1)
	v101 = v60 + v99
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v102 != int32(95) {
		v114 = v102
		v115 = v101
		v116 = v99
		v117 = v62
		v118 = v63
		v119 = v65
		goto L15
	} else {
		goto L24
	}
L24:
	;
	goto L5
L25:
	;
	v114 = v107
	v115 = v106
	v116 = v61
	v117 = v62
	v118 = v63
	v119 = v65
	goto L15
L26:
	;
	goto L14
L27:
	;
	v482 = int32(0)
	v483 = F_errsave_start(m, l4)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L7
	} else {
		goto L86
	}
L28:
	;
	if int32(0) <= v238 {
		goto L53
	} else {
		goto L54
	}
L29:
	;
	v238 = v130
	v239 = v127
	v240 = v132
	goto L28
L30:
	;
	goto L31
L31:
	;
	v147 = int32(0)
	v149 = v127 + int32(1)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	switch v150 - int32(43) {
	case 0:
		goto L34
	default:
		v158 = v147
		v159 = v149
		goto L32
	case 2:
		goto L33
	}
L32:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if base.Ui32(int32(9)) < base.Ui32((v160-int32(48))&int32(255)) {
		goto L5
	} else {
		goto L35
	}
L33:
	;
	v158 = int32(1)
	v159 = v127 + int32(2)
	goto L32
L34:
	;
	v158 = v147
	v159 = v127 + int32(2)
	goto L32
L35:
	;
	v168 = v160
	v176 = v159
	v181 = int64(0)
	goto L36
L36:
	;
	if base.Ui32((v168-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	if v158 != 0 {
		goto L47
	} else {
		goto L48
	}
L38:
	;
	goto L37
L39:
	;
	if v212&int32(255) != 0 {
		v168 = v212
		v176 = v213
		v181 = v214
		goto L36
	} else {
		goto L46
	}
L40:
	;
	v193 = v181*int64(10) + base.I64_extend_i32_u(v168)&int64(15)
	if int64(1073741823) < v193 {
		goto L27
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v168&int32(255) != int32(95) {
		v218 = v176
		v219 = v181
		goto L38
	} else {
		goto L44
	}
L43:
	;
	v197 = v176 + int32(1)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v212 = v198
	v213 = v197
	v214 = v193
	goto L39
L44:
	;
	v204 = v176 + int32(1)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if base.Ui32(int32(9)) < base.Ui32((v205-int32(48))&int32(255)) {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v212 = v205
	v213 = v204
	v214 = v181
	goto L39
L46:
	;
	v218 = v213
	v219 = v214
	goto L38
L47:
	;
	v222 = int64(0) - v219
	goto L49
L48:
	;
	v222 = v219
	goto L49
L49:
	;
	v223 = base.I32_wrap_i64(v222)
	v224 = v132 - v223
	v225 = int32(0)
	if v225 < v224 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v228 = v224
	goto L52
L51:
	;
	v228 = v225
	goto L52
L52:
	;
	v238 = v223 + v130
	v239 = v218
	v240 = v228
	goto L28
L53:
	;
	v247 = int32(4)
	v250 = base.I32_div_s(v238+v247, v247)
	v255 = v250 - int32(1)
	goto L55
L54:
	;
	v255 = v238 >> (uint(int32(2)) % 32)
	goto L55
L55:
	;
	v256 = int32(2)
	v258 = v255<<(uint(v256)%32) - v238
	v259 = v129 + v258
	v261 = v259 + v256
	v263 = base.I32_div_s(v261, int32(4))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v264 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_pfree(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L7
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v271 = F_palloc(m, v263<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L7
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v271
	v274 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v271))) = uint16(v274)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v263
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v255
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v282 = v280 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v282
	if v261 < int32(4) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_pfree(m, v44)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L7
	} else {
		goto L70
	}
L62:
	;
	v286 = int32(1)
	v287 = v286 - v258
	if v263&v286 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v290 = v287 + v44
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	v292 = int32(10)
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)))
	v303 = ((v291*v292+v294)*v292+v298)*v292 + v302
	*(*uint16)(unsafe.Add(mBase, uint32(v280)+2)) = uint16(v303)
	v311 = v280 + int32(4)
	v312 = int32(5) - v258
	v313 = v263 - int32(1)
	goto L65
L64:
	;
	v311 = v282
	v312 = v287
	v313 = v263
	goto L65
L65:
	;
	if base.Ui32(v259-int32(2)) < base.Ui32(int32(4)) {
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v323 = v311
	v324 = v312
	v325 = v313
	goto L67
L67:
	;
	v333 = v324 + v44
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	v335 = int32(10)
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+1)))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+2)))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+3)))
	v346 = ((v334*v335+v337)*v335+v341)*v335 + v345
	*(*uint16)(unsafe.Add(mBase, uint32(v323))) = uint16(v346)
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+7)))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+6)))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+5)))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+4)))
	v360 = v348 + (v349+(v350+v351*v335)*v335)*v335
	*(*uint16)(unsafe.Add(mBase, uint32(v323)+2)) = uint16(v360)
	v366 = int32(2)
	if v366 < v325 {
		v323 = v323 + int32(4)
		v324 = v324 + int32(8)
		v325 = v325 - v366
		goto L67
	} else {
		goto L69
	}
L68:
	;
	goto L61
L69:
	;
	goto L68
L70:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v388 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v464
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v239
	v534 = int32(1)
	goto L4
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v464 = int32(0)
	v468 = v450
	goto L71
L73:
	;
	v395 = v388
	v399 = v387
	goto L77
L74:
	;
	goto L75
L75:
	;
	if v388 != 0 {
		v464 = v388
		v468 = v387
		goto L71
	} else {
		goto L85
	}
L76:
	;
	v423 = v395
	goto L81
L77:
	;
	v409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399))))
	if v409 != 0 {
		goto L76
	} else {
		goto L79
	}
L78:
	;
	v450 = v387 + v388<<(uint(int32(1))%32)
	goto L72
L79:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v411 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v410 - v411
	if v411 < v395 {
		v395 = v395 - v411
		v399 = v399 + int32(2)
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v440 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v399-int32(2)+v423<<(uint(int32(1))%32)))))
	if v440 != 0 {
		v464 = v423
		v468 = v399
		goto L71
	} else {
		goto L83
	}
L82:
	;
	v450 = v399
	goto L72
L83:
	;
	v441 = int32(1)
	if v441 < v423 {
		v423 = v423 - v441
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v450 = v387
	goto L72
L86:
	;
	if v483 == int32(0) {
		v534 = v482
		goto L4
	} else {
		goto L87
	}
L87:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	F_errmsg(m, int32(118998), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	F_errsave_finish(m, l4, int32(520760), int32(7320), int32(215533))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v534 = v482
	goto L4
L91:
	;
	if v515 == int32(0) {
		v534 = v514
		goto L4
	} else {
		goto L92
	}
L92:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L7
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(509593)
	F_errmsg(m, int32(753253), v18)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	F_errsave_finish(m, l4, int32(520760), int32(7326), int32(215533))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	v534 = v514
	goto L4
}
