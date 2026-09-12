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
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = l1
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v16 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[143])) = uint8(v16)
	*(*uint8)(unsafe.Add(mBase, _consts[144])) = uint8(v16)
	v22 = v14 << (uint(int32(6)) % 32)
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_consts[151])))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_consts[151]))) = v25 + int64(1)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v30
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v33 = m.T0[v29].(func(*base.Module, int64, int64) int32)(m, v30, v32)
	mBase = m.M
	if v33 == int32(0) {
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
	v37 = v32
	goto L5
L3:
	;
	goto L4
L4:
	;
	v152 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L34
	}
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v45 = F_LWLockAcquire(m, v43, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
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
	v47 = int32(0)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v47 < v49 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	F_LWLockRelease(m, v130+v75<<(uint(int32(7))%32))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
	} else {
		goto L32
	}
L10:
	;
	F_SimpleLruWaitIO(m, l0, v55)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L7
	} else {
		goto L31
	}
L11:
	;
	F_SlruInternalWritePage(m, l0, v55, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L30
	}
L12:
	;
	v55 = v47
	v58 = v47
	goto L15
L13:
	;
	v113 = v47
	goto L14
L14:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	F_LWLockRelease(m, v114+v113<<(uint(int32(7))%32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L28
	}
L15:
	;
	v60 = int32(base.Ui32(v55) >> (uint(int32(4)) % 32))
	if v58 != v60 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v113 = v75
	goto L14
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	F_LWLockRelease(m, v62+v58<<(uint(int32(7))%32))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	v75 = v58
	goto L19
L19:
	;
	v77 = v55 << (uint(int32(2)) % 32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77+v78)))
	if v80 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v73 = F_LWLockAcquire(m, v68+v60<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v75 = v60
	goto L19
L22:
	;
	v104 = v55 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v104 < v105 {
		v55 = v104
		v58 = v75
		goto L15
	} else {
		goto L27
	}
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v83+v55<<(uint(int32(3))%32))))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v89 = m.T0[v88].(func(*base.Module, int64, int64) int32)(m, v87, v37)
	mBase = m.M
	if v89 == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v93 = v92 + v77
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 != int32(2) {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v55))))
	if v99 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = int32(0)
	goto L22
L27:
	;
	goto L16
L28:
	;
	v123 = F_SlruScanDirectory(m, l0, int32(391), v10+int32(8))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
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
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v13)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v137
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
	v140 = m.T0[v136].(func(*base.Module, int64, int64) int32)(m, v137, v139)
	mBase = m.M
	if v140 == int32(0) {
		v37 = v139
		goto L5
	} else {
		goto L33
	}
L33:
	;
	goto L6
L34:
	;
	if v152 == int32(0) {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0 + int32(16)
	F_errmsg(m, int32(420664), v10)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(489398), int32(1435), int32(355259))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int64
	_ = v121
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(208)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v15 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[143])) = uint8(v15)
	*(*uint8)(unsafe.Add(mBase, _consts[144])) = uint8(v15)
	v21 = v13 << (uint(int32(6)) % 32)
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[150])))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[150]))) = v24 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v2
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v32 = F_LWLockAcquire(m, v30, v2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if int32(0) < v34 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v39 = v2
	v41 = v2
	goto L6
L4:
	;
	v78 = v2
	goto L5
L5:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_LWLockRelease(m, v81+v78<<(uint(int32(7))%32))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L18
	}
L6:
	;
	v45 = int32(base.Ui32(v39) >> (uint(int32(4)) % 32))
	if v41 != v45 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v78 = v60
	goto L5
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_LWLockRelease(m, v47+v41<<(uint(int32(7))%32))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v60 = v41
	goto L10
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v39<<(uint(int32(2))%32))))
	if v65 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v58 = F_LWLockAcquire(m, v53+v45<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v60 = v45
	goto L10
L13:
	;
	F_SlruInternalWritePage(m, l0, v39, v10+int32(8))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v71 = v39 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v71 < v72 {
		v39 = v71
		v41 = v60
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
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v87 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v143 != int32(5) {
		goto L30
	} else {
		goto L31
	}
L20:
	;
	v100 = int32(0)
	v102 = int32(1)
	v104 = int64(0)
	goto L21
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(8)|int32(4)+v100<<(uint(int32(2))%32))))
	v109 = F_CloseTransientFile(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	if v125&int32(1) != 0 {
		goto L19
	} else {
		goto L28
	}
L23:
	;
	if v109 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[146])) = int32(5)
	v116 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	*(*int32)(unsafe.Add(mBase, _consts[147])) = v116
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v10+int32(80)+v100<<(uint(int32(3))%32))))
	v125 = int32(0)
	v126 = v121 << (uint(int64(5)) % 64)
	goto L26
L25:
	;
	v125 = v102
	v126 = v104
	goto L26
L26:
	;
	v128 = v100 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v128 < v129 {
		v100 = v128
		v102 = v125
		v104 = v126
		goto L21
	} else {
		goto L27
	}
L27:
	;
	goto L22
L28:
	;
	F_SlruReportIOError(m, l0, v126, int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L19
L30:
	;
	F_fsync_fname(m, l0+int32(16), int32(1))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	m.G0 = v10 + int32(208)
	return
L33:
	;
	goto L32
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
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
	v137 = m.ExcPending
	if v137 != 0 {
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
		v130 = int32(628513)
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
	v123 = int32(0)
	goto L4
L4:
	;
	m.G0 = v9 + int32(16)
	return v123
L5:
	;
	v114 = F_expression_tree_walker_impl(m, l0, int32(569), l1)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
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
	if base.Ui32(v23) <= base.Ui32(int32(16383)) {
		goto L6
	} else {
		goto L13
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v19) <= base.Ui32(int32(16383)) {
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
	v130 = int32(628415)
	goto L1
L12:
	;
	v130 = int32(628274)
	goto L1
L13:
	;
	v130 = int32(628274)
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
	if base.Ui32(v48) <= base.Ui32(int32(16383)) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v130 = int32(628274)
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
	if base.Ui32(int32(16383)) < base.Ui32(v61) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v130 = int32(628477)
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
	v130 = int32(628314)
	goto L1
L32:
	;
	goto L33
L33:
	;
	v72 = int32(628374)
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
	if base.Ui32(int32(16383)) < base.Ui32(v73) {
		v130 = v72
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v77 = int32(0)
	if l0 == v77 {
		v103 = v77
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if base.Ui32(int32(16384)) <= base.Ui32(v103) {
		v130 = v72
		goto L1
	} else {
		goto L46
	}
L37:
	;
	goto L36
L38:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v83 = v81 - int32(9)
	if base.Ui32(int32(30)) < base.Ui32(v83) {
		v103 = v77
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v87 = int32(1) << (uint(v83) % 32)
	if v87&int32(3904) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0+v99)))
	v103 = v101
	goto L37
L41:
	;
	if v87&int32(5) != 0 {
		v99 = int32(16)
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v99 = int32(24)
	goto L40
L44:
	;
	if v83 != int32(30) {
		v103 = v77
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v99 = int32(12)
	goto L40
L46:
	;
	goto L5
L47:
	;
	v123 = v114
	goto L4
L48:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L25
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(268852), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v130
	F_errdetail_internal(m, int32(205224), v9)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L25
	} else {
		goto L51
	}
L51:
	;
	v149 = F_exprLocation(m, l0)
	mBase = m.M
	F_parser_errposition(m, l1, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(491199), int32(680), int32(220543))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	v15 = v14
	goto L6
L5:
	;
	v15 = v2
	goto L6
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v19 = v17
	goto L9
L8:
	;
	v19 = int32(0)
	goto L9
L9:
	;
	if v19 <= v15 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v23 = v2
	goto L13
L11:
	;
	goto L12
L12:
	;
	v60 = F_palloc0(m, int32(52))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L25
	} else {
		goto L27
	}
L13:
	;
	v27 = int32(0)
	if v16 == v27 {
		v37 = v27
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
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v31 <= v23 {
		v37 = int32(0)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v37 = v33 + v23<<(uint(int32(2))%32)
	goto L15
L18:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v40 <= v23 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	if v37 == int32(0) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v47 = v44 + v23<<(uint(int32(2))%32)
	if v47 == int32(0) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v51 != int32(2) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	F_set_relation_column_names(m, l0, v50, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v23 = v23 + int32(1)
	goto L13
L25:
	;
	return
L26:
	;
	goto L24
L27:
	;
	v62 = F_lappend(m, v8, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v8 = v62
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v57 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v26 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v27 <= v26 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v32 = v26
	v33 = int32(1)
	goto L10
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v37 = int32(2)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v32<<(uint(v37)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v36+v33<<(uint(v37)%32)))) = v44
	v46 = int32(1)
	v49 = v32 + v46
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v49 < v50 {
		v32 = v49
		v33 = v33 + v46
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
	v62 = F_palloc0(m, v15)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v62
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v65 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L25
	}
L18:
	;
	return
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v68 <= int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v73 = int32(0)
	goto L21
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v79 = int32(2)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78+v73<<(uint(v79)%32))))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v86 = v77 + v83<<(uint(v79)%32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v87 != 0 {
		goto L17
	} else {
		goto L23
	}
L22:
	;
	goto L18
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v82
	v90 = v73 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v90 < v91 {
		v73 = v90
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	F_errmsg_internal(m, int32(115451), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(496220), int32(147), int32(112585))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
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
