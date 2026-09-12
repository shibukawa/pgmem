package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RI_FKey_cascade_del(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
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
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	v16 = m.G0
	v18 = v16 - int32(624)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ri_CheckTrigger(m, l0, int32(320450), int32(3))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v30 = F_ri_FetchConstraintInfo(m, v27, v28, int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+88))
	v34 = F_table_open(m, v32, int32(3))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+620)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+616)) = v41
	v46 = v18 + int32(616)
	v49 = F_ri_FetchPreparedPlan(m, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v49 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v314 = v49
	goto L9
L8:
	;
	F_initStringInfo(m, v18+int32(600))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v315 = int32(0)
	v319 = F_ri_PerformCheck(m, v30, v46, v314, v34, v37, v36, v315, v315, int32(1), int32(8))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L53
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+119)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+68))
	v58 = F_get_namespace_name(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v60 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+336)) = uint8(v60)
	v64 = v18 + int32(336)
	v65 = v58
	goto L12
L12:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v79 != int32(34) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v96 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v64)+1)) = uint16(v96)
	v99 = v18 + int32(336)
	v100 = F_strlen(m, v99)
	mBase = m.M
	v103 = v100 + v99
	v104 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v104)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v108 = v103 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v96)
	v113 = v106 + int32(4)
	v114 = v108
	goto L20
L14:
	;
	goto L13
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v92))) = uint8(v91)
	v64 = v92
	v65 = v65 + int32(1)
	goto L12
L16:
	;
	if v79 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v86 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+1)) = uint8(v86)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v91 = v88
	v92 = v64 + int32(2)
	goto L15
L19:
	;
	v91 = v79
	v92 = v64 + int32(1)
	goto L15
L20:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v128 != int32(34) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v145 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v114)+1)) = uint16(v145)
	if v56&int32(255) == int32(112) {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	goto L21
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v140)
	v113 = v113 + int32(1)
	v114 = v141
	goto L20
L24:
	;
	if v128 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v135 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)) = uint8(v135)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	v140 = v137
	v141 = v114 + int32(2)
	goto L23
L27:
	;
	v140 = v128
	v141 = v114 + int32(1)
	goto L23
L28:
	;
	v153 = int32(785690)
	goto L30
L29:
	;
	v153 = int32(772114)
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v18 + int32(336)
	F_appendStringInfo(m, v18+int32(600), int32(184716), v18+int32(32))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v30)+168))
	if int32(0) < v165 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v182 = int32(0)
	v184 = int32(560790)
	goto L35
L33:
	;
	v277 = v165
	goto L34
L34:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v18)+600))
	v297 = F_ri_PlanCheck(m, v292, v277, v18+int32(48), v18+int32(616), v34, v37)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L52
	}
L35:
	;
	v192 = v182 << (uint(int32(1)) % 32)
	v194 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30+int32(172)+v192))))
	v195 = F_attnumTypeId(m, v37, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	v277 = v275
	goto L34
L37:
	;
	v197 = v192 + (v30 + int32(236))
	v198 = int32(*(*int16)(unsafe.Add(mBase, uint32(v197))))
	v199 = F_attnumTypeId(m, v34, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v201 = int32(*(*int16)(unsafe.Add(mBase, uint32(v197))))
	v202 = F_attnumAttName(m, v34, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v204 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+192)) = uint8(v204)
	v208 = v18 + int32(192)
	v209 = v202
	goto L40
L40:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v223 != int32(34) {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v240 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v208)+1)) = uint16(v240)
	v243 = v182 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v243
	v250 = F_pg_sprintf(m, v18+int32(176), int32(485008), v18+int32(16))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L48
	}
L42:
	;
	goto L41
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v236))) = uint8(v235)
	v208 = v236
	v209 = v209 + int32(1)
	goto L40
L44:
	;
	if v223 == int32(0) {
		goto L42
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v230 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)) = uint8(v230)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	v235 = v232
	v236 = v208 + int32(2)
	goto L43
L47:
	;
	v235 = v223
	v236 = v208 + int32(1)
	goto L43
L48:
	;
	v253 = v182 << (uint(int32(2)) % 32)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v30+int32(300)+v253)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v184
	F_appendStringInfo(m, v18+int32(600), int32(766605), v18)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_generate_operator_clause(m, v18+int32(600), v18+int32(176), v195, v255, v18+int32(192), v199)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(48)+v253))) = v195
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v30)+168))
	if v243 < v275 {
		v182 = v243
		v184 = int32(564022)
		goto L35
	} else {
		goto L51
	}
L51:
	;
	goto L36
L52:
	;
	v314 = v297
	goto L9
L53:
	;
	v321 = F_SPI_finish(m)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v321 != int32(2) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	F_sequence_close(m, v34, int32(3))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L61
	}
L58:
	;
	F_errmsg_internal(m, int32(471975), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(513194), int32(1003), int32(320450))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	m.G0 = v18 + int32(624)
	return int32(0)
}
func F_RI_FKey_cascade_upd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v74 int32
	_ = v74
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
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
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v317 int32
	_ = v317
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	v19 = m.G0
	v21 = v19 - int32(784)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ri_CheckTrigger(m, l0, int32(438979), int32(2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v33 = F_ri_FetchConstraintInfo(m, v30, v31, int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+88))
	v37 = F_table_open(m, v35, int32(3))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+780)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+776)) = v45
	v50 = v21 + int32(776)
	v53 = F_ri_FetchPreparedPlan(m, v50)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v53 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v371 = v53
	goto L9
L8:
	;
	F_initStringInfo(m, v21+int32(760))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v375 = F_ri_PerformCheck(m, v33, v50, v371, v37, v41, v39, v40, int32(0), int32(1), int32(9))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L56
	}
L10:
	;
	F_initStringInfo(m, v21+int32(744))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v37)+48))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+119)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+68))
	v66 = F_get_namespace_name(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v68 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+480)) = uint8(v68)
	v72 = v21 + int32(480)
	v74 = v66
	goto L13
L13:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v90 != int32(34) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v107 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+1)) = uint16(v107)
	v110 = v21 + int32(480)
	v111 = F_strlen(m, v110)
	mBase = m.M
	v114 = v111 + v110
	v115 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v115)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v37)+48))
	v119 = v114 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v119))) = uint8(v107)
	v124 = v117 + int32(4)
	v126 = v119
	goto L21
L15:
	;
	goto L14
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v102)
	v72 = v103
	v74 = v74 + int32(1)
	goto L13
L17:
	;
	if v90 == int32(0) {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v97 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)) = uint8(v97)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	v102 = v99
	v103 = v72 + int32(2)
	goto L16
L20:
	;
	v102 = v90
	v103 = v72 + int32(1)
	goto L16
L21:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v142 != int32(34) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v159 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v126)+1)) = uint16(v159)
	v161 = int32(785690)
	if v64&int32(255) == int32(112) {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v154)
	v124 = v124 + int32(1)
	v126 = v155
	goto L21
L25:
	;
	if v142 == int32(0) {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v149 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)) = uint8(v149)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v154 = v151
	v155 = v126 + int32(2)
	goto L24
L28:
	;
	v154 = v142
	v155 = v126 + int32(1)
	goto L24
L29:
	;
	v168 = v161
	goto L31
L30:
	;
	v168 = int32(772114)
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v21 + int32(480)
	F_appendStringInfo(m, v21+int32(760), int32(542341), v21+int32(48))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v33)+168))
	if int32(0) < v180 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v196 = int32(0)
	v199 = v180
	v201 = v161
	v202 = int32(560790)
	goto L36
L34:
	;
	goto L35
L35:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v21)+744))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v21)+748))
	F_appendBinaryStringInfo(m, v21+int32(760), v339, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L54
	}
L36:
	;
	v210 = v196 << (uint(int32(1)) % 32)
	v212 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33+int32(172)+v210))))
	v213 = F_attnumTypeId(m, v41, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	v215 = v210 + (v33 + int32(236))
	v216 = int32(*(*int16)(unsafe.Add(mBase, uint32(v215))))
	v217 = F_attnumTypeId(m, v37, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v219 = int32(*(*int16)(unsafe.Add(mBase, uint32(v215))))
	v220 = F_attnumAttName(m, v37, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v222 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+336)) = uint8(v222)
	v226 = v21 + int32(336)
	v228 = v220
	goto L41
L41:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	if v244 != int32(34) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v261 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v226)+1)) = uint16(v261)
	v264 = v196 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v21 + int32(336)
	F_appendStringInfo(m, v21+int32(760), int32(485000), v21+int32(32))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L49
	}
L43:
	;
	goto L42
L44:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v257))) = uint8(v256)
	v226 = v257
	v228 = v228 + int32(1)
	goto L41
L45:
	;
	if v244 == int32(0) {
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v251 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v226)+1)) = uint8(v251)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	v256 = v253
	v257 = v226 + int32(2)
	goto L44
L48:
	;
	v256 = v244
	v257 = v226 + int32(1)
	goto L44
L49:
	;
	v278 = v199 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v278
	v285 = F_pg_sprintf(m, v21+int32(320), int32(485008), v21+int32(16))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v288 = v196 << (uint(int32(2)) % 32)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v33+int32(300)+v288)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v202
	F_appendStringInfo(m, v21+int32(744), int32(766605), v21)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_generate_operator_clause(m, v21+int32(744), v21+int32(320), v213, v290, v21+int32(336), v217)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v306 = v21 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v306+v288))) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v306+v199<<(uint(int32(2))%32)))) = v213
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v33)+168))
	if v264 < v317 {
		v196 = v264
		v199 = v278
		v201 = int32(694808)
		v202 = int32(564022)
		goto L36
	} else {
		goto L53
	}
L53:
	;
	goto L37
L54:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v21)+760))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v33)+168))
	v351 = F_ri_PlanCheck(m, v343, v344<<(uint(int32(1))%32), v21-int32(-64), v21+int32(776), v37, v41)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v371 = v351
	goto L9
L56:
	;
	v377 = F_SPI_finish(m)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v377 != int32(2) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	F_sequence_close(m, v37, int32(3))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L64
	}
L61:
	;
	F_errmsg_internal(m, int32(471975), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(513194), int32(1120), int32(438979))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	m.G0 = v21 + int32(784)
	return int32(0)
}
func F_RI_FKey_check(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
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
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v310 int32
	_ = v310
	var v321 int32
	_ = v321
	var v337 int32
	_ = v337
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v500 int32
	_ = v500
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	v16 = m.G0
	v18 = v16 - int32(688)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = F_ri_FetchConstraintInfo(m, v20, v21, int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v28&int32(3) == int32(2) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v18 + int32(688)
	return
L4:
	;
	v33 = int32(28)
	goto L6
L5:
	;
	v33 = int32(24)
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0+v33)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v25)+188))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+72))
	v39 = m.T0[v38].(func(*base.Module, int32, int32, int32) int32)(m, v25, v35, int32(4206968))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v39 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v46 = F_table_open(m, v44, int32(2))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if v48 <= int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L109
	}
L11:
	;
	F_sequence_close(m, v46, int32(2))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L108
	}
L12:
	;
	v52 = v23 + int32(236)
	v54 = int32(1)
	v56 = int32(0)
	v59 = v54
	v61 = v54
	v64 = v48
	goto L13
L13:
	;
	v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52+v56<<(uint(int32(1))%32)))))
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+6)))
	if v75 < v74 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v86&int32(1) != 0 {
		goto L11
	} else {
		goto L20
	}
L15:
	;
	F_slot_getsomeattrs_int(m, v35, v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v80 = v64
	goto L17
L17:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v83 = int32(1)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v74-v83))))
	v86 = v85 & v59
	v89 = (v85 ^ v83) & v61
	v91 = v56 + v83
	if v91 < v80 {
		v56 = v91
		v59 = v86
		v61 = v89
		v64 = v80
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v80 = v79
	goto L17
L19:
	;
	goto L14
L20:
	;
	if v89&int32(1) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L32
	}
L22:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+164)))
	switch v97 - int32(102) {
	case 0:
		goto L24
	default:
		goto L21
	case 13:
		goto L23
	}
L23:
	;
	F_sequence_close(m, v46, int32(2))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L31
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(50352322))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v43)+48))
	v109 = v23 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v107 + int32(4)
	F_errmsg(m, int32(726643), v18+int32(96))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errdetail(m, int32(618639), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errtableconstraint(m, v43, v109)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(513194), int32(319), int32(330977))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	goto L3
L32:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+684)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+680)) = v136
	v141 = v18 + int32(680)
	v144 = F_ri_FetchPreparedPlan(m, v141)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v144 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v539 = v144
	goto L36
L35:
	;
	F_initStringInfo(m, v18+int32(664))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	v540 = int32(0)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+119)))
	v547 = F_ri_PerformCheck(m, v23, v141, v539, v43, v46, v540, v35, v540, base.B2i32(v543 == int32(112)), int32(5))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L104
	}
L37:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+119)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v150)+68))
	v153 = F_get_namespace_name(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v155 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+400)) = uint8(v155)
	v159 = v153
	v161 = v18 + int32(400)
	goto L39
L39:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v174 != int32(34) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v191 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+1)) = uint16(v191)
	v194 = v18 + int32(400)
	v195 = F_strlen(m, v194)
	mBase = m.M
	v198 = v195 + v194
	v199 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v199)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v203 = v198 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v191)
	v208 = v203
	v210 = v201 + int32(4)
	goto L47
L41:
	;
	goto L40
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v186)
	v159 = v159 + int32(1)
	v161 = v187
	goto L39
L43:
	;
	if v174 == int32(0) {
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v181 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)) = uint8(v181)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	v186 = v183
	v187 = v161 + int32(2)
	goto L42
L46:
	;
	v186 = v174
	v187 = v161 + int32(1)
	goto L42
L47:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	if v223 != int32(34) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v240 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v208)+1)) = uint16(v240)
	if v151&int32(255) == int32(112) {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	goto L48
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v236))) = uint8(v235)
	v208 = v236
	v210 = v210 + int32(1)
	goto L47
L51:
	;
	if v223 == int32(0) {
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v230 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)) = uint8(v230)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	v235 = v232
	v236 = v208 + int32(2)
	goto L50
L54:
	;
	v235 = v223
	v236 = v208 + int32(1)
	goto L50
L55:
	;
	v248 = int32(785690)
	goto L57
L56:
	;
	v248 = int32(772114)
	goto L57
L57:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+165)))
	if v249 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if int32(0) < v337 {
		goto L73
	} else {
		goto L74
	}
L59:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v256 = int32(*(*int16)(unsafe.Add(mBase, uint32(v252<<(uint(int32(1))%32)+v23)+170)))
	v257 = F_attnumAttName(m, v46, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v18 + int32(400)
	F_appendStringInfo(m, v18+int32(664), int32(31006), v18+int32(80))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L72
	}
L62:
	;
	v259 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+256)) = uint8(v259)
	v263 = v257
	v265 = v18 + int32(256)
	goto L63
L63:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	if v278 != int32(34) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v295 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v265)+1)) = uint16(v295)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v18 + int32(400)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v18 + int32(256)
	F_appendStringInfo(m, v18+int32(664), int32(30964), v18-int32(-64))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L71
	}
L65:
	;
	goto L64
L66:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v291))) = uint8(v290)
	v263 = v263 + int32(1)
	v265 = v291
	goto L63
L67:
	;
	if v278 == int32(0) {
		goto L65
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v285 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v265)+1)) = uint8(v285)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	v290 = v287
	v291 = v265 + int32(2)
	goto L66
L70:
	;
	v290 = v278
	v291 = v265 + int32(1)
	goto L66
L71:
	;
	goto L58
L72:
	;
	goto L58
L73:
	;
	v351 = int32(0)
	v357 = int32(560790)
	goto L76
L74:
	;
	goto L75
L75:
	;
	F_appendStringInfoString(m, v18+int32(664), int32(31027))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L93
	}
L76:
	;
	v362 = v351 << (uint(int32(1)) % 32)
	v363 = v23 + int32(172) + v362
	v364 = int32(*(*int16)(unsafe.Add(mBase, uint32(v363))))
	v365 = F_attnumTypeId(m, v46, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	goto L75
L78:
	;
	v368 = int32(*(*int16)(unsafe.Add(mBase, uint32(v362+v52))))
	v369 = F_attnumTypeId(m, v43, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v371 = int32(*(*int16)(unsafe.Add(mBase, uint32(v363))))
	v372 = F_attnumAttName(m, v46, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v374 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+256)) = uint8(v374)
	v378 = v372
	v380 = v18 + int32(256)
	goto L81
L81:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	if v393 != int32(34) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v410 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v380)+1)) = uint16(v410)
	v413 = v351 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v413
	v420 = F_pg_sprintf(m, v18+int32(240), int32(485008), v18+int32(48))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L89
	}
L83:
	;
	goto L82
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v406))) = uint8(v405)
	v378 = v378 + int32(1)
	v380 = v406
	goto L81
L85:
	;
	if v393 == int32(0) {
		goto L83
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v400 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v380)+1)) = uint8(v400)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	v405 = v402
	v406 = v380 + int32(2)
	goto L84
L88:
	;
	v405 = v393
	v406 = v380 + int32(1)
	goto L84
L89:
	;
	v423 = v351 << (uint(int32(2)) % 32)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(300)+v423)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v357
	F_appendStringInfo(m, v18+int32(664), int32(766605), v18+int32(32))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_generate_operator_clause(m, v18+int32(664), v18+int32(256), v365, v425, v18+int32(240), v369)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+int32(112)+v423))) = v369
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if v413 < v447 {
		v351 = v413
		v357 = int32(564022)
		goto L76
	} else {
		goto L92
	}
L92:
	;
	goto L77
L93:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+165)))
	if v469 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v476 = int32(*(*int16)(unsafe.Add(mBase, uint32(v470<<(uint(int32(1))%32)+v52-int32(2)))))
	v477 = F_attnumTypeId(m, v43, v476)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v18)+664))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v522 = F_ri_PlanCheck(m, v516, v517, v18+int32(112), v18+int32(680), v43, v46)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L103
	}
L97:
	;
	F_appendStringInfoString(m, v18+int32(664), int32(773029))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v484
	v491 = F_pg_sprintf(m, v18+int32(240), int32(485008), v18+int32(16))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v23)+688))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(785690)
	F_appendStringInfo(m, v18+int32(664), int32(766605), v18)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_generate_operator_clause(m, v18+int32(664), v18+int32(240), v477, v493, int32(351775), int32(4537))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_appendStringInfoString(m, v18+int32(664), int32(701288))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	goto L96
L103:
	;
	v539 = v522
	goto L36
L104:
	;
	v549 = F_SPI_finish(m)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v549 != int32(2) {
		goto L10
	} else {
		goto L106
	}
L106:
	;
	F_sequence_close(m, v46, int32(2))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	goto L3
L108:
	;
	goto L3
L109:
	;
	F_errmsg_internal(m, int32(471975), int32(0))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(513194), int32(460), int32(330977))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RI_FKey_restrict_del(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	F_ri_CheckTrigger(m, l0, int32(320388), int32(3))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_restrict(m, v8, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_RI_FKey_restrict_upd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	F_ri_CheckTrigger(m, l0, int32(438899), int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_restrict(m, v8, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_ri_CheckTrigger(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return
		} else {
			F_errcode(m, int32(16908867))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg(m, int32(234756), v7)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					F_errfinish(m, int32(513194), int32(2175), int32(234030))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		if v12 != int32(442) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				F_errcode(m, int32(16908867))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
					F_errmsg(m, int32(234756), v7)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						F_errfinish(m, int32(513194), int32(2175), int32(234030))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if v15&int32(28) != int32(4) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return
				} else {
					F_errcode(m, int32(16908867))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = l1
						F_errmsg(m, int32(537330), v7-int32(-64))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							F_errfinish(m, int32(513194), int32(2184), int32(234030))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v21 = v15 & int32(3)
				switch l2 - int32(2) {
				case 0:
					if v21 == int32(2) {
						m.G0 = v7 + int32(80)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							F_errcode(m, int32(16908867))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l1
								F_errmsg(m, int32(559860), v7+int32(32))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_errfinish(m, int32(513194), int32(2198), int32(234030))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				case 1:
					if v21 != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return
						} else {
							F_errcode(m, int32(16908867))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = l1
								F_errmsg(m, int32(559329), v7+int32(48))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									F_errfinish(m, int32(513194), int32(2204), int32(234030))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						m.G0 = v7 + int32(80)
						return
					}
				default:
					if v21 == int32(0) {
						m.G0 = v7 + int32(80)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							F_errcode(m, int32(16908867))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
								F_errmsg(m, int32(538668), v7+int32(16))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									F_errfinish(m, int32(513194), int32(2192), int32(234030))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
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
			}
		}
	}
}
func F_ri_FetchPreparedPlan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1145]))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return v103
L2:
	;
	v50 = v13
	goto L4
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(3023656976388)
	v20 = F_hash_create(m, int32(415185), int32(64), v10, int32(40))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v51 = int32(0)
	v53 = F_hash_search(m, v50, l0, v51, v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L10
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1146])) = v20
	F_CacheRegisterSyscacheCallback(m, int32(19), int32(1501), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(51539607560)
	v36 = F_hash_create(m, int32(415086), int32(256), v10, int32(40))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1145])) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(292057776136)
	v45 = F_hash_create(m, int32(415499), int32(256), v10, int32(40))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1147])) = v45
	v49 = *(*int32)(unsafe.Add(mBase, _consts[1145]))
	v50 = v49
	goto L4
L10:
	;
	if v53 == int32(0) {
		v103 = v2
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v57 == int32(0) {
		v103 = v2
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v60 = int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v61 == int32(0) {
		v91 = v60
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v91 != 0 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 <= int32(0) {
		v91 = v60
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v71 = v2
	goto L16
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v71<<(uint(int32(2))%32))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+95)))
	if v79 == int32(0) {
		v91 = v79
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v91 = v79
	goto L13
L18:
	;
	v83 = v71 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v83 < v84 {
		v71 = v83
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v103 = v57
	goto L1
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = int32(0)
	F_SPI_freeplan(m, v57)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v103 = v2
	goto L1
}
