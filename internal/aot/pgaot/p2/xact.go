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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(1112)
	v15 = int32(_a_F_XactLockTableWait_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_XactLockTableWait[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_XactLockTableWait[0])) = v8 + int32(8)
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
	v32 = v8 + int32(32)
	v34 = int32(0)
	v36 = F_LockAcquire(m, v32, int32(5), v34, v34)
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
	v40 = F_LockRelease(m, v32, int32(5), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v42 = F_TransactionIdIsInProgress(m, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	if l3 != 0 {
		goto L21
	} else {
		goto L22
	}
L8:
	;
	if v42 == int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v46 = l0
	goto L10
L10:
	;
	v51 = F_SubTransGetTopmostTransaction(m, v46)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(17104896)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v51
	v59 = v8 + int32(32)
	v61 = int32(0)
	v63 = F_LockAcquire(m, v59, int32(5), v61, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v67 = F_LockRelease(m, v59, int32(5), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v69 = F_TransactionIdIsInProgress(m, v51)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if v69 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_XactLockTableWait[1]))
	if v74 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
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
	v46 = v51
	goto L10
L20:
	;
	goto L19
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_XactLockTableWait[0])) = v85
	goto L23
L22:
	;
	goto L23
L23:
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
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XactLogCommitRecord[0])))
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
	if v28|v34 != 0 {
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
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_XactLogCommitRecord[1]))
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
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_XactLogCommitRecord[2]))
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
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_XactLogCommitRecord[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v60
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_XactLogCommitRecord[4]))
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
	v113 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_XactLogCommitRecord[5])))
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
	v118 = *(*int64)(unsafe.Add(mBase, _c_F_XactLogCommitRecord[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v118
	v121 = *(*int64)(unsafe.Add(mBase, _c_F_XactLogCommitRecord[7]))
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
	v235 = int32(_a_F_XactLogCommitRecord_0)
	v237 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_XactLogCommitRecord[8])))
	v238 = v237 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_XactLogCommitRecord[8])) = uint8(v238)
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int64
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v283 int32
	_ = v283
	var v284 int64
	_ = v284
	var v285 int64
	_ = v285
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int64
	_ = v338
	var v339 int64
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int64
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int64
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int64
	_ = v445
	var v446 int64
	_ = v446
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v461 int64
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v546 int32
	_ = v546
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v568 int32
	_ = v568
	var v571 int64
	_ = v571
	var v572 int64
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int64
	_ = v578
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v613 int32
	_ = v613
	var v614 int64
	_ = v614
	var v615 int64
	_ = v615
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v632 int64
	_ = v632
	var v634 int64
	_ = v634
	var v636 int64
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v811 int32
	_ = v811
	var v812 int64
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v957 int32
	_ = v957
	var v958 int64
	_ = v958
	var v959 int64
	_ = v959
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v1008 int32
	_ = v1008
	var v1009 int64
	_ = v1009
	var v1010 int64
	_ = v1010
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1056 int32
	_ = v1056
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1087 int64
	_ = v1087
	var v1088 int64
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1094 int64
	_ = v1094
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1119 int32
	_ = v1119
	var v1133 int32
	_ = v1133
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
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
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v1147 = int32(0)
	F_standby_desc_invalidations(m, l0, v1144, v18+int32(4), v1147, v1147, v1147)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L32
	} else {
		goto L207
	}
L3:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+416)) = v1102
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_0), v15+int32(416))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L32
	} else {
		goto L200
	}
L4:
	;
	v626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+56)))
	base.MemoryFill(m, v15+int32(432), int32(0), int32(248))
	v632 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+424)) = v632
	v634 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+696)) = v634
	v636 = *(*int64)(unsafe.Add(mBase, uint32(v18)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+704)) = v636
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+476)) = v638
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+436)) = v640
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+444)) = v642
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+452)) = v644
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+680)) = v646
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+460)) = v648
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+688)) = v650
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+468)) = v652
	v655 = v15 + int32(480)
	v657 = v18 + int32(72)
	v658 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+54)))
	if (v657^v655)&int32(3) != 0 {
		v729 = v657
		v730 = v658
		v731 = v655
		goto L130
	} else {
		goto L131
	}
L5:
	;
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+56)))
	v353 = v15 + int32(424)
	v354 = int32(0)
	base.MemoryFill(m, v353, v354, int32(264))
	v360 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v353))) = v360
	if v354 <= base.I32_extend8_s(v19) {
		goto L74
	} else {
		goto L75
	}
L6:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+56)))
	v28 = v15 + int32(424)
	v29 = int32(0)
	base.MemoryFill(m, v28, v29, int32(288))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v35
	if v29 <= base.I32_extend8_s(v19) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v15)+476))
	if v143 != 0 {
		goto L29
	} else {
		goto L30
	}
L8:
	;
	goto L7
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v40
	if v40&int32(1) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v46
	v52 = v18 + int32(20)
	goto L12
L11:
	;
	v52 = v18 + int32(12)
	goto L12
L12:
	;
	if v40&int32(2) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v57 = v52 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v55
	v63 = v57 + v55<<(uint(int32(2))%32)
	goto L15
L14:
	;
	v63 = v52
	goto L15
L15:
	;
	if v40&int32(4) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v69 = v63 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v67
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v76 = v69 + v72*int32(12)
	goto L18
L17:
	;
	v76 = v63
	goto L18
L18:
	;
	if v40&int32(256) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v82 = int32(4)
	v83 = v76 + v82
	*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v81
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v90 = v83 + v86<<(uint(v82)%32)
	goto L21
L20:
	;
	v90 = v76
	goto L21
L21:
	;
	if v40&int32(8) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v96 = int32(4)
	v97 = v90 + v96
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v28)+44)) = v95
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v104 = v97 + v100<<(uint(v96)%32)
	goto L24
L23:
	;
	v104 = v90
	goto L24
L24:
	;
	if v40&int32(16) == int32(0) {
		v128 = v40
		v129 = v104
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v128&int32(32) == int32(0) {
		goto L8
	} else {
		goto L28
	}
L26:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+52)) = v111
	v114 = v104 + int32(4)
	if v40&int32(128) == int32(0) {
		v128 = v40
		v129 = v114
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v122 = F_strlcpy(m, v15+int32(480), v114, int32(200))
	mBase = m.M
	v123 = F_strlen(m, v114)
	mBase = m.M
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v128 = v127
	v129 = v123 + v114 + int32(1)
	goto L25
L28:
	;
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v129)))
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v129)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+280)) = v135
	*(*int64)(unsafe.Add(mBase, uint32(v28)+272)) = v134
	goto L8
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v143
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_1), v15+int32(96))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v151 = F_timestamptz_to_str(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
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
	F_appendStringInfoString(m, l0, v151)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v15)+452))
	if int32(0) < v155 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v15)+456))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = int32(_a_F_xact_desc_2)
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_3), v15+int32(80))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L32
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	if int32(0) < v212 {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	v168 = int32(0)
	goto L40
L40:
	;
	v180 = v15 + int32(712)
	v183 = v158 + v168*int32(12)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	F_GetRelationPath(m, v180, v184, v185, v186, int32(-1), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L32
	} else {
		goto L42
	}
L41:
	;
	goto L38
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v180
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_4), v15-int32(-64))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L32
	} else {
		goto L43
	}
L43:
	;
	v198 = v168 + int32(1)
	if v198 != v155 {
		v168 = v198
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v15)+448))
	F_appendStringInfoString(m, l0, int32(_a_F_xact_desc_5))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L32
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v15)+460))
	if int32(0) < v257 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v221 = int32(0)
	goto L49
L49:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v215+v221<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v235
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_6), v15+int32(48))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L32
	} else {
		goto L51
	}
L50:
	;
	goto L47
L51:
	;
	v243 = v221 + int32(1)
	if v243 != v212 {
		v221 = v243
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v15)+464))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(_a_F_xact_desc_7)
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_8), v15+int32(32))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L32
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v15)+468))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v15)+472))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v15)+436))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v15)+440))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v15)+432))
	F_standby_desc_invalidations(m, l0, v308, v309, v310, v311, int32(base.Ui32(v312&int32(1073741824))>>(uint(int32(30))%32)))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L32
	} else {
		goto L61
	}
L56:
	;
	v270 = int32(0)
	goto L57
L57:
	;
	v283 = v260 + v270<<(uint(int32(4))%32)
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v283)))
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v283)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v285
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v284
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_9), v15+int32(16))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L32
	} else {
		goto L59
	}
L58:
	;
	goto L55
L59:
	;
	v294 = v270 + int32(1)
	if v294 != v257 {
		v270 = v294
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v15)+432))
	if v319&int32(536870912) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_xact_desc_10))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L32
	} else {
		goto L65
	}
L63:
	;
	v326 = v319
	goto L64
L64:
	;
	if v326 < int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v15)+432))
	v326 = v325
	goto L64
L66:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_xact_desc_11))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L32
	} else {
		goto L69
	}
L67:
	;
	v333 = v326
	goto L68
L68:
	;
	if v333&int32(32) == int32(0) {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v15)+432))
	v333 = v332
	goto L68
L70:
	;
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v15)+696))
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v15)+704))
	v340 = F_timestamptz_to_str(m, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L32
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v340
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+8)) = uint32(v338)
	v345 = int64(base.Ui64(v338) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+4)) = uint32(v345)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v26
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_12), v15)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L32
	} else {
		goto L72
	}
L72:
	;
	goto L1
L73:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v15)+468))
	if v454 != 0 {
		goto L92
	} else {
		goto L93
	}
L74:
	;
	goto L73
L75:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v353)+8)) = v365
	if v365&int32(1) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v353)+12)) = v369
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v353)+16)) = v371
	v377 = v18 + int32(20)
	goto L78
L77:
	;
	v377 = v18 + int32(12)
	goto L78
L78:
	;
	if v365&int32(2) != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	v382 = v377 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v353)+24)) = v382
	*(*int32)(unsafe.Add(mBase, uint32(v353)+20)) = v380
	v388 = v382 + v380<<(uint(int32(2))%32)
	goto L81
L80:
	;
	v388 = v377
	goto L81
L81:
	;
	if v365&int32(4) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v394 = v388 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v353)+32)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v353)+28)) = v392
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v401 = v394 + v397*int32(12)
	goto L84
L83:
	;
	v401 = v388
	goto L84
L84:
	;
	if v365&int32(256) != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	v407 = int32(4)
	v408 = v401 + v407
	*(*int32)(unsafe.Add(mBase, uint32(v353)+40)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v353)+36)) = v406
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	v415 = v408 + v411<<(uint(v407)%32)
	goto L87
L86:
	;
	v415 = v401
	goto L87
L87:
	;
	if v365&int32(16) == int32(0) {
		v439 = v365
		v440 = v415
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v439&int32(32) == int32(0) {
		goto L74
	} else {
		goto L91
	}
L89:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	*(*int32)(unsafe.Add(mBase, uint32(v353)+44)) = v422
	v425 = v415 + int32(4)
	if v365&int32(128) == int32(0) {
		v439 = v365
		v440 = v425
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v433 = F_strlcpy(m, v15+int32(472), v425, int32(200))
	mBase = m.M
	v434 = F_strlen(m, v425)
	mBase = m.M
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v353)+8))
	v439 = v438
	v440 = v434 + v425 + int32(1)
	goto L88
L91:
	;
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v440)))
	v446 = *(*int64)(unsafe.Add(mBase, uint32(v440)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v353)+256)) = v446
	*(*int64)(unsafe.Add(mBase, uint32(v353)+248)) = v445
	goto L74
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v454
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_1), v15+int32(208))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L32
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v461 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v462 = F_timestamptz_to_str(m, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L32
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	F_appendStringInfoString(m, l0, v462)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L32
	} else {
		goto L97
	}
L97:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v15)+452))
	if int32(0) < v466 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v15)+456))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = int32(_a_F_xact_desc_2)
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_3), v15+int32(192))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L32
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	if int32(0) < v523 {
		goto L107
	} else {
		goto L108
	}
L101:
	;
	v479 = int32(0)
	goto L102
L102:
	;
	v491 = v15 + int32(712)
	v494 = v469 + v479*int32(12)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v494)+8))
	F_GetRelationPath(m, v491, v495, v496, v497, int32(-1), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L32
	} else {
		goto L104
	}
L103:
	;
	goto L100
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = v491
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_4), v15+int32(176))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L32
	} else {
		goto L105
	}
L105:
	;
	v509 = v479 + int32(1)
	if v509 != v466 {
		v479 = v509
		goto L102
	} else {
		goto L106
	}
L106:
	;
	goto L103
L107:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v15)+448))
	F_appendStringInfoString(m, l0, int32(_a_F_xact_desc_5))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L32
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+432)))
	if v568&int32(32) != 0 {
		goto L115
	} else {
		goto L116
	}
L110:
	;
	v532 = int32(0)
	goto L111
L111:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v526+v532<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v546
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_6), v15+int32(160))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L32
	} else {
		goto L113
	}
L112:
	;
	goto L109
L113:
	;
	v554 = v532 + int32(1)
	if v554 != v523 {
		v532 = v554
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v571 = *(*int64)(unsafe.Add(mBase, uint32(v15)+672))
	v572 = *(*int64)(unsafe.Add(mBase, uint32(v15)+680))
	v573 = F_timestamptz_to_str(m, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L32
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v15)+460))
	if v587 <= int32(0) {
		goto L1
	} else {
		goto L120
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+156)) = v573
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+152)) = uint32(v571)
	v578 = int64(base.Ui64(v571) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+148)) = uint32(v578)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v351
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_12), v15+int32(144))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L32
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v15)+464))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = int32(_a_F_xact_desc_7)
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_8), v15+int32(128))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L32
	} else {
		goto L121
	}
L121:
	;
	v600 = int32(0)
	goto L122
L122:
	;
	v613 = v590 + v600<<(uint(int32(4))%32)
	v614 = *(*int64)(unsafe.Add(mBase, uint32(v613)))
	v615 = *(*int64)(unsafe.Add(mBase, uint32(v613)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+120)) = v615
	*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v614
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_9), v15+int32(112))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L32
	} else {
		goto L124
	}
L123:
	;
	goto L1
L124:
	;
	v624 = v600 + int32(1)
	if v624 != v587 {
		v600 = v624
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+384)) = v655
	v769 = int32(7)
	v773 = v657 + (v658+v769)&int32(_a_F_xact_desc_13)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+448)) = v773
	v779 = int32(-8)
	v781 = v773 + (v642<<(uint(int32(2))%32)+v769)&v779
	*(*int32)(unsafe.Add(mBase, uint32(v15)+456)) = v781
	v783 = int32(12)
	v789 = v781 + (v644*v783+v769)&v779
	*(*int32)(unsafe.Add(mBase, uint32(v15)+684)) = v789
	v797 = v789 + (v646*v783+v769)&v779
	*(*int32)(unsafe.Add(mBase, uint32(v15)+464)) = v797
	v799 = int32(4)
	v801 = v797 + v648<<(uint(v799)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+692)) = v801
	*(*int32)(unsafe.Add(mBase, uint32(v15)+472)) = v801 + v650<<(uint(v799)%32)
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_14), v15+int32(384))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L32
	} else {
		goto L151
	}
L127:
	;
	F___memset(m, v764, int32(0), v763)
	mBase = m.M
	goto L126
L128:
	;
	v763 = int32(0)
	v764 = v758
	goto L127
L129:
	;
	v741 = v736
	v742 = v737
	v743 = v738
	goto L147
L130:
	;
	if v730 == int32(0) {
		v758 = v731
		goto L128
	} else {
		goto L146
	}
L131:
	;
	v664 = int32(0)
	if base.B2i32(v657&int32(3) == v664)|base.B2i32(v658 == v664) != 0 {
		v695 = v657
		v696 = v658
		v697 = v655
		v698 = base.B2i32(v658 != v664)
		goto L132
	} else {
		goto L133
	}
L132:
	;
	if v698 == int32(0) {
		v758 = v697
		goto L128
	} else {
		goto L139
	}
L133:
	;
	v674 = v657
	v675 = v658
	v676 = v655
	goto L134
L134:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	*(*uint8)(unsafe.Add(mBase, uint32(v676))) = uint8(v678)
	if v678 == int32(0) {
		v763 = v675
		v764 = v676
		goto L127
	} else {
		goto L136
	}
L135:
	;
	v695 = v689
	v696 = v685
	v697 = v683
	v698 = v687
	goto L132
L136:
	;
	v682 = int32(1)
	v683 = v676 + v682
	v685 = v675 - v682
	v686 = int32(0)
	v687 = base.B2i32(v685 != v686)
	v689 = v674 + v682
	if v689&int32(3) == v686 {
		v695 = v689
		v696 = v685
		v697 = v683
		v698 = v687
		goto L132
	} else {
		goto L137
	}
L137:
	;
	if v685 != 0 {
		v674 = v689
		v675 = v685
		v676 = v683
		goto L134
	} else {
		goto L138
	}
L138:
	;
	goto L135
L139:
	;
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v695))))
	if v701 == int32(0) {
		v763 = v696
		v764 = v697
		goto L127
	} else {
		goto L140
	}
L140:
	;
	if base.Ui32(v696) < base.Ui32(int32(4)) {
		v729 = v695
		v730 = v696
		v731 = v697
		goto L130
	} else {
		goto L141
	}
L141:
	;
	v707 = v695
	v708 = v696
	v709 = v697
	goto L142
L142:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	v715 = int32(-2139062144)
	if (int32(16843008)-v712|v712)&v715 != v715 {
		v736 = v707
		v737 = v708
		v738 = v709
		goto L129
	} else {
		goto L144
	}
L143:
	;
	v729 = v723
	v730 = v725
	v731 = v721
	goto L130
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v709))) = v712
	v720 = int32(4)
	v721 = v709 + v720
	v723 = v707 + v720
	v725 = v708 - v720
	if base.Ui32(int32(3)) < base.Ui32(v725) {
		v707 = v723
		v708 = v725
		v709 = v721
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v736 = v729
	v737 = v730
	v738 = v731
	goto L129
L147:
	;
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v741))))
	*(*uint8)(unsafe.Add(mBase, uint32(v743))) = uint8(v745)
	if v745 == int32(0) {
		v763 = v742
		v764 = v743
		goto L127
	} else {
		goto L149
	}
L148:
	;
	v758 = v750
	goto L128
L149:
	;
	v749 = int32(1)
	v750 = v743 + v749
	v754 = v742 - v749
	if v754 != 0 {
		v741 = v741 + v749
		v742 = v754
		v743 = v750
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v812 = *(*int64)(unsafe.Add(mBase, uint32(v15)+424))
	v813 = F_timestamptz_to_str(m, v812)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L32
	} else {
		goto L152
	}
L152:
	;
	F_appendStringInfoString(m, l0, v813)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L32
	} else {
		goto L153
	}
L153:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v15)+452))
	if int32(0) < v817 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v15)+456))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+368)) = int32(_a_F_xact_desc_15)
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_3), v15+int32(368))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L32
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v15)+680))
	if int32(0) < v874 {
		goto L163
	} else {
		goto L164
	}
L157:
	;
	v830 = int32(0)
	goto L158
L158:
	;
	v842 = v15 + int32(712)
	v845 = v820 + v830*int32(12)
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v845)+4))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v845)))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v845)+8))
	F_GetRelationPath(m, v842, v846, v847, v848, int32(-1), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L32
	} else {
		goto L160
	}
L159:
	;
	goto L156
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+352)) = v842
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_4), v15+int32(352))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L32
	} else {
		goto L161
	}
L161:
	;
	v860 = v830 + int32(1)
	if v860 != v817 {
		v830 = v860
		goto L158
	} else {
		goto L162
	}
L162:
	;
	goto L159
L163:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v15)+684))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+336)) = int32(_a_F_xact_desc_16)
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_3), v15+int32(336))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L32
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v15)+460))
	if int32(0) < v931 {
		goto L172
	} else {
		goto L173
	}
L166:
	;
	v887 = int32(0)
	goto L167
L167:
	;
	v899 = v15 + int32(712)
	v902 = v877 + v887*int32(12)
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v902)+4))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v902)))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v902)+8))
	F_GetRelationPath(m, v899, v903, v904, v905, int32(-1), int32(0))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L32
	} else {
		goto L169
	}
L168:
	;
	goto L165
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+320)) = v899
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_4), v15+int32(320))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L32
	} else {
		goto L170
	}
L170:
	;
	v917 = v887 + int32(1)
	if v917 != v874 {
		v887 = v917
		goto L167
	} else {
		goto L171
	}
L171:
	;
	goto L168
L172:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v15)+464))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+304)) = int32(_a_F_xact_desc_17)
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_8), v15+int32(304))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L32
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v15)+688))
	if int32(0) < v982 {
		goto L180
	} else {
		goto L181
	}
L175:
	;
	v944 = int32(0)
	goto L176
L176:
	;
	v957 = v934 + v944<<(uint(int32(4))%32)
	v958 = *(*int64)(unsafe.Add(mBase, uint32(v957)))
	v959 = *(*int64)(unsafe.Add(mBase, uint32(v957)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+296)) = v959
	*(*int64)(unsafe.Add(mBase, uint32(v15)+288)) = v958
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_9), v15+int32(288))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L32
	} else {
		goto L178
	}
L177:
	;
	goto L174
L178:
	;
	v968 = v944 + int32(1)
	if v968 != v931 {
		v944 = v968
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v15)+692))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+272)) = int32(_a_F_xact_desc_18)
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_8), v15+int32(272))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L32
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	if int32(0) < v1033 {
		goto L188
	} else {
		goto L189
	}
L183:
	;
	v995 = int32(0)
	goto L184
L184:
	;
	v1008 = v985 + v995<<(uint(int32(4))%32)
	v1009 = *(*int64)(unsafe.Add(mBase, uint32(v1008)))
	v1010 = *(*int64)(unsafe.Add(mBase, uint32(v1008)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+264)) = v1010
	*(*int64)(unsafe.Add(mBase, uint32(v15)+256)) = v1009
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_9), v15+int32(256))
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L32
	} else {
		goto L186
	}
L185:
	;
	goto L182
L186:
	;
	v1019 = v995 + int32(1)
	if v1019 != v982 {
		v995 = v1019
		goto L184
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v15)+448))
	F_appendStringInfoString(m, l0, int32(_a_F_xact_desc_5))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L32
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v15)+468))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v15)+472))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v15)+436))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v15)+440))
	v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+52)))
	F_standby_desc_invalidations(m, l0, v1078, v1079, v1080, v1081, v1082)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L32
	} else {
		goto L196
	}
L191:
	;
	v1042 = int32(0)
	goto L192
L192:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1036+v1042<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v1056
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_6), v15+int32(240))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L32
	} else {
		goto L194
	}
L193:
	;
	goto L190
L194:
	;
	v1064 = v1042 + int32(1)
	if v1064 != v1033 {
		v1042 = v1064
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	if v626 == int32(0) {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	v1087 = *(*int64)(unsafe.Add(mBase, uint32(v15)+696))
	v1088 = *(*int64)(unsafe.Add(mBase, uint32(v15)+704))
	v1089 = F_timestamptz_to_str(m, v1088)
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L32
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+236)) = v1089
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+232)) = uint32(v1087)
	v1094 = int64(base.Ui64(v1087) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+228)) = uint32(v1094)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v626
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_12), v15+int32(224))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L32
	} else {
		goto L199
	}
L199:
	;
	goto L1
L200:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_xact_desc_19))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L32
	} else {
		goto L201
	}
L201:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v1112 <= int32(0) {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v1119 = int32(0)
	goto L203
L203:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(8)+v1119<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+400)) = v1133
	F_appendStringInfo(m, l0, int32(_a_F_xact_desc_6), v15+int32(400))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L32
	} else {
		goto L205
	}
L204:
	;
	goto L1
L205:
	;
	v1141 = v1119 + int32(1)
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v1141 < v1142 {
		v1119 = v1141
		goto L203
	} else {
		goto L206
	}
L206:
	;
	goto L204
L207:
	;
	goto L1
}
