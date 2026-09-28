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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
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
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
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
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	v2 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = l0 - v10 + int32(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 == v2 {
		v209 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v215 != 0 {
		goto L56
	} else {
		goto L57
	}
L2:
	;
	v18 = v14
	goto L4
L3:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v133 = l0 - v130 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v26)+19))
	v137 = v133 + v134<<(uint(int32(12))%32)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	if v139 != 0 {
		goto L34
	} else {
		goto L35
	}
L4:
	;
	v26 = v18 + v13
	v27 = int32(1)
	v28 = v26 - v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+3))
	if v29 != v27 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v119 = int32(129)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if base.Ui32(v119) <= base.Ui32(v120) {
		goto L31
	} else {
		goto L32
	}
L6:
	;
	goto L5
L7:
	;
	if v29 != int32(2) {
		v209 = v2
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v66 - int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v70 == int32(-1729435864) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v34 != int32(-1729435864) {
		v209 = v2
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+19))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+11))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v26)+15))
	v40 = v38 + v39
	v41 = int32(1)
	if base.B2i32(v37 != v40+v41)|base.B2i32(v40 != int32(base.Ui32(v18-v41)>>(uint(int32(12))%32))) != 0 {
		v209 = v2
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v53 = l0 - v50 + int32(1)
	v56 = v53 + v38<<(uint(int32(12))%32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v58 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53+v58)+7)) = v57
	goto L15
L14:
	;
	goto L15
L15:
	;
	if v57 == int32(0) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v57+v53)+11)) = v64
	goto L3
L17:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = int32(1)
	v89 = l0 - v86 + v88
	v93 = (v18 - v88) & int32(-4096)
	v94 = v89 + v93
	*(*int64)(unsafe.Add(mBase, uint32(v94))) = int64(8225038576)
	v97 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v97
	v99 = v85 + v89
	if v85 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v26)+11))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v26)+15))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v77
	goto L17
L19:
	;
	goto L20
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v26)+15))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v79+v13)+7)) = int32(0)
	goto L17
L21:
	;
	v103 = v99 - v88
	goto L23
L22:
	;
	v103 = v97
	goto L23
L23:
	;
	if v85 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v108 = v103 - v89 + int32(1)
	goto L26
L25:
	;
	v108 = int32(0)
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v108
	v111 = v93 | int32(1)
	if v85 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+7)) = v111
	goto L29
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v111
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v114 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v118 != 0 {
		v18 = v118
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v209 = v2
	goto L1
L31:
	;
	v123 = v119
	goto L33
L32:
	;
	v123 = v120
	goto L33
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v123<<(uint(int32(2))%32))+32)) = v127
	goto L3
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133+v139)+7)) = v138
	goto L36
L35:
	;
	goto L36
L36:
	;
	if v138 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v26)+11))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v156
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v26)+23))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v26)+15))
	v160 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v160
	v166 = v158 + v159 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v166
	v168 = int32(129)
	if base.Ui32(v168) <= base.Ui32(v166) {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v138+v133)+11)) = v143
	goto L37
L39:
	;
	goto L40
L40:
	;
	v145 = int32(129)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if base.Ui32(v145) <= base.Ui32(v146) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v149 = v145
	goto L43
L42:
	;
	v149 = v146
	goto L43
L43:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0+v149<<(uint(int32(2))%32))+32)) = v153
	goto L37
L44:
	;
	v171 = v168
	goto L46
L45:
	;
	v171 = v166
	goto L46
L46:
	;
	v176 = l0 + v171<<(uint(int32(2))%32) + int32(32)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v180 = int32(1)
	v181 = l0 - v178 + v180
	v183 = v156 << (uint(int32(12)) % 32)
	v184 = v181 + v183
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = int32(-364896016)
	v188 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v184)+8)) = v188
	v190 = v177 + v181
	if v177 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v194 = v190 - v180
	goto L49
L48:
	;
	v194 = v188
	goto L49
L49:
	;
	if v177 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v199 = v194 - v181 + int32(1)
	goto L52
L51:
	;
	v199 = int32(0)
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184)+12)) = v199
	v202 = v183 | int32(1)
	if v177 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+7)) = v202
	goto L55
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = v202
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v209 = v205
	goto L1
L56:
	;
	v218 = v215
	v219 = v209
	goto L59
L57:
	;
	v304 = v209
	goto L58
L58:
	;
	return v304
L59:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v229 = l0 - v226 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v231 = v229 + v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+11))
	if v232 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v304 = v299
	goto L58
L61:
	;
	v233 = v229 + v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231)+7))
	*(*int32)(unsafe.Add(mBase, uint32(v233)+7)) = v234
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v239 = v233 - int32(1)
	v241 = v238
	goto L63
L62:
	;
	v239 = int32(0)
	v241 = v218
	goto L63
L63:
	;
	v242 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v241 - v242
	if v232 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v249 = v239 - v229 + v242
	goto L66
L65:
	;
	v249 = int32(0)
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v249
	if v230 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v254 = v231 - int32(1)
	goto L69
L68:
	;
	v254 = int32(0)
	goto L69
L69:
	;
	v255 = v254 - v13
	v258 = int32(1)
	v260 = F_FreePageManagerPutInternal(m, l0, int32(base.Ui32(v255)>>(uint(int32(12))%32)), v258, v258)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	return int32(0)
L71:
	;
	if v260 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v269 = int32(1)
	v270 = l0 - v267 + v269
	v272 = v255 & int32(-4096)
	v273 = v270 + v272
	*(*int64)(unsafe.Add(mBase, uint32(v273))) = int64(8225038576)
	v276 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v273)+8)) = v276
	v278 = v270 + v266
	if v266 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	if base.Ui32(v219) < base.Ui32(v260) {
		goto L84
	} else {
		goto L85
	}
L75:
	;
	v282 = v278 - v269
	goto L77
L76:
	;
	v282 = v276
	goto L77
L77:
	;
	if v266 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v287 = v282 - v270 + int32(1)
	goto L80
L79:
	;
	v287 = int32(0)
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+12)) = v287
	v290 = v272 | int32(1)
	if v266 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+7)) = v290
	goto L83
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v290
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v293 + int32(1)
	return v219
L84:
	;
	v299 = v260
	goto L86
L85:
	;
	v299 = v219
	goto L86
L86:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v300 != 0 {
		v218 = v300
		v219 = v299
		goto L59
	} else {
		goto L87
	}
L87:
	;
	goto L60
}
func F_FreePageBtreeRemove(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 == int32(1) {
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
	v20 = v15 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v20
	if base.Ui32(v20) <= base.Ui32(l2) {
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
	v24 = l1 + int32(12)
	v27 = (v20 - l2) << (uint(int32(3)) % 32)
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v30 = v24 + l2<<(uint(int32(3))%32)
	base.MemoryCopy(m, v30, v30+int32(8), v27)
	goto L8
L7:
	;
	goto L8
L8:
	;
	if l2 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = l0 - v35 + int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v44 = l1
	goto L10
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v54 == int32(0) {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L4
L12:
	;
	v57 = v54 + v38
	v59 = v57 - int32(1)
	if v54 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v61 = v59
	goto L15
L14:
	;
	v61 = int32(0)
	goto L15
L15:
	;
	v63 = v57 + int32(11)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+3))
	v68 = int32(0)
	v71 = v65
	goto L16
L16:
	;
	if base.Ui32(v71) <= base.Ui32(v68) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if base.Ui32(v95) < base.Ui32(v65) {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	goto L17
L19:
	;
	v95 = v68
	goto L18
L20:
	;
	goto L21
L21:
	;
	v82 = int32(1)
	v83 = int32(base.Ui32(v68+v71) >> (uint(v82) % 32))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v63+v83<<(uint(int32(3))%32))))
	v90 = base.B2i32(base.Ui32(v39) < base.Ui32(v89))
	if base.Ui32(v39) < base.Ui32(v89) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v91 = v68
	goto L24
L23:
	;
	v91 = v83 + v82
	goto L24
L24:
	;
	if base.Ui32(v39) < base.Ui32(v89) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v92 = v83
	goto L27
L26:
	;
	v92 = v71
	goto L27
L27:
	;
	if v39 != v89 {
		v68 = v91
		v71 = v92
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v95 = v83
	goto L18
L29:
	;
	v101 = int32(0)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v59+v95<<(uint(int32(3))%32))+16))
	if v105 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v115 = int32(-1)
	goto L31
L31:
	;
	v116 = v115 + v95
	*(*int32)(unsafe.Add(mBase, uint32(v63+v116<<(uint(int32(3))%32)))) = v39
	if v116 == int32(0) {
		v44 = v61
		goto L10
	} else {
		goto L38
	}
L32:
	;
	v110 = v38 + v105 - int32(1)
	goto L34
L33:
	;
	v110 = v101
	goto L34
L34:
	;
	if v110 != v44 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v112 = int32(-1)
	goto L37
L36:
	;
	v112 = v101
	goto L37
L37:
	;
	v115 = v112
	goto L31
L38:
	;
	goto L11
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
	if base.Ui32(int32(_a_F_GetPageWithFreeSpace_0)) <= base.Ui32(l1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
			F_errmsg_internal(m, int32(_a_F_GetPageWithFreeSpace_1), v6)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_GetPageWithFreeSpace_2), int32(438), int32(_a_F_GetPageWithFreeSpace_3))
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
	var v36 int32
	_ = v36
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v522 int32
	_ = v522
	var v531 int32
	_ = v531
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	v4 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(_a_F_PageIndexMultiDelete_0)
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
	v642 = m.ExcPending
	if v642 != 0 {
		goto L12
	} else {
		goto L126
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L12
	} else {
		goto L123
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L12
	} else {
		goto L119
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L12
	} else {
		goto L115
	}
L5:
	;
	m.G0 = v23 + int32(_a_F_PageIndexMultiDelete_0)
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
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(base.Ui32(v59) < base.Ui32(int32(24)))|base.B2i32(base.Ui32(v62) < base.Ui32(v59))|(base.B2i32(base.Ui32(v65) < base.Ui32(v62))|base.B2i32(base.Ui32(int32(_a_F_PageIndexMultiDelete_1)) < base.Ui32(v65)))|base.B2i32((v65+int32(7))&int32(_a_F_PageIndexMultiDelete_2) != v65) != 0 {
		goto L4
	} else {
		goto L15
	}
L9:
	;
	v36 = v28
	goto L10
L10:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v36<<(uint(int32(1))%32)))))
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
	if v36 != 0 {
		v36 = v36 - int32(1)
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	if base.Ui32(v59) < base.Ui32(int32(25)) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v84 = int32(base.Ui32(v59+int32(_a_F_PageIndexMultiDelete_3))>>(uint(int32(2))%32)) & int32(_a_F_PageIndexMultiDelete_4)
	if v84 == int32(0) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v91 = int32(1)
	v98 = v23 + int32(1680)
	v100 = v4
	v103 = v65
	v104 = v4
	v106 = v91
	v107 = v91
	v108 = v4
	goto L18
L18:
	;
	v114 = v107 & int32(_a_F_PageIndexMultiDelete_4)
	v117 = l0 + int32(20) + v114<<(uint(int32(2))%32)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v120 = int32(base.Ui32(v118) >> (uint(int32(17)) % 32))
	v122 = v118 & int32(_a_F_PageIndexMultiDelete_5)
	if base.B2i32(base.Ui32(v122) < base.Ui32(v62))|base.B2i32(base.Ui32(v65) < base.Ui32(v122+v120))|base.B2i32(v122 != (v122+int32(7))&int32(_a_F_PageIndexMultiDelete_6)) != 0 {
		goto L3
	} else {
		goto L20
	}
L19:
	;
	if l2 != v168 {
		goto L2
	} else {
		goto L29
	}
L20:
	;
	if l2 <= v104 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v172 = v107 + int32(1)
	if base.Ui32(v172&int32(_a_F_PageIndexMultiDelete_4)) <= base.Ui32(v84) {
		v98 = v164
		v100 = v165
		v103 = v167
		v104 = v168
		v106 = v169
		v107 = v172
		v108 = v170
		goto L18
	} else {
		goto L28
	}
L22:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v98)+2)) = uint16(v122)
	*(*uint16)(unsafe.Add(mBase, uint32(v98))) = uint16(v100)
	v146 = (v120 + int32(7)) & int32(_a_F_PageIndexMultiDelete_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v98)+4)) = uint16(v146)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(48)+v100<<(uint(int32(2))%32)))) = v153
	if v103 < v122 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v104<<(uint(int32(1))%32)))))
	if v114 != v137 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v164 = v98
	v165 = v100
	v167 = v103
	v168 = v104 + int32(1)
	v169 = v106
	v170 = v108
	goto L21
L25:
	;
	v163 = v103
	goto L27
L26:
	;
	v163 = v122
	goto L27
L27:
	;
	v164 = v98 + int32(6)
	v165 = v100 + int32(1)
	v167 = v163
	v168 = v104
	v169 = base.B2i32(v122 < v103) & v106
	v170 = v146 + v108
	goto L21
L28:
	;
	goto L19
L29:
	;
	v177 = v65 - v59
	if base.Ui32(v177) < base.Ui32(v170) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v180 = v165 << (uint(int32(2)) % 32)
	if v180 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	base.MemoryCopy(m, l0+int32(24), v23+int32(48), v180)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v187 = v180 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v187)
	if int32(0) < v165 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v192 = v23 + int32(1680)
	v193 = int32(0)
	v202 = m.G0
	v204 = v202 + int32(-8192)
	m.G0 = v204
	if v169 != 0 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	goto L36
L36:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v65)
	goto L5
L37:
	;
	goto L5
L38:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v531)
	m.G0 = v204 - int32(-8192)
	goto L37
L39:
	;
	v206 = int32(1)
	if v165 <= v206 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v306) {
		goto L67
	} else {
		goto L68
	}
L42:
	;
	v209 = v206
	goto L44
L43:
	;
	v209 = v165
	goto L44
L44:
	;
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v214 = v210
	v217 = v193
	goto L46
L45:
	;
	v300 = v294 - v291
	if v300 == int32(0) {
		v531 = v290
		goto L38
	} else {
		goto L65
	}
L46:
	;
	v226 = v192 + v217*int32(6)
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226)+4)))
	v228 = int32(*(*int16)(unsafe.Add(mBase, uint32(v226)+2)))
	v229 = v227 + v228
	if v214 == v229 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v165 <= v217 {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	v231 = v214 - v227
	v233 = v217 + int32(1)
	if v233 != v209 {
		v214 = v231
		v217 = v233
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	v290 = v231
	v291 = v214
	v294 = v214
	goto L45
L52:
	;
	v290 = v214
	v291 = v229
	v294 = v229
	goto L45
L53:
	;
	goto L54
L54:
	;
	v241 = v214
	v242 = v229
	v244 = v217
	v245 = v229
	goto L55
L55:
	;
	v253 = v192 + v244*int32(6)
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253))))
	v261 = l0 + int32(20) + (v254+int32(1))&int32(_a_F_PageIndexMultiDelete_4)<<(uint(int32(2))%32)
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253)+4)))
	v263 = int32(*(*int16)(unsafe.Add(mBase, uint32(v253)+2)))
	if v262+v263 == v242 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v290 = v279
	v291 = v273
	v294 = v275
	goto L45
L57:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	v279 = v241 - v274
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v276&int32(-32768) | v279&int32(_a_F_PageIndexMultiDelete_5)
	v285 = v244 + int32(1)
	if v285 != v165 {
		v241 = v279
		v242 = v273
		v244 = v285
		v245 = v275
		goto L55
	} else {
		goto L64
	}
L58:
	;
	v273 = v263
	v274 = v262
	v275 = v245
	goto L57
L59:
	;
	goto L60
L60:
	;
	v266 = v245 - v242
	if v266 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	base.MemoryCopy(m, l0+v241, l0+v242, v266)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253)+4)))
	v271 = int32(*(*int16)(unsafe.Add(mBase, uint32(v253)+2)))
	v273 = v271
	v274 = v270
	v275 = v270 + v271
	goto L57
L64:
	;
	goto L56
L65:
	;
	base.MemoryCopy(m, l0+v290, l0+v291, v300)
	v531 = v290
	goto L38
L66:
	;
	if v165 <= v450 {
		goto L101
	} else {
		goto L102
	}
L67:
	;
	v316 = int32(base.Ui32(v306+int32(_a_F_PageIndexMultiDelete_3))>>(uint(int32(4))%32)) & int32(_a_F_PageIndexMultiDelete_7)
	goto L69
L68:
	;
	v316 = int32(0)
	goto L69
L69:
	;
	if v165 < v316 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if int32(2) <= v165 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L72
L72:
	;
	v406 = int32(1)
	if base.Ui32(v165) <= base.Ui32(v406) {
		goto L91
	} else {
		goto L92
	}
L73:
	;
	v401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192)+4)))
	v402 = int32(*(*int16)(unsafe.Add(mBase, uint32(v192)+2)))
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v447 = v404
	v449 = v401 + v402
	v450 = int32(0)
	goto L66
L74:
	;
	v321 = int32(1)
	if v165 <= v321 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	v368 = int32(0)
	goto L76
L76:
	;
	v380 = v192 + v368*int32(6)
	v381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v380)+4)))
	if v381 == int32(0) {
		goto L73
	} else {
		goto L90
	}
L77:
	;
	v324 = v321
	goto L79
L78:
	;
	v324 = v165
	goto L79
L79:
	;
	v333 = int32(0)
	v339 = v193
	goto L80
L80:
	;
	v345 = v192 + v333*int32(6)
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v345)+4)))
	if v346 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v324&int32(1) == int32(0) {
		goto L73
	} else {
		goto L89
	}
L82:
	;
	v347 = int32(*(*int16)(unsafe.Add(mBase, uint32(v345)+2)))
	base.MemoryCopy(m, v204+v347, l0+v347, v346)
	goto L84
L83:
	;
	goto L84
L84:
	;
	v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v345)+10)))
	if v352 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v353 = int32(*(*int16)(unsafe.Add(mBase, uint32(v345)+8)))
	base.MemoryCopy(m, v204+v353, l0+v353, v352)
	goto L87
L86:
	;
	goto L87
L87:
	;
	v358 = int32(2)
	v359 = v333 + v358
	v361 = v339 + v358
	if v361 != v324&int32(2147483646) {
		v333 = v359
		v339 = v361
		goto L80
	} else {
		goto L88
	}
L88:
	;
	goto L81
L89:
	;
	v368 = v359
	goto L76
L90:
	;
	v384 = int32(*(*int16)(unsafe.Add(mBase, uint32(v380)+2)))
	base.MemoryCopy(m, v204+v384, l0+v384, v381)
	goto L73
L91:
	;
	v409 = v406
	goto L93
L92:
	;
	v409 = v165
	goto L93
L93:
	;
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v414 = v410
	v417 = v193
	goto L95
L94:
	;
	v437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v438 = v435 - v437
	if v438 == int32(0) {
		v447 = v435
		v449 = v429
		v450 = v436
		goto L66
	} else {
		goto L99
	}
L95:
	;
	v426 = v192 + v417*int32(6)
	v427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v426)+4)))
	v428 = int32(*(*int16)(unsafe.Add(mBase, uint32(v426)+2)))
	v429 = v427 + v428
	if v414 != v429 {
		v435 = v414
		v436 = v417
		goto L94
	} else {
		goto L97
	}
L96:
	;
	v435 = v431
	v436 = v409
	goto L94
L97:
	;
	v431 = v414 - v427
	v433 = v417 + int32(1)
	if v433 != v409 {
		v414 = v431
		v417 = v433
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	base.MemoryCopy(m, v437+v204, l0+v437, v438)
	v447 = v435
	v449 = v429
	v450 = v436
	goto L66
L100:
	;
	v522 = v514 - v513
	if v522 == int32(0) {
		v531 = v512
		goto L38
	} else {
		goto L114
	}
L101:
	;
	v512 = v447
	v513 = v449
	v514 = v449
	goto L100
L102:
	;
	goto L103
L103:
	;
	v463 = v447
	v464 = v449
	v465 = v449
	v466 = v450
	goto L104
L104:
	;
	v475 = v192 + v466*int32(6)
	v476 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v475))))
	v483 = l0 + int32(20) + (v476+int32(1))&int32(_a_F_PageIndexMultiDelete_4)<<(uint(int32(2))%32)
	v484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v475)+4)))
	v485 = int32(*(*int16)(unsafe.Add(mBase, uint32(v475)+2)))
	if v484+v485 == v464 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v512 = v501
	v513 = v495
	v514 = v496
	goto L100
L106:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	v501 = v463 - v497
	*(*int32)(unsafe.Add(mBase, uint32(v483))) = v498&int32(-32768) | v501&int32(_a_F_PageIndexMultiDelete_5)
	v507 = v466 + int32(1)
	if v507 != v165 {
		v463 = v501
		v464 = v495
		v465 = v496
		v466 = v507
		goto L104
	} else {
		goto L113
	}
L107:
	;
	v495 = v485
	v496 = v465
	v497 = v484
	goto L106
L108:
	;
	goto L109
L109:
	;
	v488 = v465 - v464
	if v488 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	base.MemoryCopy(m, l0+v463, v464+v204, v488)
	goto L112
L111:
	;
	goto L112
L112:
	;
	v492 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v475)+4)))
	v493 = int32(*(*int16)(unsafe.Add(mBase, uint32(v475)+2)))
	v495 = v493
	v496 = v492 + v493
	v497 = v492
	goto L106
L113:
	;
	goto L105
L114:
	;
	base.MemoryCopy(m, l0+v512, v513+v204, v522)
	v531 = v512
	goto L38
L115:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L12
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v59
	F_errmsg(m, int32(_a_F_PageIndexMultiDelete_8), v23)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L12
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_PageIndexMultiDelete_9), int32(1211), int32(_a_F_PageIndexMultiDelete_10))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L12
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L12
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v122
	F_errmsg(m, int32(_a_F_PageIndexMultiDelete_11), v23+int32(16))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L12
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_PageIndexMultiDelete_9), int32(1236), int32(_a_F_PageIndexMultiDelete_10))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L12
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errmsg_internal(m, int32(_a_F_PageIndexMultiDelete_12), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L12
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_PageIndexMultiDelete_9), int32(1263), int32(_a_F_PageIndexMultiDelete_10))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L12
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L12
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v170
	F_errmsg(m, int32(_a_F_PageIndexMultiDelete_13), v23+int32(32))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L12
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_PageIndexMultiDelete_9), int32(1269), int32(_a_F_PageIndexMultiDelete_10))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L12
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PageInit(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	v6 = int32(3)
	if l1&v6|(l0&v6|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(l1))) == int32(0) {
		if l1 == int32(0) {
		} else {
			v20 = l0 + l1
			v22 = l0 + int32(4)
			if base.Ui32(v22) < base.Ui32(v20) {
				v24 = v20
			} else {
				v24 = v22
			}
			v29 = (l0^int32(-1)+v24)&int32(-4) + int32(4)
			if v29 == int32(0) {
			} else {
				base.MemoryFill(m, l0, int32(0), v29)
			}
		}
	} else {
		if l1 == int32(0) {
		} else {
			base.MemoryFill(m, l0, int32(0), l1)
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_PageInit_0)
	v43 = l1 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v43)
	v49 = l1 - (l2+int32(7))&int32(_a_F_PageInit_1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v49)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v49)
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
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
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v233 int32
	_ = v233
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
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	if v17 == int32(0) {
		v67 = v5
		v68 = v5
		v69 = v5
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v13 + int32(16)
	return v233
L5:
	;
	v70 = int32(0)
	v73 = (v70 - l0) & int32(3)
	if v73 == v70 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_PageIsVerified[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+268))
	goto L8
L7:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	if base.Ui32(int32(7)) < base.Ui32(v43) {
		v67 = v41
		v68 = v42
		v69 = v5
		goto L5
	} else {
		goto L13
	}
L8:
	;
	if base.B2i32(v22 != int32(0)) == int32(0) {
		v41 = v5
		v42 = v5
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v27 = F_pg_checksum_page(m, l0, l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	if base.B2i32(l3 == int32(0))|base.B2i32(v31 == v27) != 0 {
		v41 = base.B2i32(v27 != v31)
		v42 = v27
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v37)
	v41 = v37
	v42 = v27
	goto L7
L13:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v46) < base.Ui32(v47) {
		v67 = v41
		v68 = v42
		v69 = v5
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(base.Ui32(v49) < base.Ui32(v46))|base.B2i32(base.Ui32(int32(_a_F_PageIsVerified_0)) < base.Ui32(v49)) != 0 {
		v67 = v41
		v68 = v42
		v69 = v5
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v58 = (v49 + int32(7)) & int32(_a_F_PageIsVerified_1)
	if v41|base.B2i32(v58 != v49) == int32(0) {
		v233 = int32(1)
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v67 = v41
	v68 = v42
	v69 = base.B2i32(v49 == v58)
	goto L5
L17:
	;
	if v67 != 0 {
		goto L46
	} else {
		goto L47
	}
L18:
	;
	v86 = l0 + v73
	v90 = (l0 - int32(-8192)) & int32(-4)
	v92 = v90 - int32(28)
	if base.Ui32(v86) < base.Ui32(v92) {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v76 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	if v73 == int32(1) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v79 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	if v73 == int32(2) {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if v82|base.B2i32(v73 != int32(3)) != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	v95 = v73
	v97 = v86
	goto L28
L26:
	;
	v124 = v73
	goto L27
L27:
	;
	v133 = l0 + v124
	if base.Ui32(v133) < base.Ui32(v90) {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v97)+28))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)+24))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v104|(v105|(v106|(v107|(v108|(v109|(v110|v111)))))) != 0 {
		goto L17
	} else {
		goto L30
	}
L29:
	;
	v124 = v120
	goto L27
L30:
	;
	v120 = v95 + int32(32)
	v121 = l0 + v120
	if base.Ui32(v121) < base.Ui32(v92) {
		v95 = v120
		v97 = v121
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v136 = v124
	v138 = v133
	goto L35
L33:
	;
	v151 = v124
	goto L34
L34:
	;
	v160 = int32(_a_F_PageIsVerified_0)
	if base.Ui32(v151) <= base.Ui32(v160) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if v145 != 0 {
		goto L17
	} else {
		goto L37
	}
L36:
	;
	v151 = v147
	goto L34
L37:
	;
	v147 = v136 + int32(4)
	v148 = l0 + v147
	if base.Ui32(v148) < base.Ui32(v90) {
		v136 = v147
		v138 = v148
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v163 = v160
	goto L41
L40:
	;
	v163 = v151
	goto L41
L41:
	;
	v166 = v151
	goto L42
L42:
	;
	if v166 == v163 {
		v233 = int32(1)
		goto L4
	} else {
		goto L44
	}
L43:
	;
	goto L17
L44:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v166))))
	if v179 == int32(0) {
		v166 = v166 + int32(1)
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
	v233 = int32(0)
	goto L4
L49:
	;
	if int32(base.Ui32(l2)>>(uint(int32(2))%32))&v69 != 0 {
		v233 = int32(1)
		goto L4
	} else {
		goto L62
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
	v203 = m.ExcPending
	if v203 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	if v202 == int32(0) {
		goto L49
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v68
	if l2&int32(8) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v216 = int32(_a_F_PageIsVerified_2)
	goto L59
L58:
	;
	v216 = int32(_a_F_PageIsVerified_3)
	goto L59
L59:
	;
	F_errmsg(m, v216, v13)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_PageIsVerified_4), int32(158), int32(_a_F_PageIsVerified_5))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	goto L49
L62:
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	v2 = int32(0)
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v8) < base.Ui32(int32(25)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v75 = v73 & int32(_a_F_PageTruncateLinePointerArray_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v75)
	return
L2:
	;
	v16 = int32(base.Ui32(v8+int32(_a_F_PageTruncateLinePointerArray_1))>>(uint(int32(2))%32)) & int32(_a_F_PageTruncateLinePointerArray_2)
	if v16 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = v16
	v24 = v2
	v27 = v2
	goto L5
L4:
	;
	if int32(0) < v51 {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v22<<(uint(int32(2))%32))))
	v33 = v31 & int32(_a_F_PageTruncateLinePointerArray_3)
	if base.B2i32(v22 == int32(1))|v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v51 = v45
	v53 = int32(0)
	goto L4
L7:
	;
	v48 = v22 - int32(1)
	if v48 != 0 {
		v22 = v48
		v24 = v45
		v27 = v46
		goto L5
	} else {
		goto L12
	}
L8:
	;
	v39 = int32(0)
	v45 = v24 + base.B2i32(v33 == v39)
	v46 = base.B2i32(v33 != v39)
	goto L7
L9:
	;
	goto L10
L10:
	;
	if v33 != 0 {
		v45 = v24
		v46 = v27
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v51 = v24
	v53 = int32(1)
	goto L4
L12:
	;
	goto L6
L13:
	;
	v58 = v8 - v51<<(uint(int32(2))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v58)
	goto L15
L14:
	;
	goto L15
L15:
	;
	if v53 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v64 = v62 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v64)
	return
}
func F_UnlockPage(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(16973824)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v11
	v17 = F_LockRelease(m, v7, l2, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
