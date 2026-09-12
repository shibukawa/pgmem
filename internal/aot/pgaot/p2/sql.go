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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v260 int32
	_ = v260
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
	v19 = int32(4547032)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v14 + int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = int32(686)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v20
	v30 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v35 = F_AllocSetContextCreateInternal(m, v30, int32(266536), v6, int32(1024), int32(8192))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v41 = F_AllocSetContextCreateInternal(m, v35, int32(181156), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+80)) = v41
	v44 = v16 + v17
	v47 = F_MemoryContextStrdup(m, v35, v44+int32(4))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+32)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v35)+36)) = v47
	goto L5
L5:
	;
	v55 = F_get_call_result_type(m, l0, v14+int32(32), v14+int32(28))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+48)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	if v59 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v35
	v62 = F_CreateTupleDescCopy(m, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v68 = v57
	goto L9
L9:
	;
	F_get_typlenbyval(m, v68, l3+int32(52), l3+int32(54))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v62
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v30
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v68 = v67
	goto L9
L11:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+55)) = uint8(v75)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+101)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+57)) = uint8(base.B2i32(v77 != int32(118)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+58)) = uint8(v81)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v35
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+24))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v88 = F_prepare_sql_fn_parse_info(m, l1, v86, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v88
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v30
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v96 = F_MemoryContextAlloc(m, v35, v93<<(uint(int32(1))%32))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+44)) = v96
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if int32(0) < v100 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v103 = v99
	v107 = v6
	goto L17
L15:
	;
	goto L16
L16:
	;
	v144 = F_SysCacheGetAttrNotNull(m, int32(47), l1, int32(26))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L21
	}
L17:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v107<<(uint(int32(2))%32))))
	v119 = F_get_typlen(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v122 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v121+v107<<(uint(v122)%32)))) = uint16(v119)
	v127 = v107 + v122
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v127 < v129 {
		v103 = v128
		v107 = v127
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v146 = F_text_to_cstring(m, v144)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v148 = F_MemoryContextStrdup(m, v35, v146)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v148
	v155 = F_SysCacheGetAttr(m, int32(47), l1, int32(28), v14+int32(27))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+27)))
	if v157 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+72)) = uint8(v182)
	if v184 != 0 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v160 = F_text_to_cstring(m, v155)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l3)+36))
	v180 = F_pg_parse_query(m, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L35
	}
L29:
	;
	v162 = F_stringToNode(m, v160)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v164 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v182 = int32(0)
	v184 = v169
	goto L25
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v162
	v176 = F_list_make1_impl(m, int32(1), v14+int32(16))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v182 = int32(0)
	v184 = v176
	goto L25
L35:
	;
	v182 = int32(1)
	v184 = v180
	goto L25
L36:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v188 = v186
	goto L38
L37:
	;
	v188 = int32(0)
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+68)) = v188
	if v188 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v41
	v218 = F_copyObjectImpl(m, v184)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L48
	}
L40:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if v190 == int32(2278) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	v201 = F_format_type_be(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v201
	F_errmsg(m, int32(193573), v14)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errdetail(m, int32(687226), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(516223), int32(1187), int32(333692))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+64)) = v218
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v30
	v224 = *(*int32)(unsafe.Add(mBase, _consts[403]))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	if v228 != v224 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+84)) = v35
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v260
	m.G0 = v14 + int32(48)
	return
L50:
	;
	if v228 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	if v224 != 0 {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
	if v233 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v232 == int32(0) {
		goto L53
	} else {
		goto L59
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v233)+28)) = v232
	goto L55
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228)+20)) = v232
	goto L55
L59:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+24)) = v238
	goto L53
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v224
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v224)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v245
	if v245 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = int32(0)
	goto L52
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+24)) = v35
	goto L65
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = v35
	goto L49
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
				F_errcontext_msg(m, int32(736102), v5)
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
