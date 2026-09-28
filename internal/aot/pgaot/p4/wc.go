package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_wc_isalpha_libc_mb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	if base.Ui32(l0) <= base.Ui32(int32(_a_F_wc_isalpha_libc_mb_0)) {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_wc_isalpha_libc_mb[0]))))
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v12<<(uint(int32(5))%32))+uint32(_c_F_wc_isalpha_libc_mb[0]))))
		v24 = int32(base.Ui32(v16)>>(uint(l0&int32(7))%32)) & int32(1)
	} else {
		v24 = base.B2i32(base.Ui32(l0) < base.Ui32(int32(_a_F_wc_isalpha_libc_mb_1)))
	}
	return base.B2i32(v24 != int32(0))
}
func F_wc_iscased_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	if base.Ui32(int32(255)) < base.Ui32(l0) {
		return int32(0)
	} else {
		if base.Ui32(l0-int32(65)) < base.Ui32(int32(26)) {
			return int32(1)
		} else {
			return base.B2i32(base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(26))) != int32(0))
		}
	}
}
func F_wc_isdigit_builtin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if (v6^int32(-1))&int32(1) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v63 = base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
	goto L3
L2:
	;
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	return v63
L4:
	;
	v63 = base.B2i32(v53&int32(255) == int32(9))
	goto L3
L5:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_wc_isdigit_builtin[0]))))
	v53 = v47
	goto L4
L6:
	;
	v19 = int32(3408)
	v20 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0<<(uint(int32(1))%32))+uint32(_c_F_wc_isdigit_builtin[1]))))
	v53 = v46
	goto L4
L9:
	;
	v25 = base.I32_div_s(v19+v20, int32(2))
	v27 = v25 * int32(12)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_wc_isdigit_builtin[2])))
	if base.Ui32(v30) < base.Ui32(l0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v53 = int32(0)
	goto L4
L11:
	;
	if v41 <= v40 {
		v19 = v40
		v20 = v41
		goto L9
	} else {
		goto L16
	}
L12:
	;
	v40 = v19
	v41 = v25 + int32(1)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_wc_isdigit_builtin[3])))
	if base.Ui32(v36) <= base.Ui32(l0) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v40 = v25 - int32(1)
	v41 = v20
	goto L11
L16:
	;
	goto L10
}
func F_wc_isgraph_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	var v11 int32
	_ = v11
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		v11 = base.B2i32(base.Ui32(l0-int32(33)) < base.Ui32(int32(94)))
	} else {
		v11 = int32(0)
	}
	return v11
}
func F_wc_isspace_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	var v16 int32
	_ = v16
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		v16 = base.B2i32(base.B2i32(l0 == int32(32))|base.B2i32(base.Ui32(l0-int32(9)) < base.Ui32(int32(5))) != int32(0))
	} else {
		v16 = int32(0)
	}
	return v16
}
