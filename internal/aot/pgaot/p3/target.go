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
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v151 != 0 {
		goto L49
	} else {
		goto L50
	}
L2:
	;
	v25 = v4
	v28 = v4
	goto L7
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v15 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v146 = v4
	goto L1
L6:
	;
	goto L5
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v28<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	if base.B2i32(l2 == int32(16)) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v146 = v136
	goto L1
L9:
	;
	v138 = v28 + int32(1)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v138 < v139 {
		v25 = v136
		v28 = v138
		goto L7
	} else {
		goto L48
	}
L10:
	;
	if v107 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L11:
	;
	v105 = F_transformExpr(m, l0, v35, l2)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L19
	} else {
		goto L38
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	switch v38 - int32(69) {
	case 0:
		goto L17
	default:
		goto L15
	case 10:
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v100 == int32(57) {
		v107 = v99
		v109 = v35
		goto L10
	} else {
		goto L37
	}
L15:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v103 = v98
	goto L11
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61+v62<<(uint(int32(2))%32)-int32(4))))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v69 != int32(77) {
		goto L15
	} else {
		goto L22
	}
L17:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42+v43<<(uint(int32(2))%32)-int32(4))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 != int32(77) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v54 = F_ExpandColumnRefStar(m, l0, v35, int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	v58 = F_list_concat(m, v25, v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v136 = v58
	goto L9
L22:
	;
	v72 = F_copyObjectImpl(m, v35)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v74 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v79 = v75 - int32(1)
	goto L26
L25:
	;
	v79 = int32(-1)
	goto L26
L26:
	;
	v80 = int32(0)
	if v74 == v80 {
		v88 = v80
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v88
	v90 = F_transformExpr(m, l0, v72, l2)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L19
	} else {
		goto L34
	}
L28:
	;
	goto L27
L29:
	;
	if v79 <= int32(0) {
		v88 = v80
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v79 < v85 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v79
	goto L33
L32:
	;
	goto L33
L33:
	;
	v88 = v74
	goto L28
L34:
	;
	v93 = F_ExpandRowReference(m, l0, v90, int32(1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	v95 = F_list_concat(m, v25, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	v136 = v95
	goto L9
L37:
	;
	v103 = v99
	goto L11
L38:
	;
	v107 = v103
	v109 = v105
	goto L10
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
	v116 = F_FigureColnameInternal(m, v35, v13+int32(12))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L19
	} else {
		goto L42
	}
L40:
	;
	v122 = v107
	goto L41
L41:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v123 + int32(1)
	v129 = F_makeTargetEntry(m, v109, base.I32_extend16_s(v123), v122, int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L19
	} else {
		goto L46
	}
L42:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v118 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v120 = v118
	goto L45
L44:
	;
	v120 = int32(567288)
	goto L45
L45:
	;
	v122 = v120
	goto L41
L46:
	;
	v131 = F_lappend(m, v25, v129)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L19
	} else {
		goto L47
	}
L47:
	;
	v136 = v131
	goto L9
L48:
	;
	goto L8
L49:
	;
	v152 = F_list_concat(m, v146, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L19
	} else {
		goto L52
	}
L50:
	;
	v156 = v146
	goto L51
L51:
	;
	m.G0 = v13 + int32(16)
	return v156
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = int32(0)
	v156 = v152
	goto L51
}
