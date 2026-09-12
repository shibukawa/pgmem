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
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[882])))
	v14 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[883]))
	F_read_relmap_file(m, int32(4433700), v17, int32(0), int32(22))
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
	v22 = int32(4443856)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
	v29 = *(*int32)(unsafe.Add(mBase, _consts[205]))
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
	v72 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	if v72 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	F_formrdesc(m, int32(121901), int32(83), int32(0), int32(34), int32(1704192))
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
	F_formrdesc(m, int32(333199), int32(75), int32(0), int32(25), int32(1707600))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_formrdesc(m, int32(467765), int32(81), int32(0), int32(30), int32(1710112))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_formrdesc(m, int32(350286), int32(71), int32(0), int32(32), int32(1713120))
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
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L108
	}
L16:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L104
	}
L17:
	;
	m.G0 = v9 + int32(48)
	return
L18:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _consts[881])))
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
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _consts[882])))
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
	*(*uint8)(unsafe.Add(mBase, _consts[881])) = uint8(v108)
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
	v144 = *(*int32)(unsafe.Add(mBase, _consts[884]))
	F_hash_seq_init(m, v9+int32(28), v144)
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
	*(*uint8)(unsafe.Add(mBase, _consts[882])) = uint8(v139)
	goto L31
L38:
	;
	v149 = F_hash_seq_search(m, v9+int32(28))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v149 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v151 = v149
	goto L43
L41:
	;
	goto L42
L42:
	;
	if v68&int32(1) == int32(0) {
		goto L17
	} else {
		goto L97
	}
L43:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	v159 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	F_ResourceOwnerEnlarge(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+16)) = v162 + int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	if v167 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	F_ResourceOwnerRemember(m, v169, v157, int32(1698240))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v157)+48))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+80))
	if v174 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L48
L50:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v157)+56))
	v179 = F_SearchSysCache1(m, int32(57), v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	v201 = v173
	goto L52
L52:
	;
	v204 = base.B2i32(v174 == int32(0))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+124)))
	if v205 != int32(1) {
		v217 = v201
		v218 = v204
		goto L66
	} else {
		goto L67
	}
L53:
	;
	if v179 == int32(0) {
		goto L16
	} else {
		goto L54
	}
L54:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v157)+48))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+22)))
	goto L56
L55:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v157)+180))
	if v190 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v188 = F__emscripten_memcpy_bulkmem(m, v183, v184+v185, int32(144))
	mBase = m.M
	goto L58
L58:
	;
	goto L55
L59:
	;
	F_pfree(m, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	F_RelationParseRelOptions(m, v157, v179)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	F_ReleaseCatCache(m, v179)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v157)+48))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+80))
	if v198 == int32(0) {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	v201 = v197
	goto L52
L66:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+125)))
	if v219 != int32(1) {
		v231 = v217
		v232 = v218
		goto L71
	} else {
		goto L72
	}
L67:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v157)+68))
	if v208 != 0 {
		v217 = v201
		v218 = v204
		goto L66
	} else {
		goto L68
	}
L68:
	;
	F_RelationBuildRuleLock(m, v157)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v157)+48))
	v212 = int32(1)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v157)+68))
	if v213 != 0 {
		v217 = v211
		v218 = v212
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v214 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+124)) = uint8(v214)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v157)+48))
	v217 = v216
	v218 = v212
	goto L66
L71:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+127)))
	if v233 != int32(1) {
		v240 = v232
		goto L76
	} else {
		goto L77
	}
L72:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v157)+76))
	if v222 != 0 {
		v231 = v217
		v232 = v218
		goto L71
	} else {
		goto L73
	}
L73:
	;
	F_RelationBuildTriggers(m, v157)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v157)+48))
	v226 = int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v157)+76))
	if v227 != 0 {
		v231 = v225
		v232 = v226
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+125)) = uint8(v228)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v157)+48))
	v231 = v230
	v232 = v226
	goto L71
L76:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v157)+188))
	if v241 != 0 {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v157)+80))
	if v236 != 0 {
		v240 = v232
		goto L76
	} else {
		goto L78
	}
L78:
	;
	F_RelationBuildRowSecurity(m, v157)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v240 = int32(1)
	goto L76
L80:
	;
	v286 = F_hash_seq_search(m, v9+int32(28))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L95
	}
L81:
	;
	F_hash_seq_term(m, v9+int32(28))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L93
	}
L82:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+16)) = v261 - int32(1)
	v266 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	if v266 != 0 {
		goto L88
	} else {
		goto L89
	}
L83:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v157)+48))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+119)))
	switch v243 - int32(83) {
	case 0, 26, 31, 33:
		goto L84
	default:
		goto L82
	}
L84:
	;
	F_RelationInitTableAccessMethod(m, v157)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+16)) = v248 - int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	if v253 == int32(0) {
		goto L81
	} else {
		goto L86
	}
L86:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	F_ResourceOwnerForget(m, v257, v157, int32(1698240))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	goto L81
L88:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	F_ResourceOwnerForget(m, v268, v157, int32(1698240))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	if v240 == int32(0) {
		goto L80
	} else {
		goto L92
	}
L91:
	;
	goto L90
L92:
	;
	goto L81
L93:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[884]))
	F_hash_seq_init(m, v9+int32(28), v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	goto L80
L95:
	;
	if v286 != 0 {
		v151 = v286
		goto L43
	} else {
		goto L96
	}
L96:
	;
	goto L44
L97:
	;
	v299 = int32(0)
	goto L98
L98:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v299<<(uint(int32(2))%32))+uint32(_consts[885])))
	F_InitCatCachePhase2(m, v309, int32(1))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L100
	}
L99:
	;
	F_write_relcache_init_file(m, int32(1))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L102
	}
L100:
	;
	v314 = v299 + int32(1)
	if v314 != int32(85) {
		v299 = v314
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	F_write_relcache_init_file(m, int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	goto L17
L104:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v157)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v339
	F_errmsg_internal(m, int32(43981), v9)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(477505), int32(4273), int32(529412))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v157)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v353 + int32(4)
	F_errmsg_internal(m, int32(660159), v9+int32(16))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(477505), int32(4301), int32(529412))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
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
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v5 < v15 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = base.B2i32(l2 == int32(3))&base.B2i32(l3 == int32(117)) | base.B2i32(l3 == int32(112))
	goto L3
L2:
	;
	v27 = int32(0)
	goto L3
L3:
	;
	v29 = F_palloc(m, int32(424))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)) = uint8(v27)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = l1
	v36 = F_smgrnblocks(m, l1, l2)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+400)) = v36
	v39 = F_GetRedoRecPtr(m)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+408)) = v39
	v43 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+416)) = v43
	v45 = F_smgrnblocks(m, l0, l2)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v12+int32(20), v129, v130, v131, v132, l2)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L34
	}
L9:
	;
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v52 = v5
	goto L13
L11:
	;
	goto L12
L12:
	;
	F_smgr_bulk_finish(m, v29)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L33
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v57 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v60 = F_smgr_bulk_get_buf(m, v29)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v60
	v63 = int32(4438508)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[83])) = v65 + v66
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v72*int32(80))+uint32(_consts[231])))
	m.T0[v77].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, l2, v52, v12+int32(20), v66)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v80 = int32(4438508)
	v82 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	v83 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[83])) = v82 - v83
	v89 = int32(*(*uint8)(unsafe.Add(mBase, _consts[232])))
	if v89 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v90 = int32(5)
	goto L23
L22:
	;
	v90 = v83
	goto L23
L23:
	;
	v93 = F_PageIsVerified(m, v60, v52, v90, v12+int32(95))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+95)))
	if v95 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_pgstat_prepare_report_checksum_failure(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v93 == int32(0) {
		goto L8
	} else {
		goto L30
	}
L28:
	;
	F_pgstat_report_checksum_failures_in_db(m, v98, int32(1))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	F_smgr_bulk_write(m, v29, v52, v60, int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v111 = v52 + int32(1)
	if v111 != v45 {
		v52 = v111
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L14
L33:
	;
	m.G0 = v12 + int32(96)
	return
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v12 + int32(20)
	F_errmsg(m, int32(667487), v12)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(477574), int32(550), int32(388924))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	v5 = int32(0)
	v17 = m.G0
	v21 = (v17 - int32(12288)) & int32(-4096)
	m.G0 = v21
	v28 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v5 < v28 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v31 = l3 | base.B2i32(l2 == int32(3))
	goto L3
L2:
	;
	v31 = v5
	goto L3
L3:
	;
	v35 = l0 + int32(8)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(4080)))) = v36
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+4072)) = v38
	v43 = F_smgropen(m, v21+int32(4072), int32(-1))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v45 = F_smgrnblocks(m, v43, l2)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v45 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v52 = F__emscripten_memset_bulkmem(m, v21+int32(4096), base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L10
L8:
	;
	goto L9
L9:
	;
	m.G0 = v17
	return
L10:
	;
	v56 = l1 + int32(8)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(4064)))) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+4056)) = v59
	v64 = F_smgropen(m, v21+int32(4056), int32(-1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v66 = int32(1)
	F_smgrextend(m, v64, l2, v45-v66, v21+int32(4096), v66)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v74 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v77 = F_GetAccessStrategy(m, int32(2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4092)) = v45
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4088)) = v80
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(4048)))) = v84
	v86 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+4040)) = v86
	v93 = F_smgropen(m, v21+int32(4040), int32(-1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if l3 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v97 = int32(112)
	goto L18
L17:
	;
	v97 = int32(117)
	goto L18
L18:
	;
	v102 = F_read_stream_begin_impl(m, int32(12), v74, v80, v93, v97, l2, int32(120), v21+int32(4088), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v119 = v5
	goto L20
L20:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v121 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_read_stream_end(m, v102)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L57
	}
L22:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v125 = F_read_stream_next_buffer(m, v102, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)+16))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(4032)))) = v170
	v172 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+4024)) = v172
	v177 = F_ReadBufferWithoutRelcache(m, v21+int32(4024), l2, v167, int32(1), v77, l3)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L33
	}
L27:
	;
	if v125 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v132 = v125 ^ int32(-1)
	v137 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137+v132<<(uint(int32(2))%32))))
	v164 = v130 + v132<<(uint(int32(6))%32)
	v166 = v141
	goto L26
L29:
	;
	goto L30
L30:
	;
	v143 = v125 << (uint(int32(6)) % 32)
	v145 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v150 = F_LWLockAcquire(m, v143+v145-int32(16), int32(1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v158 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v164 = v153 + v143 + int32(-64)
	v166 = v158 + v125<<(uint(int32(13))%32) + int32(-8192)
	goto L26
L32:
	;
	v197 = int32(4438516)
	v199 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v199 + int32(1)
	goto L38
L33:
	;
	if v177 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v182+(v177^int32(-1))<<(uint(int32(2))%32))))
	v196 = v188
	goto L32
L35:
	;
	goto L36
L36:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v196 = v190 + v177<<(uint(int32(13))%32) + int32(-8192)
	goto L32
L37:
	;
	F_MarkBufferDirty(m, v177)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L41
	}
L38:
	;
	v204 = F__emscripten_memcpy_bulkmem(m, v196, v166, int32(8192))
	mBase = m.M
	goto L40
L40:
	;
	goto L37
L41:
	;
	if v31 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_log_newpage_buffer(m, v177, int32(1))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v211 = int32(4438516)
	v213 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v213 - int32(1)
	if int32(0) <= v177 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	F_LWLockRelease(m, v220+v177<<(uint(int32(6))%32)-int32(16))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_ReleaseBuffer(m, v177)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	if int32(0) <= v125 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	F_LWLockRelease(m, v233+v125<<(uint(int32(6))%32)-int32(16))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_ReleaseBuffer(m, v125)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	v244 = v119 + int32(1)
	if v244 != v45 {
		v119 = v244
		goto L20
	} else {
		goto L56
	}
L56:
	;
	goto L21
L57:
	;
	F_bms_free(m, v74)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_bms_free(m, v77)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[87]))
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
		v18 = *(*int32)(unsafe.Add(mBase, _consts[37]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v19
		v21 = int32(4340688)
		v22 = *(*int32)(unsafe.Add(mBase, _consts[228]))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v22
		*(*int32)(unsafe.Add(mBase, _consts[228])) = v7
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v26 != 0 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
			v31 = v29 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v31
			if v31 == int32(0) {
				v36 = v26 + int32(76)
				v38 = *(*int32)(unsafe.Add(mBase, _consts[229]))
				if v38 != 0 {
					v40 = *(*int32)(unsafe.Add(mBase, _consts[230]))
					v45 = v40
				} else {
					v42 = int32(4367336)
					*(*int32)(unsafe.Add(mBase, _consts[229])) = v42
					v45 = v42
				}
				*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v45
				v47 = int32(4367336)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v47
				*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v36
				*(*int32)(unsafe.Add(mBase, _consts[230])) = v36
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
	var v135 int32
	_ = v135
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
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
		goto L6
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v17 + int32(112)
	return
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	F_ScanKeyInit(m, v17-int32(-64), int32(9), int32(3), int32(184), v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L21
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	if v22 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v36 = F__emscripten_memcpy_bulkmem(m, v23, v35, v22)
	mBase = m.M
	goto L12
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v30 != 0 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v39 = F__emscripten_memcpy_bulkmem(m, v26, v38, v22)
	mBase = m.M
	goto L16
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	goto L5
L18:
	;
	v42 = F__emscripten_memcpy_bulkmem(m, v31, v41, v30)
	mBase = m.M
	goto L20
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v55 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L27
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L95
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L92
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L89
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L86
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L83
	}
L27:
	;
	v58 = int32(1)
	v63 = F_systable_beginscan(m, v55, int32(2665), v58, int32(0), v58, v17-int32(-64))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v65 = F_systable_getnext(m, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v65 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v68 = v65
	v78 = int32(0)
	goto L33
L31:
	;
	goto L32
L32:
	;
	F_systable_endscan(m, v63)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L81
	}
L33:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+22)))
	v83 = v81 + v82
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+72)))
	if v84 == int32(120) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	F_systable_endscan(m, v63)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L54
	}
L35:
	;
	v122 = F_systable_getnext(m, v63)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L52
	}
L36:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v83)+88))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v92 != v93 {
		v121 = v78
		goto L35
	} else {
		goto L39
	}
L37:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+107)))
	if v87 != int32(1) {
		v121 = v78
		goto L35
	} else {
		goto L38
	}
L38:
	;
	switch v84 - int32(112) {
	case 0, 5:
		goto L36
	default:
		v121 = v78
		goto L35
	}
L39:
	;
	if v78 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v55)+52))
	v99 = F_fastgetattr_3(m, v68, int32(27), v96, v17+int32(63))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+63)))
	if v101 == int32(1) {
		goto L24
	} else {
		goto L42
	}
L42:
	;
	v104 = F_pg_detoast_datum(m, v99)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v106 != int32(1) {
		goto L23
	} else {
		goto L44
	}
L44:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	if v109 != v20 {
		goto L23
	} else {
		goto L45
	}
L45:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	if v111 != 0 {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	if v112 != int32(26) {
		goto L23
	} else {
		goto L47
	}
L47:
	;
	if v22 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v121 = int32(1)
	goto L35
L49:
	;
	v117 = F__emscripten_memcpy_bulkmem(m, v23, v104+int32(24), v22)
	mBase = m.M
	goto L51
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	if v122 != 0 {
		v68 = v122
		v78 = v121
		goto L33
	} else {
		goto L53
	}
L53:
	;
	goto L34
L54:
	;
	F_sequence_close(m, v55, int32(1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	if v121 == int32(0) {
		goto L26
	} else {
		goto L56
	}
L56:
	;
	if int32(0) < v20 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v135 = int32(0)
	goto L60
L58:
	;
	goto L59
L59:
	;
	v187 = int32(4443856)
	v188 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v190
	v192 = F_palloc(m, v22)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L66
	}
L60:
	;
	v149 = v135 << (uint(int32(2)) % 32)
	v151 = v149 + v23
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v153 = F_get_opcode(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	goto L59
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+v149))) = v153
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v160+v149)))
	v163 = F_get_op_opfamily_strategy(m, v159, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v31+v135<<(uint(int32(1))%32)))) = uint16(v163)
	if v163&int32(65535) == int32(0) {
		goto L22
	} else {
		goto L64
	}
L64:
	;
	v171 = v135 + int32(1)
	if v171 != v20 {
		v135 = v171
		goto L60
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+236)) = v192
	v195 = F_palloc(m, v22)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v195
	v198 = F_palloc(m, v30)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+244)) = v198
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v22 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	if v22 != 0 {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v202 = F__emscripten_memcpy_bulkmem(m, v201, v23, v22)
	mBase = m.M
	goto L72
L71:
	;
	goto L72
L72:
	;
	goto L69
L73:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v30 != 0 {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v205 = F__emscripten_memcpy_bulkmem(m, v204, v26, v22)
	mBase = m.M
	goto L76
L75:
	;
	goto L76
L76:
	;
	goto L73
L77:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v188
	goto L5
L78:
	;
	v208 = F__emscripten_memcpy_bulkmem(m, v207, v31, v30)
	mBase = m.M
	goto L80
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	F_sequence_close(m, v55, int32(1))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L26
L83:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v235 + int32(4)
	F_errmsg_internal(m, int32(176322), v17)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(477505), int32(5750), int32(230998))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v251 + int32(4)
	F_errmsg_internal(m, int32(176369), v17+int32(16))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(477505), int32(5723), int32(230998))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v269 + int32(4)
	F_errmsg_internal(m, int32(176273), v17+int32(32))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(477505), int32(5732), int32(230998))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errmsg_internal(m, int32(24282), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(477505), int32(5740), int32(230998))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v301+v135<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v300
	F_errmsg_internal(m, int32(37624), v17+int32(48))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(477505), int32(5761), int32(230998))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
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
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int64
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
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
	v235 = m.ExcPending
	if v235 != 0 {
		goto L29
	} else {
		goto L71
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L29
	} else {
		goto L68
	}
L3:
	;
	m.G0 = v8 + int32(48)
	return
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v17 = *(*int32)(unsafe.Add(mBase, _consts[271]))
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
	v21 = *(*int32)(unsafe.Add(mBase, _consts[226]))
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
	v197 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v197 < int32(0) {
		goto L3
	} else {
		goto L65
	}
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[529]))
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
		goto L37
	} else {
		goto L38
	}
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v87
	v194 = v87
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
	v35 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	if base.B2i32(base.Ui32(v42) < base.Ui32(int32(12000))) == int32(0) {
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
	v59 = *(*int32)(unsafe.Add(mBase, _consts[37]))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v190
	if v190 == int32(0) {
		goto L1
	} else {
		goto L64
	}
L34:
	;
	goto L33
L35:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v190 = v185
	goto L34
L36:
	;
	v168 = int32(0)
	goto L60
L37:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if int32(0) < v97 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	if int32(0) < v125 {
		goto L48
	} else {
		goto L49
	}
L40:
	;
	v102 = v91
	goto L43
L41:
	;
	goto L42
L42:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[256]))
	if v120 <= int32(0) {
		v190 = v91
		goto L34
	} else {
		goto L47
	}
L43:
	;
	v106 = v102 << (uint(int32(3)) % 32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+uint32(_consts[257])))
	if v89 == v109 {
		v181 = v106 + int32(4433184)
		goto L35
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	v112 = v102 + int32(1)
	if v112 != v97 {
		v102 = v112
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L36
L48:
	;
	v130 = v91
	goto L51
L49:
	;
	goto L50
L50:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	if v148 <= int32(0) {
		v190 = v91
		goto L34
	} else {
		goto L55
	}
L51:
	;
	v134 = v130 << (uint(int32(3)) % 32)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v134)+uint32(_consts[259])))
	if v89 == v137 {
		v181 = v134 + int32(4432136)
		goto L35
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	v140 = v130 + int32(1)
	if v140 != v125 {
		v130 = v140
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v154 = int32(0)
	goto L56
L56:
	;
	v158 = v154 << (uint(int32(3)) % 32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)+uint32(_consts[260])))
	if v89 == v161 {
		v181 = v158 + int32(4432660)
		goto L35
	} else {
		goto L58
	}
L57:
	;
	v190 = v91
	goto L34
L58:
	;
	v164 = v154 + int32(1)
	if v148 != v164 {
		v154 = v164
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v172 = v168 << (uint(int32(3)) % 32)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172)+uint32(_consts[261])))
	if v89 == v175 {
		v181 = v172 + int32(4433708)
		goto L35
	} else {
		goto L62
	}
L61:
	;
	v190 = v91
	goto L34
L62:
	;
	v178 = v168 + int32(1)
	if v120 != v178 {
		v168 = v178
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v194 = v190
	goto L11
L65:
	;
	if v194 == v14 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v201
	v203 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v203
	v207 = F_RelFileLocatorSkippingWAL(m, v8+int32(16))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L29
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v207
	goto L3
L68:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v220
	F_errmsg_internal(m, int32(41251), v8+int32(32))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L29
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(477505), int32(1380), int32(218577))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L29
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v236 + int32(4)
	F_errmsg_internal(m, int32(54355), v8)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L29
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(477505), int32(1398), int32(218577))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L29
	} else {
		goto L73
	}
L73:
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
	var v43 int64
	_ = v43
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
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int64
	_ = v237
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int64
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int64
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
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
	v43 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v42)+32)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v43
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
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13-int32(-64)))) = v62
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v64
	v68 = F_smgropen(m, v13+int32(56), v59)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	v85 = v56
	goto L13
L13:
	;
	v88 = v13 + int32(136)
	v90 = v13 + int32(124)
	v92 = v13 + int32(112)
	v93 = int32(1)
	v95 = F_smgrexists(m, v85, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L20
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v68
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+72))
	if v72 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v85 = v84
	goto L13
L16:
	;
	v80 = v72
	goto L18
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+76))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v68)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v68)+72))
	v80 = v78
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+72)) = v80 + int32(1)
	goto L15
L19:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v124 != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	if v95 == int32(0) {
		v119 = v93
		v120 = v3
		v121 = v88
		v122 = v90
		v123 = v92
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v99 = F_FreeSpaceMapPrepareTruncateRel(m, l0, l1)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v99
	if v99 == int32(-1) {
		v119 = v93
		v120 = v3
		v121 = v88
		v122 = v90
		v123 = v92
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v110 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+136)) = v110
	v114 = F_smgrnblocks(m, v42, v110)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+124)) = v114
	v119 = int32(2)
	v120 = v110
	v121 = v13 + int32(140)
	v122 = v13 + int32(128)
	v123 = v13 + int32(116)
	goto L19
L25:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v172 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L26:
	;
	v150 = v124
	goto L28
L27:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v126
	v128 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v128
	v132 = F_smgropen(m, v13+int32(40), v125)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	v152 = F_smgrexists(m, v150, int32(2))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L34
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v132
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)+72))
	if v136 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v150 = v148
	goto L28
L31:
	;
	v144 = v136
	goto L33
L32:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)+76))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+4)) = v138
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v132)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v132)+72))
	v144 = v142
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+72)) = v144 + int32(1)
	goto L30
L34:
	;
	if v152 == int32(0) {
		v170 = v119
		goto L25
	} else {
		goto L35
	}
L35:
	;
	v156 = F_visibilitymap_prepare_truncate(m, l0, l1)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v156
	if v156 == int32(-1) {
		v170 = v119
		goto L25
	} else {
		goto L37
	}
L37:
	;
	v161 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v161
	v164 = F_smgrnblocks(m, v42, v161)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v164
	v170 = v119 + int32(1)
	goto L25
L39:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v213)+120)) = v214 | int32(3)
	v218 = int32(4438516)
	v220 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v220 + int32(1)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+118)))
	if v225 != int32(112) {
		goto L51
	} else {
		goto L52
	}
L40:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v175 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v201 = v175
	goto L43
L42:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v177
	v179 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v179
	v183 = F_smgropen(m, v13+int32(24), v176)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L44
	}
L43:
	;
	v202 = int32(0)
	v204 = F_hash_search(m, v172, v201, v202, v202)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L49
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v183
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)+72))
	if v187 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v201 = v199
	goto L43
L46:
	;
	v195 = v187
	goto L48
L47:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183)+76))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v183)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v183)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v183)+72))
	v195 = v193
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+72)) = v195 + int32(1)
	goto L45
L49:
	;
	if v204 == int32(0) {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	v208 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v204)+12)) = uint8(v208)
	goto L39
L51:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v254 != 0 {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v229 <= int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v232 != 0 {
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
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = v235
	v237 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+92)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = int32(7)
	F_XLogBeginInsert(m)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v233 != 0 {
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
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v250 = F_XLogInsert(m, int32(2), int32(33))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	F_XLogFlush(m, v250)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	goto L51
L62:
	;
	v280 = v254
	goto L64
L63:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v256
	v258 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v258
	v262 = F_smgropen(m, v13+int32(8), v255)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L65
	}
L64:
	;
	F_smgrtruncate(m, v280, v13+int32(132), v170, v13+int32(120), v13+int32(108))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L4
	} else {
		goto L70
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v262
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v262)+72))
	if v266 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v280 = v278
	goto L64
L67:
	;
	v274 = v266
	goto L69
L68:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v262)+76))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v262)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v267)+4)) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v262)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v268))) = v270
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v262)+72))
	v274 = v272
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+72)) = v274 + int32(1)
	goto L66
L70:
	;
	v289 = int32(4438516)
	v291 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v291 - int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v296)+120)) = v297 & int32(-4)
	if v120 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_FreeSpaceMapVacuumRange(m, l0, l1, int32(-1))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
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
							F_sequence_close(m, v12, int32(3))
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
							F_sequence_close(m, v12, int32(3))
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
					F_errmsg_internal(m, int32(43981), v8)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						F_errfinish(m, int32(473018), int32(3664), int32(121869))
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v10
	v15 = F_LockRelease(m, v6, l1, int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
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
	var v35 int32
	_ = v35
	var v51 int32
	_ = v51
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v79 = v11
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v79 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v79 = int32(0)
				} else {
					v79 = v11
				}
			}
		} else {
			v23 = l0 - int32(2671)
			if base.Ui32(int32(27)) < base.Ui32(v23) {
				if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
					v79 = v11
				} else {
					if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
						v79 = v11
					} else {
						v79 = int32(0)
					}
				}
			} else {
				if int32(1)<<(uint(v23)%32)&int32(226492515) == int32(0) {
					if base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)) {
						v79 = v11
					} else {
						if base.Ui32(l0-int32(2846)) < base.Ui32(int32(2)) {
							v79 = v11
						} else {
							v79 = int32(0)
						}
					}
				} else {
					v79 = v11
				}
			}
		}
	} else {
		if l0 <= int32(5999) {
			v35 = l0 - int32(4177)
			if base.Ui32(int32(9)) < base.Ui32(v35) {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v79 = v11
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v79 = int32(0)
					} else {
						v79 = v11
					}
				}
			} else {
				if int32(1)<<(uint(v35)%32)&int32(963) == int32(0) {
					if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
						v79 = v11
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
							v79 = int32(0)
						} else {
							v79 = v11
						}
					}
				} else {
					v79 = v11
				}
			}
		} else {
			switch l0 - int32(6243) {
			case 0, 1, 2, 3, 4, 59, 60:
				v79 = v11
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v79 = int32(0)
			default:
				if base.Ui32(l0-int32(6000)) < base.Ui32(int32(3)) {
					v79 = v11
				} else {
					v51 = l0 - int32(6100)
					if base.Ui32(int32(15)) < base.Ui32(v51) {
						v79 = int32(0)
					} else {
						if int32(1)<<(uint(v51)%32)&int32(49153) != 0 {
							v79 = v11
						} else {
							v79 = int32(0)
						}
					}
				}
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
	v85 = *(*int32)(unsafe.Add(mBase, _consts[226]))
	if v79 != 0 {
		v86 = int32(0)
	} else {
		v86 = v85
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v86
	v89 = F_LockRelease(m, v7, l1, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
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
	var v113 int32
	_ = v113
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
																							v113 = m.ExcPending
																							if v113 != 0 {
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
																				v113 = m.ExcPending
																				if v113 != 0 {
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
																				v113 = m.ExcPending
																				if v113 != 0 {
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
																	v113 = m.ExcPending
																	if v113 != 0 {
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
																				v113 = m.ExcPending
																				if v113 != 0 {
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
																	v113 = m.ExcPending
																	if v113 != 0 {
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
																	v113 = m.ExcPending
																	if v113 != 0 {
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
														v113 = m.ExcPending
														if v113 != 0 {
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
																				v113 = m.ExcPending
																				if v113 != 0 {
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
																	v113 = m.ExcPending
																	if v113 != 0 {
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
																	v113 = m.ExcPending
																	if v113 != 0 {
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
														v113 = m.ExcPending
														if v113 != 0 {
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
																	v113 = m.ExcPending
																	if v113 != 0 {
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
														v113 = m.ExcPending
														if v113 != 0 {
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
														v113 = m.ExcPending
														if v113 != 0 {
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
											v113 = m.ExcPending
											if v113 != 0 {
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
	var v51 int32
	_ = v51
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
	F_sequence_close(m, v18, int32(3))
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
	F_sequence_close(m, v44, int32(3))
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
	v50 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v51 <= v50 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v55 = v50
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
	v98 = *(*int32)(unsafe.Add(mBase, _consts[206]))
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
	F_errmsg_internal(m, int32(665005), v14)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(473018), int32(18421), int32(9331))
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
	F_errmsg_internal(m, int32(37833), v14+int32(16))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(473018), int32(18443), int32(9331))
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 float32
	_ = v354
	var v355 float32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
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
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L229
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L226
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L223
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L220
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L214
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L208
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L205
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L202
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L199
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L196
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L193
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L190
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L187
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L184
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
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L181
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+88))
	v45 = v32 + v33
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+84))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+88))
	if v47 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v303 = F_relation_open(m, l0, int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L103
	}
L24:
	;
	if v44|v47 != 0 {
		goto L15
	} else {
		goto L28
	}
L25:
	;
	if v44 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+88)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v42)+88)) = v47
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v45)+92))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+92)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v42)+92)) = v54
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+84))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v42)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+84)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v42)+84)) = v58
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+118)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+118)))
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+118)) = uint8(v63)
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+118)) = uint8(v62)
	if l3 != 0 {
		v299 = l7
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v45)+112))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v42)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+112)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v42)+112)) = v66
	v299 = l7
	goto L23
L28:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v45)+92))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)+92))
	if v71 != v72 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+118)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+118)))
	if v74 != v75 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	if v43 != v46 {
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
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v45)+112))
	if v80 != 0 {
		goto L11
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+117)))
	v83 = int32(0)
	if v82 == v83 {
		goto L41
	} else {
		goto L42
	}
L35:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v42)+112))
	if v81 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	if v182 == int32(0) {
		goto L10
	} else {
		goto L68
	}
L38:
	;
	goto L37
L39:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v182 = v177
	goto L38
L40:
	;
	v160 = int32(0)
	goto L64
L41:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if int32(0) < v89 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	if int32(0) < v117 {
		goto L52
	} else {
		goto L53
	}
L44:
	;
	v94 = v83
	goto L47
L45:
	;
	goto L46
L46:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[256]))
	if v112 <= int32(0) {
		v182 = v83
		goto L38
	} else {
		goto L51
	}
L47:
	;
	v98 = v94 << (uint(int32(3)) % 32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)+uint32(_consts[257])))
	if l0 == v101 {
		v173 = v98 + int32(4433184)
		goto L39
	} else {
		goto L49
	}
L48:
	;
	goto L46
L49:
	;
	v104 = v94 + int32(1)
	if v104 != v89 {
		v94 = v104
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	goto L40
L52:
	;
	v122 = v83
	goto L55
L53:
	;
	goto L54
L54:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	if v140 <= int32(0) {
		v182 = v83
		goto L38
	} else {
		goto L59
	}
L55:
	;
	v126 = v122 << (uint(int32(3)) % 32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)+uint32(_consts[259])))
	if l0 == v129 {
		v173 = v126 + int32(4432136)
		goto L39
	} else {
		goto L57
	}
L56:
	;
	goto L54
L57:
	;
	v132 = v122 + int32(1)
	if v132 != v117 {
		v122 = v132
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v146 = int32(0)
	goto L60
L60:
	;
	v150 = v146 << (uint(int32(3)) % 32)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)+uint32(_consts[260])))
	if l0 == v153 {
		v173 = v150 + int32(4432660)
		goto L39
	} else {
		goto L62
	}
L61:
	;
	v182 = v83
	goto L38
L62:
	;
	v156 = v146 + int32(1)
	if v140 != v156 {
		v146 = v156
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v164 = v160 << (uint(int32(3)) % 32)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)+uint32(_consts[261])))
	if l0 == v167 {
		v173 = v164 + int32(4433708)
		goto L39
	} else {
		goto L66
	}
L65:
	;
	v182 = v83
	goto L38
L66:
	;
	v170 = v160 + int32(1)
	if v112 != v170 {
		v160 = v170
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+117)))
	v186 = int32(0)
	if v185 == v186 {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	if v285 == int32(0) {
		goto L9
	} else {
		goto L100
	}
L70:
	;
	goto L69
L71:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v285 = v280
	goto L70
L72:
	;
	v263 = int32(0)
	goto L96
L73:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if int32(0) < v192 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	if int32(0) < v220 {
		goto L84
	} else {
		goto L85
	}
L76:
	;
	v197 = v186
	goto L79
L77:
	;
	goto L78
L78:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _consts[256]))
	if v215 <= int32(0) {
		v285 = v186
		goto L70
	} else {
		goto L83
	}
L79:
	;
	v201 = v197 << (uint(int32(3)) % 32)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v201)+uint32(_consts[257])))
	if l1 == v204 {
		v276 = v201 + int32(4433184)
		goto L71
	} else {
		goto L81
	}
L80:
	;
	goto L78
L81:
	;
	v207 = v197 + int32(1)
	if v207 != v192 {
		v197 = v207
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	goto L72
L84:
	;
	v225 = v186
	goto L87
L85:
	;
	goto L86
L86:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	if v243 <= int32(0) {
		v285 = v186
		goto L70
	} else {
		goto L91
	}
L87:
	;
	v229 = v225 << (uint(int32(3)) % 32)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v229)+uint32(_consts[259])))
	if l1 == v232 {
		v276 = v229 + int32(4432136)
		goto L71
	} else {
		goto L89
	}
L88:
	;
	goto L86
L89:
	;
	v235 = v225 + int32(1)
	if v235 != v220 {
		v225 = v235
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v249 = int32(0)
	goto L92
L92:
	;
	v253 = v249 << (uint(int32(3)) % 32)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v253)+uint32(_consts[260])))
	if l1 == v256 {
		v276 = v253 + int32(4432660)
		goto L71
	} else {
		goto L94
	}
L93:
	;
	v285 = v186
	goto L70
L94:
	;
	v259 = v249 + int32(1)
	if v243 != v259 {
		v249 = v259
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v267 = v263 << (uint(int32(3)) % 32)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v267)+uint32(_consts[261])))
	if l1 == v270 {
		v276 = v267 + int32(4433708)
		goto L71
	} else {
		goto L98
	}
L97:
	;
	v285 = v186
	goto L70
L98:
	;
	v273 = v263 + int32(1)
	if v215 != v273 {
		v263 = v273
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+117)))
	F_RelationMapUpdateMap(m, l0, v285, v288, int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+117)))
	F_RelationMapUpdateMap(m, l1, v182, v292, int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = l1
	v299 = l7 + int32(4)
	goto L23
L103:
	;
	v306 = F_relation_open(m, l1, int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v303)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+32)) = v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v303)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+36)) = v310
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v303)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+40)) = v312
	v315 = F_GetCurrentSubTransactionId(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v303)+36)) = v315
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v303)+40))
	if v317 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	F_relation_close(m, v303, int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L112
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v303)+40)) = v315
	goto L108
L107:
	;
	goto L108
L108:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	if v322 <= int32(31) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v303)+56))
	*(*int32)(unsafe.Add(mBase, _consts[262])) = v322 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v322<<(uint(int32(2))%32))+uint32(_consts[263]))) = v325
	goto L105
L110:
	;
	goto L111
L111:
	;
	v336 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[264])) = uint8(v336)
	goto L105
L112:
	;
	F_relation_close(m, v306, int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+119)))
	if v345 != int32(105) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+140)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v45)+136)) = l5
	goto L116
L115:
	;
	goto L116
L116:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v45)+96))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+96)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v42)+96)) = v350
	v354 = *(*float32)(unsafe.Add(mBase, uint32(v45)+100))
	v355 = *(*float32)(unsafe.Add(mBase, uint32(v42)+100))
	*(*float32)(unsafe.Add(mBase, uint32(v45)+100)) = v355
	*(*float32)(unsafe.Add(mBase, uint32(v42)+100)) = v354
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v45)+104))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v42)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+104)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v42)+104)) = v358
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v45)+108))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v42)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+108)) = v363
	*(*int32)(unsafe.Add(mBase, uint32(v42)+108)) = v362
	if l2 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	if v43 != v46 {
		goto L127
	} else {
		goto L128
	}
L118:
	;
	v370 = F_CatalogOpenIndexes(m, v26)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	F_CacheInvalidateRelcacheByTuple(m, v30)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L125
	}
L121:
	;
	F_CatalogTupleUpdateWithInfo(m, v26, v30+int32(4), v30, v370)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_CatalogTupleUpdateWithInfo(m, v26, v36+int32(4), v36, v370)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_CatalogCloseIndexes(m, v370)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L117
L125:
	;
	F_CacheInvalidateRelcacheByTuple(m, v36)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	goto L117
L127:
	;
	v388 = F_changeDependencyFor(m, int32(1259), l0, int32(2601), v46, v43)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v399 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L130:
	;
	if v388 != int32(1) {
		goto L8
	} else {
		goto L131
	}
L131:
	;
	v394 = F_changeDependencyFor(m, int32(1259), l1, int32(2601), v43, v46)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	if v394 != int32(1) {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	goto L129
L134:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v45)+112))
	if v417 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L135:
	;
	v403 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l0, v403, v403, l4)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v408 == int32(0) {
		goto L134
	} else {
		goto L137
	}
L137:
	;
	v412 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l1, v412, v412, int32(1))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	goto L134
L139:
	;
	F_pfree(m, v30)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L178
	}
L140:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+119)))
	if v492 != int32(116) {
		goto L139
	} else {
		goto L173
	}
L141:
	;
	if l3 == int32(0) {
		goto L139
	} else {
		goto L172
	}
L142:
	;
	v434 = int32(1)
	if base.Ui32(l0) < base.Ui32(int32(12000)) {
		v442 = v434
		goto L152
	} else {
		goto L153
	}
L143:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v42)+112))
	if v420 == int32(0) {
		goto L141
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	if l3 == int32(0) {
		goto L142
	} else {
		goto L148
	}
L146:
	;
	if l3 == int32(0) {
		goto L142
	} else {
		goto L147
	}
L147:
	;
	goto L3
L148:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v42)+112))
	if v427 == int32(0) {
		goto L3
	} else {
		goto L149
	}
L149:
	;
	F_swap_relation_files(m, v417, v427, l2, int32(1), l4, l5, l6, v299)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	goto L140
L151:
	;
	if v442 != 0 {
		goto L6
	} else {
		goto L155
	}
L152:
	;
	goto L151
L153:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
	if v437 == int32(99) {
		v442 = v434
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v440 = F_isTempToastNamespace(m, v437)
	mBase = m.M
	v442 = v440
	goto L152
L155:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v45)+112))
	if v443 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v446 = F_deleteDependencyRecordsFor(m, int32(1259), v443, int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v42)+112))
	if v451 != 0 {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	if v446 != int32(1) {
		goto L5
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v454 = F_deleteDependencyRecordsFor(m, int32(1259), v451, int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	v459 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+220)) = v459
	v461 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v461
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v45)+112))
	if v467 != 0 {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	if v454 != int32(1) {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+204)) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = l0
	F_recordDependencyOn(m, v22+int32(200), v22+int32(212), int32(105))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v42)+112))
	if v477 == int32(0) {
		goto L139
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+204)) = v477
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = l1
	F_recordDependencyOn(m, v22+int32(200), v22+int32(212), int32(105))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	goto L139
L172:
	;
	goto L140
L173:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
	if v495 != int32(116) {
		goto L139
	} else {
		goto L174
	}
L174:
	;
	v499 = F_toast_get_valid_index(m, l0, int32(8))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v502 = F_toast_get_valid_index(m, l1, int32(8))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v505 = int32(0)
	F_swap_relation_files(m, v499, v502, l2, int32(1), l4, v505, v505, v299)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	goto L139
L178:
	;
	F_pfree(m, v36)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	F_sequence_close(m, v26, int32(3))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	m.G0 = v22 + int32(224)
	return
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = l0
	F_errmsg_internal(m, int32(43981), v22)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(473708), int32(1087), int32(155617))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l1
	F_errmsg_internal(m, int32(43981), v22+int32(16))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(473708), int32(1092), int32(155617))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v45 + int32(4)
	F_errmsg_internal(m, int32(252167), v22+int32(96))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(473708), int32(1142), int32(155617))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v45 + int32(4)
	F_errmsg_internal(m, int32(668904), v22+int32(192))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(473708), int32(1153), int32(155617))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v45 + int32(4)
	F_errmsg_internal(m, int32(668854), v22+int32(176))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(473708), int32(1156), int32(155617))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v45 + int32(4)
	F_errmsg_internal(m, int32(668953), v22+int32(160))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(473708), int32(1159), int32(155617))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v45 + int32(4)
	F_errmsg_internal(m, int32(668802), v22+int32(144))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(473708), int32(1163), int32(155617))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v45 + int32(4)
	F_errmsg_internal(m, int32(54355), v22+int32(112))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(473708), int32(1171), int32(155617))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+132)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v42 + int32(4)
	F_errmsg_internal(m, int32(54355), v22+int32(128))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(473708), int32(1175), int32(155617))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	v675 = F_get_rel_namespace(m, l0)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v677 = F_get_namespace_name(m, v675)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v679 = F_get_rel_name(m, l0)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v679
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v677
	F_errmsg_internal(m, int32(651118), v22+int32(80))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(473708), int32(1289), int32(155617))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	v697 = F_get_rel_namespace(m, l1)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v699 = F_get_namespace_name(m, v697)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	v701 = F_get_rel_name(m, l1)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v701
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v699
	F_errmsg_internal(m, int32(651118), v22-int32(-64))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(473708), int32(1297), int32(155617))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	F_errmsg_internal(m, int32(146355), int32(0))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(473708), int32(1359), int32(155617))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v446
	F_errmsg_internal(m, int32(413061), v22+int32(48))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	F_errfinish(m, int32(473708), int32(1369), int32(155617))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v454
	F_errmsg_internal(m, int32(413061), v22+int32(32))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(473708), int32(1378), int32(155617))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	F_errmsg_internal(m, int32(356305), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(473708), int32(1332), int32(155617))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
