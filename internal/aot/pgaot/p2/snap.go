package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

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
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
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
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L67
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L64
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L61
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L58
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L55
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
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L52
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+40))
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
	F_pg_qsort(m, v59, v62, int32(4), int32(185))
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
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[4]))
	v88 = F_LWLockAcquire(m, v84+int32(512), int32(1))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v91 = F_GetOldestSafeDecodingTransactionId(m, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[4]))
	F_LWLockRelease(m, v94+int32(512))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v99))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v91)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v111 != 0 {
		goto L4
	} else {
		goto L33
	}
L30:
	;
	v111 = base.B2i32(base.Ui32(v99) < base.Ui32(v91))
	goto L29
L31:
	;
	goto L32
L32:
	;
	v111 = base.B2i32(int32(0) < v91-v99)
	goto L29
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[3]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+40)) = v114
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[5]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	goto L34
L34:
	;
	v121 = F_palloc(m, v118<<(uint(int32(2))%32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v123-v125 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v132 = v76
	goto L39
L37:
	;
	v171 = v76
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+16)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+12)) = v121
	m.G0 = v9 + int32(16)
	return v51
L39:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	v141 = F_bsearch(m, v9+int32(12), v137, v138, int32(4), int32(185))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	v171 = v156
	goto L38
L41:
	;
	v157 = int32(3)
	v159 = v155 + int32(1)
	if base.Ui32(v159) <= base.Ui32(v157) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	if v141 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v155 = v143
	v156 = v132
	goto L41
L44:
	;
	goto L45
L45:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildInitialSnapshot[5]))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	goto L46
L46:
	;
	if v146 <= v132 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v121+v132<<(uint(int32(2))%32)))) = v151
	v155 = v151
	v156 = v132 + int32(1)
	goto L41
L48:
	;
	v162 = v157
	goto L50
L49:
	;
	v162 = v159
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v162-v164 < int32(0) {
		v132 = v156
		goto L39
	} else {
		goto L51
	}
L51:
	;
	goto L40
L52:
	;
	F_errmsg_internal(m, int32(_a_F_SnapBuildInitialSnapshot_0), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_SnapBuildInitialSnapshot_1), int32(454), int32(_a_F_SnapBuildInitialSnapshot_2))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errmsg_internal(m, int32(_a_F_SnapBuildInitialSnapshot_3), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_SnapBuildInitialSnapshot_1), int32(458), int32(_a_F_SnapBuildInitialSnapshot_2))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errmsg_internal(m, int32(_a_F_SnapBuildInitialSnapshot_4), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_SnapBuildInitialSnapshot_1), int32(461), int32(_a_F_SnapBuildInitialSnapshot_2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_SnapBuildInitialSnapshot_5), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_SnapBuildInitialSnapshot_1), int32(465), int32(_a_F_SnapBuildInitialSnapshot_2))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v91
	F_errmsg_internal(m, int32(_a_F_SnapBuildInitialSnapshot_6), v9)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_SnapBuildInitialSnapshot_1), int32(484), int32(_a_F_SnapBuildInitialSnapshot_2))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(_a_F_SnapBuildInitialSnapshot_7), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_SnapBuildInitialSnapshot_1), int32(514), int32(_a_F_SnapBuildInitialSnapshot_2))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
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
