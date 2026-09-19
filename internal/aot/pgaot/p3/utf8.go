package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_skip_utf8(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	if l3 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	goto L3
L3:
	;
	if l3 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v11 = l1
	v13 = l3
	goto L7
L5:
	;
	v48 = l1
	goto L6
L6:
	;
	return v48
L7:
	;
	if l2 <= v11 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v48 = v42
	goto L6
L9:
	;
	return int32(-1)
L10:
	;
	goto L11
L11:
	;
	v19 = v11 + int32(1)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v11))))
	if base.Ui32(v21) < base.Ui32(int32(192)) {
		v42 = v19
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = int32(1)
	if v43 < v13 {
		v11 = v42
		v13 = v13 - v43
		goto L7
	} else {
		goto L19
	}
L13:
	;
	if l2 <= v19 {
		v42 = v19
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v28 = v19
	goto L15
L15:
	;
	v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0+v28))))
	if int32(-65) < v31 {
		v42 = v28
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v42 = l2
	goto L12
L17:
	;
	v35 = v28 + int32(1)
	if v35 != l2 {
		v28 = v35
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L8
}
func F_utf8_to_big5(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14015(m, l0, int32(36), v3, v3, v3, int32(_a_F_utf8_to_big5_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_utf8_to_euc_jis_2004(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14015(m, l0, int32(5), int32(0), int32(25), int32(_a_F_utf8_to_euc_jis_2004_0), int32(_a_F_utf8_to_euc_jis_2004_1))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_utf8_to_euc_kr(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14015(m, l0, int32(3), v3, v3, v3, int32(_a_F_utf8_to_euc_kr_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_utf8_to_iso8859_1(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn14016(m, l0, int32(_a_F_utf8_to_iso8859_1_0), int32(_a_F_utf8_to_iso8859_1_1), int32(169), int32(_a_F_utf8_to_iso8859_1_2), int32(_a_F_utf8_to_iso8859_1_3), int32(_a_F_utf8_to_iso8859_1_4), int32(19), int32(9))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_utf8_to_koi8r(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14015(m, l0, int32(22), v3, v3, v3, int32(_a_F_utf8_to_koi8r_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
