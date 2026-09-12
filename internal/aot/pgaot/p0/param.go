package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BuildParamLogString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 != 0 {
		v185 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return v185
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[37]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	goto L3
L3:
	;
	if (v19-int32(7))&int32(-9) == int32(0) {
		v185 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_initStringInfo(m, v14+int32(32))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v38 = F_AllocSetContextCreateInternal(m, v33, int32(347757), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v40 = int32(4562080)
	v41 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v38
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v44 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v41
	F_MemoryContextDelete(m, v38)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L42
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(791891)
	F_appendStringInfo(m, v14+int32(32), int32(779723), v14+int32(16))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v59 = l0 + int32(32)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v60 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v96 = int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v97 <= v96 {
		goto L8
	} else {
		goto L23
	}
L12:
	;
	F_appendStringInfoString(m, v14+int32(32), int32(559295))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L22
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v61 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if l1 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_getTypeOutputInfo(m, v61, v14+int32(28), v14+int32(27))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L19
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v66 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_appendStringInfoStringQuoted(m, v14+int32(32), v66, l2)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v84 = F_OidOutputFunctionCall(m, v82, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_appendStringInfoStringQuoted(m, v14+int32(32), v84, l2)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	goto L11
L22:
	;
	goto L11
L23:
	;
	v104 = v96
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(780599)
	v114 = v104 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v114
	F_appendStringInfo(m, v14+int32(32), int32(779723), v14)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	goto L8
L26:
	;
	v123 = v59 + v104*int32(12)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+4)))
	if v124 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v114 < v163 {
		v104 = v114
		goto L24
	} else {
		goto L41
	}
L28:
	;
	if l1 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123)+8))
	if v127 != 0 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_appendStringInfoString(m, v14+int32(32), int32(559295))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L5
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	goto L27
L34:
	;
	F_getTypeOutputInfo(m, v127, v14+int32(28), v14+int32(27))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L38
	}
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1+v104<<(uint(int32(2))%32))))
	if v139 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	F_appendStringInfoStringQuoted(m, v14+int32(32), v139, l2)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	goto L27
L38:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v157 = F_OidOutputFunctionCall(m, v155, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	F_appendStringInfoStringQuoted(m, v14+int32(32), v157, l2)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	goto L27
L41:
	;
	goto L25
L42:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v185 = v180
	goto L1
}
func F_ExecSetParamPlanMulti(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if int32(0) <= v61 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v61 = base.I32_ctz(v47) | v48<<(uint(int32(5))%32)
	goto L1
L3:
	;
	v61 = int32(-2)
	goto L1
L4:
	;
	v14 = base.I32_div_s(int32(0), int32(32))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= v14 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v18 = l0 + int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v14<<(uint(int32(2))%32))))
	v25 = v22 & int32(-1)
	if v25 != 0 {
		v47 = v25
		v48 = v14
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v27 = v14 + int32(1)
	if v27 == v15 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v30 = v27
	goto L8
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v18+v30<<(uint(int32(2))%32))))
	if v37 != 0 {
		v47 = v37
		v48 = v30
		goto L2
	} else {
		goto L10
	}
L9:
	;
	goto L3
L10:
	;
	v39 = v30 + int32(1)
	if v39 != v15 {
		v30 = v39
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v66 = v61
	goto L15
L13:
	;
	goto L14
L14:
	;
	return
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v66*int32(12))))
	if v72 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	F_ExecSetParamPlan(m, v72, l1)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	if l0 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	return
L21:
	;
	goto L19
L22:
	;
	if int32(0) <= v130 {
		v66 = v130
		goto L15
	} else {
		goto L33
	}
L23:
	;
	v130 = base.I32_ctz(v116) | v117<<(uint(int32(5))%32)
	goto L22
L24:
	;
	v130 = int32(-2)
	goto L22
L25:
	;
	v81 = v66 + int32(1)
	v83 = base.I32_div_s(v81, int32(32))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v84 <= v83 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v87 = l0 + int32(8)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v83<<(uint(int32(2))%32))))
	v94 = v91 & (int32(-1) << (uint(v81) % 32))
	if v94 != 0 {
		v116 = v94
		v117 = v83
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v96 = v83 + int32(1)
	if v96 == v84 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v99 = v96
	goto L29
L29:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v87+v99<<(uint(int32(2))%32))))
	if v106 != 0 {
		v116 = v106
		v117 = v99
		goto L23
	} else {
		goto L31
	}
L30:
	;
	goto L24
L31:
	;
	v108 = v99 + int32(1)
	if v108 != v84 {
		v99 = v108
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	goto L16
}
func F_SerializeParamExecParams(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v399 int32
	_ = v399
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(4)
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if int32(0) <= v71 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v71 = base.I32_ctz(v57) | v58<<(uint(int32(5))%32)
	goto L1
L3:
	;
	v71 = int32(-2)
	goto L1
L4:
	;
	v24 = base.I32_div_s(int32(0), int32(32))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v25 <= v24 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v28 = l1 + int32(8)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v24<<(uint(int32(2))%32))))
	v35 = v32 & int32(-1)
	if v35 != 0 {
		v57 = v35
		v58 = v24
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v37 = v24 + int32(1)
	if v37 == v25 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v40 = v37
	goto L8
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v28+v40<<(uint(int32(2))%32))))
	if v47 != 0 {
		v57 = v47
		v58 = v40
		goto L2
	} else {
		goto L10
	}
L9:
	;
	goto L3
L10:
	;
	v49 = v40 + int32(1)
	if v49 != v25 {
		v40 = v49
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v78 = v71
	v79 = v14
	goto L15
L13:
	;
	v186 = v14
	goto L14
L14:
	;
	v191 = F_dsa_allocate_extended(m, l2, v186, int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L17
	} else {
		goto L38
	}
L15:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v86 = v83 + v78*int32(12)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+84))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v78<<(uint(int32(2))%32))))
	v95 = F_add_size(m, v79, int32(4))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v186 = v121
	goto L14
L17:
	;
	return int32(0)
L18:
	;
	if v93 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+8)))
	v119 = F_datumEstimateSpace(m, v115, v116, v113&int32(1), v114)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L17
	} else {
		goto L24
	}
L20:
	;
	F_get_typlenbyval(m, v93, v12+int32(14), v12+int32(13))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L17
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v107)
	v110 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+14)) = uint16(v110)
	v113 = v107
	v114 = v110
	goto L19
L23:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)))
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+14)))
	v113 = v105
	v114 = v106
	goto L19
L24:
	;
	v121 = F_add_size(m, v95, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	if l1 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	if int32(0) <= v178 {
		v78 = v178
		v79 = v121
		goto L15
	} else {
		goto L37
	}
L27:
	;
	v178 = base.I32_ctz(v164) | v165<<(uint(int32(5))%32)
	goto L26
L28:
	;
	v178 = int32(-2)
	goto L26
L29:
	;
	v129 = v78 + int32(1)
	v131 = base.I32_div_s(v129, int32(32))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v132 <= v131 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v135 = l1 + int32(8)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v131<<(uint(int32(2))%32))))
	v142 = v139 & (int32(-1) << (uint(v129) % 32))
	if v142 != 0 {
		v164 = v142
		v165 = v131
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v144 = v131 + int32(1)
	if v144 == v132 {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v147 = v144
	goto L33
L33:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v135+v147<<(uint(int32(2))%32))))
	if v154 != 0 {
		v164 = v154
		v165 = v147
		goto L27
	} else {
		goto L35
	}
L34:
	;
	goto L28
L35:
	;
	v156 = v147 + int32(1)
	if v156 != v132 {
		v147 = v156
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	goto L16
L38:
	;
	v193 = F_dsa_get_address(m, l2, v191)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L17
	} else {
		goto L39
	}
L39:
	;
	v195 = int32(0)
	if l1 == v195 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v193 + int32(4)
	if l1 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L41:
	;
	v230 = int32(0)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v202 = int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v203 <= v202 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v206 = v202
	goto L46
L45:
	;
	v206 = v203
	goto L46
L46:
	;
	v210 = int32(0)
	v212 = v195
	goto L47
L47:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(8)+v210<<(uint(int32(2))%32))))
	if v218 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v230 = v221
	goto L40
L49:
	;
	v221 = v212 + base.I32_popcnt(v218)
	goto L51
L50:
	;
	v221 = v212
	goto L51
L51:
	;
	v223 = v210 + int32(1)
	if v223 != v206 {
		v210 = v223
		v212 = v221
		goto L47
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	if int32(0) <= v291 {
		goto L64
	} else {
		goto L65
	}
L54:
	;
	v291 = base.I32_ctz(v277) | v278<<(uint(int32(5))%32)
	goto L53
L55:
	;
	v291 = int32(-2)
	goto L53
L56:
	;
	v244 = base.I32_div_s(int32(0), int32(32))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v245 <= v244 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v248 = l1 + int32(8)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v248+v244<<(uint(int32(2))%32))))
	v255 = v252 & int32(-1)
	if v255 != 0 {
		v277 = v255
		v278 = v244
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v257 = v244 + int32(1)
	if v257 == v245 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v260 = v257
	goto L60
L60:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v248+v260<<(uint(int32(2))%32))))
	if v267 != 0 {
		v277 = v267
		v278 = v260
		goto L54
	} else {
		goto L62
	}
L61:
	;
	goto L55
L62:
	;
	v269 = v260 + int32(1)
	if v269 != v245 {
		v260 = v269
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v298 = v291
	goto L67
L65:
	;
	goto L66
L66:
	;
	m.G0 = v12 + int32(16)
	return v191
L67:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)+84))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+12))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v305+v298<<(uint(int32(2))%32))))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v298
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v313 + int32(4)
	v319 = v310 + v298*int32(12)
	if v309 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L66
L69:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+8)))
	F_datumSerialize(m, v336, v337, v334&int32(1), v335, v12+int32(8))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L17
	} else {
		goto L74
	}
L70:
	;
	F_get_typlenbyval(m, v309, v12+int32(6), v12+int32(5))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L17
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v328 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v328)
	v331 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+6)) = uint16(v331)
	v334 = v328
	v335 = v331
	goto L69
L73:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)))
	v327 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
	v334 = v326
	v335 = v327
	goto L69
L74:
	;
	if l1 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	if int32(0) <= v399 {
		v298 = v399
		goto L67
	} else {
		goto L86
	}
L76:
	;
	v399 = base.I32_ctz(v385) | v386<<(uint(int32(5))%32)
	goto L75
L77:
	;
	v399 = int32(-2)
	goto L75
L78:
	;
	v350 = v298 + int32(1)
	v352 = base.I32_div_s(v350, int32(32))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v353 <= v352 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v356 = l1 + int32(8)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v356+v352<<(uint(int32(2))%32))))
	v363 = v360 & (int32(-1) << (uint(v350) % 32))
	if v363 != 0 {
		v385 = v363
		v386 = v352
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v365 = v352 + int32(1)
	if v365 == v353 {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v368 = v365
	goto L82
L82:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v356+v368<<(uint(int32(2))%32))))
	if v375 != 0 {
		v385 = v375
		v386 = v368
		goto L76
	} else {
		goto L84
	}
L83:
	;
	goto L77
L84:
	;
	v377 = v368 + int32(1)
	if v377 != v353 {
		v368 = v377
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	goto L68
}
func F_copyParamList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l0 == v2 {
		v92 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v92
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v16 <= int32(0) {
		v92 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = F_palloc(m, v16*int32(12)+int32(32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v27 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v16
	v30 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(815)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v23
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v37 <= v30 {
		v92 = v23
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v40 = int32(32)
	v49 = v2
	goto L7
L7:
	;
	v54 = v49 * int32(12)
	v55 = v23 + v40 + v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v56 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v92 = v23
	goto L1
L9:
	;
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	*(*int64)(unsafe.Add(mBase, uint32(v55))) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+4)))
	if v70 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v62 = m.T0[v56].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v49+int32(1), int32(0), v12+int32(4))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v65 = l0 + v40 + v54
	goto L9
L13:
	;
	v65 = v62
	goto L9
L14:
	;
	v88 = v49 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v88 < v89 {
		v49 = v88
		goto L7
	} else {
		goto L19
	}
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v71 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	F_get_typlenbyval(m, v71, v12+int32(2), v12+int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+2)))
	v83 = F_datumCopy(m, v80, v81, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v83
	goto L14
L19:
	;
	goto L8
}
func F_count_param_references(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v3 = int32(0)
	if l0 == v3 {
		v30 = v3
		return v30
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v6 == int32(8) {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v9 != 0 {
				return int32(0)
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v10 != v11 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l0
					v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v15 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v14 + v15
					if int32(0) < v14 {
						v30 = v15
						return v30
					} else {
						return int32(0)
					}
				}
			}
		} else {
			v25 = F_expression_tree_walker_impl(m, l0, int32(7340), l1)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v30 = v25
				return v30
			}
		}
	}
}
func F_get_param_path_clause_serials(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
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
	var v69 int32
	_ = v69
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v7 == v2 {
		v88 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v88
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v10 - int32(290) {
	case 0:
		goto L4
	default:
		goto L3
	case 8, 9, 10:
		goto L5
	}
L3:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v88 = v85
	goto L1
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v51 == int32(0) {
		v88 = v2
		goto L1
	} else {
		goto L17
	}
L5:
	;
	v13 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v16 = F_get_param_path_clause_serials(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v20 = F_bms_add_members(m, v13, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v23 = F_get_param_path_clause_serials(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v25 = F_bms_add_members(m, v20, v23)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v27 == int32(0) {
		v88 = v25
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v30 <= int32(0) {
		v88 = v25
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v34 = v13
	v35 = v25
	goto L13
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v34<<(uint(int32(2))%32))))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+56))
	v45 = F_bms_add_member(m, v35, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L15
	}
L14:
	;
	v88 = v45
	goto L1
L15:
	;
	v48 = v34 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v48 < v49 {
		v34 = v48
		v35 = v45
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v54 <= int32(0) {
		v88 = v2
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v59 = int32(0)
	v60 = v2
	goto L19
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v67 = v64 + v59<<(uint(int32(2))%32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v69 = F_get_param_path_clause_serials(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	v88 = v80
	goto L1
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v71 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v82 = v59 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v82 < v83 {
		v59 = v82
		v60 = v80
		goto L19
	} else {
		goto L31
	}
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v74 = v72
	goto L25
L24:
	;
	v74 = int32(0)
	goto L25
L25:
	;
	if v67 == v74 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v76 = F_bms_copy(m, v69)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v78 = F_bms_int_members(m, v60, v69)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L30
	}
L29:
	;
	v80 = v76
	goto L22
L30:
	;
	v80 = v78
	goto L22
L31:
	;
	goto L20
}
