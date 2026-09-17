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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int64
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v279 int64
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v355 int64
	_ = v355
	var v356 int64
	_ = v356
	var v360 int32
	_ = v360
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
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[0]))
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
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if int32(0) < v61 {
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
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[1])))
	if v32&int32(1) == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v37 = int32(_a_F_WaitForOlderSnapshots_0)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[2]))
	v40 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[2])) = v39 + v40
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v43 + v40
	*(*int64)(unsafe.Add(mBase, uint32(v28+int32(24))+232)) = v25
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v51 + v40
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[2])) = v57 - v40
	goto L7
L10:
	;
	v64 = l0
	v65 = l1
	v66 = v18
	v70 = v22
	v78 = int64(0)
	goto L13
L11:
	;
	v360 = v18
	goto L12
L12:
	;
	m.G0 = v360 + int32(16)
	return
L13:
	;
	v79 = base.I32_wrap_i64(v78)
	v82 = v70 + v79<<(uint(int32(3))%32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v83 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v360 = v66
	goto L12
L15:
	;
	v355 = v78 + int64(1)
	v356 = int64(*(*int32)(unsafe.Add(mBase, uint32(v66)+12)))
	if v355 < v356 {
		v78 = v355
		goto L13
	} else {
		goto L62
	}
L16:
	;
	if v78 != int64(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v65 == int32(0) {
		goto L15
	} else {
		goto L57
	}
L18:
	;
	v90 = F_GetCurrentVirtualXIDs(m, v64, v66+int32(8))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v65 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L21:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	if v78 < base.I64_extend_i32_s(v92) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v98 = v79
	v103 = v92
	goto L25
L23:
	;
	goto L24
L24:
	;
	F_pfree(m, v90)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L40
	}
L25:
	;
	v112 = v70 + v98<<(uint(int32(3))%32)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v113 == int32(0) {
		v172 = v103
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	v180 = v98 + int32(1)
	if v180 < v172 {
		v98 = v180
		v103 = v172
		goto L25
	} else {
		goto L39
	}
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if int32(0) < v116 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v126 = int32(0)
	goto L32
L30:
	;
	goto L31
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v112))) = int64(4294967295)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v172 = v163
	goto L27
L32:
	;
	v138 = v90 + v126<<(uint(int32(3))%32)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if v139 == v119 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v113 == v141 {
		v172 = v103
		goto L27
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v144 = v126 + int32(1)
	if v144 != v116 {
		v126 = v144
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
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v199 == int32(0) {
		goto L17
	} else {
		goto L41
	}
L41:
	;
	goto L20
L42:
	;
	v279 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
	*(*int64)(unsafe.Add(mBase, uint32(v66))) = v279
	v282 = F_VirtualXactLock(m, v66, int32(1))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L56
	}
L43:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v220 = int32(0)
	if v219 < v220 {
		v238 = v220
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v238 == int32(0) {
		goto L42
	} else {
		goto L51
	}
L45:
	;
	goto L44
L46:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[3]))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
	if base.Ui32(v227) <= base.Ui32(v219) {
		v238 = int32(0)
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v232 = v229 + v219*int32(640)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v232)+44))
	if v234 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v235 = v232
	goto L50
L49:
	;
	v235 = int32(0)
	goto L50
L50:
	;
	v238 = v235
	goto L45
L51:
	;
	v242 = int64(*(*int32)(unsafe.Add(mBase, uint32(v238)+44)))
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[0]))
	if v245 == int32(0) {
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
	v249 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[1])))
	if v249&int32(1) == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v254 = int32(_a_F_WaitForOlderSnapshots_0)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[2]))
	v257 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[2])) = v256 + v257
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v260 + v257
	*(*int64)(unsafe.Add(mBase, uint32(v245+int32(40))+232)) = v242
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = v268 + v257
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[2])) = v274 - v257
	goto L53
L56:
	;
	goto L17
L57:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[0]))
	if v306 == int32(0) {
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
	v310 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[1])))
	if v310&int32(1) == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v315 = int32(_a_F_WaitForOlderSnapshots_0)
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[2]))
	v318 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[2])) = v317 + v318
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v321 + v318
	*(*int64)(unsafe.Add(mBase, uint32(v306+int32(32))+232)) = v78 + int64(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v329 + v318
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForOlderSnapshots[2])) = v335 - v318
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
	var v49 int64
	_ = v49
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
	var v100 int64
	_ = v100
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v130 int64
	_ = v130
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v169 int64
	_ = v169
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v190 int64
	_ = v190
	var v199 int64
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v239 int64
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v19 = base.I32_wrap_i64(int64(base.Ui64(l0) >> (uint(int64(32)) % 64)))
	v20 = base.I32_wrap_i64(l0)
	v24 = m.G0
	v25 = int32(16)
	v26 = v24 - v25
	m.G0 = v26
	F_gettimeofday(m, v26)
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
	v49 = int64(0)
	goto L4
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L9
	} else {
		goto L47
	}
L3:
	;
	m.G0 = v15 - int32(-64)
	return
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForWalSummarization[0]))
	if v52 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L9
	} else {
		goto L46
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
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForWalSummarization[1])))
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
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForWalSummarization[2]))
	v64 = F_LWLockAcquire(m, v60+int32(_a_F_WaitForWalSummarization_0), int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForWalSummarization[3]))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v67)+24))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForWalSummarization[2]))
	F_LWLockRelease(m, v71+int32(_a_F_WaitForWalSummarization_0))
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
	F_gettimeofday(m, v82)
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
	if v94 <= v183 {
		v202 = int32(0)
		goto L42
	} else {
		goto L43
	}
L18:
	;
	if v94 <= v47 {
		v112 = int32(0)
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v112 < int32(_a_F_WaitForWalSummarization_1) {
		v182 = v44
		v183 = v47
		v184 = v49
		goto L17
	} else {
		goto L23
	}
L20:
	;
	goto L19
L21:
	;
	v100 = v94 - v47
	if base.B2i32(int64(0) < v47)^base.B2i32(v100 < v94)|base.B2i32(int64(2147483646000) < v100) != 0 {
		v112 = int32(2147483647)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v109 = base.I64_div_s(v100+int64(999), int64(1000))
	v112 = base.I32_wrap_i64(v109)
	goto L20
L23:
	;
	v118 = base.B2i32(base.Ui64(v49) < base.Ui64(v68))
	if base.Ui64(v49) < base.Ui64(v68) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v119 = int32(0)
	goto L26
L25:
	;
	v119 = v44 + int32(1)
	goto L26
L26:
	;
	if int32(6) <= v119 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	if base.Ui64(v49) < base.Ui64(v68) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v122 = v68
	goto L30
L29:
	;
	v122 = v49
	goto L30
L30:
	;
	v124 = v47 + int64(10000000)
	if v94 <= v38 {
		v142 = int32(0)
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v145 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L9
	} else {
		goto L35
	}
L32:
	;
	goto L31
L33:
	;
	v130 = v94 - v38
	if base.B2i32(int64(0) < v38)^base.B2i32(v130 < v94)|base.B2i32(int64(2147483646000) < v130) != 0 {
		v142 = int32(2147483647)
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v139 = base.I64_div_s(v130+int64(999), int64(1000))
	v142 = base.I32_wrap_i64(v139)
	goto L32
L35:
	;
	if v145 == int32(0) {
		v182 = v119
		v183 = v124
		v184 = v122
		goto L17
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	v153 = base.I32_div_s(v142, int32(1000))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v19
	F_errmsg_plural(m, int32(_a_F_WaitForWalSummarization_2), int32(_a_F_WaitForWalSummarization_3), v153, v13+int32(-16))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+44)) = uint32(v68)
	v164 = int64(32)
	v165 = int64(base.Ui64(v68) >> (uint(v164) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+40)) = uint32(v165)
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+36)) = uint32(v69)
	v169 = int64(base.Ui64(v69) >> (uint(v164) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+32)) = uint32(v169)
	F_errdetail(m, int32(_a_F_WaitForWalSummarization_4), v13+int32(-32))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_WaitForWalSummarization_5), int32(765), int32(_a_F_WaitForWalSummarization_6))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	v182 = v119
	v183 = v124
	v184 = v122
	goto L17
L41:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForWalSummarization[3]))
	v210 = F_ConditionVariableTimedSleep(m, v204+int32(32), int32(_a_F_WaitForWalSummarization_1)-v202, int32(134217783))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L9
	} else {
		goto L45
	}
L42:
	;
	goto L41
L43:
	;
	v190 = v94 - v183
	if base.B2i32(int64(0) < v183)^base.B2i32(v190 < v94)|base.B2i32(int64(2147483646000) < v190) != 0 {
		v202 = int32(2147483647)
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v199 = base.I64_div_s(v190+int64(999), int64(1000))
	v202 = base.I32_wrap_i64(v199)
	goto L42
L45:
	;
	v44 = v182
	v47 = v183
	v49 = v184
	goto L4
L46:
	;
	goto L3
L47:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(_a_F_WaitForWalSummarization_7), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+20)) = uint32(v68)
	v232 = int64(32)
	v233 = int64(base.Ui64(v68) >> (uint(v232) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+16)) = uint32(v233)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v19
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+12)) = uint32(v69)
	v239 = int64(base.Ui64(v69) >> (uint(v232) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)) = uint32(v239)
	F_errdetail(m, int32(_a_F_WaitForWalSummarization_8), v15)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_WaitForWalSummarization_5), int32(747), int32(_a_F_WaitForWalSummarization_6))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
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
	var v15 int32
	_ = v15
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
	v15 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v21 = l0 + int32(8) + v15*int32(20)
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
	v36 = v15 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+648))
	if base.Ui32(v36) <= base.Ui32(v37) {
		v15 = v36
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
