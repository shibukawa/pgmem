package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_wc_isalpha_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	var v13 int32
	_ = v13
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		v13 = base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26)))
	} else {
		v13 = int32(0)
	}
	return v13
}
func F_wc_isgraph_libc_mb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v39 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v39 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v39 = base.B2i32(v32 != int32(0))
	goto L1
L6:
	;
	v12 = int32(_a_F_wc_isgraph_libc_mb_0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v22 = int32(_a_F_wc_isgraph_libc_mb_0)
	v23 = F_wcslen(m, v22)
	mBase = m.M
	v32 = v23<<(uint(int32(2))%32) + v22
	goto L5
L9:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v15 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v15 != 0 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	if l0 != v15 {
		v12 = v12 + int32(4)
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L13
L15:
	;
	v21 = v12
	goto L17
L16:
	;
	v21 = int32(0)
	goto L17
L17:
	;
	v32 = v21
	goto L5
L18:
	;
	v77 = int32(0)
	goto L20
L19:
	;
	if base.Ui32(l0) <= base.Ui32(int32(254)) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	return base.B2i32(v77 != int32(0))
L21:
	;
	v77 = base.B2i32(v74 != int32(0))
	goto L20
L22:
	;
	v74 = base.B2i32(base.Ui32(int32(32)) < base.Ui32((l0+int32(1))&int32(127)))
	goto L21
L23:
	;
	goto L24
L24:
	;
	v51 = int32(_a_F_wc_isgraph_libc_mb_1)
	v74 = base.B2i32(l0&v51 != v51)&base.B2i32(base.Ui32(l0-int32(_a_F_wc_isgraph_libc_mb_2)) < base.Ui32(int32(_a_F_wc_isgraph_libc_mb_3))) | (base.B2i32(base.Ui32(l0-int32(_a_F_wc_isgraph_libc_mb_4)) < base.Ui32(int32(_a_F_wc_isgraph_libc_mb_5))) | base.B2i32(base.Ui32(l0) < base.Ui32(int32(_a_F_wc_isgraph_libc_mb_6))) | base.B2i32(base.Ui32(l0-int32(_a_F_wc_isgraph_libc_mb_7)) < base.Ui32(int32(_a_F_wc_isgraph_libc_mb_8))))
	goto L21
}
func F_wc_islower_builtin(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v53
L2:
	;
	goto L1
L3:
	;
	v13 = int32(691)
	v18 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	v53 = base.B2i32(base.Ui32(l0-int32(97)) < base.Ui32(int32(26)))
	goto L2
L6:
	;
	v23 = base.I32_div_s(v13+v18, int32(2))
	v25 = v23 << (uint(int32(3)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_wc_islower_builtin[0])))
	if base.Ui32(v27) < base.Ui32(l0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v53 = int32(0)
	goto L2
L8:
	;
	if v38 <= v37 {
		v13 = v37
		v18 = v38
		goto L6
	} else {
		goto L13
	}
L9:
	;
	v37 = v13
	v38 = v23 + int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_wc_islower_builtin[1])))
	if base.Ui32(v33) <= base.Ui32(l0) {
		v53 = int32(1)
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v37 = v23 - int32(1)
	v38 = v18
	goto L8
L13:
	;
	goto L7
}
func F_wc_ispunct_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		if base.Ui32(l0-int32(33)) <= base.Ui32(int32(93)) {
			v10 = F_isalnum(m, l0)
			v12 = v10
		} else {
			v12 = int32(1)
		}
		v18 = base.B2i32(base.B2i32(v12 == int32(0)) != int32(0))
	} else {
		v18 = int32(0)
	}
	return v18
}
func F_wc_isspace_libc_mb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.B2i32(v39 != int32(0))
L2:
	;
	v39 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v39 = base.B2i32(v32 != int32(0))
	goto L1
L6:
	;
	v12 = int32(_a_F_wc_isspace_libc_mb_0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v22 = int32(_a_F_wc_isspace_libc_mb_0)
	v23 = F_wcslen(m, v22)
	mBase = m.M
	v32 = v23<<(uint(int32(2))%32) + v22
	goto L5
L9:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v15 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v15 != 0 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	if l0 != v15 {
		v12 = v12 + int32(4)
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L13
L15:
	;
	v21 = v12
	goto L17
L16:
	;
	v21 = int32(0)
	goto L17
L17:
	;
	v32 = v21
	goto L5
}
func F_wc_isupper_libc_mb(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	v4 = F_towlower(m, l0)
	return base.B2i32(base.B2i32(v4 != l0) != int32(0))
}
func F_wc_isupper_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	var v13 int32
	_ = v13
	if base.Ui32(l0) <= base.Ui32(int32(255)) {
		v13 = base.B2i32(base.B2i32(base.Ui32(l0-int32(65)) < base.Ui32(int32(26))) != int32(0))
	} else {
		v13 = int32(0)
	}
	return v13
}
