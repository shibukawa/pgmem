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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
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
	return v200
L3:
	;
	if v35 != int32(11) {
		v200 = v8
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v23 = v209
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
		v200 = v8
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
		v200 = v8
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
		v200 = v8
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
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(462)
	v74 = F_OidFunctionCall1Coll(m, v45, int32(0), v17)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	if v74 == int32(0) {
		v200 = v8
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	if v78 == int32(0) {
		v200 = v8
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v82 = F_get_op_index_interpretation(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	if v82 == int32(0) {
		v200 = v8
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v86 <= int32(0) {
		v200 = v8
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v89 = int32(0)
	if v89 < v86 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v92 = v86
	goto L26
L25:
	;
	v92 = v89
	goto L26
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v102 = int32(0)
	goto L29
L27:
	;
	v170 = F_palloc0(m, int32(20))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L8
	} else {
		goto L57
	}
L28:
	;
	v164 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v164)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v168 = v166
	goto L27
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v93+v102<<(uint(int32(2))%32))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v114 = int32(1)
	if base.Ui32(v113-v114) <= base.Ui32(v114) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v143 = int32(3)
	if v142&v143 == v143 {
		goto L28
	} else {
		goto L51
	}
L31:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
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
	if v113&int32(-2) == int32(4) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	if v118&int32(1) != 0 {
		goto L28
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v118&int32(2) != 0 {
		goto L28
	} else {
		goto L38
	}
L37:
	;
	v200 = int32(0)
	goto L6
L38:
	;
	v200 = int32(0)
	goto L6
L39:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
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
	if v113 != int32(3) {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	if v129&int32(2) != 0 {
		goto L28
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v129&int32(1) != 0 {
		goto L28
	} else {
		goto L46
	}
L45:
	;
	v200 = int32(0)
	goto L6
L46:
	;
	v200 = int32(0)
	goto L6
L47:
	;
	v140 = v102 + int32(1)
	if v140 == v92 {
		v200 = int32(0)
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
	v102 = v140
	goto L29
L51:
	;
	v147 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v147)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	if v5^base.B2i32(v142&v147 == int32(0)) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v159 = int32(2)
	goto L54
L53:
	;
	v159 = int32(4)
	goto L54
L54:
	;
	v160 = F_get_opfamily_member_for_cmptype(m, v149, v150, v151, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	if l3 != 0 {
		v168 = v160
		goto L27
	} else {
		goto L56
	}
L56:
	;
	v200 = int32(0)
	goto L6
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+4)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = int32(12)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v170)+12)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = v175
	v178 = F_copyObjectImpl(m, v55)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+16)) = v178
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v182 = F_lappend(m, v181, v170)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v182
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v188 = F_bms_add_member(m, v185, l1+int32(7))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v188
	v200 = int32(1)
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
	var v177 int32
	_ = v177
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
		v177 = v28
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
	if v52 == v28&int32(65535) {
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
		v177 = v97
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
	if v121 == v97&int32(65535) {
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
	F_appendStringInfoString(m, l0, int32(783295))
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
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v177
	F_errmsg_internal(m, int32(490571), v13)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(522511), int32(2974), int32(120822))
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
func F_window_dense_rank(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)+176))
	v8 = F_WinGetPartitionLocalMemory(m, v4, int32(8))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		if v12 == int64(0) {
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = int64(1)
			v24 = int32(0)
			F_WinSetMarkPosition(m, v4, v6)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v28 = F_WinGetPartitionLocalMemory(m, v4, int32(8))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
					if v24 != 0 {
						v32 = v30 + int64(1)
						*(*int64)(unsafe.Add(mBase, uint32(v28))) = v32
						v34 = v32
					} else {
						v34 = v30
					}
					v35 = F_Int64GetDatum(m, v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						return v35
					}
				}
			}
		} else {
			v20 = F_WinRowsArePeers(m, v4, v6-int64(1), v6)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v24 = v20 ^ int32(1)
				F_WinSetMarkPosition(m, v4, v6)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v28 = F_WinGetPartitionLocalMemory(m, v4, int32(8))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
						if v24 != 0 {
							v32 = v30 + int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(v28))) = v32
							v34 = v32
						} else {
							v34 = v30
						}
						v35 = F_Int64GetDatum(m, v34)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							return v35
						}
					}
				}
			}
		}
	}
}
func F_window_lag(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = F_WinGetFuncArgInPartition(m, v8, int32(-1), int32(1), v6+int32(15), v6+int32(14))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v19 == int32(1) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			v25 = int32(0)
		} else {
			v25 = v15
		}
		m.G0 = v6 + int32(16)
		return v25
	}
}
func F_window_last_value(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_WinGetFuncArgInFrame(m, v8, int32(0), int32(2), int32(1), v6+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v24 = int32(0)
		} else {
			v24 = v14
		}
		m.G0 = v6 + int32(16)
		return v24
	}
}
func F_window_nth_value(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_WinGetFuncArgCurrent(m, v10, int32(1), v8+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v18 == int32(0) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v23 = int32(0)
			if v21 == v23 {
				v68 = v23
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
				if v27 == int32(0) {
					v68 = v23
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					v32 = v30 - int32(11)
					if base.Ui32(int32(9)) < base.Ui32(v32) {
						v68 = v23
					} else {
						if int32(base.Ui32(int32(977))>>(uint(v32)%32))&int32(1) == int32(0) {
							v68 = v23
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_consts[1074])))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v27+v47)))
							if v49 == int32(0) {
								v68 = v23
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
								if v52 <= int32(1) {
									v68 = v23
								} else {
									v54 = int32(1)
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+int32(4))))
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
									switch v60 - int32(7) {
									case 0:
										v68 = v54
									case 1:
										v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
										if v63 == int32(0) {
											v68 = v54
										} else {
											v68 = int32(0)
										}
									default:
										v68 = int32(0)
									}
								}
							}
						}
					}
				}
			}
			if v14 <= int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v96 = m.ExcPending
				if v96 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(100925570))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(252710), int32(0))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(520136), int32(707), int32(363832))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v72 = int32(1)
				v77 = F_WinGetFuncArgInFrame(m, v10, v14-v72, v72, v68, v8+int32(15))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					if v79 != int32(1) {
						v87 = v77
					} else {
						v84 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
						v87 = int32(0)
					}
					m.G0 = v8 + int32(16)
					return v87
				}
			}
		} else {
			v84 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
			v87 = int32(0)
			m.G0 = v8 + int32(16)
			return v87
		}
	}
}
func F_window_percent_rank(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = F_WinGetPartitionRowCount(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)+176))
		v15 = F_WinGetPartitionLocalMemory(m, v7, int32(8))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
			if v17 == int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(1)
				v29 = int32(0)
				F_WinSetMarkPosition(m, v7, v13)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v33 = F_WinGetPartitionLocalMemory(m, v7, int32(8))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						if v29 != 0 {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
							v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+176))
							*(*int64)(unsafe.Add(mBase, uint32(v33))) = v36 + int64(1)
						} else {
						}
						if int64(2) <= v8 {
							v42 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
							v43 = int64(1)
							v50 = base.F64_div(base.F64_convert_i64_s(v42-v43), base.F64_convert_i64_u(v8-v43))
						} else {
							v50 = float64(0)
						}
						v51 = F_Float8GetDatum(m, v50)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							return v51
						}
					}
				}
			} else {
				v25 = F_WinRowsArePeers(m, v7, v13-int64(1), v13)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v29 = v25 ^ int32(1)
					F_WinSetMarkPosition(m, v7, v13)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						v33 = F_WinGetPartitionLocalMemory(m, v7, int32(8))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							if v29 != 0 {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
								v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+176))
								*(*int64)(unsafe.Add(mBase, uint32(v33))) = v36 + int64(1)
							} else {
							}
							if int64(2) <= v8 {
								v42 = *(*int64)(unsafe.Add(mBase, uint32(v33)))
								v43 = int64(1)
								v50 = base.F64_div(base.F64_convert_i64_s(v42-v43), base.F64_convert_i64_u(v8-v43))
							} else {
								v50 = float64(0)
							}
							v51 = F_Float8GetDatum(m, v50)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								return v51
							}
						}
					}
				}
			}
		}
	}
}
func F_window_rank(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)+176))
	v9 = F_WinGetPartitionLocalMemory(m, v5, int32(8))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
		if v13 == int64(0) {
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = int64(1)
			v25 = int32(0)
			F_WinSetMarkPosition(m, v5, v7)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v29 = F_WinGetPartitionLocalMemory(m, v5, int32(8))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					if v25 == int32(0) {
						v33 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
						v34 = F_Int64GetDatum(m, v33)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							return v34
						}
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
						v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+176))
						v40 = v38 + int64(1)
						*(*int64)(unsafe.Add(mBase, uint32(v29))) = v40
						v42 = F_Int64GetDatum(m, v40)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							return v42
						}
					}
				}
			}
		} else {
			v21 = F_WinRowsArePeers(m, v5, v7-int64(1), v7)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v25 = v21 ^ int32(1)
				F_WinSetMarkPosition(m, v5, v7)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v29 = F_WinGetPartitionLocalMemory(m, v5, int32(8))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						if v25 == int32(0) {
							v33 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
							v34 = F_Int64GetDatum(m, v33)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								return v34
							}
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
							v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)+176))
							v40 = v38 + int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(v29))) = v40
							v42 = F_Int64GetDatum(m, v40)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								return v42
							}
						}
					}
				}
			}
		}
	}
}
