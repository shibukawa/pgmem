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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v10&int32(2) != 0 {
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
	v17 = int32(5646)
	goto L6
L5:
	;
	v17 = int32(5645)
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+10)))
	v20 = v18 + v19
	v23 = v13 + int32(8)
	v26 = v14
	goto L7
L7:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v33&int32(4) != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L3
L9:
	;
	v60 = int32(1)
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	if v60 < v26 {
		v23 = v23 + (v62+int32(9))&int32(131064)
		v26 = v26 - v60
		goto L7
	} else {
		goto L24
	}
L10:
	;
	return v57
L11:
	;
	v36 = int32(1)
	v39 = F_compare_subnode(m, v23, v20, v32, v17, v33&v36)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	if v43 != v32 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	return int32(0)
L15:
	;
	if v39 != 0 {
		v57 = v36
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	if v33&int32(1) == int32(0) {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v52 = m.T0[v17].(func(*base.Module, int32, int32, int32, int32) int32)(m, v20, v32, v23+int32(2), v43)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L14
	} else {
		goto L22
	}
L20:
	;
	if base.Ui32(v43) <= base.Ui32(v32) {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if v52 == int32(0) {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v57 = int32(1)
	goto L10
L24:
	;
	goto L8
}
