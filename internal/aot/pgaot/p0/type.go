package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AssignTypeArrayOid(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AssignTypeArrayOid[0])))
	if v4 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_AssignTypeArrayOid[1]))
		if v8 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_AssignTypeArrayOid_0), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_AssignTypeArrayOid_1), int32(2502), int32(_a_F_AssignTypeArrayOid_2))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
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
			*(*int32)(unsafe.Add(mBase, _c_F_AssignTypeArrayOid[1])) = int32(0)
			return v8
		}
	} else {
		v17 = F_table_open(m, int32(1247), int32(1))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v23 = F_GetNewOidWithIndex(m, v17, int32(2703), int32(1))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				F_relation_close(m, v17, int32(1))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					return v23
				}
			}
		}
	}
}
func F_GenerateTypeDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
	v20 = v18 + v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if l2 != 0 {
		v35 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l3 != 0 {
		v47 = l3
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v27 = F_heap_getattr_5(m, l0, int32(30), v24, v16+int32(31))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+31)))
	if v29 != 0 {
		v35 = int32(0)
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v31 = F_text_to_cstring(m, base.I32_wrap_i64(v27))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v33 = F_stringToNode(m, v31)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v35 = v33
	goto L1
L8:
	;
	if l8 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v41 = F_heap_getattr_5(m, l0, int32(32), v38, v16+int32(31))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+31)))
	if v43 != 0 {
		v47 = int32(0)
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v45 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v41))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v47 = v45
	goto L8
L13:
	;
	v50 = F_deleteDependencyRecordsFor(m, int32(1247), v21, int32(1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(1247)
	v61 = F_new_object_addresses(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	F_deleteSharedDependencyRecordsFor(m, int32(1247), v21, int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	if l6 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if l7 != 0 {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+79)))
	if v63 != int32(109) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(2615)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v78
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(2615)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v68
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L19
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	F_recordDependencyOnOwner(m, int32(1247), v21, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	F_recordDependencyOnNewAcl(m, int32(1247), v21, v91, v47)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	F_recordDependencyOnCurrentExtension(m, v16+int32(16), l8)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v20)+100))
	if v99 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	if v109 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L3
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	if v119 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v20)+112))
	if v129 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v20)+116))
	if v139 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v20)+120))
	if v149 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v20)+124))
	if v159 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	if v169 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L3
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v20)+132))
	if v179 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L62
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1247)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L3
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v20)+144))
	v190 = int32(0)
	if base.B2i32(v189 == v190)|base.B2i32(v189 == int32(100)) == v190 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(3456)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L3
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v207 = v16 + int32(16)
	F_record_object_address_dependencies(m, v207, v61, int32(110))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L3
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	F_free_object_addresses(m, v61)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	if v35 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_recordDependencyOnExpr(m, v207, v35, int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L3
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	if v216 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L76
L78:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v20)+92))
	if v240 != 0 {
		goto L85
	} else {
		goto L86
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1259)
	if l4 != int32(99) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_recordDependencyOn(m, v16+int32(16), v16+int32(4), int32(105))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L3
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_recordDependencyOn(m, v16+int32(4), v16+int32(16), int32(105))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L3
	} else {
		goto L84
	}
L83:
	;
	goto L78
L84:
	;
	goto L78
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1247)
	if l5 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	m.G0 = v16 + int32(32)
	return
L88:
	;
	v252 = int32(105)
	goto L90
L89:
	;
	v252 = int32(110)
	goto L90
L90:
	;
	F_recordDependencyOn(m, v16+int32(16), v16+int32(4), v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	goto L87
}
func F_LookupTypeNameExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int64
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v279 int32
	_ = v279
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v442 int64
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v485 int32
	_ = v485
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(160)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 == v6 {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L16
	} else {
		goto L135
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L16
	} else {
		goto L129
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L16
	} else {
		goto L123
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L16
	} else {
		goto L120
	}
L5:
	;
	m.G0 = v14 + int32(160)
	return v485
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v470
	v485 = v473
	goto L5
L7:
	;
	v328 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v306))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L16
	} else {
		goto L88
	}
L8:
	;
	v325 = int32(0)
	if l2 != 0 {
		v470 = int32(-1)
		v473 = v325
		goto L6
	} else {
		goto L87
	}
L9:
	;
	if v306 != 0 {
		goto L7
	} else {
		goto L86
	}
L10:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v306 = v19
	goto L9
L11:
	;
	goto L12
L12:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	if v20 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v23 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v26 = F_makeRangeVar(m, v23, v23, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	F_DeconstructQualifiedName(m, v16, v14+int32(136), v14+int32(132))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L16
	} else {
		goto L56
	}
L16:
	;
	return int32(0)
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v30 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115+v109)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v119 = int32(0)
	v122 = F_RangeVarGetRelidExtended(m, v26, v119, l4, v119, v119)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L16
	} else {
		goto L37
	}
L19:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v109 = int32(4)
	v110 = v108
	goto L18
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L31
	}
L21:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	switch v33 - int32(1) {
	case 0:
		goto L24
	case 1:
		goto L19
	case 2:
		goto L23
	case 3:
		goto L22
	default:
		goto L20
	}
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v76
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v109 = int32(12)
	v110 = v80 + int32(8)
	goto L18
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v62
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v109 = int32(8)
	v110 = v66 + int32(4)
	goto L18
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L16
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v44 = F_NameListToString(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v44
	F_errmsg(m, int32(_a_F_LookupTypeNameExtended_0), v14+int32(80))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L16
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_LookupTypeNameExtended_1), int32(102), int32(_a_F_LookupTypeNameExtended_2))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v91 = F_NameListToString(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v91
	F_errmsg(m, int32(_a_F_LookupTypeNameExtended_3), v14-int32(-64))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_LookupTypeNameExtended_1), int32(124), int32(_a_F_LookupTypeNameExtended_2))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v124 = F_get_attnum(m, v122, v118)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L16
	} else {
		goto L38
	}
L38:
	;
	if v124 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if l4 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v151 = F_get_atttype(m, v122, v124)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L16
	} else {
		goto L48
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L16
	} else {
		goto L43
	}
L43:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v118
	F_errmsg(m, int32(_a_F_LookupTypeNameExtended_4), v14+int32(96))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_LookupTypeNameExtended_1), int32(146), int32(_a_F_LookupTypeNameExtended_2))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L16
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
	v155 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L16
	} else {
		goto L49
	}
L49:
	;
	if v155 == int32(0) {
		v306 = v151
		goto L9
	} else {
		goto L50
	}
L50:
	;
	v160 = v14 + int32(140)
	F_initStringInfo(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L16
	} else {
		goto L51
	}
L51:
	;
	F_appendTypeNameToBuffer(m, l1, v160)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L16
	} else {
		goto L52
	}
L52:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v14)+140))
	v166 = F_format_type_be(m, v151)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v165
	F_errmsg(m, int32(_a_F_LookupTypeNameExtended_5), v14+int32(112))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L16
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_LookupTypeNameExtended_1), int32(159), int32(_a_F_LookupTypeNameExtended_2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L16
	} else {
		goto L55
	}
L55:
	;
	v306 = v151
	goto L9
L56:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v14)+136))
	if v186 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v297 == int32(0) {
		v306 = v290
		goto L9
	} else {
		goto L84
	}
L58:
	;
	v188 = v14 + int32(140)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = int32(524)
	*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v188)+16)) = v188
	v195 = int32(_a_F_LookupTypeNameExtended_6)
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = v196
	*(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[0])) = v14 + int32(148)
	goto L61
L59:
	;
	goto L60
L60:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v14)+132))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L16
	} else {
		goto L68
	}
L61:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v14)+136))
	v203 = F_LookupExplicitNamespace(m, v202, l4)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	if v203 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v206 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v14)+132)))
	v208 = int64(0)
	v210 = F_GetSysCacheOid(m, int32(81), v206, base.I64_extend_i32_u(v203), v208, v208)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L16
	} else {
		goto L66
	}
L64:
	;
	v213 = int32(0)
	goto L65
L65:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(140))+8))
	*(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[0])) = v217
	goto L67
L66:
	;
	v213 = v210
	goto L65
L67:
	;
	v290 = v213
	goto L57
L68:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[1]))
	if v224 == int32(0) {
		v279 = int32(0)
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v290 = v279
	goto L57
L70:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if int32(0) < v227 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v238 = v6
	goto L74
L72:
	;
	goto L73
L73:
	;
	v279 = int32(0)
	goto L69
L74:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v224)+12))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242+v238<<(uint(int32(2))%32))))
	if l3 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L73
L76:
	;
	v260 = v238 + int32(1)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v260 < v261 {
		v238 = v260
		goto L74
	} else {
		goto L83
	}
L77:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[2]))
	if v246 == v250 {
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v254 = int64(0)
	v256 = F_GetSysCacheOid(m, int32(81), base.I64_extend_i32_u(v219), base.I64_extend_i32_u(v246), v254, v254)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L16
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	if v256 != 0 {
		v279 = v256
		goto L69
	} else {
		goto L82
	}
L82:
	;
	goto L76
L83:
	;
	goto L75
L84:
	;
	v300 = F_get_array_type(m, v290)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L16
	} else {
		goto L85
	}
L85:
	;
	v306 = v300
	goto L9
L86:
	;
	goto L8
L87:
	;
	v485 = v325
	goto L5
L88:
	;
	if v328 == int32(0) {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v332 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if l2 == int32(0) {
		v485 = v328
		goto L5
	} else {
		goto L119
	}
L91:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v463 = v335
	goto L90
L92:
	;
	goto L93
L93:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v328)+16))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+22)))
	v338 = v336 + v337
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+82)))
	if v339 == int32(0) {
		goto L3
	} else {
		goto L94
	}
L94:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v338)+116))
	if v342 == int32(0) {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	v347 = F_palloc_mul(m, int32(8), v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L16
	} else {
		goto L96
	}
L96:
	;
	v349 = int32(0)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v350 == v349 {
		v415 = v349
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v423 = F_construct_array_builtin(m, v347, v415, int32(2275))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L16
	} else {
		goto L113
	}
L98:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	if v353 <= int32(0) {
		v415 = v349
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v360 = v349
	goto L100
L100:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v350)+12))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367+v360<<(uint(int32(2))%32))))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	switch v372 - int32(69) {
	case 0:
		goto L103
	default:
		goto L1
	case 3:
		goto L104
	}
L101:
	;
	v415 = v408
	goto L97
L102:
	;
	if v399 == int32(0) {
		goto L1
	} else {
		goto L111
	}
L103:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	if v386 == int32(0) {
		goto L1
	} else {
		goto L108
	}
L104:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	switch v375 - int32(473) {
	case 0:
		goto L106
	case 1, 3:
		goto L105
	default:
		goto L1
	}
L105:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v371)+8))
	v399 = v385
	goto L102
L106:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v371)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v378
	v383 = F_psprintf(m, int32(_a_F_LookupTypeNameExtended_7), v14+int32(32))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L16
	} else {
		goto L107
	}
L107:
	;
	v399 = v383
	goto L102
L108:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	if v389 != int32(1) {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)))
	if v394 != int32(476) {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	v399 = v397
	goto L102
L111:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v347+v360<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v399)
	v408 = v360 + int32(1)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	if v408 < v409 {
		v360 = v408
		goto L100
	} else {
		goto L112
	}
L112:
	;
	goto L101
L113:
	;
	v426 = v14 + int32(140)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v426)+12)) = int32(524)
	*(*int32)(unsafe.Add(mBase, uint32(v426)+4)) = v427
	*(*int32)(unsafe.Add(mBase, uint32(v426))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v426)+16)) = v426
	v433 = int32(_a_F_LookupTypeNameExtended_6)
	v434 = *(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v426)+8)) = v434
	*(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[0])) = v14 + int32(148)
	goto L114
L114:
	;
	v442 = F_OidFunctionCall1Coll(m, v342, int32(0), base.I64_extend_i32_u(v423))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L16
	} else {
		goto L115
	}
L115:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v426)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[0])) = v445
	goto L116
L116:
	;
	F_pfree(m, v347)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L16
	} else {
		goto L117
	}
L117:
	;
	F_pfree(m, v423)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L16
	} else {
		goto L118
	}
L118:
	;
	v463 = base.I32_wrap_i64(v442)
	goto L90
L119:
	;
	v470 = v463
	v473 = v328
	goto L6
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v306
	F_errmsg_internal(m, int32(_a_F_LookupTypeNameExtended_8), v14)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L16
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_LookupTypeNameExtended_1), int32(209), int32(_a_F_LookupTypeNameExtended_2))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L16
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L16
	} else {
		goto L124
	}
L124:
	;
	v513 = F_TypeNameToString(m, l1)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L16
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v513
	F_errmsg(m, int32(_a_F_LookupTypeNameExtended_9), v14+int32(48))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L16
	} else {
		goto L126
	}
L126:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v521)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L16
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_LookupTypeNameExtended_1), int32(356), int32(_a_F_LookupTypeNameExtended_10))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L16
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L16
	} else {
		goto L130
	}
L130:
	;
	v536 = F_TypeNameToString(m, l1)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L16
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v536
	F_errmsg(m, int32(_a_F_LookupTypeNameExtended_11), v14+int32(16))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L16
	} else {
		goto L132
	}
L132:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v544)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L16
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_LookupTypeNameExtended_1), int32(365), int32(_a_F_LookupTypeNameExtended_10))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L16
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L16
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_LookupTypeNameExtended_12), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L16
	} else {
		goto L137
	}
L137:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v564)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L16
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_LookupTypeNameExtended_1), int32(410), int32(_a_F_LookupTypeNameExtended_10))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L16
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_TypeCacheConstrCallback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheConstrCallback[0]))
	if v5 != 0 {
		v7 = v5
		for {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+312))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+312)) = v9 & int32(-524289)
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+320))
			if v13 != 0 {
				v7 = v13
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_TypeCacheRelCallback(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l1
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheRelCallback[0]))
	v16 = int32(0)
	v18 = F_hash_search(m, v13, v9+int32(24), v16, v16)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	v96 = v9 + int32(4)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheRelCallback[1]))
	F_hash_seq_init(m, v96, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L29
	}
L5:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheRelCallback[2]))
	if v79 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L6:
	;
	return
L7:
	;
	if v18 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheRelCallback[1]))
	v26 = int32(0)
	v28 = F_hash_search(m, v23, v18+int32(4), v26, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v28 == int32(0) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+188))
	if v32 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+13)))
	if v58&int32(1)|base.B2i32(v62 != int32(99)) != 0 {
		goto L5
	} else {
		goto L20
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v35 = v33 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v35
	if v35 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v28)+312))
	v52 = v50 & int32(_a_F_TypeCacheRelCallback_0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v52
	if v50&int32(-1572866) == int32(0) {
		goto L5
	} else {
		goto L19
	}
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+188))
	F_FreeTupleDesc(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+192)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+188)) = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v28)+312))
	v48 = v46 & int32(_a_F_TypeCacheRelCallback_0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+312)) = v48
	v58 = v48
	goto L11
L18:
	;
	goto L17
L19:
	;
	v58 = v52
	goto L11
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheRelCallback[0]))
	v73 = F_hash_search(m, v67, v28+int32(16), int32(2), v9+int32(4))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	goto L5
L22:
	;
	v83 = v79
	goto L23
L23:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+312))
	if v88&int32(_a_F_TypeCacheRelCallback_1) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L1
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+312)) = v88 & int32(_a_F_TypeCacheRelCallback_0)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v83)+320))
	if v94 != 0 {
		v83 = v94
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	v101 = F_hash_seq_search(m, v96)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	if v101 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v106 = v101
	goto L32
L32:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+13)))
	switch v111 - int32(99) {
	case 0:
		goto L36
	case 1:
		goto L35
	default:
		goto L34
	}
L33:
	;
	goto L1
L34:
	;
	v174 = F_hash_seq_search(m, v9+int32(4))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L6
	} else {
		goto L49
	}
L35:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v106)+312))
	if v161&int32(_a_F_TypeCacheRelCallback_1) == int32(0) {
		goto L34
	} else {
		goto L48
	}
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v106)+188))
	if v114 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if base.B2i32(v144 == int32(0))|v145&int32(1) != 0 {
		goto L34
	} else {
		goto L46
	}
L38:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v117 = v115 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v114)+12)) = v117
	if v117 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v106)+312))
	v137 = v135 & int32(_a_F_TypeCacheRelCallback_0)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+312)) = v137
	if v135&int32(-1572866) == int32(0) {
		goto L34
	} else {
		goto L45
	}
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v106)+188))
	F_FreeTupleDesc(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v106)+192)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+188)) = int32(0)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v106)+312))
	v130 = v128 & int32(_a_F_TypeCacheRelCallback_0)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+312)) = v130
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+13)))
	v144 = base.B2i32(v132 == int32(99))
	v145 = v130
	goto L37
L44:
	;
	goto L43
L45:
	;
	v144 = int32(1)
	v145 = v137
	goto L37
L46:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheRelCallback[0]))
	v159 = F_hash_search(m, v153, v106+int32(16), int32(2), v9+int32(31))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	goto L34
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+312)) = v161 & int32(_a_F_TypeCacheRelCallback_0)
	goto L34
L49:
	;
	if v174 != 0 {
		v106 = v174
		goto L32
	} else {
		goto L50
	}
L50:
	;
	goto L33
}
func F_TypeCacheTypCallback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheTypCallback[0]))
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v23 = F_hash_seq_search(m, v7+int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	F_hash_seq_init(m, v7+int32(8), v10)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_hash_seq_init_with_hash_value(m, v7+int32(8), v10, l2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return
L6:
	;
	goto L1
L7:
	;
	goto L1
L8:
	;
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = v23
	goto L12
L10:
	;
	goto L11
L11:
	;
	m.G0 = v7 + int32(32)
	return
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v29 & int32(-524290)
	if base.B2i32(v29&int32(1) == int32(0))|v29&int32(-1572866) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v55 = F_hash_seq_search(m, v7+int32(8))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L19
	}
L15:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+13)))
	if v40 != int32(99) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+188))
	if v43 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheTypCallback[1]))
	v51 = F_hash_search(m, v45, v27+int32(16), int32(2), v7+int32(31))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if v55 != 0 {
		v27 = v55
		goto L12
	} else {
		goto L20
	}
L20:
	;
	goto L13
}
func F_TypeCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32, l21 int32, l22 int32, l23 int32, l24 int32, l25 int32, l26 int32, l27 int32, l28 int32, l29 int32, l30 int32, l31 int32, l32 int32) {
	mBase := m.M
	_ = mBase
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v213 int64
	_ = v213
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v284 int64
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	v34 = int32(0)
	v38 = m.G0
	v40 = v38 - int32(528)
	m.G0 = v40
	if base.B2i32(l7 <= v34)&base.B2i32(base.Ui32(l7) <= base.Ui32(int32(-3))) == v34 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L21
	} else {
		goto L131
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L21
	} else {
		goto L128
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L21
	} else {
		goto L124
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L21
	} else {
		goto L120
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L21
	} else {
		goto L116
	}
L6:
	;
	if l26 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L21
	} else {
		goto L112
	}
L9:
	;
	if base.B2i32(l8 == int32(109))|l21 != 0 {
		goto L57
	} else {
		goto L58
	}
L10:
	;
	if l7 == int32(-1) {
		goto L9
	} else {
		goto L55
	}
L11:
	;
	v50 = l7 - int32(1)
	v52 = int32(_a_F_TypeCreate_0)
	if base.Ui32((l7^v50)&v52) <= base.Ui32(v50&v52) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if l7 == int32(-1) {
		goto L45
	} else {
		goto L46
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L21
	} else {
		goto L41
	}
L15:
	;
	switch base.I32_ctz(l7) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	default:
		goto L14
	}
L16:
	;
	if l27 == int32(100) {
		goto L10
	} else {
		goto L36
	}
L17:
	;
	if l27 == int32(105) {
		goto L10
	} else {
		goto L31
	}
L18:
	;
	if l27 == int32(115) {
		goto L10
	} else {
		goto L26
	}
L19:
	;
	if l27 == int32(99) {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return
L22:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = l27
	F_errmsg(m, int32(_a_F_TypeCreate_8), v40+int32(32))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(273), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+52)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = l27
	F_errmsg(m, int32(_a_F_TypeCreate_8), v40+int32(48))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L21
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(281), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L21
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L21
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = l27
	F_errmsg(m, int32(_a_F_TypeCreate_8), v40-int32(-64))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L21
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(289), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L21
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+80)) = l27
	F_errmsg(m, int32(_a_F_TypeCreate_8), v40+int32(80))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L21
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(297), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L21
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L21
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = l7
	F_errmsg(m, int32(_a_F_TypeCreate_1), v40)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L21
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(303), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	switch l27 - int32(100) {
	case 0, 5:
		goto L9
	default:
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if l7 != int32(-2) {
		goto L10
	} else {
		goto L53
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L21
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+96)) = l27
	F_errmsg(m, int32(_a_F_TypeCreate_9), v40+int32(96))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L21
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(313), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L21
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	if l27 != int32(99) {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	goto L10
L55:
	;
	if l28 != int32(112) {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	goto L9
L57:
	;
	v204 = int32(1)
	goto L59
L58:
	;
	v204 = base.B2i32(l4 != int32(0)) & base.B2i32(l5 != int32(99))
	goto L59
L59:
	;
	v205 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+520)) = v205
	*(*int64)(unsafe.Add(mBase, uint32(v40)+512)) = v205
	*(*int64)(unsafe.Add(mBase, uint32(v40)+504)) = v205
	*(*int64)(unsafe.Add(mBase, uint32(v40)+496)) = v205
	v213 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+464)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v40)+472)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v40)+480)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v40)+488)) = v213
	v223 = int32(0)
	base.MemoryFill(m, v40+int32(208), v223, int32(256))
	v227 = v40 + int32(144)
	v229 = F_strncpy(m, v227, l2, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v229)+63)) = uint8(v223)
	goto L60
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+432)) = base.I64_extend_i32_u(l32)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+424)) = base.I64_extend_i32_s(l30)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+416)) = base.I64_extend_i32_s(l29)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+408)) = base.I64_extend_i32_u(l23)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+400)) = base.I64_extend_i32_u(l31)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+392)) = base.I64_extend_i32_s(l28)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+384)) = base.I64_extend_i32_s(l27)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+376)) = base.I64_extend_i32_u(l18)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+368)) = base.I64_extend_i32_u(l17)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+360)) = base.I64_extend_i32_u(l16)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+352)) = base.I64_extend_i32_u(l15)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+344)) = base.I64_extend_i32_u(l14)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+336)) = base.I64_extend_i32_u(l13)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+328)) = base.I64_extend_i32_u(l12)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+320)) = base.I64_extend_i32_u(l22)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+312)) = base.I64_extend_i32_u(l20)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+304)) = base.I64_extend_i32_u(l19)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+296)) = base.I64_extend_i32_u(l4)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+288)) = base.I64_extend_i32_s(l11)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+280)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+272)) = base.I64_extend_i32_u(l10)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+264)) = base.I64_extend_i32_s(l9)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+256)) = base.I64_extend_i32_s(l8)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+248)) = base.I64_extend_i32_u(l26)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+240)) = base.I64_extend_i32_s(l7)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+232)) = base.I64_extend_i32_u(l6)
	v284 = base.I64_extend_i32_u(l3)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+224)) = v284
	*(*int64)(unsafe.Add(mBase, uint32(v40)+216)) = base.I64_extend_i32_u(v227)
	if l25 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if l24 != 0 {
		goto L67
	} else {
		goto L68
	}
L62:
	;
	v288 = F_cstring_to_text(m, l25)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L21
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v292 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+525)) = uint8(v292)
	goto L61
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+440)) = base.I64_extend_i32_u(v288)
	goto L61
L66:
	;
	if v204 != 0 {
		goto L72
	} else {
		goto L73
	}
L67:
	;
	v294 = F_cstring_to_text(m, l24)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L21
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v298 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+526)) = uint8(v298)
	goto L66
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+448)) = base.I64_extend_i32_u(v294)
	goto L66
L71:
	;
	v314 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L21
	} else {
		goto L76
	}
L72:
	;
	v308 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+527)) = uint8(v308)
	v311 = int32(0)
	goto L71
L73:
	;
	v301 = F_get_user_default_acl(m, int32(50), l6, l3)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L21
	} else {
		goto L74
	}
L74:
	;
	if v301 == int32(0) {
		goto L72
	} else {
		goto L75
	}
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+456)) = base.I64_extend_i32_u(v301)
	v311 = v301
	goto L71
L76:
	;
	v318 = F_SearchSysCacheCopy(m, int32(81), base.I64_extend_i32_u(l2), v284)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L21
	} else {
		goto L78
	}
L77:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCreate[0]))
	if v379 != 0 {
		goto L99
	} else {
		goto L100
	}
L78:
	;
	if v318 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v318)+16))
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+22)))
	v322 = v320 + v321
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322)+82)))
	if v323 == int32(1) {
		goto L3
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if l1 != 0 {
		v363 = l1
		goto L90
	} else {
		goto L91
	}
L82:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v322)+72))
	if l6 != v326 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	F_aclcheck_error(m, int32(2), int32(50), l2)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L21
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if l1 != 0 {
		goto L2
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	v332 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+464)) = uint8(v332)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v314)+52))
	v341 = F_heap_modify_tuple(m, v318, v334, v40+int32(208), v40+int32(496), v40+int32(464))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L21
	} else {
		goto L88
	}
L88:
	;
	F_CatalogTupleUpdate(m, v314, v341+int32(4), v341)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L21
	} else {
		goto L89
	}
L89:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v375 = v347
	v376 = v341
	goto L77
L90:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+208)) = base.I64_extend_i32_u(v363)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v314)+52))
	v371 = F_heap_form_tuple(m, v366, v40+int32(208), v40+int32(496))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L21
	} else {
		goto L97
	}
L91:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TypeCreate[2])))
	if v349 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCreate[3]))
	if v353 == int32(0) {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v361 = F_GetNewOidWithIndex(m, v314, int32(2703), int32(1))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L21
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TypeCreate[3])) = int32(0)
	v363 = v353
	goto L90
L96:
	;
	v363 = v361
	goto L90
L97:
	;
	F_CatalogTupleInsert(m, v314, v371)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L21
	} else {
		goto L98
	}
L98:
	;
	v375 = v363
	v376 = v371
	goto L77
L99:
	;
	if l25 != 0 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCreate[1]))
	if v390 != 0 {
		goto L107
	} else {
		goto L108
	}
L102:
	;
	v380 = F_stringToNode(m, l25)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L21
	} else {
		goto L105
	}
L103:
	;
	v383 = int32(0)
	goto L104
L104:
	;
	F_GenerateTypeDependencies(m, v376, v314, v383, v311, l5, l21, v204, int32(1), base.B2i32(v318 != int32(0)))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L21
	} else {
		goto L106
	}
L105:
	;
	v383 = v380
	goto L104
L106:
	;
	goto L101
L107:
	;
	v392 = int32(0)
	F_RunObjectPostCreateHook(m, int32(1247), v375, v392, v392)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L21
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
	F_relation_close(m, v314, int32(3))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L21
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	m.G0 = v40 + int32(528)
	return
L112:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L21
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+128)) = l7
	F_errmsg(m, int32(_a_F_TypeCreate_10), v40+int32(128))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L21
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(257), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L21
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L21
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+112)) = l27
	F_errmsg(m, int32(_a_F_TypeCreate_9), v40+int32(112))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L21
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(319), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L21
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L21
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(_a_F_TypeCreate_11), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L21
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(326), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L21
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(_a_F_TypeCreate_4))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L21
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = l2
	F_errmsg(m, int32(_a_F_TypeCreate_5), v40+int32(16))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L21
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(435), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L21
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errmsg_internal(m, int32(_a_F_TypeCreate_6), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L21
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(445), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L21
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L21
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(_a_F_TypeCreate_7), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L21
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(475), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L21
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_coerce_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v600 int32
	_ = v600
	v9 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(32)
	m.G0 = v23
	if l1 == v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v23 + int32(32)
	return v600
L2:
	;
	v600 = l1
	goto L1
L3:
	;
	goto L4
L4:
	;
	if l2 == l3 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v600 = l1
	goto L1
L6:
	;
	goto L7
L7:
	;
	switch l3 - int32(2276) {
	case 0, 7:
		v600 = l1
		goto L1
	case 1, 2, 3, 4, 5, 6:
		goto L8
	default:
		goto L9
	}
L8:
	;
	v54 = int32(0)
	if base.B2i32(base.B2i32(l3 == int32(2277))|base.B2i32(l3 == int32(3500))|base.B2i32(l3 == int32(3831))|base.B2i32(l3 == int32(_a_F_coerce_type_0))|base.B2i32(l3 == int32(_a_F_coerce_type_1))|base.B2i32(l3 == int32(_a_F_coerce_type_2))|base.B2i32(l3 == int32(_a_F_coerce_type_3)) == v54)|base.B2i32(l2 == int32(705)) == v54 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	switch l3 - int32(_a_F_coerce_type_4) {
	case 0, 2:
		v600 = l1
		goto L1
	case 1:
		goto L8
	default:
		goto L10
	}
L10:
	;
	if l3 != int32(2776) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v600 = l1
	goto L1
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v590)+24)) = l7
	v600 = v590
	goto L1
L13:
	;
	v586 = F_makeRelabelType(m, l1, v61, int32(-1), int32(0), l6)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L17
	} else {
		goto L188
	}
L14:
	;
	v61 = F_getBaseType(m, l2)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(l2 != int32(705))|base.B2i32(v68 != int32(7)) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	return int32(0)
L18:
	;
	if v61 != l2 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v600 = l1
	goto L1
L20:
	;
	v75 = F_palloc0(m, int32(40))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L17
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if base.B2i32(l0 == int32(0))|base.B2i32(v68 != int32(8)) != 0 {
		goto L48
	} else {
		goto L49
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = l4
	v82 = F_getBaseTypeAndTypmod(m, l3, v23+int32(28))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v85 = F_typeidType(m, v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	if v82 != int32(1186) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v90 = int32(-1)
	goto L28
L27:
	;
	v90 = v84
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = v82
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+22)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+22)))
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98+v99)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+22)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+v104)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+33)) = uint8(v106)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+32)) = uint8(v108)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+36)) = v110
	v113 = v23 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = int32(524)
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v113)+16)) = v113
	v119 = int32(_a_F_coerce_type_5)
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_coerce_type[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v120
	*(*int32)(unsafe.Add(mBase, _c_F_coerce_type[0])) = v23 + int32(16)
	goto L29
L29:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+22)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126+v127)+100))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v130 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v133 = int32(0)
	goto L32
L31:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v133 = v132
	goto L32
L32:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+22)))
	v136 = v134 + v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+92))
	if v137 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v139 = v137
	goto L35
L34:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v139 = v138
	goto L35
L35:
	;
	v140 = F_OidInputFunctionCall(m, v129, v133, v139, v90)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L17
	} else {
		goto L36
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v75)+24)) = v140
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v143 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(8))+8))
	*(*int32)(unsafe.Add(mBase, _c_F_coerce_type[0])) = v155
	goto L41
L38:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	if v144 != int32(-1) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v148 = F_pg_detoast_datum(m, base.I32_wrap_i64(v140))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L17
	} else {
		goto L40
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v75)+24)) = base.I64_extend_i32_u(v148)
	goto L37
L41:
	;
	if l3 != v82 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v160 = F_coerce_to_domain(m, v75, v82, v158, l3, l5, l6, l7, int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L17
	} else {
		goto L45
	}
L43:
	;
	v162 = v75
	goto L44
L44:
	;
	F_ReleaseCatCache(m, v85)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L17
	} else {
		goto L46
	}
L45:
	;
	v162 = v160
	goto L44
L46:
	;
	v600 = v162
	goto L1
L47:
	;
	v200 = F_find_coercion_pathway(m, l3, l2, l5, v23+int32(8))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L17
	} else {
		goto L61
	}
L48:
	;
	v177 = v68
	goto L50
L49:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v170 == int32(0) {
		goto L47
	} else {
		goto L51
	}
L50:
	;
	if v177 != int32(31) {
		goto L47
	} else {
		goto L54
	}
L51:
	;
	v173 = m.T0[v170].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l3, l4, l7)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	if v173 != 0 {
		v600 = v173
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v177 = v175
	goto L50
L54:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v181 = F_coerce_type(m, l0, v180, l2, l3, l4, l5, l6, l7)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L17
	} else {
		goto L55
	}
L55:
	;
	v183 = F_type_is_collatable(m, l3)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L17
	} else {
		goto L56
	}
L56:
	;
	if v183 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v600 = v181
	goto L1
L58:
	;
	goto L59
L59:
	;
	v188 = F_palloc0(m, int32(16))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L17
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+4)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = int32(31)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+8)) = v193
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = v195
	v600 = v188
	goto L1
L61:
	;
	if v200 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = l4
	v205 = F_getBaseTypeAndTypmod(m, l3, v23+int32(28))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L17
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if l2 != int32(2249) {
		goto L75
	} else {
		goto L76
	}
L65:
	;
	if v200 != int32(2) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v211 = F_build_coercion_expression(m, l1, v200, v209, v205, v210, l5, l6, l7)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L17
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v220 = F_coerce_to_domain(m, l1, v205, v218, l3, l5, l6, l7, int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L17
	} else {
		goto L72
	}
L69:
	;
	if v205 == l3 {
		v600 = v211
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v216 = F_coerce_to_domain(m, v211, v205, v214, l3, l5, l6, l7, int32(1))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L17
	} else {
		goto L71
	}
L71:
	;
	v600 = v216
	goto L1
L72:
	;
	if v220 != l1 {
		v600 = v220
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v225 = F_makeRelabelType(m, v220, l3, int32(-1), int32(0), l6)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L17
	} else {
		goto L74
	}
L74:
	;
	v590 = v225
	goto L12
L75:
	;
	if l3 != int32(2287) {
		goto L162
	} else {
		goto L163
	}
L76:
	;
	v229 = F_typeOrDomainTypeRelid(m, l3)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L17
	} else {
		goto L77
	}
L77:
	;
	if v229 == int32(0) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v233 = m.G0
	v235 = v233 - int32(80)
	m.G0 = v235
	*(*int32)(unsafe.Add(mBase, uint32(v235)+76)) = int32(-1)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v239 != int32(6) {
		goto L85
	} else {
		goto L86
	}
L79:
	;
	v600 = v400
	goto L1
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L17
	} else {
		goto L153
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L17
	} else {
		goto L143
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L17
	} else {
		goto L135
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L17
	} else {
		goto L128
	}
L84:
	;
	v256 = F_getBaseTypeAndTypmod(m, l3, v235+int32(76))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L17
	} else {
		goto L92
	}
L85:
	;
	if v239 != int32(36) {
		goto L83
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	if v245 != 0 {
		goto L83
	} else {
		goto L89
	}
L88:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v253 = v244
	goto L84
L89:
	;
	v246 = F_GetNSItemByVar(m, l0, l1)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L17
	} else {
		goto L90
	}
L90:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v251 = F_expandNSItemVars(m, l0, v246, v248, v249, int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L17
	} else {
		goto L91
	}
L91:
	;
	v253 = v251
	goto L84
L92:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v235)+76))
	v259 = F_lookup_rowtype_tupdesc(m, v256, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L17
	} else {
		goto L93
	}
L93:
	;
	if v253 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	v262 = v261
	goto L96
L95:
	;
	v262 = v9
	goto L96
L96:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if int32(0) < v263 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v270 = int32(0)
	v279 = v262
	v280 = v263
	v282 = v9
	v284 = int32(1)
	goto L100
L98:
	;
	v355 = v262
	v358 = v9
	goto L99
L99:
	;
	if v355 != 0 {
		goto L80
	} else {
		goto L117
	}
L100:
	;
	v293 = v259 + v280<<(uint(int32(3))%32) + v270*int32(100)
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+119)))
	if v294 == int32(1) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v355 = v333
	v358 = v336
	goto L99
L102:
	;
	v341 = v270 + int32(1)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if v341 < v342 {
		v270 = v341
		v279 = v333
		v280 = v342
		v282 = v336
		v284 = v337
		goto L100
	} else {
		goto L116
	}
L103:
	;
	v300 = F_makeNullConst(m, int32(23), int32(-1), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L17
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if v279 == int32(0) {
		goto L82
	} else {
		goto L108
	}
L106:
	;
	v302 = F_lappend(m, v282, v300)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L17
	} else {
		goto L107
	}
L107:
	;
	v333 = v279
	v336 = v302
	v337 = v284
	goto L102
L108:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v307 = F_exprType(m, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L17
	} else {
		goto L109
	}
L109:
	;
	v310 = v293 + int32(28)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+68))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v310)+76))
	v315 = F_coerce_to_target_type(m, l0, v306, v307, v311, v312, l5, int32(2), int32(-1))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L17
	} else {
		goto L110
	}
L110:
	;
	if v315 == int32(0) {
		goto L81
	} else {
		goto L111
	}
L111:
	;
	v319 = F_lappend(m, v282, v315)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L17
	} else {
		goto L112
	}
L112:
	;
	v322 = v279 + int32(4)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	if base.Ui32(v322) < base.Ui32(v324+v325<<(uint(int32(2))%32)) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v330 = v322
	goto L115
L114:
	;
	v330 = int32(0)
	goto L115
L115:
	;
	v333 = v330
	v336 = v319
	v337 = v284 + int32(1)
	goto L102
L116:
	;
	goto L101
L117:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	if int32(0) <= v364 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_DecrTupleDescRefCount(m, v259)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L17
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v370 = F_palloc0(m, int32(24))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L17
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370)+20)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v370)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v370)+12)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v370)+8)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v370)+4)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v370))) = int32(36)
	if l3 != v256 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v381 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v370)+12)) = v381
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v235)+76))
	v386 = F_coerce_type_typmod(m, v370, v256, v383, l5, v381, l7, int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L17
	} else {
		goto L126
	}
L124:
	;
	v400 = v370
	goto L125
L125:
	;
	m.G0 = v235 + int32(80)
	goto L79
L126:
	;
	v389 = F_palloc0(m, int32(28))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L17
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v389)+24)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v389)+20)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v389)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v389)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v389)+4)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = int32(55)
	v400 = v389
	goto L125
L128:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L17
	} else {
		goto L129
	}
L129:
	;
	v412 = F_format_type_be(m, int32(2249))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L17
	} else {
		goto L130
	}
L130:
	;
	v414 = F_format_type_be(m, l3)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L17
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235)+4)) = v414
	*(*int32)(unsafe.Add(mBase, uint32(v235))) = v412
	F_errmsg(m, int32(_a_F_coerce_type_6), v235)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L17
	} else {
		goto L132
	}
L132:
	;
	F_parser_coercion_errposition(m, l0, l7, l1)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L17
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_coerce_type_7), int32(1055), int32(_a_F_coerce_type_8))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L17
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L17
	} else {
		goto L136
	}
L136:
	;
	v436 = F_format_type_be(m, int32(2249))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L17
	} else {
		goto L137
	}
L137:
	;
	v438 = F_format_type_be(m, l3)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L17
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235)+36)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v235)+32)) = v436
	F_errmsg(m, int32(_a_F_coerce_type_6), v235+int32(32))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L17
	} else {
		goto L139
	}
L139:
	;
	v449 = F_errdetail(m, int32(_a_F_coerce_type_9), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L17
	} else {
		goto L140
	}
L140:
	;
	F_parser_coercion_errposition(m, l0, l7, l1)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L17
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_coerce_type_7), int32(1094), int32(_a_F_coerce_type_8))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L17
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L17
	} else {
		goto L144
	}
L144:
	;
	v466 = F_format_type_be(m, int32(2249))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L17
	} else {
		goto L145
	}
L145:
	;
	v468 = F_format_type_be(m, l3)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L17
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235)+68)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v235)+64)) = v466
	F_errmsg(m, int32(_a_F_coerce_type_6), v235-int32(-64))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L17
	} else {
		goto L147
	}
L147:
	;
	v477 = F_format_type_be(m, v307)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L17
	} else {
		goto L148
	}
L148:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v310)+68))
	v480 = F_format_type_be(m, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L17
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235)+56)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v235)+52)) = v480
	*(*int32)(unsafe.Add(mBase, uint32(v235)+48)) = v477
	v488 = F_errdetail(m, int32(_a_F_coerce_type_10), v235+int32(48))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L17
	} else {
		goto L150
	}
L150:
	;
	F_parser_coercion_errposition(m, l0, l7, v306)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L17
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_coerce_type_7), int32(1115), int32(_a_F_coerce_type_8))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L17
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L17
	} else {
		goto L154
	}
L154:
	;
	v505 = F_format_type_be(m, int32(2249))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L17
	} else {
		goto L155
	}
L155:
	;
	v507 = F_format_type_be(m, l3)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L17
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235)+20)) = v507
	*(*int32)(unsafe.Add(mBase, uint32(v235)+16)) = v505
	F_errmsg(m, int32(_a_F_coerce_type_6), v235+int32(16))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L17
	} else {
		goto L157
	}
L157:
	;
	v518 = F_errdetail(m, int32(_a_F_coerce_type_11), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L17
	} else {
		goto L158
	}
L158:
	;
	F_parser_coercion_errposition(m, l0, l7, l1)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L17
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_coerce_type_7), int32(1127), int32(_a_F_coerce_type_8))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L17
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	v539 = F_typeInheritsFrom(m, l2, l3)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L17
	} else {
		goto L171
	}
L162:
	;
	if l3 != int32(2249) {
		goto L161
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v535 = F_is_complex_array(m, l2)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L17
	} else {
		goto L168
	}
L165:
	;
	v531 = F_typeOrDomainTypeRelid(m, l2)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L17
	} else {
		goto L166
	}
L166:
	;
	if v531 == int32(0) {
		goto L161
	} else {
		goto L167
	}
L167:
	;
	v600 = l1
	goto L1
L168:
	;
	if v535 == int32(0) {
		goto L161
	} else {
		goto L169
	}
L169:
	;
	v600 = l1
	goto L1
L170:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L17
	} else {
		goto L183
	}
L171:
	;
	if v539 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v543 = F_typeIsOfTypedTable(m, l2, l3)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L17
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v547 = F_getBaseType(m, l2)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L17
	} else {
		goto L177
	}
L175:
	;
	if v543 == int32(0) {
		goto L170
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v550 = F_palloc0(m, int32(20))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L17
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = int32(30)
	if v547 != l2 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v558 = F_makeRelabelType(m, l1, v547, int32(-1), int32(0), int32(2))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L17
	} else {
		goto L182
	}
L180:
	;
	v561 = l1
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550)+16)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v550)+12)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v550)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v550)+4)) = v561
	v600 = v550
	goto L1
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v558)+24)) = l7
	v561 = v558
	goto L181
L183:
	;
	v570 = F_format_type_be(m, l2)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L17
	} else {
		goto L184
	}
L184:
	;
	v572 = F_format_type_be(m, l3)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L17
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v572
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v570
	F_errmsg_internal(m, int32(_a_F_coerce_type_12), v23)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L17
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_coerce_type_7), int32(545), int32(_a_F_coerce_type_13))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L17
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	v590 = v586
	goto L12
}
func F_coerce_type_typmod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
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
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v11 = F_exprTypmod(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 != l2 {
			if l6 != 0 {
				F_hide_coercion_node(m, l0)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					if l2 < int32(0) {
						v62 = F_exprCollation(m, l0)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v65 = F_applyRelabelType(m, l0, l1, l2, v62, l4, l5, int32(0))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								v71 = v65
								return v71
							}
						}
					} else {
						v20 = F_typeidType(m, l1)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
							v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
							v24 = v22 + v23
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
							if v25 == int32(0) {
								v36 = l1
								v38 = int32(1)
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
								v31 = base.B2i32(v29 == int32(_a_F_coerce_type_typmod_0))
								if v29 == int32(_a_F_coerce_type_typmod_0) {
									v32 = v25
								} else {
									v32 = l1
								}
								if v29 == int32(_a_F_coerce_type_typmod_0) {
									v35 = int32(3)
								} else {
									v35 = int32(1)
								}
								v36 = v32
								v38 = v35
							}
							F_ReleaseCatCache(m, v20)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v42 = base.I64_extend_i32_u(v36)
								v43 = F_SearchSysCache2(m, int32(12), v42, v42)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									if v43 == int32(0) {
										v62 = F_exprCollation(m, l0)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v65 = F_applyRelabelType(m, l0, l1, l2, v62, l4, l5, int32(0))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v71 = v65
												return v71
											}
										}
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
										v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+22)))
										v50 = *(*int32)(unsafe.Add(mBase, uint32(v47+v48)+12))
										F_ReleaseCatCache(m, v43)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											if v50 == int32(0) {
												v62 = F_exprCollation(m, l0)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int32(0)
												} else {
													v65 = F_applyRelabelType(m, l0, l1, l2, v62, l4, l5, int32(0))
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return int32(0)
													} else {
														v71 = v65
														return v71
													}
												}
											} else {
												v55 = F_build_coercion_expression(m, l0, v38, v50, l1, l2, l3, l4, l5)
												mBase = m.M
												v56 = m.ExcPending
												if v56 != 0 {
													return int32(0)
												} else {
													return v55
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				if l2 < int32(0) {
					v62 = F_exprCollation(m, l0)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v65 = F_applyRelabelType(m, l0, l1, l2, v62, l4, l5, int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v71 = v65
							return v71
						}
					}
				} else {
					v20 = F_typeidType(m, l1)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
						v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
						v24 = v22 + v23
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
						if v25 == int32(0) {
							v36 = l1
							v38 = int32(1)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+88))
							v31 = base.B2i32(v29 == int32(_a_F_coerce_type_typmod_0))
							if v29 == int32(_a_F_coerce_type_typmod_0) {
								v32 = v25
							} else {
								v32 = l1
							}
							if v29 == int32(_a_F_coerce_type_typmod_0) {
								v35 = int32(3)
							} else {
								v35 = int32(1)
							}
							v36 = v32
							v38 = v35
						}
						F_ReleaseCatCache(m, v20)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v42 = base.I64_extend_i32_u(v36)
							v43 = F_SearchSysCache2(m, int32(12), v42, v42)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								if v43 == int32(0) {
									v62 = F_exprCollation(m, l0)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v65 = F_applyRelabelType(m, l0, l1, l2, v62, l4, l5, int32(0))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											v71 = v65
											return v71
										}
									}
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
									v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+22)))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v47+v48)+12))
									F_ReleaseCatCache(m, v43)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										if v50 == int32(0) {
											v62 = F_exprCollation(m, l0)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												v65 = F_applyRelabelType(m, l0, l1, l2, v62, l4, l5, int32(0))
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													v71 = v65
													return v71
												}
											}
										} else {
											v55 = F_build_coercion_expression(m, l0, v38, v50, l1, l2, l3, l4, l5)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												return v55
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v71 = l0
			return v71
		}
	}
}
func F_fillTypeDesc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v12 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(_a_F_fillTypeDesc_0), v7)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_fillTypeDesc_1), int32(175), int32(_a_F_fillTypeDesc_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v31 = v29 + v30
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+76)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v32)
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+78)))
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v34)
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+128)))
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)) = uint8(v36)
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+129)))
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v38)
			F_ReleaseCatCache(m, v12)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_findTypeReceiveFunction(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+60)) = int32(23)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+52)) = int64(111669151977)
	v14 = int32(1)
	v16 = v6 + int32(-12)
	v18 = F_LookupFuncName(m, l0, v14, v16, v14)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v24 = F_LookupFuncName(m, l0, int32(3), v16, int32(1))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				if v24 == int32(0) {
					v50 = v18
					v51 = F_get_func_rettype(m, v50)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						if v51 != l1 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									v113 = F_NameListToString(m, l0)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										v115 = F_format_type_be(m, l1)
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v115
											*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v113
											F_errmsg(m, int32(_a_F_findTypeReceiveFunction_0), v6+int32(-32))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_findTypeReceiveFunction_1), int32(2174), int32(_a_F_findTypeReceiveFunction_2))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							}
						} else {
							v54 = F_func_volatile(m, v50)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								if v54 != int32(118) {
									m.G0 = v8 - int32(-64)
									return v50
								} else {
									v60 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										if v60 == int32(0) {
											m.G0 = v8 - int32(-64)
											return v50
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v67 = F_NameListToString(m, l0)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v67
													F_errmsg(m, int32(_a_F_findTypeReceiveFunction_3), v6+int32(-48))
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_findTypeReceiveFunction_1), int32(2181), int32(_a_F_findTypeReceiveFunction_2))
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 - int32(-64)
															return v50
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(84439172))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = F_NameListToString(m, l0)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v35
								F_errmsg(m, int32(_a_F_findTypeReceiveFunction_4), v6+int32(-16))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_findTypeReceiveFunction_1), int32(2156), int32(_a_F_findTypeReceiveFunction_2))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			} else {
				if v24 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							v95 = F_func_signature_string(m, l0, int32(1), int32(0), v6+int32(-12))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v95
								F_errmsg(m, int32(_a_F_findTypeReceiveFunction_5), v8)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_findTypeReceiveFunction_1), int32(2166), int32(_a_F_findTypeReceiveFunction_2))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				} else {
					v50 = v24
					v51 = F_get_func_rettype(m, v50)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						if v51 != l1 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									v113 = F_NameListToString(m, l0)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										v115 = F_format_type_be(m, l1)
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v115
											*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v113
											F_errmsg(m, int32(_a_F_findTypeReceiveFunction_0), v6+int32(-32))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_findTypeReceiveFunction_1), int32(2174), int32(_a_F_findTypeReceiveFunction_2))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							}
						} else {
							v54 = F_func_volatile(m, v50)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								if v54 != int32(118) {
									m.G0 = v8 - int32(-64)
									return v50
								} else {
									v60 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										if v60 == int32(0) {
											m.G0 = v8 - int32(-64)
											return v50
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v67 = F_NameListToString(m, l0)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v67
													F_errmsg(m, int32(_a_F_findTypeReceiveFunction_3), v6+int32(-48))
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_findTypeReceiveFunction_1), int32(2181), int32(_a_F_findTypeReceiveFunction_2))
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 - int32(-64)
															return v50
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_format_type(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v4 == int32(1) {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v7)
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v12 != 0 {
			v17 = int32(-1)
			v18 = int32(2)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v17 = v15
			v18 = int32(3)
		}
		v19 = F_format_type_extended(m, v11, v17, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			v23 = F_cstring_to_text(m, v19)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v23)
			}
		}
	}
}
func F_format_type_be(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_format_type_extended(m, l0, int32(-1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_lookup_type_cache(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int64
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v995 int32
	_ = v995
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1227 int32
	_ = v1227
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+124)) = l0
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[0]))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[1]))
	if v17 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v85 = v13
	goto L3
L3:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[5]))
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[4]))
	if v89 <= v87 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+88)) = int32(1812)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = int64(1408749273092)
	v27 = int32(72)
	v30 = F_hash_create(m, int32(_a_F_lookup_type_cache_0), int64(64), v9+v27, v27)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[2]))
	if v36 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return int32(0)
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[1])) = v30
	goto L6
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = int64(34359738372)
	v47 = F_hash_create(m, int32(_a_F_lookup_type_cache_1), int64(64), v9+int32(72), int32(40))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[3]))
	if v51 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[2])) = v47
	goto L11
L13:
	;
	v56 = v51
	goto L15
L14:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L16
	}
L15:
	;
	v58 = F_MemoryContextAlloc(m, v56, int32(16))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[3]))
	v56 = v55
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[4])) = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[0])) = v58
	F_CacheRegisterRelcacheCallback(m, int32(1813))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	F_CacheRegisterSyscacheCallback(m, int32(82), int32(1814), int64(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	F_CacheRegisterSyscacheCallback(m, int32(14), int32(1815), int64(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	F_CacheRegisterSyscacheCallback(m, int32(19), int32(1816), int64(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[0]))
	v85 = v84
	goto L3
L22:
	;
	v93 = F_repalloc(m, v85, v89<<(uint(int32(3))%32))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L7
	} else {
		goto L25
	}
L23:
	;
	v103 = v85
	v104 = v87
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[5])) = v104 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v103+v104<<(uint(int32(2))%32)))) = v112
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[1]))
	v117 = v9 + int32(124)
	v118 = int32(0)
	v120 = F_hash_search(m, v115, v117, v118, v118)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L33
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[4])) = v89 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[0])) = v93
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[5]))
	v103 = v93
	v104 = v102
	goto L24
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L7
	} else {
		goto L448
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L7
	} else {
		goto L444
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L7
	} else {
		goto L440
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L7
	} else {
		goto L436
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L7
	} else {
		goto L432
	}
L31:
	;
	if l1&int32(623) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L32:
	;
	F_ReleaseCatCache(m, v227)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L47
	}
L33:
	;
	if v120 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v125 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+124)))
	v126 = F_SearchSysCache1(m, int32(82), v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+312)))
	if v186&int32(1) != 0 {
		v231 = v120
		goto L31
	} else {
		goto L43
	}
L37:
	;
	if v126 == int32(0) {
		goto L30
	} else {
		goto L38
	}
L38:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+22)))
	v132 = v130 + v131
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+82)))
	if v133 == int32(0) {
		goto L29
	} else {
		goto L39
	}
L39:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[1]))
	v141 = F_hash_search(m, v137, v117, int32(1), v9+int32(123))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	base.MemoryFill(m, v141+int32(4), int32(0), int32(324))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v148
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[1]))
	v152 = F_get_hash_value(m, v151, v117)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = v152
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v141)+8)) = uint16(v155)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+10)) = uint8(v157)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+11)) = uint8(v159)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+129)))
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+12)) = uint8(v161)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+13)) = uint8(v163)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v132)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+16)) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v132)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+20)) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v132)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+24)) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v132)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+28)) = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v132)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+32)) = v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v141)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+312)) = v175 | int32(1)
	if v163 != int32(100) {
		v225 = v141
		v227 = v126
		goto L32
	} else {
		goto L42
	}
L42:
	;
	v181 = int32(_a_F_lookup_type_cache_6)
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+320)) = v182
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[6])) = v141
	v225 = v141
	v227 = v126
	goto L32
L43:
	;
	v190 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+124)))
	v191 = F_SearchSysCache1(m, int32(82), v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	if v191 == int32(0) {
		goto L28
	} else {
		goto L45
	}
L45:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191)+16))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+22)))
	v197 = v195 + v196
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+82)))
	if v198 == int32(0) {
		goto L27
	} else {
		goto L46
	}
L46:
	;
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+8)) = uint16(v201)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+10)) = uint8(v203)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+11)) = uint8(v205)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+129)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v207)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+13)) = uint8(v209)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v197)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v197)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+20)) = v213
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v197)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+24)) = v215
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v197)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+28)) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v197)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+32)) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v120)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+312)) = v221 | int32(1)
	v225 = v120
	v227 = v191
	goto L32
L47:
	;
	v231 = v225
	goto L31
L48:
	;
	if l1&int32(33) == int32(0) {
		v274 = l1
		goto L58
	} else {
		goto L59
	}
L49:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+312)))
	if v239&int32(2) != 0 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	v244 = F_GetDefaultOpClass(m, v242, int32(403))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L7
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+40)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v256&int32(-123) | int32(2)
	goto L48
L52:
	;
	if v244 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v246 = F_get_opclass_family(m, v244)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L7
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v251 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v251
	v254 = v251
	goto L51
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+36)) = v246
	v249 = F_get_opclass_input_type(m, v244)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	v254 = v249
	goto L51
L58:
	;
	if v274&int32(_a_F_lookup_type_cache_7) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+312)))
	if v267&int32(8) != 0 {
		v274 = l1
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	if v272 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v273 = l1
	goto L63
L62:
	;
	v273 = l1 | int32(1024)
	goto L63
L63:
	;
	v274 = v273
	goto L58
L64:
	;
	if v274&int32(33) == int32(0) {
		goto L74
	} else {
		goto L75
	}
L65:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+312)))
	if v279&int32(4) != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	v284 = F_GetDefaultOpClass(m, v282, int32(405))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L7
	} else {
		goto L68
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+48)) = v294
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v296&int32(-389) | int32(4)
	goto L64
L68:
	;
	if v284 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v286 = F_get_opclass_family(m, v284)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L7
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v291 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+44)) = v291
	v294 = v291
	goto L67
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+44)) = v286
	v289 = F_get_opclass_input_type(m, v284)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	v294 = v289
	goto L67
L74:
	;
	if v274&int32(2) == int32(0) {
		goto L120
	} else {
		goto L121
	}
L75:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+312)))
	if v307&int32(8) != 0 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	if v310 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	if v390 != v392 {
		goto L117
	} else {
		goto L118
	}
L78:
	;
	if v324 != int32(2988) {
		goto L88
	} else {
		goto L89
	}
L79:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v313 = F_get_opfamily_member(m, v310, v311, v311, int32(3))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L7
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	if v316 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	if v313 != 0 {
		v324 = v313
		goto L78
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v390 = int32(0)
	goto L77
L85:
	;
	goto L86
L86:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v231)+48))
	v322 = F_get_opfamily_member(m, v316, v320, v320, int32(1))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	v324 = v322
	goto L78
L88:
	;
	if v324 != int32(1070) {
		v390 = v324
		goto L77
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v377&int32(_a_F_lookup_type_cache_16) != 0 {
		goto L113
	} else {
		goto L114
	}
L91:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v329&int32(512) == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v335 = F_get_base_element_type(m, v334)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L7
	} else {
		goto L96
	}
L93:
	;
	v369 = v329
	goto L94
L94:
	;
	v390 = v369 << (uint(int32(21)) % 32) >> (uint(int32(31)) % 32) & int32(1070)
	goto L77
L95:
	;
	v367 = v365 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v367
	v369 = v367
	goto L94
L96:
	;
	if v335 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v365 = v339
	goto L95
L98:
	;
	goto L99
L99:
	;
	v341 = F_lookup_type_cache(m, v335, int32(_a_F_lookup_type_cache_17))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v341)+52))
	if v344 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v346 = v343 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v346
	v348 = v346
	goto L103
L102:
	;
	v348 = v343
	goto L103
L103:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v341)+64))
	if v349 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v351 = v348 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v351
	v353 = v351
	goto L106
L105:
	;
	v353 = v348
	goto L106
L106:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v341)+68))
	if v354 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v356 = v353 | int32(_a_F_lookup_type_cache_11)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v356
	v358 = v356
	goto L109
L108:
	;
	v358 = v353
	goto L109
L109:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v341)+72))
	if v361 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v362 = v358 | int32(_a_F_lookup_type_cache_12)
	goto L112
L111:
	;
	v362 = v358
	goto L112
L112:
	;
	v365 = v362
	goto L95
L113:
	;
	v383 = v377
	goto L115
L114:
	;
	F_cache_record_field_properties(m, v231)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L7
	} else {
		goto L116
	}
L115:
	;
	v390 = v383 << (uint(int32(16)) % 32) >> (uint(int32(31)) % 32) & int32(2988)
	goto L77
L116:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v383 = v382
	goto L115
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+80)) = int32(0)
	goto L119
L118:
	;
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+52)) = v390
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v397&int32(-393) | int32(8)
	goto L74
L120:
	;
	if v274&int32(4) == int32(0) {
		goto L157
	} else {
		goto L158
	}
L121:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v409&int32(16) != 0 {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	if v412 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+56)) = v486
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v485 | int32(16)
	goto L120
L124:
	;
	v485 = v409
	v486 = int32(0)
	goto L123
L125:
	;
	goto L126
L126:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v418 = F_get_opfamily_member(m, v412, v416, v416, int32(1))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v418 != int32(2990) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if v418 != int32(1072) {
		v485 = v420
		v486 = v418
		goto L123
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	if v420&int32(_a_F_lookup_type_cache_16) == int32(0) {
		goto L153
	} else {
		goto L154
	}
L131:
	;
	if v420&int32(512) == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v430 = F_get_base_element_type(m, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L7
	} else {
		goto L136
	}
L133:
	;
	v463 = v420
	goto L134
L134:
	;
	v485 = v463
	v486 = v463 << (uint(int32(20)) % 32) >> (uint(int32(31)) % 32) & int32(1072)
	goto L123
L135:
	;
	v463 = v460 | int32(512)
	goto L134
L136:
	;
	if v430 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v460 = v434
	goto L135
L138:
	;
	goto L139
L139:
	;
	v436 = F_lookup_type_cache(m, v430, int32(_a_F_lookup_type_cache_17))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v436)+52))
	if v439 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v441 = v438 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v441
	v443 = v441
	goto L143
L142:
	;
	v443 = v438
	goto L143
L143:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v436)+64))
	if v444 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v446 = v443 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v446
	v448 = v446
	goto L146
L145:
	;
	v448 = v443
	goto L146
L146:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v436)+68))
	if v449 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v451 = v448 | int32(_a_F_lookup_type_cache_11)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v451
	v453 = v451
	goto L149
L148:
	;
	v453 = v448
	goto L149
L149:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v436)+72))
	if v456 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v457 = v453 | int32(_a_F_lookup_type_cache_12)
	goto L152
L151:
	;
	v457 = v453
	goto L152
L152:
	;
	v460 = v457
	goto L135
L153:
	;
	F_cache_record_field_properties(m, v231)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L7
	} else {
		goto L156
	}
L154:
	;
	v478 = v420
	goto L155
L155:
	;
	v485 = v478
	v486 = v478 << (uint(int32(15)) % 32) >> (uint(int32(31)) % 32) & int32(2990)
	goto L123
L156:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v478 = v477
	goto L155
L157:
	;
	if v274&int32(72) == int32(0) {
		goto L194
	} else {
		goto L195
	}
L158:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v497&int32(32) != 0 {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	if v500 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+60)) = v574
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v573 | int32(32)
	goto L157
L161:
	;
	v573 = v497
	v574 = int32(0)
	goto L160
L162:
	;
	goto L163
L163:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v506 = F_get_opfamily_member(m, v500, v504, v504, int32(5))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L7
	} else {
		goto L164
	}
L164:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v506 != int32(2991) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	if v506 != int32(1073) {
		v573 = v508
		v574 = v506
		goto L160
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	if v508&int32(_a_F_lookup_type_cache_16) == int32(0) {
		goto L190
	} else {
		goto L191
	}
L168:
	;
	if v508&int32(512) == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v518 = F_get_base_element_type(m, v517)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L7
	} else {
		goto L173
	}
L170:
	;
	v551 = v508
	goto L171
L171:
	;
	v573 = v551
	v574 = v551 << (uint(int32(20)) % 32) >> (uint(int32(31)) % 32) & int32(1073)
	goto L160
L172:
	;
	v551 = v548 | int32(512)
	goto L171
L173:
	;
	if v518 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v548 = v522
	goto L172
L175:
	;
	goto L176
L176:
	;
	v524 = F_lookup_type_cache(m, v518, int32(_a_F_lookup_type_cache_17))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L7
	} else {
		goto L177
	}
L177:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v524)+52))
	if v527 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v529 = v526 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v529
	v531 = v529
	goto L180
L179:
	;
	v531 = v526
	goto L180
L180:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v524)+64))
	if v532 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v534 = v531 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v534
	v536 = v534
	goto L183
L182:
	;
	v536 = v531
	goto L183
L183:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v524)+68))
	if v537 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v539 = v536 | int32(_a_F_lookup_type_cache_11)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v539
	v541 = v539
	goto L186
L185:
	;
	v541 = v536
	goto L186
L186:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v524)+72))
	if v544 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v545 = v541 | int32(_a_F_lookup_type_cache_12)
	goto L189
L188:
	;
	v545 = v541
	goto L189
L189:
	;
	v548 = v545
	goto L172
L190:
	;
	F_cache_record_field_properties(m, v231)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L7
	} else {
		goto L193
	}
L191:
	;
	v566 = v508
	goto L192
L192:
	;
	v573 = v566
	v574 = v566 << (uint(int32(15)) % 32) >> (uint(int32(31)) % 32) & int32(2991)
	goto L160
L193:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v566 = v565
	goto L192
L194:
	;
	if v274&int32(144) == int32(0) {
		goto L234
	} else {
		goto L235
	}
L195:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+312)))
	if v585&int32(64) != 0 {
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	if v588 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v231)+64))
	if v661 != v663 {
		goto L231
	} else {
		goto L232
	}
L198:
	;
	v661 = int32(0)
	goto L197
L199:
	;
	goto L200
L200:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v231)+40))
	v594 = F_get_opfamily_proc(m, v588, v592, v592, int32(1))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L7
	} else {
		goto L201
	}
L201:
	;
	if v594 != int32(2987) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	if v594 != int32(382) {
		v661 = v594
		goto L197
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v648&int32(_a_F_lookup_type_cache_16) != 0 {
		goto L227
	} else {
		goto L228
	}
L205:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v600&int32(512) == int32(0) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v606 = F_get_base_element_type(m, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L7
	} else {
		goto L210
	}
L207:
	;
	v640 = v600
	goto L208
L208:
	;
	v661 = v640 << (uint(int32(20)) % 32) >> (uint(int32(31)) % 32) & int32(382)
	goto L197
L209:
	;
	v638 = v636 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v638
	v640 = v638
	goto L208
L210:
	;
	if v606 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v636 = v610
	goto L209
L212:
	;
	goto L213
L213:
	;
	v612 = F_lookup_type_cache(m, v606, int32(_a_F_lookup_type_cache_17))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L7
	} else {
		goto L214
	}
L214:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v612)+52))
	if v615 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v617 = v614 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v617
	v619 = v617
	goto L217
L216:
	;
	v619 = v614
	goto L217
L217:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v612)+64))
	if v620 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v622 = v619 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v622
	v624 = v622
	goto L220
L219:
	;
	v624 = v619
	goto L220
L220:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v612)+68))
	if v625 != 0 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v627 = v624 | int32(_a_F_lookup_type_cache_11)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v627
	v629 = v627
	goto L223
L222:
	;
	v629 = v624
	goto L223
L223:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v612)+72))
	if v632 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v633 = v629 | int32(_a_F_lookup_type_cache_12)
	goto L226
L225:
	;
	v633 = v629
	goto L226
L226:
	;
	v636 = v633
	goto L209
L227:
	;
	v654 = v648
	goto L229
L228:
	;
	F_cache_record_field_properties(m, v231)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L7
	} else {
		goto L230
	}
L229:
	;
	v661 = v654 << (uint(int32(15)) % 32) >> (uint(int32(31)) % 32) & int32(2987)
	goto L197
L230:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v654 = v653
	goto L229
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+108)) = int32(0)
	goto L233
L232:
	;
	goto L233
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+64)) = v661
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v668 | int32(64)
	goto L194
L234:
	;
	if v274&int32(_a_F_lookup_type_cache_8) == int32(0) {
		goto L303
	} else {
		goto L304
	}
L235:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+312)))
	if v678&int32(128) != 0 {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v681 = int32(0)
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	if v682 == v681 {
		v823 = v681
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v231)+68))
	if v823 != v826 {
		goto L300
	} else {
		goto L301
	}
L238:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	if v685 != 0 {
		goto L244
	} else {
		goto L245
	}
L239:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v810&int32(512) != 0 {
		goto L296
	} else {
		goto L297
	}
L240:
	;
	v823 = v802 << (uint(int32(19)) % 32) >> (uint(int32(31)) % 32) & int32(3902)
	goto L237
L241:
	;
	v800 = v798 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v800
	v802 = v800
	goto L240
L242:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	v784 = F_lookup_type_cache(m, v782, int32(_a_F_lookup_type_cache_18))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L7
	} else {
		goto L289
	}
L243:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v733&int32(512) == int32(0) {
		goto L268
	} else {
		goto L269
	}
L244:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v231)+48))
	v688 = F_get_opfamily_member(m, v682, v686, v686, int32(1))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L7
	} else {
		goto L247
	}
L245:
	;
	v693 = v682
	goto L246
L246:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v231)+48))
	v696 = F_get_opfamily_proc(m, v693, v694, v694, int32(1))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L7
	} else {
		goto L249
	}
L247:
	;
	if v688 != v685 {
		v823 = v681
		goto L237
	} else {
		goto L248
	}
L248:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	v693 = v691
	goto L246
L249:
	;
	if v696 <= int32(_a_F_lookup_type_cache_19) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	if v696 == int32(626) {
		goto L243
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	if v696 == int32(_a_F_lookup_type_cache_20) {
		goto L239
	} else {
		goto L262
	}
L253:
	;
	if v696 != int32(3902) {
		v823 = v696
		goto L237
	} else {
		goto L254
	}
L254:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v704&int32(512) != 0 {
		v802 = v704
		goto L240
	} else {
		goto L255
	}
L255:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v231)+200))
	if v707 != 0 {
		v781 = v707
		goto L242
	} else {
		goto L256
	}
L256:
	;
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+13)))
	if v708 == int32(114) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	F_load_rangetype_info(m, v231)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L7
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v798 = v715
	goto L241
L260:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v231)+200))
	if v713 != 0 {
		v781 = v713
		goto L242
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	if v696 != int32(_a_F_lookup_type_cache_21) {
		v823 = v696
		goto L237
	} else {
		goto L263
	}
L263:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v720&int32(_a_F_lookup_type_cache_16) != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v726 = v720
	goto L266
L265:
	;
	F_cache_record_field_properties(m, v231)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L7
	} else {
		goto L267
	}
L266:
	;
	v823 = v726 << (uint(int32(14)) % 32) >> (uint(int32(31)) % 32) & int32(_a_F_lookup_type_cache_21)
	goto L237
L267:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v726 = v725
	goto L266
L268:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v739 = F_get_base_element_type(m, v738)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L7
	} else {
		goto L272
	}
L269:
	;
	v773 = v733
	goto L270
L270:
	;
	v823 = v773 << (uint(int32(19)) % 32) >> (uint(int32(31)) % 32) & int32(626)
	goto L237
L271:
	;
	v771 = v769 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v771
	v773 = v771
	goto L270
L272:
	;
	if v739 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v769 = v743
	goto L271
L274:
	;
	goto L275
L275:
	;
	v745 = F_lookup_type_cache(m, v739, int32(_a_F_lookup_type_cache_17))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L7
	} else {
		goto L276
	}
L276:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v745)+52))
	if v748 != 0 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v750 = v747 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v750
	v752 = v750
	goto L279
L278:
	;
	v752 = v747
	goto L279
L279:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v745)+64))
	if v753 != 0 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v755 = v752 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v755
	v757 = v755
	goto L282
L281:
	;
	v757 = v752
	goto L282
L282:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v745)+68))
	if v758 != 0 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v760 = v757 | int32(_a_F_lookup_type_cache_11)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v760
	v762 = v760
	goto L285
L284:
	;
	v762 = v757
	goto L285
L285:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v745)+72))
	if v765 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v766 = v762 | int32(_a_F_lookup_type_cache_12)
	goto L288
L287:
	;
	v766 = v762
	goto L288
L288:
	;
	v769 = v766
	goto L271
L289:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v784)+68))
	if v787 != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v789 = v786 | int32(_a_F_lookup_type_cache_11)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v789
	v791 = v789
	goto L292
L291:
	;
	v791 = v786
	goto L292
L292:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v784)+72))
	if v794 != 0 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v795 = v791 | int32(_a_F_lookup_type_cache_12)
	goto L295
L294:
	;
	v795 = v791
	goto L295
L295:
	;
	v798 = v795
	goto L241
L296:
	;
	v816 = v810
	goto L298
L297:
	;
	F_cache_multirange_element_properties(m, v231)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L7
	} else {
		goto L299
	}
L298:
	;
	v823 = v816 << (uint(int32(19)) % 32) >> (uint(int32(31)) % 32) & int32(_a_F_lookup_type_cache_20)
	goto L237
L299:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v816 = v815
	goto L298
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+136)) = int32(0)
	goto L302
L301:
	;
	goto L302
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+68)) = v823
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v831 | int32(128)
	goto L234
L303:
	;
	if v274&int32(32) == int32(0) {
		goto L372
	} else {
		goto L373
	}
L304:
	;
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+313)))
	if v842&int32(1) != 0 {
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v845 = int32(0)
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	if v846 == v845 {
		v987 = v845
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v231)+72))
	if v987 != v990 {
		goto L369
	} else {
		goto L370
	}
L307:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	if v849 != 0 {
		goto L313
	} else {
		goto L314
	}
L308:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v974&int32(512) != 0 {
		goto L365
	} else {
		goto L366
	}
L309:
	;
	v987 = v966 << (uint(int32(18)) % 32) >> (uint(int32(31)) % 32) & int32(3417)
	goto L306
L310:
	;
	v964 = v962 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v964
	v966 = v964
	goto L309
L311:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v945)))
	v948 = F_lookup_type_cache(m, v946, int32(_a_F_lookup_type_cache_18))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L7
	} else {
		goto L358
	}
L312:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v897&int32(512) == int32(0) {
		goto L337
	} else {
		goto L338
	}
L313:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v231)+48))
	v852 = F_get_opfamily_member(m, v846, v850, v850, int32(1))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L7
	} else {
		goto L316
	}
L314:
	;
	v857 = v846
	goto L315
L315:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v231)+48))
	v860 = F_get_opfamily_proc(m, v857, v858, v858, int32(2))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L7
	} else {
		goto L318
	}
L316:
	;
	if v852 != v849 {
		v987 = v845
		goto L306
	} else {
		goto L317
	}
L317:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v231)+44))
	v857 = v855
	goto L315
L318:
	;
	if v860 <= int32(_a_F_lookup_type_cache_20) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	if v860 == int32(782) {
		goto L312
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	if v860 == int32(_a_F_lookup_type_cache_22) {
		goto L308
	} else {
		goto L331
	}
L322:
	;
	if v860 != int32(3417) {
		v987 = v860
		goto L306
	} else {
		goto L323
	}
L323:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v868&int32(512) != 0 {
		v966 = v868
		goto L309
	} else {
		goto L324
	}
L324:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v231)+200))
	if v871 != 0 {
		v945 = v871
		goto L311
	} else {
		goto L325
	}
L325:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+13)))
	if v872 == int32(114) {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	F_load_rangetype_info(m, v231)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L7
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v962 = v879
	goto L310
L329:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v231)+200))
	if v877 != 0 {
		v945 = v877
		goto L311
	} else {
		goto L330
	}
L330:
	;
	goto L328
L331:
	;
	if v860 != int32(_a_F_lookup_type_cache_23) {
		v987 = v860
		goto L306
	} else {
		goto L332
	}
L332:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v884&int32(_a_F_lookup_type_cache_16) != 0 {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v890 = v884
	goto L335
L334:
	;
	F_cache_record_field_properties(m, v231)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L7
	} else {
		goto L336
	}
L335:
	;
	v987 = v890 << (uint(int32(13)) % 32) >> (uint(int32(31)) % 32) & int32(_a_F_lookup_type_cache_23)
	goto L306
L336:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v890 = v889
	goto L335
L337:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v903 = F_get_base_element_type(m, v902)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L7
	} else {
		goto L341
	}
L338:
	;
	v937 = v897
	goto L339
L339:
	;
	v987 = v937 << (uint(int32(18)) % 32) >> (uint(int32(31)) % 32) & int32(782)
	goto L306
L340:
	;
	v935 = v933 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v935
	v937 = v935
	goto L339
L341:
	;
	if v903 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v933 = v907
	goto L340
L343:
	;
	goto L344
L344:
	;
	v909 = F_lookup_type_cache(m, v903, int32(_a_F_lookup_type_cache_17))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L7
	} else {
		goto L345
	}
L345:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v909)+52))
	if v912 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v914 = v911 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v914
	v916 = v914
	goto L348
L347:
	;
	v916 = v911
	goto L348
L348:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v909)+64))
	if v917 != 0 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v919 = v916 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v919
	v921 = v919
	goto L351
L350:
	;
	v921 = v916
	goto L351
L351:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v909)+68))
	if v922 != 0 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v924 = v921 | int32(_a_F_lookup_type_cache_11)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v924
	v926 = v924
	goto L354
L353:
	;
	v926 = v921
	goto L354
L354:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v909)+72))
	if v929 != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v930 = v926 | int32(_a_F_lookup_type_cache_12)
	goto L357
L356:
	;
	v930 = v926
	goto L357
L357:
	;
	v933 = v930
	goto L340
L358:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v948)+68))
	if v951 != 0 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v953 = v950 | int32(_a_F_lookup_type_cache_11)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v953
	v955 = v953
	goto L361
L360:
	;
	v955 = v950
	goto L361
L361:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v948)+72))
	if v958 != 0 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v959 = v955 | int32(_a_F_lookup_type_cache_12)
	goto L364
L363:
	;
	v959 = v955
	goto L364
L364:
	;
	v962 = v959
	goto L310
L365:
	;
	v980 = v974
	goto L367
L366:
	;
	F_cache_multirange_element_properties(m, v231)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L7
	} else {
		goto L368
	}
L367:
	;
	v987 = v980 << (uint(int32(18)) % 32) >> (uint(int32(31)) % 32) & int32(_a_F_lookup_type_cache_22)
	goto L306
L368:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	v980 = v979
	goto L367
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+164)) = int32(0)
	goto L371
L370:
	;
	goto L371
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+72)) = v987
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+312)) = v995 | int32(256)
	goto L303
L372:
	;
	if v274&int32(64) == int32(0) {
		goto L379
	} else {
		goto L380
	}
L373:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v231)+80))
	if v1006 != 0 {
		goto L372
	} else {
		goto L374
	}
L374:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v231)+52))
	if v1007 == int32(0) {
		goto L372
	} else {
		goto L375
	}
L375:
	;
	v1010 = F_get_opcode(m, v1007)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L7
	} else {
		goto L376
	}
L376:
	;
	if v1010 == int32(0) {
		goto L372
	} else {
		goto L377
	}
L377:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[3]))
	F_fmgr_info_cxt(m, v1010, v231+int32(76), v1017)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L7
	} else {
		goto L378
	}
L378:
	;
	goto L372
L379:
	;
	if v274&int32(128) == int32(0) {
		goto L384
	} else {
		goto L385
	}
L380:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v231)+108))
	if v1025 != 0 {
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v231)+64))
	if v1026 == int32(0) {
		goto L379
	} else {
		goto L382
	}
L382:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[3]))
	F_fmgr_info_cxt(m, v1026, v231+int32(104), v1032)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L7
	} else {
		goto L383
	}
L383:
	;
	goto L379
L384:
	;
	if v274&int32(_a_F_lookup_type_cache_9) == int32(0) {
		goto L389
	} else {
		goto L390
	}
L385:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v231)+136))
	if v1040 != 0 {
		goto L384
	} else {
		goto L386
	}
L386:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v231)+68))
	if v1041 == int32(0) {
		goto L384
	} else {
		goto L387
	}
L387:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[3]))
	F_fmgr_info_cxt(m, v1041, v231+int32(132), v1047)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L7
	} else {
		goto L388
	}
L388:
	;
	goto L384
L389:
	;
	if v274&int32(256) == int32(0) {
		goto L394
	} else {
		goto L395
	}
L390:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v231)+164))
	if v1055 != 0 {
		goto L389
	} else {
		goto L391
	}
L391:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v231)+72))
	if v1056 == int32(0) {
		goto L389
	} else {
		goto L392
	}
L392:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[3]))
	F_fmgr_info_cxt(m, v1056, v231+int32(160), v1062)
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L7
	} else {
		goto L393
	}
L393:
	;
	goto L389
L394:
	;
	if v274&int32(2048) == int32(0) {
		goto L399
	} else {
		goto L400
	}
L395:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v231)+188))
	if v1070 != 0 {
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+13)))
	if v1071 != int32(99) {
		goto L394
	} else {
		goto L397
	}
L397:
	;
	F_load_typcache_tupdesc(m, v231)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L7
	} else {
		goto L398
	}
L398:
	;
	goto L394
L399:
	;
	if v274&int32(_a_F_lookup_type_cache_10) == int32(0) {
		goto L408
	} else {
		goto L409
	}
L400:
	;
	v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+13)))
	if v1080 != int32(114) {
		goto L399
	} else {
		goto L401
	}
L401:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v231)+200))
	if v1083 == int32(0) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	F_load_rangetype_info(m, v231)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L7
	} else {
		goto L405
	}
L403:
	;
	goto L404
L404:
	;
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083)+312)))
	if v1088&int32(1) != 0 {
		goto L399
	} else {
		goto L406
	}
L405:
	;
	goto L399
L406:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1083)))
	v1093 = F_lookup_type_cache(m, v1091, int32(0))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L7
	} else {
		goto L407
	}
L407:
	;
	goto L399
L408:
	;
	if v274&int32(_a_F_lookup_type_cache_11) == int32(0) {
		goto L415
	} else {
		goto L416
	}
L409:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v231)+296))
	if v1100 != 0 {
		goto L408
	} else {
		goto L410
	}
L410:
	;
	v1101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+13)))
	if v1101 != int32(109) {
		goto L408
	} else {
		goto L411
	}
L411:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v1105 = F_get_multirange_range(m, v1104)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L7
	} else {
		goto L412
	}
L412:
	;
	if v1105 == int32(0) {
		goto L26
	} else {
		goto L413
	}
L413:
	;
	v1110 = F_lookup_type_cache(m, v1105, int32(2048))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L7
	} else {
		goto L414
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+296)) = v1110
	goto L408
L415:
	;
	if v274&int32(_a_F_lookup_type_cache_12) == int32(0) {
		goto L420
	} else {
		goto L421
	}
L416:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v231)+300))
	if v1118 != 0 {
		goto L415
	} else {
		goto L417
	}
L417:
	;
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+13)))
	if v1119 != int32(100) {
		goto L415
	} else {
		goto L418
	}
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+304)) = int32(-1)
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	v1127 = F_getBaseTypeAndTypmod(m, v1124, v231+int32(304))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L7
	} else {
		goto L419
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+300)) = v1127
	goto L415
L420:
	;
	v1142 = int32(_a_F_lookup_type_cache_13)
	v1144 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[5])) = v1144 - int32(1)
	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+13)))
	if v1148 != int32(99) {
		goto L425
	} else {
		goto L426
	}
L421:
	;
	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+314)))
	if v1134&int32(8) != 0 {
		goto L420
	} else {
		goto L422
	}
L422:
	;
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+13)))
	if v1137 != int32(100) {
		goto L420
	} else {
		goto L423
	}
L423:
	;
	F_load_domaintype_info(m, v231)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L7
	} else {
		goto L424
	}
L424:
	;
	goto L420
L425:
	;
	m.G0 = v9 + int32(128)
	return v231
L426:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v231)+312))
	if v1151&int32(-1572865) == int32(0) {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v231)+188))
	if v1156 == int32(0) {
		goto L425
	} else {
		goto L430
	}
L428:
	;
	goto L429
L429:
	;
	v1160 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[2]))
	v1166 = F_hash_search(m, v1160, v231+int32(16), int32(1), v9+int32(72))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L7
	} else {
		goto L431
	}
L430:
	;
	goto L429
L431:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v231)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1166))) = v1168
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	*(*int32)(unsafe.Add(mBase, uint32(v1166)+4)) = v1170
	goto L425
L432:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L7
	} else {
		goto L433
	}
L433:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v1184
	F_errmsg(m, int32(_a_F_lookup_type_cache_2), v9)
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L7
	} else {
		goto L434
	}
L434:
	;
	F_errfinish(m, int32(_a_F_lookup_type_cache_3), int32(485), int32(_a_F_lookup_type_cache_4))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L7
	} else {
		goto L435
	}
L435:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L436:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L7
	} else {
		goto L437
	}
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v132 + int32(4)
	F_errmsg(m, int32(_a_F_lookup_type_cache_5), v9+int32(32))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L7
	} else {
		goto L438
	}
L438:
	;
	F_errfinish(m, int32(_a_F_lookup_type_cache_3), int32(491), int32(_a_F_lookup_type_cache_4))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L7
	} else {
		goto L439
	}
L439:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L440:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L7
	} else {
		goto L441
	}
L441:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v1221
	F_errmsg(m, int32(_a_F_lookup_type_cache_2), v9+int32(48))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L7
	} else {
		goto L442
	}
L442:
	;
	F_errfinish(m, int32(_a_F_lookup_type_cache_3), int32(540), int32(_a_F_lookup_type_cache_4))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L7
	} else {
		goto L443
	}
L443:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L444:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L7
	} else {
		goto L445
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v197 + int32(4)
	F_errmsg(m, int32(_a_F_lookup_type_cache_5), v9-int32(-64))
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L7
	} else {
		goto L446
	}
L446:
	;
	F_errfinish(m, int32(_a_F_lookup_type_cache_3), int32(546), int32(_a_F_lookup_type_cache_4))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L7
	} else {
		goto L447
	}
L447:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L448:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v1257
	F_errmsg_internal(m, int32(_a_F_lookup_type_cache_14), v9+int32(16))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L7
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(_a_F_lookup_type_cache_3), int32(1073), int32(_a_F_lookup_type_cache_15))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L7
	} else {
		goto L450
	}
L450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_makeTypeNameFromNameList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc0(m, int32(32))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = int32(-1)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = int64(-4294967296)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(68)
		return v4
	}
}
func F_typeInheritsFrom(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = F_typeOrDomainTypeRelid(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L3
	} else {
		goto L59
	}
L2:
	;
	m.G0 = v12 + int32(80)
	return v194
L3:
	;
	return int32(0)
L4:
	;
	if v14 == int32(0) {
		v194 = v3
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v20 = F_typeidTypeRelid(m, l1)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if v20 == int32(0) {
		v194 = v3
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v26 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(v20))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v26 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v31)+126)))
	F_ReleaseCatCache(m, v26)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v33 != int32(1) {
		v194 = v3
		goto L2
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v14
	v43 = F_list_make1_impl(m, int32(480), v12+int32(12))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v47 = F_table_open(m, int32(2611), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	if v43 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	F_relation_close(m, v47, int32(1))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L3
	} else {
		goto L56
	}
L15:
	;
	v172 = v162
	v179 = v169
	v181 = int32(0)
	goto L14
L16:
	;
	v162 = int32(0)
	v169 = v3
	goto L15
L17:
	;
	goto L18
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v52 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v162 = v43
	v169 = v3
	goto L15
L20:
	;
	goto L21
L21:
	;
	v55 = v43
	v60 = v3
	v62 = v3
	goto L22
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v60<<(uint(int32(2))%32))))
	v69 = int32(0)
	if v62 == v69 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v162 = v149
	v169 = v156
	goto L15
L24:
	;
	if v107 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	v107 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v75 <= int32(0) {
		v101 = v69
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v107 = v101
	goto L24
L29:
	;
	v78 = int32(0)
	if v78 < v75 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v81 = v75
	goto L32
L31:
	;
	v81 = v78
	goto L32
L32:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v84 = int32(0)
	goto L33
L33:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v82+v84<<(uint(int32(2))%32))))
	v93 = base.B2i32(v92 == v68)
	if v92 == v68 {
		v101 = v93
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v101 = v93
	goto L28
L35:
	;
	v95 = v84 + int32(1)
	if v95 != v81 {
		v84 = v95
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v110 = F_lappend_oid(m, v62, v68)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L3
	} else {
		goto L40
	}
L38:
	;
	v149 = v55
	v156 = v62
	goto L39
L39:
	;
	v159 = v60 + int32(1)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v159 < v160 {
		v55 = v149
		v60 = v159
		v62 = v156
		goto L22
	} else {
		goto L55
	}
L40:
	;
	v113 = v12 + int32(16)
	F_ScanKeyInit(m, v113, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v68))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v121 = int32(1)
	v124 = F_systable_beginscan(m, v47, int32(2680), v121, int32(0), v121, v113)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v126 = v55
	goto L43
L43:
	;
	v135 = F_systable_getnext(m, v124)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L45
	}
L44:
	;
	F_systable_endscan(m, v124)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L54
	}
L45:
	;
	if v135 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+22)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v137+v138)+4))
	if v20 == v140 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	goto L44
L49:
	;
	F_systable_endscan(m, v124)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L3
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v145 = F_lappend_oid(m, v126, v140)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L53
	}
L52:
	;
	v172 = v126
	v179 = v110
	v181 = int32(1)
	goto L14
L53:
	;
	v126 = v145
	goto L43
L54:
	;
	v149 = v126
	v156 = v110
	goto L39
L55:
	;
	goto L23
L56:
	;
	F_list_free(m, v179)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	F_list_free(m, v172)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v194 = v181
	goto L2
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v20
	F_errmsg_internal(m, int32(_a_F_typeInheritsFrom_0), v12)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_typeInheritsFrom_1), int32(363), int32(_a_F_typeInheritsFrom_2))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
