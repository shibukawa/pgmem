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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v14 - int32(26) {
	case 0:
		v42 = v5
		goto L2
	default:
		goto L3
	case 3:
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(80)
	return v157
L2:
	;
	v46 = F_jspGetNext(m, l2, v9+int32(52))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L13
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = m.T0[v39].(func(*base.Module, int32, int32) int32)(m, v9+int32(16), l2)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L10
	}
L4:
	;
	F_jspGetArg(m, l2, v9+int32(24))
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
	v30 = F_extract_jsp_bool_expr(m, l0, v9+int32(12), v9+int32(24), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v30 == int32(0) {
		v42 = v5
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v35 = F_lappend(m, int32(0), v30)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v42 = v35
	goto L2
L10:
	;
	if v40 != 0 {
		v42 = v5
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v157 = v5
	goto L1
L12:
	;
	if v107 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L13:
	;
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v49 = v42
	goto L17
L15:
	;
	v94 = v42
	goto L16
L16:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v100
	v104 = m.T0[v99].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v9+int32(4), l3, v94)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L30
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	switch v54 - int32(26) {
	case 0:
		v85 = v49
		goto L19
	default:
		goto L20
	case 3:
		goto L21
	}
L18:
	;
	v94 = v85
	goto L16
L19:
	;
	v88 = v9 + int32(52)
	v91 = F_jspGetNext(m, v88, v88)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L28
	}
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = m.T0[v80].(func(*base.Module, int32, int32) int32)(m, v9+int32(16), v9+int32(52))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L26
	}
L21:
	;
	F_jspGetArg(m, v9+int32(52), v9+int32(24))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v63
	v70 = F_extract_jsp_bool_expr(m, l0, v9+int32(8), v9+int32(24), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	if v70 == int32(0) {
		v85 = v49
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v74 = F_lappend(m, v49, v70)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v85 = v74
	goto L19
L26:
	;
	if v81 == int32(0) {
		v107 = v49
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v85 = v49
	goto L19
L28:
	;
	if v91 != 0 {
		v49 = v85
		goto L17
	} else {
		goto L29
	}
L29:
	;
	goto L18
L30:
	;
	v107 = v104
	goto L12
L31:
	;
	v157 = int32(0)
	goto L1
L32:
	;
	goto L33
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v115 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v157 = v119
	goto L1
L35:
	;
	goto L36
L36:
	;
	v124 = F_palloc(m, v115<<(uint(int32(2))%32)+int32(8))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v129 <= int32(0) {
		v157 = v124
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v135 = int32(0)
	goto L39
L39:
	;
	v142 = v135 << (uint(int32(2)) % 32)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v144+v142)))
	*(*int32)(unsafe.Add(mBase, uint32(v124+int32(8)+v142))) = v146
	v149 = v135 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v149 < v150 {
		v135 = v149
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v157 = v124
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
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = l0 - int32(4)
	if base.Ui32(int32(50)) <= base.Ui32(v9) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
			F_errmsg_internal(m, int32(485419), v6)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(497981), int32(918), int32(382207))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		if base.I32_wrap_i64(int64(base.Ui64(int64(1125796693540851))>>(uint(base.I64_extend_i32_u(v9))%64)))&int32(1) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(485419), v6)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(497981), int32(918), int32(382207))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v9<<(uint(int32(2))%32))+uint32(_consts[1090])))
			m.G0 = v6 + int32(16)
			return v24
		}
	}
}
