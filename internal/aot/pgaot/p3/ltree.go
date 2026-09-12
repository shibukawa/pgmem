package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltree_r_risparent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(5599), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_ltree_addtext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_text_to_cstring(m, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = F_DirectFunctionCall1Coll(m, int32(5636), int32(0), v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v16)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = F_ltree_concat(m, v7, v18)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v18)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v26 != v7 {
									F_pfree(m, v7)
									mBase = m.M
									v29 = m.ExcPending
									if v29 != 0 {
										return int32(0)
									} else {
										v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if v30 != v14 {
											F_pfree(m, v14)
											mBase = m.M
											v33 = m.ExcPending
											if v33 != 0 {
												return int32(0)
											} else {
												return v22
											}
										} else {
											return v22
										}
									}
								} else {
									v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if v30 != v14 {
										F_pfree(m, v14)
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return int32(0)
										} else {
											return v22
										}
									} else {
										return v22
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
func F_ltree_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+14)))
	if v6 != int32(1) {
		return v5
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v12 != 0 {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v21 = int32(base.Ui32(v16)>>(uint(int32(2))%32)) + int32(8)
			} else {
				v21 = int32(8)
			}
			v22 = F_palloc(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = int32(1)
				v26 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v22))) = v21 << (uint(v26) % 32)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v33 = int32(base.Ui32(v31) >> (uint(v26) % 32))
				if v33 != 0 {
					v34 = F__emscripten_memcpy_bulkmem(m, v22+int32(8), v12, v33)
					mBase = m.M
				} else {
				}
				v37 = F_palloc(m, int32(16))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v37))) = v22
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v40
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v42
					v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+12)))
					v45 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v37)+14)) = uint8(v45)
					*(*uint16)(unsafe.Add(mBase, uint32(v37)+12)) = uint16(v44)
					return v37
				}
			}
		}
	}
}
func F_ltree_gist_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(282619), int32(159635), int32(8), int32(4), int32(2024))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
		v18 = F_lappend(m, v16, int32(6977))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v2)+4)) = v18
			return int32(0)
		}
	}
}
func F_ltree_gist_relopts_validator(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 != (v8+int32(3))&int32(-4) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(4)
				F_errmsg(m, int32(475816), v6)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errfinish(m, int32(493871), int32(731), int32(209940))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
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
		m.G0 = v6 + int32(16)
		return
	}
}
func F_ltree_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	v2 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v24 != int32(3) {
		v38 = v2
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
	if v39 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if int32(0) <= v27 {
		v38 = v27
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
	v32 = int32(0)
	if base.Ui32(v32-v27) < base.Ui32(v30) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v36 = v27 + v30
	goto L9
L8:
	;
	v36 = v32
	goto L9
L9:
	;
	v38 = v36
	goto L4
L10:
	;
	return v248
L11:
	;
	F_pfree(m, v22)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L59
	}
L12:
	;
	v55 = int32(-1)
	if base.Ui32(v42) < base.Ui32(v39) {
		v210 = v55
		goto L22
	} else {
		goto L23
	}
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v48 != v17 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if v39 <= v42-v38 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	F_pfree(m, v17)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v52 = int32(-1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v22 != v53 {
		v231 = v52
		goto L11
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v248 = v52
	goto L10
L22:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v218 != v17 {
		goto L54
	} else {
		goto L55
	}
L23:
	;
	v57 = int32(8)
	v70 = v2
	v72 = v17 + v57
	goto L24
L24:
	;
	if v70 < v38 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v210 = v55
	goto L22
L26:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72))))
	v201 = v70 + int32(1)
	if v201 != v42-v39+int32(1) {
		v70 = v201
		v72 = v72 + (v194+int32(9))&int32(131064)
		goto L24
	} else {
		goto L53
	}
L27:
	;
	v82 = v22 + v57
	v90 = v72
	v91 = int32(0)
	goto L29
L28:
	;
	if v39 != v91 {
		goto L26
	} else {
		goto L52
	}
L29:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v90))))
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82))))
	if v96 != v97 {
		goto L28
	} else {
		goto L31
	}
L30:
	;
	v210 = v70
	goto L22
L31:
	;
	v99 = int32(2)
	v100 = v90 + v99
	v102 = v82 + v99
	if base.Ui32(int32(4)) <= base.Ui32(v96) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	if v164 != 0 {
		goto L28
	} else {
		goto L50
	}
L33:
	;
	v164 = int32(0)
	goto L32
L34:
	;
	v138 = v133
	v139 = v134
	v140 = v135
	goto L44
L35:
	;
	if (v100|v102)&int32(3) != 0 {
		v133 = v100
		v134 = v102
		v135 = v96
		goto L34
	} else {
		goto L38
	}
L36:
	;
	v126 = v100
	v127 = v102
	v128 = v96
	goto L37
L37:
	;
	if v128 == int32(0) {
		goto L33
	} else {
		goto L43
	}
L38:
	;
	v110 = v100
	v111 = v102
	v112 = v96
	goto L39
L39:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v115 != v116 {
		v133 = v110
		v134 = v111
		v135 = v112
		goto L34
	} else {
		goto L41
	}
L40:
	;
	v126 = v121
	v127 = v119
	v128 = v123
	goto L37
L41:
	;
	v118 = int32(4)
	v119 = v111 + v118
	v121 = v110 + v118
	v123 = v112 - v118
	if base.Ui32(int32(3)) < base.Ui32(v123) {
		v110 = v121
		v111 = v119
		v112 = v123
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v133 = v126
	v134 = v127
	v135 = v128
	goto L34
L44:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v143 == v144 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v164 = v143 - v144
	goto L32
L46:
	;
	v146 = int32(1)
	v151 = v140 - v146
	if v151 != 0 {
		v138 = v138 + v146
		v139 = v139 + v146
		v140 = v151
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	goto L45
L49:
	;
	goto L33
L50:
	;
	v165 = int32(9)
	v167 = int32(131064)
	v176 = v91 + int32(1)
	if v176 != v39 {
		v82 = v82 + (v97+v165)&v167
		v90 = v90 + (v96+v165)&v167
		v91 = v176
		goto L29
	} else {
		goto L51
	}
L51:
	;
	goto L30
L52:
	;
	v210 = v70
	goto L22
L53:
	;
	goto L25
L54:
	;
	F_pfree(m, v17)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v22 == v222 {
		v248 = v210
		goto L10
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v231 = v210
	goto L11
L59:
	;
	v248 = v231
	goto L10
}
func F_ltree_isparent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v23 int32
	_ = v23
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = F_pg_detoast_datum(m, v9)
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v129 != v10 {
		goto L38
	} else {
		goto L39
	}
L4:
	;
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
	if base.Ui32(v18) < base.Ui32(v17) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v122 = int32(0)
	goto L3
L6:
	;
	goto L7
L7:
	;
	if v17 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v122 = int32(1)
	goto L3
L9:
	;
	goto L10
L10:
	;
	v23 = int32(8)
	v30 = v17
	v31 = v15 + v23
	v32 = v10 + v23
	goto L11
L11:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32))))
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31))))
	if v35 != v36 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v122 = v116
	goto L3
L13:
	;
	v122 = int32(0)
	goto L3
L14:
	;
	goto L15
L15:
	;
	v39 = int32(2)
	v40 = v32 + v39
	v42 = v31 + v39
	if base.Ui32(int32(4)) <= base.Ui32(v35) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v104 != 0 {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	v104 = int32(0)
	goto L16
L18:
	;
	v78 = v73
	v79 = v74
	v80 = v75
	goto L28
L19:
	;
	if (v40|v42)&int32(3) != 0 {
		v73 = v40
		v74 = v42
		v75 = v35
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v66 = v40
	v67 = v42
	v68 = v35
	goto L21
L21:
	;
	if v68 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v50 = v40
	v51 = v42
	v52 = v35
	goto L23
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v55 != v56 {
		v73 = v50
		v74 = v51
		v75 = v52
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v66 = v61
	v67 = v59
	v68 = v63
	goto L21
L25:
	;
	v58 = int32(4)
	v59 = v51 + v58
	v61 = v50 + v58
	v63 = v52 - v58
	if base.Ui32(int32(3)) < base.Ui32(v63) {
		v50 = v61
		v51 = v59
		v52 = v63
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v73 = v66
	v74 = v67
	v75 = v68
	goto L18
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v83 == v84 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v104 = v83 - v84
	goto L16
L30:
	;
	v86 = int32(1)
	v91 = v80 - v86
	if v91 != 0 {
		v78 = v78 + v86
		v79 = v79 + v86
		v80 = v91
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	v122 = int32(0)
	goto L3
L35:
	;
	goto L36
L36:
	;
	v106 = int32(9)
	v108 = int32(131064)
	v116 = int32(1)
	if v116 < v30 {
		v30 = v30 - v116
		v31 = v31 + (v36+v106)&v108
		v32 = v32 + (v35+v106)&v108
		goto L11
	} else {
		goto L37
	}
L37:
	;
	goto L12
L38:
	;
	F_pfree(m, v10)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v133 != v15 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	F_pfree(m, v15)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	return v122
L45:
	;
	goto L44
}
func F_ltree_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v15 = F_palloc(m, int32(base.Ui32(v12)>>(uint(int32(2))%32)))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
			if v17 == int32(0) {
				v66 = v15
			} else {
				v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+8)))
				if v22 != 0 {
					v23 = F__emscripten_memcpy_bulkmem(m, v15, v8+int32(10), v22)
					mBase = m.M
					v24 = v23
				} else {
					v24 = v15
				}
				v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+8)))
				v26 = v24 + v25
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
				if base.Ui32(v27) < base.Ui32(int32(2)) {
					v66 = v26
				} else {
					v38 = v8 + int32(8) + (v25+int32(9))&int32(131064)
					v40 = v26
					v42 = int32(1)
					for {
						v44 = int32(46)
						*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v44)
						v47 = v40 + int32(1)
						v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
						if v50 != 0 {
							v51 = F__emscripten_memcpy_bulkmem(m, v47, v38+int32(2), v50)
							mBase = m.M
							v52 = v51
						} else {
							v52 = v47
						}
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
						v54 = v52 + v53
						v61 = v42 + int32(1)
						v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
						if base.Ui32(v61) < base.Ui32(v62) {
							v38 = v38 + (v53+int32(9))&int32(131064)
							v40 = v54
							v42 = v61
							continue
						} else {
							break
						}
						break
					}
					v66 = v54
				}
			}
			v70 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v70)
			return v15
		}
	}
}
func F_ltree_recv(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pq_getmsgint(m, v8, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
				F_errmsg_internal(m, int32(472980), v6)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(497313), int32(238), int32(36896))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			v38 = F_pq_getmsgtext(m, v8, v33-v34, v6+int32(12))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v41 = F_parse_ltree(m, v38, int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v38)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(16)
						return v41
					}
				}
			}
		}
	}
}
