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
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int64
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int64
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
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
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
	F_LWLockRelease(m, v310+int32(512))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L6
	} else {
		goto L68
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
	v274 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v274
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+73)) = uint8(v274)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)))
	if v282&int32(14) == v274 {
		goto L1
	} else {
		goto L66
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18+v19<<(uint(int32(2))%32)))) = v13
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v13
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+73)) = uint8(v13)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)))
	if v33&int32(14) != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+544)) = l1
	v88 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+536)) = uint8(v88)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+540)) = v90
	v94 = base.I32_div_s(l0-v86, int32(640))
	v96 = base.AtomicRmwCmpxchg32(m, v85, int32(52), v90, v94)
	if v96 != v90 {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	goto L2
L12:
	;
	v37 = v33 & int32(241)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)) = uint8(v37)
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v41+v42))) = uint8(v37)
	goto L14
L13:
	;
	goto L14
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+276)))
	if v46 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v69 = int32(_a_F_ProcArrayEndTransaction_0)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[2]))
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)+48))
	v72 = base.I32_wrap_i64(v71)
	v73 = F_TransactionIdPrecedes(m, v72, l1)
	mBase = m.M
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[2]))
	if v73 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+277)))
	if v49 != int32(1) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v53 = v19 << (uint(int32(1)) % 32)
	v54 = int32(_a_F_ProcArrayEndTransaction_1)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53+v56))) = uint8(v58)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v53)+1)) = uint8(v58)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+276)) = uint16(v58)
	goto L15
L19:
	;
	goto L18
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v75)+48)) = v71 + base.I64_extend_i32_s(l1-v72)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v75)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v75)+56)) = v80 + int64(1)
	goto L11
L23:
	;
	v100 = v96
	goto L26
L24:
	;
	v108 = v90
	goto L25
L25:
	;
	if v108 != int32(-1) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+540)) = v100
	v105 = base.AtomicRmwCmpxchg32(m, v85, int32(52), v100, v94)
	if v100 != v105 {
		v100 = v105
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v108 = v100
	goto L25
L28:
	;
	goto L27
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = int32(134217769)
	v121 = int32(0)
	goto L32
L30:
	;
	goto L31
L31:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
	v153 = F_LWLockAcquire(m, v149+int32(512), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L6
	} else {
		goto L41
	}
L32:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_PGSemaphoreLock(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L34
	}
L33:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[3]))
	v132 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v132
	if v121 <= v132 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+536)))
	if v129 != 0 {
		v121 = v121 + int32(1)
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v137 = v121
	goto L37
L37:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_PGSemaphoreUnlock(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L39
	}
L38:
	;
	goto L1
L39:
	;
	v144 = int32(1)
	if base.Ui32(v144) < base.Ui32(v137) {
		v137 = v137 - v144
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v155 = int32(-1)
	v157 = base.AtomicRmwXchg32(m, v85, int32(52), v155)
	if v157 == v155 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v160 = v157
	goto L43
L43:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[4]))
	v169 = v166 + v160*int32(640)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+544))
	v171 = int32(0)
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v169)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v176+v177<<(uint(int32(2))%32)))) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v169)+56)) = v171
	*(*int64)(unsafe.Add(mBase, uint32(v169)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v169)+120)) = v171
	*(*uint8)(unsafe.Add(mBase, uint32(v169)+73)) = uint8(v171)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+124)))
	if v191&int32(14) != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
	F_LWLockRelease(m, v246+int32(512))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L6
	} else {
		goto L58
	}
L45:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v169)+540))
	if v242 != int32(-1) {
		v160 = v242
		goto L43
	} else {
		goto L57
	}
L46:
	;
	v195 = v191 & int32(241)
	*(*uint8)(unsafe.Add(mBase, uint32(v169)+124)) = uint8(v195)
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v169)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v199+v200))) = uint8(v195)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+276)))
	if v204 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v227 = int32(_a_F_ProcArrayEndTransaction_0)
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[2]))
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v228)+48))
	v230 = base.I32_wrap_i64(v229)
	v231 = F_TransactionIdPrecedes(m, v230, v170)
	mBase = m.M
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[2]))
	if v231 != 0 {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+277)))
	if v207 != int32(1) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v211 = v177 << (uint(int32(1)) % 32)
	v212 = int32(_a_F_ProcArrayEndTransaction_1)
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	v216 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v211+v214))) = uint8(v216)
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v220+v211)+1)) = uint8(v216)
	*(*uint16)(unsafe.Add(mBase, uint32(v169)+276)) = uint16(v216)
	goto L49
L53:
	;
	goto L52
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v233)+48)) = v229 + base.I64_extend_i32_s(v170-v230)
	goto L56
L55:
	;
	goto L56
L56:
	;
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v233)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v233)+56)) = v238 + int64(1)
	goto L45
L57:
	;
	goto L44
L58:
	;
	v252 = v157
	goto L59
L59:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[4]))
	v260 = v257 + v252*int32(640)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+540))
	*(*int32)(unsafe.Add(mBase, uint32(v260)+540)) = int32(-1)
	v264 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v260)+536)) = uint8(v264)
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[5]))
	if v267 != v260 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L1
L61:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	F_PGSemaphoreUnlock(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v261 != int32(-1) {
		v252 = v261
		goto L59
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	goto L60
L66:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[0]))
	v292 = F_LWLockAcquire(m, v288+int32(512), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)))
	v296 = v294 & int32(-15)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)) = uint8(v296)
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_ProcArrayEndTransaction[1]))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v300+v301))) = uint8(v296)
	goto L2
L68:
	;
	goto L1
}
func F_ProcLockWakeup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v9 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = l1 + int32(32)
	if v12 == v16 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = v12
	v23 = v3
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32))))
	if v32&v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L1
L7:
	;
	if v26 != v16 {
		v22 = v26
		v23 = v141
		goto L5
	} else {
		goto L32
	}
L8:
	;
	v141 = int32(1)<<(uint(v28)%32) | v23
	goto L7
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
	v35 = F_LockCheckConflicts(m, l0, v28, l1, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v35 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	v41 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+128)) = v40 + v41
	v46 = l1 + v28<<(uint(int32(2))%32)
	v48 = v46 + int32(88)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v49 + v41
	v54 = v41 << (uint(v28) % 32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v54 | v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	if v58 == v59 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v69 == int32(0) {
		v141 = v23
		goto L7
	} else {
		goto L17
	}
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v61 & (v54 ^ int32(-1))
	goto L16
L15:
	;
	goto L16
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v66 | v54
	goto L13
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v69
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v75
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+40)) = v79 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+92)) = v77
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_ProcLockWakeup[0]))
	v91 = base.AtomicRmwXchg64(m, v88, int32(112), v77)
	v93 = v22 + int32(20)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v141 = v23
	goto L7
L19:
	;
	goto L18
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v97 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	if v100 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_ProcLockWakeup[1]))
	if v104 == v100 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v106 = m.G0
	v108 = v106 - int32(16)
	m.G0 = v108
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_ProcLockWakeup[2]))
	if v111 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v134 = F_pgmem_kill(m, v100, int32(23))
	mBase = m.M
	goto L19
L26:
	;
	m.G0 = v108 + int32(16)
	goto L18
L27:
	;
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v108)+15)) = uint8(v114)
	goto L28
L28:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_ProcLockWakeup[3]))
	v122 = F_write(m, v118, v108+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v122 {
		goto L26
	} else {
		goto L30
	}
L29:
	;
	goto L26
L30:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_ProcLockWakeup[4]))
	if v126 == int32(27) {
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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	if int32(0) <= l0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v28 = v23 + l0*int32(640) + int32(20)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v29 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_ProcSendSignal_0), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_errfinish(m, int32(_a_F_ProcSendSignal_1), int32(1989), int32(_a_F_ProcSendSignal_2))
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
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v32 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v35 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSendSignal[1]))
	if v39 == v35 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v41 = m.G0
	v43 = v41 - int32(16)
	m.G0 = v43
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSendSignal[2]))
	if v46 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v69 = F_pgmem_kill(m, v35, int32(23))
	mBase = m.M
	goto L11
L18:
	;
	m.G0 = v43 + int32(16)
	goto L10
L19:
	;
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+15)) = uint8(v49)
	goto L20
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSendSignal[3]))
	v57 = F_write(m, v53, v43+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v57 {
		goto L18
	} else {
		goto L22
	}
L21:
	;
	goto L18
L22:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_ProcSendSignal[4]))
	if v61 == int32(27) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ProcWaitForSignal[0]))
	v6 = F_WaitLatch(m, v3, int32(33), int32(0), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_ProcWaitForSignal[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_ProcWaitForSignal[1]))
		if v13 != 0 {
			F_ProcessInterrupts(m)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
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
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
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
	goto L4
L2:
	;
	m.G0 = v9 + int32(192)
	return
L3:
	;
	goto L2
L4:
	;
	if v16 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v153 = int32(m.ExcTag)
	v154 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v153 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L7:
	;
	v134 = int32(_a_F_ProcessProcSignalBarrier_0)
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0]))
	v137 = base.AtomicRmwXchg64(m, v135, int32(104), v133)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v9)+184)) = v133
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0]))
	F_ConditionVariableBroadcast(m, v142+int32(116))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L32
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[1]))
	if v21 == int32(0) {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	v61 = v13
	v62 = v14
	v63 = v15
	v64 = v17
	goto L10
L10:
	;
	if v63 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v25 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[1])) = v25
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0]))
	v29 = int64(0)
	v32 = base.AtomicRmwCmpxchg64(m, v28, int32(104), v29, v29)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[2]))
	v38 = base.AtomicRmwCmpxchg64(m, v34, v25, v29, v29)
	if v32 == v38 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0]))
	v42 = int32(0)
	v44 = base.AtomicRmwXchg32(m, v41, int32(112), v42)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v44
	if v44 == v42 {
		v129 = v13
		v130 = v14
		v133 = v38
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[3]))
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[4]))
	goto L14
L14:
	;
	v55 = v9 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v9 + int32(12)
	goto L17
L15:
	;
	v61 = v51
	v62 = v53
	v63 = int32(0)
	v64 = v38
	goto L10
L17:
	;
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[3])) = v61
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[4])) = v62
	v129 = v61
	v130 = v62
	v133 = v64
	goto L7
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[4])) = v9 + int32(16)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	if v72 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[3])) = v61
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[4])) = v62
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[0]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	v112 = base.AtomicRmwOr32(m, v109, int32(112), v110)
	v114 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[1])) = v114
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[5])) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v9)+184)) = v64
	F_pg_re_throw(m)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L31
	}
L22:
	;
	goto L23
L23:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	v82 = base.I32_ctz(v81)
	if v82 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[3])) = v61
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessProcSignalBarrier[4])) = v62
	v129 = v61
	v130 = v62
	v133 = v64
	goto L7
L25:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	if v99 != 0 {
		goto L23
	} else {
		goto L30
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+180)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v9)+184)) = v64
	F_smgrreleaseall(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v94 & base.I32_rotl(int32(-2), v82)
	goto L25
L29:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v9)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v90 & int32(-2)
	goto L25
L30:
	;
	goto L24
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	goto L5
L33:
	;
	v158 = int32(v154)
	m.G0 = v9
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v9+int32(12) == v164 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	m.ExcPending = 1
	goto L42
L35:
	;
	if v168 != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v168 = v166
	goto L38
L37:
	;
	v168 = int32(0)
	goto L38
L38:
	;
	goto L35
L39:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v9)+184))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v9)+180))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v9)+176))
	v13 = v171
	v14 = v170
	v15 = v160
	v16 = v168
	v17 = v169
	goto L1
L40:
	;
	goto L41
L41:
	;
	F___wasm_longjmp(m, v161, v160)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	return
L43:
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
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
	v96 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v16))), uint32(v96))
	goto L1
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[1]))
	v12 = v9 + l2<<(uint(int32(7))%32)
	v14 = v12 + int32(8)
	v16 = v12 + int32(104)
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
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[2]))
	v43 = v41 + int32(37)
	if v43 < int32(0) {
		goto L1
	} else {
		goto L12
	}
L6:
	;
	F_s_lock(m, v16, int32(_a_F_SendProcSignal_0), int32(293), int32(_a_F_SendProcSignal_1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v27 != l0 {
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
	v34 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14)+96)), uint32(v34))
	v38 = F_pgmem_kill(m, l0, int32(10))
	mBase = m.M
	return v38
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[1]))
	v50 = v43
	v51 = v47
	goto L13
L13:
	;
	v55 = v51 + v50<<(uint(int32(7))%32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if l0 == v56 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59+l1<<(uint(int32(2))%32))+40)) = int32(1)
	v90 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v59)+96)), uint32(v90))
	v94 = F_pgmem_kill(m, l0, int32(10))
	mBase = m.M
	return v94
L15:
	;
	goto L14
L16:
	;
	v59 = v55 + int32(8)
	v61 = v55 + int32(104)
	v64 = base.AtomicRmwXchg32(m, v59, int32(96), int32(1))
	if v64 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v77 = v51
	goto L18
L18:
	;
	v79 = int32(0)
	if base.B2i32(v50 <= v79) == v79 {
		v50 = v50 - int32(1)
		v51 = v77
		goto L13
	} else {
		goto L24
	}
L19:
	;
	F_s_lock(m, v61, int32(_a_F_SendProcSignal_0), int32(321), int32(_a_F_SendProcSignal_1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v70 == l0 {
		goto L15
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v72 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v61))), uint32(v72))
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_SendProcSignal[1]))
	v77 = v76
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
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
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1241), int32(_a_F_assignProcTypes_4))
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
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1241), int32(_a_F_assignProcTypes_4))
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
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1241), int32(_a_F_assignProcTypes_4))
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
										F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1224), int32(_a_F_assignProcTypes_4))
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
									F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1224), int32(_a_F_assignProcTypes_4))
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
									F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1231), int32(_a_F_assignProcTypes_4))
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
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1241), int32(_a_F_assignProcTypes_4))
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
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1241), int32(_a_F_assignProcTypes_4))
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
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1241), int32(_a_F_assignProcTypes_4))
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
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1259), int32(_a_F_assignProcTypes_4))
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
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1263), int32(_a_F_assignProcTypes_4))
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
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1280), int32(_a_F_assignProcTypes_4))
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
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1280), int32(_a_F_assignProcTypes_4))
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
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1284), int32(_a_F_assignProcTypes_4))
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
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1295), int32(_a_F_assignProcTypes_4))
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
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1299), int32(_a_F_assignProcTypes_4))
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
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1315), int32(_a_F_assignProcTypes_4))
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
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1319), int32(_a_F_assignProcTypes_4))
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
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1332), int32(_a_F_assignProcTypes_4))
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
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
											F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1340), int32(_a_F_assignProcTypes_4))
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
												F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1340), int32(_a_F_assignProcTypes_4))
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
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1344), int32(_a_F_assignProcTypes_4))
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
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
																	F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1357), int32(_a_F_assignProcTypes_4))
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
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1367), int32(_a_F_assignProcTypes_4))
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
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
																	F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1371), int32(_a_F_assignProcTypes_4))
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
													F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1378), int32(_a_F_assignProcTypes_4))
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
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1382), int32(_a_F_assignProcTypes_4))
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
																F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
																	F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
														F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
															F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1407), int32(_a_F_assignProcTypes_4))
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
					F_errfinish(m, int32(_a_F_assignProcTypes_3), int32(1212), int32(_a_F_assignProcTypes_4))
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
