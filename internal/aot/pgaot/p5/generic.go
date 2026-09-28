package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generic_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+2)))
	v18 = v16 + int32(4)
	if base.Ui32(v18) < base.Ui32(v13) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(32)
	return
L4:
	;
	v23 = v14 + v18
	v26 = v15
	v27 = v16
	goto L7
L5:
	;
	v46 = v15
	v47 = v16
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v46
	F_appendStringInfo(m, l0, int32(_a_F_generic_desc_0), v10)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L12
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v26
	F_appendStringInfo(m, l0, int32(_a_F_generic_desc_1), v10+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v46 = v36
	v47 = v37
	goto L6
L9:
	;
	return
L10:
	;
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+2)))
	v40 = v23 + v37 + int32(4)
	if base.Ui32(v40) < base.Ui32(v14+v13) {
		v23 = v40
		v26 = v36
		v27 = v37
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L3
}
func F_transformGenericOptions(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v239 int64
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
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
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int64
	_ = v308
	var v309 int32
	_ = v309
	var v311 int64
	_ = v311
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = F_untransformRelOptions(m, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if l2 == int32(0) {
		v233 = v16
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v239 = int64(0)
	if v233 == int32(0) {
		v311 = v239
		goto L69
	} else {
		goto L70
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v22 <= int32(0) {
		v233 = v16
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v30 = v16
	v34 = int32(0)
	goto L9
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L65
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L61
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L57
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v34<<(uint(int32(2))%32))))
	if v30 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L53
	}
L11:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	switch v109 {
	case 0, 2:
		goto L33
	case 1:
		goto L34
	case 3:
		goto L35
	default:
		goto L6
	}
L12:
	;
	v104 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v45 = int32(0)
	if v45 < v44 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v48 = v44
	goto L17
L16:
	;
	v48 = v45
	goto L17
L17:
	;
	v58 = int32(0)
	goto L18
L18:
	;
	if v58 == v48 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v104 = v68
	goto L11
L20:
	;
	v104 = int32(0)
	goto L11
L21:
	;
	goto L22
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v68 = v58<<(uint(int32(2))%32) + v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if base.B2i32(v74 == int32(0))|base.B2i32(v74 != v77) != 0 {
		v95 = v74
		v96 = v77
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v95-v96 != 0 {
		v58 = v58 + int32(1)
		goto L18
	} else {
		goto L30
	}
L24:
	;
	goto L23
L25:
	;
	v80 = v70
	v81 = v71
	goto L26
L26:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	if v85 == int32(0) {
		v95 = v85
		v96 = v84
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v95 = v85
	v96 = v84
	goto L24
L28:
	;
	v88 = int32(1)
	if v85 == v84 {
		v80 = v80 + v88
		v81 = v81 + v88
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L19
L31:
	;
	goto L10
L32:
	;
	v150 = v34 + int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v150 < v151 {
		v30 = v147
		v34 = v150
		goto L9
	} else {
		goto L52
	}
L33:
	;
	if v104 != 0 {
		goto L7
	} else {
		goto L50
	}
L34:
	;
	if v104 == int32(0) {
		goto L8
	} else {
		goto L49
	}
L35:
	;
	if v104 == int32(0) {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v113 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v147 = v140
	goto L32
L38:
	;
	if v30+int32(16) != v112 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v127 = int32(2)
	v131 = (v113 + int32(base.Ui32(v104-v112^int32(-1))>>(uint(v127)%32))) << (uint(v127) % 32)
	if v131 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	F_pfree(m, v112)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_pfree(m, v30)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	v140 = int32(0)
	goto L37
L46:
	;
	base.MemoryCopy(m, v104, v104+int32(4), v131)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v135 - int32(1)
	v140 = v30
	goto L37
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104))) = v40
	v147 = v30
	goto L32
L50:
	;
	v144 = F_lappend(m, v30, v40)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v147 = v144
	goto L32
L52:
	;
	v233 = v147
	goto L3
L53:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v160
	F_errmsg(m, int32(_a_F_transformGenericOptions_0), v14+int32(48))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_transformGenericOptions_1), int32(160), int32(_a_F_transformGenericOptions_2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v179
	F_errmsg(m, int32(_a_F_transformGenericOptions_0), v14-int32(-64))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_transformGenericOptions_1), int32(169), int32(_a_F_transformGenericOptions_2))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(_a_F_transformGenericOptions_3))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v198
	F_errmsg(m, int32(_a_F_transformGenericOptions_4), v14+int32(80))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_transformGenericOptions_1), int32(179), int32(_a_F_transformGenericOptions_2))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v214
	F_errmsg_internal(m, int32(_a_F_transformGenericOptions_5), v14+int32(32))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_transformGenericOptions_1), int32(185), int32(_a_F_transformGenericOptions_2))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L94
	}
L69:
	;
	if l3 != 0 {
		goto L86
	} else {
		goto L87
	}
L70:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v242 <= int32(0) {
		v311 = v239
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v245 = int32(0)
	v253 = v245
	v256 = v245
	goto L72
L72:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v258+v253<<(uint(int32(2))%32))))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	v264 = F_defGetString(m, v262)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	if v298 == int32(0) {
		v311 = v239
		goto L69
	} else {
		goto L84
	}
L74:
	;
	v266 = int32(61)
	v267 = F___strchrnul(m, v263, v266)
	mBase = m.M
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	if v269 == v266 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v273 != 0 {
		goto L68
	} else {
		goto L79
	}
L76:
	;
	v273 = v267
	goto L78
L77:
	;
	v273 = int32(0)
	goto L78
L78:
	;
	goto L75
L79:
	;
	v274 = F_strlen(m, v263)
	mBase = m.M
	v275 = F_strlen(m, v264)
	mBase = m.M
	v276 = v274 + v275
	v279 = F_palloc(m, v276+int32(6))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = v276<<(uint(int32(2))%32) + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v263
	v291 = F_pg_sprintf(m, v279+int32(4), int32(_a_F_transformGenericOptions_6), v14)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_transformGenericOptions[0]))
	v298 = F_accumArrayResult(m, v256, base.I64_extend_i32_u(v279), int32(0), int32(25), v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v301 = v253 + int32(1)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v301 < v302 {
		v253 = v301
		v256 = v298
		goto L72
	} else {
		goto L83
	}
L83:
	;
	goto L73
L84:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_transformGenericOptions[0]))
	v308 = F_makeArrayResult(m, v298, v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v311 = v308
	goto L69
L86:
	;
	if base.I32_wrap_i64(v311) != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	m.G0 = v14 + int32(96)
	return v311
L89:
	;
	v327 = v311
	goto L91
L90:
	;
	v324 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L92
	}
L91:
	;
	v329 = F_OidFunctionCall2Coll(m, l3, int32(0), v327, base.I64_extend_i32_u(l0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L93
	}
L92:
	;
	v327 = base.I64_extend_i32_u(v324)
	goto L91
L93:
	;
	goto L88
L94:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v263
	F_errmsg(m, int32(_a_F_transformGenericOptions_7), v14+int32(16))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_transformGenericOptions_1), int32(87), int32(_a_F_transformGenericOptions_8))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
