package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F___toread(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v4 - int32(1) | v4
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v9 != v10 {
		v12 = int32(0)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v15 = m.T0[v14].(func(*base.Module, int32, int32, int32) int32)(m, l0, v12, v12)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v23&int32(4) != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v23 | int32(32)
				return int32(-1)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v33 = v31 + v32
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v33
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33
				return v23 << (uint(int32(27)) % 32) >> (uint(int32(31)) % 32)
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v23&int32(4) != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v23 | int32(32)
			return int32(-1)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v33 = v31 + v32
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v33
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33
			return v23 << (uint(int32(27)) % 32) >> (uint(int32(31)) % 32)
		}
	}
}
func F_t_isalnum_cstr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = F_pg_mblen_cstr(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = F_pg_mb2wchar_with_len(m, l0, v5+int32(8), v9)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			v16 = F_pg_database_locale(m)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = F_pg_iswalnum(m, v15, v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(16)
					return v18
				}
			}
		}
	}
}
func F_t_isalnum_with_len(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_pg_mb2wchar_with_len(m, l0, v6+int32(8), l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		v15 = F_pg_database_locale(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_pg_iswalnum(m, v14, v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v17
			}
		}
	}
}
func F_tconvert(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_from_text(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_terminate_brin_buildstate(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v4 != 0 {
		if v4 < int32(0) {
			v8 = *(*int32)(unsafe.Add(mBase, _c_F_terminate_brin_buildstate[0]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8+(v4^int32(-1))<<(uint(int32(2))%32))))
			v22 = v14
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_terminate_brin_buildstate[1]))
			v22 = v16 + v4<<(uint(int32(13))%32) + int32(-8192)
		}
		v23 = int32(4)
		v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+14)))
		v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+12)))
		v26 = v24 - v25
		if v26 <= v23 {
			v29 = v23
		} else {
			v29 = v26
		}
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v32 < int32(0) {
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_terminate_brin_buildstate[2]))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+(v32^int32(-1))*int32(56))+16))
			v51 = v42
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_terminate_brin_buildstate[3]))
			v45 = int32(56)
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+v32*v45-v45)+16))
			v51 = v50
		}
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		F_ReleaseBuffer(m, v52)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_RecordPageWithFreeSpace(m, v55, v51, v29-int32(4))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				F_FreeSpaceMapVacuumRange(m, v58, v51, v51+int32(1))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
					F_MemoryContextDelete(m, v66)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						F_pfree(m, v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	} else {
		v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
		F_MemoryContextDelete(m, v66)
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return
		} else {
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			F_pfree(m, v69)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_textarray_to_strvaluelist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_deconstruct_array_builtin(m, l0, int32(25), v6+int32(12), v6+int32(8), v6+int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v19 < v20 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L14
	}
L4:
	;
	v23 = v19
	v25 = v2
	goto L7
L5:
	;
	v48 = v2
	goto L6
L6:
	;
	m.G0 = v6 + int32(16)
	return v48
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v23))))
	if v28 == int32(1) {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v48 = v40
	goto L6
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v23<<(uint(int32(3))%32))))
	v36 = F_text_to_cstring(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v38 = F_makeString(m, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v40 = F_lappend(m, v25, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v43 = v23 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v43 < v44 {
		v23 = v43
		v25 = v40
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L8
L14:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errmsg(m, int32(_a_F_textarray_to_strvaluelist_0), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_textarray_to_strvaluelist_1), int32(2145), int32(_a_F_textarray_to_strvaluelist_2))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_texteq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = F_pg_newlocale_from_collation(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L82
	}
L4:
	;
	return int64(0)
L5:
	;
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v16 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return base.I64_extend_i32_u(v213)
L7:
	;
	F_pfree(m, v207)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L81
	}
L8:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = F_toast_raw_datum_size(m, v15)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v118 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v15))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L45
	}
L11:
	;
	v22 = F_toast_raw_datum_size(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v20 != v22 {
		v213 = int32(0)
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v26 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v15))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v29 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v19))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v31 = int32(1)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v33&v31 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v36 = v31
	goto L18
L17:
	;
	v36 = int32(4)
	goto L18
L18:
	;
	v37 = v26 + v36
	v38 = int32(1)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v40&v38 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v43 = v38
	goto L21
L20:
	;
	v43 = int32(4)
	goto L21
L21:
	;
	v44 = v29 + v43
	v45 = int32(4)
	v46 = v20 - v45
	if base.Ui32(v45) <= base.Ui32(v46) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v109 != v26 {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	v108 = int32(0)
	goto L22
L24:
	;
	v82 = v77
	v83 = v78
	v84 = v79
	goto L34
L25:
	;
	if (v37|v44)&int32(3) != 0 {
		v77 = v37
		v78 = v44
		v79 = v46
		goto L24
	} else {
		goto L28
	}
L26:
	;
	v70 = v37
	v71 = v44
	v72 = v46
	goto L27
L27:
	;
	if v72 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v54 = v37
	v55 = v44
	v56 = v46
	goto L29
L29:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v59 != v60 {
		v77 = v54
		v78 = v55
		v79 = v56
		goto L24
	} else {
		goto L31
	}
L30:
	;
	v70 = v65
	v71 = v63
	v72 = v67
	goto L27
L31:
	;
	v62 = int32(4)
	v63 = v55 + v62
	v65 = v54 + v62
	v67 = v56 - v62
	if base.Ui32(int32(3)) < base.Ui32(v67) {
		v54 = v65
		v55 = v63
		v56 = v67
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v77 = v70
	v78 = v71
	v79 = v72
	goto L24
L34:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v87 == v88 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v108 = v87 - v88
	goto L22
L36:
	;
	v90 = int32(1)
	v95 = v84 - v90
	if v95 != 0 {
		v82 = v82 + v90
		v83 = v83 + v90
		v84 = v95
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L35
L39:
	;
	goto L23
L40:
	;
	F_pfree(m, v26)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v114 = base.B2i32(v108 == int32(0))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v29 != v115 {
		v205 = v114
		v207 = v29
		goto L7
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v213 = v114
	goto L6
L45:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v121 = F_pg_detoast_datum_packed(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v124 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v154 = int32(1)
	if v124&v154 != 0 {
		goto L58
	} else {
		goto L59
	}
L48:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	if v130 == int32(18) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v141 = int32(1)
	if v124&v141 != 0 {
		v153 = int32(base.Ui32(v124)>>(uint(v141)%32)) - v141
		goto L47
	} else {
		goto L57
	}
L51:
	;
	v133 = int32(16)
	goto L53
L52:
	;
	v133 = int32(0)
	goto L53
L53:
	;
	if base.Ui32((v130-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v140 = int32(4)
	goto L56
L55:
	;
	v140 = v133
	goto L56
L56:
	;
	v153 = v140
	goto L47
L57:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v153 = int32(base.Ui32(v147)>>(uint(int32(2))%32)) - int32(4)
	goto L47
L58:
	;
	v158 = v154
	goto L60
L59:
	;
	v158 = int32(4)
	goto L60
L60:
	;
	v160 = int32(1)
	if v123&v160 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v164 = v160
	goto L63
L62:
	;
	v164 = int32(4)
	goto L63
L63:
	;
	if v123 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v195 = F_varstr_cmp(m, v118+v158, v153, v121+v164, v194, v10)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L75
	}
L65:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)))
	if v171 == int32(18) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v182 = int32(1)
	if v123&v182 != 0 {
		v194 = int32(base.Ui32(v123)>>(uint(v182)%32)) - v182
		goto L64
	} else {
		goto L74
	}
L68:
	;
	v174 = int32(16)
	goto L70
L69:
	;
	v174 = int32(0)
	goto L70
L70:
	;
	if base.Ui32((v171-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v181 = int32(4)
	goto L73
L72:
	;
	v181 = v174
	goto L73
L73:
	;
	v194 = v181
	goto L64
L74:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v194 = int32(base.Ui32(v188)>>(uint(int32(2))%32)) - int32(4)
	goto L64
L75:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v197 != v118 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_pfree(m, v118)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v202 = base.B2i32(v195 == int32(0))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v121 == v203 {
		v213 = v202
		goto L6
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v205 = v202
	v207 = v121
	goto L7
L81:
	;
	v213 = v205
	goto L6
L82:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_texteq_0), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errhint(m, int32(_a_F_texteq_1), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_texteq_2), int32(1337), int32(_a_F_texteq_3))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_texteqfast(m *base.Module, l0 int64, l1 int64) int32 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_DirectFunctionCall2Coll(m, int32(1757), int32(950), l0, l1)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v5 != int64(0))
	}
}
func F_textgtname(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_DirectFunctionCall2Coll(m, int32(1756), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) < base.I32_wrap_i64(v6)))
	}
}
func F_texthashfast(m *base.Module, l0 int64) int32 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_DirectFunctionCall1Coll(m, int32(1783), int32(950), l0)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v4)
	}
}
func F_texticregexne(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14379(m, l0, int32(27))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_textoverlay_no_len(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = F_pg_detoast_datum_packed(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v13 = F_text_length(m, base.I64_extend_i32_u(v9))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = F_text_overlay(m, v4, v9, v11, v13)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v15)
				}
			}
		}
	}
}
func F_textpos(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int64
	_ = v138
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v8 = m.G0
	v10 = v8 - int32(1072)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v20 != 0 {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				if v21 == int32(1) {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
					if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						if v51 == int32(1) {
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
							if v57 == int32(18) {
								v60 = int32(16)
							} else {
								v60 = int32(0)
							}
							if base.Ui32((v57-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v67 = int32(4)
							} else {
								v67 = v60
							}
							v80 = v67
						} else {
							v68 = int32(1)
							if v51&v68 != 0 {
								v80 = int32(base.Ui32(v51)>>(uint(v68)%32)) - v68
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								v80 = int32(base.Ui32(v74)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if v21 == int32(1) {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
							if v86 == int32(18) {
								v89 = int32(16)
							} else {
								v89 = int32(0)
							}
							if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v96 = int32(4)
							} else {
								v96 = v89
							}
							v109 = v96
						} else {
							v97 = int32(1)
							if v21&v97 != 0 {
								v109 = int32(base.Ui32(v21)>>(uint(v97)%32)) - v97
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
								v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if base.Ui32(v109) <= base.Ui32(v80) {
							F_text_position_setup(m, v13, v18, v20, v10)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int64(0)
							} else {
								v119 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+5)) = uint8(v119)
								v122 = F_text_position_next(m, v10)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int64(0)
								} else {
									if v122 == int32(0) {
										v138 = int64(0)
										m.G0 = v10 + int32(1072)
										return v138
									} else {
										v126 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1064))
										v127 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1052))
										v129 = F_pg_mbstrlen_with_len(m, v126, v127-v126)
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int64(0)
										} else {
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1068))
											v138 = base.I64_extend_i32_s(v129 + v131 + int32(1))
											m.G0 = v10 + int32(1072)
											return v138
										}
									}
								}
							}
						} else {
							v111 = F_pg_newlocale_from_collation(m, v20)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int64(0)
							} else {
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
								if v113 == int32(0) {
									F_text_position_setup(m, v13, v18, v20, v10)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int64(0)
									} else {
										v119 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+5)) = uint8(v119)
										v122 = F_text_position_next(m, v10)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int64(0)
										} else {
											if v122 == int32(0) {
												v138 = int64(0)
												m.G0 = v10 + int32(1072)
												return v138
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1064))
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1052))
												v129 = F_pg_mbstrlen_with_len(m, v126, v127-v126)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return int64(0)
												} else {
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1068))
													v138 = base.I64_extend_i32_s(v129 + v131 + int32(1))
													m.G0 = v10 + int32(1072)
													return v138
												}
											}
										}
									}
								} else {
									v138 = int64(0)
									m.G0 = v10 + int32(1072)
									return v138
								}
							}
						}
					} else {
						if v24 == int32(18) {
							v35 = int32(16)
						} else {
							v35 = int32(0)
						}
						v48 = v35
						if v48 != 0 {
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
							if v51 == int32(1) {
								v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
								if v57 == int32(18) {
									v60 = int32(16)
								} else {
									v60 = int32(0)
								}
								if base.Ui32((v57-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v67 = int32(4)
								} else {
									v67 = v60
								}
								v80 = v67
							} else {
								v68 = int32(1)
								if v51&v68 != 0 {
									v80 = int32(base.Ui32(v51)>>(uint(v68)%32)) - v68
								} else {
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
									v80 = int32(base.Ui32(v74)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							if v21 == int32(1) {
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
								if v86 == int32(18) {
									v89 = int32(16)
								} else {
									v89 = int32(0)
								}
								if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v96 = int32(4)
								} else {
									v96 = v89
								}
								v109 = v96
							} else {
								v97 = int32(1)
								if v21&v97 != 0 {
									v109 = int32(base.Ui32(v21)>>(uint(v97)%32)) - v97
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
									v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							if base.Ui32(v109) <= base.Ui32(v80) {
								F_text_position_setup(m, v13, v18, v20, v10)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int64(0)
								} else {
									v119 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+5)) = uint8(v119)
									v122 = F_text_position_next(m, v10)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int64(0)
									} else {
										if v122 == int32(0) {
											v138 = int64(0)
											m.G0 = v10 + int32(1072)
											return v138
										} else {
											v126 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1064))
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1052))
											v129 = F_pg_mbstrlen_with_len(m, v126, v127-v126)
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int64(0)
											} else {
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1068))
												v138 = base.I64_extend_i32_s(v129 + v131 + int32(1))
												m.G0 = v10 + int32(1072)
												return v138
											}
										}
									}
								}
							} else {
								v111 = F_pg_newlocale_from_collation(m, v20)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int64(0)
								} else {
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
									if v113 == int32(0) {
										F_text_position_setup(m, v13, v18, v20, v10)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int64(0)
										} else {
											v119 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+5)) = uint8(v119)
											v122 = F_text_position_next(m, v10)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int64(0)
											} else {
												if v122 == int32(0) {
													v138 = int64(0)
													m.G0 = v10 + int32(1072)
													return v138
												} else {
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1064))
													v127 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1052))
													v129 = F_pg_mbstrlen_with_len(m, v126, v127-v126)
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int64(0)
													} else {
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1068))
														v138 = base.I64_extend_i32_s(v129 + v131 + int32(1))
														m.G0 = v10 + int32(1072)
														return v138
													}
												}
											}
										}
									} else {
										v138 = int64(0)
										m.G0 = v10 + int32(1072)
										return v138
									}
								}
							}
						} else {
							v138 = int64(1)
							m.G0 = v10 + int32(1072)
							return v138
						}
					}
				} else {
					v36 = int32(1)
					if v21&v36 != 0 {
						v48 = int32(base.Ui32(v21)>>(uint(v36)%32)) - v36
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
						v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
					}
					if v48 != 0 {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						if v51 == int32(1) {
							v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
							if v57 == int32(18) {
								v60 = int32(16)
							} else {
								v60 = int32(0)
							}
							if base.Ui32((v57-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v67 = int32(4)
							} else {
								v67 = v60
							}
							v80 = v67
						} else {
							v68 = int32(1)
							if v51&v68 != 0 {
								v80 = int32(base.Ui32(v51)>>(uint(v68)%32)) - v68
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								v80 = int32(base.Ui32(v74)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if v21 == int32(1) {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
							if v86 == int32(18) {
								v89 = int32(16)
							} else {
								v89 = int32(0)
							}
							if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v96 = int32(4)
							} else {
								v96 = v89
							}
							v109 = v96
						} else {
							v97 = int32(1)
							if v21&v97 != 0 {
								v109 = int32(base.Ui32(v21)>>(uint(v97)%32)) - v97
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
								v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if base.Ui32(v109) <= base.Ui32(v80) {
							F_text_position_setup(m, v13, v18, v20, v10)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int64(0)
							} else {
								v119 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+5)) = uint8(v119)
								v122 = F_text_position_next(m, v10)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int64(0)
								} else {
									if v122 == int32(0) {
										v138 = int64(0)
										m.G0 = v10 + int32(1072)
										return v138
									} else {
										v126 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1064))
										v127 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1052))
										v129 = F_pg_mbstrlen_with_len(m, v126, v127-v126)
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int64(0)
										} else {
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1068))
											v138 = base.I64_extend_i32_s(v129 + v131 + int32(1))
											m.G0 = v10 + int32(1072)
											return v138
										}
									}
								}
							}
						} else {
							v111 = F_pg_newlocale_from_collation(m, v20)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int64(0)
							} else {
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
								if v113 == int32(0) {
									F_text_position_setup(m, v13, v18, v20, v10)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int64(0)
									} else {
										v119 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+5)) = uint8(v119)
										v122 = F_text_position_next(m, v10)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int64(0)
										} else {
											if v122 == int32(0) {
												v138 = int64(0)
												m.G0 = v10 + int32(1072)
												return v138
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1064))
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1052))
												v129 = F_pg_mbstrlen_with_len(m, v126, v127-v126)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return int64(0)
												} else {
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v10)+1068))
													v138 = base.I64_extend_i32_s(v129 + v131 + int32(1))
													m.G0 = v10 + int32(1072)
													return v138
												}
											}
										}
									}
								} else {
									v138 = int64(0)
									m.G0 = v10 + int32(1072)
									return v138
								}
							}
						}
					} else {
						v138 = int64(1)
						m.G0 = v10 + int32(1072)
						return v138
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(34209924))
					mBase = m.M
					v149 = m.ExcPending
					if v149 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_textpos_0), int32(0))
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return int64(0)
						} else {
							F_errhint(m, int32(_a_F_textpos_1), int32(0))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_textpos_2), int32(1337), int32(_a_F_textpos_3))
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
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
			}
		}
	}
}
func F_textrecv(m *base.Module, l0 int32) int64 {
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
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
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
	var v32 int32
	_ = v32
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v16 = F_pq_getmsgtext(m, v10, v11-v12, v8+int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v22 = v20 + int32(4)
		v23 = F_palloc(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = v22 << (uint(int32(2)) % 32)
			if v20 != 0 {
				base.MemoryCopy(m, v23+int32(4), v16, v20)
			} else {
			}
			F_pfree(m, v16)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int64(0)
			} else {
				m.G0 = v8 + int32(16)
				return base.I64_extend_i32_u(v23)
			}
		}
	}
}
func F_thesaurus_init(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
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
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v499 int32
	_ = v499
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v615 int64
	_ = v615
	var v617 int32
	_ = v617
	var v620 int64
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var __phi690 int32
	_ = __phi690
	var v691 int32
	_ = v691
	var __phi691 int32
	_ = __phi691
	var v693 int32
	_ = v693
	var __phi693 int32
	_ = __phi693
	var v695 int32
	_ = v695
	var __phi695 int32
	_ = __phi695
	var v697 int32
	_ = v697
	var __phi697 int32
	_ = __phi697
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var __phi824 int32
	_ = __phi824
	var v826 int32
	_ = v826
	var __phi826 int32
	_ = __phi826
	var v827 int32
	_ = v827
	var __phi827 int32
	_ = __phi827
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 int64
	_ = v900
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v957 int32
	_ = v957
	var v970 int32
	_ = v970
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v999 int32
	_ = v999
	var __phi999 int32
	_ = __phi999
	var v1001 int32
	_ = v1001
	var __phi1001 int32
	_ = __phi1001
	var v1002 int32
	_ = v1002
	var __phi1002 int32
	_ = __phi1002
	var v1006 int32
	_ = v1006
	var __phi1006 int32
	_ = __phi1006
	var v1008 int32
	_ = v1008
	var __phi1008 int32
	_ = __phi1008
	var v1013 int32
	_ = v1013
	var v1016 int64
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1024 int32
	_ = v1024
	var v1028 int64
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1033 int64
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var __phi1048 int32
	_ = __phi1048
	var v1051 int32
	_ = v1051
	var __phi1051 int32
	_ = __phi1051
	var v1052 int32
	_ = v1052
	var __phi1052 int32
	_ = __phi1052
	var v1059 int32
	_ = v1059
	var __phi1059 int32
	_ = __phi1059
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int64
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1214 int32
	_ = v1214
	var v1219 int32
	_ = v1219
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1260 int32
	_ = v1260
	var v1264 int32
	_ = v1264
	var v1269 int32
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1292 int32
	_ = v1292
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1313 int32
	_ = v1313
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1344 int32
	_ = v1344
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(160)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = F_palloc0(m, int32(28))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v20 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L1
	} else {
		goto L317
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v28 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v34 = v2
	v43 = v2
	v44 = v2
	goto L6
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v44<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v52 = int32(_a_F_thesaurus_init_0)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_thesaurus_init[0])))
	if base.B2i32(v55 == int32(0))|base.B2i32(v55 != v58) != 0 {
		v76 = v55
		v77 = v58
		goto L11
	} else {
		goto L12
	}
L7:
	;
	if v515&int32(1) == int32(0) {
		goto L3
	} else {
		goto L146
	}
L8:
	;
	v528 = v44 + int32(1)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v528 < v529 {
		v34 = v515
		v43 = v524
		v44 = v528
		goto L6
	} else {
		goto L145
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v499
	F_tsearch_readline_end(m, v18+int32(112))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L143
	}
L10:
	;
	if v76-v77 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	v61 = v51
	v62 = v52
	goto L13
L13:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v66 == int32(0) {
		v76 = v66
		v77 = v65
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v76 = v66
	v77 = v65
	goto L11
L15:
	;
	v69 = int32(1)
	if v66 == v65 {
		v61 = v61 + v69
		v62 = v62 + v69
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	if v34&int32(1) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v419 = int32(_a_F_thesaurus_init_1)
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v425 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_thesaurus_init[1])))
	if base.B2i32(v422 == int32(0))|base.B2i32(v422 != v425) != 0 {
		v443 = v422
		v444 = v425
		goto L121
	} else {
		goto L122
	}
L20:
	;
	v86 = v18 + int32(112)
	v87 = F_defGetString(m, v50)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L116
	}
L23:
	;
	v90 = F_get_tsearch_config_filename(m, v87, int32(_a_F_thesaurus_init_2))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v92 = F_tsearch_readline_begin(m, v86, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v92 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v94 = int32(0)
	v96 = F_tsearch_readline(m, v86)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L112
	}
L29:
	;
	if v96 == int32(0) {
		v499 = v94
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v101 = v96
	v108 = v94
	v110 = v94
	goto L31
L31:
	;
	v115 = v101
	goto L33
L33:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if base.Ui32(v130-int32(9)) < base.Ui32(int32(5)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v382 = F_pg_mblen_cstr(m, v115)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L111
	}
L36:
	;
	v136 = int32(0)
	switch v130 - int32(32) {
	case 0:
		goto L35
	case 1, 2:
		goto L41
	case 3:
		v317 = v108
		v319 = v110
		goto L40
	default:
		goto L42
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L107
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L103
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L99
	}
L40:
	;
	F_pfree(m, v101)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L96
	}
L41:
	;
	v143 = v115
	v145 = v136
	v146 = int32(1)
	v147 = v130
	v149 = v136
	v151 = v108
	v154 = v136
	goto L44
L42:
	;
	if v130 == int32(0) {
		v317 = v108
		v319 = v110
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	switch v146 - int32(2) {
	case 0:
		goto L52
	case 1:
		goto L51
	case 2:
		goto L50
	default:
		goto L53
	}
L45:
	;
	if v274 == int32(4) {
		goto L89
	} else {
		goto L90
	}
L46:
	;
	v279 = F_pg_mblen_cstr(m, v143)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L87
	}
L47:
	;
	v268 = F_pg_mblen_cstr(m, v143)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L86
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L82
	}
L49:
	;
	v273 = v248
	v274 = int32(3)
	v276 = v250
	v277 = v151
	v278 = v154
	goto L46
L50:
	;
	v226 = v147 & int32(255)
	if base.B2i32(base.Ui32(v226-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v226 == int32(32)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L51:
	;
	v213 = v147 & int32(255)
	switch v213 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		v273 = v145
		v274 = int32(3)
		v276 = v149
		v277 = v151
		v278 = v154
		goto L46
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 26, 27, 28, 29, 30, 31, 32:
		goto L72
	case 33:
		goto L74
	default:
		goto L73
	}
L52:
	;
	switch v147&int32(255) - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		goto L68
	default:
		v273 = v145
		v274 = int32(2)
		v276 = v149
		v277 = v151
		v278 = v154
		goto L46
	case 49:
		goto L69
	}
L53:
	;
	v161 = v147 & int32(255)
	if v161 == int32(58) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if v145 != 0 {
		v248 = v145
		v250 = v149
		goto L49
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v188 = base.B2i32(v161 != int32(32)) & base.B2i32(base.Ui32((v147-int32(14))&int32(255)) < base.Ui32(int32(251)))
	if v188 != 0 {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_thesaurus_init_3), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(212), int32(_a_F_thesaurus_init_5))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v189 = v143
	goto L64
L63:
	;
	v189 = v154
	goto L64
L64:
	;
	if v188 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v192 = int32(2)
	goto L67
L66:
	;
	v192 = int32(1)
	goto L67
L67:
	;
	v273 = v145
	v274 = v192
	v276 = v149
	v277 = v151
	v278 = v189
	goto L46
L68:
	;
	F_newLexeme(m, v22, v154, v143, v110, v145&int32(_a_F_thesaurus_init_6))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	F_newLexeme(m, v22, v154, v143, v110, v145&int32(_a_F_thesaurus_init_6))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v248 = v145 + int32(1)
	v250 = v149
	goto L49
L71:
	;
	v208 = int32(1)
	v273 = v145 + v208
	v274 = v208
	v276 = v149
	v277 = v151
	v278 = v154
	goto L46
L72:
	;
	v273 = v145
	v274 = int32(4)
	v276 = v149
	v277 = int32(0)
	v278 = v143
	goto L46
L73:
	;
	if v213 == int32(92) {
		goto L47
	} else {
		goto L76
	}
L74:
	;
	v216 = F_pg_mblen_cstr(m, v143)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v273 = v145
	v274 = int32(4)
	v276 = v149
	v277 = int32(1)
	v278 = v216 + v143
	goto L46
L76:
	;
	goto L72
L77:
	;
	v273 = v145
	v274 = int32(4)
	v276 = v149
	v277 = v151
	v278 = v154
	goto L46
L78:
	;
	goto L79
L79:
	;
	if v143 == v154 {
		goto L48
	} else {
		goto L80
	}
L80:
	;
	v238 = int32(_a_F_thesaurus_init_6)
	F_addWrd(m, v22, v154, v143, v110, v149&v238, v145&v238, v151&int32(1))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v248 = v145
	v250 = v149 + int32(1)
	goto L49
L82:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_thesaurus_init_7), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(262), int32(_a_F_thesaurus_init_5))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v273 = v145
	v274 = int32(4)
	v276 = v149
	v277 = int32(0)
	v278 = v268 + v143
	goto L46
L87:
	;
	v281 = v279 + v143
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v282 != 0 {
		v143 = v281
		v145 = v273
		v146 = v274
		v147 = v282
		v149 = v276
		v151 = v277
		v154 = v278
		goto L44
	} else {
		goto L88
	}
L88:
	;
	goto L45
L89:
	;
	if v281 == v278 {
		goto L39
	} else {
		goto L92
	}
L90:
	;
	v296 = v276
	goto L91
L91:
	;
	v297 = int32(0)
	if base.B2i32(v296 == v297)|base.B2i32(v273 == v297) != 0 {
		goto L38
	} else {
		goto L94
	}
L92:
	;
	v286 = int32(_a_F_thesaurus_init_6)
	F_addWrd(m, v22, v278, v281, v110, v276&v286, v273&v286, v277&int32(1))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v296 = v276 + int32(1)
	goto L91
L94:
	;
	if base.B2i32(base.Ui32(int32(_a_F_thesaurus_init_6)) < base.Ui32(v296))|base.B2i32(base.Ui32(int32(_a_F_thesaurus_init_8)) <= base.Ui32(v273)) != 0 {
		goto L37
	} else {
		goto L95
	}
L95:
	;
	v317 = v277
	v319 = v110 + int32(1)
	goto L40
L96:
	;
	v328 = F_tsearch_readline(m, v18+int32(112))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v328 != 0 {
		v101 = v328
		v108 = v317
		v110 = v319
		goto L31
	} else {
		goto L98
	}
L98:
	;
	v499 = v319
	goto L9
L99:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errmsg(m, int32(_a_F_thesaurus_init_7), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(278), int32(_a_F_thesaurus_init_5))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errmsg(m, int32(_a_F_thesaurus_init_9), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(287), int32(_a_F_thesaurus_init_5))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errmsg(m, int32(_a_F_thesaurus_init_10), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(292), int32(_a_F_thesaurus_init_5))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	v115 = v382 + v115
	goto L33
L112:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v90
	F_errmsg(m, int32(_a_F_thesaurus_init_11), v18+int32(80))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(180), int32(_a_F_thesaurus_init_5))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errmsg(m, int32(_a_F_thesaurus_init_12), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(617), int32(_a_F_thesaurus_init_13))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	if v443-v444 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L121:
	;
	goto L120
L122:
	;
	v428 = v51
	v429 = v419
	goto L123
L123:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+1)))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428)+1)))
	if v433 == int32(0) {
		v443 = v433
		v444 = v432
		goto L121
	} else {
		goto L125
	}
L124:
	;
	v443 = v433
	v444 = v432
	goto L121
L125:
	;
	v436 = int32(1)
	if v433 == v432 {
		v428 = v428 + v436
		v429 = v429 + v436
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	if v43 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L139
	}
L130:
	;
	v450 = F_defGetString(m, v50)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	v452 = F_pstrdup(m, v450)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v515 = v34
	v524 = v452
	goto L8
L135:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_thesaurus_init_14), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(626), int32(_a_F_thesaurus_init_13))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v477
	F_errmsg(m, int32(_a_F_thesaurus_init_15), v18+int32(96))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(634), int32(_a_F_thesaurus_init_13))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_pfree(m, v90)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v515 = int32(1)
	v524 = v43
	goto L8
L145:
	;
	goto L7
L146:
	;
	if v524 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v535 = int32(0)
	v537 = F_stringToQualifiedNameList(m, v524, v535)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L1
	} else {
		goto L313
	}
L150:
	;
	v540 = F_get_ts_dict_oid(m, v537, int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v540
	v543 = F_lookup_ts_dictionary_cache(m, v540)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v543
	v546 = int32(16)
	v549 = F_palloc_mul(m, int32(8), v546)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if int32(0) < v551 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L1
	} else {
		goto L309
	}
L155:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L1
	} else {
		goto L304
	}
L156:
	;
	v558 = v535
	v561 = v546
	v563 = v549
	v564 = int32(0)
	goto L159
L157:
	;
	v794 = v535
	v797 = v546
	v799 = v549
	goto L158
L158:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v806 != 0 {
		goto L202
	} else {
		goto L203
	}
L159:
	;
	v571 = v564 << (uint(int32(3)) % 32)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v573 = v571 + v572
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	if v575 != int32(63) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	v794 = v765
	v797 = v768
	v799 = v770
	goto L158
L161:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v777+v571)))
	F_pfree(m, v779)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L199
	}
L162:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v615 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v611)+44)))
	v617 = F_strlen(m, v574)
	mBase = m.M
	v620 = F_FunctionCall4Coll(m, v611+int32(12), int32(0), v615, base.I64_extend_i32_u(v574), base.I64_extend_i32_s(v617), int64(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L170
	}
L163:
	;
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+1)))
	if v578 != 0 {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v573)+4))
	if v561 <= v558 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v583 = F_repalloc(m, v563, v561<<(uint(int32(4))%32))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	v587 = v561
	v588 = v563
	goto L167
L167:
	;
	v590 = F_palloc(m, int32(16))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L169
	}
L168:
	;
	v587 = v561 << (uint(int32(1)) % 32)
	v588 = v583
	goto L167
L169:
	;
	v594 = v588 + v558<<(uint(int32(3))%32)
	v595 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v594))) = v595
	*(*int32)(unsafe.Add(mBase, uint32(v594)+4)) = v590
	v598 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v590)+6)) = uint16(v598)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v594)+4))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	*(*int32)(unsafe.Add(mBase, uint32(v600))) = v601
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v594)+4))
	v604 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v579)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v603)+4)) = uint16(v604)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v594)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v606)+8)) = v595
	v765 = v558 + v598
	v768 = v587
	v770 = v588
	goto L161
L170:
	;
	v622 = base.I32_wrap_i64(v620)
	if v622 == int32(0) {
		goto L154
	} else {
		goto L171
	}
L171:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v622)+4))
	if v625 == int32(0) {
		goto L155
	} else {
		goto L172
	}
L172:
	;
	v630 = v622
	v631 = v558
	v634 = v561
	v636 = v563
	goto L173
L173:
	;
	v643 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v630))))
	v644 = int32(1)
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v630)+12))
	if v645 == int32(0) {
		v683 = v644
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v765 = v754
	v768 = v756
	v770 = v757
	goto L161
L175:
	;
	__phi690 = v631
	__phi691 = v630
	__phi693 = v634
	__phi695 = v636
	__phi697 = v630 + int32(4)
	v690 = __phi690
	v691 = __phi691
	v693 = __phi693
	v695 = __phi695
	v697 = __phi697
	goto L181
L176:
	;
	v648 = v630
	v661 = v644
	goto L177
L177:
	;
	v663 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v648)+8)))
	if v663 != v643 {
		v683 = v661
		goto L175
	} else {
		goto L179
	}
L178:
	;
	v683 = v666
	goto L175
L179:
	;
	v666 = v661 + int32(1)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v648)+20))
	if v667 != 0 {
		v648 = v648 + int32(8)
		v661 = v666
		goto L177
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v702 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v691))))
	if v643 != v702 {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v753)+4))
	if v761 != 0 {
		v630 = v753
		v631 = v754
		v634 = v756
		v636 = v757
		goto L173
	} else {
		goto L198
	}
L183:
	;
	goto L182
L184:
	;
	v753 = v691
	v754 = v690
	v756 = v693
	v757 = v695
	goto L183
L185:
	;
	goto L186
L186:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v704+v571)+4))
	if v693 <= v690 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v710 = F_repalloc(m, v695, v693<<(uint(int32(4))%32))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L190
	}
L188:
	;
	v714 = v693
	v715 = v695
	goto L189
L189:
	;
	v718 = v715 + v690<<(uint(int32(3))%32)
	v720 = F_palloc(m, int32(16))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L191
	}
L190:
	;
	v714 = v693 << (uint(int32(1)) % 32)
	v715 = v710
	goto L189
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v718)+4)) = v720
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v697)))
	if v723 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = v731
	*(*uint16)(unsafe.Add(mBase, uint32(v732)+6)) = uint16(v733)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v718)+4))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	*(*int32)(unsafe.Add(mBase, uint32(v736))) = v737
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v718)+4))
	v740 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v706)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v739)+4)) = uint16(v740)
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v718)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v742)+8)) = int32(0)
	v748 = v690 + int32(1)
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v691)+12))
	v751 = v691 + int32(8)
	if v749 != 0 {
		__phi690 = v748
		__phi691 = v751
		__phi693 = v714
		__phi695 = v715
		__phi697 = v691 + int32(12)
		v690 = __phi690
		v691 = __phi691
		v693 = __phi693
		v695 = __phi695
		v697 = __phi697
		goto L181
	} else {
		goto L197
	}
L193:
	;
	v731 = int32(0)
	v732 = v720
	v733 = int32(1)
	goto L192
L194:
	;
	goto L195
L195:
	;
	v728 = F_pstrdup(m, v723)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v718)+4))
	v731 = v728
	v732 = v730
	v733 = v683
	goto L192
L197:
	;
	v753 = v751
	v754 = v748
	v756 = v714
	v757 = v715
	goto L183
L198:
	;
	goto L174
L199:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v782+v571)+4))
	F_pfree(m, v784)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v788 = v564 + int32(1)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v788 < v789 {
		v558 = v765
		v561 = v768
		v563 = v770
		v564 = v788
		goto L159
	} else {
		goto L201
	}
L201:
	;
	goto L160
L202:
	;
	F_pfree(m, v806)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v794
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v799
	if int32(2) <= v794 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	goto L204
L206:
	;
	F_pg_qsort(m, v799, v794, int32(8), int32(1264))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if int32(0) < v957 {
		goto L248
	} else {
		goto L249
	}
L209:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v819 < int32(2) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v932 = int32(3)
	v935 = (v916-v920)>>(uint(v932)%32) + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v935
	v939 = F_repalloc(m, v920, v935<<(uint(v932)%32))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L1
	} else {
		goto L244
	}
L211:
	;
	v916 = v818
	v920 = v818
	goto L210
L212:
	;
	goto L213
L213:
	;
	__phi824 = v818
	__phi826 = v818 + int32(8)
	__phi827 = v818
	v824 = __phi824
	v826 = __phi826
	v827 = __phi827
	goto L214
L214:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v824)))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v827)+8))
	if v840 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L215:
	;
	v916 = v904
	v920 = v911
	goto L210
L216:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v910 = v826 + int32(8)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if (v910-v911)>>(uint(int32(3))%32) < v908 {
		__phi824 = v904
		__phi826 = v910
		__phi827 = v826
		v824 = __phi824
		v826 = __phi826
		v827 = __phi827
		goto L214
	} else {
		goto L243
	}
L217:
	;
	v900 = *(*int64)(unsafe.Add(mBase, uint32(v827)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v824)+8)) = v900
	v904 = v824 + int32(8)
	goto L216
L218:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v827)+12))
	if v873 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L219:
	;
	if v839 == int32(0) {
		goto L218
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	if v839 == int32(0) {
		goto L217
	} else {
		goto L223
	}
L222:
	;
	goto L217
L223:
	;
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v840))))
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839))))
	if base.B2i32(v849 == int32(0))|base.B2i32(v849 != v852) != 0 {
		v870 = v849
		v871 = v852
		goto L225
	} else {
		goto L226
	}
L224:
	;
	if v870-v871 != 0 {
		goto L217
	} else {
		goto L231
	}
L225:
	;
	goto L224
L226:
	;
	v855 = v840
	v856 = v839
	goto L227
L227:
	;
	v859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v856)+1)))
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v855)+1)))
	if v860 == int32(0) {
		v870 = v860
		v871 = v859
		goto L225
	} else {
		goto L229
	}
L228:
	;
	v870 = v860
	v871 = v859
	goto L225
L229:
	;
	v863 = int32(1)
	if v860 == v859 {
		v855 = v855 + v863
		v856 = v856 + v863
		goto L227
	} else {
		goto L230
	}
L230:
	;
	goto L228
L231:
	;
	goto L218
L232:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v827)+8))
	if v895 == int32(0) {
		v904 = v824
		goto L216
	} else {
		goto L241
	}
L233:
	;
	F_pfree(m, v873)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L1
	} else {
		goto L240
	}
L234:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v824)+4))
	if v876 == int32(0) {
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v873)))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v876)))
	if v879 != v880 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v873)+8)) = v876
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v827)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v824)+4)) = v889
	goto L232
L237:
	;
	v882 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+4)))
	v883 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v876)+4)))
	if v882 != v883 {
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v885 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v873)+6)))
	v886 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v876)+6)))
	if v885 == v886 {
		goto L233
	} else {
		goto L239
	}
L239:
	;
	goto L236
L240:
	;
	goto L232
L241:
	;
	F_pfree(m, v895)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	v904 = v824
	goto L216
L243:
	;
	goto L215
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v939
	goto L208
L245:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L300
	}
L246:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L1
	} else {
		goto L296
	}
L247:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L292
	}
L248:
	;
	v970 = int32(0)
	goto L251
L249:
	;
	goto L250
L250:
	;
	m.G0 = v18 + int32(160)
	return base.I64_extend_i32_u(v22)
L251:
	;
	v977 = v970 << (uint(int32(3)) % 32)
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v977+v978)+4))
	v983 = F_palloc_mul(m, int32(8), int32(2))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L1
	} else {
		goto L253
	}
L252:
	;
	goto L250
L253:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v985+v977)+4)) = v983
	v988 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v983)+4)) = v988
	if v980 == v988 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v1143 = v1142 + v977
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1143)+4))
	if v1127 == v1144 {
		goto L245
	} else {
		goto L289
	}
L255:
	;
	v1127 = v983
	goto L254
L256:
	;
	goto L257
L257:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v980)+4))
	if v992 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1127 = v983
	goto L254
L259:
	;
	goto L260
L260:
	;
	__phi999 = v983
	__phi1001 = int32(2)
	__phi1002 = v992
	__phi1006 = v980
	__phi1008 = v980 + int32(4)
	v999 = __phi999
	v1001 = __phi1001
	v1002 = __phi1002
	v1006 = __phi1006
	v1008 = __phi1008
	goto L261
L261:
	;
	v1013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+3)))
	if v1013&int32(16) != 0 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	v1127 = v1096
	goto L254
L263:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1039)+4))
	if v1040 == int32(0) {
		goto L247
	} else {
		goto L269
	}
L264:
	;
	v1016 = *(*int64)(unsafe.Add(mBase, uint32(v1006)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v1016
	v1018 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v1018
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+114)) = uint16(v1018)
	v1039 = v18 + int32(112)
	goto L263
L265:
	;
	goto L266
L266:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v1028 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1024)+44)))
	v1030 = F_strlen(m, v1002)
	mBase = m.M
	v1033 = F_FunctionCall4Coll(m, v1024+int32(12), int32(0), v1028, base.I64_extend_i32_u(v1002), base.I64_extend_i32_s(v1030), int64(0))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	v1035 = base.I32_wrap_i64(v1033)
	if v1035 == int32(0) {
		goto L246
	} else {
		goto L268
	}
L268:
	;
	v1039 = v1035
	goto L263
L269:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1045+v977)+4))
	__phi1048 = v999
	__phi1051 = v1001
	__phi1052 = v1039
	__phi1059 = v1039 + int32(4)
	v1048 = __phi1048
	v1051 = __phi1051
	v1052 = __phi1052
	v1059 = __phi1059
	goto L270
L270:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1063+v977)+4))
	v1066 = v1048 - v1065
	if v1051 <= v1066>>(uint(int32(3))%32)+int32(1) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	if v999 == v1047 {
		goto L278
	} else {
		goto L279
	}
L272:
	;
	v1074 = F_repalloc(m, v1065, v1051<<(uint(int32(4))%32))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L275
	}
L273:
	;
	v1085 = v1048
	v1086 = v1051
	goto L274
L274:
	;
	v1087 = *(*int64)(unsafe.Add(mBase, uint32(v1052)))
	*(*int64)(unsafe.Add(mBase, uint32(v1085))) = v1087
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1059)))
	v1090 = F_pstrdup(m, v1089)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L1
	} else {
		goto L276
	}
L275:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1076+v977)+4)) = v1074
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1081+v977)+4))
	v1085 = v1083 + v1066
	v1086 = v1051 << (uint(int32(1)) % 32)
	goto L274
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1085)+4)) = v1090
	v1095 = int32(8)
	v1096 = v1085 + v1095
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+12))
	if v1097 != 0 {
		__phi1048 = v1096
		__phi1051 = v1086
		__phi1052 = v1052 + v1095
		__phi1059 = v1052 + int32(12)
		v1048 = __phi1048
		v1051 = __phi1051
		v1052 = __phi1052
		v1059 = __phi1059
		goto L270
	} else {
		goto L277
	}
L277:
	;
	goto L271
L278:
	;
	v1105 = int32(-1)
	goto L280
L279:
	;
	v1105 = (v999 - v1047) >> (uint(int32(3)) % 32)
	goto L280
L280:
	;
	if int32(0) < v1105 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1108+v977)+4))
	v1113 = v1110 + v1105<<(uint(int32(3))%32)
	v1114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1113)+2)))
	v1116 = v1114 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1113)+2)) = uint16(v1116)
	goto L283
L282:
	;
	goto L283
L283:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1008)))
	if v1119 != 0 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	F_pfree(m, v1119)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L1
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+12))
	if v1124 != 0 {
		__phi999 = v1096
		__phi1001 = v1086
		__phi1002 = v1124
		__phi1006 = v1006 + int32(8)
		__phi1008 = v1006 + int32(12)
		v999 = __phi999
		v1001 = __phi1001
		v1002 = __phi1002
		v1006 = __phi1006
		v1008 = __phi1008
		goto L261
	} else {
		goto L288
	}
L287:
	;
	goto L286
L288:
	;
	goto L262
L289:
	;
	v1148 = int32(base.Ui32(v1127-v1144) >> (uint(int32(3)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v1143)+2)) = uint16(v1148)
	F_pfree(m, v980)
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	v1153 = v970 + int32(1)
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v1153 < v1154 {
		v970 = v1153
		goto L251
	} else {
		goto L291
	}
L291:
	;
	goto L252
L292:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1008)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v970 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v1183
	F_errmsg(m, int32(_a_F_thesaurus_init_16), v18+int32(32))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(569), int32(_a_F_thesaurus_init_17))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L295
	}
L295:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L296:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1008)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v970 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1205
	F_errmsg(m, int32(_a_F_thesaurus_init_18), v18+int32(16))
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(576), int32(_a_F_thesaurus_init_17))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L1
	} else {
		goto L299
	}
L299:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L300:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v970 + int32(1)
	F_errmsg(m, int32(_a_F_thesaurus_init_19), v18)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(588), int32(_a_F_thesaurus_init_17))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L304:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v1248 = v1245 + v564<<(uint(int32(3))%32)
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1248)+4))
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1249)))
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1248)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v1251
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v1250 + int32(1)
	F_errmsg(m, int32(_a_F_thesaurus_init_20), v18-int32(-64))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errhint(m, int32(_a_F_thesaurus_init_21), int32(0))
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(426), int32(_a_F_thesaurus_init_22))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L309:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v1280 = v1277 + v564<<(uint(int32(3))%32)
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+4))
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1281)))
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1280)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v1283
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v1282 + int32(1)
	F_errmsg(m, int32(_a_F_thesaurus_init_23), v18+int32(48))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(419), int32(_a_F_thesaurus_init_22))
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L313:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	F_errmsg(m, int32(_a_F_thesaurus_init_24), int32(0))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(645), int32(_a_F_thesaurus_init_13))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L317:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	F_errmsg(m, int32(_a_F_thesaurus_init_25), int32(0))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(641), int32(_a_F_thesaurus_init_13))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tideq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+2)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2))))
	v9 = int32(16)
	v11 = v7 | v8<<(uint(v9)%32)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v16 = v12 | v13<<(uint(v9)%32)
	if base.Ui32(v11) < base.Ui32(v16) {
		v27 = int32(-1)
	} else {
		if base.Ui32(v16) < base.Ui32(v11) {
			v27 = int32(1)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+4)))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v21) < base.Ui32(v22) {
				v27 = int32(-1)
			} else {
				v27 = base.B2i32(base.Ui32(v22) < base.Ui32(v21))
			}
		}
	}
	return base.I64_extend_i32_u(base.B2i32(v27 == int32(0)))
}
func F_tidout(m *base.Module, l0 int32) int64 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+2)))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9))))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v12
	v14 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10 | v11<<(uint(v14)%32)
	v19 = v7 + v14
	v22 = F_pg_snprintf(m, v19, int32(32), int32(_a_F_tidout_0), v7)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int64(0)
	} else {
		v26 = F_pstrdup(m, v19)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(48)
			return base.I64_extend_i32_u(v26)
		}
	}
}
func F_tidrecv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pq_getmsgint(m, v4, int32(4))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v11 = F_pq_getmsgint(m, v4, int32(2))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v14 = F_palloc(m, int32(6))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)) = uint16(v11)
				*(*uint16)(unsafe.Add(mBase, uint32(v14)+2)) = uint16(v6)
				v19 = int32(base.Ui32(v6) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v14))) = uint16(v19)
				return base.I64_extend_i32_u(v14)
			}
		}
	}
}
func F_tidsend(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v8)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10))))
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+2)))
		F_enlargeStringInfo(m, v8, int32(4))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v25 = v16 | v15<<(uint(int32(16))%32)
			v26 = int32(16711935)
			*(*int32)(unsafe.Add(mBase, uint32(v20+v21))) = base.I32_rotr(v25&v26, int32(8)) | base.I32_rotr(v25, int32(24))&v26
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v20 + int32(4)
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
			F_enlargeStringInfo(m, v8, int32(2))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int64(0)
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v46 = int32(8)
				v50 = v39<<(uint(v46)%32) | int32(base.Ui32(v39)>>(uint(v46)%32))
				*(*uint16)(unsafe.Add(mBase, uint32(v43+v44))) = uint16(v50)
				v52 = int32(2)
				v53 = v43 + v52
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v53
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				*(*int32)(unsafe.Add(mBase, uint32(v56))) = v53 << (uint(v52) % 32)
				m.G0 = v8 + int32(16)
				return base.I64_extend_i32_u(v56)
			}
		}
	}
}
func F_tidsmaller(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = base.I32_wrap_i64(v4)
	v7 = base.I32_wrap_i64(v5)
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+2)))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6))))
	v13 = int32(16)
	v15 = v11 | v12<<(uint(v13)%32)
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+2)))
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7))))
	v20 = v16 | v17<<(uint(v13)%32)
	if base.Ui32(v15) < base.Ui32(v20) {
		v31 = int32(-1)
	} else {
		if base.Ui32(v20) < base.Ui32(v15) {
			v31 = int32(1)
		} else {
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
			if base.Ui32(v25) < base.Ui32(v26) {
				v31 = int32(-1)
			} else {
				v31 = base.B2i32(base.Ui32(v26) < base.Ui32(v25))
			}
		}
	}
	if v31 <= int32(0) {
		v34 = v4
	} else {
		v34 = v5
	}
	return v34 & int64(4294967295)
}
func F_timesub(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v61 int64
	_ = v61
	var v66 int32
	_ = v66
	var __phi66 int32
	_ = __phi66
	var v71 int64
	_ = v71
	var __phi71 int64
	_ = __phi71
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v185 int64
	_ = v185
	var v188 int64
	_ = v188
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v213 int64
	_ = v213
	var v216 int64
	_ = v216
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int64
	_ = v343
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v14 = v13
	goto L3
L2:
	;
	v14 = int32(0)
	goto L3
L3:
	;
	v20 = v14
	goto L6
L4:
	;
	v57 = int64(86400)
	v58 = base.I64_div_s(v53, v57)
	v61 = v53 - v58*v57
	__phi66 = int32(1970)
	__phi71 = v58
	v66 = __phi66
	v71 = __phi71
	goto L13
L5:
	;
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v53 = v50
	v54 = int64(0)
	v56 = int32(0)
	goto L4
L6:
	;
	v30 = v20 - int32(1)
	if v30 < int32(0) {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v36)+8))
	if v33 != v37 {
		v53 = v33
		v54 = v39
		v56 = int32(0)
		goto L4
	} else {
		goto L10
	}
L8:
	;
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v36 = l2 + int32(_a_F_timesub_0) + v30<<(uint(int32(4))%32)
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	if v33 < v37 {
		v20 = v30
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	if v30 == int32(0) {
		v53 = v33
		v54 = v39
		v56 = base.B2i32(int64(0) < v39)
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v36-int32(8))))
	v53 = v33
	v54 = v39
	v56 = base.B2i32(v48 < v39)
	goto L4
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_timesub[0])) = int32(61)
	return int32(0)
L13:
	;
	v76 = base.B2i32(v71 < int64(0))
	if v76 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v179 = base.I32_wrap_i64(v71)
	v180 = base.I64_extend_i32_s(l1)
	v182 = v180 - v54 + v61
	if v182 < int64(0) {
		goto L44
	} else {
		goto L45
	}
L15:
	;
	goto L14
L16:
	;
	if v66&int32(3) != 0 {
		v89 = int32(0)
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if base.Ui64(int64(1571958030700)) < base.Ui64(v71+int64(785979015533)) {
		goto L12
	} else {
		goto L23
	}
L19:
	;
	v92 = int64(*(*int32)(unsafe.Add(mBase, uint32(v89<<(uint(int32(2))%32))+uint32(_c_F_timesub[1]))))
	if v71 < v92 {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	v84 = base.I32_rem_s(v66, int32(100))
	if v84 != 0 {
		v89 = int32(1)
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v86 = base.I32_rem_s(v66, int32(400))
	v89 = base.B2i32(v86 == int32(0))
	goto L19
L22:
	;
	goto L18
L23:
	;
	if v71 < int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v100 = int32(-1)
	goto L26
L25:
	;
	v100 = int32(1)
	goto L26
L26:
	;
	v102 = base.I64_div_s(v71, int64(366))
	if base.Ui64(v71+int64(365)) < base.Ui64(int64(731)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v108 = v100
	goto L29
L28:
	;
	v108 = base.I32_wrap_i64(v102)
	goto L29
L29:
	;
	if int32(0) <= v66 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v117 = v108 + v66
	v119 = v117 - int32(1)
	if v119 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	if v108 <= v66^int32(2147483647) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v108 < int32(-2147483648)-v66 {
		goto L12
	} else {
		goto L35
	}
L34:
	;
	goto L12
L35:
	;
	goto L30
L36:
	;
	v151 = v66 - int32(1)
	if v151 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v123 = int32(0) - v117
	v127 = base.I32_div_u_s(v123, int32(100))
	v130 = base.I32_div_u_s(v123, int32(400))
	v143 = int32(base.Ui32(v123)>>(uint(int32(2))%32)) - v127 + v130 ^ int32(-1)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v137 = base.I32_div_u_s(v119, int32(100))
	v140 = base.I32_div_u_s(v119, int32(400))
	v143 = int32(base.Ui32(v119)>>(uint(int32(2))%32)) - v137 + v140
	goto L36
L40:
	;
	__phi66 = v117
	__phi71 = (base.I64_extend_i32_s(v117)-base.I64_extend_i32_s(v66))*int64(-365) + v71 - base.I64_extend_i32_s(v143-v175)
	v66 = __phi66
	v71 = __phi71
	goto L13
L41:
	;
	v155 = int32(0) - v66
	v159 = base.I32_div_u_s(v155, int32(100))
	v162 = base.I32_div_u_s(v155, int32(400))
	v175 = int32(base.Ui32(v155)>>(uint(int32(2))%32)) - v159 + v162 ^ int32(-1)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v169 = base.I32_div_u_s(v151, int32(100))
	v172 = base.I32_div_u_s(v151, int32(400))
	v175 = int32(base.Ui32(v151)>>(uint(int32(2))%32)) - v169 + v172
	goto L40
L44:
	;
	v185 = int64(-86400)
	if base.Ui64(v182) <= base.Ui64(v185) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v209 = v179
	v210 = v182
	goto L46
L46:
	;
	if base.Ui64(int64(86400)) <= base.Ui64(v210) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v188 = v185
	goto L49
L48:
	;
	v188 = v182
	goto L49
L49:
	;
	v190 = v54 + v188 - v61
	v192 = base.I64_extend_i32_u(base.B2i32(v190 != v180))
	v195 = int64(86400)
	v196 = base.I64_div_u_s(v190-(v192+v180), v195)
	v197 = v196 + v192
	v209 = base.I32_wrap_i64(v197) ^ int32(-1) + v179
	v210 = v61 + v197*v195 + v180 - v54 + v195
	goto L46
L50:
	;
	v213 = int64(172799)
	if v213 <= v210 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v231 = v209
	v232 = v210
	goto L52
L52:
	;
	if v231 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v216 = v213
	goto L55
L54:
	;
	v216 = v210
	goto L55
L55:
	;
	v220 = int64(86400)
	v221 = base.I64_div_u_s(v210-v216+int64(86399), v220)
	v231 = v209 + base.I32_wrap_i64(v221) + int32(1)
	v232 = v210 + v221*int64(-86400) - v220
	goto L52
L56:
	;
	v236 = v231
	v239 = v66
	goto L59
L57:
	;
	v269 = v231
	v272 = v66
	goto L58
L58:
	;
	v281 = v269
	v284 = v272
	goto L66
L59:
	;
	if v239 == int32(-2147483648) {
		goto L12
	} else {
		goto L61
	}
L60:
	;
	v269 = v266
	v272 = v252
	goto L58
L61:
	;
	v252 = v239 - int32(1)
	if v252&int32(3) != 0 {
		v262 = int32(0)
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v262<<(uint(int32(2))%32))+uint32(_c_F_timesub[1])))
	v266 = v265 + v236
	if v266 < int32(0) {
		v236 = v266
		v239 = v252
		goto L59
	} else {
		goto L65
	}
L63:
	;
	v257 = base.I32_rem_s(v252, int32(100))
	if v257 != 0 {
		v262 = int32(1)
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v259 = base.I32_rem_s(v252, int32(400))
	v262 = base.B2i32(v259 == int32(0))
	goto L62
L65:
	;
	goto L60
L66:
	;
	v294 = v284 & int32(3)
	if v294 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_timesub[2])) = v284
	if v284 < int32(-2147481748) {
		goto L12
	} else {
		goto L80
	}
L68:
	;
	goto L67
L69:
	;
	if v284 == int32(2147483647) {
		goto L12
	} else {
		goto L79
	}
L70:
	;
	v298 = base.I32_rem_s(v284, int32(100))
	if v298 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if v281 < int32(365) {
		goto L68
	} else {
		goto L78
	}
L73:
	;
	v302 = base.I32_rem_s(v284, int32(400))
	v304 = base.B2i32(v302 == int32(0))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v304<<(uint(int32(2))%32))+uint32(_c_F_timesub[1])))
	if v281 < v307 {
		goto L68
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if v281 < int32(366) {
		goto L68
	} else {
		goto L77
	}
L76:
	;
	v316 = v304
	goto L69
L77:
	;
	v316 = int32(1)
	goto L69
L78:
	;
	v316 = int32(0)
	goto L69
L79:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v316<<(uint(int32(2))%32))+uint32(_c_F_timesub[1])))
	v281 = v281 - v323
	v284 = v284 + int32(1)
	goto L66
L80:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_timesub[3])) = v281
	*(*int32)(unsafe.Add(mBase, _c_F_timesub[2])) = v284 - int32(1900)
	v339 = base.I32_rem_s(v284-int32(1970), int32(7))
	v340 = int32(0)
	v343 = base.I64_div_u_s(v232, int64(3600))
	*(*uint32)(unsafe.Add(mBase, _c_F_timesub[4])) = uint32(v343)
	if v284 <= v340 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v376 = int32(7)
	v377 = base.I32_rem_s(v371+(v281+v339)-int32(473), v376)
	if v377 < int32(0) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v349 = int32(0) - v284
	v353 = base.I32_div_u_s(v349, int32(100))
	v356 = base.I32_div_u_s(v349, int32(400))
	v371 = int32(base.Ui32(v349)>>(uint(int32(2))%32)) - v353 + v356 ^ int32(-1)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v361 = v284 - int32(1)
	v365 = base.I32_div_u_s(v361, int32(100))
	v368 = base.I32_div_u_s(v361, int32(400))
	v371 = int32(base.Ui32(v361)>>(uint(int32(2))%32)) - v365 + v368
	goto L81
L85:
	;
	v382 = v377 + v376
	goto L87
L86:
	;
	v382 = v377
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_timesub[5])) = v382
	v388 = base.I32_wrap_i64(v232 - v343*int64(3600))
	v389 = int32(_a_F_timesub_1)
	v391 = int32(60)
	v392 = base.I32_div_u_s(v388&v389, v391)
	*(*int32)(unsafe.Add(mBase, _c_F_timesub[6])) = v392
	*(*int32)(unsafe.Add(mBase, _c_F_timesub[7])) = v56 + (v388-v392*v391)&v389
	if v294 != 0 {
		v410 = int32(0)
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v412 = v410 * int32(48)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)+uint32(_c_F_timesub[8])))
	if v413 <= v281 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v405 = base.I32_rem_s(v284, int32(100))
	if v405 != 0 {
		v410 = int32(1)
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v407 = base.I32_rem_s(v284, int32(400))
	v410 = base.B2i32(v407 == int32(0))
	goto L88
L91:
	;
	v417 = v281
	v419 = v340
	v420 = v413
	goto L94
L92:
	;
	v437 = v281
	v439 = v340
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_timesub[9])) = l1
	*(*int32)(unsafe.Add(mBase, _c_F_timesub[10])) = v439
	*(*int32)(unsafe.Add(mBase, _c_F_timesub[11])) = v437 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_timesub[12])) = int32(0)
	return int32(_a_F_timesub_2)
L94:
	;
	v429 = v417 - v420
	v431 = v419 + int32(1)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v412+int32(_a_F_timesub_3)+v431<<(uint(int32(2))%32))))
	if v435 <= v429 {
		v417 = v429
		v419 = v431
		v420 = v435
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v437 = v429
	v439 = v431
	goto L93
L96:
	;
	goto L95
}
func F_tliSwitchPoint(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l1 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
	m.G0 = v11 + int32(16)
	return v63
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v17 <= int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v24 = int32(0)
	goto L8
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v24<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v33 == l0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L5
L10:
	;
	if l2 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v33
	goto L13
L12:
	;
	goto L13
L13:
	;
	v37 = v24 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v37 < v38 {
		v24 = v37
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	return int64(0)
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg(m, int32(_a_F_tliSwitchPoint_0), v11)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_tliSwitchPoint_1), int32(591), int32(_a_F_tliSwitchPoint_2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tok_is_keyword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	v6 = int32(1)
	if l0 == l2 {
		v45 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v45
L2:
	;
	if l0 != int32(277) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v45 = int32(0)
	goto L1
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v10 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if base.B2i32(v16 == int32(0))|base.B2i32(v16 != v19) != 0 {
		v37 = v16
		v38 = v19
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v37-v38 == int32(0) {
		v45 = v6
		goto L1
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v22 = v11
	v23 = l3
	goto L10
L10:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v27 == int32(0) {
		v37 = v27
		v38 = v26
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v37 = v27
	v38 = v26
	goto L8
L12:
	;
	v30 = int32(1)
	if v27 == v26 {
		v22 = v22 + v30
		v23 = v23 + v30
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L3
}
func F_tokenize_auth_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v477 int32
	_ = v477
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v521 int32
	_ = v521
	var v547 int32
	_ = v547
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v575 int32
	_ = v575
	var v586 int32
	_ = v586
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v616 int32
	_ = v616
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v777 int32
	_ = v777
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v898 int32
	_ = v898
	var v911 int32
	_ = v911
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v954 int32
	_ = v954
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v982 int32
	_ = v982
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v1010 int32
	_ = v1010
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1067 int32
	_ = v1067
	var v1074 int32
	_ = v1074
	var v1085 int32
	_ = v1085
	var v1095 int32
	_ = v1095
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1255 int32
	_ = v1255
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1345 int32
	_ = v1345
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1500 int32
	_ = v1500
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	v25 = m.G0
	v27 = v25 - int32(80)
	m.G0 = v27
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = int32(846)
	v31 = int32(_a_F_tokenize_auth_file_0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[0])) = v27 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v27 + int32(24)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1]))
	v50 = F_AllocSetContextCreateInternal(m, v45, int32(_a_F_tokenize_auth_file_1), int32(0), int32(1024), int32(_a_F_tokenize_auth_file_2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v52 = int32(_a_F_tokenize_auth_file_3)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v50
	F_initStringInfo(m, v27+int32(44))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l4 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	goto L8
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v1518
	F_MemoryContextDelete(m, v1517)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L1
	} else {
		goto L235
	}
L8:
	;
	if int32(base.Ui32(v64)>>(uint(int32(4))%32))&int32(1) != 0 {
		v1500 = v27
		v1517 = v50
		v1518 = v53
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v69 = int32(1)
	v72 = l0
	v73 = l1
	v74 = l2
	v75 = l3
	v77 = v27
	v88 = l4 + v69
	v92 = v69
	v94 = v50
	v95 = v53
	goto L10
L10:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	goto L12
L11:
	;
	v1500 = v1130
	v1517 = v1147
	v1518 = v1148
	goto L7
L12:
	;
	if int32(base.Ui32(v96)>>(uint(int32(5))%32))&int32(1) != 0 {
		v1500 = v77
		v1517 = v94
		v1518 = v95
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v101 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = v101
	v105 = v77 + int32(44)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	*(*uint8)(unsafe.Add(mBase, uint32(v106))) = uint8(v101)
	*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v101
	goto L14
L14:
	;
	v113 = int32(0)
	v114 = F_pg_get_line_append(m, v73, v105)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	goto L33
L16:
	;
	if v114 == int32(0) {
		v255 = v113
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v126 = v101
	v139 = v113
	goto L18
L18:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v77)+44))
	v143 = F_strlen(m, v142)
	mBase = m.M
	if v143 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v255 = v228
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+48)) = v211
	if v211 <= v126 {
		v255 = v139
		goto L15
	} else {
		goto L29
	}
L21:
	;
	v211 = v143
	goto L20
L22:
	;
	goto L23
L23:
	;
	v150 = v143
	goto L24
L24:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v142-int32(1)))))
	switch v173 - int32(10) {
	case 0, 3:
		goto L27
	default:
		v185 = v150
		goto L26
	}
L25:
	;
	v211 = v185
	goto L20
L26:
	;
	goto L25
L27:
	;
	v176 = int32(0)
	v177 = int32(1)
	v178 = v150 - v177
	*(*uint8)(unsafe.Add(mBase, uint32(v142+v178))) = uint8(v176)
	if v177 < v150 {
		v150 = v178
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v185 = v176
	goto L26
L29:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v77)+44))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214+v211-int32(1)))))
	if v218 != int32(92) {
		v255 = v139
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v221 = int32(1)
	v222 = v211 - v221
	*(*int32)(unsafe.Add(mBase, uint32(v77)+48)) = v222
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v222+v214))) = uint8(v225)
	v228 = v139 + v221
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v77)+48))
	v232 = F_pg_get_line_append(m, v73, v77+int32(44))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v232 != 0 {
		v126 = v229
		v139 = v228
		goto L18
	} else {
		goto L32
	}
L32:
	;
	goto L19
L33:
	;
	if int32(base.Ui32(v258)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[2]))
	v266 = F_errstart(m, v75, int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v288 = int32(0)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v77)+44))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if base.B2i32(v292 == v288)|v287 != 0 {
		v1125 = v72
		v1126 = v73
		v1127 = v74
		v1128 = v75
		v1130 = v77
		v1133 = base.B2i32(v287 == v288)
		v1139 = v288
		v1141 = v88
		v1145 = v92
		v1146 = v255
		v1147 = v94
		v1148 = v95
		goto L45
	} else {
		goto L46
	}
L37:
	;
	if v266 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[2])) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v72
	v285 = F_psprintf(m, int32(_a_F_tokenize_auth_file_4), v77)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L44
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v72
	F_errmsg(m, int32(_a_F_tokenize_auth_file_4), v77+int32(16))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_tokenize_auth_file_5), int32(765), int32(_a_F_tokenize_auth_file_1))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	v1500 = v77
	v1517 = v94
	v1518 = v95
	goto L7
L45:
	;
	if v1133 != 0 {
		goto L161
	} else {
		goto L162
	}
L46:
	;
	v296 = v72
	v297 = v73
	v298 = v74
	v299 = v75
	v301 = v77
	v303 = v291
	v310 = v288
	v312 = v88
	v316 = v92
	v317 = v255
	v318 = v94
	v319 = v95
	goto L47
L47:
	;
	F_initStringInfo(m, v301+int32(60))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	v1125 = v296
	v1126 = v297
	v1127 = v298
	v1128 = v299
	v1130 = v301
	v1133 = v1119
	v1139 = v1116
	v1141 = v312
	v1145 = v316
	v1146 = v317
	v1147 = v318
	v1148 = v319
	goto L45
L49:
	;
	v332 = v303
	v342 = int32(0)
	goto L50
L50:
	;
	v350 = v301 + int32(60)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	v352 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v351))) = uint8(v352)
	*(*int32)(unsafe.Add(mBase, uint32(v350)+12)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v350)+4)) = v352
	goto L52
L51:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v301)+60))
	F_pfree(m, v1102)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L1
	} else {
		goto L153
	}
L52:
	;
	v365 = int32(0)
	v366 = v332
	goto L56
L53:
	;
	goto L51
L54:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v301)+60))
	v785 = v771 & int32(1)
	if v785 != 0 {
		goto L109
	} else {
		goto L110
	}
L55:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v301)+64))
	if v756 <= int32(0) {
		v1085 = v739
		v1095 = v342
		goto L53
	} else {
		goto L107
	}
L56:
	;
	v383 = int32(*(*int8)(unsafe.Add(mBase, uint32(v366))))
	if v383 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v407 = int32(0)
	v415 = v383
	v418 = v401
	v419 = v407
	v421 = v407
	v422 = v388
	v423 = v407
	v430 = v407
	goto L66
L58:
	;
	v386 = int32(0)
	v739 = v366
	v744 = v386
	v750 = v386
	goto L55
L59:
	;
	goto L60
L60:
	;
	v388 = int32(0)
	if v388 <= v383 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	goto L64
L62:
	;
	v399 = v388
	goto L63
L63:
	;
	v400 = int32(1)
	v401 = v366 + v400
	if base.B2i32(v383 == int32(44))|v399 != 0 {
		v365 = v365 + v400
		v366 = v401
		goto L56
	} else {
		goto L65
	}
L64:
	;
	v399 = base.B2i32(base.B2i32(v383 == int32(32))|base.B2i32(v383 == int32(9)) != int32(0))
	goto L63
L65:
	;
	goto L57
L66:
	;
	v435 = base.I32_extend8_s(v415)
	if int32(0) <= v435 {
		goto L74
	} else {
		goto L75
	}
L67:
	;
	v766 = v719
	v771 = v723
	v777 = int32(0)
	goto L54
L68:
	;
	v720 = int32(1)
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v301)+64))
	if v722 != 0 {
		goto L103
	} else {
		goto L104
	}
L69:
	;
	v705 = v419
	v706 = v682
	v719 = v678
	goto L68
L70:
	;
	v662 = F_strlen(m, v645)
	mBase = m.M
	v667 = v662 + (v649 + v332 + v365) + int32(1)
	v668 = int32(0)
	if v430 == v668 {
		v739 = v667
		v744 = v423
		v750 = v668
		goto L55
	} else {
		goto L102
	}
L71:
	;
	v633 = int32(1)
	v634 = v616 - v633
	if v430 == int32(0) {
		v739 = v634
		v744 = v423
		v750 = v633
		goto L55
	} else {
		goto L101
	}
L72:
	;
	v606 = int32(0)
	if v430 == v606 {
		v739 = v489
		v744 = v423
		v750 = v606
		goto L55
	} else {
		goto L100
	}
L73:
	;
	v604 = v586 - int32(1)
	v605 = int32(0)
	if v430 != 0 {
		v766 = v604
		v771 = v423
		v777 = v605
		goto L54
	} else {
		goto L99
	}
L74:
	;
	v439 = v415 & int32(255)
	goto L77
L75:
	;
	goto L76
L76:
	;
	F_appendStringInfoChar(m, v301+int32(60), v435)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L82
	}
L77:
	;
	v445 = int32(0)
	if (base.B2i32(base.B2i32(v439 == int32(32))|base.B2i32(v439 == int32(9)) == v445)|v419)&int32(1) == v445 {
		v586 = v418
		goto L73
	} else {
		goto L78
	}
L78:
	;
	if (base.B2i32(v439 != int32(35))|v419)&int32(1) == int32(0) {
		v645 = v418
		v649 = v422
		goto L70
	} else {
		goto L79
	}
L79:
	;
	if (base.B2i32(v439 != int32(44))|v419)&int32(1) == int32(0) {
		v616 = v418
		goto L71
	} else {
		goto L80
	}
L80:
	;
	if (base.B2i32(v439 != int32(34))|v421)&int32(1) == int32(0) {
		v678 = v418
		v682 = v422
		goto L69
	} else {
		goto L81
	}
L81:
	;
	goto L76
L82:
	;
	if v415&int32(255) != int32(34) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v489 = v418
	v493 = v422
	goto L86
L84:
	;
	v558 = v418
	v562 = v422
	v575 = v421
	goto L85
L85:
	;
	v705 = v419 & (v575 ^ int32(1))
	v706 = v562
	v719 = v558
	goto L68
L86:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
	if v506 == int32(0) {
		goto L72
	} else {
		goto L88
	}
L87:
	;
	v558 = v512
	v562 = v510
	v575 = int32(0)
	goto L85
L88:
	;
	v509 = int32(1)
	v510 = v493 + v509
	v512 = v489 + v509
	v513 = base.I32_extend8_s(v506)
	if int32(0) <= v513 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	goto L92
L90:
	;
	goto L91
L91:
	;
	F_appendStringInfoChar(m, v301+int32(60), v513)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L97
	}
L92:
	;
	v521 = int32(0)
	if (base.B2i32(base.B2i32(v506 == int32(32))|base.B2i32(v506 == int32(9)) == v521)|v419)&int32(1) == v521 {
		v586 = v512
		goto L73
	} else {
		goto L93
	}
L93:
	;
	if (base.B2i32(v506 != int32(35))|v419)&int32(1) == int32(0) {
		v645 = v512
		v649 = v510
		goto L70
	} else {
		goto L94
	}
L94:
	;
	if (base.B2i32(v506 != int32(44))|v419)&int32(1) == int32(0) {
		v616 = v512
		goto L71
	} else {
		goto L95
	}
L95:
	;
	if v506 == int32(34) {
		v678 = v512
		v682 = v510
		goto L69
	} else {
		goto L96
	}
L96:
	;
	goto L91
L97:
	;
	if v506 != int32(34) {
		v489 = v512
		v493 = v510
		goto L86
	} else {
		goto L98
	}
L98:
	;
	goto L87
L99:
	;
	v739 = v604
	v744 = v423
	v750 = v605
	goto L55
L100:
	;
	v766 = v489
	v771 = v423
	v777 = v606
	goto L54
L101:
	;
	v766 = v634
	v771 = v423
	v777 = v633
	goto L54
L102:
	;
	v766 = v667
	v771 = v423
	v777 = v668
	goto L54
L103:
	;
	v723 = v423
	goto L105
L104:
	;
	v723 = v720
	goto L105
L105:
	;
	v724 = int32(1)
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v719))))
	if v730 != 0 {
		v415 = v730
		v418 = v719 + v724
		v419 = v419 ^ v724
		v421 = v705
		v422 = v706 + v724
		v423 = v723
		v430 = v720
		goto L66
	} else {
		goto L106
	}
L106:
	;
	goto L67
L107:
	;
	v766 = v739
	v771 = v744
	v777 = v750
	goto L54
L108:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	if v777&base.B2i32(v1074 == int32(0)) != 0 {
		v332 = v766
		v342 = v1067
		goto L50
	} else {
		goto L152
	}
L109:
	;
	v1026 = int32(_a_F_tokenize_auth_file_3)
	v1027 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1]))
	v1030 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v1030
	v1032 = F_strlen(m, v783)
	mBase = m.M
	v1035 = F_palloc0(m, v1032+int32(13))
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L1
	} else {
		goto L147
	}
L110:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v301)+64))
	if v786 < int32(2) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783))))
	if v789 != int32(64) {
		goto L109
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+76)) = int32(0)
	v796 = F_AbsoluteConfigLocation(m, v783+int32(1), v296)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v800 = F_open_auth_file(m, v796, v299, v312, v301+int32(20))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	if v800 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	F_pfree(m, v796)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	F_tokenize_auth_file(m, v796, v800, v301+int32(76), v299, v312)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L1
	} else {
		goto L119
	}
L118:
	;
	v1067 = v342
	goto L108
L119:
	;
	F_pfree(m, v796)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v301)+76))
	if v812 == int32(0) {
		v1010 = v342
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v1017 = F_FreeFile(m, v800)
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L144
	}
L122:
	;
	v815 = int32(0)
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	if v816 <= v815 {
		v1010 = v342
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v831 = v815
	v836 = v342
	goto L124
L124:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v812)+12))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v843+v831<<(uint(int32(2))%32))))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v847)+16))
	if v848 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v1010 = v982
	goto L121
L126:
	;
	v849 = F_pstrdup(m, v848)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v847)))
	if v852 == int32(0) {
		v982 = v836
		goto L130
	} else {
		goto L131
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = v849
	v1010 = v836
	goto L121
L130:
	;
	v990 = v831 + int32(1)
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v812)+4))
	if v990 < v991 {
		v831 = v990
		v836 = v982
		goto L124
	} else {
		goto L143
	}
L131:
	;
	v855 = int32(0)
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v852)+4))
	if v856 <= v855 {
		v982 = v836
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v870 = v855
	v876 = v836
	goto L133
L133:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v852)+12))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v883+v870<<(uint(int32(2))%32))))
	if v887 == int32(0) {
		v954 = v876
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v982 = v954
	goto L130
L135:
	;
	v962 = v870 + int32(1)
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v852)+4))
	if v962 < v963 {
		v870 = v962
		v876 = v954
		goto L133
	} else {
		goto L142
	}
L136:
	;
	v890 = int32(0)
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v887)+4))
	if v891 <= v890 {
		v954 = v876
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v898 = v890
	v911 = v876
	goto L138
L138:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v887)+12))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v918+v898<<(uint(int32(2))%32))))
	v923 = int32(_a_F_tokenize_auth_file_3)
	v924 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1]))
	v927 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v927
	v929 = F_lappend(m, v911, v922)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L140
	}
L139:
	;
	v954 = v929
	goto L135
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v924
	v934 = v898 + int32(1)
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v887)+4))
	if v934 < v935 {
		v898 = v934
		v911 = v929
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	goto L134
L143:
	;
	goto L125
L144:
	;
	if v312 != 0 {
		v1067 = v1010
		goto L108
	} else {
		goto L145
	}
L145:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[3]))
	F_MemoryContextDelete(m, v1020)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[3])) = int32(0)
	v1067 = v1010
	goto L108
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1035)+8)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1035)+4)) = uint8(v785)
	v1041 = v1035 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v1035))) = v1041
	v1044 = v1032 + int32(1)
	if v1044 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	base.MemoryCopy(m, v1041, v783, v1044)
	goto L150
L149:
	;
	goto L150
L150:
	;
	v1046 = F_lappend(m, v342, v1035)
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v1027
	v1067 = v1046
	goto L108
L152:
	;
	v1085 = v766
	v1095 = v1067
	goto L53
L153:
	;
	if v1095 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v1105 = int32(_a_F_tokenize_auth_file_3)
	v1106 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1]))
	v1109 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v1109
	v1111 = F_lappend(m, v310, v1095)
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L157
	}
L155:
	;
	v1116 = v310
	goto L156
L156:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	v1118 = int32(0)
	v1119 = base.B2i32(v1117 == v1118)
	v1120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1085))))
	if v1120 == v1118 {
		v1125 = v296
		v1126 = v297
		v1127 = v298
		v1128 = v299
		v1130 = v301
		v1133 = v1119
		v1139 = v1116
		v1141 = v312
		v1145 = v316
		v1146 = v317
		v1147 = v318
		v1148 = v319
		goto L45
	} else {
		goto L158
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v1106
	v1116 = v1111
	goto L156
L158:
	;
	if v1117 == int32(0) {
		v303 = v1085
		v310 = v1116
		goto L47
	} else {
		goto L159
	}
L159:
	;
	goto L48
L160:
	;
	v1485 = int32(1)
	v1486 = v1145 + v1146 + v1485
	*(*int32)(unsafe.Add(mBase, uint32(v1130)+28)) = v1486
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1126)))
	goto L233
L161:
	;
	v1150 = v1139
	goto L163
L162:
	;
	v1150 = int32(1)
	goto L163
L163:
	;
	if v1150 == int32(0) {
		goto L160
	} else {
		goto L164
	}
L164:
	;
	if base.B2i32(v1139 == int32(0))|(v1133^int32(1)) != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v1430 = int32(_a_F_tokenize_auth_file_3)
	v1431 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1]))
	v1434 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v1434
	v1437 = F_palloc0(m, int32(20))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L1
	} else {
		goto L225
	}
L166:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1139)+4))
	if v1158 != int32(2) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1139)+12))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1161)+4))
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1162)+12))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1163)))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1161)))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+12))
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1166)))
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1167)))
	v1169 = int32(_a_F_tokenize_auth_file_6)
	v1172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168))))
	v1175 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tokenize_auth_file[4])))
	if base.B2i32(v1172 == int32(0))|base.B2i32(v1172 != v1175) != 0 {
		v1193 = v1172
		v1194 = v1175
		goto L169
	} else {
		goto L170
	}
L168:
	;
	if v1193-v1194 == int32(0) {
		goto L175
	} else {
		goto L176
	}
L169:
	;
	goto L168
L170:
	;
	v1178 = v1168
	v1179 = v1169
	goto L171
L171:
	;
	v1182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179)+1)))
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1178)+1)))
	if v1183 == int32(0) {
		v1193 = v1183
		v1194 = v1182
		goto L169
	} else {
		goto L173
	}
L172:
	;
	v1193 = v1183
	v1194 = v1182
	goto L169
L173:
	;
	v1186 = int32(1)
	if v1183 == v1182 {
		v1178 = v1178 + v1186
		v1179 = v1179 + v1186
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1164)))
	F_tokenize_include_file(m, v1125, v1198, v1127, v1128, v1141, int32(0), v1130+int32(20))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v1205 = int32(_a_F_tokenize_auth_file_7)
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168))))
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tokenize_auth_file[5])))
	if base.B2i32(v1208 == int32(0))|base.B2i32(v1208 != v1211) != 0 {
		v1229 = v1208
		v1230 = v1211
		goto L183
	} else {
		goto L184
	}
L178:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+20))
	if v1204 != 0 {
		goto L165
	} else {
		goto L179
	}
L179:
	;
	goto L160
L180:
	;
	F_pfree(m, v1239)
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L1
	} else {
		goto L223
	}
L181:
	;
	v1345 = v1302
	goto L219
L182:
	;
	if v1229-v1230 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L183:
	;
	goto L182
L184:
	;
	v1214 = v1168
	v1215 = v1205
	goto L185
L185:
	;
	v1218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1215)+1)))
	v1219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1214)+1)))
	if v1219 == int32(0) {
		v1229 = v1219
		v1230 = v1218
		goto L183
	} else {
		goto L187
	}
L186:
	;
	v1229 = v1219
	v1230 = v1218
	goto L183
L187:
	;
	v1222 = int32(1)
	if v1219 == v1218 {
		v1214 = v1214 + v1222
		v1215 = v1215 + v1222
		goto L185
	} else {
		goto L188
	}
L188:
	;
	goto L186
L189:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1164)))
	v1239 = F_GetConfFilesInDir(m, v1234, v1125, v1128, v1130+int32(76), v1130+int32(20))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L1
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v1305 = int32(_a_F_tokenize_auth_file_8)
	v1308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168))))
	v1311 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tokenize_auth_file[6])))
	if base.B2i32(v1308 == int32(0))|base.B2i32(v1308 != v1311) != 0 {
		v1329 = v1308
		v1330 = v1311
		goto L210
	} else {
		goto L211
	}
L192:
	;
	if v1239 == int32(0) {
		goto L165
	} else {
		goto L193
	}
L193:
	;
	F_initStringInfo(m, v1130+int32(60))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v1247 = int32(0)
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+76))
	if v1248 <= v1247 {
		goto L180
	} else {
		goto L195
	}
L195:
	;
	v1255 = v1247
	goto L196
L196:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1239+v1255<<(uint(int32(2))%32))))
	F_tokenize_include_file(m, v1125, v1278, v1127, v1128, v1141, int32(0), v1130+int32(20))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L1
	} else {
		goto L198
	}
L197:
	;
	v1302 = int32(0)
	if v1302 < v1300 {
		goto L181
	} else {
		goto L208
	}
L198:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+20))
	if v1284 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v1286 = v1130 + int32(60)
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+64))
	if int32(0) < v1287 {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	goto L201
L201:
	;
	v1299 = v1255 + int32(1)
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+76))
	if v1299 < v1300 {
		v1255 = v1299
		goto L196
	} else {
		goto L207
	}
L202:
	;
	F_appendStringInfoChar(m, v1286, int32(10))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	v1294 = v1284
	goto L204
L204:
	;
	F_appendStringInfoString(m, v1286, v1294)
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L1
	} else {
		goto L206
	}
L205:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+20))
	v1294 = v1293
	goto L204
L206:
	;
	goto L201
L207:
	;
	goto L197
L208:
	;
	goto L180
L209:
	;
	if v1329-v1330 != 0 {
		goto L165
	} else {
		goto L216
	}
L210:
	;
	goto L209
L211:
	;
	v1314 = v1168
	v1315 = v1305
	goto L212
L212:
	;
	v1318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315)+1)))
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1314)+1)))
	if v1319 == int32(0) {
		v1329 = v1319
		v1330 = v1318
		goto L210
	} else {
		goto L214
	}
L213:
	;
	v1329 = v1319
	v1330 = v1318
	goto L210
L214:
	;
	v1322 = int32(1)
	if v1319 == v1318 {
		v1314 = v1314 + v1322
		v1315 = v1315 + v1322
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1164)))
	F_tokenize_include_file(m, v1125, v1332, v1127, v1128, v1141, int32(1), v1130+int32(20))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+20))
	if v1338 == int32(0) {
		goto L160
	} else {
		goto L218
	}
L218:
	;
	goto L165
L219:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1239+v1345<<(uint(int32(2))%32))))
	F_pfree(m, v1368)
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L1
	} else {
		goto L221
	}
L220:
	;
	goto L180
L221:
	;
	v1372 = v1345 + int32(1)
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+76))
	if v1372 < v1373 {
		v1345 = v1372
		goto L219
	} else {
		goto L222
	}
L222:
	;
	goto L220
L223:
	;
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+64))
	if v1401 == int32(0) {
		goto L160
	} else {
		goto L224
	}
L224:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1130)+20)) = v1404
	goto L165
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1437))) = v1139
	v1440 = F_pstrdup(m, v1125)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1437)+8)) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(v1437)+4)) = v1440
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+44))
	v1445 = F_pstrdup(m, v1444)
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1437)+12)) = v1445
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1130)+20))
	if v1448 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1449 = F_pstrdup(m, v1448)
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L1
	} else {
		goto L231
	}
L229:
	;
	v1452 = int32(0)
	goto L230
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1437)+16)) = v1452
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1127)))
	v1455 = F_lappend(m, v1454, v1437)
	mBase = m.M
	v1456 = m.ExcPending
	if v1456 != 0 {
		goto L1
	} else {
		goto L232
	}
L231:
	;
	v1452 = v1449
	goto L230
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1127))) = v1455
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v1431
	goto L160
L233:
	;
	if int32(base.Ui32(v1488)>>(uint(int32(4))%32))&v1485 == int32(0) {
		v72 = v1125
		v73 = v1126
		v74 = v1127
		v75 = v1128
		v77 = v1130
		v88 = v1141
		v92 = v1486
		v94 = v1147
		v95 = v1148
		goto L10
	} else {
		goto L234
	}
L234:
	;
	goto L11
L235:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(v1500)+32))
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[0])) = v1524
	m.G0 = v1500 + int32(80)
	return
}
func F_tolower_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		if base.Ui32(l0) <= base.Ui32(int32(255)) {
			if base.Ui32(l0-int32(65)) < base.Ui32(int32(26)) {
				v29 = l0 | int32(32)
			} else {
				v29 = l0
			}
			v30 = v29
		} else {
			v30 = l0
		}
		return v30
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
		if v5&int32(1) == int32(0) {
			if base.Ui32(l0) <= base.Ui32(int32(255)) {
				if base.Ui32(l0-int32(65)) < base.Ui32(int32(26)) {
					v29 = l0 | int32(32)
				} else {
					v29 = l0
				}
				v30 = v29
			} else {
				v30 = l0
			}
			return v30
		} else {
			if base.Ui32((l0-int32(65))&int32(255)) < base.Ui32(int32(26)) {
				v18 = l0 | int32(32)
			} else {
				v18 = l0
			}
			return v18
		}
	}
}
func F_transformAggregateCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v6
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+50)))
	if v19 != int32(110) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v256
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v266 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L2:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v98 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v98
	if l2 == v98 {
		v151 = int32(1)
		goto L33
	} else {
		goto L34
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v24 = v22
	goto L7
L6:
	;
	v24 = int32(0)
	goto L7
L7:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v27 = v25
	goto L10
L9:
	;
	v27 = int32(0)
	goto L10
L10:
	;
	v28 = v24 - v27
	v29 = F_list_copy_tail(m, l2, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v31 = int32(0)
	if base.B2i32(l2 == v31)|base.B2i32(v28 <= v31) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v41
	v50 = int32(0)
	v51 = v6
	v53 = int32(1)
	v54 = v6
	goto L20
L14:
	;
	v41 = int32(0)
	goto L16
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v28 < v38 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L13
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v28
	goto L19
L18:
	;
	goto L19
L19:
	;
	v41 = l2
	goto L16
L20:
	;
	v57 = int32(0)
	if v29 == v57 {
		v67 = v57
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if l3 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v61 <= v51 {
		v67 = int32(0)
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v67 = v63 + v51<<(uint(int32(2))%32)
	goto L22
L25:
	;
	v256 = v50
	v260 = int32(0)
	v262 = v6
	goto L1
L26:
	;
	goto L27
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if base.B2i32(v67 == int32(0))|base.B2i32(v73 <= v51) != 0 {
		v256 = v50
		v260 = v54
		v262 = v6
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v76 == int32(0) {
		v256 = v50
		v260 = v54
		v262 = v6
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+v51<<(uint(int32(2))%32))))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v85 = int32(0)
	v87 = F_makeTargetEntry(m, v83, base.I32_extend16_s(v53), v85, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v89 = F_lappend(m, v50, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L11
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v89
	v92 = int32(1)
	v96 = F_addTargetToSortList(m, l0, v87, v54, v89, v82)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	v50 = v89
	v51 = v51 + v92
	v53 = v53 + v92
	v54 = v96
	goto L20
L33:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v151
	v155 = v15 + int32(12)
	v157 = F_transformSortClause(m, l0, l3, v155, int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L11
	} else {
		goto L41
	}
L34:
	;
	v103 = int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v104 <= int32(0) {
		v151 = v103
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v112 = v103
	v113 = v6
	v115 = v6
	goto L36
L36:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v113<<(uint(int32(2))%32))))
	v125 = int32(0)
	v127 = F_makeTargetEntry(m, v123, base.I32_extend16_s(v112), v125, v125)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L11
	} else {
		goto L38
	}
L37:
	;
	v151 = base.I32_extend16_s(v133)
	goto L33
L38:
	;
	v129 = F_lappend(m, v115, v127)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v129
	v132 = int32(1)
	v133 = v112 + v132
	v135 = v113 + v132
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v135 < v136 {
		v112 = v133
		v113 = v135
		v115 = v129
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	if l4 == int32(0) {
		v248 = v6
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v152
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v256 = v250
	v260 = v157
	v262 = v248
	goto L1
L43:
	;
	v162 = F_transformDistinctClause(m, l0, v155, v157, int32(1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L11
	} else {
		goto L44
	}
L44:
	;
	if v162 == int32(0) {
		v248 = v6
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v166 <= int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v248 = v162
	goto L42
L47:
	;
	v169 = int32(0)
	if v169 < v166 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v173 = v166
	goto L50
L49:
	;
	v173 = v169
	goto L50
L50:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v180 = v169
	goto L51
L51:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v174+v180<<(uint(int32(2))%32))))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	if v191 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v196 = F_get_sortgroupclause_expr(m, v190, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L11
	} else {
		goto L57
	}
L53:
	;
	v193 = v180 + int32(1)
	if v173 != v193 {
		v180 = v193
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	goto L46
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	v205 = F_exprType(m, v196)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	v207 = F_format_type_be(m, v205)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v207
	F_errmsg(m, int32(_a_F_transformAggregateCall_0), v15)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	v215 = F_errdetail(m, int32(_a_F_transformAggregateCall_1), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	v217 = F_exprLocation(m, v196)
	mBase = m.M
	F_parser_errposition(m, l0, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_transformAggregateCall_2), int32(222), int32(_a_F_transformAggregateCall_3))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	if v319 == int32(0) {
		v368 = v320
		goto L79
	} else {
		goto L80
	}
L67:
	;
	v319 = v256
	v320 = int32(0)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v270 = int32(0)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v271 <= v270 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v319 = v313
	v320 = v307
	goto L66
L71:
	;
	v307 = int32(0)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v281 = v270
	v282 = int32(0)
	goto L74
L74:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288+v281<<(uint(int32(2))%32))))
	v293 = F_exprType(m, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L11
	} else {
		goto L76
	}
L75:
	;
	v307 = v295
	goto L70
L76:
	;
	v295 = F_lappend_oid(m, v282, v293)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	v298 = v281 + int32(1)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	if v298 < v299 {
		v281 = v298
		v282 = v295
		goto L74
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v368
	F_check_agglevels_and_constraints(m, l0, l1)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L11
	} else {
		goto L90
	}
L80:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v328 <= int32(0) {
		v368 = v320
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v334 = int32(0)
	v338 = v320
	goto L82
L82:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v344+v334<<(uint(int32(2))%32))))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+26)))
	if v349 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v368 = v357
	goto L79
L84:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	v353 = F_exprType(m, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L11
	} else {
		goto L87
	}
L85:
	;
	v357 = v338
	goto L86
L86:
	;
	v359 = v334 + int32(1)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v359 < v360 {
		v334 = v359
		v338 = v357
		goto L82
	} else {
		goto L89
	}
L87:
	;
	v355 = F_lappend_oid(m, v338, v353)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L11
	} else {
		goto L88
	}
L88:
	;
	v357 = v355
	goto L86
L89:
	;
	goto L83
L90:
	;
	m.G0 = v15 + int32(16)
	return
}
func F_transformAssignedExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l2
	if int32(0) < l4 {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v23 = F_attnumTypeId(m, v22, l4)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
			v34 = v27 + v28<<(uint(int32(3))%32) + l4*int32(100)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
			if l1 == int32(0) {
				v72 = F_exprType(m, l1)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					if l5 == int32(0) {
						v122 = v72
						v126 = F_coerce_to_target_type(m, l0, l1, v122, v23, v36, int32(1), int32(2), int32(-1))
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return int32(0)
						} else {
							if v126 != 0 {
								v159 = v126
								*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v18
								m.G0 = v16 + int32(32)
								return v159
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(67141764))
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return int32(0)
									} else {
										v135 = F_format_type_be(m, v23)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											v137 = F_format_type_be(m, v122)
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v137
												*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v135
												*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l3
												F_errmsg(m, int32(_a_F_transformAssignedExpr_0), v16+int32(16))
												mBase = m.M
												v146 = m.ExcPending
												if v146 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(_a_F_transformAssignedExpr_1), int32(0))
													mBase = m.M
													v150 = m.ExcPending
													if v150 != 0 {
														return int32(0)
													} else {
														v151 = F_exprLocation(m, l1)
														mBase = m.M
														F_parser_errposition(m, l0, v151)
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_transformAssignedExpr_2), int32(597), int32(_a_F_transformAssignedExpr_3))
															mBase = m.M
															v158 = m.ExcPending
															if v158 != 0 {
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
							}
						}
					} else {
						if l2 == int32(15) {
							v78 = F_makeNullConst(m, v23, v36, v35)
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								v87 = v78
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
								v91 = F_transformAssignmentIndirection(m, l0, v87, l3, int32(0), v23, v36, v35, l5, v89, l1, int32(1), l6)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v159 = v91
									*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v18
									m.G0 = v16 + int32(32)
									return v159
								}
							}
						} else {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
							v84 = F_makeVar(m, v81, base.I32_extend16_s(l4), v23, v36, v35, int32(0))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v84)+44)) = l6
								v87 = v84
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
								v91 = F_transformAssignmentIndirection(m, l0, v87, l3, int32(0), v23, v36, v35, l5, v89, l1, int32(1), l6)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v159 = v91
									*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v18
									m.G0 = v16 + int32(32)
									return v159
								}
							}
						}
					}
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v39 != int32(57) {
					v72 = F_exprType(m, l1)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						if l5 == int32(0) {
							v122 = v72
							v126 = F_coerce_to_target_type(m, l0, l1, v122, v23, v36, int32(1), int32(2), int32(-1))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return int32(0)
							} else {
								if v126 != 0 {
									v159 = v126
									*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v18
									m.G0 = v16 + int32(32)
									return v159
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67141764))
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
											return int32(0)
										} else {
											v135 = F_format_type_be(m, v23)
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return int32(0)
											} else {
												v137 = F_format_type_be(m, v122)
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v137
													*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v135
													*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l3
													F_errmsg(m, int32(_a_F_transformAssignedExpr_0), v16+int32(16))
													mBase = m.M
													v146 = m.ExcPending
													if v146 != 0 {
														return int32(0)
													} else {
														F_errhint(m, int32(_a_F_transformAssignedExpr_1), int32(0))
														mBase = m.M
														v150 = m.ExcPending
														if v150 != 0 {
															return int32(0)
														} else {
															v151 = F_exprLocation(m, l1)
															mBase = m.M
															F_parser_errposition(m, l0, v151)
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_transformAssignedExpr_2), int32(597), int32(_a_F_transformAssignedExpr_3))
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
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
								}
							}
						} else {
							if l2 == int32(15) {
								v78 = F_makeNullConst(m, v23, v36, v35)
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									v87 = v78
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
									v91 = F_transformAssignmentIndirection(m, l0, v87, l3, int32(0), v23, v36, v35, l5, v89, l1, int32(1), l6)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v159 = v91
										*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v18
										m.G0 = v16 + int32(32)
										return v159
									}
								}
							} else {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
								v84 = F_makeVar(m, v81, base.I32_extend16_s(l4), v23, v36, v35, int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v84)+44)) = l6
									v87 = v84
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
									v91 = F_transformAssignmentIndirection(m, l0, v87, l3, int32(0), v23, v36, v35, l5, v89, l1, int32(1), l6)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v159 = v91
										*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v18
										m.G0 = v16 + int32(32)
										return v159
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v36
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v23
					if l5 == int32(0) {
						v47 = F_exprType(m, l1)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v122 = v47
							v126 = F_coerce_to_target_type(m, l0, l1, v122, v23, v36, int32(1), int32(2), int32(-1))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return int32(0)
							} else {
								if v126 != 0 {
									v159 = v126
									*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v18
									m.G0 = v16 + int32(32)
									return v159
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67141764))
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
											return int32(0)
										} else {
											v135 = F_format_type_be(m, v23)
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return int32(0)
											} else {
												v137 = F_format_type_be(m, v122)
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v137
													*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v135
													*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l3
													F_errmsg(m, int32(_a_F_transformAssignedExpr_0), v16+int32(16))
													mBase = m.M
													v146 = m.ExcPending
													if v146 != 0 {
														return int32(0)
													} else {
														F_errhint(m, int32(_a_F_transformAssignedExpr_1), int32(0))
														mBase = m.M
														v150 = m.ExcPending
														if v150 != 0 {
															return int32(0)
														} else {
															v151 = F_exprLocation(m, l1)
															mBase = m.M
															F_parser_errposition(m, l0, v151)
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_transformAssignedExpr_2), int32(597), int32(_a_F_transformAssignedExpr_3))
																mBase = m.M
																v158 = m.ExcPending
																if v158 != 0 {
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
								}
							}
						}
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								if v51 == int32(78) {
									F_errmsg(m, int32(_a_F_transformAssignedExpr_4), int32(0))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, l0, l6)
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_transformAssignedExpr_2), int32(513), int32(_a_F_transformAssignedExpr_3))
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
								} else {
									F_errmsg(m, int32(_a_F_transformAssignedExpr_5), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, l0, l6)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_transformAssignedExpr_2), int32(518), int32(_a_F_transformAssignedExpr_3))
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
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v96 = m.ExcPending
		if v96 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = l3
				F_errmsg(m, int32(_a_F_transformAssignedExpr_6), v16)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					F_parser_errposition(m, l0, l6)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_transformAssignedExpr_2), int32(486), int32(_a_F_transformAssignedExpr_3))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
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
func F_transformAssignmentSubscripts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l4
	v24 = v19 + int32(12)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v28 = F_getBaseTypeAndTypmod(m, v25, v19+int32(8))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v24))) = v28
		switch v28 - int32(22) {
		case 0:
			v37 = int32(1005)
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = v37
		default:
		case 8:
			v37 = int32(1028)
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = v37
		}
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
		v43 = F_transformContainerSubscripts(m, l0, l1, v40, v41, l6, int32(1))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
			if l3 != v49 {
				v51 = F_get_typcollation(m, v49)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					v53 = v51
					v54 = F_transformAssignmentIndirection(m, l0, int32(0), l2, int32(1), v48, v45, v53, l7, l8, l9, l10, l11)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v43)+36)) = v54
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v57
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v59
						if v57 == l3 {
							v92 = v43
							m.G0 = v19 + int32(16)
							return v92
						} else {
							v62 = F_exprType(m, v43)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								v66 = F_coerce_to_target_type(m, l0, v43, v62, l3, l4, l10, int32(2), int32(-1))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									if v66 != 0 {
										v92 = v66
										m.G0 = v19 + int32(16)
										return v92
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(101744772))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v75 = F_format_type_be(m, v62)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													v77 = F_format_type_be(m, l3)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v77
														*(*int32)(unsafe.Add(mBase, uint32(v19))) = v75
														F_errmsg(m, int32(_a_F_transformAssignmentSubscripts_0), v19)
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return int32(0)
														} else {
															F_parser_errposition(m, l0, l11)
															mBase = m.M
															v85 = m.ExcPending
															if v85 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_transformAssignmentSubscripts_1), int32(1005), int32(_a_F_transformAssignmentSubscripts_2))
																mBase = m.M
																v90 = m.ExcPending
																if v90 != 0 {
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
								}
							}
						}
					}
				}
			} else {
				v53 = l5
				v54 = F_transformAssignmentIndirection(m, l0, int32(0), l2, int32(1), v48, v45, v53, l7, l8, l9, l10, l11)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v43)+36)) = v54
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v57
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v43)+16)) = v59
					if v57 == l3 {
						v92 = v43
						m.G0 = v19 + int32(16)
						return v92
					} else {
						v62 = F_exprType(m, v43)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v66 = F_coerce_to_target_type(m, l0, v43, v62, l3, l4, l10, int32(2), int32(-1))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								if v66 != 0 {
									v92 = v66
									m.G0 = v19 + int32(16)
									return v92
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(101744772))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											v75 = F_format_type_be(m, v62)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												v77 = F_format_type_be(m, l3)
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v77
													*(*int32)(unsafe.Add(mBase, uint32(v19))) = v75
													F_errmsg(m, int32(_a_F_transformAssignmentSubscripts_0), v19)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int32(0)
													} else {
														F_parser_errposition(m, l0, l11)
														mBase = m.M
														v85 = m.ExcPending
														if v85 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_transformAssignmentSubscripts_1), int32(1005), int32(_a_F_transformAssignmentSubscripts_2))
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
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
							}
						}
					}
				}
			}
		}
	}
}
func F_transformSetOperationTree(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
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
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int64
	_ = v235
	var v246 int32
	_ = v246
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v524 int32
	_ = v524
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	F_check_stack_depth(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v27 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L161
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L153
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v30 != 0 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L148
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v31 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	m.G0 = v21 + int32(16)
	return v584
L10:
	;
	v490 = F_make_parsestate(m, l0)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L127
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v34 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v35 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v36 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v37 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v39 = F_palloc0(m, int32(36))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(142)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)))
	v45 = v44
	goto L19
L18:
	;
	v45 = v5
	goto L19
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v46
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)) = uint8(v48)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v54 = F_transformSetOperationTree(m, l0, v50, int32(0), v21+int32(12))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v58 = int32(0)
	if base.B2i32(l2 == v58)|base.B2i32(v45&int32(1) == v58) == v58 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v67 == int32(142) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	if v46 == int32(2) {
		goto L47
	} else {
		goto L48
	}
L24:
	;
	v77 = v54
	goto L27
L25:
	;
	v99 = v54
	goto L26
L26:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v111+v112<<(uint(int32(2))%32)-int32(4))))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+36))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+76))
	v130 = int32(0)
	v137 = int32(1)
	v140 = v5
	goto L30
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v89 == int32(142) {
		v77 = v88
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v99 = v88
	goto L26
L29:
	;
	goto L28
L30:
	;
	v141 = int32(0)
	if v57 == v141 {
		v151 = v141
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L23
L32:
	;
	goto L31
L33:
	;
	if v120 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v145 <= v130 {
		v151 = int32(0)
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v151 = v147 + v130<<(uint(int32(2))%32)
	goto L33
L36:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v160+v130<<(uint(int32(2))%32))))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
	v172 = F_pstrdup(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L44
	}
L37:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_analyzeCTETargetList(m, l0, v163, v162)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L43
	}
L38:
	;
	v162 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if base.B2i32(v151 == int32(0))|base.B2i32(v157 <= v130) != 0 {
		v162 = v140
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	if v160 != 0 {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v162 = v140
	goto L37
L43:
	;
	goto L32
L44:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v177 = F_makeTargetEntry(m, v174, base.I32_extend16_s(v137), v172, int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v179 = F_lappend(m, v140, v177)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v181 = int32(1)
	v130 = v130 + v181
	v137 = v137 + v181
	v140 = v179
	goto L30
L47:
	;
	v208 = int32(_a_F_transformSetOperationTree_0)
	goto L49
L48:
	;
	v208 = int32(_a_F_transformSetOperationTree_1)
	goto L49
L49:
	;
	if v46 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v211 = int32(_a_F_transformSetOperationTree_2)
	goto L52
L51:
	;
	v211 = v208
	goto L52
L52:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v216 = F_transformSetOperationTree(m, l0, v212, int32(0), v21+int32(8))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v216
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v221 = m.G0
	v223 = v221 - int32(96)
	m.G0 = v223
	if v57 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v226 = v225
	goto L56
L55:
	;
	v226 = v5
	goto L56
L56:
	;
	if v219 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	m.G0 = v223 + int32(96)
	v584 = v39
	goto L9
L58:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v231 = v229
	goto L60
L59:
	;
	v231 = int32(0)
	goto L60
L60:
	;
	if v231 == v226 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if l3 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L122
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L66
L65:
	;
	goto L66
L66:
	;
	v235 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+28)) = v235
	*(*int64)(unsafe.Add(mBase, uint32(v39)+20)) = v235
	v246 = int32(0)
	goto L67
L67:
	;
	v257 = int32(0)
	if v57 == v257 {
		v267 = v257
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v219 == int32(0) {
		goto L57
	} else {
		goto L72
	}
L70:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v261 <= v246 {
		v267 = int32(0)
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v267 = v263 + v246<<(uint(int32(2))%32)
	goto L69
L72:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if base.B2i32(v267 == int32(0))|base.B2i32(v272 <= v246) != 0 {
		goto L57
	} else {
		goto L73
	}
L73:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	if v275 == int32(0) {
		goto L57
	} else {
		goto L74
	}
L74:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275+v246<<(uint(int32(2))%32))))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	v285 = F_exprType(m, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v287 = F_exprType(m, v282)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+76)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v223)+72)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v223)+28)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v223)+24)) = v282
	v297 = F_list_make2_impl(m, v223+int32(28), v223+int32(24))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v301 = F_select_common_type(m, l0, v297, v211, v223+int32(80))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v223)+80))
	v304 = F_exprLocation(m, v303)
	mBase = m.M
	if v285 != int32(705) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v287 != int32(705) {
		goto L87
	} else {
		goto L88
	}
L80:
	;
	v307 = F_coerce_to_common_type(m, l0, v284, v301, v211)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	if base.Ui32(int32(1)) < base.Ui32(v309-int32(7)) {
		v317 = v284
		goto L79
	} else {
		goto L84
	}
L83:
	;
	v317 = v307
	goto L79
L84:
	;
	v314 = F_coerce_to_common_type(m, l0, v284, v301, v211)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283)+4)) = v314
	v317 = v314
	goto L79
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+64)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v223)+68)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v223)+20)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v223)+16)) = v330
	v339 = F_list_make2_impl(m, v223+int32(20), v223+int32(16))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L93
	}
L87:
	;
	v320 = F_coerce_to_common_type(m, l0, v282, v301, v211)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if base.Ui32(int32(1)) < base.Ui32(v322-int32(7)) {
		v330 = v282
		goto L86
	} else {
		goto L91
	}
L90:
	;
	v330 = v320
	goto L86
L91:
	;
	v327 = F_coerce_to_common_type(m, l0, v282, v301, v211)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v281)+4)) = v327
	v330 = v327
	goto L86
L93:
	;
	v341 = F_select_common_typmod(m, v339, v301)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+56)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v223)+60)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v223)+12)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v223)+8)) = v330
	v351 = F_list_make2_impl(m, v223+int32(12), v223+int32(8))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v353 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)))
	v358 = v356
	goto L98
L97:
	;
	v358 = int32(0)
	goto L98
L98:
	;
	v361 = F_select_common_collation(m, l0, v351, v358&int32(1))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v364 = F_lappend_oid(m, v363, v301)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v364
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v368 = F_lappend_int(m, v367, v341)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = v368
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	v372 = F_lappend_oid(m, v371, v361)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+28)) = v372
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v375 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if l3 != 0 {
		goto L116
	} else {
		goto L117
	}
L104:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)))
	if v378 != 0 {
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v380 = v223 + int32(36)
	*(*int32)(unsafe.Add(mBase, uint32(v380)+12)) = int32(524)
	*(*int32)(unsafe.Add(mBase, uint32(v380)+4)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v380)+16)) = v380
	v386 = int32(_a_F_transformSetOperationTree_3)
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_transformSetOperationTree[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+8)) = v387
	*(*int32)(unsafe.Add(mBase, _c_F_transformSetOperationTree[0])) = v223 + int32(44)
	goto L108
L107:
	;
	goto L106
L108:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v395 = F_palloc0(m, int32(20))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = int32(106)
	v399 = int32(0)
	F_get_sort_group_operators(m, v301, v399, int32(1), v399, v223+int32(92), v223+int32(88), v399, v223+int32(87))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v411 = int32(0)
	if base.B2i32(v45&int32(1) == v411)|base.B2i32(v301 != int32(2287))&base.B2i32(v301 != int32(2249)) == v411 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v421 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+87)) = uint8(v421)
	goto L113
L112:
	;
	goto L113
L113:
	;
	v423 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+4)) = v423
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v223)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v395)+8)) = v425
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v223)+92))
	*(*uint16)(unsafe.Add(mBase, uint32(v395)+16)) = uint16(v423)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+12)) = v427
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+87)))
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+18)) = uint8(v431)
	v433 = F_lappend(m, v393, v395)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v433
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v223+int32(36))+8))
	*(*int32)(unsafe.Add(mBase, _c_F_transformSetOperationTree[0])) = v439
	goto L115
L115:
	;
	goto L103
L116:
	;
	v445 = F_palloc0(m, int32(20))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v246 = v246 + int32(1)
	goto L67
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v445)+16)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v445)+12)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v445)+8)) = v341
	*(*int32)(unsafe.Add(mBase, uint32(v445)+4)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(v445))) = int32(57)
	v453 = int32(0)
	v456 = F_makeTargetEntry(m, v445, v453, v453, v453)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v459 = F_lappend(m, v458, v456)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v459
	goto L118
L122:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+32)) = v211
	F_errmsg(m, int32(_a_F_transformSetOperationTree_4), v223+int32(32))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v478 = F_exprLocation(m, v219)
	mBase = m.M
	F_parser_errposition(m, l0, v478)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_transformSetOperationTree_5), int32(2298), int32(_a_F_transformSetOperationTree_6))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	v492 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v490)+80)) = uint16(v492)
	*(*int32)(unsafe.Add(mBase, uint32(v490)+44)) = v492
	v496 = F_transformStmt(m, v490, l1)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_free_parsestate(m, v490)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v500 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v502 = F_contain_vars_of_level(m, v496, int32(1))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	if l3 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	if v502 != 0 {
		goto L3
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	v567 = int32(0)
	v570 = F_addRangeTableEntryForSubquery(m, l0, v496, v567, v567, v567)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L146
	}
L136:
	;
	v506 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v506
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v496)+76))
	if v508 == v506 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
	if v511 <= int32(0) {
		goto L135
	} else {
		goto L138
	}
L138:
	;
	v516 = int32(0)
	v524 = v5
	goto L139
L139:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v508)+12))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v533+v516<<(uint(int32(2))%32))))
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537)+26)))
	if v538 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	goto L135
L141:
	;
	v541 = F_lappend(m, v524, v537)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	v544 = v524
	goto L143
L143:
	;
	v546 = v516 + int32(1)
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
	if v546 < v547 {
		v516 = v546
		v524 = v544
		goto L139
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v541
	v544 = v541
	goto L143
L145:
	;
	goto L140
L146:
	;
	v573 = F_palloc0(m, int32(8))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v573))) = int32(63)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v570)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v573)+4)) = v577
	v584 = v573
	goto L9
L148:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errmsg(m, int32(_a_F_transformSetOperationTree_7), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v613 = F_exprLocation(m, v612)
	mBase = m.M
	F_parser_errposition(m, l0, v613)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_transformSetOperationTree_5), int32(2115), int32(_a_F_transformSetOperationTree_8))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)+12))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v629)))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)+8))
	v633 = v631 - int32(1)
	if base.Ui32(v633) <= base.Ui32(int32(3)) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v640
	F_errmsg(m, int32(_a_F_transformSetOperationTree_9), v21)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L159
	}
L156:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v633<<(uint(int32(2))%32))+uint32(_c_F_transformSetOperationTree[1])))
	v640 = v638
	goto L158
L157:
	;
	v640 = int32(_a_F_transformSetOperationTree_10)
	goto L158
L158:
	;
	goto L155
L159:
	;
	F_errfinish(m, int32(_a_F_transformSetOperationTree_5), int32(2125), int32(_a_F_transformSetOperationTree_8))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(_a_F_transformSetOperationTree_11))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errmsg(m, int32(_a_F_transformSetOperationTree_12), int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v662 = F_locate_var_of_level(m, v496, int32(1))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_parser_errposition(m, l0, v662)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	F_errfinish(m, int32(_a_F_transformSetOperationTree_5), int32(2185), int32(_a_F_transformSetOperationTree_8))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_transformWithClause(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
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
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int64
	_ = v325
	var v327 int32
	_ = v327
	var v329 int64
	_ = v329
	var v331 int32
	_ = v331
	var v333 int64
	_ = v333
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v693 int32
	_ = v693
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v160 != 0 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = v3
	goto L4
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v34 = int32(2)
	v36 = v33 + v29<<(uint(v34)%32)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v39 = v36 + int32(4)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v47 = base.B2i32(base.Ui32(v39) < base.Ui32(v42+v43<<(uint(v34)%32)))
	if base.Ui32(v39) < base.Ui32(v42+v43<<(uint(v34)%32)) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L1
L6:
	;
	v135 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+36)) = v135
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+32)) = uint8(v135)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v140 != int32(141) {
		goto L33
	} else {
		goto L34
	}
L7:
	;
	v48 = v39
	goto L9
L8:
	;
	v48 = int32(0)
	goto L9
L9:
	;
	if base.Ui32(v39) < base.Ui32(v42+v43<<(uint(v34)%32)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v52 = (v48 - v42) >> (uint(int32(2)) % 32)
	goto L12
L11:
	;
	v52 = v43
	goto L12
L12:
	;
	if v43 <= v52 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v57 = v52
	goto L14
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v42+v57<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if base.B2i32(v73 == int32(0))|base.B2i32(v73 != v76) != 0 {
		v94 = v73
		v95 = v76
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	if v94-v95 != 0 {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	v79 = v54
	v80 = v70
	goto L19
L19:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	if v84 == int32(0) {
		v94 = v84
		v95 = v83
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v94 = v84
	v95 = v83
	goto L17
L21:
	;
	v87 = int32(1)
	if v84 == v83 {
		v79 = v79 + v87
		v80 = v80 + v87
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v98 = v57 + int32(1)
	if v43 != v98 {
		v57 = v98
		goto L14
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L15
L26:
	;
	goto L6
L27:
	;
	return int32(0)
L28:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v109
	F_errmsg(m, int32(_a_F_transformWithClause_0), v14+int32(32))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
	F_parser_errposition(m, l0, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_transformWithClause_1), int32(139), int32(_a_F_transformWithClause_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v143 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v143)
	goto L35
L34:
	;
	goto L35
L35:
	;
	v146 = v29 + int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v146 < v147 {
		v29 = v146
		goto L4
	} else {
		goto L36
	}
L36:
	;
	goto L5
L37:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	m.G0 = v14 + int32(80)
	return v693
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l0
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v162 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v642 = F_list_copy(m, v641)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L27
	} else {
		goto L153
	}
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v165 = v163
	goto L43
L42:
	;
	v165 = int32(0)
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v165
	v169 = F_palloc0(m, v165*int32(12))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L27
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v169
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v172 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if int32(0) < v173 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v228 = v165
	goto L47
L47:
	;
	if v228 <= int32(0) {
		goto L37
	} else {
		goto L54
	}
L48:
	;
	v179 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v228 = v216
	goto L47
L51:
	;
	v189 = v179 * int32(12)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+v179<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v189+v190))) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v198+v189)+4)) = v179
	v202 = v179 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v202 < v203 {
		v179 = v202
		goto L51
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	goto L52
L54:
	;
	v234 = int32(0)
	goto L55
L55:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v243+v234*int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v234
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247)+16))
	v254 = F_makeDependencyGraphWalker(m, v251, v14+int32(36))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L27
	} else {
		goto L57
	}
L56:
	;
	v260 = int32(0)
	if v258 <= v260 {
		goto L37
	} else {
		goto L59
	}
L57:
	;
	v257 = v234 + int32(1)
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v257 < v258 {
		v234 = v257
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v266 = v260
	goto L60
L60:
	;
	v278 = v266
	goto L63
L61:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v376 <= int32(0) {
		goto L37
	} else {
		goto L83
	}
L62:
	;
	if v278 != v266 {
		goto L72
	} else {
		goto L73
	}
L63:
	;
	v289 = v263 + v278*int32(12)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
	if v290 == int32(0) {
		goto L62
	} else {
		goto L65
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L27
	} else {
		goto L67
	}
L65:
	;
	v294 = v278 + int32(1)
	if v294 != v258 {
		v278 = v294
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L27
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(_a_F_transformWithClause_3), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L27
	} else {
		goto L69
	}
L69:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v263+v266*int32(12))))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+28))
	F_parser_errposition(m, v264, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L27
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_transformWithClause_1), int32(883), int32(_a_F_transformWithClause_4))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L27
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v322 = v263 + v266*int32(12)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v323
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v322)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v325
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v322)+8)) = v327
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v289)))
	*(*int64)(unsafe.Add(mBase, uint32(v322))) = v329
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+8)) = v331
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v14)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v289))) = v333
	goto L74
L73:
	;
	goto L74
L74:
	;
	v337 = v266 + int32(1)
	if v337 < v258 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v344 = v337
	goto L78
L76:
	;
	goto L77
L77:
	;
	if v337 != v258 {
		v266 = v337
		goto L60
	} else {
		goto L82
	}
L78:
	;
	v355 = v263 + v344*int32(12)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+8))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v263+v266*int32(12))+4))
	v358 = F_bms_del_member(m, v356, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L27
	} else {
		goto L80
	}
L79:
	;
	goto L77
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355)+8)) = v358
	v362 = v344 + int32(1)
	if v362 != v258 {
		v344 = v362
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	goto L61
L83:
	;
	v382 = v376
	v384 = int32(0)
	goto L91
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L27
	} else {
		goto L150
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L27
	} else {
		goto L145
	}
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L27
	} else {
		goto L140
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L27
	} else {
		goto L135
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L27
	} else {
		goto L130
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L27
	} else {
		goto L125
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L27
	} else {
		goto L120
	}
L91:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v391+v384*int32(12))))
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+32)))
	if v396 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	if v444 <= int32(0) {
		goto L37
	} else {
		goto L110
	}
L93:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v395)+16))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	if v400 != int32(141) {
		goto L90
	} else {
		goto L96
	}
L94:
	;
	v444 = v382
	goto L95
L95:
	;
	v446 = v384 + int32(1)
	if v446 < v444 {
		v382 = v444
		v384 = v446
		goto L91
	} else {
		goto L109
	}
L96:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v399)+68))
	if v403 != int32(1) {
		goto L89
	} else {
		goto L97
	}
L97:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v399)+64))
	if v406 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v384
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	v415 = F_checkWellFormedRecursionWalker(m, v412, v14+int32(36))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L27
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v399)+44))
	if v417 != 0 {
		goto L88
	} else {
		goto L102
	}
L101:
	;
	goto L100
L102:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v399)+48))
	if v418 != 0 {
		goto L87
	} else {
		goto L103
	}
L103:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v399)+52))
	if v419 != 0 {
		goto L86
	} else {
		goto L104
	}
L104:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v399)+60))
	if v420 != 0 {
		goto L85
	} else {
		goto L105
	}
L105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = int32(1)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v399)+76))
	v428 = v14 + int32(36)
	v429 = F_checkWellFormedRecursionWalker(m, v426, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L27
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v384
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v399)+80))
	v437 = F_checkWellFormedRecursionWalker(m, v436, v428)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L27
	} else {
		goto L107
	}
L107:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	if v439 != int32(1) {
		goto L84
	} else {
		goto L108
	}
L108:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v444 = v442
	goto L95
L109:
	;
	goto L92
L110:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v454 = int32(0)
	v456 = v450
	goto L111
L111:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v463+v454*int32(12))))
	v468 = F_lappend(m, v456, v467)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L27
	} else {
		goto L113
	}
L112:
	;
	v475 = int32(0)
	if v473 <= v475 {
		goto L37
	} else {
		goto L115
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v468
	v472 = v454 + int32(1)
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v472 < v473 {
		v454 = v472
		v456 = v468
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v480 = v475
	goto L116
L116:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v489+v480*int32(12))))
	F_analyzeCTE(m, l0, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L27
	} else {
		goto L118
	}
L117:
	;
	goto L37
L118:
	;
	v497 = v480 + int32(1)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v497 < v498 {
		v480 = v497
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L27
	} else {
		goto L121
	}
L121:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v507
	F_errmsg(m, int32(_a_F_transformWithClause_5), v14+int32(16))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L27
	} else {
		goto L122
	}
L122:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v395)+28))
	F_parser_errposition(m, v514, v515)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L27
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_transformWithClause_1), int32(936), int32(_a_F_transformWithClause_6))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L27
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L27
	} else {
		goto L126
	}
L126:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v530
	F_errmsg(m, int32(_a_F_transformWithClause_7), v14)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L27
	} else {
		goto L127
	}
L127:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v395)+28))
	F_parser_errposition(m, v535, v536)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L27
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_transformWithClause_1), int32(944), int32(_a_F_transformWithClause_6))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L27
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L27
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_transformWithClause_8), int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L27
	} else {
		goto L132
	}
L132:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v399)+44))
	v557 = F_exprLocation(m, v556)
	mBase = m.M
	F_parser_errposition(m, v555, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L27
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_transformWithClause_1), int32(979), int32(_a_F_transformWithClause_6))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L27
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L27
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_transformWithClause_9), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L27
	} else {
		goto L137
	}
L137:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v399)+48))
	v578 = F_exprLocation(m, v577)
	mBase = m.M
	F_parser_errposition(m, v576, v578)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L27
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_transformWithClause_1), int32(985), int32(_a_F_transformWithClause_6))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L27
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L27
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(_a_F_transformWithClause_10), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L27
	} else {
		goto L142
	}
L142:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v399)+52))
	v599 = F_exprLocation(m, v598)
	mBase = m.M
	F_parser_errposition(m, v597, v599)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L27
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_transformWithClause_1), int32(991), int32(_a_F_transformWithClause_6))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L27
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L145:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L27
	} else {
		goto L146
	}
L146:
	;
	F_errmsg(m, int32(_a_F_transformWithClause_11), int32(0))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L27
	} else {
		goto L147
	}
L147:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v399)+60))
	v620 = F_exprLocation(m, v619)
	mBase = m.M
	F_parser_errposition(m, v618, v620)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L27
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_transformWithClause_1), int32(997), int32(_a_F_transformWithClause_6))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L27
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errmsg_internal(m, int32(_a_F_transformWithClause_12), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L27
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_transformWithClause_1), int32(1019), int32(_a_F_transformWithClause_6))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L27
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v642
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v645 == int32(0) {
		goto L37
	} else {
		goto L154
	}
L154:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v645)+4))
	if v648 <= int32(0) {
		goto L37
	} else {
		goto L155
	}
L155:
	;
	v654 = int32(0)
	goto L156
L156:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v645)+12))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v663+v654<<(uint(int32(2))%32))))
	F_analyzeCTE(m, l0, v667)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L27
	} else {
		goto L158
	}
L157:
	;
	goto L37
L158:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v671 = F_lappend(m, v670, v667)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L27
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v671
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v675 = F_list_delete_first(m, v674)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L27
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v675
	v679 = v654 + int32(1)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v645)+4))
	if v679 < v680 {
		v654 = v679
		goto L156
	} else {
		goto L161
	}
L161:
	;
	goto L157
}
func F_transientrel_receive(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+188))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
	m.T0[v9].(func(*base.Module, int32, int32, int32, int32, int32))(m, v4, l0, v5, v6, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return int32(1)
	}
}
func F_transientrel_startup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_table_open(m, v4, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v6
		v10 = F_GetCurrentCommandId(m, int32(1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v10
			v15 = F_GetBulkInsertState(m)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v15
				return
			}
		}
	}
}
func F_trgm2int(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	return v2 | (v3<<(uint(int32(8))%32) | v6<<(uint(int32(16))%32))
}
func F_trueTriConsistentFn(m *base.Module, l0 int32) int32 {
	return int32(1)
}
func F_tsm_system_rows_handler(m *base.Module, l0 int32) int64 {
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14392(m, l0, int32(256), int32(_a_F_tsm_system_rows_handler_0), int32(_a_F_tsm_system_rows_handler_1), int32(_a_F_tsm_system_rows_handler_2), int32(_a_F_tsm_system_rows_handler_3), int32(_a_F_tsm_system_rows_handler_4), int32(20))
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		return v9
	}
}
func F_tsq_mcontains(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
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
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v381 int32
	_ = v381
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = F_palloc_mul(m, int32(4), v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(0) < v19 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = v13 + int32(8)
	v27 = v23
	v28 = v19
	v31 = v2
	v35 = v2
	goto L6
L4:
	;
	v74 = v2
	goto L5
L5:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v82 = F_palloc_mul(m, int32(4), v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v37 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v74 = v64
	goto L5
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	v42 = v40 & int32(4095)
	v45 = F_palloc(m, v42+int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v62 = v28
	v64 = v31
	goto L10
L10:
	;
	v68 = v35 + int32(1)
	if v68 < v62 {
		v27 = v27 + int32(12)
		v28 = v62
		v31 = v64
		v35 = v68
		goto L6
	} else {
		goto L15
	}
L11:
	;
	if v42 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	base.MemoryCopy(m, v45, v23+v14*int32(12)+int32(base.Ui32(v47)>>(uint(int32(12))%32)), v42)
	goto L14
L13:
	;
	goto L14
L14:
	;
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42+v45))) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v15+v31<<(uint(int32(2))%32)))) = v45
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v62 = v61
	v64 = v31 + int32(1)
	goto L10
L15:
	;
	goto L7
L16:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v84 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v88 = v11 + int32(8)
	v93 = v88
	v94 = int32(0)
	v95 = v84
	v98 = v2
	goto L20
L18:
	;
	v141 = v2
	goto L19
L19:
	;
	F_pg_qsort(m, v15, v74, int32(4), int32(1726))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L30
	}
L20:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v103 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v141 = v129
	goto L19
L22:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v108 = v106 & int32(4095)
	v111 = F_palloc(m, v108+int32(1))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v128 = v95
	v129 = v98
	goto L24
L24:
	;
	v134 = v94 + int32(1)
	if v134 < v128 {
		v93 = v93 + int32(12)
		v94 = v134
		v95 = v128
		v98 = v129
		goto L20
	} else {
		goto L29
	}
L25:
	;
	if v108 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	base.MemoryCopy(m, v111, v88+v81*int32(12)+int32(base.Ui32(v113)>>(uint(int32(12))%32)), v108)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v108+v111))) = uint8(v119)
	*(*int32)(unsafe.Add(mBase, uint32(v82+v98<<(uint(int32(2))%32)))) = v111
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v128 = v127
	v129 = v98 + int32(1)
	goto L24
L29:
	;
	goto L21
L30:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v74) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v154 = int32(1)
	v155 = int32(0)
	goto L34
L32:
	;
	v218 = v74
	goto L33
L33:
	;
	F_pg_qsort(m, v82, v141, int32(4), int32(1726))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L47
	}
L34:
	;
	v164 = int32(2)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v15+v154<<(uint(v164)%32))))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v15+v155<<(uint(v164)%32))))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	if base.B2i32(v174 == int32(0))|base.B2i32(v174 != v177) != 0 {
		v195 = v174
		v196 = v177
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v218 = v207 + int32(1)
	goto L33
L36:
	;
	v210 = v154 + int32(1)
	if v210 != v74 {
		v154 = v210
		v155 = v207
		goto L34
	} else {
		goto L46
	}
L37:
	;
	if v195-v196 == int32(0) {
		v207 = v155
		goto L36
	} else {
		goto L44
	}
L38:
	;
	goto L37
L39:
	;
	v180 = v167
	v181 = v171
	goto L40
L40:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+1)))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
	if v185 == int32(0) {
		v195 = v185
		v196 = v184
		goto L38
	} else {
		goto L42
	}
L41:
	;
	v195 = v185
	v196 = v184
	goto L38
L42:
	;
	v188 = int32(1)
	if v185 == v184 {
		v180 = v180 + v188
		v181 = v181 + v188
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v201 = v155 + int32(1)
	if v154 == v201 {
		v207 = v154
		goto L36
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v201<<(uint(int32(2))%32)))) = v167
	v207 = v201
	goto L36
L46:
	;
	goto L35
L47:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v141) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v232 = int32(1)
	v233 = int32(0)
	goto L51
L49:
	;
	v297 = v141
	goto L50
L50:
	;
	if v218 < v297 {
		goto L64
	} else {
		goto L65
	}
L51:
	;
	v242 = int32(2)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v82+v232<<(uint(v242)%32))))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v82+v233<<(uint(v242)%32))))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if base.B2i32(v252 == int32(0))|base.B2i32(v252 != v255) != 0 {
		v273 = v252
		v274 = v255
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v297 = v285 + int32(1)
	goto L50
L53:
	;
	v288 = v232 + int32(1)
	if v288 != v141 {
		v232 = v288
		v233 = v285
		goto L51
	} else {
		goto L63
	}
L54:
	;
	if v273-v274 == int32(0) {
		v285 = v233
		goto L53
	} else {
		goto L61
	}
L55:
	;
	goto L54
L56:
	;
	v258 = v245
	v259 = v249
	goto L57
L57:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+1)))
	if v263 == int32(0) {
		v273 = v263
		v274 = v262
		goto L55
	} else {
		goto L59
	}
L58:
	;
	v273 = v263
	v274 = v262
	goto L55
L59:
	;
	v266 = int32(1)
	if v263 == v262 {
		v258 = v258 + v266
		v259 = v259 + v266
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v279 = v233 + int32(1)
	if v232 == v279 {
		v285 = v232
		goto L53
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82+v279<<(uint(int32(2))%32)))) = v245
	v285 = v279
	goto L53
L63:
	;
	goto L52
L64:
	;
	return int64(0)
L65:
	;
	v303 = int32(0)
	if v297 <= v303 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	return int64(1)
L67:
	;
	goto L68
L68:
	;
	v309 = v303
	v311 = int32(0)
	goto L69
L69:
	;
	if v218 <= v309 {
		v369 = v309
		goto L71
	} else {
		goto L72
	}
L70:
	;
	return int64(1)
L71:
	;
	if v369 == v218 {
		goto L64
	} else {
		goto L84
	}
L72:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v82+v311<<(uint(int32(2))%32))))
	v324 = v309
	goto L73
L73:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v15+v324<<(uint(int32(2))%32))))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337))))
	if base.B2i32(v340 == int32(0))|base.B2i32(v340 != v343) != 0 {
		v361 = v340
		v362 = v343
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L64
L75:
	;
	if v361-v362 == int32(0) {
		v369 = v324
		goto L71
	} else {
		goto L82
	}
L76:
	;
	goto L75
L77:
	;
	v346 = v323
	v347 = v337
	goto L78
L78:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+1)))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+1)))
	if v351 == int32(0) {
		v361 = v351
		v362 = v350
		goto L76
	} else {
		goto L80
	}
L79:
	;
	v361 = v351
	v362 = v350
	goto L76
L80:
	;
	v354 = int32(1)
	if v351 == v350 {
		v346 = v346 + v354
		v347 = v347 + v354
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v367 = v324 + int32(1)
	if v367 != v218 {
		v324 = v367
		goto L73
	} else {
		goto L83
	}
L83:
	;
	goto L74
L84:
	;
	v381 = v311 + int32(1)
	if v381 != v297 {
		v309 = v369
		v311 = v381
		goto L69
	} else {
		goto L85
	}
L85:
	;
	goto L70
}
func F_tsqueryout(m *base.Module, l0 int32) int64 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == int32(0) {
		v15 = F_palloc(m, int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v19)
			v50 = v15
			m.G0 = v8 + int32(32)
			return base.I64_extend_i32_u(v50)
		}
	} else {
		v21 = int32(32)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v21
		v24 = v10 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v24
		v28 = F_palloc_mul(m, int32(1), v21)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v28
			v32 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v32)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v35 = int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v24 + v34*v35
			F_infix_1(m, v8+v35, int32(-1), v32)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v45 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int64(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
						v50 = v49
						m.G0 = v8 + int32(32)
						return base.I64_extend_i32_u(v50)
					}
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					v50 = v49
					m.G0 = v8 + int32(32)
					return base.I64_extend_i32_u(v50)
				}
			}
		}
	}
}
func F_tstoreReceiveSlot_detoast(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int64
	_ = v60
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	v3 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v13 < v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	m.T0[v16].(func(*base.Module, int32, int32))(m, l0, v12)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if int32(0) < v12 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v27 = v3
	v29 = v3
	goto L9
L7:
	;
	v77 = v3
	goto L8
L8:
	;
	v83 = int32(_a_F_tstoreReceiveSlot_detoast_0)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_tstoreReceiveSlot_detoast[0]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_tstoreReceiveSlot_detoast[0])) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_tuplestore_putvalues(m, v88, v11, v89, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L18
	}
L9:
	;
	v36 = v27 << (uint(int32(3)) % 32)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v36+v37)))
	v40 = v36 + (v11 + int32(28))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+6)))
	if v41&int32(4) != 0 {
		v65 = v29
		v66 = v39
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v77 = v65
	goto L8
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v67+v36))) = v66
	v71 = v27 + int32(1)
	if v71 != v12 {
		v27 = v71
		v29 = v65
		goto L9
	} else {
		goto L17
	}
L12:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+2)))
	if v44 != int32(_a_F_tstoreReceiveSlot_detoast_1) {
		v65 = v29
		v66 = v39
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v27))))
	if v49 != 0 {
		v65 = v29
		v66 = v39
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v50 = base.I32_wrap_i64(v39)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v51 != int32(1) {
		v65 = v29
		v66 = v39
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v54 = F_detoast_external_attr(m, v50)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v60 = base.I64_extend_i32_u(v54)
	*(*int64)(unsafe.Add(mBase, uint32(v56+v29<<(uint(int32(3))%32)))) = v60
	v65 = v29 + int32(1)
	v66 = v60
	goto L11
L17:
	;
	goto L10
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tstoreReceiveSlot_detoast[0])) = v84
	if int32(0) < v77 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v100 = int32(0)
	goto L22
L20:
	;
	goto L21
L21:
	;
	return int32(1)
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v100<<(uint(int32(3))%32))))
	F_pfree(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	v116 = v100 + int32(1)
	if v116 != v77 {
		v100 = v116
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
}
func F_tsvectorin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var __phi260 int32
	_ = __phi260
	var v266 int32
	_ = v266
	var __phi266 int32
	_ = __phi266
	var v267 int32
	_ = v267
	var __phi267 int32
	_ = __phi267
	var v274 int32
	_ = v274
	var __phi274 int32
	_ = __phi274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v420 int32
	_ = v420
	var v431 int32
	_ = v431
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int64
	_ = v463
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v593 int32
	_ = v593
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v621 int32
	_ = v621
	var v636 int32
	_ = v636
	var v653 int32
	_ = v653
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v842 int64
	_ = v842
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	v2 = int32(0)
	v18 = int64(0)
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = F_init_tsvector_parser(m, v25, v2, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v34 = F_palloc_mul(m, int32(12), int32(64))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v38 = F_palloc_mul(m, int32(1), int32(256))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v49 = F_gettoken_tsvector(m, v28, v19+int32(-4), v19+int32(-8), v19+int32(-12), v19+int32(-16), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L157
	}
L6:
	;
	m.G0 = v21 - int32(-64)
	return v842
L7:
	;
	if v49 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v52 = v38
	v55 = int32(64)
	v58 = int32(256)
	v61 = v2
	v62 = v34
	v64 = v38
	goto L11
L9:
	;
	v230 = v2
	v231 = v34
	v233 = v38
	goto L10
L10:
	;
	F_close_tsvector_parser(m, v28)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L49
	}
L11:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if int32(2048) <= v69 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v230 = v208
	v231 = v124
	v233 = v166
	goto L10
L13:
	;
	v72 = F_errsave_start(m, v27)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v91 = v52 - v64
	if int32(_a_F_tsvectorin_0) <= v91 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	if v72 == int32(0) {
		v842 = v18
		goto L6
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(2047)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v81
	F_errmsg(m, int32(_a_F_tsvectorin_1), v21)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errsave_finish(m, v27, int32(_a_F_tsvectorin_2), int32(215), int32(_a_F_tsvectorin_3))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v842 = v18
	goto L6
L21:
	;
	v94 = F_errsave_start(m, v27)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v55 <= v61 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	if v94 == int32(0) {
		v842 = v18
		goto L6
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = int32(_a_F_tsvectorin_4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v91
	F_errmsg(m, int32(_a_F_tsvectorin_5), v19+int32(-48))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errsave_finish(m, v27, int32(_a_F_tsvectorin_2), int32(221), int32(_a_F_tsvectorin_3))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v842 = v18
	goto L6
L29:
	;
	v117 = F_repalloc(m, v62, v55*int32(24))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	v122 = v69
	v123 = v55
	v124 = v62
	goto L31
L31:
	;
	if v58 <= v122+v91 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v122 = v121
	v123 = v55 << (uint(int32(1)) % 32)
	v124 = v117
	goto L31
L33:
	;
	v134 = v58
	v140 = v64
	goto L36
L34:
	;
	v154 = v52
	v155 = v122
	v160 = v58
	v166 = v64
	goto L35
L35:
	;
	v171 = int32(12)
	v173 = v124 + v61*v171
	v174 = int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v155<<(uint(v174)%32)&int32(4094) | v178&v174 | v91<<(uint(v171)%32)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v186 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v146 = v134 << (uint(int32(1)) % 32)
	v147 = F_repalloc(m, v140, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v154 = v91 + v147
	v155 = v149
	v160 = v146
	v166 = v147
	goto L35
L38:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	if v146 <= v91+v149 {
		v134 = v146
		v140 = v147
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	base.MemoryCopy(m, v154, v187, v186)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	if v192 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = v205
	v208 = v61 + int32(1)
	v218 = F_gettoken_tsvector(m, v28, v19+int32(-4), v19+int32(-8), v19+int32(-12), v19+int32(-16), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L47
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v189 | int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v205 = v198
	goto L43
L45:
	;
	goto L46
L46:
	;
	v199 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v189 & int32(-2)
	v205 = v199
	goto L43
L47:
	;
	if v218 != 0 {
		v52 = v190 + v154
		v55 = v123
		v58 = v160
		v61 = v208
		v62 = v124
		v64 = v166
		goto L11
	} else {
		goto L48
	}
L48:
	;
	goto L12
L49:
	;
	if v27 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v230 <= int32(0) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v242 != int32(453) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)))
	if v245 != int32(1) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v248 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v248)
	v842 = v18
	goto L6
L54:
	;
	v709 = v700 << (uint(int32(2)) % 32)
	v712 = v709 + v698 + int32(8)
	v713 = F_palloc0(m, v712)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L139
	}
L55:
	;
	v698 = v2
	v700 = v230
	goto L54
L56:
	;
	goto L57
L57:
	;
	if v230 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v543 = int32(1)
	v547 = int32(base.Ui32(v542)>>(uint(v543)%32))&int32(2047) + v538
	if v542&v543 != 0 {
		goto L115
	} else {
		goto L116
	}
L59:
	;
	v524 = v231
	v538 = v2
	goto L58
L60:
	;
	goto L61
L61:
	;
	F_qsort_arg(m, v231, v230, int32(12), int32(1732), v233)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	__phi260 = v231
	__phi266 = v231
	__phi267 = v231 + int32(12)
	__phi274 = v2
	v260 = __phi260
	v266 = __phi266
	v267 = __phi267
	v274 = __phi274
	goto L63
L63:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
	v279 = int32(1)
	v281 = int32(2047)
	v282 = int32(base.Ui32(v278)>>(uint(v279)%32)) & v281
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v287 = int32(base.Ui32(v283)>>(uint(v279)%32)) & v281
	if v282 == v287 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v524 = v500
	v538 = v514
	goto L58
L65:
	;
	v518 = int32(12)
	v519 = v267 + v518
	v522 = base.I32_div_s(v519-v231, v518)
	if v522 < v230 {
		__phi260 = v500
		__phi266 = v267
		__phi267 = v519
		__phi274 = v514
		v260 = __phi260
		v266 = __phi266
		v267 = __phi267
		v274 = __phi274
		goto L63
	} else {
		goto L114
	}
L66:
	;
	if v278&int32(1) == int32(0) {
		v500 = v260
		v514 = v274
		goto L65
	} else {
		goto L105
	}
L67:
	;
	v289 = int32(12)
	v291 = v233 + int32(base.Ui32(v278)>>(uint(v289)%32))
	v294 = v233 + int32(base.Ui32(v283)>>(uint(v289)%32))
	if v282 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	goto L69
L69:
	;
	v342 = v287 + v274
	if v283&int32(1) != 0 {
		goto L84
	} else {
		goto L85
	}
L70:
	;
	if v339 == int32(0) {
		goto L66
	} else {
		goto L83
	}
L71:
	;
	v339 = int32(0)
	goto L70
L72:
	;
	goto L73
L73:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if v300 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v301 = v291
	v302 = v294
	v303 = v282
	v304 = v300
	goto L78
L75:
	;
	v327 = v294
	v331 = int32(0)
	goto L76
L76:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	v339 = v331 - v332
	goto L70
L77:
	;
	v327 = v322
	v331 = v324
	goto L76
L78:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if base.B2i32(v304 != v306)|base.B2i32(v306 == int32(0)) != 0 {
		v322 = v302
		v324 = v304
		goto L77
	} else {
		goto L80
	}
L79:
	;
	v322 = v316
	v324 = int32(0)
	goto L77
L80:
	;
	v312 = v303 - int32(1)
	if v312 == int32(0) {
		v322 = v302
		v324 = v304
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v315 = int32(1)
	v316 = v302 + v315
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+1)))
	if v317 != 0 {
		v301 = v301 + v315
		v302 = v316
		v303 = v312
		v304 = v317
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	goto L69
L84:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	if int32(2) <= v345 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v454 = v342
	goto L86
L86:
	;
	v459 = v260 + int32(12)
	if v260 == v266 {
		goto L102
	} else {
		goto L103
	}
L87:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	F_pg_qsort(m, v348, v345, int32(2), int32(1733))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	v420 = v345
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+8)) = v420
	v431 = int32(1)
	v454 = (v342+v431)&int32(-2) + v420<<(uint(v431)%32) + int32(2)
	goto L86
L90:
	;
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v348))))
	v357 = v355
	v358 = v348
	v363 = v348 + int32(2)
	goto L91
L91:
	;
	v374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v363))))
	v375 = int32(_a_F_tsvectorin_6)
	v376 = v374 & v375
	if v376 != v357&v375 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v420 = (v405 - v348 + int32(2)) >> (uint(int32(1)) % 32)
	goto L89
L93:
	;
	goto L92
L94:
	;
	v399 = v363 + int32(2)
	if (v399-v348)>>(uint(int32(1))%32) < v345 {
		v357 = v396
		v358 = v397
		v363 = v399
		goto L91
	} else {
		goto L101
	}
L95:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v358)+2)) = uint16(v374)
	v382 = v358 + int32(2)
	if int32(508) < v382-v348 {
		v405 = v382
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v388 = int32(14)
	if base.Ui32(int32(base.Ui32(v374)>>(uint(v388)%32))) <= base.Ui32(int32(base.Ui32(v357&int32(_a_F_tsvectorin_7))>>(uint(v388)%32))) {
		v396 = v357
		v397 = v358
		goto L94
	} else {
		goto L100
	}
L98:
	;
	if v376 != int32(_a_F_tsvectorin_6) {
		v396 = v374
		v397 = v382
		goto L94
	} else {
		goto L99
	}
L99:
	;
	v405 = v382
	goto L93
L100:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v358))) = uint16(v374)
	v396 = v374
	v397 = v358
	goto L94
L101:
	;
	v405 = v397
	goto L93
L102:
	;
	v500 = v459
	v514 = v454
	goto L65
L103:
	;
	goto L104
L104:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v459)+8)) = v461
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v267)))
	*(*int64)(unsafe.Add(mBase, uint32(v459))) = v463
	v500 = v459
	v514 = v454
	goto L65
L105:
	;
	if v283&int32(1) != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	v474 = v472 + v473
	v477 = F_repalloc(m, v471, v474<<(uint(int32(1))%32))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v283 | int32(1)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v266)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v260)+4)) = v496
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v260)+8)) = v498
	v500 = v260
	v514 = v274
	goto L65
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+4)) = v477
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v266)+20))
	v482 = v480 << (uint(int32(1)) % 32)
	if v482 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v266)+16))
	base.MemoryCopy(m, v477+v483<<(uint(int32(1))%32), v487, v482)
	goto L112
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+8)) = v474
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v266)+16))
	F_pfree(m, v490)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v500 = v260
	v514 = v274
	goto L65
L114:
	;
	goto L64
L115:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v524)+8))
	if int32(2) <= v550 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v653 = v547
	goto L117
L117:
	;
	v664 = int32(12)
	v667 = base.I32_div_s(v524-v231+v664, v664)
	if v653 < int32(_a_F_tsvectorin_0) {
		v698 = v653
		v700 = v667
		goto L54
	} else {
		goto L133
	}
L118:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	F_pg_qsort(m, v553, v550, int32(2), int32(1733))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	v621 = v550
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v524)+8)) = v621
	v636 = int32(1)
	v653 = (v547+v636)&int32(-2) + v621<<(uint(v636)%32) + int32(2)
	goto L117
L121:
	;
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v553))))
	v562 = v560
	v563 = v553
	v568 = v553 + int32(2)
	goto L122
L122:
	;
	v579 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v568))))
	v580 = int32(_a_F_tsvectorin_6)
	v581 = v579 & v580
	if v581 != v562&v580 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v621 = (v610 - v553 + int32(2)) >> (uint(int32(1)) % 32)
	goto L120
L124:
	;
	goto L123
L125:
	;
	v604 = v568 + int32(2)
	if (v604-v553)>>(uint(int32(1))%32) < v550 {
		v562 = v601
		v563 = v602
		v568 = v604
		goto L122
	} else {
		goto L132
	}
L126:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v563)+2)) = uint16(v579)
	v587 = v563 + int32(2)
	if int32(508) < v587-v553 {
		v610 = v587
		goto L124
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v593 = int32(14)
	if base.Ui32(int32(base.Ui32(v579)>>(uint(v593)%32))) <= base.Ui32(int32(base.Ui32(v562&int32(_a_F_tsvectorin_7))>>(uint(v593)%32))) {
		v601 = v562
		v602 = v563
		goto L125
	} else {
		goto L131
	}
L129:
	;
	if v581 != int32(_a_F_tsvectorin_6) {
		v601 = v579
		v602 = v587
		goto L125
	} else {
		goto L130
	}
L130:
	;
	v610 = v587
	goto L124
L131:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v563))) = uint16(v579)
	v601 = v579
	v602 = v563
	goto L125
L132:
	;
	v610 = v602
	goto L124
L133:
	;
	v670 = F_errsave_start(m, v27)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	if v670 == int32(0) {
		v842 = v18
		goto L6
	} else {
		goto L135
	}
L135:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = int32(_a_F_tsvectorin_4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v653
	F_errmsg(m, int32(_a_F_tsvectorin_5), v19+int32(-32))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errsave_finish(m, v27, int32(_a_F_tsvectorin_2), int32(275), int32(_a_F_tsvectorin_3))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v842 = v18
	goto L6
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v713)+4)) = v700
	*(*int32)(unsafe.Add(mBase, uint32(v713))) = v712 << (uint(int32(2)) % 32)
	if int32(0) < v700 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v722 = v713 + int32(8)
	v723 = v722 + v709
	v727 = int32(0)
	v740 = v2
	goto L143
L141:
	;
	goto L142
L142:
	;
	v842 = base.I64_extend_i32_u(v713)
	goto L6
L143:
	;
	v745 = v231 + v727*int32(12)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v745)))
	v750 = int32(base.Ui32(v746)>>(uint(int32(1))%32)) & int32(2047)
	if v750 != 0 {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	goto L142
L145:
	;
	base.MemoryCopy(m, v723+v740, v233+int32(base.Ui32(v746)>>(uint(int32(12))%32)), v750)
	goto L147
L146:
	;
	goto L147
L147:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v745)))
	v761 = v756&int32(4095) | v740<<(uint(int32(12))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v745))) = v761
	v763 = int32(1)
	v767 = int32(base.Ui32(v756)>>(uint(v763)%32))&int32(2047) + v740
	if v756&v763 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v745)+8))
	if int32(_a_F_tsvectorin_8) <= v773 {
		goto L5
	} else {
		goto L151
	}
L149:
	;
	v800 = v767
	v801 = v761
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v722+v727<<(uint(int32(2))%32)))) = v801
	v804 = v727 + int32(1)
	if v804 != v700 {
		v727 = v804
		v740 = v800
		goto L143
	} else {
		goto L156
	}
L151:
	;
	v776 = int32(1)
	v779 = (v767 + v776) & int32(-2)
	*(*uint16)(unsafe.Add(mBase, uint32(v723+v779))) = uint16(v773)
	v783 = v779 + int32(2)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v745)+8))
	v786 = v784 << (uint(v776) % 32)
	if v786 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v745)+4))
	base.MemoryCopy(m, v783+v723, v788, v786)
	goto L154
L153:
	;
	goto L154
L154:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v745)+8))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v745)+4))
	F_pfree(m, v791)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v745)))
	v800 = v790<<(uint(int32(1))%32) + v783
	v801 = v797
	goto L150
L156:
	;
	goto L144
L157:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorin_9), int32(0))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(_a_F_tsvectorin_2), int32(293), int32(_a_F_tsvectorin_3))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tt_process_call(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
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
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int64
	_ = v58
	v5 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == int32(0) {
		v58 = v5
		m.G0 = v8 + int32(48)
		return v58
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v11+v14*int32(12))))
		if v18 == int32(0) {
			v58 = v5
			m.G0 = v8 + int32(48)
			return v58
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
			v23 = v8 + int32(16)
			v25 = F_pg_sprintf(m, v23, int32(_a_F_tt_process_call_0), v8)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v23
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v34 = v30 + v31*int32(12)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v35
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v42 = F_BuildTupleFromCStrings(m, v39, v8+int32(36))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
					v45 = F_HeapTupleHeaderGetDatum(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int64(0)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
						F_pfree(m, v47)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
							F_pfree(m, v50)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int64(0)
							} else {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v53 + int32(1)
								v58 = v45
								m.G0 = v8 + int32(48)
								return v58
							}
						}
					}
				}
			}
		}
	}
}
func F_tt_setup_firstcall(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_lookup_ts_parser_cache(m, l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
		if v13 != 0 {
			v14 = int32(_a_F_tt_setup_firstcall_0)
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_tt_setup_firstcall[0]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_tt_setup_firstcall[0])) = v17
			v20 = F_palloc(m, int32(8))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
				v27 = F_OidFunctionCall1Coll(m, v24, v22, int64(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					*(*uint32)(unsafe.Add(mBase, uint32(v20)+4)) = uint32(v27)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v20
					v34 = F_get_call_result_type(m, l1, int32(0), v9+int32(12))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						if v34 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_tt_setup_firstcall_1), int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_tt_setup_firstcall_2), int32(69), int32(_a_F_tt_setup_firstcall_3))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v38
							v40 = F_TupleDescGetAttInMetadata(m, v38)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v40
								*(*int32)(unsafe.Add(mBase, _c_F_tt_setup_firstcall[0])) = v15
								m.G0 = v9 + int32(16)
								return
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
				F_errmsg_internal(m, int32(_a_F_tt_setup_firstcall_4), v9)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_tt_setup_firstcall_2), int32(57), int32(_a_F_tt_setup_firstcall_3))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
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
func F_tuplehash_insert_hash_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v84 int64
	_ = v84
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v220 int32
	_ = v220
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
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
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int64
	_ = v339
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int64
	_ = v389
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int64
	_ = v412
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
	var v472 int32
	_ = v472
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	goto L1
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v34) <= base.Ui32(v33) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L22
	} else {
		goto L97
	}
L3:
	;
	goto L2
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L1
L5:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v472)
	m.G0 = v17 + int32(16)
	return v462
L6:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v449 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v448 + v449
	*(*int32)(unsafe.Add(mBase, uint32(v438)+8)) = l1
	v453 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v438))) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v438)+4)) = v449
	v462 = v438
	v472 = v453
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L22
	} else {
		goto L94
	}
L8:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v36 == int64(4294967296) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v235 = int32(0)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v238 = v237 & l1
	v241 = v236 + v238*int32(12)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+4))
	if v242 == v235 {
		v438 = v241
		goto L6
	} else {
		goto L53
	}
L11:
	;
	v39 = int32(0)
	v41 = int64(2)
	v43 = v36 << (uint(int64(1)) % 64)
	if base.Ui64(v43) <= base.Ui64(v41) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v46 = v41
	goto L15
L14:
	;
	v46 = v43
	goto L15
L15:
	;
	v47 = int64(1)
	if v46&(v46-v47) == int64(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v57 = v46
	goto L18
L17:
	;
	v57 = v47 << (uint(int64(64)-base.I64_clz(v46)) % 64)
	goto L18
L18:
	;
	if base.Ui64(v57*int64(12)) < base.Ui64(int64(2147483647)) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v69 = F_MemoryContextAllocExtended(m, v64, base.I32_wrap_i64(v57)*int32(12), int32(5))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	goto L3
L22:
	;
	return int32(0)
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v69
	v74 = int64(1)
	if v57&(v57-v74) == int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v84 = v57
	goto L26
L25:
	;
	v84 = v74 << (uint(int64(64)-base.I64_clz(v57)) % 64)
	goto L26
L26:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v84*int64(12)) {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v84
	v92 = base.I32_wrap_i64(v84) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v92
	if v84 == int64(4294967296) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v101 = int32(-85899346)
	goto L30
L29:
	;
	v101 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v84), float64(0.9)))
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v101
	if v63 != int64(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v109 = v39
	goto L35
L32:
	;
	goto L33
L33:
	;
	F_pfree(m, v62)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L22
	} else {
		goto L52
	}
L34:
	;
	v138 = v133
	v139 = v39
	goto L40
L35:
	;
	v121 = v62 + v109*int32(12)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	if v122 != int32(1) {
		v133 = v109
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v133 = int32(0)
	goto L34
L37:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v125&v92 == v109 {
		v133 = v109
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v129 = v109 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v129)) < base.Ui64(v63) {
		v109 = v129
		goto L35
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	v150 = v62 + v138*int32(12)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	if v151 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L33
L42:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	v162 = v155
	goto L45
L43:
	;
	goto L44
L44:
	;
	v196 = v138 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v196)) < base.Ui64(v63) {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v170 = v162 & v154
	v175 = v69 + v170*int32(12)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v176 != 0 {
		v162 = v170 + int32(1)
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+8)) = v177
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v150)))
	*(*int64)(unsafe.Add(mBase, uint32(v175))) = v179
	goto L44
L47:
	;
	goto L46
L48:
	;
	v200 = v196
	goto L50
L49:
	;
	v200 = int32(0)
	goto L50
L50:
	;
	v202 = v139 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v202)) < base.Ui64(v63) {
		v138 = v200
		v139 = v202
		goto L40
	} else {
		goto L51
	}
L51:
	;
	goto L41
L52:
	;
	goto L12
L53:
	;
	v249 = v241
	v251 = v235
	v252 = v237
	v253 = v238
	goto L54
L54:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v249)+8))
	if v259 == l1 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v438 = v419
	goto L6
L56:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+52))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261)+36))
	v266 = F_ExecStoreMinimalTuple(m, v263, v264, int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L22
	} else {
		goto L59
	}
L57:
	;
	v300 = v252
	v303 = v259
	goto L58
L58:
	;
	v304 = v303 & v300
	if base.Ui32(v253) < base.Ui32(v304) {
		goto L67
	} else {
		goto L68
	}
L59:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v261)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v262)+12)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v262)+8)) = v268
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v261)+48))
	if v271 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v262)+20))
	F_MemoryContextReset(m, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L22
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v278 = int32(_a_F_tuplehash_insert_hash_internal_0)
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_tuplehash_insert_hash_internal[0]))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v262)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplehash_insert_hash_internal[0])) = v281
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v271)+24))
	v286 = m.T0[v285].(func(*base.Module, int32, int32, int32) int64)(m, v271, v262, v17+int32(15))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L22
	} else {
		goto L64
	}
L63:
	;
	v462 = v249
	v472 = int32(1)
	goto L5
L64:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tuplehash_insert_hash_internal[0])) = v279
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v262)+20))
	F_MemoryContextReset(m, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	if v286 != int64(0) {
		v462 = v249
		v472 = int32(1)
		goto L5
	} else {
		goto L66
	}
L66:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v249)+8))
	v300 = v296
	v303 = v297
	goto L58
L67:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v308 = v253 + v306
	goto L69
L68:
	;
	v308 = v253
	goto L69
L69:
	;
	v311 = (v253 + int32(1)) & v300
	if base.Ui32(v308-v304) < base.Ui32(v251) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v317 = v236 + v311*int32(12)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
	if v318 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v407 = v251 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v407) {
		goto L89
	} else {
		goto L90
	}
L73:
	;
	v322 = v311
	v324 = int32(0)
	goto L76
L74:
	;
	v354 = v311
	v357 = v317
	goto L75
L75:
	;
	if v354 != v253 {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	v334 = v324 + int32(1)
	if int32(151) <= v334 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v354 = v346
	v357 = v349
	goto L75
L78:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v339 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v337), base.F64_convert_i64_u(v339)), float64(0.1)) != 0 {
		goto L4
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v346 = (v322 + int32(1)) & v300
	v349 = v236 + v346*int32(12)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)+4))
	if v350 != 0 {
		v322 = v346
		v324 = v334
		goto L76
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	goto L77
L83:
	;
	v369 = v354
	v372 = v357
	goto L86
L84:
	;
	goto L85
L85:
	;
	v438 = v249
	goto L6
L86:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v383 = v380 & (v369 - int32(1))
	v386 = v236 + v383*int32(12)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v372)+8)) = v387
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v386)))
	*(*int64)(unsafe.Add(mBase, uint32(v372))) = v389
	if v383 != v253 {
		v369 = v383
		v372 = v386
		goto L86
	} else {
		goto L88
	}
L87:
	;
	goto L85
L88:
	;
	goto L87
L89:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v412 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v410), base.F64_convert_i64_u(v412)), float64(0.1)) != 0 {
		goto L4
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v419 = v236 + v311*int32(12)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
	if v420 != 0 {
		v249 = v419
		v251 = v407
		v252 = v300
		v253 = v311
		goto L54
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	goto L55
L94:
	;
	F_errmsg_internal(m, int32(_a_F_tuplehash_insert_hash_internal_1), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L22
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_tuplehash_insert_hash_internal_2), int32(635), int32(_a_F_tuplehash_insert_hash_internal_3))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L22
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errmsg_internal(m, int32(_a_F_tuplehash_insert_hash_internal_4), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L22
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_tuplehash_insert_hash_internal_2), int32(332), int32(_a_F_tuplehash_insert_hash_internal_5))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L22
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_typenameType(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
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
	var v46 int32
	_ = v46
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v12 = F_LookupTypeNameExtended(m, l0, l1, l2, int32(1), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v17)+82)))
			if v19 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = F_TypeNameToString(m, l1)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v54
							F_errmsg(m, int32(_a_F_typenameType_0), v8+int32(16))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
								F_parser_errposition(m, l0, v62)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_typenameType_1), int32(280), int32(_a_F_typenameType_2))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
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
			} else {
				m.G0 = v8 + int32(32)
				return v12
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = F_TypeNameToString(m, l1)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v33
						F_errmsg(m, int32(_a_F_typenameType_3), v8)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
							F_parser_errposition(m, l0, v39)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_typenameType_1), int32(274), int32(_a_F_typenameType_2))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
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
}
