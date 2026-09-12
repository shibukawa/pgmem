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
				v29 = *(*int32)(unsafe.Add(mBase, _consts[50]))
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
				F_errmsg_internal(m, int32(264466), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495182), int32(832), int32(264932))
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
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
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	v9 = m.G0
	v11 = v9 - int32(224)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, int32(57), l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L13
	} else {
		goto L14
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L13
	} else {
		goto L110
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L13
	} else {
		goto L105
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L13
	} else {
		goto L100
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L13
	} else {
		goto L95
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L13
	} else {
		goto L91
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L13
	} else {
		goto L87
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L13
	} else {
		goto L83
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L13
	} else {
		goto L79
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L13
	} else {
		goto L75
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L13
	} else {
		goto L71
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L13
	} else {
		goto L68
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L13
	} else {
		goto L64
	}
L13:
	;
	return
L14:
	;
	if v14 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
	v18 = v16 + v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+119)))
	v22 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v23 = F_object_ownercheck(m, int32(1259), l1, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
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
	if v23 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v28 = F_get_rel_relkind(m, l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, _consts[383])))
	if v46 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	switch v28 - int32(73) {
	case 0, 32:
		goto L29
	default:
		v39 = int32(41)
		goto L24
	case 10:
		goto L28
	case 29:
		goto L25
	case 36:
		goto L26
	case 45:
		goto L27
	}
L23:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_aclcheck_error(m, int32(2), v41, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L30
	}
L24:
	;
	v41 = v39
	goto L23
L25:
	;
	v39 = int32(18)
	goto L24
L26:
	;
	v41 = int32(23)
	goto L23
L27:
	;
	v41 = int32(51)
	goto L23
L28:
	;
	v41 = int32(37)
	goto L23
L29:
	;
	v41 = int32(20)
	goto L23
L30:
	;
	goto L21
L31:
	;
	v50 = int32(1)
	if base.Ui32(l1) < base.Ui32(int32(12000)) {
		v58 = v50
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	v59 = int32(4)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v60 - int32(215) {
	case 0:
		goto L41
	case 1:
		goto L11
	case 2:
		v81 = v59
		goto L39
	default:
		goto L40
	}
L34:
	;
	if v58 != 0 {
		goto L12
	} else {
		goto L38
	}
L35:
	;
	goto L34
L36:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)+68))
	if v53 == int32(99) {
		v58 = v50
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v56 = F_isTempToastNamespace(m, v53)
	mBase = m.M
	v58 = v56
	goto L35
L38:
	;
	goto L33
L39:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v81+l3)))
	if base.B2i32(v84 == int32(37))&base.B2i32(v19&int32(255) != int32(83)) != 0 {
		goto L10
	} else {
		goto L47
	}
L40:
	;
	if v60 != int32(146) {
		goto L11
	} else {
		goto L46
	}
L41:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v18)+68))
	v66 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v68 = F_object_aclcheck(m, int32(2615), v64, v66, int64(512))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	if v68 == int32(0) {
		v81 = v59
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v18)+68))
	v74 = F_get_namespace_name(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	F_aclcheck_error(m, v68, int32(36), v74)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v81 = v59
	goto L39
L46:
	;
	v81 = int32(12)
	goto L39
L47:
	;
	if base.B2i32(v84 == int32(51))&base.B2i32(v19&int32(255) != int32(118)) != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	if base.B2i32(v84 == int32(23))&base.B2i32(v19&int32(255) != int32(109)) != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	if base.B2i32(v84 == int32(18))&base.B2i32(v19&int32(255) != int32(102)) != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	if base.B2i32(v84 == int32(49))&base.B2i32(v19&int32(255) != int32(99)) != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	v121 = v19 & int32(223)
	if v84 != int32(20) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	if base.B2i32(v19&int32(255) == int32(99))&base.B2i32(v84 != int32(49)) != 0 {
		goto L4
	} else {
		goto L56
	}
L53:
	;
	if v121 == int32(73) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v126 != int32(215) {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v136 == int32(217) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v121 == int32(73) {
		goto L3
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L63
	}
L60:
	;
	v142 = v19 & int32(255)
	if v142 == int32(99) {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	if v142 == int32(116) {
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
	v163 = m.ExcPending
	if v163 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v164
	F_errmsg(m, int32(328059), v11+int32(208))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L13
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(495648), int32(19611), int32(264730))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
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
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v180
	F_errmsg_internal(m, int32(487111), v11)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(495648), int32(19635), int32(264730))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
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
	v196 = m.ExcPending
	if v196 != 0 {
		goto L13
	} else {
		goto L72
	}
L72:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v197
	F_errmsg(m, int32(417149), v11+int32(16))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(495648), int32(19649), int32(264730))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
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
	v215 = m.ExcPending
	if v215 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v216
	F_errmsg(m, int32(32727), v11+int32(32))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(495648), int32(19654), int32(264730))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
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
	v234 = m.ExcPending
	if v234 != 0 {
		goto L13
	} else {
		goto L80
	}
L80:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v235
	F_errmsg(m, int32(32672), v11+int32(48))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L13
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(495648), int32(19659), int32(264730))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
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
	v253 = m.ExcPending
	if v253 != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v254
	F_errmsg(m, int32(394121), v11-int32(-64))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L13
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(495648), int32(19664), int32(264730))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
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
	v272 = m.ExcPending
	if v272 != 0 {
		goto L13
	} else {
		goto L88
	}
L88:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v273
	F_errmsg(m, int32(370768), v11+int32(80))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(495648), int32(19669), int32(264730))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
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
	v291 = m.ExcPending
	if v291 != 0 {
		goto L13
	} else {
		goto L92
	}
L92:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v292
	F_errmsg(m, int32(28604), v11+int32(192))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L13
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(495648), int32(19676), int32(264730))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
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
	v310 = m.ExcPending
	if v310 != 0 {
		goto L13
	} else {
		goto L96
	}
L96:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v311
	F_errmsg(m, int32(370797), v11+int32(112))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L13
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = int32(540847)
	F_errhint(m, int32(653134), v11+int32(96))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(495648), int32(19688), int32(264730))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
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
	v336 = m.ExcPending
	if v336 != 0 {
		goto L13
	} else {
		goto L101
	}
L101:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v337
	F_errmsg(m, int32(697993), v11+int32(128))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L13
	} else {
		goto L102
	}
L102:
	;
	F_errhint(m, int32(653576), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L13
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(495648), int32(19701), int32(264730))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
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
	v359 = m.ExcPending
	if v359 != 0 {
		goto L13
	} else {
		goto L106
	}
L106:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v360
	F_errmsg(m, int32(715737), v11+int32(160))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L13
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = int32(540847)
	F_errhint(m, int32(653134), v11+int32(144))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L13
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(495648), int32(19709), int32(264730))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
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
	v385 = m.ExcPending
	if v385 != 0 {
		goto L13
	} else {
		goto L111
	}
L111:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v386
	F_errmsg(m, int32(722748), v11+int32(176))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	F_errhint(m, int32(653576), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L13
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(495648), int32(19715), int32(264730))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
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
	var v12 int32
	_ = v12
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
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
	v12 = F_SearchSysCache1(m, int32(57), l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v16 = v14 + v15
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+119)))
			v20 = *(*int32)(unsafe.Add(mBase, _consts[279]))
			v21 = F_object_ownercheck(m, int32(1259), l1, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if v21 == int32(0) {
					v26 = F_get_rel_relkind(m, l1)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						switch v26 - int32(73) {
						case 0, 32:
							v39 = int32(20)
						default:
							v37 = int32(41)
							v39 = v37
						case 10:
							v39 = int32(37)
						case 29:
							v37 = int32(18)
							v39 = v37
						case 36:
							v39 = int32(23)
						case 45:
							v39 = int32(51)
						}
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						F_aclcheck_error(m, int32(2), v39, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, _consts[383])))
							if v44 == int32(0) {
								v48 = int32(1)
								if base.Ui32(l1) < base.Ui32(int32(12000)) {
									v56 = v48
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
									if v51 == int32(99) {
										v56 = v48
									} else {
										v54 = F_isTempToastNamespace(m, v51)
										mBase = m.M
										v56 = v54
									}
								}
								if v56 != 0 {
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
											F_errmsg(m, int32(328059), v9+int32(16))
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return
											} else {
												F_errfinish(m, int32(493538), int32(87), int32(23384))
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
									if v17&int32(253) != int32(112) {
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
												F_errmsg(m, int32(396094), v9)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return
												} else {
													F_errfinish(m, int32(493538), int32(93), int32(23384))
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
										F_ReleaseCatCache(m, v12)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							} else {
								if v17&int32(253) != int32(112) {
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
											F_errmsg(m, int32(396094), v9)
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												F_errfinish(m, int32(493538), int32(93), int32(23384))
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
									F_ReleaseCatCache(m, v12)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
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
					v44 = int32(*(*uint8)(unsafe.Add(mBase, _consts[383])))
					if v44 == int32(0) {
						v48 = int32(1)
						if base.Ui32(l1) < base.Ui32(int32(12000)) {
							v56 = v48
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
							if v51 == int32(99) {
								v56 = v48
							} else {
								v54 = F_isTempToastNamespace(m, v51)
								mBase = m.M
								v56 = v54
							}
						}
						if v56 != 0 {
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
									F_errmsg(m, int32(328059), v9+int32(16))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										F_errfinish(m, int32(493538), int32(87), int32(23384))
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
							if v17&int32(253) != int32(112) {
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
										F_errmsg(m, int32(396094), v9)
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return
										} else {
											F_errfinish(m, int32(493538), int32(93), int32(23384))
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
								F_ReleaseCatCache(m, v12)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					} else {
						if v17&int32(253) != int32(112) {
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
									F_errmsg(m, int32(396094), v9)
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return
									} else {
										F_errfinish(m, int32(493538), int32(93), int32(23384))
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
							F_ReleaseCatCache(m, v12)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
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
			v53 = int32(0)
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			if v55 != 0 {
				v59 = v53
				*(*int32)(unsafe.Add(mBase, uint32(l6))) = v59
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				if v61 != 0 {
					v65 = v53
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
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v57 = F_copyObjectImpl(m, v56)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					v59 = v57
					*(*int32)(unsafe.Add(mBase, uint32(l6))) = v59
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					if v61 != 0 {
						v65 = v53
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
				F_errmsg_internal(m, int32(145022), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					F_errfinish(m, int32(495385), int32(4652), int32(167738))
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
				v53 = int32(0)
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
				if v55 != 0 {
					v59 = v53
					*(*int32)(unsafe.Add(mBase, uint32(l6))) = v59
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					if v61 != 0 {
						v65 = v53
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
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					v57 = F_copyObjectImpl(m, v56)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						v59 = v57
						*(*int32)(unsafe.Add(mBase, uint32(l6))) = v59
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						if v61 != 0 {
							v65 = v53
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v9 == v10 {
		F_range_deserialize(m, l0, l1, v7+int32(40), v7+int32(24), v7+int32(15))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_range_deserialize(m, l0, l2, v7+int32(32), v7+int32(16), v7+int32(14))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(0)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
				if v31 != 0 {
					v104 = v30
					m.G0 = v7 + int32(48)
					return v104
				} else {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
					if v32 != 0 {
						v104 = v30
						m.G0 = v7 + int32(48)
						return v104
					} else {
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+44)))
						if v34 == int32(1) {
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+46)))
							if v33&int32(1) != 0 {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
								if v40 == v37&int32(255) {
									v99 = int32(0)
								} else {
									v46 = int32(1)
									if v37&v46 != 0 {
										v49 = int32(-1)
									} else {
										v49 = v46
									}
									v99 = v49
								}
							} else {
								v51 = int32(1)
								if v37&v51 != 0 {
									v54 = int32(-1)
								} else {
									v54 = v51
								}
								v99 = v54
							}
							v104 = base.B2i32(int32(0) < v99)
							m.G0 = v7 + int32(48)
							return v104
						} else {
							if v33&int32(1) != 0 {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
								if v59 != 0 {
									v60 = int32(1)
								} else {
									v60 = int32(-1)
								}
								v99 = v60
								v104 = base.B2i32(int32(0) < v99)
								m.G0 = v7 + int32(48)
								return v104
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
								v66 = F_FunctionCall2Coll(m, l0+int32(212), v63, v64, v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									if v66 != 0 {
										v99 = v66
									} else {
										v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)))
										v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+45)))
										if v69 == int32(0) {
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+46)))
											if v68&int32(1) == int32(0) {
												v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
												if v77 == v72&int32(255) {
													v99 = int32(0)
												} else {
													v82 = int32(1)
													if v72&v82 != 0 {
														v86 = v82
													} else {
														v86 = int32(-1)
													}
													v99 = v86
												}
											} else {
												v87 = int32(1)
												if v72&v87 != 0 {
													v91 = v87
												} else {
													v91 = int32(-1)
												}
												v99 = v91
											}
										} else {
											if v68&int32(1) != 0 {
												v99 = int32(0)
											} else {
												v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
												if v97 != 0 {
													v98 = int32(-1)
												} else {
													v98 = int32(1)
												}
												v99 = v98
											}
										}
									}
									v104 = base.B2i32(int32(0) < v99)
									m.G0 = v7 + int32(48)
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
			F_errmsg_internal(m, int32(326275), int32(0))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(495320), int32(715), int32(311881))
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
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
	v17 = v14
	v18 = l0
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
	switch v17 {
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
	v21 = v18 + int32(1)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	v17 = v22
	v18 = v21
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
	if v34 == int32(34) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	F_appendStringInfoChar(m, v8, v35)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L23
	}
L14:
	;
	F_appendStringInfoChar(m, v8, v35)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	if v34 == int32(92) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if v34 != 0 {
		goto L13
	} else {
		goto L17
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
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	m.G0 = v8 + int32(16)
	return v43
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
func F_range_cmp(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
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
	F_check_stack_depth(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v22 == v23 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L117
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v26 != 0 {
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
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L114
	}
L9:
	;
	F_range_deserialize(m, v37, v13, v10+int32(40), v10+int32(24), v10+int32(15))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v27 == v22 {
		v37 = v26
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v30 = F_lookup_type_cache(m, v22, int32(2048))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+200))
	if v32 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v30
	v37 = v30
	goto L9
L16:
	;
	F_range_deserialize(m, v37, v18, v10+int32(32), v10+int32(16), v10+int32(14))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v54 = int32(1)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	if v56 == v54 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v195 != v13 {
		goto L106
	} else {
		goto L107
	}
L19:
	;
	v192 = int32(0) - (v55 ^ int32(1))
	goto L18
L20:
	;
	goto L21
L21:
	;
	if v55&int32(1) != 0 {
		v192 = v54
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+36)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+44)))
	if v66 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
	if v127 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L24:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+46)))
	if v65&int32(1) != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	if v65&int32(1) != 0 {
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)))
	if v69 == v72 {
		goto L23
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v80 = int32(1)
	if v69&v80 != 0 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v75 = int32(1)
	if v69&v75 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v78 = int32(-1)
	goto L33
L32:
	;
	v78 = v75
	goto L33
L33:
	;
	v192 = v78
	goto L18
L34:
	;
	v83 = int32(-1)
	goto L36
L35:
	;
	v83 = v80
	goto L36
L36:
	;
	v192 = v83
	goto L18
L37:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)))
	if v88 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v37)+208))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v95 = F_FunctionCall2Coll(m, v37+int32(212), v92, v93, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v89 = int32(1)
	goto L42
L41:
	;
	v89 = int32(-1)
	goto L42
L42:
	;
	v192 = v89
	goto L18
L43:
	;
	if v95 != 0 {
		v192 = v95
		goto L18
	} else {
		goto L44
	}
L44:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+37)))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+45)))
	if v98 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+46)))
	if v97&int32(1) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	if v97&int32(1) != 0 {
		goto L23
	} else {
		goto L58
	}
L48:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)))
	if v101 == v106 {
		goto L23
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v113 = int32(1)
	if v101&v113 != 0 {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v108 = int32(1)
	if v101&v108 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v112 = v108
	goto L54
L53:
	;
	v112 = int32(-1)
	goto L54
L54:
	;
	v192 = v112
	goto L18
L55:
	;
	v117 = v113
	goto L57
L56:
	;
	v117 = int32(-1)
	goto L57
L57:
	;
	v192 = v117
	goto L18
L58:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)))
	if v122 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v123 = int32(-1)
	goto L61
L60:
	;
	v123 = int32(1)
	goto L61
L61:
	;
	v192 = v123
	goto L18
L62:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+30)))
	if v126&int32(1) != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L64
L64:
	;
	if v126&int32(1) != 0 {
		goto L77
	} else {
		goto L78
	}
L65:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	if v133 == v130&int32(255) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v144 = int32(1)
	if v130&v144 != 0 {
		goto L74
	} else {
		goto L75
	}
L68:
	;
	v192 = int32(0)
	goto L18
L69:
	;
	goto L70
L70:
	;
	v139 = int32(1)
	if v130&v139 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v142 = int32(-1)
	goto L73
L72:
	;
	v142 = v139
	goto L73
L73:
	;
	v192 = v142
	goto L18
L74:
	;
	v147 = int32(-1)
	goto L76
L75:
	;
	v147 = v144
	goto L76
L76:
	;
	v192 = v147
	goto L18
L77:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	if v152 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v37)+208))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v159 = F_FunctionCall2Coll(m, v37+int32(212), v156, v157, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L83
	}
L80:
	;
	v153 = int32(1)
	goto L82
L81:
	;
	v153 = int32(-1)
	goto L82
L82:
	;
	v192 = v153
	goto L18
L83:
	;
	if v159 != 0 {
		v192 = v159
		goto L18
	} else {
		goto L84
	}
L84:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)))
	if v162 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+30)))
	if v161&int32(1) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	if v161&int32(1) != 0 {
		goto L100
	} else {
		goto L101
	}
L88:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	if v170 == v165&int32(255) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v180 = int32(1)
	if v165&v180 != 0 {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	v192 = int32(0)
	goto L18
L92:
	;
	goto L93
L93:
	;
	v175 = int32(1)
	if v165&v175 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v179 = v175
	goto L96
L95:
	;
	v179 = int32(-1)
	goto L96
L96:
	;
	v192 = v179
	goto L18
L97:
	;
	v184 = v180
	goto L99
L98:
	;
	v184 = int32(-1)
	goto L99
L99:
	;
	v192 = v184
	goto L18
L100:
	;
	v192 = int32(0)
	goto L18
L101:
	;
	goto L102
L102:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	if v190 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v191 = int32(-1)
	goto L105
L104:
	;
	v191 = int32(1)
	goto L105
L105:
	;
	v192 = v191
	goto L18
L106:
	;
	F_pfree(m, v13)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v199 != v18 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	F_pfree(m, v18)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	m.G0 = v10 + int32(48)
	return v192
L113:
	;
	goto L112
L114:
	;
	F_errmsg_internal(m, int32(326275), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(495320), int32(1268), int32(236579))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v22
	F_errmsg_internal(m, int32(371313), v10)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(495320), int32(1776), int32(399876))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
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
func F_range_eq(m *base.Module, l0 int32) int32 {
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
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
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
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v33
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(371313), v9)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(495320), int32(1776), int32(399876))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
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
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v33
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(371313), v9)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495320), int32(1776), int32(399876))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
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
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v33
						}
					}
				}
			}
		}
	}
}
func F_range_ge(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_range_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2^int32(-1)) >> (uint(int32(31)) % 32))
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
			F_errmsg_internal(m, int32(481758), v8)
			mBase = m.M
			v173 = m.ExcPending
			if v173 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493796), int32(968), int32(402318))
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	switch l1 - int32(1) {
	case 0:
		v107 = F_range_before_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v108 = m.ExcPending
		if v108 != 0 {
			return int32(0)
		} else {
			v109 = v107
			m.G0 = v8 + int32(48)
			return v109
		}
	case 1:
		v12 = F_range_overleft_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v109 = v12
			m.G0 = v8 + int32(48)
			return v109
		}
	case 2:
		v16 = F_range_overlaps_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v109 = v16
			m.G0 = v8 + int32(48)
			return v109
		}
	case 3:
		v18 = F_range_overright_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v109 = v18
			m.G0 = v8 + int32(48)
			return v109
		}
	case 4:
		v20 = F_range_after_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v109 = v20
			m.G0 = v8 + int32(48)
			return v109
		}
	case 5:
		v22 = F_range_adjacent_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v109 = v22
			m.G0 = v8 + int32(48)
			return v109
		}
	case 6:
		v24 = F_range_contains_multirange_internal(m, l0, l2, l3)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v109 = v24
			m.G0 = v8 + int32(48)
			return v109
		}
	case 7:
		v26 = F_multirange_contains_range_internal(m, l0, l3, l2)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v109 = v26
			m.G0 = v8 + int32(48)
			return v109
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v97 = m.ExcPending
		if v97 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
			F_errmsg_internal(m, int32(481758), v8)
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493796), int32(1119), int32(402024))
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 17:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v34 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v28)>>(uint(int32(2))%32))-int32(1)))))
		if v34&int32(1) == int32(0) {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			if v39 != 0 {
				F_range_deserialize(m, l0, l2, v8+int32(40), v8+int32(32), v8+int32(7))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v63 = int32(0)
					F_multirange_get_bounds(m, l0, l3, v63, v8+int32(24), v8+int32(8))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
						F_multirange_get_bounds(m, l0, l3, v71-int32(1), v8+int32(8), v8+int32(16))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							v84 = F_range_cmp_bounds(m, l0, v8+int32(40), v8+int32(24))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								if v84 != 0 {
									v109 = v63
									m.G0 = v8 + int32(48)
									return v109
								} else {
									v90 = F_range_cmp_bounds(m, l0, v8+int32(32), v8+int32(16))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										v109 = base.B2i32(v90 == int32(0))
										m.G0 = v8 + int32(48)
										return v109
									}
								}
							}
						}
					}
				}
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v41)>>(uint(int32(2))%32))-int32(1)))))
				if v47&int32(1) == int32(0) {
					v109 = int32(0)
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					v109 = base.B2i32(v52 == int32(0))
				}
				m.G0 = v8 + int32(48)
				return v109
			}
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v41)>>(uint(int32(2))%32))-int32(1)))))
			if v47&int32(1) == int32(0) {
				v109 = int32(0)
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				v109 = base.B2i32(v52 == int32(0))
			}
			m.G0 = v8 + int32(48)
			return v109
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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v23 = (v19 - int32(1)) & int32(65535)
	v26 = F_palloc(m, v23*int32(12))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = v19 & int32(65535)
	if v29 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v40 = int32(1)
	goto L6
L4:
	;
	goto L5
L5:
	;
	F_qsort_arg(m, v26, v23, int32(12), int32(1488), l0)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(4)+v40<<(uint(int32(4))%32))))
	v53 = F_pg_detoast_datum(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v55 = int32(12)
	v57 = v26 + v40*v55
	*(*int32)(unsafe.Add(mBase, uint32(v57-v55))) = v40
	v62 = v57 - int32(8)
	if l3 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v78 = (v40 + int32(1)) & int32(65535)
	if base.Ui32(v78) <= base.Ui32(v23) {
		v40 = v78
		goto L6
	} else {
		goto L15
	}
L10:
	;
	F_range_deserialize(m, l0, v53, v17+int32(8), v62, v17+int32(7))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_range_deserialize(m, l0, v53, v62, v17+int32(8), v17+int32(7))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	goto L9
L15:
	;
	goto L7
L16:
	;
	v98 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v98
	if v29 != int32(1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v104 = int32(1)
	if base.Ui32(v23) <= base.Ui32(v104) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v185 = v5
	v186 = v5
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v186
	m.G0 = v17 + int32(16)
	return
L20:
	;
	v109 = v104
	goto L22
L21:
	;
	v109 = v23
	goto L22
L22:
	;
	v118 = int32(0)
	v122 = v5
	v123 = v5
	goto L23
L23:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v26+v118*int32(12))))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(4)+v130<<(uint(int32(4))%32))))
	v135 = F_pg_detoast_datum(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	v185 = v171
	v186 = v172
	goto L19
L25:
	;
	if base.Ui32(v118) < base.Ui32(int32(base.Ui32(v23)>>(uint(v104)%32))) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v174 = v118 + int32(1)
	if v174 != v109 {
		v118 = v174
		v122 = v171
		v123 = v172
		goto L23
	} else {
		goto L40
	}
L27:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v138 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v154 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v146 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v144 + v146
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v149+v144<<(uint(v146)%32)))) = uint16(v130)
	v171 = v122
	v172 = v145
	goto L26
L31:
	;
	v144 = v138
	v145 = v135
	goto L30
L32:
	;
	goto L33
L33:
	;
	v141 = F_range_super_union(m, l0, v123, v135)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v144 = v143
	v145 = v141
	goto L30
L35:
	;
	v162 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v160 + v162
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v165+v160<<(uint(v162)%32)))) = uint16(v130)
	v171 = v161
	v172 = v123
	goto L26
L36:
	;
	v160 = v154
	v161 = v135
	goto L35
L37:
	;
	goto L38
L38:
	;
	v157 = F_range_super_union(m, l0, v122, v135)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v160 = v159
	v161 = v157
	goto L35
L40:
	;
	goto L24
}
func F_range_merge_from_multirange(m *base.Module, l0 int32) int32 {
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
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
						return int32(0)
					} else {
						v64 = v31
						m.G0 = v8 + int32(48)
						return v64
					}
				case 1:
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
					v35 = F_multirange_get_range(m, v33, v11, int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v64 = v35
						m.G0 = v8 + int32(48)
						return v64
					}
				default:
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
					F_multirange_get_bounds(m, v37, v11, int32(0), v8+int32(40), v8+int32(32))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
						F_multirange_get_bounds(m, v45, v11, v46-int32(1), v8+int32(24), v8+int32(16))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
							v60 = int32(0)
							v62 = F_make_range(m, v55, v8+int32(40), v8+int32(16), v60, v60)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v64 = v62
								m.G0 = v8 + int32(48)
								return v64
							}
						}
					}
				}
			} else {
				v21 = F_lookup_type_cache(m, v15, int32(65536))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+296))
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
							F_errmsg_internal(m, int32(371186), v8)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495315), int32(558), int32(399871))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
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
								return int32(0)
							} else {
								v64 = v31
								m.G0 = v8 + int32(48)
								return v64
							}
						case 1:
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
							v35 = F_multirange_get_range(m, v33, v11, int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v64 = v35
								m.G0 = v8 + int32(48)
								return v64
							}
						default:
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
							F_multirange_get_bounds(m, v37, v11, int32(0), v8+int32(40), v8+int32(32))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
								F_multirange_get_bounds(m, v45, v11, v46-int32(1), v8+int32(24), v8+int32(16))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
									v60 = int32(0)
									v62 = F_make_range(m, v55, v8+int32(40), v8+int32(16), v60, v60)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v64 = v62
										m.G0 = v8 + int32(48)
										return v64
									}
								}
							}
						}
					}
				}
			}
		} else {
			v21 = F_lookup_type_cache(m, v15, int32(65536))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+296))
				if v23 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
						F_errmsg_internal(m, int32(371186), v8)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495315), int32(558), int32(399871))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
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
							return int32(0)
						} else {
							v64 = v31
							m.G0 = v8 + int32(48)
							return v64
						}
					case 1:
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
						v35 = F_multirange_get_range(m, v33, v11, int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v64 = v35
							m.G0 = v8 + int32(48)
							return v64
						}
					default:
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
						F_multirange_get_bounds(m, v37, v11, int32(0), v8+int32(40), v8+int32(32))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							F_multirange_get_bounds(m, v45, v11, v46-int32(1), v8+int32(24), v8+int32(16))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v28)+296))
								v60 = int32(0)
								v62 = F_make_range(m, v55, v8+int32(40), v8+int32(16), v60, v60)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v64 = v62
									m.G0 = v8 + int32(48)
									return v64
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_range_minus(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v39 int32
	_ = v39
	var v42 int32
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
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			if v19 == v20 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
				if v23 != 0 {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					if v24 == v19 {
						v34 = v23
						v35 = F_range_minus_internal(m, v34, v12, v17)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							if v35 == int32(0) {
								v39 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v39)
								v42 = int32(0)
							} else {
								v42 = v35
							}
							m.G0 = v9 + int32(16)
							return v42
						}
					} else {
						v27 = F_lookup_type_cache(m, v19, int32(2048))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+200))
							if v29 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
									F_errmsg_internal(m, int32(371313), v9)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(495320), int32(1776), int32(399876))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v27
								v34 = v27
								v35 = F_range_minus_internal(m, v34, v12, v17)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									if v35 == int32(0) {
										v39 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v39)
										v42 = int32(0)
									} else {
										v42 = v35
									}
									m.G0 = v9 + int32(16)
									return v42
								}
							}
						}
					}
				} else {
					v27 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+200))
						if v29 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(371313), v9)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(495320), int32(1776), int32(399876))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v27
							v34 = v27
							v35 = F_range_minus_internal(m, v34, v12, v17)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								if v35 == int32(0) {
									v39 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v39)
									v42 = int32(0)
								} else {
									v42 = v35
								}
								m.G0 = v9 + int32(16)
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
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(326275), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495320), int32(983), int32(115370))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
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
func F_range_out(m *base.Module, l0 int32) int32 {
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
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
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		F_check_stack_depth(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v19 = F_get_range_io_data(m, l0, v17, int32(1))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				F_range_deserialize(m, v21, v11, v8+int32(24), v8+int32(16), v8+int32(15))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+int32(base.Ui32(v30)>>(uint(int32(2))%32))-int32(1)))))
					if v36&int32(41) == int32(0) {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
						v44 = F_OutputFunctionCall(m, v19+int32(4), v43)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = v44
							if v36&int32(81) == int32(0) {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
								v54 = F_OutputFunctionCall(m, v19+int32(4), v53)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									v56 = v54
									if v36&int32(1) != 0 {
										v60 = F_pstrdup(m, int32(8965))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v110 = v60
											m.G0 = v8 + int32(48)
											return v110
										}
									} else {
										F_initStringInfo(m, v8+int32(32))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											if v36&int32(2) != 0 {
												v72 = int32(91)
											} else {
												v72 = int32(40)
											}
											F_appendStringInfoChar(m, v8+int32(32), v72)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												if v36&int32(40) == int32(0) {
													v81 = F_range_bound_escape(m, v46)
													mBase = m.M
													v82 = m.ExcPending
													if v82 != 0 {
														return int32(0)
													} else {
														F_appendStringInfoString(m, v8+int32(32), v81)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int32(0)
														} else {
															F_appendStringInfoChar(m, v8+int32(32), int32(44))
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return int32(0)
															} else {
																if v36&int32(80) == int32(0) {
																	v96 = F_range_bound_escape(m, v56)
																	mBase = m.M
																	v97 = m.ExcPending
																	if v97 != 0 {
																		return int32(0)
																	} else {
																		F_appendStringInfoString(m, v8+int32(32), v96)
																		mBase = m.M
																		v99 = m.ExcPending
																		if v99 != 0 {
																			return int32(0)
																		} else {
																			if v36&int32(4) != 0 {
																				v106 = int32(93)
																			} else {
																				v106 = int32(41)
																			}
																			F_appendStringInfoChar(m, v8+int32(32), v106)
																			mBase = m.M
																			v108 = m.ExcPending
																			if v108 != 0 {
																				return int32(0)
																			} else {
																				v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																				v110 = v109
																				m.G0 = v8 + int32(48)
																				return v110
																			}
																		}
																	}
																} else {
																	if v36&int32(4) != 0 {
																		v106 = int32(93)
																	} else {
																		v106 = int32(41)
																	}
																	F_appendStringInfoChar(m, v8+int32(32), v106)
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return int32(0)
																	} else {
																		v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																		v110 = v109
																		m.G0 = v8 + int32(48)
																		return v110
																	}
																}
															}
														}
													}
												} else {
													F_appendStringInfoChar(m, v8+int32(32), int32(44))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														if v36&int32(80) == int32(0) {
															v96 = F_range_bound_escape(m, v56)
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int32(0)
															} else {
																F_appendStringInfoString(m, v8+int32(32), v96)
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	if v36&int32(4) != 0 {
																		v106 = int32(93)
																	} else {
																		v106 = int32(41)
																	}
																	F_appendStringInfoChar(m, v8+int32(32), v106)
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return int32(0)
																	} else {
																		v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																		v110 = v109
																		m.G0 = v8 + int32(48)
																		return v110
																	}
																}
															}
														} else {
															if v36&int32(4) != 0 {
																v106 = int32(93)
															} else {
																v106 = int32(41)
															}
															F_appendStringInfoChar(m, v8+int32(32), v106)
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return int32(0)
															} else {
																v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																v110 = v109
																m.G0 = v8 + int32(48)
																return v110
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
									v60 = F_pstrdup(m, int32(8965))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v110 = v60
										m.G0 = v8 + int32(48)
										return v110
									}
								} else {
									F_initStringInfo(m, v8+int32(32))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										if v36&int32(2) != 0 {
											v72 = int32(91)
										} else {
											v72 = int32(40)
										}
										F_appendStringInfoChar(m, v8+int32(32), v72)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											if v36&int32(40) == int32(0) {
												v81 = F_range_bound_escape(m, v46)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													F_appendStringInfoString(m, v8+int32(32), v81)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														F_appendStringInfoChar(m, v8+int32(32), int32(44))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int32(0)
														} else {
															if v36&int32(80) == int32(0) {
																v96 = F_range_bound_escape(m, v56)
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int32(0)
																} else {
																	F_appendStringInfoString(m, v8+int32(32), v96)
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int32(0)
																	} else {
																		if v36&int32(4) != 0 {
																			v106 = int32(93)
																		} else {
																			v106 = int32(41)
																		}
																		F_appendStringInfoChar(m, v8+int32(32), v106)
																		mBase = m.M
																		v108 = m.ExcPending
																		if v108 != 0 {
																			return int32(0)
																		} else {
																			v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																			v110 = v109
																			m.G0 = v8 + int32(48)
																			return v110
																		}
																	}
																}
															} else {
																if v36&int32(4) != 0 {
																	v106 = int32(93)
																} else {
																	v106 = int32(41)
																}
																F_appendStringInfoChar(m, v8+int32(32), v106)
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return int32(0)
																} else {
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																	v110 = v109
																	m.G0 = v8 + int32(48)
																	return v110
																}
															}
														}
													}
												}
											} else {
												F_appendStringInfoChar(m, v8+int32(32), int32(44))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													if v36&int32(80) == int32(0) {
														v96 = F_range_bound_escape(m, v56)
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															F_appendStringInfoString(m, v8+int32(32), v96)
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																if v36&int32(4) != 0 {
																	v106 = int32(93)
																} else {
																	v106 = int32(41)
																}
																F_appendStringInfoChar(m, v8+int32(32), v106)
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return int32(0)
																} else {
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																	v110 = v109
																	m.G0 = v8 + int32(48)
																	return v110
																}
															}
														}
													} else {
														if v36&int32(4) != 0 {
															v106 = int32(93)
														} else {
															v106 = int32(41)
														}
														F_appendStringInfoChar(m, v8+int32(32), v106)
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return int32(0)
														} else {
															v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
															v110 = v109
															m.G0 = v8 + int32(48)
															return v110
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
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
							v54 = F_OutputFunctionCall(m, v19+int32(4), v53)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v56 = v54
								if v36&int32(1) != 0 {
									v60 = F_pstrdup(m, int32(8965))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v110 = v60
										m.G0 = v8 + int32(48)
										return v110
									}
								} else {
									F_initStringInfo(m, v8+int32(32))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										if v36&int32(2) != 0 {
											v72 = int32(91)
										} else {
											v72 = int32(40)
										}
										F_appendStringInfoChar(m, v8+int32(32), v72)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											if v36&int32(40) == int32(0) {
												v81 = F_range_bound_escape(m, v46)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return int32(0)
												} else {
													F_appendStringInfoString(m, v8+int32(32), v81)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int32(0)
													} else {
														F_appendStringInfoChar(m, v8+int32(32), int32(44))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int32(0)
														} else {
															if v36&int32(80) == int32(0) {
																v96 = F_range_bound_escape(m, v56)
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int32(0)
																} else {
																	F_appendStringInfoString(m, v8+int32(32), v96)
																	mBase = m.M
																	v99 = m.ExcPending
																	if v99 != 0 {
																		return int32(0)
																	} else {
																		if v36&int32(4) != 0 {
																			v106 = int32(93)
																		} else {
																			v106 = int32(41)
																		}
																		F_appendStringInfoChar(m, v8+int32(32), v106)
																		mBase = m.M
																		v108 = m.ExcPending
																		if v108 != 0 {
																			return int32(0)
																		} else {
																			v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																			v110 = v109
																			m.G0 = v8 + int32(48)
																			return v110
																		}
																	}
																}
															} else {
																if v36&int32(4) != 0 {
																	v106 = int32(93)
																} else {
																	v106 = int32(41)
																}
																F_appendStringInfoChar(m, v8+int32(32), v106)
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return int32(0)
																} else {
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																	v110 = v109
																	m.G0 = v8 + int32(48)
																	return v110
																}
															}
														}
													}
												}
											} else {
												F_appendStringInfoChar(m, v8+int32(32), int32(44))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													if v36&int32(80) == int32(0) {
														v96 = F_range_bound_escape(m, v56)
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															F_appendStringInfoString(m, v8+int32(32), v96)
															mBase = m.M
															v99 = m.ExcPending
															if v99 != 0 {
																return int32(0)
															} else {
																if v36&int32(4) != 0 {
																	v106 = int32(93)
																} else {
																	v106 = int32(41)
																}
																F_appendStringInfoChar(m, v8+int32(32), v106)
																mBase = m.M
																v108 = m.ExcPending
																if v108 != 0 {
																	return int32(0)
																} else {
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																	v110 = v109
																	m.G0 = v8 + int32(48)
																	return v110
																}
															}
														}
													} else {
														if v36&int32(4) != 0 {
															v106 = int32(93)
														} else {
															v106 = int32(41)
														}
														F_appendStringInfoChar(m, v8+int32(32), v106)
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return int32(0)
														} else {
															v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
															v110 = v109
															m.G0 = v8 + int32(48)
															return v110
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
								v60 = F_pstrdup(m, int32(8965))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v110 = v60
									m.G0 = v8 + int32(48)
									return v110
								}
							} else {
								F_initStringInfo(m, v8+int32(32))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									if v36&int32(2) != 0 {
										v72 = int32(91)
									} else {
										v72 = int32(40)
									}
									F_appendStringInfoChar(m, v8+int32(32), v72)
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										if v36&int32(40) == int32(0) {
											v81 = F_range_bound_escape(m, v46)
											mBase = m.M
											v82 = m.ExcPending
											if v82 != 0 {
												return int32(0)
											} else {
												F_appendStringInfoString(m, v8+int32(32), v81)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													F_appendStringInfoChar(m, v8+int32(32), int32(44))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														if v36&int32(80) == int32(0) {
															v96 = F_range_bound_escape(m, v56)
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int32(0)
															} else {
																F_appendStringInfoString(m, v8+int32(32), v96)
																mBase = m.M
																v99 = m.ExcPending
																if v99 != 0 {
																	return int32(0)
																} else {
																	if v36&int32(4) != 0 {
																		v106 = int32(93)
																	} else {
																		v106 = int32(41)
																	}
																	F_appendStringInfoChar(m, v8+int32(32), v106)
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return int32(0)
																	} else {
																		v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																		v110 = v109
																		m.G0 = v8 + int32(48)
																		return v110
																	}
																}
															}
														} else {
															if v36&int32(4) != 0 {
																v106 = int32(93)
															} else {
																v106 = int32(41)
															}
															F_appendStringInfoChar(m, v8+int32(32), v106)
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return int32(0)
															} else {
																v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																v110 = v109
																m.G0 = v8 + int32(48)
																return v110
															}
														}
													}
												}
											}
										} else {
											F_appendStringInfoChar(m, v8+int32(32), int32(44))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												if v36&int32(80) == int32(0) {
													v96 = F_range_bound_escape(m, v56)
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														F_appendStringInfoString(m, v8+int32(32), v96)
														mBase = m.M
														v99 = m.ExcPending
														if v99 != 0 {
															return int32(0)
														} else {
															if v36&int32(4) != 0 {
																v106 = int32(93)
															} else {
																v106 = int32(41)
															}
															F_appendStringInfoChar(m, v8+int32(32), v106)
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return int32(0)
															} else {
																v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
																v110 = v109
																m.G0 = v8 + int32(48)
																return v110
															}
														}
													}
												} else {
													if v36&int32(4) != 0 {
														v106 = int32(93)
													} else {
														v106 = int32(41)
													}
													F_appendStringInfoChar(m, v8+int32(32), v106)
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
														return int32(0)
													} else {
														v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
														v110 = v109
														m.G0 = v8 + int32(48)
														return v110
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
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
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
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
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v12 == v13 {
		goto L19
	} else {
		goto L20
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return v314
L2:
	;
	v314 = int32(0)
	goto L1
L3:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+30)))
	if v305 != 0 {
		goto L2
	} else {
		goto L127
	}
L4:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v271 = F_FunctionCall2Coll(m, l0+int32(212), v269, v265, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L22
	} else {
		goto L111
	}
L5:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
	if v262 != 0 {
		goto L3
	} else {
		goto L110
	}
L6:
	;
	v243 = int32(1)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
	if v244 == v243 {
		goto L104
	} else {
		goto L105
	}
L7:
	;
	if v178&int32(1) != 0 {
		goto L2
	} else {
		goto L103
	}
L8:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v207 = F_FunctionCall2Coll(m, l0+int32(212), v205, v206, v200)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L22
	} else {
		goto L86
	}
L9:
	;
	if v37 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L10:
	;
	v185 = int32(255)
	if v183&v185 == v182&v185 {
		v240 = v182
		goto L6
	} else {
		goto L80
	}
L11:
	;
	if v36&int32(1) == int32(0) {
		goto L9
	} else {
		goto L78
	}
L12:
	;
	v157 = int32(1)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
	if v158 == v157 {
		goto L72
	} else {
		goto L73
	}
L13:
	;
	if v119 <= int32(0) {
		v314 = int32(1)
		goto L1
	} else {
		goto L69
	}
L14:
	;
	v140 = int32(1)
	if v121&v140 != 0 {
		v314 = v140
		goto L1
	} else {
		goto L67
	}
L15:
	;
	if v124&int32(1) != 0 {
		goto L11
	} else {
		goto L66
	}
L16:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	if v136 != 0 {
		goto L11
	} else {
		goto L65
	}
L17:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v119 = F_FunctionCall2Coll(m, l0+int32(212), v117, v112, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L22
	} else {
		goto L59
	}
L18:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
	if v110 != 0 {
		goto L16
	} else {
		goto L58
	}
L19:
	;
	F_range_deserialize(m, l0, l1, v10+int32(40), v10+int32(24), v10+int32(15))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
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
	v97 = m.ExcPending
	if v97 != 0 {
		goto L22
	} else {
		goto L55
	}
L22:
	;
	return int32(0)
L23:
	;
	F_range_deserialize(m, l0, l2, v10+int32(32), v10+int32(16), v10+int32(14))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v33 = int32(0)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	if v34 != 0 {
		v314 = v33
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)))
	if v35 != 0 {
		v314 = v33
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+36)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+44)))
	if v37 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v40&int32(1) == int32(0) {
		goto L12
	} else {
		goto L54
	}
L28:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+46)))
	if v36&int32(1) == int32(0) {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v36&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)))
	if v45 == v40&int32(255) {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	if v40&int32(1) != 0 {
		v182 = v40
		v183 = v45
		goto L10
	} else {
		goto L33
	}
L33:
	;
	goto L12
L34:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)))
	if v53 != 0 {
		goto L18
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v59 = F_FunctionCall2Coll(m, l0+int32(212), v56, v57, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L22
	} else {
		goto L39
	}
L37:
	;
	v240 = v53
	goto L6
L38:
	;
	if v67&int32(1) != 0 {
		goto L18
	} else {
		goto L53
	}
L39:
	;
	if v59 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+37)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+45)))
	if v64 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if v59 < int32(0) {
		v200 = v57
		goto L8
	} else {
		goto L51
	}
L43:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+46)))
	if v63&int32(1) != 0 {
		goto L38
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if v63&int32(1) != 0 {
		goto L18
	} else {
		goto L49
	}
L46:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)))
	if v70 == v67&int32(255) {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	if v67&int32(1) == int32(0) {
		v200 = v57
		goto L8
	} else {
		goto L48
	}
L48:
	;
	goto L18
L49:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)))
	if v80 == int32(0) {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	v200 = v57
	goto L8
L51:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
	if v85 == int32(0) {
		v112 = v57
		goto L17
	} else {
		goto L52
	}
L52:
	;
	goto L16
L53:
	;
	v200 = v57
	goto L8
L54:
	;
	goto L5
L55:
	;
	F_errmsg_internal(m, int32(326275), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L22
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(495320), int32(854), int32(311833))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
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
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v112 = v111
	goto L17
L59:
	;
	if v119 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+45)))
	if v122 != 0 {
		goto L14
	} else {
		goto L61
	}
L61:
	;
	v123 = int32(1)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+46)))
	if v121&v123 != 0 {
		goto L15
	} else {
		goto L62
	}
L62:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	if v127 == v124&int32(255) {
		v314 = v123
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v124&int32(1) != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v314 = v123
	goto L1
L65:
	;
	v314 = int32(1)
	goto L1
L66:
	;
	v314 = v123
	goto L1
L67:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	if v143 == int32(0) {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	v314 = v140
	goto L1
L69:
	;
	if v36&int32(1) == int32(0) {
		v200 = v112
		goto L8
	} else {
		goto L70
	}
L70:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)))
	if v153 == int32(0) {
		v240 = v153
		goto L6
	} else {
		goto L71
	}
L71:
	;
	goto L2
L72:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+22)))
	if v161 == v40&int32(255) {
		v314 = v157
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v40&int32(1) != 0 {
		v314 = v157
		goto L1
	} else {
		goto L77
	}
L75:
	;
	if v40&int32(1) == int32(0) {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	v314 = v157
	goto L1
L77:
	;
	goto L11
L78:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)))
	if v37 == int32(0) {
		goto L7
	} else {
		goto L79
	}
L79:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+46)))
	v182 = v181
	v183 = v178
	goto L10
L80:
	;
	if v183&int32(1) == int32(0) {
		v240 = v183
		goto L6
	} else {
		goto L81
	}
L81:
	;
	goto L2
L82:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v200 = v196
	goto L8
L83:
	;
	goto L84
L84:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+46)))
	if v197&int32(1) != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	goto L2
L86:
	;
	if v207 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+45)))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+37)))
	if v212 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	if v207 < int32(0) {
		goto L2
	} else {
		goto L101
	}
L90:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)))
	if v211&int32(1) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	if v211&int32(1) != 0 {
		goto L5
	} else {
		goto L99
	}
L93:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+46)))
	if v220 == v215&int32(255) {
		goto L5
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v215&int32(1) != 0 {
		goto L5
	} else {
		goto L98
	}
L96:
	;
	if v215&int32(1) != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	goto L2
L98:
	;
	goto L2
L99:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+46)))
	if v230 == int32(0) {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	goto L2
L101:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
	if v235 == int32(0) {
		v265 = v206
		goto L4
	} else {
		goto L102
	}
L102:
	;
	goto L3
L103:
	;
	v240 = v178
	goto L6
L104:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+30)))
	if v247 == v240&int32(255) {
		v314 = v243
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v240&int32(1) == int32(0) {
		goto L2
	} else {
		goto L109
	}
L107:
	;
	if v240&int32(1) == int32(0) {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	v314 = v243
	goto L1
L109:
	;
	v314 = v243
	goto L1
L110:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v265 = v263
	goto L4
L111:
	;
	if v271 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+37)))
	if v276 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L114
L114:
	;
	if v271 <= int32(0) {
		v314 = int32(1)
		goto L1
	} else {
		goto L126
	}
L115:
	;
	v279 = int32(1)
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+38)))
	if v275&v279 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L117
L117:
	;
	v293 = int32(1)
	if v275&v293 != 0 {
		v314 = v293
		goto L1
	} else {
		goto L124
	}
L118:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+30)))
	if v285 == v280&int32(255) {
		v314 = v279
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	if v280&int32(1) != 0 {
		goto L2
	} else {
		goto L123
	}
L121:
	;
	if v280&int32(1) != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v314 = v279
	goto L1
L123:
	;
	v314 = v279
	goto L1
L124:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+30)))
	if v296 == int32(0) {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	v314 = v293
	goto L1
L126:
	;
	goto L2
L127:
	;
	v314 = int32(1)
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v10)>>(uint(int32(2))%32))-int32(1)))))
	if v16&int32(1) != 0 {
		v49 = v4
		m.G0 = v8 + int32(48)
		return v49
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		if v19 == int32(0) {
			v49 = v4
			m.G0 = v8 + int32(48)
			return v49
		} else {
			F_range_deserialize(m, l0, l1, v8+int32(40), v8+int32(32), v8+int32(15))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				F_multirange_get_bounds(m, l0, l2, v32-int32(1), v8+int32(24), v8+int32(16))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v45 = F_range_cmp_bounds(m, l0, v8+int32(32), v8+int32(16))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v49 = base.B2i32(v45 <= int32(0))
						m.G0 = v8 + int32(48)
						return v49
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
	var v36 int32
	_ = v36
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
	v36 = int32(0)
	goto L8
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v36 != 0 {
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
	if v36 == int32(0) {
		v30 = v44
		v36 = int32(1)
		goto L8
	} else {
		goto L38
	}
L19:
	;
	if v38 != 0 {
		v109 = v44
		v110 = v38
		v111 = v36
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
	F_errmsg(m, int32(727925), v12)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	F_errdetail(m, int32(579619), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	F_errsave_finish(m, l4, int32(495320), int32(2528), int32(425431))
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
	v111 = v36
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
	F_errmsg(m, int32(727925), v12+int32(16))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	F_errdetail(m, int32(579619), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	F_errsave_finish(m, l4, int32(495320), int32(2536), int32(425431))
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
		v36 = int32(0)
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
	v36 = v111
	goto L8
}
func F_range_send(m *base.Module, l0 int32) int32 {
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
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
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = F_makeStringInfo(m)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			F_check_stack_depth(m)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				v23 = F_get_range_io_data(m, l0, v21, int32(3))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					F_range_deserialize(m, v25, v13, v10+int32(24), v10+int32(16), v10+int32(15))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(base.Ui32(v34)>>(uint(int32(2))%32))-int32(1)))))
						F_pq_begintypsend(m, v17)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_enlargeStringInfo(m, v17, int32(1))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
								*(*uint8)(unsafe.Add(mBase, uint32(v46+v47))) = uint8(v40)
								*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v46 + int32(1)
								if v40&int32(41) == int32(0) {
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
									v60 = F_SendFunctionCall(m, v23+int32(4), v59)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
										F_enlargeStringInfo(m, v17, int32(4))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
											v71 = int32(4)
											v72 = int32(base.Ui32(v62)>>(uint(int32(2))%32)) - v71
											v73 = int32(24)
											v75 = int32(65280)
											v77 = int32(8)
											*(*int32)(unsafe.Add(mBase, uint32(v66+v67))) = v72<<(uint(v73)%32) | v72&v75<<(uint(v77)%32) | (int32(base.Ui32(v72)>>(uint(v77)%32))&v75 | int32(base.Ui32(v72)>>(uint(v73)%32)))
											*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v66 + v71
											F_pq_sendbytes(m, v17, v60+v71, v72)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												if v40&int32(81) == int32(0) {
													v105 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
													v106 = F_SendFunctionCall(m, v23+int32(4), v105)
													mBase = m.M
													v107 = m.ExcPending
													if v107 != 0 {
														return int32(0)
													} else {
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
														F_enlargeStringInfo(m, v17, int32(4))
														mBase = m.M
														v111 = m.ExcPending
														if v111 != 0 {
															return int32(0)
														} else {
															v112 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
															v113 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
															v117 = int32(4)
															v118 = int32(base.Ui32(v108)>>(uint(int32(2))%32)) - v117
															v119 = int32(24)
															v121 = int32(65280)
															v123 = int32(8)
															*(*int32)(unsafe.Add(mBase, uint32(v112+v113))) = v118<<(uint(v119)%32) | v118&v121<<(uint(v123)%32) | (int32(base.Ui32(v118)>>(uint(v123)%32))&v121 | int32(base.Ui32(v118)>>(uint(v119)%32)))
															*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v112 + v117
															F_pq_sendbytes(m, v17, v106+v117, v118)
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return int32(0)
															} else {
																v146 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
																v147 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v146))) = v147 << (uint(int32(2)) % 32)
																m.G0 = v10 + int32(32)
																return v146
															}
														}
													}
												} else {
													v146 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
													v147 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v146))) = v147 << (uint(int32(2)) % 32)
													m.G0 = v10 + int32(32)
													return v146
												}
											}
										}
									}
								} else {
									if v40&int32(81) == int32(0) {
										v105 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
										v106 = F_SendFunctionCall(m, v23+int32(4), v105)
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int32(0)
										} else {
											v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
											F_enlargeStringInfo(m, v17, int32(4))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int32(0)
											} else {
												v112 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
												v113 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
												v117 = int32(4)
												v118 = int32(base.Ui32(v108)>>(uint(int32(2))%32)) - v117
												v119 = int32(24)
												v121 = int32(65280)
												v123 = int32(8)
												*(*int32)(unsafe.Add(mBase, uint32(v112+v113))) = v118<<(uint(v119)%32) | v118&v121<<(uint(v123)%32) | (int32(base.Ui32(v118)>>(uint(v123)%32))&v121 | int32(base.Ui32(v118)>>(uint(v119)%32)))
												*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v112 + v117
												F_pq_sendbytes(m, v17, v106+v117, v118)
												mBase = m.M
												v141 = m.ExcPending
												if v141 != 0 {
													return int32(0)
												} else {
													v146 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
													v147 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v146))) = v147 << (uint(int32(2)) % 32)
													m.G0 = v10 + int32(32)
													return v146
												}
											}
										}
									} else {
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
										v147 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v146))) = v147 << (uint(int32(2)) % 32)
										m.G0 = v10 + int32(32)
										return v146
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
func F_range_upper(m *base.Module, l0 int32) int32 {
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
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
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			if v18 == v15 {
				v28 = v17
				F_range_deserialize(m, v28, v11, v8+int32(24), v8+int32(16), v8+int32(15))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					if v37 == int32(0) {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
						if v40 != int32(1) {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
							v47 = v46
						} else {
							v43 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
							v47 = int32(0)
						}
					} else {
						v43 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
						v47 = int32(0)
					}
					m.G0 = v8 + int32(32)
					return v47
				}
			} else {
				v21 = F_lookup_type_cache(m, v15, int32(2048))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
							F_errmsg_internal(m, int32(371313), v8)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495320), int32(1776), int32(399876))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
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
						F_range_deserialize(m, v28, v11, v8+int32(24), v8+int32(16), v8+int32(15))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							if v37 == int32(0) {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
								if v40 != int32(1) {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
									v47 = v46
								} else {
									v43 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
									v47 = int32(0)
								}
							} else {
								v43 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
								v47 = int32(0)
							}
							m.G0 = v8 + int32(32)
							return v47
						}
					}
				}
			}
		} else {
			v21 = F_lookup_type_cache(m, v15, int32(2048))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
				if v23 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
						F_errmsg_internal(m, int32(371313), v8)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495320), int32(1776), int32(399876))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
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
					F_range_deserialize(m, v28, v11, v8+int32(24), v8+int32(16), v8+int32(15))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						if v37 == int32(0) {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
							if v40 != int32(1) {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
								v47 = v46
							} else {
								v43 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
								v47 = int32(0)
							}
						} else {
							v43 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v43)
							v47 = int32(0)
						}
						m.G0 = v8 + int32(32)
						return v47
					}
				}
			}
		}
	}
}
