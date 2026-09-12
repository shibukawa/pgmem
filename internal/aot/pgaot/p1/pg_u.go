package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_u_isalnum(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	if base.Ui32(int32(128)) <= base.Ui32(l0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v111
L2:
	;
	v111 = base.B2i32(v100&int32(255) == int32(9))
	goto L1
L3:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[644]))))
	v100 = v93
	goto L2
L4:
	;
	return base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
L5:
	;
	v12 = int32(1178)
	v13 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v68 = int32(1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(v68)%32))+uint32(_consts[637]))))
	if v73&v68 != 0 {
		v111 = v68
		goto L1
	} else {
		goto L28
	}
L8:
	;
	v18 = base.I32_div_s(v12+v13, int32(2))
	v20 = v18 << (uint(int32(3)) % 32)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[638])))
	if base.Ui32(v23) < base.Ui32(l0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if l1 != 0 {
		goto L4
	} else {
		goto L18
	}
L10:
	;
	if v36 <= v35 {
		v12 = v35
		v13 = v36
		goto L8
	} else {
		goto L17
	}
L11:
	;
	v35 = v12
	v36 = v18 + int32(1)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+uint32(_consts[639])))
	if base.Ui32(v29) <= base.Ui32(l0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(1)
L15:
	;
	goto L16
L16:
	;
	v35 = v18 - int32(1)
	v36 = v13
	goto L10
L17:
	;
	goto L9
L18:
	;
	v42 = int32(3367)
	v43 = int32(0)
	goto L20
L19:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+uint32(_consts[643]))))
	v100 = v67
	goto L2
L20:
	;
	v48 = base.I32_div_s(v42+v43, int32(2))
	v50 = v48 * int32(12)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+uint32(_consts[645])))
	if base.Ui32(v53) < base.Ui32(l0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v100 = int32(0)
	goto L2
L22:
	;
	if v64 <= v63 {
		v42 = v63
		v43 = v64
		goto L20
	} else {
		goto L27
	}
L23:
	;
	v63 = v42
	v64 = v48 + int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)+uint32(_consts[646])))
	if base.Ui32(v59) <= base.Ui32(l0) {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v63 = v48 - int32(1)
	v64 = v43
	goto L22
L27:
	;
	goto L21
L28:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L4
}
func F_pg_u_isdigit(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
L2:
	;
	goto L3
L3:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return base.B2i32(v52&int32(255) == int32(9))
L5:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[643]))))
	v52 = v46
	goto L4
L6:
	;
	v16 = int32(0)
	v17 = int32(3367)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_consts[644]))))
	v52 = v45
	goto L4
L9:
	;
	v22 = base.I32_div_s(v16+v17, int32(2))
	v24 = v22 * int32(12)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[645])))
	if base.Ui32(v27) < base.Ui32(l0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v52 = int32(0)
	goto L4
L11:
	;
	if v37 <= v38 {
		v16 = v37
		v17 = v38
		goto L9
	} else {
		goto L16
	}
L12:
	;
	v37 = v22 + int32(1)
	v38 = v17
	goto L11
L13:
	;
	goto L14
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_consts[646])))
	if base.Ui32(v33) <= base.Ui32(l0) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v37 = v16
	v38 = v22 - int32(1)
	goto L11
L16:
	;
	goto L10
}
func F_pg_u_islower(m *base.Module, l0 int32) int32 {
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
	v11 = int32(689)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(26)))
L4:
	;
	v16 = base.I32_div_s(v10+v11, int32(2))
	v18 = v16 << (uint(int32(3)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[1260])))
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[1261])))
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
