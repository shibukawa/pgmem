package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ScanKeyEntryInitialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	v3 = l2
	v4 = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l4
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v16 = l0 + int32(16)
	if l6 != 0 {
		F_fmgr_info(m, l6, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			return
		}
	} else {
		if v16&int32(3) == int32(0) {
			v25 = l0 + int32(44)
			v27 = l0 + int32(20)
			if base.Ui32(v27) < base.Ui32(v25) {
				v29 = v25
			} else {
				v29 = v27
			}
			v38 = F__emscripten_memset_bulkmem(m, v16, base.I32_extend8_s(int32(0)), (v29-l0-int32(17))&int32(-4)+int32(4))
			mBase = m.M
			return
		} else {
			v39 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v16))) = v39
			*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v39
			*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v39
			return
		}
	}
}
func F_SerialPagePrecedesLogically(m *base.Module, l0 int64, l1 int64) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v3 = int32(0)
	v7 = int32(10)
	v9 = int32(4)
	v10 = base.I32_wrap_i64(l0)<<(uint(v7)%32) | v9
	v13 = base.I32_wrap_i64(l1) << (uint(v7) % 32)
	v15 = v13 | v9
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v15))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == v3 {
		v27 = base.B2i32(base.Ui32(v10) < base.Ui32(v15))
	} else {
		v27 = int32(base.Ui32(v10-v15) >> (uint(int32(31)) % 32))
	}
	if v27 != 0 {
		v29 = v13 + int32(1027)
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v29))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v10)) == int32(0) {
			v41 = base.B2i32(base.Ui32(v10) < base.Ui32(v29))
		} else {
			v41 = int32(base.Ui32(v10-v29) >> (uint(int32(31)) % 32))
		}
		v42 = v41
	} else {
		v42 = v3
	}
	return v42
}
func F_ShmemAllocNoError(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v6 = *(*int32)(unsafe.Add(mBase, _consts[973]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
	if v7 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[973]))
		F_s_lock(m, v15, int32(520728), int32(208), int32(33935))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _consts[960]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
			v26 = v25 + (l0+int32(127))&int32(-128)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
			if base.Ui32(v26) <= base.Ui32(v27) {
				v30 = *(*int32)(unsafe.Add(mBase, _consts[959]))
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v26
				v33 = v30 + v25
			} else {
				v33 = int32(0)
			}
			v35 = *(*int32)(unsafe.Add(mBase, _consts[973]))
			*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(0)
			return v33
		}
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _consts[960]))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
		v26 = v25 + (l0+int32(127))&int32(-128)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
		if base.Ui32(v26) <= base.Ui32(v27) {
			v30 = *(*int32)(unsafe.Add(mBase, _consts[959]))
			*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v26
			v33 = v30 + v25
		} else {
			v33 = int32(0)
		}
		v35 = *(*int32)(unsafe.Add(mBase, _consts[973]))
		*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(0)
		return v33
	}
}
func F_SignalHandlerForConfigReload(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, _consts[715])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[506]))
	F_SetLatch(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_SplitIdentifierString(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
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
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
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
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v389 int32
	_ = v389
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v10 = l0
	goto L1
L1:
	;
	v19 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
	goto L3
L2:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v29 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	if base.B2i32(v19 == int32(32))|base.B2i32(base.Ui32((v19-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v10 = v10 + int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	return int32(1)
L6:
	;
	v32 = v29
	v35 = v10
	goto L7
L7:
	;
	v40 = v32 & int32(255)
	if v40 != int32(34) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v447 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v377))) = uint8(v447)
	v449 = F_strlen(m, v379)
	mBase = m.M
	F_truncate_identifier(m, v379, v449, v447)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L26
	} else {
		goto L126
	}
L10:
	;
	return v437
L11:
	;
	v437 = int32(0)
	goto L10
L12:
	;
	v380 = v373
	goto L114
L13:
	;
	if v40 == int32(0) {
		v75 = v35
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v195 = int32(0)
	v197 = v35 + int32(1)
	v198 = int32(34)
	v199 = F___strchrnul(m, v197, v198)
	mBase = m.M
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	if v201 == v198 {
		goto L56
	} else {
		goto L57
	}
L16:
	;
	if v35 == v75 {
		goto L11
	} else {
		goto L25
	}
L17:
	;
	v46 = l1 & int32(255)
	if v40 == v46 {
		v75 = v35
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v48 = v32
	v52 = v35
	goto L19
L19:
	;
	v55 = base.I32_extend8_s(v48)
	goto L21
L20:
	;
	v75 = v66
	goto L16
L21:
	;
	if base.B2i32(v55 == int32(32))|base.B2i32(base.Ui32((v55-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v75 = v52
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v66 = v52 + int32(1)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v67 == int32(0) {
		v75 = v66
		goto L16
	} else {
		goto L23
	}
L23:
	;
	if v67 != v46 {
		v48 = v67
		v52 = v66
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v79 = v75 - v35
	v81 = F_downcase_truncate_identifier(m, v35, v79, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	return int32(0)
L27:
	;
	if (v81^v35)&int32(3) != 0 {
		v154 = v81
		v155 = v79
		v156 = v35
		goto L32
	} else {
		goto L33
	}
L28:
	;
	F_pfree(m, v81)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L26
	} else {
		goto L54
	}
L29:
	;
	v192 = F___memset(m, v189, int32(0), v188)
	mBase = m.M
	goto L28
L30:
	;
	v188 = int32(0)
	v189 = v183
	goto L29
L31:
	;
	v166 = v161
	v167 = v162
	v168 = v163
	goto L50
L32:
	;
	if v155 == int32(0) {
		v183 = v156
		goto L30
	} else {
		goto L49
	}
L33:
	;
	v90 = int32(0)
	v91 = base.B2i32(v79 != v90)
	if v81&int32(3) == v90 {
		v120 = v81
		v121 = v79
		v122 = v35
		v123 = v91
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v123 == int32(0) {
		v183 = v122
		goto L30
	} else {
		goto L42
	}
L35:
	;
	if v79 == int32(0) {
		v120 = v81
		v121 = v79
		v122 = v35
		v123 = v91
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v99 = v81
	v100 = v79
	v101 = v35
	goto L37
L37:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v103)
	if v103 == int32(0) {
		v188 = v100
		v189 = v101
		goto L29
	} else {
		goto L39
	}
L38:
	;
	v120 = v114
	v121 = v110
	v122 = v108
	v123 = v112
	goto L34
L39:
	;
	v107 = int32(1)
	v108 = v101 + v107
	v110 = v100 - v107
	v111 = int32(0)
	v112 = base.B2i32(v110 != v111)
	v114 = v99 + v107
	if v114&int32(3) == v111 {
		v120 = v114
		v121 = v110
		v122 = v108
		v123 = v112
		goto L34
	} else {
		goto L40
	}
L40:
	;
	if v110 != 0 {
		v99 = v114
		v100 = v110
		v101 = v108
		goto L37
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v126 == int32(0) {
		v188 = v121
		v189 = v122
		goto L29
	} else {
		goto L43
	}
L43:
	;
	if base.Ui32(v121) < base.Ui32(int32(4)) {
		v154 = v120
		v155 = v121
		v156 = v122
		goto L32
	} else {
		goto L44
	}
L44:
	;
	v132 = v120
	v133 = v121
	v134 = v122
	goto L45
L45:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v140 = int32(-2139062144)
	if (int32(16843008)-v137|v137)&v140 != v140 {
		v161 = v132
		v162 = v133
		v163 = v134
		goto L31
	} else {
		goto L47
	}
L46:
	;
	v154 = v148
	v155 = v150
	v156 = v146
	goto L32
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v137
	v145 = int32(4)
	v146 = v134 + v145
	v148 = v132 + v145
	v150 = v133 - v145
	if base.Ui32(int32(3)) < base.Ui32(v150) {
		v132 = v148
		v133 = v150
		v134 = v146
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v161 = v154
	v162 = v155
	v163 = v156
	goto L31
L50:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v170)
	if v170 == int32(0) {
		v188 = v167
		v189 = v168
		goto L29
	} else {
		goto L52
	}
L51:
	;
	v183 = v175
	goto L30
L52:
	;
	v174 = int32(1)
	v175 = v168 + v174
	v179 = v167 - v174
	if v179 != 0 {
		v166 = v166 + v174
		v167 = v179
		v168 = v175
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v373 = v75
	v377 = v75
	v379 = v35
	goto L12
L55:
	;
	if v205 == int32(0) {
		v437 = v195
		goto L10
	} else {
		goto L59
	}
L56:
	;
	v205 = v199
	goto L58
L57:
	;
	v205 = v195
	goto L58
L58:
	;
	goto L55
L59:
	;
	v212 = v205
	goto L60
L60:
	;
	v216 = v212 + int32(1)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	if v217 != int32(34) {
		v373 = v216
		v377 = v212
		v379 = v197
		goto L12
	} else {
		goto L62
	}
L61:
	;
	v437 = v195
	goto L10
L62:
	;
	v220 = F_strlen(m, v212)
	mBase = m.M
	if v212 == v216 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v365 = int32(34)
	v366 = F___strchrnul(m, v216, v365)
	mBase = m.M
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	if v368 == v365 {
		goto L110
	} else {
		goto L111
	}
L64:
	;
	goto L63
L65:
	;
	v224 = v212 + v220
	if base.Ui32(v216-v224) <= base.Ui32(int32(0)-v220<<(uint(int32(1))%32)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v231 = F___memcpy(m, v212, v216, v220)
	mBase = m.M
	goto L63
L67:
	;
	goto L68
L68:
	;
	v234 = (v212 ^ v216) & int32(3)
	if base.Ui32(v212) < base.Ui32(v216) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	if v336 == int32(0) {
		goto L64
	} else {
		goto L105
	}
L70:
	;
	if base.Ui32(v314) <= base.Ui32(int32(3)) {
		v335 = v313
		v336 = v314
		v337 = v315
		goto L69
	} else {
		goto L101
	}
L71:
	;
	if v234 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if v234 != 0 {
		v296 = v220
		goto L84
	} else {
		goto L85
	}
L74:
	;
	v335 = v216
	v336 = v220
	v337 = v212
	goto L69
L75:
	;
	goto L76
L76:
	;
	if v212&int32(3) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v313 = v216
	v314 = v220
	v315 = v212
	goto L70
L78:
	;
	goto L79
L79:
	;
	v241 = v216
	v242 = v220
	v243 = v212
	goto L80
L80:
	;
	if v242 == int32(0) {
		goto L64
	} else {
		goto L82
	}
L81:
	;
	v313 = v250
	v314 = v252
	v315 = v254
	goto L70
L82:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	*(*uint8)(unsafe.Add(mBase, uint32(v243))) = uint8(v247)
	v249 = int32(1)
	v250 = v241 + v249
	v252 = v242 - v249
	v254 = v243 + v249
	if v254&int32(3) != 0 {
		v241 = v250
		v242 = v252
		v243 = v254
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	if v296 == int32(0) {
		goto L64
	} else {
		goto L97
	}
L85:
	;
	if v224&int32(3) != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v261 = v220
	goto L89
L87:
	;
	v276 = v220
	goto L88
L88:
	;
	if base.Ui32(v276) <= base.Ui32(int32(3)) {
		v296 = v276
		goto L84
	} else {
		goto L93
	}
L89:
	;
	if v261 == int32(0) {
		goto L64
	} else {
		goto L91
	}
L90:
	;
	v276 = v267
	goto L88
L91:
	;
	v267 = v261 - int32(1)
	v268 = v212 + v267
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216+v267))))
	*(*uint8)(unsafe.Add(mBase, uint32(v268))) = uint8(v270)
	if v268&int32(3) != 0 {
		v261 = v267
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v283 = v276
	goto L94
L94:
	;
	v287 = v283 - int32(4)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v216+v287)))
	*(*int32)(unsafe.Add(mBase, uint32(v212+v287))) = v290
	if base.Ui32(int32(3)) < base.Ui32(v287) {
		v283 = v287
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v296 = v287
	goto L84
L96:
	;
	goto L95
L97:
	;
	v303 = v296
	goto L98
L98:
	;
	v307 = v303 - int32(1)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216+v307))))
	*(*uint8)(unsafe.Add(mBase, uint32(v212+v307))) = uint8(v310)
	if v307 != 0 {
		v303 = v307
		goto L98
	} else {
		goto L100
	}
L99:
	;
	goto L64
L100:
	;
	goto L99
L101:
	;
	v320 = v313
	v321 = v314
	v322 = v315
	goto L102
L102:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	*(*int32)(unsafe.Add(mBase, uint32(v322))) = v324
	v326 = int32(4)
	v327 = v320 + v326
	v329 = v322 + v326
	v331 = v321 - v326
	if base.Ui32(int32(3)) < base.Ui32(v331) {
		v320 = v327
		v321 = v331
		v322 = v329
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v335 = v327
	v336 = v331
	v337 = v329
	goto L69
L104:
	;
	goto L103
L105:
	;
	v342 = v335
	v343 = v336
	v344 = v337
	goto L106
L106:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	*(*uint8)(unsafe.Add(mBase, uint32(v344))) = uint8(v346)
	v348 = int32(1)
	v353 = v343 - v348
	if v353 != 0 {
		v342 = v342 + v348
		v343 = v353
		v344 = v344 + v348
		goto L106
	} else {
		goto L108
	}
L107:
	;
	goto L64
L108:
	;
	goto L107
L109:
	;
	if v372 != 0 {
		v212 = v372
		goto L60
	} else {
		goto L113
	}
L110:
	;
	v372 = v366
	goto L112
L111:
	;
	v372 = int32(0)
	goto L112
L112:
	;
	goto L109
L113:
	;
	goto L61
L114:
	;
	v389 = int32(*(*int8)(unsafe.Add(mBase, uint32(v380))))
	goto L116
L115:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	v401 = l1 & int32(255)
	if v399 == v401 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	if base.B2i32(v389 == int32(32))|base.B2i32(base.Ui32((v389-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v380 = v380 + int32(1)
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	v406 = v380
	goto L121
L119:
	;
	goto L120
L120:
	;
	if v399 == int32(0) {
		v443 = v380
		goto L9
	} else {
		goto L125
	}
L121:
	;
	v411 = v406 + int32(1)
	v412 = int32(*(*int8)(unsafe.Add(mBase, uint32(v411))))
	goto L123
L123:
	;
	if base.B2i32(v412 == int32(32))|base.B2i32(base.Ui32((v412-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v406 = v411
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v443 = v411
	goto L9
L125:
	;
	goto L11
L126:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v454 = F_lappend(m, v453, v379)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L26
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v454
	if v399 != v401 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	v32 = v458
	v35 = v443
	goto L7
}
func F_StatementTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, _consts[712]))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[226])))
	if v9 != 0 {
		v10 = int32(15)
	} else {
		v10 = int32(2)
	}
	v11 = F_kill(m, int32(0)-v4, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[712]))
		v15 = F_kill(m, v14, v10)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			return
		}
	}
}
func F___sin(m *base.Module, l0 float64, l1 float64, l2 int32) float64 {
	var v7 float64
	_ = v7
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	v7 = base.F64_mul(l0, l0)
	v22 = base.F64_add(base.F64_mul(base.F64_mul(v7, base.F64_mul(v7, v7)), base.F64_add(base.F64_mul(v7, float64(1.58969099521155e-10)), float64(-2.5050760253406863e-08))), base.F64_add(base.F64_mul(v7, base.F64_add(base.F64_mul(v7, float64(2.7557313707070068e-06)), float64(-0.0001984126982985795))), float64(0.00833333333332249)))
	v23 = base.F64_mul(l0, v7)
	if l2 == int32(0) {
		return base.F64_add(base.F64_mul(v23, base.F64_add(base.F64_mul(v7, v22), float64(-0.16666666666666632))), l0)
	} else {
		return base.F64_sub(l0, base.F64_add(base.F64_sub(base.F64_mul(v7, base.F64_sub(base.F64_mul(l1, float64(0.5)), base.F64_mul(v23, v22))), l1), base.F64_mul(v23, float64(0.16666666666666632))))
	}
}
func F___stdio_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	v4 = int32(0)
	v8 = m.G0
	v9 = int32(32)
	v10 = v8 - v9
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l2 - base.B2i32(v13 != v4)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v18
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v28 = m.Wasi_snapshot_preview1.Fd_read(m, v22, v10+int32(16), int32(2), v10+int32(12))
	mBase = m.M
	if v28 == v4 {
		v35 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[140])) = v28
		v35 = int32(-1)
	}
	if v35 != 0 {
		v43 = v9
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v43 | v44
		v64 = v4
	} else {
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		if int32(0) < v36 {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			if base.Ui32(v36) <= base.Ui32(v47) {
				v64 = v36
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v49 + (v36 - v47)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if v54 != 0 {
					v55 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v49 + v55
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
					*(*uint8)(unsafe.Add(mBase, uint32(l1+l2-v55))) = uint8(v61)
				} else {
				}
				v64 = l2
			}
		} else {
			if v36 != 0 {
				v41 = int32(32)
			} else {
				v41 = int32(16)
			}
			v43 = v41
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v43 | v44
			v64 = v4
		}
	}
	m.G0 = v10 + int32(32)
	return v64
}
func F___stdio_seek(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v5 = F___lseek(m, v4, l1, l2)
	mBase = m.M
	return v5
}
func F___strerror_l(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	if base.Ui32(l0) <= base.Ui32(int32(153)) {
		v5 = l0
	} else {
		v5 = int32(0)
	}
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5<<(uint(int32(1))%32))+uint32(_consts[1637]))))
	return v10 + int32(4139812)
}
func F_s_lock_stuck(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		if l2 != 0 {
			v16 = l2
		} else {
			v16 = int32(705831)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v16
		F_errmsg_internal(m, int32(487365), v7)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_errfinish(m, int32(521218), int32(90), int32(330781))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_save_ps_display_args(m *base.Module, l0 int32, l1 int32) int32 {
	return l1
}
func F_sbrk(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1649]))
	v9 = (l0 + int32(7)) & int32(-8)
	v10 = v5 + v9
	if base.Ui32(v10) <= base.Ui32(v5) {
		v13 = v9
	} else {
		v13 = int32(0)
	}
	if v13 == int32(0) {
		if base.Ui32(v10) <= base.Ui32(base.MemorySize(m)<<(uint(int32(16))%32)) {
			*(*int32)(unsafe.Add(mBase, _consts[1649])) = v10
			return v5
		} else {
			v20 = m.Env.Emscripten_resize_heap(m, v10)
			mBase = m.M
			if v20 != 0 {
				*(*int32)(unsafe.Add(mBase, _consts[1649])) = v10
				return v5
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(48)
				return int32(-1)
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(48)
		return int32(-1)
	}
}
func F_scalarltjoinsel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_Float8GetDatum(m, float64(0.3333333333333333))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_scalarltsel(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_scalarineqsel_wrapper(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_searchstoplist(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 == v3 {
		v28 = v3
		m.G0 = v7 + int32(16)
		return v28
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v14 <= int32(0) {
			v28 = v3
			m.G0 = v7 + int32(16)
			return v28
		} else {
			v21 = F_bsearch(m, v7+int32(12), v11, v14, int32(4), int32(1186))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v28 = base.B2i32(v21 != int32(0))
				m.G0 = v7 + int32(16)
				return v28
			}
		}
	}
}
func F_send(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v4 = int32(0)
	v6 = F_sendto(m, l0, l1, l2, v4, v4)
	return v6
}
func F_send_feedback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v81 int64
	_ = v81
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v142 int64
	_ = v142
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
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
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v212 int64
	_ = v212
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v256 int64
	_ = v256
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v263 int64
	_ = v263
	var v265 int64
	_ = v265
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int64
	_ = v299
	var v301 int64
	_ = v301
	var v303 int64
	_ = v303
	var v306 int64
	_ = v306
	var v308 int64
	_ = v308
	var v310 int64
	_ = v310
	var v312 int64
	_ = v312
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int64
	_ = v346
	var v348 int64
	_ = v348
	var v350 int64
	_ = v350
	var v353 int64
	_ = v353
	var v355 int64
	_ = v355
	var v357 int64
	_ = v357
	var v359 int64
	_ = v359
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int64
	_ = v402
	var v403 int64
	_ = v403
	var v408 int64
	_ = v408
	var v412 int64
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int64
	_ = v434
	var v439 int64
	_ = v439
	var v444 int64
	_ = v444
	v3 = l2
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	if l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[857]))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v21 = *(*int64)(unsafe.Add(mBase, _consts[858]))
	if base.Ui64(v21) < base.Ui64(l0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v23 = l0
	goto L8
L7:
	;
	v23 = v21
	goto L8
L8:
	;
	v27 = int32(4449808)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+280)) = v29
	*(*int64)(unsafe.Add(mBase, _consts[276])) = v29
	v34 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v34)+272)) = v35
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v35
	goto L11
L9:
	;
	v45 = int64(0)
	v47 = *(*int32)(unsafe.Add(mBase, _consts[859]))
	if v47 == int32(0) {
		v92 = v47
		v95 = v45
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	v44 = *(*int64)(unsafe.Add(mBase, _consts[276]))
	goto L9
L13:
	;
	v113 = *(*int64)(unsafe.Add(mBase, _consts[860]))
	if base.Ui64(v113) < base.Ui64(v108) {
		goto L33
	} else {
		goto L34
	}
L14:
	;
	if v92 != int32(4159816) {
		goto L27
	} else {
		goto L28
	}
L15:
	;
	if v47 == int32(4159816) {
		v92 = v47
		v95 = v45
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v47)+8))
	if base.Ui64(v52) <= base.Ui64(v44) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[859]))
	v92 = v88
	v95 = v63
	goto L14
L18:
	;
	v57 = v47
	goto L21
L19:
	;
	v81 = v45
	goto L20
L20:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[861]))
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
	v108 = v81
	v111 = v86
	goto L13
L21:
	;
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v57)+16))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v67
	F_pfree(m, v57)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v81 = v63
	goto L20
L23:
	;
	return
L24:
	;
	if v65 == int32(4159816) {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	if base.Ui64(v73) <= base.Ui64(v44) {
		v57 = v65
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v100 = v95
	goto L29
L28:
	;
	v100 = v23
	goto L29
L29:
	;
	if v92 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v101 = v100
	goto L32
L31:
	;
	v101 = v23
	goto L32
L32:
	;
	v108 = v101
	v111 = v101
	goto L13
L33:
	;
	v115 = v108
	goto L35
L34:
	;
	v115 = v113
	goto L35
L35:
	;
	v117 = *(*int64)(unsafe.Add(mBase, _consts[862]))
	if base.Ui64(v117) < base.Ui64(v111) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v119 = v111
	goto L38
L37:
	;
	v119 = v117
	goto L38
L38:
	;
	v123 = m.G0
	v124 = int32(16)
	v125 = v123 - v124
	m.G0 = v125
	F___gettimeofday(m, v125)
	mBase = m.M
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v125)))
	v129 = int64(*(*int32)(unsafe.Add(mBase, uint32(v125)+8)))
	m.G0 = v125 + v124
	v137 = v129 + v128*int64(1000000) - int64(946684800000000)
	goto L39
L39:
	;
	if l1 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, _consts[863])) = v137
	v160 = *(*int32)(unsafe.Add(mBase, _consts[864]))
	if v160 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	v139 = *(*int64)(unsafe.Add(mBase, _consts[862]))
	if v119 != v139 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v142 = *(*int64)(unsafe.Add(mBase, _consts[860]))
	if v115 != v142 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v145 = *(*int64)(unsafe.Add(mBase, _consts[863]))
	v147 = *(*int32)(unsafe.Add(mBase, _consts[857]))
	goto L44
L44:
	;
	if base.B2i32(base.I64_extend_i32_s(v147*int32(1000))*int64(1000) <= v137-v145) == int32(0) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L40
L46:
	;
	F_enlargeStringInfo(m, v184, int32(1))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L23
	} else {
		goto L52
	}
L47:
	;
	v163 = int32(4554240)
	v164 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v167 = *(*int32)(unsafe.Add(mBase, _consts[855]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v167
	v169 = F_makeStringInfo(m)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L23
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v176 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v176)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+12)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = v176
	goto L51
L50:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v164
	*(*int32)(unsafe.Add(mBase, _consts[864])) = v169
	v184 = v169
	goto L46
L51:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _consts[864]))
	v184 = v183
	goto L46
L52:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v192 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v189+v190))) = uint8(v192)
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v189 + int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, _consts[864]))
	F_enlargeStringInfo(m, v198, int32(8))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v205 = int64(56)
	v207 = int64(65280)
	v209 = int64(40)
	v212 = int64(16711680)
	v214 = int64(24)
	v216 = int64(4278190080)
	v218 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v202+v203))) = v23<<(uint(v205)%64) | v23&v207<<(uint(v209)%64) | (v23&v212<<(uint(v214)%64) | v23&v216<<(uint(v218)%64)) | (int64(base.Ui64(v23)>>(uint(v218)%64))&v216 | int64(base.Ui64(v23)>>(uint(v214)%64))&v212 | (int64(base.Ui64(v23)>>(uint(v209)%64))&v207 | int64(base.Ui64(v23)>>(uint(v205)%64))))
	v241 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v198)+4)) = v202 + v241
	v245 = *(*int32)(unsafe.Add(mBase, _consts[864]))
	F_enlargeStringInfo(m, v245, v241)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L23
	} else {
		goto L54
	}
L54:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v252 = int64(56)
	v254 = int64(65280)
	v256 = int64(40)
	v259 = int64(16711680)
	v261 = int64(24)
	v263 = int64(4278190080)
	v265 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v249+v250))) = v115<<(uint(v252)%64) | v115&v254<<(uint(v256)%64) | (v115&v259<<(uint(v261)%64) | v115&v263<<(uint(v265)%64)) | (int64(base.Ui64(v115)>>(uint(v265)%64))&v263 | int64(base.Ui64(v115)>>(uint(v261)%64))&v259 | (int64(base.Ui64(v115)>>(uint(v256)%64))&v254 | int64(base.Ui64(v115)>>(uint(v252)%64))))
	v288 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v245)+4)) = v249 + v288
	v292 = *(*int32)(unsafe.Add(mBase, _consts[864]))
	F_enlargeStringInfo(m, v292, v288)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L23
	} else {
		goto L55
	}
L55:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v299 = int64(56)
	v301 = int64(65280)
	v303 = int64(40)
	v306 = int64(16711680)
	v308 = int64(24)
	v310 = int64(4278190080)
	v312 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v296+v297))) = v119<<(uint(v299)%64) | v119&v301<<(uint(v303)%64) | (v119&v306<<(uint(v308)%64) | v119&v310<<(uint(v312)%64)) | (int64(base.Ui64(v119)>>(uint(v312)%64))&v310 | int64(base.Ui64(v119)>>(uint(v308)%64))&v306 | (int64(base.Ui64(v119)>>(uint(v303)%64))&v301 | int64(base.Ui64(v119)>>(uint(v299)%64))))
	v335 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v292)+4)) = v296 + v335
	v339 = *(*int32)(unsafe.Add(mBase, _consts[864]))
	F_enlargeStringInfo(m, v339, v335)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L23
	} else {
		goto L56
	}
L56:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v346 = int64(56)
	v348 = int64(65280)
	v350 = int64(40)
	v353 = int64(16711680)
	v355 = int64(24)
	v357 = int64(4278190080)
	v359 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v343+v344))) = v137<<(uint(v346)%64) | v137&v348<<(uint(v350)%64) | (v137&v353<<(uint(v355)%64) | v137&v357<<(uint(v359)%64)) | (int64(base.Ui64(v137)>>(uint(v359)%64))&v357 | int64(base.Ui64(v137)>>(uint(v355)%64))&v353 | (int64(base.Ui64(v137)>>(uint(v350)%64))&v348 | int64(base.Ui64(v137)>>(uint(v346)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v339)+4)) = v343 + int32(8)
	v386 = *(*int32)(unsafe.Add(mBase, _consts[864]))
	F_enlargeStringInfo(m, v386, int32(1))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L23
	} else {
		goto L57
	}
L57:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	*(*uint8)(unsafe.Add(mBase, uint32(v390+v391))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v386)+4)) = v390 + int32(1)
	v399 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L23
	} else {
		goto L58
	}
L58:
	;
	if v399 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+24)) = uint32(v115)
	v402 = int64(32)
	v403 = int64(base.Ui64(v115) >> (uint(v402) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+20)) = uint32(v403)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+16)) = uint32(v119)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	v408 = int64(base.Ui64(v119) >> (uint(v402) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+12)) = uint32(v408)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+8)) = uint32(v23)
	v412 = int64(base.Ui64(v23) >> (uint(v402) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+4)) = uint32(v412)
	F_errmsg_internal(m, int32(539951), v12)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L23
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _consts[865]))
	v425 = *(*int32)(unsafe.Add(mBase, _consts[864]))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	v429 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)+44))
	m.T0[v430].(func(*base.Module, int32, int32, int32))(m, v423, v426, v427)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L23
	} else {
		goto L64
	}
L62:
	;
	F_errfinish(m, int32(518511), int32(3910), int32(334520))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L23
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v434 = *(*int64)(unsafe.Add(mBase, _consts[858]))
	if base.Ui64(v434) < base.Ui64(v23) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	*(*int64)(unsafe.Add(mBase, _consts[858])) = v23
	goto L67
L66:
	;
	goto L67
L67:
	;
	v439 = *(*int64)(unsafe.Add(mBase, _consts[862]))
	if base.Ui64(v439) < base.Ui64(v119) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, _consts[862])) = v119
	goto L70
L69:
	;
	goto L70
L70:
	;
	v444 = *(*int64)(unsafe.Add(mBase, _consts[860]))
	if base.Ui64(v115) <= base.Ui64(v444) {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, _consts[860])) = v115
	goto L1
}
func F_setCompoundAffixFlagValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v13 == int32(2) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L15
	} else {
		goto L47
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L15
	} else {
		goto L43
	}
L3:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v139
	m.G0 = v11 + int32(32)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v23 = F_strtox_2(m, l2, v11+int32(28), int32(10), int64(2147483648))
	mBase = m.M
	v24 = base.I32_wrap_i64(v23)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v34 = F_strlen(m, l2)
	mBase = m.M
	v36 = v34 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v36) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if l2 == v25 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v28 == int32(68) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if base.Ui32(int32(65537)) <= base.Ui32(v24) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v24
	goto L3
L11:
	;
	if (l2^v58)&int32(3) != 0 {
		goto L25
	} else {
		goto L26
	}
L12:
	;
	v39 = F_palloc0(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v44 = (v34 + int32(8)) & int32(4088)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v44) <= base.Ui32(v45) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	return
L16:
	;
	v58 = v39
	goto L11
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v52 - v44
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v53 + v44
	v58 = v53
	goto L11
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v52 = v45
	v53 = v47
	goto L17
L19:
	;
	goto L20
L20:
	;
	v48 = int32(8192)
	v50 = F_palloc0(m, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v52 = v48
	v53 = v50
	goto L17
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v58
	goto L3
L23:
	;
	goto L22
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v114)
	if v114&int32(255) == int32(0) {
		goto L23
	} else {
		goto L39
	}
L25:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v113 = l2
	v114 = v66
	v115 = v58
	goto L24
L26:
	;
	goto L27
L27:
	;
	if l2&int32(3) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v70 = l2
	v72 = v58
	goto L31
L29:
	;
	v84 = l2
	v86 = v58
	goto L30
L30:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v91 = int32(-2139062144)
	if (int32(16843008)-v88|v88)&v91 != v91 {
		v113 = v84
		v114 = v88
		v115 = v86
		goto L24
	} else {
		goto L35
	}
L31:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v73)
	if v73 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L32:
	;
	v84 = v80
	v86 = v78
	goto L30
L33:
	;
	v77 = int32(1)
	v78 = v72 + v77
	v80 = v70 + v77
	if v80&int32(3) != 0 {
		v70 = v80
		v72 = v78
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v96 = v84
	v97 = v88
	v98 = v86
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v97
	v100 = int32(4)
	v101 = v98 + v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v104 = v96 + v100
	v108 = int32(-2139062144)
	if (v102|(int32(16843008)-v102))&v108 == v108 {
		v96 = v104
		v97 = v102
		v98 = v101
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v113 = v104
	v114 = v102
	v115 = v101
	goto L24
L38:
	;
	goto L37
L39:
	;
	v122 = v113
	v124 = v115
	goto L40
L40:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)) = uint8(v125)
	v127 = int32(1)
	if v125 != 0 {
		v122 = v122 + v127
		v124 = v124 + v127
		goto L40
	} else {
		goto L42
	}
L41:
	;
	goto L23
L42:
	;
	goto L41
L43:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
	F_errmsg(m, int32(745516), v11)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L15
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(520914), int32(1041), int32(364311))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L15
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l2
	F_errmsg(m, int32(420692), v11+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(520914), int32(1045), int32(364311))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L15
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_set_deparse_for_query(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v906 int32
	_ = v906
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v954 int32
	_ = v954
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	v4 = int32(0)
	v22 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(v4), int32(80))
	mBase = m.M
	goto L1
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v23
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v27
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v33
	F_set_rtable_names(m, v22, l2, v24)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	v48 = v4
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v48
	if v48 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v70 != 0 {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v59 = v58
	goto L8
L7:
	;
	v59 = int32(0)
	goto L8
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v60 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v63 = v61
	goto L11
L10:
	;
	v63 = int32(0)
	goto L11
L11:
	;
	if v59 < v63 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v66 = F_palloc0(m, int32(52))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L5
L15:
	;
	v68 = F_lappend(m, v48, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v48 = v68
	goto L4
L17:
	;
	v71 = F_has_dangerous_join_using(m, v22, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	v80 = v48
	v81 = v60
	goto L19
L19:
	;
	v94 = v4
	goto L22
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+32)) = uint8(v71)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	F_set_using_names(m, v22, v74, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v80 = v79
	v81 = v78
	goto L19
L22:
	;
	v100 = int32(0)
	if v81 == v100 {
		v110 = v100
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v80 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v104 <= v94 {
		v110 = int32(0)
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v110 = v106 + v94<<(uint(int32(2))%32)
	goto L24
L27:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	if int32(0) < v698 {
		goto L160
	} else {
		goto L161
	}
L28:
	;
	return
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v113 <= v94 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	if v110 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v120 = v117 + v94<<(uint(int32(2))%32)
	if v120 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	if v125 == int32(2) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v123)+32))
	v131 = int32(2)
	v134 = int32(4)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v129+v130<<(uint(v131)%32)-v134)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v123)+28))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v129+v137<<(uint(v131)%32)-v134)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	if v145 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	F_set_relation_column_names(m, v22, v124, v123)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L2
	} else {
		goto L159
	}
L36:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v148 = v146
	goto L38
L37:
	;
	v148 = int32(0)
	goto L38
L38:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v149 < v148 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v151 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L41
L41:
	;
	F_build_colinfo_names_hash(m, v123)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L48
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v164
	goto L41
L43:
	;
	v156 = F_palloc0(m, v148<<(uint(int32(2))%32))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v158 = int32(2)
	v162 = F_repalloc0(m, v151, v149<<(uint(v158)%32), v148<<(uint(v158)%32))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L47
	}
L46:
	;
	v164 = v156
	goto L42
L47:
	;
	v164 = v162
	goto L42
L48:
	;
	v171 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v123)+44))
	if v173 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v175 = v174
	goto L51
L50:
	;
	v175 = v171
	goto L51
L51:
	;
	if v175 < v148 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v177 = v175
	v186 = v171
	goto L55
L53:
	;
	v312 = v173
	v319 = v171
	goto L54
L54:
	;
	v328 = int32(0)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	if v312 != 0 {
		goto L94
	} else {
		goto L95
	}
L55:
	;
	v196 = v177 << (uint(int32(2)) % 32)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+36))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v196+v197)))
	if int32(0) < v199 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v123)+44))
	v312 = v309
	v319 = v305
	goto L54
L57:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v229 = v228 + v196
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	if v230 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L58:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v227 = v202 + v199<<(uint(int32(2))%32) - int32(4)
	goto L57
L59:
	;
	goto L60
L60:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v123)+40))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v208+v196)))
	if int32(0) < v210 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v227 = v213 + v210<<(uint(int32(2))%32) - int32(4)
	goto L57
L62:
	;
	goto L63
L63:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v221+v196)))
	v227 = v223 + int32(4)
	goto L57
L64:
	;
	v307 = v177 + int32(1)
	if v307 != v148 {
		v177 = v307
		v186 = v305
		goto L55
	} else {
		goto L93
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = int32(0)
	v305 = v186
	goto L64
L66:
	;
	goto L67
L67:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v235 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = v230
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v239 == int32(0) {
		v305 = v186
		goto L64
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	if v246 != 0 {
		v270 = v246
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v244 = F_hash_search(m, v239, v230, int32(1), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v305 = v186
	goto L64
L73:
	;
	v272 = int32(1)
	if v186&v272 != 0 {
		v305 = v272
		goto L64
	} else {
		goto L84
	}
L74:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v235)+8))
	if v247 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+4))
	v250 = v248
	goto L77
L76:
	;
	v250 = int32(0)
	goto L77
L77:
	;
	if v177 < v250 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v252+v196)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+4))
	v256 = v255
	goto L80
L79:
	;
	v256 = v230
	goto L80
L80:
	;
	v257 = F_make_colname_unique(m, v256, v22, v123)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v259+v196))) = v257
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v262 == int32(0) {
		v270 = v257
		goto L73
	} else {
		goto L82
	}
L82:
	;
	v267 = F_hash_search(m, v262, v257, int32(1), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	v270 = v257
	goto L73
L84:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v278 == int32(0) {
		v297 = v277
		v298 = v278
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v305 = base.B2i32(v298-v297 != int32(0))
	goto L64
L86:
	;
	goto L85
L87:
	;
	if v277 != v278 {
		v297 = v277
		v298 = v278
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v282 = v270
	v283 = v230
	goto L89
L89:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)))
	if v287 == int32(0) {
		v297 = v286
		v298 = v287
		goto L86
	} else {
		goto L91
	}
L90:
	;
	v297 = v286
	v298 = v287
	goto L86
L91:
	;
	v290 = int32(1)
	if v286 == v287 {
		v282 = v282 + v290
		v283 = v283 + v290
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	goto L56
L94:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v334 = v333
	goto L96
L95:
	;
	v334 = v328
	goto L96
L96:
	;
	v335 = v330 + v331 - v334
	*(*int32)(unsafe.Add(mBase, uint32(v123)+8)) = v335
	v339 = F_palloc0(m, v335<<(uint(int32(2))%32))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = v339
	v342 = F_palloc0(m, v335)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L2
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+16)) = v342
	if v148 <= int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	if v428 <= int32(0) {
		goto L116
	} else {
		goto L117
	}
L100:
	;
	v347 = int32(0)
	v414 = v347
	v425 = v328
	v426 = v347
	goto L99
L101:
	;
	goto L102
L102:
	;
	v349 = int32(0)
	v355 = v349
	v366 = v328
	v367 = v349
	goto L103
L103:
	;
	v370 = v355 << (uint(int32(2)) % 32)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v123)+36))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v370+v371)))
	if v373 == int32(0) {
		v414 = v355
		v425 = v366
		v426 = v367
		goto L99
	} else {
		goto L105
	}
L104:
	;
	v414 = v148
	v425 = v406
	v426 = v398
	goto L99
L105:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v123)+40))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v376+v370)))
	if v378 == int32(0) {
		v414 = v355
		v425 = v366
		v426 = v367
		goto L99
	} else {
		goto L106
	}
L106:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v383+v370)))
	*(*int32)(unsafe.Add(mBase, uint32(v381+v370))) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v389 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v387+v355))) = uint8(v389)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v123)+36))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v391+v370)))
	if v389 < v393 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v396 = F_bms_add_member(m, v367, v393)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L2
	} else {
		goto L110
	}
L108:
	;
	v398 = v367
	goto L109
L109:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v123)+40))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v399+v370)))
	if int32(0) < v401 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v398 = v396
	goto L109
L111:
	;
	v404 = F_bms_add_member(m, v366, v401)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L2
	} else {
		goto L114
	}
L112:
	;
	v406 = v366
	goto L113
L113:
	;
	v408 = v355 + int32(1)
	if v408 != v148 {
		v355 = v408
		v366 = v406
		v367 = v398
		goto L103
	} else {
		goto L115
	}
L114:
	;
	v406 = v404
	goto L113
L115:
	;
	goto L104
L116:
	;
	v680 = v414
	v684 = v414
	v689 = v319
	goto L27
L117:
	;
	goto L118
L118:
	;
	v431 = int32(0)
	v433 = v414
	v435 = v431
	v437 = v414
	v441 = v431
	v442 = v319
	goto L119
L119:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451+v441))))
	if v453 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v680 = v653
	v684 = v657
	v689 = v662
	goto L27
L121:
	;
	v672 = v441 + int32(1)
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v143)+8))
	if v672 < v673 {
		v433 = v653
		v435 = v655
		v437 = v657
		v441 = v672
		v442 = v662
		goto L119
	} else {
		goto L158
	}
L122:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v647+v441))))
	*(*uint8)(unsafe.Add(mBase, uint32(v645+v437))) = uint8(v649)
	v653 = v627
	v655 = v629
	v657 = v437 + int32(1)
	v662 = v636
	goto L121
L123:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	if v456 <= v435 {
		v486 = v435
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v143)+12))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v563+v441<<(uint(int32(2))%32))))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v568 != 0 {
		goto L141
	} else {
		goto L142
	}
L126:
	;
	v503 = v486 + int32(1)
	v504 = F_bms_is_member(m, v503, v426)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L2
	} else {
		goto L132
	}
L127:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v461 = v435
	goto L128
L128:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v458+v461<<(uint(int32(2))%32))))
	if v480 != 0 {
		v486 = v461
		goto L126
	} else {
		goto L130
	}
L129:
	;
	v486 = v456
	goto L126
L130:
	;
	v482 = v461 + int32(1)
	if v482 != v456 {
		v461 = v482
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	if v504 != 0 {
		v653 = v433
		v655 = v503
		v657 = v437
		v662 = v442
		goto L121
	} else {
		goto L133
	}
L133:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v507 <= v433 {
		v534 = v433
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v553 = int32(2)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v506+v534<<(uint(v553)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v552+v437<<(uint(v553)%32)))) = v559
	v627 = v534 + int32(1)
	v629 = v503
	v636 = v442
	goto L122
L135:
	;
	v509 = v433
	goto L136
L136:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v506+v509<<(uint(int32(2))%32))))
	if v530 != 0 {
		v534 = v509
		goto L134
	} else {
		goto L138
	}
L137:
	;
	v534 = v507
	goto L134
L138:
	;
	v532 = v509 + int32(1)
	if v532 != v507 {
		v509 = v532
		goto L136
	} else {
		goto L139
	}
L139:
	;
	goto L137
L140:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v617 != 0 {
		goto L154
	} else {
		goto L155
	}
L141:
	;
	v569 = F_make_colname_unique(m, v567, v22, v123)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L2
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v609+v437<<(uint(int32(2))%32)))) = v567
	v616 = v442
	goto L140
L144:
	;
	v572 = v437 << (uint(int32(2)) % 32)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v572+v573))) = v569
	v576 = int32(1)
	if v442&v576 != 0 {
		v616 = v576
		goto L140
	} else {
		goto L145
	}
L145:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v579+v572)))
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567))))
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581))))
	if v585 == int32(0) {
		v604 = v584
		v605 = v585
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v616 = base.B2i32(v605-v604 != int32(0))
	goto L140
L147:
	;
	goto L146
L148:
	;
	if v584 != v585 {
		v604 = v584
		v605 = v585
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v589 = v581
	v590 = v567
	goto L150
L150:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590)+1)))
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589)+1)))
	if v594 == int32(0) {
		v604 = v593
		v605 = v594
		goto L147
	} else {
		goto L152
	}
L151:
	;
	v604 = v593
	v605 = v594
	goto L147
L152:
	;
	v597 = int32(1)
	if v593 == v594 {
		v589 = v589 + v597
		v590 = v590 + v597
		goto L150
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v618+v437<<(uint(int32(2))%32))))
	v625 = F_hash_search(m, v617, v622, int32(1), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L2
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v627 = v433
	v629 = v435
	v636 = v616
	goto L122
L157:
	;
	goto L156
L158:
	;
	goto L120
L159:
	;
	v94 = v94 + int32(1)
	goto L22
L160:
	;
	v701 = int32(0)
	v703 = v680
	v705 = v701
	v707 = v684
	v709 = v701
	v712 = v689
	goto L163
L161:
	;
	v954 = v689
	goto L162
L162:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v963 != 0 {
		goto L203
	} else {
		goto L204
	}
L163:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721+v709))))
	if v723 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	v954 = v932
	goto L162
L165:
	;
	v942 = v709 + int32(1)
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	if v942 < v943 {
		v703 = v923
		v705 = v925
		v707 = v927
		v709 = v942
		v712 = v932
		goto L163
	} else {
		goto L202
	}
L166:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v917+v709))))
	*(*uint8)(unsafe.Add(mBase, uint32(v915+v707))) = uint8(v919)
	v923 = v897
	v925 = v899
	v927 = v707 + int32(1)
	v932 = v906
	goto L165
L167:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	if v726 <= v705 {
		v756 = v705
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L169
L169:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v833+v709<<(uint(int32(2))%32))))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v838 != 0 {
		goto L185
	} else {
		goto L186
	}
L170:
	;
	v773 = v756 + int32(1)
	v774 = F_bms_is_member(m, v773, v425)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L2
	} else {
		goto L176
	}
L171:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v731 = v705
	goto L172
L172:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v728+v731<<(uint(int32(2))%32))))
	if v750 != 0 {
		v756 = v731
		goto L170
	} else {
		goto L174
	}
L173:
	;
	v756 = v726
	goto L170
L174:
	;
	v752 = v731 + int32(1)
	if v752 != v726 {
		v731 = v752
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	if v774 != 0 {
		v923 = v703
		v925 = v773
		v927 = v707
		v932 = v712
		goto L165
	} else {
		goto L177
	}
L177:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v777 <= v703 {
		v804 = v703
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v823 = int32(2)
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v776+v804<<(uint(v823)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v822+v707<<(uint(v823)%32)))) = v829
	v897 = v804 + int32(1)
	v899 = v773
	v906 = v712
	goto L166
L179:
	;
	v779 = v703
	goto L180
L180:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v776+v779<<(uint(int32(2))%32))))
	if v800 != 0 {
		v804 = v779
		goto L178
	} else {
		goto L182
	}
L181:
	;
	v804 = v777
	goto L178
L182:
	;
	v802 = v779 + int32(1)
	if v802 != v777 {
		v779 = v802
		goto L180
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v887 != 0 {
		goto L198
	} else {
		goto L199
	}
L185:
	;
	v839 = F_make_colname_unique(m, v837, v22, v123)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L2
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v879+v707<<(uint(int32(2))%32)))) = v837
	v886 = v712
	goto L184
L188:
	;
	v842 = v707 << (uint(int32(2)) % 32)
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v842+v843))) = v839
	v846 = int32(1)
	if v712&v846 != 0 {
		v886 = v846
		goto L184
	} else {
		goto L189
	}
L189:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v849+v842)))
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837))))
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851))))
	if v855 == int32(0) {
		v874 = v854
		v875 = v855
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v886 = base.B2i32(v875-v874 != int32(0))
	goto L184
L191:
	;
	goto L190
L192:
	;
	if v854 != v855 {
		v874 = v854
		v875 = v855
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v859 = v851
	v860 = v837
	goto L194
L194:
	;
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860)+1)))
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859)+1)))
	if v864 == int32(0) {
		v874 = v863
		v875 = v864
		goto L191
	} else {
		goto L196
	}
L195:
	;
	v874 = v863
	v875 = v864
	goto L191
L196:
	;
	v867 = int32(1)
	if v863 == v864 {
		v859 = v859 + v867
		v860 = v860 + v867
		goto L194
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v888+v707<<(uint(int32(2))%32))))
	v895 = F_hash_search(m, v887, v892, int32(1), int32(0))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L2
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v897 = v703
	v899 = v705
	v906 = v886
	goto L166
L201:
	;
	goto L200
L202:
	;
	goto L164
L203:
	;
	F_hash_destroy(m, v963)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L2
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v971 = base.B2i32(v968 != int32(0)) & v954
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+20)) = uint8(v971)
	v94 = v94 + int32(1)
	goto L22
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+48)) = int32(0)
	goto L205
}
func F_set_pathtarget_cost_width(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int64
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v68 float64
	_ = v68
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	v3 = int32(0)
	v9 = int64(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v9
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 == v3 {
		v85 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v88 = int64(1073741823)
	if v88 <= v85 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v23 <= int32(0) {
		v85 = v9
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = v3
	v38 = v9
	goto L4
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v34<<(uint(int32(2))%32))))
	v46 = F_get_expr_width(m, l0, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v85 = v72
	goto L1
L6:
	;
	return int32(0)
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v51 != int32(6) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l0
	v55 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(24)))) = v55
	*(*int64)(unsafe.Add(mBase, uint32(v14+int32(16)))) = v55
	v61 = F_cost_qual_eval_walker(m, v45, v14+int32(8))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v72 = v38 + base.I64_extend_i32_s(v46)
	v74 = v34 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v74 < v75 {
		v34 = v74
		v38 = v72
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v65 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = base.F64_add(v64, v65)
	v68 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = base.F64_add(v63, v68)
	goto L10
L12:
	;
	goto L5
L13:
	;
	v91 = v88
	goto L15
L14:
	;
	v91 = v85
	goto L15
L15:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(l1)+32)) = uint32(v91)
	m.G0 = v14 + int32(32)
	return l1
}
func F_set_rtable_names(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
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
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v313 int32
	_ = v313
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = int64(292057776192)
	v24 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v31 = F_hash_create(m, int32(172999), v27, v16+int32(32), int32(1048))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v16 + int32(80)
	return
L4:
	;
	return
L5:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v122 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v45 = v4
	goto L9
L9:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v45<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v56 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v106 = v45 + int32(1)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v106 < v107 {
		v45 = v106
		goto L9
	} else {
		goto L21
	}
L12:
	;
	v59 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v60 <= v59 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v66 = v59
	goto L14
L14:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v66<<(uint(int32(2))%32))))
	if v80 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L11
L16:
	;
	v84 = F_hash_search(m, v31, v80, int32(1), v16+int32(31))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v89 = v66 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v89 < v90 {
		v66 = v89
		goto L14
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+64)) = int32(0)
	goto L18
L20:
	;
	goto L15
L21:
	;
	goto L10
L22:
	;
	F_hash_destroy(m, v31)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L4
	} else {
		goto L74
	}
L23:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v125 <= int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v139 = v4
	v140 = int32(1)
	goto L25
L25:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142+v139<<(uint(int32(2))%32))))
	v148 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v148 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L22
L27:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if l2 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L29
L31:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v290 = F_lappend(m, v289, v281)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L72
	}
L32:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	if v156 != 0 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v153 = F_bms_is_member(m, v140, l2)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v153 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v281 = int32(0)
	goto L31
L36:
	;
	if v166 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v166 = v157
	goto L36
L38:
	;
	goto L39
L39:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	switch v159 {
	case 0:
		goto L41
	default:
		goto L40
	case 2:
		v281 = int32(0)
		goto L31
	}
L40:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v166 = v164
	goto L36
L41:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v146)+16))
	v161 = F_get_rel_name(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v166 = v161
	goto L36
L43:
	;
	v281 = int32(0)
	goto L31
L44:
	;
	goto L45
L45:
	;
	v173 = F_hash_search(m, v31, v166, int32(1), v16+int32(31))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+31)))
	if v175 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v178 = F_strlen(m, v166)
	mBase = m.M
	v181 = F_palloc(m, v178+int32(16))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L50
	}
L48:
	;
	v265 = v166
	v273 = v173
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+64)) = int32(0)
	v281 = v265
	goto L31
L50:
	;
	v186 = v178
	goto L51
L51:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v173)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+64)) = v196 + int32(1)
	if v186 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v265 = v201
	v273 = v257
	goto L49
L53:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v173)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v202
	v208 = F_pg_sprintf(m, v186+v201, int32(486910), v16+int32(16))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L57
	}
L54:
	;
	v200 = F__emscripten_memcpy_bulkmem(m, v181, v166, v186)
	mBase = m.M
	v201 = v200
	goto L56
L55:
	;
	v201 = v181
	goto L56
L56:
	;
	goto L53
L57:
	;
	v210 = F_strlen(m, v201)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v210) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v216 = v186
	goto L61
L59:
	;
	v244 = v186
	goto L60
L60:
	;
	v257 = F_hash_search(m, v31, v201, int32(1), v16+int32(31))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L70
	}
L61:
	;
	v228 = F_pg_mbcliplen(m, v166, v216, v216-int32(1))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L63
	}
L62:
	;
	v244 = v228
	goto L60
L63:
	;
	if v228 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v173)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v232
	v236 = F_pg_sprintf(m, v228+v231, int32(486910), v16)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L68
	}
L65:
	;
	v230 = F__emscripten_memcpy_bulkmem(m, v201, v166, v228)
	mBase = m.M
	v231 = v230
	goto L67
L66:
	;
	v231 = v201
	goto L67
L67:
	;
	goto L64
L68:
	;
	v238 = F_strlen(m, v231)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v238) {
		v216 = v228
		goto L61
	} else {
		goto L69
	}
L69:
	;
	goto L62
L70:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+31)))
	if v259 != 0 {
		v186 = v244
		goto L51
	} else {
		goto L71
	}
L71:
	;
	goto L52
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v290
	v293 = int32(1)
	v296 = v139 + v293
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v296 < v297 {
		v139 = v296
		v140 = v140 + v293
		goto L25
	} else {
		goto L73
	}
L73:
	;
	goto L26
L74:
	;
	goto L3
}
func F_set_using_names(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	v24 = l1
	v25 = l2
	goto L1
L1:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v41 != int32(64) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v115)+60))
	if v251 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L4:
	;
	v207 = v170
	goto L40
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L15
	} else {
		goto L37
	}
L6:
	;
	switch v41 - int32(63) {
	case 0:
		goto L9
	default:
		goto L5
	case 2:
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v106 = int32(4)
	v107 = v103<<(uint(int32(2))%32) - v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v109)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113+v107)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	switch v119 - int32(63) {
	case 0:
		v140 = v106
		goto L18
	case 1:
		goto L19
	default:
		goto L20
	}
L9:
	;
	m.G0 = v21 + int32(48)
	return
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v46 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v49 <= int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v58 = int32(0)
	goto L13
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v58<<(uint(int32(2))%32))))
	F_set_using_names(m, l0, v75, v25)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L9
L15:
	;
	return
L16:
	;
	v79 = v58 + int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v79 < v80 {
		v58 = v79
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v140+v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+28)) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	switch v145 - int32(63) {
	case 0:
		v166 = v106
		goto L24
	case 1:
		goto L25
	default:
		goto L26
	}
L19:
	;
	v140 = int32(36)
	goto L18
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v127
	F_errmsg_internal(m, int32(508183), v21+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(516795), int32(5080), int32(156086))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v144+v166)))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+32)) = v168
	v170 = int32(0)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v115)+52))
	if v172 != 0 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v166 = int32(36)
	goto L24
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v153
	F_errmsg_internal(m, int32(508183), v21+int32(32))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(516795), int32(5087), int32(156086))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v176 = v173 << (uint(int32(2)) % 32)
	goto L32
L31:
	;
	v176 = v170
	goto L32
L32:
	;
	v177 = F_palloc0(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v177
	v180 = F_palloc0(m, v176)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+40)) = v180
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v115)+56))
	if v183 == int32(0) {
		v237 = v170
		goto L3
	} else {
		goto L35
	}
L35:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	if int32(0) < v186 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v237 = v170
	goto L3
L37:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v193
	F_errmsg_internal(m, int32(507644), v21)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(516795), int32(4364), int32(172943))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v222 = v207 << (uint(int32(2)) % 32)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v225+v222)))
	*(*int32)(unsafe.Add(mBase, uint32(v222+v223))) = v227
	v230 = v207 + int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	if v230 < v231 {
		v207 = v230
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v237 = v230
	goto L3
L42:
	;
	goto L41
L43:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
	v315 = int32(2)
	v318 = int32(4)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v313+v314<<(uint(v315)%32)-v318)))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v313+v321<<(uint(v315)%32)-v318)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v111)+40))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v330 != 0 {
		goto L52
	} else {
		goto L53
	}
L44:
	;
	v254 = int32(0)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	if v255 <= v254 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v261 = v254
	v262 = v237
	goto L46
L46:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v111)+40))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v115)+48))
	v278 = base.B2i32(v277 <= v261)
	if v277 <= v261 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L43
L48:
	;
	v279 = v262
	goto L50
L49:
	;
	v279 = v261
	goto L50
L50:
	;
	v280 = int32(2)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283+v261<<(uint(v280)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v276+v279<<(uint(v280)%32)))) = v287
	v291 = v261 + int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	if v291 < v292 {
		v261 = v291
		v262 = v262 + v278
		goto L46
	} else {
		goto L51
	}
L51:
	;
	goto L47
L52:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v451 == int32(0) {
		v623 = v25
		goto L82
	} else {
		goto L83
	}
L53:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v331 <= int32(0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v339 = int32(0)
	goto L55
L55:
	;
	v354 = v339 << (uint(int32(2)) % 32)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v354+v355)))
	if v357 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L52
L57:
	;
	v430 = v339 + int32(1)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v430 < v431 {
		v339 = v430
		goto L55
	} else {
		goto L81
	}
L58:
	;
	v360 = v354 + v329
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	if int32(0) < v361 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	if v365 < v361 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v393 = v354 + v328
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	if v394 <= int32(0) {
		goto L57
	} else {
		goto L71
	}
L62:
	;
	if v364 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v383 = v364
	v384 = v361
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384<<(uint(int32(2))%32)+v383-int32(4)))) = v357
	goto L61
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v327)+4)) = v379
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v360)))
	v383 = v379
	v384 = v382
	goto L64
L66:
	;
	v371 = F_palloc0(m, v361<<(uint(int32(2))%32))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L15
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v373 = int32(2)
	v377 = F_repalloc0(m, v364, v365<<(uint(v373)%32), v361<<(uint(v373)%32))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L15
	} else {
		goto L70
	}
L69:
	;
	v379 = v371
	goto L65
L70:
	;
	v379 = v377
	goto L65
L71:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	if v398 < v394 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	if v397 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v416 = v397
	v417 = v394
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v417<<(uint(int32(2))%32)+v416-int32(4)))) = v357
	goto L57
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v320)+4)) = v412
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	v416 = v412
	v417 = v415
	goto L74
L76:
	;
	v404 = F_palloc0(m, v394<<(uint(int32(2))%32))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L15
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v406 = int32(2)
	v410 = F_repalloc0(m, v397, v398<<(uint(v406)%32), v394<<(uint(v406)%32))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L15
	} else {
		goto L80
	}
L79:
	;
	v412 = v404
	goto L75
L80:
	;
	v412 = v410
	goto L75
L81:
	;
	goto L56
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327)+24)) = v623
	*(*int32)(unsafe.Add(mBase, uint32(v320)+24)) = v623
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	F_set_using_names(m, l0, v641, v623)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L15
	} else {
		goto L142
	}
L83:
	;
	v454 = F_list_copy(m, v25)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L15
	} else {
		goto L84
	}
L84:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	if v456 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	v459 = v457
	goto L87
L86:
	;
	v459 = int32(0)
	goto L87
L87:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v460 < v459 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v462 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v481 = v456
	goto L90
L90:
	;
	if v481 == int32(0) {
		v623 = v454
		goto L82
	} else {
		goto L97
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v475
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v481 = v478
	goto L90
L92:
	;
	v467 = F_palloc0(m, v459<<(uint(int32(2))%32))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L15
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v469 = int32(2)
	v473 = F_repalloc0(m, v462, v460<<(uint(v469)%32), v459<<(uint(v469)%32))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L15
	} else {
		goto L96
	}
L95:
	;
	v475 = v467
	goto L91
L96:
	;
	v475 = v473
	goto L91
L97:
	;
	v484 = int32(0)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	if v485 <= v484 {
		v623 = v454
		goto L82
	} else {
		goto L98
	}
L98:
	;
	v490 = v454
	v494 = v484
	goto L99
L99:
	;
	v507 = v494 << (uint(int32(2)) % 32)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v507+v508)))
	if v510 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v623 = v549
	goto L82
L101:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v481)+12))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v513+v507)))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)+4))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v517 == int32(0) {
		v529 = v516
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v543 = v510
	goto L103
L103:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v111)+44))
	v546 = F_lappend(m, v545, v543)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L15
	} else {
		goto L115
	}
L104:
	;
	v531 = F_make_colname_unique(m, v529, l0, v111)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L15
	} else {
		goto L110
	}
L105:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v517)+8))
	if v520 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	v523 = v521
	goto L108
L107:
	;
	v523 = int32(0)
	goto L108
L108:
	;
	if v523 <= v494 {
		v529 = v516
		goto L104
	} else {
		goto L109
	}
L109:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v520)+12))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v525+v507)))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
	v529 = v528
	goto L104
L110:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v533 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v537 = F_lappend(m, v536, v531)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L15
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v540+v507))) = v531
	v543 = v531
	goto L103
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v537
	goto L113
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111)+44)) = v546
	v549 = F_lappend(m, v490, v543)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L15
	} else {
		goto L116
	}
L116:
	;
	v551 = v507 + v329
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)))
	if int32(0) < v552 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	if v556 < v552 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	v584 = v507 + v328
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v584)))
	if int32(0) < v585 {
		goto L129
	} else {
		goto L130
	}
L120:
	;
	if v555 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	v574 = v555
	v575 = v552
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v575<<(uint(int32(2))%32)+v574-int32(4)))) = v543
	goto L119
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v552
	*(*int32)(unsafe.Add(mBase, uint32(v327)+4)) = v570
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v551)))
	v574 = v570
	v575 = v573
	goto L122
L124:
	;
	v562 = F_palloc0(m, v552<<(uint(int32(2))%32))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L15
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v564 = int32(2)
	v568 = F_repalloc0(m, v555, v556<<(uint(v564)%32), v552<<(uint(v564)%32))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L15
	} else {
		goto L128
	}
L127:
	;
	v570 = v562
	goto L123
L128:
	;
	v570 = v568
	goto L123
L129:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	if v589 < v585 {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	goto L131
L131:
	;
	v618 = v494 + int32(1)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	if v618 < v619 {
		v490 = v549
		v494 = v618
		goto L99
	} else {
		goto L141
	}
L132:
	;
	if v588 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v607 = v588
	v608 = v585
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v608<<(uint(int32(2))%32)+v607-int32(4)))) = v543
	goto L131
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v585
	*(*int32)(unsafe.Add(mBase, uint32(v320)+4)) = v603
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v584)))
	v607 = v603
	v608 = v606
	goto L134
L136:
	;
	v595 = F_palloc0(m, v585<<(uint(int32(2))%32))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L15
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v597 = int32(2)
	v601 = F_repalloc0(m, v588, v589<<(uint(v597)%32), v585<<(uint(v597)%32))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L15
	} else {
		goto L140
	}
L139:
	;
	v603 = v595
	goto L135
L140:
	;
	v603 = v601
	goto L135
L141:
	;
	goto L100
L142:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v24 = v644
	v25 = v623
	goto L1
}
func F_setop_load_group(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v4)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v24 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, v8, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L9
	}
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
	if v9&int32(2) == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	m.T0[v16].(func(*base.Module, int32))(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
	goto L1
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = F_ExecStoreMinimalTuple(m, v24, v26, int32(1))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v35 = int64(1)
	goto L11
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_ExecReScan(m, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v41 = m.T0[v40].(func(*base.Module, int32) int32)(m, l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v41
	if v41 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v46&int32(2) != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v50 = F_setop_compare_slots(m, v49, v41, l2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	if v50 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = v52 + int64(1)
	goto L11
}
func F_setval_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	F_do_setval(m, v3, v5, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_Int64GetDatum(m, v5)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_shdepDropDependency(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v58 int32
	_ = v58
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	v12 = m.G0
	v14 = v12 - int32(192)
	m.G0 = v14
	v18 = int32(1)
	if l1 <= int32(3591) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v87 = int32(3)
	v93 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	if v86 != 0 {
		goto L29
	} else {
		goto L30
	}
L2:
	;
	goto L1
L3:
	;
	v86 = int32(0)
	goto L2
L4:
	;
	if base.Ui32(l1-int32(2964)) < base.Ui32(int32(4)) {
		v86 = v18
		goto L2
	} else {
		goto L27
	}
L5:
	;
	if l1 <= int32(2670) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	if l1 <= int32(5999) {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	switch l1 - int32(1213) {
	case 0, 1, 19, 20, 47, 48, 49:
		v86 = v18
		goto L2
	case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
		goto L3
	default:
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v30 = l1 - int32(2671)
	if base.Ui32(int32(27)) < base.Ui32(v30) {
		goto L4
	} else {
		goto L13
	}
L11:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l1-int32(2396)) {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v86 = v18
	goto L2
L13:
	;
	if int32(1)<<(uint(v30)%32)&int32(226492515) == int32(0) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v86 = v18
	goto L2
L15:
	;
	if base.Ui32(l1-int32(3592)) < base.Ui32(int32(2)) {
		v86 = v18
		goto L2
	} else {
		goto L25
	}
L16:
	;
	v42 = l1 - int32(4177)
	if base.Ui32(int32(9)) < base.Ui32(v42) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	switch l1 - int32(6243) {
	case 0, 1, 2, 3, 4, 59, 60:
		v86 = v18
		goto L2
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
		goto L3
	default:
		goto L21
	}
L19:
	;
	if int32(1)<<(uint(v42)%32)&int32(963) == int32(0) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v86 = v18
	goto L2
L21:
	;
	if base.Ui32(l1-int32(6000)) < base.Ui32(int32(3)) {
		v86 = v18
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v58 = l1 - int32(6100)
	if base.Ui32(int32(15)) < base.Ui32(v58) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	if int32(1)<<(uint(v58)%32)&int32(49153) != 0 {
		v86 = v18
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L3
L25:
	;
	if base.Ui32(int32(2)) <= base.Ui32(l1-int32(4060)) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v86 = v18
	goto L2
L27:
	;
	if base.Ui32(l1-int32(2846)) < base.Ui32(int32(2)) {
		v86 = v18
		goto L2
	} else {
		goto L28
	}
L28:
	;
	goto L3
L29:
	;
	v94 = int32(0)
	goto L31
L30:
	;
	v94 = v93
	goto L31
L31:
	;
	F_ScanKeyInit(m, v14, int32(1), v87, int32(184), v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return
L33:
	;
	F_ScanKeyInit(m, v14+int32(48), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v106 = int32(3)
	F_ScanKeyInit(m, v14+int32(96), v106, v106, int32(184), l2)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	if l4 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v113 = int32(4)
	F_ScanKeyInit(m, v14+int32(144), v113, int32(3), int32(65), l3)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L32
	} else {
		goto L39
	}
L37:
	;
	v121 = v87
	goto L38
L38:
	;
	v125 = F_systable_beginscan(m, l0, int32(1232), int32(1), int32(0), v121, v14)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L32
	} else {
		goto L40
	}
L39:
	;
	v121 = v113
	goto L38
L40:
	;
	v127 = F_systable_getnext(m, v125)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	if v127 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v130 = v127
	goto L45
L43:
	;
	goto L44
L44:
	;
	F_systable_endscan(m, v125)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L32
	} else {
		goto L63
	}
L45:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+22)))
	v142 = v140 + v141
	if l5 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L44
L47:
	;
	v153 = F_systable_getnext(m, v125)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L32
	} else {
		goto L61
	}
L48:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	if v143 != l5 {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if l6 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+20))
	if v145 != l6 {
		goto L47
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if l7 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v142)+24)))
	if l7 != v147 {
		goto L47
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	F_CatalogTupleDelete(m, l0, v130+int32(4))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L32
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	goto L47
L61:
	;
	if v153 != 0 {
		v130 = v153
		goto L45
	} else {
		goto L62
	}
L62:
	;
	goto L46
L63:
	;
	m.G0 = v14 + int32(192)
	return
}
func F_shell_archive_init(m *base.Module) int32 {
	return int32(792624)
}
func F_shell_archive_shutdown(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v4 = F_errstart(m, int32(14), int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		if v4 != 0 {
			F_errmsg_internal(m, int32(256306), int32(0))
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				F_errfinish(m, int32(521990), int32(141), int32(255741))
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			return
		}
	}
}
func F_shift_jis_2004_to_euc_jis_2004(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v15, v16, v17, int32(41), int32(5))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v17 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v305 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v299))) = uint8(v305)
	return v296 - v14
L4:
	;
	v296 = v14
	v299 = v13
	goto L3
L5:
	;
	goto L6
L6:
	;
	v26 = v17
	v28 = v14
	v31 = v13
	goto L7
L7:
	;
	v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28))))
	if int32(0) <= v37 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v296 = v284
	v299 = v287
	goto L3
L9:
	;
	if int32(0) < v291 {
		v26 = v291
		v28 = v284
		v31 = v287
		goto L7
	} else {
		goto L114
	}
L10:
	;
	if v37 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v53 = F_pg_encoding_verifymbchar(m, int32(41), v28, v26)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v12 != 0 {
		v296 = v28
		v299 = v31
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v37)
	v46 = int32(1)
	v284 = v28 + v46
	v287 = v31 + v46
	v291 = v26 - v46
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(41), v28, v26)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	if base.Ui32(v26) < base.Ui32(v53) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v12 != 0 {
		v296 = v28
		v299 = v31
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if base.Ui32(int32(62)) < base.Ui32((v37+int32(95))&int32(255)) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	F_report_invalid_encoding(m, int32(41), v28, v26)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v284 = v28 + v53
	v287 = v278
	v291 = v26 - v53
	goto L9
L25:
	;
	v278 = v270 + int32(2)
	goto L24
L26:
	;
	if v53 != int32(2) {
		v278 = v31
		goto L24
	} else {
		goto L29
	}
L27:
	;
	if v53 != int32(1) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)) = uint8(v37)
	v68 = int32(142)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v68)
	v270 = v31
	goto L25
L29:
	;
	v73 = v37 & int32(255)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v75 = base.I32_extend8_s(v74)
	if v37 == int32(-128) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v261 = int32(96)
	v262 = v259 - v261
	*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)) = uint8(v262)
	v265 = v257 - v261
	*(*uint8)(unsafe.Add(mBase, uint32(v260))) = uint8(v265)
	v270 = v260
	goto L25
L31:
	;
	v257 = v73<<(uint(int32(1))%32) + v250 - int32(256)
	v259 = v251
	v260 = v31
	goto L30
L32:
	;
	if v37&int32(-16) == int32(-32) {
		goto L47
	} else {
		goto L48
	}
L33:
	;
	if base.Ui32(int32(-97)) < base.Ui32(v37) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v81 = v74 + int32(-64)
	if base.Ui32(v81) <= base.Ui32(int32(62)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v250 = int32(-1)
	v251 = v74 - int32(63)
	goto L31
L36:
	;
	goto L37
L37:
	;
	if v75 < int32(-97) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v12 != 0 {
		v296 = v28
		v299 = v31
		goto L3
	} else {
		goto L44
	}
L39:
	;
	v99 = v81
	v100 = int32(-1)
	goto L41
L40:
	;
	if base.Ui32(int32(93)) < base.Ui32((v75+int32(97))&int32(255)) {
		goto L38
	} else {
		goto L42
	}
L41:
	;
	if int32(0) <= v99 {
		v250 = v100
		v251 = v99
		goto L31
	} else {
		goto L43
	}
L42:
	;
	v99 = v74 - int32(158)
	v100 = int32(0)
	goto L41
L43:
	;
	goto L38
L44:
	;
	F_report_invalid_encoding(m, int32(41), v28, v26)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	v257 = v73<<(uint(int32(1))%32) + v243 - int32(384)
	v259 = v244
	v260 = v31
	goto L30
L47:
	;
	v113 = v74 + int32(-64)
	if base.Ui32(v113) <= base.Ui32(int32(62)) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	if v37&int32(-4) == int32(-16) {
		goto L63
	} else {
		goto L64
	}
L50:
	;
	v243 = int32(-1)
	v244 = v74 - int32(63)
	goto L46
L51:
	;
	goto L52
L52:
	;
	if v75 < int32(-97) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v12 != 0 {
		v296 = v28
		v299 = v31
		goto L3
	} else {
		goto L59
	}
L54:
	;
	v131 = v113
	v132 = int32(-1)
	goto L56
L55:
	;
	if base.Ui32(int32(93)) < base.Ui32((v75+int32(97))&int32(255)) {
		goto L53
	} else {
		goto L57
	}
L56:
	;
	if int32(0) <= v131 {
		v243 = v132
		v244 = v131
		goto L46
	} else {
		goto L58
	}
L57:
	;
	v131 = v74 - int32(158)
	v132 = int32(0)
	goto L56
L58:
	;
	goto L53
L59:
	;
	F_report_invalid_encoding(m, int32(41), v28, v26)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
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
	v239 = int32(143)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v239)
	v257 = v238
	v259 = v237
	v260 = v31 + int32(1)
	goto L30
L62:
	;
	switch v73 - int32(240) {
	case 0:
		goto L98
	case 1:
		goto L101
	case 2:
		goto L100
	default:
		goto L99
	}
L63:
	;
	v145 = v74 + int32(-64)
	if base.Ui32(v145) <= base.Ui32(int32(62)) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	if base.Ui32((v37+int32(12))&int32(255)) <= base.Ui32(int32(8)) {
		goto L78
	} else {
		goto L79
	}
L66:
	;
	v219 = int32(1)
	v220 = v74 - int32(63)
	goto L62
L67:
	;
	goto L68
L68:
	;
	if v75 < int32(-97) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v12 != 0 {
		v296 = v28
		v299 = v31
		goto L3
	} else {
		goto L75
	}
L70:
	;
	v163 = v145
	v164 = int32(1)
	goto L72
L71:
	;
	if base.Ui32(int32(93)) < base.Ui32((v75+int32(97))&int32(255)) {
		goto L69
	} else {
		goto L73
	}
L72:
	;
	if int32(0) <= v163 {
		v219 = v164
		v220 = v163
		goto L62
	} else {
		goto L74
	}
L73:
	;
	v163 = v74 - int32(158)
	v164 = int32(0)
	goto L72
L74:
	;
	goto L69
L75:
	;
	F_report_invalid_encoding(m, int32(41), v28, v26)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	if v37 == int32(-12) {
		goto L94
	} else {
		goto L95
	}
L78:
	;
	v179 = v74 + int32(-64)
	if base.Ui32(v179) <= base.Ui32(int32(62)) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	if v12 != 0 {
		v296 = v28
		v299 = v31
		goto L3
	} else {
		goto L92
	}
L81:
	;
	v209 = int32(1)
	v210 = v74 - int32(63)
	goto L77
L82:
	;
	goto L83
L83:
	;
	if v75 < int32(-97) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v12 != 0 {
		v296 = v28
		v299 = v31
		goto L3
	} else {
		goto L90
	}
L85:
	;
	v197 = v179
	v198 = int32(1)
	goto L87
L86:
	;
	if base.Ui32(int32(93)) < base.Ui32((v75+int32(97))&int32(255)) {
		goto L84
	} else {
		goto L88
	}
L87:
	;
	if int32(0) <= v197 {
		v209 = v198
		v210 = v197
		goto L77
	} else {
		goto L89
	}
L88:
	;
	v197 = v74 - int32(158)
	v198 = int32(0)
	goto L87
L89:
	;
	goto L84
L90:
	;
	F_report_invalid_encoding(m, int32(41), v28, v26)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_report_invalid_encoding(m, int32(41), v28, v26)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	if v209 != 0 {
		v237 = v210
		v238 = int32(15)
		goto L61
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v237 = v210
	v238 = v73<<(uint(int32(1))%32) - v209 - int32(410)
	goto L61
L97:
	;
	goto L96
L98:
	;
	if v219 != 0 {
		goto L111
	} else {
		goto L112
	}
L99:
	;
	if v219 != 0 {
		goto L108
	} else {
		goto L109
	}
L100:
	;
	if v219 != 0 {
		goto L105
	} else {
		goto L106
	}
L101:
	;
	if v219 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v225 = int32(3)
	goto L104
L103:
	;
	v225 = int32(4)
	goto L104
L104:
	;
	v237 = v220
	v238 = v225
	goto L61
L105:
	;
	v228 = int32(5)
	goto L107
L106:
	;
	v228 = int32(12)
	goto L107
L107:
	;
	v237 = v220
	v238 = v228
	goto L61
L108:
	;
	v231 = int32(13)
	goto L110
L109:
	;
	v231 = int32(14)
	goto L110
L110:
	;
	v237 = v220
	v238 = v231
	goto L61
L111:
	;
	v234 = int32(1)
	goto L113
L112:
	;
	v234 = int32(8)
	goto L113
L113:
	;
	v237 = v220
	v238 = v234
	goto L61
L114:
	;
	goto L8
}
func F_shortest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v552 int32
	_ = v552
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v628 int32
	_ = v628
	v8 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if l5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l6 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v22 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v628
L8:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+8)))
	if v134&int32(2) != 0 {
		goto L48
	} else {
		goto L49
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v28 = v25 + v22<<(uint(int32(3))%32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v29 < int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+66)))
	v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+64)))
	if base.I32_extend16_s(v33) <= base.I32_extend16_s(v32) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if l5 == int32(0) {
		v628 = v124
		goto L7
	} else {
		goto L46
	}
L12:
	;
	v38 = l2
	goto L14
L13:
	;
	v38 = int32(0)
	goto L14
L14:
	;
	if l2 == l3 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v41 = v38
	goto L17
L16:
	;
	v41 = int32(0)
	goto L17
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v29 == v42 {
		v124 = v41
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v48 = v42 - v29
	v49 = base.I32_div_u_s((l4-l2)>>(uint(int32(2))%32), v48)
	if base.Ui32(v49) < base.Ui32(v32) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v51 = v49
	goto L21
L20:
	;
	v51 = v32
	goto L21
L21:
	;
	if v32 == int32(256) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v54 = v49
	goto L24
L23:
	;
	v54 = v51
	goto L24
L24:
	;
	if base.Ui32(l3) <= base.Ui32(l2) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v65 = int32(0)
	goto L27
L26:
	;
	v60 = int32(1)
	v62 = base.I32_div_u_s((l3-l2)>>(uint(int32(2))%32)-v60, v48)
	v65 = v62 + v60
	goto L27
L27:
	;
	if base.Ui32(v33) < base.Ui32(v65) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v67 = v65
	goto L30
L29:
	;
	v67 = v33
	goto L30
L30:
	;
	if base.Ui32(v54) < base.Ui32(v67) {
		v628 = v8
		goto L7
	} else {
		goto L31
	}
L31:
	;
	if v67 == int32(0) {
		v124 = l2
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v71 = int32(2)
	v85 = l2
	v86 = int32(0)
	goto L34
L33:
	;
	if base.Ui32(v67) <= base.Ui32(v103) {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	if v54 == v86 {
		v103 = v54
		v104 = v85
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v103 = v67
	v104 = v99
	goto L33
L36:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+420))
	v95 = m.T0[v94].(func(*base.Module, int32, int32, int32) int32)(m, v44+v29<<(uint(v71)%32), v85, v48)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return int32(0)
L38:
	;
	if v95 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v103 = v86
	v104 = v85
	goto L33
L40:
	;
	goto L41
L41:
	;
	v99 = v85 + v48<<(uint(v71)%32)
	v101 = v86 + int32(1)
	if v101 != v67 {
		v85 = v99
		v86 = v101
		goto L34
	} else {
		goto L42
	}
L42:
	;
	goto L35
L43:
	;
	v108 = v104
	goto L45
L44:
	;
	v108 = int32(0)
	goto L45
L45:
	;
	v124 = v108
	goto L11
L46:
	;
	if v124 == int32(0) {
		v628 = v124
		goto L7
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = l2
	return v124
L48:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133)+44))
	v142 = (l3 - l2) >> (uint(int32(2)) % 32)
	if base.B2i32(v137 != int32(256))&base.B2i32(base.Ui32(v137) < base.Ui32(v142)) != 0 {
		v628 = v8
		goto L7
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v159 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v159 < v167 {
		goto L60
	} else {
		goto L61
	}
L51:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v133)+40))
	if (l4-l2)>>(uint(int32(2))%32) < v145 {
		v628 = v8
		goto L7
	} else {
		goto L52
	}
L52:
	;
	if base.Ui32(v142) < base.Ui32(v145) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v154 = l2 + v145<<(uint(int32(2))%32)
	goto L55
L54:
	;
	v154 = l3
	goto L55
L55:
	;
	if l5 == int32(0) {
		v628 = v154
		goto L7
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = l2
	return v154
L57:
	;
	if v380 == int32(0) {
		v628 = v8
		goto L7
	} else {
		goto L94
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+20)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = int64(0)
	v380 = v359
	goto L57
L59:
	;
	v334 = int32(0)
	goto L91
L60:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+8)))
	if v171&int32(1) != 0 {
		v326 = v170
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v175 = F_getvacant(m, l0, l1, l2, l2)
	mBase = m.M
	if v175 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L62
L64:
	;
	v380 = int32(0)
	goto L57
L65:
	;
	goto L66
L66:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if int32(0) < v179 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v183 = int32(0)
	goto L70
L68:
	;
	goto L69
L69:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+12))
	v222 = v215 + int32(base.Ui32(v217)>>(uint(int32(3))%32))&int32(536870908)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v224 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v223 | v224<<(uint(v217)%32)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v229 == v224 {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	*(*int32)(unsafe.Add(mBase, uint32(v194+v183<<(uint(int32(2))%32)))) = int32(0)
	v201 = v183 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v201 < v202 {
		v183 = v201
		goto L70
	} else {
		goto L72
	}
L71:
	;
	goto L69
L72:
	;
	goto L71
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+8)) = int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+4)) = v308
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v319 <= int32(0) {
		v359 = v175
		goto L58
	} else {
		goto L90
	}
L74:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v308 = v232
	goto L73
L75:
	;
	goto L76
L76:
	;
	if v229 <= int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v308 = int32(0)
	goto L73
L78:
	;
	goto L79
L79:
	;
	v237 = v229 & int32(3)
	v238 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v229) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v244 = v238
	v247 = v238
	v253 = v159
	goto L83
L81:
	;
	v271 = v238
	v274 = v238
	goto L82
L82:
	;
	if v237 == int32(0) {
		v308 = v274
		goto L73
	} else {
		goto L86
	}
L83:
	;
	v257 = v228 + v244<<(uint(int32(2))%32)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v265 = v258 ^ (v259 ^ (v260 ^ (v261 ^ v247)))
	v266 = int32(4)
	v267 = v244 + v266
	v269 = v253 + v266
	if v269 != v229&int32(2147483644) {
		v244 = v267
		v247 = v265
		v253 = v269
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v271 = v267
	v274 = v265
	goto L82
L85:
	;
	goto L84
L86:
	;
	v284 = v271
	v287 = v274
	v292 = v159
	goto L87
L87:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v228+v284<<(uint(int32(2))%32))))
	v299 = v298 ^ v287
	v300 = int32(1)
	v303 = v292 + v300
	if v303 != v237 {
		v284 = v284 + v300
		v287 = v299
		v292 = v303
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v308 = v299
	goto L73
L89:
	;
	goto L88
L90:
	;
	v326 = v175
	goto L59
L91:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v345+v334<<(uint(int32(5))%32))+20)) = int32(0)
	v352 = v334 + int32(1)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v352 < v353 {
		v334 = v352
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v359 = v326
	goto L58
L93:
	;
	goto L92
L94:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v383 == l2 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v409 = F_miss(m, l0, l1, v380, base.I32_extend16_s(v407), l2, l2)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L37
	} else {
		goto L102
	}
L96:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v389 = int32(1)
	v394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v385+(v386^int32(-1))&v389<<(uint(v389)%32))+20)))
	v407 = v394
	goto L95
L97:
	;
	goto L98
L98:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l2-int32(4))))
	if base.Ui32(v397) <= base.Ui32(int32(2047)) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400+v397<<(uint(int32(1))%32)))))
	v407 = v404
	goto L95
L100:
	;
	goto L101
L101:
	;
	v405 = F_pg_reg_getcolor(m, v17, v397)
	mBase = m.M
	v407 = v405
	goto L95
L102:
	;
	if v409 == int32(0) {
		v628 = v8
		goto L7
	} else {
		goto L103
	}
L103:
	;
	v414 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v409)+20)) = l2
	v429 = v409
	v430 = l2
	goto L104
L104:
	;
	if base.Ui32(l4+base.B2i32(l4 != v16)<<(uint(v414)%32)) <= base.Ui32(v430) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	if l5 != 0 {
		goto L121
	} else {
		goto L122
	}
L106:
	;
	goto L105
L107:
	;
	v473 = v430
	v474 = v429
	goto L106
L108:
	;
	goto L109
L109:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	if base.Ui32(v438) <= base.Ui32(int32(2047)) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v429)+24))
	v449 = base.I32_extend16_s(v447)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v448+v449<<(uint(int32(2))%32))))
	if v453 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v445 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v441+v438<<(uint(int32(1))%32)))))
	v447 = v445
	goto L110
L112:
	;
	goto L113
L113:
	;
	v446 = F_pg_reg_getcolor(m, v17, v438)
	mBase = m.M
	v447 = v446
	goto L110
L114:
	;
	v458 = F_miss(m, l0, l1, v429, v449, v430+int32(4), l2)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L37
	} else {
		goto L117
	}
L115:
	;
	v462 = v453
	goto L116
L116:
	;
	v464 = v430 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v462)+20)) = v464
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+8)))
	if v466&int32(2) == int32(0) {
		v429 = v462
		v430 = v464
		goto L104
	} else {
		goto L119
	}
L117:
	;
	if v458 == int32(0) {
		v628 = v8
		goto L7
	} else {
		goto L118
	}
L118:
	;
	v462 = v458
	goto L116
L119:
	;
	if base.Ui32(v464) < base.Ui32(l3+base.B2i32(l3 != v16)<<(uint(v414)%32)) {
		v429 = v462
		v430 = v464
		goto L104
	} else {
		goto L120
	}
L120:
	;
	v473 = v464
	v474 = v462
	goto L106
L121:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v476 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v474)+8))
	v575 = v573 & int32(2)
	if base.Ui32(v473) <= base.Ui32(l3) {
		goto L154
	} else {
		goto L155
	}
L124:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v480 = v479
	goto L126
L125:
	;
	v480 = v476
	goto L126
L126:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v481 <= int32(0) {
		v552 = v480
		goto L127
	} else {
		goto L128
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v552
	goto L123
L128:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v481&int32(1) != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484)+8)))
	if v487&int32(8) != 0 {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v499 = v484
	v500 = v480
	v502 = v481
	goto L131
L131:
	;
	if v481 == int32(1) {
		v552 = v500
		goto L127
	} else {
		goto L138
	}
L132:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v484)+20))
	if base.Ui32(v480) < base.Ui32(v490) {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	v493 = v480
	goto L134
L134:
	;
	v499 = v484 + int32(32)
	v500 = v493
	v502 = v481 - int32(1)
	goto L131
L135:
	;
	v492 = v490
	goto L137
L136:
	;
	v492 = v480
	goto L137
L137:
	;
	v493 = v492
	goto L134
L138:
	;
	v512 = v499
	v515 = v500
	v518 = v502
	goto L139
L139:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+8)))
	if v520&int32(8) != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v552 = v534
	goto L127
L141:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v512)+20))
	if base.Ui32(v515) < base.Ui32(v523) {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	v526 = v515
	goto L143
L143:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+40)))
	if v528&int32(8) != 0 {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	v525 = v523
	goto L146
L145:
	;
	v525 = v515
	goto L146
L146:
	;
	v526 = v525
	goto L143
L147:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v512)+52))
	if base.Ui32(v526) < base.Ui32(v531) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v534 = v526
	goto L149
L149:
	;
	v538 = int32(2)
	if v538 < v518 {
		v512 = v512 - int32(-64)
		v515 = v534
		v518 = v518 - v538
		goto L139
	} else {
		goto L153
	}
L150:
	;
	v533 = v531
	goto L152
L151:
	;
	v533 = v526
	goto L152
L152:
	;
	v534 = v533
	goto L149
L153:
	;
	goto L140
L154:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v473 != v582 {
		v607 = v575
		goto L158
	} else {
		goto L159
	}
L155:
	;
	if v575 == int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	return v473 - int32(4)
L157:
	;
	if l6 == int32(0) {
		v628 = v8
		goto L7
	} else {
		goto L168
	}
L158:
	;
	if v607 != 0 {
		goto L165
	} else {
		goto L166
	}
L159:
	;
	if l4 != v582 {
		v607 = v575
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v592 = int32(*(*int16)(unsafe.Add(mBase, uint32(v585+(v586^int32(-1))&int32(2))+24)))
	v593 = F_miss(m, l0, l1, v474, v592, v473, l2)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L37
	} else {
		goto L161
	}
L161:
	;
	if v593 == int32(0) {
		goto L157
	} else {
		goto L162
	}
L162:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v593)+8))
	v599 = v597 & int32(2)
	if l6 == int32(0) {
		v607 = v599
		goto L158
	} else {
		goto L163
	}
L163:
	;
	if v599 != 0 {
		v607 = v599
		goto L158
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(1)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v593)+8))
	v607 = v604 & int32(2)
	goto L158
L165:
	;
	v610 = v473
	goto L167
L166:
	;
	v610 = int32(0)
	goto L167
L167:
	;
	return v610
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(1)
	v628 = v8
	goto L7
}
func F_show_incremental_sort_group_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v66 int32
	_ = v66
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
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v170 int64
	_ = v170
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v183 int64
	_ = v183
	var v190 int32
	_ = v190
	var v195 int64
	_ = v195
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v208 int64
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int64
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int64
	_ = v239
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int64
	_ = v267
	var v269 int32
	_ = v269
	var v272 int64
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v284 int64
	_ = v284
	var v287 int64
	_ = v287
	var v288 int64
	_ = v288
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int64
	_ = v312
	var v314 int32
	_ = v314
	var v317 int64
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	v13 = m.G0
	v15 = v13 - int32(160)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v17&int32(1) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L5
L2:
	;
	v35 = int32(0)
	v36 = v17
	goto L3
L3:
	;
	if v36&int32(2) != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v32 = F_lappend(m, int32(0), v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[436]))
	goto L7
L7:
	;
	goto L4
L8:
	;
	return
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v35 = v32
	v36 = v34
	goto L3
L10:
	;
	goto L14
L11:
	;
	v53 = v35
	v54 = v36
	goto L12
L12:
	;
	if v54&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v50 = F_lappend(m, v35, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L17
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[437]))
	goto L16
L16:
	;
	goto L13
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v53 = v50
	v54 = v52
	goto L12
L18:
	;
	goto L22
L19:
	;
	v71 = v53
	v72 = v54
	goto L20
L20:
	;
	if v72&int32(8) != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v68 = F_lappend(m, v53, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L25
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[438]))
	goto L24
L24:
	;
	goto L21
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v71 = v68
	v72 = v70
	goto L20
L26:
	;
	goto L30
L27:
	;
	v88 = v71
	goto L28
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v89 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v86 = F_lappend(m, v71, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L33
	}
L30:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[439]))
	goto L32
L32:
	;
	goto L29
L33:
	;
	v88 = v86
	goto L28
L34:
	;
	m.G0 = v15 + int32(160)
	return
L35:
	;
	if l2 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	F_initStringInfo(m, v15+int32(144))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L8
	} else {
		goto L75
	}
L38:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	F_appendStringInfoSpaces(m, v92, v93<<(uint(int32(1))%32))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+72)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = l1
	F_appendStringInfo(m, v98, int32(442852), v15-int32(-64))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	if v88 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if int64(0) < v170 {
		goto L61
	} else {
		goto L62
	}
L44:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_appendStringInfoString(m, v109, int32(778714))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L8
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if int32(1) < v116 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L43
L48:
	;
	v119 = int32(778358)
	goto L50
L49:
	;
	v119 = int32(778714)
	goto L50
L50:
	;
	F_appendStringInfoString(m, v113, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v122 = int32(0)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v123 <= v122 {
		goto L43
	} else {
		goto L52
	}
L52:
	;
	v132 = v122
	goto L53
L53:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+v132<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v138, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L8
	} else {
		goto L55
	}
L54:
	;
	goto L43
L55:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v132 < v146-int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_appendStringInfoString(m, v150, int32(778962))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L8
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v155 = v132 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v155 < v156 {
		v132 = v155
		goto L53
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	goto L54
L61:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v174 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v176 = int32(14278)
	goto L65
L62:
	;
	goto L63
L63:
	;
	v195 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if v195 <= int64(0) {
		goto L34
	} else {
		goto L69
	}
L64:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v180 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v176
	v183 = base.I64_div_s(v174, v173)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v176
	F_appendStringInfo(m, v179, int32(570053), v15+int32(32))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L8
	} else {
		goto L68
	}
L65:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	goto L63
L69:
	;
	v198 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v199 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v202 = int32(329506)
	goto L72
L70:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v205 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v202
	v208 = base.I64_div_s(v199, v198)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v202
	F_appendStringInfo(m, v204, int32(570053), v15)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L8
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	goto L70
L74:
	;
	goto L34
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = l1
	F_appendStringInfo(m, v15+int32(144), int32(144673), v15+int32(112))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
	F_ExplainOpenGroup(m, int32(144649), v227, int32(1), l3)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	v233 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	F_ExplainPropertyInteger(m, int32(94180), int32(0), v233, l3)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	F_ExplainPropertyList(m, int32(470501), v88, l3)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v239 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if int64(0) < v239 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v242 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v243 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	goto L84
L81:
	;
	goto L82
L82:
	;
	v284 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	if int64(0) < v284 {
		goto L93
	} else {
		goto L94
	}
L83:
	;
	F_initStringInfo(m, v15+int32(128))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L8
	} else {
		goto L87
	}
L84:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = int32(14278)
	F_appendStringInfo(m, v15+int32(128), int32(206192), v15+int32(96))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
	F_ExplainOpenGroup(m, int32(439340), v261, int32(1), l3)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	v267 = base.I64_div_s(v243, v242)
	F_ExplainPropertyInteger(m, int32(470540), int32(570190), v267, l3)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	v272 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_ExplainPropertyInteger(m, int32(470519), int32(570190), v272, l3)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	F_ExplainCloseGroup(m, int32(439340), int32(1), l3)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	goto L82
L93:
	;
	v287 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v288 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	goto L98
L94:
	;
	goto L95
L95:
	;
	F_ExplainCloseGroup(m, int32(144649), int32(1), l3)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L8
	} else {
		goto L106
	}
L96:
	;
	F_initStringInfo(m, v15+int32(128))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L8
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = int32(329506)
	F_appendStringInfo(m, v15+int32(128), int32(206192), v15+int32(80))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v15)+128))
	F_ExplainOpenGroup(m, int32(439340), v306, int32(1), l3)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	v312 = base.I64_div_s(v288, v287)
	F_ExplainPropertyInteger(m, int32(470540), int32(570190), v312, l3)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	v317 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExplainPropertyInteger(m, int32(470519), int32(570190), v317, l3)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	F_ExplainCloseGroup(m, int32(439340), int32(1), l3)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	goto L95
L106:
	;
	goto L34
}
func F_sigemptyset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	return
}
func F_slice_to(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v7 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = v10 - v7
	v38 = l1 - int32(8)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v39 < v36 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	if l1 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 < v7 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v12 < v10 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14-int32(4))))
	if v12 <= v19 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L2
L8:
	;
	return int32(0)
L9:
	;
	goto L10
L10:
	;
	F_pfree(m, l1-int32(8))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	return int32(0)
L13:
	;
	F_pfree(m, v38)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L11
	} else {
		goto L65
	}
L14:
	;
	v43 = F_repalloc(m, v38, v36+int32(29))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L11
	} else {
		goto L17
	}
L15:
	;
	v55 = v7
	v56 = v14
	v57 = l1
	goto L16
L16:
	;
	v58 = v55 + v56
	if v57 == v58 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if v43 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v36 + int32(20)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v55 = v50
	v56 = v51
	v57 = v43 + int32(8)
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202-int32(4)))) = v36
	return v202
L20:
	;
	v202 = v57
	goto L19
L21:
	;
	v62 = v57 + v36
	if base.Ui32(v58-v62) <= base.Ui32(int32(0)-v36<<(uint(int32(1))%32)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v69 = F___memcpy(m, v57, v58, v36)
	mBase = m.M
	v202 = v69
	goto L19
L23:
	;
	goto L24
L24:
	;
	v72 = (v57 ^ v58) & int32(3)
	if base.Ui32(v57) < base.Ui32(v58) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if v174 == int32(0) {
		goto L20
	} else {
		goto L61
	}
L26:
	;
	if base.Ui32(v152) <= base.Ui32(int32(3)) {
		v173 = v151
		v174 = v152
		v175 = v153
		goto L25
	} else {
		goto L57
	}
L27:
	;
	if v72 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	if v72 != 0 {
		v134 = v36
		goto L40
	} else {
		goto L41
	}
L30:
	;
	v173 = v58
	v174 = v36
	v175 = v57
	goto L25
L31:
	;
	goto L32
L32:
	;
	if v57&int32(3) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v151 = v58
	v152 = v36
	v153 = v57
	goto L26
L34:
	;
	goto L35
L35:
	;
	v79 = v58
	v80 = v36
	v81 = v57
	goto L36
L36:
	;
	if v80 == int32(0) {
		goto L20
	} else {
		goto L38
	}
L37:
	;
	v151 = v88
	v152 = v90
	v153 = v92
	goto L26
L38:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v85)
	v87 = int32(1)
	v88 = v79 + v87
	v90 = v80 - v87
	v92 = v81 + v87
	if v92&int32(3) != 0 {
		v79 = v88
		v80 = v90
		v81 = v92
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	if v134 == int32(0) {
		goto L20
	} else {
		goto L53
	}
L41:
	;
	if v62&int32(3) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v99 = v36
	goto L45
L43:
	;
	v114 = v36
	goto L44
L44:
	;
	if base.Ui32(v114) <= base.Ui32(int32(3)) {
		v134 = v114
		goto L40
	} else {
		goto L49
	}
L45:
	;
	if v99 == int32(0) {
		goto L20
	} else {
		goto L47
	}
L46:
	;
	v114 = v105
	goto L44
L47:
	;
	v105 = v99 - int32(1)
	v106 = v57 + v105
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v105))))
	*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v108)
	if v106&int32(3) != 0 {
		v99 = v105
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v121 = v114
	goto L50
L50:
	;
	v125 = v121 - int32(4)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v58+v125)))
	*(*int32)(unsafe.Add(mBase, uint32(v57+v125))) = v128
	if base.Ui32(int32(3)) < base.Ui32(v125) {
		v121 = v125
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v134 = v125
	goto L40
L52:
	;
	goto L51
L53:
	;
	v141 = v134
	goto L54
L54:
	;
	v145 = v141 - int32(1)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v145))))
	*(*uint8)(unsafe.Add(mBase, uint32(v57+v145))) = uint8(v148)
	if v145 != 0 {
		v141 = v145
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L20
L56:
	;
	goto L55
L57:
	;
	v158 = v151
	v159 = v152
	v160 = v153
	goto L58
L58:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v162
	v164 = int32(4)
	v165 = v158 + v164
	v167 = v160 + v164
	v169 = v159 - v164
	if base.Ui32(int32(3)) < base.Ui32(v169) {
		v158 = v165
		v159 = v169
		v160 = v167
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v173 = v165
	v174 = v169
	v175 = v167
	goto L25
L60:
	;
	goto L59
L61:
	;
	v180 = v173
	v181 = v174
	v182 = v175
	goto L62
L62:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	*(*uint8)(unsafe.Add(mBase, uint32(v182))) = uint8(v184)
	v186 = int32(1)
	v191 = v181 - v186
	if v191 != 0 {
		v180 = v180 + v186
		v181 = v191
		v182 = v182 + v186
		goto L62
	} else {
		goto L64
	}
L63:
	;
	goto L20
L64:
	;
	goto L63
L65:
	;
	return int32(0)
}
func F_slotsync_worker_onexit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v4 = *(*int32)(unsafe.Add(mBase, _consts[841]))
	if v4 != 0 {
		F_ReplicationSlotRelease(m)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_ReplicationSlotCleanup(m, int32(0))
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, _consts[839]))
				v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(1)
				if v12 != 0 {
					v16 = *(*int32)(unsafe.Add(mBase, _consts[839]))
					F_s_lock(m, v16+int32(16), int32(523799), int32(1231), int32(105931))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, _consts[839]))
						*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(-1)
						v29 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
						if v29 != 0 {
							v30 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v30)
							*(*uint8)(unsafe.Add(mBase, _consts[840])) = uint8(v30)
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = int32(0)
						return
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _consts[839]))
					*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(-1)
					v29 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
					if v29 != 0 {
						v30 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v30)
						*(*uint8)(unsafe.Add(mBase, _consts[840])) = uint8(v30)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = int32(0)
					return
				}
			}
		}
	} else {
		F_ReplicationSlotCleanup(m, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _consts[839]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(1)
			if v12 != 0 {
				v16 = *(*int32)(unsafe.Add(mBase, _consts[839]))
				F_s_lock(m, v16+int32(16), int32(523799), int32(1231), int32(105931))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _consts[839]))
					*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(-1)
					v29 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
					if v29 != 0 {
						v30 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v30)
						*(*uint8)(unsafe.Add(mBase, _consts[840])) = uint8(v30)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = int32(0)
					return
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _consts[839]))
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(-1)
				v29 = int32(*(*uint8)(unsafe.Add(mBase, _consts[840])))
				if v29 != 0 {
					v30 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v25)+5)) = uint8(v30)
					*(*uint8)(unsafe.Add(mBase, _consts[840])) = uint8(v30)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = int32(0)
				return
			}
		}
	}
}
func F_smgrdestroyall(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	v6 = int32(4548892)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v8 + int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1200]))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L15
	}
L2:
	;
	v98 = int32(4548892)
	v100 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v100 - int32(1)
	return
L3:
	;
	if v13 == int32(4477720) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = v13
	goto L5
L5:
	;
	v24 = v18 + int32(4)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = int32(4548892)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v28 + int32(1)
	v33 = v18 - int32(76)
	v36 = v18 - int32(40)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37*int32(80))+uint32(_consts[1201])))
	m.T0[v42].(func(*base.Module, int32, int32))(m, v33, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L2
L7:
	;
	return
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46*int32(80))+uint32(_consts[1201])))
	m.T0[v51].(func(*base.Module, int32, int32))(m, v33, int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55*int32(80))+uint32(_consts[1201])))
	m.T0[v60].(func(*base.Module, int32, int32))(m, v33, int32(2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64*int32(80))+uint32(_consts[1201])))
	m.T0[v69].(func(*base.Module, int32, int32))(m, v33, int32(3))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v75
	v78 = *(*int32)(unsafe.Add(mBase, _consts[1202]))
	v81 = F_hash_search(m, v78, v33, int32(2), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v81 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v85 = int32(4548892)
	v87 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v87 - int32(1)
	if v25 != int32(4477720) {
		v18 = v25
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L6
L15:
	;
	F_errmsg_internal(m, int32(465202), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(518193), int32(339), int32(18621))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_smgrnblocks_cached(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[158])))
	if v4 == int32(1) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+20))
		if v10 != int32(-1) {
			v15 = v10
		} else {
			v15 = int32(-1)
		}
	} else {
		v15 = int32(-1)
	}
	return v15
}
func F_smgrprefetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v5 = int32(4548892)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[171]))
	*(*int32)(unsafe.Add(mBase, _consts[171])) = v7 + int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11*int32(80))+uint32(_consts[1203])))
	v17 = m.T0[v16].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = int32(4548892)
		v23 = *(*int32)(unsafe.Add(mBase, _consts[171]))
		*(*int32)(unsafe.Add(mBase, _consts[171])) = v23 - int32(1)
		return v17
	}
}
func F_sort(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(1)
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v18 != int32(2) {
			v67 = int32(0)
			v68 = v17
			v69 = v2
			v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			if v71 != 0 {
				v72 = F_array_contains_nulls(m, v13)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					if v72 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(161982), int32(0))
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(518874), int32(206), int32(84776))
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						v76 = v13 + int32(16)
						v77 = F_ArrayGetNItems(m, v74, v76)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							if int32(2) <= v77 {
								v81 = int32(1)
								if v68 != 0 {
									v120 = v81
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v122 = F_ArrayGetNItems(m, v121, v76)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v120)
										v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
										if v125 != 0 {
											v133 = v125
										} else {
											v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v133 = (v126<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										F_isort(m, v133+v13, v122, v10+int32(15))
										mBase = m.M
										m.G0 = v10 + int32(16)
										return v13
									}
								} else {
									switch v69 - int32(3) {
									case 0:
										v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
										if v84|int32(32) != int32(97) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(763223), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(518874), int32(224), int32(84776))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
											if v89|int32(32) != int32(115) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(763223), int32(0))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(518874), int32(224), int32(84776))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+2)))
												if v94|int32(32) == int32(99) {
													v120 = v81
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v122 = F_ArrayGetNItems(m, v121, v76)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v120)
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
														if v125 != 0 {
															v133 = v125
														} else {
															v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															v133 = (v126<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														}
														F_isort(m, v133+v13, v122, v10+int32(15))
														mBase = m.M
														m.G0 = v10 + int32(16)
														return v13
													}
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(763223), int32(0))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(518874), int32(224), int32(84776))
																mBase = m.M
																v185 = m.ExcPending
																if v185 != 0 {
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
										}
									case 1:
										v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
										if v99|int32(32) != int32(100) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(763223), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(518874), int32(224), int32(84776))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
											if v104|int32(32) != int32(101) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(763223), int32(0))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(518874), int32(224), int32(84776))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+2)))
												if v109|int32(32) != int32(115) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(763223), int32(0))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(518874), int32(224), int32(84776))
																mBase = m.M
																v185 = m.ExcPending
																if v185 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+3)))
													if v115|int32(32) != int32(99) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(763223), int32(0))
																mBase = m.M
																v178 = m.ExcPending
																if v178 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(518874), int32(224), int32(84776))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														v120 = int32(0)
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														v122 = F_ArrayGetNItems(m, v121, v76)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return int32(0)
														} else {
															*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v120)
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
															if v125 != 0 {
																v133 = v125
															} else {
																v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
																v133 = (v126<<(uint(int32(3))%32) + int32(23)) & int32(-8)
															}
															F_isort(m, v133+v13, v122, v10+int32(15))
															mBase = m.M
															m.G0 = v10 + int32(16)
															return v13
														}
													}
												}
											}
										}
									default:
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(763223), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(518874), int32(224), int32(84776))
													mBase = m.M
													v185 = m.ExcPending
													if v185 != 0 {
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
							} else {
								m.G0 = v10 + int32(16)
								return v13
							}
						}
					}
				}
			} else {
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				v76 = v13 + int32(16)
				v77 = F_ArrayGetNItems(m, v74, v76)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					if int32(2) <= v77 {
						v81 = int32(1)
						if v68 != 0 {
							v120 = v81
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v122 = F_ArrayGetNItems(m, v121, v76)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v120)
								v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
								if v125 != 0 {
									v133 = v125
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
									v133 = (v126<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								F_isort(m, v133+v13, v122, v10+int32(15))
								mBase = m.M
								m.G0 = v10 + int32(16)
								return v13
							}
						} else {
							switch v69 - int32(3) {
							case 0:
								v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
								if v84|int32(32) != int32(97) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(763223), int32(0))
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(518874), int32(224), int32(84776))
												mBase = m.M
												v185 = m.ExcPending
												if v185 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
									if v89|int32(32) != int32(115) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(763223), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(518874), int32(224), int32(84776))
													mBase = m.M
													v185 = m.ExcPending
													if v185 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+2)))
										if v94|int32(32) == int32(99) {
											v120 = v81
											v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
											v122 = F_ArrayGetNItems(m, v121, v76)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v120)
												v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
												if v125 != 0 {
													v133 = v125
												} else {
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v133 = (v126<<(uint(int32(3))%32) + int32(23)) & int32(-8)
												}
												F_isort(m, v133+v13, v122, v10+int32(15))
												mBase = m.M
												m.G0 = v10 + int32(16)
												return v13
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(763223), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(518874), int32(224), int32(84776))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
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
								}
							case 1:
								v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
								if v99|int32(32) != int32(100) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(763223), int32(0))
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(518874), int32(224), int32(84776))
												mBase = m.M
												v185 = m.ExcPending
												if v185 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
									if v104|int32(32) != int32(101) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(763223), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(518874), int32(224), int32(84776))
													mBase = m.M
													v185 = m.ExcPending
													if v185 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+2)))
										if v109|int32(32) != int32(115) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(763223), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(518874), int32(224), int32(84776))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+3)))
											if v115|int32(32) != int32(99) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(763223), int32(0))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(518874), int32(224), int32(84776))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v120 = int32(0)
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												v122 = F_ArrayGetNItems(m, v121, v76)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v120)
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
													if v125 != 0 {
														v133 = v125
													} else {
														v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														v133 = (v126<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													}
													F_isort(m, v133+v13, v122, v10+int32(15))
													mBase = m.M
													m.G0 = v10 + int32(16)
													return v13
												}
											}
										}
									}
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v169 = m.ExcPending
								if v169 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(763223), int32(0))
										mBase = m.M
										v178 = m.ExcPending
										if v178 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(518874), int32(224), int32(84776))
											mBase = m.M
											v185 = m.ExcPending
											if v185 != 0 {
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
					} else {
						m.G0 = v10 + int32(16)
						return v13
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v23 = F_pg_detoast_datum_packed(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v67 = int32(0)
					v68 = v17
					v69 = v2
				} else {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
					if v28 == int32(1) {
						v31 = int32(4)
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
						if v33&int32(254) == int32(2) {
							v42 = v31
						} else {
							v42 = base.B2i32(v33 == int32(18)) << (uint(v31) % 32)
						}
						if v33 == int32(1) {
							v45 = v31
						} else {
							v45 = v42
						}
						v57 = v17
						v59 = v45
					} else {
						if v28&int32(1) != 0 {
							v48 = int32(1)
							v57 = v28
							v59 = int32(base.Ui32(v28)>>(uint(v48)%32)) - v48
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
							v57 = v52
							v59 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v60 = int32(1)
					if v57&v60 != 0 {
						v64 = v60
					} else {
						v64 = int32(4)
					}
					v67 = v23 + v64
					v68 = int32(0)
					v69 = v59
				}
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				if v71 != 0 {
					v72 = F_array_contains_nulls(m, v13)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						if v72 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67108994))
								mBase = m.M
								v151 = m.ExcPending
								if v151 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(161982), int32(0))
									mBase = m.M
									v157 = m.ExcPending
									if v157 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(518874), int32(206), int32(84776))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v76 = v13 + int32(16)
							v77 = F_ArrayGetNItems(m, v74, v76)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								if int32(2) <= v77 {
									v81 = int32(1)
									if v68 != 0 {
										v120 = v81
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v122 = F_ArrayGetNItems(m, v121, v76)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v120)
											v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
											if v125 != 0 {
												v133 = v125
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												v133 = (v126<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											F_isort(m, v133+v13, v122, v10+int32(15))
											mBase = m.M
											m.G0 = v10 + int32(16)
											return v13
										}
									} else {
										switch v69 - int32(3) {
										case 0:
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
											if v84|int32(32) != int32(97) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(763223), int32(0))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(518874), int32(224), int32(84776))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
												if v89|int32(32) != int32(115) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(763223), int32(0))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(518874), int32(224), int32(84776))
																mBase = m.M
																v185 = m.ExcPending
																if v185 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+2)))
													if v94|int32(32) == int32(99) {
														v120 = v81
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														v122 = F_ArrayGetNItems(m, v121, v76)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return int32(0)
														} else {
															*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v120)
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
															if v125 != 0 {
																v133 = v125
															} else {
																v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
																v133 = (v126<<(uint(int32(3))%32) + int32(23)) & int32(-8)
															}
															F_isort(m, v133+v13, v122, v10+int32(15))
															mBase = m.M
															m.G0 = v10 + int32(16)
															return v13
														}
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(763223), int32(0))
																mBase = m.M
																v178 = m.ExcPending
																if v178 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(518874), int32(224), int32(84776))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
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
											}
										case 1:
											v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
											if v99|int32(32) != int32(100) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(763223), int32(0))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(518874), int32(224), int32(84776))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
												if v104|int32(32) != int32(101) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(763223), int32(0))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(518874), int32(224), int32(84776))
																mBase = m.M
																v185 = m.ExcPending
																if v185 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+2)))
													if v109|int32(32) != int32(115) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(763223), int32(0))
																mBase = m.M
																v178 = m.ExcPending
																if v178 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(518874), int32(224), int32(84776))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+3)))
														if v115|int32(32) != int32(99) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50856066))
																mBase = m.M
																v172 = m.ExcPending
																if v172 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(763223), int32(0))
																	mBase = m.M
																	v178 = m.ExcPending
																	if v178 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(518874), int32(224), int32(84776))
																		mBase = m.M
																		v185 = m.ExcPending
																		if v185 != 0 {
																			return int32(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															}
														} else {
															v120 = int32(0)
															v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															v122 = F_ArrayGetNItems(m, v121, v76)
															mBase = m.M
															v123 = m.ExcPending
															if v123 != 0 {
																return int32(0)
															} else {
																*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v120)
																v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
																if v125 != 0 {
																	v133 = v125
																} else {
																	v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
																	v133 = (v126<<(uint(int32(3))%32) + int32(23)) & int32(-8)
																}
																F_isort(m, v133+v13, v122, v10+int32(15))
																mBase = m.M
																m.G0 = v10 + int32(16)
																return v13
															}
														}
													}
												}
											}
										default:
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(763223), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(518874), int32(224), int32(84776))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
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
								} else {
									m.G0 = v10 + int32(16)
									return v13
								}
							}
						}
					}
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v76 = v13 + int32(16)
					v77 = F_ArrayGetNItems(m, v74, v76)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						if int32(2) <= v77 {
							v81 = int32(1)
							if v68 != 0 {
								v120 = v81
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v122 = F_ArrayGetNItems(m, v121, v76)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v120)
									v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
									if v125 != 0 {
										v133 = v125
									} else {
										v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v133 = (v126<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									F_isort(m, v133+v13, v122, v10+int32(15))
									mBase = m.M
									m.G0 = v10 + int32(16)
									return v13
								}
							} else {
								switch v69 - int32(3) {
								case 0:
									v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
									if v84|int32(32) != int32(97) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(763223), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(518874), int32(224), int32(84776))
													mBase = m.M
													v185 = m.ExcPending
													if v185 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
										if v89|int32(32) != int32(115) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(763223), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(518874), int32(224), int32(84776))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+2)))
											if v94|int32(32) == int32(99) {
												v120 = v81
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
												v122 = F_ArrayGetNItems(m, v121, v76)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v120)
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
													if v125 != 0 {
														v133 = v125
													} else {
														v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
														v133 = (v126<<(uint(int32(3))%32) + int32(23)) & int32(-8)
													}
													F_isort(m, v133+v13, v122, v10+int32(15))
													mBase = m.M
													m.G0 = v10 + int32(16)
													return v13
												}
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(763223), int32(0))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(518874), int32(224), int32(84776))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
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
									}
								case 1:
									v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
									if v99|int32(32) != int32(100) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v169 = m.ExcPending
										if v169 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(763223), int32(0))
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(518874), int32(224), int32(84776))
													mBase = m.M
													v185 = m.ExcPending
													if v185 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
										if v104|int32(32) != int32(101) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(763223), int32(0))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(518874), int32(224), int32(84776))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+2)))
											if v109|int32(32) != int32(115) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v172 = m.ExcPending
													if v172 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(763223), int32(0))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(518874), int32(224), int32(84776))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+3)))
												if v115|int32(32) != int32(99) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(763223), int32(0))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(518874), int32(224), int32(84776))
																mBase = m.M
																v185 = m.ExcPending
																if v185 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v120 = int32(0)
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
													v122 = F_ArrayGetNItems(m, v121, v76)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v120)
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
														if v125 != 0 {
															v133 = v125
														} else {
															v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
															v133 = (v126<<(uint(int32(3))%32) + int32(23)) & int32(-8)
														}
														F_isort(m, v133+v13, v122, v10+int32(15))
														mBase = m.M
														m.G0 = v10 + int32(16)
														return v13
													}
												}
											}
										}
									}
								default:
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(763223), int32(0))
											mBase = m.M
											v178 = m.ExcPending
											if v178 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(518874), int32(224), int32(84776))
												mBase = m.M
												v185 = m.ExcPending
												if v185 != 0 {
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
						} else {
							m.G0 = v10 + int32(16)
							return v13
						}
					}
				}
			}
		}
	}
}
func F_sort_pending_writebacks(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v63 int32
	_ = v63
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v147 int64
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int64
	_ = v727
	var v729 int64
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int64
	_ = v737
	var v739 int64
	_ = v739
	var v741 int32
	_ = v741
	var v743 int64
	_ = v743
	var v745 int64
	_ = v745
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int64
	_ = v812
	var v814 int64
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int64
	_ = v822
	var v824 int64
	_ = v824
	var v826 int32
	_ = v826
	var v828 int64
	_ = v828
	var v830 int64
	_ = v830
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v862 int32
	_ = v862
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int64
	_ = v904
	var v906 int64
	_ = v906
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int64
	_ = v914
	var v916 int64
	_ = v916
	var v918 int32
	_ = v918
	var v920 int64
	_ = v920
	var v922 int64
	_ = v922
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v935 int32
	_ = v935
	var v946 int32
	_ = v946
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v972 int32
	_ = v972
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int64
	_ = v990
	var v992 int64
	_ = v992
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1001 int64
	_ = v1001
	var v1003 int64
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int64
	_ = v1007
	var v1009 int64
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1051 int32
	_ = v1051
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int64
	_ = v1073
	var v1075 int64
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1084 int64
	_ = v1084
	var v1086 int64
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1090 int64
	_ = v1090
	var v1092 int64
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int64
	_ = v1140
	var v1142 int64
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int64
	_ = v1150
	var v1152 int64
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int64
	_ = v1156
	var v1158 int64
	_ = v1158
	var v1160 int32
	_ = v1160
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = l0
	v24 = l1
	goto L1
L1:
	;
	v42 = v23 + int32(20)
	v44 = v24
	goto L3
L2:
	;
	m.G0 = v21 + int32(32)
	return
L3:
	;
	v63 = v23 + v44*int32(20)
	if base.Ui32(v44) <= base.Ui32(int32(6)) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	goto L4
L6:
	;
	if base.Ui32(v63) <= base.Ui32(v42) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if base.Ui32(v63) <= base.Ui32(v42) {
		goto L5
	} else {
		goto L28
	}
L9:
	;
	v81 = v42
	goto L10
L10:
	;
	if base.Ui32(v81) <= base.Ui32(v23) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L5
L12:
	;
	v185 = v81 + int32(20)
	if base.Ui32(v185) < base.Ui32(v63) {
		v81 = v185
		goto L10
	} else {
		goto L27
	}
L13:
	;
	v89 = v81
	goto L14
L14:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v89-int32(12))))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if base.Ui32(v106) < base.Ui32(v107) {
		goto L12
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	v110 = v89 - int32(20)
	if base.Ui32(v107) < base.Ui32(v106) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v136 = v21 + int32(24)
	v137 = int32(16)
	v138 = v89 + v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v139
	v142 = v21 + v137
	v143 = int32(8)
	v144 = v89 + v143
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v144)))
	*(*int64)(unsafe.Add(mBase, uint32(v142))) = v145
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v147
	v150 = v110 + v137
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v151
	v154 = v110 + v143
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v154)))
	*(*int64)(unsafe.Add(mBase, uint32(v144))) = v155
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v110)))
	*(*int64)(unsafe.Add(mBase, uint32(v89))) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v159
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v142)))
	*(*int64)(unsafe.Add(mBase, uint32(v154))) = v161
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v110))) = v163
	if base.Ui32(v23) < base.Ui32(v110) {
		v89 = v110
		goto L14
	} else {
		goto L26
	}
L18:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v89-int32(16))))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if base.Ui32(v114) < base.Ui32(v115) {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v115) < base.Ui32(v114) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if base.Ui32(v118) < base.Ui32(v119) {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	if base.Ui32(v119) < base.Ui32(v118) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v89-int32(8))))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v124 < v125 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	if v125 < v124 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v89-int32(4))))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if base.Ui32(v130) <= base.Ui32(v131) {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	goto L17
L26:
	;
	goto L15
L27:
	;
	goto L11
L28:
	;
	v190 = v42
	goto L29
L29:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v190-int32(12))))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
	if base.Ui32(v208) < base.Ui32(v209) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v246 = v23 + int32(base.Ui32(v44)>>(uint(int32(1))%32))*int32(20)
	if v44 != int32(7) {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	goto L30
L32:
	;
	v238 = v190 + int32(20)
	if base.Ui32(v238) < base.Ui32(v63) {
		v190 = v238
		goto L29
	} else {
		goto L42
	}
L33:
	;
	if base.Ui32(v209) < base.Ui32(v208) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v190-int32(16))))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if base.Ui32(v214) < base.Ui32(v215) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	if base.Ui32(v215) < base.Ui32(v214) {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v190-int32(20))))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	if base.Ui32(v220) < base.Ui32(v221) {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	if base.Ui32(v221) < base.Ui32(v220) {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v190-int32(8))))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	if v226 < v227 {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	if v227 < v226 {
		goto L31
	} else {
		goto L40
	}
L40:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v190-int32(4))))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v190)+16))
	if base.Ui32(v233) < base.Ui32(v232) {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	goto L32
L42:
	;
	goto L5
L43:
	;
	v250 = v63 - int32(20)
	if base.Ui32(v44) < base.Ui32(int32(41)) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v711 = v246
	goto L45
L45:
	;
	v718 = v21 + int32(24)
	v719 = int32(16)
	v720 = v23 + v719
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)))
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = v721
	v724 = v21 + v719
	v725 = int32(8)
	v726 = v23 + v725
	v727 = *(*int64)(unsafe.Add(mBase, uint32(v726)))
	*(*int64)(unsafe.Add(mBase, uint32(v724))) = v727
	v729 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v729
	v732 = v711 + v719
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	*(*int32)(unsafe.Add(mBase, uint32(v720))) = v733
	v736 = v711 + v725
	v737 = *(*int64)(unsafe.Add(mBase, uint32(v736)))
	*(*int64)(unsafe.Add(mBase, uint32(v726))) = v737
	v739 = *(*int64)(unsafe.Add(mBase, uint32(v711)))
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v739
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v718)))
	*(*int32)(unsafe.Add(mBase, uint32(v732))) = v741
	v743 = *(*int64)(unsafe.Add(mBase, uint32(v724)))
	*(*int64)(unsafe.Add(mBase, uint32(v736))) = v743
	v745 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v711))) = v745
	v748 = v63 - int32(20)
	v751 = v748
	v752 = v42
	v758 = v42
	v762 = v748
	goto L398
L46:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v596)+4))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v596)))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v600)+4))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v600)))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v600)+8))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v596)+8))
	if base.Ui32(v615) < base.Ui32(v616) {
		goto L314
	} else {
		goto L315
	}
L47:
	;
	v596 = v246
	v597 = v250
	v600 = v23
	goto L46
L48:
	;
	goto L49
L49:
	;
	v254 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
	v256 = v254 * int32(20)
	v257 = v23 + v256
	v260 = v23 + v254*int32(40)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v257)+8))
	if base.Ui32(v274) < base.Ui32(v275) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v371 = v254 * int32(-20)
	v372 = v246 + v371
	v373 = v246 + v256
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v246)+8))
	if base.Ui32(v387) < base.Ui32(v388) {
		goto L140
	} else {
		goto L141
	}
L51:
	;
	v369 = v357
	goto L50
L52:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	if base.Ui32(v275) < base.Ui32(v327) {
		goto L100
	} else {
		goto L101
	}
L53:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	if base.Ui32(v275) < base.Ui32(v291) {
		v357 = v257
		goto L51
	} else {
		goto L63
	}
L54:
	;
	if base.Ui32(v275) < base.Ui32(v274) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	if base.Ui32(v272) < base.Ui32(v270) {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	if base.Ui32(v270) < base.Ui32(v272) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	if base.Ui32(v273) < base.Ui32(v271) {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	if base.Ui32(v271) < base.Ui32(v273) {
		goto L52
	} else {
		goto L59
	}
L59:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	if v282 < v283 {
		goto L53
	} else {
		goto L60
	}
L60:
	;
	if v283 < v282 {
		goto L52
	} else {
		goto L61
	}
L61:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v257)+16))
	if base.Ui32(v287) <= base.Ui32(v286) {
		goto L52
	} else {
		goto L62
	}
L62:
	;
	goto L53
L63:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	if base.Ui32(v291) < base.Ui32(v275) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if base.Ui32(v274) < base.Ui32(v291) {
		goto L73
	} else {
		goto L74
	}
L65:
	;
	if base.Ui32(v270) < base.Ui32(v293) {
		v357 = v257
		goto L51
	} else {
		goto L66
	}
L66:
	;
	if base.Ui32(v293) < base.Ui32(v270) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	if base.Ui32(v271) < base.Ui32(v294) {
		v357 = v257
		goto L51
	} else {
		goto L68
	}
L68:
	;
	if base.Ui32(v294) < base.Ui32(v271) {
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	if v300 < v301 {
		v357 = v257
		goto L51
	} else {
		goto L70
	}
L70:
	;
	if v301 < v300 {
		goto L64
	} else {
		goto L71
	}
L71:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v257)+16))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	if base.Ui32(v304) < base.Ui32(v305) {
		v357 = v257
		goto L51
	} else {
		goto L72
	}
L72:
	;
	goto L64
L73:
	;
	v369 = v260
	goto L50
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(v291) < base.Ui32(v274) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v369 = v23
	goto L50
L77:
	;
	goto L78
L78:
	;
	if base.Ui32(v272) < base.Ui32(v293) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v369 = v260
	goto L50
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(v293) < base.Ui32(v272) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v369 = v23
	goto L50
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(v273) < base.Ui32(v294) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v369 = v260
	goto L50
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(v294) < base.Ui32(v273) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v369 = v23
	goto L50
L89:
	;
	goto L90
L90:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	if v315 < v316 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v369 = v260
	goto L50
L92:
	;
	goto L93
L93:
	;
	if v316 < v315 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v369 = v23
	goto L50
L95:
	;
	goto L96
L96:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	if base.Ui32(v319) < base.Ui32(v320) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v322 = v260
	goto L99
L98:
	;
	v322 = v23
	goto L99
L99:
	;
	v369 = v322
	goto L50
L100:
	;
	if base.Ui32(v274) < base.Ui32(v327) {
		goto L110
	} else {
		goto L111
	}
L101:
	;
	if base.Ui32(v327) < base.Ui32(v275) {
		v357 = v257
		goto L51
	} else {
		goto L102
	}
L102:
	;
	if base.Ui32(v270) < base.Ui32(v325) {
		goto L100
	} else {
		goto L103
	}
L103:
	;
	if base.Ui32(v325) < base.Ui32(v270) {
		v357 = v257
		goto L51
	} else {
		goto L104
	}
L104:
	;
	if base.Ui32(v271) < base.Ui32(v326) {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	if base.Ui32(v326) < base.Ui32(v271) {
		v357 = v257
		goto L51
	} else {
		goto L106
	}
L106:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	if v334 < v335 {
		goto L100
	} else {
		goto L107
	}
L107:
	;
	if v335 < v334 {
		v357 = v257
		goto L51
	} else {
		goto L108
	}
L108:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v257)+16))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	if base.Ui32(v339) < base.Ui32(v338) {
		v357 = v257
		goto L51
	} else {
		goto L109
	}
L109:
	;
	goto L100
L110:
	;
	v369 = v23
	goto L50
L111:
	;
	goto L112
L112:
	;
	if base.Ui32(v327) < base.Ui32(v274) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v369 = v260
	goto L50
L114:
	;
	goto L115
L115:
	;
	if base.Ui32(v272) < base.Ui32(v325) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v369 = v23
	goto L50
L117:
	;
	goto L118
L118:
	;
	if base.Ui32(v325) < base.Ui32(v272) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v369 = v260
	goto L50
L120:
	;
	goto L121
L121:
	;
	if base.Ui32(v273) < base.Ui32(v326) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v369 = v23
	goto L50
L123:
	;
	goto L124
L124:
	;
	if base.Ui32(v326) < base.Ui32(v273) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v369 = v260
	goto L50
L126:
	;
	goto L127
L127:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	if v349 < v350 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v369 = v23
	goto L50
L129:
	;
	goto L130
L130:
	;
	if v350 < v349 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v369 = v260
	goto L50
L132:
	;
	goto L133
L133:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v260)+16))
	if base.Ui32(v353) < base.Ui32(v354) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v356 = v23
	goto L136
L135:
	;
	v356 = v260
	goto L136
L136:
	;
	v357 = v356
	goto L51
L137:
	;
	v485 = v250 + v254*int32(-40)
	v486 = v250 + v371
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v486)+4))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v485)+4))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v485)+8))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v486)+8))
	if base.Ui32(v500) < base.Ui32(v501) {
		goto L227
	} else {
		goto L228
	}
L138:
	;
	v482 = v470
	goto L137
L139:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	if base.Ui32(v388) < base.Ui32(v440) {
		goto L187
	} else {
		goto L188
	}
L140:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	if base.Ui32(v388) < base.Ui32(v404) {
		v470 = v246
		goto L138
	} else {
		goto L150
	}
L141:
	;
	if base.Ui32(v388) < base.Ui32(v387) {
		goto L139
	} else {
		goto L142
	}
L142:
	;
	if base.Ui32(v385) < base.Ui32(v383) {
		goto L140
	} else {
		goto L143
	}
L143:
	;
	if base.Ui32(v383) < base.Ui32(v385) {
		goto L139
	} else {
		goto L144
	}
L144:
	;
	if base.Ui32(v386) < base.Ui32(v384) {
		goto L140
	} else {
		goto L145
	}
L145:
	;
	if base.Ui32(v384) < base.Ui32(v386) {
		goto L139
	} else {
		goto L146
	}
L146:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	if v395 < v396 {
		goto L140
	} else {
		goto L147
	}
L147:
	;
	if v396 < v395 {
		goto L139
	} else {
		goto L148
	}
L148:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v246)+16))
	if base.Ui32(v400) <= base.Ui32(v399) {
		goto L139
	} else {
		goto L149
	}
L149:
	;
	goto L140
L150:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	if base.Ui32(v404) < base.Ui32(v388) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	if base.Ui32(v387) < base.Ui32(v404) {
		goto L160
	} else {
		goto L161
	}
L152:
	;
	if base.Ui32(v383) < base.Ui32(v406) {
		v470 = v246
		goto L138
	} else {
		goto L153
	}
L153:
	;
	if base.Ui32(v406) < base.Ui32(v383) {
		goto L151
	} else {
		goto L154
	}
L154:
	;
	if base.Ui32(v384) < base.Ui32(v407) {
		v470 = v246
		goto L138
	} else {
		goto L155
	}
L155:
	;
	if base.Ui32(v407) < base.Ui32(v384) {
		goto L151
	} else {
		goto L156
	}
L156:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	if v413 < v414 {
		v470 = v246
		goto L138
	} else {
		goto L157
	}
L157:
	;
	if v414 < v413 {
		goto L151
	} else {
		goto L158
	}
L158:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v246)+16))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	if base.Ui32(v417) < base.Ui32(v418) {
		v470 = v246
		goto L138
	} else {
		goto L159
	}
L159:
	;
	goto L151
L160:
	;
	v482 = v373
	goto L137
L161:
	;
	goto L162
L162:
	;
	if base.Ui32(v404) < base.Ui32(v387) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v482 = v372
	goto L137
L164:
	;
	goto L165
L165:
	;
	if base.Ui32(v385) < base.Ui32(v406) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v482 = v373
	goto L137
L167:
	;
	goto L168
L168:
	;
	if base.Ui32(v406) < base.Ui32(v385) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v482 = v372
	goto L137
L170:
	;
	goto L171
L171:
	;
	if base.Ui32(v386) < base.Ui32(v407) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v482 = v373
	goto L137
L173:
	;
	goto L174
L174:
	;
	if base.Ui32(v407) < base.Ui32(v386) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v482 = v372
	goto L137
L176:
	;
	goto L177
L177:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	if v428 < v429 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v482 = v373
	goto L137
L179:
	;
	goto L180
L180:
	;
	if v429 < v428 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v482 = v372
	goto L137
L182:
	;
	goto L183
L183:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	if base.Ui32(v432) < base.Ui32(v433) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v435 = v373
	goto L186
L185:
	;
	v435 = v372
	goto L186
L186:
	;
	v482 = v435
	goto L137
L187:
	;
	if base.Ui32(v387) < base.Ui32(v440) {
		goto L197
	} else {
		goto L198
	}
L188:
	;
	if base.Ui32(v440) < base.Ui32(v388) {
		v470 = v246
		goto L138
	} else {
		goto L189
	}
L189:
	;
	if base.Ui32(v383) < base.Ui32(v438) {
		goto L187
	} else {
		goto L190
	}
L190:
	;
	if base.Ui32(v438) < base.Ui32(v383) {
		v470 = v246
		goto L138
	} else {
		goto L191
	}
L191:
	;
	if base.Ui32(v384) < base.Ui32(v439) {
		goto L187
	} else {
		goto L192
	}
L192:
	;
	if base.Ui32(v439) < base.Ui32(v384) {
		v470 = v246
		goto L138
	} else {
		goto L193
	}
L193:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	if v447 < v448 {
		goto L187
	} else {
		goto L194
	}
L194:
	;
	if v448 < v447 {
		v470 = v246
		goto L138
	} else {
		goto L195
	}
L195:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v246)+16))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	if base.Ui32(v452) < base.Ui32(v451) {
		v470 = v246
		goto L138
	} else {
		goto L196
	}
L196:
	;
	goto L187
L197:
	;
	v482 = v372
	goto L137
L198:
	;
	goto L199
L199:
	;
	if base.Ui32(v440) < base.Ui32(v387) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v482 = v373
	goto L137
L201:
	;
	goto L202
L202:
	;
	if base.Ui32(v385) < base.Ui32(v438) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v482 = v372
	goto L137
L204:
	;
	goto L205
L205:
	;
	if base.Ui32(v438) < base.Ui32(v385) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v482 = v373
	goto L137
L207:
	;
	goto L208
L208:
	;
	if base.Ui32(v386) < base.Ui32(v439) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v482 = v372
	goto L137
L210:
	;
	goto L211
L211:
	;
	if base.Ui32(v439) < base.Ui32(v386) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v482 = v373
	goto L137
L213:
	;
	goto L214
L214:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	if v462 < v463 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v482 = v372
	goto L137
L216:
	;
	goto L217
L217:
	;
	if v463 < v462 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v482 = v373
	goto L137
L219:
	;
	goto L220
L220:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	if base.Ui32(v466) < base.Ui32(v467) {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v469 = v372
	goto L223
L222:
	;
	v469 = v373
	goto L223
L223:
	;
	v470 = v469
	goto L138
L224:
	;
	v596 = v482
	v597 = v595
	v600 = v369
	goto L46
L225:
	;
	v595 = v583
	goto L224
L226:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v250)+8))
	if base.Ui32(v501) < base.Ui32(v553) {
		goto L274
	} else {
		goto L275
	}
L227:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v250)+8))
	if base.Ui32(v501) < base.Ui32(v517) {
		v583 = v486
		goto L225
	} else {
		goto L237
	}
L228:
	;
	if base.Ui32(v501) < base.Ui32(v500) {
		goto L226
	} else {
		goto L229
	}
L229:
	;
	if base.Ui32(v498) < base.Ui32(v496) {
		goto L227
	} else {
		goto L230
	}
L230:
	;
	if base.Ui32(v496) < base.Ui32(v498) {
		goto L226
	} else {
		goto L231
	}
L231:
	;
	if base.Ui32(v499) < base.Ui32(v497) {
		goto L227
	} else {
		goto L232
	}
L232:
	;
	if base.Ui32(v497) < base.Ui32(v499) {
		goto L226
	} else {
		goto L233
	}
L233:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v485)+12))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v486)+12))
	if v508 < v509 {
		goto L227
	} else {
		goto L234
	}
L234:
	;
	if v509 < v508 {
		goto L226
	} else {
		goto L235
	}
L235:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v485)+16))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v486)+16))
	if base.Ui32(v513) <= base.Ui32(v512) {
		goto L226
	} else {
		goto L236
	}
L236:
	;
	goto L227
L237:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if base.Ui32(v517) < base.Ui32(v501) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	if base.Ui32(v500) < base.Ui32(v517) {
		goto L247
	} else {
		goto L248
	}
L239:
	;
	if base.Ui32(v496) < base.Ui32(v519) {
		v583 = v486
		goto L225
	} else {
		goto L240
	}
L240:
	;
	if base.Ui32(v519) < base.Ui32(v496) {
		goto L238
	} else {
		goto L241
	}
L241:
	;
	if base.Ui32(v497) < base.Ui32(v520) {
		v583 = v486
		goto L225
	} else {
		goto L242
	}
L242:
	;
	if base.Ui32(v520) < base.Ui32(v497) {
		goto L238
	} else {
		goto L243
	}
L243:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v486)+12))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v250)+12))
	if v526 < v527 {
		v583 = v486
		goto L225
	} else {
		goto L244
	}
L244:
	;
	if v527 < v526 {
		goto L238
	} else {
		goto L245
	}
L245:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v486)+16))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v250)+16))
	if base.Ui32(v530) < base.Ui32(v531) {
		v583 = v486
		goto L225
	} else {
		goto L246
	}
L246:
	;
	goto L238
L247:
	;
	v595 = v250
	goto L224
L248:
	;
	goto L249
L249:
	;
	if base.Ui32(v517) < base.Ui32(v500) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v595 = v485
	goto L224
L251:
	;
	goto L252
L252:
	;
	if base.Ui32(v498) < base.Ui32(v519) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v595 = v250
	goto L224
L254:
	;
	goto L255
L255:
	;
	if base.Ui32(v519) < base.Ui32(v498) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v595 = v485
	goto L224
L257:
	;
	goto L258
L258:
	;
	if base.Ui32(v499) < base.Ui32(v520) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v595 = v250
	goto L224
L260:
	;
	goto L261
L261:
	;
	if base.Ui32(v520) < base.Ui32(v499) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v595 = v485
	goto L224
L263:
	;
	goto L264
L264:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v485)+12))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v250)+12))
	if v541 < v542 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v595 = v250
	goto L224
L266:
	;
	goto L267
L267:
	;
	if v542 < v541 {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v595 = v485
	goto L224
L269:
	;
	goto L270
L270:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v485)+16))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v250)+16))
	if base.Ui32(v545) < base.Ui32(v546) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v548 = v250
	goto L273
L272:
	;
	v548 = v485
	goto L273
L273:
	;
	v595 = v548
	goto L224
L274:
	;
	if base.Ui32(v500) < base.Ui32(v553) {
		goto L284
	} else {
		goto L285
	}
L275:
	;
	if base.Ui32(v553) < base.Ui32(v501) {
		v583 = v486
		goto L225
	} else {
		goto L276
	}
L276:
	;
	if base.Ui32(v496) < base.Ui32(v551) {
		goto L274
	} else {
		goto L277
	}
L277:
	;
	if base.Ui32(v551) < base.Ui32(v496) {
		v583 = v486
		goto L225
	} else {
		goto L278
	}
L278:
	;
	if base.Ui32(v497) < base.Ui32(v552) {
		goto L274
	} else {
		goto L279
	}
L279:
	;
	if base.Ui32(v552) < base.Ui32(v497) {
		v583 = v486
		goto L225
	} else {
		goto L280
	}
L280:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v486)+12))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v250)+12))
	if v560 < v561 {
		goto L274
	} else {
		goto L281
	}
L281:
	;
	if v561 < v560 {
		v583 = v486
		goto L225
	} else {
		goto L282
	}
L282:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v486)+16))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v250)+16))
	if base.Ui32(v565) < base.Ui32(v564) {
		v583 = v486
		goto L225
	} else {
		goto L283
	}
L283:
	;
	goto L274
L284:
	;
	v595 = v485
	goto L224
L285:
	;
	goto L286
L286:
	;
	if base.Ui32(v553) < base.Ui32(v500) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v595 = v250
	goto L224
L288:
	;
	goto L289
L289:
	;
	if base.Ui32(v498) < base.Ui32(v551) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v595 = v485
	goto L224
L291:
	;
	goto L292
L292:
	;
	if base.Ui32(v551) < base.Ui32(v498) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v595 = v250
	goto L224
L294:
	;
	goto L295
L295:
	;
	if base.Ui32(v499) < base.Ui32(v552) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v595 = v485
	goto L224
L297:
	;
	goto L298
L298:
	;
	if base.Ui32(v552) < base.Ui32(v499) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v595 = v250
	goto L224
L300:
	;
	goto L301
L301:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v485)+12))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v250)+12))
	if v575 < v576 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v595 = v485
	goto L224
L303:
	;
	goto L304
L304:
	;
	if v576 < v575 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v595 = v250
	goto L224
L306:
	;
	goto L307
L307:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v485)+16))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v250)+16))
	if base.Ui32(v579) < base.Ui32(v580) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v582 = v485
	goto L310
L309:
	;
	v582 = v250
	goto L310
L310:
	;
	v583 = v582
	goto L225
L311:
	;
	v711 = v710
	goto L45
L312:
	;
	v710 = v698
	goto L311
L313:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v597)))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v597)+8))
	if base.Ui32(v616) < base.Ui32(v668) {
		goto L361
	} else {
		goto L362
	}
L314:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v597)+8))
	if base.Ui32(v616) < base.Ui32(v632) {
		v698 = v596
		goto L312
	} else {
		goto L324
	}
L315:
	;
	if base.Ui32(v616) < base.Ui32(v615) {
		goto L313
	} else {
		goto L316
	}
L316:
	;
	if base.Ui32(v613) < base.Ui32(v611) {
		goto L314
	} else {
		goto L317
	}
L317:
	;
	if base.Ui32(v611) < base.Ui32(v613) {
		goto L313
	} else {
		goto L318
	}
L318:
	;
	if base.Ui32(v614) < base.Ui32(v612) {
		goto L314
	} else {
		goto L319
	}
L319:
	;
	if base.Ui32(v612) < base.Ui32(v614) {
		goto L313
	} else {
		goto L320
	}
L320:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v600)+12))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v596)+12))
	if v623 < v624 {
		goto L314
	} else {
		goto L321
	}
L321:
	;
	if v624 < v623 {
		goto L313
	} else {
		goto L322
	}
L322:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v600)+16))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v596)+16))
	if base.Ui32(v628) <= base.Ui32(v627) {
		goto L313
	} else {
		goto L323
	}
L323:
	;
	goto L314
L324:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v597)))
	if base.Ui32(v632) < base.Ui32(v616) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	if base.Ui32(v615) < base.Ui32(v632) {
		goto L334
	} else {
		goto L335
	}
L326:
	;
	if base.Ui32(v611) < base.Ui32(v634) {
		v698 = v596
		goto L312
	} else {
		goto L327
	}
L327:
	;
	if base.Ui32(v634) < base.Ui32(v611) {
		goto L325
	} else {
		goto L328
	}
L328:
	;
	if base.Ui32(v612) < base.Ui32(v635) {
		v698 = v596
		goto L312
	} else {
		goto L329
	}
L329:
	;
	if base.Ui32(v635) < base.Ui32(v612) {
		goto L325
	} else {
		goto L330
	}
L330:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v596)+12))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v597)+12))
	if v641 < v642 {
		v698 = v596
		goto L312
	} else {
		goto L331
	}
L331:
	;
	if v642 < v641 {
		goto L325
	} else {
		goto L332
	}
L332:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v596)+16))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v597)+16))
	if base.Ui32(v645) < base.Ui32(v646) {
		v698 = v596
		goto L312
	} else {
		goto L333
	}
L333:
	;
	goto L325
L334:
	;
	v710 = v597
	goto L311
L335:
	;
	goto L336
L336:
	;
	if base.Ui32(v632) < base.Ui32(v615) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v710 = v600
	goto L311
L338:
	;
	goto L339
L339:
	;
	if base.Ui32(v613) < base.Ui32(v634) {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v710 = v597
	goto L311
L341:
	;
	goto L342
L342:
	;
	if base.Ui32(v634) < base.Ui32(v613) {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v710 = v600
	goto L311
L344:
	;
	goto L345
L345:
	;
	if base.Ui32(v614) < base.Ui32(v635) {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v710 = v597
	goto L311
L347:
	;
	goto L348
L348:
	;
	if base.Ui32(v635) < base.Ui32(v614) {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v710 = v600
	goto L311
L350:
	;
	goto L351
L351:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v600)+12))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v597)+12))
	if v656 < v657 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v710 = v597
	goto L311
L353:
	;
	goto L354
L354:
	;
	if v657 < v656 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v710 = v600
	goto L311
L356:
	;
	goto L357
L357:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v600)+16))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v597)+16))
	if base.Ui32(v660) < base.Ui32(v661) {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v663 = v597
	goto L360
L359:
	;
	v663 = v600
	goto L360
L360:
	;
	v710 = v663
	goto L311
L361:
	;
	if base.Ui32(v615) < base.Ui32(v668) {
		goto L371
	} else {
		goto L372
	}
L362:
	;
	if base.Ui32(v668) < base.Ui32(v616) {
		v698 = v596
		goto L312
	} else {
		goto L363
	}
L363:
	;
	if base.Ui32(v611) < base.Ui32(v666) {
		goto L361
	} else {
		goto L364
	}
L364:
	;
	if base.Ui32(v666) < base.Ui32(v611) {
		v698 = v596
		goto L312
	} else {
		goto L365
	}
L365:
	;
	if base.Ui32(v612) < base.Ui32(v667) {
		goto L361
	} else {
		goto L366
	}
L366:
	;
	if base.Ui32(v667) < base.Ui32(v612) {
		v698 = v596
		goto L312
	} else {
		goto L367
	}
L367:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v596)+12))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v597)+12))
	if v675 < v676 {
		goto L361
	} else {
		goto L368
	}
L368:
	;
	if v676 < v675 {
		v698 = v596
		goto L312
	} else {
		goto L369
	}
L369:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v596)+16))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v597)+16))
	if base.Ui32(v680) < base.Ui32(v679) {
		v698 = v596
		goto L312
	} else {
		goto L370
	}
L370:
	;
	goto L361
L371:
	;
	v710 = v600
	goto L311
L372:
	;
	goto L373
L373:
	;
	if base.Ui32(v668) < base.Ui32(v615) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v710 = v597
	goto L311
L375:
	;
	goto L376
L376:
	;
	if base.Ui32(v613) < base.Ui32(v666) {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v710 = v600
	goto L311
L378:
	;
	goto L379
L379:
	;
	if base.Ui32(v666) < base.Ui32(v613) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v710 = v597
	goto L311
L381:
	;
	goto L382
L382:
	;
	if base.Ui32(v614) < base.Ui32(v667) {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v710 = v600
	goto L311
L384:
	;
	goto L385
L385:
	;
	if base.Ui32(v667) < base.Ui32(v614) {
		goto L386
	} else {
		goto L387
	}
L386:
	;
	v710 = v597
	goto L311
L387:
	;
	goto L388
L388:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v600)+12))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v597)+12))
	if v690 < v691 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v710 = v600
	goto L311
L390:
	;
	goto L391
L391:
	;
	if v691 < v690 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v710 = v597
	goto L311
L393:
	;
	goto L394
L394:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v600)+16))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v597)+16))
	if base.Ui32(v694) < base.Ui32(v695) {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v697 = v600
	goto L397
L396:
	;
	v697 = v597
	goto L397
L397:
	;
	v698 = v697
	goto L312
L398:
	;
	if base.Ui32(v751) < base.Ui32(v752) {
		v844 = v752
		v850 = v758
		goto L400
	} else {
		goto L401
	}
L400:
	;
	if base.Ui32(v844) <= base.Ui32(v751) {
		goto L417
	} else {
		goto L418
	}
L401:
	;
	v771 = v752
	v777 = v758
	goto L402
L402:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v771)+8))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if base.Ui32(v786) < base.Ui32(v787) {
		v836 = v777
		goto L404
	} else {
		goto L405
	}
L403:
	;
	v844 = v839
	v850 = v836
	goto L400
L404:
	;
	v839 = v771 + int32(20)
	if base.Ui32(v839) <= base.Ui32(v751) {
		v771 = v839
		v777 = v836
		goto L402
	} else {
		goto L415
	}
L405:
	;
	if base.Ui32(v787) < base.Ui32(v786) {
		v844 = v771
		v850 = v777
		goto L400
	} else {
		goto L406
	}
L406:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v771)+4))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if base.Ui32(v790) < base.Ui32(v791) {
		v836 = v777
		goto L404
	} else {
		goto L407
	}
L407:
	;
	if base.Ui32(v791) < base.Ui32(v790) {
		v844 = v771
		v850 = v777
		goto L400
	} else {
		goto L408
	}
L408:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v771)))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if base.Ui32(v794) < base.Ui32(v795) {
		v836 = v777
		goto L404
	} else {
		goto L409
	}
L409:
	;
	if base.Ui32(v795) < base.Ui32(v794) {
		v844 = v771
		v850 = v777
		goto L400
	} else {
		goto L410
	}
L410:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v771)+12))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v798 < v799 {
		v836 = v777
		goto L404
	} else {
		goto L411
	}
L411:
	;
	if v799 < v798 {
		v844 = v771
		v850 = v777
		goto L400
	} else {
		goto L412
	}
L412:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v771)+16))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if base.Ui32(v802) < base.Ui32(v803) {
		v836 = v777
		goto L404
	} else {
		goto L413
	}
L413:
	;
	if base.Ui32(v803) < base.Ui32(v802) {
		v844 = v771
		v850 = v777
		goto L400
	} else {
		goto L414
	}
L414:
	;
	v806 = int32(16)
	v807 = v777 + v806
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v807)))
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = v808
	v810 = int32(8)
	v811 = v777 + v810
	v812 = *(*int64)(unsafe.Add(mBase, uint32(v811)))
	*(*int64)(unsafe.Add(mBase, uint32(v724))) = v812
	v814 = *(*int64)(unsafe.Add(mBase, uint32(v777)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v814
	v817 = v771 + v806
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	*(*int32)(unsafe.Add(mBase, uint32(v807))) = v818
	v821 = v771 + v810
	v822 = *(*int64)(unsafe.Add(mBase, uint32(v821)))
	*(*int64)(unsafe.Add(mBase, uint32(v811))) = v822
	v824 = *(*int64)(unsafe.Add(mBase, uint32(v771)))
	*(*int64)(unsafe.Add(mBase, uint32(v777))) = v824
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v718)))
	*(*int32)(unsafe.Add(mBase, uint32(v817))) = v826
	v828 = *(*int64)(unsafe.Add(mBase, uint32(v724)))
	*(*int64)(unsafe.Add(mBase, uint32(v821))) = v828
	v830 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v771))) = v830
	v836 = v777 + int32(20)
	goto L404
L415:
	;
	goto L403
L416:
	;
	v1134 = int32(16)
	v1135 = v844 + v1134
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1135)))
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = v1136
	v1138 = int32(8)
	v1139 = v844 + v1138
	v1140 = *(*int64)(unsafe.Add(mBase, uint32(v1139)))
	*(*int64)(unsafe.Add(mBase, uint32(v724))) = v1140
	v1142 = *(*int64)(unsafe.Add(mBase, uint32(v844)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v1142
	v1145 = v862 + v1134
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1145)))
	*(*int32)(unsafe.Add(mBase, uint32(v1135))) = v1146
	v1149 = v862 + v1138
	v1150 = *(*int64)(unsafe.Add(mBase, uint32(v1149)))
	*(*int64)(unsafe.Add(mBase, uint32(v1139))) = v1150
	v1152 = *(*int64)(unsafe.Add(mBase, uint32(v862)))
	*(*int64)(unsafe.Add(mBase, uint32(v844))) = v1152
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v718)))
	*(*int32)(unsafe.Add(mBase, uint32(v1145))) = v1154
	v1156 = *(*int64)(unsafe.Add(mBase, uint32(v724)))
	*(*int64)(unsafe.Add(mBase, uint32(v1149))) = v1156
	v1158 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v862))) = v1158
	v1160 = int32(20)
	v751 = v862 - v1160
	v752 = v844 + v1160
	v758 = v850
	v762 = v873
	goto L398
L417:
	;
	v862 = v751
	v873 = v762
	goto L420
L418:
	;
	v935 = v751
	v946 = v762
	goto L419
L419:
	;
	v952 = int32(20)
	v953 = base.I32_div_s(v850-v23, v952)
	v956 = base.I32_div_s(v844-v850, v952)
	if v953 < v956 {
		goto L434
	} else {
		goto L435
	}
L420:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v862)+8))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if base.Ui32(v878) < base.Ui32(v879) {
		goto L416
	} else {
		goto L422
	}
L421:
	;
	v935 = v931
	v946 = v929
	goto L419
L422:
	;
	if base.Ui32(v879) < base.Ui32(v878) {
		v929 = v873
		goto L423
	} else {
		goto L424
	}
L423:
	;
	v931 = v862 - int32(20)
	if base.Ui32(v844) <= base.Ui32(v931) {
		v862 = v931
		v873 = v929
		goto L420
	} else {
		goto L433
	}
L424:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v862)+4))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if base.Ui32(v882) < base.Ui32(v883) {
		goto L416
	} else {
		goto L425
	}
L425:
	;
	if base.Ui32(v883) < base.Ui32(v882) {
		v929 = v873
		goto L423
	} else {
		goto L426
	}
L426:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v862)))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if base.Ui32(v886) < base.Ui32(v887) {
		goto L416
	} else {
		goto L427
	}
L427:
	;
	if base.Ui32(v887) < base.Ui32(v886) {
		v929 = v873
		goto L423
	} else {
		goto L428
	}
L428:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v862)+12))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v890 < v891 {
		goto L416
	} else {
		goto L429
	}
L429:
	;
	if v891 < v890 {
		v929 = v873
		goto L423
	} else {
		goto L430
	}
L430:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v862)+16))
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if base.Ui32(v894) < base.Ui32(v895) {
		goto L416
	} else {
		goto L431
	}
L431:
	;
	if base.Ui32(v895) < base.Ui32(v894) {
		v929 = v873
		goto L423
	} else {
		goto L432
	}
L432:
	;
	v898 = int32(16)
	v899 = v862 + v898
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v899)))
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = v900
	v902 = int32(8)
	v903 = v862 + v902
	v904 = *(*int64)(unsafe.Add(mBase, uint32(v903)))
	*(*int64)(unsafe.Add(mBase, uint32(v724))) = v904
	v906 = *(*int64)(unsafe.Add(mBase, uint32(v862)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v906
	v909 = v873 + v898
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v909)))
	*(*int32)(unsafe.Add(mBase, uint32(v899))) = v910
	v913 = v873 + v902
	v914 = *(*int64)(unsafe.Add(mBase, uint32(v913)))
	*(*int64)(unsafe.Add(mBase, uint32(v903))) = v914
	v916 = *(*int64)(unsafe.Add(mBase, uint32(v873)))
	*(*int64)(unsafe.Add(mBase, uint32(v862))) = v916
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v718)))
	*(*int32)(unsafe.Add(mBase, uint32(v909))) = v918
	v920 = *(*int64)(unsafe.Add(mBase, uint32(v724)))
	*(*int64)(unsafe.Add(mBase, uint32(v913))) = v920
	v922 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v873))) = v922
	v929 = v873 - int32(20)
	goto L423
L433:
	;
	goto L421
L434:
	;
	v958 = v953
	goto L436
L435:
	;
	v958 = v956
	goto L436
L436:
	;
	if v958 != 0 {
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v972 = int32(0)
	goto L440
L438:
	;
	goto L439
L439:
	;
	v1033 = int32(20)
	v1034 = base.I32_div_s(v946-v935, v1033)
	v1037 = base.I32_div_s(v63-v946, v1033)
	v1039 = v1037 - int32(1)
	if v1034 < v1039 {
		goto L443
	} else {
		goto L444
	}
L440:
	;
	v982 = v972 * int32(20)
	v983 = v23 + v982
	v984 = int32(16)
	v985 = v983 + v984
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v985)))
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = v986
	v988 = int32(8)
	v989 = v983 + v988
	v990 = *(*int64)(unsafe.Add(mBase, uint32(v989)))
	*(*int64)(unsafe.Add(mBase, uint32(v724))) = v990
	v992 = *(*int64)(unsafe.Add(mBase, uint32(v983)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v992
	v994 = v982 + (v844 + v958*int32(-20))
	v996 = v994 + v984
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v996)))
	*(*int32)(unsafe.Add(mBase, uint32(v985))) = v997
	v1000 = v994 + v988
	v1001 = *(*int64)(unsafe.Add(mBase, uint32(v1000)))
	*(*int64)(unsafe.Add(mBase, uint32(v989))) = v1001
	v1003 = *(*int64)(unsafe.Add(mBase, uint32(v994)))
	*(*int64)(unsafe.Add(mBase, uint32(v983))) = v1003
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v718)))
	*(*int32)(unsafe.Add(mBase, uint32(v996))) = v1005
	v1007 = *(*int64)(unsafe.Add(mBase, uint32(v724)))
	*(*int64)(unsafe.Add(mBase, uint32(v1000))) = v1007
	v1009 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v994))) = v1009
	v1012 = v972 + int32(1)
	if v1012 != v958 {
		v972 = v1012
		goto L440
	} else {
		goto L442
	}
L441:
	;
	goto L439
L442:
	;
	goto L441
L443:
	;
	v1041 = v1034
	goto L445
L444:
	;
	v1041 = v1039
	goto L445
L445:
	;
	if v1041 != 0 {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v1051 = int32(0)
	goto L449
L447:
	;
	goto L448
L448:
	;
	if base.Ui32(v956) <= base.Ui32(v1034) {
		goto L452
	} else {
		goto L453
	}
L449:
	;
	v1065 = v1051 * int32(20)
	v1066 = v844 + v1065
	v1067 = int32(16)
	v1068 = v1066 + v1067
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1068)))
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = v1069
	v1071 = int32(8)
	v1072 = v1066 + v1071
	v1073 = *(*int64)(unsafe.Add(mBase, uint32(v1072)))
	*(*int64)(unsafe.Add(mBase, uint32(v724))) = v1073
	v1075 = *(*int64)(unsafe.Add(mBase, uint32(v1066)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v1075
	v1077 = v1065 + (v63 + v1041*int32(-20))
	v1079 = v1077 + v1067
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1079)))
	*(*int32)(unsafe.Add(mBase, uint32(v1068))) = v1080
	v1083 = v1077 + v1071
	v1084 = *(*int64)(unsafe.Add(mBase, uint32(v1083)))
	*(*int64)(unsafe.Add(mBase, uint32(v1072))) = v1084
	v1086 = *(*int64)(unsafe.Add(mBase, uint32(v1077)))
	*(*int64)(unsafe.Add(mBase, uint32(v1066))) = v1086
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v718)))
	*(*int32)(unsafe.Add(mBase, uint32(v1079))) = v1088
	v1090 = *(*int64)(unsafe.Add(mBase, uint32(v724)))
	*(*int64)(unsafe.Add(mBase, uint32(v1083))) = v1090
	v1092 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1077))) = v1092
	v1095 = v1051 + int32(1)
	if v1095 != v1041 {
		v1051 = v1095
		goto L449
	} else {
		goto L451
	}
L450:
	;
	goto L448
L451:
	;
	goto L450
L452:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v956) {
		goto L455
	} else {
		goto L456
	}
L453:
	;
	goto L454
L454:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v1034) {
		goto L459
	} else {
		goto L460
	}
L455:
	;
	F_sort_pending_writebacks(m, v23, v956)
	mBase = m.M
	goto L457
L456:
	;
	goto L457
L457:
	;
	if base.Ui32(v1034) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L458
	}
L458:
	;
	v23 = v63 + v1034*int32(-20)
	v24 = v1034
	goto L1
L459:
	;
	F_sort_pending_writebacks(m, v63+v1034*int32(-20), v1034)
	mBase = m.M
	goto L461
L460:
	;
	goto L461
L461:
	;
	if base.Ui32(int32(1)) < base.Ui32(v956) {
		v44 = v956
		goto L3
	} else {
		goto L462
	}
L462:
	;
	goto L5
}
func F_sortouts_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v8 < v11 {
		return int32(-1)
	} else {
		v15 = int32(1)
		if v11 < v8 {
			v30 = v15
			return v30
		} else {
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6)+4)))
			v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
			if v17 < v18 {
				return int32(-1)
			} else {
				if v18 < v17 {
					v30 = v15
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					if v24 < v25 {
						v30 = int32(-1)
					} else {
						v30 = base.B2i32(v25 < v24)
					}
				}
				return v30
			}
		}
	}
}
func F_soundex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var __phi64 int32
	_ = __phi64
	var v65 int32
	_ = v65
	var __phi65 int32
	_ = __phi65
	var v67 int32
	_ = v67
	var __phi67 int32
	_ = __phi67
	var v68 int32
	_ = v68
	var __phi68 int32
	_ = __phi68
	var v69 int32
	_ = v69
	var __phi69 int32
	_ = __phi69
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = F_text_to_cstring(m, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = v5 + int32(11)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v174 = F_cstring_to_text(m, v5+int32(11))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L38
	}
L5:
	;
	v56 = F_toupper(m, v25&int32(255))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v56)
	v58 = int32(1)
	v60 = v5 + int32(12)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v61 != 0 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v23 = v12
	v25 = v22
	goto L9
L7:
	;
	goto L8
L8:
	;
	v50 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v50)
	goto L4
L9:
	;
	if base.Ui32(int32(229)) < base.Ui32((v25|int32(32)-int32(123))&int32(255)) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v39 != 0 {
		v23 = v23 + int32(1)
		v25 = v39
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v162 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v157))) = uint8(v162)
	goto L4
L14:
	;
	__phi64 = v23
	__phi65 = v61
	__phi67 = v60
	__phi68 = v58
	__phi69 = v23 + int32(1)
	v64 = __phi64
	v65 = __phi65
	v67 = __phi67
	v68 = __phi68
	v69 = __phi69
	goto L17
L15:
	;
	v144 = v60
	v145 = v58
	goto L16
L16:
	;
	v151 = int32(4) - v145
	v152 = F___memset(m, v144, int32(48), v151)
	mBase = m.M
	v157 = v152 + v151
	goto L13
L17:
	;
	if base.Ui32(int32(25)) < base.Ui32((v65|int32(32)-int32(97))&int32(255)) {
		v127 = v67
		v128 = v68
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if int32(3) < v128 {
		v157 = v127
		goto L13
	} else {
		goto L37
	}
L19:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v132 != 0 {
		goto L33
	} else {
		goto L34
	}
L20:
	;
	v82 = F_toupper(m, v65&int32(255))
	mBase = m.M
	v83 = base.I32_extend8_s(v82)
	v87 = base.B2i32(base.Ui32(int32(25)) < base.Ui32(v83-int32(65)))
	if v87 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v97 = F_toupper(m, v96)
	mBase = m.M
	v98 = base.I32_extend8_s(v97)
	if base.Ui32(v98-int32(65)) <= base.Ui32(int32(25)) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[1567]))))
	v93 = v92
	goto L21
L23:
	;
	goto L24
L24:
	;
	v93 = v82
	goto L21
L25:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+uint32(_consts[1567]))))
	v106 = v105
	goto L27
L26:
	;
	v106 = v97
	goto L27
L27:
	;
	if v93&int32(255) == v106&int32(255) {
		v127 = v67
		v128 = v68
		goto L19
	} else {
		goto L28
	}
L28:
	;
	if v87 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+uint32(_consts[1567]))))
	v115 = v114
	goto L31
L30:
	;
	v115 = v82
	goto L31
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v115)
	if v115&int32(255) == int32(48) {
		v127 = v67
		v128 = v68
		goto L19
	} else {
		goto L32
	}
L32:
	;
	v121 = int32(1)
	v127 = v67 + v121
	v128 = v68 + v121
	goto L19
L33:
	;
	if v128 < int32(4) {
		__phi64 = v69
		__phi65 = v132
		__phi67 = v127
		__phi68 = v128
		__phi69 = v69 + int32(1)
		v64 = __phi64
		v65 = __phi65
		v67 = __phi67
		v68 = __phi68
		v69 = __phi69
		goto L17
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L18
L36:
	;
	goto L35
L37:
	;
	v144 = v127
	v145 = v128
	goto L16
L38:
	;
	m.G0 = v5 + int32(16)
	return v174
}
func F_spgbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 float64
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int64
	_ = v224
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v14 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 == int32(0) {
			v20 = F_SpGistNewBuffer(m, l1)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = F_SpGistNewBuffer(m, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = F_SpGistNewBuffer(m, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = int32(4548900)
						v28 = *(*int32)(unsafe.Add(mBase, _consts[26]))
						*(*int32)(unsafe.Add(mBase, _consts[26])) = v28 + int32(1)
						if v20 < int32(0) {
							v35 = *(*int32)(unsafe.Add(mBase, _consts[1]))
							v41 = *(*int32)(unsafe.Add(mBase, uint32(v35+(v20^int32(-1))<<(uint(int32(2))%32))))
							v49 = v41
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, _consts[2]))
							v49 = v43 + v20<<(uint(int32(13))%32) + int32(-8192)
						}
						F_PageInit(m, v49, int32(8192), int32(8))
						mBase = m.M
						v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+16)))
						v55 = v49 + v54
						v56 = int32(65410)
						*(*uint16)(unsafe.Add(mBase, uint32(v55)+6)) = uint16(v56)
						v58 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v55))) = uint16(v58)
						v60 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v49)+80)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+72)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49-int32(-64)))) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+56)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+48)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+40)) = v60
						*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = v60
						*(*int32)(unsafe.Add(mBase, uint32(v49)+88)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = int64(-1173640210)
						v80 = int32(92)
						*(*uint16)(unsafe.Add(mBase, uint32(v49)+12)) = uint16(v80)
						v82 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v49)+84)) = v82
						*(*int32)(unsafe.Add(mBase, uint32(v49)+76)) = v82
						*(*int32)(unsafe.Add(mBase, uint32(v49)+68)) = v82
						*(*int32)(unsafe.Add(mBase, uint32(v49)+60)) = v82
						*(*int32)(unsafe.Add(mBase, uint32(v49)+52)) = v82
						*(*int32)(unsafe.Add(mBase, uint32(v49)+44)) = v82
						*(*int32)(unsafe.Add(mBase, uint32(v49)+36)) = v82
						F_MarkBufferDirty(m, v20)
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							v98 = int32(4)
							if v22 < int32(0) {
								v102 = *(*int32)(unsafe.Add(mBase, _consts[1]))
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v102+(v22^int32(-1))<<(uint(int32(2))%32))))
								v116 = v108
							} else {
								v110 = *(*int32)(unsafe.Add(mBase, _consts[2]))
								v116 = v110 + v22<<(uint(int32(13))%32) + int32(-8192)
							}
							F_PageInit(m, v116, int32(8192), int32(8))
							mBase = m.M
							v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+16)))
							v121 = v116 + v120
							v122 = int32(65410)
							*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)) = uint16(v122)
							*(*uint16)(unsafe.Add(mBase, uint32(v121))) = uint16(v98)
							F_MarkBufferDirty(m, v22)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								v127 = int32(12)
								if v24 < int32(0) {
									v131 = *(*int32)(unsafe.Add(mBase, _consts[1]))
									v137 = *(*int32)(unsafe.Add(mBase, uint32(v131+(v24^int32(-1))<<(uint(int32(2))%32))))
									v145 = v137
								} else {
									v139 = *(*int32)(unsafe.Add(mBase, _consts[2]))
									v145 = v139 + v24<<(uint(int32(13))%32) + int32(-8192)
								}
								F_PageInit(m, v145, int32(8192), int32(8))
								mBase = m.M
								v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+16)))
								v150 = v145 + v149
								v151 = int32(65410)
								*(*uint16)(unsafe.Add(mBase, uint32(v150)+6)) = uint16(v151)
								*(*uint16)(unsafe.Add(mBase, uint32(v150))) = uint16(v127)
								F_MarkBufferDirty(m, v24)
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
									return int32(0)
								} else {
									v156 = int32(4548900)
									v158 = *(*int32)(unsafe.Add(mBase, _consts[26]))
									*(*int32)(unsafe.Add(mBase, _consts[26])) = v158 - int32(1)
									F_UnlockReleaseBuffer(m, v20)
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return int32(0)
									} else {
										F_UnlockReleaseBuffer(m, v22)
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return int32(0)
										} else {
											F_UnlockReleaseBuffer(m, v24)
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												F_initSpGistState(m, v11+int32(8), l1)
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = int64(0)
													v174 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+88)) = uint8(v174)
													v177 = *(*int32)(unsafe.Add(mBase, _consts[28]))
													v182 = F_AllocSetContextCreateInternal(m, v177, int32(65643), int32(0), int32(8192), int32(8388608))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = v182
														v185 = int32(1)
														v186 = int32(0)
														v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
														v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+140))
														v196 = m.T0[v195].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v185, v186, v185, v186, int32(-1), int32(245), v11+int32(8), v186)
														mBase = m.M
														v197 = m.ExcPending
														if v197 != 0 {
															return int32(0)
														} else {
															v198 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
															F_MemoryContextDelete(m, v198)
															mBase = m.M
															v200 = m.ExcPending
															if v200 != 0 {
																return int32(0)
															} else {
																F_SpGistUpdateMetaPage(m, l1)
																mBase = m.M
																v202 = m.ExcPending
																if v202 != 0 {
																	return int32(0)
																} else {
																	v203 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
																	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+118)))
																	if v204 != int32(112) {
																		v221 = F_palloc0(m, int32(16))
																		mBase = m.M
																		v222 = m.ExcPending
																		if v222 != 0 {
																			return int32(0)
																		} else {
																			*(*float64)(unsafe.Add(mBase, uint32(v221))) = v196
																			v224 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																			*(*float64)(unsafe.Add(mBase, uint32(v221)+8)) = base.F64_convert_i64_s(v224)
																			m.G0 = v11 + int32(112)
																			return v221
																		}
																	} else {
																		v208 = *(*int32)(unsafe.Add(mBase, _consts[27]))
																		if v208 <= int32(0) {
																			v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
																			if v211 != 0 {
																				v221 = F_palloc0(m, int32(16))
																				mBase = m.M
																				v222 = m.ExcPending
																				if v222 != 0 {
																					return int32(0)
																				} else {
																					*(*float64)(unsafe.Add(mBase, uint32(v221))) = v196
																					v224 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																					*(*float64)(unsafe.Add(mBase, uint32(v221)+8)) = base.F64_convert_i64_s(v224)
																					m.G0 = v11 + int32(112)
																					return v221
																				}
																			} else {
																				v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
																				if v212 != 0 {
																					v221 = F_palloc0(m, int32(16))
																					mBase = m.M
																					v222 = m.ExcPending
																					if v222 != 0 {
																						return int32(0)
																					} else {
																						*(*float64)(unsafe.Add(mBase, uint32(v221))) = v196
																						v224 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																						*(*float64)(unsafe.Add(mBase, uint32(v221)+8)) = base.F64_convert_i64_s(v224)
																						m.G0 = v11 + int32(112)
																						return v221
																					}
																				} else {
																					v213 = int32(0)
																					v215 = F_RelationGetNumberOfBlocksInFork(m, l1, v213)
																					mBase = m.M
																					v216 = m.ExcPending
																					if v216 != 0 {
																						return int32(0)
																					} else {
																						F_log_newpage_range(m, l1, v213, v215, int32(1))
																						mBase = m.M
																						v219 = m.ExcPending
																						if v219 != 0 {
																							return int32(0)
																						} else {
																							v221 = F_palloc0(m, int32(16))
																							mBase = m.M
																							v222 = m.ExcPending
																							if v222 != 0 {
																								return int32(0)
																							} else {
																								*(*float64)(unsafe.Add(mBase, uint32(v221))) = v196
																								v224 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																								*(*float64)(unsafe.Add(mBase, uint32(v221)+8)) = base.F64_convert_i64_s(v224)
																								m.G0 = v11 + int32(112)
																								return v221
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v213 = int32(0)
																			v215 = F_RelationGetNumberOfBlocksInFork(m, l1, v213)
																			mBase = m.M
																			v216 = m.ExcPending
																			if v216 != 0 {
																				return int32(0)
																			} else {
																				F_log_newpage_range(m, l1, v213, v215, int32(1))
																				mBase = m.M
																				v219 = m.ExcPending
																				if v219 != 0 {
																					return int32(0)
																				} else {
																					v221 = F_palloc0(m, int32(16))
																					mBase = m.M
																					v222 = m.ExcPending
																					if v222 != 0 {
																						return int32(0)
																					} else {
																						*(*float64)(unsafe.Add(mBase, uint32(v221))) = v196
																						v224 = *(*int64)(unsafe.Add(mBase, uint32(v11)+96))
																						*(*float64)(unsafe.Add(mBase, uint32(v221)+8)) = base.F64_convert_i64_s(v224)
																						m.G0 = v11 + int32(112)
																						return v221
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
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v234 = m.ExcPending
			if v234 != 0 {
				return int32(0)
			} else {
				v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v235 + int32(4)
				F_errmsg_internal(m, int32(528903), v11)
				mBase = m.M
				v241 = m.ExcPending
				if v241 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515302), int32(84), int32(450405))
					mBase = m.M
					v246 = m.ExcPending
					if v246 != 0 {
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
func F_spgcanreturn(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = int32(1)
	if l1 <= v4 {
		v7 = F_spgGetCache(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
			v12 = v11
			return v12 & int32(1)
		}
	} else {
		v12 = v4
		return v12 & int32(1)
	}
}
func F_spgendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+88))
	F_MemoryContextDelete(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+92))
	F_MemoryContextDelete(m, v8)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+104))
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_pfree(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+68))
	if v14 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v4)+72))
	if v22 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v14 == v18 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_FreeTupleDesc(m, v14)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	F_pfree(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v25 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v4)+120))
	F_pfree(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_pfree(m, v4)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L25
	}
L19:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v4)+124))
	F_pfree(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v4)+188))
	F_pfree(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v4)+192))
	F_pfree(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_pfree(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	F_pfree(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	return
}
func F_spggettuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	if l1 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+208)) = uint8(v13)
	v16 = v12 + int32(3488)
	v18 = v12 + int32(5120)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
	v21 = v19
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
	} else {
		goto L33
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	if v29 < v21 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return base.B2i32(v29 < v21)
L6:
	;
	v35 = v12 + v29*int32(6)
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+228)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint16(v36)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+224))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v40+int32(2672)))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v44)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v16+v46<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if int32(0) < v52 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v12)+108))
	if v76 <= int32(0) {
		v107 = v21
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v18+v56<<(uint(int32(2))%32))))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v12+int32(3080)))))
	F_index_store_float8_orderby_distances(m, l0, v55, v60, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+220)) = v70 + int32(1)
	return base.B2i32(v29 < v21)
L12:
	;
	return int32(0)
L13:
	;
	goto L11
L14:
	;
	if v107 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	v79 = int32(0)
	if v21 <= v79 {
		v107 = v21
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v83 = v79
	v87 = v21
	goto L17
L17:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v18+v83<<(uint(int32(2))%32))))
	if v94 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v107 = v98
	goto L14
L19:
	;
	F_pfree(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L22
	}
L20:
	;
	v98 = v87
	goto L21
L21:
	;
	v100 = v83 + int32(1)
	if v100 < v98 {
		v83 = v100
		v87 = v98
		goto L17
	} else {
		goto L23
	}
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
	v98 = v97
	goto L21
L23:
	;
	goto L18
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+216)) = int64(0)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_spgWalk(m, v149, v12, int32(0), int32(258))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L12
	} else {
		goto L31
	}
L25:
	;
	v113 = int32(0)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+208)))
	if v114&int32(1) == v113 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v120 = v113
	goto L27
L27:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v16+v120<<(uint(int32(2))%32))))
	F_pfree(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L29
	}
L28:
	;
	goto L24
L29:
	;
	v135 = v120 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
	if v135 < v136 {
		v120 = v135
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
	if v154 != 0 {
		v21 = v154
		goto L4
	} else {
		goto L32
	}
L32:
	;
	goto L5
L33:
	;
	F_errmsg_internal(m, int32(267179), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(520154), int32(1031), int32(400985))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_split_pathtarget_at_srfs_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
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
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v407 int32
	_ = v407
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v526 int32
	_ = v526
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v627 int32
	_ = v627
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	v6 = l5
	v7 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(96)
	m.G0 = v21
	if l1 == l2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v21 + int32(96)
	return
L2:
	;
	if v149&int32(1) != 0 {
		goto L47
	} else {
		goto L48
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v211
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = l1
	v29 = F_list_make1_impl(m, int32(1), v21+int32(4))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+60)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = l0
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v29
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v32
	v37 = F_list_make1_impl(m, int32(471), v21)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v211 = v37
	goto L3
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v43 = v42
	goto L12
L11:
	;
	v43 = int32(0)
	goto L12
L12:
	;
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v44
	v52 = F_list_make1_impl(m, int32(1), v21+int32(24))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v54
	v62 = F_list_make1_impl(m, int32(1), v21+int32(20))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v64 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v64
	v72 = F_list_make1_impl(m, int32(1), v21+int32(16))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+80)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v72
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v77 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v78 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l1
	v181 = F_list_make1_impl(m, int32(1), v21+int32(12))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L44
	}
L19:
	;
	if v145 != 0 {
		goto L2
	} else {
		goto L43
	}
L20:
	;
	v145 = int32(0)
	v149 = v7
	goto L19
L21:
	;
	goto L22
L22:
	;
	v88 = int32(0)
	v91 = v7
	v92 = v7
	goto L23
L23:
	;
	v102 = v91 << (uint(int32(2)) % 32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v102+v103)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v107 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v145 = v133
	v149 = v134
	goto L19
L25:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107+v102)))
	v110 = v109
	goto L27
L26:
	;
	v110 = int32(0)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v110
	v116 = F_split_pathtarget_walker(m, v105, v21+int32(56))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v21)+88))
	if v118 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v118 < v88 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v133 = v88
	v134 = v92
	goto L31
L31:
	;
	v137 = v91 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v137 < v138 {
		v88 = v133
		v91 = v137
		v92 = v134
		goto L23
	} else {
		goto L42
	}
L32:
	;
	v120 = v88
	goto L34
L33:
	;
	v120 = v118
	goto L34
L34:
	;
	v122 = base.B2i32(v118 <= v88) & v92
	if v118 < v88 {
		v132 = v122
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v133 = v120
	v134 = v132
	goto L31
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	switch v124 - int32(15) {
	case 0:
		goto L39
	default:
		goto L37
	case 2:
		goto L38
	}
L37:
	;
	v132 = int32(1)
	goto L35
L38:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+16)))
	if v130 != 0 {
		v132 = v122
		goto L35
	} else {
		goto L41
	}
L39:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+12)))
	if v127 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v132 = v122
	goto L35
L41:
	;
	goto L37
L42:
	;
	goto L24
L43:
	;
	goto L18
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v181
	v184 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v184
	v191 = F_list_make1_impl(m, int32(471), v21+int32(8))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	v211 = v191
	goto L3
L46:
	;
	v252 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v252
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v252
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	v271 = v7
	v272 = v252
	goto L55
L47:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	v217 = F_lappend(m, v215, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v231 = v145 << (uint(int32(2)) % 32)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v234 = v231 + v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	v237 = F_list_concat(m, v235, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L7
	} else {
		goto L53
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v217
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	v222 = F_lappend(m, v220, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v222
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
	v227 = F_lappend(m, v225, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v227
	v251 = v227
	goto L46
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v237
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v242 = v241 + v231
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
	v245 = F_list_concat(m, v243, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v245
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	v251 = v248
	goto L46
L55:
	;
	v277 = int32(0)
	if v257 == v277 {
		v287 = v277
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v288 = int32(0)
	if v256 == v288 {
		v297 = v288
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	if v281 <= v271 {
		v287 = int32(0)
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v287 = v283 + v271<<(uint(int32(2))%32)
	goto L57
L60:
	;
	if v251 == int32(0) {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v291 <= v271 {
		v297 = v288
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v297 = v293 + v271<<(uint(int32(2))%32)
	goto L60
L63:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	if v300 <= v271 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v287 == int32(0) {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v297 == int32(0) {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	v309 = v306 + v271<<(uint(int32(2))%32)
	if v309 == int32(0) {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v314 = v287 + int32(4)
	if v314 == int32(0) {
		v627 = l1
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v638 = F_lappend(m, v637, v627)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L7
	} else {
		goto L122
	}
L69:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+12))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
	if base.Ui32(v318+v319<<(uint(int32(2))%32)) <= base.Ui32(v314) {
		v627 = l1
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v325 = F_palloc0(m, int32(40))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = int32(277)
	if v312 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	if v382 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L73:
	;
	v331 = int32(0)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	if v332 <= v331 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v337 = v331
	goto L75
L75:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v353+v337<<(uint(int32(2))%32))))
	F_add_sp_item_to_pathtarget(m, v325, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L7
	} else {
		goto L77
	}
L76:
	;
	goto L72
L77:
	;
	v361 = v337 + int32(1)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	if v361 < v362 {
		v337 = v361
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	if v498 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L80:
	;
	v386 = v297 + int32(4)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v382)+12))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if base.Ui32(v386) < base.Ui32(v388+v389<<(uint(int32(2))%32)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v394 = v386
	goto L83
L82:
	;
	v394 = int32(0)
	goto L83
L83:
	;
	if v394 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v398 = (v394 - v388) >> (uint(int32(2)) % 32)
	goto L86
L85:
	;
	v398 = v389
	goto L86
L86:
	;
	if v389 <= v398 {
		goto L79
	} else {
		goto L87
	}
L87:
	;
	v407 = v398
	goto L88
L88:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v382)+12))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v418+v407<<(uint(int32(2))%32))))
	if v422 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L79
L90:
	;
	v477 = v407 + int32(1)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if v477 < v478 {
		v407 = v477
		goto L88
	} else {
		goto L97
	}
L91:
	;
	v425 = int32(0)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v422)+4))
	if v426 <= v425 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v431 = v425
	goto L93
L93:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v422)+12))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v447+v431<<(uint(int32(2))%32))))
	F_add_sp_item_to_pathtarget(m, v325, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L7
	} else {
		goto L95
	}
L94:
	;
	goto L90
L95:
	;
	v455 = v431 + int32(1)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v422)+4))
	if v455 < v456 {
		v431 = v455
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	goto L89
L98:
	;
	v617 = F_set_pathtarget_cost_width(m, l0, v325)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L7
	} else {
		goto L121
	}
L99:
	;
	v502 = v309 + int32(4)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v498)+12))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if base.Ui32(v502) < base.Ui32(v504+v505<<(uint(int32(2))%32)) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v510 = v502
	goto L102
L101:
	;
	v510 = int32(0)
	goto L102
L102:
	;
	if v510 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v514 = (v510 - v504) >> (uint(int32(2)) % 32)
	goto L105
L104:
	;
	v514 = v505
	goto L105
L105:
	;
	if v505 <= v514 {
		goto L98
	} else {
		goto L106
	}
L106:
	;
	v526 = v514
	goto L107
L107:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v498)+12))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v534+v526<<(uint(int32(2))%32))))
	if v538 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	goto L98
L109:
	;
	v596 = v526 + int32(1)
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if v596 < v597 {
		v526 = v596
		goto L107
	} else {
		goto L120
	}
L110:
	;
	v541 = int32(0)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	if v542 <= v541 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v547 = v541
	goto L112
L112:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v538)+12))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v563+v547<<(uint(int32(2))%32))))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	v569 = F_list_member(m, v272, v568)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L7
	} else {
		goto L114
	}
L113:
	;
	goto L109
L114:
	;
	if v569 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	F_add_sp_item_to_pathtarget(m, v325, v567)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L7
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v574 = v547 + int32(1)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	if v574 < v575 {
		v547 = v574
		goto L112
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	goto L113
L120:
	;
	goto L108
L121:
	;
	v627 = v325
	goto L68
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v638
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v644 = F_lappend_int(m, v641, base.B2i32(v312 != int32(0)))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L7
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v644
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v627)+4))
	v271 = v271 + int32(1)
	v272 = v649
	goto L55
}
func F_str_tolower(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v4 {
		v66 = v4
		m.G0 = v10 + int32(16)
		return v66
	} else {
		if l2 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(713732)
					F_errmsg(m, int32(264303), v10)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(601156), int32(0))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(521728), int32(1655), int32(224592))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
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
		} else {
			v16 = F_pg_newlocale_from_collation(m, l2)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+3)))
				if v20 == int32(1) {
					v23 = F_pnstrdup(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
						if v25 == int32(0) {
							v66 = v23
						} else {
							v28 = v23
							v30 = v25
							for {
								v35 = int32(255)
								v36 = v30 & v35
								if base.Ui32((v36-int32(65))&v35) < base.Ui32(int32(26)) {
									v45 = v36 | int32(32)
								} else {
									v45 = v36
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v45)
								v48 = v28 + int32(1)
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
								if v49 != 0 {
									v28 = v48
									v30 = v49
									continue
								} else {
									break
								}
								break
							}
							v66 = v23
						}
						m.G0 = v10 + int32(16)
						return v66
					}
				} else {
					v51 = l1 + int32(1)
					v52 = F_palloc(m, v51)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = F_pg_strlower(m, v52, v51, l0, l1, v16)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v57 = v54 + int32(1)
							if base.Ui32(v57) <= base.Ui32(v51) {
								v66 = v52
								m.G0 = v10 + int32(16)
								return v66
							} else {
								v59 = F_repalloc(m, v52, v57)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v61 = F_pg_strlower(m, v59, v57, l0, l1, v16)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v66 = v59
										m.G0 = v10 + int32(16)
										return v66
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
func F_strict_word_similarity(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 float32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v17 = v10 + int32(1)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v22 = int32(1)
			v23 = v21 & v22
			if v21 == v22 {
				v26 = int32(4)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v28&int32(254) == int32(2) {
					v37 = v26
				} else {
					v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
				}
				if v28 == int32(1) {
					v40 = v26
				} else {
					v40 = v37
				}
				v51 = v40
			} else {
				v41 = int32(1)
				if v23 != 0 {
					v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v23 != 0 {
				v52 = v17
			} else {
				v52 = v10 + int32(4)
			}
			v53 = int32(1)
			v54 = v19 + v53
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			v59 = v57 & v53
			if v59 != 0 {
				v60 = v54
			} else {
				v60 = v19 + int32(4)
			}
			if v57 == int32(1) {
				v63 = int32(4)
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				if v65&int32(254) == int32(2) {
					v74 = v63
				} else {
					v74 = base.B2i32(v65 == int32(18)) << (uint(v63) % 32)
				}
				if v65 == int32(1) {
					v77 = v63
				} else {
					v77 = v74
				}
				v88 = v77
			} else {
				v78 = int32(1)
				if v59 != 0 {
					v88 = int32(base.Ui32(v57)>>(uint(v78)%32)) - v78
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v90 = F_calc_word_similarity(m, v52, v51, v60, v88, int32(2))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v92 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v96 != v19 {
							F_pfree(m, v19)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								return base.I32_reinterpret_f32(v90)
							}
						} else {
							return base.I32_reinterpret_f32(v90)
						}
					}
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v96 != v19 {
						F_pfree(m, v19)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(v90)
						}
					} else {
						return base.I32_reinterpret_f32(v90)
					}
				}
			}
		}
	}
}
func F_strict_word_similarity_op(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 float32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 float64
	_ = v101
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v17 = v10 + int32(1)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v22 = int32(1)
			v23 = v21 & v22
			if v21 == v22 {
				v26 = int32(4)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v28&int32(254) == int32(2) {
					v37 = v26
				} else {
					v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
				}
				if v28 == int32(1) {
					v40 = v26
				} else {
					v40 = v37
				}
				v51 = v40
			} else {
				v41 = int32(1)
				if v23 != 0 {
					v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v23 != 0 {
				v52 = v17
			} else {
				v52 = v10 + int32(4)
			}
			v53 = int32(1)
			v54 = v19 + v53
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			v59 = v57 & v53
			if v59 != 0 {
				v60 = v54
			} else {
				v60 = v19 + int32(4)
			}
			if v57 == int32(1) {
				v63 = int32(4)
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				if v65&int32(254) == int32(2) {
					v74 = v63
				} else {
					v74 = base.B2i32(v65 == int32(18)) << (uint(v63) % 32)
				}
				if v65 == int32(1) {
					v77 = v63
				} else {
					v77 = v74
				}
				v88 = v77
			} else {
				v78 = int32(1)
				if v59 != 0 {
					v88 = int32(base.Ui32(v57)>>(uint(v78)%32)) - v78
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v90 = F_calc_word_similarity(m, v52, v51, v60, v88, int32(3))
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v92 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v96 != v19 {
							F_pfree(m, v19)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v101 = *(*float64)(unsafe.Add(mBase, _consts[1561]))
								return base.F64_le(v101, base.F64_promote_f32(v90))
							}
						} else {
							v101 = *(*float64)(unsafe.Add(mBase, _consts[1561]))
							return base.F64_le(v101, base.F64_promote_f32(v90))
						}
					}
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v96 != v19 {
						F_pfree(m, v19)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v101 = *(*float64)(unsafe.Add(mBase, _consts[1561]))
							return base.F64_le(v101, base.F64_promote_f32(v90))
						}
					} else {
						v101 = *(*float64)(unsafe.Add(mBase, _consts[1561]))
						return base.F64_le(v101, base.F64_promote_f32(v90))
					}
				}
			}
		}
	}
}
func F_strip_implicit_coercions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v40
L2:
	;
	v2 = l0
	goto L5
L3:
	;
	goto L4
L4:
	;
	v40 = int32(0)
	goto L1
L5:
	;
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	switch v3 - int32(15) {
	case 0:
		goto L13
	default:
		v40 = v2
		goto L1
	case 12:
		goto L12
	case 13:
		goto L11
	case 14:
		goto L10
	case 15:
		goto L9
	case 40:
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v37 != 0 {
		v2 = v37
		goto L5
	} else {
		goto L20
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	if v31 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v2)+12))
	if v26 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L18
	}
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v2)+24))
	if v21 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	if v16 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v2)+20))
	if v11 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L15
	}
L13:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	if v6 != int32(2) {
		v40 = v2
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v2)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v36 = v10
	goto L7
L15:
	;
	v36 = v2 + int32(4)
	goto L7
L16:
	;
	v36 = v2 + int32(4)
	goto L7
L17:
	;
	v36 = v2 + int32(4)
	goto L7
L18:
	;
	v36 = v2 + int32(4)
	goto L7
L19:
	;
	v36 = v2 + int32(4)
	goto L7
L20:
	;
	goto L6
}
func F_strlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	if l0&int32(3) == int32(0) {
		v26 = l0
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v51 - l0
L2:
	;
	v30 = v26
	goto L11
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v15 = l0
	goto L7
L7:
	;
	v19 = v15 + int32(1)
	if v19&int32(3) == int32(0) {
		v26 = v19
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v51 = v19
	goto L1
L9:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v24 != 0 {
		v15 = v19
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v39 = int32(-2139062144)
	if (int32(16843008)-v36|v36)&v39 == v39 {
		v30 = v30 + int32(4)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v45 = v30
	goto L14
L13:
	;
	goto L12
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 != 0 {
		v45 = v45 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v51 = v45
	goto L1
L16:
	;
	goto L15
}
func F_strpbrk(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v2 = int32(532416)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*int8)(unsafe.Add(mBase, _consts[1638])))
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v69 = v63 - l0 + l0
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v71 != 0 {
		goto L21
	} else {
		goto L22
	}
L2:
	;
	m.G0 = v8 + int32(32)
	goto L1
L3:
	;
	v15 = F___memset(m, v8, int32(0), int32(32))
	mBase = m.M
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1638])))
	if v16 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1639])))
	if v11 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v12 = F___strchrnul(m, l0, v10)
	mBase = m.M
	v63 = v12
	goto L2
L7:
	;
	goto L6
L8:
	;
	v18 = v2
	v19 = v16
	goto L11
L9:
	;
	goto L10
L10:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v40 == int32(0) {
		v63 = l0
		goto L2
	} else {
		goto L14
	}
L11:
	;
	v26 = v8 + int32(base.Ui32(v19)>>(uint(int32(3))%32))&int32(28)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v28 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v27 | v28<<(uint(v19)%32)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v32 != 0 {
		v18 = v18 + v28
		v19 = v32
		goto L11
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	goto L12
L14:
	;
	v44 = l0
	v45 = v40
	goto L15
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(base.Ui32(v45)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v53)>>(uint(v45)%32))&int32(1) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v63 = v59
	goto L2
L17:
	;
	v63 = v44
	goto L2
L18:
	;
	goto L19
L19:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v59 = v44 + int32(1)
	if v57 != 0 {
		v44 = v59
		v45 = v57
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	v72 = v69
	goto L23
L22:
	;
	v72 = int32(0)
	goto L23
L23:
	;
	return v72
}
func F_strtod(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v42 int64
	_ = v42
	var v47 int64
	_ = v47
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v95 int64
	_ = v95
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v117 int64
	_ = v117
	var v127 int64
	_ = v127
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v141 int64
	_ = v141
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_strtox_1(m, v7, l0, l1, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return float64(0)
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		v23 = m.G0
		v25 = v23 - int32(32)
		m.G0 = v25
		v28 = v15 & int64(281474976710655)
		v32 = int64(base.Ui64(v15)>>(uint(int64(48))%64)) & int64(32767)
		v33 = base.I32_wrap_i64(v32)
		if base.Ui32(v33-int32(15361)) <= base.Ui32(int32(2045)) {
			v42 = v28<<(uint(int64(4))%64) | int64(base.Ui64(v14)>>(uint(int64(60))%64))
			v47 = v14 & int64(1152921504606846975)
			if base.Ui64(int64(576460752303423489)) <= base.Ui64(v47) {
				v57 = v42 + int64(1)
			} else {
				if v47 != int64(576460752303423488) {
					v57 = v42
				} else {
					v57 = v42&int64(1) + v42
				}
			}
			v60 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v57))
			if base.Ui64(int64(4503599627370495)) < base.Ui64(v57) {
				v61 = int64(0)
			} else {
				v61 = v57
			}
			v134 = v61
			v141 = base.I64_extend_i32_u(v60) + base.I64_extend_i32_u(v33-int32(15360))
		} else {
			if v14|v28 == int64(0) {
				if base.Ui32(int32(17406)) < base.Ui32(v33) {
					v134 = int64(0)
					v141 = int64(2047)
				} else {
					v84 = base.B2i32(v32 == int64(0))
					if v32 == int64(0) {
						v85 = int32(15360)
					} else {
						v85 = int32(15361)
					}
					v86 = v85 - v33
					if int32(112) < v86 {
						v89 = int64(0)
						v134 = v89
						v141 = v89
					} else {
						if v32 == int64(0) {
							v95 = v28
						} else {
							v95 = v28 | int64(281474976710656)
						}
						F___ashlti3(m, v25+int32(16), v14, v95, int32(128)-v86)
						mBase = m.M
						F___lshrti3(m, v25, v14, v95, v86)
						mBase = m.M
						v100 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
						v103 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
						v106 = v100<<(uint(int64(4))%64) | int64(base.Ui64(v103)>>(uint(int64(60))%64))
						v108 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
						v117 = base.I64_extend_i32_u(base.B2i32(v33 != v85)&base.B2i32(v108|v109 != int64(0))) | v103&int64(1152921504606846975)
						if base.Ui64(int64(576460752303423489)) <= base.Ui64(v117) {
							v127 = v106 + int64(1)
						} else {
							if v117 != int64(576460752303423488) {
								v127 = v106
							} else {
								v127 = v106&int64(1) + v106
							}
						}
						v131 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v127))
						if base.Ui64(int64(4503599627370495)) < base.Ui64(v127) {
							v132 = v127 ^ int64(4503599627370496)
						} else {
							v132 = v127
						}
						v134 = v132
						v141 = base.I64_extend_i32_u(v131)
					}
				}
			} else {
				if v32 != int64(32767) {
					if base.Ui32(int32(17406)) < base.Ui32(v33) {
						v134 = int64(0)
						v141 = int64(2047)
					} else {
						v84 = base.B2i32(v32 == int64(0))
						if v32 == int64(0) {
							v85 = int32(15360)
						} else {
							v85 = int32(15361)
						}
						v86 = v85 - v33
						if int32(112) < v86 {
							v89 = int64(0)
							v134 = v89
							v141 = v89
						} else {
							if v32 == int64(0) {
								v95 = v28
							} else {
								v95 = v28 | int64(281474976710656)
							}
							F___ashlti3(m, v25+int32(16), v14, v95, int32(128)-v86)
							mBase = m.M
							F___lshrti3(m, v25, v14, v95, v86)
							mBase = m.M
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
							v103 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
							v106 = v100<<(uint(int64(4))%64) | int64(base.Ui64(v103)>>(uint(int64(60))%64))
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
							v117 = base.I64_extend_i32_u(base.B2i32(v33 != v85)&base.B2i32(v108|v109 != int64(0))) | v103&int64(1152921504606846975)
							if base.Ui64(int64(576460752303423489)) <= base.Ui64(v117) {
								v127 = v106 + int64(1)
							} else {
								if v117 != int64(576460752303423488) {
									v127 = v106
								} else {
									v127 = v106&int64(1) + v106
								}
							}
							v131 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v127))
							if base.Ui64(int64(4503599627370495)) < base.Ui64(v127) {
								v132 = v127 ^ int64(4503599627370496)
							} else {
								v132 = v127
							}
							v134 = v132
							v141 = base.I64_extend_i32_u(v131)
						}
					}
				} else {
					v134 = v28<<(uint(int64(4))%64) | int64(base.Ui64(v14)>>(uint(int64(60))%64)) | int64(2251799813685248)
					v141 = int64(2047)
				}
			}
		}
		m.G0 = v25 + int32(32)
		m.G0 = v7 + int32(16)
		return base.F64_reinterpret_i64(v15&int64(-9223372036854775807-1) | v141<<(uint(int64(52))%64) | v134)
	}
}
func F_strtok_r(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	if l0 != 0 {
		v7 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = int32(779145)
	v12 = m.G0
	v14 = v12 - int32(32)
	v15 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v15
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1640])))
	if v23 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v4 != 0 {
		v7 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	return int32(0)
L4:
	;
	v92 = v91 + v7
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v93 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L5:
	;
	v91 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1641])))
	if v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v31 = v7
	goto L11
L9:
	;
	goto L10
L10:
	;
	v41 = v8
	v42 = v23
	goto L14
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v37 == v23 {
		v31 = v31 + int32(1)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v91 = v31 - v7
	goto L4
L13:
	;
	goto L12
L14:
	;
	v49 = v14 + int32(base.Ui32(v42)>>(uint(int32(3))%32))&int32(28)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v51 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v50 | v51<<(uint(v42)%32)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v55 != 0 {
		v41 = v41 + v51
		v42 = v55
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v58 == int32(0) {
		v83 = v7
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v91 = v83 - v7
	goto L4
L18:
	;
	v62 = v7
	v63 = v58
	goto L19
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(base.Ui32(v63)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v71)>>(uint(v63)%32))&int32(1) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v83 = v79
	goto L17
L21:
	;
	v83 = v62
	goto L17
L22:
	;
	goto L23
L23:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v79 = v62 + int32(1)
	if v77 != 0 {
		v62 = v79
		v63 = v77
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	v96 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v96
	return v96
L26:
	;
	goto L27
L27:
	;
	v100 = int32(779145)
	v104 = m.G0
	v106 = v104 - int32(32)
	m.G0 = v106
	v108 = int32(*(*int8)(unsafe.Add(mBase, _consts[1640])))
	if v108 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v167 = v161 - v92 + v92
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v168 != 0 {
		goto L48
	} else {
		goto L49
	}
L29:
	;
	m.G0 = v106 + int32(32)
	goto L28
L30:
	;
	v113 = F___memset(m, v106, int32(0), int32(32))
	mBase = m.M
	v114 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1640])))
	if v114 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1641])))
	if v109 != 0 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v110 = F___strchrnul(m, v92, v108)
	mBase = m.M
	v161 = v110
	goto L29
L34:
	;
	goto L33
L35:
	;
	v116 = v100
	v117 = v114
	goto L38
L36:
	;
	goto L37
L37:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v138 == int32(0) {
		v161 = v92
		goto L29
	} else {
		goto L41
	}
L38:
	;
	v124 = v106 + int32(base.Ui32(v117)>>(uint(int32(3))%32))&int32(28)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v126 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v125 | v126<<(uint(v117)%32)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v130 != 0 {
		v116 = v116 + v126
		v117 = v130
		goto L38
	} else {
		goto L40
	}
L39:
	;
	goto L37
L40:
	;
	goto L39
L41:
	;
	v142 = v92
	v143 = v138
	goto L42
L42:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v106+int32(base.Ui32(v143)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v151)>>(uint(v143)%32))&int32(1) != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v161 = v157
	goto L29
L44:
	;
	v161 = v142
	goto L29
L45:
	;
	goto L46
L46:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)))
	v157 = v142 + int32(1)
	if v155 != 0 {
		v142 = v157
		v143 = v155
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v167 + int32(1)
	v172 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v172)
	return v92
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	return v92
}
func F_strtoll(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v5 int64
	_ = v5
	v5 = F_strtox_2(m, l0, l1, l2, int64(-9223372036854775807-1))
	return v5
}
func F_strtox_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v8 = m.G0
	v10 = v8 - int32(160)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(-1)
	v17 = v10 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = int64(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = base.I64_extend_i32_s(v22 - v23)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+104)) = v27
	F___floatscan(m, v10, v10+int32(16), l3, int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return
	} else {
		v42 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
		v43 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		if l2 != 0 {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+136))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v44 + (l1 + (v45 - v46))
		} else {
		}
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v42
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v43
		m.G0 = v10 + int32(160)
		return
	}
}
func F_subltree(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_inner_subltree(m, v5, v9, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v13 != v5 {
				F_pfree(m, v5)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					return v11
				}
			} else {
				return v11
			}
		}
	}
}
func F_substitute_grouped_columns_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
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
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v585 int32
	_ = v585
	var v594 int32
	_ = v594
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v670 int32
	_ = v670
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	v3 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v22 + int32(32)
	return v867
L2:
	;
	v867 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v27 - int32(9) {
	case 0:
		goto L8
	case 1:
		goto L7
	default:
		goto L6
	}
L5:
	;
	v857 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L20
	} else {
		goto L187
	}
L6:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v39 != int32(1) {
		v167 = v27
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v35 <= v34 {
		v867 = l0
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v30 == v31 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	if v31 < v30 {
		v867 = l0
		goto L1
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	goto L6
L12:
	;
	switch v167 - int32(6) {
	case 0:
		goto L43
	case 1, 2:
		v867 = l0
		goto L1
	default:
		goto L42
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v42 != 0 {
		v167 = v27
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v43 == int32(0) {
		v167 = v27
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v49 = int32(0)
	goto L17
L16:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v167 = v164
	goto L12
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v66 <= v49 {
		goto L16
	} else {
		goto L19
	}
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+56))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v88 = v85 + v71<<(uint(int32(5))%32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(32))))
	v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88-int32(28)))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(24))))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(20))))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(16))))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v105 = F_makeVar(m, v91, v94, v97, v100, v103, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L20
	} else {
		goto L23
	}
L19:
	;
	v71 = v49 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v49<<(uint(int32(2))%32)+v72)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v76 = F_equal(m, l0, v75)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	if v76 == int32(0) {
		v49 = v71
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+36)) = v109
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88-int32(4)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v105)+40)) = uint16(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+108))
	if v116 == int32(0) {
		v867 = v105
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v120 = int32(0)
	if v119 == v120 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v158 != 0 {
		v867 = v105
		goto L1
	} else {
		goto L38
	}
L26:
	;
	v158 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v126 <= int32(0) {
		v151 = v120
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v158 = v151
	goto L25
L30:
	;
	v129 = int32(0)
	if v129 < v126 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v132 = v126
	goto L33
L32:
	;
	v132 = v129
	goto L33
L33:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v135 = int32(0)
	goto L34
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133+v135<<(uint(int32(2))%32))))
	v144 = base.B2i32(v143 == v82)
	if v143 == v82 {
		v151 = v144
		goto L29
	} else {
		goto L36
	}
L35:
	;
	v151 = v144
	goto L29
L36:
	;
	v146 = v135 + int32(1)
	if v146 != v132 {
		v135 = v146
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	v161 = F_bms_add_member(m, v159, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v161
	v867 = v105
	goto L1
L40:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v845 + int32(1)
	v851 = F_query_tree_mutator_impl(m, l0, int32(481), l1, int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L20
	} else {
		goto L186
	}
L41:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v230)+16))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v764)+56))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v765)+16))
	v769 = v766 + v226<<(uint(int32(5))%32)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v769-int32(32))))
	v775 = int32(*(*int16)(unsafe.Add(mBase, uint32(v769-int32(28)))))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v769-int32(24))))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v769-int32(20))))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v769-int32(16))))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v786 = F_makeVar(m, v772, v775, v778, v781, v784, v785)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L20
	} else {
		goto L168
	}
L42:
	;
	if v167 == int32(67) {
		goto L40
	} else {
		goto L166
	}
L43:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v186 != v187 {
		v867 = l0
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v186 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v267 = int32(0)
	if v265 == v267 {
		goto L64
	} else {
		goto L65
	}
L46:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v191&int32(1) != 0 {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v194 == int32(0) {
		goto L45
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	if v197 <= int32(0) {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	v200 = int32(0)
	if v200 < v197 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v204 = v197
	goto L54
L53:
	;
	v204 = v200
	goto L54
L54:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	v208 = v200
	goto L55
L55:
	;
	v226 = v208 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v205+v208<<(uint(int32(2))%32))))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	if v232 != int32(6) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L45
L57:
	;
	if v226 != v204 {
		v208 = v226
		goto L55
	} else {
		goto L62
	}
L58:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v235 != v236 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231)+8)))
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v238 != v239 {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v231)+28))
	if v241 == int32(0) {
		goto L41
	} else {
		goto L61
	}
L61:
	;
	goto L57
L62:
	;
	goto L56
L63:
	;
	if v305 != 0 {
		v867 = l0
		goto L1
	} else {
		goto L76
	}
L64:
	;
	v305 = int32(0)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	if v273 <= int32(0) {
		v298 = v267
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v305 = v298
	goto L63
L68:
	;
	v276 = int32(0)
	if v276 < v273 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v279 = v273
	goto L71
L70:
	;
	v279 = v276
	goto L71
L71:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	v282 = int32(0)
	goto L72
L72:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v280+v282<<(uint(int32(2))%32))))
	v291 = base.B2i32(v290 == v266)
	if v290 == v266 {
		v298 = v291
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v298 = v291
	goto L67
L74:
	;
	v293 = v282 + int32(1)
	if v293 != v279 {
		v282 = v293
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+8))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+12))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v308+v309<<(uint(int32(2))%32)-int32(4))))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	if v316 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v706 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	v707 = F_get_rte_attribute_name(m, v315, v706)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L20
	} else {
		goto L150
	}
L78:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v315)+16))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v321 = v319 + int32(148)
	v322 = int32(0)
	v323 = m.G0
	v325 = v323 - int32(16)
	m.G0 = v325
	v327 = m.G0
	v329 = v327 + int32(-64)
	m.G0 = v329
	v332 = v325 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v332))) = v322
	v337 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L20
	} else {
		goto L81
	}
L79:
	;
	m.G0 = v325 + int32(16)
	if v670 == int32(0) {
		goto L77
	} else {
		goto L148
	}
L80:
	;
	if v475 == int32(0) {
		v670 = v3
		goto L79
	} else {
		goto L120
	}
L81:
	;
	F_ScanKeyInit(m, v327+int32(-48), int32(9), int32(3), int32(184), v317)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L20
	} else {
		goto L82
	}
L82:
	;
	v347 = int32(1)
	v352 = F_systable_beginscan(m, v337, int32(2665), v347, int32(0), v347, v327+int32(-48))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L20
	} else {
		goto L86
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L20
	} else {
		goto L117
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L20
	} else {
		goto L114
	}
L85:
	;
	F_systable_endscan(m, v352)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L20
	} else {
		goto L112
	}
L86:
	;
	v354 = F_systable_getnext(m, v352)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	if v354 == int32(0) {
		v475 = v322
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v361 = v354
	goto L89
L89:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v361)+16))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+22)))
	v379 = v377 + v378
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+72)))
	if v380 == int32(112) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v475 = v322
	goto L85
L91:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+73)))
	if v383&int32(1) != 0 {
		v475 = v322
		goto L85
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v467 = F_systable_getnext(m, v352)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L20
	} else {
		goto L110
	}
L94:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v337)+52))
	v389 = F_heap_getattr_3(m, v361, v386, v327+int32(-49))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L20
	} else {
		goto L95
	}
L95:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+15)))
	if v391 == int32(1) {
		goto L84
	} else {
		goto L96
	}
L96:
	;
	v394 = F_pg_detoast_datum(m, v389)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L20
	} else {
		goto L97
	}
L97:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v396 != int32(1) {
		goto L83
	} else {
		goto L98
	}
L98:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v394)+16))
	if v399 < int32(0) {
		goto L83
	} else {
		goto L99
	}
L99:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v394)+8))
	if v402 != 0 {
		goto L83
	} else {
		goto L100
	}
L100:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v394)+12))
	if v403 != int32(21) {
		goto L83
	} else {
		goto L101
	}
L101:
	;
	v406 = int32(0)
	if v399 == v406 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v361)+16))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+22)))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v462+v463)))
	*(*int32)(unsafe.Add(mBase, uint32(v332))) = v465
	v475 = v449
	goto L85
L103:
	;
	v449 = int32(0)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v418 = v406
	v419 = int32(0)
	goto L106
L106:
	;
	v435 = int32(*(*int16)(unsafe.Add(mBase, uint32(v394+int32(24)+v418<<(uint(int32(1))%32)))))
	v438 = F_bms_add_member(m, v419, v435+int32(7))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L20
	} else {
		goto L108
	}
L107:
	;
	v449 = v438
	goto L102
L108:
	;
	v441 = v418 + int32(1)
	if v441 != v399 {
		v418 = v441
		v419 = v438
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	if v467 != 0 {
		v361 = v467
		goto L89
	} else {
		goto L111
	}
L111:
	;
	goto L90
L112:
	;
	F_sequence_close(m, v337, int32(1))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L20
	} else {
		goto L113
	}
L113:
	;
	m.G0 = v329 - int32(-64)
	goto L80
L114:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v361)+16))
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+22)))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v500+v501)))
	*(*int32)(unsafe.Add(mBase, uint32(v329))) = v503
	F_errmsg_internal(m, int32(43656), v329)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L20
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(515494), int32(1499), int32(144929))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L20
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errmsg_internal(m, int32(25335), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L20
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(515494), int32(1506), int32(144929))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L20
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	if v318 == int32(0) {
		v585 = v3
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v594 = int32(0)
	if v475 == v594 {
		goto L133
	} else {
		goto L134
	}
L122:
	;
	v531 = int32(0)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	if v532 <= v531 {
		v585 = v3
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v538 = v531
	v545 = v3
	goto L124
L124:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v318)+12))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v554+v538<<(uint(int32(2))%32))))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	if v559 != int32(6) {
		v570 = v545
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v585 = v570
	goto L121
L126:
	;
	v572 = v538 + int32(1)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	if v572 < v573 {
		v538 = v572
		v545 = v570
		goto L124
	} else {
		goto L131
	}
L127:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v558)+4))
	if v562 != v309 {
		v570 = v545
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v558)+28))
	if v564 != 0 {
		v570 = v545
		goto L126
	} else {
		goto L129
	}
L129:
	;
	v565 = int32(*(*int16)(unsafe.Add(mBase, uint32(v558)+8)))
	v568 = F_bms_add_member(m, v545, v565+int32(7))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L20
	} else {
		goto L130
	}
L130:
	;
	v570 = v568
	goto L126
L131:
	;
	goto L125
L132:
	;
	if v647 == int32(0) {
		v670 = v3
		goto L79
	} else {
		goto L146
	}
L133:
	;
	v647 = int32(1)
	goto L132
L134:
	;
	goto L135
L135:
	;
	if v585 == int32(0) {
		v638 = v594
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v647 = v638
	goto L132
L137:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v585)+4))
	if v604 < v603 {
		v638 = v594
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v606 = int32(1)
	if v603 <= v606 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v609 = v606
	goto L141
L140:
	;
	v609 = v603
	goto L141
L141:
	;
	v610 = int32(8)
	v615 = int32(0)
	goto L142
L142:
	;
	v622 = v615 << (uint(int32(2)) % 32)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v475+v610+v622)))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v622+(v585+v610))))
	v629 = v624 & (v626 ^ int32(-1))
	v631 = base.B2i32(v629 == int32(0))
	if v629 != 0 {
		v638 = v631
		goto L136
	} else {
		goto L144
	}
L143:
	;
	v638 = v631
	goto L136
L144:
	;
	v633 = v615 + int32(1)
	if v633 != v609 {
		v615 = v633
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v325)+12))
	v652 = F_lappend_oid(m, v650, v651)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L20
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = v652
	v670 = int32(1)
	goto L79
L148:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v680)))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v683 = F_lappend_int(m, v681, v682)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L20
	} else {
		goto L149
	}
L149:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v685))) = v683
	v867 = l0
	goto L1
L150:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L20
	} else {
		goto L151
	}
L151:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L20
	} else {
		goto L152
	}
L152:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+4))
	if v709 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v718
	F_errmsg(m, int32(265841), v22)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L20
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v707
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v718
	F_errmsg(m, int32(16525), v22+int32(16))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L20
	} else {
		goto L163
	}
L156:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)))
	if v726 == int32(1) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	F_errdetail(m, int32(619631), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L20
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_parser_errposition(m, v733, v734)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L20
	} else {
		goto L161
	}
L160:
	;
	goto L159
L161:
	;
	F_errfinish(m, int32(521798), int32(1556), int32(219580))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L20
	} else {
		goto L162
	}
L162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L163:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_parser_errposition(m, v749, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L20
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(521798), int32(1562), int32(219580))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L20
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	v761 = F_expression_tree_mutator_impl(m, l0, int32(481), l1)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L20
	} else {
		goto L167
	}
L167:
	;
	v867 = v761
	goto L1
L168:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v769-int32(8))))
	*(*int32)(unsafe.Add(mBase, uint32(v786)+36)) = v790
	v794 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v769-int32(4)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v786)+40)) = uint16(v794)
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v796)+108))
	if v797 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v867 = v786
	goto L1
L170:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v801 = int32(0)
	if v800 == v801 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if v839 != 0 {
		goto L169
	} else {
		goto L184
	}
L172:
	;
	v839 = int32(0)
	goto L171
L173:
	;
	goto L174
L174:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v800)+4))
	if v807 <= int32(0) {
		v832 = v801
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v839 = v832
	goto L171
L176:
	;
	v810 = int32(0)
	if v810 < v807 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v813 = v807
	goto L179
L178:
	;
	v813 = v810
	goto L179
L179:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v800)+12))
	v816 = int32(0)
	goto L180
L180:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v814+v816<<(uint(int32(2))%32))))
	v825 = base.B2i32(v824 == v763)
	if v824 == v763 {
		v832 = v825
		goto L175
	} else {
		goto L182
	}
L181:
	;
	v832 = v825
	goto L175
L182:
	;
	v827 = v816 + int32(1)
	if v827 != v813 {
		v816 = v827
		goto L180
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v786)+24))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v765)+8))
	v842 = F_bms_add_member(m, v840, v841)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L20
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v786)+24)) = v842
	goto L169
L186:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v853 - int32(1)
	v867 = v851
	goto L1
L187:
	;
	v859 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v859)
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v857)+28))
	v862 = F_substitute_grouped_columns_mutator(m, v861, l1)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L20
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v857)+28)) = v862
	v865 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v865)
	v867 = v857
	goto L1
}
func F_switchToPresortedPrefixMode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
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
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v186 int64
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
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
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int64
	_ = v237
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int64
	_ = v248
	var v249 int64
	_ = v249
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v268 int64
	_ = v268
	var v272 int32
	_ = v272
	var v282 int32
	_ = v282
	var v285 int64
	_ = v285
	var v289 int64
	_ = v289
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int64
	_ = v306
	var v307 int64
	_ = v307
	var v310 int64
	_ = v310
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v317 int64
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int64
	_ = v326
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int64
	_ = v337
	var v338 int64
	_ = v338
	var v340 int32
	_ = v340
	var v341 int64
	_ = v341
	var v343 int64
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v357 int64
	_ = v357
	var v361 int32
	_ = v361
	var v371 int32
	_ = v371
	var v374 int64
	_ = v374
	var v378 int64
	_ = v378
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int64
	_ = v395
	var v396 int64
	_ = v396
	var v399 int64
	_ = v399
	var v402 int64
	_ = v402
	var v403 int64
	_ = v403
	var v406 int64
	_ = v406
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v420 int64
	_ = v420
	var v421 int64
	_ = v421
	var v422 int64
	_ = v422
	var v424 int64
	_ = v424
	v7 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+56))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v18 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v50 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v28 = int32(2)
	v29 = v22 << (uint(v28) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+84))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
	v37 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	v38 = int32(0)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v41 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_tuplesort_reset(m, v18)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L10
	}
L5:
	;
	v42 = v28
	goto L7
L6:
	;
	v42 = v38
	goto L7
L7:
	;
	v43 = F_tuplesort_begin_heap(m, v17, v21-v22, v24+v22<<(uint(int32(1))%32), v29+v30, v32+v29, v34+v22, v37, v38, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v43
	goto L1
L10:
	;
	goto L1
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v56 = v54 - v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+236))
	if v59 != 0 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	goto L13
L13:
	;
	v85 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v85 <= int64(0) {
		v199 = v85
		v200 = v7
		goto L26
	} else {
		goto L27
	}
L14:
	;
	goto L13
L15:
	;
	goto L14
L16:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v53)+72)) = uint32(v56)
	v68 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+68)) = uint8(v68)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+32))
	if v74 != 0 {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	if int64(1073741823) < v56 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if int64(1073741823) < v56 {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53)+232))
	if v62 != int32(-1) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L15
L22:
	;
	goto L16
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
	v77 = v76
	goto L25
L24:
	;
	v77 = v73
	goto L25
L25:
	;
	v78 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+28)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+32)) = v78
	goto L15
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v199 - v200
	if v199 == v200 {
		goto L61
	} else {
		goto L62
	}
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	if v88 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	m.T0[v189].(func(*base.Module, int32))(m, v187)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L8
	} else {
		goto L59
	}
L29:
	;
	v133 = int64(1)
	v134 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v134 < int64(2) {
		v199 = v134
		v200 = v133
		goto L26
	} else {
		goto L45
	}
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v106 = int32(0)
	v108 = F_tuplesort_gettupleslot(m, v103, base.B2i32(v15 == int32(1)), v106, v88, v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L35
	}
L31:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+4)))
	if v91&int32(2) != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_tuplesort_puttupleslot(m, v94, v88)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+32))
	m.T0[v100].(func(*base.Module, int32, int32))(m, v97, v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v110 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v124 = F_isCurrentGroup(m, l0, v122, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L8
	} else {
		goto L42
	}
L37:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+4)))
	if v111&int32(2) == int32(0) {
		v122 = v110
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+32))
	m.T0[v118].(func(*base.Module, int32, int32))(m, v110, v116)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v122 = v121
	goto L36
L42:
	;
	if v124 == int32(0) {
		v186 = v7
		goto L28
	} else {
		goto L43
	}
L43:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	F_tuplesort_puttupleslot(m, v128, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	goto L29
L45:
	;
	v146 = v133
	goto L46
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v148 = int32(0)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v151 = F_tuplesort_gettupleslot(m, v147, base.B2i32(v15 == int32(1)), v148, v149, v148)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L8
	} else {
		goto L48
	}
L47:
	;
	v199 = v177
	v200 = v176
	goto L26
L48:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v153 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v167 = F_isCurrentGroup(m, l0, v165, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L8
	} else {
		goto L55
	}
L50:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+4)))
	if v154&int32(2) == int32(0) {
		v165 = v153
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+32))
	m.T0[v161].(func(*base.Module, int32, int32))(m, v153, v159)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v165 = v164
	goto L49
L55:
	;
	if v167 == int32(0) {
		v186 = v146
		goto L28
	} else {
		goto L56
	}
L56:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	F_tuplesort_puttupleslot(m, v171, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v176 = v146 + int64(1)
	v177 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	if v176 < v177 {
		v146 = v176
		goto L46
	} else {
		goto L58
	}
L58:
	;
	goto L47
L59:
	;
	v192 = *(*int64)(unsafe.Add(mBase, uint32(l0)+152))
	v199 = v192
	v200 = v186
	goto L26
L60:
	;
	m.G0 = v11 + int32(16)
	return
L61:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+32))
	m.T0[v207].(func(*base.Module, int32, int32))(m, v204, v205)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L8
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_tuplesort_performsort(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L8
	} else {
		goto L66
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(1)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	m.T0[v214].(func(*base.Module, int32))(m, v212)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L8
	} else {
		goto L65
	}
L65:
	;
	goto L60
L66:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v220 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v417 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L68:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v223 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v326 = *(*int64)(unsafe.Add(mBase, uint32(l0)+224))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+224)) = v326 + int64(1)
	v330 = int32(0)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v325)+128))
	if v334 == v330 {
		goto L100
	} else {
		goto L101
	}
L70:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+280)))
	if v226 != int32(1) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v231 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	v236 = v223 + v231*int32(96) + int32(56)
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v236)))
	*(*int64)(unsafe.Add(mBase, uint32(v236))) = v237 + int64(1)
	v241 = int32(0)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v229)+128))
	if v245 == v241 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v305 {
	case 0:
		goto L94
	case 1:
		goto L93
	default:
		goto L92
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v282
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v229)+112))
	v289 = base.I64_div_s(v285+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v289
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v229)+124))
	switch v291 - int32(3) {
	case 0:
		goto L88
	case 1:
		v302 = v291
		goto L85
	case 2:
		goto L87
	default:
		goto L86
	}
L74:
	;
	if v261&int32(255) != base.B2i32(v245 != int32(0)) {
		goto L80
	} else {
		goto L81
	}
L75:
	;
	v248 = *(*int64)(unsafe.Add(mBase, uint32(v229)+96))
	v249 = *(*int64)(unsafe.Add(mBase, uint32(v229)+88))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+120)))
	v261 = v251
	v262 = v248 - v249
	goto L74
L76:
	;
	goto L77
L77:
	;
	v252 = F_LogicalTapeSetBlocks(m, v245)
	mBase = m.M
	v254 = v252 << (uint(int64(13)) % 64)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+120)))
	if v255 != 0 {
		v261 = v255
		v262 = v254
		goto L74
	} else {
		goto L78
	}
L78:
	;
	v256 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v229)+120)) = uint8(v256)
	*(*int64)(unsafe.Add(mBase, uint32(v229)+112)) = v254
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v229)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+124)) = v259
	v282 = v241
	goto L73
L79:
	;
	v282 = int32(1)
	goto L73
L80:
	;
	if v261&int32(1) != 0 {
		v282 = v241
		goto L73
	} else {
		goto L84
	}
L81:
	;
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v229)+112))
	if v262 <= v268 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v229)+120)) = uint8(v261)
	*(*int64)(unsafe.Add(mBase, uint32(v229)+112)) = v262
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v229)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+124)) = v272
	if v261&int32(1) == int32(0) {
		goto L79
	} else {
		goto L83
	}
L83:
	;
	v282 = v241
	goto L73
L84:
	;
	goto L79
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v302
	goto L72
L86:
	;
	v302 = int32(0)
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(8)
	goto L72
L88:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+69)))
	if v296 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v297 = int32(1)
	goto L91
L90:
	;
	v297 = int32(2)
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v297
	goto L72
L92:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v236)+40))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+40)) = v321 | v322
	goto L67
L93:
	;
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v236)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v236)+32)) = v313 + v314
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v236)+24))
	if v313 <= v317 {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v236)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v236)+16)) = v306 + v307
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v236)+8))
	if v306 <= v310 {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v306
	goto L92
L96:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v236)+24)) = v313
	goto L92
L97:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v394 {
	case 0:
		goto L119
	case 1:
		goto L118
	default:
		goto L117
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v371
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v325)+112))
	v378 = base.I64_div_s(v374+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v378
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v325)+124))
	switch v380 - int32(3) {
	case 0:
		goto L113
	case 1:
		v391 = v380
		goto L110
	case 2:
		goto L112
	default:
		goto L111
	}
L99:
	;
	if v350&int32(255) != base.B2i32(v334 != int32(0)) {
		goto L105
	} else {
		goto L106
	}
L100:
	;
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v325)+96))
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v325)+88))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+120)))
	v350 = v340
	v351 = v337 - v338
	goto L99
L101:
	;
	goto L102
L102:
	;
	v341 = F_LogicalTapeSetBlocks(m, v334)
	mBase = m.M
	v343 = v341 << (uint(int64(13)) % 64)
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+120)))
	if v344 != 0 {
		v350 = v344
		v351 = v343
		goto L99
	} else {
		goto L103
	}
L103:
	;
	v345 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v325)+120)) = uint8(v345)
	*(*int64)(unsafe.Add(mBase, uint32(v325)+112)) = v343
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v325)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v325)+124)) = v348
	v371 = v330
	goto L98
L104:
	;
	v371 = int32(1)
	goto L98
L105:
	;
	if v350&int32(1) != 0 {
		v371 = v330
		goto L98
	} else {
		goto L109
	}
L106:
	;
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v325)+112))
	if v351 <= v357 {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v325)+120)) = uint8(v350)
	*(*int64)(unsafe.Add(mBase, uint32(v325)+112)) = v351
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v325)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v325)+124)) = v361
	if v350&int32(1) == int32(0) {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v371 = v330
	goto L98
L109:
	;
	goto L104
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v391
	goto L97
L111:
	;
	v391 = int32(0)
	goto L110
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(8)
	goto L97
L113:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+69)))
	if v385 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v386 = int32(1)
	goto L116
L115:
	;
	v386 = int32(2)
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v386
	goto L97
L117:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v410 | v411
	goto L67
L118:
	;
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v403 = *(*int64)(unsafe.Add(mBase, uint32(l0)+256))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+256)) = v402 + v403
	v406 = *(*int64)(unsafe.Add(mBase, uint32(l0)+248))
	if v402 <= v406 {
		goto L117
	} else {
		goto L121
	}
L119:
	;
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	v396 = *(*int64)(unsafe.Add(mBase, uint32(l0)+240))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+240)) = v395 + v396
	v399 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	if v395 <= v399 {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v395
	goto L117
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+248)) = v402
	goto L117
L122:
	;
	v420 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v421 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	v422 = v421 + v200
	if v420 < v422 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(3)
	goto L60
L125:
	;
	v424 = v420
	goto L127
L126:
	;
	v424 = v422
	goto L127
L127:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = v424
	goto L124
}
func F_synchronize_slots(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
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
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
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
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int64
	_ = v374
	var v377 int64
	_ = v377
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v755 int32
	_ = v755
	var v763 int32
	_ = v763
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v795 int64
	_ = v795
	var v796 int32
	_ = v796
	var v799 int64
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v807 int64
	_ = v807
	var v809 int64
	_ = v809
	var v810 int64
	_ = v810
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int64
	_ = v828
	var v829 int32
	_ = v829
	var v831 int64
	_ = v831
	var v832 int64
	_ = v832
	var v837 int64
	_ = v837
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int64
	_ = v896
	var v897 int64
	_ = v897
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v937 int32
	_ = v937
	var v939 int64
	_ = v939
	var v941 int64
	_ = v941
	var v943 int64
	_ = v943
	var v945 int64
	_ = v945
	var v947 int64
	_ = v947
	var v949 int64
	_ = v949
	var v951 int64
	_ = v951
	var v953 int64
	_ = v953
	var v955 int32
	_ = v955
	var v957 int64
	_ = v957
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int64
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v987 int64
	_ = v987
	var v988 int32
	_ = v988
	var v992 int64
	_ = v992
	var v995 int64
	_ = v995
	var v1000 int32
	_ = v1000
	var v1004 int64
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1009 int64
	_ = v1009
	var v1010 int64
	_ = v1010
	var v1011 int64
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int64
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1098 int32
	_ = v1098
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1142 int32
	_ = v1142
	var v1143 int64
	_ = v1143
	var v1144 int64
	_ = v1144
	var v1147 int64
	_ = v1147
	var v1148 int64
	_ = v1148
	var v1151 int64
	_ = v1151
	var v1157 int32
	_ = v1157
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1174 int32
	_ = v1174
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1190 int32
	_ = v1190
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	v2 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(240)
	m.G0 = v25
	v28 = *(*int64)(unsafe.Add(mBase, _consts[842]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+160)) = v28
	v31 = *(*int64)(unsafe.Add(mBase, _consts[843]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+152)) = v31
	v34 = *(*int64)(unsafe.Add(mBase, _consts[844]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+144)) = v34
	v37 = *(*int64)(unsafe.Add(mBase, _consts[845]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+136)) = v37
	v40 = *(*int64)(unsafe.Add(mBase, _consts[846]))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+128)) = v40
	v43 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v46 = base.B2i32(v44 == int32(2))
	goto L1
L1:
	;
	if v46 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[458]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+60))
	v60 = m.T0[v59].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, int32(17481), int32(10), v25+int32(128))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L9
	}
L5:
	;
	return int32(0)
L6:
	;
	goto L4
L7:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v1198 != 0 {
		goto L330
	} else {
		goto L331
	}
L8:
	;
	F_pfree(m, v1112)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L5
	} else {
		goto L329
	}
L9:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v62 == int32(2) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v67 = F_MakeSingleTupleTableSlot(m, v65, int32(1650812))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L5
	} else {
		goto L326
	}
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v72 = F_tuplestore_gettupleslot(m, v69, int32(1), int32(0), v67)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	if v72 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v81 = v2
	goto L18
L16:
	;
	v405 = v2
	goto L17
L17:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v425 = F_LWLockAcquire(m, v421+int32(4736), int32(1))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L5
	} else {
		goto L140
	}
L18:
	;
	v97 = F_palloc0(m, int32(48))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	v405 = v388
	goto L17
L20:
	;
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+6)))
	if v99 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_slot_getsomeattrs_int(m, v67, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v107 = F_text_to_cstring(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v107
	v110 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+6)))
	if v110 <= int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_slot_getsomeattrs_int(m, v67, int32(2))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v118 = F_text_to_cstring(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v118
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+6)))
	if v121 <= int32(2) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_slot_getsomeattrs_int(m, v67, int32(3))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+2)))
	if v128 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v133 = int64(0)
	goto L37
L36:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v131)))
	v133 = v132
	goto L37
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v97)+24)) = v133
	v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+6)))
	if v135 <= int32(3) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_slot_getsomeattrs_int(m, v67, int32(4))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+3)))
	if v142 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v147 = int64(0)
	goto L44
L43:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
	v147 = v146
	goto L44
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v97)+16)) = v147
	v149 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+6)))
	if v149 <= int32(4) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_slot_getsomeattrs_int(m, v67, int32(5))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+4)))
	if v159 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v160 = int32(0)
	goto L51
L50:
	;
	v160 = v157
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+40)) = v160
	v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+6)))
	if v162 <= int32(5) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_slot_getsomeattrs_int(m, v67, int32(6))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+12)) = uint8(base.B2i32(v169 != int32(0)))
	v173 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+6)))
	if v173 <= int32(6) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	F_slot_getsomeattrs_int(m, v67, int32(7))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+6)))
	if v180 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	v185 = int64(0)
	goto L62
L61:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+24))
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
	v185 = v184
	goto L62
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v97)+32)) = v185
	v187 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+6)))
	if v187 <= int32(7) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_slot_getsomeattrs_int(m, v67, int32(8))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+13)) = uint8(base.B2i32(v194 != int32(0)))
	v198 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+6)))
	if v198 <= int32(8) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	F_slot_getsomeattrs_int(m, v67, int32(9))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L5
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+32))
	v206 = F_text_to_cstring(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L5
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v206
	v209 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+6)))
	if v209 <= int32(9) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_slot_getsomeattrs_int(m, v67, int32(10))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L5
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v215 = int32(0)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+9)))
	if v217 == v215 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+36))
	v222 = F_text_to_cstring(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	v371 = v215
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+44)) = v371
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v97)+16))
	if v374 == int64(0) {
		goto L130
	} else {
		goto L131
	}
L79:
	;
	v224 = int32(0)
	v226 = int32(390159)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v230 = int32(*(*uint8)(unsafe.Add(mBase, _consts[847])))
	if v230 == v224 {
		v249 = v229
		v250 = v230
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v371 = v370
	goto L78
L81:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	v370 = v369
	goto L80
L82:
	;
	if v250-v249 == int32(0) {
		v368 = int32(1659504)
		goto L81
	} else {
		goto L90
	}
L83:
	;
	goto L82
L84:
	;
	if v229 != v230 {
		v249 = v229
		v250 = v230
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v234 = v226
	v235 = v222
	goto L86
L86:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+1)))
	if v239 == int32(0) {
		v249 = v238
		v250 = v239
		goto L83
	} else {
		goto L88
	}
L87:
	;
	v249 = v238
	v250 = v239
	goto L83
L88:
	;
	v242 = int32(1)
	if v238 == v239 {
		v234 = v234 + v242
		v235 = v235 + v242
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v255 = int32(460460)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, _consts[848])))
	if v259 == int32(0) {
		v278 = v258
		v279 = v259
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v279-v278 == int32(0) {
		v368 = int32(1659512)
		goto L81
	} else {
		goto L99
	}
L92:
	;
	goto L91
L93:
	;
	if v258 != v259 {
		v278 = v258
		v279 = v259
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v263 = v255
	v264 = v222
	goto L95
L95:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+1)))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+1)))
	if v268 == int32(0) {
		v278 = v267
		v279 = v268
		goto L92
	} else {
		goto L97
	}
L96:
	;
	v278 = v267
	v279 = v268
	goto L92
L97:
	;
	v271 = int32(1)
	if v267 == v268 {
		v263 = v263 + v271
		v264 = v264 + v271
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v284 = int32(460447)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, _consts[849])))
	if v288 == int32(0) {
		v307 = v287
		v308 = v288
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if v308-v307 == int32(0) {
		v368 = int32(1659520)
		goto L81
	} else {
		goto L108
	}
L101:
	;
	goto L100
L102:
	;
	if v287 != v288 {
		v307 = v287
		v308 = v288
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v292 = v284
	v293 = v222
	goto L104
L104:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+1)))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+1)))
	if v297 == int32(0) {
		v307 = v296
		v308 = v297
		goto L101
	} else {
		goto L106
	}
L105:
	;
	v307 = v296
	v308 = v297
	goto L101
L106:
	;
	v300 = int32(1)
	if v296 == v297 {
		v292 = v292 + v300
		v293 = v293 + v300
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v313 = int32(103238)
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v317 = int32(*(*uint8)(unsafe.Add(mBase, _consts[850])))
	if v317 == int32(0) {
		v336 = v316
		v337 = v317
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if v337-v336 == int32(0) {
		v368 = int32(1659528)
		goto L81
	} else {
		goto L117
	}
L110:
	;
	goto L109
L111:
	;
	if v316 != v317 {
		v336 = v316
		v337 = v317
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v321 = v313
	v322 = v222
	goto L113
L113:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322)+1)))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321)+1)))
	if v326 == int32(0) {
		v336 = v325
		v337 = v326
		goto L110
	} else {
		goto L115
	}
L114:
	;
	v336 = v325
	v337 = v326
	goto L110
L115:
	;
	v329 = int32(1)
	if v325 == v326 {
		v321 = v321 + v329
		v322 = v322 + v329
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v341 = int32(71526)
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, _consts[851])))
	if v345 == int32(0) {
		v364 = v344
		v365 = v345
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v365-v364 != 0 {
		v370 = v224
		goto L80
	} else {
		goto L126
	}
L119:
	;
	goto L118
L120:
	;
	if v344 != v345 {
		v364 = v344
		v365 = v345
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v349 = v341
	v350 = v222
	goto L122
L122:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	if v354 == int32(0) {
		v364 = v353
		v365 = v354
		goto L119
	} else {
		goto L124
	}
L123:
	;
	v364 = v353
	v365 = v354
	goto L119
L124:
	;
	v357 = int32(1)
	if v353 == v354 {
		v349 = v349 + v357
		v350 = v350 + v357
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v368 = int32(1659536)
	goto L81
L127:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+12))
	m.T0[v390].(func(*base.Module, int32))(m, v67)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L5
	} else {
		goto L137
	}
L128:
	;
	v386 = F_lappend(m, v81, v97)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L5
	} else {
		goto L136
	}
L129:
	;
	F_pfree(m, v97)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L5
	} else {
		goto L135
	}
L130:
	;
	if v371 != 0 {
		goto L128
	} else {
		goto L134
	}
L131:
	;
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v97)+24))
	if v377 == int64(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v97)+40))
	if v380|v371 == int32(0) {
		goto L129
	} else {
		goto L133
	}
L133:
	;
	goto L128
L134:
	;
	goto L129
L135:
	;
	v388 = v81
	goto L127
L136:
	;
	v388 = v386
	goto L127
L137:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v396 = F_tuplestore_gettupleslot(m, v393, int32(1), int32(0), v67)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L5
	} else {
		goto L138
	}
L138:
	;
	if v396 != 0 {
		v81 = v388
		goto L18
	} else {
		goto L139
	}
L139:
	;
	goto L19
L140:
	;
	v428 = *(*int32)(unsafe.Add(mBase, _consts[852]))
	if v428 <= int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	if v405 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L142:
	;
	v432 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v432+int32(4736))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L5
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _consts[853]))
	v440 = int32(0)
	v441 = v428
	v446 = v438
	v449 = v2
	goto L146
L145:
	;
	goto L141
L146:
	;
	v464 = v446 + v440*int32(288)
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+4)))
	if v465 != int32(1) {
		v477 = v441
		v478 = v446
		v479 = v449
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v484 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v484+int32(4736))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L5
	} else {
		goto L153
	}
L148:
	;
	v481 = v440 + int32(1)
	if v481 < v477 {
		v440 = v481
		v441 = v477
		v446 = v478
		v449 = v479
		goto L146
	} else {
		goto L152
	}
L149:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+201)))
	if v468 == int32(0) {
		v477 = v441
		v478 = v446
		v479 = v449
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v471 = F_lappend(m, v449, v464)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L5
	} else {
		goto L151
	}
L151:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _consts[852]))
	v476 = *(*int32)(unsafe.Add(mBase, _consts[853]))
	v477 = v474
	v478 = v476
	v479 = v471
	goto L148
L152:
	;
	goto L147
L153:
	;
	if v479 == int32(0) {
		goto L141
	} else {
		goto L154
	}
L154:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	if v491 <= int32(0) {
		goto L141
	} else {
		goto L155
	}
L155:
	;
	v499 = v2
	goto L156
L156:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v479)+12))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v516+v499<<(uint(int32(2))%32))))
	v522 = v520 + int32(24)
	if v405 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	goto L141
L158:
	;
	v705 = v499 + int32(1)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	if v705 < v706 {
		v499 = v705
		goto L156
	} else {
		goto L201
	}
L159:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v520)+88))
	F_LockSharedObject(m, int32(1262), v624, int32(1))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L5
	} else {
		goto L184
	}
L160:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v525 <= int32(0) {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v405)+12))
	v530 = int32(0)
	goto L162
L162:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v528+v530<<(uint(int32(2))%32))))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	if v560 == int32(0) {
		v579 = v559
		v580 = v560
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = int32(1)
	if v585 != 0 {
		goto L176
	} else {
		goto L177
	}
L164:
	;
	if v580-v579 != 0 {
		goto L172
	} else {
		goto L173
	}
L165:
	;
	goto L164
L166:
	;
	if v559 != v560 {
		v579 = v559
		v580 = v560
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v564 = v556
	v565 = v522
	goto L168
L168:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+1)))
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564)+1)))
	if v569 == int32(0) {
		v579 = v568
		v580 = v569
		goto L165
	} else {
		goto L170
	}
L169:
	;
	v579 = v568
	v580 = v569
	goto L165
L170:
	;
	v572 = int32(1)
	if v568 == v569 {
		v564 = v564 + v572
		v565 = v565 + v572
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	v583 = v530 + int32(1)
	if v583 != v525 {
		v530 = v583
		goto L162
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	goto L163
L175:
	;
	goto L159
L176:
	;
	F_s_lock(m, v520, int32(523799), int32(395), int32(470833))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L5
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v555)+44))
	if v593 != 0 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L178
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = int32(0)
	goto L158
L181:
	;
	goto L182
L182:
	;
	v596 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = v596
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v520)+112))
	if v598 == v596 {
		goto L158
	} else {
		goto L183
	}
L183:
	;
	goto L159
L184:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v520)))
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = int32(1)
	if v628 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	F_s_lock(m, v520, int32(523799), int32(461), int32(126267))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L5
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520)+4)))
	if v636 == int32(1) {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L187
L189:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v520)+88))
	F_UnlockSharedObject(m, int32(1262), v659, int32(1))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L5
	} else {
		goto L196
	}
L190:
	;
	v639 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = v639
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520)+201)))
	if v641 == v639 {
		goto L189
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = int32(0)
	goto L189
L193:
	;
	F_ReplicationSlotAcquire(m, v522, int32(1), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	v648 = int32(4464484)
	v649 = *(*int32)(unsafe.Add(mBase, _consts[841]))
	*(*int32)(unsafe.Add(mBase, _consts[841])) = int32(0)
	F_ReplicationSlotDropPtr(m, v649)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	goto L189
L196:
	;
	v665 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L5
	} else {
		goto L197
	}
L197:
	;
	if v665 == int32(0) {
		goto L158
	} else {
		goto L198
	}
L198:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v520)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+100)) = v669
	*(*int32)(unsafe.Add(mBase, uint32(v25)+96)) = v522
	F_errmsg(m, int32(61665), v25+int32(96))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L5
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(523799), int32(477), int32(126267))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L5
	} else {
		goto L200
	}
L200:
	;
	goto L158
L201:
	;
	goto L157
L202:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L5
	} else {
		goto L323
	}
L203:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L5
	} else {
		goto L319
	}
L204:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L5
	} else {
		goto L315
	}
L205:
	;
	F_list_free_deep(m, v405)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L5
	} else {
		goto L313
	}
L206:
	;
	v1098 = int32(0)
	goto L205
L207:
	;
	goto L208
L208:
	;
	v733 = int32(0)
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v734 <= v733 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1098 = int32(0)
	goto L205
L210:
	;
	goto L211
L211:
	;
	v755 = v733
	v763 = int32(0)
	goto L212
L212:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v405)+12))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v776+v755<<(uint(int32(2))%32))))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v780)+8))
	v783 = F_get_database_oid(m, v781, int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L5
	} else {
		goto L214
	}
L213:
	;
	v1098 = v1083
	goto L205
L214:
	;
	F_LockSharedObject(m, int32(1262), v783, int32(1))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L5
	} else {
		goto L215
	}
L215:
	;
	v788 = m.G0
	v790 = v788 - int32(16)
	m.G0 = v790
	v795 = F_GetWalRcvFlushRecPtr(m, int32(0), v790+int32(8))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L5
	} else {
		goto L216
	}
L216:
	;
	v799 = F_GetXLogReplayRecPtr(m, v790+int32(12))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L5
	} else {
		goto L217
	}
L217:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v790)+12))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v790)+8))
	m.G0 = v790 + int32(16)
	if base.Ui64(v799) < base.Ui64(v795) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v807 = v795
	goto L220
L219:
	;
	v807 = v799
	goto L220
L220:
	;
	if v801 == v802 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v809 = v807
	goto L223
L222:
	;
	v809 = v799
	goto L223
L223:
	;
	v810 = *(*int64)(unsafe.Add(mBase, uint32(v780)+24))
	if base.Ui64(v809) < base.Ui64(v810) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	F_UnlockSharedObject(m, int32(1262), v783, int32(1))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L5
	} else {
		goto L311
	}
L225:
	;
	v812 = int32(0)
	v816 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	if v816 == int32(7) {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	goto L227
L227:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	v849 = F_SearchNamedReplicationSlot(m, v847, int32(1))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L5
	} else {
		goto L236
	}
L228:
	;
	v819 = int32(15)
	goto L230
L229:
	;
	v819 = int32(21)
	goto L230
L230:
	;
	v821 = F_errstart(m, v819, int32(0))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L5
	} else {
		goto L231
	}
L231:
	;
	if v821 == int32(0) {
		v1072 = v812
		goto L224
	} else {
		goto L232
	}
L232:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L5
	} else {
		goto L233
	}
L233:
	;
	v828 = *(*int64)(unsafe.Add(mBase, uint32(v780)+24))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	*(*uint32)(unsafe.Add(mBase, uint32(v25+int32(16)))) = uint32(v809)
	v831 = int64(32)
	v832 = int64(base.Ui64(v809) >> (uint(v831) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v25)+12)) = uint32(v832)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v829
	*(*uint32)(unsafe.Add(mBase, uint32(v25)+4)) = uint32(v828)
	v837 = int64(base.Ui64(v828) >> (uint(v831) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v25))) = uint32(v837)
	F_errmsg(m, int32(539561), v25)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L5
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(523799), int32(649), int32(91314))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L5
	} else {
		goto L235
	}
L235:
	;
	v1072 = v812
	goto L224
L236:
	;
	if v849 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v849)))
	*(*int32)(unsafe.Add(mBase, uint32(v849))) = int32(1)
	if v851 != 0 {
		goto L240
	} else {
		goto L241
	}
L238:
	;
	goto L239
L239:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v780)+44))
	if v910 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L240:
	;
	F_s_lock(m, v849, int32(523799), int32(659), int32(91314))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L5
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	v859 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v849))) = v859
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849)+201)))
	if v861 == v859 {
		goto L204
	} else {
		goto L244
	}
L243:
	;
	goto L242
L244:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	F_ReplicationSlotAcquire(m, v864, int32(1), int32(0))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L5
	} else {
		goto L245
	}
L245:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v849)+112))
	if v869 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L5
	} else {
		goto L266
	}
L247:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v780)+44))
	if v870 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v849)))
	*(*int32)(unsafe.Add(mBase, uint32(v849))) = int32(1)
	if v871 != 0 {
		goto L251
	} else {
		goto L252
	}
L249:
	;
	goto L250
L250:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v849)+92))
	if v889 == int32(2) {
		goto L258
	} else {
		goto L259
	}
L251:
	;
	F_s_lock(m, v849, int32(523799), int32(698), int32(91314))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L5
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v780)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v849))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v849)+112)) = v879
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L5
	} else {
		goto L255
	}
L254:
	;
	goto L253
L255:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L5
	} else {
		goto L256
	}
L256:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v849)+112))
	if v887 != 0 {
		goto L246
	} else {
		goto L257
	}
L257:
	;
	goto L250
L258:
	;
	v892 = F_update_and_persist_local_synced_slot(m, v780, v783)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L5
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	v896 = *(*int64)(unsafe.Add(mBase, uint32(v780)+24))
	v897 = *(*int64)(unsafe.Add(mBase, uint32(v849)+120))
	if base.Ui64(v896) < base.Ui64(v897) {
		goto L203
	} else {
		goto L263
	}
L261:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L5
	} else {
		goto L262
	}
L262:
	;
	v1072 = v892
	goto L224
L263:
	;
	v899 = int32(0)
	v901 = F_update_local_synced_slot(m, v780, v783, v899, v899)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L5
	} else {
		goto L264
	}
L264:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L5
	} else {
		goto L265
	}
L265:
	;
	v1072 = v901
	goto L224
L266:
	;
	v1072 = base.B2i32(v869 == int32(0))
	goto L224
L267:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	v914 = int32(1)
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780)+12)))
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780)+13)))
	F_ReplicationSlotCreate(m, v913, v914, int32(2), v916, v917, v914)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L5
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v1072 = int32(0)
	goto L224
L270:
	;
	v922 = *(*int32)(unsafe.Add(mBase, _consts[841]))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v780)+4))
	v927 = F_strncpy(m, v25+int32(176), v925, int32(64))
	mBase = m.M
	v928 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v927)+63)) = uint8(v928)
	goto L271
L271:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v922)))
	*(*int32)(unsafe.Add(mBase, uint32(v922))) = int32(1)
	if v930 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	F_s_lock(m, v922, int32(523799), int32(773), int32(91314))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L5
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v922)+88)) = v783
	v939 = *(*int64)(unsafe.Add(mBase, uint32(v25)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v922)+137)) = v939
	v941 = *(*int64)(unsafe.Add(mBase, uint32(v25+int32(184))))
	*(*int64)(unsafe.Add(mBase, uint32(v922)+145)) = v941
	v943 = *(*int64)(unsafe.Add(mBase, uint32(v25)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v922)+153)) = v943
	v945 = *(*int64)(unsafe.Add(mBase, uint32(v25+int32(200))))
	*(*int64)(unsafe.Add(mBase, uint32(v922)+161)) = v945
	v947 = *(*int64)(unsafe.Add(mBase, uint32(v25+int32(208))))
	*(*int64)(unsafe.Add(mBase, uint32(v922)+169)) = v947
	v949 = *(*int64)(unsafe.Add(mBase, uint32(v25+int32(216))))
	*(*int64)(unsafe.Add(mBase, uint32(v922)+177)) = v949
	v951 = *(*int64)(unsafe.Add(mBase, uint32(v25+int32(224))))
	*(*int64)(unsafe.Add(mBase, uint32(v922)+185)) = v951
	v953 = *(*int64)(unsafe.Add(mBase, uint32(v25+int32(232))))
	*(*int64)(unsafe.Add(mBase, uint32(v922)+193)) = v953
	v955 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v922))) = v955
	v957 = *(*int64)(unsafe.Add(mBase, uint32(v780)+16))
	v959 = *(*int32)(unsafe.Add(mBase, _consts[841]))
	v961 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v965 = F_LWLockAcquire(m, v961+int32(4608), v955)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L5
	} else {
		goto L276
	}
L275:
	;
	goto L274
L276:
	;
	v967 = F_GetRedoRecPtr(m)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L5
	} else {
		goto L277
	}
L277:
	;
	v970 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v970)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v970)+440)) = int32(1)
	if v971 != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v975 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	F_s_lock(m, v975+int32(440), int32(521668), int32(2685), int32(552770))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L5
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v984 = *(*int32)(unsafe.Add(mBase, _consts[199]))
	*(*int32)(unsafe.Add(mBase, uint32(v984)+440)) = int32(0)
	v987 = *(*int64)(unsafe.Add(mBase, uint32(v984)+224))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v959)))
	*(*int32)(unsafe.Add(mBase, uint32(v959))) = int32(1)
	if base.Ui64(v967) < base.Ui64(v987) {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	goto L280
L282:
	;
	v992 = v967
	goto L284
L283:
	;
	v992 = v987
	goto L284
L284:
	;
	if v987 == int64(0) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v995 = v967
	goto L287
L286:
	;
	v995 = v992
	goto L287
L287:
	;
	if v988 != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	F_s_lock(m, v959, int32(523799), int32(539), int32(91275))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L5
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v959))) = int32(0)
	if base.Ui64(v995) < base.Ui64(v957) {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	goto L290
L292:
	;
	v1004 = v957
	goto L294
L293:
	;
	v1004 = v995
	goto L294
L294:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v959)+104)) = v1004
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L5
	} else {
		goto L295
	}
L295:
	;
	v1009 = int64(*(*int32)(unsafe.Add(mBase, _consts[271])))
	v1010 = *(*int64)(unsafe.Add(mBase, uint32(v959)+104))
	v1011 = F_XLogGetLastRemovedSegno(m)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L5
	} else {
		goto L296
	}
L296:
	;
	v1013 = base.I64_div_u_s(v1010, v1009)
	if base.Ui64(v1013) <= base.Ui64(v1011) {
		goto L202
	} else {
		goto L297
	}
L297:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v1016+int32(4608))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L5
	} else {
		goto L298
	}
L298:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v1026 = F_LWLockAcquire(m, v1022+int32(4736), int32(0))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L5
	} else {
		goto L299
	}
L299:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v1033 = F_LWLockAcquire(m, v1029+int32(512), int32(0))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L5
	} else {
		goto L300
	}
L300:
	;
	v1036 = F_GetOldestSafeDecodingTransactionId(m, int32(1))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L5
	} else {
		goto L301
	}
L301:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v922)))
	*(*int32)(unsafe.Add(mBase, uint32(v922))) = int32(1)
	if v1038 != 0 {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	F_s_lock(m, v922, int32(523799), int32(783), int32(91314))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L5
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v922)+100)) = v1036
	*(*int32)(unsafe.Add(mBase, uint32(v922)+20)) = v1036
	*(*int32)(unsafe.Add(mBase, uint32(v922))) = int32(0)
	v1050 = int32(1)
	F_ReplicationSlotsComputeRequiredXmin(m, v1050)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L5
	} else {
		goto L306
	}
L305:
	;
	goto L304
L306:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v1055+int32(512))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L5
	} else {
		goto L307
	}
L307:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v1061+int32(4736))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L5
	} else {
		goto L308
	}
L308:
	;
	v1066 = F_update_and_persist_local_synced_slot(m, v780, v783)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L5
	} else {
		goto L309
	}
L309:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L5
	} else {
		goto L310
	}
L310:
	;
	v1072 = v1050
	goto L224
L311:
	;
	v1083 = v1072 | v763
	v1085 = v755 + int32(1)
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v1085 < v1086 {
		v755 = v1085
		v763 = v1083
		goto L212
	} else {
		goto L312
	}
L312:
	;
	goto L213
L313:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	if v1112 != 0 {
		goto L8
	} else {
		goto L314
	}
L314:
	;
	goto L7
L315:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L5
	} else {
		goto L316
	}
L316:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v1120
	F_errmsg(m, int32(24163), v25+int32(48))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L5
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(523799), int32(669), int32(91314))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L5
	} else {
		goto L318
	}
L318:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L319:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+80)) = v1136
	F_errmsg_internal(m, int32(730730), v25+int32(80))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L5
	} else {
		goto L320
	}
L320:
	;
	v1143 = *(*int64)(unsafe.Add(mBase, uint32(v849)+120))
	v1144 = *(*int64)(unsafe.Add(mBase, uint32(v780)+24))
	*(*uint32)(unsafe.Add(mBase, uint32(v25)+76)) = uint32(v1144)
	*(*uint32)(unsafe.Add(mBase, uint32(v25)+68)) = uint32(v1143)
	v1147 = int64(32)
	v1148 = int64(base.Ui64(v1144) >> (uint(v1147) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v25)+72)) = uint32(v1148)
	v1151 = int64(base.Ui64(v1143) >> (uint(v1147) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v25)+64)) = uint32(v1151)
	F_errdetail_internal(m, int32(691174), v25-int32(-64))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L5
	} else {
		goto L321
	}
L321:
	;
	F_errfinish(m, int32(523799), int32(739), int32(91314))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L5
	} else {
		goto L322
	}
L322:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v959 + int32(24)
	F_errmsg_internal(m, int32(19173), v25+int32(32))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L5
	} else {
		goto L324
	}
L324:
	;
	F_errfinish(m, int32(523799), int32(548), int32(91275))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L5
	} else {
		goto L325
	}
L325:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L326:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+112)) = v1184
	F_errmsg(m, int32(211276), v25+int32(112))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L5
	} else {
		goto L327
	}
L327:
	;
	F_errfinish(m, int32(523799), int32(839), int32(126249))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L5
	} else {
		goto L328
	}
L328:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L329:
	;
	goto L7
L330:
	;
	F_tuplestore_end(m, v1198)
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L5
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if v1201 != 0 {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	goto L332
L334:
	;
	F_FreeTupleDesc(m, v1201)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L5
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	F_pfree(m, v60)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L5
	} else {
		goto L338
	}
L337:
	;
	goto L336
L338:
	;
	if v46 == int32(0) {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L5
	} else {
		goto L342
	}
L340:
	;
	goto L341
L341:
	;
	m.G0 = v25 + int32(240)
	return v1098 & int32(1)
L342:
	;
	goto L341
}
