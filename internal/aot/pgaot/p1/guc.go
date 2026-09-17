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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_validate_option_array_item(m, l1, v3, v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = F_find_option(m, l1, int32(0), int32(1), int32(19))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v23 = v22
	goto L6
L5:
	;
	v23 = l1
	goto L6
L6:
	;
	if l0 == int32(0) {
		v132 = v3
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v9 + int32(16)
	return v132
L8:
	;
	v26 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v26
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v30 <= int32(0) {
		v132 = v3
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v37 = v3
	goto L10
L10:
	;
	v43 = F_array_ref(m, l0, v9+int32(12), v9+int32(3))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v132 = v120
	goto L7
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v43
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+3)))
	if v46 != 0 {
		v120 = v37
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v124 = v122 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v124 <= v126 {
		v37 = v120
		goto L10
	} else {
		goto L39
	}
L14:
	;
	v47 = F_text_to_cstring(m, v43)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v49 = F_strlen(m, v23)
	mBase = m.M
	if v49 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v94 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L17:
	;
	v94 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v55 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v56 = v47
	v57 = v23
	v58 = v49
	v59 = v55
	goto L24
L21:
	;
	v82 = v23
	v86 = int32(0)
	goto L22
L22:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v94 = v86 - v87
	goto L16
L23:
	;
	v82 = v77
	v86 = v79
	goto L22
L24:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if base.B2i32(v59 != v61)|base.B2i32(v61 == int32(0)) != 0 {
		v77 = v57
		v79 = v59
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v77 = v71
	v79 = int32(0)
	goto L23
L26:
	;
	v67 = v58 - int32(1)
	if v67 == int32(0) {
		v77 = v57
		v79 = v59
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v70 = int32(1)
	v71 = v57 + v70
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v72 != 0 {
		v56 = v56 + v70
		v57 = v71
		v58 = v67
		v59 = v72
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v49))))
	if v98 == int32(61) {
		v120 = v37
		goto L13
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v37 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L31
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v115 + int32(1)
	v120 = v114
	goto L13
L34:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v106 = F_array_set(m, v37, v9+int32(8), v103, int32(-1), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v112 = F_construct_array_builtin(m, v9+int32(4), int32(1), int32(25))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v114 = v106
	goto L33
L38:
	;
	v114 = v112
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
