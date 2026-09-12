package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_int_penalty(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 float32
	_ = v27
	var v28 float32
	_ = v28
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = F_inner_int_union(m, v11, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		F_rt__int_size(m, v14, v7+int32(12))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			F_rt__int_size(m, v22, v7+int32(8))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = *(*float32)(unsafe.Add(mBase, uint32(v7)+12))
				v28 = *(*float32)(unsafe.Add(mBase, uint32(v7)+8))
				*(*float32)(unsafe.Add(mBase, uint32(v9))) = base.F32_sub(v27, v28)
				F_pfree(m, v14)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v9
				}
			}
		}
	}
}
func F_g_int_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v56 int32
	_ = v56
	var v68 float32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 float32
	_ = v92
	var v93 float32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 float32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 float32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v129 float32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 float32
	_ = v149
	var v150 float32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 float32
	_ = v155
	var v156 int32
	_ = v156
	var v157 float32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v183 float32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 float32
	_ = v318
	var v319 float32
	_ = v319
	var v321 float32
	_ = v321
	var v322 float32
	_ = v322
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 float32
	_ = v403
	var v404 float32
	_ = v404
	var v407 float32
	_ = v407
	var v408 float32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 float32
	_ = v425
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 float32
	_ = v438
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v31 = (v27 + int32(65534)) & int32(65535)
	v35 = v31<<(uint(int32(1))%32) + int32(4)
	v36 = F_palloc(m, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v36
	v41 = F_palloc(m, v35)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v41
	if base.Ui32(int32(2)) <= base.Ui32(v31) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v47 = v26 + int32(4)
	v48 = int32(1)
	v50 = v48
	v54 = v48
	v55 = v2
	v56 = v2
	v68 = float32(0)
	goto L7
L5:
	;
	v192 = v41
	v193 = v2
	v194 = v2
	goto L6
L6:
	;
	v208 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v208
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v214 = v26 + int32(4)
	v216 = int32(65535)
	v224 = base.B2i32(v193&v216 == v208) | base.B2i32(v194&v216 == v208)
	if v224 != 0 {
		goto L46
	} else {
		goto L47
	}
L7:
	;
	v70 = int32(4)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v47+v54<<(uint(v70)%32))))
	v75 = v54 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v47+v75<<(uint(v70)%32))))
	v80 = F_inner_int_union(m, v73, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v192 = v187
	v193 = v170
	v194 = v171
	goto L6
L9:
	;
	F_rt__int_size(m, v80, v23+int32(20))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v86 = F_inner_int_inter(m, v73, v79)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_rt__int_size(m, v86, v23+int32(16))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v92 = *(*float32)(unsafe.Add(mBase, uint32(v23)+16))
	v93 = *(*float32)(unsafe.Add(mBase, uint32(v23)+20))
	F_pfree(m, v80)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_pfree(m, v86)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v98 = base.F32_sub(v93, v92)
	v102 = (base.F32_gt(v98, v68) | v50) & int32(1)
	if v102 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v103 = v75
	goto L17
L16:
	;
	v103 = v56
	goto L17
L17:
	;
	if v102 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v104 = v98
	goto L20
L19:
	;
	v104 = v68
	goto L20
L20:
	;
	if v102 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v105 = v54
	goto L23
L22:
	;
	v105 = v55
	goto L23
L23:
	;
	v107 = v54 + int32(2)
	if base.Ui32(v107&int32(65535)) <= base.Ui32(v31) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v111 = v107
	v116 = v105
	v117 = v103
	v129 = v104
	goto L27
L25:
	;
	v170 = v105
	v171 = v103
	v183 = v104
	goto L26
L26:
	;
	if v75 != v31 {
		v50 = int32(0)
		v54 = v75
		v55 = v170
		v56 = v171
		v68 = v183
		goto L7
	} else {
		goto L45
	}
L27:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v47+v111&int32(65535)<<(uint(int32(4))%32))))
	v137 = F_inner_int_union(m, v73, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v170 = v159
	v171 = v158
	v183 = v157
	goto L26
L29:
	;
	F_rt__int_size(m, v137, v23+int32(20))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v143 = F_inner_int_inter(m, v73, v136)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_rt__int_size(m, v143, v23+int32(16))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v149 = *(*float32)(unsafe.Add(mBase, uint32(v23)+16))
	v150 = *(*float32)(unsafe.Add(mBase, uint32(v23)+20))
	F_pfree(m, v137)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_pfree(m, v143)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v155 = base.F32_sub(v150, v149)
	v156 = base.F32_gt(v155, v129)
	if v156 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v157 = v155
	goto L37
L36:
	;
	v157 = v129
	goto L37
L37:
	;
	if v156 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v158 = v111
	goto L40
L39:
	;
	v158 = v117
	goto L40
L40:
	;
	if v156 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v159 = v54
	goto L43
L42:
	;
	v159 = v116
	goto L43
L43:
	;
	v161 = v111 + int32(1)
	if base.Ui32(v161&int32(65535)) <= base.Ui32(v31) {
		v111 = v161
		v116 = v159
		v117 = v158
		v129 = v157
		goto L27
	} else {
		goto L44
	}
L44:
	;
	goto L28
L45:
	;
	goto L8
L46:
	;
	v225 = int32(1)
	goto L48
L47:
	;
	v225 = v193
	goto L48
L48:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v214+v225&int32(65535)<<(uint(int32(4))%32))))
	v232 = F_copy_intArrayType(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_rt__int_size(m, v232, v23+int32(12))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if v224 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v239 = int32(2)
	goto L53
L52:
	;
	v239 = v194
	goto L53
L53:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v214+v239&int32(65535)<<(uint(int32(4))%32))))
	v246 = F_copy_intArrayType(m, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_rt__int_size(m, v246, v23+int32(8))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v252 = int32(65535)
	v253 = v27 + v252
	v255 = v253 & v252
	v258 = F_palloc(m, v255<<(uint(int32(3))%32))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if v27&int32(65535) == int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	F_pfree(m, v258)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L96
	}
L58:
	;
	F_pg_qsort(m, v258, v255, int32(8), int32(6813))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v268 = int32(1)
	v270 = v268
	v272 = v268
	goto L62
L61:
	;
	v461 = v192
	v464 = v212
	v465 = v246
	v466 = v232
	goto L57
L62:
	;
	v292 = v258 + v270<<(uint(int32(3))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v292-int32(8)))) = uint16(v272)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v214+v270<<(uint(int32(4))%32))))
	v300 = F_inner_int_union(m, v232, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	F_pg_qsort(m, v258, v255, int32(8), int32(6813))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L71
	}
L64:
	;
	F_rt__int_size(m, v300, v23+int32(28))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_pfree(m, v300)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v308 = F_inner_int_union(m, v246, v299)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_rt__int_size(m, v308, v23+int32(24))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v308)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v318 = *(*float32)(unsafe.Add(mBase, uint32(v23)+28))
	v319 = *(*float32)(unsafe.Add(mBase, uint32(v23)+12))
	v321 = *(*float32)(unsafe.Add(mBase, uint32(v23)+24))
	v322 = *(*float32)(unsafe.Add(mBase, uint32(v23)+8))
	*(*float32)(unsafe.Add(mBase, uint32(v292-int32(4)))) = base.F32_abs(base.F32_sub(base.F32_sub(v318, v319), base.F32_sub(v321, v322)))
	v328 = v272 + int32(1)
	v329 = int32(65535)
	v330 = v328 & v329
	if base.Ui32(v330) <= base.Ui32(v253&v329) {
		v270 = v330
		v272 = v328
		goto L62
	} else {
		goto L70
	}
L70:
	;
	goto L63
L71:
	;
	v338 = int32(1)
	if base.Ui32(v255) <= base.Ui32(v338) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v341 = v338
	goto L74
L73:
	;
	v341 = v255
	goto L74
L74:
	;
	v345 = int32(0)
	v349 = v192
	v352 = v212
	v353 = v246
	v354 = v232
	goto L75
L75:
	;
	v368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258+v345<<(uint(int32(3))%32)))))
	if v225&int32(65535) == v368 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v461 = v447
	v464 = v448
	v465 = v449
	v466 = v450
	goto L57
L77:
	;
	v455 = v345 + int32(1)
	if v455 != v341 {
		v345 = v455
		v349 = v447
		v352 = v448
		v353 = v449
		v354 = v450
		goto L75
	} else {
		goto L95
	}
L78:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v352))) = uint16(v225)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v371 + int32(1)
	v447 = v349
	v448 = v352 + int32(2)
	v449 = v353
	v450 = v354
	goto L77
L79:
	;
	goto L80
L80:
	;
	if v239&int32(65535) == v368 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v349))) = uint16(v239)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v381 + int32(1)
	v447 = v349 + int32(2)
	v448 = v352
	v449 = v353
	v450 = v354
	goto L77
L82:
	;
	goto L83
L83:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v214+v368<<(uint(int32(4))%32))))
	v391 = F_inner_int_union(m, v354, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v393 = F_inner_int_union(m, v353, v390)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_rt__int_size(m, v391, v23+int32(28))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_rt__int_size(m, v393, v23+int32(24))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v403 = *(*float32)(unsafe.Add(mBase, uint32(v23)+28))
	v404 = *(*float32)(unsafe.Add(mBase, uint32(v23)+12))
	v407 = *(*float32)(unsafe.Add(mBase, uint32(v23)+24))
	v408 = *(*float32)(unsafe.Add(mBase, uint32(v23)+8))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v413 = v411 - v412
	if base.F64_lt(base.F64_promote_f32(base.F32_sub(v403, v404)), base.F64_add(base.F64_promote_f32(base.F32_sub(v407, v408)), base.F64_mul(base.F64_convert_i32_s(v413*v413*v413), float64(-0.01)))) != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F_pfree(m, v354)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	F_pfree(m, v353)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	F_pfree(m, v393)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v425 = *(*float32)(unsafe.Add(mBase, uint32(v23)+28))
	*(*float32)(unsafe.Add(mBase, uint32(v23)+12)) = v425
	*(*uint16)(unsafe.Add(mBase, uint32(v352))) = uint16(v368)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v428 + int32(1)
	v447 = v349
	v448 = v352 + int32(2)
	v449 = v353
	v450 = v391
	goto L77
L93:
	;
	F_pfree(m, v391)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v438 = *(*float32)(unsafe.Add(mBase, uint32(v23)+24))
	*(*float32)(unsafe.Add(mBase, uint32(v23)+8)) = v438
	*(*uint16)(unsafe.Add(mBase, uint32(v349))) = uint16(v368)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v441 + int32(1)
	v447 = v349 + int32(2)
	v448 = v352
	v449 = v393
	v450 = v354
	goto L77
L95:
	;
	goto L76
L96:
	;
	v479 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v464))) = uint16(v479)
	*(*uint16)(unsafe.Add(mBase, uint32(v461))) = uint16(v479)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v465
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v466
	m.G0 = v23 + int32(32)
	return v25
}
