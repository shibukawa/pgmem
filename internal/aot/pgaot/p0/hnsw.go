package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_HnswCheckNorm(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = F_FunctionCall1Coll(m, v3, v4, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
		return base.F64_gt(v9, float64(0))
	}
}
func F_HnswFindElementNeighbors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 float64
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 float64
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int64
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+65)))
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v34
	if l3 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v34 = v25
	goto L1
L3:
	;
	goto L4
L4:
	;
	v26 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v27 == v26 {
		v34 = v26
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = l0 + v27 - int32(1)
	goto L1
L6:
	;
	m.G0 = v20 + int32(32)
	return
L7:
	;
	if l7 != 0 {
		goto L21
	} else {
		goto L22
	}
L8:
	;
	v38 = l1 - l0
	v39 = int32(16)
	v43 = (int32(base.Ui32(v38)>>(uint(v39)%32)) ^ v38) * int32(-2048144789)
	v48 = (int32(base.Ui32(v43)>>(uint(int32(13))%32)) ^ v43) * int32(-1028477387)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = int32(base.Ui32(v48)>>(uint(v39)%32)) ^ v48
	if l2 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if l2 == int32(0) {
		goto L6
	} else {
		goto L19
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if l0 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v69 = F_FunctionCall2Coll(m, v55, v56, v34, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+88))
	v68 = v59
	goto L12
L14:
	;
	goto L15
L15:
	;
	v60 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+88))
	if v61 == v60 {
		v68 = v60
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v68 = l0 + v61 - int32(1)
	goto L12
L17:
	;
	return
L18:
	;
	v88 = v69
	goto L7
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = l2
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+80)))
	v77 = v20 + int32(16)
	F_HnswLoadElementImpl(m, v74, v75, v77, v20+int32(12), l3, l4, int32(1), int32(0), v20+int32(28))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v88 = v77
	goto L7
L21:
	;
	v90 = l1
	goto L23
L22:
	;
	v90 = int32(0)
	goto L23
L23:
	;
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v88)))
	v93 = F_palloc(m, int32(40))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v93)+32)) = v91
	if l0 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v99 = l2 - l0 + int32(1)
	goto L27
L26:
	;
	v99 = l2
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+24)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v93
	v106 = F_list_make1_impl(m, int32(1), v20+int32(4))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+65)))
	if base.Ui32(v22) < base.Ui32(v108) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v112 = v108
	v120 = v106
	goto L32
L30:
	;
	v150 = v106
	goto L31
L31:
	;
	if base.Ui32(v22) < base.Ui32(v108) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v129 = int32(1)
	v131 = int32(0)
	v135 = F_HnswSearchLayer(m, l0, v20+int32(12), v120, v129, v112, l3, l4, l5, v129, v90, v131, v131, v129, v131)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L17
	} else {
		goto L34
	}
L33:
	;
	v150 = v135
	goto L31
L34:
	;
	v138 = v112 - int32(1)
	if base.Ui32(v22) < base.Ui32(v138) {
		v112 = v138
		v120 = v135
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v158 = v22
	goto L38
L37:
	;
	v158 = v108
	goto L38
L38:
	;
	v162 = v158
	v170 = v150
	goto L39
L39:
	;
	v179 = int32(1)
	v180 = int32(0)
	v184 = F_HnswSearchLayer(m, l0, v20+int32(12), v170, l6+l7, v162, l3, l4, l5, v179, v90, v180, v180, v179, v180)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L17
	} else {
		goto L43
	}
L40:
	;
	goto L6
L41:
	;
	v334 = l5 << (uint(base.B2i32(v162 == int32(0))) % 32)
	if l0 != 0 {
		goto L75
	} else {
		goto L76
	}
L42:
	;
	v251 = int32(0)
	v255 = base.AtomicRmwOr32(m, v251, int32(_a_F_HnswFindElementNeighbors_0), v251)
	if v234 == v251 {
		v321 = v251
		goto L41
	} else {
		goto L57
	}
L43:
	;
	if v184 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v186 = int32(0)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v186 < v188 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v244 = int32(0)
	if l3 == v244 {
		v321 = v244
		goto L41
	} else {
		goto L56
	}
L47:
	;
	v193 = v186
	v198 = v186
	goto L50
L48:
	;
	v234 = v186
	goto L49
L49:
	;
	if l3 != 0 {
		goto L42
	} else {
		goto L55
	}
L50:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208+v193<<(uint(int32(2))%32))))
	v214 = F_palloc(m, int32(12))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L17
	} else {
		goto L52
	}
L51:
	;
	v234 = v221
	goto L49
L52:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v216
	v218 = *(*float64)(unsafe.Add(mBase, uint32(v212)+32))
	*(*float32)(unsafe.Add(mBase, uint32(v214)+4)) = base.F32_demote_f64(v218)
	v221 = F_lappend(m, v198, v214)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L17
	} else {
		goto L53
	}
L53:
	;
	v224 = v193 + int32(1)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v224 < v225 {
		v193 = v224
		v198 = v221
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	v321 = v234
	goto L41
L56:
	;
	v247 = int32(0)
	v250 = base.AtomicRmwOr32(m, v247, int32(_a_F_HnswFindElementNeighbors_0), v247)
	v321 = v244
	goto L41
L57:
	;
	v258 = int32(0)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v259 <= v258 {
		v321 = v251
		goto L41
	} else {
		goto L58
	}
L58:
	;
	v264 = v258
	v268 = v251
	v275 = v259
	goto L59
L59:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279+v264<<(uint(int32(2))%32))))
	if l0 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v321 = v310
	goto L41
L61:
	;
	if v90 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L62:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	v295 = v286
	goto L61
L63:
	;
	goto L64
L64:
	;
	v287 = int32(0)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v288 == v287 {
		v295 = v287
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v295 = l0 + v288 - int32(1)
	goto L61
L66:
	;
	v313 = v264 + int32(1)
	if v313 < v311 {
		v264 = v313
		v268 = v310
		v275 = v311
		goto L59
	} else {
		goto L73
	}
L67:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+64)))
	if v304 == int32(0) {
		v310 = v268
		v311 = v275
		goto L66
	} else {
		goto L71
	}
L68:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v295)+76))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v298 != v299 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295)+80)))
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+80)))
	if v301 == v302 {
		v310 = v268
		v311 = v275
		goto L66
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	v307 = F_lappend(m, v268, v283)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L17
	} else {
		goto L72
	}
L72:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v310 = v307
	v311 = v309
	goto L66
L73:
	;
	goto L60
L74:
	;
	if v380 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L75:
	;
	v335 = int32(0)
	v337 = v162 << (uint(int32(2)) % 32)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v337+(l0+v338)-int32(1))))
	v350 = F_SelectNeighbors(m, l0, v321, v334, l4, l0+v343+int32(3), v335, v335, v335)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L17
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v363 = int32(0)
	v365 = v162 << (uint(int32(2)) % 32)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v365+v366)))
	v374 = F_SelectNeighbors(m, v363, v321, v334, l4, v368+int32(4), v363, v363, v363)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L17
	} else {
		goto L80
	}
L78:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0+v352+v337-int32(1))))
	if v357 == int32(0) {
		v380 = v350
		v381 = v335
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v380 = v350
	v381 = l0 + v357 - int32(1)
	goto L74
L80:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v376+v365)))
	v380 = v374
	v381 = v378
	goto L74
L81:
	;
	if int32(0) < v162 {
		v162 = v162 - int32(1)
		v170 = v184
		goto L39
	} else {
		goto L87
	}
L82:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	if v384 <= int32(0) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v392 = int32(0)
	goto L84
L84:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v380)+12))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	v409 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v381))) = v408 + v409
	v414 = v381 + int32(8) + v408*int32(12)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v407+v392<<(uint(int32(2))%32))))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v414)+8)) = v419
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v418)))
	*(*int64)(unsafe.Add(mBase, uint32(v414))) = v421
	v424 = v392 + v409
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	if v424 < v425 {
		v392 = v424
		goto L84
	} else {
		goto L86
	}
L85:
	;
	goto L81
L86:
	;
	goto L85
L87:
	;
	goto L40
}
func F_HnswGetMetaPageInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v8 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		F_LockBuffer(m, v8, int32(1))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			if v8 < int32(0) {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_HnswGetMetaPageInfo[0]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+(v8^int32(-1))<<(uint(int32(2))%32))))
				v30 = v22
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_HnswGetMetaPageInfo[1]))
				v30 = v24 + v8<<(uint(int32(13))%32) + int32(-8192)
			}
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
			if v31 == int32(-1454134957) {
				if l1 != 0 {
					v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+36)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v34
				} else {
				}
				if l2 != 0 {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+40))
					if v36 != int32(-1) {
						v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+44)))
						v41 = F_palloc(m, int32(108))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(v41)+80)) = uint16(v39)
							*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v36
							v45 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+46)))
							*(*uint8)(unsafe.Add(mBase, uint32(v41)+65)) = uint8(v50)
							F_UnlockReleaseBuffer(m, v8)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
						F_UnlockReleaseBuffer(m, v8)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_UnlockReleaseBuffer(m, v8)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_HnswGetMetaPageInfo_0), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_HnswGetMetaPageInfo_1), int32(311), int32(_a_F_HnswGetMetaPageInfo_2))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
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
func F_HnswInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v101 float64
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_HnswInit[0])))
	if v8 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[1]))
		v16 = F_LWLockAcquire(m, v12+int32(2688), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v22 = F_ShmemInitStruct(m, int32(_a_F_HnswInit_0), int32(4), v5+int32(15))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+15)))
				if v24 == int32(1) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					v31 = v27
					*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[2])) = v31
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[1]))
					F_LWLockRelease(m, v35+int32(2688))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[2]))
						F_LWLockRegisterTranche(m, v41, int32(_a_F_HnswInit_1))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							v47 = F_add_reloption_kind(m)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3])) = v47
								F_add_int_reloption(m, v47, int32(_a_F_HnswInit_2), int32(_a_F_HnswInit_3), int32(16), int32(2), int32(100))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3]))
									F_add_int_reloption(m, v58, int32(_a_F_HnswInit_4), int32(_a_F_HnswInit_5), int32(64), int32(4), int32(1000))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_6), int32(_a_F_HnswInit_7), int32(_a_F_HnswInit_8), int32(_a_F_HnswInit_9), int32(40), int32(1), int32(1000), int32(6), int32(0))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return
										} else {
											v79 = int32(0)
											F_DefineCustomEnumVariable(m, int32(_a_F_HnswInit_10), int32(_a_F_HnswInit_11), v79, int32(_a_F_HnswInit_12), v79, int32(_a_F_HnswInit_13), int32(6))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return
											} else {
												v88 = int32(0)
												F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_14), int32(_a_F_HnswInit_15), v88, int32(_a_F_HnswInit_16), int32(_a_F_HnswInit_17), int32(1), int32(2147483647), int32(6), v88)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return
												} else {
													v101 = float64(1)
													F_DefineCustomRealVariable(m, int32(_a_F_HnswInit_18), int32(_a_F_HnswInit_19), int32(0), int32(_a_F_HnswInit_20), v101, v101, float64(1000), int32(6))
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return
													} else {
														F_MarkGUCPrefixReserved(m, int32(_a_F_HnswInit_21))
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return
														} else {
															m.G0 = v5 + int32(16)
															return
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
				} else {
					v28 = F_LWLockNewTrancheId(m)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22))) = v28
						v31 = v28
						*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[2])) = v31
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[1]))
						F_LWLockRelease(m, v35+int32(2688))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[2]))
							F_LWLockRegisterTranche(m, v41, int32(_a_F_HnswInit_1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								v47 = F_add_reloption_kind(m)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3])) = v47
									F_add_int_reloption(m, v47, int32(_a_F_HnswInit_2), int32(_a_F_HnswInit_3), int32(16), int32(2), int32(100))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3]))
										F_add_int_reloption(m, v58, int32(_a_F_HnswInit_4), int32(_a_F_HnswInit_5), int32(64), int32(4), int32(1000))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_6), int32(_a_F_HnswInit_7), int32(_a_F_HnswInit_8), int32(_a_F_HnswInit_9), int32(40), int32(1), int32(1000), int32(6), int32(0))
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												v79 = int32(0)
												F_DefineCustomEnumVariable(m, int32(_a_F_HnswInit_10), int32(_a_F_HnswInit_11), v79, int32(_a_F_HnswInit_12), v79, int32(_a_F_HnswInit_13), int32(6))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return
												} else {
													v88 = int32(0)
													F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_14), int32(_a_F_HnswInit_15), v88, int32(_a_F_HnswInit_16), int32(_a_F_HnswInit_17), int32(1), int32(2147483647), int32(6), v88)
													mBase = m.M
													v96 = m.ExcPending
													if v96 != 0 {
														return
													} else {
														v101 = float64(1)
														F_DefineCustomRealVariable(m, int32(_a_F_HnswInit_18), int32(_a_F_HnswInit_19), int32(0), int32(_a_F_HnswInit_20), v101, v101, float64(1000), int32(6))
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return
														} else {
															F_MarkGUCPrefixReserved(m, int32(_a_F_HnswInit_21))
															mBase = m.M
															v109 = m.ExcPending
															if v109 != 0 {
																return
															} else {
																m.G0 = v5 + int32(16)
																return
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
	} else {
		v47 = F_add_reloption_kind(m)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3])) = v47
			F_add_int_reloption(m, v47, int32(_a_F_HnswInit_2), int32(_a_F_HnswInit_3), int32(16), int32(2), int32(100))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3]))
				F_add_int_reloption(m, v58, int32(_a_F_HnswInit_4), int32(_a_F_HnswInit_5), int32(64), int32(4), int32(1000))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_6), int32(_a_F_HnswInit_7), int32(_a_F_HnswInit_8), int32(_a_F_HnswInit_9), int32(40), int32(1), int32(1000), int32(6), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						v79 = int32(0)
						F_DefineCustomEnumVariable(m, int32(_a_F_HnswInit_10), int32(_a_F_HnswInit_11), v79, int32(_a_F_HnswInit_12), v79, int32(_a_F_HnswInit_13), int32(6))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							v88 = int32(0)
							F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_14), int32(_a_F_HnswInit_15), v88, int32(_a_F_HnswInit_16), int32(_a_F_HnswInit_17), int32(1), int32(2147483647), int32(6), v88)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								v101 = float64(1)
								F_DefineCustomRealVariable(m, int32(_a_F_HnswInit_18), int32(_a_F_HnswInit_19), int32(0), int32(_a_F_HnswInit_20), v101, v101, float64(1000), int32(6))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									F_MarkGUCPrefixReserved(m, int32(_a_F_HnswInit_21))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										m.G0 = v5 + int32(16)
										return
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
func F_HnswInitLockTranche(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[0]))
	v12 = F_LWLockAcquire(m, v8+int32(2688), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v18 = F_ShmemInitStruct(m, int32(_a_F_HnswInitLockTranche_0), int32(4), v5+int32(15))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+15)))
			if v20 == int32(1) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v27 = v23
				*(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[1])) = v27
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[0]))
				F_LWLockRelease(m, v31+int32(2688))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[1]))
					F_LWLockRegisterTranche(m, v37, int32(_a_F_HnswInitLockTranche_1))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				}
			} else {
				v24 = F_LWLockNewTrancheId(m)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v24
					v27 = v24
					*(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[1])) = v27
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[0]))
					F_LWLockRelease(m, v31+int32(2688))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[1]))
						F_LWLockRegisterTranche(m, v37, int32(_a_F_HnswInitLockTranche_1))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							m.G0 = v5 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_HnswLoadElementImpl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 float64
	_ = v54
	var v56 float64
	_ = v56
	var v60 float64
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v2 = l1
	v12 = F_ReadBuffer(m, l4, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		F_LockBuffer(m, v12, int32(1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v12 < int32(0) {
				v20 = *(*int32)(unsafe.Add(mBase, _c_F_HnswLoadElementImpl[0]))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v20+(v12^int32(-1))<<(uint(int32(2))%32))))
				v34 = v26
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_HnswLoadElementImpl[1]))
				v34 = v28 + v12<<(uint(int32(13))%32) + int32(-8192)
			}
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v2<<(uint(int32(2))%32))+20))
			v41 = v38&int32(_a_F_HnswLoadElementImpl_0) + v34
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)))
			if v42 == int32(0) {
				if l2 == int32(0) {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
					if v66 == int32(0) {
						v70 = F_palloc(m, int32(108))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(v70)+80)) = uint16(v2)
							*(*int32)(unsafe.Add(mBase, uint32(v70)+76)) = l0
							v74 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v70)+88)) = v74
							*(*int32)(unsafe.Add(mBase, uint32(v70)+72)) = v74
							*(*int32)(unsafe.Add(mBase, uint32(l8))) = v70
							v79 = v70
							F_HnswLoadElementFromTuple(m, v79, v41, int32(1), l6)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v79 = v66
						F_HnswLoadElementFromTuple(m, v79, v41, int32(1), l6)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							F_UnlockReleaseBuffer(m, v12)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					if v47 != 0 {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
						v52 = F_FunctionCall2Coll(m, v48, v49, v47, v41+int32(72))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v54 = *(*float64)(unsafe.Add(mBase, uint32(v52)))
							v56 = v54
							*(*float64)(unsafe.Add(mBase, uint32(l2))) = v56
							if l7 == int32(0) {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
								if v66 == int32(0) {
									v70 = F_palloc(m, int32(108))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										*(*uint16)(unsafe.Add(mBase, uint32(v70)+80)) = uint16(v2)
										*(*int32)(unsafe.Add(mBase, uint32(v70)+76)) = l0
										v74 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v70)+88)) = v74
										*(*int32)(unsafe.Add(mBase, uint32(v70)+72)) = v74
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v70
										v79 = v70
										F_HnswLoadElementFromTuple(m, v79, v41, int32(1), l6)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											F_UnlockReleaseBuffer(m, v12)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												return
											}
										}
									}
								} else {
									v79 = v66
									F_HnswLoadElementFromTuple(m, v79, v41, int32(1), l6)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										F_UnlockReleaseBuffer(m, v12)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								v60 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
								if base.F64_lt(v56, v60) == int32(0) {
									F_UnlockReleaseBuffer(m, v12)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										return
									}
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
									if v66 == int32(0) {
										v70 = F_palloc(m, int32(108))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											*(*uint16)(unsafe.Add(mBase, uint32(v70)+80)) = uint16(v2)
											*(*int32)(unsafe.Add(mBase, uint32(v70)+76)) = l0
											v74 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v70)+88)) = v74
											*(*int32)(unsafe.Add(mBase, uint32(v70)+72)) = v74
											*(*int32)(unsafe.Add(mBase, uint32(l8))) = v70
											v79 = v70
											F_HnswLoadElementFromTuple(m, v79, v41, int32(1), l6)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return
											} else {
												F_UnlockReleaseBuffer(m, v12)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													return
												}
											}
										}
									} else {
										v79 = v66
										F_HnswLoadElementFromTuple(m, v79, v41, int32(1), l6)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											F_UnlockReleaseBuffer(m, v12)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						}
					} else {
						v56 = float64(0)
						*(*float64)(unsafe.Add(mBase, uint32(l2))) = v56
						if l7 == int32(0) {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
							if v66 == int32(0) {
								v70 = F_palloc(m, int32(108))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									*(*uint16)(unsafe.Add(mBase, uint32(v70)+80)) = uint16(v2)
									*(*int32)(unsafe.Add(mBase, uint32(v70)+76)) = l0
									v74 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v70)+88)) = v74
									*(*int32)(unsafe.Add(mBase, uint32(v70)+72)) = v74
									*(*int32)(unsafe.Add(mBase, uint32(l8))) = v70
									v79 = v70
									F_HnswLoadElementFromTuple(m, v79, v41, int32(1), l6)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										F_UnlockReleaseBuffer(m, v12)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								v79 = v66
								F_HnswLoadElementFromTuple(m, v79, v41, int32(1), l6)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return
								} else {
									F_UnlockReleaseBuffer(m, v12)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							v60 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
							if base.F64_lt(v56, v60) == int32(0) {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return
								} else {
									return
								}
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
								if v66 == int32(0) {
									v70 = F_palloc(m, int32(108))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										*(*uint16)(unsafe.Add(mBase, uint32(v70)+80)) = uint16(v2)
										*(*int32)(unsafe.Add(mBase, uint32(v70)+76)) = l0
										v74 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v70)+88)) = v74
										*(*int32)(unsafe.Add(mBase, uint32(v70)+72)) = v74
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v70
										v79 = v70
										F_HnswLoadElementFromTuple(m, v79, v41, int32(1), l6)
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											F_UnlockReleaseBuffer(m, v12)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												return
											}
										}
									}
								} else {
									v79 = v66
									F_HnswLoadElementFromTuple(m, v79, v41, int32(1), l6)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										F_UnlockReleaseBuffer(m, v12)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_HnswLoadElementImpl_1), int32(0))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_HnswLoadElementImpl_2), int32(550), int32(_a_F_HnswLoadElementImpl_3))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
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
func F_HnswParallelBuildMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v10 = F_shm_toc_lookup(m, l1, int64(-6917529027641081853), int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_HnswParallelBuildMain[0])) = v10
		F_pgstat_report_activity(m, int32(3), v10)
		mBase = m.M
		v17 = F_shm_toc_lookup(m, l1, int64(-6917529027641081855), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+8)))
			if v22 != 0 {
				v23 = int32(4)
			} else {
				v23 = int32(5)
			}
			v24 = F_table_open(m, v19, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				if v22 != 0 {
					v29 = int32(3)
				} else {
					v29 = int32(8)
				}
				v30 = F_index_open(m, v26, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v34 = F_shm_toc_lookup(m, l1, int64(-6917529027641081854), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_HnswParallelScanAndInsert(m, v24, v30, v17, v34, int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							F_relation_close(m, v30, v29)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								F_relation_close(m, v24, v23)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_HnswParallelScanAndInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 float64
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 float64
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v10 = m.G0
	v12 = v10 - int32(224)
	m.G0 = v12
	v14 = F_BuildIndexInfo(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+121)) = uint8(v16)
		v19 = v12 + int32(16)
		F_InitBuildState_1(m, v19, l0, l1, v14, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+220)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v12)+176)) = l2 + int32(40)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = int32(_a_F_HnswParallelScanAndInsert_0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+208)) = v19
			v31 = int32(0)
			v37 = F_table_beginscan_parallel(m, l0, l2+int32(160))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+140))
				v41 = m.T0[v40].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, v14, int32(1), v31, l4, v31, int32(-1), int32(_a_F_HnswParallelScanAndInsert_1), v19, v37)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v45 = base.AtomicRmwXchg32(m, l2, int32(24), int32(1))
					if v45 != 0 {
						F_s_lock(m, l2+int32(24), int32(_a_F_HnswParallelScanAndInsert_2), int32(818), int32(_a_F_HnswParallelScanAndInsert_3))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v53 + int32(1)
							v57 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
							*(*float64)(unsafe.Add(mBase, uint32(l2)+32)) = base.F64_add(v57, v41)
							v60 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+24)), uint32(v60))
							v65 = F_errstart(m, int32(14), v60)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								if v65 != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(v12))) = base.I64_trunc_sat_f64_s(v41)
									if l4 != 0 {
										v71 = int32(_a_F_HnswParallelScanAndInsert_4)
									} else {
										v71 = int32(_a_F_HnswParallelScanAndInsert_5)
									}
									F_errmsg(m, v71, v12)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										if l4 != 0 {
											v77 = int32(825)
										} else {
											v77 = int32(827)
										}
										F_errfinish(m, int32(_a_F_HnswParallelScanAndInsert_2), v77, int32(_a_F_HnswParallelScanAndInsert_3))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											F_ConditionVariableSignal(m, l2+int32(12))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												v85 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
												F_MemoryContextDelete(m, v85)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
													F_MemoryContextDelete(m, v88)
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return
													} else {
														m.G0 = v12 + int32(224)
														return
													}
												}
											}
										}
									}
								} else {
									F_ConditionVariableSignal(m, l2+int32(12))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return
									} else {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
										F_MemoryContextDelete(m, v85)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
											F_MemoryContextDelete(m, v88)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												m.G0 = v12 + int32(224)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v53 + int32(1)
						v57 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
						*(*float64)(unsafe.Add(mBase, uint32(l2)+32)) = base.F64_add(v57, v41)
						v60 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+24)), uint32(v60))
						v65 = F_errstart(m, int32(14), v60)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							if v65 != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(v12))) = base.I64_trunc_sat_f64_s(v41)
								if l4 != 0 {
									v71 = int32(_a_F_HnswParallelScanAndInsert_4)
								} else {
									v71 = int32(_a_F_HnswParallelScanAndInsert_5)
								}
								F_errmsg(m, v71, v12)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									if l4 != 0 {
										v77 = int32(825)
									} else {
										v77 = int32(827)
									}
									F_errfinish(m, int32(_a_F_HnswParallelScanAndInsert_2), v77, int32(_a_F_HnswParallelScanAndInsert_3))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										F_ConditionVariableSignal(m, l2+int32(12))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return
										} else {
											v85 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
											F_MemoryContextDelete(m, v85)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
												F_MemoryContextDelete(m, v88)
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return
												} else {
													m.G0 = v12 + int32(224)
													return
												}
											}
										}
									}
								}
							} else {
								F_ConditionVariableSignal(m, l2+int32(12))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
									F_MemoryContextDelete(m, v85)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
										F_MemoryContextDelete(m, v88)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											m.G0 = v12 + int32(224)
											return
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
func F_hnsw_bit_support(m *base.Module, l0 int32) int32 {
	return int32(_a_F_hnsw_bit_support_0)
}
