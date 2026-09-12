package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RollbackToSavepoint(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+72))
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L14
	} else {
		goto L81
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L14
	} else {
		goto L77
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L14
	} else {
		goto L73
	}
L4:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+76)))
	if v14 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	if int32(0) <= v16 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if base.Ui32(int32(19)) < base.Ui32(v19) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v80 = v12
	goto L27
L8:
	;
	v23 = int32(1) << (uint(v19) % 32)
	if v23&int32(1011559) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v23&int32(136) != 0 {
		goto L2
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L14
	} else {
		goto L19
	}
L12:
	;
	if v19 != int32(4) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return
L15:
	;
	F_errcode(m, int32(16908610))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = int32(542689)
	F_errmsg(m, int32(163338), v9+int32(80))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(515908), int32(4604), int32(94714))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if base.Ui32(v55) <= base.Ui32(int32(19)) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v65
	F_errmsg_internal(m, int32(197273), v9+int32(96))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L14
	} else {
		goto L24
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v55<<(uint(int32(2))%32))+uint32(_consts[164])))
	v65 = v64
	goto L23
L22:
	;
	v65 = int32(568776)
	goto L23
L23:
	;
	goto L20
L24:
	;
	F_errfinish(m, int32(515908), int32(4631), int32(94714))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L14
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
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v128 != v129 {
		goto L1
	} else {
		goto L46
	}
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	if v83 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L14
	} else {
		goto L42
	}
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v87 == int32(0) {
		v106 = v86
		v107 = v87
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v80)+80))
	if v111 != 0 {
		v80 = v111
		goto L27
	} else {
		goto L41
	}
L32:
	;
	if v107-v106 == int32(0) {
		goto L26
	} else {
		goto L40
	}
L33:
	;
	goto L32
L34:
	;
	if v86 != v87 {
		v106 = v86
		v107 = v87
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v91 = v83
	v92 = l0
	goto L36
L36:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if v96 == int32(0) {
		v106 = v95
		v107 = v96
		goto L33
	} else {
		goto L38
	}
L37:
	;
	v106 = v95
	v107 = v96
	goto L33
L38:
	;
	v99 = int32(1)
	if v95 == v96 {
		v91 = v91 + v99
		v92 = v92 + v99
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	goto L31
L41:
	;
	goto L28
L42:
	;
	F_errcode(m, int32(16778371))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg(m, int32(76603), v9)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(515908), int32(4644), int32(94714))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L14
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
	if v12 != v80 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v134 = v12
	goto L50
L48:
	;
	v176 = v12
	v178 = v19
	goto L49
L49:
	;
	switch v178 - int32(12) {
	case 0:
		v210 = int32(18)
		goto L63
	default:
		goto L65
	case 3:
		goto L64
	}
L50:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
	switch v139 - int32(12) {
	case 0:
		v169 = int32(17)
		goto L52
	default:
		goto L54
	case 3:
		goto L53
	}
L51:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v171)+24))
	v176 = v171
	v178 = v173
	goto L49
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+24)) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v134)+80))
	if v171 != v80 {
		v134 = v171
		goto L50
	} else {
		goto L62
	}
L53:
	;
	v169 = int32(16)
	goto L52
L54:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L14
	} else {
		goto L55
	}
L55:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
	if base.Ui32(v146) <= base.Ui32(int32(19)) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v156
	F_errmsg_internal(m, int32(197273), v9+int32(32))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L14
	} else {
		goto L60
	}
L57:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v146<<(uint(int32(2))%32))+uint32(_consts[164])))
	v156 = v155
	goto L59
L58:
	;
	v156 = int32(568776)
	goto L59
L59:
	;
	goto L56
L60:
	;
	F_errfinish(m, int32(515908), int32(4668), int32(94714))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L14
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	goto L51
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+24)) = v210
	m.G0 = v9 + int32(112)
	return
L64:
	;
	v210 = int32(19)
	goto L63
L65:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L14
	} else {
		goto L66
	}
L66:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v176)+24))
	if base.Ui32(v187) <= base.Ui32(int32(19)) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v197
	F_errmsg_internal(m, int32(197273), v9+int32(16))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L14
	} else {
		goto L71
	}
L68:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v187<<(uint(int32(2))%32))+uint32(_consts[164])))
	v197 = v196
	goto L70
L69:
	;
	v197 = int32(568776)
	goto L70
L70:
	;
	goto L67
L71:
	;
	F_errfinish(m, int32(515908), int32(4680), int32(94714))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L14
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L14
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(272433), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L14
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(515908), int32(4583), int32(94714))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L14
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errcode(m, int32(16778371))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L14
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = l0
	F_errmsg(m, int32(76603), v9-int32(-64))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L14
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(515908), int32(4595), int32(94714))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L14
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_errcode(m, int32(16778371))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L14
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l0
	F_errmsg(m, int32(320460), v9+int32(48))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L14
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(515908), int32(4650), int32(94714))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L14
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_add_to_tsvector(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v19 = v8
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_parsetext(m, v20, v6, l1, l2)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if v19 < v23 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v25 + int32(1)
			} else {
			}
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(16)
		v12 = F_palloc(m, int32(256))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v12
			v19 = int32(0)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F_parsetext(m, v20, v6, l1, l2)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				if v19 < v23 {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v25 + int32(1)
				} else {
				}
				return
			}
		}
	}
}
func F_coerce_to_common_type(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l2
	v12 = F_exprType(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v12
		if l2 != v12 {
			v24 = F_can_coerce_type(m, int32(1), v9+int32(24), v9+int32(28), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v24 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(101744772))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = F_format_type_be(m, v12)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = F_format_type_be(m, l2)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v48
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v46
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
									F_errmsg(m, int32(193280), v9)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int32(0)
									} else {
										v56 = F_exprLocation(m, l1)
										mBase = m.M
										F_parser_errposition(m, l0, v56)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(523107), int32(1590), int32(383929))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
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
				} else {
					v28 = int32(-1)
					v32 = F_coerce_type(m, l0, l1, v12, l2, v28, int32(0), int32(2), v28)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = v32
						m.G0 = v9 + int32(32)
						return v34
					}
				}
			}
		} else {
			v34 = l1
			m.G0 = v9 + int32(32)
			return v34
		}
	}
}
func F_convert_to_base_unit(m *base.Module, l0 float64, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 float64
	_ = v112
	var v116 float64
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v16 {
	case 0, 9, 10, 11, 12, 13, 32:
		v35 = l1
		v36 = v12 + int32(12)
		goto L1
	default:
		goto L2
	}
L1:
	;
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v37)
	v46 = v35
	goto L5
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v16)
	v21 = l1 + int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	switch v22 {
	case 0, 9, 10, 11, 12, 13, 32:
		v35 = v21
		v36 = v12 + int32(13)
		goto L1
	default:
		goto L3
	}
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v22)
	v27 = l1 + int32(2)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	switch v28 {
	case 0, 9, 10, 11, 12, 13, 32:
		v35 = v27
		v36 = v12 + int32(14)
		goto L1
	default:
		goto L4
	}
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v28)
	v35 = l1 + int32(3)
	v36 = v12 + int32(15)
	goto L1
L5:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if base.Ui32(v49-int32(9)) < base.Ui32(int32(5)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v46 = v46 + int32(1)
	goto L5
L8:
	;
	if v49 == int32(32) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v49 != 0 {
		v131 = v37
		goto L10
	} else {
		goto L11
	}
L10:
	;
	m.G0 = v12 + int32(16)
	return v131
L11:
	;
	if l2&int32(251658240) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v60 = int32(1796288)
	goto L14
L13:
	;
	v60 = int32(1796704)
	goto L14
L14:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v61 == int32(0) {
		v131 = v37
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v68 = v37
	goto L16
L16:
	;
	v75 = v60 + v68<<(uint(int32(4))%32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if l2 != v76 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v131 = int32(0)
	goto L10
L18:
	;
	v121 = v68 + int32(1)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v121<<(uint(int32(4))%32)))))
	if v125 != 0 {
		v68 = v121
		goto L16
	} else {
		goto L32
	}
L19:
	;
	v79 = v12 + int32(12)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v83 == int32(0) {
		v102 = v82
		v103 = v83
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v103-v102 != 0 {
		goto L18
	} else {
		goto L28
	}
L21:
	;
	goto L20
L22:
	;
	if v82 != v83 {
		v102 = v82
		v103 = v83
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v87 = v79
	v88 = v75
	goto L24
L24:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	if v92 == int32(0) {
		v102 = v91
		v103 = v92
		goto L21
	} else {
		goto L26
	}
L25:
	;
	v102 = v91
	v103 = v92
	goto L21
L26:
	;
	v95 = int32(1)
	if v91 == v92 {
		v87 = v87 + v95
		v88 = v88 + v95
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v105 = *(*float64)(unsafe.Add(mBase, uint32(v75)+8))
	v106 = base.F64_mul(l0, v105)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+16)))
	if v107 == int32(0) {
		v116 = v106
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v116
	v131 = int32(1)
	goto L10
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	if l2 != v110 {
		v116 = v106
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v112 = *(*float64)(unsafe.Add(mBase, uint32(v75)+24))
	v116 = base.F64_mul(v112, base.F64_nearest(base.F64_div(v106, v112)))
	goto L29
L32:
	;
	goto L17
}
func F_push_to_sink(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
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
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	if l4 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v91 != 0 {
		goto L28
	} else {
		goto L29
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L10
	} else {
		goto L24
	}
L3:
	;
	return
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = v10 + v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = v13 - v10
	if base.Ui32(l4) < base.Ui32(v14) {
		v90 = l3
		v91 = l4
		v93 = v12
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v14 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	m.T0[v20].(func(*base.Module, int32, int32))(m, l0, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v16 = F__emscripten_memcpy_bulkmem(m, v12, l3, v14)
	mBase = m.M
	goto L9
L8:
	;
	goto L9
L9:
	;
	goto L6
L10:
	;
	return
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = F_pg_checksum_update(m, l1, v23, v24)
	mBase = m.M
	if v25 < int32(0) {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v28 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v28
	v30 = l4 - v14
	if v30 == v28 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v37 = l3 + v14
	v38 = v30
	goto L14
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v38) < base.Ui32(v42) {
		v90 = v37
		v91 = v38
		v93 = v41
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L3
L16:
	;
	if v42 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	m.T0[v48].(func(*base.Module, int32, int32))(m, l0, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L21
	}
L18:
	;
	v44 = F__emscripten_memcpy_bulkmem(m, v41, v37, v42)
	mBase = m.M
	goto L20
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v53 = F_pg_checksum_update(m, l1, v51, v52)
	mBase = m.M
	if v53 < int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v59 = v38 - v42
	if v59 != 0 {
		v37 = v37 + v42
		v38 = v59
		goto L14
	} else {
		goto L23
	}
L23:
	;
	goto L15
L24:
	;
	F_errmsg_internal(m, int32(299746), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(518781), int32(1980), int32(330230))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v96 + v91
	return
L28:
	;
	v94 = F__emscripten_memcpy_bulkmem(m, v93, v90, v91)
	mBase = m.M
	goto L30
L29:
	;
	goto L30
L30:
	;
	goto L27
}
func F_to_ascii_enc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_copy(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if base.Ui32(int32(42)) <= base.Ui32(v13) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v13
					F_errmsg(m, int32(433071), v6)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(521323), int32(146), int32(512061))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
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
			v32 = F_encode_to_ascii(m, v9, v13)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v32
			}
		}
	}
}
func F_to_hex32(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	v7 = m.G0
	v8 = int32(-64)
	v9 = v7 + v8
	m.G0 = v9
	v11 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+20)))
	v13 = v9 - v8
	v14 = v13
	v15 = v11
	goto L1
L1:
	;
	v21 = v14 - int32(1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v15)&int32(15))+uint32(_consts[1155]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v27)
	if base.Ui64(v15) < base.Ui64(int64(16)) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v35 = v13 - v21
	v37 = v35 + int32(4)
	v38 = F_palloc(m, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(v9) < base.Ui32(v21) {
		v14 = v21
		v15 = int64(base.Ui64(v15) >> (uint(int64(4)) % 64))
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v37 << (uint(int32(2)) % 32)
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v9 - int32(-64)
	return v38
L9:
	;
	v47 = F__emscripten_memcpy_bulkmem(m, v38+int32(4), v21, v35)
	mBase = m.M
	goto L11
L10:
	;
	goto L11
L11:
	;
	goto L8
}
func F_to_json(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = F_get_fn_expr_argtype(m, v11, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(388219), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519339), int32(749), int32(257504))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
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
			F_json_categorize_type(m, v13, int32(0), v8+int32(12), v8+int32(8))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				v45 = F_makeStringInfo(m)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_datum_to_json_internal(m, v10, int32(0), v45, v43, v42, int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
						v52 = F_cstring_to_text_with_len(m, v50, v51)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return v52
						}
					}
				}
			}
		}
	}
}
func F_to_jsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = F_get_fn_expr_argtype(m, v10, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(388219), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(524345), int32(1098), int32(527100))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
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
			F_json_categorize_type(m, v12, int32(1), v7+int32(12), v7+int32(8))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v43 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v43
				*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v43
				v47 = int32(0)
				F_datum_to_jsonb_internal(m, v9, v47, v7+int32(16), v42, v41, v47)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
					v54 = F_JsonbValueToJsonb(m, v53)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(32)
						return v54
					}
				}
			}
		}
	}
}
func F_to_regproc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_text_to_cstring(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[1141]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _consts[1142]))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v20
			v26 = F_DirectInputFunctionCallSafe(m, int32(1494), v14, int32(-1), v7, v7+int32(12))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v26 == int32(0) {
					v30 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
					v33 = int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					v33 = v32
				}
				m.G0 = v7 + int32(16)
				return v33
			}
		}
	}
}
func F_to_regtypemod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_text_to_cstring(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _consts[1141]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _consts[1142]))
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v20
			v28 = F_parseTypeString(m, v14, v7+int32(28), v7+int32(24), v7+int32(8))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				if v28 == int32(0) {
					v32 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
					v35 = int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
					v35 = v34
				}
				m.G0 = v7 + int32(32)
				return v35
			}
		}
	}
}
func F_to_timestamp(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v120 int32
	_ = v120
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	v3 = m.G0
	v5 = v3 - int32(80)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v16 = int32(0)
			v27 = F_do_to_timestamp(m, v8, v13, v15, v16, v5+int32(24), v5+int32(12), v5+int32(16), v5+int32(8), v16, v16)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+16)))
				if v29 == int32(1) {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
					v48 = v32
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, _consts[1067]))
					v38 = m.G0
					v39 = int32(16)
					v40 = v38 - v39
					m.G0 = v40
					v44 = F_DetermineTimeZoneOffsetInternal(m, v5+int32(24), v36, v40+int32(8))
					mBase = m.M
					m.G0 = v40 + v39
					v48 = v44
				}
				*(*int32)(unsafe.Add(mBase, uint32(v5)+68)) = v48
				v51 = v5 + int32(24)
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				v54 = v5 + int32(68)
				v56 = v5 + int32(72)
				v63 = m.G0
				v65 = v63 - int32(16)
				m.G0 = v65
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
				if v67 <= int32(-4713) {
					if v67 != int32(-4713) {
						*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
						v144 = int32(-1)
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
						if int32(10) < v72 {
							v83 = v72
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
							v85 = F_date2j(m, v67, v83, v84)
							mBase = m.M
							v88 = base.I64_extend_i32_s(v85 - int32(2451545))
							v89 = int64(63)
							F___multi3(m, v65, v88, v88>>(uint(v89)%64), int64(86400000000), int64(0))
							mBase = m.M
							v94 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
							v95 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
							if v94 != v95>>(uint(v89)%64) {
								*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
								v144 = int32(-1)
							} else {
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
								v103 = int32(60)
								v112 = base.I64_extend_i32_s(v52) + base.I64_extend_i32_s(v100+(v101+v102*v103)*v103)*int64(1000000)
								v113 = v95 + v112
								*(*int64)(unsafe.Add(mBase, uint32(v56))) = v113
								if base.B2i32(v112 < int64(0))^base.B2i32(v113 < v95) != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
									v144 = int32(-1)
								} else {
									if v54 != 0 {
										v120 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
										v125 = base.I64_extend_i32_s(int32(0)-v120)*int64(-1000000) + v113
										*(*int64)(unsafe.Add(mBase, uint32(v56))) = v125
										v127 = v125
									} else {
										v127 = v113
									}
									if base.Ui64(v127+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
										v144 = int32(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
										v144 = int32(-1)
									}
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
							v144 = int32(-1)
						}
					}
				} else {
					if v67 <= int32(5874897) {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
						v83 = v77
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
						v85 = F_date2j(m, v67, v83, v84)
						mBase = m.M
						v88 = base.I64_extend_i32_s(v85 - int32(2451545))
						v89 = int64(63)
						F___multi3(m, v65, v88, v88>>(uint(v89)%64), int64(86400000000), int64(0))
						mBase = m.M
						v94 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
						v95 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
						if v94 != v95>>(uint(v89)%64) {
							*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
							v144 = int32(-1)
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
							v103 = int32(60)
							v112 = base.I64_extend_i32_s(v52) + base.I64_extend_i32_s(v100+(v101+v102*v103)*v103)*int64(1000000)
							v113 = v95 + v112
							*(*int64)(unsafe.Add(mBase, uint32(v56))) = v113
							if base.B2i32(v112 < int64(0))^base.B2i32(v113 < v95) != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
								v144 = int32(-1)
							} else {
								if v54 != 0 {
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
									v125 = base.I64_extend_i32_s(int32(0)-v120)*int64(-1000000) + v113
									*(*int64)(unsafe.Add(mBase, uint32(v56))) = v125
									v127 = v125
								} else {
									v127 = v113
								}
								if base.Ui64(v127+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
									v144 = int32(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
									v144 = int32(-1)
								}
							}
						}
					} else {
						if v67 != int32(5874898) {
							*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
							v144 = int32(-1)
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
							if int32(5) < v80 {
								*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
								v144 = int32(-1)
							} else {
								v83 = v80
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
								v85 = F_date2j(m, v67, v83, v84)
								mBase = m.M
								v88 = base.I64_extend_i32_s(v85 - int32(2451545))
								v89 = int64(63)
								F___multi3(m, v65, v88, v88>>(uint(v89)%64), int64(86400000000), int64(0))
								mBase = m.M
								v94 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
								v95 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
								if v94 != v95>>(uint(v89)%64) {
									*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
									v144 = int32(-1)
								} else {
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
									v103 = int32(60)
									v112 = base.I64_extend_i32_s(v52) + base.I64_extend_i32_s(v100+(v101+v102*v103)*v103)*int64(1000000)
									v113 = v95 + v112
									*(*int64)(unsafe.Add(mBase, uint32(v56))) = v113
									if base.B2i32(v112 < int64(0))^base.B2i32(v113 < v95) != 0 {
										*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
										v144 = int32(-1)
									} else {
										if v54 != 0 {
											v120 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
											v125 = base.I64_extend_i32_s(int32(0)-v120)*int64(-1000000) + v113
											*(*int64)(unsafe.Add(mBase, uint32(v56))) = v125
											v127 = v125
										} else {
											v127 = v113
										}
										if base.Ui64(v127+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
											v144 = int32(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v56))) = int64(0)
											v144 = int32(-1)
										}
									}
								}
							}
						}
					}
				}
				m.G0 = v65 + int32(16)
				if v144 == int32(0) {
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
					if v150 != 0 {
						F_AdjustTimestampForTypmod(m, v5+int32(72), v150, int32(0))
						mBase = m.M
						v155 = m.ExcPending
						if v155 != 0 {
							return int32(0)
						} else {
							v156 = *(*int64)(unsafe.Add(mBase, uint32(v5)+72))
							v157 = F_Int64GetDatum(m, v156)
							mBase = m.M
							v158 = m.ExcPending
							if v158 != 0 {
								return int32(0)
							} else {
								m.G0 = v5 + int32(80)
								return v157
							}
						}
					} else {
						v156 = *(*int64)(unsafe.Add(mBase, uint32(v5)+72))
						v157 = F_Int64GetDatum(m, v156)
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return int32(0)
						} else {
							m.G0 = v5 + int32(80)
							return v157
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(420796), int32(0))
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(521728), int32(4140), int32(248378))
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
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
