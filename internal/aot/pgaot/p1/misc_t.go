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
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_pg_mblen_cstr(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_pg_mblen_with_len(m, l0, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v12 != int32(1) {
				v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_t_isalnum_cstr[0])))
				if v17&int32(1) == int32(0) {
					F_char2wchar(m, v6+int32(4), int32(3), l0, v12, int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
						if base.Ui32(int32(10)) <= base.Ui32(v40-int32(48)) {
							v45 = F_iswalpha(m, v40)
							mBase = m.M
							v49 = base.B2i32(v45 != int32(0))
						} else {
							v49 = int32(1)
						}
						v50 = v49
						m.G0 = v6 + int32(16)
						return v50
					}
				} else {
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					v50 = base.B2i32(base.Ui32(v22-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v22|int32(32)-int32(97)) < base.Ui32(int32(26)))
					m.G0 = v6 + int32(16)
					return v50
				}
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				v50 = base.B2i32(base.Ui32(v22-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v22|int32(32)-int32(97)) < base.Ui32(int32(26)))
				m.G0 = v6 + int32(16)
				return v50
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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_pg_mblen_with_len(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != int32(1) {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_t_isalnum_with_len[0])))
			if v15&int32(1) == int32(0) {
				F_char2wchar(m, v6+int32(4), int32(3), l0, v8, int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
					if base.Ui32(int32(10)) <= base.Ui32(v38-int32(48)) {
						v43 = F_iswalpha(m, v38)
						mBase = m.M
						v47 = base.B2i32(v43 != int32(0))
					} else {
						v47 = int32(1)
					}
					v48 = v47
					m.G0 = v6 + int32(16)
					return v48
				}
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				v48 = base.B2i32(base.Ui32(v20-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v20|int32(32)-int32(97)) < base.Ui32(int32(26)))
				m.G0 = v6 + int32(16)
				return v48
			}
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			v48 = base.B2i32(base.Ui32(v20-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v20|int32(32)-int32(97)) < base.Ui32(int32(26)))
			m.G0 = v6 + int32(16)
			return v48
		}
	}
}
func F_tconvert(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_from_text(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
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
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+(v32^int32(-1))<<(uint(int32(6))%32))+16))
			v51 = v42
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, _c_F_terminate_brin_buildstate[3]))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+v32<<(uint(int32(6))%32)+int32(-64))+16))
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v23<<(uint(int32(2))%32))))
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
	F_errfinish(m, int32(_a_F_textarray_to_strvaluelist_1), int32(2098), int32(_a_F_textarray_to_strvaluelist_2))
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
func F_textcat(m *base.Module, l0 int32) int32 {
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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_text_catenate(m, v3, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
func F_texteq(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = F_pg_newlocale_from_collation(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
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
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L82
	}
L4:
	;
	return int32(0)
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v15 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return v211
L7:
	;
	F_pfree(m, v203)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L81
	}
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_toast_raw_datum_size(m, v14)
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
	v115 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L45
	}
L11:
	;
	v22 = F_toast_raw_datum_size(m, v18)
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
		v211 = int32(0)
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v25 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v27 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v29 = int32(1)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v31&v29 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v34 = v29
	goto L18
L17:
	;
	v34 = int32(4)
	goto L18
L18:
	;
	v35 = v25 + v34
	v36 = int32(1)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v38&v36 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v41 = v36
	goto L21
L20:
	;
	v41 = int32(4)
	goto L21
L21:
	;
	v42 = v27 + v41
	v43 = int32(4)
	v44 = v20 - v43
	if base.Ui32(v43) <= base.Ui32(v44) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v107 != v25 {
		goto L40
	} else {
		goto L41
	}
L23:
	;
	v106 = int32(0)
	goto L22
L24:
	;
	v80 = v75
	v81 = v76
	v82 = v77
	goto L34
L25:
	;
	if (v35|v42)&int32(3) != 0 {
		v75 = v35
		v76 = v42
		v77 = v44
		goto L24
	} else {
		goto L28
	}
L26:
	;
	v68 = v35
	v69 = v42
	v70 = v44
	goto L27
L27:
	;
	if v70 == int32(0) {
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v52 = v35
	v53 = v42
	v54 = v44
	goto L29
L29:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v57 != v58 {
		v75 = v52
		v76 = v53
		v77 = v54
		goto L24
	} else {
		goto L31
	}
L30:
	;
	v68 = v63
	v69 = v61
	v70 = v65
	goto L27
L31:
	;
	v60 = int32(4)
	v61 = v53 + v60
	v63 = v52 + v60
	v65 = v54 - v60
	if base.Ui32(int32(3)) < base.Ui32(v65) {
		v52 = v63
		v53 = v61
		v54 = v65
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v75 = v68
	v76 = v69
	v77 = v70
	goto L24
L34:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v85 == v86 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v106 = v85 - v86
	goto L22
L36:
	;
	v88 = int32(1)
	v93 = v82 - v88
	if v93 != 0 {
		v80 = v80 + v88
		v81 = v81 + v88
		v82 = v93
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
	F_pfree(m, v25)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v112 = base.B2i32(v106 == int32(0))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v27 != v113 {
		v202 = v112
		v203 = v27
		goto L7
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v211 = v112
	goto L6
L45:
	;
	v118 = v115 + int32(1)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v120 = F_pg_detoast_datum_packed(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v127 = v125 & int32(1)
	if v127 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v128 = v118
	goto L49
L48:
	;
	v128 = v115 + int32(4)
	goto L49
L49:
	;
	if v125 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v156 = int32(1)
	v157 = v120 + v156
	if v122&v156 != 0 {
		goto L61
	} else {
		goto L62
	}
L51:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v134 == int32(18) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v145 = int32(1)
	if v127 != 0 {
		v155 = int32(base.Ui32(v125)>>(uint(v145)%32)) - v145
		goto L50
	} else {
		goto L60
	}
L54:
	;
	v137 = int32(16)
	goto L56
L55:
	;
	v137 = int32(0)
	goto L56
L56:
	;
	if base.Ui32((v134-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v144 = int32(4)
	goto L59
L58:
	;
	v144 = v137
	goto L59
L59:
	;
	v155 = v144
	goto L50
L60:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v155 = int32(base.Ui32(v149)>>(uint(int32(2))%32)) - int32(4)
	goto L50
L61:
	;
	v162 = v157
	goto L63
L62:
	;
	v162 = v120 + int32(4)
	goto L63
L63:
	;
	if v122 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v192 = F_varstr_cmp(m, v128, v155, v162, v191, v9)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L75
	}
L65:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v168 == int32(18) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v179 = int32(1)
	if v122&v179 != 0 {
		v191 = int32(base.Ui32(v122)>>(uint(v179)%32)) - v179
		goto L64
	} else {
		goto L74
	}
L68:
	;
	v171 = int32(16)
	goto L70
L69:
	;
	v171 = int32(0)
	goto L70
L70:
	;
	if base.Ui32((v168-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v178 = int32(4)
	goto L73
L72:
	;
	v178 = v171
	goto L73
L73:
	;
	v191 = v178
	goto L64
L74:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v191 = int32(base.Ui32(v185)>>(uint(int32(2))%32)) - int32(4)
	goto L64
L75:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v194 != v115 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_pfree(m, v115)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v199 = base.B2i32(v192 == int32(0))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v120 == v200 {
		v211 = v199
		goto L6
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v202 = v199
	v203 = v120
	goto L7
L81:
	;
	v211 = v202
	goto L6
L82:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_texteq_0), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errhint(m, int32(_a_F_texteq_1), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_texteq_2), int32(1648), int32(_a_F_texteq_3))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
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
func F_texteqfast(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_DirectFunctionCall2Coll(m, int32(1544), int32(100), l0, l1)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v5 != int32(0))
	}
}
func F_textgtname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(1543), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(int32(0) < v6)
	}
}
func F_texthashfast(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_DirectFunctionCall1Coll(m, int32(1570), int32(100), l0)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_texticregexne(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14018(m, l0, int32(27))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_textoverlay_no_len(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_textoverlay_no_len[0]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v16*int32(28))+uint32(_c_F_textoverlay_no_len[1])))
			if v21 == int32(1) {
				v24 = F_toast_raw_datum_size(m, v11)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v28 = F_text_overlay(m, v6, v11, v13, v24-int32(4))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						return v28
					}
				}
			} else {
				v31 = F_pg_detoast_datum_packed(m, v11)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = int32(1)
					v34 = v31 + v33
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
					v39 = v37 & v33
					if v39 != 0 {
						v40 = v34
					} else {
						v40 = v31 + int32(4)
					}
					if v37 == int32(1) {
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
						if v46 == int32(18) {
							v49 = int32(16)
						} else {
							v49 = int32(0)
						}
						if base.Ui32((v46-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v56 = int32(4)
						} else {
							v56 = v49
						}
						v67 = v56
					} else {
						v57 = int32(1)
						if v39 != 0 {
							v67 = int32(base.Ui32(v37)>>(uint(v57)%32)) - v57
						} else {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
							v67 = int32(base.Ui32(v61)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v68 = F_pg_mbstrlen_with_len(m, v40, v67)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						v70 = F_text_overlay(m, v6, v11, v13, v68)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							return v70
						}
					}
				}
			}
		}
	}
}
func F_textpos(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	v7 = m.G0
	v9 = v7 - int32(1072)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v19 != 0 {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v20 == int32(1) {
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
					if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
						if v50 == int32(1) {
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
							if v56 == int32(18) {
								v59 = int32(16)
							} else {
								v59 = int32(0)
							}
							if base.Ui32((v56-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v66 = int32(4)
							} else {
								v66 = v59
							}
							v79 = v66
						} else {
							v67 = int32(1)
							if v50&v67 != 0 {
								v79 = int32(base.Ui32(v50)>>(uint(v67)%32)) - v67
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if v20 == int32(1) {
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
							if v85 == int32(18) {
								v88 = int32(16)
							} else {
								v88 = int32(0)
							}
							if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v95 = int32(4)
							} else {
								v95 = v88
							}
							v108 = v95
						} else {
							v96 = int32(1)
							if v20&v96 != 0 {
								v108 = int32(base.Ui32(v20)>>(uint(v96)%32)) - v96
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
								v108 = int32(base.Ui32(v102)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if base.Ui32(v108) <= base.Ui32(v79) {
							F_text_position_setup(m, v12, v17, v19, v9)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								v118 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v118)
								v121 = F_text_position_next(m, v9)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int32(0)
								} else {
									if v121 == int32(0) {
										v134 = v118
										m.G0 = v9 + int32(1072)
										return v134
									} else {
										v125 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
										v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1052))
										v128 = F_pg_mbstrlen_with_len(m, v125, v126-v125)
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
											v134 = v128 + v130 + int32(1)
											m.G0 = v9 + int32(1072)
											return v134
										}
									}
								}
							}
						} else {
							v110 = F_pg_newlocale_from_collation(m, v19)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
								if v112 == int32(0) {
									F_text_position_setup(m, v12, v17, v19, v9)
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int32(0)
									} else {
										v118 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v118)
										v121 = F_text_position_next(m, v9)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int32(0)
										} else {
											if v121 == int32(0) {
												v134 = v118
												m.G0 = v9 + int32(1072)
												return v134
											} else {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1052))
												v128 = F_pg_mbstrlen_with_len(m, v125, v126-v125)
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
													v134 = v128 + v130 + int32(1)
													m.G0 = v9 + int32(1072)
													return v134
												}
											}
										}
									}
								} else {
									v134 = int32(0)
									m.G0 = v9 + int32(1072)
									return v134
								}
							}
						}
					} else {
						if v23 == int32(18) {
							v34 = int32(16)
						} else {
							v34 = int32(0)
						}
						v47 = v34
						if v47 != 0 {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
							if v50 == int32(1) {
								v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
								if v56 == int32(18) {
									v59 = int32(16)
								} else {
									v59 = int32(0)
								}
								if base.Ui32((v56-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v66 = int32(4)
								} else {
									v66 = v59
								}
								v79 = v66
							} else {
								v67 = int32(1)
								if v50&v67 != 0 {
									v79 = int32(base.Ui32(v50)>>(uint(v67)%32)) - v67
								} else {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							if v20 == int32(1) {
								v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
								if v85 == int32(18) {
									v88 = int32(16)
								} else {
									v88 = int32(0)
								}
								if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v95 = int32(4)
								} else {
									v95 = v88
								}
								v108 = v95
							} else {
								v96 = int32(1)
								if v20&v96 != 0 {
									v108 = int32(base.Ui32(v20)>>(uint(v96)%32)) - v96
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
									v108 = int32(base.Ui32(v102)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							if base.Ui32(v108) <= base.Ui32(v79) {
								F_text_position_setup(m, v12, v17, v19, v9)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return int32(0)
								} else {
									v118 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v118)
									v121 = F_text_position_next(m, v9)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int32(0)
									} else {
										if v121 == int32(0) {
											v134 = v118
											m.G0 = v9 + int32(1072)
											return v134
										} else {
											v125 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
											v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1052))
											v128 = F_pg_mbstrlen_with_len(m, v125, v126-v125)
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
												v134 = v128 + v130 + int32(1)
												m.G0 = v9 + int32(1072)
												return v134
											}
										}
									}
								}
							} else {
								v110 = F_pg_newlocale_from_collation(m, v19)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
									if v112 == int32(0) {
										F_text_position_setup(m, v12, v17, v19, v9)
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
											return int32(0)
										} else {
											v118 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v118)
											v121 = F_text_position_next(m, v9)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return int32(0)
											} else {
												if v121 == int32(0) {
													v134 = v118
													m.G0 = v9 + int32(1072)
													return v134
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1052))
													v128 = F_pg_mbstrlen_with_len(m, v125, v126-v125)
													mBase = m.M
													v129 = m.ExcPending
													if v129 != 0 {
														return int32(0)
													} else {
														v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
														v134 = v128 + v130 + int32(1)
														m.G0 = v9 + int32(1072)
														return v134
													}
												}
											}
										}
									} else {
										v134 = int32(0)
										m.G0 = v9 + int32(1072)
										return v134
									}
								}
							}
						} else {
							v134 = int32(1)
							m.G0 = v9 + int32(1072)
							return v134
						}
					}
				} else {
					v35 = int32(1)
					if v20&v35 != 0 {
						v47 = int32(base.Ui32(v20)>>(uint(v35)%32)) - v35
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
					}
					if v47 != 0 {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
						if v50 == int32(1) {
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
							if v56 == int32(18) {
								v59 = int32(16)
							} else {
								v59 = int32(0)
							}
							if base.Ui32((v56-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v66 = int32(4)
							} else {
								v66 = v59
							}
							v79 = v66
						} else {
							v67 = int32(1)
							if v50&v67 != 0 {
								v79 = int32(base.Ui32(v50)>>(uint(v67)%32)) - v67
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if v20 == int32(1) {
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
							if v85 == int32(18) {
								v88 = int32(16)
							} else {
								v88 = int32(0)
							}
							if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v95 = int32(4)
							} else {
								v95 = v88
							}
							v108 = v95
						} else {
							v96 = int32(1)
							if v20&v96 != 0 {
								v108 = int32(base.Ui32(v20)>>(uint(v96)%32)) - v96
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
								v108 = int32(base.Ui32(v102)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if base.Ui32(v108) <= base.Ui32(v79) {
							F_text_position_setup(m, v12, v17, v19, v9)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int32(0)
							} else {
								v118 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v118)
								v121 = F_text_position_next(m, v9)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int32(0)
								} else {
									if v121 == int32(0) {
										v134 = v118
										m.G0 = v9 + int32(1072)
										return v134
									} else {
										v125 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
										v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1052))
										v128 = F_pg_mbstrlen_with_len(m, v125, v126-v125)
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
											v134 = v128 + v130 + int32(1)
											m.G0 = v9 + int32(1072)
											return v134
										}
									}
								}
							}
						} else {
							v110 = F_pg_newlocale_from_collation(m, v19)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
								if v112 == int32(0) {
									F_text_position_setup(m, v12, v17, v19, v9)
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int32(0)
									} else {
										v118 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v118)
										v121 = F_text_position_next(m, v9)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int32(0)
										} else {
											if v121 == int32(0) {
												v134 = v118
												m.G0 = v9 + int32(1072)
												return v134
											} else {
												v125 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1052))
												v128 = F_pg_mbstrlen_with_len(m, v125, v126-v125)
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
													v134 = v128 + v130 + int32(1)
													m.G0 = v9 + int32(1072)
													return v134
												}
											}
										}
									}
								} else {
									v134 = int32(0)
									m.G0 = v9 + int32(1072)
									return v134
								}
							}
						}
					} else {
						v134 = int32(1)
						m.G0 = v9 + int32(1072)
						return v134
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v143 = m.ExcPending
				if v143 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(34209924))
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_textpos_0), int32(0))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_textpos_1), int32(0))
							mBase = m.M
							v154 = m.ExcPending
							if v154 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_textpos_2), int32(1648), int32(_a_F_textpos_3))
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
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
func F_textrecv(m *base.Module, l0 int32) int32 {
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v16 = F_pq_getmsgtext(m, v10, v11-v12, v8+int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v22 = v20 + int32(4)
		v23 = F_palloc(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
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
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return v23
			}
		}
	}
}
func F_thesaurus_init(m *base.Module, l0 int32) int32 {
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v109 int32
	_ = v109
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
	var v153 int32
	_ = v153
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
	var v318 int32
	_ = v318
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
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var __phi684 int32
	_ = __phi684
	var v685 int32
	_ = v685
	var __phi685 int32
	_ = __phi685
	var v687 int32
	_ = v687
	var __phi687 int32
	_ = __phi687
	var v689 int32
	_ = v689
	var __phi689 int32
	_ = __phi689
	var v690 int32
	_ = v690
	var __phi690 int32
	_ = __phi690
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var __phi818 int32
	_ = __phi818
	var v820 int32
	_ = v820
	var __phi820 int32
	_ = __phi820
	var v821 int32
	_ = v821
	var __phi821 int32
	_ = __phi821
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v894 int64
	_ = v894
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v951 int32
	_ = v951
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v992 int32
	_ = v992
	var __phi992 int32
	_ = __phi992
	var v994 int32
	_ = v994
	var __phi994 int32
	_ = __phi994
	var v995 int32
	_ = v995
	var __phi995 int32
	_ = __phi995
	var v999 int32
	_ = v999
	var __phi999 int32
	_ = __phi999
	var v1000 int32
	_ = v1000
	var __phi1000 int32
	_ = __phi1000
	var v1006 int32
	_ = v1006
	var v1009 int64
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var __phi1038 int32
	_ = __phi1038
	var v1041 int32
	_ = v1041
	var __phi1041 int32
	_ = __phi1041
	var v1042 int32
	_ = v1042
	var __phi1042 int32
	_ = __phi1042
	var v1048 int32
	_ = v1048
	var __phi1048 int32
	_ = __phi1048
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int64
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1181 int32
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1203 int32
	_ = v1203
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1249 int32
	_ = v1249
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1281 int32
	_ = v1281
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1302 int32
	_ = v1302
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(160)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v1321 = m.ExcPending
	if v1321 != 0 {
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
	v44 = v2
	v45 = v2
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
	if v513&int32(1) == int32(0) {
		goto L3
	} else {
		goto L145
	}
L8:
	;
	v526 = v44 + int32(1)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v526 < v527 {
		v34 = v513
		v44 = v526
		v45 = v524
		goto L6
	} else {
		goto L144
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v498
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
		v498 = v94
		goto L9
	} else {
		goto L30
	}
L30:
	;
	v101 = v96
	v108 = v94
	v109 = v94
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
		v318 = v109
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
	v153 = v136
	goto L44
L42:
	;
	if v130 == int32(0) {
		v317 = v108
		v318 = v109
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
	v278 = v153
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
		v278 = v153
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
		v278 = v153
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
	v189 = v153
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
	F_newLexeme(m, v22, v153, v143, v109, v145&int32(_a_F_thesaurus_init_6))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	F_newLexeme(m, v22, v153, v143, v109, v145&int32(_a_F_thesaurus_init_6))
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
	v278 = v153
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
	v278 = v153
	goto L46
L78:
	;
	goto L79
L79:
	;
	if v143 == v153 {
		goto L48
	} else {
		goto L80
	}
L80:
	;
	v238 = int32(_a_F_thesaurus_init_6)
	F_addWrd(m, v22, v153, v143, v109, v149&v238, v145&v238, v151&int32(1))
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
		v153 = v278
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
	F_addWrd(m, v22, v278, v281, v109, v276&v286, v273&v286, v277&int32(1))
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
	v318 = v109 + int32(1)
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
		v109 = v318
		goto L31
	} else {
		goto L98
	}
L98:
	;
	v498 = v318
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
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(616), int32(_a_F_thesaurus_init_13))
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
	if v45 == int32(0) {
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
	v513 = v34
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
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(625), int32(_a_F_thesaurus_init_13))
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
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(633), int32(_a_F_thesaurus_init_13))
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
	v513 = int32(1)
	v524 = v45
	goto L8
L144:
	;
	goto L7
L145:
	;
	if v524 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v533 = int32(0)
	v535 = F_stringToQualifiedNameList(m, v524, v533)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L1
	} else {
		goto L313
	}
L149:
	;
	v538 = F_get_ts_dict_oid(m, v535, int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v538
	v541 = F_lookup_ts_dictionary_cache(m, v538)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v541
	v545 = F_palloc(m, int32(128))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v547 <= int32(0) {
		goto L156
	} else {
		goto L157
	}
L153:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L1
	} else {
		goto L309
	}
L154:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L1
	} else {
		goto L304
	}
L155:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v800 != 0 {
		goto L202
	} else {
		goto L203
	}
L156:
	;
	v788 = v533
	v791 = int32(16)
	v793 = v545
	goto L155
L157:
	;
	goto L158
L158:
	;
	v555 = v533
	v558 = int32(16)
	v560 = v545
	v564 = v2
	goto L159
L159:
	;
	v568 = v564 << (uint(int32(3)) % 32)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v570 = v568 + v569
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v570)))
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571))))
	if v572 != int32(63) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	v788 = v759
	v791 = v762
	v793 = v764
	goto L155
L161:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v771+v568)))
	F_pfree(m, v773)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L199
	}
L162:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v611 = int32(0)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v608)+44))
	v613 = F_strlen(m, v571)
	mBase = m.M
	v615 = F_FunctionCall4Coll(m, v608+int32(12), v611, v612, v571, v613, v611)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L170
	}
L163:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571)+1)))
	if v575 != 0 {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v570)+4))
	if v558 <= v555 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v580 = F_repalloc(m, v560, v558<<(uint(int32(4))%32))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	v584 = v558
	v585 = v560
	goto L167
L167:
	;
	v587 = F_palloc(m, int32(16))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L169
	}
L168:
	;
	v584 = v558 << (uint(int32(1)) % 32)
	v585 = v580
	goto L167
L169:
	;
	v591 = v585 + v555<<(uint(int32(3))%32)
	v592 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v591))) = v592
	*(*int32)(unsafe.Add(mBase, uint32(v591)+4)) = v587
	v595 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v587)+6)) = uint16(v595)
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v591)+4))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v576)))
	*(*int32)(unsafe.Add(mBase, uint32(v597))) = v598
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v591)+4))
	v601 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v576)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v600)+4)) = uint16(v601)
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v591)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v603)+8)) = v592
	v759 = v555 + v595
	v762 = v584
	v764 = v585
	goto L161
L170:
	;
	if v615 == int32(0) {
		goto L153
	} else {
		goto L171
	}
L171:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v615)+4))
	if v619 == int32(0) {
		goto L154
	} else {
		goto L172
	}
L172:
	;
	v624 = v615
	v625 = v555
	v628 = v558
	v630 = v560
	goto L173
L173:
	;
	v637 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v624))))
	v638 = int32(1)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v624)+12))
	if v639 == int32(0) {
		v677 = v638
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v759 = v748
	v762 = v750
	v764 = v751
	goto L161
L175:
	;
	__phi684 = v625
	__phi685 = v624
	__phi687 = v628
	__phi689 = v630
	__phi690 = v624 + int32(4)
	v684 = __phi684
	v685 = __phi685
	v687 = __phi687
	v689 = __phi689
	v690 = __phi690
	goto L181
L176:
	;
	v642 = v624
	v655 = v638
	goto L177
L177:
	;
	v657 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v642)+8)))
	if v657 != v637 {
		v677 = v655
		goto L175
	} else {
		goto L179
	}
L178:
	;
	v677 = v660
	goto L175
L179:
	;
	v660 = v655 + int32(1)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v642)+20))
	if v661 != 0 {
		v642 = v642 + int32(8)
		v655 = v660
		goto L177
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v696 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v685))))
	if v637 != v696 {
		goto L184
	} else {
		goto L185
	}
L182:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v747)+4))
	if v755 != 0 {
		v624 = v747
		v625 = v748
		v628 = v750
		v630 = v751
		goto L173
	} else {
		goto L198
	}
L183:
	;
	goto L182
L184:
	;
	v747 = v685
	v748 = v684
	v750 = v687
	v751 = v689
	goto L183
L185:
	;
	goto L186
L186:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v698+v568)+4))
	if v687 <= v684 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v704 = F_repalloc(m, v689, v687<<(uint(int32(4))%32))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L190
	}
L188:
	;
	v708 = v687
	v709 = v689
	goto L189
L189:
	;
	v712 = v709 + v684<<(uint(int32(3))%32)
	v714 = F_palloc(m, int32(16))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L191
	}
L190:
	;
	v708 = v687 << (uint(int32(1)) % 32)
	v709 = v704
	goto L189
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v712)+4)) = v714
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v690)))
	if v717 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v712))) = v725
	*(*uint16)(unsafe.Add(mBase, uint32(v726)+6)) = uint16(v727)
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v700)))
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = v731
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	v734 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v700)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v733)+4)) = uint16(v734)
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v736)+8)) = int32(0)
	v742 = v684 + int32(1)
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v685)+12))
	v745 = v685 + int32(8)
	if v743 != 0 {
		__phi684 = v742
		__phi685 = v745
		__phi687 = v708
		__phi689 = v709
		__phi690 = v685 + int32(12)
		v684 = __phi684
		v685 = __phi685
		v687 = __phi687
		v689 = __phi689
		v690 = __phi690
		goto L181
	} else {
		goto L197
	}
L193:
	;
	v725 = int32(0)
	v726 = v714
	v727 = int32(1)
	goto L192
L194:
	;
	goto L195
L195:
	;
	v722 = F_pstrdup(m, v717)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	v725 = v722
	v726 = v724
	v727 = v677
	goto L192
L197:
	;
	v747 = v745
	v748 = v742
	v750 = v708
	v751 = v709
	goto L183
L198:
	;
	goto L174
L199:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v776+v568)+4))
	F_pfree(m, v778)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v782 = v564 + int32(1)
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v782 < v783 {
		v555 = v759
		v558 = v762
		v560 = v764
		v564 = v782
		goto L159
	} else {
		goto L201
	}
L201:
	;
	goto L160
L202:
	;
	F_pfree(m, v800)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v793
	if int32(2) <= v788 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	goto L204
L206:
	;
	F_pg_qsort(m, v793, v788, int32(8), int32(1149))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L1
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if int32(0) < v951 {
		goto L248
	} else {
		goto L249
	}
L209:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v813 < int32(2) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v926 = int32(3)
	v929 = (v910-v914)>>(uint(v926)%32) + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v929
	v933 = F_repalloc(m, v914, v929<<(uint(v926)%32))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L1
	} else {
		goto L244
	}
L211:
	;
	v910 = v812
	v914 = v812
	goto L210
L212:
	;
	goto L213
L213:
	;
	__phi818 = v812
	__phi820 = v812 + int32(8)
	__phi821 = v812
	v818 = __phi818
	v820 = __phi820
	v821 = __phi821
	goto L214
L214:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v818)))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v821)+8))
	if v834 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L215:
	;
	v910 = v898
	v914 = v905
	goto L210
L216:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v904 = v820 + int32(8)
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if (v904-v905)>>(uint(int32(3))%32) < v902 {
		__phi818 = v898
		__phi820 = v904
		__phi821 = v820
		v818 = __phi818
		v820 = __phi820
		v821 = __phi821
		goto L214
	} else {
		goto L243
	}
L217:
	;
	v894 = *(*int64)(unsafe.Add(mBase, uint32(v821)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v818)+8)) = v894
	v898 = v818 + int32(8)
	goto L216
L218:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v821)+12))
	if v867 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L219:
	;
	if v833 == int32(0) {
		goto L218
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	if v833 == int32(0) {
		goto L217
	} else {
		goto L223
	}
L222:
	;
	goto L217
L223:
	;
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834))))
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v833))))
	if base.B2i32(v843 == int32(0))|base.B2i32(v843 != v846) != 0 {
		v864 = v843
		v865 = v846
		goto L225
	} else {
		goto L226
	}
L224:
	;
	if v864-v865 != 0 {
		goto L217
	} else {
		goto L231
	}
L225:
	;
	goto L224
L226:
	;
	v849 = v834
	v850 = v833
	goto L227
L227:
	;
	v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850)+1)))
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849)+1)))
	if v854 == int32(0) {
		v864 = v854
		v865 = v853
		goto L225
	} else {
		goto L229
	}
L228:
	;
	v864 = v854
	v865 = v853
	goto L225
L229:
	;
	v857 = int32(1)
	if v854 == v853 {
		v849 = v849 + v857
		v850 = v850 + v857
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
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v821)+8))
	if v889 == int32(0) {
		v898 = v818
		goto L216
	} else {
		goto L241
	}
L233:
	;
	F_pfree(m, v867)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L1
	} else {
		goto L240
	}
L234:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v818)+4))
	if v870 == int32(0) {
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v867)))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v870)))
	if v873 != v874 {
		goto L236
	} else {
		goto L237
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v867)+8)) = v870
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v821)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v818)+4)) = v883
	goto L232
L237:
	;
	v876 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v867)+4)))
	v877 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v870)+4)))
	if v876 != v877 {
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v879 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v867)+6)))
	v880 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v870)+6)))
	if v879 == v880 {
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
	F_pfree(m, v889)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	v898 = v818
	goto L216
L243:
	;
	goto L215
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v933
	goto L208
L245:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L1
	} else {
		goto L300
	}
L246:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L1
	} else {
		goto L296
	}
L247:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L292
	}
L248:
	;
	v967 = int32(0)
	goto L251
L249:
	;
	goto L250
L250:
	;
	m.G0 = v18 + int32(160)
	return v22
L251:
	;
	v971 = v967 << (uint(int32(3)) % 32)
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v971+v972)+4))
	v976 = F_palloc(m, int32(16))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L253
	}
L252:
	;
	goto L250
L253:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v978+v971)+4)) = v976
	v981 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v976)+4)) = v981
	if v974 == v981 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v1133 = v1132 + v971
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+4))
	if v1117 == v1134 {
		goto L245
	} else {
		goto L289
	}
L255:
	;
	v1117 = v976
	goto L254
L256:
	;
	goto L257
L257:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v974)+4))
	if v985 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1117 = v976
	goto L254
L259:
	;
	goto L260
L260:
	;
	__phi992 = v976
	__phi994 = int32(2)
	__phi995 = v985
	__phi999 = v974
	__phi1000 = v974 + int32(4)
	v992 = __phi992
	v994 = __phi994
	v995 = __phi995
	v999 = __phi999
	v1000 = __phi1000
	goto L261
L261:
	;
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v999)+3)))
	if v1006&int32(16) != 0 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	v1117 = v1086
	goto L254
L263:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+4))
	if v1030 == int32(0) {
		goto L247
	} else {
		goto L269
	}
L264:
	;
	v1009 = *(*int64)(unsafe.Add(mBase, uint32(v999)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v1009
	v1011 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v1011
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+114)) = uint16(v1011)
	v1029 = v18 + int32(112)
	goto L263
L265:
	;
	goto L266
L266:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v1020 = int32(0)
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1017)+44))
	v1022 = F_strlen(m, v995)
	mBase = m.M
	v1024 = F_FunctionCall4Coll(m, v1017+int32(12), v1020, v1021, v995, v1022, v1020)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	if v1024 == int32(0) {
		goto L246
	} else {
		goto L268
	}
L268:
	;
	v1029 = v1024
	goto L263
L269:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1035+v971)+4))
	__phi1038 = v992
	__phi1041 = v994
	__phi1042 = v1029
	__phi1048 = v1029 + int32(4)
	v1038 = __phi1038
	v1041 = __phi1041
	v1042 = __phi1042
	v1048 = __phi1048
	goto L270
L270:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1053+v971)+4))
	v1056 = v1038 - v1055
	if v1041 <= v1056>>(uint(int32(3))%32)+int32(1) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	if v992 == v1037 {
		goto L278
	} else {
		goto L279
	}
L272:
	;
	v1064 = F_repalloc(m, v1055, v1041<<(uint(int32(4))%32))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L1
	} else {
		goto L275
	}
L273:
	;
	v1075 = v1038
	v1076 = v1041
	goto L274
L274:
	;
	v1077 = *(*int64)(unsafe.Add(mBase, uint32(v1042)))
	*(*int64)(unsafe.Add(mBase, uint32(v1075))) = v1077
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1048)))
	v1080 = F_pstrdup(m, v1079)
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L1
	} else {
		goto L276
	}
L275:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1066+v971)+4)) = v1064
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1071+v971)+4))
	v1075 = v1073 + v1056
	v1076 = v1041 << (uint(int32(1)) % 32)
	goto L274
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1075)+4)) = v1080
	v1085 = int32(8)
	v1086 = v1075 + v1085
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1042)+12))
	if v1087 != 0 {
		__phi1038 = v1086
		__phi1041 = v1076
		__phi1042 = v1042 + v1085
		__phi1048 = v1042 + int32(12)
		v1038 = __phi1038
		v1041 = __phi1041
		v1042 = __phi1042
		v1048 = __phi1048
		goto L270
	} else {
		goto L277
	}
L277:
	;
	goto L271
L278:
	;
	v1095 = int32(-1)
	goto L280
L279:
	;
	v1095 = (v992 - v1037) >> (uint(int32(3)) % 32)
	goto L280
L280:
	;
	if int32(0) < v1095 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1098+v971)+4))
	v1103 = v1100 + v1095<<(uint(int32(3))%32)
	v1104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1103)+2)))
	v1106 = v1104 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1103)+2)) = uint16(v1106)
	goto L283
L282:
	;
	goto L283
L283:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1000)))
	if v1109 != 0 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	F_pfree(m, v1109)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v999)+12))
	if v1114 != 0 {
		__phi992 = v1086
		__phi994 = v1076
		__phi995 = v1114
		__phi999 = v999 + int32(8)
		__phi1000 = v999 + int32(12)
		v992 = __phi992
		v994 = __phi994
		v995 = __phi995
		v999 = __phi999
		v1000 = __phi1000
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
	v1138 = int32(base.Ui32(v1117-v1134) >> (uint(int32(3)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v1133)+2)) = uint16(v1138)
	F_pfree(m, v974)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	v1143 = v967 + int32(1)
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v1143 < v1144 {
		v967 = v1143
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
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1000)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v967 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v1172
	F_errmsg(m, int32(_a_F_thesaurus_init_16), v18+int32(32))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(568), int32(_a_F_thesaurus_init_17))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
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
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L297
	}
L297:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1000)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v967 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1194
	F_errmsg(m, int32(_a_F_thesaurus_init_18), v18+int32(16))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(575), int32(_a_F_thesaurus_init_17))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
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
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v967 + int32(1)
	F_errmsg(m, int32(_a_F_thesaurus_init_19), v18)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L1
	} else {
		goto L302
	}
L302:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(587), int32(_a_F_thesaurus_init_17))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
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
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v1237 = v1234 + v564<<(uint(int32(3))%32)
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1237)+4))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1238)))
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1237)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v1240
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v1239 + int32(1)
	F_errmsg(m, int32(_a_F_thesaurus_init_20), v18-int32(-64))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errhint(m, int32(_a_F_thesaurus_init_21), int32(0))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(425), int32(_a_F_thesaurus_init_22))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
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
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v1269 = v1266 + v564<<(uint(int32(3))%32)
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1269)+4))
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1270)))
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1269)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v1272
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v1271 + int32(1)
	F_errmsg(m, int32(_a_F_thesaurus_init_23), v18+int32(48))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(418), int32(_a_F_thesaurus_init_22))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
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
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	F_errmsg(m, int32(_a_F_thesaurus_init_24), int32(0))
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(644), int32(_a_F_thesaurus_init_13))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
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
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	F_errmsg(m, int32(_a_F_thesaurus_init_25), int32(0))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	F_errfinish(m, int32(_a_F_thesaurus_init_4), int32(640), int32(_a_F_thesaurus_init_13))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
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
func F_tideq(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	return base.B2i32(v27 == int32(0))
}
func F_tidout(m *base.Module, l0 int32) int32 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
		return int32(0)
	} else {
		v26 = F_pstrdup(m, v19)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(48)
			return v26
		}
	}
}
func F_tidrecv(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pq_getmsgint(m, v4, int32(4))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = F_pq_getmsgint(m, v4, int32(2))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v14 = F_palloc(m, int32(6))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)) = uint16(v11)
				*(*uint16)(unsafe.Add(mBase, uint32(v14)+2)) = uint16(v6)
				v19 = int32(base.Ui32(v6) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v14))) = uint16(v19)
				return v14
			}
		}
	}
}
func F_tidsend(m *base.Module, l0 int32) int32 {
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v8)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10))))
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+2)))
		F_enlargeStringInfo(m, v8, int32(4))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
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
				return int32(0)
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
				return v56
			}
		}
	}
}
func F_tidsmaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v10 = int32(16)
	v12 = v8 | v9<<(uint(v10)%32)
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+2)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4))))
	v17 = v13 | v14<<(uint(v10)%32)
	if base.Ui32(v12) < base.Ui32(v17) {
		v28 = int32(-1)
	} else {
		if base.Ui32(v17) < base.Ui32(v12) {
			v28 = int32(1)
		} else {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
			if base.Ui32(v22) < base.Ui32(v23) {
				v28 = int32(-1)
			} else {
				v28 = base.B2i32(base.Ui32(v23) < base.Ui32(v22))
			}
		}
	}
	if v28 <= int32(0) {
		v31 = v3
	} else {
		v31 = v4
	}
	return v31
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
	F_errfinish(m, int32(_a_F_tliSwitchPoint_1), int32(590), int32(_a_F_tliSwitchPoint_2))
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
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
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
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
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
	var v254 int32
	_ = v254
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
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v440 int32
	_ = v440
	var v465 int32
	_ = v465
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v608 int32
	_ = v608
	var v618 int32
	_ = v618
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v653 int32
	_ = v653
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v696 int32
	_ = v696
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v724 int32
	_ = v724
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v752 int32
	_ = v752
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v785 int32
	_ = v785
	var v793 int32
	_ = v793
	var v813 int32
	_ = v813
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v974 int32
	_ = v974
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1064 int32
	_ = v1064
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1219 int32
	_ = v1219
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	v25 = m.G0
	v27 = v25 - int32(80)
	m.G0 = v27
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = int32(799)
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
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v1237
	F_MemoryContextDelete(m, v1235)
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L1
	} else {
		goto L211
	}
L8:
	;
	if int32(base.Ui32(v64)>>(uint(int32(4))%32))&int32(1) != 0 {
		v1219 = v27
		v1235 = v50
		v1237 = v53
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
	v87 = l4 + v69
	v91 = v69
	v93 = v50
	v95 = v53
	goto L10
L10:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	goto L12
L11:
	;
	v1219 = v849
	v1235 = v865
	v1237 = v867
	goto L7
L12:
	;
	if int32(base.Ui32(v96)>>(uint(int32(5))%32))&int32(1) != 0 {
		v1219 = v77
		v1235 = v93
		v1237 = v95
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
		v254 = v113
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v125 = v101
	v138 = v113
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
	v254 = v228
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+48)) = v211
	if v211 <= v125 {
		v254 = v138
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
		v254 = v138
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
	v228 = v138 + v221
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
		v125 = v229
		v138 = v228
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
		v844 = v72
		v845 = v73
		v846 = v74
		v847 = v75
		v849 = v77
		v851 = base.B2i32(v287 == v288)
		v857 = v288
		v859 = v87
		v863 = v91
		v864 = v254
		v865 = v93
		v867 = v95
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
	F_errfinish(m, int32(_a_F_tokenize_auth_file_5), int32(769), int32(_a_F_tokenize_auth_file_1))
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
	v1219 = v77
	v1235 = v93
	v1237 = v95
	goto L7
L45:
	;
	if v851 != 0 {
		goto L137
	} else {
		goto L138
	}
L46:
	;
	v296 = v72
	v297 = v73
	v298 = v74
	v299 = v75
	v301 = v77
	v307 = v291
	v309 = v288
	v311 = v87
	v315 = v91
	v316 = v254
	v317 = v93
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
	v844 = v296
	v845 = v297
	v846 = v298
	v847 = v299
	v849 = v301
	v851 = v838
	v857 = v835
	v859 = v311
	v863 = v315
	v864 = v316
	v865 = v317
	v867 = v319
	goto L45
L49:
	;
	v336 = v307
	v341 = int32(0)
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
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v301)+60))
	F_pfree(m, v821)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L129
	}
L52:
	;
	v366 = v336
	v368 = int32(0)
	goto L54
L53:
	;
	v398 = int32(0)
	v407 = v383
	v410 = v386
	v411 = v398
	v412 = v398
	v415 = v397
	v417 = v398
	v420 = v398
	goto L59
L54:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	v384 = int32(1)
	v386 = v366 + v384
	switch v383 {
	case 0:
		v397 = v383
		goto L53
	default:
		goto L57
	case 9, 13, 32:
		v388 = v384
		goto L56
	}
L55:
	;
	v397 = int32(0)
	goto L53
L56:
	;
	if base.B2i32(v383 == int32(44))|v388 != 0 {
		v366 = v386
		v368 = v368 + int32(1)
		goto L54
	} else {
		goto L58
	}
L57:
	;
	v388 = int32(0)
	goto L56
L58:
	;
	goto L55
L59:
	;
	v428 = v407 & int32(255)
	switch v428 {
	case 0:
		v492 = v410
		v493 = v398
		goto L61
	default:
		goto L65
	case 9, 13, 32:
		goto L66
	}
L60:
	;
	v495 = v492 - int32(1)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v301)+64))
	if v420|base.B2i32(int32(0) < v496) != 0 {
		goto L80
	} else {
		goto L81
	}
L61:
	;
	goto L60
L62:
	;
	v487 = int32(1)
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	v407 = v489
	v410 = v410 + v487
	v411 = v483
	v412 = v484
	v415 = v485
	v417 = v417 + v487
	v420 = v486
	goto L59
L63:
	;
	v477 = int32(1)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v301)+64))
	if v479 != 0 {
		goto L77
	} else {
		goto L78
	}
L64:
	;
	F_appendStringInfoChar(m, v301+int32(60), base.I32_extend8_s(v407))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L75
	}
L65:
	;
	if (base.B2i32(v428 != int32(35))|v412)&int32(1) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	if v412&int32(1) == int32(0) {
		v492 = v410
		v493 = v398
		goto L61
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	v440 = F_strlen(m, v410)
	mBase = m.M
	v492 = v440 + (v336 + v417 + v368) + int32(2)
	v493 = v398
	goto L61
L69:
	;
	goto L70
L70:
	;
	if (base.B2i32(v428 != int32(44))|v412)&int32(1) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v492 = v410
	v493 = int32(1)
	goto L61
L72:
	;
	goto L73
L73:
	;
	if (base.B2i32(v428 != int32(34))|v411)&int32(1) == int32(0) {
		v476 = v412
		goto L63
	} else {
		goto L74
	}
L74:
	;
	goto L64
L75:
	;
	if v428 != int32(34) {
		v483 = int32(0)
		v484 = v412
		v485 = v415
		v486 = v420
		goto L62
	} else {
		goto L76
	}
L76:
	;
	v476 = v412 & base.B2i32(v428 == int32(34)) & (v411 ^ int32(1))
	goto L63
L77:
	;
	v480 = v415
	goto L79
L78:
	;
	v480 = v477
	goto L79
L79:
	;
	v483 = v476
	v484 = v412 ^ int32(1)
	v485 = v480
	v486 = v477
	goto L62
L80:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v301)+60))
	if v415&int32(1)|base.B2i32(v496 < int32(2)) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L81:
	;
	v813 = v341
	goto L82
L82:
	;
	goto L51
L83:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	if v493&base.B2i32(v793 == int32(0)) != 0 {
		v336 = v495
		v341 = v785
		goto L50
	} else {
		goto L128
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+76)) = int32(0)
	v539 = F_AbsoluteConfigLocation(m, v500+int32(1), v296)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L94
	}
L85:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v500))))
	if v508 == int32(64) {
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v511 = int32(_a_F_tokenize_auth_file_3)
	v512 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1]))
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v515
	v517 = F_strlen(m, v500)
	mBase = m.M
	v520 = F_palloc0(m, v517+int32(13))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v520)+8)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v520)+4)) = uint8(v415)
	v526 = v520 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v520))) = v526
	v529 = v517 + int32(1)
	if v529 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	base.MemoryCopy(m, v526, v500, v529)
	goto L92
L91:
	;
	goto L92
L92:
	;
	v531 = F_lappend(m, v341, v520)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v512
	v785 = v531
	goto L83
L94:
	;
	v543 = F_open_auth_file(m, v539, v299, v311, v301+int32(20))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	if v543 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_pfree(m, v539)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	F_tokenize_auth_file(m, v539, v543, v301+int32(76), v299, v311)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L100
	}
L99:
	;
	v785 = v341
	goto L83
L100:
	;
	F_pfree(m, v539)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v301)+76))
	if v555 == int32(0) {
		v752 = v341
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v760 = F_FreeFile(m, v543)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L125
	}
L103:
	;
	v558 = int32(0)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	if v559 <= v558 {
		v752 = v341
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v578 = v341
	v579 = v558
	goto L105
L105:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v555)+12))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v586+v579<<(uint(int32(2))%32))))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+16))
	if v591 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v752 = v724
	goto L102
L107:
	;
	v592 = F_pstrdup(m, v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v590)))
	if v595 == int32(0) {
		v724 = v578
		goto L111
	} else {
		goto L112
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = v592
	v752 = v578
	goto L102
L111:
	;
	v733 = v579 + int32(1)
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	if v733 < v734 {
		v578 = v724
		v579 = v733
		goto L105
	} else {
		goto L124
	}
L112:
	;
	v598 = int32(0)
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v595)+4))
	if v599 <= v598 {
		v724 = v578
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v608 = v598
	v618 = v578
	goto L114
L114:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v595)+12))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v626+v608<<(uint(int32(2))%32))))
	if v630 == int32(0) {
		v696 = v618
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v724 = v696
	goto L111
L116:
	;
	v705 = v608 + int32(1)
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v595)+4))
	if v705 < v706 {
		v608 = v705
		v618 = v696
		goto L114
	} else {
		goto L123
	}
L117:
	;
	v633 = int32(0)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	if v634 <= v633 {
		v696 = v618
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v641 = v633
	v653 = v618
	goto L119
L119:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v630)+12))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v661+v641<<(uint(int32(2))%32))))
	v666 = int32(_a_F_tokenize_auth_file_3)
	v667 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1]))
	v670 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v670
	v672 = F_lappend(m, v653, v665)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L121
	}
L120:
	;
	v696 = v672
	goto L116
L121:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v667
	v677 = v641 + int32(1)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	if v677 < v678 {
		v641 = v677
		v653 = v672
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	goto L115
L124:
	;
	goto L106
L125:
	;
	if v311 != 0 {
		v785 = v752
		goto L83
	} else {
		goto L126
	}
L126:
	;
	v763 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[3]))
	F_MemoryContextDelete(m, v763)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[3])) = int32(0)
	v785 = v752
	goto L83
L128:
	;
	v813 = v785
	goto L82
L129:
	;
	if v813 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v824 = int32(_a_F_tokenize_auth_file_3)
	v825 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1]))
	v828 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v828
	v830 = F_lappend(m, v309, v813)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	v835 = v309
	goto L132
L132:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v301)+20))
	v837 = int32(0)
	v838 = base.B2i32(v836 == v837)
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495))))
	if v839 == v837 {
		v844 = v296
		v845 = v297
		v846 = v298
		v847 = v299
		v849 = v301
		v851 = v838
		v857 = v835
		v859 = v311
		v863 = v315
		v864 = v316
		v865 = v317
		v867 = v319
		goto L45
	} else {
		goto L134
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v825
	v835 = v830
	goto L132
L134:
	;
	if v836 == int32(0) {
		v307 = v495
		v309 = v835
		goto L47
	} else {
		goto L135
	}
L135:
	;
	goto L48
L136:
	;
	v1204 = int32(1)
	v1205 = v863 + v864 + v1204
	*(*int32)(unsafe.Add(mBase, uint32(v849)+28)) = v1205
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v845)))
	goto L209
L137:
	;
	v869 = v857
	goto L139
L138:
	;
	v869 = int32(1)
	goto L139
L139:
	;
	if v869 == int32(0) {
		goto L136
	} else {
		goto L140
	}
L140:
	;
	if base.B2i32(v857 == int32(0))|(v851^int32(1)) != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v1149 = int32(_a_F_tokenize_auth_file_3)
	v1150 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1]))
	v1153 = *(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v1153
	v1156 = F_palloc0(m, int32(20))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L1
	} else {
		goto L201
	}
L142:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v857)+4))
	if v877 != int32(2) {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v857)+12))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v880)+4))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v881)+12))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v882)))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v880)))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v884)+12))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v885)))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v886)))
	v888 = int32(_a_F_tokenize_auth_file_6)
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887))))
	v894 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tokenize_auth_file[4])))
	if base.B2i32(v891 == int32(0))|base.B2i32(v891 != v894) != 0 {
		v912 = v891
		v913 = v894
		goto L145
	} else {
		goto L146
	}
L144:
	;
	if v912-v913 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L145:
	;
	goto L144
L146:
	;
	v897 = v887
	v898 = v888
	goto L147
L147:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898)+1)))
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v897)+1)))
	if v902 == int32(0) {
		v912 = v902
		v913 = v901
		goto L145
	} else {
		goto L149
	}
L148:
	;
	v912 = v902
	v913 = v901
	goto L145
L149:
	;
	v905 = int32(1)
	if v902 == v901 {
		v897 = v897 + v905
		v898 = v898 + v905
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v883)))
	F_tokenize_include_file(m, v844, v917, v846, v847, v859, int32(0), v849+int32(20))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L1
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v924 = int32(_a_F_tokenize_auth_file_7)
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887))))
	v930 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tokenize_auth_file[5])))
	if base.B2i32(v927 == int32(0))|base.B2i32(v927 != v930) != 0 {
		v948 = v927
		v949 = v930
		goto L159
	} else {
		goto L160
	}
L154:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v849)+20))
	if v923 != 0 {
		goto L141
	} else {
		goto L155
	}
L155:
	;
	goto L136
L156:
	;
	F_pfree(m, v958)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L1
	} else {
		goto L199
	}
L157:
	;
	v1064 = v1021
	goto L195
L158:
	;
	if v948-v949 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L159:
	;
	goto L158
L160:
	;
	v933 = v887
	v934 = v924
	goto L161
L161:
	;
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934)+1)))
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933)+1)))
	if v938 == int32(0) {
		v948 = v938
		v949 = v937
		goto L159
	} else {
		goto L163
	}
L162:
	;
	v948 = v938
	v949 = v937
	goto L159
L163:
	;
	v941 = int32(1)
	if v938 == v937 {
		v933 = v933 + v941
		v934 = v934 + v941
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v883)))
	v958 = F_GetConfFilesInDir(m, v953, v844, v847, v849+int32(76), v849+int32(20))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v1024 = int32(_a_F_tokenize_auth_file_8)
	v1027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887))))
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_tokenize_auth_file[6])))
	if base.B2i32(v1027 == int32(0))|base.B2i32(v1027 != v1030) != 0 {
		v1048 = v1027
		v1049 = v1030
		goto L186
	} else {
		goto L187
	}
L168:
	;
	if v958 == int32(0) {
		goto L141
	} else {
		goto L169
	}
L169:
	;
	F_initStringInfo(m, v849+int32(60))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v966 = int32(0)
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v849)+76))
	if v967 <= v966 {
		goto L156
	} else {
		goto L171
	}
L171:
	;
	v974 = v966
	goto L172
L172:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v958+v974<<(uint(int32(2))%32))))
	F_tokenize_include_file(m, v844, v997, v846, v847, v859, int32(0), v849+int32(20))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L174
	}
L173:
	;
	v1021 = int32(0)
	if v1021 < v1019 {
		goto L157
	} else {
		goto L184
	}
L174:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v849)+20))
	if v1003 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v1005 = v849 + int32(60)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v849)+64))
	if int32(0) < v1006 {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	goto L177
L177:
	;
	v1018 = v974 + int32(1)
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v849)+76))
	if v1018 < v1019 {
		v974 = v1018
		goto L172
	} else {
		goto L183
	}
L178:
	;
	F_appendStringInfoChar(m, v1005, int32(10))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L181
	}
L179:
	;
	v1013 = v1003
	goto L180
L180:
	;
	F_appendStringInfoString(m, v1005, v1013)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L1
	} else {
		goto L182
	}
L181:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v849)+20))
	v1013 = v1012
	goto L180
L182:
	;
	goto L177
L183:
	;
	goto L173
L184:
	;
	goto L156
L185:
	;
	if v1048-v1049 != 0 {
		goto L141
	} else {
		goto L192
	}
L186:
	;
	goto L185
L187:
	;
	v1033 = v887
	v1034 = v1024
	goto L188
L188:
	;
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034)+1)))
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033)+1)))
	if v1038 == int32(0) {
		v1048 = v1038
		v1049 = v1037
		goto L186
	} else {
		goto L190
	}
L189:
	;
	v1048 = v1038
	v1049 = v1037
	goto L186
L190:
	;
	v1041 = int32(1)
	if v1038 == v1037 {
		v1033 = v1033 + v1041
		v1034 = v1034 + v1041
		goto L188
	} else {
		goto L191
	}
L191:
	;
	goto L189
L192:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v883)))
	F_tokenize_include_file(m, v844, v1051, v846, v847, v859, int32(1), v849+int32(20))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v849)+20))
	if v1057 == int32(0) {
		goto L136
	} else {
		goto L194
	}
L194:
	;
	goto L141
L195:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v958+v1064<<(uint(int32(2))%32))))
	F_pfree(m, v1087)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L1
	} else {
		goto L197
	}
L196:
	;
	goto L156
L197:
	;
	v1091 = v1064 + int32(1)
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v849)+76))
	if v1091 < v1092 {
		v1064 = v1091
		goto L195
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v849)+64))
	if v1120 == int32(0) {
		goto L136
	} else {
		goto L200
	}
L200:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v849)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v849)+20)) = v1123
	goto L141
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1156))) = v857
	v1159 = F_pstrdup(m, v844)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1156)+8)) = v863
	*(*int32)(unsafe.Add(mBase, uint32(v1156)+4)) = v1159
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v849)+44))
	v1164 = F_pstrdup(m, v1163)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1156)+12)) = v1164
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v849)+20))
	if v1167 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v1168 = F_pstrdup(m, v1167)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L1
	} else {
		goto L207
	}
L205:
	;
	v1171 = int32(0)
	goto L206
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1156)+16)) = v1171
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v846)))
	v1174 = F_lappend(m, v1173, v1156)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L1
	} else {
		goto L208
	}
L207:
	;
	v1171 = v1168
	goto L206
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v846))) = v1174
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[1])) = v1150
	goto L136
L209:
	;
	if int32(base.Ui32(v1207)>>(uint(int32(4))%32))&v1204 == int32(0) {
		v72 = v844
		v73 = v845
		v74 = v846
		v75 = v847
		v77 = v849
		v87 = v859
		v91 = v1205
		v93 = v865
		v95 = v867
		goto L10
	} else {
		goto L210
	}
L210:
	;
	goto L11
L211:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1219)+32))
	*(*int32)(unsafe.Add(mBase, _c_F_tokenize_auth_file[0])) = v1243
	m.G0 = v1219 + int32(80)
	return
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
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v151
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v152
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
	F_errdetail(m, int32(_a_F_transformAggregateCall_1), int32(0))
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
	F_errfinish(m, int32(_a_F_transformAggregateCall_2), int32(221), int32(_a_F_transformAggregateCall_3))
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = l2
	if int32(0) < l4 {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v22 = F_attnumTypeId(m, v21, l4)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v28 = int32(4)
			v33 = v26 + v27<<(uint(v28)%32) + l4*int32(100)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v33-v28)))
			if l1 == int32(0) {
				v73 = F_exprType(m, l1)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					if l5 == int32(0) {
						v124 = v73
						v128 = F_coerce_to_target_type(m, l0, l1, v124, v22, v37, int32(1), int32(2), int32(-1))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							if v128 != 0 {
								v161 = v128
								*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v17
								m.G0 = v15 + int32(32)
								return v161
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(67141764))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return int32(0)
									} else {
										v137 = F_format_type_be(m, v22)
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return int32(0)
										} else {
											v139 = F_format_type_be(m, v124)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v139
												*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v137
												*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l3
												F_errmsg(m, int32(_a_F_transformAssignedExpr_0), v15+int32(16))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(_a_F_transformAssignedExpr_1), int32(0))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														v153 = F_exprLocation(m, l1)
														mBase = m.M
														F_parser_errposition(m, l0, v153)
														mBase = m.M
														v155 = m.ExcPending
														if v155 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_transformAssignedExpr_2), int32(596), int32(_a_F_transformAssignedExpr_3))
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
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
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)))
						if v77 == int32(1) {
							v80 = F_makeNullConst(m, v22, v37, v34)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v89 = v80
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
								v93 = F_transformAssignmentIndirection(m, l0, v89, l3, int32(0), v22, v37, v34, l5, v91, l1, int32(1), l6)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									v161 = v93
									*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v17
									m.G0 = v15 + int32(32)
									return v161
								}
							}
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
							v86 = F_makeVar(m, v83, base.I32_extend16_s(l4), v22, v37, v34, int32(0))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v86)+44)) = l6
								v89 = v86
								v91 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
								v93 = F_transformAssignmentIndirection(m, l0, v89, l3, int32(0), v22, v37, v34, l5, v91, l1, int32(1), l6)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									v161 = v93
									*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v17
									m.G0 = v15 + int32(32)
									return v161
								}
							}
						}
					}
				}
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v40 != int32(57) {
					v73 = F_exprType(m, l1)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						if l5 == int32(0) {
							v124 = v73
							v128 = F_coerce_to_target_type(m, l0, l1, v124, v22, v37, int32(1), int32(2), int32(-1))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								if v128 != 0 {
									v161 = v128
									*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v17
									m.G0 = v15 + int32(32)
									return v161
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67141764))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											v137 = F_format_type_be(m, v22)
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												v139 = F_format_type_be(m, v124)
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v139
													*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v137
													*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l3
													F_errmsg(m, int32(_a_F_transformAssignedExpr_0), v15+int32(16))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errhint(m, int32(_a_F_transformAssignedExpr_1), int32(0))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															v153 = F_exprLocation(m, l1)
															mBase = m.M
															F_parser_errposition(m, l0, v153)
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_transformAssignedExpr_2), int32(596), int32(_a_F_transformAssignedExpr_3))
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
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
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)))
							if v77 == int32(1) {
								v80 = F_makeNullConst(m, v22, v37, v34)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v89 = v80
									v91 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
									v93 = F_transformAssignmentIndirection(m, l0, v89, l3, int32(0), v22, v37, v34, l5, v91, l1, int32(1), l6)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										v161 = v93
										*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v17
										m.G0 = v15 + int32(32)
										return v161
									}
								}
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
								v86 = F_makeVar(m, v83, base.I32_extend16_s(l4), v22, v37, v34, int32(0))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v86)+44)) = l6
									v89 = v86
									v91 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
									v93 = F_transformAssignmentIndirection(m, l0, v89, l3, int32(0), v22, v37, v34, l5, v91, l1, int32(1), l6)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										v161 = v93
										*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v17
										m.G0 = v15 + int32(32)
										return v161
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v37
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v22
					if l5 == int32(0) {
						v48 = F_exprType(m, l1)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v124 = v48
							v128 = F_coerce_to_target_type(m, l0, l1, v124, v22, v37, int32(1), int32(2), int32(-1))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								if v128 != 0 {
									v161 = v128
									*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v17
									m.G0 = v15 + int32(32)
									return v161
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67141764))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											v137 = F_format_type_be(m, v22)
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												v139 = F_format_type_be(m, v124)
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v139
													*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v137
													*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l3
													F_errmsg(m, int32(_a_F_transformAssignedExpr_0), v15+int32(16))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errhint(m, int32(_a_F_transformAssignedExpr_1), int32(0))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															v153 = F_exprLocation(m, l1)
															mBase = m.M
															F_parser_errposition(m, l0, v153)
															mBase = m.M
															v155 = m.ExcPending
															if v155 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_transformAssignedExpr_2), int32(596), int32(_a_F_transformAssignedExpr_3))
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
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
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								if v52 == int32(78) {
									F_errmsg(m, int32(_a_F_transformAssignedExpr_4), int32(0))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, l0, l6)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_transformAssignedExpr_2), int32(512), int32(_a_F_transformAssignedExpr_3))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
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
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, l0, l6)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_transformAssignedExpr_2), int32(517), int32(_a_F_transformAssignedExpr_3))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
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
		v98 = m.ExcPending
		if v98 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = l3
				F_errmsg(m, int32(_a_F_transformAssignedExpr_6), v15)
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return int32(0)
				} else {
					F_parser_errposition(m, l0, l6)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_transformAssignedExpr_2), int32(485), int32(_a_F_transformAssignedExpr_3))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
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
																F_errfinish(m, int32(_a_F_transformAssignmentSubscripts_1), int32(1004), int32(_a_F_transformAssignmentSubscripts_2))
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
															F_errfinish(m, int32(_a_F_transformAssignmentSubscripts_1), int32(1004), int32(_a_F_transformAssignmentSubscripts_2))
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
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
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int64
	_ = v224
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
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
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(144)
	m.G0 = v20
	F_check_stack_depth(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v26 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L164
	}
L4:
	;
	m.G0 = v20 + int32(144)
	return v635
L5:
	;
	v528 = F_make_parsestate(m, l0)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L138
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L133
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L125
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v29 != 0 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L120
	}
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v30 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v33 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v34 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v35 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v36 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v38 = F_palloc0(m, int32(36))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = int32(142)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v42 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+32)))
	v44 = v43
	goto L20
L19:
	;
	v44 = v5
	goto L20
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v45
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)))
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+8)) = uint8(v47)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v53 = F_transformSetOperationTree(m, l0, v49, int32(0), v20+int32(92))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v20)+92))
	v57 = int32(0)
	if base.B2i32(l2 == v57)|base.B2i32(v44&int32(1) == v57) == v57 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v66 == int32(142) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v204 = F_transformSetOperationTree(m, l0, v200, int32(0), v20+int32(88))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L48
	}
L25:
	;
	v75 = v53
	goto L28
L26:
	;
	v96 = v53
	goto L27
L27:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v108+v109<<(uint(int32(2))%32)-int32(4))))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+36))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+76))
	v126 = int32(0)
	v127 = int32(1)
	v128 = v5
	goto L31
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v87 == int32(142) {
		v75 = v86
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v96 = v86
	goto L27
L30:
	;
	goto L29
L31:
	;
	v137 = int32(0)
	if v56 == v137 {
		v147 = v137
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L24
L33:
	;
	goto L32
L34:
	;
	if v117 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v141 <= v126 {
		v147 = int32(0)
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v147 = v143 + v126<<(uint(int32(2))%32)
	goto L34
L37:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v156+v126<<(uint(int32(2))%32))))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+12))
	v168 = F_pstrdup(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L45
	}
L38:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_analyzeCTETargetList(m, l0, v159, v157)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L44
	}
L39:
	;
	v157 = int32(0)
	goto L38
L40:
	;
	goto L41
L41:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if base.B2i32(v147 == int32(0))|base.B2i32(v153 <= v126) != 0 {
		v157 = v128
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	if v156 != 0 {
		goto L37
	} else {
		goto L43
	}
L43:
	;
	v157 = v128
	goto L38
L44:
	;
	goto L33
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v173 = F_makeTargetEntry(m, v170, base.I32_extend16_s(v127), v168, int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v175 = F_lappend(m, v128, v173)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v177 = int32(1)
	v126 = v126 + v177
	v127 = v127 + v177
	v128 = v175
	goto L31
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v204
	if v45 == int32(2) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v210 = int32(_a_F_transformSetOperationTree_0)
	goto L51
L50:
	;
	v210 = int32(_a_F_transformSetOperationTree_1)
	goto L51
L51:
	;
	if v45 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v213 = int32(_a_F_transformSetOperationTree_2)
	goto L54
L53:
	;
	v213 = v210
	goto L54
L54:
	;
	if v56 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v216 = v214
	goto L57
L56:
	;
	v216 = int32(0)
	goto L57
L57:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	if v217 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v220 = v218
	goto L60
L59:
	;
	v220 = int32(0)
	goto L60
L60:
	;
	if v216 != v220 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	if l3 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v224 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+28)) = v224
	*(*int64)(unsafe.Add(mBase, uint32(v38)+20)) = v224
	v235 = int32(0)
	goto L65
L65:
	;
	v246 = int32(0)
	if v56 == v246 {
		v256 = v246
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if v217 == int32(0) {
		v635 = v38
		goto L4
	} else {
		goto L70
	}
L68:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v250 <= v235 {
		v256 = int32(0)
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v256 = v252 + v235<<(uint(int32(2))%32)
	goto L67
L70:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if base.B2i32(v256 == int32(0))|base.B2i32(v261 <= v235) != 0 {
		v635 = v38
		goto L4
	} else {
		goto L71
	}
L71:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	if v264 == int32(0) {
		v635 = v38
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v264+v235<<(uint(int32(2))%32))))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	v274 = F_exprType(m, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v276 = F_exprType(m, v271)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v271
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v273
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v271
	v286 = F_list_make2_impl(m, v20+int32(28), v20+int32(24))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v290 = F_select_common_type(m, l0, v286, v213, v20+int32(84))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v293 = F_exprLocation(m, v292)
	mBase = m.M
	if v274 != int32(705) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v276 != int32(705) {
		goto L85
	} else {
		goto L86
	}
L78:
	;
	v296 = F_coerce_to_common_type(m, l0, v273, v290, v213)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	if base.Ui32(int32(1)) < base.Ui32(v298-int32(7)) {
		v306 = v273
		goto L77
	} else {
		goto L82
	}
L81:
	;
	v306 = v296
	goto L77
L82:
	;
	v303 = F_coerce_to_common_type(m, l0, v273, v290, v213)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272)+4)) = v303
	v306 = v303
	goto L77
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v319
	v328 = F_list_make2_impl(m, v20+int32(20), v20+int32(16))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L91
	}
L85:
	;
	v309 = F_coerce_to_common_type(m, l0, v271, v290, v213)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	if base.Ui32(int32(1)) < base.Ui32(v311-int32(7)) {
		v319 = v271
		goto L84
	} else {
		goto L89
	}
L88:
	;
	v319 = v309
	goto L84
L89:
	;
	v316 = F_coerce_to_common_type(m, l0, v271, v290, v213)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v270)+4)) = v316
	v319 = v316
	goto L84
L91:
	;
	v330 = F_select_common_typmod(m, v328, v290)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v319
	v340 = F_list_make2_impl(m, v20+int32(12), v20+int32(8))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v342 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+8)))
	v347 = v345
	goto L96
L95:
	;
	v347 = int32(0)
	goto L96
L96:
	;
	v350 = F_select_common_collation(m, l0, v340, v347&int32(1))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v353 = F_lappend_oid(m, v352, v290)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v353
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v357 = F_lappend_int(m, v356, v330)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v357
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v361 = F_lappend_oid(m, v360, v350)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = v361
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v364 == int32(1) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if l3 != 0 {
		goto L114
	} else {
		goto L115
	}
L102:
	;
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+8)))
	if v367 != 0 {
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v369 = v20 + int32(96)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v369)+4)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v369)+16)) = v369
	v375 = int32(_a_F_transformSetOperationTree_3)
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_transformSetOperationTree[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+8)) = v376
	*(*int32)(unsafe.Add(mBase, _c_F_transformSetOperationTree[0])) = v20 + int32(104)
	goto L106
L105:
	;
	goto L104
L106:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	v384 = F_palloc0(m, int32(20))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = int32(106)
	v388 = int32(0)
	F_get_sort_group_operators(m, v290, v388, int32(1), v388, v20+int32(140), v20+int32(136), v388, v20+int32(135))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v402 = int32(0)
	if base.B2i32(v44&int32(1) == v402)|base.B2i32(v290 != int32(2287))&base.B2i32(v290 != int32(2249)) == v402 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v412 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+135)) = uint8(v412)
	goto L111
L110:
	;
	goto L111
L111:
	;
	v414 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v384)+4)) = v414
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v20)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v384)+8)) = v416
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v20)+140))
	*(*uint16)(unsafe.Add(mBase, uint32(v384)+16)) = uint16(v414)
	*(*int32)(unsafe.Add(mBase, uint32(v384)+12)) = v418
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+135)))
	*(*uint8)(unsafe.Add(mBase, uint32(v384)+18)) = uint8(v422)
	v424 = F_lappend(m, v382, v384)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v424
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(96))+8))
	*(*int32)(unsafe.Add(mBase, _c_F_transformSetOperationTree[0])) = v430
	goto L113
L113:
	;
	goto L101
L114:
	;
	v436 = F_palloc0(m, int32(20))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v235 = v235 + int32(1)
	goto L65
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436)+16)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v436)+12)) = v350
	*(*int32)(unsafe.Add(mBase, uint32(v436)+8)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v436)+4)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v436))) = int32(57)
	v444 = int32(0)
	v447 = F_makeTargetEntry(m, v436, v444, v444, v444)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v450 = F_lappend(m, v449, v447)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v450
	goto L116
L120:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(_a_F_transformSetOperationTree_4), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v468 = F_exprLocation(m, v467)
	mBase = m.M
	F_parser_errposition(m, l0, v468)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_transformSetOperationTree_5), int32(2066), int32(_a_F_transformSetOperationTree_6))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+12))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)+8))
	v488 = v486 - int32(1)
	if base.Ui32(v488) <= base.Ui32(int32(3)) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v495
	F_errmsg(m, int32(_a_F_transformSetOperationTree_7), v20+int32(48))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L131
	}
L128:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v488<<(uint(int32(2))%32))+uint32(_c_F_transformSetOperationTree[1])))
	v495 = v493
	goto L130
L129:
	;
	v495 = int32(_a_F_transformSetOperationTree_8)
	goto L130
L130:
	;
	goto L127
L131:
	;
	F_errfinish(m, int32(_a_F_transformSetOperationTree_5), int32(2076), int32(_a_F_transformSetOperationTree_6))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v213
	F_errmsg(m, int32(_a_F_transformSetOperationTree_9), v20+int32(32))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v520 = F_exprLocation(m, v217)
	mBase = m.M
	F_parser_errposition(m, l0, v520)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_transformSetOperationTree_5), int32(2226), int32(_a_F_transformSetOperationTree_6))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	v530 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v528)+84)) = uint16(v530)
	*(*int32)(unsafe.Add(mBase, uint32(v528)+44)) = v530
	v534 = F_transformStmt(m, v528, l1)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_free_parsestate(m, v528)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v538 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v540 = F_contain_vars_of_level(m, v534, int32(1))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	if l3 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	if v540 != 0 {
		goto L3
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v603 != 0 {
		goto L157
	} else {
		goto L158
	}
L147:
	;
	v544 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v544
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v534)+76))
	if v546 == v544 {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	if v549 <= int32(0) {
		goto L146
	} else {
		goto L149
	}
L149:
	;
	v554 = int32(0)
	v560 = v5
	goto L150
L150:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v546)+12))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v570+v554<<(uint(int32(2))%32))))
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+26)))
	if v575 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L146
L152:
	;
	v578 = F_lappend(m, v560, v574)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L155
	}
L153:
	;
	v581 = v560
	goto L154
L154:
	;
	v583 = v554 + int32(1)
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	if v583 < v584 {
		v554 = v583
		v560 = v581
		goto L150
	} else {
		goto L156
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v578
	v581 = v578
	goto L154
L156:
	;
	goto L151
L157:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v608 = v604 + int32(1)
	goto L159
L158:
	;
	v608 = int32(1)
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v608
	v611 = v20 + int32(96)
	v614 = F_pg_snprintf(m, v611, int32(32), int32(_a_F_transformSetOperationTree_10), v20)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v617 = F_makeAlias(m, v611, int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v619 = int32(0)
	v621 = F_addRangeTableEntryForSubquery(m, l0, v534, v617, v619, v619)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v624 = F_palloc0(m, int32(8))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v624))) = int32(63)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v621)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v624)+4)) = v628
	v635 = v624
	goto L4
L164:
	;
	F_errcode(m, int32(_a_F_transformSetOperationTree_11))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	F_errmsg(m, int32(_a_F_transformSetOperationTree_12), int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v663 = F_locate_var_of_level(m, v534, int32(1))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	F_parser_errposition(m, l0, v663)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_transformSetOperationTree_5), int32(2138), int32(_a_F_transformSetOperationTree_6))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
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
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)) = uint8(v143)
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
func F_tsm_system_rows_handler(m *base.Module, l0 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14029(m, l0, int32(256), int32(_a_F_tsm_system_rows_handler_0), int32(_a_F_tsm_system_rows_handler_1), int32(_a_F_tsm_system_rows_handler_2), int32(_a_F_tsm_system_rows_handler_3), int32(_a_F_tsm_system_rows_handler_4), int32(20))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_tsq_mcontains(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v398 int32
	_ = v398
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v16 = F_palloc(m, v13<<(uint(int32(2))%32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if int32(0) < v20 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = v12 + int32(8)
	v28 = v24
	v29 = v20
	v32 = v2
	v36 = v2
	goto L6
L4:
	;
	v75 = v2
	goto L5
L5:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v84 = F_palloc(m, v81<<(uint(int32(2))%32))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v38 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v75 = v65
	goto L5
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v43 = v41 & int32(4095)
	v46 = F_palloc(m, v43+int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v63 = v29
	v65 = v32
	goto L10
L10:
	;
	v69 = v36 + int32(1)
	if v69 < v63 {
		v28 = v28 + int32(12)
		v29 = v63
		v32 = v65
		v36 = v69
		goto L6
	} else {
		goto L15
	}
L11:
	;
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	base.MemoryCopy(m, v46, v24+v13*int32(12)+int32(base.Ui32(v48)>>(uint(int32(12))%32)), v43)
	goto L14
L13:
	;
	goto L14
L14:
	;
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43+v46))) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v16+v32<<(uint(int32(2))%32)))) = v46
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v63 = v62
	v65 = v32 + int32(1)
	goto L10
L15:
	;
	goto L7
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v86 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v90 = v11 + int32(8)
	v95 = v90
	v96 = int32(0)
	v97 = v86
	v100 = v2
	goto L20
L18:
	;
	v143 = v2
	goto L19
L19:
	;
	F_pg_qsort(m, v16, v75, int32(4), int32(1513))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L30
	}
L20:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v105 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v143 = v131
	goto L19
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v110 = v108 & int32(4095)
	v113 = F_palloc(m, v110+int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v130 = v97
	v131 = v100
	goto L24
L24:
	;
	v136 = v96 + int32(1)
	if v136 < v130 {
		v95 = v95 + int32(12)
		v96 = v136
		v97 = v130
		v100 = v131
		goto L20
	} else {
		goto L29
	}
L25:
	;
	if v110 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	base.MemoryCopy(m, v113, v90+v81*int32(12)+int32(base.Ui32(v115)>>(uint(int32(12))%32)), v110)
	goto L28
L27:
	;
	goto L28
L28:
	;
	v121 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v110+v113))) = uint8(v121)
	*(*int32)(unsafe.Add(mBase, uint32(v84+v100<<(uint(int32(2))%32)))) = v113
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v130 = v129
	v131 = v100 + int32(1)
	goto L24
L29:
	;
	goto L21
L30:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v75) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v156 = int32(1)
	v157 = int32(0)
	goto L34
L32:
	;
	v220 = v75
	goto L33
L33:
	;
	F_pg_qsort(m, v84, v143, int32(4), int32(1513))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L47
	}
L34:
	;
	v166 = int32(2)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v16+v156<<(uint(v166)%32))))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v16+v157<<(uint(v166)%32))))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	if base.B2i32(v176 == int32(0))|base.B2i32(v176 != v179) != 0 {
		v197 = v176
		v198 = v179
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v220 = v209 + int32(1)
	goto L33
L36:
	;
	v212 = v156 + int32(1)
	if v212 != v75 {
		v156 = v212
		v157 = v209
		goto L34
	} else {
		goto L46
	}
L37:
	;
	if v197-v198 == int32(0) {
		v209 = v157
		goto L36
	} else {
		goto L44
	}
L38:
	;
	goto L37
L39:
	;
	v182 = v169
	v183 = v173
	goto L40
L40:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
	if v187 == int32(0) {
		v197 = v187
		v198 = v186
		goto L38
	} else {
		goto L42
	}
L41:
	;
	v197 = v187
	v198 = v186
	goto L38
L42:
	;
	v190 = int32(1)
	if v187 == v186 {
		v182 = v182 + v190
		v183 = v183 + v190
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v203 = v157 + int32(1)
	if v156 == v203 {
		v209 = v156
		goto L36
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v203<<(uint(int32(2))%32)))) = v169
	v209 = v203
	goto L36
L46:
	;
	goto L35
L47:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v143) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v234 = int32(1)
	v235 = int32(0)
	goto L51
L49:
	;
	v299 = v143
	goto L50
L50:
	;
	if v220 < v299 {
		goto L65
	} else {
		goto L66
	}
L51:
	;
	v244 = int32(2)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v84+v234<<(uint(v244)%32))))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v84+v235<<(uint(v244)%32))))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if base.B2i32(v254 == int32(0))|base.B2i32(v254 != v257) != 0 {
		v275 = v254
		v276 = v257
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v299 = v287 + int32(1)
	goto L50
L53:
	;
	v290 = v234 + int32(1)
	if v290 != v143 {
		v234 = v290
		v235 = v287
		goto L51
	} else {
		goto L63
	}
L54:
	;
	if v275-v276 == int32(0) {
		v287 = v235
		goto L53
	} else {
		goto L61
	}
L55:
	;
	goto L54
L56:
	;
	v260 = v247
	v261 = v251
	goto L57
L57:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	if v265 == int32(0) {
		v275 = v265
		v276 = v264
		goto L55
	} else {
		goto L59
	}
L58:
	;
	v275 = v265
	v276 = v264
	goto L55
L59:
	;
	v268 = int32(1)
	if v265 == v264 {
		v260 = v260 + v268
		v261 = v261 + v268
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v281 = v235 + int32(1)
	if v234 == v281 {
		v287 = v234
		goto L53
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84+v281<<(uint(int32(2))%32)))) = v247
	v287 = v281
	goto L53
L63:
	;
	goto L52
L64:
	;
	return v398
L65:
	;
	v398 = int32(0)
	goto L64
L66:
	;
	v305 = int32(0)
	if v299 <= v305 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	return int32(1)
L68:
	;
	goto L69
L69:
	;
	v311 = v305
	v313 = int32(0)
	goto L70
L70:
	;
	if v220 <= v311 {
		v371 = v311
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v398 = v382
	goto L64
L72:
	;
	if v371 == v220 {
		goto L65
	} else {
		goto L85
	}
L73:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v84+v313<<(uint(int32(2))%32))))
	v326 = v311
	goto L74
L74:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v16+v326<<(uint(int32(2))%32))))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	if base.B2i32(v342 == int32(0))|base.B2i32(v342 != v345) != 0 {
		v363 = v342
		v364 = v345
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L65
L76:
	;
	if v363-v364 == int32(0) {
		v371 = v326
		goto L72
	} else {
		goto L83
	}
L77:
	;
	goto L76
L78:
	;
	v348 = v325
	v349 = v339
	goto L79
L79:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+1)))
	if v353 == int32(0) {
		v363 = v353
		v364 = v352
		goto L77
	} else {
		goto L81
	}
L80:
	;
	v363 = v353
	v364 = v352
	goto L77
L81:
	;
	v356 = int32(1)
	if v353 == v352 {
		v348 = v348 + v356
		v349 = v349 + v356
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v369 = v326 + int32(1)
	if v369 != v220 {
		v326 = v369
		goto L74
	} else {
		goto L84
	}
L84:
	;
	goto L75
L85:
	;
	v382 = int32(1)
	v384 = v313 + v382
	if v299 != v384 {
		v311 = v371
		v313 = v384
		goto L70
	} else {
		goto L86
	}
L86:
	;
	goto L71
}
func F_tsqueryout(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == int32(0) {
		v15 = F_palloc(m, int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v19)
			v49 = v15
			m.G0 = v8 + int32(32)
			return v49
		}
	} else {
		v21 = int32(32)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v21
		v24 = v10 + int32(8)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v24
		v27 = F_palloc(m, v21)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v27
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v27
			v31 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v31)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v34 = int32(12)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v24 + v33*v34
			F_infix_1(m, v8+v34, int32(-1), v31)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v44 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
						v49 = v48
						m.G0 = v8 + int32(32)
						return v49
					}
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					v49 = v48
					m.G0 = v8 + int32(32)
					return v49
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
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
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
	F_slot_getsomeattrs_int(m, l0, v12)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
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
	v25 = v3
	v27 = v3
	goto L9
L7:
	;
	v72 = v3
	goto L8
L8:
	;
	v78 = int32(_a_F_tstoreReceiveSlot_detoast_0)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_tstoreReceiveSlot_detoast[0]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_tstoreReceiveSlot_detoast[0])) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_tuplestore_putvalues(m, v83, v11, v84, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L18
	}
L9:
	;
	v34 = v25 << (uint(int32(2)) % 32)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34+v35)))
	v40 = v11 + int32(20) + v25<<(uint(int32(4))%32)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+9)))
	if v41 != 0 {
		v60 = v37
		v61 = v27
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v72 = v61
	goto L8
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v62+v34))) = v60
	v66 = v25 + int32(1)
	if v66 != v12 {
		v25 = v66
		v27 = v61
		goto L9
	} else {
		goto L17
	}
L12:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)))
	if v42 != int32(_a_F_tstoreReceiveSlot_detoast_1) {
		v60 = v37
		v61 = v27
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v25))))
	if v47 != 0 {
		v60 = v37
		v61 = v27
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v48 != int32(1) {
		v60 = v37
		v61 = v27
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v51 = F_detoast_external_attr(m, v37)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v53+v27<<(uint(int32(2))%32)))) = v51
	v60 = v51
	v61 = v27 + int32(1)
	goto L11
L17:
	;
	goto L10
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tstoreReceiveSlot_detoast[0])) = v79
	if int32(0) < v72 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v95 = int32(0)
	goto L22
L20:
	;
	goto L21
L21:
	;
	return int32(1)
L22:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v95<<(uint(int32(2))%32))))
	F_pfree(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	v111 = v95 + int32(1)
	if v111 != v72 {
		v95 = v111
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
}
func F_tsvectorin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var __phi256 int32
	_ = __phi256
	var v260 int32
	_ = v260
	var __phi260 int32
	_ = __phi260
	var v262 int32
	_ = v262
	var __phi262 int32
	_ = __phi262
	var v270 int32
	_ = v270
	var __phi270 int32
	_ = __phi270
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
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
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
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v414 int32
	_ = v414
	var v424 int32
	_ = v424
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int64
	_ = v455
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v615 int32
	_ = v615
	var v624 int32
	_ = v624
	var v639 int32
	_ = v639
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v800 int32
	_ = v800
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = F_init_tsvector_parser(m, v23, v2, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v31 = F_palloc(m, int32(768))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = F_palloc(m, int32(256))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v45 = F_gettoken_tsvector(m, v26, v18+int32(-4), v18+int32(-8), v18+int32(-12), v18+int32(-16), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L1
	} else {
		goto L153
	}
L6:
	;
	m.G0 = v20 - int32(-64)
	return v800
L7:
	;
	if v45 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v49 = v34
	v52 = int32(256)
	v57 = int32(64)
	v58 = v2
	v59 = v31
	v61 = v34
	goto L11
L9:
	;
	v226 = v2
	v227 = v31
	v229 = v34
	goto L10
L10:
	;
	F_close_tsvector_parser(m, v26)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L49
	}
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	if int32(2047) <= v65 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v226 = v204
	v227 = v122
	v229 = v163
	goto L10
L13:
	;
	v68 = int32(0)
	v69 = F_errsave_start(m, v25)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v88 = v49 - v61
	if int32(_a_F_tsvectorin_0) <= v88 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	if v69 == int32(0) {
		v800 = v68
		goto L6
	} else {
		goto L17
	}
L17:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(2046)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v78
	F_errmsg(m, int32(_a_F_tsvectorin_1), v20)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errsave_finish(m, v25, int32(_a_F_tsvectorin_2), int32(215), int32(_a_F_tsvectorin_3))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v800 = v68
	goto L6
L21:
	;
	v91 = int32(0)
	v92 = F_errsave_start(m, v25)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v57 <= v58 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	if v92 == int32(0) {
		v800 = v91
		goto L6
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(_a_F_tsvectorin_4)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v88
	F_errmsg(m, int32(_a_F_tsvectorin_5), v18+int32(-48))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errsave_finish(m, v25, int32(_a_F_tsvectorin_2), int32(221), int32(_a_F_tsvectorin_3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v800 = v91
	goto L6
L29:
	;
	v115 = F_repalloc(m, v59, v57*int32(24))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	v120 = v65
	v121 = v57
	v122 = v59
	goto L31
L31:
	;
	if v52 <= v120+v88 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v120 = v119
	v121 = v57 << (uint(int32(1)) % 32)
	v122 = v115
	goto L31
L33:
	;
	v129 = v52
	v138 = v61
	goto L36
L34:
	;
	v151 = v49
	v152 = v120
	v154 = v52
	v163 = v61
	goto L35
L35:
	;
	v167 = int32(12)
	v169 = v122 + v58*v167
	v170 = int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v152<<(uint(v170)%32)&int32(4094) | v174&v170 | v88<<(uint(v167)%32)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	if v182 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v143 = v129 << (uint(int32(1)) % 32)
	v144 = F_repalloc(m, v138, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v151 = v88 + v144
	v152 = v146
	v154 = v143
	v163 = v144
	goto L35
L38:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	if v143 <= v88+v146 {
		v129 = v143
		v138 = v144
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	base.MemoryCopy(m, v151, v183, v182)
	goto L42
L41:
	;
	goto L42
L42:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	if v188 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169)+8)) = v201
	v204 = v58 + int32(1)
	v214 = F_gettoken_tsvector(m, v26, v18+int32(-4), v18+int32(-8), v18+int32(-12), v18+int32(-16), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L47
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v185 | int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+4)) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v201 = v194
	goto L43
L45:
	;
	goto L46
L46:
	;
	v195 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v169)+4)) = v195
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v185 & int32(-2)
	v201 = v195
	goto L43
L47:
	;
	if v214 != 0 {
		v49 = v186 + v151
		v52 = v154
		v57 = v121
		v58 = v204
		v59 = v122
		v61 = v163
		goto L11
	} else {
		goto L48
	}
L48:
	;
	goto L12
L49:
	;
	if v25 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v226 <= int32(0) {
		v684 = v2
		v688 = v226
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v237 != int32(447) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+4)))
	if v240 != int32(1) {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v243 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v243)
	v800 = int32(0)
	goto L6
L54:
	;
	v696 = v688 << (uint(int32(2)) % 32)
	v699 = v696 + v684 + int32(8)
	v700 = F_palloc0(m, v699)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L137
	}
L55:
	;
	if v226 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	v533 = int32(1)
	v537 = int32(base.Ui32(v532)>>(uint(v533)%32))&int32(2047) + v529
	if v532&v533 != 0 {
		goto L113
	} else {
		goto L114
	}
L57:
	;
	v515 = v227
	v529 = v2
	goto L56
L58:
	;
	goto L59
L59:
	;
	F_qsort_arg(m, v227, v226, int32(12), int32(1519), v229)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	__phi256 = v227
	__phi260 = v227 + int32(12)
	__phi262 = v227
	__phi270 = v2
	v256 = __phi256
	v260 = __phi260
	v262 = __phi262
	v270 = __phi270
	goto L61
L61:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v274 = int32(1)
	v276 = int32(2047)
	v277 = int32(base.Ui32(v273)>>(uint(v274)%32)) & v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v282 = int32(base.Ui32(v278)>>(uint(v274)%32)) & v276
	if v277 == v282 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v515 = v492
	v529 = v506
	goto L56
L63:
	;
	v509 = int32(12)
	v510 = v260 + v509
	v513 = base.I32_div_s(v510-v227, v509)
	if v513 < v226 {
		__phi256 = v492
		__phi260 = v510
		__phi262 = v260
		__phi270 = v506
		v256 = __phi256
		v260 = __phi260
		v262 = __phi262
		v270 = __phi270
		goto L61
	} else {
		goto L112
	}
L64:
	;
	if v273&int32(1) == int32(0) {
		v492 = v256
		v506 = v270
		goto L63
	} else {
		goto L103
	}
L65:
	;
	v284 = int32(12)
	v286 = v229 + int32(base.Ui32(v273)>>(uint(v284)%32))
	v289 = v229 + int32(base.Ui32(v278)>>(uint(v284)%32))
	if v277 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	goto L67
L67:
	;
	v337 = v282 + v270
	if v278&int32(1) != 0 {
		goto L82
	} else {
		goto L83
	}
L68:
	;
	if v334 == int32(0) {
		goto L64
	} else {
		goto L81
	}
L69:
	;
	v334 = int32(0)
	goto L68
L70:
	;
	goto L71
L71:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v295 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v296 = v286
	v297 = v289
	v298 = v277
	v299 = v295
	goto L76
L73:
	;
	v322 = v289
	v326 = int32(0)
	goto L74
L74:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	v334 = v326 - v327
	goto L68
L75:
	;
	v322 = v317
	v326 = v319
	goto L74
L76:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	if base.B2i32(v299 != v301)|base.B2i32(v301 == int32(0)) != 0 {
		v317 = v297
		v319 = v299
		goto L75
	} else {
		goto L78
	}
L77:
	;
	v317 = v311
	v319 = int32(0)
	goto L75
L78:
	;
	v307 = v298 - int32(1)
	if v307 == int32(0) {
		v317 = v297
		v319 = v299
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v310 = int32(1)
	v311 = v297 + v310
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+1)))
	if v312 != 0 {
		v296 = v296 + v310
		v297 = v311
		v298 = v307
		v299 = v312
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	goto L67
L82:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	if int32(2) <= v340 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v447 = v337
	goto L84
L84:
	;
	v451 = v256 + int32(12)
	if v256 == v262 {
		goto L100
	} else {
		goto L101
	}
L85:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	F_pg_qsort(m, v343, v340, int32(2), int32(1520))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	v414 = v340
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+8)) = v414
	v424 = int32(1)
	v447 = (v337+v424)&int32(-2) + v414<<(uint(v424)%32) + int32(2)
	goto L84
L88:
	;
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v343))))
	v352 = v350
	v353 = v343
	v355 = v343 + int32(2)
	goto L89
L89:
	;
	v368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v355))))
	v369 = int32(_a_F_tsvectorin_6)
	v370 = v368 & v369
	if v370 != v352&v369 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v414 = (v399 - v343 + int32(2)) >> (uint(int32(1)) % 32)
	goto L87
L91:
	;
	goto L90
L92:
	;
	v393 = v355 + int32(2)
	if (v393-v343)>>(uint(int32(1))%32) < v340 {
		v352 = v390
		v353 = v391
		v355 = v393
		goto L89
	} else {
		goto L99
	}
L93:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v353)+2)) = uint16(v368)
	v376 = v353 + int32(2)
	if int32(508) < v376-v343 {
		v399 = v376
		goto L91
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v382 = int32(14)
	if base.Ui32(int32(base.Ui32(v368)>>(uint(v382)%32))) <= base.Ui32(int32(base.Ui32(v352&int32(_a_F_tsvectorin_7))>>(uint(v382)%32))) {
		v390 = v352
		v391 = v353
		goto L92
	} else {
		goto L98
	}
L96:
	;
	if v370 != int32(_a_F_tsvectorin_6) {
		v390 = v368
		v391 = v376
		goto L92
	} else {
		goto L97
	}
L97:
	;
	v399 = v376
	goto L91
L98:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v353))) = uint16(v368)
	v390 = v368
	v391 = v353
	goto L92
L99:
	;
	v399 = v391
	goto L91
L100:
	;
	v492 = v451
	v506 = v447
	goto L63
L101:
	;
	goto L102
L102:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v451)+8)) = v453
	v455 = *(*int64)(unsafe.Add(mBase, uint32(v260)))
	*(*int64)(unsafe.Add(mBase, uint32(v451))) = v455
	v492 = v451
	v506 = v447
	goto L63
L103:
	;
	if v278&int32(1) != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v262)+20))
	v466 = v464 + v465
	v469 = F_repalloc(m, v463, v466<<(uint(int32(1))%32))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v278 | int32(1)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v262)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+4)) = v488
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v262)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+8)) = v490
	v492 = v256
	v506 = v270
	goto L63
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+4)) = v469
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v262)+20))
	v474 = v472 << (uint(int32(1)) % 32)
	if v474 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v262)+16))
	base.MemoryCopy(m, v469+v475<<(uint(int32(1))%32), v479, v474)
	goto L110
L109:
	;
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+8)) = v466
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v262)+16))
	F_pfree(m, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v492 = v256
	v506 = v270
	goto L63
L112:
	;
	goto L62
L113:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v515)+8))
	if int32(2) <= v540 {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v639 = v537
	goto L115
L115:
	;
	v651 = int32(12)
	v654 = base.I32_div_s(v515-v227+v651, v651)
	if v639 < int32(_a_F_tsvectorin_0) {
		v684 = v639
		v688 = v654
		goto L54
	} else {
		goto L131
	}
L116:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v515)+4))
	F_pg_qsort(m, v543, v540, int32(2), int32(1520))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	v615 = v540
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v515)+8)) = v615
	v624 = int32(1)
	v639 = (v537+v624)&int32(-2) + v615<<(uint(v624)%32) + int32(2)
	goto L115
L119:
	;
	v550 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v543))))
	v552 = v550
	v553 = v543
	v555 = v543 + int32(2)
	goto L120
L120:
	;
	v568 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v555))))
	v569 = int32(_a_F_tsvectorin_6)
	v570 = v568 & v569
	if v570 != v552&v569 {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	v615 = (v599 - v543 + int32(2)) >> (uint(int32(1)) % 32)
	goto L118
L122:
	;
	goto L121
L123:
	;
	v593 = v555 + int32(2)
	if (v593-v543)>>(uint(int32(1))%32) < v540 {
		v552 = v590
		v553 = v591
		v555 = v593
		goto L120
	} else {
		goto L130
	}
L124:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v553)+2)) = uint16(v568)
	v576 = v553 + int32(2)
	if int32(508) < v576-v543 {
		v599 = v576
		goto L122
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v582 = int32(14)
	if base.Ui32(int32(base.Ui32(v568)>>(uint(v582)%32))) <= base.Ui32(int32(base.Ui32(v552&int32(_a_F_tsvectorin_7))>>(uint(v582)%32))) {
		v590 = v552
		v591 = v553
		goto L123
	} else {
		goto L129
	}
L127:
	;
	if v570 != int32(_a_F_tsvectorin_6) {
		v590 = v568
		v591 = v576
		goto L123
	} else {
		goto L128
	}
L128:
	;
	v599 = v576
	goto L122
L129:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v553))) = uint16(v568)
	v590 = v568
	v591 = v553
	goto L123
L130:
	;
	v599 = v591
	goto L122
L131:
	;
	v657 = int32(0)
	v658 = F_errsave_start(m, v25)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	if v658 == int32(0) {
		v800 = v657
		goto L6
	} else {
		goto L133
	}
L133:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(_a_F_tsvectorin_4)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v639
	F_errmsg(m, int32(_a_F_tsvectorin_8), v18+int32(-32))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errsave_finish(m, v25, int32(_a_F_tsvectorin_2), int32(274), int32(_a_F_tsvectorin_3))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v800 = v657
	goto L6
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700)+4)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v700))) = v699 << (uint(int32(2)) % 32)
	if v688 <= int32(0) {
		v800 = v700
		goto L6
	} else {
		goto L138
	}
L138:
	;
	v709 = v700 + int32(8)
	v710 = v709 + v696
	v714 = int32(0)
	v727 = v2
	goto L139
L139:
	;
	v731 = v227 + v714*int32(12)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v736 = int32(base.Ui32(v732)>>(uint(int32(1))%32)) & int32(2047)
	if v736 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v800 = v700
	goto L6
L141:
	;
	base.MemoryCopy(m, v710+v727, v229+int32(base.Ui32(v732)>>(uint(int32(12))%32)), v736)
	goto L143
L142:
	;
	goto L143
L143:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v747 = v742&int32(4095) | v727<<(uint(int32(12))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v731))) = v747
	v749 = int32(1)
	v753 = int32(base.Ui32(v742)>>(uint(v749)%32))&int32(2047) + v727
	if v742&v749 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v731)+8))
	if int32(_a_F_tsvectorin_9) <= v759 {
		goto L5
	} else {
		goto L147
	}
L145:
	;
	v786 = v753
	v787 = v747
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v709+v714<<(uint(int32(2))%32)))) = v787
	v790 = v714 + int32(1)
	if v790 != v688 {
		v714 = v790
		v727 = v786
		goto L139
	} else {
		goto L152
	}
L147:
	;
	v762 = int32(1)
	v765 = (v753 + v762) & int32(-2)
	*(*uint16)(unsafe.Add(mBase, uint32(v710+v765))) = uint16(v759)
	v769 = v765 + int32(2)
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v731)+8))
	v772 = v770 << (uint(v762) % 32)
	if v772 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v731)+4))
	base.MemoryCopy(m, v710+v769, v774, v772)
	goto L150
L149:
	;
	goto L150
L150:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v731)+8))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v731)+4))
	F_pfree(m, v777)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v786 = v776<<(uint(int32(1))%32) + v769
	v787 = v783
	goto L146
L152:
	;
	goto L140
L153:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorin_10), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_tsvectorin_2), int32(292), int32(_a_F_tsvectorin_3))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tt_process_call(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v45 int32
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
	var v58 int32
	_ = v58
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == v2 {
		v58 = v2
		m.G0 = v8 + int32(48)
		return v58
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v11+v14*int32(12))))
		if v18 == int32(0) {
			v58 = v2
			m.G0 = v8 + int32(48)
			return v58
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
			v23 = v8 + int32(16)
			v25 = F_pg_sprintf(m, v23, int32(_a_F_tt_process_call_0), v8)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
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
					return int32(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
					v45 = F_HeapTupleHeaderGetDatum(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
						F_pfree(m, v47)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
							F_pfree(m, v50)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
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
	var v27 int32
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
				v27 = F_OidFunctionCall1Coll(m, v24, v22, v22)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v27
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
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int64
	_ = v336
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int64
	_ = v386
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int64
	_ = v409
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v459 int32
	_ = v459
	var v469 int32
	_ = v469
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
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
	v498 = m.ExcPending
	if v498 != 0 {
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
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v469)
	m.G0 = v17 + int32(16)
	return v459
L6:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v446 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v445 + v446
	*(*int32)(unsafe.Add(mBase, uint32(v435)+8)) = l1
	v450 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v435))) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v435)+4)) = v446
	v459 = v435
	v469 = v450
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
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
		v435 = v241
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
	v435 = v416
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
	v298 = v252
	v300 = v259
	goto L58
L58:
	;
	v301 = v300 & v298
	if base.Ui32(v253) < base.Ui32(v301) {
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
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v271)+20))
	v286 = m.T0[v285].(func(*base.Module, int32, int32, int32) int32)(m, v271, v262, v17+int32(15))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L22
	} else {
		goto L64
	}
L63:
	;
	v459 = v249
	v469 = int32(1)
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
	if v286 != 0 {
		v459 = v249
		v469 = int32(1)
		goto L5
	} else {
		goto L66
	}
L66:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v249)+8))
	v298 = v294
	v300 = v295
	goto L58
L67:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v305 = v253 + v303
	goto L69
L68:
	;
	v305 = v253
	goto L69
L69:
	;
	v308 = (v253 + int32(1)) & v298
	if base.Ui32(v305-v301) < base.Ui32(v251) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v314 = v236 + v308*int32(12)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	if v315 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v404 = v251 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v404) {
		goto L89
	} else {
		goto L90
	}
L73:
	;
	v319 = v308
	v321 = int32(0)
	goto L76
L74:
	;
	v351 = v308
	v354 = v314
	goto L75
L75:
	;
	if v351 != v253 {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	v331 = v321 + int32(1)
	if int32(151) <= v331 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v351 = v343
	v354 = v346
	goto L75
L78:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v336 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v334), base.F64_convert_i64_u(v336)), float64(0.1)) != 0 {
		goto L4
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v343 = (v319 + int32(1)) & v298
	v346 = v236 + v343*int32(12)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if v347 != 0 {
		v319 = v343
		v321 = v331
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
	v366 = v351
	v369 = v354
	goto L86
L84:
	;
	goto L85
L85:
	;
	v435 = v249
	goto L6
L86:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v380 = v377 & (v366 - int32(1))
	v383 = v236 + v380*int32(12)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+8)) = v384
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v383)))
	*(*int64)(unsafe.Add(mBase, uint32(v369))) = v386
	if v380 != v253 {
		v366 = v380
		v369 = v383
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
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v409 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v407), base.F64_convert_i64_u(v409)), float64(0.1)) != 0 {
		goto L4
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v416 = v236 + v308*int32(12)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
	if v417 != 0 {
		v249 = v416
		v251 = v404
		v252 = v298
		v253 = v308
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
	v425 = m.ExcPending
	if v425 != 0 {
		goto L22
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_tuplehash_insert_hash_internal_2), int32(630), int32(_a_F_tuplehash_insert_hash_internal_3))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
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
	v502 = m.ExcPending
	if v502 != 0 {
		goto L22
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_tuplehash_insert_hash_internal_2), int32(327), int32(_a_F_tuplehash_insert_hash_internal_5))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
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
