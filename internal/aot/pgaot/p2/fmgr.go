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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int64
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int64
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
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
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v311 int32
	_ = v311
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v5
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[0]))
	if base.Ui32(v22) < base.Ui32(l0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L9
	} else {
		goto L86
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L9
	} else {
		goto L83
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L9
	} else {
		goto L79
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L9
	} else {
		goto L76
	}
L5:
	;
	m.G0 = v12 + int32(112)
	return
L6:
	;
	v43 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_fmgr_info_cxt_security[1]))))
	if v26 == int32(_a_F_fmgr_info_cxt_security_0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v29 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v29)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
	v33 = v26 << (uint(int32(4)) % 32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_fmgr_info_cxt_security[2])))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_fmgr_info_cxt_security[3])))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v38
	goto L5
L9:
	;
	return
L10:
	;
	if v43 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+22)))
	v49 = v47 + v48
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+104)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v50)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+99)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)) = uint8(v52)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)) = uint8(v54)
	if l3 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	F_ReleaseCatCache(m, v43)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L9
	} else {
		goto L75
	}
L13:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v49)+76))
	switch v78 - int32(12) {
	case 0:
		goto L26
	case 1:
		goto L25
	case 2:
		goto L24
	default:
		goto L23
	}
L14:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+97)))
	if v56 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v72 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v72)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1826)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
	goto L12
L16:
	;
	v59 = F_heap_attisnull(m, v43, int32(29), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	if v59 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[4]))
	if v64 == int32(0) {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v67 = m.T0[v64].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if v67 == int32(0) {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v292)
	goto L12
L23:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+22)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v265+v266)+76))
	v270 = F_SearchSysCache1(m, int32(36), base.I64_extend_i32_u(v268))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L9
	} else {
		goto L71
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1827)
	v292 = int32(1)
	goto L22
L25:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+22)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141+v142)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v144
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[5]))
	if v147 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L26:
	;
	v84 = F_SysCacheGetAttrNotNull(m, int32(47), v43, int32(26))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v87 = F_text_to_cstring(m, base.I32_wrap_i64(v84))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[6]))
	if v90 <= int32(0) {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v95 = int32(0)
	goto L30
L30:
	;
	v103 = v95 << (uint(int32(4)) % 32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+uint32(_c_F_fmgr_info_cxt_security[7])))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if base.B2i32(v109 == int32(0))|base.B2i32(v109 != v112) != 0 {
		v130 = v109
		v131 = v112
		goto L33
	} else {
		goto L34
	}
L31:
	;
	F_pfree(m, v87)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L43
	}
L32:
	;
	if v130-v131 != 0 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	goto L32
L34:
	;
	v115 = v87
	v116 = v106
	goto L35
L35:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	if v120 == int32(0) {
		v130 = v120
		v131 = v119
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v130 = v120
	v131 = v119
	goto L33
L37:
	;
	v123 = int32(1)
	if v120 == v119 {
		v115 = v115 + v123
		v116 = v116 + v123
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v134 = v95 + int32(1)
	if v90 != v134 {
		v95 = v134
		goto L30
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L31
L42:
	;
	goto L3
L43:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v103)+uint32(_c_F_fmgr_info_cxt_security[3])))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v138
	v292 = int32(2)
	goto L22
L44:
	;
	v256 = int32(1)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	if v257 != v256 {
		goto L2
	} else {
		goto L70
	}
L45:
	;
	v190 = F_SysCacheGetAttrNotNull(m, int32(47), v43, int32(26))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L9
	} else {
		goto L57
	}
L46:
	;
	v152 = int32(0)
	v154 = F_hash_search(m, v147, v12+int32(56), v152, v152)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	if v154 == int32(0) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v158 != v160 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v163 = v154 + int32(8)
	v165 = v43 + int32(4)
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+2)))
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163))))
	v168 = int32(16)
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+2)))
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165))))
	if v166|v167<<(uint(v168)%32) == v171|v172<<(uint(v168)%32) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v182 == int32(0) {
		goto L45
	} else {
		goto L56
	}
L51:
	;
	goto L50
L52:
	;
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+4)))
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+4)))
	if v178 == v179 {
		v182 = int32(1)
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v182 = int32(0)
	goto L51
L55:
	;
	goto L54
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v154)+20))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	v252 = v185
	v253 = v186
	goto L44
L57:
	;
	v193 = F_text_to_cstring(m, base.I32_wrap_i64(v190))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v197 = F_SysCacheGetAttrNotNull(m, int32(47), v43, int32(27))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	v200 = F_text_to_cstring(m, base.I32_wrap_i64(v197))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	v205 = F_load_external_function(m, v200, v193, int32(1), v12+int32(52))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v208 = F_fetch_finfo_record(m, v207, v193)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+22)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v210+v211)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v213
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[5]))
	if v216 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = int64(103079215108)
	v227 = F_hash_create(m, int32(_a_F_fmgr_info_cxt_security_1), int64(100), v12+int32(56), int32(40))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L9
	} else {
		goto L66
	}
L64:
	;
	v230 = v216
	goto L65
L65:
	;
	v236 = F_hash_search(m, v230, v12+int32(108), int32(1), v12+int32(56))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L9
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[5])) = v227
	v230 = v227
	goto L65
L67:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+4)) = v239
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+8)) = v241
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v236)+12)) = uint16(v243)
	*(*int32)(unsafe.Add(mBase, uint32(v236)+20)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v236)+16)) = v205
	F_pfree(m, v193)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v200)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v252 = v208
	v253 = v205
	goto L44
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v253
	v292 = v256
	goto L22
L71:
	;
	if v270 == int32(0) {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v270)+16))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+22)))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275+v276)+76))
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[8]))
	F_fmgr_info_cxt_security(m, v278, v12+int32(56), v282, int32(1))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v286
	F_ReleaseCatCache(m, v270)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	v292 = int32(0)
	goto L22
L75:
	;
	goto L5
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg_internal(m, int32(_a_F_fmgr_info_cxt_security_2), v12)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_fmgr_info_cxt_security_3), int32(185), int32(_a_F_fmgr_info_cxt_security_4))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L9
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
	F_errcode(m, int32(52461700))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v87
	F_errmsg(m, int32(_a_F_fmgr_info_cxt_security_5), v12+int32(32))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_fmgr_info_cxt_security_3), int32(239), int32(_a_F_fmgr_info_cxt_security_4))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L9
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v368
	F_errmsg_internal(m, int32(_a_F_fmgr_info_cxt_security_6), v12+int32(48))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L9
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_fmgr_info_cxt_security_3), int32(410), int32(_a_F_fmgr_info_cxt_security_7))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L9
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v268
	F_errmsg_internal(m, int32(_a_F_fmgr_info_cxt_security_8), v12+int32(16))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_fmgr_info_cxt_security_3), int32(430), int32(_a_F_fmgr_info_cxt_security_9))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fmgr_sql_validator(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v427 int32
	_ = v427
	var v443 int32
	_ = v443
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = base.I32_wrap_i64(v20)
	v22 = F_CheckFunctionValidatorAccess(m, v19, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L4
	} else {
		goto L112
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L4
	} else {
		goto L107
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L104
	}
L4:
	;
	return int64(0)
L5:
	;
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = F_SearchSysCache1(m, int32(47), v20&int64(4294967295))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	m.G0 = v16 + int32(80)
	return int64(0)
L9:
	;
	if v29 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v35 = v33 + v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+108))
	v37 = F_get_typtype(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+104)))
	if int32(0) < v63 {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	if v37 != int32(112) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+108))
	if v41 <= int32(3830) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	switch v41 - int32(2249) {
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
	if base.B2i32(base.Ui32(v41-int32(_a_F_fmgr_sql_validator_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v41-int32(_a_F_fmgr_sql_validator_1)) < base.Ui32(int32(2))) != 0 {
		goto L11
	} else {
		goto L19
	}
L17:
	;
	if base.B2i32(v41 == int32(2776))|base.B2i32(v41 == int32(3500)) != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	goto L1
L19:
	;
	if v41 != int32(3831) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L11
L21:
	;
	v69 = int32(0)
	v76 = v2
	goto L24
L22:
	;
	v125 = v2
	goto L23
L23:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_fmgr_sql_validator[0])))
	if v132 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L24:
	;
	v84 = v35 + int32(136) + v69<<(uint(int32(2))%32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = F_get_typtype(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	v125 = v113
	goto L23
L26:
	;
	v115 = v69 + int32(1)
	v116 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+104)))
	if v115 < v116 {
		v69 = v115
		v76 = v113
		goto L24
	} else {
		goto L36
	}
L27:
	;
	if v86 != int32(112) {
		v113 = v76
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v90 = int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v91 <= int32(3830) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	switch v91 - int32(2277) {
	case 0, 6:
		v113 = v90
		goto L26
	case 1, 2, 3, 4, 5:
		goto L2
	default:
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if base.B2i32(base.Ui32(v91-int32(_a_F_fmgr_sql_validator_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v91-int32(_a_F_fmgr_sql_validator_1)) < base.Ui32(int32(2))) != 0 {
		v113 = v90
		goto L26
	} else {
		goto L34
	}
L32:
	;
	if base.B2i32(v91 == int32(2776))|base.B2i32(v91 == int32(3500)) != 0 {
		v113 = v90
		goto L26
	} else {
		goto L33
	}
L33:
	;
	goto L2
L34:
	;
	if v91 != int32(3831) {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v113 = v90
	goto L26
L36:
	;
	goto L25
L37:
	;
	v137 = F_SysCacheGetAttrNotNull(m, int32(47), v29, int32(26))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_ReleaseCatCache(m, v29)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L103
	}
L40:
	;
	v140 = F_text_to_cstring(m, base.I32_wrap_i64(v137))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v35 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = int32(508)
	v148 = int32(_a_F_fmgr_sql_validator_2)
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql_validator[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql_validator[1])) = v16 + int32(56)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v16 + int32(68)
	v162 = F_SysCacheGetAttr(m, int32(47), v29, int32(28), v16+int32(79))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+79)))
	if v164 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v16)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql_validator[1])) = v427
	goto L39
L44:
	;
	v290 = int32(0)
	if v277 == v290 {
		goto L77
	} else {
		goto L78
	}
L45:
	;
	if v125 != 0 {
		goto L43
	} else {
		goto L76
	}
L46:
	;
	v168 = F_text_to_cstring(m, base.I32_wrap_i64(v162))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v223 = F_pg_parse_query(m, v140)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L66
	}
L49:
	;
	if v184 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v170 = F_stringToNode(m, v168)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	if v172 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v184 = v176
	goto L49
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v170
	v182 = F_list_make1_impl(m, int32(1), v16+int32(40))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v184 = v182
	goto L49
L56:
	;
	v264 = int32(0)
	goto L45
L57:
	;
	goto L58
L58:
	;
	v188 = int32(0)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v189 <= v188 {
		v264 = v188
		goto L45
	} else {
		goto L59
	}
L59:
	;
	v193 = v188
	v196 = int32(0)
	goto L60
L60:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206+v196<<(uint(int32(2))%32))))
	F_AcquireRewriteLocks(m, v210, int32(1), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L62
	}
L61:
	;
	v264 = v217
	goto L45
L62:
	;
	v215 = F_pg_rewrite_query(m, v210)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v217 = F_lappend(m, v193, v215)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v220 = v196 + int32(1)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v220 < v221 {
		v193 = v217
		v196 = v220
		goto L60
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	if v125 != 0 {
		goto L43
	} else {
		goto L67
	}
L67:
	;
	v225 = int32(0)
	v228 = F_prepare_sql_fn_parse_info(m, v29, v225, v225)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	if v223 == int32(0) {
		v277 = v225
		goto L44
	} else {
		goto L69
	}
L69:
	;
	v232 = int32(0)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v233 <= v232 {
		v264 = v225
		goto L45
	} else {
		goto L70
	}
L70:
	;
	v236 = v225
	v239 = v232
	goto L71
L71:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v249+v239<<(uint(int32(2))%32))))
	v256 = F_pg_analyze_and_rewrite_withcb(m, v253, v140, int32(509), v228, int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L73
	}
L72:
	;
	v264 = v258
	goto L45
L73:
	;
	v258 = F_lappend(m, v236, v256)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	v261 = v239 + int32(1)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v261 < v262 {
		v236 = v258
		v239 = v261
		goto L71
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	v277 = v264
	goto L44
L77:
	;
	v399 = int32(0)
	v405 = F_internal_get_result_type(m, v21, v399, v399, v16+int32(48), v16+int32(44))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L4
	} else {
		goto L101
	}
L78:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	if v293 <= int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v296 = int32(0)
	if v296 < v293 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v299 = v293
	goto L82
L81:
	;
	v299 = v296
	goto L82
L82:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v277)+12))
	v304 = v290
	goto L84
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L4
	} else {
		goto L97
	}
L84:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v300+v304<<(uint(int32(2))%32))))
	if v317 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L77
L86:
	;
	v368 = v304 + int32(1)
	if v368 != v299 {
		v304 = v368
		goto L84
	} else {
		goto L96
	}
L87:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
	if v320 <= int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v317)+12))
	v327 = int32(0)
	goto L89
L89:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v323+v327<<(uint(int32(2))%32))))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	if v342 != int32(6) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L86
L91:
	;
	v352 = v327 + int32(1)
	if v320 != v352 {
		v327 = v352
		goto L89
	} else {
		goto L95
	}
L92:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341)+28))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	if v346 != int32(213) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	if v349 != 0 {
		goto L83
	} else {
		goto L94
	}
L94:
	;
	goto L91
L95:
	;
	goto L90
L96:
	;
	goto L85
L97:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_fmgr_sql_validator_3), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_fmgr_sql_validator_4), int32(2076), int32(_a_F_fmgr_sql_validator_5))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v409 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35)+96)))
	v411 = F_check_sql_fn_retval(m, v277, v407, v408, v409, int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	goto L43
L103:
	;
	goto L8
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v21
	F_errmsg_internal(m, int32(_a_F_fmgr_sql_validator_6), v16)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_fmgr_sql_validator_7), int32(884), int32(_a_F_fmgr_sql_validator_8))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
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
	F_errcode(m, int32(50724996))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v483 = F_format_type_be(m, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v483
	F_errmsg(m, int32(_a_F_fmgr_sql_validator_9), v16+int32(32))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_fmgr_sql_validator_7), int32(911), int32(_a_F_fmgr_sql_validator_8))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v35)+108))
	v504 = F_format_type_be(m, v503)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v504
	F_errmsg(m, int32(_a_F_fmgr_sql_validator_10), v16+int32(16))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_fmgr_sql_validator_7), int32(896), int32(_a_F_fmgr_sql_validator_8))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
