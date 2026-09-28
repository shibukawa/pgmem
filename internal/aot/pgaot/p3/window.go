package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_find_window_run_conditions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	v5 = l4
	v8 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v19)
	v23 = l2
	goto L1
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v35 != int32(27) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v201
L3:
	;
	if v35 != int32(11) {
		v201 = v8
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v23 = v211
	goto L1
L5:
	;
	goto L2
L6:
	;
	m.G0 = v17 + int32(16)
	goto L5
L7:
	;
	v40 = F_contain_subplans(m, v23)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	if v40 != 0 {
		v201 = v8
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v45 = F_get_func_support(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v45 == int32(0) {
		v201 = v8
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)+28))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	if v5 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v53 = int32(4)
	goto L15
L14:
	;
	v53 = int32(0)
	goto L15
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50+v53)))
	v56 = F_is_pseudo_constant_clause(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	if v56 == int32(0) {
		v201 = v8
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61+v62<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(470)
	v75 = F_OidFunctionCall1Coll(m, v45, int32(0), base.I64_extend_i32_u(v17))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v77 = base.I32_wrap_i64(v75)
	if v77 == int32(0) {
		v201 = v8
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v80 == int32(0) {
		v201 = v8
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v84 = F_get_op_index_interpretation(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	if v84 == int32(0) {
		v201 = v8
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v88 <= int32(0) {
		v201 = v8
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v91 = int32(0)
	if v91 < v88 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v94 = v88
	goto L26
L25:
	;
	v94 = v91
	goto L26
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v105 = int32(0)
	goto L29
L27:
	;
	v172 = F_palloc0(m, int32(20))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L57
	}
L28:
	;
	v166 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v166)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v170 = v168
	goto L27
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v95+v105<<(uint(int32(2))%32))))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v116 = int32(1)
	if base.Ui32(v115-v116) <= base.Ui32(v116) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v145 = int32(3)
	if v144&v145 == v145 {
		goto L28
	} else {
		goto L51
	}
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v5 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v115&int32(-2) == int32(4) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	if v120&int32(1) != 0 {
		goto L28
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v120&int32(2) != 0 {
		goto L28
	} else {
		goto L38
	}
L37:
	;
	v201 = int32(0)
	goto L6
L38:
	;
	v201 = int32(0)
	goto L6
L39:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v5 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	if v115 != int32(3) {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	if v131&int32(2) != 0 {
		goto L28
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v131&int32(1) != 0 {
		goto L28
	} else {
		goto L46
	}
L45:
	;
	v201 = int32(0)
	goto L6
L46:
	;
	v201 = int32(0)
	goto L6
L47:
	;
	v142 = v105 + int32(1)
	if v142 == v94 {
		v201 = int32(0)
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	goto L30
L50:
	;
	v105 = v142
	goto L29
L51:
	;
	v149 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v149)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	if v5^base.B2i32(v144&v149 == int32(0)) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v161 = int32(2)
	goto L54
L53:
	;
	v161 = int32(4)
	goto L54
L54:
	;
	v162 = F_get_opfamily_member_for_cmptype(m, v151, v152, v153, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	if l3 != 0 {
		v170 = v162
		goto L27
	} else {
		goto L56
	}
L56:
	;
	v201 = int32(0)
	goto L6
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = int32(12)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v172)+12)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v172)+8)) = v177
	v180 = F_copyObjectImpl(m, v55)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+16)) = v180
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v184 = F_lappend(m, v183, v172)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v184
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v190 = F_bms_add_member(m, v187, l1+int32(7))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v190
	v201 = int32(1)
	goto L6
}
func F_show_window_keys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v17 = F_set_deparse_context_plan(m, v15, v16, l4)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l5)+56))
	if v20 <= v19 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+4)))
	v24 = v23
	goto L5
L4:
	;
	v24 = v19
	goto L5
L5:
	;
	if l2 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L48
	}
L7:
	;
	m.G0 = v13 + int32(16)
	return
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3))))
	if v27 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v66 == int32(0) {
		v174 = v28
		goto L6
	} else {
		goto L22
	}
L10:
	;
	goto L9
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v32 <= int32(0) {
		v66 = int32(0)
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v66 = int32(0)
	goto L10
L14:
	;
	v35 = int32(0)
	if v35 < v32 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v38 = v32
	goto L17
L16:
	;
	v38 = v35
	goto L17
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v43 = int32(0)
	goto L18
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v39+v43<<(uint(int32(2))%32))))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+8)))
	if v52 == v28&int32(_a_F_show_window_keys_0) {
		v66 = v51
		goto L10
	} else {
		goto L20
	}
L19:
	;
	goto L13
L20:
	;
	v55 = v43 + int32(1)
	if v55 != v38 {
		v43 = v55
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v70 = int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v75 = F_deparse_expression(m, v71, v17, v24&v70, v70)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_appendStringInfoString(m, l0, v75)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_pfree(m, v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if l2 == int32(1) {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v88 = v70
	goto L27
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3+v88<<(uint(int32(1))%32)))))
	if v93 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L7
L29:
	;
	if v135 == int32(0) {
		v174 = v97
		goto L6
	} else {
		goto L42
	}
L30:
	;
	goto L29
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v101 <= int32(0) {
		v135 = int32(0)
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v135 = int32(0)
	goto L30
L34:
	;
	v104 = int32(0)
	if v104 < v101 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v107 = v101
	goto L37
L36:
	;
	v107 = v104
	goto L37
L37:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v112 = int32(0)
	goto L38
L38:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v108+v112<<(uint(int32(2))%32))))
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+8)))
	if v121 == v97&int32(_a_F_show_window_keys_0) {
		v135 = v120
		goto L30
	} else {
		goto L40
	}
L39:
	;
	goto L33
L40:
	;
	v124 = v112 + int32(1)
	if v124 != v107 {
		v112 = v124
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v140 = int32(1)
	v143 = F_deparse_expression(m, v139, v17, v24&v140, v140)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_show_window_keys_1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_appendStringInfoString(m, l0, v143)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_pfree(m, v143)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v153 = v88 + int32(1)
	if v153 != l2 {
		v88 = v153
		goto L27
	} else {
		goto L47
	}
L47:
	;
	goto L28
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v174
	F_errmsg_internal(m, int32(_a_F_show_window_keys_2), v13)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_show_window_keys_3), int32(2991), int32(_a_F_show_window_keys_4))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_window_dense_rank(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_WinCheckAndInitializeNullTreatment(m, v5, v3, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+176))
		v14 = F_WinGetPartitionLocalMemory(m, v5, int32(8))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
			if v16 == int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(1)
				v27 = v3
				F_WinSetMarkPosition(m, v5, v12)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v31 = F_WinGetPartitionLocalMemory(m, v5, int32(8))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						v33 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
						if v27 != 0 {
							v35 = v33 + int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(v31))) = v35
							v37 = v35
						} else {
							v37 = v33
						}
						return v37
					}
				}
			} else {
				v23 = F_WinRowsArePeers(m, v5, v12-int64(1), v12)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v27 = v23 ^ int32(1)
					F_WinSetMarkPosition(m, v5, v12)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						v31 = F_WinGetPartitionLocalMemory(m, v5, int32(8))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							v33 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
							if v27 != 0 {
								v35 = v33 + int64(1)
								*(*int64)(unsafe.Add(mBase, uint32(v31))) = v35
								v37 = v35
							} else {
								v37 = v33
							}
							return v37
						}
					}
				}
			}
		}
	}
}
func F_window_lag(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14405(m, l0, int32(-1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_window_last_value(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14404(m, l0, int32(2))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_window_nth_value(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int64
	_ = v94
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_WinCheckAndInitializeNullTreatment(m, v12, int32(1), l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v20 = v10 + int32(15)
		v21 = F_WinGetFuncArgCurrent(m, v12, int32(1), v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
			if v23 == int32(0) {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v28 = int32(0)
				if v26 == v28 {
					v74 = v28
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
					if v32 == int32(0) {
						v74 = v28
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
						v37 = v35 - int32(11)
						if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v37))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v37)%32))&int32(1) == int32(0))|int32(0) != 0 {
							v74 = v28
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v37<<(uint(int32(2))%32))+uint32(_c_F_window_nth_value[0])))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v32+v52)))
							if v54 == int32(0) {
								v74 = v28
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
								if v57 <= int32(1) {
									v74 = v28
								} else {
									v59 = int32(1)
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+int32(4))))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
									switch v65 - int32(7) {
									case 0:
										v74 = v59
									case 1:
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
										if v68 == int32(0) {
											v74 = v59
										} else {
											v74 = int32(0)
										}
									default:
										v74 = int32(0)
									}
								}
							}
						}
					}
				}
				v75 = base.I32_wrap_i64(v21)
				if v75 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(100925570))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_window_nth_value_0), int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_window_nth_value_1), int32(717), int32(_a_F_window_nth_value_2))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v78 = int32(1)
					v81 = F_WinGetFuncArgInFrame(m, v12, v75-v78, v78, v74, v20)
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int64(0)
					} else {
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v83 != int32(1) {
							v94 = v81
						} else {
							v89 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
							v94 = int64(0)
						}
						m.G0 = v10 + int32(16)
						return v94
					}
				}
			} else {
				v89 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
				v94 = int64(0)
				m.G0 = v10 + int32(16)
				return v94
			}
		}
	}
}
func F_window_percent_rank(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v53 int64
	_ = v53
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = F_WinGetPartitionRowCount(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		F_WinCheckAndInitializeNullTreatment(m, v6, int32(0), l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)+176))
			v17 = F_WinGetPartitionLocalMemory(m, v6, int32(8))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				if v19 == int64(0) {
					*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(1)
					v30 = int32(0)
					F_WinSetMarkPosition(m, v6, v15)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						v34 = F_WinGetPartitionLocalMemory(m, v6, int32(8))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							if v30 != 0 {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
								v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)+176))
								*(*int64)(unsafe.Add(mBase, uint32(v34))) = v37 + int64(1)
							} else {
							}
							if int64(2) <= v7 {
								v43 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
								v44 = int64(1)
								v53 = base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v43-v44), base.F64_convert_i64_u(v7-v44)))
							} else {
								v53 = int64(0)
							}
							return v53
						}
					}
				} else {
					v26 = F_WinRowsArePeers(m, v6, v15-int64(1), v15)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v30 = v26 ^ int32(1)
						F_WinSetMarkPosition(m, v6, v15)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							v34 = F_WinGetPartitionLocalMemory(m, v6, int32(8))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								if v30 != 0 {
									v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
									v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)+176))
									*(*int64)(unsafe.Add(mBase, uint32(v34))) = v37 + int64(1)
								} else {
								}
								if int64(2) <= v7 {
									v43 = *(*int64)(unsafe.Add(mBase, uint32(v34)))
									v44 = int64(1)
									v53 = base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v43-v44), base.F64_convert_i64_u(v7-v44)))
								} else {
									v53 = int64(0)
								}
								return v53
							}
						}
					}
				}
			}
		}
	}
}
func F_window_rank(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_WinCheckAndInitializeNullTreatment(m, v5, v2, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+176))
		v14 = F_WinGetPartitionLocalMemory(m, v5, int32(8))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
			if v16 == int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(1)
				v27 = v2
				F_WinSetMarkPosition(m, v5, v12)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v31 = F_WinGetPartitionLocalMemory(m, v5, int32(8))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						if v27 == int32(0) {
							v35 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
							return v35
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
							v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+176))
							v40 = v38 + int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(v31))) = v40
							return v40
						}
					}
				}
			} else {
				v23 = F_WinRowsArePeers(m, v5, v12-int64(1), v12)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v27 = v23 ^ int32(1)
					F_WinSetMarkPosition(m, v5, v12)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						v31 = F_WinGetPartitionLocalMemory(m, v5, int32(8))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							if v27 == int32(0) {
								v35 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
								return v35
							} else {
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
								v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+176))
								v40 = v38 + int64(1)
								*(*int64)(unsafe.Add(mBase, uint32(v31))) = v40
								return v40
							}
						}
					}
				}
			}
		}
	}
}
