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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v217 int64
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_check_stack_depth(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v21 == int32(18) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v14 - int32(-64)
	return v287
L4:
	;
	v34 = F_JsonbIteratorInit(m, v16)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v24&int32(1342177280) == int32(1073741824) {
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
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	v287 = int32(0)
	goto L3
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v34
	v38 = v12 + int32(-8)
	v40 = v12 + int32(-40)
	v42 = F_JsonbIteratorNext(m, v38, v40, int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v45 = F_JsonbIteratorNext(m, v38, v40, int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(0) < v47 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	F_populate_array_report_expected_array(m, l0, l2)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L56
	}
L14:
	;
	v275 = int32(1)
	v281 = F_JsonbIteratorNext(m, v12+int32(-8), v12+int32(-40), v275)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L55
	}
L15:
	;
	goto L36
L16:
	;
	v164 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v164)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v12 + int32(-40)
	if v45 != int32(3) {
		goto L14
	} else {
		goto L35
	}
L17:
	;
	switch v45 - int32(3) {
	case 0:
		goto L19
	default:
		goto L14
	case 2:
		goto L18
	}
L18:
	;
	if l2 <= int32(0) {
		goto L13
	} else {
		goto L22
	}
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	if v52 != int32(18) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+3)))
	if v56&int32(64) == int32(0) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v61 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v61)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v12 + int32(-40)
	goto L15
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l2
	v70 = F_palloc_mul(m, int32(4), l2)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v70
	v74 = F_palloc0_mul(m, int32(4), l2)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v74
	v78 = l2 & int32(3)
	v79 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v85 = v79
	v92 = v4
	goto L28
L26:
	;
	v121 = v79
	goto L27
L27:
	;
	v132 = v121
	v140 = v4
	goto L32
L28:
	;
	v96 = v85 << (uint(int32(2)) % 32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v99 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v96+v97))) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v101+v96)+4)) = v99
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v105+v96)+8)) = v99
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v109+v96)+12)) = v99
	v113 = int32(4)
	v114 = v85 + v113
	v116 = v92 + v113
	if v116 != l2&int32(2147483644) {
		v85 = v114
		v92 = v116
		goto L28
	} else {
		goto L30
	}
L29:
	;
	if v78 == int32(0) {
		goto L16
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v121 = v114
	goto L27
L32:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v142+v132<<(uint(int32(2))%32)))) = int32(-1)
	v148 = int32(1)
	v151 = v140 + v148
	if v151 != v78 {
		v132 = v132 + v148
		v140 = v151
		goto L32
	} else {
		goto L34
	}
L33:
	;
	goto L16
L34:
	;
	goto L33
L35:
	;
	goto L15
L36:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v198 = int32(0)
	if base.B2i32(v197 <= v198)|base.B2i32(l2 < v197) == v198 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L14
L38:
	;
	v260 = F_JsonbIteratorNext(m, v12+int32(-8), v12+int32(-40), int32(1))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L53
	}
L39:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v204)+8))
	v208 = int32(0)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v217 = F_populate_record_field(m, v205, v206, v207, v208, v209, int64(0), v12+int32(-56), v12+int32(-1), v215, v208)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v242 = int32(0)
	v245 = F_populate_array_dim_jsonb(m, l0, v12+int32(-40), l2+int32(1))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L49
	}
L42:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v219 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v287 = int32(0)
	goto L3
L44:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+63)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v231 = F_accumArrayResult(m, v226, v217, v227, v229, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	if v222 != int32(453) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+4)))
	if v225 != 0 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v236 = v233 + l2<<(uint(int32(2))%32) - int32(4)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v237 + int32(1)
	goto L38
L49:
	;
	if v245 == int32(0) {
		v287 = v242
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v249 = F_populate_array_check_dimension(m, l0, l2)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v249 == int32(0) {
		v287 = v242
		goto L3
	} else {
		goto L52
	}
L52:
	;
	goto L38
L53:
	;
	if v260 == int32(3) {
		goto L36
	} else {
		goto L54
	}
L54:
	;
	goto L37
L55:
	;
	v287 = v275
	goto L3
L56:
	;
	v287 = int32(0)
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
					F_errmsg(m, int32(_a_F_populate_recordset_array_element_start_0), v6)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_populate_recordset_array_element_start_1), int32(_a_F_populate_recordset_array_element_start_2), int32(_a_F_populate_recordset_array_element_start_3))
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
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
	if v7 <= int32(2) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
		switch v9 - int32(3) {
		case 0, 2:
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v15 = v14
		default:
			v15 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v15
	} else {
	}
	return int32(0)
}
