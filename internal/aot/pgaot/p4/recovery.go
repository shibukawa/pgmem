package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RecoveryRestartPoint(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v36 int32
	_ = v36
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v130 int32
	_ = v130
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	if v12 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+412))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+376))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+364))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+352))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+340))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+328))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+316))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+304))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+292))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+280))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+268))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+256))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+244))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+232))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+220))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+208))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v15)+196))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v15)+184))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v15)+172))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+160))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+148))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+136))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+124))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)+112))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+88))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64))))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v82 = v18 + (v19 + (v20 + (v21 + (v22 + (v23 + (v24 + (v25 + (v26 + (v27 + (v28 + (v29 + (v30 + (v31 + (v32 + (v33 + (v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v46 + (v47 + (v48 + (v49 + (v50 + v16))))))))))))))))))))))))))))))
		} else {
			v82 = v16
		}
		if int32(0) < v82 {
			v87 = int32(1)
		} else {
			v87 = int32(0)
		}
	} else {
		v87 = int32(0)
	}
	if v87 != 0 {
		v90 = F_errstart(m, int32(13), int32(0))
		mBase = m.M
		v91 = m.ExcPending
		if v91 != 0 {
			return
		} else {
			if v90 == int32(0) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v94 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*uint32)(unsafe.Add(mBase, uint32(v9)+4)) = uint32(v94)
				v97 = int64(base.Ui64(v94) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v9))) = uint32(v97)
				F_errmsg_internal(m, int32(174856), v9)
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return
				} else {
					F_errfinish(m, int32(510097), int32(7606), int32(90814))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v108 = *(*int32)(unsafe.Add(mBase, _consts[199]))
		v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+440))
		*(*int32)(unsafe.Add(mBase, uint32(v108)+440)) = int32(1)
		if v109 != 0 {
			v113 = *(*int32)(unsafe.Add(mBase, _consts[199]))
			F_s_lock(m, v113+int32(440), int32(510097), int32(7614), int32(90814))
			mBase = m.M
			v120 = m.ExcPending
			if v120 != 0 {
				return
			} else {
				v122 = *(*int32)(unsafe.Add(mBase, _consts[199]))
				v123 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
				*(*int64)(unsafe.Add(mBase, uint32(v122)+328)) = v123
				v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
				*(*int64)(unsafe.Add(mBase, uint32(v122)+336)) = v125
				v130 = F__emscripten_memcpy_bulkmem(m, v122+int32(344), l0, int32(88))
				mBase = m.M
				*(*int32)(unsafe.Add(mBase, uint32(v122)+440)) = int32(0)
				m.G0 = v9 + int32(16)
				return
			}
		} else {
			v122 = *(*int32)(unsafe.Add(mBase, _consts[199]))
			v123 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v122)+328)) = v123
			v125 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
			*(*int64)(unsafe.Add(mBase, uint32(v122)+336)) = v125
			v130 = F__emscripten_memcpy_bulkmem(m, v122+int32(344), l0, int32(88))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(v122)+440)) = int32(0)
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_ResolveRecoveryConflictWithVirtualXIDs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v73 int64
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int64
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v129 int64
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int64
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int64
	_ = v315
	var v316 int64
	_ = v316
	var v325 int64
	_ = v325
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int64
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int64
	_ = v410
	var v411 int64
	_ = v411
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	v5 = int32(0)
	v14 = int64(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v18 + int32(48)
	return
L2:
	;
	if l3 == int32(0) {
		v54 = v14
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v55 = l0
	v59 = v5
	v61 = v5
	goto L11
L4:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _consts[926])))
	if v26 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, _consts[887])))
	if v30 != int32(1) {
		v54 = v14
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v36 = m.G0
	v37 = int32(16)
	v38 = v36 - v37
	m.G0 = v38
	F___gettimeofday(m, v38)
	mBase = m.M
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
	v42 = int64(*(*int32)(unsafe.Add(mBase, uint32(v38)+8)))
	m.G0 = v38 + v37
	goto L9
L8:
	;
	goto L7
L9:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v51 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v54 = v42 + v41*int64(1000000) - int64(946684800000000)
	goto L3
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1167])) = int32(1000)
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v73
	v78 = F_VirtualXactLock(m, v18+int32(24), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v384&int32(1) != 0 {
		goto L82
	} else {
		goto L83
	}
L13:
	;
	return
L14:
	;
	if v78 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v86 = v59
	v88 = v61
	goto L18
L16:
	;
	v384 = v59
	v386 = v61
	goto L17
L17:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v55+int32(12))))
	if v399 != 0 {
		v55 = v55 + int32(8)
		v59 = v384
		v61 = v386
		goto L11
	} else {
		goto L80
	}
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v98 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v384 = v366
	v386 = v370
	goto L17
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L13
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v106 = *(*int64)(unsafe.Add(mBase, _consts[308]))
	*(*int64)(unsafe.Add(mBase, uint32(v18+int32(40)))) = v106
	v109 = *(*int32)(unsafe.Add(mBase, _consts[306]))
	*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(39)))) = uint8(base.B2i32(v109 == int32(3)))
	goto L24
L23:
	;
	goto L22
L24:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+39)))
	if v113 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	if v54 == int64(0) {
		v355 = v88
		goto L58
	} else {
		goto L59
	}
L26:
	;
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v174
	v177 = v18 + int32(16)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v181 = *(*int32)(unsafe.Add(mBase, _consts[1131]))
	v183 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v187 = F_LWLockAcquire(m, v183+int32(512), int32(1))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L13
	} else {
		goto L41
	}
L27:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = l2
	v157 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
	F_pg_usleep(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L37
	}
L28:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
	v129 = v125 + base.I64_extend_i32_u(v124)*int64(1000)
	if v129 == int64(0) {
		goto L27
	} else {
		goto L34
	}
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[927]))
	if int32(0) <= v117 {
		v124 = v117
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[928]))
	if v121 < int32(0) {
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L27
L33:
	;
	v124 = v121
	goto L28
L34:
	;
	v135 = m.G0
	v136 = int32(16)
	v137 = v135 - v136
	m.G0 = v137
	F___gettimeofday(m, v137)
	mBase = m.M
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v137)))
	v141 = int64(*(*int32)(unsafe.Add(mBase, uint32(v137)+8)))
	m.G0 = v137 + v136
	goto L35
L35:
	;
	if v129 <= v141+v140*int64(1000000)-int64(946684800000000) {
		goto L26
	} else {
		goto L36
	}
L36:
	;
	goto L27
L37:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = int32(0)
	v164 = int32(4142604)
	v165 = int32(1000000)
	v167 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
	v169 = v167 << (uint(int32(1)) % 32)
	if v165 <= v169 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v172 = v165
	goto L40
L39:
	;
	v172 = v169
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1167])) = v172
	goto L25
L41:
	;
	v189 = int32(0)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	if v190 <= v189 {
		v242 = v189
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v251+int32(512))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L13
	} else {
		goto L53
	}
L43:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _consts[1132]))
	v204 = v189
	goto L44
L44:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v181+int32(36)+v204<<(uint(int32(2))%32))))
	v218 = v196 + v215*int32(640)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+52))
	if v219 != v179 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v242 = int32(0)
	goto L42
L46:
	;
	goto L45
L47:
	;
	v231 = v204 + int32(1)
	if v231 != v190 {
		v204 = v231
		goto L44
	} else {
		goto L52
	}
L48:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v218)+56))
	if v221 != v178 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v223 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v218)+73)) = uint8(v223)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v218)+44))
	if v225 == int32(0) {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v228 = F_SendProcSignal(m, v225, l1, v179)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	v242 = v225
	goto L42
L52:
	;
	goto L46
L53:
	;
	if v242 == int32(0) {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	F_pg_usleep(m, int32(5000))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	goto L25
L56:
	;
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v371
	v376 = F_VirtualXactLock(m, v18+int32(8), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L13
	} else {
		goto L78
	}
L57:
	;
	v366 = v361
	v370 = v360
	goto L56
L58:
	;
	v360 = v355
	v361 = v86
	goto L57
L59:
	;
	v278 = int32(1)
	v280 = v88 ^ v278
	v282 = v86 ^ v278
	if v282&v278 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v287 = int32(1)
	if v280&v287 == int32(0) {
		v366 = v278
		v370 = v287
		goto L56
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, _consts[887])))
	v294 = v293 & v280
	v296 = int32(*(*uint8)(unsafe.Add(mBase, _consts[926])))
	v299 = v282 & v296 & int32(1)
	if v299 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L62
L64:
	;
	if v294&int32(1) == int32(0) {
		v339 = v88
		goto L70
	} else {
		goto L71
	}
L65:
	;
	if v294&int32(1) == int32(0) {
		v325 = int64(0)
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v310 = m.G0
	v311 = int32(16)
	v312 = v310 - v311
	m.G0 = v312
	F___gettimeofday(m, v312)
	mBase = m.M
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v312)))
	v316 = int64(*(*int32)(unsafe.Add(mBase, uint32(v312)+8)))
	m.G0 = v312 + v311
	goto L69
L68:
	;
	goto L67
L69:
	;
	v325 = v316 + v315*int64(1000000) - int64(946684800000000)
	goto L64
L70:
	;
	if v299 == int32(0) {
		v355 = v339
		goto L58
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if base.B2i32(base.I64_extend_i32_s(int32(500))*int64(1000) <= v325-v54) == int32(0) {
		v339 = v88
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v339 = int32(1)
	goto L70
L74:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _consts[925]))
	goto L75
L75:
	;
	if base.B2i32(base.I64_extend_i32_s(v343)*int64(1000) <= v325-v54) == int32(0) {
		v355 = v339
		goto L58
	} else {
		goto L76
	}
L76:
	;
	v351 = int32(1)
	F_LogRecoveryConflict(m, l1, v54, v325, v55, v351)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	v360 = v339
	v361 = v351
	goto L57
L78:
	;
	if v376 == int32(0) {
		v86 = v366
		v88 = v370
		goto L18
	} else {
		goto L79
	}
L79:
	;
	goto L19
L80:
	;
	goto L12
L81:
	;
	goto L1
L82:
	;
	v405 = m.G0
	v406 = int32(16)
	v407 = v405 - v406
	m.G0 = v407
	F___gettimeofday(m, v407)
	mBase = m.M
	v410 = *(*int64)(unsafe.Add(mBase, uint32(v407)))
	v411 = int64(*(*int32)(unsafe.Add(mBase, uint32(v407)+8)))
	m.G0 = v407 + v406
	goto L85
L83:
	;
	goto L84
L84:
	;
	if v386&int32(1) == int32(0) {
		goto L1
	} else {
		goto L88
	}
L85:
	;
	v420 = int32(0)
	F_LogRecoveryConflict(m, l1, v54, v411+v410*int64(1000000)-int64(946684800000000), v420, v420)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	if v386&int32(1) != 0 {
		goto L81
	} else {
		goto L87
	}
L87:
	;
	goto L1
L88:
	;
	goto L81
}
func F_check_recovery_prefetch(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return int32(1)
}
func F_check_recovery_target(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = int32(363031)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[309])))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v10 == int32(0) {
		v29 = v9
		v30 = v10
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v50
L2:
	;
	if v30-v29 == int32(0) {
		v50 = v4
		goto L1
	} else {
		goto L10
	}
L3:
	;
	goto L2
L4:
	;
	if v9 != v10 {
		v29 = v9
		v30 = v10
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v14 = v5
	v15 = v6
	goto L6
L6:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v19 == int32(0) {
		v29 = v18
		v30 = v19
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v29 = v18
	v30 = v19
	goto L3
L8:
	;
	v22 = int32(1)
	if v18 == v19 {
		v14 = v14 + v22
		v15 = v15 + v22
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v34 == int32(0) {
		v50 = v4
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v39
	goto L12
L12:
	;
	v45 = F_format_elog_string(m, int32(682704), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v45
	v50 = int32(0)
	goto L1
}
func F_check_recovery_target_lsn(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v218 int64
	_ = v218
	var v226 int64
	_ = v226
	var v232 int64
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v250
L2:
	;
	v15 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v15)
	v19 = v11 + int32(15)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v15)
	v22 = int32(549891)
	v26 = m.G0
	v28 = v26 - int32(32)
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v29
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _consts[310])))
	if v37 == v15 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	goto L4
L4:
	;
	v250 = int32(1)
	goto L1
L5:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v233 != 0 {
		v250 = v15
		goto L1
	} else {
		goto L60
	}
L6:
	;
	v119 = v110 + int32(1)
	v120 = int32(549891)
	v124 = m.G0
	v126 = v124 - int32(32)
	v127 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v126)+24)) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v126)+16)) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = v127
	*(*int64)(unsafe.Add(mBase, uint32(v126))) = v127
	v135 = int32(*(*uint8)(unsafe.Add(mBase, _consts[310])))
	if v135 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L7:
	;
	if base.Ui32(int32(-8)) <= base.Ui32(v105-int32(9)) {
		goto L28
	} else {
		goto L29
	}
L8:
	;
	v105 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, _consts[311])))
	if v41 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = v13
	goto L14
L12:
	;
	goto L13
L13:
	;
	v55 = v22
	v56 = v37
	goto L17
L14:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v51 == v37 {
		v45 = v45 + int32(1)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v105 = v45 - v13
	goto L7
L16:
	;
	goto L15
L17:
	;
	v63 = v28 + int32(base.Ui32(v56)>>(uint(int32(3))%32))&int32(28)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v64 | v65<<(uint(v56)%32)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v69 != 0 {
		v55 = v55 + v65
		v56 = v69
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v72 == int32(0) {
		v97 = v13
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v105 = v97 - v13
	goto L7
L21:
	;
	v76 = v13
	v77 = v72
	goto L22
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v28+int32(base.Ui32(v77)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v85)>>(uint(v77)%32))&int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v97 = v93
	goto L20
L24:
	;
	v97 = v76
	goto L20
L25:
	;
	goto L26
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v93 = v76 + int32(1)
	if v91 != 0 {
		v76 = v93
		v77 = v91
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v110 = v13 + v105
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v111 == int32(47) {
		goto L6
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v115)
	v232 = int64(0)
	goto L5
L31:
	;
	goto L30
L32:
	;
	v218 = F_strtox_2(m, v13, int32(0), int32(16), int64(4294967295))
	mBase = m.M
	goto L58
L33:
	;
	if base.Ui32(int32(-8)) <= base.Ui32(v203-int32(9)) {
		goto L54
	} else {
		goto L55
	}
L34:
	;
	v203 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, _consts[311])))
	if v139 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v143 = v119
	goto L40
L38:
	;
	goto L39
L39:
	;
	v153 = v120
	v154 = v135
	goto L43
L40:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v149 == v135 {
		v143 = v143 + int32(1)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v203 = v143 - v119
	goto L33
L42:
	;
	goto L41
L43:
	;
	v161 = v126 + int32(base.Ui32(v154)>>(uint(int32(3))%32))&int32(28)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v163 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v161))) = v162 | v163<<(uint(v154)%32)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	if v167 != 0 {
		v153 = v153 + v163
		v154 = v167
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v170 == int32(0) {
		v195 = v119
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	v203 = v195 - v119
	goto L33
L47:
	;
	v174 = v119
	v175 = v170
	goto L48
L48:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v126+int32(base.Ui32(v175)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v183)>>(uint(v175)%32))&int32(1) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v195 = v191
	goto L46
L50:
	;
	v195 = v174
	goto L46
L51:
	;
	goto L52
L52:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
	v191 = v174 + int32(1)
	if v189 != 0 {
		v174 = v191
		v175 = v189
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119+v203))))
	if v209 == int32(0) {
		goto L32
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v212 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v212)
	v232 = int64(0)
	goto L5
L57:
	;
	goto L56
L58:
	;
	v226 = F_strtox_2(m, v119, int32(0), int32(16), int64(4294967295))
	mBase = m.M
	goto L59
L59:
	;
	v232 = base.I64_extend_i32_u(base.I32_wrap_i64(v218))<<(uint(int64(32))%64) | base.I64_extend_i32_u(base.I32_wrap_i64(v226))
	goto L5
L60:
	;
	v235 = F_guc_malloc(m, int32(8))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	return int32(0)
L62:
	;
	if v235 == int32(0) {
		v250 = v15
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v235))) = v232
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v235
	goto L4
}
func F_check_recovery_target_timeline(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int64
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(95213)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[312])))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v11 == v4 {
		v30 = v10
		v31 = v11
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v74
	goto L31
L2:
	;
	v83 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L26
	} else {
		goto L27
	}
L3:
	;
	if v31-v30 == int32(0) {
		v81 = v4
		goto L2
	} else {
		goto L11
	}
L4:
	;
	goto L3
L5:
	;
	if v10 != v11 {
		v30 = v10
		v31 = v11
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v15 = v6
	v16 = v7
	goto L7
L7:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v20 == int32(0) {
		v30 = v19
		v31 = v20
		goto L4
	} else {
		goto L9
	}
L8:
	;
	v30 = v19
	v31 = v20
	goto L4
L9:
	;
	v23 = int32(1)
	if v19 == v20 {
		v15 = v15 + v23
		v16 = v16 + v23
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v35 = int32(78799)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _consts[313])))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v39 == int32(0) {
		v58 = v38
		v59 = v39
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v59-v58 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	goto L12
L14:
	;
	if v38 != v39 {
		v58 = v38
		v59 = v39
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v43 = v6
	v44 = v35
	goto L16
L16:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v48 == int32(0) {
		v58 = v47
		v59 = v48
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v58 = v47
	v59 = v48
	goto L13
L18:
	;
	v51 = int32(1)
	if v47 == v48 {
		v43 = v43 + v51
		v44 = v44 + v51
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v81 = int32(1)
	goto L2
L21:
	;
	goto L22
L22:
	;
	v65 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v71 = F_strtox_2(m, v67, v65, v65, int64(4294967295))
	mBase = m.M
	goto L23
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v74 == int32(68) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v74 == int32(28) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v81 = int32(2)
	goto L2
L26:
	;
	return int32(0)
L27:
	;
	if v83 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v81
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v83
	v91 = int32(1)
	goto L30
L29:
	;
	v91 = int32(0)
	goto L30
L30:
	;
	return v91
L31:
	;
	v98 = F_format_elog_string(m, int32(624407), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v98
	return int32(0)
}
func F_check_recovery_target_xid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v7 == int32(0) {
		v36 = int32(1)
		return v36
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, _consts[140])) = v10
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v18 = F_strtox_2(m, v14, v10, v10, int64(-1))
		mBase = m.M
		v20 = *(*int32)(unsafe.Add(mBase, _consts[140]))
		if v20 == int32(28) {
			v36 = v10
			return v36
		} else {
			if v20 == int32(68) {
				v36 = v10
				return v36
			} else {
				v26 = F_guc_malloc(m, int32(4))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					if v26 == int32(0) {
						v36 = v10
					} else {
						*(*uint32)(unsafe.Add(mBase, uint32(v26))) = uint32(v18)
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v26
						v36 = int32(1)
					}
					return v36
				}
			}
		}
	}
}
