package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TargetPrivilegesCheck(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v13 = F_pg_class_aclcheck(m, v10, v12, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+119)))
			switch v16 - int32(73) {
			case 0, 32:
				v28 = int32(20)
			default:
				v26 = int32(41)
				v28 = v26
			case 10:
				v28 = int32(37)
			case 29:
				v26 = int32(18)
				v28 = v26
			case 36:
				v28 = int32(23)
			case 45:
				v28 = int32(51)
			}
			v29 = F_get_rel_name(m, v10)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				F_aclcheck_error(m, v13, v28, v29)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = int32(0)
					v35 = F_check_enable_rls(m, v10, v33, v33)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						if v35 == int32(2) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, _consts[239]))
									v49 = F_GetUserNameFromId(m, v47, int32(1))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
										*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v51 + int32(4)
										F_errmsg(m, int32(711606), v8)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											F_errfinish(m, int32(494837), int32(2380), int32(317890))
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
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
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		} else {
			v33 = int32(0)
			v35 = F_check_enable_rls(m, v10, v33, v33)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				if v35 == int32(2) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, _consts[239]))
							v49 = F_GetUserNameFromId(m, v47, int32(1))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v51 + int32(4)
								F_errmsg(m, int32(711606), v8)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_errfinish(m, int32(494837), int32(2380), int32(317890))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
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
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	}
}
func F_addTargetToSortList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
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
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	v16 = m.G0
	v18 = v16 - int32(80)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v21 = F_exprType(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v21 == int32(705) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = int32(25)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v31 = int32(-1)
	v35 = F_coerce_type(m, l0, v28, int32(705), v27, v31, int32(0), int32(2), v31)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v38 = v21
	goto L5
L5:
	;
	v40 = v18 + int32(48)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	if v41 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v35
	v38 = v27
	goto L5
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v45 = F_exprLocation(m, v44)
	mBase = m.M
	v46 = v45
	goto L9
L8:
	;
	v46 = v41
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v40
	v52 = int32(4482344)
	v53 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v53
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v18 + int32(56)
	goto L10
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	switch v59 {
	case 0, 1:
		goto L13
	case 2:
		goto L16
	case 3:
		goto L15
	default:
		goto L14
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L84
	}
L12:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(48))+8))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v120
	goto L26
L13:
	;
	v102 = int32(1)
	v104 = int32(0)
	F_get_sort_group_operators(m, v38, v102, v102, v104, v18+int32(76), v18+int32(72), v104, v18+int32(71))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L25
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v75 = F_compatible_oper_opid(m, v74, v38, v38)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	v60 = int32(0)
	v61 = int32(1)
	F_get_sort_group_operators(m, v38, v60, v61, v61, v60, v18+int32(72), v18+int32(76), v18+int32(71))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+70)) = uint8(v72)
	goto L12
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v75
	v80 = F_get_equality_op_for_ordering_op(m, v75, v18+int32(70))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v80
	if v80 == int32(0) {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	v85 = F_op_hashjoinable(m, v80, v38)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+71)) = uint8(v85)
	goto L12
L22:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v92
	F_errmsg_internal(m, int32(481494), v18)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(498071), int32(3472), int32(75913))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+70)) = uint8(v114)
	goto L12
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v122 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	m.G0 = v18 + int32(80)
	return v370
L28:
	;
	v182 = F_palloc0(m, int32(20))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L42
	}
L29:
	;
	if l2 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v127 <= int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
	v132 = v127
	v137 = int32(0)
	goto L32
L32:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147+v137<<(uint(int32(2))%32))))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v122 == v152 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L28
L34:
	;
	if v130 == int32(0) {
		v370 = l2
		goto L27
	} else {
		goto L37
	}
L35:
	;
	v162 = v132
	goto L36
L36:
	;
	v164 = v137 + int32(1)
	if v164 < v162 {
		v132 = v162
		v137 = v164
		goto L32
	} else {
		goto L41
	}
L37:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	if v130 == v156 {
		v370 = l2
		goto L27
	} else {
		goto L38
	}
L38:
	;
	v158 = F_get_commutator(m, v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v158 == v130 {
		v370 = l2
		goto L27
	} else {
		goto L40
	}
L40:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v162 = v161
	goto L36
L41:
	;
	goto L33
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = int32(106)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v186 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if l3 == int32(0) {
		v319 = int32(1)
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v326 = v186
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v326
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = v337
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v182)+12)) = v339
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+71)))
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+18)) = uint8(v341)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+70)))
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+16)) = uint8(v343)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	switch v345 {
	case 0:
		v364 = v343
		goto L76
	case 1:
		goto L77
	case 2:
		goto L79
	default:
		goto L78
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v319
	v326 = v319
	goto L45
L47:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v193 <= int32(0) {
		v319 = int32(1)
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v197 = v193 & int32(3)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v199 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v193) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v206 = v199
	v211 = v199
	v216 = int32(0)
	goto L52
L50:
	;
	v245 = v199
	v250 = v199
	goto L51
L51:
	;
	if v197 != 0 {
		goto L67
	} else {
		goto L68
	}
L52:
	;
	v223 = v198 + v206<<(uint(int32(2))%32)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+16))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+16))
	if base.Ui32(v211) < base.Ui32(v231) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v245 = v241
	v250 = v239
	goto L51
L54:
	;
	v233 = v231
	goto L56
L55:
	;
	v233 = v211
	goto L56
L56:
	;
	if base.Ui32(v233) < base.Ui32(v229) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v235 = v229
	goto L59
L58:
	;
	v235 = v233
	goto L59
L59:
	;
	if base.Ui32(v235) < base.Ui32(v227) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v237 = v227
	goto L62
L61:
	;
	v237 = v235
	goto L62
L62:
	;
	if base.Ui32(v237) < base.Ui32(v225) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v239 = v225
	goto L65
L64:
	;
	v239 = v237
	goto L65
L65:
	;
	v240 = int32(4)
	v241 = v206 + v240
	v243 = v216 + v240
	if v243 != v193&int32(2147483644) {
		v206 = v241
		v211 = v239
		v216 = v243
		goto L52
	} else {
		goto L66
	}
L66:
	;
	goto L53
L67:
	;
	v260 = v245
	v265 = v250
	v269 = v199
	goto L70
L68:
	;
	v292 = v250
	goto L69
L69:
	;
	v319 = v292 + int32(1)
	goto L46
L70:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v198+v260<<(uint(int32(2))%32))))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	if base.Ui32(v265) < base.Ui32(v279) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v292 = v281
	goto L69
L72:
	;
	v281 = v279
	goto L74
L73:
	;
	v281 = v265
	goto L74
L74:
	;
	v282 = int32(1)
	v285 = v269 + v282
	if v285 != v197 {
		v260 = v260 + v282
		v265 = v281
		v269 = v285
		goto L70
	} else {
		goto L75
	}
L75:
	;
	goto L71
L76:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+17)) = uint8(v364)
	v366 = F_lappend(m, l2, v182)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L83
	}
L77:
	;
	v364 = int32(1)
	goto L76
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v364 = int32(0)
	goto L76
L80:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v351
	F_errmsg_internal(m, int32(481256), v18+int32(16))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(498071), int32(3508), int32(75913))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
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
	v370 = v366
	goto L27
L84:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+12))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v395+v396<<(uint(int32(2))%32)-int32(4))))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v403
	F_errmsg(m, int32(208997), v18+int32(32))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errhint(m, int32(585397), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(498071), int32(3464), int32(75913))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_rewriteTargetListIU(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
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
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int64
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int64
	_ = v278
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v304 int32
	_ = v304
	var v306 int64
	_ = v306
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int64
	_ = v319
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v346 int32
	_ = v346
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v450 int32
	_ = v450
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	v8 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(176)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+120)))
	v33 = F_palloc0(m, v30<<(uint(int32(2))%32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l0 == int32(0) {
		v450 = v8
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if int32(0) < v30 {
		goto L96
	} else {
		goto L97
	}
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v39 <= int32(0) {
		v450 = v8
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v56 = v8
	v60 = v30 + int32(1)
	v62 = v8
	goto L7
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L89
	}
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v56<<(uint(int32(2))%32))))
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v72)+8)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+26)))
	if v74 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L85
	}
L9:
	;
	goto L8
L10:
	;
	v382 = v56 + int32(1)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v382 < v383 {
		v56 = v382
		v60 = v373
		v62 = v375
		goto L7
	} else {
		goto L84
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v346
	v373 = v60
	v375 = v62
	goto L10
L12:
	;
	switch v175 - int32(14) {
	case 0:
		goto L49
	default:
		goto L9
	case 12:
		v187 = int32(4)
		goto L48
	}
L13:
	;
	if v168 == int32(0) {
		goto L9
	} else {
		goto L47
	}
L14:
	;
	switch v150 - int32(14) {
	case 0:
		goto L45
	default:
		v167 = v151
		v168 = v152
		v169 = v154
		v170 = v106
		goto L13
	case 12:
		v162 = int32(4)
		goto L44
	}
L15:
	;
	v150 = v107
	v151 = v100
	v152 = v99
	v154 = int32(0)
	goto L14
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L41
	}
L17:
	;
	if v73 <= int32(0) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v73 != v60 {
		goto L36
	} else {
		goto L37
	}
L20:
	;
	if v30 < v73 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v87 = v80 + v81<<(uint(int32(4))%32) + v73*int32(100)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+11)))
	if v88 != 0 {
		v373 = v60
		v375 = v62
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v93 = v33 + (v73-int32(1))<<(uint(int32(2))%32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 == int32(0) {
		v346 = v72
		goto L11
	} else {
		goto L23
	}
L23:
	;
	v98 = v87 - int32(80)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v100 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v103 = int32(0)
	v167 = v103
	v168 = v99
	v169 = v103
	v170 = v103
	goto L13
L25:
	;
	goto L26
L26:
	;
	v106 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v107 != int32(55) {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	if v99 == int32(0) {
		goto L15
	} else {
		goto L28
	}
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v113 != int32(55) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v175 = v113
	v176 = v100
	v177 = v99
	v178 = int32(0)
	v179 = v106
	goto L12
L30:
	;
	goto L31
L31:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	if v116 != v117 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v120 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v167 = int32(0)
	v168 = v119
	v169 = v100
	v170 = v106
	goto L13
L34:
	;
	goto L35
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v150 = v124
	v151 = v120
	v152 = v119
	v154 = v100
	goto L14
L36:
	;
	v126 = F_flatCopyTargetEntry(m, v72)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v129 = v72
	goto L38
L38:
	;
	v132 = F_lappend(m, v62, v129)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v126)+8)) = uint16(v60)
	v129 = v126
	goto L38
L40:
	;
	v373 = v60 + int32(1)
	v375 = v132
	goto L10
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+128)) = v73
	F_errmsg_internal(m, int32(73248), v27+int32(128))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(494780), int32(815), int32(515987))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
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
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162+v151)))
	v167 = v151
	v168 = v152
	v169 = v154
	v170 = v164
	goto L13
L45:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v151)+36))
	if v158 == int32(0) {
		v167 = v151
		v168 = v152
		v169 = v154
		v170 = v106
		goto L13
	} else {
		goto L46
	}
L46:
	;
	v162 = int32(32)
	goto L44
L47:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v175 = v173
	v176 = v167
	v177 = v168
	v178 = v169
	v179 = v170
	goto L12
L48:
	;
	if v179 == int32(0) {
		goto L9
	} else {
		goto L51
	}
L49:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v177)+36))
	if v183 == int32(0) {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	v187 = int32(32)
	goto L48
L51:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187+v177)))
	if v191 == int32(0) {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	v194 = F_exprType(m, v176)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v196 = F_exprType(m, v177)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v194 != v196 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	v210 = v191
	goto L56
L56:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	switch v224 - int32(14) {
	case 0:
		goto L60
	default:
		goto L58
	case 12:
		v231 = int32(4)
		goto L59
	}
L57:
	;
	v236 = F_equal(m, v210, v179)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L63
	}
L58:
	;
	goto L57
L59:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v210+v231)))
	if v233 != 0 {
		v210 = v233
		goto L56
	} else {
		goto L62
	}
L60:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v210)+36))
	if v227 == int32(0) {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v231 = int32(32)
	goto L59
L62:
	;
	goto L58
L63:
	;
	if v236 == int32(0) {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	switch v240 - int32(14) {
	case 0:
		goto L68
	default:
		goto L67
	case 12:
		goto L69
	}
L65:
	;
	if v178 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L66:
	;
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v176)))
	*(*int64)(unsafe.Add(mBase, uint32(v244))) = v302
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v176)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+16)) = v304
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v176)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v244)+8)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v244)+4)) = v177
	v309 = v244
	goto L65
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L75
	}
L68:
	;
	v272 = F_palloc0(m, int32(40))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L74
	}
L69:
	;
	v244 = F_palloc0(m, int32(20))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v246 = int32(26)
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	if v248 != v246 {
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v177)))
	*(*int64)(unsafe.Add(mBase, uint32(v244))) = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+16)) = v253
	v255 = int32(8)
	v256 = v244 + v255
	v258 = v177 + v255
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v258)))
	*(*int64)(unsafe.Add(mBase, uint32(v256))) = v259
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
	v263 = F_list_concat_copy(m, v261, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v263
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v268 = F_list_concat_copy(m, v266, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244)+12)) = v268
	v309 = v244
	goto L65
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = int32(14)
	v277 = v272 + int32(32)
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v176)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v277))) = v278
	v280 = *(*int64)(unsafe.Add(mBase, uint32(v176)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+8)) = v280
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v176)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+16)) = v282
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v176)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v272)+24)) = v284
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v176)))
	*(*int64)(unsafe.Add(mBase, uint32(v272))) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v177
	v309 = v272
	goto L65
L75:
	;
	F_errmsg_internal(m, int32(281171), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(494780), int32(1178), int32(382725))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
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
	v329 = F_flatCopyTargetEntry(m, v72)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L83
	}
L79:
	;
	v328 = v309
	goto L78
L80:
	;
	goto L81
L81:
	;
	v315 = F_palloc0(m, int32(28))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315))) = int32(55)
	v319 = *(*int64)(unsafe.Add(mBase, uint32(v178)))
	*(*int64)(unsafe.Add(mBase, uint32(v315))) = v319
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v178)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v315)+8)) = v321
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v178)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v315)+16)) = v323
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v178)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+24)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(v315)+4)) = v309
	v328 = v315
	goto L78
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+4)) = v328
	v346 = v329
	goto L11
L84:
	;
	v450 = v375
	goto L3
L85:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+144)) = v98 + int32(4)
	F_errmsg(m, int32(695782), v27+int32(144))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(494780), int32(1122), int32(382725))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+160)) = v98 + int32(4)
	F_errmsg(m, int32(695782), v27+int32(160))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(494780), int32(1140), int32(382725))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L185
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L180
	}
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L175
	}
L96:
	;
	v459 = base.B2i32(l1 != int32(3))
	v460 = int32(1)
	v472 = v460
	v473 = int32(0)
	v483 = v8
	goto L99
L97:
	;
	v719 = v8
	goto L98
L98:
	;
	F_pfree(m, v33)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L173
	}
L99:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	v495 = v472*int32(100) + (v490 + v491<<(uint(int32(4))%32))
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495)+11)))
	if v496 != 0 {
		v691 = v473
		v696 = v483
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v719 = v696
	goto L98
L101:
	;
	if v472 != v30 {
		v472 = v472 + int32(1)
		v473 = v691
		v483 = v696
		goto L99
	} else {
		goto L172
	}
L102:
	;
	v498 = v495 - int32(80)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v472<<(uint(int32(2))%32)+v33-int32(4))))
	if l1 != int32(3) {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+90)))
	if v652 != 0 {
		goto L153
	} else {
		goto L154
	}
L104:
	;
	if l1 != int32(2) {
		v647 = v473
		v648 = v516
		goto L103
	} else {
		goto L150
	}
L105:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+89)))
	v588 = base.B2i32(l2 == int32(1))&base.B2i32(v584 == int32(100)) | v580
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+90)))
	v590 = int32(0)
	v592 = v588 | base.B2i32(v589 != v590)
	if v589 == v590 {
		v609 = v579
		v610 = v592
		goto L136
	} else {
		goto L137
	}
L106:
	;
	v507 = int32(0)
	if v504 == v507 {
		v516 = v507
		goto L109
	} else {
		goto L110
	}
L107:
	;
	if v504 != 0 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v579 = v473
	v580 = int32(1)
	v581 = int32(0)
	goto L105
L109:
	;
	if l1 != int32(3) {
		goto L104
	} else {
		goto L112
	}
L110:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
	if v510 == int32(0) {
		v516 = v507
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v510)))
	v516 = base.B2i32(v513 == int32(57))
	goto L109
L112:
	;
	v518 = int32(0)
	if l4 == v518 {
		v531 = v518
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+89)))
	if (base.B2i32(v532 != int32(97))|v516)&int32(1) != 0 {
		v579 = v473
		v580 = v516
		v581 = v531
		goto L105
	} else {
		goto L118
	}
L114:
	;
	if v504 == int32(0) {
		v531 = v518
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v504)+4))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	if v524 != int32(6) {
		v531 = v518
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	if v527 != l5 {
		v531 = v518
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v529 = int32(*(*int16)(unsafe.Add(mBase, uint32(v523)+8)))
	v531 = v529
	goto L113
L118:
	;
	v538 = int32(1)
	switch l2 - v460 {
	case 0:
		v579 = v473
		v580 = v538
		v581 = v531
		goto L105
	case 1:
		goto L119
	default:
		goto L120
	}
L119:
	;
	v579 = v473
	v580 = int32(0)
	v581 = v531
	goto L105
L120:
	;
	if v531 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	if v473 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L130
	}
L124:
	;
	v541 = F_findDefaultOnlyColumns(m, l4)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	v543 = v473
	goto L126
L126:
	;
	v544 = F_bms_is_member(m, v531, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L128
	}
L127:
	;
	v543 = v541
	goto L126
L128:
	;
	if v544 != 0 {
		v579 = v543
		v580 = v538
		v581 = v531
		goto L105
	} else {
		goto L129
	}
L129:
	;
	goto L123
L130:
	;
	F_errcode(m, int32(156008580))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v555 = v498 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v555
	F_errmsg(m, int32(694613), v27+int32(48))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v555
	F_errdetail(m, int32(641374), v27+int32(32))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errhint(m, int32(627934), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(494780), int32(915), int32(515987))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	if v581 == int32(0) {
		v647 = v609
		v648 = v610
		goto L103
	} else {
		goto L146
	}
L137:
	;
	if v588&int32(1) != 0 {
		v609 = v579
		v610 = v592
		goto L136
	} else {
		goto L138
	}
L138:
	;
	if v581 == int32(0) {
		goto L95
	} else {
		goto L139
	}
L139:
	;
	if v579 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v601 = F_findDefaultOnlyColumns(m, l4)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	v603 = v579
	goto L142
L142:
	;
	v605 = F_bms_is_member(m, v581, v603)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L144
	}
L143:
	;
	v603 = v601
	goto L142
L144:
	;
	if v605 == int32(0) {
		goto L95
	} else {
		goto L145
	}
L145:
	;
	v609 = v603
	v610 = int32(1)
	goto L136
L146:
	;
	if l6 == int32(0) {
		v647 = v609
		v648 = v610
		goto L103
	} else {
		goto L147
	}
L147:
	;
	if v610&int32(1) == int32(0) {
		v647 = v609
		v648 = v610
		goto L103
	} else {
		goto L148
	}
L148:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v620 = F_bms_add_member(m, v619, v581)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v620
	v647 = v609
	v648 = int32(1)
	goto L103
L150:
	;
	v626 = int32(0)
	v627 = base.B2i32(v504 == v626)
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+89)))
	if (v627|base.B2i32(v628 != int32(97))|v516)&int32(1) == v626 {
		goto L94
	} else {
		goto L151
	}
L151:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+90)))
	v638 = int32(0)
	if (base.B2i32(v637 == v638)|v627|v516)&int32(1) == v638 {
		goto L93
	} else {
		goto L152
	}
L152:
	;
	v647 = v473
	v648 = v516
	goto L103
L153:
	;
	v653 = int32(0)
	goto L155
L154:
	;
	v653 = v504
	goto L155
L155:
	;
	if v652 != 0 {
		v684 = v653
		goto L156
	} else {
		goto L157
	}
L156:
	;
	if v684 == int32(0) {
		v691 = v647
		v696 = v483
		goto L101
	} else {
		goto L170
	}
L157:
	;
	if v648&int32(1) == int32(0) {
		v684 = v653
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v659 = F_build_column_default(m, l3, v472)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	if v659 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v661 = v504
	goto L162
L161:
	;
	v661 = int32(0)
	goto L162
L162:
	;
	if l1 == int32(3) {
		v671 = v659
		v672 = v661
		goto L163
	} else {
		goto L164
	}
L163:
	;
	if v671 == int32(0) {
		v684 = v672
		goto L156
	} else {
		goto L167
	}
L164:
	;
	if v659 != 0 {
		v671 = v659
		v672 = v661
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v498)+68))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v498)+76))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v498)+96))
	v667 = int32(*(*int16)(unsafe.Add(mBase, uint32(v498)+72)))
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+82)))
	v669 = F_coerce_null_to_domain(m, v664, v665, v666, v667, v668)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v671 = v669
	v672 = v504
	goto L163
L167:
	;
	v678 = F_pstrdup(m, v498+int32(4))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v681 = F_makeTargetEntry(m, v671, base.I32_extend16_s(v472), v678, int32(0))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v684 = v681
	goto L156
L170:
	;
	v687 = F_lappend(m, v483, v684)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v691 = v647
	v696 = v687
	goto L101
L172:
	;
	goto L100
L173:
	;
	v726 = F_list_concat(m, v719, v450)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	m.G0 = v27 + int32(176)
	return v726
L175:
	;
	F_errcode(m, int32(156008580))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v742 = v498 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v742
	F_errmsg(m, int32(694613), v27+int32(16))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v742
	F_errdetail(m, int32(603538), v27)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(494780), int32(953), int32(515987))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L180:
	;
	F_errcode(m, int32(156008580))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v766 = v498 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v766
	F_errmsg(m, int32(519788), v27+int32(112))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v766
	F_errdetail(m, int32(641374), v27+int32(96))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(494780), int32(980), int32(515987))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L185:
	;
	F_errcode(m, int32(156008580))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v792 = v498 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v792
	F_errmsg(m, int32(519788), v27+int32(80))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v792
	F_errdetail(m, int32(603538), v27-int32(-64))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(494780), int32(988), int32(515987))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
