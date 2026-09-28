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
	var v89 int32
	_ = v89
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
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
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
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
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
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
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
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = int32(_a_F_ParseConfigFile_0)
	v20 = m.G0
	v22 = v20 - int32(32)
	v23 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v23
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParseConfigFile[0])))
	if v31 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(96)
	return v346
L2:
	;
	v100 = F_strlen(m, l0)
	mBase = m.M
	if v99 == v100 {
		goto L21
	} else {
		goto L22
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
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParseConfigFile[1])))
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
		v89 = l0
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v99 = v89 - l0
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
		v89 = v70
		goto L15
	} else {
		goto L19
	}
L18:
	;
	v89 = v87
	goto L15
L19:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	v87 = v70 + int32(1)
	if v85 != 0 {
		v70 = v87
		v71 = v85
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v102 = int32(0)
	v104 = F_errstart(m, l5, v102)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	if int32(11) <= l4 {
		goto L42
	} else {
		goto L43
	}
L24:
	;
	return int32(0)
L25:
	;
	if v104 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v121 = F_palloc(m, int32(28))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L24
	} else {
		goto L32
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg(m, int32(_a_F_ParseConfigFile_1), v14)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_ParseConfigFile_2), int32(194), int32(_a_F_ParseConfigFile_3))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v121))) = int64(0)
	v126 = F_pstrdup(m, int32(_a_F_ParseConfigFile_4))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = v126
	if l2 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v129 = F_pstrdup(m, l2)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L24
	} else {
		goto L37
	}
L35:
	;
	v131 = v102
	goto L36
L36:
	;
	v132 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+24)) = v132
	v134 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v121)+20)) = uint16(v134)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v121)+12)) = v131
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v138 == v132 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v131 = v129
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v121
	v346 = int32(0)
	goto L1
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v121
	goto L38
L40:
	;
	goto L41
L41:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(v142)+24)) = v121
	goto L38
L42:
	;
	v148 = int32(0)
	v150 = F_errstart(m, l5, v148)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L24
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v192 = F_AbsoluteConfigLocation(m, l0, l2)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L24
	} else {
		goto L62
	}
L45:
	;
	if v150 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L24
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v167 = F_palloc(m, int32(28))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L24
	} else {
		goto L52
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l0
	F_errmsg(m, int32(_a_F_ParseConfigFile_5), v14+int32(16))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L24
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_ParseConfigFile_2), int32(211), int32(_a_F_ParseConfigFile_3))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L24
	} else {
		goto L51
	}
L51:
	;
	goto L48
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v167))) = int64(0)
	v172 = F_pstrdup(m, int32(_a_F_ParseConfigFile_6))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L24
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = v172
	if l2 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v175 = F_pstrdup(m, l2)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L24
	} else {
		goto L57
	}
L55:
	;
	v177 = v148
	goto L56
L56:
	;
	v178 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+24)) = v178
	v180 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v167)+20)) = uint16(v180)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v167)+12)) = v177
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v184 == v178 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v177 = v175
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v167
	v346 = int32(0)
	goto L1
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v167
	goto L58
L60:
	;
	goto L61
L61:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+24)) = v167
	goto L58
L62:
	;
	if l2 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v267 = F_AllocateFile(m, v192, int32(_a_F_ParseConfigFile_7))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L24
	} else {
		goto L89
	}
L64:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if base.B2i32(v198 == int32(0))|base.B2i32(v198 != v201) != 0 {
		v219 = v198
		v220 = v201
		goto L66
	} else {
		goto L67
	}
L65:
	;
	if v219-v220 != 0 {
		goto L63
	} else {
		goto L72
	}
L66:
	;
	goto L65
L67:
	;
	v204 = v192
	v205 = l2
	goto L68
L68:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+1)))
	if v209 == int32(0) {
		v219 = v209
		v220 = v208
		goto L66
	} else {
		goto L70
	}
L69:
	;
	v219 = v209
	v220 = v208
	goto L66
L70:
	;
	v212 = int32(1)
	if v209 == v208 {
		v204 = v204 + v212
		v205 = v205 + v212
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v223 = F_errstart(m, l5, int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L24
	} else {
		goto L73
	}
L73:
	;
	if v223 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L24
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v240 = F_palloc(m, int32(28))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L24
	} else {
		goto L80
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = l2
	F_errmsg(m, int32(_a_F_ParseConfigFile_8), v14+int32(80))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L24
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_ParseConfigFile_2), int32(231), int32(_a_F_ParseConfigFile_3))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L24
	} else {
		goto L79
	}
L79:
	;
	goto L76
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v240))) = int64(0)
	v245 = F_pstrdup(m, int32(_a_F_ParseConfigFile_9))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L24
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+8)) = v245
	v248 = F_pstrdup(m, l2)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L24
	} else {
		goto L82
	}
L82:
	;
	v250 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v240)+24)) = v250
	v252 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v240)+20)) = uint16(v252)
	*(*int32)(unsafe.Add(mBase, uint32(v240)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v240)+12)) = v248
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v256 == v250 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v240
	F_pfree(m, v192)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L24
	} else {
		goto L87
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v240
	goto L83
L85:
	;
	goto L86
L86:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(v260)+24)) = v240
	goto L83
L87:
	;
	v346 = int32(0)
	goto L1
L88:
	;
	F_pfree(m, v192)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L24
	} else {
		goto L120
	}
L89:
	;
	if v267 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if l1 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	v337 = F_ParseConfigFp(m, v267, v192, l4, l5, l6, l7)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L24
	} else {
		goto L118
	}
L93:
	;
	v271 = int32(0)
	v273 = F_errstart(m, l5, v271)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L24
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v319 = int32(1)
	v322 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L24
	} else {
		goto L114
	}
L96:
	;
	if v273 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L24
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v192
	v292 = F_psprintf(m, int32(_a_F_ParseConfigFile_10), v14+int32(32))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L24
	} else {
		goto L103
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v192
	F_errmsg(m, int32(_a_F_ParseConfigFile_11), v14+int32(48))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L24
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_ParseConfigFile_2), int32(247), int32(_a_F_ParseConfigFile_3))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L24
	} else {
		goto L102
	}
L102:
	;
	goto L99
L103:
	;
	v295 = F_palloc(m, int32(28))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L24
	} else {
		goto L104
	}
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v295))) = int64(0)
	v299 = F_pstrdup(m, v292)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L24
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295)+8)) = v299
	if l2 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v302 = F_pstrdup(m, l2)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L24
	} else {
		goto L109
	}
L107:
	;
	v304 = v271
	goto L108
L108:
	;
	v305 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v295)+24)) = v305
	v307 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v295)+20)) = uint16(v307)
	*(*int32)(unsafe.Add(mBase, uint32(v295)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v295)+12)) = v304
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v311 == v305 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v304 = v302
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v295
	v341 = int32(0)
	goto L88
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v295
	goto L110
L112:
	;
	goto L113
L113:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+24)) = v295
	goto L110
L114:
	;
	if v322 == int32(0) {
		v341 = v319
		goto L88
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v192
	F_errmsg(m, int32(_a_F_ParseConfigFile_12), v14-int32(-64))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L24
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_ParseConfigFile_2), int32(258), int32(_a_F_ParseConfigFile_3))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L24
	} else {
		goto L117
	}
L117:
	;
	v341 = v319
	goto L88
L118:
	;
	v339 = F_FreeFile(m, v267)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L24
	} else {
		goto L119
	}
L119:
	;
	v341 = v337
	goto L88
L120:
	;
	v346 = v341
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
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_SetConfigOption[0]))
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
func F_set_config_by_name(m *base.Module, l0 int32) int64 {
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
	var v23 int64
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
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v2 = int32(0)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v5 != int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v9 = F_text_to_cstring(m, v8)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v13 == int32(0) {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v17 = F_text_to_cstring(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					v19 = v17
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
					if v20 == int32(0) {
						v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
						v26 = base.B2i32(v23 != int64(0))
					} else {
						v26 = v2
					}
					v29 = F_superuser(m)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
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
							return int64(0)
						} else {
							v36 = int32(0)
							v38 = F_GetConfigOptionByName(m, v9, v36, v36)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
							} else {
								v40 = F_cstring_to_text(m, v38)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v40)
								}
							}
						}
					}
				}
			} else {
				v19 = v2
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
				if v20 == int32(0) {
					v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
					v26 = base.B2i32(v23 != int64(0))
				} else {
					v26 = v2
				}
				v29 = F_superuser(m)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
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
						return int64(0)
					} else {
						v36 = int32(0)
						v38 = F_GetConfigOptionByName(m, v9, v36, v36)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v40 = F_cstring_to_text(m, v38)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v40)
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(67108994))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_set_config_by_name_0), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_set_config_by_name_1), int32(370), int32(_a_F_set_config_by_name_2))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int64(0)
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
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
	var v195 int32
	_ = v195
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
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
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
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v608 int32
	_ = v608
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v691 int32
	_ = v691
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v726 int32
	_ = v726
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v804 int32
	_ = v804
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v957 int32
	_ = v957
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v990 int32
	_ = v990
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1040 int32
	_ = v1040
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1075 int32
	_ = v1075
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1118 float64
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 float64
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1153 int32
	_ = v1153
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 float64
	_ = v1192
	var v1193 float64
	_ = v1193
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 float64
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1232 float64
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 float64
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1283 int32
	_ = v1283
	var v1285 float64
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1308 int32
	_ = v1308
	var v1317 int32
	_ = v1317
	var v1319 float64
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1341 int32
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1391 int32
	_ = v1391
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1426 int32
	_ = v1426
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1527 int32
	_ = v1527
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1554 int32
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1585 int32
	_ = v1585
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1604 int32
	_ = v1604
	var v1622 int32
	_ = v1622
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1641 int32
	_ = v1641
	var v1646 int32
	_ = v1646
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1674 int32
	_ = v1674
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1712 int32
	_ = v1712
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1740 int32
	_ = v1740
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1753 int32
	_ = v1753
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1772 int32
	_ = v1772
	var v1781 int32
	_ = v1781
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1808 int32
	_ = v1808
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1858 int32
	_ = v1858
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1915 int32
	_ = v1915
	var v1933 int32
	_ = v1933
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1947 int32
	_ = v1947
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1955 int32
	_ = v1955
	var v1957 int32
	_ = v1957
	var v1961 int32
	_ = v1961
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1969 int32
	_ = v1969
	var v1981 int32
	_ = v1981
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1996 int32
	_ = v1996
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2033 int32
	_ = v2033
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2049 int32
	_ = v2049
	var v2051 int32
	_ = v2051
	var v2059 int32
	_ = v2059
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2075 int32
	_ = v2075
	var v2078 int32
	_ = v2078
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2119 int32
	_ = v2119
	var v2124 int32
	_ = v2124
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2149 int32
	_ = v2149
	var v2155 int32
	_ = v2155
	var v2157 int32
	_ = v2157
	var v2159 int32
	_ = v2159
	var v2161 int32
	_ = v2161
	var v2164 int32
	_ = v2164
	var v2166 int32
	_ = v2166
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2187 int32
	_ = v2187
	var v2189 int32
	_ = v2189
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2199 int32
	_ = v2199
	var v2212 int32
	_ = v2212
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2230 int32
	_ = v2230
	var v2232 int32
	_ = v2232
	var v2245 int32
	_ = v2245
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2254 int32
	_ = v2254
	var v2257 int32
	_ = v2257
	var v2295 int32
	_ = v2295
	var v2313 int32
	_ = v2313
	var v2316 int32
	_ = v2316
	var v2318 int32
	_ = v2318
	var v2330 int32
	_ = v2330
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2345 int32
	_ = v2345
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2388 int32
	_ = v2388
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2412 int32
	_ = v2412
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
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_set_config_with_handle[0])))
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
	return v2412
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
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[1]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+72))
	if v51 != 0 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	return int32(0)
L16:
	;
	if v41 == int32(0) {
		v2412 = v11
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
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	switch v91 {
	case 0:
		goto L40
	case 1:
		goto L39
	case 2:
		goto L38
	case 3:
		goto L37
	case 4:
		goto L36
	case 5:
		goto L33
	default:
		v263 = v90
		goto L31
	}
L19:
	;
	v57 = int32(0)
	if base.B2i32(v54&int32(1) == v57)|(base.B2i32(l7 == v57)|base.B2i32(l6 == int32(2))) != 0 {
		goto L18
	} else {
		goto L23
	}
L20:
	;
	v54 = int32(1)
	goto L22
L21:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+76)))
	v54 = v53
	goto L22
L22:
	;
	goto L19
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[2]))
	if v66 == int32(4) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+21)))
	if v69&int32(128) != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v73 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	if v73 == int32(0) {
		v2412 = v11
		goto L11
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v80
	F_errmsg(m, int32(_a_F_set_config_with_handle_0), v20)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3377), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v2412 = v11
	goto L11
L31:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	if v264&int32(_a_F_set_config_with_handle_3) != 0 {
		goto L91
	} else {
		goto L92
	}
L32:
	;
	v263 = int32(0)
	goto L31
L33:
	;
	if l3&int32(-3) != int32(4) {
		goto L32
	} else {
		goto L80
	}
L34:
	;
	if l4 == int32(9) {
		goto L32
	} else {
		goto L70
	}
L35:
	;
	if l9|base.B2i32(l7 == int32(0)) != 0 {
		v263 = v90
		goto L31
	} else {
		goto L68
	}
L36:
	;
	if l3 != int32(2) {
		goto L34
	} else {
		goto L67
	}
L37:
	;
	switch l3 - int32(2) {
	case 0:
		goto L35
	default:
		goto L34
	case 2:
		goto L59
	}
L38:
	;
	if base.Ui32(int32(-3)) < base.Ui32(l3-int32(3)) {
		v263 = v90
		goto L31
	} else {
		goto L53
	}
L39:
	;
	v114 = int32(1)
	switch l3 - v114 {
	case 0:
		goto L32
	case 1:
		v263 = v114
		goto L31
	default:
		goto L47
	}
L40:
	;
	if l3 == int32(0) {
		v263 = v90
		goto L31
	} else {
		goto L41
	}
L41:
	;
	v95 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	if v95 == int32(0) {
		v2412 = v11
		goto L11
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v102
	F_errmsg(m, int32(_a_F_set_config_with_handle_4), v20+int32(176))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3393), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	v2412 = v11
	goto L11
L47:
	;
	v118 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	if v118 == int32(0) {
		v2412 = v11
		goto L11
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L15
	} else {
		goto L50
	}
L50:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v125
	F_errmsg(m, int32(_a_F_set_config_with_handle_5), v20+int32(192))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L15
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3416), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L15
	} else {
		goto L52
	}
L52:
	;
	v2412 = v11
	goto L11
L53:
	;
	v142 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L15
	} else {
		goto L54
	}
L54:
	;
	if v142 == int32(0) {
		v2412 = v11
		goto L11
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L15
	} else {
		goto L56
	}
L56:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v149
	F_errmsg(m, int32(_a_F_set_config_with_handle_6), v20+int32(208))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L15
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3426), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L15
	} else {
		goto L58
	}
L58:
	;
	v2412 = v11
	goto L11
L59:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v165 = F_pg_parameter_aclcheck(m, v163, l5, int64(4096))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L15
	} else {
		goto L60
	}
L60:
	;
	if v165 == int32(0) {
		goto L34
	} else {
		goto L61
	}
L61:
	;
	v170 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L15
	} else {
		goto L62
	}
L62:
	;
	if v170 == int32(0) {
		v2412 = v11
		goto L11
	} else {
		goto L63
	}
L63:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L15
	} else {
		goto L64
	}
L64:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v177
	F_errmsg(m, int32(_a_F_set_config_with_handle_7), v20+int32(240))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3453), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L15
	} else {
		goto L66
	}
L66:
	;
	v2412 = v11
	goto L11
L67:
	;
	goto L35
L68:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_set_config_with_handle[0])))
	if v195&int32(1) == int32(0) {
		v263 = v90
		goto L31
	} else {
		goto L69
	}
L69:
	;
	v2412 = int32(-1)
	goto L11
L70:
	;
	if int32(1)<<(uint(l3)%32)&int32(26) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v210 = base.B2i32(base.Ui32(l3) <= base.Ui32(int32(4)))
	goto L73
L72:
	;
	v210 = int32(0)
	goto L73
L73:
	;
	if v210 != 0 {
		v263 = v90
		goto L31
	} else {
		goto L74
	}
L74:
	;
	v212 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L15
	} else {
		goto L75
	}
L75:
	;
	if v212 == int32(0) {
		v2412 = v11
		goto L11
	} else {
		goto L76
	}
L76:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L15
	} else {
		goto L77
	}
L77:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = v219
	F_errmsg(m, int32(_a_F_set_config_with_handle_8), v20+int32(224))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L15
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3495), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L15
	} else {
		goto L79
	}
L79:
	;
	v2412 = v11
	goto L11
L80:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v237 = F_pg_parameter_aclcheck(m, v235, l5, int64(4096))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L15
	} else {
		goto L81
	}
L81:
	;
	if v237 == int32(0) {
		v263 = v90
		goto L31
	} else {
		goto L82
	}
L82:
	;
	v242 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L15
	} else {
		goto L83
	}
L83:
	;
	if v242 == int32(0) {
		v2412 = v11
		goto L11
	} else {
		goto L84
	}
L84:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L15
	} else {
		goto L85
	}
L85:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v249
	F_errmsg(m, int32(_a_F_set_config_with_handle_7), v20+int32(256))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L15
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3515), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L15
	} else {
		goto L87
	}
L87:
	;
	v2412 = v11
	goto L11
L88:
	;
	v369 = int32(0)
	v381 = base.B2i32(base.B2i32(l7 == v369)|base.B2i32(base.Ui32(int32(10)) < base.Ui32(l4)) == v369) & (base.B2i32(l4 == v369) | base.B2i32(l2 != v369))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if base.Ui32(v382) <= base.Ui32(l4) {
		v407 = l7
		goto L124
	} else {
		goto L125
	}
L89:
	;
	if l2 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L90:
	;
	v302 = int32(0)
	v304 = F_errstart(m, v36, v302)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L15
	} else {
		goto L105
	}
L91:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[3]))
	if v268&int32(1) != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v299 = v264
	goto L93
L93:
	;
	if v299&int32(8) != 0 {
		goto L89
	} else {
		goto L104
	}
L94:
	;
	v271 = int32(0)
	v273 = F_errstart(m, v36, v271)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L15
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_set_config_with_handle[3])))
	goto L102
L97:
	;
	if v273 == int32(0) {
		v2412 = v271
		goto L11
	} else {
		goto L98
	}
L98:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L15
	} else {
		goto L99
	}
L99:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v280
	F_errmsg(m, int32(_a_F_set_config_with_handle_9), v20+int32(144))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L15
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3554), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L15
	} else {
		goto L101
	}
L101:
	;
	v2412 = v271
	goto L11
L102:
	;
	if int32(base.Ui32(v293&int32(2))>>(uint(int32(1))%32)) != 0 {
		goto L90
	} else {
		goto L103
	}
L103:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v299 = v298
	goto L93
L104:
	;
	goto L88
L105:
	;
	if v304 == int32(0) {
		v2412 = v302
		goto L11
	} else {
		goto L106
	}
L106:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L15
	} else {
		goto L107
	}
L107:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v311
	F_errmsg(m, int32(_a_F_set_config_with_handle_10), v20+int32(160))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L15
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3562), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L15
	} else {
		goto L109
	}
L109:
	;
	v2412 = v302
	goto L11
L110:
	;
	v325 = int32(0)
	v327 = F_errstart(m, v36, v325)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L15
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	if l6 != int32(2) {
		goto L88
	} else {
		goto L118
	}
L113:
	;
	if v327 == int32(0) {
		v2412 = v325
		goto L11
	} else {
		goto L114
	}
L114:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L15
	} else {
		goto L115
	}
L115:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v334
	F_errmsg(m, int32(_a_F_set_config_with_handle_11), v20+int32(112))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L15
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3574), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L15
	} else {
		goto L117
	}
L117:
	;
	v2412 = v325
	goto L11
L118:
	;
	v348 = int32(0)
	v350 = F_errstart(m, v36, v348)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L15
	} else {
		goto L119
	}
L119:
	;
	if v350 == int32(0) {
		v2412 = v348
		goto L11
	} else {
		goto L120
	}
L120:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L15
	} else {
		goto L121
	}
L121:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v357
	F_errmsg(m, int32(_a_F_set_config_with_handle_12), v20+int32(128))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L15
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3582), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L15
	} else {
		goto L123
	}
L123:
	;
	v2412 = v348
	goto L11
L124:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v408 {
	case 0:
		goto L137
	case 1:
		goto L136
	case 2:
		goto L135
	case 3:
		goto L134
	case 4:
		goto L133
	default:
		goto L132
	}
L125:
	;
	if l7^int32(1)|v381 != 0 {
		v407 = int32(0)
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v388 = int32(-1)
	v391 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L15
	} else {
		goto L127
	}
L127:
	;
	if v391 == int32(0) {
		v2412 = v388
		goto L11
	} else {
		goto L128
	}
L128:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v395
	F_errmsg_internal(m, int32(_a_F_set_config_with_handle_13), v20+int32(96))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L15
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3608), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L15
	} else {
		goto L130
	}
L130:
	;
	v2412 = v388
	goto L11
L131:
	;
	v2382 = int32(1)
	v2383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+20)))
	if v2383&int32(64) == int32(0) {
		v2412 = v2382
		goto L11
	} else {
		goto L691
	}
L132:
	;
	if v407 != 0 {
		goto L131
	} else {
		goto L690
	}
L133:
	;
	if l2 != 0 {
		goto L596
	} else {
		goto L597
	}
L134:
	;
	if l2 != 0 {
		goto L424
	} else {
		goto L425
	}
L135:
	;
	if l2 != 0 {
		goto L329
	} else {
		goto L330
	}
L136:
	;
	if l2 != 0 {
		goto L234
	} else {
		goto L235
	}
L137:
	;
	if l2 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if v263 != 0 {
		goto L149
	} else {
		goto L150
	}
L139:
	;
	v413 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L15
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	if l4 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	if v413 != 0 {
		v436 = l3
		v437 = l4
		v438 = l5
		goto L138
	} else {
		goto L143
	}
L143:
	;
	v2412 = int32(0)
	goto L11
L144:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)) = uint8(v418)
	v420 = int32(0)
	v427 = F_call_bool_check_hook(m, v47, v20+int32(264), v20+int32(260), v420, v36)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L15
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+116)))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)) = uint8(v429)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v431
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v436 = v434
	v437 = v435
	v438 = v433
	goto L138
L147:
	;
	if v427 != 0 {
		v436 = l3
		v437 = v420
		v438 = l5
		goto L138
	} else {
		goto L148
	}
L148:
	;
	v2412 = v420
	goto L11
L149:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v440 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	goto L151
L151:
	;
	if v407 != 0 {
		goto L172
	} else {
		goto L173
	}
L152:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491))))
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	if v492 != v493 {
		goto L164
	} else {
		goto L165
	}
L153:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v440 == v443 {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v440 == v445 {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v453 = v47 + int32(56)
	goto L156
L156:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	if v466 != 0 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	F_pfree(m, v440)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L15
	} else {
		goto L163
	}
L158:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v466)+40))
	if v440 == v467 {
		goto L152
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	goto L157
L161:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v466)+56))
	if v440 != v469 {
		v453 = v466
		goto L156
	} else {
		goto L162
	}
L162:
	;
	goto L152
L163:
	;
	goto L152
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v490 | int32(2)
	v498 = int32(0)
	v500 = F_errstart(m, v36, v498)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L15
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v490 & int32(-3)
	v2412 = int32(-1)
	goto L11
L167:
	;
	if v500 == int32(0) {
		v2412 = v498
		goto L11
	} else {
		goto L168
	}
L168:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L15
	} else {
		goto L169
	}
L169:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v507
	F_errmsg(m, int32(_a_F_set_config_with_handle_5), v20+int32(16))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L15
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3660), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L15
	} else {
		goto L171
	}
L171:
	;
	v2412 = v498
	goto L11
L172:
	;
	if v381 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	goto L174
L174:
	;
	if v381 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L175:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L15
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v47)+108))
	if v528 != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	goto L177
L179:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	m.T0[v528].(func(*base.Module, int32, int32))(m, v527&int32(1), v531)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L15
	} else {
		goto L182
	}
L180:
	;
	v535 = v527
	goto L181
L181:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	*(*uint8)(unsafe.Add(mBase, uint32(v536))) = uint8(v535)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(60), v540)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L15
	} else {
		goto L183
	}
L182:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	v535 = v534
	goto L181
L183:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v545 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v436
	goto L174
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v437
	goto L184
L186:
	;
	if v437 == int32(0) {
		goto L185
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	if v437 != 0 {
		goto L185
	} else {
		goto L194
	}
L189:
	;
	v551 = v47 + int32(68)
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4]))
	if v553 != 0 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v560
	v562 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+72)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(v560)+4)) = v551
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[5])) = v551
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v437
	goto L184
L191:
	;
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[5]))
	v560 = v555
	goto L190
L192:
	;
	goto L193
L193:
	;
	v557 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4])) = v557
	v560 = v557
	goto L190
L194:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v47)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v568)+4)) = v569
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = v571
	goto L185
L195:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v709 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L196:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v583) <= base.Ui32(v437) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+116)) = uint8(v585)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47-int32(-64), v589)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L15
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v595 == int32(0) {
		goto L195
	} else {
		goto L201
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v437
	goto L199
L201:
	;
	v608 = v595
	goto L202
L202:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v608)+12))
	if base.Ui32(v617) <= base.Ui32(v437) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	goto L195
L204:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	*(*uint8)(unsafe.Add(mBase, uint32(v608)+32)) = uint8(v619)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v608)+40))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v608)+40)) = v622
	if v621 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	goto L206
L206:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v608)))
	if v691 != 0 {
		v608 = v691
		goto L202
	} else {
		goto L219
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v608)+24)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v608)+16)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v608)+12)) = v437
	goto L206
L208:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v621 == v626 {
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v621 == v628 {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	v641 = v47 + int32(56)
	goto L211
L211:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	if v647 != 0 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	F_pfree(m, v621)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L15
	} else {
		goto L218
	}
L213:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v647)+40))
	if v621 == v648 {
		goto L207
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	goto L212
L216:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v647)+56))
	if v621 != v650 {
		v641 = v647
		goto L211
	} else {
		goto L217
	}
L217:
	;
	goto L207
L218:
	;
	goto L207
L219:
	;
	goto L203
L220:
	;
	if v407 != 0 {
		goto L131
	} else {
		goto L232
	}
L221:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v709 == v712 {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v709 == v714 {
		goto L220
	} else {
		goto L223
	}
L223:
	;
	v726 = v47 + int32(56)
	goto L224
L224:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	if v735 != 0 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	F_pfree(m, v709)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L15
	} else {
		goto L231
	}
L226:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v735)+40))
	if v709 == v736 {
		goto L220
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	goto L225
L229:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v735)+56))
	if v709 != v738 {
		v726 = v735
		goto L224
	} else {
		goto L230
	}
L230:
	;
	goto L220
L231:
	;
	goto L220
L232:
	;
	v2412 = int32(-1)
	goto L11
L233:
	;
	if v263 != 0 {
		goto L244
	} else {
		goto L245
	}
L234:
	;
	v764 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L15
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	if l4 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L237:
	;
	if v764 != 0 {
		v787 = l3
		v788 = l4
		v789 = l5
		goto L233
	} else {
		goto L238
	}
L238:
	;
	v2412 = int32(0)
	goto L11
L239:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v47)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v769
	v771 = int32(0)
	v778 = F_call_int_check_hook(m, v47, v20+int32(264), v20+int32(260), v771, v36)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L15
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v780
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v782
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v787 = v785
	v788 = v786
	v789 = v784
	goto L233
L242:
	;
	if v778 != 0 {
		v787 = l3
		v788 = v771
		v789 = l5
		goto L233
	} else {
		goto L243
	}
L243:
	;
	v2412 = v771
	goto L11
L244:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v791 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	goto L246
L246:
	;
	if v407 != 0 {
		goto L267
	} else {
		goto L268
	}
L247:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v842)))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if v843 != v844 {
		goto L259
	} else {
		goto L260
	}
L248:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v791 == v794 {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v791 == v796 {
		goto L247
	} else {
		goto L250
	}
L250:
	;
	v804 = v47 + int32(56)
	goto L251
L251:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v804)))
	if v817 != 0 {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	F_pfree(m, v791)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L15
	} else {
		goto L258
	}
L253:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v817)+40))
	if v791 == v818 {
		goto L247
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	goto L252
L256:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v817)+56))
	if v791 != v820 {
		v804 = v817
		goto L251
	} else {
		goto L257
	}
L257:
	;
	goto L247
L258:
	;
	goto L247
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v841 | int32(2)
	v849 = int32(0)
	v851 = F_errstart(m, v36, v849)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L15
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v841 & int32(-3)
	v2412 = int32(-1)
	goto L11
L262:
	;
	if v851 == int32(0) {
		v2412 = v849
		goto L11
	} else {
		goto L263
	}
L263:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L15
	} else {
		goto L264
	}
L264:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v858
	F_errmsg(m, int32(_a_F_set_config_with_handle_5), v20+int32(32))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L15
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3756), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L15
	} else {
		goto L266
	}
L266:
	;
	v2412 = v849
	goto L11
L267:
	;
	if v381 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	goto L269
L269:
	;
	if v381 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L270:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L15
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v879 != 0 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	goto L272
L274:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	m.T0[v879].(func(*base.Module, int32, int32))(m, v878, v880)
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L15
	} else {
		goto L277
	}
L275:
	;
	v884 = v878
	goto L276
L276:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v885))) = v884
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(60), v889)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L15
	} else {
		goto L278
	}
L277:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v884 = v883
	goto L276
L278:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v894 == int32(0) {
		goto L281
	} else {
		goto L282
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v789
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v787
	goto L269
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v788
	goto L279
L281:
	;
	if v788 == int32(0) {
		goto L280
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	if v788 != 0 {
		goto L280
	} else {
		goto L289
	}
L284:
	;
	v900 = v47 + int32(68)
	v902 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4]))
	if v902 != 0 {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v909
	v911 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+72)) = v911
	*(*int32)(unsafe.Add(mBase, uint32(v909)+4)) = v900
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[5])) = v900
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v788
	goto L279
L286:
	;
	v904 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[5]))
	v909 = v904
	goto L285
L287:
	;
	goto L288
L288:
	;
	v906 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4])) = v906
	v909 = v906
	goto L285
L289:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v47)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v917)+4)) = v918
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v918))) = v920
	goto L280
L290:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1058 == int32(0) {
		goto L315
	} else {
		goto L316
	}
L291:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v932) <= base.Ui32(v788) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+124)) = v934
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47-int32(-64), v938)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L15
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v944 == int32(0) {
		goto L290
	} else {
		goto L296
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v789
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v788
	goto L294
L296:
	;
	v957 = v944
	goto L297
L297:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v957)+12))
	if base.Ui32(v966) <= base.Ui32(v788) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	goto L290
L299:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v957)+32)) = v968
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v957)+40))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v957)+40)) = v971
	if v970 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	goto L301
L301:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v957)))
	if v1040 != 0 {
		v957 = v1040
		goto L297
	} else {
		goto L314
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v957)+24)) = v789
	*(*int32)(unsafe.Add(mBase, uint32(v957)+16)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v957)+12)) = v788
	goto L301
L303:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v970 == v975 {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v970 == v977 {
		goto L302
	} else {
		goto L305
	}
L305:
	;
	v990 = v47 + int32(56)
	goto L306
L306:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v990)))
	if v996 != 0 {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	F_pfree(m, v970)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L15
	} else {
		goto L313
	}
L308:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v996)+40))
	if v970 == v997 {
		goto L302
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	goto L307
L311:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v996)+56))
	if v970 != v999 {
		v990 = v996
		goto L306
	} else {
		goto L312
	}
L312:
	;
	goto L302
L313:
	;
	goto L302
L314:
	;
	goto L298
L315:
	;
	if v407 != 0 {
		goto L131
	} else {
		goto L327
	}
L316:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1058 == v1061 {
		goto L315
	} else {
		goto L317
	}
L317:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v1058 == v1063 {
		goto L315
	} else {
		goto L318
	}
L318:
	;
	v1075 = v47 + int32(56)
	goto L319
L319:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1075)))
	if v1084 != 0 {
		goto L321
	} else {
		goto L322
	}
L320:
	;
	F_pfree(m, v1058)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L15
	} else {
		goto L326
	}
L321:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+40))
	if v1058 == v1085 {
		goto L315
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	goto L320
L324:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+56))
	if v1058 != v1087 {
		v1075 = v1084
		goto L319
	} else {
		goto L325
	}
L325:
	;
	goto L315
L326:
	;
	goto L315
L327:
	;
	v2412 = int32(-1)
	goto L11
L328:
	;
	if v263 != 0 {
		goto L339
	} else {
		goto L340
	}
L329:
	;
	v1113 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L15
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	if l4 == int32(0) {
		goto L334
	} else {
		goto L335
	}
L332:
	;
	if v1113 != 0 {
		v1136 = l3
		v1137 = l4
		v1138 = l5
		goto L328
	} else {
		goto L333
	}
L333:
	;
	v2412 = int32(0)
	goto L11
L334:
	;
	v1118 = *(*float64)(unsafe.Add(mBase, uint32(v47)+104))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+264)) = v1118
	v1120 = int32(0)
	v1127 = F_call_real_check_hook(m, v47, v20+int32(264), v20+int32(260), v1120, v36)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L15
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	v1129 = *(*float64)(unsafe.Add(mBase, uint32(v47)+144))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+264)) = v1129
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v1131
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v1136 = v1134
	v1137 = v1135
	v1138 = v1133
	goto L328
L337:
	;
	if v1127 != 0 {
		v1136 = l3
		v1137 = v1120
		v1138 = l5
		goto L328
	} else {
		goto L338
	}
L338:
	;
	v2412 = v1120
	goto L11
L339:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1140 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L340:
	;
	goto L341
L341:
	;
	if v407 != 0 {
		goto L362
	} else {
		goto L363
	}
L342:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	v1192 = *(*float64)(unsafe.Add(mBase, uint32(v1191)))
	v1193 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	if base.F64_ne(v1192, v1193) != 0 {
		goto L354
	} else {
		goto L355
	}
L343:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1140 == v1143 {
		goto L342
	} else {
		goto L344
	}
L344:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v1140 == v1145 {
		goto L342
	} else {
		goto L345
	}
L345:
	;
	v1153 = v47 + int32(56)
	goto L346
L346:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1153)))
	if v1166 != 0 {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	F_pfree(m, v1140)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L15
	} else {
		goto L353
	}
L348:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+40))
	if v1140 == v1167 {
		goto L342
	} else {
		goto L351
	}
L349:
	;
	goto L350
L350:
	;
	goto L347
L351:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+56))
	if v1140 != v1169 {
		v1153 = v1166
		goto L346
	} else {
		goto L352
	}
L352:
	;
	goto L342
L353:
	;
	goto L342
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v1190 | int32(2)
	v1198 = int32(0)
	v1200 = F_errstart(m, v36, v1198)
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L15
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v1190 & int32(-3)
	v2412 = int32(-1)
	goto L11
L357:
	;
	if v1200 == int32(0) {
		v2412 = v1198
		goto L11
	} else {
		goto L358
	}
L358:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L15
	} else {
		goto L359
	}
L359:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v1207
	F_errmsg(m, int32(_a_F_set_config_with_handle_5), v20+int32(48))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L15
	} else {
		goto L360
	}
L360:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3852), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L15
	} else {
		goto L361
	}
L361:
	;
	v2412 = v1198
	goto L11
L362:
	;
	if v381 == int32(0) {
		goto L365
	} else {
		goto L366
	}
L363:
	;
	goto L364
L364:
	;
	if v381 == int32(0) {
		goto L385
	} else {
		goto L386
	}
L365:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L15
	} else {
		goto L368
	}
L366:
	;
	goto L367
L367:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	v1228 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v47)+132))
	if v1229 != 0 {
		goto L369
	} else {
		goto L370
	}
L368:
	;
	goto L367
L369:
	;
	m.T0[v1229].(func(*base.Module, float64, int32))(m, v1228, v1227)
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L15
	} else {
		goto L372
	}
L370:
	;
	v1234 = v1227
	v1235 = v1228
	goto L371
L371:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v1236))) = v1235
	F_set_extra_field(m, v47, v47+int32(60), v1234)
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L15
	} else {
		goto L373
	}
L372:
	;
	v1232 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	v1234 = v1233
	v1235 = v1232
	goto L371
L373:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v1244 == int32(0) {
		goto L376
	} else {
		goto L377
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v1138
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v1136
	goto L364
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1137
	goto L374
L376:
	;
	if v1137 == int32(0) {
		goto L375
	} else {
		goto L379
	}
L377:
	;
	goto L378
L378:
	;
	if v1137 != 0 {
		goto L375
	} else {
		goto L384
	}
L379:
	;
	v1250 = v47 + int32(68)
	v1252 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4]))
	if v1252 != 0 {
		goto L381
	} else {
		goto L382
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v1259
	v1261 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+72)) = v1261
	*(*int32)(unsafe.Add(mBase, uint32(v1259)+4)) = v1250
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[5])) = v1250
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1137
	goto L374
L381:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[5]))
	v1259 = v1254
	goto L380
L382:
	;
	goto L383
L383:
	;
	v1256 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4])) = v1256
	v1259 = v1256
	goto L380
L384:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v47)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1267)+4)) = v1268
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1268))) = v1270
	goto L375
L385:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1409 == int32(0) {
		goto L410
	} else {
		goto L411
	}
L386:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v1283) <= base.Ui32(v1137) {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1285 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	*(*float64)(unsafe.Add(mBase, uint32(v47)+144)) = v1285
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47-int32(-64), v1289)
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L15
	} else {
		goto L390
	}
L388:
	;
	goto L389
L389:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v1295 == int32(0) {
		goto L385
	} else {
		goto L391
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v1138
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v1136
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v1137
	goto L389
L391:
	;
	v1308 = v1295
	goto L392
L392:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+12))
	if base.Ui32(v1317) <= base.Ui32(v1137) {
		goto L394
	} else {
		goto L395
	}
L393:
	;
	goto L385
L394:
	;
	v1319 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	*(*float64)(unsafe.Add(mBase, uint32(v1308)+32)) = v1319
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1308)+40))
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v1308)+40)) = v1322
	if v1321 == int32(0) {
		goto L397
	} else {
		goto L398
	}
L395:
	;
	goto L396
L396:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1308)))
	if v1391 != 0 {
		v1308 = v1391
		goto L392
	} else {
		goto L409
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1308)+24)) = v1138
	*(*int32)(unsafe.Add(mBase, uint32(v1308)+16)) = v1136
	*(*int32)(unsafe.Add(mBase, uint32(v1308)+12)) = v1137
	goto L396
L398:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1321 == v1326 {
		goto L397
	} else {
		goto L399
	}
L399:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v1321 == v1328 {
		goto L397
	} else {
		goto L400
	}
L400:
	;
	v1341 = v47 + int32(56)
	goto L401
L401:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1341)))
	if v1347 != 0 {
		goto L403
	} else {
		goto L404
	}
L402:
	;
	F_pfree(m, v1321)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L15
	} else {
		goto L408
	}
L403:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+40))
	if v1321 == v1348 {
		goto L397
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	goto L402
L406:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1347)+56))
	if v1321 != v1350 {
		v1341 = v1347
		goto L401
	} else {
		goto L407
	}
L407:
	;
	goto L397
L408:
	;
	goto L397
L409:
	;
	goto L393
L410:
	;
	if v407 != 0 {
		goto L131
	} else {
		goto L422
	}
L411:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1409 == v1412 {
		goto L410
	} else {
		goto L412
	}
L412:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v1409 == v1414 {
		goto L410
	} else {
		goto L413
	}
L413:
	;
	v1426 = v47 + int32(56)
	goto L414
L414:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v1426)))
	if v1435 != 0 {
		goto L416
	} else {
		goto L417
	}
L415:
	;
	F_pfree(m, v1409)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L15
	} else {
		goto L421
	}
L416:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+40))
	if v1409 == v1436 {
		goto L410
	} else {
		goto L419
	}
L417:
	;
	goto L418
L418:
	;
	goto L415
L419:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+56))
	if v1409 != v1438 {
		v1426 = v1435
		goto L414
	} else {
		goto L420
	}
L420:
	;
	goto L410
L421:
	;
	goto L410
L422:
	;
	v2412 = int32(-1)
	goto L11
L423:
	;
	if v263 != 0 {
		goto L442
	} else {
		goto L443
	}
L424:
	;
	v1464 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L15
	} else {
		goto L427
	}
L425:
	;
	goto L426
L426:
	;
	if l4 == int32(0) {
		goto L429
	} else {
		goto L430
	}
L427:
	;
	if v1464 != 0 {
		v1498 = l4
		v1499 = l5
		v1500 = l3
		goto L423
	} else {
		goto L428
	}
L428:
	;
	v2412 = int32(0)
	goto L11
L429:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v47)+100))
	if v1469 != 0 {
		goto L433
	} else {
		goto L434
	}
L430:
	;
	goto L431
L431:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v1490
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v1492
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v1498 = v1496
	v1499 = v1494
	v1500 = v1495
	goto L423
L432:
	;
	v1482 = F_call_string_check_hook(m, v47, v20+int32(264), v20+int32(260), int32(0), v36)
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L15
	} else {
		goto L438
	}
L433:
	;
	v1470 = F_guc_strdup(m, v36, v1469)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L15
	} else {
		goto L436
	}
L434:
	;
	goto L435
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = int32(0)
	goto L432
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v1470
	if v1470 != 0 {
		goto L432
	} else {
		goto L437
	}
L437:
	;
	v2412 = int32(0)
	goto L11
L438:
	;
	if v1482 != 0 {
		v1498 = v11
		v1499 = l5
		v1500 = l3
		goto L423
	} else {
		goto L439
	}
L439:
	;
	v1484 = int32(0)
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if v1485 == v1484 {
		v2412 = v1484
		goto L11
	} else {
		goto L440
	}
L440:
	;
	F_pfree(m, v1485)
	mBase = m.M
	v1489 = m.ExcPending
	if v1489 != 0 {
		goto L15
	} else {
		goto L441
	}
L441:
	;
	v2412 = v1484
	goto L11
L442:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1501)))
	v1503 = int32(0)
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if base.B2i32(v1502 == v1503)|base.B2i32(v1505 == v1503) == v1503 {
		goto L447
	} else {
		goto L448
	}
L443:
	;
	goto L444
L444:
	;
	if v407 == int32(0) {
		goto L490
	} else {
		goto L491
	}
L445:
	;
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1572 == int32(0) {
		goto L470
	} else {
		goto L471
	}
L446:
	;
	v1544 = int32(1)
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1545)))
	if v1505 == v1546 {
		v1568 = v1544
		goto L459
	} else {
		goto L460
	}
L447:
	;
	v1513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1502))))
	v1516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1505))))
	if base.B2i32(v1513 == int32(0))|base.B2i32(v1513 != v1516) != 0 {
		v1534 = v1513
		v1535 = v1516
		goto L451
	} else {
		goto L452
	}
L448:
	;
	goto L449
L449:
	;
	v1539 = int32(1)
	if v1505 == int32(0) {
		v1571 = v1539
		goto L445
	} else {
		goto L457
	}
L450:
	;
	v1542 = base.B2i32(v1534-v1535 != int32(0))
	goto L446
L451:
	;
	goto L450
L452:
	;
	v1519 = v1502
	v1520 = v1505
	goto L453
L453:
	;
	v1523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1520)+1)))
	v1524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1519)+1)))
	if v1524 == int32(0) {
		v1534 = v1524
		v1535 = v1523
		goto L451
	} else {
		goto L455
	}
L454:
	;
	v1534 = v1524
	v1535 = v1523
	goto L451
L455:
	;
	v1527 = int32(1)
	if v1524 == v1523 {
		v1519 = v1519 + v1527
		v1520 = v1520 + v1527
		goto L453
	} else {
		goto L456
	}
L456:
	;
	goto L454
L457:
	;
	v1542 = v1539
	goto L446
L458:
	;
	if v1568 != 0 {
		v1571 = v1542
		goto L445
	} else {
		goto L468
	}
L459:
	;
	goto L458
L460:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1505 == v1548 {
		v1568 = v1544
		goto L459
	} else {
		goto L461
	}
L461:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v47)+100))
	if v1505 == v1550 {
		v1568 = v1544
		goto L459
	} else {
		goto L462
	}
L462:
	;
	v1554 = v47 + int32(56)
	goto L463
L463:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1554)))
	v1558 = int32(0)
	v1559 = base.B2i32(v1557 != v1558)
	if v1557 == v1558 {
		v1568 = v1559
		goto L459
	} else {
		goto L465
	}
L464:
	;
	v1568 = v1559
	goto L459
L465:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1557)+32))
	if v1505 == v1562 {
		v1568 = v1559
		goto L459
	} else {
		goto L466
	}
L466:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1557)+48))
	if v1505 != v1564 {
		v1554 = v1557
		goto L463
	} else {
		goto L467
	}
L467:
	;
	goto L464
L468:
	;
	F_pfree(m, v1505)
	mBase = m.M
	v1570 = m.ExcPending
	if v1570 != 0 {
		goto L15
	} else {
		goto L469
	}
L469:
	;
	v1571 = v1542
	goto L445
L470:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	if v1571 != 0 {
		goto L482
	} else {
		goto L483
	}
L471:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1572 == v1575 {
		goto L470
	} else {
		goto L472
	}
L472:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v1572 == v1577 {
		goto L470
	} else {
		goto L473
	}
L473:
	;
	v1585 = v47 + int32(56)
	goto L474
L474:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v1585)))
	if v1598 != 0 {
		goto L476
	} else {
		goto L477
	}
L475:
	;
	F_pfree(m, v1572)
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L15
	} else {
		goto L481
	}
L476:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v1598)+40))
	if v1572 == v1599 {
		goto L470
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	goto L475
L479:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1598)+56))
	if v1572 != v1601 {
		v1585 = v1598
		goto L474
	} else {
		goto L480
	}
L480:
	;
	goto L470
L481:
	;
	goto L470
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v1622 | int32(2)
	v1626 = int32(0)
	v1628 = F_errstart(m, v36, v1626)
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L15
	} else {
		goto L485
	}
L483:
	;
	goto L484
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v1622 & int32(-3)
	v2412 = int32(-1)
	goto L11
L485:
	;
	if v1628 == int32(0) {
		v2412 = v1626
		goto L11
	} else {
		goto L486
	}
L486:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v1634 = m.ExcPending
	if v1634 != 0 {
		goto L15
	} else {
		goto L487
	}
L487:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v1635
	F_errmsg(m, int32(_a_F_set_config_with_handle_5), v20-int32(-64))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L15
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3977), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L15
	} else {
		goto L489
	}
L489:
	;
	v2412 = v1626
	goto L11
L490:
	;
	if v381 == int32(0) {
		goto L529
	} else {
		goto L530
	}
L491:
	;
	if v381 == int32(0) {
		goto L492
	} else {
		goto L493
	}
L492:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L15
	} else {
		goto L495
	}
L493:
	;
	goto L494
L494:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v47)+108))
	if v1658 != 0 {
		goto L496
	} else {
		goto L497
	}
L495:
	;
	goto L494
L496:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	m.T0[v1658].(func(*base.Module, int32, int32))(m, v1657, v1659)
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L15
	} else {
		goto L499
	}
L497:
	;
	v1663 = v1657
	goto L498
L498:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	F_set_string_field(m, v47, v1664, v1663)
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L15
	} else {
		goto L500
	}
L499:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v1663 = v1662
	goto L498
L500:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(60), v1669)
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L15
	} else {
		goto L501
	}
L501:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v1674 == int32(0) {
		goto L504
	} else {
		goto L505
	}
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v1499
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v1500
	if l9 != 0 {
		goto L490
	} else {
		goto L513
	}
L503:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1498
	goto L502
L504:
	;
	if v1498 == int32(0) {
		goto L503
	} else {
		goto L507
	}
L505:
	;
	goto L506
L506:
	;
	if v1498 != 0 {
		goto L503
	} else {
		goto L512
	}
L507:
	;
	v1680 = v47 + int32(68)
	v1682 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4]))
	if v1682 != 0 {
		goto L509
	} else {
		goto L510
	}
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v1689
	v1691 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+72)) = v1691
	*(*int32)(unsafe.Add(mBase, uint32(v1689)+4)) = v1680
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[5])) = v1680
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1498
	goto L502
L509:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[5]))
	v1689 = v1684
	goto L508
L510:
	;
	goto L511
L511:
	;
	v1686 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4])) = v1686
	v1689 = v1686
	goto L508
L512:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v47)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1697)+4)) = v1698
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1698))) = v1700
	goto L503
L513:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v1709 = int32(_a_F_set_config_with_handle_15)
	v1712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1708))))
	v1715 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_set_config_with_handle[6])))
	if base.B2i32(v1712 == int32(0))|base.B2i32(v1712 != v1715) != 0 {
		v1733 = v1712
		v1734 = v1715
		goto L515
	} else {
		goto L516
	}
L514:
	;
	if v1733-v1734 != 0 {
		goto L490
	} else {
		goto L521
	}
L515:
	;
	goto L514
L516:
	;
	v1718 = v1708
	v1719 = v1709
	goto L517
L517:
	;
	v1722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1719)+1)))
	v1723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1718)+1)))
	if v1723 == int32(0) {
		v1733 = v1723
		v1734 = v1722
		goto L515
	} else {
		goto L519
	}
L518:
	;
	v1733 = v1723
	v1734 = v1722
	goto L515
L519:
	;
	v1726 = int32(1)
	if v1723 == v1722 {
		v1718 = v1718 + v1726
		v1719 = v1719 + v1726
		goto L517
	} else {
		goto L520
	}
L520:
	;
	goto L518
L521:
	;
	v1737 = int32(0)
	if l2 != 0 {
		goto L522
	} else {
		goto L523
	}
L522:
	;
	v1740 = int32(_a_F_set_config_with_handle_16)
	goto L524
L523:
	;
	v1740 = v1737
	goto L524
L524:
	;
	if l4 == int32(10) {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v1744 = int32(1)
	goto L527
L526:
	;
	v1744 = l4
	goto L527
L527:
	;
	v1747 = F_set_config_with_handle(m, int32(_a_F_set_config_with_handle_17), v1737, v1740, l3, v1744, l5, l6, int32(1), v36, int32(0))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L15
	} else {
		goto L528
	}
L528:
	;
	goto L490
L529:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if v1933 == int32(0) {
		goto L568
	} else {
		goto L569
	}
L530:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v1753) <= base.Ui32(v1498) {
		goto L531
	} else {
		goto L532
	}
L531:
	;
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	F_set_string_field(m, v47, v47+int32(116), v1757)
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L15
	} else {
		goto L534
	}
L532:
	;
	goto L533
L533:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v1768 == int32(0) {
		goto L529
	} else {
		goto L536
	}
L534:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47-int32(-64), v1762)
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L15
	} else {
		goto L535
	}
L535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v1499
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v1500
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v1498
	goto L533
L536:
	;
	v1772 = v47 + int32(56)
	v1781 = v1768
	goto L537
L537:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v1781)+12))
	if base.Ui32(v1790) <= base.Ui32(v1498) {
		goto L539
	} else {
		goto L540
	}
L538:
	;
	goto L529
L539:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1781)+32))
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+32)) = v1793
	if v1792 == int32(0) {
		goto L542
	} else {
		goto L543
	}
L540:
	;
	goto L541
L541:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1781)))
	if v1915 != 0 {
		v1781 = v1915
		goto L537
	} else {
		goto L567
	}
L542:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1781)+40))
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+40)) = v1846
	if v1845 == int32(0) {
		goto L555
	} else {
		goto L556
	}
L543:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1797)))
	if v1792 == v1798 {
		goto L542
	} else {
		goto L544
	}
L544:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1792 == v1800 {
		goto L542
	} else {
		goto L545
	}
L545:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v47)+100))
	if v1792 == v1802 {
		goto L542
	} else {
		goto L546
	}
L546:
	;
	v1808 = v1772
	goto L547
L547:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1808)))
	if v1821 != 0 {
		goto L549
	} else {
		goto L550
	}
L548:
	;
	F_pfree(m, v1792)
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L15
	} else {
		goto L554
	}
L549:
	;
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+32))
	if v1792 == v1822 {
		goto L542
	} else {
		goto L552
	}
L550:
	;
	goto L551
L551:
	;
	goto L548
L552:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+48))
	if v1792 != v1824 {
		v1808 = v1821
		goto L547
	} else {
		goto L553
	}
L553:
	;
	goto L542
L554:
	;
	goto L542
L555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+24)) = v1499
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+16)) = v1500
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+12)) = v1498
	goto L541
L556:
	;
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1845 == v1850 {
		goto L555
	} else {
		goto L557
	}
L557:
	;
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v1845 == v1852 {
		goto L555
	} else {
		goto L558
	}
L558:
	;
	v1858 = v1772
	goto L559
L559:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1858)))
	if v1871 != 0 {
		goto L561
	} else {
		goto L562
	}
L560:
	;
	F_pfree(m, v1845)
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L15
	} else {
		goto L566
	}
L561:
	;
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1871)+40))
	if v1845 == v1872 {
		goto L555
	} else {
		goto L564
	}
L562:
	;
	goto L563
L563:
	;
	goto L560
L564:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1871)+56))
	if v1845 != v1874 {
		v1858 = v1871
		goto L559
	} else {
		goto L565
	}
L565:
	;
	goto L555
L566:
	;
	goto L555
L567:
	;
	goto L538
L568:
	;
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1964 == int32(0) {
		goto L582
	} else {
		goto L583
	}
L569:
	;
	v1937 = int32(1)
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v1938)))
	if v1933 == v1939 {
		v1961 = v1937
		goto L571
	} else {
		goto L572
	}
L570:
	;
	if v1961 != 0 {
		goto L568
	} else {
		goto L580
	}
L571:
	;
	goto L570
L572:
	;
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1933 == v1941 {
		v1961 = v1937
		goto L571
	} else {
		goto L573
	}
L573:
	;
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v47)+100))
	if v1933 == v1943 {
		v1961 = v1937
		goto L571
	} else {
		goto L574
	}
L574:
	;
	v1947 = v47 + int32(56)
	goto L575
L575:
	;
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v1947)))
	v1951 = int32(0)
	v1952 = base.B2i32(v1950 != v1951)
	if v1950 == v1951 {
		v1961 = v1952
		goto L571
	} else {
		goto L577
	}
L576:
	;
	v1961 = v1952
	goto L571
L577:
	;
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v1950)+32))
	if v1933 == v1955 {
		v1961 = v1952
		goto L571
	} else {
		goto L578
	}
L578:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v1950)+48))
	if v1933 != v1957 {
		v1947 = v1950
		goto L575
	} else {
		goto L579
	}
L579:
	;
	goto L576
L580:
	;
	F_pfree(m, v1933)
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L15
	} else {
		goto L581
	}
L581:
	;
	goto L568
L582:
	;
	if v407 != 0 {
		goto L131
	} else {
		goto L594
	}
L583:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1964 == v1967 {
		goto L582
	} else {
		goto L584
	}
L584:
	;
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v1964 == v1969 {
		goto L582
	} else {
		goto L585
	}
L585:
	;
	v1981 = v47 + int32(56)
	goto L586
L586:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v1981)))
	if v1990 != 0 {
		goto L588
	} else {
		goto L589
	}
L587:
	;
	F_pfree(m, v1964)
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L15
	} else {
		goto L593
	}
L588:
	;
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+40))
	if v1964 == v1991 {
		goto L582
	} else {
		goto L591
	}
L589:
	;
	goto L590
L590:
	;
	goto L587
L591:
	;
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1990)+56))
	if v1964 != v1993 {
		v1981 = v1990
		goto L586
	} else {
		goto L592
	}
L592:
	;
	goto L582
L593:
	;
	goto L582
L594:
	;
	v2412 = int32(-1)
	goto L11
L595:
	;
	if v263 != 0 {
		goto L606
	} else {
		goto L607
	}
L596:
	;
	v2019 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L15
	} else {
		goto L599
	}
L597:
	;
	goto L598
L598:
	;
	if l4 == int32(0) {
		goto L601
	} else {
		goto L602
	}
L599:
	;
	if v2019 != 0 {
		v2042 = l3
		v2043 = l4
		v2044 = l5
		goto L595
	} else {
		goto L600
	}
L600:
	;
	v2412 = int32(0)
	goto L11
L601:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v47)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v2024
	v2026 = int32(0)
	v2033 = F_call_enum_check_hook(m, v47, v20+int32(264), v20+int32(260), v2026, v36)
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L15
	} else {
		goto L604
	}
L602:
	;
	goto L603
L603:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v2035
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v2037
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v2040 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v2042 = v2040
	v2043 = v2041
	v2044 = v2039
	goto L595
L604:
	;
	if v2033 != 0 {
		v2042 = l3
		v2043 = v2026
		v2044 = l5
		goto L595
	} else {
		goto L605
	}
L605:
	;
	v2412 = v2026
	goto L11
L606:
	;
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v2046 == int32(0) {
		goto L609
	} else {
		goto L610
	}
L607:
	;
	goto L608
L608:
	;
	if v407 != 0 {
		goto L629
	} else {
		goto L630
	}
L609:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v2097)))
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if v2098 != v2099 {
		goto L621
	} else {
		goto L622
	}
L610:
	;
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v2046 == v2049 {
		goto L609
	} else {
		goto L611
	}
L611:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v2046 == v2051 {
		goto L609
	} else {
		goto L612
	}
L612:
	;
	v2059 = v47 + int32(56)
	goto L613
L613:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v2059)))
	if v2072 != 0 {
		goto L615
	} else {
		goto L616
	}
L614:
	;
	F_pfree(m, v2046)
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L15
	} else {
		goto L620
	}
L615:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+40))
	if v2046 == v2073 {
		goto L609
	} else {
		goto L618
	}
L616:
	;
	goto L617
L617:
	;
	goto L614
L618:
	;
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v2072)+56))
	if v2046 != v2075 {
		v2059 = v2072
		goto L613
	} else {
		goto L619
	}
L619:
	;
	goto L609
L620:
	;
	goto L609
L621:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v2096 | int32(2)
	v2104 = int32(0)
	v2106 = F_errstart(m, v36, v2104)
	mBase = m.M
	v2107 = m.ExcPending
	if v2107 != 0 {
		goto L15
	} else {
		goto L624
	}
L622:
	;
	goto L623
L623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v2096 & int32(-3)
	v2412 = int32(-1)
	goto L11
L624:
	;
	if v2106 == int32(0) {
		v2412 = v2104
		goto L11
	} else {
		goto L625
	}
L625:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v2112 = m.ExcPending
	if v2112 != 0 {
		goto L15
	} else {
		goto L626
	}
L626:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v2113
	F_errmsg(m, int32(_a_F_set_config_with_handle_5), v20+int32(80))
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L15
	} else {
		goto L627
	}
L627:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(_a_F_set_config_with_handle_18), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v2124 = m.ExcPending
	if v2124 != 0 {
		goto L15
	} else {
		goto L628
	}
L628:
	;
	v2412 = v2104
	goto L11
L629:
	;
	if v381 == int32(0) {
		goto L632
	} else {
		goto L633
	}
L630:
	;
	goto L631
L631:
	;
	if v381 == int32(0) {
		goto L652
	} else {
		goto L653
	}
L632:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L15
	} else {
		goto L635
	}
L633:
	;
	goto L634
L634:
	;
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	if v2134 != 0 {
		goto L636
	} else {
		goto L637
	}
L635:
	;
	goto L634
L636:
	;
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	m.T0[v2134].(func(*base.Module, int32, int32))(m, v2133, v2135)
	mBase = m.M
	v2137 = m.ExcPending
	if v2137 != 0 {
		goto L15
	} else {
		goto L639
	}
L637:
	;
	v2139 = v2133
	goto L638
L638:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v2140))) = v2139
	v2144 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(60), v2144)
	mBase = m.M
	v2146 = m.ExcPending
	if v2146 != 0 {
		goto L15
	} else {
		goto L640
	}
L639:
	;
	v2138 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v2139 = v2138
	goto L638
L640:
	;
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v2149 == int32(0) {
		goto L643
	} else {
		goto L644
	}
L641:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v2044
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v2042
	goto L631
L642:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v2043
	goto L641
L643:
	;
	if v2043 == int32(0) {
		goto L642
	} else {
		goto L646
	}
L644:
	;
	goto L645
L645:
	;
	if v2043 != 0 {
		goto L642
	} else {
		goto L651
	}
L646:
	;
	v2155 = v47 + int32(68)
	v2157 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4]))
	if v2157 != 0 {
		goto L648
	} else {
		goto L649
	}
L647:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v2164
	v2166 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+72)) = v2166
	*(*int32)(unsafe.Add(mBase, uint32(v2164)+4)) = v2155
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[5])) = v2155
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v2043
	goto L641
L648:
	;
	v2159 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[5]))
	v2164 = v2159
	goto L647
L649:
	;
	goto L650
L650:
	;
	v2161 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4])) = v2161
	v2164 = v2161
	goto L647
L651:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v47)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v2172)+4)) = v2173
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v2173))) = v2175
	goto L642
L652:
	;
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v2313 == int32(0) {
		goto L677
	} else {
		goto L678
	}
L653:
	;
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v2187) <= base.Ui32(v2043) {
		goto L654
	} else {
		goto L655
	}
L654:
	;
	v2189 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+120)) = v2189
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47-int32(-64), v2193)
	mBase = m.M
	v2195 = m.ExcPending
	if v2195 != 0 {
		goto L15
	} else {
		goto L657
	}
L655:
	;
	goto L656
L656:
	;
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v2199 == int32(0) {
		goto L652
	} else {
		goto L658
	}
L657:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v2044
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v2042
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v2043
	goto L656
L658:
	;
	v2212 = v2199
	goto L659
L659:
	;
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2212)+12))
	if base.Ui32(v2221) <= base.Ui32(v2043) {
		goto L661
	} else {
		goto L662
	}
L660:
	;
	goto L652
L661:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v2212)+32)) = v2223
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v2212)+40))
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v2212)+40)) = v2226
	if v2225 == int32(0) {
		goto L664
	} else {
		goto L665
	}
L662:
	;
	goto L663
L663:
	;
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(v2212)))
	if v2295 != 0 {
		v2212 = v2295
		goto L659
	} else {
		goto L676
	}
L664:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2212)+24)) = v2044
	*(*int32)(unsafe.Add(mBase, uint32(v2212)+16)) = v2042
	*(*int32)(unsafe.Add(mBase, uint32(v2212)+12)) = v2043
	goto L663
L665:
	;
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v2225 == v2230 {
		goto L664
	} else {
		goto L666
	}
L666:
	;
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v2225 == v2232 {
		goto L664
	} else {
		goto L667
	}
L667:
	;
	v2245 = v47 + int32(56)
	goto L668
L668:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v2245)))
	if v2251 != 0 {
		goto L670
	} else {
		goto L671
	}
L669:
	;
	F_pfree(m, v2225)
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L15
	} else {
		goto L675
	}
L670:
	;
	v2252 = *(*int32)(unsafe.Add(mBase, uint32(v2251)+40))
	if v2225 == v2252 {
		goto L664
	} else {
		goto L673
	}
L671:
	;
	goto L672
L672:
	;
	goto L669
L673:
	;
	v2254 = *(*int32)(unsafe.Add(mBase, uint32(v2251)+56))
	if v2225 != v2254 {
		v2245 = v2251
		goto L668
	} else {
		goto L674
	}
L674:
	;
	goto L664
L675:
	;
	goto L664
L676:
	;
	goto L660
L677:
	;
	if v407 != 0 {
		goto L131
	} else {
		goto L689
	}
L678:
	;
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v2313 == v2316 {
		goto L677
	} else {
		goto L679
	}
L679:
	;
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	if v2313 == v2318 {
		goto L677
	} else {
		goto L680
	}
L680:
	;
	v2330 = v47 + int32(56)
	goto L681
L681:
	;
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(v2330)))
	if v2339 != 0 {
		goto L683
	} else {
		goto L684
	}
L682:
	;
	F_pfree(m, v2313)
	mBase = m.M
	v2345 = m.ExcPending
	if v2345 != 0 {
		goto L15
	} else {
		goto L688
	}
L683:
	;
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v2339)+40))
	if v2313 == v2340 {
		goto L677
	} else {
		goto L686
	}
L684:
	;
	goto L685
L685:
	;
	goto L682
L686:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v2339)+56))
	if v2313 != v2342 {
		v2330 = v2339
		goto L681
	} else {
		goto L687
	}
L687:
	;
	goto L677
L688:
	;
	goto L677
L689:
	;
	v2412 = int32(-1)
	goto L11
L690:
	;
	v2412 = int32(-1)
	goto L11
L691:
	;
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	if v2388&int32(4) != 0 {
		v2412 = v2382
		goto L11
	} else {
		goto L692
	}
L692:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v2388 | int32(4)
	v2394 = int32(_a_F_set_config_with_handle_19)
	v2395 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+80)) = v2395
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[7])) = v47 + int32(80)
	v2412 = v2382
	goto L11
}
