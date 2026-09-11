package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_out_grouping_b_U(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = v18
	goto L2
L1:
	;
	return v125
L2:
	;
	if v28 <= v19 {
		v125 = int32(-1)
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v125 = int32(0)
	goto L1
L4:
	;
	v36 = int32(1)
	v37 = v28 - v36
	v39 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15+v37))))
	v41 = v39 & int32(255)
	if v37 == v19 {
		v96 = v41
		v97 = v36
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l3 < v96 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	if int32(0) <= v39 {
		v96 = v41
		v97 = v36
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v47 = v41 & int32(63)
	v49 = v28 - int32(2)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v49))))
	v53 = v51 << (uint(int32(6)) % 32)
	if base.B2i32(v49 != v19)&base.B2i32(base.Ui32(v51) < base.Ui32(int32(192))) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v96 = v53&int32(1984) | v47
	v97 = int32(2)
	goto L5
L9:
	;
	goto L10
L10:
	;
	v66 = v53&int32(4032) | v47
	v68 = v28 - int32(3)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v68))))
	if base.B2i32(v68 != v19)&base.B2i32(base.Ui32(v70) < base.Ui32(int32(224))) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v96 = v70<<(uint(int32(12))%32)&int32(61440) | v66
	v97 = int32(3)
	goto L5
L12:
	;
	goto L13
L13:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+(v15-int32(4))))))
	v96 = v70<<(uint(int32(12))%32)&int32(258048) | v88&int32(7)<<(uint(int32(18))%32) | v66
	v97 = int32(4)
	goto L5
L14:
	;
	v117 = v28 - v97
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
	if l4 != 0 {
		v28 = v117
		goto L2
	} else {
		goto L18
	}
L15:
	;
	v101 = v96 - l2
	if v101 < int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v101)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v107)>>(uint(v101&int32(7))%32))&int32(1) == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	return v97
L18:
	;
	goto L3
}
