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
				v26 = int32(42)
				v28 = v26
			case 10:
				v28 = int32(38)
			case 29:
				v28 = int32(18)
			case 36:
				v28 = int32(23)
			case 45:
				v28 = int32(52)
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
											F_errfinish(m, int32(_a_F_TargetPrivilegesCheck_1), int32(2651), int32(_a_F_TargetPrivilegesCheck_2))
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
									F_errfinish(m, int32(_a_F_TargetPrivilegesCheck_1), int32(2651), int32(_a_F_TargetPrivilegesCheck_2))
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
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = int32(524)
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
	F_errfinish(m, int32(_a_F_addTargetToSortList_2), int32(3481), int32(_a_F_addTargetToSortList_3))
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
	F_errfinish(m, int32(_a_F_addTargetToSortList_2), int32(3517), int32(_a_F_addTargetToSortList_3))
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
	F_errfinish(m, int32(_a_F_addTargetToSortList_2), int32(3473), int32(_a_F_addTargetToSortList_3))
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
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
	var v223 int32
	_ = v223
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
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v474 int32
	_ = v474
	var v480 int32
	_ = v480
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
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
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
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
	v55 = v8
	v58 = v30 + int32(1)
	v65 = v8
	goto L7
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L105
	}
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v55<<(uint(int32(2))%32))))
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
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L101
	}
L9:
	;
	goto L8
L10:
	;
	v386 = v55 + int32(1)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v386 < v387 {
		v55 = v386
		v58 = v375
		v65 = v382
		goto L7
	} else {
		goto L100
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v351
	v375 = v58
	v382 = v65
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
	if v73 != v58 {
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
	v105 = v98 + v99<<(uint(int32(3))%32) + v73*int32(100)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+19)))
	if v106 != 0 {
		v375 = v58
		v382 = v65
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
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(861), int32(_a_F_rewriteTargetListIU_2))
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
	v116 = v105 - int32(72)
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
	v151 = F_lappend(m, v65, v148)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v145)+8)) = uint16(v58)
	v148 = v145
	goto L40
L42:
	;
	v375 = v58 + int32(1)
	v382 = v151
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
	v223 = v199
	goto L65
L65:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v234 = v232 - int32(14)
	if v234 != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v246 = F_equal(m, v223, v185)
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
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v223+v241)))
	if v243 != 0 {
		v223 = v243
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
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v223)+36))
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
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(1225), int32(_a_F_rewriteTargetListIU_4))
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
	v457 = v382
	goto L3
L101:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+144)) = v116 + int32(4)
	F_errmsg(m, int32(_a_F_rewriteTargetListIU_5), v27+int32(144))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(1169), int32(_a_F_rewriteTargetListIU_4))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
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
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+160)) = v116 + int32(4)
	F_errmsg(m, int32(_a_F_rewriteTargetListIU_5), v27+int32(160))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(1187), int32(_a_F_rewriteTargetListIU_4))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
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
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L197
	}
L110:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L192
	}
L111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L187
	}
L112:
	;
	v462 = int32(1)
	v474 = v462
	v480 = int32(0)
	v488 = v8
	goto L115
L113:
	;
	v734 = v8
	goto L114
L114:
	;
	F_pfree(m, v33)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L1
	} else {
		goto L185
	}
L115:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)))
	v497 = v490 + v491<<(uint(int32(3))%32) + v474*int32(100)
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+19)))
	if v498 != 0 {
		v707 = v480
		v708 = v488
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v734 = v708
	goto L114
L117:
	;
	if v474 != v30 {
		v474 = v474 + int32(1)
		v480 = v707
		v488 = v708
		goto L115
	} else {
		goto L184
	}
L118:
	;
	v500 = v497 - int32(72)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v474<<(uint(int32(2))%32)+v33-int32(4))))
	v508 = base.B2i32(l1 != int32(3))
	if v506|v508 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+90)))
	if v661 != 0 {
		goto L166
	} else {
		goto L167
	}
L120:
	;
	if l1 != int32(2) {
		v656 = v523
		v659 = v480
		goto L119
	} else {
		goto L163
	}
L121:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+90)))
	v591 = int32(0)
	v593 = int32(1)
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+89)))
	v599 = base.B2i32(l2 == v593)&base.B2i32(v595 == int32(100)) | v586
	if base.B2i32(v590 == v591)|v599&v593 == v591 {
		goto L151
	} else {
		goto L152
	}
L122:
	;
	v586 = int32(1)
	v587 = int32(0)
	v589 = v480
	goto L121
L123:
	;
	goto L124
L124:
	;
	v514 = int32(0)
	if v506 == v514 {
		v523 = v514
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
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	if v517 == int32(0) {
		v523 = v514
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v517)))
	v523 = base.B2i32(v520 == int32(57))
	goto L125
L128:
	;
	v525 = int32(0)
	if base.B2i32(l4 == v525)|base.B2i32(v506 == v525) != 0 {
		v538 = v525
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+89)))
	if (base.B2i32(v540 != int32(97))|v523)&int32(1) != 0 {
		v586 = v523
		v587 = v538
		v589 = v480
		goto L121
	} else {
		goto L133
	}
L130:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v531)))
	if v532 != int32(6) {
		v538 = v525
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v531)+4))
	if v535 != l5 {
		v538 = v525
		goto L129
	} else {
		goto L132
	}
L132:
	;
	v537 = int32(*(*int16)(unsafe.Add(mBase, uint32(v531)+8)))
	v538 = v537
	goto L129
L133:
	;
	v546 = int32(1)
	switch l2 - v462 {
	case 0:
		v586 = v546
		v587 = v538
		v589 = v480
		goto L121
	case 1:
		goto L134
	default:
		goto L135
	}
L134:
	;
	v586 = int32(0)
	v587 = v538
	v589 = v480
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
	if v480 == int32(0) {
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
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L145
	}
L139:
	;
	v549 = F_findDefaultOnlyColumns(m, l4)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	v551 = v480
	goto L141
L141:
	;
	v552 = F_bms_is_member(m, v538, v551)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L143
	}
L142:
	;
	v551 = v549
	goto L141
L143:
	;
	if v552 != 0 {
		v586 = v546
		v587 = v538
		v589 = v551
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
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v563 = v500 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v563
	F_errmsg(m, int32(_a_F_rewriteTargetListIU_6), v27+int32(48))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v563
	v574 = F_errdetail(m, int32(_a_F_rewriteTargetListIU_7), v27+int32(32))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errhint(m, int32(_a_F_rewriteTargetListIU_8), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(961), int32(_a_F_rewriteTargetListIU_2))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
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
	if v587 == int32(0) {
		goto L111
	} else {
		goto L154
	}
L152:
	;
	v617 = v599
	v618 = v589
	goto L153
L153:
	;
	v621 = int32(0)
	if base.B2i32(v617&int32(1) == v621)|(base.B2i32(l6 == v621)|base.B2i32(v587 == v621)) != 0 {
		v656 = v617
		v659 = v618
		goto L119
	} else {
		goto L161
	}
L154:
	;
	if v589 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v609 = F_findDefaultOnlyColumns(m, l4)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	v611 = v589
	goto L157
L157:
	;
	v613 = F_bms_is_member(m, v587, v611)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L159
	}
L158:
	;
	v611 = v609
	goto L157
L159:
	;
	if v613 == int32(0) {
		goto L111
	} else {
		goto L160
	}
L160:
	;
	v617 = int32(1)
	v618 = v611
	goto L153
L161:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v630 = F_bms_add_member(m, v629, v587)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v630
	v656 = int32(1)
	v659 = v618
	goto L119
L163:
	;
	v636 = int32(0)
	v637 = base.B2i32(v506 == v636)
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+89)))
	if (v637|base.B2i32(v638 != int32(97))|v523)&int32(1) == v636 {
		goto L110
	} else {
		goto L164
	}
L164:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+90)))
	v648 = int32(0)
	if (base.B2i32(v647 == v648)|v637|v523)&int32(1) == v648 {
		goto L109
	} else {
		goto L165
	}
L165:
	;
	v656 = v523
	v659 = v480
	goto L119
L166:
	;
	v662 = int32(0)
	goto L168
L167:
	;
	v662 = v506
	goto L168
L168:
	;
	if v661|base.B2i32(v656&int32(1) == int32(0)) != 0 {
		v697 = v662
		goto L169
	} else {
		goto L170
	}
L169:
	;
	if v697 == int32(0) {
		v707 = v659
		v708 = v488
		goto L117
	} else {
		goto L182
	}
L170:
	;
	v669 = F_build_column_default(m, l3, v474)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	if v669 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v671 = v506
	goto L174
L173:
	;
	v671 = int32(0)
	goto L174
L174:
	;
	if base.B2i32(l1 == int32(3))|v669 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v500)+68))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v500)+76))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v500)+96))
	v680 = int32(*(*int16)(unsafe.Add(mBase, uint32(v500)+72)))
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500)+82)))
	v682 = F_coerce_null_to_domain(m, v677, v678, v679, v680, v681)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	v684 = v669
	v685 = v671
	goto L177
L177:
	;
	if v684 == int32(0) {
		v697 = v685
		goto L169
	} else {
		goto L179
	}
L178:
	;
	v684 = v682
	v685 = v506
	goto L177
L179:
	;
	v691 = F_pstrdup(m, v500+int32(4))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v694 = F_makeTargetEntry(m, v684, base.I32_extend16_s(v474), v691, int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v697 = v694
	goto L169
L182:
	;
	v700 = F_lappend(m, v488, v697)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v707 = v659
	v708 = v700
	goto L117
L184:
	;
	goto L116
L185:
	;
	v738 = F_list_concat(m, v734, v457)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	m.G0 = v27 + int32(176)
	return v738
L187:
	;
	F_errcode(m, int32(156008580))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	v754 = v500 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v754
	F_errmsg(m, int32(_a_F_rewriteTargetListIU_6), v27+int32(16))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v754
	v763 = F_errdetail(m, int32(_a_F_rewriteTargetListIU_9), v27)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(1000), int32(_a_F_rewriteTargetListIU_2))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
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
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	v778 = v500 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v778
	F_errmsg(m, int32(_a_F_rewriteTargetListIU_10), v27+int32(112))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v778
	v789 = F_errdetail(m, int32(_a_F_rewriteTargetListIU_7), v27+int32(96))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(1027), int32(_a_F_rewriteTargetListIU_2))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
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
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v804 = v500 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v804
	F_errmsg(m, int32(_a_F_rewriteTargetListIU_10), v27+int32(80))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = v804
	v815 = F_errdetail(m, int32(_a_F_rewriteTargetListIU_9), v27-int32(-64))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_rewriteTargetListIU_1), int32(1035), int32(_a_F_rewriteTargetListIU_2))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
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
