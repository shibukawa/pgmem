package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_checkcondition_gin_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = base.I32_div_s(l1-v6, int32(3))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v5+v9)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+v11))))
	if v13 == int32(1) {
		v16 = int32(2)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v19 != 0 {
			v20 = v16
		} else {
			v20 = int32(1)
		}
		if l2 != 0 {
			v21 = v16
		} else {
			v21 = v20
		}
		v22 = v21
	} else {
		v22 = v13
	}
	return base.I32_extend8_s(v22)
}
func F_checkcondition_str_2(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v12 = v10 & int32(1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+10)))
	v15 = v13 + v14
	v21 = base.B2i32(v10&int32(2) != int32(0))
	v22 = v8 + int32(8)
	v26 = v9
	goto L4
L2:
	;
	goto L3
L3:
	;
	return int32(0)
L4:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v30&int32(4) != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22))))
	v56 = int32(1)
	if v56 < v26 {
		v22 = v22 + (v50+int32(9))&int32(_a_F_checkcondition_str_2_0)
		v26 = v26 - v56
		goto L4
	} else {
		goto L15
	}
L7:
	;
	v33 = F_compare_subnode(m, v22, v15, v29, v12, v21)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22))))
	v44 = F_ltree_label_match(m, v15, v29, v22+int32(2), v43, v12, v21)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L13
	}
L10:
	;
	return int32(0)
L11:
	;
	if v33 == int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	return int32(1)
L13:
	;
	if v44 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	return int32(1)
L15:
	;
	goto L5
}
