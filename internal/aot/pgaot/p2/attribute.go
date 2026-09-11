package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyAttributeOutText(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
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
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v9 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1&int32(3) == int32(0) {
		v35 = l1
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v72 = l1
	goto L3
L3:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
	if v74 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L4:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v70 = F_pg_server_to_any(m, l1, v68, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v68 = v60 - l1
	goto L4
L6:
	;
	v39 = v35
	goto L15
L7:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v19 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v68 = int32(0)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v24 = l1
	goto L11
L11:
	;
	v28 = v24 + int32(1)
	if v28&int32(3) == int32(0) {
		v35 = v28
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v60 = v28
	goto L5
L13:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v33 != 0 {
		v24 = v28
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v48 = int32(-2139062144)
	if (int32(16843008)-v45|v45)&v48 == v48 {
		v39 = v39 + int32(4)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v54 = v39
	goto L18
L17:
	;
	goto L16
L18:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v58 != 0 {
		v54 = v54 + int32(1)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v60 = v54
	goto L5
L20:
	;
	goto L19
L21:
	;
	return
L22:
	;
	v72 = v70
	goto L3
L23:
	;
	return
L24:
	;
	if base.Ui32(v347) <= base.Ui32(v349) {
		goto L23
	} else {
		goto L111
	}
L25:
	;
	v219 = v73
	v220 = v72
	v221 = v72
	goto L72
L26:
	;
	if v73&int32(255) == int32(0) {
		goto L23
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v73&int32(255) == int32(0) {
		goto L23
	} else {
		goto L30
	}
L29:
	;
	goto L25
L30:
	;
	v86 = v72
	v87 = v73
	v88 = v72
	goto L31
L31:
	;
	v92 = v87 & int32(255)
	if base.Ui32(v92) <= base.Ui32(int32(31)) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v347 = v215
	v349 = v213
	goto L24
L33:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v216 != 0 {
		v86 = v215
		v87 = v216
		v88 = v213
		goto L31
	} else {
		goto L71
	}
L34:
	;
	v213 = v88
	v215 = v86 + int32(1)
	goto L33
L35:
	;
	v96 = v87 - int32(8)
	if base.Ui32(int32(6)) <= base.Ui32(v96&int32(255)) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	goto L37
L37:
	;
	if base.B2i32(v92 != int32(92))&base.B2i32(v92 != v8) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L38:
	;
	if base.Ui32(v88) < base.Ui32(v86) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	if v92 != v8 {
		goto L34
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v110 = base.I32_wrap_i64(int64(base.Ui64(int64(125784399180898)) >> (uint(base.I64_extend_i32_u(v96<<(uint(int32(3))%32))&int64(248)) % 64)))
	goto L38
L42:
	;
	v110 = v8
	goto L38
L43:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v112, v88, v86-v88)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L21
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	if v120 <= v117+int32(1) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L45
L47:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140)+8))
	if v144 <= v141+int32(1) {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	F_appendStringInfoChar(m, v116, int32(92))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L21
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v127 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v125+v117))) = uint8(v127)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v132 = v130 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v134+v132))) = uint8(v136)
	goto L47
L51:
	;
	goto L47
L52:
	;
	v164 = v86 + int32(1)
	v213 = v164
	v215 = v164
	goto L33
L53:
	;
	F_appendStringInfoChar(m, v140, base.I32_extend8_s(v110))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L21
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	*(*uint8)(unsafe.Add(mBase, uint32(v149+v141))) = uint8(v110)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v155 = v153 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v152)+4)) = v155
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v157+v155))) = uint8(v159)
	goto L52
L56:
	;
	goto L52
L57:
	;
	if base.Ui32(v88) < base.Ui32(v86) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	if int32(0) <= base.I32_extend8_s(v87) {
		goto L34
	} else {
		goto L69
	}
L60:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v172, v88, v86-v88)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L21
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
	if v180 <= v177+int32(1) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L62
L64:
	;
	v213 = v86
	v215 = v86 + int32(1)
	goto L33
L65:
	;
	F_appendStringInfoChar(m, v176, int32(92))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L21
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v187 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v185+v177))) = uint8(v187)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v192 = v190 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v194+v192))) = uint8(v196)
	goto L64
L68:
	;
	goto L64
L69:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v206 = F_pg_encoding_mblen(m, v205, v86)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L21
	} else {
		goto L70
	}
L70:
	;
	v213 = v88
	v215 = v206 + v86
	goto L33
L71:
	;
	goto L32
L72:
	;
	v224 = v219 & int32(255)
	if base.Ui32(v224) <= base.Ui32(int32(31)) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	if v343 == int32(0) {
		v347 = v342
		v349 = v341
		goto L24
	} else {
		goto L110
	}
L75:
	;
	v228 = v219 - int32(8)
	if base.Ui32(int32(6)) <= base.Ui32(v228&int32(255)) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	goto L77
L77:
	;
	if base.B2i32(v224 != int32(92))&base.B2i32(v224 != v8) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L78:
	;
	if base.Ui32(v220) < base.Ui32(v221) {
		goto L84
	} else {
		goto L85
	}
L79:
	;
	if v224 == v8 {
		v245 = v8
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v245 = base.I32_wrap_i64(int64(base.Ui64(int64(125784399180898)) >> (uint(base.I64_extend_i32_u(v228<<(uint(int32(3))%32))&int64(248)) % 64)))
	goto L78
L82:
	;
	v235 = v221 + int32(1)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	if v236 != 0 {
		v219 = v236
		v221 = v235
		goto L72
	} else {
		goto L83
	}
L83:
	;
	v347 = v235
	v349 = v220
	goto L24
L84:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v247, v220, v221-v220)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L21
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	if v255 <= v252+int32(1) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L86
L88:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	if v279 <= v276+int32(1) {
		goto L94
	} else {
		goto L95
	}
L89:
	;
	F_appendStringInfoChar(m, v251, int32(92))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L21
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v262 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v260+v252))) = uint8(v262)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v267 = v265 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v271 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v269+v267))) = uint8(v271)
	goto L88
L92:
	;
	goto L88
L93:
	;
	v299 = v221 + int32(1)
	v341 = v299
	v342 = v299
	goto L74
L94:
	;
	F_appendStringInfoChar(m, v275, base.I32_extend8_s(v245))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L21
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	*(*uint8)(unsafe.Add(mBase, uint32(v284+v276))) = uint8(v245)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	v290 = v288 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v287)+4)) = v290
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v294 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v292+v290))) = uint8(v294)
	goto L93
L97:
	;
	goto L93
L98:
	;
	if base.Ui32(v220) < base.Ui32(v221) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	v341 = v220
	v342 = v221 + int32(1)
	goto L74
L101:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v307, v220, v221-v220)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L21
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v311)+8))
	if v315 <= v312+int32(1) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L103
L105:
	;
	v341 = v221
	v342 = v221 + int32(1)
	goto L74
L106:
	;
	F_appendStringInfoChar(m, v311, int32(92))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L21
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v322 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v320+v312))) = uint8(v322)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	v327 = v325 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v324)+4)) = v327
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
	v331 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v329+v327))) = uint8(v331)
	goto L105
L109:
	;
	goto L105
L110:
	;
	v219 = v343
	v220 = v341
	v221 = v342
	goto L72
L111:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_appendBinaryStringInfo(m, v353, v349, v347-v349)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L21
	} else {
		goto L112
	}
L112:
	;
	goto L23
}
func F_attribute_reloptions(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(64), int32(24), int32(_a_F_attribute_reloptions_0), int32(2))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_get_attribute_options(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
	if v11 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1571)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(51539607560)
		v24 = F_hash_create(m, int32(_a_F_get_attribute_options_0), int32(256), v6+int32(-48), int32(72))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0])) = v24
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[1]))
			if v30 == int32(0) {
				F_CreateCacheMemoryContext(m)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_CacheRegisterSyscacheCallback(m, int32(7), int32(1572), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
						v42 = v41
						*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
						v47 = int32(0)
						v49 = F_hash_search(m, v42, v6+int32(-48), v47, v47)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 != 0 {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
								v97 = v51
								v98 = v49
								if v97 == int32(0) {
									v113 = int32(0)
									m.G0 = v8 - int32(-64)
									return v113
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
									v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
										v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
										if v110 != 0 {
											v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
											mBase = m.M
										} else {
										}
										v113 = v105
										m.G0 = v8 - int32(-64)
										return v113
									}
								}
							} else {
								v52 = int32(0)
								v55 = F_SearchSysCache2(m, int32(7), l0, base.I32_extend16_s(l1))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									if v55 != 0 {
										v61 = F_SysCacheGetAttr(m, int32(7), v55, int32(23), v6+int32(-49))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
											if v63 == int32(0) {
												v67 = F_attribute_reloptions(m, v61, int32(0))
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v70 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[1]))
													v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
													v74 = F_MemoryContextAlloc(m, v70, int32(base.Ui32(v71)>>(uint(int32(2))%32)))
													mBase = m.M
													v75 = m.ExcPending
													if v75 != 0 {
														return int32(0)
													} else {
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
														v78 = int32(base.Ui32(v76) >> (uint(int32(2)) % 32))
														if v78 != 0 {
															v79 = F__emscripten_memcpy_bulkmem(m, v74, v67, v78)
															mBase = m.M
														} else {
														}
														v82 = v74
														F_ReleaseCatCache(m, v55)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															v86 = v82
															v88 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
															v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
																v97 = v86
																v98 = v93
																if v97 == int32(0) {
																	v113 = int32(0)
																	m.G0 = v8 - int32(-64)
																	return v113
																} else {
																	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
																	v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
																	mBase = m.M
																	v106 = m.ExcPending
																	if v106 != 0 {
																		return int32(0)
																	} else {
																		v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
																		v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
																		v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
																		if v110 != 0 {
																			v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
																			mBase = m.M
																		} else {
																		}
																		v113 = v105
																		m.G0 = v8 - int32(-64)
																		return v113
																	}
																}
															}
														}
													}
												}
											} else {
												v82 = v52
												F_ReleaseCatCache(m, v55)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													v86 = v82
													v88 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
													v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
														v97 = v86
														v98 = v93
														if v97 == int32(0) {
															v113 = int32(0)
															m.G0 = v8 - int32(-64)
															return v113
														} else {
															v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
															v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return int32(0)
															} else {
																v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
																v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
																v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
																if v110 != 0 {
																	v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
																	mBase = m.M
																} else {
																}
																v113 = v105
																m.G0 = v8 - int32(-64)
																return v113
															}
														}
													}
												}
											}
										}
									} else {
										v86 = v52
										v88 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
										v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
											v97 = v86
											v98 = v93
											if v97 == int32(0) {
												v113 = int32(0)
												m.G0 = v8 - int32(-64)
												return v113
											} else {
												v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
												v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
												mBase = m.M
												v106 = m.ExcPending
												if v106 != 0 {
													return int32(0)
												} else {
													v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
													v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
													v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
													if v110 != 0 {
														v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
														mBase = m.M
													} else {
													}
													v113 = v105
													m.G0 = v8 - int32(-64)
													return v113
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
				F_CacheRegisterSyscacheCallback(m, int32(7), int32(1572), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
					v42 = v41
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
					v47 = int32(0)
					v49 = F_hash_search(m, v42, v6+int32(-48), v47, v47)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if v49 != 0 {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
							v97 = v51
							v98 = v49
							if v97 == int32(0) {
								v113 = int32(0)
								m.G0 = v8 - int32(-64)
								return v113
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
								v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
									v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
									if v110 != 0 {
										v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
										mBase = m.M
									} else {
									}
									v113 = v105
									m.G0 = v8 - int32(-64)
									return v113
								}
							}
						} else {
							v52 = int32(0)
							v55 = F_SearchSysCache2(m, int32(7), l0, base.I32_extend16_s(l1))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								if v55 != 0 {
									v61 = F_SysCacheGetAttr(m, int32(7), v55, int32(23), v6+int32(-49))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
										if v63 == int32(0) {
											v67 = F_attribute_reloptions(m, v61, int32(0))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v70 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[1]))
												v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
												v74 = F_MemoryContextAlloc(m, v70, int32(base.Ui32(v71)>>(uint(int32(2))%32)))
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return int32(0)
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
													v78 = int32(base.Ui32(v76) >> (uint(int32(2)) % 32))
													if v78 != 0 {
														v79 = F__emscripten_memcpy_bulkmem(m, v74, v67, v78)
														mBase = m.M
													} else {
													}
													v82 = v74
													F_ReleaseCatCache(m, v55)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														v86 = v82
														v88 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
														v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
															v97 = v86
															v98 = v93
															if v97 == int32(0) {
																v113 = int32(0)
																m.G0 = v8 - int32(-64)
																return v113
															} else {
																v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
																v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return int32(0)
																} else {
																	v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
																	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
																	v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
																	if v110 != 0 {
																		v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
																		mBase = m.M
																	} else {
																	}
																	v113 = v105
																	m.G0 = v8 - int32(-64)
																	return v113
																}
															}
														}
													}
												}
											}
										} else {
											v82 = v52
											F_ReleaseCatCache(m, v55)
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return int32(0)
											} else {
												v86 = v82
												v88 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
												v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
													v97 = v86
													v98 = v93
													if v97 == int32(0) {
														v113 = int32(0)
														m.G0 = v8 - int32(-64)
														return v113
													} else {
														v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
														v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
															v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
															v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
															if v110 != 0 {
																v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
																mBase = m.M
															} else {
															}
															v113 = v105
															m.G0 = v8 - int32(-64)
															return v113
														}
													}
												}
											}
										}
									}
								} else {
									v86 = v52
									v88 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
									v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
										v97 = v86
										v98 = v93
										if v97 == int32(0) {
											v113 = int32(0)
											m.G0 = v8 - int32(-64)
											return v113
										} else {
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
											v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
												v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
												if v110 != 0 {
													v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
													mBase = m.M
												} else {
												}
												v113 = v105
												m.G0 = v8 - int32(-64)
												return v113
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
		v42 = v11
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
		v47 = int32(0)
		v49 = F_hash_search(m, v42, v6+int32(-48), v47, v47)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			if v49 != 0 {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
				v97 = v51
				v98 = v49
				if v97 == int32(0) {
					v113 = int32(0)
					m.G0 = v8 - int32(-64)
					return v113
				} else {
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
					v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
						v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
						if v110 != 0 {
							v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
							mBase = m.M
						} else {
						}
						v113 = v105
						m.G0 = v8 - int32(-64)
						return v113
					}
				}
			} else {
				v52 = int32(0)
				v55 = F_SearchSysCache2(m, int32(7), l0, base.I32_extend16_s(l1))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					if v55 != 0 {
						v61 = F_SysCacheGetAttr(m, int32(7), v55, int32(23), v6+int32(-49))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							if v63 == int32(0) {
								v67 = F_attribute_reloptions(m, v61, int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[1]))
									v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
									v74 = F_MemoryContextAlloc(m, v70, int32(base.Ui32(v71)>>(uint(int32(2))%32)))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
										v78 = int32(base.Ui32(v76) >> (uint(int32(2)) % 32))
										if v78 != 0 {
											v79 = F__emscripten_memcpy_bulkmem(m, v74, v67, v78)
											mBase = m.M
										} else {
										}
										v82 = v74
										F_ReleaseCatCache(m, v55)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											v86 = v82
											v88 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
											v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
												v97 = v86
												v98 = v93
												if v97 == int32(0) {
													v113 = int32(0)
													m.G0 = v8 - int32(-64)
													return v113
												} else {
													v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
													v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return int32(0)
													} else {
														v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
														v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
														if v110 != 0 {
															v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
															mBase = m.M
														} else {
														}
														v113 = v105
														m.G0 = v8 - int32(-64)
														return v113
													}
												}
											}
										}
									}
								}
							} else {
								v82 = v52
								F_ReleaseCatCache(m, v55)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									v86 = v82
									v88 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
									v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
										v97 = v86
										v98 = v93
										if v97 == int32(0) {
											v113 = int32(0)
											m.G0 = v8 - int32(-64)
											return v113
										} else {
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
											v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return int32(0)
											} else {
												v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
												v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
												if v110 != 0 {
													v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
													mBase = m.M
												} else {
												}
												v113 = v105
												m.G0 = v8 - int32(-64)
												return v113
											}
										}
									}
								}
							}
						}
					} else {
						v86 = v52
						v88 = *(*int32)(unsafe.Add(mBase, _c_F_get_attribute_options[0]))
						v93 = F_hash_search(m, v88, v6+int32(-48), int32(1), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v86
							v97 = v86
							v98 = v93
							if v97 == int32(0) {
								v113 = int32(0)
								m.G0 = v8 - int32(-64)
								return v113
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
								v105 = F_palloc(m, int32(base.Ui32(v102)>>(uint(int32(2))%32)))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
									v110 = int32(base.Ui32(v108) >> (uint(int32(2)) % 32))
									if v110 != 0 {
										v111 = F__emscripten_memcpy_bulkmem(m, v105, v107, v110)
										mBase = m.M
									} else {
									}
									v113 = v105
									m.G0 = v8 - int32(-64)
									return v113
								}
							}
						}
					}
				}
			}
		}
	}
}
