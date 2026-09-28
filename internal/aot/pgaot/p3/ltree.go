package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__ltree_r_risparent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F__ltree_r_risparent_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_ltree_addtext(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = F_text_to_cstring(m, v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v20 = F_DirectFunctionCall1Coll(m, int32(_a_F_ltree_addtext_0), int32(0), base.I64_extend_i32_u(v17))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					F_pfree(m, v17)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						v24 = base.I32_wrap_i64(v20)
						v25 = F_ltree_concat(m, v8, v24)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v24)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int64(0)
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v29 != v8 {
									F_pfree(m, v8)
									mBase = m.M
									v32 = m.ExcPending
									if v32 != 0 {
										return int64(0)
									} else {
										v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v33 != v15 {
											F_pfree(m, v15)
											mBase = m.M
											v36 = m.ExcPending
											if v36 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v25)
											}
										} else {
											return base.I64_extend_i32_u(v25)
										}
									}
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v33 != v15 {
										F_pfree(m, v15)
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return int64(0)
										} else {
											return base.I64_extend_i32_u(v25)
										}
									} else {
										return base.I64_extend_i32_u(v25)
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
func F_ltree_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+18)))
	if v6 != int32(1) {
		return base.I64_extend_i32_u(v5)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			if v13 != 0 {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v22 = int32(base.Ui32(v17)>>(uint(int32(2))%32)) + int32(8)
			} else {
				v22 = int32(8)
			}
			v23 = F_palloc(m, v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(1)
				v27 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v22 << (uint(v27) % 32)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v32 = int32(base.Ui32(v30) >> (uint(v27) % 32))
				if v32 != 0 {
					base.MemoryCopy(m, v23+int32(8), v13, v32)
				} else {
				}
				v37 = F_palloc(m, int32(24))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v37))) = base.I64_extend_i32_u(v23)
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v41
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v43
					v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+16)))
					v46 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v37)+18)) = uint8(v46)
					*(*uint16)(unsafe.Add(mBase, uint32(v37)+16)) = uint16(v45)
					return base.I64_extend_i32_u(v37)
				}
			}
		}
	}
}
func F_ltree_gist_options(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(_a_F_ltree_gist_options_0), int32(_a_F_ltree_gist_options_1), int32(8), int32(4), int32(2024))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
		v18 = F_lappend(m, v16, int32(_a_F_ltree_gist_options_2))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v2)+4)) = v18
			return int64(0)
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v8&int32(3) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(4)
				F_errmsg(m, int32(_a_F_ltree_gist_relopts_validator_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ltree_gist_relopts_validator_1), int32(732), int32(_a_F_ltree_gist_relopts_validator_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
func F_ltree_index(m *base.Module, l0 int32) int64 {
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
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
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
	var v212 int32
	_ = v212
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v255 int64
	_ = v255
	var v257 int32
	_ = v257
	var v272 int64
	_ = v272
	v2 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v36 = v30 + v27
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
	return v272
L11:
	;
	F_pfree(m, v22)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L62
	}
L12:
	;
	if base.Ui32(v39) <= base.Ui32(v42) {
		goto L22
	} else {
		goto L23
	}
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v52 = int64(-1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v22 != v53 {
		v255 = v52
		goto L11
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v272 = v52
	goto L10
L22:
	;
	v57 = int32(8)
	v67 = v17 + v57
	v72 = v2
	goto L26
L23:
	;
	v234 = int64(-1)
	goto L24
L24:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v235 != v17 {
		goto L57
	} else {
		goto L58
	}
L25:
	;
	v234 = base.I64_extend_i32_s(v212)
	goto L24
L26:
	;
	if v38 <= v72 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v212 = int32(-1)
	goto L25
L28:
	;
	v82 = v22 + v57
	v88 = v67
	v90 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67))))
	v201 = v72 + int32(1)
	if v201 != v42-v39+int32(1) {
		v67 = v67 + (v194+int32(9))&int32(_a_F_ltree_index_0)
		v72 = v201
		goto L26
	} else {
		goto L56
	}
L31:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88))))
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82))))
	if v96 != v97 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v39 == v90 {
		v212 = v72
		goto L25
	} else {
		goto L55
	}
L33:
	;
	goto L32
L34:
	;
	v99 = int32(2)
	v100 = v88 + v99
	v102 = v82 + v99
	if base.Ui32(int32(4)) <= base.Ui32(v96) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	if v164 != 0 {
		goto L33
	} else {
		goto L53
	}
L36:
	;
	v164 = int32(0)
	goto L35
L37:
	;
	v138 = v133
	v139 = v134
	v140 = v135
	goto L47
L38:
	;
	if (v100|v102)&int32(3) != 0 {
		v133 = v100
		v134 = v102
		v135 = v96
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v126 = v100
	v127 = v102
	v128 = v96
	goto L40
L40:
	;
	if v128 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v110 = v100
	v111 = v102
	v112 = v96
	goto L42
L42:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v115 != v116 {
		v133 = v110
		v134 = v111
		v135 = v112
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v126 = v121
	v127 = v119
	v128 = v123
	goto L40
L44:
	;
	v118 = int32(4)
	v119 = v111 + v118
	v121 = v110 + v118
	v123 = v112 - v118
	if base.Ui32(int32(3)) < base.Ui32(v123) {
		v110 = v121
		v111 = v119
		v112 = v123
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v133 = v126
	v134 = v127
	v135 = v128
	goto L37
L47:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v143 == v144 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v164 = v143 - v144
	goto L35
L49:
	;
	v146 = int32(1)
	v151 = v140 - v146
	if v151 != 0 {
		v138 = v138 + v146
		v139 = v139 + v146
		v140 = v151
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	v165 = int32(9)
	v167 = int32(_a_F_ltree_index_0)
	v176 = v90 + int32(1)
	if v176 != v39 {
		v82 = v82 + (v97+v165)&v167
		v88 = v88 + (v96+v165)&v167
		v90 = v176
		goto L31
	} else {
		goto L54
	}
L54:
	;
	v212 = v72
	goto L25
L55:
	;
	goto L30
L56:
	;
	goto L27
L57:
	;
	F_pfree(m, v17)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v22 == v239 {
		v272 = v234
		goto L10
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v255 = v234
	goto L11
L62:
	;
	v272 = v255
	goto L10
}
func F_ltree_isparent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	v9 = int64(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v136 != v11 {
		goto L32
	} else {
		goto L33
	}
L4:
	;
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
	if base.Ui32(v19) < base.Ui32(v18) {
		v135 = v9
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v21 = int32(8)
	v27 = v18
	v28 = v16 + v21
	v29 = v11 + v21
	goto L9
L7:
	;
	goto L8
L8:
	;
	v135 = int64(1)
	goto L3
L9:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29))))
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28))))
	if v34 != v35 {
		v135 = v9
		goto L3
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v37 = int32(2)
	v38 = v29 + v37
	v40 = v28 + v37
	if base.Ui32(int32(4)) <= base.Ui32(v34) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if v102 != 0 {
		v135 = v9
		goto L3
	} else {
		goto L30
	}
L13:
	;
	v102 = int32(0)
	goto L12
L14:
	;
	v76 = v71
	v77 = v72
	v78 = v73
	goto L24
L15:
	;
	if (v38|v40)&int32(3) != 0 {
		v71 = v38
		v72 = v40
		v73 = v34
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v64 = v38
	v65 = v40
	v66 = v34
	goto L17
L17:
	;
	if v66 == int32(0) {
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v48 = v38
	v49 = v40
	v50 = v34
	goto L19
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v53 != v54 {
		v71 = v48
		v72 = v49
		v73 = v50
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v64 = v59
	v65 = v57
	v66 = v61
	goto L17
L21:
	;
	v56 = int32(4)
	v57 = v49 + v56
	v59 = v48 + v56
	v61 = v50 - v56
	if base.Ui32(int32(3)) < base.Ui32(v61) {
		v48 = v59
		v49 = v57
		v50 = v61
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v71 = v64
	v72 = v65
	v73 = v66
	goto L14
L24:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v81 == v82 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v102 = v81 - v82
	goto L12
L26:
	;
	v84 = int32(1)
	v89 = v78 - v84
	if v89 != 0 {
		v76 = v76 + v84
		v77 = v77 + v84
		v78 = v89
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L13
L30:
	;
	v103 = int32(9)
	v105 = int32(_a_F_ltree_isparent_0)
	v113 = int32(1)
	if v113 < v27 {
		v27 = v27 - v113
		v28 = v28 + (v35+v103)&v105
		v29 = v29 + (v34+v103)&v105
		goto L9
	} else {
		goto L31
	}
L31:
	;
	goto L10
L32:
	;
	F_pfree(m, v11)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v140 != v16 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	F_pfree(m, v16)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	return v135
L39:
	;
	goto L38
}
func F_ltree_out(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v15 = F_palloc(m, int32(base.Ui32(v12)>>(uint(int32(2))%32)))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
			if v17 == int32(0) {
				v63 = v15
			} else {
				v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+8)))
				if v20 != 0 {
					base.MemoryCopy(m, v15, v8+int32(10), v20)
				} else {
				}
				v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+8)))
				v25 = v15 + v24
				v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
				if base.Ui32(v26) < base.Ui32(int32(2)) {
					v63 = v25
				} else {
					v37 = v8 + int32(8) + (v24+int32(9))&int32(_a_F_ltree_out_0)
					v38 = v25
					v42 = int32(1)
					for {
						v43 = int32(46)
						*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v43)
						v46 = v38 + int32(1)
						v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37))))
						if v47 != 0 {
							base.MemoryCopy(m, v46, v37+int32(2), v47)
						} else {
						}
						v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37))))
						v52 = v46 + v51
						v59 = v42 + int32(1)
						v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
						if base.Ui32(v59) < base.Ui32(v60) {
							v37 = v37 + (v51+int32(9))&int32(_a_F_ltree_out_0)
							v38 = v52
							v42 = v59
							continue
						} else {
							break
						}
						break
					}
					v63 = v52
				}
			}
			v68 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v68)
			return base.I64_extend_i32_u(v15)
		}
	}
}
func F_ltree_recv(m *base.Module, l0 int32) int64 {
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pq_getmsgint(m, v8, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		if v10 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
				F_errmsg_internal(m, int32(_a_F_ltree_recv_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_ltree_recv_1), int32(239), int32(_a_F_ltree_recv_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			v34 = F_pq_getmsgtext(m, v8, v29-v30, v6+int32(12))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				v37 = F_parse_ltree(m, v34, int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int64(0)
				} else {
					F_pfree(m, v34)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						m.G0 = v6 + int32(16)
						return base.I64_extend_i32_u(v37)
					}
				}
			}
		}
	}
}
