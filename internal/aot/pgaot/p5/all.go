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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v82 int32
	_ = v82
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = int64(34359738372)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v19
	v22 = int32(32)
	v26 = F_hash_create(m, int32(389876), v22, v14+v22, int32(1064))
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
		v150 = v4
		v152 = v44
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
	v150 = v35
	v152 = v44
	goto L5
L8:
	;
	goto L9
L9:
	;
	v55 = v35
	v57 = v44
	v59 = v4
	goto L10
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v59<<(uint(int32(2))%32))))
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
	v150 = v135
	v152 = v137
	goto L5
L12:
	;
	v143 = v59 + int32(1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v143 < v144 {
		v55 = v135
		v57 = v137
		v59 = v143
		goto L10
	} else {
		goto L29
	}
L13:
	;
	if v70 == int32(0) {
		v135 = v55
		v137 = v57
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v74 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v75 <= v74 {
		v135 = v55
		v137 = v57
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v78 = v74
	v82 = v55
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
	v135 = v124
	v137 = v125
	goto L12
L18:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+19)))
	if v102 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v128 = v78 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v128 < v129 {
		v78 = v128
		v82 = v124
		v84 = v125
		goto L16
	} else {
		goto L28
	}
L20:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v109 = v105 + v106<<(uint(int32(2))%32)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v110 + int32(1)
	v124 = v82
	v125 = v84
	goto L19
L21:
	;
	goto L22
L22:
	;
	if v82 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v116 = v114
	goto L25
L24:
	;
	v116 = int32(0)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v119 = F_lappend_oid(m, v82, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v122 = F_lappend_int(m, v84, int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v124 = v119
	v125 = v122
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
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L35
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v152
	goto L30
L32:
	;
	goto L33
L33:
	;
	F_list_free(m, v152)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
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
	return v150
}
