package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QTN2QT(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v2
	F_cntsize(m, l0, v9+int32(28), v9+int32(24))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		if base.Ui32(v23) <= base.Ui32(int32(1073741815)) {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
			v30 = base.I32_div_u_s(int32(1073741815)-v23, int32(12))
			if base.Ui32(v26) <= base.Ui32(v30) {
				v50 = v26 * int32(12)
				v53 = v23 + v50 + int32(8)
				v54 = F_palloc0(m, v53)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v54))) = v53 << (uint(int32(2)) % 32)
					v61 = v54 + int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v61
					v63 = v61 + v50
					*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v63
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v63
					F_fillQT(m, v9+int32(12), l0)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(32)
						return v54
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_QTN2QT_0), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_QTN2QT_1), int32(383), int32(_a_F_QTN2QT_2))
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_QTN2QT_0), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_QTN2QT_1), int32(383), int32(_a_F_QTN2QT_2))
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
func F_QueueFKConstraintValidation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v157 int32
	_ = v157
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v196 int32
	_ = v196
	var v200 int64
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
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
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	F_check_stack_depth(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
	v23 = v21 + v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+119)))
	if v25 == int32(114) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v253 = F_heap_copytuple(m, l4)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L41
	}
L4:
	;
	v196 = v15 + int32(-56)
	v200 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23))))
	F_ScanKeyInit(m, v196, int32(12), int32(3), int32(184), v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L28
	}
L5:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
	v177 = F_get_rel_relkind(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L26
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
	if v28 != l3 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v157 = v25
	goto L8
L8:
	;
	if v157&int32(255) == int32(112) {
		goto L4
	} else {
		goto L25
	}
L9:
	;
	v31 = F_palloc0(m, int32(108))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(161)
	v37 = F_pstrdup(m, v23+int32(4))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v37
	v41 = F_palloc0(m, int32(32))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v43
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v23)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = v51
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v55 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v129)+64))
	v138 = F_lappend(m, v137, v41)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L24
	}
L14:
	;
	v101 = F_palloc0(m, int32(144))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L21
	}
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v58 <= int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v72 = int32(0)
	goto L17
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v61+v72<<(uint(int32(2))%32))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v81 == v54 {
		v129 = v80
		goto L13
	} else {
		goto L19
	}
L18:
	;
	goto L14
L19:
	;
	v84 = v72 + int32(1)
	if v58 != v84 {
		v72 = v84
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v54
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)) = uint8(v107)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v110 = F_CreateTupleDescCopyConstr(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+8)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v101)+88)) = int64(0)
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+84)) = uint8(v115)
	v117 = int32(_a_F_QueueFKConstraintValidation_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v101)+96)) = uint16(v117)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v120 = F_lappend(m, v119, v101)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v120
	v129 = v101
	goto L13
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+64)) = v138
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+119)))
	v157 = v142
	goto L8
L25:
	;
	goto L5
L26:
	;
	if v177 != int32(112) {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	goto L4
L28:
	;
	v204 = int32(1)
	v207 = F_systable_beginscan(m, l1, int32(2579), v204, int32(0), v204, v196)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L30
L30:
	;
	v223 = F_systable_getnext(m, v207)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	F_systable_endscan(m, v207)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L40
	}
L32:
	;
	if v223 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+22)))
	v227 = v225 + v226
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+76)))
	if v228 != 0 {
		goto L30
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L31
L36:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v227)+80))
	v230 = F_table_open(m, v229, l5)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_QueueFKConstraintValidation(m, l0, l1, v230, l3, v223, l5)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_relation_close(m, v230, int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L30
L40:
	;
	goto L3
L41:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v253)+16))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+22)))
	v258 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v255+v256)+76)) = uint8(v258)
	F_CatalogTupleUpdate(m, l1, v253+int32(4), v253)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_QueueFKConstraintValidation[0]))
	if v265 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v268 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2606), v267, v268, v268, v268)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_pfree(m, v253)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	m.G0 = v17 - int32(-64)
	return
}
func F_querytree(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	F_errstart_cold(m, int32(21), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		F_errmsg_internal(m, int32(_a_F_querytree_0), int32(0))
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			F_errfinish(m, int32(_a_F_querytree_1), int32(713), int32(_a_F_querytree_2))
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_quote_identifier(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v16 = base.B2i32(v7 == int32(95)) | base.B2i32(base.Ui32((v7-int32(97))&int32(255)) < base.Ui32(int32(26)))
	if v7 != 0 {
		v18 = v7
		v19 = l0
		v21 = v16
		v22 = int32(0)
		for {
			if base.Ui32((v18-int32(97))&int32(255)) < base.Ui32(int32(26)) {
				v45 = v21
				v46 = v22
			} else {
				v29 = int32(255)
				v30 = v18 & v29
				if base.B2i32(v30 == int32(95))|base.B2i32(base.Ui32((v18-int32(48))&v29) < base.Ui32(int32(10))) != 0 {
					v45 = v21
					v46 = v22
				} else {
					v45 = int32(0)
					v46 = v22 + base.B2i32(v30 == int32(34))
				}
			}
			v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
			if v49 != 0 {
				v18 = v49
				v19 = v19 + int32(1)
				v21 = v45
				v22 = v46
				continue
			} else {
				break
			}
			break
		}
		v57 = v45
		v59 = v46 + int32(3)
	} else {
		v57 = v16
		v59 = int32(3)
	}
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_quote_identifier[0])))
	if v61|base.B2i32(v57 == int32(0)) != 0 {
		v76 = F_strlen(m, l0)
		mBase = m.M
		v78 = F_palloc(m, v76+v59)
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return int32(0)
		} else {
			v80 = int32(34)
			*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v80)
			v82 = l0
			v83 = v78
			for {
				v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
				if v88 != int32(34) {
				} else {
					v95 = int32(34)
					*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)) = uint8(v95)
					v99 = v83 + int32(2)
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v88)
					v82 = v82 + int32(1)
					v83 = v99
					continue
				}
				if v88 == int32(0) {
					break
				} else {
					v99 = v83 + int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v88)
					v82 = v82 + int32(1)
					v83 = v99
					continue
				}
				break
			}
			v103 = int32(34)
			*(*uint16)(unsafe.Add(mBase, uint32(v83)+1)) = uint16(v103)
			return v78
		}
	} else {
		v66 = F_ScanKeywordLookup(m, l0, int32(_a_F_quote_identifier_0))
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return int32(0)
		} else {
			if v66 < int32(0) {
				return l0
			} else {
				v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+uint32(_c_F_quote_identifier[1]))))
				if v73 != 0 {
					v76 = F_strlen(m, l0)
					mBase = m.M
					v78 = F_palloc(m, v76+v59)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						v80 = int32(34)
						*(*uint8)(unsafe.Add(mBase, uint32(v78))) = uint8(v80)
						v82 = l0
						v83 = v78
						for {
							v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
							if v88 != int32(34) {
							} else {
								v95 = int32(34)
								*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)) = uint8(v95)
								v99 = v83 + int32(2)
								*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v88)
								v82 = v82 + int32(1)
								v83 = v99
								continue
							}
							if v88 == int32(0) {
								break
							} else {
								v99 = v83 + int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v88)
								v82 = v82 + int32(1)
								v83 = v99
								continue
							}
							break
						}
						v103 = int32(34)
						*(*uint16)(unsafe.Add(mBase, uint32(v83)+1)) = uint16(v103)
						return v78
					}
				} else {
					return l0
				}
			}
		}
	}
}
func F_quote_nullable(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v2 == int32(1) {
		v6 = F_cstring_to_text(m, int32(_a_F_quote_nullable_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v6)
		}
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v15 = F_DirectFunctionCall1Coll(m, int32(1679), int32(0), v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			return v15
		}
	}
}
