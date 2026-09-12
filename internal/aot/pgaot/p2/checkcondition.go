package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_checkcondition_HL(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v79 int32
	_ = v79
	v4 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v4 < v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = v9
	v17 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	if l2 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	v21 = v17 << (uint(int32(4)) % 32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = v21 + v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v24 != l1 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v66 = v17 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v66 < v67 {
		v16 = v67
		v17 = v66
		goto L4
	} else {
		goto L17
	}
L7:
	;
	if l2 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(1)
L9:
	;
	goto L10
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = F_palloc(m, v16<<(uint(int32(1))%32))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v52 = v30 + v49<<(uint(int32(1))%32)
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52-int32(2)))))
	if base.Ui32(v48) <= base.Ui32(v55) {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	return int32(0)
L15:
	;
	v39 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v39)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v39
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44+v21)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v46)
	goto L6
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v49 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v52))) = uint16(v48)
	goto L6
L17:
	;
	goto L5
L18:
	;
	return int32(0)
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v79 <= int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	return int32(1)
}
func F_checkcondition_bit_2(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v5&int32(21) != 0 {
		v24 = int32(1)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v11 = int32(3)
		v13 = base.I32_rem_u_s(v9, v10<<(uint(v11)%32))
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(base.Ui32(v13)>>(uint(v11)%32))))))
		v24 = int32(base.Ui32(v17)>>(uint(v13&int32(7))%32)) & int32(1)
	}
	return v24
}
