package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PinBuffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v51 int32
	_ = v51
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = v13 + int32(1)
	v16 = F_GetPrivateRefCountEntry(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v133 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerRemember(m, v138, v15, int32(1629760))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L40
	}
L2:
	;
	return int32(0)
L3:
	;
	if v16 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = int32(4436184)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[278])) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v15
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v33 = v30
	goto L7
L5:
	;
	goto L6
L6:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v130 = v16
	v132 = v124
	goto L1
L7:
	;
	if v33&int32(4194304) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = int32(457275)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(6287)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(496301)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v51&int32(4194304) != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v102 = v33
	goto L11
L11:
	;
	v111 = v102 + int32(1)
	v113 = v111 & int32(3932160)
	if l1 != 0 {
		goto L30
	} else {
		goto L31
	}
L12:
	;
	goto L15
L13:
	;
	v71 = v51
	goto L14
L14:
	;
	v80 = int32(4126684)
	v81 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(8))+8))
	if v83 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	F_perform_spin_delay(m, v11+int32(8))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L17
	}
L16:
	;
	v71 = v66
	goto L14
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v66&int32(4194304) != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v102 = v71
	goto L11
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[276])) = v98
	goto L20
L22:
	;
	if int32(999) < v81 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v81 < int32(11) {
		goto L20
	} else {
		goto L29
	}
L25:
	;
	v88 = int32(900)
	if v88 <= v81 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v91 = v88
	goto L28
L27:
	;
	v91 = v81
	goto L28
L28:
	;
	v98 = v91 + int32(100)
	goto L21
L29:
	;
	v98 = v81 - int32(1)
	goto L21
L30:
	;
	v118 = base.B2i32(v113 == int32(0))
	goto L32
L31:
	;
	v118 = base.B2i32(base.Ui32(v113) < base.Ui32(int32(1310720)))
	goto L32
L32:
	;
	if v118 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v119 = v102 + int32(262145)
	goto L35
L34:
	;
	v119 = v111
	goto L35
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v121 = base.B2i32(v102 == v120)
	if v102 == v120 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v122 = v119
	goto L38
L37:
	;
	v122 = v120
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v122
	if v102 == v120 {
		v130 = v23
		v132 = v119
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v33 = v120
	goto L7
L40:
	;
	m.G0 = v11 + int32(32)
	return int32(base.Ui32(v132&int32(16777216)) >> (uint(int32(24)) % 32))
}
func F_ReadBufferExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
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
	var v73 int32
	_ = v73
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
	var v90 int64
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
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
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int64
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v244 int64
	_ = v244
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int64
	_ = v264
	var v281 int32
	_ = v281
	var v287 int64
	_ = v287
	var v293 int64
	_ = v293
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+118)))
	if v17 == int32(116) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L9
	} else {
		goto L93
	}
L2:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v23 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v27
	v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v29
	v33 = F_smgropen(m, v14+int32(16), v26)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v52 = v23
	goto L8
L8:
	;
	if l2 == int32(-1) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v33
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
	if v39 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v52 = v51
	goto L8
L12:
	;
	v47 = v39
	goto L14
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
	v47 = v45
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v47 + int32(1)
	goto L11
L15:
	;
	m.G0 = v14 + int32(112)
	return v345
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+36)) = int64(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l0
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v60
	v63 = int32(1)
	if base.Ui32(l3-v63) < base.Ui32(int32(2)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+118)))
	v73 = int32(1)
	if base.Ui32(l3-v73) <= base.Ui32(v73) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v68 = int32(9)
	goto L21
L20:
	;
	v68 = v63
	goto L21
L21:
	;
	v69 = F_ExtendBufferedRel(m, v14, l1, l4, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v345 = v69
	goto L15
L23:
	;
	if v72 == int32(116) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)) = uint8(v72)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v52
	if l3 == int32(3) {
		goto L85
	} else {
		goto L86
	}
L26:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v228 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L27:
	;
	v79 = int32(1)
	v82 = F_LocalBufferAlloc(m, v52, l1, l2, v14+int32(28))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L9
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v94 = F_IOContextForStrategy(m, l4)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L32
	}
L30:
	;
	v84 = int32(3)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)))
	if v85 != int32(1) {
		v223 = v79
		v224 = v82
		v225 = v85
		v227 = v84
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v88 = int32(4418568)
	v90 = *(*int64)(unsafe.Add(mBase, _consts[266]))
	*(*int64)(unsafe.Add(mBase, _consts[266])) = v90 + int64(1)
	v223 = v79
	v224 = v82
	v225 = v85
	v227 = v84
	goto L26
L32:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerEnlarge(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v106
	v112 = F_BufTableHashCode(m, v14+int32(32))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v122 = v115 + v112&int32(127)<<(uint(int32(7))%32) + int32(6912)
	v124 = F_LWLockAcquire(m, v122, int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v128 = F_BufTableLookup(m, v14+int32(32), v112)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L9
	} else {
		goto L38
	}
L37:
	;
	v215 = int32(4418536)
	v217 = *(*int64)(unsafe.Add(mBase, _consts[277]))
	*(*int64)(unsafe.Add(mBase, _consts[277])) = v217 + int64(1)
	v223 = int32(0)
	v224 = v212
	v225 = int32(1)
	v227 = v94
	goto L26
L38:
	;
	if int32(0) <= v128 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v136 = v133 + v128<<(uint(int32(6))%32)
	v137 = F_PinBuffer(m, v136, l4)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L9
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_LWLockRelease(m, v122)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L9
	} else {
		goto L45
	}
L42:
	;
	F_LWLockRelease(m, v122)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)) = uint8(v137)
	if v137 != 0 {
		v212 = v136
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v223 = int32(0)
	v224 = v136
	v225 = int32(0)
	v227 = v94
	goto L26
L45:
	;
	v145 = F_GetVictimBuffer(m, l4, v94)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v150 = F_LWLockAcquire(m, v122, int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	v154 = v148 + v145<<(uint(int32(6))%32)
	v156 = v154 + int32(-64)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v154-int32(44))))
	v162 = F_BufTableInsert(m, v14+int32(32), v112, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	if v162 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v166 = F_LockBufHdr(m, v156)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L9
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_UnpinBuffer(m, v156)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L9
	} else {
		goto L60
	}
L52:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v156)+16)) = v168
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v156)+8)) = v170
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v156))) = v172
	v178 = int32(-2113667072)
	if v72 == int32(112) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v183 = v178
	goto L55
L54:
	;
	v183 = int32(33816576)
	goto L55
L55:
	;
	if l1 == int32(3) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v186 = v178
	goto L58
L57:
	;
	v186 = v183
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154-int32(40)))) = v166&int32(-38010881) | v186
	F_LWLockRelease(m, v122)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)) = uint8(v191)
	v223 = v191
	v224 = v156
	v225 = v191
	v227 = v94
	goto L26
L60:
	;
	F_StrategyFreeBuffer(m, v156)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v204 = v201 + v162<<(uint(int32(6))%32)
	v205 = F_PinBuffer(m, v204, l4)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	F_LWLockRelease(m, v122)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)) = uint8(v205)
	if v205 != 0 {
		v212 = v204
		goto L37
	} else {
		goto L64
	}
L64:
	;
	v223 = int32(0)
	v224 = v204
	v225 = int32(0)
	v227 = v94
	goto L26
L65:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v224)+20))
	v320 = v318 + int32(1)
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)))
	F_ZeroAndLockBuffer(m, v320, l3, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L9
	} else {
		goto L84
	}
L66:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v252 != 0 {
		goto L76
	} else {
		goto L77
	}
L67:
	;
	if v225 == int32(0) {
		goto L65
	} else {
		goto L74
	}
L68:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v231 != int32(1) {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v228)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v228)+112)) = v244 + int64(1)
	goto L67
L71:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v237)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v237)+112)) = v238 + int64(1)
	if v236&int32(1) != 0 {
		goto L66
	} else {
		goto L73
	}
L73:
	;
	goto L65
L74:
	;
	goto L66
L75:
	;
	v281 = v223*int32(320) + v227<<(uint(int32(6))%32)
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_consts[723])))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_consts[723]))) = v287 + int64(1)
	v293 = *(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_consts[724])))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+uint32(_consts[724]))) = v293
	v295 = int32(1)
	F_pgstat_count_backend_io_op(m, v223, v227, int32(2), v295, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _consts[725])) = uint8(v295)
	*(*uint8)(unsafe.Add(mBase, _consts[726])) = uint8(v295)
	goto L82
L76:
	;
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v252)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v252)+120)) = v253 + int64(1)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v257 != int32(1) {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v263)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v263)+120)) = v264 + int64(1)
	if v262&int32(1) == int32(0) {
		goto L65
	} else {
		goto L81
	}
L81:
	;
	goto L75
L82:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, _consts[409])))
	if v305 != int32(1) {
		goto L65
	} else {
		goto L83
	}
L83:
	;
	v308 = int32(4514992)
	v310 = *(*int32)(unsafe.Add(mBase, _consts[410]))
	v312 = *(*int32)(unsafe.Add(mBase, _consts[727]))
	*(*int32)(unsafe.Add(mBase, _consts[410])) = v310 + v312
	goto L65
L84:
	;
	v345 = v320
	goto L15
L85:
	;
	v337 = int32(9)
	goto L87
L86:
	;
	v337 = int32(8)
	goto L87
L87:
	;
	v338 = F_StartReadBuffer(m, v14+int32(32), v14+int32(28), l2, v337)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	if v338 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F_WaitReadBuffers(m, v14+int32(32))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L9
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v345 = v344
	goto L15
L92:
	;
	goto L91
L93:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(144242), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(496301), int32(818), int32(462499))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L9
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
