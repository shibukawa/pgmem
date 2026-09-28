package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_finalize_agg_primnode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v3 == int32(9) {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v7 = F_finalize_primnode(m, v6, l1)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v12 = F_finalize_primnode(m, v11, l1)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		} else {
			v17 = F_expression_tree_walker_impl(m, l0, int32(897), l1)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v20 = v17
				return v20
			}
		}
	} else {
		v20 = int32(0)
		return v20
	}
}
func F_lookup_agg_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v29 = F_func_get_detail(m, l0, v6, v6, l1, l2, v6, v6, v6, v9+int32(56), v9+int32(76), l4, v9+int32(75), v9+int32(68), v9-int32(-64), v9+int32(60), v6)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L44
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L39
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L34
	}
L4:
	;
	return int32(0)
L5:
	;
	if v29 != int32(2) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	if v35 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+75)))
	if v38 == int32(1) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	if l3 == int32(2276) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
	if v43 != int32(2276) {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v49 = F_enforce_generic_type_consistency(m, l2, v46, l1, v47, int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v49
	v52 = int32(0)
	if l1 <= v52 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_agg_function[0]))
	v106 = F_object_aclcheck(m, int32(1255), v102, v104, int64(128))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L28
	}
L15:
	;
	v59 = v52
	goto L16
L16:
	;
	v62 = v59 << (uint(int32(2)) % 32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2+v62)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v65+v62)))
	v68 = F_IsBinaryCoercible(m, v64, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L23
	}
L18:
	;
	if v68 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v71 = v59 + int32(1)
	if l1 != v71 {
		v59 = v71
		goto L16
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L17
L22:
	;
	goto L14
L23:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v82 = F_func_signature_string(m, l0, l1, int32(0), v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v82
	F_errmsg(m, int32(_a_F_lookup_agg_function_0), v9+int32(32))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_lookup_agg_function_1), int32(939), int32(_a_F_lookup_agg_function_2))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	if v106 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	v110 = F_get_func_name(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	m.G0 = v9 + int32(80)
	return v114
L32:
	;
	F_aclcheck_error(m, v106, int32(19), v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v127 = F_func_signature_string(m, l0, l1, int32(0), l2)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v127
	F_errmsg(m, int32(_a_F_lookup_agg_function_3), v9+int32(48))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_lookup_agg_function_1), int32(894), int32(_a_F_lookup_agg_function_2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v148 = F_func_signature_string(m, l0, l1, int32(0), l2)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v148
	F_errmsg(m, int32(_a_F_lookup_agg_function_4), v9)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_lookup_agg_function_1), int32(900), int32(_a_F_lookup_agg_function_2))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v167 = F_func_signature_string(m, l0, l1, int32(0), l2)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v167
	F_errmsg(m, int32(_a_F_lookup_agg_function_5), v9+int32(16))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_lookup_agg_function_1), int32(915), int32(_a_F_lookup_agg_function_2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
