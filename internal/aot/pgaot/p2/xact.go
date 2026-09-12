package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_XactLockTableWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(1109)
	v15 = int32(4508392)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v8 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v8 + int32(20)
	goto L3
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l0
	v34 = int32(0)
	v36 = F_LockAcquire(m, v8+int32(32), int32(5), v34, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v42 = F_LockRelease(m, v8+int32(32), int32(5), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v44 = F_TransactionIdIsInProgress(m, l0)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	if l3 != 0 {
		goto L22
	} else {
		goto L23
	}
L8:
	;
	if v44 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v48 = l0
	goto L10
L10:
	;
	v53 = F_SubTransGetTopmostTransaction(m, v48)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v53
	v63 = int32(0)
	v65 = F_LockAcquire(m, v8+int32(32), int32(5), v63, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v71 = F_LockRelease(m, v8+int32(32), int32(5), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v73 = F_TransactionIdIsInProgress(m, v53)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if v73 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v78 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_pg_usleep(m, int32(1000))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v48 = v53
	goto L10
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v90
	goto L24
L23:
	;
	goto L24
L24:
	;
	m.G0 = v8 + int32(48)
	return
}
func F_XactLogCommitRecord(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int64
	_ = v118
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	v14 = int32(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v14
	if l9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = int32(1073741824)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v23
	v26 = v23
	goto L3
L2:
	;
	v26 = v14
	goto L3
L3:
	;
	v28 = l10 & int32(2)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _consts[171])))
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = v26 | int32(-2147483648)
	goto L6
L5:
	;
	v35 = v26
	goto L6
L6:
	;
	v36 = v28<<(uint(int32(5))%32) | v35
	if v34|v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v36
	goto L9
L8:
	;
	goto L9
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if int32(4) <= v40 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = v36 | int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v44
	v46 = v44
	goto L12
L11:
	;
	v46 = v36
	goto L12
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v51 = int32(0)
	if base.B2i32(v48 < int32(2))&base.B2i32(l7 <= v51) == v51 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v57 = v46 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v57
	v60 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v60
	v63 = *(*int32)(unsafe.Add(mBase, _consts[173]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v63
	v65 = v57
	goto L15
L14:
	;
	v65 = v46
	goto L15
L15:
	;
	if int32(0) < l1 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = l1
	v70 = v65 | int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v70
	v72 = v70
	goto L18
L17:
	;
	v72 = v65
	goto L18
L18:
	;
	if l11 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v75 = int32(48)
	goto L21
L20:
	;
	v75 = int32(0)
	goto L21
L21:
	;
	if int32(0) < l3 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = l3
	v80 = v72 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v80
	v84 = v75 | int32(1)
	v85 = v80
	goto L24
L23:
	;
	v84 = v75
	v85 = v72
	goto L24
L24:
	;
	if int32(0) < l5 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = l5
	v90 = v85 | int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v90
	v92 = v90
	goto L27
L26:
	;
	v92 = v85
	goto L27
L27:
	;
	if int32(0) < l7 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = l7
	v97 = v92 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v97
	v99 = v97
	goto L30
L29:
	;
	v99 = v92
	goto L30
L30:
	;
	if l11 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, _consts[174])))
	if v113 != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v111 = v99
	goto L31
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = l11
	v104 = v99 | int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v104
	if v48 < int32(2) {
		v111 = v104
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v109 = v99 | int32(144)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v109
	v111 = v109
	goto L31
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v111 | int32(32)
	v118 = *(*int64)(unsafe.Add(mBase, _consts[175]))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v118
	v121 = *(*int64)(unsafe.Add(mBase, _consts[176]))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v121
	v124 = int32(1)
	goto L38
L37:
	;
	v124 = v111
	goto L38
L38:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	return int64(0)
L40:
	;
	F_XLogRegisterData(m, v16+int32(-8), int32(8))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	if v124 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v235 = int32(4411332)
	v237 = int32(*(*uint8)(unsafe.Add(mBase, _consts[64])))
	v238 = v237 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[64])) = uint8(v238)
	goto L76
L43:
	;
	F_XLogRegisterData(m, v16+int32(-12), int32(4))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v141&int32(1) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_XLogRegisterData(m, v16+int32(-20), int32(8))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L39
	} else {
		goto L48
	}
L46:
	;
	v150 = v141
	goto L47
L47:
	;
	if v150&int32(2) != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v150 = v149
	goto L47
L49:
	;
	F_XLogRegisterData(m, v16+int32(-24), int32(4))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L39
	} else {
		goto L52
	}
L50:
	;
	v163 = v150
	goto L51
L51:
	;
	if v163&int32(4) != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	F_XLogRegisterData(m, l2, l1<<(uint(int32(2))%32))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L39
	} else {
		goto L53
	}
L53:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v163 = v162
	goto L51
L54:
	;
	F_XLogRegisterData(m, v16+int32(-28), int32(4))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L39
	} else {
		goto L57
	}
L55:
	;
	v176 = v163
	goto L56
L56:
	;
	if v176&int32(256) != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	F_XLogRegisterData(m, l4, l3*int32(12))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L39
	} else {
		goto L58
	}
L58:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v176 = v175
	goto L56
L59:
	;
	F_XLogRegisterData(m, v16+int32(-32), int32(4))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L39
	} else {
		goto L62
	}
L60:
	;
	v189 = v176
	goto L61
L61:
	;
	if v189&int32(8) != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	F_XLogRegisterData(m, l6, l5<<(uint(int32(4))%32))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L39
	} else {
		goto L63
	}
L63:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v189 = v188
	goto L61
L64:
	;
	F_XLogRegisterData(m, v16+int32(-36), int32(4))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L39
	} else {
		goto L67
	}
L65:
	;
	v202 = v189
	goto L66
L66:
	;
	if v202&int32(16) == int32(0) {
		v223 = v202
		goto L69
	} else {
		goto L70
	}
L67:
	;
	F_XLogRegisterData(m, l8, l7<<(uint(int32(4))%32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L39
	} else {
		goto L68
	}
L68:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v202 = v201
	goto L66
L69:
	;
	if v223&int32(32) == int32(0) {
		goto L42
	} else {
		goto L74
	}
L70:
	;
	F_XLogRegisterData(m, v16+int32(-40), int32(4))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L39
	} else {
		goto L71
	}
L71:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	if v212&int32(128) == int32(0) {
		v223 = v212
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v217 = F_strlen(m, l12)
	mBase = m.M
	F_XLogRegisterData(m, l12, v217+int32(1))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L39
	} else {
		goto L73
	}
L73:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v223 = v222
	goto L69
L74:
	;
	F_XLogRegisterData(m, v16+int32(-56), int32(16))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L39
	} else {
		goto L75
	}
L75:
	;
	goto L42
L76:
	;
	if v124 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v243 = v84 | int32(128)
	goto L79
L78:
	;
	v243 = v84
	goto L79
L79:
	;
	v244 = F_XLogInsert(m, int32(1), v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L39
	} else {
		goto L80
	}
L80:
	;
	m.G0 = v18 - int32(-64)
	return v244
}
func F_xact_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v288 int64
	_ = v288
	var v289 int64
	_ = v289
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int64
	_ = v342
	var v343 int64
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int64
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int64
	_ = v451
	var v452 int64
	_ = v452
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v467 int64
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v576 int32
	_ = v576
	var v579 int64
	_ = v579
	var v580 int64
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int64
	_ = v586
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v621 int32
	_ = v621
	var v622 int64
	_ = v622
	var v623 int64
	_ = v623
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v641 int64
	_ = v641
	var v643 int64
	_ = v643
	var v645 int64
	_ = v645
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v819 int32
	_ = v819
	var v820 int64
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v969 int32
	_ = v969
	var v970 int64
	_ = v970
	var v971 int64
	_ = v971
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1020 int32
	_ = v1020
	var v1021 int64
	_ = v1021
	var v1022 int64
	_ = v1022
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1099 int64
	_ = v1099
	var v1100 int64
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 int64
	_ = v1106
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1131 int32
	_ = v1131
	var v1145 int32
	_ = v1145
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	v13 = m.G0
	v15 = v13 - int32(784)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+48)))
	switch int32(base.Ui32(v19)>>(uint(int32(4))%32))&int32(7) - int32(1) {
	case 0:
		goto L4
	case 1, 3:
		goto L5
	default:
		goto L6
	case 4:
		goto L3
	case 5:
		goto L2
	case 6:
		goto L1
	}
L1:
	;
	m.G0 = v15 + int32(784)
	return
L2:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v1159 = int32(0)
	F_standby_desc_invalidations(m, l0, v1156, v18+int32(4), v1159, v1159, v1159)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L32
	} else {
		goto L209
	}
L3:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+416)) = v1114
	F_appendStringInfo(m, l0, int32(745516), v15+int32(416))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L32
	} else {
		goto L202
	}
L4:
	;
	v634 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+56)))
	v640 = F__emscripten_memset_bulkmem(m, v15+int32(432), base.I32_extend8_s(int32(0)), int32(248))
	mBase = m.M
	goto L126
L5:
	;
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+56)))
	v360 = int32(0)
	v365 = F___memset(m, v15+int32(424), v360, int32(264))
	mBase = m.M
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v365))) = v366
	if v360 <= base.I32_extend8_s(v19&int32(255)) {
		goto L74
	} else {
		goto L75
	}
L6:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+56)))
	v31 = int32(0)
	v36 = F___memset(m, v15+int32(424), v31, int32(288))
	mBase = m.M
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v37
	if v31 <= base.I32_extend8_s(v19&int32(255)) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v15)+476))
	if v145 != 0 {
		goto L29
	} else {
		goto L30
	}
L8:
	;
	goto L7
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v42
	if v42&int32(1) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v48
	v54 = v18 + int32(20)
	goto L12
L11:
	;
	v54 = v18 + int32(12)
	goto L12
L12:
	;
	if v42&int32(2) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v59 = v54 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v57
	v65 = v59 + v57<<(uint(int32(2))%32)
	goto L15
L14:
	;
	v65 = v54
	goto L15
L15:
	;
	if v42&int32(4) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v71 = v65 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = v69
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v78 = v71 + v74*int32(12)
	goto L18
L17:
	;
	v78 = v65
	goto L18
L18:
	;
	if v42&int32(256) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v84 = int32(4)
	v85 = v78 + v84
	*(*int32)(unsafe.Add(mBase, uint32(v36)+40)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v36)+36)) = v83
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v92 = v85 + v88<<(uint(v84)%32)
	goto L21
L20:
	;
	v92 = v78
	goto L21
L21:
	;
	if v42&int32(8) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v98 = int32(4)
	v99 = v92 + v98
	*(*int32)(unsafe.Add(mBase, uint32(v36)+48)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v36)+44)) = v97
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v106 = v99 + v102<<(uint(v98)%32)
	goto L24
L23:
	;
	v106 = v92
	goto L24
L24:
	;
	if v42&int32(16) == int32(0) {
		v130 = v42
		v131 = v106
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v130&int32(32) == int32(0) {
		goto L8
	} else {
		goto L28
	}
L26:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+52)) = v113
	v116 = v106 + int32(4)
	if v42&int32(128) == int32(0) {
		v130 = v42
		v131 = v116
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v124 = F_strlcpy(m, v36+int32(56), v116, int32(200))
	mBase = m.M
	v125 = F_strlen(m, v116)
	mBase = m.M
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v130 = v129
	v131 = v125 + v116 + int32(1)
	goto L25
L28:
	;
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v131)))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v131)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v36)+280)) = v137
	*(*int64)(unsafe.Add(mBase, uint32(v36)+272)) = v136
	goto L8
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v145
	F_appendStringInfo(m, l0, int32(745521), v15+int32(96))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v153 = F_timestamptz_to_str(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L32
	} else {
		goto L34
	}
L32:
	;
	return
L33:
	;
	goto L31
L34:
	;
	F_appendStringInfoString(m, l0, v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v15)+452))
	if int32(0) < v157 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v15)+456))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = int32(153122)
	F_appendStringInfo(m, l0, int32(547348), v15+int32(80))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L32
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	if int32(0) < v216 {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	v170 = int32(0)
	goto L40
L40:
	;
	v185 = v160 + v170*int32(12)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
	F_GetRelationPath(m, v15+int32(712), v186, v187, v188, int32(-1), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L32
	} else {
		goto L42
	}
L41:
	;
	goto L38
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v15 + int32(712)
	F_appendStringInfo(m, l0, int32(206163), v15-int32(-64))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L32
	} else {
		goto L43
	}
L43:
	;
	v202 = v170 + int32(1)
	if v202 != v157 {
		v170 = v202
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v15)+448))
	F_appendStringInfoString(m, l0, int32(547229))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L32
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v15)+460))
	if int32(0) < v261 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v225 = int32(0)
	goto L49
L49:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v219+v225<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v239
	F_appendStringInfo(m, l0, int32(59440), v15+int32(48))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L32
	} else {
		goto L51
	}
L50:
	;
	goto L47
L51:
	;
	v247 = v225 + int32(1)
	if v247 != v216 {
		v225 = v247
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v15)+464))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(757461)
	F_appendStringInfo(m, l0, int32(547253), v15+int32(32))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L32
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v15)+468))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v15)+472))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v15)+436))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v15)+440))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v15)+432))
	F_standby_desc_invalidations(m, l0, v312, v313, v314, v315, int32(base.Ui32(v316&int32(1073741824))>>(uint(int32(30))%32)))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L32
	} else {
		goto L61
	}
L56:
	;
	v274 = int32(0)
	goto L57
L57:
	;
	v287 = v264 + v274<<(uint(int32(4))%32)
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v287)))
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v287)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v289
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v288
	F_appendStringInfo(m, l0, int32(37913), v15+int32(16))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L32
	} else {
		goto L59
	}
L58:
	;
	goto L55
L59:
	;
	v298 = v274 + int32(1)
	if v298 != v261 {
		v274 = v298
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v15)+432))
	if v323&int32(536870912) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_appendStringInfoString(m, l0, int32(319489))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L32
	} else {
		goto L65
	}
L63:
	;
	v330 = v323
	goto L64
L64:
	;
	if v330 < int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v15)+432))
	v330 = v329
	goto L64
L66:
	;
	F_appendStringInfoString(m, l0, int32(489362))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L32
	} else {
		goto L69
	}
L67:
	;
	v337 = v330
	goto L68
L68:
	;
	if v337&int32(32) == int32(0) {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v15)+432))
	v337 = v336
	goto L68
L70:
	;
	v342 = *(*int64)(unsafe.Add(mBase, uint32(v15)+696))
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v15)+704))
	v344 = F_timestamptz_to_str(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L32
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v344
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)) = uint32(v342)
	v349 = int64(base.Ui64(v342) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+4)) = uint32(v349)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v26
	F_appendStringInfo(m, l0, int32(179766), v15)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L32
	} else {
		goto L72
	}
L72:
	;
	goto L1
L73:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v15)+468))
	if v460 != 0 {
		goto L92
	} else {
		goto L93
	}
L74:
	;
	goto L73
L75:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v365)+8)) = v371
	if v371&int32(1) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v365)+12)) = v375
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v365)+16)) = v377
	v383 = v18 + int32(20)
	goto L78
L77:
	;
	v383 = v18 + int32(12)
	goto L78
L78:
	;
	if v371&int32(2) != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	v388 = v383 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v365)+24)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v365)+20)) = v386
	v394 = v388 + v386<<(uint(int32(2))%32)
	goto L81
L80:
	;
	v394 = v383
	goto L81
L81:
	;
	if v371&int32(4) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	v400 = v394 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v365)+32)) = v400
	*(*int32)(unsafe.Add(mBase, uint32(v365)+28)) = v398
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	v407 = v400 + v403*int32(12)
	goto L84
L83:
	;
	v407 = v394
	goto L84
L84:
	;
	if v371&int32(256) != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v413 = int32(4)
	v414 = v407 + v413
	*(*int32)(unsafe.Add(mBase, uint32(v365)+40)) = v414
	*(*int32)(unsafe.Add(mBase, uint32(v365)+36)) = v412
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v421 = v414 + v417<<(uint(v413)%32)
	goto L87
L86:
	;
	v421 = v407
	goto L87
L87:
	;
	if v371&int32(16) == int32(0) {
		v445 = v371
		v446 = v421
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v445&int32(32) == int32(0) {
		goto L74
	} else {
		goto L91
	}
L89:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	*(*int32)(unsafe.Add(mBase, uint32(v365)+44)) = v428
	v431 = v421 + int32(4)
	if v371&int32(128) == int32(0) {
		v445 = v371
		v446 = v431
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v439 = F_strlcpy(m, v365+int32(48), v431, int32(200))
	mBase = m.M
	v440 = F_strlen(m, v431)
	mBase = m.M
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v365)+8))
	v445 = v444
	v446 = v440 + v431 + int32(1)
	goto L88
L91:
	;
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v446)))
	v452 = *(*int64)(unsafe.Add(mBase, uint32(v446)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v365)+256)) = v452
	*(*int64)(unsafe.Add(mBase, uint32(v365)+248)) = v451
	goto L74
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v460
	F_appendStringInfo(m, l0, int32(745521), v15+int32(208))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L32
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v467 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v468 = F_timestamptz_to_str(m, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L32
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	F_appendStringInfoString(m, l0, v468)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L32
	} else {
		goto L97
	}
L97:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v15)+452))
	if int32(0) < v472 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v15)+456))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = int32(153122)
	F_appendStringInfo(m, l0, int32(547348), v15+int32(192))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L32
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	if int32(0) < v531 {
		goto L107
	} else {
		goto L108
	}
L101:
	;
	v485 = int32(0)
	goto L102
L102:
	;
	v500 = v475 + v485*int32(12)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)+4))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v500)+8))
	F_GetRelationPath(m, v15+int32(712), v501, v502, v503, int32(-1), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L32
	} else {
		goto L104
	}
L103:
	;
	goto L100
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v15 + int32(712)
	F_appendStringInfo(m, l0, int32(206163), v15+int32(176))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L32
	} else {
		goto L105
	}
L105:
	;
	v517 = v485 + int32(1)
	if v517 != v472 {
		v485 = v517
		goto L102
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v15)+448))
	F_appendStringInfoString(m, l0, int32(547229))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L32
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+432)))
	if v576&int32(32) != 0 {
		goto L115
	} else {
		goto L116
	}
L110:
	;
	v540 = int32(0)
	goto L111
L111:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v534+v540<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v554
	F_appendStringInfo(m, l0, int32(59440), v15+int32(160))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L32
	} else {
		goto L113
	}
L112:
	;
	goto L109
L113:
	;
	v562 = v540 + int32(1)
	if v562 != v531 {
		v540 = v562
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v579 = *(*int64)(unsafe.Add(mBase, uint32(v15)+672))
	v580 = *(*int64)(unsafe.Add(mBase, uint32(v15)+680))
	v581 = F_timestamptz_to_str(m, v580)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L32
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v15)+460))
	if v595 <= int32(0) {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+156)) = v581
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+152)) = uint32(v579)
	v586 = int64(base.Ui64(v579) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+148)) = uint32(v586)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v355
	F_appendStringInfo(m, l0, int32(179766), v15+int32(144))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L32
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v15)+464))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = int32(757461)
	F_appendStringInfo(m, l0, int32(547253), v15+int32(128))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L32
	} else {
		goto L121
	}
L121:
	;
	v608 = int32(0)
	goto L122
L122:
	;
	v621 = v598 + v608<<(uint(int32(4))%32)
	v622 = *(*int64)(unsafe.Add(mBase, uint32(v621)))
	v623 = *(*int64)(unsafe.Add(mBase, uint32(v621)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v623
	*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v622
	F_appendStringInfo(m, l0, int32(37913), v15+int32(112))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L32
	} else {
		goto L124
	}
L123:
	;
	goto L1
L124:
	;
	v632 = v608 + int32(1)
	if v632 != v595 {
		v608 = v632
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v641 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+424)) = v641
	v643 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+696)) = v643
	v645 = *(*int64)(unsafe.Add(mBase, uint32(v18)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+704)) = v645
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+476)) = v647
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+436)) = v649
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+444)) = v651
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+452)) = v653
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+680)) = v655
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+460)) = v657
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+688)) = v659
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+468)) = v661
	v664 = v15 + int32(480)
	v666 = v18 + int32(72)
	v667 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+54)))
	if (v666^v664)&int32(3) != 0 {
		v737 = v666
		v738 = v667
		v739 = v664
		goto L131
	} else {
		goto L132
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+384)) = v664
	v777 = int32(7)
	v781 = v666 + (v667+v777)&int32(131064)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+448)) = v781
	v787 = int32(-8)
	v789 = v781 + (v651<<(uint(int32(2))%32)+v777)&v787
	*(*int32)(unsafe.Add(mBase, uint32(v15)+456)) = v789
	v791 = int32(12)
	v797 = v789 + (v653*v791+v777)&v787
	*(*int32)(unsafe.Add(mBase, uint32(v15)+684)) = v797
	v805 = v797 + (v655*v791+v777)&v787
	*(*int32)(unsafe.Add(mBase, uint32(v15)+464)) = v805
	v807 = int32(4)
	v809 = v805 + v657<<(uint(v807)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+692)) = v809
	*(*int32)(unsafe.Add(mBase, uint32(v15)+472)) = v809 + v659<<(uint(v807)%32)
	F_appendStringInfo(m, l0, int32(745615), v15+int32(384))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L32
	} else {
		goto L153
	}
L128:
	;
	v775 = F___memset(m, v772, int32(0), v771)
	mBase = m.M
	goto L127
L129:
	;
	v771 = int32(0)
	v772 = v766
	goto L128
L130:
	;
	v749 = v744
	v750 = v745
	v751 = v746
	goto L149
L131:
	;
	if v738 == int32(0) {
		v766 = v739
		goto L129
	} else {
		goto L148
	}
L132:
	;
	v673 = int32(0)
	v674 = base.B2i32(v667 != v673)
	if v666&int32(3) == v673 {
		v703 = v666
		v704 = v667
		v705 = v664
		v706 = v674
		goto L133
	} else {
		goto L134
	}
L133:
	;
	if v706 == int32(0) {
		v766 = v705
		goto L129
	} else {
		goto L141
	}
L134:
	;
	if v667 == int32(0) {
		v703 = v666
		v704 = v667
		v705 = v664
		v706 = v674
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v682 = v666
	v683 = v667
	v684 = v664
	goto L136
L136:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
	*(*uint8)(unsafe.Add(mBase, uint32(v684))) = uint8(v686)
	if v686 == int32(0) {
		v771 = v683
		v772 = v684
		goto L128
	} else {
		goto L138
	}
L137:
	;
	v703 = v697
	v704 = v693
	v705 = v691
	v706 = v695
	goto L133
L138:
	;
	v690 = int32(1)
	v691 = v684 + v690
	v693 = v683 - v690
	v694 = int32(0)
	v695 = base.B2i32(v693 != v694)
	v697 = v682 + v690
	if v697&int32(3) == v694 {
		v703 = v697
		v704 = v693
		v705 = v691
		v706 = v695
		goto L133
	} else {
		goto L139
	}
L139:
	;
	if v693 != 0 {
		v682 = v697
		v683 = v693
		v684 = v691
		goto L136
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703))))
	if v709 == int32(0) {
		v771 = v704
		v772 = v705
		goto L128
	} else {
		goto L142
	}
L142:
	;
	if base.Ui32(v704) < base.Ui32(int32(4)) {
		v737 = v703
		v738 = v704
		v739 = v705
		goto L131
	} else {
		goto L143
	}
L143:
	;
	v715 = v703
	v716 = v704
	v717 = v705
	goto L144
L144:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	v723 = int32(-2139062144)
	if (int32(16843008)-v720|v720)&v723 != v723 {
		v744 = v715
		v745 = v716
		v746 = v717
		goto L130
	} else {
		goto L146
	}
L145:
	;
	v737 = v731
	v738 = v733
	v739 = v729
	goto L131
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v717))) = v720
	v728 = int32(4)
	v729 = v717 + v728
	v731 = v715 + v728
	v733 = v716 - v728
	if base.Ui32(int32(3)) < base.Ui32(v733) {
		v715 = v731
		v716 = v733
		v717 = v729
		goto L144
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v744 = v737
	v745 = v738
	v746 = v739
	goto L130
L149:
	;
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749))))
	*(*uint8)(unsafe.Add(mBase, uint32(v751))) = uint8(v753)
	if v753 == int32(0) {
		v771 = v750
		v772 = v751
		goto L128
	} else {
		goto L151
	}
L150:
	;
	v766 = v758
	goto L129
L151:
	;
	v757 = int32(1)
	v758 = v751 + v757
	v762 = v750 - v757
	if v762 != 0 {
		v749 = v749 + v757
		v750 = v762
		v751 = v758
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	v820 = *(*int64)(unsafe.Add(mBase, uint32(v15)+424))
	v821 = F_timestamptz_to_str(m, v820)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L32
	} else {
		goto L154
	}
L154:
	;
	F_appendStringInfoString(m, l0, v821)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L32
	} else {
		goto L155
	}
L155:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v15)+452))
	if int32(0) < v825 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v15)+456))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+368)) = int32(672500)
	F_appendStringInfo(m, l0, int32(547348), v15+int32(368))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L32
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v15)+680))
	if int32(0) < v884 {
		goto L165
	} else {
		goto L166
	}
L159:
	;
	v838 = int32(0)
	goto L160
L160:
	;
	v853 = v828 + v838*int32(12)
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v853)+4))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v853)))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v853)+8))
	F_GetRelationPath(m, v15+int32(712), v854, v855, v856, int32(-1), int32(0))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L32
	} else {
		goto L162
	}
L161:
	;
	goto L158
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+352)) = v15 + int32(712)
	F_appendStringInfo(m, l0, int32(206163), v15+int32(352))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L32
	} else {
		goto L163
	}
L163:
	;
	v870 = v838 + int32(1)
	if v870 != v825 {
		v838 = v870
		goto L160
	} else {
		goto L164
	}
L164:
	;
	goto L161
L165:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v15)+684))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+336)) = int32(672320)
	F_appendStringInfo(m, l0, int32(547348), v15+int32(336))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L32
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v15)+460))
	if int32(0) < v943 {
		goto L174
	} else {
		goto L175
	}
L168:
	;
	v897 = int32(0)
	goto L169
L169:
	;
	v912 = v887 + v897*int32(12)
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v912)+4))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v912)))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v912)+8))
	F_GetRelationPath(m, v15+int32(712), v913, v914, v915, int32(-1), int32(0))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L32
	} else {
		goto L171
	}
L170:
	;
	goto L167
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+320)) = v15 + int32(712)
	F_appendStringInfo(m, l0, int32(206163), v15+int32(320))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L32
	} else {
		goto L172
	}
L172:
	;
	v929 = v897 + int32(1)
	if v929 != v884 {
		v897 = v929
		goto L169
	} else {
		goto L173
	}
L173:
	;
	goto L170
L174:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v15)+464))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+304)) = int32(732842)
	F_appendStringInfo(m, l0, int32(547253), v15+int32(304))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L32
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v15)+688))
	if int32(0) < v994 {
		goto L182
	} else {
		goto L183
	}
L177:
	;
	v956 = int32(0)
	goto L178
L178:
	;
	v969 = v946 + v956<<(uint(int32(4))%32)
	v970 = *(*int64)(unsafe.Add(mBase, uint32(v969)))
	v971 = *(*int64)(unsafe.Add(mBase, uint32(v969)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+296)) = v971
	*(*int64)(unsafe.Add(mBase, uint32(v15)+288)) = v970
	F_appendStringInfo(m, l0, int32(37913), v15+int32(288))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L32
	} else {
		goto L180
	}
L179:
	;
	goto L176
L180:
	;
	v980 = v956 + int32(1)
	if v980 != v943 {
		v956 = v980
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v15)+692))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+272)) = int32(732627)
	F_appendStringInfo(m, l0, int32(547253), v15+int32(272))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L32
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	if int32(0) < v1045 {
		goto L190
	} else {
		goto L191
	}
L185:
	;
	v1007 = int32(0)
	goto L186
L186:
	;
	v1020 = v997 + v1007<<(uint(int32(4))%32)
	v1021 = *(*int64)(unsafe.Add(mBase, uint32(v1020)))
	v1022 = *(*int64)(unsafe.Add(mBase, uint32(v1020)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+264)) = v1022
	*(*int64)(unsafe.Add(mBase, uint32(v15)+256)) = v1021
	F_appendStringInfo(m, l0, int32(37913), v15+int32(256))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L32
	} else {
		goto L188
	}
L187:
	;
	goto L184
L188:
	;
	v1031 = v1007 + int32(1)
	if v1031 != v994 {
		v1007 = v1031
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v15)+448))
	F_appendStringInfoString(m, l0, int32(547229))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L32
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v15)+468))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v15)+472))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v15)+436))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v15)+440))
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+52)))
	F_standby_desc_invalidations(m, l0, v1090, v1091, v1092, v1093, v1094)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L32
	} else {
		goto L198
	}
L193:
	;
	v1054 = int32(0)
	goto L194
L194:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1048+v1054<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v1068
	F_appendStringInfo(m, l0, int32(59440), v15+int32(240))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L32
	} else {
		goto L196
	}
L195:
	;
	goto L192
L196:
	;
	v1076 = v1054 + int32(1)
	if v1076 != v1045 {
		v1054 = v1076
		goto L194
	} else {
		goto L197
	}
L197:
	;
	goto L195
L198:
	;
	if v634 == int32(0) {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	v1099 = *(*int64)(unsafe.Add(mBase, uint32(v15)+696))
	v1100 = *(*int64)(unsafe.Add(mBase, uint32(v15)+704))
	v1101 = F_timestamptz_to_str(m, v1100)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L32
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+236)) = v1101
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+232)) = uint32(v1099)
	v1106 = int64(base.Ui64(v1099) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+228)) = uint32(v1106)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v634
	F_appendStringInfo(m, l0, int32(179766), v15+int32(224))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L32
	} else {
		goto L201
	}
L201:
	;
	goto L1
L202:
	;
	F_appendStringInfoString(m, l0, int32(547231))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L32
	} else {
		goto L203
	}
L203:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v1124 <= int32(0) {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v1131 = int32(0)
	goto L205
L205:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(8)+v1131<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+400)) = v1145
	F_appendStringInfo(m, l0, int32(59440), v15+int32(400))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L32
	} else {
		goto L207
	}
L206:
	;
	goto L1
L207:
	;
	v1153 = v1131 + int32(1)
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v1153 < v1154 {
		v1131 = v1153
		goto L205
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	goto L1
}
