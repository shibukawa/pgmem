package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PathNameOpenFilePerm(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v96 int64
	_ = v96
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int64
	_ = v117
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	v4 = int32(0)
	v18 = F_strlen(m, l0)
	mBase = m.M
	v20 = v18 + int32(1)
	v21 = F_emscripten_builtin_malloc(m, v20)
	mBase = m.M
	if v21 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L42
	} else {
		goto L57
	}
L2:
	;
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v26 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v25 = F___memcpy(m, v21, l0, v20)
	mBase = m.M
	v26 = v25
	goto L2
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[0]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v29 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L42
	} else {
		goto L53
	}
L9:
	;
	v32 = int32(32)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[1]))
	v36 = v34 << (uint(int32(1)) % 32)
	if base.Ui32(v36) <= base.Ui32(v32) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v240 = v29
	v243 = v28
	goto L11
L11:
	;
	v252 = v243 + v240*int32(48)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+12)) = v253
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[2]))
	if v256 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L12:
	;
	v39 = v32
	goto L14
L13:
	;
	v39 = v36
	goto L14
L14:
	;
	v42 = F_emscripten_builtin_realloc(m, v28, v39*int32(48))
	mBase = m.M
	if v42 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[0])) = v42
	if base.Ui32(v39) <= base.Ui32(v34) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42+v39*int32(48)-int32(36)))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v34
	*(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[1])) = v39
	v240 = v34
	v243 = v42
	goto L11
L17:
	;
	v49 = v42 & int32(3)
	if v34+int32(1) != v39 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v53 = v39 - v34
	v64 = v34
	v65 = v4
	v68 = v4
	goto L21
L19:
	;
	v164 = v34
	v168 = v4
	goto L20
L20:
	;
	v175 = v42 + v164*int32(48)
	if v49 != 0 {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	v75 = v42 + v64*int32(48)
	if v49 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	if v53&int32(1) == int32(0) {
		goto L16
	} else {
		goto L32
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = int32(-1)
	v110 = v64 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v110
	v114 = v42 + v110*int32(48)
	if v49 != 0 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v81 = v42 + (v34+v68)*int32(48)
	v82 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v81)+40)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81)+32)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81)+16)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = v82
	goto L23
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+44)) = int32(0)
	v96 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v75)+36)) = v96
	*(*int64)(unsafe.Add(mBase, uint32(v75)+28)) = v96
	*(*int64)(unsafe.Add(mBase, uint32(v75)+20)) = v96
	*(*int64)(unsafe.Add(mBase, uint32(v75)+12)) = v96
	*(*int64)(unsafe.Add(mBase, uint32(v75)+4)) = v96
	goto L23
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(-1)
	v148 = int32(2)
	v149 = v64 + v148
	*(*int32)(unsafe.Add(mBase, uint32(v114)+12)) = v149
	v152 = v68 + v148
	v154 = v65 + v148
	if v154 != v53&int32(-2) {
		v64 = v149
		v65 = v154
		v68 = v152
		goto L21
	} else {
		goto L31
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+44)) = int32(0)
	v117 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v114)+36)) = v117
	*(*int64)(unsafe.Add(mBase, uint32(v114)+28)) = v117
	*(*int64)(unsafe.Add(mBase, uint32(v114)+20)) = v117
	*(*int64)(unsafe.Add(mBase, uint32(v114)+12)) = v117
	*(*int64)(unsafe.Add(mBase, uint32(v114)+4)) = v117
	goto L27
L29:
	;
	goto L30
L30:
	;
	v130 = (v34+v68)*int32(48) + v42
	v131 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v130)+88)) = v131
	*(*int64)(unsafe.Add(mBase, uint32(v130)+80)) = v131
	*(*int64)(unsafe.Add(mBase, uint32(v130)+72)) = v131
	*(*int64)(unsafe.Add(mBase, uint32(v130-int32(-64)))) = v131
	*(*int64)(unsafe.Add(mBase, uint32(v130)+56)) = v131
	*(*int64)(unsafe.Add(mBase, uint32(v130)+48)) = v131
	goto L27
L31:
	;
	goto L22
L32:
	;
	v164 = v149
	v168 = v152
	goto L20
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+12)) = v164 + int32(1)
	goto L16
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+44)) = int32(0)
	v178 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v175)+36)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v175)+28)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v175)+20)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v175)+12)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v175)+4)) = v178
	goto L33
L35:
	;
	goto L36
L36:
	;
	v191 = v42 + (v34+v168)*int32(48)
	v192 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v191)+40)) = v192
	*(*int64)(unsafe.Add(mBase, uint32(v191)+32)) = v192
	*(*int64)(unsafe.Add(mBase, uint32(v191)+24)) = v192
	*(*int64)(unsafe.Add(mBase, uint32(v191)+16)) = v192
	*(*int64)(unsafe.Add(mBase, uint32(v191)+8)) = v192
	*(*int64)(unsafe.Add(mBase, uint32(v191))) = v192
	goto L33
L37:
	;
	v319 = l1 | int32(_a_F_PathNameOpenFilePerm_0)
	v320 = F_BasicOpenFilePerm(m, l0, v319, l2)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L42
	} else {
		goto L46
	}
L38:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[3]))
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[4]))
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[5]))
	if v262+(v264+v256) < v260 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	goto L40
L40:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[0]))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	F_LruDelete(m, v285)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L37
L42:
	;
	return int32(0)
L43:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[2]))
	if v291 <= int32(0) {
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[3]))
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[4]))
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[5]))
	if v295 <= v297+(v299+v291) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = v320
	if v320 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[6]))
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[0]))
	v331 = v328 + v240*int32(48)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+32))
	if v332 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v346 = int32(_a_F_PathNameOpenFilePerm_1)
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[2])) = v348 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v252)+40)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v252)+36)) = v319 & int32(-705)
	*(*int32)(unsafe.Add(mBase, uint32(v252)+32)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v252)+24)) = int64(0)
	v359 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v252)+8)) = v359
	*(*uint16)(unsafe.Add(mBase, uint32(v252)+4)) = uint16(v359)
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[0]))
	v365 = int32(48)
	v367 = v364 + v240*v365
	*(*int32)(unsafe.Add(mBase, uint32(v367)+16)) = v359
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v364)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v367)+20)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v364)+20)) = v240
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v367)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v364+v373*v365)+16)) = v240
	return v240
L50:
	;
	F_emscripten_builtin_free(m, v332)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v331)+32)) = int32(0)
	goto L52
L51:
	;
	goto L52
L52:
	;
	v336 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v331)+4)) = uint16(v336)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v328)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v331)+12)) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v328)+12)) = v240
	F_emscripten_builtin_free(m, v26)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_PathNameOpenFilePerm[6])) = v326
	return int32(-1)
L53:
	;
	F_errcode(m, int32(_a_F_PathNameOpenFilePerm_2))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L42
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(_a_F_PathNameOpenFilePerm_3), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L42
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_PathNameOpenFilePerm_4), int32(1597), int32(_a_F_PathNameOpenFilePerm_5))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L42
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
	F_errcode(m, int32(_a_F_PathNameOpenFilePerm_2))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L42
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(_a_F_PathNameOpenFilePerm_3), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L42
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_PathNameOpenFilePerm_4), int32(1436), int32(_a_F_PathNameOpenFilePerm_6))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L42
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_assign_search_path(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_assign_search_path[0])) = uint8(v4)
	return
}
func F_compare_path_costs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v48 int32
	_ = v48
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v8 != v9 {
		if v8 < v9 {
			v14 = int32(-1)
		} else {
			v14 = int32(1)
		}
		return v14
	} else {
		if l2 == int32(0) {
			v18 = int32(-1)
			v19 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
			v20 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
			if base.F64_lt(v19, v20) != 0 {
				v48 = v18
				return v48
			} else {
				if base.F64_gt(v19, v20) != 0 {
					return int32(1)
				} else {
					v25 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
					v26 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
					if base.F64_lt(v25, v26) != 0 {
						v48 = v18
					} else {
						if base.F64_gt(v25, v26) == int32(0) {
							v48 = int32(0)
						} else {
							v48 = int32(1)
						}
					}
					return v48
				}
			}
		} else {
			v32 = int32(-1)
			v33 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
			v34 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			if base.F64_lt(v33, v34) != 0 {
				v48 = v32
				return v48
			} else {
				if base.F64_gt(v33, v34) != 0 {
					return int32(1)
				} else {
					v39 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
					v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
					if base.F64_lt(v39, v40) != 0 {
						v48 = v32
					} else {
						if base.F64_gt(v39, v40) != 0 {
							v48 = int32(1)
						} else {
							v48 = int32(0)
						}
					}
					return v48
				}
			}
		}
	}
}
func F_create_gather_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	v14 = F_palloc0(m, int32(88))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(1597727834410)
		v23 = F_get_baserel_parampathinfo(m, l0, l1, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = l2
			v26 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v26
			*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v26
			*(*uint16)(unsafe.Add(mBase, uint32(v14)+20)) = uint16(v26)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v23
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v26)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v33
			if v33 == v26 {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
				v40 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v40
				*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v39
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(v40)
			} else {
			}
			if l4 != 0 {
				v50 = l4
			} else {
				if v23 != 0 {
					v50 = v23 + int32(8)
				} else {
					v50 = l1 + int32(16)
				}
			}
			v51 = *(*float64)(unsafe.Add(mBase, uint32(v50)))
			*(*float64)(unsafe.Add(mBase, uint32(v14)+32)) = v51
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+40))
			v55 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
			v57 = *(*float64)(unsafe.Add(mBase, _c_F_create_gather_path[0]))
			v58 = *(*float64)(unsafe.Add(mBase, uint32(v53)+56))
			v59 = *(*float64)(unsafe.Add(mBase, uint32(v53)+48))
			v61 = *(*float64)(unsafe.Add(mBase, _c_F_create_gather_path[1]))
			v62 = base.F64_add(v59, v61)
			*(*float64)(unsafe.Add(mBase, uint32(v14)+48)) = v62
			*(*float64)(unsafe.Add(mBase, uint32(v14)+56)) = base.F64_add(v62, base.F64_add(base.F64_mul(v57, v51), base.F64_sub(v58, v59)))
			*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v54 + base.B2i32(v55&int64(16384) == int64(0))
			return v14
		}
	}
}
func F_get_path_all(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
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
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v131 int32
	_ = v131
	var v147 int64
	_ = v147
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v16 + int32(16)
	return v147
L4:
	;
	v131 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v131)
	v147 = int64(0)
	goto L3
L5:
	;
	v26 = F_array_contains_nulls(m, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	goto L4
L8:
	;
	goto L9
L9:
	;
	F_deconstruct_array_builtin(m, v24, int32(25), v16+int32(12), v16+int32(8), v16+int32(4))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v39 = F_palloc_mul(m, int32(4), v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v43 = F_palloc_mul(m, int32(4), v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if int32(0) < v45 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = int32(0)
	goto L16
L14:
	;
	v106 = v45
	goto L15
L15:
	;
	v115 = F_get_worker(m, v19, v39, v43, v106, l1)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L31
	}
L16:
	;
	v63 = v52 << (uint(int32(2)) % 32)
	v64 = v39 + v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v52<<(uint(int32(3))%32))))
	v70 = F_text_to_cstring(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	v106 = v100
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v70
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v73 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v99 = v52 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v99 < v100 {
		v52 = v99
		goto L16
	} else {
		goto L30
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_get_path_all[0])) = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v79 = F_strtol(m, v77, v16, int32(10))
	mBase = m.M
	goto L23
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63+v43))) = int32(-2147483648)
	goto L19
L23:
	;
	v80 = int32(-2147483648)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v81 == v82 {
		v89 = v80
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63+v43))) = v89
	goto L19
L25:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v84 != 0 {
		v89 = v80
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_get_path_all[0]))
	if v87 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v88 = int32(-2147483648)
	goto L29
L28:
	;
	v88 = v79
	goto L29
L29:
	;
	v89 = v88
	goto L24
L30:
	;
	goto L17
L31:
	;
	if v115 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v147 = base.I64_extend_i32_u(v115)
	goto L3
L33:
	;
	goto L34
L34:
	;
	goto L4
}
func F_path_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v205 int32
	_ = v205
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	v10 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = l0
	goto L1
L1:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v29-int32(9)))&base.B2i32(v29 != int32(32)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v17
	v43 = base.B2i32(v29 == int32(91))
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v43)
	if v29 == int32(91) {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	v17 = v17 + int32(1)
	goto L1
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	m.G0 = v15 + int32(16)
	return v238
L7:
	;
	v219 = int32(0)
	v220 = F_errsave_start(m, l8)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L28
	} else {
		goto L50
	}
L8:
	;
	v123 = int32(0)
	v125 = v111
	v128 = l3
	v134 = v123
	goto L26
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v106
	v111 = v106
	v122 = int32(1)
	goto L8
L10:
	;
	if l1 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v49 != int32(40) {
		v111 = v17
		v122 = v10
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v106 = v17 + int32(1)
	goto L9
L14:
	;
	v61 = v17
	goto L15
L15:
	;
	v65 = v61 + int32(1)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if base.Ui32(v66-int32(9)) < base.Ui32(int32(5)) {
		v61 = v65
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v76 = F_strlen(m, v17)
	mBase = m.M
	v83 = v76 + int32(1)
	goto L21
L17:
	;
	switch v66 - int32(32) {
	case 0:
		v61 = v65
		goto L15
	default:
		goto L18
	case 8:
		v106 = v65
		goto L9
	}
L18:
	;
	goto L16
L19:
	;
	if v95 != v17 {
		v111 = v17
		v122 = v10
		goto L8
	} else {
		goto L25
	}
L20:
	;
	goto L19
L21:
	;
	v85 = int32(0)
	if v83 == v85 {
		v95 = v85
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v95 = v90
	goto L20
L23:
	;
	v89 = v83 - int32(1)
	v90 = v17 + v89
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v91 != int32(40) {
		v83 = v89
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v106 = v65
	goto L9
L26:
	;
	v141 = F_pair_decode(m, v125, v128, v128+int32(8), v15+int32(12), l6, l7, l8)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v122 != 0 {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	return int32(0)
L29:
	;
	if v141 == int32(0) {
		v238 = v123
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v148 == int32(44) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v152 = v147 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v152
	v154 = v152
	goto L33
L32:
	;
	v154 = v147
	goto L33
L33:
	;
	v158 = v134 + int32(1)
	if v158 != l2 {
		v125 = v154
		v128 = v128 + int32(16)
		v134 = v158
		goto L26
	} else {
		goto L34
	}
L34:
	;
	goto L27
L35:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
	if v160 != int32(41) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v190 = v154
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v190
	if l5 != 0 {
		goto L46
	} else {
		goto L47
	}
L38:
	;
	if v160 != int32(93) {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v168 = v154
	goto L43
L41:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v165 != int32(1) {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v181 = v168 + int32(1)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	if base.B2i32(base.Ui32(v182-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v182 == int32(32)) != 0 {
		v168 = v181
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v190 = v181
	goto L37
L45:
	;
	goto L44
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v190
	v238 = int32(1)
	goto L6
L47:
	;
	goto L48
L48:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v205 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v238 = int32(1)
	goto L6
L50:
	;
	if v220 == int32(0) {
		v238 = v219
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L28
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l6
	F_errmsg(m, int32(_a_F_path_decode_0), v15)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L28
	} else {
		goto L53
	}
L53:
	;
	F_errsave_finish(m, l8, int32(_a_F_path_decode_1), int32(337), int32(_a_F_path_decode_2))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L28
	} else {
		goto L54
	}
L54:
	;
	v238 = v219
	goto L6
}
func F_path_isclosed(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		return base.I64_extend_i32_u(base.B2i32(v7 != int32(0)))
	}
}
func F_path_isopen(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		return base.I64_extend_i32_u(base.B2i32(v7 == int32(0)))
	}
}
func F_path_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v49 float64
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pq_getmsgbyte(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v13 = F_pq_getmsgint(m, v7, int32(4))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.Ui32(int32(-134217726)) < base.Ui32(v13-int32(134217726)) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v13<<(uint(int32(4))%32) + int32(16)
	v24 = F_palloc(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v26 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = base.B2i32(v8 != v26)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v23 << (uint(int32(2)) % 32)
	v37 = int32(0)
	goto L8
L8:
	;
	v45 = v24 + int32(16) + v37<<(uint(int32(4))%32)
	v46 = F_pq_getmsgfloat8(m, v7)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	return base.I64_extend_i32_u(v24)
L10:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v45))) = v46
	v49 = F_pq_getmsgfloat8(m, v7)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v45)+8)) = v49
	v53 = v37 + int32(1)
	if v53 != v13 {
		v37 = v53
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(_a_F_path_recv_0), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_path_recv_1), int32(1551), int32(_a_F_path_recv_2))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_push_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v19 = l4 - l1
	v20 = F_palloc0_mul(m, int32(4), v19)
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
	v23 = l1 + int32(1)
	if l4 <= v23 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v155 = int32(2)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v20+v19<<(uint(v155)%32)-int32(4))))
	if v161 == int32(16) {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	v34 = v23
	goto L5
L5:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v34))))
	if v39 != 0 {
		goto L3
	} else {
		goto L7
	}
L6:
	;
	goto L3
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2+v34<<(uint(int32(3))%32))))
	v44 = F_text_to_cstring(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_push_path[0])) = int32(0)
	v52 = F_strtol(m, v44, v16+int32(12), int32(10))
	mBase = m.M
	goto L9
L9:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v44 == v57 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+(v34-l1)<<(uint(int32(2))%32)))) = v136
	v139 = v34 + int32(1)
	if v139 != l4 {
		v34 = v139
		goto L5
	} else {
		goto L26
	}
L11:
	;
	F_pushJsonbValue(m, l0, int32(4), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L18
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(1)
	v67 = F_strlen(m, v44)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v67
	F_pushJsonbValue(m, l0, int32(6), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v59 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_push_path[0]))
	if v61 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	F_pushJsonbValue(m, l0, int32(1), v16+int32(16))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v136 = int32(17)
	goto L10
L18:
	;
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v83
	if v83 < v52 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v94 = v52
	goto L22
L20:
	;
	goto L21
L21:
	;
	v136 = int32(16)
	goto L10
L22:
	;
	F_pushJsonbValue(m, l0, int32(3), v16+int32(48))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	v105 = int32(1)
	if base.Ui32(v105) < base.Ui32(v94) {
		v94 = v94 - v105
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	goto L6
L27:
	;
	v164 = int32(3)
	goto L29
L28:
	;
	v164 = v155
	goto L29
L29:
	;
	F_pushJsonbValue(m, l0, v164, l5)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v168 = l4 - int32(1)
	if v168 <= l1 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	m.G0 = v16 + int32(80)
	return
L32:
	;
	v177 = v168
	goto L33
L33:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v177))))
	if v184 != 0 {
		goto L31
	} else {
		goto L35
	}
L34:
	;
	goto L31
L35:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v20+(v177-l1)<<(uint(int32(2))%32))))
	if v191 == int32(17) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v194 = int32(7)
	goto L38
L37:
	;
	v194 = int32(5)
	goto L38
L38:
	;
	F_pushJsonbValue(m, l0, v194, int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v199 = v177 - int32(1)
	if l1 < v199 {
		v177 = v199
		goto L33
	} else {
		goto L40
	}
L40:
	;
	goto L34
}
