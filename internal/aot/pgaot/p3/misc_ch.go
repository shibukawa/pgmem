package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
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
func F_CheckElement_2(m *base.Module, l0 float32) {
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	if base.Ui32(base.I32_reinterpret_f32(l0)&int32(2147483647)) < base.Ui32(int32(2139095041)) {
		if base.F32_eq(base.F32_abs(l0), math.Float32frombits(uint32(0x7f800000))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(130))
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_CheckElement_2_0), int32(0))
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CheckElement_2_1), int32(147), int32(_a_F_CheckElement_2_2))
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_CheckElement_2_3), int32(0))
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CheckElement_2_1), int32(142), int32(_a_F_CheckElement_2_2))
					v25 = m.ExcPending
					if v25 != 0 {
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
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
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
	return v311
L2:
	;
	v311 = int32(1)
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
	v300 = m.ExcPending
	if v300 != 0 {
		goto L6
	} else {
		goto L72
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
	v297 = v4
	goto L5
L14:
	;
	v297 = int32(1)
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
		v297 = v4
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
	v297 = v4
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
	v156 = F_guc_malloc(m, v150)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L40
	}
L34:
	;
	v150 = int32(4)
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
	v140 = F_strlen(m, v139)
	mBase = m.M
	v142 = int32(1)
	v143 = v140 + v130 + v142
	v145 = v127 + v142
	if v145 != v120 {
		v127 = v145
		v130 = v143
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v150 = v143
	goto L33
L39:
	;
	goto L38
L40:
	;
	if v156 == int32(0) {
		v311 = v4
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v160 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v163 = v161
	goto L44
L43:
	;
	v163 = int32(0)
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v165 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v156
	goto L14
L46:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v168 <= int32(0) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v174 = int32(0)
	v178 = v156 + int32(4)
	goto L48
L48:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183+v174<<(uint(int32(2))%32))))
	if (v187^v178)&int32(3) != 0 {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	goto L45
L50:
	;
	v262 = F_strlen(m, v187)
	mBase = m.M
	v264 = int32(1)
	v267 = v174 + v264
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v267 < v268 {
		v174 = v267
		v178 = v178 + v262 + v264
		goto L48
	} else {
		goto L71
	}
L51:
	;
	goto L50
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v242))) = uint8(v241)
	if v241&int32(255) == int32(0) {
		goto L51
	} else {
		goto L67
	}
L53:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v240 = v187
	v241 = v193
	v242 = v178
	goto L52
L54:
	;
	goto L55
L55:
	;
	if v187&int32(3) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v197 = v187
	v199 = v178
	goto L59
L57:
	;
	v211 = v187
	v213 = v178
	goto L58
L58:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v218 = int32(-2139062144)
	if (int32(16843008)-v215|v215)&v218 != v218 {
		v240 = v211
		v241 = v215
		v242 = v213
		goto L52
	} else {
		goto L63
	}
L59:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v200)
	if v200 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L60:
	;
	v211 = v207
	v213 = v205
	goto L58
L61:
	;
	v204 = int32(1)
	v205 = v199 + v204
	v207 = v197 + v204
	if v207&int32(3) != 0 {
		v197 = v207
		v199 = v205
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v223 = v211
	v224 = v215
	v225 = v213
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = v224
	v227 = int32(4)
	v228 = v225 + v227
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v231 = v223 + v227
	v235 = int32(-2139062144)
	if (v229|(int32(16843008)-v229))&v235 == v235 {
		v223 = v231
		v224 = v229
		v225 = v228
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v240 = v231
	v241 = v229
	v242 = v228
	goto L52
L66:
	;
	goto L65
L67:
	;
	v249 = v240
	v251 = v242
	goto L68
L68:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v251)+1)) = uint8(v252)
	v254 = int32(1)
	if v252 != 0 {
		v249 = v249 + v254
		v251 = v251 + v254
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
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	F_list_free(m, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	v311 = v297
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
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
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
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v654 int32
	_ = v654
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v736 int32
	_ = v736
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v1014 int32
	_ = v1014
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1083 int32
	_ = v1083
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1130 int32
	_ = v1130
	var v1151 int32
	_ = v1151
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1174 int32
	_ = v1174
	var v1198 int32
	_ = v1198
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1347 int32
	_ = v1347
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1380 int32
	_ = v1380
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1520 int32
	_ = v1520
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1590 int32
	_ = v1590
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1611 int32
	_ = v1611
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1649 int32
	_ = v1649
	var v1653 int32
	_ = v1653
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1688 int32
	_ = v1688
	var v1694 int32
	_ = v1694
	var v1697 int32
	_ = v1697
	var v1702 int32
	_ = v1702
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1733 int32
	_ = v1733
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1759 int32
	_ = v1759
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1791 int32
	_ = v1791
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1800 int32
	_ = v1800
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1829 int32
	_ = v1829
	var v1832 int32
	_ = v1832
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1872 int32
	_ = v1872
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1906 int32
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1971 int32
	_ = v1971
	var v1980 int32
	_ = v1980
	var v1992 int32
	_ = v1992
	var v1995 int32
	_ = v1995
	var v1997 int32
	_ = v1997
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2018 int32
	_ = v2018
	var v2025 int32
	_ = v2025
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2050 int32
	_ = v2050
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2078 int32
	_ = v2078
	var v2084 int32
	_ = v2084
	var v2106 int32
	_ = v2106
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2133 int32
	_ = v2133
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2152 int32
	_ = v2152
	var v2155 int32
	_ = v2155
	var v2173 int32
	_ = v2173
	var v2176 int32
	_ = v2176
	var v2178 int32
	_ = v2178
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2199 int32
	_ = v2199
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2210 int32
	_ = v2210
	var v2231 int32
	_ = v2231
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2253 int32
	_ = v2253
	var v2254 int32
	_ = v2254
	var v2259 int32
	_ = v2259
	var v2265 int32
	_ = v2265
	var v2287 int32
	_ = v2287
	var v2291 int32
	_ = v2291
	var v2293 int32
	_ = v2293
	var v2297 int32
	_ = v2297
	var v2302 int32
	_ = v2302
	var v2309 int32
	_ = v2309
	var v2332 int32
	_ = v2332
	var v2336 int32
	_ = v2336
	var v2339 int32
	_ = v2339
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2364 int32
	_ = v2364
	var v2367 int32
	_ = v2367
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2382 int32
	_ = v2382
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2392 int32
	_ = v2392
	var v2395 int32
	_ = v2395
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2402 int32
	_ = v2402
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2412 int32
	_ = v2412
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2422 int32
	_ = v2422
	var v2428 int32
	_ = v2428
	var v2431 int32
	_ = v2431
	var v2433 int32
	_ = v2433
	var v2436 int32
	_ = v2436
	var v2439 int32
	_ = v2439
	var v2442 int32
	_ = v2442
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2448 int32
	_ = v2448
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2457 int32
	_ = v2457
	var v2460 int32
	_ = v2460
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2466 int32
	_ = v2466
	var v2469 int32
	_ = v2469
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2493 int32
	_ = v2493
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
	var v2510 int32
	_ = v2510
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
	var v2523 int32
	_ = v2523
	var v2529 int32
	_ = v2529
	var v2531 int32
	_ = v2531
	var v2533 int32
	_ = v2533
	var v2535 int32
	_ = v2535
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2549 int32
	_ = v2549
	var v2563 int32
	_ = v2563
	var v2586 int32
	_ = v2586
	var v2588 int32
	_ = v2588
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2603 int32
	_ = v2603
	var v2605 int32
	_ = v2605
	var v2618 int32
	_ = v2618
	var v2622 int32
	_ = v2622
	var v2624 int32
	_ = v2624
	var v2627 int32
	_ = v2627
	var v2632 int32
	_ = v2632
	var v2634 int32
	_ = v2634
	var v2638 int32
	_ = v2638
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2653 int32
	_ = v2653
	var v2657 int32
	_ = v2657
	var v2660 int32
	_ = v2660
	var v2661 int32
	_ = v2661
	var v2667 int32
	_ = v2667
	var v2668 int32
	_ = v2668
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2686 int32
	_ = v2686
	var v2713 int32
	_ = v2713
	var v2737 int32
	_ = v2737
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
	m.G0 = v2713 + int32(48)
	return v2737
L2:
	;
	v2713 = v2686
	v2737 = int32(1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v2686 = v29
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
	v92 = F_strlen(m, v31)
	mBase = m.M
	if base.Ui32(v92) < base.Ui32(int32(-2)) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v363 = m.G0
	v365 = v363 - int32(1232)
	m.G0 = v365
	v369 = v365 + int32(816)
	v371 = v365 + int32(16)
	v373 = v369
	v374 = l1
	v375 = v29
	v376 = v362
	v378 = int32(-2)
	v384 = v371
	v386 = v369
	v388 = v371
	v390 = v365
	v391 = v4
	v393 = v29 + int32(36)
	v394 = int32(200)
	v398 = v29 + int32(40)
	goto L58
L19:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_3))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L55
	}
L20:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_4))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L54
	}
L21:
	;
	v97 = v92 + int32(2)
	v98 = F_palloc(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_5))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L6
	} else {
		goto L53
	}
L24:
	;
	if v98 == int32(0) {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	if v92 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v253 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v98+v92))) = uint16(v253)
	if base.Ui32(v97) < base.Ui32(int32(2)) {
		v341 = v253
		goto L40
	} else {
		goto L41
	}
L27:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v92) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v116 = v4
	v123 = v76
	goto L31
L29:
	;
	v169 = v4
	goto L30
L30:
	;
	v188 = v92 & int32(3)
	if v188 == int32(0) {
		goto L26
	} else {
		goto L34
	}
L31:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v116))))
	*(*uint8)(unsafe.Add(mBase, uint32(v98+v116))) = uint8(v136)
	v139 = v116 | int32(1)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v139))))
	*(*uint8)(unsafe.Add(mBase, uint32(v98+v139))) = uint8(v142)
	v145 = v116 | int32(2)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v145))))
	*(*uint8)(unsafe.Add(mBase, uint32(v98+v145))) = uint8(v148)
	v151 = v116 | int32(3)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v151))))
	*(*uint8)(unsafe.Add(mBase, uint32(v98+v151))) = uint8(v154)
	v156 = int32(4)
	v157 = v116 + v156
	v159 = v123 + v156
	if v159 != v92&int32(-4) {
		v116 = v157
		v123 = v159
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v169 = v157
	goto L30
L33:
	;
	goto L32
L34:
	;
	v194 = v4
	v199 = v169
	goto L35
L35:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v199))))
	*(*uint8)(unsafe.Add(mBase, uint32(v98+v199))) = uint8(v219)
	v221 = int32(1)
	v224 = v194 + v221
	if v224 != v188 {
		v194 = v224
		v199 = v199 + v221
		goto L35
	} else {
		goto L37
	}
L36:
	;
	goto L26
L37:
	;
	goto L36
L38:
	;
	if v341 == int32(0) {
		goto L19
	} else {
		goto L52
	}
L39:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_6))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L6
	} else {
		goto L51
	}
L40:
	;
	goto L38
L41:
	;
	v259 = v97 - int32(2)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+v259))))
	if v261 != 0 {
		v341 = v253
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v98-int32(1)))))
	if v265 != 0 {
		v341 = v253
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v267 = F_palloc(m, int32(48))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	if v267 == int32(0) {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v271 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v267)+20)) = v271
	*(*int32)(unsafe.Add(mBase, uint32(v267)+8)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v267)+4)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v267)+12)) = v259
	*(*int64)(unsafe.Add(mBase, uint32(v267)+40)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v267)+24)) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v267)+16)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = v271
	F_syncrep_yyensure_buffer_stack(m, v90)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v285+v286<<(uint(int32(2))%32))))
	if v290 == v267 {
		v341 = v267
		goto L40
	} else {
		goto L47
	}
L47:
	;
	if v290 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v90)+36))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v292))) = uint8(v293)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v297 = int32(2)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v295+v296<<(uint(v297)%32))))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v90)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v300)+8)) = v301
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v303+v304<<(uint(v297)%32))))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v90)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v308)+16)) = v309
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v313 = v311
	v314 = v312
	goto L50
L49:
	;
	v313 = v285
	v314 = v286
	goto L50
L50:
	;
	v315 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v314<<(uint(v315)%32)+v313))) = v267
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v323 = v319 + v320<<(uint(v315)%32)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+28)) = v325
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+36)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v90)+80)) = v328
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v332
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+24)) = uint8(v334)
	*(*int32)(unsafe.Add(mBase, uint32(v90)+48)) = int32(1)
	v341 = v267
	goto L40
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v341)+20)) = int32(1)
	goto L18
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	if v2605+int32(816) != v2588 {
		goto L412
	} else {
		goto L413
	}
L57:
	;
	F_syncrep_yyerror(m, v393, v376, int32(_a_F_check_synchronous_standby_names_7))
	mBase = m.M
	v2586 = m.ExcPending
	if v2586 != 0 {
		goto L6
	} else {
		goto L411
	}
L58:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v386))) = uint16(v391)
	v401 = v394 << (uint(int32(1)) % 32)
	if base.Ui32(v373+v401-int32(2)) <= base.Ui32(v386) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	F_syncrep_yyerror(m, v2417, v2400, int32(_a_F_check_synchronous_standby_names_8))
	mBase = m.M
	v2549 = m.ExcPending
	if v2549 != 0 {
		goto L6
	} else {
		goto L407
	}
L60:
	;
	if base.Ui32(int32(_a_F_check_synchronous_standby_names_9)) < base.Ui32(v394) {
		goto L57
	} else {
		goto L63
	}
L61:
	;
	v454 = v373
	v456 = v384
	v458 = v386
	v459 = v388
	v460 = v394
	goto L62
L62:
	;
	v462 = int32(1) << (uint(v391) % 32)
	if v462&int32(13390146) != 0 {
		v2397 = v454
		v2398 = v374
		v2399 = v375
		v2400 = v376
		v2402 = v378
		v2408 = v456
		v2409 = v462
		v2410 = v458
		v2412 = v459
		v2414 = v390
		v2415 = v391
		v2417 = v393
		v2418 = v460
		v2422 = v398
		goto L86
	} else {
		goto L87
	}
L63:
	;
	v408 = int32(_a_F_check_synchronous_standby_names_10)
	if base.Ui32(v408) <= base.Ui32(v401) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v411 = v408
	goto L66
L65:
	;
	v411 = v401
	goto L66
L66:
	;
	v416 = F_palloc(m, v411*int32(6)|int32(3))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	if v416 == int32(0) {
		goto L57
	} else {
		goto L68
	}
L68:
	;
	v421 = int32(1)
	v424 = (v386-v373)>>(uint(v421)%32) + v421
	v426 = v424 << (uint(v421) % 32)
	if v426 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v431 = v428 + v411<<(uint(int32(1))%32)
	v433 = v424 << (uint(int32(2)) % 32)
	if v433 != 0 {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v427 = F__emscripten_memcpy_bulkmem(m, v416, v373, v426)
	mBase = m.M
	v428 = v427
	goto L72
L71:
	;
	v428 = v416
	goto L72
L72:
	;
	goto L69
L73:
	;
	if v390+int32(816) != v373 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v434 = F__emscripten_memcpy_bulkmem(m, v431, v388, v433)
	mBase = m.M
	v435 = v434
	goto L76
L75:
	;
	v435 = v431
	goto L76
L76:
	;
	goto L73
L77:
	;
	F_pfree(m, v373)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L6
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v441 = int32(1)
	v444 = v428 + v424<<(uint(v441)%32)
	if base.Ui32(v428+v411<<(uint(v441)%32)) <= base.Ui32(v444) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	v2588 = v428
	v2589 = v374
	v2590 = v375
	v2603 = v441
	v2605 = v390
	goto L56
L82:
	;
	goto L83
L83:
	;
	v454 = v428
	v456 = v435 + v433 - int32(4)
	v458 = v444 - int32(2)
	v459 = v435
	v460 = v411
	goto L62
L84:
	;
	goto L59
L85:
	;
	v373 = v2518
	v374 = v2519
	v375 = v2520
	v376 = v2521
	v378 = v2523
	v384 = v2529
	v386 = v2531 + int32(2)
	v388 = v2533
	v390 = v2535
	v391 = v2544
	v393 = v2538
	v394 = v2539
	v398 = v2543
	goto L58
L86:
	;
	if v2409&int32(_a_F_check_synchronous_standby_names_11) != 0 {
		goto L84
	} else {
		goto L388
	}
L87:
	;
	v467 = int32(*(*int8)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_check_synchronous_standby_names[1]))))
	if v378 == int32(-2) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v2377 = v467 + v2376
	if base.Ui32(int32(22)) < base.Ui32(v2377) {
		v2397 = v2339
		v2398 = v2340
		v2399 = v2341
		v2400 = v2342
		v2402 = v2375
		v2408 = v2350
		v2409 = v2351
		v2410 = v2352
		v2412 = v2354
		v2414 = v2356
		v2415 = v2357
		v2417 = v2359
		v2418 = v2360
		v2422 = v2364
		goto L86
	} else {
		goto L380
	}
L89:
	;
	v470 = m.G0
	v472 = v470 - int32(32)
	m.G0 = v472
	*(*int32)(unsafe.Add(mBase, uint32(v376)+92)) = v390 + int32(1228)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v376)+40))
	if v477 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v2339 = v454
	v2340 = v374
	v2341 = v375
	v2342 = v376
	v2344 = v378
	v2350 = v456
	v2351 = v462
	v2352 = v458
	v2354 = v459
	v2356 = v390
	v2357 = v391
	v2359 = v393
	v2360 = v460
	v2364 = v398
	goto L91
L91:
	;
	if v2344 <= int32(0) {
		goto L376
	} else {
		goto L377
	}
L92:
	;
	v2339 = v544
	v2340 = v545
	v2341 = v546
	v2342 = v547
	v2344 = v1688
	v2350 = v555
	v2351 = v556
	v2352 = v557
	v2354 = v559
	v2356 = v561
	v2357 = v562
	v2359 = v564
	v2360 = v565
	v2364 = v569
	goto L91
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v376)+40)) = int32(1)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v376)+44))
	if v482 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	v544 = v454
	v545 = v374
	v546 = v375
	v547 = v376
	v555 = v456
	v556 = v462
	v557 = v458
	v559 = v459
	v561 = v390
	v562 = v391
	v563 = v472
	v564 = v393
	v565 = v460
	v569 = v398
	goto L112
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v376)+44)) = int32(1)
	goto L98
L97:
	;
	goto L98
L98:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	if v487 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v491 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v376)+4)) = v491
	goto L101
L100:
	;
	goto L101
L101:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v376)+8))
	if v493 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v376)+8)) = v497
	goto L104
L103:
	;
	goto L104
L104:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v376)+20))
	if v499 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v376)+28)) = v527
	v531 = v524 + v525<<(uint(int32(2))%32)
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v376)+80)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v376)+36)) = v533
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	*(*int32)(unsafe.Add(mBase, uint32(v376)+4)) = v537
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
	*(*uint8)(unsafe.Add(mBase, uint32(v376)+24)) = uint8(v539)
	goto L95
L106:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v499+v500<<(uint(int32(2))%32))))
	if v504 != 0 {
		v524 = v499
		v525 = v500
		v526 = v504
		goto L105
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	F_syncrep_yyensure_buffer_stack(m, v376)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L6
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	v510 = F_syncrep_yy_create_buffer(m, v509, v376)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v376)+20))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	v514 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v512+v513<<(uint(v514)%32)))) = v510
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v376)+20))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v518+v519<<(uint(v514)%32))))
	v524 = v518
	v525 = v519
	v526 = v523
	goto L105
L112:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v547)+36))
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v570))) = uint8(v571)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v547)+44))
	v579 = v573
	v580 = v570
	v582 = v570
	goto L114
L114:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582))))
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600)+uint32(_c_F_check_synchronous_standby_names[4]))))
	if base.Ui32(int32(-26)) <= base.Ui32(v579&int32(2147483647)-int32(31)) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+68)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v547)+64)) = v579
	goto L118
L117:
	;
	goto L118
L118:
	;
	v612 = int32(1)
	v616 = int32(*(*int16)(unsafe.Add(mBase, uint32(v579<<(uint(v612)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v617 = v616 + v603
	v622 = int32(*(*int16)(unsafe.Add(mBase, uint32(v617<<(uint(v612)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v622 != v579 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v629 = v579
	v631 = v603
	v633 = v603
	goto L122
L120:
	;
	v688 = v617
	goto L121
L121:
	;
	v710 = int32(1)
	v716 = int32(*(*int16)(unsafe.Add(mBase, uint32(v688<<(uint(v710)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v688&int32(2147483647)))%64)&int64(-32985348833280) == int64(0) {
		v579 = v716
		v582 = v582 + v710
		goto L114
	} else {
		goto L128
	}
L122:
	;
	v654 = int32(*(*int16)(unsafe.Add(mBase, uint32(v629<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v629&int32(2147483647)))%64)&int64(2076672024) != int64(0) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v688 = v677
	goto L121
L124:
	;
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631)+uint32(_c_F_check_synchronous_standby_names[9]))))
	v667 = v666
	goto L126
L125:
	;
	v667 = v633
	goto L126
L126:
	;
	v671 = v667 & int32(255)
	v672 = int32(1)
	v676 = int32(*(*int16)(unsafe.Add(mBase, uint32(v654<<(uint(v672)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v677 = v671 + v676
	v682 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v677<<(uint(v672)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v654&int32(_a_F_check_synchronous_standby_names_12) != v682 {
		v629 = v654
		v631 = v671
		v633 = v667
		goto L122
	} else {
		goto L127
	}
L127:
	;
	goto L123
L128:
	;
	v736 = v580
	goto L129
L129:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v547)+64))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v547)+68))
	v759 = v752
	v764 = v736
	v768 = v753
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+80)) = v764
	*(*int32)(unsafe.Add(mBase, uint32(v547)+32)) = v768 - v764
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768))))
	*(*uint8)(unsafe.Add(mBase, uint32(v547)+24)) = uint8(v783)
	v785 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v768))) = uint8(v785)
	*(*int32)(unsafe.Add(mBase, uint32(v547)+36)) = v768
	v792 = int32(*(*int16)(unsafe.Add(mBase, uint32(v759<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[10]))))
	v797 = v792
	goto L133
L133:
	;
	v819 = int32(260)
	switch v797 {
	case 0:
		goto L162
	case 1:
		goto L112
	case 2:
		goto L149
	case 3:
		goto L148
	case 4:
		goto L161
	case 5:
		goto L160
	case 6:
		goto L159
	case 7:
		goto L158
	case 8:
		goto L156
	case 9:
		goto L155
	case 10:
		goto L154
	case 11:
		goto L147
	case 12:
		goto L146
	case 13:
		goto L145
	case 14:
		v1688 = v819
		goto L144
	case 15:
		goto L153
	case 16:
		goto L151
	case 17:
		goto L152
	case 18:
		goto L157
	default:
		goto L150
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+36)) = v2309
	*(*int32)(unsafe.Add(mBase, uint32(v547)+48)) = int32(0)
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v547)+44))
	v2336 = base.I32_div_s(v2332-int32(1), int32(2))
	v797 = v2336 + int32(17)
	goto L133
L136:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_13))
	mBase = m.M
	v2302 = m.ExcPending
	if v2302 != 0 {
		goto L6
	} else {
		goto L375
	}
L137:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_14))
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
		goto L6
	} else {
		goto L374
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	v2143 = v2121 + v2133
	*(*int32)(unsafe.Add(mBase, uint32(v547)+36)) = v2143
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v547)+44))
	if base.Ui32(v2143) <= base.Ui32(v2123) {
		v759 = v2145
		v764 = v2123
		v768 = v2143
		goto L131
	} else {
		goto L355
	}
L140:
	;
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1755)))
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+16)) = v1733
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v547)+28))
	if v1759 != 0 {
		v1881 = int32(0)
		goto L299
	} else {
		goto L300
	}
L141:
	;
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1733 = v1702
	v1755 = v1724 + v1725<<(uint(int32(2))%32)
	goto L140
L142:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_15))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L6
	} else {
		goto L298
	}
L143:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_16))
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L6
	} else {
		goto L297
	}
L144:
	;
	m.G0 = v563 + int32(32)
	goto L92
L145:
	;
	v1688 = int32(41)
	goto L144
L146:
	;
	v1688 = int32(40)
	goto L144
L147:
	;
	v1688 = int32(44)
	goto L144
L148:
	;
	v1688 = int32(262)
	goto L144
L149:
	;
	v1688 = int32(261)
	goto L144
L150:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_17))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L6
	} else {
		goto L296
	}
L151:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v547)+80))
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v768))) = uint8(v884)
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v890 = v886 + v887<<(uint(int32(2))%32)
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v891)+44))
	if v892 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L152:
	;
	v1688 = int32(0)
	goto L144
L153:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_18))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L6
	} else {
		goto L174
	}
L154:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v547)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v875))) = int32(_a_F_check_synchronous_standby_names_19)
	v1688 = int32(258)
	goto L144
L155:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v547)+80))
	v870 = F_pstrdup(m, v869)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L6
	} else {
		goto L173
	}
L156:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v547)+80))
	v864 = F_pstrdup(m, v863)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L6
	} else {
		goto L172
	}
L157:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	if v845 != 0 {
		v1688 = v819
		goto L144
	} else {
		goto L166
	}
L158:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v547)+92))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v836)))
	*(*int32)(unsafe.Add(mBase, uint32(v835))) = v837
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	*(*int32)(unsafe.Add(mBase, uint32(v839))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v547)+44)) = int32(1)
	v1688 = int32(258)
	goto L144
L159:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v547)+80))
	F_appendStringInfoString(m, v831, v832)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L6
	} else {
		goto L165
	}
L160:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	F_appendStringInfoChar(m, v827, int32(34))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L6
	} else {
		goto L164
	}
L161:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	F_initStringInfo(m, v822)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L6
	} else {
		goto L163
	}
L162:
	;
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v768))) = uint8(v820)
	v736 = v764
	goto L129
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+44)) = int32(3)
	goto L112
L164:
	;
	goto L112
L165:
	;
	goto L112
L166:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v547)+80))
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846))))
	if v847 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v563)+20)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v563)+16)) = int32(_a_F_check_synchronous_standby_names_20)
	v854 = F_psprintf(m, int32(_a_F_check_synchronous_standby_names_21), v563+int32(16))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L6
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v563))) = int32(_a_F_check_synchronous_standby_names_20)
	v860 = F_psprintf(m, int32(_a_F_check_synchronous_standby_names_22), v563)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L6
	} else {
		goto L171
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v564))) = v854
	v1688 = v819
	goto L144
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v564))) = v860
	v1688 = v819
	goto L144
L172:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v547)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v866))) = v864
	v1688 = int32(258)
	goto L144
L173:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v547)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v872))) = v870
	v1688 = int32(259)
	goto L144
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v891)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v547)+28)) = v895
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v897))) = v898
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v902 = int32(2)
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v900+v901<<(uint(v902)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v905)+44)) = int32(1)
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v908+v909<<(uint(v902)%32))))
	v914 = v913
	v915 = v908
	v916 = v909
	goto L177
L176:
	;
	v914 = v891
	v915 = v886
	v916 = v887
	goto L177
L177:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v547)+36))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v914)+4))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v547)+28))
	v920 = v918 + v919
	if base.Ui32(v917) <= base.Ui32(v920) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v547)+80))
	v926 = v922 + (v883 ^ int32(-1)) + v768
	*(*int32)(unsafe.Add(mBase, uint32(v547)+36)) = v926
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v547)+44))
	if base.Ui32(v922) < base.Ui32(v926) {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	goto L180
L180:
	;
	if base.Ui32(v920+int32(1)) < base.Ui32(v917) {
		goto L143
	} else {
		goto L213
	}
L181:
	;
	v935 = v928
	v938 = v922
	goto L184
L182:
	;
	v1083 = v928
	goto L183
L183:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v1083&int32(2147483647)-int32(31)) {
		goto L202
	} else {
		goto L203
	}
L184:
	;
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v938))))
	if v956 != 0 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v1083 = v1074
	goto L183
L186:
	;
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956)+uint32(_c_F_check_synchronous_standby_names[4]))))
	v961 = v959
	goto L188
L187:
	;
	v961 = int32(1)
	goto L188
L188:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v935&int32(2147483647)-int32(31)) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+68)) = v938
	*(*int32)(unsafe.Add(mBase, uint32(v547)+64)) = v935
	goto L191
L190:
	;
	goto L191
L191:
	;
	v971 = v961 & int32(255)
	v972 = int32(1)
	v976 = int32(*(*int16)(unsafe.Add(mBase, uint32(v935<<(uint(v972)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v977 = v971 + v976
	v982 = int32(*(*int16)(unsafe.Add(mBase, uint32(v977<<(uint(v972)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v982 != v935 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v989 = v935
	v991 = v961
	v993 = v971
	goto L195
L193:
	;
	v1048 = v977
	goto L194
L194:
	;
	v1070 = int32(1)
	v1074 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1048<<(uint(v1070)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v1076 = v938 + v1070
	if v1076 != v926 {
		v935 = v1074
		v938 = v1076
		goto L184
	} else {
		goto L201
	}
L195:
	;
	v1014 = int32(*(*int16)(unsafe.Add(mBase, uint32(v989<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v989&int32(2147483647)))%64)&int64(2076672024) != int64(0) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v1048 = v1037
	goto L194
L197:
	;
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v993)+uint32(_c_F_check_synchronous_standby_names[9]))))
	v1027 = v1026
	goto L199
L198:
	;
	v1027 = v991
	goto L199
L199:
	;
	v1031 = v1027 & int32(255)
	v1032 = int32(1)
	v1036 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1014<<(uint(v1032)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v1037 = v1031 + v1036
	v1042 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1037<<(uint(v1032)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v1014&int32(_a_F_check_synchronous_standby_names_12) != v1042 {
		v989 = v1014
		v991 = v1027
		v993 = v1031
		goto L195
	} else {
		goto L200
	}
L200:
	;
	goto L196
L201:
	;
	goto L185
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+68)) = v926
	*(*int32)(unsafe.Add(mBase, uint32(v547)+64)) = v1083
	goto L204
L203:
	;
	goto L204
L204:
	;
	v1112 = int32(1)
	v1116 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1083<<(uint(v1112)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v1118 = v1116 + v1112
	v1123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1118<<(uint(v1112)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v1123 != v1083 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v1130 = v1083
	goto L208
L206:
	;
	v1174 = v1118
	goto L207
L207:
	;
	v1198 = v1174 & int32(2147483647)
	if int64(1)<<(uint(base.I64_extend_i32_u(v1198))%64)&int64(-32985348833280) != int64(0) {
		v736 = v922
		goto L129
	} else {
		goto L211
	}
L208:
	;
	v1151 = int32(1)
	v1155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1130<<(uint(v1151)%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	v1156 = base.I32_extend16_s(v1155)
	v1161 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1156<<(uint(v1151)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v1163 = v1161 + v1151
	v1168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1163<<(uint(v1151)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v1155 != v1168 {
		v1130 = v1156
		goto L208
	} else {
		goto L210
	}
L209:
	;
	v1174 = v1163
	goto L207
L210:
	;
	goto L209
L211:
	;
	if v1198 == int32(0) {
		v736 = v922
		goto L129
	} else {
		goto L212
	}
L212:
	;
	v1207 = int32(1)
	v1208 = v926 + v1207
	*(*int32)(unsafe.Add(mBase, uint32(v547)+36)) = v1208
	v1214 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1174<<(uint(v1207)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v579 = v1214
	v580 = v922
	v582 = v1208
	goto L114
L213:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v547)+80))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v914)+40))
	if v1219 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	if v917-v1218 != int32(1) {
		v2121 = v918
		v2123 = v1218
		v2133 = v919
		goto L139
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v1227 = v1218 ^ int32(-1) + v917
	if v1227 != 0 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v2309 = v1218
	goto L135
L218:
	;
	v1228 = int32(7)
	v1229 = v1227 & v1228
	if base.Ui32(v917-v1218-int32(2)) < base.Ui32(v1228) {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v1386 = v914
	v1388 = v915
	v1390 = v916
	goto L220
L220:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+44))
	if v1407 == int32(2) {
		goto L234
	} else {
		goto L235
	}
L221:
	;
	if v1229 != 0 {
		goto L228
	} else {
		goto L229
	}
L222:
	;
	v1291 = v918
	v1292 = v1218
	goto L221
L223:
	;
	goto L224
L224:
	;
	v1242 = v918
	v1243 = v1218
	v1245 = int32(0)
	goto L225
L225:
	;
	v1264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1242))) = uint8(v1264)
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1242)+1)) = uint8(v1266)
	v1268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1242)+2)) = uint8(v1268)
	v1270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1242)+3)) = uint8(v1270)
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1242)+4)) = uint8(v1272)
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1242)+5)) = uint8(v1274)
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1242)+6)) = uint8(v1276)
	v1278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1242)+7)) = uint8(v1278)
	v1280 = int32(8)
	v1281 = v1242 + v1280
	v1283 = v1243 + v1280
	v1285 = v1245 + v1280
	if v1285 != v1227&int32(-8) {
		v1242 = v1281
		v1243 = v1283
		v1245 = v1285
		goto L225
	} else {
		goto L227
	}
L226:
	;
	v1291 = v1281
	v1292 = v1283
	goto L221
L227:
	;
	goto L226
L228:
	;
	v1318 = v1291
	v1319 = v1292
	v1321 = int32(0)
	goto L231
L229:
	;
	goto L230
L230:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1375+v1376<<(uint(int32(2))%32))))
	v1386 = v1380
	v1388 = v1375
	v1390 = v1376
	goto L220
L231:
	;
	v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1319))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1318))) = uint8(v1340)
	v1342 = int32(1)
	v1347 = v1321 + v1342
	if v1347 != v1229 {
		v1318 = v1318 + v1342
		v1319 = v1319 + v1342
		v1321 = v1347
		goto L231
	} else {
		goto L233
	}
L232:
	;
	goto L230
L233:
	;
	goto L232
L234:
	;
	v1410 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v547)+28)) = v1410
	v1733 = v1410
	v1755 = v1388 + v1390<<(uint(int32(2))%32)
	goto L140
L235:
	;
	goto L236
L236:
	;
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+12))
	v1417 = v1218 - v917
	v1418 = v1416 + v1417
	if v1418 == int32(0) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v547)+36))
	v1426 = v1416
	v1427 = v1386
	v1431 = v1421
	goto L240
L238:
	;
	v1489 = v1386
	v1491 = v1418
	goto L239
L239:
	;
	v1510 = int32(_a_F_check_synchronous_standby_names_23)
	if base.Ui32(v1510) <= base.Ui32(v1491) {
		goto L256
	} else {
		goto L257
	}
L240:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1427)+20))
	if v1448 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v1489 = v1479
	v1491 = v1481
	goto L239
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1427)+4)) = int32(0)
	goto L136
L243:
	;
	goto L244
L244:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v1427)+4))
	if base.Ui32(int32(2147483646)) <= base.Ui32(v1426) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1459 = int32(-3)
	goto L247
L246:
	;
	v1459 = v1426 << (uint(int32(1)) % 32)
	goto L247
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1427)+12)) = v1459
	v1462 = v1459 + int32(2)
	if v1453 != 0 {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1427)+4)) = v1467
	if v1467 == int32(0) {
		goto L136
	} else {
		goto L254
	}
L249:
	;
	v1463 = F_repalloc(m, v1453, v1462)
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L6
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v1465 = F_palloc(m, v1462)
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L6
	} else {
		goto L253
	}
L252:
	;
	v1467 = v1463
	goto L248
L253:
	;
	v1467 = v1465
	goto L248
L254:
	;
	v1472 = v1467 + (v1431 - v1453)
	*(*int32)(unsafe.Add(mBase, uint32(v547)+36)) = v1472
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1474+v1475<<(uint(int32(2))%32))))
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1479)+12))
	v1481 = v1480 + v1417
	if v1481 == int32(0) {
		v1426 = v1480
		v1427 = v1479
		v1431 = v1472
		goto L240
	} else {
		goto L255
	}
L255:
	;
	goto L241
L256:
	;
	v1513 = v1510
	goto L258
L257:
	;
	v1513 = v1491
	goto L258
L258:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+24))
	if v1515 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1520 = int32(0)
	goto L263
L260:
	;
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = int32(0)
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1595+v1596<<(uint(int32(2))%32))))
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1600)+4))
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	v1605 = F_fread(m, v1601+v1227, int32(1), v1513, v1604)
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L6
	} else {
		goto L278
	}
L262:
	;
	switch v1546 {
	case 0:
		goto L270
	default:
		v1590 = v1560
		goto L268
	case 11:
		goto L269
	}
L263:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	v1543 = F_do_getc(m, v1542)
	mBase = m.M
	v1544 = m.ExcPending
	if v1544 != 0 {
		goto L6
	} else {
		goto L266
	}
L264:
	;
	v1560 = v1513
	goto L262
L265:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v1547+v1548<<(uint(int32(2))%32))))
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1552)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1553+v1227+v1520))) = uint8(v1543)
	v1558 = v1520 + int32(1)
	if v1558 != v1513 {
		v1520 = v1558
		goto L263
	} else {
		goto L267
	}
L266:
	;
	v1546 = v1543 + int32(1)
	switch v1546 {
	case 0, 11:
		v1560 = v1520
		goto L262
	default:
		goto L265
	}
L267:
	;
	goto L264
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+28)) = v1590
	v1702 = v1590
	goto L141
L269:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1577+v1578<<(uint(int32(2))%32))))
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1582)+4))
	v1586 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v1583+v1227+v1560))) = uint8(v1586)
	v1590 = v1560 + int32(1)
	goto L268
L270:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+76))
	if v1562 < int32(0) {
		goto L273
	} else {
		goto L274
	}
L271:
	;
	if int32(base.Ui32(v1567)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v1590 = v1560
		goto L268
	} else {
		goto L276
	}
L272:
	;
	goto L271
L273:
	;
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1561)))
	v1567 = v1565
	goto L272
L274:
	;
	goto L275
L275:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1561)))
	v1567 = v1566
	goto L272
L276:
	;
	F_yy_fatal_error_4(m, int32(_a_F_check_synchronous_standby_names_15))
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L6
	} else {
		goto L277
	}
L277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L278:
	;
	v1611 = v1605
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+28)) = v1611
	if v1611 != 0 {
		v1702 = v1611
		goto L141
	} else {
		goto L281
	}
L281:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+76))
	if v1635 < int32(0) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	if int32(base.Ui32(v1640)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L287
	} else {
		goto L288
	}
L283:
	;
	goto L282
L284:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1634)))
	v1640 = v1638
	goto L283
L285:
	;
	goto L286
L286:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1634)))
	v1640 = v1639
	goto L283
L287:
	;
	v1702 = int32(0)
	goto L141
L288:
	;
	goto L289
L289:
	;
	v1649 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	if v1649 != int32(27) {
		goto L142
	} else {
		goto L290
	}
L290:
	;
	v1653 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = v1653
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1655)+76))
	if v1653 <= v1656 {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1667+v1668<<(uint(int32(2))%32))))
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v1672)+4))
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	v1677 = F_fread(m, v1673+v1227, int32(1), v1513, v1676)
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L6
	} else {
		goto L295
	}
L292:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1655)))
	*(*int32)(unsafe.Add(mBase, uint32(v1655))) = v1659 & int32(-49)
	goto L291
L293:
	;
	goto L294
L294:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1655)))
	*(*int32)(unsafe.Add(mBase, uint32(v1655))) = v1663 & int32(-49)
	goto L291
L295:
	;
	v1611 = v1677
	goto L279
L296:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L297:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v547)+28))
	v1883 = v1882 + v1227
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1885 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1884+v1885<<(uint(int32(2))%32))))
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v1889)+12))
	if base.Ui32(v1890) < base.Ui32(v1883) {
		goto L323
	} else {
		goto L324
	}
L300:
	;
	if v1227 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	if v1763 != 0 {
		goto L306
	} else {
		goto L307
	}
L302:
	;
	goto L303
L303:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1869 = int32(2)
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1867+v1868<<(uint(v1869)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1872)+44)) = v1869
	v1881 = v1869
	goto L299
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1829)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1829))) = v1762
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	if v1836 != 0 {
		goto L319
	} else {
		goto L320
	}
L305:
	;
	v1786 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1784+v1787<<(uint(int32(2))%32))))
	if v1791 == int32(0) {
		goto L313
	} else {
		goto L314
	}
L306:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1763+v1764<<(uint(int32(2))%32))))
	if v1768 != 0 {
		v1784 = v1763
		goto L305
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	F_syncrep_yyensure_buffer_stack(m, v547)
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L6
	} else {
		goto L310
	}
L309:
	;
	goto L308
L310:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	v1772 = F_syncrep_yy_create_buffer(m, v1771, v547)
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L6
	} else {
		goto L311
	}
L311:
	;
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1775 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1774+v1775<<(uint(int32(2))%32)))) = v1772
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	if v1780 != 0 {
		v1784 = v1780
		goto L305
	} else {
		goto L312
	}
L312:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	v1829 = int32(0)
	v1832 = v1782
	goto L304
L313:
	;
	v1829 = int32(0)
	v1832 = v1786
	goto L304
L314:
	;
	goto L315
L315:
	;
	v1795 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1791)+16)) = v1795
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1791)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1797))) = uint8(v1795)
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v1791)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v1800)+1)) = uint8(v1795)
	*(*int32)(unsafe.Add(mBase, uint32(v1791)+44)) = v1795
	*(*int32)(unsafe.Add(mBase, uint32(v1791)+28)) = int32(1)
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1791)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1791)+8)) = v1807
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	if v1809 == v1795 {
		v1829 = v1791
		v1832 = v1786
		goto L304
	} else {
		goto L316
	}
L316:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1815 = v1809 + v1812<<(uint(int32(2))%32)
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1815)))
	if v1791 != v1816 {
		v1829 = v1791
		v1832 = v1786
		goto L304
	} else {
		goto L317
	}
L317:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v547)+28)) = v1818
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v1815)))
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1820)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v547)+80)) = v1821
	*(*int32)(unsafe.Add(mBase, uint32(v547)+36)) = v1821
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1815)))
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1824)))
	*(*int32)(unsafe.Add(mBase, uint32(v547)+4)) = v1825
	v1827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1821))))
	*(*uint8)(unsafe.Add(mBase, uint32(v547)+24)) = uint8(v1827)
	v1829 = v1791
	v1832 = v1786
	goto L304
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1829)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0])) = v1832
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1853 = v1849 + v1850<<(uint(int32(2))%32)
	v1854 = *(*int32)(unsafe.Add(mBase, uint32(v1853)))
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1854)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v547)+28)) = v1855
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v1853)))
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1857)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v547)+36)) = v1858
	*(*int32)(unsafe.Add(mBase, uint32(v547)+80)) = v1858
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(v1853)))
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(v1861)))
	*(*int32)(unsafe.Add(mBase, uint32(v547)+4)) = v1862
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1858))))
	*(*uint8)(unsafe.Add(mBase, uint32(v547)+24)) = uint8(v1864)
	v1881 = int32(1)
	goto L299
L319:
	;
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1836+v1837<<(uint(int32(2))%32))))
	if v1829 == v1841 {
		goto L318
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1829)+32)) = int64(1)
	goto L318
L322:
	;
	goto L321
L323:
	;
	v1894 = v1883 + int32(base.Ui32(v1882)>>(uint(int32(1))%32))
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1889)+4))
	if v1895 != 0 {
		goto L327
	} else {
		goto L328
	}
L324:
	;
	v1924 = v1883
	v1925 = v1884
	v1926 = v1885
	goto L325
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+28)) = v1924
	v1928 = int32(2)
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1925+v1926<<(uint(v1928)%32))))
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1931)+4))
	v1934 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1932+v1924))) = uint8(v1934)
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1937 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1936+v1937<<(uint(v1928)%32))))
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1941)+4))
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v547)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v1942+v1943)+1)) = uint8(v1934)
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1951 = v1947 + v1948<<(uint(v1928)%32)
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1951)))
	v1953 = *(*int32)(unsafe.Add(mBase, uint32(v1952)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v547)+80)) = v1953
	if v1881 == int32(1) {
		v2309 = v1953
		goto L135
	} else {
		goto L333
	}
L326:
	;
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1902 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1903 = int32(2)
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v1901+v1902<<(uint(v1903)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+4)) = v1900
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v1908+v1909<<(uint(v1903)%32))))
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v1913)+4))
	if v1914 == int32(0) {
		goto L137
	} else {
		goto L332
	}
L327:
	;
	v1896 = F_repalloc(m, v1895, v1894)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L6
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	v1898 = F_palloc(m, v1894)
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L6
	} else {
		goto L331
	}
L330:
	;
	v1900 = v1896
	goto L326
L331:
	;
	v1900 = v1898
	goto L326
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1913)+12)) = v1894 - int32(2)
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(v547)+12))
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v547)+28))
	v1923 = *(*int32)(unsafe.Add(mBase, uint32(v547)+20))
	v1924 = v1921 + v1227
	v1925 = v1923
	v1926 = v1920
	goto L325
L333:
	;
	switch v1881 - int32(1) {
	case 0:
		goto L138
	case 1:
		goto L334
	default:
		goto L335
	}
L334:
	;
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v547)+28))
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v1951)))
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v2115)+4))
	v2121 = v2116
	v2123 = v1953
	v2133 = v2114
	goto L139
L335:
	;
	v1962 = v1953 + (v883 ^ int32(-1)) + v768
	*(*int32)(unsafe.Add(mBase, uint32(v547)+36)) = v1962
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v547)+44))
	if base.Ui32(v1962) <= base.Ui32(v1953) {
		v579 = v1964
		v580 = v1953
		v582 = v1962
		goto L114
	} else {
		goto L336
	}
L336:
	;
	v1971 = v1964
	v1980 = v1953
	goto L337
L337:
	;
	v1992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1980))))
	if v1992 != 0 {
		goto L339
	} else {
		goto L340
	}
L338:
	;
	v579 = v2110
	v580 = v1953
	v582 = v1962
	goto L114
L339:
	;
	v1995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1992)+uint32(_c_F_check_synchronous_standby_names[4]))))
	v1997 = v1995
	goto L341
L340:
	;
	v1997 = int32(1)
	goto L341
L341:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v1971&int32(2147483647)-int32(31)) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+68)) = v1980
	*(*int32)(unsafe.Add(mBase, uint32(v547)+64)) = v1971
	goto L344
L343:
	;
	goto L344
L344:
	;
	v2007 = v1997 & int32(255)
	v2008 = int32(1)
	v2012 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1971<<(uint(v2008)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v2013 = v2007 + v2012
	v2018 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2013<<(uint(v2008)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v2018 != v1971 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v2025 = v1971
	v2027 = v1997
	v2029 = v2007
	goto L348
L346:
	;
	v2084 = v2013
	goto L347
L347:
	;
	v2106 = int32(1)
	v2110 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2084<<(uint(v2106)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v2112 = v1980 + v2106
	if v1962 != v2112 {
		v1971 = v2110
		v1980 = v2112
		goto L337
	} else {
		goto L354
	}
L348:
	;
	v2050 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2025<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v2025&int32(2147483647)))%64)&int64(2076672024) != int64(0) {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	v2084 = v2073
	goto L347
L350:
	;
	v2062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2029)+uint32(_c_F_check_synchronous_standby_names[9]))))
	v2063 = v2062
	goto L352
L351:
	;
	v2063 = v2027
	goto L352
L352:
	;
	v2067 = v2063 & int32(255)
	v2068 = int32(1)
	v2072 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2050<<(uint(v2068)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v2073 = v2067 + v2072
	v2078 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2073<<(uint(v2068)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v2050&int32(_a_F_check_synchronous_standby_names_12) != v2078 {
		v2025 = v2050
		v2027 = v2063
		v2029 = v2067
		goto L348
	} else {
		goto L353
	}
L353:
	;
	goto L349
L354:
	;
	goto L338
L355:
	;
	v2152 = v2145
	v2155 = v2123
	goto L356
L356:
	;
	v2173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2155))))
	if v2173 != 0 {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	v759 = v2291
	v764 = v2123
	v768 = v2143
	goto L131
L358:
	;
	v2176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2173)+uint32(_c_F_check_synchronous_standby_names[4]))))
	v2178 = v2176
	goto L360
L359:
	;
	v2178 = int32(1)
	goto L360
L360:
	;
	if base.Ui32(int32(-26)) <= base.Ui32(v2152&int32(2147483647)-int32(31)) {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+68)) = v2155
	*(*int32)(unsafe.Add(mBase, uint32(v547)+64)) = v2152
	goto L363
L362:
	;
	goto L363
L363:
	;
	v2188 = v2178 & int32(255)
	v2189 = int32(1)
	v2193 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2152<<(uint(v2189)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v2194 = v2188 + v2193
	v2199 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2194<<(uint(v2189)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v2199 != v2152 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v2206 = v2152
	v2208 = v2178
	v2210 = v2188
	goto L367
L365:
	;
	v2265 = v2194
	goto L366
L366:
	;
	v2287 = int32(1)
	v2291 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2265<<(uint(v2287)%32))+uint32(_c_F_check_synchronous_standby_names[7]))))
	v2293 = v2155 + v2287
	if v2293 != v2143 {
		v2152 = v2291
		v2155 = v2293
		goto L356
	} else {
		goto L373
	}
L367:
	;
	v2231 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2206<<(uint(int32(1))%32))+uint32(_c_F_check_synchronous_standby_names[8]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v2206&int32(2147483647)))%64)&int64(2076672024) != int64(0) {
		goto L369
	} else {
		goto L370
	}
L368:
	;
	v2265 = v2254
	goto L366
L369:
	;
	v2243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2210)+uint32(_c_F_check_synchronous_standby_names[9]))))
	v2244 = v2243
	goto L371
L370:
	;
	v2244 = v2208
	goto L371
L371:
	;
	v2248 = v2244 & int32(255)
	v2249 = int32(1)
	v2253 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2231<<(uint(v2249)%32))+uint32(_c_F_check_synchronous_standby_names[5]))))
	v2254 = v2248 + v2253
	v2259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2254<<(uint(v2249)%32))+uint32(_c_F_check_synchronous_standby_names[6]))))
	if v2231&int32(_a_F_check_synchronous_standby_names_12) != v2259 {
		v2206 = v2231
		v2208 = v2244
		v2210 = v2248
		goto L367
	} else {
		goto L372
	}
L372:
	;
	goto L368
L373:
	;
	goto L357
L374:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L375:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L376:
	;
	v2367 = int32(0)
	v2375 = v2367
	v2376 = v2367
	goto L88
L377:
	;
	goto L378
L378:
	;
	if base.Ui32(int32(262)) < base.Ui32(v2344) {
		v2375 = v2344
		v2376 = int32(2)
		goto L88
	} else {
		goto L379
	}
L379:
	;
	v2374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2344)+uint32(_c_F_check_synchronous_standby_names[11]))))
	v2375 = v2344
	v2376 = v2374
	goto L88
L380:
	;
	v2382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2377)+uint32(_c_F_check_synchronous_standby_names[12]))))
	if v2376 != v2382 {
		v2397 = v2339
		v2398 = v2340
		v2399 = v2341
		v2400 = v2342
		v2402 = v2375
		v2408 = v2350
		v2409 = v2351
		v2410 = v2352
		v2412 = v2354
		v2414 = v2356
		v2415 = v2357
		v2417 = v2359
		v2418 = v2360
		v2422 = v2364
		goto L86
	} else {
		goto L381
	}
L381:
	;
	if v2377 != int32(19) {
		goto L382
	} else {
		goto L383
	}
L382:
	;
	v2387 = v2350 + int32(4)
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v2356)+1228))
	*(*int32)(unsafe.Add(mBase, uint32(v2387))) = v2388
	if v2375 != 0 {
		goto L385
	} else {
		goto L386
	}
L383:
	;
	goto L384
L384:
	;
	v2588 = v2339
	v2589 = v2340
	v2590 = v2341
	v2603 = int32(0)
	v2605 = v2356
	goto L56
L385:
	;
	v2392 = int32(-2)
	goto L387
L386:
	;
	v2392 = int32(0)
	goto L387
L387:
	;
	v2395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2377)+uint32(_c_F_check_synchronous_standby_names[13]))))
	v2518 = v2339
	v2519 = v2340
	v2520 = v2341
	v2521 = v2342
	v2523 = v2392
	v2529 = v2387
	v2531 = v2352
	v2533 = v2354
	v2535 = v2356
	v2538 = v2359
	v2539 = v2360
	v2543 = v2364
	v2544 = v2395
	goto L85
L388:
	;
	v2428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2415)+uint32(_c_F_check_synchronous_standby_names[14]))))
	v2431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2428)+uint32(_c_F_check_synchronous_standby_names[15]))))
	v2433 = int32(2)
	v2436 = *(*int32)(unsafe.Add(mBase, uint32(v2408+(int32(1)-v2431)<<(uint(v2433)%32))))
	switch v2428 - v2433 {
	case 0:
		goto L397
	case 1:
		goto L396
	case 2:
		goto L395
	case 3:
		goto L394
	case 4:
		goto L393
	case 5:
		goto L392
	case 6:
		goto L391
	case 7, 8:
		goto L390
	default:
		v2488 = v2436
		goto L389
	}
L389:
	;
	v2493 = v2408 - v2431<<(uint(int32(2))%32) + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2493))) = v2488
	v2497 = v2410 - v2431<<(uint(int32(1))%32)
	v2498 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2497))))
	v2501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2428)+uint32(_c_F_check_synchronous_standby_names[16]))))
	v2504 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2501)+uint32(_c_F_check_synchronous_standby_names[16]))))
	v2505 = v2498 + v2504
	if base.Ui32(int32(22)) < base.Ui32(v2505) {
		goto L404
	} else {
		goto L405
	}
L390:
	;
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(v2408)))
	v2488 = v2487
	goto L389
L391:
	;
	v2483 = *(*int32)(unsafe.Add(mBase, uint32(v2408-int32(8))))
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(v2408)))
	v2485 = F_lappend(m, v2483, v2484)
	mBase = m.M
	v2486 = m.ExcPending
	if v2486 != 0 {
		goto L6
	} else {
		goto L403
	}
L392:
	;
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v2408)))
	*(*int32)(unsafe.Add(mBase, uint32(v2414)+8)) = v2473
	*(*int32)(unsafe.Add(mBase, uint32(v2414)+12)) = v2473
	v2479 = F_list_make1_impl(m, int32(1), v2414+int32(8))
	mBase = m.M
	v2480 = m.ExcPending
	if v2480 != 0 {
		goto L6
	} else {
		goto L402
	}
L393:
	;
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(v2408-int32(12))))
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v2408-int32(4))))
	v2471 = F_create_syncrep_config(m, v2466, v2469, int32(0))
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L6
	} else {
		goto L401
	}
L394:
	;
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(v2408-int32(12))))
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(v2408-int32(4))))
	v2462 = F_create_syncrep_config(m, v2457, v2460, int32(1))
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L6
	} else {
		goto L400
	}
L395:
	;
	v2448 = *(*int32)(unsafe.Add(mBase, uint32(v2408-int32(12))))
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v2408-int32(4))))
	v2453 = F_create_syncrep_config(m, v2448, v2451, int32(0))
	mBase = m.M
	v2454 = m.ExcPending
	if v2454 != 0 {
		goto L6
	} else {
		goto L399
	}
L396:
	;
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(v2408)))
	v2444 = F_create_syncrep_config(m, int32(_a_F_check_synchronous_standby_names_24), v2442, int32(0))
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L6
	} else {
		goto L398
	}
L397:
	;
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(v2408)))
	*(*int32)(unsafe.Add(mBase, uint32(v2422))) = v2439
	v2488 = v2436
	goto L389
L398:
	;
	v2488 = v2444
	goto L389
L399:
	;
	v2488 = v2453
	goto L389
L400:
	;
	v2488 = v2462
	goto L389
L401:
	;
	v2488 = v2471
	goto L389
L402:
	;
	v2488 = v2479
	goto L389
L403:
	;
	v2488 = v2485
	goto L389
L404:
	;
	v2517 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2501)+uint32(_c_F_check_synchronous_standby_names[17]))))
	v2518 = v2397
	v2519 = v2398
	v2520 = v2399
	v2521 = v2400
	v2523 = v2402
	v2529 = v2493
	v2531 = v2497
	v2533 = v2412
	v2535 = v2414
	v2538 = v2417
	v2539 = v2418
	v2543 = v2422
	v2544 = v2517
	goto L85
L405:
	;
	v2510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2505)+uint32(_c_F_check_synchronous_standby_names[12]))))
	if v2510 != v2498 {
		goto L404
	} else {
		goto L406
	}
L406:
	;
	v2514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2505)+uint32(_c_F_check_synchronous_standby_names[13]))))
	v2518 = v2397
	v2519 = v2398
	v2520 = v2399
	v2521 = v2400
	v2523 = v2402
	v2529 = v2493
	v2531 = v2497
	v2533 = v2412
	v2535 = v2414
	v2538 = v2417
	v2539 = v2418
	v2543 = v2422
	v2544 = v2514
	goto L85
L407:
	;
	v2563 = v2410
	goto L408
L408:
	;
	if base.B2i32(v2397 == v2563) == int32(0) {
		v2563 = v2563 - int32(2)
		goto L408
	} else {
		goto L410
	}
L409:
	;
	v2588 = v2397
	v2589 = v2398
	v2590 = v2399
	v2603 = int32(1)
	v2605 = v2414
	goto L56
L410:
	;
	goto L409
L411:
	;
	v2588 = v373
	v2589 = v374
	v2590 = v375
	v2603 = int32(2)
	v2605 = v390
	goto L56
L412:
	;
	F_pfree(m, v2588)
	mBase = m.M
	v2618 = m.ExcPending
	if v2618 != 0 {
		goto L6
	} else {
		goto L415
	}
L413:
	;
	goto L414
L414:
	;
	m.G0 = v2605 + int32(1232)
	v2622 = *(*int32)(unsafe.Add(mBase, uint32(v2590)+44))
	F_replication_scanner_finish(m, v2622)
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L6
	} else {
		goto L416
	}
L415:
	;
	goto L414
L416:
	;
	if v2603 == int32(0) {
		goto L418
	} else {
		goto L419
	}
L417:
	;
	v2653 = *(*int32)(unsafe.Add(mBase, uint32(v2627)+4))
	if v2653 <= int32(0) {
		goto L430
	} else {
		goto L431
	}
L418:
	;
	v2627 = *(*int32)(unsafe.Add(mBase, uint32(v2590)+40))
	if v2627 != 0 {
		goto L417
	} else {
		goto L421
	}
L419:
	;
	goto L420
L420:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[18])) = int32(16801924)
	goto L422
L421:
	;
	goto L420
L422:
	;
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v2590)+36))
	v2634 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[19])) = v2634
	goto L423
L423:
	;
	if v2632 != 0 {
		goto L425
	} else {
		goto L426
	}
L424:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[20])) = v2650
	v2713 = v2590
	v2737 = int32(0)
	goto L1
L425:
	;
	v2638 = *(*int32)(unsafe.Add(mBase, uint32(v2590)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v2590)+16)) = v2638
	v2643 = F_format_elog_string(m, int32(_a_F_check_synchronous_standby_names_25), v2590+int32(16))
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L6
	} else {
		goto L428
	}
L426:
	;
	goto L427
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2590))) = int32(_a_F_check_synchronous_standby_names_26)
	v2648 = F_format_elog_string(m, int32(_a_F_check_synchronous_standby_names_27), v2590)
	mBase = m.M
	v2649 = m.ExcPending
	if v2649 != 0 {
		goto L6
	} else {
		goto L429
	}
L428:
	;
	v2650 = v2643
	goto L424
L429:
	;
	v2650 = v2648
	goto L424
L430:
	;
	v2657 = *(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[19])) = v2657
	goto L433
L431:
	;
	goto L432
L432:
	;
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v2627)))
	v2673 = F_guc_malloc(m, v2672)
	mBase = m.M
	v2674 = m.ExcPending
	if v2674 != 0 {
		goto L6
	} else {
		goto L435
	}
L433:
	;
	v2660 = *(*int32)(unsafe.Add(mBase, uint32(v2590)+40))
	v2661 = *(*int32)(unsafe.Add(mBase, uint32(v2660)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2590)+32)) = v2661
	v2667 = F_format_elog_string(m, int32(_a_F_check_synchronous_standby_names_28), v2590+int32(32))
	mBase = m.M
	v2668 = m.ExcPending
	if v2668 != 0 {
		goto L6
	} else {
		goto L434
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_synchronous_standby_names[21])) = v2667
	v2713 = v2590
	v2737 = int32(0)
	goto L1
L435:
	;
	if v2673 == int32(0) {
		v2713 = v2590
		v2737 = int32(0)
		goto L1
	} else {
		goto L436
	}
L436:
	;
	v2677 = *(*int32)(unsafe.Add(mBase, uint32(v2590)+40))
	v2678 = *(*int32)(unsafe.Add(mBase, uint32(v2677)))
	if v2678 != 0 {
		goto L438
	} else {
		goto L439
	}
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2589))) = v2680
	v2686 = v2590
	goto L2
L438:
	;
	v2679 = F__emscripten_memcpy_bulkmem(m, v2673, v2677, v2678)
	mBase = m.M
	v2680 = v2679
	goto L440
L439:
	;
	v2680 = v2673
	goto L440
L440:
	;
	goto L437
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
