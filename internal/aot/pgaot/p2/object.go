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
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_RunObjectPostAlterHookStr[0]))
	m.T0[v19].(func(*base.Module, int32, int32, int32, int32, int32))(m, int32(2), int32(_a_F_RunObjectPostAlterHookStr_0), l0, l1, v7+int32(8))
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_attnum_namespace[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+24)))
	m.G0 = v8 + int32(16)
	return v52
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == l0 {
		v48 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_attnum_namespace[0])) = v45
	v48 = v45
	goto L2
L8:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_namespace_0)
	goto L7
L9:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_namespace_1)
	goto L7
L10:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_namespace_2)
	goto L7
L11:
	;
	v21 = v18 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_namespace[1])))
	if l0 != v22 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_namespace_3)
	goto L7
L13:
	;
	if v18 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_namespace[2])))
	if v28 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_namespace[3])))
	if v30 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_namespace[4])))
	if v32 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v18 = v18 + int32(4)
	goto L11
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_attnum_namespace_4), v8)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_object_attnum_namespace_5), int32(2777), int32(_a_F_get_object_attnum_namespace_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_object_attnum_owner(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_object_attnum_owner[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+26)))
	m.G0 = v8 + int32(16)
	return v52
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v12 == l0 {
		v48 = v11
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = int32(0)
	goto L11
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_object_attnum_owner[0])) = v45
	v48 = v45
	goto L2
L8:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_owner_0)
	goto L7
L9:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_owner_1)
	goto L7
L10:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_owner_2)
	goto L7
L11:
	;
	v21 = v18 * int32(40)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_owner[1])))
	if l0 != v22 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v45 = v21 + int32(_a_F_get_object_attnum_owner_3)
	goto L7
L13:
	;
	if v18 == int32(36) {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L12
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_owner[2])))
	if v28 == l0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_owner[3])))
	if v30 == l0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_get_object_attnum_owner[4])))
	if v32 == l0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v18 = v18 + int32(4)
	goto L11
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_get_object_attnum_owner_4), v8)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_get_object_attnum_owner_5), int32(2777), int32(_a_F_get_object_attnum_owner_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	v9 = F_strlen(m, l0)
	mBase = m.M
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = F_strlen(m, l1)
	mBase = m.M
	v13 = int32(1)
	v14 = v11
	goto L3
L2:
	;
	v13 = int32(0)
	v14 = int32(0)
	goto L3
L3:
	;
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v15 = F_strlen(m, l2)
	mBase = m.M
	v19 = v15 + v13 + int32(1)
	goto L6
L5:
	;
	v19 = v13
	goto L6
L6:
	;
	v21 = int32(63) - v19
	if v21 < v9+v14 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v27 = v9
	v28 = v14
	goto L10
L8:
	;
	v41 = v9
	v42 = v14
	goto L9
L9:
	;
	v46 = F_pg_mbcliplen(m, l0, v41, v41)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v34 = v28 - base.B2i32(v27 <= v28)
	v35 = v27 - base.B2i32(v28 < v27)
	if v21 < v34+v35 {
		v27 = v35
		v28 = v34
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v41 = v35
	v42 = v34
	goto L9
L12:
	;
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	if l1 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v50 = F_pg_mbcliplen(m, l1, v42, v42)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	v52 = v42
	goto L17
L17:
	;
	v57 = F_palloc(m, v46+v19+v52+int32(1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v52 = v50
	goto L17
L19:
	;
	if v46 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	base.MemoryCopy(m, v57, l0, v46)
	goto L22
L21:
	;
	goto L22
L22:
	;
	if l1 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v61 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v46+v57))) = uint8(v61)
	v64 = v46 + int32(1)
	if v52 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v69 = v46
	goto L25
L25:
	;
	v70 = v69 + v57
	if l2 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	base.MemoryCopy(m, v64+v57, l1, v52)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v69 = v64 + v52
	goto L25
L29:
	;
	v71 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v71)
	v74 = v70 + int32(1)
	if (l2^v74)&int32(3) != 0 {
		goto L35
	} else {
		goto L36
	}
L30:
	;
	goto L31
L31:
	;
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v70))) = uint8(v150)
	return v57
L32:
	;
	return v57
L33:
	;
	goto L32
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v128)
	if v128&int32(255) == int32(0) {
		goto L33
	} else {
		goto L49
	}
L35:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v127 = l2
	v128 = v80
	v129 = v74
	goto L34
L36:
	;
	goto L37
L37:
	;
	if l2&int32(3) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v84 = l2
	v86 = v74
	goto L41
L39:
	;
	v98 = l2
	v100 = v74
	goto L40
L40:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v105 = int32(-2139062144)
	if (int32(16843008)-v102|v102)&v105 != v105 {
		v127 = v98
		v128 = v102
		v129 = v100
		goto L34
	} else {
		goto L45
	}
L41:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v87)
	if v87 == int32(0) {
		goto L33
	} else {
		goto L43
	}
L42:
	;
	v98 = v94
	v100 = v92
	goto L40
L43:
	;
	v91 = int32(1)
	v92 = v86 + v91
	v94 = v84 + v91
	if v94&int32(3) != 0 {
		v84 = v94
		v86 = v92
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v110 = v98
	v111 = v102
	v112 = v100
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v111
	v114 = int32(4)
	v115 = v112 + v114
	v117 = v110 + v114
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v122 = int32(-2139062144)
	if (int32(16843008)-v119|v119)&v122 == v122 {
		v110 = v117
		v111 = v119
		v112 = v115
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v127 = v117
	v128 = v119
	v129 = v115
	goto L34
L48:
	;
	goto L47
L49:
	;
	v136 = v127
	v138 = v129
	goto L50
L50:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)) = uint8(v139)
	v141 = int32(1)
	if v139 != 0 {
		v136 = v136 + v141
		v138 = v138 + v141
		goto L50
	} else {
		goto L52
	}
L51:
	;
	goto L33
L52:
	;
	goto L51
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9 = v7 - int32(1)
	if v9 < v3 {
		v45 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v45
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
	v45 = v39
	goto L1
L5:
	;
	v39 = int32(0)
	if v39 < v16 {
		v16 = v16 - int32(1)
		goto L3
	} else {
		goto L9
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
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.B2i32(v29 == v30)|base.B2i32(v29 == int32(0)) != 0 {
		v45 = int32(1)
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
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
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
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
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v49
						v51 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
						*(*int64)(unsafe.Add(mBase, uint32(v27)+12)) = v51
						v57 = v27 + int32(12)
						v58 = v28 + int32(1)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
						if v40 != v41 {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v49
							v51 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
							*(*int64)(unsafe.Add(mBase, uint32(v27)+12)) = v51
							v57 = v27 + int32(12)
							v58 = v28 + int32(1)
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							if v43 == v44 {
								v57 = v27
								v58 = v28
							} else {
								if v43 != 0 {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v49
									v51 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
									*(*int64)(unsafe.Add(mBase, uint32(v27)+12)) = v51
									v57 = v27 + int32(12)
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
