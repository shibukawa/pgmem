package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_expression_returns_set_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v3 = int32(0)
	if l0 == v3 {
		v29 = v3
		return v29
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v7 - int32(15) {
		case 0:
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
			if v10 == int32(0) {
				v24 = F_expression_tree_walker_impl_x2especialized_x2e2(m, l0, l1)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v29 = v24
					return v29
				}
			} else {
				return int32(1)
			}
		default:
			if base.Ui32(v7-int32(9)) < base.Ui32(int32(3)) {
				v29 = v3
				return v29
			} else {
				v24 = F_expression_tree_walker_impl_x2especialized_x2e2(m, l0, l1)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v29 = v24
					return v29
				}
			}
		case 2:
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
			if v15 == int32(0) {
				v24 = F_expression_tree_walker_impl_x2especialized_x2e2(m, l0, l1)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v29 = v24
					return v29
				}
			} else {
				return int32(1)
			}
		}
	}
}
func F_transformExpressionList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	v5 = int32(0)
	if l1 == v5 {
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v14 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = v5
	v25 = v5
	goto L7
L5:
	;
	v117 = v5
	goto L6
L6:
	;
	return v117
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v25<<(uint(int32(2))%32))))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	switch v31 - int32(69) {
	case 0:
		goto L12
	default:
		goto L10
	case 10:
		goto L11
	}
L8:
	;
	v117 = v105
	goto L6
L9:
	;
	v107 = v25 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v107 < v108 {
		v24 = v105
		v25 = v107
		goto L7
	} else {
		goto L40
	}
L10:
	;
	if v31 == int32(57) {
		goto L32
	} else {
		goto L33
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v54+v55<<(uint(int32(2))%32)-int32(4))))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v62 != int32(77) {
		goto L10
	} else {
		goto L17
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36<<(uint(int32(2))%32)-int32(4))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 != int32(77) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v47 = F_ExpandColumnRefStar(m, l0, v30, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v51 = F_list_concat(m, v24, v47)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v105 = v51
	goto L9
L17:
	;
	v65 = F_copyObjectImpl(m, v30)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if v67 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v72 = v68 - int32(1)
	goto L21
L20:
	;
	v72 = int32(-1)
	goto L21
L21:
	;
	v73 = int32(0)
	if base.B2i32(v67 == v73)|base.B2i32(v72 <= v73) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v83
	v85 = F_transformExpr(m, l0, v65, l2)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L14
	} else {
		goto L29
	}
L23:
	;
	v83 = int32(0)
	goto L25
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v72 < v80 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = v72
	goto L28
L27:
	;
	goto L28
L28:
	;
	v83 = v67
	goto L25
L29:
	;
	v88 = F_ExpandRowReference(m, l0, v85, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v90 = F_list_concat(m, v24, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	v105 = v90
	goto L9
L32:
	;
	v96 = l3
	goto L34
L33:
	;
	v96 = int32(0)
	goto L34
L34:
	;
	if v96 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v99 = v30
	goto L37
L36:
	;
	v97 = F_transformExpr(m, l0, v30, l2)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L14
	} else {
		goto L38
	}
L37:
	;
	v100 = F_lappend(m, v24, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L14
	} else {
		goto L39
	}
L38:
	;
	v99 = v97
	goto L37
L39:
	;
	v105 = v100
	goto L9
L40:
	;
	goto L8
}
