package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetMaxSnapshotXidCount(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_GetMaxSnapshotXidCount[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	return v3
}
func F_max_parallel_hazard_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v9 = l0
	goto L4
L4:
	;
	v13 = int32(1)
	v15 = F_check_functions_in_node(m, v9, int32(864), l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	return int32(0)
L6:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v9+v131)))
	if v134 != 0 {
		v9 = v134
		goto L4
	} else {
		goto L53
	}
L7:
	;
	return v128
L8:
	;
	return int32(0)
L9:
	;
	if v15 != 0 {
		v128 = v13
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	switch v19 - int32(8) {
	case 0:
		goto L14
	case 1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 48, 49, 50, 52, 53, 54, 55, 56, 57, 58:
		goto L11
	case 3:
		goto L17
	case 14:
		goto L16
	case 15:
		goto L15
	case 47:
		goto L19
	case 51:
		goto L18
	case 59:
		goto L13
	default:
		goto L12
	}
L11:
	;
	v126 = F_expression_tree_walker_impl(m, v9, int32(865), l1)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L8
	} else {
		goto L52
	}
L12:
	;
	if v19 == int32(318) {
		v131 = int32(4)
		goto L6
	} else {
		goto L51
	}
L13:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v9)+140))
	if v111 != 0 {
		goto L47
	} else {
		goto L48
	}
L14:
	;
	v62 = int32(0)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	switch v63 {
	case 0:
		v128 = v62
		goto L7
	case 1:
		goto L32
	default:
		goto L31
	}
L15:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+38)))
	if v41 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v36 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v36)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v38 != v36 {
		goto L11
	} else {
		goto L22
	}
L17:
	;
	v31 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v31)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v33 != v31 {
		goto L11
	} else {
		goto L21
	}
L18:
	;
	v27 = int32(117)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v27)
	return int32(1)
L19:
	;
	v22 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v22)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v24 != v22 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	v128 = v13
	goto L7
L21:
	;
	v128 = v13
	goto L7
L22:
	;
	v128 = v13
	goto L7
L23:
	;
	v44 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v44)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v46 == v44 {
		v128 = v13
		goto L7
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v51 = F_list_concat_copy(m, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v51
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v55 = F_max_parallel_hazard_walker(m, v54, l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	if v55 != 0 {
		v128 = v13
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_list_free(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v49
	v131 = int32(48)
	goto L6
L31:
	;
	v105 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v105)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	return base.B2i32(v107 == v105)
L32:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v66 = int32(0)
	if v64 == v66 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v104 != 0 {
		v128 = v62
		goto L7
	} else {
		goto L46
	}
L34:
	;
	v104 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v72 <= int32(0) {
		v97 = v66
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v104 = v97
	goto L33
L38:
	;
	v75 = int32(0)
	if v75 < v72 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v78 = v72
	goto L41
L40:
	;
	v78 = v75
	goto L41
L41:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v81 = int32(0)
	goto L42
L42:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v79+v81<<(uint(int32(2))%32))))
	v90 = base.B2i32(v89 == v65)
	if v89 == v65 {
		v97 = v90
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v97 = v90
	goto L37
L44:
	;
	v92 = v81 + int32(1)
	if v92 != v78 {
		v81 = v92
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	goto L31
L47:
	;
	v112 = int32(117)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v112)
	return int32(1)
L48:
	;
	goto L49
L49:
	;
	v118 = F_query_tree_walker_impl(m, v9, int32(865), l1, int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	return v118
L51:
	;
	goto L11
L52:
	;
	v128 = v126
	goto L7
L53:
	;
	goto L5
}
