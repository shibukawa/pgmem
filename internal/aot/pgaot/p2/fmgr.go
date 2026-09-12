package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fmgr_info_cxt_security(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
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
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
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
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v315 int32
	_ = v315
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v5
	v22 = *(*int32)(unsafe.Add(mBase, _consts[1213]))
	if base.Ui32(v22) < base.Ui32(l0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L10
	} else {
		goto L88
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L10
	} else {
		goto L85
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L10
	} else {
		goto L82
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L10
	} else {
		goto L79
	}
L5:
	;
	m.G0 = v12 + int32(112)
	return
L6:
	;
	v46 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[1214]))))
	if v28 == int32(65535) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v32 = v28 << (uint(int32(4)) % 32)
	if v32+int32(1732448) == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v37 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v37)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_consts[1215])))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_consts[1216])))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
	goto L5
L10:
	;
	return
L11:
	;
	if v46 == int32(0) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+22)))
	v52 = v50 + v51
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+104)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v53)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+99)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)) = uint8(v55)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)) = uint8(v57)
	if l3 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	F_ReleaseCatCache(m, v46)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L10
	} else {
		goto L78
	}
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v52)+76))
	switch v81 - int32(12) {
	case 0:
		goto L27
	case 1:
		goto L26
	case 2:
		goto L25
	default:
		goto L24
	}
L15:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+97)))
	if v59 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v75 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v75)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1626)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
	goto L13
L17:
	;
	v62 = F_heap_attisnull(m, v46, int32(29), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	if v62 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[1217]))
	if v67 == int32(0) {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v70 = m.T0[v67].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	if v70 == int32(0) {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L16
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v296)
	goto L13
L24:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+22)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v270+v271)+76))
	v274 = F_SearchSysCache1(m, int32(36), v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L10
	} else {
		goto L74
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1627)
	v296 = int32(1)
	goto L23
L26:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+22)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148+v149)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v151
	v154 = *(*int32)(unsafe.Add(mBase, _consts[1218]))
	if v154 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L27:
	;
	v87 = F_SysCacheGetAttrNotNull(m, int32(47), v46, int32(26))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v89 = F_text_to_cstring(m, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[1219]))
	if v92 <= int32(0) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v97 = int32(0)
	goto L31
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v97<<(uint(int32(4))%32))+uint32(_consts[1220])))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v112 == int32(0) {
		v131 = v111
		v132 = v112
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v138 = v97 << (uint(int32(4)) % 32)
	if v138+int32(1732448) == int32(0) {
		goto L1
	} else {
		goto L45
	}
L33:
	;
	if v132-v131 != 0 {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	goto L33
L35:
	;
	if v111 != v112 {
		v131 = v111
		v132 = v112
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v116 = v89
	v117 = v108
	goto L37
L37:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v121 == int32(0) {
		v131 = v120
		v132 = v121
		goto L34
	} else {
		goto L39
	}
L38:
	;
	v131 = v120
	v132 = v121
	goto L34
L39:
	;
	v124 = int32(1)
	if v120 == v121 {
		v116 = v116 + v124
		v117 = v117 + v124
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v135 = v97 + int32(1)
	if v135 != v92 {
		v97 = v135
		goto L31
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	goto L32
L44:
	;
	goto L1
L45:
	;
	F_pfree(m, v89)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v138)+uint32(_consts[1216])))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v145
	v296 = int32(2)
	goto L23
L47:
	;
	v261 = int32(1)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	if v262 != v261 {
		goto L3
	} else {
		goto L73
	}
L48:
	;
	v197 = F_SysCacheGetAttrNotNull(m, int32(47), v46, int32(26))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L10
	} else {
		goto L60
	}
L49:
	;
	v159 = int32(0)
	v161 = F_hash_search(m, v154, v12+int32(60), v159, v159)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	if v161 == int32(0) {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	if v165 != v167 {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v170 = v161 + int32(8)
	v172 = v46 + int32(4)
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+2)))
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170))))
	v175 = int32(16)
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172)+2)))
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172))))
	if v173|v174<<(uint(v175)%32) == v178|v179<<(uint(v175)%32) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	if v189 == int32(0) {
		goto L48
	} else {
		goto L59
	}
L54:
	;
	goto L53
L55:
	;
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+4)))
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172)+4)))
	if v185 == v186 {
		v189 = int32(1)
		goto L54
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v189 = int32(0)
	goto L54
L58:
	;
	goto L57
L59:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v161)+20))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
	v258 = v192
	v259 = v193
	goto L47
L60:
	;
	v199 = F_text_to_cstring(m, v197)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	v203 = F_SysCacheGetAttrNotNull(m, int32(47), v46, int32(27))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v205 = F_text_to_cstring(m, v203)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	v210 = F_load_external_function(m, v205, v199, int32(1), v12+int32(56))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v213 = F_fetch_finfo_record(m, v212, v199)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+22)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v215+v216)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v218
	v221 = *(*int32)(unsafe.Add(mBase, _consts[1218]))
	if v221 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+76)) = int64(103079215108)
	v232 = F_hash_create(m, int32(310753), int32(100), v12+int32(60), int32(40))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L10
	} else {
		goto L69
	}
L67:
	;
	v235 = v221
	goto L68
L68:
	;
	v241 = F_hash_search(m, v235, v12+int32(108), int32(1), v12+int32(60))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L10
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1218])) = v232
	v235 = v232
	goto L68
L70:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	*(*int32)(unsafe.Add(mBase, uint32(v241)+4)) = v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v241)+8)) = v246
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v241)+12)) = uint16(v248)
	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v241)+16)) = v210
	F_pfree(m, v199)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L10
	} else {
		goto L71
	}
L71:
	;
	F_pfree(m, v205)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	v258 = v213
	v259 = v210
	goto L47
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v259
	v296 = v261
	goto L23
L74:
	;
	if v274 == int32(0) {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v274)+16))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+22)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v279+v280)+76))
	v286 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	F_fmgr_info_cxt_security(m, v282, v12+int32(60), v286, int32(1))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v290
	F_ReleaseCatCache(m, v274)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L10
	} else {
		goto L77
	}
L77:
	;
	v296 = int32(0)
	goto L23
L78:
	;
	goto L5
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg_internal(m, int32(42736), v12)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(476043), int32(183), int32(9973))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L10
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
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v345
	F_errmsg_internal(m, int32(464454), v12+int32(48))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(476043), int32(408), int32(323287))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L10
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v273
	F_errmsg_internal(m, int32(50006), v12+int32(16))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(476043), int32(428), int32(323266))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v89
	F_errmsg(m, int32(377288), v12+int32(32))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(476043), int32(237), int32(9973))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fmgr_sql_validator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
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
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v263 int32
	_ = v263
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v407 int32
	_ = v407
	var v422 int32
	_ = v422
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_CheckFunctionValidatorAccess(m, v18, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L118
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L113
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L110
	}
L4:
	;
	return int32(0)
L5:
	;
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = F_SearchSysCache1(m, int32(47), v19)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	m.G0 = v15 + int32(80)
	return int32(0)
L9:
	;
	if v25 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
	v31 = v29 + v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	v33 = F_get_typtype(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31)+104)))
	if int32(0) < v57 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	if v33 != int32(112) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	if v37 <= int32(3830) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	switch v37 - int32(2249) {
	case 0, 28, 29, 34:
		goto L11
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 30, 31, 32, 33:
		goto L1
	default:
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(v37-int32(5077)) < base.Ui32(int32(4)) {
		goto L11
	} else {
		goto L20
	}
L17:
	;
	if v37 == int32(2776) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	if v37 == int32(3500) {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	if base.Ui32(v37-int32(4537)) < base.Ui32(int32(2)) {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	if v37 != int32(3831) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L11
L23:
	;
	v63 = int32(0)
	v70 = v2
	goto L26
L24:
	;
	v116 = v2
	goto L25
L25:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, _consts[292])))
	if v122 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	v77 = v31 + int32(136) + v63<<(uint(int32(2))%32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = F_get_typtype(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	v116 = v104
	goto L25
L28:
	;
	v106 = v63 + int32(1)
	v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31)+104)))
	if v106 < v107 {
		v63 = v106
		v70 = v104
		goto L26
	} else {
		goto L40
	}
L29:
	;
	if v79 != int32(112) {
		v104 = v70
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v83 = int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v84 <= int32(3830) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	switch v84 - int32(2277) {
	case 0, 6:
		v104 = v83
		goto L28
	case 1, 2, 3, 4, 5:
		goto L2
	default:
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if base.Ui32(v84-int32(5077)) < base.Ui32(int32(4)) {
		v104 = v83
		goto L28
	} else {
		goto L37
	}
L34:
	;
	if v84 == int32(2776) {
		v104 = v83
		goto L28
	} else {
		goto L35
	}
L35:
	;
	if v84 == int32(3500) {
		v104 = v83
		goto L28
	} else {
		goto L36
	}
L36:
	;
	goto L2
L37:
	;
	if base.Ui32(v84-int32(4537)) < base.Ui32(int32(2)) {
		v104 = v83
		goto L28
	} else {
		goto L38
	}
L38:
	;
	if v84 != int32(3831) {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v104 = v83
	goto L28
L40:
	;
	goto L27
L41:
	;
	v127 = F_SysCacheGetAttrNotNull(m, int32(47), v25, int32(26))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L4
	} else {
		goto L109
	}
L44:
	;
	v129 = F_text_to_cstring(m, v127)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v31 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = int32(473)
	v137 = int32(4442424)
	v138 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	*(*int32)(unsafe.Add(mBase, _consts[84])) = v15 + int32(56)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v15 + int32(68)
	v151 = F_SysCacheGetAttr(m, int32(47), v25, int32(28), v15+int32(79))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+79)))
	if v153 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	*(*int32)(unsafe.Add(mBase, _consts[84])) = v407
	goto L43
L48:
	;
	v275 = int32(0)
	if v263 == v275 {
		goto L83
	} else {
		goto L84
	}
L49:
	;
	if v116 != 0 {
		goto L47
	} else {
		goto L82
	}
L50:
	;
	v156 = F_text_to_cstring(m, v151)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v211 = F_pg_parse_query(m, v129)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L72
	}
L53:
	;
	if v172 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	v158 = F_stringToNode(m, v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v160 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v172 = v164
	goto L53
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v158
	v170 = F_list_make1_impl(m, int32(1), v15+int32(40))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v172 = v170
	goto L53
L60:
	;
	v251 = int32(0)
	goto L49
L61:
	;
	goto L62
L62:
	;
	v176 = int32(0)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v177 <= v176 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v251 = int32(0)
	goto L49
L64:
	;
	goto L65
L65:
	;
	v182 = int32(0)
	v185 = v176
	goto L66
L66:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194+v185<<(uint(int32(2))%32))))
	F_AcquireRewriteLocks(m, v198, int32(1), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L68
	}
L67:
	;
	v251 = v205
	goto L49
L68:
	;
	v203 = F_pg_rewrite_query(m, v198)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v205 = F_lappend(m, v182, v203)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v208 = v185 + int32(1)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v208 < v209 {
		v182 = v205
		v185 = v208
		goto L66
	} else {
		goto L71
	}
L71:
	;
	goto L67
L72:
	;
	if v116 != 0 {
		goto L47
	} else {
		goto L73
	}
L73:
	;
	v213 = int32(0)
	v216 = F_prepare_sql_fn_parse_info(m, v25, v213, v213)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	if v211 == int32(0) {
		v263 = v213
		goto L48
	} else {
		goto L75
	}
L75:
	;
	v220 = int32(0)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if v221 <= v220 {
		v251 = v213
		goto L49
	} else {
		goto L76
	}
L76:
	;
	v224 = v213
	v227 = v220
	goto L77
L77:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v236+v227<<(uint(int32(2))%32))))
	v243 = F_pg_analyze_and_rewrite_withcb(m, v240, v129, int32(474), v216, int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L79
	}
L78:
	;
	v251 = v245
	goto L49
L79:
	;
	v245 = F_lappend(m, v224, v243)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	v248 = v227 + int32(1)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if v248 < v249 {
		v224 = v245
		v227 = v248
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	v263 = v251
	goto L48
L83:
	;
	v380 = int32(0)
	v386 = F_internal_get_result_type(m, v19, v380, v380, v15+int32(48), v15+int32(44))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L107
	}
L84:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v278 <= int32(0) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v281 = int32(0)
	if v281 < v278 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v284 = v278
	goto L88
L87:
	;
	v284 = v281
	goto L88
L88:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v290 = v275
	goto L90
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L103
	}
L90:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v285+v290<<(uint(int32(2))%32))))
	if v301 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L83
L92:
	;
	v350 = v290 + int32(1)
	if v350 != v284 {
		v290 = v350
		goto L90
	} else {
		goto L102
	}
L93:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if v304 <= int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	v311 = int32(0)
	goto L95
L95:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v307+v311<<(uint(int32(2))%32))))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
	if v325 != int32(6) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L92
L97:
	;
	v335 = v311 + int32(1)
	if v304 != v335 {
		v311 = v335
		goto L95
	} else {
		goto L101
	}
L98:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v324)+28))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	if v329 != int32(213) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	if v332 != 0 {
		goto L89
	} else {
		goto L100
	}
L100:
	;
	goto L97
L101:
	;
	goto L96
L102:
	;
	goto L91
L103:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	F_errmsg(m, int32(133239), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(474739), int32(2075), int32(90389))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v390 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31)+96)))
	v392 = F_check_sql_fn_retval(m, v263, v388, v389, v390, int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	goto L47
L109:
	;
	goto L8
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
	F_errmsg_internal(m, int32(42736), v15)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(480601), int32(853), int32(200756))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v461 = F_format_type_be(m, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v461
	F_errmsg(m, int32(183373), v15+int32(32))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(480601), int32(880), int32(200756))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	v482 = F_format_type_be(m, v481)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v482
	F_errmsg(m, int32(182834), v15+int32(16))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(480601), int32(865), int32(200756))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
