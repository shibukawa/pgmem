package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FetchPortalTargetList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v2 != int32(4) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v8 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v46 = int32(0)
	goto L3
L3:
	;
	return v46
L4:
	;
	v41 = F_FetchStatementTargetList(m, v39)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	goto L4
L6:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v9 <= int32(0) {
		v39 = int32(0)
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v39 = int32(0)
	goto L5
L9:
	;
	v12 = int32(0)
	if v12 < v9 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v15 = v9
	goto L12
L11:
	;
	v15 = v12
	goto L12
L12:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v18 = int32(0)
	goto L13
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16+v18<<(uint(int32(2))%32))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+26)))
	if v26 == int32(1) {
		v39 = v25
		goto L5
	} else {
		goto L15
	}
L14:
	;
	goto L8
L15:
	;
	v30 = v18 + int32(1)
	if v30 != v15 {
		v18 = v30
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	return int32(0)
L18:
	;
	v46 = v41
	goto L3
}
func F_PortalRunFetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v140 int64
	_ = v140
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int64
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v205 int32
	_ = v205
	var v208 int64
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int64
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	var v259 int64
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
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
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v294 int64
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int64
	_ = v306
	var v307 int32
	_ = v307
	var v308 int64
	_ = v308
	var v312 int32
	_ = v312
	var v314 int64
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int64
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int64
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v340 int64
	_ = v340
	var v369 int32
	_ = v369
	var v382 int32
	_ = v382
	var v393 int32
	_ = v393
	var v394 int64
	_ = v394
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(192)
	m.G0 = v18
	v26 = v5
	v27 = v5
	v28 = v5
	v29 = v5
	v30 = v5
	v31 = v5
	v32 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v32 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v393 = int32(m.ExcTag)
	v394 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v393 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v18)+172)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v18)+180)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v31
	F_MarkPortalActive(m, l0)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v63 = v26
	v64 = v27
	v65 = v28
	v66 = v29
	v67 = v30
	v68 = v31
	goto L9
L9:
	;
	if v63 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[0]))
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[1]))
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[2]))
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[3]))
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[4]))
	goto L11
L11:
	;
	v57 = v18 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v18 + int32(12)
	goto L14
L12:
	;
	v63 = int32(0)
	v64 = v49
	v65 = v47
	v66 = v51
	v67 = v53
	v68 = v55
	goto L9
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[4])) = l0
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[1])) = v18 + int32(16)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v77 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[0])) = v65
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[1])) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v18)+172)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v18)+180)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v68
	F_MarkPortalFailed(m, l0)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L6
	} else {
		goto L118
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[3])) = v77
	goto L20
L19:
	;
	goto L20
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[2])) = v81
	v83 = int32(_a_F_PortalRunFetch_0)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[5])) = v81
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if base.Ui32(int32(3)) <= base.Ui32(v87-int32(1)) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v18)+172)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v18)+180)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v68
	switch l1 {
	case 0:
		goto L38
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	default:
		goto L39
	}
L22:
	;
	if v87 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v122 != 0 {
		goto L21
	} else {
		goto L29
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v18)+172)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v18)+180)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v68
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v18)+172)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v18)+180)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v68
	F_errmsg_internal(m, int32(_a_F_PortalRunFetch_1), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v18)+172)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v18)+180)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v68
	F_errfinish(m, int32(_a_F_PortalRunFetch_2), int32(1434), int32(_a_F_PortalRunFetch_3))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L3
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v18)+172)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v18)+180)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v68
	F_FillPortalStore(m, l0, int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	goto L21
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[0])) = v65
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[1])) = v64
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[5])) = v84
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[3])) = v67
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[4])) = v68
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[2])) = v66
	m.G0 = v18 + int32(192)
	return v340
L32:
	;
	v340 = base.I64_extend_i32_u(v332) & int64(1)
	goto L31
L33:
	;
	v330 = F_PortalRunSelect(m, l0, v329, v328, l3)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L6
	} else {
		goto L117
	}
L34:
	;
	v320 = int32(1)
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[6]))
	v325 = F_PortalRunSelect(m, l0, int32(0), v320, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L6
	} else {
		goto L116
	}
L35:
	;
	v297 = l2 >> (uint(int32(31)) % 32)
	v299 = l2 ^ v297 - v297
	if base.B2i32(v299 != int32(2147483647))|v276 != 0 {
		v328 = v299
		v329 = v276
		goto L33
	} else {
		goto L106
	}
L36:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v278 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L37:
	;
	if l2 != 0 {
		goto L35
	} else {
		goto L97
	}
L38:
	;
	v276 = base.B2i32(int32(0) <= l2)
	goto L37
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L94
	}
L40:
	;
	if int32(0) < l2 {
		goto L80
	} else {
		goto L81
	}
L41:
	;
	if int32(0) < l2 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v276 = int32(base.Ui32(l2) >> (uint(int32(31)) % 32))
	goto L37
L43:
	;
	v140 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v144 = l2 - int32(1)
	if base.B2i32(base.Ui64(v140) <= base.Ui64(int64(2147483646)))&base.B2i32(base.Ui64(int64(base.Ui64(v140)>>(uint(int64(1))%64))) < base.Ui64(base.I64_extend_i32_u(v144))) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	goto L45
L45:
	;
	if l2 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L46:
	;
	v185 = int32(1)
	v187 = F_PortalRunSelect(m, l0, v185, v185, l3)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L6
	} else {
		goto L59
	}
L47:
	;
	F_DoPortalRewind(m, l0)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v163 = v161 + base.I32_wrap_i64(v140)
	if base.Ui32(l2) <= base.Ui32(v163) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	if l2 == int32(1) {
		goto L46
	} else {
		goto L51
	}
L51:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[6]))
	v159 = F_PortalRunSelect(m, l0, int32(1), v144, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	goto L46
L53:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[6]))
	v171 = F_PortalRunSelect(m, l0, int32(0), v163-l2+int32(1), v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(l2) <= base.Ui32(v163+int32(1)) {
		goto L46
	} else {
		goto L57
	}
L56:
	;
	goto L46
L57:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[6]))
	v182 = F_PortalRunSelect(m, l0, int32(1), l2+(v163^int32(-1)), v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	goto L46
L59:
	;
	v340 = v187
	goto L31
L60:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[6]))
	v195 = F_PortalRunSelect(m, l0, int32(1), int32(2147483647), v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L6
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_DoPortalRewind(m, l0)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L6
	} else {
		goto L69
	}
L63:
	;
	if l2 != int32(-1) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[6]))
	v204 = F_PortalRunSelect(m, l0, int32(0), l2^int32(-1), v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v208 = F_PortalRunSelect(m, l0, int32(0), int32(1), l3)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L6
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	v340 = v208
	goto L31
L69:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v212 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+20)) = l3
	goto L72
L71:
	;
	goto L72
L72:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v214 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v217 = F_RunFromStore(m, l0, int32(0), int64(0), l3)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L6
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v212)+12))
	F_PushActiveSnapshot(m, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L77
	}
L76:
	;
	v340 = v217
	goto L31
L77:
	;
	F_ExecutorRun(m, v212, int32(0), int64(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v212)+40))
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v226)+112))
	F_PopActiveSnapshot(m)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	v340 = v227
	goto L31
L80:
	;
	if l2 != int32(1) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	if int32(0) <= l2 {
		v277 = int32(1)
		goto L36
	} else {
		goto L88
	}
L83:
	;
	v235 = int32(1)
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[6]))
	v240 = F_PortalRunSelect(m, l0, v235, l2-v235, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v242 = int32(1)
	v244 = F_PortalRunSelect(m, l0, v242, v242, l3)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L6
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	v340 = v244
	goto L31
L88:
	;
	if l2 != int32(-1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[6]))
	v255 = F_PortalRunSelect(m, l0, int32(0), l2^int32(-1), v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v259 = F_PortalRunSelect(m, l0, int32(0), int32(1), l3)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L6
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	v340 = v259
	goto L31
L94:
	;
	F_errmsg_internal(m, int32(_a_F_PortalRunFetch_4), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_PortalRunFetch_2), int32(1605), int32(_a_F_PortalRunFetch_5))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v277 = v276
	goto L36
L98:
	;
	if v281&int32(1) == int32(0) {
		goto L34
	} else {
		goto L104
	}
L99:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v282 != 0 {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v285 = int32(0)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v286 == v285 {
		v332 = v285
		goto L32
	} else {
		goto L103
	}
L102:
	;
	v332 = v281 ^ int32(1)
	goto L32
L103:
	;
	v328 = v285
	v329 = v277
	goto L33
L104:
	;
	v294 = F_PortalRunSelect(m, l0, v277, int32(0), l3)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	v340 = v294
	goto L31
L106:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v303 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v306 = F_PortalRunSelect(m, l0, int32(0), int32(2147483647), l3)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v308 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	if v308 == int64(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v340 = v306
	goto L31
L111:
	;
	F_DoPortalRewind(m, l0)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L6
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v314 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	F_DoPortalRewind(m, l0)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L6
	} else {
		goto L115
	}
L114:
	;
	v340 = int64(0)
	goto L31
L115:
	;
	v340 = v308 - (v314 ^ int64(1))
	goto L31
L116:
	;
	v328 = int32(1)
	v329 = v320
	goto L33
L117:
	;
	v340 = v330
	goto L31
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[3])) = v67
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[4])) = v68
	*(*int32)(unsafe.Add(mBase, _c_F_PortalRunFetch[2])) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v18)+172)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v18)+176)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v18)+180)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v68
	F_pg_re_throw(m)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	goto L5
L120:
	;
	v398 = int32(v394)
	m.G0 = v18
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v398)+4))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	if v18+int32(12) == v404 {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	m.ExcPending = 1
	goto L129
L122:
	;
	if v408 != 0 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	v408 = v406
	goto L125
L124:
	;
	v408 = int32(0)
	goto L125
L125:
	;
	goto L122
L126:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v18)+184))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v18)+180))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v18)+176))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v18)+172))
	v26 = v400
	v27 = v412
	v28 = v413
	v29 = v411
	v30 = v410
	v31 = v409
	v32 = v408
	goto L1
L127:
	;
	goto L128
L128:
	;
	F___wasm_longjmp(m, v401, v400)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	return int64(0)
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
