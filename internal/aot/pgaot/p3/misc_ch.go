package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckBuiltinCryptoMode(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBuiltinCryptoMode[0]))
	if v3 == int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_F_CheckBuiltinCryptoMode_0), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_CheckBuiltinCryptoMode_1), int32(735), int32(_a_F_CheckBuiltinCryptoMode_2))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
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
func F_CheckSASLAuth(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
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
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v5
	F_initStringInfo(m, v12-int32(-64))
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	m.T0[v26].(func(*base.Module, int32, int32))(m, l1, v12-int32(-64))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_appendStringInfoChar(m, v12-int32(-64), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)+68))
	F_sendAuthRequest(m, int32(10), v35, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	F_pfree(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v50 = int32(1)
	v52 = v5
	goto L10
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L66
	}
L8:
	;
	m.G0 = v12 + int32(80)
	return v194
L9:
	;
	v194 = int32(-1)
	goto L8
L10:
	;
	F_pq_startmsgread(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	if v145 == int32(1) {
		v194 = int32(0)
		goto L8
	} else {
		goto L65
	}
L12:
	;
	v55 = F_pq_getbyte(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v55 != int32(112) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v55 == int32(-1) {
		v194 = int32(-2)
		goto L8
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_initStringInfo(m, v12+int32(48))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L22
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v55
	F_errmsg(m, int32(_a_F_CheckSASLAuth_0), v12)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(90), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v84 = F_pq_getmessage(m, v12+int32(48), v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v84 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	F_pfree(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v91 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L9
L28:
	;
	if v91 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v93
	F_errmsg_internal(m, int32(_a_F_CheckSASLAuth_3), v12+int32(32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v50&int32(1) != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(105), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_pq_getmsgend(m, v12+int32(48))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L46
	}
L35:
	;
	v130 = F_pq_getmsgbytes(m, v12+int32(48), v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L45
	}
L36:
	;
	v112 = F_pq_getmsgrawstring(m, v12+int32(48))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v128 = v52
	v129 = v125
	goto L35
L39:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v115 = m.T0[v114].(func(*base.Module, int32, int32, int32) int32)(m, l1, v112, l2)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v120 = F_pq_getmsgint(m, v12+int32(48), int32(4))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v120 != int32(-1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v128 = v115
	v129 = v120
	goto L35
L43:
	;
	goto L44
L44:
	;
	v132 = int32(-1)
	v134 = v115
	v135 = int32(0)
	goto L34
L45:
	;
	v132 = v129
	v134 = v128
	v135 = v130
	goto L34
L46:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = m.T0[v144].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v134, v135, v132, v12+int32(44), v12+int32(40), l3)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	F_pfree(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v150 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v145 == int32(2) {
		goto L7
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v181 = int32(0)
	if v145 == v181 {
		v50 = v181
		v52 = v134
		goto L10
	} else {
		goto L64
	}
L52:
	;
	v155 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v155 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v157
	F_errmsg_internal(m, int32(_a_F_CheckSASLAuth_4), v12+int32(16))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v145 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(176), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v173 = int32(12)
	goto L61
L60:
	;
	v173 = int32(11)
	goto L61
L61:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	F_sendAuthRequest(m, v173, v174, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	F_pfree(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L51
L64:
	;
	goto L11
L65:
	;
	goto L9
L66:
	;
	F_errmsg_internal(m, int32(_a_F_CheckSASLAuth_5), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_CheckSASLAuth_1), int32(171), int32(_a_F_CheckSASLAuth_2))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
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
				v42 = v13
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
				v17 = *(*int32)(unsafe.Add(mBase, _c_F_char2wchar[0]))
				if v14 != 0 {
					if v14 == int32(-1) {
						v22 = int32(_a_F_char2wchar_0)
					} else {
						v22 = v14
					}
					*(*int32)(unsafe.Add(mBase, _c_F_char2wchar[0])) = v22
				} else {
				}
				if v17 == int32(_a_F_char2wchar_0) {
					v27 = int32(-1)
				} else {
					v27 = v17
				}
				v28 = F_mbstowcs(m, l0, v9, l1)
				mBase = m.M
				if v27 != 0 {
					if v27 == int32(-1) {
						v36 = int32(_a_F_char2wchar_0)
					} else {
						v36 = v27
					}
					*(*int32)(unsafe.Add(mBase, _c_F_char2wchar[0])) = v36
				} else {
				}
				v42 = v28
			}
			F_pfree(m, v9)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				if v42 != int32(-1) {
					return
				} else {
					F_pg_verifymbstr(m, l2, l3)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							F_errcode(m, int32(17301634))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_char2wchar_1), int32(0))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_char2wchar_2), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_char2wchar_3), int32(1001), int32(_a_F_char2wchar_4))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
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
func F_charne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 != v3)
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
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
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
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
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
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
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
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v451 int32
	_ = v451
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	if l0 == v3 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v16 + int32(48)
	return base.B2i32(v589 == int32(0)) & v591
L2:
	;
	v582 = F_raw_expression_tree_walker_impl(m, v568, int32(485), l1)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L23
	} else {
		goto L169
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L23
	} else {
		goto L164
	}
L4:
	;
	v589 = v531
	v591 = int32(0)
	goto L1
L5:
	;
	v531 = v3
	goto L4
L6:
	;
	goto L7
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 <= int32(109) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v335)+8))
	if v348 != 0 {
		v589 = v340
		v591 = v3
		goto L1
	} else {
		goto L120
	}
L9:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v194)+64))
	if v207 != 0 {
		goto L89
	} else {
		goto L90
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L23
	} else {
		goto L86
	}
L11:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82+l0)))
	if v84 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v41 {
	case 0:
		goto L26
	case 1:
		goto L27
	case 2:
		goto L28
	case 3:
		goto L29
	default:
		v165 = l0
		goto L10
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(2)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v35 = F_checkWellFormedRecursionWalker(m, v34, l1)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L23
	} else {
		goto L24
	}
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	switch v20 - int32(3) {
	case 0:
		v335 = l0
		v337 = v23
		v340 = v3
		goto L8
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
		v568 = l0
		v573 = v3
		goto L2
	case 19:
		goto L13
	default:
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v20 == int32(110) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v20 == int32(64) {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v568 = l0
	v573 = v3
	goto L2
L19:
	;
	v589 = v3
	v591 = v3
	goto L1
L20:
	;
	goto L21
L21:
	;
	if v20 == int32(141) {
		v194 = l0
		v199 = v3
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v568 = l0
	v573 = v3
	goto L2
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	v82 = int32(12)
	goto L11
L25:
	;
	v82 = int32(28)
	goto L11
L26:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v76 = F_checkWellFormedRecursionWalker(m, v75, l1)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L23
	} else {
		goto L45
	}
L27:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v65 = F_checkWellFormedRecursionWalker(m, v64, l1)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L23
	} else {
		goto L40
	}
L28:
	;
	if v23 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	if v23 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v47 = F_checkWellFormedRecursionWalker(m, v46, l1)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L23
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v51 = F_checkWellFormedRecursionWalker(m, v50, l1)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L23
	} else {
		goto L34
	}
L34:
	;
	goto L25
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v58 = F_checkWellFormedRecursionWalker(m, v57, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L23
	} else {
		goto L38
	}
L38:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v61 = F_checkWellFormedRecursionWalker(m, v60, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L23
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	goto L25
L40:
	;
	if v23 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L43
L42:
	;
	goto L43
L43:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v72 = F_checkWellFormedRecursionWalker(m, v71, l1)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L23
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v23
	goto L25
L45:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v79 = F_checkWellFormedRecursionWalker(m, v78, l1)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	goto L25
L47:
	;
	v589 = int32(1)
	v591 = v3
	goto L1
L48:
	;
	goto L49
L49:
	;
	v88 = v84
	goto L50
L50:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v102 = int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v103 <= int32(63) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v589 = v102
	v591 = v3
	goto L1
L52:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162+v88)))
	if v164 != 0 {
		v88 = v164
		goto L50
	} else {
		goto L85
	}
L53:
	;
	v162 = int32(28)
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(2)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	v157 = F_checkWellFormedRecursionWalker(m, v156, l1)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L23
	} else {
		goto L84
	}
L55:
	;
	switch v103 - int32(3) {
	case 0:
		v335 = v88
		v337 = v101
		v340 = v102
		goto L8
	default:
		v568 = v88
		v573 = v102
		goto L2
	case 19:
		goto L54
	}
L56:
	;
	goto L57
L57:
	;
	if v103 != int32(64) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	if v103 == int32(110) {
		v589 = v102
		v591 = v3
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	switch v114 {
	case 0:
		goto L66
	case 1:
		goto L65
	case 2:
		goto L64
	case 3:
		goto L63
	default:
		v165 = v88
		goto L10
	}
L61:
	;
	if v103 == int32(141) {
		v194 = v88
		v199 = v102
		goto L9
	} else {
		goto L62
	}
L62:
	;
	v568 = v88
	v573 = v102
	goto L2
L63:
	;
	if v101 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L64:
	;
	if v101 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L65:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v122 = F_checkWellFormedRecursionWalker(m, v121, l1)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L23
	} else {
		goto L69
	}
L66:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v116 = F_checkWellFormedRecursionWalker(m, v115, l1)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v119 = F_checkWellFormedRecursionWalker(m, v118, l1)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L23
	} else {
		goto L68
	}
L68:
	;
	goto L53
L69:
	;
	if v101 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L72
L71:
	;
	goto L72
L72:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v129 = F_checkWellFormedRecursionWalker(m, v128, l1)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L23
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	goto L53
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L76
L75:
	;
	goto L76
L76:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v137 = F_checkWellFormedRecursionWalker(m, v136, l1)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L23
	} else {
		goto L77
	}
L77:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v140 = F_checkWellFormedRecursionWalker(m, v139, l1)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L23
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	goto L53
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(3)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v148 = F_checkWellFormedRecursionWalker(m, v147, l1)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L23
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v152 = F_checkWellFormedRecursionWalker(m, v151, l1)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L23
	} else {
		goto L83
	}
L83:
	;
	goto L53
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v101
	v162 = int32(12)
	goto L52
L85:
	;
	goto L51
L86:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v182
	F_errmsg_internal(m, int32(_a_F_checkWellFormedRecursionWalker_0), v16+int32(32))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L23
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_checkWellFormedRecursionWalker_1), int32(1179), int32(_a_F_checkWellFormedRecursionWalker_2))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L23
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+8)))
	if v208 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	F_checkWellFormedSelectStmt(m, v194, l1)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L23
	} else {
		goto L119
	}
L92:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v213 = F_lcons(m, v211, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L23
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v268 = int32(0)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v271 = F_lcons(m, v268, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L23
	} else {
		goto L105
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v213
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v194)+64))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	if v217 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_checkWellFormedSelectStmt(m, v194, l1)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L23
	} else {
		goto L103
	}
L97:
	;
	v220 = int32(0)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v221 <= v220 {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v226 = v220
	goto L99
L99:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237+v226<<(uint(int32(2))%32))))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+16))
	v243 = F_checkWellFormedRecursionWalker(m, v242, l1)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L23
	} else {
		goto L101
	}
L100:
	;
	goto L96
L101:
	;
	v246 = v226 + int32(1)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v246 < v247 {
		v226 = v246
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v265 = F_list_delete_first(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L23
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v265
	v589 = v199
	v591 = v3
	goto L1
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v271
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v194)+64))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v275 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_checkWellFormedSelectStmt(m, v194, l1)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L23
	} else {
		goto L117
	}
L107:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v278 <= int32(0) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v283 = v268
	goto L109
L109:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294+v283<<(uint(int32(2))%32))))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+16))
	v300 = F_checkWellFormedRecursionWalker(m, v299, l1)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L23
	} else {
		goto L111
	}
L110:
	;
	goto L106
L111:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v302 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+12))
	v305 = v303
	goto L114
L113:
	;
	v305 = int32(0)
	goto L114
L114:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v307 = F_lappend(m, v306, v298)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L23
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = v307
	v311 = v283 + int32(1)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v311 < v312 {
		v283 = v311
		goto L109
	} else {
		goto L116
	}
L116:
	;
	goto L110
L117:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v330 = F_list_delete_first(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L23
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v330
	v531 = v199
	goto L4
L119:
	;
	v589 = v199
	v591 = v3
	goto L1
L120:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v349 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v335)+12))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v467+v468*int32(12))))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473))))
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466))))
	if v477 == int32(0) {
		v496 = v476
		v497 = v477
		goto L149
	} else {
		goto L150
	}
L122:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	if v352 <= int32(0) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v355 = int32(0)
	if v355 < v352 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v358 = v352
	goto L126
L125:
	;
	v358 = v355
	goto L126
L126:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v349)+12))
	v368 = v3
	goto L127
L127:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v359+v368<<(uint(int32(2))%32))))
	if v376 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	goto L121
L129:
	;
	v451 = v368 + int32(1)
	if v451 != v358 {
		v368 = v451
		goto L127
	} else {
		goto L147
	}
L130:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	if v379 <= int32(0) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v382 = int32(0)
	if v382 < v379 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v385 = v379
	goto L134
L133:
	;
	v385 = v382
	goto L134
L134:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v335)+12))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	v392 = int32(0)
	goto L135
L135:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v387+v392<<(uint(int32(2))%32))))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386))))
	if v410 == int32(0) {
		v429 = v409
		v430 = v410
		goto L138
	} else {
		goto L139
	}
L136:
	;
	goto L129
L137:
	;
	if v430-v429 == int32(0) {
		v531 = v340
		goto L4
	} else {
		goto L145
	}
L138:
	;
	goto L137
L139:
	;
	if v409 != v410 {
		v429 = v409
		v430 = v410
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v414 = v386
	v415 = v406
	goto L141
L141:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+1)))
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414)+1)))
	if v419 == int32(0) {
		v429 = v418
		v430 = v419
		goto L138
	} else {
		goto L143
	}
L142:
	;
	v429 = v418
	v430 = v419
	goto L138
L143:
	;
	v422 = int32(1)
	if v418 == v419 {
		v414 = v414 + v422
		v415 = v415 + v422
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v435 = v392 + int32(1)
	if v435 != v385 {
		v392 = v435
		goto L135
	} else {
		goto L146
	}
L146:
	;
	goto L136
L147:
	;
	goto L128
L148:
	;
	if v497-v496 != 0 {
		v589 = v340
		v591 = v3
		goto L1
	} else {
		goto L156
	}
L149:
	;
	goto L148
L150:
	;
	if v476 != v477 {
		v496 = v476
		v497 = v477
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v481 = v466
	v482 = v473
	goto L152
L152:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+1)))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481)+1)))
	if v486 == int32(0) {
		v496 = v485
		v497 = v486
		goto L149
	} else {
		goto L154
	}
L153:
	;
	v496 = v485
	v497 = v486
	goto L149
L154:
	;
	v489 = int32(1)
	if v485 == v486 {
		v481 = v481 + v489
		v482 = v482 + v489
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	if v337 != 0 {
		goto L3
	} else {
		goto L157
	}
L157:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v501 = v499 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v501
	if v501 < int32(2) {
		v589 = v340
		v591 = v3
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L23
	} else {
		goto L159
	}
L159:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L23
	} else {
		goto L160
	}
L160:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v512
	F_errmsg(m, int32(_a_F_checkWellFormedRecursionWalker_3), v16)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L23
	} else {
		goto L161
	}
L161:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v335)+24))
	F_parser_errposition(m, v517, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L23
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_checkWellFormedRecursionWalker_1), int32(1077), int32(_a_F_checkWellFormedRecursionWalker_2))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L23
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
	v546 = m.ExcPending
	if v546 != 0 {
		goto L23
	} else {
		goto L165
	}
L165:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v472)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v548
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v547<<(uint(int32(2))%32))+uint32(_c_F_checkWellFormedRecursionWalker[0])))
	F_errmsg(m, v554, v16+int32(16))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L23
	} else {
		goto L166
	}
L166:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v335)+24))
	F_parser_errposition(m, v559, v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L23
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_checkWellFormedRecursionWalker_1), int32(1069), int32(_a_F_checkWellFormedRecursionWalker_2))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L23
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
	v589 = v573
	v591 = v582
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
		v10 = F_raw_expression_tree_walker_impl(m, l0, int32(485), l1)
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
			v80 = F_raw_expression_tree_walker_impl(m, l0, int32(485), l1)
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = F_pstrdup(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v191
L2:
	;
	return int32(0)
L3:
	;
	v20 = F_SplitIdentifierString(m, v13, int32(44), v10+int32(12))
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
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[1])) = v25
	goto L8
L6:
	;
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v39 == int32(0) {
		v153 = v4
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v31 = F_format_elog_string(m, int32(_a_F_check_createrole_self_grant_0), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[2])) = v31
	F_pfree(m, v13)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_list_free(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v191 = v4
	goto L1
L12:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[1])) = v171
	goto L53
L13:
	;
	F_pfree(m, v13)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L49
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v42 <= int32(0) {
		v153 = v4
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v46 = int32(0)
	v49 = v4
	goto L16
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v46<<(uint(int32(2))%32))))
	v61 = v57
	v62 = int32(_a_F_check_createrole_self_grant_1)
	goto L19
L17:
	;
	v153 = v145
	goto L13
L18:
	;
	if v99 != 0 {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v65 == v66 {
		v88 = v65
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v99 = int32(0)
	goto L18
L21:
	;
	v90 = int32(1)
	if v88 != 0 {
		v61 = v61 + v90
		v62 = v62 + v90
		goto L19
	} else {
		goto L30
	}
L22:
	;
	if base.Ui32((v65-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v76 = v65 | int32(32)
	goto L25
L24:
	;
	v76 = v65
	goto L25
L25:
	;
	if base.Ui32((v66-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v85 = v66 | int32(32)
	goto L28
L27:
	;
	v85 = v66
	goto L28
L28:
	;
	if v76 == v85 {
		v88 = v76
		goto L21
	} else {
		goto L29
	}
L29:
	;
	v99 = v76 - v85
	goto L18
L30:
	;
	goto L20
L31:
	;
	v103 = v57
	v104 = int32(_a_F_check_createrole_self_grant_2)
	goto L35
L32:
	;
	v144 = int32(4)
	goto L33
L33:
	;
	v145 = v49 | v144
	v147 = v46 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v147 < v148 {
		v46 = v147
		v49 = v145
		goto L16
	} else {
		goto L48
	}
L34:
	;
	if v141 != 0 {
		goto L12
	} else {
		goto L47
	}
L35:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v107 == v108 {
		v130 = v107
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v141 = int32(0)
	goto L34
L37:
	;
	v132 = int32(1)
	if v130 != 0 {
		v103 = v103 + v132
		v104 = v104 + v132
		goto L35
	} else {
		goto L46
	}
L38:
	;
	if base.Ui32((v107-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v118 = v107 | int32(32)
	goto L41
L40:
	;
	v118 = v107
	goto L41
L41:
	;
	if base.Ui32((v108-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v127 = v108 | int32(32)
	goto L44
L43:
	;
	v127 = v108
	goto L44
L44:
	;
	if v118 == v127 {
		v130 = v118
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v141 = v118 - v127
	goto L34
L46:
	;
	goto L36
L47:
	;
	v144 = int32(2)
	goto L33
L48:
	;
	goto L17
L49:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_list_free(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v163 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	if v163 == int32(0) {
		v191 = v4
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v153
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v163
	v191 = int32(1)
	goto L1
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v57
	v177 = F_format_elog_string(m, int32(_a_F_check_createrole_self_grant_3), v10)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_createrole_self_grant[2])) = v177
	F_pfree(m, v13)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	F_list_free(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v191 = v4
	goto L1
}
func F_check_debug_io_direct(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = F_pstrdup(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return v240
L2:
	;
	return int32(0)
L3:
	;
	v19 = F_SplitGUCList(m, v13, v10+int32(28))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v19 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[1])) = v24
	goto L8
L6:
	;
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v41 == int32(0) {
		v202 = v4
		goto L13
	} else {
		goto L14
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_check_debug_io_direct_0)
	v33 = F_format_elog_string(m, int32(_a_F_check_debug_io_direct_1), v10+int32(16))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[2])) = v33
	F_pfree(m, v13)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_list_free(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v240 = v4
	goto L1
L12:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[1])) = v220
	goto L66
L13:
	;
	F_pfree(m, v13)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L2
	} else {
		goto L62
	}
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v44 <= int32(0) {
		v202 = v4
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v48 = int32(0)
	v51 = v4
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v48<<(uint(int32(2))%32))))
	v64 = v60
	v65 = int32(_a_F_check_debug_io_direct_2)
	goto L20
L17:
	;
	v202 = v194
	goto L13
L18:
	;
	v194 = v51 | v193
	v196 = v48 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v196 < v197 {
		v48 = v196
		v51 = v194
		goto L16
	} else {
		goto L61
	}
L19:
	;
	if v102 == int32(0) {
		v193 = int32(1)
		goto L18
	} else {
		goto L32
	}
L20:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v68 == v69 {
		v91 = v68
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v102 = int32(0)
	goto L19
L22:
	;
	v93 = int32(1)
	if v91 != 0 {
		v64 = v64 + v93
		v65 = v65 + v93
		goto L20
	} else {
		goto L31
	}
L23:
	;
	if base.Ui32((v68-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v79 = v68 | int32(32)
	goto L26
L25:
	;
	v79 = v68
	goto L26
L26:
	;
	if base.Ui32((v69-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v88 = v69 | int32(32)
	goto L29
L28:
	;
	v88 = v69
	goto L29
L29:
	;
	if v79 == v88 {
		v91 = v79
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v102 = v79 - v88
	goto L19
L31:
	;
	goto L21
L32:
	;
	v109 = v60
	v110 = int32(_a_F_check_debug_io_direct_3)
	goto L34
L33:
	;
	if v147 == int32(0) {
		v193 = int32(2)
		goto L18
	} else {
		goto L46
	}
L34:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v113 == v114 {
		v136 = v113
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v147 = int32(0)
	goto L33
L36:
	;
	v138 = int32(1)
	if v136 != 0 {
		v109 = v109 + v138
		v110 = v110 + v138
		goto L34
	} else {
		goto L45
	}
L37:
	;
	if base.Ui32((v113-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v124 = v113 | int32(32)
	goto L40
L39:
	;
	v124 = v113
	goto L40
L40:
	;
	if base.Ui32((v114-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v133 = v114 | int32(32)
	goto L43
L42:
	;
	v133 = v114
	goto L43
L43:
	;
	if v124 == v133 {
		v136 = v124
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v147 = v124 - v133
	goto L33
L45:
	;
	goto L35
L46:
	;
	v153 = v60
	v154 = int32(_a_F_check_debug_io_direct_4)
	goto L48
L47:
	;
	if v191 != 0 {
		goto L12
	} else {
		goto L60
	}
L48:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v157 == v158 {
		v180 = v157
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v191 = int32(0)
	goto L47
L50:
	;
	v182 = int32(1)
	if v180 != 0 {
		v153 = v153 + v182
		v154 = v154 + v182
		goto L48
	} else {
		goto L59
	}
L51:
	;
	if base.Ui32((v157-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v168 = v157 | int32(32)
	goto L54
L53:
	;
	v168 = v157
	goto L54
L54:
	;
	if base.Ui32((v158-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v177 = v158 | int32(32)
	goto L57
L56:
	;
	v177 = v158
	goto L57
L57:
	;
	if v168 == v177 {
		v180 = v168
		goto L50
	} else {
		goto L58
	}
L58:
	;
	v191 = v168 - v177
	goto L47
L59:
	;
	goto L49
L60:
	;
	v193 = int32(4)
	goto L18
L61:
	;
	goto L17
L62:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_list_free(m, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v212 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v212
	if v212 == int32(0) {
		v240 = v4
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v202
	v240 = int32(1)
	goto L1
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v60
	v226 = F_format_elog_string(m, int32(_a_F_check_debug_io_direct_5), v10)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_debug_io_direct[2])) = v226
	F_pfree(m, v13)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_list_free(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v240 = v4
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
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v56 int32
	_ = v56
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_scanner_yyerror(m, int32(_a_F_check_indirection_0), l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	return l0
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = int32(0)
	if v13 < v10 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v16 = v10
	goto L7
L6:
	;
	v16 = v13
	goto L7
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = v3
	goto L8
L8:
	;
	v30 = v17 + v24<<(uint(int32(2))%32)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 != int32(77) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v42 = v24 + int32(1)
	if v42 != v16 {
		v24 = v42
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v36 = v30 + int32(4)
	if v36 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if base.Ui32(v36) < base.Ui32(v17+v10<<(uint(int32(2))%32)) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	goto L9
L15:
	;
	return int32(0)
L16:
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
							F_errdetail(m, int32(_a_F_check_lateral_ref_ok_2), int32(0))
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
									F_errfinish(m, int32(_a_F_check_lateral_ref_ok_3), int32(509), int32(_a_F_check_lateral_ref_ok_4))
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
								F_errdetail(m, int32(_a_F_check_lateral_ref_ok_2), int32(0))
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
										F_errfinish(m, int32(_a_F_check_lateral_ref_ok_3), int32(509), int32(_a_F_check_lateral_ref_ok_4))
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
										F_errfinish(m, int32(_a_F_check_lateral_ref_ok_3), int32(509), int32(_a_F_check_lateral_ref_ok_4))
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
func F_check_publications_origin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
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
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
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
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = int32(25)
	if l2 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L19
	} else {
		goto L77
	}
L2:
	;
	m.G0 = v11 + int32(80)
	return
L3:
	;
	if l3 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = l3
	v23 = int32(_a_F_check_publications_origin_0)
	goto L6
L5:
	;
	if v60 != 0 {
		goto L2
	} else {
		goto L18
	}
L6:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v26 == v27 {
		v49 = v26
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v60 = int32(0)
	goto L5
L8:
	;
	v51 = int32(1)
	if v49 != 0 {
		v22 = v22 + v51
		v23 = v23 + v51
		goto L6
	} else {
		goto L17
	}
L9:
	;
	if base.Ui32((v26-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = v26 | int32(32)
	goto L12
L11:
	;
	v37 = v26
	goto L12
L12:
	;
	if base.Ui32((v27-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v27 | int32(32)
	goto L15
L14:
	;
	v46 = v27
	goto L15
L15:
	;
	if v37 == v46 {
		v49 = v37
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v60 = v37 - v46
	goto L5
L17:
	;
	goto L7
L18:
	;
	F_initStringInfo(m, v11-int32(-64))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	F_appendStringInfoString(m, v11-int32(-64), int32(_a_F_check_publications_origin_1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_GetPublicationsStr(m, l1, v11-int32(-64), int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_appendStringInfoString(m, v11-int32(-64), int32(_a_F_check_publications_origin_2))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	if int32(0) < l5 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v86 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_check_publications_origin[0]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+60))
	v128 = m.T0[v127].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v121, int32(1), v11+int32(60))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L19
	} else {
		goto L34
	}
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l4+v86<<(uint(int32(2))%32))))
	v95 = F_get_rel_namespace(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L19
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v97 = F_get_namespace_name(m, v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v99 = F_get_rel_name(m, v94)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v97
	F_appendStringInfo(m, v11-int32(-64), int32(_a_F_check_publications_origin_3), v11+int32(48))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	v111 = v86 + int32(1)
	if v111 != l5 {
		v86 = v111
		goto L27
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
	F_pfree(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v133 != int32(2) {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	v138 = F_MakeSingleTupleTableSlot(m, v136, int32(_a_F_check_publications_origin_4))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v143 = F_tuplestore_gettupleslot(m, v140, int32(1), int32(0), v138)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L19
	} else {
		goto L39
	}
L38:
	;
	F_ExecDropSingleTupleTableSlot(m, v138)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L19
	} else {
		goto L63
	}
L39:
	;
	if v143 == int32(0) {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v150 = int32(0)
	goto L41
L41:
	;
	v156 = int32(*(*int16)(unsafe.Add(mBase, uint32(v138)+6)))
	if v156 <= int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v172 == int32(0) {
		goto L38
	} else {
		goto L53
	}
L43:
	;
	F_slot_getsomeattrs_int(m, v138, int32(1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L19
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v164 = F_text_to_cstring(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L19
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	m.T0[v167].(func(*base.Module, int32))(m, v138)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L19
	} else {
		goto L48
	}
L48:
	;
	v170 = F_makeString(m, v164)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L19
	} else {
		goto L49
	}
L49:
	;
	v172 = F_list_append_unique(m, v150, v170)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L19
	} else {
		goto L50
	}
L50:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v177 = F_tuplestore_gettupleslot(m, v174, int32(1), int32(0), v138)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L19
	} else {
		goto L51
	}
L51:
	;
	if v177 != 0 {
		v150 = v172
		goto L41
	} else {
		goto L52
	}
L52:
	;
	goto L42
L53:
	;
	v181 = F_makeStringInfo(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L19
	} else {
		goto L54
	}
L54:
	;
	F_GetPublicationsStr(m, v172, v181, int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L19
	} else {
		goto L55
	}
L55:
	;
	v188 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L19
	} else {
		goto L56
	}
L56:
	;
	if v188 == int32(0) {
		goto L38
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L19
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l6
	F_errmsg(m, int32(_a_F_check_publications_origin_5), v11+int32(16))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v202
	F_errdetail_plural(m, int32(_a_F_check_publications_origin_6), int32(_a_F_check_publications_origin_7), v201, v11)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	F_errhint(m, int32(_a_F_check_publications_origin_8), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L19
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_check_publications_origin_9), int32(2193), int32(_a_F_check_publications_origin_10))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	goto L38
L63:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	if v227 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_pfree(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L19
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	if v230 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L66
L68:
	;
	F_tuplestore_end(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L19
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	if v233 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	F_FreeTupleDesc(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L19
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	F_pfree(m, v128)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L19
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	goto L2
L77:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L19
	} else {
		goto L78
	}
L78:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v256
	F_errmsg(m, int32(_a_F_check_publications_origin_11), v11+int32(32))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L19
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_check_publications_origin_9), int32(2153), int32(_a_F_check_publications_origin_10))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L19
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_check_stack_depth[0]))
	if v8 == int32(0) {
		m.G0 = v5 + int32(16)
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_check_stack_depth[1]))
		v15 = v8 - (v5 + int32(15))
		v17 = v15 >> (uint(int32(31)) % 32)
		if v15^v17-v17 <= v12 {
			m.G0 = v5 + int32(16)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				F_errcode(m, int32(16777477))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_check_stack_depth_0), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_check_stack_depth[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v33
						F_errhint(m, int32(_a_F_check_stack_depth_1), v5)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_stack_depth_2), int32(104), int32(_a_F_check_stack_depth_3))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v15 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return v423
L2:
	;
	v423 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v19 = F_pstrdup(m, v14)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_pfree(m, v19)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L6
	} else {
		goto L106
	}
L6:
	;
	return int32(0)
L7:
	;
	v26 = F_SplitIdentifierString(m, v19, int32(44), v12+int32(32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[1])) = v31
	goto L12
L10:
	;
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v40 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v37 = F_format_elog_string(m, int32(_a_F_check_synchronized_standby_slots_0), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[2])) = v37
	v409 = v4
	goto L5
L14:
	;
	v409 = int32(1)
	goto L5
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if int32(0) < v43 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v47 = int32(0)
	goto L19
L17:
	;
	v114 = v40
	goto L18
L18:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v120 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v47<<(uint(int32(2))%32))))
	v61 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v61
	v71 = F_ReplicationSlotValidateNameInternal(m, v60, v12+int32(44), v12+int32(40), v12+int32(36))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v108 == int32(0) {
		goto L14
	} else {
		goto L32
	}
L21:
	;
	if v71 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[3])) = v75
	goto L25
L23:
	;
	goto L24
L24:
	;
	v105 = v47 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v105 < v106 {
		v47 = v105
		goto L19
	} else {
		goto L31
	}
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[1])) = v79
	goto L26
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v82
	v88 = F_format_elog_string(m, int32(_a_F_check_synchronized_standby_slots_1), v12+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[2])) = v88
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	if v91 == int32(0) {
		v409 = v4
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[1])) = v95
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v91
	v101 = F_format_elog_string(m, int32(_a_F_check_synchronized_standby_slots_1), v12)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronized_standby_slots[4])) = v101
	v409 = v4
	goto L5
L31:
	;
	goto L20
L32:
	;
	v114 = v108
	goto L18
L33:
	;
	v212 = F_guc_malloc(m, v206)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L57
	}
L34:
	;
	v206 = int32(4)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v127 = int32(0)
	v130 = int32(4)
	goto L37
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v124+v127<<(uint(int32(2))%32))))
	if v139&int32(3) == int32(0) {
		v163 = v139
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v206 = v199
	goto L33
L39:
	;
	v198 = int32(1)
	v199 = v196 + v130 + v198
	v201 = v127 + v198
	if v201 != v120 {
		v127 = v201
		v130 = v199
		goto L37
	} else {
		goto L56
	}
L40:
	;
	v196 = v188 - v139
	goto L39
L41:
	;
	v167 = v163
	goto L50
L42:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v147 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v196 = int32(0)
	goto L39
L44:
	;
	goto L45
L45:
	;
	v152 = v139
	goto L46
L46:
	;
	v156 = v152 + int32(1)
	if v156&int32(3) == int32(0) {
		v163 = v156
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v188 = v156
	goto L40
L48:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v161 != 0 {
		v152 = v156
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v176 = int32(-2139062144)
	if (int32(16843008)-v173|v173)&v176 == v176 {
		v167 = v167 + int32(4)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v182 = v167
	goto L53
L52:
	;
	goto L51
L53:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v186 != 0 {
		v182 = v182 + int32(1)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v188 = v182
	goto L40
L55:
	;
	goto L54
L56:
	;
	goto L38
L57:
	;
	if v212 == int32(0) {
		v423 = v4
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v216 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+4))
	v219 = v217
	goto L61
L60:
	;
	v219 = int32(0)
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v221 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v212
	goto L14
L63:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v224 <= int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v230 = int32(0)
	v234 = v212 + int32(4)
	goto L65
L65:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239+v230<<(uint(int32(2))%32))))
	if (v243^v234)&int32(3) != 0 {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	goto L62
L67:
	;
	if v243&int32(3) == int32(0) {
		v341 = v243
		goto L90
	} else {
		goto L91
	}
L68:
	;
	goto L67
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v297)
	if v297&int32(255) == int32(0) {
		goto L68
	} else {
		goto L84
	}
L70:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v296 = v243
	v297 = v249
	v298 = v234
	goto L69
L71:
	;
	goto L72
L72:
	;
	if v243&int32(3) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v253 = v243
	v255 = v234
	goto L76
L74:
	;
	v267 = v243
	v269 = v234
	goto L75
L75:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v274 = int32(-2139062144)
	if (int32(16843008)-v271|v271)&v274 != v274 {
		v296 = v267
		v297 = v271
		v298 = v269
		goto L69
	} else {
		goto L80
	}
L76:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	*(*uint8)(unsafe.Add(mBase, uint32(v255))) = uint8(v256)
	if v256 == int32(0) {
		goto L68
	} else {
		goto L78
	}
L77:
	;
	v267 = v263
	v269 = v261
	goto L75
L78:
	;
	v260 = int32(1)
	v261 = v255 + v260
	v263 = v253 + v260
	if v263&int32(3) != 0 {
		v253 = v263
		v255 = v261
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v279 = v267
	v280 = v271
	v281 = v269
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281))) = v280
	v283 = int32(4)
	v284 = v281 + v283
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	v287 = v279 + v283
	v291 = int32(-2139062144)
	if (v285|(int32(16843008)-v285))&v291 == v291 {
		v279 = v287
		v280 = v285
		v281 = v284
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v296 = v287
	v297 = v285
	v298 = v284
	goto L69
L83:
	;
	goto L82
L84:
	;
	v305 = v296
	v307 = v298
	goto L85
L85:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v307)+1)) = uint8(v308)
	v310 = int32(1)
	if v308 != 0 {
		v305 = v305 + v310
		v307 = v307 + v310
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L68
L87:
	;
	goto L86
L88:
	;
	v376 = int32(1)
	v379 = v230 + v376
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v379 < v380 {
		v230 = v379
		v234 = v234 + v374 + v376
		goto L65
	} else {
		goto L105
	}
L89:
	;
	v374 = v366 - v243
	goto L88
L90:
	;
	v345 = v341
	goto L99
L91:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	if v325 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v374 = int32(0)
	goto L88
L93:
	;
	goto L94
L94:
	;
	v330 = v243
	goto L95
L95:
	;
	v334 = v330 + int32(1)
	if v334&int32(3) == int32(0) {
		v341 = v334
		goto L90
	} else {
		goto L97
	}
L96:
	;
	v366 = v334
	goto L89
L97:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v339 != 0 {
		v330 = v334
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v354 = int32(-2139062144)
	if (int32(16843008)-v351|v351)&v354 == v354 {
		v345 = v345 + int32(4)
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v360 = v345
	goto L102
L101:
	;
	goto L100
L102:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360))))
	if v364 != 0 {
		v360 = v360 + int32(1)
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v366 = v360
	goto L89
L104:
	;
	goto L103
L105:
	;
	goto L66
L106:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	F_list_free(m, v413)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	v423 = v409
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
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
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
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
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v710 int32
	_ = v710
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v792 int32
	_ = v792
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1070 int32
	_ = v1070
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1104 int32
	_ = v1104
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1179 int32
	_ = v1179
	var v1186 int32
	_ = v1186
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1230 int32
	_ = v1230
	var v1254 int32
	_ = v1254
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1403 int32
	_ = v1403
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1436 int32
	_ = v1436
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1487 int32
	_ = v1487
	var v1504 int32
	_ = v1504
	var v1509 int32
	_ = v1509
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1576 int32
	_ = v1576
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1646 int32
	_ = v1646
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1667 int32
	_ = v1667
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1705 int32
	_ = v1705
	var v1709 int32
	_ = v1709
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1715 int32
	_ = v1715
	var v1719 int32
	_ = v1719
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1744 int32
	_ = v1744
	var v1750 int32
	_ = v1750
	var v1753 int32
	_ = v1753
	var v1758 int32
	_ = v1758
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1789 int32
	_ = v1789
	var v1811 int32
	_ = v1811
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1847 int32
	_ = v1847
	var v1851 int32
	_ = v1851
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1885 int32
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1897 int32
	_ = v1897
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1976 int32
	_ = v1976
	var v1977 int32
	_ = v1977
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1990 int32
	_ = v1990
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2027 int32
	_ = v2027
	var v2036 int32
	_ = v2036
	var v2048 int32
	_ = v2048
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2074 int32
	_ = v2074
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2106 int32
	_ = v2106
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2124 int32
	_ = v2124
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2134 int32
	_ = v2134
	var v2140 int32
	_ = v2140
	var v2162 int32
	_ = v2162
	var v2166 int32
	_ = v2166
	var v2168 int32
	_ = v2168
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2177 int32
	_ = v2177
	var v2179 int32
	_ = v2179
	var v2189 int32
	_ = v2189
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2208 int32
	_ = v2208
	var v2211 int32
	_ = v2211
	var v2229 int32
	_ = v2229
	var v2232 int32
	_ = v2232
	var v2234 int32
	_ = v2234
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2255 int32
	_ = v2255
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2287 int32
	_ = v2287
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2315 int32
	_ = v2315
	var v2321 int32
	_ = v2321
	var v2343 int32
	_ = v2343
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2353 int32
	_ = v2353
	var v2358 int32
	_ = v2358
	var v2365 int32
	_ = v2365
	var v2388 int32
	_ = v2388
	var v2392 int32
	_ = v2392
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2400 int32
	_ = v2400
	var v2406 int32
	_ = v2406
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2410 int32
	_ = v2410
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2420 int32
	_ = v2420
	var v2423 int32
	_ = v2423
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2438 int32
	_ = v2438
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2448 int32
	_ = v2448
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2478 int32
	_ = v2478
	var v2484 int32
	_ = v2484
	var v2487 int32
	_ = v2487
	var v2489 int32
	_ = v2489
	var v2492 int32
	_ = v2492
	var v2495 int32
	_ = v2495
	var v2498 int32
	_ = v2498
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2504 int32
	_ = v2504
	var v2507 int32
	_ = v2507
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2513 int32
	_ = v2513
	var v2516 int32
	_ = v2516
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2522 int32
	_ = v2522
	var v2525 int32
	_ = v2525
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2549 int32
	_ = v2549
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2557 int32
	_ = v2557
	var v2560 int32
	_ = v2560
	var v2561 int32
	_ = v2561
	var v2566 int32
	_ = v2566
	var v2570 int32
	_ = v2570
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2579 int32
	_ = v2579
	var v2585 int32
	_ = v2585
	var v2587 int32
	_ = v2587
	var v2589 int32
	_ = v2589
	var v2591 int32
	_ = v2591
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2599 int32
	_ = v2599
	var v2600 int32
	_ = v2600
	var v2605 int32
	_ = v2605
	var v2619 int32
	_ = v2619
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2645 int32
	_ = v2645
	var v2646 int32
	_ = v2646
	var v2659 int32
	_ = v2659
	var v2661 int32
	_ = v2661
	var v2674 int32
	_ = v2674
	var v2678 int32
	_ = v2678
	var v2680 int32
	_ = v2680
	var v2683 int32
	_ = v2683
	var v2688 int32
	_ = v2688
	var v2690 int32
	_ = v2690
	var v2694 int32
	_ = v2694
	var v2699 int32
	_ = v2699
	var v2700 int32
	_ = v2700
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2709 int32
	_ = v2709
	var v2713 int32
	_ = v2713
	var v2716 int32
	_ = v2716
	var v2717 int32
	_ = v2717
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2733 int32
	_ = v2733
	var v2734 int32
	_ = v2734
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2742 int32
	_ = v2742
	var v2769 int32
	_ = v2769
	var v2793 int32
	_ = v2793
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
	m.G0 = v2769 + int32(48)
	return v2793
L2:
	;
	v2769 = v2742
	v2793 = int32(1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v2742 = v29
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
	v74 = F__emscripten_memset_bulkmem(m, v50, base.I32_extend8_s(int32(0)), int32(96))
	mBase = m.M
	goto L17
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
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v76 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+60)) = v76
	v78 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v75)+52)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v75)+44)) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v75)+36)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v75)+4)) = v78
	*(*int64)(unsafe.Add(mBase, uint32(v75)+12)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v76
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v44
	if v31&int32(3) == v76 {
		v115 = v31
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if base.Ui32(v148) < base.Ui32(int32(-2)) {
		goto L38
	} else {
		goto L39
	}
L19:
	;
	v148 = v140 - v31
	goto L18
L20:
	;
	v119 = v115
	goto L29
L21:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if v99 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v148 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v104 = v31
	goto L25
L25:
	;
	v108 = v104 + int32(1)
	if v108&int32(3) == int32(0) {
		v115 = v108
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v140 = v108
	goto L19
L27:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v113 != 0 {
		v104 = v108
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v128 = int32(-2139062144)
	if (int32(16843008)-v125|v125)&v128 == v128 {
		v119 = v119 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v134 = v119
	goto L32
L31:
	;
	goto L30
L32:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v138 != 0 {
		v134 = v134 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v140 = v134
	goto L19
L34:
	;
	goto L33
L35:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v419 = m.G0
	v421 = v419 - int32(1232)
	m.G0 = v421
	v425 = v421 + int32(816)
	v427 = v421 + int32(16)
	v429 = v425
	v430 = l1
	v431 = v29
	v432 = v418
	v434 = int32(-2)
	v440 = v427
	v442 = v425
	v444 = v427
	v446 = v421
	v447 = v4
	v449 = v29 + int32(36)
	v450 = int32(200)
	v454 = v29 + int32(40)
	goto L75
L36:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_3))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L6
	} else {
		goto L72
	}
L37:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_4))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L71
	}
L38:
	;
	v153 = v148 + int32(2)
	v154 = F_palloc(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_5))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L6
	} else {
		goto L70
	}
L41:
	;
	if v154 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	if v148 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v309 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v154+v148))) = uint16(v309)
	if base.Ui32(v153) < base.Ui32(int32(2)) {
		v397 = v309
		goto L57
	} else {
		goto L58
	}
L44:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v148) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v172 = v4
	v179 = int32(0)
	goto L48
L46:
	;
	v225 = v4
	goto L47
L47:
	;
	v244 = v148 & int32(3)
	if v244 == int32(0) {
		goto L43
	} else {
		goto L51
	}
L48:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v172))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v172))) = uint8(v192)
	v195 = v172 | int32(1)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v195))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v195))) = uint8(v198)
	v201 = v172 | int32(2)
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v201))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v201))) = uint8(v204)
	v207 = v172 | int32(3)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v207))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v207))) = uint8(v210)
	v212 = int32(4)
	v213 = v172 + v212
	v215 = v179 + v212
	if v215 != v148&int32(-4) {
		v172 = v213
		v179 = v215
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v225 = v213
	goto L47
L50:
	;
	goto L49
L51:
	;
	v250 = v4
	v255 = v225
	goto L52
L52:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v255))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154+v255))) = uint8(v275)
	v277 = int32(1)
	v280 = v250 + v277
	if v280 != v244 {
		v250 = v280
		v255 = v255 + v277
		goto L52
	} else {
		goto L54
	}
L53:
	;
	goto L43
L54:
	;
	goto L53
L55:
	;
	if v397 == int32(0) {
		goto L36
	} else {
		goto L69
	}
L56:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_6))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L6
	} else {
		goto L68
	}
L57:
	;
	goto L55
L58:
	;
	v315 = v153 - int32(2)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v315))))
	if v317 != 0 {
		v397 = v309
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153+v154-int32(1)))))
	if v321 != 0 {
		v397 = v309
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v323 = F_palloc(m, int32(48))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	if v323 == int32(0) {
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v327 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v323)+20)) = v327
	*(*int32)(unsafe.Add(mBase, uint32(v323)+8)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v323)+4)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v323)+12)) = v315
	*(*int64)(unsafe.Add(mBase, uint32(v323)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v323)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v323)+16)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v327
	F_syncrep_yyensure_buffer_stack(m, v90)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v341+v342<<(uint(int32(2))%32))))
	if v346 == v323 {
		v397 = v323
		goto L57
	} else {
		goto L64
	}
L64:
	;
	if v346 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v90)+36))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v348))) = uint8(v349)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v353 = int32(2)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v351+v352<<(uint(v353)%32))))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v90)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v356)+8)) = v357
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v359+v360<<(uint(v353)%32))))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v90)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+16)) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v369 = v367
	v370 = v368
	goto L67
L66:
	;
	v369 = v341
	v370 = v342
	goto L67
L67:
	;
	v371 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v370<<(uint(v371)%32)+v369))) = v323
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v379 = v375 + v376<<(uint(v371)%32)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = v381
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v90)+80)) = v384
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v388
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+24)) = uint8(v390)
	*(*int32)(unsafe.Add(mBase, uint32(v90)+48)) = int32(1)
	v397 = v323
	goto L57
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397)+20)) = int32(1)
	goto L35
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	if v2661+int32(816) != v2644 {
		goto L429
	} else {
		goto L430
	}
L74:
	;
	F_syncrep_yyerror(m, v449, v432, int32(_a_F_check_synchronous_standby_names_7))
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L6
	} else {
		goto L428
	}
L75:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v442))) = uint16(v447)
	v457 = v450 << (uint(int32(1)) % 32)
	if base.Ui32(v429+v457-int32(2)) <= base.Ui32(v442) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	F_syncrep_yyerror(m, v2473, v2456, int32(_a_F_check_synchronous_standby_names_8))
	mBase = m.M
	v2605 = m.ExcPending
	if v2605 != 0 {
		goto L6
	} else {
		goto L424
	}
L77:
	;
	if base.Ui32(int32(_a_F_check_synchronous_standby_names_9)) < base.Ui32(v450) {
		goto L74
	} else {
		goto L80
	}
L78:
	;
	v510 = v429
	v512 = v440
	v514 = v442
	v515 = v444
	v516 = v450
	goto L79
L79:
	;
	v518 = int32(1) << (uint(v447) % 32)
	if v518&int32(13390146) != 0 {
		v2453 = v510
		v2454 = v430
		v2455 = v431
		v2456 = v432
		v2458 = v434
		v2464 = v512
		v2465 = v518
		v2466 = v514
		v2468 = v515
		v2470 = v446
		v2471 = v447
		v2473 = v449
		v2474 = v516
		v2478 = v454
		goto L103
	} else {
		goto L104
	}
L80:
	;
	v464 = int32(_a_F_check_synchronous_standby_names_10)
	if base.Ui32(v464) <= base.Ui32(v457) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v467 = v464
	goto L83
L82:
	;
	v467 = v457
	goto L83
L83:
	;
	v472 = F_palloc(m, v467*int32(6)|int32(3))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	if v472 == int32(0) {
		goto L74
	} else {
		goto L85
	}
L85:
	;
	v477 = int32(1)
	v480 = (v442-v429)>>(uint(v477)%32) + v477
	v482 = v480 << (uint(v477) % 32)
	if v482 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v487 = v484 + v467<<(uint(int32(1))%32)
	v489 = v480 << (uint(int32(2)) % 32)
	if v489 != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v483 = F__emscripten_memcpy_bulkmem(m, v472, v429, v482)
	mBase = m.M
	v484 = v483
	goto L89
L88:
	;
	v484 = v472
	goto L89
L89:
	;
	goto L86
L90:
	;
	if v446+int32(816) != v429 {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v490 = F__emscripten_memcpy_bulkmem(m, v487, v444, v489)
	mBase = m.M
	v491 = v490
	goto L93
L92:
	;
	v491 = v487
	goto L93
L93:
	;
	goto L90
L94:
	;
	F_pfree(m, v429)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L6
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v497 = int32(1)
	v500 = v484 + v480<<(uint(v497)%32)
	if base.Ui32(v484+v467<<(uint(v497)%32)) <= base.Ui32(v500) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L96
L98:
	;
	v2644 = v484
	v2645 = v430
	v2646 = v431
	v2659 = v497
	v2661 = v446
	goto L73
L99:
	;
	goto L100
L100:
	;
	v510 = v484
	v512 = v491 + v489 - int32(4)
	v514 = v500 - int32(2)
	v515 = v491
	v516 = v467
	goto L79
L101:
	;
	goto L76
L102:
	;
	v429 = v2574
	v430 = v2575
	v431 = v2576
	v432 = v2577
	v434 = v2579
	v440 = v2585
	v442 = v2587 + int32(2)
	v444 = v2589
	v446 = v2591
	v447 = v2600
	v449 = v2594
	v450 = v2595
	v454 = v2599
	goto L75
L103:
	;
	if v2465&int32(_a_F_check_synchronous_standby_names_11) != 0 {
		goto L101
	} else {
		goto L405
	}
L104:
	;
	v523 = int32(*(*int8)(unsafe.Add(mBase, uint32(v447)+uint32(_c_F_check_synchronous_standby_names[1]))))
	if v434 == int32(-2) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v2433 = v523 + v2432
	if base.Ui32(int32(22)) < base.Ui32(v2433) {
		v2453 = v2395
		v2454 = v2396
		v2455 = v2397
		v2456 = v2398
		v2458 = v2431
		v2464 = v2406
		v2465 = v2407
		v2466 = v2408
		v2468 = v2410
		v2470 = v2412
		v2471 = v2413
		v2473 = v2415
		v2474 = v2416
		v2478 = v2420
		goto L103
	} else {
		goto L397
	}
L106:
	;
	v526 = m.G0
	v528 = v526 - int32(32)
	m.G0 = v528
	*(*int32)(unsafe.Add(mBase, uint32(v432)+92)) = v446 + int32(1228)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v432)+40))
	if v533 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v2395 = v510
	v2396 = v430
	v2397 = v431
	v2398 = v432
	v2400 = v434
	v2406 = v512
	v2407 = v518
	v2408 = v514
	v2410 = v515
	v2412 = v446
	v2413 = v447
	v2415 = v449
	v2416 = v516
	v2420 = v454
	goto L108
L108:
	;
	if v2400 <= int32(0) {
		goto L393
	} else {
		goto L394
	}
L109:
	;
	v2395 = v600
	v2396 = v601
	v2397 = v602
	v2398 = v603
	v2400 = v1744
	v2406 = v611
	v2407 = v612
	v2408 = v613
	v2410 = v615
	v2412 = v617
	v2413 = v618
	v2415 = v620
	v2416 = v621
	v2420 = v625
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432)+40)) = int32(1)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v432)+44))
	if v538 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	v600 = v510
	v601 = v430
	v602 = v431
	v603 = v432
	v611 = v512
	v612 = v518
	v613 = v514
	v615 = v515
	v617 = v446
	v618 = v447
	v619 = v528
	v620 = v449
	v621 = v516
	v625 = v454
	goto L129
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432)+44)) = int32(1)
	goto L115
L114:
	;
	goto L115
L115:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	if v543 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v547 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+4)) = v547
	goto L118
L117:
	;
	goto L118
L118:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v432)+8))
	if v549 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v553 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+8)) = v553
	goto L121
L120:
	;
	goto L121
L121:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v432)+20))
	if v555 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v582)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+28)) = v583
	v587 = v580 + v581<<(uint(int32(2))%32)
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v588)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+80)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v432)+36)) = v589
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v587)))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v592)))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+4)) = v593
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+24)) = uint8(v595)
	goto L112
L123:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v555+v556<<(uint(int32(2))%32))))
	if v560 != 0 {
		v580 = v555
		v581 = v556
		v582 = v560
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	F_syncrep_yyensure_buffer_stack(m, v432)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L6
	} else {
		goto L127
	}
L126:
	;
	goto L125
L127:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	v566 = F_syncrep_yy_create_buffer(m, v565, v432)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v432)+20))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v570 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v568+v569<<(uint(v570)%32)))) = v566
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v432)+20))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v432)+12))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v574+v575<<(uint(v570)%32))))
	v580 = v574
	v581 = v575
	v582 = v579
	goto L122
L129:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v603)+36))
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v626))) = uint8(v627)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v603)+44))
	v635 = v629
	v636 = v626
	v638 = v626
	goto L131
L131:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638))))
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656)+uint32(_c_F_check_synchronous_standby_names[4]))))
	if base.Ui32(int32(-26)) <= base.Ui32(v635&int32(2147483647)-int32(31)) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+68)) = v638
	*(*int32)(unsafe.Add(mBase, uint32(v603)+64)) = v635
	goto L135
L134:
	;
	goto L135
L135:
	;
	v668 = int32(1)
	v672 = int32(*(*int16)(unsafe.Add(mBase, uint32(v635<<(uint(v668)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v673 = v672 + v659
	v678 = int32(*(*int16)(unsafe.Add(mBase, uint32(v673<<(uint(v668)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v678 != v635 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v685 = v635
	v687 = v659
	v689 = v659
	goto L139
L137:
	;
	v744 = v673
	goto L138
L138:
	;
	v766 = int32(1)
	v772 = int32(*(*int16)(unsafe.Add(mBase, uint32(v744<<(uint(v766)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v744&int32(2147483647)))%64)&int64(-32985348833280) == int64(0) {
		v635 = v772
		v638 = v638 + v766
		goto L131
	} else {
		goto L145
	}
L139:
	;
	v710 = int32(*(*int16)(unsafe.Add(mBase, uint32(v685<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v685&int32(2147483647)))%64)&int64(2076672024) != int64(0) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v744 = v733
	goto L138
L141:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687)+uint32(_c_F_check_synchronous_standby_names[9]))))
	v723 = v722
	goto L143
L142:
	;
	v723 = v689
	goto L143
L143:
	;
	v727 = v723 & int32(255)
	v728 = int32(1)
	v732 = int32(*(*int16)(unsafe.Add(mBase, uint32(v710<<(uint(v728)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v733 = v727 + v732
	v738 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v733<<(uint(v728)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v710&int32(_a_F_check_synchronous_standby_names_12) != v738 {
		v685 = v710
		v687 = v727
		v689 = v723
		goto L139
	} else {
		goto L144
	}
L144:
	;
	goto L140
L145:
	;
	v792 = v636
	goto L146
L146:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v603)+64))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v603)+68))
	v815 = v808
	v820 = v792
	v824 = v809
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+80)) = v820
	*(*int32)(unsafe.Add(mBase, uint32(v603)+32)) = v824 - v820
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v824))))
	*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)) = uint8(v839)
	v841 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v824))) = uint8(v841)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v824
	v848 = int32(*(*int16)(unsafe.Add(mBase, uint32(v815<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[10]))))
	v853 = v848
	goto L150
L150:
	;
	v875 = int32(260)
	switch v853 {
	case 0:
		goto L179
	case 1:
		goto L129
	case 2:
		goto L166
	case 3:
		goto L165
	case 4:
		goto L178
	case 5:
		goto L177
	case 6:
		goto L176
	case 7:
		goto L175
	case 8:
		goto L173
	case 9:
		goto L172
	case 10:
		goto L171
	case 11:
		goto L164
	case 12:
		goto L163
	case 13:
		goto L162
	case 14:
		v1744 = v875
		goto L161
	case 15:
		goto L170
	case 16:
		goto L168
	case 17:
		goto L169
	case 18:
		goto L174
	default:
		goto L167
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v2365
	*(*int32)(unsafe.Add(mBase, uint32(v603)+48)) = int32(0)
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v603)+44))
	v2392 = base.I32_div_s(v2388-int32(1), int32(2))
	v853 = v2392 + int32(17)
	goto L150
L153:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_13))
	mBase = m.M
	v2358 = m.ExcPending
	if v2358 != 0 {
		goto L6
	} else {
		goto L392
	}
L154:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_14))
	mBase = m.M
	v2353 = m.ExcPending
	if v2353 != 0 {
		goto L6
	} else {
		goto L391
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	v2199 = v2177 + v2189
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v2199
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(v603)+44))
	if base.Ui32(v2199) <= base.Ui32(v2179) {
		v815 = v2201
		v820 = v2179
		v824 = v2199
		goto L148
	} else {
		goto L372
	}
L157:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1811)))
	*(*int32)(unsafe.Add(mBase, uint32(v1812)+16)) = v1789
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	if v1815 != 0 {
		v1937 = int32(0)
		goto L316
	} else {
		goto L317
	}
L158:
	;
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1789 = v1758
	v1811 = v1780 + v1781<<(uint(int32(2))%32)
	goto L157
L159:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_15))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L6
	} else {
		goto L315
	}
L160:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_16))
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L6
	} else {
		goto L314
	}
L161:
	;
	m.G0 = v619 + int32(32)
	goto L109
L162:
	;
	v1744 = int32(41)
	goto L161
L163:
	;
	v1744 = int32(40)
	goto L161
L164:
	;
	v1744 = int32(44)
	goto L161
L165:
	;
	v1744 = int32(262)
	goto L161
L166:
	;
	v1744 = int32(261)
	goto L161
L167:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_17))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L6
	} else {
		goto L313
	}
L168:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v824))) = uint8(v940)
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v946 = v942 + v943<<(uint(int32(2))%32)
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v947)+44))
	if v948 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L169:
	;
	v1744 = int32(0)
	goto L161
L170:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_18))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L6
	} else {
		goto L191
	}
L171:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v603)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v931))) = int32(_a_F_check_synchronous_standby_names_19)
	v1744 = int32(258)
	goto L161
L172:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	v926 = F_pstrdup(m, v925)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L6
	} else {
		goto L190
	}
L173:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	v920 = F_pstrdup(m, v919)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L6
	} else {
		goto L189
	}
L174:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	if v901 != 0 {
		v1744 = v875
		goto L161
	} else {
		goto L183
	}
L175:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v603)+92))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v892)))
	*(*int32)(unsafe.Add(mBase, uint32(v891))) = v893
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	*(*int32)(unsafe.Add(mBase, uint32(v895))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+44)) = int32(1)
	v1744 = int32(258)
	goto L161
L176:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	F_appendStringInfoString(m, v887, v888)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L6
	} else {
		goto L182
	}
L177:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	F_appendStringInfoChar(m, v883, int32(34))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L6
	} else {
		goto L181
	}
L178:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	F_initStringInfo(m, v878)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L6
	} else {
		goto L180
	}
L179:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v824))) = uint8(v876)
	v792 = v820
	goto L146
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+44)) = int32(3)
	goto L129
L181:
	;
	goto L129
L182:
	;
	goto L129
L183:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902))))
	if v903 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v619)+20)) = v902
	*(*int32)(unsafe.Add(mBase, uint32(v619)+16)) = int32(_a_F_check_synchronous_standby_names_20)
	v910 = F_psprintf(m, int32(_a_F_check_synchronous_standby_names_21), v619+int32(16))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L6
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v619))) = int32(_a_F_check_synchronous_standby_names_20)
	v916 = F_psprintf(m, int32(_a_F_check_synchronous_standby_names_22), v619)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L6
	} else {
		goto L188
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v620))) = v910
	v1744 = v875
	goto L161
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v620))) = v916
	v1744 = v875
	goto L161
L189:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v603)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v922))) = v920
	v1744 = int32(258)
	goto L161
L190:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v603)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v928))) = v926
	v1744 = int32(259)
	goto L161
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v947)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v951
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v953))) = v954
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v958 = int32(2)
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v956+v957<<(uint(v958)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v961)+44)) = int32(1)
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v964+v965<<(uint(v958)%32))))
	v970 = v969
	v971 = v964
	v972 = v965
	goto L194
L193:
	;
	v970 = v947
	v971 = v942
	v972 = v943
	goto L194
L194:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v603)+36))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v970)+4))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	v976 = v974 + v975
	if base.Ui32(v973) <= base.Ui32(v976) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	v982 = v978 + (v939 ^ int32(-1)) + v824
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v982
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v603)+44))
	if base.Ui32(v978) < base.Ui32(v982) {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	goto L197
L197:
	;
	if base.Ui32(v976+int32(1)) < base.Ui32(v973) {
		goto L160
	} else {
		goto L230
	}
L198:
	;
	v991 = v984
	v994 = v978
	goto L201
L199:
	;
	v1139 = v984
	goto L200
L200:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v1139&int32(2147483647)-int32(31)) {
		goto L219
	} else {
		goto L220
	}
L201:
	;
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v994))))
	if v1012 != 0 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v1139 = v1130
	goto L200
L203:
	;
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1012)+uint32(_c_F_check_synchronous_standby_names[4]))))
	v1017 = v1015
	goto L205
L204:
	;
	v1017 = int32(1)
	goto L205
L205:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v991&int32(2147483647)-int32(31)) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+68)) = v994
	*(*int32)(unsafe.Add(mBase, uint32(v603)+64)) = v991
	goto L208
L207:
	;
	goto L208
L208:
	;
	v1027 = v1017 & int32(255)
	v1028 = int32(1)
	v1032 = int32(*(*int16)(unsafe.Add(mBase, uint32(v991<<(uint(v1028)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v1033 = v1027 + v1032
	v1038 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1033<<(uint(v1028)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v1038 != v991 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1045 = v991
	v1047 = v1017
	v1049 = v1027
	goto L212
L210:
	;
	v1104 = v1033
	goto L211
L211:
	;
	v1126 = int32(1)
	v1130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1104<<(uint(v1126)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v1132 = v994 + v1126
	if v1132 != v982 {
		v991 = v1130
		v994 = v1132
		goto L201
	} else {
		goto L218
	}
L212:
	;
	v1070 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1045<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v1045&int32(2147483647)))%64)&int64(2076672024) != int64(0) {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	v1104 = v1093
	goto L211
L214:
	;
	v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1049)+uint32(_c_F_check_synchronous_standby_names[9]))))
	v1083 = v1082
	goto L216
L215:
	;
	v1083 = v1047
	goto L216
L216:
	;
	v1087 = v1083 & int32(255)
	v1088 = int32(1)
	v1092 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1070<<(uint(v1088)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v1093 = v1087 + v1092
	v1098 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1093<<(uint(v1088)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v1070&int32(_a_F_check_synchronous_standby_names_12) != v1098 {
		v1045 = v1070
		v1047 = v1083
		v1049 = v1087
		goto L212
	} else {
		goto L217
	}
L217:
	;
	goto L213
L218:
	;
	goto L202
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+68)) = v982
	*(*int32)(unsafe.Add(mBase, uint32(v603)+64)) = v1139
	goto L221
L220:
	;
	goto L221
L221:
	;
	v1168 = int32(1)
	v1172 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1139<<(uint(v1168)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v1174 = v1172 + v1168
	v1179 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1174<<(uint(v1168)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v1179 != v1139 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1186 = v1139
	goto L225
L223:
	;
	v1230 = v1174
	goto L224
L224:
	;
	v1254 = v1230 & int32(2147483647)
	if int64(1)<<(uint(base.I64_extend_i32_u(v1254))%64)&int64(-32985348833280) != int64(0) {
		v792 = v978
		goto L146
	} else {
		goto L228
	}
L225:
	;
	v1207 = int32(1)
	v1211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1186<<(uint(v1207)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	v1212 = base.I32_extend16_s(v1211)
	v1217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1212<<(uint(v1207)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v1219 = v1217 + v1207
	v1224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1219<<(uint(v1207)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v1211 != v1224 {
		v1186 = v1212
		goto L225
	} else {
		goto L227
	}
L226:
	;
	v1230 = v1219
	goto L224
L227:
	;
	goto L226
L228:
	;
	if v1254 == int32(0) {
		v792 = v978
		goto L146
	} else {
		goto L229
	}
L229:
	;
	v1263 = int32(1)
	v1264 = v982 + v1263
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v1264
	v1270 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1230<<(uint(v1263)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v635 = v1270
	v636 = v978
	v638 = v1264
	goto L131
L230:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v603)+80))
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v970)+40))
	if v1275 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	if v973-v1274 != int32(1) {
		v2177 = v974
		v2179 = v1274
		v2189 = v975
		goto L156
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	v1283 = v1274 ^ int32(-1) + v973
	if v1283 != 0 {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v2365 = v1274
	goto L152
L235:
	;
	v1284 = int32(7)
	v1285 = v1283 & v1284
	if base.Ui32(v973-v1274-int32(2)) < base.Ui32(v1284) {
		goto L239
	} else {
		goto L240
	}
L236:
	;
	v1442 = v970
	v1444 = v971
	v1446 = v972
	goto L237
L237:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+44))
	if v1463 == int32(2) {
		goto L251
	} else {
		goto L252
	}
L238:
	;
	if v1285 != 0 {
		goto L245
	} else {
		goto L246
	}
L239:
	;
	v1347 = v974
	v1348 = v1274
	goto L238
L240:
	;
	goto L241
L241:
	;
	v1298 = v974
	v1299 = v1274
	v1301 = int32(0)
	goto L242
L242:
	;
	v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298))) = uint8(v1320)
	v1322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+1)) = uint8(v1322)
	v1324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+2)) = uint8(v1324)
	v1326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+3)) = uint8(v1326)
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+4)) = uint8(v1328)
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+5)) = uint8(v1330)
	v1332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+6)) = uint8(v1332)
	v1334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1298)+7)) = uint8(v1334)
	v1336 = int32(8)
	v1337 = v1298 + v1336
	v1339 = v1299 + v1336
	v1341 = v1301 + v1336
	if v1341 != v1283&int32(-8) {
		v1298 = v1337
		v1299 = v1339
		v1301 = v1341
		goto L242
	} else {
		goto L244
	}
L243:
	;
	v1347 = v1337
	v1348 = v1339
	goto L238
L244:
	;
	goto L243
L245:
	;
	v1374 = v1347
	v1375 = v1348
	v1377 = int32(0)
	goto L248
L246:
	;
	goto L247
L247:
	;
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1431+v1432<<(uint(int32(2))%32))))
	v1442 = v1436
	v1444 = v1431
	v1446 = v1432
	goto L237
L248:
	;
	v1396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1375))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1374))) = uint8(v1396)
	v1398 = int32(1)
	v1403 = v1377 + v1398
	if v1403 != v1285 {
		v1374 = v1374 + v1398
		v1375 = v1375 + v1398
		v1377 = v1403
		goto L248
	} else {
		goto L250
	}
L249:
	;
	goto L247
L250:
	;
	goto L249
L251:
	;
	v1466 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v1466
	v1789 = v1466
	v1811 = v1444 + v1446<<(uint(int32(2))%32)
	goto L157
L252:
	;
	goto L253
L253:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+12))
	v1473 = v1274 - v973
	v1474 = v1472 + v1473
	if v1474 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v603)+36))
	v1482 = v1472
	v1483 = v1442
	v1487 = v1477
	goto L257
L255:
	;
	v1545 = v1442
	v1547 = v1474
	goto L256
L256:
	;
	v1566 = int32(_a_F_check_synchronous_standby_names_23)
	if base.Ui32(v1566) <= base.Ui32(v1547) {
		goto L273
	} else {
		goto L274
	}
L257:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1483)+20))
	if v1504 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v1545 = v1535
	v1547 = v1537
	goto L256
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1483)+4)) = int32(0)
	goto L153
L260:
	;
	goto L261
L261:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1483)+4))
	if base.Ui32(int32(2147483646)) <= base.Ui32(v1482) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1515 = int32(-3)
	goto L264
L263:
	;
	v1515 = v1482 << (uint(int32(1)) % 32)
	goto L264
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1483)+12)) = v1515
	v1518 = v1515 + int32(2)
	if v1509 != 0 {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1483)+4)) = v1523
	if v1523 == int32(0) {
		goto L153
	} else {
		goto L271
	}
L266:
	;
	v1519 = F_repalloc(m, v1509, v1518)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L6
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v1521 = F_palloc(m, v1518)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L6
	} else {
		goto L270
	}
L269:
	;
	v1523 = v1519
	goto L265
L270:
	;
	v1523 = v1521
	goto L265
L271:
	;
	v1528 = v1523 + (v1487 - v1509)
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v1528
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1530+v1531<<(uint(int32(2))%32))))
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1535)+12))
	v1537 = v1536 + v1473
	if v1537 == int32(0) {
		v1482 = v1536
		v1483 = v1535
		v1487 = v1528
		goto L257
	} else {
		goto L272
	}
L272:
	;
	goto L258
L273:
	;
	v1569 = v1566
	goto L275
L274:
	;
	v1569 = v1547
	goto L275
L275:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1545)+24))
	if v1571 != 0 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1576 = int32(0)
	goto L280
L277:
	;
	goto L278
L278:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = int32(0)
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1651+v1652<<(uint(int32(2))%32))))
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1656)+4))
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1661 = F_fread(m, v1657+v1283, int32(1), v1569, v1660)
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L6
	} else {
		goto L295
	}
L279:
	;
	switch v1602 {
	case 0:
		goto L287
	default:
		v1646 = v1616
		goto L285
	case 11:
		goto L286
	}
L280:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1599 = F_do_getc(m, v1598)
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L6
	} else {
		goto L283
	}
L281:
	;
	v1616 = v1569
	goto L279
L282:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1603+v1604<<(uint(int32(2))%32))))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1608)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1609+v1283+v1576))) = uint8(v1599)
	v1614 = v1576 + int32(1)
	if v1614 != v1569 {
		v1576 = v1614
		goto L280
	} else {
		goto L284
	}
L283:
	;
	v1602 = v1599 + int32(1)
	switch v1602 {
	case 0, 11:
		v1616 = v1576
		goto L279
	default:
		goto L282
	}
L284:
	;
	goto L281
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v1646
	v1758 = v1646
	goto L158
L286:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1633+v1634<<(uint(int32(2))%32))))
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+4))
	v1642 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v1639+v1283+v1616))) = uint8(v1642)
	v1646 = v1616 + int32(1)
	goto L285
L287:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1617)+76))
	if v1618 < int32(0) {
		goto L290
	} else {
		goto L291
	}
L288:
	;
	if int32(base.Ui32(v1623)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v1646 = v1616
		goto L285
	} else {
		goto L293
	}
L289:
	;
	goto L288
L290:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1617)))
	v1623 = v1621
	goto L289
L291:
	;
	goto L292
L292:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1617)))
	v1623 = v1622
	goto L289
L293:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_15))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L6
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	v1667 = v1661
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v1667
	if v1667 != 0 {
		v1758 = v1667
		goto L158
	} else {
		goto L298
	}
L298:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+76))
	if v1691 < int32(0) {
		goto L301
	} else {
		goto L302
	}
L299:
	;
	if int32(base.Ui32(v1696)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L304
	} else {
		goto L305
	}
L300:
	;
	goto L299
L301:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1690)))
	v1696 = v1694
	goto L300
L302:
	;
	goto L303
L303:
	;
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1690)))
	v1696 = v1695
	goto L300
L304:
	;
	v1758 = int32(0)
	goto L158
L305:
	;
	goto L306
L306:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	if v1705 != int32(27) {
		goto L159
	} else {
		goto L307
	}
L307:
	;
	v1709 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = v1709
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1711)+76))
	if v1709 <= v1712 {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1723+v1724<<(uint(int32(2))%32))))
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1728)+4))
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1733 = F_fread(m, v1729+v1283, int32(1), v1569, v1732)
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L6
	} else {
		goto L312
	}
L309:
	;
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1711)))
	*(*int32)(unsafe.Add(mBase, uint32(v1711))) = v1715 & int32(-49)
	goto L308
L310:
	;
	goto L311
L311:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1711)))
	*(*int32)(unsafe.Add(mBase, uint32(v1711))) = v1719 & int32(-49)
	goto L308
L312:
	;
	v1667 = v1733
	goto L296
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L314:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L316:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	v1939 = v1938 + v1283
	v1940 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v1940+v1941<<(uint(int32(2))%32))))
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+12))
	if base.Ui32(v1946) < base.Ui32(v1939) {
		goto L340
	} else {
		goto L341
	}
L317:
	;
	if v1283 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	if v1819 != 0 {
		goto L323
	} else {
		goto L324
	}
L319:
	;
	goto L320
L320:
	;
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1925 = int32(2)
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v1923+v1924<<(uint(v1925)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1928)+44)) = v1925
	v1937 = v1925
	goto L316
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1885)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1885))) = v1818
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	if v1892 != 0 {
		goto L336
	} else {
		goto L337
	}
L322:
	;
	v1842 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1847 = *(*int32)(unsafe.Add(mBase, uint32(v1840+v1843<<(uint(int32(2))%32))))
	if v1847 == int32(0) {
		goto L330
	} else {
		goto L331
	}
L323:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1819+v1820<<(uint(int32(2))%32))))
	if v1824 != 0 {
		v1840 = v1819
		goto L322
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	F_syncrep_yyensure_buffer_stack(m, v603)
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L6
	} else {
		goto L327
	}
L326:
	;
	goto L325
L327:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v1828 = F_syncrep_yy_create_buffer(m, v1827, v603)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L6
	} else {
		goto L328
	}
L328:
	;
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1830+v1831<<(uint(int32(2))%32)))) = v1828
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	if v1836 != 0 {
		v1840 = v1836
		goto L322
	} else {
		goto L329
	}
L329:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	v1885 = int32(0)
	v1888 = v1838
	goto L321
L330:
	;
	v1885 = int32(0)
	v1888 = v1842
	goto L321
L331:
	;
	goto L332
L332:
	;
	v1851 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1847)+16)) = v1851
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1847)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1853))) = uint8(v1851)
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1847)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1856)+1)) = uint8(v1851)
	*(*int32)(unsafe.Add(mBase, uint32(v1847)+44)) = v1851
	*(*int32)(unsafe.Add(mBase, uint32(v1847)+28)) = int32(1)
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v1847)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1847)+8)) = v1863
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	if v1865 == v1851 {
		v1885 = v1847
		v1888 = v1842
		goto L321
	} else {
		goto L333
	}
L333:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1871 = v1865 + v1868<<(uint(int32(2))%32)
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1871)))
	if v1847 != v1872 {
		v1885 = v1847
		v1888 = v1842
		goto L321
	} else {
		goto L334
	}
L334:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1872)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v1874
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1871)))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1876)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+80)) = v1877
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v1877
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1871)))
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1880)))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+4)) = v1881
	v1883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1877))))
	*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)) = uint8(v1883)
	v1885 = v1847
	v1888 = v1842
	goto L321
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1885)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = v1888
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1909 = v1905 + v1906<<(uint(int32(2))%32)
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v1909)))
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1910)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v1911
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v1909)))
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v1913)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v1914
	*(*int32)(unsafe.Add(mBase, uint32(v603)+80)) = v1914
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1909)))
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v1917)))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+4)) = v1918
	v1920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1914))))
	*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)) = uint8(v1920)
	v1937 = int32(1)
	goto L316
L336:
	;
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1892+v1893<<(uint(int32(2))%32))))
	if v1885 == v1897 {
		goto L335
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1885)+32)) = int64(1)
	goto L335
L339:
	;
	goto L338
L340:
	;
	v1950 = v1939 + int32(base.Ui32(v1938)>>(uint(int32(1))%32))
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1945)+4))
	if v1951 != 0 {
		goto L344
	} else {
		goto L345
	}
L341:
	;
	v1980 = v1939
	v1981 = v1940
	v1982 = v1941
	goto L342
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+28)) = v1980
	v1984 = int32(2)
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v1981+v1982<<(uint(v1984)%32))))
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1987)+4))
	v1990 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1988+v1980))) = uint8(v1990)
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1997 = *(*int32)(unsafe.Add(mBase, uint32(v1992+v1993<<(uint(v1984)%32))))
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v1997)+4))
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1998+v1999)+1)) = uint8(v1990)
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v2007 = v2003 + v2004<<(uint(v1984)%32)
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v2007)))
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v2008)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+80)) = v2009
	if v1937 == int32(1) {
		v2365 = v2009
		goto L152
	} else {
		goto L350
	}
L343:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1959 = int32(2)
	v1962 = *(*int32)(unsafe.Add(mBase, uint32(v1957+v1958<<(uint(v1959)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1962)+4)) = v1956
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v1964+v1965<<(uint(v1959)%32))))
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v1969)+4))
	if v1970 == int32(0) {
		goto L154
	} else {
		goto L349
	}
L344:
	;
	v1952 = F_repalloc(m, v1951, v1950)
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L6
	} else {
		goto L347
	}
L345:
	;
	goto L346
L346:
	;
	v1954 = F_palloc(m, v1950)
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L6
	} else {
		goto L348
	}
L347:
	;
	v1956 = v1952
	goto L343
L348:
	;
	v1956 = v1954
	goto L343
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1969)+12)) = v1950 - int32(2)
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v603)+20))
	v1980 = v1977 + v1283
	v1981 = v1979
	v1982 = v1976
	goto L342
L350:
	;
	switch v1937 - int32(1) {
	case 0:
		goto L155
	case 1:
		goto L351
	default:
		goto L352
	}
L351:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v603)+28))
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v2007)))
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2171)+4))
	v2177 = v2172
	v2179 = v2009
	v2189 = v2170
	goto L156
L352:
	;
	v2018 = v2009 + (v939 ^ int32(-1)) + v824
	*(*int32)(unsafe.Add(mBase, uint32(v603)+36)) = v2018
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v603)+44))
	if base.Ui32(v2018) <= base.Ui32(v2009) {
		v635 = v2020
		v636 = v2009
		v638 = v2018
		goto L131
	} else {
		goto L353
	}
L353:
	;
	v2027 = v2020
	v2036 = v2009
	goto L354
L354:
	;
	v2048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2036))))
	if v2048 != 0 {
		goto L356
	} else {
		goto L357
	}
L355:
	;
	v635 = v2166
	v636 = v2009
	v638 = v2018
	goto L131
L356:
	;
	v2051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2048)+uint32(_c_F_check_synchronous_standby_names[4]))))
	v2053 = v2051
	goto L358
L357:
	;
	v2053 = int32(1)
	goto L358
L358:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v2027&int32(2147483647)-int32(31)) {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+68)) = v2036
	*(*int32)(unsafe.Add(mBase, uint32(v603)+64)) = v2027
	goto L361
L360:
	;
	goto L361
L361:
	;
	v2063 = v2053 & int32(255)
	v2064 = int32(1)
	v2068 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2027<<(uint(v2064)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v2069 = v2063 + v2068
	v2074 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2069<<(uint(v2064)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v2074 != v2027 {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v2081 = v2027
	v2083 = v2053
	v2085 = v2063
	goto L365
L363:
	;
	v2140 = v2069
	goto L364
L364:
	;
	v2162 = int32(1)
	v2166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2140<<(uint(v2162)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v2168 = v2036 + v2162
	if v2018 != v2168 {
		v2027 = v2166
		v2036 = v2168
		goto L354
	} else {
		goto L371
	}
L365:
	;
	v2106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2081<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v2081&int32(2147483647)))%64)&int64(2076672024) != int64(0) {
		goto L367
	} else {
		goto L368
	}
L366:
	;
	v2140 = v2129
	goto L364
L367:
	;
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2085)+uint32(_c_F_check_synchronous_standby_names[9]))))
	v2119 = v2118
	goto L369
L368:
	;
	v2119 = v2083
	goto L369
L369:
	;
	v2123 = v2119 & int32(255)
	v2124 = int32(1)
	v2128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2106<<(uint(v2124)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v2129 = v2123 + v2128
	v2134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2129<<(uint(v2124)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v2106&int32(_a_F_check_synchronous_standby_names_12) != v2134 {
		v2081 = v2106
		v2083 = v2119
		v2085 = v2123
		goto L365
	} else {
		goto L370
	}
L370:
	;
	goto L366
L371:
	;
	goto L355
L372:
	;
	v2208 = v2201
	v2211 = v2179
	goto L373
L373:
	;
	v2229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2211))))
	if v2229 != 0 {
		goto L375
	} else {
		goto L376
	}
L374:
	;
	v815 = v2347
	v820 = v2179
	v824 = v2199
	goto L148
L375:
	;
	v2232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2229)+uint32(_c_F_check_synchronous_standby_names[4]))))
	v2234 = v2232
	goto L377
L376:
	;
	v2234 = int32(1)
	goto L377
L377:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v2208&int32(2147483647)-int32(31)) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v603)+68)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v603)+64)) = v2208
	goto L380
L379:
	;
	goto L380
L380:
	;
	v2244 = v2234 & int32(255)
	v2245 = int32(1)
	v2249 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2208<<(uint(v2245)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v2250 = v2244 + v2249
	v2255 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2250<<(uint(v2245)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v2255 != v2208 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v2262 = v2208
	v2264 = v2234
	v2266 = v2244
	goto L384
L382:
	;
	v2321 = v2250
	goto L383
L383:
	;
	v2343 = int32(1)
	v2347 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2321<<(uint(v2343)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v2349 = v2211 + v2343
	if v2349 != v2199 {
		v2208 = v2347
		v2211 = v2349
		goto L373
	} else {
		goto L390
	}
L384:
	;
	v2287 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2262<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v2262&int32(2147483647)))%64)&int64(2076672024) != int64(0) {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	v2321 = v2310
	goto L383
L386:
	;
	v2299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2266)+uint32(_c_F_check_synchronous_standby_names[9]))))
	v2300 = v2299
	goto L388
L387:
	;
	v2300 = v2264
	goto L388
L388:
	;
	v2304 = v2300 & int32(255)
	v2305 = int32(1)
	v2309 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2287<<(uint(v2305)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v2310 = v2304 + v2309
	v2315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2310<<(uint(v2305)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v2287&int32(_a_F_check_synchronous_standby_names_12) != v2315 {
		v2262 = v2287
		v2264 = v2300
		v2266 = v2304
		goto L384
	} else {
		goto L389
	}
L389:
	;
	goto L385
L390:
	;
	goto L374
L391:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L392:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L393:
	;
	v2423 = int32(0)
	v2431 = v2423
	v2432 = v2423
	goto L105
L394:
	;
	goto L395
L395:
	;
	if base.Ui32(int32(262)) < base.Ui32(v2400) {
		v2431 = v2400
		v2432 = int32(2)
		goto L105
	} else {
		goto L396
	}
L396:
	;
	v2430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2400)+uint32(_c_F_check_synchronous_standby_names[11]))))
	v2431 = v2400
	v2432 = v2430
	goto L105
L397:
	;
	v2438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2433)+uint32(_c_F_check_synchronous_standby_names[12]))))
	if v2432 != v2438 {
		v2453 = v2395
		v2454 = v2396
		v2455 = v2397
		v2456 = v2398
		v2458 = v2431
		v2464 = v2406
		v2465 = v2407
		v2466 = v2408
		v2468 = v2410
		v2470 = v2412
		v2471 = v2413
		v2473 = v2415
		v2474 = v2416
		v2478 = v2420
		goto L103
	} else {
		goto L398
	}
L398:
	;
	if v2433 != int32(19) {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v2443 = v2406 + int32(4)
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v2412)+1228))
	*(*int32)(unsafe.Add(mBase, uint32(v2443))) = v2444
	if v2431 != 0 {
		goto L402
	} else {
		goto L403
	}
L400:
	;
	goto L401
L401:
	;
	v2644 = v2395
	v2645 = v2396
	v2646 = v2397
	v2659 = int32(0)
	v2661 = v2412
	goto L73
L402:
	;
	v2448 = int32(-2)
	goto L404
L403:
	;
	v2448 = int32(0)
	goto L404
L404:
	;
	v2451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2433)+uint32(_c_F_check_synchronous_standby_names[13]))))
	v2574 = v2395
	v2575 = v2396
	v2576 = v2397
	v2577 = v2398
	v2579 = v2448
	v2585 = v2443
	v2587 = v2408
	v2589 = v2410
	v2591 = v2412
	v2594 = v2415
	v2595 = v2416
	v2599 = v2420
	v2600 = v2451
	goto L102
L405:
	;
	v2484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2471)+uint32(_c_F_check_synchronous_standby_names[14]))))
	v2487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2484)+uint32(_c_F_check_synchronous_standby_names[15]))))
	v2489 = int32(2)
	v2492 = *(*int32)(unsafe.Add(mBase, uint32(v2464+(int32(1)-v2487)<<(uint(v2489)%32))))
	switch v2484 - v2489 {
	case 0:
		goto L414
	case 1:
		goto L413
	case 2:
		goto L412
	case 3:
		goto L411
	case 4:
		goto L410
	case 5:
		goto L409
	case 6:
		goto L408
	case 7, 8:
		goto L407
	default:
		v2544 = v2492
		goto L406
	}
L406:
	;
	v2549 = v2464 - v2487<<(uint(int32(2))%32) + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2549))) = v2544
	v2553 = v2466 - v2487<<(uint(int32(1))%32)
	v2554 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2553))))
	v2557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2484)+uint32(_c_F_check_synchronous_standby_names[16]))))
	v2560 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2557)+uint32(_c_F_check_synchronous_standby_names[16]))))
	v2561 = v2554 + v2560
	if base.Ui32(int32(22)) < base.Ui32(v2561) {
		goto L421
	} else {
		goto L422
	}
L407:
	;
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(v2464)))
	v2544 = v2543
	goto L406
L408:
	;
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(8))))
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(v2464)))
	v2541 = F_lappend(m, v2539, v2540)
	mBase = m.M
	v2542 = m.ExcPending
	if v2542 != 0 {
		goto L6
	} else {
		goto L420
	}
L409:
	;
	v2529 = *(*int32)(unsafe.Add(mBase, uint32(v2464)))
	*(*int32)(unsafe.Add(mBase, uint32(v2470)+8)) = v2529
	*(*int32)(unsafe.Add(mBase, uint32(v2470)+12)) = v2529
	v2535 = F_list_make1_impl(m, int32(1), v2470+int32(8))
	mBase = m.M
	v2536 = m.ExcPending
	if v2536 != 0 {
		goto L6
	} else {
		goto L419
	}
L410:
	;
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(12))))
	v2525 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(4))))
	v2527 = F_create_syncrep_config(m, v2522, v2525, int32(0))
	mBase = m.M
	v2528 = m.ExcPending
	if v2528 != 0 {
		goto L6
	} else {
		goto L418
	}
L411:
	;
	v2513 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(12))))
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(4))))
	v2518 = F_create_syncrep_config(m, v2513, v2516, int32(1))
	mBase = m.M
	v2519 = m.ExcPending
	if v2519 != 0 {
		goto L6
	} else {
		goto L417
	}
L412:
	;
	v2504 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(12))))
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v2464-int32(4))))
	v2509 = F_create_syncrep_config(m, v2504, v2507, int32(0))
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L6
	} else {
		goto L416
	}
L413:
	;
	v2498 = *(*int32)(unsafe.Add(mBase, uint32(v2464)))
	v2500 = F_create_syncrep_config(m, int32(_a_F_check_synchronous_standby_names_24), v2498, int32(0))
	mBase = m.M
	v2501 = m.ExcPending
	if v2501 != 0 {
		goto L6
	} else {
		goto L415
	}
L414:
	;
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(v2464)))
	*(*int32)(unsafe.Add(mBase, uint32(v2478))) = v2495
	v2544 = v2492
	goto L406
L415:
	;
	v2544 = v2500
	goto L406
L416:
	;
	v2544 = v2509
	goto L406
L417:
	;
	v2544 = v2518
	goto L406
L418:
	;
	v2544 = v2527
	goto L406
L419:
	;
	v2544 = v2535
	goto L406
L420:
	;
	v2544 = v2541
	goto L406
L421:
	;
	v2573 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2557)+uint32(_c_F_check_synchronous_standby_names[17]))))
	v2574 = v2453
	v2575 = v2454
	v2576 = v2455
	v2577 = v2456
	v2579 = v2458
	v2585 = v2549
	v2587 = v2553
	v2589 = v2468
	v2591 = v2470
	v2594 = v2473
	v2595 = v2474
	v2599 = v2478
	v2600 = v2573
	goto L102
L422:
	;
	v2566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2561)+uint32(_c_F_check_synchronous_standby_names[12]))))
	if v2566 != v2554 {
		goto L421
	} else {
		goto L423
	}
L423:
	;
	v2570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2561)+uint32(_c_F_check_synchronous_standby_names[13]))))
	v2574 = v2453
	v2575 = v2454
	v2576 = v2455
	v2577 = v2456
	v2579 = v2458
	v2585 = v2549
	v2587 = v2553
	v2589 = v2468
	v2591 = v2470
	v2594 = v2473
	v2595 = v2474
	v2599 = v2478
	v2600 = v2570
	goto L102
L424:
	;
	v2619 = v2466
	goto L425
L425:
	;
	if base.B2i32(v2453 == v2619) == int32(0) {
		v2619 = v2619 - int32(2)
		goto L425
	} else {
		goto L427
	}
L426:
	;
	v2644 = v2453
	v2645 = v2454
	v2646 = v2455
	v2659 = int32(1)
	v2661 = v2470
	goto L73
L427:
	;
	goto L426
L428:
	;
	v2644 = v429
	v2645 = v430
	v2646 = v431
	v2659 = int32(2)
	v2661 = v446
	goto L73
L429:
	;
	F_pfree(m, v2644)
	mBase = m.M
	v2674 = m.ExcPending
	if v2674 != 0 {
		goto L6
	} else {
		goto L432
	}
L430:
	;
	goto L431
L431:
	;
	m.G0 = v2661 + int32(1232)
	v2678 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+44))
	F_replication_scanner_finish(m, v2678)
	mBase = m.M
	v2680 = m.ExcPending
	if v2680 != 0 {
		goto L6
	} else {
		goto L433
	}
L432:
	;
	goto L431
L433:
	;
	if v2659 == int32(0) {
		goto L435
	} else {
		goto L436
	}
L434:
	;
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(v2683)+4))
	if v2709 <= int32(0) {
		goto L447
	} else {
		goto L448
	}
L435:
	;
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+40))
	if v2683 != 0 {
		goto L434
	} else {
		goto L438
	}
L436:
	;
	goto L437
L437:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[18])) = int32(16801924)
	goto L439
L438:
	;
	goto L437
L439:
	;
	v2688 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+36))
	v2690 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[19])) = v2690
	goto L440
L440:
	;
	if v2688 != 0 {
		goto L442
	} else {
		goto L443
	}
L441:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[20])) = v2706
	v2769 = v2646
	v2793 = int32(0)
	goto L1
L442:
	;
	v2694 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v2646)+16)) = v2694
	v2699 = F_format_elog_string(m, int32(_a_F_check_synchronous_standby_names_25), v2646+int32(16))
	mBase = m.M
	v2700 = m.ExcPending
	if v2700 != 0 {
		goto L6
	} else {
		goto L445
	}
L443:
	;
	goto L444
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2646))) = int32(_a_F_check_synchronous_standby_names_26)
	v2704 = F_format_elog_string(m, int32(_a_F_check_synchronous_standby_names_27), v2646)
	mBase = m.M
	v2705 = m.ExcPending
	if v2705 != 0 {
		goto L6
	} else {
		goto L446
	}
L445:
	;
	v2706 = v2699
	goto L441
L446:
	;
	v2706 = v2704
	goto L441
L447:
	;
	v2713 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[19])) = v2713
	goto L450
L448:
	;
	goto L449
L449:
	;
	v2728 = *(*int32)(unsafe.Add(mBase, uint32(v2683)))
	v2729 = F_guc_malloc(m, v2728)
	mBase = m.M
	v2730 = m.ExcPending
	if v2730 != 0 {
		goto L6
	} else {
		goto L452
	}
L450:
	;
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+40))
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2646)+32)) = v2717
	v2723 = F_format_elog_string(m, int32(_a_F_check_synchronous_standby_names_28), v2646+int32(32))
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L6
	} else {
		goto L451
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[21])) = v2723
	v2769 = v2646
	v2793 = int32(0)
	goto L1
L452:
	;
	if v2729 == int32(0) {
		v2769 = v2646
		v2793 = int32(0)
		goto L1
	} else {
		goto L453
	}
L453:
	;
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(v2646)+40))
	v2734 = *(*int32)(unsafe.Add(mBase, uint32(v2733)))
	if v2734 != 0 {
		goto L455
	} else {
		goto L456
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2645))) = v2736
	v2742 = v2646
	goto L2
L455:
	;
	v2735 = F__emscripten_memcpy_bulkmem(m, v2729, v2733, v2734)
	mBase = m.M
	v2736 = v2735
	goto L457
L456:
	;
	v2736 = v2729
	goto L457
L457:
	;
	goto L454
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
func F_chr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_chr[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(0) <= v11 {
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v217 = m.ExcPending
			if v217 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v220 = m.ExcPending
				if v220 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_chr_0), int32(0))
					mBase = m.M
					v224 = m.ExcPending
					if v224 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_chr_1), int32(1049), int32(_a_F_chr_2))
						mBase = m.M
						v229 = m.ExcPending
						if v229 != 0 {
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
			if v14 != int32(6) {
				if base.Ui32(v14) <= base.Ui32(int32(41)) {
					v177 = *(*int32)(unsafe.Add(mBase, uint32(v14*int32(28))+uint32(_c_F_chr[1])))
					v178 = v177
				} else {
					v178 = int32(1)
				}
				if int32(1) < v178 {
					v181 = base.B2i32(base.Ui32(v11) < base.Ui32(int32(128)))
				} else {
					v181 = base.B2i32(base.Ui32(v11) < base.Ui32(int32(256)))
				}
				if v181 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v249 = m.ExcPending
					if v249 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v11
							F_errmsg(m, int32(_a_F_chr_3), v9+int32(32))
							mBase = m.M
							v258 = m.ExcPending
							if v258 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_chr_1), int32(1121), int32(_a_F_chr_2))
								mBase = m.M
								v263 = m.ExcPending
								if v263 != 0 {
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
					v185 = F_palloc(m, int32(5))
					mBase = m.M
					v186 = m.ExcPending
					if v186 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v185)+4)) = uint8(v11)
						*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(20)
						v190 = v185
						m.G0 = v9 + int32(48)
						return v190
					}
				}
			} else {
				if base.Ui32(v11) < base.Ui32(int32(128)) {
					if base.Ui32(v14) <= base.Ui32(int32(41)) {
						v177 = *(*int32)(unsafe.Add(mBase, uint32(v14*int32(28))+uint32(_c_F_chr[1])))
						v178 = v177
					} else {
						v178 = int32(1)
					}
					if int32(1) < v178 {
						v181 = base.B2i32(base.Ui32(v11) < base.Ui32(int32(128)))
					} else {
						v181 = base.B2i32(base.Ui32(v11) < base.Ui32(int32(256)))
					}
					if v181 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v249 = m.ExcPending
						if v249 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v252 = m.ExcPending
							if v252 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v11
								F_errmsg(m, int32(_a_F_chr_3), v9+int32(32))
								mBase = m.M
								v258 = m.ExcPending
								if v258 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_chr_1), int32(1121), int32(_a_F_chr_2))
									mBase = m.M
									v263 = m.ExcPending
									if v263 != 0 {
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
						v185 = F_palloc(m, int32(5))
						mBase = m.M
						v186 = m.ExcPending
						if v186 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v185)+4)) = uint8(v11)
							*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(20)
							v190 = v185
							m.G0 = v9 + int32(48)
							return v190
						}
					}
				} else {
					if base.Ui32(int32(_a_F_chr_4)) <= base.Ui32(v11) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v233 = m.ExcPending
						if v233 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v236 = m.ExcPending
							if v236 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
								F_errmsg(m, int32(_a_F_chr_3), v9)
								mBase = m.M
								v240 = m.ExcPending
								if v240 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_chr_1), int32(1068), int32(_a_F_chr_2))
									mBase = m.M
									v245 = m.ExcPending
									if v245 != 0 {
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
						v29 = base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v11))
						if base.Ui32(int32(2047)) < base.Ui32(v11) {
							v30 = int32(3)
						} else {
							v30 = int32(2)
						}
						if base.Ui32(int32(_a_F_chr_5)) < base.Ui32(v11) {
							v33 = int32(4)
						} else {
							v33 = v30
						}
						v35 = v33 + int32(4)
						v36 = F_palloc(m, v35)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v36))) = v35 << (uint(int32(2)) % 32)
							v44 = v36 + int32(4)
							if v29 == int32(0) {
								v50 = int32(base.Ui32(v11)>>(uint(int32(6))%32)) | int32(192)
								*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v50)
								v90 = int32(5)
							} else {
								if base.Ui32(v11-int32(2048)) <= base.Ui32(int32(_a_F_chr_6)) {
									v60 = int32(base.Ui32(v11)>>(uint(int32(12))%32)) | int32(224)
									*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)) = uint8(v60)
									v62 = int32(6)
									v67 = int32(base.Ui32(v11)>>(uint(v62)%32))&int32(63) | int32(128)
									*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)) = uint8(v67)
									v90 = v62
								} else {
									v73 = int32(base.Ui32(v11)>>(uint(int32(18))%32)) | int32(240)
									*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)) = uint8(v73)
									v77 = int32(63)
									v79 = int32(128)
									v80 = int32(base.Ui32(v11)>>(uint(int32(6))%32))&v77 | v79
									*(*uint8)(unsafe.Add(mBase, uint32(v36)+6)) = uint8(v80)
									v87 = int32(base.Ui32(v11)>>(uint(int32(12))%32))&v77 | v79
									*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)) = uint8(v87)
									v90 = int32(7)
								}
							}
							v95 = v11&int32(63) | int32(128)
							*(*uint8)(unsafe.Add(mBase, uint32(v90+v36))) = uint8(v95)
							v97 = int32(0)
							switch v33 - int32(1) {
							case 0:
								v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
								v133 = v132
								if base.I32_extend8_s(v133) < int32(-62) {
									v146 = v97
								} else {
									v138 = v133
									v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
								}
							case 1:
								v106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+1)))
								v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
								switch v107 - int32(224) {
								case 0:
									v110 = int32(224)
									if base.Ui32(v110) <= base.Ui32((v106-int32(-64))&int32(255)) {
										v138 = v110
										v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
									} else {
										v146 = v97
									}
								default:
									if v106 <= int32(-65) {
										v133 = v107
										if base.I32_extend8_s(v133) < int32(-62) {
											v146 = v97
										} else {
											v138 = v133
											v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
										}
									} else {
										v146 = v97
									}
								case 13:
									if int32(-97) < v106 {
										v146 = v97
									} else {
										v138 = int32(237)
										v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
									}
								case 16:
									if base.Ui32((v106-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
										v146 = v97
									} else {
										v138 = int32(240)
										v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
									}
								case 20:
									if int32(-113) < v106 {
										v146 = v97
									} else {
										v138 = int32(244)
										v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
									}
								}
							case 2:
								v103 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+2)))
								if int32(-65) < v103 {
									v146 = v97
								} else {
									v106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+1)))
									v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
									switch v107 - int32(224) {
									case 0:
										v110 = int32(224)
										if base.Ui32(v110) <= base.Ui32((v106-int32(-64))&int32(255)) {
											v138 = v110
											v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
										} else {
											v146 = v97
										}
									default:
										if v106 <= int32(-65) {
											v133 = v107
											if base.I32_extend8_s(v133) < int32(-62) {
												v146 = v97
											} else {
												v138 = v133
												v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
											}
										} else {
											v146 = v97
										}
									case 13:
										if int32(-97) < v106 {
											v146 = v97
										} else {
											v138 = int32(237)
											v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
										}
									case 16:
										if base.Ui32((v106-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
											v146 = v97
										} else {
											v138 = int32(240)
											v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
										}
									case 20:
										if int32(-113) < v106 {
											v146 = v97
										} else {
											v138 = int32(244)
											v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
										}
									}
								}
							case 3:
								v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+3)))
								if int32(-65) < v100 {
									v146 = v97
								} else {
									v103 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+2)))
									if int32(-65) < v103 {
										v146 = v97
									} else {
										v106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+1)))
										v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
										switch v107 - int32(224) {
										case 0:
											v110 = int32(224)
											if base.Ui32(v110) <= base.Ui32((v106-int32(-64))&int32(255)) {
												v138 = v110
												v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
											} else {
												v146 = v97
											}
										default:
											if v106 <= int32(-65) {
												v133 = v107
												if base.I32_extend8_s(v133) < int32(-62) {
													v146 = v97
												} else {
													v138 = v133
													v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
												}
											} else {
												v146 = v97
											}
										case 13:
											if int32(-97) < v106 {
												v146 = v97
											} else {
												v138 = int32(237)
												v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
											}
										case 16:
											if base.Ui32((v106-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
												v146 = v97
											} else {
												v138 = int32(240)
												v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
											}
										case 20:
											if int32(-113) < v106 {
												v146 = v97
											} else {
												v138 = int32(244)
												v146 = base.B2i32(base.Ui32(v138&int32(255)) < base.Ui32(int32(245)))
											}
										}
									}
								}
							default:
								v146 = v97
							}
							if v146 != 0 {
								v190 = v36
								m.G0 = v9 + int32(48)
								return v190
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(261))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v11
										F_errmsg(m, int32(_a_F_chr_7), v9+int32(16))
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_chr_1), int32(1109), int32(_a_F_chr_2))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
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
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v201 = m.ExcPending
		if v201 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v204 = m.ExcPending
			if v204 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_chr_8), int32(0))
				mBase = m.M
				v208 = m.ExcPending
				if v208 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_chr_1), int32(1045), int32(_a_F_chr_2))
					mBase = m.M
					v213 = m.ExcPending
					if v213 != 0 {
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
