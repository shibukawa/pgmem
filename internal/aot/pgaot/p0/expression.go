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
	var v28 int32
	_ = v28
	v3 = int32(0)
	if l0 == v3 {
		v28 = v3
		return v28
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
					v28 = v24
					return v28
				}
			} else {
				return int32(1)
			}
		default:
			if base.Ui32(v7-int32(9)) < base.Ui32(int32(3)) {
				v28 = v3
				return v28
			} else {
				v24 = F_expression_tree_walker_impl_x2especialized_x2e2(m, l0, l1)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = v24
					return v28
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
					v28 = v24
					return v28
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
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	v5 = int32(0)
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v109
L2:
	;
	v18 = v5
	v21 = v5
	goto L7
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v10 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v109 = v5
	goto L1
L6:
	;
	goto L5
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v21<<(uint(int32(2))%32))))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	switch v27 - int32(69) {
	case 0:
		goto L12
	default:
		goto L10
	case 10:
		goto L11
	}
L8:
	;
	v109 = v99
	goto L1
L9:
	;
	v101 = v21 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v101 < v102 {
		v18 = v99
		v21 = v101
		goto L7
	} else {
		goto L40
	}
L10:
	;
	if v27 == int32(57) {
		goto L32
	} else {
		goto L33
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50+v51<<(uint(int32(2))%32)-int32(4))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v58 != int32(77) {
		goto L10
	} else {
		goto L17
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31+v32<<(uint(int32(2))%32)-int32(4))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v39 != int32(77) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v43 = F_ExpandColumnRefStar(m, l0, v26, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v47 = F_list_concat(m, v18, v43)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v99 = v47
	goto L9
L17:
	;
	v61 = F_copyObjectImpl(m, v26)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	if v63 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v68 = v64 - int32(1)
	goto L21
L20:
	;
	v68 = int32(-1)
	goto L21
L21:
	;
	v69 = int32(0)
	if v63 == v69 {
		v77 = v69
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = v77
	v79 = F_transformExpr(m, l0, v61, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L14
	} else {
		goto L29
	}
L23:
	;
	goto L22
L24:
	;
	if v68 <= int32(0) {
		v77 = v69
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v68 < v74 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v68
	goto L28
L27:
	;
	goto L28
L28:
	;
	v77 = v63
	goto L23
L29:
	;
	v82 = F_ExpandRowReference(m, l0, v79, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	v84 = F_list_concat(m, v18, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	v99 = v84
	goto L9
L32:
	;
	v90 = l3
	goto L34
L33:
	;
	v90 = int32(0)
	goto L34
L34:
	;
	if v90 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v93 = v26
	goto L37
L36:
	;
	v91 = F_transformExpr(m, l0, v26, l2)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L14
	} else {
		goto L38
	}
L37:
	;
	v94 = F_lappend(m, v18, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L14
	} else {
		goto L39
	}
L38:
	;
	v93 = v91
	goto L37
L39:
	;
	v99 = v94
	goto L9
L40:
	;
	goto L8
}
