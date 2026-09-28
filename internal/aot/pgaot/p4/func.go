package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LookupFuncNameInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v133 int32
	_ = v133
	v8 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v8
	v21 = F_FuncnameGetCandidates(m, l1, l2, v8, v8, v8, l4, l5, v12+int32(12))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v133
L2:
	;
	return int32(0)
L3:
	;
	if v21 == int32(0) {
		v133 = v8
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = l2 << (uint(int32(2)) % 32)
	v38 = v21
	v40 = v8
	goto L5
L5:
	;
	if base.B2i32(l2 <= int32(0)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(1)
	v133 = int32(0)
	goto L1
L7:
	;
	goto L6
L8:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v122 != 0 {
		v38 = v122
		v40 = v121
		goto L5
	} else {
		goto L40
	}
L9:
	;
	v45 = v38 + int32(32)
	if base.Ui32(int32(4)) <= base.Ui32(v28) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	goto L11
L11:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v108 == int32(0) {
		goto L7
	} else {
		goto L31
	}
L12:
	;
	if v107 != 0 {
		v121 = v40
		goto L8
	} else {
		goto L30
	}
L13:
	;
	v107 = int32(0)
	goto L12
L14:
	;
	v81 = v76
	v82 = v77
	v83 = v78
	goto L24
L15:
	;
	if (l3|v45)&int32(3) != 0 {
		v76 = l3
		v77 = v45
		v78 = v28
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v69 = l3
	v70 = v45
	v71 = v28
	goto L17
L17:
	;
	if v71 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v53 = l3
	v54 = v45
	v55 = v28
	goto L19
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v58 != v59 {
		v76 = v53
		v77 = v54
		v78 = v55
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v69 = v64
	v70 = v62
	v71 = v66
	goto L17
L21:
	;
	v61 = int32(4)
	v62 = v54 + v61
	v64 = v53 + v61
	v66 = v55 - v61
	if base.Ui32(int32(3)) < base.Ui32(v66) {
		v53 = v64
		v54 = v62
		v55 = v66
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v76 = v69
	v77 = v70
	v78 = v71
	goto L14
L24:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v86 == v87 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v107 = v86 - v87
	goto L12
L26:
	;
	v89 = int32(1)
	v94 = v83 - v89
	if v94 != 0 {
		v81 = v81 + v89
		v82 = v82 + v89
		v83 = v94
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L13
L30:
	;
	goto L11
L31:
	;
	switch l0 - int32(1) {
	case 0, 18:
		goto L34
	default:
		goto L32
	case 28:
		goto L33
	}
L32:
	;
	if v40 != 0 {
		goto L7
	} else {
		goto L39
	}
L33:
	;
	v115 = F_get_func_prokind(m, v108)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L37
	}
L34:
	;
	v111 = F_get_func_prokind(m, v108)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	if v111 != int32(112) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v121 = v40
	goto L8
L37:
	;
	if v115 != int32(112) {
		v121 = v40
		goto L8
	} else {
		goto L38
	}
L38:
	;
	goto L32
L39:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v121 = v119
	goto L8
L40:
	;
	v133 = v121
	goto L1
}
func F_func_strict(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_func_strict_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_func_strict_1), int32(2080), int32(_a_F_func_strict_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v30)+99)))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v32
			}
		}
	}
}
