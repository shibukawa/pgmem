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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	v4 = int32(0)
	v21 = F_strlen(m, l0)
	mBase = m.M
	v23 = v21 + int32(1)
	v24 = F_emscripten_builtin_malloc(m, v23)
	mBase = m.M
	if v24 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L36
	} else {
		goto L51
	}
L2:
	;
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v29 = int32(0)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v28 = F___memcpy(m, v24, l0, v23)
	mBase = m.M
	v29 = v28
	goto L2
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v32 == int32(0) {
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
	v314 = m.ExcPending
	if v314 != 0 {
		goto L36
	} else {
		goto L47
	}
L9:
	;
	v35 = int32(32)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[759]))
	v39 = v37 << (uint(int32(1)) % 32)
	if base.Ui32(v39) <= base.Ui32(v35) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v163 = v31
	v164 = v32
	goto L11
L11:
	;
	v178 = v163 + v164*int32(48)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+12)) = v179
	v182 = *(*int32)(unsafe.Add(mBase, _consts[751]))
	if v182 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L12:
	;
	v42 = v35
	goto L14
L13:
	;
	v42 = v39
	goto L14
L14:
	;
	v45 = F_emscripten_builtin_realloc(m, v31, v42*int32(48))
	mBase = m.M
	if v45 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[262])) = v45
	if base.Ui32(v37) < base.Ui32(v42) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v53 = int32(48)
	v54 = v37 * v53
	v59 = v54 + v45
	v68 = v37
	v71 = v4
	goto L19
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45+v42*int32(48)-int32(36)))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+12)) = v37
	*(*int32)(unsafe.Add(mBase, _consts[759])) = v42
	v163 = v45
	v164 = v37
	goto L11
L19:
	;
	v84 = v45 + v68*int32(48)
	if v45&int32(3) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L18
L21:
	;
	v122 = int32(1)
	v123 = v68 + v122
	*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = int32(-1)
	v128 = v71 + v122
	if v128 != v42-v37 {
		v68 = v123
		v71 = v128
		goto L19
	} else {
		goto L30
	}
L22:
	;
	if base.Ui32(v84+int32(48)) <= base.Ui32(v84) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v108 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v84)+4)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v84)+44)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v84)+36)) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v84)+28)) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v84)+20)) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v84)+12)) = v108
	goto L21
L25:
	;
	v91 = v71 * int32(48)
	v97 = v91 + (v59 + v53)
	v98 = v91 + (v59 + int32(4))
	if base.Ui32(v98) < base.Ui32(v97) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v100 = v97
	goto L28
L27:
	;
	v100 = v98
	goto L28
L28:
	;
	v107 = F__emscripten_memset_bulkmem(m, v59+v91, base.I32_extend8_s(int32(0)), (v45^int32(-1)-v54+v71*int32(-48)+v100)&int32(-4)+int32(4))
	mBase = m.M
	goto L29
L29:
	;
	goto L21
L30:
	;
	goto L20
L31:
	;
	v251 = l1 | int32(524288)
	v252 = F_BasicOpenFilePerm(m, l0, v251, l2)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L36
	} else {
		goto L40
	}
L32:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	v188 = *(*int32)(unsafe.Add(mBase, _consts[753]))
	v190 = *(*int32)(unsafe.Add(mBase, _consts[754]))
	if v188+(v190+v182) < v186 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	goto L34
L34:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+16))
	F_LruDelete(m, v214)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L31
L36:
	;
	return int32(0)
L37:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _consts[751]))
	if v220 <= int32(0) {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[752]))
	v226 = *(*int32)(unsafe.Add(mBase, _consts[753]))
	v228 = *(*int32)(unsafe.Add(mBase, _consts[754]))
	if v224 <= v226+(v228+v220) {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v252
	if v252 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	v260 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	v263 = v260 + v164*int32(48)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+32))
	if v264 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v278 = int32(4359856)
	v280 = *(*int32)(unsafe.Add(mBase, _consts[751]))
	*(*int32)(unsafe.Add(mBase, _consts[751])) = v280 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v178)+40)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v178)+36)) = v251 & int32(-705)
	*(*int32)(unsafe.Add(mBase, uint32(v178)+32)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v178)+24)) = int64(0)
	v291 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v178)+8)) = v291
	*(*uint16)(unsafe.Add(mBase, uint32(v178)+4)) = uint16(v291)
	v296 = *(*int32)(unsafe.Add(mBase, _consts[262]))
	v297 = int32(48)
	v299 = v296 + v164*v297
	*(*int32)(unsafe.Add(mBase, uint32(v299)+16)) = v291
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v296)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v299)+20)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v296)+20)) = v164
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v299)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v296+v305*v297)+16)) = v164
	return v164
L44:
	;
	F_emscripten_builtin_free(m, v264)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v263)+32)) = int32(0)
	goto L46
L45:
	;
	goto L46
L46:
	;
	v268 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v263)+4)) = uint16(v268)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+12)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v164
	F_emscripten_builtin_free(m, v29)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[155])) = v258
	return int32(-1)
L47:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L36
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(12890), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L36
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(478009), int32(1613), int32(275077))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L36
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
	F_errcode(m, int32(8389))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L36
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(12890), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L36
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(478009), int32(1452), int32(418326))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L36
	} else {
		goto L54
	}
L54:
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
	*(*uint8)(unsafe.Add(mBase, _consts[288])) = uint8(v4)
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v50 float64
	_ = v50
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v61 float64
	_ = v61
	v13 = F_palloc0(m, int32(88))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(1580547965224)
		v22 = F_get_baserel_parampathinfo(m, l0, l1, int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = l2
			v25 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v25
			*(*uint16)(unsafe.Add(mBase, uint32(v13)+20)) = uint16(v25)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v22
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v25)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v32
			if v32 == v25 {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
				v39 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v39
				*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v38
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+76)) = uint8(v39)
			} else {
			}
			if l4 != 0 {
				v49 = l4
			} else {
				if v22 != 0 {
					v49 = v22 + int32(8)
				} else {
					v49 = l1 + int32(16)
				}
			}
			v50 = *(*float64)(unsafe.Add(mBase, uint32(v49)))
			*(*float64)(unsafe.Add(mBase, uint32(v13)+32)) = v50
			v53 = *(*float64)(unsafe.Add(mBase, _consts[499]))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
			v55 = *(*float64)(unsafe.Add(mBase, uint32(v54)+56))
			v57 = *(*float64)(unsafe.Add(mBase, _consts[498]))
			v58 = *(*float64)(unsafe.Add(mBase, uint32(v54)+48))
			v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+40))
			*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v59
			v61 = base.F64_add(v58, v57)
			*(*float64)(unsafe.Add(mBase, uint32(v13)+48)) = v61
			*(*float64)(unsafe.Add(mBase, uint32(v13)+56)) = base.F64_add(v61, base.F64_add(base.F64_mul(v53, v50), base.F64_sub(v55, v58)))
			return v13
		}
	}
}
func F_get_path_all(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v134 int32
	_ = v134
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v15 + int32(16)
	return v134
L4:
	;
	v134 = int32(0)
	goto L3
L5:
	;
	v25 = F_array_contains_nulls(m, v23)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
	goto L4
L8:
	;
	goto L9
L9:
	;
	F_deconstruct_array_builtin(m, v23, int32(25), v15+int32(12), v15+int32(8), v15+int32(4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v41 = F_palloc(m, v38<<(uint(int32(2))%32))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v46 = F_palloc(m, v43<<(uint(int32(2))%32))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if int32(0) < v48 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v55 = int32(0)
	goto L16
L14:
	;
	v106 = v48
	goto L15
L15:
	;
	v114 = F_get_worker(m, v18, v41, v46, v106, l1)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L31
	}
L16:
	;
	v65 = v55 << (uint(int32(2)) % 32)
	v66 = v41 + v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67+v65)))
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
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v70
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v73 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v99 = v55 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v99 < v100 {
		v55 = v99
		goto L16
	} else {
		goto L30
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v79 = F_strtol(m, v77, v15, int32(10))
	mBase = m.M
	goto L23
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v46))) = int32(-2147483648)
	goto L19
L23:
	;
	v80 = int32(-2147483648)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v81 == v82 {
		v89 = v80
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65+v46))) = v89
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
	v87 = *(*int32)(unsafe.Add(mBase, _consts[155]))
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
	if v114 != 0 {
		v134 = v114
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v116 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v116)
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
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
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
	return v239
L7:
	;
	v220 = int32(0)
	v221 = F_errsave_start(m, l8)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L28
	} else {
		goto L51
	}
L8:
	;
	v125 = int32(0)
	v127 = v113
	v130 = l3
	v136 = v125
	goto L26
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v108
	v113 = v108
	v124 = int32(1)
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
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v51 != int32(40) {
		v113 = v17
		v124 = v10
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v108 = v17 + int32(1)
	goto L9
L14:
	;
	v63 = v17
	goto L15
L15:
	;
	v67 = v63 + int32(1)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if base.Ui32(v68-int32(9)) < base.Ui32(int32(5)) {
		v63 = v67
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v78 = F_strlen(m, v17)
	mBase = m.M
	v85 = v78 + int32(1)
	goto L21
L17:
	;
	switch v68 - int32(32) {
	case 0:
		v63 = v67
		goto L15
	default:
		goto L18
	case 8:
		v108 = v67
		goto L9
	}
L18:
	;
	goto L16
L19:
	;
	if v97 != v17 {
		v113 = v17
		v124 = v10
		goto L8
	} else {
		goto L25
	}
L20:
	;
	goto L19
L21:
	;
	v87 = int32(0)
	if v85 == v87 {
		v97 = v87
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v97 = v92
	goto L20
L23:
	;
	v91 = v85 - int32(1)
	v92 = v17 + v91
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v93 != int32(40) {
		v85 = v91
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v108 = v67
	goto L9
L26:
	;
	v143 = F_pair_decode(m, v127, v130, v130+int32(8), v15+int32(12), l6, l7, l8)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v124 != 0 {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	return int32(0)
L29:
	;
	if v143 == int32(0) {
		v239 = v125
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v150 == int32(44) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v154 = v149 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v154
	v156 = v154
	goto L33
L32:
	;
	v156 = v149
	goto L33
L33:
	;
	v160 = v136 + int32(1)
	if v160 != l2 {
		v127 = v156
		v130 = v130 + int32(16)
		v136 = v160
		goto L26
	} else {
		goto L34
	}
L34:
	;
	goto L27
L35:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v162 != int32(41) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v192 = v156
	goto L37
L37:
	;
	if l5 != 0 {
		goto L47
	} else {
		goto L48
	}
L38:
	;
	if v162 != int32(93) {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v170 = v156
	goto L43
L41:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v167 != int32(1) {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v183 = v170 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v183
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if base.Ui32(v185-int32(9)) < base.Ui32(int32(5)) {
		v170 = v183
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v192 = v183
	goto L37
L45:
	;
	if v185 == int32(32) {
		v170 = v183
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v192
	v239 = int32(1)
	goto L6
L48:
	;
	goto L49
L49:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v206 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	v239 = int32(1)
	goto L6
L51:
	;
	if v221 == int32(0) {
		v239 = v220
		goto L6
	} else {
		goto L52
	}
L52:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L28
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l6
	F_errmsg(m, int32(685107), v15)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L28
	} else {
		goto L54
	}
L54:
	;
	F_errsave_finish(m, l8, int32(472338), int32(336), int32(395567))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L28
	} else {
		goto L55
	}
L55:
	;
	v239 = v220
	goto L6
}
func F_path_isclosed(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		return base.B2i32(v7 != int32(0))
	}
}
func F_path_isopen(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		return base.B2i32(v7 == int32(0))
	}
}
func F_path_recv(m *base.Module, l0 int32) int32 {
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v59 = m.ExcPending
	if v59 != 0 {
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
	return v24
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
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(331434), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(472338), int32(1502), int32(34291))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v20 = (l4 - l1) << (uint(int32(2)) % 32)
	v21 = F_palloc0(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = l1 + int32(1)
	if l4 <= v24 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v21+v20-int32(4))))
	if v216 == int32(16) {
		goto L44
	} else {
		goto L45
	}
L4:
	;
	v35 = v24
	goto L5
L5:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v35))))
	if v40 != 0 {
		goto L3
	} else {
		goto L7
	}
L6:
	;
	goto L3
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2+v35<<(uint(int32(2))%32))))
	v45 = F_text_to_cstring(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[155])) = int32(0)
	v53 = F_strtol(m, v45, v16+int32(4), int32(10))
	mBase = m.M
	goto L9
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v45 == v58 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+(v35-l1)<<(uint(int32(2))%32)))) = v193
	v196 = v35 + int32(1)
	if v196 != l4 {
		v35 = v196
		goto L5
	} else {
		goto L43
	}
L11:
	;
	v138 = F_pushJsonbValue(m, l0, int32(4), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L35
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(1)
	if v45&int32(3) == int32(0) {
		v91 = v45
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v60 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[155]))
	if v62 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v124
	v128 = F_pushJsonbValue(m, l0, int32(6), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L33
	}
L17:
	;
	v124 = v116 - v45
	goto L16
L18:
	;
	v95 = v91
	goto L27
L19:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v75 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v124 = int32(0)
	goto L16
L21:
	;
	goto L22
L22:
	;
	v80 = v45
	goto L23
L23:
	;
	v84 = v80 + int32(1)
	if v84&int32(3) == int32(0) {
		v91 = v84
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v116 = v84
	goto L17
L25:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v89 != 0 {
		v80 = v84
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v104 = int32(-2139062144)
	if (int32(16843008)-v101|v101)&v104 == v104 {
		v95 = v95 + int32(4)
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v110 = v95
	goto L30
L29:
	;
	goto L28
L30:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v114 != 0 {
		v110 = v110 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v116 = v110
	goto L17
L32:
	;
	goto L31
L33:
	;
	v133 = F_pushJsonbValue(m, l0, int32(1), v16+int32(8))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v193 = int32(17)
	goto L10
L35:
	;
	v140 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v140
	if v140 < v53 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v151 = v53
	goto L39
L37:
	;
	goto L38
L38:
	;
	v193 = int32(16)
	goto L10
L39:
	;
	v160 = F_pushJsonbValue(m, l0, int32(3), v16+int32(28))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	goto L38
L41:
	;
	v162 = int32(1)
	if base.Ui32(v162) < base.Ui32(v151) {
		v151 = v151 - v162
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L6
L44:
	;
	v219 = int32(3)
	goto L46
L45:
	;
	v219 = int32(2)
	goto L46
L46:
	;
	v220 = F_pushJsonbValue(m, l0, v219, l5)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v223 = l4 - int32(1)
	if v223 <= l1 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	m.G0 = v16 + int32(48)
	return
L49:
	;
	v232 = v223
	goto L50
L50:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v232))))
	if v239 != 0 {
		goto L48
	} else {
		goto L52
	}
L51:
	;
	goto L48
L52:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v21+(v232-l1)<<(uint(int32(2))%32))))
	if v246 == int32(17) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v249 = int32(7)
	goto L55
L54:
	;
	v249 = int32(5)
	goto L55
L55:
	;
	v251 = F_pushJsonbValue(m, l0, v249, int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v254 = v232 - int32(1)
	if l1 < v254 {
		v232 = v254
		goto L50
	} else {
		goto L57
	}
L57:
	;
	goto L51
}
