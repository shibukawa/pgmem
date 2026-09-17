package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_relation_filenode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_SearchSysCache1(m, int32(57), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v139
L2:
	;
	v136 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v136)
	v139 = int32(0)
	goto L1
L3:
	;
	return int32(0)
L4:
	;
	if v8 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	goto L2
L6:
	;
	goto L7
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
	v16 = v14 + v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+119)))
	switch v17 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L9
	default:
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
	if v22 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_ReleaseCatCache(m, v8)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	F_ReleaseCatCache(m, v8)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L3
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+117)))
	v27 = int32(0)
	if v26 == v27 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	return v22
L16:
	;
	F_ReleaseCatCache(m, v8)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L3
	} else {
		goto L54
	}
L17:
	;
	goto L16
L18:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v129 = v124
	goto L17
L19:
	;
	v32 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filenode[0]))
	if v32 < v34 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v75 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filenode[1]))
	if v75 < v77 {
		goto L40
	} else {
		goto L41
	}
L22:
	;
	v38 = v32
	goto L25
L23:
	;
	goto L24
L24:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filenode[2]))
	if v57 <= int32(0) {
		v129 = v27
		goto L17
	} else {
		goto L31
	}
L25:
	;
	v43 = v38 << (uint(int32(3)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_pg_relation_filenode[3])))
	if v44 == v7 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	v123 = v43 + int32(_a_F_pg_relation_filenode_0)
	goto L18
L28:
	;
	goto L29
L29:
	;
	v49 = v38 + int32(1)
	if v49 != v34 {
		v38 = v49
		goto L25
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	v62 = int32(0)
	goto L32
L32:
	;
	v67 = v62 << (uint(int32(3)) % 32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_pg_relation_filenode[4])))
	if v68 != v7 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v123 = v67 + int32(_a_F_pg_relation_filenode_1)
	goto L18
L34:
	;
	v71 = v62 + int32(1)
	if v57 != v71 {
		v62 = v71
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	v129 = v27
	goto L17
L38:
	;
	v105 = int32(0)
	goto L48
L39:
	;
	v123 = v86 + int32(_a_F_pg_relation_filenode_2)
	goto L18
L40:
	;
	v81 = v75
	goto L43
L41:
	;
	goto L42
L42:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filenode[5]))
	if v98 <= int32(0) {
		v129 = v27
		goto L17
	} else {
		goto L47
	}
L43:
	;
	v86 = v81 << (uint(int32(3)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+uint32(_c_F_pg_relation_filenode[6])))
	if v7 == v87 {
		goto L39
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	v90 = v81 + int32(1)
	if v90 != v77 {
		v81 = v90
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L38
L48:
	;
	v110 = v105 << (uint(int32(3)) % 32)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+uint32(_c_F_pg_relation_filenode[7])))
	if v111 != v7 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v123 = v110 + int32(_a_F_pg_relation_filenode_3)
	goto L18
L50:
	;
	v114 = v105 + int32(1)
	if v98 != v114 {
		v105 = v114
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	v129 = v27
	goto L17
L54:
	;
	if v129 != 0 {
		v139 = v129
		goto L1
	} else {
		goto L55
	}
L55:
	;
	goto L8
}
