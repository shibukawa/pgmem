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
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v38 != 0 {
		F_DeleteExpandedObject(m, v38+int32(12))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
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
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v176 int64
	_ = v176
	var v188 int64
	_ = v188
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int64
	_ = v211
	var v214 int64
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v443 int32
	_ = v443
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
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
	v29 = int32(_a_F_set_var_from_str_0)
	goto L1
L3:
	;
	v28 = l1 + int32(1)
	v29 = v6
	goto L1
L4:
	;
	m.G0 = v18 + int32(16)
	return v527
L5:
	;
	v507 = int32(0)
	v508 = F_errsave_start(m, l4)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L7
	} else {
		goto L90
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
		v118 = v33
		v124 = v50
		v125 = v51
		v128 = v6
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v132 = v124 + v44
	v133 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v132)+2)) = uint8(v133)
	*(*uint16)(unsafe.Add(mBase, uint32(v132))) = uint16(v133)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v137|int32(32) != int32(101) {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	v56 = v33
	v60 = v52
	v61 = v32
	v62 = v50
	v63 = v51
	v66 = v6
	goto L11
L11:
	;
	v71 = v60 - int32(48)
	if base.Ui32(v71&int32(255)) <= base.Ui32(int32(9)) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v118 = v114
	v124 = v110
	v125 = v111
	v128 = v112
	goto L9
L13:
	;
	v114 = v56 + int32(1)
	if v108&int32(255) != 0 {
		v56 = v114
		v60 = v108
		v61 = v109
		v62 = v110
		v63 = v111
		v66 = v112
		goto L11
	} else {
		goto L24
	}
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v44))) = uint8(v71)
	v78 = int32(1)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v108 = v88
	v109 = v61
	v110 = v62 + v78
	v111 = v63 + (v61^int32(-1))&v78
	v112 = v66 + v61&v78
	goto L13
L15:
	;
	goto L16
L16:
	;
	v90 = v60 & int32(255)
	if v90 != int32(95) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if v90 != int32(46) {
		v118 = v56
		v124 = v62
		v125 = v63
		v128 = v66
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if base.Ui32(int32(9)) < base.Ui32((v101-int32(48))&int32(255)) {
		goto L5
	} else {
		goto L23
	}
L20:
	;
	if v61&int32(1) != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v97 == int32(95) {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v108 = v97
	v109 = int32(1)
	v110 = v62
	v111 = v63
	v112 = v66
	goto L13
L23:
	;
	v108 = v101
	v109 = v61
	v110 = v62
	v111 = v63
	v112 = v66
	goto L13
L24:
	;
	goto L12
L25:
	;
	v475 = int32(0)
	v476 = F_errsave_start(m, l4)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L7
	} else {
		goto L85
	}
L26:
	;
	v237 = int32(4)
	v240 = base.I32_div_s(v230+v237, v237)
	if int32(0) <= v230 {
		goto L51
	} else {
		goto L52
	}
L27:
	;
	v228 = v118
	v230 = v125
	v233 = v128
	goto L26
L28:
	;
	goto L29
L29:
	;
	v144 = int32(0)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	switch v145 - int32(43) {
	case 0:
		goto L32
	default:
		v153 = v144
		v154 = v118 + int32(1)
		goto L30
	case 2:
		goto L31
	}
L30:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if base.Ui32(int32(9)) < base.Ui32((v155-int32(48))&int32(255)) {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	v153 = int32(1)
	v154 = v118 + int32(2)
	goto L30
L32:
	;
	v153 = v144
	v154 = v118 + int32(2)
	goto L30
L33:
	;
	v163 = v155
	v168 = v154
	v176 = int64(0)
	goto L34
L34:
	;
	if base.Ui32((v163-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	if v153 != 0 {
		goto L45
	} else {
		goto L46
	}
L36:
	;
	goto L35
L37:
	;
	v206 = v168 + int32(1)
	if v203&int32(255) != 0 {
		v163 = v203
		v168 = v206
		v176 = v204
		goto L34
	} else {
		goto L44
	}
L38:
	;
	v188 = v176*int64(10) + base.I64_extend_i32_u(v163)&int64(15)
	if int64(1073741823) < v188 {
		goto L25
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v163&int32(255) != int32(95) {
		v210 = v168
		v211 = v176
		goto L36
	} else {
		goto L42
	}
L41:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	v203 = v191
	v204 = v188
	goto L37
L42:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	if base.Ui32(int32(9)) < base.Ui32((v196-int32(48))&int32(255)) {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v203 = v196
	v204 = v176
	goto L37
L44:
	;
	v210 = v206
	v211 = v204
	goto L36
L45:
	;
	v214 = int64(0) - v211
	goto L47
L46:
	;
	v214 = v211
	goto L47
L47:
	;
	v215 = base.I32_wrap_i64(v214)
	v216 = v128 - v215
	v217 = int32(0)
	if v217 < v216 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v220 = v216
	goto L50
L49:
	;
	v220 = v217
	goto L50
L50:
	;
	v228 = v210
	v230 = v215 + v125
	v233 = v220
	goto L26
L51:
	;
	v247 = v240 - int32(1)
	goto L53
L52:
	;
	v247 = v230 >> (uint(int32(2)) % 32)
	goto L53
L53:
	;
	v248 = int32(2)
	v250 = v247<<(uint(v248)%32) - v230
	v251 = v124 + v250
	v253 = v251 + v248
	v255 = base.I32_div_s(v253, int32(4))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v256 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_pfree(m, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L7
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v263 = F_palloc(m, v255<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L7
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v263
	v266 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v263))) = uint16(v266)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v255
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v233
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v247
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v274 = v272 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v274
	if v253 < int32(4) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_pfree(m, v44)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L7
	} else {
		goto L68
	}
L60:
	;
	v278 = int32(1)
	v279 = v278 - v250
	if v255&v278 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v282 = v279 + v44
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	v284 = int32(10)
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+2)))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+3)))
	v295 = ((v283*v284+v286)*v284+v290)*v284 + v294
	*(*uint16)(unsafe.Add(mBase, uint32(v272)+2)) = uint16(v295)
	v304 = int32(5) - v250
	v305 = v272 + int32(4)
	v306 = v255 - int32(1)
	goto L63
L62:
	;
	v304 = v279
	v305 = v274
	v306 = v255
	goto L63
L63:
	;
	if base.Ui32(v251-int32(2)) < base.Ui32(int32(4)) {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v312 = v304
	v316 = v305
	v318 = v306
	goto L65
L65:
	;
	v326 = v312 + v44
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	v328 = int32(10)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+1)))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+2)))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+3)))
	v339 = ((v327*v328+v330)*v328+v334)*v328 + v338
	*(*uint16)(unsafe.Add(mBase, uint32(v316))) = uint16(v339)
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+7)))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+6)))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+5)))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326)+4)))
	v353 = v341 + (v342+(v343+v344*v328)*v328)*v328
	*(*uint16)(unsafe.Add(mBase, uint32(v316)+2)) = uint16(v353)
	v359 = int32(2)
	if v359 < v318 {
		v312 = v312 + int32(8)
		v316 = v316 + int32(4)
		v318 = v318 - v359
		goto L65
	} else {
		goto L67
	}
L66:
	;
	goto L59
L67:
	;
	goto L66
L68:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v381 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v457
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v228
	v527 = int32(1)
	goto L4
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v457 = int32(0)
	v461 = v443
	goto L69
L71:
	;
	v388 = v381
	v392 = v380
	goto L74
L72:
	;
	goto L73
L73:
	;
	if v381 != 0 {
		v457 = v381
		v461 = v380
		goto L69
	} else {
		goto L84
	}
L74:
	;
	v402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392))))
	if v402 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v443 = v380 + v381<<(uint(int32(1))%32)
	goto L70
L76:
	;
	v404 = v388
	goto L79
L77:
	;
	goto L78
L78:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v429 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v428 - v429
	if v429 < v388 {
		v388 = v388 - v429
		v392 = v392 + int32(2)
		goto L74
	} else {
		goto L83
	}
L79:
	;
	v423 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392+v404<<(uint(int32(1))%32)-int32(2)))))
	if v423 != 0 {
		v457 = v404
		v461 = v392
		goto L69
	} else {
		goto L81
	}
L81:
	;
	v424 = int32(1)
	if v424 < v404 {
		v404 = v404 - v424
		goto L79
	} else {
		goto L82
	}
L82:
	;
	v443 = v392
	goto L70
L83:
	;
	goto L75
L84:
	;
	v443 = v380
	goto L70
L85:
	;
	if v476 == int32(0) {
		v527 = v475
		goto L4
	} else {
		goto L86
	}
L86:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	F_errmsg(m, int32(_a_F_set_var_from_str_1), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	F_errsave_finish(m, l4, int32(_a_F_set_var_from_str_2), int32(_a_F_set_var_from_str_3), int32(_a_F_set_var_from_str_4))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	v527 = v475
	goto L4
L90:
	;
	if v508 == int32(0) {
		v527 = v507
		goto L4
	} else {
		goto L91
	}
L91:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(_a_F_set_var_from_str_5)
	F_errmsg(m, int32(_a_F_set_var_from_str_6), v18)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L7
	} else {
		goto L93
	}
L93:
	;
	F_errsave_finish(m, l4, int32(_a_F_set_var_from_str_2), int32(_a_F_set_var_from_str_7), int32(_a_F_set_var_from_str_4))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	v527 = v507
	goto L4
}
