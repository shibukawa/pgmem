package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_has_server_privilege_id(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13915(m, l0, int32(_a_F_has_server_privilege_id_0), int32(1417))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_has_server_privilege_id_id(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13916(m, l0, int32(_a_F_has_server_privilege_id_id_0), int32(1417))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_has_server_privilege_id_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v14 = F_text_to_cstring(m, v6)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = F_get_foreign_server_oid(m, v14, int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v20 = F_convert_any_priv_string(m, v11, int32(_a_F_has_server_privilege_id_name_0))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = F_object_aclcheck(m, int32(1417), v17, v4, v20)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v22 == int32(0))
						}
					}
				}
			}
		}
	}
}
func F_has_server_privilege_name_id(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13917(m, l0, int32(_a_F_has_server_privilege_name_id_0), int32(1417))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_server_get_sink(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v15 = F_palloc0(m, int32(40))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v15
L2:
	;
	return int32(0)
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_server_get_sink_0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = l0
	F_StartTransactionCommand(m)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_server_get_sink[0]))
	v28 = F_has_privs_of_role(m, v26, int32(_a_F_server_get_sink_1))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L2
	} else {
		goto L87
	}
L6:
	;
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L2
	} else {
		goto L82
	}
L10:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v32 != int32(47) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v35 = F_opendir(m, l1)
	mBase = m.M
	if v35 != 0 {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	m.G0 = v12 - int32(-64)
	goto L1
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L2
	} else {
		goto L78
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L74
	}
L15:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_server_get_sink[1]))
	v158 = F_mkdir(m, l1, v157)
	mBase = m.M
	goto L68
L16:
	;
	switch v155 {
	case 0:
		goto L15
	case 1:
		goto L12
	case 2, 3, 4:
		goto L14
	default:
		goto L13
	}
L17:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_server_get_sink[2]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v127 = F_close(m, v126)
	mBase = m.M
	F_emscripten_builtin_free(m, v35)
	mBase = m.M
	goto L49
L18:
	;
	v48 = v40
	v55 = v3
	v56 = v3
	goto L26
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_server_get_sink[2])) = int32(0)
	v39 = int32(1)
	v40 = F_readdir(m, v35)
	mBase = m.M
	if v40 != 0 {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_server_get_sink[2]))
	if v44 != int32(44) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v119 = v39
	v120 = v3
	v121 = v3
	goto L17
L23:
	;
	v47 = int32(-1)
	goto L25
L24:
	;
	v47 = int32(0)
	goto L25
L25:
	;
	v155 = v47
	goto L16
L26:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+19)))
	if v57 == int32(46) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v119 = v39
	v120 = v107
	v121 = v108
	goto L17
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_server_get_sink[2])) = int32(0)
	v112 = F_readdir(m, v35)
	mBase = m.M
	if v112 != 0 {
		v48 = v112
		v55 = v107
		v56 = v108
		goto L26
	} else {
		goto L48
	}
L29:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+20)))
	if v60 == int32(0) {
		v107 = v55
		v108 = v56
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v75 = int32(_a_F_server_get_sink_2)
	v77 = v48 + int32(19)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_server_get_sink[3])))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if base.B2i32(v80 == int32(0))|base.B2i32(v80 != v83) != 0 {
		v101 = v80
		v102 = v83
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+20)))
	if v63 != int32(46) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v71 == int32(0) {
		v107 = v55
		v108 = v56
		goto L28
	} else {
		goto L37
	}
L34:
	;
	v71 = int32(46) - v63
	goto L33
L35:
	;
	goto L36
L36:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+21)))
	v71 = int32(0) - v69
	goto L33
L37:
	;
	v107 = int32(1)
	v108 = v56
	goto L28
L38:
	;
	if v101-v102 != 0 {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	goto L38
L40:
	;
	v86 = v75
	v87 = v77
	goto L41
L41:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if v91 == int32(0) {
		v101 = v91
		v102 = v90
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v101 = v91
	v102 = v90
	goto L39
L43:
	;
	v94 = int32(1)
	if v91 == v90 {
		v86 = v86 + v94
		v87 = v87 + v94
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v119 = int32(4)
	v120 = v55
	v121 = v56
	goto L17
L46:
	;
	goto L47
L47:
	;
	v107 = v55
	v108 = int32(1)
	goto L28
L48:
	;
	goto L27
L49:
	;
	if v127 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_server_get_sink[2])) = v123
	if v123 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v135 = int32(-1)
	goto L52
L52:
	;
	if v121 != 0 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v134 = int32(-1)
	goto L55
L54:
	;
	v134 = v119
	goto L55
L55:
	;
	v135 = v134
	goto L52
L56:
	;
	v138 = int32(3)
	goto L58
L57:
	;
	v138 = v135
	goto L58
L58:
	;
	if v135 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v141 = v138
	goto L61
L60:
	;
	v141 = v135
	goto L61
L61:
	;
	if v120 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v142 = int32(2)
	goto L64
L63:
	;
	v142 = v141
	goto L64
L64:
	;
	if v141 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v145 = v142
	goto L67
L66:
	;
	v145 = v141
	goto L67
L67:
	;
	v155 = v145
	goto L16
L68:
	;
	if int32(0) <= v158 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
	F_errmsg(m, int32(_a_F_server_get_sink_3), v10+int32(-48))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_server_get_sink_4), int32(102), int32(_a_F_server_get_sink_5))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errcode(m, int32(33686021))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l1
	F_errmsg(m, int32(_a_F_server_get_sink_6), v10+int32(-32))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_server_get_sink_4), int32(116), int32(_a_F_server_get_sink_5))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	F_errmsg(m, int32(_a_F_server_get_sink_7), v12)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_server_get_sink_4), int32(124), int32(_a_F_server_get_sink_5))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_server_get_sink_8), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(_a_F_server_get_sink_9)
	F_errdetail(m, int32(_a_F_server_get_sink_10), v10+int32(-16))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_server_get_sink_4), int32(75), int32(_a_F_server_get_sink_5))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L88
	}
L88:
	;
	F_errmsg(m, int32(_a_F_server_get_sink_11), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_server_get_sink_4), int32(89), int32(_a_F_server_get_sink_5))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
