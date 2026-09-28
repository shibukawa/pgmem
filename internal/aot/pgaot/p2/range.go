package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecGetRangeTableRelation(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v6 = F_bms_is_member(m, l1, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v12 = l1 - int32(1)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v10+v12<<(uint(int32(2))%32))))
			if v16 == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v20+l1<<(uint(int32(2))%32)-int32(4))))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGetRangeTableRelation[0]))
				if int32(0) <= v29 {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
					v34 = v32
				} else {
					v34 = int32(0)
				}
				v35 = F_table_open(m, v27, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v37+v12<<(uint(int32(2))%32)))) = v35
					v43 = v35
					return v43
				}
			} else {
				v43 = v16
				return v43
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_ExecGetRangeTableRelation_0), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_ExecGetRangeTableRelation_1), int32(858), int32(_a_F_ExecGetRangeTableRelation_2))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
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
func F_RangeVarCallbackForAlterRelation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	v9 = m.G0
	v11 = v9 - int32(224)
	m.G0 = v11
	v15 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L13
	} else {
		goto L14
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L13
	} else {
		goto L110
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L13
	} else {
		goto L105
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L13
	} else {
		goto L100
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L13
	} else {
		goto L95
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L13
	} else {
		goto L91
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L13
	} else {
		goto L87
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L13
	} else {
		goto L83
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L13
	} else {
		goto L79
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L13
	} else {
		goto L75
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L13
	} else {
		goto L71
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L68
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L13
	} else {
		goto L64
	}
L13:
	;
	return
L14:
	;
	if v15 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
	v19 = v17 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+119)))
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForAlterRelation[0]))
	v24 = F_object_ownercheck(m, int32(1259), l1, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	m.G0 = v11 + int32(224)
	return
L18:
	;
	if v24 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v29 = F_get_rel_relkind(m, l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarCallbackForAlterRelation[1])))
	if v47 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	switch v29 - int32(73) {
	case 0, 32:
		v40 = int32(20)
		goto L24
	default:
		goto L25
	case 10:
		goto L29
	case 29:
		goto L26
	case 36:
		goto L27
	case 45:
		goto L28
	}
L23:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_aclcheck_error(m, int32(2), v42, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L13
	} else {
		goto L30
	}
L24:
	;
	v42 = v40
	goto L23
L25:
	;
	v40 = int32(42)
	goto L24
L26:
	;
	v42 = int32(18)
	goto L23
L27:
	;
	v42 = int32(23)
	goto L23
L28:
	;
	v42 = int32(52)
	goto L23
L29:
	;
	v42 = int32(38)
	goto L23
L30:
	;
	goto L21
L31:
	;
	v51 = int32(1)
	if base.Ui32(l1) < base.Ui32(int32(_a_F_RangeVarCallbackForAlterRelation_0)) {
		v59 = v51
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	v60 = int32(4)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v61 - int32(215) {
	case 0:
		goto L41
	case 1:
		goto L11
	case 2:
		v82 = v60
		goto L39
	default:
		goto L40
	}
L34:
	;
	if v59 != 0 {
		goto L12
	} else {
		goto L38
	}
L35:
	;
	goto L34
L36:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
	if v54 == int32(99) {
		v59 = v51
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v57 = F_isTempToastNamespace(m, v54)
	mBase = m.M
	v59 = v57
	goto L35
L38:
	;
	goto L33
L39:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82+l3)))
	if base.B2i32(v85 == int32(38))&base.B2i32(v20 != int32(83)) != 0 {
		goto L10
	} else {
		goto L47
	}
L40:
	;
	if v61 != int32(146) {
		goto L11
	} else {
		goto L46
	}
L41:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForAlterRelation[0]))
	v69 = F_object_aclcheck(m, int32(2615), v65, v67, int64(512))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	if v69 == int32(0) {
		v82 = v60
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
	v75 = F_get_namespace_name(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	F_aclcheck_error(m, v69, int32(37), v75)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v82 = v60
	goto L39
L46:
	;
	v82 = int32(12)
	goto L39
L47:
	;
	if base.B2i32(v85 == int32(52))&base.B2i32(v20 != int32(118)) != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	if base.B2i32(v85 == int32(23))&base.B2i32(v20 != int32(109)) != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	if base.B2i32(v85 == int32(18))&base.B2i32(v20 != int32(102)) != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	if base.B2i32(v85 == int32(50))&base.B2i32(v20 != int32(99)) != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	v112 = v20 & int32(223)
	if base.B2i32(v112 == int32(73))|base.B2i32(v85 != int32(20)) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v120 != int32(215) {
		goto L5
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v124 = base.B2i32(v20 == int32(99))
	if v124&base.B2i32(v85 != int32(50)) != 0 {
		goto L4
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v128 == int32(217) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v112 == int32(73) {
		goto L3
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_ReleaseCatCache(m, v15)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L13
	} else {
		goto L63
	}
L60:
	;
	if v20 == int32(99) {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	if v20 == int32(116) {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	goto L17
L64:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v151
	F_errmsg(m, int32(_a_F_RangeVarCallbackForAlterRelation_1), v11+int32(208))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L13
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForAlterRelation_2), int32(_a_F_RangeVarCallbackForAlterRelation_3), int32(_a_F_RangeVarCallbackForAlterRelation_4))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
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
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v167
	F_errmsg_internal(m, int32(_a_F_RangeVarCallbackForAlterRelation_28), v11)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForAlterRelation_2), int32(_a_F_RangeVarCallbackForAlterRelation_29), int32(_a_F_RangeVarCallbackForAlterRelation_4))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v184
	F_errmsg(m, int32(_a_F_RangeVarCallbackForAlterRelation_5), v11+int32(16))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForAlterRelation_2), int32(_a_F_RangeVarCallbackForAlterRelation_6), int32(_a_F_RangeVarCallbackForAlterRelation_4))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L13
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v203
	F_errmsg(m, int32(_a_F_RangeVarCallbackForAlterRelation_7), v11+int32(32))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForAlterRelation_2), int32(_a_F_RangeVarCallbackForAlterRelation_8), int32(_a_F_RangeVarCallbackForAlterRelation_4))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L13
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v222
	F_errmsg(m, int32(_a_F_RangeVarCallbackForAlterRelation_9), v11+int32(48))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForAlterRelation_2), int32(_a_F_RangeVarCallbackForAlterRelation_10), int32(_a_F_RangeVarCallbackForAlterRelation_4))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L13
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v241
	F_errmsg(m, int32(_a_F_RangeVarCallbackForAlterRelation_11), v11-int32(-64))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForAlterRelation_2), int32(_a_F_RangeVarCallbackForAlterRelation_12), int32(_a_F_RangeVarCallbackForAlterRelation_4))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L13
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L13
	} else {
		goto L88
	}
L88:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v260
	F_errmsg(m, int32(_a_F_RangeVarCallbackForAlterRelation_13), v11+int32(80))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForAlterRelation_2), int32(_a_F_RangeVarCallbackForAlterRelation_14), int32(_a_F_RangeVarCallbackForAlterRelation_4))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L13
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L13
	} else {
		goto L92
	}
L92:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v279
	F_errmsg(m, int32(_a_F_RangeVarCallbackForAlterRelation_15), v11+int32(192))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L13
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForAlterRelation_2), int32(_a_F_RangeVarCallbackForAlterRelation_16), int32(_a_F_RangeVarCallbackForAlterRelation_4))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L13
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L13
	} else {
		goto L96
	}
L96:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v298
	F_errmsg(m, int32(_a_F_RangeVarCallbackForAlterRelation_17), v11+int32(112))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L13
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = int32(_a_F_RangeVarCallbackForAlterRelation_18)
	F_errhint(m, int32(_a_F_RangeVarCallbackForAlterRelation_19), v11+int32(96))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForAlterRelation_2), int32(_a_F_RangeVarCallbackForAlterRelation_20), int32(_a_F_RangeVarCallbackForAlterRelation_4))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L13
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L13
	} else {
		goto L101
	}
L101:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v324
	F_errmsg(m, int32(_a_F_RangeVarCallbackForAlterRelation_21), v11+int32(128))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L13
	} else {
		goto L102
	}
L102:
	;
	F_errhint(m, int32(_a_F_RangeVarCallbackForAlterRelation_22), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L13
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForAlterRelation_2), int32(_a_F_RangeVarCallbackForAlterRelation_23), int32(_a_F_RangeVarCallbackForAlterRelation_4))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L13
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L13
	} else {
		goto L106
	}
L106:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v347
	F_errmsg(m, int32(_a_F_RangeVarCallbackForAlterRelation_24), v11+int32(160))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L13
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = int32(_a_F_RangeVarCallbackForAlterRelation_18)
	F_errhint(m, int32(_a_F_RangeVarCallbackForAlterRelation_19), v11+int32(144))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L13
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForAlterRelation_2), int32(_a_F_RangeVarCallbackForAlterRelation_25), int32(_a_F_RangeVarCallbackForAlterRelation_4))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L13
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L13
	} else {
		goto L111
	}
L111:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v373
	F_errmsg(m, int32(_a_F_RangeVarCallbackForAlterRelation_26), v11+int32(176))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	F_errhint(m, int32(_a_F_RangeVarCallbackForAlterRelation_22), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L13
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForAlterRelation_2), int32(_a_F_RangeVarCallbackForAlterRelation_27), int32(_a_F_RangeVarCallbackForAlterRelation_4))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L13
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RangeVarCallbackForPolicy(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v17 = v15 + v16
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForPolicy[0]))
			v22 = F_object_ownercheck(m, int32(1259), l1, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if v22 == int32(0) {
					v27 = F_get_rel_relkind(m, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						switch v27 - int32(73) {
						case 0, 32:
							v38 = int32(20)
							v40 = v38
						default:
							v38 = int32(42)
							v40 = v38
						case 10:
							v40 = int32(38)
						case 29:
							v40 = int32(18)
						case 36:
							v40 = int32(23)
						case 45:
							v40 = int32(52)
						}
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						F_aclcheck_error(m, int32(2), v40, v41)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarCallbackForPolicy[1])))
							if v45 == int32(0) {
								v49 = int32(1)
								if base.Ui32(l1) < base.Ui32(int32(_a_F_RangeVarCallbackForPolicy_0)) {
									v57 = v49
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
									if v52 == int32(99) {
										v57 = v49
									} else {
										v55 = F_isTempToastNamespace(m, v52)
										mBase = m.M
										v57 = v55
									}
								}
								if v57 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										F_errcode(m, int32(16797828))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v75
											F_errmsg(m, int32(_a_F_RangeVarCallbackForPolicy_1), v9+int32(16))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_RangeVarCallbackForPolicy_2), int32(87), int32(_a_F_RangeVarCallbackForPolicy_3))
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									if v18&int32(253) != int32(112) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											F_errcode(m, int32(151027844))
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return
											} else {
												v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v9))) = v94
												F_errmsg(m, int32(_a_F_RangeVarCallbackForPolicy_4), v9)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_RangeVarCallbackForPolicy_2), int32(93), int32(_a_F_RangeVarCallbackForPolicy_3))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										F_ReleaseCatCache(m, v13)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							} else {
								if v18&int32(253) != int32(112) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = v94
											F_errmsg(m, int32(_a_F_RangeVarCallbackForPolicy_4), v9)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_RangeVarCallbackForPolicy_2), int32(93), int32(_a_F_RangeVarCallbackForPolicy_3))
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									F_ReleaseCatCache(m, v13)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						}
					}
				} else {
					v45 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarCallbackForPolicy[1])))
					if v45 == int32(0) {
						v49 = int32(1)
						if base.Ui32(l1) < base.Ui32(int32(_a_F_RangeVarCallbackForPolicy_0)) {
							v57 = v49
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
							if v52 == int32(99) {
								v57 = v49
							} else {
								v55 = F_isTempToastNamespace(m, v52)
								mBase = m.M
								v57 = v55
							}
						}
						if v57 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								F_errcode(m, int32(16797828))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v75
									F_errmsg(m, int32(_a_F_RangeVarCallbackForPolicy_1), v9+int32(16))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_RangeVarCallbackForPolicy_2), int32(87), int32(_a_F_RangeVarCallbackForPolicy_3))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							if v18&int32(253) != int32(112) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v94
										F_errmsg(m, int32(_a_F_RangeVarCallbackForPolicy_4), v9)
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_RangeVarCallbackForPolicy_2), int32(93), int32(_a_F_RangeVarCallbackForPolicy_3))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								F_ReleaseCatCache(m, v13)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					} else {
						if v18&int32(253) != int32(112) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return
								} else {
									v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v94
									F_errmsg(m, int32(_a_F_RangeVarCallbackForPolicy_4), v9)
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_RangeVarCallbackForPolicy_2), int32(93), int32(_a_F_RangeVarCallbackForPolicy_3))
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							F_ReleaseCatCache(m, v13)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			}
		} else {
			m.G0 = v9 + int32(32)
			return
		}
	}
}
func F__equalRangeTableSample(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v28 = v3
			return v28
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_equal(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v28 = v3
					return v28
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v20 = F_equal(m, v18, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						if v20 == int32(0) {
							v28 = v3
							return v28
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v26 = F_equal(m, v24, v25)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								v28 = v26
								return v28
							}
						}
					}
				}
			}
		}
	}
}
func F_get_range_key_properties(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10+l1<<(uint(int32(1))%32)))))
	if v14 != 0 {
		v17 = l1 << (uint(int32(2)) % 32)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v21+v17)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v24+v17)))
		v28 = F_makeVar(m, int32(1), v14, v20, v23, v26, int32(0))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l5))) = v28
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			if v53 != 0 {
				v58 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l6))) = v58
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				if v60 != 0 {
					v65 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l7))) = v65
					return
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					v63 = F_copyObjectImpl(m, v62)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						v65 = v63
						*(*int32)(unsafe.Add(mBase, uint32(l7))) = v65
						return
					}
				}
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v56 = F_copyObjectImpl(m, v55)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					v58 = v56
					*(*int32)(unsafe.Add(mBase, uint32(l6))) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					if v60 != 0 {
						v65 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l7))) = v65
						return
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
						v63 = F_copyObjectImpl(m, v62)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							v65 = v63
							*(*int32)(unsafe.Add(mBase, uint32(l7))) = v65
							return
						}
					}
				}
			}
		}
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
		if v31 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_get_range_key_properties_0), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_get_range_key_properties_1), int32(_a_F_get_range_key_properties_2), int32(_a_F_get_range_key_properties_3))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
			v35 = F_copyObjectImpl(m, v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l5))) = v35
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
				v40 = v38 + int32(4)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
				if base.Ui32(v40) < base.Ui32(v43+v44<<(uint(int32(2))%32)) {
					v49 = v40
				} else {
					v49 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v49
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if v53 != 0 {
					v58 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l6))) = v58
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					if v60 != 0 {
						v65 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l7))) = v65
						return
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
						v63 = F_copyObjectImpl(m, v62)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							v65 = v63
							*(*int32)(unsafe.Add(mBase, uint32(l7))) = v65
							return
						}
					}
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					v56 = F_copyObjectImpl(m, v55)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						v58 = v56
						*(*int32)(unsafe.Add(mBase, uint32(l6))) = v58
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						if v60 != 0 {
							v65 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l7))) = v65
							return
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
							v63 = F_copyObjectImpl(m, v62)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								v65 = v63
								*(*int32)(unsafe.Add(mBase, uint32(l7))) = v65
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_range_after_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v9 == v10 {
		F_range_deserialize(m, l0, l1, v7-int32(-64), v7+int32(32), v7+int32(15))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_range_deserialize(m, l0, l2, v7+int32(48), v7+int32(16), v7+int32(14))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(0)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
				if v31 != 0 {
					v104 = v30
					m.G0 = v7 + int32(80)
					return v104
				} else {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
					if v32&int32(1) != 0 {
						v104 = v30
						m.G0 = v7 + int32(80)
						return v104
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+24)))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+72)))
						if v36 == int32(1) {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+74)))
							if v35&int32(1) != 0 {
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+26)))
								if v42 == v39 {
									v98 = int32(0)
								} else {
									v46 = int32(1)
									if v39&v46 != 0 {
										v49 = int32(-1)
									} else {
										v49 = v46
									}
									v98 = v49
								}
							} else {
								v51 = int32(1)
								if v39&v51 != 0 {
									v54 = int32(-1)
								} else {
									v54 = v51
								}
								v98 = v54
							}
							v104 = base.B2i32(int32(0) < v98)
							m.G0 = v7 + int32(80)
							return v104
						} else {
							if v35&int32(1) != 0 {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+26)))
								if v59 != 0 {
									v60 = int32(1)
								} else {
									v60 = int32(-1)
								}
								v98 = v60
								v104 = base.B2i32(int32(0) < v98)
								m.G0 = v7 + int32(80)
								return v104
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v64 = *(*int64)(unsafe.Add(mBase, uint32(v7)+64))
								v65 = *(*int64)(unsafe.Add(mBase, uint32(v7)+16))
								v66 = F_FunctionCall2Coll(m, l0+int32(212), v63, v64, v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v68 = base.I32_wrap_i64(v66)
									if v68 != 0 {
										v98 = v68
									} else {
										v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+25)))
										v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+73)))
										if v70 == int32(0) {
											v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+74)))
											if v69&int32(1) == int32(0) {
												v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+26)))
												if v78 == v73 {
													v98 = int32(0)
												} else {
													v81 = int32(1)
													if v73&v81 != 0 {
														v85 = v81
													} else {
														v85 = int32(-1)
													}
													v98 = v85
												}
											} else {
												v86 = int32(1)
												if v73&v86 != 0 {
													v90 = v86
												} else {
													v90 = int32(-1)
												}
												v98 = v90
											}
										} else {
											if v69&int32(1) != 0 {
												v98 = int32(0)
											} else {
												v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+26)))
												if v96 != 0 {
													v97 = int32(-1)
												} else {
													v97 = int32(1)
												}
												v98 = v97
											}
										}
									}
									v104 = base.B2i32(int32(0) < v98)
									m.G0 = v7 + int32(80)
									return v104
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v112 = m.ExcPending
		if v112 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_range_after_internal_0), int32(0))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_after_internal_1), int32(719), int32(_a_F_range_after_internal_2))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
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
func F_range_bound_escape(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_initStringInfo(m, v8)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v17 = l0
	v18 = v14
	goto L5
L3:
	;
	v29 = l0
	goto L11
L4:
	;
	F_appendStringInfoChar(m, v8, int32(34))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	switch v18 {
	case 0:
		goto L7
	default:
		goto L8
	case 9, 10, 11, 12, 13, 32, 34, 40, 41, 44, 91, 92, 93:
		goto L4
	}
L6:
	;
	if v14 != 0 {
		v28 = int32(0)
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	v17 = v17 + int32(1)
	v18 = v20
	goto L5
L9:
	;
	goto L4
L10:
	;
	v28 = int32(1)
	goto L3
L11:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v35 = base.I32_extend8_s(v34)
	if base.B2i32(v34 == int32(34))|base.B2i32(v34 == int32(92)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	F_appendStringInfoChar(m, v8, v35)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L23
	}
L14:
	;
	if v34 != 0 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_appendStringInfoChar(m, v8, v35)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L22
	}
L17:
	;
	if v28 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_appendStringInfoChar(m, v8, int32(34))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	m.G0 = v8 + int32(16)
	return v46
L21:
	;
	goto L20
L22:
	;
	goto L13
L23:
	;
	v29 = v29 + int32(1)
	goto L11
}
func F_range_cmp(m *base.Module, l0 int32) int64 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int64
	_ = v79
	var v84 int64
	_ = v84
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v114 int64
	_ = v114
	var v119 int64
	_ = v119
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int64
	_ = v143
	var v148 int64
	_ = v148
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v179 int64
	_ = v179
	var v184 int64
	_ = v184
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v23 == v24 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L117
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L114
	}
L9:
	;
	F_range_deserialize(m, v38, v14, v11-int32(-64), v11+int32(32), v11+int32(15))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v28 == v23 {
		v38 = v27
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v31 = F_lookup_type_cache(m, v23, int32(2048))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+200))
	if v33 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v31
	v38 = v31
	goto L9
L16:
	;
	F_range_deserialize(m, v38, v19, v11+int32(48), v11+int32(16), v11+int32(14))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)))
	v59 = int64(1)
	if v55 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v62 = base.I64_extend_i32_u(v55&v56) - v59
	goto L20
L19:
	;
	v62 = v59
	goto L20
L20:
	;
	if v55|v56&int32(1) != 0 {
		v195 = v62
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v196 != v14 {
		goto L106
	} else {
		goto L107
	}
L22:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+56)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+72)))
	if v67 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)))
	if v130 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L24:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+74)))
	if v66&int32(1) != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	if v66&int32(1) != 0 {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)))
	if v70 == v73 {
		goto L23
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v70&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	if v70&int32(1) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v79 = int64(-1)
	goto L33
L32:
	;
	v79 = int64(1)
	goto L33
L33:
	;
	v195 = v79
	goto L21
L34:
	;
	v84 = int64(-1)
	goto L36
L35:
	;
	v84 = int64(1)
	goto L36
L36:
	;
	v195 = v84
	goto L21
L37:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)))
	if v89 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v38)+208))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v11)+64))
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
	v96 = F_FunctionCall2Coll(m, v38+int32(212), v93, v94, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v90 = int64(1)
	goto L42
L41:
	;
	v90 = int64(-1)
	goto L42
L42:
	;
	v195 = v90
	goto L21
L43:
	;
	if base.I32_wrap_i64(v96) != 0 {
		v195 = v96
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+57)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+73)))
	if v100 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+74)))
	if v99&int32(1) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	if v99&int32(1) != 0 {
		goto L23
	} else {
		goto L58
	}
L48:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)))
	if v103 == v108 {
		goto L23
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v103&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	if v103&int32(1) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v114 = int64(1)
	goto L54
L53:
	;
	v114 = int64(-1)
	goto L54
L54:
	;
	v195 = v114
	goto L21
L55:
	;
	v119 = int64(1)
	goto L57
L56:
	;
	v119 = int64(-1)
	goto L57
L57:
	;
	v195 = v119
	goto L21
L58:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+58)))
	if v124 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v125 = int64(-1)
	goto L61
L60:
	;
	v125 = int64(1)
	goto L61
L61:
	;
	v195 = v125
	goto L21
L62:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)))
	if v129&int32(1) != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	if v129&int32(1) != 0 {
		goto L77
	} else {
		goto L78
	}
L65:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
	if v136 == v133 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if v133&int32(1) != 0 {
		goto L74
	} else {
		goto L75
	}
L68:
	;
	v195 = int64(0)
	goto L21
L69:
	;
	goto L70
L70:
	;
	if v133&int32(1) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v143 = int64(-1)
	goto L73
L72:
	;
	v143 = int64(1)
	goto L73
L73:
	;
	v195 = v143
	goto L21
L74:
	;
	v148 = int64(-1)
	goto L76
L75:
	;
	v148 = int64(1)
	goto L76
L76:
	;
	v195 = v148
	goto L21
L77:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
	if v153 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v38)+208))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
	v160 = F_FunctionCall2Coll(m, v38+int32(212), v157, v158, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L83
	}
L80:
	;
	v154 = int64(1)
	goto L82
L81:
	;
	v154 = int64(-1)
	goto L82
L82:
	;
	v195 = v154
	goto L21
L83:
	;
	if base.I32_wrap_i64(v160) != 0 {
		v195 = v160
		goto L21
	} else {
		goto L84
	}
L84:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)))
	if v164 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)))
	if v163&int32(1) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	if v163&int32(1) != 0 {
		goto L100
	} else {
		goto L101
	}
L88:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
	if v172 == v167 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if v167&int32(1) != 0 {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	v195 = int64(0)
	goto L21
L92:
	;
	goto L93
L93:
	;
	if v167&int32(1) != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v179 = int64(1)
	goto L96
L95:
	;
	v179 = int64(-1)
	goto L96
L96:
	;
	v195 = v179
	goto L21
L97:
	;
	v184 = int64(1)
	goto L99
L98:
	;
	v184 = int64(-1)
	goto L99
L99:
	;
	v195 = v184
	goto L21
L100:
	;
	v195 = int64(0)
	goto L21
L101:
	;
	goto L102
L102:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)))
	if v190 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v191 = int64(-1)
	goto L105
L104:
	;
	v191 = int64(1)
	goto L105
L105:
	;
	v195 = v191
	goto L21
L106:
	;
	F_pfree(m, v14)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v200 != v19 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	F_pfree(m, v19)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	m.G0 = v11 + int32(80)
	return base.I64_extend32_s(v195)
L113:
	;
	goto L112
L114:
	;
	F_errmsg_internal(m, int32(_a_F_range_cmp_3), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_range_cmp_1), int32(1438), int32(_a_F_range_cmp_4))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v23
	F_errmsg_internal(m, int32(_a_F_range_cmp_0), v11)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_range_cmp_1), int32(1946), int32(_a_F_range_cmp_2))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_range_eq(m *base.Module, l0 int32) int64 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = F_range_eq_internal(m, v32, v12, v17)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v33)
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_eq_0), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_range_eq_1), int32(1946), int32(_a_F_range_eq_2))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = F_range_eq_internal(m, v32, v12, v17)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int64(0)
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v33)
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_eq_0), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_range_eq_1), int32(1946), int32(_a_F_range_eq_2))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = F_range_eq_internal(m, v32, v12, v17)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v33)
						}
					}
				}
			}
		}
	}
}
func F_range_ge(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_range_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return int64(base.Ui64(v2^int64(-1)) >> (uint(int64(63)) % 64))
	}
}
func F_range_get_flags(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0+int32(base.Ui32(v2)>>(uint(int32(2))%32))-int32(1)))))
	return v8
}
func F_range_gist_consistent_int_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	switch l1 - int32(1) {
	case 0:
		v12 = int32(0)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v19 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v13)>>(uint(int32(2))%32))-int32(1)))))
		if v19&int32(1) != 0 {
			v182 = v12
			m.G0 = v8 + int32(16)
			return v182
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3+int32(base.Ui32(v22)>>(uint(int32(2))%32))-int32(1)))))
			if v28&int32(1) != 0 {
				v182 = v12
				m.G0 = v8 + int32(16)
				return v182
			} else {
				v31 = F_range_overright_internal(m, l0, l2, l3)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v182 = v31 ^ int32(1)
					m.G0 = v8 + int32(16)
					return v182
				}
			}
		}
	case 1:
		v37 = int32(0)
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v44 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v38)>>(uint(int32(2))%32))-int32(1)))))
		if v44&int32(1) != 0 {
			v182 = v37
			m.G0 = v8 + int32(16)
			return v182
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3+int32(base.Ui32(v47)>>(uint(int32(2))%32))-int32(1)))))
			if v53&int32(1) != 0 {
				v182 = v37
				m.G0 = v8 + int32(16)
				return v182
			} else {
				v56 = F_range_after_internal(m, l0, l2, l3)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v182 = v56 ^ int32(1)
					m.G0 = v8 + int32(16)
					return v182
				}
			}
		}
	case 2:
		v180 = F_range_overlaps_internal(m, l0, l2, l3)
		mBase = m.M
		v181 = m.ExcPending
		if v181 != 0 {
			return int32(0)
		} else {
			v182 = v180
			m.G0 = v8 + int32(16)
			return v182
		}
	case 3:
		v60 = int32(0)
		v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v61)>>(uint(int32(2))%32))-int32(1)))))
		if v67&int32(1) != 0 {
			v182 = v60
			m.G0 = v8 + int32(16)
			return v182
		} else {
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3+int32(base.Ui32(v70)>>(uint(int32(2))%32))-int32(1)))))
			if v76&int32(1) != 0 {
				v182 = v60
				m.G0 = v8 + int32(16)
				return v182
			} else {
				v79 = F_range_before_internal(m, l0, l2, l3)
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					v182 = v79 ^ int32(1)
					m.G0 = v8 + int32(16)
					return v182
				}
			}
		}
	case 4:
		v83 = int32(0)
		v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v90 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v84)>>(uint(int32(2))%32))-int32(1)))))
		if v90&int32(1) != 0 {
			v182 = v83
			m.G0 = v8 + int32(16)
			return v182
		} else {
			v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v99 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3+int32(base.Ui32(v93)>>(uint(int32(2))%32))-int32(1)))))
			if v99&int32(1) != 0 {
				v182 = v83
				m.G0 = v8 + int32(16)
				return v182
			} else {
				v102 = F_range_overleft_internal(m, l0, l2, l3)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					v182 = v102 ^ int32(1)
					m.G0 = v8 + int32(16)
					return v182
				}
			}
		}
	case 5:
		v106 = int32(0)
		v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v113 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v107)>>(uint(int32(2))%32))-int32(1)))))
		if v113&int32(1) != 0 {
			v182 = v106
			m.G0 = v8 + int32(16)
			return v182
		} else {
			v116 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v122 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3+int32(base.Ui32(v116)>>(uint(int32(2))%32))-int32(1)))))
			if v122&int32(1) != 0 {
				v182 = v106
				m.G0 = v8 + int32(16)
				return v182
			} else {
				v125 = F_range_adjacent_internal(m, l0, l2, l3)
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					if v125 == int32(0) {
						v180 = F_range_overlaps_internal(m, l0, l2, l3)
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
							return int32(0)
						} else {
							v182 = v180
							m.G0 = v8 + int32(16)
							return v182
						}
					} else {
						v182 = int32(1)
						m.G0 = v8 + int32(16)
						return v182
					}
				}
			}
		}
	case 6:
		v130 = F_range_contains_internal(m, l0, l2, l3)
		mBase = m.M
		v131 = m.ExcPending
		if v131 != 0 {
			return int32(0)
		} else {
			v182 = v130
			m.G0 = v8 + int32(16)
			return v182
		}
	case 7:
		v132 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v132)>>(uint(int32(2))%32))-int32(1)))))
		if v138&int32(-127) == int32(0) {
			v180 = F_range_overlaps_internal(m, l0, l2, l3)
			mBase = m.M
			v181 = m.ExcPending
			if v181 != 0 {
				return int32(0)
			} else {
				v182 = v180
				m.G0 = v8 + int32(16)
				return v182
			}
		} else {
			v182 = int32(1)
			m.G0 = v8 + int32(16)
			return v182
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v169 = m.ExcPending
		if v169 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
			F_errmsg_internal(m, int32(_a_F_range_gist_consistent_int_range_0), v8)
			mBase = m.M
			v173 = m.ExcPending
			if v173 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_gist_consistent_int_range_1), int32(968), int32(_a_F_range_gist_consistent_int_range_2))
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 17:
		v144 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v150 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3+int32(base.Ui32(v144)>>(uint(int32(2))%32))-int32(1)))))
		if v150&int32(1) != 0 {
			v153 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v159 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v153)>>(uint(int32(2))%32))-int32(1)))))
			v182 = base.B2i32(v159&int32(-127) != int32(0))
			m.G0 = v8 + int32(16)
			return v182
		} else {
			v164 = F_range_contains_internal(m, l0, l2, l3)
			mBase = m.M
			v165 = m.ExcPending
			if v165 != 0 {
				return int32(0)
			} else {
				v182 = v164
				m.G0 = v8 + int32(16)
				return v182
			}
		}
	}
}
func F_range_gist_consistent_leaf_multirange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	switch l1 - int32(1) {
	case 0:
		v100 = F_range_before_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v101 = m.ExcPending
		if v101 != 0 {
			return int32(0)
		} else {
			v102 = v100
			m.G0 = v11 + int32(96)
			return v102
		}
	case 1:
		v15 = F_range_overleft_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v102 = v15
			m.G0 = v11 + int32(96)
			return v102
		}
	case 2:
		v19 = F_range_overlaps_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v102 = v19
			m.G0 = v11 + int32(96)
			return v102
		}
	case 3:
		v21 = F_range_overright_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v102 = v21
			m.G0 = v11 + int32(96)
			return v102
		}
	case 4:
		v23 = F_range_after_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v102 = v23
			m.G0 = v11 + int32(96)
			return v102
		}
	case 5:
		v25 = F_range_adjacent_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v102 = v25
			m.G0 = v11 + int32(96)
			return v102
		}
	case 6:
		v27 = F_range_contains_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v102 = v27
			m.G0 = v11 + int32(96)
			return v102
		}
	case 7:
		v29 = F_multirange_contains_range_internal(m, l0, l3, l2)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v102 = v29
			m.G0 = v11 + int32(96)
			return v102
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v90 = m.ExcPending
		if v90 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
			F_errmsg_internal(m, int32(_a_F_range_gist_consistent_leaf_multirange_0), v11)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_gist_consistent_leaf_multirange_1), int32(1119), int32(_a_F_range_gist_consistent_leaf_multirange_2))
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 17:
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v31)>>(uint(int32(2))%32))-int32(1)))))
		if v37&int32(1) == int32(0) {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			if v42 != 0 {
				v59 = v11 + int32(80)
				v61 = v11 - int32(-64)
				F_range_deserialize(m, l0, l2, v59, v61, v11+int32(15))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					v66 = int32(0)
					v69 = v11 + int32(48)
					v71 = v11 + int32(16)
					F_multirange_get_bounds(m, l0, l3, v66, v69, v71)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
						v78 = v11 + int32(32)
						F_multirange_get_bounds(m, l0, l3, v74-int32(1), v71, v78)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v81 = F_range_cmp_bounds(m, l0, v59, v69)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								if v81 != 0 {
									v102 = v66
									m.G0 = v11 + int32(96)
									return v102
								} else {
									v83 = F_range_cmp_bounds(m, l0, v61, v78)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										v102 = base.B2i32(v83 == int32(0))
										m.G0 = v11 + int32(96)
										return v102
									}
								}
							}
						}
					}
				}
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v44)>>(uint(int32(2))%32))-int32(1)))))
				if v50&int32(1) == int32(0) {
					v102 = int32(0)
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					v102 = base.B2i32(v55 == int32(0))
				}
				m.G0 = v11 + int32(96)
				return v102
			}
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v44)>>(uint(int32(2))%32))-int32(1)))))
			if v50&int32(1) == int32(0) {
				v102 = int32(0)
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				v102 = base.B2i32(v55 == int32(0))
			}
			m.G0 = v11 + int32(96)
			return v102
		}
	}
}
func F_range_gist_single_sorting_split(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
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
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
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
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v24 = (v20 - int32(1)) & int32(_a_F_range_gist_single_sorting_split_0)
	v25 = F_palloc_mul(m, int32(24), v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v20&int32(_a_F_range_gist_single_sorting_split_0) != int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+32)) = v182
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v181
	m.G0 = v17 + int32(32)
	return
L4:
	;
	v38 = int32(1)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_qsort_arg(m, v25, v24, int32(24), int32(1684), l0)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L39
	}
L7:
	;
	v49 = v38 * int32(24)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(8)+v49)))
	v52 = F_pg_detoast_datum(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	F_qsort_arg(m, v25, v24, int32(24), int32(1684), l0)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L17
	}
L9:
	;
	v54 = v49 + v25
	*(*int32)(unsafe.Add(mBase, uint32(v54-int32(24)))) = v38
	v59 = v54 - int32(16)
	if l3 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v75 = (v38 + int32(1)) & int32(_a_F_range_gist_single_sorting_split_0)
	if base.Ui32(v75) <= base.Ui32(v24) {
		v38 = v75
		goto L7
	} else {
		goto L16
	}
L11:
	;
	F_range_deserialize(m, l0, v52, v17+int32(16), v59, v17+int32(15))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_range_deserialize(m, l0, v52, v59, v17+int32(16), v17+int32(15))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L10
L15:
	;
	goto L10
L16:
	;
	goto L8
L17:
	;
	v81 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v81
	v85 = int32(1)
	if base.Ui32(v24) <= base.Ui32(v85) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v88 = v85
	goto L20
L19:
	;
	v88 = v24
	goto L20
L20:
	;
	v98 = int32(0)
	v105 = v5
	v106 = v5
	goto L21
L21:
	;
	v108 = int32(24)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v25+v98*v108)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(8)+v111*v108)))
	v116 = F_pg_detoast_datum(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v181 = base.I64_extend_i32_u(v152)
	v182 = base.I64_extend_i32_u(v153)
	goto L3
L23:
	;
	if base.Ui32(v98) < base.Ui32(int32(base.Ui32(v24)>>(uint(int32(1))%32))) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v155 = v98 + int32(1)
	if v155 != v88 {
		v98 = v155
		v105 = v152
		v106 = v153
		goto L21
	} else {
		goto L38
	}
L25:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v119 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	goto L27
L27:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v135 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v127 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v125 + v127
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v130+v125<<(uint(v127)%32)))) = uint16(v111)
	v152 = v126
	v153 = v106
	goto L24
L29:
	;
	v125 = v119
	v126 = v116
	goto L28
L30:
	;
	goto L31
L31:
	;
	v122 = F_range_super_union(m, l0, v105, v116)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v125 = v124
	v126 = v122
	goto L28
L33:
	;
	v143 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v141 + v143
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v146+v141<<(uint(v143)%32)))) = uint16(v111)
	v152 = v105
	v153 = v142
	goto L24
L34:
	;
	v141 = v135
	v142 = v116
	goto L33
L35:
	;
	goto L36
L36:
	;
	v138 = F_range_super_union(m, l0, v106, v116)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v141 = v140
	v142 = v138
	goto L33
L38:
	;
	goto L22
L39:
	;
	v163 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v163
	v181 = int64(0)
	v182 = int64(0)
	goto L3
}
func F_range_merge_from_multirange(m *base.Module, l0 int32) int64 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			if v18 == v15 {
				v28 = v17
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				switch v29 {
				case 0:
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
					v31 = F_make_empty_range(m, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						v62 = v31
						m.G0 = v8 + int32(80)
						return base.I64_extend_i32_u(v62)
					}
				case 1:
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
					v35 = F_multirange_get_range(m, v33, v11, int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						v62 = v35
						m.G0 = v8 + int32(80)
						return base.I64_extend_i32_u(v62)
					}
				default:
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
					v40 = v8 - int32(-64)
					F_multirange_get_bounds(m, v37, v11, int32(0), v40, v8+int32(48))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int64(0)
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
						v52 = v8 + int32(16)
						F_multirange_get_bounds(m, v45, v11, v46-int32(1), v8+int32(32), v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
							v56 = int32(0)
							v58 = F_make_range(m, v55, v40, v52, v56, v56)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int64(0)
							} else {
								v62 = v58
								m.G0 = v8 + int32(80)
								return base.I64_extend_i32_u(v62)
							}
						}
					}
				}
			} else {
				v21 = F_lookup_type_cache(m, v15, int32(_a_F_range_merge_from_multirange_0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+296))
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
							F_errmsg_internal(m, int32(_a_F_range_merge_from_multirange_1), v8)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_range_merge_from_multirange_2), int32(561), int32(_a_F_range_merge_from_multirange_3))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v21
						v28 = v21
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
						switch v29 {
						case 0:
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
							v31 = F_make_empty_range(m, v30)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int64(0)
							} else {
								v62 = v31
								m.G0 = v8 + int32(80)
								return base.I64_extend_i32_u(v62)
							}
						case 1:
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
							v35 = F_multirange_get_range(m, v33, v11, int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int64(0)
							} else {
								v62 = v35
								m.G0 = v8 + int32(80)
								return base.I64_extend_i32_u(v62)
							}
						default:
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
							v40 = v8 - int32(-64)
							F_multirange_get_bounds(m, v37, v11, int32(0), v40, v8+int32(48))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
								v52 = v8 + int32(16)
								F_multirange_get_bounds(m, v45, v11, v46-int32(1), v8+int32(32), v52)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
									v56 = int32(0)
									v58 = F_make_range(m, v55, v40, v52, v56, v56)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int64(0)
									} else {
										v62 = v58
										m.G0 = v8 + int32(80)
										return base.I64_extend_i32_u(v62)
									}
								}
							}
						}
					}
				}
			}
		} else {
			v21 = F_lookup_type_cache(m, v15, int32(_a_F_range_merge_from_multirange_0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+296))
				if v23 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
						F_errmsg_internal(m, int32(_a_F_range_merge_from_multirange_1), v8)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_range_merge_from_multirange_2), int32(561), int32(_a_F_range_merge_from_multirange_3))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v21
					v28 = v21
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					switch v29 {
					case 0:
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
						v31 = F_make_empty_range(m, v30)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							v62 = v31
							m.G0 = v8 + int32(80)
							return base.I64_extend_i32_u(v62)
						}
					case 1:
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
						v35 = F_multirange_get_range(m, v33, v11, int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							v62 = v35
							m.G0 = v8 + int32(80)
							return base.I64_extend_i32_u(v62)
						}
					default:
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
						v40 = v8 - int32(-64)
						F_multirange_get_bounds(m, v37, v11, int32(0), v40, v8+int32(48))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							v52 = v8 + int32(16)
							F_multirange_get_bounds(m, v45, v11, v46-int32(1), v8+int32(32), v52)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
								v56 = int32(0)
								v58 = F_make_range(m, v55, v40, v52, v56, v56)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int64(0)
								} else {
									v62 = v58
									m.G0 = v8 + int32(80)
									return base.I64_extend_i32_u(v62)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_range_minus(m *base.Module, l0 int32) int64 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			if v20 == v21 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				if v24 != 0 {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
					if v25 == v20 {
						v35 = v24
						v36 = F_range_minus_internal(m, v35, v13, v18)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int64(0)
						} else {
							if v36 != 0 {
								v42 = base.I64_extend_i32_u(v36)
							} else {
								v39 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v39)
								v42 = int64(0)
							}
							m.G0 = v10 + int32(16)
							return v42
						}
					} else {
						v28 = F_lookup_type_cache(m, v20, int32(2048))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int64(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+200))
							if v30 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
									F_errmsg_internal(m, int32(_a_F_range_minus_0), v10)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_range_minus_1), int32(1946), int32(_a_F_range_minus_2))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v28
								v35 = v28
								v36 = F_range_minus_internal(m, v35, v13, v18)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int64(0)
								} else {
									if v36 != 0 {
										v42 = base.I64_extend_i32_u(v36)
									} else {
										v39 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v39)
										v42 = int64(0)
									}
									m.G0 = v10 + int32(16)
									return v42
								}
							}
						}
					}
				} else {
					v28 = F_lookup_type_cache(m, v20, int32(2048))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+200))
						if v30 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
								F_errmsg_internal(m, int32(_a_F_range_minus_0), v10)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_range_minus_1), int32(1946), int32(_a_F_range_minus_2))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v28
							v35 = v28
							v36 = F_range_minus_internal(m, v35, v13, v18)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int64(0)
							} else {
								if v36 != 0 {
									v42 = base.I64_extend_i32_u(v36)
								} else {
									v39 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v39)
									v42 = int64(0)
								}
								m.G0 = v10 + int32(16)
								return v42
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_range_minus_3), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_range_minus_1), int32(987), int32(_a_F_range_minus_4))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int64(0)
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
func F_range_out(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		F_check_stack_depth(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v19 = F_get_range_io_data(m, l0, v17, int32(1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				F_range_deserialize(m, v21, v11, v6+int32(-32), v6+int32(-48), v6+int32(-49))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(base.Ui32(v30)>>(uint(int32(2))%32))-int32(1)))))
					if v36&int32(41) == int32(0) {
						v43 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
						v44 = F_OutputFunctionCall(m, v19+int32(4), v43)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							v46 = v44
							if v36&int32(81) == int32(0) {
								v53 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
								v54 = F_OutputFunctionCall(m, v19+int32(4), v53)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									v56 = v54
									if v36&int32(1) != 0 {
										v60 = F_pstrdup(m, int32(_a_F_range_out_0))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int64(0)
										} else {
											v105 = v60
											m.G0 = v8 - int32(-64)
											return base.I64_extend_i32_u(v105)
										}
									} else {
										v63 = v6 + int32(-16)
										F_initStringInfo(m, v63)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int64(0)
										} else {
											if v36&int32(2) != 0 {
												v70 = int32(91)
											} else {
												v70 = int32(40)
											}
											F_appendStringInfoChar(m, v63, v70)
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int64(0)
											} else {
												if v36&int32(40) == int32(0) {
													v77 = F_range_bound_escape(m, v46)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int64(0)
													} else {
														F_appendStringInfoString(m, v63, v77)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int64(0)
														} else {
															v82 = v6 + int32(-16)
															F_appendStringInfoChar(m, v82, int32(44))
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return int64(0)
															} else {
																if v36&int32(80) == int32(0) {
																	v90 = F_range_bound_escape(m, v56)
																	mBase = m.M
																	v91 = m.ExcPending
																	if v91 != 0 {
																		return int64(0)
																	} else {
																		F_appendStringInfoString(m, v82, v90)
																		mBase = m.M
																		v93 = m.ExcPending
																		if v93 != 0 {
																			return int64(0)
																		} else {
																			if v36&int32(4) != 0 {
																				v100 = int32(93)
																			} else {
																				v100 = int32(41)
																			}
																			F_appendStringInfoChar(m, v6+int32(-16), v100)
																			mBase = m.M
																			v102 = m.ExcPending
																			if v102 != 0 {
																				return int64(0)
																			} else {
																				v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																				v105 = v103
																				m.G0 = v8 - int32(-64)
																				return base.I64_extend_i32_u(v105)
																			}
																		}
																	}
																} else {
																	if v36&int32(4) != 0 {
																		v100 = int32(93)
																	} else {
																		v100 = int32(41)
																	}
																	F_appendStringInfoChar(m, v6+int32(-16), v100)
																	mBase = m.M
																	v102 = m.ExcPending
																	if v102 != 0 {
																		return int64(0)
																	} else {
																		v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																		v105 = v103
																		m.G0 = v8 - int32(-64)
																		return base.I64_extend_i32_u(v105)
																	}
																}
															}
														}
													}
												} else {
													v82 = v6 + int32(-16)
													F_appendStringInfoChar(m, v82, int32(44))
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int64(0)
													} else {
														if v36&int32(80) == int32(0) {
															v90 = F_range_bound_escape(m, v56)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int64(0)
															} else {
																F_appendStringInfoString(m, v82, v90)
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return int64(0)
																} else {
																	if v36&int32(4) != 0 {
																		v100 = int32(93)
																	} else {
																		v100 = int32(41)
																	}
																	F_appendStringInfoChar(m, v6+int32(-16), v100)
																	mBase = m.M
																	v102 = m.ExcPending
																	if v102 != 0 {
																		return int64(0)
																	} else {
																		v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																		v105 = v103
																		m.G0 = v8 - int32(-64)
																		return base.I64_extend_i32_u(v105)
																	}
																}
															}
														} else {
															if v36&int32(4) != 0 {
																v100 = int32(93)
															} else {
																v100 = int32(41)
															}
															F_appendStringInfoChar(m, v6+int32(-16), v100)
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return int64(0)
															} else {
																v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																v105 = v103
																m.G0 = v8 - int32(-64)
																return base.I64_extend_i32_u(v105)
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v56 = v2
								if v36&int32(1) != 0 {
									v60 = F_pstrdup(m, int32(_a_F_range_out_0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int64(0)
									} else {
										v105 = v60
										m.G0 = v8 - int32(-64)
										return base.I64_extend_i32_u(v105)
									}
								} else {
									v63 = v6 + int32(-16)
									F_initStringInfo(m, v63)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int64(0)
									} else {
										if v36&int32(2) != 0 {
											v70 = int32(91)
										} else {
											v70 = int32(40)
										}
										F_appendStringInfoChar(m, v63, v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int64(0)
										} else {
											if v36&int32(40) == int32(0) {
												v77 = F_range_bound_escape(m, v46)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int64(0)
												} else {
													F_appendStringInfoString(m, v63, v77)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int64(0)
													} else {
														v82 = v6 + int32(-16)
														F_appendStringInfoChar(m, v82, int32(44))
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int64(0)
														} else {
															if v36&int32(80) == int32(0) {
																v90 = F_range_bound_escape(m, v56)
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return int64(0)
																} else {
																	F_appendStringInfoString(m, v82, v90)
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return int64(0)
																	} else {
																		if v36&int32(4) != 0 {
																			v100 = int32(93)
																		} else {
																			v100 = int32(41)
																		}
																		F_appendStringInfoChar(m, v6+int32(-16), v100)
																		mBase = m.M
																		v102 = m.ExcPending
																		if v102 != 0 {
																			return int64(0)
																		} else {
																			v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																			v105 = v103
																			m.G0 = v8 - int32(-64)
																			return base.I64_extend_i32_u(v105)
																		}
																	}
																}
															} else {
																if v36&int32(4) != 0 {
																	v100 = int32(93)
																} else {
																	v100 = int32(41)
																}
																F_appendStringInfoChar(m, v6+int32(-16), v100)
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
																	return int64(0)
																} else {
																	v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																	v105 = v103
																	m.G0 = v8 - int32(-64)
																	return base.I64_extend_i32_u(v105)
																}
															}
														}
													}
												}
											} else {
												v82 = v6 + int32(-16)
												F_appendStringInfoChar(m, v82, int32(44))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int64(0)
												} else {
													if v36&int32(80) == int32(0) {
														v90 = F_range_bound_escape(m, v56)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int64(0)
														} else {
															F_appendStringInfoString(m, v82, v90)
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int64(0)
															} else {
																if v36&int32(4) != 0 {
																	v100 = int32(93)
																} else {
																	v100 = int32(41)
																}
																F_appendStringInfoChar(m, v6+int32(-16), v100)
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
																	return int64(0)
																} else {
																	v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																	v105 = v103
																	m.G0 = v8 - int32(-64)
																	return base.I64_extend_i32_u(v105)
																}
															}
														}
													} else {
														if v36&int32(4) != 0 {
															v100 = int32(93)
														} else {
															v100 = int32(41)
														}
														F_appendStringInfoChar(m, v6+int32(-16), v100)
														mBase = m.M
														v102 = m.ExcPending
														if v102 != 0 {
															return int64(0)
														} else {
															v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
															v105 = v103
															m.G0 = v8 - int32(-64)
															return base.I64_extend_i32_u(v105)
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
						v46 = v2
						if v36&int32(81) == int32(0) {
							v53 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
							v54 = F_OutputFunctionCall(m, v19+int32(4), v53)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								v56 = v54
								if v36&int32(1) != 0 {
									v60 = F_pstrdup(m, int32(_a_F_range_out_0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int64(0)
									} else {
										v105 = v60
										m.G0 = v8 - int32(-64)
										return base.I64_extend_i32_u(v105)
									}
								} else {
									v63 = v6 + int32(-16)
									F_initStringInfo(m, v63)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int64(0)
									} else {
										if v36&int32(2) != 0 {
											v70 = int32(91)
										} else {
											v70 = int32(40)
										}
										F_appendStringInfoChar(m, v63, v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int64(0)
										} else {
											if v36&int32(40) == int32(0) {
												v77 = F_range_bound_escape(m, v46)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int64(0)
												} else {
													F_appendStringInfoString(m, v63, v77)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int64(0)
													} else {
														v82 = v6 + int32(-16)
														F_appendStringInfoChar(m, v82, int32(44))
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int64(0)
														} else {
															if v36&int32(80) == int32(0) {
																v90 = F_range_bound_escape(m, v56)
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return int64(0)
																} else {
																	F_appendStringInfoString(m, v82, v90)
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return int64(0)
																	} else {
																		if v36&int32(4) != 0 {
																			v100 = int32(93)
																		} else {
																			v100 = int32(41)
																		}
																		F_appendStringInfoChar(m, v6+int32(-16), v100)
																		mBase = m.M
																		v102 = m.ExcPending
																		if v102 != 0 {
																			return int64(0)
																		} else {
																			v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																			v105 = v103
																			m.G0 = v8 - int32(-64)
																			return base.I64_extend_i32_u(v105)
																		}
																	}
																}
															} else {
																if v36&int32(4) != 0 {
																	v100 = int32(93)
																} else {
																	v100 = int32(41)
																}
																F_appendStringInfoChar(m, v6+int32(-16), v100)
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
																	return int64(0)
																} else {
																	v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																	v105 = v103
																	m.G0 = v8 - int32(-64)
																	return base.I64_extend_i32_u(v105)
																}
															}
														}
													}
												}
											} else {
												v82 = v6 + int32(-16)
												F_appendStringInfoChar(m, v82, int32(44))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int64(0)
												} else {
													if v36&int32(80) == int32(0) {
														v90 = F_range_bound_escape(m, v56)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int64(0)
														} else {
															F_appendStringInfoString(m, v82, v90)
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int64(0)
															} else {
																if v36&int32(4) != 0 {
																	v100 = int32(93)
																} else {
																	v100 = int32(41)
																}
																F_appendStringInfoChar(m, v6+int32(-16), v100)
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
																	return int64(0)
																} else {
																	v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																	v105 = v103
																	m.G0 = v8 - int32(-64)
																	return base.I64_extend_i32_u(v105)
																}
															}
														}
													} else {
														if v36&int32(4) != 0 {
															v100 = int32(93)
														} else {
															v100 = int32(41)
														}
														F_appendStringInfoChar(m, v6+int32(-16), v100)
														mBase = m.M
														v102 = m.ExcPending
														if v102 != 0 {
															return int64(0)
														} else {
															v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
															v105 = v103
															m.G0 = v8 - int32(-64)
															return base.I64_extend_i32_u(v105)
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v56 = v2
							if v36&int32(1) != 0 {
								v60 = F_pstrdup(m, int32(_a_F_range_out_0))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int64(0)
								} else {
									v105 = v60
									m.G0 = v8 - int32(-64)
									return base.I64_extend_i32_u(v105)
								}
							} else {
								v63 = v6 + int32(-16)
								F_initStringInfo(m, v63)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int64(0)
								} else {
									if v36&int32(2) != 0 {
										v70 = int32(91)
									} else {
										v70 = int32(40)
									}
									F_appendStringInfoChar(m, v63, v70)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int64(0)
									} else {
										if v36&int32(40) == int32(0) {
											v77 = F_range_bound_escape(m, v46)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int64(0)
											} else {
												F_appendStringInfoString(m, v63, v77)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int64(0)
												} else {
													v82 = v6 + int32(-16)
													F_appendStringInfoChar(m, v82, int32(44))
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int64(0)
													} else {
														if v36&int32(80) == int32(0) {
															v90 = F_range_bound_escape(m, v56)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int64(0)
															} else {
																F_appendStringInfoString(m, v82, v90)
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return int64(0)
																} else {
																	if v36&int32(4) != 0 {
																		v100 = int32(93)
																	} else {
																		v100 = int32(41)
																	}
																	F_appendStringInfoChar(m, v6+int32(-16), v100)
																	mBase = m.M
																	v102 = m.ExcPending
																	if v102 != 0 {
																		return int64(0)
																	} else {
																		v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																		v105 = v103
																		m.G0 = v8 - int32(-64)
																		return base.I64_extend_i32_u(v105)
																	}
																}
															}
														} else {
															if v36&int32(4) != 0 {
																v100 = int32(93)
															} else {
																v100 = int32(41)
															}
															F_appendStringInfoChar(m, v6+int32(-16), v100)
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return int64(0)
															} else {
																v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																v105 = v103
																m.G0 = v8 - int32(-64)
																return base.I64_extend_i32_u(v105)
															}
														}
													}
												}
											}
										} else {
											v82 = v6 + int32(-16)
											F_appendStringInfoChar(m, v82, int32(44))
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int64(0)
											} else {
												if v36&int32(80) == int32(0) {
													v90 = F_range_bound_escape(m, v56)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int64(0)
													} else {
														F_appendStringInfoString(m, v82, v90)
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int64(0)
														} else {
															if v36&int32(4) != 0 {
																v100 = int32(93)
															} else {
																v100 = int32(41)
															}
															F_appendStringInfoChar(m, v6+int32(-16), v100)
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return int64(0)
															} else {
																v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
																v105 = v103
																m.G0 = v8 - int32(-64)
																return base.I64_extend_i32_u(v105)
															}
														}
													}
												} else {
													if v36&int32(4) != 0 {
														v100 = int32(93)
													} else {
														v100 = int32(41)
													}
													F_appendStringInfoChar(m, v6+int32(-16), v100)
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return int64(0)
													} else {
														v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
														v105 = v103
														m.G0 = v8 - int32(-64)
														return base.I64_extend_i32_u(v105)
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
		}
	}
}
func F_range_overlaps_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v199 int64
	_ = v199
	var v200 int32
	_ = v200
	var v204 int64
	_ = v204
	var v207 int32
	_ = v207
	var v208 int64
	_ = v208
	var v209 int64
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v273 int64
	_ = v273
	var v276 int32
	_ = v276
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v14 == v15 {
		goto L19
	} else {
		goto L20
	}
L1:
	;
	m.G0 = v12 + int32(80)
	return v326
L2:
	;
	v326 = int32(0)
	goto L1
L3:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+42)))
	if v315 != 0 {
		goto L2
	} else {
		goto L125
	}
L4:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
	v278 = F_FunctionCall2Coll(m, l0+int32(212), v276, v273, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L22
	} else {
		goto L109
	}
L5:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+40)))
	if v267 != 0 {
		goto L3
	} else {
		goto L108
	}
L6:
	;
	v248 = int32(1)
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+40)))
	if v249 == v248 {
		goto L102
	} else {
		goto L103
	}
L7:
	;
	if v180&int32(1) != 0 {
		goto L2
	} else {
		goto L101
	}
L8:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
	v209 = F_FunctionCall2Coll(m, l0+int32(212), v207, v208, v204)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L22
	} else {
		goto L85
	}
L9:
	;
	if v41 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L10:
	;
	if base.B2i32(v188&int32(255) == v186)|base.B2i32(v186&int32(1) == int32(0)) != 0 {
		v246 = v186
		goto L6
	} else {
		goto L80
	}
L11:
	;
	if v40&int32(1) == int32(0) {
		goto L9
	} else {
		goto L78
	}
L12:
	;
	v160 = int32(1)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
	if v161 == v160 {
		goto L72
	} else {
		goto L73
	}
L13:
	;
	if v125 <= int32(0) {
		v326 = int32(1)
		goto L1
	} else {
		goto L69
	}
L14:
	;
	v143 = int32(1)
	if v126&v143 != 0 {
		v326 = v143
		goto L1
	} else {
		goto L67
	}
L15:
	;
	if v129&int32(1) != 0 {
		goto L11
	} else {
		goto L66
	}
L16:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+26)))
	if v139 != 0 {
		goto L11
	} else {
		goto L65
	}
L17:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	v123 = F_FunctionCall2Coll(m, l0+int32(212), v121, v118, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L22
	} else {
		goto L59
	}
L18:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
	if v114 != 0 {
		goto L16
	} else {
		goto L58
	}
L19:
	;
	F_range_deserialize(m, l0, l1, v12-int32(-64), v12+int32(32), v12+int32(15))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L22
	} else {
		goto L55
	}
L22:
	;
	return int32(0)
L23:
	;
	F_range_deserialize(m, l0, l2, v12+int32(48), v12+int32(16), v12+int32(14))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v35 = int32(0)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v36 != 0 {
		v326 = v35
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)))
	if v37&int32(1) != 0 {
		v326 = v35
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+56)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+72)))
	if v41 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v44&int32(1) == int32(0) {
		goto L12
	} else {
		goto L54
	}
L28:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+74)))
	if v40&int32(1) == int32(0) {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v40&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+58)))
	if v49 == v44 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	v51 = int32(1)
	if v44&v51 != 0 {
		v186 = v49
		v188 = v51
		goto L10
	} else {
		goto L33
	}
L33:
	;
	goto L12
L34:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+58)))
	if v56 != 0 {
		goto L18
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v12)+64))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
	v62 = F_FunctionCall2Coll(m, l0+int32(212), v59, v60, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L22
	} else {
		goto L39
	}
L37:
	;
	v246 = v4
	goto L6
L38:
	;
	if v71&int32(1) != 0 {
		goto L18
	} else {
		goto L53
	}
L39:
	;
	v64 = base.I32_wrap_i64(v62)
	if v64 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+57)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+73)))
	if v68 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if v64 < int32(0) {
		v204 = v60
		goto L8
	} else {
		goto L51
	}
L43:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+74)))
	if v67&int32(1) != 0 {
		goto L38
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if v67&int32(1) != 0 {
		goto L18
	} else {
		goto L49
	}
L46:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+58)))
	if v74 == v71 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	if v71&int32(1) == int32(0) {
		v204 = v60
		goto L8
	} else {
		goto L48
	}
L48:
	;
	goto L18
L49:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+58)))
	if v82&int32(1) == int32(0) {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	v204 = v60
	goto L8
L51:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
	if v89 == int32(0) {
		v118 = v60
		goto L17
	} else {
		goto L52
	}
L52:
	;
	goto L16
L53:
	;
	v204 = v60
	goto L8
L54:
	;
	goto L5
L55:
	;
	F_errmsg_internal(m, int32(_a_F_range_overlaps_internal_0), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L22
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_range_overlaps_internal_1), int32(858), int32(_a_F_range_overlaps_internal_2))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L22
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v12)+64))
	v118 = v115
	goto L17
L59:
	;
	v125 = base.I32_wrap_i64(v123)
	if v125 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+25)))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+73)))
	if v127 != 0 {
		goto L14
	} else {
		goto L61
	}
L61:
	;
	v128 = int32(1)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+74)))
	if v126&v128 != 0 {
		goto L15
	} else {
		goto L62
	}
L62:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+26)))
	if v129 == v132 {
		v326 = v128
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v129&int32(1) != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v326 = v128
	goto L1
L65:
	;
	v326 = int32(1)
	goto L1
L66:
	;
	v326 = v128
	goto L1
L67:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+26)))
	if v146&int32(1) == int32(0) {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	v326 = v143
	goto L1
L69:
	;
	if v40&int32(1) == int32(0) {
		v204 = v118
		goto L8
	} else {
		goto L70
	}
L70:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+58)))
	if v158 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v246 = v4
	goto L6
L72:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+26)))
	if v164 == v44 {
		v326 = v160
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v44&int32(1) != 0 {
		v326 = v160
		goto L1
	} else {
		goto L77
	}
L75:
	;
	if v44&int32(1) == int32(0) {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	v326 = v160
	goto L1
L77:
	;
	goto L11
L78:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+58)))
	if v41 == int32(0) {
		goto L7
	} else {
		goto L79
	}
L79:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+74)))
	v186 = v180
	v188 = v183
	goto L10
L80:
	;
	goto L2
L81:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v12)+64))
	v204 = v199
	goto L8
L82:
	;
	goto L83
L83:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+74)))
	if v200 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	goto L2
L85:
	;
	v211 = base.I32_wrap_i64(v209)
	if v211 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+73)))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+57)))
	if v215 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	if v211 < int32(0) {
		goto L2
	} else {
		goto L99
	}
L89:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+58)))
	if v214&int32(1) == int32(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	if v214&int32(1) != 0 {
		goto L5
	} else {
		goto L97
	}
L92:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+74)))
	if v218&int32(1)|base.B2i32(v225 == v218) != 0 {
		goto L5
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if v218&int32(1) != 0 {
		goto L5
	} else {
		goto L96
	}
L95:
	;
	goto L2
L96:
	;
	goto L2
L97:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+74)))
	if v232&int32(1) == int32(0) {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	goto L2
L99:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+40)))
	if v239 == int32(0) {
		v273 = v208
		goto L4
	} else {
		goto L100
	}
L100:
	;
	goto L3
L101:
	;
	v246 = v180
	goto L6
L102:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+42)))
	if v246 == v252 {
		v326 = v248
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	if v246&int32(1) == int32(0) {
		goto L2
	} else {
		goto L107
	}
L105:
	;
	if v246&int32(1) == int32(0) {
		goto L2
	} else {
		goto L106
	}
L106:
	;
	v326 = v248
	goto L1
L107:
	;
	v326 = v248
	goto L1
L108:
	;
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
	v273 = v268
	goto L4
L109:
	;
	v280 = base.I32_wrap_i64(v278)
	if v280 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+41)))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+57)))
	if v284 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if v280 <= int32(0) {
		v326 = int32(1)
		goto L1
	} else {
		goto L124
	}
L113:
	;
	v287 = int32(1)
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+58)))
	if v283&v287 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	goto L115
L115:
	;
	v299 = int32(1)
	if v283&v299 != 0 {
		v326 = v299
		goto L1
	} else {
		goto L122
	}
L116:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+42)))
	if v293 == v288 {
		v326 = v287
		goto L1
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if v288&int32(1) != 0 {
		goto L2
	} else {
		goto L121
	}
L119:
	;
	if v288&int32(1) != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	v326 = v287
	goto L1
L121:
	;
	v326 = v287
	goto L1
L122:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+42)))
	if v302&int32(1) == int32(0) {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	v326 = v299
	goto L1
L124:
	;
	goto L2
L125:
	;
	v326 = int32(1)
	goto L1
}
func F_range_overleft_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v10)>>(uint(int32(2))%32))-int32(1)))))
	if v16&int32(1) != 0 {
		v47 = v4
		m.G0 = v8 + int32(80)
		return v47
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		if v19 == int32(0) {
			v47 = v4
			m.G0 = v8 + int32(80)
			return v47
		} else {
			v25 = v8 + int32(48)
			F_range_deserialize(m, l0, l1, v8-int32(-64), v25, v8+int32(15))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v38 = v8 + int32(16)
				F_multirange_get_bounds(m, l0, l2, v32-int32(1), v8+int32(32), v38)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v41 = F_range_cmp_bounds(m, l0, v25, v38)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v47 = base.B2i32(v41 <= int32(0))
						m.G0 = v8 + int32(80)
						return v47
					}
				}
			}
		}
	}
}
func F_range_parse_bound(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	switch v14 - int32(41) {
	case 0, 3:
		goto L3
	case 1, 2:
		goto L2
	default:
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return v122
L2:
	;
	F_initStringInfo(m, v12+int32(32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v21)
	v122 = l1
	goto L1
L4:
	;
	if v14 != int32(93) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	v30 = l1
	v37 = int32(0)
	goto L8
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v117
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v119)
	v122 = v30
	goto L1
L10:
	;
	goto L9
L11:
	;
	v44 = v30 + int32(1)
	if v38 != int32(34) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	switch v38 - int32(41) {
	case 0, 3:
		goto L10
	case 1, 2:
		goto L11
	default:
		goto L13
	}
L13:
	;
	if v38 == int32(93) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_appendStringInfoChar(m, v12+int32(32), base.I32_extend8_s(v110))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L6
	} else {
		goto L40
	}
L16:
	;
	if v38 != int32(92) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if v37 == int32(0) {
		v30 = v44
		v37 = int32(1)
		goto L8
	} else {
		goto L38
	}
L19:
	;
	if v38 != 0 {
		v109 = v44
		v110 = v38
		v111 = v37
		goto L15
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v70 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	v49 = int32(0)
	v50 = F_errsave_start(m, l4)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	if v50 == int32(0) {
		v122 = v49
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg(m, int32(_a_F_range_parse_bound_0), v12)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v63 = F_errdetail(m, int32(_a_F_range_parse_bound_1), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	F_errsave_finish(m, l4, int32(_a_F_range_parse_bound_2), int32(2698), int32(_a_F_range_parse_bound_3))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v122 = v49
	goto L1
L29:
	;
	v73 = int32(0)
	v74 = F_errsave_start(m, l4)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v109 = v30 + int32(2)
	v110 = v70
	v111 = v37
	goto L15
L32:
	;
	if v74 == int32(0) {
		v122 = v73
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
	F_errmsg(m, int32(_a_F_range_parse_bound_0), v12+int32(16))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v89 = F_errdetail(m, int32(_a_F_range_parse_bound_1), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	F_errsave_finish(m, l4, int32(_a_F_range_parse_bound_2), int32(2706), int32(_a_F_range_parse_bound_3))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v122 = v73
	goto L1
L38:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v102 != int32(34) {
		v30 = v44
		v37 = int32(0)
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v109 = v30 + int32(2)
	v110 = int32(34)
	v111 = int32(1)
	goto L15
L40:
	;
	v30 = v109
	v37 = v111
	goto L8
}
func F_range_send(m *base.Module, l0 int32) int64 {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v91 int32
	_ = v91
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		F_check_stack_depth(m)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v20 = v8 + int32(-16)
			F_initStringInfo(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				v25 = F_get_range_io_data(m, l0, v23, int32(3))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
					F_range_deserialize(m, v27, v13, v8+int32(-32), v8+int32(-48), v8+int32(-49))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(base.Ui32(v36)>>(uint(int32(2))%32))-int32(1)))))
						F_pq_begintypsend(m, v20)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							F_enlargeStringInfo(m, v20, int32(1))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int64(0)
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
								*(*uint8)(unsafe.Add(mBase, uint32(v48+v49))) = uint8(v42)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v48 + int32(1)
								if v42&int32(41) == int32(0) {
									v61 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
									v62 = F_SendFunctionCall(m, v25+int32(4), v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int64(0)
									} else {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
										F_enlargeStringInfo(m, v20, int32(4))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int64(0)
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
											v73 = int32(4)
											v74 = int32(base.Ui32(v64)>>(uint(int32(2))%32)) - v73
											v75 = int32(16711935)
											*(*int32)(unsafe.Add(mBase, uint32(v68+v69))) = base.I32_rotr(v74&v75, int32(8)) | base.I32_rotr(v74, int32(24))&v75
											*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v68 + v73
											F_appendBinaryStringInfo(m, v20, v62+v73, v74)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int64(0)
											} else {
												if v42&int32(81) == int32(0) {
													v101 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
													v102 = F_SendFunctionCall(m, v25+int32(4), v101)
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int64(0)
													} else {
														v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
														v106 = v8 + int32(-16)
														F_enlargeStringInfo(m, v106, int32(4))
														mBase = m.M
														v109 = m.ExcPending
														if v109 != 0 {
															return int64(0)
														} else {
															v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
															v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
															v115 = int32(4)
															v116 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - v115
															v117 = int32(16711935)
															*(*int32)(unsafe.Add(mBase, uint32(v110+v111))) = base.I32_rotr(v116&v117, int32(8)) | base.I32_rotr(v116, int32(24))&v117
															*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v110 + v115
															F_appendBinaryStringInfo(m, v106, v102+v115, v116)
															mBase = m.M
															v133 = m.ExcPending
															if v133 != 0 {
																return int64(0)
															} else {
																v139 = v8 + int32(-16)
																v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
																v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v141))) = v142 << (uint(int32(2)) % 32)
																m.G0 = v10 - int32(-64)
																return base.I64_extend_i32_u(v141)
															}
														}
													}
												} else {
													v139 = v8 + int32(-16)
													v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
													v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v141))) = v142 << (uint(int32(2)) % 32)
													m.G0 = v10 - int32(-64)
													return base.I64_extend_i32_u(v141)
												}
											}
										}
									}
								} else {
									if v42&int32(81) == int32(0) {
										v101 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
										v102 = F_SendFunctionCall(m, v25+int32(4), v101)
										mBase = m.M
										v103 = m.ExcPending
										if v103 != 0 {
											return int64(0)
										} else {
											v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
											v106 = v8 + int32(-16)
											F_enlargeStringInfo(m, v106, int32(4))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int64(0)
											} else {
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
												v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
												v115 = int32(4)
												v116 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - v115
												v117 = int32(16711935)
												*(*int32)(unsafe.Add(mBase, uint32(v110+v111))) = base.I32_rotr(v116&v117, int32(8)) | base.I32_rotr(v116, int32(24))&v117
												*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v110 + v115
												F_appendBinaryStringInfo(m, v106, v102+v115, v116)
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int64(0)
												} else {
													v139 = v8 + int32(-16)
													v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
													v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v141))) = v142 << (uint(int32(2)) % 32)
													m.G0 = v10 - int32(-64)
													return base.I64_extend_i32_u(v141)
												}
											}
										}
									} else {
										v139 = v8 + int32(-16)
										v141 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
										v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v141))) = v142 << (uint(int32(2)) % 32)
										m.G0 = v10 - int32(-64)
										return base.I64_extend_i32_u(v141)
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
func F_range_upper(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
		if v18 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			if v19 == v16 {
				v29 = v18
				F_range_deserialize(m, v29, v12, v9+int32(32), v9+int32(16), v9+int32(15))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int64(0)
				} else {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v38 == int32(0) {
						v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)))
						if v41&int32(1) == int32(0) {
							v49 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
							v50 = v49
						} else {
							v46 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v46)
							v50 = int64(0)
						}
					} else {
						v46 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v46)
						v50 = int64(0)
					}
					m.G0 = v9 + int32(48)
					return v50
				}
			} else {
				v22 = F_lookup_type_cache(m, v16, int32(2048))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+200))
					if v24 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
							F_errmsg_internal(m, int32(_a_F_range_upper_0), v9)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_range_upper_1), int32(1946), int32(_a_F_range_upper_2))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v22
						v29 = v22
						F_range_deserialize(m, v29, v12, v9+int32(32), v9+int32(16), v9+int32(15))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int64(0)
						} else {
							v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
							if v38 == int32(0) {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)))
								if v41&int32(1) == int32(0) {
									v49 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
									v50 = v49
								} else {
									v46 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v46)
									v50 = int64(0)
								}
							} else {
								v46 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v46)
								v50 = int64(0)
							}
							m.G0 = v9 + int32(48)
							return v50
						}
					}
				}
			}
		} else {
			v22 = F_lookup_type_cache(m, v16, int32(2048))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+200))
				if v24 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
						F_errmsg_internal(m, int32(_a_F_range_upper_0), v9)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_range_upper_1), int32(1946), int32(_a_F_range_upper_2))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v22
					v29 = v22
					F_range_deserialize(m, v29, v12, v9+int32(32), v9+int32(16), v9+int32(15))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int64(0)
					} else {
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v38 == int32(0) {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)))
							if v41&int32(1) == int32(0) {
								v49 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
								v50 = v49
							} else {
								v46 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v46)
								v50 = int64(0)
							}
						} else {
							v46 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v46)
							v50 = int64(0)
						}
						m.G0 = v9 + int32(48)
						return v50
					}
				}
			}
		}
	}
}
