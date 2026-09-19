package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetJsonBehaviorValueString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(2))%32))+uint32(_c_F_GetJsonBehaviorValueString[0])))
	v6 = F_pstrdup(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_JsonTablePlanNextRow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
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
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	switch v13 - int32(50) {
	case 0:
		goto L5
	case 1:
		v180 = l0
		v182 = int32(0)
		goto L2
	default:
		v141 = l0
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v200
L2:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v180)+52))
	v188 = F_JsonTablePlanNextRow(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L9
	} else {
		goto L51
	}
L3:
	;
	v174 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v163+int32(36)))) = v174
	v178 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v163+int32(40)))) = uint8(v178)
	v200 = base.B2i32(v28 != v174)
	goto L1
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L48
	}
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v29 != 0 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v17 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v20 = F_JsonTablePlanNextRow(m, v17)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if v20 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v200 = int32(1)
	goto L1
L12:
	;
	if v28 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v30
	v33 = v29 + int32(4)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if base.Ui32(v33) < base.Ui32(v36+v37<<(uint(int32(2))%32)) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	goto L12
L16:
	;
	v42 = v33
	goto L18
L17:
	;
	v42 = int32(0)
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v42
	goto L12
L19:
	;
	v163 = l0
	goto L3
L20:
	;
	goto L21
L21:
	;
	v50 = int32(_a_F_JsonTablePlanNextRow_0)
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_JsonTablePlanNextRow[0]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_JsonTablePlanNextRow[0])) = v53
	v55 = F_JsonbValueToJsonb(m, v28)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v55
	*(*int32)(unsafe.Add(mBase, _c_F_JsonTablePlanNextRow[0])) = v51
	v62 = int32(1)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v63 + v62
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v67 == v57 {
		v200 = v62
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v76 = v67
	v77 = l0 + int32(48)
	goto L24
L24:
	;
	F_JsonTableResetNestedPlan(m, v76)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L26
	}
L25:
	;
	v200 = v133
	goto L1
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v84 != int32(50) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v84 != int32(51) {
		v141 = v82
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+40)))
	if v89 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v180 = v82
	v182 = int32(1)
	goto L2
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v82)+32))
	if v100 != 0 {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
	if v90 == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v93 = F_JsonTablePlanNextRow(m, v90)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	if v93 == int32(0) {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v200 = int32(1)
	goto L1
L36:
	;
	if v99 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = v101
	v104 = v100 + int32(4)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if base.Ui32(v104) < base.Ui32(v107+v108<<(uint(int32(2))%32)) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = int32(0)
	goto L36
L40:
	;
	v113 = v104
	goto L42
L41:
	;
	v113 = int32(0)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+32)) = v113
	goto L36
L43:
	;
	v163 = v82
	goto L3
L44:
	;
	goto L45
L45:
	;
	v121 = int32(_a_F_JsonTablePlanNextRow_0)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_JsonTablePlanNextRow[0]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_JsonTablePlanNextRow[0])) = v124
	v126 = F_JsonbValueToJsonb(m, v99)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	v128 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v82)+40)) = uint8(v128)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v126
	*(*int32)(unsafe.Add(mBase, _c_F_JsonTablePlanNextRow[0])) = v122
	v133 = int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+44)) = v134 + v133
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
	if v140 != 0 {
		v76 = v140
		v77 = v82 + int32(48)
		goto L24
	} else {
		goto L47
	}
L47:
	;
	goto L25
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v153
	F_errmsg_internal(m, int32(_a_F_JsonTablePlanNextRow_1), v10)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_JsonTablePlanNextRow_2), int32(_a_F_JsonTablePlanNextRow_3), int32(_a_F_JsonTablePlanNextRow_4))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	if v188 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v180)+56))
	v193 = F_JsonTablePlanNextRow(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L9
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v200 = int32(1)
	goto L1
L55:
	;
	if v193 == int32(0) {
		v200 = v182
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L54
}
func F__equalJsonTablePath(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v48
L2:
	;
	return int32(0)
L3:
	;
	if v6 == int32(0) {
		v48 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v13 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v48 = int32(1)
	goto L1
L6:
	;
	if v12 == int32(0) {
		v48 = v3
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v13 != v12 {
		v48 = v3
		goto L1
	} else {
		goto L18
	}
L9:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if base.B2i32(v18 == int32(0))|base.B2i32(v18 != v21) != 0 {
		v39 = v18
		v40 = v21
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v39-v40 == int32(0) {
		goto L5
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v24 = v13
	v25 = v12
	goto L13
L13:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 == int32(0) {
		v39 = v29
		v40 = v28
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v39 = v29
	v40 = v28
	goto L11
L15:
	;
	v32 = int32(1)
	if v29 == v28 {
		v24 = v24 + v32
		v25 = v25 + v32
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v48 = v3
	goto L1
L18:
	;
	goto L5
}
func F_getJsonEncodingConst(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(_a_F_getJsonEncodingConst_0)
	v11 = F_palloc(m, int32(64))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l0 == int32(0) {
			v27 = v9
			v29 = F_strncpy(m, v11, v27, int32(64))
			mBase = m.M
			v30 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v29)+63)) = uint8(v30)
			v34 = int32(0)
			v38 = F_makeConst(m, int32(19), int32(-1), v34, int32(64), v11, v34, v34)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v38
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v17 == int32(0) {
				v27 = v9
				v29 = F_strncpy(m, v11, v27, int32(64))
				mBase = m.M
				v30 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+63)) = uint8(v30)
				v34 = int32(0)
				v38 = F_makeConst(m, int32(19), int32(-1), v34, int32(64), v11, v34, v34)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v38
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(int32(4)) <= base.Ui32(v20) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v20
						F_errmsg_internal(m, int32(_a_F_getJsonEncodingConst_1), v7)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_getJsonEncodingConst_2), int32(3274), int32(_a_F_getJsonEncodingConst_3))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v20<<(uint(int32(2))%32))+uint32(_c_F_getJsonEncodingConst[0])))
					v27 = v25
					v29 = F_strncpy(m, v11, v27, int32(64))
					mBase = m.M
					v30 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v29)+63)) = uint8(v30)
					v34 = int32(0)
					v38 = F_makeConst(m, int32(19), int32(-1), v34, int32(64), v11, v34, v34)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v38
					}
				}
			}
		}
	}
}
func F_get_json_behavior(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v10) < base.Ui32(int32(9)) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v10<<(uint(int32(2))%32))+uint32(_c_F_get_json_behavior[0])))
		F_appendStringInfoString(m, v13, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v19 == int32(8) {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				F_get_rule_expr(m, v22, l1, int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
					F_appendStringInfo(m, v26, int32(_a_F_get_json_behavior_0), v8+int32(16))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						m.G0 = v8 + int32(32)
						return
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
				F_appendStringInfo(m, v26, int32(_a_F_get_json_behavior_0), v8+int32(16))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					m.G0 = v8 + int32(32)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v40
			F_errmsg_internal(m, int32(_a_F_get_json_behavior_1), v8)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_get_json_behavior_2), int32(_a_F_get_json_behavior_3), int32(_a_F_get_json_behavior_4))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
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
func F_get_json_table(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_appendStringInfoString(m, v18, int32(_a_F_get_json_table_0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v22&int32(2) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v25 + int32(4)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v30 = int32(0)
	F_appendContextKeyword(m, l1, int32(_a_F_get_json_table_1), v30, v30, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_get_rule_expr(m, v35, l1, l2)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_appendStringInfoString(m, v18, int32(_a_F_get_json_table_2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	F_get_const_expr(m, v42, l1, int32(-1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v48 = F_quote_identifier(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v48
	F_appendStringInfo(m, v18, int32(_a_F_get_json_table_3), v14+int32(32))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v56 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_get_json_table_columns(m, l0, v204, l1, l2)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L46
	}
L13:
	;
	F_appendStringInfoChar(m, v18, int32(32))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v63 = int32(0)
	F_appendContextKeyword(m, l1, int32(_a_F_get_json_table_4), v63, v63, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v70 = v68 & int32(2)
	if v70 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v71 + int32(4)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if v76 == int32(0) {
		v83 = v4
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v75 == int32(0) {
		v179 = v70
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 <= int32(0) {
		v83 = v4
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v83 = v82
	goto L19
L22:
	;
	if v179 == int32(0) {
		goto L12
	} else {
		goto L45
	}
L23:
	;
	v86 = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if base.B2i32(v83 == v86)|base.B2i32(v88 <= v86) != 0 {
		v179 = v70
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	if v92 == int32(0) {
		v179 = v70
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v96 = int32(0)
	F_appendContextKeyword(m, l1, int32(_a_F_get_json_table_1), v96, v96, v96)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	F_get_rule_expr(m, v101, l1, int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	v107 = F_quote_identifier(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v107
	F_appendStringInfo(m, v18, int32(_a_F_get_json_table_3), v14+int32(16))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v119 = int32(1)
	goto L30
L30:
	;
	v127 = int32(0)
	if v76 == v127 {
		v137 = v127
		goto L33
	} else {
		goto L34
	}
L32:
	;
	F_appendStringInfoString(m, v18, int32(_a_F_get_json_table_2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L40
	}
L33:
	;
	v138 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if base.B2i32(v137 == v138)|base.B2i32(v140 <= v119) == v138 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v131 <= v119 {
		v137 = int32(0)
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v137 = v133 + v119<<(uint(int32(2))%32)
	goto L33
L36:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	if v145 != 0 {
		goto L32
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v179 = v147 & int32(2)
	goto L22
L39:
	;
	goto L38
L40:
	;
	v154 = int32(0)
	F_appendContextKeyword(m, l1, int32(_a_F_get_json_table_1), v154, v154, v154)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v145+v119<<(uint(int32(2))%32))))
	F_get_rule_expr(m, v162, l1, int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v168 = F_quote_identifier(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v168
	F_appendStringInfo(m, v18, int32(_a_F_get_json_table_3), v14)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v119 = v119 + int32(1)
	goto L30
L45:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v189 - int32(4)
	goto L12
L46:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if v208 != int32(6) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_get_json_behavior(m, v207, l1, int32(_a_F_get_json_table_5))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v214&int32(2) != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v217 - int32(4)
	goto L53
L52:
	;
	goto L53
L53:
	;
	v222 = int32(0)
	F_appendContextKeyword(m, l1, int32(_a_F_get_json_table_6), v222, v222, v222)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	m.G0 = v14 + int32(48)
	return
}
func F_get_json_table_columns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
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
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	v17 = m.G0
	v18 = int32(32)
	v19 = v17 - v18
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v21, v18)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v26 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_json_table_columns_0), v26, v26, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v31&int32(2) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v34 + int32(4)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v46 = int32(0)
	goto L7
L7:
	;
	v58 = int32(0)
	if v41 == v58 {
		v69 = v58
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v200 != 0 {
		goto L57
	} else {
		goto L58
	}
L9:
	;
	if v40 == int32(0) {
		v78 = v58
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v63 <= v46 {
		v69 = int32(0)
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v69 = v65 + v46<<(uint(int32(2))%32)
	goto L9
L12:
	;
	v79 = int32(0)
	if v39 == v79 {
		v90 = v79
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v72 <= v46 {
		v78 = v58
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v78 = v74 + v46<<(uint(int32(2))%32)
	goto L12
L15:
	;
	if v38 == int32(0) {
		v99 = v79
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v84 <= v46 {
		v90 = int32(0)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v90 = v86 + v46<<(uint(int32(2))%32)
	goto L15
L18:
	;
	v100 = int32(0)
	if base.B2i32(v69 == v100)|base.B2i32(v78 == v100)|(base.B2i32(v90 == v100)|base.B2i32(v99 == v100)) != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v93 <= v46 {
		v99 = v79
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v99 = v95 + v46<<(uint(int32(2))%32)
	goto L18
L21:
	;
	goto L8
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v112 = int32(0)
	if base.B2i32(v111 < v112)|base.B2i32(v46 < v111) == v112 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v118 < v46 {
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v46 = v46 + int32(1)
	goto L7
L26:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v111 < v46 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_appendStringInfoString(m, v21, int32(_a_F_get_json_table_columns_1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v130 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_json_table_columns_2), v130, v130, v130)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v135 = F_quote_identifier(m, v124)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v120 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	F_appendStringInfoString(m, v21, int32(_a_F_get_json_table_columns_3))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L49
	}
L34:
	;
	F_get_type_category_preferred(m, v122, v19+int32(31), v19+int32(30))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L43
	}
L35:
	;
	F_appendStringInfoString(m, v21, int32(_a_F_get_json_table_columns_4))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L42
	}
L36:
	;
	v137 = F_format_type_with_typemod(m, v122, v121)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(_a_F_get_json_table_columns_5)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v135
	F_appendStringInfo(m, v21, int32(_a_F_get_json_table_columns_6), v19)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v135
	F_appendStringInfo(m, v21, int32(_a_F_get_json_table_columns_6), v19+int32(16))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v146 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	switch v147 {
	case 0:
		goto L35
	case 1:
		goto L34
	default:
		v178 = v146
		goto L33
	}
L41:
	;
	v46 = v46 + int32(1)
	goto L7
L42:
	;
	v178 = int32(4)
	goto L33
L43:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+31)))
	if v166 != int32(83) {
		v178 = v146
		goto L33
	} else {
		goto L44
	}
L44:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	if v172 == int32(2) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v175 = int32(_a_F_get_json_table_columns_7)
	goto L47
L46:
	;
	v175 = int32(_a_F_get_json_table_columns_8)
	goto L47
L47:
	;
	F_appendStringInfoString(m, v21, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v178 = v146
	goto L33
L49:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v120)+20))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	if v183 == int32(7) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	F_get_json_expr_options(m, v120, l2, v178)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L56
	}
L51:
	;
	F_get_const_expr(m, v182, l2, int32(-1))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_get_rule_expr(m, v182, l2, l3)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	goto L50
L55:
	;
	goto L50
L56:
	;
	goto L25
L57:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_get_json_table_nested_columns(m, l0, v200, l2, l3, int32(base.Ui32(v201^int32(-1))>>(uint(int32(31))%32)))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v208&int32(2) != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v211 - int32(4)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v216 = int32(0)
	F_appendContextKeyword(m, l2, int32(_a_F_get_json_table_columns_9), v216, v216, v216)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	m.G0 = v19 + int32(32)
	return
}
func F_json_build_array_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	v9 = F_makeStringInfo(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_appendStringInfoChar(m, v9, int32(91))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < l0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = int32(_a_F_json_build_array_worker_0)
	v26 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_appendStringInfoChar(m, v9, int32(93))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	if l4 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v45 = v26 + int32(1)
	if v45 != l0 {
		v25 = v43
		v26 = v45
		goto L7
	} else {
		goto L16
	}
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v26))))
	if v28 != 0 {
		v43 = v25
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_appendStringInfoString(m, v9, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v32 = v26 << (uint(int32(2)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1+v32)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v26))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l3+v32)))
	F_add_json(m, v34, v36, v9, v38, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v43 = int32(_a_F_json_build_array_worker_1)
	goto L9
L16:
	;
	goto L8
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v60 = F_cstring_to_text_with_len(m, v58, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	return v60
}
func F_json_build_object_noargs(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_cstring_to_text_with_len(m, int32(_a_F_json_build_object_noargs_0), int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_json_build_object_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	v7 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	if l0&int32(1) == v7 {
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
		goto L6
	} else {
		goto L53
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L6
	} else {
		goto L49
	}
L3:
	;
	v24 = F_makeStringInfo(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L44
	}
L6:
	;
	return int32(0)
L7:
	;
	F_appendStringInfoChar(m, v24, int32(123))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if l5 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v31 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v18)+80)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = int64(51539607564)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_json_build_object_worker[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = int32(1305)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = int32(1306)
	v55 = F_hash_create(m, int32(_a_F_json_build_object_worker_0), int32(32), v18+int32(56), int32(1224))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if int32(0) < l0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_json_build_object_worker[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v61
	goto L11
L13:
	;
	v66 = v18 + int32(36)
	v77 = v7
	v81 = int32(_a_F_json_build_object_worker_1)
	goto L16
L14:
	;
	goto L15
L15:
	;
	F_appendStringInfoChar(m, v24, int32(125))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L6
	} else {
		goto L42
	}
L16:
	;
	if l4 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L15
L18:
	;
	v174 = v77 + int32(2)
	if v174 < l0 {
		v77 = v174
		v81 = v171
		goto L16
	} else {
		goto L41
	}
L19:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v77))))
	if v116 == int32(1) {
		goto L2
	} else {
		goto L30
	}
L20:
	;
	F_appendStringInfoString(m, v24, v81)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L29
	}
L21:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v77)+1)))
	if v86 != int32(1) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if l5 == int32(0) {
		v171 = v81
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v91 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v112 = int32(1)
	v113 = v81
	v114 = v66
	goto L19
L25:
	;
	v94 = int32(_a_F_json_build_object_worker_2)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_json_build_object_worker[0]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	*(*int32)(unsafe.Add(mBase, _c_F_json_build_object_worker[0])) = v97
	F_initStringInfo(m, v66)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = int32(0)
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_json_build_object_worker[0])) = v95
	goto L24
L29:
	;
	v112 = int32(0)
	v113 = int32(_a_F_json_build_object_worker_3)
	v114 = v24
	goto L19
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v121 = v77 << (uint(int32(2)) % 32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1+v121)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l3+v121)))
	F_add_json(m, v123, int32(0), v114, v126, int32(1))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	if l5 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v132 = F_pstrdup(m, v130+v119)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_appendStringInfoString(m, v24, int32(_a_F_json_build_object_worker_4))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L6
	} else {
		goto L39
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v132
	v135 = F_strlen(m, v132)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v135
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v145 = F_hash_search(m, v139, v18+int32(56), int32(1), v18+int32(111))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+111)))
	if v147 == int32(1) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v112 != 0 {
		v171 = v113
		goto L18
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v156 = v77 | int32(1)
	v158 = v156 << (uint(int32(2)) % 32)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1+v158)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v156))))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l3+v158)))
	F_add_json(m, v160, v162, v24, v164, int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v171 = v113
	goto L18
L41:
	;
	goto L17
L42:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v196 = F_cstring_to_text_with_len(m, v194, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	m.G0 = v18 + int32(112)
	return v196
L44:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(_a_F_json_build_object_worker_5), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(_a_F_json_build_object_worker_6)
	F_errhint(m, int32(_a_F_json_build_object_worker_7), v18+int32(16))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_json_build_object_worker_8), int32(1238), int32(_a_F_json_build_object_worker_9))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_json_build_object_worker_10), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_json_build_object_worker_8), int32(1275), int32(_a_F_json_build_object_worker_9))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errcode(m, int32(_a_F_json_build_object_worker_11))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v132
	F_errmsg(m, int32(_a_F_json_build_object_worker_12), v18)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_json_build_object_worker_8), int32(1297), int32(_a_F_json_build_object_worker_9))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_json_each_text(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	F_each_worker(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_json_errsave_error(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	switch l0 - int32(17) {
	case 0, 2, 3:
		goto L5
	default:
		goto L3
	case 6:
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return
L2:
	;
	v64 = F_json_errdetail(m, l0, l1)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L22
	}
L3:
	;
	v48 = F_errsave_start(m, l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L18
	}
L4:
	;
	if l2 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v17 = F_errsave_start(m, l2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	if v17 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_errcode(m, int32(84017282))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(_a_F_json_errsave_error_0), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v63 = int32(651)
	goto L2
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L15
	}
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v31 != int32(447) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v34 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_errmsg_internal(m, int32(_a_F_json_errsave_error_1), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_json_errsave_error_2), int32(656), int32(_a_F_json_errsave_error_3))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	if v48 == int32(0) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_json_errsave_error_4)
	F_errmsg(m, int32(_a_F_json_errsave_error_5), v13+int32(16))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v63 = int32(663)
	goto L2
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v64
	F_errdetail_internal(m, int32(_a_F_json_errsave_error_6), v13)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v70 = m.G0
	v72 = v70 - int32(16)
	m.G0 = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v74-v75 < int32(50) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v100-v75 < int32(4) {
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v100 = v75
	goto L24
L26:
	;
	goto L27
L27:
	;
	v79 = v75
	goto L28
L28:
	;
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79))))
	if v89 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v100 = v96
	goto L24
L30:
	;
	v92 = F_pg_mblen_range(m, v79, v74)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L33
	}
L31:
	;
	v95 = int32(1)
	goto L32
L32:
	;
	v96 = v95 + v79
	if int32(49) < v74-v96 {
		v79 = v96
		goto L28
	} else {
		goto L34
	}
L33:
	;
	v95 = v92
	goto L32
L34:
	;
	goto L29
L35:
	;
	v113 = v75
	goto L37
L36:
	;
	v113 = v100
	goto L37
L37:
	;
	v114 = v74 - v113
	v117 = F_palloc(m, v114+int32(1))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	if v114 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	base.MemoryCopy(m, v117, v113, v114)
	goto L41
L40:
	;
	goto L41
L41:
	;
	v121 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v114+v117))) = uint8(v121)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v124 == int32(12) {
		v142 = int32(_a_F_json_errsave_error_7)
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L49
	}
L43:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v128) <= base.Ui32(v74-v129) {
		v142 = int32(_a_F_json_errsave_error_7)
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v133 == int32(10) {
		v142 = int32(_a_F_json_errsave_error_7)
		goto L42
	} else {
		goto L45
	}
L45:
	;
	if v133 == int32(13) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v140 = int32(_a_F_json_errsave_error_7)
	goto L48
L47:
	;
	v140 = int32(_a_F_json_errsave_error_8)
	goto L48
L48:
	;
	v142 = v140
	goto L42
L49:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v117
	if base.Ui32(v75) < base.Ui32(v113) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v152 = int32(_a_F_json_errsave_error_8)
	goto L52
L51:
	;
	v152 = int32(_a_F_json_errsave_error_7)
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v146
	F_errcontext_msg(m, int32(_a_F_json_errsave_error_9), v72)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	m.G0 = v72 + int32(16)
	F_errsave_finish(m, l2, int32(_a_F_json_errsave_error_2), v63, int32(_a_F_json_errsave_error_3))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	goto L1
}
func F_json_extract_path(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_get_path_all(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_json_manifest_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v9 - int32(2) {
	case 0:
		goto L8
	default:
		goto L1
	case 5:
		goto L7
	case 9:
		goto L6
	}
L1:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_json_manifest_object_field_start_0)
	m.T0[v491].(func(*base.Module, int32, int32, int32))(m, v490, int32(_a_F_json_manifest_object_field_start_1), v7)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L60
	} else {
		goto L140
	}
L2:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = int32(_a_F_json_manifest_object_field_start_2)
	m.T0[v482].(func(*base.Module, int32, int32, int32))(m, v481, int32(_a_F_json_manifest_object_field_start_1), v7-int32(-64))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L60
	} else {
		goto L139
	}
L3:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v472)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(_a_F_json_manifest_object_field_start_3)
	m.T0[v473].(func(*base.Module, int32, int32, int32))(m, v472, int32(_a_F_json_manifest_object_field_start_1), v7+int32(48))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L60
	} else {
		goto L138
	}
L4:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_json_manifest_object_field_start_4)
	m.T0[v464].(func(*base.Module, int32, int32, int32))(m, v463, int32(_a_F_json_manifest_object_field_start_1), v7+int32(32))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L60
	} else {
		goto L137
	}
L5:
	;
	F_pfree(m, l1)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L60
	} else {
		goto L136
	}
L6:
	;
	v363 = int32(0)
	v364 = int32(_a_F_json_manifest_object_field_start_5)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[0])))
	if base.B2i32(v367 == v363)|base.B2i32(v367 != v370) != 0 {
		v388 = v367
		v389 = v370
		goto L113
	} else {
		goto L114
	}
L7:
	;
	v181 = int32(0)
	v182 = int32(_a_F_json_manifest_object_field_start_6)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[1])))
	if base.B2i32(v185 == v181)|base.B2i32(v185 != v188) != 0 {
		v206 = v185
		v207 = v188
		goto L64
	} else {
		goto L65
	}
L8:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
	if v12 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v15 = int32(_a_F_json_manifest_object_field_start_7)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[2])))
	if base.B2i32(v18 == int32(0))|base.B2i32(v18 != v21) != 0 {
		v39 = v18
		v40 = v21
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v46 = int32(_a_F_json_manifest_object_field_start_8)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[3])))
	if base.B2i32(v49 == int32(0))|base.B2i32(v49 != v52) != 0 {
		v70 = v49
		v71 = v52
		goto L21
	} else {
		goto L22
	}
L12:
	;
	if v39-v40 != 0 {
		goto L4
	} else {
		goto L19
	}
L13:
	;
	goto L12
L14:
	;
	v24 = l1
	v25 = v15
	goto L15
L15:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 == int32(0) {
		v39 = v29
		v40 = v28
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v39 = v29
	v40 = v28
	goto L13
L17:
	;
	v32 = int32(1)
	if v29 == v28 {
		v24 = v24 + v32
		v25 = v25 + v32
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v42 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v42)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(3)
	goto L5
L20:
	;
	if v70-v71 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	goto L20
L22:
	;
	v55 = l1
	v56 = v46
	goto L23
L23:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v60 == int32(0) {
		v70 = v60
		v71 = v59
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v70 = v60
	v71 = v59
	goto L21
L25:
	;
	v63 = int32(1)
	if v60 == v59 {
		v55 = v55 + v63
		v56 = v56 + v63
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(4)
	goto L5
L28:
	;
	goto L29
L29:
	;
	v77 = int32(_a_F_json_manifest_object_field_start_9)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[4])))
	if base.B2i32(v80 == int32(0))|base.B2i32(v80 != v83) != 0 {
		v101 = v80
		v102 = v83
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v101-v102 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	goto L30
L32:
	;
	v86 = l1
	v87 = v77
	goto L33
L33:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if v91 == int32(0) {
		v101 = v91
		v102 = v90
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v101 = v91
	v102 = v90
	goto L31
L35:
	;
	v94 = int32(1)
	if v91 == v90 {
		v86 = v86 + v94
		v87 = v87 + v94
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
	goto L5
L38:
	;
	goto L39
L39:
	;
	v108 = int32(_a_F_json_manifest_object_field_start_10)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[5])))
	if base.B2i32(v111 == int32(0))|base.B2i32(v111 != v114) != 0 {
		v132 = v111
		v133 = v114
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v132-v133 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v117 = l1
	v118 = v108
	goto L43
L43:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v122 == int32(0) {
		v132 = v122
		v133 = v121
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v132 = v122
	v133 = v121
	goto L41
L45:
	;
	v125 = int32(1)
	if v122 == v121 {
		v117 = v117 + v125
		v118 = v118 + v125
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
	goto L5
L48:
	;
	goto L49
L49:
	;
	v139 = int32(_a_F_json_manifest_object_field_start_11)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[6])))
	if base.B2i32(v142 == int32(0))|base.B2i32(v142 != v145) != 0 {
		v163 = v142
		v164 = v145
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v163-v164 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	goto L50
L52:
	;
	v148 = l1
	v149 = v139
	goto L53
L53:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	if v153 == int32(0) {
		v163 = v153
		v164 = v152
		goto L51
	} else {
		goto L55
	}
L54:
	;
	v163 = v153
	v164 = v152
	goto L51
L55:
	;
	v156 = int32(1)
	if v153 == v152 {
		v148 = v148 + v156
		v149 = v149 + v156
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(13)
	goto L5
L58:
	;
	goto L59
L59:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(_a_F_json_manifest_object_field_start_12)
	m.T0[v171].(func(*base.Module, int32, int32, int32))(m, v170, int32(_a_F_json_manifest_object_field_start_1), v7+int32(16))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	return int32(0)
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v359
	goto L5
L63:
	;
	if v206-v207 == int32(0) {
		v359 = v181
		goto L62
	} else {
		goto L70
	}
L64:
	;
	goto L63
L65:
	;
	v191 = l1
	v192 = v182
	goto L66
L66:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	if v196 == int32(0) {
		v206 = v196
		v207 = v195
		goto L64
	} else {
		goto L68
	}
L67:
	;
	v206 = v196
	v207 = v195
	goto L64
L68:
	;
	v199 = int32(1)
	if v196 == v195 {
		v191 = v191 + v199
		v192 = v192 + v199
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v212 = int32(_a_F_json_manifest_object_field_start_13)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[7])))
	if base.B2i32(v215 == int32(0))|base.B2i32(v215 != v218) != 0 {
		v236 = v215
		v237 = v218
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v236-v237 == int32(0) {
		v359 = int32(1)
		goto L62
	} else {
		goto L78
	}
L72:
	;
	goto L71
L73:
	;
	v221 = l1
	v222 = v212
	goto L74
L74:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	if v226 == int32(0) {
		v236 = v226
		v237 = v225
		goto L72
	} else {
		goto L76
	}
L75:
	;
	v236 = v226
	v237 = v225
	goto L72
L76:
	;
	v229 = int32(1)
	if v226 == v225 {
		v221 = v221 + v229
		v222 = v222 + v229
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v242 = int32(_a_F_json_manifest_object_field_start_14)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[8])))
	if base.B2i32(v245 == int32(0))|base.B2i32(v245 != v248) != 0 {
		v266 = v245
		v267 = v248
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v266-v267 == int32(0) {
		v359 = int32(2)
		goto L62
	} else {
		goto L86
	}
L80:
	;
	goto L79
L81:
	;
	v251 = l1
	v252 = v242
	goto L82
L82:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+1)))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+1)))
	if v256 == int32(0) {
		v266 = v256
		v267 = v255
		goto L80
	} else {
		goto L84
	}
L83:
	;
	v266 = v256
	v267 = v255
	goto L80
L84:
	;
	v259 = int32(1)
	if v256 == v255 {
		v251 = v251 + v259
		v252 = v252 + v259
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v272 = int32(_a_F_json_manifest_object_field_start_15)
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[9])))
	if base.B2i32(v275 == int32(0))|base.B2i32(v275 != v278) != 0 {
		v296 = v275
		v297 = v278
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v296-v297 == int32(0) {
		v359 = int32(3)
		goto L62
	} else {
		goto L94
	}
L88:
	;
	goto L87
L89:
	;
	v281 = l1
	v282 = v272
	goto L90
L90:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)))
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+1)))
	if v286 == int32(0) {
		v296 = v286
		v297 = v285
		goto L88
	} else {
		goto L92
	}
L91:
	;
	v296 = v286
	v297 = v285
	goto L88
L92:
	;
	v289 = int32(1)
	if v286 == v285 {
		v281 = v281 + v289
		v282 = v282 + v289
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v302 = int32(_a_F_json_manifest_object_field_start_16)
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[10])))
	if base.B2i32(v305 == int32(0))|base.B2i32(v305 != v308) != 0 {
		v326 = v305
		v327 = v308
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v326-v327 == int32(0) {
		v359 = int32(4)
		goto L62
	} else {
		goto L102
	}
L96:
	;
	goto L95
L97:
	;
	v311 = l1
	v312 = v302
	goto L98
L98:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+1)))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+1)))
	if v316 == int32(0) {
		v326 = v316
		v327 = v315
		goto L96
	} else {
		goto L100
	}
L99:
	;
	v326 = v316
	v327 = v315
	goto L96
L100:
	;
	v319 = int32(1)
	if v316 == v315 {
		v311 = v311 + v319
		v312 = v312 + v319
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v331 = int32(_a_F_json_manifest_object_field_start_17)
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[11])))
	if base.B2i32(v334 == int32(0))|base.B2i32(v334 != v337) != 0 {
		v355 = v334
		v356 = v337
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v355-v356 != 0 {
		goto L3
	} else {
		goto L110
	}
L104:
	;
	goto L103
L105:
	;
	v340 = l1
	v341 = v331
	goto L106
L106:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+1)))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+1)))
	if v345 == int32(0) {
		v355 = v345
		v356 = v344
		goto L104
	} else {
		goto L108
	}
L107:
	;
	v355 = v345
	v356 = v344
	goto L104
L108:
	;
	v348 = int32(1)
	if v345 == v344 {
		v340 = v340 + v348
		v341 = v341 + v348
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v359 = int32(5)
	goto L62
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v451
	goto L5
L112:
	;
	if v388-v389 == int32(0) {
		v451 = v363
		goto L111
	} else {
		goto L119
	}
L113:
	;
	goto L112
L114:
	;
	v373 = l1
	v374 = v364
	goto L115
L115:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+1)))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+1)))
	if v378 == int32(0) {
		v388 = v378
		v389 = v377
		goto L113
	} else {
		goto L117
	}
L116:
	;
	v388 = v378
	v389 = v377
	goto L113
L117:
	;
	v381 = int32(1)
	if v378 == v377 {
		v373 = v373 + v381
		v374 = v374 + v381
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v394 = int32(_a_F_json_manifest_object_field_start_18)
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[12])))
	if base.B2i32(v397 == int32(0))|base.B2i32(v397 != v400) != 0 {
		v418 = v397
		v419 = v400
		goto L121
	} else {
		goto L122
	}
L120:
	;
	if v418-v419 == int32(0) {
		v451 = int32(1)
		goto L111
	} else {
		goto L127
	}
L121:
	;
	goto L120
L122:
	;
	v403 = l1
	v404 = v394
	goto L123
L123:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+1)))
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+1)))
	if v408 == int32(0) {
		v418 = v408
		v419 = v407
		goto L121
	} else {
		goto L125
	}
L124:
	;
	v418 = v408
	v419 = v407
	goto L121
L125:
	;
	v411 = int32(1)
	if v408 == v407 {
		v403 = v403 + v411
		v404 = v404 + v411
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v423 = int32(_a_F_json_manifest_object_field_start_19)
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v429 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_json_manifest_object_field_start[13])))
	if base.B2i32(v426 == int32(0))|base.B2i32(v426 != v429) != 0 {
		v447 = v426
		v448 = v429
		goto L129
	} else {
		goto L130
	}
L128:
	;
	if v447-v448 != 0 {
		goto L2
	} else {
		goto L135
	}
L129:
	;
	goto L128
L130:
	;
	v432 = l1
	v433 = v423
	goto L131
L131:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+1)))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432)+1)))
	if v437 == int32(0) {
		v447 = v437
		v448 = v436
		goto L129
	} else {
		goto L133
	}
L132:
	;
	v447 = v437
	v448 = v436
	goto L129
L133:
	;
	v440 = int32(1)
	if v437 == v436 {
		v432 = v432 + v440
		v433 = v433 + v440
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	v451 = int32(2)
	goto L111
L136:
	;
	m.G0 = v7 + int32(80)
	return int32(0)
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_json_object_agg_transfn_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
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
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v13 = v10 + int32(36)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == v4 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L26
	} else {
		goto L82
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L26
	} else {
		goto L78
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L26
	} else {
		goto L74
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L26
	} else {
		goto L70
	}
L5:
	;
	if v43 != 0 {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v43 = v40
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v35
	v40 = v36
	goto L6
L8:
	;
	v32 = int32(0)
	if v13 == v32 {
		v40 = v32
		goto L6
	} else {
		goto L18
	}
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	switch v18 - int32(429) {
	case 0:
		goto L11
	case 1:
		goto L10
	default:
		goto L8
	}
L10:
	;
	if v13 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	if v13 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = int32(1)
	goto L5
L13:
	;
	goto L14
L14:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v35 = v25
	v36 = int32(1)
	goto L7
L15:
	;
	v43 = int32(2)
	goto L5
L16:
	;
	goto L17
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+368))
	v35 = v30
	v36 = int32(2)
	goto L7
L18:
	;
	v35 = v32
	v36 = v4
	goto L7
L19:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v44 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L26
	} else {
		goto L67
	}
L22:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v137 == int32(1) {
		goto L2
	} else {
		goto L41
	}
L23:
	;
	v47 = int32(_a_F_json_object_agg_transfn_worker_0)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_json_object_agg_transfn_worker[0]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_json_object_agg_transfn_worker[0])) = v50
	v53 = F_palloc(m, int32(44))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v134 = v133
	goto L22
L26:
	;
	return int32(0)
L27:
	;
	v57 = F_makeStringInfo(m)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v57
	if l2 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_json_object_agg_transfn_worker[0])) = v48
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v105 = F_get_fn_expr_argtype(m, v103, int32(1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L26
	} else {
		goto L34
	}
L30:
	;
	v60 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = int64(51539607564)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_json_object_agg_transfn_worker[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = int32(1305)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = int32(1306)
	v84 = F_hash_create(m, int32(_a_F_json_object_agg_transfn_worker_1), int32(32), v10+int32(40), int32(1224))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L26
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v93 = v53 + int32(20)
	v94 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+16)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v93)+8)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v93))) = v94
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v84
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_json_object_agg_transfn_worker[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v53)+40)) = v88
	goto L29
L34:
	;
	if v105 == int32(0) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_json_categorize_type(m, v105, int32(0), v53+int32(4), v53+int32(8))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L26
	} else {
		goto L36
	}
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v118 = F_get_fn_expr_argtype(m, v116, int32(2))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L26
	} else {
		goto L37
	}
L37:
	;
	if v118 == int32(0) {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	F_json_categorize_type(m, v118, int32(0), v53+int32(12), v53+int32(16))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L26
	} else {
		goto L39
	}
L39:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	F_appendStringInfoString(m, v129, int32(_a_F_json_object_agg_transfn_worker_5))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v134 = v53
	goto L22
L41:
	;
	if l1 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	m.G0 = v10 + int32(96)
	return v134
L43:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	F_datum_to_json_internal(m, v176, int32(0), v173, v178, v179, int32(1))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L26
	} else {
		goto L54
	}
L44:
	;
	v165 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	if v167 < int32(3) {
		v173 = v166
		v174 = v165
		goto L43
	} else {
		goto L52
	}
L45:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v142 != int32(1) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	if l2 == int32(0) {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v148 = v134 + int32(24)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
	if v149 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v152 = int32(_a_F_json_object_agg_transfn_worker_0)
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_json_object_agg_transfn_worker[0]))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v134)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_json_object_agg_transfn_worker[0])) = v155
	F_initStringInfo(m, v148)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L26
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+28)) = int32(0)
	v173 = v148
	v174 = int32(1)
	goto L43
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_json_object_agg_transfn_worker[0])) = v153
	v173 = v148
	v174 = int32(1)
	goto L43
L52:
	;
	F_appendStringInfoString(m, v166, int32(_a_F_json_object_agg_transfn_worker_10))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L26
	} else {
		goto L53
	}
L53:
	;
	v173 = v166
	v174 = v165
	goto L43
L54:
	;
	if l2 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v186 = F_MemoryContextStrdup(m, v183, v184+v175)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L26
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	F_appendStringInfoString(m, v206, int32(_a_F_json_object_agg_transfn_worker_9))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L26
	} else {
		goto L62
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v186
	v189 = F_strlen(m, v186)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v189
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v134)+20))
	v199 = F_hash_search(m, v193, v10+int32(40), int32(1), v10+int32(95))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L26
	} else {
		goto L59
	}
L59:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+95)))
	if v201 == int32(1) {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v174 != 0 {
		goto L42
	} else {
		goto L61
	}
L61:
	;
	goto L57
L62:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v210 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v213 = int32(0)
	goto L65
L64:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v213 = v212
	goto L65
L65:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v134)+16))
	F_datum_to_json_internal(m, v213, v210, v214, v215, v216, int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L26
	} else {
		goto L66
	}
L66:
	;
	goto L42
L67:
	;
	F_errmsg_internal(m, int32(_a_F_json_object_agg_transfn_worker_11), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L26
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_json_object_agg_transfn_worker_3), int32(1016), int32(_a_F_json_object_agg_transfn_worker_4))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L26
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
	v247 = m.ExcPending
	if v247 != 0 {
		goto L26
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(1)
	F_errmsg(m, int32(_a_F_json_object_agg_transfn_worker_2), v10)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L26
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_json_object_agg_transfn_worker_3), int32(1043), int32(_a_F_json_object_agg_transfn_worker_4))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L26
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L26
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(2)
	F_errmsg(m, int32(_a_F_json_object_agg_transfn_worker_2), v10+int32(16))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L26
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_json_object_agg_transfn_worker_3), int32(1053), int32(_a_F_json_object_agg_transfn_worker_4))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L26
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L26
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(_a_F_json_object_agg_transfn_worker_6), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L26
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_json_object_agg_transfn_worker_3), int32(1076), int32(_a_F_json_object_agg_transfn_worker_4))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L26
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errcode(m, int32(_a_F_json_object_agg_transfn_worker_7))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L26
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v186
	F_errmsg(m, int32(_a_F_json_object_agg_transfn_worker_8), v10+int32(32))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L26
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_json_object_agg_transfn_worker_3), int32(1126), int32(_a_F_json_object_agg_transfn_worker_4))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L26
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_json_object_two_arg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
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
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(1) < v20 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L89
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L85
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L81
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v20 != v23 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if v20 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	m.G0 = v10 + int32(48)
	return v229
L10:
	;
	v28 = F_cstring_to_text(m, int32(_a_F_json_object_two_arg_0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_deconstruct_array_builtin(m, v13, int32(25), v10+int32(28), v10+int32(20), v10+int32(12))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v229 = v28
	goto L9
L14:
	;
	F_deconstruct_array_builtin(m, v18, int32(25), v10+int32(24), v10+int32(16), v10+int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v48 != v49 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v52 = v10 + int32(32)
	F_initStringInfo(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_appendStringInfoChar(m, v52, int32(123))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if int32(0) < v58 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v62 = int32(0)
	goto L22
L20:
	;
	goto L21
L21:
	;
	F_appendStringInfoChar(m, v10+int32(32), int32(125))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L74
	}
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+v62))))
	if v71 == int32(1) {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	if v62 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_appendStringInfoString(m, v10+int32(32), int32(_a_F_json_object_two_arg_1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v80 = v62 << (uint(int32(2)) % 32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80+v81)))
	v84 = F_pg_detoast_datum_packed(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v118 = int32(1)
	if v86&v118 != 0 {
		goto L41
	} else {
		goto L42
	}
L30:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v86 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v92 == int32(18) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v103 = int32(1)
	if v86&v103 != 0 {
		v115 = int32(base.Ui32(v86)>>(uint(v103)%32)) - v103
		goto L29
	} else {
		goto L40
	}
L34:
	;
	v95 = int32(16)
	goto L36
L35:
	;
	v95 = int32(0)
	goto L36
L36:
	;
	if base.Ui32((v92-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v102 = int32(4)
	goto L39
L38:
	;
	v102 = v95
	goto L39
L39:
	;
	v115 = v102
	goto L29
L40:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v115 = int32(base.Ui32(v109)>>(uint(int32(2))%32)) - int32(4)
	goto L29
L41:
	;
	v122 = v118
	goto L43
L42:
	;
	v122 = int32(4)
	goto L43
L43:
	;
	F_escape_json_with_len(m, v10+int32(32), v84+v122, v115)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v84 != v83 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_pfree(m, v84)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v130 = v10 + int32(32)
	F_appendStringInfoString(m, v130, int32(_a_F_json_object_two_arg_2))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134+v62))))
	if v136 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v195 = v62 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v195 < v196 {
		v62 = v195
		goto L22
	} else {
		goto L73
	}
L51:
	;
	F_appendStringInfoString(m, v130, int32(_a_F_json_object_two_arg_3))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v142+v80)))
	v145 = F_pg_detoast_datum_packed(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L50
L55:
	;
	v179 = int32(1)
	if v147&v179 != 0 {
		goto L67
	} else {
		goto L68
	}
L56:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v147 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+1)))
	if v153 == int32(18) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v164 = int32(1)
	if v147&v164 != 0 {
		v176 = int32(base.Ui32(v147)>>(uint(v164)%32)) - v164
		goto L55
	} else {
		goto L66
	}
L60:
	;
	v156 = int32(16)
	goto L62
L61:
	;
	v156 = int32(0)
	goto L62
L62:
	;
	if base.Ui32((v153-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v163 = int32(4)
	goto L65
L64:
	;
	v163 = v156
	goto L65
L65:
	;
	v176 = v163
	goto L55
L66:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v176 = int32(base.Ui32(v170)>>(uint(int32(2))%32)) - int32(4)
	goto L55
L67:
	;
	v183 = v179
	goto L69
L68:
	;
	v183 = int32(4)
	goto L69
L69:
	;
	F_escape_json_with_len(m, v10+int32(32), v145+v183, v176)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v145 == v144 {
		goto L50
	} else {
		goto L71
	}
L71:
	;
	F_pfree(m, v145)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L50
L73:
	;
	goto L23
L74:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_pfree(m, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	F_pfree(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	F_pfree(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v224 = F_cstring_to_text_with_len(m, v222, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	F_pfree(m, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v229 = v224
	goto L9
L81:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_json_object_two_arg_4), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_json_object_two_arg_5), int32(1509), int32(_a_F_json_object_two_arg_6))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_json_object_two_arg_7), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_json_object_two_arg_5), int32(1520), int32(_a_F_json_object_two_arg_6))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(_a_F_json_object_two_arg_8), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_json_object_two_arg_5), int32(1531), int32(_a_F_json_object_two_arg_6))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_json_string_to_tsvector_byid(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v9
		v16 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v16
		v21 = v7 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v21
		F_iterate_json_values(m, v11, int32(2), v7+int32(24))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = F_make_tsvector(m, v21)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v30 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(32)
						return v28
					}
				} else {
					m.G0 = v7 + int32(32)
					return v28
				}
			}
		}
	}
}
func F_json_strip_nulls(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v18 == int32(2) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v24 = base.B2i32(v21 != int32(0))
		} else {
			v24 = int32(0)
		}
		v26 = F_palloc0(m, int32(12))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v29 = F_palloc0(m, int32(40))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v33 = F_pg_detoast_datum_packed(m, v14)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = int32(1)
					v36 = v33 + v35
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
					v41 = v39 & v35
					if v41 != 0 {
						v42 = v36
					} else {
						v42 = v33 + int32(4)
					}
					if v39 == int32(1) {
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
						if v48 == int32(18) {
							v51 = int32(16)
						} else {
							v51 = int32(0)
						}
						if base.Ui32((v48-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v58 = int32(4)
						} else {
							v58 = v51
						}
						v69 = v58
					} else {
						v59 = int32(1)
						if v41 != 0 {
							v69 = int32(base.Ui32(v39)>>(uint(v59)%32)) - v59
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
							v69 = int32(base.Ui32(v63)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v71 = *(*int32)(unsafe.Add(mBase, _c_F_json_strip_nulls[0]))
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
					v74 = F_makeJsonLexContextCstringLen(m, v11+int32(12), v42, v69, v72, int32(1))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v74
						v77 = F_makeStringInfo(m)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v26)+9)) = uint8(v24)
							v80 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v26)+8)) = uint8(v80)
							*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v77
							*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(1371)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = int32(1372)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = int32(1373)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = int32(1374)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = int32(1375)
							*(*int32)(unsafe.Add(mBase, uint32(v29))) = v26
							*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = int32(1376)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = int32(1377)
							v99 = v11 + int32(12)
							v100 = F_pg_parse_json(m, v99, v29)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								if v100 != 0 {
									F_json_errsave_error(m, v100, v99, int32(0))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
										v108 = F_cstring_to_text_with_len(m, v106, v107)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											m.G0 = v11 + int32(80)
											return v108
										}
									}
								} else {
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
									v108 = F_cstring_to_text_with_len(m, v106, v107)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(80)
										return v108
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
func F_makeJsonKeyValue(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_palloc0(m, int32(12))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(126)
		return v5
	}
}
func F_makeJsonLexContext(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v6 = F_pg_detoast_datum_packed(m, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = int32(1)
		v9 = v6 + v8
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		v14 = v12 & v8
		if v14 != 0 {
			v15 = v9
		} else {
			v15 = v6 + int32(4)
		}
		if v12 == int32(1) {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v21 == int32(18) {
				v24 = int32(16)
			} else {
				v24 = int32(0)
			}
			if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v31 = int32(4)
			} else {
				v31 = v24
			}
			v42 = v31
		} else {
			v32 = int32(1)
			if v14 != 0 {
				v42 = int32(base.Ui32(v12)>>(uint(v32)%32)) - v32
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v42 = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v44 = *(*int32)(unsafe.Add(mBase, _c_F_makeJsonLexContext[0]))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
		v46 = F_makeJsonLexContextCstringLen(m, l0, v15, v42, v45, l2)
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return
		} else {
			return
		}
	}
}
func F_makeJsonValueExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13945(m, l0, l1, l2, int32(44))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_transformJsonBehavior(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v31 int32
	_ = v31
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	if l2 == v6 {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v316 = F_format_type_be(m, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L14
	} else {
		goto L117
	}
L2:
	;
	if l2 != 0 {
		goto L113
	} else {
		goto L114
	}
L3:
	;
	v206 = int32(0)
	v207 = F_exprType(m, v202)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L14
	} else {
		goto L69
	}
L4:
	;
	v198 = F_makeConst(m, v194, int32(-1), int32(0), v192, v195, v193, v191)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L14
	} else {
		goto L68
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L14
	} else {
		goto L60
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L14
	} else {
		goto L55
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L14
	} else {
		goto L50
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L14
	} else {
		goto L45
	}
L9:
	;
	if v56|base.B2i32(v55 == int32(1)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L10:
	;
	v55 = l3
	v56 = v6
	v58 = int32(-1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v20 != int32(8) {
		v55 = v20
		v56 = v6
		v58 = v19
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v25 = F_transformExprRecurse(m, l0, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	v30 = F_ValidJsonBehaviorDefaultExpr(m, v25, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if v30 == int32(0) {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v34 = F_contain_var_clause(m, v25)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	if v34 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v36 = F_expression_returns_set(m, v25)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	if v36 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v38 = F_exprCollation(m, v25)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	if v38 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v42 = F_exprType(m, v25)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L14
	} else {
		goto L26
	}
L24:
	;
	v46 = v38
	goto L25
L25:
	;
	v47 = int32(8)
	v48 = int32(0)
	if base.B2i32(v23 == v48)|base.B2i32(v46 == v48) != 0 {
		v55 = v47
		v56 = v25
		v58 = v19
		goto L9
	} else {
		goto L28
	}
L26:
	;
	v44 = F_get_typcollation(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	v46 = v44
	goto L25
L28:
	;
	if v46 != v23 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v55 = v47
	v56 = v25
	v58 = v19
	goto L9
L30:
	;
	v64 = int32(3802)
	v65 = int32(-1)
	v66 = int32(0)
	switch v55 {
	case 0, 2, 5:
		goto L35
	default:
		goto L34
	case 3:
		goto L33
	case 4:
		goto L36
	case 6:
		goto L38
	case 7:
		goto L37
	case 8:
		v191 = v66
		v192 = v65
		v193 = v66
		v194 = v64
		v195 = v6
		goto L4
	}
L31:
	;
	goto L32
L32:
	;
	if v56 != 0 {
		v202 = v56
		goto L3
	} else {
		goto L44
	}
L33:
	;
	v98 = int32(1)
	v191 = v98
	v192 = v98
	v193 = v66
	v194 = int32(16)
	v195 = v98
	goto L4
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L14
	} else {
		goto L41
	}
L35:
	;
	v81 = int32(1)
	v191 = v81
	v192 = int32(4)
	v193 = v81
	v194 = int32(23)
	v195 = v6
	goto L4
L36:
	;
	v78 = int32(1)
	v191 = v78
	v192 = v78
	v193 = v66
	v194 = int32(16)
	v195 = v6
	goto L4
L37:
	;
	v76 = F_DirectFunctionCall1Coll(m, int32(486), int32(0), int32(_a_F_transformJsonBehavior_11))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L14
	} else {
		goto L40
	}
L38:
	;
	v71 = F_DirectFunctionCall1Coll(m, int32(486), int32(0), int32(_a_F_transformJsonBehavior_12))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	v191 = v66
	v192 = v65
	v193 = v66
	v194 = v64
	v195 = v71
	goto L4
L40:
	;
	v191 = v66
	v192 = v65
	v193 = v66
	v194 = v64
	v195 = v76
	goto L4
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v55
	F_errmsg_internal(m, int32(_a_F_transformJsonBehavior_8), v14)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_transformJsonBehavior_4), int32(_a_F_transformJsonBehavior_9), int32(_a_F_transformJsonBehavior_10))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v102 = int32(0)
	v301 = v102
	v303 = v102
	goto L2
L45:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L14
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_transformJsonBehavior_13), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L14
	} else {
		goto L47
	}
L47:
	;
	v115 = F_exprLocation(m, v25)
	mBase = m.M
	F_parser_errposition(m, l0, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L14
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_transformJsonBehavior_4), int32(_a_F_transformJsonBehavior_14), int32(_a_F_transformJsonBehavior_6))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L14
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L14
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(_a_F_transformJsonBehavior_15), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	v134 = F_exprLocation(m, v25)
	mBase = m.M
	F_parser_errposition(m, l0, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L14
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_transformJsonBehavior_4), int32(_a_F_transformJsonBehavior_16), int32(_a_F_transformJsonBehavior_6))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L14
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L14
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_transformJsonBehavior_17), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L14
	} else {
		goto L57
	}
L57:
	;
	v153 = F_exprLocation(m, v25)
	mBase = m.M
	F_parser_errposition(m, l0, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L14
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_transformJsonBehavior_4), int32(_a_F_transformJsonBehavior_18), int32(_a_F_transformJsonBehavior_6))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L14
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L14
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(_a_F_transformJsonBehavior_19), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L14
	} else {
		goto L62
	}
L62:
	;
	v172 = F_get_collation_name(m, v46)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	v174 = F_get_collation_name(m, v23)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L14
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v172
	F_errdetail(m, int32(_a_F_transformJsonBehavior_20), v12+int32(-16))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	v183 = F_exprLocation(m, v25)
	mBase = m.M
	F_parser_errposition(m, l0, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L14
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_transformJsonBehavior_4), int32(_a_F_transformJsonBehavior_21), int32(_a_F_transformJsonBehavior_6))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L14
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+28)) = v58
	v202 = v198
	goto L3
L69:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v207 == v209 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v301 = v206
	v303 = v202
	goto L2
L71:
	;
	goto L72
L72:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	if v211 == int32(7) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v285 = int32(-1)
	v286 = int32(0)
	if v55 == int32(3) {
		goto L107
	} else {
		goto L108
	}
L74:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v234 = F_TypeCategory(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L14
	} else {
		goto L88
	}
L75:
	;
	v228 = int32(1)
	v229 = F_exprType(m, v202)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L14
	} else {
		goto L86
	}
L76:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+24)))
	if v214 != 0 {
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v215 = F_exprType(m, v202)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L14
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	if v215 == int32(3802) {
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v219 = F_exprType(m, v202)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L14
	} else {
		goto L82
	}
L82:
	;
	if v219 != int32(16) {
		goto L74
	} else {
		goto L83
	}
L83:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v224 = F_getBaseType(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L14
	} else {
		goto L84
	}
L84:
	;
	if v224 == int32(23) {
		goto L74
	} else {
		goto L85
	}
L85:
	;
	goto L75
L86:
	;
	if v229 == int32(16) {
		goto L73
	} else {
		goto L87
	}
L87:
	;
	v301 = v228
	v303 = v202
	goto L2
L88:
	;
	v236 = F_exprType(m, v202)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L14
	} else {
		goto L89
	}
L89:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v240 = int32(1)
	if v234 == int32(86) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v245 = v240
	goto L92
L91:
	;
	v245 = int32(3)
	goto L92
L92:
	;
	if v234 == int32(83) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v248 = v240
	goto L95
L94:
	;
	v248 = v245
	goto L95
L95:
	;
	v250 = F_exprLocation(m, l2)
	mBase = m.M
	v251 = F_coerce_to_target_type(m, l0, v202, v236, v238, v239, v248, int32(1), v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L14
	} else {
		goto L96
	}
L96:
	;
	if v251 != 0 {
		v301 = v206
		v303 = v251
		goto L2
	} else {
		goto L97
	}
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L14
	} else {
		goto L98
	}
L98:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L14
	} else {
		goto L99
	}
L99:
	;
	v260 = F_exprType(m, v202)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L14
	} else {
		goto L100
	}
L100:
	;
	v262 = F_format_type_be(m, v260)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L14
	} else {
		goto L101
	}
L101:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v265 = F_format_type_be(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L14
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v262
	F_errmsg(m, int32(_a_F_transformJsonBehavior_2), v12+int32(-32))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L14
	} else {
		goto L103
	}
L103:
	;
	if v55 == int32(8) {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v276 = F_exprLocation(m, v202)
	mBase = m.M
	F_parser_errposition(m, l0, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L14
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_transformJsonBehavior_4), int32(_a_F_transformJsonBehavior_7), int32(_a_F_transformJsonBehavior_6))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L14
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	v294 = int32(_a_F_transformJsonBehavior_0)
	goto L109
L108:
	;
	v294 = int32(_a_F_transformJsonBehavior_1)
	goto L109
L109:
	;
	v295 = F_DirectFunctionCall1Coll(m, int32(486), v286, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L14
	} else {
		goto L110
	}
L110:
	;
	v297 = int32(0)
	v299 = F_makeConst(m, int32(3802), v285, v286, v285, v295, v297, v297)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L14
	} else {
		goto L111
	}
L111:
	;
	v301 = v228
	v303 = v299
	goto L2
L112:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v309)+12)) = uint8(v301)
	m.G0 = v14 - int32(-64)
	return v309
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v303
	v309 = l2
	goto L112
L114:
	;
	goto L115
L115:
	;
	v307 = F_makeJsonBehavior(m, v55, v303, v58)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L14
	} else {
		goto L116
	}
L116:
	;
	v309 = v307
	goto L112
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v316
	F_errhint(m, int32(_a_F_transformJsonBehavior_3), v12+int32(-48))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L14
	} else {
		goto L118
	}
L118:
	;
	v324 = F_exprLocation(m, v202)
	mBase = m.M
	F_parser_errposition(m, l0, v324)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L14
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_transformJsonBehavior_4), int32(_a_F_transformJsonBehavior_5), int32(_a_F_transformJsonBehavior_6))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L14
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_transformJsonReturning(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l1 != 0 {
		v12 = F_transformJsonOutput(m, l0, l1, int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if base.B2i32(v16 == int32(114))|base.B2i32(v16 == int32(3802)) != 0 {
				v64 = v12
				m.G0 = v9 + int32(16)
				return v64
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67141764))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						v30 = F_format_type_be(m, v29)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v30
							F_errmsg(m, int32(_a_F_transformJsonReturning_0), v9)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								F_errhint(m, int32(_a_F_transformJsonReturning_1), int32(0))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
									F_parser_errposition(m, l0, v42)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_transformJsonReturning_2), int32(_a_F_transformJsonReturning_3), int32(_a_F_transformJsonReturning_4))
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
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
	} else {
		v51 = F_palloc0(m, int32(16))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v51))) = int32(43)
			v58 = F_makeJsonFormat(m, int32(1), int32(0), int32(-1))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v51)+8)) = int64(-4294967182)
				*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v58
				v64 = v51
				m.G0 = v9 + int32(16)
				return v64
			}
		}
	}
}
func F_transformJsonValueExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
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
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v17 = F_transformExprRecurse(m, l0, v16)
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
	v21 = F_exprType(m, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v21 == int32(705) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = F_coerce_to_specific_type(m, l0, v17, int32(25), l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v28 = v17
	goto L6
L6:
	;
	v29 = F_exprType(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v28 = v26
	goto L6
L8:
	;
	v31 = F_exprLocation(m, v28)
	mBase = m.M
	F_get_type_category_preferred(m, v29, v14+int32(39), v14+int32(38))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v39 != 0 {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L106
	}
L11:
	;
	m.G0 = v14 + int32(48)
	return v208
L12:
	;
	if l4|l5|base.B2i32(v29 == int32(17)) != 0 {
		goto L64
	} else {
		goto L65
	}
L13:
	;
	if l4 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L14:
	;
	if l3 != 0 {
		v113 = l3
		goto L12
	} else {
		goto L59
	}
L15:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+39)))
	if v105 != int32(83) {
		goto L14
	} else {
		goto L58
	}
L16:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v29-int32(700)) {
		goto L15
	} else {
		goto L57
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L52
	}
L18:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v29 != int32(17) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if l5 != 0 {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v44 = v40
	goto L23
L22:
	;
	v44 = int32(0)
	goto L23
L23:
	;
	if v44 != 0 {
		goto L17
	} else {
		goto L24
	}
L24:
	;
	if v29 == int32(114) {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	if v29 == int32(3802) {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	v113 = v39
	goto L12
L27:
	;
	if v29 <= int32(1042) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	if base.B2i32(v29 == int32(114))|base.B2i32(v29 == int32(3802)) != 0 {
		goto L13
	} else {
		goto L51
	}
L30:
	;
	if base.B2i32(int32(1)<<(uint(v29)%32)&int32(45154304) == int32(0))|base.B2i32(base.Ui32(int32(25)) < base.Ui32(v29)) != 0 {
		goto L16
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v29 <= int32(1183) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v208 = v28
	goto L11
L34:
	;
	if base.Ui32(v29-int32(1082)) < base.Ui32(int32(2)) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	if v29 == int32(1184) {
		goto L44
	} else {
		goto L45
	}
L37:
	;
	v208 = v28
	goto L11
L38:
	;
	goto L39
L39:
	;
	if v29 == int32(1043) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v208 = v28
	goto L11
L41:
	;
	goto L42
L42:
	;
	if v29 != int32(1114) {
		goto L15
	} else {
		goto L43
	}
L43:
	;
	v208 = v28
	goto L11
L44:
	;
	v208 = v28
	goto L11
L45:
	;
	goto L46
L46:
	;
	if v29 == int32(1266) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v208 = v28
	goto L11
L48:
	;
	goto L49
L49:
	;
	if v29 != int32(1700) {
		goto L15
	} else {
		goto L50
	}
L50:
	;
	v208 = v28
	goto L11
L51:
	;
	goto L14
L52:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(_a_F_transformJsonValueExpr_0), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	F_parser_errposition(m, l0, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_transformJsonValueExpr_1), int32(3337), int32(_a_F_transformJsonValueExpr_2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	v208 = v28
	goto L11
L58:
	;
	v208 = v28
	goto L11
L59:
	;
	goto L13
L60:
	;
	v208 = v28
	goto L11
L61:
	;
	goto L62
L62:
	;
	if v29 != l4 {
		v113 = int32(0)
		goto L12
	} else {
		goto L63
	}
L63:
	;
	v208 = v28
	goto L11
L64:
	;
	if v113 != int32(1) {
		goto L79
	} else {
		goto L80
	}
L65:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+39)))
	if v118 == int32(83) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v131 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v132 = int32(_a_F_transformJsonValueExpr_4)
	goto L71
L70:
	;
	v132 = int32(_a_F_transformJsonValueExpr_5)
	goto L71
L71:
	;
	F_errmsg(m, v132, int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	if v137 < int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v140 = v31
	goto L75
L74:
	;
	v140 = v137
	goto L75
L75:
	;
	F_parser_errposition(m, l0, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_transformJsonValueExpr_1), int32(3403), int32(_a_F_transformJsonValueExpr_2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	v177 = base.B2i32(v113 == int32(2))
	if v113 == int32(2) {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	v172 = v29
	v173 = v28
	goto L78
L80:
	;
	goto L81
L81:
	;
	if v29 != int32(17) {
		v172 = v29
		v173 = v28
		goto L78
	} else {
		goto L82
	}
L82:
	;
	v152 = F_getJsonEncodingConst(m, v38)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v152
	v158 = int32(25)
	v165 = F_list_make2_impl(m, v14+int32(28), v14+int32(24))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v167 = int32(0)
	v169 = F_makeFuncExpr(m, int32(1714), v158, v165, v167, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169)+32)) = v31
	v172 = v158
	v173 = v169
	goto L78
L86:
	;
	v178 = int32(3802)
	goto L88
L87:
	;
	v178 = int32(114)
	goto L88
L88:
	;
	if l4 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v179 = l4
	goto L91
L90:
	;
	v179 = v178
	goto L91
L91:
	;
	v183 = F_coerce_to_target_type(m, l0, v173, v172, v179, int32(-1), int32(3), int32(1), v31)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v183 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if l4 != 0 {
		goto L10
	} else {
		goto L96
	}
L94:
	;
	v202 = v183
	goto L95
L95:
	;
	if v173 == v202 {
		goto L102
	} else {
		goto L103
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v173
	if v113 == int32(2) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v191 = int32(3787)
	goto L99
L98:
	;
	v191 = int32(3176)
	goto L99
L99:
	;
	v195 = F_list_make1_impl(m, int32(1), v14+int32(20))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v197 = int32(0)
	v199 = F_makeFuncExpr(m, v191, v178, v195, v197, v197)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+32)) = v31
	v202 = v199
	goto L95
L102:
	;
	v208 = v28
	goto L11
L103:
	;
	goto L104
L104:
	;
	v204 = F_copyObjectImpl(m, l2)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+8)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v204)+4)) = v28
	v208 = v204
	goto L11
L106:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v224 = F_format_type_be(m, v172)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v226 = F_format_type_be(m, l4)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v224
	F_errmsg(m, int32(_a_F_transformJsonValueExpr_3), v14)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_parser_errposition(m, l0, v31)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_transformJsonValueExpr_1), int32(3438), int32(_a_F_transformJsonValueExpr_2))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
