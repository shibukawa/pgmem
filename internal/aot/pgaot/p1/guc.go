package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GUCArrayDelete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int64
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v14 = F_validate_option_array_item(m, l1, v3, v3)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = F_find_option(m, l1, int32(0), int32(1), int32(19))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v24 = v23
	goto L6
L5:
	;
	v24 = l1
	goto L6
L6:
	;
	if l0 == int32(0) {
		v135 = v3
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v10 + int32(32)
	return v135
L8:
	;
	v27 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v27
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v31 <= int32(0) {
		v135 = v3
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v38 = v3
	goto L10
L10:
	;
	v45 = F_array_ref(m, l0, v10+int32(24), v10+int32(15))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v135 = v123
	goto L7
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v45
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	if v48 != 0 {
		v123 = v38
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v127 = v125 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v127 <= v129 {
		v38 = v123
		goto L10
	} else {
		goto L39
	}
L14:
	;
	v50 = F_text_to_cstring(m, base.I32_wrap_i64(v45))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v52 = F_strlen(m, v24)
	mBase = m.M
	if v52 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v97 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L17:
	;
	v97 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v58 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v59 = v50
	v60 = v24
	v61 = v52
	v62 = v58
	goto L24
L21:
	;
	v85 = v24
	v89 = int32(0)
	goto L22
L22:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v97 = v89 - v90
	goto L16
L23:
	;
	v85 = v80
	v89 = v82
	goto L22
L24:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if base.B2i32(v62 != v64)|base.B2i32(v64 == int32(0)) != 0 {
		v80 = v60
		v82 = v62
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v80 = v74
	v82 = int32(0)
	goto L23
L26:
	;
	v70 = v61 - int32(1)
	if v70 == int32(0) {
		v80 = v60
		v82 = v62
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v73 = int32(1)
	v74 = v60 + v73
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	if v75 != 0 {
		v59 = v59 + v73
		v60 = v74
		v61 = v70
		v62 = v75
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v52))))
	if v101 == int32(61) {
		v123 = v38
		goto L13
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v38 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L31
L33:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v118 + int32(1)
	v123 = v117
	goto L13
L34:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
	v109 = F_array_set(m, v38, v10+int32(28), v106, int32(-1), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v115 = F_construct_array_builtin(m, v10+int32(16), int32(1), int32(25))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v117 = v109
	goto L33
L38:
	;
	v117 = v115
	goto L33
L39:
	;
	goto L11
}
func F_GUC_check_errcode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_GUC_check_errcode[0])) = l0
	return
}
func F_InitializeGUCOptions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v9 = F_pg_tzset(m, int32(_a_F_InitializeGUCOptions_0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeGUCOptions[0])) = v9
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeGUCOptions[1])) = v9
	F_build_guc_variables(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = v5 + int32(12)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeGUCOptions[2]))
	F_hash_seq_init(m, v17, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = F_hash_seq_search(m, v17)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v24 = v22
	goto L9
L7:
	;
	goto L8
L8:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitializeGUCOptions[3])) = uint8(v36)
	v41 = int32(1)
	v42 = int32(10)
	v48 = F_set_config_with_handle(m, int32(_a_F_InitializeGUCOptions_1), v36, int32(_a_F_InitializeGUCOptions_2), v41, v42, v42, v36, v41, v36, v36)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	F_InitializeOneGUCOption(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v31 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v31 != 0 {
		v24 = v31
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v51 = int32(0)
	v53 = int32(1)
	v54 = int32(10)
	v60 = F_set_config_with_handle(m, int32(_a_F_InitializeGUCOptions_3), v51, int32(_a_F_InitializeGUCOptions_4), v53, v54, v54, v51, v53, v51, v51)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v63 = int32(0)
	v65 = int32(1)
	v66 = int32(10)
	v72 = F_set_config_with_handle(m, int32(_a_F_InitializeGUCOptions_5), v63, int32(_a_F_InitializeGUCOptions_4), v65, v66, v66, v63, v65, v63, v63)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_InitializeGUCOptionsFromEnvironment(m)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	m.G0 = v5 + int32(32)
	return
}
func F_NewGUCNestLevel(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = int32(_a_F_NewGUCNestLevel_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_NewGUCNestLevel[0]))
	v6 = v4 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_NewGUCNestLevel[0])) = v6
	return v6
}
func F_guc_name_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	v5 = l0
	v6 = l1
	goto L1
L1:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return base.I32_extend8_s(v29) - base.I32_extend8_s(v38)
L3:
	;
	v17 = int32(1)
	if base.Ui32((v10-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	if v9 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v9 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(1)
L8:
	;
	v15 = int32(-1)
	goto L10
L9:
	;
	v15 = int32(0)
	goto L10
L10:
	;
	return v15
L11:
	;
	v29 = v10 | int32(32)
	goto L13
L12:
	;
	v29 = v10
	goto L13
L13:
	;
	if base.Ui32((v9-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = v9 | int32(32)
	goto L16
L15:
	;
	v38 = v9
	goto L16
L16:
	;
	if v29 == v38&int32(255) {
		v5 = v5 + v17
		v6 = v6 + v17
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L2
}
func F_guc_name_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = v6
	v9 = v5
	goto L1
L1:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return base.I32_extend8_s(v31) - base.I32_extend8_s(v40)
L3:
	;
	v19 = int32(1)
	if base.Ui32((v12-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	if v11 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v11 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(1)
L8:
	;
	v17 = int32(-1)
	goto L10
L9:
	;
	v17 = int32(0)
	goto L10
L10:
	;
	return v17
L11:
	;
	v31 = v12 | int32(32)
	goto L13
L12:
	;
	v31 = v12
	goto L13
L13:
	;
	if base.Ui32((v11-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = v11 | int32(32)
	goto L16
L15:
	;
	v40 = v11
	goto L16
L16:
	;
	if v31 == v40&int32(255) {
		v8 = v8 + v19
		v9 = v9 + v19
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L2
}
func F_guc_restore_error_context_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v14 int32
	_ = v14
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	if l0 != 0 {
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v5))) = v10
			F_errcontext_msg(m, int32(_a_F_guc_restore_error_context_callback_0), v5)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_guc_var_compare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v10 = v6
	v12 = v8
	goto L1
L1:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return base.I32_extend8_s(v33) - base.I32_extend8_s(v42)
L3:
	;
	v21 = int32(1)
	if base.Ui32((v14-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	if v13 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v13 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(1)
L8:
	;
	v19 = int32(-1)
	goto L10
L9:
	;
	v19 = int32(0)
	goto L10
L10:
	;
	return v19
L11:
	;
	v33 = v14 | int32(32)
	goto L13
L12:
	;
	v33 = v14
	goto L13
L13:
	;
	if base.Ui32((v13-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = v13 | int32(32)
	goto L16
L15:
	;
	v42 = v13
	goto L16
L16:
	;
	if v33 == v42&int32(255) {
		v10 = v10 + v21
		v12 = v12 + v21
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L2
}
