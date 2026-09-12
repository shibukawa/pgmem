package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WaitForLockers(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	v13 = F_list_make1_impl(m, int32(1), v6+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		F_WaitForLockersMultiple(m, v13, l1, int32(1))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_list_free(m, v13)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
func F_WaitForOlderSnapshots(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int64
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v275 int64
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v349 int64
	_ = v349
	var v350 int64
	_ = v350
	var v354 int32
	_ = v354
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v22 = F_GetCurrentVirtualXIDs(m, l0, v18+int32(12))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+12)))
	v28 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v28 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if int32(0) < v59 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	goto L5
L7:
	;
	goto L6
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v32 != int32(1) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v35 = int32(4543684)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v38 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v37 + v38
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v41 + v38
	*(*int64)(unsafe.Add(mBase, uint32(v28+int32(24))+232)) = v25
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v49 + v38
	v55 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v55 - v38
	goto L7
L10:
	;
	v62 = l0
	v63 = l1
	v64 = v18
	v68 = v22
	v76 = int64(0)
	goto L13
L11:
	;
	v354 = v18
	goto L12
L12:
	;
	m.G0 = v354 + int32(16)
	return
L13:
	;
	v77 = base.I32_wrap_i64(v76)
	v80 = v68 + v77<<(uint(int32(3))%32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v81 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v354 = v64
	goto L12
L15:
	;
	v349 = v76 + int64(1)
	v350 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+12)))
	if v349 < v350 {
		v76 = v349
		goto L13
	} else {
		goto L62
	}
L16:
	;
	if v76 != int64(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v63 == int32(0) {
		goto L15
	} else {
		goto L57
	}
L18:
	;
	v88 = F_GetCurrentVirtualXIDs(m, v62, v64+int32(8))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v63 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L21:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	if v76 < base.I64_extend_i32_s(v90) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v97 = v77
	v101 = v90
	goto L25
L23:
	;
	goto L24
L24:
	;
	F_pfree(m, v88)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L40
	}
L25:
	;
	v110 = v68 + v97<<(uint(int32(3))%32)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v111 == int32(0) {
		v170 = v101
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	v178 = v97 + int32(1)
	if v178 < v170 {
		v97 = v178
		v101 = v170
		goto L25
	} else {
		goto L39
	}
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if int32(0) < v114 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v122 = int32(0)
	goto L32
L30:
	;
	goto L31
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v110))) = int64(4294967295)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v170 = v161
	goto L27
L32:
	;
	v136 = v88 + v122<<(uint(int32(3))%32)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	if v137 == v117 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v111 == v139 {
		v170 = v101
		goto L27
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v142 = v122 + int32(1)
	if v142 != v114 {
		v122 = v142
		goto L32
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	goto L33
L39:
	;
	goto L26
L40:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v197 == int32(0) {
		goto L17
	} else {
		goto L41
	}
L41:
	;
	goto L20
L42:
	;
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v80)))
	*(*int64)(unsafe.Add(mBase, uint32(v64))) = v275
	v278 = F_VirtualXactLock(m, v64, int32(1))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L56
	}
L43:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v218 = int32(0)
	if v217 < v218 {
		v236 = v218
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v236 == int32(0) {
		goto L42
	} else {
		goto L51
	}
L45:
	;
	goto L44
L46:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
	if base.Ui32(v225) <= base.Ui32(v217) {
		v236 = int32(0)
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v230 = v227 + v217*int32(640)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v230)+44))
	if v232 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v233 = v230
	goto L50
L49:
	;
	v233 = int32(0)
	goto L50
L50:
	;
	v236 = v233
	goto L45
L51:
	;
	v240 = int64(*(*int32)(unsafe.Add(mBase, uint32(v236)+44)))
	v243 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v243 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L42
L53:
	;
	goto L52
L54:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v247 != int32(1) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v250 = int32(4543684)
	v252 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v253 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v252 + v253
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v256 + v253
	*(*int64)(unsafe.Add(mBase, uint32(v243+int32(40))+232)) = v240
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v264 + v253
	v270 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v270 - v253
	goto L53
L56:
	;
	goto L17
L57:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	if v302 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L15
L59:
	;
	goto L58
L60:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
	if v306 != int32(1) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v309 = int32(4543684)
	v311 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v312 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v311 + v312
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	*(*int32)(unsafe.Add(mBase, uint32(v302))) = v315 + v312
	*(*int64)(unsafe.Add(mBase, uint32(v302+int32(32))+232)) = v76 + int64(1)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	*(*int32)(unsafe.Add(mBase, uint32(v302))) = v323 + v312
	v329 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v329 - v312
	goto L59
L62:
	;
	goto L14
}
func F_WaitForWalSummarization(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v38 int64
	_ = v38
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v100 int64
	_ = v100
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v129 int64
	_ = v129
	var v137 int64
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v167 int64
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v185 int32
	_ = v185
	var v188 int64
	_ = v188
	var v196 int64
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int64
	_ = v229
	var v230 int64
	_ = v230
	var v236 int64
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v19 = base.I32_wrap_i64(int64(base.Ui64(l0) >> (uint(int64(32)) % 64)))
	v20 = base.I32_wrap_i64(l0)
	v24 = m.G0
	v25 = int32(16)
	v26 = v24 - v25
	m.G0 = v26
	F___gettimeofday(m, v26)
	mBase = m.M
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
	v30 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+8)))
	m.G0 = v26 + v25
	v38 = v30 + v29*int64(1000000) - int64(946684800000000)
	goto L1
L1:
	;
	v44 = int32(0)
	v47 = v38
	v48 = int64(0)
	goto L4
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L9
	} else {
		goto L50
	}
L3:
	;
	m.G0 = v15 - int32(-64)
	return
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v52 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L9
	} else {
		goto L49
	}
L6:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _consts[533])))
	if v56 != int32(1) {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v64 = F_LWLockAcquire(m, v60+int32(6272), int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[634]))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v67)+24))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	v71 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v71+int32(6272))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	if base.Ui64(v69) < base.Ui64(l0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v80 = m.G0
	v81 = int32(16)
	v82 = v80 - v81
	m.G0 = v82
	F___gettimeofday(m, v82)
	mBase = m.M
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
	v86 = int64(*(*int32)(unsafe.Add(mBase, uint32(v82)+8)))
	m.G0 = v82 + v81
	v94 = v86 + v85*int64(1000000) - int64(946684800000000)
	goto L18
L15:
	;
	goto L16
L16:
	;
	goto L5
L17:
	;
	if v94 <= v181 {
		v199 = int32(0)
		goto L44
	} else {
		goto L45
	}
L18:
	;
	if v94 <= v47 {
		v111 = int32(0)
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v111 < int32(10000) {
		v180 = v44
		v181 = v47
		v182 = v48
		goto L17
	} else {
		goto L24
	}
L20:
	;
	goto L19
L21:
	;
	v97 = int32(2147483647)
	v100 = v94 - v47
	if base.B2i32(int64(0) < v47)^base.B2i32(v100 < v94) != 0 {
		v111 = v97
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if int64(2147483646000) < v100 {
		v111 = v97
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v108 = base.I64_div_s(v100+int64(999), int64(1000))
	v111 = base.I32_wrap_i64(v108)
	goto L20
L24:
	;
	v117 = base.B2i32(base.Ui64(v48) < base.Ui64(v68))
	if base.Ui64(v48) < base.Ui64(v68) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v118 = int32(0)
	goto L27
L26:
	;
	v118 = v44 + int32(1)
	goto L27
L27:
	;
	if int32(6) <= v118 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	if base.Ui64(v48) < base.Ui64(v68) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v121 = v68
	goto L31
L30:
	;
	v121 = v48
	goto L31
L31:
	;
	v123 = v47 + int64(10000000)
	if v94 <= v38 {
		v140 = int32(0)
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v143 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L9
	} else {
		goto L37
	}
L33:
	;
	goto L32
L34:
	;
	v126 = int32(2147483647)
	v129 = v94 - v38
	if base.B2i32(int64(0) < v38)^base.B2i32(v129 < v94) != 0 {
		v140 = v126
		goto L33
	} else {
		goto L35
	}
L35:
	;
	if int64(2147483646000) < v129 {
		v140 = v126
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v137 = base.I64_div_s(v129+int64(999), int64(1000))
	v140 = base.I32_wrap_i64(v137)
	goto L33
L37:
	;
	if v143 == int32(0) {
		v180 = v118
		v181 = v123
		v182 = v121
		goto L17
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v151 = base.I32_div_s(v140, int32(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v19
	F_errmsg_plural(m, int32(442024), int32(181580), v151, v13+int32(-16))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+44)) = uint32(v68)
	v162 = int64(32)
	v163 = int64(base.Ui64(v68) >> (uint(v162) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+40)) = uint32(v163)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+36)) = uint32(v69)
	v167 = int64(base.Ui64(v69) >> (uint(v162) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+32)) = uint32(v167)
	F_errdetail(m, int32(594407), v13+int32(-32))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(515130), int32(765), int32(269237))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	v180 = v118
	v181 = v123
	v182 = v121
	goto L17
L43:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _consts[634]))
	v207 = F_ConditionVariableTimedSleep(m, v201+int32(32), int32(10000)-v199, int32(134217783))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L9
	} else {
		goto L48
	}
L44:
	;
	goto L43
L45:
	;
	v185 = int32(2147483647)
	v188 = v94 - v181
	if base.B2i32(int64(0) < v181)^base.B2i32(v188 < v94) != 0 {
		v199 = v185
		goto L44
	} else {
		goto L46
	}
L46:
	;
	if int64(2147483646000) < v188 {
		v199 = v185
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v196 = base.I64_div_s(v188+int64(999), int64(1000))
	v199 = base.I32_wrap_i64(v196)
	goto L44
L48:
	;
	v44 = v180
	v47 = v181
	v48 = v182
	goto L4
L49:
	;
	goto L3
L50:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(343401), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+20)) = uint32(v68)
	v229 = int64(32)
	v230 = int64(base.Ui64(v68) >> (uint(v229) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+16)) = uint32(v230)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+12)) = uint32(v69)
	v236 = int64(base.Ui64(v69) >> (uint(v229) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)) = uint32(v236)
	F_errdetail(m, int32(594317), v15)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(515130), int32(747), int32(269237))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_for_freed_segments_locked(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1468))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	if v8 != v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v21 = l0 + int32(8) + v14*int32(20)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v22 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+652)) = v8
	goto L3
L6:
	;
	v36 = v14 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v36) <= base.Ui32(v37) {
		v14 = v36
		goto L4
	} else {
		goto L11
	}
L7:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)))
	if v25 != int32(1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	F_dsm_detach(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = int64(0)
	goto L6
L11:
	;
	goto L5
}
