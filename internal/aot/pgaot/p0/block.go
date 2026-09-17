package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BlockRefTableGetEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
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
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int64
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int64
	_ = v300
	var v302 int64
	_ = v302
	var v303 int64
	_ = v303
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	v8 = m.G0
	v9 = int32(16)
	v10 = v8 - v9
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l2
	v24 = int32(-1636608416)
	if v10&int32(3) != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v285 = (v278 ^ v270 - base.I32_rotl(v278, int32(24))) & v284
	v288 = v283 + v285*int32(40)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+20)))
	if v289 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L2:
	;
	v256 = int32(14)
	v258 = v252 ^ v253 - base.I32_rotl(v252, v256)
	v262 = v258 ^ v251 - base.I32_rotl(v258, int32(11))
	v266 = v262 ^ v252 - base.I32_rotl(v262, int32(25))
	v270 = v266 ^ v258 - base.I32_rotl(v266, int32(16))
	v274 = v270 ^ v262 - base.I32_rotl(v270, int32(4))
	v278 = v274 ^ v266 - base.I32_rotl(v274, v256)
	goto L1
L3:
	;
	switch v178 - int32(1) {
	case 0:
		v244 = v169
		v245 = v170
		v246 = v174
		goto L30
	case 1:
		v237 = v169
		v238 = v170
		v239 = v174
		goto L31
	case 2:
		v230 = v169
		v231 = v170
		v232 = v174
		goto L32
	case 3:
		v224 = v170
		v225 = v174
		goto L33
	case 4:
		v220 = v170
		v221 = v174
		goto L34
	case 5:
		v214 = v170
		v215 = v174
		goto L35
	case 6:
		v208 = v170
		v209 = v174
		goto L36
	case 7:
		v203 = v174
		goto L37
	case 8:
		v198 = v174
		goto L38
	case 9:
		v193 = v174
		goto L39
	case 10:
		goto L40
	default:
		v251 = v169
		v252 = v170
		v253 = v174
		goto L2
	}
L4:
	;
	v133 = v10
	v134 = v9
	v135 = v24
	v136 = v24
	v137 = v24
	goto L27
L5:
	;
	goto L4
L6:
	;
	goto L7
L7:
	;
	goto L11
L9:
	;
	switch v76 - int32(1) {
	case 0:
		v130 = v67
		goto L16
	case 1:
		v125 = v67
		goto L17
	case 2:
		goto L18
	case 3:
		v118 = v68
		goto L19
	case 4:
		v115 = v68
		goto L20
	case 5:
		v110 = v68
		goto L21
	case 6:
		goto L22
	case 7:
		v101 = v72
		goto L23
	case 8:
		v96 = v72
		goto L24
	case 9:
		v91 = v72
		goto L25
	case 10:
		goto L26
	default:
		v251 = v67
		v252 = v68
		v253 = v72
		goto L2
	}
L11:
	;
	goto L12
L12:
	;
	v31 = v10
	v32 = v9
	v33 = v24
	v34 = v24
	v35 = v24
	goto L13
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v38 = v37 + v34
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v42 = v41 + v35
	v44 = int32(4)
	v46 = v39 + v33 - v42 ^ base.I32_rotl(v42, v44)
	v50 = v38 - v46 ^ base.I32_rotl(v46, int32(6))
	v51 = v42 + v38
	v52 = v46 + v51
	v53 = v50 + v52
	v57 = v51 - v50 ^ base.I32_rotl(v50, int32(8))
	v61 = v52 - v57 ^ base.I32_rotl(v57, int32(16))
	v65 = v53 - v61 ^ base.I32_rotl(v61, int32(19))
	v66 = v57 + v53
	v67 = v61 + v66
	v68 = v65 + v67
	v72 = v66 - v65 ^ base.I32_rotl(v65, v44)
	v73 = int32(12)
	v74 = v31 + v73
	v76 = v32 - v73
	if base.Ui32(int32(11)) < base.Ui32(v76) {
		v31 = v74
		v32 = v76
		v33 = v67
		v34 = v68
		v35 = v72
		goto L13
	} else {
		goto L15
	}
L14:
	;
	goto L9
L15:
	;
	goto L14
L16:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	v251 = v130 + v131
	v252 = v68
	v253 = v72
	goto L2
L17:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	v130 = v126<<(uint(int32(8))%32) + v125
	goto L16
L18:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+2)))
	v125 = v121<<(uint(int32(16))%32) + v67
	goto L17
L19:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v251 = v119 + v67
	v252 = v118
	v253 = v72
	goto L2
L20:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)))
	v118 = v115 + v116
	goto L19
L21:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+5)))
	v115 = v111<<(uint(int32(8))%32) + v110
	goto L20
L22:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)))
	v110 = v106<<(uint(int32(16))%32) + v68
	goto L21
L23:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v251 = v102 + v67
	v252 = v104 + v68
	v253 = v101
	goto L2
L24:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+8)))
	v101 = v97<<(uint(int32(8))%32) + v96
	goto L23
L25:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+9)))
	v96 = v92<<(uint(int32(16))%32) + v91
	goto L24
L26:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+10)))
	v91 = v87<<(uint(int32(24))%32) + v72
	goto L25
L27:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	v140 = v139 + v136
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	v144 = v143 + v137
	v146 = int32(4)
	v148 = v141 + v135 - v144 ^ base.I32_rotl(v144, v146)
	v152 = v140 - v148 ^ base.I32_rotl(v148, int32(6))
	v153 = v144 + v140
	v154 = v148 + v153
	v155 = v152 + v154
	v159 = v153 - v152 ^ base.I32_rotl(v152, int32(8))
	v163 = v154 - v159 ^ base.I32_rotl(v159, int32(16))
	v167 = v155 - v163 ^ base.I32_rotl(v163, int32(19))
	v168 = v159 + v155
	v169 = v163 + v168
	v170 = v167 + v169
	v174 = v168 - v167 ^ base.I32_rotl(v167, v146)
	v175 = int32(12)
	v176 = v133 + v175
	v178 = v134 - v175
	if base.Ui32(int32(11)) < base.Ui32(v178) {
		v133 = v176
		v134 = v178
		v135 = v169
		v136 = v170
		v137 = v174
		goto L27
	} else {
		goto L29
	}
L28:
	;
	goto L3
L29:
	;
	goto L28
L30:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v251 = v244 + v247
	v252 = v245
	v253 = v246
	goto L2
L31:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+1)))
	v244 = v240<<(uint(int32(8))%32) + v237
	v245 = v238
	v246 = v239
	goto L30
L32:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+2)))
	v237 = v233<<(uint(int32(16))%32) + v230
	v238 = v231
	v239 = v232
	goto L31
L33:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+3)))
	v230 = v226<<(uint(int32(24))%32) + v169
	v231 = v224
	v232 = v225
	goto L32
L34:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+4)))
	v224 = v220 + v222
	v225 = v221
	goto L33
L35:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+5)))
	v220 = v216<<(uint(int32(8))%32) + v214
	v221 = v215
	goto L34
L36:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+6)))
	v214 = v210<<(uint(int32(16))%32) + v208
	v215 = v209
	goto L35
L37:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+7)))
	v208 = v204<<(uint(int32(24))%32) + v170
	v209 = v203
	goto L36
L38:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+8)))
	v203 = v199<<(uint(int32(8))%32) + v198
	goto L37
L39:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+9)))
	v198 = v194<<(uint(int32(16))%32) + v193
	goto L38
L40:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+10)))
	v193 = v189<<(uint(int32(24))%32) + v174
	goto L39
L41:
	;
	m.G0 = v10 + int32(16)
	return v326
L42:
	;
	v326 = int32(0)
	goto L41
L43:
	;
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v293 = v285
	v294 = v288
	goto L44
L44:
	;
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v294)))
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v294)+8))
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	if v300^v292|(v302^v303) != int64(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v294)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v315
	v326 = v294
	goto L41
L46:
	;
	v310 = (v293 + int32(1)) & v284
	v313 = v283 + v310*int32(40)
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+20)))
	if v314 != 0 {
		v293 = v310
		v294 = v313
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	goto L45
L49:
	;
	goto L42
}
func F_BlockRefTableRead(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if int32(0) < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = l0 + int32(8)
	v17 = l1
	v18 = l2
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(32)
	return
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[0])))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[1])))
	if v24 < v23 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	if int32(0) < v75 {
		v17 = v74
		v18 = v75
		goto L4
	} else {
		goto L26
	}
L7:
	;
	v26 = v23 - v24
	if v18 < v26 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(_a_F_BlockRefTableRead_0)) <= base.Ui32(v18) {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v28 = v18
	goto L12
L11:
	;
	v28 = v26
	goto L12
L12:
	;
	if v28 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	base.MemoryCopy(m, v17, v24+v15, v28)
	goto L15
L14:
	;
	goto L15
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[2])))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[1])))
	v34 = m.Env.Pgmem_crc32c(m, v31, v15+v32, v28)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[2]))) = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[1])))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[1]))) = v36 + v28
	v74 = v17 + v28
	v75 = v18 - v28
	goto L6
L16:
	;
	v45 = m.T0[v42].(func(*base.Module, int32, int32, int32) int32)(m, v41, v17, v18)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v60 = m.T0[v42].(func(*base.Module, int32, int32, int32) int32)(m, v41, v15, int32(_a_F_BlockRefTableRead_0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L19
	} else {
		goto L23
	}
L19:
	;
	return
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[2])))
	v48 = m.Env.Pgmem_crc32c(m, v47, v17, v45)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[2]))) = v48
	v50 = v18 - v45
	v51 = v17 + v45
	if v45 != 0 {
		v74 = v51
		v75 = v50
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[3])))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[4])))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[5])))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v54
	m.T0[v53].(func(*base.Module, int32, int32, int32))(m, v52, int32(_a_F_BlockRefTableRead_1), v10)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v74 = v51
	v75 = v50
	goto L6
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[1]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[0]))) = v60
	if v60 != 0 {
		v74 = v17
		v75 = v18
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[3])))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[4])))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableRead[5])))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v67
	m.T0[v66].(func(*base.Module, int32, int32, int32))(m, v65, int32(_a_F_BlockRefTableRead_1), v10+int32(16))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	v74 = v17
	v75 = v18
	goto L6
L26:
	;
	goto L5
}
func F_BlockRefTableSetLimitBlock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l2
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v16
	v29 = F_blockreftable_insert(m, v21, v14+int32(8), v14+int32(31))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return
	} else {
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+31)))
		if v31 == int32(0) {
			v34 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v34
			*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = l3
			*(*int64)(unsafe.Add(mBase, uint32(v29)+32)) = v34
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
			if base.Ui32(v39) <= base.Ui32(l3) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = l3
				v43 = int32(base.Ui32(l3) >> (uint(int32(16)) % 32))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
				if base.Ui32(v44) <= base.Ui32(v43) {
				} else {
					v47 = v43 + int32(1)
					if base.Ui32(v47) < base.Ui32(v44) {
						v51 = v47
						for {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
							v61 = int32(1)
							v64 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v60+v51<<(uint(v61)%32)))) = uint16(v64)
							v67 = v51 + v61
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
							if base.Ui32(v67) < base.Ui32(v68) {
								v51 = v67
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
					v84 = v81 + v43<<(uint(int32(1))%32)
					v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84))))
					if v85 != 0 {
						v87 = l3 & int32(_a_F_BlockRefTableSetLimitBlock_0)
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v43<<(uint(int32(2))%32))))
						if v85 == int32(_a_F_BlockRefTableSetLimitBlock_1) {
							if l3&int32(1) != 0 {
								v147 = v92 + int32(base.Ui32(v87)>>(uint(int32(3))%32))&int32(_a_F_BlockRefTableSetLimitBlock_2)
								v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147))))
								v153 = v148 & base.I32_rotl(int32(-2), l3&int32(15))
								*(*uint16)(unsafe.Add(mBase, uint32(v147))) = uint16(v153)
								v157 = v87 + int32(1)
							} else {
								v157 = v87
							}
							if v87 == int32(_a_F_BlockRefTableSetLimitBlock_0) {
							} else {
								v162 = v157
								for {
									v172 = int32(3)
									v174 = int32(536870910)
									v176 = v92 + int32(base.Ui32(v162)>>(uint(v172)%32))&v174
									v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176))))
									v178 = int32(-2)
									v179 = int32(15)
									v182 = v177 & base.I32_rotl(v178, v162&v179)
									*(*uint16)(unsafe.Add(mBase, uint32(v176))) = uint16(v182)
									v185 = v162 + int32(1)
									v190 = v92 + int32(base.Ui32(v185)>>(uint(v172)%32))&v174
									v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190))))
									v196 = v191 & base.I32_rotl(v178, v185&v179)
									*(*uint16)(unsafe.Add(mBase, uint32(v190))) = uint16(v196)
									v199 = v162 + int32(2)
									if v199 != int32(_a_F_BlockRefTableSetLimitBlock_3) {
										v162 = v199
										continue
									} else {
										break
									}
									break
								}
							}
						} else {
							v95 = int32(0)
							v99 = v95
							v102 = v95
							v105 = v81
							for {
								v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92+v99<<(uint(int32(1))%32)))))
								if base.Ui32(v111) < base.Ui32(v87) {
									v113 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v92+v102<<(uint(v113)%32)))) = uint16(v111)
									v117 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
									v120 = v102 + v113
									v121 = v117
								} else {
									v120 = v102
									v121 = v105
								}
								v122 = int32(1)
								v123 = v99 + v122
								v126 = v121 + v43<<(uint(v122)%32)
								v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126))))
								if base.Ui32(v123) < base.Ui32(v127) {
									v99 = v123
									v102 = v120
									v105 = v121
									continue
								} else {
									break
								}
								break
							}
							v134 = v120
							v138 = v126
							*(*uint16)(unsafe.Add(mBase, uint32(v138))) = uint16(v134)
						}
					} else {
						v134 = v85
						v138 = v84
						*(*uint16)(unsafe.Add(mBase, uint32(v138))) = uint16(v134)
					}
				}
			}
		}
		m.G0 = v14 + int32(48)
		return
	}
}
