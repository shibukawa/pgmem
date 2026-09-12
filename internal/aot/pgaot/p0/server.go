package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_has_server_privilege_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v16)
		v20 = *(*int32)(unsafe.Add(mBase, _consts[4]))
		v22 = F_convert_any_priv_string(m, v12, int32(1694144))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, int32(1417), v10, v20, v22, v8+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v8 + int32(16)
				return v35
			}
		}
	}
}
func F_has_server_privilege_id_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v22 = F_convert_any_priv_string(m, v14, int32(1694144))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, int32(1417), v11, v12, v22, v9+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v9 + int32(16)
				return v35
			}
		}
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
					v20 = F_convert_any_priv_string(m, v11, int32(1694144))
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
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v21 = F_get_role_oid_or_public(m, v12)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = F_convert_any_priv_string(m, v14, int32(1694144))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v28 = F_object_aclcheck_ext(m, int32(1417), v11, v21, v24, v9+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v30 == int32(1) {
						v33 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
						v37 = int32(0)
					} else {
						v37 = base.B2i32(v28 == int32(0))
					}
					m.G0 = v9 + int32(16)
					return v37
				}
			}
		}
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
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
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(796036)
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
	v26 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v28 = F_has_privs_of_role(m, v26, int32(4570))
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
		goto L88
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
		goto L83
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
		goto L79
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L75
	}
L15:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	v158 = F_mkdir(m, l1, v157)
	mBase = m.M
	goto L69
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
	v122 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v125 = F_close(m, v124)
	mBase = m.M
	F_emscripten_builtin_free(m, v35)
	mBase = m.M
	goto L50
L18:
	;
	v48 = v40
	v54 = v3
	v55 = v3
	goto L26
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(0)
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
	v44 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v44 != int32(44) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v118 = v3
	v119 = v3
	v120 = v39
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
	v118 = v106
	v119 = v107
	v120 = v39
	goto L17
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(0)
	v111 = F_readdir(m, v35)
	mBase = m.M
	if v111 != 0 {
		v48 = v111
		v54 = v106
		v55 = v107
		goto L26
	} else {
		goto L49
	}
L29:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+20)))
	if v60 == int32(0) {
		v106 = v54
		v107 = v55
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v75 = int32(445185)
	v77 = v48 + int32(19)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _consts[193])))
	if v81 == int32(0) {
		v100 = v80
		v101 = v81
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
		v106 = v54
		v107 = v55
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
	v106 = int32(1)
	v107 = v55
	goto L28
L38:
	;
	if v101-v100 != 0 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	goto L38
L40:
	;
	if v80 != v81 {
		v100 = v80
		v101 = v81
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v85 = v75
	v86 = v77
	goto L42
L42:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v90 == int32(0) {
		v100 = v89
		v101 = v90
		goto L39
	} else {
		goto L44
	}
L43:
	;
	v100 = v89
	v101 = v90
	goto L39
L44:
	;
	v93 = int32(1)
	if v89 == v90 {
		v85 = v85 + v93
		v86 = v86 + v93
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v118 = v54
	v119 = v55
	v120 = int32(4)
	goto L17
L47:
	;
	goto L48
L48:
	;
	v106 = v54
	v107 = int32(1)
	goto L28
L49:
	;
	goto L27
L50:
	;
	if v125 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v145 = int32(-1)
	goto L53
L52:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = v122
	if v118 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v155 = v145
	goto L16
L54:
	;
	v132 = int32(2)
	goto L56
L55:
	;
	v132 = int32(1)
	goto L56
L56:
	;
	if v119 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v135 = int32(3)
	goto L59
L58:
	;
	v135 = int32(1)
	goto L59
L59:
	;
	if v122 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v137 = int32(-1)
	goto L62
L61:
	;
	v137 = v120
	goto L62
L62:
	;
	if v137 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v140 = v135
	goto L65
L64:
	;
	v140 = v137
	goto L65
L65:
	;
	if v140 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v143 = v132
	goto L68
L67:
	;
	v143 = v140
	goto L68
L68:
	;
	v145 = v143
	goto L53
L69:
	;
	if int32(0) <= v158 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
	F_errmsg(m, int32(311450), v10+int32(-48))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(520514), int32(102), int32(33555))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errcode(m, int32(33686021))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l1
	F_errmsg(m, int32(8638), v10+int32(-32))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(520514), int32(116), int32(33555))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	F_errmsg(m, int32(311113), v12)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(520514), int32(124), int32(33555))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(225944), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(175608)
	F_errdetail(m, int32(638622), v10+int32(-16))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L2
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(520514), int32(75), int32(33555))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(225890), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(520514), int32(89), int32(33555))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
