package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SnapBuildGetTwoPhaseAt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	return v2
}
func F_SnapBuildInitialSnapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
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
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
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
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
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
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[0]))
	if v17 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
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
		goto L65
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L62
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L59
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L56
	}
L8:
	;
	if v33 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v33 = int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[1]))
	v21 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[2]))
	if base.B2i32(v20 == v21)|base.B2i32(v24 == v21) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v33 = base.B2i32(v24 != int32(0))
	goto L8
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v33 = int32(0)
	goto L8
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v36 != int32(2) {
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L53
	}
L18:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
	if v39 == int32(0) {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[3]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
	if v44 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v51 = F_MemoryContextAllocZero(m, v45, v46<<(uint(int32(2))%32)+int32(76))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(5)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = v51 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = v57
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v62
	v65 = v62 << (uint(int32(2)) % 32)
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	base.MemoryCopy(m, v59, v66, v65)
	goto L24
L23:
	;
	goto L24
L24:
	;
	F_pg_qsort(m, v59, v62, int32(4), int32(187))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v72 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v51)+64)) = v72
	*(*int64)(unsafe.Add(mBase, uint32(v51)+44)) = v72
	v76 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+32)) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v51)+20)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v51)+27)) = v76
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[4]))
	v87 = F_LWLockAcquire(m, v83+int32(512), int32(1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v90 = F_GetOldestSafeDecodingTransactionId(m, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[4]))
	F_LWLockRelease(m, v93+int32(512))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v98 = int32(3)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if base.B2i32(base.Ui32(v90) < base.Ui32(v98))|base.B2i32(base.Ui32(v100) < base.Ui32(v98)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v111)+52)) = v100
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[5]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	goto L35
L30:
	;
	if v90-v100 <= int32(0) {
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(v100) < base.Ui32(v90) {
		goto L3
	} else {
		goto L34
	}
L33:
	;
	goto L3
L34:
	;
	goto L29
L35:
	;
	v117 = F_palloc_mul(m, int32(4), v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v119
	v121 = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v119-v122 < v121 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v128 = v121
	goto L40
L38:
	;
	v167 = v121
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v117
	m.G0 = v9 + int32(16)
	return v51
L40:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	v138 = F_bsearch(m, v9+int32(12), v134, v135, int32(4), int32(187))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	v167 = v153
	goto L39
L42:
	;
	v154 = int32(3)
	v156 = v152 + int32(1)
	if base.Ui32(v156) <= base.Ui32(v154) {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	if v138 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v152 = v140
	v153 = v128
	goto L42
L45:
	;
	goto L46
L46:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[5]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	goto L47
L47:
	;
	if v143 <= v128 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v117+v128<<(uint(int32(2))%32)))) = v148
	v152 = v148
	v153 = v128 + int32(1)
	goto L42
L49:
	;
	v159 = v154
	goto L51
L50:
	;
	v159 = v156
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v159-v161 < int32(0) {
		v128 = v153
		goto L40
	} else {
		goto L52
	}
L52:
	;
	goto L41
L53:
	;
	F_errmsg_internal(m, int32(_a_F_SnapBuildInitialSnapshot_0), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_SnapBuildInitialSnapshot_1), int32(458), int32(_a_F_SnapBuildInitialSnapshot_2))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	F_errmsg_internal(m, int32(_a_F_SnapBuildInitialSnapshot_3), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_SnapBuildInitialSnapshot_1), int32(462), int32(_a_F_SnapBuildInitialSnapshot_2))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errmsg_internal(m, int32(_a_F_SnapBuildInitialSnapshot_4), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_SnapBuildInitialSnapshot_1), int32(465), int32(_a_F_SnapBuildInitialSnapshot_2))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errmsg_internal(m, int32(_a_F_SnapBuildInitialSnapshot_5), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_SnapBuildInitialSnapshot_1), int32(469), int32(_a_F_SnapBuildInitialSnapshot_2))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_SnapBuildInitialSnapshot_6), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_SnapBuildInitialSnapshot_1), int32(517), int32(_a_F_SnapBuildInitialSnapshot_2))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
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
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v90
	F_errmsg_internal(m, int32(_a_F_SnapBuildInitialSnapshot_7), v9)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_SnapBuildInitialSnapshot_1), int32(488), int32(_a_F_SnapBuildInitialSnapshot_2))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
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
