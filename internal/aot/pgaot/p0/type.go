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
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[200])))
	if v4 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[303]))
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
					F_errmsg(m, int32(412866), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(494610), int32(2458), int32(437508))
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
			*(*int32)(unsafe.Add(mBase, _consts[303])) = int32(0)
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
				F_sequence_close(m, v17, int32(1))
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
	var v38 int32
	_ = v38
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
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
		v34 = l2
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
	v44 = F_pg_detoast_datum_copy(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v47 = v44
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
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v20)+100))
	if v100 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v20)+104))
	if v110 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	if v120 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v20)+112))
	if v130 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L3
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v20)+116))
	if v140 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v20)+120))
	if v150 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v20)+124))
	if v160 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	if v170 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1255)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L3
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v20)+132))
	if v180 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1247)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v20)+144))
	if v190 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L66
L68:
	;
	F_record_object_address_dependencies(m, v16+int32(16), v61, int32(110))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L72
	}
L69:
	;
	if v190 == int32(100) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(3456)
	F_add_exact_object_address(m, v16+int32(4), v61)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	F_free_object_addresses(m, v61)
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
	F_recordDependencyOnExpr(m, v16+int32(16), v34, int32(0))
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
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v278 int32
	_ = v278
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v481 int32
	_ = v481
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
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
	v550 = m.ExcPending
	if v550 != 0 {
		goto L16
	} else {
		goto L135
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L16
	} else {
		goto L129
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L16
	} else {
		goto L123
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L16
	} else {
		goto L120
	}
L5:
	;
	m.G0 = v14 + int32(160)
	return v481
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v464
	v481 = v469
	goto L5
L7:
	;
	v325 = F_SearchSysCache1(m, int32(82), v305)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L16
	} else {
		goto L88
	}
L8:
	;
	if l2 != 0 {
		v464 = int32(-1)
		v469 = v6
		goto L6
	} else {
		goto L87
	}
L9:
	;
	if v305 != 0 {
		goto L7
	} else {
		goto L86
	}
L10:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v305 = v19
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
	v187 = m.ExcPending
	if v187 != 0 {
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
	F_errmsg(m, int32(204872), v14+int32(80))
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
	F_errfinish(m, int32(498972), int32(102), int32(461629))
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
	F_errmsg(m, int32(204818), v14-int32(-64))
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
	F_errfinish(m, int32(498972), int32(124), int32(461629))
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
	F_errmsg(m, int32(71628), v14+int32(96))
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
	F_errfinish(m, int32(498972), int32(146), int32(461629))
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
		v305 = v151
		goto L9
	} else {
		goto L50
	}
L50:
	;
	F_initStringInfo(m, v14+int32(140))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L16
	} else {
		goto L51
	}
L51:
	;
	F_appendTypeNameToBuffer(m, l1, v14+int32(140))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L16
	} else {
		goto L52
	}
L52:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v14)+140))
	v168 = F_format_type_be(m, v151)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+116)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v167
	F_errmsg(m, int32(183202), v14+int32(112))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L16
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(498972), int32(159), int32(461629))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L16
	} else {
		goto L55
	}
L55:
	;
	v305 = v151
	goto L9
L56:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v14)+136))
	if v188 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v296 == int32(0) {
		v305 = v289
		goto L9
	} else {
		goto L84
	}
L58:
	;
	v190 = v14 + int32(140)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v190)+16)) = v190
	v197 = int32(4508392)
	v198 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v198
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v14 + int32(148)
	goto L61
L59:
	;
	goto L60
L60:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v14)+132))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L16
	} else {
		goto L68
	}
L61:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v14)+136))
	v205 = F_LookupExplicitNamespace(m, v204, l4)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	if v205 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v14)+132))
	v209 = int32(0)
	v211 = F_GetSysCacheOid(m, int32(81), v208, v205, v209, v209)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L16
	} else {
		goto L66
	}
L64:
	;
	v214 = int32(0)
	goto L65
L65:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(140))+8))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v218
	goto L67
L66:
	;
	v214 = v211
	goto L65
L67:
	;
	v289 = v214
	goto L57
L68:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _consts[216]))
	if v225 == int32(0) {
		v278 = int32(0)
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v289 = v278
	goto L57
L70:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if int32(0) < v228 {
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
	v278 = int32(0)
	goto L69
L74:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
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
	v259 = v238 + int32(1)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v259 < v260 {
		v238 = v259
		goto L74
	} else {
		goto L83
	}
L77:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _consts[125]))
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
	v253 = int32(0)
	v255 = F_GetSysCacheOid(m, int32(81), v220, v246, v253, v253)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L16
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	if v255 != 0 {
		v278 = v255
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
	v299 = F_get_array_type(m, v289)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L16
	} else {
		goto L85
	}
L85:
	;
	v305 = v299
	goto L9
L86:
	;
	goto L8
L87:
	;
	v481 = v6
	goto L5
L88:
	;
	if v325 == int32(0) {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v329 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if l2 == int32(0) {
		v481 = v325
		goto L5
	} else {
		goto L119
	}
L91:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v451 = v332
	goto L90
L92:
	;
	goto L93
L93:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v325)+16))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+22)))
	v335 = v333 + v334
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+82)))
	if v336 == int32(0) {
		goto L3
	} else {
		goto L94
	}
L94:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335)+116))
	if v339 == int32(0) {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	v345 = F_palloc(m, v342<<(uint(int32(2))%32))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L16
	} else {
		goto L96
	}
L96:
	;
	v347 = int32(0)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v348 == v347 {
		v410 = v347
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v418 = F_construct_array_builtin(m, v345, v410, int32(2275))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L16
	} else {
		goto L113
	}
L98:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	if v351 <= int32(0) {
		v410 = v347
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v358 = v347
	goto L100
L100:
	;
	v366 = v358 << (uint(int32(2)) % 32)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v348)+12))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v366+v367)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	switch v370 - int32(69) {
	case 0:
		goto L103
	default:
		goto L1
	case 3:
		goto L104
	}
L101:
	;
	v410 = v403
	goto L97
L102:
	;
	if v397 == int32(0) {
		goto L1
	} else {
		goto L111
	}
L103:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v384 == int32(0) {
		goto L1
	} else {
		goto L108
	}
L104:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	switch v373 - int32(465) {
	case 0:
		goto L106
	case 1, 3:
		goto L105
	default:
		goto L1
	}
L105:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v369)+8))
	v397 = v383
	goto L102
L106:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v369)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v376
	v381 = F_psprintf(m, int32(432330), v14+int32(32))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L16
	} else {
		goto L107
	}
L107:
	;
	v397 = v381
	goto L102
L108:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v384)+4))
	if v387 != int32(1) {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v384)+12))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	if v392 != int32(468) {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	v397 = v395
	goto L102
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366+v345))) = v397
	v403 = v358 + int32(1)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	if v403 < v404 {
		v358 = v403
		goto L100
	} else {
		goto L112
	}
L112:
	;
	goto L101
L113:
	;
	v421 = v14 + int32(140)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v421)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v421)+4)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v421))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v421)+16)) = v421
	v428 = int32(4508392)
	v429 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	*(*int32)(unsafe.Add(mBase, uint32(v421)+8)) = v429
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v14 + int32(148)
	goto L114
L114:
	;
	v436 = F_OidFunctionCall1Coll(m, v339, int32(0), v418)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L16
	} else {
		goto L115
	}
L115:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(140))+8))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v441
	goto L116
L116:
	;
	F_pfree(m, v345)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L16
	} else {
		goto L117
	}
L117:
	;
	F_pfree(m, v418)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L16
	} else {
		goto L118
	}
L118:
	;
	v451 = v436
	goto L90
L119:
	;
	v464 = v451
	v469 = v325
	goto L6
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v305
	F_errmsg_internal(m, int32(50356), v14)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L16
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(498972), int32(209), int32(461629))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
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
	v506 = m.ExcPending
	if v506 != 0 {
		goto L16
	} else {
		goto L124
	}
L124:
	;
	v507 = F_TypeNameToString(m, l1)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L16
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v507
	F_errmsg(m, int32(714178), v14+int32(48))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L16
	} else {
		goto L126
	}
L126:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L16
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(498972), int32(356), int32(423346))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
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
	v529 = m.ExcPending
	if v529 != 0 {
		goto L16
	} else {
		goto L130
	}
L130:
	;
	v530 = F_TypeNameToString(m, l1)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L16
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v530
	F_errmsg(m, int32(714106), v14+int32(16))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L16
	} else {
		goto L132
	}
L132:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L16
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(498972), int32(365), int32(423346))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
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
	v553 = m.ExcPending
	if v553 != 0 {
		goto L16
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(134206), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L16
	} else {
		goto L137
	}
L137:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	F_parser_errposition(m, l0, v558)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L16
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(498972), int32(410), int32(423346))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[894]))
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
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
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
	v12 = *(*int32)(unsafe.Add(mBase, _consts[895]))
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
	v95 = *(*int32)(unsafe.Add(mBase, _consts[890]))
	F_hash_seq_init(m, v8+int32(4), v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L30
	}
L5:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[894]))
	if v77 == int32(0) {
		goto L1
	} else {
		goto L23
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
	v22 = *(*int32)(unsafe.Add(mBase, _consts[890]))
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
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+13)))
	if v59 != int32(99) {
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
	v51 = v49 & int32(1572865)
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
	v47 = v45 & int32(1572865)
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
	if v57&int32(1) != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[895]))
	v71 = F_hash_search(m, v65, v27+int32(16), int32(2), v8+int32(4))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	goto L5
L23:
	;
	v81 = v77
	goto L24
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+312))
	if v85&int32(1048576) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L1
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+312)) = v85 & int32(1572865)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v81)+320))
	if v91 != 0 {
		v81 = v91
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v100 = F_hash_seq_search(m, v8+int32(4))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	if v100 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v105 = v100
	goto L33
L33:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+13)))
	switch v109 - int32(99) {
	case 0:
		goto L37
	case 1:
		goto L36
	default:
		goto L35
	}
L34:
	;
	goto L1
L35:
	;
	v171 = F_hash_seq_search(m, v8+int32(4))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L51
	}
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v105)+312))
	if v158&int32(1048576) == int32(0) {
		goto L35
	} else {
		goto L50
	}
L37:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+188))
	if v112 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v142 == int32(0) {
		goto L35
	} else {
		goto L47
	}
L39:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v115 = v113 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = v115
	if v115 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v105)+312))
	v135 = v133 & int32(1572865)
	*(*int32)(unsafe.Add(mBase, uint32(v105)+312)) = v135
	if v133&int32(-1572866) == int32(0) {
		goto L35
	} else {
		goto L46
	}
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v105)+188))
	F_FreeTupleDesc(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v105)+192)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v105)+188)) = int32(0)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v105)+312))
	v128 = v126 & int32(1572865)
	*(*int32)(unsafe.Add(mBase, uint32(v105)+312)) = v128
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+13)))
	v142 = base.B2i32(v130 == int32(99))
	v143 = v128
	goto L38
L45:
	;
	goto L44
L46:
	;
	v142 = int32(1)
	v143 = v135
	goto L38
L47:
	;
	if v143&int32(1) != 0 {
		goto L35
	} else {
		goto L48
	}
L48:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[895]))
	v156 = F_hash_search(m, v150, v105+int32(16), int32(2), v8+int32(31))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	goto L35
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+312)) = v158 & int32(1572865)
	goto L35
L51:
	;
	if v171 != 0 {
		v105 = v171
		goto L33
	} else {
		goto L52
	}
L52:
	;
	goto L34
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[890]))
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
	if v27&int32(1) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v52 = F_hash_seq_search(m, v6+int32(8))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L20
	}
L15:
	;
	if v27&int32(-1572866) != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+13)))
	if v37 != int32(99) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v26)+188))
	if v40 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[895]))
	v48 = F_hash_search(m, v42, v26+int32(16), int32(2), v6+int32(31))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	goto L14
L20:
	;
	if v52 != 0 {
		v26 = v52
		goto L12
	} else {
		goto L21
	}
L21:
	;
	goto L13
}
func F_TypeCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32, l19 int32, l20 int32, l21 int32, l22 int32, l23 int32, l24 int32, l25 int32, l26 int32, l27 int32, l28 int32, l29 int32, l30 int32, l31 int32, l32 int32) {
	mBase := m.M
	_ = mBase
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v175 int64
	_ = v175
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
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
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	v34 = int32(0)
	v36 = m.G0
	v38 = v36 - int32(384)
	m.G0 = v38
	if base.B2i32(l7 <= v34)&base.B2i32(base.Ui32(l7) <= base.Ui32(int32(-3))) == v34 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L19
	} else {
		goto L125
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L19
	} else {
		goto L122
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L19
	} else {
		goto L118
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L19
	} else {
		goto L114
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L19
	} else {
		goto L110
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
	v347 = m.ExcPending
	if v347 != 0 {
		goto L19
	} else {
		goto L106
	}
L9:
	;
	if l8 == int32(109) {
		v170 = int32(1)
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
	v118 = m.ExcPending
	if v118 != 0 {
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
	v54 = m.ExcPending
	if v54 != 0 {
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
	v57 = m.ExcPending
	if v57 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = l27
	F_errmsg(m, int32(475112), v38+int32(32))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(498962), int32(270), int32(355895))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
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
	v76 = m.ExcPending
	if v76 != 0 {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = l27
	F_errmsg(m, int32(475112), v38+int32(48))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(498962), int32(278), int32(355895))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
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
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = l27
	F_errmsg(m, int32(475112), v38-int32(-64))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(498962), int32(286), int32(355895))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
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
	v121 = m.ExcPending
	if v121 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = l7
	F_errmsg(m, int32(369589), v38)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(498962), int32(302), int32(355895))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
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
	v138 = m.ExcPending
	if v138 != 0 {
		goto L19
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = l27
	F_errmsg(m, int32(369434), v38+int32(80))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L19
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(498962), int32(312), int32(355895))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
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
	v171 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+376)) = v171
	*(*int64)(unsafe.Add(mBase, uint32(v38)+368)) = v171
	v175 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+336)) = v175
	*(*int64)(unsafe.Add(mBase, uint32(v38)+344)) = v175
	*(*int64)(unsafe.Add(mBase, uint32(v38)+360)) = v171
	*(*int64)(unsafe.Add(mBase, uint32(v38)+352)) = v171
	*(*int64)(unsafe.Add(mBase, uint32(v38)+320)) = v175
	*(*int64)(unsafe.Add(mBase, uint32(v38)+328)) = v175
	v192 = F__emscripten_memset_bulkmem(m, v38+int32(192), base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L53
L51:
	;
	if l21 != 0 {
		v170 = int32(1)
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v170 = base.B2i32(l4 != int32(0)) & base.B2i32(l5 != int32(99))
	goto L50
L53:
	;
	v196 = F_strncpy(m, v38+int32(128), l2, int32(64))
	mBase = m.M
	v197 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+63)) = uint8(v197)
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+304)) = l32
	*(*int32)(unsafe.Add(mBase, uint32(v38)+300)) = l30
	*(*int32)(unsafe.Add(mBase, uint32(v38)+296)) = l29
	*(*int32)(unsafe.Add(mBase, uint32(v38)+292)) = l23
	*(*int32)(unsafe.Add(mBase, uint32(v38)+288)) = l31
	*(*int32)(unsafe.Add(mBase, uint32(v38)+284)) = l28
	*(*int32)(unsafe.Add(mBase, uint32(v38)+280)) = l27
	*(*int32)(unsafe.Add(mBase, uint32(v38)+276)) = l18
	*(*int32)(unsafe.Add(mBase, uint32(v38)+272)) = l17
	*(*int32)(unsafe.Add(mBase, uint32(v38)+268)) = l16
	*(*int32)(unsafe.Add(mBase, uint32(v38)+264)) = l15
	*(*int32)(unsafe.Add(mBase, uint32(v38)+260)) = l14
	*(*int32)(unsafe.Add(mBase, uint32(v38)+256)) = l13
	*(*int32)(unsafe.Add(mBase, uint32(v38)+252)) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v38)+248)) = l22
	*(*int32)(unsafe.Add(mBase, uint32(v38)+244)) = l20
	*(*int32)(unsafe.Add(mBase, uint32(v38)+240)) = l19
	*(*int32)(unsafe.Add(mBase, uint32(v38)+236)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v38)+232)) = l11
	*(*int32)(unsafe.Add(mBase, uint32(v38)+228)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+224)) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v38)+220)) = l9
	*(*int32)(unsafe.Add(mBase, uint32(v38)+216)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v38)+212)) = l26
	*(*int32)(unsafe.Add(mBase, uint32(v38)+208)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v38)+204)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v38)+200)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v38)+196)) = v38 + int32(128)
	if l25 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if l24 != 0 {
		goto L61
	} else {
		goto L62
	}
L56:
	;
	v230 = F_cstring_to_text(m, l25)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L19
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v233 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+381)) = uint8(v233)
	goto L55
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+308)) = v230
	goto L55
L60:
	;
	if v170 != 0 {
		goto L66
	} else {
		goto L67
	}
L61:
	;
	v235 = F_cstring_to_text(m, l24)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L19
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v238 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+382)) = uint8(v238)
	goto L60
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+312)) = v235
	goto L60
L65:
	;
	v253 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L19
	} else {
		goto L70
	}
L66:
	;
	v247 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+383)) = uint8(v247)
	v250 = int32(0)
	goto L65
L67:
	;
	v241 = F_get_user_default_acl(m, int32(49), l6, l3)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L19
	} else {
		goto L68
	}
L68:
	;
	if v241 == int32(0) {
		goto L66
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+316)) = v241
	v250 = v241
	goto L65
L70:
	;
	v256 = F_SearchSysCacheCopy(m, int32(81), l2, l3)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L19
	} else {
		goto L72
	}
L71:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _consts[205]))
	if v317 != 0 {
		goto L93
	} else {
		goto L94
	}
L72:
	;
	if v256 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+22)))
	v260 = v258 + v259
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+82)))
	if v261 == int32(1) {
		goto L3
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if l1 != 0 {
		v301 = l1
		goto L84
	} else {
		goto L85
	}
L76:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v260)+72))
	if l6 != v264 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_aclcheck_error(m, int32(2), int32(49), l2)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L19
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if l1 != 0 {
		goto L2
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	v270 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+320)) = uint8(v270)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v253)+52))
	v279 = F_heap_modify_tuple(m, v256, v272, v38+int32(192), v38+int32(352), v38+int32(320))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L19
	} else {
		goto L82
	}
L82:
	;
	F_CatalogTupleUpdate(m, v253, v279+int32(4), v279)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L19
	} else {
		goto L83
	}
L83:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v312 = v285
	v313 = v279
	goto L71
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+192)) = v301
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v253)+52))
	v308 = F_heap_form_tuple(m, v303, v38+int32(192), v38+int32(352))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L19
	} else {
		goto L91
	}
L85:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, _consts[200])))
	if v287 == int32(1) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _consts[227]))
	if v291 == int32(0) {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v299 = F_GetNewOidWithIndex(m, v253, int32(2703), int32(1))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L19
	} else {
		goto L90
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, _consts[227])) = int32(0)
	v301 = v291
	goto L84
L90:
	;
	v301 = v299
	goto L84
L91:
	;
	F_CatalogTupleInsert(m, v253, v308)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L19
	} else {
		goto L92
	}
L92:
	;
	v312 = v301
	v313 = v308
	goto L71
L93:
	;
	if l25 != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v327 != 0 {
		goto L101
	} else {
		goto L102
	}
L96:
	;
	v318 = F_stringToNode(m, l25)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L19
	} else {
		goto L99
	}
L97:
	;
	v320 = int32(0)
	goto L98
L98:
	;
	F_GenerateTypeDependencies(m, v313, v253, v320, v250, l5, l21, v170, int32(1), base.B2i32(v256 != int32(0)))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L19
	} else {
		goto L100
	}
L99:
	;
	v320 = v318
	goto L98
L100:
	;
	goto L95
L101:
	;
	v329 = int32(0)
	F_RunObjectPostCreateHook(m, int32(1247), v312, v329, v329)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L19
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
	F_sequence_close(m, v253, int32(3))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L19
	} else {
		goto L105
	}
L104:
	;
	goto L103
L105:
	;
	m.G0 = v38 + int32(384)
	return
L106:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L19
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+112)) = l7
	F_errmsg(m, int32(475082), v38+int32(112))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L19
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(498962), int32(254), int32(355895))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L19
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L19
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = l27
	F_errmsg(m, int32(369434), v38+int32(96))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L19
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(498962), int32(318), int32(355895))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L19
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L19
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(530712), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L19
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(498962), int32(325), int32(355895))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L19
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L19
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = l2
	F_errmsg(m, int32(116985), v38+int32(16))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L19
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(498962), int32(434), int32(355895))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L19
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_errmsg_internal(m, int32(369218), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L19
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(498962), int32(444), int32(355895))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L19
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L19
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(413172), int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L19
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(498962), int32(474), int32(355895))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L19
	} else {
		goto L128
	}
L128:
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
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
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
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
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v392 int32
	_ = v392
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v594 int32
	_ = v594
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
	return v594
L2:
	;
	v594 = l1
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
	v594 = l1
	goto L1
L6:
	;
	goto L7
L7:
	;
	switch l3 - int32(2276) {
	case 0, 7:
		v594 = l1
		goto L1
	case 1, 2, 3, 4, 5, 6:
		goto L8
	default:
		goto L9
	}
L8:
	;
	if l2 == int32(705) {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	switch l3 - int32(5077) {
	case 0, 2:
		v594 = l1
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
	v594 = l1
	goto L1
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v584)+24)) = l7
	v594 = v584
	goto L1
L13:
	;
	v579 = F_makeRelabelType(m, l1, v58, int32(-1), int32(0), l6)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L17
	} else {
		goto L188
	}
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l2 != int32(705) {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	if base.B2i32(l3 == int32(2277))|base.B2i32(l3 == int32(3500))|base.B2i32(l3 == int32(3831))|base.B2i32(l3 == int32(4537))|base.B2i32(l3 == int32(5078))|base.B2i32(l3 == int32(5080))|base.B2i32(l3 == int32(4538)) == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v58 = F_getBaseType(m, l2)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	if v58 != l2 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v594 = l1
	goto L1
L20:
	;
	if l0 == int32(0) {
		v167 = v63
		goto L48
	} else {
		goto L49
	}
L21:
	;
	if v63 != int32(7) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v69 = F_palloc0(m, int32(32))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = l4
	v76 = F_getBaseTypeAndTypmod(m, l3, v23+int32(28))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v79 = F_typeidType(m, v76)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L17
	} else {
		goto L25
	}
L25:
	;
	if v76 != int32(1186) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v84 = int32(-1)
	goto L28
L27:
	;
	v84 = v78
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v76
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87+v88)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
	v95 = int32(*(*int16)(unsafe.Add(mBase, uint32(v92+v93)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+22)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v98)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v69)+25)) = uint8(v100)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v69)+24)) = uint8(v102)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+28)) = v104
	v107 = v23 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v107)+16)) = v107
	v113 = int32(4508392)
	v114 = *(*int32)(unsafe.Add(mBase, _consts[77]))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v114
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v23 + int32(16)
	goto L29
L29:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+22)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120+v121)+100))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v124 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v127 = int32(0)
	goto L32
L31:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v127 = v126
	goto L32
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+22)))
	v130 = v128 + v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+92))
	if v131 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v133 = v131
	goto L35
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v133 = v132
	goto L35
L35:
	;
	v134 = F_OidInputFunctionCall(m, v123, v127, v133, v84)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L17
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+20)) = v134
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v137 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v23+int32(8))+8))
	*(*int32)(unsafe.Add(mBase, _consts[77])) = v147
	goto L41
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	if v138 != int32(-1) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v141 = F_pg_detoast_datum(m, v134)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L17
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+20)) = v141
	goto L37
L41:
	;
	if l3 != v76 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v152 = F_coerce_to_domain(m, v69, v76, v150, l3, l5, l6, l7, int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L17
	} else {
		goto L45
	}
L43:
	;
	v154 = v69
	goto L44
L44:
	;
	F_ReleaseCatCache(m, v79)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L17
	} else {
		goto L46
	}
L45:
	;
	v154 = v152
	goto L44
L46:
	;
	v594 = v154
	goto L1
L47:
	;
	v190 = F_find_coercion_pathway(m, l3, l2, l5, v23+int32(8))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L17
	} else {
		goto L61
	}
L48:
	;
	if v167 != int32(31) {
		goto L47
	} else {
		goto L54
	}
L49:
	;
	if v63 != int32(8) {
		v167 = v63
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v161 == int32(0) {
		goto L47
	} else {
		goto L51
	}
L51:
	;
	v164 = m.T0[v161].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l3, l4, l7)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	if v164 != 0 {
		v594 = v164
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v167 = v166
	goto L48
L54:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v171 = F_coerce_type(m, l0, v170, l2, l3, l4, l5, l6, l7)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L17
	} else {
		goto L55
	}
L55:
	;
	v173 = F_type_is_collatable(m, l3)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L17
	} else {
		goto L56
	}
L56:
	;
	if v173 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v594 = v171
	goto L1
L58:
	;
	goto L59
L59:
	;
	v178 = F_palloc0(m, int32(16))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L17
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+4)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = int32(31)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v178)+8)) = v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v178)+12)) = v185
	v594 = v178
	goto L1
L61:
	;
	if v190 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = l4
	v195 = F_getBaseTypeAndTypmod(m, l3, v23+int32(28))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
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
	if v190 != int32(2) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v201 = F_build_coercion_expression(m, l1, v190, v199, v195, v200, l5, l6, l7)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L17
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v210 = F_coerce_to_domain(m, l1, v195, v208, l3, l5, l6, l7, int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L17
	} else {
		goto L72
	}
L69:
	;
	if v195 == l3 {
		v594 = v201
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v206 = F_coerce_to_domain(m, v201, v195, v204, l3, l5, l6, l7, int32(1))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L17
	} else {
		goto L71
	}
L71:
	;
	v594 = v206
	goto L1
L72:
	;
	if v210 != l1 {
		v594 = v210
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v215 = F_makeRelabelType(m, v210, l3, int32(-1), int32(0), l6)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L17
	} else {
		goto L74
	}
L74:
	;
	v584 = v215
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
	v219 = F_typeOrDomainTypeRelid(m, l3)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L17
	} else {
		goto L77
	}
L77:
	;
	if v219 == int32(0) {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v223 = m.G0
	v225 = v223 - int32(80)
	m.G0 = v225
	*(*int32)(unsafe.Add(mBase, uint32(v225)+76)) = int32(-1)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v229 != int32(6) {
		goto L85
	} else {
		goto L86
	}
L79:
	;
	v594 = v392
	goto L1
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L17
	} else {
		goto L153
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L17
	} else {
		goto L143
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L17
	} else {
		goto L135
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L17
	} else {
		goto L128
	}
L84:
	;
	v247 = int32(0)
	v250 = F_getBaseTypeAndTypmod(m, l3, v225+int32(76))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L17
	} else {
		goto L92
	}
L85:
	;
	if v229 != int32(36) {
		goto L83
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	if v235 != 0 {
		goto L83
	} else {
		goto L89
	}
L88:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v246 = v234
	goto L84
L89:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v239 = F_GetNSItemByRangeTablePosn(m, l0, v237, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L17
	} else {
		goto L90
	}
L90:
	;
	v242 = F_expandNSItemVars(m, l0, v239, v238, v236, int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L17
	} else {
		goto L91
	}
L91:
	;
	v246 = v242
	goto L84
L92:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v225)+76))
	v253 = F_lookup_rowtype_tupdesc(m, v250, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L17
	} else {
		goto L93
	}
L93:
	;
	if v246 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v256 = v255
	goto L96
L95:
	;
	v256 = v9
	goto L96
L96:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	if int32(0) < v257 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v266 = int32(0)
	v272 = v247
	v274 = v257
	v275 = v256
	v279 = int32(1)
	goto L100
L98:
	;
	v345 = v247
	v348 = v256
	goto L99
L99:
	;
	if v348 != 0 {
		goto L80
	} else {
		goto L117
	}
L100:
	;
	v289 = v253 + int32(20) + v274<<(uint(int32(4))%32) + v266*int32(100)
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+91)))
	if v290 == int32(1) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v345 = v327
	v348 = v328
	goto L99
L102:
	;
	v334 = v266 + int32(1)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	if v334 < v335 {
		v266 = v334
		v272 = v327
		v274 = v335
		v275 = v328
		v279 = v329
		goto L100
	} else {
		goto L116
	}
L103:
	;
	v296 = F_makeNullConst(m, int32(23), int32(-1), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L17
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if v275 == int32(0) {
		goto L82
	} else {
		goto L108
	}
L106:
	;
	v298 = F_lappend(m, v272, v296)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L17
	} else {
		goto L107
	}
L107:
	;
	v327 = v298
	v328 = v275
	v329 = v279
	goto L102
L108:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	v303 = F_exprType(m, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L17
	} else {
		goto L109
	}
L109:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v289)+68))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v289)+76))
	v309 = F_coerce_to_target_type(m, l0, v302, v303, v305, v306, l5, int32(2), int32(-1))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L17
	} else {
		goto L110
	}
L110:
	;
	if v309 == int32(0) {
		goto L81
	} else {
		goto L111
	}
L111:
	;
	v313 = F_lappend(m, v272, v309)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L17
	} else {
		goto L112
	}
L112:
	;
	v316 = v275 + int32(4)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	if base.Ui32(v316) < base.Ui32(v318+v319<<(uint(int32(2))%32)) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v324 = v316
	goto L115
L114:
	;
	v324 = int32(0)
	goto L115
L115:
	;
	v327 = v313
	v328 = v324
	v329 = v279 + int32(1)
	goto L102
L116:
	;
	goto L101
L117:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	if int32(0) <= v357 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_DecrTupleDescRefCount(m, v253)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L17
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v363 = F_palloc0(m, int32(24))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L17
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+20)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v363)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v363)+12)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v363)+8)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v363)+4)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v363))) = int32(36)
	if l3 != v250 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v374 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v363)+12)) = v374
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v225)+76))
	v379 = F_coerce_type_typmod(m, v363, v250, v376, l5, v374, l7, int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L17
	} else {
		goto L126
	}
L124:
	;
	v392 = v363
	goto L125
L125:
	;
	m.G0 = v225 + int32(80)
	goto L79
L126:
	;
	v382 = F_palloc0(m, int32(28))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L17
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v382)+24)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v382)+20)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v382)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v382)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v382)+4)) = v379
	*(*int32)(unsafe.Add(mBase, uint32(v382))) = int32(55)
	v392 = v382
	goto L125
L128:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L17
	} else {
		goto L129
	}
L129:
	;
	v405 = F_format_type_be(m, int32(2249))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L17
	} else {
		goto L130
	}
L130:
	;
	v407 = F_format_type_be(m, l3)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L17
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = v405
	F_errmsg(m, int32(183015), v225)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L17
	} else {
		goto L132
	}
L132:
	;
	F_parser_coercion_errposition(m, l0, l7, l1)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L17
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(499643), int32(1051), int32(27625))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
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
	v427 = m.ExcPending
	if v427 != 0 {
		goto L17
	} else {
		goto L136
	}
L136:
	;
	v429 = F_format_type_be(m, int32(2249))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L17
	} else {
		goto L137
	}
L137:
	;
	v431 = F_format_type_be(m, l3)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L17
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+36)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v225)+32)) = v429
	F_errmsg(m, int32(183015), v225+int32(32))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L17
	} else {
		goto L139
	}
L139:
	;
	F_errdetail(m, int32(592412), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L17
	} else {
		goto L140
	}
L140:
	;
	F_parser_coercion_errposition(m, l0, l7, l1)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L17
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(499643), int32(1090), int32(27625))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
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
	v457 = m.ExcPending
	if v457 != 0 {
		goto L17
	} else {
		goto L144
	}
L144:
	;
	v459 = F_format_type_be(m, int32(2249))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L17
	} else {
		goto L145
	}
L145:
	;
	v461 = F_format_type_be(m, l3)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L17
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+68)) = v461
	*(*int32)(unsafe.Add(mBase, uint32(v225)+64)) = v459
	F_errmsg(m, int32(183015), v225-int32(-64))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L17
	} else {
		goto L147
	}
L147:
	;
	v470 = F_format_type_be(m, v303)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L17
	} else {
		goto L148
	}
L148:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v289)+68))
	v473 = F_format_type_be(m, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L17
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+56)) = v279
	*(*int32)(unsafe.Add(mBase, uint32(v225)+52)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v225)+48)) = v470
	F_errdetail(m, int32(654154), v225+int32(48))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L17
	} else {
		goto L150
	}
L150:
	;
	F_parser_coercion_errposition(m, l0, l7, v302)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L17
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(499643), int32(1111), int32(27625))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
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
	v496 = m.ExcPending
	if v496 != 0 {
		goto L17
	} else {
		goto L154
	}
L154:
	;
	v498 = F_format_type_be(m, int32(2249))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L17
	} else {
		goto L155
	}
L155:
	;
	v500 = F_format_type_be(m, l3)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L17
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+20)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v225)+16)) = v498
	F_errmsg(m, int32(183015), v225+int32(16))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L17
	} else {
		goto L157
	}
L157:
	;
	F_errdetail(m, int32(592298), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L17
	} else {
		goto L158
	}
L158:
	;
	F_parser_coercion_errposition(m, l0, l7, l1)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L17
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(499643), int32(1123), int32(27625))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
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
	v532 = F_typeInheritsFrom(m, l2, l3)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
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
	v528 = F_is_complex_array(m, l2)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L17
	} else {
		goto L168
	}
L165:
	;
	v524 = F_typeOrDomainTypeRelid(m, l2)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L17
	} else {
		goto L166
	}
L166:
	;
	if v524 == int32(0) {
		goto L161
	} else {
		goto L167
	}
L167:
	;
	v594 = l1
	goto L1
L168:
	;
	if v528 == int32(0) {
		goto L161
	} else {
		goto L169
	}
L169:
	;
	v594 = l1
	goto L1
L170:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L17
	} else {
		goto L183
	}
L171:
	;
	if v532 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v536 = F_typeIsOfTypedTable(m, l2, l3)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L17
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v540 = F_getBaseType(m, l2)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L17
	} else {
		goto L177
	}
L175:
	;
	if v536 == int32(0) {
		goto L170
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v543 = F_palloc0(m, int32(20))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L17
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v543))) = int32(30)
	if v540 != l2 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v551 = F_makeRelabelType(m, l1, v540, int32(-1), int32(0), int32(2))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L17
	} else {
		goto L182
	}
L180:
	;
	v554 = l1
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v543)+16)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v543)+12)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v543)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v543)+4)) = v554
	v594 = v543
	goto L1
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v551)+24)) = l7
	v554 = v551
	goto L181
L183:
	;
	v563 = F_format_type_be(m, l2)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L17
	} else {
		goto L184
	}
L184:
	;
	v565 = F_format_type_be(m, l3)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L17
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v565
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v563
	F_errmsg_internal(m, int32(182829), v23)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L17
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(499643), int32(544), int32(366782))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
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
	v584 = v579
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
								v30 = base.B2i32(v28 == int32(6179))
								if v28 == int32(6179) {
									v31 = v24
								} else {
									v31 = l1
								}
								if v28 == int32(6179) {
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
							v30 = base.B2i32(v28 == int32(6179))
							if v28 == int32(6179) {
								v31 = v24
							} else {
								v31 = l1
							}
							if v28 == int32(6179) {
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
				F_errmsg_internal(m, int32(50356), v7)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(494053), int32(174), int32(488740))
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+60)) = int32(23)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+52)) = int64(111669151977)
	v14 = int32(1)
	v18 = F_LookupFuncName(m, l0, v14, v6+int32(-12), v14)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v26 = F_LookupFuncName(m, l0, int32(3), v6+int32(-12), int32(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				if v26 == int32(0) {
					v52 = v18
					v53 = F_get_func_rettype(m, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						if v53 != l1 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int32(0)
								} else {
									v115 = F_NameListToString(m, l0)
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										v117 = F_format_type_be(m, l1)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v117
											*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v115
											F_errmsg(m, int32(191352), v6+int32(-32))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(494610), int32(2130), int32(254773))
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
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
							v56 = F_func_volatile(m, v52)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 != int32(118) {
									m.G0 = v8 - int32(-64)
									return v52
								} else {
									v62 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										if v62 == int32(0) {
											m.G0 = v8 - int32(-64)
											return v52
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v69 = F_NameListToString(m, l0)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v69
													F_errmsg(m, int32(386140), v6+int32(-48))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(494610), int32(2137), int32(254773))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 - int32(-64)
															return v52
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
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(84439172))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = F_NameListToString(m, l0)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v37
								F_errmsg(m, int32(168883), v6+int32(-16))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(494610), int32(2112), int32(254773))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
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
				if v26 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							v97 = F_func_signature_string(m, l0, int32(1), int32(0), v6+int32(-12))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v97
								F_errmsg(m, int32(69765), v8)
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(494610), int32(2122), int32(254773))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
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
					v52 = v26
					v53 = F_get_func_rettype(m, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						if v53 != l1 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int32(0)
								} else {
									v115 = F_NameListToString(m, l0)
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										v117 = F_format_type_be(m, l1)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v117
											*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v115
											F_errmsg(m, int32(191352), v6+int32(-32))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(494610), int32(2130), int32(254773))
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
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
							v56 = F_func_volatile(m, v52)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 != int32(118) {
									m.G0 = v8 - int32(-64)
									return v52
								} else {
									v62 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										if v62 == int32(0) {
											m.G0 = v8 - int32(-64)
											return v52
										} else {
											F_errcode(m, int32(117833860))
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return int32(0)
											} else {
												v69 = F_NameListToString(m, l0)
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v69
													F_errmsg(m, int32(386140), v6+int32(-48))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(494610), int32(2137), int32(254773))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 - int32(-64)
															return v52
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
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
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v926 int32
	_ = v926
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
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
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
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
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	v8 = m.G0
	v10 = v8 - int32(128)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+124)) = l0
	v14 = *(*int32)(unsafe.Add(mBase, _consts[890]))
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _consts[893]))
	if v82 <= v84 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[891]))
	v18 = *(*int32)(unsafe.Add(mBase, _consts[892]))
	v81 = v16
	v82 = v18
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = int32(1614)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+88)) = int64(1408749273092)
	v26 = int32(72)
	v29 = F_hash_create(m, int32(399508), int32(64), v10+v26, v26)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[890])) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v10)+88)) = int64(34359738372)
	v42 = F_hash_create(m, int32(369801), int32(64), v10+int32(72), int32(40))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[895])) = v42
	F_CacheRegisterRelcacheCallback(m, int32(1615))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	F_CacheRegisterSyscacheCallback(m, int32(82), int32(1616), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_CacheRegisterSyscacheCallback(m, int32(14), int32(1617), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	F_CacheRegisterSyscacheCallback(m, int32(19), int32(1618), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	if v64 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	v71 = v64
	goto L14
L14:
	;
	v74 = F_MemoryContextAlloc(m, v71, int32(16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	v71 = v70
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[892])) = int32(4)
	*(*int32)(unsafe.Add(mBase, _consts[891])) = v74
	v81 = v74
	v82 = int32(4)
	goto L1
L17:
	;
	v88 = F_repalloc(m, v81, v82<<(uint(int32(3))%32))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	v98 = v81
	v99 = v84
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[893])) = v99 + int32(1)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v98+v99<<(uint(int32(2))%32)))) = v107
	v110 = *(*int32)(unsafe.Add(mBase, _consts[890]))
	v113 = int32(0)
	v115 = F_hash_search(m, v110, v10+int32(124), v113, v113)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L28
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[892])) = v82 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, _consts[891])) = v88
	v97 = *(*int32)(unsafe.Add(mBase, _consts[893]))
	v98 = v88
	v99 = v97
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L5
	} else {
		goto L453
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L5
	} else {
		goto L449
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L5
	} else {
		goto L445
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L5
	} else {
		goto L441
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L5
	} else {
		goto L437
	}
L26:
	;
	if l1&int32(623) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L27:
	;
	F_ReleaseCatCache(m, v250)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L5
	} else {
		goto L52
	}
L28:
	;
	if v115 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	v121 = F_SearchSysCache1(m, int32(82), v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+312)))
	if v209&int32(1) != 0 {
		v255 = v115
		goto L26
	} else {
		goto L48
	}
L32:
	;
	if v121 == int32(0) {
		goto L25
	} else {
		goto L33
	}
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+22)))
	v127 = v125 + v126
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+82)))
	if v128 == int32(0) {
		goto L24
	} else {
		goto L34
	}
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[890]))
	v138 = F_hash_search(m, v132, v10+int32(124), int32(1), v10+int32(123))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L36
	}
L35:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v169
	v172 = *(*int32)(unsafe.Add(mBase, _consts[890]))
	v175 = F_get_hash_value(m, v172, v10+int32(124))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L5
	} else {
		goto L46
	}
L36:
	;
	if v138&int32(3) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v145 = v138 + int32(328)
	if base.Ui32(v145) <= base.Ui32(v138) {
		goto L35
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v166 = F__emscripten_memset_bulkmem(m, v138+int32(4), base.I32_extend8_s(int32(0)), int32(324))
	mBase = m.M
	goto L45
L40:
	;
	v151 = v138 + int32(4)
	if base.Ui32(v151) < base.Ui32(v145) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v153 = v145
	goto L43
L42:
	;
	v153 = v151
	goto L43
L43:
	;
	v160 = F__emscripten_memset_bulkmem(m, v138, base.I32_extend8_s(int32(0)), (v138^int32(-1)+v153)&int32(-4)+int32(4))
	mBase = m.M
	goto L44
L44:
	;
	goto L35
L45:
	;
	goto L35
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v175
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+8)) = uint16(v178)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+10)) = uint8(v180)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+11)) = uint8(v182)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+129)))
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+12)) = uint8(v184)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+13)) = uint8(v186)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v127)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+16)) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v127)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+20)) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v127)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+24)) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v127)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+28)) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v127)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+32)) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v138)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+312)) = v198 | int32(1)
	if v186 != int32(100) {
		v248 = v138
		v250 = v121
		goto L27
	} else {
		goto L47
	}
L47:
	;
	v204 = int32(4508364)
	v205 = *(*int32)(unsafe.Add(mBase, _consts[894]))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+320)) = v205
	*(*int32)(unsafe.Add(mBase, _consts[894])) = v138
	v248 = v138
	v250 = v121
	goto L27
L48:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	v214 = F_SearchSysCache1(m, int32(82), v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	if v214 == int32(0) {
		goto L23
	} else {
		goto L50
	}
L50:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+22)))
	v220 = v218 + v219
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+82)))
	if v221 == int32(0) {
		goto L22
	} else {
		goto L51
	}
L51:
	;
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v115)+8)) = uint16(v224)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v115)+10)) = uint8(v226)
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v115)+11)) = uint8(v228)
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+129)))
	*(*uint8)(unsafe.Add(mBase, uint32(v115)+12)) = uint8(v230)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(v115)+13)) = uint8(v232)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v220)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+16)) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v220)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+20)) = v236
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v220)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+24)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v220)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+28)) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v220)+144))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+32)) = v242
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v115)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+312)) = v244 | int32(1)
	v248 = v115
	v250 = v214
	goto L27
L52:
	;
	v255 = v248
	goto L26
L53:
	;
	if l1&int32(33) == int32(0) {
		v299 = l1
		goto L63
	} else {
		goto L64
	}
L54:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+312)))
	if v264&int32(2) != 0 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	v269 = F_GetDefaultOpClass(m, v267, int32(403))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L5
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+40)) = v279
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v281&int32(-123) | int32(2)
	goto L53
L57:
	;
	if v269 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v271 = F_get_opclass_family(m, v269)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L5
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v276 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+36)) = v276
	v279 = v276
	goto L56
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+36)) = v271
	v274 = F_get_opclass_input_type(m, v269)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v279 = v274
	goto L56
L63:
	;
	if v299&int32(50320) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+312)))
	if v292&int32(8) != 0 {
		v299 = l1
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v255)+36))
	if v297 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v298 = l1
	goto L68
L67:
	;
	v298 = l1 | int32(1024)
	goto L68
L68:
	;
	v299 = v298
	goto L63
L69:
	;
	if v299&int32(33) == int32(0) {
		goto L79
	} else {
		goto L80
	}
L70:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+312)))
	if v304&int32(4) != 0 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	v309 = F_GetDefaultOpClass(m, v307, int32(405))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+48)) = v319
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v321&int32(-389) | int32(4)
	goto L69
L73:
	;
	if v309 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v311 = F_get_opclass_family(m, v309)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L5
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v316 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+44)) = v316
	v319 = v316
	goto L72
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+44)) = v311
	v314 = F_get_opclass_input_type(m, v309)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v319 = v314
	goto L72
L79:
	;
	if v299&int32(2) == int32(0) {
		goto L125
	} else {
		goto L126
	}
L80:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+312)))
	if v332&int32(8) != 0 {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v255)+36))
	if v335 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v255)+52))
	if v417 != v419 {
		goto L122
	} else {
		goto L123
	}
L83:
	;
	if v350 != int32(2988) {
		goto L93
	} else {
		goto L94
	}
L84:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v255)+40))
	v338 = F_get_opfamily_member(m, v335, v336, v336, int32(3))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L5
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v255)+44))
	if v342 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	if v338 != 0 {
		v350 = v338
		goto L83
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v417 = int32(0)
	goto L82
L90:
	;
	goto L91
L91:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v255)+48))
	v348 = F_get_opfamily_member(m, v342, v346, v346, int32(1))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	v350 = v348
	goto L83
L93:
	;
	if v350 != int32(1070) {
		v417 = v350
		goto L82
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v404&int32(16384) != 0 {
		goto L118
	} else {
		goto L119
	}
L96:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v356&int32(512) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v362 = F_get_base_element_type(m, v361)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L101
	}
L98:
	;
	v396 = v356
	goto L99
L99:
	;
	v417 = v396 << (uint(int32(21)) % 32) >> (uint(int32(31)) % 32) & int32(1070)
	goto L82
L100:
	;
	v394 = v392 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v394
	v396 = v394
	goto L99
L101:
	;
	if v362 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v392 = v366
	goto L100
L103:
	;
	goto L104
L104:
	;
	v368 = F_lookup_type_cache(m, v362, int32(16409))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v368)+52))
	if v371 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v373 = v370 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v373
	v375 = v373
	goto L108
L107:
	;
	v375 = v370
	goto L108
L108:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v368)+64))
	if v376 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v378 = v375 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v378
	v380 = v378
	goto L111
L110:
	;
	v380 = v375
	goto L111
L111:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v368)+68))
	if v381 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v383 = v380 | int32(4096)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v383
	v385 = v383
	goto L114
L113:
	;
	v385 = v380
	goto L114
L114:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v368)+72))
	if v388 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v389 = v385 | int32(8192)
	goto L117
L116:
	;
	v389 = v385
	goto L117
L117:
	;
	v392 = v389
	goto L100
L118:
	;
	v410 = v404
	goto L120
L119:
	;
	F_cache_record_field_properties(m, v255)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L5
	} else {
		goto L121
	}
L120:
	;
	v417 = v410 << (uint(int32(16)) % 32) >> (uint(int32(31)) % 32) & int32(2988)
	goto L82
L121:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v410 = v409
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+80)) = int32(0)
	goto L124
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+52)) = v417
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v424&int32(-393) | int32(8)
	goto L79
L125:
	;
	if v299&int32(4) == int32(0) {
		goto L162
	} else {
		goto L163
	}
L126:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v436&int32(16) != 0 {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v255)+36))
	if v439 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+56)) = v513
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v512 | int32(16)
	goto L125
L129:
	;
	v512 = v436
	v513 = int32(0)
	goto L128
L130:
	;
	goto L131
L131:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v255)+40))
	v445 = F_get_opfamily_member(m, v439, v443, v443, int32(1))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v445 != int32(2990) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	if v445 != int32(1072) {
		v512 = v447
		v513 = v445
		goto L128
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	if v447&int32(16384) == int32(0) {
		goto L158
	} else {
		goto L159
	}
L136:
	;
	if v447&int32(512) == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v457 = F_get_base_element_type(m, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L5
	} else {
		goto L141
	}
L138:
	;
	v490 = v447
	goto L139
L139:
	;
	v512 = v490
	v513 = v490 << (uint(int32(20)) % 32) >> (uint(int32(31)) % 32) & int32(1072)
	goto L128
L140:
	;
	v490 = v487 | int32(512)
	goto L139
L141:
	;
	if v457 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v487 = v461
	goto L140
L143:
	;
	goto L144
L144:
	;
	v463 = F_lookup_type_cache(m, v457, int32(16409))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L5
	} else {
		goto L145
	}
L145:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v463)+52))
	if v466 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v468 = v465 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v468
	v470 = v468
	goto L148
L147:
	;
	v470 = v465
	goto L148
L148:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v463)+64))
	if v471 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v473 = v470 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v473
	v475 = v473
	goto L151
L150:
	;
	v475 = v470
	goto L151
L151:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v463)+68))
	if v476 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v478 = v475 | int32(4096)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v478
	v480 = v478
	goto L154
L153:
	;
	v480 = v475
	goto L154
L154:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v463)+72))
	if v483 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v484 = v480 | int32(8192)
	goto L157
L156:
	;
	v484 = v480
	goto L157
L157:
	;
	v487 = v484
	goto L140
L158:
	;
	F_cache_record_field_properties(m, v255)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L5
	} else {
		goto L161
	}
L159:
	;
	v505 = v447
	goto L160
L160:
	;
	v512 = v505
	v513 = v505 << (uint(int32(15)) % 32) >> (uint(int32(31)) % 32) & int32(2990)
	goto L128
L161:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v505 = v504
	goto L160
L162:
	;
	if v299&int32(72) == int32(0) {
		goto L199
	} else {
		goto L200
	}
L163:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v524&int32(32) != 0 {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v255)+36))
	if v527 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+60)) = v601
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v600 | int32(32)
	goto L162
L166:
	;
	v600 = v524
	v601 = int32(0)
	goto L165
L167:
	;
	goto L168
L168:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v255)+40))
	v533 = F_get_opfamily_member(m, v527, v531, v531, int32(5))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L5
	} else {
		goto L169
	}
L169:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v533 != int32(2991) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	if v533 != int32(1073) {
		v600 = v535
		v601 = v533
		goto L165
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	if v535&int32(16384) == int32(0) {
		goto L195
	} else {
		goto L196
	}
L173:
	;
	if v535&int32(512) == int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v545 = F_get_base_element_type(m, v544)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L5
	} else {
		goto L178
	}
L175:
	;
	v578 = v535
	goto L176
L176:
	;
	v600 = v578
	v601 = v578 << (uint(int32(20)) % 32) >> (uint(int32(31)) % 32) & int32(1073)
	goto L165
L177:
	;
	v578 = v575 | int32(512)
	goto L176
L178:
	;
	if v545 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v575 = v549
	goto L177
L180:
	;
	goto L181
L181:
	;
	v551 = F_lookup_type_cache(m, v545, int32(16409))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L5
	} else {
		goto L182
	}
L182:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v551)+52))
	if v554 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v556 = v553 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v556
	v558 = v556
	goto L185
L184:
	;
	v558 = v553
	goto L185
L185:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v551)+64))
	if v559 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v561 = v558 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v561
	v563 = v561
	goto L188
L187:
	;
	v563 = v558
	goto L188
L188:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v551)+68))
	if v564 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v566 = v563 | int32(4096)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v566
	v568 = v566
	goto L191
L190:
	;
	v568 = v563
	goto L191
L191:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v551)+72))
	if v571 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v572 = v568 | int32(8192)
	goto L194
L193:
	;
	v572 = v568
	goto L194
L194:
	;
	v575 = v572
	goto L177
L195:
	;
	F_cache_record_field_properties(m, v255)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L5
	} else {
		goto L198
	}
L196:
	;
	v593 = v535
	goto L197
L197:
	;
	v600 = v593
	v601 = v593 << (uint(int32(15)) % 32) >> (uint(int32(31)) % 32) & int32(2991)
	goto L165
L198:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v593 = v592
	goto L197
L199:
	;
	if v299&int32(144) == int32(0) {
		goto L239
	} else {
		goto L240
	}
L200:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+312)))
	if v612&int32(64) != 0 {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v255)+36))
	if v615 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v255)+64))
	if v688 != v690 {
		goto L236
	} else {
		goto L237
	}
L203:
	;
	v688 = int32(0)
	goto L202
L204:
	;
	goto L205
L205:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v255)+40))
	v621 = F_get_opfamily_proc(m, v615, v619, v619, int32(1))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	if v621 != int32(2987) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	if v621 != int32(382) {
		v688 = v621
		goto L202
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v675&int32(16384) != 0 {
		goto L232
	} else {
		goto L233
	}
L210:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v627&int32(512) == int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v633 = F_get_base_element_type(m, v632)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L5
	} else {
		goto L215
	}
L212:
	;
	v667 = v627
	goto L213
L213:
	;
	v688 = v667 << (uint(int32(20)) % 32) >> (uint(int32(31)) % 32) & int32(382)
	goto L202
L214:
	;
	v665 = v663 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v665
	v667 = v665
	goto L213
L215:
	;
	if v633 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v663 = v637
	goto L214
L217:
	;
	goto L218
L218:
	;
	v639 = F_lookup_type_cache(m, v633, int32(16409))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L5
	} else {
		goto L219
	}
L219:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v639)+52))
	if v642 != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v644 = v641 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v644
	v646 = v644
	goto L222
L221:
	;
	v646 = v641
	goto L222
L222:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v639)+64))
	if v647 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v649 = v646 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v649
	v651 = v649
	goto L225
L224:
	;
	v651 = v646
	goto L225
L225:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v639)+68))
	if v652 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v654 = v651 | int32(4096)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v654
	v656 = v654
	goto L228
L227:
	;
	v656 = v651
	goto L228
L228:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v639)+72))
	if v659 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v660 = v656 | int32(8192)
	goto L231
L230:
	;
	v660 = v656
	goto L231
L231:
	;
	v663 = v660
	goto L214
L232:
	;
	v681 = v675
	goto L234
L233:
	;
	F_cache_record_field_properties(m, v255)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L5
	} else {
		goto L235
	}
L234:
	;
	v688 = v681 << (uint(int32(15)) % 32) >> (uint(int32(31)) % 32) & int32(2987)
	goto L202
L235:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v681 = v680
	goto L234
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+108)) = int32(0)
	goto L238
L237:
	;
	goto L238
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+64)) = v688
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v695 | int32(64)
	goto L199
L239:
	;
	if v299&int32(49152) == int32(0) {
		goto L308
	} else {
		goto L309
	}
L240:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+312)))
	if v705&int32(128) != 0 {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v708 = int32(0)
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v255)+44))
	if v709 == v708 {
		v850 = v708
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v255)+68))
	if v850 != v854 {
		goto L305
	} else {
		goto L306
	}
L243:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v255)+52))
	if v712 != 0 {
		goto L249
	} else {
		goto L250
	}
L244:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v837&int32(512) != 0 {
		goto L301
	} else {
		goto L302
	}
L245:
	;
	v850 = v829 << (uint(int32(19)) % 32) >> (uint(int32(31)) % 32) & int32(3902)
	goto L242
L246:
	;
	v827 = v825 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v827
	v829 = v827
	goto L245
L247:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v808)))
	v811 = F_lookup_type_cache(m, v809, int32(16400))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L5
	} else {
		goto L294
	}
L248:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v760&int32(512) == int32(0) {
		goto L273
	} else {
		goto L274
	}
L249:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v255)+48))
	v715 = F_get_opfamily_member(m, v709, v713, v713, int32(1))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L5
	} else {
		goto L252
	}
L250:
	;
	v720 = v709
	goto L251
L251:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v255)+48))
	v723 = F_get_opfamily_proc(m, v720, v721, v721, int32(1))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L5
	} else {
		goto L254
	}
L252:
	;
	if v715 != v712 {
		v850 = v708
		goto L242
	} else {
		goto L253
	}
L253:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v255)+44))
	v720 = v718
	goto L251
L254:
	;
	if v723 <= int32(4277) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	if v723 == int32(626) {
		goto L248
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	if v723 == int32(4278) {
		goto L244
	} else {
		goto L267
	}
L258:
	;
	if v723 != int32(3902) {
		v850 = v723
		goto L242
	} else {
		goto L259
	}
L259:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v731&int32(512) != 0 {
		v829 = v731
		goto L245
	} else {
		goto L260
	}
L260:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v255)+200))
	if v734 != 0 {
		v808 = v734
		goto L247
	} else {
		goto L261
	}
L261:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+13)))
	if v735 == int32(114) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	F_load_rangetype_info(m, v255)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L5
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v825 = v742
	goto L246
L265:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v255)+200))
	if v740 != 0 {
		v808 = v740
		goto L247
	} else {
		goto L266
	}
L266:
	;
	goto L264
L267:
	;
	if v723 != int32(6192) {
		v850 = v723
		goto L242
	} else {
		goto L268
	}
L268:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v747&int32(16384) != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v753 = v747
	goto L271
L270:
	;
	F_cache_record_field_properties(m, v255)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L5
	} else {
		goto L272
	}
L271:
	;
	v850 = v753 << (uint(int32(14)) % 32) >> (uint(int32(31)) % 32) & int32(6192)
	goto L242
L272:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v753 = v752
	goto L271
L273:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v766 = F_get_base_element_type(m, v765)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L5
	} else {
		goto L277
	}
L274:
	;
	v800 = v760
	goto L275
L275:
	;
	v850 = v800 << (uint(int32(19)) % 32) >> (uint(int32(31)) % 32) & int32(626)
	goto L242
L276:
	;
	v798 = v796 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v798
	v800 = v798
	goto L275
L277:
	;
	if v766 == int32(0) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v796 = v770
	goto L276
L279:
	;
	goto L280
L280:
	;
	v772 = F_lookup_type_cache(m, v766, int32(16409))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L5
	} else {
		goto L281
	}
L281:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v772)+52))
	if v775 != 0 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	v777 = v774 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v777
	v779 = v777
	goto L284
L283:
	;
	v779 = v774
	goto L284
L284:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v772)+64))
	if v780 != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v782 = v779 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v782
	v784 = v782
	goto L287
L286:
	;
	v784 = v779
	goto L287
L287:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v772)+68))
	if v785 != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v787 = v784 | int32(4096)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v787
	v789 = v787
	goto L290
L289:
	;
	v789 = v784
	goto L290
L290:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v772)+72))
	if v792 != 0 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v793 = v789 | int32(8192)
	goto L293
L292:
	;
	v793 = v789
	goto L293
L293:
	;
	v796 = v793
	goto L276
L294:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v811)+68))
	if v814 != 0 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v816 = v813 | int32(4096)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v816
	v818 = v816
	goto L297
L296:
	;
	v818 = v813
	goto L297
L297:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v811)+72))
	if v821 != 0 {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v822 = v818 | int32(8192)
	goto L300
L299:
	;
	v822 = v818
	goto L300
L300:
	;
	v825 = v822
	goto L246
L301:
	;
	v843 = v837
	goto L303
L302:
	;
	F_cache_multirange_element_properties(m, v255)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L5
	} else {
		goto L304
	}
L303:
	;
	v850 = v843 << (uint(int32(19)) % 32) >> (uint(int32(31)) % 32) & int32(4278)
	goto L242
L304:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v843 = v842
	goto L303
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+136)) = int32(0)
	goto L307
L306:
	;
	goto L307
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+68)) = v850
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v859 | int32(128)
	goto L239
L308:
	;
	if v299&int32(32) == int32(0) {
		goto L377
	} else {
		goto L378
	}
L309:
	;
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+313)))
	if v871&int32(1) != 0 {
		goto L308
	} else {
		goto L310
	}
L310:
	;
	v874 = int32(0)
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v255)+44))
	if v875 == v874 {
		v1016 = v874
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v255)+72))
	if v1016 != v1020 {
		goto L374
	} else {
		goto L375
	}
L312:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v255)+52))
	if v878 != 0 {
		goto L318
	} else {
		goto L319
	}
L313:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v1003&int32(512) != 0 {
		goto L370
	} else {
		goto L371
	}
L314:
	;
	v1016 = v995 << (uint(int32(18)) % 32) >> (uint(int32(31)) % 32) & int32(3417)
	goto L311
L315:
	;
	v993 = v991 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v993
	v995 = v993
	goto L314
L316:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v974)))
	v977 = F_lookup_type_cache(m, v975, int32(16400))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L5
	} else {
		goto L363
	}
L317:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v926&int32(512) == int32(0) {
		goto L342
	} else {
		goto L343
	}
L318:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v255)+48))
	v881 = F_get_opfamily_member(m, v875, v879, v879, int32(1))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L5
	} else {
		goto L321
	}
L319:
	;
	v886 = v875
	goto L320
L320:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v255)+48))
	v889 = F_get_opfamily_proc(m, v886, v887, v887, int32(2))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L5
	} else {
		goto L323
	}
L321:
	;
	if v881 != v878 {
		v1016 = v874
		goto L311
	} else {
		goto L322
	}
L322:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v255)+44))
	v886 = v884
	goto L320
L323:
	;
	if v889 <= int32(4278) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	if v889 == int32(782) {
		goto L317
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	if v889 == int32(4279) {
		goto L313
	} else {
		goto L336
	}
L327:
	;
	if v889 != int32(3417) {
		v1016 = v889
		goto L311
	} else {
		goto L328
	}
L328:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v897&int32(512) != 0 {
		v995 = v897
		goto L314
	} else {
		goto L329
	}
L329:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v255)+200))
	if v900 != 0 {
		v974 = v900
		goto L316
	} else {
		goto L330
	}
L330:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+13)))
	if v901 == int32(114) {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	F_load_rangetype_info(m, v255)
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L5
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v991 = v908
	goto L315
L334:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v255)+200))
	if v906 != 0 {
		v974 = v906
		goto L316
	} else {
		goto L335
	}
L335:
	;
	goto L333
L336:
	;
	if v889 != int32(6193) {
		v1016 = v889
		goto L311
	} else {
		goto L337
	}
L337:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v913&int32(16384) != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v919 = v913
	goto L340
L339:
	;
	F_cache_record_field_properties(m, v255)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L5
	} else {
		goto L341
	}
L340:
	;
	v1016 = v919 << (uint(int32(13)) % 32) >> (uint(int32(31)) % 32) & int32(6193)
	goto L311
L341:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v919 = v918
	goto L340
L342:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v932 = F_get_base_element_type(m, v931)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L5
	} else {
		goto L346
	}
L343:
	;
	v966 = v926
	goto L344
L344:
	;
	v1016 = v966 << (uint(int32(18)) % 32) >> (uint(int32(31)) % 32) & int32(782)
	goto L311
L345:
	;
	v964 = v962 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v964
	v966 = v964
	goto L344
L346:
	;
	if v932 == int32(0) {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v962 = v936
	goto L345
L348:
	;
	goto L349
L349:
	;
	v938 = F_lookup_type_cache(m, v932, int32(16409))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L5
	} else {
		goto L350
	}
L350:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v938)+52))
	if v941 != 0 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v943 = v940 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v943
	v945 = v943
	goto L353
L352:
	;
	v945 = v940
	goto L353
L353:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v938)+64))
	if v946 != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v948 = v945 | int32(2048)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v948
	v950 = v948
	goto L356
L355:
	;
	v950 = v945
	goto L356
L356:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v938)+68))
	if v951 != 0 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v953 = v950 | int32(4096)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v953
	v955 = v953
	goto L359
L358:
	;
	v955 = v950
	goto L359
L359:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v938)+72))
	if v958 != 0 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v959 = v955 | int32(8192)
	goto L362
L361:
	;
	v959 = v955
	goto L362
L362:
	;
	v962 = v959
	goto L345
L363:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v977)+68))
	if v980 != 0 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v982 = v979 | int32(4096)
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v982
	v984 = v982
	goto L366
L365:
	;
	v984 = v979
	goto L366
L366:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v977)+72))
	if v987 != 0 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v988 = v984 | int32(8192)
	goto L369
L368:
	;
	v988 = v984
	goto L369
L369:
	;
	v991 = v988
	goto L315
L370:
	;
	v1009 = v1003
	goto L372
L371:
	;
	F_cache_multirange_element_properties(m, v255)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L5
	} else {
		goto L373
	}
L372:
	;
	v1016 = v1009 << (uint(int32(18)) % 32) >> (uint(int32(31)) % 32) & int32(4279)
	goto L311
L373:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	v1009 = v1008
	goto L372
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+164)) = int32(0)
	goto L376
L375:
	;
	goto L376
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+72)) = v1016
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+312)) = v1025 | int32(256)
	goto L308
L377:
	;
	if v299&int32(64) == int32(0) {
		goto L384
	} else {
		goto L385
	}
L378:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v255)+80))
	if v1037 != 0 {
		goto L377
	} else {
		goto L379
	}
L379:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v255)+52))
	if v1038 == int32(0) {
		goto L377
	} else {
		goto L380
	}
L380:
	;
	v1041 = F_get_opcode(m, v1038)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L5
	} else {
		goto L381
	}
L381:
	;
	if v1041 == int32(0) {
		goto L377
	} else {
		goto L382
	}
L382:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	F_fmgr_info_cxt(m, v1041, v255+int32(76), v1048)
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L5
	} else {
		goto L383
	}
L383:
	;
	goto L377
L384:
	;
	if v299&int32(128) == int32(0) {
		goto L389
	} else {
		goto L390
	}
L385:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v255)+108))
	if v1056 != 0 {
		goto L384
	} else {
		goto L386
	}
L386:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v255)+64))
	if v1057 == int32(0) {
		goto L384
	} else {
		goto L387
	}
L387:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	F_fmgr_info_cxt(m, v1057, v255+int32(104), v1063)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L5
	} else {
		goto L388
	}
L388:
	;
	goto L384
L389:
	;
	if v299&int32(32768) == int32(0) {
		goto L394
	} else {
		goto L395
	}
L390:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v255)+136))
	if v1071 != 0 {
		goto L389
	} else {
		goto L391
	}
L391:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v255)+68))
	if v1072 == int32(0) {
		goto L389
	} else {
		goto L392
	}
L392:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	F_fmgr_info_cxt(m, v1072, v255+int32(132), v1078)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L5
	} else {
		goto L393
	}
L393:
	;
	goto L389
L394:
	;
	if v299&int32(256) == int32(0) {
		goto L399
	} else {
		goto L400
	}
L395:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v255)+164))
	if v1086 != 0 {
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v255)+72))
	if v1087 == int32(0) {
		goto L394
	} else {
		goto L397
	}
L397:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, _consts[207]))
	F_fmgr_info_cxt(m, v1087, v255+int32(160), v1093)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L5
	} else {
		goto L398
	}
L398:
	;
	goto L394
L399:
	;
	if v299&int32(2048) == int32(0) {
		goto L404
	} else {
		goto L405
	}
L400:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v255)+188))
	if v1101 != 0 {
		goto L399
	} else {
		goto L401
	}
L401:
	;
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+13)))
	if v1102 != int32(99) {
		goto L399
	} else {
		goto L402
	}
L402:
	;
	F_load_typcache_tupdesc(m, v255)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L5
	} else {
		goto L403
	}
L403:
	;
	goto L399
L404:
	;
	if v299&int32(65536) == int32(0) {
		goto L413
	} else {
		goto L414
	}
L405:
	;
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+13)))
	if v1111 != int32(114) {
		goto L404
	} else {
		goto L406
	}
L406:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v255)+200))
	if v1114 == int32(0) {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	F_load_rangetype_info(m, v255)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L5
	} else {
		goto L410
	}
L408:
	;
	goto L409
L409:
	;
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1114)+312)))
	if v1119&int32(1) != 0 {
		goto L404
	} else {
		goto L411
	}
L410:
	;
	goto L404
L411:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1114)))
	v1124 = F_lookup_type_cache(m, v1122, int32(0))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L5
	} else {
		goto L412
	}
L412:
	;
	goto L404
L413:
	;
	if v299&int32(4096) == int32(0) {
		goto L420
	} else {
		goto L421
	}
L414:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v255)+296))
	if v1131 != 0 {
		goto L413
	} else {
		goto L415
	}
L415:
	;
	v1132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+13)))
	if v1132 != int32(109) {
		goto L413
	} else {
		goto L416
	}
L416:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v1136 = F_get_multirange_range(m, v1135)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L5
	} else {
		goto L417
	}
L417:
	;
	if v1136 == int32(0) {
		goto L21
	} else {
		goto L418
	}
L418:
	;
	v1141 = F_lookup_type_cache(m, v1136, int32(2048))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L5
	} else {
		goto L419
	}
L419:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+296)) = v1141
	goto L413
L420:
	;
	if v299&int32(8192) == int32(0) {
		goto L425
	} else {
		goto L426
	}
L421:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v255)+300))
	if v1149 != 0 {
		goto L420
	} else {
		goto L422
	}
L422:
	;
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+13)))
	if v1150 != int32(100) {
		goto L420
	} else {
		goto L423
	}
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+304)) = int32(-1)
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	v1158 = F_getBaseTypeAndTypmod(m, v1155, v255+int32(304))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L5
	} else {
		goto L424
	}
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+300)) = v1158
	goto L420
L425:
	;
	v1173 = int32(4508360)
	v1175 = *(*int32)(unsafe.Add(mBase, _consts[893]))
	*(*int32)(unsafe.Add(mBase, _consts[893])) = v1175 - int32(1)
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+13)))
	if v1179 != int32(99) {
		goto L430
	} else {
		goto L431
	}
L426:
	;
	v1165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+314)))
	if v1165&int32(8) != 0 {
		goto L425
	} else {
		goto L427
	}
L427:
	;
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+13)))
	if v1168 != int32(100) {
		goto L425
	} else {
		goto L428
	}
L428:
	;
	F_load_domaintype_info(m, v255)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L5
	} else {
		goto L429
	}
L429:
	;
	goto L425
L430:
	;
	m.G0 = v10 + int32(128)
	return v255
L431:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v255)+312))
	if v1182&int32(-1572865) == int32(0) {
		goto L432
	} else {
		goto L433
	}
L432:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v255)+188))
	if v1187 == int32(0) {
		goto L430
	} else {
		goto L435
	}
L433:
	;
	goto L434
L434:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, _consts[895]))
	v1197 = F_hash_search(m, v1191, v255+int32(16), int32(1), v10+int32(72))
	mBase = m.M
	v1198 = m.ExcPending
	if v1198 != 0 {
		goto L5
	} else {
		goto L436
	}
L435:
	;
	goto L434
L436:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1197))) = v1199
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	*(*int32)(unsafe.Add(mBase, uint32(v1197)+4)) = v1201
	goto L430
L437:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L5
	} else {
		goto L438
	}
L438:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v1215
	F_errmsg(m, int32(69359), v10)
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L5
	} else {
		goto L439
	}
L439:
	;
	F_errfinish(m, int32(499349), int32(473), int32(399149))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L5
	} else {
		goto L440
	}
L440:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L441:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L5
	} else {
		goto L442
	}
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v127 + int32(4)
	F_errmsg(m, int32(304473), v10+int32(32))
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L5
	} else {
		goto L443
	}
L443:
	;
	F_errfinish(m, int32(499349), int32(479), int32(399149))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L5
	} else {
		goto L444
	}
L444:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L445:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L5
	} else {
		goto L446
	}
L446:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v1252
	F_errmsg(m, int32(69359), v10+int32(48))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L5
	} else {
		goto L447
	}
L447:
	;
	F_errfinish(m, int32(499349), int32(528), int32(399149))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L5
	} else {
		goto L448
	}
L448:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L449:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L5
	} else {
		goto L450
	}
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = v220 + int32(4)
	F_errmsg(m, int32(304473), v10-int32(-64))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L5
	} else {
		goto L451
	}
L451:
	;
	F_errfinish(m, int32(499349), int32(534), int32(399149))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L5
	} else {
		goto L452
	}
L452:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L453:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v1288
	F_errmsg_internal(m, int32(50465), v10+int32(16))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L5
	} else {
		goto L454
	}
L454:
	;
	F_errfinish(m, int32(499349), int32(1068), int32(241951))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L5
	} else {
		goto L455
	}
L455:
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
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
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
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
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
	v204 = m.ExcPending
	if v204 != 0 {
		goto L3
	} else {
		goto L59
	}
L2:
	;
	m.G0 = v12 - int32(-64)
	return v193
L3:
	;
	return int32(0)
L4:
	;
	if v14 == int32(0) {
		v193 = v3
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
		v193 = v3
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
		v193 = v3
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
	F_sequence_close(m, v46, int32(1))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L3
	} else {
		goto L56
	}
L15:
	;
	v172 = v162
	v177 = int32(0)
	v179 = v169
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
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v51 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v172 = v42
	v177 = v3
	v179 = v3
	goto L14
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
	v162 = v149
	v169 = v156
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
		v99 = v68
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v106 = v99
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
		v99 = v92
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v99 = v92
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
	v149 = v54
	v156 = v61
	goto L39
L39:
	;
	v159 = v59 + int32(1)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v159 < v160 {
		v54 = v149
		v59 = v159
		v61 = v156
		goto L22
	} else {
		goto L55
	}
L40:
	;
	F_ScanKeyInit(m, v10+int32(-52), int32(1), int32(3), int32(184), v67)
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
	v124 = F_systable_beginscan(m, v46, int32(2680), v119, int32(0), v119, v10+int32(-52))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v126 = v54
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
	v177 = int32(1)
	v179 = v109
	goto L14
L53:
	;
	v126 = v145
	goto L43
L54:
	;
	v149 = v126
	v156 = v109
	goto L39
L55:
	;
	goto L23
L56:
	;
	F_list_free(m, v179)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	F_list_free(m, v172)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v193 = v177
	goto L2
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v20
	F_errmsg_internal(m, int32(46291), v12)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(493765), int32(362), int32(130589))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
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
