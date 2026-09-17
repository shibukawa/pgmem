package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_rule_groupingset(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v12 = int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v13 {
	case 0:
		goto L7
	case 1:
		goto L6
	case 2:
		v103 = int32(_a_F_get_rule_groupingset_0)
		v104 = v12
		goto L3
	case 3:
		goto L5
	case 4:
		goto L4
	default:
		v108 = v12
		goto L2
	}
L1:
	;
	return
L2:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v109 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L3:
	;
	F_appendStringInfoString(m, v10, v103)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L32
	}
L4:
	;
	v103 = int32(_a_F_get_rule_groupingset_1)
	v104 = int32(0)
	goto L3
L5:
	;
	v103 = int32(_a_F_get_rule_groupingset_2)
	v104 = v12
	goto L3
L6:
	;
	if l2 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	F_appendStringInfoString(m, v10, int32(_a_F_get_rule_groupingset_3))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	return
L10:
	;
	if l2 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	F_appendStringInfoString(m, v10, int32(_a_F_get_rule_groupingset_4))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L19
	}
L12:
	;
	F_appendStringInfoChar(m, v10, int32(40))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L16
	}
L13:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v19 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 != int32(1) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v39 = v19
	v40 = l0 + int32(8)
	goto L11
L16:
	;
	v32 = l0 + int32(8)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v33 == int32(0) {
		v85 = v32
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 <= int32(0) {
		v85 = v32
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v39 = v33
	v40 = v32
	goto L11
L19:
	;
	v49 = F_get_rule_sortgroupclause(m, v44, l1, int32(0), l3)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v51 < int32(2) {
		v85 = v40
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v55 = int32(1)
	goto L22
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v55<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v10, int32(_a_F_get_rule_groupingset_5))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L8
	} else {
		goto L24
	}
L23:
	;
	v85 = v40
	goto L10
L24:
	;
	v73 = F_get_rule_sortgroupclause(m, v68, l1, int32(0), l3)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v76 = v55 + int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v39+int32(4))))
	if v76 < v77 {
		v55 = v76
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	F_appendStringInfoChar(m, v10, int32(41))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L31
	}
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v90 == int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v93 == int32(1) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	return
L32:
	;
	v108 = v104
	goto L2
L33:
	;
	F_appendStringInfoChar(m, v10, int32(41))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L8
	} else {
		goto L44
	}
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v113 <= int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	F_appendStringInfoString(m, v10, int32(_a_F_get_rule_groupingset_4))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	F_get_rule_groupingset(m, v120, l1, v108, l3)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v123 < int32(2) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v126 = int32(1)
	goto L39
L39:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	F_appendStringInfoString(m, v10, int32(_a_F_get_rule_groupingset_5))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L8
	} else {
		goto L41
	}
L40:
	;
	goto L33
L41:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v135+v126<<(uint(int32(2))%32))))
	F_get_rule_groupingset(m, v142, l1, v108, l3)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v146 = v126 + int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v146 < v147 {
		v126 = v146
		goto L39
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	goto L1
}
func F_get_rule_windowspec(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v8, int32(40))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v13 = F_quote_identifier(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v17 = int32(0)
	v18 = base.B2i32(v12 != v17)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v19 == v17 {
		v79 = v18
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_appendStringInfoString(m, v8, v13)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v81 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 != 0 {
		v79 = v18
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if v12 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_appendStringInfoChar(m, v8, int32(32))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_appendStringInfoString(m, v8, int32(_a_F_get_rule_windowspec_0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v29 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v79 = int32(1)
	goto L8
L17:
	;
	goto L18
L18:
	;
	v33 = int32(1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v34 <= int32(0) {
		v79 = v33
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	F_appendStringInfoString(m, v8, int32(_a_F_get_rule_windowspec_1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v44 = F_get_rule_sortgroupclause(m, v42, l1, int32(0), l2)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v47 < int32(2) {
		v79 = v33
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v53 = int32(1)
	goto L23
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v53<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v8, int32(_a_F_get_rule_windowspec_2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v79 = v69
	goto L8
L25:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v67 = F_get_rule_sortgroupclause(m, v65, l1, int32(0), l2)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v69 = int32(1)
	v71 = v53 + v69
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v71 < v72 {
		v53 = v71
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	F_appendStringInfoChar(m, v8, int32(41))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L45
	}
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_get_window_frame_options(m, v114, v115, v116, l2)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L44
	}
L30:
	;
	F_appendStringInfoChar(m, v8, int32(32))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L43
	}
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v100&int32(1) == int32(0) {
		goto L28
	} else {
		goto L41
	}
L32:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v84 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if v79 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_appendStringInfoChar(m, v8, int32(32))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_appendStringInfoString(m, v8, int32(_a_F_get_rule_windowspec_3))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_get_rule_orderby(m, v91, l1, int32(0), l2)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v95&int32(1) == int32(0) {
		goto L28
	} else {
		goto L40
	}
L40:
	;
	goto L30
L41:
	;
	if v79 == int32(0) {
		v114 = v100
		goto L29
	} else {
		goto L42
	}
L42:
	;
	goto L30
L43:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20))))
	v114 = v113
	goto L29
L44:
	;
	goto L28
L45:
	;
	return
}
