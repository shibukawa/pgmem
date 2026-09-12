package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RunObjectPostAlterHookStr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l2
	v19 = *(*int32)(unsafe.Add(mBase, _consts[289]))
	m.T0[v19].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(2), int32(6243), l0, l1, v7+int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_get_object_attnum_namespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[290]))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64)+24)))
	m.G0 = v7 + int32(16)
	return v67
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == l0 {
		v64 = v10
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[290])) = v61
	v64 = v61
	goto L2
L8:
	;
	v19 = v16 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[291])))
	if l0 != v22 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = v19 + int32(745696)
	goto L7
L10:
	;
	if v16 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v29 = (v16 | int32(1)) * int32(40)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[291])))
	if l0 == v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v29 + int32(745696)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v39 = (v16 | int32(2)) * int32(40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[291])))
	if l0 == v42 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v39 + int32(745696)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v49 = (v16 | int32(3)) * int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[291])))
	if l0 == v52 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = v49 + int32(745696)
	goto L7
L21:
	;
	v16 = v16 + int32(4)
	goto L8
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(59130), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(488742), int32(2777), int32(498863))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_attnum_owner(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[290]))
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L2:
	;
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v64)+26)))
	m.G0 = v7 + int32(16)
	return v67
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == l0 {
		v64 = v10
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(0)
	goto L8
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[290])) = v61
	v64 = v61
	goto L2
L8:
	;
	v19 = v16 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[291])))
	if l0 != v22 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = v19 + int32(745696)
	goto L7
L10:
	;
	if v16 == int32(36) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v29 = (v16 | int32(1)) * int32(40)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+uint32(_consts[291])))
	if l0 == v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v29 + int32(745696)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v39 = (v16 | int32(2)) * int32(40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[291])))
	if l0 == v42 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v39 + int32(745696)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v49 = (v16 | int32(3)) * int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_consts[291])))
	if l0 == v52 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = v49 + int32(745696)
	goto L7
L21:
	;
	v16 = v16 + int32(4)
	goto L8
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	F_errmsg_internal(m, int32(59130), v7)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(488742), int32(2777), int32(498863))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	if v4 != 0 {
		return int32(0)
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v5 != 0 {
			return int32(0)
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+20))
			v9 = F_cstring_to_text_with_len(m, v6, v7-v6)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
				return int32(0)
			}
		}
	}
}
func F_makeObjectName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
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
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	v4 = int32(0)
	v10 = F_strlen(m, l0)
	mBase = m.M
	if l1 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v16 = v4
	v17 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v15 = F_strlen(m, l1)
	mBase = m.M
	v16 = int32(1)
	v17 = v15
	goto L1
L5:
	;
	v18 = F_strlen(m, l2)
	mBase = m.M
	v22 = v18 + v16 + int32(1)
	goto L7
L6:
	;
	v22 = v16
	goto L7
L7:
	;
	v24 = int32(63) - v22
	if v17+v10 <= v24 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v51 = F_pg_mbcliplen(m, l0, v47, v47)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v45 = v17
	v47 = v10
	goto L8
L10:
	;
	goto L11
L11:
	;
	v30 = v17
	v31 = v10
	goto L12
L12:
	;
	v38 = v31 - base.B2i32(v30 < v31)
	v39 = v30 - base.B2i32(v31 <= v30)
	if v24 < v38+v39 {
		v30 = v39
		v31 = v38
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v45 = v39
	v47 = v38
	goto L8
L14:
	;
	goto L13
L15:
	;
	return int32(0)
L16:
	;
	if l1 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v55 = F_pg_mbcliplen(m, l1, v45, v45)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	v57 = v45
	goto L19
L19:
	;
	v62 = F_palloc(m, v51+v22+v57+int32(1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v57 = v55
	goto L19
L21:
	;
	if v51 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if l1 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v64 = F__emscripten_memcpy_bulkmem(m, v62, l0, v51)
	mBase = m.M
	v65 = v64
	goto L25
L24:
	;
	v65 = v62
	goto L25
L25:
	;
	goto L22
L26:
	;
	v67 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v51+v65))) = uint8(v67)
	v70 = v51 + int32(1)
	if v57 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v76 = v51
	goto L28
L28:
	;
	v77 = v76 + v65
	if l2 != 0 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v76 = v57 + v70
	goto L28
L30:
	;
	v72 = F__emscripten_memcpy_bulkmem(m, v65+v70, l1, v57)
	mBase = m.M
	goto L32
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	v78 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v78)
	v81 = v77 + int32(1)
	if (l2^v81)&int32(3) != 0 {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	goto L35
L35:
	;
	v157 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v157)
	return v65
L36:
	;
	return v65
L37:
	;
	goto L36
L38:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v135)
	if v135&int32(255) == int32(0) {
		goto L37
	} else {
		goto L53
	}
L39:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v134 = l2
	v135 = v87
	v136 = v81
	goto L38
L40:
	;
	goto L41
L41:
	;
	if l2&int32(3) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v91 = l2
	v93 = v81
	goto L45
L43:
	;
	v105 = l2
	v107 = v81
	goto L44
L44:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v112 = int32(-2139062144)
	if (int32(16843008)-v109|v109)&v112 != v112 {
		v134 = v105
		v135 = v109
		v136 = v107
		goto L38
	} else {
		goto L49
	}
L45:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v94)
	if v94 == int32(0) {
		goto L37
	} else {
		goto L47
	}
L46:
	;
	v105 = v101
	v107 = v99
	goto L44
L47:
	;
	v98 = int32(1)
	v99 = v93 + v98
	v101 = v91 + v98
	if v101&int32(3) != 0 {
		v91 = v101
		v93 = v99
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v117 = v105
	v118 = v109
	v119 = v107
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v118
	v121 = int32(4)
	v122 = v119 + v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v125 = v117 + v121
	v129 = int32(-2139062144)
	if (v123|(int32(16843008)-v123))&v129 == v129 {
		v117 = v125
		v118 = v123
		v119 = v122
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v134 = v125
	v135 = v123
	v136 = v122
	goto L38
L52:
	;
	goto L51
L53:
	;
	v143 = v134
	v145 = v136
	goto L54
L54:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v145)+1)) = uint8(v146)
	v148 = int32(1)
	if v146 != 0 {
		v143 = v143 + v148
		v145 = v145 + v148
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L37
L56:
	;
	goto L55
}
func F_object_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_object_aclmask_ext(m, l0, l1, l2, l3, int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int64(0))
	}
}
func F_object_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_object_aclmask_ext(m, l0, l1, l2, l3, l4)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int64(0))
	}
}
func F_object_address_present(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9 = v7 - int32(1)
	if v9 < v3 {
		v44 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v44
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = v9
	goto L3
L3:
	;
	v22 = v13 + v16*int32(12)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v12 != v23 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v44 = v38
	goto L1
L5:
	;
	v38 = int32(0)
	if v38 < v16 {
		v16 = v16 - int32(1)
		goto L3
	} else {
		goto L10
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v25 != v26 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v28 = int32(1)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v29 == v30 {
		v44 = v28
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v29 == int32(0) {
		v44 = v28
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	goto L4
}
func F_record_object_address_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(2) <= v11 {
		F_pg_qsort(m, v10, v11, int32(12), int32(462))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v19 = int32(1)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if int32(2) <= v20 {
				v27 = v18
				v28 = v19
				v30 = int32(1)
				for {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v37 = v34 + v30*int32(12)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
					if v33 != v38 {
						v50 = v27 + int32(12)
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
						*(*int64)(unsafe.Add(mBase, uint32(v50))) = v51
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v53
						v57 = v50
						v58 = v28 + int32(1)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
						if v40 != v41 {
							v50 = v27 + int32(12)
							v51 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
							*(*int64)(unsafe.Add(mBase, uint32(v50))) = v51
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v53
							v57 = v50
							v58 = v28 + int32(1)
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							if v43 == v44 {
								v57 = v27
								v58 = v28
							} else {
								if v43 != 0 {
									v50 = v27 + int32(12)
									v51 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
									*(*int64)(unsafe.Add(mBase, uint32(v50))) = v51
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v53
									v57 = v50
									v58 = v28 + int32(1)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v44
									v57 = v27
									v58 = v28
								}
							}
						}
					}
					v62 = v30 + int32(1)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if v62 < v63 {
						v27 = v57
						v28 = v58
						v30 = v62
						continue
					} else {
						break
					}
					break
				}
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v69 = v65
				v70 = v58
			} else {
				v69 = v18
				v70 = v19
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v70
			v79 = v69
			v80 = v70
			F_recordMultipleDependencies(m, l0, v79, v80, l2)
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v79 = v10
		v80 = v11
		F_recordMultipleDependencies(m, l0, v79, v80, l2)
		mBase = m.M
		v86 = m.ExcPending
		if v86 != 0 {
			return
		} else {
			return
		}
	}
}
