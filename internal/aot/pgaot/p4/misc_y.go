package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__yconv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v32 int32
	_ = v32
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(100)
	v13 = base.I32_div_s(l1, v12)
	v15 = base.I32_div_s(l0, v12)
	v23 = l1 - v13*v12 + (l0 - v15*v12)
	v26 = base.I32_div_s(base.I32_extend16_s(v23), v12)
	v28 = v13 + v15 + base.I32_extend16_s(v26)
	v32 = base.I32_extend16_s(v23 - v26*v12)
	if int32(0) <= v32 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v49 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	if v32 <= int32(0) {
		v49 = v28
		v50 = v32
		goto L1
	} else {
		goto L5
	}
L3:
	;
	if v28 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v49 = v28 - int32(1)
	v50 = v32 + int32(100)
	goto L1
L5:
	;
	if int32(0) <= v28 {
		v49 = v28
		v50 = v32
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v49 = v28 + int32(1)
	v50 = v32 - int32(100)
	goto L1
L7:
	;
	v110 = v50 >> (uint(int32(31)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v50 ^ v110 - v110
	v117 = F_pg_sprintf(m, v10+int32(20), int32(466076), v10)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L15
	} else {
		goto L22
	}
L8:
	;
	v104 = l3
	goto L7
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v49
	v72 = F_pg_sprintf(m, v10+int32(20), int32(466076), v10+int32(16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	if int32(0) <= v50 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if base.Ui32(l3) <= base.Ui32(l2) {
		v104 = l2
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v54 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v54)
	v57 = l2 + int32(1)
	if v57 == l3 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v59 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v59)
	v62 = l2 + int32(2)
	if v62 == l3 {
		v104 = l3
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v64)
	v104 = v62
	goto L7
L15:
	;
	return int32(0)
L16:
	;
	if base.Ui32(l3) <= base.Ui32(l2) {
		v104 = l2
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v79 = v10 + int32(20)
	v81 = l2
	goto L18
L18:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v86)
	if v86 == int32(0) {
		v104 = v81
		goto L7
	} else {
		goto L20
	}
L19:
	;
	goto L8
L20:
	;
	v90 = int32(1)
	v93 = v81 + v90
	if v93 != l3 {
		v79 = v79 + v90
		v81 = v93
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if base.Ui32(l3) <= base.Ui32(v104) {
		v140 = v104
		goto L23
	} else {
		goto L24
	}
L23:
	;
	m.G0 = v10 + int32(32)
	return v140
L24:
	;
	v122 = v10 + int32(20)
	v124 = v104
	goto L25
L25:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v129)
	if v129 == int32(0) {
		v140 = v124
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v140 = l3
	goto L23
L27:
	;
	v133 = int32(1)
	v136 = v124 + v133
	if v136 != l3 {
		v122 = v122 + v133
		v124 = v136
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
}
func F_yy_fatal_error_5(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		F_errmsg_internal(m, int32(206200), v5)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errfinish(m, int32(314771), int32(51), int32(80852))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
