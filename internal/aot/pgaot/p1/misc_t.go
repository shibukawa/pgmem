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
	var v20 int32
	_ = v20
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
				v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[916])))
				if v17 != int32(1) {
					F_char2wchar(m, v6+int32(4), int32(3), l0, v12, int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
						if base.Ui32(int32(10)) <= base.Ui32(v38-int32(48)) {
							v45 = F_iswalpha(m, v38)
							mBase = m.M
							v48 = base.B2i32(v45 != int32(0))
						} else {
							v48 = int32(1)
						}
						v49 = v48
						m.G0 = v6 + int32(16)
						return v49
					}
				} else {
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					v49 = base.B2i32(base.Ui32(v20-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v20|int32(32)-int32(97)) < base.Ui32(int32(26)))
					m.G0 = v6 + int32(16)
					return v49
				}
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				v49 = base.B2i32(base.Ui32(v20-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v20|int32(32)-int32(97)) < base.Ui32(int32(26)))
				m.G0 = v6 + int32(16)
				return v49
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
	var v18 int32
	_ = v18
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
			v15 = int32(*(*uint8)(unsafe.Add(mBase, _consts[916])))
			if v15 != int32(1) {
				F_char2wchar(m, v6+int32(4), int32(3), l0, v8, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
					if base.Ui32(int32(10)) <= base.Ui32(v36-int32(48)) {
						v43 = F_iswalpha(m, v36)
						mBase = m.M
						v46 = base.B2i32(v43 != int32(0))
					} else {
						v46 = int32(1)
					}
					v47 = v46
					m.G0 = v6 + int32(16)
					return v47
				}
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				v47 = base.B2i32(base.Ui32(v18-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v18|int32(32)-int32(97)) < base.Ui32(int32(26)))
				m.G0 = v6 + int32(16)
				return v47
			}
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			v47 = base.B2i32(base.Ui32(v18-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v18|int32(32)-int32(97)) < base.Ui32(int32(26)))
			m.G0 = v6 + int32(16)
			return v47
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
			v8 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8+(v4^int32(-1))<<(uint(int32(2))%32))))
			v22 = v14
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
			v36 = *(*int32)(unsafe.Add(mBase, _consts[8]))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+(v32^int32(-1))<<(uint(int32(6))%32))+16))
			v51 = v42
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
	if v20 <= v19 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L15
	}
L4:
	;
	m.G0 = v6 + int32(16)
	return v48
L5:
	;
	v48 = v2
	goto L4
L6:
	;
	goto L7
L7:
	;
	v23 = v19
	v25 = v2
	goto L8
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v23))))
	if v28 == int32(1) {
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v48 = v40
	goto L4
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v23<<(uint(int32(2))%32))))
	v36 = F_text_to_cstring(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v38 = F_makeString(m, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v40 = F_lappend(m, v25, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v43 = v23 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v43 < v44 {
		v23 = v43
		v25 = v40
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(161930), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(516112), int32(2098), int32(79666))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
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
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
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
	v224 = m.ExcPending
	if v224 != 0 {
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
	return v213
L7:
	;
	F_pfree(m, v205)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
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
		v213 = int32(0)
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
		v204 = v112
		v205 = v27
		goto L7
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v213 = v112
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
	v157 = int32(1)
	v158 = v120 + v157
	if v122&v157 != 0 {
		goto L61
	} else {
		goto L62
	}
L51:
	;
	v131 = int32(4)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v133&int32(254) == int32(2) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v146 = int32(1)
	if v127 != 0 {
		v156 = int32(base.Ui32(v125)>>(uint(v146)%32)) - v146
		goto L50
	} else {
		goto L60
	}
L54:
	;
	v142 = v131
	goto L56
L55:
	;
	v142 = base.B2i32(v133 == int32(18)) << (uint(v131) % 32)
	goto L56
L56:
	;
	if v133 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v145 = v131
	goto L59
L58:
	;
	v145 = v142
	goto L59
L59:
	;
	v156 = v145
	goto L50
L60:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v156 = int32(base.Ui32(v150)>>(uint(int32(2))%32)) - int32(4)
	goto L50
L61:
	;
	v163 = v158
	goto L63
L62:
	;
	v163 = v120 + int32(4)
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
	v194 = F_varstr_cmp(m, v128, v156, v163, v193, v9)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L75
	}
L65:
	;
	v166 = int32(4)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v168&int32(254) == int32(2) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v181 = int32(1)
	if v122&v181 != 0 {
		v193 = int32(base.Ui32(v122)>>(uint(v181)%32)) - v181
		goto L64
	} else {
		goto L74
	}
L68:
	;
	v177 = v166
	goto L70
L69:
	;
	v177 = base.B2i32(v168 == int32(18)) << (uint(v166) % 32)
	goto L70
L70:
	;
	if v168 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v180 = v166
	goto L73
L72:
	;
	v180 = v177
	goto L73
L73:
	;
	v193 = v180
	goto L64
L74:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v193 = int32(base.Ui32(v187)>>(uint(int32(2))%32)) - int32(4)
	goto L64
L75:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v196 != v115 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_pfree(m, v115)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v201 = base.B2i32(v194 == int32(0))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v120 == v202 {
		v213 = v201
		goto L6
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v204 = v201
	v205 = v120
	goto L7
L81:
	;
	v213 = v204
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
	F_errmsg(m, int32(257570), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errhint(m, int32(601086), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(524313), int32(1648), int32(113106))
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
func F_texteqfast(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_DirectFunctionCall2Coll(m, int32(1560), int32(100), l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1559), v3, v4, v5)
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
	v4 = F_DirectFunctionCall1Coll(m, int32(1586), int32(100), l0)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_texticregexne(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = v8 + int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			v18 = int32(1)
			v19 = v17 & v18
			if v17 == v18 {
				v22 = int32(4)
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				if v24&int32(254) == int32(2) {
					v33 = v22
				} else {
					v33 = base.B2i32(v24 == int32(18)) << (uint(v22) % 32)
				}
				if v24 == int32(1) {
					v36 = v22
				} else {
					v36 = v33
				}
				v47 = v36
			} else {
				v37 = int32(1)
				if v19 != 0 {
					v47 = int32(base.Ui32(v17)>>(uint(v37)%32)) - v37
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v50 = F_RE_compile_and_cache(m, v15, int32(27), v49)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v56 = F_palloc(m, v47<<(uint(int32(2))%32)+int32(4))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					if v19 != 0 {
						v60 = v13
					} else {
						v60 = v8 + int32(4)
					}
					v61 = F_pg_mb2wchar_with_len(m, v60, v56, v47)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = int32(0)
						v66 = F_RE_wchar_execute(m, v56, v61, v63, v63, v63)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v56)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								return v66 ^ int32(1)
							}
						}
					}
				}
			}
		}
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
			v15 = *(*int32)(unsafe.Add(mBase, _consts[356]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v16*int32(28))+uint32(_consts[355])))
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
						v43 = int32(4)
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
						if v45&int32(254) == int32(2) {
							v54 = v43
						} else {
							v54 = base.B2i32(v45 == int32(18)) << (uint(v43) % 32)
						}
						if v45 == int32(1) {
							v57 = v43
						} else {
							v57 = v54
						}
						v68 = v57
					} else {
						v58 = int32(1)
						if v39 != 0 {
							v68 = int32(base.Ui32(v37)>>(uint(v58)%32)) - v58
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
							v68 = int32(base.Ui32(v62)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v69 = F_pg_mbstrlen_with_len(m, v40, v68)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v71 = F_text_overlay(m, v6, v11, v13, v69)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							return v71
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
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
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
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
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
						v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
						if v49 == int32(1) {
							v52 = int32(4)
							v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
							if v54&int32(254) == int32(2) {
								v63 = v52
							} else {
								v63 = base.B2i32(v54 == int32(18)) << (uint(v52) % 32)
							}
							if v54 == int32(1) {
								v66 = v52
							} else {
								v66 = v63
							}
							v79 = v66
						} else {
							v67 = int32(1)
							if v49&v67 != 0 {
								v79 = int32(base.Ui32(v49)>>(uint(v67)%32)) - v67
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if v20 == int32(1) {
							v82 = int32(4)
							v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
							if v84&int32(254) == int32(2) {
								v93 = v82
							} else {
								v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
							}
							if v84 == int32(1) {
								v96 = v82
							} else {
								v96 = v93
							}
							v109 = v96
						} else {
							v97 = int32(1)
							if v20&v97 != 0 {
								v109 = int32(base.Ui32(v20)>>(uint(v97)%32)) - v97
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
								v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if base.Ui32(v109) <= base.Ui32(v79) {
							F_text_position_setup(m, v12, v17, v19, v9)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								v119 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v119)
								v122 = F_text_position_next(m, v9)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int32(0)
								} else {
									if v122 == int32(0) {
										v135 = v119
										m.G0 = v9 + int32(1072)
										return v135
									} else {
										v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
										v127 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1052))
										v129 = F_pg_mbstrlen_with_len(m, v126, v127-v126)
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int32(0)
										} else {
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
											v135 = v129 + v131 + int32(1)
											m.G0 = v9 + int32(1072)
											return v135
										}
									}
								}
							}
						} else {
							v111 = F_pg_newlocale_from_collation(m, v19)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
								if v113 == int32(0) {
									F_text_position_setup(m, v12, v17, v19, v9)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										v119 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v119)
										v122 = F_text_position_next(m, v9)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int32(0)
										} else {
											if v122 == int32(0) {
												v135 = v119
												m.G0 = v9 + int32(1072)
												return v135
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1052))
												v129 = F_pg_mbstrlen_with_len(m, v126, v127-v126)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return int32(0)
												} else {
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
													v135 = v129 + v131 + int32(1)
													m.G0 = v9 + int32(1072)
													return v135
												}
											}
										}
									}
								} else {
									v135 = int32(0)
									m.G0 = v9 + int32(1072)
									return v135
								}
							}
						}
					} else {
						v46 = base.B2i32(v23 == int32(18)) << (uint(int32(4)) % 32)
						if v46 != 0 {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
							if v49 == int32(1) {
								v52 = int32(4)
								v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
								if v54&int32(254) == int32(2) {
									v63 = v52
								} else {
									v63 = base.B2i32(v54 == int32(18)) << (uint(v52) % 32)
								}
								if v54 == int32(1) {
									v66 = v52
								} else {
									v66 = v63
								}
								v79 = v66
							} else {
								v67 = int32(1)
								if v49&v67 != 0 {
									v79 = int32(base.Ui32(v49)>>(uint(v67)%32)) - v67
								} else {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							if v20 == int32(1) {
								v82 = int32(4)
								v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
								if v84&int32(254) == int32(2) {
									v93 = v82
								} else {
									v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
								}
								if v84 == int32(1) {
									v96 = v82
								} else {
									v96 = v93
								}
								v109 = v96
							} else {
								v97 = int32(1)
								if v20&v97 != 0 {
									v109 = int32(base.Ui32(v20)>>(uint(v97)%32)) - v97
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
									v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							if base.Ui32(v109) <= base.Ui32(v79) {
								F_text_position_setup(m, v12, v17, v19, v9)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									v119 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v119)
									v122 = F_text_position_next(m, v9)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return int32(0)
									} else {
										if v122 == int32(0) {
											v135 = v119
											m.G0 = v9 + int32(1072)
											return v135
										} else {
											v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1052))
											v129 = F_pg_mbstrlen_with_len(m, v126, v127-v126)
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int32(0)
											} else {
												v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
												v135 = v129 + v131 + int32(1)
												m.G0 = v9 + int32(1072)
												return v135
											}
										}
									}
								}
							} else {
								v111 = F_pg_newlocale_from_collation(m, v19)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
									if v113 == int32(0) {
										F_text_position_setup(m, v12, v17, v19, v9)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int32(0)
										} else {
											v119 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v119)
											v122 = F_text_position_next(m, v9)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												if v122 == int32(0) {
													v135 = v119
													m.G0 = v9 + int32(1072)
													return v135
												} else {
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
													v127 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1052))
													v129 = F_pg_mbstrlen_with_len(m, v126, v127-v126)
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
														v135 = v129 + v131 + int32(1)
														m.G0 = v9 + int32(1072)
														return v135
													}
												}
											}
										}
									} else {
										v135 = int32(0)
										m.G0 = v9 + int32(1072)
										return v135
									}
								}
							}
						} else {
							v135 = int32(1)
							m.G0 = v9 + int32(1072)
							return v135
						}
					}
				} else {
					v34 = int32(1)
					if v20&v34 != 0 {
						v46 = int32(base.Ui32(v20)>>(uint(v34)%32)) - v34
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
					}
					if v46 != 0 {
						v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
						if v49 == int32(1) {
							v52 = int32(4)
							v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
							if v54&int32(254) == int32(2) {
								v63 = v52
							} else {
								v63 = base.B2i32(v54 == int32(18)) << (uint(v52) % 32)
							}
							if v54 == int32(1) {
								v66 = v52
							} else {
								v66 = v63
							}
							v79 = v66
						} else {
							v67 = int32(1)
							if v49&v67 != 0 {
								v79 = int32(base.Ui32(v49)>>(uint(v67)%32)) - v67
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if v20 == int32(1) {
							v82 = int32(4)
							v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
							if v84&int32(254) == int32(2) {
								v93 = v82
							} else {
								v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
							}
							if v84 == int32(1) {
								v96 = v82
							} else {
								v96 = v93
							}
							v109 = v96
						} else {
							v97 = int32(1)
							if v20&v97 != 0 {
								v109 = int32(base.Ui32(v20)>>(uint(v97)%32)) - v97
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
								v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						if base.Ui32(v109) <= base.Ui32(v79) {
							F_text_position_setup(m, v12, v17, v19, v9)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								v119 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v119)
								v122 = F_text_position_next(m, v9)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int32(0)
								} else {
									if v122 == int32(0) {
										v135 = v119
										m.G0 = v9 + int32(1072)
										return v135
									} else {
										v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
										v127 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1052))
										v129 = F_pg_mbstrlen_with_len(m, v126, v127-v126)
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int32(0)
										} else {
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
											v135 = v129 + v131 + int32(1)
											m.G0 = v9 + int32(1072)
											return v135
										}
									}
								}
							}
						} else {
							v111 = F_pg_newlocale_from_collation(m, v19)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
								if v113 == int32(0) {
									F_text_position_setup(m, v12, v17, v19, v9)
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int32(0)
									} else {
										v119 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+5)) = uint8(v119)
										v122 = F_text_position_next(m, v9)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int32(0)
										} else {
											if v122 == int32(0) {
												v135 = v119
												m.G0 = v9 + int32(1072)
												return v135
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1064))
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1052))
												v129 = F_pg_mbstrlen_with_len(m, v126, v127-v126)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return int32(0)
												} else {
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
													v135 = v129 + v131 + int32(1)
													m.G0 = v9 + int32(1072)
													return v135
												}
											}
										}
									}
								} else {
									v135 = int32(0)
									m.G0 = v9 + int32(1072)
									return v135
								}
							}
						}
					} else {
						v135 = int32(1)
						m.G0 = v9 + int32(1072)
						return v135
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v144 = m.ExcPending
				if v144 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(34209924))
					mBase = m.M
					v147 = m.ExcPending
					if v147 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(257570), int32(0))
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(601086), int32(0))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524313), int32(1648), int32(113106))
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
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
				v30 = F__emscripten_memcpy_bulkmem(m, v23+int32(4), v16, v20)
				mBase = m.M
			} else {
			}
			F_pfree(m, v16)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
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
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
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
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
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
	var v497 int32
	_ = v497
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v543 int32
	_ = v543
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
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
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v732 int32
	_ = v732
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var __phi884 int32
	_ = __phi884
	var v885 int32
	_ = v885
	var __phi885 int32
	_ = __phi885
	var v886 int32
	_ = v886
	var __phi886 int32
	_ = __phi886
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v963 int64
	_ = v963
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1020 int32
	_ = v1020
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1078 int64
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int64
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1251 int32
	_ = v1251
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1317 int32
	_ = v1317
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1367 int32
	_ = v1367
	var v1372 int32
	_ = v1372
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(160)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = F_palloc0(m, int32(28))
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
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v543 != 0 {
		goto L162
	} else {
		goto L163
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v27 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L155
	}
L7:
	;
	if v533&int32(1) != 0 {
		goto L3
	} else {
		goto L154
	}
L8:
	;
	v533 = v2
	v543 = v2
	goto L7
L9:
	;
	goto L10
L10:
	;
	v32 = v2
	v42 = v2
	v44 = v2
	goto L11
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v44<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v52 = int32(404220)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, _consts[914])))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v56 == int32(0) {
		v75 = v55
		v76 = v56
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v533 = v513
	v543 = v523
	goto L7
L13:
	;
	v528 = v44 + int32(1)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v528 < v529 {
		v32 = v513
		v42 = v523
		v44 = v528
		goto L11
	} else {
		goto L153
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v497
	F_tsearch_readline_end(m, v19+int32(112))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L152
	}
L15:
	;
	if v76-v75 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	goto L15
L17:
	;
	if v55 != v56 {
		v75 = v55
		v76 = v56
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v60 = v51
	v61 = v52
	goto L19
L19:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	if v65 == int32(0) {
		v75 = v64
		v76 = v65
		goto L16
	} else {
		goto L21
	}
L20:
	;
	v75 = v64
	v76 = v65
	goto L16
L21:
	;
	v68 = int32(1)
	if v64 == v65 {
		v60 = v60 + v68
		v61 = v61 + v68
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v32&int32(1) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v420 = int32(17824)
	v423 = int32(*(*uint8)(unsafe.Add(mBase, _consts[915])))
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v424 == int32(0) {
		v443 = v423
		v444 = v424
		goto L129
	} else {
		goto L130
	}
L26:
	;
	v86 = F_defGetString(m, v50)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
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
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L124
	}
L29:
	;
	v89 = F_get_tsearch_config_filename(m, v86, int32(164562))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v91 = F_tsearch_readline_begin(m, v19+int32(112), v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v91 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v93 = int32(0)
	v97 = F_tsearch_readline(m, v19+int32(112))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L120
	}
L35:
	;
	if v97 == int32(0) {
		v497 = v93
		goto L14
	} else {
		goto L36
	}
L36:
	;
	v108 = v93
	v109 = v93
	v111 = v97
	goto L37
L37:
	;
	v117 = v111
	goto L39
L39:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if base.Ui32(v133-int32(9)) < base.Ui32(int32(5)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v383 = F_pg_mblen_cstr(m, v117)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L119
	}
L42:
	;
	v139 = int32(0)
	switch v133 - int32(32) {
	case 0:
		goto L41
	case 1, 2:
		goto L47
	case 3:
		v316 = v108
		v317 = v109
		goto L46
	default:
		goto L48
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L115
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L111
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L107
	}
L46:
	;
	F_pfree(m, v111)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L104
	}
L47:
	;
	v146 = v117
	v147 = v133
	v148 = int32(1)
	v149 = v139
	v151 = v139
	v153 = v108
	v155 = v139
	goto L50
L48:
	;
	if v133 == int32(0) {
		v316 = v108
		v317 = v109
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	switch v148 - int32(2) {
	case 0:
		goto L58
	case 1:
		goto L57
	case 2:
		goto L56
	default:
		goto L59
	}
L51:
	;
	if v276 == int32(4) {
		goto L95
	} else {
		goto L96
	}
L52:
	;
	v281 = F_pg_mblen_cstr(m, v146)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L93
	}
L53:
	;
	v270 = F_pg_mblen_cstr(m, v146)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L92
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L88
	}
L55:
	;
	v276 = int32(3)
	v277 = v251
	v278 = v151
	v279 = v153
	v280 = v252
	goto L52
L56:
	;
	v230 = v147 & int32(255)
	if base.Ui32(v230-int32(9)) < base.Ui32(int32(5)) {
		goto L83
	} else {
		goto L84
	}
L57:
	;
	v217 = v147 & int32(255)
	switch v217 - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		v276 = int32(3)
		v277 = v149
		v278 = v151
		v279 = v153
		v280 = v155
		goto L52
	case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 26, 27, 28, 29, 30, 31, 32:
		goto L78
	case 33:
		goto L80
	default:
		goto L79
	}
L58:
	;
	switch v147&int32(255) - int32(9) {
	case 0, 1, 2, 3, 4, 23:
		goto L74
	default:
		v276 = int32(2)
		v277 = v149
		v278 = v151
		v279 = v153
		v280 = v155
		goto L52
	case 49:
		goto L75
	}
L59:
	;
	v165 = v147 & int32(255)
	if v165 == int32(58) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v149 != 0 {
		v251 = v149
		v252 = v155
		goto L55
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v192 = base.B2i32(v165 != int32(32)) & base.B2i32(base.Ui32((v147-int32(14))&int32(255)) < base.Ui32(int32(251)))
	if v192 != 0 {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(226864), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(515997), int32(212), int32(485369))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	v193 = v146
	goto L70
L69:
	;
	v193 = v151
	goto L70
L70:
	;
	if v192 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v196 = int32(2)
	goto L73
L72:
	;
	v196 = int32(1)
	goto L73
L73:
	;
	v276 = v196
	v277 = v149
	v278 = v193
	v279 = v153
	v280 = v155
	goto L52
L74:
	;
	F_newLexeme(m, v23, v151, v146, v109, v149&int32(65535))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	F_newLexeme(m, v23, v151, v146, v109, v149&int32(65535))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v251 = v149 + int32(1)
	v252 = v155
	goto L55
L77:
	;
	v212 = int32(1)
	v276 = v212
	v277 = v149 + v212
	v278 = v151
	v279 = v153
	v280 = v155
	goto L52
L78:
	;
	v276 = int32(4)
	v277 = v149
	v278 = v146
	v279 = int32(0)
	v280 = v155
	goto L52
L79:
	;
	if v217 == int32(92) {
		goto L53
	} else {
		goto L82
	}
L80:
	;
	v220 = F_pg_mblen_cstr(m, v146)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v276 = int32(4)
	v277 = v149
	v278 = v220 + v146
	v279 = int32(1)
	v280 = v155
	goto L52
L82:
	;
	goto L78
L83:
	;
	if v146 == v151 {
		goto L54
	} else {
		goto L86
	}
L84:
	;
	if v230 == int32(32) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v276 = int32(4)
	v277 = v149
	v278 = v151
	v279 = v153
	v280 = v155
	goto L52
L86:
	;
	v239 = int32(65535)
	F_addWrd(m, v23, v151, v146, v109, v155&v239, v149&v239, v153&int32(1))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v251 = v149
	v252 = v155 + int32(1)
	goto L55
L88:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(393910), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(515997), int32(262), int32(485369))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v276 = int32(4)
	v277 = v149
	v278 = v270 + v146
	v279 = int32(0)
	v280 = v155
	goto L52
L93:
	;
	v283 = v281 + v146
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	if v284 != 0 {
		v146 = v283
		v147 = v284
		v148 = v276
		v149 = v277
		v151 = v278
		v153 = v279
		v155 = v280
		goto L50
	} else {
		goto L94
	}
L94:
	;
	goto L51
L95:
	;
	if v283 == v278 {
		goto L45
	} else {
		goto L98
	}
L96:
	;
	v298 = v280
	goto L97
L97:
	;
	if v298 == int32(0) {
		goto L44
	} else {
		goto L100
	}
L98:
	;
	v288 = int32(65535)
	F_addWrd(m, v23, v278, v283, v109, v280&v288, v277&v288, v279&int32(1))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v298 = v280 + int32(1)
	goto L97
L100:
	;
	if v277 == int32(0) {
		goto L44
	} else {
		goto L101
	}
L101:
	;
	if base.Ui32(int32(65535)) < base.Ui32(v298) {
		goto L43
	} else {
		goto L102
	}
L102:
	;
	if base.Ui32(int32(65536)) <= base.Ui32(v277) {
		goto L43
	} else {
		goto L103
	}
L103:
	;
	v316 = v279
	v317 = v109 + int32(1)
	goto L46
L104:
	;
	v329 = F_tsearch_readline(m, v19+int32(112))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v329 != 0 {
		v108 = v316
		v109 = v317
		v111 = v329
		goto L37
	} else {
		goto L106
	}
L106:
	;
	v497 = v317
	goto L14
L107:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errmsg(m, int32(393910), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(515997), int32(278), int32(485369))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
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
	F_errcode(m, int32(22))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(391185), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(515997), int32(287), int32(485369))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(12385), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(515997), int32(292), int32(485369))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	v117 = v383 + v117
	goto L39
L120:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v89
	F_errmsg(m, int32(311458), v19+int32(80))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(515997), int32(180), int32(485369))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(140864), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(515997), int32(616), int32(106439))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	if v444-v443 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L129:
	;
	goto L128
L130:
	;
	if v423 != v424 {
		v443 = v423
		v444 = v424
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v428 = v51
	v429 = v420
	goto L132
L132:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429)+1)))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v428)+1)))
	if v433 == int32(0) {
		v443 = v432
		v444 = v433
		goto L129
	} else {
		goto L134
	}
L133:
	;
	v443 = v432
	v444 = v433
	goto L129
L134:
	;
	v436 = int32(1)
	if v432 == v433 {
		v428 = v428 + v436
		v429 = v429 + v436
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	if v42 == int32(0) {
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
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L148
	}
L139:
	;
	v450 = F_defGetString(m, v50)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	v452 = F_pstrdup(m, v450)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v513 = v32
	v523 = v452
	goto L13
L144:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(140649), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(515997), int32(625), int32(106439))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v477
	F_errmsg(m, int32(758058), v19+int32(96))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(515997), int32(633), int32(106439))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	v513 = int32(1)
	v523 = v42
	goto L13
L153:
	;
	goto L12
L154:
	;
	goto L6
L155:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errmsg(m, int32(227498), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(515997), int32(640), int32(106439))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L323
	}
L160:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L1
	} else {
		goto L318
	}
L161:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L1
	} else {
		goto L314
	}
L162:
	;
	v581 = int32(0)
	v583 = F_stringToQualifiedNameList(m, v543, v581)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L1
	} else {
		goto L310
	}
L165:
	;
	v586 = F_get_ts_dict_oid(m, v583, int32(0))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v586
	v589 = F_lookup_ts_dictionary_cache(m, v586)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v589
	v593 = F_palloc(m, int32(128))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v595 <= int32(0) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v866 != 0 {
		goto L216
	} else {
		goto L217
	}
L170:
	;
	v852 = v581
	v857 = v593
	v858 = int32(16)
	goto L169
L171:
	;
	goto L172
L172:
	;
	v602 = v581
	v607 = v593
	v608 = int32(16)
	v615 = v2
	goto L173
L173:
	;
	v617 = v615 << (uint(int32(3)) % 32)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v619 = v617 + v618
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620))))
	if v621 != int32(63) {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v852 = v822
	v857 = v827
	v858 = v828
	goto L169
L175:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v836+v617)))
	F_pfree(m, v838)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L213
	}
L176:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v662 = int32(0)
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v659)+44))
	v664 = F_strlen(m, v620)
	mBase = m.M
	v666 = F_FunctionCall4Coll(m, v659+int32(12), v662, v663, v620, v664, v662)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L184
	}
L177:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620)+1)))
	if v624 != 0 {
		goto L176
	} else {
		goto L178
	}
L178:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v619)+4))
	if v608 <= v602 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v629 = F_repalloc(m, v607, v608<<(uint(int32(4))%32))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	v633 = v607
	v634 = v608
	goto L181
L181:
	;
	v636 = F_palloc(m, int32(16))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L183
	}
L182:
	;
	v633 = v629
	v634 = v608 << (uint(int32(1)) % 32)
	goto L181
L183:
	;
	v640 = v633 + v602<<(uint(int32(3))%32)
	v641 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v640))) = v641
	v644 = v640 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = v636
	v646 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v636)+6)) = uint16(v646)
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v625)))
	*(*int32)(unsafe.Add(mBase, uint32(v648))) = v649
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	v652 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v625)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v651)+4)) = uint16(v652)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v644)))
	*(*int32)(unsafe.Add(mBase, uint32(v654)+8)) = v641
	v822 = v602 + v646
	v827 = v633
	v828 = v634
	goto L175
L184:
	;
	if v666 == int32(0) {
		goto L161
	} else {
		goto L185
	}
L185:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v666)+4))
	if v670 == int32(0) {
		goto L160
	} else {
		goto L186
	}
L186:
	;
	v675 = v602
	v678 = v666
	v680 = v607
	v681 = v608
	goto L187
L187:
	;
	v689 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v678))))
	v690 = int32(1)
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v678)+12))
	if v691 == int32(0) {
		v732 = v690
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v822 = v811
	v827 = v814
	v828 = v815
	goto L175
L189:
	;
	v742 = v678
	v743 = v675
	v744 = v678 + int32(4)
	v748 = v680
	v749 = v681
	goto L195
L190:
	;
	v694 = v678
	v705 = v690
	goto L191
L191:
	;
	v711 = v694 + int32(8)
	v712 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v711))))
	if v712 != v689&int32(65535) {
		v732 = v705
		goto L189
	} else {
		goto L193
	}
L192:
	;
	v732 = v717
	goto L189
L193:
	;
	v717 = v705 + int32(1)
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v694+int32(20))))
	if v720 != 0 {
		v694 = v711
		v705 = v717
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	v757 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v742))))
	if v689&int32(65535) != v757 {
		goto L198
	} else {
		goto L199
	}
L196:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v813)+4))
	if v819 != 0 {
		v675 = v811
		v678 = v813
		v680 = v814
		v681 = v815
		goto L187
	} else {
		goto L212
	}
L197:
	;
	goto L196
L198:
	;
	v811 = v743
	v813 = v742
	v814 = v748
	v815 = v749
	goto L197
L199:
	;
	goto L200
L200:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v759+v617)+4))
	if v749 <= v743 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v765 = F_repalloc(m, v748, v749<<(uint(int32(4))%32))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L204
	}
L202:
	;
	v769 = v748
	v770 = v749
	goto L203
L203:
	;
	v773 = v769 + v743<<(uint(int32(3))%32)
	v775 = v773 + int32(4)
	v777 = F_palloc(m, int32(16))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L205
	}
L204:
	;
	v769 = v765
	v770 = v749 << (uint(int32(1)) % 32)
	goto L203
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v775))) = v777
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v744)))
	if v780 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v773))) = v789
	*(*uint16)(unsafe.Add(mBase, uint32(v788)+6)) = uint16(v790)
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v761)))
	*(*int32)(unsafe.Add(mBase, uint32(v793))) = v794
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	v797 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v796)+4)) = uint16(v797)
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	*(*int32)(unsafe.Add(mBase, uint32(v799)+8)) = int32(0)
	v803 = v743 + int32(1)
	v805 = v742 + int32(12)
	v807 = v742 + int32(8)
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v805)))
	if v808 != 0 {
		v742 = v807
		v743 = v803
		v744 = v805
		v748 = v769
		v749 = v770
		goto L195
	} else {
		goto L211
	}
L207:
	;
	v788 = v777
	v789 = int32(0)
	v790 = int32(1)
	goto L206
L208:
	;
	goto L209
L209:
	;
	v785 = F_pstrdup(m, v780)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	v788 = v787
	v789 = v785
	v790 = v732
	goto L206
L211:
	;
	v811 = v803
	v813 = v807
	v814 = v769
	v815 = v770
	goto L197
L212:
	;
	goto L188
L213:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v841+v617)+4))
	F_pfree(m, v843)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v847 = v615 + int32(1)
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v847 < v848 {
		v602 = v822
		v607 = v827
		v608 = v828
		v615 = v847
		goto L173
	} else {
		goto L215
	}
L215:
	;
	goto L174
L216:
	;
	F_pfree(m, v866)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v858
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v852
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v857
	if int32(2) <= v852 {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	goto L218
L220:
	;
	F_pg_qsort(m, v857, v852, int32(8), int32(1165))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L1
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	if int32(0) < v1020 {
		goto L260
	} else {
		goto L261
	}
L223:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v879 < int32(2) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v994 = int32(3)
	v997 = (v978-v980)>>(uint(v994)%32) + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v997
	v1001 = F_repalloc(m, v980, v997<<(uint(v994)%32))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L257
	}
L225:
	;
	v978 = v878
	v980 = v878
	goto L224
L226:
	;
	goto L227
L227:
	;
	__phi884 = v878
	__phi885 = v878
	__phi886 = v878 + int32(8)
	v884 = __phi884
	v885 = __phi885
	v886 = __phi886
	goto L228
L228:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v885)))
	v902 = v884 + int32(8)
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v902)))
	if v903 != 0 {
		goto L232
	} else {
		goto L233
	}
L229:
	;
	v978 = v965
	v980 = v972
	goto L224
L230:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v971 = v886 + int32(8)
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if (v971-v972)>>(uint(int32(3))%32) < v969 {
		__phi884 = v886
		__phi885 = v965
		__phi886 = v971
		v884 = __phi884
		v885 = __phi885
		v886 = __phi886
		goto L228
	} else {
		goto L256
	}
L231:
	;
	v962 = v885 + int32(8)
	v963 = *(*int64)(unsafe.Add(mBase, uint32(v902)))
	*(*int64)(unsafe.Add(mBase, uint32(v962))) = v963
	v965 = v962
	goto L230
L232:
	;
	if v900 == int32(0) {
		goto L231
	} else {
		goto L235
	}
L233:
	;
	v931 = v900
	goto L234
L234:
	;
	if v931 != 0 {
		goto L231
	} else {
		goto L244
	}
L235:
	;
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v900))))
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v903))))
	if v909 == int32(0) {
		v928 = v908
		v929 = v909
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v931 = v929 - v928
	goto L234
L237:
	;
	goto L236
L238:
	;
	if v908 != v909 {
		v928 = v908
		v929 = v909
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v913 = v903
	v914 = v900
	goto L240
L240:
	;
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914)+1)))
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v913)+1)))
	if v918 == int32(0) {
		v928 = v917
		v929 = v918
		goto L237
	} else {
		goto L242
	}
L241:
	;
	v928 = v917
	v929 = v918
	goto L237
L242:
	;
	v921 = int32(1)
	if v917 == v918 {
		v913 = v913 + v921
		v914 = v914 + v921
		goto L240
	} else {
		goto L243
	}
L243:
	;
	goto L241
L244:
	;
	v933 = v884 + int32(12)
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v933)))
	if v934 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v902)))
	if v956 == int32(0) {
		v965 = v885
		goto L230
	} else {
		goto L254
	}
L246:
	;
	F_pfree(m, v934)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L253
	}
L247:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v885)+4))
	if v937 == int32(0) {
		goto L246
	} else {
		goto L248
	}
L248:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v934)))
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v937)))
	if v940 != v941 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v934)+8)) = v937
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v933)))
	*(*int32)(unsafe.Add(mBase, uint32(v885)+4)) = v950
	goto L245
L250:
	;
	v943 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+4)))
	v944 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937)+4)))
	if v943 != v944 {
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v946 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+6)))
	v947 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937)+6)))
	if v946 == v947 {
		goto L246
	} else {
		goto L252
	}
L252:
	;
	goto L249
L253:
	;
	goto L245
L254:
	;
	F_pfree(m, v956)
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v965 = v885
	goto L230
L256:
	;
	goto L229
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v1001
	goto L222
L258:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L1
	} else {
		goto L306
	}
L259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L302
	}
L260:
	;
	v1037 = int32(0)
	goto L263
L261:
	;
	goto L262
L262:
	;
	m.G0 = v19 + int32(160)
	return v23
L263:
	;
	v1041 = v1037 << (uint(int32(3)) % 32)
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1041+v1042)+4))
	v1046 = F_palloc(m, int32(16))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L265
	}
L264:
	;
	goto L262
L265:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1048+v1041)+4)) = v1046
	v1051 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1046)+4)) = v1051
	if v1044 == v1051 {
		v1185 = v1046
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1202 = v1201 + v1041
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+4))
	if v1185 == v1203 {
		goto L159
	} else {
		goto L299
	}
L267:
	;
	v1059 = int32(2)
	v1063 = v1044
	v1066 = v1046
	goto L268
L268:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+4))
	if v1072 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	v1185 = v1156
	goto L266
L270:
	;
	v1185 = v1066
	goto L266
L271:
	;
	goto L272
L272:
	;
	v1075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1063)+3)))
	if v1075&int32(16) != 0 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1097)+4))
	if v1098 == int32(0) {
		goto L259
	} else {
		goto L279
	}
L274:
	;
	v1078 = *(*int64)(unsafe.Add(mBase, uint32(v1063)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v1078
	v1080 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+124)) = v1080
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+114)) = uint16(v1080)
	v1097 = v19 + int32(112)
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v1089 = int32(0)
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1086)+44))
	v1091 = F_strlen(m, v1072)
	mBase = m.M
	v1093 = F_FunctionCall4Coll(m, v1086+int32(12), v1089, v1090, v1072, v1091, v1089)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	if v1093 == int32(0) {
		goto L258
	} else {
		goto L278
	}
L278:
	;
	v1097 = v1093
	goto L273
L279:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1103+v1041)+4))
	v1109 = v1066
	v1110 = v1097
	v1111 = v1097 + int32(4)
	v1112 = v1059
	goto L280
L280:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1125+v1041)+4))
	v1128 = v1109 - v1127
	if v1112 <= v1128>>(uint(int32(3))%32)+int32(1) {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	if v1066 == v1105 {
		goto L288
	} else {
		goto L289
	}
L282:
	;
	v1136 = F_repalloc(m, v1127, v1112<<(uint(int32(4))%32))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L1
	} else {
		goto L285
	}
L283:
	;
	v1147 = v1109
	v1148 = v1112
	goto L284
L284:
	;
	v1149 = *(*int64)(unsafe.Add(mBase, uint32(v1110)))
	*(*int64)(unsafe.Add(mBase, uint32(v1147))) = v1149
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1111)))
	v1152 = F_pstrdup(m, v1151)
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L1
	} else {
		goto L286
	}
L285:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1138+v1041)+4)) = v1136
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1143+v1041)+4))
	v1147 = v1145 + v1128
	v1148 = v1112 << (uint(int32(1)) % 32)
	goto L284
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1147)+4)) = v1152
	v1155 = int32(8)
	v1156 = v1147 + v1155
	v1158 = v1110 + int32(12)
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1158)))
	if v1161 != 0 {
		v1109 = v1156
		v1110 = v1110 + v1155
		v1111 = v1158
		v1112 = v1148
		goto L280
	} else {
		goto L287
	}
L287:
	;
	goto L281
L288:
	;
	v1164 = int32(-1)
	goto L290
L289:
	;
	v1164 = (v1066 - v1105) >> (uint(int32(3)) % 32)
	goto L290
L290:
	;
	if int32(0) < v1164 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1167+v1041)+4))
	v1174 = v1169 + v1164<<(uint(int32(3))%32) + int32(2)
	v1175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1174))))
	v1177 = v1175 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1174))) = uint16(v1177)
	goto L293
L292:
	;
	goto L293
L293:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+4))
	if v1180 != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	F_pfree(m, v1180)
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L1
	} else {
		goto L297
	}
L295:
	;
	goto L296
L296:
	;
	v1184 = v1063 + int32(8)
	if v1184 != 0 {
		v1059 = v1148
		v1063 = v1184
		v1066 = v1156
		goto L268
	} else {
		goto L298
	}
L297:
	;
	goto L296
L298:
	;
	goto L269
L299:
	;
	v1207 = int32(base.Ui32(v1185-v1203) >> (uint(int32(3)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v1202)+2)) = uint16(v1207)
	F_pfree(m, v1044)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	v1212 = v1037 + int32(1)
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	if v1212 < v1213 {
		v1037 = v1212
		goto L263
	} else {
		goto L301
	}
L301:
	;
	goto L264
L302:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v1037 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v1242
	F_errmsg(m, int32(710913), v19+int32(32))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(515997), int32(568), int32(365369))
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L306:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v1037 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v1264
	F_errmsg(m, int32(710767), v19+int32(16))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	F_errfinish(m, int32(515997), int32(575), int32(365369))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L310:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_errmsg(m, int32(227213), int32(0))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	F_errfinish(m, int32(515997), int32(644), int32(106439))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L314:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1305 = v1302 + v615<<(uint(int32(3))%32)
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+4))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1306)))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1305)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v1308
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v1307 + int32(1)
	F_errmsg(m, int32(710842), v19+int32(48))
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	F_errfinish(m, int32(515997), int32(418), int32(393943))
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L318:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1333 = v1330 + v615<<(uint(int32(3))%32)
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+4))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1334)))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1333)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v1336
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v1335 + int32(1)
	F_errmsg(m, int32(710969), v19-int32(-64))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	F_errhint(m, int32(660170), int32(0))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	F_errfinish(m, int32(515997), int32(425), int32(393943))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L323:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1037 + int32(1)
	F_errmsg(m, int32(710720), v19)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	F_errfinish(m, int32(515997), int32(587), int32(365369))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L1
	} else {
		goto L326
	}
L326:
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	v22 = F_pg_snprintf(m, v7+v14, int32(32), int32(700295), v7)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v28 = F_pstrdup(m, v7+int32(16))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(48)
			return v28
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
			v23 = int32(24)
			v25 = int32(65280)
			v27 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v20+v21))) = v16<<(uint(v23)%32) | v16&v25<<(uint(v27)%32) | (v15<<(uint(v27)%32)&v25 | int32(base.Ui32(v15<<(uint(int32(16))%32))>>(uint(v23)%32)))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v20 + int32(4)
			v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
			F_enlargeStringInfo(m, v8, int32(2))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v51 = int32(8)
				v55 = v44<<(uint(v51)%32) | int32(base.Ui32(v44)>>(uint(v51)%32))
				*(*uint16)(unsafe.Add(mBase, uint32(v48+v49))) = uint16(v55)
				v57 = int32(2)
				v58 = v48 + v57
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v58
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				*(*int32)(unsafe.Add(mBase, uint32(v61))) = v58 << (uint(v57) % 32)
				m.G0 = v8 + int32(16)
				return v61
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+2)))
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4))))
	v10 = int32(16)
	v12 = v8 | v9<<(uint(v10)%32)
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v17 = v13 | v14<<(uint(v10)%32)
	if base.Ui32(v12) < base.Ui32(v17) {
		v28 = int32(-1)
	} else {
		if base.Ui32(v17) < base.Ui32(v12) {
			v28 = int32(1)
		} else {
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v22) < base.Ui32(v23) {
				v28 = int32(-1)
			} else {
				v28 = base.B2i32(base.Ui32(v23) < base.Ui32(v22))
			}
		}
	}
	if v28 <= int32(0) {
		v31 = v4
	} else {
		v31 = v3
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
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
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
	var v94 int64
	_ = v94
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v187 int64
	_ = v187
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v194 int64
	_ = v194
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v211 int32
	_ = v211
	var v212 int64
	_ = v212
	var v216 int64
	_ = v216
	var v219 int64
	_ = v219
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int64
	_ = v379
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
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
	v58 = base.I64_div_s(v54, v57)
	v61 = v54 - v58*v57
	__phi66 = int32(1970)
	__phi71 = v58
	v66 = __phi66
	v71 = __phi71
	goto L13
L5:
	;
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v54 = v50
	v55 = int64(0)
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
	if v37 != v33 {
		v54 = v33
		v55 = v39
		v56 = int32(0)
		goto L4
	} else {
		goto L10
	}
L8:
	;
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v36 = l2 + int32(22632) + v30<<(uint(int32(4))%32)
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
		v54 = v33
		v55 = v39
		v56 = base.B2i32(int64(0) < v39)
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v36-int32(8))))
	v54 = v33
	v55 = v39
	v56 = base.B2i32(v48 < v39)
	goto L4
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(61)
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
	v181 = base.I32_wrap_i64(v71)
	v182 = base.I64_extend_i32_s(l1)
	v184 = v182 - v55 + v61
	if v184 < int64(0) {
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
	v94 = int64(*(*int32)(unsafe.Add(mBase, uint32(v89<<(uint(int32(2))%32))+uint32(_consts[1236]))))
	if v71 < v94 {
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
	v102 = int32(-1)
	goto L26
L25:
	;
	v102 = int32(1)
	goto L26
L26:
	;
	v104 = base.I64_div_s(v71, int64(366))
	if base.Ui64(v71+int64(365)) < base.Ui64(int64(731)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v110 = v102
	goto L29
L28:
	;
	v110 = base.I32_wrap_i64(v104)
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
	v119 = v110 + v66
	v121 = v119 - int32(1)
	if v121 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L31:
	;
	if v110 <= v66^int32(2147483647) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v110 < int32(-2147483648)-v66 {
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
	v153 = v66 - int32(1)
	if v153 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v125 = int32(0) - v119
	v129 = base.I32_div_u_s(v125, int32(100))
	v132 = base.I32_div_u_s(v125, int32(400))
	v145 = int32(base.Ui32(v125)>>(uint(int32(2))%32)) - v129 + v132 ^ int32(-1)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v139 = base.I32_div_u_s(v121, int32(100))
	v142 = base.I32_div_u_s(v121, int32(400))
	v145 = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v139 + v142
	goto L36
L40:
	;
	__phi66 = v119
	__phi71 = (base.I64_extend_i32_s(v119)-base.I64_extend_i32_s(v66))*int64(-365) + v71 - base.I64_extend_i32_s(v145-v177)
	v66 = __phi66
	v71 = __phi71
	goto L13
L41:
	;
	v157 = int32(0) - v66
	v161 = base.I32_div_u_s(v157, int32(100))
	v164 = base.I32_div_u_s(v157, int32(400))
	v177 = int32(base.Ui32(v157)>>(uint(int32(2))%32)) - v161 + v164 ^ int32(-1)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v171 = base.I32_div_u_s(v153, int32(100))
	v174 = base.I32_div_u_s(v153, int32(400))
	v177 = int32(base.Ui32(v153)>>(uint(int32(2))%32)) - v171 + v174
	goto L40
L44:
	;
	v187 = int64(-86400)
	if base.Ui64(v184) <= base.Ui64(v187) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v211 = v181
	v212 = v184
	goto L46
L46:
	;
	if int64(86400) <= v212 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v190 = v187
	goto L49
L48:
	;
	v190 = v184
	goto L49
L49:
	;
	v192 = v55 + v190 - v61
	v194 = base.I64_extend_i32_u(base.B2i32(v192 != v182))
	v197 = int64(86400)
	v198 = base.I64_div_u_s(v192-(v194+v182), v197)
	v199 = v198 + v194
	v211 = base.I32_wrap_i64(v199) ^ int32(-1) + v181
	v212 = v61 + v199*v197 + v182 - v55 + v197
	goto L46
L50:
	;
	v216 = v212 - int64(172799)
	if base.Ui64(v216) <= base.Ui64(v212) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v233 = v211
	v234 = v212
	goto L52
L52:
	;
	if v233 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v219 = v216
	goto L55
L54:
	;
	v219 = int64(0)
	goto L55
L55:
	;
	v222 = int64(86400)
	v223 = base.I64_div_u_s(v219+int64(86399), v222)
	v233 = v211 + base.I32_wrap_i64(v223) + int32(1)
	v234 = v212 + v223*int64(-86400) - v222
	goto L52
L56:
	;
	v238 = v233
	v241 = v66
	goto L59
L57:
	;
	v273 = v233
	v276 = v66
	goto L58
L58:
	;
	v285 = v273
	v288 = v276
	goto L66
L59:
	;
	if v241 == int32(-2147483648) {
		goto L12
	} else {
		goto L61
	}
L60:
	;
	v273 = v270
	v276 = v254
	goto L58
L61:
	;
	v254 = v241 - int32(1)
	if v254&int32(3) != 0 {
		v264 = int32(0)
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v264<<(uint(int32(2))%32))+uint32(_consts[1236])))
	v270 = v269 + v238
	if v270 < int32(0) {
		v238 = v270
		v241 = v254
		goto L59
	} else {
		goto L65
	}
L63:
	;
	v259 = base.I32_rem_s(v254, int32(100))
	if v259 != 0 {
		v264 = int32(1)
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v261 = base.I32_rem_s(v254, int32(400))
	v264 = base.B2i32(v261 == int32(0))
	goto L62
L65:
	;
	goto L60
L66:
	;
	v298 = v288 & int32(3)
	if v298 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1237])) = v288
	if v288 < int32(-2147481748) {
		goto L12
	} else {
		goto L80
	}
L68:
	;
	goto L67
L69:
	;
	if v288 == int32(2147483647) {
		goto L12
	} else {
		goto L79
	}
L70:
	;
	v302 = base.I32_rem_s(v288, int32(100))
	if v302 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if v285 < int32(365) {
		goto L68
	} else {
		goto L78
	}
L73:
	;
	v306 = base.I32_rem_s(v288, int32(400))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(base.B2i32(v306 == int32(0))<<(uint(int32(2))%32))+uint32(_consts[1236])))
	if v285 < v313 {
		goto L68
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if v285 < int32(366) {
		goto L68
	} else {
		goto L77
	}
L76:
	;
	v316 = base.I32_rem_s(v288, int32(400))
	v325 = base.B2i32(v316 == int32(0))
	goto L69
L77:
	;
	v325 = int32(1)
	goto L69
L78:
	;
	v325 = int32(0)
	goto L69
L79:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v325<<(uint(int32(2))%32))+uint32(_consts[1236])))
	v285 = v285 - v334
	v288 = v288 + int32(1)
	goto L66
L80:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1238])) = v285
	*(*int32)(unsafe.Add(mBase, _consts[1237])) = v288 - int32(1900)
	v349 = base.I32_rem_s(v288-int32(1970), int32(7))
	if v288 <= int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v376 = int32(0)
	v379 = base.I64_div_u_s(v234, int64(3600))
	*(*uint32)(unsafe.Add(mBase, _consts[1239])) = uint32(v379)
	v386 = int32(7)
	v387 = base.I32_rem_s(v285+v349+v375-int32(473), v386)
	if v387 < v376 {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v353 = int32(0) - v288
	v357 = base.I32_div_u_s(v353, int32(100))
	v360 = base.I32_div_u_s(v353, int32(400))
	v375 = int32(base.Ui32(v353)>>(uint(int32(2))%32)) - v357 + v360 ^ int32(-1)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v365 = v288 - int32(1)
	v369 = base.I32_div_u_s(v365, int32(100))
	v372 = base.I32_div_u_s(v365, int32(400))
	v375 = int32(base.Ui32(v365)>>(uint(int32(2))%32)) - v369 + v372
	goto L81
L85:
	;
	v392 = v387 + v386
	goto L87
L86:
	;
	v392 = v387
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1240])) = v392
	v398 = base.I32_wrap_i64(v234 - v379*int64(3600))
	v399 = int32(65535)
	v401 = int32(60)
	v402 = base.I32_div_u_s(v398&v399, v401)
	*(*int32)(unsafe.Add(mBase, _consts[1241])) = v402
	*(*int32)(unsafe.Add(mBase, _consts[1242])) = v56 + (v398-v402*v401)&v399
	if v298 != 0 {
		v420 = int32(0)
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v422 = v420 * int32(48)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v422)+uint32(_consts[1243])))
	if v425 <= v285 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v415 = base.I32_rem_s(v288, int32(100))
	if v415 != 0 {
		v420 = int32(1)
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v417 = base.I32_rem_s(v288, int32(400))
	v420 = base.B2i32(v417 == int32(0))
	goto L88
L91:
	;
	v427 = v285
	v430 = v425
	v431 = v376
	goto L94
L92:
	;
	v447 = v285
	v451 = v376
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1244])) = l1
	*(*int32)(unsafe.Add(mBase, _consts[1245])) = v451
	*(*int32)(unsafe.Add(mBase, _consts[1246])) = v447 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1247])) = int32(0)
	return int32(4554472)
L94:
	;
	v439 = v427 - v430
	v441 = v431 + int32(1)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v422+int32(1864608)+v441<<(uint(int32(2))%32))))
	if v445 <= v439 {
		v427 = v439
		v430 = v445
		v431 = v441
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v447 = v439
	v451 = v441
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
	F_errmsg(m, int32(13120), v11)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(522343), int32(590), int32(95444))
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
func F_tokenize_auth_file(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
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
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v376 int32
	_ = v376
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v449 int32
	_ = v449
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v522 int32
	_ = v522
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
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
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v611 int32
	_ = v611
	var v619 int32
	_ = v619
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v726 int32
	_ = v726
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v753 int32
	_ = v753
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v780 int32
	_ = v780
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v812 int32
	_ = v812
	var v828 int32
	_ = v828
	var v839 int32
	_ = v839
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1004 int32
	_ = v1004
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1093 int32
	_ = v1093
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1249 int32
	_ = v1249
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	v24 = m.G0
	v26 = v24 - int32(80)
	m.G0 = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(799)
	v30 = int32(4547032)
	v31 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v26 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v26 + int32(24)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v49 = F_AllocSetContextCreateInternal(m, v44, int32(404629), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v51 = int32(4554128)
	v52 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v49
	F_initStringInfo(m, v26+int32(44))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
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
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v63 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v1266
	F_MemoryContextDelete(m, v1263)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L1
	} else {
		goto L237
	}
L8:
	;
	if int32(base.Ui32(v68)>>(uint(int32(4))%32))&int32(1) != 0 {
		v1249 = v26
		v1263 = v49
		v1266 = v52
		goto L7
	} else {
		goto L13
	}
L9:
	;
	goto L8
L10:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v68 = v66
	goto L9
L11:
	;
	goto L12
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v68 = v67
	goto L9
L13:
	;
	v73 = int32(1)
	v76 = l0
	v77 = l1
	v78 = l2
	v79 = l3
	v81 = v26
	v90 = l4 + v73
	v92 = v73
	v95 = v49
	v98 = v52
	goto L14
L14:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v77)+76))
	if v99 < int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v1249 = v883
	v1263 = v897
	v1266 = v900
	goto L7
L16:
	;
	if int32(base.Ui32(v104)>>(uint(int32(5))%32))&int32(1) != 0 {
		v1249 = v81
		v1263 = v95
		v1266 = v98
		goto L7
	} else {
		goto L21
	}
L17:
	;
	goto L16
L18:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v104 = v102
	goto L17
L19:
	;
	goto L20
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v104 = v103
	goto L17
L21:
	;
	v109 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+20)) = v109
	v113 = v81 + int32(44)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v109
	goto L22
L22:
	;
	v121 = int32(0)
	v124 = F_pg_get_line_append(m, v77, v81+int32(44))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v77)+76))
	if v265 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L24:
	;
	if v124 == int32(0) {
		v259 = v121
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v134 = v109
	v145 = v121
	goto L26
L26:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v81)+44))
	v152 = F_strlen(m, v151)
	mBase = m.M
	if v152 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v259 = v236
	goto L23
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+48)) = v219
	if v219 <= v134 {
		v259 = v145
		goto L23
	} else {
		goto L37
	}
L29:
	;
	v219 = v152
	goto L28
L30:
	;
	goto L31
L31:
	;
	v161 = v152
	goto L32
L32:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+(v151-int32(1))))))
	switch v181 - int32(10) {
	case 0, 3:
		goto L35
	default:
		v193 = v161
		goto L34
	}
L33:
	;
	v219 = v193
	goto L28
L34:
	;
	goto L33
L35:
	;
	v184 = int32(0)
	v185 = int32(1)
	v186 = v161 - v185
	*(*uint8)(unsafe.Add(mBase, uint32(v151+v186))) = uint8(v184)
	if v185 < v161 {
		v161 = v186
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v193 = v184
	goto L34
L37:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v81)+44))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222+v219-int32(1)))))
	if v226 != int32(92) {
		v259 = v145
		goto L23
	} else {
		goto L38
	}
L38:
	;
	v229 = int32(1)
	v230 = v219 - v229
	*(*int32)(unsafe.Add(mBase, uint32(v81)+48)) = v230
	v233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v230+v222))) = uint8(v233)
	v236 = v145 + v229
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v81)+48))
	v240 = F_pg_get_line_append(m, v77, v81+int32(44))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v240 != 0 {
		v134 = v237
		v145 = v236
		goto L26
	} else {
		goto L40
	}
L40:
	;
	goto L27
L41:
	;
	if int32(base.Ui32(v270)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	goto L41
L43:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v270 = v268
	goto L42
L44:
	;
	goto L45
L45:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v270 = v269
	goto L42
L46:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v278 = F_errstart(m, v79, int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	v300 = int32(0)
	v301 = base.B2i32(v299 == v300)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v81)+44))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	if v304 == v300 {
		v878 = v76
		v879 = v77
		v880 = v78
		v881 = v79
		v883 = v81
		v884 = v301
		v891 = v300
		v892 = v90
		v894 = v92
		v895 = v259
		v897 = v95
		v900 = v98
		goto L57
	} else {
		goto L58
	}
L49:
	;
	if v278 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v76
	v297 = F_psprintf(m, int32(313155), v81)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L56
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v76
	F_errmsg(m, int32(313155), v81+int32(16))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(524345), int32(769), int32(404629))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	v1249 = v81
	v1263 = v95
	v1266 = v98
	goto L7
L57:
	;
	if v884 != 0 {
		goto L155
	} else {
		goto L156
	}
L58:
	;
	if v299 != 0 {
		v878 = v76
		v879 = v77
		v880 = v78
		v881 = v79
		v883 = v81
		v884 = v301
		v891 = v300
		v892 = v90
		v894 = v92
		v895 = v259
		v897 = v95
		v900 = v98
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v307 = v76
	v308 = v77
	v309 = v78
	v310 = v79
	v312 = v81
	v316 = v303
	v320 = v300
	v321 = v90
	v323 = v92
	v324 = v259
	v326 = v95
	v329 = v98
	goto L60
L60:
	;
	F_initStringInfo(m, v312+int32(60))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	v878 = v307
	v879 = v308
	v880 = v309
	v881 = v310
	v883 = v312
	v884 = v872
	v891 = v869
	v892 = v321
	v894 = v323
	v895 = v324
	v897 = v326
	v900 = v329
	goto L57
L62:
	;
	v342 = int32(0)
	v344 = v316
	goto L63
L63:
	;
	v359 = v312 + int32(60)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	v361 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v360))) = uint8(v361)
	*(*int32)(unsafe.Add(mBase, uint32(v359)+12)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v359)+4)) = v361
	goto L65
L64:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v312)+60))
	F_pfree(m, v855)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L147
	}
L65:
	;
	v376 = v344
	goto L67
L66:
	;
	v401 = int32(0)
	v409 = v390
	v411 = v401
	v414 = v393
	v415 = v400
	v417 = v401
	v420 = v401
	goto L73
L67:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	v391 = int32(1)
	v393 = v376 + v391
	switch v390 {
	case 0:
		v400 = v390
		goto L66
	default:
		goto L70
	case 9, 13, 32:
		v395 = v391
		goto L69
	}
L68:
	;
	v400 = int32(0)
	goto L66
L69:
	;
	if v390 == int32(44) {
		v376 = v393
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v395 = int32(0)
	goto L69
L71:
	;
	if v395 != 0 {
		v376 = v393
		goto L67
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	v429 = v409 & int32(255)
	switch v429 {
	case 0:
		v522 = v414
		v533 = v401
		goto L75
	default:
		goto L79
	case 9, 13, 32:
		goto L80
	}
L74:
	;
	v536 = int32(1)
	v537 = v522 - v536
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v312)+64))
	if (v420|base.B2i32(int32(0) < v538))&v536 != 0 {
		goto L97
	} else {
		goto L98
	}
L75:
	;
	goto L74
L76:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	v409 = v510
	v411 = v505
	v414 = v414 + int32(1)
	v415 = v507
	v417 = v506
	v420 = v509
	goto L73
L77:
	;
	v499 = int32(1)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v312)+64))
	if v501 != 0 {
		goto L94
	} else {
		goto L95
	}
L78:
	;
	F_appendStringInfoChar(m, v312+int32(60), base.I32_extend8_s(v409))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L92
	}
L79:
	;
	if (base.B2i32(v429 != int32(35))|v411)&int32(1) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	if v411&int32(1) != 0 {
		v483 = int32(0)
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v522 = v414
	v533 = v401
	goto L75
L82:
	;
	v449 = v414
	goto L85
L83:
	;
	goto L84
L84:
	;
	if (base.B2i32(v429 != int32(44))|v411)&int32(1) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	v465 = v449 + int32(1)
	if v463 != 0 {
		v449 = v465
		goto L85
	} else {
		goto L87
	}
L86:
	;
	v522 = v465
	v533 = v401
	goto L75
L87:
	;
	goto L86
L88:
	;
	v522 = v414
	v533 = int32(1)
	goto L75
L89:
	;
	goto L90
L90:
	;
	if (base.B2i32(v429 != int32(34))|v417)&int32(1) == int32(0) {
		v498 = v411
		goto L77
	} else {
		goto L91
	}
L91:
	;
	v483 = base.B2i32(v429 == int32(34))
	goto L78
L92:
	;
	v489 = int32(0)
	if v483 == v489 {
		v505 = v411
		v506 = v489
		v507 = v415
		v509 = v420
		goto L76
	} else {
		goto L93
	}
L93:
	;
	v498 = v411 & v483 & (v417 ^ int32(1))
	goto L77
L94:
	;
	v502 = v415
	goto L96
L95:
	;
	v502 = v499
	goto L96
L96:
	;
	v505 = v411 ^ int32(1)
	v506 = v498
	v507 = v502
	v509 = v499
	goto L76
L97:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v312)+60))
	v546 = v415 & int32(1)
	if v546 != 0 {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	v839 = v342
	goto L99
L99:
	;
	goto L64
L100:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v312)+20))
	if v533&base.B2i32(v828 == int32(0)) != 0 {
		v342 = v812
		v344 = v537
		goto L63
	} else {
		goto L146
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+76)) = int32(0)
	v581 = F_AbsoluteConfigLocation(m, v544+int32(1), v307)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L112
	}
L102:
	;
	v552 = int32(4554128)
	v553 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v556 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v556
	v558 = F_strlen(m, v544)
	mBase = m.M
	v561 = F_palloc0(m, v558+int32(13))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L106
	}
L103:
	;
	if v538 < int32(2) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544))))
	if v549 == int32(64) {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	goto L102
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v561)+8)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v561)+4)) = uint8(v546)
	v567 = v561 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v561))) = v567
	v570 = v558 + int32(1)
	if v570 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v573 = F_lappend(m, v342, v561)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L111
	}
L108:
	;
	v571 = F__emscripten_memcpy_bulkmem(m, v567, v544, v570)
	mBase = m.M
	goto L110
L109:
	;
	goto L110
L110:
	;
	goto L107
L111:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v553
	v812 = v573
	goto L100
L112:
	;
	v585 = F_open_auth_file(m, v581, v310, v321, v312+int32(20))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	if v585 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_pfree(m, v581)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	F_tokenize_auth_file(m, v581, v585, v312+int32(76), v310, v321)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L118
	}
L117:
	;
	v812 = v342
	goto L100
L118:
	;
	F_pfree(m, v581)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v312)+76))
	if v597 == int32(0) {
		v780 = v342
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v796 = F_FreeFile(m, v585)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L143
	}
L121:
	;
	v600 = int32(0)
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	if v601 <= v600 {
		v780 = v342
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v611 = v342
	v619 = v600
	goto L123
L123:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v597)+12))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v627+v619<<(uint(int32(2))%32))))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v631)+16))
	if v632 != 0 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v780 = v753
	goto L120
L125:
	;
	v633 = F_pstrdup(m, v632)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v631)))
	if v636 == int32(0) {
		v753 = v611
		goto L129
	} else {
		goto L130
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+20)) = v633
	v780 = v611
	goto L120
L129:
	;
	v770 = v619 + int32(1)
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v597)+4))
	if v770 < v771 {
		v611 = v753
		v619 = v770
		goto L123
	} else {
		goto L142
	}
L130:
	;
	v639 = int32(0)
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v636)+4))
	if v640 <= v639 {
		v753 = v611
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v650 = v611
	v653 = v639
	goto L132
L132:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v636)+12))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v666+v653<<(uint(int32(2))%32))))
	if v670 == int32(0) {
		v726 = v650
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v753 = v726
	goto L129
L134:
	;
	v743 = v653 + int32(1)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v636)+4))
	if v743 < v744 {
		v650 = v726
		v653 = v743
		goto L132
	} else {
		goto L141
	}
L135:
	;
	v673 = int32(0)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v670)+4))
	if v674 <= v673 {
		v726 = v650
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v681 = v673
	v684 = v650
	goto L137
L137:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v670)+12))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v700+v681<<(uint(int32(2))%32))))
	v705 = int32(4554128)
	v706 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v709 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v709
	v711 = F_lappend(m, v684, v704)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L139
	}
L138:
	;
	v726 = v711
	goto L134
L139:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v706
	v716 = v681 + int32(1)
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v670)+4))
	if v716 < v717 {
		v681 = v716
		v684 = v711
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	goto L133
L142:
	;
	goto L124
L143:
	;
	if v321 != 0 {
		v812 = v780
		goto L100
	} else {
		goto L144
	}
L144:
	;
	v799 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	F_MemoryContextDelete(m, v799)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, _consts[444])) = int32(0)
	v812 = v780
	goto L100
L146:
	;
	v839 = v812
	goto L99
L147:
	;
	if v839 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v858 = int32(4554128)
	v859 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v862 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v862
	v864 = F_lappend(m, v320, v839)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L151
	}
L149:
	;
	v869 = v320
	goto L150
L150:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v312)+20))
	v871 = int32(0)
	v872 = base.B2i32(v870 == v871)
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537))))
	if v873 == v871 {
		v878 = v307
		v879 = v308
		v880 = v309
		v881 = v310
		v883 = v312
		v884 = v872
		v891 = v869
		v892 = v321
		v894 = v323
		v895 = v324
		v897 = v326
		v900 = v329
		goto L57
	} else {
		goto L152
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v859
	v869 = v864
	goto L150
L152:
	;
	if v870 == int32(0) {
		v316 = v537
		v320 = v869
		goto L60
	} else {
		goto L153
	}
L153:
	;
	goto L61
L154:
	;
	v1230 = v894 + v895 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v883)+28)) = v1230
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v879)+76))
	if v1232 < int32(0) {
		goto L233
	} else {
		goto L234
	}
L155:
	;
	v902 = v891
	goto L157
L156:
	;
	v902 = int32(1)
	goto L157
L157:
	;
	if v902 == int32(0) {
		goto L154
	} else {
		goto L158
	}
L158:
	;
	if v891 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v1175 = int32(4554128)
	v1176 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v1179 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v1179
	v1182 = F_palloc0(m, int32(20))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L223
	}
L160:
	;
	if v884^int32(1) != 0 {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v891)+4))
	if v909 != int32(2) {
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v891)+12))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v912)+4))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v913)+12))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v914)))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v912)))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v916)+12))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v917)))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v918)))
	v920 = int32(429927)
	v923 = int32(*(*uint8)(unsafe.Add(mBase, _consts[445])))
	v924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919))))
	if v924 == int32(0) {
		v943 = v923
		v944 = v924
		goto L164
	} else {
		goto L165
	}
L163:
	;
	if v944-v943 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L164:
	;
	goto L163
L165:
	;
	if v923 != v924 {
		v943 = v923
		v944 = v924
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v928 = v919
	v929 = v920
	goto L167
L167:
	;
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v929)+1)))
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v928)+1)))
	if v933 == int32(0) {
		v943 = v932
		v944 = v933
		goto L164
	} else {
		goto L169
	}
L168:
	;
	v943 = v932
	v944 = v933
	goto L164
L169:
	;
	v936 = int32(1)
	if v932 == v933 {
		v928 = v928 + v936
		v929 = v929 + v936
		goto L167
	} else {
		goto L170
	}
L170:
	;
	goto L168
L171:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v915)))
	F_tokenize_include_file(m, v878, v948, v880, v881, v892, int32(0), v883+int32(20))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v955 = int32(224090)
	v958 = int32(*(*uint8)(unsafe.Add(mBase, _consts[446])))
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919))))
	if v959 == int32(0) {
		v978 = v958
		v979 = v959
		goto L179
	} else {
		goto L180
	}
L174:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v883)+20))
	if v954 != 0 {
		goto L159
	} else {
		goto L175
	}
L175:
	;
	goto L154
L176:
	;
	F_pfree(m, v988)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L221
	}
L177:
	;
	v1093 = v1051
	goto L217
L178:
	;
	if v979-v978 == int32(0) {
		goto L186
	} else {
		goto L187
	}
L179:
	;
	goto L178
L180:
	;
	if v958 != v959 {
		v978 = v958
		v979 = v959
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v963 = v919
	v964 = v955
	goto L182
L182:
	;
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964)+1)))
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963)+1)))
	if v968 == int32(0) {
		v978 = v967
		v979 = v968
		goto L179
	} else {
		goto L184
	}
L183:
	;
	v978 = v967
	v979 = v968
	goto L179
L184:
	;
	v971 = int32(1)
	if v967 == v968 {
		v963 = v963 + v971
		v964 = v964 + v971
		goto L182
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v915)))
	v988 = F_GetConfFilesInDir(m, v983, v878, v881, v883+int32(76), v883+int32(20))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v1054 = int32(123353)
	v1057 = int32(*(*uint8)(unsafe.Add(mBase, _consts[447])))
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919))))
	if v1058 == int32(0) {
		v1077 = v1057
		v1078 = v1058
		goto L207
	} else {
		goto L208
	}
L189:
	;
	if v988 == int32(0) {
		goto L159
	} else {
		goto L190
	}
L190:
	;
	F_initStringInfo(m, v883+int32(60))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v996 = int32(0)
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v883)+76))
	if v997 <= v996 {
		goto L176
	} else {
		goto L192
	}
L192:
	;
	v1004 = v996
	goto L193
L193:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v988+v1004<<(uint(int32(2))%32))))
	F_tokenize_include_file(m, v878, v1026, v880, v881, v892, int32(0), v883+int32(20))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L1
	} else {
		goto L195
	}
L194:
	;
	v1051 = int32(0)
	if v1051 < v1049 {
		goto L177
	} else {
		goto L205
	}
L195:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v883)+20))
	if v1032 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v883)+64))
	if int32(0) < v1035 {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	goto L198
L198:
	;
	v1048 = v1004 + int32(1)
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v883)+76))
	if v1048 < v1049 {
		v1004 = v1048
		goto L193
	} else {
		goto L204
	}
L199:
	;
	F_appendStringInfoChar(m, v883+int32(60), int32(10))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L202
	}
L200:
	;
	v1044 = v1032
	goto L201
L201:
	;
	F_appendStringInfoString(m, v883+int32(60), v1044)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L203
	}
L202:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v883)+20))
	v1044 = v1043
	goto L201
L203:
	;
	goto L198
L204:
	;
	goto L194
L205:
	;
	goto L176
L206:
	;
	if v1078-v1077 != 0 {
		goto L159
	} else {
		goto L214
	}
L207:
	;
	goto L206
L208:
	;
	if v1057 != v1058 {
		v1077 = v1057
		v1078 = v1058
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v1062 = v919
	v1063 = v1054
	goto L210
L210:
	;
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1063)+1)))
	v1067 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1062)+1)))
	if v1067 == int32(0) {
		v1077 = v1066
		v1078 = v1067
		goto L207
	} else {
		goto L212
	}
L211:
	;
	v1077 = v1066
	v1078 = v1067
	goto L207
L212:
	;
	v1070 = int32(1)
	if v1066 == v1067 {
		v1062 = v1062 + v1070
		v1063 = v1063 + v1070
		goto L210
	} else {
		goto L213
	}
L213:
	;
	goto L211
L214:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v915)))
	F_tokenize_include_file(m, v878, v1080, v880, v881, v892, int32(1), v883+int32(20))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v883)+20))
	if v1086 == int32(0) {
		goto L154
	} else {
		goto L216
	}
L216:
	;
	goto L159
L217:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v988+v1093<<(uint(int32(2))%32))))
	F_pfree(m, v1115)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L1
	} else {
		goto L219
	}
L218:
	;
	goto L176
L219:
	;
	v1119 = v1093 + int32(1)
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v883)+76))
	if v1119 < v1120 {
		v1093 = v1119
		goto L217
	} else {
		goto L220
	}
L220:
	;
	goto L218
L221:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v883)+64))
	if v1147 == int32(0) {
		goto L154
	} else {
		goto L222
	}
L222:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v883)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v883)+20)) = v1150
	goto L159
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1182))) = v891
	v1185 = F_pstrdup(m, v878)
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1182)+8)) = v894
	*(*int32)(unsafe.Add(mBase, uint32(v1182)+4)) = v1185
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v883)+44))
	v1190 = F_pstrdup(m, v1189)
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1182)+12)) = v1190
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v883)+20))
	if v1193 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v1194 = F_pstrdup(m, v1193)
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L1
	} else {
		goto L229
	}
L227:
	;
	v1197 = int32(0)
	goto L228
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1182)+16)) = v1197
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v880)))
	v1200 = F_lappend(m, v1199, v1182)
	mBase = m.M
	v1201 = m.ExcPending
	if v1201 != 0 {
		goto L1
	} else {
		goto L230
	}
L229:
	;
	v1197 = v1194
	goto L228
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v880))) = v1200
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v1176
	goto L154
L231:
	;
	if int32(base.Ui32(v1237)>>(uint(int32(4))%32))&int32(1) == int32(0) {
		v76 = v878
		v77 = v879
		v78 = v880
		v79 = v881
		v81 = v883
		v90 = v892
		v92 = v1230
		v95 = v897
		v98 = v900
		goto L14
	} else {
		goto L236
	}
L232:
	;
	goto L231
L233:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v879)))
	v1237 = v1235
	goto L232
L234:
	;
	goto L235
L235:
	;
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v879)))
	v1237 = v1236
	goto L232
L236:
	;
	goto L15
L237:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1249)+32))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v1272
	m.G0 = v1249 + int32(80)
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
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
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v253
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v263 == int32(0) {
		goto L68
	} else {
		goto L69
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
	v93 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v93
	if l2 == v93 {
		v146 = int32(1)
		goto L34
	} else {
		goto L35
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v23 = v22
	goto L7
L6:
	;
	v23 = v6
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v25 = v24
	goto L10
L9:
	;
	v25 = v6
	goto L10
L10:
	;
	v26 = v23 - v25
	v27 = F_list_copy_tail(m, l2, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v29 = int32(0)
	if l2 == v29 {
		v37 = v29
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v37
	v46 = int32(0)
	v47 = v6
	v49 = int32(1)
	v50 = v6
	goto L20
L14:
	;
	goto L13
L15:
	;
	if v26 <= int32(0) {
		v37 = v29
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v26 < v34 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v26
	goto L19
L18:
	;
	goto L19
L19:
	;
	v37 = l2
	goto L14
L20:
	;
	v53 = int32(0)
	if v27 == v53 {
		v63 = v53
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v57 <= v47 {
		v63 = int32(0)
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v63 = v59 + v47<<(uint(int32(2))%32)
	goto L22
L25:
	;
	v253 = v46
	v257 = int32(0)
	v258 = v6
	goto L1
L26:
	;
	goto L27
L27:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v67 <= v47 {
		v253 = v46
		v257 = v50
		v258 = v6
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v63 == int32(0) {
		v253 = v46
		v257 = v50
		v258 = v6
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v74 = v71 + v47<<(uint(int32(2))%32)
	if v74 == int32(0) {
		v253 = v46
		v257 = v50
		v258 = v6
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v80 = int32(0)
	v82 = F_makeTargetEntry(m, v78, base.I32_extend16_s(v49), v80, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L11
	} else {
		goto L31
	}
L31:
	;
	v84 = F_lappend(m, v46, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v84
	v87 = int32(1)
	v91 = F_addTargetToSortList(m, l0, v82, v50, v84, v77)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	v46 = v84
	v47 = v47 + v87
	v49 = v49 + v87
	v50 = v91
	goto L20
L34:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v146
	v152 = F_transformSortClause(m, l0, l3, v15+int32(12), int32(1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L11
	} else {
		goto L42
	}
L35:
	;
	v98 = int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v99 <= int32(0) {
		v146 = v98
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v107 = v98
	v108 = v6
	v110 = v6
	goto L37
L37:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v108<<(uint(int32(2))%32))))
	v120 = int32(0)
	v122 = F_makeTargetEntry(m, v118, base.I32_extend16_s(v107), v120, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L11
	} else {
		goto L39
	}
L38:
	;
	v146 = base.I32_extend16_s(v128)
	goto L34
L39:
	;
	v124 = F_lappend(m, v110, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v124
	v127 = int32(1)
	v128 = v107 + v127
	v130 = v108 + v127
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v130 < v131 {
		v107 = v128
		v108 = v130
		v110 = v124
		goto L37
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	if l4 == int32(0) {
		v244 = v6
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v147
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v253 = v247
	v257 = v152
	v258 = v244
	goto L1
L44:
	;
	v159 = F_transformDistinctClause(m, l0, v15+int32(12), v152, int32(1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L11
	} else {
		goto L45
	}
L45:
	;
	if v159 == int32(0) {
		v244 = v6
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v163 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v244 = v159
	goto L43
L48:
	;
	v166 = int32(0)
	if v166 < v163 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v170 = v163
	goto L51
L50:
	;
	v170 = v166
	goto L51
L51:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v177 = v166
	goto L52
L52:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v171+v177<<(uint(int32(2))%32))))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	if v188 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v193 = F_get_sortgroupclause_expr(m, v187, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L11
	} else {
		goto L58
	}
L54:
	;
	v190 = v177 + int32(1)
	if v170 != v190 {
		v177 = v190
		goto L52
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	goto L47
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	v202 = F_exprType(m, v193)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	v204 = F_format_type_be(m, v202)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v204
	F_errmsg(m, int32(199802), v15)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	F_errdetail(m, int32(610075), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v214 = F_exprLocation(m, v193)
	mBase = m.M
	F_parser_errposition(m, l0, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L11
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(521728), int32(221), int32(319092))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	if v316 == int32(0) {
		v365 = v317
		goto L80
	} else {
		goto L81
	}
L68:
	;
	v316 = v253
	v317 = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v267 = int32(0)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v268 <= v267 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v316 = v310
	v317 = v304
	goto L67
L72:
	;
	v304 = int32(0)
	goto L71
L73:
	;
	goto L74
L74:
	;
	v278 = v267
	v279 = int32(0)
	goto L75
L75:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v285+v278<<(uint(int32(2))%32))))
	v290 = F_exprType(m, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L11
	} else {
		goto L77
	}
L76:
	;
	v304 = v292
	goto L71
L77:
	;
	v292 = F_lappend_oid(m, v279, v290)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	v295 = v278 + int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v295 < v296 {
		v278 = v295
		v279 = v292
		goto L75
	} else {
		goto L79
	}
L79:
	;
	goto L76
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v365
	F_check_agglevels_and_constraints(m, l0, l1)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L11
	} else {
		goto L91
	}
L81:
	;
	v325 = int32(0)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v316)+4))
	if v326 <= v325 {
		v365 = v317
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v331 = v325
	v335 = v317
	goto L83
L83:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v316)+12))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341+v331<<(uint(int32(2))%32))))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+26)))
	if v346 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v365 = v354
	goto L80
L85:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v350 = F_exprType(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L11
	} else {
		goto L88
	}
L86:
	;
	v354 = v335
	goto L87
L87:
	;
	v356 = v331 + int32(1)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v316)+4))
	if v356 < v357 {
		v331 = v356
		v335 = v354
		goto L83
	} else {
		goto L90
	}
L88:
	;
	v352 = F_lappend_oid(m, v335, v350)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L11
	} else {
		goto L89
	}
L89:
	;
	v354 = v352
	goto L87
L90:
	;
	goto L84
L91:
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
												F_errmsg(m, int32(202346), v15+int32(16))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(644826), int32(0))
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
															F_errfinish(m, int32(515748), int32(596), int32(217855))
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
													F_errmsg(m, int32(202346), v15+int32(16))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errhint(m, int32(644826), int32(0))
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
																F_errfinish(m, int32(515748), int32(596), int32(217855))
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
													F_errmsg(m, int32(202346), v15+int32(16))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errhint(m, int32(644826), int32(0))
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
																F_errfinish(m, int32(515748), int32(596), int32(217855))
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
									F_errmsg(m, int32(544983), int32(0))
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
											F_errfinish(m, int32(515748), int32(512), int32(217855))
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
									F_errmsg(m, int32(545022), int32(0))
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
											F_errfinish(m, int32(515748), int32(517), int32(217855))
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
				F_errmsg(m, int32(743695), v15)
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
						F_errfinish(m, int32(515748), int32(485), int32(217855))
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
						if l3 == v57 {
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
														F_errmsg(m, int32(193254), v19)
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
																F_errfinish(m, int32(515748), int32(1004), int32(126066))
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
					if l3 == v57 {
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
													F_errmsg(m, int32(193254), v19)
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
															F_errfinish(m, int32(515748), int32(1004), int32(126066))
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int64
	_ = v221
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
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
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v395 int32
	_ = v395
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
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
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L166
	}
L4:
	;
	m.G0 = v20 + int32(144)
	return v634
L5:
	;
	v525 = F_make_parsestate(m, l0)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L140
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L135
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L127
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
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L122
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
	if l2 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v200 = F_transformSetOperationTree(m, l0, v196, int32(0), v20+int32(88))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L49
	}
L23:
	;
	if v44&int32(1) == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v20)+92))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v63 == int32(142) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v72 = v53
	goto L28
L26:
	;
	v93 = v53
	goto L27
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v105+v106<<(uint(int32(2))%32)-int32(4))))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+36))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+76))
	v123 = int32(0)
	v125 = v5
	v127 = int32(1)
	goto L31
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v84 == int32(142) {
		v72 = v83
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v93 = v83
	goto L27
L30:
	;
	goto L29
L31:
	;
	v134 = int32(0)
	if v62 == v134 {
		v144 = v134
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L22
L33:
	;
	goto L32
L34:
	;
	if v114 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v138 <= v123 {
		v144 = int32(0)
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v144 = v140 + v123<<(uint(int32(2))%32)
	goto L34
L37:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v166 = F_pstrdup(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L46
	}
L38:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	F_analyzeCTETargetList(m, l0, v158, v156)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L45
	}
L39:
	;
	v156 = int32(0)
	goto L38
L40:
	;
	goto L41
L41:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v148 <= v123 {
		v156 = v125
		goto L38
	} else {
		goto L42
	}
L42:
	;
	if v144 == int32(0) {
		v156 = v125
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v155 = v152 + v123<<(uint(int32(2))%32)
	if v155 != 0 {
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v156 = v125
	goto L38
L45:
	;
	goto L33
L46:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	v171 = F_makeTargetEntry(m, v168, base.I32_extend16_s(v127), v166, int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v173 = F_lappend(m, v125, v171)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v123 = v123 + int32(1)
	v125 = v173
	v127 = v127 + int32(1)
	goto L31
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v200
	if v45 == int32(2) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v206 = int32(546087)
	goto L52
L51:
	;
	v206 = int32(542429)
	goto L52
L52:
	;
	if v45 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v209 = int32(554723)
	goto L55
L54:
	;
	v209 = v206
	goto L55
L55:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v20)+92))
	if v210 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v213 = v211
	goto L58
L57:
	;
	v213 = int32(0)
	goto L58
L58:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	if v214 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	v217 = v215
	goto L61
L60:
	;
	v217 = int32(0)
	goto L61
L61:
	;
	if v213 != v217 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	if l3 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v221 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+20)) = v221
	*(*int64)(unsafe.Add(mBase, uint32(v38)+28)) = v221
	v232 = int32(0)
	goto L66
L66:
	;
	v243 = int32(0)
	if v210 == v243 {
		v253 = v243
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if v214 == int32(0) {
		v634 = v38
		goto L4
	} else {
		goto L71
	}
L69:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v247 <= v232 {
		v253 = int32(0)
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v253 = v249 + v232<<(uint(int32(2))%32)
	goto L68
L71:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v256 <= v232 {
		v634 = v38
		goto L4
	} else {
		goto L72
	}
L72:
	;
	if v253 == int32(0) {
		v634 = v38
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v263 = v260 + v232<<(uint(int32(2))%32)
	if v263 == int32(0) {
		v634 = v38
		goto L4
	} else {
		goto L74
	}
L74:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	v270 = F_exprType(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v272 = F_exprType(m, v267)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v267
	v282 = F_list_make2_impl(m, v20+int32(28), v20+int32(24))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v286 = F_select_common_type(m, l0, v282, v209, v20+int32(84))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	v289 = F_exprLocation(m, v288)
	mBase = m.M
	if v270 != int32(705) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v272 != int32(705) {
		goto L87
	} else {
		goto L88
	}
L80:
	;
	v292 = F_coerce_to_common_type(m, l0, v269, v286, v209)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	if base.Ui32(int32(1)) < base.Ui32(v294-int32(7)) {
		v302 = v269
		goto L79
	} else {
		goto L84
	}
L83:
	;
	v302 = v292
	goto L79
L84:
	;
	v299 = F_coerce_to_common_type(m, l0, v269, v286, v209)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v299
	v302 = v299
	goto L79
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v315
	v324 = F_list_make2_impl(m, v20+int32(20), v20+int32(16))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L93
	}
L87:
	;
	v305 = F_coerce_to_common_type(m, l0, v267, v286, v209)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if base.Ui32(int32(1)) < base.Ui32(v307-int32(7)) {
		v315 = v267
		goto L86
	} else {
		goto L91
	}
L90:
	;
	v315 = v305
	goto L86
L91:
	;
	v312 = F_coerce_to_common_type(m, l0, v267, v286, v209)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v266)+4)) = v312
	v315 = v312
	goto L86
L93:
	;
	v326 = F_select_common_typmod(m, v324, v286)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v315
	v336 = F_list_make2_impl(m, v20+int32(12), v20+int32(8))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v338 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+8)))
	v343 = v341
	goto L98
L97:
	;
	v343 = int32(0)
	goto L98
L98:
	;
	v346 = F_select_common_collation(m, l0, v336, v343&int32(1))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v349 = F_lappend_oid(m, v348, v286)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v349
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v353 = F_lappend_int(m, v352, v326)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v353
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v357 = F_lappend_oid(m, v356, v346)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = v357
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v360 == int32(1) {
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
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+8)))
	if v363 != 0 {
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v365 = v20 + int32(96)
	*(*int32)(unsafe.Add(mBase, uint32(v365)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v365)+4)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v365)+16)) = v365
	v371 = int32(4547032)
	v372 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	*(*int32)(unsafe.Add(mBase, uint32(v365)+8)) = v372
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v20 + int32(104)
	goto L108
L107:
	;
	goto L106
L108:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	v380 = F_palloc0(m, int32(20))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = int32(106)
	v384 = int32(0)
	F_get_sort_group_operators(m, v286, v384, int32(1), v384, v20+int32(140), v20+int32(136), v384, v20+int32(135))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v44&int32(1) == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v407 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v380)+4)) = v407
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v20)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+8)) = v409
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v20)+140))
	*(*uint16)(unsafe.Add(mBase, uint32(v380)+16)) = uint16(v407)
	*(*int32)(unsafe.Add(mBase, uint32(v380)+12)) = v411
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+135)))
	*(*uint8)(unsafe.Add(mBase, uint32(v380)+18)) = uint8(v415)
	v417 = F_lappend(m, v378, v380)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L114
	}
L112:
	;
	if base.B2i32(v286 != int32(2287))&base.B2i32(v286 != int32(2249)) != 0 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v405 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+135)) = uint8(v405)
	goto L111
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v417
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v20+int32(96))+8))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v423
	goto L115
L115:
	;
	goto L103
L116:
	;
	v429 = F_palloc0(m, int32(20))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v232 = v232 + int32(1)
	goto L66
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+16)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v429)+12)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v429)+8)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v429)+4)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = int32(57)
	v437 = int32(0)
	v440 = F_makeTargetEntry(m, v429, v437, v437, v437)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v443 = F_lappend(m, v442, v440)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v443
	goto L118
L122:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errmsg(m, int32(542309), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v461 = F_exprLocation(m, v460)
	mBase = m.M
	F_parser_errposition(m, l0, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(521846), int32(2066), int32(429666))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+12))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+8))
	v483 = v479 - int32(1)
	if base.Ui32(v483) <= base.Ui32(int32(3)) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v491
	F_errmsg(m, int32(542263), v20+int32(48))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L133
	}
L130:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v483<<(uint(int32(2))%32))+uint32(_consts[338])))
	v491 = v490
	goto L132
L131:
	;
	v491 = int32(392056)
	goto L132
L132:
	;
	goto L129
L133:
	;
	F_errfinish(m, int32(521846), int32(2076), int32(429666))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v209
	F_errmsg(m, int32(156851), v20+int32(32))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v20)+88))
	v517 = F_exprLocation(m, v516)
	mBase = m.M
	F_parser_errposition(m, l0, v517)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(521846), int32(2226), int32(429666))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
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
	v527 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v525)+84)) = uint16(v527)
	*(*int32)(unsafe.Add(mBase, uint32(v525)+44)) = v527
	v531 = F_transformStmt(m, v525, l1)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_free_parsestate(m, v525)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v535 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v537 = F_contain_vars_of_level(m, v531, int32(1))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	if l3 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	if v537 != 0 {
		goto L3
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v600 != 0 {
		goto L159
	} else {
		goto L160
	}
L149:
	;
	v541 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v541
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v531)+76))
	if v544 == v541 {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v544)+4))
	if v547 <= int32(0) {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	v551 = v541
	v557 = v5
	goto L152
L152:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v544)+12))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v567+v551<<(uint(int32(2))%32))))
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571)+26)))
	if v572 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L148
L154:
	;
	v575 = F_lappend(m, v557, v571)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L157
	}
L155:
	;
	v578 = v557
	goto L156
L156:
	;
	v580 = v551 + int32(1)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v544)+4))
	if v580 < v581 {
		v551 = v580
		v557 = v578
		goto L152
	} else {
		goto L158
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v575
	v578 = v575
	goto L156
L158:
	;
	goto L153
L159:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v600)+4))
	v605 = v601 + int32(1)
	goto L161
L160:
	;
	v605 = int32(1)
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v605
	v611 = F_pg_snprintf(m, v20+int32(96), int32(32), int32(509842), v20)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v616 = F_makeAlias(m, v20+int32(96), int32(0))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v618 = int32(0)
	v620 = F_addRangeTableEntryForSubquery(m, l0, v531, v616, v618, v618)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v623 = F_palloc0(m, int32(8))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623))) = int32(63)
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v620)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+4)) = v627
	v634 = v623
	goto L4
L166:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	F_errmsg(m, int32(320194), int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v662 = F_locate_var_of_level(m, v531, int32(1))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_parser_errposition(m, l0, v662)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(521846), int32(2138), int32(429666))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
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
	var v28 int32
	_ = v28
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
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int64
	_ = v312
	var v314 int32
	_ = v314
	var v316 int64
	_ = v316
	var v318 int32
	_ = v318
	var v320 int64
	_ = v320
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v710 int32
	_ = v710
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 == v3 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	m.G0 = v14 + int32(80)
	return v710
L2:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v659 = F_list_copy(m, v658)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L38
	} else {
		goto L153
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L38
	} else {
		goto L148
	}
L4:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v137 == int32(0) {
		goto L2
	} else {
		goto L34
	}
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v28 = v3
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v34 = int32(2)
	v36 = v33 + v28<<(uint(v34)%32)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v39 = v36 + int32(4)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if base.Ui32(v39) < base.Ui32(v42+v43<<(uint(v34)%32)) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	v48 = v39
	goto L11
L10:
	;
	v48 = int32(0)
	goto L11
L11:
	;
	if v48 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v52 = (v48 - v42) >> (uint(int32(2)) % 32)
	goto L14
L13:
	;
	v52 = v43
	goto L14
L14:
	;
	if v52 < v43 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v57 = v52
	goto L18
L16:
	;
	goto L17
L17:
	;
	v112 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+36)) = v112
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+32)) = uint8(v112)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v117 != int32(141) {
		goto L30
	} else {
		goto L31
	}
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v42+v57<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v74 == int32(0) {
		v93 = v73
		v94 = v74
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L17
L20:
	;
	if v94-v93 == int32(0) {
		goto L3
	} else {
		goto L28
	}
L21:
	;
	goto L20
L22:
	;
	if v73 != v74 {
		v93 = v73
		v94 = v74
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v78 = v54
	v79 = v70
	goto L24
L24:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v83 == int32(0) {
		v93 = v82
		v94 = v83
		goto L21
	} else {
		goto L26
	}
L25:
	;
	v93 = v82
	v94 = v83
	goto L21
L26:
	;
	v86 = int32(1)
	if v82 == v83 {
		v78 = v78 + v86
		v79 = v79 + v86
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v99 = v57 + int32(1)
	if v99 != v43 {
		v57 = v99
		goto L18
	} else {
		goto L29
	}
L29:
	;
	goto L19
L30:
	;
	v120 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)) = uint8(v120)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v123 = v28 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v123 < v124 {
		v28 = v123
		goto L7
	} else {
		goto L33
	}
L33:
	;
	goto L8
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l0
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v141 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v144 = v142
	goto L37
L36:
	;
	v144 = int32(0)
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v144
	v148 = F_palloc0(m, v144*int32(12))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return int32(0)
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v148
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v153 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v154 = int32(0)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v154 < v155 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v209 = v144
	goto L42
L42:
	;
	if v209 <= int32(0) {
		goto L1
	} else {
		goto L49
	}
L43:
	;
	v160 = v154
	goto L46
L44:
	;
	goto L45
L45:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v209 = v197
	goto L42
L46:
	;
	v170 = v160 * int32(12)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173+v160<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v170+v171))) = v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v179+v170)+4)) = v160
	v183 = v160 + int32(1)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v183 < v184 {
		v160 = v183
		goto L46
	} else {
		goto L48
	}
L47:
	;
	goto L45
L48:
	;
	goto L47
L49:
	;
	v215 = int32(0)
	goto L50
L50:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v224+v215*int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v215
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)+16))
	v235 = F_makeDependencyGraphWalker(m, v232, v14+int32(36))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L38
	} else {
		goto L52
	}
L51:
	;
	v241 = int32(0)
	if v239 <= v241 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	v238 = v215 + int32(1)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v238 < v239 {
		v215 = v238
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v252 = v241
	goto L55
L55:
	;
	v259 = v252
	goto L58
L56:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v369 <= int32(0) {
		goto L1
	} else {
		goto L78
	}
L57:
	;
	if v259 != v252 {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	v270 = v244 + v259*int32(12)
	v272 = v270 + int32(8)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	if v273 == int32(0) {
		goto L57
	} else {
		goto L60
	}
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L38
	} else {
		goto L62
	}
L60:
	;
	v277 = v259 + int32(1)
	if v277 != v239 {
		v259 = v277
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L38
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(465666), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L38
	} else {
		goto L64
	}
L64:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v244+v252*int32(12))))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+28))
	F_parser_errposition(m, v245, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L38
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(521976), int32(883), int32(87408))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L38
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v304 = v14 + int32(72)
	v307 = v244 + v252*int32(12)
	v309 = v307 + int32(8)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	*(*int32)(unsafe.Add(mBase, uint32(v304))) = v310
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v307)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	*(*int32)(unsafe.Add(mBase, uint32(v309))) = v314
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v270)))
	*(*int64)(unsafe.Add(mBase, uint32(v307))) = v316
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v304)))
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = v318
	v320 = *(*int64)(unsafe.Add(mBase, uint32(v14)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v270))) = v320
	goto L69
L68:
	;
	goto L69
L69:
	;
	v326 = v252 + int32(1)
	if v326 < v239 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v335 = v326
	goto L73
L71:
	;
	goto L72
L72:
	;
	if v239 != v326 {
		v252 = v326
		goto L55
	} else {
		goto L77
	}
L73:
	;
	v348 = v244 + v335*int32(12) + int32(8)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v244+v252*int32(12)+int32(4))))
	v351 = F_bms_del_member(m, v349, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L38
	} else {
		goto L75
	}
L74:
	;
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v348))) = v351
	v355 = v335 + int32(1)
	if v355 != v239 {
		v335 = v355
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	goto L56
L78:
	;
	v375 = v369
	v377 = int32(0)
	goto L86
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L38
	} else {
		goto L145
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L38
	} else {
		goto L140
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L38
	} else {
		goto L135
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L38
	} else {
		goto L130
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L38
	} else {
		goto L125
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L38
	} else {
		goto L120
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L38
	} else {
		goto L115
	}
L86:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v384+v377*int32(12))))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388)+32)))
	if v389 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v438 <= int32(0) {
		goto L1
	} else {
		goto L105
	}
L88:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v388)+16))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	if v393 != int32(141) {
		goto L85
	} else {
		goto L91
	}
L89:
	;
	v438 = v375
	goto L90
L90:
	;
	v441 = v377 + int32(1)
	if v441 < v438 {
		v375 = v438
		v377 = v441
		goto L86
	} else {
		goto L104
	}
L91:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v392)+68))
	if v396 != int32(1) {
		goto L84
	} else {
		goto L92
	}
L92:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v392)+64))
	if v399 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v377
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	v408 = F_checkWellFormedRecursionWalker(m, v405, v14+int32(36))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L38
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v392)+44))
	if v410 != 0 {
		goto L83
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v392)+48))
	if v411 != 0 {
		goto L82
	} else {
		goto L98
	}
L98:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v392)+52))
	if v412 != 0 {
		goto L81
	} else {
		goto L99
	}
L99:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v392)+60))
	if v413 != 0 {
		goto L80
	} else {
		goto L100
	}
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = int32(1)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v392)+76))
	v422 = F_checkWellFormedRecursionWalker(m, v419, v14+int32(36))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L38
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+52)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v377
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v392)+80))
	v432 = F_checkWellFormedRecursionWalker(m, v429, v14+int32(36))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L38
	} else {
		goto L102
	}
L102:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	if v434 != int32(1) {
		goto L79
	} else {
		goto L103
	}
L103:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v438 = v437
	goto L90
L104:
	;
	goto L87
L105:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v449 = int32(0)
	v451 = v445
	goto L106
L106:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v458+v449*int32(12))))
	v463 = F_lappend(m, v451, v462)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L38
	} else {
		goto L108
	}
L107:
	;
	v470 = int32(0)
	if v468 <= v470 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v463
	v467 = v449 + int32(1)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v467 < v468 {
		v449 = v467
		v451 = v463
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v475 = v470
	goto L111
L111:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v484+v475*int32(12))))
	F_analyzeCTE(m, l0, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L38
	} else {
		goto L113
	}
L112:
	;
	goto L1
L113:
	;
	v492 = v475 + int32(1)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v492 < v493 {
		v475 = v492
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L38
	} else {
		goto L116
	}
L116:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v502
	F_errmsg(m, int32(130666), v14+int32(16))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L38
	} else {
		goto L117
	}
L117:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v388)+28))
	F_parser_errposition(m, v509, v510)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L38
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(521976), int32(936), int32(283363))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L38
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
	F_errcode(m, int32(151388292))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L38
	} else {
		goto L121
	}
L121:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v525
	F_errmsg(m, int32(301359), v14)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L38
	} else {
		goto L122
	}
L122:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v388)+28))
	F_parser_errposition(m, v530, v531)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L38
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(521976), int32(944), int32(283363))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L38
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
	v545 = m.ExcPending
	if v545 != 0 {
		goto L38
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(465421), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L38
	} else {
		goto L127
	}
L127:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v392)+44))
	v552 = F_exprLocation(m, v551)
	mBase = m.M
	F_parser_errposition(m, v550, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L38
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(521976), int32(979), int32(283363))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L38
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
	v566 = m.ExcPending
	if v566 != 0 {
		goto L38
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(465516), int32(0))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L38
	} else {
		goto L132
	}
L132:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v392)+48))
	v573 = F_exprLocation(m, v572)
	mBase = m.M
	F_parser_errposition(m, v571, v573)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L38
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(521976), int32(985), int32(283363))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L38
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
	v587 = m.ExcPending
	if v587 != 0 {
		goto L38
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(465470), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L38
	} else {
		goto L137
	}
L137:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v392)+52))
	v594 = F_exprLocation(m, v593)
	mBase = m.M
	F_parser_errposition(m, v592, v594)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L38
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(521976), int32(991), int32(283363))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L38
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
	v608 = m.ExcPending
	if v608 != 0 {
		goto L38
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(465563), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L38
	} else {
		goto L142
	}
L142:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v392)+60))
	v615 = F_exprLocation(m, v614)
	mBase = m.M
	F_parser_errposition(m, v613, v615)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L38
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(521976), int32(997), int32(283363))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L38
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
	F_errmsg_internal(m, int32(435942), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L38
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(521976), int32(1019), int32(283363))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L38
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L38
	} else {
		goto L149
	}
L149:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v643
	F_errmsg(m, int32(434642), v14+int32(32))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L38
	} else {
		goto L150
	}
L150:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
	F_parser_errposition(m, l0, v650)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L38
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(521976), int32(139), int32(377021))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L38
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v659
	v662 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v662 == int32(0) {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v662)+4))
	if v665 <= int32(0) {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v671 = int32(0)
	goto L156
L156:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v662)+12))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v680+v671<<(uint(int32(2))%32))))
	F_analyzeCTE(m, l0, v684)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L38
	} else {
		goto L158
	}
L157:
	;
	goto L1
L158:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v688 = F_lappend(m, v687, v684)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L38
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v688
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v692 = F_list_delete_first(m, v691)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L38
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v692
	v696 = v671 + int32(1)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v662)+4))
	if v696 < v697 {
		v671 = v696
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
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_palloc0(m, int32(36))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(440)
		v16 = int32(20)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v16
		v23 = F_list_make1_impl(m, int32(472), v7+int32(8))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(7626)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(7627)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(7628)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(7629)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(7630)
			v43 = int32(256)
			*(*uint16)(unsafe.Add(mBase, uint32(v10)+8)) = uint16(v43)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v23
			m.G0 = v7 + int32(16)
			return v10
		}
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
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
	v29 = v24
	v30 = v20
	v32 = v2
	v33 = int32(0)
	goto L6
L4:
	;
	v76 = v2
	goto L5
L5:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v86 = F_palloc(m, v83<<(uint(int32(2))%32))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L17
	}
L6:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v39 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v76 = v67
	goto L5
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v44 = v42 & int32(4095)
	v47 = F_palloc(m, v44+int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v65 = v30
	v67 = v32
	goto L10
L10:
	;
	v71 = v33 + int32(1)
	if v71 < v65 {
		v29 = v29 + int32(12)
		v30 = v65
		v32 = v67
		v33 = v71
		goto L6
	} else {
		goto L16
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	if v44 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v56 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54+v44))) = uint8(v56)
	*(*int32)(unsafe.Add(mBase, uint32(v16+v32<<(uint(int32(2))%32)))) = v54
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v65 = v64
	v67 = v32 + int32(1)
	goto L10
L13:
	;
	v53 = F__emscripten_memcpy_bulkmem(m, v47, v24+v13*int32(12)+int32(base.Ui32(v49)>>(uint(int32(12))%32)), v44)
	mBase = m.M
	v54 = v53
	goto L15
L14:
	;
	v54 = v47
	goto L15
L15:
	;
	goto L12
L16:
	;
	goto L7
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v88 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v92 = v11 + int32(8)
	v97 = v92
	v98 = int32(0)
	v99 = v88
	v102 = v2
	goto L21
L19:
	;
	v146 = v2
	goto L20
L20:
	;
	F_pg_qsort(m, v16, v76, int32(4), int32(1529))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L32
	}
L21:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v107 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v146 = v134
	goto L20
L23:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v112 = v110 & int32(4095)
	v115 = F_palloc(m, v112+int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v133 = v99
	v134 = v102
	goto L25
L25:
	;
	v139 = v98 + int32(1)
	if v139 < v133 {
		v97 = v97 + int32(12)
		v98 = v139
		v99 = v133
		v102 = v134
		goto L21
	} else {
		goto L31
	}
L26:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	if v112 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v122+v112))) = uint8(v124)
	*(*int32)(unsafe.Add(mBase, uint32(v86+v102<<(uint(int32(2))%32)))) = v122
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v133 = v132
	v134 = v102 + int32(1)
	goto L25
L28:
	;
	v121 = F__emscripten_memcpy_bulkmem(m, v115, v92+v83*int32(12)+int32(base.Ui32(v117)>>(uint(int32(12))%32)), v112)
	mBase = m.M
	v122 = v121
	goto L30
L29:
	;
	v122 = v115
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L22
L32:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v76) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v159 = int32(1)
	v160 = int32(0)
	goto L36
L34:
	;
	v221 = v76
	goto L35
L35:
	;
	F_pg_qsort(m, v86, v146, int32(4), int32(1529))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L50
	}
L36:
	;
	v169 = int32(2)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v16+v159<<(uint(v169)%32))))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v16+v160<<(uint(v169)%32))))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v180 == int32(0) {
		v199 = v179
		v200 = v180
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v221 = v211 + int32(1)
	goto L35
L38:
	;
	v214 = v159 + int32(1)
	if v214 != v76 {
		v159 = v214
		v160 = v211
		goto L36
	} else {
		goto L49
	}
L39:
	;
	if v200-v199 == int32(0) {
		v211 = v160
		goto L38
	} else {
		goto L47
	}
L40:
	;
	goto L39
L41:
	;
	if v179 != v180 {
		v199 = v179
		v200 = v180
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v184 = v172
	v185 = v176
	goto L43
L43:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+1)))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+1)))
	if v189 == int32(0) {
		v199 = v188
		v200 = v189
		goto L40
	} else {
		goto L45
	}
L44:
	;
	v199 = v188
	v200 = v189
	goto L40
L45:
	;
	v192 = int32(1)
	if v188 == v189 {
		v184 = v184 + v192
		v185 = v185 + v192
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v205 = v160 + int32(1)
	if v159 == v205 {
		v211 = v159
		goto L38
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v205<<(uint(int32(2))%32)))) = v172
	v211 = v205
	goto L38
L49:
	;
	goto L37
L50:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v146) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v236 = int32(1)
	v237 = int32(0)
	goto L54
L52:
	;
	v300 = v146
	goto L53
L53:
	;
	if v221 < v300 {
		goto L69
	} else {
		goto L70
	}
L54:
	;
	v246 = int32(2)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v86+v236<<(uint(v246)%32))))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v86+v237<<(uint(v246)%32))))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v257 == int32(0) {
		v276 = v256
		v277 = v257
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v300 = v288 + int32(1)
	goto L53
L56:
	;
	v291 = v236 + int32(1)
	if v291 != v146 {
		v236 = v291
		v237 = v288
		goto L54
	} else {
		goto L67
	}
L57:
	;
	if v277-v276 == int32(0) {
		v288 = v237
		goto L56
	} else {
		goto L65
	}
L58:
	;
	goto L57
L59:
	;
	if v256 != v257 {
		v276 = v256
		v277 = v257
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v261 = v249
	v262 = v253
	goto L61
L61:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	if v266 == int32(0) {
		v276 = v265
		v277 = v266
		goto L58
	} else {
		goto L63
	}
L62:
	;
	v276 = v265
	v277 = v266
	goto L58
L63:
	;
	v269 = int32(1)
	if v265 == v266 {
		v261 = v261 + v269
		v262 = v262 + v269
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v282 = v237 + int32(1)
	if v236 == v282 {
		v288 = v236
		goto L56
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86+v282<<(uint(int32(2))%32)))) = v249
	v288 = v282
	goto L56
L67:
	;
	goto L55
L68:
	;
	return v398
L69:
	;
	v398 = int32(0)
	goto L68
L70:
	;
	v306 = int32(0)
	if v300 <= v306 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	return int32(1)
L72:
	;
	goto L73
L73:
	;
	v312 = v306
	v314 = int32(0)
	goto L74
L74:
	;
	if v221 <= v312 {
		v371 = v312
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v398 = v382
	goto L68
L76:
	;
	if v371 == v221 {
		goto L69
	} else {
		goto L90
	}
L77:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v86+v314<<(uint(int32(2))%32))))
	v327 = v312
	goto L78
L78:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v16+v327<<(uint(int32(2))%32))))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340))))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	if v344 == int32(0) {
		v363 = v343
		v364 = v344
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L69
L80:
	;
	if v364-v363 == int32(0) {
		v371 = v327
		goto L76
	} else {
		goto L88
	}
L81:
	;
	goto L80
L82:
	;
	if v343 != v344 {
		v363 = v343
		v364 = v344
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v348 = v326
	v349 = v340
	goto L84
L84:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+1)))
	if v353 == int32(0) {
		v363 = v352
		v364 = v353
		goto L81
	} else {
		goto L86
	}
L85:
	;
	v363 = v352
	v364 = v353
	goto L81
L86:
	;
	v356 = int32(1)
	if v352 == v353 {
		v348 = v348 + v356
		v349 = v349 + v356
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v369 = v327 + int32(1)
	if v369 != v221 {
		v327 = v369
		goto L78
	} else {
		goto L89
	}
L89:
	;
	goto L79
L90:
	;
	v382 = int32(1)
	v384 = v314 + v382
	if v300 != v384 {
		v312 = v371
		v314 = v384
		goto L74
	} else {
		goto L91
	}
L91:
	;
	goto L75
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
	var v26 int32
	_ = v26
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
	var v71 int32
	_ = v71
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
	v26 = v3
	goto L9
L7:
	;
	v71 = v3
	goto L8
L8:
	;
	v78 = int32(4554128)
	v79 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v81
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
		v60 = v26
		v61 = v37
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v71 = v60
	goto L8
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v62+v34))) = v61
	v66 = v25 + int32(1)
	if v66 != v12 {
		v25 = v66
		v26 = v60
		goto L9
	} else {
		goto L17
	}
L12:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+4)))
	if v42 != int32(65535) {
		v60 = v26
		v61 = v37
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v25))))
	if v47 != 0 {
		v60 = v26
		v61 = v37
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v48 != int32(1) {
		v60 = v26
		v61 = v37
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
	*(*int32)(unsafe.Add(mBase, uint32(v53+v26<<(uint(int32(2))%32)))) = v51
	v60 = v26 + int32(1)
	v61 = v51
	goto L11
L17:
	;
	goto L10
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v79
	if int32(0) < v71 {
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
	if v111 != v71 {
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
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
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
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
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
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var __phi260 int32
	_ = __phi260
	var v261 int32
	_ = v261
	var __phi261 int32
	_ = __phi261
	var v269 int32
	_ = v269
	var __phi269 int32
	_ = __phi269
	var v275 int32
	_ = v275
	var __phi275 int32
	_ = __phi275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v416 int32
	_ = v416
	var v427 int32
	_ = v427
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v456 int64
	_ = v456
	var v458 int32
	_ = v458
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
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
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v650 int32
	_ = v650
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v809 int32
	_ = v809
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
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
		goto L8
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L161
	}
L6:
	;
	m.G0 = v20 - int32(-64)
	return v809
L7:
	;
	F_close_tsvector_parser(m, v26)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L51
	}
L8:
	;
	if v45 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v227 = v2
	v231 = v31
	v233 = v34
	goto L7
L10:
	;
	goto L11
L11:
	;
	v51 = int32(256)
	v55 = v34
	v58 = v2
	v60 = int32(64)
	v62 = v31
	v64 = v34
	goto L12
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	if int32(2047) <= v67 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v227 = v207
	v231 = v124
	v233 = v166
	goto L7
L14:
	;
	v70 = int32(0)
	v71 = F_errsave_start(m, v25)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v90 = v55 - v64
	if int32(1048576) <= v90 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	if v71 == int32(0) {
		v809 = v70
		goto L6
	} else {
		goto L18
	}
L18:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(2046)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v80
	F_errmsg(m, int32(703616), v20)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errsave_finish(m, v25, int32(517964), int32(215), int32(288090))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v809 = v70
	goto L6
L22:
	;
	v93 = int32(0)
	v94 = F_errsave_start(m, v25)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v60 <= v58 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	if v94 == int32(0) {
		v809 = v93
		goto L6
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(1048575)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v90
	F_errmsg(m, int32(703557), v18+int32(-48))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errsave_finish(m, v25, int32(517964), int32(221), int32(288090))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v809 = v93
	goto L6
L30:
	;
	v117 = F_repalloc(m, v62, v60*int32(24))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	v122 = v67
	v123 = v60
	v124 = v62
	goto L32
L32:
	;
	if v51 <= v90+v122 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v122 = v121
	v123 = v60 << (uint(int32(1)) % 32)
	v124 = v117
	goto L32
L34:
	;
	v128 = v51
	v141 = v64
	goto L37
L35:
	;
	v153 = v51
	v155 = v122
	v157 = v55
	v166 = v64
	goto L36
L36:
	;
	v169 = int32(12)
	v171 = v124 + v58*v169
	v172 = int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v155<<(uint(v172)%32)&int32(4094) | v176&v172 | v90<<(uint(v169)%32)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	if v185 != 0 {
		goto L42
	} else {
		goto L43
	}
L37:
	;
	v145 = v128 << (uint(int32(1)) % 32)
	v146 = F_repalloc(m, v141, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v153 = v145
	v155 = v148
	v157 = v90 + v146
	v166 = v146
	goto L36
L39:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	if v145 <= v90+v148 {
		v128 = v145
		v141 = v146
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	if v190 != 0 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v186 = F__emscripten_memcpy_bulkmem(m, v157, v184, v185)
	mBase = m.M
	v187 = v186
	goto L44
L43:
	;
	v187 = v157
	goto L44
L44:
	;
	goto L41
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171)+8)) = v203
	v207 = v58 + int32(1)
	v217 = F_gettoken_tsvector(m, v26, v18+int32(-4), v18+int32(-8), v18+int32(-12), v18+int32(-16), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v189 | int32(1)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v171)+4)) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v203 = v196
	goto L45
L47:
	;
	goto L48
L48:
	;
	v197 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v171)+4)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v171))) = v189 & int32(-2)
	v203 = v197
	goto L45
L49:
	;
	if v217 != 0 {
		v51 = v153
		v55 = v188 + v187
		v58 = v207
		v60 = v123
		v62 = v124
		v64 = v166
		goto L12
	} else {
		goto L50
	}
L50:
	;
	goto L13
L51:
	;
	if v25 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v249 = int32(0)
	if v227 <= v249 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v240 != int32(447) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+4)))
	if v243 != int32(1) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v246 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v246)
	v809 = int32(0)
	goto L6
L56:
	;
	v704 = v694 << (uint(int32(2)) % 32)
	v707 = v704 + v695 + int32(8)
	v708 = F_palloc0(m, v707)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L143
	}
L57:
	;
	v694 = v227
	v695 = v2
	goto L56
L58:
	;
	goto L59
L59:
	;
	if v227 == int32(1) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	v541 = int32(1)
	v545 = int32(base.Ui32(v540)>>(uint(v541)%32))&int32(2047) + v538
	if v540&v541 != 0 {
		goto L119
	} else {
		goto L120
	}
L61:
	;
	v523 = v231
	v538 = v2
	goto L60
L62:
	;
	goto L63
L63:
	;
	F_qsort_arg(m, v231, v227, int32(12), int32(1535), v233)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	__phi260 = v231
	__phi261 = v231 + int32(12)
	__phi269 = v231
	__phi275 = v2
	v260 = __phi260
	v261 = __phi261
	v269 = __phi269
	v275 = __phi275
	goto L65
L65:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	v278 = int32(1)
	v280 = int32(2047)
	v281 = int32(base.Ui32(v277)>>(uint(v278)%32)) & v280
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v286 = int32(base.Ui32(v282)>>(uint(v278)%32)) & v280
	if v281 == v286 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v523 = v500
	v538 = v515
	goto L60
L67:
	;
	v517 = int32(12)
	v518 = v261 + v517
	v521 = base.I32_div_s(v518-v231, v517)
	if v521 < v227 {
		__phi260 = v500
		__phi261 = v518
		__phi269 = v261
		__phi275 = v515
		v260 = __phi260
		v261 = __phi261
		v269 = __phi269
		v275 = __phi275
		goto L65
	} else {
		goto L118
	}
L68:
	;
	if v277&int32(1) == int32(0) {
		v500 = v260
		v515 = v275
		goto L67
	} else {
		goto L108
	}
L69:
	;
	v288 = int32(12)
	v290 = v233 + int32(base.Ui32(v277)>>(uint(v288)%32))
	v293 = v233 + int32(base.Ui32(v282)>>(uint(v288)%32))
	if v281 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	goto L71
L71:
	;
	v340 = v286 + v275
	if v282&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L72:
	;
	if v337 == int32(0) {
		goto L68
	} else {
		goto L86
	}
L73:
	;
	v337 = int32(0)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	if v299 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v300 = v290
	v301 = v293
	v302 = v281
	v303 = v299
	goto L80
L77:
	;
	v325 = v293
	v329 = int32(0)
	goto L78
L78:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
	v337 = v329 - v330
	goto L72
L79:
	;
	v325 = v320
	v329 = v322
	goto L78
L80:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v303 != v305 {
		v320 = v301
		v322 = v303
		goto L79
	} else {
		goto L82
	}
L81:
	;
	v320 = v314
	v322 = int32(0)
	goto L79
L82:
	;
	if v305 == int32(0) {
		v320 = v301
		v322 = v303
		goto L79
	} else {
		goto L83
	}
L83:
	;
	v310 = v302 - int32(1)
	if v310 == int32(0) {
		v320 = v301
		v322 = v303
		goto L79
	} else {
		goto L84
	}
L84:
	;
	v313 = int32(1)
	v314 = v301 + v313
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+1)))
	if v315 != 0 {
		v300 = v300 + v313
		v301 = v314
		v302 = v310
		v303 = v315
		goto L80
	} else {
		goto L85
	}
L85:
	;
	goto L81
L86:
	;
	goto L71
L87:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	if int32(2) <= v343 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v451 = v340
	goto L89
L89:
	;
	v454 = v260 + int32(12)
	if v260 == v269 {
		goto L105
	} else {
		goto L106
	}
L90:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	F_pg_qsort(m, v346, v343, int32(2), int32(1536))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	v416 = v343
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+8)) = v416
	v427 = int32(1)
	v451 = (v340+v427)&int32(-2) + v416<<(uint(v427)%32) + int32(2)
	goto L89
L93:
	;
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v346))))
	v355 = v346 + int32(2)
	v356 = v353
	v360 = v346
	goto L94
L94:
	;
	v371 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v355))))
	v372 = int32(16383)
	v373 = v371 & v372
	if v373 != v356&v372 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v416 = (v403 - v346 + int32(2)) >> (uint(int32(1)) % 32)
	goto L92
L96:
	;
	goto L95
L97:
	;
	v396 = v355 + int32(2)
	if (v396-v346)>>(uint(int32(1))%32) < v343 {
		v355 = v396
		v356 = v393
		v360 = v394
		goto L94
	} else {
		goto L104
	}
L98:
	;
	v378 = v360 + int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v378))) = uint16(v371)
	if int32(508) < v378-v346 {
		v403 = v378
		goto L96
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v385 = int32(14)
	if base.Ui32(int32(base.Ui32(v371)>>(uint(v385)%32))) <= base.Ui32(int32(base.Ui32(v356&int32(65535))>>(uint(v385)%32))) {
		v393 = v356
		v394 = v360
		goto L97
	} else {
		goto L103
	}
L101:
	;
	if v373 != int32(16383) {
		v393 = v371
		v394 = v378
		goto L97
	} else {
		goto L102
	}
L102:
	;
	v403 = v378
	goto L96
L103:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v360))) = uint16(v371)
	v393 = v371
	v394 = v360
	goto L97
L104:
	;
	v403 = v394
	goto L96
L105:
	;
	v500 = v454
	v515 = v451
	goto L67
L106:
	;
	goto L107
L107:
	;
	v456 = *(*int64)(unsafe.Add(mBase, uint32(v261)))
	*(*int64)(unsafe.Add(mBase, uint32(v454))) = v456
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v454)+8)) = v458
	v500 = v454
	v515 = v451
	goto L67
L108:
	;
	if v282&int32(1) != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	v469 = v269 + int32(20)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v471 = v467 + v470
	v474 = F_repalloc(m, v466, v471<<(uint(int32(1))%32))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v282 | int32(1)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v269)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v260)+4)) = v496
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v269)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v260)+8)) = v498
	v500 = v260
	v515 = v275
	goto L67
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+4)) = v474
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	v478 = int32(1)
	v482 = v269 + int32(16)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v486 = v484 << (uint(v478) % 32)
	if v486 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+8)) = v471
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v482)))
	F_pfree(m, v490)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L117
	}
L114:
	;
	v487 = F__emscripten_memcpy_bulkmem(m, v474+v477<<(uint(v478)%32), v483, v486)
	mBase = m.M
	goto L116
L115:
	;
	goto L116
L116:
	;
	goto L113
L117:
	;
	v500 = v260
	v515 = v275
	goto L67
L118:
	;
	goto L66
L119:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v523)+8))
	if int32(2) <= v548 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v650 = v545
	goto L121
L121:
	;
	v659 = int32(12)
	v662 = base.I32_div_s(v523-v231+v659, v659)
	if v650 < int32(1048576) {
		v694 = v662
		v695 = v650
		goto L56
	} else {
		goto L137
	}
L122:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v523)+4))
	F_pg_qsort(m, v551, v548, int32(2), int32(1536))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L1
	} else {
		goto L125
	}
L123:
	;
	v625 = v548
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v523)+8)) = v625
	v632 = int32(1)
	v650 = (v545+v632)&int32(-2) + v625<<(uint(v632)%32) + int32(2)
	goto L121
L125:
	;
	v558 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551))))
	v560 = v551 + int32(2)
	v561 = v558
	v565 = v551
	goto L126
L126:
	;
	v576 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v560))))
	v577 = int32(16383)
	v578 = v576 & v577
	if v578 != v561&v577 {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	v625 = (v608 - v551 + int32(2)) >> (uint(int32(1)) % 32)
	goto L124
L128:
	;
	goto L127
L129:
	;
	v601 = v560 + int32(2)
	if (v601-v551)>>(uint(int32(1))%32) < v548 {
		v560 = v601
		v561 = v598
		v565 = v599
		goto L126
	} else {
		goto L136
	}
L130:
	;
	v583 = v565 + int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v583))) = uint16(v576)
	if int32(508) < v583-v551 {
		v608 = v583
		goto L128
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v590 = int32(14)
	if base.Ui32(int32(base.Ui32(v576)>>(uint(v590)%32))) <= base.Ui32(int32(base.Ui32(v561&int32(65535))>>(uint(v590)%32))) {
		v598 = v561
		v599 = v565
		goto L129
	} else {
		goto L135
	}
L133:
	;
	if v578 != int32(16383) {
		v598 = v576
		v599 = v583
		goto L129
	} else {
		goto L134
	}
L134:
	;
	v608 = v583
	goto L128
L135:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v565))) = uint16(v576)
	v598 = v576
	v599 = v565
	goto L129
L136:
	;
	v608 = v599
	goto L128
L137:
	;
	v665 = int32(0)
	v666 = F_errsave_start(m, v25)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	if v666 == int32(0) {
		v809 = v665
		goto L6
	} else {
		goto L139
	}
L139:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(1048575)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v650
	F_errmsg(m, int32(703781), v18+int32(-32))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errsave_finish(m, v25, int32(517964), int32(274), int32(288090))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v809 = v665
	goto L6
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v708)+4)) = v694
	*(*int32)(unsafe.Add(mBase, uint32(v708))) = v707 << (uint(int32(2)) % 32)
	if v694 <= int32(0) {
		v809 = v708
		goto L6
	} else {
		goto L144
	}
L144:
	;
	v717 = v708 + int32(8)
	v718 = v717 + v704
	v723 = int32(0)
	v730 = v249
	goto L145
L145:
	;
	v738 = int32(12)
	v740 = v231 + v723*v738
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	v748 = int32(base.Ui32(v741)>>(uint(int32(1))%32)) & int32(2047)
	if v748 != 0 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v809 = v708
	goto L6
L147:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	v756 = v751&int32(4095) | v730<<(uint(int32(12))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v740))) = v756
	v758 = int32(1)
	v762 = int32(base.Ui32(v751)>>(uint(v758)%32))&int32(2047) + v730
	if v751&v758 != 0 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v749 = F__emscripten_memcpy_bulkmem(m, v718+v730, v233+int32(base.Ui32(v741)>>(uint(v738)%32)), v748)
	mBase = m.M
	goto L150
L149:
	;
	goto L150
L150:
	;
	goto L147
L151:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v740)+8))
	if int32(65536) <= v768 {
		goto L5
	} else {
		goto L154
	}
L152:
	;
	v796 = v762
	v797 = v756
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v717+v723<<(uint(int32(2))%32)))) = v797
	v800 = v723 + int32(1)
	if v800 != v694 {
		v723 = v800
		v730 = v796
		goto L145
	} else {
		goto L160
	}
L154:
	;
	v771 = int32(1)
	v774 = (v762 + v771) & int32(-2)
	*(*uint16)(unsafe.Add(mBase, uint32(v718+v774))) = uint16(v768)
	v778 = v774 + int32(2)
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v740)+4))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v740)+8))
	v783 = v781 << (uint(v771) % 32)
	if v783 != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v740)+8))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v740)+4))
	F_pfree(m, v787)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L159
	}
L156:
	;
	v784 = F__emscripten_memcpy_bulkmem(m, v718+v778, v780, v783)
	mBase = m.M
	goto L158
L157:
	;
	goto L158
L158:
	;
	goto L155
L159:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	v796 = v786<<(uint(int32(1))%32) + v778
	v797 = v793
	goto L153
L160:
	;
	goto L146
L161:
	;
	F_errmsg_internal(m, int32(343702), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(517964), int32(292), int32(288090))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v11 == v2 {
		v59 = v2
		m.G0 = v8 + int32(48)
		return v59
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v11+v14*int32(12))))
		if v18 == int32(0) {
			v59 = v2
			m.G0 = v8 + int32(48)
			return v59
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
			v25 = F_pg_sprintf(m, v8+int32(16), int32(509851), v8)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v8 + int32(16)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v36 = v32 + v33*int32(12)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v39
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v44 = F_BuildTupleFromCStrings(m, v41, v8+int32(36))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
					v47 = F_HeapTupleHeaderGetDatum(m, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
						F_pfree(m, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
							F_pfree(m, v52)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v55 + int32(1)
								v59 = v47
								m.G0 = v8 + int32(48)
								return v59
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
			v14 = int32(4554128)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v17
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
								F_errmsg_internal(m, int32(384849), int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									F_errfinish(m, int32(518285), int32(69), int32(318667))
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
								*(*int32)(unsafe.Add(mBase, _consts[3])) = v15
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
				F_errmsg_internal(m, int32(47262), v9)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					F_errfinish(m, int32(518285), int32(57), int32(318667))
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v59 int64
	_ = v59
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v86 int64
	_ = v86
	var v94 int32
	_ = v94
	var v101 float64
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v191 int32
	_ = v191
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int64
	_ = v378
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int64
	_ = v428
	var v430 int32
	_ = v430
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v463 int64
	_ = v463
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v514 int32
	_ = v514
	var v530 int32
	_ = v530
	var v540 int32
	_ = v540
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	goto L1
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v36) <= base.Ui32(v35) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L1
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v540)
	m.G0 = v18 + int32(16)
	return v530
L5:
	;
	v530 = v514
	v540 = int32(0)
	goto L4
L6:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v501 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v500 + v501
	*(*int32)(unsafe.Add(mBase, uint32(v490)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v490))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v490)+4)) = v501
	v514 = v490
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L23
	} else {
		goto L105
	}
L8:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v38 == int64(4294967296) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v276 = int32(0)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v279 = v278 & l1
	v282 = v277 + v279*int32(12)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v283 == v276 {
		v490 = v282
		goto L6
	} else {
		goto L64
	}
L11:
	;
	v41 = int32(0)
	v43 = int64(2)
	v45 = v38 << (uint(int64(1)) % 64)
	if base.Ui64(v45) <= base.Ui64(v43) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L10
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L23
	} else {
		goto L61
	}
L14:
	;
	v48 = v43
	goto L16
L15:
	;
	v48 = v45
	goto L16
L16:
	;
	v49 = int64(1)
	if v48&(v48-v49) == int64(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v59 = v48
	goto L19
L18:
	;
	v59 = v49 << (uint(int64(64)-base.I64_clz(v48)) % 64)
	goto L19
L19:
	;
	if base.Ui64(v59*int64(12)) < base.Ui64(int64(2147483647)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v71 = F_MemoryContextAllocExtended(m, v66, base.I32_wrap_i64(v59)*int32(12), int32(5))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L23
	} else {
		goto L58
	}
L23:
	;
	return int32(0)
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v71
	v76 = int64(1)
	if v59&(v59-v76) == int64(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v86 = v59
	goto L27
L26:
	;
	v86 = v76 << (uint(int64(64)-base.I64_clz(v59)) % 64)
	goto L27
L27:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v86*int64(12)) {
		goto L13
	} else {
		goto L28
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v86
	v94 = base.I32_wrap_i64(v86) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v94
	v101 = base.F64_mul(base.F64_convert_i64_u(v86), float64(0.9))
	if base.F64_lt(v101, float64(4.294967296e+09))&base.F64_ge(v101, float64(0)) != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v86 == int64(4294967296) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v107 = base.I32_trunc_f64_u(v101)
	v109 = v107
	goto L29
L31:
	;
	goto L32
L32:
	;
	v109 = int32(0)
	goto L29
L33:
	;
	v110 = int32(-85899346)
	goto L35
L34:
	;
	v110 = v109
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v110
	if v65 != int64(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v120 = v41
	goto L40
L37:
	;
	goto L38
L38:
	;
	F_pfree(m, v64)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L23
	} else {
		goto L57
	}
L39:
	;
	v150 = v143
	v151 = v41
	goto L45
L40:
	;
	v131 = v64 + v120*int32(12)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	if v132 != int32(1) {
		v143 = v120
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v143 = int32(0)
	goto L39
L42:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	if v135&v94 == v120 {
		v143 = v120
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v139 = v120 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v139)) < base.Ui64(v65) {
		v120 = v139
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v161 = v64 + v150*int32(12)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v162 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L38
L47:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	v172 = v166
	goto L50
L48:
	;
	goto L49
L49:
	;
	v209 = v150 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v209)) < base.Ui64(v65) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v182 = v172 & v165
	v187 = v71 + v182*int32(12)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v188 != 0 {
		v172 = v182 + int32(1)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v161)))
	*(*int64)(unsafe.Add(mBase, uint32(v187))) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v187)+8)) = v191
	goto L49
L52:
	;
	goto L51
L53:
	;
	v213 = v209
	goto L55
L54:
	;
	v213 = int32(0)
	goto L55
L55:
	;
	v215 = v151 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v215)) < base.Ui64(v65) {
		v150 = v213
		v151 = v215
		goto L45
	} else {
		goto L56
	}
L56:
	;
	goto L46
L57:
	;
	goto L12
L58:
	;
	F_errmsg_internal(m, int32(418921), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(342598), int32(327), int32(357659))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errmsg_internal(m, int32(418921), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L23
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(342598), int32(327), int32(357659))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L23
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v290 = v278
	v291 = v282
	v294 = v276
	v296 = v279
	goto L65
L65:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	if v301 == l1 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v490 = v470
	goto L6
L67:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)+52))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v303)+36))
	v308 = F_ExecStoreMinimalTuple(m, v305, v306, int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L23
	} else {
		goto L70
	}
L68:
	;
	v338 = v301
	v339 = v290
	goto L69
L69:
	;
	v342 = v338 & v339
	if base.Ui32(v296) < base.Ui32(v342) {
		goto L78
	} else {
		goto L79
	}
L70:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v303)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v304)+12)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v304)+8)) = v310
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v303)+48))
	if v313 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v304)+20))
	F_MemoryContextReset(m, v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L23
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v320 = int32(4554128)
	v321 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v304)+20))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v323
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v313)+20))
	v328 = m.T0[v327].(func(*base.Module, int32, int32, int32) int32)(m, v313, v304, v18+int32(15))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L23
	} else {
		goto L75
	}
L74:
	;
	v530 = v291
	v540 = int32(1)
	goto L4
L75:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v321
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v304)+20))
	F_MemoryContextReset(m, v332)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L23
	} else {
		goto L76
	}
L76:
	;
	if v328 != 0 {
		v530 = v291
		v540 = int32(1)
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v338 = v336
	v339 = v337
	goto L69
L78:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v346 = v296 + v344
	goto L80
L79:
	;
	v346 = v296
	goto L80
L80:
	;
	v349 = (v296 + int32(1)) & v339
	if base.Ui32(v346-v342) < base.Ui32(v294) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v355 = v277 + v349*int32(12)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+4))
	if v356 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v458 = v294 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v458) {
		goto L100
	} else {
		goto L101
	}
L84:
	;
	v360 = v349
	v364 = int32(0)
	goto L87
L85:
	;
	v393 = v349
	v398 = v355
	goto L86
L86:
	;
	if v393 != v296 {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	v373 = v364 + int32(1)
	if int32(151) <= v373 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v393 = v385
	v398 = v388
	goto L86
L89:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v378 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v376), base.F64_convert_i64_u(v378)), float64(0.1)) != 0 {
		goto L3
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v385 = (v360 + int32(1)) & v339
	v388 = v277 + v385*int32(12)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	if v389 != 0 {
		v360 = v385
		v364 = v373
		goto L87
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	goto L88
L94:
	;
	v409 = v393
	v414 = v398
	goto L97
L95:
	;
	goto L96
L96:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v449 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v448 + v449
	*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v291))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v449
	v514 = v291
	goto L5
L97:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v424 = v421 & (v409 - int32(1))
	v427 = v277 + v424*int32(12)
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v427)))
	*(*int64)(unsafe.Add(mBase, uint32(v414))) = v428
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v427)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v414)+8)) = v430
	if v424 != v296 {
		v409 = v424
		v414 = v427
		goto L97
	} else {
		goto L99
	}
L98:
	;
	goto L96
L99:
	;
	goto L98
L100:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v463 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v461), base.F64_convert_i64_u(v463)), float64(0.1)) != 0 {
		goto L3
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v470 = v277 + v349*int32(12)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	if v471 != 0 {
		v290 = v339
		v291 = v470
		v294 = v458
		v296 = v349
		goto L65
	} else {
		goto L104
	}
L103:
	;
	goto L102
L104:
	;
	goto L66
L105:
	;
	F_errmsg_internal(m, int32(483103), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L23
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(342598), int32(630), int32(326211))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L23
	} else {
		goto L107
	}
L107:
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
							F_errmsg(m, int32(318262), v8+int32(16))
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
									F_errfinish(m, int32(522308), int32(280), int32(389145))
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
						F_errmsg(m, int32(78176), v8)
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
								F_errfinish(m, int32(522308), int32(274), int32(389145))
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
