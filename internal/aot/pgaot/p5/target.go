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
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_TargetPrivilegesCheck[0]))
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
				v26 = int32(20)
				v28 = v26
			default:
				v26 = int32(41)
				v28 = v26
			case 10:
				v28 = int32(37)
			case 29:
				v28 = int32(18)
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
									v47 = *(*int32)(unsafe.Add(mBase, _c_F_TargetPrivilegesCheck[0]))
									v49 = F_GetUserNameFromId(m, v47, int32(1))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
										*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v51 + int32(4)
										F_errmsg(m, int32(_a_F_TargetPrivilegesCheck_0), v8)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_TargetPrivilegesCheck_1), int32(2380), int32(_a_F_TargetPrivilegesCheck_2))
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
							v47 = *(*int32)(unsafe.Add(mBase, _c_F_TargetPrivilegesCheck[0]))
							v49 = F_GetUserNameFromId(m, v47, int32(1))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v51 + int32(4)
								F_errmsg(m, int32(_a_F_TargetPrivilegesCheck_0), v8)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_TargetPrivilegesCheck_1), int32(2380), int32(_a_F_TargetPrivilegesCheck_2))
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
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
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
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
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
	v52 = int32(_a_F_addTargetToSortList_0)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_addTargetToSortList[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v53
	*(*int32)(unsafe.Add(mBase, _c_F_addTargetToSortList[0])) = v18 + int32(56)
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
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L82
	}
L12:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(48))+8))
	*(*int32)(unsafe.Add(mBase, _c_F_addTargetToSortList[0])) = v120
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
	F_errmsg_internal(m, int32(_a_F_addTargetToSortList_1), v18)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_addTargetToSortList_2), int32(3472), int32(_a_F_addTargetToSortList_3))
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
	v123 = int32(0)
	if base.B2i32(v122 == v123)|base.B2i32(l2 == v123) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	m.G0 = v18 + int32(80)
	return v373
L28:
	;
	v183 = F_palloc0(m, int32(20))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L41
	}
L29:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v128 <= int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
	v133 = v128
	v139 = int32(0)
	goto L31
L31:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148+v139<<(uint(int32(2))%32))))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v122 == v153 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L28
L33:
	;
	if v131 == int32(0) {
		v373 = l2
		goto L27
	} else {
		goto L36
	}
L34:
	;
	v163 = v133
	goto L35
L35:
	;
	v165 = v139 + int32(1)
	if v165 < v163 {
		v133 = v163
		v139 = v165
		goto L31
	} else {
		goto L40
	}
L36:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	if v131 == v157 {
		v373 = l2
		goto L27
	} else {
		goto L37
	}
L37:
	;
	v159 = F_get_commutator(m, v157)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v159 == v131 {
		v373 = l2
		goto L27
	} else {
		goto L39
	}
L39:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v163 = v162
	goto L35
L40:
	;
	goto L32
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = int32(106)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v187 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if l3 == int32(0) {
		v322 = int32(1)
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v330 = v187
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v330
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+8)) = v340
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v18)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+12)) = v342
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+71)))
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+18)) = uint8(v344)
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+70)))
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+16)) = uint8(v346)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	switch v348 {
	case 0:
		v367 = v346
		goto L74
	case 1:
		goto L75
	case 2:
		goto L77
	default:
		goto L76
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v322
	v330 = v322
	goto L44
L46:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v194 <= int32(0) {
		v322 = int32(1)
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v198 = v194 & int32(3)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v200 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v194) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v322 = v296 + int32(1)
	goto L45
L49:
	;
	v207 = v200
	v213 = v200
	v221 = int32(0)
	goto L52
L50:
	;
	v248 = v200
	v254 = v200
	goto L51
L51:
	;
	v263 = v248
	v269 = v254
	v272 = v200
	goto L68
L52:
	;
	v224 = v199 + v207<<(uint(int32(2))%32)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+12))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+16))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+16))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+16))
	if base.Ui32(v213) < base.Ui32(v232) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v198 == int32(0) {
		v296 = v240
		goto L48
	} else {
		goto L67
	}
L54:
	;
	v234 = v232
	goto L56
L55:
	;
	v234 = v213
	goto L56
L56:
	;
	if base.Ui32(v234) < base.Ui32(v230) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v236 = v230
	goto L59
L58:
	;
	v236 = v234
	goto L59
L59:
	;
	if base.Ui32(v236) < base.Ui32(v228) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v238 = v228
	goto L62
L61:
	;
	v238 = v236
	goto L62
L62:
	;
	if base.Ui32(v238) < base.Ui32(v226) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v240 = v226
	goto L65
L64:
	;
	v240 = v238
	goto L65
L65:
	;
	v241 = int32(4)
	v242 = v207 + v241
	v244 = v221 + v241
	if v244 != v194&int32(2147483644) {
		v207 = v242
		v213 = v240
		v221 = v244
		goto L52
	} else {
		goto L66
	}
L66:
	;
	goto L53
L67:
	;
	v248 = v242
	v254 = v240
	goto L51
L68:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v199+v263<<(uint(int32(2))%32))))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+16))
	if base.Ui32(v269) < base.Ui32(v282) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v296 = v284
	goto L48
L70:
	;
	v284 = v282
	goto L72
L71:
	;
	v284 = v269
	goto L72
L72:
	;
	v285 = int32(1)
	v288 = v272 + v285
	if v288 != v198 {
		v263 = v263 + v285
		v269 = v284
		v272 = v288
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+17)) = uint8(v367)
	v369 = F_lappend(m, l2, v183)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L81
	}
L75:
	;
	v367 = int32(1)
	goto L74
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	v367 = int32(0)
	goto L74
L78:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v354
	F_errmsg_internal(m, int32(_a_F_addTargetToSortList_4), v18+int32(16))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_addTargetToSortList_2), int32(3508), int32(_a_F_addTargetToSortList_3))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	v373 = v369
	goto L27
L82:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v398+v399<<(uint(int32(2))%32)-int32(4))))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v406
	F_errmsg(m, int32(_a_F_addTargetToSortList_5), v18+int32(32))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errhint(m, int32(_a_F_addTargetToSortList_6), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_addTargetToSortList_2), int32(3464), int32(_a_F_addTargetToSortList_3))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
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
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int64
	_ = v265
	var v267 int64
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v288 int64
	_ = v288
	var v290 int64
	_ = v290
	var v292 int64
	_ = v292
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int64
	_ = v308
	var v310 int32
	_ = v310
	var v312 int64
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v327 int64
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v351 int32
	_ = v351
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
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
		v457 = v8
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if int32(0) < v30 {
		goto L112
	} else {
		goto L113
	}
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v39 <= int32(0) {
		v457 = v8
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v56 = v8
	v62 = v30 + int32(1)
	v66 = v8
	goto L7
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L105
	}
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v56<<(uint(int32(2))%32))))
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v72)+8)))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+26)))
	if v74 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L101
	}
L9:
	;
	goto L8
L10:
	;
	v386 = v56 + int32(1)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v386 < v387 {
		v56 = v386
		v62 = v379
		v66 = v383
		goto L7
	} else {
		goto L100
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v351
	v379 = v62
	v383 = v66
	goto L10
L12:
	;
	v188 = v181 - int32(14)
	if v188 != 0 {
		goto L53
	} else {
		goto L54
	}
L13:
	;
	if v174 == int32(0) {
		goto L9
	} else {
		goto L51
	}
L14:
	;
	v161 = v154 - int32(14)
	if v161 != 0 {
		goto L44
	} else {
		goto L45
	}
L15:
	;
	v154 = v127
	v155 = v118
	v156 = v117
	v158 = int32(0)
	goto L14
L16:
	;
	v78 = int32(0)
	if base.B2i32(v73 <= v30)&base.B2i32(v78 < v73) == v78 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if v73 != v62 {
		goto L38
	} else {
		goto L39
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v105 = v98 + v99<<(uint(int32(4))%32) + v73*int32(100)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+11)))
	if v106 != 0 {
		v379 = v62
		v383 = v66
		goto L10
	} else {
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+128)) = v73
	F_errmsg_internal(m, int32(_a_F_rewriteTargetListIU_0), v27+int32(128))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(815), int32(_a_F_rewriteTargetListIU_2))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
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
	v111 = v33 + (v73-int32(1))<<(uint(int32(2))%32)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v112 == int32(0) {
		v351 = v72
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v116 = v105 - int32(80)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v118 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v121 = int32(0)
	v173 = v121
	v174 = v117
	v175 = v121
	v176 = v121
	goto L13
L28:
	;
	goto L29
L29:
	;
	v124 = int32(0)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if base.B2i32(v117 == v124)|base.B2i32(v127 != int32(55)) != 0 {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v132 != int32(55) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v181 = v132
	v182 = v118
	v183 = v117
	v184 = int32(0)
	v185 = v124
	goto L12
L32:
	;
	goto L33
L33:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	if v135 != v136 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v139 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v173 = int32(0)
	v174 = v138
	v175 = v118
	v176 = v124
	goto L13
L36:
	;
	goto L37
L37:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v154 = v143
	v155 = v139
	v156 = v138
	v158 = v118
	goto L14
L38:
	;
	v145 = F_flatCopyTargetEntry(m, v72)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v148 = v72
	goto L40
L40:
	;
	v151 = F_lappend(m, v66, v148)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v145)+8)) = uint16(v62)
	v148 = v145
	goto L40
L42:
	;
	v379 = v62 + int32(1)
	v383 = v151
	goto L10
L43:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v168+v155)))
	v173 = v155
	v174 = v156
	v175 = v158
	v176 = v170
	goto L13
L44:
	;
	if v161 == int32(12) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v155)+36))
	if v164 == int32(0) {
		v173 = v155
		v174 = v156
		v175 = v158
		v176 = v124
		goto L13
	} else {
		goto L50
	}
L47:
	;
	v168 = int32(4)
	goto L43
L48:
	;
	v173 = v155
	v174 = v156
	v175 = v158
	v176 = v124
	goto L13
L50:
	;
	v168 = int32(32)
	goto L43
L51:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v181 = v179
	v182 = v173
	v183 = v174
	v184 = v175
	v185 = v176
	goto L12
L52:
	;
	if v185 == int32(0) {
		goto L9
	} else {
		goto L60
	}
L53:
	;
	if v188 == int32(12) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v183)+36))
	if v191 == int32(0) {
		goto L9
	} else {
		goto L59
	}
L56:
	;
	v195 = int32(4)
	goto L52
L57:
	;
	goto L9
L59:
	;
	v195 = int32(32)
	goto L52
L60:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195+v183)))
	if v199 == int32(0) {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	v202 = F_exprType(m, v182)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v204 = F_exprType(m, v183)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v202 != v204 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v214 = v199
	goto L65
L65:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	v234 = v232 - int32(14)
	if v234 != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v246 = F_equal(m, v214, v185)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L77
	}
L67:
	;
	goto L66
L68:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v241+v214)))
	if v243 != 0 {
		v214 = v243
		goto L65
	} else {
		goto L76
	}
L69:
	;
	if v234 == int32(12) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v214)+36))
	if v237 == int32(0) {
		goto L67
	} else {
		goto L75
	}
L72:
	;
	v241 = int32(4)
	goto L68
L73:
	;
	goto L67
L75:
	;
	v241 = int32(32)
	goto L68
L76:
	;
	goto L67
L77:
	;
	if v246 == int32(0) {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v252 = v250 - int32(14)
	if v252 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	if v184 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L80:
	;
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v182)))
	*(*int64)(unsafe.Add(mBase, uint32(v256))) = v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v182)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+16)) = v310
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v182)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v256)+8)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v256)+4)) = v183
	v315 = v256
	goto L79
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L91
	}
L82:
	;
	if v252 != int32(12) {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v280 = F_palloc0(m, int32(40))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L90
	}
L85:
	;
	v256 = F_palloc0(m, int32(20))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v258 = int32(26)
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v258
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v260 != v258 {
		goto L80
	} else {
		goto L87
	}
L87:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+16)) = v263
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v183)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v256)+8)) = v265
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
	*(*int64)(unsafe.Add(mBase, uint32(v256))) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v182)+8))
	v271 = F_list_concat_copy(m, v269, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+8)) = v271
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	v276 = F_list_concat_copy(m, v274, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+12)) = v276
	v315 = v256
	goto L79
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280))) = int32(14)
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v182)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v280)+32)) = v284
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v182)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v280)+8)) = v286
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v182)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v280)+16)) = v288
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v182)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v280)+24)) = v290
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v182)))
	*(*int64)(unsafe.Add(mBase, uint32(v280))) = v292
	*(*int32)(unsafe.Add(mBase, uint32(v280)+32)) = v183
	v315 = v280
	goto L79
L91:
	;
	F_errmsg_internal(m, int32(_a_F_rewriteTargetListIU_3), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(1178), int32(_a_F_rewriteTargetListIU_4))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	v333 = F_flatCopyTargetEntry(m, v72)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L99
	}
L95:
	;
	v332 = v315
	goto L94
L96:
	;
	goto L97
L97:
	;
	v319 = F_palloc0(m, int32(28))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319))) = int32(55)
	v323 = *(*int64)(unsafe.Add(mBase, uint32(v184)))
	*(*int64)(unsafe.Add(mBase, uint32(v319))) = v323
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v184)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v319)+8)) = v325
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v184)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v319)+16)) = v327
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v184)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v319)+24)) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v319)+4)) = v315
	v332 = v319
	goto L94
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+4)) = v332
	v351 = v333
	goto L11
L100:
	;
	v457 = v383
	goto L3
L101:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+144)) = v116 + int32(4)
	F_errmsg(m, int32(_a_F_rewriteTargetListIU_5), v27+int32(144))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(1122), int32(_a_F_rewriteTargetListIU_4))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+160)) = v116 + int32(4)
	F_errmsg(m, int32(_a_F_rewriteTargetListIU_5), v27+int32(160))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(1140), int32(_a_F_rewriteTargetListIU_4))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L197
	}
L110:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L192
	}
L111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L187
	}
L112:
	;
	v462 = base.B2i32(l1 != int32(3))
	v463 = int32(1)
	v475 = v463
	v477 = int32(0)
	v490 = v8
	goto L115
L113:
	;
	v734 = v8
	goto L114
L114:
	;
	F_pfree(m, v33)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L185
	}
L115:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	v498 = v491 + v492<<(uint(int32(4))%32) + v475*int32(100)
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+11)))
	if v499 != 0 {
		v704 = v477
		v707 = v490
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v734 = v707
	goto L114
L117:
	;
	if v475 != v30 {
		v475 = v475 + int32(1)
		v477 = v704
		v490 = v707
		goto L115
	} else {
		goto L184
	}
L118:
	;
	v501 = v498 - int32(80)
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v475<<(uint(int32(2))%32)+v33-int32(4))))
	if v507|v462 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+90)))
	if v660 != 0 {
		goto L166
	} else {
		goto L167
	}
L120:
	;
	if l1 != int32(2) {
		v655 = v522
		v657 = v477
		goto L119
	} else {
		goto L163
	}
L121:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+90)))
	v590 = int32(0)
	v592 = int32(1)
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+89)))
	v598 = base.B2i32(l2 == v592)&base.B2i32(v594 == int32(100)) | v585
	if base.B2i32(v589 == v590)|v598&v592 == v590 {
		goto L151
	} else {
		goto L152
	}
L122:
	;
	v585 = int32(1)
	v587 = v477
	v588 = int32(0)
	goto L121
L123:
	;
	goto L124
L124:
	;
	v513 = int32(0)
	if v507 == v513 {
		v522 = v513
		goto L125
	} else {
		goto L126
	}
L125:
	;
	if l1 != int32(3) {
		goto L120
	} else {
		goto L128
	}
L126:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	if v516 == int32(0) {
		v522 = v513
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v516)))
	v522 = base.B2i32(v519 == int32(57))
	goto L125
L128:
	;
	v524 = int32(0)
	if base.B2i32(l4 == v524)|base.B2i32(v507 == v524) != 0 {
		v538 = v524
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+89)))
	if (base.B2i32(v539 != int32(97))|v522)&int32(1) != 0 {
		v585 = v522
		v587 = v477
		v588 = v538
		goto L121
	} else {
		goto L133
	}
L130:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v507)+4))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v530)))
	if v531 != int32(6) {
		v538 = v524
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v530)+4))
	if v534 != l5 {
		v538 = v524
		goto L129
	} else {
		goto L132
	}
L132:
	;
	v536 = int32(*(*int16)(unsafe.Add(mBase, uint32(v530)+8)))
	v538 = v536
	goto L129
L133:
	;
	v545 = int32(1)
	switch l2 - v463 {
	case 0:
		v585 = v545
		v587 = v477
		v588 = v538
		goto L121
	case 1:
		goto L134
	default:
		goto L135
	}
L134:
	;
	v585 = int32(0)
	v587 = v477
	v588 = v538
	goto L121
L135:
	;
	if v538 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	if v477 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	goto L138
L138:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L145
	}
L139:
	;
	v548 = F_findDefaultOnlyColumns(m, l4)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	v550 = v477
	goto L141
L141:
	;
	v551 = F_bms_is_member(m, v538, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L143
	}
L142:
	;
	v550 = v548
	goto L141
L143:
	;
	if v551 != 0 {
		v585 = v545
		v587 = v550
		v588 = v538
		goto L121
	} else {
		goto L144
	}
L144:
	;
	goto L138
L145:
	;
	F_errcode(m, int32(156008580))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v562 = v501 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v562
	F_errmsg(m, int32(_a_F_rewriteTargetListIU_6), v27+int32(48))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v562
	F_errdetail(m, int32(_a_F_rewriteTargetListIU_7), v27+int32(32))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errhint(m, int32(_a_F_rewriteTargetListIU_8), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(915), int32(_a_F_rewriteTargetListIU_2))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	if v588 == int32(0) {
		goto L111
	} else {
		goto L154
	}
L152:
	;
	v616 = v598
	v617 = v587
	goto L153
L153:
	;
	v620 = int32(0)
	if base.B2i32(v616&int32(1) == v620)|(base.B2i32(l6 == v620)|base.B2i32(v588 == v620)) != 0 {
		v655 = v616
		v657 = v617
		goto L119
	} else {
		goto L161
	}
L154:
	;
	if v587 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v608 = F_findDefaultOnlyColumns(m, l4)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	v610 = v587
	goto L157
L157:
	;
	v612 = F_bms_is_member(m, v588, v610)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L159
	}
L158:
	;
	v610 = v608
	goto L157
L159:
	;
	if v612 == int32(0) {
		goto L111
	} else {
		goto L160
	}
L160:
	;
	v616 = int32(1)
	v617 = v610
	goto L153
L161:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v629 = F_bms_add_member(m, v628, v588)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v629
	v655 = int32(1)
	v657 = v617
	goto L119
L163:
	;
	v635 = int32(0)
	v636 = base.B2i32(v507 == v635)
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+89)))
	if (v636|base.B2i32(v637 != int32(97))|v522)&int32(1) == v635 {
		goto L110
	} else {
		goto L164
	}
L164:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+90)))
	v647 = int32(0)
	if (base.B2i32(v646 == v647)|v636|v522)&int32(1) == v647 {
		goto L109
	} else {
		goto L165
	}
L165:
	;
	v655 = v522
	v657 = v477
	goto L119
L166:
	;
	v661 = int32(0)
	goto L168
L167:
	;
	v661 = v507
	goto L168
L168:
	;
	if v660|base.B2i32(v655&int32(1) == int32(0)) != 0 {
		v696 = v661
		goto L169
	} else {
		goto L170
	}
L169:
	;
	if v696 == int32(0) {
		v704 = v657
		v707 = v490
		goto L117
	} else {
		goto L182
	}
L170:
	;
	v668 = F_build_column_default(m, l3, v475)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	if v668 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v670 = v507
	goto L174
L173:
	;
	v670 = int32(0)
	goto L174
L174:
	;
	if base.B2i32(l1 == int32(3))|v668 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v501)+68))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v501)+76))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v501)+96))
	v679 = int32(*(*int16)(unsafe.Add(mBase, uint32(v501)+72)))
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+82)))
	v681 = F_coerce_null_to_domain(m, v676, v677, v678, v679, v680)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	v683 = v668
	v684 = v670
	goto L177
L177:
	;
	if v683 == int32(0) {
		v696 = v684
		goto L169
	} else {
		goto L179
	}
L178:
	;
	v683 = v681
	v684 = v507
	goto L177
L179:
	;
	v690 = F_pstrdup(m, v501+int32(4))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v693 = F_makeTargetEntry(m, v683, base.I32_extend16_s(v475), v690, int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v696 = v693
	goto L169
L182:
	;
	v699 = F_lappend(m, v490, v696)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v704 = v657
	v707 = v699
	goto L117
L184:
	;
	goto L116
L185:
	;
	v737 = F_list_concat(m, v734, v457)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	m.G0 = v27 + int32(176)
	return v737
L187:
	;
	F_errcode(m, int32(156008580))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	v753 = v501 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v753
	F_errmsg(m, int32(_a_F_rewriteTargetListIU_6), v27+int32(16))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v753
	F_errdetail(m, int32(_a_F_rewriteTargetListIU_9), v27)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(953), int32(_a_F_rewriteTargetListIU_2))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	F_errcode(m, int32(156008580))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	v777 = v501 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v777
	F_errmsg(m, int32(_a_F_rewriteTargetListIU_10), v27+int32(112))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v777
	F_errdetail(m, int32(_a_F_rewriteTargetListIU_7), v27+int32(96))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(980), int32(_a_F_rewriteTargetListIU_2))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	F_errcode(m, int32(156008580))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v803 = v501 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v803
	F_errmsg(m, int32(_a_F_rewriteTargetListIU_10), v27+int32(80))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v803
	F_errdetail(m, int32(_a_F_rewriteTargetListIU_9), v27-int32(-64))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(988), int32(_a_F_rewriteTargetListIU_2))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
