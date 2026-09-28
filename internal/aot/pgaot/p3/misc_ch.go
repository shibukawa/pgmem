package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckBuiltinCryptoMode(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBuiltinCryptoMode[0]))
	if v2 == int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_F_CheckBuiltinCryptoMode_0), int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_CheckBuiltinCryptoMode_1), int32(735), int32(_a_F_CheckBuiltinCryptoMode_2))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		return
	}
}
func F_CheckElement_2(m *base.Module, l0 float32) {
	var v8 int32
	_ = v8
	Fn14209(m, l0, int32(147), int32(_a_F_CheckElement_2_0), int32(_a_F_CheckElement_2_1), int32(142), int32(_a_F_CheckElement_2_2))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_CheckSASLAuth(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
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
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v6
	v19 = v12 - int32(-64)
	F_initStringInfo(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	m.T0[v24].(func(*base.Module, int32, int32))(m, l1, v19)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_appendStringInfoChar(m, v19, int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	F_sendAuthRequest(m, int32(10), v31, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	F_pfree(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v46 = int32(1)
	v47 = v6
	goto L11
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L71
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L68
	}
L9:
	;
	m.G0 = v12 + int32(80)
	return v187
L10:
	;
	v187 = int32(-1)
	goto L9
L11:
	;
	F_pq_startmsgread(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	switch v135 - int32(1) {
	case 0:
		v187 = v173
		goto L9
	default:
		goto L10
	case 2:
		goto L66
	}
L13:
	;
	v50 = F_pq_getbyte(m)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v50 != int32(112) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v50 == int32(-1) {
		v187 = int32(-2)
		goto L9
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v74 = v12 + int32(48)
	F_initStringInfo(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v50
	F_errmsg(m, int32(_a_F_CheckSASLAuth_0), v12)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(96), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v78 = F_pq_getmessage(m, v74, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v78 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	F_pfree(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v85 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L10
L29:
	;
	if v85 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v87
	F_errmsg_internal(m, int32(_a_F_CheckSASLAuth_3), v12+int32(32))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v100 = v12 + int32(48)
	if v46&int32(1) != 0 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(111), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	F_pq_getmsgend(m, v12+int32(48))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L47
	}
L36:
	;
	v120 = F_pq_getmsgbytes(m, v100, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L46
	}
L37:
	;
	v104 = F_pq_getmsgrawstring(m, v100)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v118 = v47
	v119 = v115
	goto L36
L40:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v107 = m.T0[v106].(func(*base.Module, int32, int32, int32) int32)(m, l1, v104, l2)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v110 = F_pq_getmsgint(m, v100, int32(4))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v110 != int32(-1) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v118 = v107
	v119 = v110
	goto L36
L44:
	;
	goto L45
L45:
	;
	v123 = int32(-1)
	v124 = v107
	v125 = int32(0)
	goto L35
L46:
	;
	v123 = v119
	v124 = v118
	v125 = v120
	goto L35
L47:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v135 = m.T0[v134].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v124, v125, v123, v12+int32(44), v12+int32(40), l3)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	F_pfree(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v140 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v135&int32(-2) == int32(2) {
		goto L8
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v173 = int32(0)
	if v135 == v173 {
		v46 = v173
		v47 = v124
		goto L11
	} else {
		goto L65
	}
L53:
	;
	v147 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v147 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v149
	F_errmsg_internal(m, int32(_a_F_CheckSASLAuth_4), v12+int32(16))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if v135 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(182), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v165 = int32(12)
	goto L62
L61:
	;
	v165 = int32(11)
	goto L62
L62:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	F_sendAuthRequest(m, v165, v166, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	F_pfree(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L52
L65:
	;
	goto L12
L66:
	;
	if l4 == int32(0) {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	v180 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v180)
	goto L10
L68:
	;
	F_errmsg_internal(m, int32(_a_F_CheckSASLAuth_5), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(177), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errmsg_internal(m, int32(_a_F_CheckSASLAuth_6), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(201), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_char2wchar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	if l1 == int32(0) {
		return
	} else {
		v9 = F_pnstrdup(m, l2, l3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			if l4 == int32(0) {
				v13 = F_mbstowcs(m, l0, v9, l1)
				mBase = m.M
				v41 = v13
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_char2wchar[0]))
				if l4 != 0 {
					if l4 == int32(-1) {
						v21 = int32(_a_F_char2wchar_0)
					} else {
						v21 = l4
					}
					*(*int32)(unsafe.Add(mBase, _c_F_char2wchar[0])) = v21
				} else {
				}
				if v16 == int32(_a_F_char2wchar_0) {
					v26 = int32(-1)
				} else {
					v26 = v16
				}
				v27 = F_mbstowcs(m, l0, v9, l1)
				mBase = m.M
				if v26 != 0 {
					if v26 == int32(-1) {
						v35 = int32(_a_F_char2wchar_0)
					} else {
						v35 = v26
					}
					*(*int32)(unsafe.Add(mBase, _c_F_char2wchar[0])) = v35
				} else {
				}
				v41 = v27
			}
			F_pfree(m, v9)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				if v41 != int32(-1) {
					return
				} else {
					F_pg_verifymbstr(m, l2, l3)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_errcode(m, int32(17301634))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_char2wchar_1), int32(0))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_char2wchar_2), int32(0))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_char2wchar_3), int32(1367), int32(_a_F_char2wchar_4))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
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
				}
			}
		}
	}
}
func F_charne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 != v3))
}
func F_checkWellFormedRecursionWalker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
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
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
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
	var v308 int32
	_ = v308
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
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
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
	var v394 int32
	_ = v394
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v455 int32
	_ = v455
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v538 int32
	_ = v538
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	if l0 == v3 {
		v538 = v3
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v16 + int32(48)
	return base.B2i32(v593 == int32(0)) & v594
L2:
	;
	v584 = F_raw_expression_tree_walker_impl(m, v570, int32(520), l1)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L19
	} else {
		goto L169
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L19
	} else {
		goto L164
	}
L4:
	;
	v593 = v538
	v594 = v3
	goto L1
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 <= int32(109) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v338)+8))
	if v351 != 0 {
		v593 = v345
		v594 = v3
		goto L1
	} else {
		goto L122
	}
L7:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v197)+64))
	if v210 != 0 {
		goto L91
	} else {
		goto L92
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L19
	} else {
		goto L88
	}
L9:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82+l0)))
	if v84 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v41 {
	case 0:
		goto L22
	case 1:
		goto L23
	case 2:
		goto L24
	case 3:
		goto L25
	default:
		v168 = l0
		goto L8
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(2)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = F_checkWellFormedRecursionWalker(m, v34, l1)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	switch v20 - int32(3) {
	case 0:
		v338 = l0
		v341 = v23
		v345 = v3
		goto L6
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
		v570 = l0
		v577 = v3
		goto L2
	case 19:
		goto L11
	default:
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v20 == int32(110) {
		v593 = v3
		v594 = v3
		goto L1
	} else {
		goto L17
	}
L15:
	;
	if v20 == int32(64) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v570 = l0
	v577 = v3
	goto L2
L17:
	;
	if v20 == int32(141) {
		v197 = l0
		v204 = v3
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v570 = l0
	v577 = v3
	goto L2
L19:
	;
	return int32(0)
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	v82 = int32(12)
	goto L9
L21:
	;
	v82 = int32(28)
	goto L9
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v76 = F_checkWellFormedRecursionWalker(m, v75, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L19
	} else {
		goto L41
	}
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = F_checkWellFormedRecursionWalker(m, v64, l1)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L19
	} else {
		goto L36
	}
L24:
	;
	if v23 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	if v23 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v47 = F_checkWellFormedRecursionWalker(m, v46, l1)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v51 = F_checkWellFormedRecursionWalker(m, v50, l1)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	goto L21
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v58 = F_checkWellFormedRecursionWalker(m, v57, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v61 = F_checkWellFormedRecursionWalker(m, v60, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	goto L21
L36:
	;
	if v23 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v72 = F_checkWellFormedRecursionWalker(m, v71, l1)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	goto L21
L41:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v79 = F_checkWellFormedRecursionWalker(m, v78, l1)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L19
	} else {
		goto L42
	}
L42:
	;
	goto L21
L43:
	;
	v593 = int32(1)
	v594 = v3
	goto L1
L44:
	;
	goto L45
L45:
	;
	v88 = v84
	goto L46
L46:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v102 = int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v103 <= int32(63) {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v593 = v102
	v594 = v3
	goto L1
L48:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v165+v88)))
	if v167 != 0 {
		v88 = v167
		goto L46
	} else {
		goto L87
	}
L49:
	;
	v165 = int32(28)
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(2)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	v159 = F_checkWellFormedRecursionWalker(m, v158, l1)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L19
	} else {
		goto L86
	}
L51:
	;
	v107 = v103 - int32(3)
	if v107 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	if v103 != int32(64) {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	if v107 == int32(19) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v338 = v88
	v341 = v101
	v345 = v102
	goto L6
L57:
	;
	goto L50
L58:
	;
	v570 = v88
	v577 = v102
	goto L2
L60:
	;
	if v103 == int32(110) {
		v593 = v102
		v594 = v3
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	switch v116 {
	case 0:
		goto L68
	case 1:
		goto L67
	case 2:
		goto L66
	case 3:
		goto L65
	default:
		v168 = v88
		goto L8
	}
L63:
	;
	if v103 == int32(141) {
		v197 = v88
		v204 = v102
		goto L7
	} else {
		goto L64
	}
L64:
	;
	v570 = v88
	v577 = v102
	goto L2
L65:
	;
	if v101 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L66:
	;
	if v101 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L67:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v124 = F_checkWellFormedRecursionWalker(m, v123, l1)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L71
	}
L68:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v118 = F_checkWellFormedRecursionWalker(m, v117, l1)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L19
	} else {
		goto L69
	}
L69:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v121 = F_checkWellFormedRecursionWalker(m, v120, l1)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L19
	} else {
		goto L70
	}
L70:
	;
	goto L49
L71:
	;
	if v101 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L74
L73:
	;
	goto L74
L74:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v131 = F_checkWellFormedRecursionWalker(m, v130, l1)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L19
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	goto L49
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L78
L77:
	;
	goto L78
L78:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v139 = F_checkWellFormedRecursionWalker(m, v138, l1)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L19
	} else {
		goto L79
	}
L79:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v142 = F_checkWellFormedRecursionWalker(m, v141, l1)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	goto L49
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L83
L82:
	;
	goto L83
L83:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v150 = F_checkWellFormedRecursionWalker(m, v149, l1)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L19
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v154 = F_checkWellFormedRecursionWalker(m, v153, l1)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L19
	} else {
		goto L85
	}
L85:
	;
	goto L49
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	v165 = int32(12)
	goto L48
L87:
	;
	goto L47
L88:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v185
	F_errmsg_internal(m, int32(_a_F_checkWellFormedRecursionWalker_0), v16+int32(32))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L19
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_checkWellFormedRecursionWalker_1), int32(1179), int32(_a_F_checkWellFormedRecursionWalker_2))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L19
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+8)))
	if v211 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	F_checkWellFormedSelectStmt(m, v197, l1)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L19
	} else {
		goto L121
	}
L94:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v216 = F_lcons(m, v214, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L19
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v271 = int32(0)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v274 = F_lcons(m, v271, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L19
	} else {
		goto L107
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v216
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v197)+64))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v220 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	F_checkWellFormedSelectStmt(m, v197, l1)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L19
	} else {
		goto L105
	}
L99:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v223 <= int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v230 = int32(0)
	goto L101
L101:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240+v230<<(uint(int32(2))%32))))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+16))
	v246 = F_checkWellFormedRecursionWalker(m, v245, l1)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L19
	} else {
		goto L103
	}
L102:
	;
	goto L98
L103:
	;
	v249 = v230 + int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	if v249 < v250 {
		v230 = v249
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v268 = F_list_delete_first(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L19
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v268
	v593 = v204
	v594 = v3
	goto L1
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v274
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v197)+64))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	if v278 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F_checkWellFormedSelectStmt(m, v197, l1)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L19
	} else {
		goto L119
	}
L109:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v281 <= int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v287 = v271
	goto L111
L111:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297+v287<<(uint(int32(2))%32))))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+16))
	v303 = F_checkWellFormedRecursionWalker(m, v302, l1)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L19
	} else {
		goto L113
	}
L112:
	;
	goto L108
L113:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v305 != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+12))
	v308 = v306
	goto L116
L115:
	;
	v308 = int32(0)
	goto L116
L116:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	v310 = F_lappend(m, v309, v301)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L19
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v308))) = v310
	v314 = v287 + int32(1)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v314 < v315 {
		v287 = v314
		goto L111
	} else {
		goto L118
	}
L118:
	;
	goto L112
L119:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v333 = F_list_delete_first(m, v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L19
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v333
	v538 = v204
	goto L4
L121:
	;
	v593 = v204
	v594 = v3
	goto L1
L122:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v352 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v471+v472*int32(12))))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470))))
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477))))
	if base.B2i32(v480 == int32(0))|base.B2i32(v480 != v483) != 0 {
		v501 = v480
		v502 = v483
		goto L150
	} else {
		goto L151
	}
L124:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	if v355 <= int32(0) {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v358 = int32(0)
	if v358 < v355 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v361 = v355
	goto L128
L127:
	;
	v361 = v358
	goto L128
L128:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v352)+12))
	v369 = v3
	goto L129
L129:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v362+v369<<(uint(int32(2))%32))))
	if v379 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	goto L123
L131:
	;
	v455 = v369 + int32(1)
	if v455 != v361 {
		v369 = v455
		goto L129
	} else {
		goto L148
	}
L132:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	if v382 <= int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v385 = int32(0)
	if v385 < v382 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v388 = v382
	goto L136
L135:
	;
	v388 = v385
	goto L136
L136:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	v394 = int32(0)
	goto L137
L137:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v390+v394<<(uint(int32(2))%32))))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	if base.B2i32(v412 == int32(0))|base.B2i32(v412 != v415) != 0 {
		v433 = v412
		v434 = v415
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L131
L139:
	;
	if v433-v434 == int32(0) {
		v538 = v345
		goto L4
	} else {
		goto L146
	}
L140:
	;
	goto L139
L141:
	;
	v418 = v389
	v419 = v409
	goto L142
L142:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+1)))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)))
	if v423 == int32(0) {
		v433 = v423
		v434 = v422
		goto L140
	} else {
		goto L144
	}
L143:
	;
	v433 = v423
	v434 = v422
	goto L140
L144:
	;
	v426 = int32(1)
	if v423 == v422 {
		v418 = v418 + v426
		v419 = v419 + v426
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v439 = v394 + int32(1)
	if v439 != v388 {
		v394 = v439
		goto L137
	} else {
		goto L147
	}
L147:
	;
	goto L138
L148:
	;
	goto L130
L149:
	;
	if v501-v502 != 0 {
		v593 = v345
		v594 = v3
		goto L1
	} else {
		goto L156
	}
L150:
	;
	goto L149
L151:
	;
	v486 = v470
	v487 = v477
	goto L152
L152:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+1)))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486)+1)))
	if v491 == int32(0) {
		v501 = v491
		v502 = v490
		goto L150
	} else {
		goto L154
	}
L153:
	;
	v501 = v491
	v502 = v490
	goto L150
L154:
	;
	v494 = int32(1)
	if v491 == v490 {
		v486 = v486 + v494
		v487 = v487 + v494
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	if v341 != 0 {
		goto L3
	} else {
		goto L157
	}
L157:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v506 = v504 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v506
	if v506 < int32(2) {
		v593 = v345
		v594 = v3
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L19
	} else {
		goto L159
	}
L159:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L19
	} else {
		goto L160
	}
L160:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v517
	F_errmsg(m, int32(_a_F_checkWellFormedRecursionWalker_3), v16)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L19
	} else {
		goto L161
	}
L161:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v338)+24))
	F_parser_errposition(m, v522, v523)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L19
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_checkWellFormedRecursionWalker_1), int32(1077), int32(_a_F_checkWellFormedRecursionWalker_2))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L19
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L19
	} else {
		goto L165
	}
L165:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v552
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v551<<(uint(int32(2))%32))+uint32(_c_F_checkWellFormedRecursionWalker[0])))
	F_errmsg(m, v556, v16+int32(16))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L19
	} else {
		goto L166
	}
L166:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v338)+24))
	F_parser_errposition(m, v561, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L19
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_checkWellFormedRecursionWalker_1), int32(1069), int32(_a_F_checkWellFormedRecursionWalker_2))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L19
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	v593 = v577
	v594 = v584
	goto L1
}
func F_checkWellFormedSelectStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
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
	var v26 int32
	_ = v26
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v8 != 0 {
		v10 = F_raw_expression_tree_walker_impl(m, l0, int32(520), l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		switch v12 {
		case 0, 1:
			v80 = F_raw_expression_tree_walker_impl(m, l0, int32(520), l1)
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		case 2:
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
			if v13 == int32(1) {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(4)
			} else {
			}
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v19 = F_checkWellFormedRecursionWalker(m, v18, l1)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				v22 = F_checkWellFormedRecursionWalker(m, v21, l1)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v27 = F_checkWellFormedRecursionWalker(m, v26, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v30 = F_checkWellFormedRecursionWalker(m, v29, l1)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v33 = F_checkWellFormedRecursionWalker(m, v32, l1)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
								v36 = F_checkWellFormedRecursionWalker(m, v35, l1)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									m.G0 = v6 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		case 3:
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
			if v38 == int32(1) {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(5)
			} else {
			}
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v44 = F_checkWellFormedRecursionWalker(m, v43, l1)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(5)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
				v49 = F_checkWellFormedRecursionWalker(m, v48, l1)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(0)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v54 = F_checkWellFormedRecursionWalker(m, v53, l1)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v57 = F_checkWellFormedRecursionWalker(m, v56, l1)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v60 = F_checkWellFormedRecursionWalker(m, v59, l1)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
								v63 = F_checkWellFormedRecursionWalker(m, v62, l1)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									m.G0 = v6 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return
			} else {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v69
				F_errmsg_internal(m, int32(_a_F_checkWellFormedSelectStmt_0), v6)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_checkWellFormedSelectStmt_1), int32(1267), int32(_a_F_checkWellFormedSelectStmt_2))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
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
func F_check_createrole_self_grant(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_pstrdup(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v195
L2:
	;
	return int32(0)
L3:
	;
	v21 = F_SplitIdentifierString(m, v14, int32(44), v11+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[1])) = v26
	goto L8
L6:
	;
	goto L7
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v40 == int32(0) {
		v158 = v4
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v32 = F_format_elog_string(m, int32(_a_F_check_createrole_self_grant_0), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[2])) = v32
	F_pfree(m, v14)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v195 = v4
	goto L1
L12:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[1])) = v174
	goto L53
L13:
	;
	F_pfree(m, v14)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L49
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v43 <= int32(0) {
		v158 = v4
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v47 = int32(0)
	v53 = v4
	goto L16
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v47<<(uint(int32(2))%32))))
	v63 = v59
	v64 = int32(_a_F_check_createrole_self_grant_1)
	goto L19
L17:
	;
	v158 = v147
	goto L13
L18:
	;
	if v101 != 0 {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v67 == v68 {
		v90 = v67
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v101 = int32(0)
	goto L18
L21:
	;
	v92 = int32(1)
	if v90 != 0 {
		v63 = v63 + v92
		v64 = v64 + v92
		goto L19
	} else {
		goto L30
	}
L22:
	;
	if base.Ui32((v67-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v78 = v67 | int32(32)
	goto L25
L24:
	;
	v78 = v67
	goto L25
L25:
	;
	if base.Ui32((v68-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v87 = v68 | int32(32)
	goto L28
L27:
	;
	v87 = v68
	goto L28
L28:
	;
	if v78 == v87 {
		v90 = v78
		goto L21
	} else {
		goto L29
	}
L29:
	;
	v101 = v78 - v87
	goto L18
L30:
	;
	goto L20
L31:
	;
	v105 = v59
	v106 = int32(_a_F_check_createrole_self_grant_2)
	goto L35
L32:
	;
	v146 = int32(4)
	goto L33
L33:
	;
	v147 = v146 | v53
	v149 = v47 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v149 < v150 {
		v47 = v149
		v53 = v147
		goto L16
	} else {
		goto L48
	}
L34:
	;
	if v143 != 0 {
		goto L12
	} else {
		goto L47
	}
L35:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v109 == v110 {
		v132 = v109
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v143 = int32(0)
	goto L34
L37:
	;
	v134 = int32(1)
	if v132 != 0 {
		v105 = v105 + v134
		v106 = v106 + v134
		goto L35
	} else {
		goto L46
	}
L38:
	;
	if base.Ui32((v109-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v120 = v109 | int32(32)
	goto L41
L40:
	;
	v120 = v109
	goto L41
L41:
	;
	if base.Ui32((v110-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v129 = v110 | int32(32)
	goto L44
L43:
	;
	v129 = v110
	goto L44
L44:
	;
	if v120 == v129 {
		v132 = v120
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v143 = v120 - v129
	goto L34
L46:
	;
	goto L36
L47:
	;
	v146 = int32(2)
	goto L33
L48:
	;
	goto L17
L49:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v166 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	if v166 == int32(0) {
		v195 = v4
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v158
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v166
	v195 = int32(1)
	goto L1
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v59
	v180 = F_format_elog_string(m, int32(_a_F_check_createrole_self_grant_3), v11)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[2])) = v180
	F_pfree(m, v14)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	F_list_free(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v195 = v4
	goto L1
}
func F_check_debug_io_direct(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_pstrdup(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v244
L2:
	;
	return int32(0)
L3:
	;
	v20 = F_SplitGUCList(m, v14, v11+int32(28))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v20 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[1])) = v25
	goto L8
L6:
	;
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v42 == int32(0) {
		v207 = v4
		goto L13
	} else {
		goto L14
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_check_debug_io_direct_0)
	v34 = F_format_elog_string(m, int32(_a_F_check_debug_io_direct_1), v11+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[2])) = v34
	F_pfree(m, v14)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_list_free(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v244 = v4
	goto L1
L12:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[1])) = v223
	goto L66
L13:
	;
	F_pfree(m, v14)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L62
	}
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v45 <= int32(0) {
		v207 = v4
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v49 = int32(0)
	v55 = v4
	goto L16
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v49<<(uint(int32(2))%32))))
	v66 = v62
	v67 = int32(_a_F_check_debug_io_direct_2)
	goto L20
L17:
	;
	v207 = v196
	goto L13
L18:
	;
	v196 = v195 | v55
	v198 = v49 + int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v198 < v199 {
		v49 = v198
		v55 = v196
		goto L16
	} else {
		goto L61
	}
L19:
	;
	if v104 == int32(0) {
		v195 = int32(1)
		goto L18
	} else {
		goto L32
	}
L20:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v70 == v71 {
		v93 = v70
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v104 = int32(0)
	goto L19
L22:
	;
	v95 = int32(1)
	if v93 != 0 {
		v66 = v66 + v95
		v67 = v67 + v95
		goto L20
	} else {
		goto L31
	}
L23:
	;
	if base.Ui32((v70-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v81 = v70 | int32(32)
	goto L26
L25:
	;
	v81 = v70
	goto L26
L26:
	;
	if base.Ui32((v71-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v90 = v71 | int32(32)
	goto L29
L28:
	;
	v90 = v71
	goto L29
L29:
	;
	if v81 == v90 {
		v93 = v81
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v104 = v81 - v90
	goto L19
L31:
	;
	goto L21
L32:
	;
	v111 = v62
	v112 = int32(_a_F_check_debug_io_direct_3)
	goto L34
L33:
	;
	if v149 == int32(0) {
		v195 = int32(2)
		goto L18
	} else {
		goto L46
	}
L34:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v115 == v116 {
		v138 = v115
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v149 = int32(0)
	goto L33
L36:
	;
	v140 = int32(1)
	if v138 != 0 {
		v111 = v111 + v140
		v112 = v112 + v140
		goto L34
	} else {
		goto L45
	}
L37:
	;
	if base.Ui32((v115-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v126 = v115 | int32(32)
	goto L40
L39:
	;
	v126 = v115
	goto L40
L40:
	;
	if base.Ui32((v116-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v135 = v116 | int32(32)
	goto L43
L42:
	;
	v135 = v116
	goto L43
L43:
	;
	if v126 == v135 {
		v138 = v126
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v149 = v126 - v135
	goto L33
L45:
	;
	goto L35
L46:
	;
	v155 = v62
	v156 = int32(_a_F_check_debug_io_direct_4)
	goto L48
L47:
	;
	if v193 != 0 {
		goto L12
	} else {
		goto L60
	}
L48:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v159 == v160 {
		v182 = v159
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v193 = int32(0)
	goto L47
L50:
	;
	v184 = int32(1)
	if v182 != 0 {
		v155 = v155 + v184
		v156 = v156 + v184
		goto L48
	} else {
		goto L59
	}
L51:
	;
	if base.Ui32((v159-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v170 = v159 | int32(32)
	goto L54
L53:
	;
	v170 = v159
	goto L54
L54:
	;
	if base.Ui32((v160-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v179 = v160 | int32(32)
	goto L57
L56:
	;
	v179 = v160
	goto L57
L57:
	;
	if v170 == v179 {
		v182 = v170
		goto L50
	} else {
		goto L58
	}
L58:
	;
	v193 = v170 - v179
	goto L47
L59:
	;
	goto L49
L60:
	;
	v195 = int32(4)
	goto L18
L61:
	;
	goto L17
L62:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_list_free(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v215 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v215
	if v215 == int32(0) {
		v244 = v4
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v207
	v244 = int32(1)
	goto L1
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v62
	v229 = F_format_elog_string(m, int32(_a_F_check_debug_io_direct_5), v11)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[2])) = v229
	F_pfree(m, v14)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_list_free(m, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v244 = v4
	goto L1
}
func F_check_exclusion_constraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v9 = int32(0)
	v12 = F_check_exclusion_or_unique_constraint(m, l0, l1, l2, l3, l4, l5, l6, l7, v9, v9, v9)
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		return
	}
}
func F_check_indirection(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(0)
	if v13 < v10 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v10
	goto L6
L5:
	;
	v16 = v13
	goto L6
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = v3
	goto L7
L7:
	;
	v28 = v23 << (uint(int32(2)) % 32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v19+v28)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if base.B2i32(v31 == int32(77))&base.B2i32(v28+int32(4) < v10<<(uint(int32(2))%32)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_scanner_yyerror(m, int32(_a_F_check_indirection_0), l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v41 = v23 + int32(1)
	if v16 != v41 {
		v23 = v41
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	goto L1
L13:
	;
	return int32(0)
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_lateral_ref_ok(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+22)))
	if v11 != int32(1) {
		m.G0 = v9 + int32(32)
		return
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+23)))
		if v14 != 0 {
			m.G0 = v9 + int32(32)
			return
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_check_lateral_ref_ok_0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v17
					F_errmsg(m, int32(_a_F_check_lateral_ref_ok_1), v9+int32(16))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						if v31 == int32(0) {
							v42 = F_errdetail(m, int32(_a_F_check_lateral_ref_ok_2), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								F_parser_errposition(m, l0, l2)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_check_lateral_ref_ok_3), int32(505), int32(_a_F_check_lateral_ref_ok_4))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
							if v15 != v34 {
								v42 = F_errdetail(m, int32(_a_F_check_lateral_ref_ok_2), int32(0))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return
								} else {
									F_parser_errposition(m, l0, l2)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_check_lateral_ref_ok_3), int32(505), int32(_a_F_check_lateral_ref_ok_4))
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
								F_errhint(m, int32(_a_F_check_lateral_ref_ok_5), v9)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									F_parser_errposition(m, l0, l2)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_check_lateral_ref_ok_3), int32(505), int32(_a_F_check_lateral_ref_ok_4))
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
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
				}
			}
		}
	}
}
func F_check_stack_depth(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_check_stack_depth[0]))
	if v8 == int32(0) {
		m.G0 = v5 + int32(16)
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_check_stack_depth[1]))
		v13 = v8 - v5
		v15 = v13 >> (uint(int32(31)) % 32)
		if v13^v15-v15 <= v12 {
			m.G0 = v5 + int32(16)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_errcode(m, int32(16777477))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_check_stack_depth_0), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_check_stack_depth[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v31
						F_errhint(m, int32(_a_F_check_stack_depth_1), v5)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_stack_depth_2), int32(105), int32(_a_F_check_stack_depth_3))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
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
	}
}
func F_check_synchronized_standby_slots(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v14 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v301
L2:
	;
	v301 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v18 = F_pstrdup(m, v13)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_pfree(m, v18)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L6
	} else {
		goto L72
	}
L6:
	;
	return int32(0)
L7:
	;
	v25 = F_SplitIdentifierString(m, v18, int32(44), v11+int32(32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[1])) = v30
	goto L12
L10:
	;
	goto L11
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v39 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v36 = F_format_elog_string(m, int32(_a_F_check_synchronized_standby_slots_0), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[2])) = v36
	v288 = v4
	goto L5
L14:
	;
	v288 = int32(1)
	goto L5
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if int32(0) < v42 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v46 = int32(0)
	goto L19
L17:
	;
	v113 = v39
	goto L18
L18:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v118 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54+v46<<(uint(int32(2))%32))))
	v59 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v59
	v70 = F_ReplicationSlotValidateNameInternal(m, v58, v59, v11+int32(44), v11+int32(40), v11+int32(36))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v107 == int32(0) {
		goto L14
	} else {
		goto L32
	}
L21:
	;
	if v70 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[3])) = v74
	goto L25
L23:
	;
	goto L24
L24:
	;
	v104 = v46 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v104 < v105 {
		v46 = v104
		goto L19
	} else {
		goto L31
	}
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[1])) = v78
	goto L26
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v81
	v87 = F_format_elog_string(m, int32(_a_F_check_synchronized_standby_slots_1), v11+int32(16))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[2])) = v87
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	if v90 == int32(0) {
		v288 = v4
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[1])) = v94
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v90
	v100 = F_format_elog_string(m, int32(_a_F_check_synchronized_standby_slots_1), v11)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[4])) = v100
	v288 = v4
	goto L5
L31:
	;
	goto L20
L32:
	;
	v113 = v107
	goto L18
L33:
	;
	v152 = F_guc_malloc(m, v147)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L40
	}
L34:
	;
	v147 = int32(4)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	v125 = int32(0)
	v128 = int32(4)
	goto L37
L37:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v122+v125<<(uint(int32(2))%32))))
	v137 = F_strlen(m, v136)
	mBase = m.M
	v139 = int32(1)
	v140 = v137 + v128 + v139
	v142 = v125 + v139
	if v142 != v118 {
		v125 = v142
		v128 = v140
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v147 = v140
	goto L33
L39:
	;
	goto L38
L40:
	;
	if v152 == int32(0) {
		v301 = v4
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v156 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v159 = v157
	goto L44
L43:
	;
	v159 = int32(0)
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v161 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v152
	goto L14
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v164 <= int32(0) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v170 = int32(0)
	v174 = v152 + int32(4)
	goto L48
L48:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178+v170<<(uint(int32(2))%32))))
	if (v182^v174)&int32(3) != 0 {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	goto L45
L50:
	;
	v257 = F_strlen(m, v182)
	mBase = m.M
	v259 = int32(1)
	v262 = v170 + v259
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v262 < v263 {
		v170 = v262
		v174 = v174 + v257 + v259
		goto L48
	} else {
		goto L71
	}
L51:
	;
	goto L50
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v237))) = uint8(v236)
	if v236&int32(255) == int32(0) {
		goto L51
	} else {
		goto L67
	}
L53:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	v235 = v182
	v236 = v188
	v237 = v174
	goto L52
L54:
	;
	goto L55
L55:
	;
	if v182&int32(3) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v192 = v182
	v194 = v174
	goto L59
L57:
	;
	v206 = v182
	v208 = v174
	goto L58
L58:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v213 = int32(-2139062144)
	if (int32(16843008)-v210|v210)&v213 != v213 {
		v235 = v206
		v236 = v210
		v237 = v208
		goto L52
	} else {
		goto L63
	}
L59:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v195)
	if v195 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L60:
	;
	v206 = v202
	v208 = v200
	goto L58
L61:
	;
	v199 = int32(1)
	v200 = v194 + v199
	v202 = v192 + v199
	if v202&int32(3) != 0 {
		v192 = v202
		v194 = v200
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v218 = v206
	v219 = v210
	v220 = v208
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220))) = v219
	v222 = int32(4)
	v223 = v220 + v222
	v225 = v218 + v222
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	v230 = int32(-2139062144)
	if (int32(16843008)-v227|v227)&v230 == v230 {
		v218 = v225
		v219 = v227
		v220 = v223
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v235 = v225
	v236 = v227
	v237 = v223
	goto L52
L66:
	;
	goto L65
L67:
	;
	v244 = v235
	v246 = v237
	goto L68
L68:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)) = uint8(v247)
	v249 = int32(1)
	if v247 != 0 {
		v244 = v244 + v249
		v246 = v246 + v249
		goto L68
	} else {
		goto L70
	}
L69:
	;
	goto L51
L70:
	;
	goto L69
L71:
	;
	goto L49
L72:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	F_list_free(m, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	v301 = v288
	goto L1
}
func F_check_synchronous_standby_names(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
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
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
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
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v645 int32
	_ = v645
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v677 int32
	_ = v677
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v718 int32
	_ = v718
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v996 int32
	_ = v996
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1028 int32
	_ = v1028
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1099 int32
	_ = v1099
	var v1105 int32
	_ = v1105
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1152 int32
	_ = v1152
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1326 int32
	_ = v1326
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1364 int32
	_ = v1364
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1427 int32
	_ = v1427
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1503 int32
	_ = v1503
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1589 int32
	_ = v1589
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1620 int32
	_ = v1620
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1651 int32
	_ = v1651
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1668 int32
	_ = v1668
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1699 int32
	_ = v1699
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1792 int32
	_ = v1792
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1806 int32
	_ = v1806
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1930 int32
	_ = v1930
	var v1937 int32
	_ = v1937
	var v1944 int32
	_ = v1944
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1962 int32
	_ = v1962
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1975 int32
	_ = v1975
	var v1976 int32
	_ = v1976
	var v1981 int32
	_ = v1981
	var v1987 int32
	_ = v1987
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v2013 int32
	_ = v2013
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2035 int32
	_ = v2035
	var v2045 int32
	_ = v2045
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2089 int32
	_ = v2089
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2135 int32
	_ = v2135
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2154 int32
	_ = v2154
	var v2160 int32
	_ = v2160
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2186 int32
	_ = v2186
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2208 int32
	_ = v2208
	var v2218 int32
	_ = v2218
	var v2238 int32
	_ = v2238
	var v2242 int32
	_ = v2242
	var v2244 int32
	_ = v2244
	var v2248 int32
	_ = v2248
	var v2254 int32
	_ = v2254
	var v2260 int32
	_ = v2260
	var v2284 int32
	_ = v2284
	var v2288 int32
	_ = v2288
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2303 int32
	_ = v2303
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2319 int32
	_ = v2319
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2335 int32
	_ = v2335
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2353 int32
	_ = v2353
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2381 int32
	_ = v2381
	var v2387 int32
	_ = v2387
	var v2391 int32
	_ = v2391
	var v2399 int32
	_ = v2399
	var v2401 int32
	_ = v2401
	var v2403 int32
	_ = v2403
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2433 int32
	_ = v2433
	var v2436 int32
	_ = v2436
	var v2441 int32
	_ = v2441
	var v2444 int32
	_ = v2444
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2450 int32
	_ = v2450
	var v2453 int32
	_ = v2453
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2459 int32
	_ = v2459
	var v2462 int32
	_ = v2462
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2468 int32
	_ = v2468
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2481 int32
	_ = v2481
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2490 int32
	_ = v2490
	var v2493 int32
	_ = v2493
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2501 int32
	_ = v2501
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2508 int32
	_ = v2508
	var v2514 int32
	_ = v2514
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2530 int32
	_ = v2530
	var v2532 int32
	_ = v2532
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2536 int32
	_ = v2536
	var v2538 int32
	_ = v2538
	var v2540 int32
	_ = v2540
	var v2544 int32
	_ = v2544
	var v2550 int32
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2554 int32
	_ = v2554
	var v2556 int32
	_ = v2556
	var v2563 int32
	_ = v2563
	var v2579 int32
	_ = v2579
	var v2583 int32
	_ = v2583
	var v2585 int32
	_ = v2585
	var v2588 int32
	_ = v2588
	var v2593 int32
	_ = v2593
	var v2595 int32
	_ = v2595
	var v2599 int32
	_ = v2599
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2611 int32
	_ = v2611
	var v2614 int32
	_ = v2614
	var v2618 int32
	_ = v2618
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2646 int32
	_ = v2646
	var v2673 int32
	_ = v2673
	var v2697 int32
	_ = v2697
	v4 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(48)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v31 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v2673 + int32(48)
	return v2697
L2:
	;
	v2673 = v2646
	v2697 = int32(1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v2646 = v29
	goto L2
L4:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v34 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v37
	v42 = v29 + int32(44)
	v44 = F_palloc0(m, int32(16))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v71 = int32(0)
	base.MemoryFill(m, v50, v71, int32(96))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+60)) = v71
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v74)+52)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v74)+44)) = v71
	*(*int64)(unsafe.Add(mBase, uint32(v74)+36)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v74)+4)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v74)+12)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = v71
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v44
	v91 = F_strlen(m, v31)
	mBase = m.M
	v93 = v91 + int32(2)
	v94 = F_palloc(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L19
	}
L9:
	;
	v50 = F_palloc(m, int32(96))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v56 = int32(28)
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = v56
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v50
	if v50 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v56 = int32(48)
	goto L11
L14:
	;
	F_errmsg_internal(m, int32(_a_F_check_synchronous_standby_names_0), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_check_synchronous_standby_names_1), int32(179), int32(_a_F_check_synchronous_standby_names_2))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v354 = m.G0
	v356 = v354 - int32(1024)
	m.G0 = v356
	*(*int32)(unsafe.Add(mBase, uint32(v356)+1020)) = int32(0)
	v362 = v356 + int32(816)
	v364 = v356 + int32(16)
	v366 = v362
	v367 = l1
	v368 = v29
	v369 = v353
	v370 = int32(-2)
	v378 = v364
	v380 = v356
	v381 = v362
	v382 = int32(200)
	v384 = v29 + int32(36)
	v385 = v4
	v386 = v364
	v388 = v29 + int32(40)
	goto L53
L18:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_3))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L6
	} else {
		goto L51
	}
L19:
	;
	if v94 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v91 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_4))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L6
	} else {
		goto L50
	}
L23:
	;
	v247 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v94+v91))) = uint16(v247)
	if base.Ui32(v93) < base.Ui32(int32(2)) {
		v334 = v247
		goto L37
	} else {
		goto L38
	}
L24:
	;
	v99 = v91 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v91) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v111 = v4
	v114 = v4
	goto L28
L26:
	;
	v166 = v4
	goto L27
L27:
	;
	v192 = v166
	v196 = v4
	goto L32
L28:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v111))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94+v111))) = uint8(v132)
	v135 = v111 | int32(1)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v31))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94+v135))) = uint8(v138)
	v141 = v111 | int32(2)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v31))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94+v141))) = uint8(v144)
	v147 = v111 | int32(3)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v31))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94+v147))) = uint8(v150)
	v152 = int32(4)
	v153 = v111 + v152
	v155 = v114 + v152
	if v155 != v91&int32(2147483644) {
		v111 = v153
		v114 = v155
		goto L28
	} else {
		goto L30
	}
L29:
	;
	if v99 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v166 = v153
	goto L27
L32:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v192))))
	*(*uint8)(unsafe.Add(mBase, uint32(v94+v192))) = uint8(v213)
	v215 = int32(1)
	v218 = v196 + v215
	if v218 != v99 {
		v192 = v192 + v215
		v196 = v218
		goto L32
	} else {
		goto L34
	}
L33:
	;
	goto L23
L34:
	;
	goto L33
L35:
	;
	if v334 == int32(0) {
		goto L18
	} else {
		goto L49
	}
L36:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_5))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L6
	} else {
		goto L48
	}
L37:
	;
	goto L35
L38:
	;
	v253 = v93 - int32(2)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v253))))
	if v255 != 0 {
		v334 = v247
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v93-int32(1)))))
	if v259 != 0 {
		v334 = v247
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v261 = F_palloc(m, int32(48))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	if v261 == int32(0) {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v265 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v261)+20)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v261)+8)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v261)+4)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v261)+12)) = v253
	*(*int64)(unsafe.Add(mBase, uint32(v261)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v261)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v261)+16)) = v253
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v265
	F_syncrep_yyensure_buffer_stack(m, v89)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v279+v280<<(uint(int32(2))%32))))
	if v284 == v261 {
		v334 = v261
		goto L37
	} else {
		goto L44
	}
L44:
	;
	if v284 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v286))) = uint8(v287)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v291 = int32(2)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v289+v290<<(uint(v291)%32))))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v89)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v294)+8)) = v295
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v297+v298<<(uint(v291)%32))))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v302)+16)) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v307 = v305
	v308 = v306
	goto L47
L46:
	;
	v307 = v279
	v308 = v280
	goto L47
L47:
	;
	v309 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v308<<(uint(v309)%32)+v307))) = v261
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v317 = v313 + v314<<(uint(v309)%32)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v319
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+36)) = v322
	*(*int32)(unsafe.Add(mBase, uint32(v89)+80)) = v322
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v326
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+24)) = uint8(v328)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+48)) = int32(1)
	v334 = v261
	goto L37
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v334)+20)) = int32(1)
	goto L17
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	if v2563+int32(816) != v2556 {
		goto L398
	} else {
		goto L399
	}
L53:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v381))) = uint8(v385)
	if base.Ui32(v366+v382-int32(1)) <= base.Ui32(v381) {
		goto L59
	} else {
		goto L60
	}
L54:
	;
	v2550 = v367
	v2551 = v368
	v2554 = int32(1)
	v2556 = v411
	v2563 = v380
	goto L52
L55:
	;
	goto L54
L56:
	;
	v366 = v2518
	v367 = v2519
	v368 = v2520
	v369 = v2521
	v370 = v2522
	v378 = v2530
	v380 = v2532
	v381 = v2533 + int32(1)
	v382 = v2534
	v384 = v2536
	v385 = base.I32_extend8_s(v2544)
	v386 = v2538
	v388 = v2540
	goto L53
L57:
	;
	v2428 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2360)+uint32(_c_F_check_synchronous_standby_names[1]))))
	v2431 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2428)+uint32(_c_F_check_synchronous_standby_names[2]))))
	v2433 = int32(2)
	v2436 = *(*int32)(unsafe.Add(mBase, uint32(v2353+(int32(1)-v2431)<<(uint(v2433)%32))))
	switch v2428&int32(255) - v2433 {
	case 0:
		goto L386
	case 1:
		goto L385
	case 2:
		goto L384
	case 3:
		goto L383
	case 4:
		goto L382
	case 5:
		goto L381
	case 6:
		goto L380
	case 7, 8:
		goto L379
	default:
		v2490 = v2436
		goto L378
	}
L58:
	;
	v2399 = m.G0
	v2401 = v2399 - int32(32)
	m.G0 = v2401
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2391)))
	if v2403 == int32(0) {
		goto L369
	} else {
		goto L370
	}
L59:
	;
	v397 = int32(2)
	v398 = int32(_a_F_check_synchronous_standby_names_6)
	if int32(_a_F_check_synchronous_standby_names_7) < v382 {
		v2373 = v366
		v2374 = v367
		v2375 = v368
		v2376 = v369
		v2378 = v397
		v2381 = v398
		v2387 = v380
		v2391 = v384
		goto L58
	} else {
		goto L62
	}
L60:
	;
	v441 = v366
	v446 = v378
	v447 = v381
	v448 = v382
	v449 = v386
	goto L61
L61:
	;
	if v385 == int32(12) {
		goto L79
	} else {
		goto L80
	}
L62:
	;
	v401 = int32(_a_F_check_synchronous_standby_names_8)
	v403 = v382 << (uint(int32(1)) % 32)
	if v401 <= v403 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v406 = v401
	goto L65
L64:
	;
	v406 = v403
	goto L65
L65:
	;
	v411 = F_palloc(m, v406*int32(5)+int32(3))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	if v411 == int32(0) {
		v2373 = v366
		v2374 = v367
		v2375 = v368
		v2376 = v369
		v2378 = v397
		v2381 = v398
		v2387 = v380
		v2391 = v384
		goto L58
	} else {
		goto L67
	}
L67:
	;
	v415 = v381 - v366
	v417 = v415 + int32(1)
	if v417 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	base.MemoryCopy(m, v411, v366, v417)
	goto L70
L69:
	;
	goto L70
L70:
	;
	v422 = base.I32_div_s(v406+int32(3), int32(4))
	v423 = int32(2)
	v425 = v411 + v422<<(uint(v423)%32)
	v427 = v417 << (uint(v423) % 32)
	if v427 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	base.MemoryCopy(m, v425, v386, v427)
	goto L73
L72:
	;
	goto L73
L73:
	;
	if v380+int32(816) != v366 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_pfree(m, v366)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L6
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if v406-int32(1) <= v415 {
		goto L55
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	v441 = v411
	v446 = v425 + v427 - int32(4)
	v447 = v411 + v415
	v448 = v406
	v449 = v425
	goto L61
L79:
	;
	v2550 = v367
	v2551 = v368
	v2554 = int32(0)
	v2556 = v441
	v2563 = v380
	goto L52
L80:
	;
	goto L81
L81:
	;
	v454 = int32(1) << (uint(v385) % 32)
	if v454&int32(13390146) != 0 {
		v2341 = v441
		v2342 = v367
		v2343 = v368
		v2344 = v369
		v2345 = v370
		v2353 = v446
		v2355 = v380
		v2356 = v447
		v2357 = v448
		v2359 = v384
		v2360 = v385
		v2361 = v449
		v2363 = v388
		v2364 = v454
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if v2364&int32(_a_F_check_synchronous_standby_names_9) == int32(0) {
		goto L57
	} else {
		goto L368
	}
L83:
	;
	v459 = int32(*(*int8)(unsafe.Add(mBase, uint32(v385)+uint32(_c_F_check_synchronous_standby_names[3]))))
	if v370 == int32(-2) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v2330 = v459 + v2329
	if base.Ui32(int32(22)) < base.Ui32(v2330) {
		v2341 = v2291
		v2342 = v2292
		v2343 = v2293
		v2344 = v2294
		v2345 = v2328
		v2353 = v2303
		v2355 = v2305
		v2356 = v2306
		v2357 = v2307
		v2359 = v2309
		v2360 = v2310
		v2361 = v2311
		v2363 = v2313
		v2364 = v2314
		goto L82
	} else {
		goto L366
	}
L85:
	;
	v462 = m.G0
	v464 = v462 - int32(32)
	m.G0 = v464
	*(*int32)(unsafe.Add(mBase, uint32(v369)+92)) = v380 + int32(1020)
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v369)+40))
	if v469 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v2291 = v441
	v2292 = v367
	v2293 = v368
	v2294 = v369
	v2295 = v370
	v2303 = v446
	v2305 = v380
	v2306 = v447
	v2307 = v448
	v2309 = v384
	v2310 = v385
	v2311 = v449
	v2313 = v388
	v2314 = v454
	goto L87
L87:
	;
	if v2295 <= int32(0) {
		goto L359
	} else {
		goto L360
	}
L88:
	;
	v2291 = v569
	v2292 = v570
	v2293 = v571
	v2294 = v572
	v2295 = v1651
	v2303 = v581
	v2305 = v583
	v2306 = v584
	v2307 = v585
	v2309 = v587
	v2310 = v588
	v2311 = v589
	v2313 = v591
	v2314 = v592
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+40)) = int32(1)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v369)+44))
	if v474 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v539 = v441
	v540 = v367
	v541 = v368
	v542 = v369
	v551 = v446
	v553 = v380
	v554 = v447
	v555 = v448
	v556 = v464
	v557 = v384
	v558 = v385
	v559 = v449
	v561 = v388
	v562 = v454
	goto L108
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v369)+44)) = int32(1)
	goto L94
L93:
	;
	goto L94
L94:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v479 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v483 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v483
	goto L97
L96:
	;
	goto L97
L97:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v369)+8))
	if v485 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+8)) = v489
	goto L100
L99:
	;
	goto L100
L100:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v369)+20))
	if v491 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+28)) = v520
	v524 = v517 + v516<<(uint(int32(2))%32)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+80)) = v526
	*(*int32)(unsafe.Add(mBase, uint32(v369)+36)) = v526
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v529)))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v530
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	*(*uint8)(unsafe.Add(mBase, uint32(v369)+24)) = uint8(v532)
	goto L91
L102:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v491+v492<<(uint(int32(2))%32))))
	if v496 != 0 {
		v516 = v492
		v517 = v491
		v519 = v496
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	F_syncrep_yyensure_buffer_stack(m, v369)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L6
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	v502 = F_syncrep_yy_create_buffer(m, v501, v369)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v369)+20))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v506 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v504+v505<<(uint(v506)%32)))) = v502
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v369)+20))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v510+v511<<(uint(v506)%32))))
	v516 = v511
	v517 = v510
	v519 = v515
	goto L101
L108:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v542)+36))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v565))) = uint8(v566)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v542)+44))
	v569 = v539
	v570 = v540
	v571 = v541
	v572 = v542
	v573 = v568
	v574 = v565
	v577 = v565
	v581 = v551
	v583 = v553
	v584 = v554
	v585 = v555
	v586 = v556
	v587 = v557
	v588 = v558
	v589 = v559
	v591 = v561
	v592 = v562
	goto L110
L110:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577))))
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595)+uint32(_c_F_check_synchronous_standby_names[6]))))
	if base.Ui32(int32(-26)) <= base.Ui32(v573-int32(31)) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+68)) = v577
	*(*int32)(unsafe.Add(mBase, uint32(v572)+64)) = v573
	goto L114
L113:
	;
	goto L114
L114:
	;
	v603 = int32(1)
	v607 = int32(*(*int16)(unsafe.Add(mBase, uint32(v573<<(uint(v603)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v608 = v607 + v596
	v613 = int32(*(*int16)(unsafe.Add(mBase, uint32(v608<<(uint(v603)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v613 != v573 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v619 = v573
	v624 = v596
	v625 = v596
	goto L118
L116:
	;
	v677 = v608
	goto L117
L117:
	;
	v697 = int32(1)
	v703 = int32(*(*int16)(unsafe.Add(mBase, uint32(v677<<(uint(v697)%32))+uint32(_c_F_check_synchronous_standby_names[9]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v677))%64)&int64(-32985348833280) == int64(0) {
		v573 = v703
		v577 = v577 + v697
		goto L110
	} else {
		goto L124
	}
L118:
	;
	v645 = int32(*(*int16)(unsafe.Add(mBase, uint32(v619<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[10]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v619))%64)&int64(2076672024) != int64(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v677 = v662
	goto L117
L120:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625)+uint32(_c_F_check_synchronous_standby_names[11]))))
	v654 = v653
	goto L122
L121:
	;
	v654 = v624
	goto L122
L122:
	;
	v656 = v654 & int32(255)
	v657 = int32(1)
	v661 = int32(*(*int16)(unsafe.Add(mBase, uint32(v645<<(uint(v657)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v662 = v656 + v661
	v667 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v662<<(uint(v657)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v667 != v645&int32(_a_F_check_synchronous_standby_names_10) {
		v619 = v645
		v624 = v654
		v625 = v656
		goto L118
	} else {
		goto L123
	}
L123:
	;
	goto L119
L124:
	;
	v718 = v574
	goto L125
L125:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v572)+64))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v572)+68))
	v743 = v737
	v746 = v718
	v750 = v738
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+80)) = v746
	*(*int32)(unsafe.Add(mBase, uint32(v572)+32)) = v750 - v746
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750))))
	*(*uint8)(unsafe.Add(mBase, uint32(v572)+24)) = uint8(v768)
	v770 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v750))) = uint8(v770)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v750
	v777 = int32(*(*int16)(unsafe.Add(mBase, uint32(v743<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[12]))))
	v784 = v777
	goto L129
L129:
	;
	v804 = int32(260)
	switch v784 {
	case 0:
		goto L158
	case 1:
		v539 = v569
		v540 = v570
		v541 = v571
		v542 = v572
		v551 = v581
		v553 = v583
		v554 = v584
		v555 = v585
		v556 = v586
		v557 = v587
		v558 = v588
		v559 = v589
		v561 = v591
		v562 = v592
		goto L108
	case 2:
		goto L145
	case 3:
		goto L144
	case 4:
		goto L157
	case 5:
		goto L156
	case 6:
		goto L155
	case 7:
		goto L154
	case 8:
		goto L152
	case 9:
		goto L151
	case 10:
		goto L150
	case 11:
		goto L143
	case 12:
		goto L142
	case 13:
		goto L141
	case 14:
		v1651 = v804
		goto L140
	case 15:
		goto L149
	case 16:
		goto L147
	case 17:
		goto L148
	case 18:
		goto L153
	default:
		goto L146
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v2260
	*(*int32)(unsafe.Add(mBase, uint32(v572)+48)) = int32(0)
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(v572)+44))
	v2288 = base.I32_div_s(v2284-int32(1), int32(2))
	v784 = v2288 + int32(17)
	goto L129
L132:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_11))
	mBase = m.M
	v2254 = m.ExcPending
	if v2254 != 0 {
		goto L6
	} else {
		goto L358
	}
L133:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_12))
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L6
	} else {
		goto L357
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	v2102 = v2082 + v2089
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v2102
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v572)+44))
	if base.Ui32(v2102) <= base.Ui32(v2081) {
		v743 = v2104
		v746 = v2081
		v750 = v2102
		goto L127
	} else {
		goto L338
	}
L136:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1719)))
	*(*int32)(unsafe.Add(mBase, uint32(v1720)+16)) = v1699
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	if v1723 != 0 {
		v1846 = int32(0)
		goto L282
	} else {
		goto L283
	}
L137:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1699 = v1668
	v1719 = v1688 + v1689<<(uint(int32(2))%32)
	goto L136
L138:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_13))
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L6
	} else {
		goto L281
	}
L139:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_14))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L6
	} else {
		goto L280
	}
L140:
	;
	m.G0 = v586 + int32(32)
	goto L88
L141:
	;
	v1651 = int32(41)
	goto L140
L142:
	;
	v1651 = int32(40)
	goto L140
L143:
	;
	v1651 = int32(44)
	goto L140
L144:
	;
	v1651 = int32(262)
	goto L140
L145:
	;
	v1651 = int32(261)
	goto L140
L146:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_15))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L6
	} else {
		goto L279
	}
L147:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v750))) = uint8(v869)
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v875 = v871 + v872<<(uint(int32(2))%32)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v876)+44))
	if v877 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L148:
	;
	v1651 = int32(0)
	goto L140
L149:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_16))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L6
	} else {
		goto L170
	}
L150:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v572)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v860))) = int32(_a_F_check_synchronous_standby_names_17)
	v1651 = int32(258)
	goto L140
L151:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	v855 = F_pstrdup(m, v854)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L6
	} else {
		goto L169
	}
L152:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	v849 = F_pstrdup(m, v848)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L6
	} else {
		goto L168
	}
L153:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	if v830 != 0 {
		v1651 = v804
		goto L140
	} else {
		goto L162
	}
L154:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v572)+92))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	*(*int32)(unsafe.Add(mBase, uint32(v820))) = v822
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+44)) = int32(1)
	v1651 = int32(258)
	goto L140
L155:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	F_appendStringInfoString(m, v816, v817)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L6
	} else {
		goto L161
	}
L156:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	F_appendStringInfoChar(m, v812, int32(34))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L6
	} else {
		goto L160
	}
L157:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v572)))
	F_initStringInfo(m, v807)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L6
	} else {
		goto L159
	}
L158:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v750))) = uint8(v805)
	v718 = v746
	goto L125
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+44)) = int32(3)
	v539 = v569
	v540 = v570
	v541 = v571
	v542 = v572
	v551 = v581
	v553 = v583
	v554 = v584
	v555 = v585
	v556 = v586
	v557 = v587
	v558 = v588
	v559 = v589
	v561 = v591
	v562 = v592
	goto L108
L160:
	;
	v539 = v569
	v540 = v570
	v541 = v571
	v542 = v572
	v551 = v581
	v553 = v583
	v554 = v584
	v555 = v585
	v556 = v586
	v557 = v587
	v558 = v588
	v559 = v589
	v561 = v591
	v562 = v592
	goto L108
L161:
	;
	v539 = v569
	v540 = v570
	v541 = v571
	v542 = v572
	v551 = v581
	v553 = v583
	v554 = v584
	v555 = v585
	v556 = v586
	v557 = v587
	v558 = v588
	v559 = v589
	v561 = v591
	v562 = v592
	goto L108
L162:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v831))))
	if v832 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v586)+20)) = v831
	*(*int32)(unsafe.Add(mBase, uint32(v586)+16)) = int32(_a_F_check_synchronous_standby_names_18)
	v839 = F_psprintf(m, int32(_a_F_check_synchronous_standby_names_19), v586+int32(16))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L6
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v586))) = int32(_a_F_check_synchronous_standby_names_18)
	v845 = F_psprintf(m, int32(_a_F_check_synchronous_standby_names_20), v586)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L6
	} else {
		goto L167
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v587))) = v839
	v1651 = v804
	goto L140
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v587))) = v845
	v1651 = v804
	goto L140
L168:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v572)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v851))) = v849
	v1651 = int32(258)
	goto L140
L169:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v572)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v857))) = v855
	v1651 = int32(259)
	goto L140
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v876)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v880
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v882))) = v883
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v887 = int32(2)
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v885+v886<<(uint(v887)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v890)+44)) = int32(1)
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v893+v894<<(uint(v887)%32))))
	v899 = v898
	v900 = v893
	v901 = v894
	goto L173
L172:
	;
	v899 = v876
	v900 = v871
	v901 = v872
	goto L173
L173:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v572)+36))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v899)+4))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	v905 = v903 + v904
	if base.Ui32(v902) <= base.Ui32(v905) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	v910 = v868 ^ int32(-1) + v750
	v911 = v907 + v910
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v911
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v572)+44))
	if int32(0) < v910 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	goto L176
L176:
	;
	if base.Ui32(v905+int32(1)) < base.Ui32(v902) {
		goto L139
	} else {
		goto L208
	}
L177:
	;
	v920 = v913
	v924 = v907
	goto L180
L178:
	;
	v1060 = v913
	goto L179
L179:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v1060-int32(31)) {
		goto L198
	} else {
		goto L199
	}
L180:
	;
	v942 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v924))))
	if v942 != 0 {
		goto L182
	} else {
		goto L183
	}
L181:
	;
	v1060 = v1052
	goto L179
L182:
	;
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v942)+uint32(_c_F_check_synchronous_standby_names[6]))))
	v945 = v943
	goto L184
L183:
	;
	v945 = int32(1)
	goto L184
L184:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v920-int32(31)) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+68)) = v924
	*(*int32)(unsafe.Add(mBase, uint32(v572)+64)) = v920
	goto L187
L186:
	;
	goto L187
L187:
	;
	v953 = v945 & int32(255)
	v954 = int32(1)
	v958 = int32(*(*int16)(unsafe.Add(mBase, uint32(v920<<(uint(v954)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v959 = v953 + v958
	v964 = int32(*(*int16)(unsafe.Add(mBase, uint32(v959<<(uint(v954)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v964 != v920 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v970 = v920
	v975 = v945
	v976 = v953
	goto L191
L189:
	;
	v1028 = v959
	goto L190
L190:
	;
	v1048 = int32(1)
	v1052 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1028<<(uint(v1048)%32))+uint32(_c_F_check_synchronous_standby_names[9]))))
	v1054 = v924 + v1048
	if v1054 != v911 {
		v920 = v1052
		v924 = v1054
		goto L180
	} else {
		goto L197
	}
L191:
	;
	v996 = int32(*(*int16)(unsafe.Add(mBase, uint32(v970<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[10]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v970))%64)&int64(2076672024) != int64(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v1028 = v1013
	goto L190
L193:
	;
	v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v976)+uint32(_c_F_check_synchronous_standby_names[11]))))
	v1005 = v1004
	goto L195
L194:
	;
	v1005 = v975
	goto L195
L195:
	;
	v1007 = v1005 & int32(255)
	v1008 = int32(1)
	v1012 = int32(*(*int16)(unsafe.Add(mBase, uint32(v996<<(uint(v1008)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v1013 = v1007 + v1012
	v1018 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1013<<(uint(v1008)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v1018 != v996&int32(_a_F_check_synchronous_standby_names_10) {
		v970 = v996
		v975 = v1005
		v976 = v1007
		goto L191
	} else {
		goto L196
	}
L196:
	;
	goto L192
L197:
	;
	goto L181
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+68)) = v911
	*(*int32)(unsafe.Add(mBase, uint32(v572)+64)) = v1060
	goto L200
L199:
	;
	goto L200
L200:
	;
	v1088 = int32(1)
	v1092 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1060<<(uint(v1088)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v1094 = v1092 + v1088
	v1099 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1094<<(uint(v1088)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v1099 != v1060 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1105 = v1060
	goto L204
L202:
	;
	v1152 = v1094
	goto L203
L203:
	;
	if base.B2i32(v1152 == int32(0))|base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v1152))%64)&int64(-32985348833280) != int64(0)) != 0 {
		v718 = v907
		goto L125
	} else {
		goto L207
	}
L204:
	;
	v1127 = int32(1)
	v1131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1105<<(uint(v1127)%32))+uint32(_c_F_check_synchronous_standby_names[10]))))
	v1132 = base.I32_extend16_s(v1131)
	v1137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1132<<(uint(v1127)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v1139 = v1137 + v1127
	v1144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1139<<(uint(v1127)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v1131 != v1144 {
		v1105 = v1132
		goto L204
	} else {
		goto L206
	}
L205:
	;
	v1152 = v1139
	goto L203
L206:
	;
	goto L205
L207:
	;
	v1182 = int32(1)
	v1183 = v911 + v1182
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v1183
	v1189 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1152<<(uint(v1182)%32))+uint32(_c_F_check_synchronous_standby_names[9]))))
	v573 = v1189
	v574 = v907
	v577 = v1183
	goto L110
L208:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v572)+80))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v899)+40))
	if v1194 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	if v902-v1193 != int32(1) {
		v2081 = v1193
		v2082 = v903
		v2089 = v904
		goto L135
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v1202 = v1193 ^ int32(-1) + v902
	if int32(0) < v1202 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v2260 = v1193
	goto L131
L213:
	;
	v1205 = int32(7)
	v1206 = v1202 & v1205
	if base.Ui32(v902-v1193-int32(2)) < base.Ui32(v1205) {
		goto L218
	} else {
		goto L219
	}
L214:
	;
	v1364 = v899
	v1369 = v900
	v1370 = v901
	goto L215
L215:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+44))
	if v1386 == int32(2) {
		goto L228
	} else {
		goto L229
	}
L216:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1354+v1355<<(uint(int32(2))%32))))
	v1364 = v1359
	v1369 = v1354
	v1370 = v1355
	goto L215
L217:
	;
	v1297 = v1270
	v1299 = v1272
	v1302 = int32(0)
	goto L225
L218:
	;
	v1270 = v1193
	v1272 = v903
	goto L217
L219:
	;
	goto L220
L220:
	;
	v1219 = v1193
	v1221 = v903
	v1224 = int32(0)
	goto L221
L221:
	;
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221))) = uint8(v1241)
	v1243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+1)) = uint8(v1243)
	v1245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+2)) = uint8(v1245)
	v1247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+3)) = uint8(v1247)
	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+4)) = uint8(v1249)
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+5)) = uint8(v1251)
	v1253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+6)) = uint8(v1253)
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221)+7)) = uint8(v1255)
	v1257 = int32(8)
	v1258 = v1221 + v1257
	v1260 = v1219 + v1257
	v1262 = v1224 + v1257
	if v1262 != v1202&int32(2147483640) {
		v1219 = v1260
		v1221 = v1258
		v1224 = v1262
		goto L221
	} else {
		goto L223
	}
L222:
	;
	if v1206 == int32(0) {
		goto L216
	} else {
		goto L224
	}
L223:
	;
	goto L222
L224:
	;
	v1270 = v1260
	v1272 = v1258
	goto L217
L225:
	;
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1297))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1299))) = uint8(v1319)
	v1321 = int32(1)
	v1326 = v1302 + v1321
	if v1326 != v1206 {
		v1297 = v1297 + v1321
		v1299 = v1299 + v1321
		v1302 = v1326
		goto L225
	} else {
		goto L227
	}
L226:
	;
	goto L216
L227:
	;
	goto L226
L228:
	;
	v1389 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v1389
	v1699 = v1389
	v1719 = v1369 + v1370<<(uint(int32(2))%32)
	goto L136
L229:
	;
	goto L230
L230:
	;
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1364)+12))
	v1396 = v1193 - v902
	v1397 = v1395 + v1396
	if v1397 <= int32(0) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v572)+36))
	v1405 = v1364
	v1409 = v1400
	v1411 = v1395
	goto L234
L232:
	;
	v1469 = v1364
	v1471 = v1397
	goto L233
L233:
	;
	v1491 = int32(_a_F_check_synchronous_standby_names_21)
	if base.Ui32(v1491) <= base.Ui32(v1471) {
		goto L250
	} else {
		goto L251
	}
L234:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+20))
	if v1427 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v1469 = v1460
	v1471 = v1462
	goto L233
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1405)+4)) = int32(0)
	goto L132
L237:
	;
	goto L238
L238:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1405)+4))
	v1434 = v1411 << (uint(int32(1)) % 32)
	if v1434 <= int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1438 = base.I32_div_s(v1411, int32(8))
	v1440 = v1438 + v1411
	goto L241
L240:
	;
	v1440 = v1434
	goto L241
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1405)+12)) = v1440
	v1443 = v1440 + int32(2)
	if v1432 != 0 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1405)+4)) = v1448
	if v1448 == int32(0) {
		goto L132
	} else {
		goto L248
	}
L243:
	;
	v1444 = F_repalloc(m, v1432, v1443)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L6
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v1446 = F_palloc(m, v1443)
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L6
	} else {
		goto L247
	}
L246:
	;
	v1448 = v1444
	goto L242
L247:
	;
	v1448 = v1446
	goto L242
L248:
	;
	v1453 = v1448 + (v1409 - v1432)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v1453
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1455+v1456<<(uint(int32(2))%32))))
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1460)+12))
	v1462 = v1461 + v1396
	if v1462 <= int32(0) {
		v1405 = v1460
		v1409 = v1453
		v1411 = v1461
		goto L234
	} else {
		goto L249
	}
L249:
	;
	goto L235
L250:
	;
	v1494 = v1491
	goto L252
L251:
	;
	v1494 = v1471
	goto L252
L252:
	;
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+24))
	if v1496 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1503 = int32(0)
	goto L257
L254:
	;
	goto L255
L255:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = int32(0)
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1571+v1572<<(uint(int32(2))%32))))
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1576)+4))
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1581 = F_fread(m, v1577+v1202, int32(1), v1494, v1580)
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L6
	} else {
		goto L268
	}
L256:
	;
	switch v1527 {
	case 0:
		goto L264
	default:
		v1566 = v1541
		goto L262
	case 11:
		goto L263
	}
L257:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1524 = F_do_getc(m, v1523)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L6
	} else {
		goto L260
	}
L258:
	;
	v1541 = v1494
	goto L256
L259:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1528+v1529<<(uint(int32(2))%32))))
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1533)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1534+v1202+v1503))) = uint8(v1524)
	v1539 = v1503 + int32(1)
	if v1539 != v1494 {
		v1503 = v1539
		goto L257
	} else {
		goto L261
	}
L260:
	;
	v1527 = v1524 + int32(1)
	switch v1527 {
	case 0, 11:
		v1541 = v1503
		goto L256
	default:
		goto L259
	}
L261:
	;
	goto L258
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v1566
	v1668 = v1566
	goto L137
L263:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1554 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v1553+v1554<<(uint(int32(2))%32))))
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1558)+4))
	v1562 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v1559+v1202+v1541))) = uint8(v1562)
	v1566 = v1541 + int32(1)
	goto L262
L264:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1542)))
	goto L265
L265:
	;
	if int32(base.Ui32(v1543)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v1566 = v1541
		goto L262
	} else {
		goto L266
	}
L266:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_13))
	mBase = m.M
	v1552 = m.ExcPending
	if v1552 != 0 {
		goto L6
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	v1589 = v1581
	goto L269
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v1589
	if v1589 != 0 {
		v1668 = v1589
		goto L137
	} else {
		goto L271
	}
L271:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1610)))
	goto L272
L272:
	;
	if int32(base.Ui32(v1611)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v1668 = int32(0)
	goto L137
L274:
	;
	goto L275
L275:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	if v1620 != int32(27) {
		goto L138
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = int32(0)
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v1626)))
	*(*int32)(unsafe.Add(mBase, uint32(v1626))) = v1627 & int32(-49)
	goto L277
L277:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1631+v1632<<(uint(int32(2))%32))))
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1636)+4))
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1641 = F_fread(m, v1637+v1202, int32(1), v1494, v1640)
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L6
	} else {
		goto L278
	}
L278:
	;
	v1589 = v1641
	goto L269
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	v1848 = v1847 + v1202
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1849+v1850<<(uint(int32(2))%32))))
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1854)+12))
	if v1855 < v1848 {
		goto L306
	} else {
		goto L307
	}
L283:
	;
	if v1202 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	if v1727 != 0 {
		goto L289
	} else {
		goto L290
	}
L285:
	;
	goto L286
L286:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1834 = int32(2)
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1832+v1833<<(uint(v1834)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1837)+44)) = v1834
	v1846 = v1834
	goto L282
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1796)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1796))) = v1726
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	if v1801 != 0 {
		goto L302
	} else {
		goto L303
	}
L288:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1749+v1752<<(uint(int32(2))%32))))
	if v1756 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L289:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v1727+v1728<<(uint(int32(2))%32))))
	if v1732 != 0 {
		v1749 = v1727
		goto L288
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	F_syncrep_yyensure_buffer_stack(m, v572)
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L6
	} else {
		goto L293
	}
L292:
	;
	goto L291
L293:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v572)+4))
	v1736 = F_syncrep_yy_create_buffer(m, v1735, v572)
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L6
	} else {
		goto L294
	}
L294:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1738+v1739<<(uint(int32(2))%32)))) = v1736
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	if v1744 != 0 {
		v1749 = v1744
		goto L288
	} else {
		goto L295
	}
L295:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	v1796 = int32(0)
	v1797 = v1746
	goto L287
L296:
	;
	v1796 = int32(0)
	v1797 = v1751
	goto L287
L297:
	;
	goto L298
L298:
	;
	v1760 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+16)) = v1760
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1762))) = uint8(v1760)
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1765)+1)) = uint8(v1760)
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+44)) = v1760
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+28)) = int32(1)
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1756)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+8)) = v1772
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	if v1774 == v1760 {
		v1796 = v1756
		v1797 = v1751
		goto L287
	} else {
		goto L299
	}
L299:
	;
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1780 = v1774 + v1777<<(uint(int32(2))%32)
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v1780)))
	if v1756 != v1781 {
		v1796 = v1756
		v1797 = v1751
		goto L287
	} else {
		goto L300
	}
L300:
	;
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v1781)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v1783
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v1780)))
	v1786 = *(*int32)(unsafe.Add(mBase, uint32(v1785)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+80)) = v1786
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v1786
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1780)))
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v1789)))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+4)) = v1790
	v1792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1786))))
	*(*uint8)(unsafe.Add(mBase, uint32(v572)+24)) = uint8(v1792)
	v1796 = v1756
	v1797 = v1751
	goto L287
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1796)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = v1797
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1818 = v1814 + v1815<<(uint(int32(2))%32)
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v1819)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v1820
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v1823
	*(*int32)(unsafe.Add(mBase, uint32(v572)+80)) = v1823
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1826)))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+4)) = v1827
	v1829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1823))))
	*(*uint8)(unsafe.Add(mBase, uint32(v572)+24)) = uint8(v1829)
	v1846 = int32(1)
	goto L282
L302:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1801+v1802<<(uint(int32(2))%32))))
	if v1796 == v1806 {
		goto L301
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1796)+32)) = int64(1)
	goto L301
L305:
	;
	goto L304
L306:
	;
	v1859 = v1848 + v1847>>(uint(int32(1))%32)
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1854)+4))
	if v1860 != 0 {
		goto L310
	} else {
		goto L311
	}
L307:
	;
	v1889 = v1849
	v1891 = v1848
	v1892 = v1850
	goto L308
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+28)) = v1891
	v1894 = int32(2)
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1889+v1892<<(uint(v1894)%32))))
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1897)+4))
	v1900 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1898+v1891))) = uint8(v1900)
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1902+v1903<<(uint(v1894)%32))))
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1907)+4))
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1908+v1909)+1)) = uint8(v1900)
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1917 = v1913 + v1914<<(uint(v1894)%32)
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v1917)))
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v1918)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v572)+80)) = v1919
	if v1846 == int32(1) {
		v2260 = v1919
		goto L131
	} else {
		goto L316
	}
L309:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1868 = int32(2)
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1866+v1867<<(uint(v1868)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1871)+4)) = v1865
	v1873 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1873+v1874<<(uint(v1868)%32))))
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1878)+4))
	if v1879 == int32(0) {
		goto L133
	} else {
		goto L315
	}
L310:
	;
	v1861 = F_repalloc(m, v1860, v1859)
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L6
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	v1863 = F_palloc(m, v1859)
	mBase = m.M
	v1864 = m.ExcPending
	if v1864 != 0 {
		goto L6
	} else {
		goto L314
	}
L313:
	;
	v1865 = v1861
	goto L309
L314:
	;
	v1865 = v1863
	goto L309
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1878)+12)) = v1859 - int32(2)
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v572)+12))
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v572)+20))
	v1889 = v1888
	v1891 = v1886 + v1202
	v1892 = v1885
	goto L308
L316:
	;
	switch v1846 - int32(1) {
	case 0:
		goto L134
	case 1:
		goto L317
	default:
		goto L318
	}
L317:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, uint32(v572)+28))
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v1917)))
	v2075 = *(*int32)(unsafe.Add(mBase, uint32(v2074)+4))
	v2081 = v1919
	v2082 = v2075
	v2089 = v2073
	goto L135
L318:
	;
	v1927 = v868 ^ int32(-1) + v750
	v1928 = v1919 + v1927
	*(*int32)(unsafe.Add(mBase, uint32(v572)+36)) = v1928
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v572)+44))
	if v1927 <= int32(0) {
		v573 = v1930
		v574 = v1919
		v577 = v1928
		goto L110
	} else {
		goto L319
	}
L319:
	;
	v1937 = v1930
	v1944 = v1919
	goto L320
L320:
	;
	v1959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1944))))
	if v1959 != 0 {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	v573 = v2069
	v574 = v1919
	v577 = v1928
	goto L110
L322:
	;
	v1960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1959)+uint32(_c_F_check_synchronous_standby_names[6]))))
	v1962 = v1960
	goto L324
L323:
	;
	v1962 = int32(1)
	goto L324
L324:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v1937-int32(31)) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+68)) = v1944
	*(*int32)(unsafe.Add(mBase, uint32(v572)+64)) = v1937
	goto L327
L326:
	;
	goto L327
L327:
	;
	v1970 = v1962 & int32(255)
	v1971 = int32(1)
	v1975 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1937<<(uint(v1971)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v1976 = v1970 + v1975
	v1981 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1976<<(uint(v1971)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v1981 != v1937 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	v1987 = v1937
	v1992 = v1962
	v1993 = v1970
	goto L331
L329:
	;
	v2045 = v1976
	goto L330
L330:
	;
	v2065 = int32(1)
	v2069 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2045<<(uint(v2065)%32))+uint32(_c_F_check_synchronous_standby_names[9]))))
	v2071 = v1944 + v2065
	if v1928 != v2071 {
		v1937 = v2069
		v1944 = v2071
		goto L320
	} else {
		goto L337
	}
L331:
	;
	v2013 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1987<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[10]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v1987))%64)&int64(2076672024) != int64(0) {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	v2045 = v2030
	goto L330
L333:
	;
	v2021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1993)+uint32(_c_F_check_synchronous_standby_names[11]))))
	v2022 = v2021
	goto L335
L334:
	;
	v2022 = v1992
	goto L335
L335:
	;
	v2024 = v2022 & int32(255)
	v2025 = int32(1)
	v2029 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2013<<(uint(v2025)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v2030 = v2024 + v2029
	v2035 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2030<<(uint(v2025)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v2035 != v2013&int32(_a_F_check_synchronous_standby_names_10) {
		v1987 = v2013
		v1992 = v2022
		v1993 = v2024
		goto L331
	} else {
		goto L336
	}
L336:
	;
	goto L332
L337:
	;
	goto L321
L338:
	;
	v2110 = v2104
	v2114 = v2081
	goto L339
L339:
	;
	v2132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2114))))
	if v2132 != 0 {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	v743 = v2242
	v746 = v2081
	v750 = v2102
	goto L127
L341:
	;
	v2133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2132)+uint32(_c_F_check_synchronous_standby_names[6]))))
	v2135 = v2133
	goto L343
L342:
	;
	v2135 = int32(1)
	goto L343
L343:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v2110-int32(31)) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v572)+68)) = v2114
	*(*int32)(unsafe.Add(mBase, uint32(v572)+64)) = v2110
	goto L346
L345:
	;
	goto L346
L346:
	;
	v2143 = v2135 & int32(255)
	v2144 = int32(1)
	v2148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2110<<(uint(v2144)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v2149 = v2143 + v2148
	v2154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2149<<(uint(v2144)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v2154 != v2110 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v2160 = v2110
	v2165 = v2135
	v2166 = v2143
	goto L350
L348:
	;
	v2218 = v2149
	goto L349
L349:
	;
	v2238 = int32(1)
	v2242 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2218<<(uint(v2238)%32))+uint32(_c_F_check_synchronous_standby_names[9]))))
	v2244 = v2114 + v2238
	if v2244 != v2102 {
		v2110 = v2242
		v2114 = v2244
		goto L339
	} else {
		goto L356
	}
L350:
	;
	v2186 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2160<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[10]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v2160))%64)&int64(2076672024) != int64(0) {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	v2218 = v2203
	goto L349
L352:
	;
	v2194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2166)+uint32(_c_F_check_synchronous_standby_names[11]))))
	v2195 = v2194
	goto L354
L353:
	;
	v2195 = v2165
	goto L354
L354:
	;
	v2197 = v2195 & int32(255)
	v2198 = int32(1)
	v2202 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2186<<(uint(v2198)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v2203 = v2197 + v2202
	v2208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2203<<(uint(v2198)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if v2208 != v2186&int32(_a_F_check_synchronous_standby_names_10) {
		v2160 = v2186
		v2165 = v2195
		v2166 = v2197
		goto L350
	} else {
		goto L355
	}
L355:
	;
	goto L351
L356:
	;
	goto L340
L357:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L358:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L359:
	;
	v2319 = int32(0)
	v2328 = v2319
	v2329 = v2319
	goto L84
L360:
	;
	goto L361
L361:
	;
	if v2295 == int32(256) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v2550 = v2292
	v2551 = v2293
	v2554 = int32(1)
	v2556 = v2291
	v2563 = v2305
	goto L52
L363:
	;
	goto L364
L364:
	;
	if base.Ui32(int32(262)) < base.Ui32(v2295) {
		v2328 = v2295
		v2329 = int32(2)
		goto L84
	} else {
		goto L365
	}
L365:
	;
	v2327 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2295)+uint32(_c_F_check_synchronous_standby_names[13]))))
	v2328 = v2295
	v2329 = v2327
	goto L84
L366:
	;
	v2333 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2330)+uint32(_c_F_check_synchronous_standby_names[14]))))
	if v2329 != v2333 {
		v2341 = v2291
		v2342 = v2292
		v2343 = v2293
		v2344 = v2294
		v2345 = v2328
		v2353 = v2303
		v2355 = v2305
		v2356 = v2306
		v2357 = v2307
		v2359 = v2309
		v2360 = v2310
		v2361 = v2311
		v2363 = v2313
		v2364 = v2314
		goto L82
	} else {
		goto L367
	}
L367:
	;
	v2335 = *(*int32)(unsafe.Add(mBase, uint32(v2305)+1020))
	*(*int32)(unsafe.Add(mBase, uint32(v2303)+4)) = v2335
	v2340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2330)+uint32(_c_F_check_synchronous_standby_names[15]))))
	v2518 = v2291
	v2519 = v2292
	v2520 = v2293
	v2521 = v2294
	v2522 = int32(-2)
	v2530 = v2303 + int32(4)
	v2532 = v2305
	v2533 = v2306
	v2534 = v2307
	v2536 = v2309
	v2538 = v2311
	v2540 = v2313
	v2544 = v2340
	goto L56
L368:
	;
	v2373 = v2341
	v2374 = v2342
	v2375 = v2343
	v2376 = v2344
	v2378 = int32(1)
	v2381 = int32(_a_F_check_synchronous_standby_names_22)
	v2387 = v2355
	v2391 = v2359
	goto L58
L369:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v2376)+80))
	v2407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2406))))
	if v2407 != 0 {
		goto L373
	} else {
		goto L374
	}
L370:
	;
	goto L371
L371:
	;
	m.G0 = v2401 + int32(32)
	v2550 = v2374
	v2551 = v2375
	v2554 = v2378
	v2556 = v2373
	v2563 = v2387
	goto L52
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2391))) = v2419
	goto L371
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2401)+20)) = v2406
	*(*int32)(unsafe.Add(mBase, uint32(v2401)+16)) = v2381
	v2413 = F_psprintf(m, int32(_a_F_check_synchronous_standby_names_19), v2401+int32(16))
	mBase = m.M
	v2414 = m.ExcPending
	if v2414 != 0 {
		goto L6
	} else {
		goto L376
	}
L374:
	;
	goto L375
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2401))) = v2381
	v2417 = F_psprintf(m, int32(_a_F_check_synchronous_standby_names_20), v2401)
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L6
	} else {
		goto L377
	}
L376:
	;
	v2419 = v2413
	goto L372
L377:
	;
	v2419 = v2417
	goto L372
L378:
	;
	v2493 = v2353 - v2431<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v2493)+4)) = v2490
	v2496 = v2493 + int32(4)
	v2497 = v2356 - v2431
	v2498 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2497))))
	v2501 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2428)+uint32(_c_F_check_synchronous_standby_names[16]))))
	v2504 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2501)+uint32(_c_F_check_synchronous_standby_names[16]))))
	v2505 = v2498 + v2504
	if base.Ui32(v2505) <= base.Ui32(int32(22)) {
		goto L394
	} else {
		goto L395
	}
L379:
	;
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(v2353)))
	v2490 = v2489
	goto L378
L380:
	;
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(8))))
	v2486 = *(*int32)(unsafe.Add(mBase, uint32(v2353)))
	v2487 = F_lappend(m, v2485, v2486)
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L6
	} else {
		goto L392
	}
L381:
	;
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v2353)))
	*(*int32)(unsafe.Add(mBase, uint32(v2355)+8)) = v2475
	*(*int32)(unsafe.Add(mBase, uint32(v2355)+12)) = v2475
	v2481 = F_list_make1_impl(m, int32(1), v2355+int32(8))
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L6
	} else {
		goto L391
	}
L382:
	;
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(12))))
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(4))))
	v2473 = F_create_syncrep_config(m, v2468, v2471, int32(0))
	mBase = m.M
	v2474 = m.ExcPending
	if v2474 != 0 {
		goto L6
	} else {
		goto L390
	}
L383:
	;
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(12))))
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(4))))
	v2464 = F_create_syncrep_config(m, v2459, v2462, int32(1))
	mBase = m.M
	v2465 = m.ExcPending
	if v2465 != 0 {
		goto L6
	} else {
		goto L389
	}
L384:
	;
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(12))))
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v2353-int32(4))))
	v2455 = F_create_syncrep_config(m, v2450, v2453, int32(0))
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L6
	} else {
		goto L388
	}
L385:
	;
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v2353)))
	v2446 = F_create_syncrep_config(m, int32(_a_F_check_synchronous_standby_names_23), v2444, int32(0))
	mBase = m.M
	v2447 = m.ExcPending
	if v2447 != 0 {
		goto L6
	} else {
		goto L387
	}
L386:
	;
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(v2353)))
	*(*int32)(unsafe.Add(mBase, uint32(v2363))) = v2441
	v2490 = v2436
	goto L378
L387:
	;
	v2490 = v2446
	goto L378
L388:
	;
	v2490 = v2455
	goto L378
L389:
	;
	v2490 = v2464
	goto L378
L390:
	;
	v2490 = v2473
	goto L378
L391:
	;
	v2490 = v2481
	goto L378
L392:
	;
	v2490 = v2487
	goto L378
L393:
	;
	v2517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2505)+uint32(_c_F_check_synchronous_standby_names[15]))))
	v2518 = v2341
	v2519 = v2342
	v2520 = v2343
	v2521 = v2344
	v2522 = v2345
	v2530 = v2496
	v2532 = v2355
	v2533 = v2497
	v2534 = v2357
	v2536 = v2359
	v2538 = v2361
	v2540 = v2363
	v2544 = v2517
	goto L56
L394:
	;
	v2508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2505)+uint32(_c_F_check_synchronous_standby_names[14]))))
	if v2508 == v2498&int32(255) {
		goto L393
	} else {
		goto L397
	}
L395:
	;
	goto L396
L396:
	;
	v2514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2501)+uint32(_c_F_check_synchronous_standby_names[17]))))
	v2518 = v2341
	v2519 = v2342
	v2520 = v2343
	v2521 = v2344
	v2522 = v2345
	v2530 = v2496
	v2532 = v2355
	v2533 = v2497
	v2534 = v2357
	v2536 = v2359
	v2538 = v2361
	v2540 = v2363
	v2544 = v2514
	goto L56
L397:
	;
	goto L396
L398:
	;
	F_pfree(m, v2556)
	mBase = m.M
	v2579 = m.ExcPending
	if v2579 != 0 {
		goto L6
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	m.G0 = v2563 + int32(1024)
	v2583 = *(*int32)(unsafe.Add(mBase, uint32(v2551)+44))
	F_replication_scanner_finish(m, v2583)
	mBase = m.M
	v2585 = m.ExcPending
	if v2585 != 0 {
		goto L6
	} else {
		goto L402
	}
L401:
	;
	goto L400
L402:
	;
	if v2554 == int32(0) {
		goto L404
	} else {
		goto L405
	}
L403:
	;
	v2614 = *(*int32)(unsafe.Add(mBase, uint32(v2588)+4))
	if v2614 <= int32(0) {
		goto L416
	} else {
		goto L417
	}
L404:
	;
	v2588 = *(*int32)(unsafe.Add(mBase, uint32(v2551)+40))
	if v2588 != 0 {
		goto L403
	} else {
		goto L407
	}
L405:
	;
	goto L406
L406:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[18])) = int32(16801924)
	goto L408
L407:
	;
	goto L406
L408:
	;
	v2593 = *(*int32)(unsafe.Add(mBase, uint32(v2551)+36))
	v2595 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[19])) = v2595
	goto L409
L409:
	;
	if v2593 != 0 {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[20])) = v2611
	v2673 = v2551
	v2697 = int32(0)
	goto L1
L411:
	;
	v2599 = *(*int32)(unsafe.Add(mBase, uint32(v2551)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v2551)+16)) = v2599
	v2604 = F_format_elog_string(m, int32(_a_F_check_synchronous_standby_names_24), v2551+int32(16))
	mBase = m.M
	v2605 = m.ExcPending
	if v2605 != 0 {
		goto L6
	} else {
		goto L414
	}
L412:
	;
	goto L413
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2551))) = int32(_a_F_check_synchronous_standby_names_25)
	v2609 = F_format_elog_string(m, int32(_a_F_check_synchronous_standby_names_26), v2551)
	mBase = m.M
	v2610 = m.ExcPending
	if v2610 != 0 {
		goto L6
	} else {
		goto L415
	}
L414:
	;
	v2611 = v2604
	goto L410
L415:
	;
	v2611 = v2609
	goto L410
L416:
	;
	v2618 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[19])) = v2618
	goto L419
L417:
	;
	goto L418
L418:
	;
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(v2588)))
	v2634 = F_guc_malloc(m, v2633)
	mBase = m.M
	v2635 = m.ExcPending
	if v2635 != 0 {
		goto L6
	} else {
		goto L421
	}
L419:
	;
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v2551)+40))
	v2622 = *(*int32)(unsafe.Add(mBase, uint32(v2621)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2551)+32)) = v2622
	v2628 = F_format_elog_string(m, int32(_a_F_check_synchronous_standby_names_27), v2551+int32(32))
	mBase = m.M
	v2629 = m.ExcPending
	if v2629 != 0 {
		goto L6
	} else {
		goto L420
	}
L420:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[21])) = v2628
	v2673 = v2551
	v2697 = int32(0)
	goto L1
L421:
	;
	if v2634 == int32(0) {
		v2673 = v2551
		v2697 = int32(0)
		goto L1
	} else {
		goto L422
	}
L422:
	;
	v2638 = *(*int32)(unsafe.Add(mBase, uint32(v2551)+40))
	v2639 = *(*int32)(unsafe.Add(mBase, uint32(v2638)))
	if v2639 != 0 {
		goto L423
	} else {
		goto L424
	}
L423:
	;
	base.MemoryCopy(m, v2634, v2638, v2639)
	goto L425
L424:
	;
	goto L425
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2550))) = v2634
	v2646 = v2551
	goto L2
}
func F_chmod(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	v3 = m.Env.X__syscall_chmod(m, l0, l1)
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v3) {
		*(*int32)(unsafe.Add(mBase, _c_F_chmod[0])) = int32(0) - v3
		v11 = int32(-1)
	} else {
		v11 = v3
	}
	return v11
}
func F_chr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_chr[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = base.I32_wrap_i64(v11)
	if int32(0) <= v15 {
		if v15 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v232 = m.ExcPending
			if v232 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v235 = m.ExcPending
				if v235 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_chr_0), int32(0))
					mBase = m.M
					v239 = m.ExcPending
					if v239 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_chr_1), int32(1064), int32(_a_F_chr_2))
						mBase = m.M
						v244 = m.ExcPending
						if v244 != 0 {
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
			if base.B2i32(v14 != int32(6))|base.B2i32(base.Ui32(v15) < base.Ui32(int32(128))) == int32(0) {
				if base.Ui32(int32(_a_F_chr_3)) <= base.Ui32(v15) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v248 = m.ExcPending
					if v248 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v251 = m.ExcPending
						if v251 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v15
							F_errmsg(m, int32(_a_F_chr_4), v9)
							mBase = m.M
							v255 = m.ExcPending
							if v255 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_chr_1), int32(1083), int32(_a_F_chr_2))
								mBase = m.M
								v260 = m.ExcPending
								if v260 != 0 {
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
					if base.Ui32(int32(2047)) < base.Ui32(v15) {
						v34 = int32(3)
					} else {
						v34 = int32(2)
					}
					if base.Ui32(int32(_a_F_chr_5)) < base.Ui32(v15) {
						v37 = int32(4)
					} else {
						v37 = v34
					}
					v39 = v37 + int32(4)
					v40 = F_palloc(m, v39)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v40))) = v39 << (uint(int32(2)) % 32)
						v48 = v40 + int32(4)
						if v11&int64(2095104) == int64(0) {
							v56 = v15&int32(63) | int32(128)
							*(*uint8)(unsafe.Add(mBase, uint32(v40)+5)) = uint8(v56)
							v61 = int32(base.Ui32(v15)>>(uint(int32(6))%32)) | int32(192)
							*(*uint8)(unsafe.Add(mBase, uint32(v40)+4)) = uint8(v61)
						} else {
							if base.Ui32(v15-int32(2048)) <= base.Ui32(int32(_a_F_chr_6)) {
								v67 = int32(63)
								v69 = int32(128)
								v70 = v15&v67 | v69
								*(*uint8)(unsafe.Add(mBase, uint32(v40)+6)) = uint8(v70)
								v77 = int32(base.Ui32(v15)>>(uint(int32(6))%32))&v67 | v69
								*(*uint8)(unsafe.Add(mBase, uint32(v40)+5)) = uint8(v77)
								v84 = int32(base.Ui32(v15)>>(uint(int32(12))%32))&int32(15) | int32(224)
								*(*uint8)(unsafe.Add(mBase, uint32(v40)+4)) = uint8(v84)
							} else {
								v86 = int32(63)
								v88 = int32(128)
								v89 = v15&v86 | v88
								*(*uint8)(unsafe.Add(mBase, uint32(v40)+7)) = uint8(v89)
								v94 = int32(base.Ui32(v15)>>(uint(int32(18))%32)) | int32(240)
								*(*uint8)(unsafe.Add(mBase, uint32(v40)+4)) = uint8(v94)
								v101 = int32(base.Ui32(v15)>>(uint(int32(6))%32))&v86 | v88
								*(*uint8)(unsafe.Add(mBase, uint32(v40)+6)) = uint8(v101)
								v108 = int32(base.Ui32(v15)>>(uint(int32(12))%32))&v86 | v88
								*(*uint8)(unsafe.Add(mBase, uint32(v40)+5)) = uint8(v108)
							}
						}
						v110 = int32(0)
						switch v37 - int32(1) {
						case 0:
							v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
							v146 = v145
							if base.I32_extend8_s(v146) < int32(-62) {
								v159 = v110
							} else {
								v151 = v146
								v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
							}
						case 1:
							v119 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48)+1)))
							v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
							switch v120 - int32(224) {
							case 0:
								v123 = int32(224)
								if base.Ui32(v123) <= base.Ui32((v119-int32(-64))&int32(255)) {
									v151 = v123
									v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
								} else {
									v159 = v110
								}
							default:
								if v119 <= int32(-65) {
									v146 = v120
									if base.I32_extend8_s(v146) < int32(-62) {
										v159 = v110
									} else {
										v151 = v146
										v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
									}
								} else {
									v159 = v110
								}
							case 13:
								if int32(-97) < v119 {
									v159 = v110
								} else {
									v151 = int32(237)
									v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
								}
							case 16:
								if base.Ui32((v119-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
									v159 = v110
								} else {
									v151 = int32(240)
									v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
								}
							case 20:
								if int32(-113) < v119 {
									v159 = v110
								} else {
									v151 = int32(244)
									v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
								}
							}
						case 2:
							v116 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48)+2)))
							if int32(-65) < v116 {
								v159 = v110
							} else {
								v119 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48)+1)))
								v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
								switch v120 - int32(224) {
								case 0:
									v123 = int32(224)
									if base.Ui32(v123) <= base.Ui32((v119-int32(-64))&int32(255)) {
										v151 = v123
										v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
									} else {
										v159 = v110
									}
								default:
									if v119 <= int32(-65) {
										v146 = v120
										if base.I32_extend8_s(v146) < int32(-62) {
											v159 = v110
										} else {
											v151 = v146
											v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
										}
									} else {
										v159 = v110
									}
								case 13:
									if int32(-97) < v119 {
										v159 = v110
									} else {
										v151 = int32(237)
										v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
									}
								case 16:
									if base.Ui32((v119-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
										v159 = v110
									} else {
										v151 = int32(240)
										v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
									}
								case 20:
									if int32(-113) < v119 {
										v159 = v110
									} else {
										v151 = int32(244)
										v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
									}
								}
							}
						case 3:
							v113 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48)+3)))
							if int32(-65) < v113 {
								v159 = v110
							} else {
								v116 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48)+2)))
								if int32(-65) < v116 {
									v159 = v110
								} else {
									v119 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48)+1)))
									v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
									switch v120 - int32(224) {
									case 0:
										v123 = int32(224)
										if base.Ui32(v123) <= base.Ui32((v119-int32(-64))&int32(255)) {
											v151 = v123
											v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
										} else {
											v159 = v110
										}
									default:
										if v119 <= int32(-65) {
											v146 = v120
											if base.I32_extend8_s(v146) < int32(-62) {
												v159 = v110
											} else {
												v151 = v146
												v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
											}
										} else {
											v159 = v110
										}
									case 13:
										if int32(-97) < v119 {
											v159 = v110
										} else {
											v151 = int32(237)
											v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
										}
									case 16:
										if base.Ui32((v119-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
											v159 = v110
										} else {
											v151 = int32(240)
											v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
										}
									case 20:
										if int32(-113) < v119 {
											v159 = v110
										} else {
											v151 = int32(244)
											v159 = base.B2i32(base.Ui32(v151&int32(255)) < base.Ui32(int32(245)))
										}
									}
								}
							}
						default:
							v159 = v110
						}
						if v159 != 0 {
							v205 = v40
							m.G0 = v9 + int32(48)
							return base.I64_extend_i32_u(v205)
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v15
									F_errmsg(m, int32(_a_F_chr_7), v9+int32(16))
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_chr_1), int32(1124), int32(_a_F_chr_2))
										mBase = m.M
										v177 = m.ExcPending
										if v177 != 0 {
											return int64(0)
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
				}
			} else {
				if base.B2i32(v14 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v14)) != 0 {
					v193 = int32(1)
				} else {
					v192 = *(*int32)(unsafe.Add(mBase, uint32(v14*int32(28))+uint32(_c_F_chr[1])))
					v193 = v192
				}
				if int32(1) < v193 {
					v196 = base.B2i32(base.Ui32(v15) < base.Ui32(int32(128)))
				} else {
					v196 = base.B2i32(base.Ui32(v15) < base.Ui32(int32(256)))
				}
				if v196 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v264 = m.ExcPending
					if v264 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v267 = m.ExcPending
						if v267 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v15
							F_errmsg(m, int32(_a_F_chr_4), v9+int32(32))
							mBase = m.M
							v273 = m.ExcPending
							if v273 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_chr_1), int32(1136), int32(_a_F_chr_2))
								mBase = m.M
								v278 = m.ExcPending
								if v278 != 0 {
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
					v200 = F_palloc(m, int32(5))
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return int64(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v200)+4)) = uint8(v11)
						*(*int32)(unsafe.Add(mBase, uint32(v200))) = int32(20)
						v205 = v200
						m.G0 = v9 + int32(48)
						return base.I64_extend_i32_u(v205)
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v216 = m.ExcPending
		if v216 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v219 = m.ExcPending
			if v219 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_chr_8), int32(0))
				mBase = m.M
				v223 = m.ExcPending
				if v223 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_chr_1), int32(1060), int32(_a_F_chr_2))
					mBase = m.M
					v228 = m.ExcPending
					if v228 != 0 {
						return int64(0)
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
