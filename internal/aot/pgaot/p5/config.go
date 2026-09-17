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
				F_errmsg(m, int32(_a_F_set_config_by_name_0), int32(0))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_set_config_by_name_1), int32(342), int32(_a_F_set_config_by_name_2))
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
	var v65 int32
	_ = v65
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
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
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v591 int32
	_ = v591
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v683 int32
	_ = v683
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v923 int32
	_ = v923
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v965 int32
	_ = v965
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v1015 int32
	_ = v1015
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1080 float64
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 float64
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 float64
	_ = v1141
	var v1142 float64
	_ = v1142
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1162 int32
	_ = v1162
	var v1167 int32
	_ = v1167
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 float64
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 float64
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 float64
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1232 int32
	_ = v1232
	var v1234 float64
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1257 int32
	_ = v1257
	var v1266 int32
	_ = v1266
	var v1268 float64
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1299 int32
	_ = v1299
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1349 int32
	_ = v1349
	var v1367 int32
	_ = v1367
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1387 int32
	_ = v1387
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1573 int32
	_ = v1573
	var v1578 int32
	_ = v1578
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1685 int32
	_ = v1685
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1700 int32
	_ = v1700
	var v1704 int32
	_ = v1704
	var v1713 int32
	_ = v1713
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1740 int32
	_ = v1740
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1759 int32
	_ = v1759
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1793 int32
	_ = v1793
	var v1799 int32
	_ = v1799
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1856 int32
	_ = v1856
	var v1874 int32
	_ = v1874
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1884 int32
	_ = v1884
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2019 int32
	_ = v2019
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2034 int32
	_ = v2034
	var v2039 int32
	_ = v2039
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2064 int32
	_ = v2064
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2127 int32
	_ = v2127
	var v2136 int32
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2156 int32
	_ = v2156
	var v2169 int32
	_ = v2169
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2178 int32
	_ = v2178
	var v2181 int32
	_ = v2181
	var v2219 int32
	_ = v2219
	var v2237 int32
	_ = v2237
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2247 int32
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2257 int32
	_ = v2257
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2265 int32
	_ = v2265
	var v2267 int32
	_ = v2267
	var v2271 int32
	_ = v2271
	var v2273 int32
	_ = v2273
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2299 int32
	_ = v2299
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2323 int32
	_ = v2323
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
	return v2323
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
		v2323 = v11
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
	v86 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	switch v87 {
	case 0:
		goto L39
	case 1:
		goto L38
	case 2:
		goto L37
	case 3:
		goto L36
	case 4:
		goto L35
	case 5:
		goto L32
	default:
		v259 = v86
		goto L30
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
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+21)))
	if v65&int32(128) != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v69 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	if v69 == int32(0) {
		v2323 = v11
		goto L11
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v76
	F_errmsg(m, int32(_a_F_set_config_with_handle_0), v20)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3465), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v2323 = v11
	goto L11
L30:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	if v260&int32(_a_F_set_config_with_handle_3) != 0 {
		goto L90
	} else {
		goto L91
	}
L31:
	;
	v259 = int32(0)
	goto L30
L32:
	;
	if l3&int32(-3) != int32(4) {
		goto L31
	} else {
		goto L79
	}
L33:
	;
	if l4 == int32(9) {
		goto L31
	} else {
		goto L69
	}
L34:
	;
	if l9|base.B2i32(l7 == int32(0)) != 0 {
		v259 = v86
		goto L30
	} else {
		goto L67
	}
L35:
	;
	if l3 != int32(2) {
		goto L33
	} else {
		goto L66
	}
L36:
	;
	switch l3 - int32(2) {
	case 0:
		goto L34
	default:
		goto L33
	case 2:
		goto L58
	}
L37:
	;
	if base.Ui32(int32(-3)) < base.Ui32(l3-int32(3)) {
		v259 = v86
		goto L30
	} else {
		goto L52
	}
L38:
	;
	v110 = int32(1)
	switch l3 - v110 {
	case 0:
		goto L31
	case 1:
		v259 = v110
		goto L30
	default:
		goto L46
	}
L39:
	;
	if l3 == int32(0) {
		v259 = v86
		goto L30
	} else {
		goto L40
	}
L40:
	;
	v91 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L15
	} else {
		goto L41
	}
L41:
	;
	if v91 == int32(0) {
		v2323 = v11
		goto L11
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L15
	} else {
		goto L43
	}
L43:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v98
	F_errmsg(m, int32(_a_F_set_config_with_handle_4), v20+int32(176))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3481), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	v2323 = v11
	goto L11
L46:
	;
	v114 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L15
	} else {
		goto L47
	}
L47:
	;
	if v114 == int32(0) {
		v2323 = v11
		goto L11
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v121
	F_errmsg(m, int32(_a_F_set_config_with_handle_5), v20+int32(192))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L15
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3504), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L15
	} else {
		goto L51
	}
L51:
	;
	v2323 = v11
	goto L11
L52:
	;
	v138 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L15
	} else {
		goto L53
	}
L53:
	;
	if v138 == int32(0) {
		v2323 = v11
		goto L11
	} else {
		goto L54
	}
L54:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L15
	} else {
		goto L55
	}
L55:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v145
	F_errmsg(m, int32(_a_F_set_config_with_handle_6), v20+int32(208))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L15
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3514), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L15
	} else {
		goto L57
	}
L57:
	;
	v2323 = v11
	goto L11
L58:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v161 = F_pg_parameter_aclcheck(m, v159, l5, int64(4096))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L15
	} else {
		goto L59
	}
L59:
	;
	if v161 == int32(0) {
		goto L33
	} else {
		goto L60
	}
L60:
	;
	v166 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L15
	} else {
		goto L61
	}
L61:
	;
	if v166 == int32(0) {
		v2323 = v11
		goto L11
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L15
	} else {
		goto L63
	}
L63:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v173
	F_errmsg(m, int32(_a_F_set_config_with_handle_7), v20+int32(240))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L15
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3541), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	v2323 = v11
	goto L11
L66:
	;
	goto L34
L67:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_set_config_with_handle[0])))
	if v191&int32(1) == int32(0) {
		v259 = v86
		goto L30
	} else {
		goto L68
	}
L68:
	;
	v2323 = int32(-1)
	goto L11
L69:
	;
	if int32(1)<<(uint(l3)%32)&int32(26) != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v206 = base.B2i32(base.Ui32(l3) <= base.Ui32(int32(4)))
	goto L72
L71:
	;
	v206 = int32(0)
	goto L72
L72:
	;
	if v206 != 0 {
		v259 = v86
		goto L30
	} else {
		goto L73
	}
L73:
	;
	v208 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L15
	} else {
		goto L74
	}
L74:
	;
	if v208 == int32(0) {
		v2323 = v11
		goto L11
	} else {
		goto L75
	}
L75:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L15
	} else {
		goto L76
	}
L76:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = v215
	F_errmsg(m, int32(_a_F_set_config_with_handle_8), v20+int32(224))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L15
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3583), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L15
	} else {
		goto L78
	}
L78:
	;
	v2323 = v11
	goto L11
L79:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v233 = F_pg_parameter_aclcheck(m, v231, l5, int64(4096))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L15
	} else {
		goto L80
	}
L80:
	;
	if v233 == int32(0) {
		v259 = v86
		goto L30
	} else {
		goto L81
	}
L81:
	;
	v238 = F_errstart(m, v36, int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L15
	} else {
		goto L82
	}
L82:
	;
	if v238 == int32(0) {
		v2323 = v11
		goto L11
	} else {
		goto L83
	}
L83:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L15
	} else {
		goto L84
	}
L84:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v245
	F_errmsg(m, int32(_a_F_set_config_with_handle_7), v20+int32(256))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L15
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3603), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L15
	} else {
		goto L86
	}
L86:
	;
	v2323 = v11
	goto L11
L87:
	;
	v365 = int32(0)
	v377 = base.B2i32(base.B2i32(l7 == v365)|base.B2i32(base.Ui32(int32(10)) < base.Ui32(l4)) == v365) & (base.B2i32(l4 == v365) | base.B2i32(l2 != v365))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if base.Ui32(v378) <= base.Ui32(l4) {
		v403 = l7
		goto L123
	} else {
		goto L124
	}
L88:
	;
	if l2 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L89:
	;
	v298 = int32(0)
	v300 = F_errstart(m, v36, v298)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L15
	} else {
		goto L104
	}
L90:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[2]))
	if v264&int32(1) != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v295 = v260
	goto L92
L92:
	;
	if v295&int32(8) != 0 {
		goto L88
	} else {
		goto L103
	}
L93:
	;
	v267 = int32(0)
	v269 = F_errstart(m, v36, v267)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L15
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_set_config_with_handle[2])))
	goto L101
L96:
	;
	if v269 == int32(0) {
		v2323 = v267
		goto L11
	} else {
		goto L97
	}
L97:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L15
	} else {
		goto L98
	}
L98:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v276
	F_errmsg(m, int32(_a_F_set_config_with_handle_9), v20+int32(144))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L15
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3642), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L15
	} else {
		goto L100
	}
L100:
	;
	v2323 = v267
	goto L11
L101:
	;
	if int32(base.Ui32(v289&int32(2))>>(uint(int32(1))%32)) != 0 {
		goto L89
	} else {
		goto L102
	}
L102:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v295 = v294
	goto L92
L103:
	;
	goto L87
L104:
	;
	if v300 == int32(0) {
		v2323 = v298
		goto L11
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L15
	} else {
		goto L106
	}
L106:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v307
	F_errmsg(m, int32(_a_F_set_config_with_handle_10), v20+int32(160))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L15
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3650), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L15
	} else {
		goto L108
	}
L108:
	;
	v2323 = v298
	goto L11
L109:
	;
	v321 = int32(0)
	v323 = F_errstart(m, v36, v321)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L15
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	if l6 != int32(2) {
		goto L87
	} else {
		goto L117
	}
L112:
	;
	if v323 == int32(0) {
		v2323 = v321
		goto L11
	} else {
		goto L113
	}
L113:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L15
	} else {
		goto L114
	}
L114:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v330
	F_errmsg(m, int32(_a_F_set_config_with_handle_11), v20+int32(112))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L15
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3662), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L15
	} else {
		goto L116
	}
L116:
	;
	v2323 = v321
	goto L11
L117:
	;
	v344 = int32(0)
	v346 = F_errstart(m, v36, v344)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L15
	} else {
		goto L118
	}
L118:
	;
	if v346 == int32(0) {
		v2323 = v344
		goto L11
	} else {
		goto L119
	}
L119:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L15
	} else {
		goto L120
	}
L120:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v353
	F_errmsg(m, int32(_a_F_set_config_with_handle_12), v20+int32(128))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L15
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3670), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L15
	} else {
		goto L122
	}
L122:
	;
	v2323 = v344
	goto L11
L123:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v404 {
	case 0:
		goto L136
	case 1:
		goto L135
	case 2:
		goto L134
	case 3:
		goto L133
	case 4:
		goto L132
	default:
		goto L131
	}
L124:
	;
	if l7^int32(1)|v377 != 0 {
		v403 = int32(0)
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v384 = int32(-1)
	v387 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L15
	} else {
		goto L126
	}
L126:
	;
	if v387 == int32(0) {
		v2323 = v384
		goto L11
	} else {
		goto L127
	}
L127:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v391
	F_errmsg_internal(m, int32(_a_F_set_config_with_handle_13), v20+int32(96))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L15
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3696), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L15
	} else {
		goto L129
	}
L129:
	;
	v2323 = v384
	goto L11
L130:
	;
	v2293 = int32(1)
	v2294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+20)))
	if v2294&int32(64) == int32(0) {
		v2323 = v2293
		goto L11
	} else {
		goto L850
	}
L131:
	;
	if v403 != 0 {
		goto L130
	} else {
		goto L849
	}
L132:
	;
	if l2 != 0 {
		goto L723
	} else {
		goto L724
	}
L133:
	;
	if l2 != 0 {
		goto L519
	} else {
		goto L520
	}
L134:
	;
	if l2 != 0 {
		goto L392
	} else {
		goto L393
	}
L135:
	;
	if l2 != 0 {
		goto L265
	} else {
		goto L266
	}
L136:
	;
	if l2 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v259 != 0 {
		goto L148
	} else {
		goto L149
	}
L138:
	;
	v409 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L15
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	if l4 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	if v409 != 0 {
		v432 = l3
		v433 = l4
		v434 = l5
		goto L137
	} else {
		goto L142
	}
L142:
	;
	v2323 = int32(0)
	goto L11
L143:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)) = uint8(v414)
	v416 = int32(0)
	v423 = F_call_bool_check_hook(m, v47, v20+int32(264), v20+int32(260), v416, v36)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L15
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+112)))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)) = uint8(v425)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v427
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v432 = v430
	v433 = v431
	v434 = v429
	goto L137
L146:
	;
	if v423 != 0 {
		v432 = l3
		v433 = v416
		v434 = l5
		goto L137
	} else {
		goto L147
	}
L147:
	;
	v2323 = v416
	goto L11
L148:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v436 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	if v403 != 0 {
		goto L182
	} else {
		goto L183
	}
L151:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	if v475 != v476 {
		goto L174
	} else {
		goto L175
	}
L152:
	;
	v440 = int32(1)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v436 == v441 {
		v470 = v440
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if v470 != 0 {
		goto L151
	} else {
		goto L172
	}
L154:
	;
	goto L153
L155:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v443 {
	case 0:
		goto L161
	case 1:
		goto L160
	case 2:
		goto L159
	case 3:
		goto L158
	case 4:
		goto L157
	default:
		goto L156
	}
L156:
	;
	v456 = v47 + int32(56)
	goto L167
L157:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v436 == v452 {
		v470 = v440
		goto L154
	} else {
		goto L166
	}
L158:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v436 != v450 {
		goto L156
	} else {
		goto L165
	}
L159:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v436 != v448 {
		goto L156
	} else {
		goto L164
	}
L160:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v436 != v446 {
		goto L156
	} else {
		goto L163
	}
L161:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v436 != v444 {
		goto L156
	} else {
		goto L162
	}
L162:
	;
	v470 = v440
	goto L154
L163:
	;
	v470 = v440
	goto L154
L164:
	;
	v470 = v440
	goto L154
L165:
	;
	v470 = v440
	goto L154
L166:
	;
	goto L156
L167:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	v460 = int32(0)
	v461 = base.B2i32(v459 != v460)
	if v459 == v460 {
		v470 = v461
		goto L154
	} else {
		goto L169
	}
L168:
	;
	v470 = v461
	goto L154
L169:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v459)+40))
	if v436 == v464 {
		v470 = v461
		goto L154
	} else {
		goto L170
	}
L170:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v459)+56))
	if v436 != v466 {
		v456 = v459
		goto L167
	} else {
		goto L171
	}
L171:
	;
	goto L168
L172:
	;
	F_pfree(m, v436)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L15
	} else {
		goto L173
	}
L173:
	;
	goto L151
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v473 | int32(2)
	v481 = int32(0)
	v483 = F_errstart(m, v36, v481)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L15
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v473 & int32(-3)
	v2323 = int32(-1)
	goto L11
L177:
	;
	if v483 == int32(0) {
		v2323 = v481
		goto L11
	} else {
		goto L178
	}
L178:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L15
	} else {
		goto L179
	}
L179:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v490
	F_errmsg(m, int32(_a_F_set_config_with_handle_5), v20+int32(16))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L15
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3748), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L15
	} else {
		goto L181
	}
L181:
	;
	v2323 = v481
	goto L11
L182:
	;
	if v377 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	goto L184
L184:
	;
	if v377 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L185:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L15
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v47)+104))
	if v511 != 0 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	goto L187
L189:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	m.T0[v511].(func(*base.Module, int32, int32))(m, v510&int32(1), v514)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L15
	} else {
		goto L192
	}
L190:
	;
	v518 = v510
	goto L191
L191:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	*(*uint8)(unsafe.Add(mBase, uint32(v519))) = uint8(v518)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(60), v523)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L15
	} else {
		goto L193
	}
L192:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	v518 = v517
	goto L191
L193:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v528 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v432
	goto L184
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v433
	goto L194
L196:
	;
	if v433 == int32(0) {
		goto L195
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	if v433 != 0 {
		goto L195
	} else {
		goto L204
	}
L199:
	;
	v534 = v47 - int32(-64)
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[3]))
	if v536 != 0 {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+64)) = v543
	v545 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v543)+4)) = v534
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4])) = v534
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v433
	goto L194
L201:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4]))
	v543 = v538
	goto L200
L202:
	;
	goto L203
L203:
	;
	v540 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[3])) = v540
	v543 = v540
	goto L200
L204:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v551)+4)) = v552
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v552))) = v554
	goto L195
L205:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v701 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L206:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v566) <= base.Ui32(v433) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+112)) = uint8(v568)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(116), v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L15
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v578 == int32(0) {
		goto L205
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v433
	goto L209
L211:
	;
	v591 = v578
	goto L212
L212:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v591)+12))
	if base.Ui32(v600) <= base.Ui32(v433) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	goto L205
L214:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+264)))
	*(*uint8)(unsafe.Add(mBase, uint32(v591)+32)) = uint8(v602)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v591)+40))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v591)+40)) = v605
	if v604 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	goto L216
L216:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v591)))
	if v683 != 0 {
		v591 = v683
		goto L212
	} else {
		goto L239
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v591)+24)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v591)+16)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v591)+12)) = v433
	goto L216
L218:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v604 == v609 {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v611 {
	case 0:
		goto L225
	case 1:
		goto L224
	case 2:
		goto L223
	case 3:
		goto L222
	case 4:
		goto L221
	default:
		goto L220
	}
L220:
	;
	v633 = v47 + int32(56)
	goto L231
L221:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v604 == v620 {
		goto L217
	} else {
		goto L230
	}
L222:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v604 != v618 {
		goto L220
	} else {
		goto L229
	}
L223:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v604 != v616 {
		goto L220
	} else {
		goto L228
	}
L224:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v604 != v614 {
		goto L220
	} else {
		goto L227
	}
L225:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v604 != v612 {
		goto L220
	} else {
		goto L226
	}
L226:
	;
	goto L217
L227:
	;
	goto L217
L228:
	;
	goto L217
L229:
	;
	goto L217
L230:
	;
	goto L220
L231:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	if v639 != 0 {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	F_pfree(m, v604)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L15
	} else {
		goto L238
	}
L233:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v639)+40))
	if v604 == v640 {
		goto L217
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	goto L232
L236:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v639)+56))
	if v604 != v642 {
		v633 = v639
		goto L231
	} else {
		goto L237
	}
L237:
	;
	goto L217
L238:
	;
	goto L217
L239:
	;
	goto L213
L240:
	;
	if v403 != 0 {
		goto L130
	} else {
		goto L263
	}
L241:
	;
	v705 = int32(1)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v701 == v706 {
		v735 = v705
		goto L243
	} else {
		goto L244
	}
L242:
	;
	if v735 != 0 {
		goto L240
	} else {
		goto L261
	}
L243:
	;
	goto L242
L244:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v708 {
	case 0:
		goto L250
	case 1:
		goto L249
	case 2:
		goto L248
	case 3:
		goto L247
	case 4:
		goto L246
	default:
		goto L245
	}
L245:
	;
	v721 = v47 + int32(56)
	goto L256
L246:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v701 == v717 {
		v735 = v705
		goto L243
	} else {
		goto L255
	}
L247:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v701 != v715 {
		goto L245
	} else {
		goto L254
	}
L248:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v701 != v713 {
		goto L245
	} else {
		goto L253
	}
L249:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v701 != v711 {
		goto L245
	} else {
		goto L252
	}
L250:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v701 != v709 {
		goto L245
	} else {
		goto L251
	}
L251:
	;
	v735 = v705
	goto L243
L252:
	;
	v735 = v705
	goto L243
L253:
	;
	v735 = v705
	goto L243
L254:
	;
	v735 = v705
	goto L243
L255:
	;
	goto L245
L256:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	v725 = int32(0)
	v726 = base.B2i32(v724 != v725)
	if v724 == v725 {
		v735 = v726
		goto L243
	} else {
		goto L258
	}
L257:
	;
	v735 = v726
	goto L243
L258:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v724)+40))
	if v701 == v729 {
		v735 = v726
		goto L243
	} else {
		goto L259
	}
L259:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v724)+56))
	if v701 != v731 {
		v721 = v724
		goto L256
	} else {
		goto L260
	}
L260:
	;
	goto L257
L261:
	;
	F_pfree(m, v701)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L15
	} else {
		goto L262
	}
L262:
	;
	goto L240
L263:
	;
	v2323 = int32(-1)
	goto L11
L264:
	;
	if v259 != 0 {
		goto L275
	} else {
		goto L276
	}
L265:
	;
	v743 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L15
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	if l4 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	if v743 != 0 {
		v766 = l3
		v767 = l4
		v768 = l5
		goto L264
	} else {
		goto L269
	}
L269:
	;
	v2323 = int32(0)
	goto L11
L270:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v748
	v750 = int32(0)
	v757 = F_call_int_check_hook(m, v47, v20+int32(264), v20+int32(260), v750, v36)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L15
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v759
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v761
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v766 = v764
	v767 = v765
	v768 = v763
	goto L264
L273:
	;
	if v757 != 0 {
		v766 = l3
		v767 = v750
		v768 = l5
		goto L264
	} else {
		goto L274
	}
L274:
	;
	v2323 = v750
	goto L11
L275:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v770 == int32(0) {
		goto L278
	} else {
		goto L279
	}
L276:
	;
	goto L277
L277:
	;
	if v403 != 0 {
		goto L309
	} else {
		goto L310
	}
L278:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v808)))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if v809 != v810 {
		goto L301
	} else {
		goto L302
	}
L279:
	;
	v774 = int32(1)
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v770 == v775 {
		v804 = v774
		goto L281
	} else {
		goto L282
	}
L280:
	;
	if v804 != 0 {
		goto L278
	} else {
		goto L299
	}
L281:
	;
	goto L280
L282:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v777 {
	case 0:
		goto L288
	case 1:
		goto L287
	case 2:
		goto L286
	case 3:
		goto L285
	case 4:
		goto L284
	default:
		goto L283
	}
L283:
	;
	v790 = v47 + int32(56)
	goto L294
L284:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v770 == v786 {
		v804 = v774
		goto L281
	} else {
		goto L293
	}
L285:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v770 != v784 {
		goto L283
	} else {
		goto L292
	}
L286:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v770 != v782 {
		goto L283
	} else {
		goto L291
	}
L287:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v770 != v780 {
		goto L283
	} else {
		goto L290
	}
L288:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v770 != v778 {
		goto L283
	} else {
		goto L289
	}
L289:
	;
	v804 = v774
	goto L281
L290:
	;
	v804 = v774
	goto L281
L291:
	;
	v804 = v774
	goto L281
L292:
	;
	v804 = v774
	goto L281
L293:
	;
	goto L283
L294:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v790)))
	v794 = int32(0)
	v795 = base.B2i32(v793 != v794)
	if v793 == v794 {
		v804 = v795
		goto L281
	} else {
		goto L296
	}
L295:
	;
	v804 = v795
	goto L281
L296:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v793)+40))
	if v770 == v798 {
		v804 = v795
		goto L281
	} else {
		goto L297
	}
L297:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v793)+56))
	if v770 != v800 {
		v790 = v793
		goto L294
	} else {
		goto L298
	}
L298:
	;
	goto L295
L299:
	;
	F_pfree(m, v770)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L15
	} else {
		goto L300
	}
L300:
	;
	goto L278
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v807 | int32(2)
	v815 = int32(0)
	v817 = F_errstart(m, v36, v815)
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L15
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v807 & int32(-3)
	v2323 = int32(-1)
	goto L11
L304:
	;
	if v817 == int32(0) {
		v2323 = v815
		goto L11
	} else {
		goto L305
	}
L305:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L15
	} else {
		goto L306
	}
L306:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v824
	F_errmsg(m, int32(_a_F_set_config_with_handle_5), v20+int32(32))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L15
	} else {
		goto L307
	}
L307:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3846), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L15
	} else {
		goto L308
	}
L308:
	;
	v2323 = v815
	goto L11
L309:
	;
	if v377 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L310:
	;
	goto L311
L311:
	;
	if v377 == int32(0) {
		goto L332
	} else {
		goto L333
	}
L312:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L15
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	if v845 != 0 {
		goto L316
	} else {
		goto L317
	}
L315:
	;
	goto L314
L316:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	m.T0[v845].(func(*base.Module, int32, int32))(m, v844, v846)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L15
	} else {
		goto L319
	}
L317:
	;
	v850 = v844
	goto L318
L318:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v851))) = v850
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(60), v855)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L15
	} else {
		goto L320
	}
L319:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v850 = v849
	goto L318
L320:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v860 == int32(0) {
		goto L323
	} else {
		goto L324
	}
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v766
	goto L311
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v767
	goto L321
L323:
	;
	if v767 == int32(0) {
		goto L322
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	if v767 != 0 {
		goto L322
	} else {
		goto L331
	}
L326:
	;
	v866 = v47 - int32(-64)
	v868 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[3]))
	if v868 != 0 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+64)) = v875
	v877 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v877
	*(*int32)(unsafe.Add(mBase, uint32(v875)+4)) = v866
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4])) = v866
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v767
	goto L321
L328:
	;
	v870 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4]))
	v875 = v870
	goto L327
L329:
	;
	goto L330
L330:
	;
	v872 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[3])) = v872
	v875 = v872
	goto L327
L331:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v883)+4)) = v884
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v884))) = v886
	goto L322
L332:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1033 == int32(0) {
		goto L367
	} else {
		goto L368
	}
L333:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v898) <= base.Ui32(v767) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+120)) = v900
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(124), v904)
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L15
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v910 == int32(0) {
		goto L332
	} else {
		goto L338
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v766
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v767
	goto L336
L338:
	;
	v923 = v910
	goto L339
L339:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v923)+12))
	if base.Ui32(v932) <= base.Ui32(v767) {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	goto L332
L341:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v923)+32)) = v934
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v923)+40))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v923)+40)) = v937
	if v936 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L342:
	;
	goto L343
L343:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v923)))
	if v1015 != 0 {
		v923 = v1015
		goto L339
	} else {
		goto L366
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v923)+24)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v923)+16)) = v766
	*(*int32)(unsafe.Add(mBase, uint32(v923)+12)) = v767
	goto L343
L345:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v936 == v941 {
		goto L344
	} else {
		goto L346
	}
L346:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v943 {
	case 0:
		goto L352
	case 1:
		goto L351
	case 2:
		goto L350
	case 3:
		goto L349
	case 4:
		goto L348
	default:
		goto L347
	}
L347:
	;
	v965 = v47 + int32(56)
	goto L358
L348:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v936 == v952 {
		goto L344
	} else {
		goto L357
	}
L349:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v936 != v950 {
		goto L347
	} else {
		goto L356
	}
L350:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v936 != v948 {
		goto L347
	} else {
		goto L355
	}
L351:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v936 != v946 {
		goto L347
	} else {
		goto L354
	}
L352:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v936 != v944 {
		goto L347
	} else {
		goto L353
	}
L353:
	;
	goto L344
L354:
	;
	goto L344
L355:
	;
	goto L344
L356:
	;
	goto L344
L357:
	;
	goto L347
L358:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v965)))
	if v971 != 0 {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	F_pfree(m, v936)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L15
	} else {
		goto L365
	}
L360:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v971)+40))
	if v936 == v972 {
		goto L344
	} else {
		goto L363
	}
L361:
	;
	goto L362
L362:
	;
	goto L359
L363:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v971)+56))
	if v936 != v974 {
		v965 = v971
		goto L358
	} else {
		goto L364
	}
L364:
	;
	goto L344
L365:
	;
	goto L344
L366:
	;
	goto L340
L367:
	;
	if v403 != 0 {
		goto L130
	} else {
		goto L390
	}
L368:
	;
	v1037 = int32(1)
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1033 == v1038 {
		v1067 = v1037
		goto L370
	} else {
		goto L371
	}
L369:
	;
	if v1067 != 0 {
		goto L367
	} else {
		goto L388
	}
L370:
	;
	goto L369
L371:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1040 {
	case 0:
		goto L377
	case 1:
		goto L376
	case 2:
		goto L375
	case 3:
		goto L374
	case 4:
		goto L373
	default:
		goto L372
	}
L372:
	;
	v1053 = v47 + int32(56)
	goto L383
L373:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1033 == v1049 {
		v1067 = v1037
		goto L370
	} else {
		goto L382
	}
L374:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1033 != v1047 {
		goto L372
	} else {
		goto L381
	}
L375:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1033 != v1045 {
		goto L372
	} else {
		goto L380
	}
L376:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1033 != v1043 {
		goto L372
	} else {
		goto L379
	}
L377:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1033 != v1041 {
		goto L372
	} else {
		goto L378
	}
L378:
	;
	v1067 = v1037
	goto L370
L379:
	;
	v1067 = v1037
	goto L370
L380:
	;
	v1067 = v1037
	goto L370
L381:
	;
	v1067 = v1037
	goto L370
L382:
	;
	goto L372
L383:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1053)))
	v1057 = int32(0)
	v1058 = base.B2i32(v1056 != v1057)
	if v1056 == v1057 {
		v1067 = v1058
		goto L370
	} else {
		goto L385
	}
L384:
	;
	v1067 = v1058
	goto L370
L385:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+40))
	if v1033 == v1061 {
		v1067 = v1058
		goto L370
	} else {
		goto L386
	}
L386:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+56))
	if v1033 != v1063 {
		v1053 = v1056
		goto L383
	} else {
		goto L387
	}
L387:
	;
	goto L384
L388:
	;
	F_pfree(m, v1033)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L15
	} else {
		goto L389
	}
L389:
	;
	goto L367
L390:
	;
	v2323 = int32(-1)
	goto L11
L391:
	;
	if v259 != 0 {
		goto L402
	} else {
		goto L403
	}
L392:
	;
	v1075 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L15
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	if l4 == int32(0) {
		goto L397
	} else {
		goto L398
	}
L395:
	;
	if v1075 != 0 {
		v1098 = l3
		v1099 = l4
		v1100 = l5
		goto L391
	} else {
		goto L396
	}
L396:
	;
	v2323 = int32(0)
	goto L11
L397:
	;
	v1080 = *(*float64)(unsafe.Add(mBase, uint32(v47)+96))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+264)) = v1080
	v1082 = int32(0)
	v1089 = F_call_real_check_hook(m, v47, v20+int32(264), v20+int32(260), v1082, v36)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L15
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	v1091 = *(*float64)(unsafe.Add(mBase, uint32(v47)+136))
	*(*float64)(unsafe.Add(mBase, uint32(v20)+264)) = v1091
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v1093
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v1098 = v1096
	v1099 = v1097
	v1100 = v1095
	goto L391
L400:
	;
	if v1089 != 0 {
		v1098 = l3
		v1099 = v1082
		v1100 = l5
		goto L391
	} else {
		goto L401
	}
L401:
	;
	v2323 = v1082
	goto L11
L402:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1102 == int32(0) {
		goto L405
	} else {
		goto L406
	}
L403:
	;
	goto L404
L404:
	;
	if v403 != 0 {
		goto L436
	} else {
		goto L437
	}
L405:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v1141 = *(*float64)(unsafe.Add(mBase, uint32(v1140)))
	v1142 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	if base.F64_ne(v1141, v1142) != 0 {
		goto L428
	} else {
		goto L429
	}
L406:
	;
	v1106 = int32(1)
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1102 == v1107 {
		v1136 = v1106
		goto L408
	} else {
		goto L409
	}
L407:
	;
	if v1136 != 0 {
		goto L405
	} else {
		goto L426
	}
L408:
	;
	goto L407
L409:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1109 {
	case 0:
		goto L415
	case 1:
		goto L414
	case 2:
		goto L413
	case 3:
		goto L412
	case 4:
		goto L411
	default:
		goto L410
	}
L410:
	;
	v1122 = v47 + int32(56)
	goto L421
L411:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1102 == v1118 {
		v1136 = v1106
		goto L408
	} else {
		goto L420
	}
L412:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1102 != v1116 {
		goto L410
	} else {
		goto L419
	}
L413:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1102 != v1114 {
		goto L410
	} else {
		goto L418
	}
L414:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1102 != v1112 {
		goto L410
	} else {
		goto L417
	}
L415:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1102 != v1110 {
		goto L410
	} else {
		goto L416
	}
L416:
	;
	v1136 = v1106
	goto L408
L417:
	;
	v1136 = v1106
	goto L408
L418:
	;
	v1136 = v1106
	goto L408
L419:
	;
	v1136 = v1106
	goto L408
L420:
	;
	goto L410
L421:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1122)))
	v1126 = int32(0)
	v1127 = base.B2i32(v1125 != v1126)
	if v1125 == v1126 {
		v1136 = v1127
		goto L408
	} else {
		goto L423
	}
L422:
	;
	v1136 = v1127
	goto L408
L423:
	;
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+40))
	if v1102 == v1130 {
		v1136 = v1127
		goto L408
	} else {
		goto L424
	}
L424:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1125)+56))
	if v1102 != v1132 {
		v1122 = v1125
		goto L421
	} else {
		goto L425
	}
L425:
	;
	goto L422
L426:
	;
	F_pfree(m, v1102)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L15
	} else {
		goto L427
	}
L427:
	;
	goto L405
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v1139 | int32(2)
	v1147 = int32(0)
	v1149 = F_errstart(m, v36, v1147)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L15
	} else {
		goto L431
	}
L429:
	;
	goto L430
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v1139 & int32(-3)
	v2323 = int32(-1)
	goto L11
L431:
	;
	if v1149 == int32(0) {
		v2323 = v1147
		goto L11
	} else {
		goto L432
	}
L432:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L15
	} else {
		goto L433
	}
L433:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v1156
	F_errmsg(m, int32(_a_F_set_config_with_handle_5), v20+int32(48))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L15
	} else {
		goto L434
	}
L434:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(3944), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L15
	} else {
		goto L435
	}
L435:
	;
	v2323 = v1147
	goto L11
L436:
	;
	if v377 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L437:
	;
	goto L438
L438:
	;
	if v377 == int32(0) {
		goto L459
	} else {
		goto L460
	}
L439:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L15
	} else {
		goto L442
	}
L440:
	;
	goto L441
L441:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	v1177 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1178 != 0 {
		goto L443
	} else {
		goto L444
	}
L442:
	;
	goto L441
L443:
	;
	m.T0[v1178].(func(*base.Module, float64, int32))(m, v1177, v1176)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L15
	} else {
		goto L446
	}
L444:
	;
	v1183 = v1176
	v1184 = v1177
	goto L445
L445:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	*(*float64)(unsafe.Add(mBase, uint32(v1185))) = v1184
	F_set_extra_field(m, v47, v47+int32(60), v1183)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L15
	} else {
		goto L447
	}
L446:
	;
	v1181 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	v1183 = v1182
	v1184 = v1181
	goto L445
L447:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v1193 == int32(0) {
		goto L450
	} else {
		goto L451
	}
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v1098
	goto L438
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1099
	goto L448
L450:
	;
	if v1099 == int32(0) {
		goto L449
	} else {
		goto L453
	}
L451:
	;
	goto L452
L452:
	;
	if v1099 != 0 {
		goto L449
	} else {
		goto L458
	}
L453:
	;
	v1199 = v47 - int32(-64)
	v1201 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[3]))
	if v1201 != 0 {
		goto L455
	} else {
		goto L456
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+64)) = v1208
	v1210 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v1208)+4)) = v1199
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4])) = v1199
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1099
	goto L448
L455:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4]))
	v1208 = v1203
	goto L454
L456:
	;
	goto L457
L457:
	;
	v1205 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[3])) = v1205
	v1208 = v1205
	goto L454
L458:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1216)+4)) = v1217
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v1217))) = v1219
	goto L449
L459:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1367 == int32(0) {
		goto L494
	} else {
		goto L495
	}
L460:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v1232) <= base.Ui32(v1099) {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v1234 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	*(*float64)(unsafe.Add(mBase, uint32(v47)+136)) = v1234
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(144), v1238)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L15
	} else {
		goto L464
	}
L462:
	;
	goto L463
L463:
	;
	v1244 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v1244 == int32(0) {
		goto L459
	} else {
		goto L465
	}
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v1098
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v1099
	goto L463
L465:
	;
	v1257 = v1244
	goto L466
L466:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+12))
	if base.Ui32(v1266) <= base.Ui32(v1099) {
		goto L468
	} else {
		goto L469
	}
L467:
	;
	goto L459
L468:
	;
	v1268 = *(*float64)(unsafe.Add(mBase, uint32(v20)+264))
	*(*float64)(unsafe.Add(mBase, uint32(v1257)+32)) = v1268
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1257)+40))
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+40)) = v1271
	if v1270 == int32(0) {
		goto L471
	} else {
		goto L472
	}
L469:
	;
	goto L470
L470:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1257)))
	if v1349 != 0 {
		v1257 = v1349
		goto L466
	} else {
		goto L493
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+24)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+16)) = v1098
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+12)) = v1099
	goto L470
L472:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1270 == v1275 {
		goto L471
	} else {
		goto L473
	}
L473:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1277 {
	case 0:
		goto L479
	case 1:
		goto L478
	case 2:
		goto L477
	case 3:
		goto L476
	case 4:
		goto L475
	default:
		goto L474
	}
L474:
	;
	v1299 = v47 + int32(56)
	goto L485
L475:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1270 == v1286 {
		goto L471
	} else {
		goto L484
	}
L476:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1270 != v1284 {
		goto L474
	} else {
		goto L483
	}
L477:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1270 != v1282 {
		goto L474
	} else {
		goto L482
	}
L478:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1270 != v1280 {
		goto L474
	} else {
		goto L481
	}
L479:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1270 != v1278 {
		goto L474
	} else {
		goto L480
	}
L480:
	;
	goto L471
L481:
	;
	goto L471
L482:
	;
	goto L471
L483:
	;
	goto L471
L484:
	;
	goto L474
L485:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1299)))
	if v1305 != 0 {
		goto L487
	} else {
		goto L488
	}
L486:
	;
	F_pfree(m, v1270)
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L15
	} else {
		goto L492
	}
L487:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+40))
	if v1270 == v1306 {
		goto L471
	} else {
		goto L490
	}
L488:
	;
	goto L489
L489:
	;
	goto L486
L490:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+56))
	if v1270 != v1308 {
		v1299 = v1305
		goto L485
	} else {
		goto L491
	}
L491:
	;
	goto L471
L492:
	;
	goto L471
L493:
	;
	goto L467
L494:
	;
	if v403 != 0 {
		goto L130
	} else {
		goto L517
	}
L495:
	;
	v1371 = int32(1)
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1367 == v1372 {
		v1401 = v1371
		goto L497
	} else {
		goto L498
	}
L496:
	;
	if v1401 != 0 {
		goto L494
	} else {
		goto L515
	}
L497:
	;
	goto L496
L498:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1374 {
	case 0:
		goto L504
	case 1:
		goto L503
	case 2:
		goto L502
	case 3:
		goto L501
	case 4:
		goto L500
	default:
		goto L499
	}
L499:
	;
	v1387 = v47 + int32(56)
	goto L510
L500:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1367 == v1383 {
		v1401 = v1371
		goto L497
	} else {
		goto L509
	}
L501:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1367 != v1381 {
		goto L499
	} else {
		goto L508
	}
L502:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1367 != v1379 {
		goto L499
	} else {
		goto L507
	}
L503:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1367 != v1377 {
		goto L499
	} else {
		goto L506
	}
L504:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1367 != v1375 {
		goto L499
	} else {
		goto L505
	}
L505:
	;
	v1401 = v1371
	goto L497
L506:
	;
	v1401 = v1371
	goto L497
L507:
	;
	v1401 = v1371
	goto L497
L508:
	;
	v1401 = v1371
	goto L497
L509:
	;
	goto L499
L510:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	v1391 = int32(0)
	v1392 = base.B2i32(v1390 != v1391)
	if v1390 == v1391 {
		v1401 = v1392
		goto L497
	} else {
		goto L512
	}
L511:
	;
	v1401 = v1392
	goto L497
L512:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1390)+40))
	if v1367 == v1395 {
		v1401 = v1392
		goto L497
	} else {
		goto L513
	}
L513:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1390)+56))
	if v1367 != v1397 {
		v1387 = v1390
		goto L510
	} else {
		goto L514
	}
L514:
	;
	goto L511
L515:
	;
	F_pfree(m, v1367)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L15
	} else {
		goto L516
	}
L516:
	;
	goto L494
L517:
	;
	v2323 = int32(-1)
	goto L11
L518:
	;
	if v259 != 0 {
		goto L537
	} else {
		goto L538
	}
L519:
	;
	v1409 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L15
	} else {
		goto L522
	}
L520:
	;
	goto L521
L521:
	;
	if l4 == int32(0) {
		goto L524
	} else {
		goto L525
	}
L522:
	;
	if v1409 != 0 {
		v1443 = l4
		v1444 = l5
		v1445 = l3
		goto L518
	} else {
		goto L523
	}
L523:
	;
	v2323 = int32(0)
	goto L11
L524:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	if v1414 != 0 {
		goto L528
	} else {
		goto L529
	}
L525:
	;
	goto L526
L526:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v1435
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v1437
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v1443 = v1441
	v1444 = v1439
	v1445 = v1440
	goto L518
L527:
	;
	v1427 = F_call_string_check_hook(m, v47, v20+int32(264), v20+int32(260), int32(0), v36)
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L15
	} else {
		goto L533
	}
L528:
	;
	v1415 = F_guc_strdup(m, v36, v1414)
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L15
	} else {
		goto L531
	}
L529:
	;
	goto L530
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = int32(0)
	goto L527
L531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v1415
	if v1415 != 0 {
		goto L527
	} else {
		goto L532
	}
L532:
	;
	v2323 = int32(0)
	goto L11
L533:
	;
	if v1427 != 0 {
		v1443 = v11
		v1444 = l5
		v1445 = l3
		goto L518
	} else {
		goto L534
	}
L534:
	;
	v1429 = int32(0)
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if v1430 == v1429 {
		v2323 = v1429
		goto L11
	} else {
		goto L535
	}
L535:
	;
	F_pfree(m, v1430)
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L15
	} else {
		goto L536
	}
L536:
	;
	v2323 = v1429
	goto L11
L537:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1446)))
	v1448 = int32(0)
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if base.B2i32(v1447 == v1448)|base.B2i32(v1450 == v1448) == v1448 {
		goto L542
	} else {
		goto L543
	}
L538:
	;
	goto L539
L539:
	;
	if v403 == int32(0) {
		goto L596
	} else {
		goto L597
	}
L540:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1517 == int32(0) {
		goto L565
	} else {
		goto L566
	}
L541:
	;
	v1489 = int32(1)
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1490)))
	if v1450 == v1491 {
		v1513 = v1489
		goto L554
	} else {
		goto L555
	}
L542:
	;
	v1458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1447))))
	v1461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1450))))
	if base.B2i32(v1458 == int32(0))|base.B2i32(v1458 != v1461) != 0 {
		v1479 = v1458
		v1480 = v1461
		goto L546
	} else {
		goto L547
	}
L543:
	;
	goto L544
L544:
	;
	v1484 = int32(1)
	if v1450 == int32(0) {
		v1516 = v1484
		goto L540
	} else {
		goto L552
	}
L545:
	;
	v1487 = base.B2i32(v1479-v1480 != int32(0))
	goto L541
L546:
	;
	goto L545
L547:
	;
	v1464 = v1447
	v1465 = v1450
	goto L548
L548:
	;
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1465)+1)))
	v1469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1464)+1)))
	if v1469 == int32(0) {
		v1479 = v1469
		v1480 = v1468
		goto L546
	} else {
		goto L550
	}
L549:
	;
	v1479 = v1469
	v1480 = v1468
	goto L546
L550:
	;
	v1472 = int32(1)
	if v1469 == v1468 {
		v1464 = v1464 + v1472
		v1465 = v1465 + v1472
		goto L548
	} else {
		goto L551
	}
L551:
	;
	goto L549
L552:
	;
	v1487 = v1484
	goto L541
L553:
	;
	if v1513 != 0 {
		v1516 = v1487
		goto L540
	} else {
		goto L563
	}
L554:
	;
	goto L553
L555:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	if v1450 == v1493 {
		v1513 = v1489
		goto L554
	} else {
		goto L556
	}
L556:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	if v1450 == v1495 {
		v1513 = v1489
		goto L554
	} else {
		goto L557
	}
L557:
	;
	v1499 = v47 + int32(56)
	goto L558
L558:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1499)))
	v1503 = int32(0)
	v1504 = base.B2i32(v1502 != v1503)
	if v1502 == v1503 {
		v1513 = v1504
		goto L554
	} else {
		goto L560
	}
L559:
	;
	v1513 = v1504
	goto L554
L560:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1502)+32))
	if v1450 == v1507 {
		v1513 = v1504
		goto L554
	} else {
		goto L561
	}
L561:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1502)+48))
	if v1450 != v1509 {
		v1499 = v1502
		goto L558
	} else {
		goto L562
	}
L562:
	;
	goto L559
L563:
	;
	F_pfree(m, v1450)
	mBase = m.M
	v1515 = m.ExcPending
	if v1515 != 0 {
		goto L15
	} else {
		goto L564
	}
L564:
	;
	v1516 = v1487
	goto L540
L565:
	;
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	if v1516 != 0 {
		goto L588
	} else {
		goto L589
	}
L566:
	;
	v1521 = int32(1)
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1517 == v1522 {
		v1551 = v1521
		goto L568
	} else {
		goto L569
	}
L567:
	;
	if v1551 != 0 {
		goto L565
	} else {
		goto L586
	}
L568:
	;
	goto L567
L569:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1524 {
	case 0:
		goto L575
	case 1:
		goto L574
	case 2:
		goto L573
	case 3:
		goto L572
	case 4:
		goto L571
	default:
		goto L570
	}
L570:
	;
	v1537 = v47 + int32(56)
	goto L581
L571:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1517 == v1533 {
		v1551 = v1521
		goto L568
	} else {
		goto L580
	}
L572:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1517 != v1531 {
		goto L570
	} else {
		goto L579
	}
L573:
	;
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1517 != v1529 {
		goto L570
	} else {
		goto L578
	}
L574:
	;
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1517 != v1527 {
		goto L570
	} else {
		goto L577
	}
L575:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1517 != v1525 {
		goto L570
	} else {
		goto L576
	}
L576:
	;
	v1551 = v1521
	goto L568
L577:
	;
	v1551 = v1521
	goto L568
L578:
	;
	v1551 = v1521
	goto L568
L579:
	;
	v1551 = v1521
	goto L568
L580:
	;
	goto L570
L581:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v1537)))
	v1541 = int32(0)
	v1542 = base.B2i32(v1540 != v1541)
	if v1540 == v1541 {
		v1551 = v1542
		goto L568
	} else {
		goto L583
	}
L582:
	;
	v1551 = v1542
	goto L568
L583:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1540)+40))
	if v1517 == v1545 {
		v1551 = v1542
		goto L568
	} else {
		goto L584
	}
L584:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1540)+56))
	if v1517 != v1547 {
		v1537 = v1540
		goto L581
	} else {
		goto L585
	}
L585:
	;
	goto L582
L586:
	;
	F_pfree(m, v1517)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L15
	} else {
		goto L587
	}
L587:
	;
	goto L565
L588:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v1554 | int32(2)
	v1558 = int32(0)
	v1560 = F_errstart(m, v36, v1558)
	mBase = m.M
	v1561 = m.ExcPending
	if v1561 != 0 {
		goto L15
	} else {
		goto L591
	}
L589:
	;
	goto L590
L590:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v1554 & int32(-3)
	v2323 = int32(-1)
	goto L11
L591:
	;
	if v1560 == int32(0) {
		v2323 = v1558
		goto L11
	} else {
		goto L592
	}
L592:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L15
	} else {
		goto L593
	}
L593:
	;
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v1567
	F_errmsg(m, int32(_a_F_set_config_with_handle_5), v20-int32(-64))
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L15
	} else {
		goto L594
	}
L594:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(4071), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v1578 = m.ExcPending
	if v1578 != 0 {
		goto L15
	} else {
		goto L595
	}
L595:
	;
	v2323 = v1558
	goto L11
L596:
	;
	if v377 == int32(0) {
		goto L635
	} else {
		goto L636
	}
L597:
	;
	if v377 == int32(0) {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v1588 = m.ExcPending
	if v1588 != 0 {
		goto L15
	} else {
		goto L601
	}
L599:
	;
	goto L600
L600:
	;
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v47)+104))
	if v1590 != 0 {
		goto L602
	} else {
		goto L603
	}
L601:
	;
	goto L600
L602:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	m.T0[v1590].(func(*base.Module, int32, int32))(m, v1589, v1591)
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L15
	} else {
		goto L605
	}
L603:
	;
	v1595 = v1589
	goto L604
L604:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	F_set_string_field(m, v47, v1596, v1595)
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L15
	} else {
		goto L606
	}
L605:
	;
	v1594 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v1595 = v1594
	goto L604
L606:
	;
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(60), v1601)
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L15
	} else {
		goto L607
	}
L607:
	;
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v1606 == int32(0) {
		goto L610
	} else {
		goto L611
	}
L608:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v1444
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v1445
	if l9 != 0 {
		goto L596
	} else {
		goto L619
	}
L609:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1443
	goto L608
L610:
	;
	if v1443 == int32(0) {
		goto L609
	} else {
		goto L613
	}
L611:
	;
	goto L612
L612:
	;
	if v1443 != 0 {
		goto L609
	} else {
		goto L618
	}
L613:
	;
	v1612 = v47 - int32(-64)
	v1614 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[3]))
	if v1614 != 0 {
		goto L615
	} else {
		goto L616
	}
L614:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+64)) = v1621
	v1623 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v1623
	*(*int32)(unsafe.Add(mBase, uint32(v1621)+4)) = v1612
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4])) = v1612
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1443
	goto L608
L615:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4]))
	v1621 = v1616
	goto L614
L616:
	;
	goto L617
L617:
	;
	v1618 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[3])) = v1618
	v1621 = v1618
	goto L614
L618:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+4)) = v1630
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v1630))) = v1632
	goto L609
L619:
	;
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v1641 = int32(_a_F_set_config_with_handle_15)
	v1644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1640))))
	v1647 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_set_config_with_handle[5])))
	if base.B2i32(v1644 == int32(0))|base.B2i32(v1644 != v1647) != 0 {
		v1665 = v1644
		v1666 = v1647
		goto L621
	} else {
		goto L622
	}
L620:
	;
	if v1665-v1666 != 0 {
		goto L596
	} else {
		goto L627
	}
L621:
	;
	goto L620
L622:
	;
	v1650 = v1640
	v1651 = v1641
	goto L623
L623:
	;
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1651)+1)))
	v1655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1650)+1)))
	if v1655 == int32(0) {
		v1665 = v1655
		v1666 = v1654
		goto L621
	} else {
		goto L625
	}
L624:
	;
	v1665 = v1655
	v1666 = v1654
	goto L621
L625:
	;
	v1658 = int32(1)
	if v1655 == v1654 {
		v1650 = v1650 + v1658
		v1651 = v1651 + v1658
		goto L623
	} else {
		goto L626
	}
L626:
	;
	goto L624
L627:
	;
	v1669 = int32(0)
	if l2 != 0 {
		goto L628
	} else {
		goto L629
	}
L628:
	;
	v1672 = int32(_a_F_set_config_with_handle_16)
	goto L630
L629:
	;
	v1672 = v1669
	goto L630
L630:
	;
	if l4 == int32(10) {
		goto L631
	} else {
		goto L632
	}
L631:
	;
	v1676 = int32(1)
	goto L633
L632:
	;
	v1676 = l4
	goto L633
L633:
	;
	v1679 = F_set_config_with_handle(m, int32(_a_F_set_config_with_handle_17), v1669, v1672, l3, v1676, l5, l6, int32(1), v36, int32(0))
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L15
	} else {
		goto L634
	}
L634:
	;
	goto L596
L635:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if v1874 == int32(0) {
		goto L684
	} else {
		goto L685
	}
L636:
	;
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v1685) <= base.Ui32(v1443) {
		goto L637
	} else {
		goto L638
	}
L637:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	F_set_string_field(m, v47, v47+int32(112), v1689)
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L15
	} else {
		goto L640
	}
L638:
	;
	goto L639
L639:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v1700 == int32(0) {
		goto L635
	} else {
		goto L642
	}
L640:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(116), v1694)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L15
	} else {
		goto L641
	}
L641:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v1444
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v1445
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v1443
	goto L639
L642:
	;
	v1704 = v47 + int32(56)
	v1713 = v1700
	goto L643
L643:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+12))
	if base.Ui32(v1722) <= base.Ui32(v1443) {
		goto L645
	} else {
		goto L646
	}
L644:
	;
	goto L635
L645:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+32))
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+32)) = v1725
	if v1724 == int32(0) {
		goto L648
	} else {
		goto L649
	}
L646:
	;
	goto L647
L647:
	;
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1713)))
	if v1856 != 0 {
		v1713 = v1856
		goto L643
	} else {
		goto L683
	}
L648:
	;
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+40))
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+40)) = v1778
	if v1777 == int32(0) {
		goto L661
	} else {
		goto L662
	}
L649:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1729)))
	if v1724 == v1730 {
		goto L648
	} else {
		goto L650
	}
L650:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	if v1724 == v1732 {
		goto L648
	} else {
		goto L651
	}
L651:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	if v1724 == v1734 {
		goto L648
	} else {
		goto L652
	}
L652:
	;
	v1740 = v1704
	goto L653
L653:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(v1740)))
	if v1753 != 0 {
		goto L655
	} else {
		goto L656
	}
L654:
	;
	F_pfree(m, v1724)
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L15
	} else {
		goto L660
	}
L655:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1753)+32))
	if v1724 == v1754 {
		goto L648
	} else {
		goto L658
	}
L656:
	;
	goto L657
L657:
	;
	goto L654
L658:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1753)+48))
	if v1724 != v1756 {
		v1740 = v1753
		goto L653
	} else {
		goto L659
	}
L659:
	;
	goto L648
L660:
	;
	goto L648
L661:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+24)) = v1444
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+16)) = v1445
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+12)) = v1443
	goto L647
L662:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1777 == v1782 {
		goto L661
	} else {
		goto L663
	}
L663:
	;
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1784 {
	case 0:
		goto L669
	case 1:
		goto L668
	case 2:
		goto L667
	case 3:
		goto L666
	case 4:
		goto L665
	default:
		goto L664
	}
L664:
	;
	v1799 = v1704
	goto L675
L665:
	;
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1777 == v1793 {
		goto L661
	} else {
		goto L674
	}
L666:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1777 != v1791 {
		goto L664
	} else {
		goto L673
	}
L667:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1777 != v1789 {
		goto L664
	} else {
		goto L672
	}
L668:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1777 != v1787 {
		goto L664
	} else {
		goto L671
	}
L669:
	;
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1777 != v1785 {
		goto L664
	} else {
		goto L670
	}
L670:
	;
	goto L661
L671:
	;
	goto L661
L672:
	;
	goto L661
L673:
	;
	goto L661
L674:
	;
	goto L664
L675:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1799)))
	if v1812 != 0 {
		goto L677
	} else {
		goto L678
	}
L676:
	;
	F_pfree(m, v1777)
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L15
	} else {
		goto L682
	}
L677:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+40))
	if v1777 == v1813 {
		goto L661
	} else {
		goto L680
	}
L678:
	;
	goto L679
L679:
	;
	goto L676
L680:
	;
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v1812)+56))
	if v1777 != v1815 {
		v1799 = v1812
		goto L675
	} else {
		goto L681
	}
L681:
	;
	goto L661
L682:
	;
	goto L661
L683:
	;
	goto L644
L684:
	;
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1905 == int32(0) {
		goto L698
	} else {
		goto L699
	}
L685:
	;
	v1878 = int32(1)
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1879)))
	if v1874 == v1880 {
		v1902 = v1878
		goto L687
	} else {
		goto L688
	}
L686:
	;
	if v1902 != 0 {
		goto L684
	} else {
		goto L696
	}
L687:
	;
	goto L686
L688:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v47)+112))
	if v1874 == v1882 {
		v1902 = v1878
		goto L687
	} else {
		goto L689
	}
L689:
	;
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	if v1874 == v1884 {
		v1902 = v1878
		goto L687
	} else {
		goto L690
	}
L690:
	;
	v1888 = v47 + int32(56)
	goto L691
L691:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1888)))
	v1892 = int32(0)
	v1893 = base.B2i32(v1891 != v1892)
	if v1891 == v1892 {
		v1902 = v1893
		goto L687
	} else {
		goto L693
	}
L692:
	;
	v1902 = v1893
	goto L687
L693:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+32))
	if v1874 == v1896 {
		v1902 = v1893
		goto L687
	} else {
		goto L694
	}
L694:
	;
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1891)+48))
	if v1874 != v1898 {
		v1888 = v1891
		goto L691
	} else {
		goto L695
	}
L695:
	;
	goto L692
L696:
	;
	F_pfree(m, v1874)
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L15
	} else {
		goto L697
	}
L697:
	;
	goto L684
L698:
	;
	if v403 != 0 {
		goto L130
	} else {
		goto L721
	}
L699:
	;
	v1909 = int32(1)
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1905 == v1910 {
		v1939 = v1909
		goto L701
	} else {
		goto L702
	}
L700:
	;
	if v1939 != 0 {
		goto L698
	} else {
		goto L719
	}
L701:
	;
	goto L700
L702:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1912 {
	case 0:
		goto L708
	case 1:
		goto L707
	case 2:
		goto L706
	case 3:
		goto L705
	case 4:
		goto L704
	default:
		goto L703
	}
L703:
	;
	v1925 = v47 + int32(56)
	goto L714
L704:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1905 == v1921 {
		v1939 = v1909
		goto L701
	} else {
		goto L713
	}
L705:
	;
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1905 != v1919 {
		goto L703
	} else {
		goto L712
	}
L706:
	;
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1905 != v1917 {
		goto L703
	} else {
		goto L711
	}
L707:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1905 != v1915 {
		goto L703
	} else {
		goto L710
	}
L708:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1905 != v1913 {
		goto L703
	} else {
		goto L709
	}
L709:
	;
	v1939 = v1909
	goto L701
L710:
	;
	v1939 = v1909
	goto L701
L711:
	;
	v1939 = v1909
	goto L701
L712:
	;
	v1939 = v1909
	goto L701
L713:
	;
	goto L703
L714:
	;
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v1925)))
	v1929 = int32(0)
	v1930 = base.B2i32(v1928 != v1929)
	if v1928 == v1929 {
		v1939 = v1930
		goto L701
	} else {
		goto L716
	}
L715:
	;
	v1939 = v1930
	goto L701
L716:
	;
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+40))
	if v1905 == v1933 {
		v1939 = v1930
		goto L701
	} else {
		goto L717
	}
L717:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1928)+56))
	if v1905 != v1935 {
		v1925 = v1928
		goto L714
	} else {
		goto L718
	}
L718:
	;
	goto L715
L719:
	;
	F_pfree(m, v1905)
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L15
	} else {
		goto L720
	}
L720:
	;
	goto L698
L721:
	;
	v2323 = int32(-1)
	goto L11
L722:
	;
	if v259 != 0 {
		goto L733
	} else {
		goto L734
	}
L723:
	;
	v1947 = F_parse_and_validate_value(m, v47, l2, l4, v36, v20+int32(264), v20+int32(260))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L15
	} else {
		goto L726
	}
L724:
	;
	goto L725
L725:
	;
	if l4 == int32(0) {
		goto L728
	} else {
		goto L729
	}
L726:
	;
	if v1947 != 0 {
		v1970 = l3
		v1971 = l4
		v1972 = l5
		goto L722
	} else {
		goto L727
	}
L727:
	;
	v2323 = int32(0)
	goto L11
L728:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v1952
	v1954 = int32(0)
	v1961 = F_call_enum_check_hook(m, v47, v20+int32(264), v20+int32(260), v1954, v36)
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		goto L15
	} else {
		goto L731
	}
L729:
	;
	goto L730
L730:
	;
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v1963
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v1965
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v1970 = v1968
	v1971 = v1969
	v1972 = v1967
	goto L722
L731:
	;
	if v1961 != 0 {
		v1970 = l3
		v1971 = v1954
		v1972 = l5
		goto L722
	} else {
		goto L732
	}
L732:
	;
	v2323 = v1954
	goto L11
L733:
	;
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v1974 == int32(0) {
		goto L736
	} else {
		goto L737
	}
L734:
	;
	goto L735
L735:
	;
	if v403 != 0 {
		goto L767
	} else {
		goto L768
	}
L736:
	;
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v2012 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v2012)))
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	if v2013 != v2014 {
		goto L759
	} else {
		goto L760
	}
L737:
	;
	v1978 = int32(1)
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v1974 == v1979 {
		v2008 = v1978
		goto L739
	} else {
		goto L740
	}
L738:
	;
	if v2008 != 0 {
		goto L736
	} else {
		goto L757
	}
L739:
	;
	goto L738
L740:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v1981 {
	case 0:
		goto L746
	case 1:
		goto L745
	case 2:
		goto L744
	case 3:
		goto L743
	case 4:
		goto L742
	default:
		goto L741
	}
L741:
	;
	v1994 = v47 + int32(56)
	goto L752
L742:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v1974 == v1990 {
		v2008 = v1978
		goto L739
	} else {
		goto L751
	}
L743:
	;
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1974 != v1988 {
		goto L741
	} else {
		goto L750
	}
L744:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v1974 != v1986 {
		goto L741
	} else {
		goto L749
	}
L745:
	;
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v1974 != v1984 {
		goto L741
	} else {
		goto L748
	}
L746:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v1974 != v1982 {
		goto L741
	} else {
		goto L747
	}
L747:
	;
	v2008 = v1978
	goto L739
L748:
	;
	v2008 = v1978
	goto L739
L749:
	;
	v2008 = v1978
	goto L739
L750:
	;
	v2008 = v1978
	goto L739
L751:
	;
	goto L741
L752:
	;
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1994)))
	v1998 = int32(0)
	v1999 = base.B2i32(v1997 != v1998)
	if v1997 == v1998 {
		v2008 = v1999
		goto L739
	} else {
		goto L754
	}
L753:
	;
	v2008 = v1999
	goto L739
L754:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v1997)+40))
	if v1974 == v2002 {
		v2008 = v1999
		goto L739
	} else {
		goto L755
	}
L755:
	;
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v1997)+56))
	if v1974 != v2004 {
		v1994 = v1997
		goto L752
	} else {
		goto L756
	}
L756:
	;
	goto L753
L757:
	;
	F_pfree(m, v1974)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L15
	} else {
		goto L758
	}
L758:
	;
	goto L736
L759:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v2011 | int32(2)
	v2019 = int32(0)
	v2021 = F_errstart(m, v36, v2019)
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L15
	} else {
		goto L762
	}
L760:
	;
	goto L761
L761:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v2011 & int32(-3)
	v2323 = int32(-1)
	goto L11
L762:
	;
	if v2021 == int32(0) {
		v2323 = v2019
		goto L11
	} else {
		goto L763
	}
L763:
	;
	F_errcode(m, int32(33685829))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L15
	} else {
		goto L764
	}
L764:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v2028
	F_errmsg(m, int32(_a_F_set_config_with_handle_5), v20+int32(80))
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L15
	} else {
		goto L765
	}
L765:
	;
	F_errfinish(m, int32(_a_F_set_config_with_handle_1), int32(_a_F_set_config_with_handle_18), int32(_a_F_set_config_with_handle_2))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L15
	} else {
		goto L766
	}
L766:
	;
	v2323 = v2019
	goto L11
L767:
	;
	if v377 == int32(0) {
		goto L770
	} else {
		goto L771
	}
L768:
	;
	goto L769
L769:
	;
	if v377 == int32(0) {
		goto L790
	} else {
		goto L791
	}
L770:
	;
	F_push_old_value(m, v47, l6)
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L15
	} else {
		goto L773
	}
L771:
	;
	goto L772
L772:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v47)+108))
	if v2049 != 0 {
		goto L774
	} else {
		goto L775
	}
L773:
	;
	goto L772
L774:
	;
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	m.T0[v2049].(func(*base.Module, int32, int32))(m, v2048, v2050)
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L15
	} else {
		goto L777
	}
L775:
	;
	v2054 = v2048
	goto L776
L776:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v47)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2055))) = v2054
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(60), v2059)
	mBase = m.M
	v2061 = m.ExcPending
	if v2061 != 0 {
		goto L15
	} else {
		goto L778
	}
L777:
	;
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v2054 = v2053
	goto L776
L778:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	if v2064 == int32(0) {
		goto L781
	} else {
		goto L782
	}
L779:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v1972
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v1970
	goto L769
L780:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1971
	goto L779
L781:
	;
	if v1971 == int32(0) {
		goto L780
	} else {
		goto L784
	}
L782:
	;
	goto L783
L783:
	;
	if v1971 != 0 {
		goto L780
	} else {
		goto L789
	}
L784:
	;
	v2070 = v47 - int32(-64)
	v2072 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[3]))
	if v2072 != 0 {
		goto L786
	} else {
		goto L787
	}
L785:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+64)) = v2079
	v2081 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v2081
	*(*int32)(unsafe.Add(mBase, uint32(v2079)+4)) = v2070
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4])) = v2070
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v1971
	goto L779
L786:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[4]))
	v2079 = v2074
	goto L785
L787:
	;
	goto L788
L788:
	;
	v2076 = int32(_a_F_set_config_with_handle_14)
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[3])) = v2076
	v2079 = v2076
	goto L785
L789:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v2087)+4)) = v2088
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v2088))) = v2090
	goto L780
L790:
	;
	v2237 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	if v2237 == int32(0) {
		goto L825
	} else {
		goto L826
	}
L791:
	;
	v2102 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	if base.Ui32(v2102) <= base.Ui32(v1971) {
		goto L792
	} else {
		goto L793
	}
L792:
	;
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+116)) = v2104
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	F_set_extra_field(m, v47, v47+int32(120), v2108)
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L15
	} else {
		goto L795
	}
L793:
	;
	goto L794
L794:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v2114 == int32(0) {
		goto L790
	} else {
		goto L796
	}
L795:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v1972
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v1970
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v1971
	goto L794
L796:
	;
	v2127 = v2114
	goto L797
L797:
	;
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v2127)+12))
	if base.Ui32(v2136) <= base.Ui32(v1971) {
		goto L799
	} else {
		goto L800
	}
L798:
	;
	goto L790
L799:
	;
	v2138 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v2127)+32)) = v2138
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v2127)+40))
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(v20)+260))
	*(*int32)(unsafe.Add(mBase, uint32(v2127)+40)) = v2141
	if v2140 == int32(0) {
		goto L802
	} else {
		goto L803
	}
L800:
	;
	goto L801
L801:
	;
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v2127)))
	if v2219 != 0 {
		v2127 = v2219
		goto L797
	} else {
		goto L824
	}
L802:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2127)+24)) = v1972
	*(*int32)(unsafe.Add(mBase, uint32(v2127)+16)) = v1970
	*(*int32)(unsafe.Add(mBase, uint32(v2127)+12)) = v1971
	goto L801
L803:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v2140 == v2145 {
		goto L802
	} else {
		goto L804
	}
L804:
	;
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v2147 {
	case 0:
		goto L810
	case 1:
		goto L809
	case 2:
		goto L808
	case 3:
		goto L807
	case 4:
		goto L806
	default:
		goto L805
	}
L805:
	;
	v2169 = v47 + int32(56)
	goto L816
L806:
	;
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v2140 == v2156 {
		goto L802
	} else {
		goto L815
	}
L807:
	;
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v2140 != v2154 {
		goto L805
	} else {
		goto L814
	}
L808:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v2140 != v2152 {
		goto L805
	} else {
		goto L813
	}
L809:
	;
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v2140 != v2150 {
		goto L805
	} else {
		goto L812
	}
L810:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v2140 != v2148 {
		goto L805
	} else {
		goto L811
	}
L811:
	;
	goto L802
L812:
	;
	goto L802
L813:
	;
	goto L802
L814:
	;
	goto L802
L815:
	;
	goto L805
L816:
	;
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v2169)))
	if v2175 != 0 {
		goto L818
	} else {
		goto L819
	}
L817:
	;
	F_pfree(m, v2140)
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L15
	} else {
		goto L823
	}
L818:
	;
	v2176 = *(*int32)(unsafe.Add(mBase, uint32(v2175)+40))
	if v2140 == v2176 {
		goto L802
	} else {
		goto L821
	}
L819:
	;
	goto L820
L820:
	;
	goto L817
L821:
	;
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v2175)+56))
	if v2140 != v2178 {
		v2169 = v2175
		goto L816
	} else {
		goto L822
	}
L822:
	;
	goto L802
L823:
	;
	goto L802
L824:
	;
	goto L798
L825:
	;
	if v403 != 0 {
		goto L130
	} else {
		goto L848
	}
L826:
	;
	v2241 = int32(1)
	v2242 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	if v2237 == v2242 {
		v2271 = v2241
		goto L828
	} else {
		goto L829
	}
L827:
	;
	if v2271 != 0 {
		goto L825
	} else {
		goto L846
	}
L828:
	;
	goto L827
L829:
	;
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	switch v2244 {
	case 0:
		goto L835
	case 1:
		goto L834
	case 2:
		goto L833
	case 3:
		goto L832
	case 4:
		goto L831
	default:
		goto L830
	}
L830:
	;
	v2257 = v47 + int32(56)
	goto L841
L831:
	;
	v2253 = *(*int32)(unsafe.Add(mBase, uint32(v47)+120))
	if v2237 == v2253 {
		v2271 = v2241
		goto L828
	} else {
		goto L840
	}
L832:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v2237 != v2251 {
		goto L830
	} else {
		goto L839
	}
L833:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v2237 != v2249 {
		goto L830
	} else {
		goto L838
	}
L834:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(v47)+124))
	if v2237 != v2247 {
		goto L830
	} else {
		goto L837
	}
L835:
	;
	v2245 = *(*int32)(unsafe.Add(mBase, uint32(v47)+116))
	if v2237 != v2245 {
		goto L830
	} else {
		goto L836
	}
L836:
	;
	v2271 = v2241
	goto L828
L837:
	;
	v2271 = v2241
	goto L828
L838:
	;
	v2271 = v2241
	goto L828
L839:
	;
	v2271 = v2241
	goto L828
L840:
	;
	goto L830
L841:
	;
	v2260 = *(*int32)(unsafe.Add(mBase, uint32(v2257)))
	v2261 = int32(0)
	v2262 = base.B2i32(v2260 != v2261)
	if v2260 == v2261 {
		v2271 = v2262
		goto L828
	} else {
		goto L843
	}
L842:
	;
	v2271 = v2262
	goto L828
L843:
	;
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(v2260)+40))
	if v2237 == v2265 {
		v2271 = v2262
		goto L828
	} else {
		goto L844
	}
L844:
	;
	v2267 = *(*int32)(unsafe.Add(mBase, uint32(v2260)+56))
	if v2237 != v2267 {
		v2257 = v2260
		goto L841
	} else {
		goto L845
	}
L845:
	;
	goto L842
L846:
	;
	F_pfree(m, v2237)
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L15
	} else {
		goto L847
	}
L847:
	;
	goto L825
L848:
	;
	v2323 = int32(-1)
	goto L11
L849:
	;
	v2323 = int32(-1)
	goto L11
L850:
	;
	v2299 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	if v2299&int32(4) != 0 {
		v2323 = v2293
		goto L11
	} else {
		goto L851
	}
L851:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v2299 | int32(4)
	v2305 = int32(_a_F_set_config_with_handle_19)
	v2306 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+76)) = v2306
	*(*int32)(unsafe.Add(mBase, _c_F_set_config_with_handle[6])) = v47 + int32(76)
	v2323 = v2293
	goto L11
}
