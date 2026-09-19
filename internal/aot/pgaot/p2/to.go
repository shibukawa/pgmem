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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_RollbackToSavepoint[0]))
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
	v244 = m.ExcPending
	if v244 != 0 {
		goto L14
	} else {
		goto L80
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L14
	} else {
		goto L76
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L14
	} else {
		goto L72
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
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_RollbackToSavepoint[1]))
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
	v77 = v12
	goto L27
L8:
	;
	v23 = int32(1) << (uint(v19) % 32)
	if v23&int32(_a_F_RollbackToSavepoint_0) == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = int32(_a_F_RollbackToSavepoint_1)
	F_errmsg(m, int32(_a_F_RollbackToSavepoint_2), v9+int32(80))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_RollbackToSavepoint_3), int32(_a_F_RollbackToSavepoint_4), int32(_a_F_RollbackToSavepoint_5))
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
	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v62
	F_errmsg_internal(m, int32(_a_F_RollbackToSavepoint_6), v9+int32(96))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L14
	} else {
		goto L24
	}
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55<<(uint(int32(2))%32))+uint32(_c_F_RollbackToSavepoint[2])))
	v62 = v60
	goto L23
L22:
	;
	v62 = int32(_a_F_RollbackToSavepoint_7)
	goto L23
L23:
	;
	goto L20
L24:
	;
	F_errfinish(m, int32(_a_F_RollbackToSavepoint_3), int32(_a_F_RollbackToSavepoint_8), int32(_a_F_RollbackToSavepoint_5))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
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
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v126 != v127 {
		goto L1
	} else {
		goto L45
	}
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v80 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L14
	} else {
		goto L41
	}
L29:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v83 == int32(0))|base.B2i32(v83 != v86) != 0 {
		v104 = v83
		v105 = v86
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v77)+80))
	if v109 != 0 {
		v77 = v109
		goto L27
	} else {
		goto L40
	}
L32:
	;
	if v104-v105 == int32(0) {
		goto L26
	} else {
		goto L39
	}
L33:
	;
	goto L32
L34:
	;
	v89 = v80
	v90 = l0
	goto L35
L35:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	if v94 == int32(0) {
		v104 = v94
		v105 = v93
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v104 = v94
	v105 = v93
	goto L33
L37:
	;
	v97 = int32(1)
	if v94 == v93 {
		v89 = v89 + v97
		v90 = v90 + v97
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L31
L40:
	;
	goto L28
L41:
	;
	F_errcode(m, int32(16778371))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg(m, int32(_a_F_RollbackToSavepoint_9), v9)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_RollbackToSavepoint_3), int32(_a_F_RollbackToSavepoint_10), int32(_a_F_RollbackToSavepoint_5))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L14
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
	if v12 != v77 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v132 = v12
	goto L49
L47:
	;
	v171 = v12
	v173 = v19
	goto L48
L48:
	;
	switch v173 - int32(12) {
	case 0:
		v202 = int32(18)
		goto L62
	default:
		goto L64
	case 3:
		goto L63
	}
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v132)+24))
	switch v137 - int32(12) {
	case 0:
		v164 = int32(17)
		goto L51
	default:
		goto L53
	case 3:
		goto L52
	}
L50:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v166)+24))
	v171 = v166
	v173 = v168
	goto L48
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132)+24)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v132)+80))
	if v166 != v77 {
		v132 = v166
		goto L49
	} else {
		goto L61
	}
L52:
	;
	v164 = int32(16)
	goto L51
L53:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L14
	} else {
		goto L54
	}
L54:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v132)+24))
	if base.Ui32(v144) <= base.Ui32(int32(19)) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v151
	F_errmsg_internal(m, int32(_a_F_RollbackToSavepoint_6), v9+int32(32))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L14
	} else {
		goto L59
	}
L56:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v144<<(uint(int32(2))%32))+uint32(_c_F_RollbackToSavepoint[2])))
	v151 = v149
	goto L58
L57:
	;
	v151 = int32(_a_F_RollbackToSavepoint_7)
	goto L58
L58:
	;
	goto L55
L59:
	;
	F_errfinish(m, int32(_a_F_RollbackToSavepoint_3), int32(_a_F_RollbackToSavepoint_11), int32(_a_F_RollbackToSavepoint_5))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	goto L50
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+24)) = v202
	m.G0 = v9 + int32(112)
	return
L63:
	;
	v202 = int32(19)
	goto L62
L64:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v171)+24))
	if base.Ui32(v182) <= base.Ui32(int32(19)) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v189
	F_errmsg_internal(m, int32(_a_F_RollbackToSavepoint_6), v9+int32(16))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L14
	} else {
		goto L70
	}
L67:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182<<(uint(int32(2))%32))+uint32(_c_F_RollbackToSavepoint[2])))
	v189 = v187
	goto L69
L68:
	;
	v189 = int32(_a_F_RollbackToSavepoint_7)
	goto L69
L69:
	;
	goto L66
L70:
	;
	F_errfinish(m, int32(_a_F_RollbackToSavepoint_3), int32(_a_F_RollbackToSavepoint_12), int32(_a_F_RollbackToSavepoint_5))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L14
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L14
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(_a_F_RollbackToSavepoint_13), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L14
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_RollbackToSavepoint_3), int32(_a_F_RollbackToSavepoint_14), int32(_a_F_RollbackToSavepoint_5))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L14
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errcode(m, int32(16778371))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L14
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = l0
	F_errmsg(m, int32(_a_F_RollbackToSavepoint_9), v9-int32(-64))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L14
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_RollbackToSavepoint_3), int32(_a_F_RollbackToSavepoint_15), int32(_a_F_RollbackToSavepoint_5))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L14
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_errcode(m, int32(16778371))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L14
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = l0
	F_errmsg(m, int32(_a_F_RollbackToSavepoint_16), v9+int32(48))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L14
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_RollbackToSavepoint_3), int32(_a_F_RollbackToSavepoint_17), int32(_a_F_RollbackToSavepoint_5))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L14
	} else {
		goto L83
	}
L83:
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
									F_errmsg(m, int32(_a_F_coerce_to_common_type_0), v9)
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
											F_errfinish(m, int32(_a_F_coerce_to_common_type_1), int32(1590), int32(_a_F_coerce_to_common_type_2))
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
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
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
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 float64
	_ = v116
	var v120 float64
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v16 {
	case 0, 9, 10, 11, 12, 13, 32:
		v34 = l1
		v35 = v12 + int32(12)
		goto L1
	default:
		goto L2
	}
L1:
	;
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35))) = uint8(v37)
	v44 = v34
	goto L5
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v16)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	switch v22 {
	case 0, 9, 10, 11, 12, 13, 32:
		v34 = l1 + int32(1)
		v35 = v12 + int32(13)
		goto L1
	default:
		goto L3
	}
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v22)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
	switch v28 {
	case 0, 9, 10, 11, 12, 13, 32:
		v34 = l1 + int32(2)
		v35 = v12 + int32(14)
		goto L1
	default:
		goto L4
	}
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v28)
	v34 = l1 + int32(3)
	v35 = v12 + int32(15)
	goto L1
L5:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if base.B2i32(base.Ui32(v49-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v49 == int32(32)) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return v132
L7:
	;
	v44 = v44 + int32(1)
	goto L5
L8:
	;
	if v49 != 0 {
		v132 = v37
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L6
L10:
	;
	m.G0 = v12 + int32(16)
	goto L9
L11:
	;
	if l2&int32(251658240) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v63 = int32(_a_F_convert_to_base_unit_0)
	goto L14
L13:
	;
	v63 = int32(_a_F_convert_to_base_unit_1)
	goto L14
L14:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v64 == int32(0) {
		v132 = v37
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
	v78 = v63 + v68<<(uint(int32(4))%32)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if l2 != v79 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v132 = int32(0)
	goto L10
L18:
	;
	v125 = v68 + int32(1)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v125<<(uint(int32(4))%32)))))
	if v129 != 0 {
		v68 = v125
		goto L16
	} else {
		goto L31
	}
L19:
	;
	v82 = v12 + int32(12)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if base.B2i32(v85 == int32(0))|base.B2i32(v85 != v88) != 0 {
		v106 = v85
		v107 = v88
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v106-v107 != 0 {
		goto L18
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v91 = v82
	v92 = v78
	goto L23
L23:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if v96 == int32(0) {
		v106 = v96
		v107 = v95
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v106 = v96
	v107 = v95
	goto L21
L25:
	;
	v99 = int32(1)
	if v96 == v95 {
		v91 = v91 + v99
		v92 = v92 + v99
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v109 = *(*float64)(unsafe.Add(mBase, uint32(v78)+8))
	v110 = base.F64_mul(l0, v109)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+16)))
	if v111 == int32(0) {
		v120 = v110
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l3))) = v120
	v132 = int32(1)
	goto L10
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	if l2 != v114 {
		v120 = v110
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v116 = *(*float64)(unsafe.Add(mBase, uint32(v78)+24))
	v120 = base.F64_mul(v116, base.F64_nearest(base.F64_div(v110, v116)))
	goto L28
L31:
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	if l4 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if v89 != 0 {
		goto L25
	} else {
		goto L26
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L9
	} else {
		goto L22
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
		v88 = l3
		v89 = l4
		v91 = v12
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	base.MemoryCopy(m, v12, l3, v14)
	goto L8
L7:
	;
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	m.T0[v19].(func(*base.Module, int32, int32))(m, l0, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = F_pg_checksum_update(m, l1, v22, v23)
	mBase = m.M
	if v24 < int32(0) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v27 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v27
	v29 = l4 - v14
	if v29 == v27 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v36 = l3 + v14
	v37 = v29
	goto L13
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v37) < base.Ui32(v41) {
		v88 = v36
		v89 = v37
		v91 = v40
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L3
L15:
	;
	if v41 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	base.MemoryCopy(m, v40, v36, v41)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	m.T0[v46].(func(*base.Module, int32, int32))(m, l0, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v51 = F_pg_checksum_update(m, l1, v49, v50)
	mBase = m.M
	if v51 < int32(0) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v57 = v37 - v41
	if v57 != 0 {
		v36 = v36 + v41
		v37 = v57
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	F_errmsg_internal(m, int32(_a_F_push_to_sink_0), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_push_to_sink_1), int32(1980), int32(_a_F_push_to_sink_2))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	base.MemoryCopy(m, v91, v88, v89)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v93 + v89
	return
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
					F_errmsg(m, int32(_a_F_to_ascii_enc_0), v6)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_to_ascii_enc_1), int32(146), int32(_a_F_to_ascii_enc_2))
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14025(m, l0, int64(4), int64(16), int32(15))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_to_json(m *base.Module, l0 int32) int32 {
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v5 = m.G0
	v7 = v5 - int32(16)
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
					F_errmsg(m, int32(_a_F_to_json_0), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_to_json_1), int32(749), int32(_a_F_to_json_2))
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
			F_json_categorize_type(m, v12, int32(0), v7+int32(12), v7+int32(8))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v44 = F_makeStringInfo(m)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_datum_to_json_internal(m, v9, int32(0), v44, v42, v41, int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
						v51 = F_cstring_to_text_with_len(m, v49, v50)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(16)
							return v51
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
					F_errmsg(m, int32(_a_F_to_jsonb_0), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_to_jsonb_1), int32(1098), int32(_a_F_to_jsonb_2))
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
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14027(m, l0, int32(1478))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
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
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_to_regtypemod[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v17
			v20 = *(*int64)(unsafe.Add(mBase, _c_F_to_regtypemod[1]))
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v121 int32
	_ = v121
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v17 = int32(0)
			v28 = F_do_to_timestamp(m, v9, v14, v16, v17, v6+int32(24), v6+int32(12), v6+int32(16), v6+int32(8), v17, v17)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)))
				if v30 == int32(1) {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
					v49 = v33
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_to_timestamp[0]))
					v39 = m.G0
					v40 = int32(16)
					v41 = v39 - v40
					m.G0 = v41
					v45 = F_DetermineTimeZoneOffsetInternal(m, v6+int32(24), v37, v41+int32(8))
					mBase = m.M
					m.G0 = v41 + v40
					v49 = v45
				}
				*(*int32)(unsafe.Add(mBase, uint32(v6)+68)) = v49
				v52 = v6 + int32(24)
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v55 = v6 + int32(68)
				v57 = v6 + int32(72)
				v64 = m.G0
				v66 = v64 - int32(16)
				m.G0 = v66
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
				if v68 <= int32(-4713) {
					if v68 != int32(-4713) {
						*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
						v145 = int32(-1)
					} else {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
						if int32(10) < v73 {
							v84 = v73
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
							v86 = F_date2j(m, v68, v84, v85)
							mBase = m.M
							v89 = base.I64_extend_i32_s(v86 - int32(_a_F_to_timestamp_0))
							v90 = int64(63)
							F___multi3(m, v66, v89, v89>>(uint(v90)%64), int64(86400000000), int64(0))
							mBase = m.M
							v95 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
							if v95 != v96>>(uint(v90)%64) {
								*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
								v145 = int32(-1)
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
								v104 = int32(60)
								v113 = base.I64_extend_i32_s(v53) + base.I64_extend_i32_s(v101+(v102+v103*v104)*v104)*int64(1000000)
								v114 = v96 + v113
								*(*int64)(unsafe.Add(mBase, uint32(v57))) = v114
								if base.B2i32(v113 < int64(0))^base.B2i32(v114 < v96) != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
									v145 = int32(-1)
								} else {
									if v55 != 0 {
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
										v126 = base.I64_extend_i32_s(int32(0)-v121)*int64(-1000000) + v114
										*(*int64)(unsafe.Add(mBase, uint32(v57))) = v126
										v128 = v126
									} else {
										v128 = v114
									}
									if base.Ui64(v128+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
										v145 = int32(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
										v145 = int32(-1)
									}
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
							v145 = int32(-1)
						}
					}
				} else {
					if v68 <= int32(_a_F_to_timestamp_1) {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
						v84 = v78
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
						v86 = F_date2j(m, v68, v84, v85)
						mBase = m.M
						v89 = base.I64_extend_i32_s(v86 - int32(_a_F_to_timestamp_0))
						v90 = int64(63)
						F___multi3(m, v66, v89, v89>>(uint(v90)%64), int64(86400000000), int64(0))
						mBase = m.M
						v95 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
						v96 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
						if v95 != v96>>(uint(v90)%64) {
							*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
							v145 = int32(-1)
						} else {
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
							v104 = int32(60)
							v113 = base.I64_extend_i32_s(v53) + base.I64_extend_i32_s(v101+(v102+v103*v104)*v104)*int64(1000000)
							v114 = v96 + v113
							*(*int64)(unsafe.Add(mBase, uint32(v57))) = v114
							if base.B2i32(v113 < int64(0))^base.B2i32(v114 < v96) != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
								v145 = int32(-1)
							} else {
								if v55 != 0 {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
									v126 = base.I64_extend_i32_s(int32(0)-v121)*int64(-1000000) + v114
									*(*int64)(unsafe.Add(mBase, uint32(v57))) = v126
									v128 = v126
								} else {
									v128 = v114
								}
								if base.Ui64(v128+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
									v145 = int32(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
									v145 = int32(-1)
								}
							}
						}
					} else {
						if v68 != int32(_a_F_to_timestamp_2) {
							*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
							v145 = int32(-1)
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
							if int32(5) < v81 {
								*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
								v145 = int32(-1)
							} else {
								v84 = v81
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
								v86 = F_date2j(m, v68, v84, v85)
								mBase = m.M
								v89 = base.I64_extend_i32_s(v86 - int32(_a_F_to_timestamp_0))
								v90 = int64(63)
								F___multi3(m, v66, v89, v89>>(uint(v90)%64), int64(86400000000), int64(0))
								mBase = m.M
								v95 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
								v96 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
								if v95 != v96>>(uint(v90)%64) {
									*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
									v145 = int32(-1)
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
									v104 = int32(60)
									v113 = base.I64_extend_i32_s(v53) + base.I64_extend_i32_s(v101+(v102+v103*v104)*v104)*int64(1000000)
									v114 = v96 + v113
									*(*int64)(unsafe.Add(mBase, uint32(v57))) = v114
									if base.B2i32(v113 < int64(0))^base.B2i32(v114 < v96) != 0 {
										*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
										v145 = int32(-1)
									} else {
										if v55 != 0 {
											v121 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
											v126 = base.I64_extend_i32_s(int32(0)-v121)*int64(-1000000) + v114
											*(*int64)(unsafe.Add(mBase, uint32(v57))) = v126
											v128 = v126
										} else {
											v128 = v114
										}
										if base.Ui64(v128+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
											v145 = int32(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
											v145 = int32(-1)
										}
									}
								}
							}
						}
					}
				}
				m.G0 = v66 + int32(16)
				if v145 == int32(0) {
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
					if v151 != 0 {
						F_AdjustTimestampForTypmod(m, v57, v151, int32(0))
						mBase = m.M
						v154 = m.ExcPending
						if v154 != 0 {
							return int32(0)
						} else {
							v155 = *(*int64)(unsafe.Add(mBase, uint32(v6)+72))
							v156 = F_Int64GetDatum(m, v155)
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return int32(0)
							} else {
								m.G0 = v6 + int32(80)
								return v156
							}
						}
					} else {
						v155 = *(*int64)(unsafe.Add(mBase, uint32(v6)+72))
						v156 = F_Int64GetDatum(m, v155)
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return int32(0)
						} else {
							m.G0 = v6 + int32(80)
							return v156
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_to_timestamp_3), int32(0))
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_to_timestamp_4), int32(_a_F_to_timestamp_5), int32(_a_F_to_timestamp_6))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
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
