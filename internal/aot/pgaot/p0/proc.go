package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ProcArrayEndTransaction(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v86 int64
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
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
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int64
	_ = v235
	var v236 int32
	_ = v236
	var v250 int64
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
	F_LWLockRelease(m, v324+int32(512))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L6
	} else {
		goto L76
	}
L3:
	;
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
	v11 = F_LWLockConditionalAcquire(m, v7+int32(512), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v290 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+336)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v290
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v296&int32(14) == v290 {
		goto L1
	} else {
		goto L74
	}
L6:
	;
	return
L7:
	;
	if v11 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v13 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v18+v19<<(uint(int32(2))%32)))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(l0)+336)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = int64(0)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v31&int32(14) != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+612)) = l1
	v94 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+604)) = uint8(v94)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+608)) = v96
	v100 = base.I32_div_s(l0-v92, int32(768))
	v102 = base.AtomicRmwCmpxchg32(m, v91, int32(56), v96, v100)
	if v102 != v96 {
		goto L27
	} else {
		goto L28
	}
L11:
	;
	goto L2
L12:
	;
	v35 = v31 & int32(241)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v35)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v39+v40))) = uint8(v35)
	goto L14
L13:
	;
	goto L14
L14:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v44 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v67 = int32(3)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[2]))
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)+48))
	v72 = base.I32_wrap_i64(v71)
	if base.B2i32(base.Ui32(l1) < base.Ui32(v67))|base.B2i32(base.Ui32(v72) < base.Ui32(v67)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+57)))
	if v47 != int32(1) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v51 = v19 << (uint(int32(1)) % 32)
	v52 = int32(_a_F_ProcArrayEndTransaction_0)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v56 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51+v54))) = uint8(v56)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v60+v51)+1)) = uint8(v56)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)) = uint16(v56)
	goto L15
L19:
	;
	goto L18
L20:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v70)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v70)+56)) = v86 + int64(1)
	goto L11
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v70)+48)) = v71 + base.I64_extend_i32_s(l1-v72)
	goto L20
L22:
	;
	if v72-l1 < int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(l1) <= base.Ui32(v72) {
		goto L20
	} else {
		goto L26
	}
L25:
	;
	goto L20
L26:
	;
	goto L21
L27:
	;
	v106 = v102
	goto L30
L28:
	;
	v114 = v96
	goto L29
L29:
	;
	if v114 != int32(-1) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+608)) = v106
	v111 = base.AtomicRmwCmpxchg32(m, v91, int32(56), v106, v100)
	if v106 != v111 {
		v106 = v111
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v114 = v106
	goto L29
L32:
	;
	goto L31
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = int32(134217769)
	v127 = int32(0)
	goto L36
L34:
	;
	goto L35
L35:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
	v159 = F_LWLockAcquire(m, v155+int32(512), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L45
	}
L36:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	F_PGSemaphoreLock(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L38
	}
L37:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[3]))
	v138 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v138
	if v127 <= v138 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+604)))
	if v135 != 0 {
		v127 = v127 + int32(1)
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v143 = v127
	goto L41
L41:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	F_PGSemaphoreUnlock(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L6
	} else {
		goto L43
	}
L42:
	;
	goto L1
L43:
	;
	v150 = int32(1)
	if base.Ui32(v150) < base.Ui32(v143) {
		v143 = v143 - v150
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v161 = int32(-1)
	v163 = base.AtomicRmwXchg32(m, v91, int32(56), v161)
	if v163 == v161 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[4]))
	v168 = v163
	goto L47
L47:
	;
	v175 = v167 + v168*int32(768)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+612))
	v177 = int32(0)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v175)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v182+v183<<(uint(int32(2))%32)))) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v175)+336)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v175)+52)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v175)+44)) = int64(0)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+36)))
	if v195&int32(14) != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
	F_LWLockRelease(m, v258+int32(512))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L6
	} else {
		goto L66
	}
L49:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v175)+608))
	if v254 != int32(-1) {
		v168 = v254
		goto L47
	} else {
		goto L65
	}
L50:
	;
	v199 = v195 & int32(241)
	*(*uint8)(unsafe.Add(mBase, uint32(v175)+36)) = uint8(v199)
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v175)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v203+v204))) = uint8(v199)
	goto L52
L51:
	;
	goto L52
L52:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+56)))
	if v208 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v231 = int32(3)
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[2]))
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v234)+48))
	v236 = base.I32_wrap_i64(v235)
	if base.B2i32(base.Ui32(v176) < base.Ui32(v231))|base.B2i32(base.Ui32(v236) < base.Ui32(v231)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+57)))
	if v211 != int32(1) {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v215 = v183 << (uint(int32(1)) % 32)
	v216 = int32(_a_F_ProcArrayEndTransaction_0)
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v220 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v215+v218))) = uint8(v220)
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v224+v215)+1)) = uint8(v220)
	*(*uint16)(unsafe.Add(mBase, uint32(v175)+56)) = uint16(v220)
	goto L53
L57:
	;
	goto L56
L58:
	;
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v234)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v234)+56)) = v250 + int64(1)
	goto L49
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v234)+48)) = v235 + base.I64_extend_i32_s(v176-v236)
	goto L58
L60:
	;
	if v236-v176 < int32(0) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	if base.Ui32(v176) <= base.Ui32(v236) {
		goto L58
	} else {
		goto L64
	}
L63:
	;
	goto L58
L64:
	;
	goto L59
L65:
	;
	goto L48
L66:
	;
	v264 = v163
	goto L67
L67:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[4]))
	v272 = v269 + v264*int32(768)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+608))
	*(*int32)(unsafe.Add(mBase, uint32(v272)+608)) = int32(-1)
	v276 = int32(0)
	v279 = base.AtomicRmwOr32(m, v276, int32(_a_F_ProcArrayEndTransaction_1), v276)
	*(*uint8)(unsafe.Add(mBase, uint32(v272)+604)) = uint8(v276)
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[5]))
	if v283 != v272 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L1
L69:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v272)+332))
	F_PGSemaphoreUnlock(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L6
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v273 != int32(-1) {
		v264 = v273
		goto L67
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	goto L68
L74:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
	v306 = F_LWLockAcquire(m, v302+int32(512), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	v310 = v308 & int32(-15)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v310)
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+12))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v314+v315))) = uint8(v310)
	goto L2
L76:
	;
	goto L1
}
func F_ProcLockWakeup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v95 int64
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	v3 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v10 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = l1 + int32(32)
	if v13 == v17 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = v13
	v24 = v3
	goto L5
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29+v30<<(uint(int32(2))%32))))
	if v34&v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L1
L7:
	;
	if v28 != v17 {
		v23 = v28
		v24 = v153
		goto L5
	} else {
		goto L32
	}
L8:
	;
	v153 = int32(1)<<(uint(v30)%32) | v24
	goto L7
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v37 = F_LockCheckConflicts(m, l0, v30, l1, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v37 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	v43 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+128)) = v42 + v43
	v48 = l1 + v30<<(uint(int32(2))%32)
	v50 = v48 + int32(88)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v51 + v43
	v56 = v43 << (uint(v30) % 32)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v56 | v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+44))
	if v60 == v61 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v71 == int32(0) {
		v153 = v24
		goto L7
	} else {
		goto L17
	}
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v63 & (v56 ^ int32(-1))
	goto L16
L15:
	;
	goto L16
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v68 | v56
	goto L13
L17:
	;
	v75 = v23 - int32(4)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v71
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v79
	v81 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v76)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+40)) = v83 - int32(1)
	v87 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v87
	v95 = base.AtomicRmwXchg64(m, v23, int32(20), v81)
	v97 = v23 - int32(72)
	v101 = base.AtomicRmwOr32(m, v87, int32(_a_F_ProcLockWakeup_0), v87)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v102 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v153 = v24
	goto L7
L19:
	;
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(1)
	v105 = int32(0)
	v108 = base.AtomicRmwOr32(m, v105, int32(_a_F_ProcLockWakeup_0), v105)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	if v109 == v105 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	if v112 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_ProcLockWakeup[0]))
	if v116 == v112 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v118 = m.G0
	v120 = v118 - int32(16)
	m.G0 = v120
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_ProcLockWakeup[1]))
	if v123 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v146 = F_pgmem_kill(m, v112, int32(23))
	mBase = m.M
	goto L19
L26:
	;
	m.G0 = v120 + int32(16)
	goto L18
L27:
	;
	v126 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+15)) = uint8(v126)
	goto L28
L28:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcLockWakeup[2]))
	v134 = F_write(m, v130, v120+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v134 {
		goto L26
	} else {
		goto L30
	}
L29:
	;
	goto L26
L30:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_ProcLockWakeup[3]))
	if v138 == int32(27) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L6
}
func F_ProcSendSignal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	if int32(0) <= l0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v28 = v23 + l0*int32(768) + int32(316)
	v29 = int32(0)
	v32 = base.AtomicRmwOr32(m, v29, int32(_a_F_ProcSendSignal_0), v29)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v33 != 0 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSendSignal[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if base.Ui32(l0) < base.Ui32(v7) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	return
L7:
	;
	F_errmsg_internal(m, int32(_a_F_ProcSendSignal_1), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_errfinish(m, int32(_a_F_ProcSendSignal_2), int32(2058), int32(_a_F_ProcSendSignal_3))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	return
L11:
	;
	goto L10
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(1)
	v36 = int32(0)
	v39 = base.AtomicRmwOr32(m, v36, int32(_a_F_ProcSendSignal_0), v36)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v40 == v36 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v43 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSendSignal[1]))
	if v47 == v43 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v49 = m.G0
	v51 = v49 - int32(16)
	m.G0 = v51
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSendSignal[2]))
	if v54 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v77 = F_pgmem_kill(m, v43, int32(23))
	mBase = m.M
	goto L11
L18:
	;
	m.G0 = v51 + int32(16)
	goto L10
L19:
	;
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51)+15)) = uint8(v57)
	goto L20
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSendSignal[3]))
	v65 = F_write(m, v61, v51+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v65 {
		goto L18
	} else {
		goto L22
	}
L21:
	;
	goto L18
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSendSignal[4]))
	if v69 == int32(27) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
}
func F_ProcSignalShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int64
	_ = v25
	var v29 int64
	_ = v29
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalShmemInit[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalShmemInit[1]))
	if v2 < v8+int32(38) {
		v14 = v2
		for {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalShmemInit[0]))
			v19 = v16 + v14*int32(112)
			v20 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v19)+88)), uint32(v20))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v20
			v25 = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v20
			v29 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v29
			*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v29
			*(*int64)(unsafe.Add(mBase, uint32(v19-int32(-64)))) = v29
			*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v29
			*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v29
			*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v20
			v44 = v19 + int32(108)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v44))), uint32(v20))
			*(*int64)(unsafe.Add(mBase, uint32(v44)+4)) = v25
			v51 = v14 + int32(1)
			v53 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSignalShmemInit[1]))
			if v51 < v53+int32(38) {
				v14 = v51
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_ProcWaitForSignal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ProcWaitForSignal[0]))
	v6 = F_WaitLatch(m, v3, int32(33), int32(0), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_ProcWaitForSignal[0]))
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v10
		v15 = base.AtomicRmwOr32(m, v10, int32(_a_F_ProcWaitForSignal_0), v10)
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_ProcWaitForSignal[1]))
		if v17 != 0 {
			F_ProcessInterrupts(m)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_ProcessProcSignalBarrier(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
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
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	v1 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(192)
	m.G0 = v9
	v13 = v1
	v14 = v1
	v15 = v1
	v16 = int32(-1)
	v17 = int64(0)
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v16 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	goto L2
L5:
	;
	v191 = int32(m.ExcTag)
	v192 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v191 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0])) = v61
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[1])) = v63
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[2]))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	v173 = base.AtomicRmwOr32(m, v170, int32(96), v171)
	v175 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[3])) = v175
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[4])) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v9)+184)) = v64
	F_pg_re_throw(m)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L5
	} else {
		goto L39
	}
L7:
	;
	m.G0 = v9 + int32(192)
	return
L8:
	;
	v143 = int32(_a_F_ProcessProcSignalBarrier_0)
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[2]))
	v146 = base.AtomicRmwXchg64(m, v144, int32(88), v142)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v9)+184)) = v142
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[2]))
	F_ConditionVariableBroadcast(m, v151+int32(100))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L38
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[3]))
	if v21 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	v61 = v13
	v62 = v14
	v63 = v15
	v64 = v17
	goto L11
L11:
	;
	if v62 != 0 {
		goto L6
	} else {
		goto L19
	}
L12:
	;
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[3])) = v25
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[2]))
	v29 = int64(0)
	v32 = base.AtomicRmwCmpxchg64(m, v28, int32(88), v29, v29)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[5]))
	v38 = base.AtomicRmwCmpxchg64(m, v34, v25, v29, v29)
	if v32 == v38 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[2]))
	v42 = int32(0)
	v44 = base.AtomicRmwXchg32(m, v41, int32(96), v42)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v44
	if v44 == v42 {
		v138 = v13
		v140 = v15
		v142 = v38
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0]))
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[1]))
	goto L15
L15:
	;
	v55 = v9 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v9 + int32(12)
	goto L18
L16:
	;
	v61 = v51
	v62 = int32(0)
	v63 = v53
	v64 = v38
	goto L11
L18:
	;
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[1])) = v9 + int32(16)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	if v70 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	goto L23
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0])) = v61
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[1])) = v63
	v138 = v61
	v140 = v63
	v142 = v64
	goto L8
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	v78 = base.I32_ctz(v77)
	switch v78 {
	case 0:
		goto L27
	case 1:
		goto L28
	default:
		goto L26
	}
L25:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	if v126 != 0 {
		goto L23
	} else {
		goto L36
	}
L26:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v120 & base.I32_rotl(int32(-2), v78)
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v9)+184)) = v64
	F_smgrreleaseall(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L35
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v9)+184)) = v64
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[6]))
	if v83 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v107 & int32(-3)
	goto L25
L30:
	;
	v85 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[7])) = uint8(v85)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[8]))
	v92 = F_LWLockAcquire(m, v88+int32(_a_F_ProcessProcSignalBarrier_1), int32(1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[9]))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[8]))
	F_LWLockRelease(m, v98+int32(_a_F_ProcessProcSignalBarrier_1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[10])) = uint8(v96)
	goto L29
L35:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v116 & int32(-2)
	goto L25
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0])) = v61
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[1])) = v63
	v138 = v61
	v140 = v63
	v142 = v64
	goto L8
L38:
	;
	goto L7
L39:
	;
	goto L4
L40:
	;
	v196 = int32(v192)
	m.G0 = v9
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	if v9+int32(12) == v202 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	m.ExcPending = 1
	goto L49
L42:
	;
	if v206 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	v206 = v204
	goto L45
L44:
	;
	v206 = int32(0)
	goto L45
L45:
	;
	goto L42
L46:
	;
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v9)+184))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v9)+180))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v9)+176))
	v13 = v209
	v14 = v198
	v15 = v208
	v16 = v206
	v17 = v207
	goto L1
L47:
	;
	goto L48
L48:
	;
	F___wasm_longjmp(m, v199, v198)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SendProcSignal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	if l2 != int32(-1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[0])) = int32(71)
	return int32(-1)
L2:
	;
	v92 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16))), uint32(v92))
	goto L1
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[1]))
	v12 = v9 + l2*int32(112)
	v14 = v12 + int32(8)
	v16 = v12 + int32(88)
	v19 = base.AtomicRmwXchg32(m, v16, int32(0), int32(1))
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[2]))
	v41 = v39 + int32(37)
	if v41 < int32(0) {
		goto L1
	} else {
		goto L12
	}
L6:
	;
	F_s_lock(m, v16, int32(_a_F_SendProcSignal_0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v25 != l0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+l1<<(uint(int32(2))%32))+40)) = int32(1)
	v32 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14)+80)), uint32(v32))
	v36 = F_pgmem_kill(m, l0, int32(10))
	mBase = m.M
	return v36
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[1]))
	v48 = v41
	v49 = v45
	goto L13
L13:
	;
	v53 = v49 + v48*int32(112)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if l0 == v54 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57+l1<<(uint(int32(2))%32))+40)) = int32(1)
	v86 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v57)+80)), uint32(v86))
	v90 = F_pgmem_kill(m, l0, int32(10))
	mBase = m.M
	return v90
L15:
	;
	goto L14
L16:
	;
	v57 = v53 + int32(8)
	v59 = v53 + int32(88)
	v62 = base.AtomicRmwXchg32(m, v57, int32(80), int32(1))
	if v62 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v73 = v49
	goto L18
L18:
	;
	v75 = int32(0)
	if base.B2i32(v48 <= v75) == v75 {
		v48 = v48 - int32(1)
		v49 = v73
		goto L13
	} else {
		goto L24
	}
L19:
	;
	F_s_lock(m, v59, int32(_a_F_SendProcSignal_0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v66 == l0 {
		goto L15
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v68 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v59))), uint32(v68))
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[1]))
	v73 = v72
	goto L18
L24:
	;
	goto L1
}
func F_assignProcTypes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int64
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
	v14 = F_SearchSysCache1(m, int32(47), v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v18 = v16 + v17
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if l3 == v19 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if l2 != 0 {
					if v21 != l2 {
						v24 = v21
					} else {
						v24 = int32(0)
					}
					if v24 == int32(0) {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if base.B2i32(v27 == int32(0))|base.B2i32(v27 == l2) != 0 {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
							if v52 != int32(2278) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_0), int32(0))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_assignProcTypes_1)
											F_errhint(m, int32(_a_F_assignProcTypes_2), v10+int32(16))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1259), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
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
							} else {
								v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
								if v55 != int32(1) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_0), int32(0))
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_assignProcTypes_1)
												F_errhint(m, int32(_a_F_assignProcTypes_2), v10+int32(16))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1259), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
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
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
									if v58 == int32(2281) {
										v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v245 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
											v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v249 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											} else {
											}
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
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
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										} else {
											v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v252 != 0 {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
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
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_0), int32(0))
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_assignProcTypes_1)
													F_errhint(m, int32(_a_F_assignProcTypes_2), v10+int32(16))
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1259), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
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
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_assignProcTypes_6), int32(0))
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1242), int32(_a_F_assignProcTypes_4))
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_assignProcTypes_6), int32(0))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1242), int32(_a_F_assignProcTypes_4))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
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
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v21 != v49 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v278 = m.ExcPending
						if v278 != 0 {
							return
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v281 = m.ExcPending
							if v281 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_assignProcTypes_7), int32(0))
								mBase = m.M
								v285 = m.ExcPending
								if v285 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1249), int32(_a_F_assignProcTypes_4))
									mBase = m.M
									v290 = m.ExcPending
									if v290 != 0 {
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
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
						if v52 != int32(2278) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_assignProcTypes_0), int32(0))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_assignProcTypes_1)
										F_errhint(m, int32(_a_F_assignProcTypes_2), v10+int32(16))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1259), int32(_a_F_assignProcTypes_4))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
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
						} else {
							v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v55 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_0), int32(0))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_assignProcTypes_1)
											F_errhint(m, int32(_a_F_assignProcTypes_2), v10+int32(16))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1259), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
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
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
								if v58 == int32(2281) {
									v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v245 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
										v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v249 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										} else {
										}
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
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
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									} else {
										v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v252 != 0 {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
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
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										}
									}
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_0), int32(0))
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_assignProcTypes_1)
												F_errhint(m, int32(_a_F_assignProcTypes_2), v10+int32(16))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1259), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
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
			} else {
				v85 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+10)))
					if v87 == int32(1) {
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						switch v90 - int32(1) {
						case 0:
							v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v93 != int32(2) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v294 = m.ExcPending
								if v294 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v297 = m.ExcPending
									if v297 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_8), int32(0))
										mBase = m.M
										v301 = m.ExcPending
										if v301 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1277), int32(_a_F_assignProcTypes_4))
											mBase = m.M
											v306 = m.ExcPending
											if v306 != 0 {
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
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
								if v96 != int32(23) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v310 = m.ExcPending
									if v310 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v313 = m.ExcPending
										if v313 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_9), int32(0))
											mBase = m.M
											v317 = m.ExcPending
											if v317 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1281), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v322 = m.ExcPending
												if v322 != 0 {
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
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v99 == int32(0) {
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v102
									} else {
									}
									v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v104 != 0 {
									} else {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v18)+140))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v105
									}
									v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v245 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
										v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v249 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										} else {
										}
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
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
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									} else {
										v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v252 != 0 {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
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
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						case 1:
							v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v107 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v326 = m.ExcPending
								if v326 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v329 = m.ExcPending
									if v329 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_10), int32(0))
										mBase = m.M
										v333 = m.ExcPending
										if v333 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1298), int32(_a_F_assignProcTypes_4))
											mBase = m.M
											v338 = m.ExcPending
											if v338 != 0 {
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
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
								if v110 != int32(2281) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v326 = m.ExcPending
									if v326 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v329 = m.ExcPending
										if v329 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_10), int32(0))
											mBase = m.M
											v333 = m.ExcPending
											if v333 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1298), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v338 = m.ExcPending
												if v338 != 0 {
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
									v113 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
									if v113 == int32(2278) {
										v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v245 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
											v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v249 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											} else {
											}
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
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
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										} else {
											v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v252 != 0 {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
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
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_11), int32(0))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1302), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
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
						case 2:
							v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v132 != int32(5) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v342 = m.ExcPending
								if v342 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v345 = m.ExcPending
									if v345 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_12), int32(0))
										mBase = m.M
										v349 = m.ExcPending
										if v349 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1313), int32(_a_F_assignProcTypes_4))
											mBase = m.M
											v354 = m.ExcPending
											if v354 != 0 {
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
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
								if v135 != int32(16) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v358 = m.ExcPending
									if v358 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v361 = m.ExcPending
										if v361 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_13), int32(0))
											mBase = m.M
											v365 = m.ExcPending
											if v365 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1317), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v370 = m.ExcPending
												if v370 != 0 {
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
									v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v138 == int32(0) {
										v141 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v141
									} else {
									}
									v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v143 != 0 {
									} else {
										v144 = *(*int32)(unsafe.Add(mBase, uint32(v18)+144))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v144
									}
									v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v245 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
										v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v249 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										} else {
										}
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
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
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									} else {
										v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v252 != 0 {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
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
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										}
									}
								}
							}
						case 3:
							v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v146 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v374 = m.ExcPending
								if v374 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v377 = m.ExcPending
									if v377 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_14), int32(0))
										mBase = m.M
										v381 = m.ExcPending
										if v381 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1333), int32(_a_F_assignProcTypes_4))
											mBase = m.M
											v386 = m.ExcPending
											if v386 != 0 {
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
								v149 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
								if v149 != int32(16) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v390 = m.ExcPending
									if v390 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v393 = m.ExcPending
										if v393 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_15), int32(0))
											mBase = m.M
											v397 = m.ExcPending
											if v397 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1337), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v402 = m.ExcPending
												if v402 != 0 {
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
									v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v152 == v153 {
										v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v245 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
											v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v249 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											} else {
											}
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
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
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											}
										} else {
											v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v252 != 0 {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
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
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											}
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v158 = m.ExcPending
										if v158 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_16), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1350), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v170 = m.ExcPending
													if v170 != 0 {
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
						default:
							v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v245 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
								v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v249 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
								} else {
								}
								if l2 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v486 = m.ExcPending
									if v486 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v489 = m.ExcPending
										if v489 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
											mBase = m.M
											v493 = m.ExcPending
											if v493 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v498 = m.ExcPending
												if v498 != 0 {
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
									F_ReleaseCatCache(m, v14)
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return
									} else {
										m.G0 = v10 + int32(32)
										return
									}
								}
							} else {
								v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
								if v252 != 0 {
									F_ReleaseCatCache(m, v14)
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return
									} else {
										m.G0 = v10 + int32(32)
										return
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
									if l2 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v486 = m.ExcPending
										if v486 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v489 = m.ExcPending
											if v489 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
												mBase = m.M
												v493 = m.ExcPending
												if v493 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v498 = m.ExcPending
													if v498 != 0 {
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
										F_ReleaseCatCache(m, v14)
										mBase = m.M
										v257 = m.ExcPending
										if v257 != 0 {
											return
										} else {
											m.G0 = v10 + int32(32)
											return
										}
									}
								}
							}
						case 5:
							v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
							if v171 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v406 = m.ExcPending
								if v406 != 0 {
									return
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v409 = m.ExcPending
									if v409 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_assignProcTypes_17), int32(0))
										mBase = m.M
										v413 = m.ExcPending
										if v413 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1358), int32(_a_F_assignProcTypes_4))
											mBase = m.M
											v418 = m.ExcPending
											if v418 != 0 {
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
								v174 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
								if v174 != int32(2281) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v406 = m.ExcPending
									if v406 != 0 {
										return
									} else {
										F_errcode(m, int32(117833860))
										mBase = m.M
										v409 = m.ExcPending
										if v409 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_assignProcTypes_17), int32(0))
											mBase = m.M
											v413 = m.ExcPending
											if v413 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1358), int32(_a_F_assignProcTypes_4))
												mBase = m.M
												v418 = m.ExcPending
												if v418 != 0 {
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
									v177 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
									if v177 != int32(2278) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v422 = m.ExcPending
										if v422 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v425 = m.ExcPending
											if v425 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_18), int32(0))
												mBase = m.M
												v429 = m.ExcPending
												if v429 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1362), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v434 = m.ExcPending
													if v434 != 0 {
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
										v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v180 == v181 {
											v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v245 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
												v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v249 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												} else {
												}
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
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
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											} else {
												v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v252 != 0 {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
													if l2 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v486 = m.ExcPending
														if v486 != 0 {
															return
														} else {
															F_errcode(m, int32(117833860))
															mBase = m.M
															v489 = m.ExcPending
															if v489 != 0 {
																return
															} else {
																F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
																mBase = m.M
																v493 = m.ExcPending
																if v493 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
																	mBase = m.M
																	v498 = m.ExcPending
																	if v498 != 0 {
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
														F_ReleaseCatCache(m, v14)
														mBase = m.M
														v257 = m.ExcPending
														if v257 != 0 {
															return
														} else {
															m.G0 = v10 + int32(32)
															return
														}
													}
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v186 = m.ExcPending
											if v186 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v189 = m.ExcPending
												if v189 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_19), int32(0))
													mBase = m.M
													v193 = m.ExcPending
													if v193 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1375), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v198 = m.ExcPending
														if v198 != 0 {
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
					} else {
						v200 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return
						} else {
							v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+12)))
							if v202 != int32(1) {
								v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v245 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
									v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v249 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
									} else {
									}
									if l2 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v486 = m.ExcPending
										if v486 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v489 = m.ExcPending
											if v489 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
												mBase = m.M
												v493 = m.ExcPending
												if v493 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v498 = m.ExcPending
													if v498 != 0 {
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
										F_ReleaseCatCache(m, v14)
										mBase = m.M
										v257 = m.ExcPending
										if v257 != 0 {
											return
										} else {
											m.G0 = v10 + int32(32)
											return
										}
									}
								} else {
									v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v252 != 0 {
										F_ReleaseCatCache(m, v14)
										mBase = m.M
										v257 = m.ExcPending
										if v257 != 0 {
											return
										} else {
											m.G0 = v10 + int32(32)
											return
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
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
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								switch v205 - int32(1) {
								case 0:
									v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
									if v208 != int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v438 = m.ExcPending
										if v438 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v441 = m.ExcPending
											if v441 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_20), int32(0))
												mBase = m.M
												v445 = m.ExcPending
												if v445 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1385), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v450 = m.ExcPending
													if v450 != 0 {
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
										v211 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
										if v211 == int32(23) {
											v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v236 == int32(0) {
												v239 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v239
											} else {
											}
											v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v241 != 0 {
											} else {
												v242 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v242
											}
											v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v245 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
												v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v249 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												} else {
												}
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
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
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											} else {
												v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v252 != 0 {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
													if l2 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v486 = m.ExcPending
														if v486 != 0 {
															return
														} else {
															F_errcode(m, int32(117833860))
															mBase = m.M
															v489 = m.ExcPending
															if v489 != 0 {
																return
															} else {
																F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
																mBase = m.M
																v493 = m.ExcPending
																if v493 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
																	mBase = m.M
																	v498 = m.ExcPending
																	if v498 != 0 {
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
														F_ReleaseCatCache(m, v14)
														mBase = m.M
														v257 = m.ExcPending
														if v257 != 0 {
															return
														} else {
															m.G0 = v10 + int32(32)
															return
														}
													}
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v217 = m.ExcPending
											if v217 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v220 = m.ExcPending
												if v220 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_21), int32(0))
													mBase = m.M
													v224 = m.ExcPending
													if v224 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1389), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v229 = m.ExcPending
														if v229 != 0 {
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
								case 1:
									v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+104)))
									if v230 != int32(2) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v454 = m.ExcPending
										if v454 != 0 {
											return
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v457 = m.ExcPending
											if v457 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_assignProcTypes_22), int32(0))
												mBase = m.M
												v461 = m.ExcPending
												if v461 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1396), int32(_a_F_assignProcTypes_4))
													mBase = m.M
													v466 = m.ExcPending
													if v466 != 0 {
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
										v233 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
										if v233 != int32(20) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v470 = m.ExcPending
											if v470 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v473 = m.ExcPending
												if v473 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_23), int32(0))
													mBase = m.M
													v477 = m.ExcPending
													if v477 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1400), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v482 = m.ExcPending
														if v482 != 0 {
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
											v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v236 == int32(0) {
												v239 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v239
											} else {
											}
											v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
											if v241 != 0 {
											} else {
												v242 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v242
											}
											v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v245 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
												v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v249 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
												} else {
												}
												if l2 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v486 = m.ExcPending
													if v486 != 0 {
														return
													} else {
														F_errcode(m, int32(117833860))
														mBase = m.M
														v489 = m.ExcPending
														if v489 != 0 {
															return
														} else {
															F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
															mBase = m.M
															v493 = m.ExcPending
															if v493 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
																mBase = m.M
																v498 = m.ExcPending
																if v498 != 0 {
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
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												}
											} else {
												v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
												if v252 != 0 {
													F_ReleaseCatCache(m, v14)
													mBase = m.M
													v257 = m.ExcPending
													if v257 != 0 {
														return
													} else {
														m.G0 = v10 + int32(32)
														return
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
													if l2 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v486 = m.ExcPending
														if v486 != 0 {
															return
														} else {
															F_errcode(m, int32(117833860))
															mBase = m.M
															v489 = m.ExcPending
															if v489 != 0 {
																return
															} else {
																F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
																mBase = m.M
																v493 = m.ExcPending
																if v493 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
																	mBase = m.M
																	v498 = m.ExcPending
																	if v498 != 0 {
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
														F_ReleaseCatCache(m, v14)
														mBase = m.M
														v257 = m.ExcPending
														if v257 != 0 {
															return
														} else {
															m.G0 = v10 + int32(32)
															return
														}
													}
												}
											}
										}
									}
								default:
									v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v236 == int32(0) {
										v239 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v239
									} else {
									}
									v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									if v241 != 0 {
									} else {
										v242 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v242
									}
									v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v245 == int32(0) {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l2
										v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v249 == int32(0) {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
										} else {
										}
										if l2 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v486 = m.ExcPending
											if v486 != 0 {
												return
											} else {
												F_errcode(m, int32(117833860))
												mBase = m.M
												v489 = m.ExcPending
												if v489 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
													mBase = m.M
													v493 = m.ExcPending
													if v493 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
														mBase = m.M
														v498 = m.ExcPending
														if v498 != 0 {
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
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										}
									} else {
										v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
										if v252 != 0 {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v257 = m.ExcPending
											if v257 != 0 {
												return
											} else {
												m.G0 = v10 + int32(32)
												return
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l2
											if l2 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v486 = m.ExcPending
												if v486 != 0 {
													return
												} else {
													F_errcode(m, int32(117833860))
													mBase = m.M
													v489 = m.ExcPending
													if v489 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_assignProcTypes_5), int32(0))
														mBase = m.M
														v493 = m.ExcPending
														if v493 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1425), int32(_a_F_assignProcTypes_4))
															mBase = m.M
															v498 = m.ExcPending
															if v498 != 0 {
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
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v257 = m.ExcPending
												if v257 != 0 {
													return
												} else {
													m.G0 = v10 + int32(32)
													return
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
			v264 = m.ExcPending
			if v264 != 0 {
				return
			} else {
				v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v265
				F_errmsg_internal(m, int32(_a_F_assignProcTypes_24), v10)
				mBase = m.M
				v269 = m.ExcPending
				if v269 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1230), int32(_a_F_assignProcTypes_4))
					mBase = m.M
					v274 = m.ExcPending
					if v274 != 0 {
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
