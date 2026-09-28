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
	var v27 int32
	_ = v27
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 != 0 {
		v181 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return v181
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_BuildParamLogString[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	goto L3
L3:
	;
	if (v19-int32(7))&int32(-9) == int32(0) {
		v181 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = v14 + int32(32)
	F_initStringInfo(m, v27)
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
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_BuildParamLogString[1]))
	v38 = F_AllocSetContextCreateInternal(m, v33, int32(_a_F_BuildParamLogString_0), int32(0), int32(_a_F_BuildParamLogString_1), int32(_a_F_BuildParamLogString_2))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v40 = int32(_a_F_BuildParamLogString_3)
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_BuildParamLogString[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildParamLogString[1])) = v38
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v44 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BuildParamLogString[1])) = v41
	F_MemoryContextDelete(m, v38)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L42
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(_a_F_BuildParamLogString_4)
	F_appendStringInfo(m, v27, int32(_a_F_BuildParamLogString_5), v14+int32(16))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v57 = l0 + int32(32)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v58 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v92 = int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v93 <= v92 {
		goto L8
	} else {
		goto L23
	}
L12:
	;
	F_appendStringInfoString(m, v14+int32(32), int32(_a_F_BuildParamLogString_6))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L22
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v59 == int32(0) {
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
	F_getTypeOutputInfo(m, v59, v14+int32(28), v14+int32(27))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L19
	}
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v64 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_appendStringInfoStringQuoted(m, v27, v64, l2)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	v80 = F_OidOutputFunctionCall(m, v78, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_appendStringInfoStringQuoted(m, v14+int32(32), v80, l2)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
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
	v100 = v92
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_BuildParamLogString_7)
	v110 = v100 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v110
	F_appendStringInfo(m, v14+int32(32), int32(_a_F_BuildParamLogString_5), v14)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	goto L8
L26:
	;
	v119 = v57 + v100<<(uint(int32(4))%32)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+8)))
	if v120 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v110 < v159 {
		v100 = v110
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
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	if v123 != 0 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_appendStringInfoString(m, v14+int32(32), int32(_a_F_BuildParamLogString_6))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
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
	F_getTypeOutputInfo(m, v123, v14+int32(28), v14+int32(27))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L38
	}
L35:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1+v100<<(uint(int32(2))%32))))
	if v135 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	F_appendStringInfoStringQuoted(m, v14+int32(32), v135, l2)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	goto L27
L38:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
	v153 = F_OidOutputFunctionCall(m, v151, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	F_appendStringInfoStringQuoted(m, v14+int32(32), v153, l2)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
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
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v181 = v176
	goto L1
}
func F_ExecSetParamPlanMulti(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	v12 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= v12 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v18 = l0 + int32(8)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v25 = v22 & int32(-1)
	if v25 != 0 {
		v47 = v25
		v48 = v12
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v26 = int32(1)
	if v26 == v15 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v30 = v26
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
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v66*int32(24))))
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
	v83 = int32(base.Ui32(v81) >> (uint(int32(5)) % 32))
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v115 int64
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
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int64
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v220 int32
	_ = v220
	var v221 int64
	_ = v221
	var v222 int32
	_ = v222
	var v223 int64
	_ = v223
	var v227 int64
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v235 int64
	_ = v235
	var v240 int32
	_ = v240
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int64
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v409 int32
	_ = v409
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
	v22 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v25 <= v22 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v28 = l1 + int32(8)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v35 = v32 & int32(-1)
	if v35 != 0 {
		v57 = v35
		v58 = v22
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v36 = int32(1)
	if v36 == v25 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v40 = v36
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
	v78 = v14
	v79 = v71
	goto L15
L13:
	;
	v185 = v14
	goto L14
L14:
	;
	v191 = F_dsa_allocate_extended(m, l2, v185, int32(0))
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
	v86 = v83 + v79*int32(24)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+96))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v79<<(uint(int32(2))%32))))
	v95 = F_add_size(m, v78, int32(4))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v185 = v121
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
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v86)+8))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+16)))
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
	v110 = int32(8)
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
		v78 = v121
		v79 = v178
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
	v129 = v79 + int32(1)
	v131 = int32(base.Ui32(v129) >> (uint(int32(5)) % 32))
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
	v196 = int64(0)
	if l1 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v193 + int32(4)
	if l1 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L41:
	;
	v240 = int32(0)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v201 = l1 + int32(8)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v202 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v240 = base.I32_popcnt(v205)
	goto L40
L45:
	;
	goto L46
L46:
	;
	v208 = v202 << (uint(int32(2)) % 32)
	if v208 <= int32(7) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v240 = base.I32_wrap_i64(v235)
	goto L40
L48:
	;
	if v208 == int32(0) {
		v235 = v196
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v232 = F_pg_popcount_optimized(m, v201, v208)
	mBase = m.M
	v235 = v232
	goto L47
L51:
	;
	v213 = v208
	v214 = v201
	v215 = v196
	goto L52
L52:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+3)))
	v217 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v216)+uint32(_c_F_SerializeParamExecParams[0]))))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+2)))
	v219 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v218)+uint32(_c_F_SerializeParamExecParams[0]))))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)))
	v221 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v220)+uint32(_c_F_SerializeParamExecParams[0]))))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	v223 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v222)+uint32(_c_F_SerializeParamExecParams[0]))))
	v227 = v217 + (v219 + (v221 + (v215 + v223)))
	v228 = int32(4)
	v231 = v213 - v228
	if v231 != 0 {
		v213 = v231
		v214 = v214 + v228
		v215 = v227
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v235 = v227
	goto L47
L54:
	;
	goto L53
L55:
	;
	if int32(0) <= v301 {
		goto L66
	} else {
		goto L67
	}
L56:
	;
	v301 = base.I32_ctz(v287) | v288<<(uint(int32(5))%32)
	goto L55
L57:
	;
	v301 = int32(-2)
	goto L55
L58:
	;
	v252 = int32(0)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v255 <= v252 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v258 = l1 + int32(8)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v265 = v262 & int32(-1)
	if v265 != 0 {
		v287 = v265
		v288 = v252
		goto L56
	} else {
		goto L60
	}
L60:
	;
	v266 = int32(1)
	if v266 == v255 {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v270 = v266
	goto L62
L62:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v258+v270<<(uint(int32(2))%32))))
	if v277 != 0 {
		v287 = v277
		v288 = v270
		goto L56
	} else {
		goto L64
	}
L63:
	;
	goto L57
L64:
	;
	v279 = v270 + int32(1)
	if v279 != v255 {
		v270 = v279
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v309 = v301
	goto L69
L67:
	;
	goto L68
L68:
	;
	m.G0 = v12 + int32(16)
	return v191
L69:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+96))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+12))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v315+v309<<(uint(int32(2))%32))))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v309
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v323 + int32(4)
	v329 = v320 + v309*int32(24)
	if v319 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L68
L71:
	;
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v329)+8))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+16)))
	F_datumSerialize(m, v346, v347, v344&int32(1), v345, v12+int32(8))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L17
	} else {
		goto L76
	}
L72:
	;
	F_get_typlenbyval(m, v319, v12+int32(6), v12+int32(5))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L17
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v338 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)) = uint8(v338)
	v341 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+6)) = uint16(v341)
	v344 = v338
	v345 = v341
	goto L71
L75:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+5)))
	v337 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
	v344 = v336
	v345 = v337
	goto L71
L76:
	;
	if l1 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if int32(0) <= v409 {
		v309 = v409
		goto L69
	} else {
		goto L88
	}
L78:
	;
	v409 = base.I32_ctz(v395) | v396<<(uint(int32(5))%32)
	goto L77
L79:
	;
	v409 = int32(-2)
	goto L77
L80:
	;
	v360 = v309 + int32(1)
	v362 = int32(base.Ui32(v360) >> (uint(int32(5)) % 32))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v363 <= v362 {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v366 = l1 + int32(8)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v366+v362<<(uint(int32(2))%32))))
	v373 = v370 & (int32(-1) << (uint(v360) % 32))
	if v373 != 0 {
		v395 = v373
		v396 = v362
		goto L78
	} else {
		goto L82
	}
L82:
	;
	v375 = v362 + int32(1)
	if v375 == v363 {
		goto L79
	} else {
		goto L83
	}
L83:
	;
	v378 = v375
	goto L84
L84:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v366+v378<<(uint(int32(2))%32))))
	if v385 != 0 {
		v395 = v385
		v396 = v378
		goto L78
	} else {
		goto L86
	}
L85:
	;
	goto L79
L86:
	;
	v387 = v378 + int32(1)
	if v387 != v363 {
		v378 = v387
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	goto L70
}
func F_copyParamList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l0 == v2 {
		v96 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v96
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v17 <= int32(0) {
		v96 = v2
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = F_palloc(m, v17<<(uint(int32(4))%32)+int32(32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v28 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v17
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = int32(862)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v24
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v38 <= v33 {
		v96 = v24
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v41 = int32(32)
	v50 = v2
	goto L7
L7:
	;
	v56 = v50 << (uint(int32(4)) % 32)
	v57 = v24 + v41 + v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v58 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v96 = v24
	goto L1
L9:
	;
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+8)) = v68
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v57))) = v70
	if base.I32_wrap_i64(v68)&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v64 = m.T0[v58].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v50+int32(1), int32(0), v13+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v67 = v56 + (l0 + v41)
	goto L9
L13:
	;
	v67 = v64
	goto L9
L14:
	;
	v92 = v50 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v92 < v93 {
		v50 = v92
		goto L7
	} else {
		goto L19
	}
L15:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	if v75 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	F_get_typlenbyval(m, v75, v13+int32(14), v13+int32(13))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)))
	v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+14)))
	v87 = F_datumCopy(m, v84, v85, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v57))) = v87
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
			v25 = F_expression_tree_walker_impl(m, l0, int32(_a_F_count_param_references_0), l1)
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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v8 == v2 {
		v91 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v91
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v11 - int32(293) {
	case 0:
		goto L4
	default:
		goto L3
	case 7, 8, 9:
		goto L5
	}
L3:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v91 = v88
	goto L1
L4:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v53 == int32(0) {
		v91 = v2
		goto L1
	} else {
		goto L17
	}
L5:
	;
	v14 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v17 = F_get_param_path_clause_serials(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v21 = F_bms_add_members(m, v14, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v24 = F_get_param_path_clause_serials(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v26 = F_bms_add_members(m, v21, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v28 == int32(0) {
		v91 = v26
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v31 <= int32(0) {
		v91 = v26
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v35 = v14
	v36 = v26
	goto L13
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v35<<(uint(int32(2))%32))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
	v47 = F_bms_add_member(m, v36, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L15
	}
L14:
	;
	v91 = v47
	goto L1
L15:
	;
	v50 = v35 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v50 < v51 {
		v35 = v50
		v36 = v47
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v56 <= int32(0) {
		v91 = v2
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v61 = int32(0)
	v62 = v2
	goto L19
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v70 = v67 + v61<<(uint(int32(2))%32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = F_get_param_path_clause_serials(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	v91 = v83
	goto L1
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v74 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v85 = v61 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v85 < v86 {
		v61 = v85
		v62 = v83
		goto L19
	} else {
		goto L28
	}
L23:
	;
	v81 = F_bms_int_members(m, v62, v72)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L6
	} else {
		goto L27
	}
L24:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	if v70 != v77 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v79 = F_bms_copy(m, v72)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v83 = v79
	goto L22
L27:
	;
	v83 = v81
	goto L22
L28:
	;
	goto L20
}
