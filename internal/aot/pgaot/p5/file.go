package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FileSetDelete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	v5 = m.G0
	v7 = v5 - int32(2080)
	m.G0 = v7
	if l1&int32(3) == int32(0) {
		v34 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v73 = v67 - int32(1636608432)
	if l1&int32(3) != 0 {
		goto L22
	} else {
		goto L23
	}
L2:
	;
	v67 = v59 - l1
	goto L1
L3:
	;
	v38 = v34
	goto L12
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v18 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v67 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v23 = l1
	goto L8
L8:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v59 = v27
	goto L2
L10:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v53 = v38
	goto L15
L14:
	;
	goto L13
L15:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v59 = v53
	goto L2
L17:
	;
	goto L16
L18:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v333 = base.I32_rem_u_s(v327^v319-base.I32_rotl(v327, int32(24)), v332)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0+v333<<(uint(int32(2))%32))+12))
	F_TempTablespacePath(m, v7+int32(1056), v337)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L58
	} else {
		goto L59
	}
L19:
	;
	v305 = int32(14)
	v307 = v301 ^ v302 - base.I32_rotl(v301, v305)
	v311 = v307 ^ v300 - base.I32_rotl(v307, int32(11))
	v315 = v311 ^ v301 - base.I32_rotl(v311, int32(25))
	v319 = v315 ^ v307 - base.I32_rotl(v315, int32(16))
	v323 = v319 ^ v311 - base.I32_rotl(v319, int32(4))
	v327 = v323 ^ v315 - base.I32_rotl(v323, v305)
	goto L18
L20:
	;
	switch v231 - int32(1) {
	case 0:
		v293 = v232
		v294 = v233
		v295 = v234
		goto L47
	case 1:
		v286 = v232
		v287 = v233
		v288 = v234
		goto L48
	case 2:
		v279 = v232
		v280 = v233
		v281 = v234
		goto L49
	case 3:
		v273 = v233
		v274 = v234
		goto L50
	case 4:
		v269 = v233
		v270 = v234
		goto L51
	case 5:
		v263 = v233
		v264 = v234
		goto L52
	case 6:
		v257 = v233
		v258 = v234
		goto L53
	case 7:
		v252 = v234
		goto L54
	case 8:
		v247 = v234
		goto L55
	case 9:
		v242 = v234
		goto L56
	case 10:
		goto L57
	default:
		v300 = v232
		v301 = v233
		v302 = v234
		goto L19
	}
L21:
	;
	v182 = l1
	v183 = v67
	v184 = v73
	v185 = v73
	v186 = v73
	goto L44
L22:
	;
	if base.Ui32(int32(11)) < base.Ui32(v67) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v67) < base.Ui32(int32(12)) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v230 = l1
	v231 = v67
	v232 = v73
	v233 = v73
	v234 = v73
	goto L20
L26:
	;
	switch v129 - int32(1) {
	case 0:
		v179 = v130
		goto L33
	case 1:
		v174 = v130
		goto L34
	case 2:
		goto L35
	case 3:
		v167 = v131
		goto L36
	case 4:
		v164 = v131
		goto L37
	case 5:
		v159 = v131
		goto L38
	case 6:
		goto L39
	case 7:
		v150 = v132
		goto L40
	case 8:
		v145 = v132
		goto L41
	case 9:
		v140 = v132
		goto L42
	case 10:
		goto L43
	default:
		v300 = v130
		v301 = v131
		v302 = v132
		goto L19
	}
L27:
	;
	v128 = l1
	v129 = v67
	v130 = v73
	v131 = v73
	v132 = v73
	goto L26
L28:
	;
	goto L29
L29:
	;
	v80 = l1
	v81 = v67
	v82 = v73
	v83 = v73
	v84 = v73
	goto L30
L30:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v87 = v86 + v83
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	v91 = v90 + v84
	v93 = int32(4)
	v95 = v88 + v82 - v91 ^ base.I32_rotl(v91, v93)
	v99 = v87 - v95 ^ base.I32_rotl(v95, int32(6))
	v100 = v91 + v87
	v101 = v95 + v100
	v102 = v99 + v101
	v106 = v100 - v99 ^ base.I32_rotl(v99, int32(8))
	v110 = v101 - v106 ^ base.I32_rotl(v106, int32(16))
	v114 = v102 - v110 ^ base.I32_rotl(v110, int32(19))
	v115 = v106 + v102
	v116 = v110 + v115
	v117 = v114 + v116
	v121 = v115 - v114 ^ base.I32_rotl(v114, v93)
	v122 = int32(12)
	v123 = v80 + v122
	v125 = v81 - v122
	if base.Ui32(int32(11)) < base.Ui32(v125) {
		v80 = v123
		v81 = v125
		v82 = v116
		v83 = v117
		v84 = v121
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v128 = v123
	v129 = v125
	v130 = v116
	v131 = v117
	v132 = v121
	goto L26
L32:
	;
	goto L31
L33:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	v300 = v179 + v180
	v301 = v131
	v302 = v132
	goto L19
L34:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)))
	v179 = v175<<(uint(int32(8))%32) + v174
	goto L33
L35:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+2)))
	v174 = v170<<(uint(int32(16))%32) + v130
	goto L34
L36:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v300 = v168 + v130
	v301 = v167
	v302 = v132
	goto L19
L37:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+4)))
	v167 = v164 + v165
	goto L36
L38:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+5)))
	v164 = v160<<(uint(int32(8))%32) + v159
	goto L37
L39:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+6)))
	v159 = v155<<(uint(int32(16))%32) + v131
	goto L38
L40:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v300 = v151 + v130
	v301 = v153 + v131
	v302 = v150
	goto L19
L41:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+8)))
	v150 = v146<<(uint(int32(8))%32) + v145
	goto L40
L42:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+9)))
	v145 = v141<<(uint(int32(16))%32) + v140
	goto L41
L43:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+10)))
	v140 = v136<<(uint(int32(24))%32) + v132
	goto L42
L44:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v189 = v188 + v185
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v182)+8))
	v193 = v192 + v186
	v195 = int32(4)
	v197 = v190 + v184 - v193 ^ base.I32_rotl(v193, v195)
	v201 = v189 - v197 ^ base.I32_rotl(v197, int32(6))
	v202 = v193 + v189
	v203 = v197 + v202
	v204 = v201 + v203
	v208 = v202 - v201 ^ base.I32_rotl(v201, int32(8))
	v212 = v203 - v208 ^ base.I32_rotl(v208, int32(16))
	v216 = v204 - v212 ^ base.I32_rotl(v212, int32(19))
	v217 = v208 + v204
	v218 = v212 + v217
	v219 = v216 + v218
	v223 = v217 - v216 ^ base.I32_rotl(v216, v195)
	v224 = int32(12)
	v225 = v182 + v224
	v227 = v183 - v224
	if base.Ui32(int32(11)) < base.Ui32(v227) {
		v182 = v225
		v183 = v227
		v184 = v218
		v185 = v219
		v186 = v223
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v230 = v225
	v231 = v227
	v232 = v218
	v233 = v219
	v234 = v223
	goto L20
L46:
	;
	goto L45
L47:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v300 = v293 + v296
	v301 = v294
	v302 = v295
	goto L19
L48:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	v293 = v289<<(uint(int32(8))%32) + v286
	v294 = v287
	v295 = v288
	goto L47
L49:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+2)))
	v286 = v282<<(uint(int32(16))%32) + v279
	v287 = v280
	v288 = v281
	goto L48
L50:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+3)))
	v279 = v275<<(uint(int32(24))%32) + v232
	v280 = v273
	v281 = v274
	goto L49
L51:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+4)))
	v273 = v269 + v271
	v274 = v270
	goto L50
L52:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+5)))
	v269 = v265<<(uint(int32(8))%32) + v263
	v270 = v264
	goto L51
L53:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+6)))
	v263 = v259<<(uint(int32(16))%32) + v257
	v264 = v258
	goto L52
L54:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+7)))
	v257 = v253<<(uint(int32(24))%32) + v233
	v258 = v252
	goto L53
L55:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+8)))
	v252 = v248<<(uint(int32(8))%32) + v247
	goto L54
L56:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+9)))
	v247 = v243<<(uint(int32(16))%32) + v242
	goto L55
L57:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+10)))
	v242 = v238<<(uint(int32(24))%32) + v234
	goto L56
L58:
	;
	return int32(0)
L59:
	;
	v342 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(224323)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v7 + int32(1056)
	v355 = F_pg_snprintf(m, v7+int32(32), int32(1024), int32(98965), v7+int32(16))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(32)
	v365 = F_pg_snprintf(m, v7+int32(1056), int32(1024), int32(167509), v7)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v370 = F_PathNameDeleteTemporaryFile(m, v7+int32(1056), int32(1))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L58
	} else {
		goto L62
	}
L62:
	;
	m.G0 = v7 + int32(2080)
	return v370
}
func F_FileWriteV(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int64
	_ = v52
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int64
	_ = v82
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int64
	_ = v100
	var v105 int32
	_ = v105
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v124 int64
	_ = v124
	var v133 int64
	_ = v133
	var v136 int64
	_ = v136
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	v8 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = F_FileAccess(m, l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L3
	} else {
		goto L37
	}
L2:
	;
	m.G0 = v17 + int32(16)
	return v213
L3:
	;
	return int32(0)
L4:
	;
	if v19 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v213 = int32(-1)
	goto L2
L6:
	;
	goto L7
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[552]))
	v30 = v27 + l0*int32(48)
	v32 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	if v32 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	goto L25
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
	if v35&int32(4) == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if l2 <= int32(0) {
		v124 = l3
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v30)+24))
	if v124 <= v133 {
		goto L8
	} else {
		goto L23
	}
L12:
	;
	v42 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = v42
	v52 = l3
	v60 = v8
	goto L16
L14:
	;
	v77 = v42
	v82 = l3
	goto L15
L15:
	;
	v92 = l2 & int32(3)
	if v92 == int32(0) {
		v124 = v82
		goto L11
	} else {
		goto L19
	}
L16:
	;
	v63 = l1 + v47<<(uint(int32(3))%32)
	v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v63)+4)))
	v66 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v63)+12)))
	v68 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v63)+20)))
	v70 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v63)+28)))
	v71 = v52 + v64 + v66 + v68 + v70
	v72 = int32(4)
	v73 = v47 + v72
	v75 = v60 + v72
	if v75 != l2&int32(2147483644) {
		v47 = v73
		v52 = v71
		v60 = v75
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v77 = v73
	v82 = v71
	goto L15
L18:
	;
	goto L17
L19:
	;
	v95 = v77
	v100 = v82
	v105 = v8
	goto L20
L20:
	;
	v112 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1+v95<<(uint(int32(3))%32))+4)))
	v113 = v100 + v112
	v114 = int32(1)
	v117 = v105 + v114
	if v117 != v92 {
		v95 = v95 + v114
		v100 = v113
		v105 = v117
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v124 = v113
	goto L11
L22:
	;
	goto L21
L23:
	;
	v136 = *(*int64)(unsafe.Add(mBase, _consts[653]))
	if base.Ui64(base.I64_extend_i32_u(v32)<<(uint(int64(10))%64)) < base.Ui64(v136+(v124-v133)) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L8
L25:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[96]))
	*(*int32)(unsafe.Add(mBase, uint32(v174))) = l4
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if base.B2i32(l2 != int32(1)) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v213 = v183
	goto L2
L27:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _consts[96]))
	v186 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = v186
	if v186 <= v183 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v181 = F_pwrite(m, v176, v179, v180, l3)
	mBase = m.M
	v183 = v181
	goto L27
L29:
	;
	goto L30
L30:
	;
	v182 = F_pwritev(m, v176, l1, l2, l3)
	mBase = m.M
	v183 = v182
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(51)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
	if v193&int32(4) == int32(0) {
		v213 = v183
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v210 == int32(27) {
		goto L25
	} else {
		goto L36
	}
L34:
	;
	v199 = l3 + base.I64_extend_i32_u(v183)
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v30)+24))
	if v199 <= v200 {
		v213 = v183
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v202 = int32(4359864)
	v204 = *(*int64)(unsafe.Add(mBase, _consts[653]))
	*(*int64)(unsafe.Add(mBase, _consts[653])) = v204 + (v199 - v200)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v199
	v213 = v183
	goto L2
L36:
	;
	goto L26
L37:
	;
	F_errcode(m, int32(16581))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _consts[652]))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v239
	F_errmsg(m, int32(642466), v17)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(478009), int32(2295), int32(494257))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
