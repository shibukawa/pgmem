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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
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
	var v106 int32
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
	v10 = v8 - int32(16)
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
	m.G0 = v10 + int32(16)
	return v135
L8:
	;
	v27 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v27
	v32 = l0 + int32(16)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v33 <= int32(0) {
		v135 = v3
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = v3
	goto L10
L10:
	;
	v47 = F_array_ref(m, l0, v10+int32(12), v10+int32(3))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v47
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+3)))
	if v50 != 0 {
		v123 = v40
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v127 = v125 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v127 <= v129 {
		v40 = v123
		goto L10
	} else {
		goto L40
	}
L14:
	;
	v51 = F_text_to_cstring(m, v47)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v53 = F_strlen(m, v24)
	mBase = m.M
	if v53 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v97 == int32(0) {
		goto L30
	} else {
		goto L31
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
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v59 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v60 = v51
	v61 = v24
	v62 = v53
	v63 = v59
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
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v63 != v65 {
		v80 = v61
		v82 = v63
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
	if v65 == int32(0) {
		v80 = v61
		v82 = v63
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v70 = v62 - int32(1)
	if v70 == int32(0) {
		v80 = v61
		v82 = v63
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v73 = int32(1)
	v74 = v61 + v73
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	if v75 != 0 {
		v60 = v60 + v73
		v61 = v74
		v62 = v70
		v63 = v75
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v53))))
	if v101 == int32(61) {
		v123 = v40
		goto L13
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v40 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L32
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v118 + int32(1)
	v123 = v117
	goto L13
L35:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v109 = F_array_set(m, v40, v10+int32(8), v106, int32(-1), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v115 = F_construct_array_builtin(m, v10+int32(4), int32(1), int32(25))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v117 = v109
	goto L34
L39:
	;
	v117 = v115
	goto L34
L40:
	;
	goto L11
}
func F_GUC_check_errcode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _consts[427])) = l0
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v9 = F_pg_tzset(m, int32(520075))
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
	*(*int32)(unsafe.Add(mBase, _consts[1206])) = v9
	*(*int32)(unsafe.Add(mBase, _consts[1064])) = v9
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
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1207]))
	F_hash_seq_init(m, v5+int32(12), v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = v24
	goto L9
L7:
	;
	goto L8
L8:
	;
	v38 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1208])) = uint8(v38)
	v43 = int32(1)
	v44 = int32(10)
	v50 = F_set_config_with_handle(m, int32(261967), v38, int32(440671), v43, v44, v44, v38, v43, v38, v38)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	F_InitializeOneGUCOption(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v33 = F_hash_seq_search(m, v5+int32(12))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v33 != 0 {
		v27 = v33
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v53 = int32(0)
	v55 = int32(1)
	v56 = int32(10)
	v62 = F_set_config_with_handle(m, int32(19570), v53, int32(240899), v55, v56, v56, v53, v55, v53, v53)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v65 = int32(0)
	v67 = int32(1)
	v68 = int32(10)
	v74 = F_set_config_with_handle(m, int32(395092), v65, int32(240899), v67, v68, v68, v65, v67, v65, v65)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_InitializeGUCOptionsFromEnvironment(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
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
	v2 = int32(4513288)
	v4 = *(*int32)(unsafe.Add(mBase, _consts[309]))
	v6 = v4 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[309])) = v6
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
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
	return base.I32_extend8_s(v33) - base.I32_extend8_s(v42)
L3:
	;
	v21 = int32(1)
	if base.Ui32((v10-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	if v9&int32(255) != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v9&int32(255) != 0 {
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
	v33 = v10 | int32(32)
	goto L13
L12:
	;
	v33 = v10
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
	v42 = v9 | int32(32)
	goto L16
L15:
	;
	v42 = v9
	goto L16
L16:
	;
	if v33 == v42&int32(255) {
		v5 = v5 + v21
		v6 = v6 + v21
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = v5
	v10 = v6
	goto L1
L1:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return base.I32_extend8_s(v35) - base.I32_extend8_s(v44)
L3:
	;
	v23 = int32(1)
	if base.Ui32((v12-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	if v11&int32(255) != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v11&int32(255) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(1)
L8:
	;
	v21 = int32(-1)
	goto L10
L9:
	;
	v21 = int32(0)
	goto L10
L10:
	;
	return v21
L11:
	;
	v35 = v12 | int32(32)
	goto L13
L12:
	;
	v35 = v12
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
	v44 = v11 | int32(32)
	goto L16
L15:
	;
	v44 = v11
	goto L16
L16:
	;
	if v35 == v44&int32(255) {
		v9 = v9 + v23
		v10 = v10 + v23
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
			F_errcontext_msg(m, int32(702111), v5)
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v11 = v6
	v12 = v8
	goto L1
L1:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	return base.I32_extend8_s(v37) - base.I32_extend8_s(v46)
L3:
	;
	v25 = int32(1)
	if base.Ui32((v14-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	if v13&int32(255) != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v13&int32(255) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(1)
L8:
	;
	v23 = int32(-1)
	goto L10
L9:
	;
	v23 = int32(0)
	goto L10
L10:
	;
	return v23
L11:
	;
	v37 = v14 | int32(32)
	goto L13
L12:
	;
	v37 = v14
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
	v46 = v13 | int32(32)
	goto L16
L15:
	;
	v46 = v13
	goto L16
L16:
	;
	if v37 == v46&int32(255) {
		v11 = v11 + v25
		v12 = v12 + v25
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L2
}
