package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_find_all_inheritors(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = int64(34359738372)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_find_all_inheritors[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v19
	v22 = int32(32)
	v26 = F_hash_create(m, int32(_a_F_find_all_inheritors_0), v22, v14+v22, int32(1064))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l0
	v35 = F_list_make1_impl(m, int32(472), v14+int32(12))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v37
	v44 = F_list_make1_impl(m, int32(471), v14+int32(8))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v35 == int32(0) {
		v149 = v4
		v150 = v44
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l2 != 0 {
		goto L31
	} else {
		goto L32
	}
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v48 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v149 = v35
	v150 = v44
	goto L5
L8:
	;
	goto L9
L9:
	;
	v56 = v35
	v57 = v44
	v61 = v4
	goto L10
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v61<<(uint(int32(2))%32))))
	v68 = int32(0)
	v70 = F_find_inheritance_children_extended(m, v66, int32(1), l1, v68, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v149 = v134
	v150 = v135
	goto L5
L12:
	;
	v141 = v61 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v141 < v142 {
		v56 = v134
		v57 = v135
		v61 = v141
		goto L10
	} else {
		goto L29
	}
L13:
	;
	if v70 == int32(0) {
		v134 = v56
		v135 = v57
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v74 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v75 <= v74 {
		v134 = v56
		v135 = v57
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v78 = v74
	v83 = v56
	v84 = v57
	goto L16
L16:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v78<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v93
	v100 = F_hash_search(m, v26, v14+int32(20), int32(1), v14+int32(19))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	v134 = v122
	v135 = v123
	goto L12
L18:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+19)))
	if v102 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v126 = v78 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v126 < v127 {
		v78 = v126
		v83 = v122
		v84 = v123
		goto L16
	} else {
		goto L28
	}
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v107 = v103 + v104<<(uint(int32(2))%32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v108 + int32(1)
	v122 = v83
	v123 = v84
	goto L19
L21:
	;
	goto L22
L22:
	;
	if v83 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v114 = v112
	goto L25
L24:
	;
	v114 = int32(0)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v117 = F_lappend_oid(m, v83, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v120 = F_lappend_int(m, v84, int32(1))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v122 = v117
	v123 = v120
	goto L19
L28:
	;
	goto L17
L29:
	;
	goto L11
L30:
	;
	F_hash_destroy(m, v26)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L35
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v150
	goto L30
L32:
	;
	goto L33
L33:
	;
	F_list_free(m, v150)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	m.G0 = v14 + int32(80)
	return v149
}
