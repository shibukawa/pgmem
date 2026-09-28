package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_define_custom_variable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_define_custom_variable[0]))
	v19 = F_hash_search(m, v14, v9+int32(8), v2, v2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L57
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L53
	}
L3:
	;
	m.G0 = v9 + int32(16)
	return
L4:
	;
	return
L5:
	;
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_InitializeOneGUCOption(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+21)))
	if v36&int32(2) == int32(0) {
		goto L1
	} else {
		goto L12
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_define_custom_variable[0]))
	v30 = F_hash_search(m, v26, l0, int32(3), v9+int32(15))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v30 == int32(0) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = l0
	goto L3
L12:
	;
	F_InitializeOneGUCOption(m, l0)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v35)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v50
	goto L16
L15:
	;
	goto L16
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	if v53 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v58 = int32(_a_F_define_custom_variable_0)
	goto L22
L18:
	;
	goto L19
L19:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+28)))
	if v67&int32(4) != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	goto L20
L22:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v61 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v65
	goto L21
L24:
	;
	if v61 != v35+int32(76) {
		v58 = v61
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v74 = int32(_a_F_define_custom_variable_1)
	goto L31
L27:
	;
	goto L28
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v35)+116))
	if v83 != 0 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	goto L29
L31:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v77 == int32(0) {
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v81
	goto L30
L33:
	;
	if v77 != v35+int32(80) {
		v74 = v77
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v85 = int32(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v35)+44))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v35)+36))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v35)+52))
	v93 = F_set_config_with_handle(m, v84, v85, v83, v86, v87, v88, v85, int32(1), int32(19), v85)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v35)+96))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v35)+40))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
	F_reapply_stacked_values(m, l0, v35, v95, v97, v98, v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v35)+88))
	if v103 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_free_placeholder(m, v35)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L52
	}
L41:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v35)+92))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_define_custom_variable[1])))
	if v113 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v114 = int32(12)
	goto L44
L43:
	;
	v114 = int32(15)
	goto L44
L44:
	;
	v115 = F_find_option(m, v107, int32(1), int32(0), v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	if v115 == int32(0) {
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v119 = F_guc_strdup(m, v114, v103)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)+88))
	if v121 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_pfree(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+92)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v115)+88)) = v119
	goto L40
L51:
	;
	goto L50
L52:
	;
	goto L3
L53:
	;
	F_errcode(m, int32(_a_F_define_custom_variable_2))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(_a_F_define_custom_variable_3), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_define_custom_variable_4), int32(941), int32(_a_F_define_custom_variable_5))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v163
	F_errmsg(m, int32(_a_F_define_custom_variable_6), v9)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_define_custom_variable_4), int32(_a_F_define_custom_variable_7), int32(_a_F_define_custom_variable_8))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
