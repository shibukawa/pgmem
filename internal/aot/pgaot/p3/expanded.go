package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DeleteExpandedObject(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
	F_MemoryContextDelete(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_expanded_record_fetch_field(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	if int32(0) < l1 {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
		if v6&int32(5) == int32(0) {
			v36 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v36)
			return int32(0)
		} else {
			F_deconstruct_expanded_record(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				if v15 < l1 {
					v36 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v36)
					return int32(0)
				} else {
					v18 = l1 - int32(1)
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v19))))
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v21)
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v18<<(uint(int32(2))%32))))
					return v27
				}
			}
		}
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		if v29 == int32(0) {
			v36 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v36)
			return int32(0)
		} else {
			v32 = F_heap_getsysattr(m, v29, l1, l2)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				return v32
			}
		}
	}
}
func F_expanded_record_get_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3&int32(1) != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		return v6
	} else {
		if v3&int32(4) != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			v13 = F_heap_form_tuple(m, v10, v11, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v18 = v13
				return v18
			}
		} else {
			v18 = int32(0)
			return v18
		}
	}
}
func F_expanded_record_set_field_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
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
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	v4 = l3
	v7 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.B2i32(l5 == v7)|base.B2i32(v16&int32(64) == v7) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v134 = v16
	goto L3
L2:
	;
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	F_build_dummy_expanded_header(m, l0)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v134&int32(4) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L4:
	;
	return
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v29&int32(5) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v61 | int32(4)
	if l1 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	F_deconstruct_expanded_record(m, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	v51 = v49 << (uint(int32(2)) % 32)
	if v51 != 0 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	v36 = v34 << (uint(int32(2)) % 32)
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+56))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	base.MemoryCopy(m, v37, v38, v36)
	goto L13
L12:
	;
	goto L13
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	base.MemoryCopy(m, v41, v42, v40)
	goto L16
L15:
	;
	goto L16
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v61 = v44 | v45&int32(16)
	goto L6
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v28)+56))
	base.MemoryFill(m, v52, int32(0), v51)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	if v55 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	base.MemoryFill(m, v56, int32(1), v55)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v61 = v59
	goto L6
L23:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v134 = v130
	goto L3
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L34
	}
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	if v67 < l1 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v28)+56))
	v71 = l1 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v69+v71<<(uint(int32(2))%32)))) = l2
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v76+v71))) = uint8(v4)
	if v4 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v95 = int32(_a_F_expanded_record_set_field_internal_3)
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0])) = v98
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_domain_check(m, v28+int32(18), int32(0), v103, l0+int32(104), v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L32
	}
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v82 = v79 + v71<<(uint(int32(4))%32)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+26)))
	if v83 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+24)))
	if v84 != int32(_a_F_expanded_record_set_field_internal_4) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v87 != int32(1) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v90 | int32(16)
	goto L27
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0])) = v96
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_MemoryContextReset(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	m.G0 = v24 + int32(16)
	goto L23
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = l1
	F_errmsg_internal(m, int32(_a_F_expanded_record_set_field_internal_0), v24)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_expanded_record_set_field_internal_1), int32(1530), int32(_a_F_expanded_record_set_field_internal_7))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
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
	F_deconstruct_expanded_record(m, l0)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if l1 <= int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L78
	}
L42:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v143 < l1 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v147 = l1 - int32(1)
	v152 = v145 + v147<<(uint(int32(4))%32) + int32(20)
	if v4 != 0 {
		v220 = l2
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v223 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v223
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v225 & int32(-2)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+6)))
	if v231 == v223 {
		goto L67
	} else {
		goto L68
	}
L45:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+6)))
	if v153 != 0 {
		v220 = l2
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v154 = int32(0)
	if l4 == v154 {
		v189 = l2
		v192 = v154
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v193 = int32(_a_F_expanded_record_set_field_internal_3)
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0]))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0])) = v196
	v199 = int32(*(*int16)(unsafe.Add(mBase, uint32(v152)+4)))
	v200 = F_datumCopy(m, v189, int32(0), v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L58
	}
L48:
	;
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+4)))
	if v158 != int32(_a_F_expanded_record_set_field_internal_4) {
		v189 = l2
		v192 = int32(0)
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v162 != int32(1) {
		v189 = l2
		v192 = int32(0)
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v165 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v180 = int32(_a_F_expanded_record_set_field_internal_3)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0])) = v179
	v184 = F_detoast_external_attr(m, l2)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L57
	}
L52:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v173 = F_AllocSetContextCreateInternal(m, v168, int32(_a_F_expanded_record_set_field_internal_5), int32(0), int32(1024), int32(_a_F_expanded_record_set_field_internal_6))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	F_MemoryContextReset(m, v165)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v173
	v179 = v173
	goto L51
L56:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v179 = v178
	goto L51
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0])) = v181
	v189 = v184
	v192 = int32(1)
	goto L47
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0])) = v194
	if v192 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_MemoryContextReset(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v207 | int32(8)
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+4)))
	if v211 != int32(_a_F_expanded_record_set_field_internal_4) {
		v220 = v200
		goto L44
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if v214 != int32(1) {
		v220 = v200
		goto L44
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v207 | int32(24)
	v220 = v200
	goto L44
L65:
	;
	m.G0 = v12 + int32(16)
	return
L66:
	;
	v247 = v230 + v147<<(uint(int32(2))%32)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v220
	*(*uint8)(unsafe.Add(mBase, uint32(v234))) = uint8(v4)
	if v248 == int32(0) {
		goto L65
	} else {
		goto L71
	}
L67:
	;
	v234 = v229 + v147
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	if v235 != int32(1) {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230+v147<<(uint(int32(2))%32)))) = v220
	*(*uint8)(unsafe.Add(mBase, uint32(v229+v147))) = uint8(v4)
	goto L65
L70:
	;
	goto L69
L71:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v253&int32(128) != 0 {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if base.Ui32(v256) <= base.Ui32(v248) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if base.Ui32(v248) < base.Ui32(v258) {
		goto L65
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_pfree(m, v248)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	goto L65
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	F_errmsg_internal(m, int32(_a_F_expanded_record_set_field_internal_0), v12)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_expanded_record_set_field_internal_1), int32(1143), int32(_a_F_expanded_record_set_field_internal_2))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_make_expanded_record_for_rec(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v6 != int32(2249) {
		F_revalidate_rectypeid(m, l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			if l2 == int32(0) {
				v22 = F_make_expanded_record_from_typeid(m, v13, int32(-1), v5)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return v22
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
				if v13 != v16 {
					v22 = F_make_expanded_record_from_typeid(m, v13, int32(-1), v5)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						return v22
					}
				} else {
					v18 = F_make_expanded_record_from_exprecord(m, l2, v5)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						return v18
					}
				}
			}
		}
	} else {
		if l2 == int32(0) {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
			if v33 != 0 {
				v36 = v33
				v37 = F_make_expanded_record_from_tupdesc(m, v36, v5)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					return v37
				}
			} else {
				v34 = F_expanded_record_fetch_tupdesc(m, l2)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v36 = v34
					v37 = F_make_expanded_record_from_tupdesc(m, v36, v5)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						return v37
					}
				}
			}
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)))
			if v27&int32(64) != 0 {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
				if v33 != 0 {
					v36 = v33
					v37 = F_make_expanded_record_from_tupdesc(m, v36, v5)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						return v37
					}
				} else {
					v34 = F_expanded_record_fetch_tupdesc(m, l2)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = v34
						v37 = F_make_expanded_record_from_tupdesc(m, v36, v5)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							return v37
						}
					}
				}
			} else {
				v30 = F_make_expanded_record_from_exprecord(m, l2, v5)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					return v30
				}
			}
		}
	}
}
