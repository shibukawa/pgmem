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
						F_errfinish(m, int32(_a_F_AssignTypeArrayOid_1), int32(2458), int32(_a_F_AssignTypeArrayOid_2))
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
	var v27 int32
	_ = v27
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
	v20 = v18 + v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if l2 != 0 {
		v34 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l3 != 0 {
		v45 = l3
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
		v34 = int32(0)
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v30 = F_text_to_cstring(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v32 = F_stringToNode(m, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v34 = v32
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
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v40 = F_heap_getattr_5(m, l0, int32(32), v37, v16+int32(31))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+31)))
	if v42 != 0 {
		v45 = int32(0)
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v43 = F_pg_detoast_datum_copy(m, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v45 = v43
	goto L8
L13:
	;
	v48 = F_deleteDependencyRecordsFor(m, int32(1247), v21, int32(1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
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
	v59 = F_new_object_addresses(m)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L18
	}
L16:
	;
	F_deleteSharedDependencyRecordsFor(m, int32(1247), v21, int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
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
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+79)))
	if v61 != int32(109) {
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
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v76
	F_add_exact_object_address(m, v16+int32(4), v59)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(2615)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v66
	F_add_exact_object_address(m, v16+int32(4), v59)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L19
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	F_recordDependencyOnOwner(m, int32(1247), v21, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	F_recordDependencyOnNewAcl(m, int32(1247), v21, v89, v45)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
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
	v96 = m.ExcPending
	if v96 != 0 {
		goto L3
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v20)+100))
	if v97 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v59)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	if v107 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v59)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	if v117 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v59)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v20)+112))
	if v127 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v59)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v20)+116))
	if v137 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v59)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v20)+120))
	if v147 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v59)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L3
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v20)+124))
	if v157 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v59)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	if v167 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v59)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v20)+132))
	if v177 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1247)
	F_add_exact_object_address(m, v16+int32(4), v59)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v20)+144))
	v188 = int32(0)
	if base.B2i32(v187 == v188)|base.B2i32(v187 == int32(100)) == v188 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(3456)
	F_add_exact_object_address(m, v16+int32(4), v59)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v205 = v16 + int32(16)
	F_record_object_address_dependencies(m, v205, v59, int32(110))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	F_free_object_addresses(m, v59)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	if v34 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_recordDependencyOnExpr(m, v205, v34, int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L3
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	if v214 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L76
L78:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v20)+92))
	if v238 != 0 {
		goto L85
	} else {
		goto L86
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v214
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
	v230 = m.ExcPending
	if v230 != 0 {
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
	v237 = m.ExcPending
	if v237 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v238
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
	v250 = int32(105)
	goto L90
L89:
	;
	v250 = int32(110)
	goto L90
L90:
	;
	F_recordDependencyOn(m, v16+int32(16), v16+int32(4), v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
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
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v276 int32
	_ = v276
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v477 int32
	_ = v477
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
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
	v546 = m.ExcPending
	if v546 != 0 {
		goto L16
	} else {
		goto L135
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L16
	} else {
		goto L129
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L16
	} else {
		goto L123
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L16
	} else {
		goto L120
	}
L5:
	;
	m.G0 = v14 + int32(160)
	return v477
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v460
	v477 = v465
	goto L5
L7:
	;
	v323 = F_SearchSysCache1(m, int32(82), v303)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L16
	} else {
		goto L88
	}
L8:
	;
	if l2 != 0 {
		v460 = int32(-1)
		v465 = v6
		goto L6
	} else {
		goto L87
	}
L9:
	;
	if v303 != 0 {
		goto L7
	} else {
		goto L86
	}
L10:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v303 = v19
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
		v303 = v151
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
	v303 = v151
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
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v294 == int32(0) {
		v303 = v287
		goto L9
	} else {
		goto L84
	}
L58:
	;
	v188 = v14 + int32(140)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = int32(489)
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
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v14)+132))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
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
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v14)+132))
	v207 = int32(0)
	v209 = F_GetSysCacheOid(m, int32(81), v206, v203, v207, v207)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L16
	} else {
		goto L66
	}
L64:
	;
	v212 = int32(0)
	goto L65
L65:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(140))+8))
	*(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[0])) = v216
	goto L67
L66:
	;
	v212 = v209
	goto L65
L67:
	;
	v287 = v212
	goto L57
L68:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[1]))
	if v223 == int32(0) {
		v276 = int32(0)
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v287 = v276
	goto L57
L70:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if int32(0) < v226 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v236 = v6
	goto L74
L72:
	;
	goto L73
L73:
	;
	v276 = int32(0)
	goto L69
L74:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240+v236<<(uint(int32(2))%32))))
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
	v257 = v236 + int32(1)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v257 < v258 {
		v236 = v257
		goto L74
	} else {
		goto L83
	}
L77:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[2]))
	if v244 == v248 {
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v251 = int32(0)
	v253 = F_GetSysCacheOid(m, int32(81), v218, v244, v251, v251)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L16
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	if v253 != 0 {
		v276 = v253
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
	v297 = F_get_array_type(m, v287)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L16
	} else {
		goto L85
	}
L85:
	;
	v303 = v297
	goto L9
L86:
	;
	goto L8
L87:
	;
	v477 = v6
	goto L5
L88:
	;
	if v323 == int32(0) {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v327 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if l2 == int32(0) {
		v477 = v323
		goto L5
	} else {
		goto L119
	}
L91:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v447 = v330
	goto L90
L92:
	;
	goto L93
L93:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v323)+16))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+22)))
	v333 = v331 + v332
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+82)))
	if v334 == int32(0) {
		goto L3
	} else {
		goto L94
	}
L94:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v333)+116))
	if v337 == int32(0) {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	v343 = F_palloc(m, v340<<(uint(int32(2))%32))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L16
	} else {
		goto L96
	}
L96:
	;
	v345 = int32(0)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v346 == v345 {
		v408 = v345
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v416 = F_construct_array_builtin(m, v343, v408, int32(2275))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L16
	} else {
		goto L113
	}
L98:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if v349 <= int32(0) {
		v408 = v345
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v356 = v345
	goto L100
L100:
	;
	v364 = v356 << (uint(int32(2)) % 32)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v346)+12))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v364+v365)))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	switch v368 - int32(69) {
	case 0:
		goto L103
	default:
		goto L1
	case 3:
		goto L104
	}
L101:
	;
	v408 = v401
	goto L97
L102:
	;
	if v395 == int32(0) {
		goto L1
	} else {
		goto L111
	}
L103:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v382 == int32(0) {
		goto L1
	} else {
		goto L108
	}
L104:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	switch v371 - int32(465) {
	case 0:
		goto L106
	case 1, 3:
		goto L105
	default:
		goto L1
	}
L105:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v367)+8))
	v395 = v381
	goto L102
L106:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v367)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v374
	v379 = F_psprintf(m, int32(_a_F_LookupTypeNameExtended_7), v14+int32(32))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L16
	} else {
		goto L107
	}
L107:
	;
	v395 = v379
	goto L102
L108:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	if v385 != int32(1) {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v382)+12))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	if v390 != int32(468) {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v395 = v393
	goto L102
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v343+v364))) = v395
	v401 = v356 + int32(1)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if v401 < v402 {
		v356 = v401
		goto L100
	} else {
		goto L112
	}
L112:
	;
	goto L101
L113:
	;
	v419 = v14 + int32(140)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v419)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v419)+4)) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v419))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v419)+16)) = v419
	v426 = int32(_a_F_LookupTypeNameExtended_6)
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v419)+8)) = v427
	*(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[0])) = v14 + int32(148)
	goto L114
L114:
	;
	v434 = F_OidFunctionCall1Coll(m, v337, int32(0), v416)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L16
	} else {
		goto L115
	}
L115:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_LookupTypeNameExtended[0])) = v437
	goto L116
L116:
	;
	F_pfree(m, v343)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L16
	} else {
		goto L117
	}
L117:
	;
	F_pfree(m, v416)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L16
	} else {
		goto L118
	}
L118:
	;
	v447 = v434
	goto L90
L119:
	;
	v460 = v447
	v465 = v323
	goto L6
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v303
	F_errmsg_internal(m, int32(_a_F_LookupTypeNameExtended_8), v14)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L16
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_LookupTypeNameExtended_1), int32(209), int32(_a_F_LookupTypeNameExtended_2))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
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
	v502 = m.ExcPending
	if v502 != 0 {
		goto L16
	} else {
		goto L124
	}
L124:
	;
	v503 = F_TypeNameToString(m, l1)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L16
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v503
	F_errmsg(m, int32(_a_F_LookupTypeNameExtended_9), v14+int32(48))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L16
	} else {
		goto L126
	}
L126:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L16
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_LookupTypeNameExtended_1), int32(356), int32(_a_F_LookupTypeNameExtended_10))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
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
	v525 = m.ExcPending
	if v525 != 0 {
		goto L16
	} else {
		goto L130
	}
L130:
	;
	v526 = F_TypeNameToString(m, l1)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L16
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v526
	F_errmsg(m, int32(_a_F_LookupTypeNameExtended_11), v14+int32(16))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L16
	} else {
		goto L132
	}
L132:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L16
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_LookupTypeNameExtended_1), int32(365), int32(_a_F_LookupTypeNameExtended_10))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
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
	v549 = m.ExcPending
	if v549 != 0 {
		goto L16
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_LookupTypeNameExtended_12), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L16
	} else {
		goto L137
	}
L137:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v554)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L16
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_LookupTypeNameExtended_1), int32(410), int32(_a_F_LookupTypeNameExtended_10))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
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
func F_TypeCacheConstrCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheConstrCallback[0]))
	if v5 != 0 {
		v6 = v5
		for {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+312))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+312)) = v9 & int32(-524289)
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+320))
			if v13 != 0 {
				v6 = v13
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
func F_TypeCacheRelCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = l1
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheRelCallback[0]))
	v15 = int32(0)
	v17 = F_hash_search(m, v12, v8+int32(24), v15, v15)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	v94 = v8 + int32(4)
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheRelCallback[1]))
	F_hash_seq_init(m, v94, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L6
	} else {
		goto L29
	}
L5:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheRelCallback[2]))
	if v78 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L6:
	;
	return
L7:
	;
	if v17 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheRelCallback[1]))
	v25 = int32(0)
	v27 = F_hash_search(m, v22, v17+int32(4), v25, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v27 == int32(0) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+188))
	if v31 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+13)))
	if v57&int32(1)|base.B2i32(v61 != int32(99)) != 0 {
		goto L5
	} else {
		goto L20
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v34 = v32 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v34
	if v34 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v27)+312))
	v51 = v49 & int32(_a_F_TypeCacheRelCallback_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v51
	if v49&int32(-1572866) == int32(0) {
		goto L5
	} else {
		goto L19
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+188))
	F_FreeTupleDesc(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+192)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+188)) = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+312))
	v47 = v45 & int32(_a_F_TypeCacheRelCallback_0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+312)) = v47
	v57 = v47
	goto L11
L18:
	;
	goto L17
L19:
	;
	v57 = v51
	goto L11
L20:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheRelCallback[0]))
	v72 = F_hash_search(m, v66, v27+int32(16), int32(2), v8+int32(4))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	goto L5
L22:
	;
	v82 = v78
	goto L23
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+312))
	if v86&int32(_a_F_TypeCacheRelCallback_1) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L1
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+312)) = v86 & int32(_a_F_TypeCacheRelCallback_0)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v82)+320))
	if v92 != 0 {
		v82 = v92
		goto L23
	} else {
		goto L28
	}
L28:
	;
	goto L24
L29:
	;
	v99 = F_hash_seq_search(m, v94)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	if v99 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v104 = v99
	goto L32
L32:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+13)))
	switch v108 - int32(99) {
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
	v171 = F_hash_seq_search(m, v8+int32(4))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L49
	}
L35:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v104)+312))
	if v158&int32(_a_F_TypeCacheRelCallback_1) == int32(0) {
		goto L34
	} else {
		goto L48
	}
L36:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v104)+188))
	if v111 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if base.B2i32(v141 == int32(0))|v142&int32(1) != 0 {
		goto L34
	} else {
		goto L46
	}
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v114 = v112 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = v114
	if v114 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v104)+312))
	v134 = v132 & int32(_a_F_TypeCacheRelCallback_0)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+312)) = v134
	if v132&int32(-1572866) == int32(0) {
		goto L34
	} else {
		goto L45
	}
L41:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v104)+188))
	F_FreeTupleDesc(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v104)+192)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+188)) = int32(0)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v104)+312))
	v127 = v125 & int32(_a_F_TypeCacheRelCallback_0)
	*(*int32)(unsafe.Add(mBase, uint32(v104)+312)) = v127
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+13)))
	v141 = base.B2i32(v129 == int32(99))
	v142 = v127
	goto L37
L44:
	;
	goto L43
L45:
	;
	v141 = int32(1)
	v142 = v134
	goto L37
L46:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheRelCallback[0]))
	v156 = F_hash_search(m, v150, v104+int32(16), int32(2), v8+int32(31))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	goto L34
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+312)) = v158 & int32(_a_F_TypeCacheRelCallback_0)
	goto L34
L49:
	;
	if v171 != 0 {
		v104 = v171
		goto L32
	} else {
		goto L50
	}
L50:
	;
	goto L33
}
func F_TypeCacheTypCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheTypCallback[0]))
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v22 = F_hash_seq_search(m, v6+int32(8))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	F_hash_seq_init(m, v6+int32(8), v9)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_hash_seq_init_with_hash_value(m, v6+int32(8), v9, l2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
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
	if v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v26 = v22
	goto L12
L10:
	;
	goto L11
L11:
	;
	m.G0 = v6 + int32(32)
	return
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+312)) = v27 & int32(-524290)
	if base.B2i32(v27&int32(1) == int32(0))|v27&int32(-1572866) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v53 = F_hash_seq_search(m, v6+int32(8))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L19
	}
L15:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+13)))
	if v38 != int32(99) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+188))
	if v41 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCacheTypCallback[1]))
	v49 = F_hash_search(m, v43, v26+int32(16), int32(2), v6+int32(31))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	if v53 != 0 {
		v26 = v53
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v180 int64
	_ = v180
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	v34 = int32(0)
	v37 = m.G0
	v39 = v37 - int32(384)
	m.G0 = v39
	if base.B2i32(l7 <= v34)&base.B2i32(base.Ui32(l7) <= base.Ui32(int32(-3))) == v34 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L19
	} else {
		goto L124
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L19
	} else {
		goto L121
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L19
	} else {
		goto L117
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L19
	} else {
		goto L113
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L19
	} else {
		goto L109
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
	v345 = m.ExcPending
	if v345 != 0 {
		goto L19
	} else {
		goto L105
	}
L9:
	;
	if l8 == int32(109) {
		v171 = int32(1)
		goto L50
	} else {
		goto L51
	}
L10:
	;
	if l7 == int32(-1) {
		goto L9
	} else {
		goto L48
	}
L11:
	;
	switch l7 - int32(1) {
	case 0:
		goto L17
	case 1:
		goto L16
	default:
		goto L14
	case 3:
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if l7 == int32(-1) {
		goto L38
	} else {
		goto L39
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L19
	} else {
		goto L34
	}
L15:
	;
	if l27 == int32(105) {
		goto L10
	} else {
		goto L29
	}
L16:
	;
	if l27 == int32(115) {
		goto L10
	} else {
		goto L24
	}
L17:
	;
	if l27 == int32(99) {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = l27
	F_errmsg(m, int32(_a_F_TypeCreate_6), v39+int32(32))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(270), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = l27
	F_errmsg(m, int32(_a_F_TypeCreate_6), v39+int32(48))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(278), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = l27
	F_errmsg(m, int32(_a_F_TypeCreate_6), v39-int32(-64))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(286), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = l7
	F_errmsg(m, int32(_a_F_TypeCreate_7), v39)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(302), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	switch l27 - int32(100) {
	case 0, 5:
		goto L9
	default:
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if l7 != int32(-2) {
		goto L10
	} else {
		goto L46
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L19
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+80)) = l27
	F_errmsg(m, int32(_a_F_TypeCreate_8), v39+int32(80))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L19
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(312), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	if l27 != int32(99) {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	goto L10
L48:
	;
	if l28 != int32(112) {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	goto L9
L50:
	;
	v172 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+376)) = v172
	*(*int64)(unsafe.Add(mBase, uint32(v39)+368)) = v172
	*(*int64)(unsafe.Add(mBase, uint32(v39)+360)) = v172
	*(*int64)(unsafe.Add(mBase, uint32(v39)+352)) = v172
	v180 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+320)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v39)+328)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v39)+336)) = v180
	*(*int64)(unsafe.Add(mBase, uint32(v39)+344)) = v180
	v190 = int32(0)
	v191 = int32(128)
	base.MemoryFill(m, v39+int32(192), v190, v191)
	v194 = v39 + v191
	v196 = F_strncpy(m, v194, l2, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+63)) = uint8(v190)
	goto L53
L51:
	;
	if l21 != 0 {
		v171 = int32(1)
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v171 = base.B2i32(l4 != int32(0)) & base.B2i32(l5 != int32(99))
	goto L50
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+304)) = l32
	*(*int32)(unsafe.Add(mBase, uint32(v39)+300)) = l30
	*(*int32)(unsafe.Add(mBase, uint32(v39)+296)) = l29
	*(*int32)(unsafe.Add(mBase, uint32(v39)+292)) = l23
	*(*int32)(unsafe.Add(mBase, uint32(v39)+288)) = l31
	*(*int32)(unsafe.Add(mBase, uint32(v39)+284)) = l28
	*(*int32)(unsafe.Add(mBase, uint32(v39)+280)) = l27
	*(*int32)(unsafe.Add(mBase, uint32(v39)+276)) = l18
	*(*int32)(unsafe.Add(mBase, uint32(v39)+272)) = l17
	*(*int32)(unsafe.Add(mBase, uint32(v39)+268)) = l16
	*(*int32)(unsafe.Add(mBase, uint32(v39)+264)) = l15
	*(*int32)(unsafe.Add(mBase, uint32(v39)+260)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v39)+256)) = l13
	*(*int32)(unsafe.Add(mBase, uint32(v39)+252)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v39)+248)) = l22
	*(*int32)(unsafe.Add(mBase, uint32(v39)+244)) = l20
	*(*int32)(unsafe.Add(mBase, uint32(v39)+240)) = l19
	*(*int32)(unsafe.Add(mBase, uint32(v39)+236)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v39)+232)) = l11
	*(*int32)(unsafe.Add(mBase, uint32(v39)+228)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+224)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v39)+220)) = l9
	*(*int32)(unsafe.Add(mBase, uint32(v39)+216)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v39)+212)) = l26
	*(*int32)(unsafe.Add(mBase, uint32(v39)+208)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v39)+204)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v39)+200)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v39)+196)) = v194
	if l25 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if l24 != 0 {
		goto L60
	} else {
		goto L61
	}
L55:
	;
	v228 = F_cstring_to_text(m, l25)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L19
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v231 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+381)) = uint8(v231)
	goto L54
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+308)) = v228
	goto L54
L59:
	;
	if v171 != 0 {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v233 = F_cstring_to_text(m, l24)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L19
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v236 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+382)) = uint8(v236)
	goto L59
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+312)) = v233
	goto L59
L64:
	;
	v251 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L19
	} else {
		goto L69
	}
L65:
	;
	v245 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+383)) = uint8(v245)
	v248 = int32(0)
	goto L64
L66:
	;
	v239 = F_get_user_default_acl(m, int32(49), l6, l3)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L19
	} else {
		goto L67
	}
L67:
	;
	if v239 == int32(0) {
		goto L65
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+316)) = v239
	v248 = v239
	goto L64
L69:
	;
	v254 = F_SearchSysCacheCopy(m, int32(81), l2, l3)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L19
	} else {
		goto L71
	}
L70:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCreate[0]))
	if v314 != 0 {
		goto L92
	} else {
		goto L93
	}
L71:
	;
	if v254 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v254)+16))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+22)))
	v258 = v256 + v257
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+82)))
	if v259 == int32(1) {
		goto L3
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if l1 != 0 {
		v299 = l1
		goto L83
	} else {
		goto L84
	}
L75:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v258)+72))
	if l6 != v262 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_aclcheck_error(m, int32(2), int32(49), l2)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L19
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if l1 != 0 {
		goto L2
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v268 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+320)) = uint8(v268)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v251)+52))
	v277 = F_heap_modify_tuple(m, v254, v270, v39+int32(192), v39+int32(352), v39+int32(320))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L19
	} else {
		goto L81
	}
L81:
	;
	F_CatalogTupleUpdate(m, v251, v277+int32(4), v277)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L19
	} else {
		goto L82
	}
L82:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v310 = v283
	v311 = v277
	goto L70
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+192)) = v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v251)+52))
	v306 = F_heap_form_tuple(m, v301, v39+int32(192), v39+int32(352))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L19
	} else {
		goto L90
	}
L84:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TypeCreate[2])))
	if v285 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCreate[3]))
	if v289 == int32(0) {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v297 = F_GetNewOidWithIndex(m, v251, int32(2703), int32(1))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L19
	} else {
		goto L89
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TypeCreate[3])) = int32(0)
	v299 = v289
	goto L83
L89:
	;
	v299 = v297
	goto L83
L90:
	;
	F_CatalogTupleInsert(m, v251, v306)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L19
	} else {
		goto L91
	}
L91:
	;
	v310 = v299
	v311 = v306
	goto L70
L92:
	;
	if l25 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_TypeCreate[1]))
	if v325 != 0 {
		goto L100
	} else {
		goto L101
	}
L95:
	;
	v315 = F_stringToNode(m, l25)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L19
	} else {
		goto L98
	}
L96:
	;
	v318 = int32(0)
	goto L97
L97:
	;
	F_GenerateTypeDependencies(m, v311, v251, v318, v248, l5, l21, v171, int32(1), base.B2i32(v254 != int32(0)))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L19
	} else {
		goto L99
	}
L98:
	;
	v318 = v315
	goto L97
L99:
	;
	goto L94
L100:
	;
	v327 = int32(0)
	F_RunObjectPostCreateHook(m, int32(1247), v310, v327, v327)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L19
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
	F_relation_close(m, v251, int32(3))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L19
	} else {
		goto L104
	}
L103:
	;
	goto L102
L104:
	;
	m.G0 = v39 + int32(384)
	return
L105:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L19
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+112)) = l7
	F_errmsg(m, int32(_a_F_TypeCreate_9), v39+int32(112))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L19
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(254), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L19
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L19
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+96)) = l27
	F_errmsg(m, int32(_a_F_TypeCreate_8), v39+int32(96))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L19
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(318), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L19
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L19
	} else {
		goto L114
	}
L114:
	;
	F_errmsg(m, int32(_a_F_TypeCreate_10), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L19
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(325), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L19
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errcode(m, int32(_a_F_TypeCreate_0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L19
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = l2
	F_errmsg(m, int32(_a_F_TypeCreate_1), v39+int32(16))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L19
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(434), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L19
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_errmsg_internal(m, int32(_a_F_TypeCreate_4), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L19
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(444), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L19
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L19
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(_a_F_TypeCreate_5), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L19
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_TypeCreate_2), int32(474), int32(_a_F_TypeCreate_3))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L19
	} else {
		goto L127
	}
L127:
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
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
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
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
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
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
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v598 int32
	_ = v598
	v9 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	if l1 == v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v22 + int32(32)
	return v598
L2:
	;
	v598 = l1
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
	v598 = l1
	goto L1
L6:
	;
	goto L7
L7:
	;
	switch l3 - int32(2276) {
	case 0, 7:
		v598 = l1
		goto L1
	case 1, 2, 3, 4, 5, 6:
		goto L8
	default:
		goto L9
	}
L8:
	;
	v53 = int32(0)
	if base.B2i32(base.B2i32(l3 == int32(2277))|base.B2i32(l3 == int32(3500))|base.B2i32(l3 == int32(3831))|base.B2i32(l3 == int32(_a_F_coerce_type_0))|base.B2i32(l3 == int32(_a_F_coerce_type_1))|base.B2i32(l3 == int32(_a_F_coerce_type_2))|base.B2i32(l3 == int32(_a_F_coerce_type_3)) == v53)|base.B2i32(l2 == int32(705)) == v53 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	switch l3 - int32(_a_F_coerce_type_4) {
	case 0, 2:
		v598 = l1
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
	v598 = l1
	goto L1
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v588)+24)) = l7
	v598 = v588
	goto L1
L13:
	;
	v584 = F_makeRelabelType(m, l1, v60, int32(-1), int32(0), l6)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L17
	} else {
		goto L188
	}
L14:
	;
	v60 = F_getBaseType(m, l2)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.B2i32(l2 != int32(705))|base.B2i32(v67 != int32(7)) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	return int32(0)
L18:
	;
	if v60 != l2 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v598 = l1
	goto L1
L20:
	;
	v74 = F_palloc0(m, int32(32))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L17
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if base.B2i32(l0 == int32(0))|base.B2i32(v67 != int32(8)) != 0 {
		goto L48
	} else {
		goto L49
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = l4
	v81 = F_getBaseTypeAndTypmod(m, l3, v22+int32(28))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v84 = F_typeidType(m, v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	if v81 != int32(1186) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v89 = int32(-1)
	goto L28
L27:
	;
	v89 = v83
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v81
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92+v93)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+22)))
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97+v98)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+22)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102+v103)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+25)) = uint8(v105)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+24)) = uint8(v107)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+28)) = v109
	v112 = v22 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v112)+16)) = v112
	v118 = int32(_a_F_coerce_type_5)
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_coerce_type[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v119
	*(*int32)(unsafe.Add(mBase, _c_F_coerce_type[0])) = v22 + int32(16)
	goto L29
L29:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+22)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125+v126)+100))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v129 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v132 = int32(0)
	goto L32
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v132 = v131
	goto L32
L32:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+22)))
	v135 = v133 + v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+92))
	if v136 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v138 = v136
	goto L35
L34:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v138 = v137
	goto L35
L35:
	;
	v139 = F_OidInputFunctionCall(m, v128, v132, v138, v89)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L17
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = v139
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v142 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(8))+8))
	*(*int32)(unsafe.Add(mBase, _c_F_coerce_type[0])) = v152
	goto L41
L38:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	if v143 != int32(-1) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v146 = F_pg_detoast_datum(m, v139)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L17
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = v146
	goto L37
L41:
	;
	if l3 != v81 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v157 = F_coerce_to_domain(m, v74, v81, v155, l3, l5, l6, l7, int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L17
	} else {
		goto L45
	}
L43:
	;
	v159 = v74
	goto L44
L44:
	;
	F_ReleaseCatCache(m, v84)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L17
	} else {
		goto L46
	}
L45:
	;
	v159 = v157
	goto L44
L46:
	;
	v598 = v159
	goto L1
L47:
	;
	v197 = F_find_coercion_pathway(m, l3, l2, l5, v22+int32(8))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L17
	} else {
		goto L61
	}
L48:
	;
	v174 = v67
	goto L50
L49:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v167 == int32(0) {
		goto L47
	} else {
		goto L51
	}
L50:
	;
	if v174 != int32(31) {
		goto L47
	} else {
		goto L54
	}
L51:
	;
	v170 = m.T0[v167].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l3, l4, l7)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	if v170 != 0 {
		v598 = v170
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v174 = v172
	goto L50
L54:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v178 = F_coerce_type(m, l0, v177, l2, l3, l4, l5, l6, l7)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L17
	} else {
		goto L55
	}
L55:
	;
	v180 = F_type_is_collatable(m, l3)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L17
	} else {
		goto L56
	}
L56:
	;
	if v180 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v598 = v178
	goto L1
L58:
	;
	goto L59
L59:
	;
	v185 = F_palloc0(m, int32(16))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L17
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+4)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(31)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+8)) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+12)) = v192
	v598 = v185
	goto L1
L61:
	;
	if v197 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = l4
	v202 = F_getBaseTypeAndTypmod(m, l3, v22+int32(28))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
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
	if v197 != int32(2) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v208 = F_build_coercion_expression(m, l1, v197, v206, v202, v207, l5, l6, l7)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L17
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v217 = F_coerce_to_domain(m, l1, v202, v215, l3, l5, l6, l7, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L17
	} else {
		goto L72
	}
L69:
	;
	if v202 == l3 {
		v598 = v208
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v213 = F_coerce_to_domain(m, v208, v202, v211, l3, l5, l6, l7, int32(1))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L17
	} else {
		goto L71
	}
L71:
	;
	v598 = v213
	goto L1
L72:
	;
	if v217 != l1 {
		v598 = v217
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v222 = F_makeRelabelType(m, v217, l3, int32(-1), int32(0), l6)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L17
	} else {
		goto L74
	}
L74:
	;
	v588 = v222
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
	v226 = F_typeOrDomainTypeRelid(m, l3)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L17
	} else {
		goto L77
	}
L77:
	;
	if v226 == int32(0) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v230 = m.G0
	v232 = v230 - int32(80)
	m.G0 = v232
	*(*int32)(unsafe.Add(mBase, uint32(v232)+76)) = int32(-1)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v236 != int32(6) {
		goto L85
	} else {
		goto L86
	}
L79:
	;
	v598 = v398
	goto L1
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L17
	} else {
		goto L153
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L17
	} else {
		goto L143
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L17
	} else {
		goto L135
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L17
	} else {
		goto L128
	}
L84:
	;
	v256 = F_getBaseTypeAndTypmod(m, l3, v232+int32(76))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L17
	} else {
		goto L92
	}
L85:
	;
	if v236 != int32(36) {
		goto L83
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	if v242 != 0 {
		goto L83
	} else {
		goto L89
	}
L88:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v253 = v241
	goto L84
L89:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v246 = F_GetNSItemByRangeTablePosn(m, l0, v244, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L17
	} else {
		goto L90
	}
L90:
	;
	v249 = F_expandNSItemVars(m, l0, v246, v245, v243, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L17
	} else {
		goto L91
	}
L91:
	;
	v253 = v249
	goto L84
L92:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v232)+76))
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
	v354 = v262
	v357 = v9
	goto L99
L99:
	;
	if v354 != 0 {
		goto L80
	} else {
		goto L117
	}
L100:
	;
	v292 = v259 + v280<<(uint(int32(4))%32) + v270*int32(100)
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+111)))
	if v293 == int32(1) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v354 = v332
	v357 = v335
	goto L99
L102:
	;
	v340 = v270 + int32(1)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if v340 < v341 {
		v270 = v340
		v279 = v332
		v280 = v341
		v282 = v335
		v284 = v336
		goto L100
	} else {
		goto L116
	}
L103:
	;
	v299 = F_makeNullConst(m, int32(23), int32(-1), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
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
	v301 = F_lappend(m, v282, v299)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L17
	} else {
		goto L107
	}
L107:
	;
	v332 = v279
	v335 = v301
	v336 = v284
	goto L102
L108:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v306 = F_exprType(m, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L17
	} else {
		goto L109
	}
L109:
	;
	v309 = v292 + int32(20)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+68))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v309)+76))
	v314 = F_coerce_to_target_type(m, l0, v305, v306, v310, v311, l5, int32(2), int32(-1))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L17
	} else {
		goto L110
	}
L110:
	;
	if v314 == int32(0) {
		goto L81
	} else {
		goto L111
	}
L111:
	;
	v318 = F_lappend(m, v282, v314)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L17
	} else {
		goto L112
	}
L112:
	;
	v321 = v279 + int32(4)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	if base.Ui32(v321) < base.Ui32(v323+v324<<(uint(int32(2))%32)) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v329 = v321
	goto L115
L114:
	;
	v329 = int32(0)
	goto L115
L115:
	;
	v332 = v329
	v335 = v318
	v336 = v284 + int32(1)
	goto L102
L116:
	;
	goto L101
L117:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	if int32(0) <= v362 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_DecrTupleDescRefCount(m, v259)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L17
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v368 = F_palloc0(m, int32(24))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L17
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368)+20)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v368)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v368)+12)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v368)+8)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v368)+4)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = int32(36)
	if l3 != v256 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v379 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v368)+12)) = v379
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v232)+76))
	v384 = F_coerce_type_typmod(m, v368, v256, v381, l5, v379, l7, int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L17
	} else {
		goto L126
	}
L124:
	;
	v398 = v368
	goto L125
L125:
	;
	m.G0 = v232 + int32(80)
	goto L79
L126:
	;
	v387 = F_palloc0(m, int32(28))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L17
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+24)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v387)+20)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v387)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v387)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v387)+4)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = int32(55)
	v398 = v387
	goto L125
L128:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L17
	} else {
		goto L129
	}
L129:
	;
	v410 = F_format_type_be(m, int32(2249))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L17
	} else {
		goto L130
	}
L130:
	;
	v412 = F_format_type_be(m, l3)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L17
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+4)) = v412
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = v410
	F_errmsg(m, int32(_a_F_coerce_type_6), v232)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L17
	} else {
		goto L132
	}
L132:
	;
	F_parser_coercion_errposition(m, l0, l7, l1)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L17
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_coerce_type_7), int32(1051), int32(_a_F_coerce_type_8))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
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
	v432 = m.ExcPending
	if v432 != 0 {
		goto L17
	} else {
		goto L136
	}
L136:
	;
	v434 = F_format_type_be(m, int32(2249))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L17
	} else {
		goto L137
	}
L137:
	;
	v436 = F_format_type_be(m, l3)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L17
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+36)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v232)+32)) = v434
	F_errmsg(m, int32(_a_F_coerce_type_6), v232+int32(32))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L17
	} else {
		goto L139
	}
L139:
	;
	F_errdetail(m, int32(_a_F_coerce_type_9), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L17
	} else {
		goto L140
	}
L140:
	;
	F_parser_coercion_errposition(m, l0, l7, l1)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L17
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_coerce_type_7), int32(1090), int32(_a_F_coerce_type_8))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
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
	v462 = m.ExcPending
	if v462 != 0 {
		goto L17
	} else {
		goto L144
	}
L144:
	;
	v464 = F_format_type_be(m, int32(2249))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L17
	} else {
		goto L145
	}
L145:
	;
	v466 = F_format_type_be(m, l3)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L17
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+68)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v232)+64)) = v464
	F_errmsg(m, int32(_a_F_coerce_type_6), v232-int32(-64))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L17
	} else {
		goto L147
	}
L147:
	;
	v475 = F_format_type_be(m, v306)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L17
	} else {
		goto L148
	}
L148:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v309)+68))
	v478 = F_format_type_be(m, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L17
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+56)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v232)+52)) = v478
	*(*int32)(unsafe.Add(mBase, uint32(v232)+48)) = v475
	F_errdetail(m, int32(_a_F_coerce_type_10), v232+int32(48))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L17
	} else {
		goto L150
	}
L150:
	;
	F_parser_coercion_errposition(m, l0, l7, v305)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L17
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_coerce_type_7), int32(1111), int32(_a_F_coerce_type_8))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
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
	v501 = m.ExcPending
	if v501 != 0 {
		goto L17
	} else {
		goto L154
	}
L154:
	;
	v503 = F_format_type_be(m, int32(2249))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L17
	} else {
		goto L155
	}
L155:
	;
	v505 = F_format_type_be(m, l3)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L17
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+20)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v232)+16)) = v503
	F_errmsg(m, int32(_a_F_coerce_type_6), v232+int32(16))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L17
	} else {
		goto L157
	}
L157:
	;
	F_errdetail(m, int32(_a_F_coerce_type_11), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L17
	} else {
		goto L158
	}
L158:
	;
	F_parser_coercion_errposition(m, l0, l7, l1)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L17
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_coerce_type_7), int32(1123), int32(_a_F_coerce_type_8))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
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
	v537 = F_typeInheritsFrom(m, l2, l3)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
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
	v533 = F_is_complex_array(m, l2)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L17
	} else {
		goto L168
	}
L165:
	;
	v529 = F_typeOrDomainTypeRelid(m, l2)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L17
	} else {
		goto L166
	}
L166:
	;
	if v529 == int32(0) {
		goto L161
	} else {
		goto L167
	}
L167:
	;
	v598 = l1
	goto L1
L168:
	;
	if v533 == int32(0) {
		goto L161
	} else {
		goto L169
	}
L169:
	;
	v598 = l1
	goto L1
L170:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L17
	} else {
		goto L183
	}
L171:
	;
	if v537 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v541 = F_typeIsOfTypedTable(m, l2, l3)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L17
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v545 = F_getBaseType(m, l2)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L17
	} else {
		goto L177
	}
L175:
	;
	if v541 == int32(0) {
		goto L170
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v548 = F_palloc0(m, int32(20))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L17
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v548))) = int32(30)
	if v545 != l2 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v556 = F_makeRelabelType(m, l1, v545, int32(-1), int32(0), int32(2))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L17
	} else {
		goto L182
	}
L180:
	;
	v559 = l1
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v548)+16)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v548)+12)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v548)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v548)+4)) = v559
	v598 = v548
	goto L1
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v556)+24)) = l7
	v559 = v556
	goto L181
L183:
	;
	v568 = F_format_type_be(m, l2)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L17
	} else {
		goto L184
	}
L184:
	;
	v570 = F_format_type_be(m, l3)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L17
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v568
	F_errmsg_internal(m, int32(_a_F_coerce_type_12), v22)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L17
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_coerce_type_7), int32(544), int32(_a_F_coerce_type_13))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
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
	v588 = v584
	goto L12
}
func F_coerce_type_typmod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	v10 = F_exprTypmod(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != l2 {
			if l6 != 0 {
				F_hide_coercion_node(m, l0)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					if l2 < int32(0) {
						v59 = F_exprCollation(m, l0)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							v62 = F_applyRelabelType(m, l0, l1, l2, v59, l4, l5, int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v67 = v62
								return v67
							}
						}
					} else {
						v19 = F_typeidType(m, l1)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return int32(0)
						} else {
							v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
							v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
							v23 = v21 + v22
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
							if v24 == int32(0) {
								v35 = l1
								v37 = int32(1)
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
								v30 = base.B2i32(v28 == int32(_a_F_coerce_type_typmod_0))
								if v28 == int32(_a_F_coerce_type_typmod_0) {
									v31 = v24
								} else {
									v31 = l1
								}
								if v28 == int32(_a_F_coerce_type_typmod_0) {
									v34 = int32(3)
								} else {
									v34 = int32(1)
								}
								v35 = v31
								v37 = v34
							}
							F_ReleaseCatCache(m, v19)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v41 = F_SearchSysCache2(m, int32(12), v35, v35)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									if v41 == int32(0) {
										v59 = F_exprCollation(m, l0)
										mBase = m.M
										v60 = m.ExcPending
										if v60 != 0 {
											return int32(0)
										} else {
											v62 = F_applyRelabelType(m, l0, l1, l2, v59, l4, l5, int32(0))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												v67 = v62
												return v67
											}
										}
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
										v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+22)))
										v48 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46)+12))
										F_ReleaseCatCache(m, v41)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											if v48 == int32(0) {
												v59 = F_exprCollation(m, l0)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													v62 = F_applyRelabelType(m, l0, l1, l2, v59, l4, l5, int32(0))
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return int32(0)
													} else {
														v67 = v62
														return v67
													}
												}
											} else {
												v53 = F_build_coercion_expression(m, l0, v37, v48, l1, l2, l3, l4, l5)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int32(0)
												} else {
													return v53
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
					v59 = F_exprCollation(m, l0)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						v62 = F_applyRelabelType(m, l0, l1, l2, v59, l4, l5, int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v67 = v62
							return v67
						}
					}
				} else {
					v19 = F_typeidType(m, l1)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
						v23 = v21 + v22
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
						if v24 == int32(0) {
							v35 = l1
							v37 = int32(1)
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
							v30 = base.B2i32(v28 == int32(_a_F_coerce_type_typmod_0))
							if v28 == int32(_a_F_coerce_type_typmod_0) {
								v31 = v24
							} else {
								v31 = l1
							}
							if v28 == int32(_a_F_coerce_type_typmod_0) {
								v34 = int32(3)
							} else {
								v34 = int32(1)
							}
							v35 = v31
							v37 = v34
						}
						F_ReleaseCatCache(m, v19)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v41 = F_SearchSysCache2(m, int32(12), v35, v35)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								if v41 == int32(0) {
									v59 = F_exprCollation(m, l0)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v62 = F_applyRelabelType(m, l0, l1, l2, v59, l4, l5, int32(0))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v67 = v62
											return v67
										}
									}
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
									v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+22)))
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46)+12))
									F_ReleaseCatCache(m, v41)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										if v48 == int32(0) {
											v59 = F_exprCollation(m, l0)
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int32(0)
											} else {
												v62 = F_applyRelabelType(m, l0, l1, l2, v59, l4, l5, int32(0))
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int32(0)
												} else {
													v67 = v62
													return v67
												}
											}
										} else {
											v53 = F_build_coercion_expression(m, l0, v37, v48, l1, l2, l3, l4, l5)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												return v53
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
			v67 = l0
			return v67
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v11 = F_SearchSysCache1(m, int32(82), l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg_internal(m, int32(_a_F_fillTypeDesc_0), v7)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_fillTypeDesc_1), int32(174), int32(_a_F_fillTypeDesc_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v30 = v28 + v29
			v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+76)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v31)
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+78)))
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)) = uint8(v33)
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+128)))
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)) = uint8(v35)
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+129)))
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v37)
			F_ReleaseCatCache(m, v11)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
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
												F_errfinish(m, int32(_a_F_findTypeReceiveFunction_1), int32(2130), int32(_a_F_findTypeReceiveFunction_2))
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
														F_errfinish(m, int32(_a_F_findTypeReceiveFunction_1), int32(2137), int32(_a_F_findTypeReceiveFunction_2))
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
									F_errfinish(m, int32(_a_F_findTypeReceiveFunction_1), int32(2112), int32(_a_F_findTypeReceiveFunction_2))
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
									F_errfinish(m, int32(_a_F_findTypeReceiveFunction_1), int32(2122), int32(_a_F_findTypeReceiveFunction_2))
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
												F_errfinish(m, int32(_a_F_findTypeReceiveFunction_1), int32(2130), int32(_a_F_findTypeReceiveFunction_2))
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
														F_errfinish(m, int32(_a_F_findTypeReceiveFunction_1), int32(2137), int32(_a_F_findTypeReceiveFunction_2))
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
func F_format_type(m *base.Module, l0 int32) int32 {
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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v4 == int32(1) {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v7)
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v12 != 0 {
			v17 = int32(-1)
			v18 = int32(2)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v17 = v15
			v18 = int32(3)
		}
		v19 = F_format_type_extended(m, v11, v17, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = F_cstring_to_text(m, v19)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v23
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
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
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
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
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
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
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v823 int32
	_ = v823
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1200 int32
	_ = v1200
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1239 int32
	_ = v1239
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1255 int32
	_ = v1255
	var v1260 int32
	_ = v1260
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+124)) = l0
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[0]))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[3]))
	if v79 <= v81 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[1]))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[2]))
	v78 = v15
	v79 = v17
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = int32(1598)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = int64(1408749273092)
	v25 = int32(72)
	v26 = v9 + v25
	v28 = F_hash_create(m, int32(_a_F_lookup_type_cache_14), int32(64), v26, v25)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[0])) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = int64(34359738372)
	v39 = F_hash_create(m, int32(_a_F_lookup_type_cache_15), int32(64), v26, int32(40))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[5])) = v39
	F_CacheRegisterRelcacheCallback(m, int32(1599))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	F_CacheRegisterSyscacheCallback(m, int32(82), int32(1600), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_CacheRegisterSyscacheCallback(m, int32(14), int32(1601), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	F_CacheRegisterSyscacheCallback(m, int32(19), int32(1602), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[6]))
	if v61 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	v68 = v61
	goto L14
L14:
	;
	v71 = F_MemoryContextAlloc(m, v68, int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[6]))
	v68 = v67
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[2])) = int32(4)
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[1])) = v71
	v78 = v71
	v79 = int32(4)
	goto L1
L17:
	;
	v85 = F_repalloc(m, v78, v79<<(uint(int32(3))%32))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	v95 = v78
	v96 = v81
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[3])) = v96 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v95+v96<<(uint(int32(2))%32)))) = v104
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[0]))
	v109 = v9 + int32(124)
	v110 = int32(0)
	v112 = F_hash_search(m, v107, v109, v110, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L28
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[2])) = v79 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[1])) = v85
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[3]))
	v95 = v85
	v96 = v94
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L5
	} else {
		goto L443
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L5
	} else {
		goto L439
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L5
	} else {
		goto L435
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L5
	} else {
		goto L431
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L5
	} else {
		goto L427
	}
L26:
	;
	if l1&int32(623) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L27:
	;
	F_ReleaseCatCache(m, v219)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L5
	} else {
		goto L42
	}
L28:
	;
	if v112 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	v118 = F_SearchSysCache1(m, int32(82), v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+312)))
	if v178&int32(1) != 0 {
		v223 = v112
		goto L26
	} else {
		goto L38
	}
L32:
	;
	if v118 == int32(0) {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+22)))
	v124 = v122 + v123
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+82)))
	if v125 == int32(0) {
		goto L24
	} else {
		goto L34
	}
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[0]))
	v133 = F_hash_search(m, v129, v109, int32(1), v9+int32(123))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	base.MemoryFill(m, v133+int32(4), int32(0), int32(324))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v140
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[0]))
	v144 = F_get_hash_value(m, v143, v109)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v144
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+8)) = uint16(v147)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+10)) = uint8(v149)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+11)) = uint8(v151)
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+129)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+12)) = uint8(v153)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+13)) = uint8(v155)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v124)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+16)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v124)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+20)) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v124)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+24)) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v124)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+28)) = v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v124)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+32)) = v165
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v133)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+312)) = v167 | int32(1)
	if v155 != int32(100) {
		v217 = v133
		v219 = v118
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v173 = int32(_a_F_lookup_type_cache_4)
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+320)) = v174
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[4])) = v133
	v217 = v133
	v219 = v118
	goto L27
L38:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	v183 = F_SearchSysCache1(m, int32(82), v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	if v183 == int32(0) {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+22)))
	v189 = v187 + v188
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+82)))
	if v190 == int32(0) {
		goto L22
	} else {
		goto L41
	}
L41:
	;
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v112)+8)) = uint16(v193)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+10)) = uint8(v195)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+11)) = uint8(v197)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+129)))
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+12)) = uint8(v199)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+13)) = uint8(v201)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v189)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+16)) = v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v189)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v189)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+24)) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v189)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v189)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+32)) = v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v112)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+312)) = v213 | int32(1)
	v217 = v112
	v219 = v183
	goto L27
L42:
	;
	v223 = v217
	goto L26
L43:
	;
	if l1&int32(33) == int32(0) {
		v266 = l1
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+312)))
	if v231&int32(2) != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	v236 = F_GetDefaultOpClass(m, v234, int32(403))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+40)) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v248&int32(-123) | int32(2)
	goto L43
L47:
	;
	if v236 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v238 = F_get_opclass_family(m, v236)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L5
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v243 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+36)) = v243
	v246 = v243
	goto L46
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+36)) = v238
	v241 = F_get_opclass_input_type(m, v236)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	v246 = v241
	goto L46
L53:
	;
	if v266&int32(_a_F_lookup_type_cache_5) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L54:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+312)))
	if v259&int32(8) != 0 {
		v266 = l1
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v223)+36))
	if v264 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v265 = l1
	goto L58
L57:
	;
	v265 = l1 | int32(1024)
	goto L58
L58:
	;
	v266 = v265
	goto L53
L59:
	;
	if v266&int32(33) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L60:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+312)))
	if v271&int32(4) != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	v276 = F_GetDefaultOpClass(m, v274, int32(405))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L5
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+48)) = v286
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v288&int32(-389) | int32(4)
	goto L59
L63:
	;
	if v276 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v278 = F_get_opclass_family(m, v276)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v283 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+44)) = v283
	v286 = v283
	goto L62
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+44)) = v278
	v281 = F_get_opclass_input_type(m, v276)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v286 = v281
	goto L62
L69:
	;
	if v266&int32(2) == int32(0) {
		goto L115
	} else {
		goto L116
	}
L70:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+312)))
	if v299&int32(8) != 0 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v223)+36))
	if v302 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v223)+52))
	if v382 != v384 {
		goto L112
	} else {
		goto L113
	}
L73:
	;
	if v316 != int32(2988) {
		goto L83
	} else {
		goto L84
	}
L74:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v223)+40))
	v305 = F_get_opfamily_member(m, v302, v303, v303, int32(3))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L5
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v223)+44))
	if v308 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if v305 != 0 {
		v316 = v305
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v382 = int32(0)
	goto L72
L80:
	;
	goto L81
L81:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v223)+48))
	v314 = F_get_opfamily_member(m, v308, v312, v312, int32(1))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	v316 = v314
	goto L73
L83:
	;
	if v316 != int32(1070) {
		v382 = v316
		goto L72
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v369&int32(_a_F_lookup_type_cache_16) != 0 {
		goto L108
	} else {
		goto L109
	}
L86:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v321&int32(512) == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v327 = F_get_base_element_type(m, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L91
	}
L88:
	;
	v361 = v321
	goto L89
L89:
	;
	v382 = v361 << (uint(int32(21)) % 32) >> (uint(int32(31)) % 32) & int32(1070)
	goto L72
L90:
	;
	v359 = v357 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v359
	v361 = v359
	goto L89
L91:
	;
	if v327 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v357 = v331
	goto L90
L93:
	;
	goto L94
L94:
	;
	v333 = F_lookup_type_cache(m, v327, int32(_a_F_lookup_type_cache_17))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v333)+52))
	if v336 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v338 = v335 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v338
	v340 = v338
	goto L98
L97:
	;
	v340 = v335
	goto L98
L98:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v333)+64))
	if v341 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v343 = v340 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v343
	v345 = v343
	goto L101
L100:
	;
	v345 = v340
	goto L101
L101:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v333)+68))
	if v346 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v348 = v345 | int32(_a_F_lookup_type_cache_9)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v348
	v350 = v348
	goto L104
L103:
	;
	v350 = v345
	goto L104
L104:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v333)+72))
	if v353 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v354 = v350 | int32(_a_F_lookup_type_cache_10)
	goto L107
L106:
	;
	v354 = v350
	goto L107
L107:
	;
	v357 = v354
	goto L90
L108:
	;
	v375 = v369
	goto L110
L109:
	;
	F_cache_record_field_properties(m, v223)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L111
	}
L110:
	;
	v382 = v375 << (uint(int32(16)) % 32) >> (uint(int32(31)) % 32) & int32(2988)
	goto L72
L111:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v375 = v374
	goto L110
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+80)) = int32(0)
	goto L114
L113:
	;
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+52)) = v382
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v389&int32(-393) | int32(8)
	goto L69
L115:
	;
	if v266&int32(4) == int32(0) {
		goto L152
	} else {
		goto L153
	}
L116:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v401&int32(16) != 0 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v223)+36))
	if v404 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+56)) = v478
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v477 | int32(16)
	goto L115
L119:
	;
	v477 = v401
	v478 = int32(0)
	goto L118
L120:
	;
	goto L121
L121:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v223)+40))
	v410 = F_get_opfamily_member(m, v404, v408, v408, int32(1))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L5
	} else {
		goto L122
	}
L122:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v410 != int32(2990) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	if v410 != int32(1072) {
		v477 = v412
		v478 = v410
		goto L118
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if v412&int32(_a_F_lookup_type_cache_16) == int32(0) {
		goto L148
	} else {
		goto L149
	}
L126:
	;
	if v412&int32(512) == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v422 = F_get_base_element_type(m, v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L5
	} else {
		goto L131
	}
L128:
	;
	v455 = v412
	goto L129
L129:
	;
	v477 = v455
	v478 = v455 << (uint(int32(20)) % 32) >> (uint(int32(31)) % 32) & int32(1072)
	goto L118
L130:
	;
	v455 = v452 | int32(512)
	goto L129
L131:
	;
	if v422 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v452 = v426
	goto L130
L133:
	;
	goto L134
L134:
	;
	v428 = F_lookup_type_cache(m, v422, int32(_a_F_lookup_type_cache_17))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L5
	} else {
		goto L135
	}
L135:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v428)+52))
	if v431 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v433 = v430 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v433
	v435 = v433
	goto L138
L137:
	;
	v435 = v430
	goto L138
L138:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v428)+64))
	if v436 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v438 = v435 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v438
	v440 = v438
	goto L141
L140:
	;
	v440 = v435
	goto L141
L141:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v428)+68))
	if v441 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v443 = v440 | int32(_a_F_lookup_type_cache_9)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v443
	v445 = v443
	goto L144
L143:
	;
	v445 = v440
	goto L144
L144:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v428)+72))
	if v448 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v449 = v445 | int32(_a_F_lookup_type_cache_10)
	goto L147
L146:
	;
	v449 = v445
	goto L147
L147:
	;
	v452 = v449
	goto L130
L148:
	;
	F_cache_record_field_properties(m, v223)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L5
	} else {
		goto L151
	}
L149:
	;
	v470 = v412
	goto L150
L150:
	;
	v477 = v470
	v478 = v470 << (uint(int32(15)) % 32) >> (uint(int32(31)) % 32) & int32(2990)
	goto L118
L151:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v470 = v469
	goto L150
L152:
	;
	if v266&int32(72) == int32(0) {
		goto L189
	} else {
		goto L190
	}
L153:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v489&int32(32) != 0 {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v223)+36))
	if v492 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+60)) = v566
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v565 | int32(32)
	goto L152
L156:
	;
	v565 = v489
	v566 = int32(0)
	goto L155
L157:
	;
	goto L158
L158:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v223)+40))
	v498 = F_get_opfamily_member(m, v492, v496, v496, int32(5))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L5
	} else {
		goto L159
	}
L159:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v498 != int32(2991) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	if v498 != int32(1073) {
		v565 = v500
		v566 = v498
		goto L155
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if v500&int32(_a_F_lookup_type_cache_16) == int32(0) {
		goto L185
	} else {
		goto L186
	}
L163:
	;
	if v500&int32(512) == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v510 = F_get_base_element_type(m, v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L5
	} else {
		goto L168
	}
L165:
	;
	v543 = v500
	goto L166
L166:
	;
	v565 = v543
	v566 = v543 << (uint(int32(20)) % 32) >> (uint(int32(31)) % 32) & int32(1073)
	goto L155
L167:
	;
	v543 = v540 | int32(512)
	goto L166
L168:
	;
	if v510 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v540 = v514
	goto L167
L170:
	;
	goto L171
L171:
	;
	v516 = F_lookup_type_cache(m, v510, int32(_a_F_lookup_type_cache_17))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L5
	} else {
		goto L172
	}
L172:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v516)+52))
	if v519 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v521 = v518 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v521
	v523 = v521
	goto L175
L174:
	;
	v523 = v518
	goto L175
L175:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v516)+64))
	if v524 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v526 = v523 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v526
	v528 = v526
	goto L178
L177:
	;
	v528 = v523
	goto L178
L178:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v516)+68))
	if v529 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v531 = v528 | int32(_a_F_lookup_type_cache_9)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v531
	v533 = v531
	goto L181
L180:
	;
	v533 = v528
	goto L181
L181:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v516)+72))
	if v536 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v537 = v533 | int32(_a_F_lookup_type_cache_10)
	goto L184
L183:
	;
	v537 = v533
	goto L184
L184:
	;
	v540 = v537
	goto L167
L185:
	;
	F_cache_record_field_properties(m, v223)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L5
	} else {
		goto L188
	}
L186:
	;
	v558 = v500
	goto L187
L187:
	;
	v565 = v558
	v566 = v558 << (uint(int32(15)) % 32) >> (uint(int32(31)) % 32) & int32(2991)
	goto L155
L188:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v558 = v557
	goto L187
L189:
	;
	if v266&int32(144) == int32(0) {
		goto L229
	} else {
		goto L230
	}
L190:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+312)))
	if v577&int32(64) != 0 {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v223)+36))
	if v580 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v223)+64))
	if v653 != v655 {
		goto L226
	} else {
		goto L227
	}
L193:
	;
	v653 = int32(0)
	goto L192
L194:
	;
	goto L195
L195:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v223)+40))
	v586 = F_get_opfamily_proc(m, v580, v584, v584, int32(1))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L5
	} else {
		goto L196
	}
L196:
	;
	if v586 != int32(2987) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	if v586 != int32(382) {
		v653 = v586
		goto L192
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v640&int32(_a_F_lookup_type_cache_16) != 0 {
		goto L222
	} else {
		goto L223
	}
L200:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v592&int32(512) == int32(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v598 = F_get_base_element_type(m, v597)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L5
	} else {
		goto L205
	}
L202:
	;
	v632 = v592
	goto L203
L203:
	;
	v653 = v632 << (uint(int32(20)) % 32) >> (uint(int32(31)) % 32) & int32(382)
	goto L192
L204:
	;
	v630 = v628 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v630
	v632 = v630
	goto L203
L205:
	;
	if v598 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v628 = v602
	goto L204
L207:
	;
	goto L208
L208:
	;
	v604 = F_lookup_type_cache(m, v598, int32(_a_F_lookup_type_cache_17))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L5
	} else {
		goto L209
	}
L209:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v604)+52))
	if v607 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v609 = v606 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v609
	v611 = v609
	goto L212
L211:
	;
	v611 = v606
	goto L212
L212:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v604)+64))
	if v612 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v614 = v611 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v614
	v616 = v614
	goto L215
L214:
	;
	v616 = v611
	goto L215
L215:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v604)+68))
	if v617 != 0 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v619 = v616 | int32(_a_F_lookup_type_cache_9)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v619
	v621 = v619
	goto L218
L217:
	;
	v621 = v616
	goto L218
L218:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v604)+72))
	if v624 != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v625 = v621 | int32(_a_F_lookup_type_cache_10)
	goto L221
L220:
	;
	v625 = v621
	goto L221
L221:
	;
	v628 = v625
	goto L204
L222:
	;
	v646 = v640
	goto L224
L223:
	;
	F_cache_record_field_properties(m, v223)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L5
	} else {
		goto L225
	}
L224:
	;
	v653 = v646 << (uint(int32(15)) % 32) >> (uint(int32(31)) % 32) & int32(2987)
	goto L192
L225:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v646 = v645
	goto L224
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+108)) = int32(0)
	goto L228
L227:
	;
	goto L228
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+64)) = v653
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v660 | int32(64)
	goto L189
L229:
	;
	if v266&int32(_a_F_lookup_type_cache_6) == int32(0) {
		goto L298
	} else {
		goto L299
	}
L230:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+312)))
	if v670&int32(128) != 0 {
		goto L229
	} else {
		goto L231
	}
L231:
	;
	v673 = int32(0)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v223)+44))
	if v674 == v673 {
		v815 = v673
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v223)+68))
	if v815 != v818 {
		goto L295
	} else {
		goto L296
	}
L233:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v223)+52))
	if v677 != 0 {
		goto L239
	} else {
		goto L240
	}
L234:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v802&int32(512) != 0 {
		goto L291
	} else {
		goto L292
	}
L235:
	;
	v815 = v794 << (uint(int32(19)) % 32) >> (uint(int32(31)) % 32) & int32(3902)
	goto L232
L236:
	;
	v792 = v790 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v792
	v794 = v792
	goto L235
L237:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v773)))
	v776 = F_lookup_type_cache(m, v774, int32(_a_F_lookup_type_cache_18))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L5
	} else {
		goto L284
	}
L238:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v725&int32(512) == int32(0) {
		goto L263
	} else {
		goto L264
	}
L239:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v223)+48))
	v680 = F_get_opfamily_member(m, v674, v678, v678, int32(1))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L5
	} else {
		goto L242
	}
L240:
	;
	v685 = v674
	goto L241
L241:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v223)+48))
	v688 = F_get_opfamily_proc(m, v685, v686, v686, int32(1))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L5
	} else {
		goto L244
	}
L242:
	;
	if v680 != v677 {
		v815 = v673
		goto L232
	} else {
		goto L243
	}
L243:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v223)+44))
	v685 = v683
	goto L241
L244:
	;
	if v688 <= int32(_a_F_lookup_type_cache_19) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	if v688 == int32(626) {
		goto L238
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	if v688 == int32(_a_F_lookup_type_cache_20) {
		goto L234
	} else {
		goto L257
	}
L248:
	;
	if v688 != int32(3902) {
		v815 = v688
		goto L232
	} else {
		goto L249
	}
L249:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v696&int32(512) != 0 {
		v794 = v696
		goto L235
	} else {
		goto L250
	}
L250:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v223)+200))
	if v699 != 0 {
		v773 = v699
		goto L237
	} else {
		goto L251
	}
L251:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+13)))
	if v700 == int32(114) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	F_load_rangetype_info(m, v223)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L5
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v790 = v707
	goto L236
L255:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v223)+200))
	if v705 != 0 {
		v773 = v705
		goto L237
	} else {
		goto L256
	}
L256:
	;
	goto L254
L257:
	;
	if v688 != int32(_a_F_lookup_type_cache_21) {
		v815 = v688
		goto L232
	} else {
		goto L258
	}
L258:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v712&int32(_a_F_lookup_type_cache_16) != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v718 = v712
	goto L261
L260:
	;
	F_cache_record_field_properties(m, v223)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L5
	} else {
		goto L262
	}
L261:
	;
	v815 = v718 << (uint(int32(14)) % 32) >> (uint(int32(31)) % 32) & int32(_a_F_lookup_type_cache_21)
	goto L232
L262:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v718 = v717
	goto L261
L263:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v731 = F_get_base_element_type(m, v730)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L5
	} else {
		goto L267
	}
L264:
	;
	v765 = v725
	goto L265
L265:
	;
	v815 = v765 << (uint(int32(19)) % 32) >> (uint(int32(31)) % 32) & int32(626)
	goto L232
L266:
	;
	v763 = v761 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v763
	v765 = v763
	goto L265
L267:
	;
	if v731 == int32(0) {
		goto L268
	} else {
		goto L269
	}
L268:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v761 = v735
	goto L266
L269:
	;
	goto L270
L270:
	;
	v737 = F_lookup_type_cache(m, v731, int32(_a_F_lookup_type_cache_17))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L5
	} else {
		goto L271
	}
L271:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v737)+52))
	if v740 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v742 = v739 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v742
	v744 = v742
	goto L274
L273:
	;
	v744 = v739
	goto L274
L274:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v737)+64))
	if v745 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v747 = v744 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v747
	v749 = v747
	goto L277
L276:
	;
	v749 = v744
	goto L277
L277:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v737)+68))
	if v750 != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v752 = v749 | int32(_a_F_lookup_type_cache_9)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v752
	v754 = v752
	goto L280
L279:
	;
	v754 = v749
	goto L280
L280:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v737)+72))
	if v757 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v758 = v754 | int32(_a_F_lookup_type_cache_10)
	goto L283
L282:
	;
	v758 = v754
	goto L283
L283:
	;
	v761 = v758
	goto L266
L284:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v776)+68))
	if v779 != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v781 = v778 | int32(_a_F_lookup_type_cache_9)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v781
	v783 = v781
	goto L287
L286:
	;
	v783 = v778
	goto L287
L287:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v776)+72))
	if v786 != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v787 = v783 | int32(_a_F_lookup_type_cache_10)
	goto L290
L289:
	;
	v787 = v783
	goto L290
L290:
	;
	v790 = v787
	goto L236
L291:
	;
	v808 = v802
	goto L293
L292:
	;
	F_cache_multirange_element_properties(m, v223)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L5
	} else {
		goto L294
	}
L293:
	;
	v815 = v808 << (uint(int32(19)) % 32) >> (uint(int32(31)) % 32) & int32(_a_F_lookup_type_cache_20)
	goto L232
L294:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v808 = v807
	goto L293
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+136)) = int32(0)
	goto L297
L296:
	;
	goto L297
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+68)) = v815
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v823 | int32(128)
	goto L229
L298:
	;
	if v266&int32(32) == int32(0) {
		goto L367
	} else {
		goto L368
	}
L299:
	;
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+313)))
	if v834&int32(1) != 0 {
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v837 = int32(0)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v223)+44))
	if v838 == v837 {
		v979 = v837
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v223)+72))
	if v979 != v982 {
		goto L364
	} else {
		goto L365
	}
L302:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v223)+52))
	if v841 != 0 {
		goto L308
	} else {
		goto L309
	}
L303:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v966&int32(512) != 0 {
		goto L360
	} else {
		goto L361
	}
L304:
	;
	v979 = v958 << (uint(int32(18)) % 32) >> (uint(int32(31)) % 32) & int32(3417)
	goto L301
L305:
	;
	v956 = v954 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v956
	v958 = v956
	goto L304
L306:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v937)))
	v940 = F_lookup_type_cache(m, v938, int32(_a_F_lookup_type_cache_18))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L5
	} else {
		goto L353
	}
L307:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v889&int32(512) == int32(0) {
		goto L332
	} else {
		goto L333
	}
L308:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v223)+48))
	v844 = F_get_opfamily_member(m, v838, v842, v842, int32(1))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L5
	} else {
		goto L311
	}
L309:
	;
	v849 = v838
	goto L310
L310:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v223)+48))
	v852 = F_get_opfamily_proc(m, v849, v850, v850, int32(2))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L5
	} else {
		goto L313
	}
L311:
	;
	if v844 != v841 {
		v979 = v837
		goto L301
	} else {
		goto L312
	}
L312:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v223)+44))
	v849 = v847
	goto L310
L313:
	;
	if v852 <= int32(_a_F_lookup_type_cache_20) {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	if v852 == int32(782) {
		goto L307
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	if v852 == int32(_a_F_lookup_type_cache_22) {
		goto L303
	} else {
		goto L326
	}
L317:
	;
	if v852 != int32(3417) {
		v979 = v852
		goto L301
	} else {
		goto L318
	}
L318:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v860&int32(512) != 0 {
		v958 = v860
		goto L304
	} else {
		goto L319
	}
L319:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v223)+200))
	if v863 != 0 {
		v937 = v863
		goto L306
	} else {
		goto L320
	}
L320:
	;
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+13)))
	if v864 == int32(114) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	F_load_rangetype_info(m, v223)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L5
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v954 = v871
	goto L305
L324:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v223)+200))
	if v869 != 0 {
		v937 = v869
		goto L306
	} else {
		goto L325
	}
L325:
	;
	goto L323
L326:
	;
	if v852 != int32(_a_F_lookup_type_cache_23) {
		v979 = v852
		goto L301
	} else {
		goto L327
	}
L327:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v876&int32(_a_F_lookup_type_cache_16) != 0 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v882 = v876
	goto L330
L329:
	;
	F_cache_record_field_properties(m, v223)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L5
	} else {
		goto L331
	}
L330:
	;
	v979 = v882 << (uint(int32(13)) % 32) >> (uint(int32(31)) % 32) & int32(_a_F_lookup_type_cache_23)
	goto L301
L331:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v882 = v881
	goto L330
L332:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v895 = F_get_base_element_type(m, v894)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L5
	} else {
		goto L336
	}
L333:
	;
	v929 = v889
	goto L334
L334:
	;
	v979 = v929 << (uint(int32(18)) % 32) >> (uint(int32(31)) % 32) & int32(782)
	goto L301
L335:
	;
	v927 = v925 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v927
	v929 = v927
	goto L334
L336:
	;
	if v895 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v925 = v899
	goto L335
L338:
	;
	goto L339
L339:
	;
	v901 = F_lookup_type_cache(m, v895, int32(_a_F_lookup_type_cache_17))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L5
	} else {
		goto L340
	}
L340:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v901)+52))
	if v904 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v906 = v903 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v906
	v908 = v906
	goto L343
L342:
	;
	v908 = v903
	goto L343
L343:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v901)+64))
	if v909 != 0 {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v911 = v908 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v911
	v913 = v911
	goto L346
L345:
	;
	v913 = v908
	goto L346
L346:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v901)+68))
	if v914 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v916 = v913 | int32(_a_F_lookup_type_cache_9)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v916
	v918 = v916
	goto L349
L348:
	;
	v918 = v913
	goto L349
L349:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v901)+72))
	if v921 != 0 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v922 = v918 | int32(_a_F_lookup_type_cache_10)
	goto L352
L351:
	;
	v922 = v918
	goto L352
L352:
	;
	v925 = v922
	goto L335
L353:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v940)+68))
	if v943 != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v945 = v942 | int32(_a_F_lookup_type_cache_9)
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v945
	v947 = v945
	goto L356
L355:
	;
	v947 = v942
	goto L356
L356:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v940)+72))
	if v950 != 0 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v951 = v947 | int32(_a_F_lookup_type_cache_10)
	goto L359
L358:
	;
	v951 = v947
	goto L359
L359:
	;
	v954 = v951
	goto L305
L360:
	;
	v972 = v966
	goto L362
L361:
	;
	F_cache_multirange_element_properties(m, v223)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L5
	} else {
		goto L363
	}
L362:
	;
	v979 = v972 << (uint(int32(18)) % 32) >> (uint(int32(31)) % 32) & int32(_a_F_lookup_type_cache_22)
	goto L301
L363:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	v972 = v971
	goto L362
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+164)) = int32(0)
	goto L366
L365:
	;
	goto L366
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+72)) = v979
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+312)) = v987 | int32(256)
	goto L298
L367:
	;
	if v266&int32(64) == int32(0) {
		goto L374
	} else {
		goto L375
	}
L368:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v223)+80))
	if v998 != 0 {
		goto L367
	} else {
		goto L369
	}
L369:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v223)+52))
	if v999 == int32(0) {
		goto L367
	} else {
		goto L370
	}
L370:
	;
	v1002 = F_get_opcode(m, v999)
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L5
	} else {
		goto L371
	}
L371:
	;
	if v1002 == int32(0) {
		goto L367
	} else {
		goto L372
	}
L372:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[6]))
	F_fmgr_info_cxt(m, v1002, v223+int32(76), v1009)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L5
	} else {
		goto L373
	}
L373:
	;
	goto L367
L374:
	;
	if v266&int32(128) == int32(0) {
		goto L379
	} else {
		goto L380
	}
L375:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v223)+108))
	if v1017 != 0 {
		goto L374
	} else {
		goto L376
	}
L376:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v223)+64))
	if v1018 == int32(0) {
		goto L374
	} else {
		goto L377
	}
L377:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[6]))
	F_fmgr_info_cxt(m, v1018, v223+int32(104), v1024)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L5
	} else {
		goto L378
	}
L378:
	;
	goto L374
L379:
	;
	if v266&int32(_a_F_lookup_type_cache_7) == int32(0) {
		goto L384
	} else {
		goto L385
	}
L380:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v223)+136))
	if v1032 != 0 {
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v223)+68))
	if v1033 == int32(0) {
		goto L379
	} else {
		goto L382
	}
L382:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[6]))
	F_fmgr_info_cxt(m, v1033, v223+int32(132), v1039)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L5
	} else {
		goto L383
	}
L383:
	;
	goto L379
L384:
	;
	if v266&int32(256) == int32(0) {
		goto L389
	} else {
		goto L390
	}
L385:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v223)+164))
	if v1047 != 0 {
		goto L384
	} else {
		goto L386
	}
L386:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v223)+72))
	if v1048 == int32(0) {
		goto L384
	} else {
		goto L387
	}
L387:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[6]))
	F_fmgr_info_cxt(m, v1048, v223+int32(160), v1054)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L5
	} else {
		goto L388
	}
L388:
	;
	goto L384
L389:
	;
	if v266&int32(2048) == int32(0) {
		goto L394
	} else {
		goto L395
	}
L390:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v223)+188))
	if v1062 != 0 {
		goto L389
	} else {
		goto L391
	}
L391:
	;
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+13)))
	if v1063 != int32(99) {
		goto L389
	} else {
		goto L392
	}
L392:
	;
	F_load_typcache_tupdesc(m, v223)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L5
	} else {
		goto L393
	}
L393:
	;
	goto L389
L394:
	;
	if v266&int32(_a_F_lookup_type_cache_8) == int32(0) {
		goto L403
	} else {
		goto L404
	}
L395:
	;
	v1072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+13)))
	if v1072 != int32(114) {
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v223)+200))
	if v1075 == int32(0) {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	F_load_rangetype_info(m, v223)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L5
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+312)))
	if v1080&int32(1) != 0 {
		goto L394
	} else {
		goto L401
	}
L400:
	;
	goto L394
L401:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1075)))
	v1085 = F_lookup_type_cache(m, v1083, int32(0))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L5
	} else {
		goto L402
	}
L402:
	;
	goto L394
L403:
	;
	if v266&int32(_a_F_lookup_type_cache_9) == int32(0) {
		goto L410
	} else {
		goto L411
	}
L404:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v223)+296))
	if v1092 != 0 {
		goto L403
	} else {
		goto L405
	}
L405:
	;
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+13)))
	if v1093 != int32(109) {
		goto L403
	} else {
		goto L406
	}
L406:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v1097 = F_get_multirange_range(m, v1096)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L5
	} else {
		goto L407
	}
L407:
	;
	if v1097 == int32(0) {
		goto L21
	} else {
		goto L408
	}
L408:
	;
	v1102 = F_lookup_type_cache(m, v1097, int32(2048))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L5
	} else {
		goto L409
	}
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+296)) = v1102
	goto L403
L410:
	;
	if v266&int32(_a_F_lookup_type_cache_10) == int32(0) {
		goto L415
	} else {
		goto L416
	}
L411:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v223)+300))
	if v1110 != 0 {
		goto L410
	} else {
		goto L412
	}
L412:
	;
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+13)))
	if v1111 != int32(100) {
		goto L410
	} else {
		goto L413
	}
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+304)) = int32(-1)
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	v1119 = F_getBaseTypeAndTypmod(m, v1116, v223+int32(304))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L5
	} else {
		goto L414
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+300)) = v1119
	goto L410
L415:
	;
	v1134 = int32(_a_F_lookup_type_cache_11)
	v1136 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[3])) = v1136 - int32(1)
	v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+13)))
	if v1140 != int32(99) {
		goto L420
	} else {
		goto L421
	}
L416:
	;
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+314)))
	if v1126&int32(8) != 0 {
		goto L415
	} else {
		goto L417
	}
L417:
	;
	v1129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+13)))
	if v1129 != int32(100) {
		goto L415
	} else {
		goto L418
	}
L418:
	;
	F_load_domaintype_info(m, v223)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L5
	} else {
		goto L419
	}
L419:
	;
	goto L415
L420:
	;
	m.G0 = v9 + int32(128)
	return v223
L421:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v223)+312))
	if v1143&int32(-1572865) == int32(0) {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v223)+188))
	if v1148 == int32(0) {
		goto L420
	} else {
		goto L425
	}
L423:
	;
	goto L424
L424:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_type_cache[5]))
	v1158 = F_hash_search(m, v1152, v223+int32(16), int32(1), v9+int32(72))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L5
	} else {
		goto L426
	}
L425:
	;
	goto L424
L426:
	;
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1158))) = v1160
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	*(*int32)(unsafe.Add(mBase, uint32(v1158)+4)) = v1162
	goto L420
L427:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L5
	} else {
		goto L428
	}
L428:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v1176
	F_errmsg(m, int32(_a_F_lookup_type_cache_0), v9)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L5
	} else {
		goto L429
	}
L429:
	;
	F_errfinish(m, int32(_a_F_lookup_type_cache_1), int32(473), int32(_a_F_lookup_type_cache_2))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L5
	} else {
		goto L430
	}
L430:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L431:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L5
	} else {
		goto L432
	}
L432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v124 + int32(4)
	F_errmsg(m, int32(_a_F_lookup_type_cache_3), v9+int32(32))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L5
	} else {
		goto L433
	}
L433:
	;
	F_errfinish(m, int32(_a_F_lookup_type_cache_1), int32(479), int32(_a_F_lookup_type_cache_2))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L5
	} else {
		goto L434
	}
L434:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L435:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L5
	} else {
		goto L436
	}
L436:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v1213
	F_errmsg(m, int32(_a_F_lookup_type_cache_0), v9+int32(48))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L5
	} else {
		goto L437
	}
L437:
	;
	F_errfinish(m, int32(_a_F_lookup_type_cache_1), int32(528), int32(_a_F_lookup_type_cache_2))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L5
	} else {
		goto L438
	}
L438:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L439:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L5
	} else {
		goto L440
	}
L440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v189 + int32(4)
	F_errmsg(m, int32(_a_F_lookup_type_cache_3), v9-int32(-64))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L5
	} else {
		goto L441
	}
L441:
	;
	F_errfinish(m, int32(_a_F_lookup_type_cache_1), int32(534), int32(_a_F_lookup_type_cache_2))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L5
	} else {
		goto L442
	}
L442:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L443:
	;
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v1249
	F_errmsg_internal(m, int32(_a_F_lookup_type_cache_12), v9+int32(16))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L5
	} else {
		goto L444
	}
L444:
	;
	F_errfinish(m, int32(_a_F_lookup_type_cache_1), int32(1068), int32(_a_F_lookup_type_cache_13))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L5
	} else {
		goto L445
	}
L445:
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
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
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L59
	}
L2:
	;
	m.G0 = v12 - int32(-64)
	return v192
L3:
	;
	return int32(0)
L4:
	;
	if v14 == int32(0) {
		v192 = v3
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
		v192 = v3
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v25 = F_SearchSysCache1(m, int32(57), v20)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v25 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v30)+126)))
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v32 != int32(1) {
		v192 = v3
		goto L2
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v14
	v42 = F_list_make1_impl(m, int32(472), v10+int32(-56))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v46 = F_table_open(m, int32(2611), int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	if v42 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	F_relation_close(m, v46, int32(1))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L3
	} else {
		goto L56
	}
L15:
	;
	v170 = v160
	v177 = v167
	v179 = int32(0)
	goto L14
L16:
	;
	v160 = int32(0)
	v167 = v3
	goto L15
L17:
	;
	goto L18
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v51 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v160 = v42
	v167 = v3
	goto L15
L20:
	;
	goto L21
L21:
	;
	v54 = v42
	v59 = v3
	v61 = v3
	goto L22
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v59<<(uint(int32(2))%32))))
	v68 = int32(0)
	if v61 == v68 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v160 = v147
	v167 = v154
	goto L15
L24:
	;
	if v106 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	v106 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v74 <= int32(0) {
		v100 = v68
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v106 = v100
	goto L24
L29:
	;
	v77 = int32(0)
	if v77 < v74 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v80 = v74
	goto L32
L31:
	;
	v80 = v77
	goto L32
L32:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v83 = int32(0)
	goto L33
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v81+v83<<(uint(int32(2))%32))))
	v92 = base.B2i32(v91 == v67)
	if v91 == v67 {
		v100 = v92
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v100 = v92
	goto L28
L35:
	;
	v94 = v83 + int32(1)
	if v94 != v80 {
		v83 = v94
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v109 = F_lappend_oid(m, v61, v67)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L40
	}
L38:
	;
	v147 = v54
	v154 = v61
	goto L39
L39:
	;
	v157 = v59 + int32(1)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v157 < v158 {
		v54 = v147
		v59 = v157
		v61 = v154
		goto L22
	} else {
		goto L55
	}
L40:
	;
	v112 = v10 + int32(-52)
	F_ScanKeyInit(m, v112, int32(1), int32(3), int32(184), v67)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v119 = int32(1)
	v122 = F_systable_beginscan(m, v46, int32(2680), v119, int32(0), v119, v112)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v124 = v54
	goto L43
L43:
	;
	v133 = F_systable_getnext(m, v122)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L3
	} else {
		goto L45
	}
L44:
	;
	F_systable_endscan(m, v122)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L54
	}
L45:
	;
	if v133 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+22)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v135+v136)+4))
	if v20 == v138 {
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
	F_systable_endscan(m, v122)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v143 = F_lappend_oid(m, v124, v138)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L53
	}
L52:
	;
	v170 = v124
	v177 = v109
	v179 = int32(1)
	goto L14
L53:
	;
	v124 = v143
	goto L43
L54:
	;
	v147 = v124
	v154 = v109
	goto L39
L55:
	;
	goto L23
L56:
	;
	F_list_free(m, v177)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	F_list_free(m, v170)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v192 = v179
	goto L2
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v20
	F_errmsg_internal(m, int32(_a_F_typeInheritsFrom_0), v12)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_typeInheritsFrom_1), int32(362), int32(_a_F_typeInheritsFrom_2))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
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
