package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonbIteratorInit(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_iteratorFromContainer(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_JsonbToJsonbValue(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(18)
	v5 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l0 + v5
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(base.Ui32(v8)>>(uint(int32(2))%32)) - v5
	return
}
func F_add_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l3 != 0 {
		if l1 != 0 {
			v12 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v12
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v12
			v26 = int32(0)
			v27 = v12
			F_datum_to_jsonb_internal(m, l0, l1, l2, v26, v27, l4)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		} else {
			F_json_categorize_type(m, l3, int32(1), v10+int32(12), v10+int32(8))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				v26 = v24
				v27 = v25
				F_datum_to_jsonb_internal(m, l0, l1, l2, v26, v27, l4)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					m.G0 = v10 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				F_errmsg(m, int32(352917), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_errfinish(m, int32(476670), int32(1025), int32(479327))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
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
func F_fillJsonbValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v119 int32
	_ = v119
	v9 = l0 + int32(4)
	v12 = v9 + l1<<(uint(int32(2))%32)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	switch int32(base.Ui32(v13)>>(uint(int32(28))%32)) & int32(7) {
	case 0:
		goto L5
	case 1:
		goto L4
	case 2:
		goto L2
	case 3:
		goto L3
	case 4:
		goto L6
	default:
		goto L1
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(18)
	v78 = (l3 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = l2 + v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v81 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)) = uint8(v69)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(3)
	return
L3:
	;
	v65 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)) = uint8(v65)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(3)
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = l2 + (l3+int32(3))&int32(-4)
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = l2 + l3
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v24 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	return
L7:
	;
	v29 = l1
	v30 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v24 & int32(268435455)
	return
L10:
	;
	v36 = v29 - int32(1)
	if int32(0) <= v36 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v24&int32(268435455) - v48
	return
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9+v36<<(uint(int32(2))%32))))
	v45 = v42&int32(268435455) + v30
	if int32(0) <= v42 {
		v29 = v36
		v30 = v45
		goto L10
	} else {
		goto L15
	}
L13:
	;
	v48 = v30
	goto L14
L14:
	;
	goto L11
L15:
	;
	v48 = v45
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v119 + (l3 - v78)
	return
L17:
	;
	v86 = l1
	v87 = int32(0)
	goto L20
L18:
	;
	goto L19
L19:
	;
	v119 = v81 & int32(268435455)
	goto L16
L20:
	;
	v93 = v86 - int32(1)
	if int32(0) <= v93 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v119 = v81&int32(268435455) - v105
	goto L16
L22:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v9+v93<<(uint(int32(2))%32))))
	v102 = v99&int32(268435455) + v87
	if int32(0) <= v99 {
		v86 = v93
		v87 = v102
		goto L20
	} else {
		goto L25
	}
L23:
	;
	v105 = v87
	goto L24
L24:
	;
	goto L21
L25:
	;
	v105 = v102
	goto L24
}
func F_jsonb_agg_transfn_worker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v13 = v8 + int32(-4)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == v3 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L23
	} else {
		goto L70
	}
L2:
	;
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	v43 = v40
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v35
	v40 = v36
	goto L3
L5:
	;
	v32 = int32(0)
	if v13 == v32 {
		v40 = v32
		goto L3
	} else {
		goto L15
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	switch v18 - int32(429) {
	case 0:
		goto L8
	case 1:
		goto L7
	default:
		goto L5
	}
L7:
	;
	if v13 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	if v13 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v43 = int32(1)
	goto L2
L10:
	;
	goto L11
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v35 = v25
	v36 = int32(1)
	goto L4
L12:
	;
	v43 = int32(2)
	goto L2
L13:
	;
	goto L14
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+368))
	v35 = v30
	v36 = int32(2)
	goto L4
L15:
	;
	v35 = v32
	v36 = v3
	goto L4
L16:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v44 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L23
	} else {
		goto L67
	}
L19:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if l1 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = F_get_fn_expr_argtype(m, v47, int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v83 = v82
	v85 = v81
	goto L19
L23:
	;
	return int32(0)
L24:
	;
	if v49 == int32(0) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v55 = int32(4442576)
	v56 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v58
	v61 = F_palloc(m, int32(20))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v64 = F_palloc0(m, int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v64
	v69 = F_pushJsonbValue(m, v64, int32(4), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v69
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v56
	F_json_categorize_type(m, v49, int32(1), v61+int32(12), v61+int32(16))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v83 = v64
	v85 = v61
	goto L19
L30:
	;
	m.G0 = v10 - int32(-64)
	return v85
L31:
	;
	v109 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v109
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	F_datum_to_jsonb_internal(m, v108, v106, v8+int32(-24), v115, v116, int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L23
	} else {
		goto L39
	}
L32:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v106 = int32(0)
	v108 = v105
	goto L31
L33:
	;
	v90 = int32(1)
	if v87&v90 == int32(0) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v96 = int32(1)
	v98 = v87 & v96
	if v98 != 0 {
		goto L30
	} else {
		goto L37
	}
L36:
	;
	v106 = v90
	v108 = int32(0)
	goto L31
L37:
	;
	if v98 == int32(0) {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	v106 = v96
	v108 = int32(0)
	goto L31
L39:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v121 = F_JsonbValueToJsonb(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v123 = int32(4442576)
	v124 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v126
	v130 = F_JsonbIteratorInit(m, v121+int32(4))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L23
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v130
	v134 = int32(0)
	goto L43
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v124
	goto L30
L43:
	;
	v146 = F_JsonbIteratorNext(m, v8+int32(-28), v8+int32(-48), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L23
	} else {
		goto L50
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L23
	} else {
		goto L64
	}
L45:
	;
	goto L44
L46:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	switch v168 - int32(1) {
	case 0:
		goto L58
	case 1:
		goto L57
	default:
		goto L56
	}
L47:
	;
	v165 = F_pushJsonbValue(m, v83, v146, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L23
	} else {
		goto L55
	}
L48:
	;
	v155 = int32(1)
	if v134&v155 != 0 {
		v134 = v155
		goto L43
	} else {
		goto L53
	}
L49:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
	if v149 != 0 {
		v134 = int32(1)
		goto L43
	} else {
		goto L51
	}
L50:
	;
	switch v146 {
	case 0:
		goto L42
	case 1, 2, 3:
		goto L46
	case 4:
		goto L49
	case 5:
		goto L48
	case 6, 7:
		goto L47
	default:
		goto L45
	}
L51:
	;
	v152 = F_pushJsonbValue(m, v83, int32(4), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L23
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v152
	goto L43
L53:
	;
	v158 = int32(0)
	v161 = F_pushJsonbValue(m, v83, int32(5), v158)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L23
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v161
	v134 = v158
	goto L43
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v165
	goto L43
L56:
	;
	v196 = F_pushJsonbValue(m, v83, v146, v8+int32(-48))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L23
	} else {
		goto L63
	}
L57:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v188 = F_DirectFunctionCall1Coll(m, int32(1331), int32(0), v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L23
	} else {
		goto L61
	}
L58:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v174 = F_palloc(m, v171+int32(1))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v182 = F_pg_snprintf(m, v174, v178+int32(1), int32(195784), v10)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v174
	goto L56
L61:
	;
	v190 = F_pg_detoast_datum(m, v188)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L23
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v190
	goto L56
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v196
	goto L43
L64:
	;
	F_errmsg_internal(m, int32(351107), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L23
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(476670), int32(1612), int32(209367))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L23
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errmsg_internal(m, int32(58958), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L23
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(476670), int32(1518), int32(209367))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L23
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L23
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(352917), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L23
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(476670), int32(1530), int32(209367))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L23
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_contains(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			if (v17^v18)&int32(536870912) == int32(0) {
				v26 = F_JsonbIteratorInit(m, v10+int32(4))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v26
					v31 = F_JsonbIteratorInit(m, v15+int32(4))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v31
						v38 = F_JsonbDeepContains(m, v7+int32(12), v7+int32(8))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = v38
							m.G0 = v7 + int32(16)
							return v40
						}
					}
				}
			} else {
				v40 = int32(0)
				m.G0 = v7 + int32(16)
				return v40
			}
		}
	}
}
func F_jsonb_delete(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
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
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(1)
	v22 = v19 + v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v25 = v23 & v21
	if v23 == v21 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v54 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v56&int32(268435456) == v54 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v28 = int32(4)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v30&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v39 = v28
	goto L10
L9:
	;
	v39 = base.B2i32(v30 == int32(18)) << (uint(v28) % 32)
	goto L10
L10:
	;
	if v30 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v42 = v28
	goto L13
L12:
	;
	v42 = v39
	goto L13
L13:
	;
	v53 = v42
	goto L4
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	if v56&int32(268435455) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L66
	}
L18:
	;
	v65 = F_JsonbIteratorInit(m, v14+int32(4))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v206 = v14
	goto L20
L20:
	;
	m.G0 = v11 + int32(32)
	return v206
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
	v73 = F_JsonbIteratorNext(m, v11+int32(24), v11+int32(4), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v202 = F_JsonbValueToJsonb(m, v199)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L65
	}
L23:
	;
	if v73 == int32(0) {
		v199 = v2
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v25 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v79 = v22
	goto L27
L26:
	;
	v79 = v19 + int32(4)
	goto L27
L27:
	;
	v82 = v73
	v85 = v2
	goto L28
L28:
	;
	v89 = base.B2i32(v82 != int32(1))
	if v89&base.B2i32(v82 != int32(3)) != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v199 = v185
	goto L22
L30:
	;
	v179 = int32(4)
	if base.Ui32(v82) < base.Ui32(v179) {
		goto L59
	} else {
		goto L60
	}
L31:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v93 != int32(1) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v53 != v96 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if base.Ui32(int32(4)) <= base.Ui32(v53) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	if v160 != 0 {
		goto L30
	} else {
		goto L52
	}
L35:
	;
	v160 = int32(0)
	goto L34
L36:
	;
	v134 = v129
	v135 = v130
	v136 = v131
	goto L46
L37:
	;
	if (v79|v98)&int32(3) != 0 {
		v129 = v79
		v130 = v98
		v131 = v53
		goto L36
	} else {
		goto L40
	}
L38:
	;
	v122 = v79
	v123 = v98
	v124 = v53
	goto L39
L39:
	;
	if v124 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v106 = v79
	v107 = v98
	v108 = v53
	goto L41
L41:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v111 != v112 {
		v129 = v106
		v130 = v107
		v131 = v108
		goto L36
	} else {
		goto L43
	}
L42:
	;
	v122 = v117
	v123 = v115
	v124 = v119
	goto L39
L43:
	;
	v114 = int32(4)
	v115 = v107 + v114
	v117 = v106 + v114
	v119 = v108 - v114
	if base.Ui32(int32(3)) < base.Ui32(v119) {
		v106 = v117
		v107 = v115
		v108 = v119
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v129 = v122
	v130 = v123
	v131 = v124
	goto L36
L46:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v139 == v140 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v160 = v139 - v140
	goto L34
L48:
	;
	v142 = int32(1)
	v147 = v136 - v142
	if v147 != 0 {
		v134 = v134 + v142
		v135 = v135 + v142
		v136 = v147
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L35
L52:
	;
	if v89 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v168 = F_JsonbIteratorNext(m, v11+int32(24), v11+int32(4), int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v175 = F_JsonbIteratorNext(m, v11+int32(24), v11+int32(4), int32(1))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	if v175 != 0 {
		v82 = v175
		goto L28
	} else {
		goto L58
	}
L58:
	;
	v199 = v85
	goto L22
L59:
	;
	v184 = v11 + v179
	goto L61
L60:
	;
	v184 = int32(0)
	goto L61
L61:
	;
	v185 = F_pushJsonbValue(m, v11+int32(28), v82, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v192 = F_JsonbIteratorNext(m, v11+int32(24), v11+int32(4), int32(1))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v192 != 0 {
		v82 = v192
		v85 = v185
		goto L28
	} else {
		goto L64
	}
L64:
	;
	goto L29
L65:
	;
	v206 = v202
	goto L20
L66:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(218369), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(471558), int32(4680), int32(334174))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_each_text(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_each_worker_jsonb(m, l0, int32(60370), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_jsonb_exists_any(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v112 int32
	_ = v112
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_deconstruct_array_builtin(m, v16, int32(25), v8+int32(28), v8+int32(24), v8+int32(20))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if int32(0) < v27 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	m.G0 = v8 + int32(32)
	return v112
L6:
	;
	v112 = int32(1)
	goto L5
L7:
	;
	v33 = int32(0)
	v35 = v27
	goto L10
L8:
	;
	goto L9
L9:
	;
	v112 = int32(0)
	goto L5
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v33))))
	if v40 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v44 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v44
	v48 = v43 + v33<<(uint(int32(2))%32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v52&v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v95 = v35
	goto L14
L14:
	;
	v98 = v33 + int32(1)
	if v98 < v95 {
		v33 = v98
		v35 = v95
		goto L10
	} else {
		goto L31
	}
L15:
	;
	v55 = v44
	goto L17
L16:
	;
	v55 = int32(4)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v49 + v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v59 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v89
	v92 = F_findJsonbValueFromContainer(m, v11+int32(4), int32(1610612736), v8)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L29
	}
L19:
	;
	v62 = int32(4)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v64&int32(254) == int32(2) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v77 = int32(1)
	if v59&v77 != 0 {
		v89 = int32(base.Ui32(v59)>>(uint(v77)%32)) - v77
		goto L18
	} else {
		goto L28
	}
L22:
	;
	v73 = v62
	goto L24
L23:
	;
	v73 = base.B2i32(v64 == int32(18)) << (uint(v62) % 32)
	goto L24
L24:
	;
	if v64 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v76 = v62
	goto L27
L26:
	;
	v76 = v73
	goto L27
L27:
	;
	v89 = v76
	goto L18
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v89 = int32(base.Ui32(v83)>>(uint(int32(2))%32)) - int32(4)
	goto L18
L29:
	;
	if v92 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v95 = v94
	goto L14
L31:
	;
	goto L11
}
func F_jsonb_extract_path(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_get_jsonb_path_all(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v14 = l0 + int32(28)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = F_compareJsonbContainers(m, v7+int32(4), v16+int32(4))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v22 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						if v26 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								return base.B2i32(int32(0) < v20)
							}
						} else {
							return base.B2i32(int32(0) < v20)
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if v26 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							return base.B2i32(int32(0) < v20)
						}
					} else {
						return base.B2i32(int32(0) < v20)
					}
				}
			}
		}
	}
}
func F_jsonb_in_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(1)
	if l1&int32(3) == int32(0) {
		v34 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v67
	if base.Ui32(int32(268435456)) <= base.Ui32(v67) {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v67 = v59 - l1
	goto L1
L3:
	;
	v38 = v34
	goto L12
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v18 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v67 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v23 = l1
	goto L8
L8:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v59 = v27
	goto L2
L10:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v53 = v38
	goto L15
L14:
	;
	goto L13
L15:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v59 = v53
	goto L2
L17:
	;
	goto L16
L18:
	;
	m.G0 = v7 + int32(32)
	return v105
L19:
	;
	v71 = int32(23)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v73 = F_errsave_start(m, v72)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l1
	v100 = F_pushJsonbValue(m, l0, int32(1), v7+int32(12))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L22
	} else {
		goto L29
	}
L22:
	;
	return int32(0)
L23:
	;
	if v73 == int32(0) {
		v105 = v71
		goto L18
	} else {
		goto L24
	}
L24:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(313898), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(268435455)
	F_errdetail(m, int32(556922), v7)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	F_errsave_finish(m, v72, int32(476670), int32(284), int32(268424))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v105 = v71
	goto L18
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v100
	v105 = int32(0)
	goto L18
}
func F_jsonb_numeric(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v19 = F_JsonbExtractScalar(m, v11+int32(4), v8+int32(12))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			if v19 != 0 {
				switch v21 {
				case 0:
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v22 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v26 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
							v39 = int32(0)
							m.G0 = v8 + int32(32)
							return v39
						}
					} else {
						v26 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
						v39 = int32(0)
						m.G0 = v8 + int32(32)
						return v39
					}
				default:
					F_cannotCastJsonbValue(m, v21, int32(468051))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 2:
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					v33 = F_pg_detoast_datum_copy(m, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v11 == v35 {
							v39 = v33
							m.G0 = v8 + int32(32)
							return v39
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = v33
								m.G0 = v8 + int32(32)
								return v39
							}
						}
					}
				}
			} else {
				F_cannotCastJsonbValue(m, v21, int32(468051))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
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
func F_jsonb_object(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v17 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v17
	v25 = F_pushJsonbValue(m, v7+int32(-32), int32(6), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	switch v16 {
	case 0:
		goto L5
	case 1:
		goto L9
	case 2:
		goto L8
	default:
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L77
	}
L5:
	;
	v295 = F_pushJsonbValue(m, v7+int32(-32), int32(7), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L75
	}
L6:
	;
	F_deconstruct_array_builtin(m, v12, int32(25), v7+int32(-4), v7+int32(-8), v7+int32(-12))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L24
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L20
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v48 == int32(2) {
		goto L6
	} else {
		goto L15
	}
L9:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)))
	if v27&int32(1) == int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	F_errmsg(m, int32(115041), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(476670), int32(1304), int32(102894))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errmsg(m, int32(138249), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(476670), int32(1311), int32(102894))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(110307), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(476670), int32(1317), int32(102894))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
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
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	v93 = int32(2)
	v94 = base.I32_div_s(v92, v93)
	if v93 <= v92 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v98 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	F_pfree(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L73
	}
L28:
	;
	v104 = int32(1)
	v105 = v98 << (uint(v104) % 32)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105+v106))))
	if v108 == v104 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	goto L27
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v105<<(uint(int32(2))%32))))
	v116 = F_text_to_cstring(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v116&int32(3) == int32(0) {
		v141 = v116
		goto L34
	} else {
		goto L35
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v174
	v177 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v177
	v184 = F_pushJsonbValue(m, v7+int32(-32), v177, v7+int32(-52))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L49
	}
L33:
	;
	v174 = v166 - v116
	goto L32
L34:
	;
	v145 = v141
	goto L43
L35:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v125 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v174 = int32(0)
	goto L32
L37:
	;
	goto L38
L38:
	;
	v130 = v116
	goto L39
L39:
	;
	v134 = v130 + int32(1)
	if v134&int32(3) == int32(0) {
		v141 = v134
		goto L34
	} else {
		goto L41
	}
L40:
	;
	v166 = v134
	goto L33
L41:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v139 != 0 {
		v130 = v134
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v154 = int32(-2139062144)
	if (int32(16843008)-v151|v151)&v154 == v154 {
		v145 = v145 + int32(4)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v160 = v145
	goto L46
L45:
	;
	goto L44
L46:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v164 != 0 {
		v160 = v160 + int32(1)
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v166 = v160
	goto L33
L48:
	;
	goto L47
L49:
	;
	v188 = v105 | int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v189))))
	if v191 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v261 = int32(0)
	goto L52
L51:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+v188<<(uint(int32(2))%32))))
	v197 = F_text_to_cstring(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v261
	v268 = F_pushJsonbValue(m, v7+int32(-32), int32(2), v7+int32(-52))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L71
	}
L53:
	;
	if v197&int32(3) == int32(0) {
		v222 = v197
		goto L56
	} else {
		goto L57
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v255
	v261 = int32(1)
	goto L52
L55:
	;
	v255 = v247 - v197
	goto L54
L56:
	;
	v226 = v222
	goto L65
L57:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if v206 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v255 = int32(0)
	goto L54
L59:
	;
	goto L60
L60:
	;
	v211 = v197
	goto L61
L61:
	;
	v215 = v211 + int32(1)
	if v215&int32(3) == int32(0) {
		v222 = v215
		goto L56
	} else {
		goto L63
	}
L62:
	;
	v247 = v215
	goto L55
L63:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v220 != 0 {
		v211 = v215
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v235 = int32(-2139062144)
	if (int32(16843008)-v232|v232)&v235 == v235 {
		v226 = v226 + int32(4)
		goto L65
	} else {
		goto L67
	}
L66:
	;
	v241 = v226
	goto L68
L67:
	;
	goto L66
L68:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	if v245 != 0 {
		v241 = v241 + int32(1)
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v247 = v241
	goto L55
L70:
	;
	goto L69
L71:
	;
	v271 = v98 + int32(1)
	if v271 != v94 {
		v98 = v271
		goto L28
	} else {
		goto L72
	}
L72:
	;
	goto L29
L73:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	F_pfree(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L5
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v295
	v298 = F_JsonbValueToJsonb(m, v295)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	m.G0 = v9 - int32(-64)
	return v298
L77:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errmsg(m, int32(19978), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(476670), int32(1333), int32(102894))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_ops__add_path_item(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
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
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v13 - int32(21) {
	case 0, 1, 2, 3:
		v318 = v3
		v323 = F_palloc(m, int32(12))
		mBase = m.M
		v324 = m.ExcPending
		if v324 != 0 {
			return int32(0)
		} else {
			v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v323)+4)) = v318
			*(*int32)(unsafe.Add(mBase, uint32(v323)+8)) = v325
			v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v323))) = v328
			v332 = v323
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v332
			v340 = int32(1)
			m.G0 = v11 + int32(32)
			return v340
		}
	case 4:
		v17 = v11 + int32(16)
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v18
		} else {
		}
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
		if int32(126) <= v22 {
			v30 = v22 - int32(1636608432)
			if v20&int32(3) != 0 {
				if base.Ui32(int32(11)) < base.Ui32(v22) {
					v139 = v20
					v140 = v22
					v141 = v30
					v142 = v30
					v143 = v30
					for {
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
						v146 = v145 + v142
						v147 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
						v149 = *(*int32)(unsafe.Add(mBase, uint32(v139)+8))
						v150 = v149 + v143
						v152 = int32(4)
						v154 = v147 + v141 - v150 ^ base.I32_rotl(v150, v152)
						v158 = v146 - v154 ^ base.I32_rotl(v154, int32(6))
						v159 = v150 + v146
						v160 = v154 + v159
						v161 = v158 + v160
						v165 = v159 - v158 ^ base.I32_rotl(v158, int32(8))
						v169 = v160 - v165 ^ base.I32_rotl(v165, int32(16))
						v173 = v161 - v169 ^ base.I32_rotl(v169, int32(19))
						v174 = v165 + v161
						v175 = v169 + v174
						v176 = v173 + v175
						v180 = v174 - v173 ^ base.I32_rotl(v173, v152)
						v181 = int32(12)
						v182 = v139 + v181
						v184 = v140 - v181
						if base.Ui32(int32(11)) < base.Ui32(v184) {
							v139 = v182
							v140 = v184
							v141 = v175
							v142 = v176
							v143 = v180
							continue
						} else {
							break
						}
						break
					}
					v187 = v182
					v188 = v184
					v189 = v175
					v190 = v176
					v191 = v180
				} else {
					v187 = v20
					v188 = v22
					v189 = v30
					v190 = v30
					v191 = v30
				}
				switch v188 - int32(1) {
				case 0:
					v250 = v189
					v251 = v190
					v252 = v191
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
					v257 = v250 + v253
					v258 = v251
					v259 = v252
				case 1:
					v243 = v189
					v244 = v190
					v245 = v191
					v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
					v250 = v246<<(uint(int32(8))%32) + v243
					v251 = v244
					v252 = v245
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
					v257 = v250 + v253
					v258 = v251
					v259 = v252
				case 2:
					v236 = v189
					v237 = v190
					v238 = v191
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
					v243 = v239<<(uint(int32(16))%32) + v236
					v244 = v237
					v245 = v238
					v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
					v250 = v246<<(uint(int32(8))%32) + v243
					v251 = v244
					v252 = v245
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
					v257 = v250 + v253
					v258 = v251
					v259 = v252
				case 3:
					v230 = v190
					v231 = v191
					v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+3)))
					v236 = v232<<(uint(int32(24))%32) + v189
					v237 = v230
					v238 = v231
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
					v243 = v239<<(uint(int32(16))%32) + v236
					v244 = v237
					v245 = v238
					v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
					v250 = v246<<(uint(int32(8))%32) + v243
					v251 = v244
					v252 = v245
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
					v257 = v250 + v253
					v258 = v251
					v259 = v252
				case 4:
					v226 = v190
					v227 = v191
					v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
					v230 = v226 + v228
					v231 = v227
					v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+3)))
					v236 = v232<<(uint(int32(24))%32) + v189
					v237 = v230
					v238 = v231
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
					v243 = v239<<(uint(int32(16))%32) + v236
					v244 = v237
					v245 = v238
					v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
					v250 = v246<<(uint(int32(8))%32) + v243
					v251 = v244
					v252 = v245
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
					v257 = v250 + v253
					v258 = v251
					v259 = v252
				case 5:
					v220 = v190
					v221 = v191
					v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
					v226 = v222<<(uint(int32(8))%32) + v220
					v227 = v221
					v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
					v230 = v226 + v228
					v231 = v227
					v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+3)))
					v236 = v232<<(uint(int32(24))%32) + v189
					v237 = v230
					v238 = v231
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
					v243 = v239<<(uint(int32(16))%32) + v236
					v244 = v237
					v245 = v238
					v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
					v250 = v246<<(uint(int32(8))%32) + v243
					v251 = v244
					v252 = v245
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
					v257 = v250 + v253
					v258 = v251
					v259 = v252
				case 6:
					v214 = v190
					v215 = v191
					v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
					v220 = v216<<(uint(int32(16))%32) + v214
					v221 = v215
					v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
					v226 = v222<<(uint(int32(8))%32) + v220
					v227 = v221
					v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
					v230 = v226 + v228
					v231 = v227
					v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+3)))
					v236 = v232<<(uint(int32(24))%32) + v189
					v237 = v230
					v238 = v231
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
					v243 = v239<<(uint(int32(16))%32) + v236
					v244 = v237
					v245 = v238
					v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
					v250 = v246<<(uint(int32(8))%32) + v243
					v251 = v244
					v252 = v245
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
					v257 = v250 + v253
					v258 = v251
					v259 = v252
				case 7:
					v209 = v191
					v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+7)))
					v214 = v210<<(uint(int32(24))%32) + v190
					v215 = v209
					v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
					v220 = v216<<(uint(int32(16))%32) + v214
					v221 = v215
					v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
					v226 = v222<<(uint(int32(8))%32) + v220
					v227 = v221
					v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
					v230 = v226 + v228
					v231 = v227
					v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+3)))
					v236 = v232<<(uint(int32(24))%32) + v189
					v237 = v230
					v238 = v231
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
					v243 = v239<<(uint(int32(16))%32) + v236
					v244 = v237
					v245 = v238
					v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
					v250 = v246<<(uint(int32(8))%32) + v243
					v251 = v244
					v252 = v245
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
					v257 = v250 + v253
					v258 = v251
					v259 = v252
				case 8:
					v204 = v191
					v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
					v209 = v205<<(uint(int32(8))%32) + v204
					v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+7)))
					v214 = v210<<(uint(int32(24))%32) + v190
					v215 = v209
					v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
					v220 = v216<<(uint(int32(16))%32) + v214
					v221 = v215
					v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
					v226 = v222<<(uint(int32(8))%32) + v220
					v227 = v221
					v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
					v230 = v226 + v228
					v231 = v227
					v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+3)))
					v236 = v232<<(uint(int32(24))%32) + v189
					v237 = v230
					v238 = v231
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
					v243 = v239<<(uint(int32(16))%32) + v236
					v244 = v237
					v245 = v238
					v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
					v250 = v246<<(uint(int32(8))%32) + v243
					v251 = v244
					v252 = v245
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
					v257 = v250 + v253
					v258 = v251
					v259 = v252
				case 9:
					v199 = v191
					v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
					v204 = v200<<(uint(int32(16))%32) + v199
					v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
					v209 = v205<<(uint(int32(8))%32) + v204
					v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+7)))
					v214 = v210<<(uint(int32(24))%32) + v190
					v215 = v209
					v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
					v220 = v216<<(uint(int32(16))%32) + v214
					v221 = v215
					v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
					v226 = v222<<(uint(int32(8))%32) + v220
					v227 = v221
					v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
					v230 = v226 + v228
					v231 = v227
					v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+3)))
					v236 = v232<<(uint(int32(24))%32) + v189
					v237 = v230
					v238 = v231
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
					v243 = v239<<(uint(int32(16))%32) + v236
					v244 = v237
					v245 = v238
					v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
					v250 = v246<<(uint(int32(8))%32) + v243
					v251 = v244
					v252 = v245
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
					v257 = v250 + v253
					v258 = v251
					v259 = v252
				case 10:
					v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+10)))
					v199 = v195<<(uint(int32(24))%32) + v191
					v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+9)))
					v204 = v200<<(uint(int32(16))%32) + v199
					v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+8)))
					v209 = v205<<(uint(int32(8))%32) + v204
					v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+7)))
					v214 = v210<<(uint(int32(24))%32) + v190
					v215 = v209
					v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
					v220 = v216<<(uint(int32(16))%32) + v214
					v221 = v215
					v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+5)))
					v226 = v222<<(uint(int32(8))%32) + v220
					v227 = v221
					v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+4)))
					v230 = v226 + v228
					v231 = v227
					v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+3)))
					v236 = v232<<(uint(int32(24))%32) + v189
					v237 = v230
					v238 = v231
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+2)))
					v243 = v239<<(uint(int32(16))%32) + v236
					v244 = v237
					v245 = v238
					v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
					v250 = v246<<(uint(int32(8))%32) + v243
					v251 = v244
					v252 = v245
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
					v257 = v250 + v253
					v258 = v251
					v259 = v252
				default:
					v257 = v189
					v258 = v190
					v259 = v191
				}
			} else {
				if base.Ui32(v22) < base.Ui32(int32(12)) {
					v85 = v20
					v86 = v22
					v87 = v30
					v88 = v30
					v89 = v30
				} else {
					v37 = v20
					v38 = v22
					v39 = v30
					v40 = v30
					v41 = v30
					for {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
						v44 = v43 + v40
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
						v48 = v47 + v41
						v50 = int32(4)
						v52 = v45 + v39 - v48 ^ base.I32_rotl(v48, v50)
						v56 = v44 - v52 ^ base.I32_rotl(v52, int32(6))
						v57 = v48 + v44
						v58 = v52 + v57
						v59 = v56 + v58
						v63 = v57 - v56 ^ base.I32_rotl(v56, int32(8))
						v67 = v58 - v63 ^ base.I32_rotl(v63, int32(16))
						v71 = v59 - v67 ^ base.I32_rotl(v67, int32(19))
						v72 = v63 + v59
						v73 = v67 + v72
						v74 = v71 + v73
						v78 = v72 - v71 ^ base.I32_rotl(v71, v50)
						v79 = int32(12)
						v80 = v37 + v79
						v82 = v38 - v79
						if base.Ui32(int32(11)) < base.Ui32(v82) {
							v37 = v80
							v38 = v82
							v39 = v73
							v40 = v74
							v41 = v78
							continue
						} else {
							break
						}
						break
					}
					v85 = v80
					v86 = v82
					v87 = v73
					v88 = v74
					v89 = v78
				}
				switch v86 - int32(1) {
				case 0:
					v136 = v87
					v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
					v257 = v136 + v137
					v258 = v88
					v259 = v89
				case 1:
					v131 = v87
					v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
					v136 = v132<<(uint(int32(8))%32) + v131
					v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
					v257 = v136 + v137
					v258 = v88
					v259 = v89
				case 2:
					v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+2)))
					v131 = v127<<(uint(int32(16))%32) + v87
					v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
					v136 = v132<<(uint(int32(8))%32) + v131
					v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
					v257 = v136 + v137
					v258 = v88
					v259 = v89
				case 3:
					v124 = v88
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
					v257 = v125 + v87
					v258 = v124
					v259 = v89
				case 4:
					v121 = v88
					v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
					v124 = v121 + v122
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
					v257 = v125 + v87
					v258 = v124
					v259 = v89
				case 5:
					v116 = v88
					v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+5)))
					v121 = v117<<(uint(int32(8))%32) + v116
					v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
					v124 = v121 + v122
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
					v257 = v125 + v87
					v258 = v124
					v259 = v89
				case 6:
					v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+6)))
					v116 = v112<<(uint(int32(16))%32) + v88
					v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+5)))
					v121 = v117<<(uint(int32(8))%32) + v116
					v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
					v124 = v121 + v122
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
					v257 = v125 + v87
					v258 = v124
					v259 = v89
				case 7:
					v107 = v89
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
					v257 = v108 + v87
					v258 = v110 + v88
					v259 = v107
				case 8:
					v102 = v89
					v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+8)))
					v107 = v103<<(uint(int32(8))%32) + v102
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
					v257 = v108 + v87
					v258 = v110 + v88
					v259 = v107
				case 9:
					v97 = v89
					v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+9)))
					v102 = v98<<(uint(int32(16))%32) + v97
					v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+8)))
					v107 = v103<<(uint(int32(8))%32) + v102
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
					v257 = v108 + v87
					v258 = v110 + v88
					v259 = v107
				case 10:
					v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+10)))
					v97 = v93<<(uint(int32(24))%32) + v89
					v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+9)))
					v102 = v98<<(uint(int32(16))%32) + v97
					v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+8)))
					v107 = v103<<(uint(int32(8))%32) + v102
					v108 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
					v257 = v108 + v87
					v258 = v110 + v88
					v259 = v107
				default:
					v257 = v87
					v258 = v88
					v259 = v89
				}
			}
			v262 = int32(14)
			v264 = v258 ^ v259 - base.I32_rotl(v258, v262)
			v268 = v264 ^ v257 - base.I32_rotl(v264, int32(11))
			v272 = v268 ^ v258 - base.I32_rotl(v268, int32(25))
			v276 = v272 ^ v264 - base.I32_rotl(v272, int32(16))
			v280 = v276 ^ v268 - base.I32_rotl(v276, int32(4))
			v284 = v280 ^ v272 - base.I32_rotl(v280, v262)
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v284 ^ v276 - base.I32_rotl(v284, int32(24))
			v294 = F_pg_snprintf(m, v11+int32(22), int32(10), int32(27372), v11)
			mBase = m.M
			v297 = m.ExcPending
			if v297 != 0 {
				return int32(0)
			} else {
				v302 = int32(8)
				v303 = v11 + int32(22)
				v304 = int32(17)
				v306 = v302 + int32(5)
				v307 = F_palloc(m, v306)
				mBase = m.M
				v308 = m.ExcPending
				if v308 != 0 {
					return int32(0)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v307)+4)) = uint8(v304)
					*(*int32)(unsafe.Add(mBase, uint32(v307))) = v306 << (uint(int32(2)) % 32)
					if v302 != 0 {
						v315 = F__emscripten_memcpy_bulkmem(m, v307+int32(5), v303, v302)
						mBase = m.M
					} else {
					}
					v318 = v307
					v323 = F_palloc(m, int32(12))
					mBase = m.M
					v324 = m.ExcPending
					if v324 != 0 {
						return int32(0)
					} else {
						v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						*(*int32)(unsafe.Add(mBase, uint32(v323)+4)) = v318
						*(*int32)(unsafe.Add(mBase, uint32(v323)+8)) = v325
						v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v323))) = v328
						v332 = v323
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v332
						v340 = int32(1)
						m.G0 = v11 + int32(32)
						return v340
					}
				}
			}
		} else {
			v302 = v22
			v303 = v20
			v304 = int32(1)
			v306 = v302 + int32(5)
			v307 = F_palloc(m, v306)
			mBase = m.M
			v308 = m.ExcPending
			if v308 != 0 {
				return int32(0)
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v307)+4)) = uint8(v304)
				*(*int32)(unsafe.Add(mBase, uint32(v307))) = v306 << (uint(int32(2)) % 32)
				if v302 != 0 {
					v315 = F__emscripten_memcpy_bulkmem(m, v307+int32(5), v303, v302)
					mBase = m.M
				} else {
				}
				v318 = v307
				v323 = F_palloc(m, int32(12))
				mBase = m.M
				v324 = m.ExcPending
				if v324 != 0 {
					return int32(0)
				} else {
					v325 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v323)+4)) = v318
					*(*int32)(unsafe.Add(mBase, uint32(v323)+8)) = v325
					v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v323))) = v328
					v332 = v323
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v332
					v340 = int32(1)
					m.G0 = v11 + int32(32)
					return v340
				}
			}
		}
	default:
		v340 = v3
		m.G0 = v11 + int32(32)
		return v340
	case 6:
		v332 = v3
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v332
		v340 = int32(1)
		m.G0 = v11 + int32(32)
		return v340
	}
}
func F_jsonb_ops__extract_nodes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v111 = l3
	goto L3
L3:
	;
	return v111
L4:
	;
	v12 = l3
	v13 = v8
	goto L7
L5:
	;
	v37 = l3
	goto L6
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v41 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v16 == int32(25) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v37 = v30
	goto L6
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v21 = F_palloc(m, int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v30 = v12
	goto L11
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v33 != 0 {
		v12 = v30
		v13 = v33
		goto L7
	} else {
		goto L15
	}
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(2)
	v28 = F_lappend(m, v12, v21)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v30 = v28
	goto L11
L15:
	;
	goto L8
L16:
	;
	v102 = F_lappend(m, v37, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L12
	} else {
		goto L34
	}
L17:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v44 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	goto L19
L19:
	;
	v91 = F_make_scalar_key(m, l2, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L12
	} else {
		goto L32
	}
L20:
	;
	v82 = F_make_scalar_key(m, l2, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L12
	} else {
		goto L30
	}
L21:
	;
	v81 = int32(0)
	goto L20
L22:
	;
	v56 = F_make_scalar_key(m, l2, int32(1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L25
	}
L23:
	;
	v45 = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v46 == v45 {
		v81 = v45
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	switch v50 - int32(21) {
	case 0, 2:
		v81 = int32(1)
		goto L20
	default:
		goto L21
	case 3:
		goto L22
	}
L25:
	;
	v59 = F_palloc(m, int32(8))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(2)
	v65 = F_make_scalar_key(m, l2, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v68 = F_palloc(m, int32(8))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(2)
	v74 = F_palloc(m, int32(16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v74))) = int64(8589934592)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v68
	v99 = v74
	goto L16
L30:
	;
	v85 = F_palloc(m, int32(8))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(2)
	v99 = v85
	goto L16
L32:
	;
	v94 = F_palloc(m, int32(8))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(2)
	v99 = v94
	goto L16
L34:
	;
	v111 = v102
	goto L3
}
func F_jsonb_path_ops__add_path_item(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v10 - int32(21) {
	case 0, 2:
		v30 = v9
		m.G0 = v7 + int32(32)
		return v30
	default:
		v30 = int32(0)
		m.G0 = v7 + int32(32)
		return v30
	case 4:
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(1)
		v18 = v7 + int32(16)
		if v18 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = v19
		} else {
		}
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v21
		F_JsonbHashScalarValue(m, v7+int32(12), l0)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v30 = v9
			m.G0 = v7 + int32(32)
			return v30
		}
	case 6:
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
		v30 = v9
		m.G0 = v7 + int32(32)
		return v30
	}
}
func F_jsonb_path_query_first_internal(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v21 = F_pg_detoast_datum(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v30 = F_executeJsonPath(m, v16, v21, int32(1409), int32(1410), v11, base.B2i32(v25 == int32(0)), v8+int32(8), l1)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					if v32 != 0 {
						v42 = v32
						v43 = F_JsonbValueToJsonb(m, v42)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v49 = v43
							m.G0 = v8 + int32(16)
							return v49
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						if v33 == int32(0) {
							v45 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
							v49 = int32(0)
							m.G0 = v8 + int32(16)
							return v49
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
							if v36 <= int32(0) {
								v45 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
								v49 = int32(0)
								m.G0 = v8 + int32(16)
								return v49
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
								v42 = v40
								v43 = F_JsonbValueToJsonb(m, v42)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v49 = v43
									m.G0 = v8 + int32(16)
									return v49
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_jsonb_path_query_tz(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_query_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_populate_record_valid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[1091]))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v8
	v11 = *(*int64)(unsafe.Add(mBase, _consts[1092]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v11
	v16 = F_populate_record_worker(m, l0, int32(400237), int32(0), int32(1), v5)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)))
		m.G0 = v5 + int32(16)
		return v20 ^ int32(1)
	}
}
func F_jsonb_put_escaped_value(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v3 {
	case 0:
		F_appendBinaryStringInfo(m, l0, int32(288456), int32(4))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			return
		}
	case 1:
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		F_escape_json_with_len(m, l0, v4, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	case 2:
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v11 = F_DirectFunctionCall1Coll(m, int32(617), int32(0), v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_appendStringInfoString(m, l0, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		}
	case 3:
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
		if v15 == int32(1) {
			F_appendBinaryStringInfo(m, l0, int32(327357), int32(4))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				return
			}
		} else {
			F_appendBinaryStringInfo(m, l0, int32(343932), int32(5))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				return
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(350633), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				F_errfinish(m, int32(476670), int32(371), int32(328307))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
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
func F_jsonb_string_to_tsvector_byid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = l0 + int32(28)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v9
		v18 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v7 + int32(8)
		F_iterate_jsonb_values(m, v13, int32(2), v7+int32(24))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v32 = F_make_tsvector(m, v7+int32(8))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				if v34 != v13 {
					F_pfree(m, v13)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(32)
						return v32
					}
				} else {
					m.G0 = v7 + int32(32)
					return v32
				}
			}
		}
	}
}
func F_jsonb_strip_nulls(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(0)
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v17 == int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = base.B2i32(v20 != int32(0))
	goto L5
L4:
	;
	v23 = v2
	goto L5
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
	if v24&int32(16) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = F_JsonbIteratorInit(m, v11+int32(4))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v99 = v11
	goto L8
L8:
	;
	m.G0 = v8 + int32(48)
	return v99
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v31
	v38 = v2
	goto L10
L10:
	;
	v40 = int32(0)
	goto L13
L11:
	;
	v95 = F_JsonbValueToJsonb(m, v38)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L34
	}
L12:
	;
	goto L11
L13:
	;
	v50 = F_JsonbIteratorNext(m, v8+int32(44), v8+int32(20), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L14:
	;
	if v50&int32(-2) == int32(2) {
		goto L30
	} else {
		goto L31
	}
L15:
	;
	if v40&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v8)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v8)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v56
	v40 = int32(1)
	goto L13
L17:
	;
	switch v50 {
	case 0:
		goto L12
	case 1:
		goto L16
	default:
		goto L15
	}
L18:
	;
	if v50 == int32(2) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if v23 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v63 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v64 == v63 {
		v40 = v63
		goto L13
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v71 = F_pushJsonbValue(m, v8+int32(40), int32(1), v8)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	goto L20
L26:
	;
	goto L14
L27:
	;
	if v50 != int32(3) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v78 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v79 == v78 {
		v40 = v78
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v92 = v8 + int32(20)
	goto L32
L31:
	;
	v92 = int32(0)
	goto L32
L32:
	;
	v93 = F_pushJsonbValue(m, v8+int32(40), v50, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v38 = v93
	goto L10
L34:
	;
	v99 = v95
	goto L8
}
func F_jsonb_typeof(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v9 + int32(4)
		v17 = F_JsonbExtractScalar(m, v14, v6+int32(12))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 != 0 {
				v21 = F_JsonbTypeName(m, v6+int32(12))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v33 = v21
					v34 = F_cstring_to_text(m, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(32)
						return v34
					}
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				if v24&int32(1073741824) != 0 {
					v33 = int32(24367)
					v34 = F_cstring_to_text(m, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(32)
						return v34
					}
				} else {
					if v24&int32(536870912) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v44
							F_errmsg_internal(m, int32(27340), v6)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(476670), int32(171), int32(363230))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v33 = int32(103833)
						v34 = F_cstring_to_text(m, v33)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v6 + int32(32)
							return v34
						}
					}
				}
			}
		}
	}
}
