package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RemoveSQLFunctionCache(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 != 0 {
		v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v3)+24)) = v4 - int64(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	} else {
	}
	return
}
func F_sql_compile_callback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int64
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v264 int32
	_ = v264
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
	v20 = int32(_a_F_sql_compile_callback_0)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_sql_compile_callback[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_sql_compile_callback[0])) = v15 + int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = int32(732)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v21
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_sql_compile_callback[1]))
	v36 = F_AllocSetContextCreateInternal(m, v31, int32(_a_F_sql_compile_callback_1), v6, int32(1024), int32(_a_F_sql_compile_callback_2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v42 = F_AllocSetContextCreateInternal(m, v36, int32(_a_F_sql_compile_callback_3), int32(0), int32(1024), int32(_a_F_sql_compile_callback_2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+80)) = v42
	v45 = v17 + v18
	v48 = F_MemoryContextStrdup(m, v36, v45+int32(4))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+32)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v36)+36)) = v48
	v56 = F_get_call_result_type(m, l0, v15+int32(32), v15+int32(28))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+48)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v60 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_sql_compile_callback[1])) = v36
	v63 = F_CreateTupleDescCopy(m, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v69 = v58
	goto L8
L8:
	;
	F_get_typlenbyval(m, v69, l3+int32(52), l3+int32(54))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v63
	*(*int32)(unsafe.Add(mBase, _c_F_sql_compile_callback[1])) = v31
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v69 = v68
	goto L8
L10:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+55)) = uint8(v76)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+101)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+57)) = uint8(base.B2i32(v78 != int32(118)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+58)) = uint8(v82)
	*(*int32)(unsafe.Add(mBase, _c_F_sql_compile_callback[1])) = v36
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+24))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v89 = F_prepare_sql_fn_parse_info(m, l1, v87, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v89
	*(*int32)(unsafe.Add(mBase, _c_F_sql_compile_callback[1])) = v31
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v97 = F_MemoryContextAlloc(m, v36, v94<<(uint(int32(1))%32))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+44)) = v97
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if int32(0) < v101 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v104 = v100
	v108 = v6
	goto L16
L14:
	;
	goto L15
L15:
	;
	v147 = F_SysCacheGetAttrNotNull(m, int32(47), l1, int32(26))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L20
	}
L16:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+v108<<(uint(int32(2))%32))))
	v121 = F_get_typlen(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v124 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v123+v108<<(uint(v124)%32)))) = uint16(v121)
	v129 = v108 + v124
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v129 < v131 {
		v104 = v130
		v108 = v129
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v150 = F_text_to_cstring(m, base.I32_wrap_i64(v147))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v152 = F_MemoryContextStrdup(m, v36, v150)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v152
	v159 = F_SysCacheGetAttr(m, int32(47), l1, int32(28), v15+int32(27))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+27)))
	if v161 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+72)) = uint8(v187)
	if v189 != 0 {
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v165 = F_text_to_cstring(m, base.I32_wrap_i64(v159))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	v185 = F_pg_parse_query(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L34
	}
L28:
	;
	v167 = F_stringToNode(m, v165)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v169 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v187 = int32(0)
	v189 = v174
	goto L24
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v167
	v181 = F_list_make1_impl(m, int32(1), v15+int32(16))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v187 = int32(0)
	v189 = v181
	goto L24
L34:
	;
	v187 = int32(1)
	v189 = v185
	goto L24
L35:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v193 = v191
	goto L37
L36:
	;
	v193 = int32(0)
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+68)) = v193
	if v193 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_sql_compile_callback[1])) = v42
	v223 = F_copyObjectImpl(m, v189)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L47
	}
L39:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	if v195 == int32(2278) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v206 = F_format_type_be(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v206
	F_errmsg(m, int32(_a_F_sql_compile_callback_4), v15)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v214 = F_errdetail(m, int32(_a_F_sql_compile_callback_5), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_sql_compile_callback_6), int32(1188), int32(_a_F_sql_compile_callback_7))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+64)) = v223
	*(*int32)(unsafe.Add(mBase, _c_F_sql_compile_callback[1])) = v31
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_sql_compile_callback[2]))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	if v233 != v229 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+84)) = v36
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_sql_compile_callback[0])) = v264
	m.G0 = v15 + int32(48)
	return
L49:
	;
	if v233 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	if v229 != 0 {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	if v238 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v237 == int32(0) {
		goto L52
	} else {
		goto L58
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+28)) = v237
	goto L54
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v233)+20)) = v237
	goto L54
L58:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v237)+24)) = v243
	goto L52
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v229
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v229)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = v250
	if v250 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = int32(0)
	goto L51
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+24)) = v36
	goto L64
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229)+20)) = v36
	goto L48
}
func F_sql_function_parse_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = F_function_parse_error_transpose(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		if v8 == int32(0) {
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v15
				F_errcontext_msg(m, int32(_a_F_sql_function_parse_error_callback_0), v5)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
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
}
