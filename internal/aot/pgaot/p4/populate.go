package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_populate_array_dim_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_check_stack_depth(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v19 == int32(18) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v12 + int32(48)
	return v272
L4:
	;
	v32 = F_JsonbIteratorInit(m, v14)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v22&int32(1342177280) == int32(1073741824) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_populate_array_report_expected_array(m, l0, l2-int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	v272 = int32(0)
	goto L3
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v32
	v40 = F_JsonbIteratorNext(m, v12+int32(40), v12+int32(20), int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v47 = F_JsonbIteratorNext(m, v12+int32(40), v12+int32(20), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(0) < v49 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v263 = int32(1)
	v269 = F_JsonbIteratorNext(m, v12+int32(40), v12+int32(20), v263)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L57
	}
L14:
	;
	goto L38
L15:
	;
	v161 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v161)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v12 + int32(20)
	if v47 != int32(3) {
		goto L13
	} else {
		goto L37
	}
L16:
	;
	switch v47 - int32(3) {
	case 0:
		goto L18
	default:
		goto L13
	case 2:
		goto L17
	}
L17:
	;
	v68 = int32(0)
	if l2 <= v68 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v54 != int32(18) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+3)))
	if v58&int32(64) == int32(0) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v63 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v63)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v12 + int32(20)
	goto L14
L21:
	;
	F_populate_array_report_expected_array(m, l0, l2)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l2
	v75 = l2 << (uint(int32(2)) % 32)
	v76 = F_palloc(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v272 = v68
	goto L3
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v76
	v79 = F_palloc0(m, v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v79
	v82 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v88 = v82
	v93 = v4
	goto L30
L28:
	;
	v120 = v82
	goto L29
L29:
	;
	v129 = l2 & int32(3)
	if v129 == int32(0) {
		goto L15
	} else {
		goto L33
	}
L30:
	;
	v97 = v88 << (uint(int32(2)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v100 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v97+v98))) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v102+v97)+4)) = v100
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v106+v97)+8)) = v100
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v110+v97)+12)) = v100
	v114 = int32(4)
	v115 = v88 + v114
	v117 = v93 + v114
	if v117 != l2&int32(2147483644) {
		v88 = v115
		v93 = v117
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v120 = v115
	goto L29
L32:
	;
	goto L31
L33:
	;
	v133 = v120
	v137 = v4
	goto L34
L34:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v141+v133<<(uint(int32(2))%32)))) = int32(-1)
	v147 = int32(1)
	v150 = v137 + v147
	if v150 != v129 {
		v133 = v133 + v147
		v137 = v150
		goto L34
	} else {
		goto L36
	}
L35:
	;
	goto L15
L36:
	;
	goto L35
L37:
	;
	goto L14
L38:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v190 <= int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L13
L40:
	;
	v250 = F_JsonbIteratorNext(m, v12+int32(40), v12+int32(20), int32(1))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L55
	}
L41:
	;
	v232 = int32(0)
	v235 = F_populate_array_dim_jsonb(m, l0, v12+int32(20), l2+int32(1))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L51
	}
L42:
	;
	if l2 < v190 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	v198 = int32(0)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v207 = F_populate_record_field(m, v195, v196, v197, v198, v199, v198, v12+int32(4), v12+int32(47), v205, v198)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v209 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v272 = int32(0)
	goto L3
L46:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+47)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v221 = F_accumArrayResult(m, v216, v207, v217, v219, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L50
	}
L47:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v212 != int32(447) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
	if v215 != 0 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v226 = v223 + l2<<(uint(int32(2))%32) - int32(4)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	*(*int32)(unsafe.Add(mBase, uint32(v226))) = v227 + int32(1)
	goto L40
L51:
	;
	if v235 == int32(0) {
		v272 = v232
		goto L3
	} else {
		goto L52
	}
L52:
	;
	v239 = F_populate_array_check_dimension(m, l0, l2)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v239 == int32(0) {
		v272 = v232
		goto L3
	} else {
		goto L54
	}
L54:
	;
	goto L40
L55:
	;
	if v250 == int32(3) {
		goto L38
	} else {
		goto L56
	}
L56:
	;
	goto L39
L57:
	;
	v272 = v263
	goto L3
}
func F_populate_array_element_start(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
	if int32(0) < v6 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+32))
		if v9 != v6 {
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v13
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v13
	}
	return int32(0)
}
func F_populate_recordset_array_element_start(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v9 != int32(1) {
		m.G0 = v6 + int32(16)
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
		if v12 == int32(3) {
			m.G0 = v6 + int32(16)
			return int32(0)
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v24
					F_errmsg(m, int32(122378), v6)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(483564), int32(4277), int32(80611))
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
		}
	}
}
func F_populate_recordset_object_end(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	if v8 <= int32(1) {
		v11 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+8)) = uint8(v11)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v13
		F_populate_recordset_record(m, l0, v5+int32(8))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_hash_destroy(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
				m.G0 = v5 + int32(16)
				return int32(0)
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return int32(0)
	}
}
func F_populate_recordset_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+32))
	if v6 <= int32(2) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
		switch v9 - int32(3) {
		case 0, 2:
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			v15 = v14
		default:
			v15 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v15
	} else {
	}
	return int32(0)
}
