package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tbm_add_tuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
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
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v350 int32
	_ = v350
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	if v5 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = l0 + int32(40)
	v32 = int32(-1)
	v33 = v5
	v42 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v21 + int32(16)
	return
L4:
	;
	v48 = l1 + v42*int32(6)
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+4)))
	if base.Ui32(int32(_a_F_tbm_add_tuples_0)) < base.Ui32((v49-int32(292))&int32(_a_F_tbm_add_tuples_1)) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L3
L6:
	;
	v350 = v42 + int32(1)
	if v350 != l2 {
		v32 = v335
		v33 = v336
		v42 = v350
		goto L4
	} else {
		goto L63
	}
L7:
	;
	if v206 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L8:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch v161 {
	case 0:
		goto L31
	case 1:
		goto L30
	default:
		goto L29
	}
L9:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+2)))
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48))))
	v60 = v56 | v57<<(uint(int32(16))%32)
	if v60 == v32 {
		v205 = v32
		v206 = v33
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v62 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v69 = v60 & int32(-256)
	v72 = (v69 ^ v57) * int32(-2048144789)
	v77 = (int32(base.Ui32(v72)>>(uint(int32(13))%32)) ^ v72) * int32(-1028477387)
	v81 = v67 & (int32(base.Ui32(v77)>>(uint(int32(16))%32)) ^ v77)
	v84 = v66 + v81*int32(48)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+4)))
	if v85 == int32(0) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v92 = v81
	v95 = v84
	goto L15
L15:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v69 != v106 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+5)))
	if v115 != int32(1) {
		goto L8
	} else {
		goto L21
	}
L17:
	;
	v110 = (v92 + int32(1)) & v67
	v113 = v66 + v110*int32(48)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+4)))
	if v114 != 0 {
		v92 = v110
		v95 = v113
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L8
L21:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v95+int32(base.Ui32(v56)>>(uint(int32(3))%32))&int32(28))+8))
	if int32(base.Ui32(v123)>>(uint(v56)%32))&int32(1) == int32(0) {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v335 = v60
	v336 = int32(0)
	goto L6
L23:
	;
	return
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v49
	F_errmsg_internal(m, int32(_a_F_tbm_add_tuples_2), v21)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_tbm_add_tuples_3), int32(383), int32(_a_F_tbm_add_tuples_4))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v205 = v60
	v206 = v200
	goto L7
L28:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
	v177 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v175)+4)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v175)+12)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v175)+20)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v175)+28)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v175)+36)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v175)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)) = uint8(v176)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v192 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v191 + v192
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v195 + v192
	v200 = v175
	goto L27
L29:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v172 = F_pagetable_insert(m, v169, v60, v21+int32(15))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L23
	} else {
		goto L34
	}
L30:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v164 == v60 {
		v200 = v26
		goto L27
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
	v175 = v26
	goto L28
L32:
	;
	F_tbm_create_pagetable(m, l0)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L23
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)))
	if v174 != 0 {
		v200 = v172
		goto L27
	} else {
		goto L35
	}
L35:
	;
	v175 = v172
	goto L28
L36:
	;
	v335 = v205
	v336 = int32(0)
	goto L6
L37:
	;
	goto L38
L38:
	;
	v224 = v49 - int32(1)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+5)))
	if v227 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v228 = int32(0)
	goto L41
L40:
	;
	v228 = int32(base.Ui32(v224) >> (uint(int32(5)) % 32))
	goto L41
L41:
	;
	v231 = v206 + v228<<(uint(int32(2))%32)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v233 = int32(1)
	if v227 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v236 = v233
	goto L44
L43:
	;
	v236 = v233 << (uint(v224) % 32)
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+8)) = v232 | v236
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+6)))
	v240 = v239 | l3
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+6)) = uint8(v240)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v242 <= v243 {
		v335 = v205
		v336 = v206
		goto L6
	} else {
		goto L45
	}
L45:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v249 = v247 & v248
	v254 = int32(0)
	v257 = v249
	v258 = v242
	v259 = v246
	v262 = v243
	goto L47
L46:
	;
	v320 = int32(-1)
	v322 = base.I32_div_s(v262, int32(2))
	if v258 <= v322 {
		v335 = v320
		v336 = v206
		goto L6
	} else {
		goto L59
	}
L47:
	;
	v272 = v254
	v274 = v254
	v275 = v257
	goto L49
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v298
	v335 = int32(-1)
	v336 = v206
	goto L6
L49:
	;
	if v274&int32(1) != 0 {
		goto L46
	} else {
		goto L51
	}
L50:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+5)))
	if v304 != 0 {
		v254 = v295
		v257 = v298
		goto L47
	} else {
		goto L53
	}
L51:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v289 = int32(1)
	v290 = v275 - v289
	v294 = base.B2i32(v288&(v290^v249) == int32(0))
	v295 = v294 | v272
	v298 = v288 & v290
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v259)+20))
	v300 = v275*int32(48) + v299
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+4)))
	if v301 != v289 {
		v272 = v295
		v274 = v294
		v275 = v298
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	if v305&int32(255) == int32(0) {
		v254 = v295
		v257 = v298
		goto L47
	} else {
		goto L54
	}
L54:
	;
	F_tbm_mark_page_lossy(m, l0, v305)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L23
	} else {
		goto L55
	}
L55:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v315 = base.I32_div_s(v313, int32(2))
	if v315 < v312 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v254 = v295
	v257 = v298
	v258 = v312
	v259 = v317
	v262 = v313
	goto L47
L57:
	;
	goto L58
L58:
	;
	goto L48
L59:
	;
	v324 = int32(1073741823)
	if v324 <= v258 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v327 = v324
	goto L62
L61:
	;
	v327 = v258
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v327 << (uint(int32(1)) % 32)
	v335 = v320
	v336 = v206
	goto L6
L63:
	;
	goto L5
}
func F_tbm_union_page(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
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
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v17 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return
L2:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v258 <= v259 {
		goto L1
	} else {
		goto L45
	}
L3:
	;
	v26 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v79 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(8)+v26<<(uint(int32(2))%32))))
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L2
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v44 = v37
	v45 = v38 + v26<<(uint(int32(5))%32)
	goto L11
L9:
	;
	goto L10
L10:
	;
	v75 = v26 + int32(1)
	if v75 != int32(8) {
		v26 = v75
		goto L6
	} else {
		goto L19
	}
L11:
	;
	if v44&int32(1) != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	F_tbm_mark_page_lossy(m, l0, v45)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v58 = int32(1)
	v61 = int32(base.Ui32(v44) >> (uint(v58) % 32))
	if v61 != 0 {
		v44 = v61
		v45 = v45 + v58
		goto L11
	} else {
		goto L18
	}
L16:
	;
	return
L17:
	;
	goto L15
L18:
	;
	goto L12
L19:
	;
	goto L7
L20:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch v152 {
	case 0:
		goto L35
	case 1:
		goto L34
	default:
		goto L33
	}
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v86 = v78 & int32(-256)
	v87 = int32(16)
	v91 = (v86 ^ int32(base.Ui32(v78)>>(uint(v87)%32))) * int32(-2048144789)
	v96 = (int32(base.Ui32(v91)>>(uint(int32(13))%32)) ^ v91) * int32(-1028477387)
	v100 = v84 & (int32(base.Ui32(v96)>>(uint(v87)%32)) ^ v96)
	v103 = v83 + v100*int32(48)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
	if v104 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v109 = v103
	v110 = v100
	goto L23
L23:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v86 != v119 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+5)))
	if v128 != int32(1) {
		goto L20
	} else {
		goto L29
	}
L25:
	;
	v123 = (v110 + int32(1)) & v84
	v126 = v83 + v123*int32(48)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+4)))
	if v127 != 0 {
		v109 = v126
		v110 = v123
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	goto L20
L29:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v109+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))+8))
	if int32(base.Ui32(v136)>>(uint(v78)%32))&int32(1) != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L20
L31:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+5)))
	if v195 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L32:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+4)))
	v171 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v169)+4)) = v171
	*(*int64)(unsafe.Add(mBase, uint32(v169)+12)) = v171
	*(*int64)(unsafe.Add(mBase, uint32(v169)+20)) = v171
	*(*int64)(unsafe.Add(mBase, uint32(v169)+28)) = v171
	*(*int64)(unsafe.Add(mBase, uint32(v169)+36)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v169)+44)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v78
	*(*uint8)(unsafe.Add(mBase, uint32(v169)+4)) = uint8(v170)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v186 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v185 + v186
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v189 + v186
	v193 = v169
	goto L31
L33:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v166 = F_pagetable_insert(m, v163, v78, v15+int32(15))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L16
	} else {
		goto L40
	}
L34:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v78 == v157 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
	v169 = l0 + int32(40)
	goto L32
L36:
	;
	v193 = l0 + int32(40)
	goto L31
L37:
	;
	goto L38
L38:
	;
	F_tbm_create_pagetable(m, l0)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L16
	} else {
		goto L39
	}
L39:
	;
	goto L33
L40:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)))
	if v168 != 0 {
		v193 = v166
		goto L31
	} else {
		goto L41
	}
L41:
	;
	v169 = v166
	goto L32
L42:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+8)) = v198 | v199
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+12)) = v202 | v203
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+16)) = v206 | v207
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v193)+20))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+20)) = v210 | v211
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v193)+24))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+24)) = v214 | v215
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v193)+28))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+28)) = v218 | v219
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v193)+32))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+32)) = v222 | v223
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v193)+36))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+36)) = v226 | v227
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v193)+40))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+40)) = v230 | v231
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v193)+44))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+44)) = v234 | v235
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+6)))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	v240 = v238 | v239
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+6)) = uint8(v240)
	goto L2
L43:
	;
	goto L44
L44:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+8)) = v242 | int32(1)
	goto L2
L45:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v265 = v263 & v264
	v267 = v258
	v268 = v265
	v269 = int32(0)
	v271 = v262
	v273 = v259
	goto L47
L46:
	;
	v324 = base.I32_div_s(v273, int32(2))
	if v267 <= v324 {
		goto L1
	} else {
		goto L59
	}
L47:
	;
	v280 = v268
	v281 = v269
	v282 = v269
	goto L49
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v302
	goto L1
L49:
	;
	if v282&int32(1) != 0 {
		goto L46
	} else {
		goto L51
	}
L50:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+5)))
	if v308 != 0 {
		v268 = v302
		v269 = v299
		goto L47
	} else {
		goto L53
	}
L51:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	v293 = int32(1)
	v294 = v280 - v293
	v298 = base.B2i32(v292&(v294^v265) == int32(0))
	v299 = v298 | v281
	v302 = v292 & v294
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v271)+20))
	v304 = v280*int32(48) + v303
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+4)))
	if v305 != v293 {
		v280 = v302
		v281 = v299
		v282 = v298
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	if v309&int32(255) == int32(0) {
		v268 = v302
		v269 = v299
		goto L47
	} else {
		goto L54
	}
L54:
	;
	F_tbm_mark_page_lossy(m, l0, v309)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L16
	} else {
		goto L55
	}
L55:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v319 = base.I32_div_s(v317, int32(2))
	if v319 < v316 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v267 = v316
	v268 = v302
	v269 = v299
	v271 = v321
	v273 = v317
	goto L47
L57:
	;
	goto L58
L58:
	;
	goto L48
L59:
	;
	v326 = int32(1073741823)
	if v326 <= v267 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v329 = v326
	goto L62
L61:
	;
	v329 = v267
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v329 << (uint(int32(1)) % 32)
	goto L1
}
