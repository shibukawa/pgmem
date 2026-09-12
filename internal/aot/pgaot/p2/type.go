package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecTypeFromExprList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	v2 = int32(0)
	if l0 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v58
L2:
	;
	v10 = F_CreateTemplateTupleDesc(m, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = F_CreateTemplateTupleDesc(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v58 = v10
	goto L1
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 <= int32(0) {
		v58 = v15
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v23 = v2
	v24 = int32(1)
	goto L9
L9:
	;
	v27 = base.I32_extend16_s(v24)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v23<<(uint(int32(2))%32))))
	v34 = F_exprType(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	v58 = v15
	goto L1
L11:
	;
	v36 = F_exprTypmod(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	F_TupleDescInitEntry(m, v15, v27, int32(0), v34, v36, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v41 = F_exprCollation(m, v33)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*int32)(unsafe.Add(mBase, uint32(v15+v43<<(uint(int32(4))%32)+v27*int32(100))+16)) = v41
	goto L15
L15:
	;
	v51 = int32(1)
	v54 = v23 + v51
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v54 < v55 {
		v23 = v54
		v24 = v24 + v51
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L10
}
func F_ExecTypeFromTLInternal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v104 = F_CreateTemplateTupleDesc(m, v103)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	v8 = int32(0)
	if l0 == v8 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	if l0 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L5:
	;
	v103 = v97
	goto L1
L6:
	;
	v97 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 <= int32(0) {
		v82 = v8
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v97 = v82
	goto L5
L10:
	;
	v21 = int32(0)
	if v21 < v18 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v24 = v18
	goto L13
L12:
	;
	v24 = v21
	goto L13
L13:
	;
	v25 = int32(1)
	if v18 == v25 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v24&v25 == int32(0) {
		v82 = v63
		goto L9
	} else {
		goto L21
	}
L15:
	;
	v29 = int32(0)
	v63 = v29
	v64 = v29
	goto L14
L16:
	;
	goto L17
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = int32(0)
	v37 = v34
	v38 = v34
	v39 = v8
	goto L18
L18:
	;
	v44 = int32(2)
	v46 = v33 + v38<<(uint(v44)%32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+26)))
	v49 = int32(1)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+26)))
	v56 = v37 + (v48 ^ v49) + (v53 ^ v49)
	v58 = v38 + v44
	v60 = v39 + v44
	if v60 != v24&int32(2147483646) {
		v37 = v56
		v38 = v58
		v39 = v60
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v63 = v56
	v64 = v58
	goto L14
L20:
	;
	goto L19
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72+v64<<(uint(int32(2))%32))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+26)))
	v82 = v63 + (v77 ^ int32(1))
	goto L9
L22:
	;
	v103 = v102
	goto L1
L23:
	;
	v102 = int32(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v102 = v101
	goto L22
L26:
	;
	return int32(0)
L27:
	;
	if l0 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	return v104
L29:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v110 <= int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v118 = int32(0)
	v119 = int32(1)
	goto L31
L31:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v118<<(uint(int32(2))%32))))
	if l1 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L28
L33:
	;
	v154 = v118 + int32(1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v154 < v155 {
		v118 = v154
		v119 = v151
		goto L31
	} else {
		goto L43
	}
L34:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+26)))
	if v126 != 0 {
		v151 = v119
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v127 = base.I32_extend16_s(v119)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v130 = F_exprType(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L26
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v133 = F_exprTypmod(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L26
	} else {
		goto L39
	}
L39:
	;
	F_TupleDescInitEntry(m, v104, v127, v128, v130, v133, int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v139 = F_exprCollation(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L26
	} else {
		goto L41
	}
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	*(*int32)(unsafe.Add(mBase, uint32(v104+v141<<(uint(int32(4))%32)+v127*int32(100))+16)) = v139
	goto L42
L42:
	;
	v151 = v119 + int32(1)
	goto L33
L43:
	;
	goto L32
}
func F_TypeCategory(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_get_type_category_preferred(m, l0, v5+int32(15), v5+int32(14))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5)+15)))
		m.G0 = v5 + int32(16)
		return v15
	}
}
func F_TypeShellMake(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v8 = m.G0
	v10 = v8 - int32(224)
	m.G0 = v10
	v14 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
		v17 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+88)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v17
		v30 = F__emscripten_memset_bulkmem(m, v10+int32(96), base.I32_extend8_s(int32(0)), int32(128))
		mBase = m.M
		v32 = F_strncpy(m, v10, l1, int32(64))
		mBase = m.M
		v33 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v32)+63)) = uint8(v33)
		v35 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+148)) = v35
		*(*int64)(unsafe.Add(mBase, uint32(v10)+172)) = v35
		v39 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = v39
		*(*int64)(unsafe.Add(mBase, uint32(v10)+140)) = v35
		*(*int32)(unsafe.Add(mBase, uint32(v10)+136)) = int32(44)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+128)) = v35
		*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = int64(343597383792)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+112)) = int64(4294967300)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = l2
		*(*int64)(unsafe.Add(mBase, uint32(v10)+164)) = v35
		*(*int64)(unsafe.Add(mBase, uint32(v10)+156)) = int64(10303626545502)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+208)) = v39
		*(*int64)(unsafe.Add(mBase, uint32(v10)+200)) = int64(4294967295)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+192)) = v35
		*(*int64)(unsafe.Add(mBase, uint32(v10)+184)) = int64(481036337257)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v10
		v66 = int32(257)
		*(*uint16)(unsafe.Add(mBase, uint32(v10)+93)) = uint16(v66)
		v68 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+95)) = uint8(v68)
		v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[293])))
		if v71 == v68 {
			v75 = *(*int32)(unsafe.Add(mBase, _consts[294]))
			if v75 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v132 = m.ExcPending
					if v132 != 0 {
						return
					} else {
						F_errmsg(m, int32(406915), int32(0))
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return
						} else {
							F_errfinish(m, int32(491118), int32(133), int32(392923))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
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
				*(*int32)(unsafe.Add(mBase, _consts[294])) = int32(0)
				v85 = v75
				*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v85
				v91 = F_heap_form_tuple(m, v16, v10+int32(96), v10-int32(-64))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					F_CatalogTupleInsert(m, v14, v91)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, _consts[71]))
						if v96 != 0 {
							v97 = int32(0)
							F_GenerateTypeDependencies(m, v91, v14, v97, v97, v97, v97, v97, int32(1), v97)
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, _consts[275]))
								if v107 != 0 {
									v109 = int32(0)
									F_RunObjectPostCreateHook(m, int32(1247), v85, v109, v109)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v85
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
										F_pfree(m, v91)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											F_sequence_close(m, v14, int32(3))
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return
											} else {
												m.G0 = v10 + int32(224)
												return
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v85
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
									F_pfree(m, v91)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										F_sequence_close(m, v14, int32(3))
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return
										} else {
											m.G0 = v10 + int32(224)
											return
										}
									}
								}
							}
						} else {
							v107 = *(*int32)(unsafe.Add(mBase, _consts[275]))
							if v107 != 0 {
								v109 = int32(0)
								F_RunObjectPostCreateHook(m, int32(1247), v85, v109, v109)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v85
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
									F_pfree(m, v91)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										F_sequence_close(m, v14, int32(3))
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return
										} else {
											m.G0 = v10 + int32(224)
											return
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v85
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
								F_pfree(m, v91)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return
								} else {
									F_sequence_close(m, v14, int32(3))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										m.G0 = v10 + int32(224)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v83 = F_GetNewOidWithIndex(m, v14, int32(2703), int32(1))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return
			} else {
				v85 = v83
				*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v85
				v91 = F_heap_form_tuple(m, v16, v10+int32(96), v10-int32(-64))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					F_CatalogTupleInsert(m, v14, v91)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, _consts[71]))
						if v96 != 0 {
							v97 = int32(0)
							F_GenerateTypeDependencies(m, v91, v14, v97, v97, v97, v97, v97, int32(1), v97)
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, _consts[275]))
								if v107 != 0 {
									v109 = int32(0)
									F_RunObjectPostCreateHook(m, int32(1247), v85, v109, v109)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v85
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
										F_pfree(m, v91)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return
										} else {
											F_sequence_close(m, v14, int32(3))
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return
											} else {
												m.G0 = v10 + int32(224)
												return
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v85
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
									F_pfree(m, v91)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										F_sequence_close(m, v14, int32(3))
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return
										} else {
											m.G0 = v10 + int32(224)
											return
										}
									}
								}
							}
						} else {
							v107 = *(*int32)(unsafe.Add(mBase, _consts[275]))
							if v107 != 0 {
								v109 = int32(0)
								F_RunObjectPostCreateHook(m, int32(1247), v85, v109, v109)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v85
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
									F_pfree(m, v91)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										F_sequence_close(m, v14, int32(3))
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return
										} else {
											m.G0 = v10 + int32(224)
											return
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v85
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
								F_pfree(m, v91)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return
								} else {
									F_sequence_close(m, v14, int32(3))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										m.G0 = v10 + int32(224)
										return
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
func F_appendTypeNameToBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v50 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v9 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = F_format_type_be(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L14
	} else {
		goto L18
	}
L5:
	;
	v15 = v3
	goto L6
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v21 = v20
	goto L10
L9:
	;
	v21 = v3
	goto L10
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v25 = v22 + v15<<(uint(int32(2))%32)
	if v21 != v25 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_appendStringInfoChar(m, l1, int32(46))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	F_appendStringInfoString(m, l1, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L14
	} else {
		goto L16
	}
L14:
	;
	return
L15:
	;
	goto L13
L16:
	;
	v35 = v15 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v35 < v36 {
		v15 = v35
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L7
L18:
	;
	F_appendStringInfoString(m, l1, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	F_appendStringInfoString(m, l1, int32(531423))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v56 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	F_appendStringInfoString(m, l1, int32(499822))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L14
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	return
L27:
	;
	goto L26
}
func F_assign_record_type_identifier(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v43 int64
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 != int32(2249) {
		v12 = F_lookup_type_cache(m, l0, int32(256))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
			if v16 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int64(0)
					} else {
						v55 = F_format_type_be(m, l0)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v55
							F_errmsg(m, int32(343899), v7)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(491505), int32(2148), int32(219313))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
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
			} else {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v12)+192))
				v43 = v19
				m.G0 = v7 + int32(16)
				return v43
			}
		}
	} else {
		if l1 < int32(0) {
			v35 = int32(4089456)
			v37 = *(*int64)(unsafe.Add(mBase, _consts[1181]))
			v39 = v37 + int64(1)
			*(*int64)(unsafe.Add(mBase, _consts[1181])) = v39
			v43 = v39
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, _consts[1182]))
			if v23 <= l1 {
				v35 = int32(4089456)
				v37 = *(*int64)(unsafe.Add(mBase, _consts[1181]))
				v39 = v37 + int64(1)
				*(*int64)(unsafe.Add(mBase, _consts[1181])) = v39
				v43 = v39
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _consts[1183]))
				v29 = v26 + l1<<(uint(int32(4))%32)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				if v30 == int32(0) {
					v35 = int32(4089456)
					v37 = *(*int64)(unsafe.Add(mBase, _consts[1181]))
					v39 = v37 + int64(1)
					*(*int64)(unsafe.Add(mBase, _consts[1181])) = v39
					v43 = v39
				} else {
					v33 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
					v43 = v33
				}
			}
		}
		m.G0 = v7 + int32(16)
		return v43
	}
}
func F_findTypeSubscriptingFunction(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(2281)
	v10 = int32(1)
	v14 = F_LookupFuncName(m, l0, v10, v6+int32(44), v10)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v18 = F_get_func_rettype(m, v14)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 != int32(2281) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(117833860))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v57 = F_NameListToString(m, l0)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(307739)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v57
								F_errmsg(m, int32(188534), v6+int32(32))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(486893), int32(2296), int32(251044))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
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
					if v14 == int32(6179) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								v79 = F_NameListToString(m, l0)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v79
									F_errmsg(m, int32(180988), v6+int32(16))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(486893), int32(2306), int32(251044))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
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
						m.G0 = v6 + int32(48)
						return v14
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(52461700))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v39 = F_func_signature_string(m, l0, int32(1), int32(0), v6+int32(44))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
						F_errmsg(m, int32(68933), v6)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(486893), int32(2290), int32(251044))
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
func F_format_type_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
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
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
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
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(80)
	return v281
L2:
	;
	v27 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L10
	}
L3:
	;
	if l2&int32(8) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v281 = int32(0)
	goto L1
L5:
	;
	goto L6
L6:
	;
	if l2&int32(2) == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v22 = F_pstrdup(m, int32(646185))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v281 = v22
	goto L1
L10:
	;
	if v27 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if l2&int32(8) != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v54 = v52 + v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+92))
	if v55 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v281 = int32(0)
	goto L1
L15:
	;
	goto L16
L16:
	;
	if l2&int32(2) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v37 = F_pstrdup(m, int32(537501))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L21
	}
L20:
	;
	v281 = v37
	goto L1
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg_internal(m, int32(49916), v12)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(491104), int32(137), int32(454377))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v104 = l2 & int32(1)
	v107 = base.B2i32(int32(0) <= l1) & l2
	if v99 <= int32(1082) {
		goto L68
	} else {
		goto L69
	}
L25:
	;
	v99 = l0
	v100 = v54
	v101 = v27
	v102 = v4
	goto L24
L26:
	;
	goto L27
L27:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+88))
	if v58 != int32(6179) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v99 = l0
	v100 = v54
	v101 = v27
	v102 = v4
	goto L24
L29:
	;
	goto L30
L30:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+129)))
	if v61 == int32(112) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v99 = l0
	v100 = v54
	v101 = v27
	v102 = v4
	goto L24
L32:
	;
	goto L33
L33:
	;
	F_ReleaseCatCache(m, v27)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v67 = F_SearchSysCache1(m, int32(82), v55)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v67 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if l2&int32(8) != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+22)))
	v99 = v55
	v100 = v94 + v95
	v101 = v67
	v102 = int32(1)
	goto L24
L39:
	;
	v281 = int32(0)
	goto L1
L40:
	;
	goto L41
L41:
	;
	if l2&int32(2) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v77 = F_pstrdup(m, int32(499815))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L46
	}
L45:
	;
	v281 = v77
	goto L1
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = l0
	F_errmsg_internal(m, int32(49916), v12-int32(-64))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(491104), int32(162), int32(454377))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
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
	if v102 != 0 {
		goto L157
	} else {
		goto L158
	}
L50:
	;
	if l2&int32(4) == int32(0) {
		goto L143
	} else {
		goto L144
	}
L51:
	;
	if v229 != 0 {
		v270 = v229
		goto L49
	} else {
		goto L141
	}
L52:
	;
	if v99 != int32(114) {
		goto L50
	} else {
		goto L139
	}
L53:
	;
	if v107 != 0 {
		goto L134
	} else {
		goto L135
	}
L54:
	;
	if v107 != 0 {
		goto L129
	} else {
		goto L130
	}
L55:
	;
	if v107 != 0 {
		goto L124
	} else {
		goto L125
	}
L56:
	;
	v201 = F_pstrdup(m, int32(366846))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L8
	} else {
		goto L123
	}
L57:
	;
	if v107 != 0 {
		goto L118
	} else {
		goto L119
	}
L58:
	;
	if v107 != 0 {
		goto L113
	} else {
		goto L114
	}
L59:
	;
	if v107 != 0 {
		goto L108
	} else {
		goto L109
	}
L60:
	;
	v177 = F_pstrdup(m, int32(483384))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L8
	} else {
		goto L107
	}
L61:
	;
	v174 = F_pstrdup(m, int32(88387))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L8
	} else {
		goto L106
	}
L62:
	;
	v171 = F_pstrdup(m, int32(222286))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L8
	} else {
		goto L105
	}
L63:
	;
	v168 = F_pstrdup(m, int32(88202))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L8
	} else {
		goto L104
	}
L64:
	;
	v165 = F_pstrdup(m, int32(268150))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L103
	}
L65:
	;
	v162 = F_pstrdup(m, int32(309313))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L102
	}
L66:
	;
	if v107 != 0 {
		goto L96
	} else {
		goto L97
	}
L67:
	;
	v152 = F_pstrdup(m, int32(279925))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L8
	} else {
		goto L95
	}
L68:
	;
	if v99 <= int32(699) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	if v99 <= int32(1265) {
		goto L75
	} else {
		goto L76
	}
L71:
	;
	switch v99 - int32(16) {
	case 0:
		goto L67
	case 1, 2, 3, 6:
		goto L50
	case 4:
		goto L61
	case 5:
		goto L63
	case 7:
		goto L62
	default:
		goto L52
	}
L72:
	;
	goto L73
L73:
	;
	switch v99 - int32(700) {
	case 0:
		goto L65
	case 1:
		goto L64
	default:
		goto L74
	}
L74:
	;
	switch v99 - int32(1042) {
	case 0:
		goto L66
	case 1:
		goto L53
	default:
		goto L50
	}
L75:
	;
	switch v99 - int32(1184) {
	case 0:
		goto L55
	case 1:
		goto L50
	case 2:
		goto L59
	default:
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	switch v99 - int32(1560) {
	case 0:
		goto L83
	case 1:
		goto L50
	case 2:
		goto L54
	default:
		goto L84
	}
L78:
	;
	if v99 == int32(1083) {
		goto L58
	} else {
		goto L79
	}
L79:
	;
	if v99 != int32(1114) {
		goto L50
	} else {
		goto L80
	}
L80:
	;
	if v107 == int32(0) {
		goto L56
	} else {
		goto L81
	}
L81:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v100)+120))
	v130 = F_printTypmod(m, int32(233966), l1, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	v229 = v130
	goto L51
L83:
	;
	if v107 != 0 {
		goto L89
	} else {
		goto L90
	}
L84:
	;
	if v99 == int32(1266) {
		goto L57
	} else {
		goto L85
	}
L85:
	;
	if v99 != int32(1700) {
		goto L50
	} else {
		goto L86
	}
L86:
	;
	if v107 == int32(0) {
		goto L60
	} else {
		goto L87
	}
L87:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v100)+120))
	v142 = F_printTypmod(m, int32(483384), l1, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	v229 = v142
	goto L51
L89:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v100)+120))
	v146 = F_printTypmod(m, int32(102053), l1, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L8
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if v104 != 0 {
		goto L50
	} else {
		goto L93
	}
L92:
	;
	v229 = v146
	goto L51
L93:
	;
	v149 = F_pstrdup(m, int32(102053))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	v229 = v149
	goto L51
L95:
	;
	v229 = v152
	goto L51
L96:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v100)+120))
	v156 = F_printTypmod(m, int32(214536), l1, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L8
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v104 != 0 {
		goto L50
	} else {
		goto L100
	}
L99:
	;
	v229 = v156
	goto L51
L100:
	;
	v159 = F_pstrdup(m, int32(214536))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	v229 = v159
	goto L51
L102:
	;
	v229 = v162
	goto L51
L103:
	;
	v229 = v165
	goto L51
L104:
	;
	v229 = v168
	goto L51
L105:
	;
	v229 = v171
	goto L51
L106:
	;
	v229 = v174
	goto L51
L107:
	;
	v229 = v177
	goto L51
L108:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v100)+120))
	v181 = F_printTypmod(m, int32(304930), l1, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v184 = F_pstrdup(m, int32(304930))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L8
	} else {
		goto L112
	}
L111:
	;
	v229 = v181
	goto L51
L112:
	;
	v229 = v184
	goto L51
L113:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v100)+120))
	v188 = F_printTypmod(m, int32(370241), l1, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L8
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v191 = F_pstrdup(m, int32(366874))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L8
	} else {
		goto L117
	}
L116:
	;
	v229 = v188
	goto L51
L117:
	;
	v229 = v191
	goto L51
L118:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v100)+120))
	v195 = F_printTypmod(m, int32(370241), l1, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L8
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v198 = F_pstrdup(m, int32(366922))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L8
	} else {
		goto L122
	}
L121:
	;
	v229 = v195
	goto L51
L122:
	;
	v229 = v198
	goto L51
L123:
	;
	v229 = v201
	goto L51
L124:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v100)+120))
	v205 = F_printTypmod(m, int32(233966), l1, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L8
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v208 = F_pstrdup(m, int32(366897))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L8
	} else {
		goto L128
	}
L127:
	;
	v229 = v205
	goto L51
L128:
	;
	v229 = v208
	goto L51
L129:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v100)+120))
	v212 = F_printTypmod(m, int32(323430), l1, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L8
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v215 = F_pstrdup(m, int32(323430))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L8
	} else {
		goto L133
	}
L132:
	;
	v229 = v212
	goto L51
L133:
	;
	v229 = v215
	goto L51
L134:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v100)+120))
	v219 = F_printTypmod(m, int32(323442), l1, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v222 = F_pstrdup(m, int32(323442))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L8
	} else {
		goto L138
	}
L137:
	;
	v229 = v219
	goto L51
L138:
	;
	v229 = v222
	goto L51
L139:
	;
	v227 = F_pstrdup(m, int32(241910))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L8
	} else {
		goto L140
	}
L140:
	;
	v229 = v227
	goto L51
L141:
	;
	goto L50
L142:
	;
	v246 = F_quote_qualified_identifier(m, v243, v100+int32(4))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L8
	} else {
		goto L149
	}
L143:
	;
	v235 = int32(0)
	v237 = F_TypeIsVisibleExt(m, v99, v235)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L8
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v100)+68))
	v241 = F_get_namespace_name_or_temp(m, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L8
	} else {
		goto L148
	}
L146:
	;
	if v237 != 0 {
		v243 = v235
		goto L142
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v243 = v241
	goto L142
L149:
	;
	if v107 == int32(0) {
		v270 = v246
		goto L49
	} else {
		goto L150
	}
L150:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v100)+120))
	if v250 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v246
	v258 = F_psprintf(m, int32(654856), v12+int32(32))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v261 = F_OidFunctionCall1Coll(m, v250, int32(0), l1)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L8
	} else {
		goto L155
	}
L154:
	;
	v270 = v258
	goto L49
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v246
	v268 = F_psprintf(m, int32(173237), v12+int32(48))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L8
	} else {
		goto L156
	}
L156:
	;
	v270 = v268
	goto L49
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v270
	v276 = F_psprintf(m, int32(499774), v12+int32(16))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L8
	} else {
		goto L160
	}
L158:
	;
	v278 = v270
	goto L159
L159:
	;
	F_ReleaseCatCache(m, v101)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L8
	} else {
		goto L161
	}
L160:
	;
	v278 = v276
	goto L159
L161:
	;
	v281 = v278
	goto L1
}
func F_getTypeBinaryOutputInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v17 = v15 + v16
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+82)))
			if v18 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						v57 = F_format_type_be(m, l0)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v57
							F_errmsg(m, int32(299893), v10+int32(32))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errfinish(m, int32(491493), int32(3127), int32(238683))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
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
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+112))
				if v21 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							v77 = F_format_type_be(m, l0)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v77
								F_errmsg(m, int32(187100), v10+int32(16))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									F_errfinish(m, int32(491493), int32(3132), int32(238683))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+78)))
					if v25 != 0 {
						v30 = int32(0)
					} else {
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+76)))
						v30 = base.B2i32(v27 == int32(65535))
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v30)
					F_ReleaseCatCache(m, v13)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						m.G0 = v10 + int32(48)
						return
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(49916), v10)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errfinish(m, int32(491493), int32(3120), int32(238683))
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
}
func F_getTypeInputInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v16 = v14 + v15
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+82)))
			if v17 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						v56 = F_format_type_be(m, l0)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v56
							F_errmsg(m, int32(299893), v9+int32(32))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_errfinish(m, int32(491493), int32(3028), int32(238748))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
				if v20 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v76 = F_format_type_be(m, l0)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v76
								F_errmsg(m, int32(187236), v9+int32(16))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_errfinish(m, int32(491493), int32(3033), int32(238748))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v20
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
					v26 = v24 + v25
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
					if v27 != 0 {
						v29 = v27
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
						v29 = v28
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v29
					F_ReleaseCatCache(m, v12)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						m.G0 = v9 + int32(48)
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
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(49916), v9)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_errfinish(m, int32(491493), int32(3021), int32(238748))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
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
}
func F_get_record_type_from_argument(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_get_fn_expr_argtype(m, v9, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v11
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
		F_prepare_column_cache(m, l2+int32(4), v11, int32(-1), v17, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
			if v21|int32(32) != int32(99) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errcode(m, int32(67141764))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
						F_errmsg(m, int32(361769), v7)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_errfinish(m, int32(487084), int32(3650), int32(92656))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
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
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_parseTypeString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = F_typeStringToTypeName(m, l0, l3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v96 = int32(0)
			m.G0 = v9 + int32(48)
			return v96
		} else {
			if l3 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				v24 = base.B2i32(v20 == int32(447))
			} else {
				v24 = int32(0)
			}
			v25 = F_LookupTypeNameExtended(m, int32(0), v11, l2, int32(1), v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v25 == int32(0) {
					v29 = int32(0)
					v30 = F_errsave_start(m, l3)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 == int32(0) {
							v96 = v29
							m.G0 = v9 + int32(48)
							return v96
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								F_initStringInfo(m, v9+int32(32))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									F_appendTypeNameToBuffer(m, v11, v9+int32(32))
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v45
										F_errmsg(m, int32(71432), v9)
										mBase = m.M
										v49 = m.ExcPending
										if v49 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, l3, int32(491128), int32(802), int32(325720))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												v96 = v29
												m.G0 = v9 + int32(48)
												return v96
											}
										}
									}
								}
							}
						}
					}
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+22)))
					v57 = v55 + v56
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+82)))
					if v58 == int32(0) {
						F_ReleaseCatCache(m, v25)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = int32(0)
							v64 = F_errsave_start(m, l3)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								if v64 == int32(0) {
									v96 = v63
									m.G0 = v9 + int32(48)
									return v96
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_initStringInfo(m, v9+int32(32))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											F_appendTypeNameToBuffer(m, v11, v9+int32(32))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												v79 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v79
												F_errmsg(m, int32(299917), v9+int32(16))
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, l3, int32(491128), int32(814), int32(325720))
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return int32(0)
													} else {
														v96 = v63
														m.G0 = v9 + int32(48)
														return v96
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v91
						F_ReleaseCatCache(m, v25)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							v96 = int32(1)
							m.G0 = v9 + int32(48)
							return v96
						}
					}
				}
			}
		}
	}
}
func F_typeDepNeeded(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	goto L4
L1:
	;
	m.G0 = v10 + int32(16)
	return v85
L2:
	;
	if v26 != 0 {
		v85 = v3
		goto L1
	} else {
		goto L6
	}
L3:
	;
	goto L2
L4:
	;
	if base.Ui32(int32(11999)) < base.Ui32(l0) {
		v26 = v3
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v19 = int32(1)
	v26 = (v19 | base.B2i32(l0 != int32(2200))) & v19
	goto L3
L6:
	;
	v27 = int32(1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v28 == v27 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_get_func_signature(m, v31, v10+int32(12), v10+int32(8))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_op_input_types(m, v70, v10+int32(12), v10+int32(8))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L19
	}
L10:
	;
	return int32(0)
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v41 <= int32(0) {
		v64 = v27
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_pfree(m, v40)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L18
	}
L13:
	;
	v46 = int32(0)
	goto L14
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40+v46<<(uint(int32(2))%32))))
	v56 = base.B2i32(l0 != v55)
	if l0 == v55 {
		v64 = v56
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v64 = v56
	goto L12
L16:
	;
	v59 = v46 + int32(1)
	if v59 != v41 {
		v46 = v59
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v85 = v64
	goto L1
L19:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v85 = base.B2i32(l0 != v77) & base.B2i32(l0 != v79)
	goto L1
}
