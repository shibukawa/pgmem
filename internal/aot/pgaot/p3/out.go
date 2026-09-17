package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_out_grouping_b_U(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = v14
	goto L2
L1:
	;
	return v125
L2:
	;
	if v25 <= v15 {
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
	v32 = int32(1)
	v33 = v25 - v32
	v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16+v33))))
	v37 = v35 & int32(255)
	if base.B2i32(v33 == v15)|base.B2i32(int32(0) <= v35) != 0 {
		v95 = v37
		v99 = v32
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l3 < v95 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v44 = v37 & int32(63)
	v46 = v25 - int32(2)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v46))))
	v50 = v48 << (uint(int32(6)) % 32)
	if base.B2i32(v46 != v15)&base.B2i32(base.Ui32(v48) < base.Ui32(int32(192))) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v95 = v50&int32(1984) | v44
	v99 = int32(2)
	goto L5
L8:
	;
	goto L9
L9:
	;
	v63 = v50&int32(4032) | v44
	v65 = v25 - int32(3)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v65))))
	if base.B2i32(v65 != v15)&base.B2i32(base.Ui32(v67) < base.Ui32(int32(224))) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v95 = v67<<(uint(int32(12))%32)&int32(_a_F_out_grouping_b_U_0) | v63
	v99 = int32(3)
	goto L5
L11:
	;
	goto L12
L12:
	;
	v85 = int32(4)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v16-v85))))
	v95 = v67<<(uint(int32(12))%32)&int32(_a_F_out_grouping_b_U_1) | v87&int32(7)<<(uint(int32(18))%32) | v63
	v99 = v85
	goto L5
L13:
	;
	v117 = v25 - v99
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
	if l4 != 0 {
		v25 = v117
		goto L2
	} else {
		goto L17
	}
L14:
	;
	v101 = v95 - l2
	if v101 < int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v101)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v107)>>(uint(v101&int32(7))%32))&int32(1) == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	return v99
L17:
	;
	goto L3
}
