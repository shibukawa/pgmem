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
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
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
		v191 = v3
		goto L7
	} else {
		goto L8
	}
L7:
	;
	m.G0 = v10 + int32(16)
	return v191
L8:
	;
	v27 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v27
	v32 = l0 + int32(16)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v33 <= int32(0) {
		v191 = v3
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
	v191 = v179
	goto L7
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v47
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+3)))
	if v50 != 0 {
		v179 = v40
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v183 = v181 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v183 <= v185 {
		v40 = v179
		goto L10
	} else {
		goto L57
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
	if v24&int32(3) == int32(0) {
		v76 = v24
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v109 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	v109 = v101 - v24
	goto L16
L18:
	;
	v80 = v76
	goto L27
L19:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v60 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v109 = int32(0)
	goto L16
L21:
	;
	goto L22
L22:
	;
	v65 = v24
	goto L23
L23:
	;
	v69 = v65 + int32(1)
	if v69&int32(3) == int32(0) {
		v76 = v69
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v101 = v69
	goto L17
L25:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v74 != 0 {
		v65 = v69
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v89 = int32(-2139062144)
	if (int32(16843008)-v86|v86)&v89 == v89 {
		v80 = v80 + int32(4)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v95 = v80
	goto L30
L29:
	;
	goto L28
L30:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v99 != 0 {
		v95 = v95 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v101 = v95
	goto L17
L32:
	;
	goto L31
L33:
	;
	if v153 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L34:
	;
	v153 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v115 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v116 = v51
	v117 = v24
	v118 = v109
	v119 = v115
	goto L41
L38:
	;
	v141 = v24
	v145 = int32(0)
	goto L39
L39:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	v153 = v145 - v146
	goto L33
L40:
	;
	v141 = v136
	v145 = v138
	goto L39
L41:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v119 != v121 {
		v136 = v117
		v138 = v119
		goto L40
	} else {
		goto L43
	}
L42:
	;
	v136 = v130
	v138 = int32(0)
	goto L40
L43:
	;
	if v121 == int32(0) {
		v136 = v117
		v138 = v119
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v126 = v118 - int32(1)
	if v126 == int32(0) {
		v136 = v117
		v138 = v119
		goto L40
	} else {
		goto L45
	}
L45:
	;
	v129 = int32(1)
	v130 = v117 + v129
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v131 != 0 {
		v116 = v116 + v129
		v117 = v130
		v118 = v126
		v119 = v131
		goto L41
	} else {
		goto L46
	}
L46:
	;
	goto L42
L47:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+v109))))
	if v157 == int32(61) {
		v179 = v40
		goto L13
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v40 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L49
L51:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v174 + int32(1)
	v179 = v173
	goto L13
L52:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v165 = F_array_set(m, v40, v10+int32(8), v162, int32(-1), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v171 = F_construct_array_builtin(m, v10+int32(4), int32(1), int32(25))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	v173 = v165
	goto L51
L56:
	;
	v173 = v171
	goto L51
L57:
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
	v9 = F_pg_tzset(m, int32(498074))
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
	v50 = F_set_config_with_handle(m, int32(250484), v38, int32(421529), v43, v44, v44, v38, v43, v38, v38)
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
	v62 = F_set_config_with_handle(m, int32(18262), v53, int32(229579), v55, v56, v56, v53, v55, v53, v53)
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
	v74 = F_set_config_with_handle(m, int32(377625), v65, int32(229579), v67, v68, v68, v65, v67, v65, v65)
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
	v2 = int32(4441896)
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
			F_errcontext_msg(m, int32(662356), v5)
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
