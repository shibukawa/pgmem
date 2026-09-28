package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WinGetFuncArgInPartition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v74 int32
	_ = v74
	var v79 int64
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v167 int64
	_ = v167
	var v168 int32
	_ = v168
	var v180 int64
	_ = v180
	var v183 int32
	_ = v183
	var v196 int64
	_ = v196
	v6 = int64(0)
	v10 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)+176))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if base.B2i32(v22 == int32(3))&base.B2i32(l1 != v10) == v10 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v196
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v183)
	v196 = v180
	goto L1
L3:
	;
	v31 = v21 + base.I64_extend_i32_s(l1)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)+408))
	v33 = F_window_gettupleslot(m, l0, v31, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v53 = l1 >> (uint(int32(31)) % 32)
	v55 = l1 ^ v53 - v53
	v58 = base.B2i32(int32(0) < l1)
	if int32(0) < l1 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	return int64(0)
L7:
	;
	if v33 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v39 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v39)
	if l4 != 0 {
		v180 = v6
		v183 = v39
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v32
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	v48 = m.T0[v47].(func(*base.Module, int32, int32, int32) int64)(m, v46, v42, l3)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v196 = v6
	goto L1
L12:
	;
	if l2 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_WinSetMarkPosition(m, l0, v31)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if l4 != 0 {
		v180 = v48
		v183 = v10
		goto L2
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v196 = v48
	goto L1
L18:
	;
	v59 = v21
	goto L20
L19:
	;
	v59 = int64(0)
	goto L20
L20:
	;
	if int32(0) < l1 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v62 = int64(1)
	goto L23
L22:
	;
	v62 = int64(-1)
	goto L23
L23:
	;
	v65 = int32(0)
	v69 = v21
	v70 = v6
	v74 = v10
	goto L27
L24:
	;
	if l4 == int32(0) {
		v196 = v167
		goto L1
	} else {
		goto L57
	}
L25:
	;
	if base.B2i32(l2 == int32(0))|v156 != 0 {
		v167 = v155
		v168 = v156
		goto L24
	} else {
		goto L55
	}
L26:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+408))
	v137 = F_window_gettupleslot(m, l0, v79, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L50
	}
L27:
	;
	v79 = v69 + v62
	if int64(0) <= v79 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	if v123&int32(1) != 0 {
		v155 = v122
		v156 = v10
		goto L25
	} else {
		goto L49
	}
L29:
	;
	if v121 < v55 {
		v65 = v121
		v69 = v79
		v70 = v122
		v74 = v123
		goto L27
	} else {
		goto L48
	}
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v86)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = v87
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
	v110 = m.T0[v109].(func(*base.Module, int32, int32, int32) int64)(m, v108, v104, l3)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L40
	}
L31:
	;
	v121 = v65 + int32(1)
	v122 = v70
	v123 = v74
	goto L29
L32:
	;
	v82 = F_get_notnull_info(m, l0, v79)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L36
	}
L33:
	;
	v94 = v70
	v95 = v10
	goto L34
L34:
	;
	if v74&int32(1) == int32(0) {
		goto L26
	} else {
		goto L39
	}
L35:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+408))
	v88 = F_window_gettupleslot(m, l0, v79, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L37
	}
L36:
	;
	switch v82 - int32(1) {
	case 0:
		v121 = v65
		v122 = v70
		v123 = v74
		goto L29
	case 1:
		goto L31
	default:
		goto L35
	}
L37:
	;
	if v88 != 0 {
		goto L30
	} else {
		goto L38
	}
L38:
	;
	v90 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v90)
	v94 = int64(0)
	v95 = v90
	goto L34
L39:
	;
	v155 = v94
	v156 = v95
	goto L25
L40:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	F_put_notnull_info(m, l0, v79, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v116 = v65 + int32(1)
	if v112 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v119 = v74
	goto L44
L43:
	;
	v119 = base.B2i32(v55 <= v116) | v74
	goto L44
L44:
	;
	if v112 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v120 = v65
	goto L47
L46:
	;
	v120 = v116
	goto L47
L47:
	;
	v121 = v120
	v122 = v110
	v123 = v119
	goto L29
L48:
	;
	goto L28
L49:
	;
	goto L26
L50:
	;
	if v137 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v141 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v141)
	v167 = int64(0)
	v168 = v141
	goto L24
L52:
	;
	goto L53
L53:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v135)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v136
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+24))
	v152 = m.T0[v151].(func(*base.Module, int32, int32, int32) int64)(m, v150, v145, l3)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	v155 = v152
	v156 = int32(0)
	goto L25
L55:
	;
	F_WinSetMarkPosition(m, l0, v59)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v167 = v155
	v168 = int32(0)
	goto L24
L57:
	;
	v180 = v167
	v183 = v168
	goto L2
}
func F_WinGetPartitionRowCount(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_spool_tuples(m, v2, int64(-1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)+168))
		return v9
	}
}
