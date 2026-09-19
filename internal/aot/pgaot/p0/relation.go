package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RelationCacheInitializePhase3(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[0])))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[1]))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[2]))
	F_read_relmap_file(m, int32(_a_F_RelationCacheInitializePhase3_0), v17, int32(0), int32(22))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v22 = int32(_a_F_RelationCacheInitializePhase3_1)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[3]))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[3])) = v26
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[1]))
	if v29 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[3])) = v23
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[1]))
	if v72 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	F_formrdesc(m, int32(_a_F_RelationCacheInitializePhase3_2), int32(83), int32(0), int32(34), int32(_a_F_RelationCacheInitializePhase3_3))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L11
	}
L8:
	;
	v33 = F_load_relcache_init_file(m, int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v33 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v68 = v12 ^ int32(1)
	goto L6
L11:
	;
	F_formrdesc(m, int32(_a_F_RelationCacheInitializePhase3_4), int32(75), int32(0), int32(25), int32(_a_F_RelationCacheInitializePhase3_5))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_formrdesc(m, int32(_a_F_RelationCacheInitializePhase3_6), int32(81), int32(0), int32(30), int32(_a_F_RelationCacheInitializePhase3_7))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_formrdesc(m, int32(_a_F_RelationCacheInitializePhase3_8), int32(71), int32(0), int32(32), int32(_a_F_RelationCacheInitializePhase3_9))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v68 = int32(1)
	goto L6
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L4
	} else {
		goto L104
	}
L16:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L100
	}
L17:
	;
	m.G0 = v9 + int32(48)
	return
L18:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[5])))
	if v76 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_load_critical_index(m, int32(2662), int32(1259))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[0])))
	if v111 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	F_load_critical_index(m, int32(2659), int32(1249))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_load_critical_index(m, int32(2679), int32(2610))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_load_critical_index(m, int32(2687), int32(2616))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_load_critical_index(m, int32(2655), int32(2603))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_load_critical_index(m, int32(2693), int32(2618))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	F_load_critical_index(m, int32(2701), int32(2620))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[5])) = uint8(v108)
	goto L21
L29:
	;
	F_load_critical_index(m, int32(2671), int32(1262))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v142 = v9 + int32(28)
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[6]))
	F_hash_seq_init(m, v142, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L38
	}
L32:
	;
	F_load_critical_index(m, int32(2672), int32(1262))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_load_critical_index(m, int32(2676), int32(1260))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	F_load_critical_index(m, int32(2677), int32(1260))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_load_critical_index(m, int32(2695), int32(1261))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_load_critical_index(m, int32(3593), int32(3592))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v139 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[0])) = uint8(v139)
	goto L31
L38:
	;
	v147 = F_hash_seq_search(m, v142)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v147 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v149 = v147
	goto L43
L41:
	;
	goto L42
L42:
	;
	if v68&int32(1) == int32(0) {
		goto L17
	} else {
		goto L93
	}
L43:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[7]))
	F_ResourceOwnerEnlarge(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+16)) = v160 + int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[1]))
	if v165 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[7]))
	F_ResourceOwnerRemember(m, v167, v155, int32(_a_F_RelationCacheInitializePhase3_10))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v155)+48))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+80))
	if v172 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L48
L50:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v155)+56))
	v177 = F_SearchSysCache1(m, int32(57), v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	v198 = v171
	goto L52
L52:
	;
	v201 = base.B2i32(v172 == int32(0))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+124)))
	if v202 != int32(1) {
		v214 = v198
		v215 = v201
		goto L62
	} else {
		goto L63
	}
L53:
	;
	if v177 == int32(0) {
		goto L16
	} else {
		goto L54
	}
L54:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v155)+48))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+22)))
	base.MemoryCopy(m, v181, v182+v183, int32(144))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v155)+180))
	if v187 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_pfree(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	F_RelationParseRelOptions(m, v155, v177)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	F_ReleaseCatCache(m, v177)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v155)+48))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+80))
	if v195 == int32(0) {
		goto L15
	} else {
		goto L61
	}
L61:
	;
	v198 = v194
	goto L52
L62:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+125)))
	if v216 != int32(1) {
		v228 = v214
		v229 = v215
		goto L67
	} else {
		goto L68
	}
L63:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v155)+68))
	if v205 != 0 {
		v214 = v198
		v215 = v201
		goto L62
	} else {
		goto L64
	}
L64:
	;
	F_RelationBuildRuleLock(m, v155)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v155)+48))
	v209 = int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v155)+68))
	if v210 != 0 {
		v214 = v208
		v215 = v209
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v211 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+124)) = uint8(v211)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v155)+48))
	v214 = v213
	v215 = v209
	goto L62
L67:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+127)))
	if v230 != int32(1) {
		v237 = v229
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v155)+76))
	if v219 != 0 {
		v228 = v214
		v229 = v215
		goto L67
	} else {
		goto L69
	}
L69:
	;
	F_RelationBuildTriggers(m, v155)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v155)+48))
	v223 = int32(1)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v155)+76))
	if v224 != 0 {
		v228 = v222
		v229 = v223
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v222)+125)) = uint8(v225)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v155)+48))
	v228 = v227
	v229 = v223
	goto L67
L72:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v155)+188))
	if v238 != 0 {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v155)+80))
	if v233 != 0 {
		v237 = v229
		goto L72
	} else {
		goto L74
	}
L74:
	;
	F_RelationBuildRowSecurity(m, v155)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	v237 = int32(1)
	goto L72
L76:
	;
	v282 = F_hash_seq_search(m, v9+int32(28))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L91
	}
L77:
	;
	v272 = v9 + int32(28)
	F_hash_seq_term(m, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L89
	}
L78:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+16)) = v258 - int32(1)
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[1]))
	if v263 != 0 {
		goto L84
	} else {
		goto L85
	}
L79:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v155)+48))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+119)))
	switch v240 - int32(83) {
	case 0, 26, 31, 33:
		goto L80
	default:
		goto L78
	}
L80:
	;
	F_RelationInitTableAccessMethod(m, v155)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+16)) = v245 - int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[1]))
	if v250 == int32(0) {
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[7]))
	F_ResourceOwnerForget(m, v254, v155, int32(_a_F_RelationCacheInitializePhase3_10))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	goto L77
L84:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[7]))
	F_ResourceOwnerForget(m, v265, v155, int32(_a_F_RelationCacheInitializePhase3_10))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if v237 == int32(0) {
		goto L76
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	goto L77
L89:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[6]))
	F_hash_seq_init(m, v272, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	goto L76
L91:
	;
	if v282 != 0 {
		v149 = v282
		goto L43
	} else {
		goto L92
	}
L92:
	;
	goto L44
L93:
	;
	v295 = int32(0)
	goto L94
L94:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v295<<(uint(int32(2))%32))+uint32(_c_F_RelationCacheInitializePhase3[8])))
	F_InitCatCachePhase2(m, v303, int32(1))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L4
	} else {
		goto L96
	}
L95:
	;
	F_write_relcache_init_file(m, int32(1))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L98
	}
L96:
	;
	v308 = v295 + int32(1)
	if v308 != int32(85) {
		v295 = v308
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	F_write_relcache_init_file(m, int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	goto L17
L100:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v155)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v333
	F_errmsg_internal(m, int32(_a_F_RelationCacheInitializePhase3_11), v9)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_RelationCacheInitializePhase3_12), int32(_a_F_RelationCacheInitializePhase3_13), int32(_a_F_RelationCacheInitializePhase3_14))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v155)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v347 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationCacheInitializePhase3_15), v9+int32(16))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_RelationCacheInitializePhase3_12), int32(_a_F_RelationCacheInitializePhase3_16), int32(_a_F_RelationCacheInitializePhase3_14))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationCopyStorage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
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
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorage[0]))
	v26 = (base.B2i32(l2 == int32(3))&base.B2i32(l3 == int32(117)) | base.B2i32(l3 == int32(112))) & base.B2i32(v5 < v23)
	v28 = F_palloc(m, int32(424))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)) = uint8(v26)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = l1
	v35 = F_smgrnblocks(m, l1, l2)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+400)) = v35
	v38 = F_GetRedoRecPtr(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+408)) = v38
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorage[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+416)) = v42
	v44 = F_smgrnblocks(m, l0, l2)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v108 = v12 + int32(20)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v108, v109, v110, v111, v112, l2)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L31
	}
L6:
	;
	if v44 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v51 = v5
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_smgr_bulk_finish(m, v28)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L30
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorage[2]))
	if v56 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v59 = F_smgr_bulk_get_buf(m, v28)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v59
	F_smgrreadv(m, l0, l2, v51, v12+int32(20))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationCopyStorage[3])))
	if v69 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = int32(5)
	goto L20
L19:
	;
	v70 = int32(1)
	goto L20
L20:
	;
	v73 = F_PageIsVerified(m, v59, v51, v70, v12+int32(95))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+95)))
	if v75 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_pgstat_prepare_report_checksum_failure(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v73 == int32(0) {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	F_pgstat_report_checksum_failures_in_db(m, v78, int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	F_smgr_bulk_write(m, v28, v51, v59, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v91 = v51 + int32(1)
	if v91 != v44 {
		v51 = v91
		goto L10
	} else {
		goto L29
	}
L29:
	;
	goto L11
L30:
	;
	m.G0 = v12 + int32(96)
	return
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v108
	F_errmsg(m, int32(_a_F_RelationCopyStorage_0), v12)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_RelationCopyStorage_1), int32(550), int32(_a_F_RelationCopyStorage_2))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationCopyStorageUsingBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	v5 = int32(0)
	v16 = m.G0
	v20 = (v16 - int32(_a_F_RelationCopyStorageUsingBuffer_0)) & int32(-4096)
	m.G0 = v20
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[0]))
	if v5 < v27 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = l3 | base.B2i32(l2 == int32(3))
	goto L3
L2:
	;
	v30 = v5
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4080)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+4072)) = v33
	v38 = F_smgropen(m, v20+int32(4072), int32(-1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v40 = F_smgrnblocks(m, v38, l2)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v43 = v20 + int32(_a_F_RelationCopyStorageUsingBuffer_1)
	base.MemoryFill(m, v43, int32(0), int32(_a_F_RelationCopyStorageUsingBuffer_2))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4064)) = v47
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+4056)) = v49
	v54 = F_smgropen(m, v20+int32(4056), int32(-1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	m.G0 = v16
	return
L10:
	;
	v56 = int32(1)
	F_smgrextend(m, v54, l2, v40-v56, v43, v56)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v62 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v65 = F_GetAccessStrategy(m, int32(2))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4092)) = v40
	v68 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4088)) = v68
	v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+4040)) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4048)) = v72
	v79 = F_smgropen(m, v20+int32(4040), int32(-1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	if l3 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v83 = int32(112)
	goto L17
L16:
	;
	v83 = int32(117)
	goto L17
L17:
	;
	v88 = F_read_stream_begin_impl(m, int32(12), v62, v68, v79, v83, l2, int32(120), v20+int32(4088), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v104 = v5
	goto L19
L19:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[1]))
	if v106 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_read_stream_end(m, v88)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L52
	}
L21:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v110 = F_read_stream_next_buffer(m, v88, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)+16))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4032)) = v153
	v155 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+4024)) = v155
	v160 = F_ReadBufferWithoutRelcache(m, v20+int32(4024), l2, v152, int32(1), v65, l3)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L32
	}
L26:
	;
	if v110 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[2]))
	v117 = v110 ^ int32(-1)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[3]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v117<<(uint(int32(2))%32))))
	v149 = v115 + v117<<(uint(int32(6))%32)
	v151 = v126
	goto L25
L28:
	;
	goto L29
L29:
	;
	v128 = v110 << (uint(int32(6)) % 32)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[4]))
	v135 = F_LWLockAcquire(m, v128+v130-int32(16), int32(1))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[4]))
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[5]))
	v149 = v138 + v128 + int32(-64)
	v151 = v143 + v110<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L31:
	;
	v180 = int32(_a_F_RelationCopyStorageUsingBuffer_3)
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[6])) = v182 + int32(1)
	base.MemoryCopy(m, v179, v151, int32(_a_F_RelationCopyStorageUsingBuffer_2))
	F_MarkBufferDirty(m, v160)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L36
	}
L32:
	;
	if v160 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[3]))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165+(v160^int32(-1))<<(uint(int32(2))%32))))
	v179 = v171
	goto L31
L34:
	;
	goto L35
L35:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[5]))
	v179 = v173 + v160<<(uint(int32(13))%32) + int32(-8192)
	goto L31
L36:
	;
	if v30 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_log_newpage_buffer(m, v160, int32(1))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v193 = int32(_a_F_RelationCopyStorageUsingBuffer_3)
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[6])) = v195 - int32(1)
	if int32(0) <= v160 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[4]))
	F_LWLockRelease(m, v202+v160<<(uint(int32(6))%32)-int32(16))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_ReleaseBuffer(m, v160)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	if int32(0) <= v110 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[4]))
	F_LWLockRelease(m, v215+v110<<(uint(int32(6))%32)-int32(16))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_ReleaseBuffer(m, v110)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	v226 = v104 + int32(1)
	if v226 != v40 {
		v104 = v226
		goto L19
	} else {
		goto L51
	}
L51:
	;
	goto L20
L52:
	;
	F_bms_free(m, v62)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_bms_free(m, v65)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	goto L9
}
func F_RelationDropStorage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[0]))
	v7 = F_MemoryContextAlloc(m, v5, int32(28))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v9
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v14)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v13
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[1]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v19
		v21 = int32(_a_F_RelationDropStorage_0)
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[2]))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v22
		*(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[2])) = v7
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v26 != 0 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
			v31 = v29 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v31
			if v31 == int32(0) {
				v36 = v26 + int32(76)
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[3]))
				if v38 != 0 {
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[4]))
					v45 = v40
				} else {
					v42 = int32(_a_F_RelationDropStorage_1)
					*(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[3])) = v42
					v45 = v42
				}
				*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v45
				v47 = int32(_a_F_RelationDropStorage_1)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v47
				*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v36
				*(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[4])) = v36
			} else {
			}
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_smgrclose(m, v54)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
				return
			}
		} else {
			return
		}
	}
}
func F_RelationGetExclusionInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	v15 = m.G0
	v17 = v15 - int32(112)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+10)))
	v22 = v20 << (uint(int32(2)) % 32)
	v23 = F_palloc(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v23
	v26 = F_palloc(m, v22)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v26
	v30 = v20 << (uint(int32(1)) % 32)
	v31 = F_palloc(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v31
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L85
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L82
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L79
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L76
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L73
	}
L10:
	;
	F_systable_endscan(m, v66)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L71
	}
L11:
	;
	m.G0 = v17 + int32(112)
	return
L12:
	;
	v35 = int32(0)
	v36 = base.B2i32(v22 == v35)
	if v36 == v35 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v50 = v17 - int32(-64)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	F_ScanKeyInit(m, v50, int32(9), int32(3), int32(184), v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	base.MemoryCopy(m, v23, v39, v22)
	goto L17
L16:
	;
	goto L17
L17:
	;
	if v36 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	base.MemoryCopy(m, v26, v43, v22)
	goto L20
L19:
	;
	goto L20
L20:
	;
	if v30 == int32(0) {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	base.MemoryCopy(m, v31, v47, v30)
	goto L11
L22:
	;
	v60 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v63 = int32(1)
	v66 = F_systable_beginscan(m, v60, int32(2665), v63, int32(0), v63, v50)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v68 = F_systable_getnext(m, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v68 == int32(0) {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v73 = v68
	v84 = int32(0)
	goto L27
L27:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+22)))
	v88 = v86 + v87
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+72)))
	if v89 == int32(120) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	F_systable_endscan(m, v66)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L47
	}
L29:
	;
	v126 = F_systable_getnext(m, v66)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L45
	}
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v88)+88))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v97 != v98 {
		v125 = v84
		goto L29
	} else {
		goto L33
	}
L31:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+107)))
	if v92 != int32(1) {
		v125 = v84
		goto L29
	} else {
		goto L32
	}
L32:
	;
	switch v89 - int32(112) {
	case 0, 5:
		goto L30
	default:
		v125 = v84
		goto L29
	}
L33:
	;
	if v84 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
	v104 = F_fastgetattr_3(m, v73, int32(27), v101, v17+int32(63))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+63)))
	if v106 == int32(1) {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v109 = F_pg_detoast_datum(m, v104)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v111 != int32(1) {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)+16))
	if v114 != v20 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	if v116 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	if v117 != int32(26) {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	if v22 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	base.MemoryCopy(m, v23, v109+int32(24), v22)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v125 = int32(1)
	goto L29
L45:
	;
	if v126 != 0 {
		v73 = v126
		v84 = v125
		goto L27
	} else {
		goto L46
	}
L46:
	;
	goto L28
L47:
	;
	F_relation_close(m, v60, int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v125 == int32(0) {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	if int32(0) < v20 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v139 = int32(0)
	goto L53
L51:
	;
	goto L52
L52:
	;
	v191 = int32(_a_F_RelationGetExclusionInfo_0)
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetExclusionInfo[0]))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetExclusionInfo[0])) = v194
	v196 = F_palloc(m, v22)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L59
	}
L53:
	;
	v153 = v139 << (uint(int32(2)) % 32)
	v155 = v153 + v23
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v157 = F_get_opcode(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+v153))) = v157
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v164+v153)))
	v167 = F_get_op_opfamily_strategy(m, v163, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v31+v139<<(uint(int32(1))%32)))) = uint16(v167)
	if v167&int32(_a_F_RelationGetExclusionInfo_1) == int32(0) {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v175 = v139 + int32(1)
	if v175 != v20 {
		v139 = v175
		goto L53
	} else {
		goto L58
	}
L58:
	;
	goto L54
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+236)) = v196
	v199 = F_palloc(m, v22)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v199
	v202 = F_palloc(m, v30)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+244)) = v202
	v205 = int32(0)
	v206 = base.B2i32(v22 == v205)
	if v206 == v205 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	base.MemoryCopy(m, v209, v23, v22)
	goto L64
L63:
	;
	goto L64
L64:
	;
	if v206 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	base.MemoryCopy(m, v213, v26, v22)
	goto L67
L66:
	;
	goto L67
L67:
	;
	if v30 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	base.MemoryCopy(m, v215, v31, v30)
	goto L70
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetExclusionInfo[0])) = v192
	goto L11
L71:
	;
	F_relation_close(m, v60, int32(1))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L9
L73:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v259 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationGetExclusionInfo_2), v17)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_RelationGetExclusionInfo_3), int32(_a_F_RelationGetExclusionInfo_4), int32(_a_F_RelationGetExclusionInfo_5))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v275 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationGetExclusionInfo_6), v17+int32(16))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_RelationGetExclusionInfo_3), int32(_a_F_RelationGetExclusionInfo_7), int32(_a_F_RelationGetExclusionInfo_5))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v293 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationGetExclusionInfo_8), v17+int32(32))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_RelationGetExclusionInfo_3), int32(_a_F_RelationGetExclusionInfo_9), int32(_a_F_RelationGetExclusionInfo_5))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errmsg_internal(m, int32(_a_F_RelationGetExclusionInfo_10), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_RelationGetExclusionInfo_3), int32(_a_F_RelationGetExclusionInfo_11), int32(_a_F_RelationGetExclusionInfo_5))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v325+v139<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v324
	F_errmsg_internal(m, int32(_a_F_RelationGetExclusionInfo_12), v17+int32(48))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_RelationGetExclusionInfo_3), int32(_a_F_RelationGetExclusionInfo_13), int32(_a_F_RelationGetExclusionInfo_5))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationGetReplicaIndex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v2 == int32(0) {
		v5 = F_RelationGetIndexList(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			F_list_free(m, v5)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
				return v11
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
		return v11
	}
}
func F_RelationInitPhysicalAddr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)))
	switch v11 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L4
	default:
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L29
	} else {
		goto L77
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L29
	} else {
		goto L74
	}
L3:
	;
	m.G0 = v8 + int32(48)
	return
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[0]))
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v18 = v15
	goto L7
L6:
	;
	v18 = v17
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[1]))
	if v18 != int32(1664) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = v21
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
	if v27 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[2]))
	if base.B2i32(v197 == v14)|base.B2i32(v201 < int32(0)) != 0 {
		goto L3
	} else {
		goto L72
	}
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[3]))
	goto L16
L13:
	;
	goto L14
L14:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+117)))
	v91 = int32(0)
	if v90 == v91 {
		goto L36
	} else {
		goto L37
	}
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
	v197 = v87
	goto L11
L16:
	;
	if base.B2i32(v29 != int32(0)) == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[4]))
	if v35 < int32(2) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+118)))
	if v39 != int32(112) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L20
L20:
	;
	if base.B2i32(base.Ui32(v42) < base.Ui32(int32(_a_F_RelationInitPhysicalAddr_0))) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v47 == int32(0) {
		goto L15
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[5]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	goto L27
L24:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+119)))
	switch v51 - int32(109) {
	case 0, 5:
		goto L25
	default:
		goto L15
	}
L25:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+104)))
	if v54 != int32(1) {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	if base.B2i32(v60 == int32(2)) == int32(0) {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v69 = F_ScanPgRelation(m, v65, base.B2i32(v65 != int32(2662)), int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return
L30:
	;
	if v69 == int32(0) {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+22)))
	v76 = v74 + v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+92)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+88)) = v80
	F_pfree(m, v69)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L15
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v193
	if v193 == int32(0) {
		goto L1
	} else {
		goto L71
	}
L34:
	;
	goto L33
L35:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v193 = v188
	goto L34
L36:
	;
	v96 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[6]))
	if v96 < v98 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v139 = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[7]))
	if v139 < v141 {
		goto L57
	} else {
		goto L58
	}
L39:
	;
	v102 = v96
	goto L42
L40:
	;
	goto L41
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[8]))
	if v121 <= int32(0) {
		v193 = v91
		goto L34
	} else {
		goto L48
	}
L42:
	;
	v107 = v102 << (uint(int32(3)) % 32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+uint32(_c_F_RelationInitPhysicalAddr[9])))
	if v108 == v89 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L41
L44:
	;
	v187 = v107 + int32(_a_F_RelationInitPhysicalAddr_1)
	goto L35
L45:
	;
	goto L46
L46:
	;
	v113 = v102 + int32(1)
	if v113 != v98 {
		v102 = v113
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	v126 = int32(0)
	goto L49
L49:
	;
	v131 = v126 << (uint(int32(3)) % 32)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+uint32(_c_F_RelationInitPhysicalAddr[10])))
	if v132 != v89 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v187 = v131 + int32(_a_F_RelationInitPhysicalAddr_2)
	goto L35
L51:
	;
	v135 = v126 + int32(1)
	if v121 != v135 {
		v126 = v135
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	goto L50
L54:
	;
	v193 = v91
	goto L34
L55:
	;
	v169 = int32(0)
	goto L65
L56:
	;
	v187 = v150 + int32(_a_F_RelationInitPhysicalAddr_3)
	goto L35
L57:
	;
	v145 = v139
	goto L60
L58:
	;
	goto L59
L59:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[11]))
	if v162 <= int32(0) {
		v193 = v91
		goto L34
	} else {
		goto L64
	}
L60:
	;
	v150 = v145 << (uint(int32(3)) % 32)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+uint32(_c_F_RelationInitPhysicalAddr[12])))
	if v89 == v151 {
		goto L56
	} else {
		goto L62
	}
L61:
	;
	goto L59
L62:
	;
	v154 = v145 + int32(1)
	if v154 != v141 {
		v145 = v154
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L55
L65:
	;
	v174 = v169 << (uint(int32(3)) % 32)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+uint32(_c_F_RelationInitPhysicalAddr[13])))
	if v175 != v89 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v187 = v174 + int32(_a_F_RelationInitPhysicalAddr_4)
	goto L35
L67:
	;
	v178 = v169 + int32(1)
	if v162 != v178 {
		v169 = v178
		goto L65
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	goto L66
L70:
	;
	v193 = v91
	goto L34
L71:
	;
	v197 = v193
	goto L11
L72:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v205
	v207 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v207
	v211 = F_RelFileLocatorSkippingWAL(m, v8+int32(16))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L29
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v211
	goto L3
L74:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v224
	F_errmsg_internal(m, int32(_a_F_RelationInitPhysicalAddr_5), v8+int32(32))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L29
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_RelationInitPhysicalAddr_6), int32(1380), int32(_a_F_RelationInitPhysicalAddr_7))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L29
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v240 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationInitPhysicalAddr_8), v8)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L29
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_RelationInitPhysicalAddr_6), int32(1398), int32(_a_F_RelationInitPhysicalAddr_7))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L29
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationTruncate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int64
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int64
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int64
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
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
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(144)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v15 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v19
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v21
	v25 = F_smgropen(m, v13+int32(72), v18)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v42 = v15
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+32)) = int32(-1)
	v45 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v45
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v45
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+132)) = v49
	v52 = F_smgrnblocks(m, v42, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L10
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v25
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+72))
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = v41
	goto L3
L7:
	;
	v37 = v29
	goto L9
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+72))
	v37 = v35
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+72)) = v37 + int32(1)
	goto L6
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+120)) = v52
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v56 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v60
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v62
	v66 = F_smgropen(m, v13+int32(56), v59)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	v83 = v56
	goto L13
L13:
	;
	v86 = v13 + int32(136)
	v88 = v13 + int32(124)
	v90 = v13 + int32(112)
	v91 = int32(1)
	v93 = F_smgrexists(m, v83, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L20
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v66
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
	if v70 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v83 = v82
	goto L13
L16:
	;
	v78 = v70
	goto L18
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)+76))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v66)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
	v78 = v76
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+72)) = v78 + int32(1)
	goto L15
L19:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v122 != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	if v93 == int32(0) {
		v117 = v91
		v118 = v86
		v119 = v88
		v120 = v90
		v121 = v3
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v97 = F_FreeSpaceMapPrepareTruncateRel(m, l0, l1)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v97
	if v97 == int32(-1) {
		v117 = v91
		v118 = v86
		v119 = v88
		v120 = v90
		v121 = v3
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v108 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+136)) = v108
	v112 = F_smgrnblocks(m, v42, v108)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+124)) = v112
	v117 = int32(2)
	v118 = v13 + int32(140)
	v119 = v13 + int32(128)
	v120 = v13 + int32(116)
	v121 = v108
	goto L19
L25:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[0]))
	if v170 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L26:
	;
	v148 = v122
	goto L28
L27:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v124
	v126 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v126
	v130 = F_smgropen(m, v13+int32(40), v123)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	v150 = F_smgrexists(m, v148, int32(2))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L34
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v130
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130)+72))
	if v134 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v148 = v146
	goto L28
L31:
	;
	v142 = v134
	goto L33
L32:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+76))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+4)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v138
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v130)+72))
	v142 = v140
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+72)) = v142 + int32(1)
	goto L30
L34:
	;
	if v150 == int32(0) {
		v168 = v117
		goto L25
	} else {
		goto L35
	}
L35:
	;
	v154 = F_visibilitymap_prepare_truncate(m, l0, l1)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v154
	if v154 == int32(-1) {
		v168 = v117
		goto L25
	} else {
		goto L37
	}
L37:
	;
	v159 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v159
	v162 = F_smgrnblocks(m, v42, v159)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v162
	v168 = v117 + int32(1)
	goto L25
L39:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[1]))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v210)+120)) = v211 | int32(3)
	v215 = int32(_a_F_RelationTruncate_0)
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[2])) = v217 + int32(1)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+118)))
	if v222 != int32(112) {
		goto L51
	} else {
		goto L52
	}
L40:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v173 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v199 = v173
	goto L43
L42:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v175
	v177 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v177
	v181 = F_smgropen(m, v13+int32(24), v174)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L44
	}
L43:
	;
	v200 = int32(0)
	v202 = F_hash_search(m, v170, v199, v200, v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L49
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v181
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181)+72))
	if v185 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v199 = v197
	goto L43
L46:
	;
	v193 = v185
	goto L48
L47:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181)+76))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v181)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+4)) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v181)+72))
	v193 = v191
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+72)) = v193 + int32(1)
	goto L45
L49:
	;
	if v202 == int32(0) {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+12)) = uint8(v206)
	goto L39
L51:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v251 != 0 {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[3]))
	if v226 <= int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v229 != 0 {
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = l1
	v232 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+92)) = v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = int32(7)
	F_XLogBeginInsert(m)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v230 != 0 {
		goto L51
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	F_XLogRegisterData(m, v13+int32(88), int32(20))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v247 = F_XLogInsert(m, int32(2), int32(33))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	F_XLogFlush(m, v247)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	goto L51
L62:
	;
	v277 = v251
	goto L64
L63:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v253
	v255 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v255
	v259 = F_smgropen(m, v13+int32(8), v252)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L65
	}
L64:
	;
	F_smgrtruncate(m, v277, v13+int32(132), v168, v13+int32(120), v13+int32(108))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L70
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v259
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v259)+72))
	if v263 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v277 = v275
	goto L64
L67:
	;
	v271 = v263
	goto L69
L68:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v259)+76))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v259)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v265
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v259)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v265))) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v259)+72))
	v271 = v269
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259)+72)) = v271 + int32(1)
	goto L66
L70:
	;
	v286 = int32(_a_F_RelationTruncate_0)
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[2])) = v288 - int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[1]))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v293)+120)) = v294 & int32(-4)
	if v121 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_FreeSpaceMapVacuumRange(m, l0, l1, int32(-1))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	m.G0 = v13 + int32(144)
	return
L74:
	;
	goto L73
}
func F_SetRelationHasSubclass(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v2 = l1
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v16 = F_SearchSysCacheCopy(m, int32(57), l0, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
				v20 = v18 + v19
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+126)))
				if v2 != v21 {
					*(*uint8)(unsafe.Add(mBase, uint32(v20)+126)) = uint8(v2)
					F_CatalogTupleUpdate(m, v12, v16+int32(4), v16)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_pfree(m, v16)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_relation_close(m, v12, int32(3))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				} else {
					F_CacheInvalidateRelcacheByTuple(m, v16)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_pfree(m, v16)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_relation_close(m, v12, int32(3))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_SetRelationHasSubclass_0), v8)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_SetRelationHasSubclass_1), int32(3664), int32(_a_F_SetRelationHasSubclass_2))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
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
func F_UnlockRelation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v9
	v15 = F_LockRelease(m, v5, int32(8), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_UnlockRelationForExtension(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(72339069014638592)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v10
	v15 = F_LockRelease(m, v6, l1, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func F_UnlockRelationIdForSession(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	Fn13861(m, l0, l1, int32(1))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_UnlockRelationOid(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v53 int32
	_ = v53
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v82 = v11
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v82 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v82 = int32(0)
				} else {
					v82 = v11
				}
			}
		} else {
			v23 = l0 - int32(2671)
			if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v23))|base.B2i32(int32(1)<<(uint(v23)%32)&int32(226492515) == int32(0)) != 0 {
				if base.B2i32(base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l0-int32(2846)) < base.Ui32(int32(2))) != 0 {
					v82 = v11
				} else {
					v82 = int32(0)
				}
			} else {
				v82 = v11
			}
		}
	} else {
		if l0 <= int32(_a_F_UnlockRelationOid_0) {
			v36 = l0 - int32(_a_F_UnlockRelationOid_1)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v36))|base.B2i32(int32(1)<<(uint(v36)%32)&int32(963) == int32(0)) != 0 {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v82 = v11
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v82 = int32(0)
					} else {
						v82 = v11
					}
				}
			} else {
				v82 = v11
			}
		} else {
			switch l0 - int32(_a_F_UnlockRelationOid_2) {
			case 0, 1, 2, 3, 4, 59, 60:
				v82 = v11
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v82 = int32(0)
			default:
				if base.Ui32(l0-int32(_a_F_UnlockRelationOid_3)) < base.Ui32(int32(3)) {
					v82 = v11
				} else {
					v53 = l0 - int32(_a_F_UnlockRelationOid_4)
					if base.Ui32(int32(15)) < base.Ui32(v53) {
						v82 = int32(0)
					} else {
						if int32(1)<<(uint(v53)%32)&int32(_a_F_UnlockRelationOid_5) != 0 {
							v82 = v11
						} else {
							v82 = int32(0)
						}
					}
				}
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockRelationOid[0]))
	if v82 != 0 {
		v89 = int32(0)
	} else {
		v89 = v88
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v89
	v92 = F_LockRelease(m, v7, l1, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_get_relation_statistics_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	v11 = F_SearchSysCache2(m, int32(62), l2, l3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
			v15 = v13 + v14
			v17 = F_statext_is_kind_built(m, v11, int32(100))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				if v17 != 0 {
					v20 = F_palloc0(m, int32(28))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(271)
						v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
						v26 = int32(100)
						*(*uint8)(unsafe.Add(mBase, uint32(v20)+16)) = uint8(v26)
						*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = l1
						*(*uint8)(unsafe.Add(mBase, uint32(v20)+8)) = uint8(v25)
						v30 = F_bms_copy(m, l4)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = l5
							*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v30
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v35 = F_lappend(m, v34, v20)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v35
								v41 = F_statext_is_kind_built(m, v11, int32(102))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									if v41 != 0 {
										v44 = F_palloc0(m, int32(28))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(271)
											v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
											v50 = int32(102)
											*(*uint8)(unsafe.Add(mBase, uint32(v44)+16)) = uint8(v50)
											*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = l1
											*(*uint8)(unsafe.Add(mBase, uint32(v44)+8)) = uint8(v49)
											v54 = F_bms_copy(m, l4)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = l5
												*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v54
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v59 = F_lappend(m, v58, v44)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0))) = v59
													v65 = F_statext_is_kind_built(m, v11, int32(109))
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return
													} else {
														if v65 != 0 {
															v68 = F_palloc0(m, int32(28))
															mBase = m.M
															v69 = m.ExcPending
															if v69 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(271)
																v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
																v74 = int32(109)
																*(*uint8)(unsafe.Add(mBase, uint32(v68)+16)) = uint8(v74)
																*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = l1
																*(*uint8)(unsafe.Add(mBase, uint32(v68)+8)) = uint8(v73)
																v78 = F_bms_copy(m, l4)
																mBase = m.M
																v79 = m.ExcPending
																if v79 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = l5
																	*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v78
																	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																	v83 = F_lappend(m, v82, v68)
																	mBase = m.M
																	v84 = m.ExcPending
																	if v84 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v83
																		v89 = F_statext_is_kind_built(m, v11, int32(101))
																		mBase = m.M
																		v90 = m.ExcPending
																		if v90 != 0 {
																			return
																		} else {
																			if v89 != 0 {
																				v92 = F_palloc0(m, int32(28))
																				mBase = m.M
																				v93 = m.ExcPending
																				if v93 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = l2
																					*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(271)
																					v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
																					v98 = int32(101)
																					*(*uint8)(unsafe.Add(mBase, uint32(v92)+16)) = uint8(v98)
																					*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = l1
																					*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)) = uint8(v97)
																					v102 = F_bms_copy(m, l4)
																					mBase = m.M
																					v103 = m.ExcPending
																					if v103 != 0 {
																						return
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = l5
																						*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v102
																						v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																						v107 = F_lappend(m, v106, v92)
																						mBase = m.M
																						v108 = m.ExcPending
																						if v108 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
																							F_ReleaseCatCache(m, v11)
																							mBase = m.M
																							v114 = m.ExcPending
																							if v114 != 0 {
																								return
																							} else {
																								return
																							}
																						}
																					}
																				}
																			} else {
																				F_ReleaseCatCache(m, v11)
																				mBase = m.M
																				v114 = m.ExcPending
																				if v114 != 0 {
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
															v89 = F_statext_is_kind_built(m, v11, int32(101))
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return
															} else {
																if v89 != 0 {
																	v92 = F_palloc0(m, int32(28))
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = l2
																		*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(271)
																		v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
																		v98 = int32(101)
																		*(*uint8)(unsafe.Add(mBase, uint32(v92)+16)) = uint8(v98)
																		*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = l1
																		*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)) = uint8(v97)
																		v102 = F_bms_copy(m, l4)
																		mBase = m.M
																		v103 = m.ExcPending
																		if v103 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v102
																			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																			v107 = F_lappend(m, v106, v92)
																			mBase = m.M
																			v108 = m.ExcPending
																			if v108 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
																				F_ReleaseCatCache(m, v11)
																				mBase = m.M
																				v114 = m.ExcPending
																				if v114 != 0 {
																					return
																				} else {
																					return
																				}
																			}
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v11)
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
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
									} else {
										v65 = F_statext_is_kind_built(m, v11, int32(109))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											if v65 != 0 {
												v68 = F_palloc0(m, int32(28))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(271)
													v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
													v74 = int32(109)
													*(*uint8)(unsafe.Add(mBase, uint32(v68)+16)) = uint8(v74)
													*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = l1
													*(*uint8)(unsafe.Add(mBase, uint32(v68)+8)) = uint8(v73)
													v78 = F_bms_copy(m, l4)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v78
														v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														v83 = F_lappend(m, v82, v68)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0))) = v83
															v89 = F_statext_is_kind_built(m, v11, int32(101))
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return
															} else {
																if v89 != 0 {
																	v92 = F_palloc0(m, int32(28))
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = l2
																		*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(271)
																		v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
																		v98 = int32(101)
																		*(*uint8)(unsafe.Add(mBase, uint32(v92)+16)) = uint8(v98)
																		*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = l1
																		*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)) = uint8(v97)
																		v102 = F_bms_copy(m, l4)
																		mBase = m.M
																		v103 = m.ExcPending
																		if v103 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v102
																			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																			v107 = F_lappend(m, v106, v92)
																			mBase = m.M
																			v108 = m.ExcPending
																			if v108 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
																				F_ReleaseCatCache(m, v11)
																				mBase = m.M
																				v114 = m.ExcPending
																				if v114 != 0 {
																					return
																				} else {
																					return
																				}
																			}
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v11)
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
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
												v89 = F_statext_is_kind_built(m, v11, int32(101))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return
												} else {
													if v89 != 0 {
														v92 = F_palloc0(m, int32(28))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(271)
															v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
															v98 = int32(101)
															*(*uint8)(unsafe.Add(mBase, uint32(v92)+16)) = uint8(v98)
															*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = l1
															*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)) = uint8(v97)
															v102 = F_bms_copy(m, l4)
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = l5
																*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v102
																v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																v107 = F_lappend(m, v106, v92)
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
																	F_ReleaseCatCache(m, v11)
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
																		return
																	} else {
																		return
																	}
																}
															}
														}
													} else {
														F_ReleaseCatCache(m, v11)
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
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
					}
				} else {
					v41 = F_statext_is_kind_built(m, v11, int32(102))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						if v41 != 0 {
							v44 = F_palloc0(m, int32(28))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(271)
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
								v50 = int32(102)
								*(*uint8)(unsafe.Add(mBase, uint32(v44)+16)) = uint8(v50)
								*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = l1
								*(*uint8)(unsafe.Add(mBase, uint32(v44)+8)) = uint8(v49)
								v54 = F_bms_copy(m, l4)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = l5
									*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v54
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v59 = F_lappend(m, v58, v44)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v59
										v65 = F_statext_is_kind_built(m, v11, int32(109))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											if v65 != 0 {
												v68 = F_palloc0(m, int32(28))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(271)
													v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
													v74 = int32(109)
													*(*uint8)(unsafe.Add(mBase, uint32(v68)+16)) = uint8(v74)
													*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = l1
													*(*uint8)(unsafe.Add(mBase, uint32(v68)+8)) = uint8(v73)
													v78 = F_bms_copy(m, l4)
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v78
														v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														v83 = F_lappend(m, v82, v68)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0))) = v83
															v89 = F_statext_is_kind_built(m, v11, int32(101))
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return
															} else {
																if v89 != 0 {
																	v92 = F_palloc0(m, int32(28))
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = l2
																		*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(271)
																		v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
																		v98 = int32(101)
																		*(*uint8)(unsafe.Add(mBase, uint32(v92)+16)) = uint8(v98)
																		*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = l1
																		*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)) = uint8(v97)
																		v102 = F_bms_copy(m, l4)
																		mBase = m.M
																		v103 = m.ExcPending
																		if v103 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v102
																			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																			v107 = F_lappend(m, v106, v92)
																			mBase = m.M
																			v108 = m.ExcPending
																			if v108 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
																				F_ReleaseCatCache(m, v11)
																				mBase = m.M
																				v114 = m.ExcPending
																				if v114 != 0 {
																					return
																				} else {
																					return
																				}
																			}
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v11)
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
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
												v89 = F_statext_is_kind_built(m, v11, int32(101))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return
												} else {
													if v89 != 0 {
														v92 = F_palloc0(m, int32(28))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(271)
															v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
															v98 = int32(101)
															*(*uint8)(unsafe.Add(mBase, uint32(v92)+16)) = uint8(v98)
															*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = l1
															*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)) = uint8(v97)
															v102 = F_bms_copy(m, l4)
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = l5
																*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v102
																v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																v107 = F_lappend(m, v106, v92)
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
																	F_ReleaseCatCache(m, v11)
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
																		return
																	} else {
																		return
																	}
																}
															}
														}
													} else {
														F_ReleaseCatCache(m, v11)
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
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
						} else {
							v65 = F_statext_is_kind_built(m, v11, int32(109))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								if v65 != 0 {
									v68 = F_palloc0(m, int32(28))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(271)
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
										v74 = int32(109)
										*(*uint8)(unsafe.Add(mBase, uint32(v68)+16)) = uint8(v74)
										*(*int32)(unsafe.Add(mBase, uint32(v68)+12)) = l1
										*(*uint8)(unsafe.Add(mBase, uint32(v68)+8)) = uint8(v73)
										v78 = F_bms_copy(m, l4)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = l5
											*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = v78
											v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v83 = F_lappend(m, v82, v68)
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0))) = v83
												v89 = F_statext_is_kind_built(m, v11, int32(101))
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return
												} else {
													if v89 != 0 {
														v92 = F_palloc0(m, int32(28))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(271)
															v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
															v98 = int32(101)
															*(*uint8)(unsafe.Add(mBase, uint32(v92)+16)) = uint8(v98)
															*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = l1
															*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)) = uint8(v97)
															v102 = F_bms_copy(m, l4)
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = l5
																*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v102
																v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																v107 = F_lappend(m, v106, v92)
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
																	F_ReleaseCatCache(m, v11)
																	mBase = m.M
																	v114 = m.ExcPending
																	if v114 != 0 {
																		return
																	} else {
																		return
																	}
																}
															}
														}
													} else {
														F_ReleaseCatCache(m, v11)
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
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
									v89 = F_statext_is_kind_built(m, v11, int32(101))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										if v89 != 0 {
											v92 = F_palloc0(m, int32(28))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v92))) = int32(271)
												v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
												v98 = int32(101)
												*(*uint8)(unsafe.Add(mBase, uint32(v92)+16)) = uint8(v98)
												*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = l1
												*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)) = uint8(v97)
												v102 = F_bms_copy(m, l4)
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v92)+24)) = l5
													*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v102
													v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v107 = F_lappend(m, v106, v92)
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0))) = v107
														F_ReleaseCatCache(m, v11)
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
															return
														} else {
															return
														}
													}
												}
											}
										} else {
											F_ReleaseCatCache(m, v11)
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
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
		} else {
			return
		}
	}
}
func F_relation_mark_replica_identity(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v2 = l1
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v18 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v23 = F_SearchSysCacheCopy(m, int32(57), v21, int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L42
	}
L4:
	;
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
	v27 = v25 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+130)))
	if v28 != v2&int32(255) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L39
	}
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+130)) = uint8(v2)
	F_CatalogTupleUpdate(m, v18, v23+int32(4), v23)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_relation_close(m, v18, int32(3))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	F_pfree(m, v23)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v44 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v46 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	F_relation_close(m, v44, int32(3))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L38
	}
L16:
	;
	if v46 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v50 <= int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v55 = int32(0)
	goto L19
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v55<<(uint(int32(2))%32))))
	v72 = F_SearchSysCacheCopy(m, int32(34), v70, int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L15
L21:
	;
	if v72 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+22)))
	v78 = v76 + v77
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+22)))
	if l2 == v70 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	F_pfree(m, v72)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L36
	}
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v78)+22)) = uint8(v91)
	F_CatalogTupleUpdate(m, v44, v72+int32(4), v72)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L30
	}
L25:
	;
	v81 = int32(1)
	if v79&v81 == int32(0) {
		v91 = v81
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v86 = int32(0)
	if v79&int32(1) == v86 {
		goto L23
	} else {
		goto L29
	}
L28:
	;
	goto L23
L29:
	;
	v91 = v86
	goto L24
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_relation_mark_replica_identity[0]))
	if v98 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v100 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2610), v70, v100, v100, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_CacheInvalidateRelcache(m, l0)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	goto L23
L36:
	;
	v111 = v55 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v111 < v112 {
		v55 = v111
		goto L19
	} else {
		goto L37
	}
L37:
	;
	goto L20
L38:
	;
	m.G0 = v14 + int32(32)
	return
L39:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v135 + int32(4)
	F_errmsg_internal(m, int32(_a_F_relation_mark_replica_identity_0), v14)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_relation_mark_replica_identity_1), int32(_a_F_relation_mark_replica_identity_2), int32(_a_F_relation_mark_replica_identity_3))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v70
	F_errmsg_internal(m, int32(_a_F_relation_mark_replica_identity_4), v14+int32(16))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_relation_mark_replica_identity_1), int32(_a_F_relation_mark_replica_identity_5), int32(_a_F_relation_mark_replica_identity_3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_swap_relation_files(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 float32
	_ = v363
	var v364 float32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	v20 = m.G0
	v22 = v20 - int32(224)
	m.G0 = v22
	v26 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v30 = F_SearchSysCacheCopy(m, int32(57), l0, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L17
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L243
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L240
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L237
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L234
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L228
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L222
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L219
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L216
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L213
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L210
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L207
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L204
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L201
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L198
	}
L17:
	;
	if v30 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
	v36 = F_SearchSysCacheCopy(m, int32(57), l1, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L195
	}
L21:
	;
	if v36 == int32(0) {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
	v42 = v40 + v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+84))
	v44 = v32 + v33
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+84))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+88))
	v47 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+88))
	if base.B2i32(v46 == v47)|base.B2i32(v49 == v47) == v47 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v312 = F_relation_open(m, l0, int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L117
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+88)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v42)+88)) = v46
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v44)+92))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v42)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+92)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v42)+92)) = v57
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v44)+84))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v42)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+84)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v42)+84)) = v61
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+118)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+118)))
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+118)) = uint8(v66)
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+118)) = uint8(v65)
	if l3 != 0 {
		v308 = l7
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v46|v49 != 0 {
		goto L15
	} else {
		goto L28
	}
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v44)+112))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v42)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+112)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v42)+112)) = v69
	v308 = l7
	goto L23
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v44)+92))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)+92))
	if v74 != v75 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+118)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+118)))
	if v77 != v78 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	if v43 != v45 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	if l3 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v44)+112))
	if v83 != 0 {
		goto L11
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+117)))
	v86 = int32(0)
	if v85 == v86 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v42)+112))
	if v84 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	if v188 == int32(0) {
		goto L10
	} else {
		goto L75
	}
L38:
	;
	goto L37
L39:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v188 = v183
	goto L38
L40:
	;
	v91 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[0]))
	if v91 < v93 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v134 = int32(0)
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[1]))
	if v134 < v136 {
		goto L61
	} else {
		goto L62
	}
L43:
	;
	v97 = v91
	goto L46
L44:
	;
	goto L45
L45:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[2]))
	if v116 <= int32(0) {
		v188 = v86
		goto L38
	} else {
		goto L52
	}
L46:
	;
	v102 = v97 << (uint(int32(3)) % 32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_swap_relation_files[3])))
	if v103 == l0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L45
L48:
	;
	v182 = v102 + int32(_a_F_swap_relation_files_0)
	goto L39
L49:
	;
	goto L50
L50:
	;
	v108 = v97 + int32(1)
	if v108 != v93 {
		v97 = v108
		goto L46
	} else {
		goto L51
	}
L51:
	;
	goto L47
L52:
	;
	v121 = int32(0)
	goto L53
L53:
	;
	v126 = v121 << (uint(int32(3)) % 32)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+uint32(_c_F_swap_relation_files[4])))
	if v127 != l0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v182 = v126 + int32(_a_F_swap_relation_files_1)
	goto L39
L55:
	;
	v130 = v121 + int32(1)
	if v116 != v130 {
		v121 = v130
		goto L53
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	goto L54
L58:
	;
	v188 = v86
	goto L38
L59:
	;
	v164 = int32(0)
	goto L69
L60:
	;
	v182 = v145 + int32(_a_F_swap_relation_files_2)
	goto L39
L61:
	;
	v140 = v134
	goto L64
L62:
	;
	goto L63
L63:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[5]))
	if v157 <= int32(0) {
		v188 = v86
		goto L38
	} else {
		goto L68
	}
L64:
	;
	v145 = v140 << (uint(int32(3)) % 32)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+uint32(_c_F_swap_relation_files[6])))
	if l0 == v146 {
		goto L60
	} else {
		goto L66
	}
L65:
	;
	goto L63
L66:
	;
	v149 = v140 + int32(1)
	if v149 != v136 {
		v140 = v149
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	goto L59
L69:
	;
	v169 = v164 << (uint(int32(3)) % 32)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+uint32(_c_F_swap_relation_files[7])))
	if v170 != l0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v182 = v169 + int32(_a_F_swap_relation_files_3)
	goto L39
L71:
	;
	v173 = v164 + int32(1)
	if v157 != v173 {
		v164 = v173
		goto L69
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	goto L70
L74:
	;
	v188 = v86
	goto L38
L75:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+117)))
	v192 = int32(0)
	if v191 == v192 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	if v294 == int32(0) {
		goto L9
	} else {
		goto L114
	}
L77:
	;
	goto L76
L78:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	v294 = v289
	goto L77
L79:
	;
	v197 = int32(0)
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[0]))
	if v197 < v199 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	v240 = int32(0)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[1]))
	if v240 < v242 {
		goto L100
	} else {
		goto L101
	}
L82:
	;
	v203 = v197
	goto L85
L83:
	;
	goto L84
L84:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[2]))
	if v222 <= int32(0) {
		v294 = v192
		goto L77
	} else {
		goto L91
	}
L85:
	;
	v208 = v203 << (uint(int32(3)) % 32)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+uint32(_c_F_swap_relation_files[3])))
	if v209 == l1 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L84
L87:
	;
	v288 = v208 + int32(_a_F_swap_relation_files_0)
	goto L78
L88:
	;
	goto L89
L89:
	;
	v214 = v203 + int32(1)
	if v214 != v199 {
		v203 = v214
		goto L85
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	v227 = int32(0)
	goto L92
L92:
	;
	v232 = v227 << (uint(int32(3)) % 32)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+uint32(_c_F_swap_relation_files[4])))
	if v233 != l1 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v288 = v232 + int32(_a_F_swap_relation_files_1)
	goto L78
L94:
	;
	v236 = v227 + int32(1)
	if v222 != v236 {
		v227 = v236
		goto L92
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	goto L93
L97:
	;
	v294 = v192
	goto L77
L98:
	;
	v270 = int32(0)
	goto L108
L99:
	;
	v288 = v251 + int32(_a_F_swap_relation_files_2)
	goto L78
L100:
	;
	v246 = v240
	goto L103
L101:
	;
	goto L102
L102:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[5]))
	if v263 <= int32(0) {
		v294 = v192
		goto L77
	} else {
		goto L107
	}
L103:
	;
	v251 = v246 << (uint(int32(3)) % 32)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+uint32(_c_F_swap_relation_files[6])))
	if l1 == v252 {
		goto L99
	} else {
		goto L105
	}
L104:
	;
	goto L102
L105:
	;
	v255 = v246 + int32(1)
	if v255 != v242 {
		v246 = v255
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	goto L98
L108:
	;
	v275 = v270 << (uint(int32(3)) % 32)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_swap_relation_files[7])))
	if v276 != l1 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v288 = v275 + int32(_a_F_swap_relation_files_3)
	goto L78
L110:
	;
	v279 = v270 + int32(1)
	if v263 != v279 {
		v270 = v279
		goto L108
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	goto L109
L113:
	;
	v294 = v192
	goto L77
L114:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+117)))
	F_RelationMapUpdateMap(m, l0, v294, v297, int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+117)))
	F_RelationMapUpdateMap(m, l1, v188, v301, int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = l1
	v308 = l7 + int32(4)
	goto L23
L117:
	;
	v315 = F_relation_open(m, l1, int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v312)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+32)) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v312)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+36)) = v319
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v312)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+40)) = v321
	v324 = F_GetCurrentSubTransactionId(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v312)+36)) = v324
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v312)+40))
	if v326 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	F_relation_close(m, v312, int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L126
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+40)) = v324
	goto L122
L121:
	;
	goto L122
L122:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[8]))
	if v331 <= int32(31) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v312)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[8])) = v331 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v331<<(uint(int32(2))%32))+uint32(_c_F_swap_relation_files[9]))) = v334
	goto L119
L124:
	;
	goto L125
L125:
	;
	v345 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_swap_relation_files[10])) = uint8(v345)
	goto L119
L126:
	;
	F_relation_close(m, v315, int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+119)))
	if v354 != int32(105) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+140)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v44)+136)) = l5
	goto L130
L129:
	;
	goto L130
L130:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v44)+96))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+96)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v42)+96)) = v359
	v363 = *(*float32)(unsafe.Add(mBase, uint32(v44)+100))
	v364 = *(*float32)(unsafe.Add(mBase, uint32(v42)+100))
	*(*float32)(unsafe.Add(mBase, uint32(v44)+100)) = v364
	*(*float32)(unsafe.Add(mBase, uint32(v42)+100)) = v363
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v44)+104))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+104)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v42)+104)) = v367
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v44)+108))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v42)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+108)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v42)+108)) = v371
	if l2 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v43 != v45 {
		goto L141
	} else {
		goto L142
	}
L132:
	;
	v379 = F_CatalogOpenIndexes(m, v26)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	F_CacheInvalidateRelcacheByTuple(m, v30)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L139
	}
L135:
	;
	F_CatalogTupleUpdateWithInfo(m, v26, v30+int32(4), v30, v379)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_CatalogTupleUpdateWithInfo(m, v26, v36+int32(4), v36, v379)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_CatalogCloseIndexes(m, v379)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	goto L131
L139:
	;
	F_CacheInvalidateRelcacheByTuple(m, v36)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	goto L131
L141:
	;
	v397 = F_changeDependencyFor(m, int32(1259), l0, int32(2601), v45, v43)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[11]))
	if v408 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L144:
	;
	if v397 != int32(1) {
		goto L8
	} else {
		goto L145
	}
L145:
	;
	v403 = F_changeDependencyFor(m, int32(1259), l1, int32(2601), v43, v45)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	if v403 != int32(1) {
		goto L7
	} else {
		goto L147
	}
L147:
	;
	goto L143
L148:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v44)+112))
	if v426 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L149:
	;
	v412 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l0, v412, v412, l4)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[11]))
	if v417 == int32(0) {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	v421 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l1, v421, v421, int32(1))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	goto L148
L153:
	;
	F_pfree(m, v30)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L192
	}
L154:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+119)))
	if v501 != int32(116) {
		goto L153
	} else {
		goto L187
	}
L155:
	;
	if l3 == int32(0) {
		goto L153
	} else {
		goto L186
	}
L156:
	;
	v443 = int32(1)
	if base.Ui32(l0) < base.Ui32(int32(_a_F_swap_relation_files_4)) {
		v451 = v443
		goto L166
	} else {
		goto L167
	}
L157:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v42)+112))
	if v429 == int32(0) {
		goto L155
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	if l3 == int32(0) {
		goto L156
	} else {
		goto L162
	}
L160:
	;
	if l3 == int32(0) {
		goto L156
	} else {
		goto L161
	}
L161:
	;
	goto L3
L162:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v42)+112))
	if v436 == int32(0) {
		goto L3
	} else {
		goto L163
	}
L163:
	;
	F_swap_relation_files(m, v426, v436, l2, int32(1), l4, l5, l6, v308)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	goto L154
L165:
	;
	if v451 != 0 {
		goto L6
	} else {
		goto L169
	}
L166:
	;
	goto L165
L167:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v44)+68))
	if v446 == int32(99) {
		v451 = v443
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v449 = F_isTempToastNamespace(m, v446)
	mBase = m.M
	v451 = v449
	goto L166
L169:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v44)+112))
	if v452 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v455 = F_deleteDependencyRecordsFor(m, int32(1259), v452, int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v42)+112))
	if v460 != 0 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	if v455 != int32(1) {
		goto L5
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v463 = F_deleteDependencyRecordsFor(m, int32(1259), v460, int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v468 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+220)) = v468
	v470 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v470
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v44)+112))
	if v476 != 0 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	if v463 != int32(1) {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+204)) = v476
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = l0
	F_recordDependencyOn(m, v22+int32(200), v22+int32(212), int32(105))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v42)+112))
	if v486 == int32(0) {
		goto L153
	} else {
		goto L184
	}
L183:
	;
	goto L182
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+204)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = l1
	F_recordDependencyOn(m, v22+int32(200), v22+int32(212), int32(105))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	goto L153
L186:
	;
	goto L154
L187:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
	if v504 != int32(116) {
		goto L153
	} else {
		goto L188
	}
L188:
	;
	v508 = F_toast_get_valid_index(m, l0, int32(8))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v511 = F_toast_get_valid_index(m, l1, int32(8))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v514 = int32(0)
	F_swap_relation_files(m, v508, v511, l2, int32(1), l4, v514, v514, v308)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	goto L153
L192:
	;
	F_pfree(m, v36)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_relation_close(m, v26, int32(3))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	m.G0 = v22 + int32(224)
	return
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = l0
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_5), v22)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1087), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_5), v22+int32(16))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1092), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v44 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_8), v22+int32(96))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1142), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v44 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_9), v22+int32(192))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1153), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v44 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_10), v22+int32(176))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1156), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v44 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_11), v22+int32(160))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1159), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v44 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_12), v22+int32(144))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1163), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v44 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_13), v22+int32(112))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1171), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+132)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v42 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_13), v22+int32(128))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1175), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	v684 = F_get_rel_namespace(m, l0)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	v686 = F_get_namespace_name(m, v684)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v688 = F_get_rel_name(m, l0)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v686
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_14), v22+int32(80))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1289), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	v706 = F_get_rel_namespace(m, l1)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	v708 = F_get_namespace_name(m, v706)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	v710 = F_get_rel_name(m, l1)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v708
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_14), v22-int32(-64))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1297), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_15), int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1359), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v455
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_16), v22+int32(48))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1369), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v463
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_16), v22+int32(32))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1378), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_17), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1332), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
