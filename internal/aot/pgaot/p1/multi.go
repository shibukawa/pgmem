package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MultiXactIdIsRunning(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_GetMultiXactIdMembers(m, l0, v8+int32(12), l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if int32(0) < v12 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v21 = int32(0)
	goto L8
L4:
	;
	v191 = int32(0)
	goto L5
L5:
	;
	m.G0 = v8 + int32(16)
	return v191
L6:
	;
	F_pfree(m, v19)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L57
	}
L7:
	;
	v184 = int32(1)
	goto L6
L8:
	;
	v25 = int32(3)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v19+v21<<(uint(v25)%32))))
	if base.Ui32(v28) < base.Ui32(v25) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v166 = int32(0)
	goto L52
L10:
	;
	if v160 != 0 {
		goto L7
	} else {
		goto L50
	}
L11:
	;
	v160 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdIsRunning[0]))
	if v40 == v28 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v160 = int32(1)
	goto L10
L15:
	;
	goto L16
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdIsRunning[1]))
	if v44 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v160 = v150
	goto L10
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdIsRunning[2]))
	if v48 == int32(0) {
		v150 = int32(0)
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdIsRunning[3]))
	v120 = int32(0)
	v123 = v44 - int32(1)
	goto L40
L21:
	;
	v53 = v48
	goto L22
L22:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	if v59 == int32(4) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v150 = int32(0)
	goto L17
L24:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
	if v113 != 0 {
		v53 = v113
		goto L22
	} else {
		goto L39
	}
L25:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v62 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v65 = int32(1)
	if v28 == v62 {
		v150 = v65
		goto L17
	} else {
		goto L27
	}
L27:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
	v69 = v67 - int32(1)
	if v69 < int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v53)+48))
	v75 = int32(0)
	v78 = v69
	goto L29
L29:
	;
	v83 = int32(2)
	v84 = base.I32_div_s(v78-v75, v83)
	v85 = v84 + v75
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v72+v85<<(uint(v83)%32))))
	if v89 == v28 {
		v150 = v65
		goto L17
	} else {
		goto L31
	}
L30:
	;
	goto L24
L31:
	;
	v98 = base.B2i32(v89-v28 < int32(0)) | base.B2i32(base.Ui32(v89) < base.Ui32(int32(3)))
	if v98 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v85 + int32(1)
	goto L34
L33:
	;
	v99 = v75
	goto L34
L34:
	;
	if v98 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v102 = v78
	goto L37
L36:
	;
	v102 = v85 - int32(1)
	goto L37
L37:
	;
	if v99 <= v102 {
		v75 = v99
		v78 = v102
		goto L29
	} else {
		goto L38
	}
L38:
	;
	goto L30
L39:
	;
	goto L23
L40:
	;
	v128 = int32(2)
	v129 = base.I32_div_s(v123-v120, v128)
	v130 = v129 + v120
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v118+v130<<(uint(v128)%32))))
	v135 = base.B2i32(v134 == v28)
	if v134 == v28 {
		v150 = v135
		goto L17
	} else {
		goto L42
	}
L41:
	;
	v150 = v135
	goto L17
L42:
	;
	v138 = base.B2i32(base.Ui32(v134) < base.Ui32(v28))
	if base.Ui32(v134) < base.Ui32(v28) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v139 = v130 + int32(1)
	goto L45
L44:
	;
	v139 = v120
	goto L45
L45:
	;
	if base.Ui32(v134) < base.Ui32(v28) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v142 = v123
	goto L48
L47:
	;
	v142 = v130 - int32(1)
	goto L48
L48:
	;
	if v139 <= v142 {
		v120 = v139
		v123 = v142
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L41
L50:
	;
	v162 = v21 + int32(1)
	if v162 != v12 {
		v21 = v162
		goto L8
	} else {
		goto L51
	}
L51:
	;
	goto L9
L52:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v19+v166<<(uint(int32(3))%32))))
	v174 = F_TransactionIdIsInProgress(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	v184 = v174
	goto L6
L54:
	;
	if v174 != 0 {
		v184 = v174
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v177 = v166 + int32(1)
	if v177 != v12 {
		v166 = v177
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v191 = v184
	goto L5
}
func F_MultiXactIdSetOldestMember(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdSetOldestMember[0]))
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdSetOldestMember[1]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v2+v4<<(uint(int32(2))%32))))
	if v8 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdSetOldestMember[2]))
		v16 = F_LWLockAcquire(m, v12+int32(1664), int32(1))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdSetOldestMember[0]))
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdSetOldestMember[1]))
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdSetOldestMember[3]))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			*(*int32)(unsafe.Add(mBase, uint32(v19+v21<<(uint(int32(2))%32)))) = v27
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactIdSetOldestMember[2]))
			F_LWLockRelease(m, v30+int32(1664))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_MultiXactShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactShmemInit[0]))
	v6 = v4 + int32(56)
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactShmemInit[1])) = v6
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactShmemInit[2]))
	v11 = int32(2)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_MultiXactShmemInit[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_MultiXactShmemInit[4])) = v6 + v10<<(uint(v11)%32) + v15<<(uint(v11)%32)
	return
}
func F_SetMultiXactIdLimit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[0]))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[1]))
	v22 = F_LWLockAcquire(m, v18+int32(1664), int32(0))
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
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[2]))
	v26 = int32(1)
	v28 = l0 + int32(2147483647)
	if base.Ui32(v28) <= base.Ui32(v26) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v31 = v26
	goto L5
L4:
	;
	v31 = v28
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = v31
	v33 = int32(1)
	v34 = l0 + v16
	if base.Ui32(v34) <= base.Ui32(v33) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v37 = v33
	goto L8
L7:
	;
	v37 = v34
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+40)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = l0
	v42 = v31 - int32(_a_F_SetMultiXactIdLimit_0)
	if v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = v42
	goto L11
L10:
	;
	v44 = int32(-1)
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v44
	v47 = v31 - int32(100000000)
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v49 = v47
	goto L14
L13:
	;
	v49 = int32(-1)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[1]))
	F_LWLockRelease(m, v53+int32(1664))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v60 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v31
	F_errmsg_internal(m, int32(_a_F_SetMultiXactIdLimit_1), v13+int32(80))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[2]))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+16)))
	if v76 != int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	F_errfinish(m, int32(_a_F_SetMultiXactIdLimit_2), int32(2152), int32(_a_F_SetMultiXactIdLimit_3))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	m.G0 = v13 + int32(96)
	return
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = int64(0)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[1]))
	v86 = F_LWLockAcquire(m, v82+int32(_a_F_SetMultiXactIdLimit_4), int32(1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v88 = int32(1664)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[1]))
	v94 = F_LWLockAcquire(m, v90+v88, int32(1))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[2]))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v97)+8))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[1]))
	F_LWLockRelease(m, v102+int32(1664))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v100 == v99 {
		v132 = v98
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[1]))
	F_LWLockRelease(m, v170+v167)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L43
	}
L28:
	;
	v149 = int32(_a_F_SetMultiXactIdLimit_4)
	v152 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L39
	}
L29:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[1]))
	F_LWLockRelease(m, v134+int32(_a_F_SetMultiXactIdLimit_4))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L37
	}
L30:
	;
	v110 = F_find_multixact_start(m, v100, v13+int32(88))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v110 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v116 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+88))
	if v116 == int32(0) {
		v132 = v118
		goto L29
	} else {
		goto L34
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v118
	F_errmsg_internal(m, int32(_a_F_SetMultiXactIdLimit_13), v13+int32(48))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_SetMultiXactIdLimit_2), int32(2475), int32(_a_F_SetMultiXactIdLimit_12))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v132 = v118
	goto L29
L37:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[1]))
	v144 = F_LWLockAcquire(m, v140+int32(1664), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v147)+32)) = v132
	v167 = v88
	goto L27
L39:
	;
	if v152 == int32(0) {
		v167 = v149
		goto L27
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v100
	F_errmsg(m, int32(_a_F_SetMultiXactIdLimit_11), v13-int32(-64))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_SetMultiXactIdLimit_2), int32(2479), int32(_a_F_SetMultiXactIdLimit_12))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v167 = v149
	goto L27
L43:
	;
	if int32(0) <= v37-v51 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if int32(0) <= v49-v51 {
		goto L22
	} else {
		goto L51
	}
L45:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[3])))
	if v178&int32(1) == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[3])))
	if v185 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L44
L48:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v189+int32(16)))) = int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[5]))
	v198 = F_pgmem_kill(m, v196, int32(10))
	mBase = m.M
	goto L50
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_SetMultiXactIdLimit[6]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+20))
	goto L54
L52:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v13))) = base.F64_mul(base.F64_div(base.F64_convert_i32_u(v246), float64(2.147483647e+09)), float64(100))
	v256 = F_errdetail(m, int32(_a_F_SetMultiXactIdLimit_7), v13)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L64
	}
L53:
	;
	v232 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L61
	}
L54:
	;
	if base.B2i32(v204 == int32(2)) == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v209 = F_get_database_name(m, l1)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if v209 == int32(0) {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v215 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v215 == int32(0) {
		goto L22
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v209
	v220 = v31 - v51
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v220
	F_errmsg_plural(m, int32(_a_F_SetMultiXactIdLimit_9), int32(_a_F_SetMultiXactIdLimit_10), v220, v13+int32(32))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v246 = v220
	v248 = int32(2214)
	goto L52
L61:
	;
	if v232 == int32(0) {
		goto L22
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	v237 = v31 - v51
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v237
	F_errmsg_plural(m, int32(_a_F_SetMultiXactIdLimit_5), int32(_a_F_SetMultiXactIdLimit_6), v237, v13+int32(16))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v246 = v237
	v248 = int32(2225)
	goto L52
L64:
	;
	F_errhint(m, int32(_a_F_SetMultiXactIdLimit_8), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_SetMultiXactIdLimit_2), v248, int32(_a_F_SetMultiXactIdLimit_3))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	goto L22
}
func F_multi_sort_add_dimension(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v5 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_multi_sort_add_dimension[0]))
	v10 = l0 + l1*int32(36)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+13)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l3
	v15 = v10 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v7
	F_PrepareSortSupportFromOrderingOp(m, l2, v15)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		return
	}
}
