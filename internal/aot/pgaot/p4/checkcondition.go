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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v9&int32(2) != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return int32(0)
L4:
	;
	v16 = int32(_a_F_checkcondition_str_2_0)
	goto L6
L5:
	;
	v16 = int32(_a_F_checkcondition_str_2_1)
	goto L6
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+10)))
	v19 = v17 + v18
	v22 = v10 + int32(8)
	v24 = v11
	goto L7
L7:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v31&int32(4) != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L3
L9:
	;
	v59 = int32(1)
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22))))
	if v59 < v24 {
		v22 = v22 + (v61+int32(9))&int32(_a_F_checkcondition_str_2_2)
		v24 = v24 - v59
		goto L7
	} else {
		goto L20
	}
L10:
	;
	return int32(1)
L11:
	;
	v36 = F_compare_subnode(m, v22, v19, v30, v16, v31&int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22))))
	if base.B2i32(v40 != v30)&(base.B2i32(v31&int32(1) == int32(0))|base.B2i32(base.Ui32(v40) <= base.Ui32(v30))) != 0 {
		goto L9
	} else {
		goto L17
	}
L14:
	;
	return int32(0)
L15:
	;
	if v36 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	v51 = m.T0[v16].(func(*base.Module, int32, int32, int32, int32) int32)(m, v19, v30, v22+int32(2), v40)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	if v51 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	goto L10
L20:
	;
	goto L8
}
