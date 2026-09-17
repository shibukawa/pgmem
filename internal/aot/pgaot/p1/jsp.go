package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_extract_jsp_path_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v14 - int32(26) {
	case 0:
		v40 = v5
		goto L2
	default:
		goto L3
	case 3:
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(80)
	return v151
L2:
	;
	v44 = F_jspGetNext(m, l2, v9+int32(52))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L13
	}
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = m.T0[v37].(func(*base.Module, int32, int32) int32)(m, v9+int32(16), l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L10
	}
L4:
	;
	v18 = v9 + int32(24)
	F_jspGetArg(m, l2, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v23
	v28 = F_extract_jsp_bool_expr(m, l0, v9+int32(12), v18, int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v28 == int32(0) {
		v40 = v5
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v33 = F_lappend(m, int32(0), v28)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v40 = v33
	goto L2
L10:
	;
	if v38 != 0 {
		v40 = v5
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v151 = v5
	goto L1
L12:
	;
	if v101 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L13:
	;
	if v44 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = v40
	goto L17
L15:
	;
	v88 = v40
	goto L16
L16:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v94
	v98 = m.T0[v93].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v9+int32(4), l3, v88)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L30
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	switch v52 - int32(26) {
	case 0:
		v81 = v47
		goto L19
	default:
		goto L20
	case 3:
		goto L21
	}
L18:
	;
	v88 = v81
	goto L16
L19:
	;
	v84 = v9 + int32(52)
	v85 = F_jspGetNext(m, v84, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L28
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v77 = m.T0[v76].(func(*base.Module, int32, int32) int32)(m, v9+int32(16), v9+int32(52))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L26
	}
L21:
	;
	v58 = v9 + int32(24)
	F_jspGetArg(m, v9+int32(52), v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v61
	v66 = F_extract_jsp_bool_expr(m, l0, v9+int32(8), v58, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if v66 == int32(0) {
		v81 = v47
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v70 = F_lappend(m, v47, v66)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v81 = v70
	goto L19
L26:
	;
	if v77 == int32(0) {
		v101 = v47
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v81 = v47
	goto L19
L28:
	;
	if v85 != 0 {
		v47 = v81
		goto L17
	} else {
		goto L29
	}
L29:
	;
	goto L18
L30:
	;
	v101 = v98
	goto L12
L31:
	;
	v151 = int32(0)
	goto L1
L32:
	;
	goto L33
L33:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v109 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v151 = v113
	goto L1
L35:
	;
	goto L36
L36:
	;
	v118 = F_palloc(m, v109<<(uint(int32(2))%32)+int32(8))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v123 <= int32(0) {
		v151 = v118
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v129 = int32(0)
	goto L39
L39:
	;
	v136 = v129 << (uint(int32(2)) % 32)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v138+v136)))
	*(*int32)(unsafe.Add(mBase, uint32(v118+int32(8)+v136))) = v140
	v143 = v129 + int32(1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v143 < v144 {
		v129 = v143
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v151 = v118
	goto L1
L41:
	;
	goto L40
}
func F_jspInit(m *base.Module, l0 int32, l1 int32) {
	var v7 int32
	_ = v7
	F_jspInitByBuffer(m, l0, l1+int32(8), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_jspOperationName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = l0 - int32(4)
	if base.B2i32(base.Ui32(int32(50)) <= base.Ui32(v9))|base.B2i32(base.I32_wrap_i64(int64(base.Ui64(int64(1125796693540851))>>(uint(base.I64_extend_i32_u(v9))%64)))&int32(1) == v2) == v2 {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v9<<(uint(int32(2))%32))+uint32(_c_F_jspOperationName[0])))
		m.G0 = v6 + int32(16)
		return v25
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
			F_errmsg_internal(m, int32(_a_F_jspOperationName_0), v6)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_jspOperationName_1), int32(918), int32(_a_F_jspOperationName_2))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
