package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DeleteExpandedObject(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))+2))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	F_MemoryContextDelete(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_expanded_record_fetch_field(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v32 int64
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
			return int64(0)
		} else {
			F_deconstruct_expanded_record(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				if v15 < l1 {
					v36 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v36)
					return int64(0)
				} else {
					v18 = l1 - int32(1)
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v19))))
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v21)
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v27 = *(*int64)(unsafe.Add(mBase, uint32(v23+v18<<(uint(int32(3))%32))))
					return v27
				}
			}
		}
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		if v29 == int32(0) {
			v36 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v36)
			return int64(0)
		} else {
			v32 = F_heap_getsysattr(m, v29, l1, l2)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int64(0)
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
func F_expanded_record_set_field_internal(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int64
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v226 int64
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	v4 = l3
	v7 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.B2i32(l5 == v7)|base.B2i32(v17&int32(64) == v7) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v136 = v17
	goto L3
L2:
	;
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	F_build_dummy_expanded_header(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v136&int32(4) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L4:
	;
	return
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v30&int32(5) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v62 | int32(4)
	if l1 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	F_deconstruct_expanded_record(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v52 = v50 << (uint(int32(3)) % 32)
	if v52 != 0 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v37 = v35 << (uint(int32(3)) % 32)
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	base.MemoryCopy(m, v38, v39, v37)
	goto L13
L12:
	;
	goto L13
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	if v41 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	base.MemoryCopy(m, v42, v43, v41)
	goto L16
L15:
	;
	goto L16
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v62 = v45 | v46&int32(16)
	goto L6
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	base.MemoryFill(m, v53, int32(0), v52)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	if v56 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	base.MemoryFill(m, v57, int32(1), v56)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	v62 = v60
	goto L6
L23:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v136 = v131
	goto L3
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L34
	}
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	if v68 < l1 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v71 = l1 - int32(1)
	v73 = v71 << (uint(int32(3)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v73+v74))) = l2
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v77+v71))) = uint8(v4)
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
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_domain_check(m, base.I64_extend_i32_u(v29+int32(18)), int32(0), v104, l0+int32(104), v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L32
	}
L28:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v81 = v80 + v73
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+32)))
	if v82 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+30)))
	if v83 != int32(_a_F_expanded_record_set_field_internal_4) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l2)))))
	if v87 != int32(1) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v90 | int32(16)
	goto L27
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0])) = v96
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_MemoryContextReset(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	m.G0 = v25 + int32(16)
	goto L23
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = l1
	F_errmsg_internal(m, int32(_a_F_expanded_record_set_field_internal_0), v25)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_expanded_record_set_field_internal_1), int32(1530), int32(_a_F_expanded_record_set_field_internal_7))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
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
	v142 = m.ExcPending
	if v142 != 0 {
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
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L78
	}
L42:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v145 < l1 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v149 = l1 - int32(1)
	v154 = v147 + v149<<(uint(int32(3))%32) + int32(28)
	if v4 != 0 {
		v226 = l2
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v230 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v232 & int32(-2)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)))
	if v238 == v230 {
		goto L67
	} else {
		goto L68
	}
L45:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)))
	if v155 != 0 {
		v226 = l2
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v156 = int32(0)
	if l4 == v156 {
		v193 = l2
		v197 = v156
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v198 = int32(_a_F_expanded_record_set_field_internal_3)
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0])) = v201
	v204 = int32(*(*int16)(unsafe.Add(mBase, uint32(v154)+2)))
	v205 = F_datumCopy(m, v193, int32(0), v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L58
	}
L48:
	;
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+2)))
	if v160 != int32(_a_F_expanded_record_set_field_internal_4) {
		v193 = l2
		v197 = int32(0)
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v164 = base.I32_wrap_i64(l2)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v165 != int32(1) {
		v193 = l2
		v197 = int32(0)
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v168 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v183 = int32(_a_F_expanded_record_set_field_internal_3)
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0])) = v182
	v187 = F_detoast_external_attr(m, v164)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L57
	}
L52:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v176 = F_AllocSetContextCreateInternal(m, v171, int32(_a_F_expanded_record_set_field_internal_5), int32(0), int32(1024), int32(_a_F_expanded_record_set_field_internal_6))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	F_MemoryContextReset(m, v168)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v176
	v182 = v176
	goto L51
L56:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v182 = v181
	goto L51
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0])) = v184
	v193 = base.I64_extend_i32_u(v187)
	v197 = int32(1)
	goto L47
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_expanded_record_set_field_internal[0])) = v199
	if v197 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_MemoryContextReset(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v212 | int32(8)
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+2)))
	if v216 != int32(_a_F_expanded_record_set_field_internal_4) {
		v226 = v205
		goto L44
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v205)))))
	if v220 != int32(1) {
		v226 = v205
		goto L44
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v212 | int32(24)
	v226 = v205
	goto L44
L65:
	;
	m.G0 = v13 + int32(16)
	return
L66:
	;
	v254 = v237 + v149<<(uint(int32(3))%32)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	*(*int64)(unsafe.Add(mBase, uint32(v254))) = v226
	*(*uint8)(unsafe.Add(mBase, uint32(v241))) = uint8(v4)
	if v255 == int32(0) {
		goto L65
	} else {
		goto L71
	}
L67:
	;
	v241 = v236 + v149
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	if v242 != int32(1) {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v237+v149<<(uint(int32(3))%32)))) = v226
	*(*uint8)(unsafe.Add(mBase, uint32(v236+v149))) = uint8(v4)
	goto L65
L70:
	;
	goto L69
L71:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v260&int32(128) != 0 {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if base.Ui32(v263) <= base.Ui32(v255) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if base.Ui32(v255) < base.Ui32(v265) {
		goto L65
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_pfree(m, v255)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	F_errmsg_internal(m, int32(_a_F_expanded_record_set_field_internal_0), v13)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_expanded_record_set_field_internal_1), int32(1143), int32(_a_F_expanded_record_set_field_internal_2))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
