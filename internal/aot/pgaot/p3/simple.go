package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SimpleLruTruncate(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v39 int64
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruTruncate[0])) = uint8(v16)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruTruncate[1])) = uint8(v16)
	v22 = v14 << (uint(int32(6)) % 32)
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_SimpleLruTruncate[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_SimpleLruTruncate[2]))) = v25 + int64(1)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = int64(0)
	v33 = base.AtomicRmwCmpxchg64(m, v13, int32(48), v30, v30)
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v35 = m.T0[v29].(func(*base.Module, int64, int64) int32)(m, v33, v34)
	mBase = m.M
	if v35 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return
L2:
	;
	v39 = v34
	goto L5
L3:
	;
	goto L4
L4:
	;
	v156 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L34
	}
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v47 = F_LWLockAcquire(m, v45, int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	return
L8:
	;
	v49 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v49 < v51 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	F_LWLockRelease(m, v132+v77<<(uint(int32(7))%32))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L7
	} else {
		goto L32
	}
L10:
	;
	F_SimpleLruWaitIO(m, l0, v58)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L31
	}
L11:
	;
	F_SlruInternalWritePage(m, l0, v58, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L30
	}
L12:
	;
	v56 = v49
	v58 = v49
	goto L15
L13:
	;
	v111 = v49
	goto L14
L14:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	F_LWLockRelease(m, v116+v111<<(uint(int32(7))%32))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L28
	}
L15:
	;
	v62 = int32(base.Ui32(v58) >> (uint(int32(4)) % 32))
	if v56 != v62 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v111 = v77
	goto L14
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	F_LWLockRelease(m, v64+v56<<(uint(int32(7))%32))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	v77 = v56
	goto L19
L19:
	;
	v79 = v58 << (uint(int32(2)) % 32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79+v80)))
	if v82 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v75 = F_LWLockAcquire(m, v70+v62<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v77 = v62
	goto L19
L22:
	;
	v106 = v58 + int32(1)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v106 < v107 {
		v56 = v77
		v58 = v106
		goto L15
	} else {
		goto L27
	}
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v85+v58<<(uint(int32(3))%32))))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v91 = m.T0[v90].(func(*base.Module, int64, int64) int32)(m, v89, v39)
	mBase = m.M
	if v91 == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v95 = v94 + v79
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v96 != int32(2) {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99+v58))))
	if v101 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(0)
	goto L22
L27:
	;
	goto L16
L28:
	;
	v125 = F_SlruScanDirectory(m, l0, int32(391), v10+int32(8))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	goto L1
L30:
	;
	goto L9
L31:
	;
	goto L9
L32:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v139 = int64(0)
	v142 = base.AtomicRmwCmpxchg64(m, v13, int32(48), v139, v139)
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v144 = m.T0[v138].(func(*base.Module, int64, int64) int32)(m, v142, v143)
	mBase = m.M
	if v144 == int32(0) {
		v39 = v143
		goto L5
	} else {
		goto L33
	}
L33:
	;
	goto L6
L34:
	;
	if v156 == int32(0) {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0 + int32(16)
	F_errmsg(m, int32(_a_F_SimpleLruTruncate_0), v10)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_SimpleLruTruncate_1), int32(1435), int32(_a_F_SimpleLruTruncate_2))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	goto L1
}
func F_SimpleLruWriteAll(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int64
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v139 int32
	_ = v139
	var v143 int64
	_ = v143
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(208)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruWriteAll[0])) = uint8(v16)
	*(*uint8)(unsafe.Add(mBase, _c_F_SimpleLruWriteAll[1])) = uint8(v16)
	v22 = v14 << (uint(int32(6)) % 32)
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_SimpleLruWriteAll[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_SimpleLruWriteAll[2]))) = v25 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v2
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v33 = F_LWLockAcquire(m, v31, v2)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if int32(0) < v35 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v39 = v2
	v42 = v2
	goto L6
L4:
	;
	v80 = v2
	v81 = v2
	goto L5
L5:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	F_LWLockRelease(m, v85+v81<<(uint(int32(7))%32))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L18
	}
L6:
	;
	v47 = int32(base.Ui32(v39) >> (uint(int32(4)) % 32))
	if v42 != v47 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v80 = v76
	v81 = v62
	goto L5
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	F_LWLockRelease(m, v49+v42<<(uint(int32(7))%32))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v62 = v42
	goto L10
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v39<<(uint(int32(2))%32))))
	if v67 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v60 = F_LWLockAcquire(m, v55+v47<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v62 = v47
	goto L10
L13:
	;
	F_SlruInternalWritePage(m, l0, v39, v11+int32(8))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v73 = v39 + int32(1)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v73 < v74 {
		v39 = v73
		v42 = v62
		goto L6
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	goto L7
L18:
	;
	if v80 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v155 != int32(5) {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v102 = int32(0)
	v105 = int32(1)
	v108 = int64(0)
	goto L21
L21:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(8)|int32(4)+v102<<(uint(int32(2))%32))))
	v113 = F_CloseTransientFile(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	F_SlruReportIOError(m, l0, v143, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L31
	}
L23:
	;
	goto L22
L24:
	;
	if v113 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v118 = v102 + int32(1)
	if v118 != v80 {
		v102 = v118
		goto L21
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruWriteAll[3])) = int32(5)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruWriteAll[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_SimpleLruWriteAll[5])) = v130
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(80)+v102<<(uint(int32(3))%32))))
	v137 = v135 << (uint(int64(5)) % 64)
	v139 = v102 + int32(1)
	if v139 != v80 {
		v102 = v139
		v105 = int32(0)
		v108 = v137
		goto L21
	} else {
		goto L30
	}
L28:
	;
	if v105&int32(1) == int32(0) {
		v143 = v108
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L19
L30:
	;
	v143 = v137
	goto L23
L31:
	;
	goto L19
L32:
	;
	F_fsync_fname(m, l0+int32(16), int32(1))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	m.G0 = v11 + int32(208)
	return
L35:
	;
	goto L34
}
func F_check_simple_rowfilter_expr_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L25
	} else {
		goto L48
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v12 - int32(1) {
	case 0:
		goto L5
	default:
		v131 = int32(_a_F_check_simple_rowfilter_expr_walker_0)
		goto L1
	case 5:
		goto L10
	case 6, 14, 20, 26, 30, 31, 33, 34, 35, 37, 38, 40, 51, 52:
		goto L6
	case 16, 17, 18:
		goto L9
	case 19:
		goto L8
	case 36:
		goto L7
	}
L3:
	;
	v124 = int32(0)
	goto L4
L4:
	;
	m.G0 = v9 + int32(16)
	return v124
L5:
	;
	v115 = F_expression_tree_walker_impl(m, l0, int32(569), l1)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L25
	} else {
		goto L47
	}
L6:
	;
	v61 = F_exprType(m, l0)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L25
	} else {
		goto L26
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v27 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v23) <= base.Ui32(int32(_a_F_check_simple_rowfilter_expr_walker_1)) {
		goto L6
	} else {
		goto L13
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v19) <= base.Ui32(int32(_a_F_check_simple_rowfilter_expr_walker_1)) {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	if int32(0) <= v15 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v131 = int32(_a_F_check_simple_rowfilter_expr_walker_2)
	goto L1
L12:
	;
	v131 = int32(_a_F_check_simple_rowfilter_expr_walker_3)
	goto L1
L13:
	;
	v131 = int32(_a_F_check_simple_rowfilter_expr_walker_3)
	goto L1
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v30 <= int32(0) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v33 = int32(0)
	if v33 < v30 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v37 = v30
	goto L18
L17:
	;
	v37 = v33
	goto L18
L18:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v41 = v33
	goto L19
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v38+v41<<(uint(int32(2))%32))))
	if base.Ui32(v48) <= base.Ui32(int32(_a_F_check_simple_rowfilter_expr_walker_1)) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v131 = int32(_a_F_check_simple_rowfilter_expr_walker_3)
	goto L1
L21:
	;
	v52 = v41 + int32(1)
	if v37 != v52 {
		v41 = v52
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L6
L25:
	;
	return int32(0)
L26:
	;
	if base.Ui32(int32(_a_F_check_simple_rowfilter_expr_walker_1)) < base.Ui32(v61) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v131 = int32(_a_F_check_simple_rowfilter_expr_walker_4)
	goto L1
L28:
	;
	goto L29
L29:
	;
	v69 = F_check_functions_in_node(m, l0, int32(568), l1)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	if v69 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v131 = int32(_a_F_check_simple_rowfilter_expr_walker_5)
	goto L1
L32:
	;
	goto L33
L33:
	;
	v72 = int32(_a_F_check_simple_rowfilter_expr_walker_6)
	v73 = F_exprCollation(m, l0)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	if base.Ui32(int32(_a_F_check_simple_rowfilter_expr_walker_1)) < base.Ui32(v73) {
		v131 = v72
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v77 = int32(0)
	if l0 == v77 {
		v105 = v77
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if base.Ui32(int32(_a_F_check_simple_rowfilter_expr_walker_7)) <= base.Ui32(v105) {
		v131 = v72
		goto L1
	} else {
		goto L46
	}
L37:
	;
	goto L36
L38:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v84 = v82 - int32(9)
	if base.Ui32(int32(30)) < base.Ui32(v84) {
		v105 = v77
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v88 = int32(1) << (uint(v84) % 32)
	if v88&int32(3904) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100+l0)))
	v105 = v102
	goto L37
L41:
	;
	if v88&int32(5) != 0 {
		v100 = int32(16)
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v100 = int32(24)
	goto L40
L44:
	;
	if v84 != int32(30) {
		v105 = v77
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v100 = int32(12)
	goto L40
L46:
	;
	goto L5
L47:
	;
	v124 = v115
	goto L4
L48:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L25
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(_a_F_check_simple_rowfilter_expr_walker_8), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v131
	F_errdetail_internal(m, int32(_a_F_check_simple_rowfilter_expr_walker_9), v9)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L25
	} else {
		goto L51
	}
L51:
	;
	v150 = F_exprLocation(m, l0)
	mBase = m.M
	F_parser_errposition(m, l1, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_check_simple_rowfilter_expr_walker_10), int32(680), int32(_a_F_check_simple_rowfilter_expr_walker_11))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L25
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_simple_binary_op_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v4 == int32(0) {
		v29 = int32(0)
		return v29
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		if v7 != int32(2) {
			v29 = int32(0)
			return v29
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v14 = F_exprType(m, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = F_exprType(m, v11)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = F_generate_operator_name(m, v12, v14, v18)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = F_strlen(m, v20)
						mBase = m.M
						if v22 == int32(1) {
							v29 = v20
						} else {
							v29 = int32(0)
						}
						return v29
					}
				}
			}
		}
	}
}
func F_set_simple_column_names(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v2 = int32(0)
	v8 = v2
	goto L1
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v8
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return
L3:
	;
	goto L2
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v16 = v14
	goto L6
L5:
	;
	v16 = int32(0)
	goto L6
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v20 = v18
	goto L9
L8:
	;
	v20 = int32(0)
	goto L9
L9:
	;
	if v20 <= v16 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v24 = v2
	goto L13
L11:
	;
	goto L12
L12:
	;
	v62 = F_palloc0(m, int32(52))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L24
	} else {
		goto L26
	}
L13:
	;
	v28 = int32(0)
	if v17 == v28 {
		v38 = v28
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v8 == int32(0) {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v32 <= v24 {
		v38 = int32(0)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v38 = v34 + v24<<(uint(int32(2))%32)
	goto L15
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if base.B2i32(v38 == int32(0))|base.B2i32(v43 <= v24) != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v46 == int32(0) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	if v50 != int32(2) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46+v24<<(uint(int32(2))%32))))
	F_set_relation_column_names(m, l0, v49, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v24 = v24 + int32(1)
	goto L13
L24:
	;
	return
L25:
	;
	goto L23
L26:
	;
	v64 = F_lappend(m, v8, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v8 = v64
	goto L1
}
func F_setup_simple_rel_arrays(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v12 = v8 + int32(1)
	goto L3
L2:
	;
	v12 = int32(1)
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v12
	v15 = v12 << (uint(int32(2)) % 32)
	v16 = F_palloc0(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v16
	v19 = F_palloc0(m, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v19
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	if v23 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v56 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v27 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v31 = int32(1)
	v34 = int32(0)
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v36 = int32(2)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v34<<(uint(v36)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v35+v31<<(uint(v36)%32)))) = v43
	v45 = int32(1)
	v48 = v34 + v45
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v48 < v49 {
		v31 = v31 + v45
		v34 = v48
		goto L10
	} else {
		goto L12
	}
L11:
	;
	goto L7
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	return
L14:
	;
	goto L15
L15:
	;
	v61 = F_palloc0(m, v15)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v61
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v64 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L25
	}
L18:
	;
	return
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v67 <= int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v72 = int32(0)
	goto L21
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v78 = int32(2)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77+v72<<(uint(v78)%32))))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v85 = v76 + v82<<(uint(v78)%32)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 != 0 {
		goto L17
	} else {
		goto L23
	}
L22:
	;
	goto L18
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v81
	v89 = v72 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v89 < v90 {
		v72 = v89
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	F_errmsg_internal(m, int32(_a_F_setup_simple_rel_arrays_0), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_setup_simple_rel_arrays_1), int32(147), int32(_a_F_setup_simple_rel_arrays_2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_simple_heap_insert(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = F_GetCurrentCommandId(m, int32(1))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = int32(0)
		F_heap_insert(m, l0, l1, v4, v6, v6)
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
