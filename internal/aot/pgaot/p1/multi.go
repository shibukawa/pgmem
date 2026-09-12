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
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
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
	v178 = int32(0)
	goto L5
L5:
	;
	m.G0 = v8 + int32(16)
	return v178
L6:
	;
	F_pfree(m, v19)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L57
	}
L7:
	;
	v171 = int32(1)
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
	v154 = int32(0)
	goto L52
L10:
	;
	if v148 != 0 {
		goto L7
	} else {
		goto L50
	}
L11:
	;
	v148 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v39 == v28 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v148 = int32(1)
	goto L10
L15:
	;
	goto L16
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v43 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v148 = v140
	goto L10
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v47 == int32(0) {
		v140 = int32(0)
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v111 = int32(0)
	v113 = v43 - int32(1)
	goto L40
L21:
	;
	v52 = v47
	goto L22
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	if v57 == int32(4) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v140 = int32(0)
	goto L17
L24:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v52)+80))
	if v104 != 0 {
		v52 = v104
		goto L22
	} else {
		goto L39
	}
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v60 == int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v63 = int32(1)
	if v28 == v60 {
		v140 = v63
		goto L17
	} else {
		goto L27
	}
L27:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
	v67 = v65 - int32(1)
	if v67 < int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v72 = int32(0)
	v74 = v67
	goto L29
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
	v80 = int32(2)
	v81 = base.I32_div_s(v74-v72, v80)
	v82 = v81 + v72
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v78+v82<<(uint(v80)%32))))
	if v86 == v28 {
		v140 = v63
		goto L17
	} else {
		goto L31
	}
L30:
	;
	goto L24
L31:
	;
	v90 = F_TransactionIdPrecedes(m, v86, v28)
	mBase = m.M
	if v90 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v91 = v82 + int32(1)
	goto L34
L33:
	;
	v91 = v72
	goto L34
L34:
	;
	if v90 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v94 = v74
	goto L37
L36:
	;
	v94 = v82 - int32(1)
	goto L37
L37:
	;
	if v91 <= v94 {
		v72 = v91
		v74 = v94
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
	v118 = int32(2)
	v119 = base.I32_div_s(v113-v111, v118)
	v120 = v119 + v111
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v109+v120<<(uint(v118)%32))))
	v125 = base.B2i32(v124 == v28)
	if v124 == v28 {
		v140 = v125
		goto L17
	} else {
		goto L42
	}
L41:
	;
	v140 = v125
	goto L17
L42:
	;
	v128 = base.B2i32(base.Ui32(v124) < base.Ui32(v28))
	if base.Ui32(v124) < base.Ui32(v28) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v129 = v120 + int32(1)
	goto L45
L44:
	;
	v129 = v111
	goto L45
L45:
	;
	if base.Ui32(v124) < base.Ui32(v28) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v132 = v113
	goto L48
L47:
	;
	v132 = v120 - int32(1)
	goto L48
L48:
	;
	if v129 <= v132 {
		v111 = v129
		v113 = v132
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L41
L50:
	;
	v150 = v21 + int32(1)
	if v150 != v12 {
		v21 = v150
		goto L8
	} else {
		goto L51
	}
L51:
	;
	goto L9
L52:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v19+v154<<(uint(int32(3))%32))))
	v162 = F_TransactionIdIsInProgress(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	v171 = v162
	goto L6
L54:
	;
	if v162 != 0 {
		v171 = v162
		goto L6
	} else {
		goto L55
	}
L55:
	;
	v165 = v154 + int32(1)
	if v165 != v12 {
		v154 = v165
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v178 = v171
	goto L5
}
func F_MultiXactIdSetOldestMember(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	v5 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v3+v5<<(uint(int32(2))%32))))
	if v9 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, _consts[24]))
		v17 = F_LWLockAcquire(m, v13+int32(1664), int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _consts[128]))
			v22 = *(*int32)(unsafe.Add(mBase, _consts[126]))
			v26 = int32(1)
			v28 = *(*int32)(unsafe.Add(mBase, _consts[129]))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
			if base.Ui32(v29) <= base.Ui32(v26) {
				v32 = v26
			} else {
				v32 = v29
			}
			*(*int32)(unsafe.Add(mBase, uint32(v20+v22<<(uint(int32(2))%32)))) = v32
			v35 = *(*int32)(unsafe.Add(mBase, _consts[24]))
			F_LWLockRelease(m, v35+int32(1664))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_SetMultiXactIdLimit(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	v22 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v26 = F_LWLockAcquire(m, v22+int32(1664), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	v30 = int32(1)
	v32 = l0 + int32(2147483647)
	if base.Ui32(v32) <= base.Ui32(v30) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v35 = v30
	goto L5
L4:
	;
	v35 = v32
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v35
	v37 = int32(1)
	v38 = l0 + v20
	if base.Ui32(v38) <= base.Ui32(v37) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v41 = v37
	goto L8
L7:
	;
	v41 = v38
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = l0
	v46 = v35 - int32(3000000)
	if v46 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v48 = v46
	goto L11
L10:
	;
	v48 = int32(-1)
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v48
	v51 = v35 - int32(40000000)
	if v51 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v53 = v51
	goto L14
L13:
	;
	v53 = int32(-1)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v57 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v57+int32(1664))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v64 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v64 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+84)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v35
	F_errmsg_internal(m, int32(56061), v17+int32(80))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+8)))
	if v80 != int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	F_errfinish(m, int32(493284), int32(2503), int32(101173))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	m.G0 = v17 + int32(96)
	return
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v90 = F_LWLockAcquire(m, v86+int32(5248), int32(1))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v97 = F_LWLockAcquire(m, v93+int32(1664), int32(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+44))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)+20))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+24)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v108 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v108+int32(1664))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v106 == v105 {
		v138 = v104
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	v221 = F_LWLockAcquire(m, v217+int32(1664), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L58
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v167
	F_errmsg_internal(m, int32(42111), v17+int32(32))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L56
	}
L29:
	;
	v175 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L48
	}
L30:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v141+int32(5248))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L38
	}
L31:
	;
	v116 = F_find_multixact_start(m, v106, v17+int32(92))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v116 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v122 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	if v122 == int32(0) {
		v138 = v124
		goto L30
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v124
	F_errmsg_internal(m, int32(41211), v17+int32(48))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(493284), int32(2842), int32(101152))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v138 = v124
	goto L30
L38:
	;
	v147 = base.I32_rem_u_s(v138, int32(52352))
	if l2 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v167 = v138 - v147 - int32(52352)
	v168 = int32(1)
	v171 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L46
	}
L40:
	;
	if v103&int32(1) != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v153 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v153 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(455108), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(493284), int32(2867), int32(101152))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L39
L46:
	;
	if v171 != 0 {
		goto L28
	} else {
		goto L47
	}
L47:
	;
	v212 = v168
	v214 = v138
	v215 = v167
	goto L27
L48:
	;
	if v175 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v106
	F_errmsg(m, int32(314555), v17-int32(-64))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v189+int32(5248))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	F_errfinish(m, int32(493284), int32(2846), int32(101152))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v194 = int32(1)
	if v103&v194 != 0 {
		v212 = v194
		v214 = v102
		v215 = v101
		goto L27
	} else {
		goto L55
	}
L55:
	;
	v197 = int32(0)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	v212 = v197
	v214 = v198
	v215 = v197
	goto L27
L56:
	;
	F_errfinish(m, int32(493284), int32(2871), int32(101152))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v212 = v168
	v214 = v138
	v215 = v167
	goto L27
L58:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	*(*int32)(unsafe.Add(mBase, uint32(v224)+44)) = v215
	*(*uint8)(unsafe.Add(mBase, uint32(v224)+24)) = uint8(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = v214
	v229 = *(*int32)(unsafe.Add(mBase, _consts[24]))
	F_LWLockRelease(m, v229+int32(1664))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	if v41-v55 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if int32(0) <= v53-v55 {
		goto L22
	} else {
		goto L67
	}
L61:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, _consts[131])))
	if v243 != int32(1) {
		goto L60
	} else {
		goto L65
	}
L62:
	;
	if v212 == int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	if int32(0) <= v104-v214 {
		goto L60
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	F_SendPostmasterSignal(m, int32(4))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	goto L60
L67:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+20))
	goto L70
L68:
	;
	F_errhint(m, int32(584159), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L80
	}
L69:
	;
	v282 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L77
	}
L70:
	;
	if base.B2i32(v254 == int32(2)) == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v259 = F_get_database_name(m, l1)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v259 == int32(0) {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v265 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v265 == int32(0) {
		goto L22
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v259
	v270 = v35 - v55
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v270
	F_errmsg_plural(m, int32(449276), int32(449607), v270, v17+int32(16))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v295 = int32(2558)
	goto L68
L77:
	;
	if v282 == int32(0) {
		goto L22
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l1
	v287 = v35 - v55
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v287
	F_errmsg_plural(m, int32(449203), int32(449532), v287, v17)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v295 = int32(2567)
	goto L68
L80:
	;
	F_errfinish(m, int32(493284), v295, int32(101173))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
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
	v7 = *(*int32)(unsafe.Add(mBase, _consts[3]))
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
