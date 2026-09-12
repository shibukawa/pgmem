package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SnapBuildInitialSnapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
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
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L69
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L66
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L63
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L60
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L57
	}
L8:
	;
	if v31 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v31 = int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[653]))
	v21 = *(*int32)(unsafe.Add(mBase, _consts[654]))
	if v21 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v31 = base.B2i32(v19 != int32(0))
	goto L8
L13:
	;
	if v19 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v26 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v31 = int32(0)
	goto L8
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v34 != int32(2) {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L54
	}
L19:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v37 == int32(0) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+40))
	if v42 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v49 = F_MemoryContextAllocZero(m, v43, v44<<(uint(int32(2))%32)+int32(76))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = int32(5)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v57 = v49 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v55
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v64 = v60 << (uint(int32(2)) % 32)
	if v64 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	F_pg_qsort(m, v66, v60, int32(4), int32(185))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v65 = F__emscripten_memcpy_bulkmem(m, v57, v62, v64)
	mBase = m.M
	v66 = v65
	goto L26
L25:
	;
	v66 = v57
	goto L26
L26:
	;
	goto L23
L27:
	;
	v71 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+64)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v49)+44)) = v71
	v75 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+32)) = v75
	*(*int64)(unsafe.Add(mBase, uint32(v49)+20)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v49)+27)) = v75
	v82 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v86 = F_LWLockAcquire(m, v82+int32(512), int32(1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v89 = F_GetOldestSafeDecodingTransactionId(m, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v92+int32(512))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v97))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v89)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v109 != 0 {
		goto L4
	} else {
		goto L35
	}
L32:
	;
	v109 = base.B2i32(base.Ui32(v97) < base.Ui32(v89))
	goto L31
L33:
	;
	goto L34
L34:
	;
	v109 = base.B2i32(int32(0) < v89-v97)
	goto L31
L35:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[188]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+40)) = v112
	v115 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	goto L36
L36:
	;
	v119 = F_palloc(m, v116<<(uint(int32(2))%32))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v121
	v123 = int32(0)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	if v121-v124 < v123 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v130 = v123
	goto L41
L39:
	;
	v168 = v123
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = v119
	m.G0 = v8 + int32(16)
	return v49
L41:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v139 = F_bsearch(m, v8+int32(12), v135, v136, int32(4), int32(185))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v168 = v154
	goto L40
L43:
	;
	v155 = int32(3)
	v157 = v153 + int32(1)
	if base.Ui32(v157) <= base.Ui32(v155) {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	if v139 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v153 = v141
	v154 = v130
	goto L43
L46:
	;
	goto L47
L47:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	goto L48
L48:
	;
	if v144 <= v130 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v119+v130<<(uint(int32(2))%32)))) = v149
	v153 = v149
	v154 = v130 + int32(1)
	goto L43
L50:
	;
	v160 = v155
	goto L52
L51:
	;
	v160 = v157
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	if v160-v162 < int32(0) {
		v130 = v154
		goto L41
	} else {
		goto L53
	}
L53:
	;
	goto L42
L54:
	;
	F_errmsg_internal(m, int32(78561), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(522689), int32(454), int32(93228))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errmsg_internal(m, int32(368362), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(522689), int32(458), int32(93228))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errmsg_internal(m, int32(381269), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(522689), int32(461), int32(93228))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errmsg_internal(m, int32(455471), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(522689), int32(465), int32(93228))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v89
	F_errmsg_internal(m, int32(50557), v8)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(522689), int32(484), int32(93228))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(417820), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(522689), int32(514), int32(93228))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SnapBuildSetTwoPhaseAt(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = l1
	return
}
