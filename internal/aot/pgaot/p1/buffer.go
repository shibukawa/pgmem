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
	var v58 int64
	_ = v58
	var v60 int32
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
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int64
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int64
	_ = v239
	var v243 int64
	_ = v243
	var v250 int32
	_ = v250
	var v251 int64
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v275 int32
	_ = v275
	var v279 int64
	_ = v279
	var v283 int64
	_ = v283
	var v285 int32
	_ = v285
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	v6 = int32(0)
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
	v349 = m.ExcPending
	if v349 != 0 {
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
	return v333
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l0
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v60
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
	v333 = v69
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
	v320 = v14 + int32(32)
	if l3 == int32(3) {
		goto L85
	} else {
		goto L86
	}
L26:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v229 == int32(0) {
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
	v95 = F_IOContextForStrategy(m, l4)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L32
	}
L30:
	;
	v84 = int32(3)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)))
	if v85 != int32(1) {
		v224 = v79
		v225 = v82
		v226 = v6
		v228 = v84
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v88 = int32(_a_F_ReadBufferExtended_3)
	v90 = *(*int64)(unsafe.Add(mBase, _c_F_ReadBufferExtended[0]))
	*(*int64)(unsafe.Add(mBase, _c_F_ReadBufferExtended[0])) = v90 + int64(1)
	v224 = v79
	v225 = v82
	v226 = int32(1)
	v228 = v84
	goto L26
L32:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[8]))
	F_ResourceOwnerEnlarge(m, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v107
	v112 = v14 + int32(32)
	v113 = F_BufTableHashCode(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[9]))
	v123 = v116 + v113&int32(127)<<(uint(int32(7))%32) + int32(_a_F_ReadBufferExtended_5)
	v125 = F_LWLockAcquire(m, v123, int32(1))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v127 = F_BufTableLookup(m, v112, v113)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L9
	} else {
		goto L38
	}
L37:
	;
	v215 = int32(_a_F_ReadBufferExtended_6)
	v217 = *(*int64)(unsafe.Add(mBase, _c_F_ReadBufferExtended[11]))
	*(*int64)(unsafe.Add(mBase, _c_F_ReadBufferExtended[11])) = v217 + int64(1)
	v224 = int32(0)
	v225 = v211
	v226 = int32(1)
	v228 = v95
	goto L26
L38:
	;
	if int32(0) <= v127 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[10]))
	v135 = v132 + v127<<(uint(int32(6))%32)
	v136 = F_PinBuffer(m, v135, l4)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_LWLockRelease(m, v123)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L9
	} else {
		goto L45
	}
L42:
	;
	F_LWLockRelease(m, v123)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)) = uint8(v136)
	if v136 != 0 {
		v211 = v135
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v224 = int32(0)
	v225 = v135
	v226 = v6
	v228 = v95
	goto L26
L45:
	;
	v144 = F_GetVictimBuffer(m, l4, v95)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[10]))
	v149 = F_LWLockAcquire(m, v123, int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	v153 = v147 + v144<<(uint(int32(6))%32)
	v155 = v153 + int32(-64)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v153-int32(44))))
	v161 = F_BufTableInsert(m, v14+int32(32), v113, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	if v161 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v165 = F_LockBufHdr(m, v155)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L9
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_UnpinBuffer(m, v155)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L9
	} else {
		goto L60
	}
L52:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+16)) = v167
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v155)+8)) = v169
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v155))) = v171
	v177 = int32(-2113667072)
	if v72 == int32(112) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v182 = v177
	goto L55
L54:
	;
	v182 = int32(33816576)
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
	v185 = v177
	goto L58
L57:
	;
	v185 = v182
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153-int32(40)))) = v165&int32(-38010881) | v185
	F_LWLockRelease(m, v123)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)) = uint8(v190)
	v224 = v190
	v225 = v155
	v226 = v190
	v228 = v95
	goto L26
L60:
	;
	F_StrategyFreeBuffer(m, v155)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[10]))
	v203 = v200 + v161<<(uint(int32(6))%32)
	v204 = F_PinBuffer(m, v203, l4)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	F_LWLockRelease(m, v123)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)) = uint8(v204)
	if v204 != 0 {
		v211 = v203
		goto L37
	} else {
		goto L64
	}
L64:
	;
	v224 = int32(0)
	v225 = v203
	v226 = int32(0)
	v228 = v95
	goto L26
L65:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v225)+20))
	v310 = v308 + int32(1)
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)))
	F_ZeroAndLockBuffer(m, v310, l3, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L9
	} else {
		goto L84
	}
L66:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v250 != 0 {
		goto L76
	} else {
		goto L77
	}
L67:
	;
	if v226 == int32(0) {
		goto L65
	} else {
		goto L74
	}
L68:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v232 != int32(1) {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v229)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v229)+112)) = v243 + int64(1)
	goto L67
L71:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v238)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v238)+112)) = v239 + int64(1)
	if v237 != 0 {
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
	v275 = v224*int32(320) + v228<<(uint(int32(6))%32)
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_ReadBufferExtended[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_ReadBufferExtended[1]))) = v279 + int64(1)
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_ReadBufferExtended[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_ReadBufferExtended[2]))) = v283
	v285 = int32(1)
	F_pgstat_count_backend_io_op(m, v224, v228, int32(2), v285, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferExtended[3])) = uint8(v285)
	*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferExtended[4])) = uint8(v285)
	goto L82
L76:
	;
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v250)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v250)+120)) = v251 + int64(1)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v255 != int32(1) {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+28)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v261)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v261)+120)) = v262 + int64(1)
	if v260 != int32(1) {
		goto L65
	} else {
		goto L81
	}
L81:
	;
	goto L75
L82:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferExtended[5])))
	if v295 != int32(1) {
		goto L65
	} else {
		goto L83
	}
L83:
	;
	v298 = int32(_a_F_ReadBufferExtended_4)
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[6]))
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReadBufferExtended[6])) = v300 + v302
	goto L65
L84:
	;
	v333 = v310
	goto L15
L85:
	;
	v327 = int32(9)
	goto L87
L86:
	;
	v327 = int32(8)
	goto L87
L87:
	;
	v328 = F_StartReadBuffer(m, v320, v14+int32(28), l2, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	if v328 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F_WaitReadBuffers(m, v320)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L9
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v333 = v332
	goto L15
L92:
	;
	goto L91
L93:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_ReadBufferExtended_0), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ReadBufferExtended_1), int32(818), int32(_a_F_ReadBufferExtended_2))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
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
