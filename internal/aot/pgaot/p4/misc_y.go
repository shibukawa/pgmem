package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__yconv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(100)
	v13 = base.I32_div_s(l1, v12)
	v15 = base.I32_div_s(l0, v12)
	v23 = l1 - v13*v12 + (l0 - v15*v12)
	v26 = base.I32_div_s(base.I32_extend16_s(v23), v12)
	v28 = v13 + v15 + base.I32_extend16_s(v26)
	v34 = base.I32_extend16_s(v23 - v26*v12)
	if base.B2i32(v28 <= v5)|base.B2i32(v5 <= v34) == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v55 = int32(0)
	if v53|base.B2i32(v55 <= v54) == v55 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v53 = v28 - int32(1)
	v54 = v34 + int32(100)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v44 = int32(0)
	if base.B2i32(v34 <= v44)|base.B2i32(v44 <= v28) != 0 {
		v53 = v28
		v54 = v34
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v53 = v28 + int32(1)
	v54 = v34 - int32(100)
	goto L1
L6:
	;
	v115 = v54 >> (uint(int32(31)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v54 ^ v115 - v115
	v120 = v10 + int32(20)
	v122 = F_pg_sprintf(m, v120, int32(_a_F__yconv_0), v10)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L14
	} else {
		goto L21
	}
L7:
	;
	v109 = l3
	goto L6
L8:
	;
	if base.Ui32(l3) <= base.Ui32(l2) {
		v109 = l2
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v53
	v75 = v10 + int32(20)
	v79 = F_pg_sprintf(m, v75, int32(_a_F__yconv_0), v10+int32(16))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v61 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v61)
	if l2+int32(1) == l3 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v66 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)) = uint8(v66)
	v69 = l2 + int32(2)
	if v69 == l3 {
		v109 = l3
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v71)
	v109 = v69
	goto L6
L14:
	;
	return int32(0)
L15:
	;
	if base.Ui32(l3) <= base.Ui32(l2) {
		v109 = l2
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v84 = v75
	v86 = l2
	goto L17
L17:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v91)
	if v91 == int32(0) {
		v109 = v86
		goto L6
	} else {
		goto L19
	}
L18:
	;
	goto L7
L19:
	;
	v95 = int32(1)
	v98 = v86 + v95
	if v98 != l3 {
		v84 = v84 + v95
		v86 = v98
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	if base.Ui32(l3) <= base.Ui32(v109) {
		v143 = v109
		goto L22
	} else {
		goto L23
	}
L22:
	;
	m.G0 = v10 + int32(32)
	return v143
L23:
	;
	v125 = v120
	v127 = v109
	goto L24
L24:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	*(*uint8)(unsafe.Add(mBase, uint32(v127))) = uint8(v132)
	if v132 == int32(0) {
		v143 = v127
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v143 = l3
	goto L22
L26:
	;
	v136 = int32(1)
	v139 = v127 + v136
	if v139 != l3 {
		v125 = v125 + v136
		v127 = v139
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
}
func F_yy_fatal_error_5(m *base.Module, l0 int32) {
	var v7 int32
	_ = v7
	Fn14046(m, l0, int32(_a_F_yy_fatal_error_5_0), int32(51), int32(_a_F_yy_fatal_error_5_1), int32(_a_F_yy_fatal_error_5_2))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
