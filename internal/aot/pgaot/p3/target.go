package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_transformTargetList(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 == v4 {
		v151 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v155 != 0 {
		goto L46
	} else {
		goto L47
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 <= int32(0) {
		v151 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = v4
	v30 = v4
	goto L4
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v30<<(uint(int32(2))%32))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	if base.B2i32(l2 == int32(16)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v151 = v140
	goto L1
L6:
	;
	v142 = v30 + int32(1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v142 < v143 {
		v28 = v140
		v30 = v142
		goto L4
	} else {
		goto L45
	}
L7:
	;
	if v111 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L8:
	;
	v109 = F_transformExpr(m, l0, v37, l2)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L16
	} else {
		goto L35
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	switch v40 - int32(69) {
	case 0:
		goto L14
	default:
		goto L12
	case 10:
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v104 == int32(57) {
		v111 = v103
		v113 = v37
		goto L7
	} else {
		goto L34
	}
L12:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v107 = v102
	goto L8
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v63+v64<<(uint(int32(2))%32)-int32(4))))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v71 != int32(77) {
		goto L12
	} else {
		goto L19
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44+v45<<(uint(int32(2))%32)-int32(4))))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v52 != int32(77) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v56 = F_ExpandColumnRefStar(m, l0, v37, int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	v60 = F_list_concat(m, v28, v56)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v140 = v60
	goto L6
L19:
	;
	v74 = F_copyObjectImpl(m, v37)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	if v76 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v81 = v77 - int32(1)
	goto L23
L22:
	;
	v81 = int32(-1)
	goto L23
L23:
	;
	v82 = int32(0)
	if base.B2i32(v76 == v82)|base.B2i32(v81 <= v82) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v92
	v94 = F_transformExpr(m, l0, v74, l2)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L16
	} else {
		goto L31
	}
L25:
	;
	v92 = int32(0)
	goto L27
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v81 < v89 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = v81
	goto L30
L29:
	;
	goto L30
L30:
	;
	v92 = v76
	goto L27
L31:
	;
	v97 = F_ExpandRowReference(m, l0, v94, int32(1))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v99 = F_list_concat(m, v28, v97)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	v140 = v99
	goto L6
L34:
	;
	v107 = v103
	goto L8
L35:
	;
	v111 = v107
	v113 = v109
	goto L7
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
	v120 = F_FigureColnameInternal(m, v37, v13+int32(12))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L16
	} else {
		goto L39
	}
L37:
	;
	v126 = v111
	goto L38
L38:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v127 + int32(1)
	v133 = F_makeTargetEntry(m, v113, base.I32_extend16_s(v127), v126, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L16
	} else {
		goto L43
	}
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v122 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v124 = v122
	goto L42
L41:
	;
	v124 = int32(_a_F_transformTargetList_0)
	goto L42
L42:
	;
	v126 = v124
	goto L38
L43:
	;
	v135 = F_lappend(m, v28, v133)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	v140 = v135
	goto L6
L45:
	;
	goto L5
L46:
	;
	v156 = F_list_concat(m, v151, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L16
	} else {
		goto L49
	}
L47:
	;
	v160 = v151
	goto L48
L48:
	;
	m.G0 = v13 + int32(16)
	return v160
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = int32(0)
	v160 = v156
	goto L48
}
