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
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	if int32(0) < l1 {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
		if v6&int32(5) == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v11)
			return int32(0)
		} else {
			F_deconstruct_expanded_record(m, l0)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				if v19 < l1 {
					v21 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v21)
					return int32(0)
				} else {
					v26 = l1 - int32(1)
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v27))))
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v29)
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v26<<(uint(int32(2))%32))))
					return v35
				}
			}
		}
	} else {
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		if v37 == int32(0) {
			v40 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v40)
			return int32(0)
		} else {
			v44 = F_heap_getsysattr(m, v37, l1, l2)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				return v44
			}
		}
	}
}
func F_expanded_record_get_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v4&int32(1) != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
		return v7
	} else {
		if v4&int32(4) != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			v14 = F_heap_form_tuple(m, v11, v12, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = v14
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
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
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	v4 = l3
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if l5 == int32(0) {
		v134 = v14
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v134&int32(4) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L2:
	;
	if v14&int32(64) == int32(0) {
		v134 = v14
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	F_build_dummy_expanded_header(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v28&int32(5) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v63 | int32(4)
	if l1 <= int32(0) {
		goto L22
	} else {
		goto L23
	}
L7:
	;
	F_deconstruct_expanded_record(m, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	v56 = F__emscripten_memset_bulkmem(m, v50, base.I32_extend8_s(int32(0)), v52<<(uint(int32(2))%32))
	mBase = m.M
	goto L19
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	v37 = v35 << (uint(int32(2)) % 32)
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	if v42 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v38 = F__emscripten_memcpy_bulkmem(m, v33, v34, v37)
	mBase = m.M
	goto L14
L13:
	;
	goto L14
L14:
	;
	goto L11
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v63 = v45 | v46&int32(16)
	goto L6
L16:
	;
	v43 = F__emscripten_memcpy_bulkmem(m, v40, v41, v42)
	mBase = m.M
	goto L18
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	v61 = F__emscripten_memset_bulkmem(m, v57, base.I32_extend8_s(int32(1)), v59)
	mBase = m.M
	goto L20
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v63 = v62
	goto L6
L21:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v134 = v132
	goto L1
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L32
	}
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	if v69 < l1 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	v73 = l1 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71+v73<<(uint(int32(2))%32)))) = l2
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	*(*uint8)(unsafe.Add(mBase, uint32(v78+v73))) = uint8(v4)
	if v4 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v97 = int32(4489152)
	v98 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v100
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_domain_check(m, v27+int32(18), int32(0), v105, l0+int32(104), v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L30
	}
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v84 = v81 + v73<<(uint(int32(4))%32)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+26)))
	if v85 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+24)))
	if v86 != int32(65535) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v89 != int32(1) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v92 | int32(16)
	goto L25
L30:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v98
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_MemoryContextReset(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	m.G0 = v23 + int32(16)
	goto L21
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = l1
	F_errmsg_internal(m, int32(420098), v23)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(497773), int32(1530), int32(430231))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_deconstruct_expanded_record(m, l0)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if l1 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L76
	}
L40:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v144 < l1 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v148 = l1 - int32(1)
	v153 = v146 + v148<<(uint(int32(4))%32) + int32(20)
	if v4 != 0 {
		v221 = l2
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v224 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v226 & int32(-2)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+6)))
	if v232 == v224 {
		goto L65
	} else {
		goto L66
	}
L43:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+6)))
	if v154 != 0 {
		v221 = l2
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v155 = int32(0)
	if l4 == v155 {
		v190 = l2
		v193 = v155
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v194 = int32(4489152)
	v195 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v197
	v200 = int32(*(*int16)(unsafe.Add(mBase, uint32(v153)+4)))
	v201 = F_datumCopy(m, v190, int32(0), v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L56
	}
L46:
	;
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+4)))
	if v159 != int32(65535) {
		v190 = l2
		v193 = int32(0)
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v163 != int32(1) {
		v190 = l2
		v193 = int32(0)
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v166 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v181 = int32(4489152)
	v182 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v180
	v185 = F_detoast_external_attr(m, l2)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L55
	}
L50:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v174 = F_AllocSetContextCreateInternal(m, v169, int32(60686), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_MemoryContextReset(m, v166)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v174
	v180 = v174
	goto L49
L54:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v180 = v179
	goto L49
L55:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v182
	v190 = v185
	v193 = int32(1)
	goto L45
L56:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v195
	if v193 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	F_MemoryContextReset(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v208 | int32(8)
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+4)))
	if v212 != int32(65535) {
		v221 = v201
		goto L42
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	if v215 != int32(1) {
		v221 = v201
		goto L42
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v208 | int32(24)
	v221 = v201
	goto L42
L63:
	;
	m.G0 = v12 + int32(16)
	return
L64:
	;
	v248 = v231 + v148<<(uint(int32(2))%32)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v221
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v4)
	if v249 == int32(0) {
		goto L63
	} else {
		goto L69
	}
L65:
	;
	v235 = v148 + v230
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	if v236 != int32(1) {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231+v148<<(uint(int32(2))%32)))) = v221
	*(*uint8)(unsafe.Add(mBase, uint32(v148+v230))) = uint8(v4)
	goto L63
L68:
	;
	goto L67
L69:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v254&int32(128) != 0 {
		goto L63
	} else {
		goto L70
	}
L70:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if base.Ui32(v257) <= base.Ui32(v249) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if base.Ui32(v249) < base.Ui32(v259) {
		goto L63
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_pfree(m, v249)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	goto L63
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	F_errmsg_internal(m, int32(420098), v12)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(497773), int32(1143), int32(311405))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
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
