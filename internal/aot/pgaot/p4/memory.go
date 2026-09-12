package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetMemoryChunkSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v4&int32(15)*int32(36))+uint32(_consts[1442])))
	v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_MemoryContextDeleteChildren(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = v8
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v20 = v12
	goto L6
L5:
	;
	goto L3
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v23 != 0 {
		v20 = v23
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v67 != 0 {
		v12 = v67
		goto L4
	} else {
		goto L29
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v28 = v25
	goto L12
L10:
	;
	v46 = v24
	goto L11
L11:
	;
	if v46 != 0 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	m.T0[v36].(func(*base.Module, int32))(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v46 = v40
	goto L11
L14:
	;
	return
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v39 != 0 {
		v28 = v39
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	if v49 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(0)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	m.T0[v63].(func(*base.Module, int32))(m, v20)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L27
	}
L20:
	;
	if v48 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+28)) = v48
	goto L20
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v48
	goto L20
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v52
	goto L26
L25:
	;
	goto L26
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(0)
	goto L19
L27:
	;
	if v20 != v12 {
		v20 = v24
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L7
L29:
	;
	goto L5
}
func F_MemoryContextStatsPrint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	v14 = m.G0
	v16 = v14 - int32(144)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l3 != 0 {
		goto L101
	} else {
		goto L102
	}
L2:
	;
	if v20&int32(3) == int32(0) {
		v73 = v20
		goto L17
	} else {
		goto L18
	}
L3:
	;
	v21 = int32(317316)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1295])))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v25 == int32(0) {
		v44 = v24
		v45 = v25
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v47 = v18
	goto L5
L5:
	;
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+32)) = uint8(v48)
	v373 = v47
	goto L1
L6:
	;
	if v45-v44 != 0 {
		goto L2
	} else {
		goto L14
	}
L7:
	;
	goto L6
L8:
	;
	if v24 != v25 {
		v44 = v24
		v45 = v25
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v29 = v18
	v30 = v21
	goto L10
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	if v34 == int32(0) {
		v44 = v33
		v45 = v34
		goto L7
	} else {
		goto L12
	}
L11:
	;
	v44 = v33
	v45 = v34
	goto L7
L12:
	;
	v37 = int32(1)
	if v33 == v34 {
		v29 = v29 + v37
		v30 = v30 + v37
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v47 = v20
	goto L5
L15:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1443])))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+34)) = uint8(v108)
	v111 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1444])))
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+32)) = uint16(v111)
	v114 = v16 + int32(32)
	if v114&int32(3) == int32(0) {
		v138 = v114
		goto L34
	} else {
		goto L35
	}
L16:
	;
	v106 = v98 - v20
	goto L15
L17:
	;
	v77 = v73
	goto L26
L18:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v57 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v106 = int32(0)
	goto L15
L20:
	;
	goto L21
L21:
	;
	v62 = v20
	goto L22
L22:
	;
	v66 = v62 + int32(1)
	if v66&int32(3) == int32(0) {
		v73 = v66
		goto L17
	} else {
		goto L24
	}
L23:
	;
	v98 = v66
	goto L16
L24:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v71 != 0 {
		v62 = v66
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v86 = int32(-2139062144)
	if (int32(16843008)-v83|v83)&v86 == v86 {
		v77 = v77 + int32(4)
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v92 = v77
	goto L29
L28:
	;
	goto L27
L29:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v96 != 0 {
		v92 = v92 + int32(1)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v98 = v92
	goto L16
L31:
	;
	goto L30
L32:
	;
	if int32(101) <= v106 {
		goto L49
	} else {
		goto L50
	}
L33:
	;
	v171 = v163 - v114
	goto L32
L34:
	;
	v142 = v138
	goto L43
L35:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v122 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v171 = int32(0)
	goto L32
L37:
	;
	goto L38
L38:
	;
	v127 = v114
	goto L39
L39:
	;
	v131 = v127 + int32(1)
	if v131&int32(3) == int32(0) {
		v138 = v131
		goto L34
	} else {
		goto L41
	}
L40:
	;
	v163 = v131
	goto L33
L41:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v136 != 0 {
		v127 = v131
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v151 = int32(-2139062144)
	if (int32(16843008)-v148|v148)&v151 == v151 {
		v142 = v142 + int32(4)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v157 = v142
	goto L46
L45:
	;
	goto L44
L46:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v161 != 0 {
		v157 = v157 + int32(1)
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v163 = v157
	goto L33
L48:
	;
	goto L47
L49:
	;
	v175 = F_pg_mbcliplen(m, v20, v106, int32(100))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v177 = v106
	goto L51
L51:
	;
	if v177 <= int32(0) {
		v281 = v171
		goto L54
	} else {
		goto L55
	}
L52:
	;
	return
L53:
	;
	v177 = v175
	goto L51
L54:
	;
	v296 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16+int32(32)+v281))) = uint8(v296)
	if v106 < int32(101) {
		v373 = v18
		goto L1
	} else {
		goto L82
	}
L55:
	;
	v181 = v177 & int32(3)
	if v181 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if base.Ui32(v177) < base.Ui32(int32(4)) {
		v281 = v216
		goto L54
	} else {
		goto L66
	}
L57:
	;
	v215 = v20
	v216 = v171
	v223 = v177
	goto L56
L58:
	;
	goto L59
L59:
	;
	v184 = v20
	v185 = v171
	v189 = int32(0)
	v192 = v177
	goto L60
L60:
	;
	v197 = int32(32)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if base.Ui32(v201) <= base.Ui32(v197) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v215 = v209
	v216 = v207
	v223 = v211
	goto L56
L62:
	;
	v204 = v197
	goto L64
L63:
	;
	v204 = v201
	goto L64
L64:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16+v197+v185))) = uint8(v204)
	v206 = int32(1)
	v207 = v185 + v206
	v209 = v184 + v206
	v211 = v192 - v206
	v213 = v189 + v206
	if v213 != v181 {
		v184 = v209
		v185 = v207
		v189 = v213
		v192 = v211
		goto L60
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	v230 = v215
	v231 = v216
	v238 = v223
	goto L67
L67:
	;
	v243 = int32(32)
	v245 = v16 + v243 + v231
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	if base.Ui32(v247) <= base.Ui32(v243) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v281 = v271
	goto L54
L69:
	;
	v250 = v243
	goto L71
L70:
	;
	v250 = v247
	goto L71
L71:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v250)
	v252 = int32(32)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if base.Ui32(v253) <= base.Ui32(v252) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v256 = v252
	goto L74
L73:
	;
	v256 = v253
	goto L74
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)) = uint8(v256)
	v258 = int32(32)
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+2)))
	if base.Ui32(v259) <= base.Ui32(v258) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v262 = v258
	goto L77
L76:
	;
	v262 = v259
	goto L77
L77:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v245)+2)) = uint8(v262)
	v264 = int32(32)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+3)))
	if base.Ui32(v265) <= base.Ui32(v264) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v268 = v264
	goto L80
L79:
	;
	v268 = v265
	goto L80
L80:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v245)+3)) = uint8(v268)
	v270 = int32(4)
	v271 = v231 + v270
	if base.Ui32(v238-int32(5)) < base.Ui32(int32(-2)) {
		v230 = v230 + v270
		v231 = v271
		v238 = v238 - v270
		goto L67
	} else {
		goto L81
	}
L81:
	;
	goto L68
L82:
	;
	v301 = v16 + int32(32)
	if v301&int32(3) == int32(0) {
		v325 = v301
		goto L85
	} else {
		goto L86
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v358+(v16+int32(32))))) = int32(3026478)
	v373 = v18
	goto L1
L84:
	;
	v358 = v350 - v301
	goto L83
L85:
	;
	v329 = v325
	goto L94
L86:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v309 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v358 = int32(0)
	goto L83
L88:
	;
	goto L89
L89:
	;
	v314 = v301
	goto L90
L90:
	;
	v318 = v314 + int32(1)
	if v318&int32(3) == int32(0) {
		v325 = v318
		goto L85
	} else {
		goto L92
	}
L91:
	;
	v350 = v318
	goto L84
L92:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	if v323 != 0 {
		v314 = v318
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	v338 = int32(-2139062144)
	if (int32(16843008)-v335|v335)&v338 == v338 {
		v329 = v329 + int32(4)
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v344 = v329
	goto L97
L96:
	;
	goto L95
L97:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	if v348 != 0 {
		v344 = v344 + int32(1)
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v350 = v344
	goto L84
L99:
	;
	goto L98
L100:
	;
	m.G0 = v16 + int32(144)
	return
L101:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _consts[1410]))
	if int32(2) <= v19 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	v425 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L52
	} else {
		goto L112
	}
L104:
	;
	v382 = int32(1)
	goto L107
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v373
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v16 + int32(32)
	v421 = F_pg_fprintf(m, v378, int32(721655), v16)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L52
	} else {
		goto L111
	}
L107:
	;
	v397 = F_pg_fprintf(m, v378, int32(720101), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L52
	} else {
		goto L109
	}
L108:
	;
	goto L106
L109:
	;
	v400 = v382 + int32(1)
	if v400 != v19 {
		v382 = v400
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	goto L100
L112:
	;
	if v425 == int32(0) {
		goto L100
	} else {
		goto L113
	}
L113:
	;
	F_errhidestmt(m)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L52
	} else {
		goto L114
	}
L114:
	;
	F_errhidecontext(m)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L52
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v373
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v16 + int32(32)
	F_errmsg_internal(m, int32(172815), v16+int32(16))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L52
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(483770), int32(1044), int32(86755))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L52
	} else {
		goto L117
	}
L117:
	;
	goto L100
}
