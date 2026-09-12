package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FreePageBtreeCleanup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
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
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
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
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = l0 - v9
	v12 = v10 + int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v229 != 0 {
		goto L68
	} else {
		goto L69
	}
L2:
	;
	v224 = v2
	goto L1
L3:
	;
	goto L4
L4:
	;
	v17 = v13
	goto L6
L5:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v138 = int32(1)
	v139 = l0 - v136 + v138
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v24)+19))
	v143 = v139 + v140<<(uint(int32(12))%32)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	if v144 != 0 {
		goto L43
	} else {
		goto L44
	}
L6:
	;
	v24 = v17 + v12
	v25 = int32(1)
	v26 = v24 - v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+3))
	if v27 != v25 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v125 = int32(129)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if base.Ui32(v125) <= base.Ui32(v126) {
		goto L40
	} else {
		goto L41
	}
L8:
	;
	goto L7
L9:
	;
	if v27 != int32(2) {
		v224 = v2
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v71 - int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v75 == int32(-1729435864) {
		goto L24
	} else {
		goto L25
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v32 != int32(-1729435864) {
		v224 = v2
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)+19))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+11))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+15))
	v38 = v36 + v37
	if v35 != v38+int32(1) {
		v224 = v2
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v38 != int32(base.Ui32(v26-v12)>>(uint(int32(12))%32)) {
		v224 = v2
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = int32(1)
	v49 = l0 - v46 + v48
	v52 = v49 + v36<<(uint(int32(12))%32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v53 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v58 = v49 + v53 - v48
	goto L18
L17:
	;
	v58 = int32(0)
	goto L18
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v59 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v58 == int32(0) {
		goto L8
	} else {
		goto L22
	}
L20:
	;
	v62 = v49 + v59
	if v62 == int32(1) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+7)) = v53
	goto L19
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v69
	goto L5
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v95 = int32(1)
	v96 = l0 - v93 + v95
	v99 = (v26 - v12) & int32(-4096)
	v100 = v96 + v99
	*(*int64)(unsafe.Add(mBase, uint32(v100))) = int64(8225038576)
	v103 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v100)+8)) = v103
	if v92 != 0 {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v24)+11))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v24)+15))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v82
	goto L23
L25:
	;
	goto L26
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v24)+15))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v84
	if v84 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v88 = v84 + v10
	goto L29
L28:
	;
	v88 = int32(0)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = int32(0)
	goto L23
L30:
	;
	v109 = v92 + v96 - v95
	goto L32
L31:
	;
	v109 = v103
	goto L32
L32:
	;
	if v109 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v114 = v109 - v96 + int32(1)
	goto L35
L34:
	;
	v114 = int32(0)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+12)) = v114
	v117 = v99 | int32(1)
	if v109 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v117
	goto L38
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v117
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v120 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v124 != 0 {
		v17 = v124
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v224 = v2
	goto L1
L40:
	;
	v129 = v125
	goto L42
L41:
	;
	v129 = v126
	goto L42
L42:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v129<<(uint(int32(2))%32)+l0)+32)) = v133
	goto L5
L43:
	;
	v149 = v139 + v144 - v138
	goto L45
L44:
	;
	v149 = int32(0)
	goto L45
L45:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	if v150 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if v149 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v153 = v139 + v150
	if v153 == int32(1) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+7)) = v144
	goto L46
L49:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v24)+11))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v24)+23))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v24)+15))
	v175 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v175
	v181 = v174 + v173 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v181
	v183 = int32(129)
	if base.Ui32(v183) <= base.Ui32(v181) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v149)+12)) = v158
	goto L49
L51:
	;
	goto L52
L52:
	;
	v160 = int32(129)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	if base.Ui32(v160) <= base.Ui32(v161) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v164 = v160
	goto L55
L54:
	;
	v164 = v161
	goto L55
L55:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v164<<(uint(int32(2))%32)+l0)+32)) = v168
	goto L49
L56:
	;
	v186 = v183
	goto L58
L57:
	;
	v186 = v181
	goto L58
L58:
	;
	v191 = v186<<(uint(int32(2))%32) + l0 + int32(32)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v195 = int32(1)
	v196 = l0 - v193 + v195
	v198 = v171 << (uint(int32(12)) % 32)
	v199 = v196 + v198
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v199))) = int32(-364896016)
	v203 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+8)) = v203
	if v192 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v209 = v192 + v196 - v195
	goto L61
L60:
	;
	v209 = v203
	goto L61
L61:
	;
	if v209 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v214 = v209 - v196 + int32(1)
	goto L64
L63:
	;
	v214 = int32(0)
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+12)) = v214
	v217 = v198 | int32(1)
	if v209 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209)+8)) = v217
	goto L67
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v217
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v224 = v220
	goto L1
L68:
	;
	v233 = v224
	v234 = v229
	goto L71
L69:
	;
	v316 = v224
	goto L70
L70:
	;
	return v316
L71:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v240 = int32(1)
	v241 = l0 - v238 + v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v242 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v316 = v311
	goto L70
L73:
	;
	v247 = v241 + v242 - v240
	goto L75
L74:
	;
	v247 = int32(0)
	goto L75
L75:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	if v248 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v253 = v241 + v248 - int32(1)
	goto L78
L77:
	;
	v253 = int32(0)
	goto L78
L78:
	;
	if v253 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v247)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v257 = v256
	goto L81
L80:
	;
	v257 = v234
	goto L81
L81:
	;
	v258 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v257 - v258
	if v253 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v265 = v253 - v241 + v258
	goto L84
L83:
	;
	v265 = int32(0)
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v265
	v267 = v247 - v12
	v270 = int32(1)
	v272 = F_FreePageManagerPutInternal(m, l0, int32(base.Ui32(v267)>>(uint(int32(12))%32)), v270, v270)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	return int32(0)
L86:
	;
	if v272 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v281 = int32(1)
	v282 = l0 - v279 + v281
	v284 = v267 & int32(-4096)
	v285 = v282 + v284
	*(*int64)(unsafe.Add(mBase, uint32(v285))) = int64(8225038576)
	v288 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v285)+8)) = v288
	if v278 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	if base.Ui32(v233) < base.Ui32(v272) {
		goto L99
	} else {
		goto L100
	}
L90:
	;
	v294 = v278 + v282 - v281
	goto L92
L91:
	;
	v294 = v288
	goto L92
L92:
	;
	if v294 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v299 = v294 - v282 + int32(1)
	goto L95
L94:
	;
	v299 = int32(0)
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285)+12)) = v299
	v302 = v284 | int32(1)
	if v294 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+8)) = v302
	goto L98
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v302
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v305 + int32(1)
	return v233
L99:
	;
	v311 = v272
	goto L101
L100:
	;
	v311 = v233
	goto L101
L101:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v312 != 0 {
		v233 = v311
		v234 = v312
		goto L71
	} else {
		goto L102
	}
L102:
	;
	goto L72
}
func F_FreePageBtreeRemove(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
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
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v14 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_FreePageBtreeRemovePage(m, l0, l1)
	mBase = m.M
	return
L2:
	;
	goto L3
L3:
	;
	v19 = v14 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v19
	if base.Ui32(v19) <= base.Ui32(l2) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_FreePageBtreeConsolidate(m, l0, l1)
	mBase = m.M
	return
L5:
	;
	v23 = l1 + int32(12)
	v24 = int32(3)
	v26 = v23 + l2<<(uint(v24)%32)
	v28 = v26 + int32(8)
	v31 = (v19 - l2) << (uint(v24) % 32)
	if v26 == v28 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l2 != 0 {
		goto L4
	} else {
		goto L52
	}
L7:
	;
	goto L6
L8:
	;
	v35 = v26 + v31
	if base.Ui32(v28-v35) <= base.Ui32(int32(0)-v31<<(uint(int32(1))%32)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v42 = F___memcpy(m, v26, v28, v31)
	mBase = m.M
	goto L6
L10:
	;
	goto L11
L11:
	;
	v45 = (v26 ^ v28) & int32(3)
	if base.Ui32(v26) < base.Ui32(v28) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v147 == int32(0) {
		goto L7
	} else {
		goto L48
	}
L13:
	;
	if base.Ui32(v125) <= base.Ui32(int32(3)) {
		v146 = v124
		v147 = v125
		v148 = v126
		goto L12
	} else {
		goto L44
	}
L14:
	;
	if v45 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if v45 != 0 {
		v107 = v31
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v146 = v28
	v147 = v31
	v148 = v26
	goto L12
L18:
	;
	goto L19
L19:
	;
	if v26&int32(3) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v124 = v28
	v125 = v31
	v126 = v26
	goto L13
L21:
	;
	goto L22
L22:
	;
	v52 = v28
	v53 = v31
	v54 = v26
	goto L23
L23:
	;
	if v53 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L24:
	;
	v124 = v61
	v125 = v63
	v126 = v65
	goto L13
L25:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v58)
	v60 = int32(1)
	v61 = v52 + v60
	v63 = v53 - v60
	v65 = v54 + v60
	if v65&int32(3) != 0 {
		v52 = v61
		v53 = v63
		v54 = v65
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	if v107 == int32(0) {
		goto L7
	} else {
		goto L40
	}
L28:
	;
	if v35&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v72 = v31
	goto L32
L30:
	;
	v87 = v31
	goto L31
L31:
	;
	if base.Ui32(v87) <= base.Ui32(int32(3)) {
		v107 = v87
		goto L27
	} else {
		goto L36
	}
L32:
	;
	if v72 == int32(0) {
		goto L7
	} else {
		goto L34
	}
L33:
	;
	v87 = v78
	goto L31
L34:
	;
	v78 = v72 - int32(1)
	v79 = v26 + v78
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v78))))
	*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v81)
	if v79&int32(3) != 0 {
		v72 = v78
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v94 = v87
	goto L37
L37:
	;
	v98 = v94 - int32(4)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v28+v98)))
	*(*int32)(unsafe.Add(mBase, uint32(v26+v98))) = v101
	if base.Ui32(int32(3)) < base.Ui32(v98) {
		v94 = v98
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v107 = v98
	goto L27
L39:
	;
	goto L38
L40:
	;
	v114 = v107
	goto L41
L41:
	;
	v118 = v114 - int32(1)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v118))))
	*(*uint8)(unsafe.Add(mBase, uint32(v26+v118))) = uint8(v121)
	if v118 != 0 {
		v114 = v118
		goto L41
	} else {
		goto L43
	}
L42:
	;
	goto L7
L43:
	;
	goto L42
L44:
	;
	v131 = v124
	v132 = v125
	v133 = v126
	goto L45
L45:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v135
	v137 = int32(4)
	v138 = v131 + v137
	v140 = v133 + v137
	v142 = v132 - v137
	if base.Ui32(int32(3)) < base.Ui32(v142) {
		v131 = v138
		v132 = v142
		v133 = v140
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v146 = v138
	v147 = v142
	v148 = v140
	goto L12
L47:
	;
	goto L46
L48:
	;
	v153 = v146
	v154 = v147
	v155 = v148
	goto L49
L49:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v157)
	v159 = int32(1)
	v164 = v154 - v159
	if v164 != 0 {
		v153 = v153 + v159
		v154 = v164
		v155 = v155 + v159
		goto L49
	} else {
		goto L51
	}
L50:
	;
	goto L7
L51:
	;
	goto L50
L52:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v177 = l0 - v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v186 = l1
	goto L53
L53:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	if v192 == int32(0) {
		goto L4
	} else {
		goto L55
	}
L54:
	;
	goto L4
L55:
	;
	v195 = v192 + v177
	if v195 == int32(0) {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	v199 = v195 + int32(12)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v204 = int32(0)
	v205 = v201
	goto L57
L57:
	;
	if base.Ui32(v205) <= base.Ui32(v204) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	if base.Ui32(v231) < base.Ui32(v201) {
		goto L70
	} else {
		goto L71
	}
L59:
	;
	goto L58
L60:
	;
	v231 = v204
	goto L59
L61:
	;
	goto L62
L62:
	;
	v217 = int32(1)
	v218 = int32(base.Ui32(v204+v205) >> (uint(v217) % 32))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v199+v218<<(uint(int32(3))%32))))
	v225 = base.B2i32(base.Ui32(v178) < base.Ui32(v224))
	if base.Ui32(v178) < base.Ui32(v224) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v226 = v204
	goto L65
L64:
	;
	v226 = v218 + v217
	goto L65
L65:
	;
	if base.Ui32(v178) < base.Ui32(v224) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v227 = v218
	goto L68
L67:
	;
	v227 = v205
	goto L68
L68:
	;
	if v178 != v224 {
		v204 = v226
		v205 = v227
		goto L57
	} else {
		goto L69
	}
L69:
	;
	v231 = v218
	goto L59
L70:
	;
	v237 = int32(0)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v195+v231<<(uint(int32(3))%32))+16))
	if v241 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v248 = int32(-1)
	goto L72
L72:
	;
	v249 = v248 + v231
	*(*int32)(unsafe.Add(mBase, uint32(v199+v249<<(uint(int32(3))%32)))) = v178
	if v249 == int32(0) {
		v186 = v195
		goto L53
	} else {
		goto L79
	}
L73:
	;
	v244 = v177 + v241
	goto L75
L74:
	;
	v244 = v237
	goto L75
L75:
	;
	if v244 != v186 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v246 = int32(-1)
	goto L78
L77:
	;
	v246 = v237
	goto L78
L78:
	;
	v248 = v246
	goto L72
L79:
	;
	goto L54
}
func F_GetPageWithFreeSpace(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if base.Ui32(int32(8161)) <= base.Ui32(l1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
			F_errmsg_internal(m, int32(37378), v6)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(499271), int32(438), int32(112435))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		if l1 != 0 {
			v30 = int32(base.Ui32(l1+int32(31)) >> (uint(int32(5)) % 32))
		} else {
			v30 = int32(1)
		}
		v33 = F_fsm_search(m, l0, v30&int32(255))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			m.G0 = v6 + int32(16)
			return v33
		}
	}
}
func F_PageGetFreeSpace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = int32(4)
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v5 = v3 - v4
	if v5 <= v2 {
		v8 = v2
	} else {
		v8 = v5
	}
	return v8 - int32(4)
}
func F_PageIndexMultiDelete(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	v4 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(4128)
	m.G0 = v23
	if l2 <= int32(2) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L12
	} else {
		goto L58
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L12
	} else {
		goto L55
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L12
	} else {
		goto L51
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L12
	} else {
		goto L47
	}
L5:
	;
	m.G0 = v23 + int32(4128)
	return
L6:
	;
	v28 = l2 - int32(1)
	if v28 < int32(0) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v61) < base.Ui32(int32(24)) {
		goto L4
	} else {
		goto L15
	}
L9:
	;
	v37 = v28
	goto L10
L10:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v37<<(uint(int32(1))%32)))))
	F_PageIndexTupleDelete(m, l0, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L5
L12:
	;
	return
L13:
	;
	if v37 != 0 {
		v37 = v37 - int32(1)
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	if base.Ui32(v60) < base.Ui32(v61) {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(v59) < base.Ui32(v60) {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(int32(8192)) < base.Ui32(v59) {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if (v59+int32(7))&int32(32760) != v59 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v73 = int32(1)
	if base.Ui32(v61) < base.Ui32(int32(25)) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if l2 != v180 {
		goto L2
	} else {
		goto L38
	}
L21:
	;
	v179 = v4
	v180 = v4
	v181 = v4
	v185 = v73
	goto L20
L22:
	;
	goto L23
L23:
	;
	v81 = int32(base.Ui32(v61+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v81 == int32(0) {
		v179 = v4
		v180 = v4
		v181 = v4
		v185 = v73
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v95 = v23 + int32(1680)
	v96 = v4
	v97 = v4
	v98 = v4
	v102 = v73
	v103 = v59
	v104 = int32(1)
	goto L25
L25:
	;
	v110 = v104 & int32(65535)
	v115 = v110<<(uint(int32(2))%32) + (l0 + int32(24)) - int32(4)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v118 = int32(base.Ui32(v116) >> (uint(int32(17)) % 32))
	v120 = v116 & int32(32767)
	if base.Ui32(v120) < base.Ui32(v60) {
		goto L3
	} else {
		goto L27
	}
L26:
	;
	v179 = v161
	v180 = v162
	v181 = v163
	v185 = v165
	goto L20
L27:
	;
	if base.Ui32(v59) < base.Ui32(v120+v118) {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	if v120 != (v120+int32(7))&int32(65528) {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	if l2 <= v97 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v168 = v104 + int32(1)
	if base.Ui32(v168&int32(65535)) <= base.Ui32(v81) {
		v95 = v160
		v96 = v161
		v97 = v162
		v98 = v163
		v102 = v165
		v103 = v166
		v104 = v168
		goto L25
	} else {
		goto L37
	}
L31:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v95)+2)) = uint16(v120)
	*(*uint16)(unsafe.Add(mBase, uint32(v95))) = uint16(v96)
	v142 = (v118 + int32(7)) & int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(v95)+4)) = uint16(v142)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(48)+v96<<(uint(int32(2))%32)))) = v149
	if v103 < v120 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v97<<(uint(int32(1))%32)))))
	if v110 != v133 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v160 = v95
	v161 = v96
	v162 = v97 + int32(1)
	v163 = v98
	v165 = v102
	v166 = v103
	goto L30
L34:
	;
	v159 = v103
	goto L36
L35:
	;
	v159 = v120
	goto L36
L36:
	;
	v160 = v95 + int32(6)
	v161 = v96 + int32(1)
	v162 = v97
	v163 = v98 + v142
	v165 = base.B2i32(v120 < v103) & v102
	v166 = v159
	goto L30
L37:
	;
	goto L26
L38:
	;
	v193 = v59 - v61
	if base.Ui32(v193) < base.Ui32(v181) {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v200 = v179 << (uint(int32(2)) % 32)
	if v200 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v204 = v200 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v204)
	if int32(0) < v179 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v201 = F__emscripten_memcpy_bulkmem(m, l0+int32(24), v23+int32(48), v200)
	mBase = m.M
	goto L43
L42:
	;
	goto L43
L43:
	;
	goto L40
L44:
	;
	F_compactify_tuples(m, v23+int32(1680), v179, l0, v185)
	mBase = m.M
	goto L5
L45:
	;
	goto L46
L46:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v59)
	goto L5
L47:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v61
	F_errmsg(m, int32(57007), v23)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(499071), int32(1208), int32(351123))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v120
	F_errmsg(m, int32(57069), v23+int32(32))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(499071), int32(1233), int32(351123))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errmsg_internal(m, int32(456031), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L12
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(499071), int32(1260), int32(351123))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v181
	F_errmsg(m, int32(52858), v23+int32(16))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(499071), int32(1266), int32(351123))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PageInit(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	if l0&int32(3) != 0 {
		v30 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(0)), l1)
		mBase = m.M
	} else {
		if base.Ui32(int32(1024)) < base.Ui32(l1) {
			v30 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(0)), l1)
			mBase = m.M
		} else {
			if l1&int32(3) != 0 {
				v30 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(0)), l1)
				mBase = m.M
			} else {
				v12 = l0 + l1
				if base.Ui32(v12) <= base.Ui32(l0) {
				} else {
					v18 = l0 + int32(4)
					if base.Ui32(v18) < base.Ui32(v12) {
						v20 = v12
					} else {
						v20 = v18
					}
					v27 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(0)), (l0^int32(-1)+v20)&int32(-4)+int32(4))
					mBase = m.M
				}
			}
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(1572864)
	v36 = l1 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v36)
	v42 = l1 - (l2+int32(7))&int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v42)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v42)
	return
}
func F_PageIsVerified(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v234 int32
	_ = v234
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v15)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v17 = int32(1)
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	if v18 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v13 + int32(16)
	return v234
L5:
	;
	if l0&int32(3) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v64 = v5
	v65 = v5
	v66 = v17
	goto L5
L7:
	;
	goto L8
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+252))
	goto L10
L9:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	if base.Ui32(int32(7)) < base.Ui32(v40) {
		v64 = v38
		v65 = v39
		v66 = v17
		goto L5
	} else {
		goto L14
	}
L10:
	;
	if base.B2i32(v23 != int32(0)) == int32(0) {
		v38 = v5
		v39 = v5
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v28 = F_pg_checksum_page(m, l0, l1)
	mBase = m.M
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	v30 = base.B2i32(v28 != v29)
	if l3 == int32(0) {
		v38 = v30
		v39 = v28
		goto L9
	} else {
		goto L12
	}
L12:
	;
	if v29 == v28 {
		v38 = v30
		v39 = v28
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v34 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v34)
	v38 = v34
	v39 = v28
	goto L9
L14:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v43) < base.Ui32(v44) {
		v64 = v38
		v65 = v39
		v66 = v17
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.Ui32(v46) < base.Ui32(v43) {
		v64 = v38
		v65 = v39
		v66 = v17
		goto L5
	} else {
		goto L16
	}
L16:
	;
	if base.Ui32(int32(8192)) < base.Ui32(v46) {
		v64 = v38
		v65 = v39
		v66 = v17
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v50 = int32(1)
	v54 = (v46 + int32(7)) & int32(32760)
	v55 = base.B2i32(v54 != v46)
	if v38|v55 != v50 {
		v234 = v50
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v64 = v38 | base.B2i32(v46 == v54)
	v65 = v39
	v66 = v55
	goto L5
L19:
	;
	if v64 != 0 {
		goto L46
	} else {
		goto L47
	}
L20:
	;
	v94 = l0 + (int32(0)-l0)&int32(3)
	v96 = l0 - int32(-8192)
	v98 = v96 & int32(-4)
	v100 = v98 - int32(28)
	if base.Ui32(v94) < base.Ui32(v100) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v71 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	if (l0+int32(1))&int32(3) == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v78 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	if (l0+int32(2))&int32(3) == int32(0) {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if v85 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	if (l0-int32(1))&int32(3) != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	goto L20
L28:
	;
	v105 = v94
	goto L31
L29:
	;
	v133 = v94
	goto L30
L30:
	;
	if base.Ui32(v133) < base.Ui32(v98) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+28))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v112|(v113|(v114|(v115|(v116|(v117|(v118|v119)))))) != 0 {
		goto L19
	} else {
		goto L33
	}
L32:
	;
	v133 = v128
	goto L30
L33:
	;
	v128 = v105 + int32(32)
	if base.Ui32(v128) < base.Ui32(v100) {
		v105 = v128
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v144 = v133
	goto L38
L36:
	;
	v158 = v133
	goto L37
L37:
	;
	v169 = v158
	goto L42
L38:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	if v151 != 0 {
		goto L19
	} else {
		goto L40
	}
L39:
	;
	v158 = v153
	goto L37
L40:
	;
	v153 = v144 + int32(4)
	if base.Ui32(v153) < base.Ui32(v98) {
		v144 = v153
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	if base.Ui32(v96) <= base.Ui32(v169) {
		v234 = int32(1)
		goto L4
	} else {
		goto L44
	}
L43:
	;
	goto L19
L44:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v177 == int32(0) {
		v169 = v169 + int32(1)
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	if l2&int32(3) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v234 = int32(0)
	goto L4
L49:
	;
	if base.B2i32(v66 == int32(0))&int32(base.Ui32(l2)>>(uint(int32(2))%32)) != 0 {
		v234 = int32(1)
		goto L4
	} else {
		goto L60
	}
L50:
	;
	if l2&int32(1) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v200 = int32(19)
	goto L53
L52:
	;
	v200 = int32(15)
	goto L53
L53:
	;
	v202 = F_errstart(m, v200, int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	return int32(0)
L55:
	;
	if v202 == int32(0) {
		goto L49
	} else {
		goto L56
	}
L56:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v65
	F_errmsg(m, int32(54993), v13)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(499071), int32(155), int32(456344))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	goto L49
L60:
	;
	goto L48
}
func F_PageTruncateLinePointerArray(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v2 = int32(0)
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v8) < base.Ui32(int32(25)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v76 = v74 & int32(65534)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v76)
	return
L2:
	;
	v16 = int32(base.Ui32(v8+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v16 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = v16
	v23 = v2
	v26 = v2
	goto L5
L4:
	;
	if int32(0) < v52 {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22&int32(65535)<<(uint(int32(2))%32)+(l0+int32(24))-int32(4))))
	v37 = v35 & int32(98304)
	if v22 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v52 = v46
	v54 = int32(0)
	goto L4
L7:
	;
	v49 = v22 - int32(1)
	if v49 != 0 {
		v22 = v49
		v23 = v46
		v26 = v47
		goto L5
	} else {
		goto L12
	}
L8:
	;
	if v37 != 0 {
		v46 = v23
		v47 = v26
		goto L7
	} else {
		goto L11
	}
L9:
	;
	if v26 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v40 = int32(0)
	v46 = v23 + base.B2i32(v37 == v40)
	v47 = base.B2i32(v37 != v40)
	goto L7
L11:
	;
	v52 = v23
	v54 = int32(1)
	goto L4
L12:
	;
	goto L6
L13:
	;
	v59 = v8 - v52<<(uint(int32(2))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v59)
	goto L15
L14:
	;
	goto L15
L15:
	;
	if v54 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v65 = v63 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v65)
	return
}
func F_UnlockPage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(16973824)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v9
	v17 = F_LockRelease(m, v5, int32(7), v2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
