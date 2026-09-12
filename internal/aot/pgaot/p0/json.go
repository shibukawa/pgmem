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
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(2))%32))+uint32(_consts[324])))
	v8 = F_pstrdup(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_JsonTablePlanNextRow(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	switch v14 - int32(50) {
	case 0:
		goto L5
	case 1:
		v192 = l0
		v197 = v2
		goto L2
	default:
		v154 = l0
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v209 | v210
L2:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v192)+52))
	v201 = F_JsonTablePlanNextRow(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L9
	} else {
		goto L52
	}
L3:
	;
	v185 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v185
	v190 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v190)
	v209 = v185
	v210 = base.B2i32(v29 != v185)
	goto L1
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L48
	}
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v30 != 0 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v18 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v21 = F_JsonTablePlanNextRow(m, v18)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	if v21 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v209 = int32(1)
	v210 = v2
	goto L1
L12:
	;
	if v29 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v31
	v34 = v30 + int32(4)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if base.Ui32(v34) < base.Ui32(v37+v38<<(uint(int32(2))%32)) {
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
	v43 = v34
	goto L18
L17:
	;
	v43 = int32(0)
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v43
	goto L12
L19:
	;
	v177 = l0 + int32(36)
	v180 = l0 + int32(40)
	goto L3
L20:
	;
	goto L21
L21:
	;
	v56 = int32(4489440)
	v57 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v59
	v61 = F_JsonbValueToJsonb(m, v29)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v63 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v63)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v61
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v57
	v68 = int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v69 + v68
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v73 == v63 {
		v209 = v68
		v210 = int32(0)
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v82 = v73
	v86 = l0 + int32(48)
	goto L24
L24:
	;
	F_JsonTableResetNestedPlan(m, v82)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L26
	}
L25:
	;
	v209 = v145
	v210 = v145
	goto L1
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v91 != int32(50) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v91 != int32(51) {
		v154 = v89
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+40)))
	if v96 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v192 = v89
	v197 = int32(1)
	goto L2
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	if v108 != 0 {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)+48))
	if v97 == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v100 = F_JsonTablePlanNextRow(m, v97)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	if v100 == int32(0) {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v104 = int32(1)
	v209 = v104
	v210 = v104
	goto L1
L36:
	;
	if v107 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = v109
	v112 = v108 + int32(4)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if base.Ui32(v112) < base.Ui32(v115+v116<<(uint(int32(2))%32)) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = int32(0)
	goto L36
L40:
	;
	v121 = v112
	goto L42
L41:
	;
	v121 = int32(0)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v121
	goto L36
L43:
	;
	v177 = v89 + int32(36)
	v180 = v89 + int32(40)
	goto L3
L44:
	;
	goto L45
L45:
	;
	v133 = int32(4489440)
	v134 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v136
	v138 = F_JsonbValueToJsonb(m, v107)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+40)) = uint8(v140)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+36)) = v138
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v134
	v145 = int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v89)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+44)) = v146 + v145
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v89)+48))
	if v153 != 0 {
		v82 = v153
		v86 = v89 + int32(48)
		goto L24
	} else {
		goto L47
	}
L47:
	;
	goto L25
L48:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v167
	F_errmsg_internal(m, int32(473023), v11)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(499429), int32(4299), int32(31833))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
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
	v209 = int32(1)
	v210 = v197
	goto L1
L52:
	;
	if v201 != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v192)+56))
	v204 = F_JsonTablePlanNextRow(m, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	if v204 != 0 {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v209 = int32(0)
	v210 = v197
	goto L1
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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
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
	return v47
L2:
	;
	return int32(0)
L3:
	;
	if v6 == int32(0) {
		v47 = v3
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
	v47 = int32(1)
	goto L1
L6:
	;
	if v12 == int32(0) {
		v47 = v3
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
		v47 = v3
		goto L1
	} else {
		goto L19
	}
L9:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v19 == int32(0) {
		v38 = v18
		v39 = v19
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v39-v38 == int32(0) {
		goto L5
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	if v18 != v19 {
		v38 = v18
		v39 = v19
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v23 = v13
	v24 = v12
	goto L14
L14:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v28 == int32(0) {
		v38 = v27
		v39 = v28
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v38 = v27
	v39 = v28
	goto L11
L16:
	;
	v31 = int32(1)
	if v27 == v28 {
		v23 = v23 + v31
		v24 = v24 + v31
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v47 = v3
	goto L1
L19:
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(549035)
	v11 = F_palloc(m, int32(64))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l0 == int32(0) {
			v29 = v9
			v31 = F_strncpy(m, v11, v29, int32(64))
			mBase = m.M
			v32 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v31)+63)) = uint8(v32)
			v36 = int32(0)
			v40 = F_makeConst(m, int32(19), int32(-1), v36, int32(64), v11, v36, v36)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v40
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v17 == int32(0) {
				v29 = v9
				v31 = F_strncpy(m, v11, v29, int32(64))
				mBase = m.M
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v31)+63)) = uint8(v32)
				v36 = int32(0)
				v40 = F_makeConst(m, int32(19), int32(-1), v36, int32(64), v11, v36, v36)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v40
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if base.Ui32(int32(4)) <= base.Ui32(v20) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v20
						F_errmsg_internal(m, int32(482868), v7)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494448), int32(3274), int32(68391))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v20<<(uint(int32(2))%32))+uint32(_consts[241])))
					v29 = v27
					v31 = F_strncpy(m, v11, v29, int32(64))
					mBase = m.M
					v32 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v31)+63)) = uint8(v32)
					v36 = int32(0)
					v40 = F_makeConst(m, int32(19), int32(-1), v36, int32(64), v11, v36, v36)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v40
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v10) < base.Ui32(int32(9)) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v10<<(uint(int32(2))%32))+uint32(_consts[852])))
		F_appendStringInfoString(m, v13, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v21 == int32(8) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				F_get_rule_expr(m, v24, l1, int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
					F_appendStringInfo(m, v28, int32(197864), v8+int32(16))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						m.G0 = v8 + int32(32)
						return
					}
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l2
				F_appendStringInfo(m, v28, int32(197864), v8+int32(16))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
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
		v41 = m.ExcPending
		if v41 != 0 {
			return
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v42
			F_errmsg_internal(m, int32(484435), v8)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				F_errfinish(m, int32(493345), int32(9194), int32(212495))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
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
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_appendStringInfoString(m, v18, int32(669975))
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
	F_appendContextKeyword(m, l1, int32(741336), v30, v30, v30)
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
	F_appendStringInfoString(m, v18, int32(730094))
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
	F_appendStringInfo(m, v18, int32(197810), v14+int32(32))
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
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_get_json_table_columns(m, l0, v200, l1, l2)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L47
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
	F_appendContextKeyword(m, l1, int32(728697), v63, v63, v63)
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
		v175 = v70
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
	if v175 == int32(0) {
		goto L12
	} else {
		goto L46
	}
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v86 <= int32(0) {
		v175 = v70
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v83 == int32(0) {
		v175 = v70
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	if v91 == int32(0) {
		v175 = v70
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v95 = int32(0)
	F_appendContextKeyword(m, l1, int32(741336), v95, v95, v95)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	F_get_rule_expr(m, v100, l1, int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v106 = F_quote_identifier(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v106
	F_appendStringInfo(m, v18, int32(197810), v14+int32(16))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v118 = int32(1)
	goto L31
L31:
	;
	v126 = int32(0)
	if v76 == v126 {
		v136 = v126
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v137 <= v118 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v130 <= v118 {
		v136 = int32(0)
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v136 = v132 + v118<<(uint(int32(2))%32)
	goto L33
L36:
	;
	F_appendStringInfoString(m, v18, int32(730094))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L41
	}
L37:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v175 = v146 & int32(2)
	goto L22
L38:
	;
	if v136 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v144 = v141 + v118<<(uint(int32(2))%32)
	if v144 != 0 {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	v153 = int32(0)
	F_appendContextKeyword(m, l1, int32(741336), v153, v153, v153)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	F_get_rule_expr(m, v158, l1, int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v164 = F_quote_identifier(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v164
	F_appendStringInfo(m, v18, int32(197810), v14)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v118 = v118 + int32(1)
	goto L31
L46:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v185 - int32(4)
	goto L12
L47:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	if v204 != int32(6) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_get_json_behavior(m, v203, l1, int32(524317))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v210&int32(2) != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v213 - int32(4)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v218 = int32(0)
	F_appendContextKeyword(m, l1, int32(669555), v218, v218, v218)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	m.G0 = v14 + int32(48)
	return
}
func F_get_json_table_columns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
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
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	v18 = m.G0
	v19 = int32(32)
	v20 = v18 - v19
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_appendStringInfoChar(m, v22, v19)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = int32(0)
	F_appendContextKeyword(m, l2, int32(670213), v27, v27, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v32&int32(2) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v35 + int32(4)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v47 = int32(0)
	goto L7
L7:
	;
	v60 = int32(0)
	if v42 == v60 {
		v71 = v60
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v196 != 0 {
		goto L60
	} else {
		goto L61
	}
L9:
	;
	if v41 == int32(0) {
		v80 = v60
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v65 <= v47 {
		v71 = int32(0)
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v71 = v67 + v47<<(uint(int32(2))%32)
	goto L9
L12:
	;
	v81 = int32(0)
	if v40 == v81 {
		v92 = v81
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v74 <= v47 {
		v80 = v60
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v80 = v76 + v47<<(uint(int32(2))%32)
	goto L12
L15:
	;
	if v39 == int32(0) {
		v101 = v81
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v86 <= v47 {
		v92 = int32(0)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v92 = v88 + v47<<(uint(int32(2))%32)
	goto L15
L18:
	;
	if v71 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v95 <= v47 {
		v101 = v81
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v101 = v97 + v47<<(uint(int32(2))%32)
	goto L18
L21:
	;
	goto L8
L22:
	;
	if v80 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if v92 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	if v101 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v110 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v47 = v47 + int32(1)
	goto L7
L27:
	;
	if v47 < v110 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v114 < v47 {
		goto L21
	} else {
		goto L29
	}
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v110 < v47 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_appendStringInfoString(m, v22, int32(730094))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v126 = int32(0)
	F_appendContextKeyword(m, l2, int32(741336), v126, v126, v126)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	v131 = F_quote_identifier(m, v120)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v116 != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	F_appendStringInfoString(m, v22, int32(728647))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L52
	}
L37:
	;
	F_get_type_category_preferred(m, v118, v20+int32(31), v20+int32(30))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L46
	}
L38:
	;
	F_appendStringInfoString(m, v22, int32(521855))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L45
	}
L39:
	;
	v133 = F_format_type_with_typemod(m, v118, v117)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(507846)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v131
	F_appendStringInfo(m, v22, int32(180419), v20)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v131
	F_appendStringInfo(m, v22, int32(180419), v20+int32(16))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v142 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	switch v143 {
	case 0:
		goto L38
	case 1:
		goto L37
	default:
		v174 = v142
		goto L36
	}
L44:
	;
	v47 = v47 + int32(1)
	goto L7
L45:
	;
	v174 = int32(4)
	goto L36
L46:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+31)))
	if v162 != int32(83) {
		v174 = v142
		goto L36
	} else {
		goto L47
	}
L47:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v168 == int32(2) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v171 = int32(544145)
	goto L50
L49:
	;
	v171 = int32(527756)
	goto L50
L50:
	;
	F_appendStringInfoString(m, v22, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v174 = v142
	goto L36
L52:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v116)+20))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if v179 == int32(7) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_get_json_expr_options(m, v116, l2, v174)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L59
	}
L54:
	;
	F_get_const_expr(m, v178, l2, int32(-1))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	F_get_rule_expr(m, v178, l2, l3)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	goto L53
L58:
	;
	goto L53
L59:
	;
	goto L26
L60:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_get_json_table_nested_columns(m, l0, v196, l2, l3, int32(base.Ui32(v197^int32(-1))>>(uint(int32(31))%32)))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v204&int32(2) != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	goto L62
L64:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v207 - int32(4)
	goto L66
L65:
	;
	goto L66
L66:
	;
	v212 = int32(0)
	F_appendContextKeyword(m, l2, int32(669555), v212, v212, v212)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	m.G0 = v20 + int32(32)
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
	v25 = int32(0)
	v26 = int32(741336)
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
	v45 = v25 + int32(1)
	if v45 != l0 {
		v25 = v45
		v26 = v43
		goto L7
	} else {
		goto L16
	}
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v25))))
	if v28 != 0 {
		v43 = v26
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_appendStringInfoString(m, v9, v26)
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
	v32 = v25 << (uint(int32(2)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1+v32)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v25))))
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
	v43 = int32(730094)
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
	v4 = F_cstring_to_text_with_len(m, int32(4103), int32(2))
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
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
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
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	if l0&int32(1) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L6
	} else {
		goto L53
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
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
	v214 = m.ExcPending
	if v214 != 0 {
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
	v32 = v18 + int32(96)
	v33 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = v33
	v36 = v18 + int32(80)
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v18-int32(-64)))) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = int64(51539607564)
	v48 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(1321)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = int32(1322)
	v61 = F_hash_create(m, int32(391332), int32(32), v18+int32(56), int32(1224))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v67
	goto L11
L13:
	;
	v74 = v18 + int32(36)
	v85 = int32(0)
	v89 = int32(741336)
	goto L16
L14:
	;
	goto L15
L15:
	;
	F_appendStringInfoChar(m, v24, int32(125))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
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
	v183 = v85 + int32(2)
	if v183 < l0 {
		v85 = v183
		v89 = v179
		goto L16
	} else {
		goto L41
	}
L19:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v85))))
	if v125 == int32(1) {
		goto L2
	} else {
		goto L30
	}
L20:
	;
	F_appendStringInfoString(m, v24, v89)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L29
	}
L21:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v85)+1)))
	if v95 != int32(1) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if l5 == int32(0) {
		v179 = v89
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v100 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v121 = v89
	v122 = int32(1)
	v123 = v74
	goto L19
L25:
	;
	v103 = int32(4489440)
	v104 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v106
	F_initStringInfo(m, v74)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v104
	goto L24
L29:
	;
	v121 = int32(730094)
	v122 = int32(0)
	v123 = v24
	goto L19
L30:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	v130 = v85 << (uint(int32(2)) % 32)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1+v130)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l3+v130)))
	F_add_json(m, v132, int32(0), v123, v135, int32(1))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
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
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	v141 = F_pstrdup(m, v139+v128)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_appendStringInfoString(m, v24, int32(729850))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L39
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v141
	v144 = F_strlen(m, v141)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v144
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v154 = F_hash_search(m, v148, v18+int32(56), int32(1), v18+int32(111))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+111)))
	if v156 == int32(1) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v122 != 0 {
		v179 = v121
		goto L18
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v165 = v85 | int32(1)
	v167 = v165 << (uint(int32(2)) % 32)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1+v167)))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v165))))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l3+v167)))
	F_add_json(m, v169, v171, v24, v173, int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v179 = v121
	goto L18
L41:
	;
	goto L17
L42:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v205 = F_cstring_to_text_with_len(m, v203, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	m.G0 = v18 + int32(112)
	return v205
L44:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(122915), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(667448)
	F_errhint(m, int32(580549), v18+int32(16))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(495612), int32(1238), int32(219734))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
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
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(21521), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(495612), int32(1275), int32(219734))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
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
	F_errcode(m, int32(786562))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v141
	F_errmsg(m, int32(202794), v18)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(495612), int32(1297), int32(219734))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
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
	F_errmsg(m, int32(415563), int32(0))
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
	F_errmsg_internal(m, int32(261419), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(494070), int32(656), int32(211727))
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(245140)
	F_errmsg(m, int32(189012), v13+int32(16))
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
	F_errdetail_internal(m, int32(205923), v13)
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
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v120+v114))) = uint8(v122)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v125 == int32(12) {
		v143 = int32(741336)
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v119 = F__emscripten_memcpy_bulkmem(m, v117, v113, v114)
	mBase = m.M
	v120 = v119
	goto L42
L41:
	;
	v120 = v117
	goto L42
L42:
	;
	goto L39
L43:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L50
	}
L44:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v129) <= base.Ui32(v74-v130) {
		v143 = int32(741336)
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v134 == int32(10) {
		v143 = int32(741336)
		goto L43
	} else {
		goto L46
	}
L46:
	;
	if v134 == int32(13) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v141 = int32(741336)
	goto L49
L48:
	;
	v141 = int32(644708)
	goto L49
L49:
	;
	v143 = v141
	goto L43
L50:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v120
	if base.Ui32(v75) < base.Ui32(v113) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v153 = int32(644708)
	goto L53
L52:
	;
	v153 = int32(741336)
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v147
	F_errcontext_msg(m, int32(175425), v72)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	m.G0 = v72 + int32(16)
	F_errsave_finish(m, l2, int32(494070), v63, int32(211727))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
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
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
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
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
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
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(431151)
	m.T0[v477].(func(*base.Module, int32, int32, int32))(m, v476, int32(199215), v7)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L65
	} else {
		goto L154
	}
L2:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = int32(431276)
	m.T0[v468].(func(*base.Module, int32, int32, int32))(m, v467, int32(199215), v7-int32(-64))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L65
	} else {
		goto L153
	}
L3:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v458)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(431254)
	m.T0[v459].(func(*base.Module, int32, int32, int32))(m, v458, int32(199215), v7+int32(48))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L65
	} else {
		goto L152
	}
L4:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(209654)
	m.T0[v450].(func(*base.Module, int32, int32, int32))(m, v449, int32(199215), v7+int32(32))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L65
	} else {
		goto L151
	}
L5:
	;
	F_pfree(m, l1)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L65
	} else {
		goto L150
	}
L6:
	;
	v352 = int32(0)
	v353 = int32(372957)
	v356 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1001])))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v357 == v352 {
		v376 = v356
		v377 = v357
		goto L124
	} else {
		goto L125
	}
L7:
	;
	v176 = int32(0)
	v177 = int32(321661)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1002])))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v181 == v176 {
		v200 = v180
		v201 = v181
		goto L69
	} else {
		goto L70
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
	v15 = int32(271022)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1003])))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v19 == int32(0) {
		v38 = v18
		v39 = v19
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v45 = int32(222156)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1004])))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v49 == int32(0) {
		v68 = v48
		v69 = v49
		goto L22
	} else {
		goto L23
	}
L12:
	;
	if v39-v38 != 0 {
		goto L4
	} else {
		goto L20
	}
L13:
	;
	goto L12
L14:
	;
	if v18 != v19 {
		v38 = v18
		v39 = v19
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v23 = l1
	v24 = v15
	goto L16
L16:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v28 == int32(0) {
		v38 = v27
		v39 = v28
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v38 = v27
	v39 = v28
	goto L13
L18:
	;
	v31 = int32(1)
	if v27 == v28 {
		v23 = v23 + v31
		v24 = v24 + v31
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v41 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v41)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(3)
	goto L5
L21:
	;
	if v69-v68 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	goto L21
L23:
	;
	if v48 != v49 {
		v68 = v48
		v69 = v49
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v53 = l1
	v54 = v45
	goto L25
L25:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v58 == int32(0) {
		v68 = v57
		v69 = v58
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v68 = v57
	v69 = v58
	goto L22
L27:
	;
	v61 = int32(1)
	if v57 == v58 {
		v53 = v53 + v61
		v54 = v54 + v61
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(4)
	goto L5
L30:
	;
	goto L31
L31:
	;
	v75 = int32(165010)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1005])))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v79 == int32(0) {
		v98 = v78
		v99 = v79
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v99-v98 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	goto L32
L34:
	;
	if v78 != v79 {
		v98 = v78
		v99 = v79
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v83 = l1
	v84 = v75
	goto L36
L36:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v88 == int32(0) {
		v98 = v87
		v99 = v88
		goto L33
	} else {
		goto L38
	}
L37:
	;
	v98 = v87
	v99 = v88
	goto L33
L38:
	;
	v91 = int32(1)
	if v87 == v88 {
		v83 = v83 + v91
		v84 = v84 + v91
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(5)
	goto L5
L41:
	;
	goto L42
L42:
	;
	v105 = int32(169536)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1006])))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v109 == int32(0) {
		v128 = v108
		v129 = v109
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v129-v128 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L44:
	;
	goto L43
L45:
	;
	if v108 != v109 {
		v128 = v108
		v129 = v109
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v113 = l1
	v114 = v105
	goto L47
L47:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	if v118 == int32(0) {
		v128 = v117
		v129 = v118
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v128 = v117
	v129 = v118
	goto L44
L49:
	;
	v121 = int32(1)
	if v117 == v118 {
		v113 = v113 + v121
		v114 = v114 + v121
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
	goto L5
L52:
	;
	goto L53
L53:
	;
	v135 = int32(286188)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1007])))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v139 == int32(0) {
		v158 = v138
		v159 = v139
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v159-v158 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L55:
	;
	goto L54
L56:
	;
	if v138 != v139 {
		v158 = v138
		v159 = v139
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v143 = l1
	v144 = v135
	goto L58
L58:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)))
	if v148 == int32(0) {
		v158 = v147
		v159 = v148
		goto L55
	} else {
		goto L60
	}
L59:
	;
	v158 = v147
	v159 = v148
	goto L55
L60:
	;
	v151 = int32(1)
	if v147 == v148 {
		v143 = v143 + v151
		v144 = v144 + v151
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(13)
	goto L5
L63:
	;
	goto L64
L64:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(431175)
	m.T0[v166].(func(*base.Module, int32, int32, int32))(m, v165, int32(199215), v7+int32(16))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	return int32(0)
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v348
	goto L5
L68:
	;
	if v201-v200 == int32(0) {
		v348 = v176
		goto L67
	} else {
		goto L76
	}
L69:
	;
	goto L68
L70:
	;
	if v180 != v181 {
		v200 = v180
		v201 = v181
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v185 = l1
	v186 = v177
	goto L72
L72:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+1)))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+1)))
	if v190 == int32(0) {
		v200 = v189
		v201 = v190
		goto L69
	} else {
		goto L74
	}
L73:
	;
	v200 = v189
	v201 = v190
	goto L69
L74:
	;
	v193 = int32(1)
	if v189 == v190 {
		v185 = v185 + v193
		v186 = v186 + v193
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v206 = int32(321653)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1008])))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v210 == int32(0) {
		v229 = v209
		v230 = v210
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v230-v229 == int32(0) {
		v348 = int32(1)
		goto L67
	} else {
		goto L85
	}
L78:
	;
	goto L77
L79:
	;
	if v209 != v210 {
		v229 = v209
		v230 = v210
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v214 = l1
	v215 = v206
	goto L81
L81:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)))
	if v219 == int32(0) {
		v229 = v218
		v230 = v219
		goto L78
	} else {
		goto L83
	}
L82:
	;
	v229 = v218
	v230 = v219
	goto L78
L83:
	;
	v222 = int32(1)
	if v218 == v219 {
		v214 = v214 + v222
		v215 = v215 + v222
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	v235 = int32(341891)
	v238 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1009])))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v239 == int32(0) {
		v258 = v238
		v259 = v239
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v259-v258 == int32(0) {
		v348 = int32(2)
		goto L67
	} else {
		goto L94
	}
L87:
	;
	goto L86
L88:
	;
	if v238 != v239 {
		v258 = v238
		v259 = v239
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v243 = l1
	v244 = v235
	goto L90
L90:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+1)))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)))
	if v248 == int32(0) {
		v258 = v247
		v259 = v248
		goto L87
	} else {
		goto L92
	}
L91:
	;
	v258 = v247
	v259 = v248
	goto L87
L92:
	;
	v251 = int32(1)
	if v247 == v248 {
		v243 = v243 + v251
		v244 = v244 + v251
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v264 = int32(456244)
	v267 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1010])))
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v268 == int32(0) {
		v287 = v267
		v288 = v268
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v288-v287 == int32(0) {
		v348 = int32(3)
		goto L67
	} else {
		goto L103
	}
L96:
	;
	goto L95
L97:
	;
	if v267 != v268 {
		v287 = v267
		v288 = v268
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v272 = l1
	v273 = v264
	goto L99
L99:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+1)))
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	if v277 == int32(0) {
		v287 = v276
		v288 = v277
		goto L96
	} else {
		goto L101
	}
L100:
	;
	v287 = v276
	v288 = v277
	goto L96
L101:
	;
	v280 = int32(1)
	if v276 == v277 {
		v272 = v272 + v280
		v273 = v273 + v280
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v293 = int32(288543)
	v296 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1011])))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v297 == int32(0) {
		v316 = v296
		v317 = v297
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v317-v316 == int32(0) {
		v348 = int32(4)
		goto L67
	} else {
		goto L112
	}
L105:
	;
	goto L104
L106:
	;
	if v296 != v297 {
		v316 = v296
		v317 = v297
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v301 = l1
	v302 = v293
	goto L108
L108:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+1)))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	if v306 == int32(0) {
		v316 = v305
		v317 = v306
		goto L105
	} else {
		goto L110
	}
L109:
	;
	v316 = v305
	v317 = v306
	goto L105
L110:
	;
	v309 = int32(1)
	if v305 == v306 {
		v301 = v301 + v309
		v302 = v302 + v309
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v321 = int32(286197)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1012])))
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v325 == int32(0) {
		v344 = v324
		v345 = v325
		goto L114
	} else {
		goto L115
	}
L113:
	;
	if v345-v344 != 0 {
		goto L3
	} else {
		goto L121
	}
L114:
	;
	goto L113
L115:
	;
	if v324 != v325 {
		v344 = v324
		v345 = v325
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v329 = l1
	v330 = v321
	goto L117
L117:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+1)))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+1)))
	if v334 == int32(0) {
		v344 = v333
		v345 = v334
		goto L114
	} else {
		goto L119
	}
L118:
	;
	v344 = v333
	v345 = v334
	goto L114
L119:
	;
	v337 = int32(1)
	if v333 == v334 {
		v329 = v329 + v337
		v330 = v330 + v337
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v348 = int32(5)
	goto L67
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v437
	goto L5
L123:
	;
	if v377-v376 == int32(0) {
		v437 = v352
		goto L122
	} else {
		goto L131
	}
L124:
	;
	goto L123
L125:
	;
	if v356 != v357 {
		v376 = v356
		v377 = v357
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v361 = l1
	v362 = v353
	goto L127
L127:
	;
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+1)))
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+1)))
	if v366 == int32(0) {
		v376 = v365
		v377 = v366
		goto L124
	} else {
		goto L129
	}
L128:
	;
	v376 = v365
	v377 = v366
	goto L124
L129:
	;
	v369 = int32(1)
	if v365 == v366 {
		v361 = v361 + v369
		v362 = v362 + v369
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v382 = int32(527517)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1013])))
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v386 == int32(0) {
		v405 = v385
		v406 = v386
		goto L133
	} else {
		goto L134
	}
L132:
	;
	if v406-v405 == int32(0) {
		v437 = int32(1)
		goto L122
	} else {
		goto L140
	}
L133:
	;
	goto L132
L134:
	;
	if v385 != v386 {
		v405 = v385
		v406 = v386
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v390 = l1
	v391 = v382
	goto L136
L136:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+1)))
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+1)))
	if v395 == int32(0) {
		v405 = v394
		v406 = v395
		goto L133
	} else {
		goto L138
	}
L137:
	;
	v405 = v394
	v406 = v395
	goto L133
L138:
	;
	v398 = int32(1)
	if v394 == v395 {
		v390 = v390 + v398
		v391 = v391 + v398
		goto L136
	} else {
		goto L139
	}
L139:
	;
	goto L137
L140:
	;
	v410 = int32(527527)
	v413 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1014])))
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v414 == int32(0) {
		v433 = v413
		v434 = v414
		goto L142
	} else {
		goto L143
	}
L141:
	;
	if v434-v433 != 0 {
		goto L2
	} else {
		goto L149
	}
L142:
	;
	goto L141
L143:
	;
	if v413 != v414 {
		v433 = v413
		v434 = v414
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v418 = l1
	v419 = v410
	goto L145
L145:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+1)))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)))
	if v423 == int32(0) {
		v433 = v422
		v434 = v423
		goto L142
	} else {
		goto L147
	}
L146:
	;
	v433 = v422
	v434 = v423
	goto L142
L147:
	;
	v426 = int32(1)
	if v422 == v423 {
		v418 = v418 + v426
		v419 = v419 + v426
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v437 = int32(2)
	goto L122
L150:
	;
	m.G0 = v7 + int32(80)
	return int32(0)
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
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
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v14 = v11 + int32(36)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 == v4 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L26
	} else {
		goto L82
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L26
	} else {
		goto L78
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L26
	} else {
		goto L74
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L26
	} else {
		goto L70
	}
L5:
	;
	if v44 != 0 {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v44 = v41
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
	v41 = v37
	goto L6
L8:
	;
	v33 = int32(0)
	if v14 == v33 {
		v41 = v33
		goto L6
	} else {
		goto L18
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	switch v19 - int32(429) {
	case 0:
		goto L11
	case 1:
		goto L10
	default:
		goto L8
	}
L10:
	;
	if v14 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	if v14 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = int32(1)
	goto L5
L13:
	;
	goto L14
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+168))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v36 = v26
	v37 = int32(1)
	goto L7
L15:
	;
	v44 = int32(2)
	goto L5
L16:
	;
	goto L17
L17:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+368))
	v36 = v31
	v37 = int32(2)
	goto L7
L18:
	;
	v36 = v33
	v37 = v4
	goto L7
L19:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v45 == int32(1) {
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
	v240 = m.ExcPending
	if v240 != 0 {
		goto L26
	} else {
		goto L67
	}
L22:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v144 == int32(1) {
		goto L2
	} else {
		goto L41
	}
L23:
	;
	v48 = int32(4489440)
	v49 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v51
	v54 = F_palloc(m, int32(44))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v140 = v139
	goto L22
L26:
	;
	return int32(0)
L27:
	;
	v58 = F_makeStringInfo(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v58
	if l2 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v49
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v111 = F_get_fn_expr_argtype(m, v109, int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L26
	} else {
		goto L34
	}
L30:
	;
	v62 = v11 + int32(80)
	v63 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v62))) = v63
	v66 = v11 - int32(-64)
	*(*int64)(unsafe.Add(mBase, uint32(v66))) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v11)+72)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = int64(51539607564)
	v76 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = int32(1321)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = int32(1322)
	v89 = F_hash_create(m, int32(391332), int32(32), v11+int32(40), int32(1224))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L26
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v98 = v54 + int32(20)
	v99 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v98))) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = v99
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v89
	v93 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+40)) = v93
	goto L29
L34:
	;
	if v111 == int32(0) {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_json_categorize_type(m, v111, int32(0), v54+int32(4), v54+int32(8))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L26
	} else {
		goto L36
	}
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v124 = F_get_fn_expr_argtype(m, v122, int32(2))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L26
	} else {
		goto L37
	}
L37:
	;
	if v124 == int32(0) {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	F_json_categorize_type(m, v124, int32(0), v54+int32(12), v54+int32(16))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L26
	} else {
		goto L39
	}
L39:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	F_appendStringInfoString(m, v135, int32(714869))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v140 = v54
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
	m.G0 = v11 + int32(96)
	return v140
L43:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	v184 = int32(0)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v140)+8))
	F_datum_to_json_internal(m, v185, v184, v180, v187, v188, int32(1))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L26
	} else {
		goto L54
	}
L44:
	;
	v172 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if v174 < int32(3) {
		v180 = v173
		v181 = v172
		goto L43
	} else {
		goto L52
	}
L45:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v149 != int32(1) {
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
	v155 = v140 + int32(24)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v140)+24))
	if v156 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v159 = int32(4489440)
	v160 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v140)+40))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v162
	F_initStringInfo(m, v155)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L26
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+28)) = int32(0)
	v180 = v155
	v181 = int32(1)
	goto L43
L51:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v160
	v180 = v155
	v181 = int32(1)
	goto L43
L52:
	;
	F_appendStringInfoString(m, v173, int32(730094))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L26
	} else {
		goto L53
	}
L53:
	;
	v180 = v173
	v181 = v172
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
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v195 = F_MemoryContextStrdup(m, v192, v193+v183)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L26
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	F_appendStringInfoString(m, v215, int32(729850))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L26
	} else {
		goto L62
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v195
	v198 = F_strlen(m, v195)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v198
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v140)+20))
	v208 = F_hash_search(m, v202, v11+int32(40), int32(1), v11+int32(95))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L26
	} else {
		goto L59
	}
L59:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+95)))
	if v210 == int32(1) {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v181 != 0 {
		goto L42
	} else {
		goto L61
	}
L61:
	;
	goto L57
L62:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v219 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v221 = v184
	goto L65
L64:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v221 = v220
	goto L65
L65:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
	F_datum_to_json_internal(m, v221, v219, v222, v223, v224, int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L26
	} else {
		goto L66
	}
L66:
	;
	goto L42
L67:
	;
	F_errmsg_internal(m, int32(61282), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L26
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(495612), int32(1016), int32(219814))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
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
	v256 = m.ExcPending
	if v256 != 0 {
		goto L26
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
	F_errmsg(m, int32(467458), v11)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L26
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(495612), int32(1043), int32(219814))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
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
	v273 = m.ExcPending
	if v273 != 0 {
		goto L26
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(2)
	F_errmsg(m, int32(467458), v11+int32(16))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L26
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(495612), int32(1053), int32(219814))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
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
	v292 = m.ExcPending
	if v292 != 0 {
		goto L26
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(21521), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L26
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(495612), int32(1076), int32(219814))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
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
	F_errcode(m, int32(786562))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L26
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v195
	F_errmsg(m, int32(202794), v11+int32(32))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L26
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(495612), int32(1126), int32(219814))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
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
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
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
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
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
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
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
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L89
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L85
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
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
	return v235
L10:
	;
	v28 = F_cstring_to_text(m, int32(4103))
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
	v235 = v28
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
	F_initStringInfo(m, v10+int32(32))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_appendStringInfoChar(m, v10+int32(32), int32(123))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if int32(0) < v60 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v64 = int32(0)
	goto L22
L20:
	;
	goto L21
L21:
	;
	F_appendStringInfoChar(m, v10+int32(32), int32(125))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L74
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v64))))
	if v73 == int32(1) {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	if v64 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_appendStringInfoString(m, v10+int32(32), int32(730094))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v82 = v64 << (uint(int32(2)) % 32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82+v83)))
	v86 = F_pg_detoast_datum_packed(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v121 = int32(1)
	if v88&v121 != 0 {
		goto L41
	} else {
		goto L42
	}
L30:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	if v88 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v91 = int32(4)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if v93&int32(254) == int32(2) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v106 = int32(1)
	if v88&v106 != 0 {
		v118 = int32(base.Ui32(v88)>>(uint(v106)%32)) - v106
		goto L29
	} else {
		goto L40
	}
L34:
	;
	v102 = v91
	goto L36
L35:
	;
	v102 = base.B2i32(v93 == int32(18)) << (uint(v91) % 32)
	goto L36
L36:
	;
	if v93 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v105 = v91
	goto L39
L38:
	;
	v105 = v102
	goto L39
L39:
	;
	v118 = v105
	goto L29
L40:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v118 = int32(base.Ui32(v112)>>(uint(int32(2))%32)) - int32(4)
	goto L29
L41:
	;
	v125 = v121
	goto L43
L42:
	;
	v125 = int32(4)
	goto L43
L43:
	;
	F_escape_json_with_len(m, v10+int32(32), v86+v125, v118)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v86 != v85 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_pfree(m, v86)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	F_appendStringInfoString(m, v10+int32(32), int32(729850))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+v64))))
	if v139 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v201 = v64 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v201 < v202 {
		v64 = v201
		goto L22
	} else {
		goto L73
	}
L51:
	;
	F_appendStringInfoString(m, v10+int32(32), int32(303035))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v147+v82)))
	v150 = F_pg_detoast_datum_packed(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L50
L55:
	;
	v185 = int32(1)
	if v152&v185 != 0 {
		goto L67
	} else {
		goto L68
	}
L56:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if v152 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v155 = int32(4)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+1)))
	if v157&int32(254) == int32(2) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v170 = int32(1)
	if v152&v170 != 0 {
		v182 = int32(base.Ui32(v152)>>(uint(v170)%32)) - v170
		goto L55
	} else {
		goto L66
	}
L60:
	;
	v166 = v155
	goto L62
L61:
	;
	v166 = base.B2i32(v157 == int32(18)) << (uint(v155) % 32)
	goto L62
L62:
	;
	if v157 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v169 = v155
	goto L65
L64:
	;
	v169 = v166
	goto L65
L65:
	;
	v182 = v169
	goto L55
L66:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v182 = int32(base.Ui32(v176)>>(uint(int32(2))%32)) - int32(4)
	goto L55
L67:
	;
	v189 = v185
	goto L69
L68:
	;
	v189 = int32(4)
	goto L69
L69:
	;
	F_escape_json_with_len(m, v10+int32(32), v150+v189, v182)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v150 == v149 {
		goto L50
	} else {
		goto L71
	}
L71:
	;
	F_pfree(m, v150)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
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
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	F_pfree(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	F_pfree(m, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	F_pfree(m, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
	v230 = F_cstring_to_text_with_len(m, v228, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	F_pfree(m, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v235 = v230
	goto L9
L81:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(118141), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(495612), int32(1509), int32(326572))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
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
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(147012), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(495612), int32(1520), int32(326572))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
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
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(21521), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(495612), int32(1531), int32(326572))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
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
		F_iterate_json_values(m, v13, int32(2), v7+int32(24))
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
						v45 = int32(4)
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
						if v47&int32(254) == int32(2) {
							v56 = v45
						} else {
							v56 = base.B2i32(v47 == int32(18)) << (uint(v45) % 32)
						}
						if v47 == int32(1) {
							v59 = v45
						} else {
							v59 = v56
						}
						v70 = v59
					} else {
						v60 = int32(1)
						if v41 != 0 {
							v70 = int32(base.Ui32(v39)>>(uint(v60)%32)) - v60
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
							v70 = int32(base.Ui32(v64)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v72 = *(*int32)(unsafe.Add(mBase, _consts[109]))
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
					v75 = F_makeJsonLexContextCstringLen(m, v11+int32(12), v42, v70, v73, int32(1))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v75
						v78 = F_makeStringInfo(m)
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v26)+9)) = uint8(v24)
							v81 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v26)+8)) = uint8(v81)
							*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v78
							*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = int32(1387)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = int32(1388)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = int32(1389)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = int32(1390)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = int32(1391)
							*(*int32)(unsafe.Add(mBase, uint32(v29))) = v26
							*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = int32(1392)
							*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = int32(1393)
							v101 = F_pg_parse_json(m, v11+int32(12), v29)
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int32(0)
							} else {
								if v101 != 0 {
									F_json_errsave_error(m, v101, v11+int32(12), int32(0))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
										v111 = F_cstring_to_text_with_len(m, v109, v110)
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											m.G0 = v11 + int32(80)
											return v111
										}
									}
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
									v111 = F_cstring_to_text_with_len(m, v109, v110)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(80)
										return v111
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
			v18 = int32(4)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v20&int32(254) == int32(2) {
				v29 = v18
			} else {
				v29 = base.B2i32(v20 == int32(18)) << (uint(v18) % 32)
			}
			if v20 == int32(1) {
				v32 = v18
			} else {
				v32 = v29
			}
			v43 = v32
		} else {
			v33 = int32(1)
			if v14 != 0 {
				v43 = int32(base.Ui32(v12)>>(uint(v33)%32)) - v33
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v45 = *(*int32)(unsafe.Add(mBase, _consts[109]))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
		v47 = F_makeJsonLexContextCstringLen(m, l0, v15, v43, v46, l2)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			return
		}
	}
}
func F_makeJsonValueExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_palloc0(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(44)
		return v6
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
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
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
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
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
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
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	if l2 == v6 {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v313 = F_format_type_be(m, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L15
	} else {
		goto L119
	}
L2:
	;
	if l2 != 0 {
		goto L115
	} else {
		goto L116
	}
L3:
	;
	v203 = int32(0)
	v204 = F_exprType(m, v199)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L15
	} else {
		goto L71
	}
L4:
	;
	v195 = F_makeConst(m, v192, int32(-1), int32(0), v189, v191, v190, v188)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L15
	} else {
		goto L70
	}
L5:
	;
	v188 = v62
	v189 = v61
	v190 = v62
	v191 = v187
	v192 = v60
	goto L4
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L15
	} else {
		goto L62
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L15
	} else {
		goto L57
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L15
	} else {
		goto L52
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L15
	} else {
		goto L47
	}
L10:
	;
	if v55 != 0 {
		goto L32
	} else {
		goto L33
	}
L11:
	;
	v54 = l3
	v55 = v6
	v57 = int32(-1)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v20 != int32(8) {
		v54 = v20
		v55 = v6
		v57 = v19
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v25 = F_transformExprRecurse(m, l0, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	v30 = F_ValidJsonBehaviorDefaultExpr(m, v25, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if v30 == int32(0) {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v34 = F_contain_var_clause(m, v25)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	if v34 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v36 = F_expression_returns_set(m, v25)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	if v36 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v38 = F_exprCollation(m, v25)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	if v38 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v42 = F_exprType(m, v25)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L15
	} else {
		goto L27
	}
L25:
	;
	v46 = v38
	goto L26
L26:
	;
	v47 = int32(8)
	if v23 == int32(0) {
		v54 = v47
		v55 = v25
		v57 = v19
		goto L10
	} else {
		goto L29
	}
L27:
	;
	v44 = F_get_typcollation(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v46 = v44
	goto L26
L29:
	;
	if v46 == int32(0) {
		v54 = v47
		v55 = v25
		v57 = v19
		goto L10
	} else {
		goto L30
	}
L30:
	;
	if v46 != v23 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v54 = v47
	v55 = v25
	v57 = v19
	goto L10
L32:
	;
	if v55 != 0 {
		v199 = v55
		goto L3
	} else {
		goto L46
	}
L33:
	;
	if v54 == int32(1) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v60 = int32(3802)
	v61 = int32(-1)
	v62 = int32(0)
	switch v54 {
	case 0, 2, 5:
		goto L37
	default:
		goto L36
	case 3:
		goto L35
	case 4:
		goto L38
	case 6:
		goto L40
	case 7:
		goto L39
	case 8:
		v188 = v62
		v189 = v61
		v190 = v62
		v191 = v6
		v192 = v60
		goto L4
	}
L35:
	;
	v94 = int32(1)
	v188 = v94
	v189 = v94
	v190 = v62
	v191 = v94
	v192 = int32(16)
	goto L4
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L15
	} else {
		goto L43
	}
L37:
	;
	v77 = int32(1)
	v188 = v77
	v189 = int32(4)
	v190 = v77
	v191 = v6
	v192 = int32(23)
	goto L4
L38:
	;
	v74 = int32(1)
	v188 = v74
	v189 = v74
	v190 = v62
	v191 = v6
	v192 = int32(16)
	goto L4
L39:
	;
	v72 = F_DirectFunctionCall1Coll(m, int32(486), int32(0), int32(4103))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L15
	} else {
		goto L42
	}
L40:
	;
	v67 = F_DirectFunctionCall1Coll(m, int32(486), int32(0), int32(507054))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L15
	} else {
		goto L41
	}
L41:
	;
	v187 = v67
	goto L5
L42:
	;
	v187 = v72
	goto L5
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v54
	F_errmsg_internal(m, int32(470151), v14)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(494448), int32(4954), int32(68370))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L15
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
	v98 = int32(0)
	v298 = v98
	v300 = v98
	goto L2
L47:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(519628), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	v111 = F_exprLocation(m, v25)
	mBase = m.M
	F_parser_errposition(m, l0, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L15
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(494448), int32(4767), int32(212600))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L15
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L15
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(171256), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L15
	} else {
		goto L54
	}
L54:
	;
	v130 = F_exprLocation(m, v25)
	mBase = m.M
	F_parser_errposition(m, l0, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L15
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(494448), int32(4772), int32(212600))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L15
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L15
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(106979), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L15
	} else {
		goto L59
	}
L59:
	;
	v149 = F_exprLocation(m, v25)
	mBase = m.M
	F_parser_errposition(m, l0, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L15
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(494448), int32(4777), int32(212600))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L15
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
	F_errcode(m, int32(17432708))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L15
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(358680), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L15
	} else {
		goto L64
	}
L64:
	;
	v168 = F_get_collation_name(m, v46)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	v170 = F_get_collation_name(m, v23)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L15
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v168
	F_errdetail(m, int32(683771), v12+int32(-16))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L15
	} else {
		goto L67
	}
L67:
	;
	v179 = F_exprLocation(m, v25)
	mBase = m.M
	F_parser_errposition(m, l0, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L15
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(494448), int32(4795), int32(212600))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L15
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
	*(*int32)(unsafe.Add(mBase, uint32(v195)+28)) = v57
	v199 = v195
	goto L3
L71:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v204 == v206 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v298 = v203
	v300 = v199
	goto L2
L73:
	;
	goto L74
L74:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	if v208 == int32(7) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v282 = int32(-1)
	v283 = int32(0)
	if v54 == int32(3) {
		goto L109
	} else {
		goto L110
	}
L76:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v231 = F_TypeCategory(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L15
	} else {
		goto L90
	}
L77:
	;
	v225 = int32(1)
	v226 = F_exprType(m, v199)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L15
	} else {
		goto L88
	}
L78:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+24)))
	if v211 != 0 {
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v212 = F_exprType(m, v199)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L15
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	if v212 == int32(3802) {
		goto L77
	} else {
		goto L83
	}
L83:
	;
	v216 = F_exprType(m, v199)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L15
	} else {
		goto L84
	}
L84:
	;
	if v216 != int32(16) {
		goto L76
	} else {
		goto L85
	}
L85:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v221 = F_getBaseType(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L15
	} else {
		goto L86
	}
L86:
	;
	if v221 == int32(23) {
		goto L76
	} else {
		goto L87
	}
L87:
	;
	goto L77
L88:
	;
	if v226 == int32(16) {
		goto L75
	} else {
		goto L89
	}
L89:
	;
	v298 = v225
	v300 = v199
	goto L2
L90:
	;
	v233 = F_exprType(m, v199)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L15
	} else {
		goto L91
	}
L91:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v237 = int32(1)
	if v231 == int32(86) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v242 = v237
	goto L94
L93:
	;
	v242 = int32(3)
	goto L94
L94:
	;
	if v231 == int32(83) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v245 = v237
	goto L97
L96:
	;
	v245 = v242
	goto L97
L97:
	;
	v247 = F_exprLocation(m, l2)
	mBase = m.M
	v248 = F_coerce_to_target_type(m, l0, v199, v233, v235, v236, v245, int32(1), v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L15
	} else {
		goto L98
	}
L98:
	;
	if v248 != 0 {
		v298 = v203
		v300 = v248
		goto L2
	} else {
		goto L99
	}
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L15
	} else {
		goto L100
	}
L100:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L15
	} else {
		goto L101
	}
L101:
	;
	v257 = F_exprType(m, v199)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L15
	} else {
		goto L102
	}
L102:
	;
	v259 = F_format_type_be(m, v257)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L15
	} else {
		goto L103
	}
L103:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v262 = F_format_type_be(m, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L15
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v259
	F_errmsg(m, int32(182829), v12+int32(-32))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L15
	} else {
		goto L105
	}
L105:
	;
	if v54 == int32(8) {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v273 = F_exprLocation(m, v199)
	mBase = m.M
	F_parser_errposition(m, l0, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L15
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(494448), int32(4882), int32(212600))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L15
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	v291 = int32(343846)
	goto L111
L110:
	;
	v291 = int32(360862)
	goto L111
L111:
	;
	v292 = F_DirectFunctionCall1Coll(m, int32(486), v283, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L15
	} else {
		goto L112
	}
L112:
	;
	v294 = int32(0)
	v296 = F_makeConst(m, int32(3802), v282, v283, v282, v292, v294, v294)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L15
	} else {
		goto L113
	}
L113:
	;
	v298 = v225
	v300 = v296
	goto L2
L114:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v306)+12)) = uint8(v298)
	m.G0 = v14 - int32(-64)
	return v306
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v300
	v306 = l2
	goto L114
L116:
	;
	goto L117
L117:
	;
	v304 = F_makeJsonBehavior(m, v54, v300, v57)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L15
	} else {
		goto L118
	}
L118:
	;
	v306 = v304
	goto L114
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v313
	F_errhint(m, int32(589840), v12+int32(-48))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L15
	} else {
		goto L120
	}
L120:
	;
	v321 = F_exprLocation(m, v199)
	mBase = m.M
	F_parser_errposition(m, l0, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L15
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(494448), int32(4875), int32(212600))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L15
	} else {
		goto L122
	}
L122:
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
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
			if v16 == int32(114) {
				v63 = v12
				m.G0 = v9 + int32(16)
				return v63
			} else {
				if v16 == int32(3802) {
					v63 = v12
					m.G0 = v9 + int32(16)
					return v63
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v29 = F_format_type_be(m, v28)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v29
								F_errmsg(m, int32(186442), v9)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(638764), int32(0))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
										F_parser_errposition(m, l0, v41)
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(494448), int32(4151), int32(334148))
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
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
		}
	} else {
		v50 = F_palloc0(m, int32(16))
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(43)
			v57 = F_makeJsonFormat(m, int32(1), int32(0), int32(-1))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v50)+8)) = int64(-4294967182)
				*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = v57
				v63 = v50
				m.G0 = v9 + int32(16)
				return v63
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
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v218 int32
	_ = v218
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
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L110
	}
L11:
	;
	m.G0 = v14 + int32(48)
	return v206
L12:
	;
	if l5 != 0 {
		goto L66
	} else {
		goto L67
	}
L13:
	;
	if l4 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L14:
	;
	if l3 != 0 {
		v111 = l3
		goto L12
	} else {
		goto L61
	}
L15:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+39)))
	if v103 != int32(83) {
		goto L14
	} else {
		goto L60
	}
L16:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v29-int32(700)) {
		goto L15
	} else {
		goto L59
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L54
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
	v111 = v39
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
	if v29 == int32(114) {
		goto L13
	} else {
		goto L52
	}
L30:
	;
	if base.Ui32(int32(25)) < base.Ui32(v29) {
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
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if int32(1)<<(uint(v29)%32)&int32(45154304) == int32(0) {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	v206 = v28
	goto L11
L35:
	;
	if base.Ui32(v29-int32(1082)) < base.Ui32(int32(2)) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	if v29 == int32(1184) {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v206 = v28
	goto L11
L39:
	;
	goto L40
L40:
	;
	if v29 == int32(1043) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v206 = v28
	goto L11
L42:
	;
	goto L43
L43:
	;
	if v29 != int32(1114) {
		goto L15
	} else {
		goto L44
	}
L44:
	;
	v206 = v28
	goto L11
L45:
	;
	v206 = v28
	goto L11
L46:
	;
	goto L47
L47:
	;
	if v29 == int32(1266) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v206 = v28
	goto L11
L49:
	;
	goto L50
L50:
	;
	if v29 != int32(1700) {
		goto L15
	} else {
		goto L51
	}
L51:
	;
	v206 = v28
	goto L11
L52:
	;
	if v29 == int32(3802) {
		goto L13
	} else {
		goto L53
	}
L53:
	;
	goto L14
L54:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(367243), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	F_parser_errposition(m, l0, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(494448), int32(3337), int32(207222))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v206 = v28
	goto L11
L60:
	;
	v206 = v28
	goto L11
L61:
	;
	goto L13
L62:
	;
	v206 = v28
	goto L11
L63:
	;
	goto L64
L64:
	;
	if v29 != l4 {
		v111 = int32(0)
		goto L12
	} else {
		goto L65
	}
L65:
	;
	v206 = v28
	goto L11
L66:
	;
	if v111 != int32(1) {
		goto L83
	} else {
		goto L84
	}
L67:
	;
	if l4 != 0 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	if v29 == int32(17) {
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+39)))
	if v114 == int32(83) {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v127 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v128 = int32(358245)
	goto L75
L74:
	;
	v128 = int32(358306)
	goto L75
L75:
	;
	F_errmsg(m, v128, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	if v133 < int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v136 = v31
	goto L79
L78:
	;
	v136 = v133
	goto L79
L79:
	;
	F_parser_errposition(m, l0, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(494448), int32(3403), int32(207222))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
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
	if v111 == int32(2) {
		goto L90
	} else {
		goto L91
	}
L83:
	;
	v168 = v29
	v169 = v28
	goto L82
L84:
	;
	goto L85
L85:
	;
	if v29 != int32(17) {
		v168 = v29
		v169 = v28
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v148 = F_getJsonEncodingConst(m, v38)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v148
	v154 = int32(25)
	v161 = F_list_make2_impl(m, v14+int32(28), v14+int32(24))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v163 = int32(0)
	v165 = F_makeFuncExpr(m, int32(1714), v154, v161, v163, v163)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+32)) = v31
	v168 = v154
	v169 = v165
	goto L82
L90:
	;
	v174 = int32(3802)
	goto L92
L91:
	;
	v174 = int32(114)
	goto L92
L92:
	;
	if l4 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v175 = l4
	goto L95
L94:
	;
	v175 = v174
	goto L95
L95:
	;
	v179 = F_coerce_to_target_type(m, l0, v169, v168, v175, int32(-1), int32(3), int32(1), v31)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	if v179 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	if l4 != 0 {
		goto L10
	} else {
		goto L100
	}
L98:
	;
	v200 = v179
	goto L99
L99:
	;
	if v169 == v200 {
		goto L106
	} else {
		goto L107
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v169
	if v111 == int32(2) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v189 = int32(3787)
	goto L103
L102:
	;
	v189 = int32(3176)
	goto L103
L103:
	;
	v193 = F_list_make1_impl(m, int32(1), v14+int32(12))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v195 = int32(0)
	v197 = F_makeFuncExpr(m, v189, v174, v193, v195, v195)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+32)) = v31
	v200 = v197
	goto L99
L106:
	;
	v206 = v28
	goto L11
L107:
	;
	goto L108
L108:
	;
	v202 = F_copyObjectImpl(m, l2)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+8)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = v28
	v206 = v202
	goto L11
L110:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v222 = F_format_type_be(m, v168)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v224 = F_format_type_be(m, l4)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v224
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v222
	F_errmsg(m, int32(182768), v14+int32(16))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_parser_errposition(m, l0, v31)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(494448), int32(3438), int32(207222))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
