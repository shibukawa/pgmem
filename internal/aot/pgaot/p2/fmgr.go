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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
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
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
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
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v306 int32
	_ = v306
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
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
	v378 = m.ExcPending
	if v378 != 0 {
		goto L9
	} else {
		goto L86
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L9
	} else {
		goto L83
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L79
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
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
	v42 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
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
	if v42 == int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+22)))
	v48 = v46 + v47
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+104)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v49)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+99)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+10)) = uint8(v51)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+11)) = uint8(v53)
	if l3 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	F_ReleaseCatCache(m, v42)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L9
	} else {
		goto L75
	}
L13:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v48)+76))
	switch v77 - int32(12) {
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
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+97)))
	if v55 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v71 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1611)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
	goto L12
L16:
	;
	v58 = F_heap_attisnull(m, v42, int32(29), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	if v58 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[4]))
	if v63 == int32(0) {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v66 = m.T0[v63].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if v66 == int32(0) {
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
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)) = uint8(v287)
	goto L12
L23:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+22)))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261+v262)+76))
	v265 = F_SearchSysCache1(m, int32(36), v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L9
	} else {
		goto L71
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1612)
	v287 = int32(1)
	goto L22
L25:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+22)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v139+v140)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v142
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[5]))
	if v145 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L26:
	;
	v83 = F_SysCacheGetAttrNotNull(m, int32(47), v42, int32(26))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v85 = F_text_to_cstring(m, v83)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[6]))
	if v88 <= int32(0) {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v93 = int32(0)
	goto L30
L30:
	;
	v101 = v93 << (uint(int32(4)) % 32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_fmgr_info_cxt_security[7])))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if base.B2i32(v107 == int32(0))|base.B2i32(v107 != v110) != 0 {
		v128 = v107
		v129 = v110
		goto L33
	} else {
		goto L34
	}
L31:
	;
	F_pfree(m, v85)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L43
	}
L32:
	;
	if v128-v129 != 0 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	goto L32
L34:
	;
	v113 = v85
	v114 = v104
	goto L35
L35:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	if v118 == int32(0) {
		v128 = v118
		v129 = v117
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v128 = v118
	v129 = v117
	goto L33
L37:
	;
	v121 = int32(1)
	if v118 == v117 {
		v113 = v113 + v121
		v114 = v114 + v121
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v132 = v93 + int32(1)
	if v88 != v132 {
		v93 = v132
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
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_fmgr_info_cxt_security[3])))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v136
	v287 = int32(2)
	goto L22
L44:
	;
	v252 = int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	if v253 != v252 {
		goto L2
	} else {
		goto L70
	}
L45:
	;
	v188 = F_SysCacheGetAttrNotNull(m, int32(47), v42, int32(26))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L9
	} else {
		goto L57
	}
L46:
	;
	v150 = int32(0)
	v152 = F_hash_search(m, v145, v12+int32(60), v150, v150)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	if v152 == int32(0) {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v156 != v158 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v161 = v152 + int32(8)
	v163 = v42 + int32(4)
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161)+2)))
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161))))
	v166 = int32(16)
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+2)))
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163))))
	if v164|v165<<(uint(v166)%32) == v169|v170<<(uint(v166)%32) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v180 == int32(0) {
		goto L45
	} else {
		goto L56
	}
L51:
	;
	goto L50
L52:
	;
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161)+4)))
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+4)))
	if v176 == v177 {
		v180 = int32(1)
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v180 = int32(0)
	goto L51
L55:
	;
	goto L54
L56:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
	v248 = v183
	v249 = v184
	goto L44
L57:
	;
	v190 = F_text_to_cstring(m, v188)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v194 = F_SysCacheGetAttrNotNull(m, int32(47), v42, int32(27))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	v196 = F_text_to_cstring(m, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	v201 = F_load_external_function(m, v196, v190, int32(1), v12+int32(56))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v204 = F_fetch_finfo_record(m, v203, v190)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+22)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206+v207)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+108)) = v209
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[5]))
	if v212 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+76)) = int64(103079215108)
	v223 = F_hash_create(m, int32(_a_F_fmgr_info_cxt_security_1), int32(100), v12+int32(60), int32(40))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L9
	} else {
		goto L66
	}
L64:
	;
	v226 = v212
	goto L65
L65:
	;
	v232 = F_hash_search(m, v226, v12+int32(108), int32(1), v12+int32(60))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L9
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[5])) = v223
	v226 = v223
	goto L65
L67:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+4)) = v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+8)) = v237
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v232)+12)) = uint16(v239)
	*(*int32)(unsafe.Add(mBase, uint32(v232)+20)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v232)+16)) = v201
	F_pfree(m, v190)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v196)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v248 = v204
	v249 = v201
	goto L44
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v249
	v287 = v252
	goto L22
L71:
	;
	if v265 == int32(0) {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v265)+16))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+22)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v270+v271)+76))
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_info_cxt_security[8]))
	F_fmgr_info_cxt_security(m, v273, v12+int32(60), v277, int32(1))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v281
	F_ReleaseCatCache(m, v265)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L9
	} else {
		goto L74
	}
L74:
	;
	v287 = int32(0)
	goto L22
L75:
	;
	goto L5
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg_internal(m, int32(_a_F_fmgr_info_cxt_security_2), v12)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_fmgr_info_cxt_security_3), int32(183), int32(_a_F_fmgr_info_cxt_security_4))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
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
	v347 = m.ExcPending
	if v347 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v85
	F_errmsg(m, int32(_a_F_fmgr_info_cxt_security_5), v12+int32(32))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_fmgr_info_cxt_security_3), int32(237), int32(_a_F_fmgr_info_cxt_security_4))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
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
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v363
	F_errmsg_internal(m, int32(_a_F_fmgr_info_cxt_security_6), v12+int32(48))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L9
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_fmgr_info_cxt_security_3), int32(408), int32(_a_F_fmgr_info_cxt_security_7))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v264
	F_errmsg_internal(m, int32(_a_F_fmgr_info_cxt_security_8), v12+int32(16))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_fmgr_info_cxt_security_3), int32(428), int32(_a_F_fmgr_info_cxt_security_9))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
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
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
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
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v410 int32
	_ = v410
	var v425 int32
	_ = v425
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
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
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L112
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L107
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L4
	} else {
		goto L104
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
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31)+104)))
	if int32(0) < v59 {
		goto L21
	} else {
		goto L22
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
	if base.B2i32(base.Ui32(v37-int32(_a_F_fmgr_sql_validator_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v37-int32(_a_F_fmgr_sql_validator_1)) < base.Ui32(int32(2))) != 0 {
		goto L11
	} else {
		goto L19
	}
L17:
	;
	if base.B2i32(v37 == int32(2776))|base.B2i32(v37 == int32(3500)) != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	goto L1
L19:
	;
	if v37 != int32(3831) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L11
L21:
	;
	v65 = int32(0)
	v72 = v2
	goto L24
L22:
	;
	v120 = v2
	goto L23
L23:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_fmgr_sql_validator[0])))
	if v126 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L24:
	;
	v79 = v31 + int32(136) + v65<<(uint(int32(2))%32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v81 = F_get_typtype(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	v120 = v108
	goto L23
L26:
	;
	v110 = v65 + int32(1)
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31)+104)))
	if v110 < v111 {
		v65 = v110
		v72 = v108
		goto L24
	} else {
		goto L36
	}
L27:
	;
	if v81 != int32(112) {
		v108 = v72
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v85 = int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v86 <= int32(3830) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	switch v86 - int32(2277) {
	case 0, 6:
		v108 = v85
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
	if base.B2i32(base.Ui32(v86-int32(_a_F_fmgr_sql_validator_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v86-int32(_a_F_fmgr_sql_validator_1)) < base.Ui32(int32(2))) != 0 {
		v108 = v85
		goto L26
	} else {
		goto L34
	}
L32:
	;
	if base.B2i32(v86 == int32(2776))|base.B2i32(v86 == int32(3500)) != 0 {
		v108 = v85
		goto L26
	} else {
		goto L33
	}
L33:
	;
	goto L2
L34:
	;
	if v86 != int32(3831) {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v108 = v85
	goto L26
L36:
	;
	goto L25
L37:
	;
	v131 = F_SysCacheGetAttrNotNull(m, int32(47), v25, int32(26))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L103
	}
L40:
	;
	v133 = F_text_to_cstring(m, v131)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v31 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = int32(473)
	v141 = int32(_a_F_fmgr_sql_validator_2)
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql_validator[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql_validator[1])) = v15 + int32(56)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v15 + int32(68)
	v155 = F_SysCacheGetAttr(m, int32(47), v25, int32(28), v15+int32(79))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+79)))
	if v157 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql_validator[1])) = v410
	goto L39
L44:
	;
	v278 = int32(0)
	if v266 == v278 {
		goto L77
	} else {
		goto L78
	}
L45:
	;
	if v120 != 0 {
		goto L43
	} else {
		goto L76
	}
L46:
	;
	v160 = F_text_to_cstring(m, v155)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v214 = F_pg_parse_query(m, v133)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L66
	}
L49:
	;
	if v176 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v162 = F_stringToNode(m, v160)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v164 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v176 = v168
	goto L49
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v162
	v174 = F_list_make1_impl(m, int32(1), v15+int32(40))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v176 = v174
	goto L49
L56:
	;
	v254 = int32(0)
	goto L45
L57:
	;
	goto L58
L58:
	;
	v180 = int32(0)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v181 <= v180 {
		v254 = v180
		goto L45
	} else {
		goto L59
	}
L59:
	;
	v185 = v180
	v188 = int32(0)
	goto L60
L60:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197+v188<<(uint(int32(2))%32))))
	F_AcquireRewriteLocks(m, v201, int32(1), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L62
	}
L61:
	;
	v254 = v208
	goto L45
L62:
	;
	v206 = F_pg_rewrite_query(m, v201)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v208 = F_lappend(m, v185, v206)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v211 = v188 + int32(1)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v211 < v212 {
		v185 = v208
		v188 = v211
		goto L60
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	if v120 != 0 {
		goto L43
	} else {
		goto L67
	}
L67:
	;
	v216 = int32(0)
	v219 = F_prepare_sql_fn_parse_info(m, v25, v216, v216)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	if v214 == int32(0) {
		v266 = v216
		goto L44
	} else {
		goto L69
	}
L69:
	;
	v223 = int32(0)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v224 <= v223 {
		v254 = v216
		goto L45
	} else {
		goto L70
	}
L70:
	;
	v227 = v216
	v230 = v223
	goto L71
L71:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239+v230<<(uint(int32(2))%32))))
	v246 = F_pg_analyze_and_rewrite_withcb(m, v243, v133, int32(474), v219, int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L73
	}
L72:
	;
	v254 = v248
	goto L45
L73:
	;
	v248 = F_lappend(m, v227, v246)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	v251 = v230 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v251 < v252 {
		v227 = v248
		v230 = v251
		goto L71
	} else {
		goto L75
	}
L75:
	;
	goto L72
L76:
	;
	v266 = v254
	goto L44
L77:
	;
	v383 = int32(0)
	v389 = F_internal_get_result_type(m, v19, v383, v383, v15+int32(48), v15+int32(44))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L4
	} else {
		goto L101
	}
L78:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v281 <= int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v284 = int32(0)
	if v284 < v281 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v287 = v281
	goto L82
L81:
	;
	v287 = v284
	goto L82
L82:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	v292 = v278
	goto L84
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L97
	}
L84:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v288+v292<<(uint(int32(2))%32))))
	if v304 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L77
L86:
	;
	v353 = v292 + int32(1)
	if v353 != v287 {
		v292 = v353
		goto L84
	} else {
		goto L96
	}
L87:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	if v307 <= int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v304)+12))
	v314 = int32(0)
	goto L89
L89:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v310+v314<<(uint(int32(2))%32))))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v328 != int32(6) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L86
L91:
	;
	v338 = v314 + int32(1)
	if v307 != v338 {
		v314 = v338
		goto L89
	} else {
		goto L95
	}
L92:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v327)+28))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	if v332 != int32(213) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	if v335 != 0 {
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
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(_a_F_fmgr_sql_validator_3), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_fmgr_sql_validator_4), int32(2075), int32(_a_F_fmgr_sql_validator_5))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
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
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v393 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31)+96)))
	v395 = F_check_sql_fn_retval(m, v266, v391, v392, v393, int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
	F_errmsg_internal(m, int32(_a_F_fmgr_sql_validator_6), v15)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_fmgr_sql_validator_7), int32(853), int32(_a_F_fmgr_sql_validator_8))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
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
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v464 = F_format_type_be(m, v463)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v464
	F_errmsg(m, int32(_a_F_fmgr_sql_validator_9), v15+int32(32))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_fmgr_sql_validator_7), int32(880), int32(_a_F_fmgr_sql_validator_8))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
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
	v483 = m.ExcPending
	if v483 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	v485 = F_format_type_be(m, v484)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v485
	F_errmsg(m, int32(_a_F_fmgr_sql_validator_10), v15+int32(16))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_fmgr_sql_validator_7), int32(865), int32(_a_F_fmgr_sql_validator_8))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
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
