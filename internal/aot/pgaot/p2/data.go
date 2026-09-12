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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[203]))
	F_CreateLockFile(m, int32(434705), l0, int32(759461), int32(1), v6)
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
	*(*uint16)(unsafe.Add(mBase, _consts[19])) = uint16(v142)
	v347 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, _consts[20])) = v347
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+8)))
	*(*uint16)(unsafe.Add(mBase, _consts[21])) = uint16(v350)
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
	F_XLogRegisterBufData(m, int32(0), int32(4415028), int32(12))
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
func F_process_data_packets(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
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
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v417 int32
	_ = v417
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v597 int32
	_ = v597
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(1120)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v6
	v24 = v6
	v28 = v6
	goto L1
L1:
	;
	v36 = F_pgp_parse_pkt_hdr(m, l2, v15+int32(47), v15+int32(40), int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	m.G0 = v15 + int32(1120)
	return v597
L3:
	;
	goto L2
L4:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	F_pullf_free(m, v583)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L9
	} else {
		goto L186
	}
L5:
	;
	v576 = int32(0)
	v578 = v564
	v580 = int32(1)
	goto L4
L6:
	;
	v391 = F_pullf_read_max(m, v83, int32(4), v15+int32(1088), v15+int32(1084))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L9
	} else {
		goto L131
	}
L7:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	if v351 != 0 {
		goto L117
	} else {
		goto L118
	}
L8:
	;
	v348 = v344
	v350 = base.B2i32(v24 != int32(0))
	goto L7
L9:
	;
	return int32(0)
L10:
	;
	if v36 <= int32(0) {
		v344 = v36
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v24 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_px_debug(m, int32(493326), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v36 == int32(3) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v348 = int32(-100)
	v350 = int32(1)
	goto L7
L16:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+47)))
	switch v80 - int32(8) {
	case 0:
		goto L32
	default:
		goto L30
	case 3:
		goto L33
	case 11:
		goto L31
	}
L17:
	;
	v51 = l4
	goto L19
L18:
	;
	v51 = int32(0)
	goto L19
L19:
	;
	if v51 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	v56 = F_palloc(m, int32(8))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v73 = F_pullf_create(m, v15+int32(36), int32(4399636), l0, l2)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v36
	v63 = F_pullf_create(m, v15+int32(36), int32(4399600), v56, l2)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	if int32(0) <= v63 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	F_pfree(m, v56)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v348 = v63
	v350 = int32(0)
	goto L7
L27:
	;
	if v73 < int32(0) {
		v344 = v73
		goto L8
	} else {
		goto L28
	}
L28:
	;
	goto L16
L29:
	;
	v576 = v341
	v578 = int32(-100)
	v580 = v342
	goto L4
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v80
	F_px_debug(m, int32(467815), v15)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L9
	} else {
		goto L116
	}
L31:
	;
	if l4 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L32:
	;
	if l3 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L33:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v87 = F_pullf_read_fixed(m, v83, int32(1), v15+int32(48))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	if v87 < int32(0) {
		v564 = v87
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+48)))
	v95 = F_pullf_read_fixed(m, v83, int32(1), v15+int32(48))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	if v95 < int32(0) {
		v564 = v95
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+48)))
	if v99 == int32(0) {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v107 = v99
	goto L39
L39:
	;
	v116 = F_pullf_read(m, v83, v107, v15+int32(1088))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L41
	}
L40:
	;
	F_px_debug(m, int32(339550), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L47
	}
L41:
	;
	if v116 < int32(0) {
		v564 = v116
		goto L5
	} else {
		goto L42
	}
L42:
	;
	if v116 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v120 = v107 - v116
	if v120 <= int32(0) {
		goto L6
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	goto L40
L46:
	;
	v107 = v120
	goto L39
L47:
	;
	v564 = int32(-100)
	goto L5
L48:
	;
	v130 = int32(0)
	F_px_debug(m, int32(271441), v130)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L9
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v28 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v341 = v130
	v342 = v28
	goto L29
L52:
	;
	v135 = int32(0)
	F_px_debug(m, int32(441177), v135)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v145 = F_pullf_read_fixed(m, v141, int32(1), v15+int32(48))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L9
	} else {
		goto L56
	}
L55:
	;
	v341 = v135
	v342 = int32(1)
	goto L29
L56:
	;
	if v145 < int32(0) {
		v564 = v145
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v149
	switch v149 {
	case 0:
		goto L61
	case 1, 2:
		goto L60
	case 3:
		goto L59
	default:
		goto L58
	}
L58:
	;
	F_px_debug(m, int32(369229), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L72
	}
L59:
	;
	F_px_debug(m, int32(442832), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L9
	} else {
		goto L67
	}
L60:
	;
	v157 = F_pgp_decompress_filter(m, v15+int32(48), l0, v141)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L9
	} else {
		goto L63
	}
L61:
	;
	v151 = int32(0)
	v153 = F_process_data_packets(m, l0, l1, v141, v151, v151)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v564 = v153
	goto L5
L63:
	;
	if v157 < int32(0) {
		v564 = v157
		goto L5
	} else {
		goto L64
	}
L64:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v162 = int32(0)
	v164 = F_process_data_packets(m, l0, l1, v161, v162, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	F_pullf_free(m, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v564 = v164
	goto L5
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(1)
	goto L68
L68:
	;
	v190 = F_pullf_read(m, v141, int32(32768), v15+int32(48))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L9
	} else {
		goto L70
	}
L69:
	;
	v564 = v190
	goto L5
L70:
	;
	if int32(0) < v190 {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v564 = int32(-100)
	goto L5
L73:
	;
	v201 = int32(0)
	F_px_debug(m, int32(546033), v201)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v207 != 0 {
		v330 = int32(-12)
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v341 = v201
	v342 = v28
	goto L29
L77:
	;
	v332 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v332
	v576 = v332
	v578 = int32(0)
	v580 = v28
	goto L4
L78:
	;
	v576 = int32(0)
	v578 = v330
	v580 = v28
	goto L4
L79:
	;
	v208 = int32(-100)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v209 != int32(20) {
		v330 = v208
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
	v220 = F_pullf_read_max(m, v212, int32(20), v15+int32(1084), v15+int32(1088))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	if v220 < int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v576 = int32(0)
	v578 = v220
	v580 = v28
	goto L4
L83:
	;
	goto L84
L84:
	;
	if v220 != int32(20) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if v220 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v241)+16))
	m.T0[v244].(func(*base.Module, int32, int32))(m, v241, v15+int32(48))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L9
	} else {
		goto L93
	}
L88:
	;
	F_px_debug(m, int32(493389), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L9
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v220
	F_px_debug(m, int32(467589), v15+int32(32))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L9
	} else {
		goto L92
	}
L91:
	;
	v576 = int32(0)
	v578 = v208
	v580 = v28
	goto L4
L92:
	;
	v576 = int32(0)
	v578 = v208
	v580 = v28
	goto L4
L93:
	;
	v248 = v15 + int32(48)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1084))
	v250 = int32(20)
	goto L97
L94:
	;
	v317 = F___memset(m, v15+int32(48), int32(0), int32(20))
	mBase = m.M
	goto L112
L95:
	;
	v312 = int32(0)
	goto L94
L96:
	;
	v286 = v281
	v287 = v282
	v288 = v283
	goto L106
L97:
	;
	if (v248|v249)&int32(3) != 0 {
		v281 = v248
		v282 = v249
		v283 = v250
		goto L96
	} else {
		goto L100
	}
L99:
	;
	if v271 == int32(0) {
		goto L95
	} else {
		goto L105
	}
L100:
	;
	v258 = v248
	v259 = v249
	v260 = v250
	goto L101
L101:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if v263 != v264 {
		v281 = v258
		v282 = v259
		v283 = v260
		goto L96
	} else {
		goto L103
	}
L102:
	;
	goto L99
L103:
	;
	v266 = int32(4)
	v267 = v259 + v266
	v269 = v258 + v266
	v271 = v260 - v266
	if base.Ui32(int32(3)) < base.Ui32(v271) {
		v258 = v269
		v259 = v267
		v260 = v271
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v281 = v269
	v282 = v267
	v283 = v271
	goto L96
L106:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	if v291 == v292 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v312 = v291 - v292
	goto L94
L108:
	;
	v294 = int32(1)
	v299 = v288 - v294
	if v299 != 0 {
		v286 = v286 + v294
		v287 = v287 + v294
		v288 = v299
		goto L106
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	goto L107
L111:
	;
	goto L95
L112:
	;
	v322 = F___memset(m, v15+int32(1088), int32(0), int32(20))
	mBase = m.M
	goto L113
L113:
	;
	if v312 == int32(0) {
		goto L77
	} else {
		goto L114
	}
L114:
	;
	F_px_debug(m, int32(455744), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L9
	} else {
		goto L115
	}
L115:
	;
	v330 = v208
	goto L78
L116:
	;
	v341 = int32(0)
	v342 = v28
	goto L29
L117:
	;
	F_pullf_free(m, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L9
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v354 = int32(0)
	if v348 < v354 {
		v597 = v348
		goto L3
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	if v28 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	F_px_debug(m, int32(506072), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L9
	} else {
		goto L125
	}
L123:
	;
	v364 = v354
	goto L124
L124:
	;
	if base.B2i32(l4 == int32(0))|v350 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v364 = int32(-100)
	goto L124
L126:
	;
	v597 = v364
	goto L3
L127:
	;
	goto L128
L128:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v368 != 0 {
		v597 = v364
		goto L3
	} else {
		goto L129
	}
L129:
	;
	F_px_debug(m, int32(493363), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L9
	} else {
		goto L130
	}
L130:
	;
	v597 = int32(-100)
	goto L3
L131:
	;
	if v391 != int32(4) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	F_px_debug(m, int32(339550), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L9
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v404 = F___memset(m, v15+int32(1084), int32(0), int32(4))
	mBase = m.M
	goto L136
L135:
	;
	v564 = int32(-100)
	goto L5
L136:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v405 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = base.B2i32(v91 == int32(117))
	v432 = int32(0)
	goto L141
L138:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v91-int32(118)) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v91
	F_px_debug(m, int32(501800), v15+int32(16))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L9
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(1)
	goto L137
L141:
	;
	v439 = F_pullf_read(m, v83, int32(32768), v15+int32(1088))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L9
	} else {
		goto L143
	}
L142:
	;
	if v439 != 0 {
		v564 = v439
		goto L5
	} else {
		goto L183
	}
L143:
	;
	if int32(0) < v439 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v443 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	goto L146
L146:
	;
	goto L142
L147:
	;
	if int32(0) <= v544 {
		v432 = v545
		goto L141
	} else {
		goto L182
	}
L148:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1088))
	v535 = F_mbuf_append(m, l1, v534, v439)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L9
	} else {
		goto L181
	}
L149:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v446 == int32(0) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v15)+1088))
	v450 = int32(0)
	if v432 == v450 {
		v460 = v450
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v461 = v449 + v439
	v467 = v449
	v470 = v460
	goto L155
L152:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	if v454 == int32(10) {
		v460 = int32(0)
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v457 = int32(13)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+48)) = uint8(v457)
	v460 = int32(1)
	goto L151
L154:
	;
	if int32(0) < v515 {
		goto L175
	} else {
		goto L176
	}
L155:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467))))
	if v477 == int32(13) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v515 = v508
	v517 = int32(0)
	goto L154
L157:
	;
	v480 = int32(1)
	v482 = v467 + v480
	if base.Ui32(v461) <= base.Ui32(v482) {
		v515 = v470
		v517 = v480
		goto L154
	} else {
		goto L160
	}
L158:
	;
	v491 = v467
	v494 = v477
	goto L159
L159:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15+int32(48)+v470))) = uint8(v494)
	v497 = v470 + int32(1)
	if v470 < int32(1023) {
		goto L168
	} else {
		goto L169
	}
L160:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v467)+1)))
	v486 = base.B2i32(v484 == int32(10))
	if v484 == int32(10) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v487 = v482
	goto L163
L162:
	;
	v487 = v467
	goto L163
L163:
	;
	if v484 == int32(10) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v490 = int32(10)
	goto L166
L165:
	;
	v490 = int32(13)
	goto L166
L166:
	;
	v491 = v487
	v494 = v490
	goto L159
L167:
	;
	v510 = v491 + int32(1)
	if base.Ui32(v510) < base.Ui32(v461) {
		v467 = v510
		v470 = v508
		goto L155
	} else {
		goto L173
	}
L168:
	;
	v508 = v497
	goto L167
L169:
	;
	goto L170
L170:
	;
	v500 = int32(0)
	v503 = F_mbuf_append(m, l1, v15+int32(48), v497)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L9
	} else {
		goto L171
	}
L171:
	;
	if v503 < int32(0) {
		v544 = v503
		v545 = v500
		goto L147
	} else {
		goto L172
	}
L172:
	;
	v508 = v500
	goto L167
L173:
	;
	goto L156
L174:
	;
	v544 = v533
	v545 = v517
	goto L147
L175:
	;
	v522 = F_mbuf_append(m, l1, v15+int32(48), v515)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L9
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v527 = int32(0)
	v532 = F___memset(m, v15+int32(48), v527, int32(1024))
	mBase = m.M
	goto L180
L178:
	;
	if v522 < int32(0) {
		v533 = v522
		goto L174
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v533 = v527
	goto L174
L181:
	;
	v544 = v535
	v545 = v432
	goto L147
L182:
	;
	v564 = v544
	goto L5
L183:
	;
	if v432 == int32(0) {
		v564 = v439
		goto L5
	} else {
		goto L184
	}
L184:
	;
	v555 = F_mbuf_append(m, l1, int32(748405), int32(1))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L9
	} else {
		goto L185
	}
L185:
	;
	v564 = v555
	goto L5
L186:
	;
	v586 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v586
	if v586 <= v578 {
		v24 = v576
		v28 = v580
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v597 = v578
	goto L3
}
