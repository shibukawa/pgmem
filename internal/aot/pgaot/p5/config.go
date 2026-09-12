package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FreeConfigVariables(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v4 = l0
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v8 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	F_pfree(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if v11 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	F_pfree(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v14 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	F_pfree(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v17 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	F_pfree(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_pfree(m, v4)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	if v7 != 0 {
		v4 = v7
		goto L4
	} else {
		goto L24
	}
L24:
	;
	goto L5
}
func F_ParseConfigFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
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
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = int32(738225)
	v20 = m.G0
	v22 = v20 - int32(32)
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v23
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1283])))
	if v31 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(96)
	return v345
L2:
	;
	v100 = F_strlen(m, l0)
	mBase = m.M
	if v99 == v100 {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	v99 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1284])))
	if v35 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v39 = l0
	goto L9
L7:
	;
	goto L8
L8:
	;
	v49 = v16
	v50 = v31
	goto L12
L9:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v45 == v31 {
		v39 = v39 + int32(1)
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v99 = v39 - l0
	goto L2
L11:
	;
	goto L10
L12:
	;
	v57 = v22 + int32(base.Ui32(v50)>>(uint(int32(3))%32))&int32(28)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v58 | v59<<(uint(v50)%32)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v63 != 0 {
		v49 = v49 + v59
		v50 = v63
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v66 == int32(0) {
		v91 = l0
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v99 = v91 - l0
	goto L2
L16:
	;
	v70 = l0
	v71 = v66
	goto L17
L17:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(base.Ui32(v71)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v79)>>(uint(v71)%32))&int32(1) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v91 = v87
	goto L15
L19:
	;
	v91 = v70
	goto L15
L20:
	;
	goto L21
L21:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	v87 = v70 + int32(1)
	if v85 != 0 {
		v70 = v87
		v71 = v85
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v102 = int32(0)
	v104 = F_errstart(m, l5, v102)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if int32(11) <= l4 {
		goto L44
	} else {
		goto L45
	}
L26:
	;
	return int32(0)
L27:
	;
	if v104 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v121 = F_palloc(m, int32(28))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L26
	} else {
		goto L34
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg(m, int32(708406), v14)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(312750), int32(194), int32(387604))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v121))) = int64(0)
	v126 = F_pstrdup(m, int32(379476))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = v126
	if l2 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v129 = F_pstrdup(m, l2)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L26
	} else {
		goto L39
	}
L37:
	;
	v131 = v102
	goto L38
L38:
	;
	v132 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+24)) = v132
	v134 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+20)) = uint16(v134)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v121)+12)) = v131
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v138 == v132 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v131 = v129
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v121
	v345 = int32(0)
	goto L1
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v121
	goto L40
L42:
	;
	goto L43
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(v142)+24)) = v121
	goto L40
L44:
	;
	v148 = int32(0)
	v150 = F_errstart(m, l5, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L26
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v192 = F_AbsoluteConfigLocation(m, l0, l2)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L26
	} else {
		goto L64
	}
L47:
	;
	if v150 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L26
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v167 = F_palloc(m, int32(28))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L26
	} else {
		goto L54
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l0
	F_errmsg(m, int32(459489), v14+int32(16))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L26
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(312750), int32(211), int32(387604))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L26
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v167))) = int64(0)
	v172 = F_pstrdup(m, int32(459594))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L26
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = v172
	if l2 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v175 = F_pstrdup(m, l2)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L26
	} else {
		goto L59
	}
L57:
	;
	v177 = v148
	goto L58
L58:
	;
	v178 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+24)) = v178
	v180 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v167)+20)) = uint16(v180)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v167)+12)) = v177
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v184 == v178 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v177 = v175
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v167
	v345 = int32(0)
	goto L1
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v167
	goto L60
L62:
	;
	goto L63
L63:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+24)) = v167
	goto L60
L64:
	;
	if l2 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v266 = F_AllocateFile(m, v192, int32(229660))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L26
	} else {
		goto L92
	}
L66:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v199 == int32(0) {
		v218 = v198
		v219 = v199
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v219-v218 != 0 {
		goto L65
	} else {
		goto L75
	}
L68:
	;
	goto L67
L69:
	;
	if v198 != v199 {
		v218 = v198
		v219 = v199
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v203 = v192
	v204 = l2
	goto L71
L71:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+1)))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)))
	if v208 == int32(0) {
		v218 = v207
		v219 = v208
		goto L68
	} else {
		goto L73
	}
L72:
	;
	v218 = v207
	v219 = v208
	goto L68
L73:
	;
	v211 = int32(1)
	if v207 == v208 {
		v203 = v203 + v211
		v204 = v204 + v211
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v222 = F_errstart(m, l5, int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L26
	} else {
		goto L76
	}
L76:
	;
	if v222 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L26
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v239 = F_palloc(m, int32(28))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L26
	} else {
		goto L83
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = l2
	F_errmsg(m, int32(693474), v14+int32(80))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L26
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(312750), int32(231), int32(387604))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L26
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v239))) = int64(0)
	v244 = F_pstrdup(m, int32(269391))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L26
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+8)) = v244
	v247 = F_pstrdup(m, l2)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L26
	} else {
		goto L85
	}
L85:
	;
	v249 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+24)) = v249
	v251 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v239)+20)) = uint16(v251)
	*(*int32)(unsafe.Add(mBase, uint32(v239)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v239)+12)) = v247
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v255 == v249 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v239
	F_pfree(m, v192)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L26
	} else {
		goto L90
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v239
	goto L86
L88:
	;
	goto L89
L89:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(v259)+24)) = v239
	goto L86
L90:
	;
	v345 = int32(0)
	goto L1
L91:
	;
	F_pfree(m, v192)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L26
	} else {
		goto L123
	}
L92:
	;
	if v266 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if l1 != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	v336 = F_ParseConfigFp(m, v266, v192, l4, l5, l6, l7)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L26
	} else {
		goto L121
	}
L96:
	;
	v270 = int32(0)
	v272 = F_errstart(m, l5, v270)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L26
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v318 = int32(1)
	v321 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L26
	} else {
		goto L117
	}
L99:
	;
	if v272 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L26
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v192
	v291 = F_psprintf(m, int32(698099), v14+int32(32))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L26
	} else {
		goto L106
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v192
	F_errmsg(m, int32(296645), v14+int32(48))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L26
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(312750), int32(247), int32(387604))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L26
	} else {
		goto L105
	}
L105:
	;
	goto L102
L106:
	;
	v294 = F_palloc(m, int32(28))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L26
	} else {
		goto L107
	}
L107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v294))) = int64(0)
	v298 = F_pstrdup(m, v291)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L26
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+8)) = v298
	if l2 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v301 = F_pstrdup(m, l2)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L26
	} else {
		goto L112
	}
L110:
	;
	v303 = v270
	goto L111
L111:
	;
	v304 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v294)+24)) = v304
	v306 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v294)+20)) = uint16(v306)
	*(*int32)(unsafe.Add(mBase, uint32(v294)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v294)+12)) = v303
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v310 == v304 {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	v303 = v301
	goto L111
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v294
	v340 = int32(0)
	goto L91
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v294
	goto L113
L115:
	;
	goto L116
L116:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(v314)+24)) = v294
	goto L113
L117:
	;
	if v321 == int32(0) {
		v340 = v318
		goto L91
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v192
	F_errmsg(m, int32(697507), v14-int32(-64))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L26
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(312750), int32(258), int32(387604))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L26
	} else {
		goto L120
	}
L120:
	;
	v340 = v318
	goto L91
L121:
	;
	v338 = F_FreeFile(m, v266)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L26
	} else {
		goto L122
	}
L122:
	;
	v340 = v336
	goto L91
L123:
	;
	v345 = v340
	goto L1
}
func F_SetConfigOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	if base.B2i32(l3 != int32(9))&base.B2i32(base.Ui32(l3) <= base.Ui32(int32(10))) != 0 {
		v14 = int32(10)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[239]))
		v14 = v13
	}
	v15 = int32(0)
	v19 = F_set_config_with_handle(m, l0, int32(0), l1, l2, l3, v14, v15, int32(1), v15, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		return
	}
}
func F_set_config_by_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v2 = int32(0)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v5 != int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v9 = F_text_to_cstring(m, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v13 == int32(0) {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v17 = F_text_to_cstring(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = v17
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
					if v20 == int32(0) {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v26 = base.B2i32(v23 != int32(0))
					} else {
						v26 = v2
					}
					v29 = F_superuser(m)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						if v29 != 0 {
							v31 = int32(5)
						} else {
							v31 = int32(6)
						}
						F_set_config_option(m, v9, v19, v31, int32(13), v26, int32(1))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = int32(0)
							v38 = F_GetConfigOptionByName(m, v9, v36, v36)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = F_cstring_to_text(m, v38)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									return v40
								}
							}
						}
					}
				}
			} else {
				v19 = v2
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
				if v20 == int32(0) {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v26 = base.B2i32(v23 != int32(0))
				} else {
					v26 = v2
				}
				v29 = F_superuser(m)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v29 != 0 {
						v31 = int32(5)
					} else {
						v31 = int32(6)
					}
					F_set_config_option(m, v9, v19, v31, int32(13), v26, int32(1))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = int32(0)
						v38 = F_GetConfigOptionByName(m, v9, v36, v36)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = F_cstring_to_text(m, v38)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								return v40
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(67108994))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(379033), int32(0))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491782), int32(342), int32(376019))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
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
func F_set_config_with_handle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
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
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
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
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v594 int32
	_ = v594
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v686 int32
	_ = v686
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v926 int32
	_ = v926
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v968 int32
	_ = v968
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v1018 int32
	_ = v1018
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1083 float64
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 float64
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 float64
	_ = v1144
	var v1145 float64
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1165 int32
	_ = v1165
	var v1170 int32
	_ = v1170
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 float64
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1184 float64
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 float64
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1235 int32
	_ = v1235
	var v1237 float64
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1247 int32
	_ = v1247
	var v1260 int32
	_ = v1260
	var v1269 int32
	_ = v1269
	var v1271 float64
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1302 int32
	_ = v1302
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1352 int32
	_ = v1352
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1605 int32
	_ = v1605
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1631 int32
	_ = v1631
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1683 int32
	_ = v1683
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
	var v1702 int32
	_ = v1702
	var v1711 int32
	_ = v1711
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1738 int32
	_ = v1738
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1797 int32
	_ = v1797
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1854 int32
	_ = v1854
	var v1872 int32
	_ = v1872
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1919 int32
	_ = v1919
	var v1923 int32
	_ = v1923
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1952 int32
	_ = v1952
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v1992 int32
	_ = v1992
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v1997 int32
	_ = v1997
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2017 int32
	_ = v2017
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2032 int32
	_ = v2032
	var v2037 int32
	_ = v2037
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2057 int32
	_ = v2057
	var v2059 int32
	_ = v2059
	var v2062 int32
	_ = v2062
	var v2068 int32
	_ = v2068
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2077 int32
	_ = v2077
	var v2079 int32
	_ = v2079
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2100 int32
	_ = v2100
	var v2102 int32
	_ = v2102
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2112 int32
	_ = v2112
	var v2125 int32
	_ = v2125
	var v2134 int32
	_ = v2134
	var v2136 int32
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2167 int32
	_ = v2167
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2179 int32
	_ = v2179
	var v2217 int32
	_ = v2217
	var v2235 int32
	_ = v2235
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2245 int32
	_ = v2245
	var v2247 int32
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2263 int32
	_ = v2263
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2297 int32
	_ = v2297
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2321 int32
	_ = v2321
	v11 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(272)
	m.G0 = v20
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v11
	if l8 != 0 {
		v36 = l8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	switch l4 {
	case 0, 3:
		goto L4
	default:
		goto L3
	}
L3:
	;
	if base.Ui32(l4-int32(5)) < base.Ui32(int32(4)) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _consts[184])))
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = int32(12)
	goto L7
L6:
	;
	v28 = int32(15)
	goto L7
L7:
	;
	v36 = v28
	goto L1
L8:
	;
	v35 = int32(19)
	goto L10
L9:
	;
	v35 = int32(21)
	goto L10
L10:
	;
	v36 = v35
	goto L1
L11:
	;
	m.G0 = v20 + int32(272)
	return v2321
L12:
	;
	v41 = F_find_option(m, l0, int32(1), int32(0), v36)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v47 = l1
	goto L14
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+72))
	if v53 != 0 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	return int32(0)
L16:
	;
	if v41 == int32(0) {
		v2321 = v11
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v47 = v41
	goto L14
L18:
	;
	if l7 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v55 = int32(1)
	goto L21
L20:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+76)))
	v55 = v54
	goto L21
L21:
	;
	goto L18
L22:
	;
	v86 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	switch v87 {
	case 0:
		goto L41
	case 1:
		goto L40
	case 2:
		goto L39
	case 3:
		goto L38
	case 4:
		goto L37
	case 5:
		goto L34
	default:
		v264 = v86
		goto L32
	}
L23:
	;
	if l6 == int32(2) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v55&int32(1) == int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+21)))
	if v64&int32(128) != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v67 = int32(0)
	v69 = F_errstart(m, v36, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	if v69 == int32(0) {
		v2321 = v67
		goto L11
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v76
	F_errmsg(m, int32(258568), v20)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(496556), int32(3465), int32(387834))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	v2321 = v67
	goto L11
L32:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	if v265&int32(4096) != 0 {
		goto L93
	} else {
		goto L94
	}
L33:
	;
	v264 = int32(0)
	goto L32
L34:
	;
	if l3&int32(-3) != int32(4) {
		goto L33
	} else {
		goto L82
	}
L35:
	;
	if l4 == int32(9) {
		goto L33
	} else {
		goto L72
	}
L36:
	;
	if l9 != 0 {
		v264 = v86
		goto L32
	} else {
		goto L69
	}
L37:
	;
	if l3 != int32(2) {
		goto L35
	} else {
		goto L68
	}
L38:
	;
	switch l3 - int32(2) {
	case 0:
		goto L36
	default:
		goto L35
	case 2:
		goto L60
	}
L39:
	;
	if base.Ui32(int32(-3)) < base.Ui32(l3-int32(3)) {
		v264 = v86
		goto L32
	} else {
		goto L54
	}
L40:
	;
	v111 = int32(1)
	switch l3 - v111 {
	case 0:
		goto L33
	case 1:
		v264 = v111
		goto L32
	default:
		goto L48
	}
L41:
	;
	if l3 == int32(0) {
		v264 = v86
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v90 = int32(0)
	v92 = F_errstart(m, v36, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L15
	} else {
		goto L43
	}
L43:
	;
	if v92 == int32(0) {
		v2321 = v90
		goto L11
	} else {
		goto L44
	}
L44:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v99
	F_errmsg(m, int32(457003), v20+int32(176))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(496556), int32(3481), int32(387834))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L47
	}
L47:
	;
	v2321 = v90
	goto L11
L48:
	;
	v114 = int32(0)
	v116 = F_errstart(m, v36, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	if v116 == int32(0) {
		v2321 = v114
		goto L11
	} else {
		goto L50
	}
L50:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L15
	} else {
		goto L51
	}
L51:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v123
	F_errmsg(m, int32(213763), v20+int32(192))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L15
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(496556), int32(3504), int32(387834))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L15
	} else {
		goto L53
	}
L53:
	;
	v2321 = v114
	goto L11
L54:
	;
	v139 = int32(0)
	v141 = F_errstart(m, v36, v139)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L15
	} else {
		goto L55
	}
L55:
	;
	if v141 == int32(0) {
		v2321 = v139
		goto L11
	} else {
		goto L56
	}
L56:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L15
	} else {
		goto L57
	}
L57:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v148
	F_errmsg(m, int32(30997), v20+int32(208))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L15
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(496556), int32(3514), int32(387834))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L15
	} else {
		goto L59
	}
L59:
	;
	v2321 = v139
	goto L11
L60:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v164 = F_pg_parameter_aclcheck(m, v162, l5, int64(4096))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L15
	} else {
		goto L61
	}
L61:
	;
	if v164 == int32(0) {
		goto L35
	} else {
		goto L62
	}
L62:
	;
	v168 = int32(0)
	v170 = F_errstart(m, v36, v168)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L15
	} else {
		goto L63
	}
L63:
	;
	if v170 == int32(0) {
		v2321 = v168
		goto L11
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v177
	F_errmsg(m, int32(682205), v20+int32(240))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L15
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(496556), int32(3541), int32(387834))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L15
	} else {
		goto L67
	}
L67:
	;
	v2321 = v168
	goto L11
L68:
	;
	goto L36
L69:
	;
	if l7 == int32(0) {
		v264 = v86
		goto L32
	} else {
		goto L70
	}
L70:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, _consts[184])))
	if v194&int32(1) == int32(0) {
		v264 = v86
		goto L32
	} else {
		goto L71
	}
L71:
	;
	v2321 = int32(-1)
	goto L11
L72:
	;
	if int32(1)<<(uint(l3)%32)&int32(26) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v209 = base.B2i32(base.Ui32(l3) <= base.Ui32(int32(4)))
	goto L75
L74:
	;
	v209 = int32(0)
	goto L75
L75:
	;
	if v209 != 0 {
		v264 = v86
		goto L32
	} else {
		goto L76
	}
L76:
	;
	v210 = int32(0)
	v212 = F_errstart(m, v36, v210)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L15
	} else {
		goto L77
	}
L77:
	;
	if v212 == int32(0) {
		v2321 = v210
		goto L11
	} else {
		goto L78
	}
L78:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L15
	} else {
		goto L79
	}
L79:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = v219
	F_errmsg(m, int32(82424), v20+int32(224))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L15
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(496556), int32(3583), int32(387834))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L15
	} else {
		goto L81
	}
L81:
	;
	v2321 = v210
	goto L11
L82:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v237 = F_pg_parameter_aclcheck(m, v235, l5, int64(4096))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L15
	} else {
		goto L83
	}
L83:
	;
	if v237 == int32(0) {
		v264 = v86
		goto L32
	} else {
		goto L84
	}
L84:
	;
	v241 = int32(0)
	v243 = F_errstart(m, v36, v241)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L15
	} else {
		goto L85
	}
L85:
	;
	if v243 == int32(0) {
		v2321 = v241
		goto L11
	} else {
		goto L86
	}
L86:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L15
	} else {
		goto L87
	}
L87:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v250
	F_errmsg(m, int32(682205), v20+int32(256))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L15
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(496556), int32(3603), int32(387834))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L15
	} else {
		goto L89
	}
L89:
	;
	v2321 = v241
	goto L11
L90:
	;
	if base.Ui32(int32(10)) < base.Ui32(l4) {
		v379 = v11
		goto L126
	} else {
		goto L127
	}
L91:
	;
	if l2 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L92:
	;
	v303 = int32(0)
	v305 = F_errstart(m, v36, v303)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L15
	} else {
		goto L107
	}
L93:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _consts[240]))
	if v269&int32(1) != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v300 = v265
	goto L95
L95:
	;
	if v300&int32(8) != 0 {
		goto L91
	} else {
		goto L106
	}
L96:
	;
	v272 = int32(0)
	v274 = F_errstart(m, v36, v272)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L15
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, _consts[240])))
	goto L104
L99:
	;
	if v274 == int32(0) {
		v2321 = v272
		goto L11
	} else {
		goto L100
	}
L100:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L15
	} else {
		goto L101
	}
L101:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v281
	F_errmsg(m, int32(250857), v20+int32(144))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L15
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(496556), int32(3642), int32(387834))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L15
	} else {
		goto L103
	}
L103:
	;
	v2321 = v272
	goto L11
L104:
	;
	if int32(base.Ui32(v294&int32(2))>>(uint(int32(1))%32)) != 0 {
		goto L92
	} else {
		goto L105
	}
L105:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v300 = v299
	goto L95
L106:
	;
	goto L90
L107:
	;
	if v305 == int32(0) {
		v2321 = v303
		goto L11
	} else {
		goto L108
	}
L108:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L15
	} else {
		goto L109
	}
L109:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v312
	F_errmsg(m, int32(259581), v20+int32(160))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L15
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(496556), int32(3650), int32(387834))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L15
	} else {
		goto L111
	}
L111:
	;
	v2321 = v303
	goto L11
L112:
	;
	v326 = int32(0)
	v328 = F_errstart(m, v36, v326)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L15
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	if l6 != int32(2) {
		goto L90
	} else {
		goto L120
	}
L115:
	;
	if v328 == int32(0) {
		v2321 = v326
		goto L11
	} else {
		goto L116
	}
L116:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L15
	} else {
		goto L117
	}
L117:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v335
	F_errmsg(m, int32(105323), v20+int32(112))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L15
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(496556), int32(3662), int32(387834))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L15
	} else {
		goto L119
	}
L119:
	;
	v2321 = v326
	goto L11
L120:
	;
	v349 = int32(0)
	v351 = F_errstart(m, v36, v349)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L15
	} else {
		goto L121
	}
L121:
	;
	if v351 == int32(0) {
		v2321 = v349
		goto L11
	} else {
		goto L122
	}
L122:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L15
	} else {
		goto L123
	}
L123:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v358
	F_errmsg(m, int32(140187), v20+int32(128))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L15
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(496556), int32(3670), int32(387834))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L15
	} else {
		goto L125
	}
L125:
	;
	v2321 = v349
	goto L11
L126:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if base.Ui32(v380) <= base.Ui32(l4) {
		v405 = l7
		goto L129
	} else {
		goto L130
	}
L127:
	;
	if l7 == int32(0) {
		v379 = v11
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v374 = int32(0)
	v379 = base.B2i32(l4 == v374) | base.B2i32(l2 != v374)
	goto L126
L129:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v407 {
	case 0:
		goto L142
	case 1:
		goto L141
	case 2:
		goto L140
	case 3:
		goto L139
	case 4:
		goto L138
	default:
		goto L137
	}
L130:
	;
	if l7^int32(1)|v379 != 0 {
		v405 = int32(0)
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v386 = int32(-1)
	v389 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L15
	} else {
		goto L132
	}
L132:
	;
	if v389 == int32(0) {
		v2321 = v386
		goto L11
	} else {
		goto L133
	}
L133:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v393
	F_errmsg_internal(m, int32(11074), v20+int32(96))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L15
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(496556), int32(3696), int32(387834))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L15
	} else {
		goto L135
	}
L135:
	;
	v2321 = v386
	goto L11
L136:
	;
	v2291 = int32(1)
	v2292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+20)))
	if v2292&int32(64) == int32(0) {
		v2321 = v2291
		goto L11
	} else {
		goto L858
	}
L137:
	;
	if v405 != 0 {
		goto L136
	} else {
		goto L857
	}
L138:
	;
	if l2 != 0 {
		goto L731
	} else {
		goto L732
	}
L139:
	;
	if l2 != 0 {
		goto L525
	} else {
		goto L526
	}
L140:
	;
	if l2 != 0 {
		goto L398
	} else {
		goto L399
	}
L141:
	;
	if l2 != 0 {
		goto L271
	} else {
		goto L272
	}
L142:
	;
	if l2 != 0 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	if v264 != 0 {
		goto L154
	} else {
		goto L155
	}
L144:
	;
	v412 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L15
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	if l4 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	if v412 != 0 {
		v435 = l3
		v436 = l4
		v437 = l5
		goto L143
	} else {
		goto L148
	}
L148:
	;
	v2321 = int32(0)
	goto L11
L149:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)) = uint8(v417)
	v419 = int32(0)
	v426 = F_call_bool_check_hook(m, v47, v20+int32(264), v20+int32(260), v419, v36)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L15
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)) = uint8(v428)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v430
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v435 = v433
	v436 = v434
	v437 = v432
	goto L143
L152:
	;
	if v426 != 0 {
		v435 = l3
		v436 = v419
		v437 = l5
		goto L143
	} else {
		goto L153
	}
L153:
	;
	v2321 = v419
	goto L11
L154:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v439 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L156
L156:
	;
	if v405 != 0 {
		goto L188
	} else {
		goto L189
	}
L157:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	if v478 != v479 {
		goto L180
	} else {
		goto L181
	}
L158:
	;
	v443 = int32(1)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v439 == v444 {
		v473 = v443
		goto L160
	} else {
		goto L161
	}
L159:
	;
	if v473 != 0 {
		goto L157
	} else {
		goto L178
	}
L160:
	;
	goto L159
L161:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v446 {
	case 0:
		goto L167
	case 1:
		goto L166
	case 2:
		goto L165
	case 3:
		goto L164
	case 4:
		goto L163
	default:
		goto L162
	}
L162:
	;
	v459 = v47 + int32(56)
	goto L173
L163:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v439 == v455 {
		v473 = v443
		goto L160
	} else {
		goto L172
	}
L164:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v439 != v453 {
		goto L162
	} else {
		goto L171
	}
L165:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v439 != v451 {
		goto L162
	} else {
		goto L170
	}
L166:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v439 != v449 {
		goto L162
	} else {
		goto L169
	}
L167:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v439 != v447 {
		goto L162
	} else {
		goto L168
	}
L168:
	;
	v473 = v443
	goto L160
L169:
	;
	v473 = v443
	goto L160
L170:
	;
	v473 = v443
	goto L160
L171:
	;
	v473 = v443
	goto L160
L172:
	;
	goto L162
L173:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v459)))
	v463 = int32(0)
	v464 = base.B2i32(v462 != v463)
	if v462 == v463 {
		v473 = v464
		goto L160
	} else {
		goto L175
	}
L174:
	;
	v473 = v464
	goto L160
L175:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v462)+40))
	if v439 == v467 {
		v473 = v464
		goto L160
	} else {
		goto L176
	}
L176:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v462)+56))
	if v439 != v469 {
		v459 = v462
		goto L173
	} else {
		goto L177
	}
L177:
	;
	goto L174
L178:
	;
	F_pfree(m, v439)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L15
	} else {
		goto L179
	}
L179:
	;
	goto L157
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v476 | int32(2)
	v484 = int32(0)
	v486 = F_errstart(m, v36, v484)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L15
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v476 & int32(-3)
	v2321 = int32(-1)
	goto L11
L183:
	;
	if v486 == int32(0) {
		v2321 = v484
		goto L11
	} else {
		goto L184
	}
L184:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L15
	} else {
		goto L185
	}
L185:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v493
	F_errmsg(m, int32(213763), v20+int32(16))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L15
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(496556), int32(3748), int32(387834))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L15
	} else {
		goto L187
	}
L187:
	;
	v2321 = v484
	goto L11
L188:
	;
	if v379 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	goto L190
L190:
	;
	if v379 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L191:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L15
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v47)+104))
	if v514 != 0 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	goto L193
L195:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	m.T0[v514].(func(*base.Module, int32, int32))(m, v513&int32(1), v517)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L15
	} else {
		goto L198
	}
L196:
	;
	v521 = v513
	goto L197
L197:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	*(*uint8)(unsafe.Add(mBase, uint32(v522))) = uint8(v521)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(60), v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L15
	} else {
		goto L199
	}
L198:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	v521 = v520
	goto L197
L199:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v531 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v435
	goto L190
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v436
	goto L200
L202:
	;
	if v436 == int32(0) {
		goto L201
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	if v436 != 0 {
		goto L201
	} else {
		goto L210
	}
L205:
	;
	v537 = v47 - int32(-64)
	v539 = *(*int32)(unsafe.Add(mBase, _consts[1289]))
	if v539 != 0 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+64)) = v546
	v548 = int32(4484956)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v548
	*(*int32)(unsafe.Add(mBase, uint32(v546)+4)) = v537
	*(*int32)(unsafe.Add(mBase, _consts[1290])) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v436
	goto L200
L207:
	;
	v541 = *(*int32)(unsafe.Add(mBase, _consts[1290]))
	v546 = v541
	goto L206
L208:
	;
	goto L209
L209:
	;
	v543 = int32(4484956)
	*(*int32)(unsafe.Add(mBase, _consts[1289])) = v543
	v546 = v543
	goto L206
L210:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v554)+4)) = v555
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = v557
	goto L201
L211:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v704 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L212:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v569) <= base.Ui32(v436) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+112)) = uint8(v571)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(116), v575)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L15
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v581 == int32(0) {
		goto L211
	} else {
		goto L217
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v436
	goto L215
L217:
	;
	v594 = v581
	goto L218
L218:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v594)+12))
	if base.Ui32(v603) <= base.Ui32(v436) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	goto L211
L220:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	*(*uint8)(unsafe.Add(mBase, uint32(v594)+32)) = uint8(v605)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v594)+40))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v594)+40)) = v608
	if v607 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	goto L222
L222:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v594)))
	if v686 != 0 {
		v594 = v686
		goto L218
	} else {
		goto L245
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v594)+24)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v594)+16)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v594)+12)) = v436
	goto L222
L224:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v607 == v612 {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v614 {
	case 0:
		goto L231
	case 1:
		goto L230
	case 2:
		goto L229
	case 3:
		goto L228
	case 4:
		goto L227
	default:
		goto L226
	}
L226:
	;
	v636 = v47 + int32(56)
	goto L237
L227:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v607 == v623 {
		goto L223
	} else {
		goto L236
	}
L228:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v607 != v621 {
		goto L226
	} else {
		goto L235
	}
L229:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v607 != v619 {
		goto L226
	} else {
		goto L234
	}
L230:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v607 != v617 {
		goto L226
	} else {
		goto L233
	}
L231:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v607 != v615 {
		goto L226
	} else {
		goto L232
	}
L232:
	;
	goto L223
L233:
	;
	goto L223
L234:
	;
	goto L223
L235:
	;
	goto L223
L236:
	;
	goto L226
L237:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v636)))
	if v642 != 0 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	F_pfree(m, v607)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L15
	} else {
		goto L244
	}
L239:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+40))
	if v607 == v643 {
		goto L223
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	goto L238
L242:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v642)+56))
	if v607 != v645 {
		v636 = v642
		goto L237
	} else {
		goto L243
	}
L243:
	;
	goto L223
L244:
	;
	goto L223
L245:
	;
	goto L219
L246:
	;
	if v405 != 0 {
		goto L136
	} else {
		goto L269
	}
L247:
	;
	v708 = int32(1)
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v704 == v709 {
		v738 = v708
		goto L249
	} else {
		goto L250
	}
L248:
	;
	if v738 != 0 {
		goto L246
	} else {
		goto L267
	}
L249:
	;
	goto L248
L250:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v711 {
	case 0:
		goto L256
	case 1:
		goto L255
	case 2:
		goto L254
	case 3:
		goto L253
	case 4:
		goto L252
	default:
		goto L251
	}
L251:
	;
	v724 = v47 + int32(56)
	goto L262
L252:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v704 == v720 {
		v738 = v708
		goto L249
	} else {
		goto L261
	}
L253:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v704 != v718 {
		goto L251
	} else {
		goto L260
	}
L254:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v704 != v716 {
		goto L251
	} else {
		goto L259
	}
L255:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v704 != v714 {
		goto L251
	} else {
		goto L258
	}
L256:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v704 != v712 {
		goto L251
	} else {
		goto L257
	}
L257:
	;
	v738 = v708
	goto L249
L258:
	;
	v738 = v708
	goto L249
L259:
	;
	v738 = v708
	goto L249
L260:
	;
	v738 = v708
	goto L249
L261:
	;
	goto L251
L262:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v724)))
	v728 = int32(0)
	v729 = base.B2i32(v727 != v728)
	if v727 == v728 {
		v738 = v729
		goto L249
	} else {
		goto L264
	}
L263:
	;
	v738 = v729
	goto L249
L264:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v727)+40))
	if v704 == v732 {
		v738 = v729
		goto L249
	} else {
		goto L265
	}
L265:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v727)+56))
	if v704 != v734 {
		v724 = v727
		goto L262
	} else {
		goto L266
	}
L266:
	;
	goto L263
L267:
	;
	F_pfree(m, v704)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L15
	} else {
		goto L268
	}
L268:
	;
	goto L246
L269:
	;
	v2321 = int32(-1)
	goto L11
L270:
	;
	if v264 != 0 {
		goto L281
	} else {
		goto L282
	}
L271:
	;
	v746 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L15
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	if l4 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	if v746 != 0 {
		v769 = l3
		v770 = l4
		v771 = l5
		goto L270
	} else {
		goto L275
	}
L275:
	;
	v2321 = int32(0)
	goto L11
L276:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v751
	v753 = int32(0)
	v760 = F_call_int_check_hook(m, v47, v20+int32(264), v20+int32(260), v753, v36)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L15
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v762
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v764
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v769 = v767
	v770 = v768
	v771 = v766
	goto L270
L279:
	;
	if v760 != 0 {
		v769 = l3
		v770 = v753
		v771 = l5
		goto L270
	} else {
		goto L280
	}
L280:
	;
	v2321 = v753
	goto L11
L281:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v773 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L283
L283:
	;
	if v405 != 0 {
		goto L315
	} else {
		goto L316
	}
L284:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v811)))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if v812 != v813 {
		goto L307
	} else {
		goto L308
	}
L285:
	;
	v777 = int32(1)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v773 == v778 {
		v807 = v777
		goto L287
	} else {
		goto L288
	}
L286:
	;
	if v807 != 0 {
		goto L284
	} else {
		goto L305
	}
L287:
	;
	goto L286
L288:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v780 {
	case 0:
		goto L294
	case 1:
		goto L293
	case 2:
		goto L292
	case 3:
		goto L291
	case 4:
		goto L290
	default:
		goto L289
	}
L289:
	;
	v793 = v47 + int32(56)
	goto L300
L290:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v773 == v789 {
		v807 = v777
		goto L287
	} else {
		goto L299
	}
L291:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v773 != v787 {
		goto L289
	} else {
		goto L298
	}
L292:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v773 != v785 {
		goto L289
	} else {
		goto L297
	}
L293:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v773 != v783 {
		goto L289
	} else {
		goto L296
	}
L294:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v773 != v781 {
		goto L289
	} else {
		goto L295
	}
L295:
	;
	v807 = v777
	goto L287
L296:
	;
	v807 = v777
	goto L287
L297:
	;
	v807 = v777
	goto L287
L298:
	;
	v807 = v777
	goto L287
L299:
	;
	goto L289
L300:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v793)))
	v797 = int32(0)
	v798 = base.B2i32(v796 != v797)
	if v796 == v797 {
		v807 = v798
		goto L287
	} else {
		goto L302
	}
L301:
	;
	v807 = v798
	goto L287
L302:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v796)+40))
	if v773 == v801 {
		v807 = v798
		goto L287
	} else {
		goto L303
	}
L303:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v796)+56))
	if v773 != v803 {
		v793 = v796
		goto L300
	} else {
		goto L304
	}
L304:
	;
	goto L301
L305:
	;
	F_pfree(m, v773)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L15
	} else {
		goto L306
	}
L306:
	;
	goto L284
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v810 | int32(2)
	v818 = int32(0)
	v820 = F_errstart(m, v36, v818)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L15
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v810 & int32(-3)
	v2321 = int32(-1)
	goto L11
L310:
	;
	if v820 == int32(0) {
		v2321 = v818
		goto L11
	} else {
		goto L311
	}
L311:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L15
	} else {
		goto L312
	}
L312:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v827
	F_errmsg(m, int32(213763), v20+int32(32))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L15
	} else {
		goto L313
	}
L313:
	;
	F_errfinish(m, int32(496556), int32(3846), int32(387834))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L15
	} else {
		goto L314
	}
L314:
	;
	v2321 = v818
	goto L11
L315:
	;
	if v379 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L316:
	;
	goto L317
L317:
	;
	if v379 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L318:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L15
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	if v848 != 0 {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	goto L320
L322:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	m.T0[v848].(func(*base.Module, int32, int32))(m, v847, v849)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L15
	} else {
		goto L325
	}
L323:
	;
	v853 = v847
	goto L324
L324:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v854))) = v853
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(60), v858)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L15
	} else {
		goto L326
	}
L325:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v853 = v852
	goto L324
L326:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v863 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v769
	goto L317
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v770
	goto L327
L329:
	;
	if v770 == int32(0) {
		goto L328
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	if v770 != 0 {
		goto L328
	} else {
		goto L337
	}
L332:
	;
	v869 = v47 - int32(-64)
	v871 = *(*int32)(unsafe.Add(mBase, _consts[1289]))
	if v871 != 0 {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+64)) = v878
	v880 = int32(4484956)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v880
	*(*int32)(unsafe.Add(mBase, uint32(v878)+4)) = v869
	*(*int32)(unsafe.Add(mBase, _consts[1290])) = v869
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v770
	goto L327
L334:
	;
	v873 = *(*int32)(unsafe.Add(mBase, _consts[1290]))
	v878 = v873
	goto L333
L335:
	;
	goto L336
L336:
	;
	v875 = int32(4484956)
	*(*int32)(unsafe.Add(mBase, _consts[1289])) = v875
	v878 = v875
	goto L333
L337:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v886)+4)) = v887
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v887))) = v889
	goto L328
L338:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1036 == int32(0) {
		goto L373
	} else {
		goto L374
	}
L339:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v901) <= base.Ui32(v770) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+120)) = v903
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(124), v907)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L15
	} else {
		goto L343
	}
L341:
	;
	goto L342
L342:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v913 == int32(0) {
		goto L338
	} else {
		goto L344
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v769
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v770
	goto L342
L344:
	;
	v926 = v913
	goto L345
L345:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v926)+12))
	if base.Ui32(v935) <= base.Ui32(v770) {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	goto L338
L347:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v926)+32)) = v937
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v926)+40))
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v926)+40)) = v940
	if v939 == int32(0) {
		goto L350
	} else {
		goto L351
	}
L348:
	;
	goto L349
L349:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v926)))
	if v1018 != 0 {
		v926 = v1018
		goto L345
	} else {
		goto L372
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v926)+24)) = v771
	*(*int32)(unsafe.Add(mBase, uint32(v926)+16)) = v769
	*(*int32)(unsafe.Add(mBase, uint32(v926)+12)) = v770
	goto L349
L351:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v939 == v944 {
		goto L350
	} else {
		goto L352
	}
L352:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v946 {
	case 0:
		goto L358
	case 1:
		goto L357
	case 2:
		goto L356
	case 3:
		goto L355
	case 4:
		goto L354
	default:
		goto L353
	}
L353:
	;
	v968 = v47 + int32(56)
	goto L364
L354:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v939 == v955 {
		goto L350
	} else {
		goto L363
	}
L355:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v939 != v953 {
		goto L353
	} else {
		goto L362
	}
L356:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v939 != v951 {
		goto L353
	} else {
		goto L361
	}
L357:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v939 != v949 {
		goto L353
	} else {
		goto L360
	}
L358:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v939 != v947 {
		goto L353
	} else {
		goto L359
	}
L359:
	;
	goto L350
L360:
	;
	goto L350
L361:
	;
	goto L350
L362:
	;
	goto L350
L363:
	;
	goto L353
L364:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v968)))
	if v974 != 0 {
		goto L366
	} else {
		goto L367
	}
L365:
	;
	F_pfree(m, v939)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L15
	} else {
		goto L371
	}
L366:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v974)+40))
	if v939 == v975 {
		goto L350
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	goto L365
L369:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v974)+56))
	if v939 != v977 {
		v968 = v974
		goto L364
	} else {
		goto L370
	}
L370:
	;
	goto L350
L371:
	;
	goto L350
L372:
	;
	goto L346
L373:
	;
	if v405 != 0 {
		goto L136
	} else {
		goto L396
	}
L374:
	;
	v1040 = int32(1)
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1036 == v1041 {
		v1070 = v1040
		goto L376
	} else {
		goto L377
	}
L375:
	;
	if v1070 != 0 {
		goto L373
	} else {
		goto L394
	}
L376:
	;
	goto L375
L377:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1043 {
	case 0:
		goto L383
	case 1:
		goto L382
	case 2:
		goto L381
	case 3:
		goto L380
	case 4:
		goto L379
	default:
		goto L378
	}
L378:
	;
	v1056 = v47 + int32(56)
	goto L389
L379:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1036 == v1052 {
		v1070 = v1040
		goto L376
	} else {
		goto L388
	}
L380:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1036 != v1050 {
		goto L378
	} else {
		goto L387
	}
L381:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1036 != v1048 {
		goto L378
	} else {
		goto L386
	}
L382:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1036 != v1046 {
		goto L378
	} else {
		goto L385
	}
L383:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1036 != v1044 {
		goto L378
	} else {
		goto L384
	}
L384:
	;
	v1070 = v1040
	goto L376
L385:
	;
	v1070 = v1040
	goto L376
L386:
	;
	v1070 = v1040
	goto L376
L387:
	;
	v1070 = v1040
	goto L376
L388:
	;
	goto L378
L389:
	;
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1056)))
	v1060 = int32(0)
	v1061 = base.B2i32(v1059 != v1060)
	if v1059 == v1060 {
		v1070 = v1061
		goto L376
	} else {
		goto L391
	}
L390:
	;
	v1070 = v1061
	goto L376
L391:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+40))
	if v1036 == v1064 {
		v1070 = v1061
		goto L376
	} else {
		goto L392
	}
L392:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1059)+56))
	if v1036 != v1066 {
		v1056 = v1059
		goto L389
	} else {
		goto L393
	}
L393:
	;
	goto L390
L394:
	;
	F_pfree(m, v1036)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L15
	} else {
		goto L395
	}
L395:
	;
	goto L373
L396:
	;
	v2321 = int32(-1)
	goto L11
L397:
	;
	if v264 != 0 {
		goto L408
	} else {
		goto L409
	}
L398:
	;
	v1078 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L15
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	if l4 == int32(0) {
		goto L403
	} else {
		goto L404
	}
L401:
	;
	if v1078 != 0 {
		v1101 = l3
		v1102 = l4
		v1103 = l5
		goto L397
	} else {
		goto L402
	}
L402:
	;
	v2321 = int32(0)
	goto L11
L403:
	;
	v1083 = *(*float64)(unsafe.Add(mBase, uint32(v47)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+264)) = v1083
	v1085 = int32(0)
	v1092 = F_call_real_check_hook(m, v47, v20+int32(264), v20+int32(260), v1085, v36)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L15
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	v1094 = *(*float64)(unsafe.Add(mBase, uint32(v47)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+264)) = v1094
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v1096
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v1101 = v1099
	v1102 = v1100
	v1103 = v1098
	goto L397
L406:
	;
	if v1092 != 0 {
		v1101 = l3
		v1102 = v1085
		v1103 = l5
		goto L397
	} else {
		goto L407
	}
L407:
	;
	v2321 = v1085
	goto L11
L408:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1105 == int32(0) {
		goto L411
	} else {
		goto L412
	}
L409:
	;
	goto L410
L410:
	;
	if v405 != 0 {
		goto L442
	} else {
		goto L443
	}
L411:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v1144 = *(*float64)(unsafe.Add(mBase, uint32(v1143)))
	v1145 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	if base.F64_ne(v1144, v1145) != 0 {
		goto L434
	} else {
		goto L435
	}
L412:
	;
	v1109 = int32(1)
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1105 == v1110 {
		v1139 = v1109
		goto L414
	} else {
		goto L415
	}
L413:
	;
	if v1139 != 0 {
		goto L411
	} else {
		goto L432
	}
L414:
	;
	goto L413
L415:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1112 {
	case 0:
		goto L421
	case 1:
		goto L420
	case 2:
		goto L419
	case 3:
		goto L418
	case 4:
		goto L417
	default:
		goto L416
	}
L416:
	;
	v1125 = v47 + int32(56)
	goto L427
L417:
	;
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1105 == v1121 {
		v1139 = v1109
		goto L414
	} else {
		goto L426
	}
L418:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1105 != v1119 {
		goto L416
	} else {
		goto L425
	}
L419:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1105 != v1117 {
		goto L416
	} else {
		goto L424
	}
L420:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1105 != v1115 {
		goto L416
	} else {
		goto L423
	}
L421:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1105 != v1113 {
		goto L416
	} else {
		goto L422
	}
L422:
	;
	v1139 = v1109
	goto L414
L423:
	;
	v1139 = v1109
	goto L414
L424:
	;
	v1139 = v1109
	goto L414
L425:
	;
	v1139 = v1109
	goto L414
L426:
	;
	goto L416
L427:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1125)))
	v1129 = int32(0)
	v1130 = base.B2i32(v1128 != v1129)
	if v1128 == v1129 {
		v1139 = v1130
		goto L414
	} else {
		goto L429
	}
L428:
	;
	v1139 = v1130
	goto L414
L429:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+40))
	if v1105 == v1133 {
		v1139 = v1130
		goto L414
	} else {
		goto L430
	}
L430:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+56))
	if v1105 != v1135 {
		v1125 = v1128
		goto L427
	} else {
		goto L431
	}
L431:
	;
	goto L428
L432:
	;
	F_pfree(m, v1105)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L15
	} else {
		goto L433
	}
L433:
	;
	goto L411
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v1142 | int32(2)
	v1150 = int32(0)
	v1152 = F_errstart(m, v36, v1150)
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L15
	} else {
		goto L437
	}
L435:
	;
	goto L436
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v1142 & int32(-3)
	v2321 = int32(-1)
	goto L11
L437:
	;
	if v1152 == int32(0) {
		v2321 = v1150
		goto L11
	} else {
		goto L438
	}
L438:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L15
	} else {
		goto L439
	}
L439:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v1159
	F_errmsg(m, int32(213763), v20+int32(48))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L15
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(496556), int32(3944), int32(387834))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L15
	} else {
		goto L441
	}
L441:
	;
	v2321 = v1150
	goto L11
L442:
	;
	if v379 == int32(0) {
		goto L445
	} else {
		goto L446
	}
L443:
	;
	goto L444
L444:
	;
	if v379 == int32(0) {
		goto L465
	} else {
		goto L466
	}
L445:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L15
	} else {
		goto L448
	}
L446:
	;
	goto L447
L447:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	v1180 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1181 != 0 {
		goto L449
	} else {
		goto L450
	}
L448:
	;
	goto L447
L449:
	;
	m.T0[v1181].(func(*base.Module, float64, int32))(m, v1180, v1179)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L15
	} else {
		goto L452
	}
L450:
	;
	v1186 = v1179
	v1187 = v1180
	goto L451
L451:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	*(*float64)(unsafe.Add(mBase, uint32(v1188))) = v1187
	F_set_extra_field(m, v47, v47+int32(60), v1186)
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L15
	} else {
		goto L453
	}
L452:
	;
	v1184 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	v1186 = v1185
	v1187 = v1184
	goto L451
L453:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v1196 == int32(0) {
		goto L456
	} else {
		goto L457
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v1103
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v1101
	goto L444
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1102
	goto L454
L456:
	;
	if v1102 == int32(0) {
		goto L455
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	if v1102 != 0 {
		goto L455
	} else {
		goto L464
	}
L459:
	;
	v1202 = v47 - int32(-64)
	v1204 = *(*int32)(unsafe.Add(mBase, _consts[1289]))
	if v1204 != 0 {
		goto L461
	} else {
		goto L462
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+64)) = v1211
	v1213 = int32(4484956)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v1213
	*(*int32)(unsafe.Add(mBase, uint32(v1211)+4)) = v1202
	*(*int32)(unsafe.Add(mBase, _consts[1290])) = v1202
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1102
	goto L454
L461:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, _consts[1290]))
	v1211 = v1206
	goto L460
L462:
	;
	goto L463
L463:
	;
	v1208 = int32(4484956)
	*(*int32)(unsafe.Add(mBase, _consts[1289])) = v1208
	v1211 = v1208
	goto L460
L464:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1219)+4)) = v1220
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v1220))) = v1222
	goto L455
L465:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1370 == int32(0) {
		goto L500
	} else {
		goto L501
	}
L466:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v1235) <= base.Ui32(v1102) {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	v1237 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	*(*float64)(unsafe.Add(mBase, uint32(v47)+136)) = v1237
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(144), v1241)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L15
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v1247 == int32(0) {
		goto L465
	} else {
		goto L471
	}
L470:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v1103
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v1102
	goto L469
L471:
	;
	v1260 = v1247
	goto L472
L472:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1260)+12))
	if base.Ui32(v1269) <= base.Ui32(v1102) {
		goto L474
	} else {
		goto L475
	}
L473:
	;
	goto L465
L474:
	;
	v1271 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	*(*float64)(unsafe.Add(mBase, uint32(v1260)+32)) = v1271
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1260)+40))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v1260)+40)) = v1274
	if v1273 == int32(0) {
		goto L477
	} else {
		goto L478
	}
L475:
	;
	goto L476
L476:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1260)))
	if v1352 != 0 {
		v1260 = v1352
		goto L472
	} else {
		goto L499
	}
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1260)+24)) = v1103
	*(*int32)(unsafe.Add(mBase, uint32(v1260)+16)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(v1260)+12)) = v1102
	goto L476
L478:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1273 == v1278 {
		goto L477
	} else {
		goto L479
	}
L479:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1280 {
	case 0:
		goto L485
	case 1:
		goto L484
	case 2:
		goto L483
	case 3:
		goto L482
	case 4:
		goto L481
	default:
		goto L480
	}
L480:
	;
	v1302 = v47 + int32(56)
	goto L491
L481:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1273 == v1289 {
		goto L477
	} else {
		goto L490
	}
L482:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1273 != v1287 {
		goto L480
	} else {
		goto L489
	}
L483:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1273 != v1285 {
		goto L480
	} else {
		goto L488
	}
L484:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1273 != v1283 {
		goto L480
	} else {
		goto L487
	}
L485:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1273 != v1281 {
		goto L480
	} else {
		goto L486
	}
L486:
	;
	goto L477
L487:
	;
	goto L477
L488:
	;
	goto L477
L489:
	;
	goto L477
L490:
	;
	goto L480
L491:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1302)))
	if v1308 != 0 {
		goto L493
	} else {
		goto L494
	}
L492:
	;
	F_pfree(m, v1273)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L15
	} else {
		goto L498
	}
L493:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+40))
	if v1273 == v1309 {
		goto L477
	} else {
		goto L496
	}
L494:
	;
	goto L495
L495:
	;
	goto L492
L496:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+56))
	if v1273 != v1311 {
		v1302 = v1308
		goto L491
	} else {
		goto L497
	}
L497:
	;
	goto L477
L498:
	;
	goto L477
L499:
	;
	goto L473
L500:
	;
	if v405 != 0 {
		goto L136
	} else {
		goto L523
	}
L501:
	;
	v1374 = int32(1)
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1370 == v1375 {
		v1404 = v1374
		goto L503
	} else {
		goto L504
	}
L502:
	;
	if v1404 != 0 {
		goto L500
	} else {
		goto L521
	}
L503:
	;
	goto L502
L504:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1377 {
	case 0:
		goto L510
	case 1:
		goto L509
	case 2:
		goto L508
	case 3:
		goto L507
	case 4:
		goto L506
	default:
		goto L505
	}
L505:
	;
	v1390 = v47 + int32(56)
	goto L516
L506:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1370 == v1386 {
		v1404 = v1374
		goto L503
	} else {
		goto L515
	}
L507:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1370 != v1384 {
		goto L505
	} else {
		goto L514
	}
L508:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1370 != v1382 {
		goto L505
	} else {
		goto L513
	}
L509:
	;
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1370 != v1380 {
		goto L505
	} else {
		goto L512
	}
L510:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1370 != v1378 {
		goto L505
	} else {
		goto L511
	}
L511:
	;
	v1404 = v1374
	goto L503
L512:
	;
	v1404 = v1374
	goto L503
L513:
	;
	v1404 = v1374
	goto L503
L514:
	;
	v1404 = v1374
	goto L503
L515:
	;
	goto L505
L516:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1390)))
	v1394 = int32(0)
	v1395 = base.B2i32(v1393 != v1394)
	if v1393 == v1394 {
		v1404 = v1395
		goto L503
	} else {
		goto L518
	}
L517:
	;
	v1404 = v1395
	goto L503
L518:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+40))
	if v1370 == v1398 {
		v1404 = v1395
		goto L503
	} else {
		goto L519
	}
L519:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+56))
	if v1370 != v1400 {
		v1390 = v1393
		goto L516
	} else {
		goto L520
	}
L520:
	;
	goto L517
L521:
	;
	F_pfree(m, v1370)
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L15
	} else {
		goto L522
	}
L522:
	;
	goto L500
L523:
	;
	v2321 = int32(-1)
	goto L11
L524:
	;
	if v264 != 0 {
		goto L543
	} else {
		goto L544
	}
L525:
	;
	v1412 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L15
	} else {
		goto L528
	}
L526:
	;
	goto L527
L527:
	;
	if l4 == int32(0) {
		goto L530
	} else {
		goto L531
	}
L528:
	;
	if v1412 != 0 {
		v1446 = l4
		v1447 = l5
		v1448 = l3
		goto L524
	} else {
		goto L529
	}
L529:
	;
	v2321 = int32(0)
	goto L11
L530:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	if v1417 != 0 {
		goto L534
	} else {
		goto L535
	}
L531:
	;
	goto L532
L532:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v1438
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v1440
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v1446 = v1444
	v1447 = v1442
	v1448 = v1443
	goto L524
L533:
	;
	v1430 = F_call_string_check_hook(m, v47, v20+int32(264), v20+int32(260), int32(0), v36)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L15
	} else {
		goto L539
	}
L534:
	;
	v1418 = F_guc_strdup(m, v36, v1417)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L15
	} else {
		goto L537
	}
L535:
	;
	goto L536
L536:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = int32(0)
	goto L533
L537:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v1418
	if v1418 != 0 {
		goto L533
	} else {
		goto L538
	}
L538:
	;
	v2321 = int32(0)
	goto L11
L539:
	;
	if v1430 != 0 {
		v1446 = v11
		v1447 = l5
		v1448 = l3
		goto L524
	} else {
		goto L540
	}
L540:
	;
	v1432 = int32(0)
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if v1433 == v1432 {
		v2321 = v1432
		goto L11
	} else {
		goto L541
	}
L541:
	;
	F_pfree(m, v1433)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L15
	} else {
		goto L542
	}
L542:
	;
	v2321 = v1432
	goto L11
L543:
	;
	v1449 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1450)))
	if v1451 == int32(0) {
		goto L548
	} else {
		goto L549
	}
L544:
	;
	goto L545
L545:
	;
	if v405 == int32(0) {
		goto L603
	} else {
		goto L604
	}
L546:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1516 == int32(0) {
		goto L572
	} else {
		goto L573
	}
L547:
	;
	v1488 = int32(1)
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1489)))
	if v1449 == v1490 {
		v1512 = v1488
		goto L561
	} else {
		goto L562
	}
L548:
	;
	v1483 = int32(1)
	if v1449 == int32(0) {
		v1515 = v1483
		goto L546
	} else {
		goto L559
	}
L549:
	;
	if v1449 == int32(0) {
		goto L548
	} else {
		goto L550
	}
L550:
	;
	v1458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1449))))
	v1459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1451))))
	if v1459 == int32(0) {
		v1478 = v1458
		v1479 = v1459
		goto L552
	} else {
		goto L553
	}
L551:
	;
	v1486 = base.B2i32(v1479-v1478 != int32(0))
	goto L547
L552:
	;
	goto L551
L553:
	;
	if v1458 != v1459 {
		v1478 = v1458
		v1479 = v1459
		goto L552
	} else {
		goto L554
	}
L554:
	;
	v1463 = v1451
	v1464 = v1449
	goto L555
L555:
	;
	v1467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1464)+1)))
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1463)+1)))
	if v1468 == int32(0) {
		v1478 = v1467
		v1479 = v1468
		goto L552
	} else {
		goto L557
	}
L556:
	;
	v1478 = v1467
	v1479 = v1468
	goto L552
L557:
	;
	v1471 = int32(1)
	if v1467 == v1468 {
		v1463 = v1463 + v1471
		v1464 = v1464 + v1471
		goto L555
	} else {
		goto L558
	}
L558:
	;
	goto L556
L559:
	;
	v1486 = v1483
	goto L547
L560:
	;
	if v1512 != 0 {
		v1515 = v1486
		goto L546
	} else {
		goto L570
	}
L561:
	;
	goto L560
L562:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	if v1449 == v1492 {
		v1512 = v1488
		goto L561
	} else {
		goto L563
	}
L563:
	;
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	if v1449 == v1494 {
		v1512 = v1488
		goto L561
	} else {
		goto L564
	}
L564:
	;
	v1498 = v47 + int32(56)
	goto L565
L565:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1498)))
	v1502 = int32(0)
	v1503 = base.B2i32(v1501 != v1502)
	if v1501 == v1502 {
		v1512 = v1503
		goto L561
	} else {
		goto L567
	}
L566:
	;
	v1512 = v1503
	goto L561
L567:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1501)+32))
	if v1449 == v1506 {
		v1512 = v1503
		goto L561
	} else {
		goto L568
	}
L568:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v1501)+48))
	if v1449 != v1508 {
		v1498 = v1501
		goto L565
	} else {
		goto L569
	}
L569:
	;
	goto L566
L570:
	;
	F_pfree(m, v1449)
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L15
	} else {
		goto L571
	}
L571:
	;
	v1515 = v1486
	goto L546
L572:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	if v1515 != 0 {
		goto L595
	} else {
		goto L596
	}
L573:
	;
	v1520 = int32(1)
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1516 == v1521 {
		v1550 = v1520
		goto L575
	} else {
		goto L576
	}
L574:
	;
	if v1550 != 0 {
		goto L572
	} else {
		goto L593
	}
L575:
	;
	goto L574
L576:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1523 {
	case 0:
		goto L582
	case 1:
		goto L581
	case 2:
		goto L580
	case 3:
		goto L579
	case 4:
		goto L578
	default:
		goto L577
	}
L577:
	;
	v1536 = v47 + int32(56)
	goto L588
L578:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1516 == v1532 {
		v1550 = v1520
		goto L575
	} else {
		goto L587
	}
L579:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1516 != v1530 {
		goto L577
	} else {
		goto L586
	}
L580:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1516 != v1528 {
		goto L577
	} else {
		goto L585
	}
L581:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1516 != v1526 {
		goto L577
	} else {
		goto L584
	}
L582:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1516 != v1524 {
		goto L577
	} else {
		goto L583
	}
L583:
	;
	v1550 = v1520
	goto L575
L584:
	;
	v1550 = v1520
	goto L575
L585:
	;
	v1550 = v1520
	goto L575
L586:
	;
	v1550 = v1520
	goto L575
L587:
	;
	goto L577
L588:
	;
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1536)))
	v1540 = int32(0)
	v1541 = base.B2i32(v1539 != v1540)
	if v1539 == v1540 {
		v1550 = v1541
		goto L575
	} else {
		goto L590
	}
L589:
	;
	v1550 = v1541
	goto L575
L590:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1539)+40))
	if v1516 == v1544 {
		v1550 = v1541
		goto L575
	} else {
		goto L591
	}
L591:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1539)+56))
	if v1516 != v1546 {
		v1536 = v1539
		goto L588
	} else {
		goto L592
	}
L592:
	;
	goto L589
L593:
	;
	F_pfree(m, v1516)
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L15
	} else {
		goto L594
	}
L594:
	;
	goto L572
L595:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v1553 | int32(2)
	v1557 = int32(0)
	v1559 = F_errstart(m, v36, v1557)
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L15
	} else {
		goto L598
	}
L596:
	;
	goto L597
L597:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v1553 & int32(-3)
	v2321 = int32(-1)
	goto L11
L598:
	;
	if v1559 == int32(0) {
		v2321 = v1557
		goto L11
	} else {
		goto L599
	}
L599:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v1565 = m.ExcPending
	if v1565 != 0 {
		goto L15
	} else {
		goto L600
	}
L600:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v1566
	F_errmsg(m, int32(213763), v20-int32(-64))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L15
	} else {
		goto L601
	}
L601:
	;
	F_errfinish(m, int32(496556), int32(4071), int32(387834))
	mBase = m.M
	v1577 = m.ExcPending
	if v1577 != 0 {
		goto L15
	} else {
		goto L602
	}
L602:
	;
	v2321 = v1557
	goto L11
L603:
	;
	if v379 == int32(0) {
		goto L643
	} else {
		goto L644
	}
L604:
	;
	if v379 == int32(0) {
		goto L605
	} else {
		goto L606
	}
L605:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L15
	} else {
		goto L608
	}
L606:
	;
	goto L607
L607:
	;
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v47)+104))
	if v1589 != 0 {
		goto L609
	} else {
		goto L610
	}
L608:
	;
	goto L607
L609:
	;
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	m.T0[v1589].(func(*base.Module, int32, int32))(m, v1588, v1590)
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L15
	} else {
		goto L612
	}
L610:
	;
	v1594 = v1588
	goto L611
L611:
	;
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	F_set_string_field(m, v47, v1595, v1594)
	mBase = m.M
	v1597 = m.ExcPending
	if v1597 != 0 {
		goto L15
	} else {
		goto L613
	}
L612:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v1594 = v1593
	goto L611
L613:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(60), v1600)
	mBase = m.M
	v1602 = m.ExcPending
	if v1602 != 0 {
		goto L15
	} else {
		goto L614
	}
L614:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v1605 == int32(0) {
		goto L617
	} else {
		goto L618
	}
L615:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v1447
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v1448
	if l9 != 0 {
		goto L603
	} else {
		goto L626
	}
L616:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1446
	goto L615
L617:
	;
	if v1446 == int32(0) {
		goto L616
	} else {
		goto L620
	}
L618:
	;
	goto L619
L619:
	;
	if v1446 != 0 {
		goto L616
	} else {
		goto L625
	}
L620:
	;
	v1611 = v47 - int32(-64)
	v1613 = *(*int32)(unsafe.Add(mBase, _consts[1289]))
	if v1613 != 0 {
		goto L622
	} else {
		goto L623
	}
L621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+64)) = v1620
	v1622 = int32(4484956)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v1622
	*(*int32)(unsafe.Add(mBase, uint32(v1620)+4)) = v1611
	*(*int32)(unsafe.Add(mBase, _consts[1290])) = v1611
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1446
	goto L615
L622:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, _consts[1290]))
	v1620 = v1615
	goto L621
L623:
	;
	goto L624
L624:
	;
	v1617 = int32(4484956)
	*(*int32)(unsafe.Add(mBase, _consts[1289])) = v1617
	v1620 = v1617
	goto L621
L625:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1628)+4)) = v1629
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v1629))) = v1631
	goto L616
L626:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v1640 = int32(257191)
	v1643 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1291])))
	v1644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1639))))
	if v1644 == int32(0) {
		v1663 = v1643
		v1664 = v1644
		goto L628
	} else {
		goto L629
	}
L627:
	;
	if v1664-v1663 != 0 {
		goto L603
	} else {
		goto L635
	}
L628:
	;
	goto L627
L629:
	;
	if v1643 != v1644 {
		v1663 = v1643
		v1664 = v1644
		goto L628
	} else {
		goto L630
	}
L630:
	;
	v1648 = v1639
	v1649 = v1640
	goto L631
L631:
	;
	v1652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1649)+1)))
	v1653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1648)+1)))
	if v1653 == int32(0) {
		v1663 = v1652
		v1664 = v1653
		goto L628
	} else {
		goto L633
	}
L632:
	;
	v1663 = v1652
	v1664 = v1653
	goto L628
L633:
	;
	v1656 = int32(1)
	if v1652 == v1653 {
		v1648 = v1648 + v1656
		v1649 = v1649 + v1656
		goto L631
	} else {
		goto L634
	}
L634:
	;
	goto L632
L635:
	;
	v1667 = int32(0)
	if l2 != 0 {
		goto L636
	} else {
		goto L637
	}
L636:
	;
	v1670 = int32(370343)
	goto L638
L637:
	;
	v1670 = v1667
	goto L638
L638:
	;
	if l4 == int32(10) {
		goto L639
	} else {
		goto L640
	}
L639:
	;
	v1674 = int32(1)
	goto L641
L640:
	;
	v1674 = l4
	goto L641
L641:
	;
	v1677 = F_set_config_with_handle(m, int32(383285), v1667, v1670, l3, v1674, l5, l6, int32(1), v36, int32(0))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L15
	} else {
		goto L642
	}
L642:
	;
	goto L603
L643:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if v1872 == int32(0) {
		goto L692
	} else {
		goto L693
	}
L644:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v1683) <= base.Ui32(v1446) {
		goto L645
	} else {
		goto L646
	}
L645:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	F_set_string_field(m, v47, v47+int32(112), v1687)
	mBase = m.M
	v1689 = m.ExcPending
	if v1689 != 0 {
		goto L15
	} else {
		goto L648
	}
L646:
	;
	goto L647
L647:
	;
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v1698 == int32(0) {
		goto L643
	} else {
		goto L650
	}
L648:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(116), v1692)
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L15
	} else {
		goto L649
	}
L649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v1447
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v1448
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v1446
	goto L647
L650:
	;
	v1702 = v47 + int32(56)
	v1711 = v1698
	goto L651
L651:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1711)+12))
	if base.Ui32(v1720) <= base.Ui32(v1446) {
		goto L653
	} else {
		goto L654
	}
L652:
	;
	goto L643
L653:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1711)+32))
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v1711)+32)) = v1723
	if v1722 == int32(0) {
		goto L656
	} else {
		goto L657
	}
L654:
	;
	goto L655
L655:
	;
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1711)))
	if v1854 != 0 {
		v1711 = v1854
		goto L651
	} else {
		goto L691
	}
L656:
	;
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v1711)+40))
	v1776 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v1711)+40)) = v1776
	if v1775 == int32(0) {
		goto L669
	} else {
		goto L670
	}
L657:
	;
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1727)))
	if v1722 == v1728 {
		goto L656
	} else {
		goto L658
	}
L658:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	if v1722 == v1730 {
		goto L656
	} else {
		goto L659
	}
L659:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	if v1722 == v1732 {
		goto L656
	} else {
		goto L660
	}
L660:
	;
	v1738 = v1702
	goto L661
L661:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1738)))
	if v1751 != 0 {
		goto L663
	} else {
		goto L664
	}
L662:
	;
	F_pfree(m, v1722)
	mBase = m.M
	v1757 = m.ExcPending
	if v1757 != 0 {
		goto L15
	} else {
		goto L668
	}
L663:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+32))
	if v1722 == v1752 {
		goto L656
	} else {
		goto L666
	}
L664:
	;
	goto L665
L665:
	;
	goto L662
L666:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+48))
	if v1722 != v1754 {
		v1738 = v1751
		goto L661
	} else {
		goto L667
	}
L667:
	;
	goto L656
L668:
	;
	goto L656
L669:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1711)+24)) = v1447
	*(*int32)(unsafe.Add(mBase, uint32(v1711)+16)) = v1448
	*(*int32)(unsafe.Add(mBase, uint32(v1711)+12)) = v1446
	goto L655
L670:
	;
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1775 == v1780 {
		goto L669
	} else {
		goto L671
	}
L671:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1782 {
	case 0:
		goto L677
	case 1:
		goto L676
	case 2:
		goto L675
	case 3:
		goto L674
	case 4:
		goto L673
	default:
		goto L672
	}
L672:
	;
	v1797 = v1702
	goto L683
L673:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1775 == v1791 {
		goto L669
	} else {
		goto L682
	}
L674:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1775 != v1789 {
		goto L672
	} else {
		goto L681
	}
L675:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1775 != v1787 {
		goto L672
	} else {
		goto L680
	}
L676:
	;
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1775 != v1785 {
		goto L672
	} else {
		goto L679
	}
L677:
	;
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1775 != v1783 {
		goto L672
	} else {
		goto L678
	}
L678:
	;
	goto L669
L679:
	;
	goto L669
L680:
	;
	goto L669
L681:
	;
	goto L669
L682:
	;
	goto L672
L683:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1797)))
	if v1810 != 0 {
		goto L685
	} else {
		goto L686
	}
L684:
	;
	F_pfree(m, v1775)
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L15
	} else {
		goto L690
	}
L685:
	;
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v1810)+40))
	if v1775 == v1811 {
		goto L669
	} else {
		goto L688
	}
L686:
	;
	goto L687
L687:
	;
	goto L684
L688:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1810)+56))
	if v1775 != v1813 {
		v1797 = v1810
		goto L683
	} else {
		goto L689
	}
L689:
	;
	goto L669
L690:
	;
	goto L669
L691:
	;
	goto L652
L692:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1903 == int32(0) {
		goto L706
	} else {
		goto L707
	}
L693:
	;
	v1876 = int32(1)
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1877)))
	if v1872 == v1878 {
		v1900 = v1876
		goto L695
	} else {
		goto L696
	}
L694:
	;
	if v1900 != 0 {
		goto L692
	} else {
		goto L704
	}
L695:
	;
	goto L694
L696:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	if v1872 == v1880 {
		v1900 = v1876
		goto L695
	} else {
		goto L697
	}
L697:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	if v1872 == v1882 {
		v1900 = v1876
		goto L695
	} else {
		goto L698
	}
L698:
	;
	v1886 = v47 + int32(56)
	goto L699
L699:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1886)))
	v1890 = int32(0)
	v1891 = base.B2i32(v1889 != v1890)
	if v1889 == v1890 {
		v1900 = v1891
		goto L695
	} else {
		goto L701
	}
L700:
	;
	v1900 = v1891
	goto L695
L701:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1889)+32))
	if v1872 == v1894 {
		v1900 = v1891
		goto L695
	} else {
		goto L702
	}
L702:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1889)+48))
	if v1872 != v1896 {
		v1886 = v1889
		goto L699
	} else {
		goto L703
	}
L703:
	;
	goto L700
L704:
	;
	F_pfree(m, v1872)
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L15
	} else {
		goto L705
	}
L705:
	;
	goto L692
L706:
	;
	if v405 != 0 {
		goto L136
	} else {
		goto L729
	}
L707:
	;
	v1907 = int32(1)
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1903 == v1908 {
		v1937 = v1907
		goto L709
	} else {
		goto L710
	}
L708:
	;
	if v1937 != 0 {
		goto L706
	} else {
		goto L727
	}
L709:
	;
	goto L708
L710:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1910 {
	case 0:
		goto L716
	case 1:
		goto L715
	case 2:
		goto L714
	case 3:
		goto L713
	case 4:
		goto L712
	default:
		goto L711
	}
L711:
	;
	v1923 = v47 + int32(56)
	goto L722
L712:
	;
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1903 == v1919 {
		v1937 = v1907
		goto L709
	} else {
		goto L721
	}
L713:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1903 != v1917 {
		goto L711
	} else {
		goto L720
	}
L714:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1903 != v1915 {
		goto L711
	} else {
		goto L719
	}
L715:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1903 != v1913 {
		goto L711
	} else {
		goto L718
	}
L716:
	;
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1903 != v1911 {
		goto L711
	} else {
		goto L717
	}
L717:
	;
	v1937 = v1907
	goto L709
L718:
	;
	v1937 = v1907
	goto L709
L719:
	;
	v1937 = v1907
	goto L709
L720:
	;
	v1937 = v1907
	goto L709
L721:
	;
	goto L711
L722:
	;
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1923)))
	v1927 = int32(0)
	v1928 = base.B2i32(v1926 != v1927)
	if v1926 == v1927 {
		v1937 = v1928
		goto L709
	} else {
		goto L724
	}
L723:
	;
	v1937 = v1928
	goto L709
L724:
	;
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+40))
	if v1903 == v1931 {
		v1937 = v1928
		goto L709
	} else {
		goto L725
	}
L725:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+56))
	if v1903 != v1933 {
		v1923 = v1926
		goto L722
	} else {
		goto L726
	}
L726:
	;
	goto L723
L727:
	;
	F_pfree(m, v1903)
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L15
	} else {
		goto L728
	}
L728:
	;
	goto L706
L729:
	;
	v2321 = int32(-1)
	goto L11
L730:
	;
	if v264 != 0 {
		goto L741
	} else {
		goto L742
	}
L731:
	;
	v1945 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v1946 = m.ExcPending
	if v1946 != 0 {
		goto L15
	} else {
		goto L734
	}
L732:
	;
	goto L733
L733:
	;
	if l4 == int32(0) {
		goto L736
	} else {
		goto L737
	}
L734:
	;
	if v1945 != 0 {
		v1968 = l3
		v1969 = l4
		v1970 = l5
		goto L730
	} else {
		goto L735
	}
L735:
	;
	v2321 = int32(0)
	goto L11
L736:
	;
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v1950
	v1952 = int32(0)
	v1959 = F_call_enum_check_hook(m, v47, v20+int32(264), v20+int32(260), v1952, v36)
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L15
	} else {
		goto L739
	}
L737:
	;
	goto L738
L738:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v1961
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v1963
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v1968 = v1966
	v1969 = v1967
	v1970 = v1965
	goto L730
L739:
	;
	if v1959 != 0 {
		v1968 = l3
		v1969 = v1952
		v1970 = l5
		goto L730
	} else {
		goto L740
	}
L740:
	;
	v2321 = v1952
	goto L11
L741:
	;
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1972 == int32(0) {
		goto L744
	} else {
		goto L745
	}
L742:
	;
	goto L743
L743:
	;
	if v405 != 0 {
		goto L775
	} else {
		goto L776
	}
L744:
	;
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v2010)))
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if v2011 != v2012 {
		goto L767
	} else {
		goto L768
	}
L745:
	;
	v1976 = int32(1)
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1972 == v1977 {
		v2006 = v1976
		goto L747
	} else {
		goto L748
	}
L746:
	;
	if v2006 != 0 {
		goto L744
	} else {
		goto L765
	}
L747:
	;
	goto L746
L748:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1979 {
	case 0:
		goto L754
	case 1:
		goto L753
	case 2:
		goto L752
	case 3:
		goto L751
	case 4:
		goto L750
	default:
		goto L749
	}
L749:
	;
	v1992 = v47 + int32(56)
	goto L760
L750:
	;
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1972 == v1988 {
		v2006 = v1976
		goto L747
	} else {
		goto L759
	}
L751:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1972 != v1986 {
		goto L749
	} else {
		goto L758
	}
L752:
	;
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1972 != v1984 {
		goto L749
	} else {
		goto L757
	}
L753:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1972 != v1982 {
		goto L749
	} else {
		goto L756
	}
L754:
	;
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1972 != v1980 {
		goto L749
	} else {
		goto L755
	}
L755:
	;
	v2006 = v1976
	goto L747
L756:
	;
	v2006 = v1976
	goto L747
L757:
	;
	v2006 = v1976
	goto L747
L758:
	;
	v2006 = v1976
	goto L747
L759:
	;
	goto L749
L760:
	;
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v1992)))
	v1996 = int32(0)
	v1997 = base.B2i32(v1995 != v1996)
	if v1995 == v1996 {
		v2006 = v1997
		goto L747
	} else {
		goto L762
	}
L761:
	;
	v2006 = v1997
	goto L747
L762:
	;
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(v1995)+40))
	if v1972 == v2000 {
		v2006 = v1997
		goto L747
	} else {
		goto L763
	}
L763:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v1995)+56))
	if v1972 != v2002 {
		v1992 = v1995
		goto L760
	} else {
		goto L764
	}
L764:
	;
	goto L761
L765:
	;
	F_pfree(m, v1972)
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		goto L15
	} else {
		goto L766
	}
L766:
	;
	goto L744
L767:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v2009 | int32(2)
	v2017 = int32(0)
	v2019 = F_errstart(m, v36, v2017)
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L15
	} else {
		goto L770
	}
L768:
	;
	goto L769
L769:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v2009 & int32(-3)
	v2321 = int32(-1)
	goto L11
L770:
	;
	if v2019 == int32(0) {
		v2321 = v2017
		goto L11
	} else {
		goto L771
	}
L771:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L15
	} else {
		goto L772
	}
L772:
	;
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v2026
	F_errmsg(m, int32(213763), v20+int32(80))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L15
	} else {
		goto L773
	}
L773:
	;
	F_errfinish(m, int32(496556), int32(4212), int32(387834))
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L15
	} else {
		goto L774
	}
L774:
	;
	v2321 = v2017
	goto L11
L775:
	;
	if v379 == int32(0) {
		goto L778
	} else {
		goto L779
	}
L776:
	;
	goto L777
L777:
	;
	if v379 == int32(0) {
		goto L798
	} else {
		goto L799
	}
L778:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L15
	} else {
		goto L781
	}
L779:
	;
	goto L780
L780:
	;
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v47)+108))
	if v2047 != 0 {
		goto L782
	} else {
		goto L783
	}
L781:
	;
	goto L780
L782:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	m.T0[v2047].(func(*base.Module, int32, int32))(m, v2046, v2048)
	mBase = m.M
	v2050 = m.ExcPending
	if v2050 != 0 {
		goto L15
	} else {
		goto L785
	}
L783:
	;
	v2052 = v2046
	goto L784
L784:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2053))) = v2052
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(60), v2057)
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L15
	} else {
		goto L786
	}
L785:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v2052 = v2051
	goto L784
L786:
	;
	v2062 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v2062 == int32(0) {
		goto L789
	} else {
		goto L790
	}
L787:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v1970
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v1968
	goto L777
L788:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1969
	goto L787
L789:
	;
	if v1969 == int32(0) {
		goto L788
	} else {
		goto L792
	}
L790:
	;
	goto L791
L791:
	;
	if v1969 != 0 {
		goto L788
	} else {
		goto L797
	}
L792:
	;
	v2068 = v47 - int32(-64)
	v2070 = *(*int32)(unsafe.Add(mBase, _consts[1289]))
	if v2070 != 0 {
		goto L794
	} else {
		goto L795
	}
L793:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+64)) = v2077
	v2079 = int32(4484956)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v2079
	*(*int32)(unsafe.Add(mBase, uint32(v2077)+4)) = v2068
	*(*int32)(unsafe.Add(mBase, _consts[1290])) = v2068
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1969
	goto L787
L794:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, _consts[1290]))
	v2077 = v2072
	goto L793
L795:
	;
	goto L796
L796:
	;
	v2074 = int32(4484956)
	*(*int32)(unsafe.Add(mBase, _consts[1289])) = v2074
	v2077 = v2074
	goto L793
L797:
	;
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v2085)+4)) = v2086
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v2086))) = v2088
	goto L788
L798:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v2235 == int32(0) {
		goto L833
	} else {
		goto L834
	}
L799:
	;
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v2100) <= base.Ui32(v1969) {
		goto L800
	} else {
		goto L801
	}
L800:
	;
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+116)) = v2102
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(120), v2106)
	mBase = m.M
	v2108 = m.ExcPending
	if v2108 != 0 {
		goto L15
	} else {
		goto L803
	}
L801:
	;
	goto L802
L802:
	;
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v2112 == int32(0) {
		goto L798
	} else {
		goto L804
	}
L803:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v1970
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v1968
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v1969
	goto L802
L804:
	;
	v2125 = v2112
	goto L805
L805:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2125)+12))
	if base.Ui32(v2134) <= base.Ui32(v1969) {
		goto L807
	} else {
		goto L808
	}
L806:
	;
	goto L798
L807:
	;
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v2125)+32)) = v2136
	v2138 = *(*int32)(unsafe.Add(mBase, uint32(v2125)+40))
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v2125)+40)) = v2139
	if v2138 == int32(0) {
		goto L810
	} else {
		goto L811
	}
L808:
	;
	goto L809
L809:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v2125)))
	if v2217 != 0 {
		v2125 = v2217
		goto L805
	} else {
		goto L832
	}
L810:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2125)+24)) = v1970
	*(*int32)(unsafe.Add(mBase, uint32(v2125)+16)) = v1968
	*(*int32)(unsafe.Add(mBase, uint32(v2125)+12)) = v1969
	goto L809
L811:
	;
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v2138 == v2143 {
		goto L810
	} else {
		goto L812
	}
L812:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v2145 {
	case 0:
		goto L818
	case 1:
		goto L817
	case 2:
		goto L816
	case 3:
		goto L815
	case 4:
		goto L814
	default:
		goto L813
	}
L813:
	;
	v2167 = v47 + int32(56)
	goto L824
L814:
	;
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v2138 == v2154 {
		goto L810
	} else {
		goto L823
	}
L815:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v2138 != v2152 {
		goto L813
	} else {
		goto L822
	}
L816:
	;
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v2138 != v2150 {
		goto L813
	} else {
		goto L821
	}
L817:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v2138 != v2148 {
		goto L813
	} else {
		goto L820
	}
L818:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v2138 != v2146 {
		goto L813
	} else {
		goto L819
	}
L819:
	;
	goto L810
L820:
	;
	goto L810
L821:
	;
	goto L810
L822:
	;
	goto L810
L823:
	;
	goto L813
L824:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v2167)))
	if v2173 != 0 {
		goto L826
	} else {
		goto L827
	}
L825:
	;
	F_pfree(m, v2138)
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L15
	} else {
		goto L831
	}
L826:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+40))
	if v2138 == v2174 {
		goto L810
	} else {
		goto L829
	}
L827:
	;
	goto L828
L828:
	;
	goto L825
L829:
	;
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v2173)+56))
	if v2138 != v2176 {
		v2167 = v2173
		goto L824
	} else {
		goto L830
	}
L830:
	;
	goto L810
L831:
	;
	goto L810
L832:
	;
	goto L806
L833:
	;
	if v405 != 0 {
		goto L136
	} else {
		goto L856
	}
L834:
	;
	v2239 = int32(1)
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v2235 == v2240 {
		v2269 = v2239
		goto L836
	} else {
		goto L837
	}
L835:
	;
	if v2269 != 0 {
		goto L833
	} else {
		goto L854
	}
L836:
	;
	goto L835
L837:
	;
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v2242 {
	case 0:
		goto L843
	case 1:
		goto L842
	case 2:
		goto L841
	case 3:
		goto L840
	case 4:
		goto L839
	default:
		goto L838
	}
L838:
	;
	v2255 = v47 + int32(56)
	goto L849
L839:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v2235 == v2251 {
		v2269 = v2239
		goto L836
	} else {
		goto L848
	}
L840:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v2235 != v2249 {
		goto L838
	} else {
		goto L847
	}
L841:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v2235 != v2247 {
		goto L838
	} else {
		goto L846
	}
L842:
	;
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v2235 != v2245 {
		goto L838
	} else {
		goto L845
	}
L843:
	;
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v2235 != v2243 {
		goto L838
	} else {
		goto L844
	}
L844:
	;
	v2269 = v2239
	goto L836
L845:
	;
	v2269 = v2239
	goto L836
L846:
	;
	v2269 = v2239
	goto L836
L847:
	;
	v2269 = v2239
	goto L836
L848:
	;
	goto L838
L849:
	;
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(v2255)))
	v2259 = int32(0)
	v2260 = base.B2i32(v2258 != v2259)
	if v2258 == v2259 {
		v2269 = v2260
		goto L836
	} else {
		goto L851
	}
L850:
	;
	v2269 = v2260
	goto L836
L851:
	;
	v2263 = *(*int32)(unsafe.Add(mBase, uint32(v2258)+40))
	if v2235 == v2263 {
		v2269 = v2260
		goto L836
	} else {
		goto L852
	}
L852:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(v2258)+56))
	if v2235 != v2265 {
		v2255 = v2258
		goto L849
	} else {
		goto L853
	}
L853:
	;
	goto L850
L854:
	;
	F_pfree(m, v2235)
	mBase = m.M
	v2271 = m.ExcPending
	if v2271 != 0 {
		goto L15
	} else {
		goto L855
	}
L855:
	;
	goto L833
L856:
	;
	v2321 = int32(-1)
	goto L11
L857:
	;
	v2321 = int32(-1)
	goto L11
L858:
	;
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	if v2297&int32(4) != 0 {
		v2321 = v2291
		goto L11
	} else {
		goto L859
	}
L859:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v2297 | int32(4)
	v2303 = int32(4484964)
	v2304 = *(*int32)(unsafe.Add(mBase, _consts[1292]))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+76)) = v2304
	*(*int32)(unsafe.Add(mBase, _consts[1292])) = v47 + int32(76)
	v2321 = v2291
	goto L11
}
