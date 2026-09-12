package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_u_isupper(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(0)
	v11 = int32(655)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return base.B2i32(base.Ui32(l0-int32(65)) < base.Ui32(int32(26)))
L4:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 << (uint(int32(3)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[605])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	if v33 <= v34 {
		v10 = v33
		v11 = v34
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v33 = v16 + int32(1)
	v34 = v11
	goto L6
L8:
	;
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[606])))
	if base.Ui32(v27) <= base.Ui32(l0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(1)
L11:
	;
	goto L12
L12:
	;
	v33 = v10
	v34 = v16 - int32(1)
	goto L6
L13:
	;
	goto L5
}
func F_pg_u_prop_case_ignorable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = int32(0)
	v11 = int32(505)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[595]))))
	return int32(base.Ui32(v42&int32(16)) >> (uint(int32(4)) % 32))
L4:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 << (uint(int32(3)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[1285])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	if v33 <= v34 {
		v10 = v33
		v11 = v34
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v33 = v16 + int32(1)
	v34 = v11
	goto L6
L8:
	;
	goto L9
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[1286])))
	if base.Ui32(v27) <= base.Ui32(l0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(1)
L11:
	;
	goto L12
L12:
	;
	v33 = v10
	v34 = v16 - int32(1)
	goto L6
L13:
	;
	goto L5
}
func F_pg_u_prop_cased(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v54 = int32(689)
	v55 = int32(0)
	goto L15
L2:
	;
	v10 = int32(3367)
	v11 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[595]))))
	return int32(base.Ui32(v43&int32(8)) >> (uint(int32(3)) % 32))
L5:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 * int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[591])))
	if base.Ui32(v21) < base.Ui32(l0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[593]))))
	if v34 != int32(3) {
		goto L1
	} else {
		goto L14
	}
L7:
	;
	goto L6
L8:
	;
	if v32 <= v31 {
		v10 = v31
		v11 = v32
		goto L5
	} else {
		goto L13
	}
L9:
	;
	v31 = v10
	v32 = v16 + int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[592])))
	if base.Ui32(v27) <= base.Ui32(l0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v31 = v16 - int32(1)
	v32 = v11
	goto L8
L13:
	;
	goto L1
L14:
	;
	return int32(1)
L15:
	;
	v60 = base.I32_div_s(v54+v55, int32(2))
	v62 = v60 << (uint(int32(3)) % 32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+uint32(_consts[603])))
	if base.Ui32(v65) < base.Ui32(l0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v83 = int32(655)
	v84 = int32(0)
	goto L25
L17:
	;
	if v78 <= v77 {
		v54 = v77
		v55 = v78
		goto L15
	} else {
		goto L24
	}
L18:
	;
	v77 = v54
	v78 = v60 + int32(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v62)+uint32(_consts[604])))
	if base.Ui32(v71) <= base.Ui32(l0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(1)
L22:
	;
	goto L23
L23:
	;
	v77 = v60 - int32(1)
	v78 = v55
	goto L17
L24:
	;
	goto L16
L25:
	;
	v89 = base.I32_div_s(v83+v84, int32(2))
	v91 = v89 << (uint(int32(3)) % 32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[605])))
	if base.Ui32(v94) < base.Ui32(l0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	return int32(0)
L27:
	;
	if v107 <= v106 {
		v83 = v106
		v84 = v107
		goto L25
	} else {
		goto L34
	}
L28:
	;
	v106 = v83
	v107 = v89 + int32(1)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_consts[606])))
	if base.Ui32(v100) <= base.Ui32(l0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(1)
L32:
	;
	goto L33
L33:
	;
	v106 = v89 - int32(1)
	v107 = v84
	goto L27
L34:
	;
	goto L26
}
