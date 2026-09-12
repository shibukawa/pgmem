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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	v8 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v8
	v14 = F_FuncnameGetCandidates(m, l1, l2, v8, v8, v8, l4, l5)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v125
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		v125 = v8
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = l2 << (uint(int32(2)) % 32)
	v31 = v14
	v33 = v8
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
	v125 = int32(0)
	goto L1
L7:
	;
	goto L6
L8:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v114 != 0 {
		v31 = v114
		v33 = v113
		goto L5
	} else {
		goto L40
	}
L9:
	;
	v37 = v31 + int32(32)
	if base.Ui32(int32(4)) <= base.Ui32(v21) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	goto L11
L11:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v100 == int32(0) {
		goto L7
	} else {
		goto L31
	}
L12:
	;
	if v99 != 0 {
		v113 = v33
		goto L8
	} else {
		goto L30
	}
L13:
	;
	v99 = int32(0)
	goto L12
L14:
	;
	v73 = v68
	v74 = v69
	v75 = v70
	goto L24
L15:
	;
	if (l3|v37)&int32(3) != 0 {
		v68 = l3
		v69 = v37
		v70 = v21
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v61 = l3
	v62 = v37
	v63 = v21
	goto L17
L17:
	;
	if v63 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v45 = l3
	v46 = v37
	v47 = v21
	goto L19
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v50 != v51 {
		v68 = v45
		v69 = v46
		v70 = v47
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v61 = v56
	v62 = v54
	v63 = v58
	goto L17
L21:
	;
	v53 = int32(4)
	v54 = v46 + v53
	v56 = v45 + v53
	v58 = v47 - v53
	if base.Ui32(int32(3)) < base.Ui32(v58) {
		v45 = v56
		v46 = v54
		v47 = v58
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v68 = v61
	v69 = v62
	v70 = v63
	goto L14
L24:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v78 == v79 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v99 = v78 - v79
	goto L12
L26:
	;
	v81 = int32(1)
	v86 = v75 - v81
	if v86 != 0 {
		v73 = v73 + v81
		v74 = v74 + v81
		v75 = v86
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
	if v33 != 0 {
		goto L7
	} else {
		goto L39
	}
L33:
	;
	v107 = F_get_func_prokind(m, v100)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L37
	}
L34:
	;
	v103 = F_get_func_prokind(m, v100)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	if v103 != int32(112) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v113 = v33
	goto L8
L37:
	;
	if v107 != int32(112) {
		v113 = v33
		goto L8
	} else {
		goto L38
	}
L38:
	;
	goto L32
L39:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v113 = v111
	goto L8
L40:
	;
	v125 = v113
	goto L1
}
func F_func_strict(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(42736), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(479899), int32(1908), int32(102667))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v29)+99)))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
