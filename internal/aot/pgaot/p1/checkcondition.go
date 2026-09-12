package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_checkcondition_QueryOperand(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = base.I32_div_s(l1-v7-int32(8), int32(12))
	v15 = v6 + v12*int32(32776)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v16 != int32(1) {
		v41 = int32(0)
	} else {
		if l2 == int32(0) {
			v41 = int32(1)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v24 = v15 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v24
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
			v27 = int32(1)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
			if v28 != v27 {
				v41 = v27
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v34 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v24 + (int32(16384)-v32)<<(uint(v34)%32)
				v41 = v34
			}
		}
	}
	return v41
}
func F_checkcondition_arr_1(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	if v7 != 0 {
		v45 = int32(2)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v45
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v8) < base.Ui32(v9) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v12 = v9
	v14 = v8
	goto L6
L4:
	;
	goto L5
L5:
	;
	v45 = int32(0)
	goto L1
L6:
	;
	v17 = int32(2)
	v22 = base.I32_div_s((v12-v14)>>(uint(v17)%32), v17)
	v25 = v14 + v22<<(uint(v17)%32)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v26 == v11 {
		v45 = v17
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v30 = base.B2i32(v26 < v11)
	if v26 < v11 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = v25 + int32(4)
	goto L11
L10:
	;
	v31 = v14
	goto L11
L11:
	;
	if v26 < v11 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v32 = v12
	goto L14
L13:
	;
	v32 = v25
	goto L14
L14:
	;
	if base.Ui32(v31) < base.Ui32(v32) {
		v12 = v32
		v14 = v31
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L7
}
func F_checkcondition_bit_3(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v5 = int32(3)
	v7 = base.I32_rem_u_s(v4, l2<<(uint(v5)%32))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(base.Ui32(v7)>>(uint(v5)%32))))))
	return int32(base.Ui32(v11)>>(uint(v7&int32(7))%32)) & int32(1)
}
