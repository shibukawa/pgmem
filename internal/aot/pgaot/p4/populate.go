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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
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
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_check_stack_depth(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v20 == int32(18) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v13 + int32(48)
	return v279
L4:
	;
	v33 = F_JsonbIteratorInit(m, v15)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v23&int32(1342177280) == int32(1073741824) {
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
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	v279 = int32(0)
	goto L3
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v33
	v37 = v13 + int32(40)
	v39 = v13 + int32(20)
	v41 = F_JsonbIteratorNext(m, v37, v39, int32(1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v44 = F_JsonbIteratorNext(m, v37, v39, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(0) < v46 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	F_populate_array_report_expected_array(m, l0, l2)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L56
	}
L14:
	;
	v267 = int32(1)
	v273 = F_JsonbIteratorNext(m, v13+int32(40), v13+int32(20), v267)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L55
	}
L15:
	;
	goto L36
L16:
	;
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v159)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v13 + int32(20)
	if v44 != int32(3) {
		goto L14
	} else {
		goto L35
	}
L17:
	;
	switch v44 - int32(3) {
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
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v51 != int32(18) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+3)))
	if v55&int32(64) == int32(0) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v60 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v60)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v13 + int32(20)
	goto L15
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l2
	v69 = l2 << (uint(int32(2)) % 32)
	v70 = F_palloc(m, v69)
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
	v73 = F_palloc0(m, v69)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v73
	v77 = l2 & int32(3)
	v78 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v84 = v78
	v91 = v4
	goto L28
L26:
	;
	v119 = v78
	goto L27
L27:
	;
	v129 = v119
	v137 = v4
	goto L32
L28:
	;
	v94 = v84 << (uint(int32(2)) % 32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v97 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v94+v95))) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v99+v94)+4)) = v97
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v103+v94)+8)) = v97
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v107+v94)+12)) = v97
	v111 = int32(4)
	v112 = v84 + v111
	v114 = v91 + v111
	if v114 != l2&int32(2147483644) {
		v84 = v112
		v91 = v114
		goto L28
	} else {
		goto L30
	}
L29:
	;
	if v77 == int32(0) {
		goto L16
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v119 = v112
	goto L27
L32:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v138+v129<<(uint(int32(2))%32)))) = int32(-1)
	v144 = int32(1)
	v147 = v137 + v144
	if v147 != v77 {
		v129 = v129 + v144
		v137 = v147
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
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v191 = int32(0)
	if base.B2i32(v190 <= v191)|base.B2i32(l2 < v190) == v191 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L14
L38:
	;
	v253 = F_JsonbIteratorNext(m, v13+int32(40), v13+int32(20), int32(1))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L53
	}
L39:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	v201 = int32(0)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v210 = F_populate_record_field(m, v198, v199, v200, v201, v202, v201, v13+int32(4), v13+int32(47), v208, v201)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v235 = int32(0)
	v238 = F_populate_array_dim_jsonb(m, l0, v13+int32(20), l2+int32(1))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L49
	}
L42:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v212 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v279 = int32(0)
	goto L3
L44:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v224 = F_accumArrayResult(m, v219, v210, v220, v222, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v215 != int32(447) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+4)))
	if v218 != 0 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v229 = v226 + l2<<(uint(int32(2))%32) - int32(4)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = v230 + int32(1)
	goto L38
L49:
	;
	if v238 == int32(0) {
		v279 = v235
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v242 = F_populate_array_check_dimension(m, l0, l2)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v242 == int32(0) {
		v279 = v235
		goto L3
	} else {
		goto L52
	}
L52:
	;
	goto L38
L53:
	;
	if v253 == int32(3) {
		goto L36
	} else {
		goto L54
	}
L54:
	;
	goto L37
L55:
	;
	v279 = v267
	goto L3
L56:
	;
	v279 = int32(0)
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
