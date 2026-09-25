package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PinBuffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = v12 + int32(1)
	v15 = F_GetPrivateRefCountEntry(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v126 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_PinBuffer[0]))
	F_ResourceOwnerRemember(m, v131, v14, int32(_a_F_PinBuffer_0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L37
	}
L2:
	;
	return int32(0)
L3:
	;
	if v15 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = int32(_a_F_PinBuffer_1)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_PinBuffer[1]))
	v24 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PinBuffer[1])) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v14
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v32 = v29
	goto L7
L5:
	;
	goto L6
L6:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v123 = v15
	v125 = v118
	goto L1
L7:
	;
	if v32&int32(_a_F_PinBuffer_2) != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v123 = v22
	v125 = v114
	goto L1
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(_a_F_PinBuffer_3)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(_a_F_PinBuffer_4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(_a_F_PinBuffer_5)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v49&int32(_a_F_PinBuffer_2) != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v98 = v32
	goto L11
L11:
	;
	v106 = v98 + int32(1)
	v108 = v106 & int32(_a_F_PinBuffer_6)
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
	v68 = v49
	goto L14
L14:
	;
	v76 = int32(_a_F_PinBuffer_7)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_PinBuffer[2]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(8))+8))
	if v79 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	F_perform_spin_delay(m, v10+int32(8))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L17
	}
L16:
	;
	v68 = v63
	goto L14
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v63&int32(_a_F_PinBuffer_2) != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v98 = v68
	goto L11
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PinBuffer[2])) = v94
	goto L20
L22:
	;
	if int32(999) < v77 {
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v77 < int32(11) {
		goto L20
	} else {
		goto L29
	}
L25:
	;
	v84 = int32(900)
	if v84 <= v77 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v87 = v84
	goto L28
L27:
	;
	v87 = v77
	goto L28
L28:
	;
	v94 = v87 + int32(100)
	goto L21
L29:
	;
	v94 = v77 - int32(1)
	goto L21
L30:
	;
	v113 = base.B2i32(v108 == int32(0))
	goto L32
L31:
	;
	v113 = base.B2i32(base.Ui32(v108) < base.Ui32(int32(_a_F_PinBuffer_8)))
	goto L32
L32:
	;
	if v113 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v114 = v98 + int32(_a_F_PinBuffer_9)
	goto L35
L34:
	;
	v114 = v106
	goto L35
L35:
	;
	v116 = base.AtomicRmwCmpxchg32(m, l0, int32(24), v98, v114)
	if v98 != v116 {
		v32 = v116
		goto L7
	} else {
		goto L36
	}
L36:
	;
	goto L8
L37:
	;
	m.G0 = v10 + int32(32)
	return int32(base.Ui32(v125&int32(16777216)) >> (uint(int32(24)) % 32))
}
func F_ReadBufferExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
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
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
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
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
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
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v223 int64
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int64
	_ = v246
	var v250 int64
	_ = v250
	var v257 int32
	_ = v257
	var v258 int64
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int64
	_ = v269
	var v282 int32
	_ = v282
	var v286 int64
	_ = v286
	var v290 int64
	_ = v290
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+118)))
	if v18 == int32(116) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L9
	} else {
		goto L93
	}
L2:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v24 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v28
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v30
	v34 = F_smgropen(m, v15+int32(16), v27)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v53 = v24
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v34
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v53 = v52
	goto L8
L12:
	;
	v48 = v40
	goto L14
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v48 = v46
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v48 + int32(1)
	goto L11
L15:
	;
	m.G0 = v15 + int32(112)
	return v340
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l0
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v61
	v64 = int32(1)
	if base.Ui32(l3-v64) < base.Ui32(int32(2)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+118)))
	v74 = int32(1)
	if base.Ui32(l3-v74) <= base.Ui32(v74) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v69 = int32(9)
	goto L21
L20:
	;
	v69 = v64
	goto L21
L21:
	;
	v70 = F_ExtendBufferedRel(m, v15, l1, l4, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v340 = v70
	goto L15
L23:
	;
	if v73 == int32(116) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+40)) = uint8(v73)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v53
	v327 = v15 + int32(32)
	if l3 == int32(3) {
		goto L85
	} else {
		goto L86
	}
L26:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v236 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L27:
	;
	v80 = int32(1)
	v83 = F_LocalBufferAlloc(m, v53, l1, l2, v15+int32(28))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v96 = F_IOContextForStrategy(m, l4)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L32
	}
L30:
	;
	v85 = int32(3)
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)))
	if v86 != int32(1) {
		v230 = v80
		v231 = v83
		v232 = v6
		v234 = v85
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v89 = int32(_a_F_ReadBufferExtended_3)
	v91 = *(*int64)(unsafe.Add(mBase, _c_F_ReadBufferExtended[0]))
	*(*int64)(unsafe.Add(mBase, _c_F_ReadBufferExtended[0])) = v91 + int64(1)
	v230 = v80
	v231 = v83
	v232 = int32(1)
	v234 = v85
	goto L26
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[8]))
	F_ResourceOwnerEnlarge(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v108
	v113 = v15 + int32(32)
	v114 = F_BufTableHashCode(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[9]))
	v124 = v117 + v114&int32(127)<<(uint(int32(7))%32) + int32(_a_F_ReadBufferExtended_5)
	v126 = F_LWLockAcquire(m, v124, int32(1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v128 = F_BufTableLookup(m, v113, v114)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L9
	} else {
		goto L38
	}
L37:
	;
	v221 = int32(_a_F_ReadBufferExtended_6)
	v223 = *(*int64)(unsafe.Add(mBase, _c_F_ReadBufferExtended[11]))
	*(*int64)(unsafe.Add(mBase, _c_F_ReadBufferExtended[11])) = v223 + int64(1)
	v230 = int32(0)
	v231 = v216
	v232 = int32(1)
	v234 = v96
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
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[10]))
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
	F_LWLockRelease(m, v124)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L9
	} else {
		goto L45
	}
L42:
	;
	F_LWLockRelease(m, v124)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)) = uint8(v137)
	if v137 != 0 {
		v216 = v136
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v230 = int32(0)
	v231 = v136
	v232 = v6
	v234 = v96
	goto L26
L45:
	;
	v145 = F_GetVictimBuffer(m, l4, v96)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[10]))
	v150 = F_LWLockAcquire(m, v124, int32(0))
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
	v162 = F_BufTableInsert(m, v15+int32(32), v114, v161)
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
	v200 = m.ExcPending
	if v200 != 0 {
		goto L9
	} else {
		goto L60
	}
L52:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v156)+16)) = v168
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v15)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v156)+8)) = v170
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v156))) = v172
	v174 = int32(0)
	v178 = base.AtomicRmwOr32(m, v174, int32(_a_F_ReadBufferExtended_7), v174)
	v183 = int32(-2113667072)
	if v73 == int32(112) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v188 = v183
	goto L55
L54:
	;
	v188 = int32(33816576)
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
	v191 = v183
	goto L58
L57:
	;
	v191 = v188
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154-int32(40)))) = v166&int32(-38010881) | v191
	F_LWLockRelease(m, v124)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	v196 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)) = uint8(v196)
	v230 = v196
	v231 = v156
	v232 = v174
	v234 = v96
	goto L26
L60:
	;
	F_StrategyFreeBuffer(m, v156)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[10]))
	v208 = v205 + v162<<(uint(int32(6))%32)
	v209 = F_PinBuffer(m, v208, l4)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	F_LWLockRelease(m, v124)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)) = uint8(v209)
	if v209 != 0 {
		v216 = v208
		goto L37
	} else {
		goto L64
	}
L64:
	;
	v230 = int32(0)
	v231 = v208
	v232 = int32(0)
	v234 = v96
	goto L26
L65:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	v317 = v315 + int32(1)
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)))
	F_ZeroAndLockBuffer(m, v317, l3, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L9
	} else {
		goto L84
	}
L66:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v257 != 0 {
		goto L76
	} else {
		goto L77
	}
L67:
	;
	if v232 == int32(0) {
		goto L65
	} else {
		goto L74
	}
L68:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v239 != int32(1) {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v236)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v236)+112)) = v250 + int64(1)
	goto L67
L71:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v245)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v245)+112)) = v246 + int64(1)
	if v244 != 0 {
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
	v282 = v230*int32(320) + v234<<(uint(int32(6))%32)
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_ReadBufferExtended[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_ReadBufferExtended[1]))) = v286 + int64(1)
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_ReadBufferExtended[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v282)+uint32(_c_F_ReadBufferExtended[2]))) = v290
	v292 = int32(1)
	F_pgstat_count_backend_io_op(m, v230, v234, int32(2), v292, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferExtended[3])) = uint8(v292)
	*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferExtended[4])) = uint8(v292)
	goto L82
L76:
	;
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v257)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v257)+120)) = v258 + int64(1)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v262 != int32(1) {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+28)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v268)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v268)+120)) = v269 + int64(1)
	if v267 != int32(1) {
		goto L65
	} else {
		goto L81
	}
L81:
	;
	goto L75
L82:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferExtended[5])))
	if v302 != int32(1) {
		goto L65
	} else {
		goto L83
	}
L83:
	;
	v305 = int32(_a_F_ReadBufferExtended_4)
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[6]))
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[6])) = v307 + v309
	goto L65
L84:
	;
	v340 = v317
	goto L15
L85:
	;
	v334 = int32(9)
	goto L87
L86:
	;
	v334 = int32(8)
	goto L87
L87:
	;
	v335 = F_StartReadBuffer(m, v327, v15+int32(28), l2, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	if v335 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F_WaitReadBuffers(m, v327)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L9
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v340 = v339
	goto L15
L92:
	;
	goto L91
L93:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_ReadBufferExtended_0), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ReadBufferExtended_1), int32(818), int32(_a_F_ReadBufferExtended_2))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
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
