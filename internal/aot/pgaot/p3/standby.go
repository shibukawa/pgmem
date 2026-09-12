package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StandbyReleaseLockTree(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if int32(0) < l1 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	v15 = int32(0)
	v17 = F_hash_search(m, v12, v8+int32(8), v15, v15)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_StandbyReleaseAllLocks(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L10
	}
L5:
	;
	return
L6:
	;
	if v17 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_StandbyReleaseXidEntryLocks(m, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	v27 = F_hash_search(m, v24, v17, int32(2), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L1
L10:
	;
	goto L1
L11:
	;
	v35 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	m.G0 = v8 + int32(16)
	return
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2+v35<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v43
	if v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L13
L16:
	;
	v67 = v35 + int32(1)
	if v67 != l1 {
		v35 = v67
		goto L14
	} else {
		goto L25
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	v49 = int32(0)
	v51 = F_hash_search(m, v46, v8+int32(12), v49, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_StandbyReleaseAllLocks(m)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L24
	}
L20:
	;
	if v51 == int32(0) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	F_StandbyReleaseXidEntryLocks(m, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	v61 = F_hash_search(m, v58, v51, int32(2), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	goto L16
L24:
	;
	goto L16
L25:
	;
	goto L15
}
func F_StandbyReleaseXidEntryLocks(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = v10
	goto L4
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	m.G0 = v8 + int32(48)
	return
L4:
	;
	v18 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v20
	F_errmsg_internal(m, int32(44844), v8+int32(16))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v37
	v45 = F_LockRelease(m, v8+int32(32), int32(8), int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L14
	}
L11:
	;
	F_errfinish(m, int32(469455), int32(1046), int32(144665))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v68 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	v71 = F_hash_search(m, v68, v14, int32(2), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L20
	}
L14:
	;
	if v45 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v49 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	if v49 == int32(0) {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v54
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v53
	F_errmsg_internal(m, int32(43553), v8)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(469455), int32(1053), int32(144665))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L13
L20:
	;
	if v66 != 0 {
		v14 = v66
		goto L4
	} else {
		goto L21
	}
L21:
	;
	goto L5
}
func F_StandbySlotsHaveCaughtup(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int64
	_ = v237
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	if v18 == v3 {
		v331 = v16
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return v331
L2:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v23 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v33 != 0 {
		v331 = v16
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+316))
	v31 = base.B2i32(v29 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v31)
	v33 = v31
	goto L6
L5:
	;
	v33 = int32(0)
	goto L6
L6:
	;
	goto L3
L7:
	;
	v35 = *(*int64)(unsafe.Add(mBase, _consts[704]))
	if base.B2i32(v35 != int64(0))&base.B2i32(base.Ui64(l0) <= base.Ui64(v35)) != 0 {
		v331 = v16
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v45 = F_LWLockAcquire(m, v41+int32(4736), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v51 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v317+int32(4736))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L9
	} else {
		goto L86
	}
L12:
	;
	v312 = v3
	v313 = int64(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v63 = v50 + int32(4)
	v66 = v3
	v67 = int64(0)
	goto L15
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	if v70 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v312 = v299
	v313 = v237
	goto L11
L17:
	;
	if base.Ui64(v67-int64(1)) < base.Ui64(v178) {
		goto L65
	} else {
		goto L66
	}
L18:
	;
	F_errcode(m, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L60
	}
L19:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v89)+88))
	if v146 != 0 {
		goto L41
	} else {
		goto L42
	}
L20:
	;
	v137 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L9
	} else {
		goto L39
	}
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	v78 = int32(0)
	goto L22
L22:
	;
	v89 = v75 + v78*int32(288)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v90 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if v89 != 0 {
		goto L19
	} else {
		goto L38
	}
L24:
	;
	goto L23
L25:
	;
	v94 = v89 + int32(24)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v98 == int32(0) {
		v117 = v97
		v118 = v98
		goto L29
	} else {
		goto L30
	}
L26:
	;
	goto L27
L27:
	;
	v123 = v78 + int32(1)
	if v123 != v70 {
		v78 = v123
		goto L22
	} else {
		goto L37
	}
L28:
	;
	if v118-v117 == int32(0) {
		goto L24
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	if v97 != v98 {
		v117 = v97
		v118 = v98
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v102 = v63
	v103 = v94
	goto L32
L32:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v107 == int32(0) {
		v117 = v106
		v118 = v107
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v117 = v106
	v118 = v107
	goto L29
L34:
	;
	v110 = int32(1)
	if v106 == v107 {
		v102 = v102 + v110
		v103 = v103 + v110
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	goto L27
L37:
	;
	goto L20
L38:
	;
	goto L20
L39:
	;
	if v137 == int32(0) {
		v312 = v66
		v313 = v67
		goto L11
	} else {
		goto L40
	}
L40:
	;
	v201 = int32(2958)
	v202 = int32(624035)
	v205 = int32(67591)
	v206 = int32(623449)
	v210 = int32(50856066)
	goto L18
L41:
	;
	v148 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L9
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(1)
	if v157 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v148 == int32(0) {
		v312 = v66
		v313 = v67
		goto L11
	} else {
		goto L45
	}
L45:
	;
	v201 = int32(2972)
	v202 = int32(623817)
	v205 = int32(659044)
	v206 = int32(623377)
	v210 = int32(50856066)
	goto L18
L46:
	;
	F_s_lock(m, v89, int32(469981), int32(2976), int32(220579))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v89)+112))
	if v167 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L48
L50:
	;
	v201 = v197
	v202 = v193
	v205 = v194
	v206 = v195
	v210 = int32(325)
	goto L18
L51:
	;
	v169 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L9
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v89)+104))
	if base.B2i32(v178 != int64(0))&base.B2i32(base.Ui64(l0) <= base.Ui64(v178)) != 0 {
		goto L17
	} else {
		goto L56
	}
L54:
	;
	if v169 == int32(0) {
		v312 = v66
		v313 = v67
		goto L11
	} else {
		goto L55
	}
L55:
	;
	v193 = int32(623965)
	v194 = int32(427635)
	v195 = int32(623449)
	v197 = int32(2992)
	goto L50
L56:
	;
	if v177 != 0 {
		v312 = v66
		v313 = v67
		goto L11
	} else {
		goto L57
	}
L57:
	;
	v184 = F_errstart(m, l1, int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	if v184 == int32(0) {
		v312 = v66
		v313 = v67
		goto L11
	} else {
		goto L59
	}
L59:
	;
	v193 = int32(623879)
	v194 = int32(412487)
	v195 = int32(623449)
	v197 = int32(3007)
	goto L50
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = int32(110399)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v63
	F_errmsg(m, v205, v14+int32(32))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v63
	F_errdetail(m, v206, v14+int32(16))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(110399)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v63
	F_errhint(m, v202, v14)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(469981), v201, int32(220579))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v312 = v66
	v313 = v67
	goto L11
L65:
	;
	v237 = v67
	goto L67
L66:
	;
	v237 = v178
	goto L67
L67:
	;
	if v63&int32(3) == int32(0) {
		v261 = v63
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v296 = int32(1)
	v299 = v66 + v296
	v301 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	if v299 < v302 {
		v63 = v294 + v63 + v296
		v66 = v299
		v67 = v237
		goto L15
	} else {
		goto L85
	}
L69:
	;
	v294 = v286 - v63
	goto L68
L70:
	;
	v265 = v261
	goto L79
L71:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v245 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v294 = int32(0)
	goto L68
L73:
	;
	goto L74
L74:
	;
	v250 = v63
	goto L75
L75:
	;
	v254 = v250 + int32(1)
	if v254&int32(3) == int32(0) {
		v261 = v254
		goto L70
	} else {
		goto L77
	}
L76:
	;
	v286 = v254
	goto L69
L77:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if v259 != 0 {
		v250 = v254
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v265)))
	v274 = int32(-2139062144)
	if (int32(16843008)-v271|v271)&v274 == v274 {
		v265 = v265 + int32(4)
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v280 = v265
	goto L82
L81:
	;
	goto L80
L82:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	if v284 != 0 {
		v280 = v280 + int32(1)
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v286 = v280
	goto L69
L84:
	;
	goto L83
L85:
	;
	goto L16
L86:
	;
	v323 = *(*int32)(unsafe.Add(mBase, _consts[703]))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	if v312 != v324 {
		v331 = int32(0)
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int64)(unsafe.Add(mBase, _consts[704])) = v313
	v331 = int32(1)
	goto L1
}
