package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateDataDirLockFile(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v6 = *(*int32)(unsafe.Add(mBase, _consts[200]))
	F_CreateLockFile(m, int32(407080), l0, int32(706478), int32(1), v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_dataExecPlaceToPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int64
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int64
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	if l1 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return
L2:
	;
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	v145 = v141 + v142*int32(10)
	v147 = v145 + int32(22)
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = base.I32_rotr(l4, int32(16))
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139+v141)+4)))
	if v142 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L3:
	;
	v46 = v43 & int32(128)
	if v46 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v17 = int32(2)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+(l1^int32(-1))<<(uint(v17)%32))))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v21)+6)))
	if v23&v17 != 0 {
		v42 = v21
		v43 = v23
		v44 = v20
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v30 = v27 + l1<<(uint(int32(13))%32)
	v32 = v30 + int32(-8192)
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30-int32(8176)))))
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v35)+6)))
	if v37&int32(2) == int32(0) {
		v139 = v35
		v141 = v32
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v139 = v21
	v141 = v20
	goto L2
L8:
	;
	v42 = v35
	v43 = v37
	v44 = v32
	goto L3
L9:
	;
	v51 = v43 | int32(128)
	*(*uint16)(unsafe.Add(mBase, uint32(v42+v44)+6)) = uint16(v51)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+16)))
	v55 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v44+v53)+4)) = uint16(v55)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v57 = int32(32)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v58 == int32(0) {
		v108 = v57
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)) = uint16(v108)
	F_MarkBufferDirty(m, l1)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L28
	} else {
		goto L29
	}
L13:
	;
	if l5 == v58 {
		v108 = v57
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v62 = int32(0)
	v69 = v44 + int32(32)
	v70 = base.B2i32(v46 == v62)
	v71 = v62
	v73 = v58
	goto L15
L15:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+8)))
	v80 = base.B2i32(v77 != int32(0)) | v70
	if v77 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v108 = v98 + int32(32)
	goto L12
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+6)))
	v85 = int32(1)
	v90 = (v84+v85)&int32(131070) + int32(8)
	if v80&v85 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v97 = v69
	v98 = v71
	goto L19
L19:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v101 != l5 {
		v69 = v97
		v70 = v80
		v71 = v98
		v73 = v101
		goto L15
	} else {
		goto L27
	}
L20:
	;
	if v90 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	goto L22
L22:
	;
	v97 = v69 + v90
	v98 = v71 + v90
	goto L19
L23:
	;
	goto L22
L24:
	;
	v93 = F__emscripten_memcpy_bulkmem(m, v69, v83, v90)
	mBase = m.M
	goto L26
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L16
L28:
	;
	return
L29:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+48))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+118)))
	if v120 != int32(112) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v124 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v118)+32))
	if v127 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v129 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118)+40))
	if v128 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L28
	} else {
		goto L37
	}
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	F_XLogRegisterBufData(m, int32(0), v135, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L28
	} else {
		goto L38
	}
L38:
	;
	return
L39:
	;
	v317 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v315))) = v317
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v315)+8)) = uint16(v319)
	v323 = v152 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v316+v141)+4)) = uint16(v323)
	v328 = v152*int32(10) + int32(42)
	*(*uint16)(unsafe.Add(mBase, uint32(v141)+12)) = uint16(v328)
	F_MarkBufferDirty(m, l1)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L28
	} else {
		goto L90
	}
L40:
	;
	v315 = v141 + v152*int32(10) + int32(32)
	v316 = v139
	goto L39
L41:
	;
	goto L42
L42:
	;
	if v152+int32(1) == v142 {
		v315 = v147
		v316 = v139
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v163 = int32(10)
	v164 = v145 + int32(32)
	v169 = (v152-v142)*v163 + v163
	if v164 == v147 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+16)))
	v315 = v147
	v316 = v314
	goto L39
L45:
	;
	goto L44
L46:
	;
	v173 = v164 + v169
	if base.Ui32(v147-v173) <= base.Ui32(int32(0)-v169<<(uint(int32(1))%32)) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v180 = F___memcpy(m, v164, v147, v169)
	mBase = m.M
	goto L44
L48:
	;
	goto L49
L49:
	;
	v183 = (v164 ^ v147) & int32(3)
	if base.Ui32(v164) < base.Ui32(v147) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v285 == int32(0) {
		goto L45
	} else {
		goto L86
	}
L51:
	;
	if base.Ui32(v263) <= base.Ui32(int32(3)) {
		v284 = v262
		v285 = v263
		v286 = v264
		goto L50
	} else {
		goto L82
	}
L52:
	;
	if v183 != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v183 != 0 {
		v245 = v169
		goto L65
	} else {
		goto L66
	}
L55:
	;
	v284 = v147
	v285 = v169
	v286 = v164
	goto L50
L56:
	;
	goto L57
L57:
	;
	if v164&int32(3) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v262 = v147
	v263 = v169
	v264 = v164
	goto L51
L59:
	;
	goto L60
L60:
	;
	v190 = v147
	v191 = v169
	v192 = v164
	goto L61
L61:
	;
	if v191 == int32(0) {
		goto L45
	} else {
		goto L63
	}
L62:
	;
	v262 = v199
	v263 = v201
	v264 = v203
	goto L51
L63:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	*(*uint8)(unsafe.Add(mBase, uint32(v192))) = uint8(v196)
	v198 = int32(1)
	v199 = v190 + v198
	v201 = v191 - v198
	v203 = v192 + v198
	if v203&int32(3) != 0 {
		v190 = v199
		v191 = v201
		v192 = v203
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	if v245 == int32(0) {
		goto L45
	} else {
		goto L78
	}
L66:
	;
	if v173&int32(3) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v210 = v169
	goto L70
L68:
	;
	v225 = v169
	goto L69
L69:
	;
	if base.Ui32(v225) <= base.Ui32(int32(3)) {
		v245 = v225
		goto L65
	} else {
		goto L74
	}
L70:
	;
	if v210 == int32(0) {
		goto L45
	} else {
		goto L72
	}
L71:
	;
	v225 = v216
	goto L69
L72:
	;
	v216 = v210 - int32(1)
	v217 = v164 + v216
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v216))))
	*(*uint8)(unsafe.Add(mBase, uint32(v217))) = uint8(v219)
	if v217&int32(3) != 0 {
		v210 = v216
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v232 = v225
	goto L75
L75:
	;
	v236 = v232 - int32(4)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v147+v236)))
	*(*int32)(unsafe.Add(mBase, uint32(v164+v236))) = v239
	if base.Ui32(int32(3)) < base.Ui32(v236) {
		v232 = v236
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v245 = v236
	goto L65
L77:
	;
	goto L76
L78:
	;
	v252 = v245
	goto L79
L79:
	;
	v256 = v252 - int32(1)
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v256))))
	*(*uint8)(unsafe.Add(mBase, uint32(v164+v256))) = uint8(v259)
	if v256 != 0 {
		v252 = v256
		goto L79
	} else {
		goto L81
	}
L80:
	;
	goto L45
L81:
	;
	goto L80
L82:
	;
	v269 = v262
	v270 = v263
	v271 = v264
	goto L83
L83:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	*(*int32)(unsafe.Add(mBase, uint32(v271))) = v273
	v275 = int32(4)
	v276 = v269 + v275
	v278 = v271 + v275
	v280 = v270 - v275
	if base.Ui32(int32(3)) < base.Ui32(v280) {
		v269 = v276
		v270 = v280
		v271 = v278
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v284 = v276
	v285 = v280
	v286 = v278
	goto L50
L85:
	;
	goto L84
L86:
	;
	v291 = v284
	v292 = v285
	v293 = v286
	goto L87
L87:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	*(*uint8)(unsafe.Add(mBase, uint32(v293))) = uint8(v295)
	v297 = int32(1)
	v302 = v292 - v297
	if v302 != 0 {
		v291 = v291 + v297
		v292 = v302
		v293 = v293 + v297
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L45
L89:
	;
	goto L88
L90:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)+48))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+118)))
	if v334 != int32(112) {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v338 <= int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v332)+32))
	if v341 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+53)))
	if v343 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v332)+40))
	if v342 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	*(*uint16)(unsafe.Add(mBase, _consts[15])) = uint16(v142)
	v347 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, _consts[16])) = v347
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+8)))
	*(*uint16)(unsafe.Add(mBase, _consts[17])) = uint16(v350)
	F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L28
	} else {
		goto L98
	}
L98:
	;
	F_XLogRegisterBufData(m, int32(0), int32(4319748), int32(12))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L28
	} else {
		goto L99
	}
L99:
	;
	goto L1
}
